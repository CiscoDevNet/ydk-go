// The VLAN Bridge MIB module for managing Virtual Bridged
// Local Area Networks, as defined by IEEE 802.1Q-2003,
// including Restricted Vlan Registration defined by
// IEEE 802.1u-2001 and Vlan Classification defined by
// IEEE 802.1v-2001.
// 
// Copyright (C) The Internet Society (2006).  This version of
// this MIB module is part of RFC 4363; See the RFC itself for
// full legal notices.
package q_bridge_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package q_bridge_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:Q-BRIDGE-MIB Q-BRIDGE-MIB}", reflect.TypeOf(QBRIDGEMIB{}))
    ydk.RegisterEntity("Q-BRIDGE-MIB:Q-BRIDGE-MIB", reflect.TypeOf(QBRIDGEMIB{}))
}

// QBRIDGEMIB
type QBRIDGEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Dot1qBase QBRIDGEMIB_Dot1qBase

    
    Dot1qVlan QBRIDGEMIB_Dot1qVlan

    // A table that contains configuration and control information for each
    // Filtering Database currently operating on this device.  Entries in this
    // table appear automatically when VLANs are assigned FDB IDs in the
    // dot1qVlanCurrentTable.
    Dot1qFdbTable QBRIDGEMIB_Dot1qFdbTable

    // A table that contains information about unicast entries for which the
    // device has forwarding and/or filtering information.  This information is
    // used by the transparent bridging function in determining how to propagate a
    // received frame.
    Dot1qTpFdbTable QBRIDGEMIB_Dot1qTpFdbTable

    // A table containing filtering information for VLANs configured into the
    // bridge by (local or network) management, or learned dynamically, specifying
    // the set of ports to which frames received on a VLAN for this FDB and
    // containing a specific Group destination address are allowed to be
    // forwarded.
    Dot1qTpGroupTable QBRIDGEMIB_Dot1qTpGroupTable

    // A table containing forwarding information for each  VLAN, specifying the
    // set of ports to which forwarding of all multicasts applies, configured
    // statically by management or dynamically by GMRP.  An entry appears in this
    // table for all VLANs that are currently instantiated.
    Dot1qForwardAllTable QBRIDGEMIB_Dot1qForwardAllTable

    // A table containing forwarding information for each VLAN, specifying the set
    // of ports to which forwarding of multicast group-addressed frames for which
    // no more specific forwarding information applies.  This is configured
    // statically by management and determined dynamically by GMRP.  An entry
    // appears in this table for all VLANs that are currently instantiated.
    Dot1qForwardUnregisteredTable QBRIDGEMIB_Dot1qForwardUnregisteredTable

    // A table containing filtering information for Unicast MAC addresses for each
    // Filtering Database, configured into the device by (local or network)
    // management specifying the set of ports to which frames received from
    // specific ports and containing specific unicast destination addresses are
    // allowed to be forwarded.  A value of zero in this table (as the port number
    // from  which frames with a specific destination address are received) is
    // used to specify all ports for which there is no specific entry in this
    // table for that particular destination address.  Entries are valid for
    // unicast addresses only.
    Dot1qStaticUnicastTable QBRIDGEMIB_Dot1qStaticUnicastTable

    // A table containing filtering information for Multicast and Broadcast MAC
    // addresses for each VLAN, configured into the device by (local or network)
    // management specifying the set of ports to which frames received from
    // specific ports and containing specific Multicast and Broadcast destination
    // addresses are allowed to be forwarded.  A value of zero in this table (as
    // the port number from which frames with a specific destination address are
    // received) is used to specify all ports for which there is no specific entry
    // in this table for that particular destination address.  Entries are valid
    // for Multicast and Broadcast addresses only.
    Dot1qStaticMulticastTable QBRIDGEMIB_Dot1qStaticMulticastTable

    // A table containing current configuration information for each VLAN
    // currently configured into the device by (local or network) management, or
    // dynamically created as a result of GVRP requests received.
    Dot1qVlanCurrentTable QBRIDGEMIB_Dot1qVlanCurrentTable

    // A table containing static configuration information for each VLAN
    // configured into the device by (local or network) management.  All entries
    // are permanent and will be restored after the device is reset.
    Dot1qVlanStaticTable QBRIDGEMIB_Dot1qVlanStaticTable

    // A table containing per-port, per-VLAN statistics for traffic received. 
    // Separate objects are provided for both the most-significant and
    // least-significant bits of statistics counters for ports that are associated
    // with this transparent bridge.  The most-significant bit objects are only
    // required on high-capacity interfaces, as defined in the conformance clauses
    // for these objects.  This mechanism is provided as a way to read 64-bit
    // counters for agents that support only SNMPv1.  Note that the reporting of
    // most-significant and least- significant counter bits separately runs the
    // risk of missing an overflow of the lower bits in the interval between
    // sampling. The manager must be aware of this possibility, even within the
    // same varbindlist, when interpreting the results of a request or 
    // asynchronous notification.
    Dot1qPortVlanStatisticsTable QBRIDGEMIB_Dot1qPortVlanStatisticsTable

    // A table containing per-port, per-VLAN statistics for traffic on
    // high-capacity interfaces.
    Dot1qPortVlanHCStatisticsTable QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable

    // A table containing learning constraints for sets of Shared and Independent
    // VLANs.
    Dot1qLearningConstraintsTable QBRIDGEMIB_Dot1qLearningConstraintsTable

    // A table that contains mappings from Protocol Templates to Protocol Group
    // Identifiers used for Port-and-Protocol-based VLAN Classification.
    Dot1vProtocolGroupTable QBRIDGEMIB_Dot1vProtocolGroupTable

    // A table that contains VID sets used for Port-and-Protocol-based VLAN
    // Classification.
    Dot1vProtocolPortTable QBRIDGEMIB_Dot1vProtocolPortTable
}

func (qBRIDGEMIB *QBRIDGEMIB) GetEntityData() *types.CommonEntityData {
    qBRIDGEMIB.EntityData.YFilter = qBRIDGEMIB.YFilter
    qBRIDGEMIB.EntityData.YangName = "Q-BRIDGE-MIB"
    qBRIDGEMIB.EntityData.BundleName = "cisco_ios_xe"
    qBRIDGEMIB.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    qBRIDGEMIB.EntityData.SegmentPath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB"
    qBRIDGEMIB.EntityData.AbsolutePath = qBRIDGEMIB.EntityData.SegmentPath
    qBRIDGEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qBRIDGEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qBRIDGEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qBRIDGEMIB.EntityData.Children = types.NewOrderedMap()
    qBRIDGEMIB.EntityData.Children.Append("dot1qBase", types.YChild{"Dot1qBase", &qBRIDGEMIB.Dot1qBase})
    qBRIDGEMIB.EntityData.Children.Append("dot1qVlan", types.YChild{"Dot1qVlan", &qBRIDGEMIB.Dot1qVlan})
    qBRIDGEMIB.EntityData.Children.Append("dot1qFdbTable", types.YChild{"Dot1qFdbTable", &qBRIDGEMIB.Dot1qFdbTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qTpFdbTable", types.YChild{"Dot1qTpFdbTable", &qBRIDGEMIB.Dot1qTpFdbTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qTpGroupTable", types.YChild{"Dot1qTpGroupTable", &qBRIDGEMIB.Dot1qTpGroupTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qForwardAllTable", types.YChild{"Dot1qForwardAllTable", &qBRIDGEMIB.Dot1qForwardAllTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qForwardUnregisteredTable", types.YChild{"Dot1qForwardUnregisteredTable", &qBRIDGEMIB.Dot1qForwardUnregisteredTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qStaticUnicastTable", types.YChild{"Dot1qStaticUnicastTable", &qBRIDGEMIB.Dot1qStaticUnicastTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qStaticMulticastTable", types.YChild{"Dot1qStaticMulticastTable", &qBRIDGEMIB.Dot1qStaticMulticastTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qVlanCurrentTable", types.YChild{"Dot1qVlanCurrentTable", &qBRIDGEMIB.Dot1qVlanCurrentTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qVlanStaticTable", types.YChild{"Dot1qVlanStaticTable", &qBRIDGEMIB.Dot1qVlanStaticTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qPortVlanStatisticsTable", types.YChild{"Dot1qPortVlanStatisticsTable", &qBRIDGEMIB.Dot1qPortVlanStatisticsTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qPortVlanHCStatisticsTable", types.YChild{"Dot1qPortVlanHCStatisticsTable", &qBRIDGEMIB.Dot1qPortVlanHCStatisticsTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1qLearningConstraintsTable", types.YChild{"Dot1qLearningConstraintsTable", &qBRIDGEMIB.Dot1qLearningConstraintsTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1vProtocolGroupTable", types.YChild{"Dot1vProtocolGroupTable", &qBRIDGEMIB.Dot1vProtocolGroupTable})
    qBRIDGEMIB.EntityData.Children.Append("dot1vProtocolPortTable", types.YChild{"Dot1vProtocolPortTable", &qBRIDGEMIB.Dot1vProtocolPortTable})
    qBRIDGEMIB.EntityData.Leafs = types.NewOrderedMap()

    qBRIDGEMIB.EntityData.YListKeys = []string {}

    return &(qBRIDGEMIB.EntityData)
}

// QBRIDGEMIB_Dot1qBase
type QBRIDGEMIB_Dot1qBase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The version number of IEEE 802.1Q that this device supports. The type is
    // Dot1qVlanVersionNumber.
    Dot1qVlanVersionNumber interface{}

    // The maximum IEEE 802.1Q VLAN-ID that this device  supports. The type is
    // interface{} with range: 1..4094.
    Dot1qMaxVlanId interface{}

    // The maximum number of IEEE 802.1Q VLANs that this device supports. The type
    // is interface{} with range: 0..4294967295.
    Dot1qMaxSupportedVlans interface{}

    // The current number of IEEE 802.1Q VLANs that are configured in this device.
    // The type is interface{} with range: 0..4294967295.
    Dot1qNumVlans interface{}

    // The administrative status requested by management for GVRP.  The value
    // enabled(1) indicates that GVRP should be enabled on this device, on all
    // ports for which it has not been specifically disabled.  When disabled(2),
    // GVRP is disabled on all ports, and all GVRP packets will be forwarded
    // transparently.  This object affects all GVRP Applicant and Registrar state
    // machines.  A transition from disabled(2) to enabled(1) will cause a reset
    // of all GVRP state machines on all ports.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // EnabledStatus.
    Dot1qGvrpStatus interface{}
}

func (dot1qBase *QBRIDGEMIB_Dot1qBase) GetEntityData() *types.CommonEntityData {
    dot1qBase.EntityData.YFilter = dot1qBase.YFilter
    dot1qBase.EntityData.YangName = "dot1qBase"
    dot1qBase.EntityData.BundleName = "cisco_ios_xe"
    dot1qBase.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qBase.EntityData.SegmentPath = "dot1qBase"
    dot1qBase.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qBase.EntityData.SegmentPath
    dot1qBase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qBase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qBase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qBase.EntityData.Children = types.NewOrderedMap()
    dot1qBase.EntityData.Leafs = types.NewOrderedMap()
    dot1qBase.EntityData.Leafs.Append("dot1qVlanVersionNumber", types.YLeaf{"Dot1qVlanVersionNumber", dot1qBase.Dot1qVlanVersionNumber})
    dot1qBase.EntityData.Leafs.Append("dot1qMaxVlanId", types.YLeaf{"Dot1qMaxVlanId", dot1qBase.Dot1qMaxVlanId})
    dot1qBase.EntityData.Leafs.Append("dot1qMaxSupportedVlans", types.YLeaf{"Dot1qMaxSupportedVlans", dot1qBase.Dot1qMaxSupportedVlans})
    dot1qBase.EntityData.Leafs.Append("dot1qNumVlans", types.YLeaf{"Dot1qNumVlans", dot1qBase.Dot1qNumVlans})
    dot1qBase.EntityData.Leafs.Append("dot1qGvrpStatus", types.YLeaf{"Dot1qGvrpStatus", dot1qBase.Dot1qGvrpStatus})

    dot1qBase.EntityData.YListKeys = []string {}

    return &(dot1qBase.EntityData)
}

// QBRIDGEMIB_Dot1qBase_Dot1qVlanVersionNumber represents supports.
type QBRIDGEMIB_Dot1qBase_Dot1qVlanVersionNumber string

const (
    QBRIDGEMIB_Dot1qBase_Dot1qVlanVersionNumber_version1 QBRIDGEMIB_Dot1qBase_Dot1qVlanVersionNumber = "version1"
)

// QBRIDGEMIB_Dot1qVlan
type QBRIDGEMIB_Dot1qVlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times a VLAN entry has been deleted from the
    // dot1qVlanCurrentTable (for any reason).  If an entry is deleted, then
    // inserted, and then deleted, this counter will be incremented by 2. The type
    // is interface{} with range: 0..4294967295.
    Dot1qVlanNumDeletes interface{}

    // The next available value for dot1qVlanIndex of a local VLAN entry in
    // dot1qVlanStaticTable.  This will report values >=4096 if a new Local VLAN
    // may be created or else the value 0 if this is not possible.  A row creation
    // operation in this table for an entry with a local VlanIndex value may fail
    // if the current value of this object is not used as the index.  Even if the
    // value read is used, there is no guarantee that it will still be the valid
    // index when the create operation is attempted; another manager may have
    // already got in during the intervening time interval. In this case,
    // dot1qNextFreeLocalVlanIndex should be re-read  and the creation re-tried
    // with the new value.  This value will automatically change when the current
    // value is used to create a new row. The type is interface{} with range:
    // 0..None | 4096..2147483647.
    Dot1qNextFreeLocalVlanIndex interface{}

    // The identity of the constraint set to which a VLAN belongs, if there is not
    // an explicit entry for that VLAN in dot1qLearningConstraintsTable.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is interface{} with range: 0..65535.
    Dot1qConstraintSetDefault interface{}

    // The type of constraint set to which a VLAN belongs, if there is not an
    // explicit entry for that VLAN in dot1qLearningConstraintsTable.  The types
    // are as defined for dot1qConstraintType.  The value of this object MUST be
    // retained across  reinitializations of the management system. The type is
    // Dot1qConstraintTypeDefault.
    Dot1qConstraintTypeDefault interface{}
}

func (dot1qVlan *QBRIDGEMIB_Dot1qVlan) GetEntityData() *types.CommonEntityData {
    dot1qVlan.EntityData.YFilter = dot1qVlan.YFilter
    dot1qVlan.EntityData.YangName = "dot1qVlan"
    dot1qVlan.EntityData.BundleName = "cisco_ios_xe"
    dot1qVlan.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qVlan.EntityData.SegmentPath = "dot1qVlan"
    dot1qVlan.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qVlan.EntityData.SegmentPath
    dot1qVlan.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qVlan.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qVlan.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qVlan.EntityData.Children = types.NewOrderedMap()
    dot1qVlan.EntityData.Leafs = types.NewOrderedMap()
    dot1qVlan.EntityData.Leafs.Append("dot1qVlanNumDeletes", types.YLeaf{"Dot1qVlanNumDeletes", dot1qVlan.Dot1qVlanNumDeletes})
    dot1qVlan.EntityData.Leafs.Append("dot1qNextFreeLocalVlanIndex", types.YLeaf{"Dot1qNextFreeLocalVlanIndex", dot1qVlan.Dot1qNextFreeLocalVlanIndex})
    dot1qVlan.EntityData.Leafs.Append("dot1qConstraintSetDefault", types.YLeaf{"Dot1qConstraintSetDefault", dot1qVlan.Dot1qConstraintSetDefault})
    dot1qVlan.EntityData.Leafs.Append("dot1qConstraintTypeDefault", types.YLeaf{"Dot1qConstraintTypeDefault", dot1qVlan.Dot1qConstraintTypeDefault})

    dot1qVlan.EntityData.YListKeys = []string {}

    return &(dot1qVlan.EntityData)
}

// QBRIDGEMIB_Dot1qVlan_Dot1qConstraintTypeDefault represents reinitializations of the management system.
type QBRIDGEMIB_Dot1qVlan_Dot1qConstraintTypeDefault string

const (
    QBRIDGEMIB_Dot1qVlan_Dot1qConstraintTypeDefault_independent QBRIDGEMIB_Dot1qVlan_Dot1qConstraintTypeDefault = "independent"

    QBRIDGEMIB_Dot1qVlan_Dot1qConstraintTypeDefault_shared QBRIDGEMIB_Dot1qVlan_Dot1qConstraintTypeDefault = "shared"
)

// QBRIDGEMIB_Dot1qFdbTable
// A table that contains configuration and control
// information for each Filtering Database currently
// operating on this device.  Entries in this table appear
// automatically when VLANs are assigned FDB IDs in the
// dot1qVlanCurrentTable.
type QBRIDGEMIB_Dot1qFdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a specific Filtering Database. The type is slice of
    // QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry.
    Dot1qFdbEntry []*QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry
}

func (dot1qFdbTable *QBRIDGEMIB_Dot1qFdbTable) GetEntityData() *types.CommonEntityData {
    dot1qFdbTable.EntityData.YFilter = dot1qFdbTable.YFilter
    dot1qFdbTable.EntityData.YangName = "dot1qFdbTable"
    dot1qFdbTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qFdbTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qFdbTable.EntityData.SegmentPath = "dot1qFdbTable"
    dot1qFdbTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qFdbTable.EntityData.SegmentPath
    dot1qFdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qFdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qFdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qFdbTable.EntityData.Children = types.NewOrderedMap()
    dot1qFdbTable.EntityData.Children.Append("dot1qFdbEntry", types.YChild{"Dot1qFdbEntry", nil})
    for i := range dot1qFdbTable.Dot1qFdbEntry {
        dot1qFdbTable.EntityData.Children.Append(types.GetSegmentPath(dot1qFdbTable.Dot1qFdbEntry[i]), types.YChild{"Dot1qFdbEntry", dot1qFdbTable.Dot1qFdbEntry[i]})
    }
    dot1qFdbTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qFdbTable.EntityData.YListKeys = []string {}

    return &(dot1qFdbTable.EntityData)
}

// QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry
// Information about a specific Filtering Database.
type QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The identity of this Filtering Database. The type
    // is interface{} with range: 0..4294967295.
    Dot1qFdbId interface{}

    // The current number of dynamic entries in this Filtering Database. The type
    // is interface{} with range: 0..4294967295.
    Dot1qFdbDynamicCount interface{}
}

func (dot1qFdbEntry *QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry) GetEntityData() *types.CommonEntityData {
    dot1qFdbEntry.EntityData.YFilter = dot1qFdbEntry.YFilter
    dot1qFdbEntry.EntityData.YangName = "dot1qFdbEntry"
    dot1qFdbEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qFdbEntry.EntityData.ParentYangName = "dot1qFdbTable"
    dot1qFdbEntry.EntityData.SegmentPath = "dot1qFdbEntry" + types.AddKeyToken(dot1qFdbEntry.Dot1qFdbId, "dot1qFdbId")
    dot1qFdbEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qFdbTable/" + dot1qFdbEntry.EntityData.SegmentPath
    dot1qFdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qFdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qFdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qFdbEntry.EntityData.Children = types.NewOrderedMap()
    dot1qFdbEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qFdbEntry.EntityData.Leafs.Append("dot1qFdbId", types.YLeaf{"Dot1qFdbId", dot1qFdbEntry.Dot1qFdbId})
    dot1qFdbEntry.EntityData.Leafs.Append("dot1qFdbDynamicCount", types.YLeaf{"Dot1qFdbDynamicCount", dot1qFdbEntry.Dot1qFdbDynamicCount})

    dot1qFdbEntry.EntityData.YListKeys = []string {"Dot1qFdbId"}

    return &(dot1qFdbEntry.EntityData)
}

// QBRIDGEMIB_Dot1qTpFdbTable
// A table that contains information about unicast entries
// for which the device has forwarding and/or filtering
// information.  This information is used by the
// transparent bridging function in determining how to
// propagate a received frame.
type QBRIDGEMIB_Dot1qTpFdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a specific unicast MAC address for which the device has
    // some forwarding and/or filtering information. The type is slice of
    // QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry.
    Dot1qTpFdbEntry []*QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry
}

func (dot1qTpFdbTable *QBRIDGEMIB_Dot1qTpFdbTable) GetEntityData() *types.CommonEntityData {
    dot1qTpFdbTable.EntityData.YFilter = dot1qTpFdbTable.YFilter
    dot1qTpFdbTable.EntityData.YangName = "dot1qTpFdbTable"
    dot1qTpFdbTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qTpFdbTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qTpFdbTable.EntityData.SegmentPath = "dot1qTpFdbTable"
    dot1qTpFdbTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qTpFdbTable.EntityData.SegmentPath
    dot1qTpFdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qTpFdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qTpFdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qTpFdbTable.EntityData.Children = types.NewOrderedMap()
    dot1qTpFdbTable.EntityData.Children.Append("dot1qTpFdbEntry", types.YChild{"Dot1qTpFdbEntry", nil})
    for i := range dot1qTpFdbTable.Dot1qTpFdbEntry {
        dot1qTpFdbTable.EntityData.Children.Append(types.GetSegmentPath(dot1qTpFdbTable.Dot1qTpFdbEntry[i]), types.YChild{"Dot1qTpFdbEntry", dot1qTpFdbTable.Dot1qTpFdbEntry[i]})
    }
    dot1qTpFdbTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qTpFdbTable.EntityData.YListKeys = []string {}

    return &(dot1qTpFdbTable.EntityData)
}

// QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry
// Information about a specific unicast MAC address for
// which the device has some forwarding and/or filtering
// information.
type QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to q_bridge_mib.QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry_Dot1qFdbId
    Dot1qFdbId interface{}

    // This attribute is a key. A unicast MAC address for which the device has
    // forwarding and/or filtering information. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1qTpFdbAddress interface{}

    // Either the value '0', or the port number of the port on which a frame
    // having a source address equal to the value of the corresponding instance of
    // dot1qTpFdbAddress has been seen.  A value of '0' indicates that the port
    // number has not been learned but that the device does have some
    // forwarding/filtering information about this address (e.g., in the
    // dot1qStaticUnicastTable). Implementors are encouraged to assign the port
    // value to this object whenever it is learned, even for addresses for which
    // the corresponding value of dot1qTpFdbStatus is not learned(3). The type is
    // interface{} with range: 0..65535.
    Dot1qTpFdbPort interface{}

    // The status of this entry.  The meanings of the values are:     other(1) -
    // none of the following.  This may include         the case where some other
    // MIB object (not the         corresponding instance of dot1qTpFdbPort, nor
    // an         entry in the dot1qStaticUnicastTable) is being         used to
    // determine if and how frames addressed to         the value of the
    // corresponding instance of         dot1qTpFdbAddress are being forwarded.   
    // invalid(2) - this entry is no longer valid (e.g., it          was learned
    // but has since aged out), but has not         yet been flushed from the
    // table.     learned(3) - the value of the corresponding instance         of
    // dot1qTpFdbPort was learned and is being used.     self(4) - the value of
    // the corresponding instance of         dot1qTpFdbAddress represents one of
    // the device's         addresses.  The corresponding instance of        
    // dot1qTpFdbPort indicates which of the device's         ports has this
    // address.     mgmt(5) - the value of the corresponding instance of        
    // dot1qTpFdbAddress is also the value of an         existing instance of
    // dot1qStaticAddress. The type is Dot1qTpFdbStatus.
    Dot1qTpFdbStatus interface{}
}

func (dot1qTpFdbEntry *QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry) GetEntityData() *types.CommonEntityData {
    dot1qTpFdbEntry.EntityData.YFilter = dot1qTpFdbEntry.YFilter
    dot1qTpFdbEntry.EntityData.YangName = "dot1qTpFdbEntry"
    dot1qTpFdbEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qTpFdbEntry.EntityData.ParentYangName = "dot1qTpFdbTable"
    dot1qTpFdbEntry.EntityData.SegmentPath = "dot1qTpFdbEntry" + types.AddKeyToken(dot1qTpFdbEntry.Dot1qFdbId, "dot1qFdbId") + types.AddKeyToken(dot1qTpFdbEntry.Dot1qTpFdbAddress, "dot1qTpFdbAddress")
    dot1qTpFdbEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qTpFdbTable/" + dot1qTpFdbEntry.EntityData.SegmentPath
    dot1qTpFdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qTpFdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qTpFdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qTpFdbEntry.EntityData.Children = types.NewOrderedMap()
    dot1qTpFdbEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qTpFdbEntry.EntityData.Leafs.Append("dot1qFdbId", types.YLeaf{"Dot1qFdbId", dot1qTpFdbEntry.Dot1qFdbId})
    dot1qTpFdbEntry.EntityData.Leafs.Append("dot1qTpFdbAddress", types.YLeaf{"Dot1qTpFdbAddress", dot1qTpFdbEntry.Dot1qTpFdbAddress})
    dot1qTpFdbEntry.EntityData.Leafs.Append("dot1qTpFdbPort", types.YLeaf{"Dot1qTpFdbPort", dot1qTpFdbEntry.Dot1qTpFdbPort})
    dot1qTpFdbEntry.EntityData.Leafs.Append("dot1qTpFdbStatus", types.YLeaf{"Dot1qTpFdbStatus", dot1qTpFdbEntry.Dot1qTpFdbStatus})

    dot1qTpFdbEntry.EntityData.YListKeys = []string {"Dot1qFdbId", "Dot1qTpFdbAddress"}

    return &(dot1qTpFdbEntry.EntityData)
}

// QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus represents         existing instance of dot1qStaticAddress.
type QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus string

const (
    QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus_other QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus = "other"

    QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus_invalid QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus = "invalid"

    QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus_learned QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus = "learned"

    QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus_self QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus = "self"

    QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus_mgmt QBRIDGEMIB_Dot1qTpFdbTable_Dot1qTpFdbEntry_Dot1qTpFdbStatus = "mgmt"
)

// QBRIDGEMIB_Dot1qTpGroupTable
// A table containing filtering information for VLANs
// configured into the bridge by (local or network)
// management, or learned dynamically, specifying the set of
// ports to which frames received on a VLAN for this FDB
// and containing a specific Group destination address are
// allowed to be forwarded.
type QBRIDGEMIB_Dot1qTpGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the bridge by management, or learned
    // dynamically, specifying the set of ports to which frames received on a VLAN
    // and containing a specific Group destination address are allowed to be
    // forwarded.  The subset of these ports learned dynamically is also provided.
    // The type is slice of QBRIDGEMIB_Dot1qTpGroupTable_Dot1qTpGroupEntry.
    Dot1qTpGroupEntry []*QBRIDGEMIB_Dot1qTpGroupTable_Dot1qTpGroupEntry
}

func (dot1qTpGroupTable *QBRIDGEMIB_Dot1qTpGroupTable) GetEntityData() *types.CommonEntityData {
    dot1qTpGroupTable.EntityData.YFilter = dot1qTpGroupTable.YFilter
    dot1qTpGroupTable.EntityData.YangName = "dot1qTpGroupTable"
    dot1qTpGroupTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qTpGroupTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qTpGroupTable.EntityData.SegmentPath = "dot1qTpGroupTable"
    dot1qTpGroupTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qTpGroupTable.EntityData.SegmentPath
    dot1qTpGroupTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qTpGroupTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qTpGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qTpGroupTable.EntityData.Children = types.NewOrderedMap()
    dot1qTpGroupTable.EntityData.Children.Append("dot1qTpGroupEntry", types.YChild{"Dot1qTpGroupEntry", nil})
    for i := range dot1qTpGroupTable.Dot1qTpGroupEntry {
        dot1qTpGroupTable.EntityData.Children.Append(types.GetSegmentPath(dot1qTpGroupTable.Dot1qTpGroupEntry[i]), types.YChild{"Dot1qTpGroupEntry", dot1qTpGroupTable.Dot1qTpGroupEntry[i]})
    }
    dot1qTpGroupTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qTpGroupTable.EntityData.YListKeys = []string {}

    return &(dot1qTpGroupTable.EntityData)
}

// QBRIDGEMIB_Dot1qTpGroupTable_Dot1qTpGroupEntry
// Filtering information configured into the bridge by
// management, or learned dynamically, specifying the set of
// ports to which frames received on a VLAN and containing
// a specific Group destination address are allowed to be
// forwarded.  The subset of these ports learned dynamically
// is also provided.
type QBRIDGEMIB_Dot1qTpGroupTable_Dot1qTpGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // This attribute is a key. The destination Group MAC address in a frame to
    // which this entry's filtering information applies. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1qTpGroupAddress interface{}

    // The complete set of ports, in this VLAN, to which frames destined for this
    // Group MAC address are currently being explicitly forwarded.  This does not
    // include ports for which this address is only implicitly forwarded, in the
    // dot1qForwardAllPorts list. The type is string.
    Dot1qTpGroupEgressPorts interface{}

    // The subset of ports in dot1qTpGroupEgressPorts that were learned by GMRP or
    // some other dynamic mechanism, in this Filtering database. The type is
    // string.
    Dot1qTpGroupLearnt interface{}
}

func (dot1qTpGroupEntry *QBRIDGEMIB_Dot1qTpGroupTable_Dot1qTpGroupEntry) GetEntityData() *types.CommonEntityData {
    dot1qTpGroupEntry.EntityData.YFilter = dot1qTpGroupEntry.YFilter
    dot1qTpGroupEntry.EntityData.YangName = "dot1qTpGroupEntry"
    dot1qTpGroupEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qTpGroupEntry.EntityData.ParentYangName = "dot1qTpGroupTable"
    dot1qTpGroupEntry.EntityData.SegmentPath = "dot1qTpGroupEntry" + types.AddKeyToken(dot1qTpGroupEntry.Dot1qVlanIndex, "dot1qVlanIndex") + types.AddKeyToken(dot1qTpGroupEntry.Dot1qTpGroupAddress, "dot1qTpGroupAddress")
    dot1qTpGroupEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qTpGroupTable/" + dot1qTpGroupEntry.EntityData.SegmentPath
    dot1qTpGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qTpGroupEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qTpGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qTpGroupEntry.EntityData.Children = types.NewOrderedMap()
    dot1qTpGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qTpGroupEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qTpGroupEntry.Dot1qVlanIndex})
    dot1qTpGroupEntry.EntityData.Leafs.Append("dot1qTpGroupAddress", types.YLeaf{"Dot1qTpGroupAddress", dot1qTpGroupEntry.Dot1qTpGroupAddress})
    dot1qTpGroupEntry.EntityData.Leafs.Append("dot1qTpGroupEgressPorts", types.YLeaf{"Dot1qTpGroupEgressPorts", dot1qTpGroupEntry.Dot1qTpGroupEgressPorts})
    dot1qTpGroupEntry.EntityData.Leafs.Append("dot1qTpGroupLearnt", types.YLeaf{"Dot1qTpGroupLearnt", dot1qTpGroupEntry.Dot1qTpGroupLearnt})

    dot1qTpGroupEntry.EntityData.YListKeys = []string {"Dot1qVlanIndex", "Dot1qTpGroupAddress"}

    return &(dot1qTpGroupEntry.EntityData)
}

// QBRIDGEMIB_Dot1qForwardAllTable
// A table containing forwarding information for each
// 
// VLAN, specifying the set of ports to which forwarding of
// all multicasts applies, configured statically by
// management or dynamically by GMRP.  An entry appears in
// this table for all VLANs that are currently
// instantiated.
type QBRIDGEMIB_Dot1qForwardAllTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding information for a VLAN, specifying the set of ports to which all
    // multicasts should be forwarded, configured statically by management or
    // dynamically by GMRP. The type is slice of
    // QBRIDGEMIB_Dot1qForwardAllTable_Dot1qForwardAllEntry.
    Dot1qForwardAllEntry []*QBRIDGEMIB_Dot1qForwardAllTable_Dot1qForwardAllEntry
}

func (dot1qForwardAllTable *QBRIDGEMIB_Dot1qForwardAllTable) GetEntityData() *types.CommonEntityData {
    dot1qForwardAllTable.EntityData.YFilter = dot1qForwardAllTable.YFilter
    dot1qForwardAllTable.EntityData.YangName = "dot1qForwardAllTable"
    dot1qForwardAllTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qForwardAllTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qForwardAllTable.EntityData.SegmentPath = "dot1qForwardAllTable"
    dot1qForwardAllTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qForwardAllTable.EntityData.SegmentPath
    dot1qForwardAllTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qForwardAllTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qForwardAllTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qForwardAllTable.EntityData.Children = types.NewOrderedMap()
    dot1qForwardAllTable.EntityData.Children.Append("dot1qForwardAllEntry", types.YChild{"Dot1qForwardAllEntry", nil})
    for i := range dot1qForwardAllTable.Dot1qForwardAllEntry {
        dot1qForwardAllTable.EntityData.Children.Append(types.GetSegmentPath(dot1qForwardAllTable.Dot1qForwardAllEntry[i]), types.YChild{"Dot1qForwardAllEntry", dot1qForwardAllTable.Dot1qForwardAllEntry[i]})
    }
    dot1qForwardAllTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qForwardAllTable.EntityData.YListKeys = []string {}

    return &(dot1qForwardAllTable.EntityData)
}

// QBRIDGEMIB_Dot1qForwardAllTable_Dot1qForwardAllEntry
// Forwarding information for a VLAN, specifying the set
// of ports to which all multicasts should be forwarded,
// configured statically by management or dynamically by
// GMRP.
type QBRIDGEMIB_Dot1qForwardAllTable_Dot1qForwardAllEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // The complete set of ports in this VLAN to which all multicast
    // group-addressed frames are to be forwarded. This includes ports for which
    // this need has been determined dynamically by GMRP, or configured statically
    // by management. The type is string.
    Dot1qForwardAllPorts interface{}

    // The set of ports configured by management in this VLAN to which all
    // multicast group-addressed frames are to be forwarded.  Ports entered in
    // this list will also appear in the complete set shown by
    // dot1qForwardAllPorts.  This value will be restored after the device is
    // reset.  This only applies to ports that are members of the VLAN, defined by
    // dot1qVlanCurrentEgressPorts.  A port may not be added in this set if it is
    // already a member of the set of ports in dot1qForwardAllForbiddenPorts.  The
    // default value is a string of ones of appropriate length, to indicate the
    // standard behaviour of using basic filtering services, i.e., forward all
    // multicasts to all ports.  The value of this object MUST be retained across
    // reinitializations of the management system. The type is string.
    Dot1qForwardAllStaticPorts interface{}

    // The set of ports configured by management in this VLAN for which the
    // Service Requirement attribute Forward All Multicast Groups may not be
    // dynamically registered by GMRP.  This value will be restored after the
    // device is reset.  A port may not be added in this set if it is already a
    // member of the set of ports in dot1qForwardAllStaticPorts.  The default
    // value is a string of zeros of appropriate length.  The value of this object
    // MUST be retained across reinitializations of the management system. The
    // type is string.
    Dot1qForwardAllForbiddenPorts interface{}
}

func (dot1qForwardAllEntry *QBRIDGEMIB_Dot1qForwardAllTable_Dot1qForwardAllEntry) GetEntityData() *types.CommonEntityData {
    dot1qForwardAllEntry.EntityData.YFilter = dot1qForwardAllEntry.YFilter
    dot1qForwardAllEntry.EntityData.YangName = "dot1qForwardAllEntry"
    dot1qForwardAllEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qForwardAllEntry.EntityData.ParentYangName = "dot1qForwardAllTable"
    dot1qForwardAllEntry.EntityData.SegmentPath = "dot1qForwardAllEntry" + types.AddKeyToken(dot1qForwardAllEntry.Dot1qVlanIndex, "dot1qVlanIndex")
    dot1qForwardAllEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qForwardAllTable/" + dot1qForwardAllEntry.EntityData.SegmentPath
    dot1qForwardAllEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qForwardAllEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qForwardAllEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qForwardAllEntry.EntityData.Children = types.NewOrderedMap()
    dot1qForwardAllEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qForwardAllEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qForwardAllEntry.Dot1qVlanIndex})
    dot1qForwardAllEntry.EntityData.Leafs.Append("dot1qForwardAllPorts", types.YLeaf{"Dot1qForwardAllPorts", dot1qForwardAllEntry.Dot1qForwardAllPorts})
    dot1qForwardAllEntry.EntityData.Leafs.Append("dot1qForwardAllStaticPorts", types.YLeaf{"Dot1qForwardAllStaticPorts", dot1qForwardAllEntry.Dot1qForwardAllStaticPorts})
    dot1qForwardAllEntry.EntityData.Leafs.Append("dot1qForwardAllForbiddenPorts", types.YLeaf{"Dot1qForwardAllForbiddenPorts", dot1qForwardAllEntry.Dot1qForwardAllForbiddenPorts})

    dot1qForwardAllEntry.EntityData.YListKeys = []string {"Dot1qVlanIndex"}

    return &(dot1qForwardAllEntry.EntityData)
}

// QBRIDGEMIB_Dot1qForwardUnregisteredTable
// A table containing forwarding information for each
// VLAN, specifying the set of ports to which forwarding of
// multicast group-addressed frames for which no
// more specific forwarding information applies.  This is
// configured statically by management and determined
// dynamically by GMRP.  An entry appears in this table for
// all VLANs that are currently instantiated.
type QBRIDGEMIB_Dot1qForwardUnregisteredTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding information for a VLAN, specifying the set of ports to which all
    // multicasts for which there is no more specific forwarding information shall
    // be forwarded. This is configured statically by management or dynamically by
    // GMRP. The type is slice of
    // QBRIDGEMIB_Dot1qForwardUnregisteredTable_Dot1qForwardUnregisteredEntry.
    Dot1qForwardUnregisteredEntry []*QBRIDGEMIB_Dot1qForwardUnregisteredTable_Dot1qForwardUnregisteredEntry
}

func (dot1qForwardUnregisteredTable *QBRIDGEMIB_Dot1qForwardUnregisteredTable) GetEntityData() *types.CommonEntityData {
    dot1qForwardUnregisteredTable.EntityData.YFilter = dot1qForwardUnregisteredTable.YFilter
    dot1qForwardUnregisteredTable.EntityData.YangName = "dot1qForwardUnregisteredTable"
    dot1qForwardUnregisteredTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qForwardUnregisteredTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qForwardUnregisteredTable.EntityData.SegmentPath = "dot1qForwardUnregisteredTable"
    dot1qForwardUnregisteredTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qForwardUnregisteredTable.EntityData.SegmentPath
    dot1qForwardUnregisteredTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qForwardUnregisteredTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qForwardUnregisteredTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qForwardUnregisteredTable.EntityData.Children = types.NewOrderedMap()
    dot1qForwardUnregisteredTable.EntityData.Children.Append("dot1qForwardUnregisteredEntry", types.YChild{"Dot1qForwardUnregisteredEntry", nil})
    for i := range dot1qForwardUnregisteredTable.Dot1qForwardUnregisteredEntry {
        dot1qForwardUnregisteredTable.EntityData.Children.Append(types.GetSegmentPath(dot1qForwardUnregisteredTable.Dot1qForwardUnregisteredEntry[i]), types.YChild{"Dot1qForwardUnregisteredEntry", dot1qForwardUnregisteredTable.Dot1qForwardUnregisteredEntry[i]})
    }
    dot1qForwardUnregisteredTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qForwardUnregisteredTable.EntityData.YListKeys = []string {}

    return &(dot1qForwardUnregisteredTable.EntityData)
}

// QBRIDGEMIB_Dot1qForwardUnregisteredTable_Dot1qForwardUnregisteredEntry
// Forwarding information for a VLAN, specifying the set
// of ports to which all multicasts for which there is no
// more specific forwarding information shall be forwarded.
// This is configured statically by management or
// dynamically by GMRP.
type QBRIDGEMIB_Dot1qForwardUnregisteredTable_Dot1qForwardUnregisteredEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // The complete set of ports in this VLAN to which multicast group-addressed
    // frames for which there is no more specific forwarding information will be
    // forwarded. This includes ports for which this need has been determined
    // dynamically by GMRP, or configured statically by management. The type is
    // string.
    Dot1qForwardUnregisteredPorts interface{}

    // The set of ports configured by management, in this VLAN, to which multicast
    // group-addressed frames for which there is no more specific forwarding
    // information  are to be forwarded.  Ports entered in this list will also
    // appear in the complete set shown by dot1qForwardUnregisteredPorts.  This
    // value will be restored after the device is reset.  A port may not be added
    // in this set if it is already a member of the set of ports in
    // dot1qForwardUnregisteredForbiddenPorts.  The default value is a string of
    // zeros of appropriate length, although this has no effect with the default
    // value of dot1qForwardAllStaticPorts.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // string.
    Dot1qForwardUnregisteredStaticPorts interface{}

    // The set of ports configured by management in this VLAN for which the
    // Service Requirement attribute Forward Unregistered Multicast Groups may not
    // be dynamically registered by GMRP.  This value will be restored after the
    // device is reset.  A port may not be added in this set if it is already a
    // member of the set of ports in dot1qForwardUnregisteredStaticPorts.  The
    // default value is a string of zeros of appropriate length.  The value of
    // this object MUST be retained across reinitializations of the management
    // system. The type is string.
    Dot1qForwardUnregisteredForbiddenPorts interface{}
}

func (dot1qForwardUnregisteredEntry *QBRIDGEMIB_Dot1qForwardUnregisteredTable_Dot1qForwardUnregisteredEntry) GetEntityData() *types.CommonEntityData {
    dot1qForwardUnregisteredEntry.EntityData.YFilter = dot1qForwardUnregisteredEntry.YFilter
    dot1qForwardUnregisteredEntry.EntityData.YangName = "dot1qForwardUnregisteredEntry"
    dot1qForwardUnregisteredEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qForwardUnregisteredEntry.EntityData.ParentYangName = "dot1qForwardUnregisteredTable"
    dot1qForwardUnregisteredEntry.EntityData.SegmentPath = "dot1qForwardUnregisteredEntry" + types.AddKeyToken(dot1qForwardUnregisteredEntry.Dot1qVlanIndex, "dot1qVlanIndex")
    dot1qForwardUnregisteredEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qForwardUnregisteredTable/" + dot1qForwardUnregisteredEntry.EntityData.SegmentPath
    dot1qForwardUnregisteredEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qForwardUnregisteredEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qForwardUnregisteredEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qForwardUnregisteredEntry.EntityData.Children = types.NewOrderedMap()
    dot1qForwardUnregisteredEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qForwardUnregisteredEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qForwardUnregisteredEntry.Dot1qVlanIndex})
    dot1qForwardUnregisteredEntry.EntityData.Leafs.Append("dot1qForwardUnregisteredPorts", types.YLeaf{"Dot1qForwardUnregisteredPorts", dot1qForwardUnregisteredEntry.Dot1qForwardUnregisteredPorts})
    dot1qForwardUnregisteredEntry.EntityData.Leafs.Append("dot1qForwardUnregisteredStaticPorts", types.YLeaf{"Dot1qForwardUnregisteredStaticPorts", dot1qForwardUnregisteredEntry.Dot1qForwardUnregisteredStaticPorts})
    dot1qForwardUnregisteredEntry.EntityData.Leafs.Append("dot1qForwardUnregisteredForbiddenPorts", types.YLeaf{"Dot1qForwardUnregisteredForbiddenPorts", dot1qForwardUnregisteredEntry.Dot1qForwardUnregisteredForbiddenPorts})

    dot1qForwardUnregisteredEntry.EntityData.YListKeys = []string {"Dot1qVlanIndex"}

    return &(dot1qForwardUnregisteredEntry.EntityData)
}

// QBRIDGEMIB_Dot1qStaticUnicastTable
// A table containing filtering information for Unicast
// MAC addresses for each Filtering Database, configured
// into the device by (local or network) management
// specifying the set of ports to which frames received
// from specific ports and containing specific unicast
// destination addresses are allowed to be forwarded.  A
// value of zero in this table (as the port number from
// 
// which frames with a specific destination address are
// received) is used to specify all ports for which there
// is no specific entry in this table for that particular
// destination address.  Entries are valid for unicast
// addresses only.
type QBRIDGEMIB_Dot1qStaticUnicastTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the device by (local or network)
    // management specifying the set of ports to which frames received from a
    // specific port and containing a specific unicast destination address are
    // allowed to be forwarded. The type is slice of
    // QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry.
    Dot1qStaticUnicastEntry []*QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry
}

func (dot1qStaticUnicastTable *QBRIDGEMIB_Dot1qStaticUnicastTable) GetEntityData() *types.CommonEntityData {
    dot1qStaticUnicastTable.EntityData.YFilter = dot1qStaticUnicastTable.YFilter
    dot1qStaticUnicastTable.EntityData.YangName = "dot1qStaticUnicastTable"
    dot1qStaticUnicastTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qStaticUnicastTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qStaticUnicastTable.EntityData.SegmentPath = "dot1qStaticUnicastTable"
    dot1qStaticUnicastTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qStaticUnicastTable.EntityData.SegmentPath
    dot1qStaticUnicastTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qStaticUnicastTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qStaticUnicastTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qStaticUnicastTable.EntityData.Children = types.NewOrderedMap()
    dot1qStaticUnicastTable.EntityData.Children.Append("dot1qStaticUnicastEntry", types.YChild{"Dot1qStaticUnicastEntry", nil})
    for i := range dot1qStaticUnicastTable.Dot1qStaticUnicastEntry {
        dot1qStaticUnicastTable.EntityData.Children.Append(types.GetSegmentPath(dot1qStaticUnicastTable.Dot1qStaticUnicastEntry[i]), types.YChild{"Dot1qStaticUnicastEntry", dot1qStaticUnicastTable.Dot1qStaticUnicastEntry[i]})
    }
    dot1qStaticUnicastTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qStaticUnicastTable.EntityData.YListKeys = []string {}

    return &(dot1qStaticUnicastTable.EntityData)
}

// QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry
// Filtering information configured into the device by
// (local or network) management specifying the set of
// ports to which frames received from a specific port and
// containing a specific unicast destination address are
// allowed to be forwarded.
type QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to q_bridge_mib.QBRIDGEMIB_Dot1qFdbTable_Dot1qFdbEntry_Dot1qFdbId
    Dot1qFdbId interface{}

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object must take the
    // value of a unicast address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1qStaticUnicastAddress interface{}

    // This attribute is a key. Either the value '0' or the port number of the
    // port from which a frame must be received in order for this entry's
    // filtering information to apply.  A value of zero indicates that this entry
    // applies on all ports of the device for which there is no other applicable
    // entry. The type is interface{} with range: 0..65535.
    Dot1qStaticUnicastReceivePort interface{}

    // The set of ports for which a frame with a specific unicast address will be
    // flooded in the event that it has not been learned.  It also specifies the
    // set of ports on which a specific unicast address may be dynamically
    // learned.  The dot1qTpFdbTable will have an equivalent entry with a
    // dot1qTpFdbPort value of '0' until this address has been learned, at which
    // point it will be updated with the port the address has been seen on.  This
    // only applies to ports that are members of the VLAN, defined by
    // dot1qVlanCurrentEgressPorts.  The default value of this object is a string
    // of ones of appropriate length.  The value of this object MUST be retained
    // across reinitializations of the management system. The type is string.
    Dot1qStaticUnicastAllowedToGoTo interface{}

    // This object indicates the status of this entry. other(1) - this entry is
    // currently in use, but      the conditions under which it will remain     so
    // differ from the following values. invalid(2) - writing this value to the
    // object     removes the corresponding entry. permanent(3) - this entry is
    // currently in use     and will remain so after the next reset of     the
    // bridge. deleteOnReset(4) - this entry is currently in     use and will
    // remain so until the next     reset of the bridge. deleteOnTimeout(5) - this
    // entry is currently in     use and will remain so until it is aged out.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is Dot1qStaticUnicastStatus.
    Dot1qStaticUnicastStatus interface{}
}

func (dot1qStaticUnicastEntry *QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry) GetEntityData() *types.CommonEntityData {
    dot1qStaticUnicastEntry.EntityData.YFilter = dot1qStaticUnicastEntry.YFilter
    dot1qStaticUnicastEntry.EntityData.YangName = "dot1qStaticUnicastEntry"
    dot1qStaticUnicastEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qStaticUnicastEntry.EntityData.ParentYangName = "dot1qStaticUnicastTable"
    dot1qStaticUnicastEntry.EntityData.SegmentPath = "dot1qStaticUnicastEntry" + types.AddKeyToken(dot1qStaticUnicastEntry.Dot1qFdbId, "dot1qFdbId") + types.AddKeyToken(dot1qStaticUnicastEntry.Dot1qStaticUnicastAddress, "dot1qStaticUnicastAddress") + types.AddKeyToken(dot1qStaticUnicastEntry.Dot1qStaticUnicastReceivePort, "dot1qStaticUnicastReceivePort")
    dot1qStaticUnicastEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qStaticUnicastTable/" + dot1qStaticUnicastEntry.EntityData.SegmentPath
    dot1qStaticUnicastEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qStaticUnicastEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qStaticUnicastEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qStaticUnicastEntry.EntityData.Children = types.NewOrderedMap()
    dot1qStaticUnicastEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qStaticUnicastEntry.EntityData.Leafs.Append("dot1qFdbId", types.YLeaf{"Dot1qFdbId", dot1qStaticUnicastEntry.Dot1qFdbId})
    dot1qStaticUnicastEntry.EntityData.Leafs.Append("dot1qStaticUnicastAddress", types.YLeaf{"Dot1qStaticUnicastAddress", dot1qStaticUnicastEntry.Dot1qStaticUnicastAddress})
    dot1qStaticUnicastEntry.EntityData.Leafs.Append("dot1qStaticUnicastReceivePort", types.YLeaf{"Dot1qStaticUnicastReceivePort", dot1qStaticUnicastEntry.Dot1qStaticUnicastReceivePort})
    dot1qStaticUnicastEntry.EntityData.Leafs.Append("dot1qStaticUnicastAllowedToGoTo", types.YLeaf{"Dot1qStaticUnicastAllowedToGoTo", dot1qStaticUnicastEntry.Dot1qStaticUnicastAllowedToGoTo})
    dot1qStaticUnicastEntry.EntityData.Leafs.Append("dot1qStaticUnicastStatus", types.YLeaf{"Dot1qStaticUnicastStatus", dot1qStaticUnicastEntry.Dot1qStaticUnicastStatus})

    dot1qStaticUnicastEntry.EntityData.YListKeys = []string {"Dot1qFdbId", "Dot1qStaticUnicastAddress", "Dot1qStaticUnicastReceivePort"}

    return &(dot1qStaticUnicastEntry.EntityData)
}

// QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus represents reinitializations of the management system.
type QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus string

const (
    QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus_other QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus = "other"

    QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus_invalid QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus = "invalid"

    QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus_permanent QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus = "permanent"

    QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus_deleteOnReset QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus = "deleteOnReset"

    QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus_deleteOnTimeout QBRIDGEMIB_Dot1qStaticUnicastTable_Dot1qStaticUnicastEntry_Dot1qStaticUnicastStatus = "deleteOnTimeout"
)

// QBRIDGEMIB_Dot1qStaticMulticastTable
// A table containing filtering information for Multicast
// and Broadcast MAC addresses for each VLAN, configured
// into the device by (local or network) management
// specifying the set of ports to which frames received
// from specific ports and containing specific Multicast
// and Broadcast destination addresses are allowed to be
// forwarded.  A value of zero in this table (as the port
// number from which frames with a specific destination
// address are received) is used to specify all ports for
// which there is no specific entry in this table for that
// particular destination address.  Entries are valid for
// Multicast and Broadcast addresses only.
type QBRIDGEMIB_Dot1qStaticMulticastTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the device by (local or network)
    // management specifying the set of ports to which frames received from this
    // specific port  for this VLAN and containing this Multicast or Broadcast
    // destination address are allowed to be forwarded. The type is slice of
    // QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry.
    Dot1qStaticMulticastEntry []*QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry
}

func (dot1qStaticMulticastTable *QBRIDGEMIB_Dot1qStaticMulticastTable) GetEntityData() *types.CommonEntityData {
    dot1qStaticMulticastTable.EntityData.YFilter = dot1qStaticMulticastTable.YFilter
    dot1qStaticMulticastTable.EntityData.YangName = "dot1qStaticMulticastTable"
    dot1qStaticMulticastTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qStaticMulticastTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qStaticMulticastTable.EntityData.SegmentPath = "dot1qStaticMulticastTable"
    dot1qStaticMulticastTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qStaticMulticastTable.EntityData.SegmentPath
    dot1qStaticMulticastTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qStaticMulticastTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qStaticMulticastTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qStaticMulticastTable.EntityData.Children = types.NewOrderedMap()
    dot1qStaticMulticastTable.EntityData.Children.Append("dot1qStaticMulticastEntry", types.YChild{"Dot1qStaticMulticastEntry", nil})
    for i := range dot1qStaticMulticastTable.Dot1qStaticMulticastEntry {
        dot1qStaticMulticastTable.EntityData.Children.Append(types.GetSegmentPath(dot1qStaticMulticastTable.Dot1qStaticMulticastEntry[i]), types.YChild{"Dot1qStaticMulticastEntry", dot1qStaticMulticastTable.Dot1qStaticMulticastEntry[i]})
    }
    dot1qStaticMulticastTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qStaticMulticastTable.EntityData.YListKeys = []string {}

    return &(dot1qStaticMulticastTable.EntityData)
}

// QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry
// Filtering information configured into the device by
// (local or network) management specifying the set of
// ports to which frames received from this specific port
// 
// for this VLAN and containing this Multicast or Broadcast
// destination address are allowed to be forwarded.
type QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object must take the
    // value of a Multicast or Broadcast address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Dot1qStaticMulticastAddress interface{}

    // This attribute is a key. Either the value '0' or the port number of the
    // port from which a frame must be received in order for this entry's
    // filtering information to apply.  A value of zero indicates that this entry
    // applies on all ports of the device for which there is no other applicable
    // entry. The type is interface{} with range: 0..65535.
    Dot1qStaticMulticastReceivePort interface{}

    // The set of ports to which frames received from a specific port and destined
    // for a specific Multicast or Broadcast MAC address must be forwarded,
    // regardless of any dynamic information, e.g., from GMRP.  A port may not be
    // added in this set if it is already a member of the set of ports in
    // dot1qStaticMulticastForbiddenEgressPorts. The default value of this object
    // is a string of ones of appropriate length.  The value of this object MUST
    // be retained across reinitializations of the management system. The type is
    // string.
    Dot1qStaticMulticastStaticEgressPorts interface{}

    // The set of ports to which frames received from a specific port and destined
    // for a specific Multicast or Broadcast MAC address must not be forwarded,
    // regardless of any dynamic information, e.g., from GMRP.  A port may not be
    // added in this set if it is already a member of the set of ports in
    // dot1qStaticMulticastStaticEgressPorts. The default value of this object is
    // a string of zeros of appropriate length.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // string.
    Dot1qStaticMulticastForbiddenEgressPorts interface{}

    // This object indicates the status of this entry. other(1) - this entry is
    // currently in use, but     the conditions under which it will remain     so
    // differ from the following values.  invalid(2) - writing this value to the
    // object     removes the corresponding entry. permanent(3) - this entry is
    // currently in use     and will remain so after the next reset of     the
    // bridge. deleteOnReset(4) - this entry is currently in     use and will
    // remain so until the next     reset of the bridge. deleteOnTimeout(5) - this
    // entry is currently in     use and will remain so until it is aged out.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is Dot1qStaticMulticastStatus.
    Dot1qStaticMulticastStatus interface{}
}

func (dot1qStaticMulticastEntry *QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry) GetEntityData() *types.CommonEntityData {
    dot1qStaticMulticastEntry.EntityData.YFilter = dot1qStaticMulticastEntry.YFilter
    dot1qStaticMulticastEntry.EntityData.YangName = "dot1qStaticMulticastEntry"
    dot1qStaticMulticastEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qStaticMulticastEntry.EntityData.ParentYangName = "dot1qStaticMulticastTable"
    dot1qStaticMulticastEntry.EntityData.SegmentPath = "dot1qStaticMulticastEntry" + types.AddKeyToken(dot1qStaticMulticastEntry.Dot1qVlanIndex, "dot1qVlanIndex") + types.AddKeyToken(dot1qStaticMulticastEntry.Dot1qStaticMulticastAddress, "dot1qStaticMulticastAddress") + types.AddKeyToken(dot1qStaticMulticastEntry.Dot1qStaticMulticastReceivePort, "dot1qStaticMulticastReceivePort")
    dot1qStaticMulticastEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qStaticMulticastTable/" + dot1qStaticMulticastEntry.EntityData.SegmentPath
    dot1qStaticMulticastEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qStaticMulticastEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qStaticMulticastEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qStaticMulticastEntry.EntityData.Children = types.NewOrderedMap()
    dot1qStaticMulticastEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qStaticMulticastEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qStaticMulticastEntry.Dot1qVlanIndex})
    dot1qStaticMulticastEntry.EntityData.Leafs.Append("dot1qStaticMulticastAddress", types.YLeaf{"Dot1qStaticMulticastAddress", dot1qStaticMulticastEntry.Dot1qStaticMulticastAddress})
    dot1qStaticMulticastEntry.EntityData.Leafs.Append("dot1qStaticMulticastReceivePort", types.YLeaf{"Dot1qStaticMulticastReceivePort", dot1qStaticMulticastEntry.Dot1qStaticMulticastReceivePort})
    dot1qStaticMulticastEntry.EntityData.Leafs.Append("dot1qStaticMulticastStaticEgressPorts", types.YLeaf{"Dot1qStaticMulticastStaticEgressPorts", dot1qStaticMulticastEntry.Dot1qStaticMulticastStaticEgressPorts})
    dot1qStaticMulticastEntry.EntityData.Leafs.Append("dot1qStaticMulticastForbiddenEgressPorts", types.YLeaf{"Dot1qStaticMulticastForbiddenEgressPorts", dot1qStaticMulticastEntry.Dot1qStaticMulticastForbiddenEgressPorts})
    dot1qStaticMulticastEntry.EntityData.Leafs.Append("dot1qStaticMulticastStatus", types.YLeaf{"Dot1qStaticMulticastStatus", dot1qStaticMulticastEntry.Dot1qStaticMulticastStatus})

    dot1qStaticMulticastEntry.EntityData.YListKeys = []string {"Dot1qVlanIndex", "Dot1qStaticMulticastAddress", "Dot1qStaticMulticastReceivePort"}

    return &(dot1qStaticMulticastEntry.EntityData)
}

// QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus represents reinitializations of the management system.
type QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus string

const (
    QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus_other QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus = "other"

    QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus_invalid QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus = "invalid"

    QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus_permanent QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus = "permanent"

    QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus_deleteOnReset QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus = "deleteOnReset"

    QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus_deleteOnTimeout QBRIDGEMIB_Dot1qStaticMulticastTable_Dot1qStaticMulticastEntry_Dot1qStaticMulticastStatus = "deleteOnTimeout"
)

// QBRIDGEMIB_Dot1qVlanCurrentTable
// A table containing current configuration information
// for each VLAN currently configured into the device by
// (local or network) management, or dynamically created
// as a result of GVRP requests received.
type QBRIDGEMIB_Dot1qVlanCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information for a VLAN configured into the device by  (local or network)
    // management, or dynamically created as a result of GVRP requests received.
    // The type is slice of
    // QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry.
    Dot1qVlanCurrentEntry []*QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry
}

func (dot1qVlanCurrentTable *QBRIDGEMIB_Dot1qVlanCurrentTable) GetEntityData() *types.CommonEntityData {
    dot1qVlanCurrentTable.EntityData.YFilter = dot1qVlanCurrentTable.YFilter
    dot1qVlanCurrentTable.EntityData.YangName = "dot1qVlanCurrentTable"
    dot1qVlanCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qVlanCurrentTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qVlanCurrentTable.EntityData.SegmentPath = "dot1qVlanCurrentTable"
    dot1qVlanCurrentTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qVlanCurrentTable.EntityData.SegmentPath
    dot1qVlanCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qVlanCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qVlanCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qVlanCurrentTable.EntityData.Children = types.NewOrderedMap()
    dot1qVlanCurrentTable.EntityData.Children.Append("dot1qVlanCurrentEntry", types.YChild{"Dot1qVlanCurrentEntry", nil})
    for i := range dot1qVlanCurrentTable.Dot1qVlanCurrentEntry {
        dot1qVlanCurrentTable.EntityData.Children.Append(types.GetSegmentPath(dot1qVlanCurrentTable.Dot1qVlanCurrentEntry[i]), types.YChild{"Dot1qVlanCurrentEntry", dot1qVlanCurrentTable.Dot1qVlanCurrentEntry[i]})
    }
    dot1qVlanCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qVlanCurrentTable.EntityData.YListKeys = []string {}

    return &(dot1qVlanCurrentTable.EntityData)
}

// QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry
// Information for a VLAN configured into the device by
// 
// (local or network) management, or dynamically created
// as a result of GVRP requests received.
type QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Dot1qVlanTimeMark interface{}

    // This attribute is a key. The VLAN-ID or other identifier referring to this
    // VLAN. The type is interface{} with range: 0..4294967295.
    Dot1qVlanIndex interface{}

    // The Filtering Database used by this VLAN.  This is one of the dot1qFdbId
    // values in the dot1qFdbTable.  This value is allocated automatically by the
    // device whenever  the VLAN is created: either dynamically by GVRP, or by
    // management, in dot1qVlanStaticTable.  Allocation of this value follows the
    // learning constraints defined for this VLAN in
    // dot1qLearningConstraintsTable. The type is interface{} with range:
    // 0..4294967295.
    Dot1qVlanFdbId interface{}

    // The set of ports that are transmitting traffic for this VLAN as either
    // tagged or untagged frames. The type is string.
    Dot1qVlanCurrentEgressPorts interface{}

    // The set of ports that are transmitting traffic for this VLAN as untagged
    // frames. The type is string.
    Dot1qVlanCurrentUntaggedPorts interface{}

    // This object indicates the status of this entry. other(1) - this entry is
    // currently in use, but the     conditions under which it will remain so
    // differ     from the following values. permanent(2) - this entry,
    // corresponding to an entry     in dot1qVlanStaticTable, is currently in use
    // and     will remain so after the next reset of the     device.  The port
    // lists for this entry include     ports from the equivalent
    // dot1qVlanStaticTable     entry and ports learned dynamically.
    // dynamicGvrp(3) - this entry is currently in use      and will remain so
    // until removed by GVRP.  There     is no static entry for this VLAN, and it
    // will be     removed when the last port leaves the VLAN. The type is
    // Dot1qVlanStatus.
    Dot1qVlanStatus interface{}

    // The value of sysUpTime when this VLAN was created. The type is interface{}
    // with range: 0..4294967295.
    Dot1qVlanCreationTime interface{}
}

func (dot1qVlanCurrentEntry *QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry) GetEntityData() *types.CommonEntityData {
    dot1qVlanCurrentEntry.EntityData.YFilter = dot1qVlanCurrentEntry.YFilter
    dot1qVlanCurrentEntry.EntityData.YangName = "dot1qVlanCurrentEntry"
    dot1qVlanCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qVlanCurrentEntry.EntityData.ParentYangName = "dot1qVlanCurrentTable"
    dot1qVlanCurrentEntry.EntityData.SegmentPath = "dot1qVlanCurrentEntry" + types.AddKeyToken(dot1qVlanCurrentEntry.Dot1qVlanTimeMark, "dot1qVlanTimeMark") + types.AddKeyToken(dot1qVlanCurrentEntry.Dot1qVlanIndex, "dot1qVlanIndex")
    dot1qVlanCurrentEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qVlanCurrentTable/" + dot1qVlanCurrentEntry.EntityData.SegmentPath
    dot1qVlanCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qVlanCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qVlanCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qVlanCurrentEntry.EntityData.Children = types.NewOrderedMap()
    dot1qVlanCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qVlanCurrentEntry.EntityData.Leafs.Append("dot1qVlanTimeMark", types.YLeaf{"Dot1qVlanTimeMark", dot1qVlanCurrentEntry.Dot1qVlanTimeMark})
    dot1qVlanCurrentEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qVlanCurrentEntry.Dot1qVlanIndex})
    dot1qVlanCurrentEntry.EntityData.Leafs.Append("dot1qVlanFdbId", types.YLeaf{"Dot1qVlanFdbId", dot1qVlanCurrentEntry.Dot1qVlanFdbId})
    dot1qVlanCurrentEntry.EntityData.Leafs.Append("dot1qVlanCurrentEgressPorts", types.YLeaf{"Dot1qVlanCurrentEgressPorts", dot1qVlanCurrentEntry.Dot1qVlanCurrentEgressPorts})
    dot1qVlanCurrentEntry.EntityData.Leafs.Append("dot1qVlanCurrentUntaggedPorts", types.YLeaf{"Dot1qVlanCurrentUntaggedPorts", dot1qVlanCurrentEntry.Dot1qVlanCurrentUntaggedPorts})
    dot1qVlanCurrentEntry.EntityData.Leafs.Append("dot1qVlanStatus", types.YLeaf{"Dot1qVlanStatus", dot1qVlanCurrentEntry.Dot1qVlanStatus})
    dot1qVlanCurrentEntry.EntityData.Leafs.Append("dot1qVlanCreationTime", types.YLeaf{"Dot1qVlanCreationTime", dot1qVlanCurrentEntry.Dot1qVlanCreationTime})

    dot1qVlanCurrentEntry.EntityData.YListKeys = []string {"Dot1qVlanTimeMark", "Dot1qVlanIndex"}

    return &(dot1qVlanCurrentEntry.EntityData)
}

// QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus represents     removed when the last port leaves the VLAN.
type QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus string

const (
    QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus_other QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus = "other"

    QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus_permanent QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus = "permanent"

    QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus_dynamicGvrp QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanStatus = "dynamicGvrp"
)

// QBRIDGEMIB_Dot1qVlanStaticTable
// A table containing static configuration information for
// each VLAN configured into the device by (local or
// network) management.  All entries are permanent and will
// be restored after the device is reset.
type QBRIDGEMIB_Dot1qVlanStaticTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static information for a VLAN configured into the device by (local or
    // network) management. The type is slice of
    // QBRIDGEMIB_Dot1qVlanStaticTable_Dot1qVlanStaticEntry.
    Dot1qVlanStaticEntry []*QBRIDGEMIB_Dot1qVlanStaticTable_Dot1qVlanStaticEntry
}

func (dot1qVlanStaticTable *QBRIDGEMIB_Dot1qVlanStaticTable) GetEntityData() *types.CommonEntityData {
    dot1qVlanStaticTable.EntityData.YFilter = dot1qVlanStaticTable.YFilter
    dot1qVlanStaticTable.EntityData.YangName = "dot1qVlanStaticTable"
    dot1qVlanStaticTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qVlanStaticTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qVlanStaticTable.EntityData.SegmentPath = "dot1qVlanStaticTable"
    dot1qVlanStaticTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qVlanStaticTable.EntityData.SegmentPath
    dot1qVlanStaticTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qVlanStaticTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qVlanStaticTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qVlanStaticTable.EntityData.Children = types.NewOrderedMap()
    dot1qVlanStaticTable.EntityData.Children.Append("dot1qVlanStaticEntry", types.YChild{"Dot1qVlanStaticEntry", nil})
    for i := range dot1qVlanStaticTable.Dot1qVlanStaticEntry {
        dot1qVlanStaticTable.EntityData.Children.Append(types.GetSegmentPath(dot1qVlanStaticTable.Dot1qVlanStaticEntry[i]), types.YChild{"Dot1qVlanStaticEntry", dot1qVlanStaticTable.Dot1qVlanStaticEntry[i]})
    }
    dot1qVlanStaticTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qVlanStaticTable.EntityData.YListKeys = []string {}

    return &(dot1qVlanStaticTable.EntityData)
}

// QBRIDGEMIB_Dot1qVlanStaticTable_Dot1qVlanStaticEntry
// Static information for a VLAN configured into the
// device by (local or network) management.
type QBRIDGEMIB_Dot1qVlanStaticTable_Dot1qVlanStaticEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // An administratively assigned string, which may be used to identify the
    // VLAN. The type is string with length: 0..32.
    Dot1qVlanStaticName interface{}

    // The set of ports that are permanently assigned to the egress list for this
    // VLAN by management.  Changes to a bit in this object affect the per-port,
    // per-VLAN Registrar control for Registration Fixed for the relevant GVRP
    // state machine on each port.  A port may not be added in this set if it is
    // already a member of the set of ports in dot1qVlanForbiddenEgressPorts.  The
    // default value of this object is a string of zeros of appropriate length,
    // indicating not fixed. The type is string.
    Dot1qVlanStaticEgressPorts interface{}

    // The set of ports that are prohibited by management from being included in
    // the egress list for this VLAN. Changes to this object that cause a port to
    // be included or excluded affect the per-port, per-VLAN Registrar control for
    // Registration Forbidden for the relevant GVRP state machine on each port.  A
    // port may not be added in this set if it is already a member of the set of
    // ports in dot1qVlanStaticEgressPorts.  The default value of this object is a
    // string of zeros of appropriate length, excluding all ports from the
    // forbidden set. The type is string.
    Dot1qVlanForbiddenEgressPorts interface{}

    // The set of ports that should transmit egress packets for this VLAN as
    // untagged.  The default value of this object for the default VLAN
    // (dot1qVlanIndex = 1) is a string of appropriate length including all ports.
    // There is no specified default for other VLANs.  If a device agent cannot
    // support the set of ports being set, then it will reject the set operation
    // with an error.  For example, a manager might attempt to set more than one
    // VLAN to be untagged on egress where the device does not support this IEEE
    // 802.1Q option. The type is string.
    Dot1qVlanStaticUntaggedPorts interface{}

    // This object indicates the status of this entry. The type is RowStatus.
    Dot1qVlanStaticRowStatus interface{}
}

func (dot1qVlanStaticEntry *QBRIDGEMIB_Dot1qVlanStaticTable_Dot1qVlanStaticEntry) GetEntityData() *types.CommonEntityData {
    dot1qVlanStaticEntry.EntityData.YFilter = dot1qVlanStaticEntry.YFilter
    dot1qVlanStaticEntry.EntityData.YangName = "dot1qVlanStaticEntry"
    dot1qVlanStaticEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qVlanStaticEntry.EntityData.ParentYangName = "dot1qVlanStaticTable"
    dot1qVlanStaticEntry.EntityData.SegmentPath = "dot1qVlanStaticEntry" + types.AddKeyToken(dot1qVlanStaticEntry.Dot1qVlanIndex, "dot1qVlanIndex")
    dot1qVlanStaticEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qVlanStaticTable/" + dot1qVlanStaticEntry.EntityData.SegmentPath
    dot1qVlanStaticEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qVlanStaticEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qVlanStaticEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qVlanStaticEntry.EntityData.Children = types.NewOrderedMap()
    dot1qVlanStaticEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qVlanStaticEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qVlanStaticEntry.Dot1qVlanIndex})
    dot1qVlanStaticEntry.EntityData.Leafs.Append("dot1qVlanStaticName", types.YLeaf{"Dot1qVlanStaticName", dot1qVlanStaticEntry.Dot1qVlanStaticName})
    dot1qVlanStaticEntry.EntityData.Leafs.Append("dot1qVlanStaticEgressPorts", types.YLeaf{"Dot1qVlanStaticEgressPorts", dot1qVlanStaticEntry.Dot1qVlanStaticEgressPorts})
    dot1qVlanStaticEntry.EntityData.Leafs.Append("dot1qVlanForbiddenEgressPorts", types.YLeaf{"Dot1qVlanForbiddenEgressPorts", dot1qVlanStaticEntry.Dot1qVlanForbiddenEgressPorts})
    dot1qVlanStaticEntry.EntityData.Leafs.Append("dot1qVlanStaticUntaggedPorts", types.YLeaf{"Dot1qVlanStaticUntaggedPorts", dot1qVlanStaticEntry.Dot1qVlanStaticUntaggedPorts})
    dot1qVlanStaticEntry.EntityData.Leafs.Append("dot1qVlanStaticRowStatus", types.YLeaf{"Dot1qVlanStaticRowStatus", dot1qVlanStaticEntry.Dot1qVlanStaticRowStatus})

    dot1qVlanStaticEntry.EntityData.YListKeys = []string {"Dot1qVlanIndex"}

    return &(dot1qVlanStaticEntry.EntityData)
}

// QBRIDGEMIB_Dot1qPortVlanStatisticsTable
// A table containing per-port, per-VLAN statistics for
// traffic received.  Separate objects are provided for both the
// most-significant and least-significant bits of statistics
// counters for ports that are associated with this transparent
// bridge.  The most-significant bit objects are only required on
// high-capacity interfaces, as defined in the conformance clauses
// for these objects.  This mechanism is provided as a way to read
// 64-bit counters for agents that support only SNMPv1.
// 
// Note that the reporting of most-significant and least-
// significant counter bits separately runs the risk of missing
// an overflow of the lower bits in the interval between sampling.
// The manager must be aware of this possibility, even within the
// same varbindlist, when interpreting the results of a request or
// 
// asynchronous notification.
type QBRIDGEMIB_Dot1qPortVlanStatisticsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic statistics for a VLAN on an interface. The type is slice of
    // QBRIDGEMIB_Dot1qPortVlanStatisticsTable_Dot1qPortVlanStatisticsEntry.
    Dot1qPortVlanStatisticsEntry []*QBRIDGEMIB_Dot1qPortVlanStatisticsTable_Dot1qPortVlanStatisticsEntry
}

func (dot1qPortVlanStatisticsTable *QBRIDGEMIB_Dot1qPortVlanStatisticsTable) GetEntityData() *types.CommonEntityData {
    dot1qPortVlanStatisticsTable.EntityData.YFilter = dot1qPortVlanStatisticsTable.YFilter
    dot1qPortVlanStatisticsTable.EntityData.YangName = "dot1qPortVlanStatisticsTable"
    dot1qPortVlanStatisticsTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qPortVlanStatisticsTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qPortVlanStatisticsTable.EntityData.SegmentPath = "dot1qPortVlanStatisticsTable"
    dot1qPortVlanStatisticsTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qPortVlanStatisticsTable.EntityData.SegmentPath
    dot1qPortVlanStatisticsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qPortVlanStatisticsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qPortVlanStatisticsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qPortVlanStatisticsTable.EntityData.Children = types.NewOrderedMap()
    dot1qPortVlanStatisticsTable.EntityData.Children.Append("dot1qPortVlanStatisticsEntry", types.YChild{"Dot1qPortVlanStatisticsEntry", nil})
    for i := range dot1qPortVlanStatisticsTable.Dot1qPortVlanStatisticsEntry {
        dot1qPortVlanStatisticsTable.EntityData.Children.Append(types.GetSegmentPath(dot1qPortVlanStatisticsTable.Dot1qPortVlanStatisticsEntry[i]), types.YChild{"Dot1qPortVlanStatisticsEntry", dot1qPortVlanStatisticsTable.Dot1qPortVlanStatisticsEntry[i]})
    }
    dot1qPortVlanStatisticsTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qPortVlanStatisticsTable.EntityData.YListKeys = []string {}

    return &(dot1qPortVlanStatisticsTable.EntityData)
}

// QBRIDGEMIB_Dot1qPortVlanStatisticsTable_Dot1qPortVlanStatisticsEntry
// Traffic statistics for a VLAN on an interface.
type QBRIDGEMIB_Dot1qPortVlanStatisticsTable_Dot1qPortVlanStatisticsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1dBasePort
    Dot1dBasePort interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN.  Note that a frame received on this
    // port is counted by this object if and only if it is for a protocol being
    // processed by the local forwarding process for this VLAN.  This object
    // includes received bridge management frames classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP. The type is interface{} with range:
    // 0..4294967295.
    Dot1qTpVlanPortInFrames interface{}

    // The number of valid frames transmitted by this port to its segment from the
    // local forwarding process for this VLAN.  This includes bridge management
    // frames originated by this device that are classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP). The type is interface{} with range:
    // 0..4294967295.
    Dot1qTpVlanPortOutFrames interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN and that were discarded due to
    // VLAN-related reasons. Specifically, the IEEE 802.1Q counters for Discard
    // Inbound and Discard on Ingress Filtering. The type is interface{} with
    // range: 0..4294967295.
    Dot1qTpVlanPortInDiscards interface{}

    // The number of times the associated dot1qTpVlanPortInFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1qTpVlanPortInOverflowFrames interface{}

    // The number of times the associated dot1qTpVlanPortOutFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1qTpVlanPortOutOverflowFrames interface{}

    // The number of times the associated dot1qTpVlanPortInDiscards counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1qTpVlanPortInOverflowDiscards interface{}
}

func (dot1qPortVlanStatisticsEntry *QBRIDGEMIB_Dot1qPortVlanStatisticsTable_Dot1qPortVlanStatisticsEntry) GetEntityData() *types.CommonEntityData {
    dot1qPortVlanStatisticsEntry.EntityData.YFilter = dot1qPortVlanStatisticsEntry.YFilter
    dot1qPortVlanStatisticsEntry.EntityData.YangName = "dot1qPortVlanStatisticsEntry"
    dot1qPortVlanStatisticsEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qPortVlanStatisticsEntry.EntityData.ParentYangName = "dot1qPortVlanStatisticsTable"
    dot1qPortVlanStatisticsEntry.EntityData.SegmentPath = "dot1qPortVlanStatisticsEntry" + types.AddKeyToken(dot1qPortVlanStatisticsEntry.Dot1dBasePort, "dot1dBasePort") + types.AddKeyToken(dot1qPortVlanStatisticsEntry.Dot1qVlanIndex, "dot1qVlanIndex")
    dot1qPortVlanStatisticsEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qPortVlanStatisticsTable/" + dot1qPortVlanStatisticsEntry.EntityData.SegmentPath
    dot1qPortVlanStatisticsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qPortVlanStatisticsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qPortVlanStatisticsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qPortVlanStatisticsEntry.EntityData.Children = types.NewOrderedMap()
    dot1qPortVlanStatisticsEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", dot1qPortVlanStatisticsEntry.Dot1dBasePort})
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qPortVlanStatisticsEntry.Dot1qVlanIndex})
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortInFrames", types.YLeaf{"Dot1qTpVlanPortInFrames", dot1qPortVlanStatisticsEntry.Dot1qTpVlanPortInFrames})
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortOutFrames", types.YLeaf{"Dot1qTpVlanPortOutFrames", dot1qPortVlanStatisticsEntry.Dot1qTpVlanPortOutFrames})
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortInDiscards", types.YLeaf{"Dot1qTpVlanPortInDiscards", dot1qPortVlanStatisticsEntry.Dot1qTpVlanPortInDiscards})
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortInOverflowFrames", types.YLeaf{"Dot1qTpVlanPortInOverflowFrames", dot1qPortVlanStatisticsEntry.Dot1qTpVlanPortInOverflowFrames})
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortOutOverflowFrames", types.YLeaf{"Dot1qTpVlanPortOutOverflowFrames", dot1qPortVlanStatisticsEntry.Dot1qTpVlanPortOutOverflowFrames})
    dot1qPortVlanStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortInOverflowDiscards", types.YLeaf{"Dot1qTpVlanPortInOverflowDiscards", dot1qPortVlanStatisticsEntry.Dot1qTpVlanPortInOverflowDiscards})

    dot1qPortVlanStatisticsEntry.EntityData.YListKeys = []string {"Dot1dBasePort", "Dot1qVlanIndex"}

    return &(dot1qPortVlanStatisticsEntry.EntityData)
}

// QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable
// A table containing per-port, per-VLAN statistics for
// traffic on high-capacity interfaces.
type QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic statistics for a VLAN on a high-capacity interface. The type is
    // slice of
    // QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable_Dot1qPortVlanHCStatisticsEntry.
    Dot1qPortVlanHCStatisticsEntry []*QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable_Dot1qPortVlanHCStatisticsEntry
}

func (dot1qPortVlanHCStatisticsTable *QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable) GetEntityData() *types.CommonEntityData {
    dot1qPortVlanHCStatisticsTable.EntityData.YFilter = dot1qPortVlanHCStatisticsTable.YFilter
    dot1qPortVlanHCStatisticsTable.EntityData.YangName = "dot1qPortVlanHCStatisticsTable"
    dot1qPortVlanHCStatisticsTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qPortVlanHCStatisticsTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qPortVlanHCStatisticsTable.EntityData.SegmentPath = "dot1qPortVlanHCStatisticsTable"
    dot1qPortVlanHCStatisticsTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qPortVlanHCStatisticsTable.EntityData.SegmentPath
    dot1qPortVlanHCStatisticsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qPortVlanHCStatisticsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qPortVlanHCStatisticsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qPortVlanHCStatisticsTable.EntityData.Children = types.NewOrderedMap()
    dot1qPortVlanHCStatisticsTable.EntityData.Children.Append("dot1qPortVlanHCStatisticsEntry", types.YChild{"Dot1qPortVlanHCStatisticsEntry", nil})
    for i := range dot1qPortVlanHCStatisticsTable.Dot1qPortVlanHCStatisticsEntry {
        dot1qPortVlanHCStatisticsTable.EntityData.Children.Append(types.GetSegmentPath(dot1qPortVlanHCStatisticsTable.Dot1qPortVlanHCStatisticsEntry[i]), types.YChild{"Dot1qPortVlanHCStatisticsEntry", dot1qPortVlanHCStatisticsTable.Dot1qPortVlanHCStatisticsEntry[i]})
    }
    dot1qPortVlanHCStatisticsTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qPortVlanHCStatisticsTable.EntityData.YListKeys = []string {}

    return &(dot1qPortVlanHCStatisticsTable.EntityData)
}

// QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable_Dot1qPortVlanHCStatisticsEntry
// Traffic statistics for a VLAN on a high-capacity
// interface.
type QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable_Dot1qPortVlanHCStatisticsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1dBasePort
    Dot1dBasePort interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1qVlanCurrentTable_Dot1qVlanCurrentEntry_Dot1qVlanIndex
    Dot1qVlanIndex interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN.  Note that a frame received on this
    // port is counted by this object if and only if it is for a  protocol being
    // processed by the local forwarding process for this VLAN.  This object
    // includes received bridge management frames classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP). The type is interface{} with range:
    // 0..18446744073709551615.
    Dot1qTpVlanPortHCInFrames interface{}

    // The number of valid frames transmitted by this port to its segment from the
    // local forwarding process for this VLAN.  This includes bridge management
    // frames originated by this device that are classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP). The type is interface{} with range:
    // 0..18446744073709551615.
    Dot1qTpVlanPortHCOutFrames interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN and that were discarded due to
    // VLAN-related reasons. Specifically, the IEEE 802.1Q counters for Discard
    // Inbound and Discard on Ingress Filtering. The type is interface{} with
    // range: 0..18446744073709551615.
    Dot1qTpVlanPortHCInDiscards interface{}
}

func (dot1qPortVlanHCStatisticsEntry *QBRIDGEMIB_Dot1qPortVlanHCStatisticsTable_Dot1qPortVlanHCStatisticsEntry) GetEntityData() *types.CommonEntityData {
    dot1qPortVlanHCStatisticsEntry.EntityData.YFilter = dot1qPortVlanHCStatisticsEntry.YFilter
    dot1qPortVlanHCStatisticsEntry.EntityData.YangName = "dot1qPortVlanHCStatisticsEntry"
    dot1qPortVlanHCStatisticsEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qPortVlanHCStatisticsEntry.EntityData.ParentYangName = "dot1qPortVlanHCStatisticsTable"
    dot1qPortVlanHCStatisticsEntry.EntityData.SegmentPath = "dot1qPortVlanHCStatisticsEntry" + types.AddKeyToken(dot1qPortVlanHCStatisticsEntry.Dot1dBasePort, "dot1dBasePort") + types.AddKeyToken(dot1qPortVlanHCStatisticsEntry.Dot1qVlanIndex, "dot1qVlanIndex")
    dot1qPortVlanHCStatisticsEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qPortVlanHCStatisticsTable/" + dot1qPortVlanHCStatisticsEntry.EntityData.SegmentPath
    dot1qPortVlanHCStatisticsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qPortVlanHCStatisticsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qPortVlanHCStatisticsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qPortVlanHCStatisticsEntry.EntityData.Children = types.NewOrderedMap()
    dot1qPortVlanHCStatisticsEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qPortVlanHCStatisticsEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", dot1qPortVlanHCStatisticsEntry.Dot1dBasePort})
    dot1qPortVlanHCStatisticsEntry.EntityData.Leafs.Append("dot1qVlanIndex", types.YLeaf{"Dot1qVlanIndex", dot1qPortVlanHCStatisticsEntry.Dot1qVlanIndex})
    dot1qPortVlanHCStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortHCInFrames", types.YLeaf{"Dot1qTpVlanPortHCInFrames", dot1qPortVlanHCStatisticsEntry.Dot1qTpVlanPortHCInFrames})
    dot1qPortVlanHCStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortHCOutFrames", types.YLeaf{"Dot1qTpVlanPortHCOutFrames", dot1qPortVlanHCStatisticsEntry.Dot1qTpVlanPortHCOutFrames})
    dot1qPortVlanHCStatisticsEntry.EntityData.Leafs.Append("dot1qTpVlanPortHCInDiscards", types.YLeaf{"Dot1qTpVlanPortHCInDiscards", dot1qPortVlanHCStatisticsEntry.Dot1qTpVlanPortHCInDiscards})

    dot1qPortVlanHCStatisticsEntry.EntityData.YListKeys = []string {"Dot1dBasePort", "Dot1qVlanIndex"}

    return &(dot1qPortVlanHCStatisticsEntry.EntityData)
}

// QBRIDGEMIB_Dot1qLearningConstraintsTable
// A table containing learning constraints for sets of
// Shared and Independent VLANs.
type QBRIDGEMIB_Dot1qLearningConstraintsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A learning constraint defined for a VLAN. The type is slice of
    // QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry.
    Dot1qLearningConstraintsEntry []*QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry
}

func (dot1qLearningConstraintsTable *QBRIDGEMIB_Dot1qLearningConstraintsTable) GetEntityData() *types.CommonEntityData {
    dot1qLearningConstraintsTable.EntityData.YFilter = dot1qLearningConstraintsTable.YFilter
    dot1qLearningConstraintsTable.EntityData.YangName = "dot1qLearningConstraintsTable"
    dot1qLearningConstraintsTable.EntityData.BundleName = "cisco_ios_xe"
    dot1qLearningConstraintsTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1qLearningConstraintsTable.EntityData.SegmentPath = "dot1qLearningConstraintsTable"
    dot1qLearningConstraintsTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1qLearningConstraintsTable.EntityData.SegmentPath
    dot1qLearningConstraintsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qLearningConstraintsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qLearningConstraintsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qLearningConstraintsTable.EntityData.Children = types.NewOrderedMap()
    dot1qLearningConstraintsTable.EntityData.Children.Append("dot1qLearningConstraintsEntry", types.YChild{"Dot1qLearningConstraintsEntry", nil})
    for i := range dot1qLearningConstraintsTable.Dot1qLearningConstraintsEntry {
        dot1qLearningConstraintsTable.EntityData.Children.Append(types.GetSegmentPath(dot1qLearningConstraintsTable.Dot1qLearningConstraintsEntry[i]), types.YChild{"Dot1qLearningConstraintsEntry", dot1qLearningConstraintsTable.Dot1qLearningConstraintsEntry[i]})
    }
    dot1qLearningConstraintsTable.EntityData.Leafs = types.NewOrderedMap()

    dot1qLearningConstraintsTable.EntityData.YListKeys = []string {}

    return &(dot1qLearningConstraintsTable.EntityData)
}

// QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry
// A learning constraint defined for a VLAN.
type QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index of the row in dot1qVlanCurrentTable for
    // the VLAN constrained by this entry. The type is interface{} with range:
    // 0..4294967295.
    Dot1qConstraintVlan interface{}

    // This attribute is a key. The identity of the constraint set to which
    // dot1qConstraintVlan belongs.  These values may be chosen by the management
    // station. The type is interface{} with range: 0..65535.
    Dot1qConstraintSet interface{}

    // The type of constraint this entry defines. independent(1) - the VLAN,
    // dot1qConstraintVlan,     uses a filtering database independent from all    
    // other VLANs in the same set, defined by     dot1qConstraintSet. shared(2) -
    // the VLAN, dot1qConstraintVlan, shares     the same filtering database as
    // all other VLANs     in the same set, defined by dot1qConstraintSet. The
    // type is Dot1qConstraintType.
    Dot1qConstraintType interface{}

    // The status of this entry. The type is RowStatus.
    Dot1qConstraintStatus interface{}
}

func (dot1qLearningConstraintsEntry *QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry) GetEntityData() *types.CommonEntityData {
    dot1qLearningConstraintsEntry.EntityData.YFilter = dot1qLearningConstraintsEntry.YFilter
    dot1qLearningConstraintsEntry.EntityData.YangName = "dot1qLearningConstraintsEntry"
    dot1qLearningConstraintsEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1qLearningConstraintsEntry.EntityData.ParentYangName = "dot1qLearningConstraintsTable"
    dot1qLearningConstraintsEntry.EntityData.SegmentPath = "dot1qLearningConstraintsEntry" + types.AddKeyToken(dot1qLearningConstraintsEntry.Dot1qConstraintVlan, "dot1qConstraintVlan") + types.AddKeyToken(dot1qLearningConstraintsEntry.Dot1qConstraintSet, "dot1qConstraintSet")
    dot1qLearningConstraintsEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1qLearningConstraintsTable/" + dot1qLearningConstraintsEntry.EntityData.SegmentPath
    dot1qLearningConstraintsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1qLearningConstraintsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1qLearningConstraintsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1qLearningConstraintsEntry.EntityData.Children = types.NewOrderedMap()
    dot1qLearningConstraintsEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1qLearningConstraintsEntry.EntityData.Leafs.Append("dot1qConstraintVlan", types.YLeaf{"Dot1qConstraintVlan", dot1qLearningConstraintsEntry.Dot1qConstraintVlan})
    dot1qLearningConstraintsEntry.EntityData.Leafs.Append("dot1qConstraintSet", types.YLeaf{"Dot1qConstraintSet", dot1qLearningConstraintsEntry.Dot1qConstraintSet})
    dot1qLearningConstraintsEntry.EntityData.Leafs.Append("dot1qConstraintType", types.YLeaf{"Dot1qConstraintType", dot1qLearningConstraintsEntry.Dot1qConstraintType})
    dot1qLearningConstraintsEntry.EntityData.Leafs.Append("dot1qConstraintStatus", types.YLeaf{"Dot1qConstraintStatus", dot1qLearningConstraintsEntry.Dot1qConstraintStatus})

    dot1qLearningConstraintsEntry.EntityData.YListKeys = []string {"Dot1qConstraintVlan", "Dot1qConstraintSet"}

    return &(dot1qLearningConstraintsEntry.EntityData)
}

// QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry_Dot1qConstraintType represents     in the same set, defined by dot1qConstraintSet.
type QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry_Dot1qConstraintType string

const (
    QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry_Dot1qConstraintType_independent QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry_Dot1qConstraintType = "independent"

    QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry_Dot1qConstraintType_shared QBRIDGEMIB_Dot1qLearningConstraintsTable_Dot1qLearningConstraintsEntry_Dot1qConstraintType = "shared"
)

// QBRIDGEMIB_Dot1vProtocolGroupTable
// A table that contains mappings from Protocol
// Templates to Protocol Group Identifiers used for
// Port-and-Protocol-based VLAN Classification.
type QBRIDGEMIB_Dot1vProtocolGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A mapping from a Protocol Template to a Protocol Group Identifier. The type
    // is slice of QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry.
    Dot1vProtocolGroupEntry []*QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry
}

func (dot1vProtocolGroupTable *QBRIDGEMIB_Dot1vProtocolGroupTable) GetEntityData() *types.CommonEntityData {
    dot1vProtocolGroupTable.EntityData.YFilter = dot1vProtocolGroupTable.YFilter
    dot1vProtocolGroupTable.EntityData.YangName = "dot1vProtocolGroupTable"
    dot1vProtocolGroupTable.EntityData.BundleName = "cisco_ios_xe"
    dot1vProtocolGroupTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1vProtocolGroupTable.EntityData.SegmentPath = "dot1vProtocolGroupTable"
    dot1vProtocolGroupTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1vProtocolGroupTable.EntityData.SegmentPath
    dot1vProtocolGroupTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1vProtocolGroupTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1vProtocolGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1vProtocolGroupTable.EntityData.Children = types.NewOrderedMap()
    dot1vProtocolGroupTable.EntityData.Children.Append("dot1vProtocolGroupEntry", types.YChild{"Dot1vProtocolGroupEntry", nil})
    for i := range dot1vProtocolGroupTable.Dot1vProtocolGroupEntry {
        dot1vProtocolGroupTable.EntityData.Children.Append(types.GetSegmentPath(dot1vProtocolGroupTable.Dot1vProtocolGroupEntry[i]), types.YChild{"Dot1vProtocolGroupEntry", dot1vProtocolGroupTable.Dot1vProtocolGroupEntry[i]})
    }
    dot1vProtocolGroupTable.EntityData.Leafs = types.NewOrderedMap()

    dot1vProtocolGroupTable.EntityData.YListKeys = []string {}

    return &(dot1vProtocolGroupTable.EntityData)
}

// QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry
// A mapping from a Protocol Template to a Protocol
// Group Identifier.
type QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The data-link encapsulation format or the
    // 'detagged_frame_type' in a Protocol Template. The type is
    // Dot1vProtocolTemplateFrameType.
    Dot1vProtocolTemplateFrameType interface{}

    // This attribute is a key. The identification of the protocol above the
    // data-link layer in a Protocol Template.  Depending on the frame type, the
    // octet string will have one of the following values:  For 'ethernet',
    // 'rfc1042' and 'snap8021H',     this is the 16-bit (2-octet) IEEE 802.3 Type
    // Field. For 'snapOther',     this is the 40-bit (5-octet) PID. For
    // 'llcOther',     this is the 2-octet IEEE 802.2 Link Service Access    
    // Point (LSAP) pair: first octet for Destination Service     Access Point
    // (DSAP) and second octet for Source Service     Access Point (SSAP). The
    // type is string with length: 2 | 5.
    Dot1vProtocolTemplateProtocolValue interface{}

    // Represents a group of protocols that are associated together when assigning
    // a VID to a frame. The type is interface{} with range: 0..2147483647.
    Dot1vProtocolGroupId interface{}

    // This object indicates the status of this entry. The type is RowStatus.
    Dot1vProtocolGroupRowStatus interface{}
}

func (dot1vProtocolGroupEntry *QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry) GetEntityData() *types.CommonEntityData {
    dot1vProtocolGroupEntry.EntityData.YFilter = dot1vProtocolGroupEntry.YFilter
    dot1vProtocolGroupEntry.EntityData.YangName = "dot1vProtocolGroupEntry"
    dot1vProtocolGroupEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1vProtocolGroupEntry.EntityData.ParentYangName = "dot1vProtocolGroupTable"
    dot1vProtocolGroupEntry.EntityData.SegmentPath = "dot1vProtocolGroupEntry" + types.AddKeyToken(dot1vProtocolGroupEntry.Dot1vProtocolTemplateFrameType, "dot1vProtocolTemplateFrameType") + types.AddKeyToken(dot1vProtocolGroupEntry.Dot1vProtocolTemplateProtocolValue, "dot1vProtocolTemplateProtocolValue")
    dot1vProtocolGroupEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1vProtocolGroupTable/" + dot1vProtocolGroupEntry.EntityData.SegmentPath
    dot1vProtocolGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1vProtocolGroupEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1vProtocolGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1vProtocolGroupEntry.EntityData.Children = types.NewOrderedMap()
    dot1vProtocolGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1vProtocolGroupEntry.EntityData.Leafs.Append("dot1vProtocolTemplateFrameType", types.YLeaf{"Dot1vProtocolTemplateFrameType", dot1vProtocolGroupEntry.Dot1vProtocolTemplateFrameType})
    dot1vProtocolGroupEntry.EntityData.Leafs.Append("dot1vProtocolTemplateProtocolValue", types.YLeaf{"Dot1vProtocolTemplateProtocolValue", dot1vProtocolGroupEntry.Dot1vProtocolTemplateProtocolValue})
    dot1vProtocolGroupEntry.EntityData.Leafs.Append("dot1vProtocolGroupId", types.YLeaf{"Dot1vProtocolGroupId", dot1vProtocolGroupEntry.Dot1vProtocolGroupId})
    dot1vProtocolGroupEntry.EntityData.Leafs.Append("dot1vProtocolGroupRowStatus", types.YLeaf{"Dot1vProtocolGroupRowStatus", dot1vProtocolGroupEntry.Dot1vProtocolGroupRowStatus})

    dot1vProtocolGroupEntry.EntityData.YListKeys = []string {"Dot1vProtocolTemplateFrameType", "Dot1vProtocolTemplateProtocolValue"}

    return &(dot1vProtocolGroupEntry.EntityData)
}

// QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType represents 'detagged_frame_type' in a Protocol Template.
type QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType string

const (
    QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType_ethernet QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType = "ethernet"

    QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType_rfc1042 QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType = "rfc1042"

    QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType_snap8021H QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType = "snap8021H"

    QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType_snapOther QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType = "snapOther"

    QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType_llcOther QBRIDGEMIB_Dot1vProtocolGroupTable_Dot1vProtocolGroupEntry_Dot1vProtocolTemplateFrameType = "llcOther"
)

// QBRIDGEMIB_Dot1vProtocolPortTable
// A table that contains VID sets used for
// Port-and-Protocol-based VLAN Classification.
type QBRIDGEMIB_Dot1vProtocolPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VID set for a port. The type is slice of
    // QBRIDGEMIB_Dot1vProtocolPortTable_Dot1vProtocolPortEntry.
    Dot1vProtocolPortEntry []*QBRIDGEMIB_Dot1vProtocolPortTable_Dot1vProtocolPortEntry
}

func (dot1vProtocolPortTable *QBRIDGEMIB_Dot1vProtocolPortTable) GetEntityData() *types.CommonEntityData {
    dot1vProtocolPortTable.EntityData.YFilter = dot1vProtocolPortTable.YFilter
    dot1vProtocolPortTable.EntityData.YangName = "dot1vProtocolPortTable"
    dot1vProtocolPortTable.EntityData.BundleName = "cisco_ios_xe"
    dot1vProtocolPortTable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1vProtocolPortTable.EntityData.SegmentPath = "dot1vProtocolPortTable"
    dot1vProtocolPortTable.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/" + dot1vProtocolPortTable.EntityData.SegmentPath
    dot1vProtocolPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1vProtocolPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1vProtocolPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1vProtocolPortTable.EntityData.Children = types.NewOrderedMap()
    dot1vProtocolPortTable.EntityData.Children.Append("dot1vProtocolPortEntry", types.YChild{"Dot1vProtocolPortEntry", nil})
    for i := range dot1vProtocolPortTable.Dot1vProtocolPortEntry {
        dot1vProtocolPortTable.EntityData.Children.Append(types.GetSegmentPath(dot1vProtocolPortTable.Dot1vProtocolPortEntry[i]), types.YChild{"Dot1vProtocolPortEntry", dot1vProtocolPortTable.Dot1vProtocolPortEntry[i]})
    }
    dot1vProtocolPortTable.EntityData.Leafs = types.NewOrderedMap()

    dot1vProtocolPortTable.EntityData.YListKeys = []string {}

    return &(dot1vProtocolPortTable.EntityData)
}

// QBRIDGEMIB_Dot1vProtocolPortTable_Dot1vProtocolPortEntry
// A VID set for a port.
type QBRIDGEMIB_Dot1vProtocolPortTable_Dot1vProtocolPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1dBasePortTable_Dot1dBasePortEntry_Dot1dBasePort
    Dot1dBasePort interface{}

    // This attribute is a key. Designates a group of protocols in the Protocol
    // Group Database. The type is interface{} with range: 1..2147483647.
    Dot1vProtocolPortGroupId interface{}

    // The VID associated with a group of protocols for each port. The type is
    // interface{} with range: 1..4094.
    Dot1vProtocolPortGroupVid interface{}

    // This object indicates the status of this entry. The type is RowStatus.
    Dot1vProtocolPortRowStatus interface{}
}

func (dot1vProtocolPortEntry *QBRIDGEMIB_Dot1vProtocolPortTable_Dot1vProtocolPortEntry) GetEntityData() *types.CommonEntityData {
    dot1vProtocolPortEntry.EntityData.YFilter = dot1vProtocolPortEntry.YFilter
    dot1vProtocolPortEntry.EntityData.YangName = "dot1vProtocolPortEntry"
    dot1vProtocolPortEntry.EntityData.BundleName = "cisco_ios_xe"
    dot1vProtocolPortEntry.EntityData.ParentYangName = "dot1vProtocolPortTable"
    dot1vProtocolPortEntry.EntityData.SegmentPath = "dot1vProtocolPortEntry" + types.AddKeyToken(dot1vProtocolPortEntry.Dot1dBasePort, "dot1dBasePort") + types.AddKeyToken(dot1vProtocolPortEntry.Dot1vProtocolPortGroupId, "dot1vProtocolPortGroupId")
    dot1vProtocolPortEntry.EntityData.AbsolutePath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB/dot1vProtocolPortTable/" + dot1vProtocolPortEntry.EntityData.SegmentPath
    dot1vProtocolPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1vProtocolPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1vProtocolPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1vProtocolPortEntry.EntityData.Children = types.NewOrderedMap()
    dot1vProtocolPortEntry.EntityData.Leafs = types.NewOrderedMap()
    dot1vProtocolPortEntry.EntityData.Leafs.Append("dot1dBasePort", types.YLeaf{"Dot1dBasePort", dot1vProtocolPortEntry.Dot1dBasePort})
    dot1vProtocolPortEntry.EntityData.Leafs.Append("dot1vProtocolPortGroupId", types.YLeaf{"Dot1vProtocolPortGroupId", dot1vProtocolPortEntry.Dot1vProtocolPortGroupId})
    dot1vProtocolPortEntry.EntityData.Leafs.Append("dot1vProtocolPortGroupVid", types.YLeaf{"Dot1vProtocolPortGroupVid", dot1vProtocolPortEntry.Dot1vProtocolPortGroupVid})
    dot1vProtocolPortEntry.EntityData.Leafs.Append("dot1vProtocolPortRowStatus", types.YLeaf{"Dot1vProtocolPortRowStatus", dot1vProtocolPortEntry.Dot1vProtocolPortRowStatus})

    dot1vProtocolPortEntry.EntityData.YListKeys = []string {"Dot1dBasePort", "Dot1vProtocolPortGroupId"}

    return &(dot1vProtocolPortEntry.EntityData)
}

