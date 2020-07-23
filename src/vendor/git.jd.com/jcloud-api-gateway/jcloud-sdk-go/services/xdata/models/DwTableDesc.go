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


type DwTableDesc struct {

    /* 数据库名 (Optional) */
    DbName *string `json:"dbName"`

    /* 表名 (Optional) */
    TableName *string `json:"tableName"`

    /* 存储格式 (Optional) */
    HiveFileFormat *string `json:"hiveFileFormat"`

    /* 字段分隔符 (Optional) */
    FieldsDelimit *string `json:"fieldsDelimit"`

    /* 行分隔符 (Optional) */
    LinesDelimit *string `json:"linesDelimit"`

    /* 其他serde属性 (Optional) */
    OtherSerdeProperties *interface{} `json:"otherSerdeProperties"`

    /* 创建时间（自动生成） (Optional) */
    CreateTime *string `json:"createTime"`

    /* 所有者（自动生成） (Optional) */
    Owner *string `json:"owner"`

    /* 描述信息 (Optional) */
    Comments *string `json:"comments"`

    /* 外表位置 (Optional) */
    ExternalLocation *string `json:"externalLocation"`

    /* 参数 (Optional) */
    Parameters *interface{} `json:"parameters"`

    /* 列信息 (Optional) */
    Rows []DwTableRow `json:"rows"`
}
