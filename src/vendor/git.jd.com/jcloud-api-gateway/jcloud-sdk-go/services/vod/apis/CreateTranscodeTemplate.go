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

type CreateTranscodeTemplateRequest struct {

    core.JDCloudRequest

    /* 模板名称  */
    Name string `json:"name"`

    /* 封装格式  */
    Format string `json:"format"`

    /* 视频编码  */
    VideoCodec string `json:"videoCodec"`

    /* 分辨率-宽 (Optional) */
    Width *int `json:"width"`

    /* 分辨率-高 (Optional) */
    Height *int `json:"height"`

    /* 视频码率  */
    VideoCoderate int `json:"videoCoderate"`

    /* 视频帧率  */
    VideoFramerate string `json:"videoFramerate"`

    /* 音频编码  */
    AudioCodec string `json:"audioCodec"`

    /* 音频码率  */
    AudioCoderate int `json:"audioCoderate"`

    /* 音频采样率  */
    SampleRate int `json:"sampleRate"`

    /* 音频声道数  */
    Channel int `json:"channel"`
}

/*
 * param name: 模板名称 (Required)
 * param format: 封装格式 (Required)
 * param videoCodec: 视频编码 (Required)
 * param videoCoderate: 视频码率 (Required)
 * param videoFramerate: 视频帧率 (Required)
 * param audioCodec: 音频编码 (Required)
 * param audioCoderate: 音频码率 (Required)
 * param sampleRate: 音频采样率 (Required)
 * param channel: 音频声道数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTranscodeTemplateRequest(
    name string,
    format string,
    videoCodec string,
    videoCoderate int,
    videoFramerate string,
    audioCodec string,
    audioCoderate int,
    sampleRate int,
    channel int,
) *CreateTranscodeTemplateRequest {

	return &CreateTranscodeTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodes",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Name: name,
        Format: format,
        VideoCodec: videoCodec,
        VideoCoderate: videoCoderate,
        VideoFramerate: videoFramerate,
        AudioCodec: audioCodec,
        AudioCoderate: audioCoderate,
        SampleRate: sampleRate,
        Channel: channel,
	}
}

/*
 * param name: 模板名称 (Required)
 * param format: 封装格式 (Required)
 * param videoCodec: 视频编码 (Required)
 * param width: 分辨率-宽 (Optional)
 * param height: 分辨率-高 (Optional)
 * param videoCoderate: 视频码率 (Required)
 * param videoFramerate: 视频帧率 (Required)
 * param audioCodec: 音频编码 (Required)
 * param audioCoderate: 音频码率 (Required)
 * param sampleRate: 音频采样率 (Required)
 * param channel: 音频声道数 (Required)
 */
func NewCreateTranscodeTemplateRequestWithAllParams(
    name string,
    format string,
    videoCodec string,
    width *int,
    height *int,
    videoCoderate int,
    videoFramerate string,
    audioCodec string,
    audioCoderate int,
    sampleRate int,
    channel int,
) *CreateTranscodeTemplateRequest {

    return &CreateTranscodeTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodes",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Name: name,
        Format: format,
        VideoCodec: videoCodec,
        Width: width,
        Height: height,
        VideoCoderate: videoCoderate,
        VideoFramerate: videoFramerate,
        AudioCodec: audioCodec,
        AudioCoderate: audioCoderate,
        SampleRate: sampleRate,
        Channel: channel,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTranscodeTemplateRequestWithoutParam() *CreateTranscodeTemplateRequest {

    return &CreateTranscodeTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodes",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param name: 模板名称(Required) */
func (r *CreateTranscodeTemplateRequest) SetName(name string) {
    r.Name = name
}

/* param format: 封装格式(Required) */
func (r *CreateTranscodeTemplateRequest) SetFormat(format string) {
    r.Format = format
}

/* param videoCodec: 视频编码(Required) */
func (r *CreateTranscodeTemplateRequest) SetVideoCodec(videoCodec string) {
    r.VideoCodec = videoCodec
}

/* param width: 分辨率-宽(Optional) */
func (r *CreateTranscodeTemplateRequest) SetWidth(width int) {
    r.Width = &width
}

/* param height: 分辨率-高(Optional) */
func (r *CreateTranscodeTemplateRequest) SetHeight(height int) {
    r.Height = &height
}

/* param videoCoderate: 视频码率(Required) */
func (r *CreateTranscodeTemplateRequest) SetVideoCoderate(videoCoderate int) {
    r.VideoCoderate = videoCoderate
}

/* param videoFramerate: 视频帧率(Required) */
func (r *CreateTranscodeTemplateRequest) SetVideoFramerate(videoFramerate string) {
    r.VideoFramerate = videoFramerate
}

/* param audioCodec: 音频编码(Required) */
func (r *CreateTranscodeTemplateRequest) SetAudioCodec(audioCodec string) {
    r.AudioCodec = audioCodec
}

/* param audioCoderate: 音频码率(Required) */
func (r *CreateTranscodeTemplateRequest) SetAudioCoderate(audioCoderate int) {
    r.AudioCoderate = audioCoderate
}

/* param sampleRate: 音频采样率(Required) */
func (r *CreateTranscodeTemplateRequest) SetSampleRate(sampleRate int) {
    r.SampleRate = sampleRate
}

/* param channel: 音频声道数(Required) */
func (r *CreateTranscodeTemplateRequest) SetChannel(channel int) {
    r.Channel = channel
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTranscodeTemplateRequest) GetRegionId() string {
    return ""
}

type CreateTranscodeTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTranscodeTemplateResult `json:"result"`
}

type CreateTranscodeTemplateResult struct {
    CoderateId int64 `json:"coderateId"`
}