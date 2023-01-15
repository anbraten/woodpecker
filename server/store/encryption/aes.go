// Copyright 2023 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/google/tink/go/subtle/random"
)

type AESEncryption struct {
	cipher     cipher.AEAD
	privateKey string
}

func NewAESEncryptionService(privateKey string) EncryptionService {
	return &AESEncryption{
		privateKey: privateKey,
	}
}

func (a AESEncryption) Init() error {
	key, _ := hex.DecodeString(a.privateKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	a.cipher, err = cipher.NewGCM(block)

	return nil
}

func (a *AESEncryption) Encrypt(plaintext, associatedData string) (string, error) {
	msg := []byte(plaintext)
	aad := []byte(associatedData)

	nonce := random.GetRandomBytes(uint32(AESGCMSIVNonceSize))
	ciphertext := a.cipher.Seal(nil, nonce, msg, aad)

	result := make([]byte, 0, AESGCMSIVNonceSize+len(ciphertext))
	result = append(result, nonce...)
	result = append(result, ciphertext...)

	return base64.StdEncoding.EncodeToString(result), nil
}

func (a *AESEncryption) Decrypt(ciphertext, associatedData string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf(errTemplateBase64DecryptionFailed, err)
	}

	nonce := bytes[:AESGCMSIVNonceSize]
	message := bytes[AESGCMSIVNonceSize:]

	plaintext, err := a.cipher.Open(nil, nonce, message, []byte(associatedData))
	if err != nil {
		return "", fmt.Errorf(errTemplateDecryptionFailed, err)
	}
	return string(plaintext), nil
}
