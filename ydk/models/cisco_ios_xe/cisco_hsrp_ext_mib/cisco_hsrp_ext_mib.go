// The Extension MIB module for the CISCO-HSRP-MIB which is
// based on RFC2281.
// 
// This MIB provides an extension to the CISCO-HSRP-MIB which 
// defines Cisco's proprietary Hot Standby Routing Protocol 
// (HSRP), defined in RFC2281. The extensions cover assigning 
// of secondary HSRP ip addresses, modifying an HSRP Group's 
// priority by tracking the operational status of interfaces, 
// etc. 
// 
// Terminology:
// HSRP is a protocol used amoung a group of routers for the 
// purpose of selecting an active router and a standby router. 
// 
// An active router is the router of choice for routing 
// packets.
// 
// A standby router is a router that takes over the routing 
// duties when an active router fails, or when preset 
// conditions have been met.
// 
// A HSRP group or a standby group is a set of routers 
// which communicate using HSRP. An HSRP group has a group 
// MAC address and a group IP address. These are the 
// designated addresses. The active router assumes  
// (i.e. inherits) these group addresses. An HSRP group is
// identified by a ( ifIndex, cHsrpGrpNumber ) pair.
// 
// BIA stands for Burned In Address.
package cisco_hsrp_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_hsrp_ext_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-HSRP-EXT-MIB CISCO-HSRP-EXT-MIB}", reflect.TypeOf(CISCOHSRPEXTMIB{}))
    ydk.RegisterEntity("CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB", reflect.TypeOf(CISCOHSRPEXTMIB{}))
}

// CISCOHSRPEXTMIB
type CISCOHSRPEXTMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table containing information about tracked interfaces per HSRP group.
    Chsrpextiftrackedtable CISCOHSRPEXTMIB_Chsrpextiftrackedtable

    // A table containing information about secondary HSRP IP Addresses per
    // interface and group.
    Chsrpextsecaddrtable CISCOHSRPEXTMIB_Chsrpextsecaddrtable

    // A table containing information about standby interfaces per HSRP group.
    Chsrpextifstandbytable CISCOHSRPEXTMIB_Chsrpextifstandbytable

    // HSRP-specific configurations for each physical interface.
    Chsrpextiftable CISCOHSRPEXTMIB_Chsrpextiftable
}

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetFilter() yfilter.YFilter { return cISCOHSRPEXTMIB.YFilter }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) SetFilter(yf yfilter.YFilter) { cISCOHSRPEXTMIB.YFilter = yf }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetGoName(yname string) string {
    if yname == "cHsrpExtIfTrackedTable" { return "Chsrpextiftrackedtable" }
    if yname == "cHsrpExtSecAddrTable" { return "Chsrpextsecaddrtable" }
    if yname == "cHsrpExtIfStandbyTable" { return "Chsrpextifstandbytable" }
    if yname == "cHsrpExtIfTable" { return "Chsrpextiftable" }
    return ""
}

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetSegmentPath() string {
    return "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB"
}

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cHsrpExtIfTrackedTable" {
        return &cISCOHSRPEXTMIB.Chsrpextiftrackedtable
    }
    if childYangName == "cHsrpExtSecAddrTable" {
        return &cISCOHSRPEXTMIB.Chsrpextsecaddrtable
    }
    if childYangName == "cHsrpExtIfStandbyTable" {
        return &cISCOHSRPEXTMIB.Chsrpextifstandbytable
    }
    if childYangName == "cHsrpExtIfTable" {
        return &cISCOHSRPEXTMIB.Chsrpextiftable
    }
    return nil
}

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cHsrpExtIfTrackedTable"] = &cISCOHSRPEXTMIB.Chsrpextiftrackedtable
    children["cHsrpExtSecAddrTable"] = &cISCOHSRPEXTMIB.Chsrpextsecaddrtable
    children["cHsrpExtIfStandbyTable"] = &cISCOHSRPEXTMIB.Chsrpextifstandbytable
    children["cHsrpExtIfTable"] = &cISCOHSRPEXTMIB.Chsrpextiftable
    return children
}

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetYangName() string { return "CISCO-HSRP-EXT-MIB" }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) SetParent(parent types.Entity) { cISCOHSRPEXTMIB.parent = parent }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetParent() types.Entity { return cISCOHSRPEXTMIB.parent }

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetParentYangName() string { return "CISCO-HSRP-EXT-MIB" }

// CISCOHSRPEXTMIB_Chsrpextiftrackedtable
// A table containing information about tracked interfaces per
// HSRP group.
type CISCOHSRPEXTMIB_Chsrpextiftrackedtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each row of this table allows the tracking of one interface of the HSRP
    // group which is identified by the (ifIndex, cHsrpGrpNumber) values in this
    // table's INDEX clause. Weight(priority) is given to each and every interface
    // tracked.  When a tracked interface is unavailable, the HSRP priority of the
    // router is decreased. i.e cHsrpGrpPriority value assigned  to an HSRP group
    // will reduce by the value assigned to cHsrpExtIfTrackedPriority. This
    // reduces the likelihood  of a router with a failed key interface becoming
    // the  active router.  Setting cHsrpExtIfTrackedRowStatus to active starts
    // the tracking of cHsrpExtIfTracked by the HSRP group.  The value of
    // cHsrpExtIfTrackedRowStatus may be set  to destroy at any time.  Entries may
    // not be created via SNMP without explicitly setting
    // cHsrpExtIfTrackedRowStatus to either 'createAndGo'  or 'createAndWait'. 
    // Entries can be created and modified via the management protocol or by the
    // device's local management interface.  If the row is not active, and a local
    // management interface command modifies that row, the row may transition to
    // active state.  A row entry in the cHsrpExtIfTrackedTable can not be created
    // unless the corresponding row in the cHsrpGrpTable has been  created. If
    // that corresponding row in cHsrpGrpTable is  deleted, the interfaces it
    // tracks also get deleted.  A row which is not in active state will timeout
    // after a configurable period (five minutes by default). This timeout period
    // can be changed by setting cHsrpConfigTimeout. The type is slice of
    // CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry.
    Chsrpextiftrackedentry []CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry
}

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetFilter() yfilter.YFilter { return chsrpextiftrackedtable.YFilter }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) SetFilter(yf yfilter.YFilter) { chsrpextiftrackedtable.YFilter = yf }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetGoName(yname string) string {
    if yname == "cHsrpExtIfTrackedEntry" { return "Chsrpextiftrackedentry" }
    return ""
}

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetSegmentPath() string {
    return "cHsrpExtIfTrackedTable"
}

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cHsrpExtIfTrackedEntry" {
        for _, c := range chsrpextiftrackedtable.Chsrpextiftrackedentry {
            if chsrpextiftrackedtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry{}
        chsrpextiftrackedtable.Chsrpextiftrackedentry = append(chsrpextiftrackedtable.Chsrpextiftrackedentry, child)
        return &chsrpextiftrackedtable.Chsrpextiftrackedentry[len(chsrpextiftrackedtable.Chsrpextiftrackedentry)-1]
    }
    return nil
}

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range chsrpextiftrackedtable.Chsrpextiftrackedentry {
        children[chsrpextiftrackedtable.Chsrpextiftrackedentry[i].GetSegmentPath()] = &chsrpextiftrackedtable.Chsrpextiftrackedentry[i]
    }
    return children
}

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetYangName() string { return "cHsrpExtIfTrackedTable" }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) SetParent(parent types.Entity) { chsrpextiftrackedtable.parent = parent }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetParent() types.Entity { return chsrpextiftrackedtable.parent }

func (chsrpextiftrackedtable *CISCOHSRPEXTMIB_Chsrpextiftrackedtable) GetParentYangName() string { return "CISCO-HSRP-EXT-MIB" }

// CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry
// Each row of this table allows the tracking of one
// interface of the HSRP group which is identified by the
// (ifIndex, cHsrpGrpNumber) values in this table's INDEX clause.
// Weight(priority) is given to each and every interface tracked. 
// When a tracked interface is unavailable, the HSRP priority of
// the router is decreased. i.e cHsrpGrpPriority value assigned 
// to an HSRP group will reduce by the value assigned to
// cHsrpExtIfTrackedPriority. This reduces the likelihood 
// of a router with a failed key interface becoming the 
// active router.
// 
// Setting cHsrpExtIfTrackedRowStatus to active starts
// the tracking of cHsrpExtIfTracked by the HSRP group. 
// The value of cHsrpExtIfTrackedRowStatus may be set 
// to destroy at any time.
// 
// Entries may not be created via SNMP without explicitly setting
// cHsrpExtIfTrackedRowStatus to either 'createAndGo' 
// or 'createAndWait'.
// 
// Entries can be created and modified via the management
// protocol or by the device's local management interface.
// 
// If the row is not active, and a local management interface
// command modifies that row, the row may transition to active
// state.
// 
// A row entry in the cHsrpExtIfTrackedTable can not be created
// unless the corresponding row in the cHsrpGrpTable has been 
// created. If that corresponding row in cHsrpGrpTable is 
// deleted, the interfaces it tracks also get deleted.
// 
// A row which is not in active state will timeout after a
// configurable period (five minutes by default). This timeout
// period can be changed by setting cHsrpConfigTimeout.
type CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..255. Refers to
    // cisco_hsrp_mib.CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry_Chsrpgrpnumber
    Chsrpgrpnumber interface{}

    // This attribute is a key. The ifIndex value of the tracked interface. The
    // type is interface{} with range: 1..2147483647.
    Chsrpextiftracked interface{}

    // Priority of the tracked interface for the corresponding { ifIndex,
    // cHsrpGrpNumber } pair. In the range of 0 to 255, 0 is the lowest priority
    // and 255 is the highest. When a tracked  interface is unavailable, the
    // cHsrpGrpPriority of the router  is decreased by the value of this object
    // instance (If the  cHsrpGrpPriority is less than the 
    // cHsrpExtIfTrackedPriority, then the HSRP priority  becomes 0). This allows
    // a standby router to be configured  with a priority such that if the
    // currently active router's  priority is lowered because the tracked
    // interface goes down,  the standby router can takeover. The type is
    // interface{} with range: 0..255.
    Chsrpextiftrackedpriority interface{}

    // The control that allows modification, creation, and deletion of entries.
    // For detailed rules see the DESCRIPTION for cHsrpExtIfTrackedEntry. The type
    // is RowStatus.
    Chsrpextiftrackedrowstatus interface{}

    // This object specifies the disable HSRP IPv4 virtual IP address. The type is
    // bool.
    Chsrpextiftrackedipnone interface{}
}

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetFilter() yfilter.YFilter { return chsrpextiftrackedentry.YFilter }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) SetFilter(yf yfilter.YFilter) { chsrpextiftrackedentry.YFilter = yf }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cHsrpGrpNumber" { return "Chsrpgrpnumber" }
    if yname == "cHsrpExtIfTracked" { return "Chsrpextiftracked" }
    if yname == "cHsrpExtIfTrackedPriority" { return "Chsrpextiftrackedpriority" }
    if yname == "cHsrpExtIfTrackedRowStatus" { return "Chsrpextiftrackedrowstatus" }
    if yname == "cHsrpExtIfTrackedIpNone" { return "Chsrpextiftrackedipnone" }
    return ""
}

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetSegmentPath() string {
    return "cHsrpExtIfTrackedEntry" + "[ifIndex='" + fmt.Sprintf("%v", chsrpextiftrackedentry.Ifindex) + "']" + "[cHsrpGrpNumber='" + fmt.Sprintf("%v", chsrpextiftrackedentry.Chsrpgrpnumber) + "']" + "[cHsrpExtIfTracked='" + fmt.Sprintf("%v", chsrpextiftrackedentry.Chsrpextiftracked) + "']"
}

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = chsrpextiftrackedentry.Ifindex
    leafs["cHsrpGrpNumber"] = chsrpextiftrackedentry.Chsrpgrpnumber
    leafs["cHsrpExtIfTracked"] = chsrpextiftrackedentry.Chsrpextiftracked
    leafs["cHsrpExtIfTrackedPriority"] = chsrpextiftrackedentry.Chsrpextiftrackedpriority
    leafs["cHsrpExtIfTrackedRowStatus"] = chsrpextiftrackedentry.Chsrpextiftrackedrowstatus
    leafs["cHsrpExtIfTrackedIpNone"] = chsrpextiftrackedentry.Chsrpextiftrackedipnone
    return leafs
}

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetYangName() string { return "cHsrpExtIfTrackedEntry" }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) SetParent(parent types.Entity) { chsrpextiftrackedentry.parent = parent }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetParent() types.Entity { return chsrpextiftrackedentry.parent }

func (chsrpextiftrackedentry *CISCOHSRPEXTMIB_Chsrpextiftrackedtable_Chsrpextiftrackedentry) GetParentYangName() string { return "cHsrpExtIfTrackedTable" }

// CISCOHSRPEXTMIB_Chsrpextsecaddrtable
// A table containing information about secondary HSRP IP
// Addresses per interface and group.
type CISCOHSRPEXTMIB_Chsrpextsecaddrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The CHsrpExtSecAddrEntry allows creation of secondary IP Addresses for each
    // cHsrpGrpEntry row.  Secondary addresses can be added by setting 
    // cHsrpExtSecAddrRowStatus to be active. The value of
    // cHsrpExtSecAddrRowStatus may be set to destroy at any time.  Entries may
    // not be created via SNMP without explicitly setting cHsrpExtSecAddrRowStatus
    // to either 'createAndGo' or 'createAndWait'.  Entries can be created and
    // modified via the management protocol or by the device's local management
    // interface.  If the row is not active, and a local management interface
    // command modifies that row, the row may transition to active state.  A row
    // which is not in active state will timeout after a configurable period (five
    // minutes by default). This timeout period can be changed by setting
    // cHsrpConfigTimeout.  Before creation of a cHsrpExtSecAddrEntry row, either
    // cHsrpGrpConfiguredVirtualIpAddr or  cHsrpGrpLearnedVirtualIpAddr must have
    // a valid IP Address. This is because a secondary ip address cannot be
    // created unless the primary ip address has already been set.  To create a
    // new cHsrpExtSecAddrEntry row, a management  station should choose the
    // ifIndex of the interface which is to  be added as part of an HSRP group.
    // Also, an HSRP group number  and a cHsrpExtSecAddrAddress should be chosen. 
    // Deleting a {ifIndex, cHsrpGrpNumber} row in the cHsrpGrpTable will delete
    // all corresponding rows in the cHsrpExtSecAddrTable. Deleting a primary
    // address value in the cHsrpGrpEntry row will delete all secondary addresses
    // for the same {ifIndex, cHsrpGrpNumber} pair. The type is slice of
    // CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry.
    Chsrpextsecaddrentry []CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry
}

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetFilter() yfilter.YFilter { return chsrpextsecaddrtable.YFilter }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) SetFilter(yf yfilter.YFilter) { chsrpextsecaddrtable.YFilter = yf }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetGoName(yname string) string {
    if yname == "cHsrpExtSecAddrEntry" { return "Chsrpextsecaddrentry" }
    return ""
}

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetSegmentPath() string {
    return "cHsrpExtSecAddrTable"
}

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cHsrpExtSecAddrEntry" {
        for _, c := range chsrpextsecaddrtable.Chsrpextsecaddrentry {
            if chsrpextsecaddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry{}
        chsrpextsecaddrtable.Chsrpextsecaddrentry = append(chsrpextsecaddrtable.Chsrpextsecaddrentry, child)
        return &chsrpextsecaddrtable.Chsrpextsecaddrentry[len(chsrpextsecaddrtable.Chsrpextsecaddrentry)-1]
    }
    return nil
}

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range chsrpextsecaddrtable.Chsrpextsecaddrentry {
        children[chsrpextsecaddrtable.Chsrpextsecaddrentry[i].GetSegmentPath()] = &chsrpextsecaddrtable.Chsrpextsecaddrentry[i]
    }
    return children
}

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetYangName() string { return "cHsrpExtSecAddrTable" }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) SetParent(parent types.Entity) { chsrpextsecaddrtable.parent = parent }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetParent() types.Entity { return chsrpextsecaddrtable.parent }

func (chsrpextsecaddrtable *CISCOHSRPEXTMIB_Chsrpextsecaddrtable) GetParentYangName() string { return "CISCO-HSRP-EXT-MIB" }

// CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry
// The CHsrpExtSecAddrEntry allows creation of secondary
// IP Addresses for each cHsrpGrpEntry row.
// 
// Secondary addresses can be added by setting 
// cHsrpExtSecAddrRowStatus to be active. The value of
// cHsrpExtSecAddrRowStatus may be set to destroy at any
// time.
// 
// Entries may not be created via SNMP without explicitly setting
// cHsrpExtSecAddrRowStatus to either 'createAndGo'
// or 'createAndWait'.
// 
// Entries can be created and modified via the management
// protocol or by the device's local management interface.
// 
// If the row is not active, and a local management interface
// command modifies that row, the row may transition to active
// state.
// 
// A row which is not in active state will timeout after a
// configurable period (five minutes by default). This timeout
// period can be changed by setting cHsrpConfigTimeout.
// 
// Before creation of a cHsrpExtSecAddrEntry row,
// either cHsrpGrpConfiguredVirtualIpAddr or 
// cHsrpGrpLearnedVirtualIpAddr must have a valid IP Address.
// This is because a secondary ip address cannot be created
// unless the primary ip address has already been set.
// 
// To create a new cHsrpExtSecAddrEntry row, a management 
// station should choose the ifIndex of the interface which is to 
// be added as part of an HSRP group. Also, an HSRP group number 
// and a cHsrpExtSecAddrAddress should be chosen.
// 
// Deleting a {ifIndex, cHsrpGrpNumber} row in the
// cHsrpGrpTable will delete all corresponding
// rows in the cHsrpExtSecAddrTable.
// Deleting a primary address value in the cHsrpGrpEntry row
// will delete all secondary addresses for the same
// {ifIndex, cHsrpGrpNumber} pair.
type CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..255. Refers to
    // cisco_hsrp_mib.CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry_Chsrpgrpnumber
    Chsrpgrpnumber interface{}

    // This attribute is a key. A secondary IpAddress for the {ifIndex,
    // cHsrpGrpNumber} pair. As explained in the DESCRIPTION for
    // cHsrpExtSecAddrEntry, a primary address must exist before a secondary
    // address for  the same {ifIndex, cHsrpGrpNumber} pair can be created. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Chsrpextsecaddraddress interface{}

    // The control that allows modification, creation, and deletion of entries.
    // For detailed rules see the DESCRIPTION for cHsrpExtSecAddrEntry. The type
    // is RowStatus.
    Chsrpextsecaddrrowstatus interface{}
}

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetFilter() yfilter.YFilter { return chsrpextsecaddrentry.YFilter }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) SetFilter(yf yfilter.YFilter) { chsrpextsecaddrentry.YFilter = yf }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cHsrpGrpNumber" { return "Chsrpgrpnumber" }
    if yname == "cHsrpExtSecAddrAddress" { return "Chsrpextsecaddraddress" }
    if yname == "cHsrpExtSecAddrRowStatus" { return "Chsrpextsecaddrrowstatus" }
    return ""
}

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetSegmentPath() string {
    return "cHsrpExtSecAddrEntry" + "[ifIndex='" + fmt.Sprintf("%v", chsrpextsecaddrentry.Ifindex) + "']" + "[cHsrpGrpNumber='" + fmt.Sprintf("%v", chsrpextsecaddrentry.Chsrpgrpnumber) + "']" + "[cHsrpExtSecAddrAddress='" + fmt.Sprintf("%v", chsrpextsecaddrentry.Chsrpextsecaddraddress) + "']"
}

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = chsrpextsecaddrentry.Ifindex
    leafs["cHsrpGrpNumber"] = chsrpextsecaddrentry.Chsrpgrpnumber
    leafs["cHsrpExtSecAddrAddress"] = chsrpextsecaddrentry.Chsrpextsecaddraddress
    leafs["cHsrpExtSecAddrRowStatus"] = chsrpextsecaddrentry.Chsrpextsecaddrrowstatus
    return leafs
}

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetYangName() string { return "cHsrpExtSecAddrEntry" }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) SetParent(parent types.Entity) { chsrpextsecaddrentry.parent = parent }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetParent() types.Entity { return chsrpextsecaddrentry.parent }

func (chsrpextsecaddrentry *CISCOHSRPEXTMIB_Chsrpextsecaddrtable_Chsrpextsecaddrentry) GetParentYangName() string { return "cHsrpExtSecAddrTable" }

// CISCOHSRPEXTMIB_Chsrpextifstandbytable
// A table containing information about standby
// interfaces per HSRP group.
type CISCOHSRPEXTMIB_Chsrpextifstandbytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The cHsrpExtIfStandbyEntry allows an HSRP group interface to track one or
    // more standby interfaces.  To create a new cHsrpExtIfStandbyEntry row, a
    // management station should choose the ifIndex of the interface which is to
    // be added as part of an HSRP group. Also, an HSRP group number and a
    // cHsrpExtIfStandbyIndex should be chosen. The type is slice of
    // CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry.
    Chsrpextifstandbyentry []CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry
}

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetFilter() yfilter.YFilter { return chsrpextifstandbytable.YFilter }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) SetFilter(yf yfilter.YFilter) { chsrpextifstandbytable.YFilter = yf }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetGoName(yname string) string {
    if yname == "cHsrpExtIfStandbyEntry" { return "Chsrpextifstandbyentry" }
    return ""
}

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetSegmentPath() string {
    return "cHsrpExtIfStandbyTable"
}

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cHsrpExtIfStandbyEntry" {
        for _, c := range chsrpextifstandbytable.Chsrpextifstandbyentry {
            if chsrpextifstandbytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry{}
        chsrpextifstandbytable.Chsrpextifstandbyentry = append(chsrpextifstandbytable.Chsrpextifstandbyentry, child)
        return &chsrpextifstandbytable.Chsrpextifstandbyentry[len(chsrpextifstandbytable.Chsrpextifstandbyentry)-1]
    }
    return nil
}

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range chsrpextifstandbytable.Chsrpextifstandbyentry {
        children[chsrpextifstandbytable.Chsrpextifstandbyentry[i].GetSegmentPath()] = &chsrpextifstandbytable.Chsrpextifstandbyentry[i]
    }
    return children
}

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetYangName() string { return "cHsrpExtIfStandbyTable" }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) SetParent(parent types.Entity) { chsrpextifstandbytable.parent = parent }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetParent() types.Entity { return chsrpextifstandbytable.parent }

func (chsrpextifstandbytable *CISCOHSRPEXTMIB_Chsrpextifstandbytable) GetParentYangName() string { return "CISCO-HSRP-EXT-MIB" }

// CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry
// The cHsrpExtIfStandbyEntry allows an HSRP group
// interface to track one or more standby interfaces.
// 
// To create a new cHsrpExtIfStandbyEntry row, a
// management station should choose the ifIndex of
// the interface which is to be added as part of an
// HSRP group. Also, an HSRP group number and a
// cHsrpExtIfStandbyIndex should be chosen.
type CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..255. Refers to
    // cisco_hsrp_mib.CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry_Chsrpgrpnumber
    Chsrpgrpnumber interface{}

    // This attribute is a key. This object defines the index of the standby
    // table. The type is interface{} with range: 1..4.
    Chsrpextifstandbyindex interface{}

    // This object specifies the type of Internet address denoted by
    // cHsrpExtIfStandbyDestAddr. The type is InetAddressType.
    Chsrpextifstandbydestaddrtype interface{}

    // This object specifies the destination IP address of the standby router. The
    // type is string with length: 0..255.
    Chsrpextifstandbydestaddr interface{}

    // This object specifies the type of Internet address denoted by
    // cHsrpExtIfStandbySourceAddr. The type is InetAddressType.
    Chsrpextifstandbysourceaddrtype interface{}

    // This object specifies the source IP address of the standby router. The type
    // is string with length: 0..255.
    Chsrpextifstandbysourceaddr interface{}

    // The control that allows modification, creation, and deletion of entries.
    // Entries may not be created via SNMP without explicitly setting
    // cHsrpExtIfStandbyRowStatus to either 'createAndGo' or 'createAndWait'. The
    // type is RowStatus.
    Chsrpextifstandbyrowstatus interface{}
}

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetFilter() yfilter.YFilter { return chsrpextifstandbyentry.YFilter }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) SetFilter(yf yfilter.YFilter) { chsrpextifstandbyentry.YFilter = yf }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cHsrpGrpNumber" { return "Chsrpgrpnumber" }
    if yname == "cHsrpExtIfStandbyIndex" { return "Chsrpextifstandbyindex" }
    if yname == "cHsrpExtIfStandbyDestAddrType" { return "Chsrpextifstandbydestaddrtype" }
    if yname == "cHsrpExtIfStandbyDestAddr" { return "Chsrpextifstandbydestaddr" }
    if yname == "cHsrpExtIfStandbySourceAddrType" { return "Chsrpextifstandbysourceaddrtype" }
    if yname == "cHsrpExtIfStandbySourceAddr" { return "Chsrpextifstandbysourceaddr" }
    if yname == "cHsrpExtIfStandbyRowStatus" { return "Chsrpextifstandbyrowstatus" }
    return ""
}

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetSegmentPath() string {
    return "cHsrpExtIfStandbyEntry" + "[ifIndex='" + fmt.Sprintf("%v", chsrpextifstandbyentry.Ifindex) + "']" + "[cHsrpGrpNumber='" + fmt.Sprintf("%v", chsrpextifstandbyentry.Chsrpgrpnumber) + "']" + "[cHsrpExtIfStandbyIndex='" + fmt.Sprintf("%v", chsrpextifstandbyentry.Chsrpextifstandbyindex) + "']"
}

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = chsrpextifstandbyentry.Ifindex
    leafs["cHsrpGrpNumber"] = chsrpextifstandbyentry.Chsrpgrpnumber
    leafs["cHsrpExtIfStandbyIndex"] = chsrpextifstandbyentry.Chsrpextifstandbyindex
    leafs["cHsrpExtIfStandbyDestAddrType"] = chsrpextifstandbyentry.Chsrpextifstandbydestaddrtype
    leafs["cHsrpExtIfStandbyDestAddr"] = chsrpextifstandbyentry.Chsrpextifstandbydestaddr
    leafs["cHsrpExtIfStandbySourceAddrType"] = chsrpextifstandbyentry.Chsrpextifstandbysourceaddrtype
    leafs["cHsrpExtIfStandbySourceAddr"] = chsrpextifstandbyentry.Chsrpextifstandbysourceaddr
    leafs["cHsrpExtIfStandbyRowStatus"] = chsrpextifstandbyentry.Chsrpextifstandbyrowstatus
    return leafs
}

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetYangName() string { return "cHsrpExtIfStandbyEntry" }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) SetParent(parent types.Entity) { chsrpextifstandbyentry.parent = parent }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetParent() types.Entity { return chsrpextifstandbyentry.parent }

func (chsrpextifstandbyentry *CISCOHSRPEXTMIB_Chsrpextifstandbytable_Chsrpextifstandbyentry) GetParentYangName() string { return "cHsrpExtIfStandbyTable" }

// CISCOHSRPEXTMIB_Chsrpextiftable
// HSRP-specific configurations for each physical interface.
type CISCOHSRPEXTMIB_Chsrpextiftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If HSRP entries on this interface must use the BIA (Burned In Address),
    // there must be an entry for the interface in this  table. Entries of this
    // table are only accessible if HSRP has  been enabled i.e entries can not be
    // created if HSRP is not enabled. Also, the interfaces should be of IEEE 802
    // ones (Ethernet, Token Ring, FDDI,VLAN, LANE, TR-LANE).  Setting
    // cHsrpExtIfRowStatus to active initiates the entry with default value for
    // cHsrpExtIfUseBIA as FALSE. The value of cHsrpExtIfRowStatus may be set to
    // destroy at any time. If the row is not initiated, it is similar to having
    // cHsrpExtIfUseBIA as FALSE.  Entries may not be created via SNMP without
    // explicitly setting cHsrpExtIfRowStatus to either 'createAndGo' or
    // 'createAndWait'.  Entries can be created and modified via the management
    // protocol or by the device's local management interface.  If the row is not
    // active, and a local management interface command modifies that row, the row
    // may transition to active state.  A row which is not in active state will
    // timeout after a configurable period (five minutes by default). This timeout
    // period can be changed by setting cHsrpConfigTimeout. The type is slice of
    // CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry.
    Chsrpextifentry []CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry
}

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetFilter() yfilter.YFilter { return chsrpextiftable.YFilter }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) SetFilter(yf yfilter.YFilter) { chsrpextiftable.YFilter = yf }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetGoName(yname string) string {
    if yname == "cHsrpExtIfEntry" { return "Chsrpextifentry" }
    return ""
}

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetSegmentPath() string {
    return "cHsrpExtIfTable"
}

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cHsrpExtIfEntry" {
        for _, c := range chsrpextiftable.Chsrpextifentry {
            if chsrpextiftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry{}
        chsrpextiftable.Chsrpextifentry = append(chsrpextiftable.Chsrpextifentry, child)
        return &chsrpextiftable.Chsrpextifentry[len(chsrpextiftable.Chsrpextifentry)-1]
    }
    return nil
}

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range chsrpextiftable.Chsrpextifentry {
        children[chsrpextiftable.Chsrpextifentry[i].GetSegmentPath()] = &chsrpextiftable.Chsrpextifentry[i]
    }
    return children
}

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetYangName() string { return "cHsrpExtIfTable" }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) SetParent(parent types.Entity) { chsrpextiftable.parent = parent }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetParent() types.Entity { return chsrpextiftable.parent }

func (chsrpextiftable *CISCOHSRPEXTMIB_Chsrpextiftable) GetParentYangName() string { return "CISCO-HSRP-EXT-MIB" }

// CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry
// If HSRP entries on this interface must use the BIA (Burned
// In Address), there must be an entry for the interface in this 
// table. Entries of this table are only accessible if HSRP has 
// been enabled i.e entries can not be created if HSRP is not
// enabled. Also, the interfaces should be of IEEE 802 ones
// (Ethernet, Token Ring, FDDI,VLAN, LANE, TR-LANE).
// 
// Setting cHsrpExtIfRowStatus to active initiates the
// entry with default value for cHsrpExtIfUseBIA as FALSE.
// The value of cHsrpExtIfRowStatus may be set to destroy
// at any time. If the row is not initiated, it is similar to
// having cHsrpExtIfUseBIA as FALSE.
// 
// Entries may not be created via SNMP without explicitly setting
// cHsrpExtIfRowStatus to either 'createAndGo' or 'createAndWait'.
// 
// Entries can be created and modified via the management
// protocol or by the device's local management interface.
// 
// If the row is not active, and a local management interface
// command modifies that row, the row may transition to active
// state.
// 
// A row which is not in active state will timeout after a
// configurable period (five minutes by default). This timeout
// period can be changed by setting cHsrpConfigTimeout.
type CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // If set to TRUE, the HSRP Group MAC Address for all groups on this 
    // interface will be the burned-in-address. Otherwise, this will be determined
    // by ciscoHsrpGroupNumber. In case of sub-interfaces, UseBIA applies to all
    // sub-interfaces on an  interface and to all groups on those sub-interfaces.
    // The type is bool.
    Chsrpextifusebia interface{}

    // The control that allows modification, creation, and deletion of entries.
    // For detailed rules see the DESCRIPTION for cHsrpExtIfEntry. The type is
    // RowStatus.
    Chsrpextifrowstatus interface{}
}

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetFilter() yfilter.YFilter { return chsrpextifentry.YFilter }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) SetFilter(yf yfilter.YFilter) { chsrpextifentry.YFilter = yf }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cHsrpExtIfUseBIA" { return "Chsrpextifusebia" }
    if yname == "cHsrpExtIfRowStatus" { return "Chsrpextifrowstatus" }
    return ""
}

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetSegmentPath() string {
    return "cHsrpExtIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", chsrpextifentry.Ifindex) + "']"
}

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = chsrpextifentry.Ifindex
    leafs["cHsrpExtIfUseBIA"] = chsrpextifentry.Chsrpextifusebia
    leafs["cHsrpExtIfRowStatus"] = chsrpextifentry.Chsrpextifrowstatus
    return leafs
}

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetYangName() string { return "cHsrpExtIfEntry" }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) SetParent(parent types.Entity) { chsrpextifentry.parent = parent }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetParent() types.Entity { return chsrpextifentry.parent }

func (chsrpextifentry *CISCOHSRPEXTMIB_Chsrpextiftable_Chsrpextifentry) GetParentYangName() string { return "cHsrpExtIfTable" }

