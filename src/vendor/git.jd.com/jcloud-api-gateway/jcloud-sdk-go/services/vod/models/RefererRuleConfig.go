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


type RefererRuleConfig struct {

    /* Referer规则适用策略：允许/拒绝 (Optional) */
    Strategy *string `json:"strategy"`

    /*  (Optional) */
    Domains []string `json:"domains"`

    /* 是否允许Referer为空，如通过浏览器直接访问 (Optional) */
    AllowBlank *bool `json:"allowBlank"`
}