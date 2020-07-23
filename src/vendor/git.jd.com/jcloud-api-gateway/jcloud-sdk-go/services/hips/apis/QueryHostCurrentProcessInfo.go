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

type QueryHostCurrentProcessInfoRequest struct {

    core.JDCloudRequest

    /* 主机实例Id  */
    Host string `json:"host"`

    /* 查询的进程名 (Optional) */
    Name *string `json:"name"`

    /* 进程对应启动文件的MD5 (Optional) */
    Md5 *string `json:"md5"`

    /* 进程启动用户 (Optional) */
    User *string `json:"user"`

    /* 进程启动参数 (Optional) */
    Args *string `json:"args"`
}

/*
 * param host: 主机实例Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryHostCurrentProcessInfoRequest(
    host string,
) *QueryHostCurrentProcessInfoRequest {

	return &QueryHostCurrentProcessInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/edr:queryHostCurrentProcessInfo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Host: host,
	}
}

/*
 * param host: 主机实例Id (Required)
 * param name: 查询的进程名 (Optional)
 * param md5: 进程对应启动文件的MD5 (Optional)
 * param user: 进程启动用户 (Optional)
 * param args: 进程启动参数 (Optional)
 */
func NewQueryHostCurrentProcessInfoRequestWithAllParams(
    host string,
    name *string,
    md5 *string,
    user *string,
    args *string,
) *QueryHostCurrentProcessInfoRequest {

    return &QueryHostCurrentProcessInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/edr:queryHostCurrentProcessInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Host: host,
        Name: name,
        Md5: md5,
        User: user,
        Args: args,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryHostCurrentProcessInfoRequestWithoutParam() *QueryHostCurrentProcessInfoRequest {

    return &QueryHostCurrentProcessInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/edr:queryHostCurrentProcessInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param host: 主机实例Id(Required) */
func (r *QueryHostCurrentProcessInfoRequest) SetHost(host string) {
    r.Host = host
}

/* param name: 查询的进程名(Optional) */
func (r *QueryHostCurrentProcessInfoRequest) SetName(name string) {
    r.Name = &name
}

/* param md5: 进程对应启动文件的MD5(Optional) */
func (r *QueryHostCurrentProcessInfoRequest) SetMd5(md5 string) {
    r.Md5 = &md5
}

/* param user: 进程启动用户(Optional) */
func (r *QueryHostCurrentProcessInfoRequest) SetUser(user string) {
    r.User = &user
}

/* param args: 进程启动参数(Optional) */
func (r *QueryHostCurrentProcessInfoRequest) SetArgs(args string) {
    r.Args = &args
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryHostCurrentProcessInfoRequest) GetRegionId() string {
    return ""
}

type QueryHostCurrentProcessInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryHostCurrentProcessInfoResult `json:"result"`
}

type QueryHostCurrentProcessInfoResult struct {
    Data []hips.EdrHostCurrentProcessInfo `json:"data"`
    TotalCount int `json:"totalCount"`
    Code int `json:"code"`
    Message string `json:"message"`
}