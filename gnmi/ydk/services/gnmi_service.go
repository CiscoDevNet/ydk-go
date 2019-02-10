// Package services implements support for Go services.
//
// YANG Development Kit Copyright 2018 Cisco Systems. All rights reserved.
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package services

import (
//	"fmt"
//	"strconv"
//	"github.com/CiscoDevNet/ydk-go/ydk"
//	"github.com/CiscoDevNet/ydk-go/ydk/errors"
	"strings"
	"github.com/CiscoDevNet/ydk-go/ydk/path"
	"github.com/CiscoDevNet/ydk-go/ydk/providers"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
//	encoding "github.com/CiscoDevNet/ydk-go/ydk/types/encoding_format"
)

type GnmiService struct {
}

type GnmiSubscription struct {
	Entity types.Entity
	SubscriptionMode string
	SampleInterval uint64
	SuppressRedundant bool 
	HeartbeatInterval uint64
}

func (gs *GnmiService) Set(provider *providers.GnmiServiceProvider, entity types.Entity) bool {
	data := map[string]string {}

	readDataNode := path.ExecuteGnmiRPC(provider, "ydk:gnmi-set", entity, data, )
	return operationSucceeded(readDataNode)
}

func (gs *GnmiService) Get(provider *providers.GnmiServiceProvider, filter types.Entity, operation string) types.Entity {
	data := map[string]string {}
	data["mode"] = operation

	readDataNode := path.ExecuteGnmiRPC(provider, "ydk:gnmi-get", filter, data)
	return path.ReadDatanode(filter, readDataNode)
}

func (gs *GnmiService) Capabilities(provider *providers.GnmiServiceProvider) string {
	var caps string = path.GnmiServiceGetCapabilities(provider.Private)
	return caps
}

func (gs *GnmiService) Subscribe(provider *providers.GnmiServiceProvider, subscriptionList []GnmiSubscription, qos uint32, mode string, encode string) {

	session := provider.GetSession()
	schema := path.GetRootSchemaNode(provider.Private)

	rpc := path.CreateRpc( schema, "ydk:gnmi-subscribe")
	subscription := path.CreateDataNode( rpc.Input, "subscription", "")
	path.CreateDataNode( subscription, "mode", mode)
	path.CreateDataNode( subscription, "qos", qos)
	path.CreateDataNode( subscription, "encoding", encode)
	
	for _, sub := range subscriptionList {
		segpath := sub.Entity.GetEntityData().SegmentPath
		aliasKey := strings.Replace(segpath, "'", "_", -1)
		entTag  := "subscription-list[alias='" + aliasKey + "']"
		entitySub := path.CreateDataNode( subscription, entTag, "")

		payload := path.GetSubscribeDataPayload(provider, sub.Entity)
		path.CreateDataNode( entitySub, "entity", payload)

		smode := sub.SubscriptionMode
		if len(smode) == 0 {
			smode = "ON_CHANGE"
		}
		path.CreateDataNode( entitySub, "subscription-mode", smode)

		sinterval := sub.SampleInterval
		if sinterval == 0 {
			if sub.HeartbeatInterval > 0 {
				sinterval = sub.HeartbeatInterval
			} else {
				sinterval = 60000000000
			}
		}
		path.CreateDataNode( entitySub, "sample-interval", sinterval)

		sheartbeat := sub.HeartbeatInterval
		if sheartbeat == 0 {
			sheartbeat = 600000000000
		} else {
			if  sheartbeat < sinterval {
				sheartbeat = sinterval
			}
		}
		path.CreateDataNode( entitySub, "heartbeat-interval", sheartbeat)
		
		ssupress := "false"
		if sub.SuppressRedundant == true {
			ssupress = "true"
		}
		path.CreateDataNode( entitySub, "suppress-redundant", ssupress)
	}

	session.ExecuteSubscribeRpc(rpc)
}
