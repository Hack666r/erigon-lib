// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package txpool

import (
	"sync"
)

// Ensure, that PoolMock does implement Pool.
// If this is not the case, regenerate this file with moq.
var _ Pool = &PoolMock{}

// PoolMock is a mock implementation of Pool.
//
// 	func TestSomethingThatUsesPool(t *testing.T) {
//
// 		// make and configure a mocked Pool
// 		mockedPool := &PoolMock{
// 			AddFunc: func(newTxs TxSlots) error {
// 				panic("mock out the Add method")
// 			},
// 			AddNewGoodPeerFunc: func(peerID PeerID)  {
// 				panic("mock out the AddNewGoodPeer method")
// 			},
// 			GetRlpFunc: func(hash []byte) []byte {
// 				panic("mock out the GetRlp method")
// 			},
// 			IdHashKnownFunc: func(hash []byte) bool {
// 				panic("mock out the IdHashKnown method")
// 			},
// 		}
//
// 		// use mockedPool in code that requires Pool
// 		// and then make assertions.
//
// 	}
type PoolMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(newTxs TxSlots) error

	// AddNewGoodPeerFunc mocks the AddNewGoodPeer method.
	AddNewGoodPeerFunc func(peerID PeerID)

	// GetRlpFunc mocks the GetRlp method.
	GetRlpFunc func(hash []byte) []byte

	// IdHashKnownFunc mocks the IdHashKnown method.
	IdHashKnownFunc func(hash []byte) bool

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// NewTxs is the newTxs argument value.
			NewTxs TxSlots
		}
		// AddNewGoodPeer holds details about calls to the AddNewGoodPeer method.
		AddNewGoodPeer []struct {
			// PeerID is the peerID argument value.
			PeerID PeerID
		}
		// GetRlp holds details about calls to the GetRlp method.
		GetRlp []struct {
			// Hash is the hash argument value.
			Hash []byte
		}
		// IdHashKnown holds details about calls to the IdHashKnown method.
		IdHashKnown []struct {
			// Hash is the hash argument value.
			Hash []byte
		}
	}
	lockAdd            sync.RWMutex
	lockAddNewGoodPeer sync.RWMutex
	lockGetRlp         sync.RWMutex
	lockIdHashKnown    sync.RWMutex
}

// Add calls AddFunc.
func (mock *PoolMock) Add(newTxs TxSlots) error {
	callInfo := struct {
		NewTxs TxSlots
	}{
		NewTxs: newTxs,
	}
	mock.lockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	mock.lockAdd.Unlock()
	if mock.AddFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.AddFunc(newTxs)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedPool.AddCalls())
func (mock *PoolMock) AddCalls() []struct {
	NewTxs TxSlots
} {
	var calls []struct {
		NewTxs TxSlots
	}
	mock.lockAdd.RLock()
	calls = mock.calls.Add
	mock.lockAdd.RUnlock()
	return calls
}

// AddNewGoodPeer calls AddNewGoodPeerFunc.
func (mock *PoolMock) AddNewGoodPeer(peerID PeerID) {
	callInfo := struct {
		PeerID PeerID
	}{
		PeerID: peerID,
	}
	mock.lockAddNewGoodPeer.Lock()
	mock.calls.AddNewGoodPeer = append(mock.calls.AddNewGoodPeer, callInfo)
	mock.lockAddNewGoodPeer.Unlock()
	if mock.AddNewGoodPeerFunc == nil {
		return
	}
	mock.AddNewGoodPeerFunc(peerID)
}

// AddNewGoodPeerCalls gets all the calls that were made to AddNewGoodPeer.
// Check the length with:
//     len(mockedPool.AddNewGoodPeerCalls())
func (mock *PoolMock) AddNewGoodPeerCalls() []struct {
	PeerID PeerID
} {
	var calls []struct {
		PeerID PeerID
	}
	mock.lockAddNewGoodPeer.RLock()
	calls = mock.calls.AddNewGoodPeer
	mock.lockAddNewGoodPeer.RUnlock()
	return calls
}

// GetRlp calls GetRlpFunc.
func (mock *PoolMock) GetRlp(hash []byte) []byte {
	callInfo := struct {
		Hash []byte
	}{
		Hash: hash,
	}
	mock.lockGetRlp.Lock()
	mock.calls.GetRlp = append(mock.calls.GetRlp, callInfo)
	mock.lockGetRlp.Unlock()
	if mock.GetRlpFunc == nil {
		var (
			bytesOut []byte
		)
		return bytesOut
	}
	return mock.GetRlpFunc(hash)
}

// GetRlpCalls gets all the calls that were made to GetRlp.
// Check the length with:
//     len(mockedPool.GetRlpCalls())
func (mock *PoolMock) GetRlpCalls() []struct {
	Hash []byte
} {
	var calls []struct {
		Hash []byte
	}
	mock.lockGetRlp.RLock()
	calls = mock.calls.GetRlp
	mock.lockGetRlp.RUnlock()
	return calls
}

// IdHashKnown calls IdHashKnownFunc.
func (mock *PoolMock) IdHashKnown(hash []byte) bool {
	callInfo := struct {
		Hash []byte
	}{
		Hash: hash,
	}
	mock.lockIdHashKnown.Lock()
	mock.calls.IdHashKnown = append(mock.calls.IdHashKnown, callInfo)
	mock.lockIdHashKnown.Unlock()
	if mock.IdHashKnownFunc == nil {
		var (
			bOut bool
		)
		return bOut
	}
	return mock.IdHashKnownFunc(hash)
}

// IdHashKnownCalls gets all the calls that were made to IdHashKnown.
// Check the length with:
//     len(mockedPool.IdHashKnownCalls())
func (mock *PoolMock) IdHashKnownCalls() []struct {
	Hash []byte
} {
	var calls []struct {
		Hash []byte
	}
	mock.lockIdHashKnown.RLock()
	calls = mock.calls.IdHashKnown
	mock.lockIdHashKnown.RUnlock()
	return calls
}
