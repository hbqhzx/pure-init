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


type OffLineRimmitBill struct {

    /* 汇款人账号名称 (Optional) */
    RemittorAccount string `json:"remittorAccount"`

    /* 汇款银行账号 (Optional) */
    RemittorBankAccount string `json:"remittorBankAccount"`

    /* 汇款金额 (Optional) */
    RemitAmount int `json:"remitAmount"`

    /* 汇款银行名称 (Optional) */
    RemitBankName string `json:"remitBankName"`

    /* 汇款时间 (Optional) */
    RemitTime string `json:"remitTime"`

    /* 联系人手机 (Optional) */
    ContactsPhone string `json:"contactsPhone"`

    /* 已完成:0 待审核:1 已驳回:2 (Optional) */
    Status int `json:"status"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}
