// This module contains definitions
// for the Calvados model objects.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package snmp_framework_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_framework_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMP-FRAMEWORK-MIB/200210140000Z SNMP-FRAMEWORK-MIB}", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
    ydk.RegisterEntity("SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
}

// SnmpSecurityLevel
type SnmpSecurityLevel string

const (
    SnmpSecurityLevel_noAuthNoPriv SnmpSecurityLevel = "noAuthNoPriv"

    SnmpSecurityLevel_authNoPriv SnmpSecurityLevel = "authNoPriv"

    SnmpSecurityLevel_authPriv SnmpSecurityLevel = "authPriv"
)

// SNMPFRAMEWORKMIB
type SNMPFRAMEWORKMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SnmpEngine SNMPFRAMEWORKMIB_SnmpEngine
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetEntityData() *types.CommonEntityData {
    sNMPFRAMEWORKMIB.EntityData.YFilter = sNMPFRAMEWORKMIB.YFilter
    sNMPFRAMEWORKMIB.EntityData.YangName = "SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPFRAMEWORKMIB.EntityData.ParentYangName = "SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.SegmentPath = "SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPFRAMEWORKMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPFRAMEWORKMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPFRAMEWORKMIB.EntityData.Children = types.NewOrderedMap()
    sNMPFRAMEWORKMIB.EntityData.Children.Append("snmpEngine", types.YChild{"SnmpEngine", &sNMPFRAMEWORKMIB.SnmpEngine})
    sNMPFRAMEWORKMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPFRAMEWORKMIB.EntityData.YListKeys = []string {}

    return &(sNMPFRAMEWORKMIB.EntityData)
}

// SNMPFRAMEWORKMIB_SnmpEngine
type SNMPFRAMEWORKMIB_SnmpEngine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: (([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?.
    SnmpEngineID interface{}

    // The type is interface{} with range: 1..2147483647.
    SnmpEngineBoots interface{}

    // The type is interface{} with range: 0..2147483647.
    SnmpEngineTime interface{}

    // The type is interface{} with range: 484..2147483647.
    SnmpEngineMaxMessageSize interface{}
}

func (snmpEngine *SNMPFRAMEWORKMIB_SnmpEngine) GetEntityData() *types.CommonEntityData {
    snmpEngine.EntityData.YFilter = snmpEngine.YFilter
    snmpEngine.EntityData.YangName = "snmpEngine"
    snmpEngine.EntityData.BundleName = "cisco_ios_xr"
    snmpEngine.EntityData.ParentYangName = "SNMP-FRAMEWORK-MIB"
    snmpEngine.EntityData.SegmentPath = "snmpEngine"
    snmpEngine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpEngine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpEngine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpEngine.EntityData.Children = types.NewOrderedMap()
    snmpEngine.EntityData.Leafs = types.NewOrderedMap()
    snmpEngine.EntityData.Leafs.Append("snmpEngineID", types.YLeaf{"SnmpEngineID", snmpEngine.SnmpEngineID})
    snmpEngine.EntityData.Leafs.Append("snmpEngineBoots", types.YLeaf{"SnmpEngineBoots", snmpEngine.SnmpEngineBoots})
    snmpEngine.EntityData.Leafs.Append("snmpEngineTime", types.YLeaf{"SnmpEngineTime", snmpEngine.SnmpEngineTime})
    snmpEngine.EntityData.Leafs.Append("snmpEngineMaxMessageSize", types.YLeaf{"SnmpEngineMaxMessageSize", snmpEngine.SnmpEngineMaxMessageSize})

    snmpEngine.EntityData.YListKeys = []string {}

    return &(snmpEngine.EntityData)
}

