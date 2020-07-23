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

type ResumeHubInstanceRequest struct {

    core.JDCloudRequest

    /* IoT Hub实例ID信息  */
    InstanceId string `json:"instanceId"`
}

/*
 * param instanceId: IoT Hub实例ID信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewResumeHubInstanceRequest(
    instanceId string,
) *ResumeHubInstanceRequest {

	return &ResumeHubInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/instance/{instanceId}/resume",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        InstanceId: instanceId,
	}
}

/*
 * param instanceId: IoT Hub实例ID信息 (Required)
 */
func NewResumeHubInstanceRequestWithAllParams(
    instanceId string,
) *ResumeHubInstanceRequest {

    return &ResumeHubInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/instance/{instanceId}/resume",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewResumeHubInstanceRequestWithoutParam() *ResumeHubInstanceRequest {

    return &ResumeHubInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/instance/{instanceId}/resume",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param instanceId: IoT Hub实例ID信息(Required) */
func (r *ResumeHubInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ResumeHubInstanceRequest) GetRegionId() string {
    return ""
}

type ResumeHubInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ResumeHubInstanceResult `json:"result"`
}

type ResumeHubInstanceResult struct {
}