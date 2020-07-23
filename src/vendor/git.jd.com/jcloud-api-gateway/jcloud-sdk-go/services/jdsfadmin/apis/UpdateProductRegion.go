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

type UpdateProductRegionRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 产品区域配置的数据库 id  */
    ProductRegionId int `json:"productRegionId"`

    /* 区域标识  */
    RegionKey string `json:"regionKey"`

    /* 京东云region 名称  */
    RegionName string `json:"regionName"`

    /* 京东云region 可用区状态  */
    RegionStatus string `json:"regionStatus"`

    /* 微服务产品线产品关键词  */
    ProductLineKey string `json:"productLineKey"`
}

/*
 * param regionId: 区域id (Required)
 * param productRegionId: 产品区域配置的数据库 id (Required)
 * param regionKey: 区域标识 (Required)
 * param regionName: 京东云region 名称 (Required)
 * param regionStatus: 京东云region 可用区状态 (Required)
 * param productLineKey: 微服务产品线产品关键词 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateProductRegionRequest(
    regionId string,
    productRegionId int,
    regionKey string,
    regionName string,
    regionStatus string,
    productLineKey string,
) *UpdateProductRegionRequest {

	return &UpdateProductRegionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/productregions/{productRegionId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ProductRegionId: productRegionId,
        RegionKey: regionKey,
        RegionName: regionName,
        RegionStatus: regionStatus,
        ProductLineKey: productLineKey,
	}
}

/*
 * param regionId: 区域id (Required)
 * param productRegionId: 产品区域配置的数据库 id (Required)
 * param regionKey: 区域标识 (Required)
 * param regionName: 京东云region 名称 (Required)
 * param regionStatus: 京东云region 可用区状态 (Required)
 * param productLineKey: 微服务产品线产品关键词 (Required)
 */
func NewUpdateProductRegionRequestWithAllParams(
    regionId string,
    productRegionId int,
    regionKey string,
    regionName string,
    regionStatus string,
    productLineKey string,
) *UpdateProductRegionRequest {

    return &UpdateProductRegionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/productregions/{productRegionId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ProductRegionId: productRegionId,
        RegionKey: regionKey,
        RegionName: regionName,
        RegionStatus: regionStatus,
        ProductLineKey: productLineKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateProductRegionRequestWithoutParam() *UpdateProductRegionRequest {

    return &UpdateProductRegionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/productregions/{productRegionId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *UpdateProductRegionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param productRegionId: 产品区域配置的数据库 id(Required) */
func (r *UpdateProductRegionRequest) SetProductRegionId(productRegionId int) {
    r.ProductRegionId = productRegionId
}

/* param regionKey: 区域标识(Required) */
func (r *UpdateProductRegionRequest) SetRegionKey(regionKey string) {
    r.RegionKey = regionKey
}

/* param regionName: 京东云region 名称(Required) */
func (r *UpdateProductRegionRequest) SetRegionName(regionName string) {
    r.RegionName = regionName
}

/* param regionStatus: 京东云region 可用区状态(Required) */
func (r *UpdateProductRegionRequest) SetRegionStatus(regionStatus string) {
    r.RegionStatus = regionStatus
}

/* param productLineKey: 微服务产品线产品关键词(Required) */
func (r *UpdateProductRegionRequest) SetProductLineKey(productLineKey string) {
    r.ProductLineKey = productLineKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateProductRegionRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateProductRegionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateProductRegionResult `json:"result"`
}

type UpdateProductRegionResult struct {
    UpdateProductRegion string `json:"updateProductRegion"`
}