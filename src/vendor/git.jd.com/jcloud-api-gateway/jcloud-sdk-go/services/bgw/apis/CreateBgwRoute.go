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
    bgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/bgw/models"
)

type CreateBgwRouteRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Bgw ID  */
    BgwId string `json:"bgwId"`

    /* 创建边界网关路由参数  */
    BgwRoutes []bgw.BgwRouteSpec `json:"bgwRoutes"`
}

/*
 * param regionId: Region ID (Required)
 * param bgwId: Bgw ID (Required)
 * param bgwRoutes: 创建边界网关路由参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBgwRouteRequest(
    regionId string,
    bgwId string,
    bgwRoutes []bgw.BgwRouteSpec,
) *CreateBgwRouteRequest {

	return &CreateBgwRouteRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bgws/{bgwId}/bgwRoutes",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BgwId: bgwId,
        BgwRoutes: bgwRoutes,
	}
}

/*
 * param regionId: Region ID (Required)
 * param bgwId: Bgw ID (Required)
 * param bgwRoutes: 创建边界网关路由参数 (Required)
 */
func NewCreateBgwRouteRequestWithAllParams(
    regionId string,
    bgwId string,
    bgwRoutes []bgw.BgwRouteSpec,
) *CreateBgwRouteRequest {

    return &CreateBgwRouteRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bgws/{bgwId}/bgwRoutes",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BgwId: bgwId,
        BgwRoutes: bgwRoutes,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBgwRouteRequestWithoutParam() *CreateBgwRouteRequest {

    return &CreateBgwRouteRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bgws/{bgwId}/bgwRoutes",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateBgwRouteRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bgwId: Bgw ID(Required) */
func (r *CreateBgwRouteRequest) SetBgwId(bgwId string) {
    r.BgwId = bgwId
}

/* param bgwRoutes: 创建边界网关路由参数(Required) */
func (r *CreateBgwRouteRequest) SetBgwRoutes(bgwRoutes []bgw.BgwRouteSpec) {
    r.BgwRoutes = bgwRoutes
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBgwRouteRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBgwRouteResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBgwRouteResult `json:"result"`
}

type CreateBgwRouteResult struct {
    BgwRouteIds []string `json:"bgwRouteIds"`
}