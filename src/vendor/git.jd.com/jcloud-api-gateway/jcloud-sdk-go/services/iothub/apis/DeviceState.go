// Copyright 2018 JDCLOUD.COM
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
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
)

type DeviceStateRequest struct {

    core.JDCloudRequest

    /* Device 唯一标识  */
    DeviceId string `json:"deviceId"`

    /*  (Optional) */
    States *string `json:"states"`
}

/*
 * param deviceId: Device 唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeviceStateRequest(
    deviceId string,
) *DeviceStateRequest {

	return &DeviceStateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/device/{deviceId}/state",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        DeviceId: deviceId,
	}
}

/*
 * param deviceId: Device 唯一标识 (Required)
 * param states:  (Optional)
 */
func NewDeviceStateRequestWithAllParams(
    deviceId string,
    states *string,
) *DeviceStateRequest {

    return &DeviceStateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/device/{deviceId}/state",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        DeviceId: deviceId,
        States: states,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeviceStateRequestWithoutParam() *DeviceStateRequest {

    return &DeviceStateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/device/{deviceId}/state",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param deviceId: Device 唯一标识(Required) */
func (r *DeviceStateRequest) SetDeviceId(deviceId string) {
    r.DeviceId = deviceId
}

/* param states: (Optional) */
func (r *DeviceStateRequest) SetStates(states string) {
    r.States = &states
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeviceStateRequest) GetRegionId() string {
    return ""
}

type DeviceStateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeviceStateResult `json:"result"`
}

type DeviceStateResult struct {
    Data string `json:"data"`
}