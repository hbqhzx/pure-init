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

type DescribeInstanceMetaRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* instanceId  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: regionId (Required)
 * param instanceId: instanceId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceMetaRequest(
    regionId string,
    instanceId string,
) *DescribeInstanceMetaRequest {

	return &DescribeInstanceMetaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/metas",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: regionId (Required)
 * param instanceId: instanceId (Required)
 */
func NewDescribeInstanceMetaRequestWithAllParams(
    regionId string,
    instanceId string,
) *DescribeInstanceMetaRequest {

    return &DescribeInstanceMetaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/metas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceMetaRequestWithoutParam() *DescribeInstanceMetaRequest {

    return &DescribeInstanceMetaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/metas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *DescribeInstanceMetaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: instanceId(Required) */
func (r *DescribeInstanceMetaRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceMetaRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceMetaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceMetaResult `json:"result"`
}

type DescribeInstanceMetaResult struct {
    Pin string `json:"pin"`
    Tags interface{} `json:"tags"`
    Timestamp string `json:"timestamp"`
}