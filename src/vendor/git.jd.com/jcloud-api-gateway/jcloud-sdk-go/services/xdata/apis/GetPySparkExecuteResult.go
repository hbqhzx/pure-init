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

type GetPySparkExecuteResultRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 用户名称  */
    UserName string `json:"userName"`

    /* 查询id  */
    QueryId string `json:"queryId"`
}

/*
 * param regionId: 地域ID (Required)
 * param userName: 用户名称 (Required)
 * param queryId: 查询id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetPySparkExecuteResultRequest(
    regionId string,
    userName string,
    queryId string,
) *GetPySparkExecuteResultRequest {

	return &GetPySparkExecuteResultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dwQuery:getPySparkExecuteResult",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UserName: userName,
        QueryId: queryId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param userName: 用户名称 (Required)
 * param queryId: 查询id (Required)
 */
func NewGetPySparkExecuteResultRequestWithAllParams(
    regionId string,
    userName string,
    queryId string,
) *GetPySparkExecuteResultRequest {

    return &GetPySparkExecuteResultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:getPySparkExecuteResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UserName: userName,
        QueryId: queryId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetPySparkExecuteResultRequestWithoutParam() *GetPySparkExecuteResultRequest {

    return &GetPySparkExecuteResultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:getPySparkExecuteResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetPySparkExecuteResultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param userName: 用户名称(Required) */
func (r *GetPySparkExecuteResultRequest) SetUserName(userName string) {
    r.UserName = userName
}

/* param queryId: 查询id(Required) */
func (r *GetPySparkExecuteResultRequest) SetQueryId(queryId string) {
    r.QueryId = queryId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetPySparkExecuteResultRequest) GetRegionId() string {
    return r.RegionId
}

type GetPySparkExecuteResultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetPySparkExecuteResultResult `json:"result"`
}

type GetPySparkExecuteResultResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}