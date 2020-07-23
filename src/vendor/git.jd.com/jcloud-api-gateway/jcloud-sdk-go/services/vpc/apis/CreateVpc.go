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

type CreateVpcRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 私有网络名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。  */
    VpcName string `json:"vpcName"`

    /* 如果为空，则不限制网段，如果不为空，10.0.0.0/8、172.16.0.0/12和192.168.0.0/16及它们包含的子网，且子网掩码长度为16-28之间 (Optional) */
    AddressPrefix *string `json:"addressPrefix"`

    /* vpc描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param vpcName: 私有网络名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVpcRequest(
    regionId string,
    vpcName string,
) *CreateVpcRequest {

	return &CreateVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcs/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcName: vpcName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param vpcName: 私有网络名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Required)
 * param addressPrefix: 如果为空，则不限制网段，如果不为空，10.0.0.0/8、172.16.0.0/12和192.168.0.0/16及它们包含的子网，且子网掩码长度为16-28之间 (Optional)
 * param description: vpc描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional)
 */
func NewCreateVpcRequestWithAllParams(
    regionId string,
    vpcName string,
    addressPrefix *string,
    description *string,
) *CreateVpcRequest {

    return &CreateVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcName: vpcName,
        AddressPrefix: addressPrefix,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVpcRequestWithoutParam() *CreateVpcRequest {

    return &CreateVpcRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateVpcRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpcName: 私有网络名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。(Required) */
func (r *CreateVpcRequest) SetVpcName(vpcName string) {
    r.VpcName = vpcName
}

/* param addressPrefix: 如果为空，则不限制网段，如果不为空，10.0.0.0/8、172.16.0.0/12和192.168.0.0/16及它们包含的子网，且子网掩码长度为16-28之间(Optional) */
func (r *CreateVpcRequest) SetAddressPrefix(addressPrefix string) {
    r.AddressPrefix = &addressPrefix
}

/* param description: vpc描述，允许输入UTF-8编码下的全部字符，不超过256字符。(Optional) */
func (r *CreateVpcRequest) SetDescription(description string) {
    r.Description = &description
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
    VpcId string `json:"vpcId"`
}