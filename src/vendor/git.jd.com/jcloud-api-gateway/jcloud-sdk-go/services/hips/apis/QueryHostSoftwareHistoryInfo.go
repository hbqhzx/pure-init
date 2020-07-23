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

type QueryHostSoftwareHistoryInfoRequest struct {

    core.JDCloudRequest

    /* 主机实例Id  */
    Host string `json:"host"`

    /* 软件名称 (Optional) */
    Name *string `json:"name"`

    /* 软件版本 (Optional) */
    Version *string `json:"version"`

    /* 安装路径 (Optional) */
    Path *string `json:"path"`

    /* 查询天数  */
    Days int `json:"days"`
}

/*
 * param host: 主机实例Id (Required)
 * param days: 查询天数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryHostSoftwareHistoryInfoRequest(
    host string,
    days int,
) *QueryHostSoftwareHistoryInfoRequest {

	return &QueryHostSoftwareHistoryInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/edr:queryHostSoftwareHistoryInfo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Host: host,
        Days: days,
	}
}

/*
 * param host: 主机实例Id (Required)
 * param name: 软件名称 (Optional)
 * param version: 软件版本 (Optional)
 * param path: 安装路径 (Optional)
 * param days: 查询天数 (Required)
 */
func NewQueryHostSoftwareHistoryInfoRequestWithAllParams(
    host string,
    name *string,
    version *string,
    path *string,
    days int,
) *QueryHostSoftwareHistoryInfoRequest {

    return &QueryHostSoftwareHistoryInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/edr:queryHostSoftwareHistoryInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Host: host,
        Name: name,
        Version: version,
        Path: path,
        Days: days,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryHostSoftwareHistoryInfoRequestWithoutParam() *QueryHostSoftwareHistoryInfoRequest {

    return &QueryHostSoftwareHistoryInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/edr:queryHostSoftwareHistoryInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param host: 主机实例Id(Required) */
func (r *QueryHostSoftwareHistoryInfoRequest) SetHost(host string) {
    r.Host = host
}

/* param name: 软件名称(Optional) */
func (r *QueryHostSoftwareHistoryInfoRequest) SetName(name string) {
    r.Name = &name
}

/* param version: 软件版本(Optional) */
func (r *QueryHostSoftwareHistoryInfoRequest) SetVersion(version string) {
    r.Version = &version
}

/* param path: 安装路径(Optional) */
func (r *QueryHostSoftwareHistoryInfoRequest) SetPath(path string) {
    r.Path = &path
}

/* param days: 查询天数(Required) */
func (r *QueryHostSoftwareHistoryInfoRequest) SetDays(days int) {
    r.Days = days
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryHostSoftwareHistoryInfoRequest) GetRegionId() string {
    return ""
}

type QueryHostSoftwareHistoryInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryHostSoftwareHistoryInfoResult `json:"result"`
}

type QueryHostSoftwareHistoryInfoResult struct {
    TimeVector []string `json:"timeVector"`
    Detail []hips.NameCounts `json:"detail"`
    Code int `json:"code"`
    Message string `json:"message"`
}