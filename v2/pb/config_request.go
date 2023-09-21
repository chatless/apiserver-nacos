/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package nacos_grpc_service

import "encoding/json"

type ConfigItem struct {
	Id      json.Number `param:"id"`
	DataId  string      `param:"dataId"`
	Group   string      `param:"group"`
	Content string      `param:"content"`
	Md5     string      `param:"md5"`
	Tenant  string      `param:"tenant"`
	Appname string      `param:"appname"`
}

type ConfigPage struct {
	TotalCount     int          `param:"totalCount"`
	PageNumber     int          `param:"pageNumber"`
	PagesAvailable int          `param:"pagesAvailable"`
	PageItems      []ConfigItem `param:"pageItems"`
}

type ConfigListenContext struct {
	Group  string `json:"group"`
	Md5    string `json:"md5"`
	DataId string `json:"dataId"`
	Tenant string `json:"tenant"`
}

type ConfigContext struct {
	Group  string `json:"group"`
	DataId string `json:"dataId"`
	Tenant string `json:"tenant"`
}

type ConfigRequest struct {
	*Request
	Group  string `json:"group"`
	DataId string `json:"dataId"`
	Tenant string `json:"tenant"`
	Module string `json:"module"`
}

func (c *ConfigRequest) RequestMeta() interface{} {
	return c
}

func NewConfigRequest(group, dataId, tenant string) *ConfigRequest {
	request := Request{
		Headers: make(map[string]string, 8),
	}
	return &ConfigRequest{
		Request: &request,
		Group:   group,
		DataId:  dataId,
		Tenant:  tenant,
		Module:  "config",
	}
}

func (r *ConfigRequest) GetDataId() string {
	return r.DataId
}

func (r *ConfigRequest) GetGroup() string {
	return r.Group
}

func (r *ConfigRequest) GetTenant() string {
	return r.Tenant
}

// request of listening a batch of configs.
type ConfigBatchListenRequest struct {
	*ConfigRequest
	Listen               bool                  `json:"listen"`
	ConfigListenContexts []ConfigListenContext `json:"configListenContexts"`
}

func (c *ConfigBatchListenRequest) RequestMeta() interface{} {
	return c
}

func NewConfigBatchListenRequest(cacheLen int) *ConfigBatchListenRequest {
	return &ConfigBatchListenRequest{
		Listen:               true,
		ConfigListenContexts: make([]ConfigListenContext, 0, cacheLen),
		ConfigRequest:        NewConfigRequest("", "", ""),
	}
}

func (r *ConfigBatchListenRequest) GetRequestType() string {
	return "ConfigBatchListenRequest"
}

type ConfigChangeNotifyRequest struct {
	*ConfigRequest
}

func NewConfigChangeNotifyRequest(group, dataId, tenant string) *ConfigChangeNotifyRequest {
	return &ConfigChangeNotifyRequest{ConfigRequest: NewConfigRequest(group, dataId, tenant)}
}

func (r *ConfigChangeNotifyRequest) GetRequestType() string {
	return "ConfigChangeNotifyRequest"
}

type ConfigQueryRequest struct {
	*ConfigRequest
	Tag string `json:"tag"`
}

func (c *ConfigQueryRequest) RequestMeta() interface{} {
	return c
}

func NewConfigQueryRequest(group, dataId, tenant string) *ConfigQueryRequest {
	return &ConfigQueryRequest{ConfigRequest: NewConfigRequest(group, dataId, tenant)}
}

func (r *ConfigQueryRequest) GetRequestType() string {
	return "ConfigQueryRequest"
}

type ConfigPublishRequest struct {
	*ConfigRequest
	Content     string            `json:"content"`
	CasMd5      string            `json:"casMd5"`
	AdditionMap map[string]string `json:"additionMap"`
}

func (c *ConfigPublishRequest) RequestMeta() interface{} {
	return c
}

func NewConfigPublishRequest(group, dataId, tenant, content, casMd5 string) *ConfigPublishRequest {
	return &ConfigPublishRequest{ConfigRequest: NewConfigRequest(group, dataId, tenant),
		Content: content, CasMd5: casMd5, AdditionMap: make(map[string]string)}
}

func (r *ConfigPublishRequest) GetRequestType() string {
	return "ConfigPublishRequest"
}

type ConfigRemoveRequest struct {
	*ConfigRequest
}

func (c *ConfigRemoveRequest) RequestMeta() interface{} {
	return c
}

func NewConfigRemoveRequest(group, dataId, tenant string) *ConfigRemoveRequest {
	return &ConfigRemoveRequest{ConfigRequest: NewConfigRequest(group, dataId, tenant)}
}

func (r *ConfigRemoveRequest) GetRequestType() string {
	return "ConfigRemoveRequest"
}
