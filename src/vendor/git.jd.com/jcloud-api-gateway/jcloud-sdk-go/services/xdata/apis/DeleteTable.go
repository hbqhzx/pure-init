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

type DeleteTableRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 数据表名  */
    TableName string `json:"tableName"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`

    /* 数据库名称  */
    DatabaseName string `json:"databaseName"`
}

/*
 * param regionId: 地域ID (Required)
 * param tableName: 数据表名 (Required)
 * param instanceName: 实例名称 (Required)
 * param databaseName: 数据库名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteTableRequest(
    regionId string,
    tableName string,
    instanceName string,
    databaseName string,
) *DeleteTableRequest {

	return &DeleteTableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dwTable/{tableName}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TableName: tableName,
        InstanceName: instanceName,
        DatabaseName: databaseName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param tableName: 数据表名 (Required)
 * param instanceName: 实例名称 (Required)
 * param databaseName: 数据库名称 (Required)
 */
func NewDeleteTableRequestWithAllParams(
    regionId string,
    tableName string,
    instanceName string,
    databaseName string,
) *DeleteTableRequest {

    return &DeleteTableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwTable/{tableName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TableName: tableName,
        InstanceName: instanceName,
        DatabaseName: databaseName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteTableRequestWithoutParam() *DeleteTableRequest {

    return &DeleteTableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwTable/{tableName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteTableRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param tableName: 数据表名(Required) */
func (r *DeleteTableRequest) SetTableName(tableName string) {
    r.TableName = tableName
}

/* param instanceName: 实例名称(Required) */
func (r *DeleteTableRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

/* param databaseName: 数据库名称(Required) */
func (r *DeleteTableRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = databaseName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteTableRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteTableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteTableResult `json:"result"`
}

type DeleteTableResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}