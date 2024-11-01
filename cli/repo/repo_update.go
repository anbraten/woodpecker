// Copyright 2022 Woodpecker Authors
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

package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/urfave/cli/v3"

	"go.woodpecker-ci.org/woodpecker/v2/cli/internal"
	"go.woodpecker-ci.org/woodpecker/v2/woodpecker-go/woodpecker"
)

var repoUpdateCmd = &cli.Command{
	Name:      "update",
	Usage:     "update a repository",
	ArgsUsage: "<repo-id|repo-full-name>",
	Action:    repoUpdate,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "trusted",
			Usage: "repository is trusted",
		},
		&cli.BoolFlag{
			Name:  "gated",
			Usage: "repository is gated",
		},
		&cli.StringFlag{
			Name:  "require-approval",
			Usage: "repository requires approval for",
		},
		&cli.DurationFlag{
			Name:  "timeout",
			Usage: "repository timeout",
		},
		&cli.StringFlag{
			Name:  "visibility",
			Usage: "repository visibility",
		},
		&cli.StringFlag{
			Name:  "config",
			Usage: "repository configuration path (e.g. .woodpecker.yml)",
		},
		&cli.IntFlag{
			Name:  "pipeline-counter",
			Usage: "repository starting pipeline number",
		},
		&cli.BoolFlag{
			Name:  "unsafe",
			Usage: "validate updating the pipeline-counter is unsafe",
		},
	},
}

func repoUpdate(ctx context.Context, c *cli.Command) error {
	repoIDOrFullName := c.Args().First()
	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}
	repoID, err := internal.ParseRepo(client, repoIDOrFullName)
	if err != nil {
		return err
	}

	var (
		visibility      = c.String("visibility")
		config          = c.String("config")
		timeout         = c.Duration("timeout")
		trusted         = c.Bool("trusted")
		gated           = c.Bool("gated")
		requireApproval = c.String("require-approval")
		pipelineCounter = int(c.Int("pipeline-counter"))
		unsafe          = c.Bool("unsafe")
	)

	patch := new(woodpecker.RepoPatch)
	if c.IsSet("trusted") {
		patch.IsTrusted = &trusted
	}
	// TODO: remove isGated in next major release
	if c.IsSet("gated") {
		if gated {
			patch.RequireApproval = &woodpecker.RequireApprovalAllEvents
		} else {
			patch.RequireApproval = &woodpecker.RequireApprovalNone
		}
	}
	if c.IsSet("require-approval") {
		switch woodpecker.ApprovalMode(requireApproval) {
		case woodpecker.RequireApprovalNone, woodpecker.RequireApprovalForks, woodpecker.RequireApprovalPullRequests, woodpecker.RequireApprovalAllEvents:
			requireApproval := woodpecker.ApprovalMode(requireApproval)
			patch.RequireApproval = &requireApproval
		}

		// TODO: remove isGated in next major release
		if requireApproval == string(woodpecker.RequireApprovalAllEvents) {
			trueBool := true
			patch.IsGated = &trueBool
		}

		if requireApproval == string(woodpecker.RequireApprovalNone) {
			falseBool := false
			patch.IsGated = &falseBool
		}
	}
	if c.IsSet("timeout") {
		v := int64(timeout / time.Minute)
		patch.Timeout = &v
	}
	if c.IsSet("config") {
		patch.Config = &config
	}
	if c.IsSet("visibility") {
		switch visibility {
		case "public", "private", "internal":
			patch.Visibility = &visibility
		}
	}
	if c.IsSet("pipeline-counter") && !unsafe {
		fmt.Printf("Setting the pipeline counter is an unsafe operation that could put your repository in an inconsistent state. Please use --unsafe to proceed")
	}
	if c.IsSet("pipeline-counter") && unsafe {
		patch.PipelineCounter = &pipelineCounter
	}

	repo, err := client.RepoPatch(repoID, patch)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully updated repository %s\n", repo.FullName)
	return nil
}
