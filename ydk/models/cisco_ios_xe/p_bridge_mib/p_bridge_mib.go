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

    
    Dot1dExtBase PBRIDGEMIB_Dot1dExtBase

    // A table that contains information about every high- capacity port that is
    // associated with this transparent bridge.
    Dot1dTpHCPortTable PBRIDGEMIB_Dot1dTpHCPortTable

    // A table that contains the most-significant bits of statistics counters for
    // ports that are associated with this transparent bridge that are on
    // high-capacity interfaces, as defined in the conformance clauses for this
    // table.  This table is provided as a way to read 64-bit counters for agents
    // that support only SNMPv1.  Note that the reporting of most-significant and
    // least-significant counter bits separately runs the risk of missing an
    // overflow of the lower bits in the interval between sampling.  The manager
    // must be aware of this possibility, even within the same varbindlist, when
    // interpreting the results of a request or asynchronous notification.
    Dot1dTpPortOverflowTable PBRIDGEMIB_Dot1dTpPortOverflowTable

    // A list of Regenerated User Priorities for each received User Priority on
    // each port of a bridge.  The Regenerated User Priority value may be used to
    // index the Traffic Class Table for each input port.  This only has effect on
    // media that support native User Priority.  The default values for
    // Regenerated User Priorities are the same as the User Priorities.
    Dot1dUserPriorityRegenTable PBRIDGEMIB_Dot1dUserPriorityRegenTable

    // A table mapping evaluated User Priority to Traffic Class, for forwarding by
    // the bridge.  Traffic class is a number in the range
    // (0..(dot1dPortNumTrafficClasses-1)).
    Dot1dTrafficClassTable PBRIDGEMIB_Dot1dTrafficClassTable

    // A table mapping Regenerated User Priority to Outbound Access Priority. 
    // This is a fixed mapping for all port types, with two options for 802.5
    // Token Ring.
    Dot1dPortOutboundAccessPriorityTable PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable
}

func (pBRIDGEMIB *PBRIDGEMIB) GetEntityData() *types.CommonEntityData {
    pBRIDGEMIB.EntityData.YFilter = pBRIDGEMIB.YFilter
    pBRIDGEMIB.EntityData.YangName = "P-BRIDGE-MIB"
    pBRIDGEMIB.EntityData.BundleName = "cisco_ios_xe"
    pBRIDGEMIB.EntityData.ParentYangName = "P-BRIDGE-MIB"
    pBRIDGEMIB.EntityData.SegmentPath = "P-BRIDGE-MIB:P-BRIDGE-MIB"
    pBRIDGEMIB.EntityData.AbsolutePath = pBRIDGEMIB.EntityData.SegmentPath
    pBRIDGEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pBRIDGEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pBRIDGEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pBRIDGEMIB.EntityData.Children = types.NewOrderedMap()
    pBRIDGEMIB.EntityData.Children.Append("dot1dExtBase", types.YChild{"Dot1dExtBase", &pBRIDGEMIB.Dot1dExtBase})
    pBRIDGEMIB.EntityData.Children.Append("dot1dTpHCPortTable", types.YChild{"Dot1dTpHCPortTable", &pBRIDGEMIB.Dot1dTpHCPortTable})
    pBRIDGEMIB.EntityData.Children.Append("dot1dTpPortOverflowTable", types.YChild{"Dot1dTpPortOverflowTable", &pBRIDGEMIB.Dot1dTpPortOverflowTable})
    pBRIDGEMIB.EntityData.Children.Append("dot1dUserPriorityRegenTable", types.YChild{"Dot1dUserPriorityRegenTable", &pBRIDGEMIB.Dot1dUserPriorityRegenTable})
    pBRIDGEMIB.EntityData.Children.Append("dot1dTrafficClassTable", types.YChild{"Dot1dTrafficClassTable", &pBRIDGEMIB.Dot1dTrafficClassTable})
    pBRIDGEMIB.EntityData.Children.Append("dot1dPortOutboundAccessPriorityTable", types.YChild{"Dot1dPortOutboundAccessPriorityTable", &pBRIDGEMIB.Dot1dPortOutboundAccessPriorityTable})
    pBRIDGEMIB.EntityData.Leafs = types.NewOrderedMap()

    pBRIDGEMIB.EntityData.YListKeys = []string {}

    return &(pBRIDGEMIB.EntityData)
}

// PBRIDGEMIB_Dot1dExtBase
type PBRIDGEMIB_Dot1dExtBase struct {
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
    Dot1dDeviceCapabilities interface{}

    // The value true(1) indicates that Traffic Classes are enabled on this
    // bridge.  When false(2), the bridge operates with a single priority level
    // for all traffic.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is bool.
    Dot1dTrafficClassesEnabled interface{}

    // The administrative status requested by management for GMRP.  The value
    // enabled(1) indicates that GMRP should be enabled on this device, in all
    // VLANs, on all ports for which it has not been specifically disabled.  When
    // disabled(2), GMRP is disabled, in all VLANs and on all ports, and all GMRP
    // packets will be forwarded transparently.  This object affects both
    // Applicant and Registrar state machines.  A transition from disabled(2) to
    // enabled(1) will cause a reset of all GMRP state machines on all ports.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is EnabledStatus.
    Dot1dGmrpStatus interface{}
}

func (dot1dExtBase *PBRIDGEMIB_Dot1dExtBase) GetEntityData() *types.CommonEntityData {
    dot1dExtBase.EntityData.YFilter = dot1dExtBase.YFilter
    dot1dExtBase.EntityData.YangName = "dot1dExtBase"
    dot1dExtBase.EntityData.BundleName = "cisco_ios_xe"
    dot1dExtBase.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1dExtBase.EntityData.SegmentPath = "dot1dExtBase"
    dot1dExtBase.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/" + dot1dExtBase.EntityData.SegmentPath
    dot1dExtBase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dExtBase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dExtBase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dExtBase.EntityData.Children = types.NewOrderedMap()
    dot1dExtBase.EntityData.Leafs = types.NewOrderedMap()
    dot1dExtBase.EntityData.Leafs.Append("dot1dDeviceCapabilities", types.YLeaf{"Dot1dDeviceCapabilities", dot1dExtBase.Dot1dDeviceCapabilities})
    dot1dExtBase.EntityData.Leafs.Append("dot1dTrafficClassesEnabled", types.YLeaf{"Dot1dTrafficClassesEnabled", dot1dExtBase.Dot1dTrafficClassesEnabled})
    dot1dExtBase.EntityData.Leafs.Append("dot1dGmrpStatus", types.YLeaf{"Dot1dGmrpStatus", dot1dExtBase.Dot1dGmrpStatus})

    dot1dExtBase.EntityData.YListKeys = []string {}

    return &(dot1dExtBase.EntityData)
}

// PBRIDGEMIB_Dot1dTpHCPortTable
// A table that contains information about every high-
// capacity port that is associated with this transparent
// bridge.
type PBRIDGEMIB_Dot1dTpHCPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics information for each high-capacity port of a transparent bridge.
    // The type is slice of PBRIDGEMIB_Dot1dTpHCPortTable_Dot1dTpHCPortEntry.
    Dot1dTpHCPortEntry []*PBRIDGEMIB_Dot1dTpHCPortTable_Dot1dTpHCPortEntry
}

func (dot1dTpHCPortTable *PBRIDGEMIB_Dot1dTpHCPortTable) GetEntityData() *types.CommonEntityData {
    dot1dTpHCPortTable.EntityData.YFilter = dot1dTpHCPortTable.YFilter
    dot1dTpHCPortTable.EntityData.YangName = "dot1dTpHCPortTable"
    dot1dTpHCPortTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpHCPortTable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1dTpHCPortTable.EntityData.SegmentPath = "dot1dTpHCPortTable"
    dot1dTpHCPortTable.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/" + dot1dTpHCPortTable.EntityData.SegmentPath
    dot1dTpHCPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpHCPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpHCPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpHCPortTable.EntityData.Children = types.NewOrderedMap()
    dot1dTpHCPortTable.EntityData.Children.Append("dot1dTpHCPortEntry", types.YChild{"Dot1dTpHCPortEntry", nil})
    for i := range dot1dTpHCPortTable.Dot1dTpHCPortEntry {
        dot1dTpHCPortTable.EntityData.Children.Append(types.GetSegmentPath(dot1dTpHCPortTable.Dot1dTpHCPortEntry[i]), types.YChild{"Dot1dTpHCPortEntry", dot1dTpHCPortTable.Dot1dTpHCPortEntry[i]})
    }
    dot1dTpHCPortTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dTpHCPortTable.EntityData.YListKeys = []string {}

    return &(dot1dTpHCPortTable.EntityData)
}

// PBRIDGEMIB_Dot1dTpHCPortTable_Dot1dTpHCPortEntry
// Statistics information for each high-capacity port of a
// transparent bridge.
type PBRIDGEMIB_Dot1dTpHCPortTable_Dot1dTpHCPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dTpPortTable_Dot1dTpPortEntry_Dot1dTpPort
    Dot1dTpPort interface{}

    // The number of frames that have been received by this port from its segment.
    // Note that a frame received on the interface corresponding to this port is
    // only counted by this object if and only if it is for a protocol being
    // processed by the local bridging function, including bridge management
    // frames. The type is interface{} with range: 0..18446744073709551615.
    Dot1dTpHCPortInFrames interface{}

    // The number of frames that have been transmitted by this port to its
    // segment.  Note that a frame transmitted on the interface corresponding to
    // this port is only counted by this object if and only if it is for a
    // protocol being processed by the local bridging function, including bridge
    // management frames. The type is interface{} with range:
    // 0..18446744073709551615.
    Dot1dTpHCPortOutFrames interface{}

    // Count of valid frames that have been received by this port from its segment
    // that were discarded (i.e., filtered) by the Forwarding Process. The type is
    // interface{} with range: 0..18446744073709551615.
    Dot1dTpHCPortInDiscards interface{}
}

func (dot1dTpHCPortEntry *PBRIDGEMIB_Dot1dTpHCPortTable_Dot1dTpHCPortEntry) GetEntityData() *types.CommonEntityData {
    dot1dTpHCPortEntry.EntityData.YFilter = dot1dTpHCPortEntry.YFilter
    dot1dTpHCPortEntry.EntityData.YangName = "dot1dTpHCPortEntry"
    dot1dTpHCPortEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpHCPortEntry.EntityData.ParentYangName = "dot1dTpHCPortTable"
    dot1dTpHCPortEntry.EntityData.SegmentPath = "dot1dTpHCPortEntry" + types.AddKeyToken(dot1dTpHCPortEntry.Dot1dTpPort, "dot1dTpPort")
    dot1dTpHCPortEntry.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/dot1dTpHCPortTable/" + dot1dTpHCPortEntry.EntityData.SegmentPath
    dot1dTpHCPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpHCPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpHCPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpHCPortEntry.EntityData.Children = types.NewOrderedMap()
    dot1dTpHCPortEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dTpHCPortEntry.EntityData.Leafs.Append("dot1dTpPort", types.YLeaf{"Dot1dTpPort", dot1dTpHCPortEntry.Dot1dTpPort})
    dot1dTpHCPortEntry.EntityData.Leafs.Append("dot1dTpHCPortInFrames", types.YLeaf{"Dot1dTpHCPortInFrames", dot1dTpHCPortEntry.Dot1dTpHCPortInFrames})
    dot1dTpHCPortEntry.EntityData.Leafs.Append("dot1dTpHCPortOutFrames", types.YLeaf{"Dot1dTpHCPortOutFrames", dot1dTpHCPortEntry.Dot1dTpHCPortOutFrames})
    dot1dTpHCPortEntry.EntityData.Leafs.Append("dot1dTpHCPortInDiscards", types.YLeaf{"Dot1dTpHCPortInDiscards", dot1dTpHCPortEntry.Dot1dTpHCPortInDiscards})

    dot1dTpHCPortEntry.EntityData.YListKeys = []string {"Dot1dTpPort"}

    return &(dot1dTpHCPortEntry.EntityData)
}

// PBRIDGEMIB_Dot1dTpPortOverflowTable
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
type PBRIDGEMIB_Dot1dTpPortOverflowTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The most significant bits of statistics counters for a high- capacity
    // interface of a transparent bridge.  Each object is associated with a
    // corresponding object in dot1dTpPortTable that indicates the least
    // significant bits of the counter. The type is slice of
    // PBRIDGEMIB_Dot1dTpPortOverflowTable_Dot1dTpPortOverflowEntry.
    Dot1dTpPortOverflowEntry []*PBRIDGEMIB_Dot1dTpPortOverflowTable_Dot1dTpPortOverflowEntry
}

func (dot1dTpPortOverflowTable *PBRIDGEMIB_Dot1dTpPortOverflowTable) GetEntityData() *types.CommonEntityData {
    dot1dTpPortOverflowTable.EntityData.YFilter = dot1dTpPortOverflowTable.YFilter
    dot1dTpPortOverflowTable.EntityData.YangName = "dot1dTpPortOverflowTable"
    dot1dTpPortOverflowTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpPortOverflowTable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1dTpPortOverflowTable.EntityData.SegmentPath = "dot1dTpPortOverflowTable"
    dot1dTpPortOverflowTable.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/" + dot1dTpPortOverflowTable.EntityData.SegmentPath
    dot1dTpPortOverflowTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpPortOverflowTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpPortOverflowTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpPortOverflowTable.EntityData.Children = types.NewOrderedMap()
    dot1dTpPortOverflowTable.EntityData.Children.Append("dot1dTpPortOverflowEntry", types.YChild{"Dot1dTpPortOverflowEntry", nil})
    for i := range dot1dTpPortOverflowTable.Dot1dTpPortOverflowEntry {
        dot1dTpPortOverflowTable.EntityData.Children.Append(types.GetSegmentPath(dot1dTpPortOverflowTable.Dot1dTpPortOverflowEntry[i]), types.YChild{"Dot1dTpPortOverflowEntry", dot1dTpPortOverflowTable.Dot1dTpPortOverflowEntry[i]})
    }
    dot1dTpPortOverflowTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dTpPortOverflowTable.EntityData.YListKeys = []string {}

    return &(dot1dTpPortOverflowTable.EntityData)
}

// PBRIDGEMIB_Dot1dTpPortOverflowTable_Dot1dTpPortOverflowEntry
// The most significant bits of statistics counters for a high-
// capacity interface of a transparent bridge.  Each object is
// associated with a corresponding object in dot1dTpPortTable
// that indicates the least significant bits of the counter.
type PBRIDGEMIB_Dot1dTpPortOverflowTable_Dot1dTpPortOverflowEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dTpPortTable_Dot1dTpPortEntry_Dot1dTpPort
    Dot1dTpPort interface{}

    // The number of times the associated dot1dTpPortInFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1dTpPortInOverflowFrames interface{}

    // The number of times the associated dot1dTpPortOutFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1dTpPortOutOverflowFrames interface{}

    // The number of times the associated dot1dTpPortInDiscards counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1dTpPortInOverflowDiscards interface{}
}

func (dot1dTpPortOverflowEntry *PBRIDGEMIB_Dot1dTpPortOverflowTable_Dot1dTpPortOverflowEntry) GetEntityData() *types.CommonEntityData {
    dot1dTpPortOverflowEntry.EntityData.YFilter = dot1dTpPortOverflowEntry.YFilter
    dot1dTpPortOverflowEntry.EntityData.YangName = "dot1dTpPortOverflowEntry"
    dot1dTpPortOverflowEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dTpPortOverflowEntry.EntityData.ParentYangName = "dot1dTpPortOverflowTable"
    dot1dTpPortOverflowEntry.EntityData.SegmentPath = "dot1dTpPortOverflowEntry" + types.AddKeyToken(dot1dTpPortOverflowEntry.Dot1dTpPort, "dot1dTpPort")
    dot1dTpPortOverflowEntry.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/dot1dTpPortOverflowTable/" + dot1dTpPortOverflowEntry.EntityData.SegmentPath
    dot1dTpPortOverflowEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTpPortOverflowEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTpPortOverflowEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTpPortOverflowEntry.EntityData.Children = types.NewOrderedMap()
    dot1dTpPortOverflowEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dTpPortOverflowEntry.EntityData.Leafs.Append("dot1dTpPort", types.YLeaf{"Dot1dTpPort", dot1dTpPortOverflowEntry.Dot1dTpPort})
    dot1dTpPortOverflowEntry.EntityData.Leafs.Append("dot1dTpPortInOverflowFrames", types.YLeaf{"Dot1dTpPortInOverflowFrames", dot1dTpPortOverflowEntry.Dot1dTpPortInOverflowFrames})
    dot1dTpPortOverflowEntry.EntityData.Leafs.Append("dot1dTpPortOutOverflowFrames", types.YLeaf{"Dot1dTpPortOutOverflowFrames", dot1dTpPortOverflowEntry.Dot1dTpPortOutOverflowFrames})
    dot1dTpPortOverflowEntry.EntityData.Leafs.Append("dot1dTpPortInOverflowDiscards", types.YLeaf{"Dot1dTpPortInOverflowDiscards", dot1dTpPortOverflowEntry.Dot1dTpPortInOverflowDiscards})

    dot1dTpPortOverflowEntry.EntityData.YListKeys = []string {"Dot1dTpPort"}

    return &(dot1dTpPortOverflowEntry.EntityData)
}

// PBRIDGEMIB_Dot1dUserPriorityRegenTable
// A list of Regenerated User Priorities for each received
// User Priority on each port of a bridge.  The Regenerated
// User Priority value may be used to index the Traffic
// Class Table for each input port.  This only has effect
// on media that support native User Priority.  The default
// values for Regenerated User Priorities are the same as
// the User Priorities.
type PBRIDGEMIB_Dot1dUserPriorityRegenTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A mapping of incoming User Priority to a Regenerated User Priority. The
    // type is slice of
    // PBRIDGEMIB_Dot1dUserPriorityRegenTable_Dot1dUserPriorityRegenEntry.
    Dot1dUserPriorityRegenEntry []*PBRIDGEMIB_Dot1dUserPriorityRegenTable_Dot1dUserPriorityRegenEntry
}

func (dot1dUserPriorityRegenTable *PBRIDGEMIB_Dot1dUserPriorityRegenTable) GetEntityData() *types.CommonEntityData {
    dot1dUserPriorityRegenTable.EntityData.YFilter = dot1dUserPriorityRegenTable.YFilter
    dot1dUserPriorityRegenTable.EntityData.YangName = "dot1dUserPriorityRegenTable"
    dot1dUserPriorityRegenTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dUserPriorityRegenTable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1dUserPriorityRegenTable.EntityData.SegmentPath = "dot1dUserPriorityRegenTable"
    dot1dUserPriorityRegenTable.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/" + dot1dUserPriorityRegenTable.EntityData.SegmentPath
    dot1dUserPriorityRegenTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dUserPriorityRegenTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dUserPriorityRegenTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dUserPriorityRegenTable.EntityData.Children = types.NewOrderedMap()
    dot1dUserPriorityRegenTable.EntityData.Children.Append("dot1dUserPriorityRegenEntry", types.YChild{"Dot1dUserPriorityRegenEntry", nil})
    for i := range dot1dUserPriorityRegenTable.Dot1dUserPriorityRegenEntry {
        dot1dUserPriorityRegenTable.EntityData.Children.Append(types.GetSegmentPath(dot1dUserPriorityRegenTable.Dot1dUserPriorityRegenEntry[i]), types.YChild{"Dot1dUserPriorityRegenEntry", dot1dUserPriorityRegenTable.Dot1dUserPriorityRegenEntry[i]})
    }
    dot1dUserPriorityRegenTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dUserPriorityRegenTable.EntityData.YListKeys = []string {}

    return &(dot1dUserPriorityRegenTable.EntityData)
}

// PBRIDGEMIB_Dot1dUserPriorityRegenTable_Dot1dUserPriorityRegenEntry
// A mapping of incoming User Priority to a Regenerated
// User Priority.
type PBRIDGEMIB_Dot1dUserPriorityRegenTable_Dot1dUserPriorityRegenEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1dBasePort
    Dot1dBasePort interface{}

    // This attribute is a key. The User Priority for a frame received on this
    // port. The type is interface{} with range: 0..7.
    Dot1dUserPriority interface{}

    // The Regenerated User Priority that the incoming User  Priority is mapped to
    // for this port.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is interface{} with
    // range: 0..7.
    Dot1dRegenUserPriority interface{}
}

func (dot1dUserPriorityRegenEntry *PBRIDGEMIB_Dot1dUserPriorityRegenTable_Dot1dUserPriorityRegenEntry) GetEntityData() *types.CommonEntityData {
    dot1dUserPriorityRegenEntry.EntityData.YFilter = dot1dUserPriorityRegenEntry.YFilter
    dot1dUserPriorityRegenEntry.EntityData.YangName = "dot1dUserPriorityRegenEntry"
    dot1dUserPriorityRegenEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dUserPriorityRegenEntry.EntityData.ParentYangName = "dot1dUserPriorityRegenTable"
    dot1dUserPriorityRegenEntry.EntityData.SegmentPath = "dot1dUserPriorityRegenEntry" + types.AddKeyToken(dot1dUserPriorityRegenEntry.Dot1dBasePort, "dot1dBasePort") + types.AddKeyToken(dot1dUserPriorityRegenEntry.Dot1dUserPriority, "dot1dUserPriority")
    dot1dUserPriorityRegenEntry.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/dot1dUserPriorityRegenTable/" + dot1dUserPriorityRegenEntry.EntityData.SegmentPath
    dot1dUserPriorityRegenEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dUserPriorityRegenEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dUserPriorityRegenEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dUserPriorityRegenEntry.EntityData.Children = types.NewOrderedMap()
    dot1dUserPriorityRegenEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dUserPriorityRegenEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", dot1dUserPriorityRegenEntry.Dot1dBasePort})
    dot1dUserPriorityRegenEntry.EntityData.Leafs.Append("dot1dUserPriority", types.YLeaf{"Dot1dUserPriority", dot1dUserPriorityRegenEntry.Dot1dUserPriority})
    dot1dUserPriorityRegenEntry.EntityData.Leafs.Append("dot1dRegenUserPriority", types.YLeaf{"Dot1dRegenUserPriority", dot1dUserPriorityRegenEntry.Dot1dRegenUserPriority})

    dot1dUserPriorityRegenEntry.EntityData.YListKeys = []string {"Dot1dBasePort", "Dot1dUserPriority"}

    return &(dot1dUserPriorityRegenEntry.EntityData)
}

// PBRIDGEMIB_Dot1dTrafficClassTable
// A table mapping evaluated User Priority to Traffic
// Class, for forwarding by the bridge.  Traffic class is a
// number in the range (0..(dot1dPortNumTrafficClasses-1)).
type PBRIDGEMIB_Dot1dTrafficClassTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User Priority to Traffic Class mapping. The type is slice of
    // PBRIDGEMIB_Dot1dTrafficClassTable_Dot1dTrafficClassEntry.
    Dot1dTrafficClassEntry []*PBRIDGEMIB_Dot1dTrafficClassTable_Dot1dTrafficClassEntry
}

func (dot1dTrafficClassTable *PBRIDGEMIB_Dot1dTrafficClassTable) GetEntityData() *types.CommonEntityData {
    dot1dTrafficClassTable.EntityData.YFilter = dot1dTrafficClassTable.YFilter
    dot1dTrafficClassTable.EntityData.YangName = "dot1dTrafficClassTable"
    dot1dTrafficClassTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dTrafficClassTable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1dTrafficClassTable.EntityData.SegmentPath = "dot1dTrafficClassTable"
    dot1dTrafficClassTable.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/" + dot1dTrafficClassTable.EntityData.SegmentPath
    dot1dTrafficClassTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTrafficClassTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTrafficClassTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTrafficClassTable.EntityData.Children = types.NewOrderedMap()
    dot1dTrafficClassTable.EntityData.Children.Append("dot1dTrafficClassEntry", types.YChild{"Dot1dTrafficClassEntry", nil})
    for i := range dot1dTrafficClassTable.Dot1dTrafficClassEntry {
        dot1dTrafficClassTable.EntityData.Children.Append(types.GetSegmentPath(dot1dTrafficClassTable.Dot1dTrafficClassEntry[i]), types.YChild{"Dot1dTrafficClassEntry", dot1dTrafficClassTable.Dot1dTrafficClassEntry[i]})
    }
    dot1dTrafficClassTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dTrafficClassTable.EntityData.YListKeys = []string {}

    return &(dot1dTrafficClassTable.EntityData)
}

// PBRIDGEMIB_Dot1dTrafficClassTable_Dot1dTrafficClassEntry
// User Priority to Traffic Class mapping.
type PBRIDGEMIB_Dot1dTrafficClassTable_Dot1dTrafficClassEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1dBasePort
    Dot1dBasePort interface{}

    // This attribute is a key. The Priority value determined for the received
    // frame. This value is equivalent to the priority indicated in the tagged
    // frame received, or one of the evaluated priorities, determined according to
    // the media-type.  For untagged frames received from Ethernet media, this
    // value is equal to the dot1dPortDefaultUserPriority value for the ingress
    // port.  For untagged frames received from non-Ethernet media, this value is
    // equal to the dot1dRegenUserPriority value for the ingress port and
    // media-specific user priority. The type is interface{} with range: 0..7.
    Dot1dTrafficClassPriority interface{}

    // The Traffic Class the received frame is mapped to.  The value of this
    // object MUST be retained across reinitializations of the management system.
    // The type is interface{} with range: 0..7.
    Dot1dTrafficClass interface{}
}

func (dot1dTrafficClassEntry *PBRIDGEMIB_Dot1dTrafficClassTable_Dot1dTrafficClassEntry) GetEntityData() *types.CommonEntityData {
    dot1dTrafficClassEntry.EntityData.YFilter = dot1dTrafficClassEntry.YFilter
    dot1dTrafficClassEntry.EntityData.YangName = "dot1dTrafficClassEntry"
    dot1dTrafficClassEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dTrafficClassEntry.EntityData.ParentYangName = "dot1dTrafficClassTable"
    dot1dTrafficClassEntry.EntityData.SegmentPath = "dot1dTrafficClassEntry" + types.AddKeyToken(dot1dTrafficClassEntry.Dot1dBasePort, "dot1dBasePort") + types.AddKeyToken(dot1dTrafficClassEntry.Dot1dTrafficClassPriority, "dot1dTrafficClassPriority")
    dot1dTrafficClassEntry.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/dot1dTrafficClassTable/" + dot1dTrafficClassEntry.EntityData.SegmentPath
    dot1dTrafficClassEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dTrafficClassEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dTrafficClassEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dTrafficClassEntry.EntityData.Children = types.NewOrderedMap()
    dot1dTrafficClassEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dTrafficClassEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", dot1dTrafficClassEntry.Dot1dBasePort})
    dot1dTrafficClassEntry.EntityData.Leafs.Append("dot1dTrafficClassPriority", types.YLeaf{"Dot1dTrafficClassPriority", dot1dTrafficClassEntry.Dot1dTrafficClassPriority})
    dot1dTrafficClassEntry.EntityData.Leafs.Append("dot1dTrafficClass", types.YLeaf{"Dot1dTrafficClass", dot1dTrafficClassEntry.Dot1dTrafficClass})

    dot1dTrafficClassEntry.EntityData.YListKeys = []string {"Dot1dBasePort", "Dot1dTrafficClassPriority"}

    return &(dot1dTrafficClassEntry.EntityData)
}

// PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable
// A table mapping Regenerated User Priority to Outbound
// Access Priority.  This is a fixed mapping for all port
// types, with two options for 802.5 Token Ring.
type PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Regenerated User Priority to Outbound Access Priority mapping. The type is
    // slice of
    // PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable_Dot1dPortOutboundAccessPriorityEntry.
    Dot1dPortOutboundAccessPriorityEntry []*PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable_Dot1dPortOutboundAccessPriorityEntry
}

func (dot1dPortOutboundAccessPriorityTable *PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable) GetEntityData() *types.CommonEntityData {
    dot1dPortOutboundAccessPriorityTable.EntityData.YFilter = dot1dPortOutboundAccessPriorityTable.YFilter
    dot1dPortOutboundAccessPriorityTable.EntityData.YangName = "dot1dPortOutboundAccessPriorityTable"
    dot1dPortOutboundAccessPriorityTable.EntityData.BundleName = "cisco_ios_xe"
    dot1dPortOutboundAccessPriorityTable.EntityData.ParentYangName = "P-BRIDGE-MIB"
    dot1dPortOutboundAccessPriorityTable.EntityData.SegmentPath = "dot1dPortOutboundAccessPriorityTable"
    dot1dPortOutboundAccessPriorityTable.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/" + dot1dPortOutboundAccessPriorityTable.EntityData.SegmentPath
    dot1dPortOutboundAccessPriorityTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dPortOutboundAccessPriorityTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dPortOutboundAccessPriorityTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dPortOutboundAccessPriorityTable.EntityData.Children = types.NewOrderedMap()
    dot1dPortOutboundAccessPriorityTable.EntityData.Children.Append("dot1dPortOutboundAccessPriorityEntry", types.YChild{"Dot1dPortOutboundAccessPriorityEntry", nil})
    for i := range dot1dPortOutboundAccessPriorityTable.Dot1dPortOutboundAccessPriorityEntry {
        dot1dPortOutboundAccessPriorityTable.EntityData.Children.Append(types.GetSegmentPath(dot1dPortOutboundAccessPriorityTable.Dot1dPortOutboundAccessPriorityEntry[i]), types.YChild{"Dot1dPortOutboundAccessPriorityEntry", dot1dPortOutboundAccessPriorityTable.Dot1dPortOutboundAccessPriorityEntry[i]})
    }
    dot1dPortOutboundAccessPriorityTable.EntityData.Leafs = types.NewOrderedMap()

    dot1dPortOutboundAccessPriorityTable.EntityData.YListKeys = []string {}

    return &(dot1dPortOutboundAccessPriorityTable.EntityData)
}

// PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable_Dot1dPortOutboundAccessPriorityEntry
// Regenerated User Priority to Outbound Access Priority
// mapping.
type PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable_Dot1dPortOutboundAccessPriorityEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1dBasePort
    Dot1dBasePort interface{}

    // This attribute is a key. The type is string with range: 0..7. Refers to
    // p_bridge_mib.PBRIDGEMIB_Dot1dUserPriorityRegenTable_Dot1dUserPriorityRegenEntry_Dot1dRegenUserPriority
    Dot1dRegenUserPriority interface{}

    // The Outbound Access Priority the received frame is mapped to. The type is
    // interface{} with range: 0..7.
    Dot1dPortOutboundAccessPriority interface{}
}

func (dot1dPortOutboundAccessPriorityEntry *PBRIDGEMIB_Dot1dPortOutboundAccessPriorityTable_Dot1dPortOutboundAccessPriorityEntry) GetEntityData() *types.CommonEntityData {
    dot1dPortOutboundAccessPriorityEntry.EntityData.YFilter = dot1dPortOutboundAccessPriorityEntry.YFilter
    dot1dPortOutboundAccessPriorityEntry.EntityData.YangName = "dot1dPortOutboundAccessPriorityEntry"
    dot1dPortOutboundAccessPriorityEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1dPortOutboundAccessPriorityEntry.EntityData.ParentYangName = "dot1dPortOutboundAccessPriorityTable"
    dot1dPortOutboundAccessPriorityEntry.EntityData.SegmentPath = "dot1dPortOutboundAccessPriorityEntry" + types.AddKeyToken(dot1dPortOutboundAccessPriorityEntry.Dot1dBasePort, "dot1dBasePort") + types.AddKeyToken(dot1dPortOutboundAccessPriorityEntry.Dot1dRegenUserPriority, "dot1dRegenUserPriority")
    dot1dPortOutboundAccessPriorityEntry.EntityData.AbsolutePath = "P-BRIDGE-MIB:P-BRIDGE-MIB/dot1dPortOutboundAccessPriorityTable/" + dot1dPortOutboundAccessPriorityEntry.EntityData.SegmentPath
    dot1dPortOutboundAccessPriorityEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1dPortOutboundAccessPriorityEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1dPortOutboundAccessPriorityEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1dPortOutboundAccessPriorityEntry.EntityData.Children = types.NewOrderedMap()
    dot1dPortOutboundAccessPriorityEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1dPortOutboundAccessPriorityEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", dot1dPortOutboundAccessPriorityEntry.Dot1dBasePort})
    dot1dPortOutboundAccessPriorityEntry.EntityData.Leafs.Append("dot1dRegenUserPriority", types.YLeaf{"Dot1dRegenUserPriority", dot1dPortOutboundAccessPriorityEntry.Dot1dRegenUserPriority})
    dot1dPortOutboundAccessPriorityEntry.EntityData.Leafs.Append("dot1dPortOutboundAccessPriority", types.YLeaf{"Dot1dPortOutboundAccessPriority", dot1dPortOutboundAccessPriorityEntry.Dot1dPortOutboundAccessPriority})

    dot1dPortOutboundAccessPriorityEntry.EntityData.YListKeys = []string {"Dot1dBasePort", "Dot1dRegenUserPriority"}

    return &(dot1dPortOutboundAccessPriorityEntry.EntityData)
}

