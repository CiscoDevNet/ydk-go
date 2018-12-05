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
	isisCfg "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr/clns_isis_cfg"
	isisDatatypes "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr/clns_isis_datatypes"
	"github.com/CiscoDevNet/ydk-go/ydk/providers"
	"github.com/CiscoDevNet/ydk-go/ydk/services"
	"github.com/CiscoDevNet/ydk-go/ydk"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
	encoding "github.com/CiscoDevNet/ydk-go/ydk/types/encoding_format"
)

// configIsis adds config data to isis object.
func configIsis(isis *isisCfg.Isis) {
	// global configuration
	instance := isisCfg.Isis_Instances_Instance{}
	instance.InstanceName = "DEFAULT"
	instance.Running = types.Empty{}
	instance.IsType = isisCfg.IsisConfigurableLevels_level2
	isis.Instances.Instance = append(isis.Instances.Instance, &instance)

	net := isisCfg.Isis_Instances_Instance_Nets_Net{}
	net.NetName = "49.0000.1720.1625.5001.00"
	instance.Nets.Net = append(instance.Nets.Net, &net)

	// global address family
	af := isisCfg.Isis_Instances_Instance_Afs_Af{}
	af.AfName = isisDatatypes.IsisAddressFamily_ipv4
	af.SafName = isisDatatypes.IsisSubAddressFamily_unicast
	af.AfData = isisCfg.Isis_Instances_Instance_Afs_Af_AfData{}
	instance.Afs.Af = append(instance.Afs.Af, &af)

	metricStyle := isisCfg.Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle{}
	metricStyle.Style = isisCfg.IsisMetricStyle_new_metric_style
	metricStyle.Level = isisDatatypes.IsisInternalLevel_not_set
	af.AfData.MetricStyles.MetricStyle = append(af.AfData.MetricStyles.MetricStyle, &metricStyle)

	// loopback interface
	intrface := isisCfg.Isis_Instances_Instance_Interfaces_Interface{}
	intrface.InterfaceName = "Loopback0"
	intrface.Running = types.Empty{}
	intrface.State = isisCfg.IsisInterfaceState_passive
	instance.Interfaces.Interface = append(instance.Interfaces.Interface, &intrface)

	// interface address family
	interfaceAf := isisCfg.Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf{}
	interfaceAf.AfName = isisDatatypes.IsisAddressFamily_ipv4
	interfaceAf.SafName = isisDatatypes.IsisSubAddressFamily_unicast
	interfaceAf.InterfaceAfData.Running = types.Empty{}
	intrface.InterfaceAfs.InterfaceAf = append(intrface.InterfaceAfs.InterfaceAf, &interfaceAf)

	// gi0/0/0/0 interface
	intrface = isisCfg.Isis_Instances_Instance_Interfaces_Interface{}
	intrface.InterfaceName = "GigabitEthernet0/0/0/0"
	intrface.Running = types.Empty{}
	intrface.PointToPoint = types.Empty{}
	instance.Interfaces.Interface = append(instance.Interfaces.Interface, &intrface)

	//     interface address family
	interfaceAf = isisCfg.Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf{}
	interfaceAf.AfName = isisDatatypes.IsisAddressFamily_ipv4
	interfaceAf.SafName = isisDatatypes.IsisSubAddressFamily_unicast
	interfaceAf.InterfaceAfData.Running = types.Empty{}
	intrface.InterfaceAfs.InterfaceAf = append(intrface.InterfaceAfs.InterfaceAf, &interfaceAf)
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

	isis := isisCfg.Isis{}		// create an object
	configIsis(&isis)			// add object configuration

	// encode and print object
	payload := service.Encode(&provider, &isis)
	fmt.Println(payload)
}
