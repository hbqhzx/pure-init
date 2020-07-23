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


type ModifyDeploy struct {

    /* 描述 (Optional) */
    Desc string `json:"desc"`

    /* 部署来源：1url，2云编译，3云存储 (Optional) */
    DeploySource int `json:"deploySource"`

    /* 部署操作 (Optional) */
    DeployCmd string `json:"deployCmd"`

    /* 1使用输入的操作，2使用程序自带操作 (Optional) */
    CmdSource int `json:"cmdSource"`

    /* 下载url (Optional) */
    DownloadUrl string `json:"downloadUrl"`

    /* md5 (Optional) */
    Md5 string `json:"md5"`

    /* 云编译项目名 (Optional) */
    CompileProject string `json:"compileProject"`

    /* 云编译构建序号 (Optional) */
    CompileSeries string `json:"compileSeries"`

    /* 云存储空间 (Optional) */
    OssSpace string `json:"ossSpace"`

    /* 云存储目录 (Optional) */
    OssDir string `json:"ossDir"`

    /* 文件类型：1.tar，2.zip,3.tar.gz (Optional) */
    FileType string `json:"fileType"`
}
