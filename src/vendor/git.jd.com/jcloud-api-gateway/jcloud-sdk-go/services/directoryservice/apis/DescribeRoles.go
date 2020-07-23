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
    directoryservice "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/directoryservice/models"
)

type DescribeRolesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 资源id  */
    DirectoryId string `json:"directoryId"`

    /* 请求列表参数  */
    AdReqVo *directoryservice.AdReqVo `json:"adReqVo"`
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 * param adReqVo: 请求列表参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRolesRequest(
    regionId string,
    directoryId string,
    adReqVo *directoryservice.AdReqVo,
) *DescribeRolesRequest {

	return &DescribeRolesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/directory/{directoryId}/role:describeRoles",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DirectoryId: directoryId,
        AdReqVo: adReqVo,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 * param adReqVo: 请求列表参数 (Required)
 */
func NewDescribeRolesRequestWithAllParams(
    regionId string,
    directoryId string,
    adReqVo *directoryservice.AdReqVo,
) *DescribeRolesRequest {

    return &DescribeRolesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/role:describeRoles",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DirectoryId: directoryId,
        AdReqVo: adReqVo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRolesRequestWithoutParam() *DescribeRolesRequest {

    return &DescribeRolesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/role:describeRoles",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeRolesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param directoryId: 资源id(Required) */
func (r *DescribeRolesRequest) SetDirectoryId(directoryId string) {
    r.DirectoryId = directoryId
}

/* param adReqVo: 请求列表参数(Required) */
func (r *DescribeRolesRequest) SetAdReqVo(adReqVo *directoryservice.AdReqVo) {
    r.AdReqVo = adReqVo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRolesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeRolesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRolesResult `json:"result"`
}

type DescribeRolesResult struct {
    RoleInfo []directoryservice.AdRoleVo `json:"roleInfo"`
}