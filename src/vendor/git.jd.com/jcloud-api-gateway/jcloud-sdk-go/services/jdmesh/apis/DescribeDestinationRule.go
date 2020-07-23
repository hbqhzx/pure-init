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
    jdmesh "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdmesh/models"
)

type DescribeDestinationRuleRequest struct {

    core.JDCloudRequest

    /* 环境名，pre/product  */
    Env string `json:"env"`

    /* 应用ID  */
    AppId string `json:"appId"`
}

/*
 * param env: 环境名，pre/product (Required)
 * param appId: 应用ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDestinationRuleRequest(
    env string,
    appId string,
) *DescribeDestinationRuleRequest {

	return &DescribeDestinationRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/envs/{env}/apps/{appId}/destinationrule",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Env: env,
        AppId: appId,
	}
}

/*
 * param env: 环境名，pre/product (Required)
 * param appId: 应用ID (Required)
 */
func NewDescribeDestinationRuleRequestWithAllParams(
    env string,
    appId string,
) *DescribeDestinationRuleRequest {

    return &DescribeDestinationRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/envs/{env}/apps/{appId}/destinationrule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Env: env,
        AppId: appId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDestinationRuleRequestWithoutParam() *DescribeDestinationRuleRequest {

    return &DescribeDestinationRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/envs/{env}/apps/{appId}/destinationrule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param env: 环境名，pre/product(Required) */
func (r *DescribeDestinationRuleRequest) SetEnv(env string) {
    r.Env = env
}

/* param appId: 应用ID(Required) */
func (r *DescribeDestinationRuleRequest) SetAppId(appId string) {
    r.AppId = appId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDestinationRuleRequest) GetRegionId() string {
    return ""
}

type DescribeDestinationRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDestinationRuleResult `json:"result"`
}

type DescribeDestinationRuleResult struct {
    Destinationrule jdmesh.DestinationRule `json:"destinationrule"`
}