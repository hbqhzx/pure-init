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


type BasicConfig struct {

    /* 启用网格白名单检查  */
    EnableWhiteListCheck bool `json:"enableWhiteListCheck"`

    /* 启用被动健康检查  */
    EnableNegativeHealthCheck bool `json:"enableNegativeHealthCheck"`

    /* 被动健康检查参数 (Optional) */
    HealthCheckParameters HealthCheckParameters `json:"healthCheckParameters"`

    /* 网格外部服务依赖项 (Optional) */
    ServiceEntries []ServiceEntry `json:"serviceEntries"`
}
