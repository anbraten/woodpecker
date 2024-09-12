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

package exec

import (
	"runtime"
	"strings"

	"github.com/urfave/cli/v3"

	"go.woodpecker-ci.org/woodpecker/v2/pipeline/frontend/metadata"
	"go.woodpecker-ci.org/woodpecker/v2/server/model"
	"go.woodpecker-ci.org/woodpecker/v2/version"
)

// return the metadata from the cli context.
func metadataFromCommand(c *cli.Command, workflow *model.Workflow) metadata.Metadata {
	platform := c.String("system-platform")
	if platform == "" {
		platform = runtime.GOOS + "/" + runtime.GOARCH
	}

	fullRepoName := c.String("repo-name")
	repoOwner := ""
	repoName := ""
	if idx := strings.LastIndex(fullRepoName, "/"); idx != -1 {
		repoOwner = fullRepoName[:idx]
		repoName = fullRepoName[idx+1:]
	}

	return metadata.Metadata{
		Repo: metadata.Repo{
			Name:        repoName,
			Owner:       repoOwner,
			RemoteID:    c.String("repo-remote-id"),
			ForgeURL:    c.String("repo-url"),
			CloneURL:    c.String("repo-clone-url"),
			CloneSSHURL: c.String("repo-clone-ssh-url"),
			Private:     c.Bool("repo-private"),
			Trusted:     c.Bool("repo-trusted"),
			Branch:      "main", // TODO: allow to set the default branch
		},
		Curr: metadata.Pipeline{
			Number:     c.Int("pipeline-number"),
			Parent:     c.Int("pipeline-parent"),
			Created:    c.Int("pipeline-created"),
			Started:    c.Int("pipeline-started"),
			Finished:   c.Int("pipeline-finished"),
			Status:     c.String("pipeline-status"),
			Event:      c.String("pipeline-event"),
			ForgeURL:   c.String("pipeline-url"),
			DeployTo:   c.String("pipeline-deploy-to"),
			DeployTask: c.String("pipeline-deploy-task"),
			Commit: metadata.Commit{
				Sha:     c.String("commit-sha"),
				Ref:     c.String("commit-ref"),
				Refspec: c.String("commit-refspec"),
				Branch:  c.String("commit-branch"),
				Message: c.String("commit-message"),
				Author: metadata.Author{
					Name:   c.String("commit-author-name"),
					Email:  c.String("commit-author-email"),
					Avatar: c.String("commit-author-avatar"),
				},
			},
		},
		Prev: metadata.Pipeline{
			Number:   c.Int("prev-pipeline-number"),
			Created:  c.Int("prev-pipeline-created"),
			Started:  c.Int("prev-pipeline-started"),
			Finished: c.Int("prev-pipeline-finished"),
			Status:   c.String("prev-pipeline-status"),
			Event:    c.String("prev-pipeline-event"),
			ForgeURL: c.String("prev-pipeline-url"),
			Commit: metadata.Commit{
				Sha:     c.String("prev-commit-sha"),
				Ref:     c.String("prev-commit-ref"),
				Refspec: c.String("prev-commit-refspec"),
				Branch:  c.String("prev-commit-branch"),
				Message: c.String("prev-commit-message"),
				Author: metadata.Author{
					Name:   c.String("prev-commit-author-name"),
					Email:  c.String("prev-commit-author-email"),
					Avatar: c.String("prev-commit-author-avatar"),
				},
			},
		},
		Workflow: metadata.Workflow{
			Name:   workflow.Name,
			Number: workflow.PID,
			Matrix: workflow.Environ,
		},
		Step: metadata.Step{
			Name:   c.String("step-name"),
			Number: int(c.Int("step-number")),
		},
		Sys: metadata.System{
			Name:     c.String("system-name"),
			URL:      c.String("system-url"),
			Platform: platform,
			Version:  version.Version,
		},
		Forge: metadata.Forge{
			Type: c.String("forge-type"),
			URL:  c.String("forge-url"),
		},
	}
}
