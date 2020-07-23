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
)

type CreatePrivateVirtualInterfaceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* PrivateVif的名称参考  */
    PrivateVifName string `json:"privateVifName"`

    /* 物理连接的Id  */
    ConnectionId string `json:"connectionId"`

    /* 通道的owner:用户pin (Optional) */
    PrivateVifOwner *string `json:"privateVifOwner"`

    /* 边界网关的Id  */
    BgwId string `json:"bgwId"`

    /* 通道的业务vlan,取值范围 [1,4000]  */
    Vlan int `json:"vlan"`

    /* 通道涉及BgpPeer信息  */
    BgpPeers []bgw.BgpPeerSpec `json:"bgpPeers"`

    /* PrivateVif的描述 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param privateVifName: PrivateVif的名称参考 (Required)
 * param connectionId: 物理连接的Id (Required)
 * param bgwId: 边界网关的Id (Required)
 * param vlan: 通道的业务vlan,取值范围 [1,4000] (Required)
 * param bgpPeers: 通道涉及BgpPeer信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreatePrivateVirtualInterfaceRequest(
    regionId string,
    privateVifName string,
    connectionId string,
    bgwId string,
    vlan int,
    bgpPeers []bgw.BgpPeerSpec,
) *CreatePrivateVirtualInterfaceRequest {

	return &CreatePrivateVirtualInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/privateVifs",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PrivateVifName: privateVifName,
        ConnectionId: connectionId,
        BgwId: bgwId,
        Vlan: vlan,
        BgpPeers: bgpPeers,
	}
}

/*
 * param regionId: Region ID (Required)
 * param privateVifName: PrivateVif的名称参考 (Required)
 * param connectionId: 物理连接的Id (Required)
 * param privateVifOwner: 通道的owner:用户pin (Optional)
 * param bgwId: 边界网关的Id (Required)
 * param vlan: 通道的业务vlan,取值范围 [1,4000] (Required)
 * param bgpPeers: 通道涉及BgpPeer信息 (Required)
 * param description: PrivateVif的描述 (Optional)
 */
func NewCreatePrivateVirtualInterfaceRequestWithAllParams(
    regionId string,
    privateVifName string,
    connectionId string,
    privateVifOwner *string,
    bgwId string,
    vlan int,
    bgpPeers []bgw.BgpPeerSpec,
    description *string,
) *CreatePrivateVirtualInterfaceRequest {

    return &CreatePrivateVirtualInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/privateVifs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PrivateVifName: privateVifName,
        ConnectionId: connectionId,
        PrivateVifOwner: privateVifOwner,
        BgwId: bgwId,
        Vlan: vlan,
        BgpPeers: bgpPeers,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreatePrivateVirtualInterfaceRequestWithoutParam() *CreatePrivateVirtualInterfaceRequest {

    return &CreatePrivateVirtualInterfaceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/privateVifs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreatePrivateVirtualInterfaceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param privateVifName: PrivateVif的名称参考(Required) */
func (r *CreatePrivateVirtualInterfaceRequest) SetPrivateVifName(privateVifName string) {
    r.PrivateVifName = privateVifName
}

/* param connectionId: 物理连接的Id(Required) */
func (r *CreatePrivateVirtualInterfaceRequest) SetConnectionId(connectionId string) {
    r.ConnectionId = connectionId
}

/* param privateVifOwner: 通道的owner:用户pin(Optional) */
func (r *CreatePrivateVirtualInterfaceRequest) SetPrivateVifOwner(privateVifOwner string) {
    r.PrivateVifOwner = &privateVifOwner
}

/* param bgwId: 边界网关的Id(Required) */
func (r *CreatePrivateVirtualInterfaceRequest) SetBgwId(bgwId string) {
    r.BgwId = bgwId
}

/* param vlan: 通道的业务vlan,取值范围 [1,4000](Required) */
func (r *CreatePrivateVirtualInterfaceRequest) SetVlan(vlan int) {
    r.Vlan = vlan
}

/* param bgpPeers: 通道涉及BgpPeer信息(Required) */
func (r *CreatePrivateVirtualInterfaceRequest) SetBgpPeers(bgpPeers []bgw.BgpPeerSpec) {
    r.BgpPeers = bgpPeers
}

/* param description: PrivateVif的描述(Optional) */
func (r *CreatePrivateVirtualInterfaceRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreatePrivateVirtualInterfaceRequest) GetRegionId() string {
    return r.RegionId
}

type CreatePrivateVirtualInterfaceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreatePrivateVirtualInterfaceResult `json:"result"`
}

type CreatePrivateVirtualInterfaceResult struct {
    PrivateVifId string `json:"privateVifId"`
}