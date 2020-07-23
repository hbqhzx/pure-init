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
    jdfusion "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdfusion/models"
)

type CreateVpcSubnetRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 创建subnet  */
    Subnet *jdfusion.CreateSubnet `json:"subnet"`
}

/*
 * param regionId: 地域ID (Required)
 * param subnet: 创建subnet (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVpcSubnetRequest(
    regionId string,
    subnet *jdfusion.CreateSubnet,
) *CreateVpcSubnetRequest {

	return &CreateVpcSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpc_subnets",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Subnet: subnet,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param subnet: 创建subnet (Required)
 */
func NewCreateVpcSubnetRequestWithAllParams(
    regionId string,
    subnet *jdfusion.CreateSubnet,
) *CreateVpcSubnetRequest {

    return &CreateVpcSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_subnets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Subnet: subnet,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVpcSubnetRequestWithoutParam() *CreateVpcSubnetRequest {

    return &CreateVpcSubnetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_subnets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateVpcSubnetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param subnet: 创建subnet(Required) */
func (r *CreateVpcSubnetRequest) SetSubnet(subnet *jdfusion.CreateSubnet) {
    r.Subnet = subnet
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVpcSubnetRequest) GetRegionId() string {
    return r.RegionId
}

type CreateVpcSubnetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateVpcSubnetResult `json:"result"`
}

type CreateVpcSubnetResult struct {
    Task jdfusion.ResourceTFInfo `json:"task"`
}