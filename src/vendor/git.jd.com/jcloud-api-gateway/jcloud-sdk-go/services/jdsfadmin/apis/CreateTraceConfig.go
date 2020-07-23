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

type CreateTraceConfigRequest struct {

    core.JDCloudRequest

    /* 可用区key 例如 cn-north-1  */
    RegionId string `json:"regionId"`

    /* 配置前台显示名称  */
    TraceConfigName string `json:"traceConfigName"`

    /* 实例规格描述  */
    TpsSpec string `json:"tpsSpec"`

    /* 权重值  */
    ShowOrder int `json:"showOrder"`

    /* manager规格  */
    FlavorManager string `json:"flavorManager"`

    /* 节点规格  */
    Flavor string `json:"flavor"`

    /* 可用区一  */
    Az1 string `json:"az1"`

    /* 可用区二  */
    Az2 string `json:"az2"`

    /* 可用区三  */
    Az3 string `json:"az3"`

    /* 可用区一数量（最小值为0）  */
    Az1Num string `json:"az1Num"`

    /* 可用区二数量（最小值为0）  */
    Az2Num string `json:"az2Num"`

    /* 可用区三数量（最小值为0）  */
    Az3Num string `json:"az3Num"`

    /* es 节点规格  */
    TraceEsNodeClass string `json:"traceEsNodeClass"`

    /* es 节点数量  */
    TraceEsNodeCount int `json:"traceEsNodeCount"`

    /* es 单点存储规格  */
    TraceEsNodeDeskSize int `json:"traceEsNodeDeskSize"`
}

/*
 * param regionId: 可用区key 例如 cn-north-1 (Required)
 * param traceConfigName: 配置前台显示名称 (Required)
 * param tpsSpec: 实例规格描述 (Required)
 * param showOrder: 权重值 (Required)
 * param flavorManager: manager规格 (Required)
 * param flavor: 节点规格 (Required)
 * param az1: 可用区一 (Required)
 * param az2: 可用区二 (Required)
 * param az3: 可用区三 (Required)
 * param az1Num: 可用区一数量（最小值为0） (Required)
 * param az2Num: 可用区二数量（最小值为0） (Required)
 * param az3Num: 可用区三数量（最小值为0） (Required)
 * param traceEsNodeClass: es 节点规格 (Required)
 * param traceEsNodeCount: es 节点数量 (Required)
 * param traceEsNodeDeskSize: es 单点存储规格 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTraceConfigRequest(
    regionId string,
    traceConfigName string,
    tpsSpec string,
    showOrder int,
    flavorManager string,
    flavor string,
    az1 string,
    az2 string,
    az3 string,
    az1Num string,
    az2Num string,
    az3Num string,
    traceEsNodeClass string,
    traceEsNodeCount int,
    traceEsNodeDeskSize int,
) *CreateTraceConfigRequest {

	return &CreateTraceConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/traceconfigs",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TraceConfigName: traceConfigName,
        TpsSpec: tpsSpec,
        ShowOrder: showOrder,
        FlavorManager: flavorManager,
        Flavor: flavor,
        Az1: az1,
        Az2: az2,
        Az3: az3,
        Az1Num: az1Num,
        Az2Num: az2Num,
        Az3Num: az3Num,
        TraceEsNodeClass: traceEsNodeClass,
        TraceEsNodeCount: traceEsNodeCount,
        TraceEsNodeDeskSize: traceEsNodeDeskSize,
	}
}

/*
 * param regionId: 可用区key 例如 cn-north-1 (Required)
 * param traceConfigName: 配置前台显示名称 (Required)
 * param tpsSpec: 实例规格描述 (Required)
 * param showOrder: 权重值 (Required)
 * param flavorManager: manager规格 (Required)
 * param flavor: 节点规格 (Required)
 * param az1: 可用区一 (Required)
 * param az2: 可用区二 (Required)
 * param az3: 可用区三 (Required)
 * param az1Num: 可用区一数量（最小值为0） (Required)
 * param az2Num: 可用区二数量（最小值为0） (Required)
 * param az3Num: 可用区三数量（最小值为0） (Required)
 * param traceEsNodeClass: es 节点规格 (Required)
 * param traceEsNodeCount: es 节点数量 (Required)
 * param traceEsNodeDeskSize: es 单点存储规格 (Required)
 */
func NewCreateTraceConfigRequestWithAllParams(
    regionId string,
    traceConfigName string,
    tpsSpec string,
    showOrder int,
    flavorManager string,
    flavor string,
    az1 string,
    az2 string,
    az3 string,
    az1Num string,
    az2Num string,
    az3Num string,
    traceEsNodeClass string,
    traceEsNodeCount int,
    traceEsNodeDeskSize int,
) *CreateTraceConfigRequest {

    return &CreateTraceConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/traceconfigs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TraceConfigName: traceConfigName,
        TpsSpec: tpsSpec,
        ShowOrder: showOrder,
        FlavorManager: flavorManager,
        Flavor: flavor,
        Az1: az1,
        Az2: az2,
        Az3: az3,
        Az1Num: az1Num,
        Az2Num: az2Num,
        Az3Num: az3Num,
        TraceEsNodeClass: traceEsNodeClass,
        TraceEsNodeCount: traceEsNodeCount,
        TraceEsNodeDeskSize: traceEsNodeDeskSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTraceConfigRequestWithoutParam() *CreateTraceConfigRequest {

    return &CreateTraceConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/traceconfigs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区key 例如 cn-north-1(Required) */
func (r *CreateTraceConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param traceConfigName: 配置前台显示名称(Required) */
func (r *CreateTraceConfigRequest) SetTraceConfigName(traceConfigName string) {
    r.TraceConfigName = traceConfigName
}

/* param tpsSpec: 实例规格描述(Required) */
func (r *CreateTraceConfigRequest) SetTpsSpec(tpsSpec string) {
    r.TpsSpec = tpsSpec
}

/* param showOrder: 权重值(Required) */
func (r *CreateTraceConfigRequest) SetShowOrder(showOrder int) {
    r.ShowOrder = showOrder
}

/* param flavorManager: manager规格(Required) */
func (r *CreateTraceConfigRequest) SetFlavorManager(flavorManager string) {
    r.FlavorManager = flavorManager
}

/* param flavor: 节点规格(Required) */
func (r *CreateTraceConfigRequest) SetFlavor(flavor string) {
    r.Flavor = flavor
}

/* param az1: 可用区一(Required) */
func (r *CreateTraceConfigRequest) SetAz1(az1 string) {
    r.Az1 = az1
}

/* param az2: 可用区二(Required) */
func (r *CreateTraceConfigRequest) SetAz2(az2 string) {
    r.Az2 = az2
}

/* param az3: 可用区三(Required) */
func (r *CreateTraceConfigRequest) SetAz3(az3 string) {
    r.Az3 = az3
}

/* param az1Num: 可用区一数量（最小值为0）(Required) */
func (r *CreateTraceConfigRequest) SetAz1Num(az1Num string) {
    r.Az1Num = az1Num
}

/* param az2Num: 可用区二数量（最小值为0）(Required) */
func (r *CreateTraceConfigRequest) SetAz2Num(az2Num string) {
    r.Az2Num = az2Num
}

/* param az3Num: 可用区三数量（最小值为0）(Required) */
func (r *CreateTraceConfigRequest) SetAz3Num(az3Num string) {
    r.Az3Num = az3Num
}

/* param traceEsNodeClass: es 节点规格(Required) */
func (r *CreateTraceConfigRequest) SetTraceEsNodeClass(traceEsNodeClass string) {
    r.TraceEsNodeClass = traceEsNodeClass
}

/* param traceEsNodeCount: es 节点数量(Required) */
func (r *CreateTraceConfigRequest) SetTraceEsNodeCount(traceEsNodeCount int) {
    r.TraceEsNodeCount = traceEsNodeCount
}

/* param traceEsNodeDeskSize: es 单点存储规格(Required) */
func (r *CreateTraceConfigRequest) SetTraceEsNodeDeskSize(traceEsNodeDeskSize int) {
    r.TraceEsNodeDeskSize = traceEsNodeDeskSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTraceConfigRequest) GetRegionId() string {
    return r.RegionId
}

type CreateTraceConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTraceConfigResult `json:"result"`
}

type CreateTraceConfigResult struct {
    CreateTraceConfig string `json:"createTraceConfig"`
}