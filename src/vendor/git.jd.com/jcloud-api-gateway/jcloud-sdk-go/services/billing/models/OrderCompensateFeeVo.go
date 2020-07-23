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


type OrderCompensateFeeVo struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 应用码 (Optional) */
    AppCode string `json:"appCode"`

    /* 服务码 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 资源区域 (Optional) */
    Region string `json:"region"`

    /* 开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 计算总价 (Optional) */
    TotalFee int `json:"totalFee"`

    /* 计费类型 1:按配置 2:按用量 3:包年包月 4:一次性 (Optional) */
    BillingType int `json:"billingType"`
}
