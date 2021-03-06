// Copyright © 2019 Victor Antonovich <victor@antonovich.me>
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

package api

import (
	"strconv"
	"strings"
)

// StatusCode for API operation result.
type StatusCode int

const (
	StatusOk               StatusCode = 200
	StatusBadRequest       StatusCode = 400
	StatusNotFound         StatusCode = 404
	StatusMethodNotAllowed StatusCode = 405
	StatusServerError      StatusCode = 500
)

func (sc *StatusCode) UnmarshalJSON(s []byte) error {
	r := strings.Replace(string(s), `"`, ``, -1)

	v, err := strconv.Atoi(r)
	if err != nil {
		return err
	}

	*sc = StatusCode(v)

	return nil
}
