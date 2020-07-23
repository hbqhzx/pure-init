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

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
)

type CreateUploadTaskRequest struct {

    core.JDCloudRequest

    /* 文件名称（带后缀名）  */
    FileName string `json:"fileName"`

    /* 上传请求方式，默认：PUT (Optional) */
    HttpMethod *string `json:"httpMethod"`

    /* 视频名称  */
    MediaName string `json:"mediaName"`

    /* 文件大小 (Optional) */
    FileSize *int `json:"fileSize"`

    /* 文件MD5 (Optional) */
    Md5 *string `json:"md5"`

    /* 封面图地址 (Optional) */
    CoverImg *string `json:"coverImg"`

    /* 转码模板ID，多个时以逗号分隔 (Optional) */
    TranscodeTemplateIds *string `json:"transcodeTemplateIds"`

    /* 视频分类ID，多个时以逗号分隔 (Optional) */
    CategoryId *string `json:"categoryId"`

    /* 视频标签，多个时以逗号分隔 (Optional) */
    Tags *string `json:"tags"`

    /* 视频简介 (Optional) */
    Notes *string `json:"notes"`

    /* 上传客户端IP (Optional) */
    ClientIp *string `json:"clientIp"`
}

/*
 * param fileName: 文件名称（带后缀名） (Required)
 * param mediaName: 视频名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateUploadTaskRequest(
    fileName string,
    mediaName string,
) *CreateUploadTaskRequest {

	return &CreateUploadTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/uploadTask",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        FileName: fileName,
        MediaName: mediaName,
	}
}

/*
 * param fileName: 文件名称（带后缀名） (Required)
 * param httpMethod: 上传请求方式，默认：PUT (Optional)
 * param mediaName: 视频名称 (Required)
 * param fileSize: 文件大小 (Optional)
 * param md5: 文件MD5 (Optional)
 * param coverImg: 封面图地址 (Optional)
 * param transcodeTemplateIds: 转码模板ID，多个时以逗号分隔 (Optional)
 * param categoryId: 视频分类ID，多个时以逗号分隔 (Optional)
 * param tags: 视频标签，多个时以逗号分隔 (Optional)
 * param notes: 视频简介 (Optional)
 * param clientIp: 上传客户端IP (Optional)
 */
func NewCreateUploadTaskRequestWithAllParams(
    fileName string,
    httpMethod *string,
    mediaName string,
    fileSize *int,
    md5 *string,
    coverImg *string,
    transcodeTemplateIds *string,
    categoryId *string,
    tags *string,
    notes *string,
    clientIp *string,
) *CreateUploadTaskRequest {

    return &CreateUploadTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/uploadTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        FileName: fileName,
        HttpMethod: httpMethod,
        MediaName: mediaName,
        FileSize: fileSize,
        Md5: md5,
        CoverImg: coverImg,
        TranscodeTemplateIds: transcodeTemplateIds,
        CategoryId: categoryId,
        Tags: tags,
        Notes: notes,
        ClientIp: clientIp,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateUploadTaskRequestWithoutParam() *CreateUploadTaskRequest {

    return &CreateUploadTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/uploadTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param fileName: 文件名称（带后缀名）(Required) */
func (r *CreateUploadTaskRequest) SetFileName(fileName string) {
    r.FileName = fileName
}

/* param httpMethod: 上传请求方式，默认：PUT(Optional) */
func (r *CreateUploadTaskRequest) SetHttpMethod(httpMethod string) {
    r.HttpMethod = &httpMethod
}

/* param mediaName: 视频名称(Required) */
func (r *CreateUploadTaskRequest) SetMediaName(mediaName string) {
    r.MediaName = mediaName
}

/* param fileSize: 文件大小(Optional) */
func (r *CreateUploadTaskRequest) SetFileSize(fileSize int) {
    r.FileSize = &fileSize
}

/* param md5: 文件MD5(Optional) */
func (r *CreateUploadTaskRequest) SetMd5(md5 string) {
    r.Md5 = &md5
}

/* param coverImg: 封面图地址(Optional) */
func (r *CreateUploadTaskRequest) SetCoverImg(coverImg string) {
    r.CoverImg = &coverImg
}

/* param transcodeTemplateIds: 转码模板ID，多个时以逗号分隔(Optional) */
func (r *CreateUploadTaskRequest) SetTranscodeTemplateIds(transcodeTemplateIds string) {
    r.TranscodeTemplateIds = &transcodeTemplateIds
}

/* param categoryId: 视频分类ID，多个时以逗号分隔(Optional) */
func (r *CreateUploadTaskRequest) SetCategoryId(categoryId string) {
    r.CategoryId = &categoryId
}

/* param tags: 视频标签，多个时以逗号分隔(Optional) */
func (r *CreateUploadTaskRequest) SetTags(tags string) {
    r.Tags = &tags
}

/* param notes: 视频简介(Optional) */
func (r *CreateUploadTaskRequest) SetNotes(notes string) {
    r.Notes = &notes
}

/* param clientIp: 上传客户端IP(Optional) */
func (r *CreateUploadTaskRequest) SetClientIp(clientIp string) {
    r.ClientIp = &clientIp
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateUploadTaskRequest) GetRegionId() string {
    return ""
}

type CreateUploadTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateUploadTaskResult `json:"result"`
}

type CreateUploadTaskResult struct {
    Mid string `json:"mid"`
    UploadUrl string `json:"uploadUrl"`
}