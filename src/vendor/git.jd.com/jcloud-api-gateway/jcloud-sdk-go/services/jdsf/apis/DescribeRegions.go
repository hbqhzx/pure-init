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
    jdsf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdsf/models"
)

type DescribeRegionsRequest struct {

    core.JDCloudRequest

    /* 产品线信息，如果不传默认为 REGISTRY 可取值为（REGISTRY 注册中心）（OPEN_TRACE 调用链服务） (Optional) */
    ProductLine *string `json:"productLine"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRegionsRequest(
) *DescribeRegionsRequest {

	return &DescribeRegionsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param productLine: 产品线信息，如果不传默认为 REGISTRY 可取值为（REGISTRY 注册中心）（OPEN_TRACE 调用链服务） (Optional)
 */
func NewDescribeRegionsRequestWithAllParams(
    productLine *string,
) *DescribeRegionsRequest {

    return &DescribeRegionsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ProductLine: productLine,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRegionsRequestWithoutParam() *DescribeRegionsRequest {

    return &DescribeRegionsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param productLine: 产品线信息，如果不传默认为 REGISTRY 可取值为（REGISTRY 注册中心）（OPEN_TRACE 调用链服务）(Optional) */
func (r *DescribeRegionsRequest) SetProductLine(productLine string) {
    r.ProductLine = &productLine
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRegionsRequest) GetRegionId() string {
    return ""
}

type DescribeRegionsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRegionsResult `json:"result"`
}

type DescribeRegionsResult struct {
    Regions []jdsf.JdsfRegionInfo `json:"regions"`
}