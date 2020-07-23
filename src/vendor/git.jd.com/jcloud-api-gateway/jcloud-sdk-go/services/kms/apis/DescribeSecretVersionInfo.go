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
    kms "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/kms/models"
)

type DescribeSecretVersionInfoRequest struct {

    core.JDCloudRequest

    /* 机密ID  */
    SecretId string `json:"secretId"`

    /* 机密版本  */
    Version string `json:"version"`
}

/*
 * param secretId: 机密ID (Required)
 * param version: 机密版本 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSecretVersionInfoRequest(
    secretId string,
    version string,
) *DescribeSecretVersionInfoRequest {

	return &DescribeSecretVersionInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/secret/{secretId}/version/{version}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        SecretId: secretId,
        Version: version,
	}
}

/*
 * param secretId: 机密ID (Required)
 * param version: 机密版本 (Required)
 */
func NewDescribeSecretVersionInfoRequestWithAllParams(
    secretId string,
    version string,
) *DescribeSecretVersionInfoRequest {

    return &DescribeSecretVersionInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}/version/{version}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        SecretId: secretId,
        Version: version,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSecretVersionInfoRequestWithoutParam() *DescribeSecretVersionInfoRequest {

    return &DescribeSecretVersionInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}/version/{version}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param secretId: 机密ID(Required) */
func (r *DescribeSecretVersionInfoRequest) SetSecretId(secretId string) {
    r.SecretId = secretId
}

/* param version: 机密版本(Required) */
func (r *DescribeSecretVersionInfoRequest) SetVersion(version string) {
    r.Version = version
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSecretVersionInfoRequest) GetRegionId() string {
    return ""
}

type DescribeSecretVersionInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSecretVersionInfoResult `json:"result"`
}

type DescribeSecretVersionInfoResult struct {
    SecretVersionItem kms.SecretVersionItem `json:"secretVersionItem"`
}