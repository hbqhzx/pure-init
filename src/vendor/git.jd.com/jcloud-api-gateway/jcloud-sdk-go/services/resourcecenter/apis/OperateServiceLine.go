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

type OperateServiceLineRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    ServiceId int `json:"serviceId"`

    /*   */
    Operate int `json:"operate"`
}

/*
 * param regionId:  (Required)
 * param serviceId:  (Required)
 * param operate:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOperateServiceLineRequest(
    regionId string,
    serviceId int,
    operate int,
) *OperateServiceLineRequest {

	return &OperateServiceLineRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/serviceline/manage/operate",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceId: serviceId,
        Operate: operate,
	}
}

/*
 * param regionId:  (Required)
 * param serviceId:  (Required)
 * param operate:  (Required)
 */
func NewOperateServiceLineRequestWithAllParams(
    regionId string,
    serviceId int,
    operate int,
) *OperateServiceLineRequest {

    return &OperateServiceLineRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceline/manage/operate",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceId: serviceId,
        Operate: operate,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOperateServiceLineRequestWithoutParam() *OperateServiceLineRequest {

    return &OperateServiceLineRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceline/manage/operate",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *OperateServiceLineRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serviceId: (Required) */
func (r *OperateServiceLineRequest) SetServiceId(serviceId int) {
    r.ServiceId = serviceId
}

/* param operate: (Required) */
func (r *OperateServiceLineRequest) SetOperate(operate int) {
    r.Operate = operate
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OperateServiceLineRequest) GetRegionId() string {
    return r.RegionId
}

type OperateServiceLineResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OperateServiceLineResult `json:"result"`
}

type OperateServiceLineResult struct {
    Result bool `json:"result"`
}