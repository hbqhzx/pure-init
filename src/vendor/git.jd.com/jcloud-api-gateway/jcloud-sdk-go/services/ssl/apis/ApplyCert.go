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

type ApplyCertRequest struct {

    core.JDCloudRequest

    /* 申购记录Id  */
    RecordId string `json:"recordId"`

    /* 证书申购类型  */
    CertType string `json:"certType"`

    /* 证书品牌  */
    Brand string `json:"brand"`

    /* 证书年限  */
    CertValidity int `json:"certValidity"`

    /* 域名数量,此字段为非必需字段，默认为1 (Optional) */
    DomainCount *int `json:"domainCount"`
}

/*
 * param recordId: 申购记录Id (Required)
 * param certType: 证书申购类型 (Required)
 * param brand: 证书品牌 (Required)
 * param certValidity: 证书年限 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewApplyCertRequest(
    recordId string,
    certType string,
    brand string,
    certValidity int,
) *ApplyCertRequest {

	return &ApplyCertRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sslRecord:applyCert",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RecordId: recordId,
        CertType: certType,
        Brand: brand,
        CertValidity: certValidity,
	}
}

/*
 * param recordId: 申购记录Id (Required)
 * param certType: 证书申购类型 (Required)
 * param brand: 证书品牌 (Required)
 * param certValidity: 证书年限 (Required)
 * param domainCount: 域名数量,此字段为非必需字段，默认为1 (Optional)
 */
func NewApplyCertRequestWithAllParams(
    recordId string,
    certType string,
    brand string,
    certValidity int,
    domainCount *int,
) *ApplyCertRequest {

    return &ApplyCertRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslRecord:applyCert",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RecordId: recordId,
        CertType: certType,
        Brand: brand,
        CertValidity: certValidity,
        DomainCount: domainCount,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewApplyCertRequestWithoutParam() *ApplyCertRequest {

    return &ApplyCertRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslRecord:applyCert",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param recordId: 申购记录Id(Required) */
func (r *ApplyCertRequest) SetRecordId(recordId string) {
    r.RecordId = recordId
}

/* param certType: 证书申购类型(Required) */
func (r *ApplyCertRequest) SetCertType(certType string) {
    r.CertType = certType
}

/* param brand: 证书品牌(Required) */
func (r *ApplyCertRequest) SetBrand(brand string) {
    r.Brand = brand
}

/* param certValidity: 证书年限(Required) */
func (r *ApplyCertRequest) SetCertValidity(certValidity int) {
    r.CertValidity = certValidity
}

/* param domainCount: 域名数量,此字段为非必需字段，默认为1(Optional) */
func (r *ApplyCertRequest) SetDomainCount(domainCount int) {
    r.DomainCount = &domainCount
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ApplyCertRequest) GetRegionId() string {
    return ""
}

type ApplyCertResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ApplyCertResult `json:"result"`
}

type ApplyCertResult struct {
    RecordId string `json:"recordId"`
    OrderId string `json:"orderId"`
}