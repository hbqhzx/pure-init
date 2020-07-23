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


type ChangeStatusReqVo struct {

    /* 短信验证码,转正式合同需要短信验证 (Optional) */
    AuthCode *string `json:"authCode"`

    /* 合同号 (Optional) */
    ContractNumber *string `json:"contractNumber"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 需要转变的-->合同状态（1：正式，3：废弃） (Optional) */
    Status *int `json:"status"`
}
