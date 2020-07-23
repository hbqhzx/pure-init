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

type CreateVpcRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 创建VPC  */
    Vpc *jdfusion.VpcInfo `json:"vpc"`
}

/*
 * param regionId: 地域ID (Required)
 * param vpc: 创建VPC (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVpcRequest(
    regionId string,
    vpc *jdfusion.VpcInfo,
) *CreateVpcRequest {

	return &CreateVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpc_vpcs",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Vpc: vpc,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param vpc: 创建VPC (Required)
 */
func NewCreateVpcRequestWithAllParams(
    regionId string,
    vpc *jdfusion.VpcInfo,
) *CreateVpcRequest {

    return &CreateVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_vpcs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Vpc: vpc,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVpcRequestWithoutParam() *CreateVpcRequest {

    return &CreateVpcRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_vpcs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateVpcRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpc: 创建VPC(Required) */
func (r *CreateVpcRequest) SetVpc(vpc *jdfusion.VpcInfo) {
    r.Vpc = vpc
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVpcRequest) GetRegionId() string {
    return r.RegionId
}

type CreateVpcResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateVpcResult `json:"result"`
}

type CreateVpcResult struct {
    Task jdfusion.ResourceTFInfo `json:"task"`
}