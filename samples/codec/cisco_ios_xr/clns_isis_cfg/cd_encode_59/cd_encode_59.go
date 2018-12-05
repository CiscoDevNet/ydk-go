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
	isis.Instances = isisCfg.Isis_Instances{
		Instance: make([]isisCfg.Isis_Instances_Instance, 1) }
	instance := &isis.Instances.Instance[0]
	instance.InstanceName = "DEFAULT"
	instance.Running = types.Empty{}
	instance.IsType = isisCfg.IsisConfigurableLevels_level2

	instance.Nets = isisCfg.Isis_Instances_Instance_Nets{
		Net: make([]isisCfg.Isis_Instances_Instance_Nets_Net, 1) }
	net := &instance.Nets.Net[0]
	net.NetName = "49.0000.1720.1625.5002.00"

	// global address family
	instance.Afs = isisCfg.Isis_Instances_Instance_Afs{
		Af: make([]isisCfg.Isis_Instances_Instance_Afs_Af, 1) }
	af := &instance.Afs.Af[0]
	af.AfName = isisDatatypes.IsisAddressFamily_ipv6
	af.SafName = isisDatatypes.IsisSubAddressFamily_unicast
	af.AfData = isisCfg.Isis_Instances_Instance_Afs_Af_AfData{}

	af.AfData.MetricStyles = isisCfg.Isis_Instances_Instance_Afs_Af_AfData_MetricStyles{
		MetricStyle: make(
			[]isisCfg.Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle, 1) }
	metricStyle := &af.AfData.MetricStyles.MetricStyle[0]
	metricStyle.Style = isisCfg.IsisMetricStyle_new_metric_style
	metricStyle.Level = isisDatatypes.IsisInternalLevel_not_set

	transitionState := isisCfg.IsisMetricStyleTransition_disabled
	metricStyle.TransitionState = transitionState

	// segment routing
	mpls := isisCfg.IsisLabelPreference_ldp
	af.AfData.SegmentRouting.Mpls = mpls
	af.AfData.SegmentRouting.PrefixSidMap.Receive = true

	// loopback interface
	instance.Interfaces = isisCfg.Isis_Instances_Instance_Interfaces{
		Interface_: make([]isisCfg.Isis_Instances_Instance_Interfaces_Interface, 2) }
	intrface := &instance.Interfaces.Interface_[0]
	intrface.InterfaceName = "Loopback0"
	intrface.Running = types.Empty{}
	intrface.State = isisCfg.IsisInterfaceState_passive

	//     interface address family
	intrface.InterfaceAfs = isisCfg.Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs{
		InterfaceAf: make(
			[]isisCfg.Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf, 1) }
	interfaceAf := &intrface.InterfaceAfs.InterfaceAf[0]
	interfaceAf.AfName = isisDatatypes.IsisAddressFamily_ipv6
	interfaceAf.SafName = isisDatatypes.IsisSubAddressFamily_unicast
	interfaceAf.InterfaceAfData.Running = types.Empty{}

	//		segment routing
	prefixSID := &interfaceAf.InterfaceAfData.PrefixSid
	prefixSID.Type_ = isisCfg.Isissid1_absolute
	prefixSID.Value = 16062
	prefixSID.Php = isisCfg.IsisphpFlag_enable
	prefixSID.ExplicitNull = isisCfg.IsisexplicitNullFlag_disable
	prefixSID.NflagClear = isisCfg.NflagClear_disable

	// gi0/0/0/0 interface
	intrface = &instance.Interfaces.Interface_[1]
	intrface.InterfaceName = "GigabitEthernet0/0/0/0"
	intrface.Running = types.Empty{}
	intrface.PointToPoint = types.Empty{}

	//     interface address family
	intrface.InterfaceAfs = isisCfg.Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs{
		InterfaceAf: make(
			[]isisCfg.Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf, 1) }
	interfaceAf = &intrface.InterfaceAfs.InterfaceAf[0]
	interfaceAf.AfName = isisDatatypes.IsisAddressFamily_ipv6
	interfaceAf.SafName = isisDatatypes.IsisSubAddressFamily_unicast
	interfaceAf.InterfaceAfData.Running = types.Empty{}
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
