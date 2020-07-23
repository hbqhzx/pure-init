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

type RefreshUploadUrlRequest struct {

    core.JDCloudRequest

    /* 视频ID  */
    Mid string `json:"mid"`
}

/*
 * param mid: 视频ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRefreshUploadUrlRequest(
    mid string,
) *RefreshUploadUrlRequest {

	return &RefreshUploadUrlRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/uploadTask:refresh",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Mid: mid,
	}
}

/*
 * param mid: 视频ID (Required)
 */
func NewRefreshUploadUrlRequestWithAllParams(
    mid string,
) *RefreshUploadUrlRequest {

    return &RefreshUploadUrlRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/uploadTask:refresh",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Mid: mid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRefreshUploadUrlRequestWithoutParam() *RefreshUploadUrlRequest {

    return &RefreshUploadUrlRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/uploadTask:refresh",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param mid: 视频ID(Required) */
func (r *RefreshUploadUrlRequest) SetMid(mid string) {
    r.Mid = mid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RefreshUploadUrlRequest) GetRegionId() string {
    return ""
}

type RefreshUploadUrlResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RefreshUploadUrlResult `json:"result"`
}

type RefreshUploadUrlResult struct {
    Mid string `json:"mid"`
    UploadUrl string `json:"uploadUrl"`
}