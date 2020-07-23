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
    cms "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/cms/models"
)

type CatalogByPathRequest struct {

    core.JDCloudRequest

    /* 路径  */
    Path string `json:"path"`
}

/*
 * param path: 路径 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCatalogByPathRequest(
    path string,
) *CatalogByPathRequest {

	return &CatalogByPathRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/portal/api/help/catalogByPath",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Path: path,
	}
}

/*
 * param path: 路径 (Required)
 */
func NewCatalogByPathRequestWithAllParams(
    path string,
) *CatalogByPathRequest {

    return &CatalogByPathRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/help/catalogByPath",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Path: path,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCatalogByPathRequestWithoutParam() *CatalogByPathRequest {

    return &CatalogByPathRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/help/catalogByPath",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param path: 路径(Required) */
func (r *CatalogByPathRequest) SetPath(path string) {
    r.Path = path
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CatalogByPathRequest) GetRegionId() string {
    return ""
}

type CatalogByPathResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CatalogByPathResult `json:"result"`
}

type CatalogByPathResult struct {
    Catalog cms.Catalog `json:"catalog"`
}