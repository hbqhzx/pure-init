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

type CreateMigrationFlagRequest struct {

    core.JDCloudRequest

    /* Region ID，例如：cn-north-1  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: Region ID，例如：cn-north-1 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateMigrationFlagRequest(
    regionId string,
) *CreateMigrationFlagRequest {

	return &CreateMigrationFlagRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/migration",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID，例如：cn-north-1 (Required)
 */
func NewCreateMigrationFlagRequestWithAllParams(
    regionId string,
) *CreateMigrationFlagRequest {

    return &CreateMigrationFlagRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/migration",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateMigrationFlagRequestWithoutParam() *CreateMigrationFlagRequest {

    return &CreateMigrationFlagRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/migration",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID，例如：cn-north-1(Required) */
func (r *CreateMigrationFlagRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateMigrationFlagRequest) GetRegionId() string {
    return r.RegionId
}

type CreateMigrationFlagResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateMigrationFlagResult `json:"result"`
}

type CreateMigrationFlagResult struct {
}