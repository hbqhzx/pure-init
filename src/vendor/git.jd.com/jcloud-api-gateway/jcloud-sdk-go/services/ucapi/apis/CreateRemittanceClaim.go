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
    ucapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ucapi/models"
)

type CreateRemittanceClaimRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 汇款认领单信息 (Optional) */
    RemittanceClaim *ucapi.RemittanceClaim `json:"remittanceClaim"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRemittanceClaimRequest(
    regionId string,
) *CreateRemittanceClaimRequest {

	return &CreateRemittanceClaimRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/remittance/claims",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param remittanceClaim: 汇款认领单信息 (Optional)
 */
func NewCreateRemittanceClaimRequestWithAllParams(
    regionId string,
    remittanceClaim *ucapi.RemittanceClaim,
) *CreateRemittanceClaimRequest {

    return &CreateRemittanceClaimRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/remittance/claims",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RemittanceClaim: remittanceClaim,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRemittanceClaimRequestWithoutParam() *CreateRemittanceClaimRequest {

    return &CreateRemittanceClaimRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/remittance/claims",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateRemittanceClaimRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param remittanceClaim: 汇款认领单信息(Optional) */
func (r *CreateRemittanceClaimRequest) SetRemittanceClaim(remittanceClaim *ucapi.RemittanceClaim) {
    r.RemittanceClaim = remittanceClaim
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRemittanceClaimRequest) GetRegionId() string {
    return r.RegionId
}

type CreateRemittanceClaimResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRemittanceClaimResult `json:"result"`
}

type CreateRemittanceClaimResult struct {
}