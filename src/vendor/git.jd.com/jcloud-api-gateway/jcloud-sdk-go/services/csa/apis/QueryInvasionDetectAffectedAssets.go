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
    csa "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/csa/models"
)

type QueryInvasionDetectAffectedAssetsRequest struct {

    core.JDCloudRequest

    /* 天数，1:24时，7:7天，30:30天  */
    TimeSpan int `json:"timeSpan"`
}

/*
 * param timeSpan: 天数，1:24时，7:7天，30:30天 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryInvasionDetectAffectedAssetsRequest(
    timeSpan int,
) *QueryInvasionDetectAffectedAssetsRequest {

	return &QueryInvasionDetectAffectedAssetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/hostScreen:invasionDetectAffectedAssets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TimeSpan: timeSpan,
	}
}

/*
 * param timeSpan: 天数，1:24时，7:7天，30:30天 (Required)
 */
func NewQueryInvasionDetectAffectedAssetsRequestWithAllParams(
    timeSpan int,
) *QueryInvasionDetectAffectedAssetsRequest {

    return &QueryInvasionDetectAffectedAssetsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/hostScreen:invasionDetectAffectedAssets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TimeSpan: timeSpan,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryInvasionDetectAffectedAssetsRequestWithoutParam() *QueryInvasionDetectAffectedAssetsRequest {

    return &QueryInvasionDetectAffectedAssetsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/hostScreen:invasionDetectAffectedAssets",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param timeSpan: 天数，1:24时，7:7天，30:30天(Required) */
func (r *QueryInvasionDetectAffectedAssetsRequest) SetTimeSpan(timeSpan int) {
    r.TimeSpan = timeSpan
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryInvasionDetectAffectedAssetsRequest) GetRegionId() string {
    return ""
}

type QueryInvasionDetectAffectedAssetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryInvasionDetectAffectedAssetsResult `json:"result"`
}

type QueryInvasionDetectAffectedAssetsResult struct {
    Data []csa.HostNameCount `json:"data"`
    AffectedAssetsProportion int `json:"affectedAssetsProportion"`
    Code int `json:"code"`
    Message string `json:"message"`
}