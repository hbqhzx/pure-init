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


type SnapshotTemplate struct {

    /* 截图格式:
  - 取值: jpg, png
  - 不区分大小写
 (Optional) */
    Format string `json:"format"`

    /* 截图宽度:
  - 取值: [8,8192]
  - 等比: 如果只填写一个参数,则按参数比例等比缩放截图
  - 随源: 如果两个参数都不填写，则截取源流大小原图
 (Optional) */
    Width int `json:"width"`

    /* 截图高度:
  - 取值: [8,8192]
  - 等比: 如果只填写一个参数,则按参数比例等比缩放截图
  - 随源: 如果两个参数都不填写，则截取源流大小原图
 (Optional) */
    Height int `json:"height"`

    /* 截图与设定的宽高不匹配时的处理规则:
  - 1-拉伸
  - 2-留黑
  - 3-留白
  - 4-高斯模糊
  - 默认值1,2,3,4是等比例的缩放，1是按照设定宽高拉伸
 (Optional) */
    FillType int `json:"fillType"`

    /* 截图周期:
  - MIN_INTEGER = 5
  - MAX_INTEGER = 3600;
  - 单位: 秒
 (Optional) */
    SnapshotInterval int `json:"snapshotInterval"`

    /* 存储模式:
  - 1-覆盖
  - 2-顺序编号存储
 (Optional) */
    SaveMode int `json:"saveMode"`

    /* 存储桶 (Optional) */
    SaveBucket string `json:"saveBucket"`

    /* 存储地址 (Optional) */
    SaveEndpoint string `json:"saveEndpoint"`

    /* 截图模板自定义名称:
  - 取值要求：数字、大小写字母或短横线("-"),
              首尾不能有特殊字符("-")
  - <b>注意: 不能与已定义命名重复</b>
 (Optional) */
    Template string `json:"template"`
}