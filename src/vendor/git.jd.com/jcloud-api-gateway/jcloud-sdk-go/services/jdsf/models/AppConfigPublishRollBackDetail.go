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


type AppConfigPublishRollBackDetail struct {

    /* 配置名称 (Optional) */
    AppConfigName string `json:"appConfigName"`

    /* 应用名称 (Optional) */
    AppName string `json:"appName"`

    /* 发布信息 id (Optional) */
    AppConfigPublishVersionId string `json:"appConfigPublishVersionId"`

    /* 发布应用版本 (Optional) */
    AppConfigVersion string `json:"appConfigVersion"`

    /* 版本配置内容 (Optional) */
    AppConfigContent string `json:"appConfigContent"`

    /* 版本备注 (Optional) */
    AppConfigRemark string `json:"appConfigRemark"`

    /* 当发布方式使用注册中心时，此处未注册中心实例id (Optional) */
    PublishSourceInfo string `json:"publishSourceInfo"`
}
