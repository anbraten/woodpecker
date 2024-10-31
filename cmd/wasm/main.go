// Copyright 2024 Woodpecker Authors
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

//go:build wasm

package main

import (
	"errors"
	"fmt"
	"os"
	"syscall/js"

	term_env "github.com/muesli/termenv"

	pipeline_errors "go.woodpecker-ci.org/woodpecker/v2/pipeline/errors"
	"go.woodpecker-ci.org/woodpecker/v2/pipeline/frontend/yaml"
	"go.woodpecker-ci.org/woodpecker/v2/pipeline/frontend/yaml/linter"
)

func lint(rawConfig string) (string, error) {
	parsedConfig, err := yaml.ParseString(rawConfig)
	if err != nil {
		return "", err
	}

	file := "file.yaml"

	config := &linter.WorkflowConfig{
		File:      file,
		RawConfig: rawConfig,
		Workflow:  parsedConfig,
	}

	// TODO: lint multiple files at once to allow checks for sth like "depends_on" to work
	err = linter.New(
		linter.WithTrusted(true),
		// linter.PrivilegedPlugins(c.StringSlice("plugins-privileged")),
		// linter.WithTrustedClonePlugins(c.StringSlice("plugins-trusted-clone")),
	).Lint([]*linter.WorkflowConfig{config})
	if err != nil {
		str, err := FormatLintError(config.File, err)
		return fmt.Sprintf("%s\n❌ Config is invalid", str), err
	}

	return "✅ Config is valid", nil
}

func FormatLintError(file string, err error) (string, error) {
	if err == nil {
		return "", nil
	}

	output := term_env.NewOutput(os.Stdout)
	str := ""

	amountErrors := 0
	amountWarnings := 0
	linterErrors := pipeline_errors.GetPipelineErrors(err)
	for _, err := range linterErrors {
		line := "  "

		if err.IsWarning {
			line = fmt.Sprintf("%s ⚠️ ", line)
			amountWarnings++
		} else {
			line = fmt.Sprintf("%s ❌", line)
			amountErrors++
		}

		if data := pipeline_errors.GetLinterData(err); data != nil {
			line = fmt.Sprintf("%s %s\t%s", line, output.String(data.Field).Bold(), err.Message)
		} else {
			line = fmt.Sprintf("%s %s", line, err.Message)
		}

		// TODO: use table output
		str = fmt.Sprintf("%s%s\n", str, line)
	}

	if amountErrors > 0 {
		if amountWarnings > 0 {
			str = fmt.Sprintf("🔥 %s has %d errors and warnings:\n%s", output.String(file).Underline(), len(linterErrors), str)
		} else {
			str = fmt.Sprintf("🔥 %s has %d errors:\n%s", output.String(file).Underline(), len(linterErrors), str)
		}
		return str, errors.New("config has errors")
	}

	str = fmt.Sprintf("⚠️  %s has %d warnings:\n%s", output.String(file).Underline(), len(linterErrors), str)
	return str, nil
}

func main() {
	js.Global().Set("lint", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "expected exactly one argument"
		}

		res, err := lint(args[0].String())
		if err != nil && err.Error() != "config has errors" {
			return err.Error()
		}

		return res
	}))

	select {}
}
