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


type WeiXinPaymentRequest struct {

    /*  (Optional) */
    Amount string `json:"amount"`

    /*  (Optional) */
    AppCode string `json:"appCode"`

    /*  (Optional) */
    CodeUrl string `json:"codeUrl"`

    /*  (Optional) */
    CreateTime string `json:"createTime"`

    /*  (Optional) */
    Description string `json:"description"`

    /*  (Optional) */
    MerchantNo string `json:"merchantNo"`

    /*  (Optional) */
    NotifyUrl string `json:"notifyUrl"`

    /*  (Optional) */
    OrderNumber string `json:"orderNumber"`

    /*  (Optional) */
    PaymentNumber string `json:"paymentNumber"`

    /*  (Optional) */
    RequestUrl string `json:"requestUrl"`

    /*  (Optional) */
    SuccessCallbackUrl string `json:"successCallbackUrl"`
}