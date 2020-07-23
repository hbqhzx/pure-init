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

type AddRemmitClaimRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    Id *string `json:"id"`

    /*  (Optional) */
    Pin *string `json:"pin"`

    /*  (Optional) */
    RemittorAccount *string `json:"remittorAccount"`

    /*  (Optional) */
    RemittorBankAccount *string `json:"remittorBankAccount"`

    /*  (Optional) */
    RemitBankName *string `json:"remitBankName"`

    /*  (Optional) */
    RemitTime *string `json:"remitTime"`

    /*  (Optional) */
    ContactsPhone *string `json:"contactsPhone"`

    /*  (Optional) */
    RemitPicture *string `json:"remitPicture"`

    /*  (Optional) */
    RemitAmount *string `json:"remitAmount"`

    /*  (Optional) */
    ClaId *string `json:"claId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddRemmitClaimRequest(
    regionId string,
) *AddRemmitClaimRequest {

	return &AddRemmitClaimRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pay:addRemmitClaim",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param id:  (Optional)
 * param pin:  (Optional)
 * param remittorAccount:  (Optional)
 * param remittorBankAccount:  (Optional)
 * param remitBankName:  (Optional)
 * param remitTime:  (Optional)
 * param contactsPhone:  (Optional)
 * param remitPicture:  (Optional)
 * param remitAmount:  (Optional)
 * param claId:  (Optional)
 */
func NewAddRemmitClaimRequestWithAllParams(
    regionId string,
    id *string,
    pin *string,
    remittorAccount *string,
    remittorBankAccount *string,
    remitBankName *string,
    remitTime *string,
    contactsPhone *string,
    remitPicture *string,
    remitAmount *string,
    claId *string,
) *AddRemmitClaimRequest {

    return &AddRemmitClaimRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pay:addRemmitClaim",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        Pin: pin,
        RemittorAccount: remittorAccount,
        RemittorBankAccount: remittorBankAccount,
        RemitBankName: remitBankName,
        RemitTime: remitTime,
        ContactsPhone: contactsPhone,
        RemitPicture: remitPicture,
        RemitAmount: remitAmount,
        ClaId: claId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddRemmitClaimRequestWithoutParam() *AddRemmitClaimRequest {

    return &AddRemmitClaimRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pay:addRemmitClaim",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddRemmitClaimRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: (Optional) */
func (r *AddRemmitClaimRequest) SetId(id string) {
    r.Id = &id
}

/* param pin: (Optional) */
func (r *AddRemmitClaimRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param remittorAccount: (Optional) */
func (r *AddRemmitClaimRequest) SetRemittorAccount(remittorAccount string) {
    r.RemittorAccount = &remittorAccount
}

/* param remittorBankAccount: (Optional) */
func (r *AddRemmitClaimRequest) SetRemittorBankAccount(remittorBankAccount string) {
    r.RemittorBankAccount = &remittorBankAccount
}

/* param remitBankName: (Optional) */
func (r *AddRemmitClaimRequest) SetRemitBankName(remitBankName string) {
    r.RemitBankName = &remitBankName
}

/* param remitTime: (Optional) */
func (r *AddRemmitClaimRequest) SetRemitTime(remitTime string) {
    r.RemitTime = &remitTime
}

/* param contactsPhone: (Optional) */
func (r *AddRemmitClaimRequest) SetContactsPhone(contactsPhone string) {
    r.ContactsPhone = &contactsPhone
}

/* param remitPicture: (Optional) */
func (r *AddRemmitClaimRequest) SetRemitPicture(remitPicture string) {
    r.RemitPicture = &remitPicture
}

/* param remitAmount: (Optional) */
func (r *AddRemmitClaimRequest) SetRemitAmount(remitAmount string) {
    r.RemitAmount = &remitAmount
}

/* param claId: (Optional) */
func (r *AddRemmitClaimRequest) SetClaId(claId string) {
    r.ClaId = &claId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddRemmitClaimRequest) GetRegionId() string {
    return r.RegionId
}

type AddRemmitClaimResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddRemmitClaimResult `json:"result"`
}

type AddRemmitClaimResult struct {
    Data bool `json:"data"`
}