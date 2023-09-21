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

package config

import (
	"github.com/polaris-contrib/apiserver-nacos/core"
	"github.com/polaris-contrib/apiserver-nacos/v2/remote"
	"github.com/polarismesh/polaris/auth"
	"github.com/polarismesh/polaris/config"
	"github.com/polarismesh/polaris/namespace"
)

type ServerOption struct {
	// nacos
	ConnectionManager *remote.ConnectionManager
	Store             *core.NacosDataStorage

	// polaris
	UserSvr         auth.UserServer
	CheckerSvr      auth.StrategyServer
	NamespaceSvr    namespace.NamespaceOperateServer
	ConfigSvr       config.ConfigCenterServer
	OriginConfigSvr config.ConfigCenterServer
}

type ConfigServer struct {
}

func (h *ConfigServer) ListGRPCHandlers() map[string]*remote.RequestHandlerWarrper {
	return nil
}
