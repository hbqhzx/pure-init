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

type BackendSignDeleteRequest struct {

    core.JDCloudRequest

    /* customerId (Optional) */
    CustomerId *string `json:"customerId"`

    /* method (Optional) */
    Method *string `json:"method"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBackendSignDeleteRequest(
) *BackendSignDeleteRequest {

	return &BackendSignDeleteRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/hufubackendsign",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param customerId: customerId (Optional)
 * param method: method (Optional)
 */
func NewBackendSignDeleteRequestWithAllParams(
    customerId *string,
    method *string,
) *BackendSignDeleteRequest {

    return &BackendSignDeleteRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/hufubackendsign",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        CustomerId: customerId,
        Method: method,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBackendSignDeleteRequestWithoutParam() *BackendSignDeleteRequest {

    return &BackendSignDeleteRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/hufubackendsign",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param customerId: customerId(Optional) */
func (r *BackendSignDeleteRequest) SetCustomerId(customerId string) {
    r.CustomerId = &customerId
}

/* param method: method(Optional) */
func (r *BackendSignDeleteRequest) SetMethod(method string) {
    r.Method = &method
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BackendSignDeleteRequest) GetRegionId() string {
    return ""
}

type BackendSignDeleteResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BackendSignDeleteResult `json:"result"`
}

type BackendSignDeleteResult struct {
}