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


type Operation struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 名称 (Optional) */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 图片地址 (Optional) */
    PicUrl string `json:"picUrl"`

    /* icon类 (Optional) */
    IconClass string `json:"iconClass"`

    /* 链接地址 (Optional) */
    LinkUrl string `json:"linkUrl"`

    /* 排序 (Optional) */
    Sort int `json:"sort"`

    /* 0 底部运营位，1 运营位 (Optional) */
    Position int `json:"position"`

    /* 操作者 (Optional) */
    Operator string `json:"operator"`

    /* createTime (Optional) */
    CreateTime int `json:"createTime"`

    /* updateTime (Optional) */
    UpdateTime int `json:"updateTime"`

    /* content (Optional) */
    Content string `json:"content"`

    /* 存储文本内容的key (Optional) */
    ContentKey string `json:"contentKey"`

    /* jss保存文本的地址 (Optional) */
    ContentUrl string `json:"contentUrl"`

    /* banner编辑的类型 (Optional) */
    OperationType int `json:"operationType"`

    /* 运营位埋点 (Optional) */
    Clstag string `json:"clstag"`

    /* 运营位标签 (Optional) */
    Label string `json:"label"`
}