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

type ExportSubDistributorListRequest struct {

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

    /* 上级渠道商pin (Optional) */
    ParentPin *string `json:"parentPin"`

    /* 上级渠道商ID (Optional) */
    ParentDistributorId *string `json:"parentDistributorId"`

    /* 上级渠道商名称 (Optional) */
    ParentDistributorName *string `json:"parentDistributorName"`

    /* 所属部门(0企业线、1政府线) (Optional) */
    Dept *int `json:"dept"`

    /* 京东云负责人(京东云人员erp或名称) (Optional) */
    Erp *string `json:"erp"`

    /* 入驻条件开始日期 (Optional) */
    SettleTimeBegin *string `json:"settleTimeBegin"`

    /* 入驻条件结束日期 (Optional) */
    SettleTimeEnd *string `json:"settleTimeEnd"`

    /* 当前页序号 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 每页结果数量 (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Offset *int `json:"offset"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExportSubDistributorListRequest(
    regionId string,
) *ExportSubDistributorListRequest {

	return &ExportSubDistributorListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/distributorManage/exportSubDistributorList",
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
 * param parentPin: 上级渠道商pin (Optional)
 * param parentDistributorId: 上级渠道商ID (Optional)
 * param parentDistributorName: 上级渠道商名称 (Optional)
 * param dept: 所属部门(0企业线、1政府线) (Optional)
 * param erp: 京东云负责人(京东云人员erp或名称) (Optional)
 * param settleTimeBegin: 入驻条件开始日期 (Optional)
 * param settleTimeEnd: 入驻条件结束日期 (Optional)
 * param pageIndex: 当前页序号 (Optional)
 * param pageSize: 每页结果数量 (Optional)
 * param offset:  (Optional)
 */
func NewExportSubDistributorListRequestWithAllParams(
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
    parentPin *string,
    parentDistributorId *string,
    parentDistributorName *string,
    dept *int,
    erp *string,
    settleTimeBegin *string,
    settleTimeEnd *string,
    pageIndex *int,
    pageSize *int,
    offset *int,
) *ExportSubDistributorListRequest {

    return &ExportSubDistributorListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/distributorManage/exportSubDistributorList",
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
        ParentPin: parentPin,
        ParentDistributorId: parentDistributorId,
        ParentDistributorName: parentDistributorName,
        Dept: dept,
        Erp: erp,
        SettleTimeBegin: settleTimeBegin,
        SettleTimeEnd: settleTimeEnd,
        PageIndex: pageIndex,
        PageSize: pageSize,
        Offset: offset,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportSubDistributorListRequestWithoutParam() *ExportSubDistributorListRequest {

    return &ExportSubDistributorListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/distributorManage/exportSubDistributorList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ExportSubDistributorListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param distributorId: 渠道商ID(Optional) */
func (r *ExportSubDistributorListRequest) SetDistributorId(distributorId string) {
    r.DistributorId = &distributorId
}

/* param distributorName: 渠道商名称(Optional) */
func (r *ExportSubDistributorListRequest) SetDistributorName(distributorName string) {
    r.DistributorName = &distributorName
}

/* param pin: 京东云账户(Optional) */
func (r *ExportSubDistributorListRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param contractNo: 合同编号(Optional) */
func (r *ExportSubDistributorListRequest) SetContractNo(contractNo string) {
    r.ContractNo = &contractNo
}

/* param businessLicense: 营业执照号(Optional) */
func (r *ExportSubDistributorListRequest) SetBusinessLicense(businessLicense string) {
    r.BusinessLicense = &businessLicense
}

/* param legalRepresentative: 法定代表人(Optional) */
func (r *ExportSubDistributorListRequest) SetLegalRepresentative(legalRepresentative string) {
    r.LegalRepresentative = &legalRepresentative
}

/* param businessLicensePic: 营业执照图片(Optional) */
func (r *ExportSubDistributorListRequest) SetBusinessLicensePic(businessLicensePic string) {
    r.BusinessLicensePic = &businessLicensePic
}

/* param businessDesc: 主营业务描述(Optional) */
func (r *ExportSubDistributorListRequest) SetBusinessDesc(businessDesc string) {
    r.BusinessDesc = &businessDesc
}

/* param workAddress: 办公地址(Optional) */
func (r *ExportSubDistributorListRequest) SetWorkAddress(workAddress string) {
    r.WorkAddress = &workAddress
}

/* param contracter: 联系人姓名(Optional) */
func (r *ExportSubDistributorListRequest) SetContracter(contracter string) {
    r.Contracter = &contracter
}

/* param tel: 联系人电话(Optional) */
func (r *ExportSubDistributorListRequest) SetTel(tel string) {
    r.Tel = &tel
}

/* param email: 邮箱(Optional) */
func (r *ExportSubDistributorListRequest) SetEmail(email string) {
    r.Email = &email
}

/* param region: 所属地域(Optional) */
func (r *ExportSubDistributorListRequest) SetRegion(region string) {
    r.Region = &region
}

/* param settleTime: 入驻日期(一级渠道商手工录入、二级渠道商审批通过日期)(Optional) */
func (r *ExportSubDistributorListRequest) SetSettleTime(settleTime string) {
    r.SettleTime = &settleTime
}

/* param status: 状态(0 审批中、2驳回、1 已入驻、3已停止合作)(Optional) */
func (r *ExportSubDistributorListRequest) SetStatus(status int) {
    r.Status = &status
}

/* param reason: 驳回原因(Optional) */
func (r *ExportSubDistributorListRequest) SetReason(reason string) {
    r.Reason = &reason
}

/* param distributorLevel: 级次(0一级、1 二级)(Optional) */
func (r *ExportSubDistributorListRequest) SetDistributorLevel(distributorLevel int) {
    r.DistributorLevel = &distributorLevel
}

/* param distributorType: 渠道商类型(0合作伙伴、1 渠道代理)(Optional) */
func (r *ExportSubDistributorListRequest) SetDistributorType(distributorType int) {
    r.DistributorType = &distributorType
}

/* param parentPin: 上级渠道商pin(Optional) */
func (r *ExportSubDistributorListRequest) SetParentPin(parentPin string) {
    r.ParentPin = &parentPin
}

/* param parentDistributorId: 上级渠道商ID(Optional) */
func (r *ExportSubDistributorListRequest) SetParentDistributorId(parentDistributorId string) {
    r.ParentDistributorId = &parentDistributorId
}

/* param parentDistributorName: 上级渠道商名称(Optional) */
func (r *ExportSubDistributorListRequest) SetParentDistributorName(parentDistributorName string) {
    r.ParentDistributorName = &parentDistributorName
}

/* param dept: 所属部门(0企业线、1政府线)(Optional) */
func (r *ExportSubDistributorListRequest) SetDept(dept int) {
    r.Dept = &dept
}

/* param erp: 京东云负责人(京东云人员erp或名称)(Optional) */
func (r *ExportSubDistributorListRequest) SetErp(erp string) {
    r.Erp = &erp
}

/* param settleTimeBegin: 入驻条件开始日期(Optional) */
func (r *ExportSubDistributorListRequest) SetSettleTimeBegin(settleTimeBegin string) {
    r.SettleTimeBegin = &settleTimeBegin
}

/* param settleTimeEnd: 入驻条件结束日期(Optional) */
func (r *ExportSubDistributorListRequest) SetSettleTimeEnd(settleTimeEnd string) {
    r.SettleTimeEnd = &settleTimeEnd
}

/* param pageIndex: 当前页序号(Optional) */
func (r *ExportSubDistributorListRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 每页结果数量(Optional) */
func (r *ExportSubDistributorListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param offset: (Optional) */
func (r *ExportSubDistributorListRequest) SetOffset(offset int) {
    r.Offset = &offset
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportSubDistributorListRequest) GetRegionId() string {
    return r.RegionId
}

type ExportSubDistributorListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportSubDistributorListResult `json:"result"`
}

type ExportSubDistributorListResult struct {
    Uri string `json:"uri"`
}