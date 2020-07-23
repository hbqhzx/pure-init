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


type Connection struct {

    /* Connection的Id (Optional) */
    ConnectionId string `json:"connectionId"`

    /* 专线的名称, 只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional) */
    ConnectionName string `json:"connectionName"`

    /* 自助专线开通的地域信息；只有自助专线才有location信息 (Optional) */
    Location Location `json:"location"`

    /* 专线服务商的信息;只在type为jcloud_partner时有合作伙伴信息详情 (Optional) */
    Partner Partner `json:"partner"`

    /* 类型, jcloud_hosted:托管连接; jcloud_partner:合作伙伴连接; jcloud:自助连接。 (Optional) */
    Type string `json:"type"`

    /* 申请的专线带宽：Mbps (Optional) */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* 订单编号 (Optional) */
    OrderId string `json:"orderId"`

    /* connection上通道Id列表 (Optional) */
    VifIds []string `json:"vifIds"`

    /* 专线的状态，取值为：Ordering,Rejected,Pending,Confirming,Paying,Active,Deleting (Optional) */
    ConnectionStatus string `json:"connectionStatus"`

    /* Bgw的描述, 允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description string `json:"description"`

    /* 专线申请的时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 客户idc地址，取值范围：1~100个字符 (Optional) */
    IdcAddress string `json:"idcAddress"`

    /* 联系人名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional) */
    ContactName string `json:"contactName"`

    /* 联系人手机号码，11位数字且需要13、14、15、16、17、18、19开头 (Optional) */
    PhoneNumber string `json:"phoneNumber"`

    /* 审核不通过的原因 (Optional) */
    RejectedReason string `json:"rejectedReason"`
}
