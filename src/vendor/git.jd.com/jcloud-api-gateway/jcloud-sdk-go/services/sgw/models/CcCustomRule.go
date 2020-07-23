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


type CcCustomRule struct {

    /* 规则id (Optional) */
    Id *string `json:"id"`

    /* cc自定义规则的规则名称  */
    Name string `json:"name"`

    /* cc自定义规则的请求路径（路径不能重复）  */
    Url string `json:"url"`

    /* cc自定义规则的匹配条件（1 前缀匹配，2 完全匹配）  */
    Method int `json:"method"`

    /* cc自定义规则的规则模式（1 访问频次，2 响应时长）  */
    Mode int `json:"mode"`

    /* cc自定义规则的频次（2s 或 5/10s）  */
    Threshold string `json:"threshold"`

    /* cc自定义规则的动作处理（1 算法挑战，2 拦截，3 告警）  */
    Action int `json:"action"`

    /* 规则权重 (Optional) */
    Weight *int `json:"weight"`

    /* 防护状态( 0- 关闭，1-开启) (Optional) */
    Status *int `json:"status"`

    /* 创建时间 (Optional) */
    CreateAt *string `json:"createAt"`
}