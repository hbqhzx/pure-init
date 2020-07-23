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
    ompopenapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ompopenapi/models"
)

type OnlineApiRequest struct {

    core.JDCloudRequest

    /* 业务线名称  */
    Name string `json:"name"`

    /* 配置API列表  */
    Apis []ompopenapi.Api `json:"apis"`
}

/*
 * param name: 业务线名称 (Required)
 * param apis: 配置API列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOnlineApiRequest(
    name string,
    apis []ompopenapi.Api,
) *OnlineApiRequest {

	return &OnlineApiRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/api/{name}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Name: name,
        Apis: apis,
	}
}

/*
 * param name: 业务线名称 (Required)
 * param apis: 配置API列表 (Required)
 */
func NewOnlineApiRequestWithAllParams(
    name string,
    apis []ompopenapi.Api,
) *OnlineApiRequest {

    return &OnlineApiRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/api/{name}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Name: name,
        Apis: apis,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOnlineApiRequestWithoutParam() *OnlineApiRequest {

    return &OnlineApiRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/api/{name}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param name: 业务线名称(Required) */
func (r *OnlineApiRequest) SetName(name string) {
    r.Name = name
}

/* param apis: 配置API列表(Required) */
func (r *OnlineApiRequest) SetApis(apis []ompopenapi.Api) {
    r.Apis = apis
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OnlineApiRequest) GetRegionId() string {
    return ""
}

type OnlineApiResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OnlineApiResult `json:"result"`
}

type OnlineApiResult struct {
}