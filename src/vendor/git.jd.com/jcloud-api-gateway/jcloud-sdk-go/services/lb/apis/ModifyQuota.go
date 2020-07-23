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

type ModifyQuotaRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* lb类型，取值范围：alb、nlb，默认为alb (Optional) */
    LbType *string `json:"lbType"`

    /* 资源类型，取值范围：loadbalancer、listener、target_group、target、backend、urlMap(仅alb支持)、rules(仅alb支持)  */
    Type string `json:"type"`

    /* type为loadbalancer不设置, type为listener、backend、target_group、urlMap设置为loadbalancerId, type为target设置为targetGroupId, type为rules设置为urlMapId (Optional) */
    ParentResourceId *string `json:"parentResourceId"`

    /* 配额大小  */
    MaxLimit int `json:"maxLimit"`
}

/*
 * param regionId: Region ID (Required)
 * param type_: 资源类型，取值范围：loadbalancer、listener、target_group、target、backend、urlMap(仅alb支持)、rules(仅alb支持) (Required)
 * param maxLimit: 配额大小 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyQuotaRequest(
    regionId string,
    type_ string,
    maxLimit int,
) *ModifyQuotaRequest {

	return &ModifyQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/quotas/",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Type: type_,
        MaxLimit: maxLimit,
	}
}

/*
 * param regionId: Region ID (Required)
 * param lbType: lb类型，取值范围：alb、nlb，默认为alb (Optional)
 * param type_: 资源类型，取值范围：loadbalancer、listener、target_group、target、backend、urlMap(仅alb支持)、rules(仅alb支持) (Required)
 * param parentResourceId: type为loadbalancer不设置, type为listener、backend、target_group、urlMap设置为loadbalancerId, type为target设置为targetGroupId, type为rules设置为urlMapId (Optional)
 * param maxLimit: 配额大小 (Required)
 */
func NewModifyQuotaRequestWithAllParams(
    regionId string,
    lbType *string,
    type_ string,
    parentResourceId *string,
    maxLimit int,
) *ModifyQuotaRequest {

    return &ModifyQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas/",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LbType: lbType,
        Type: type_,
        ParentResourceId: parentResourceId,
        MaxLimit: maxLimit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyQuotaRequestWithoutParam() *ModifyQuotaRequest {

    return &ModifyQuotaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas/",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyQuotaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param lbType: lb类型，取值范围：alb、nlb，默认为alb(Optional) */
func (r *ModifyQuotaRequest) SetLbType(lbType string) {
    r.LbType = &lbType
}

/* param type_: 资源类型，取值范围：loadbalancer、listener、target_group、target、backend、urlMap(仅alb支持)、rules(仅alb支持)(Required) */
func (r *ModifyQuotaRequest) SetType(type_ string) {
    r.Type = type_
}

/* param parentResourceId: type为loadbalancer不设置, type为listener、backend、target_group、urlMap设置为loadbalancerId, type为target设置为targetGroupId, type为rules设置为urlMapId(Optional) */
func (r *ModifyQuotaRequest) SetParentResourceId(parentResourceId string) {
    r.ParentResourceId = &parentResourceId
}

/* param maxLimit: 配额大小(Required) */
func (r *ModifyQuotaRequest) SetMaxLimit(maxLimit int) {
    r.MaxLimit = maxLimit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyQuotaRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyQuotaResult `json:"result"`
}

type ModifyQuotaResult struct {
}