// This MIB module defines MIB objects which provide
// mechanisms to remotely configure the parameters
// used by a proxy forwarding application.
// 
// Copyright (C) The Internet Society (2002). This
// version of this MIB module is part of RFC 3413;
// see the RFC itself for full legal notices.
package snmp_proxy_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_proxy_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:SNMP-PROXY-MIB SNMP-PROXY-MIB}", reflect.TypeOf(SNMPPROXYMIB{}))
    ydk.RegisterEntity("SNMP-PROXY-MIB:SNMP-PROXY-MIB", reflect.TypeOf(SNMPPROXYMIB{}))
}

// SNMPPROXYMIB
type SNMPPROXYMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The table of translation parameters used by proxy forwarder applications
    // for forwarding SNMP messages.
    SnmpProxyTable SNMPPROXYMIB_SnmpProxyTable
}

func (sNMPPROXYMIB *SNMPPROXYMIB) GetEntityData() *types.CommonEntityData {
    sNMPPROXYMIB.EntityData.YFilter = sNMPPROXYMIB.YFilter
    sNMPPROXYMIB.EntityData.YangName = "SNMP-PROXY-MIB"
    sNMPPROXYMIB.EntityData.BundleName = "cisco_ios_xe"
    sNMPPROXYMIB.EntityData.ParentYangName = "SNMP-PROXY-MIB"
    sNMPPROXYMIB.EntityData.SegmentPath = "SNMP-PROXY-MIB:SNMP-PROXY-MIB"
    sNMPPROXYMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sNMPPROXYMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sNMPPROXYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sNMPPROXYMIB.EntityData.Children = types.NewOrderedMap()
    sNMPPROXYMIB.EntityData.Children.Append("snmpProxyTable", types.YChild{"SnmpProxyTable", &sNMPPROXYMIB.SnmpProxyTable})
    sNMPPROXYMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPPROXYMIB.EntityData.YListKeys = []string {}

    return &(sNMPPROXYMIB.EntityData)
}

// SNMPPROXYMIB_SnmpProxyTable
// The table of translation parameters used by proxy forwarder
// applications for forwarding SNMP messages.
type SNMPPROXYMIB_SnmpProxyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of translation parameters used by a proxy forwarder application for
    // forwarding SNMP messages.  Entries in the snmpProxyTable are created and
    // deleted using the snmpProxyRowStatus object. The type is slice of
    // SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry.
    SnmpProxyEntry []*SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry
}

func (snmpProxyTable *SNMPPROXYMIB_SnmpProxyTable) GetEntityData() *types.CommonEntityData {
    snmpProxyTable.EntityData.YFilter = snmpProxyTable.YFilter
    snmpProxyTable.EntityData.YangName = "snmpProxyTable"
    snmpProxyTable.EntityData.BundleName = "cisco_ios_xe"
    snmpProxyTable.EntityData.ParentYangName = "SNMP-PROXY-MIB"
    snmpProxyTable.EntityData.SegmentPath = "snmpProxyTable"
    snmpProxyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpProxyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpProxyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpProxyTable.EntityData.Children = types.NewOrderedMap()
    snmpProxyTable.EntityData.Children.Append("snmpProxyEntry", types.YChild{"SnmpProxyEntry", nil})
    for i := range snmpProxyTable.SnmpProxyEntry {
        snmpProxyTable.EntityData.Children.Append(types.GetSegmentPath(snmpProxyTable.SnmpProxyEntry[i]), types.YChild{"SnmpProxyEntry", snmpProxyTable.SnmpProxyEntry[i]})
    }
    snmpProxyTable.EntityData.Leafs = types.NewOrderedMap()

    snmpProxyTable.EntityData.YListKeys = []string {}

    return &(snmpProxyTable.EntityData)
}

// SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry
// A set of translation parameters used by a proxy forwarder
// application for forwarding SNMP messages.
// 
// Entries in the snmpProxyTable are created and deleted
// using the snmpProxyRowStatus object.
type SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this snmpProxyEntry. The type is string with length: 1..32.
    SnmpProxyName interface{}

    // The type of message that may be forwarded using the translation parameters
    // defined by this entry. The type is SnmpProxyType.
    SnmpProxyType interface{}

    // The contextEngineID contained in messages that may be forwarded using the
    // translation parameters defined by this entry. The type is string with
    // length: 5..32.
    SnmpProxyContextEngineID interface{}

    // The contextName contained in messages that may be forwarded using the
    // translation parameters defined by this entry.  This object is optional, and
    // if not supported, the contextName contained in a message is ignored when
    // selecting an entry in the snmpProxyTable. The type is string.
    SnmpProxyContextName interface{}

    // This object selects an entry in the snmpTargetParamsTable. The selected
    // entry is used to determine which row of the snmpProxyTable to use for
    // forwarding received messages. The type is string.
    SnmpProxyTargetParamsIn interface{}

    // This object selects a management target defined in the snmpTargetAddrTable
    // (in the SNMP-TARGET-MIB).  The selected target is defined by an entry in
    // the snmpTargetAddrTable whose index value (snmpTargetAddrName) is equal to
    // this object.  This object is only used when selection of a single target is
    // required (i.e. when forwarding an incoming read or write request). The type
    // is string.
    SnmpProxySingleTargetOut interface{}

    // This object selects a set of management targets defined in the
    // snmpTargetAddrTable (in the SNMP-TARGET-MIB).  This object is only used
    // when selection of multiple targets is required (i.e. when forwarding an
    // incoming notification). The type is string.
    SnmpProxyMultipleTargetOut interface{}

    // The storage type of this conceptual row. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    SnmpProxyStorageType interface{}

    // The status of this conceptual row.  To create a row in this table, a
    // manager must set this object to either createAndGo(4) or createAndWait(5). 
    // The following objects may not be modified while the value of this object is
    // active(1):     - snmpProxyType     - snmpProxyContextEngineID     -
    // snmpProxyContextName     - snmpProxyTargetParamsIn     -
    // snmpProxySingleTargetOut     - snmpProxyMultipleTargetOut. The type is
    // RowStatus.
    SnmpProxyRowStatus interface{}
}

func (snmpProxyEntry *SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry) GetEntityData() *types.CommonEntityData {
    snmpProxyEntry.EntityData.YFilter = snmpProxyEntry.YFilter
    snmpProxyEntry.EntityData.YangName = "snmpProxyEntry"
    snmpProxyEntry.EntityData.BundleName = "cisco_ios_xe"
    snmpProxyEntry.EntityData.ParentYangName = "snmpProxyTable"
    snmpProxyEntry.EntityData.SegmentPath = "snmpProxyEntry" + types.AddKeyToken(snmpProxyEntry.SnmpProxyName, "snmpProxyName")
    snmpProxyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpProxyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpProxyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpProxyEntry.EntityData.Children = types.NewOrderedMap()
    snmpProxyEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyName", types.YLeaf{"SnmpProxyName", snmpProxyEntry.SnmpProxyName})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyType", types.YLeaf{"SnmpProxyType", snmpProxyEntry.SnmpProxyType})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyContextEngineID", types.YLeaf{"SnmpProxyContextEngineID", snmpProxyEntry.SnmpProxyContextEngineID})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyContextName", types.YLeaf{"SnmpProxyContextName", snmpProxyEntry.SnmpProxyContextName})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyTargetParamsIn", types.YLeaf{"SnmpProxyTargetParamsIn", snmpProxyEntry.SnmpProxyTargetParamsIn})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxySingleTargetOut", types.YLeaf{"SnmpProxySingleTargetOut", snmpProxyEntry.SnmpProxySingleTargetOut})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyMultipleTargetOut", types.YLeaf{"SnmpProxyMultipleTargetOut", snmpProxyEntry.SnmpProxyMultipleTargetOut})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyStorageType", types.YLeaf{"SnmpProxyStorageType", snmpProxyEntry.SnmpProxyStorageType})
    snmpProxyEntry.EntityData.Leafs.Append("snmpProxyRowStatus", types.YLeaf{"SnmpProxyRowStatus", snmpProxyEntry.SnmpProxyRowStatus})

    snmpProxyEntry.EntityData.YListKeys = []string {"SnmpProxyName"}

    return &(snmpProxyEntry.EntityData)
}

// SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType represents the translation parameters defined by this entry.
type SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType string

const (
    SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType_read SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType = "read"

    SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType_write SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType = "write"

    SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType_trap SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType = "trap"

    SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType_inform SNMPPROXYMIB_SnmpProxyTable_SnmpProxyEntry_SnmpProxyType = "inform"
)

