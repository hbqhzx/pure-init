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

type ExecuteRasQueryRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 数据库名称 (Optional) */
    DatabaseName *string `json:"databaseName"`

    /* sql脚本  */
    Sql string `json:"sql"`

    /* 用户名称  */
    UserName string `json:"userName"`

    /* 队列名称 (Optional) */
    QueueName *string `json:"queueName"`

    /* 资源名称 (Optional) */
    Source *string `json:"source"`

    /* 回调地址名称 (Optional) */
    CallBackURL *string `json:"callBackURL"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`

    /* 实例拥有者名称 (Optional) */
    InstanceOwnerName *string `json:"instanceOwnerName"`

    /* 是否需要解释 (Optional) */
    IsExplain *string `json:"isExplain"`
}

/*
 * param regionId: 地域ID (Required)
 * param sql: sql脚本 (Required)
 * param userName: 用户名称 (Required)
 * param instanceName: 实例名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExecuteRasQueryRequest(
    regionId string,
    sql string,
    userName string,
    instanceName string,
) *ExecuteRasQueryRequest {

	return &ExecuteRasQueryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dwQuery:executeRasQuery",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Sql: sql,
        UserName: userName,
        InstanceName: instanceName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param databaseName: 数据库名称 (Optional)
 * param sql: sql脚本 (Required)
 * param userName: 用户名称 (Required)
 * param queueName: 队列名称 (Optional)
 * param source: 资源名称 (Optional)
 * param callBackURL: 回调地址名称 (Optional)
 * param instanceName: 实例名称 (Required)
 * param instanceOwnerName: 实例拥有者名称 (Optional)
 * param isExplain: 是否需要解释 (Optional)
 */
func NewExecuteRasQueryRequestWithAllParams(
    regionId string,
    databaseName *string,
    sql string,
    userName string,
    queueName *string,
    source *string,
    callBackURL *string,
    instanceName string,
    instanceOwnerName *string,
    isExplain *string,
) *ExecuteRasQueryRequest {

    return &ExecuteRasQueryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:executeRasQuery",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DatabaseName: databaseName,
        Sql: sql,
        UserName: userName,
        QueueName: queueName,
        Source: source,
        CallBackURL: callBackURL,
        InstanceName: instanceName,
        InstanceOwnerName: instanceOwnerName,
        IsExplain: isExplain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExecuteRasQueryRequestWithoutParam() *ExecuteRasQueryRequest {

    return &ExecuteRasQueryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:executeRasQuery",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ExecuteRasQueryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param databaseName: 数据库名称(Optional) */
func (r *ExecuteRasQueryRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = &databaseName
}

/* param sql: sql脚本(Required) */
func (r *ExecuteRasQueryRequest) SetSql(sql string) {
    r.Sql = sql
}

/* param userName: 用户名称(Required) */
func (r *ExecuteRasQueryRequest) SetUserName(userName string) {
    r.UserName = userName
}

/* param queueName: 队列名称(Optional) */
func (r *ExecuteRasQueryRequest) SetQueueName(queueName string) {
    r.QueueName = &queueName
}

/* param source: 资源名称(Optional) */
func (r *ExecuteRasQueryRequest) SetSource(source string) {
    r.Source = &source
}

/* param callBackURL: 回调地址名称(Optional) */
func (r *ExecuteRasQueryRequest) SetCallBackURL(callBackURL string) {
    r.CallBackURL = &callBackURL
}

/* param instanceName: 实例名称(Required) */
func (r *ExecuteRasQueryRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

/* param instanceOwnerName: 实例拥有者名称(Optional) */
func (r *ExecuteRasQueryRequest) SetInstanceOwnerName(instanceOwnerName string) {
    r.InstanceOwnerName = &instanceOwnerName
}

/* param isExplain: 是否需要解释(Optional) */
func (r *ExecuteRasQueryRequest) SetIsExplain(isExplain string) {
    r.IsExplain = &isExplain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExecuteRasQueryRequest) GetRegionId() string {
    return r.RegionId
}

type ExecuteRasQueryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExecuteRasQueryResult `json:"result"`
}

type ExecuteRasQueryResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data int `json:"data"`
}