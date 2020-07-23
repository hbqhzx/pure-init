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

package client

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
    ossopenapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ossopenapi/apis"
    "encoding/json"
    "errors"
)

type OssopenapiClient struct {
    core.JDCloudClient
}

func NewOssopenapiClient(credential *core.Credential) *OssopenapiClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("ossopenapi.jdcloud-api.com")

    return &OssopenapiClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "ossopenapi",
            Revision:    "0.5.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *OssopenapiClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *OssopenapiClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 获取下载流量情况或请求情况 */
func (c *OssopenapiClient) GetBucketMonitorStatistic(request *ossopenapi.GetBucketMonitorStatisticRequest) (*ossopenapi.GetBucketMonitorStatisticResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.GetBucketMonitorStatisticResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 自定义域名详细信息 */
func (c *OssopenapiClient) GetCName(request *ossopenapi.GetCNameRequest) (*ossopenapi.GetCNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.GetCNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除回源配置 */
func (c *OssopenapiClient) DeleteBackSourceConfiguration(request *ossopenapi.DeleteBackSourceConfigurationRequest) (*ossopenapi.DeleteBackSourceConfigurationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.DeleteBackSourceConfigurationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 关闭原图保护 */
func (c *OssopenapiClient) DeleteSourceImageProtectionConfiguration(request *ossopenapi.DeleteSourceImageProtectionConfigurationRequest) (*ossopenapi.DeleteSourceImageProtectionConfigurationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.DeleteSourceImageProtectionConfigurationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除一个bucket
 */
func (c *OssopenapiClient) DeleteBucket(request *ossopenapi.DeleteBucketRequest) (*ossopenapi.DeleteBucketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.DeleteBucketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询bucket是否存在
 */
func (c *OssopenapiClient) HeadBucket(request *ossopenapi.HeadBucketRequest) (*ossopenapi.HeadBucketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.HeadBucketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 自定义域名列表 */
func (c *OssopenapiClient) ListCNames(request *ossopenapi.ListCNamesRequest) (*ossopenapi.ListCNamesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.ListCNamesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取bucket原图保护的配置 */
func (c *OssopenapiClient) GetOriImgProtectionInfo(request *ossopenapi.GetOriImgProtectionInfoRequest) (*ossopenapi.GetOriImgProtectionInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.GetOriImgProtectionInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取回源配置 */
func (c *OssopenapiClient) GetBackSourceConfiguration(request *ossopenapi.GetBackSourceConfigurationRequest) (*ossopenapi.GetBackSourceConfigurationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.GetBackSourceConfigurationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 按bucketName删除所有域名 */
func (c *OssopenapiClient) DeleteCNamesByBucketName(request *ossopenapi.DeleteCNamesByBucketNameRequest) (*ossopenapi.DeleteCNamesByBucketNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.DeleteCNamesByBucketNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建bucket
 */
func (c *OssopenapiClient) PutBucket(request *ossopenapi.PutBucketRequest) (*ossopenapi.PutBucketResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.PutBucketResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 关闭原图保护 */
func (c *OssopenapiClient) CloseOriImgProtection(request *ossopenapi.CloseOriImgProtectionRequest) (*ossopenapi.CloseOriImgProtectionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.CloseOriImgProtectionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除自定义域名 */
func (c *OssopenapiClient) DeleteCName(request *ossopenapi.DeleteCNameRequest) (*ossopenapi.DeleteCNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.DeleteCNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取Bucket存储空间 */
func (c *OssopenapiClient) GetBucketSpaceStatistic(request *ossopenapi.GetBucketSpaceStatisticRequest) (*ossopenapi.GetBucketSpaceStatisticResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.GetBucketSpaceStatisticResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 列出当前用户的所有bucket
 */
func (c *OssopenapiClient) ListBuckets(request *ossopenapi.ListBucketsRequest) (*ossopenapi.ListBucketsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.ListBucketsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 开启原图保护 */
func (c *OssopenapiClient) CreateSourceImageProtectionConfiguration(request *ossopenapi.CreateSourceImageProtectionConfigurationRequest) (*ossopenapi.CreateSourceImageProtectionConfigurationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.CreateSourceImageProtectionConfigurationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 开启原图保护 */
func (c *OssopenapiClient) OpenOriImgProtection(request *ossopenapi.OpenOriImgProtectionRequest) (*ossopenapi.OpenOriImgProtectionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.OpenOriImgProtectionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加自定义域名 */
func (c *OssopenapiClient) AddCName(request *ossopenapi.AddCNameRequest) (*ossopenapi.AddCNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.AddCNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加修改回源配置 */
func (c *OssopenapiClient) PutBackSourceConfiguration(request *ossopenapi.PutBackSourceConfigurationRequest) (*ossopenapi.PutBackSourceConfigurationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.PutBackSourceConfigurationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 域名备案查询 */
func (c *OssopenapiClient) CheckICP(request *ossopenapi.CheckICPRequest) (*ossopenapi.CheckICPResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.CheckICPResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取bucket原图保护的配置 */
func (c *OssopenapiClient) GetSourceImageProtectionConfiguration(request *ossopenapi.GetSourceImageProtectionConfigurationRequest) (*ossopenapi.GetSourceImageProtectionConfigurationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ossopenapi.GetSourceImageProtectionConfigurationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
