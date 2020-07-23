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
    kmscap "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/kmscap/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type QueryReportInfoRequest struct {

    core.JDCloudRequest

    /* t           - 时间过滤条件，支持: 1h、6h、12h、1d、7d、30d
keyId       - 根据keyId进行过滤
keyName     - 根据keyName进行过滤
secretId    - 根据secretId进行过滤
secretName  - 根据secretName进行过滤
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryReportInfoRequest(
) *QueryReportInfoRequest {

	return &QueryReportInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/report",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param filters: t           - 时间过滤条件，支持: 1h、6h、12h、1d、7d、30d
keyId       - 根据keyId进行过滤
keyName     - 根据keyName进行过滤
secretId    - 根据secretId进行过滤
secretName  - 根据secretName进行过滤
 (Optional)
 */
func NewQueryReportInfoRequestWithAllParams(
    filters []common.Filter,
) *QueryReportInfoRequest {

    return &QueryReportInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/report",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryReportInfoRequestWithoutParam() *QueryReportInfoRequest {

    return &QueryReportInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/report",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param filters: t           - 时间过滤条件，支持: 1h、6h、12h、1d、7d、30d
keyId       - 根据keyId进行过滤
keyName     - 根据keyName进行过滤
secretId    - 根据secretId进行过滤
secretName  - 根据secretName进行过滤
(Optional) */
func (r *QueryReportInfoRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryReportInfoRequest) GetRegionId() string {
    return ""
}

type QueryReportInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryReportInfoResult `json:"result"`
}

type QueryReportInfoResult struct {
    KeyCount int `json:"keyCount"`
    SecretCount int `json:"secretCount"`
    KeyStats []kmscap.AccessStats `json:"keyStats"`
    SecretStats []kmscap.AccessStats `json:"secretStats"`
}