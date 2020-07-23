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


type ScheduledTask struct {

    /* 主键Id (Optional) */
    TaskId int64 `json:"taskId"`

    /* 作业名称 (Optional) */
    JobName string `json:"jobName"`

    /* 触发器名称 (Optional) */
    TriggerName string `json:"triggerName"`

    /* 触发器组 (Optional) */
    TriggerGroup string `json:"triggerGroup"`

    /* 作业路径 (Optional) */
    ClassPath string `json:"classPath"`

    /* 定时任务类型 (Optional) */
    TaskType string `json:"taskType"`

    /* 是否已启用，1：已启用 0：未启用 (Optional) */
    Enable int `json:"enable"`

    /* 最近一次执行状态(1-成功，0-失败，2-未执行) (Optional) */
    LatestStatus int `json:"latestStatus"`

    /* 最近一次得执行信息（失败时，此字段记录报错原因） (Optional) */
    LatestMsg string `json:"latestMsg"`

    /* 定时表达式 (Optional) */
    CronExpression string `json:"cronExpression"`

    /* 创建用户ERP (Optional) */
    CreateUserErp string `json:"createUserErp"`

    /* 最近一次修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}