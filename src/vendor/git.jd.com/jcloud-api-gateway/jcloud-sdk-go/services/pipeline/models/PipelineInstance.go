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


type PipelineInstance struct {

    /* 某一次流水线执行的uuid (Optional) */
    Uuid string `json:"uuid"`

    /* 开始执行流水线的时间 (Optional) */
    StartedAt int `json:"startedAt"`

    /* 结束执行流水线的时间 (Optional) */
    DoneAt int `json:"doneAt"`

    /* 执行持续的时间(ms) (Optional) */
    DurationMs int `json:"durationMs"`

    /* 执行状态 (Optional) */
    Status int `json:"status"`

    /* 执行状态描述 (Optional) */
    StatusHuman string `json:"statusHuman"`

    /* 执行时环境变量 (Optional) */
    Env string `json:"env"`

    /* 失败原因 (Optional) */
    FailureReason int `json:"failureReason"`

    /* 失败原因描述 (Optional) */
    FalseilureReasonHuman string `json:"falseilureReasonHuman"`

    /*  (Optional) */
    Stages []PipelineStage `json:"stages"`
}
