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

type CreateSnapshotRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 快照前缀，在4位到32位之间，必须以字母开头，可以包含字母、数字、中划线或者下划线，注意字母不能大写且不能包含其他特殊字符 (Optional) */
    SnapshotPrefix *string `json:"snapshotPrefix"`

    /* 创建快照的索引 (Optional) */
    Indices []string `json:"indices"`
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSnapshotRequest(
    regionId string,
    instanceId string,
) *CreateSnapshotRequest {

	return &CreateSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/snapshots",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 * param snapshotPrefix: 快照前缀，在4位到32位之间，必须以字母开头，可以包含字母、数字、中划线或者下划线，注意字母不能大写且不能包含其他特殊字符 (Optional)
 * param indices: 创建快照的索引 (Optional)
 */
func NewCreateSnapshotRequestWithAllParams(
    regionId string,
    instanceId string,
    snapshotPrefix *string,
    indices []string,
) *CreateSnapshotRequest {

    return &CreateSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/snapshots",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        SnapshotPrefix: snapshotPrefix,
        Indices: indices,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSnapshotRequestWithoutParam() *CreateSnapshotRequest {

    return &CreateSnapshotRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/snapshots",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *CreateSnapshotRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *CreateSnapshotRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param snapshotPrefix: 快照前缀，在4位到32位之间，必须以字母开头，可以包含字母、数字、中划线或者下划线，注意字母不能大写且不能包含其他特殊字符(Optional) */
func (r *CreateSnapshotRequest) SetSnapshotPrefix(snapshotPrefix string) {
    r.SnapshotPrefix = &snapshotPrefix
}

/* param indices: 创建快照的索引(Optional) */
func (r *CreateSnapshotRequest) SetIndices(indices []string) {
    r.Indices = indices
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSnapshotRequest) GetRegionId() string {
    return r.RegionId
}

type CreateSnapshotResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSnapshotResult `json:"result"`
}

type CreateSnapshotResult struct {
    SnapshotId string `json:"snapshotId"`
}