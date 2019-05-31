package snmpv2_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmpv2_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMPv2-MIB/200210160000Z SNMPv2-MIB}", reflect.TypeOf(SNMPv2MIB{}))
    ydk.RegisterEntity("SNMPv2-MIB:SNMPv2-MIB", reflect.TypeOf(SNMPv2MIB{}))
}

// SnmpEnableAuthenTrapsType
type SnmpEnableAuthenTrapsType string

const (
    SnmpEnableAuthenTrapsType_enabled SnmpEnableAuthenTrapsType = "enabled"

    SnmpEnableAuthenTrapsType_disabled SnmpEnableAuthenTrapsType = "disabled"
)

// SNMPv2MIB
type SNMPv2MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    System SNMPv2MIB_System

    
    Snmp SNMPv2MIB_Snmp

    
    SnmpSet SNMPv2MIB_SnmpSet

    
    SysORTable SNMPv2MIB_SysORTable
}

func (sNMPv2MIB *SNMPv2MIB) GetEntityData() *types.CommonEntityData {
    sNMPv2MIB.EntityData.YFilter = sNMPv2MIB.YFilter
    sNMPv2MIB.EntityData.YangName = "SNMPv2-MIB"
    sNMPv2MIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPv2MIB.EntityData.ParentYangName = "SNMPv2-MIB"
    sNMPv2MIB.EntityData.SegmentPath = "SNMPv2-MIB:SNMPv2-MIB"
    sNMPv2MIB.EntityData.AbsolutePath = sNMPv2MIB.EntityData.SegmentPath
    sNMPv2MIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPv2MIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPv2MIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPv2MIB.EntityData.Children = types.NewOrderedMap()
    sNMPv2MIB.EntityData.Children.Append("system", types.YChild{"System", &sNMPv2MIB.System})
    sNMPv2MIB.EntityData.Children.Append("snmp", types.YChild{"Snmp", &sNMPv2MIB.Snmp})
    sNMPv2MIB.EntityData.Children.Append("snmpSet", types.YChild{"SnmpSet", &sNMPv2MIB.SnmpSet})
    sNMPv2MIB.EntityData.Children.Append("sysORTable", types.YChild{"SysORTable", &sNMPv2MIB.SysORTable})
    sNMPv2MIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPv2MIB.EntityData.YListKeys = []string {}

    return &(sNMPv2MIB.EntityData)
}

// SNMPv2MIB_System
type SNMPv2MIB_System struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with length: 0..255.
    SysDescr interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    SysObjectID interface{}

    // The type is interface{} with range: 0..4294967295.
    SysUpTime interface{}

    // The type is string with length: 0..255.
    SysContact interface{}

    // The type is string with length: 0..255.
    SysName interface{}

    // The type is string with length: 0..255.
    SysLocation interface{}

    // The type is interface{} with range: 0..127.
    SysServices interface{}

    // The type is interface{} with range: 0..4294967295.
    SysORLastChange interface{}
}

func (system *SNMPv2MIB_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xr"
    system.EntityData.ParentYangName = "SNMPv2-MIB"
    system.EntityData.SegmentPath = "system"
    system.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + system.EntityData.SegmentPath
    system.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Leafs = types.NewOrderedMap()
    system.EntityData.Leafs.Append("sysDescr", types.YLeaf{"SysDescr", system.SysDescr})
    system.EntityData.Leafs.Append("sysObjectID", types.YLeaf{"SysObjectID", system.SysObjectID})
    system.EntityData.Leafs.Append("sysUpTime", types.YLeaf{"SysUpTime", system.SysUpTime})
    system.EntityData.Leafs.Append("sysContact", types.YLeaf{"SysContact", system.SysContact})
    system.EntityData.Leafs.Append("sysName", types.YLeaf{"SysName", system.SysName})
    system.EntityData.Leafs.Append("sysLocation", types.YLeaf{"SysLocation", system.SysLocation})
    system.EntityData.Leafs.Append("sysServices", types.YLeaf{"SysServices", system.SysServices})
    system.EntityData.Leafs.Append("sysORLastChange", types.YLeaf{"SysORLastChange", system.SysORLastChange})

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// SNMPv2MIB_Snmp
type SNMPv2MIB_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    SnmpInPkts interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpInBadVersions interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpInBadCommunityNames interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpInBadCommunityUses interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpInASNParseErrs interface{}

    // The type is SnmpEnableAuthenTrapsType. The default value is disabled.
    SnmpEnableAuthenTraps interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpSilentDrops interface{}

    // The type is interface{} with range: 0..4294967295.
    SnmpProxyDrops interface{}
}

func (snmp *SNMPv2MIB_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "SNMPv2-MIB"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + snmp.EntityData.SegmentPath
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("snmpInPkts", types.YLeaf{"SnmpInPkts", snmp.SnmpInPkts})
    snmp.EntityData.Leafs.Append("snmpInBadVersions", types.YLeaf{"SnmpInBadVersions", snmp.SnmpInBadVersions})
    snmp.EntityData.Leafs.Append("snmpInBadCommunityNames", types.YLeaf{"SnmpInBadCommunityNames", snmp.SnmpInBadCommunityNames})
    snmp.EntityData.Leafs.Append("snmpInBadCommunityUses", types.YLeaf{"SnmpInBadCommunityUses", snmp.SnmpInBadCommunityUses})
    snmp.EntityData.Leafs.Append("snmpInASNParseErrs", types.YLeaf{"SnmpInASNParseErrs", snmp.SnmpInASNParseErrs})
    snmp.EntityData.Leafs.Append("snmpEnableAuthenTraps", types.YLeaf{"SnmpEnableAuthenTraps", snmp.SnmpEnableAuthenTraps})
    snmp.EntityData.Leafs.Append("snmpSilentDrops", types.YLeaf{"SnmpSilentDrops", snmp.SnmpSilentDrops})
    snmp.EntityData.Leafs.Append("snmpProxyDrops", types.YLeaf{"SnmpProxyDrops", snmp.SnmpProxyDrops})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// SNMPv2MIB_SnmpSet
type SNMPv2MIB_SnmpSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..2147483647.
    SnmpSetSerialNo interface{}
}

func (snmpSet *SNMPv2MIB_SnmpSet) GetEntityData() *types.CommonEntityData {
    snmpSet.EntityData.YFilter = snmpSet.YFilter
    snmpSet.EntityData.YangName = "snmpSet"
    snmpSet.EntityData.BundleName = "cisco_ios_xr"
    snmpSet.EntityData.ParentYangName = "SNMPv2-MIB"
    snmpSet.EntityData.SegmentPath = "snmpSet"
    snmpSet.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + snmpSet.EntityData.SegmentPath
    snmpSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmpSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmpSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmpSet.EntityData.Children = types.NewOrderedMap()
    snmpSet.EntityData.Leafs = types.NewOrderedMap()
    snmpSet.EntityData.Leafs.Append("snmpSetSerialNo", types.YLeaf{"SnmpSetSerialNo", snmpSet.SnmpSetSerialNo})

    snmpSet.EntityData.YListKeys = []string {}

    return &(snmpSet.EntityData)
}

// SNMPv2MIB_SysORTable
type SNMPv2MIB_SysORTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SNMPv2MIB_SysORTable_SysOREntry.
    SysOREntry []*SNMPv2MIB_SysORTable_SysOREntry
}

func (sysORTable *SNMPv2MIB_SysORTable) GetEntityData() *types.CommonEntityData {
    sysORTable.EntityData.YFilter = sysORTable.YFilter
    sysORTable.EntityData.YangName = "sysORTable"
    sysORTable.EntityData.BundleName = "cisco_ios_xr"
    sysORTable.EntityData.ParentYangName = "SNMPv2-MIB"
    sysORTable.EntityData.SegmentPath = "sysORTable"
    sysORTable.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + sysORTable.EntityData.SegmentPath
    sysORTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysORTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysORTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysORTable.EntityData.Children = types.NewOrderedMap()
    sysORTable.EntityData.Children.Append("sysOREntry", types.YChild{"SysOREntry", nil})
    for i := range sysORTable.SysOREntry {
        sysORTable.EntityData.Children.Append(types.GetSegmentPath(sysORTable.SysOREntry[i]), types.YChild{"SysOREntry", sysORTable.SysOREntry[i]})
    }
    sysORTable.EntityData.Leafs = types.NewOrderedMap()

    sysORTable.EntityData.YListKeys = []string {}

    return &(sysORTable.EntityData)
}

// SNMPv2MIB_SysORTable_SysOREntry
type SNMPv2MIB_SysORTable_SysOREntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    SysORIndex interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    SysORID interface{}

    // The type is string with length: 0..255.
    SysORDescr interface{}

    // The type is interface{} with range: 0..4294967295.
    SysORUpTime interface{}
}

func (sysOREntry *SNMPv2MIB_SysORTable_SysOREntry) GetEntityData() *types.CommonEntityData {
    sysOREntry.EntityData.YFilter = sysOREntry.YFilter
    sysOREntry.EntityData.YangName = "sysOREntry"
    sysOREntry.EntityData.BundleName = "cisco_ios_xr"
    sysOREntry.EntityData.ParentYangName = "sysORTable"
    sysOREntry.EntityData.SegmentPath = "sysOREntry" + types.AddKeyToken(sysOREntry.SysORIndex, "sysORIndex")
    sysOREntry.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/sysORTable/" + sysOREntry.EntityData.SegmentPath
    sysOREntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysOREntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysOREntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysOREntry.EntityData.Children = types.NewOrderedMap()
    sysOREntry.EntityData.Leafs = types.NewOrderedMap()
    sysOREntry.EntityData.Leafs.Append("sysORIndex", types.YLeaf{"SysORIndex", sysOREntry.SysORIndex})
    sysOREntry.EntityData.Leafs.Append("sysORID", types.YLeaf{"SysORID", sysOREntry.SysORID})
    sysOREntry.EntityData.Leafs.Append("sysORDescr", types.YLeaf{"SysORDescr", sysOREntry.SysORDescr})
    sysOREntry.EntityData.Leafs.Append("sysORUpTime", types.YLeaf{"SysORUpTime", sysOREntry.SysORUpTime})

    sysOREntry.EntityData.YListKeys = []string {"SysORIndex"}

    return &(sysOREntry.EntityData)
}

