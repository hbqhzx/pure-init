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

type MandatoryDisassociateElasticIpRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkInterface ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 指定解绑的弹性Ip Id (Optional) */
    ElasticIpId *string `json:"elasticIpId"`

    /* 指定解绑的弹性Ip地址 (Optional) */
    ElasticIpAddress *string `json:"elasticIpAddress"`
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewMandatoryDisassociateElasticIpRequest(
    regionId string,
    networkInterfaceId string,
) *MandatoryDisassociateElasticIpRequest {

	return &MandatoryDisassociateElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}:mandatoryDisassociateElasticIp",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 * param elasticIpId: 指定解绑的弹性Ip Id (Optional)
 * param elasticIpAddress: 指定解绑的弹性Ip地址 (Optional)
 */
func NewMandatoryDisassociateElasticIpRequestWithAllParams(
    regionId string,
    networkInterfaceId string,
    elasticIpId *string,
    elasticIpAddress *string,
) *MandatoryDisassociateElasticIpRequest {

    return &MandatoryDisassociateElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}:mandatoryDisassociateElasticIp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
        ElasticIpId: elasticIpId,
        ElasticIpAddress: elasticIpAddress,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewMandatoryDisassociateElasticIpRequestWithoutParam() *MandatoryDisassociateElasticIpRequest {

    return &MandatoryDisassociateElasticIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}:mandatoryDisassociateElasticIp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *MandatoryDisassociateElasticIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkInterfaceId: networkInterface ID(Required) */
func (r *MandatoryDisassociateElasticIpRequest) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}

/* param elasticIpId: 指定解绑的弹性Ip Id(Optional) */
func (r *MandatoryDisassociateElasticIpRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = &elasticIpId
}

/* param elasticIpAddress: 指定解绑的弹性Ip地址(Optional) */
func (r *MandatoryDisassociateElasticIpRequest) SetElasticIpAddress(elasticIpAddress string) {
    r.ElasticIpAddress = &elasticIpAddress
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r MandatoryDisassociateElasticIpRequest) GetRegionId() string {
    return r.RegionId
}

type MandatoryDisassociateElasticIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result MandatoryDisassociateElasticIpResult `json:"result"`
}

type MandatoryDisassociateElasticIpResult struct {
}