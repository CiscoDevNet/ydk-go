/*
 * ------------------------------------------------------------------
 * YANG Development Kit
 * Copyright 2017 Cisco Systems. All rights reserved
 *
 *----------------------------------------------
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http:www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *----------------------------------------------
 */
package main

// #cgo CXXFLAGS: -g -std=c++11
// #cgo darwin LDFLAGS:  -fprofile-arcs -ftest-coverage -lydk -lxml2 -lxslt -lpcre -lssh -lssh_threads -lcurl -lpython -lc++
// #cgo linux LDFLAGS:  -fprofile-arcs -ftest-coverage --coverage -lydk -lxml2 -lxslt -lpcre -lssh -lssh_threads -lcurl -lstdc++ -lpython2.7 -ldl
// #include <ydk/ydk.h>
// #include <stdlib.h>
import "C"

import (
	"flag"
	"fmt"
	"strings"
	"strconv"
	"unsafe"
)

func perform_operation(device map[string]string, port int) {
	var cstate C.YDKStatePtr = C.YDKStateCreate()
	defer C.YDKStateFree(cstate)

	var address *C.char = C.CString(device["address"])
	defer C.free(unsafe.Pointer(address))
	var username *C.char = C.CString(device["username"])
	defer C.free(unsafe.Pointer(username))
	var password *C.char = C.CString(device["password"])
	defer C.free(unsafe.Pointer(password))
	var protocol *C.char = C.CString(device["protocol"])
	defer C.free(unsafe.Pointer(protocol))
	var runner_path *C.char = C.CString("openconfig-bgp:bgp")
	defer C.free(unsafe.Pointer(runner_path))
	var number_path *C.char = C.CString("global/config/as")
	defer C.free(unsafe.Pointer(number_path))
	var number_value *C.char = C.CString("65001")
	defer C.free(unsafe.Pointer(number_value))
	var create_path *C.char = C.CString("ietf-netconf:edit-config")
	defer C.free(unsafe.Pointer(create_path))
	var read_path *C.char = C.CString("ietf-netconf:get-config")
	defer C.free(unsafe.Pointer(read_path))
	var entity_path *C.char = C.CString("config")
	defer C.free(unsafe.Pointer(entity_path))
	var target_path *C.char = C.CString("target/running")
	defer C.free(unsafe.Pointer(target_path))
	var source_path *C.char = C.CString("source/running")
	defer C.free(unsafe.Pointer(source_path))
	var filter_path *C.char = C.CString("filter")
	defer C.free(unsafe.Pointer(filter_path))
	var empty *C.char = C.CString("")
	defer C.free(unsafe.Pointer(empty))
	var c_port C.int = C.int(port)

	codec := C.CodecInit()
	defer C.CodecFree(codec)
	provider := C.NetconfServiceProviderInit(cstate, address, username, password, c_port, protocol)
	defer C.NetconfServiceProviderFree(provider)
	root_schema := C.ServiceProviderGetRootSchema(cstate, provider)

	runner := C.RootSchemaNodeCreate(cstate, root_schema, runner_path)

	C.DataNodeCreate(cstate, runner, number_path, number_value)
	var create_xml *C.char = C.CodecEncode(cstate, codec, runner, C.XML, 0)
	defer C.free(unsafe.Pointer(create_xml))

	create_rpc := C.RootSchemaNodeRpc(cstate, root_schema, create_path)
	input := C.RpcInput(cstate, create_rpc)
	C.DataNodeCreate(cstate, input, target_path, empty)
	C.DataNodeCreate(cstate, input, entity_path, create_xml)
	C.RpcExecute(cstate, create_rpc, provider)

	read_rpc := C.RootSchemaNodeRpc(cstate, root_schema, read_path)
	input = C.RpcInput(cstate, read_rpc)
	runner_filter := C.RootSchemaNodeCreate(cstate, root_schema, runner_path)
	C.DataNodeCreate(cstate, input, source_path, empty)
	var read_xml *C.char = C.CodecEncode(cstate, codec, runner_filter, C.XML, 0)
	defer C.free(unsafe.Pointer(read_xml))

	C.DataNodeCreate(cstate, input, filter_path, read_xml)
	read_data := C.RpcExecute(cstate, read_rpc, provider)
	if read_data == nil {
		fmt.Println("No data")
	}
	var data *C.char = C.CodecEncode(cstate, codec, read_data, C.XML, 1)
	defer C.free(unsafe.Pointer(data))
	s := C.GoString(data)
	fmt.Println("Read data:", s)
}

func main() {
	verbosePtr := flag.Bool("verbose", false, "Enable verbose")
	devicePtr := flag.String(
		"device",
		"",
		"NETCONF device (ssh://user:password@host:port)")
	flag.Parse()

	if (*devicePtr == "") {
		panic("Missing device arg see --help for details")
	}

	if *verbosePtr {
		C.EnableLogging(C.INFO)
	}

	denominators := []string{"://", ":", "@", ":"}
	keys := []string {"protocol", "username", "password", "address", "port"}
	device := make(map[string]string)

	var split []string
	unprocessed := *devicePtr
	for i := 0; i < 4; i++ {
		if (!strings.Contains(unprocessed, denominators[i])) {
			panic(fmt.Sprintln("Device arg: device must be entered in",
				"ssh://user:password@host:port format"))
		}
		split = strings.SplitN(unprocessed, denominators[i], 2)

		device[keys[i]] = split[0]
		unprocessed = split[1]
	}
	device[keys[4]] = unprocessed

	port, err := strconv.Atoi(device["port"])
	if (err != nil) {
		panic("Device arg: port must be an int")
	}

    perform_operation(device, port)
}

