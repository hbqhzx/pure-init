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

type DeleteDatasourceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* channel ID  */
    Id string `json:"id"`
}

/*
 * param regionId: 地域ID (Required)
 * param id: channel ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteDatasourceRequest(
    regionId string,
    id string,
) *DeleteDatasourceRequest {

	return &DeleteDatasourceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/migration_mysqlDatasources/{id}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param id: channel ID (Required)
 */
func NewDeleteDatasourceRequestWithAllParams(
    regionId string,
    id string,
) *DeleteDatasourceRequest {

    return &DeleteDatasourceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/migration_mysqlDatasources/{id}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteDatasourceRequestWithoutParam() *DeleteDatasourceRequest {

    return &DeleteDatasourceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/migration_mysqlDatasources/{id}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteDatasourceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: channel ID(Required) */
func (r *DeleteDatasourceRequest) SetId(id string) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteDatasourceRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteDatasourceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteDatasourceResult `json:"result"`
}

type DeleteDatasourceResult struct {
}