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

type AddBusinessConfigRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 业务系统  */
    SystemCode string `json:"systemCode"`

    /* 结算对象类型  */
    SettleObjectType string `json:"settleObjectType"`

    /* 费用类型  */
    CostType string `json:"costType"`

    /* 结算系统编号  */
    RfSystemId string `json:"rfSystemId"`

    /* 结算业务  */
    SettlementBizId int `json:"settlementBizId"`

    /* 结算分组  */
    SettlementGroupId int `json:"settlementGroupId"`
}

/*
 * param regionId:  (Required)
 * param systemCode: 业务系统 (Required)
 * param settleObjectType: 结算对象类型 (Required)
 * param costType: 费用类型 (Required)
 * param rfSystemId: 结算系统编号 (Required)
 * param settlementBizId: 结算业务 (Required)
 * param settlementGroupId: 结算分组 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddBusinessConfigRequest(
    regionId string,
    systemCode string,
    settleObjectType string,
    costType string,
    rfSystemId string,
    settlementBizId int,
    settlementGroupId int,
) *AddBusinessConfigRequest {

	return &AddBusinessConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/businessConfig:add",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SystemCode: systemCode,
        SettleObjectType: settleObjectType,
        CostType: costType,
        RfSystemId: rfSystemId,
        SettlementBizId: settlementBizId,
        SettlementGroupId: settlementGroupId,
	}
}

/*
 * param regionId:  (Required)
 * param systemCode: 业务系统 (Required)
 * param settleObjectType: 结算对象类型 (Required)
 * param costType: 费用类型 (Required)
 * param rfSystemId: 结算系统编号 (Required)
 * param settlementBizId: 结算业务 (Required)
 * param settlementGroupId: 结算分组 (Required)
 */
func NewAddBusinessConfigRequestWithAllParams(
    regionId string,
    systemCode string,
    settleObjectType string,
    costType string,
    rfSystemId string,
    settlementBizId int,
    settlementGroupId int,
) *AddBusinessConfigRequest {

    return &AddBusinessConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/businessConfig:add",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SystemCode: systemCode,
        SettleObjectType: settleObjectType,
        CostType: costType,
        RfSystemId: rfSystemId,
        SettlementBizId: settlementBizId,
        SettlementGroupId: settlementGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddBusinessConfigRequestWithoutParam() *AddBusinessConfigRequest {

    return &AddBusinessConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/businessConfig:add",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AddBusinessConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param systemCode: 业务系统(Required) */
func (r *AddBusinessConfigRequest) SetSystemCode(systemCode string) {
    r.SystemCode = systemCode
}

/* param settleObjectType: 结算对象类型(Required) */
func (r *AddBusinessConfigRequest) SetSettleObjectType(settleObjectType string) {
    r.SettleObjectType = settleObjectType
}

/* param costType: 费用类型(Required) */
func (r *AddBusinessConfigRequest) SetCostType(costType string) {
    r.CostType = costType
}

/* param rfSystemId: 结算系统编号(Required) */
func (r *AddBusinessConfigRequest) SetRfSystemId(rfSystemId string) {
    r.RfSystemId = rfSystemId
}

/* param settlementBizId: 结算业务(Required) */
func (r *AddBusinessConfigRequest) SetSettlementBizId(settlementBizId int) {
    r.SettlementBizId = settlementBizId
}

/* param settlementGroupId: 结算分组(Required) */
func (r *AddBusinessConfigRequest) SetSettlementGroupId(settlementGroupId int) {
    r.SettlementGroupId = settlementGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddBusinessConfigRequest) GetRegionId() string {
    return r.RegionId
}

type AddBusinessConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddBusinessConfigResult `json:"result"`
}

type AddBusinessConfigResult struct {
    Result bool `json:"result"`
}