// Copyright 2023 Woodpecker Authors
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

package pipeline

import (
	backend "go.woodpecker-ci.org/woodpecker/pipeline/backend/types"
	"go.woodpecker-ci.org/woodpecker/pipeline/multipart"
)

// Logger handles the process logging.
type Logger interface {
	Log(*backend.Step, multipart.Reader) error
}

// LogFunc type is an adapter to allow the use of an ordinary
// function for process logging.
type LogFunc func(*backend.Step, multipart.Reader) error

// Log calls f(step, r).
func (f LogFunc) Log(step *backend.Step, r multipart.Reader) error {
	return f(step, r)
}
