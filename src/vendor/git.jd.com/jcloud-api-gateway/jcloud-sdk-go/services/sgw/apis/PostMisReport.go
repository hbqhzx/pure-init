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

type PostMisReportRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* waf 实例id  */
    InstanceId string `json:"instanceId"`

    /*  (Optional) */
    Type *int `json:"type"`

    /*  (Optional) */
    Id *string `json:"id"`
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPostMisReportRequest(
    regionId string,
    instanceId string,
) *PostMisReportRequest {

	return &PostMisReportRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:postMisReport",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 * param type:  (Optional)
 * param id:  (Optional)
 */
func NewPostMisReportRequestWithAllParams(
    regionId string,
    instanceId string,
    type *int,
    id *string,
) *PostMisReportRequest {

    return &PostMisReportRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:postMisReport",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Type: type,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPostMisReportRequestWithoutParam() *PostMisReportRequest {

    return &PostMisReportRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:postMisReport",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *PostMisReportRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: waf 实例id(Required) */
func (r *PostMisReportRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param type: (Optional) */
func (r *PostMisReportRequest) SetType(type int) {
    r.Type = &type
}

/* param id: (Optional) */
func (r *PostMisReportRequest) SetId(id string) {
    r.Id = &id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PostMisReportRequest) GetRegionId() string {
    return r.RegionId
}

type PostMisReportResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PostMisReportResult `json:"result"`
}

type PostMisReportResult struct {
}