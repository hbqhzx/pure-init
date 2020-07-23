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

type QueryThreatEventRequest struct {

    core.JDCloudRequest

    /* threatEvent事件ID  */
    ThreatEventId string `json:"threatEventId"`
}

/*
 * param threatEventId: threatEvent事件ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryThreatEventRequest(
    threatEventId string,
) *QueryThreatEventRequest {

	return &QueryThreatEventRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/threatEvents/{threatEventId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ThreatEventId: threatEventId,
	}
}

/*
 * param threatEventId: threatEvent事件ID (Required)
 */
func NewQueryThreatEventRequestWithAllParams(
    threatEventId string,
) *QueryThreatEventRequest {

    return &QueryThreatEventRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/threatEvents/{threatEventId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ThreatEventId: threatEventId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryThreatEventRequestWithoutParam() *QueryThreatEventRequest {

    return &QueryThreatEventRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/threatEvents/{threatEventId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param threatEventId: threatEvent事件ID(Required) */
func (r *QueryThreatEventRequest) SetThreatEventId(threatEventId string) {
    r.ThreatEventId = threatEventId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryThreatEventRequest) GetRegionId() string {
    return ""
}

type QueryThreatEventResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryThreatEventResult `json:"result"`
}

type QueryThreatEventResult struct {
    ThreatEventDetail csa.ThreatEventDetail `json:"threatEventDetail"`
}