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

import nc "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/nc/models"

type Secret struct {

    /* 机密数据名称 (Optional) */
    Name string `json:"name"`

    /* 私密数据的类型，目前仅支持如下类型：docker-registry：用来和docker registry认证的类型 (Optional) */
    Type string `json:"type"`

    /* 创建时间 (Optional) */
    CreatedAt string `json:"createdAt"`

    /* 机密的数据 (Optional) */
    Data nc.DockerRegistryData `json:"data"`
}