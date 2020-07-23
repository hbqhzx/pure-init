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

type UpdateSettleObjectFromBusinessRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 结算对象pin  */
    Pin string `json:"pin"`

    /* 结算对象名称 (Optional) */
    Name *string `json:"name"`

    /* 所属业务系统  */
    SystemCode string `json:"systemCode"`

    /* 结算对象类型  */
    SettleObjectType string `json:"settleObjectType"`

    /* 联系人 (Optional) */
    Contacter *string `json:"contacter"`

    /* 联系电话 (Optional) */
    Phone *string `json:"phone"`

    /* 联系邮箱 (Optional) */
    Email *string `json:"email"`

    /* 业务负责人 (Optional) */
    BusinessErp *string `json:"businessErp"`

    /* 结算周期 1：天 2：周 3：半月 4：月 5：季 (Optional) */
    SettlementCycle *int `json:"settlementCycle"`

    /* 银行开户名 (Optional) */
    BankCompanyName *string `json:"bankCompanyName"`

    /* 银行账户 (Optional) */
    BankCardNo *string `json:"bankCardNo"`

    /* 开户行支行名称 (Optional) */
    BankBranchName *string `json:"bankBranchName"`

    /* 开户行支行联行号 (Optional) */
    BankBranchNo *string `json:"bankBranchNo"`
}

/*
 * param regionId:  (Required)
 * param pin: 结算对象pin (Required)
 * param systemCode: 所属业务系统 (Required)
 * param settleObjectType: 结算对象类型 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateSettleObjectFromBusinessRequest(
    regionId string,
    pin string,
    systemCode string,
    settleObjectType string,
) *UpdateSettleObjectFromBusinessRequest {

	return &UpdateSettleObjectFromBusinessRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/settlementObject/business:modify",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
        SystemCode: systemCode,
        SettleObjectType: settleObjectType,
	}
}

/*
 * param regionId:  (Required)
 * param pin: 结算对象pin (Required)
 * param name: 结算对象名称 (Optional)
 * param systemCode: 所属业务系统 (Required)
 * param settleObjectType: 结算对象类型 (Required)
 * param contacter: 联系人 (Optional)
 * param phone: 联系电话 (Optional)
 * param email: 联系邮箱 (Optional)
 * param businessErp: 业务负责人 (Optional)
 * param settlementCycle: 结算周期 1：天 2：周 3：半月 4：月 5：季 (Optional)
 * param bankCompanyName: 银行开户名 (Optional)
 * param bankCardNo: 银行账户 (Optional)
 * param bankBranchName: 开户行支行名称 (Optional)
 * param bankBranchNo: 开户行支行联行号 (Optional)
 */
func NewUpdateSettleObjectFromBusinessRequestWithAllParams(
    regionId string,
    pin string,
    name *string,
    systemCode string,
    settleObjectType string,
    contacter *string,
    phone *string,
    email *string,
    businessErp *string,
    settlementCycle *int,
    bankCompanyName *string,
    bankCardNo *string,
    bankBranchName *string,
    bankBranchNo *string,
) *UpdateSettleObjectFromBusinessRequest {

    return &UpdateSettleObjectFromBusinessRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlementObject/business:modify",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        Name: name,
        SystemCode: systemCode,
        SettleObjectType: settleObjectType,
        Contacter: contacter,
        Phone: phone,
        Email: email,
        BusinessErp: businessErp,
        SettlementCycle: settlementCycle,
        BankCompanyName: bankCompanyName,
        BankCardNo: bankCardNo,
        BankBranchName: bankBranchName,
        BankBranchNo: bankBranchNo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateSettleObjectFromBusinessRequestWithoutParam() *UpdateSettleObjectFromBusinessRequest {

    return &UpdateSettleObjectFromBusinessRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlementObject/business:modify",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *UpdateSettleObjectFromBusinessRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 结算对象pin(Required) */
func (r *UpdateSettleObjectFromBusinessRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param name: 结算对象名称(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetName(name string) {
    r.Name = &name
}

/* param systemCode: 所属业务系统(Required) */
func (r *UpdateSettleObjectFromBusinessRequest) SetSystemCode(systemCode string) {
    r.SystemCode = systemCode
}

/* param settleObjectType: 结算对象类型(Required) */
func (r *UpdateSettleObjectFromBusinessRequest) SetSettleObjectType(settleObjectType string) {
    r.SettleObjectType = settleObjectType
}

/* param contacter: 联系人(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetContacter(contacter string) {
    r.Contacter = &contacter
}

/* param phone: 联系电话(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetPhone(phone string) {
    r.Phone = &phone
}

/* param email: 联系邮箱(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetEmail(email string) {
    r.Email = &email
}

/* param businessErp: 业务负责人(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetBusinessErp(businessErp string) {
    r.BusinessErp = &businessErp
}

/* param settlementCycle: 结算周期 1：天 2：周 3：半月 4：月 5：季(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetSettlementCycle(settlementCycle int) {
    r.SettlementCycle = &settlementCycle
}

/* param bankCompanyName: 银行开户名(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetBankCompanyName(bankCompanyName string) {
    r.BankCompanyName = &bankCompanyName
}

/* param bankCardNo: 银行账户(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetBankCardNo(bankCardNo string) {
    r.BankCardNo = &bankCardNo
}

/* param bankBranchName: 开户行支行名称(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetBankBranchName(bankBranchName string) {
    r.BankBranchName = &bankBranchName
}

/* param bankBranchNo: 开户行支行联行号(Optional) */
func (r *UpdateSettleObjectFromBusinessRequest) SetBankBranchNo(bankBranchNo string) {
    r.BankBranchNo = &bankBranchNo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateSettleObjectFromBusinessRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateSettleObjectFromBusinessResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateSettleObjectFromBusinessResult `json:"result"`
}

type UpdateSettleObjectFromBusinessResult struct {
    Result bool `json:"result"`
}