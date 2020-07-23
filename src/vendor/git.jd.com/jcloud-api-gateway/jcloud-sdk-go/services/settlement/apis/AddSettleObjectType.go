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

type AddSettleObjectTypeRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 结算对象类型编码 (Optional) */
    TypeCode *string `json:"typeCode"`

    /* 结算对象类型名称 (Optional) */
    TypeName *string `json:"typeName"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddSettleObjectTypeRequest(
    regionId string,
) *AddSettleObjectTypeRequest {

	return &AddSettleObjectTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/systemConfig:settleObjectTypeAdd",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param typeCode: 结算对象类型编码 (Optional)
 * param typeName: 结算对象类型名称 (Optional)
 */
func NewAddSettleObjectTypeRequestWithAllParams(
    regionId string,
    typeCode *string,
    typeName *string,
) *AddSettleObjectTypeRequest {

    return &AddSettleObjectTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/systemConfig:settleObjectTypeAdd",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TypeCode: typeCode,
        TypeName: typeName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddSettleObjectTypeRequestWithoutParam() *AddSettleObjectTypeRequest {

    return &AddSettleObjectTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/systemConfig:settleObjectTypeAdd",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AddSettleObjectTypeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param typeCode: 结算对象类型编码(Optional) */
func (r *AddSettleObjectTypeRequest) SetTypeCode(typeCode string) {
    r.TypeCode = &typeCode
}

/* param typeName: 结算对象类型名称(Optional) */
func (r *AddSettleObjectTypeRequest) SetTypeName(typeName string) {
    r.TypeName = &typeName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddSettleObjectTypeRequest) GetRegionId() string {
    return r.RegionId
}

type AddSettleObjectTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddSettleObjectTypeResult `json:"result"`
}

type AddSettleObjectTypeResult struct {
    Result bool `json:"result"`
}