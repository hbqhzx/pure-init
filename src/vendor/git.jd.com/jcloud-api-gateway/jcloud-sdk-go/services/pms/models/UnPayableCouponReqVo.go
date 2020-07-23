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


type UnPayableCouponReqVo struct {

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 产品线编码  */
    AppCode string `json:"appCode"`

    /* 产品编码,多产品以竖线&#124;分隔  */
    ServiceCode string `json:"serviceCode"`

    /* 计费类型  1-配置,2-包年包月,4-按用量,8-按次数  */
    BillingType string `json:"billingType"`
}
