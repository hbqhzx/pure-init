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

package models


type TraceConfigDetail struct {

    /* 主键 (Optional) */
    TraceConfigId int64 `json:"traceConfigId"`

    /* 配置名称 (Optional) */
    TraceConfigName string `json:"traceConfigName"`

    /* 实例规格描述 (Optional) */
    TpsSpec int `json:"tpsSpec"`

    /* 权重值 (Optional) */
    ShowOrder int `json:"showOrder"`

    /* 节点规格 (Optional) */
    Flavor string `json:"flavor"`

    /* manager规格 (Optional) */
    FlavorManager string `json:"flavorManager"`

    /* 可用区一 (Optional) */
    Az1 string `json:"az1"`

    /* 可用区二 (Optional) */
    Az2 string `json:"az2"`

    /* 可用区三 (Optional) */
    Az3 string `json:"az3"`

    /* 可用区一数量 (Optional) */
    Az1Num string `json:"az1Num"`

    /* 可用区二数量 (Optional) */
    Az2Num string `json:"az2Num"`

    /* 可用区三数量 (Optional) */
    Az3Num string `json:"az3Num"`

    /* es 节点规格 (Optional) */
    TraceEsNodeClass string `json:"traceEsNodeClass"`

    /* es 节点数量 (Optional) */
    TraceEsNodeCount int `json:"traceEsNodeCount"`

    /* es 节点硬盘大小 (Optional) */
    TraceEsNodeDeskSize int `json:"traceEsNodeDeskSize"`
}
