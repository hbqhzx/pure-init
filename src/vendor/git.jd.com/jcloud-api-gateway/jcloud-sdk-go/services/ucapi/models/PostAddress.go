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


type PostAddress struct {

    /* 发票邮寄地址id (Optional) */
    Id *int `json:"id"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 收件人姓名 (Optional) */
    Name *string `json:"name"`

    /* 收件人手机号码 (Optional) */
    Phone *string `json:"phone"`

    /* 省 (Optional) */
    Province *string `json:"province"`

    /* 市 (Optional) */
    City *string `json:"city"`

    /* 区/县 (Optional) */
    Town *string `json:"town"`

    /* 详细地址 (Optional) */
    Address *string `json:"address"`

    /* 是否为默认地址 1-默认地址 2-非默认地址 (Optional) */
    IsDefault *int `json:"isDefault"`

    /* 创建时间,格式为yyyy-MM-dd HH:mm:ss (Optional) */
    CreateTime *string `json:"createTime"`

    /* 更新时间,格式为yyyy-MM-dd HH:mm:ss (Optional) */
    UpdateTime *string `json:"updateTime"`

    /* 邮编 (Optional) */
    ZipCode *string `json:"zipCode"`
}