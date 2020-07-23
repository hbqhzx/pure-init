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

type DeleteTrustRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 资源id  */
    DirectoryId string `json:"directoryId"`

    /* 信任关系id  */
    TrustId string `json:"trustId"`

    /* 当前用户 (Optional) */
    CurrentUser *string `json:"currentUser"`
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 * param trustId: 信任关系id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteTrustRequest(
    regionId string,
    directoryId string,
    trustId string,
) *DeleteTrustRequest {

	return &DeleteTrustRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/directory/{directoryId}/trust/{trustId}:deleteTrust",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DirectoryId: directoryId,
        TrustId: trustId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 * param trustId: 信任关系id (Required)
 * param currentUser: 当前用户 (Optional)
 */
func NewDeleteTrustRequestWithAllParams(
    regionId string,
    directoryId string,
    trustId string,
    currentUser *string,
) *DeleteTrustRequest {

    return &DeleteTrustRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/trust/{trustId}:deleteTrust",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DirectoryId: directoryId,
        TrustId: trustId,
        CurrentUser: currentUser,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteTrustRequestWithoutParam() *DeleteTrustRequest {

    return &DeleteTrustRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/trust/{trustId}:deleteTrust",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteTrustRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param directoryId: 资源id(Required) */
func (r *DeleteTrustRequest) SetDirectoryId(directoryId string) {
    r.DirectoryId = directoryId
}

/* param trustId: 信任关系id(Required) */
func (r *DeleteTrustRequest) SetTrustId(trustId string) {
    r.TrustId = trustId
}

/* param currentUser: 当前用户(Optional) */
func (r *DeleteTrustRequest) SetCurrentUser(currentUser string) {
    r.CurrentUser = &currentUser
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteTrustRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteTrustResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteTrustResult `json:"result"`
}

type DeleteTrustResult struct {
}