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
    monitor "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type CreateTemplateRequest struct {

    core.JDCloudRequest

    /* 幂等性校验参数,最长36位  */
    ClientToken string `json:"clientToken"`

    /* 模板描述 (Optional) */
    Description *string `json:"description"`

    /* 规则的资源类型  */
    RuleServiceCode string `json:"ruleServiceCode"`

    /* 模板的资源类型  */
    ServiceCode string `json:"serviceCode"`

    /* 模板名称,长度1-32个字符，只允许中英文、数字、''-''和"_"  */
    TemplateName string `json:"templateName"`

    /* 模板内包含的规则  */
    TemplateRules []monitor.BaseRuleT `json:"templateRules"`
}

/*
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param ruleServiceCode: 规则的资源类型 (Required)
 * param serviceCode: 模板的资源类型 (Required)
 * param templateName: 模板名称,长度1-32个字符，只允许中英文、数字、''-''和"_" (Required)
 * param templateRules: 模板内包含的规则 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTemplateRequest(
    clientToken string,
    ruleServiceCode string,
    serviceCode string,
    templateName string,
    templateRules []monitor.BaseRuleT,
) *CreateTemplateRequest {

	return &CreateTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/template",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ClientToken: clientToken,
        RuleServiceCode: ruleServiceCode,
        ServiceCode: serviceCode,
        TemplateName: templateName,
        TemplateRules: templateRules,
	}
}

/*
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param description: 模板描述 (Optional)
 * param ruleServiceCode: 规则的资源类型 (Required)
 * param serviceCode: 模板的资源类型 (Required)
 * param templateName: 模板名称,长度1-32个字符，只允许中英文、数字、''-''和"_" (Required)
 * param templateRules: 模板内包含的规则 (Required)
 */
func NewCreateTemplateRequestWithAllParams(
    clientToken string,
    description *string,
    ruleServiceCode string,
    serviceCode string,
    templateName string,
    templateRules []monitor.BaseRuleT,
) *CreateTemplateRequest {

    return &CreateTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ClientToken: clientToken,
        Description: description,
        RuleServiceCode: ruleServiceCode,
        ServiceCode: serviceCode,
        TemplateName: templateName,
        TemplateRules: templateRules,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTemplateRequestWithoutParam() *CreateTemplateRequest {

    return &CreateTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param clientToken: 幂等性校验参数,最长36位(Required) */
func (r *CreateTemplateRequest) SetClientToken(clientToken string) {
    r.ClientToken = clientToken
}

/* param description: 模板描述(Optional) */
func (r *CreateTemplateRequest) SetDescription(description string) {
    r.Description = &description
}

/* param ruleServiceCode: 规则的资源类型(Required) */
func (r *CreateTemplateRequest) SetRuleServiceCode(ruleServiceCode string) {
    r.RuleServiceCode = ruleServiceCode
}

/* param serviceCode: 模板的资源类型(Required) */
func (r *CreateTemplateRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param templateName: 模板名称,长度1-32个字符，只允许中英文、数字、''-''和"_"(Required) */
func (r *CreateTemplateRequest) SetTemplateName(templateName string) {
    r.TemplateName = templateName
}

/* param templateRules: 模板内包含的规则(Required) */
func (r *CreateTemplateRequest) SetTemplateRules(templateRules []monitor.BaseRuleT) {
    r.TemplateRules = templateRules
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTemplateRequest) GetRegionId() string {
    return ""
}

type CreateTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTemplateResult `json:"result"`
}

type CreateTemplateResult struct {
    Success bool `json:"success"`
    TemplateId int64 `json:"templateId"`
}