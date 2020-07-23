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


type IpResourceProtectInfo struct {

    /* 攻击开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 攻击结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 安全状态, 0: 安全, 1: 清洗, 2: 黑洞 (Optional) */
    Status int `json:"status"`

    /* 触发原因，0->未知 1->四层 2->七层 3->四和7层 (Optional) */
    Cause int `json:"cause"`
}