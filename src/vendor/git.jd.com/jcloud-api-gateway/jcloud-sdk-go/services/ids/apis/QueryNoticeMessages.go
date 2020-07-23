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
    ids "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ids/models"
)

type QueryNoticeMessagesRequest struct {

    core.JDCloudRequest

    /* 默认值604800，显示近7天的消息 (Optional) */
    Timespan *string `json:"timespan"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryNoticeMessagesRequest(
) *QueryNoticeMessagesRequest {

	return &QueryNoticeMessagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/noticeMessages",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param timespan: 默认值604800，显示近7天的消息 (Optional)
 */
func NewQueryNoticeMessagesRequestWithAllParams(
    timespan *string,
) *QueryNoticeMessagesRequest {

    return &QueryNoticeMessagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/noticeMessages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Timespan: timespan,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryNoticeMessagesRequestWithoutParam() *QueryNoticeMessagesRequest {

    return &QueryNoticeMessagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/noticeMessages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param timespan: 默认值604800，显示近7天的消息(Optional) */
func (r *QueryNoticeMessagesRequest) SetTimespan(timespan string) {
    r.Timespan = &timespan
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryNoticeMessagesRequest) GetRegionId() string {
    return ""
}

type QueryNoticeMessagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryNoticeMessagesResult `json:"result"`
}

type QueryNoticeMessagesResult struct {
    NoticeMessages []ids.MessageNotice `json:"noticeMessages"`
}