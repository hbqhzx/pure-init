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

type DeleteSettleObjectTypeRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* id  */
    Id int `json:"id"`
}

/*
 * param regionId:  (Required)
 * param id: id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteSettleObjectTypeRequest(
    regionId string,
    id int,
) *DeleteSettleObjectTypeRequest {

	return &DeleteSettleObjectTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/systemConfig/{id}:settleObjectTypeDelete",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId:  (Required)
 * param id: id (Required)
 */
func NewDeleteSettleObjectTypeRequestWithAllParams(
    regionId string,
    id int,
) *DeleteSettleObjectTypeRequest {

    return &DeleteSettleObjectTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/systemConfig/{id}:settleObjectTypeDelete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteSettleObjectTypeRequestWithoutParam() *DeleteSettleObjectTypeRequest {

    return &DeleteSettleObjectTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/systemConfig/{id}:settleObjectTypeDelete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *DeleteSettleObjectTypeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: id(Required) */
func (r *DeleteSettleObjectTypeRequest) SetId(id int) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteSettleObjectTypeRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteSettleObjectTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteSettleObjectTypeResult `json:"result"`
}

type DeleteSettleObjectTypeResult struct {
    Result bool `json:"result"`
}