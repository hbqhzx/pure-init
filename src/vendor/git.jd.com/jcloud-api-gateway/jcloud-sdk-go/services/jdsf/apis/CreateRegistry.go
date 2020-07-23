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

type CreateRegistryRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 集群名称  */
    RegistryName string `json:"registryName"`

    /* 集群所在的vpc id  */
    VpcId string `json:"vpcId"`

    /* 子网id  */
    SubnetId string `json:"subnetId"`

    /* 服务实例数量  */
    RegistrySpec string `json:"registrySpec"`

    /* 备注信息 (Optional) */
    Remark *string `json:"remark"`

    /* 版本号  */
    RegistryVersion string `json:"registryVersion"`

    /* 可用区 (Optional) */
    Rz *string `json:"rz"`

    /* 系统分配（SYSTEM)和用户自选(CUSTOMIZE) ,如果用户自选请填写 rz 参数  */
    RzType string `json:"rzType"`
}

/*
 * param regionId: 可用区id (Required)
 * param registryName: 集群名称 (Required)
 * param vpcId: 集群所在的vpc id (Required)
 * param subnetId: 子网id (Required)
 * param registrySpec: 服务实例数量 (Required)
 * param registryVersion: 版本号 (Required)
 * param rzType: 系统分配（SYSTEM)和用户自选(CUSTOMIZE) ,如果用户自选请填写 rz 参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRegistryRequest(
    regionId string,
    registryName string,
    vpcId string,
    subnetId string,
    registrySpec string,
    registryVersion string,
    rzType string,
) *CreateRegistryRequest {

	return &CreateRegistryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RegistryName: registryName,
        VpcId: vpcId,
        SubnetId: subnetId,
        RegistrySpec: registrySpec,
        RegistryVersion: registryVersion,
        RzType: rzType,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param registryName: 集群名称 (Required)
 * param vpcId: 集群所在的vpc id (Required)
 * param subnetId: 子网id (Required)
 * param registrySpec: 服务实例数量 (Required)
 * param remark: 备注信息 (Optional)
 * param registryVersion: 版本号 (Required)
 * param rz: 可用区 (Optional)
 * param rzType: 系统分配（SYSTEM)和用户自选(CUSTOMIZE) ,如果用户自选请填写 rz 参数 (Required)
 */
func NewCreateRegistryRequestWithAllParams(
    regionId string,
    registryName string,
    vpcId string,
    subnetId string,
    registrySpec string,
    remark *string,
    registryVersion string,
    rz *string,
    rzType string,
) *CreateRegistryRequest {

    return &CreateRegistryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryName: registryName,
        VpcId: vpcId,
        SubnetId: subnetId,
        RegistrySpec: registrySpec,
        Remark: remark,
        RegistryVersion: registryVersion,
        Rz: rz,
        RzType: rzType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRegistryRequestWithoutParam() *CreateRegistryRequest {

    return &CreateRegistryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *CreateRegistryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryName: 集群名称(Required) */
func (r *CreateRegistryRequest) SetRegistryName(registryName string) {
    r.RegistryName = registryName
}

/* param vpcId: 集群所在的vpc id(Required) */
func (r *CreateRegistryRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

/* param subnetId: 子网id(Required) */
func (r *CreateRegistryRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

/* param registrySpec: 服务实例数量(Required) */
func (r *CreateRegistryRequest) SetRegistrySpec(registrySpec string) {
    r.RegistrySpec = registrySpec
}

/* param remark: 备注信息(Optional) */
func (r *CreateRegistryRequest) SetRemark(remark string) {
    r.Remark = &remark
}

/* param registryVersion: 版本号(Required) */
func (r *CreateRegistryRequest) SetRegistryVersion(registryVersion string) {
    r.RegistryVersion = registryVersion
}

/* param rz: 可用区(Optional) */
func (r *CreateRegistryRequest) SetRz(rz string) {
    r.Rz = &rz
}

/* param rzType: 系统分配（SYSTEM)和用户自选(CUSTOMIZE) ,如果用户自选请填写 rz 参数(Required) */
func (r *CreateRegistryRequest) SetRzType(rzType string) {
    r.RzType = rzType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRegistryRequest) GetRegionId() string {
    return r.RegionId
}

type CreateRegistryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRegistryResult `json:"result"`
}

type CreateRegistryResult struct {
    CreateResult string `json:"createResult"`
    RegistryId string `json:"registryId"`
}