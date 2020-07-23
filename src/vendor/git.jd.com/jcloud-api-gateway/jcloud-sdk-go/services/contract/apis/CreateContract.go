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
    contract "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/contract/models"
)

type CreateContractRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    QueryVo *contract.ContractReqVo `json:"queryVo"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateContractRequest(
    regionId string,
) *CreateContractRequest {

	return &CreateContractRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/contract:createContract",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param queryVo:  (Optional)
 */
func NewCreateContractRequestWithAllParams(
    regionId string,
    queryVo *contract.ContractReqVo,
) *CreateContractRequest {

    return &CreateContractRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contract:createContract",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        QueryVo: queryVo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateContractRequestWithoutParam() *CreateContractRequest {

    return &CreateContractRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contract:createContract",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateContractRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param queryVo: (Optional) */
func (r *CreateContractRequest) SetQueryVo(queryVo *contract.ContractReqVo) {
    r.QueryVo = queryVo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateContractRequest) GetRegionId() string {
    return r.RegionId
}

type CreateContractResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateContractResult `json:"result"`
}

type CreateContractResult struct {
    Result bool `json:"result"`
    CreateContractRes contract.CreateContractRes `json:"createContractRes"`
}