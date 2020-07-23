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
    jke "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jke/models"
)

type DescribeClusterRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 集群 ID  */
    ClusterId string `json:"clusterId"`
}

/*
 * param regionId: 地域 ID (Required)
 * param clusterId: 集群 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeClusterRequest(
    regionId string,
    clusterId string,
) *DescribeClusterRequest {

	return &DescribeClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/clusters/{clusterId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClusterId: clusterId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param clusterId: 集群 ID (Required)
 */
func NewDescribeClusterRequestWithAllParams(
    regionId string,
    clusterId string,
) *DescribeClusterRequest {

    return &DescribeClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClusterId: clusterId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeClusterRequestWithoutParam() *DescribeClusterRequest {

    return &DescribeClusterRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *DescribeClusterRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clusterId: 集群 ID(Required) */
func (r *DescribeClusterRequest) SetClusterId(clusterId string) {
    r.ClusterId = clusterId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeClusterRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeClusterResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeClusterResult `json:"result"`
}

type DescribeClusterResult struct {
    Cluster jke.Cluster `json:"cluster"`
}