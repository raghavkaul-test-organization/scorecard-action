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
// limitations under the License.

// Package policies sets a action policy to set a status check.
package policies

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type policyCriteria struct {
	Mode  string `yaml:"mode"`
	Score int    `yaml:"score"`
}

// ScorecardActionPolicy is a policy with numeric score thresholds for each check.
type ScorecardActionPolicy struct {
	Criteria map[string]policyCriteria `yaml:"policies"`
	Version  int                       `yaml:"version"`
}

type PolicyResult = bool

const (
	Pass PolicyResult = true
	Fail PolicyResult = false
)

func CheckResults(p ScorecardActionPolicy) PolicyResult {
	return true
}

// ParsePolicyFromFile takes a policy file and returns an AttestationPolicy.
func ParsePolicyFromFile(policyFile string) (*ScorecardActionPolicy, error) {
	if policyFile != "" {
		data, err := os.ReadFile(policyFile)
		if err != nil {
			return nil, fmt.Errorf("couldn't read policy file: %w", err)
		}

		sap, err := ParsePolicyFromYAML(data)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse policy: %w", err)
		}

		return sap, nil
	}

	return nil, nil
}

// ParsePolicyFromYAML parses a policy file and returns a AttestationPolicy.
func ParsePolicyFromYAML(b []byte) (*ScorecardActionPolicy, error) {
	ap := ScorecardActionPolicy{}

	err := yaml.Unmarshal(b, &ap)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse policy into yaml: %w", err)
	}

	return &ap, nil
}
