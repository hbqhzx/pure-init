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

type DynamicProductListRequest struct {

    core.JDCloudRequest

    /* productId  */
    ProductId int `json:"productId"`

    /* pageNum (Optional) */
    PageNum *int `json:"pageNum"`

    /* pageSize  */
    PageSize int `json:"pageSize"`
}

/*
 * param productId: productId (Required)
 * param pageSize: pageSize (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDynamicProductListRequest(
    productId int,
    pageSize int,
) *DynamicProductListRequest {

	return &DynamicProductListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/portal/api/product/dynamicProductList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ProductId: productId,
        PageSize: pageSize,
	}
}

/*
 * param productId: productId (Required)
 * param pageNum: pageNum (Optional)
 * param pageSize: pageSize (Required)
 */
func NewDynamicProductListRequestWithAllParams(
    productId int,
    pageNum *int,
    pageSize int,
) *DynamicProductListRequest {

    return &DynamicProductListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/product/dynamicProductList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ProductId: productId,
        PageNum: pageNum,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDynamicProductListRequestWithoutParam() *DynamicProductListRequest {

    return &DynamicProductListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/product/dynamicProductList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param productId: productId(Required) */
func (r *DynamicProductListRequest) SetProductId(productId int) {
    r.ProductId = productId
}

/* param pageNum: pageNum(Optional) */
func (r *DynamicProductListRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: pageSize(Required) */
func (r *DynamicProductListRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DynamicProductListRequest) GetRegionId() string {
    return ""
}

type DynamicProductListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DynamicProductListResult `json:"result"`
}

type DynamicProductListResult struct {
    PageInfo cms.PageInfo `json:"pageInfo"`
    List []cms.ProductDynamic `json:"list"`
}