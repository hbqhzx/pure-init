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


type AssumeRoleInfo struct {

    /* 角色资源描述  */
    RoleJrn string `json:"roleJrn"`

    /* 角色名称  */
    RoleSessionName string `json:"roleSessionName"`

    /* 权限信息 (Optional) */
    Policy *string `json:"policy"`

    /* 令牌有效期，单位秒，默认3600 (Optional) */
    DurationSeconds *int `json:"durationSeconds"`
}
