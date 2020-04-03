/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package monitor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/trustbloc/sidetree-fabric/pkg/config"
	stmocks "github.com/trustbloc/sidetree-fabric/pkg/mocks"
	obmocks "github.com/trustbloc/sidetree-fabric/pkg/observer/mocks"
)

const (
	sideTreeTxnCCName = "sidetreetxn_cc"
	dcasColl          = "dcas"
)

func TestSidetreeDCASReader_Read(t *testing.T) {
	dcasClient := obmocks.NewMockDCASClient()
	dcasClientProvider := &stmocks.DCASClientProvider{}
	dcasClientProvider.ForChannelReturns(dcasClient, nil)

	dcasCfg := config.DCAS{
		ChaincodeName: sideTreeTxnCCName,
		Collection:    dcasColl,
	}

	r := NewSidetreeDCASReader(channel1, dcasCfg, dcasClientProvider)
	require.NotNil(t, r)

	t.Run("Found", func(t *testing.T) {
		expectedValue := []byte("some value")
		key, err := dcasClient.Put(dcasCfg.ChaincodeName, dcasCfg.Collection, expectedValue)
		require.NoError(t, err)

		value, err := r.Read(key)
		require.NoError(t, err)
		require.Equal(t, expectedValue, value)
	})

	t.Run("Not found", func(t *testing.T) {
		value, err := r.Read("some key")
		require.Error(t, err)
		require.Contains(t, err.Error(), "content not found for key")
		require.Nil(t, value)

		merr, ok := err.(monitorError)
		require.True(t, ok)
		require.False(t, merr.Transient())
	})
}

func TestSidetreeDCASReader_ReadError(t *testing.T) {
	dcasClient := obmocks.NewMockDCASClient()

	dcasClientProvider := &stmocks.DCASClientProvider{}
	dcasClientProvider.ForChannelReturns(dcasClient, nil)

	dcasCfg := config.DCAS{
		ChaincodeName: sideTreeTxnCCName,
		Collection:    dcasColl,
	}

	r := NewSidetreeDCASReader(channel1, dcasCfg, dcasClientProvider)
	require.NotNil(t, r)

	errExpected := errors.New("injected Put error")
	dcasClient.GetErr = errExpected

	value, err := r.Read("some-key")
	require.EqualError(t, err, errExpected.Error())
	require.Nil(t, value)

	merr, ok := err.(monitorError)
	require.True(t, ok)
	require.True(t, merr.Transient())
}
