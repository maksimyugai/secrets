// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package messager

import (
	"sync"
)

// Ensure, that CrypterMock does implement Crypter.
// If this is not the case, regenerate this file with moq.
var _ Crypter = &CrypterMock{}

// CrypterMock is a mock implementation of Crypter.
//
// 	func TestSomethingThatUsesCrypter(t *testing.T) {
//
// 		// make and configure a mocked Crypter
// 		mockedCrypter := &CrypterMock{
// 			DecryptFunc: func(req Request) ([]byte, error) {
// 				panic("mock out the Decrypt method")
// 			},
// 			EncryptFunc: func(req Request) ([]byte, error) {
// 				panic("mock out the Encrypt method")
// 			},
// 		}
//
// 		// use mockedCrypter in code that requires Crypter
// 		// and then make assertions.
//
// 	}
type CrypterMock struct {
	// DecryptFunc mocks the Decrypt method.
	DecryptFunc func(req Request) ([]byte, error)

	// EncryptFunc mocks the Encrypt method.
	EncryptFunc func(req Request) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// Decrypt holds details about calls to the Decrypt method.
		Decrypt []struct {
			// Req is the req argument value.
			Req Request
		}
		// Encrypt holds details about calls to the Encrypt method.
		Encrypt []struct {
			// Req is the req argument value.
			Req Request
		}
	}
	lockDecrypt sync.RWMutex
	lockEncrypt sync.RWMutex
}

// Decrypt calls DecryptFunc.
func (mock *CrypterMock) Decrypt(req Request) ([]byte, error) {
	if mock.DecryptFunc == nil {
		panic("CrypterMock.DecryptFunc: method is nil but Crypter.Decrypt was just called")
	}
	callInfo := struct {
		Req Request
	}{
		Req: req,
	}
	mock.lockDecrypt.Lock()
	mock.calls.Decrypt = append(mock.calls.Decrypt, callInfo)
	mock.lockDecrypt.Unlock()
	return mock.DecryptFunc(req)
}

// DecryptCalls gets all the calls that were made to Decrypt.
// Check the length with:
//     len(mockedCrypter.DecryptCalls())
func (mock *CrypterMock) DecryptCalls() []struct {
	Req Request
} {
	var calls []struct {
		Req Request
	}
	mock.lockDecrypt.RLock()
	calls = mock.calls.Decrypt
	mock.lockDecrypt.RUnlock()
	return calls
}

// Encrypt calls EncryptFunc.
func (mock *CrypterMock) Encrypt(req Request) ([]byte, error) {
	if mock.EncryptFunc == nil {
		panic("CrypterMock.EncryptFunc: method is nil but Crypter.Encrypt was just called")
	}
	callInfo := struct {
		Req Request
	}{
		Req: req,
	}
	mock.lockEncrypt.Lock()
	mock.calls.Encrypt = append(mock.calls.Encrypt, callInfo)
	mock.lockEncrypt.Unlock()
	return mock.EncryptFunc(req)
}

// EncryptCalls gets all the calls that were made to Encrypt.
// Check the length with:
//     len(mockedCrypter.EncryptCalls())
func (mock *CrypterMock) EncryptCalls() []struct {
	Req Request
} {
	var calls []struct {
		Req Request
	}
	mock.lockEncrypt.RLock()
	calls = mock.calls.Encrypt
	mock.lockEncrypt.RUnlock()
	return calls
}
