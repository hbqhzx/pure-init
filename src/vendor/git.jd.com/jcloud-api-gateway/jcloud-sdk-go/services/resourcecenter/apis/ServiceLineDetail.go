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
    resourcecenter "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/resourcecenter/models"
)

type ServiceLineDetailRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    ServiceId int `json:"serviceId"`
}

/*
 * param regionId:  (Required)
 * param serviceId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewServiceLineDetailRequest(
    regionId string,
    serviceId int,
) *ServiceLineDetailRequest {

	return &ServiceLineDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/serviceline/manage/detail",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceId: serviceId,
	}
}

/*
 * param regionId:  (Required)
 * param serviceId:  (Required)
 */
func NewServiceLineDetailRequestWithAllParams(
    regionId string,
    serviceId int,
) *ServiceLineDetailRequest {

    return &ServiceLineDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceline/manage/detail",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceId: serviceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewServiceLineDetailRequestWithoutParam() *ServiceLineDetailRequest {

    return &ServiceLineDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceline/manage/detail",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ServiceLineDetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serviceId: (Required) */
func (r *ServiceLineDetailRequest) SetServiceId(serviceId int) {
    r.ServiceId = serviceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ServiceLineDetailRequest) GetRegionId() string {
    return r.RegionId
}

type ServiceLineDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ServiceLineDetailResult `json:"result"`
}

type ServiceLineDetailResult struct {
    ServiceLine resourcecenter.ServiceLine `json:"serviceLine"`
    Property []resourcecenter.PropertyVo `json:"property"`
}