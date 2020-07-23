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
    deploy "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/deploy/models"
)

type DescribeMilestoneRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* milestone Id  */
    MilestoneId string `json:"milestoneId"`
}

/*
 * param regionId: 地域ID (Required)
 * param milestoneId: milestone Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMilestoneRequest(
    regionId string,
    milestoneId string,
) *DescribeMilestoneRequest {

	return &DescribeMilestoneRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/milestone/{milestoneId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        MilestoneId: milestoneId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param milestoneId: milestone Id (Required)
 */
func NewDescribeMilestoneRequestWithAllParams(
    regionId string,
    milestoneId string,
) *DescribeMilestoneRequest {

    return &DescribeMilestoneRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/milestone/{milestoneId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        MilestoneId: milestoneId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMilestoneRequestWithoutParam() *DescribeMilestoneRequest {

    return &DescribeMilestoneRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/milestone/{milestoneId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeMilestoneRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param milestoneId: milestone Id(Required) */
func (r *DescribeMilestoneRequest) SetMilestoneId(milestoneId string) {
    r.MilestoneId = milestoneId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMilestoneRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMilestoneResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMilestoneResult `json:"result"`
}

type DescribeMilestoneResult struct {
    Milestone deploy.Milestone `json:"milestone"`
}