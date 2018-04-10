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
    Snmpproxytable SNMPPROXYMIB_Snmpproxytable
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

    sNMPPROXYMIB.EntityData.Children = make(map[string]types.YChild)
    sNMPPROXYMIB.EntityData.Children["snmpProxyTable"] = types.YChild{"Snmpproxytable", &sNMPPROXYMIB.Snmpproxytable}
    sNMPPROXYMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sNMPPROXYMIB.EntityData)
}

// SNMPPROXYMIB_Snmpproxytable
// The table of translation parameters used by proxy forwarder
// applications for forwarding SNMP messages.
type SNMPPROXYMIB_Snmpproxytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of translation parameters used by a proxy forwarder application for
    // forwarding SNMP messages.  Entries in the snmpProxyTable are created and
    // deleted using the snmpProxyRowStatus object. The type is slice of
    // SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry.
    Snmpproxyentry []SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry
}

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetEntityData() *types.CommonEntityData {
    snmpproxytable.EntityData.YFilter = snmpproxytable.YFilter
    snmpproxytable.EntityData.YangName = "snmpProxyTable"
    snmpproxytable.EntityData.BundleName = "cisco_ios_xe"
    snmpproxytable.EntityData.ParentYangName = "SNMP-PROXY-MIB"
    snmpproxytable.EntityData.SegmentPath = "snmpProxyTable"
    snmpproxytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpproxytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpproxytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpproxytable.EntityData.Children = make(map[string]types.YChild)
    snmpproxytable.EntityData.Children["snmpProxyEntry"] = types.YChild{"Snmpproxyentry", nil}
    for i := range snmpproxytable.Snmpproxyentry {
        snmpproxytable.EntityData.Children[types.GetSegmentPath(&snmpproxytable.Snmpproxyentry[i])] = types.YChild{"Snmpproxyentry", &snmpproxytable.Snmpproxyentry[i]}
    }
    snmpproxytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snmpproxytable.EntityData)
}

// SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry
// A set of translation parameters used by a proxy forwarder
// application for forwarding SNMP messages.
// 
// Entries in the snmpProxyTable are created and deleted
// using the snmpProxyRowStatus object.
type SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this snmpProxyEntry. The type is string with length: 1..32.
    Snmpproxyname interface{}

    // The type of message that may be forwarded using the translation parameters
    // defined by this entry. The type is Snmpproxytype.
    Snmpproxytype interface{}

    // The contextEngineID contained in messages that may be forwarded using the
    // translation parameters defined by this entry. The type is string with
    // length: 5..32.
    Snmpproxycontextengineid interface{}

    // The contextName contained in messages that may be forwarded using the
    // translation parameters defined by this entry.  This object is optional, and
    // if not supported, the contextName contained in a message is ignored when
    // selecting an entry in the snmpProxyTable. The type is string.
    Snmpproxycontextname interface{}

    // This object selects an entry in the snmpTargetParamsTable. The selected
    // entry is used to determine which row of the snmpProxyTable to use for
    // forwarding received messages. The type is string.
    Snmpproxytargetparamsin interface{}

    // This object selects a management target defined in the snmpTargetAddrTable
    // (in the SNMP-TARGET-MIB).  The selected target is defined by an entry in
    // the snmpTargetAddrTable whose index value (snmpTargetAddrName) is equal to
    // this object.  This object is only used when selection of a single target is
    // required (i.e. when forwarding an incoming read or write request). The type
    // is string.
    Snmpproxysingletargetout interface{}

    // This object selects a set of management targets defined in the
    // snmpTargetAddrTable (in the SNMP-TARGET-MIB).  This object is only used
    // when selection of multiple targets is required (i.e. when forwarding an
    // incoming notification). The type is string.
    Snmpproxymultipletargetout interface{}

    // The storage type of this conceptual row. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Snmpproxystoragetype interface{}

    // The status of this conceptual row.  To create a row in this table, a
    // manager must set this object to either createAndGo(4) or createAndWait(5). 
    // The following objects may not be modified while the value of this object is
    // active(1):     - snmpProxyType     - snmpProxyContextEngineID     -
    // snmpProxyContextName     - snmpProxyTargetParamsIn     -
    // snmpProxySingleTargetOut     - snmpProxyMultipleTargetOut. The type is
    // RowStatus.
    Snmpproxyrowstatus interface{}
}

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetEntityData() *types.CommonEntityData {
    snmpproxyentry.EntityData.YFilter = snmpproxyentry.YFilter
    snmpproxyentry.EntityData.YangName = "snmpProxyEntry"
    snmpproxyentry.EntityData.BundleName = "cisco_ios_xe"
    snmpproxyentry.EntityData.ParentYangName = "snmpProxyTable"
    snmpproxyentry.EntityData.SegmentPath = "snmpProxyEntry" + "[snmpProxyName='" + fmt.Sprintf("%v", snmpproxyentry.Snmpproxyname) + "']"
    snmpproxyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpproxyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpproxyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpproxyentry.EntityData.Children = make(map[string]types.YChild)
    snmpproxyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    snmpproxyentry.EntityData.Leafs["snmpProxyName"] = types.YLeaf{"Snmpproxyname", snmpproxyentry.Snmpproxyname}
    snmpproxyentry.EntityData.Leafs["snmpProxyType"] = types.YLeaf{"Snmpproxytype", snmpproxyentry.Snmpproxytype}
    snmpproxyentry.EntityData.Leafs["snmpProxyContextEngineID"] = types.YLeaf{"Snmpproxycontextengineid", snmpproxyentry.Snmpproxycontextengineid}
    snmpproxyentry.EntityData.Leafs["snmpProxyContextName"] = types.YLeaf{"Snmpproxycontextname", snmpproxyentry.Snmpproxycontextname}
    snmpproxyentry.EntityData.Leafs["snmpProxyTargetParamsIn"] = types.YLeaf{"Snmpproxytargetparamsin", snmpproxyentry.Snmpproxytargetparamsin}
    snmpproxyentry.EntityData.Leafs["snmpProxySingleTargetOut"] = types.YLeaf{"Snmpproxysingletargetout", snmpproxyentry.Snmpproxysingletargetout}
    snmpproxyentry.EntityData.Leafs["snmpProxyMultipleTargetOut"] = types.YLeaf{"Snmpproxymultipletargetout", snmpproxyentry.Snmpproxymultipletargetout}
    snmpproxyentry.EntityData.Leafs["snmpProxyStorageType"] = types.YLeaf{"Snmpproxystoragetype", snmpproxyentry.Snmpproxystoragetype}
    snmpproxyentry.EntityData.Leafs["snmpProxyRowStatus"] = types.YLeaf{"Snmpproxyrowstatus", snmpproxyentry.Snmpproxyrowstatus}
    return &(snmpproxyentry.EntityData)
}

// SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype represents the translation parameters defined by this entry.
type SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype string

const (
    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_read SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "read"

    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_write SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "write"

    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_trap SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "trap"

    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_inform SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "inform"
)

