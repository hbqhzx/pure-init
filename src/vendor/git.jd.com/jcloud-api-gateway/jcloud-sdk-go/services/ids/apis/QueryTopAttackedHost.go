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

type QueryTopAttackedHostRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    Ip *string `json:"ip"`

    /* 大于0的时间戳距离1970年1月1日,单位毫秒  */
    StartTime int `json:"startTime"`

    /* 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒  */
    EndTime int `json:"endTime"`
}

/*
 * param startTime: 大于0的时间戳距离1970年1月1日,单位毫秒 (Required)
 * param endTime: 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryTopAttackedHostRequest(
    startTime int,
    endTime int,
) *QueryTopAttackedHostRequest {

	return &QueryTopAttackedHostRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/attacks:queryTopAttackedHost",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param ip:  (Optional)
 * param startTime: 大于0的时间戳距离1970年1月1日,单位毫秒 (Required)
 * param endTime: 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒 (Required)
 */
func NewQueryTopAttackedHostRequestWithAllParams(
    ip *string,
    startTime int,
    endTime int,
) *QueryTopAttackedHostRequest {

    return &QueryTopAttackedHostRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/attacks:queryTopAttackedHost",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Ip: ip,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryTopAttackedHostRequestWithoutParam() *QueryTopAttackedHostRequest {

    return &QueryTopAttackedHostRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/attacks:queryTopAttackedHost",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param ip: (Optional) */
func (r *QueryTopAttackedHostRequest) SetIp(ip string) {
    r.Ip = &ip
}

/* param startTime: 大于0的时间戳距离1970年1月1日,单位毫秒(Required) */
func (r *QueryTopAttackedHostRequest) SetStartTime(startTime int) {
    r.StartTime = startTime
}

/* param endTime: 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒(Required) */
func (r *QueryTopAttackedHostRequest) SetEndTime(endTime int) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryTopAttackedHostRequest) GetRegionId() string {
    return ""
}

type QueryTopAttackedHostResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryTopAttackedHostResult `json:"result"`
}

type QueryTopAttackedHostResult struct {
    TopAttackedHosts []ids.EntryCount `json:"topAttackedHosts"`
}