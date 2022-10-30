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
	"github.com/woodpecker-ci/woodpecker/server/model"
)

// TODO: need tests before converting to builder statement
func (s storage) GetPipelineQueue() ([]*model.Feed, error) {
	feed := make([]*model.Feed, 0, perPage)
	// TODO: use builder (do not behave same as pure sql, fix that)
	err := s.engine.SQL(`
SELECT
 repo_owner
,repo_name
,repo_full_name
,pipeline_number
,pipeline_event
,pipeline_status
,pipeline_created
,pipeline_started
,pipeline_finished
,pipeline_commit
,pipeline_branch
,pipeline_ref
,pipeline_refspec
,pipeline_remote
,pipeline_title
,pipeline_message
,pipeline_author
,pipeline_email
,pipeline_avatar
FROM
 pipelines p
,repos r
WHERE p.pipeline_repo_id = r.repo_id
  AND p.pipeline_status IN ('pending','running')
`).Find(&feed)
	return feed, err
}

func (s storage) UserFeed(user *model.User) ([]*model.Feed, error) {
	feed := make([]*model.Feed, 0, perPage)
	// TODO: use builder (do not behave same as pure sql, fix that)
	return feed, s.engine.SQL(`
SELECT
 repo_owner
,repo_name
,repo_full_name
,pipeline_number
,pipeline_event
,pipeline_status
,pipeline_created
,pipeline_started
,pipeline_finished
,pipeline_commit
,pipeline_branch
,pipeline_ref
,pipeline_refspec
,pipeline_remote
,pipeline_title
,pipeline_message
,pipeline_author
,pipeline_email
,pipeline_avatar
FROM repos
INNER JOIN perms  ON perms.perm_repo_id   = repos.repo_id
INNER JOIN pipelines ON pipelines.pipeline_repo_id = repos.repo_id
WHERE perms.perm_user_id = ?
  AND (perms.perm_push = ? OR perms.perm_admin = ?)
ORDER BY pipeline_id DESC
LIMIT 50
`, user.ID, true, true).Find(&feed)
}

func (s storage) RepoListLatest(user *model.User) ([]*model.Feed, error) {
	feed := make([]*model.Feed, 0, perPage)
	// TODO: use builder (do not behave same as pure sql, fix that)
	return feed, s.engine.SQL(`
SELECT
 repo_owner
,repo_name
,repo_full_name
,pipeline_number
,pipeline_event
,pipeline_status
,pipeline_created
,pipeline_started
,pipeline_finished
,pipeline_commit
,pipeline_branch
,pipeline_ref
,pipeline_refspec
,pipeline_remote
,pipeline_title
,pipeline_message
,pipeline_author
,pipeline_email
,pipeline_avatar
FROM repos LEFT OUTER JOIN pipelines ON pipeline_id = (
	SELECT pipeline_id FROM pipelines
	WHERE pipelines.pipeline_repo_id = repos.repo_id
	ORDER BY pipeline_id DESC
	LIMIT 1
)
INNER JOIN perms ON perms.perm_repo_id = repos.repo_id
WHERE perms.perm_user_id = ?
  AND (perms.perm_push = ? OR perms.perm_admin = ?)
  AND repos.repo_active = ?
ORDER BY repo_full_name ASC;
`, user.ID, true, true, true).
		Find(&feed)
}
