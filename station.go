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

// Station type contains station-related data.
type Station struct {
	// Id is unique public station identifier.
	Id *int `json:"id,omitempty"`

	// TokenId is unique private station identifier.
	TokenId string `json:"token,omitempty"`

	// Description for this station.
	Description string `json:"desc,omitempty"`

	// Station current software version.
	Version string `json:"version,omitempty"`

	// Created timestamp.
	Created UnixTime `json:"created"`

	// Last seen timestamp (never if empty).
	Seen UnixTime `json:"seen,omitempty"`

	// Longitude of the station location.
	Longitude float64 `json:"long"`

	// Latitude of the station location.
	Latitude float64 `json:"lat"`

	// LastMeasurement contains last environment data measured by the station.
	LastMeasurement *Measurement `json:"last_measurement,omitempty"`
}
