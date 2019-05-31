package snmp_community_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_community_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMP-COMMUNITY-MIB/200308060000Z SNMP-COMMUNITY-MIB}", reflect.TypeOf(SNMPCOMMUNITYMIB{}))
    ydk.RegisterEntity("SNMP-COMMUNITY-MIB:SNMP-COMMUNITY-MIB", reflect.TypeOf(SNMPCOMMUNITYMIB{}))
}

// SNMPCOMMUNITYMIB
type SNMPCOMMUNITYMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SnmpCommunityTable SNMPCOMMUNITYMIB_SnmpCommunityTable
}

func (sNMPCOMMUNITYMIB *SNMPCOMMUNITYMIB) GetEntityData() *types.CommonEntityData {
    sNMPCOMMUNITYMIB.EntityData.YFilter = sNMPCOMMUNITYMIB.YFilter
    sNMPCOMMUNITYMIB.EntityData.YangName = "SNMP-COMMUNITY-MIB"
    sNMPCOMMUNITYMIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPCOMMUNITYMIB.EntityData.ParentYangName = "SNMP-COMMUNITY-MIB"
    sNMPCOMMUNITYMIB.EntityData.SegmentPath = "SNMP-COMMUNITY-MIB:SNMP-COMMUNITY-MIB"
    sNMPCOMMUNITYMIB.EntityData.AbsolutePath = sNMPCOMMUNITYMIB.EntityData.SegmentPath
    sNMPCOMMUNITYMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPCOMMUNITYMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPCOMMUNITYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPCOMMUNITYMIB.EntityData.Children = types.NewOrderedMap()
    sNMPCOMMUNITYMIB.EntityData.Children.Append("snmpCommunityTable", types.YChild{"SnmpCommunityTable", &sNMPCOMMUNITYMIB.SnmpCommunityTable})
    sNMPCOMMUNITYMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPCOMMUNITYMIB.EntityData.YListKeys = []string {}

    return &(sNMPCOMMUNITYMIB.EntityData)
}

// SNMPCOMMUNITYMIB_SnmpCommunityTable
type SNMPCOMMUNITYMIB_SnmpCommunityTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // SNMPCOMMUNITYMIB_SnmpCommunityTable_SnmpCommunityEntry.
    SnmpCommunityEntry []*SNMPCOMMUNITYMIB_SnmpCommunityTable_SnmpCommunityEntry
}

func (snmpCommunityTable *SNMPCOMMUNITYMIB_SnmpCommunityTable) GetEntityData() *types.CommonEntityData {
    snmpCommunityTable.EntityData.YFilter = snmpCommunityTable.YFilter
    snmpCommunityTable.EntityData.YangName = "snmpCommunityTable"
    snmpCommunityTable.EntityData.BundleName = "cisco_ios_xr"
    snmpCommunityTable.EntityData.ParentYangName = "SNMP-COMMUNITY-MIB"
    snmpCommunityTable.EntityData.SegmentPath = "snmpCommunityTable"
    snmpCommunityTable.EntityData.AbsolutePath = "SNMP-COMMUNITY-MIB:SNMP-COMMUNITY-MIB/" + snmpCommunityTable.EntityData.SegmentPath
    snmpCommunityTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpCommunityTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpCommunityTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpCommunityTable.EntityData.Children = types.NewOrderedMap()
    snmpCommunityTable.EntityData.Children.Append("snmpCommunityEntry", types.YChild{"SnmpCommunityEntry", nil})
    for i := range snmpCommunityTable.SnmpCommunityEntry {
        snmpCommunityTable.EntityData.Children.Append(types.GetSegmentPath(snmpCommunityTable.SnmpCommunityEntry[i]), types.YChild{"SnmpCommunityEntry", snmpCommunityTable.SnmpCommunityEntry[i]})
    }
    snmpCommunityTable.EntityData.Leafs = types.NewOrderedMap()

    snmpCommunityTable.EntityData.YListKeys = []string {}

    return &(snmpCommunityTable.EntityData)
}

// SNMPCOMMUNITYMIB_SnmpCommunityTable_SnmpCommunityEntry
type SNMPCOMMUNITYMIB_SnmpCommunityTable_SnmpCommunityEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    SnmpCommunityIndex interface{}

    // The type is string. This attribute is mandatory.
    SnmpCommunityName interface{}

    // The type is string with length: 1..32. This attribute is mandatory.
    SnmpCommunitySecurityName interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'. This attribute is mandatory.
    SnmpCommunityContextEngineID interface{}

    // The type is string with length: 0..32.
    SnmpCommunityContextName interface{}

    // The type is string with length: 0..255.
    SnmpCommunityTransportTag interface{}

    // The type is StorageType. The default value is permanent.
    SnmpCommunityStorageType interface{}
}

func (snmpCommunityEntry *SNMPCOMMUNITYMIB_SnmpCommunityTable_SnmpCommunityEntry) GetEntityData() *types.CommonEntityData {
    snmpCommunityEntry.EntityData.YFilter = snmpCommunityEntry.YFilter
    snmpCommunityEntry.EntityData.YangName = "snmpCommunityEntry"
    snmpCommunityEntry.EntityData.BundleName = "cisco_ios_xr"
    snmpCommunityEntry.EntityData.ParentYangName = "snmpCommunityTable"
    snmpCommunityEntry.EntityData.SegmentPath = "snmpCommunityEntry" + types.AddKeyToken(snmpCommunityEntry.SnmpCommunityIndex, "snmpCommunityIndex")
    snmpCommunityEntry.EntityData.AbsolutePath = "SNMP-COMMUNITY-MIB:SNMP-COMMUNITY-MIB/snmpCommunityTable/" + snmpCommunityEntry.EntityData.SegmentPath
    snmpCommunityEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpCommunityEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpCommunityEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpCommunityEntry.EntityData.Children = types.NewOrderedMap()
    snmpCommunityEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpCommunityEntry.EntityData.Leafs.Append("snmpCommunityIndex", types.YLeaf{"SnmpCommunityIndex", snmpCommunityEntry.SnmpCommunityIndex})
    snmpCommunityEntry.EntityData.Leafs.Append("snmpCommunityName", types.YLeaf{"SnmpCommunityName", snmpCommunityEntry.SnmpCommunityName})
    snmpCommunityEntry.EntityData.Leafs.Append("snmpCommunitySecurityName", types.YLeaf{"SnmpCommunitySecurityName", snmpCommunityEntry.SnmpCommunitySecurityName})
    snmpCommunityEntry.EntityData.Leafs.Append("snmpCommunityContextEngineID", types.YLeaf{"SnmpCommunityContextEngineID", snmpCommunityEntry.SnmpCommunityContextEngineID})
    snmpCommunityEntry.EntityData.Leafs.Append("snmpCommunityContextName", types.YLeaf{"SnmpCommunityContextName", snmpCommunityEntry.SnmpCommunityContextName})
    snmpCommunityEntry.EntityData.Leafs.Append("snmpCommunityTransportTag", types.YLeaf{"SnmpCommunityTransportTag", snmpCommunityEntry.SnmpCommunityTransportTag})
    snmpCommunityEntry.EntityData.Leafs.Append("snmpCommunityStorageType", types.YLeaf{"SnmpCommunityStorageType", snmpCommunityEntry.SnmpCommunityStorageType})

    snmpCommunityEntry.EntityData.YListKeys = []string {"SnmpCommunityIndex"}

    return &(snmpCommunityEntry.EntityData)
}

