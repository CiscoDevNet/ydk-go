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

    
    Snmpengine SNMPFRAMEWORKMIB_Snmpengine
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

    sNMPFRAMEWORKMIB.EntityData.Children = make(map[string]types.YChild)
    sNMPFRAMEWORKMIB.EntityData.Children["snmpEngine"] = types.YChild{"Snmpengine", &sNMPFRAMEWORKMIB.Snmpengine}
    sNMPFRAMEWORKMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sNMPFRAMEWORKMIB.EntityData)
}

// SNMPFRAMEWORKMIB_Snmpengine
type SNMPFRAMEWORKMIB_Snmpengine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    Snmpengineid interface{}

    // The type is interface{} with range: 1..2147483647.
    Snmpengineboots interface{}

    // The type is interface{} with range: 0..2147483647.
    Snmpenginetime interface{}

    // The type is interface{} with range: 484..2147483647.
    Snmpenginemaxmessagesize interface{}
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetEntityData() *types.CommonEntityData {
    snmpengine.EntityData.YFilter = snmpengine.YFilter
    snmpengine.EntityData.YangName = "snmpEngine"
    snmpengine.EntityData.BundleName = "cisco_ios_xr"
    snmpengine.EntityData.ParentYangName = "SNMP-FRAMEWORK-MIB"
    snmpengine.EntityData.SegmentPath = "snmpEngine"
    snmpengine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpengine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpengine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpengine.EntityData.Children = make(map[string]types.YChild)
    snmpengine.EntityData.Leafs = make(map[string]types.YLeaf)
    snmpengine.EntityData.Leafs["snmpEngineID"] = types.YLeaf{"Snmpengineid", snmpengine.Snmpengineid}
    snmpengine.EntityData.Leafs["snmpEngineBoots"] = types.YLeaf{"Snmpengineboots", snmpengine.Snmpengineboots}
    snmpengine.EntityData.Leafs["snmpEngineTime"] = types.YLeaf{"Snmpenginetime", snmpengine.Snmpenginetime}
    snmpengine.EntityData.Leafs["snmpEngineMaxMessageSize"] = types.YLeaf{"Snmpenginemaxmessagesize", snmpengine.Snmpenginemaxmessagesize}
    return &(snmpengine.EntityData)
}

