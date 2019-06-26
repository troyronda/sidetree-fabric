/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package observer

import (
	"encoding/json"
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/document"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/core/peer"
	gossipapi "github.com/hyperledger/fabric/extensions/gossip/api"
	"github.com/pkg/errors"
	dcasclient "github.com/trustbloc/fabric-peer-ext/pkg/collections/offledger/dcas/client"
	"github.com/trustbloc/fabric-peer-ext/pkg/roles"
	"github.com/trustbloc/sidetree-fabric/pkg/observer/notifier"
)

var logger = flogging.MustGetLogger("observer")

const (
	sidetreeNs   = "sidetreetxn_cc"
	sidetreeColl = "dcas"

	docNs   = "document_cc"
	docColl = "docs"

	observerRole = "observer"

	transactionTime   = "transactionTime"
	transactionNumber = "transactionNumber"
	operationIndex    = "operationIndex"
)

// getBlockPublisher returns block publisher for channel
var getBlockPublisher = func(channelID string) publisher {
	return peer.BlockPublisher.ForChannel(channelID)
}

// publisher allows clients to add handlers for various block events
type publisher interface {
	// AddWriteHandler adds a handler for KV writes
	AddWriteHandler(handler gossipapi.WriteHandler)
}

type dcasClient interface {

	// Get retrieves the DCAS value
	Get(ns, coll, key string) ([]byte, error)

	// Put puts the DCAS value and returns the key for the value
	Put(ns, coll string, value []byte) (string, error)
}

type cfg interface {
	GetChannels() []string
}

// Start starts channel observer routines
func Start(cfg cfg) error {

	if roles.HasRole(observerRole) {

		logger.Infof("peer is an observer, channels to observe: %s", cfg.GetChannels())
		for _, channelID := range cfg.GetChannels() {
			observer := &observer{channelID: channelID}
			observer.Start()
		}

	} else {
		logger.Info("peer is not an observer, nothing to do...")
	}

	return nil
}

type observer struct {
	sync.RWMutex

	channelID string
	dcas      dcasClient
}

func (o *observer) Start() {

	logger.Debugf("initialize observing on channel: %s", o.channelID)

	// register to receive Sidetree transactions from blocks
	n := notifier.New(getBlockPublisher(o.channelID))

	sidetreeTxnChannel := n.RegisterForSidetreeTxn()

	go func(channel string, txnsCh <-chan notifier.SidetreeTxn) {

		for {
			txn, ok := <-txnsCh
			if !ok {
				logger.Warnf("received !ok from channel '%s'", channel)
				return
			}

			err := o.processSidetreeTxn(txn)
			if err != nil {
				logger.Warnf("Failed to process anchor[%s] on channel[%s]: %s", txn.AnchorAddress, channel, err.Error())
				continue
			}

			logger.Debugf("Successfully processed anchor[%s] on channel[%s]", txn.AnchorAddress, channel)

		}
	}(o.channelID, sidetreeTxnChannel)

}

func (o *observer) getDCASClient() dcasClient {

	o.RLock()
	dcas := o.dcas
	o.RUnlock()

	if dcas != nil {
		return dcas
	}

	dcas = getDCAS(o.channelID)

	o.Lock()
	defer o.Unlock()

	o.dcas = dcas
	return dcas
}

func (o *observer) processSidetreeTxn(sidetreeTxn notifier.SidetreeTxn) error {

	logger.Debugf("processing sidetree txn:%+v on channel '%s'", sidetreeTxn, o.channelID)

	content, err := o.getContent(sidetreeTxn.AnchorAddress)
	if err != nil {
		return errors.Wrapf(err, "failed to retrieve content for anchor: channel[%s] sidetreeNs[%s] key[%s]", o.channelID, sidetreeNs, sidetreeTxn.AnchorAddress)
	}

	logger.Debugf("cas content for anchor[%s]: %s", sidetreeTxn.AnchorAddress, string(content))

	af, err := getAnchorFile(content)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal anchor[%s]", sidetreeTxn.AnchorAddress)
	}

	return o.processBatchFile(af.BatchFileHash, sidetreeTxn)
}

func (o *observer) processBatchFile(batchFileAddress string, sidetreeTxn notifier.SidetreeTxn) error {

	content, err := o.getContent(batchFileAddress)
	if err != nil {
		return errors.Wrapf(err, "failed to retrieve content for batch: channel[%s] sidetreeNs[%s] key[%s]", o.channelID, sidetreeNs, batchFileAddress)
	}

	bf, err := getBatchFile(content)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal batch[%s]", batchFileAddress)
	}

	logger.Debugf("batch file operations: %s", bf.Operations)

	for index, op := range bf.Operations {
		updatedOp, err := updateOperation([]byte(op), index, sidetreeTxn)
		if err != nil {
			return errors.Wrapf(err, "failed to update operation with blockchain metadata")
		}

		addr, err := o.putContent(updatedOp)
		if err != nil {
			return errors.Wrapf(err, "failed to store operation[%s] from batch[%s]", op, batchFileAddress)
		}

		logger.Debugf("successfully stored operation[%s] at address[%s]", op, addr)
	}

	return nil
}

func updateOperation(value []byte, index int, sidetreeTxn notifier.SidetreeTxn) ([]byte, error) {

	doc, err := document.FromBytes(value)
	if err != nil {
		return nil, err
	}

	// The following three constants define order of operations (when operation happened in blockchain time):
	// 1) block number
	// 2) transaction number within block (defines anchor/batch file transaction)
	// 3) operation index within batch

	//  The logical blockchain time that this operation was anchored on the blockchain
	doc[transactionTime] = sidetreeTxn.BlockNumber
	// The transaction number of the transaction this operation was batched within
	doc[transactionNumber] = sidetreeTxn.TxNum
	// The index this operation was assigned to in the batch
	doc[operationIndex] = index

	return doc.Bytes()
}

var getDCAS = func(channelID string) dcasClient {
	return dcasclient.New(channelID)
}

func (o *observer) getContent(address string) ([]byte, error) {

	logger.Debugf("retrieving content ns[%s] col[%s] address64[%s] on channel[%s]", sidetreeNs, sidetreeColl, address, o.channelID)

	return o.getDCASClient().Get(sidetreeNs, sidetreeColl, address)
}

func (o *observer) putContent(value []byte) (string, error) {

	logger.Debugf("put content ns[%s] col[%s] value[%s] on channel[%s]", docNs, docColl, string(value), o.channelID)

	return o.getDCASClient().Put(docNs, docColl, value)
}

// AnchorFile defines the schema of a Anchor File and its related operations.
type AnchorFile struct {
	// BatchFileHash is encoded hash of the batch file
	BatchFileHash string `json:"batchFileHash"`
	// MerkleRoot is encoded root hash of the Merkle tree constructed from
	// the operations included in the batch file
	MerkleRoot string `json:"merkleRoot"`
}

// getAnchorFile creates new anchor file struct from bytes
var getAnchorFile = func(bytes []byte) (*AnchorFile, error) {
	return unmarshalAnchorFile(bytes)
}

// unmarshalAnchorFile creates new anchor file struct from bytes
func unmarshalAnchorFile(bytes []byte) (*AnchorFile, error) {

	af := &AnchorFile{}
	err := json.Unmarshal(bytes, af)
	if err != nil {
		return nil, err
	}

	return af, nil
}

// BatchFile defines the schema of a Batch File and its related operations.
type BatchFile struct {
	// operations included in this batch file, each operation is an encoded string
	Operations []string `json:"operations"`
}

// getBatchFile creates new batch file struct from bytes
var getBatchFile = func(bytes []byte) (*BatchFile, error) {
	return unmarshalBatchFile(bytes)
}

// unmarshalBatchFile creates new batch file struct from bytes
func unmarshalBatchFile(bytes []byte) (*BatchFile, error) {

	bf := &BatchFile{}
	err := json.Unmarshal(bytes, bf)
	if err != nil {
		return nil, err
	}
	return bf, nil
}
