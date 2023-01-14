// Copyright 2021 Woodpecker Authors
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

package datastore

import (
	"strconv"

	"xorm.io/builder"

	"github.com/woodpecker-ci/woodpecker/server/model"
)

const orderSecretsBy = "secret_name"

func (s storage) secretDecrypt(secret *model.Secret) error {
	decryptedValue, err := s.encryption.Decrypt(secret.Value, strconv.Itoa(int(secret.ID)))
	if err != nil {
		return err
	}

	secret.Value = decryptedValue

	return nil
}

func (s storage) secretEncrypt(secret *model.Secret) error {
	encryptedValue, err := s.encryption.Encrypt(secret.Value, strconv.Itoa(int(secret.ID)))
	if err != nil {
		return err
	}

	secret.Value = encryptedValue

	return nil
}

func (s storage) SecretFind(repo *model.Repo, name string) (*model.Secret, error) {
	secret := &model.Secret{
		RepoID: repo.ID,
		Name:   name,
	}

	err := wrapGet(s.engine.Get(secret))
	if err != nil {
		return nil, err
	}

	err = s.secretDecrypt(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (s storage) SecretList(repo *model.Repo, includeGlobalAndOrgSecrets bool) ([]*model.Secret, error) {
	secrets := make([]*model.Secret, 0, perPage)
	var cond builder.Cond = builder.Eq{"secret_repo_id": repo.ID}
	if includeGlobalAndOrgSecrets {
		cond = cond.Or(builder.Eq{"secret_owner": repo.Owner}).
			Or(builder.And(builder.Eq{"secret_owner": ""}, builder.Eq{"secret_repo_id": 0}))
	}
	return secrets, s.engine.Where(cond).OrderBy(orderSecretsBy).Find(&secrets)
}

func (s storage) SecretListAll() ([]*model.Secret, error) {
	var secrets []*model.Secret
	return secrets, s.engine.Find(&secrets)
}

func (s storage) SecretCreate(secret *model.Secret) error {
	err := s.secretEncrypt(secret)
	if err != nil {
		return err
	}

	// only Insert set auto created ID back to object
	_, err = s.engine.Insert(secret)
	return err
}

func (s storage) SecretUpdate(secret *model.Secret) error {
	err := s.secretEncrypt(secret)
	if err != nil {
		return err
	}

	_, err = s.engine.ID(secret.ID).AllCols().Update(secret)
	return err
}

func (s storage) SecretDelete(secret *model.Secret) error {
	_, err := s.engine.ID(secret.ID).Delete(new(model.Secret))
	return err
}

func (s storage) OrgSecretFind(owner, name string) (*model.Secret, error) {
	secret := &model.Secret{
		Owner: owner,
		Name:  name,
	}
	return secret, wrapGet(s.engine.Get(secret))
}

func (s storage) OrgSecretList(owner string) ([]*model.Secret, error) {
	secrets := make([]*model.Secret, 0, perPage)
	return secrets, s.engine.Where("secret_owner = ?", owner).Find(&secrets)
}

func (s storage) GlobalSecretFind(name string) (*model.Secret, error) {
	secret := &model.Secret{
		Name: name,
	}
	return secret, wrapGet(s.engine.Where(builder.And(builder.Eq{"secret_owner": ""}, builder.Eq{"secret_repo_id": 0})).Get(secret))
}

func (s storage) GlobalSecretList() ([]*model.Secret, error) {
	secrets := make([]*model.Secret, 0, perPage)
	return secrets, s.engine.Where(builder.And(builder.Eq{"secret_owner": ""}, builder.Eq{"secret_repo_id": 0})).OrderBy(orderSecretsBy).Find(&secrets)
}
