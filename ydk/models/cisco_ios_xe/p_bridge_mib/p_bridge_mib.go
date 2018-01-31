// The Bridge MIB Extension module for managing Priority
// and Multicast Filtering, defined by IEEE 802.1D-1998,
// including Restricted Group Registration defined by
// IEEE 802.1t-2001.
// 
// Copyright (C) The Internet Society (2006).  This version of
// this MIB module is part of RFC 4363; See the RFC itself for
// full legal notices.
package p_bridge_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package p_bridge_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:P-BRIDGE-MIB P-BRIDGE-MIB}", reflect.TypeOf(PBRIDGEMIB{}))
    ydk.RegisterEntity("P-BRIDGE-MIB:P-BRIDGE-MIB", reflect.TypeOf(PBRIDGEMIB{}))
}

// EnabledStatus represents A simple status value for the object.
type EnabledStatus string

const (
    EnabledStatus_enabled EnabledStatus = "enabled"

    EnabledStatus_disabled EnabledStatus = "disabled"
)

// PBRIDGEMIB
type PBRIDGEMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Dot1Dextbase PBRIDGEMIB_Dot1Dextbase

    // A table that contains information about every high- capacity port that is
    // associated with this transparent bridge.
    Dot1Dtphcporttable PBRIDGEMIB_Dot1Dtphcporttable

    // A table that contains the most-significant bits of statistics counters for
    // ports that are associated with this transparent bridge that are on
    // high-capacity interfaces, as defined in the conformance clauses for this
    // table.  This table is provided as a way to read 64-bit counters for agents
    // that support only SNMPv1.  Note that the reporting of most-significant and
    // least-significant counter bits separately runs the risk of missing an
    // overflow of the lower bits in the interval between sampling.  The manager
    // must be aware of this possibility, even within the same varbindlist, when
    // interpreting the results of a request or asynchronous notification.
    Dot1Dtpportoverflowtable PBRIDGEMIB_Dot1Dtpportoverflowtable

    // A list of Regenerated User Priorities for each received User Priority on
    // each port of a bridge.  The Regenerated User Priority value may be used to
    // index the Traffic Class Table for each input port.  This only has effect on
    // media that support native User Priority.  The default values for
    // Regenerated User Priorities are the same as the User Priorities.
    Dot1Duserpriorityregentable PBRIDGEMIB_Dot1Duserpriorityregentable

    // A table mapping evaluated User Priority to Traffic Class, for forwarding by
    // the bridge.  Traffic class is a number in the range
    // (0..(dot1dPortNumTrafficClasses-1)).
    Dot1Dtrafficclasstable PBRIDGEMIB_Dot1Dtrafficclasstable

    // A table mapping Regenerated User Priority to Outbound Access Priority. 
    // This is a fixed mapping for all port types, with two options for 802.5
    // Token Ring.
    Dot1Dportoutboundaccessprioritytable PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable
}

func (pBRIDGEMIB *PBRIDGEMIB) GetFilter() yfilter.YFilter { return pBRIDGEMIB.YFilter }

func (pBRIDGEMIB *PBRIDGEMIB) SetFilter(yf yfilter.YFilter) { pBRIDGEMIB.YFilter = yf }

func (pBRIDGEMIB *PBRIDGEMIB) GetGoName(yname string) string {
    if yname == "dot1dExtBase" { return "Dot1Dextbase" }
    if yname == "dot1dTpHCPortTable" { return "Dot1Dtphcporttable" }
    if yname == "dot1dTpPortOverflowTable" { return "Dot1Dtpportoverflowtable" }
    if yname == "dot1dUserPriorityRegenTable" { return "Dot1Duserpriorityregentable" }
    if yname == "dot1dTrafficClassTable" { return "Dot1Dtrafficclasstable" }
    if yname == "dot1dPortOutboundAccessPriorityTable" { return "Dot1Dportoutboundaccessprioritytable" }
    return ""
}

func (pBRIDGEMIB *PBRIDGEMIB) GetSegmentPath() string {
    return "P-BRIDGE-MIB:P-BRIDGE-MIB"
}

func (pBRIDGEMIB *PBRIDGEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dExtBase" {
        return &pBRIDGEMIB.Dot1Dextbase
    }
    if childYangName == "dot1dTpHCPortTable" {
        return &pBRIDGEMIB.Dot1Dtphcporttable
    }
    if childYangName == "dot1dTpPortOverflowTable" {
        return &pBRIDGEMIB.Dot1Dtpportoverflowtable
    }
    if childYangName == "dot1dUserPriorityRegenTable" {
        return &pBRIDGEMIB.Dot1Duserpriorityregentable
    }
    if childYangName == "dot1dTrafficClassTable" {
        return &pBRIDGEMIB.Dot1Dtrafficclasstable
    }
    if childYangName == "dot1dPortOutboundAccessPriorityTable" {
        return &pBRIDGEMIB.Dot1Dportoutboundaccessprioritytable
    }
    return nil
}

func (pBRIDGEMIB *PBRIDGEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dot1dExtBase"] = &pBRIDGEMIB.Dot1Dextbase
    children["dot1dTpHCPortTable"] = &pBRIDGEMIB.Dot1Dtphcporttable
    children["dot1dTpPortOverflowTable"] = &pBRIDGEMIB.Dot1Dtpportoverflowtable
    children["dot1dUserPriorityRegenTable"] = &pBRIDGEMIB.Dot1Duserpriorityregentable
    children["dot1dTrafficClassTable"] = &pBRIDGEMIB.Dot1Dtrafficclasstable
    children["dot1dPortOutboundAccessPriorityTable"] = &pBRIDGEMIB.Dot1Dportoutboundaccessprioritytable
    return children
}

func (pBRIDGEMIB *PBRIDGEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pBRIDGEMIB *PBRIDGEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (pBRIDGEMIB *PBRIDGEMIB) GetYangName() string { return "P-BRIDGE-MIB" }

func (pBRIDGEMIB *PBRIDGEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pBRIDGEMIB *PBRIDGEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pBRIDGEMIB *PBRIDGEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pBRIDGEMIB *PBRIDGEMIB) SetParent(parent types.Entity) { pBRIDGEMIB.parent = parent }

func (pBRIDGEMIB *PBRIDGEMIB) GetParent() types.Entity { return pBRIDGEMIB.parent }

func (pBRIDGEMIB *PBRIDGEMIB) GetParentYangName() string { return "P-BRIDGE-MIB" }

// PBRIDGEMIB_Dot1Dextbase
type PBRIDGEMIB_Dot1Dextbase struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates the optional parts of IEEE 802.1D and 802.1Q that are implemented
    // by this device and are manageable through this MIB.  Capabilities that are
    // allowed on a per-port basis are indicated in dot1dPortCapabilities. 
    // dot1dExtendedFilteringServices(0),                       -- can perform
    // filtering of                       -- individual multicast addresses       
    // -- controlled by GMRP. dot1dTrafficClasses(1),                       -- can
    // map user priority to                       -- multiple traffic classes.
    // dot1qStaticEntryIndividualPort(2),                       --
    // dot1qStaticUnicastReceivePort &                       --
    // dot1qStaticMulticastReceivePort                       -- can represent
    // non-zero entries. dot1qIVLCapable(3),   -- Independent VLAN Learning (IVL).
    // dot1qSVLCapable(4),   -- Shared VLAN Learning (SVL). dot1qHybridCapable(5),
    // -- both IVL & SVL simultaneously. dot1qConfigurablePvidTagging(6),         
    // -- whether the implementation                       -- supports the ability
    // to                       -- override the default PVID                      
    // -- setting and its egress status                       -- (VLAN-Tagged or
    // Untagged) on                       -- each port. dot1dLocalVlanCapable(7)  
    // -- can support multiple local                       -- bridges, outside of
    // the scope                       -- of 802.1Q defined VLANs. The type is
    // map[string]bool.
    Dot1Ddevicecapabilities interface{}

    // The value true(1) indicates that Traffic Classes are enabled on this
    // bridge.  When false(2), the bridge operates with a single priority level
    // for all traffic.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is bool.
    Dot1Dtrafficclassesenabled interface{}

    // The administrative status requested by management for GMRP.  The value
    // enabled(1) indicates that GMRP should be enabled on this device, in all
    // VLANs, on all ports for which it has not been specifically disabled.  When
    // disabled(2), GMRP is disabled, in all VLANs and on all ports, and all GMRP
    // packets will be forwarded transparently.  This object affects both
    // Applicant and Registrar state machines.  A transition from disabled(2) to
    // enabled(1) will cause a reset of all GMRP state machines on all ports.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is EnabledStatus.
    Dot1Dgmrpstatus interface{}
}

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetFilter() yfilter.YFilter { return dot1Dextbase.YFilter }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) SetFilter(yf yfilter.YFilter) { dot1Dextbase.YFilter = yf }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetGoName(yname string) string {
    if yname == "dot1dDeviceCapabilities" { return "Dot1Ddevicecapabilities" }
    if yname == "dot1dTrafficClassesEnabled" { return "Dot1Dtrafficclassesenabled" }
    if yname == "dot1dGmrpStatus" { return "Dot1Dgmrpstatus" }
    return ""
}

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetSegmentPath() string {
    return "dot1dExtBase"
}

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dDeviceCapabilities"] = dot1Dextbase.Dot1Ddevicecapabilities
    leafs["dot1dTrafficClassesEnabled"] = dot1Dextbase.Dot1Dtrafficclassesenabled
    leafs["dot1dGmrpStatus"] = dot1Dextbase.Dot1Dgmrpstatus
    return leafs
}

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetYangName() string { return "dot1dExtBase" }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) SetParent(parent types.Entity) { dot1Dextbase.parent = parent }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetParent() types.Entity { return dot1Dextbase.parent }

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetParentYangName() string { return "P-BRIDGE-MIB" }

// PBRIDGEMIB_Dot1Dtphcporttable
// A table that contains information about every high-
// capacity port that is associated with this transparent
// bridge.
type PBRIDGEMIB_Dot1Dtphcporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics information for each high-capacity port of a transparent bridge.
    // The type is slice of PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry.
    Dot1Dtphcportentry []PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry
}

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetFilter() yfilter.YFilter { return dot1Dtphcporttable.YFilter }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) SetFilter(yf yfilter.YFilter) { dot1Dtphcporttable.YFilter = yf }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetGoName(yname string) string {
    if yname == "dot1dTpHCPortEntry" { return "Dot1Dtphcportentry" }
    return ""
}

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetSegmentPath() string {
    return "dot1dTpHCPortTable"
}

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dTpHCPortEntry" {
        for _, c := range dot1Dtphcporttable.Dot1Dtphcportentry {
            if dot1Dtphcporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry{}
        dot1Dtphcporttable.Dot1Dtphcportentry = append(dot1Dtphcporttable.Dot1Dtphcportentry, child)
        return &dot1Dtphcporttable.Dot1Dtphcportentry[len(dot1Dtphcporttable.Dot1Dtphcportentry)-1]
    }
    return nil
}

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dtphcporttable.Dot1Dtphcportentry {
        children[dot1Dtphcporttable.Dot1Dtphcportentry[i].GetSegmentPath()] = &dot1Dtphcporttable.Dot1Dtphcportentry[i]
    }
    return children
}

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetYangName() string { return "dot1dTpHCPortTable" }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) SetParent(parent types.Entity) { dot1Dtphcporttable.parent = parent }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetParent() types.Entity { return dot1Dtphcporttable.parent }

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetParentYangName() string { return "P-BRIDGE-MIB" }

// PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry
// Statistics information for each high-capacity port of a
// transparent bridge.
type PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry_Dot1Dtpport
    Dot1Dtpport interface{}

    // The number of frames that have been received by this port from its segment.
    // Note that a frame received on the interface corresponding to this port is
    // only counted by this object if and only if it is for a protocol being
    // processed by the local bridging function, including bridge management
    // frames. The type is interface{} with range: 0..18446744073709551615.
    Dot1Dtphcportinframes interface{}

    // The number of frames that have been transmitted by this port to its
    // segment.  Note that a frame transmitted on the interface corresponding to
    // this port is only counted by this object if and only if it is for a
    // protocol being processed by the local bridging function, including bridge
    // management frames. The type is interface{} with range:
    // 0..18446744073709551615.
    Dot1Dtphcportoutframes interface{}

    // Count of valid frames that have been received by this port from its segment
    // that were discarded (i.e., filtered) by the Forwarding Process. The type is
    // interface{} with range: 0..18446744073709551615.
    Dot1Dtphcportindiscards interface{}
}

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetFilter() yfilter.YFilter { return dot1Dtphcportentry.YFilter }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) SetFilter(yf yfilter.YFilter) { dot1Dtphcportentry.YFilter = yf }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetGoName(yname string) string {
    if yname == "dot1dTpPort" { return "Dot1Dtpport" }
    if yname == "dot1dTpHCPortInFrames" { return "Dot1Dtphcportinframes" }
    if yname == "dot1dTpHCPortOutFrames" { return "Dot1Dtphcportoutframes" }
    if yname == "dot1dTpHCPortInDiscards" { return "Dot1Dtphcportindiscards" }
    return ""
}

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetSegmentPath() string {
    return "dot1dTpHCPortEntry" + "[dot1dTpPort='" + fmt.Sprintf("%v", dot1Dtphcportentry.Dot1Dtpport) + "']"
}

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dTpPort"] = dot1Dtphcportentry.Dot1Dtpport
    leafs["dot1dTpHCPortInFrames"] = dot1Dtphcportentry.Dot1Dtphcportinframes
    leafs["dot1dTpHCPortOutFrames"] = dot1Dtphcportentry.Dot1Dtphcportoutframes
    leafs["dot1dTpHCPortInDiscards"] = dot1Dtphcportentry.Dot1Dtphcportindiscards
    return leafs
}

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetYangName() string { return "dot1dTpHCPortEntry" }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) SetParent(parent types.Entity) { dot1Dtphcportentry.parent = parent }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetParent() types.Entity { return dot1Dtphcportentry.parent }

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetParentYangName() string { return "dot1dTpHCPortTable" }

// PBRIDGEMIB_Dot1Dtpportoverflowtable
// A table that contains the most-significant bits of
// statistics counters for ports that are associated with this
// transparent bridge that are on high-capacity interfaces, as
// defined in the conformance clauses for this table.  This table
// is provided as a way to read 64-bit counters for agents that
// support only SNMPv1.
// 
// Note that the reporting of most-significant and
// least-significant counter bits separately runs the risk of
// missing an overflow of the lower bits in the interval between
// sampling.  The manager must be aware of this possibility, even
// within the same varbindlist, when interpreting the results of
// a request or asynchronous notification.
type PBRIDGEMIB_Dot1Dtpportoverflowtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The most significant bits of statistics counters for a high- capacity
    // interface of a transparent bridge.  Each object is associated with a
    // corresponding object in dot1dTpPortTable that indicates the least
    // significant bits of the counter. The type is slice of
    // PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry.
    Dot1Dtpportoverflowentry []PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry
}

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetFilter() yfilter.YFilter { return dot1Dtpportoverflowtable.YFilter }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) SetFilter(yf yfilter.YFilter) { dot1Dtpportoverflowtable.YFilter = yf }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetGoName(yname string) string {
    if yname == "dot1dTpPortOverflowEntry" { return "Dot1Dtpportoverflowentry" }
    return ""
}

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetSegmentPath() string {
    return "dot1dTpPortOverflowTable"
}

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dTpPortOverflowEntry" {
        for _, c := range dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry {
            if dot1Dtpportoverflowtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry{}
        dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry = append(dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry, child)
        return &dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry[len(dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry)-1]
    }
    return nil
}

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry {
        children[dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry[i].GetSegmentPath()] = &dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry[i]
    }
    return children
}

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetYangName() string { return "dot1dTpPortOverflowTable" }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) SetParent(parent types.Entity) { dot1Dtpportoverflowtable.parent = parent }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetParent() types.Entity { return dot1Dtpportoverflowtable.parent }

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetParentYangName() string { return "P-BRIDGE-MIB" }

// PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry
// The most significant bits of statistics counters for a high-
// capacity interface of a transparent bridge.  Each object is
// associated with a corresponding object in dot1dTpPortTable
// that indicates the least significant bits of the counter.
type PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dtpporttable_Dot1Dtpportentry_Dot1Dtpport
    Dot1Dtpport interface{}

    // The number of times the associated dot1dTpPortInFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1Dtpportinoverflowframes interface{}

    // The number of times the associated dot1dTpPortOutFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1Dtpportoutoverflowframes interface{}

    // The number of times the associated dot1dTpPortInDiscards counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1Dtpportinoverflowdiscards interface{}
}

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetFilter() yfilter.YFilter { return dot1Dtpportoverflowentry.YFilter }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) SetFilter(yf yfilter.YFilter) { dot1Dtpportoverflowentry.YFilter = yf }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetGoName(yname string) string {
    if yname == "dot1dTpPort" { return "Dot1Dtpport" }
    if yname == "dot1dTpPortInOverflowFrames" { return "Dot1Dtpportinoverflowframes" }
    if yname == "dot1dTpPortOutOverflowFrames" { return "Dot1Dtpportoutoverflowframes" }
    if yname == "dot1dTpPortInOverflowDiscards" { return "Dot1Dtpportinoverflowdiscards" }
    return ""
}

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetSegmentPath() string {
    return "dot1dTpPortOverflowEntry" + "[dot1dTpPort='" + fmt.Sprintf("%v", dot1Dtpportoverflowentry.Dot1Dtpport) + "']"
}

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dTpPort"] = dot1Dtpportoverflowentry.Dot1Dtpport
    leafs["dot1dTpPortInOverflowFrames"] = dot1Dtpportoverflowentry.Dot1Dtpportinoverflowframes
    leafs["dot1dTpPortOutOverflowFrames"] = dot1Dtpportoverflowentry.Dot1Dtpportoutoverflowframes
    leafs["dot1dTpPortInOverflowDiscards"] = dot1Dtpportoverflowentry.Dot1Dtpportinoverflowdiscards
    return leafs
}

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetYangName() string { return "dot1dTpPortOverflowEntry" }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) SetParent(parent types.Entity) { dot1Dtpportoverflowentry.parent = parent }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetParent() types.Entity { return dot1Dtpportoverflowentry.parent }

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetParentYangName() string { return "dot1dTpPortOverflowTable" }

// PBRIDGEMIB_Dot1Duserpriorityregentable
// A list of Regenerated User Priorities for each received
// User Priority on each port of a bridge.  The Regenerated
// User Priority value may be used to index the Traffic
// Class Table for each input port.  This only has effect
// on media that support native User Priority.  The default
// values for Regenerated User Priorities are the same as
// the User Priorities.
type PBRIDGEMIB_Dot1Duserpriorityregentable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A mapping of incoming User Priority to a Regenerated User Priority. The
    // type is slice of
    // PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry.
    Dot1Duserpriorityregenentry []PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry
}

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetFilter() yfilter.YFilter { return dot1Duserpriorityregentable.YFilter }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) SetFilter(yf yfilter.YFilter) { dot1Duserpriorityregentable.YFilter = yf }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetGoName(yname string) string {
    if yname == "dot1dUserPriorityRegenEntry" { return "Dot1Duserpriorityregenentry" }
    return ""
}

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetSegmentPath() string {
    return "dot1dUserPriorityRegenTable"
}

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dUserPriorityRegenEntry" {
        for _, c := range dot1Duserpriorityregentable.Dot1Duserpriorityregenentry {
            if dot1Duserpriorityregentable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry{}
        dot1Duserpriorityregentable.Dot1Duserpriorityregenentry = append(dot1Duserpriorityregentable.Dot1Duserpriorityregenentry, child)
        return &dot1Duserpriorityregentable.Dot1Duserpriorityregenentry[len(dot1Duserpriorityregentable.Dot1Duserpriorityregenentry)-1]
    }
    return nil
}

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Duserpriorityregentable.Dot1Duserpriorityregenentry {
        children[dot1Duserpriorityregentable.Dot1Duserpriorityregenentry[i].GetSegmentPath()] = &dot1Duserpriorityregentable.Dot1Duserpriorityregenentry[i]
    }
    return children
}

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetYangName() string { return "dot1dUserPriorityRegenTable" }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) SetParent(parent types.Entity) { dot1Duserpriorityregentable.parent = parent }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetParent() types.Entity { return dot1Duserpriorityregentable.parent }

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetParentYangName() string { return "P-BRIDGE-MIB" }

// PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry
// A mapping of incoming User Priority to a Regenerated
// User Priority.
type PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Dbaseport
    Dot1Dbaseport interface{}

    // This attribute is a key. The User Priority for a frame received on this
    // port. The type is interface{} with range: 0..7.
    Dot1Duserpriority interface{}

    // The Regenerated User Priority that the incoming User  Priority is mapped to
    // for this port.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is interface{} with
    // range: 0..7.
    Dot1Dregenuserpriority interface{}
}

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetFilter() yfilter.YFilter { return dot1Duserpriorityregenentry.YFilter }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) SetFilter(yf yfilter.YFilter) { dot1Duserpriorityregenentry.YFilter = yf }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetGoName(yname string) string {
    if yname == "dot1dBasePort" { return "Dot1Dbaseport" }
    if yname == "dot1dUserPriority" { return "Dot1Duserpriority" }
    if yname == "dot1dRegenUserPriority" { return "Dot1Dregenuserpriority" }
    return ""
}

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetSegmentPath() string {
    return "dot1dUserPriorityRegenEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Duserpriorityregenentry.Dot1Dbaseport) + "']" + "[dot1dUserPriority='" + fmt.Sprintf("%v", dot1Duserpriorityregenentry.Dot1Duserpriority) + "']"
}

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBasePort"] = dot1Duserpriorityregenentry.Dot1Dbaseport
    leafs["dot1dUserPriority"] = dot1Duserpriorityregenentry.Dot1Duserpriority
    leafs["dot1dRegenUserPriority"] = dot1Duserpriorityregenentry.Dot1Dregenuserpriority
    return leafs
}

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetYangName() string { return "dot1dUserPriorityRegenEntry" }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) SetParent(parent types.Entity) { dot1Duserpriorityregenentry.parent = parent }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetParent() types.Entity { return dot1Duserpriorityregenentry.parent }

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetParentYangName() string { return "dot1dUserPriorityRegenTable" }

// PBRIDGEMIB_Dot1Dtrafficclasstable
// A table mapping evaluated User Priority to Traffic
// Class, for forwarding by the bridge.  Traffic class is a
// number in the range (0..(dot1dPortNumTrafficClasses-1)).
type PBRIDGEMIB_Dot1Dtrafficclasstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // User Priority to Traffic Class mapping. The type is slice of
    // PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry.
    Dot1Dtrafficclassentry []PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry
}

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetFilter() yfilter.YFilter { return dot1Dtrafficclasstable.YFilter }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) SetFilter(yf yfilter.YFilter) { dot1Dtrafficclasstable.YFilter = yf }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetGoName(yname string) string {
    if yname == "dot1dTrafficClassEntry" { return "Dot1Dtrafficclassentry" }
    return ""
}

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetSegmentPath() string {
    return "dot1dTrafficClassTable"
}

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dTrafficClassEntry" {
        for _, c := range dot1Dtrafficclasstable.Dot1Dtrafficclassentry {
            if dot1Dtrafficclasstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry{}
        dot1Dtrafficclasstable.Dot1Dtrafficclassentry = append(dot1Dtrafficclasstable.Dot1Dtrafficclassentry, child)
        return &dot1Dtrafficclasstable.Dot1Dtrafficclassentry[len(dot1Dtrafficclasstable.Dot1Dtrafficclassentry)-1]
    }
    return nil
}

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dtrafficclasstable.Dot1Dtrafficclassentry {
        children[dot1Dtrafficclasstable.Dot1Dtrafficclassentry[i].GetSegmentPath()] = &dot1Dtrafficclasstable.Dot1Dtrafficclassentry[i]
    }
    return children
}

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetYangName() string { return "dot1dTrafficClassTable" }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) SetParent(parent types.Entity) { dot1Dtrafficclasstable.parent = parent }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetParent() types.Entity { return dot1Dtrafficclasstable.parent }

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetParentYangName() string { return "P-BRIDGE-MIB" }

// PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry
// User Priority to Traffic Class mapping.
type PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Dbaseport
    Dot1Dbaseport interface{}

    // This attribute is a key. The Priority value determined for the received
    // frame. This value is equivalent to the priority indicated in the tagged
    // frame received, or one of the evaluated priorities, determined according to
    // the media-type.  For untagged frames received from Ethernet media, this
    // value is equal to the dot1dPortDefaultUserPriority value for the ingress
    // port.  For untagged frames received from non-Ethernet media, this value is
    // equal to the dot1dRegenUserPriority value for the ingress port and
    // media-specific user priority. The type is interface{} with range: 0..7.
    Dot1Dtrafficclasspriority interface{}

    // The Traffic Class the received frame is mapped to.  The value of this
    // object MUST be retained across reinitializations of the management system.
    // The type is interface{} with range: 0..7.
    Dot1Dtrafficclass interface{}
}

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetFilter() yfilter.YFilter { return dot1Dtrafficclassentry.YFilter }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) SetFilter(yf yfilter.YFilter) { dot1Dtrafficclassentry.YFilter = yf }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetGoName(yname string) string {
    if yname == "dot1dBasePort" { return "Dot1Dbaseport" }
    if yname == "dot1dTrafficClassPriority" { return "Dot1Dtrafficclasspriority" }
    if yname == "dot1dTrafficClass" { return "Dot1Dtrafficclass" }
    return ""
}

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetSegmentPath() string {
    return "dot1dTrafficClassEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Dtrafficclassentry.Dot1Dbaseport) + "']" + "[dot1dTrafficClassPriority='" + fmt.Sprintf("%v", dot1Dtrafficclassentry.Dot1Dtrafficclasspriority) + "']"
}

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBasePort"] = dot1Dtrafficclassentry.Dot1Dbaseport
    leafs["dot1dTrafficClassPriority"] = dot1Dtrafficclassentry.Dot1Dtrafficclasspriority
    leafs["dot1dTrafficClass"] = dot1Dtrafficclassentry.Dot1Dtrafficclass
    return leafs
}

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetYangName() string { return "dot1dTrafficClassEntry" }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) SetParent(parent types.Entity) { dot1Dtrafficclassentry.parent = parent }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetParent() types.Entity { return dot1Dtrafficclassentry.parent }

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetParentYangName() string { return "dot1dTrafficClassTable" }

// PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable
// A table mapping Regenerated User Priority to Outbound
// Access Priority.  This is a fixed mapping for all port
// types, with two options for 802.5 Token Ring.
type PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Regenerated User Priority to Outbound Access Priority mapping. The type is
    // slice of
    // PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry.
    Dot1Dportoutboundaccesspriorityentry []PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry
}

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetFilter() yfilter.YFilter { return dot1Dportoutboundaccessprioritytable.YFilter }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) SetFilter(yf yfilter.YFilter) { dot1Dportoutboundaccessprioritytable.YFilter = yf }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetGoName(yname string) string {
    if yname == "dot1dPortOutboundAccessPriorityEntry" { return "Dot1Dportoutboundaccesspriorityentry" }
    return ""
}

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetSegmentPath() string {
    return "dot1dPortOutboundAccessPriorityTable"
}

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1dPortOutboundAccessPriorityEntry" {
        for _, c := range dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry {
            if dot1Dportoutboundaccessprioritytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry{}
        dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry = append(dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry, child)
        return &dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry[len(dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry)-1]
    }
    return nil
}

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry {
        children[dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry[i].GetSegmentPath()] = &dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry[i]
    }
    return children
}

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetYangName() string { return "dot1dPortOutboundAccessPriorityTable" }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) SetParent(parent types.Entity) { dot1Dportoutboundaccessprioritytable.parent = parent }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetParent() types.Entity { return dot1Dportoutboundaccessprioritytable.parent }

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetParentYangName() string { return "P-BRIDGE-MIB" }

// PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry
// Regenerated User Priority to Outbound Access Priority
// mapping.
type PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Dbaseport
    Dot1Dbaseport interface{}

    // This attribute is a key. The type is string with range: 0..7. Refers to
    // p_bridge_mib.PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry_Dot1Dregenuserpriority
    Dot1Dregenuserpriority interface{}

    // The Outbound Access Priority the received frame is mapped to. The type is
    // interface{} with range: 0..7.
    Dot1Dportoutboundaccesspriority interface{}
}

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetFilter() yfilter.YFilter { return dot1Dportoutboundaccesspriorityentry.YFilter }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) SetFilter(yf yfilter.YFilter) { dot1Dportoutboundaccesspriorityentry.YFilter = yf }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetGoName(yname string) string {
    if yname == "dot1dBasePort" { return "Dot1Dbaseport" }
    if yname == "dot1dRegenUserPriority" { return "Dot1Dregenuserpriority" }
    if yname == "dot1dPortOutboundAccessPriority" { return "Dot1Dportoutboundaccesspriority" }
    return ""
}

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetSegmentPath() string {
    return "dot1dPortOutboundAccessPriorityEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Dportoutboundaccesspriorityentry.Dot1Dbaseport) + "']" + "[dot1dRegenUserPriority='" + fmt.Sprintf("%v", dot1Dportoutboundaccesspriorityentry.Dot1Dregenuserpriority) + "']"
}

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBasePort"] = dot1Dportoutboundaccesspriorityentry.Dot1Dbaseport
    leafs["dot1dRegenUserPriority"] = dot1Dportoutboundaccesspriorityentry.Dot1Dregenuserpriority
    leafs["dot1dPortOutboundAccessPriority"] = dot1Dportoutboundaccesspriorityentry.Dot1Dportoutboundaccesspriority
    return leafs
}

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetYangName() string { return "dot1dPortOutboundAccessPriorityEntry" }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) SetParent(parent types.Entity) { dot1Dportoutboundaccesspriorityentry.parent = parent }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetParent() types.Entity { return dot1Dportoutboundaccesspriorityentry.parent }

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetParentYangName() string { return "dot1dPortOutboundAccessPriorityTable" }

