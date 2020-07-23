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

type DeleteTraceConfigRequest struct {

    core.JDCloudRequest

    /* 可用区key 例如 cn-north-1  */
    RegionId string `json:"regionId"`

    /* 配置的主键 id  */
    TraceConfigId int `json:"traceConfigId"`
}

/*
 * param regionId: 可用区key 例如 cn-north-1 (Required)
 * param traceConfigId: 配置的主键 id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteTraceConfigRequest(
    regionId string,
    traceConfigId int,
) *DeleteTraceConfigRequest {

	return &DeleteTraceConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/traceconfigs/{traceConfigId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TraceConfigId: traceConfigId,
	}
}

/*
 * param regionId: 可用区key 例如 cn-north-1 (Required)
 * param traceConfigId: 配置的主键 id (Required)
 */
func NewDeleteTraceConfigRequestWithAllParams(
    regionId string,
    traceConfigId int,
) *DeleteTraceConfigRequest {

    return &DeleteTraceConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/traceconfigs/{traceConfigId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TraceConfigId: traceConfigId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteTraceConfigRequestWithoutParam() *DeleteTraceConfigRequest {

    return &DeleteTraceConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/traceconfigs/{traceConfigId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区key 例如 cn-north-1(Required) */
func (r *DeleteTraceConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param traceConfigId: 配置的主键 id(Required) */
func (r *DeleteTraceConfigRequest) SetTraceConfigId(traceConfigId int) {
    r.TraceConfigId = traceConfigId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteTraceConfigRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteTraceConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteTraceConfigResult `json:"result"`
}

type DeleteTraceConfigResult struct {
    DeleteTraceConfig string `json:"deleteTraceConfig"`
}