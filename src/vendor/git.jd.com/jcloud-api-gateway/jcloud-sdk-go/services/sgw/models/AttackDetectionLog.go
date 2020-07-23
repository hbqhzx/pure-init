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


type AttackDetectionLog struct {

    /* 日志id (Optional) */
    Id *string `json:"id"`

    /* 产生时间 (Optional) */
    Time *string `json:"time"`

    /* 日志来源，vpc-waf:应用安全网关，cloud-waf:云WAF，sa:态势感知 (Optional) */
    Source *string `json:"source"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 实例ID (Optional) */
    WafId *string `json:"wafId"`

    /* 地域 (Optional) */
    Region *string `json:"region"`

    /* 误报类型，0：AI检出，1：用户提交 (Optional) */
    Type *int `json:"type"`

    /* 模型类型，0：默认 (Optional) */
    ModelType *int `json:"modelType"`

    /* 规则id (Optional) */
    RuleId *string `json:"ruleId"`

    /* URL (Optional) */
    Url *string `json:"url"`

    /* 匹配域 (Optional) */
    HitPosition *string `json:"hitPosition"`

    /* 恶意负载 (Optional) */
    Payload *string `json:"payload"`

    /* 处理状态，0：未处理，1：标记，2：忽略  */
    ProcessedStatus int `json:"processedStatus"`

    /* 备注 (Optional) */
    Comment *string `json:"comment"`
}