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
    pipeline "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/pipeline/models"
)

type UpdatePipelineRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 流水线任务 uuid  */
    Uuid string `json:"uuid"`

    /*  (Optional) */
    Data *pipeline.PipelineParams `json:"data"`
}

/*
 * param regionId: Region ID (Required)
 * param uuid: 流水线任务 uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdatePipelineRequest(
    regionId string,
    uuid string,
) *UpdatePipelineRequest {

	return &UpdatePipelineRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pipeline/{uuid}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Uuid: uuid,
	}
}

/*
 * param regionId: Region ID (Required)
 * param uuid: 流水线任务 uuid (Required)
 * param data:  (Optional)
 */
func NewUpdatePipelineRequestWithAllParams(
    regionId string,
    uuid string,
    data *pipeline.PipelineParams,
) *UpdatePipelineRequest {

    return &UpdatePipelineRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pipeline/{uuid}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Uuid: uuid,
        Data: data,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdatePipelineRequestWithoutParam() *UpdatePipelineRequest {

    return &UpdatePipelineRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pipeline/{uuid}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdatePipelineRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param uuid: 流水线任务 uuid(Required) */
func (r *UpdatePipelineRequest) SetUuid(uuid string) {
    r.Uuid = uuid
}

/* param data: (Optional) */
func (r *UpdatePipelineRequest) SetData(data *pipeline.PipelineParams) {
    r.Data = data
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdatePipelineRequest) GetRegionId() string {
    return r.RegionId
}

type UpdatePipelineResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdatePipelineResult `json:"result"`
}

type UpdatePipelineResult struct {
    Uuid string `json:"uuid"`
    Result bool `json:"result"`
}