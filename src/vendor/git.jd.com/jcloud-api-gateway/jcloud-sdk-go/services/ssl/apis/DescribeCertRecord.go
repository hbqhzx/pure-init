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

type DescribeCertRecordRequest struct {

    core.JDCloudRequest

    /* Record Id  */
    RecordId string `json:"recordId"`
}

/*
 * param recordId: Record Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCertRecordRequest(
    recordId string,
) *DescribeCertRecordRequest {

	return &DescribeCertRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sslRecord/{recordId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RecordId: recordId,
	}
}

/*
 * param recordId: Record Id (Required)
 */
func NewDescribeCertRecordRequestWithAllParams(
    recordId string,
) *DescribeCertRecordRequest {

    return &DescribeCertRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslRecord/{recordId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RecordId: recordId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCertRecordRequestWithoutParam() *DescribeCertRecordRequest {

    return &DescribeCertRecordRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslRecord/{recordId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param recordId: Record Id(Required) */
func (r *DescribeCertRecordRequest) SetRecordId(recordId string) {
    r.RecordId = recordId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCertRecordRequest) GetRegionId() string {
    return ""
}

type DescribeCertRecordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCertRecordResult `json:"result"`
}

type DescribeCertRecordResult struct {
    RecordDescDetail ssl.RecordDescDetail `json:"recordDescDetail"`
}