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
	"encoding/base64"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/google/tink/go/tink"
)

type TinkEncryption struct {
	keysetFilePath    string
	primaryKeyID      string
	encryption        tink.AEAD
	keysetFileWatcher *fsnotify.Watcher
}

func NewTinkEncryption(keysetFilePath string) *TinkEncryption {
	return &TinkEncryption{keysetFilePath: keysetFilePath}
}

func (t *TinkEncryption) Build() error {
	svc := &tinkEncryptionService{
		keysetFilePath:    c.keysetFilePath,
		primaryKeyID:      "",
		encryption:        nil,
		store:             c.store,
		keysetFileWatcher: nil,
		clients:           c.clients,
	}
	err := svc.initClients()
	if err != nil {
		return nil, fmt.Errorf(errTemplateFailedInitializingClients, err)
	}

	err = svc.loadKeyset()
	if err != nil {
		return nil, fmt.Errorf(errTemplateTinkFailedLoadingKeyset, err)
	}

	err = svc.validateKeyset()
	if err == errEncryptionNotEnabled {
		err = svc.enable()
	} else if err == errEncryptionKeyRotated {
		err = svc.rotate()
	}

	if err != nil {
		return nil, fmt.Errorf(errTemplateTinkFailedValidatingKeyset, err)
	}

	err = svc.initFileWatcher()
	if err != nil {
		return nil, fmt.Errorf(errTemplateTinkFailedInitializeFileWatcher, err)
	}
	return svc, nil
}

func (t *TinkEncryption) Encrypt(plaintext, associatedData string) (string, error) {
	msg := []byte(plaintext)
	aad := []byte(associatedData)
	ciphertext, err := t.encryption.Encrypt(msg, aad)
	if err != nil {
		return "", fmt.Errorf(errTemplateEncryptionFailed, err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (t *TinkEncryption) Decrypt(ciphertext, associatedData string) (string, error) {
	ct, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf(errTemplateBase64DecryptionFailed, err)
	}

	plaintext, err := svc.encryption.Decrypt(ct, []byte(associatedData))
	if err != nil {
		return "", fmt.Errorf(errTemplateDecryptionFailed, err)
	}
	return string(plaintext), nil
}
