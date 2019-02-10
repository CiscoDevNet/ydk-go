// Package providers implements support for Go service providers.
//
// YANG Development Kit Copyright 2017 Cisco Systems. All rights reserved.
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
package path

// #cgo linux LDFLAGS:  -lydk_gnmi -lgrpc++ -lprotobuf
// #include <ydk/ydk.h>
// #include <ydk/ydk_gnmi.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"strings"
	"github.com/CiscoDevNet/ydk-go/ydk"
	"github.com/CiscoDevNet/ydk-go/ydk/errors"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
	"github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
	"unsafe"
)

// GnmiSession declaration and methods
//
type GnmiSession struct {
	Repo        types.Repository
	Address     string
	Username    string
	Password    string
	Port        int
	ServerCert  string
	PrivateKey  string

	Private		interface{}
	State		errors.State
}

// Connect to gNMI Session
func (gs *GnmiSession) Connect() {

	var caddress *C.char = C.CString(gs.Address)
	defer C.free(unsafe.Pointer(caddress))
	if gs.Port == 0 {
		gs.Port = 57400
	}
	var cport C.int = C.int(gs.Port)

	var cserver *C.char = C.CString(gs.ServerCert)
	defer C.free(unsafe.Pointer(cserver))

	var cclient *C.char = C.CString(gs.PrivateKey)
	defer C.free(unsafe.Pointer(cclient))

	AddCState(&gs.State)
	cstate := GetCState(&gs.State)
	
	var repopath *C.char = C.CString(gs.Repo.Path)
	defer C.free(unsafe.Pointer(repopath))
	repo := C.RepositoryInitWithPath( *cstate, repopath)
	PanicOnCStateError(cstate)

	var cusername *C.char = C.CString(gs.Username)
	defer C.free(unsafe.Pointer(cusername))
	var cpassword *C.char = C.CString(gs.Password)
	defer C.free(unsafe.Pointer(cpassword))
	
	gs.Private = C.GnmiSessionInit( *cstate, repo, caddress, cport, cusername, cpassword, cserver, cclient);
	PanicOnCStateError(cstate)
}

// Disconnect from gNMI Session
func (gs *GnmiSession) Disconnect() {
	if gs.Private == nil {
		return
	}
	realSession := gs.Private.(C.GnmiSession)
	C.GnmiSessionFree(realSession)
	CleanUpErrorState(&gs.State)
}

func (gs *GnmiSession) GetRootSchemaNode() types.RootSchemaNode {
	cstate := GetCState(&gs.State)
	
	realSession := gs.Private.(C.GnmiSession)

	var rootSchema C.RootSchemaWrapper = C.GnmiSessionGetRootSchemaNode( *cstate, realSession)
	PanicOnCStateError(cstate)
	//ydkpath.PanicOnStateError(gs.State)
	if rootSchema == nil {
        ydk.YLogError("Root schema is nil!")
		panic(1)
	}

	rsn := types.RootSchemaNode{Private: rootSchema}
	return rsn
}

func (gs *GnmiSession) ExecuteRpc(rpc types.Rpc) types.DataNode {
	cstate := GetCState(&gs.State)

	csession := gs.Private.(C.GnmiSession)
	crpc := rpc.Private.(C.Rpc)

	cdn := C.GnmiSessionExecuteRpc( *cstate, csession, crpc)
	PanicOnCStateError(cstate)
	
	dn := types.DataNode{Private: cdn}
	return dn
}

func (gs *GnmiSession) ExecuteSubscribeRpc(rpc types.Rpc) {
	cstate := GetCState(&gs.State)

	csession := gs.Private.(C.GnmiSession)
	crpc := rpc.Private.(C.Rpc)

	C.GnmiSessionExecuteSubscribeRpc( *cstate, csession, crpc)
	PanicOnCStateError(cstate)
}

// Utility functions
//
func (gs *GnmiSession) SubscribeInProgress() bool {
	realSession := gs.Private.(C.GnmiSession)
	cstate := GetCState(&gs.State)
	
	var cresponse C.boolean = C.GnmiSessionSubscribeInProgress( *cstate, realSession)
	PanicOnCStateError(cstate)
	var response bool = false
	if cresponse == 1 {
		response = true;
	}
	return response
}

func (gs *GnmiSession) GetLastSubscribeResponse(previousResponse string) string {
	realSession := gs.Private.(C.GnmiSession)
	cstate := GetCState(&gs.State)
	
	var cprevious *C.char = C.CString(previousResponse)
	defer C.free(unsafe.Pointer(cprevious))
	
	cresponse := C.GetLastSubscribeResponse( *cstate, realSession, cprevious)
	return C.GoString(cresponse)
}

// gNMI Service Provider C-wrapper functions
//
func GnmiServiceProviderConnect(
	state *errors.State,
	repo types.Repository,
	address string, port int,
	username, password string,
	serverCert, privateKey string) types.CServiceProvider {

	AddCState(state)
	cstate := GetCState(state)

	var repopath *C.char = C.CString(repo.Path)
	crepo := C.RepositoryInitWithPath(*cstate, repopath)
	PanicOnCStateError(cstate)

	var caddress *C.char = C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	
	var cport C.int = C.int(port)

	var p C.ServiceProvider
	var cusername *C.char = C.CString(username)
	defer C.free(unsafe.Pointer(cusername))
	var cpassword *C.char = C.CString(password)
	defer C.free(unsafe.Pointer(cpassword))

	var cserver *C.char = C.CString(serverCert)
	defer C.free(unsafe.Pointer(cserver))
	var cclient *C.char = C.CString(privateKey)
	defer C.free(unsafe.Pointer(cclient))
	
	p = C.GnmiServiceProviderInit( *cstate, crepo, caddress, cport, cusername, cpassword, cserver, cclient);
	PanicOnCStateError(cstate)

	cprovider := types.CServiceProvider{Private: p}
	return cprovider
}

func GnmiServiceProviderDisconnect(provider types.CServiceProvider) {
	realProvider := provider.Private.(C.ServiceProvider)
	C.GnmiServiceProviderFree(realProvider)
}

func GnmiServiceProviderGetSession(provider types.CServiceProvider) *GnmiSession {

	var state errors.State
	AddCState(&state)
	cstate := GetCState(&state)

	realProvider := provider.Private.(C.ServiceProvider)
	realSession := C.GnmiServiceProviderGetSession( *cstate, realProvider)
	PanicOnCStateError(cstate)
	
	gnmiSession := GnmiSession{Private: realSession, State: state}
	// TODO. Possibly need to populate the rest of structure members

	return &gnmiSession
}

// GnmiService C-wrapper functions
//
func GnmiServiceGetCapabilities(provider types.CServiceProvider) string {
	realProvider := provider.Private.(C.ServiceProvider)
	var state errors.State
	AddCState(&state)
	cstate := GetCState(&state)

	var ccaps *C.char = C.GnmiServiceGetCapabilities( *cstate, realProvider)
	PanicOnCStateError(cstate)

	return C.GoString(ccaps)
}

func GnmiServiceGet(provider types.CServiceProvider, filter types.Entity, operation string) types.DataNode {
	realProvider := provider.Private.(C.ServiceProvider)
	var state errors.State
	AddCState(&state)
	cstate := GetCState(&state)

	var cmode *C.char = C.CString(operation)
	defer C.free(unsafe.Pointer(cmode))

	ec := types.EntityToCollection(filter)	
	pathList := C.GnmiPathListInit()
	for _, ent := range ec.Entities() {
		path := parseEntityToPath(ent)
		C.GnmiPathListAdd(pathList, path)
	}
	cdn := C.GnmiServiceGetFromPath( *cstate, realProvider, pathList, cmode)
	PanicOnCStateError(cstate)
	C.GnmiPathListFree(pathList)

	return types.DataNode{cdn}
}

// ExecuteRPC executes payload converted from entity for CRUD/Netconf services.
// Returns a data node (types.DataNode) representing the result of the executed rpc.
func ExecuteGnmiRPC(
	provider types.ServiceProvider,
	rpcTag string,
	entity types.Entity,
	options map[string]string) types.DataNode {

	state := provider.GetState()
	cstate := GetCState(state)
	wrappedProvider := provider.GetPrivate().(types.CServiceProvider)
	realProvider := wrappedProvider.Private.(C.ServiceProvider)
	rootSchema := C.ServiceProviderGetRootSchema(*cstate, realProvider)
	PanicOnCStateError(cstate)

	ydkRPC := C.RootSchemaNodeRpc(*cstate, rootSchema, C.CString(rpcTag))
	PanicOnCStateError(cstate)
	if rootSchema == nil {
		ydk.YLogError("ExecuteGnmiRPC: Root schema is nil!")
		panic(1)
	}

	input := C.RpcInput(*cstate, ydkRPC)
	PanicOnCStateError(cstate)

	if rpcTag == "ydk:gnmi-get" {
		mode, ok := options["mode"]
		if !ok {
			mode = "ALL"
		}
		C.DataNodeCreate(*cstate, input, C.CString("type"), C.CString(mode))
		PanicOnCStateError(cstate)
	}

	// Build RPC
	var collectionFilter yfilter.YFilter = yfilter.NotSet
	config := types.EntityToCollection(entity)
	for _, ent := range config.Entities() {
		entityData := ent.GetEntityData()
		if entityData == nil {
			continue
		}
		entityFilter := collectionFilter
		if entityData.YFilter != yfilter.NotSet {
			entityFilter = entityData.YFilter
		}

		segmentPath := entityData.SegmentPath
		aliasKey := strings.Replace(segmentPath, "'", "_", -1)
		var entityTag string
		switch entityFilter {
			case yfilter.Replace:
				entityTag = "replace[alias='" + aliasKey + "']/entity"
			case yfilter.Update:
				entityTag = "update[alias='" + aliasKey + "']/entity"
			case yfilter.Delete:
				entityTag = "delete[alias='" + aliasKey + "']/entity"
			default:
				entityTag = "request[alias='" + aliasKey + "']/entity"
		}
		if entityData.YFilter != yfilter.NotSet {
			// Remove setting of the filter before payload is calculated
			types.SetEntityFilter(ent, yfilter.NotSet)
		}
		cpayload := GetDataPayload(state, ent, rootSchema, provider)
		defer C.free(unsafe.Pointer(cpayload))
		C.DataNodeCreate(*cstate, input, C.CString(entityTag), cpayload)
		PanicOnCStateError(cstate)
	}

	// Send RPC and receive results
	dataNode := types.DataNode{C.RpcExecute(*cstate, ydkRPC, realProvider)}
	PanicOnCStateError(cstate)

	return dataNode
}

func GetSubscribeDataPayload(provider types.ServiceProvider, entity types.Entity) string {

	state := provider.GetState()
	cstate := GetCState(state)
	wrappedProvider := provider.GetPrivate().(types.CServiceProvider)
	realProvider := wrappedProvider.Private.(C.ServiceProvider)
	rootSchema := C.ServiceProviderGetRootSchema(*cstate, realProvider)
	PanicOnCStateError(cstate)

	cpayload := GetDataPayload(state, entity, rootSchema, provider)
	defer C.free(unsafe.Pointer(cpayload))
	PanicOnCStateError(cstate)

	return C.GoString(cpayload)
}

// GNMI Path utility parses entity tree and converts it to gnmi::Path
//
func parseEntityToPath(entity types.Entity) C.GnmiPath {
	path := C.GnmiPathInit()
	parseEntityPrefix(entity, path)
	entityData := entity.GetEntityData()
	if entityData.AbsolutePath == entityData.SegmentPath {
		// This is top level class
		parseEntityChildren(entity, path);
	} else {
		parseEntity(entity, path);
	}
	return path
}

func addPathElem(s string, path C.GnmiPath) C.GnmiPathElem {
	ydk.YLogDebug(fmt.Sprintf("addPathElem: Adding elem: '%s'", s))
	var pathElem *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(pathElem))
	elem := C.GnmiPathAddElem(path, pathElem)
	return elem
}

func parseEntityPrefix(entity types.Entity, path C.GnmiPath) {
	// Add origin and following containers to the path
	absPath := entity.GetEntityData().AbsolutePath
	segments := strings.Split(absPath, "/")
	s := strings.Split(segments[0], ":")
	if len(s) >= 2 {
		ydk.YLogDebug(fmt.Sprintf("parseEntityPrefix: Adding origin: '%s'", s[0]))
		var origin *C.char = C.CString(s[0])
		defer C.free(unsafe.Pointer(origin))
		C.GnmiPathAddOrigin(path, origin)
		addPathElem(s[1], path)
	} else {
		ydk.YLogWarn(fmt.Sprintf("parseEntityPrefix: Missing module in the root entity path: '%s'", s[0]))
		addPathElem(s[0], path)
	}
	// Add following segments to the path
	// The last segment belong to the entity, therefore do not process it
	for i := 1; i < len(segments)-1; i++ {
		addPathElem(segments[i], path)
	}
}

func parseEntityChildren( entity types.Entity, path C.GnmiPath) {
	children := types.GetYChildren(entity.GetEntityData())	// []YChild
	if len(children) == 0 {
		// Check if any of the leafs has YFilter
		entPath := types.GetEntityPath(entity) // EntityPath
		for _, leafData := range(entPath.ValuePaths) {
			if leafData.Data.Filter != yfilter.NotSet {
				addPathElem(leafData.Name, path)
			        return	// Only one flagged element can be requested at a time
			}
		}
	}

	// Go over children entities
	for _, ychild := range(children) {
		if types.HasDataOrFilter(ychild.Value) {
			ydk.YLogDebug(fmt.Sprintf("parseEntityChildren: Looking at child '%s'", ychild.GoName));
            		parseEntity(ychild.Value, path);
		}
	}
}

func getEntityKeys( path string) map[string]string {
	s := path
	keys := map[string]string {}
	openbrac := strings.Index(s, "[")
	for (openbrac >= 0) {
		s = s[openbrac+1:]
		equal := strings.Index(s, "=")
		if equal > 0 {
			name := s[:equal]
			s = s[equal+1:]
			closebrack := strings.Index(s, "]")
			if closebrack > 0 {
				value := s[1:closebrack-1]
				s = s[closebrack+1:]
				keys[name] = value
			}
		}
		openbrac = strings.IndexByte(s, '[')
	}
	return keys
}

func parseEntity( entity types.Entity, path C.GnmiPath) {
	entPath := types.GetEntityPath(entity) // EntityPath
	s := entPath.Path //entity.GetEntityData().SegmentPath
	ydk.YLogDebug(fmt.Sprintf("parseEntity: entity segment path: '%s'", s))

	// Collect segment keys and leafs with set filter
	keys  := map[string]string {}
	leafs := map[string]string {}
	p := strings.Index(s, "[")
	if p > 0 {
		entityKeys := getEntityKeys(s)
		s = s[0:p]		
		for _, leafData := range(entPath.ValuePaths) {
			keyValue, isKey := entityKeys[leafData.Name]
			if  isKey {
				keys[leafData.Name] = keyValue
			} else if leafData.Data.Filter != yfilter.NotSet {
				leafs[leafData.Name] = leafData.Data.Filter.String()
			}
		}
	}

	elem := addPathElem(s, path)
	for key, value := range(keys) {
		ydk.YLogDebug(fmt.Sprintf("parseEntity: Adding key value: '%s:%s'", key, value))
		var ckey *C.char = C.CString(key)
		defer C.free(unsafe.Pointer(ckey))
		var cvalue *C.char = C.CString(value)
		defer C.free(unsafe.Pointer(cvalue))
		C.GnmiPathAddElemKey(elem, ckey, cvalue)
	}

	for leaf, _ := range(leafs) {
		addPathElem(leaf, path)
		return;	// Only one leaf with operation can be processed at a time
	}

	parseEntityChildren(entity, path);
}


