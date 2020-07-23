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
    clouddnsservice "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/clouddnsservice/models"
)

type AddDomainRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 主域名的套餐类型, 0->免费 ,1->企业版, 2->高级版  */
    PackId int `json:"packId"`

    /* 要添加的主域名  */
    DomainName string `json:"domainName"`

    /* 主域名的ID，升级套餐必填，请使用getDomains获取 (Optional) */
    DomainId *int `json:"domainId"`

    /* 1->新购买、3->升级，收费套餐的域名必填 (Optional) */
    BuyType *int `json:"buyType"`

    /* 取值1，2，3 ，含义：时长，收费套餐的域名必填 (Optional) */
    TimeSpan *int `json:"timeSpan"`

    /* 时间单位，收费套餐的域名必填，1：小时，2：天，3：月，4：年 (Optional) */
    TimeUnit *int `json:"timeUnit"`

    /* 计费类型，可以不传此参数。 (Optional) */
    BillingType *int `json:"billingType"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param packId: 主域名的套餐类型, 0->免费 ,1->企业版, 2->高级版 (Required)
 * param domainName: 要添加的主域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddDomainRequest(
    regionId string,
    packId int,
    domainName string,
) *AddDomainRequest {

	return &AddDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domainAdd",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PackId: packId,
        DomainName: domainName,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param packId: 主域名的套餐类型, 0->免费 ,1->企业版, 2->高级版 (Required)
 * param domainName: 要添加的主域名 (Required)
 * param domainId: 主域名的ID，升级套餐必填，请使用getDomains获取 (Optional)
 * param buyType: 1->新购买、3->升级，收费套餐的域名必填 (Optional)
 * param timeSpan: 取值1，2，3 ，含义：时长，收费套餐的域名必填 (Optional)
 * param timeUnit: 时间单位，收费套餐的域名必填，1：小时，2：天，3：月，4：年 (Optional)
 * param billingType: 计费类型，可以不传此参数。 (Optional)
 */
func NewAddDomainRequestWithAllParams(
    regionId string,
    packId int,
    domainName string,
    domainId *int,
    buyType *int,
    timeSpan *int,
    timeUnit *int,
    billingType *int,
) *AddDomainRequest {

    return &AddDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domainAdd",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PackId: packId,
        DomainName: domainName,
        DomainId: domainId,
        BuyType: buyType,
        TimeSpan: timeSpan,
        TimeUnit: timeUnit,
        BillingType: billingType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddDomainRequestWithoutParam() *AddDomainRequest {

    return &AddDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domainAdd",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *AddDomainRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param packId: 主域名的套餐类型, 0->免费 ,1->企业版, 2->高级版(Required) */
func (r *AddDomainRequest) SetPackId(packId int) {
    r.PackId = packId
}

/* param domainName: 要添加的主域名(Required) */
func (r *AddDomainRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param domainId: 主域名的ID，升级套餐必填，请使用getDomains获取(Optional) */
func (r *AddDomainRequest) SetDomainId(domainId int) {
    r.DomainId = &domainId
}

/* param buyType: 1->新购买、3->升级，收费套餐的域名必填(Optional) */
func (r *AddDomainRequest) SetBuyType(buyType int) {
    r.BuyType = &buyType
}

/* param timeSpan: 取值1，2，3 ，含义：时长，收费套餐的域名必填(Optional) */
func (r *AddDomainRequest) SetTimeSpan(timeSpan int) {
    r.TimeSpan = &timeSpan
}

/* param timeUnit: 时间单位，收费套餐的域名必填，1：小时，2：天，3：月，4：年(Optional) */
func (r *AddDomainRequest) SetTimeUnit(timeUnit int) {
    r.TimeUnit = &timeUnit
}

/* param billingType: 计费类型，可以不传此参数。(Optional) */
func (r *AddDomainRequest) SetBillingType(billingType int) {
    r.BillingType = &billingType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddDomainRequest) GetRegionId() string {
    return r.RegionId
}

type AddDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddDomainResult `json:"result"`
}

type AddDomainResult struct {
    Data clouddnsservice.DomainAdded `json:"data"`
    Order string `json:"order"`
}