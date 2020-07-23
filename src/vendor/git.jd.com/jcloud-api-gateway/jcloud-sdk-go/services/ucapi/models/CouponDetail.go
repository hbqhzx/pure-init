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


type CouponDetail struct {

    /* 交易编号 (Optional) */
    Number string `json:"number"`

    /* 代金券编号 (Optional) */
    CouponNumber string `json:"couponNumber"`

    /* 操作类型 1生成中 2生成 3发放 4支付 5激活 6绑定 7过期 8废除9返还10退款 (Optional) */
    Type int `json:"type"`

    /* 收入（元） (Optional) */
    Income int `json:"income"`

    /* 支出（元） (Optional) */
    Expense int `json:"expense"`

    /* 余额（元） (Optional) */
    Balance int `json:"balance"`

    /* 交易时间 (Optional) */
    CreateTime string `json:"createTime"`
}
