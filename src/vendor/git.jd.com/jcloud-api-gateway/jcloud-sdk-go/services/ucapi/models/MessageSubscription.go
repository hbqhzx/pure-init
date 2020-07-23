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


type MessageSubscription struct {

    /* id (Optional) */
    Id int64 `json:"id"`

    /* 创建人 (Optional) */
    Pin string `json:"pin"`

    /* 消息设置id (Optional) */
    MessageSettingId int64 `json:"messageSettingId"`

    /* 消息接收人/组 id/主联系人账号id (Optional) */
    ReceiverId int64 `json:"receiverId"`

    /* 接收人/组名称 (Optional) */
    ReceiverName string `json:"receiverName"`

    /* 1-账号联系人; 2- 非账号联系人 (Optional) */
    IsSelf int `json:"isSelf"`

    /* 创建时间 (Optional) */
    Created string `json:"created"`

    /* 修改时间 (Optional) */
    Modified string `json:"modified"`

    /* 是否删除，0删除，1未删除 (Optional) */
    Yn int `json:"yn"`

    /* 接收人类型 1-联系人 2-联系组 (Optional) */
    ReceiverType int `json:"receiverType"`

    /* 类别 code (Optional) */
    CategoryCode string `json:"categoryCode"`
}
