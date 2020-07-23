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

type DeleteLiveStreamAppTranscodeRequest struct {

    core.JDCloudRequest

    /* 推流加速域名  */
    PublishDomain string `json:"publishDomain"`

    /* 直播流所属应用名称  */
    AppName string `json:"appName"`

    /* 转码模板自定义名称:
  - 标准质量模板：ld(h.264/640*360/15f)<br>sd(h.264/854*480/24f)<br>hd(h.264/1280*720/25f)<br>shd(h.264/1920*1080/30f)
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与标准的转码模板和已定义命名重复</b>  */
    Template string `json:"template"`
}

/*
 * param publishDomain: 推流加速域名 (Required)
 * param appName: 直播流所属应用名称 (Required)
 * param template: 转码模板自定义名称:
  - 标准质量模板：ld(h.264/640*360/15f)<br>sd(h.264/854*480/24f)<br>hd(h.264/1280*720/25f)<br>shd(h.264/1920*1080/30f)
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与标准的转码模板和已定义命名重复</b> (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveStreamAppTranscodeRequest(
    publishDomain string,
    appName string,
    template string,
) *DeleteLiveStreamAppTranscodeRequest {

	return &DeleteLiveStreamAppTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodeApps/{publishDomain}/appNames/{appName}/templates/{template}",
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
 * param template: 转码模板自定义名称:
  - 标准质量模板：ld(h.264/640*360/15f)<br>sd(h.264/854*480/24f)<br>hd(h.264/1280*720/25f)<br>shd(h.264/1920*1080/30f)
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与标准的转码模板和已定义命名重复</b> (Required)
 */
func NewDeleteLiveStreamAppTranscodeRequestWithAllParams(
    publishDomain string,
    appName string,
    template string,
) *DeleteLiveStreamAppTranscodeRequest {

    return &DeleteLiveStreamAppTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeApps/{publishDomain}/appNames/{appName}/templates/{template}",
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
func NewDeleteLiveStreamAppTranscodeRequestWithoutParam() *DeleteLiveStreamAppTranscodeRequest {

    return &DeleteLiveStreamAppTranscodeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeApps/{publishDomain}/appNames/{appName}/templates/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流加速域名(Required) */
func (r *DeleteLiveStreamAppTranscodeRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 直播流所属应用名称(Required) */
func (r *DeleteLiveStreamAppTranscodeRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param template: 转码模板自定义名称:
  - 标准质量模板：ld(h.264/640*360/15f)<br>sd(h.264/854*480/24f)<br>hd(h.264/1280*720/25f)<br>shd(h.264/1920*1080/30f)
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与标准的转码模板和已定义命名重复</b>(Required) */
func (r *DeleteLiveStreamAppTranscodeRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveStreamAppTranscodeRequest) GetRegionId() string {
    return ""
}

type DeleteLiveStreamAppTranscodeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveStreamAppTranscodeResult `json:"result"`
}

type DeleteLiveStreamAppTranscodeResult struct {
}