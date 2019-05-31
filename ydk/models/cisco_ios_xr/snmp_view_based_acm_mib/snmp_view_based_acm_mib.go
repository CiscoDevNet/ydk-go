package snmp_view_based_acm_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_view_based_acm_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMP-VIEW-BASED-ACM-MIB/200210160000Z SNMP-VIEW-BASED-ACM-MIB}", reflect.TypeOf(SNMPVIEWBASEDACMMIB{}))
    ydk.RegisterEntity("SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB", reflect.TypeOf(SNMPVIEWBASEDACMMIB{}))
}

// VacmAccessContextMatchType
type VacmAccessContextMatchType string

const (
    VacmAccessContextMatchType_exact VacmAccessContextMatchType = "exact"

    VacmAccessContextMatchType_prefix VacmAccessContextMatchType = "prefix"
)

// VacmViewTreeFamilyTypeType
type VacmViewTreeFamilyTypeType string

const (
    VacmViewTreeFamilyTypeType_included VacmViewTreeFamilyTypeType = "included"

    VacmViewTreeFamilyTypeType_excluded VacmViewTreeFamilyTypeType = "excluded"
)

// SNMPVIEWBASEDACMMIB
type SNMPVIEWBASEDACMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    VacmSecurityToGroupTable SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable

    
    VacmAccessTable SNMPVIEWBASEDACMMIB_VacmAccessTable

    
    VacmViewTreeFamilyTable SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable
}

func (sNMPVIEWBASEDACMMIB *SNMPVIEWBASEDACMMIB) GetEntityData() *types.CommonEntityData {
    sNMPVIEWBASEDACMMIB.EntityData.YFilter = sNMPVIEWBASEDACMMIB.YFilter
    sNMPVIEWBASEDACMMIB.EntityData.YangName = "SNMP-VIEW-BASED-ACM-MIB"
    sNMPVIEWBASEDACMMIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPVIEWBASEDACMMIB.EntityData.ParentYangName = "SNMP-VIEW-BASED-ACM-MIB"
    sNMPVIEWBASEDACMMIB.EntityData.SegmentPath = "SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB"
    sNMPVIEWBASEDACMMIB.EntityData.AbsolutePath = sNMPVIEWBASEDACMMIB.EntityData.SegmentPath
    sNMPVIEWBASEDACMMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPVIEWBASEDACMMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPVIEWBASEDACMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPVIEWBASEDACMMIB.EntityData.Children = types.NewOrderedMap()
    sNMPVIEWBASEDACMMIB.EntityData.Children.Append("vacmSecurityToGroupTable", types.YChild{"VacmSecurityToGroupTable", &sNMPVIEWBASEDACMMIB.VacmSecurityToGroupTable})
    sNMPVIEWBASEDACMMIB.EntityData.Children.Append("vacmAccessTable", types.YChild{"VacmAccessTable", &sNMPVIEWBASEDACMMIB.VacmAccessTable})
    sNMPVIEWBASEDACMMIB.EntityData.Children.Append("vacmViewTreeFamilyTable", types.YChild{"VacmViewTreeFamilyTable", &sNMPVIEWBASEDACMMIB.VacmViewTreeFamilyTable})
    sNMPVIEWBASEDACMMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPVIEWBASEDACMMIB.EntityData.YListKeys = []string {}

    return &(sNMPVIEWBASEDACMMIB.EntityData)
}

// SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable
type SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable_VacmSecurityToGroupEntry.
    VacmSecurityToGroupEntry []*SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable_VacmSecurityToGroupEntry
}

func (vacmSecurityToGroupTable *SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable) GetEntityData() *types.CommonEntityData {
    vacmSecurityToGroupTable.EntityData.YFilter = vacmSecurityToGroupTable.YFilter
    vacmSecurityToGroupTable.EntityData.YangName = "vacmSecurityToGroupTable"
    vacmSecurityToGroupTable.EntityData.BundleName = "cisco_ios_xr"
    vacmSecurityToGroupTable.EntityData.ParentYangName = "SNMP-VIEW-BASED-ACM-MIB"
    vacmSecurityToGroupTable.EntityData.SegmentPath = "vacmSecurityToGroupTable"
    vacmSecurityToGroupTable.EntityData.AbsolutePath = "SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB/" + vacmSecurityToGroupTable.EntityData.SegmentPath
    vacmSecurityToGroupTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vacmSecurityToGroupTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vacmSecurityToGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vacmSecurityToGroupTable.EntityData.Children = types.NewOrderedMap()
    vacmSecurityToGroupTable.EntityData.Children.Append("vacmSecurityToGroupEntry", types.YChild{"VacmSecurityToGroupEntry", nil})
    for i := range vacmSecurityToGroupTable.VacmSecurityToGroupEntry {
        vacmSecurityToGroupTable.EntityData.Children.Append(types.GetSegmentPath(vacmSecurityToGroupTable.VacmSecurityToGroupEntry[i]), types.YChild{"VacmSecurityToGroupEntry", vacmSecurityToGroupTable.VacmSecurityToGroupEntry[i]})
    }
    vacmSecurityToGroupTable.EntityData.Leafs = types.NewOrderedMap()

    vacmSecurityToGroupTable.EntityData.YListKeys = []string {}

    return &(vacmSecurityToGroupTable.EntityData)
}

// SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable_VacmSecurityToGroupEntry
type SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable_VacmSecurityToGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    VacmSecurityModel interface{}

    // This attribute is a key. The type is string with length: 1..32.
    VacmSecurityName interface{}

    // The type is string with length: 1..32. This attribute is mandatory.
    VacmGroupName interface{}

    // The type is StorageType. The default value is nonVolatile.
    VacmSecurityToGroupStorageType interface{}
}

func (vacmSecurityToGroupEntry *SNMPVIEWBASEDACMMIB_VacmSecurityToGroupTable_VacmSecurityToGroupEntry) GetEntityData() *types.CommonEntityData {
    vacmSecurityToGroupEntry.EntityData.YFilter = vacmSecurityToGroupEntry.YFilter
    vacmSecurityToGroupEntry.EntityData.YangName = "vacmSecurityToGroupEntry"
    vacmSecurityToGroupEntry.EntityData.BundleName = "cisco_ios_xr"
    vacmSecurityToGroupEntry.EntityData.ParentYangName = "vacmSecurityToGroupTable"
    vacmSecurityToGroupEntry.EntityData.SegmentPath = "vacmSecurityToGroupEntry" + types.AddKeyToken(vacmSecurityToGroupEntry.VacmSecurityModel, "vacmSecurityModel") + types.AddKeyToken(vacmSecurityToGroupEntry.VacmSecurityName, "vacmSecurityName")
    vacmSecurityToGroupEntry.EntityData.AbsolutePath = "SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB/vacmSecurityToGroupTable/" + vacmSecurityToGroupEntry.EntityData.SegmentPath
    vacmSecurityToGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vacmSecurityToGroupEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vacmSecurityToGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vacmSecurityToGroupEntry.EntityData.Children = types.NewOrderedMap()
    vacmSecurityToGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    vacmSecurityToGroupEntry.EntityData.Leafs.Append("vacmSecurityModel", types.YLeaf{"VacmSecurityModel", vacmSecurityToGroupEntry.VacmSecurityModel})
    vacmSecurityToGroupEntry.EntityData.Leafs.Append("vacmSecurityName", types.YLeaf{"VacmSecurityName", vacmSecurityToGroupEntry.VacmSecurityName})
    vacmSecurityToGroupEntry.EntityData.Leafs.Append("vacmGroupName", types.YLeaf{"VacmGroupName", vacmSecurityToGroupEntry.VacmGroupName})
    vacmSecurityToGroupEntry.EntityData.Leafs.Append("vacmSecurityToGroupStorageType", types.YLeaf{"VacmSecurityToGroupStorageType", vacmSecurityToGroupEntry.VacmSecurityToGroupStorageType})

    vacmSecurityToGroupEntry.EntityData.YListKeys = []string {"VacmSecurityModel", "VacmSecurityName"}

    return &(vacmSecurityToGroupEntry.EntityData)
}

// SNMPVIEWBASEDACMMIB_VacmAccessTable
type SNMPVIEWBASEDACMMIB_VacmAccessTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SNMPVIEWBASEDACMMIB_VacmAccessTable_VacmAccessEntry.
    VacmAccessEntry []*SNMPVIEWBASEDACMMIB_VacmAccessTable_VacmAccessEntry
}

func (vacmAccessTable *SNMPVIEWBASEDACMMIB_VacmAccessTable) GetEntityData() *types.CommonEntityData {
    vacmAccessTable.EntityData.YFilter = vacmAccessTable.YFilter
    vacmAccessTable.EntityData.YangName = "vacmAccessTable"
    vacmAccessTable.EntityData.BundleName = "cisco_ios_xr"
    vacmAccessTable.EntityData.ParentYangName = "SNMP-VIEW-BASED-ACM-MIB"
    vacmAccessTable.EntityData.SegmentPath = "vacmAccessTable"
    vacmAccessTable.EntityData.AbsolutePath = "SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB/" + vacmAccessTable.EntityData.SegmentPath
    vacmAccessTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vacmAccessTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vacmAccessTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vacmAccessTable.EntityData.Children = types.NewOrderedMap()
    vacmAccessTable.EntityData.Children.Append("vacmAccessEntry", types.YChild{"VacmAccessEntry", nil})
    for i := range vacmAccessTable.VacmAccessEntry {
        vacmAccessTable.EntityData.Children.Append(types.GetSegmentPath(vacmAccessTable.VacmAccessEntry[i]), types.YChild{"VacmAccessEntry", vacmAccessTable.VacmAccessEntry[i]})
    }
    vacmAccessTable.EntityData.Leafs = types.NewOrderedMap()

    vacmAccessTable.EntityData.YListKeys = []string {}

    return &(vacmAccessTable.EntityData)
}

// SNMPVIEWBASEDACMMIB_VacmAccessTable_VacmAccessEntry
type SNMPVIEWBASEDACMMIB_VacmAccessTable_VacmAccessEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    VacmGroupName interface{}

    // This attribute is a key. The type is string with length: 0..32.
    VacmAccessContextPrefix interface{}

    // This attribute is a key. The type is interface{} with range: 0..2147483647.
    VacmAccessSecurityModel interface{}

    // This attribute is a key. The type is SnmpSecurityLevel.
    VacmAccessSecurityLevel interface{}

    // The type is VacmAccessContextMatchType. The default value is exact.
    VacmAccessContextMatch interface{}

    // The type is string with length: 0..32.
    VacmAccessReadViewName interface{}

    // The type is string with length: 0..32.
    VacmAccessWriteViewName interface{}

    // The type is string with length: 0..32.
    VacmAccessNotifyViewName interface{}

    // The type is StorageType. The default value is nonVolatile.
    VacmAccessStorageType interface{}
}

func (vacmAccessEntry *SNMPVIEWBASEDACMMIB_VacmAccessTable_VacmAccessEntry) GetEntityData() *types.CommonEntityData {
    vacmAccessEntry.EntityData.YFilter = vacmAccessEntry.YFilter
    vacmAccessEntry.EntityData.YangName = "vacmAccessEntry"
    vacmAccessEntry.EntityData.BundleName = "cisco_ios_xr"
    vacmAccessEntry.EntityData.ParentYangName = "vacmAccessTable"
    vacmAccessEntry.EntityData.SegmentPath = "vacmAccessEntry" + types.AddKeyToken(vacmAccessEntry.VacmGroupName, "vacmGroupName") + types.AddKeyToken(vacmAccessEntry.VacmAccessContextPrefix, "vacmAccessContextPrefix") + types.AddKeyToken(vacmAccessEntry.VacmAccessSecurityModel, "vacmAccessSecurityModel") + types.AddKeyToken(vacmAccessEntry.VacmAccessSecurityLevel, "vacmAccessSecurityLevel")
    vacmAccessEntry.EntityData.AbsolutePath = "SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB/vacmAccessTable/" + vacmAccessEntry.EntityData.SegmentPath
    vacmAccessEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vacmAccessEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vacmAccessEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vacmAccessEntry.EntityData.Children = types.NewOrderedMap()
    vacmAccessEntry.EntityData.Leafs = types.NewOrderedMap()
    vacmAccessEntry.EntityData.Leafs.Append("vacmGroupName", types.YLeaf{"VacmGroupName", vacmAccessEntry.VacmGroupName})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessContextPrefix", types.YLeaf{"VacmAccessContextPrefix", vacmAccessEntry.VacmAccessContextPrefix})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessSecurityModel", types.YLeaf{"VacmAccessSecurityModel", vacmAccessEntry.VacmAccessSecurityModel})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessSecurityLevel", types.YLeaf{"VacmAccessSecurityLevel", vacmAccessEntry.VacmAccessSecurityLevel})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessContextMatch", types.YLeaf{"VacmAccessContextMatch", vacmAccessEntry.VacmAccessContextMatch})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessReadViewName", types.YLeaf{"VacmAccessReadViewName", vacmAccessEntry.VacmAccessReadViewName})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessWriteViewName", types.YLeaf{"VacmAccessWriteViewName", vacmAccessEntry.VacmAccessWriteViewName})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessNotifyViewName", types.YLeaf{"VacmAccessNotifyViewName", vacmAccessEntry.VacmAccessNotifyViewName})
    vacmAccessEntry.EntityData.Leafs.Append("vacmAccessStorageType", types.YLeaf{"VacmAccessStorageType", vacmAccessEntry.VacmAccessStorageType})

    vacmAccessEntry.EntityData.YListKeys = []string {"VacmGroupName", "VacmAccessContextPrefix", "VacmAccessSecurityModel", "VacmAccessSecurityLevel"}

    return &(vacmAccessEntry.EntityData)
}

// SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable
type SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable_VacmViewTreeFamilyEntry.
    VacmViewTreeFamilyEntry []*SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable_VacmViewTreeFamilyEntry
}

func (vacmViewTreeFamilyTable *SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable) GetEntityData() *types.CommonEntityData {
    vacmViewTreeFamilyTable.EntityData.YFilter = vacmViewTreeFamilyTable.YFilter
    vacmViewTreeFamilyTable.EntityData.YangName = "vacmViewTreeFamilyTable"
    vacmViewTreeFamilyTable.EntityData.BundleName = "cisco_ios_xr"
    vacmViewTreeFamilyTable.EntityData.ParentYangName = "SNMP-VIEW-BASED-ACM-MIB"
    vacmViewTreeFamilyTable.EntityData.SegmentPath = "vacmViewTreeFamilyTable"
    vacmViewTreeFamilyTable.EntityData.AbsolutePath = "SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB/" + vacmViewTreeFamilyTable.EntityData.SegmentPath
    vacmViewTreeFamilyTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vacmViewTreeFamilyTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vacmViewTreeFamilyTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vacmViewTreeFamilyTable.EntityData.Children = types.NewOrderedMap()
    vacmViewTreeFamilyTable.EntityData.Children.Append("vacmViewTreeFamilyEntry", types.YChild{"VacmViewTreeFamilyEntry", nil})
    for i := range vacmViewTreeFamilyTable.VacmViewTreeFamilyEntry {
        vacmViewTreeFamilyTable.EntityData.Children.Append(types.GetSegmentPath(vacmViewTreeFamilyTable.VacmViewTreeFamilyEntry[i]), types.YChild{"VacmViewTreeFamilyEntry", vacmViewTreeFamilyTable.VacmViewTreeFamilyEntry[i]})
    }
    vacmViewTreeFamilyTable.EntityData.Leafs = types.NewOrderedMap()

    vacmViewTreeFamilyTable.EntityData.YListKeys = []string {}

    return &(vacmViewTreeFamilyTable.EntityData)
}

// SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable_VacmViewTreeFamilyEntry
type SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable_VacmViewTreeFamilyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    VacmViewTreeFamilyViewName interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    VacmViewTreeFamilySubtree interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    VacmViewTreeFamilyMask interface{}

    // The type is VacmViewTreeFamilyTypeType. The default value is included.
    VacmViewTreeFamilyType interface{}

    // The type is StorageType. The default value is nonVolatile.
    VacmViewTreeFamilyStorageType interface{}
}

func (vacmViewTreeFamilyEntry *SNMPVIEWBASEDACMMIB_VacmViewTreeFamilyTable_VacmViewTreeFamilyEntry) GetEntityData() *types.CommonEntityData {
    vacmViewTreeFamilyEntry.EntityData.YFilter = vacmViewTreeFamilyEntry.YFilter
    vacmViewTreeFamilyEntry.EntityData.YangName = "vacmViewTreeFamilyEntry"
    vacmViewTreeFamilyEntry.EntityData.BundleName = "cisco_ios_xr"
    vacmViewTreeFamilyEntry.EntityData.ParentYangName = "vacmViewTreeFamilyTable"
    vacmViewTreeFamilyEntry.EntityData.SegmentPath = "vacmViewTreeFamilyEntry" + types.AddKeyToken(vacmViewTreeFamilyEntry.VacmViewTreeFamilyViewName, "vacmViewTreeFamilyViewName") + types.AddKeyToken(vacmViewTreeFamilyEntry.VacmViewTreeFamilySubtree, "vacmViewTreeFamilySubtree")
    vacmViewTreeFamilyEntry.EntityData.AbsolutePath = "SNMP-VIEW-BASED-ACM-MIB:SNMP-VIEW-BASED-ACM-MIB/vacmViewTreeFamilyTable/" + vacmViewTreeFamilyEntry.EntityData.SegmentPath
    vacmViewTreeFamilyEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vacmViewTreeFamilyEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vacmViewTreeFamilyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vacmViewTreeFamilyEntry.EntityData.Children = types.NewOrderedMap()
    vacmViewTreeFamilyEntry.EntityData.Leafs = types.NewOrderedMap()
    vacmViewTreeFamilyEntry.EntityData.Leafs.Append("vacmViewTreeFamilyViewName", types.YLeaf{"VacmViewTreeFamilyViewName", vacmViewTreeFamilyEntry.VacmViewTreeFamilyViewName})
    vacmViewTreeFamilyEntry.EntityData.Leafs.Append("vacmViewTreeFamilySubtree", types.YLeaf{"VacmViewTreeFamilySubtree", vacmViewTreeFamilyEntry.VacmViewTreeFamilySubtree})
    vacmViewTreeFamilyEntry.EntityData.Leafs.Append("vacmViewTreeFamilyMask", types.YLeaf{"VacmViewTreeFamilyMask", vacmViewTreeFamilyEntry.VacmViewTreeFamilyMask})
    vacmViewTreeFamilyEntry.EntityData.Leafs.Append("vacmViewTreeFamilyType", types.YLeaf{"VacmViewTreeFamilyType", vacmViewTreeFamilyEntry.VacmViewTreeFamilyType})
    vacmViewTreeFamilyEntry.EntityData.Leafs.Append("vacmViewTreeFamilyStorageType", types.YLeaf{"VacmViewTreeFamilyStorageType", vacmViewTreeFamilyEntry.VacmViewTreeFamilyStorageType})

    vacmViewTreeFamilyEntry.EntityData.YListKeys = []string {"VacmViewTreeFamilyViewName", "VacmViewTreeFamilySubtree"}

    return &(vacmViewTreeFamilyEntry.EntityData)
}

