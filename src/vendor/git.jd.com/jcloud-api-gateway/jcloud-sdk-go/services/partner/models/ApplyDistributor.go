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


type ApplyDistributor struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 申请时间 (Optional) */
    ApplyTime string `json:"applyTime"`

    /* 联系人姓名 (Optional) */
    Contactor string `json:"contactor"`

    /* 联系人电话 (Optional) */
    Tel string `json:"tel"`

    /* 联系邮箱 (Optional) */
    Email string `json:"email"`

    /* 合作类型 (Optional) */
    CooperateType string `json:"cooperateType"`

    /* 合作类型名称 (Optional) */
    CooperateTypeName string `json:"cooperateTypeName"`

    /* 公司名称 (Optional) */
    CompanyName string `json:"companyName"`

    /* 公司规模 (Optional) */
    CompanyScale string `json:"companyScale"`

    /* 公司规模 (Optional) */
    CompanyScaleName string `json:"companyScaleName"`

    /* 业务区域 (Optional) */
    Region string `json:"region"`

    /* 业务区域 (Optional) */
    RegionName string `json:"regionName"`

    /* 公司网址 (Optional) */
    WebSite string `json:"webSite"`

    /* 状态 (Optional) */
    Status int `json:"status"`

    /* 状态 (Optional) */
    StatusName string `json:"statusName"`
}
