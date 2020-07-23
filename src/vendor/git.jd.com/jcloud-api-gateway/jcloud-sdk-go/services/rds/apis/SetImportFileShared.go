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

type SetImportFileSharedRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 单库上云文件名  */
    FileName string `json:"fileName"`

    /* 文件是否共享<br>true:共享<br>false:不共享  */
    Shared string `json:"shared"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param fileName: 单库上云文件名 (Required)
 * param shared: 文件是否共享<br>true:共享<br>false:不共享 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetImportFileSharedRequest(
    regionId string,
    instanceId string,
    fileName string,
    shared string,
) *SetImportFileSharedRequest {

	return &SetImportFileSharedRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/importFiles/{fileName}:setShared",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        FileName: fileName,
        Shared: shared,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param fileName: 单库上云文件名 (Required)
 * param shared: 文件是否共享<br>true:共享<br>false:不共享 (Required)
 */
func NewSetImportFileSharedRequestWithAllParams(
    regionId string,
    instanceId string,
    fileName string,
    shared string,
) *SetImportFileSharedRequest {

    return &SetImportFileSharedRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/importFiles/{fileName}:setShared",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        FileName: fileName,
        Shared: shared,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetImportFileSharedRequestWithoutParam() *SetImportFileSharedRequest {

    return &SetImportFileSharedRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/importFiles/{fileName}:setShared",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *SetImportFileSharedRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *SetImportFileSharedRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param fileName: 单库上云文件名(Required) */
func (r *SetImportFileSharedRequest) SetFileName(fileName string) {
    r.FileName = fileName
}

/* param shared: 文件是否共享<br>true:共享<br>false:不共享(Required) */
func (r *SetImportFileSharedRequest) SetShared(shared string) {
    r.Shared = shared
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetImportFileSharedRequest) GetRegionId() string {
    return r.RegionId
}

type SetImportFileSharedResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetImportFileSharedResult `json:"result"`
}

type SetImportFileSharedResult struct {
}