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

type ContentByCatalogIdRequest struct {

    core.JDCloudRequest

    /* catalogId  */
    CatalogId int `json:"catalogId"`
}

/*
 * param catalogId: catalogId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewContentByCatalogIdRequest(
    catalogId int,
) *ContentByCatalogIdRequest {

	return &ContentByCatalogIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/portal/api/help/contentByCatalogId",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        CatalogId: catalogId,
	}
}

/*
 * param catalogId: catalogId (Required)
 */
func NewContentByCatalogIdRequestWithAllParams(
    catalogId int,
) *ContentByCatalogIdRequest {

    return &ContentByCatalogIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/help/contentByCatalogId",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        CatalogId: catalogId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewContentByCatalogIdRequestWithoutParam() *ContentByCatalogIdRequest {

    return &ContentByCatalogIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/help/contentByCatalogId",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param catalogId: catalogId(Required) */
func (r *ContentByCatalogIdRequest) SetCatalogId(catalogId int) {
    r.CatalogId = catalogId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ContentByCatalogIdRequest) GetRegionId() string {
    return ""
}

type ContentByCatalogIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ContentByCatalogIdResult `json:"result"`
}

type ContentByCatalogIdResult struct {
    Content cms.Content `json:"content"`
    Catalog cms.Catalog `json:"catalog"`
    ContentList []cms.Content `json:"contentList"`
}