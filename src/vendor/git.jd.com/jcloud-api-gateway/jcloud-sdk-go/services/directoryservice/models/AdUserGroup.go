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


type AdUserGroup struct {

    /* 用户名 (Optional) */
    UserName string `json:"userName"`

    /* 用户全名 (Optional) */
    UserFullName string `json:"userFullName"`

    /* 域名 (Optional) */
    DomainName string `json:"domainName"`

    /* 类型 0 用户 1 组 (Optional) */
    Type int `json:"type"`

    /* sid (Optional) */
    Sid string `json:"sid"`
}
