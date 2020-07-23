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

type AddPersonRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 生产线key（REGISTRY,OPEN_TRACE,APP_CONFIG）  */
    ProductLineKey string `json:"productLineKey"`

    /* 用户Pin,支持多用户添加  */
    UserPin string `json:"userPin"`

    /* 限制数量  */
    LimitCount int `json:"limitCount"`

    /* 产品单位(枚举值 ALL_REGION/SINGLE_REGION)  */
    ProductUnit string `json:"productUnit"`
}

/*
 * param regionId: 区域id (Required)
 * param productLineKey: 生产线key（REGISTRY,OPEN_TRACE,APP_CONFIG） (Required)
 * param userPin: 用户Pin,支持多用户添加 (Required)
 * param limitCount: 限制数量 (Required)
 * param productUnit: 产品单位(枚举值 ALL_REGION/SINGLE_REGION) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddPersonRequest(
    regionId string,
    productLineKey string,
    userPin string,
    limitCount int,
    productUnit string,
) *AddPersonRequest {

	return &AddPersonRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/person_resource_limits",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ProductLineKey: productLineKey,
        UserPin: userPin,
        LimitCount: limitCount,
        ProductUnit: productUnit,
	}
}

/*
 * param regionId: 区域id (Required)
 * param productLineKey: 生产线key（REGISTRY,OPEN_TRACE,APP_CONFIG） (Required)
 * param userPin: 用户Pin,支持多用户添加 (Required)
 * param limitCount: 限制数量 (Required)
 * param productUnit: 产品单位(枚举值 ALL_REGION/SINGLE_REGION) (Required)
 */
func NewAddPersonRequestWithAllParams(
    regionId string,
    productLineKey string,
    userPin string,
    limitCount int,
    productUnit string,
) *AddPersonRequest {

    return &AddPersonRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/person_resource_limits",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ProductLineKey: productLineKey,
        UserPin: userPin,
        LimitCount: limitCount,
        ProductUnit: productUnit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddPersonRequestWithoutParam() *AddPersonRequest {

    return &AddPersonRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/person_resource_limits",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *AddPersonRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param productLineKey: 生产线key（REGISTRY,OPEN_TRACE,APP_CONFIG）(Required) */
func (r *AddPersonRequest) SetProductLineKey(productLineKey string) {
    r.ProductLineKey = productLineKey
}

/* param userPin: 用户Pin,支持多用户添加(Required) */
func (r *AddPersonRequest) SetUserPin(userPin string) {
    r.UserPin = userPin
}

/* param limitCount: 限制数量(Required) */
func (r *AddPersonRequest) SetLimitCount(limitCount int) {
    r.LimitCount = limitCount
}

/* param productUnit: 产品单位(枚举值 ALL_REGION/SINGLE_REGION)(Required) */
func (r *AddPersonRequest) SetProductUnit(productUnit string) {
    r.ProductUnit = productUnit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddPersonRequest) GetRegionId() string {
    return r.RegionId
}

type AddPersonResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddPersonResult `json:"result"`
}

type AddPersonResult struct {
    AddPersonResult string `json:"addPersonResult"`
}