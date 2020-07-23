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


type UserInformationReq struct {

    /* 联系地址 (Optional) */
    ContactAddress *string `json:"contactAddress"`

    /* 联系邮箱 (Optional) */
    ContactEmail *string `json:"contactEmail"`

    /* 联系人 (Optional) */
    ContactPeople *string `json:"contactPeople"`

    /* 联系电话 (Optional) */
    ContactTelephone *string `json:"contactTelephone"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 省份证号 (Optional) */
    ContactIdentification *string `json:"contactIdentification"`
}