// Copyright 2019 Google LLC
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

// Code generated by "mdtogo"; DO NOT EDIT.
package fndocs

var FnShort = `Generate, transform, and validate configuration files.`
var FnLong = `
| Configuration Read From | Configuration Written To |
|-------------------------|--------------------------|
| local files or stdin    | local files or stdout    |

Functions are executables ([that you can write](#developing-functions))
packaged in container images which accept a collection of Resource
configuration as input, and emit a collection of Resource configuration as output.
`
var FnExamples = `
  # run the function defined by gcr.io/example.com/my-fn as a local container
  # against the configuration in DIR
  kpt fn run DIR/ --image gcr.io/example.com/my-fn

  # run the functions declared in files under FUNCTIONS_DIR/
  kpt fn run DIR/ --fn-path FUNCTIONS_DIR/

  # run the functions declared in files under DIR/
  kpt fn run DIR/
`

var ExportShort = `Auto-generating function pipelines for different workflow orchestrators`
var ExportLong = `
  kpt fn export ORCHESTRATOR DIR/ [--fn-path FUNCTIONS_DIR/] [--output OUTPUT_FILENAME]
  
  ORCHESTRATOR:
    Supported orchestrators are github-actions, and cloud-build.
  DIR:
    Path to a package directory. If you do not specify the --fn-path flag, this command will discover functions in DIR and run them against resources in it.
`
var ExportExamples = `
  # read functions from DIR, run them against it as one step
  # write the generated GitHub Actions pipeline to main.yaml
  kpt fn export github-actions DIR/ --output main.yaml

  # discover functions in FUNCTIONS_DIR and run them against Resource in DIR.
  kpt fn export github-actions DIR/ --fn-path FUNCTIONS_DIR/
`

var RunShort = `Locally execute one or more functions in containers`
var RunLong = `
  kpt fn run DIR [flags]

If the container exits with non-zero status code, run will fail and print the
container ` + "`" + `STDERR` + "`" + `.

  DIR:
    Path to a package directory.  Defaults to stdin if unspecified.
`
var RunExamples = `
  # read the Resources from DIR, provide them to a container my-fun as input,
  # write my-fn output back to DIR
  kpt fn run DIR/ --image gcr.io/example.com/my-fn

  # provide the my-fn with an input ConfigMap containing ` + "`" + `data: {foo: bar}` + "`" + `
  kpt fn run DIR/ --image gcr.io/example.com/my-fn:v1.0.0 -- foo=bar

  # run the functions in FUNCTIONS_DIR against the Resources in DIR
  kpt fn run DIR/ --fn-path FUNCTIONS_DIR/

  # discover functions in DIR and run them against Resource in DIR.
  # functions may be scoped to a subset of Resources -- see ` + "`" + `kpt help fn run` + "`" + `
  kpt fn run DIR/
`

var SinkShort = `Explicitly specify an output sink for a function`
var SinkLong = `
  kpt fn sink [DIR]
  
  DIR:
    Path to a package directory.  Defaults to stdout if unspecified.
`
var SinkExamples = `
  # run a function using explicit sources and sinks
  kpt fn source DIR/ | kpt fn run --image gcr.io/example.com/my-fn | kpt fn sink DIR/
`

var SourceShort = `Explicitly specify an input source for a function`
var SourceLong = `
  kpt fn source [DIR...]
  
  DIR:
    Path to a package directory.  Defaults to stdin if unspecified.
`
var SourceExamples = `
  # print to stdout configuration formatted as an input source
  kpt fn source DIR/

  # run a function using explicit sources and sinks
  kpt fn source DIR/ | kpt fn run --image gcr.io/example.com/my-fn | kpt fn sink DIR/
`
