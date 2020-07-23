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
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeBgwRoutesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Bgw ID  */
    BgwId string `json:"bgwId"`

    /* bgwRouteIds - Bgw路由的ID列表，支持多个 bgwRouteNextHopType - Bgw上路由的下一跳类型：目前支持vpc（私有网络）、privateVif（专线网关）、hostedPrivateVif（托管网关）、vpcAttachment（vpc接口） origin - Bgw上的路由是静态的还是传播的：static,propagated bgwRouteNexthop - Bgw上路由的下一跳设备资源Id,目前支持vpcId，托管通道Id,专线通道Id,vpc接口Id (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 * param bgwId: Bgw ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBgwRoutesRequest(
    regionId string,
    bgwId string,
) *DescribeBgwRoutesRequest {

	return &DescribeBgwRoutesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bgws/{bgwId}/bgwRoutes/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BgwId: bgwId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param bgwId: Bgw ID (Required)
 * param filters: bgwRouteIds - Bgw路由的ID列表，支持多个 bgwRouteNextHopType - Bgw上路由的下一跳类型：目前支持vpc（私有网络）、privateVif（专线网关）、hostedPrivateVif（托管网关）、vpcAttachment（vpc接口） origin - Bgw上的路由是静态的还是传播的：static,propagated bgwRouteNexthop - Bgw上路由的下一跳设备资源Id,目前支持vpcId，托管通道Id,专线通道Id,vpc接口Id (Optional)
 */
func NewDescribeBgwRoutesRequestWithAllParams(
    regionId string,
    bgwId string,
    filters []common.Filter,
) *DescribeBgwRoutesRequest {

    return &DescribeBgwRoutesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bgws/{bgwId}/bgwRoutes/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BgwId: bgwId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBgwRoutesRequestWithoutParam() *DescribeBgwRoutesRequest {

    return &DescribeBgwRoutesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bgws/{bgwId}/bgwRoutes/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeBgwRoutesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bgwId: Bgw ID(Required) */
func (r *DescribeBgwRoutesRequest) SetBgwId(bgwId string) {
    r.BgwId = bgwId
}

/* param filters: bgwRouteIds - Bgw路由的ID列表，支持多个 bgwRouteNextHopType - Bgw上路由的下一跳类型：目前支持vpc（私有网络）、privateVif（专线网关）、hostedPrivateVif（托管网关）、vpcAttachment（vpc接口） origin - Bgw上的路由是静态的还是传播的：static,propagated bgwRouteNexthop - Bgw上路由的下一跳设备资源Id,目前支持vpcId，托管通道Id,专线通道Id,vpc接口Id(Optional) */
func (r *DescribeBgwRoutesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBgwRoutesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeBgwRoutesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBgwRoutesResult `json:"result"`
}

type DescribeBgwRoutesResult struct {
    BgwRoutes []bgw.BgwRoute `json:"bgwRoutes"`
    StaticCount int `json:"staticCount"`
    PropagatedCount int `json:"propagatedCount"`
    InvalidPropagateCount int `json:"invalidPropagateCount"`
}