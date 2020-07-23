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


type ContactPersonGroup struct {

    /* 主键id (Optional) */
    Id int64 `json:"id"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 联系人id (Optional) */
    PersonId int64 `json:"personId"`

    /* 联系人姓名 (Optional) */
    UserName string `json:"userName"`

    /* 联系人手机号码 (Optional) */
    Mobile string `json:"mobile"`

    /* 联系人邮箱地址 (Optional) */
    Email string `json:"email"`

    /* 联系组id (Optional) */
    GroupId int `json:"groupId"`

    /* 联系组名称 (Optional) */
    GroupName string `json:"groupName"`

    /* 创建时间 (Optional) */
    Created string `json:"created"`

    /* 修改时间 (Optional) */
    Modified string `json:"modified"`

    /* 0表示删除，1表示正常 (Optional) */
    Yn int `json:"yn"`

    /* 是否为账号联系人 1-是 2-不是 (Optional) */
    IsSelf int `json:"isSelf"`
}
