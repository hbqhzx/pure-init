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


type Database struct {

    /* 数据库名称 (Optional) */
    DatabaseName string `json:"databaseName"`

    /* 数据库字符集 (Optional) */
    DatabaseCharset string `json:"databaseCharset"`

    /* 数据库状态 (Optional) */
    DatabaseStatus string `json:"databaseStatus"`

    /* 数据库权限列表 (Optional) */
    DatabasePrivileges []interface{} `json:"databasePrivileges"`

    /* 保护标志 (Optional) */
    Protection bool `json:"protection"`

    /* 创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 更新时间 (Optional) */
    UpdatedTime string `json:"updatedTime"`
}