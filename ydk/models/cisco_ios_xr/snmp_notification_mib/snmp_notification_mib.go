package snmp_notification_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_notification_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMP-NOTIFICATION-MIB/200210140000Z SNMP-NOTIFICATION-MIB}", reflect.TypeOf(SNMPNOTIFICATIONMIB{}))
    ydk.RegisterEntity("SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB", reflect.TypeOf(SNMPNOTIFICATIONMIB{}))
}

// SnmpNotifyTypeType
type SnmpNotifyTypeType string

const (
    SnmpNotifyTypeType_trap SnmpNotifyTypeType = "trap"

    SnmpNotifyTypeType_inform SnmpNotifyTypeType = "inform"
)

// SnmpNotifyFilterTypeType
type SnmpNotifyFilterTypeType string

const (
    SnmpNotifyFilterTypeType_included SnmpNotifyFilterTypeType = "included"

    SnmpNotifyFilterTypeType_excluded SnmpNotifyFilterTypeType = "excluded"
)

// SNMPNOTIFICATIONMIB
type SNMPNOTIFICATIONMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SnmpNotifyTable SNMPNOTIFICATIONMIB_SnmpNotifyTable

    
    SnmpNotifyFilterProfileTable SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable

    
    SnmpNotifyFilterTable SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable
}

func (sNMPNOTIFICATIONMIB *SNMPNOTIFICATIONMIB) GetEntityData() *types.CommonEntityData {
    sNMPNOTIFICATIONMIB.EntityData.YFilter = sNMPNOTIFICATIONMIB.YFilter
    sNMPNOTIFICATIONMIB.EntityData.YangName = "SNMP-NOTIFICATION-MIB"
    sNMPNOTIFICATIONMIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPNOTIFICATIONMIB.EntityData.ParentYangName = "SNMP-NOTIFICATION-MIB"
    sNMPNOTIFICATIONMIB.EntityData.SegmentPath = "SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB"
    sNMPNOTIFICATIONMIB.EntityData.AbsolutePath = sNMPNOTIFICATIONMIB.EntityData.SegmentPath
    sNMPNOTIFICATIONMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPNOTIFICATIONMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPNOTIFICATIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPNOTIFICATIONMIB.EntityData.Children = types.NewOrderedMap()
    sNMPNOTIFICATIONMIB.EntityData.Children.Append("snmpNotifyTable", types.YChild{"SnmpNotifyTable", &sNMPNOTIFICATIONMIB.SnmpNotifyTable})
    sNMPNOTIFICATIONMIB.EntityData.Children.Append("snmpNotifyFilterProfileTable", types.YChild{"SnmpNotifyFilterProfileTable", &sNMPNOTIFICATIONMIB.SnmpNotifyFilterProfileTable})
    sNMPNOTIFICATIONMIB.EntityData.Children.Append("snmpNotifyFilterTable", types.YChild{"SnmpNotifyFilterTable", &sNMPNOTIFICATIONMIB.SnmpNotifyFilterTable})
    sNMPNOTIFICATIONMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPNOTIFICATIONMIB.EntityData.YListKeys = []string {}

    return &(sNMPNOTIFICATIONMIB.EntityData)
}

// SNMPNOTIFICATIONMIB_SnmpNotifyTable
type SNMPNOTIFICATIONMIB_SnmpNotifyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SNMPNOTIFICATIONMIB_SnmpNotifyTable_SnmpNotifyEntry.
    SnmpNotifyEntry []*SNMPNOTIFICATIONMIB_SnmpNotifyTable_SnmpNotifyEntry
}

func (snmpNotifyTable *SNMPNOTIFICATIONMIB_SnmpNotifyTable) GetEntityData() *types.CommonEntityData {
    snmpNotifyTable.EntityData.YFilter = snmpNotifyTable.YFilter
    snmpNotifyTable.EntityData.YangName = "snmpNotifyTable"
    snmpNotifyTable.EntityData.BundleName = "cisco_ios_xr"
    snmpNotifyTable.EntityData.ParentYangName = "SNMP-NOTIFICATION-MIB"
    snmpNotifyTable.EntityData.SegmentPath = "snmpNotifyTable"
    snmpNotifyTable.EntityData.AbsolutePath = "SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB/" + snmpNotifyTable.EntityData.SegmentPath
    snmpNotifyTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpNotifyTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpNotifyTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpNotifyTable.EntityData.Children = types.NewOrderedMap()
    snmpNotifyTable.EntityData.Children.Append("snmpNotifyEntry", types.YChild{"SnmpNotifyEntry", nil})
    for i := range snmpNotifyTable.SnmpNotifyEntry {
        snmpNotifyTable.EntityData.Children.Append(types.GetSegmentPath(snmpNotifyTable.SnmpNotifyEntry[i]), types.YChild{"SnmpNotifyEntry", snmpNotifyTable.SnmpNotifyEntry[i]})
    }
    snmpNotifyTable.EntityData.Leafs = types.NewOrderedMap()

    snmpNotifyTable.EntityData.YListKeys = []string {}

    return &(snmpNotifyTable.EntityData)
}

// SNMPNOTIFICATIONMIB_SnmpNotifyTable_SnmpNotifyEntry
type SNMPNOTIFICATIONMIB_SnmpNotifyTable_SnmpNotifyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    SnmpNotifyName interface{}

    // The type is string with length: 0..255.
    SnmpNotifyTag interface{}

    // The type is SnmpNotifyTypeType. The default value is trap.
    SnmpNotifyType interface{}

    // The type is StorageType. The default value is nonVolatile.
    SnmpNotifyStorageType interface{}
}

func (snmpNotifyEntry *SNMPNOTIFICATIONMIB_SnmpNotifyTable_SnmpNotifyEntry) GetEntityData() *types.CommonEntityData {
    snmpNotifyEntry.EntityData.YFilter = snmpNotifyEntry.YFilter
    snmpNotifyEntry.EntityData.YangName = "snmpNotifyEntry"
    snmpNotifyEntry.EntityData.BundleName = "cisco_ios_xr"
    snmpNotifyEntry.EntityData.ParentYangName = "snmpNotifyTable"
    snmpNotifyEntry.EntityData.SegmentPath = "snmpNotifyEntry" + types.AddKeyToken(snmpNotifyEntry.SnmpNotifyName, "snmpNotifyName")
    snmpNotifyEntry.EntityData.AbsolutePath = "SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB/snmpNotifyTable/" + snmpNotifyEntry.EntityData.SegmentPath
    snmpNotifyEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpNotifyEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpNotifyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpNotifyEntry.EntityData.Children = types.NewOrderedMap()
    snmpNotifyEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpNotifyEntry.EntityData.Leafs.Append("snmpNotifyName", types.YLeaf{"SnmpNotifyName", snmpNotifyEntry.SnmpNotifyName})
    snmpNotifyEntry.EntityData.Leafs.Append("snmpNotifyTag", types.YLeaf{"SnmpNotifyTag", snmpNotifyEntry.SnmpNotifyTag})
    snmpNotifyEntry.EntityData.Leafs.Append("snmpNotifyType", types.YLeaf{"SnmpNotifyType", snmpNotifyEntry.SnmpNotifyType})
    snmpNotifyEntry.EntityData.Leafs.Append("snmpNotifyStorageType", types.YLeaf{"SnmpNotifyStorageType", snmpNotifyEntry.SnmpNotifyStorageType})

    snmpNotifyEntry.EntityData.YListKeys = []string {"SnmpNotifyName"}

    return &(snmpNotifyEntry.EntityData)
}

// SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable
// This type is a presence type.
type SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The type is slice of
    // SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable_SnmpNotifyFilterProfileEntry.
    SnmpNotifyFilterProfileEntry []*SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable_SnmpNotifyFilterProfileEntry
}

func (snmpNotifyFilterProfileTable *SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable) GetEntityData() *types.CommonEntityData {
    snmpNotifyFilterProfileTable.EntityData.YFilter = snmpNotifyFilterProfileTable.YFilter
    snmpNotifyFilterProfileTable.EntityData.YangName = "snmpNotifyFilterProfileTable"
    snmpNotifyFilterProfileTable.EntityData.BundleName = "cisco_ios_xr"
    snmpNotifyFilterProfileTable.EntityData.ParentYangName = "SNMP-NOTIFICATION-MIB"
    snmpNotifyFilterProfileTable.EntityData.SegmentPath = "snmpNotifyFilterProfileTable"
    snmpNotifyFilterProfileTable.EntityData.AbsolutePath = "SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB/" + snmpNotifyFilterProfileTable.EntityData.SegmentPath
    snmpNotifyFilterProfileTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpNotifyFilterProfileTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpNotifyFilterProfileTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpNotifyFilterProfileTable.EntityData.Children = types.NewOrderedMap()
    snmpNotifyFilterProfileTable.EntityData.Children.Append("snmpNotifyFilterProfileEntry", types.YChild{"SnmpNotifyFilterProfileEntry", nil})
    for i := range snmpNotifyFilterProfileTable.SnmpNotifyFilterProfileEntry {
        snmpNotifyFilterProfileTable.EntityData.Children.Append(types.GetSegmentPath(snmpNotifyFilterProfileTable.SnmpNotifyFilterProfileEntry[i]), types.YChild{"SnmpNotifyFilterProfileEntry", snmpNotifyFilterProfileTable.SnmpNotifyFilterProfileEntry[i]})
    }
    snmpNotifyFilterProfileTable.EntityData.Leafs = types.NewOrderedMap()

    snmpNotifyFilterProfileTable.EntityData.YListKeys = []string {}

    return &(snmpNotifyFilterProfileTable.EntityData)
}

// SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable_SnmpNotifyFilterProfileEntry
type SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable_SnmpNotifyFilterProfileEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    SnmpTargetParamsName interface{}

    // The type is string with length: 1..32.
    SnmpNotifyFilterProfileName interface{}

    // The type is StorageType. The default value is nonVolatile.
    SnmpNotifyFilterProfileStorType interface{}
}

func (snmpNotifyFilterProfileEntry *SNMPNOTIFICATIONMIB_SnmpNotifyFilterProfileTable_SnmpNotifyFilterProfileEntry) GetEntityData() *types.CommonEntityData {
    snmpNotifyFilterProfileEntry.EntityData.YFilter = snmpNotifyFilterProfileEntry.YFilter
    snmpNotifyFilterProfileEntry.EntityData.YangName = "snmpNotifyFilterProfileEntry"
    snmpNotifyFilterProfileEntry.EntityData.BundleName = "cisco_ios_xr"
    snmpNotifyFilterProfileEntry.EntityData.ParentYangName = "snmpNotifyFilterProfileTable"
    snmpNotifyFilterProfileEntry.EntityData.SegmentPath = "snmpNotifyFilterProfileEntry" + types.AddKeyToken(snmpNotifyFilterProfileEntry.SnmpTargetParamsName, "snmpTargetParamsName")
    snmpNotifyFilterProfileEntry.EntityData.AbsolutePath = "SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB/snmpNotifyFilterProfileTable/" + snmpNotifyFilterProfileEntry.EntityData.SegmentPath
    snmpNotifyFilterProfileEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpNotifyFilterProfileEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpNotifyFilterProfileEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpNotifyFilterProfileEntry.EntityData.Children = types.NewOrderedMap()
    snmpNotifyFilterProfileEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpNotifyFilterProfileEntry.EntityData.Leafs.Append("snmpTargetParamsName", types.YLeaf{"SnmpTargetParamsName", snmpNotifyFilterProfileEntry.SnmpTargetParamsName})
    snmpNotifyFilterProfileEntry.EntityData.Leafs.Append("snmpNotifyFilterProfileName", types.YLeaf{"SnmpNotifyFilterProfileName", snmpNotifyFilterProfileEntry.SnmpNotifyFilterProfileName})
    snmpNotifyFilterProfileEntry.EntityData.Leafs.Append("snmpNotifyFilterProfileStorType", types.YLeaf{"SnmpNotifyFilterProfileStorType", snmpNotifyFilterProfileEntry.SnmpNotifyFilterProfileStorType})

    snmpNotifyFilterProfileEntry.EntityData.YListKeys = []string {"SnmpTargetParamsName"}

    return &(snmpNotifyFilterProfileEntry.EntityData)
}

// SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable
// This type is a presence type.
type SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The type is slice of
    // SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable_SnmpNotifyFilterEntry.
    SnmpNotifyFilterEntry []*SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable_SnmpNotifyFilterEntry
}

func (snmpNotifyFilterTable *SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable) GetEntityData() *types.CommonEntityData {
    snmpNotifyFilterTable.EntityData.YFilter = snmpNotifyFilterTable.YFilter
    snmpNotifyFilterTable.EntityData.YangName = "snmpNotifyFilterTable"
    snmpNotifyFilterTable.EntityData.BundleName = "cisco_ios_xr"
    snmpNotifyFilterTable.EntityData.ParentYangName = "SNMP-NOTIFICATION-MIB"
    snmpNotifyFilterTable.EntityData.SegmentPath = "snmpNotifyFilterTable"
    snmpNotifyFilterTable.EntityData.AbsolutePath = "SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB/" + snmpNotifyFilterTable.EntityData.SegmentPath
    snmpNotifyFilterTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpNotifyFilterTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpNotifyFilterTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpNotifyFilterTable.EntityData.Children = types.NewOrderedMap()
    snmpNotifyFilterTable.EntityData.Children.Append("snmpNotifyFilterEntry", types.YChild{"SnmpNotifyFilterEntry", nil})
    for i := range snmpNotifyFilterTable.SnmpNotifyFilterEntry {
        snmpNotifyFilterTable.EntityData.Children.Append(types.GetSegmentPath(snmpNotifyFilterTable.SnmpNotifyFilterEntry[i]), types.YChild{"SnmpNotifyFilterEntry", snmpNotifyFilterTable.SnmpNotifyFilterEntry[i]})
    }
    snmpNotifyFilterTable.EntityData.Leafs = types.NewOrderedMap()

    snmpNotifyFilterTable.EntityData.YListKeys = []string {}

    return &(snmpNotifyFilterTable.EntityData)
}

// SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable_SnmpNotifyFilterEntry
type SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable_SnmpNotifyFilterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..32.
    SnmpNotifyFilterProfileName interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    SnmpNotifyFilterSubtree interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    SnmpNotifyFilterMask interface{}

    // The type is SnmpNotifyFilterTypeType. The default value is included.
    SnmpNotifyFilterType interface{}

    // The type is StorageType. The default value is nonVolatile.
    SnmpNotifyFilterStorageType interface{}
}

func (snmpNotifyFilterEntry *SNMPNOTIFICATIONMIB_SnmpNotifyFilterTable_SnmpNotifyFilterEntry) GetEntityData() *types.CommonEntityData {
    snmpNotifyFilterEntry.EntityData.YFilter = snmpNotifyFilterEntry.YFilter
    snmpNotifyFilterEntry.EntityData.YangName = "snmpNotifyFilterEntry"
    snmpNotifyFilterEntry.EntityData.BundleName = "cisco_ios_xr"
    snmpNotifyFilterEntry.EntityData.ParentYangName = "snmpNotifyFilterTable"
    snmpNotifyFilterEntry.EntityData.SegmentPath = "snmpNotifyFilterEntry" + types.AddKeyToken(snmpNotifyFilterEntry.SnmpNotifyFilterProfileName, "snmpNotifyFilterProfileName") + types.AddKeyToken(snmpNotifyFilterEntry.SnmpNotifyFilterSubtree, "snmpNotifyFilterSubtree")
    snmpNotifyFilterEntry.EntityData.AbsolutePath = "SNMP-NOTIFICATION-MIB:SNMP-NOTIFICATION-MIB/snmpNotifyFilterTable/" + snmpNotifyFilterEntry.EntityData.SegmentPath
    snmpNotifyFilterEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpNotifyFilterEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpNotifyFilterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpNotifyFilterEntry.EntityData.Children = types.NewOrderedMap()
    snmpNotifyFilterEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpNotifyFilterEntry.EntityData.Leafs.Append("snmpNotifyFilterProfileName", types.YLeaf{"SnmpNotifyFilterProfileName", snmpNotifyFilterEntry.SnmpNotifyFilterProfileName})
    snmpNotifyFilterEntry.EntityData.Leafs.Append("snmpNotifyFilterSubtree", types.YLeaf{"SnmpNotifyFilterSubtree", snmpNotifyFilterEntry.SnmpNotifyFilterSubtree})
    snmpNotifyFilterEntry.EntityData.Leafs.Append("snmpNotifyFilterMask", types.YLeaf{"SnmpNotifyFilterMask", snmpNotifyFilterEntry.SnmpNotifyFilterMask})
    snmpNotifyFilterEntry.EntityData.Leafs.Append("snmpNotifyFilterType", types.YLeaf{"SnmpNotifyFilterType", snmpNotifyFilterEntry.SnmpNotifyFilterType})
    snmpNotifyFilterEntry.EntityData.Leafs.Append("snmpNotifyFilterStorageType", types.YLeaf{"SnmpNotifyFilterStorageType", snmpNotifyFilterEntry.SnmpNotifyFilterStorageType})

    snmpNotifyFilterEntry.EntityData.YListKeys = []string {"SnmpNotifyFilterProfileName", "SnmpNotifyFilterSubtree"}

    return &(snmpNotifyFilterEntry.EntityData)
}

