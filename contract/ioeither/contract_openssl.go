// Copyright 2023 IBM Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.package Datasource

package ioeither

import (
	IOE "github.com/IBM/fp-go/ioeither"
	"github.com/ibm-hyper-protect/contract-go/contract"
	Encrypt "github.com/ibm-hyper-protect/contract-go/encrypt/ioeither"
)

// OpenSSLEncryptAndSignContract returns the OpenSSL implementation of the encryption and signer
func OpenSSLEncryptAndSignContract(cert []byte) func([]byte) func(Contract.RawMap) IOE.IOEither[error, Contract.RawMap] {
	return EncryptAndSignContract(Encrypt.OpenSSLEncryptBasic(cert), Encrypt.OpenSSLSignDigest, Encrypt.OpenSSLPublicKey)
}
