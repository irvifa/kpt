// Copyright 2020 Google LLC
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

package orchestrators

import (
	"strings"
	"testing"

	"github.com/GoogleContainerTools/kpt/internal/cmdexport/types"
	"gotest.tools/assert"
)

type testCase struct {
	description string
	config      *types.PipelineConfig
	expected    string
}

type testSuite struct {
	pipeline  Pipeline
	testCases []testCase
}

func TestPipeline(t *testing.T) {
	testSuites := []testSuite{
		githubActionsTestSuite,
		cloudBuildTestSuite,
	}

	for _, testSuite := range testSuites {
		pipeline := testSuite.pipeline
		testCases := testSuite.testCases

		for i := range testCases {
			testCase := testCases[i]

			t.Run(testCase.description, func(t *testing.T) {
				pipeline, _ := pipeline.Init(testCase.config).Generate()

				actual := string(pipeline)
				expected := strings.TrimLeft(testCase.expected, "\n")

				assert.Equal(t, expected, actual)
			})
		}
	}
}
