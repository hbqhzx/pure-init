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


type Catalog struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 栏目名称 (Optional) */
    Name string `json:"name"`

    /* 访问路径 (Optional) */
    Path string `json:"path"`

    /* 父栏目ID (Optional) */
    ParentId int `json:"parentId"`

    /* ICON 图标 (Optional) */
    IconUrl string `json:"iconUrl"`

    /* ICON 样式名 (Optional) */
    IconClass string `json:"iconClass"`

    /* 埋点名 (Optional) */
    Clstag string `json:"clstag"`

    /* 栏目简介 (Optional) */
    Summary string `json:"summary"`

    /* 索引页模板 (Optional) */
    IndexTemplate string `json:"indexTemplate"`

    /* 列表页模板 (Optional) */
    ListTemplate string `json:"listTemplate"`

    /* 详情页模板 (Optional) */
    DetailTemplate string `json:"detailTemplate"`

    /* 层级 (Optional) */
    TreeLevel int `json:"treeLevel"`

    /* 是否叶节点 0 否 1是 (Optional) */
    IsLeaf int `json:"isLeaf"`

    /* 排序值 (Optional) */
    Sort int `json:"sort"`

    /* 是否有文章 0 否 1是 (Optional) */
    HasArticle int `json:"hasArticle"`

    /* 操作人员 (Optional) */
    Operator string `json:"operator"`

    /* createTime (Optional) */
    CreateTime string `json:"createTime"`

    /* modifyTime (Optional) */
    ModifyTime string `json:"modifyTime"`

    /* 状态 0 显示 1 隐藏 (Optional) */
    Status int `json:"status"`

    /* children (Optional) */
    Children []Catalog `json:"children"`

    /* firstCatalogId (Optional) */
    FirstCatalogId int `json:"firstCatalogId"`

    /* firstCatalogName (Optional) */
    FirstCatalogName string `json:"firstCatalogName"`

    /* pdf保存地址 (Optional) */
    PdfUrl string `json:"pdfUrl"`

    /* pdf文件名称 (Optional) */
    PdfName string `json:"pdfName"`

    /* 标签 (Optional) */
    Label string `json:"label"`
}