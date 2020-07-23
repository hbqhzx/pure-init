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
    jdro "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdro/models"
)

type CreateStackRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /*   */
    Environment *jdro.Environment `json:"environment"`

    /* 模板, JSON对象  */
    Template interface{} `json:"template"`
}

/*
 * param regionId: 地域 ID (Required)
 * param environment:  (Required)
 * param template: 模板, JSON对象 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateStackRequest(
    regionId string,
    environment *jdro.Environment,
    template interface{},
) *CreateStackRequest {

	return &CreateStackRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/stacks",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Environment: environment,
        Template: template,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param environment:  (Required)
 * param template: 模板, JSON对象 (Required)
 */
func NewCreateStackRequestWithAllParams(
    regionId string,
    environment *jdro.Environment,
    template interface{},
) *CreateStackRequest {

    return &CreateStackRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stacks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Environment: environment,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateStackRequestWithoutParam() *CreateStackRequest {

    return &CreateStackRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stacks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *CreateStackRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param environment: (Required) */
func (r *CreateStackRequest) SetEnvironment(environment *jdro.Environment) {
    r.Environment = environment
}

/* param template: 模板, JSON对象(Required) */
func (r *CreateStackRequest) SetTemplate(template interface{}) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateStackRequest) GetRegionId() string {
    return r.RegionId
}

type CreateStackResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateStackResult `json:"result"`
}

type CreateStackResult struct {
    StackID string `json:"stackID"`
}