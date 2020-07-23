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
    iam "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/iam/models"
)

type DescribeSubUserPossessionPoliciesRequest struct {

    core.JDCloudRequest

    /* 子账户用户名  */
    SubUser string `json:"subUser"`
}

/*
 * param subUser: 子账户用户名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSubUserPossessionPoliciesRequest(
    subUser string,
) *DescribeSubUserPossessionPoliciesRequest {

	return &DescribeSubUserPossessionPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/subUser/{subUser}/policies:possess",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        SubUser: subUser,
	}
}

/*
 * param subUser: 子账户用户名 (Required)
 */
func NewDescribeSubUserPossessionPoliciesRequestWithAllParams(
    subUser string,
) *DescribeSubUserPossessionPoliciesRequest {

    return &DescribeSubUserPossessionPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/subUser/{subUser}/policies:possess",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        SubUser: subUser,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSubUserPossessionPoliciesRequestWithoutParam() *DescribeSubUserPossessionPoliciesRequest {

    return &DescribeSubUserPossessionPoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/subUser/{subUser}/policies:possess",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param subUser: 子账户用户名(Required) */
func (r *DescribeSubUserPossessionPoliciesRequest) SetSubUser(subUser string) {
    r.SubUser = subUser
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSubUserPossessionPoliciesRequest) GetRegionId() string {
    return ""
}

type DescribeSubUserPossessionPoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSubUserPossessionPoliciesResult `json:"result"`
}

type DescribeSubUserPossessionPoliciesResult struct {
    Total int `json:"total"`
    Policies []iam.Policy `json:"policies"`
}