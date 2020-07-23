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
    ssl "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ssl/models"
)

type DescribeMultiCertDetailsRequest struct {

    core.JDCloudRequest

    /* 证书Id,以逗号分隔多个Id  */
    CertIds string `json:"certIds"`
}

/*
 * param certIds: 证书Id,以逗号分隔多个Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMultiCertDetailsRequest(
    certIds string,
) *DescribeMultiCertDetailsRequest {

	return &DescribeMultiCertDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sslCert:queryMultiCertDetails",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        CertIds: certIds,
	}
}

/*
 * param certIds: 证书Id,以逗号分隔多个Id (Required)
 */
func NewDescribeMultiCertDetailsRequestWithAllParams(
    certIds string,
) *DescribeMultiCertDetailsRequest {

    return &DescribeMultiCertDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert:queryMultiCertDetails",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        CertIds: certIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMultiCertDetailsRequestWithoutParam() *DescribeMultiCertDetailsRequest {

    return &DescribeMultiCertDetailsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert:queryMultiCertDetails",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param certIds: 证书Id,以逗号分隔多个Id(Required) */
func (r *DescribeMultiCertDetailsRequest) SetCertIds(certIds string) {
    r.CertIds = certIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMultiCertDetailsRequest) GetRegionId() string {
    return ""
}

type DescribeMultiCertDetailsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMultiCertDetailsResult `json:"result"`
}

type DescribeMultiCertDetailsResult struct {
    CertDetails []ssl.CertDescDetail `json:"certDetails"`
}