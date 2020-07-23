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


type LowestDiscountVo struct {

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 产品线编码  */
    AppCode string `json:"appCode"`

    /* 产品编码 多个用|分隔  */
    ServiceCode string `json:"serviceCode"`

    /* 计费类型  2（包年包月）  */
    BillType int `json:"billType"`

    /* 购买时长，多个之间用逗号分隔，以月为单位  */
    BillPeriod string `json:"billPeriod"`

    /* 单位 时长类型 1:小时 2:天 3:月 4:年(注：小时和天目前不支持查询)  */
    TimeUnit int `json:"timeUnit"`

    /* 区域  */
    Region string `json:"region"`

    /* 购买方式:1新购，2续费；3配置变更  */
    BuyType int `json:"buyType"`
}