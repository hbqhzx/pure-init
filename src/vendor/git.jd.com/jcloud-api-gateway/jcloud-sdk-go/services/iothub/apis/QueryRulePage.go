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
    iothub "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/iothub/models"
)

type QueryRulePageRequest struct {

    core.JDCloudRequest

    /* 实例Id (Optional) */
    InstanceId *string `json:"instanceId"`

    /* 当前页 (Optional) */
    NowPage *int `json:"nowPage"`

    /* 每页数据条数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryRulePageRequest(
) *QueryRulePageRequest {

	return &QueryRulePageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/rule:queryPage",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param instanceId: 实例Id (Optional)
 * param nowPage: 当前页 (Optional)
 * param pageSize: 每页数据条数 (Optional)
 */
func NewQueryRulePageRequestWithAllParams(
    instanceId *string,
    nowPage *int,
    pageSize *int,
) *QueryRulePageRequest {

    return &QueryRulePageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/rule:queryPage",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        InstanceId: instanceId,
        NowPage: nowPage,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryRulePageRequestWithoutParam() *QueryRulePageRequest {

    return &QueryRulePageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/rule:queryPage",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param instanceId: 实例Id(Optional) */
func (r *QueryRulePageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}

/* param nowPage: 当前页(Optional) */
func (r *QueryRulePageRequest) SetNowPage(nowPage int) {
    r.NowPage = &nowPage
}

/* param pageSize: 每页数据条数(Optional) */
func (r *QueryRulePageRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryRulePageRequest) GetRegionId() string {
    return ""
}

type QueryRulePageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryRulePageResult `json:"result"`
}

type QueryRulePageResult struct {
    PageSize int `json:"pageSize"`
    NowPage int `json:"nowPage"`
    TotalSize int `json:"totalSize"`
    TotalPage int `json:"totalPage"`
    Data []iothub.RuleBaseInfo `json:"data"`
}