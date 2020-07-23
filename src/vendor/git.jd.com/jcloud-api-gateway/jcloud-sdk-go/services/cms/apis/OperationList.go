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
    cms "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/cms/models"
)

type OperationListRequest struct {

    core.JDCloudRequest

    /* balanceAndArrearsCallback (Optional) */
    BalanceAndArrearsCallback *string `json:"balanceAndArrearsCallback"`

    /* position (Optional) */
    Position *string `json:"position"`

    /* referer (Optional) */
    Referer *string `json:"referer"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOperationListRequest(
) *OperationListRequest {

	return &OperationListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/portal/api/operationList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param balanceAndArrearsCallback: balanceAndArrearsCallback (Optional)
 * param position: position (Optional)
 * param referer: referer (Optional)
 */
func NewOperationListRequestWithAllParams(
    balanceAndArrearsCallback *string,
    position *string,
    referer *string,
) *OperationListRequest {

    return &OperationListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/operationList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        BalanceAndArrearsCallback: balanceAndArrearsCallback,
        Position: position,
        Referer: referer,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOperationListRequestWithoutParam() *OperationListRequest {

    return &OperationListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/operationList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param balanceAndArrearsCallback: balanceAndArrearsCallback(Optional) */
func (r *OperationListRequest) SetBalanceAndArrearsCallback(balanceAndArrearsCallback string) {
    r.BalanceAndArrearsCallback = &balanceAndArrearsCallback
}

/* param position: position(Optional) */
func (r *OperationListRequest) SetPosition(position string) {
    r.Position = &position
}

/* param referer: referer(Optional) */
func (r *OperationListRequest) SetReferer(referer string) {
    r.Referer = &referer
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OperationListRequest) GetRegionId() string {
    return ""
}

type OperationListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OperationListResult `json:"result"`
}

type OperationListResult struct {
    OperationList []cms.Operation `json:"operationList"`
}