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
    gwkms "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/gwkms/apis"
    "encoding/json"
    "errors"
)

type GwkmsClient struct {
    core.JDCloudClient
}

func NewGwkmsClient(credential *core.Credential) *GwkmsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("gwkms.jdcloud-api.com")

    return &GwkmsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "gwkms",
            Revision:    "0.0.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *GwkmsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *GwkmsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 获取公钥 */
func (c *GwkmsClient) DescribePublicKeys(request *gwkms.DescribePublicKeysRequest) (*gwkms.DescribePublicKeysResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &gwkms.DescribePublicKeysResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建token */
func (c *GwkmsClient) CreateToken(request *gwkms.CreateTokenRequest) (*gwkms.CreateTokenResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &gwkms.CreateTokenResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 刷新token */
func (c *GwkmsClient) RefreshToken(request *gwkms.RefreshTokenRequest) (*gwkms.RefreshTokenResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &gwkms.RefreshTokenResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

