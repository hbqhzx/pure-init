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

type QueryRouteRequest struct {

    core.JDCloudRequest

    /* RouteId  */
    RouteId string `json:"routeId"`
}

/*
 * param routeId: RouteId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryRouteRequest(
    routeId string,
) *QueryRouteRequest {

	return &QueryRouteRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/route/{routeId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RouteId: routeId,
	}
}

/*
 * param routeId: RouteId (Required)
 */
func NewQueryRouteRequestWithAllParams(
    routeId string,
) *QueryRouteRequest {

    return &QueryRouteRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/route/{routeId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RouteId: routeId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryRouteRequestWithoutParam() *QueryRouteRequest {

    return &QueryRouteRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/route/{routeId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param routeId: RouteId(Required) */
func (r *QueryRouteRequest) SetRouteId(routeId string) {
    r.RouteId = routeId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryRouteRequest) GetRegionId() string {
    return ""
}

type QueryRouteResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryRouteResult `json:"result"`
}

type QueryRouteResult struct {
    Id int `json:"id"`
    ServiceName string `json:"serviceName"`
    Host string `json:"host"`
    Uris string `json:"uris"`
    UpstreamUrl string `json:"upstreamUrl"`
    BackendSign bool `json:"backendSign"`
    RegionId string `json:"regionId"`
    CreateTime int64 `json:"createTime"`
}