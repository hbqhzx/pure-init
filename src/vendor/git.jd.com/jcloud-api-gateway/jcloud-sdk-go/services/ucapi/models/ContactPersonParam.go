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


type ContactPersonParam struct {

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 联系人id (Optional) */
    Id *int64 `json:"id"`

    /* 联系人姓名 (Optional) */
    UserName *string `json:"userName"`

    /* 联系人手机号码 (Optional) */
    Mobile *string `json:"mobile"`

    /* 联系人邮箱地址 (Optional) */
    Email *string `json:"email"`

    /* 验证码 (Optional) */
    SmsCode *string `json:"smsCode"`

    /* 是否为账号联系人 1-是 2-不是 (Optional) */
    IsSelf *int `json:"isSelf"`
}
