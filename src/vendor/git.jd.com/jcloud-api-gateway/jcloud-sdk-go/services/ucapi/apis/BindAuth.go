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

type BindAuthRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 商城用户pin (Optional) */
    CheckPin *string `json:"checkPin"`

    /* 绑定用户pin (Optional) */
    Pin *string `json:"pin"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBindAuthRequest(
    regionId string,
) *BindAuthRequest {

	return &BindAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/personAuth/bindAuth",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param checkPin: 商城用户pin (Optional)
 * param pin: 绑定用户pin (Optional)
 */
func NewBindAuthRequestWithAllParams(
    regionId string,
    checkPin *string,
    pin *string,
) *BindAuthRequest {

    return &BindAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/personAuth/bindAuth",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CheckPin: checkPin,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBindAuthRequestWithoutParam() *BindAuthRequest {

    return &BindAuthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/personAuth/bindAuth",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *BindAuthRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param checkPin: 商城用户pin(Optional) */
func (r *BindAuthRequest) SetCheckPin(checkPin string) {
    r.CheckPin = &checkPin
}

/* param pin: 绑定用户pin(Optional) */
func (r *BindAuthRequest) SetPin(pin string) {
    r.Pin = &pin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BindAuthRequest) GetRegionId() string {
    return r.RegionId
}

type BindAuthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BindAuthResult `json:"result"`
}

type BindAuthResult struct {
}