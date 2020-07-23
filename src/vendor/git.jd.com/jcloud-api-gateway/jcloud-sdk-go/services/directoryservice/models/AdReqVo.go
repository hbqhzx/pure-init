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


type AdReqVo struct {

    /* 角色  */
    RoleName string `json:"roleName"`

    /* 当前的登录用户  */
    CurrentUser string `json:"currentUser"`

    /* 区域 (Optional) */
    Region *string `json:"region"`

    /* 页码  */
    PageNum int `json:"pageNum"`

    /* 分页大小  */
    PageSize int `json:"pageSize"`

    /* 搜索词 (Optional) */
    QueryWord *string `json:"queryWord"`
}