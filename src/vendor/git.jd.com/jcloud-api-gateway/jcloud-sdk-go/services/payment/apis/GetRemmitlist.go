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
    payment "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/payment/models"
)

type GetRemmitlistRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    AmountAccount *string `json:"amountAccount"`

    /*  (Optional) */
    AmountDate *string `json:"amountDate"`

    /*  (Optional) */
    AmountState *string `json:"amountState"`

    /*  (Optional) */
    AmountUser *string `json:"amountUser"`

    /*  (Optional) */
    BankISN *string `json:"bankISN"`

    /* null (Optional) */
    ClaId *int `json:"claId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetRemmitlistRequest(
    regionId string,
) *GetRemmitlistRequest {

	return &GetRemmitlistRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pay:getRemmitlist",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param amountAccount:  (Optional)
 * param amountDate:  (Optional)
 * param amountState:  (Optional)
 * param amountUser:  (Optional)
 * param bankISN:  (Optional)
 * param claId: null (Optional)
 */
func NewGetRemmitlistRequestWithAllParams(
    regionId string,
    amountAccount *string,
    amountDate *string,
    amountState *string,
    amountUser *string,
    bankISN *string,
    claId *int,
) *GetRemmitlistRequest {

    return &GetRemmitlistRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pay:getRemmitlist",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AmountAccount: amountAccount,
        AmountDate: amountDate,
        AmountState: amountState,
        AmountUser: amountUser,
        BankISN: bankISN,
        ClaId: claId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetRemmitlistRequestWithoutParam() *GetRemmitlistRequest {

    return &GetRemmitlistRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pay:getRemmitlist",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetRemmitlistRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param amountAccount: (Optional) */
func (r *GetRemmitlistRequest) SetAmountAccount(amountAccount string) {
    r.AmountAccount = &amountAccount
}

/* param amountDate: (Optional) */
func (r *GetRemmitlistRequest) SetAmountDate(amountDate string) {
    r.AmountDate = &amountDate
}

/* param amountState: (Optional) */
func (r *GetRemmitlistRequest) SetAmountState(amountState string) {
    r.AmountState = &amountState
}

/* param amountUser: (Optional) */
func (r *GetRemmitlistRequest) SetAmountUser(amountUser string) {
    r.AmountUser = &amountUser
}

/* param bankISN: (Optional) */
func (r *GetRemmitlistRequest) SetBankISN(bankISN string) {
    r.BankISN = &bankISN
}

/* param claId: null(Optional) */
func (r *GetRemmitlistRequest) SetClaId(claId int) {
    r.ClaId = &claId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetRemmitlistRequest) GetRegionId() string {
    return r.RegionId
}

type GetRemmitlistResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetRemmitlistResult `json:"result"`
}

type GetRemmitlistResult struct {
    Data payment.QueryCollectResponseResult `json:"data"`
}