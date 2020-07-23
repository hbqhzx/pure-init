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


type AttackDetectionLogReportData struct {

    /* 日期 (Optional) */
    Time string `json:"time"`

    /* 各指标统计数: (误报分析:array[0]访问日志数，array[1]攻击因子数，array[2]AI处理数，array[3]误报数；漏报分析:array[0]访问日志数，array[1]检测因子，array[2]AI处理数，array[3]漏报数) (Optional) */
    Data []int `json:"data"`
}