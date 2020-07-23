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

type SetNodeGroupSizeRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 节点组 ID  */
    NodeGroupId string `json:"nodeGroupId"`

    /* 创建集群请求参数模型  */
    ExpectCount int `json:"expectCount"`
}

/*
 * param regionId: 地域 ID (Required)
 * param nodeGroupId: 节点组 ID (Required)
 * param expectCount: 创建集群请求参数模型 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetNodeGroupSizeRequest(
    regionId string,
    nodeGroupId string,
    expectCount int,
) *SetNodeGroupSizeRequest {

	return &SetNodeGroupSizeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/nodeGroups/{nodeGroupId}:setNodeGroupSize",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NodeGroupId: nodeGroupId,
        ExpectCount: expectCount,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param nodeGroupId: 节点组 ID (Required)
 * param expectCount: 创建集群请求参数模型 (Required)
 */
func NewSetNodeGroupSizeRequestWithAllParams(
    regionId string,
    nodeGroupId string,
    expectCount int,
) *SetNodeGroupSizeRequest {

    return &SetNodeGroupSizeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups/{nodeGroupId}:setNodeGroupSize",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NodeGroupId: nodeGroupId,
        ExpectCount: expectCount,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetNodeGroupSizeRequestWithoutParam() *SetNodeGroupSizeRequest {

    return &SetNodeGroupSizeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups/{nodeGroupId}:setNodeGroupSize",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *SetNodeGroupSizeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param nodeGroupId: 节点组 ID(Required) */
func (r *SetNodeGroupSizeRequest) SetNodeGroupId(nodeGroupId string) {
    r.NodeGroupId = nodeGroupId
}

/* param expectCount: 创建集群请求参数模型(Required) */
func (r *SetNodeGroupSizeRequest) SetExpectCount(expectCount int) {
    r.ExpectCount = expectCount
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetNodeGroupSizeRequest) GetRegionId() string {
    return r.RegionId
}

type SetNodeGroupSizeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetNodeGroupSizeResult `json:"result"`
}

type SetNodeGroupSizeResult struct {
}