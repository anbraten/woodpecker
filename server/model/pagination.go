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

package model

import (
	"fmt"
	"strings"
)

type ListOptions struct {
	All     bool
	Page    int
	PerPage int
}

func ApplyPagination[T any](d *ListOptions, slice []T) []T {
	if d.All {
		return slice
	}
	if d.PerPage*(d.Page-1) > len(slice) {
		return []T{}
	}
	if d.PerPage*(d.Page) > len(slice) {
		return slice[d.PerPage*(d.Page-1):]
	}
	return slice[d.PerPage*(d.Page-1) : d.PerPage*(d.Page)]
}

func (d *ListOptions) Encode() string {
	query := []string{}

	if d.Page != 0 {
		query = append(query, fmt.Sprintf("page=%d", d.Page))
	}

	if d.PerPage != 0 {
		query = append(query, fmt.Sprintf("per_page=%d", d.PerPage))
	}

	if d.All {
		query = append(query, "all=true")
	}

	return strings.Join(query, "&")
}
