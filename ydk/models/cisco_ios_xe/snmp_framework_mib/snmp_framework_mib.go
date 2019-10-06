// The SNMP Management Architecture MIB
// 
// Copyright (C) The Internet Society (2002). This
// version of this MIB module is part of RFC 3411;
// see the RFC itself for full legal notices.
package snmp_framework_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_framework_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:SNMP-FRAMEWORK-MIB SNMP-FRAMEWORK-MIB}", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
    ydk.RegisterEntity("SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
}

type SnmpPrivProtocols struct {
}

func (id SnmpPrivProtocols) String() string {
	return "SNMP-FRAMEWORK-MIB:snmpPrivProtocols"
}

type SnmpAuthProtocols struct {
}

func (id SnmpAuthProtocols) String() string {
	return "SNMP-FRAMEWORK-MIB:snmpAuthProtocols"
}

// SnmpSecurityLevel represents authNoPriv is less than authPriv.
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
    sNMPFRAMEWORKMIB.EntityData.BundleName = "cisco_ios_xe"
    sNMPFRAMEWORKMIB.EntityData.ParentYangName = "SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.SegmentPath = "SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.AbsolutePath = sNMPFRAMEWORKMIB.EntityData.SegmentPath
    sNMPFRAMEWORKMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sNMPFRAMEWORKMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sNMPFRAMEWORKMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

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

    // An SNMP engine's administratively-unique identifier.  This information
    // SHOULD be stored in non-volatile storage so that it remains constant across
    // re-initializations of the SNMP engine. The type is string with length:
    // 5..32.
    SnmpEngineID interface{}

    // The number of times that the SNMP engine has (re-)initialized itself since
    // snmpEngineID was last configured. The type is interface{} with range:
    // 1..2147483647.
    SnmpEngineBoots interface{}

    // The number of seconds since the value of the snmpEngineBoots object last
    // changed. When incrementing this object's value would cause it to exceed its
    // maximum, snmpEngineBoots is incremented as if a re-initialization had
    // occurred, and this object's value consequently reverts to zero. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    SnmpEngineTime interface{}

    // The maximum length in octets of an SNMP message which this SNMP engine can
    // send or receive and process, determined as the minimum of the maximum
    // message size values supported among all of the transports available to and
    // supported by the engine. The type is interface{} with range:
    // 484..2147483647.
    SnmpEngineMaxMessageSize interface{}
}

func (snmpEngine *SNMPFRAMEWORKMIB_SnmpEngine) GetEntityData() *types.CommonEntityData {
    snmpEngine.EntityData.YFilter = snmpEngine.YFilter
    snmpEngine.EntityData.YangName = "snmpEngine"
    snmpEngine.EntityData.BundleName = "cisco_ios_xe"
    snmpEngine.EntityData.ParentYangName = "SNMP-FRAMEWORK-MIB"
    snmpEngine.EntityData.SegmentPath = "snmpEngine"
    snmpEngine.EntityData.AbsolutePath = "SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB/" + snmpEngine.EntityData.SegmentPath
    snmpEngine.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpEngine.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpEngine.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpEngine.EntityData.Children = types.NewOrderedMap()
    snmpEngine.EntityData.Leafs = types.NewOrderedMap()
    snmpEngine.EntityData.Leafs.Append("snmpEngineID", types.YLeaf{"SnmpEngineID", snmpEngine.SnmpEngineID})
    snmpEngine.EntityData.Leafs.Append("snmpEngineBoots", types.YLeaf{"SnmpEngineBoots", snmpEngine.SnmpEngineBoots})
    snmpEngine.EntityData.Leafs.Append("snmpEngineTime", types.YLeaf{"SnmpEngineTime", snmpEngine.SnmpEngineTime})
    snmpEngine.EntityData.Leafs.Append("snmpEngineMaxMessageSize", types.YLeaf{"SnmpEngineMaxMessageSize", snmpEngine.SnmpEngineMaxMessageSize})

    snmpEngine.EntityData.YListKeys = []string {}

    return &(snmpEngine.EntityData)
}

