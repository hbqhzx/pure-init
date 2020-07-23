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


type CostBillVo struct {

    /* 自增主键 (Optional) */
    Id int `json:"id"`

    /* 站点信息 (Optional) */
    Site string `json:"site"`

    /* appCode (Optional) */
    AppCode string `json:"appCode"`

    /*  (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 资源Id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 资源名称 (Optional) */
    ResourceName string `json:"resourceName"`

    /* 资源组 (Optional) */
    ResourceGroup string `json:"resourceGroup"`

    /* tag (Optional) */
    ResourceTag string `json:"resourceTag"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 配置项 (Optional) */
    Formula string `json:"formula"`

    /* 所属年月 (Optional) */
    BelongDate string `json:"belongDate"`

    /* 消费时间 (Optional) */
    ConsumeTime string `json:"consumeTime"`

    /* 消费类型 (Optional) */
    PayType string `json:"payType"`

    /* 消费明细类型 (Optional) */
    OpType string `json:"opType"`

    /* 来源数据对应唯一标识 (Optional) */
    DataSourceId int `json:"dataSourceId"`

    /* 计费类型  (包年包月、按配置、按用量、一次性计费) (Optional) */
    BillingType string `json:"billingType"`

    /* 服务开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 服务结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 总原价 (Optional) */
    TotalFee int `json:"totalFee"`

    /* 总优惠金额 (Optional) */
    TotalCouponFee int `json:"totalCouponFee"`

    /* 总应付金额 (Optional) */
    TotalPayFee int `json:"totalPayFee"`

    /* 当期应摊原价 (Optional) */
    CurrentTotalFee int `json:"currentTotalFee"`

    /* 当期应摊优惠金额 (Optional) */
    CurrentCouponFee int `json:"currentCouponFee"`

    /* 当期应摊成本 (Optional) */
    CurrentPayFee int `json:"currentPayFee"`
}
