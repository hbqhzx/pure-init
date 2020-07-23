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


type PbListQueueInfosInput struct {

    /* UUID.V4 for each request (Optional) */
    RequestId *string `json:"requestId"`

    /* region name, etc. cn-north-1 (Optional) */
    RegionId *string `json:"regionId"`

    /* senderId是查询queue所属用户的Id
 (Optional) */
    SenderId *string `json:"senderId"`

    /* queueNames的前缀，支持前缀搜索
 (Optional) */
    KeyWord *string `json:"keyWord"`

    /* queue's type, support: fifo,standard,all
 (Optional) */
    QueueType *string `json:"queueType"`
}
