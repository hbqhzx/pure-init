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


type SettlementDetailVo struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 结算对象pin (Optional) */
    Pin string `json:"pin"`

    /* 结算对象ID (Optional) */
    SettleObjectId string `json:"settleObjectId"`

    /* 业务系统 (Optional) */
    SystemCode string `json:"systemCode"`

    /* 业务系统名称 (Optional) */
    SystemName string `json:"systemName"`

    /* 结算对象类型 (Optional) */
    SettleObjectType string `json:"settleObjectType"`

    /* 结算对象类型名称 (Optional) */
    SettleObjectTypeName string `json:"settleObjectTypeName"`

    /* 来源数据id (Optional) */
    SourceId string `json:"sourceId"`

    /* 业务单据类型 1、其他2、订单3、退单 (Optional) */
    BusinessBillType int `json:"businessBillType"`

    /* 业务发生时间 (Optional) */
    BusinessTime string `json:"businessTime"`

    /* 交易单号 (Optional) */
    OrderNo string `json:"orderNo"`

    /* 退款单号 (Optional) */
    RefundNo string `json:"refundNo"`

    /* 结算单号 (Optional) */
    SettlementNo string `json:"settlementNo"`

    /* 费用类型 (Optional) */
    CostType string `json:"costType"`

    /* 费用类型名称 (Optional) */
    CostTypeName string `json:"costTypeName"`

    /* 收支方向  1应收 2应付 (Optional) */
    Direction int `json:"direction"`

    /* 结算项目 (Optional) */
    CommodityDesc string `json:"commodityDesc"`

    /* 数量 (Optional) */
    Count int `json:"count"`

    /* 结算单价 (Optional) */
    UnitPrice float64 `json:"unitPrice"`

    /* 结算价格 (Optional) */
    Amount float64 `json:"amount"`

    /* 税率 (Optional) */
    Taxrate float32 `json:"taxrate"`

    /* 备注 (Optional) */
    Remark string `json:"remark"`

    /* 合同主体 (Optional) */
    ContractSubject string `json:"contractSubject"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}
