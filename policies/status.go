// Copyright 2024 Security Scorecard Authors
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
// limitations under the License

package policies

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v61/github"
	"github.com/ossf/scorecard-action/options"
)

var actionCtx = "scorecard-action"

type Status struct {
	ctx          context.Context
	client       *github.RepositoriesService
	owner        string
	repo         string
	commitSHA    string
	repoStatusID int64
}

func (c *Status) SetupAndStart(ctx context.Context, opts options.Options) error {
	c.ctx = ctx

	r := strings.Split(opts.GithubRepository, "/")
	c.owner = r[0]
	c.repo = r[1]

	fmt.Printf("FIXME opts: %+v\n", opts)
	c.commitSHA = opts.ScorecardOpts.Commit

	token := os.Getenv(options.EnvGithubAuthToken)
	client := github.NewClient(nil).WithAuthToken(token)

	c.client = client.Repositories

	pending := "pending"
	rs := &github.RepoStatus{
		State:   &pending,
		Context: &actionCtx,
		// TargetURL is the URL of the page representing this status. It will be
		// linked from the GitHub UI to allow users to see the source of the status.
		// TargetURL *string `json:"target_url,omitempty"`

		// Description is a short high level summary of the status.
		// Description *string `json:"description,omitempty"`

		// AvatarURL is the URL of the avatar of this status.
		// AvatarURL *string `json:"avatar_url,omitempty"`

		// Creator   *User      `json:"creator,omitempty"`
		CreatedAt: &github.Timestamp{Time: time.Now()},
		// UpdatedAt: &github.Timestamp{Time: time.Now()},
	}

	rs, _, err := c.client.CreateStatus(c.ctx, c.owner, c.repo, c.commitSHA, rs)

	if err != nil {
		return fmt.Errorf("CreateStatus: %w", err)
	}

	c.repoStatusID = rs.GetID()

	return nil
}

// func (c *Status) Start() error {
// 	pending := "pending"
// 	rs := &github.RepoStatus{
// 		ID:      &c.repoStatusID,
// 		State:   &pending,
// 		Context: &actionCtx,
// 		// TargetURL is the URL of the page representing this status. It will be
// 		// linked from the GitHub UI to allow users to see the source of the status.
// 		// TargetURL *string `json:"target_url,omitempty"`

// 		// Description is a short high level summary of the status.
// 		// Description *string `json:"description,omitempty"`

// 		// AvatarURL is the URL of the avatar of this status.
// 		// AvatarURL *string `json:"avatar_url,omitempty"`

// 		// Creator   *User      `json:"creator,omitempty"`
// 		UpdatedAt: &github.Timestamp{Time: time.Now()},
// 	}

// 	rs, _, err := c.client.CreateStatus(c.ctx, c.owner, c.repo, c.commitSHA, rs)

// 	if err != nil {
// 		return fmt.Errorf("CreateStatus: %w", err)
// 	}

// 	return nil
// }

func (c *Status) Fail() error {
	failure := "failure"
	rs := &github.RepoStatus{
		ID:      &c.repoStatusID,
		State:   &failure,
		Context: &actionCtx,
		// TargetURL is the URL of the page representing this status. It will be
		// linked from the GitHub UI to allow users to see the source of the status.
		// TargetURL *string `json:"target_url,omitempty"`

		// Description is a short high level summary of the status.
		// Description *string `json:"description,omitempty"`

		// AvatarURL is the URL of the avatar of this status.
		// AvatarURL *string `json:"avatar_url,omitempty"`

		// Creator   *User      `json:"creator,omitempty"`
		UpdatedAt: &github.Timestamp{Time: time.Now()},
	}
	rs, _, err := c.client.CreateStatus(c.ctx, c.owner, c.repo, c.commitSHA, rs)

	if err != nil {
		return fmt.Errorf("CreateStatus: %w", err)
	}

	return nil
}

func (c *Status) Success() error {
	success := "success"
	rs := &github.RepoStatus{
		ID:      &c.repoStatusID,
		State:   &success,
		Context: &actionCtx,
		// TargetURL is the URL of the page representing this status. It will be
		// linked from the GitHub UI to allow users to see the source of the status.
		// TargetURL *string `json:"target_url,omitempty"`

		// Description is a short high level summary of the status.
		// Description *string `json:"description,omitempty"`

		// AvatarURL is the URL of the avatar of this status.
		// AvatarURL *string `json:"avatar_url,omitempty"`

		// Creator   *User      `json:"creator,omitempty"`
		UpdatedAt: &github.Timestamp{Time: time.Now()},
	}
	rs, _, err := c.client.CreateStatus(c.ctx, c.owner, c.repo, c.commitSHA, rs)

	if err != nil {
		return fmt.Errorf("CreateStatus: %w", err)
	}

	return nil
}
