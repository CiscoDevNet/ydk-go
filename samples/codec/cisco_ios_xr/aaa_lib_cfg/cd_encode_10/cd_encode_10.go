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

import (
	"flag"
	"fmt"
	aaaCfg "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr/aaa_lib_cfg"
	"github.com/CiscoDevNet/ydk-go/ydk/providers"
	"github.com/CiscoDevNet/ydk-go/ydk/services"
	"github.com/CiscoDevNet/ydk-go/ydk"
	encoding "github.com/CiscoDevNet/ydk-go/ydk/types/encoding_format"
)

// config_aaa adds config data to aaa object.
func config_aaa(aaa *aaaCfg.Aaa) {
}

// main executes main program.
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nError occured!!\n ", r)
		}
	}()

	vPtr := flag.Bool("v", false, "Enable verbose")
	flag.Parse()

	if *vPtr {
		ydk.EnableLogging(ydk.Info)
	}

	// create codec provider
	var provider = providers.CodecServiceProvider{}
	provider.Encoding = encoding.XML

	// create codec service
	var service = services.CodecService{}

	aaa := aaaCfg.Aaa{}		// create an object
	config_aaa(&aaa)		// add object configuration

	_ = service
}
