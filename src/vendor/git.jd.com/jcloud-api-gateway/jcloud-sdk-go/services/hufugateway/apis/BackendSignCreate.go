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

type BackendSignCreateRequest struct {

    core.JDCloudRequest

    /* customerId  */
    CustomerId string `json:"customerId"`

    /* method  */
    Method string `json:"method"`

    /* accessKey  */
    AccessKey string `json:"accessKey"`

    /* secretKey (Optional) */
    SecretKey *string `json:"secretKey"`

    /* jd_cloud 京东云用户
jd_apikms 京东云apikms
jd_hufu   虎符ak
 (Optional) */
    KeyType *string `json:"keyType"`
}

/*
 * param customerId: customerId (Required)
 * param method: method (Required)
 * param accessKey: accessKey (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBackendSignCreateRequest(
    customerId string,
    method string,
    accessKey string,
) *BackendSignCreateRequest {

	return &BackendSignCreateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/hufubackendsign",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CustomerId: customerId,
        Method: method,
        AccessKey: accessKey,
	}
}

/*
 * param customerId: customerId (Required)
 * param method: method (Required)
 * param accessKey: accessKey (Required)
 * param secretKey: secretKey (Optional)
 * param keyType: jd_cloud 京东云用户
jd_apikms 京东云apikms
jd_hufu   虎符ak
 (Optional)
 */
func NewBackendSignCreateRequestWithAllParams(
    customerId string,
    method string,
    accessKey string,
    secretKey *string,
    keyType *string,
) *BackendSignCreateRequest {

    return &BackendSignCreateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/hufubackendsign",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CustomerId: customerId,
        Method: method,
        AccessKey: accessKey,
        SecretKey: secretKey,
        KeyType: keyType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBackendSignCreateRequestWithoutParam() *BackendSignCreateRequest {

    return &BackendSignCreateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/hufubackendsign",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param customerId: customerId(Required) */
func (r *BackendSignCreateRequest) SetCustomerId(customerId string) {
    r.CustomerId = customerId
}

/* param method: method(Required) */
func (r *BackendSignCreateRequest) SetMethod(method string) {
    r.Method = method
}

/* param accessKey: accessKey(Required) */
func (r *BackendSignCreateRequest) SetAccessKey(accessKey string) {
    r.AccessKey = accessKey
}

/* param secretKey: secretKey(Optional) */
func (r *BackendSignCreateRequest) SetSecretKey(secretKey string) {
    r.SecretKey = &secretKey
}

/* param keyType: jd_cloud 京东云用户
jd_apikms 京东云apikms
jd_hufu   虎符ak
(Optional) */
func (r *BackendSignCreateRequest) SetKeyType(keyType string) {
    r.KeyType = &keyType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BackendSignCreateRequest) GetRegionId() string {
    return ""
}

type BackendSignCreateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BackendSignCreateResult `json:"result"`
}

type BackendSignCreateResult struct {
    Id string `json:"id"`
    AccessKey string `json:"accessKey"`
    SecretKey string `json:"secretKey"`
    KeyType string `json:"keyType"`
    CreateTime string `json:"createTime"`
}