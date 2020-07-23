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

type CancelPySparkJobRequest struct {

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
func NewCancelPySparkJobRequest(
    regionId string,
    userName string,
    queryId string,
) *CancelPySparkJobRequest {

	return &CancelPySparkJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dwQuery:cancelPySparkQuery",
			Method:  "POST",
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
func NewCancelPySparkJobRequestWithAllParams(
    regionId string,
    userName string,
    queryId string,
) *CancelPySparkJobRequest {

    return &CancelPySparkJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:cancelPySparkQuery",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UserName: userName,
        QueryId: queryId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCancelPySparkJobRequestWithoutParam() *CancelPySparkJobRequest {

    return &CancelPySparkJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dwQuery:cancelPySparkQuery",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CancelPySparkJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param userName: 用户名称(Required) */
func (r *CancelPySparkJobRequest) SetUserName(userName string) {
    r.UserName = userName
}

/* param queryId: 查询id(Required) */
func (r *CancelPySparkJobRequest) SetQueryId(queryId string) {
    r.QueryId = queryId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CancelPySparkJobRequest) GetRegionId() string {
    return r.RegionId
}

type CancelPySparkJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CancelPySparkJobResult `json:"result"`
}

type CancelPySparkJobResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}