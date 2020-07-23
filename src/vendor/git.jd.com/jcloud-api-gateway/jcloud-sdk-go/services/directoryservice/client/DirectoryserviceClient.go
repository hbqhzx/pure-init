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
    directoryservice "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/directoryservice/apis"
    "encoding/json"
    "errors"
)

type DirectoryserviceClient struct {
    core.JDCloudClient
}

func NewDirectoryserviceClient(credential *core.Credential) *DirectoryserviceClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("directoryservice.jdcloud-api.com")

    return &DirectoryserviceClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "directoryservice",
            Revision:    "1.0.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *DirectoryserviceClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *DirectoryserviceClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 修改域控制器数目 */
func (c *DirectoryserviceClient) UpdateNumberOfDomainControllers(request *directoryservice.UpdateNumberOfDomainControllersRequest) (*directoryservice.UpdateNumberOfDomainControllersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.UpdateNumberOfDomainControllersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用应用程序和服务url */
func (c *DirectoryserviceClient) EnableApplicationServiceUrl(request *directoryservice.EnableApplicationServiceUrlRequest) (*directoryservice.EnableApplicationServiceUrlResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.EnableApplicationServiceUrlResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 设置/启用目录服务url */
func (c *DirectoryserviceClient) EnableDirectoryUrl(request *directoryservice.EnableDirectoryUrlRequest) (*directoryservice.EnableDirectoryUrlResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.EnableDirectoryUrlResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取AD资源价格 */
func (c *DirectoryserviceClient) DescribePrice(request *directoryservice.DescribePriceRequest) (*directoryservice.DescribePriceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribePriceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加本林或信任林上的用户或组 */
func (c *DirectoryserviceClient) AddUserOrGroup(request *directoryservice.AddUserOrGroupRequest) (*directoryservice.AddUserOrGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.AddUserOrGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 检查与远程域是否已经建立信任关系 */
func (c *DirectoryserviceClient) CheckRemoteDomain(request *directoryservice.CheckRemoteDomainRequest) (*directoryservice.CheckRemoteDomainResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.CheckRemoteDomainResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加/创建信任关系 */
func (c *DirectoryserviceClient) CreateTrust(request *directoryservice.CreateTrustRequest) (*directoryservice.CreateTrustResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.CreateTrustResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取应用程序和服务列表 */
func (c *DirectoryserviceClient) DescribeApplicationServices(request *directoryservice.DescribeApplicationServicesRequest) (*directoryservice.DescribeApplicationServicesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeApplicationServicesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建目录服务 */
func (c *DirectoryserviceClient) CreateDirectory(request *directoryservice.CreateDirectoryRequest) (*directoryservice.CreateDirectoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.CreateDirectoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 禁用目录服务url */
func (c *DirectoryserviceClient) DisableDirectoryUrl(request *directoryservice.DisableDirectoryUrlRequest) (*directoryservice.DisableDirectoryUrlResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DisableDirectoryUrlResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 禁用应用程序和服务 */
func (c *DirectoryserviceClient) DisableApplicationService(request *directoryservice.DisableApplicationServiceRequest) (*directoryservice.DisableApplicationServiceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DisableApplicationServiceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 子网列表 */
func (c *DirectoryserviceClient) DescribeSubnets(request *directoryservice.DescribeSubnetsRequest) (*directoryservice.DescribeSubnetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeSubnetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 验证信任关系 */
func (c *DirectoryserviceClient) VerifyTrust(request *directoryservice.VerifyTrustRequest) (*directoryservice.VerifyTrustResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.VerifyTrustResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 目录服务详细信息 */
func (c *DirectoryserviceClient) DescribeDirectory(request *directoryservice.DescribeDirectoryRequest) (*directoryservice.DescribeDirectoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeDirectoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 角色列表 */
func (c *DirectoryserviceClient) DescribeRoles(request *directoryservice.DescribeRolesRequest) (*directoryservice.DescribeRolesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeRolesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 目录服务类型 托管式目录服务 connector等 */
func (c *DirectoryserviceClient) DescribeDirectoryType(request *directoryservice.DescribeDirectoryTypeRequest) (*directoryservice.DescribeDirectoryTypeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeDirectoryTypeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询信任关系列表 */
func (c *DirectoryserviceClient) DescribeTrusts(request *directoryservice.DescribeTrustsRequest) (*directoryservice.DescribeTrustsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeTrustsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除用户或组 */
func (c *DirectoryserviceClient) DeleteUserOrGroup(request *directoryservice.DeleteUserOrGroupRequest) (*directoryservice.DeleteUserOrGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DeleteUserOrGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 禁用应用程序和服务url */
func (c *DirectoryserviceClient) DisableApplicationServiceUrl(request *directoryservice.DisableApplicationServiceUrlRequest) (*directoryservice.DisableApplicationServiceUrlResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DisableApplicationServiceUrlResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* AD实例列表 */
func (c *DirectoryserviceClient) DescribeDirectories(request *directoryservice.DescribeDirectoriesRequest) (*directoryservice.DescribeDirectoriesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeDirectoriesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取目录服务的区域和可用区 */
func (c *DirectoryserviceClient) DescribeRegion(request *directoryservice.DescribeRegionRequest) (*directoryservice.DescribeRegionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeRegionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用应用程序和服务 */
func (c *DirectoryserviceClient) EnableApplicationService(request *directoryservice.EnableApplicationServiceRequest) (*directoryservice.EnableApplicationServiceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.EnableApplicationServiceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除信任关系 */
func (c *DirectoryserviceClient) DeleteTrust(request *directoryservice.DeleteTrustRequest) (*directoryservice.DeleteTrustResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DeleteTrustResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* VPC列表 */
func (c *DirectoryserviceClient) DescribeVpcs(request *directoryservice.DescribeVpcsRequest) (*directoryservice.DescribeVpcsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeVpcsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除目录服务 */
func (c *DirectoryserviceClient) DeleteDirectory(request *directoryservice.DeleteDirectoryRequest) (*directoryservice.DeleteDirectoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DeleteDirectoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 搜索本林或信任林上的用户或组 */
func (c *DirectoryserviceClient) SearchDomainUserOrGroup(request *directoryservice.SearchDomainUserOrGroupRequest) (*directoryservice.SearchDomainUserOrGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.SearchDomainUserOrGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 域控制器列表 */
func (c *DirectoryserviceClient) DescribeDomainControllers(request *directoryservice.DescribeDomainControllersRequest) (*directoryservice.DescribeDomainControllersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeDomainControllersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 目录服务版本 标准版 企业版等 */
func (c *DirectoryserviceClient) DescribeDirectoryVersion(request *directoryservice.DescribeDirectoryVersionRequest) (*directoryservice.DescribeDirectoryVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeDirectoryVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 用户与组列表 */
func (c *DirectoryserviceClient) DescribeUsersAndGroups(request *directoryservice.DescribeUsersAndGroupsRequest) (*directoryservice.DescribeUsersAndGroupsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeUsersAndGroupsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 检查dns是否可用 */
func (c *DirectoryserviceClient) VerifyDNS(request *directoryservice.VerifyDNSRequest) (*directoryservice.VerifyDNSResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.VerifyDNSResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 本林和信任林列表 */
func (c *DirectoryserviceClient) DescribeTrustDomains(request *directoryservice.DescribeTrustDomainsRequest) (*directoryservice.DescribeTrustDomainsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &directoryservice.DescribeTrustDomainsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

