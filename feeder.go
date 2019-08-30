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

// FeederData type contains data posted by the station to data collector.
type FeederData struct {
	// TokenId is unique private station identifier.
	TokenId string `json:"token"`

	// Station software version string.
	Version string `json:"version,omitempty"`

	// Measurements contains array of data measured by the station.
	Measurements []Measurement `json:"measurements,omitempty"`
}
