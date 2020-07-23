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

type StopJobRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*   */
    NamespaceId string `json:"namespaceId"`

    /*   */
    JobId int `json:"jobId"`
}

/*
 * param regionId: Region ID (Required)
 * param namespaceId:  (Required)
 * param jobId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStopJobRequest(
    regionId string,
    namespaceId string,
    jobId int,
) *StopJobRequest {

	return &StopJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/job:stop",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NamespaceId: namespaceId,
        JobId: jobId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param namespaceId:  (Required)
 * param jobId:  (Required)
 */
func NewStopJobRequestWithAllParams(
    regionId string,
    namespaceId string,
    jobId int,
) *StopJobRequest {

    return &StopJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/job:stop",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NamespaceId: namespaceId,
        JobId: jobId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStopJobRequestWithoutParam() *StopJobRequest {

    return &StopJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/job:stop",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *StopJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param namespaceId: (Required) */
func (r *StopJobRequest) SetNamespaceId(namespaceId string) {
    r.NamespaceId = namespaceId
}

/* param jobId: (Required) */
func (r *StopJobRequest) SetJobId(jobId int) {
    r.JobId = jobId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StopJobRequest) GetRegionId() string {
    return r.RegionId
}

type StopJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StopJobResult `json:"result"`
}

type StopJobResult struct {
    Message string `json:"message"`
}