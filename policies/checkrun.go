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
	"time"

	"github.com/google/go-github/v61/github"
)

type CheckRun struct {
	ctx        context.Context
	client     *github.ChecksService
	owner      string
	repo       string
	checkRunID int64
}

func (c *CheckRun) Setup(ctx context.Context, headSHA, owner, repo string) error {
	c.ctx = ctx
	c.owner = owner
	c.repo = repo

	token := os.Getenv("GITHUB_TOKEN")
	client := github.NewClient(nil).WithAuthToken(token)

	c.client = client.Checks

	opts := github.CreateCheckRunOptions{
		Name:    "scorecard-action",
		HeadSHA: headSHA,
		// DetailsURL: ,
		// ExternalID: ,
		// Status: ,
		// Conclusion: ,
		StartedAt: &github.Timestamp{Time: time.Now()},
		// CompletedAt: ,
		// Output: ,
		// Actions: ,
	}

	cr, _, err := c.client.CreateCheckRun(c.ctx, owner, repo, opts)
	if err != nil {
		return fmt.Errorf("CreateCheckRun: %w", err)
	}

	c.checkRunID = *cr.ID

	return nil
}

func (c *CheckRun) Start() error {
	in_progress := "in_progress"
	opts := github.UpdateCheckRunOptions{
		Status: &in_progress,
	}

	_, _, err := c.client.UpdateCheckRun(c.ctx, c.owner, c.repo, c.checkRunID, opts)

	return err
}

func (c *CheckRun) Fail() error {
	completed := "completed"
	failure := "failure"
	opts := github.UpdateCheckRunOptions{
		Status:      &completed,
		Conclusion:  &failure,
		CompletedAt: &github.Timestamp{Time: time.Now()},
		// Output: "",
	}
	_, _, err := c.client.UpdateCheckRun(c.ctx, c.owner, c.repo, c.checkRunID, opts)

	return err
}

func (c *CheckRun) Success() error {
	completed := "completed"
	success := "success"
	opts := github.UpdateCheckRunOptions{
		Status:      &completed,
		Conclusion:  &success,
		CompletedAt: &github.Timestamp{Time: time.Now()},
		// Output: "",
	}
	_, _, err := c.client.UpdateCheckRun(c.ctx, c.owner, c.repo, c.checkRunID, opts)

	return err
}
