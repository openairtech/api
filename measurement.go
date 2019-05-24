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

// Measurement type contains environment data measured by the station.
type Measurement struct {
	// Timestamp of this measurement.
	Timestamp *UnixTime `json:"timestamp,omitempty"`

	// Temperature (in Celisius degrees).
	Temperature *float32 `json:"temperature,omitempty"`

	// Relative humidity.
	Humidity *float32 `json:"humidity,omitempty"`

	// Atmospheric pressure (in hPa).
	Pressure *float32 `json:"pressure,omitempty"`

	// PM2.5 particulate matter concentration (in µg/m3).
	Pm25 *float32 `json:"pm25,omitempty"`

	// PM10 particulate matter concentration (in µg/m3).
	Pm10 *float32 `json:"pm10,omitempty"`

	// AQI air quality index value.
	Aqi *int `json:"aqi,omitempty"`
}
