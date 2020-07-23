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

package models


type ActivityInfo struct {

    /* 标题 (Optional) */
    Title string `json:"title"`

    /* 关键字 (Optional) */
    Keywords string `json:"keywords"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 0-活动已关闭，1-活动已暂停，3-活动未开始，4-活动进行中 (Optional) */
    Status int `json:"status"`

    /* 地域信息 (Optional) */
    RegionInfo []RegionInfo `json:"regionInfo"`
}