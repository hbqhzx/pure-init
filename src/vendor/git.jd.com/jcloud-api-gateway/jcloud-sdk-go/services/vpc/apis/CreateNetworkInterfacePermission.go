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

type CreateNetworkInterfacePermissionRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 弹性网卡ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 授信用户,需要存在于京东云许可的服务账号名单中 (Optional) */
    TrustUser *string `json:"trustUser"`

    /* 授权策略, 授权后，该弹性网卡可以关联的服务端账号的资源类型，取值范围，instance-attach：可以关联服务端账号的实例资源，eip-associate：可以关联服务端账号的弹性公网IP资源 (Optional) */
    Policy []string `json:"policy"`
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: 弹性网卡ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateNetworkInterfacePermissionRequest(
    regionId string,
    networkInterfaceId string,
) *CreateNetworkInterfacePermissionRequest {

	return &CreateNetworkInterfacePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkInterfacePermissions/",
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
 * param networkInterfaceId: 弹性网卡ID (Required)
 * param trustUser: 授信用户,需要存在于京东云许可的服务账号名单中 (Optional)
 * param policy: 授权策略, 授权后，该弹性网卡可以关联的服务端账号的资源类型，取值范围，instance-attach：可以关联服务端账号的实例资源，eip-associate：可以关联服务端账号的弹性公网IP资源 (Optional)
 */
func NewCreateNetworkInterfacePermissionRequestWithAllParams(
    regionId string,
    networkInterfaceId string,
    trustUser *string,
    policy []string,
) *CreateNetworkInterfacePermissionRequest {

    return &CreateNetworkInterfacePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfacePermissions/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
        TrustUser: trustUser,
        Policy: policy,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateNetworkInterfacePermissionRequestWithoutParam() *CreateNetworkInterfacePermissionRequest {

    return &CreateNetworkInterfacePermissionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfacePermissions/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateNetworkInterfacePermissionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkInterfaceId: 弹性网卡ID(Required) */
func (r *CreateNetworkInterfacePermissionRequest) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}

/* param trustUser: 授信用户,需要存在于京东云许可的服务账号名单中(Optional) */
func (r *CreateNetworkInterfacePermissionRequest) SetTrustUser(trustUser string) {
    r.TrustUser = &trustUser
}

/* param policy: 授权策略, 授权后，该弹性网卡可以关联的服务端账号的资源类型，取值范围，instance-attach：可以关联服务端账号的实例资源，eip-associate：可以关联服务端账号的弹性公网IP资源(Optional) */
func (r *CreateNetworkInterfacePermissionRequest) SetPolicy(policy []string) {
    r.Policy = policy
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateNetworkInterfacePermissionRequest) GetRegionId() string {
    return r.RegionId
}

type CreateNetworkInterfacePermissionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateNetworkInterfacePermissionResult `json:"result"`
}

type CreateNetworkInterfacePermissionResult struct {
    NetworkInterfacePermissionId string `json:"networkInterfacePermissionId"`
}