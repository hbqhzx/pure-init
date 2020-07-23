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


type ResourceChange struct {

    /*  (Optional) */
    Action string `json:"action"`

    /* 更新操作中将对资源做的改变详情 (Optional) */
    Details []ResourceChangeDetail `json:"details"`

    /* 资源在模板中的逻辑id (Optional) */
    LogicalResourceId string `json:"logicalResourceId"`

    /* 资源物理id (Optional) */
    PhysicalResourceId string `json:"physicalResourceId"`

    /*  (Optional) */
    Replacement string `json:"replacement"`

    /* 资源类型 (Optional) */
    ResourceType string `json:"resourceType"`

    /* 更新操作中修改发生的位置 (Optional) */
    Scope []string `json:"scope"`
}
