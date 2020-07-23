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

type ModifyOneLevelDistributorInfoRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 渠道商ID (Optional) */
    DistributorId *string `json:"distributorId"`

    /* 渠道商名称 (Optional) */
    DistributorName *string `json:"distributorName"`

    /* 京东云账户 (Optional) */
    Pin *string `json:"pin"`

    /* 合同编号 (Optional) */
    ContractNo *string `json:"contractNo"`

    /* 营业执照号 (Optional) */
    BusinessLicense *string `json:"businessLicense"`

    /* 法定代表人 (Optional) */
    LegalRepresentative *string `json:"legalRepresentative"`

    /* 营业执照图片 (Optional) */
    BusinessLicensePic *string `json:"businessLicensePic"`

    /* 主营业务描述 (Optional) */
    BusinessDesc *string `json:"businessDesc"`

    /* 办公地址 (Optional) */
    WorkAddress *string `json:"workAddress"`

    /* 联系人姓名 (Optional) */
    Contracter *string `json:"contracter"`

    /* 联系人电话 (Optional) */
    Tel *string `json:"tel"`

    /* 邮箱 (Optional) */
    Email *string `json:"email"`

    /* 所属地域 (Optional) */
    Region *string `json:"region"`

    /* 入驻日期(一级渠道商手工录入、二级渠道商审批通过日期) (Optional) */
    SettleTime *string `json:"settleTime"`

    /* 状态(0 审批中、2驳回、1 已入驻、3已停止合作) (Optional) */
    Status *int `json:"status"`

    /* 驳回原因 (Optional) */
    Reason *string `json:"reason"`

    /* 级次(0一级、1 二级) (Optional) */
    DistributorLevel *int `json:"distributorLevel"`

    /* 渠道商类型(0合作伙伴、1 渠道代理) (Optional) */
    DistributorType *int `json:"distributorType"`

    /* 上级渠道商ID (Optional) */
    ParentDistributorId *string `json:"parentDistributorId"`

    /* 银行开户名 (Optional) */
    BankCompanyName *string `json:"bankCompanyName"`

    /* 银行账户 (Optional) */
    BankCardNo *string `json:"bankCardNo"`

    /* 开户行支行名称 (Optional) */
    BankBranchName *string `json:"bankBranchName"`

    /* 开户行支行联行号 (Optional) */
    BankBranchNo *string `json:"bankBranchNo"`

    /* 合同主体 (Optional) */
    ContractSubject *string `json:"contractSubject"`

    /* 所属部门(0企业线、1政府线) (Optional) */
    Dept *int `json:"dept"`

    /* 京东云负责人(京东云人员erp或名称) (Optional) */
    Erp *string `json:"erp"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyOneLevelDistributorInfoRequest(
    regionId string,
) *ModifyOneLevelDistributorInfoRequest {

	return &ModifyOneLevelDistributorInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/distributorManage/modifyOneLevelDistributorInfo",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param distributorId: 渠道商ID (Optional)
 * param distributorName: 渠道商名称 (Optional)
 * param pin: 京东云账户 (Optional)
 * param contractNo: 合同编号 (Optional)
 * param businessLicense: 营业执照号 (Optional)
 * param legalRepresentative: 法定代表人 (Optional)
 * param businessLicensePic: 营业执照图片 (Optional)
 * param businessDesc: 主营业务描述 (Optional)
 * param workAddress: 办公地址 (Optional)
 * param contracter: 联系人姓名 (Optional)
 * param tel: 联系人电话 (Optional)
 * param email: 邮箱 (Optional)
 * param region: 所属地域 (Optional)
 * param settleTime: 入驻日期(一级渠道商手工录入、二级渠道商审批通过日期) (Optional)
 * param status: 状态(0 审批中、2驳回、1 已入驻、3已停止合作) (Optional)
 * param reason: 驳回原因 (Optional)
 * param distributorLevel: 级次(0一级、1 二级) (Optional)
 * param distributorType: 渠道商类型(0合作伙伴、1 渠道代理) (Optional)
 * param parentDistributorId: 上级渠道商ID (Optional)
 * param bankCompanyName: 银行开户名 (Optional)
 * param bankCardNo: 银行账户 (Optional)
 * param bankBranchName: 开户行支行名称 (Optional)
 * param bankBranchNo: 开户行支行联行号 (Optional)
 * param contractSubject: 合同主体 (Optional)
 * param dept: 所属部门(0企业线、1政府线) (Optional)
 * param erp: 京东云负责人(京东云人员erp或名称) (Optional)
 */
func NewModifyOneLevelDistributorInfoRequestWithAllParams(
    regionId string,
    distributorId *string,
    distributorName *string,
    pin *string,
    contractNo *string,
    businessLicense *string,
    legalRepresentative *string,
    businessLicensePic *string,
    businessDesc *string,
    workAddress *string,
    contracter *string,
    tel *string,
    email *string,
    region *string,
    settleTime *string,
    status *int,
    reason *string,
    distributorLevel *int,
    distributorType *int,
    parentDistributorId *string,
    bankCompanyName *string,
    bankCardNo *string,
    bankBranchName *string,
    bankBranchNo *string,
    contractSubject *string,
    dept *int,
    erp *string,
) *ModifyOneLevelDistributorInfoRequest {

    return &ModifyOneLevelDistributorInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/distributorManage/modifyOneLevelDistributorInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DistributorId: distributorId,
        DistributorName: distributorName,
        Pin: pin,
        ContractNo: contractNo,
        BusinessLicense: businessLicense,
        LegalRepresentative: legalRepresentative,
        BusinessLicensePic: businessLicensePic,
        BusinessDesc: businessDesc,
        WorkAddress: workAddress,
        Contracter: contracter,
        Tel: tel,
        Email: email,
        Region: region,
        SettleTime: settleTime,
        Status: status,
        Reason: reason,
        DistributorLevel: distributorLevel,
        DistributorType: distributorType,
        ParentDistributorId: parentDistributorId,
        BankCompanyName: bankCompanyName,
        BankCardNo: bankCardNo,
        BankBranchName: bankBranchName,
        BankBranchNo: bankBranchNo,
        ContractSubject: contractSubject,
        Dept: dept,
        Erp: erp,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyOneLevelDistributorInfoRequestWithoutParam() *ModifyOneLevelDistributorInfoRequest {

    return &ModifyOneLevelDistributorInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/distributorManage/modifyOneLevelDistributorInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ModifyOneLevelDistributorInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param distributorId: 渠道商ID(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetDistributorId(distributorId string) {
    r.DistributorId = &distributorId
}

/* param distributorName: 渠道商名称(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetDistributorName(distributorName string) {
    r.DistributorName = &distributorName
}

/* param pin: 京东云账户(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param contractNo: 合同编号(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetContractNo(contractNo string) {
    r.ContractNo = &contractNo
}

/* param businessLicense: 营业执照号(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetBusinessLicense(businessLicense string) {
    r.BusinessLicense = &businessLicense
}

/* param legalRepresentative: 法定代表人(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetLegalRepresentative(legalRepresentative string) {
    r.LegalRepresentative = &legalRepresentative
}

/* param businessLicensePic: 营业执照图片(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetBusinessLicensePic(businessLicensePic string) {
    r.BusinessLicensePic = &businessLicensePic
}

/* param businessDesc: 主营业务描述(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetBusinessDesc(businessDesc string) {
    r.BusinessDesc = &businessDesc
}

/* param workAddress: 办公地址(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetWorkAddress(workAddress string) {
    r.WorkAddress = &workAddress
}

/* param contracter: 联系人姓名(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetContracter(contracter string) {
    r.Contracter = &contracter
}

/* param tel: 联系人电话(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetTel(tel string) {
    r.Tel = &tel
}

/* param email: 邮箱(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetEmail(email string) {
    r.Email = &email
}

/* param region: 所属地域(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetRegion(region string) {
    r.Region = &region
}

/* param settleTime: 入驻日期(一级渠道商手工录入、二级渠道商审批通过日期)(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetSettleTime(settleTime string) {
    r.SettleTime = &settleTime
}

/* param status: 状态(0 审批中、2驳回、1 已入驻、3已停止合作)(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetStatus(status int) {
    r.Status = &status
}

/* param reason: 驳回原因(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetReason(reason string) {
    r.Reason = &reason
}

/* param distributorLevel: 级次(0一级、1 二级)(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetDistributorLevel(distributorLevel int) {
    r.DistributorLevel = &distributorLevel
}

/* param distributorType: 渠道商类型(0合作伙伴、1 渠道代理)(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetDistributorType(distributorType int) {
    r.DistributorType = &distributorType
}

/* param parentDistributorId: 上级渠道商ID(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetParentDistributorId(parentDistributorId string) {
    r.ParentDistributorId = &parentDistributorId
}

/* param bankCompanyName: 银行开户名(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetBankCompanyName(bankCompanyName string) {
    r.BankCompanyName = &bankCompanyName
}

/* param bankCardNo: 银行账户(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetBankCardNo(bankCardNo string) {
    r.BankCardNo = &bankCardNo
}

/* param bankBranchName: 开户行支行名称(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetBankBranchName(bankBranchName string) {
    r.BankBranchName = &bankBranchName
}

/* param bankBranchNo: 开户行支行联行号(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetBankBranchNo(bankBranchNo string) {
    r.BankBranchNo = &bankBranchNo
}

/* param contractSubject: 合同主体(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetContractSubject(contractSubject string) {
    r.ContractSubject = &contractSubject
}

/* param dept: 所属部门(0企业线、1政府线)(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetDept(dept int) {
    r.Dept = &dept
}

/* param erp: 京东云负责人(京东云人员erp或名称)(Optional) */
func (r *ModifyOneLevelDistributorInfoRequest) SetErp(erp string) {
    r.Erp = &erp
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyOneLevelDistributorInfoRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyOneLevelDistributorInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyOneLevelDistributorInfoResult `json:"result"`
}

type ModifyOneLevelDistributorInfoResult struct {
    Result bool `json:"result"`
}