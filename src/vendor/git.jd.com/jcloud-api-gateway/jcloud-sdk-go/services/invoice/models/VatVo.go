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


type VatVo struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 税号 (Optional) */
    TaxId string `json:"taxId"`

    /* 公司名称 (Optional) */
    Company string `json:"company"`

    /* 电话 (Optional) */
    Phone string `json:"phone"`

    /* 开户行 (Optional) */
    Bank string `json:"bank"`

    /* 开户行账户 (Optional) */
    Account string `json:"account"`

    /* 注册地址 (Optional) */
    Address string `json:"address"`

    /* 资质有效开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 资质有效结束时间 (Optional) */
    EndTime string `json:"endTime"`
}