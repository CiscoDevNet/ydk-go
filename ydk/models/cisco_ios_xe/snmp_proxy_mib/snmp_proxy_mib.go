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
    parent types.Entity
    YFilter yfilter.YFilter

    // The table of translation parameters used by proxy forwarder applications
    // for forwarding SNMP messages.
    Snmpproxytable SNMPPROXYMIB_Snmpproxytable
}

func (sNMPPROXYMIB *SNMPPROXYMIB) GetFilter() yfilter.YFilter { return sNMPPROXYMIB.YFilter }

func (sNMPPROXYMIB *SNMPPROXYMIB) SetFilter(yf yfilter.YFilter) { sNMPPROXYMIB.YFilter = yf }

func (sNMPPROXYMIB *SNMPPROXYMIB) GetGoName(yname string) string {
    if yname == "snmpProxyTable" { return "Snmpproxytable" }
    return ""
}

func (sNMPPROXYMIB *SNMPPROXYMIB) GetSegmentPath() string {
    return "SNMP-PROXY-MIB:SNMP-PROXY-MIB"
}

func (sNMPPROXYMIB *SNMPPROXYMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmpProxyTable" {
        return &sNMPPROXYMIB.Snmpproxytable
    }
    return nil
}

func (sNMPPROXYMIB *SNMPPROXYMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snmpProxyTable"] = &sNMPPROXYMIB.Snmpproxytable
    return children
}

func (sNMPPROXYMIB *SNMPPROXYMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sNMPPROXYMIB *SNMPPROXYMIB) GetBundleName() string { return "cisco_ios_xe" }

func (sNMPPROXYMIB *SNMPPROXYMIB) GetYangName() string { return "SNMP-PROXY-MIB" }

func (sNMPPROXYMIB *SNMPPROXYMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sNMPPROXYMIB *SNMPPROXYMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sNMPPROXYMIB *SNMPPROXYMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sNMPPROXYMIB *SNMPPROXYMIB) SetParent(parent types.Entity) { sNMPPROXYMIB.parent = parent }

func (sNMPPROXYMIB *SNMPPROXYMIB) GetParent() types.Entity { return sNMPPROXYMIB.parent }

func (sNMPPROXYMIB *SNMPPROXYMIB) GetParentYangName() string { return "SNMP-PROXY-MIB" }

// SNMPPROXYMIB_Snmpproxytable
// The table of translation parameters used by proxy forwarder
// applications for forwarding SNMP messages.
type SNMPPROXYMIB_Snmpproxytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of translation parameters used by a proxy forwarder application for
    // forwarding SNMP messages.  Entries in the snmpProxyTable are created and
    // deleted using the snmpProxyRowStatus object. The type is slice of
    // SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry.
    Snmpproxyentry []SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry
}

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetFilter() yfilter.YFilter { return snmpproxytable.YFilter }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) SetFilter(yf yfilter.YFilter) { snmpproxytable.YFilter = yf }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetGoName(yname string) string {
    if yname == "snmpProxyEntry" { return "Snmpproxyentry" }
    return ""
}

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetSegmentPath() string {
    return "snmpProxyTable"
}

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmpProxyEntry" {
        for _, c := range snmpproxytable.Snmpproxyentry {
            if snmpproxytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry{}
        snmpproxytable.Snmpproxyentry = append(snmpproxytable.Snmpproxyentry, child)
        return &snmpproxytable.Snmpproxyentry[len(snmpproxytable.Snmpproxyentry)-1]
    }
    return nil
}

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range snmpproxytable.Snmpproxyentry {
        children[snmpproxytable.Snmpproxyentry[i].GetSegmentPath()] = &snmpproxytable.Snmpproxyentry[i]
    }
    return children
}

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetBundleName() string { return "cisco_ios_xe" }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetYangName() string { return "snmpProxyTable" }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) SetParent(parent types.Entity) { snmpproxytable.parent = parent }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetParent() types.Entity { return snmpproxytable.parent }

func (snmpproxytable *SNMPPROXYMIB_Snmpproxytable) GetParentYangName() string { return "SNMP-PROXY-MIB" }

// SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry
// A set of translation parameters used by a proxy forwarder
// application for forwarding SNMP messages.
// 
// Entries in the snmpProxyTable are created and deleted
// using the snmpProxyRowStatus object.
type SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry struct {
    parent types.Entity
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

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetFilter() yfilter.YFilter { return snmpproxyentry.YFilter }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) SetFilter(yf yfilter.YFilter) { snmpproxyentry.YFilter = yf }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetGoName(yname string) string {
    if yname == "snmpProxyName" { return "Snmpproxyname" }
    if yname == "snmpProxyType" { return "Snmpproxytype" }
    if yname == "snmpProxyContextEngineID" { return "Snmpproxycontextengineid" }
    if yname == "snmpProxyContextName" { return "Snmpproxycontextname" }
    if yname == "snmpProxyTargetParamsIn" { return "Snmpproxytargetparamsin" }
    if yname == "snmpProxySingleTargetOut" { return "Snmpproxysingletargetout" }
    if yname == "snmpProxyMultipleTargetOut" { return "Snmpproxymultipletargetout" }
    if yname == "snmpProxyStorageType" { return "Snmpproxystoragetype" }
    if yname == "snmpProxyRowStatus" { return "Snmpproxyrowstatus" }
    return ""
}

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetSegmentPath() string {
    return "snmpProxyEntry" + "[snmpProxyName='" + fmt.Sprintf("%v", snmpproxyentry.Snmpproxyname) + "']"
}

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpProxyName"] = snmpproxyentry.Snmpproxyname
    leafs["snmpProxyType"] = snmpproxyentry.Snmpproxytype
    leafs["snmpProxyContextEngineID"] = snmpproxyentry.Snmpproxycontextengineid
    leafs["snmpProxyContextName"] = snmpproxyentry.Snmpproxycontextname
    leafs["snmpProxyTargetParamsIn"] = snmpproxyentry.Snmpproxytargetparamsin
    leafs["snmpProxySingleTargetOut"] = snmpproxyentry.Snmpproxysingletargetout
    leafs["snmpProxyMultipleTargetOut"] = snmpproxyentry.Snmpproxymultipletargetout
    leafs["snmpProxyStorageType"] = snmpproxyentry.Snmpproxystoragetype
    leafs["snmpProxyRowStatus"] = snmpproxyentry.Snmpproxyrowstatus
    return leafs
}

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetBundleName() string { return "cisco_ios_xe" }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetYangName() string { return "snmpProxyEntry" }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) SetParent(parent types.Entity) { snmpproxyentry.parent = parent }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetParent() types.Entity { return snmpproxyentry.parent }

func (snmpproxyentry *SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry) GetParentYangName() string { return "snmpProxyTable" }

// SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype represents the translation parameters defined by this entry.
type SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype string

const (
    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_read SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "read"

    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_write SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "write"

    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_trap SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "trap"

    SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype_inform SNMPPROXYMIB_Snmpproxytable_Snmpproxyentry_Snmpproxytype = "inform"
)

