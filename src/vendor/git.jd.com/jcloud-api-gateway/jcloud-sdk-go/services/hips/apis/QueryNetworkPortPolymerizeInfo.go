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
    hips "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/hips/models"
)

type QueryNetworkPortPolymerizeInfoRequest struct {

    core.JDCloudRequest

    /* 监听端口号  */
    Port int `json:"port"`

    /* 主机实例Id (Optional) */
    Host *string `json:"host"`

    /* 监听进程 (Optional) */
    Process *string `json:"process"`
}

/*
 * param port: 监听端口号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryNetworkPortPolymerizeInfoRequest(
    port int,
) *QueryNetworkPortPolymerizeInfoRequest {

	return &QueryNetworkPortPolymerizeInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/edr:queryNetworkPortPolymerizeInfo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Port: port,
	}
}

/*
 * param port: 监听端口号 (Required)
 * param host: 主机实例Id (Optional)
 * param process: 监听进程 (Optional)
 */
func NewQueryNetworkPortPolymerizeInfoRequestWithAllParams(
    port int,
    host *string,
    process *string,
) *QueryNetworkPortPolymerizeInfoRequest {

    return &QueryNetworkPortPolymerizeInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/edr:queryNetworkPortPolymerizeInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Port: port,
        Host: host,
        Process: process,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryNetworkPortPolymerizeInfoRequestWithoutParam() *QueryNetworkPortPolymerizeInfoRequest {

    return &QueryNetworkPortPolymerizeInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/edr:queryNetworkPortPolymerizeInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param port: 监听端口号(Required) */
func (r *QueryNetworkPortPolymerizeInfoRequest) SetPort(port int) {
    r.Port = port
}

/* param host: 主机实例Id(Optional) */
func (r *QueryNetworkPortPolymerizeInfoRequest) SetHost(host string) {
    r.Host = &host
}

/* param process: 监听进程(Optional) */
func (r *QueryNetworkPortPolymerizeInfoRequest) SetProcess(process string) {
    r.Process = &process
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryNetworkPortPolymerizeInfoRequest) GetRegionId() string {
    return ""
}

type QueryNetworkPortPolymerizeInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryNetworkPortPolymerizeInfoResult `json:"result"`
}

type QueryNetworkPortPolymerizeInfoResult struct {
    Data []hips.EdrHostNetworkPortListenInfo `json:"data"`
    TotalCount int `json:"totalCount"`
    Code int `json:"code"`
    Message string `json:"message"`
}