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
    jdfusion "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdfusion/models"
)

type GetVmImagesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 镜像来源： system：系统官方公共镜像。 self：用户自定义镜像。 others：用户共享的镜像。 marketplace：镜像市场云市场 提供的镜像。 (Optional) */
    ImageSource *string `json:"imageSource"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetVmImagesRequest(
    regionId string,
) *GetVmImagesRequest {

	return &GetVmImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vm_images",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param imageSource: 镜像来源： system：系统官方公共镜像。 self：用户自定义镜像。 others：用户共享的镜像。 marketplace：镜像市场云市场 提供的镜像。 (Optional)
 */
func NewGetVmImagesRequestWithAllParams(
    regionId string,
    imageSource *string,
) *GetVmImagesRequest {

    return &GetVmImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageSource: imageSource,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetVmImagesRequestWithoutParam() *GetVmImagesRequest {

    return &GetVmImagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetVmImagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageSource: 镜像来源： system：系统官方公共镜像。 self：用户自定义镜像。 others：用户共享的镜像。 marketplace：镜像市场云市场 提供的镜像。(Optional) */
func (r *GetVmImagesRequest) SetImageSource(imageSource string) {
    r.ImageSource = &imageSource
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetVmImagesRequest) GetRegionId() string {
    return r.RegionId
}

type GetVmImagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetVmImagesResult `json:"result"`
}

type GetVmImagesResult struct {
    Images []jdfusion.ImageInfo `json:"images"`
}