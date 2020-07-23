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


type WafCC struct {

    /* waf实例id (Optional) */
    WafId string `json:"wafId"`

    /* 攻击源ip (Optional) */
    CcIp string `json:"ccIp"`

    /* 攻击源ip情报标签 (Optional) */
    CcIpTITag string `json:"ccIpTITag"`

    /* 攻击源ip所属地区 (Optional) */
    CcIpRegion string `json:"ccIpRegion"`

    /* 攻击时间 (Optional) */
    CcDatetime string `json:"ccDatetime"`

    /* 攻击次数 (Optional) */
    CcCount int `json:"ccCount"`

    /* 被攻击URL,攻击详情 (Optional) */
    CcURL string `json:"ccURL"`

    /* 行为(block阻断，log告警，challenge_js算法挑战) (Optional) */
    CcAction string `json:"ccAction"`
}
