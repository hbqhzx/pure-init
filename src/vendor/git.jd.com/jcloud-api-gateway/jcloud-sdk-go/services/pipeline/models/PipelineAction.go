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


type PipelineAction struct {

    /* 操作(action)的UUID (Optional) */
    Uuid string `json:"uuid"`

    /* 操作(action)的名称 (Optional) */
    Name string `json:"name"`

    /* 操作(action)创建时间 (Optional) */
    CreatedAt int `json:"createdAt"`

    /* 操作(action)开始时间 (Optional) */
    StartedAt int `json:"startedAt"`

    /* 操作(action)结束时间 (Optional) */
    DoneAt int `json:"doneAt"`

    /* 操作(action)当前状态 (Optional) */
    Status int `json:"status"`

    /* 操作(action)当前状态说明 (Optional) */
    StatusHuman string `json:"statusHuman"`

    /*  (Optional) */
    Links ActionLinks `json:"links"`

    /*  (Optional) */
    ActionTypeId ActionTypeId `json:"actionTypeId"`
}
