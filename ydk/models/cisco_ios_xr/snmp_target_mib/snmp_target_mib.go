package snmp_target_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_target_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMP-TARGET-MIB/200210140000Z SNMP-TARGET-MIB}", reflect.TypeOf(SNMPTARGETMIB{}))
    ydk.RegisterEntity("SNMP-TARGET-MIB:SNMP-TARGET-MIB", reflect.TypeOf(SNMPTARGETMIB{}))
}

// SNMPTARGETMIB
type SNMPTARGETMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SnmpTargetObjects SNMPTARGETMIB_SnmpTargetObjects

    
    SnmpTargetAddrTable SNMPTARGETMIB_SnmpTargetAddrTable

    
    SnmpTargetParamsTable SNMPTARGETMIB_SnmpTargetParamsTable
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetEntityData() *types.CommonEntityData {
    sNMPTARGETMIB.EntityData.YFilter = sNMPTARGETMIB.YFilter
    sNMPTARGETMIB.EntityData.YangName = "SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPTARGETMIB.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.SegmentPath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.AbsolutePath = sNMPTARGETMIB.EntityData.SegmentPath
    sNMPTARGETMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPTARGETMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPTARGETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPTARGETMIB.EntityData.Children = types.NewOrderedMap()
    sNMPTARGETMIB.EntityData.Children.Append("snmpTargetObjects", types.YChild{"SnmpTargetObjects", &sNMPTARGETMIB.SnmpTargetObjects})
    sNMPTARGETMIB.EntityData.Children.Append("snmpTargetAddrTable", types.YChild{"SnmpTargetAddrTable", &sNMPTARGETMIB.SnmpTargetAddrTable})
    sNMPTARGETMIB.EntityData.Children.Append("snmpTargetParamsTable", types.YChild{"SnmpTargetParamsTable", &sNMPTARGETMIB.SnmpTargetParamsTable})
    sNMPTARGETMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPTARGETMIB.EntityData.YListKeys = []string {}

    return &(sNMPTARGETMIB.EntityData)
}

// SNMPTARGETMIB_SnmpTargetObjects
type SNMPTARGETMIB_SnmpTargetObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    SnmpUnavailableContexts interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpUnknownContexts interface{}
}

func (snmpTargetObjects *SNMPTARGETMIB_SnmpTargetObjects) GetEntityData() *types.CommonEntityData {
    snmpTargetObjects.EntityData.YFilter = snmpTargetObjects.YFilter
    snmpTargetObjects.EntityData.YangName = "snmpTargetObjects"
    snmpTargetObjects.EntityData.BundleName = "cisco_ios_xr"
    snmpTargetObjects.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmpTargetObjects.EntityData.SegmentPath = "snmpTargetObjects"
    snmpTargetObjects.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/" + snmpTargetObjects.EntityData.SegmentPath
    snmpTargetObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpTargetObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpTargetObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpTargetObjects.EntityData.Children = types.NewOrderedMap()
    snmpTargetObjects.EntityData.Leafs = types.NewOrderedMap()
    snmpTargetObjects.EntityData.Leafs.Append("snmpUnavailableContexts", types.YLeaf{"SnmpUnavailableContexts", snmpTargetObjects.SnmpUnavailableContexts})
    snmpTargetObjects.EntityData.Leafs.Append("snmpUnknownContexts", types.YLeaf{"SnmpUnknownContexts", snmpTargetObjects.SnmpUnknownContexts})

    snmpTargetObjects.EntityData.YListKeys = []string {}

    return &(snmpTargetObjects.EntityData)
}

// SNMPTARGETMIB_SnmpTargetAddrTable
type SNMPTARGETMIB_SnmpTargetAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry.
    SnmpTargetAddrEntry []*SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry
}

func (snmpTargetAddrTable *SNMPTARGETMIB_SnmpTargetAddrTable) GetEntityData() *types.CommonEntityData {
    snmpTargetAddrTable.EntityData.YFilter = snmpTargetAddrTable.YFilter
    snmpTargetAddrTable.EntityData.YangName = "snmpTargetAddrTable"
    snmpTargetAddrTable.EntityData.BundleName = "cisco_ios_xr"
    snmpTargetAddrTable.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmpTargetAddrTable.EntityData.SegmentPath = "snmpTargetAddrTable"
    snmpTargetAddrTable.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/" + snmpTargetAddrTable.EntityData.SegmentPath
    snmpTargetAddrTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpTargetAddrTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpTargetAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpTargetAddrTable.EntityData.Children = types.NewOrderedMap()
    snmpTargetAddrTable.EntityData.Children.Append("snmpTargetAddrEntry", types.YChild{"SnmpTargetAddrEntry", nil})
    for i := range snmpTargetAddrTable.SnmpTargetAddrEntry {
        snmpTargetAddrTable.EntityData.Children.Append(types.GetSegmentPath(snmpTargetAddrTable.SnmpTargetAddrEntry[i]), types.YChild{"SnmpTargetAddrEntry", snmpTargetAddrTable.SnmpTargetAddrEntry[i]})
    }
    snmpTargetAddrTable.EntityData.Leafs = types.NewOrderedMap()

    snmpTargetAddrTable.EntityData.YListKeys = []string {}

    return &(snmpTargetAddrTable.EntityData)
}

// SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry
type SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    SnmpTargetAddrName interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    // This attribute is mandatory.
    SnmpTargetAddrTDomain interface{}

    // The type is one of the following types: string with pattern:
    // b'(\\d*(.\\d*)*)?' This attribute is mandatory., or string with pattern:
    // b'(\\d*(.\\d*)*)?' This attribute is mandatory..
    SnmpTargetAddrTAddress interface{}

    // The type is interface{} with range: 0..2147483647. The default value is
    // 1500.
    SnmpTargetAddrTimeout interface{}

    // The type is interface{} with range: 0..255. The default value is 3.
    SnmpTargetAddrRetryCount interface{}

    // The type is string with length: 0..255.
    SnmpTargetAddrTagList interface{}

    // The type is string with length: 1..32. This attribute is mandatory.
    SnmpTargetAddrParams interface{}

    // The type is StorageType. The default value is nonVolatile.
    SnmpTargetAddrStorageType interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    SnmpTargetAddrEngineID interface{}

    // The type is one of the following types: string with pattern:
    // b'(\\d*(.\\d*)*)?', or string with pattern: b'(\\d*(.\\d*)*)?'., or string
    // with pattern: b'(\\d*(.\\d*)*)?'.
    SnmpTargetAddrTMask interface{}

    // The type is interface{} with range: 0..0 | 484..2147483647. The default
    // value is 2048.
    SnmpTargetAddrMMS interface{}

    // The type is bool. The default value is true.
    Enabled interface{}
}

func (snmpTargetAddrEntry *SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry) GetEntityData() *types.CommonEntityData {
    snmpTargetAddrEntry.EntityData.YFilter = snmpTargetAddrEntry.YFilter
    snmpTargetAddrEntry.EntityData.YangName = "snmpTargetAddrEntry"
    snmpTargetAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    snmpTargetAddrEntry.EntityData.ParentYangName = "snmpTargetAddrTable"
    snmpTargetAddrEntry.EntityData.SegmentPath = "snmpTargetAddrEntry" + types.AddKeyToken(snmpTargetAddrEntry.SnmpTargetAddrName, "snmpTargetAddrName")
    snmpTargetAddrEntry.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/snmpTargetAddrTable/" + snmpTargetAddrEntry.EntityData.SegmentPath
    snmpTargetAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpTargetAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpTargetAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpTargetAddrEntry.EntityData.Children = types.NewOrderedMap()
    snmpTargetAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrName", types.YLeaf{"SnmpTargetAddrName", snmpTargetAddrEntry.SnmpTargetAddrName})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTDomain", types.YLeaf{"SnmpTargetAddrTDomain", snmpTargetAddrEntry.SnmpTargetAddrTDomain})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTAddress", types.YLeaf{"SnmpTargetAddrTAddress", snmpTargetAddrEntry.SnmpTargetAddrTAddress})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTimeout", types.YLeaf{"SnmpTargetAddrTimeout", snmpTargetAddrEntry.SnmpTargetAddrTimeout})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrRetryCount", types.YLeaf{"SnmpTargetAddrRetryCount", snmpTargetAddrEntry.SnmpTargetAddrRetryCount})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTagList", types.YLeaf{"SnmpTargetAddrTagList", snmpTargetAddrEntry.SnmpTargetAddrTagList})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrParams", types.YLeaf{"SnmpTargetAddrParams", snmpTargetAddrEntry.SnmpTargetAddrParams})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrStorageType", types.YLeaf{"SnmpTargetAddrStorageType", snmpTargetAddrEntry.SnmpTargetAddrStorageType})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrEngineID", types.YLeaf{"SnmpTargetAddrEngineID", snmpTargetAddrEntry.SnmpTargetAddrEngineID})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTMask", types.YLeaf{"SnmpTargetAddrTMask", snmpTargetAddrEntry.SnmpTargetAddrTMask})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrMMS", types.YLeaf{"SnmpTargetAddrMMS", snmpTargetAddrEntry.SnmpTargetAddrMMS})
    snmpTargetAddrEntry.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", snmpTargetAddrEntry.Enabled})

    snmpTargetAddrEntry.EntityData.YListKeys = []string {"SnmpTargetAddrName"}

    return &(snmpTargetAddrEntry.EntityData)
}

// SNMPTARGETMIB_SnmpTargetParamsTable
type SNMPTARGETMIB_SnmpTargetParamsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry.
    SnmpTargetParamsEntry []*SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry
}

func (snmpTargetParamsTable *SNMPTARGETMIB_SnmpTargetParamsTable) GetEntityData() *types.CommonEntityData {
    snmpTargetParamsTable.EntityData.YFilter = snmpTargetParamsTable.YFilter
    snmpTargetParamsTable.EntityData.YangName = "snmpTargetParamsTable"
    snmpTargetParamsTable.EntityData.BundleName = "cisco_ios_xr"
    snmpTargetParamsTable.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmpTargetParamsTable.EntityData.SegmentPath = "snmpTargetParamsTable"
    snmpTargetParamsTable.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/" + snmpTargetParamsTable.EntityData.SegmentPath
    snmpTargetParamsTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpTargetParamsTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpTargetParamsTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpTargetParamsTable.EntityData.Children = types.NewOrderedMap()
    snmpTargetParamsTable.EntityData.Children.Append("snmpTargetParamsEntry", types.YChild{"SnmpTargetParamsEntry", nil})
    for i := range snmpTargetParamsTable.SnmpTargetParamsEntry {
        snmpTargetParamsTable.EntityData.Children.Append(types.GetSegmentPath(snmpTargetParamsTable.SnmpTargetParamsEntry[i]), types.YChild{"SnmpTargetParamsEntry", snmpTargetParamsTable.SnmpTargetParamsEntry[i]})
    }
    snmpTargetParamsTable.EntityData.Leafs = types.NewOrderedMap()

    snmpTargetParamsTable.EntityData.YListKeys = []string {}

    return &(snmpTargetParamsTable.EntityData)
}

// SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry
type SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    SnmpTargetParamsName interface{}

    // The type is interface{} with range: 0..2147483647. This attribute is
    // mandatory.
    SnmpTargetParamsMPModel interface{}

    // The type is interface{} with range: 1..2147483647. This attribute is
    // mandatory.
    SnmpTargetParamsSecurityModel interface{}

    // The type is string with length: 0..255. This attribute is mandatory.
    SnmpTargetParamsSecurityName interface{}

    // The type is SnmpSecurityLevel. This attribute is mandatory.
    SnmpTargetParamsSecurityLevel interface{}

    // The type is StorageType. The default value is nonVolatile.
    SnmpTargetParamsStorageType interface{}
}

func (snmpTargetParamsEntry *SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry) GetEntityData() *types.CommonEntityData {
    snmpTargetParamsEntry.EntityData.YFilter = snmpTargetParamsEntry.YFilter
    snmpTargetParamsEntry.EntityData.YangName = "snmpTargetParamsEntry"
    snmpTargetParamsEntry.EntityData.BundleName = "cisco_ios_xr"
    snmpTargetParamsEntry.EntityData.ParentYangName = "snmpTargetParamsTable"
    snmpTargetParamsEntry.EntityData.SegmentPath = "snmpTargetParamsEntry" + types.AddKeyToken(snmpTargetParamsEntry.SnmpTargetParamsName, "snmpTargetParamsName")
    snmpTargetParamsEntry.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/snmpTargetParamsTable/" + snmpTargetParamsEntry.EntityData.SegmentPath
    snmpTargetParamsEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpTargetParamsEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpTargetParamsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpTargetParamsEntry.EntityData.Children = types.NewOrderedMap()
    snmpTargetParamsEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsName", types.YLeaf{"SnmpTargetParamsName", snmpTargetParamsEntry.SnmpTargetParamsName})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsMPModel", types.YLeaf{"SnmpTargetParamsMPModel", snmpTargetParamsEntry.SnmpTargetParamsMPModel})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsSecurityModel", types.YLeaf{"SnmpTargetParamsSecurityModel", snmpTargetParamsEntry.SnmpTargetParamsSecurityModel})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsSecurityName", types.YLeaf{"SnmpTargetParamsSecurityName", snmpTargetParamsEntry.SnmpTargetParamsSecurityName})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsSecurityLevel", types.YLeaf{"SnmpTargetParamsSecurityLevel", snmpTargetParamsEntry.SnmpTargetParamsSecurityLevel})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsStorageType", types.YLeaf{"SnmpTargetParamsStorageType", snmpTargetParamsEntry.SnmpTargetParamsStorageType})

    snmpTargetParamsEntry.EntityData.YListKeys = []string {"SnmpTargetParamsName"}

    return &(snmpTargetParamsEntry.EntityData)
}

