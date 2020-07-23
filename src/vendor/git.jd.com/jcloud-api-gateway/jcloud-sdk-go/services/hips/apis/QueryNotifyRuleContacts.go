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
    hips "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/hips/models"
)

type QueryNotifyRuleContactsRequest struct {

    core.JDCloudRequest

    /*   */
    RuleId int `json:"ruleId"`
}

/*
 * param ruleId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryNotifyRuleContactsRequest(
    ruleId int,
) *QueryNotifyRuleContactsRequest {

	return &QueryNotifyRuleContactsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/notifyRules/{ruleId}:queryContacts",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RuleId: ruleId,
	}
}

/*
 * param ruleId:  (Required)
 */
func NewQueryNotifyRuleContactsRequestWithAllParams(
    ruleId int,
) *QueryNotifyRuleContactsRequest {

    return &QueryNotifyRuleContactsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/notifyRules/{ruleId}:queryContacts",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RuleId: ruleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryNotifyRuleContactsRequestWithoutParam() *QueryNotifyRuleContactsRequest {

    return &QueryNotifyRuleContactsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/notifyRules/{ruleId}:queryContacts",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param ruleId: (Required) */
func (r *QueryNotifyRuleContactsRequest) SetRuleId(ruleId int) {
    r.RuleId = ruleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryNotifyRuleContactsRequest) GetRegionId() string {
    return ""
}

type QueryNotifyRuleContactsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryNotifyRuleContactsResult `json:"result"`
}

type QueryNotifyRuleContactsResult struct {
    NotifyContacts hips.NotifyContacts `json:"notifyContacts"`
}