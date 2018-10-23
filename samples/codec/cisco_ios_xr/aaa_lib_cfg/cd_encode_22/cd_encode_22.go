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

const (
	expectedXML = `<aaa xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-lib-cfg">
  <usernames xmlns="http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-locald-cfg">
    <username>
      <name>netop</name>
      <ordering-index>22</ordering-index>
      <secret>$1$Z/8E$GDBQs1PtqJnwlQ9kKGpXj/</secret>
      <usergroup-under-usernames>
        <usergroup-under-username>
          <name>netadmin</name>
        </usergroup-under-username>
        <usergroup-under-username>
          <name>operator</name>
        </usergroup-under-username>
      </usergroup-under-usernames>
    </username>
  </usernames>
</aaa>
`
)

// config_aaa adds config data to aaa object.
func config_aaa(aaa *aaaCfg.Aaa) {
	username := aaaCfg.Aaa_Usernames_Username{}
	username.OrderingIndex = 22
	username.Name = "netop"
	username.Secret = "$1$Z/8E$GDBQs1PtqJnwlQ9kKGpXj/"

	// task group
	usergroupUnderUsername := aaaCfg.Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername{}
	usergroupUnderUsername.Name = "sysadmin"

	username.UsergroupUnderUsernames.UsergroupUnderUsername = append(
		username.UsergroupUnderUsernames.UsergroupUnderUsername, &usergroupUnderUsername)

	// task group
	usergroupUnderUsername = aaaCfg.Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername{}
	usergroupUnderUsername.Name = "operator"

	aaa.Usernames.Username = append(aaa.Usernames.Username, &username)
}

// main execute main program.
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

	// encode and print object
	fmt.Println(service.Encode(&provider, &aaa))
}
