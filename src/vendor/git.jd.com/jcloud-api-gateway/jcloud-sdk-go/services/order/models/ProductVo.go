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


type ProductVo struct {

    /*  (Optional) */
    ChargeMode int `json:"chargeMode"`

    /*  (Optional) */
    ChargeDuration int `json:"chargeDuration"`

    /*  (Optional) */
    ChargeUnit int `json:"chargeUnit"`

    /*  (Optional) */
    ItemId string `json:"itemId"`

    /*  (Optional) */
    AppCode string `json:"appCode"`

    /*  (Optional) */
    ServiceCode string `json:"serviceCode"`

    /*  (Optional) */
    Region string `json:"region"`

    /*  (Optional) */
    Formula string `json:"formula"`

    /*  (Optional) */
    NetworkOperator int `json:"networkOperator"`

    /*  (Optional) */
    OperatorType int `json:"operatorType"`

    /*  (Optional) */
    OperatorName string `json:"operatorName"`

    /*  (Optional) */
    ItemName string `json:"itemName"`

    /*  (Optional) */
    ItemType int `json:"itemType"`

    /*  (Optional) */
    ExtraInfo []ExtraInfoVo `json:"extraInfo"`

    /*  (Optional) */
    Quantity int `json:"quantity"`

    /*  (Optional) */
    ActualFee string `json:"actualFee"`

    /*  (Optional) */
    DiscountFee string `json:"discountFee"`

    /*  (Optional) */
    TotalFee string `json:"totalFee"`

    /*  (Optional) */
    PriceSnapshot string `json:"priceSnapshot"`

    /*  (Optional) */
    Unit string `json:"unit"`

    /*  (Optional) */
    StartTime string `json:"startTime"`

    /*  (Optional) */
    EndTime string `json:"endTime"`

    /*  (Optional) */
    ResizeFormulaType int `json:"resizeFormulaType"`

    /*  (Optional) */
    PromotionInfo string `json:"promotionInfo"`

    /*  (Optional) */
    CustomInfo string `json:"customInfo"`

    /*  (Optional) */
    AzMap string `json:"azMap"`
}
