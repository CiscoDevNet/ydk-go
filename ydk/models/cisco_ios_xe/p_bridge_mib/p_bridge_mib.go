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
    EntityData types.CommonEntityData
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

func (pBRIDGEMIB *PBRIDGEMIB) GetEntityData() *types.CommonEntityData {
    pBRIDGEMIB.EntityData.YFilter = pBRIDGEMIB.YFilter
    pBRIDGEMIB.EntityData.YangName = "P-BRIDGE-MIB"
    pBRIDGEMIB.EntityData.BundleName = "cisco_ios_xe"
    pBRIDGEMIB.EntityData.ParentYangName = "P-BRIDGE-MIB"
    pBRIDGEMIB.EntityData.SegmentPath = "P-BRIDGE-MIB:P-BRIDGE-MIB"
    pBRIDGEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pBRIDGEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pBRIDGEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pBRIDGEMIB.EntityData.Children = make(map[string]types.YChild)
    pBRIDGEMIB.EntityData.Children["dot1dExtBase"] = types.YChild{"Dot1Dextbase", &pBRIDGEMIB.Dot1Dextbase}
    pBRIDGEMIB.EntityData.Children["dot1dTpHCPortTable"] = types.YChild{"Dot1Dtphcporttable", &pBRIDGEMIB.Dot1Dtphcporttable}
    pBRIDGEMIB.EntityData.Children["dot1dTpPortOverflowTable"] = types.YChild{"Dot1Dtpportoverflowtable", &pBRIDGEMIB.Dot1Dtpportoverflowtable}
    pBRIDGEMIB.EntityData.Children["dot1dUserPriorityRegenTable"] = types.YChild{"Dot1Duserpriorityregentable", &pBRIDGEMIB.Dot1Duserpriorityregentable}
    pBRIDGEMIB.EntityData.Children["dot1dTrafficClassTable"] = types.YChild{"Dot1Dtrafficclasstable", &pBRIDGEMIB.Dot1Dtrafficclasstable}
    pBRIDGEMIB.EntityData.Children["dot1dPortOutboundAccessPriorityTable"] = types.YChild{"Dot1Dportoutboundaccessprioritytable", &pBRIDGEMIB.Dot1Dportoutboundaccessprioritytable}
    pBRIDGEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pBRIDGEMIB.EntityData)
}

// PBRIDGEMIB_Dot1Dextbase
type PBRIDGEMIB_Dot1Dextbase struct {
    EntityData types.CommonEntityData
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

func (dot1Dextbase *PBRIDGEMIB_Dot1Dextbase) GetEntityData() *types.CommonEntityData {
    dot1Dextbase.EntityData.YFilter = dot1Dextbase.YFilter
    dot1Dextbase.EntityData.YangName = "dot1dExtBase"
    dot1Dextbase.EntityData.BundleName = "cisco_ios_xe"
    dot1Dextbase.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1Dextbase.EntityData.SegmentPath = "dot1dExtBase"
    dot1Dextbase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dextbase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dextbase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dextbase.EntityData.Children = make(map[string]types.YChild)
    dot1Dextbase.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dextbase.EntityData.Leafs["dot1dDeviceCapabilities"] = types.YLeaf{"Dot1Ddevicecapabilities", dot1Dextbase.Dot1Ddevicecapabilities}
    dot1Dextbase.EntityData.Leafs["dot1dTrafficClassesEnabled"] = types.YLeaf{"Dot1Dtrafficclassesenabled", dot1Dextbase.Dot1Dtrafficclassesenabled}
    dot1Dextbase.EntityData.Leafs["dot1dGmrpStatus"] = types.YLeaf{"Dot1Dgmrpstatus", dot1Dextbase.Dot1Dgmrpstatus}
    return &(dot1Dextbase.EntityData)
}

// PBRIDGEMIB_Dot1Dtphcporttable
// A table that contains information about every high-
// capacity port that is associated with this transparent
// bridge.
type PBRIDGEMIB_Dot1Dtphcporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics information for each high-capacity port of a transparent bridge.
    // The type is slice of PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry.
    Dot1Dtphcportentry []PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry
}

func (dot1Dtphcporttable *PBRIDGEMIB_Dot1Dtphcporttable) GetEntityData() *types.CommonEntityData {
    dot1Dtphcporttable.EntityData.YFilter = dot1Dtphcporttable.YFilter
    dot1Dtphcporttable.EntityData.YangName = "dot1dTpHCPortTable"
    dot1Dtphcporttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtphcporttable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1Dtphcporttable.EntityData.SegmentPath = "dot1dTpHCPortTable"
    dot1Dtphcporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtphcporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtphcporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtphcporttable.EntityData.Children = make(map[string]types.YChild)
    dot1Dtphcporttable.EntityData.Children["dot1dTpHCPortEntry"] = types.YChild{"Dot1Dtphcportentry", nil}
    for i := range dot1Dtphcporttable.Dot1Dtphcportentry {
        dot1Dtphcporttable.EntityData.Children[types.GetSegmentPath(&dot1Dtphcporttable.Dot1Dtphcportentry[i])] = types.YChild{"Dot1Dtphcportentry", &dot1Dtphcporttable.Dot1Dtphcportentry[i]}
    }
    dot1Dtphcporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dtphcporttable.EntityData)
}

// PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry
// Statistics information for each high-capacity port of a
// transparent bridge.
type PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry struct {
    EntityData types.CommonEntityData
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

func (dot1Dtphcportentry *PBRIDGEMIB_Dot1Dtphcporttable_Dot1Dtphcportentry) GetEntityData() *types.CommonEntityData {
    dot1Dtphcportentry.EntityData.YFilter = dot1Dtphcportentry.YFilter
    dot1Dtphcportentry.EntityData.YangName = "dot1dTpHCPortEntry"
    dot1Dtphcportentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtphcportentry.EntityData.ParentYangName = "dot1dTpHCPortTable"
    dot1Dtphcportentry.EntityData.SegmentPath = "dot1dTpHCPortEntry" + "[dot1dTpPort='" + fmt.Sprintf("%v", dot1Dtphcportentry.Dot1Dtpport) + "']"
    dot1Dtphcportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtphcportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtphcportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtphcportentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dtphcportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dtphcportentry.EntityData.Leafs["dot1dTpPort"] = types.YLeaf{"Dot1Dtpport", dot1Dtphcportentry.Dot1Dtpport}
    dot1Dtphcportentry.EntityData.Leafs["dot1dTpHCPortInFrames"] = types.YLeaf{"Dot1Dtphcportinframes", dot1Dtphcportentry.Dot1Dtphcportinframes}
    dot1Dtphcportentry.EntityData.Leafs["dot1dTpHCPortOutFrames"] = types.YLeaf{"Dot1Dtphcportoutframes", dot1Dtphcportentry.Dot1Dtphcportoutframes}
    dot1Dtphcportentry.EntityData.Leafs["dot1dTpHCPortInDiscards"] = types.YLeaf{"Dot1Dtphcportindiscards", dot1Dtphcportentry.Dot1Dtphcportindiscards}
    return &(dot1Dtphcportentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The most significant bits of statistics counters for a high- capacity
    // interface of a transparent bridge.  Each object is associated with a
    // corresponding object in dot1dTpPortTable that indicates the least
    // significant bits of the counter. The type is slice of
    // PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry.
    Dot1Dtpportoverflowentry []PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry
}

func (dot1Dtpportoverflowtable *PBRIDGEMIB_Dot1Dtpportoverflowtable) GetEntityData() *types.CommonEntityData {
    dot1Dtpportoverflowtable.EntityData.YFilter = dot1Dtpportoverflowtable.YFilter
    dot1Dtpportoverflowtable.EntityData.YangName = "dot1dTpPortOverflowTable"
    dot1Dtpportoverflowtable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtpportoverflowtable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1Dtpportoverflowtable.EntityData.SegmentPath = "dot1dTpPortOverflowTable"
    dot1Dtpportoverflowtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtpportoverflowtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtpportoverflowtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtpportoverflowtable.EntityData.Children = make(map[string]types.YChild)
    dot1Dtpportoverflowtable.EntityData.Children["dot1dTpPortOverflowEntry"] = types.YChild{"Dot1Dtpportoverflowentry", nil}
    for i := range dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry {
        dot1Dtpportoverflowtable.EntityData.Children[types.GetSegmentPath(&dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry[i])] = types.YChild{"Dot1Dtpportoverflowentry", &dot1Dtpportoverflowtable.Dot1Dtpportoverflowentry[i]}
    }
    dot1Dtpportoverflowtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dtpportoverflowtable.EntityData)
}

// PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry
// The most significant bits of statistics counters for a high-
// capacity interface of a transparent bridge.  Each object is
// associated with a corresponding object in dot1dTpPortTable
// that indicates the least significant bits of the counter.
type PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry struct {
    EntityData types.CommonEntityData
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

func (dot1Dtpportoverflowentry *PBRIDGEMIB_Dot1Dtpportoverflowtable_Dot1Dtpportoverflowentry) GetEntityData() *types.CommonEntityData {
    dot1Dtpportoverflowentry.EntityData.YFilter = dot1Dtpportoverflowentry.YFilter
    dot1Dtpportoverflowentry.EntityData.YangName = "dot1dTpPortOverflowEntry"
    dot1Dtpportoverflowentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtpportoverflowentry.EntityData.ParentYangName = "dot1dTpPortOverflowTable"
    dot1Dtpportoverflowentry.EntityData.SegmentPath = "dot1dTpPortOverflowEntry" + "[dot1dTpPort='" + fmt.Sprintf("%v", dot1Dtpportoverflowentry.Dot1Dtpport) + "']"
    dot1Dtpportoverflowentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtpportoverflowentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtpportoverflowentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtpportoverflowentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dtpportoverflowentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dtpportoverflowentry.EntityData.Leafs["dot1dTpPort"] = types.YLeaf{"Dot1Dtpport", dot1Dtpportoverflowentry.Dot1Dtpport}
    dot1Dtpportoverflowentry.EntityData.Leafs["dot1dTpPortInOverflowFrames"] = types.YLeaf{"Dot1Dtpportinoverflowframes", dot1Dtpportoverflowentry.Dot1Dtpportinoverflowframes}
    dot1Dtpportoverflowentry.EntityData.Leafs["dot1dTpPortOutOverflowFrames"] = types.YLeaf{"Dot1Dtpportoutoverflowframes", dot1Dtpportoverflowentry.Dot1Dtpportoutoverflowframes}
    dot1Dtpportoverflowentry.EntityData.Leafs["dot1dTpPortInOverflowDiscards"] = types.YLeaf{"Dot1Dtpportinoverflowdiscards", dot1Dtpportoverflowentry.Dot1Dtpportinoverflowdiscards}
    return &(dot1Dtpportoverflowentry.EntityData)
}

// PBRIDGEMIB_Dot1Duserpriorityregentable
// A list of Regenerated User Priorities for each received
// User Priority on each port of a bridge.  The Regenerated
// User Priority value may be used to index the Traffic
// Class Table for each input port.  This only has effect
// on media that support native User Priority.  The default
// values for Regenerated User Priorities are the same as
// the User Priorities.
type PBRIDGEMIB_Dot1Duserpriorityregentable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A mapping of incoming User Priority to a Regenerated User Priority. The
    // type is slice of
    // PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry.
    Dot1Duserpriorityregenentry []PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry
}

func (dot1Duserpriorityregentable *PBRIDGEMIB_Dot1Duserpriorityregentable) GetEntityData() *types.CommonEntityData {
    dot1Duserpriorityregentable.EntityData.YFilter = dot1Duserpriorityregentable.YFilter
    dot1Duserpriorityregentable.EntityData.YangName = "dot1dUserPriorityRegenTable"
    dot1Duserpriorityregentable.EntityData.BundleName = "cisco_ios_xe"
    dot1Duserpriorityregentable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1Duserpriorityregentable.EntityData.SegmentPath = "dot1dUserPriorityRegenTable"
    dot1Duserpriorityregentable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Duserpriorityregentable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Duserpriorityregentable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Duserpriorityregentable.EntityData.Children = make(map[string]types.YChild)
    dot1Duserpriorityregentable.EntityData.Children["dot1dUserPriorityRegenEntry"] = types.YChild{"Dot1Duserpriorityregenentry", nil}
    for i := range dot1Duserpriorityregentable.Dot1Duserpriorityregenentry {
        dot1Duserpriorityregentable.EntityData.Children[types.GetSegmentPath(&dot1Duserpriorityregentable.Dot1Duserpriorityregenentry[i])] = types.YChild{"Dot1Duserpriorityregenentry", &dot1Duserpriorityregentable.Dot1Duserpriorityregenentry[i]}
    }
    dot1Duserpriorityregentable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Duserpriorityregentable.EntityData)
}

// PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry
// A mapping of incoming User Priority to a Regenerated
// User Priority.
type PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry struct {
    EntityData types.CommonEntityData
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

func (dot1Duserpriorityregenentry *PBRIDGEMIB_Dot1Duserpriorityregentable_Dot1Duserpriorityregenentry) GetEntityData() *types.CommonEntityData {
    dot1Duserpriorityregenentry.EntityData.YFilter = dot1Duserpriorityregenentry.YFilter
    dot1Duserpriorityregenentry.EntityData.YangName = "dot1dUserPriorityRegenEntry"
    dot1Duserpriorityregenentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Duserpriorityregenentry.EntityData.ParentYangName = "dot1dUserPriorityRegenTable"
    dot1Duserpriorityregenentry.EntityData.SegmentPath = "dot1dUserPriorityRegenEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Duserpriorityregenentry.Dot1Dbaseport) + "']" + "[dot1dUserPriority='" + fmt.Sprintf("%v", dot1Duserpriorityregenentry.Dot1Duserpriority) + "']"
    dot1Duserpriorityregenentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Duserpriorityregenentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Duserpriorityregenentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Duserpriorityregenentry.EntityData.Children = make(map[string]types.YChild)
    dot1Duserpriorityregenentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Duserpriorityregenentry.EntityData.Leafs["dot1dBasePort"] = types.YLeaf{"Dot1Dbaseport", dot1Duserpriorityregenentry.Dot1Dbaseport}
    dot1Duserpriorityregenentry.EntityData.Leafs["dot1dUserPriority"] = types.YLeaf{"Dot1Duserpriority", dot1Duserpriorityregenentry.Dot1Duserpriority}
    dot1Duserpriorityregenentry.EntityData.Leafs["dot1dRegenUserPriority"] = types.YLeaf{"Dot1Dregenuserpriority", dot1Duserpriorityregenentry.Dot1Dregenuserpriority}
    return &(dot1Duserpriorityregenentry.EntityData)
}

// PBRIDGEMIB_Dot1Dtrafficclasstable
// A table mapping evaluated User Priority to Traffic
// Class, for forwarding by the bridge.  Traffic class is a
// number in the range (0..(dot1dPortNumTrafficClasses-1)).
type PBRIDGEMIB_Dot1Dtrafficclasstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User Priority to Traffic Class mapping. The type is slice of
    // PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry.
    Dot1Dtrafficclassentry []PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry
}

func (dot1Dtrafficclasstable *PBRIDGEMIB_Dot1Dtrafficclasstable) GetEntityData() *types.CommonEntityData {
    dot1Dtrafficclasstable.EntityData.YFilter = dot1Dtrafficclasstable.YFilter
    dot1Dtrafficclasstable.EntityData.YangName = "dot1dTrafficClassTable"
    dot1Dtrafficclasstable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtrafficclasstable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1Dtrafficclasstable.EntityData.SegmentPath = "dot1dTrafficClassTable"
    dot1Dtrafficclasstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtrafficclasstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtrafficclasstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtrafficclasstable.EntityData.Children = make(map[string]types.YChild)
    dot1Dtrafficclasstable.EntityData.Children["dot1dTrafficClassEntry"] = types.YChild{"Dot1Dtrafficclassentry", nil}
    for i := range dot1Dtrafficclasstable.Dot1Dtrafficclassentry {
        dot1Dtrafficclasstable.EntityData.Children[types.GetSegmentPath(&dot1Dtrafficclasstable.Dot1Dtrafficclassentry[i])] = types.YChild{"Dot1Dtrafficclassentry", &dot1Dtrafficclasstable.Dot1Dtrafficclassentry[i]}
    }
    dot1Dtrafficclasstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dtrafficclasstable.EntityData)
}

// PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry
// User Priority to Traffic Class mapping.
type PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry struct {
    EntityData types.CommonEntityData
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

func (dot1Dtrafficclassentry *PBRIDGEMIB_Dot1Dtrafficclasstable_Dot1Dtrafficclassentry) GetEntityData() *types.CommonEntityData {
    dot1Dtrafficclassentry.EntityData.YFilter = dot1Dtrafficclassentry.YFilter
    dot1Dtrafficclassentry.EntityData.YangName = "dot1dTrafficClassEntry"
    dot1Dtrafficclassentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dtrafficclassentry.EntityData.ParentYangName = "dot1dTrafficClassTable"
    dot1Dtrafficclassentry.EntityData.SegmentPath = "dot1dTrafficClassEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Dtrafficclassentry.Dot1Dbaseport) + "']" + "[dot1dTrafficClassPriority='" + fmt.Sprintf("%v", dot1Dtrafficclassentry.Dot1Dtrafficclasspriority) + "']"
    dot1Dtrafficclassentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dtrafficclassentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dtrafficclassentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dtrafficclassentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dtrafficclassentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dtrafficclassentry.EntityData.Leafs["dot1dBasePort"] = types.YLeaf{"Dot1Dbaseport", dot1Dtrafficclassentry.Dot1Dbaseport}
    dot1Dtrafficclassentry.EntityData.Leafs["dot1dTrafficClassPriority"] = types.YLeaf{"Dot1Dtrafficclasspriority", dot1Dtrafficclassentry.Dot1Dtrafficclasspriority}
    dot1Dtrafficclassentry.EntityData.Leafs["dot1dTrafficClass"] = types.YLeaf{"Dot1Dtrafficclass", dot1Dtrafficclassentry.Dot1Dtrafficclass}
    return &(dot1Dtrafficclassentry.EntityData)
}

// PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable
// A table mapping Regenerated User Priority to Outbound
// Access Priority.  This is a fixed mapping for all port
// types, with two options for 802.5 Token Ring.
type PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Regenerated User Priority to Outbound Access Priority mapping. The type is
    // slice of
    // PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry.
    Dot1Dportoutboundaccesspriorityentry []PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry
}

func (dot1Dportoutboundaccessprioritytable *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable) GetEntityData() *types.CommonEntityData {
    dot1Dportoutboundaccessprioritytable.EntityData.YFilter = dot1Dportoutboundaccessprioritytable.YFilter
    dot1Dportoutboundaccessprioritytable.EntityData.YangName = "dot1dPortOutboundAccessPriorityTable"
    dot1Dportoutboundaccessprioritytable.EntityData.BundleName = "cisco_ios_xe"
    dot1Dportoutboundaccessprioritytable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1Dportoutboundaccessprioritytable.EntityData.SegmentPath = "dot1dPortOutboundAccessPriorityTable"
    dot1Dportoutboundaccessprioritytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dportoutboundaccessprioritytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dportoutboundaccessprioritytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dportoutboundaccessprioritytable.EntityData.Children = make(map[string]types.YChild)
    dot1Dportoutboundaccessprioritytable.EntityData.Children["dot1dPortOutboundAccessPriorityEntry"] = types.YChild{"Dot1Dportoutboundaccesspriorityentry", nil}
    for i := range dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry {
        dot1Dportoutboundaccessprioritytable.EntityData.Children[types.GetSegmentPath(&dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry[i])] = types.YChild{"Dot1Dportoutboundaccesspriorityentry", &dot1Dportoutboundaccessprioritytable.Dot1Dportoutboundaccesspriorityentry[i]}
    }
    dot1Dportoutboundaccessprioritytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Dportoutboundaccessprioritytable.EntityData)
}

// PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry
// Regenerated User Priority to Outbound Access Priority
// mapping.
type PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry struct {
    EntityData types.CommonEntityData
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

func (dot1Dportoutboundaccesspriorityentry *PBRIDGEMIB_Dot1Dportoutboundaccessprioritytable_Dot1Dportoutboundaccesspriorityentry) GetEntityData() *types.CommonEntityData {
    dot1Dportoutboundaccesspriorityentry.EntityData.YFilter = dot1Dportoutboundaccesspriorityentry.YFilter
    dot1Dportoutboundaccesspriorityentry.EntityData.YangName = "dot1dPortOutboundAccessPriorityEntry"
    dot1Dportoutboundaccesspriorityentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Dportoutboundaccesspriorityentry.EntityData.ParentYangName = "dot1dPortOutboundAccessPriorityTable"
    dot1Dportoutboundaccesspriorityentry.EntityData.SegmentPath = "dot1dPortOutboundAccessPriorityEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Dportoutboundaccesspriorityentry.Dot1Dbaseport) + "']" + "[dot1dRegenUserPriority='" + fmt.Sprintf("%v", dot1Dportoutboundaccesspriorityentry.Dot1Dregenuserpriority) + "']"
    dot1Dportoutboundaccesspriorityentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Dportoutboundaccesspriorityentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Dportoutboundaccesspriorityentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Dportoutboundaccesspriorityentry.EntityData.Children = make(map[string]types.YChild)
    dot1Dportoutboundaccesspriorityentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Dportoutboundaccesspriorityentry.EntityData.Leafs["dot1dBasePort"] = types.YLeaf{"Dot1Dbaseport", dot1Dportoutboundaccesspriorityentry.Dot1Dbaseport}
    dot1Dportoutboundaccesspriorityentry.EntityData.Leafs["dot1dRegenUserPriority"] = types.YLeaf{"Dot1Dregenuserpriority", dot1Dportoutboundaccesspriorityentry.Dot1Dregenuserpriority}
    dot1Dportoutboundaccesspriorityentry.EntityData.Leafs["dot1dPortOutboundAccessPriority"] = types.YLeaf{"Dot1Dportoutboundaccesspriority", dot1Dportoutboundaccesspriorityentry.Dot1Dportoutboundaccesspriority}
    return &(dot1Dportoutboundaccesspriorityentry.EntityData)
}

