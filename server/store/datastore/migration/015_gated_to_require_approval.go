// Copyright 2024 Woodpecker Authors
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
// limitations under the License.

package migration

import (
	"fmt"

	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

// perPage015 set the size of the slice to read per page.
var perPage015 = 100

type repo015 struct {
	ID              int64  `json:"id"               xorm:"pk autoincr 'pipeline_id'"`
	IsGated         bool   `json:"gated"            xorm:"gated"`
	RequireApproval string `json:"require_approval" xorm:"require_approval"`
}

func (repo015) TableName() string {
	return "repos"
}

var gatedToRequireApproval = xormigrate.Migration{
	ID: "gated-to-require-approval",
	MigrateSession: func(sess *xorm.Session) (err error) {
		if err := sess.Sync(new(repo015)); err != nil {
			return fmt.Errorf("sync new models failed: %w", err)
		}

		page := 0
		oldRepos := make([]*repo015, 0, perPage015)

		for {
			oldRepos = oldRepos[:0]

			err := sess.Limit(perPage015, page*perPage015).Cols("require_approval", "gated").Find(&oldRepos)
			if err != nil {
				return err
			}

			for _, oldRepo := range oldRepos {
				var newRepo repo015
				newRepo.ID = oldRepo.ID
				newRepo.RequireApproval = "forks"
				if oldRepo.IsGated {
					newRepo.RequireApproval = "all_events"
				}

				if _, err := sess.ID(oldRepo.ID).Cols("require_approval").Update(newRepo); err != nil {
					return err
				}
			}

			if len(oldRepos) < perPage015 {
				break
			}

			page++
		}

		return dropTableColumns(sess, "repos", "gated")
	},
}
