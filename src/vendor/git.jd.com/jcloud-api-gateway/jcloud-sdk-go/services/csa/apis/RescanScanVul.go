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
)

type RescanScanVulRequest struct {

    core.JDCloudRequest

    /* 扫描漏洞ID  */
    ScanVulId string `json:"scanVulId"`

    /* 扫描结果IDs (Optional) */
    ResultIds []int `json:"resultIds"`
}

/*
 * param scanVulId: 扫描漏洞ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRescanScanVulRequest(
    scanVulId string,
) *RescanScanVulRequest {

	return &RescanScanVulRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/scanVuls/{scanVulId}:rescan",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ScanVulId: scanVulId,
	}
}

/*
 * param scanVulId: 扫描漏洞ID (Required)
 * param resultIds: 扫描结果IDs (Optional)
 */
func NewRescanScanVulRequestWithAllParams(
    scanVulId string,
    resultIds []int,
) *RescanScanVulRequest {

    return &RescanScanVulRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/scanVuls/{scanVulId}:rescan",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ScanVulId: scanVulId,
        ResultIds: resultIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRescanScanVulRequestWithoutParam() *RescanScanVulRequest {

    return &RescanScanVulRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/scanVuls/{scanVulId}:rescan",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param scanVulId: 扫描漏洞ID(Required) */
func (r *RescanScanVulRequest) SetScanVulId(scanVulId string) {
    r.ScanVulId = scanVulId
}

/* param resultIds: 扫描结果IDs(Optional) */
func (r *RescanScanVulRequest) SetResultIds(resultIds []int) {
    r.ResultIds = resultIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RescanScanVulRequest) GetRegionId() string {
    return ""
}

type RescanScanVulResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RescanScanVulResult `json:"result"`
}

type RescanScanVulResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}