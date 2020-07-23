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


type AlarmConfigSpec struct {

    /* 黑洞告警邮件开关 0 关闭 1 开启 (Optional) */
    BlackHoleAlarmEmailStatus *int `json:"blackHoleAlarmEmailStatus"`

    /* 黑洞告警短信开关 0 关闭 1 开启 (Optional) */
    BlackHoleAlarmSmsStatus *int `json:"blackHoleAlarmSmsStatus"`

    /* 黑洞告警总开关  0 关闭 1 开启 (Optional) */
    BlackHoleAlarmStatus *int `json:"blackHoleAlarmStatus"`

    /* DDos 攻击告警邮件开关  0 关闭 1 开启 (Optional) */
    DdosAlarmEmailStatus *int `json:"ddosAlarmEmailStatus"`

    /* DDos 攻击告警短信开关  0 关闭 1 开启 (Optional) */
    DdosAlarmSmsStatus *int `json:"ddosAlarmSmsStatus"`

    /* DDos 告警总开关 0 关闭 1 开启 (Optional) */
    DdosAlarmStatus *int `json:"ddosAlarmStatus"`

    /* 错误码告警总开关 (Optional) */
    ErrorCodeAlarmStatus *int `json:"errorCodeAlarmStatus"`

    /* 错误码告警域名列表 (Optional) */
    ErrorCodeDomain []string `json:"errorCodeDomain"`
}
