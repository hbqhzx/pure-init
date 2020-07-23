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


type OrganizationVo struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 主体id (Optional) */
    ContractSubject string `json:"contractSubject"`

    /* org_id (Optional) */
    OrgId string `json:"orgId"`

    /* 公司名称 (Optional) */
    CompanyName string `json:"companyName"`

    /* 纳税人识别码 (Optional) */
    TaxpayerIdCode string `json:"taxpayerIdCode"`

    /* 地址电话 (Optional) */
    Address string `json:"address"`

    /* 开户行及账号 (Optional) */
    BankAccount string `json:"bankAccount"`
}
