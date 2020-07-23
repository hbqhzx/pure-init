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


type ApiGroupMonitorDetail struct {

    /* 分组名称 (Optional) */
    GroupName string `json:"groupName"`

    /* 后端服务地址 (Optional) */
    BackendUrl string `json:"backendUrl"`

    /* 分组路径 (Optional) */
    GroupPath string `json:"groupPath"`

    /* 版本号 (Optional) */
    Version string `json:"version"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 发布日期，格式为毫秒级时间戳 (Optional) */
    DeploymentDate int64 `json:"deploymentDate"`
}
