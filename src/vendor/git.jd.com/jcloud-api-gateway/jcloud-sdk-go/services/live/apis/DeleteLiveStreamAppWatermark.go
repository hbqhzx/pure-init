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

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
)

type DeleteLiveStreamAppWatermarkRequest struct {

    core.JDCloudRequest

    /* 推流加速域名  */
    PublishDomain string `json:"publishDomain"`

    /* 直播流所属应用名称  */
    AppName string `json:"appName"`

    /* 水印模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b>  */
    Template string `json:"template"`
}

/*
 * param publishDomain: 推流加速域名 (Required)
 * param appName: 直播流所属应用名称 (Required)
 * param template: 水印模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b> (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveStreamAppWatermarkRequest(
    publishDomain string,
    appName string,
    template string,
) *DeleteLiveStreamAppWatermarkRequest {

	return &DeleteLiveStreamAppWatermarkRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/watermarkApps/{publishDomain}/appNames/{appName}/templates/{template}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        AppName: appName,
        Template: template,
	}
}

/*
 * param publishDomain: 推流加速域名 (Required)
 * param appName: 直播流所属应用名称 (Required)
 * param template: 水印模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b> (Required)
 */
func NewDeleteLiveStreamAppWatermarkRequestWithAllParams(
    publishDomain string,
    appName string,
    template string,
) *DeleteLiveStreamAppWatermarkRequest {

    return &DeleteLiveStreamAppWatermarkRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarkApps/{publishDomain}/appNames/{appName}/templates/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        AppName: appName,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveStreamAppWatermarkRequestWithoutParam() *DeleteLiveStreamAppWatermarkRequest {

    return &DeleteLiveStreamAppWatermarkRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarkApps/{publishDomain}/appNames/{appName}/templates/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流加速域名(Required) */
func (r *DeleteLiveStreamAppWatermarkRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 直播流所属应用名称(Required) */
func (r *DeleteLiveStreamAppWatermarkRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param template: 水印模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b>(Required) */
func (r *DeleteLiveStreamAppWatermarkRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveStreamAppWatermarkRequest) GetRegionId() string {
    return ""
}

type DeleteLiveStreamAppWatermarkResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveStreamAppWatermarkResult `json:"result"`
}

type DeleteLiveStreamAppWatermarkResult struct {
}