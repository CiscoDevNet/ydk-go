package snmp_mpd_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_mpd_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMP-MPD-MIB/200210140000Z SNMP-MPD-MIB}", reflect.TypeOf(SNMPMPDMIB{}))
    ydk.RegisterEntity("SNMP-MPD-MIB:SNMP-MPD-MIB", reflect.TypeOf(SNMPMPDMIB{}))
}

// SNMPMPDMIB
type SNMPMPDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SnmpMPDStats SNMPMPDMIB_SnmpMPDStats
}

func (sNMPMPDMIB *SNMPMPDMIB) GetEntityData() *types.CommonEntityData {
    sNMPMPDMIB.EntityData.YFilter = sNMPMPDMIB.YFilter
    sNMPMPDMIB.EntityData.YangName = "SNMP-MPD-MIB"
    sNMPMPDMIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPMPDMIB.EntityData.ParentYangName = "SNMP-MPD-MIB"
    sNMPMPDMIB.EntityData.SegmentPath = "SNMP-MPD-MIB:SNMP-MPD-MIB"
    sNMPMPDMIB.EntityData.AbsolutePath = sNMPMPDMIB.EntityData.SegmentPath
    sNMPMPDMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPMPDMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPMPDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPMPDMIB.EntityData.Children = types.NewOrderedMap()
    sNMPMPDMIB.EntityData.Children.Append("snmpMPDStats", types.YChild{"SnmpMPDStats", &sNMPMPDMIB.SnmpMPDStats})
    sNMPMPDMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPMPDMIB.EntityData.YListKeys = []string {}

    return &(sNMPMPDMIB.EntityData)
}

// SNMPMPDMIB_SnmpMPDStats
type SNMPMPDMIB_SnmpMPDStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    SnmpUnknownSecurityModels interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpInvalidMsgs interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpUnknownPDUHandlers interface{}
}

func (snmpMPDStats *SNMPMPDMIB_SnmpMPDStats) GetEntityData() *types.CommonEntityData {
    snmpMPDStats.EntityData.YFilter = snmpMPDStats.YFilter
    snmpMPDStats.EntityData.YangName = "snmpMPDStats"
    snmpMPDStats.EntityData.BundleName = "cisco_ios_xr"
    snmpMPDStats.EntityData.ParentYangName = "SNMP-MPD-MIB"
    snmpMPDStats.EntityData.SegmentPath = "snmpMPDStats"
    snmpMPDStats.EntityData.AbsolutePath = "SNMP-MPD-MIB:SNMP-MPD-MIB/" + snmpMPDStats.EntityData.SegmentPath
    snmpMPDStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpMPDStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpMPDStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpMPDStats.EntityData.Children = types.NewOrderedMap()
    snmpMPDStats.EntityData.Leafs = types.NewOrderedMap()
    snmpMPDStats.EntityData.Leafs.Append("snmpUnknownSecurityModels", types.YLeaf{"SnmpUnknownSecurityModels", snmpMPDStats.SnmpUnknownSecurityModels})
    snmpMPDStats.EntityData.Leafs.Append("snmpInvalidMsgs", types.YLeaf{"SnmpInvalidMsgs", snmpMPDStats.SnmpInvalidMsgs})
    snmpMPDStats.EntityData.Leafs.Append("snmpUnknownPDUHandlers", types.YLeaf{"SnmpUnknownPDUHandlers", snmpMPDStats.SnmpUnknownPDUHandlers})

    snmpMPDStats.EntityData.YListKeys = []string {}

    return &(snmpMPDStats.EntityData)
}

