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
    monitor "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type DeleteSiteMonitorRequest struct {

    core.JDCloudRequest

    /* name为'list' - 站点监控id (Optional) */
    Filters []monitor.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteSiteMonitorRequest(
) *DeleteSiteMonitorRequest {

	return &DeleteSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/siteMonitor",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param filters: name为'list' - 站点监控id (Optional)
 */
func NewDeleteSiteMonitorRequestWithAllParams(
    filters []monitor.Filter,
) *DeleteSiteMonitorRequest {

    return &DeleteSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteSiteMonitorRequestWithoutParam() *DeleteSiteMonitorRequest {

    return &DeleteSiteMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param filters: name为'list' - 站点监控id(Optional) */
func (r *DeleteSiteMonitorRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteSiteMonitorRequest) GetRegionId() string {
    return ""
}

type DeleteSiteMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteSiteMonitorResult `json:"result"`
}

type DeleteSiteMonitorResult struct {
    Suc bool `json:"suc"`
}