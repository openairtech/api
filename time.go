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
	"time"
)

// UnixTime is standard time.Time with JSON (un)marshal in Unix timestamp format.
type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *UnixTime) UnmarshalJSON(s []byte) error {
	r := strings.Replace(string(s), `"`, ``, -1)

	v, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(v, 0)

	return nil
}

func (t UnixTime) String() string {
	return time.Time(t).String()
}
