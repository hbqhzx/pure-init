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


type UserInfoRes struct {

    /*  (Optional) */
    CNumber string `json:"cNumber"`

    /*  (Optional) */
    CType int `json:"cType"`

    /* 云服务账号 (Optional) */
    AccountId string `json:"accountId"`

    /* 联系地址 (Optional) */
    ContactAddress string `json:"contactAddress"`

    /* 联系邮箱 (Optional) */
    ContactEmail string `json:"contactEmail"`

    /* 联系人 (Optional) */
    ContactPeople string `json:"contactPeople"`

    /* 联系电话 (Optional) */
    ContactTelephone string `json:"contactTelephone"`

    /* 甲方名称 (Optional) */
    Name string `json:"name"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 合同号 (Optional) */
    ContractNumber string `json:"contractNumber"`

    /* 身份证 (Optional) */
    ContactIdentification string `json:"contactIdentification"`

    /* 1：个人，2：企业 (Optional) */
    Type int `json:"type"`
}
