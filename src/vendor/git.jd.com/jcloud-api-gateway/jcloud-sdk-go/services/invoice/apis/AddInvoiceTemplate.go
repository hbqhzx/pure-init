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

type AddInvoiceTemplateRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* id (Optional) */
    Id *int64 `json:"id"`

    /* 发票类型 (Optional) */
    InvoiceType *int `json:"invoiceType"`

    /* 发票抬头 (Optional) */
    InvoiceTitle *string `json:"invoiceTitle"`

    /* 发票内容 (Optional) */
    InvoiceContent *string `json:"invoiceContent"`

    /* 税号 (Optional) */
    TaxId *string `json:"taxId"`

    /* 公司 (Optional) */
    Company *string `json:"company"`

    /* 手机号 (Optional) */
    Phone *string `json:"phone"`

    /* 开户行 (Optional) */
    Bank *string `json:"bank"`

    /* 开户账号 (Optional) */
    Account *string `json:"account"`

    /* 开户行地址 (Optional) */
    Address *string `json:"address"`

    /* 邮箱 (Optional) */
    Email *string `json:"email"`

    /* vatId (Optional) */
    VatId *int `json:"vatId"`

    /* type (Optional) */
    Type *int `json:"type"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddInvoiceTemplateRequest(
    regionId string,
) *AddInvoiceTemplateRequest {

	return &AddInvoiceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/invoice_template",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param id: id (Optional)
 * param invoiceType: 发票类型 (Optional)
 * param invoiceTitle: 发票抬头 (Optional)
 * param invoiceContent: 发票内容 (Optional)
 * param taxId: 税号 (Optional)
 * param company: 公司 (Optional)
 * param phone: 手机号 (Optional)
 * param bank: 开户行 (Optional)
 * param account: 开户账号 (Optional)
 * param address: 开户行地址 (Optional)
 * param email: 邮箱 (Optional)
 * param vatId: vatId (Optional)
 * param type_: type (Optional)
 */
func NewAddInvoiceTemplateRequestWithAllParams(
    regionId string,
    id *int64,
    invoiceType *int,
    invoiceTitle *string,
    invoiceContent *string,
    taxId *string,
    company *string,
    phone *string,
    bank *string,
    account *string,
    address *string,
    email *string,
    vatId *int,
    type_ *int,
) *AddInvoiceTemplateRequest {

    return &AddInvoiceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoice_template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        InvoiceType: invoiceType,
        InvoiceTitle: invoiceTitle,
        InvoiceContent: invoiceContent,
        TaxId: taxId,
        Company: company,
        Phone: phone,
        Bank: bank,
        Account: account,
        Address: address,
        Email: email,
        VatId: vatId,
        Type: type_,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddInvoiceTemplateRequestWithoutParam() *AddInvoiceTemplateRequest {

    return &AddInvoiceTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoice_template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AddInvoiceTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: id(Optional) */
func (r *AddInvoiceTemplateRequest) SetId(id int64) {
    r.Id = &id
}

/* param invoiceType: 发票类型(Optional) */
func (r *AddInvoiceTemplateRequest) SetInvoiceType(invoiceType int) {
    r.InvoiceType = &invoiceType
}

/* param invoiceTitle: 发票抬头(Optional) */
func (r *AddInvoiceTemplateRequest) SetInvoiceTitle(invoiceTitle string) {
    r.InvoiceTitle = &invoiceTitle
}

/* param invoiceContent: 发票内容(Optional) */
func (r *AddInvoiceTemplateRequest) SetInvoiceContent(invoiceContent string) {
    r.InvoiceContent = &invoiceContent
}

/* param taxId: 税号(Optional) */
func (r *AddInvoiceTemplateRequest) SetTaxId(taxId string) {
    r.TaxId = &taxId
}

/* param company: 公司(Optional) */
func (r *AddInvoiceTemplateRequest) SetCompany(company string) {
    r.Company = &company
}

/* param phone: 手机号(Optional) */
func (r *AddInvoiceTemplateRequest) SetPhone(phone string) {
    r.Phone = &phone
}

/* param bank: 开户行(Optional) */
func (r *AddInvoiceTemplateRequest) SetBank(bank string) {
    r.Bank = &bank
}

/* param account: 开户账号(Optional) */
func (r *AddInvoiceTemplateRequest) SetAccount(account string) {
    r.Account = &account
}

/* param address: 开户行地址(Optional) */
func (r *AddInvoiceTemplateRequest) SetAddress(address string) {
    r.Address = &address
}

/* param email: 邮箱(Optional) */
func (r *AddInvoiceTemplateRequest) SetEmail(email string) {
    r.Email = &email
}

/* param vatId: vatId(Optional) */
func (r *AddInvoiceTemplateRequest) SetVatId(vatId int) {
    r.VatId = &vatId
}

/* param type_: type(Optional) */
func (r *AddInvoiceTemplateRequest) SetType(type_ int) {
    r.Type = &type_
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddInvoiceTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type AddInvoiceTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddInvoiceTemplateResult `json:"result"`
}

type AddInvoiceTemplateResult struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
}