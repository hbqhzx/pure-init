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

type DeleteCustomLiveStreamSnapshotTemplateRequest struct {

    core.JDCloudRequest

    /* 截图模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b>  */
    Template string `json:"template"`
}

/*
 * param template: 截图模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b> (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteCustomLiveStreamSnapshotTemplateRequest(
    template string,
) *DeleteCustomLiveStreamSnapshotTemplateRequest {

	return &DeleteCustomLiveStreamSnapshotTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/snapshotCustoms/{template}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        Template: template,
	}
}

/*
 * param template: 截图模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b> (Required)
 */
func NewDeleteCustomLiveStreamSnapshotTemplateRequestWithAllParams(
    template string,
) *DeleteCustomLiveStreamSnapshotTemplateRequest {

    return &DeleteCustomLiveStreamSnapshotTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotCustoms/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteCustomLiveStreamSnapshotTemplateRequestWithoutParam() *DeleteCustomLiveStreamSnapshotTemplateRequest {

    return &DeleteCustomLiveStreamSnapshotTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotCustoms/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param template: 截图模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b>(Required) */
func (r *DeleteCustomLiveStreamSnapshotTemplateRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteCustomLiveStreamSnapshotTemplateRequest) GetRegionId() string {
    return ""
}

type DeleteCustomLiveStreamSnapshotTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteCustomLiveStreamSnapshotTemplateResult `json:"result"`
}

type DeleteCustomLiveStreamSnapshotTemplateResult struct {
}