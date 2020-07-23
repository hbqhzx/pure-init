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
    jdsfadmin "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdsfadmin/models"
)

type DescriptionRegistriesRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 分页信息要请求的页数 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页信息要请求每页数据的条数 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 集群名称或者集群实例id (Optional) */
    Cluster *string `json:"cluster"`

    /* 用户主账号pin or 用户主账号id (Optional) */
    User *string `json:"user"`

    /* 数据源,线上-ONLINE,历史-HISTORY,默认为线上数据 (Optional) */
    Source *string `json:"source"`
}

/*
 * param regionId: 可用区id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescriptionRegistriesRequest(
    regionId string,
) *DescriptionRegistriesRequest {

	return &DescriptionRegistriesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param pageNumber: 分页信息要请求的页数 (Optional)
 * param pageSize: 分页信息要请求每页数据的条数 (Optional)
 * param cluster: 集群名称或者集群实例id (Optional)
 * param user: 用户主账号pin or 用户主账号id (Optional)
 * param source: 数据源,线上-ONLINE,历史-HISTORY,默认为线上数据 (Optional)
 */
func NewDescriptionRegistriesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    cluster *string,
    user *string,
    source *string,
) *DescriptionRegistriesRequest {

    return &DescriptionRegistriesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Cluster: cluster,
        User: user,
        Source: source,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescriptionRegistriesRequestWithoutParam() *DescriptionRegistriesRequest {

    return &DescriptionRegistriesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DescriptionRegistriesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 分页信息要请求的页数(Optional) */
func (r *DescriptionRegistriesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页信息要请求每页数据的条数(Optional) */
func (r *DescriptionRegistriesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param cluster: 集群名称或者集群实例id(Optional) */
func (r *DescriptionRegistriesRequest) SetCluster(cluster string) {
    r.Cluster = &cluster
}

/* param user: 用户主账号pin or 用户主账号id(Optional) */
func (r *DescriptionRegistriesRequest) SetUser(user string) {
    r.User = &user
}

/* param source: 数据源,线上-ONLINE,历史-HISTORY,默认为线上数据(Optional) */
func (r *DescriptionRegistriesRequest) SetSource(source string) {
    r.Source = &source
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescriptionRegistriesRequest) GetRegionId() string {
    return r.RegionId
}

type DescriptionRegistriesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescriptionRegistriesResult `json:"result"`
}

type DescriptionRegistriesResult struct {
    RegistryInfos []jdsfadmin.Registry `json:"registryInfos"`
    TotalCount int64 `json:"totalCount"`
}