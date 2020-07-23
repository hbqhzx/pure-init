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

type DeleteSubnetRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Subnet ID  */
    SubnetId string `json:"subnetId"`
}

/*
 * param regionId: Region ID (Required)
 * param subnetId: Subnet ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteSubnetRequest(
    regionId string,
    subnetId string,
) *DeleteSubnetRequest {

	return &DeleteSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnets/{subnetId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param subnetId: Subnet ID (Required)
 */
func NewDeleteSubnetRequestWithAllParams(
    regionId string,
    subnetId string,
) *DeleteSubnetRequest {

    return &DeleteSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SubnetId: subnetId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteSubnetRequestWithoutParam() *DeleteSubnetRequest {

    return &DeleteSubnetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteSubnetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param subnetId: Subnet ID(Required) */
func (r *DeleteSubnetRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteSubnetRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteSubnetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteSubnetResult `json:"result"`
}

type DeleteSubnetResult struct {
}