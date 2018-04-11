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

    
    Dot1Qbase QBRIDGEMIB_Dot1Qbase

    
    Dot1Qvlan QBRIDGEMIB_Dot1Qvlan

    // A table that contains configuration and control information for each
    // Filtering Database currently operating on this device.  Entries in this
    // table appear automatically when VLANs are assigned FDB IDs in the
    // dot1qVlanCurrentTable.
    Dot1Qfdbtable QBRIDGEMIB_Dot1Qfdbtable

    // A table that contains information about unicast entries for which the
    // device has forwarding and/or filtering information.  This information is
    // used by the transparent bridging function in determining how to propagate a
    // received frame.
    Dot1Qtpfdbtable QBRIDGEMIB_Dot1Qtpfdbtable

    // A table containing filtering information for VLANs configured into the
    // bridge by (local or network) management, or learned dynamically, specifying
    // the set of ports to which frames received on a VLAN for this FDB and
    // containing a specific Group destination address are allowed to be
    // forwarded.
    Dot1Qtpgrouptable QBRIDGEMIB_Dot1Qtpgrouptable

    // A table containing forwarding information for each  VLAN, specifying the
    // set of ports to which forwarding of all multicasts applies, configured
    // statically by management or dynamically by GMRP.  An entry appears in this
    // table for all VLANs that are currently instantiated.
    Dot1Qforwardalltable QBRIDGEMIB_Dot1Qforwardalltable

    // A table containing forwarding information for each VLAN, specifying the set
    // of ports to which forwarding of multicast group-addressed frames for which
    // no more specific forwarding information applies.  This is configured
    // statically by management and determined dynamically by GMRP.  An entry
    // appears in this table for all VLANs that are currently instantiated.
    Dot1Qforwardunregisteredtable QBRIDGEMIB_Dot1Qforwardunregisteredtable

    // A table containing filtering information for Unicast MAC addresses for each
    // Filtering Database, configured into the device by (local or network)
    // management specifying the set of ports to which frames received from
    // specific ports and containing specific unicast destination addresses are
    // allowed to be forwarded.  A value of zero in this table (as the port number
    // from  which frames with a specific destination address are received) is
    // used to specify all ports for which there is no specific entry in this
    // table for that particular destination address.  Entries are valid for
    // unicast addresses only.
    Dot1Qstaticunicasttable QBRIDGEMIB_Dot1Qstaticunicasttable

    // A table containing filtering information for Multicast and Broadcast MAC
    // addresses for each VLAN, configured into the device by (local or network)
    // management specifying the set of ports to which frames received from
    // specific ports and containing specific Multicast and Broadcast destination
    // addresses are allowed to be forwarded.  A value of zero in this table (as
    // the port number from which frames with a specific destination address are
    // received) is used to specify all ports for which there is no specific entry
    // in this table for that particular destination address.  Entries are valid
    // for Multicast and Broadcast addresses only.
    Dot1Qstaticmulticasttable QBRIDGEMIB_Dot1Qstaticmulticasttable

    // A table containing current configuration information for each VLAN
    // currently configured into the device by (local or network) management, or
    // dynamically created as a result of GVRP requests received.
    Dot1Qvlancurrenttable QBRIDGEMIB_Dot1Qvlancurrenttable

    // A table containing static configuration information for each VLAN
    // configured into the device by (local or network) management.  All entries
    // are permanent and will be restored after the device is reset.
    Dot1Qvlanstatictable QBRIDGEMIB_Dot1Qvlanstatictable

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
    Dot1Qportvlanstatisticstable QBRIDGEMIB_Dot1Qportvlanstatisticstable

    // A table containing per-port, per-VLAN statistics for traffic on
    // high-capacity interfaces.
    Dot1Qportvlanhcstatisticstable QBRIDGEMIB_Dot1Qportvlanhcstatisticstable

    // A table containing learning constraints for sets of Shared and Independent
    // VLANs.
    Dot1Qlearningconstraintstable QBRIDGEMIB_Dot1Qlearningconstraintstable

    // A table that contains mappings from Protocol Templates to Protocol Group
    // Identifiers used for Port-and-Protocol-based VLAN Classification.
    Dot1Vprotocolgrouptable QBRIDGEMIB_Dot1Vprotocolgrouptable

    // A table that contains VID sets used for Port-and-Protocol-based VLAN
    // Classification.
    Dot1Vprotocolporttable QBRIDGEMIB_Dot1Vprotocolporttable
}

func (qBRIDGEMIB *QBRIDGEMIB) GetEntityData() *types.CommonEntityData {
    qBRIDGEMIB.EntityData.YFilter = qBRIDGEMIB.YFilter
    qBRIDGEMIB.EntityData.YangName = "Q-BRIDGE-MIB"
    qBRIDGEMIB.EntityData.BundleName = "cisco_ios_xe"
    qBRIDGEMIB.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    qBRIDGEMIB.EntityData.SegmentPath = "Q-BRIDGE-MIB:Q-BRIDGE-MIB"
    qBRIDGEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qBRIDGEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qBRIDGEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qBRIDGEMIB.EntityData.Children = make(map[string]types.YChild)
    qBRIDGEMIB.EntityData.Children["dot1qBase"] = types.YChild{"Dot1Qbase", &qBRIDGEMIB.Dot1Qbase}
    qBRIDGEMIB.EntityData.Children["dot1qVlan"] = types.YChild{"Dot1Qvlan", &qBRIDGEMIB.Dot1Qvlan}
    qBRIDGEMIB.EntityData.Children["dot1qFdbTable"] = types.YChild{"Dot1Qfdbtable", &qBRIDGEMIB.Dot1Qfdbtable}
    qBRIDGEMIB.EntityData.Children["dot1qTpFdbTable"] = types.YChild{"Dot1Qtpfdbtable", &qBRIDGEMIB.Dot1Qtpfdbtable}
    qBRIDGEMIB.EntityData.Children["dot1qTpGroupTable"] = types.YChild{"Dot1Qtpgrouptable", &qBRIDGEMIB.Dot1Qtpgrouptable}
    qBRIDGEMIB.EntityData.Children["dot1qForwardAllTable"] = types.YChild{"Dot1Qforwardalltable", &qBRIDGEMIB.Dot1Qforwardalltable}
    qBRIDGEMIB.EntityData.Children["dot1qForwardUnregisteredTable"] = types.YChild{"Dot1Qforwardunregisteredtable", &qBRIDGEMIB.Dot1Qforwardunregisteredtable}
    qBRIDGEMIB.EntityData.Children["dot1qStaticUnicastTable"] = types.YChild{"Dot1Qstaticunicasttable", &qBRIDGEMIB.Dot1Qstaticunicasttable}
    qBRIDGEMIB.EntityData.Children["dot1qStaticMulticastTable"] = types.YChild{"Dot1Qstaticmulticasttable", &qBRIDGEMIB.Dot1Qstaticmulticasttable}
    qBRIDGEMIB.EntityData.Children["dot1qVlanCurrentTable"] = types.YChild{"Dot1Qvlancurrenttable", &qBRIDGEMIB.Dot1Qvlancurrenttable}
    qBRIDGEMIB.EntityData.Children["dot1qVlanStaticTable"] = types.YChild{"Dot1Qvlanstatictable", &qBRIDGEMIB.Dot1Qvlanstatictable}
    qBRIDGEMIB.EntityData.Children["dot1qPortVlanStatisticsTable"] = types.YChild{"Dot1Qportvlanstatisticstable", &qBRIDGEMIB.Dot1Qportvlanstatisticstable}
    qBRIDGEMIB.EntityData.Children["dot1qPortVlanHCStatisticsTable"] = types.YChild{"Dot1Qportvlanhcstatisticstable", &qBRIDGEMIB.Dot1Qportvlanhcstatisticstable}
    qBRIDGEMIB.EntityData.Children["dot1qLearningConstraintsTable"] = types.YChild{"Dot1Qlearningconstraintstable", &qBRIDGEMIB.Dot1Qlearningconstraintstable}
    qBRIDGEMIB.EntityData.Children["dot1vProtocolGroupTable"] = types.YChild{"Dot1Vprotocolgrouptable", &qBRIDGEMIB.Dot1Vprotocolgrouptable}
    qBRIDGEMIB.EntityData.Children["dot1vProtocolPortTable"] = types.YChild{"Dot1Vprotocolporttable", &qBRIDGEMIB.Dot1Vprotocolporttable}
    qBRIDGEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qBRIDGEMIB.EntityData)
}

// QBRIDGEMIB_Dot1Qbase
type QBRIDGEMIB_Dot1Qbase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The version number of IEEE 802.1Q that this device supports. The type is
    // Dot1Qvlanversionnumber.
    Dot1Qvlanversionnumber interface{}

    // The maximum IEEE 802.1Q VLAN-ID that this device  supports. The type is
    // interface{} with range: 1..4094.
    Dot1Qmaxvlanid interface{}

    // The maximum number of IEEE 802.1Q VLANs that this device supports. The type
    // is interface{} with range: 0..4294967295.
    Dot1Qmaxsupportedvlans interface{}

    // The current number of IEEE 802.1Q VLANs that are configured in this device.
    // The type is interface{} with range: 0..4294967295.
    Dot1Qnumvlans interface{}

    // The administrative status requested by management for GVRP.  The value
    // enabled(1) indicates that GVRP should be enabled on this device, on all
    // ports for which it has not been specifically disabled.  When disabled(2),
    // GVRP is disabled on all ports, and all GVRP packets will be forwarded
    // transparently.  This object affects all GVRP Applicant and Registrar state
    // machines.  A transition from disabled(2) to enabled(1) will cause a reset
    // of all GVRP state machines on all ports.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // EnabledStatus.
    Dot1Qgvrpstatus interface{}
}

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetEntityData() *types.CommonEntityData {
    dot1Qbase.EntityData.YFilter = dot1Qbase.YFilter
    dot1Qbase.EntityData.YangName = "dot1qBase"
    dot1Qbase.EntityData.BundleName = "cisco_ios_xe"
    dot1Qbase.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qbase.EntityData.SegmentPath = "dot1qBase"
    dot1Qbase.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qbase.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qbase.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qbase.EntityData.Children = make(map[string]types.YChild)
    dot1Qbase.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qbase.EntityData.Leafs["dot1qVlanVersionNumber"] = types.YLeaf{"Dot1Qvlanversionnumber", dot1Qbase.Dot1Qvlanversionnumber}
    dot1Qbase.EntityData.Leafs["dot1qMaxVlanId"] = types.YLeaf{"Dot1Qmaxvlanid", dot1Qbase.Dot1Qmaxvlanid}
    dot1Qbase.EntityData.Leafs["dot1qMaxSupportedVlans"] = types.YLeaf{"Dot1Qmaxsupportedvlans", dot1Qbase.Dot1Qmaxsupportedvlans}
    dot1Qbase.EntityData.Leafs["dot1qNumVlans"] = types.YLeaf{"Dot1Qnumvlans", dot1Qbase.Dot1Qnumvlans}
    dot1Qbase.EntityData.Leafs["dot1qGvrpStatus"] = types.YLeaf{"Dot1Qgvrpstatus", dot1Qbase.Dot1Qgvrpstatus}
    return &(dot1Qbase.EntityData)
}

// QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber represents supports.
type QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber string

const (
    QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber_version1 QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber = "version1"
)

// QBRIDGEMIB_Dot1Qvlan
type QBRIDGEMIB_Dot1Qvlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times a VLAN entry has been deleted from the
    // dot1qVlanCurrentTable (for any reason).  If an entry is deleted, then
    // inserted, and then deleted, this counter will be incremented by 2. The type
    // is interface{} with range: 0..4294967295.
    Dot1Qvlannumdeletes interface{}

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
    Dot1Qnextfreelocalvlanindex interface{}

    // The identity of the constraint set to which a VLAN belongs, if there is not
    // an explicit entry for that VLAN in dot1qLearningConstraintsTable.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is interface{} with range: 0..65535.
    Dot1Qconstraintsetdefault interface{}

    // The type of constraint set to which a VLAN belongs, if there is not an
    // explicit entry for that VLAN in dot1qLearningConstraintsTable.  The types
    // are as defined for dot1qConstraintType.  The value of this object MUST be
    // retained across  reinitializations of the management system. The type is
    // Dot1Qconstrainttypedefault.
    Dot1Qconstrainttypedefault interface{}
}

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetEntityData() *types.CommonEntityData {
    dot1Qvlan.EntityData.YFilter = dot1Qvlan.YFilter
    dot1Qvlan.EntityData.YangName = "dot1qVlan"
    dot1Qvlan.EntityData.BundleName = "cisco_ios_xe"
    dot1Qvlan.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qvlan.EntityData.SegmentPath = "dot1qVlan"
    dot1Qvlan.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qvlan.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qvlan.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qvlan.EntityData.Children = make(map[string]types.YChild)
    dot1Qvlan.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qvlan.EntityData.Leafs["dot1qVlanNumDeletes"] = types.YLeaf{"Dot1Qvlannumdeletes", dot1Qvlan.Dot1Qvlannumdeletes}
    dot1Qvlan.EntityData.Leafs["dot1qNextFreeLocalVlanIndex"] = types.YLeaf{"Dot1Qnextfreelocalvlanindex", dot1Qvlan.Dot1Qnextfreelocalvlanindex}
    dot1Qvlan.EntityData.Leafs["dot1qConstraintSetDefault"] = types.YLeaf{"Dot1Qconstraintsetdefault", dot1Qvlan.Dot1Qconstraintsetdefault}
    dot1Qvlan.EntityData.Leafs["dot1qConstraintTypeDefault"] = types.YLeaf{"Dot1Qconstrainttypedefault", dot1Qvlan.Dot1Qconstrainttypedefault}
    return &(dot1Qvlan.EntityData)
}

// QBRIDGEMIB_Dot1Qvlan_Dot1Qconstrainttypedefault represents reinitializations of the management system.
type QBRIDGEMIB_Dot1Qvlan_Dot1Qconstrainttypedefault string

const (
    QBRIDGEMIB_Dot1Qvlan_Dot1Qconstrainttypedefault_independent QBRIDGEMIB_Dot1Qvlan_Dot1Qconstrainttypedefault = "independent"

    QBRIDGEMIB_Dot1Qvlan_Dot1Qconstrainttypedefault_shared QBRIDGEMIB_Dot1Qvlan_Dot1Qconstrainttypedefault = "shared"
)

// QBRIDGEMIB_Dot1Qfdbtable
// A table that contains configuration and control
// information for each Filtering Database currently
// operating on this device.  Entries in this table appear
// automatically when VLANs are assigned FDB IDs in the
// dot1qVlanCurrentTable.
type QBRIDGEMIB_Dot1Qfdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a specific Filtering Database. The type is slice of
    // QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry.
    Dot1Qfdbentry []QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry
}

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetEntityData() *types.CommonEntityData {
    dot1Qfdbtable.EntityData.YFilter = dot1Qfdbtable.YFilter
    dot1Qfdbtable.EntityData.YangName = "dot1qFdbTable"
    dot1Qfdbtable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qfdbtable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qfdbtable.EntityData.SegmentPath = "dot1qFdbTable"
    dot1Qfdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qfdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qfdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qfdbtable.EntityData.Children = make(map[string]types.YChild)
    dot1Qfdbtable.EntityData.Children["dot1qFdbEntry"] = types.YChild{"Dot1Qfdbentry", nil}
    for i := range dot1Qfdbtable.Dot1Qfdbentry {
        dot1Qfdbtable.EntityData.Children[types.GetSegmentPath(&dot1Qfdbtable.Dot1Qfdbentry[i])] = types.YChild{"Dot1Qfdbentry", &dot1Qfdbtable.Dot1Qfdbentry[i]}
    }
    dot1Qfdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qfdbtable.EntityData)
}

// QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry
// Information about a specific Filtering Database.
type QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identity of this Filtering Database. The type
    // is interface{} with range: 0..4294967295.
    Dot1Qfdbid interface{}

    // The current number of dynamic entries in this Filtering Database. The type
    // is interface{} with range: 0..4294967295.
    Dot1Qfdbdynamiccount interface{}
}

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetEntityData() *types.CommonEntityData {
    dot1Qfdbentry.EntityData.YFilter = dot1Qfdbentry.YFilter
    dot1Qfdbentry.EntityData.YangName = "dot1qFdbEntry"
    dot1Qfdbentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qfdbentry.EntityData.ParentYangName = "dot1qFdbTable"
    dot1Qfdbentry.EntityData.SegmentPath = "dot1qFdbEntry" + "[dot1qFdbId='" + fmt.Sprintf("%v", dot1Qfdbentry.Dot1Qfdbid) + "']"
    dot1Qfdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qfdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qfdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qfdbentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qfdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qfdbentry.EntityData.Leafs["dot1qFdbId"] = types.YLeaf{"Dot1Qfdbid", dot1Qfdbentry.Dot1Qfdbid}
    dot1Qfdbentry.EntityData.Leafs["dot1qFdbDynamicCount"] = types.YLeaf{"Dot1Qfdbdynamiccount", dot1Qfdbentry.Dot1Qfdbdynamiccount}
    return &(dot1Qfdbentry.EntityData)
}

// QBRIDGEMIB_Dot1Qtpfdbtable
// A table that contains information about unicast entries
// for which the device has forwarding and/or filtering
// information.  This information is used by the
// transparent bridging function in determining how to
// propagate a received frame.
type QBRIDGEMIB_Dot1Qtpfdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a specific unicast MAC address for which the device has
    // some forwarding and/or filtering information. The type is slice of
    // QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry.
    Dot1Qtpfdbentry []QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry
}

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetEntityData() *types.CommonEntityData {
    dot1Qtpfdbtable.EntityData.YFilter = dot1Qtpfdbtable.YFilter
    dot1Qtpfdbtable.EntityData.YangName = "dot1qTpFdbTable"
    dot1Qtpfdbtable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qtpfdbtable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qtpfdbtable.EntityData.SegmentPath = "dot1qTpFdbTable"
    dot1Qtpfdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qtpfdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qtpfdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qtpfdbtable.EntityData.Children = make(map[string]types.YChild)
    dot1Qtpfdbtable.EntityData.Children["dot1qTpFdbEntry"] = types.YChild{"Dot1Qtpfdbentry", nil}
    for i := range dot1Qtpfdbtable.Dot1Qtpfdbentry {
        dot1Qtpfdbtable.EntityData.Children[types.GetSegmentPath(&dot1Qtpfdbtable.Dot1Qtpfdbentry[i])] = types.YChild{"Dot1Qtpfdbentry", &dot1Qtpfdbtable.Dot1Qtpfdbentry[i]}
    }
    dot1Qtpfdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qtpfdbtable.EntityData)
}

// QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry
// Information about a specific unicast MAC address for
// which the device has some forwarding and/or filtering
// information.
type QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to q_bridge_mib.QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry_Dot1Qfdbid
    Dot1Qfdbid interface{}

    // This attribute is a key. A unicast MAC address for which the device has
    // forwarding and/or filtering information. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Qtpfdbaddress interface{}

    // Either the value '0', or the port number of the port on which a frame
    // having a source address equal to the value of the corresponding instance of
    // dot1qTpFdbAddress has been seen.  A value of '0' indicates that the port
    // number has not been learned but that the device does have some
    // forwarding/filtering information about this address (e.g., in the
    // dot1qStaticUnicastTable). Implementors are encouraged to assign the port
    // value to this object whenever it is learned, even for addresses for which
    // the corresponding value of dot1qTpFdbStatus is not learned(3). The type is
    // interface{} with range: 0..65535.
    Dot1Qtpfdbport interface{}

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
    // dot1qStaticAddress. The type is Dot1Qtpfdbstatus.
    Dot1Qtpfdbstatus interface{}
}

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetEntityData() *types.CommonEntityData {
    dot1Qtpfdbentry.EntityData.YFilter = dot1Qtpfdbentry.YFilter
    dot1Qtpfdbentry.EntityData.YangName = "dot1qTpFdbEntry"
    dot1Qtpfdbentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qtpfdbentry.EntityData.ParentYangName = "dot1qTpFdbTable"
    dot1Qtpfdbentry.EntityData.SegmentPath = "dot1qTpFdbEntry" + "[dot1qFdbId='" + fmt.Sprintf("%v", dot1Qtpfdbentry.Dot1Qfdbid) + "']" + "[dot1qTpFdbAddress='" + fmt.Sprintf("%v", dot1Qtpfdbentry.Dot1Qtpfdbaddress) + "']"
    dot1Qtpfdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qtpfdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qtpfdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qtpfdbentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qtpfdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qtpfdbentry.EntityData.Leafs["dot1qFdbId"] = types.YLeaf{"Dot1Qfdbid", dot1Qtpfdbentry.Dot1Qfdbid}
    dot1Qtpfdbentry.EntityData.Leafs["dot1qTpFdbAddress"] = types.YLeaf{"Dot1Qtpfdbaddress", dot1Qtpfdbentry.Dot1Qtpfdbaddress}
    dot1Qtpfdbentry.EntityData.Leafs["dot1qTpFdbPort"] = types.YLeaf{"Dot1Qtpfdbport", dot1Qtpfdbentry.Dot1Qtpfdbport}
    dot1Qtpfdbentry.EntityData.Leafs["dot1qTpFdbStatus"] = types.YLeaf{"Dot1Qtpfdbstatus", dot1Qtpfdbentry.Dot1Qtpfdbstatus}
    return &(dot1Qtpfdbentry.EntityData)
}

// QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus represents         existing instance of dot1qStaticAddress.
type QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus string

const (
    QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus_other QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus = "other"

    QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus_invalid QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus = "invalid"

    QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus_learned QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus = "learned"

    QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus_self QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus = "self"

    QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus_mgmt QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry_Dot1Qtpfdbstatus = "mgmt"
)

// QBRIDGEMIB_Dot1Qtpgrouptable
// A table containing filtering information for VLANs
// configured into the bridge by (local or network)
// management, or learned dynamically, specifying the set of
// ports to which frames received on a VLAN for this FDB
// and containing a specific Group destination address are
// allowed to be forwarded.
type QBRIDGEMIB_Dot1Qtpgrouptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the bridge by management, or learned
    // dynamically, specifying the set of ports to which frames received on a VLAN
    // and containing a specific Group destination address are allowed to be
    // forwarded.  The subset of these ports learned dynamically is also provided.
    // The type is slice of QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry.
    Dot1Qtpgroupentry []QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry
}

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetEntityData() *types.CommonEntityData {
    dot1Qtpgrouptable.EntityData.YFilter = dot1Qtpgrouptable.YFilter
    dot1Qtpgrouptable.EntityData.YangName = "dot1qTpGroupTable"
    dot1Qtpgrouptable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qtpgrouptable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qtpgrouptable.EntityData.SegmentPath = "dot1qTpGroupTable"
    dot1Qtpgrouptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qtpgrouptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qtpgrouptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qtpgrouptable.EntityData.Children = make(map[string]types.YChild)
    dot1Qtpgrouptable.EntityData.Children["dot1qTpGroupEntry"] = types.YChild{"Dot1Qtpgroupentry", nil}
    for i := range dot1Qtpgrouptable.Dot1Qtpgroupentry {
        dot1Qtpgrouptable.EntityData.Children[types.GetSegmentPath(&dot1Qtpgrouptable.Dot1Qtpgroupentry[i])] = types.YChild{"Dot1Qtpgroupentry", &dot1Qtpgrouptable.Dot1Qtpgroupentry[i]}
    }
    dot1Qtpgrouptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qtpgrouptable.EntityData)
}

// QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry
// Filtering information configured into the bridge by
// management, or learned dynamically, specifying the set of
// ports to which frames received on a VLAN and containing
// a specific Group destination address are allowed to be
// forwarded.  The subset of these ports learned dynamically
// is also provided.
type QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // This attribute is a key. The destination Group MAC address in a frame to
    // which this entry's filtering information applies. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Qtpgroupaddress interface{}

    // The complete set of ports, in this VLAN, to which frames destined for this
    // Group MAC address are currently being explicitly forwarded.  This does not
    // include ports for which this address is only implicitly forwarded, in the
    // dot1qForwardAllPorts list. The type is string.
    Dot1Qtpgroupegressports interface{}

    // The subset of ports in dot1qTpGroupEgressPorts that were learned by GMRP or
    // some other dynamic mechanism, in this Filtering database. The type is
    // string.
    Dot1Qtpgrouplearnt interface{}
}

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetEntityData() *types.CommonEntityData {
    dot1Qtpgroupentry.EntityData.YFilter = dot1Qtpgroupentry.YFilter
    dot1Qtpgroupentry.EntityData.YangName = "dot1qTpGroupEntry"
    dot1Qtpgroupentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qtpgroupentry.EntityData.ParentYangName = "dot1qTpGroupTable"
    dot1Qtpgroupentry.EntityData.SegmentPath = "dot1qTpGroupEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qtpgroupentry.Dot1Qvlanindex) + "']" + "[dot1qTpGroupAddress='" + fmt.Sprintf("%v", dot1Qtpgroupentry.Dot1Qtpgroupaddress) + "']"
    dot1Qtpgroupentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qtpgroupentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qtpgroupentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qtpgroupentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qtpgroupentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qtpgroupentry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qtpgroupentry.Dot1Qvlanindex}
    dot1Qtpgroupentry.EntityData.Leafs["dot1qTpGroupAddress"] = types.YLeaf{"Dot1Qtpgroupaddress", dot1Qtpgroupentry.Dot1Qtpgroupaddress}
    dot1Qtpgroupentry.EntityData.Leafs["dot1qTpGroupEgressPorts"] = types.YLeaf{"Dot1Qtpgroupegressports", dot1Qtpgroupentry.Dot1Qtpgroupegressports}
    dot1Qtpgroupentry.EntityData.Leafs["dot1qTpGroupLearnt"] = types.YLeaf{"Dot1Qtpgrouplearnt", dot1Qtpgroupentry.Dot1Qtpgrouplearnt}
    return &(dot1Qtpgroupentry.EntityData)
}

// QBRIDGEMIB_Dot1Qforwardalltable
// A table containing forwarding information for each
// 
// VLAN, specifying the set of ports to which forwarding of
// all multicasts applies, configured statically by
// management or dynamically by GMRP.  An entry appears in
// this table for all VLANs that are currently
// instantiated.
type QBRIDGEMIB_Dot1Qforwardalltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding information for a VLAN, specifying the set of ports to which all
    // multicasts should be forwarded, configured statically by management or
    // dynamically by GMRP. The type is slice of
    // QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry.
    Dot1Qforwardallentry []QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry
}

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetEntityData() *types.CommonEntityData {
    dot1Qforwardalltable.EntityData.YFilter = dot1Qforwardalltable.YFilter
    dot1Qforwardalltable.EntityData.YangName = "dot1qForwardAllTable"
    dot1Qforwardalltable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qforwardalltable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qforwardalltable.EntityData.SegmentPath = "dot1qForwardAllTable"
    dot1Qforwardalltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qforwardalltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qforwardalltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qforwardalltable.EntityData.Children = make(map[string]types.YChild)
    dot1Qforwardalltable.EntityData.Children["dot1qForwardAllEntry"] = types.YChild{"Dot1Qforwardallentry", nil}
    for i := range dot1Qforwardalltable.Dot1Qforwardallentry {
        dot1Qforwardalltable.EntityData.Children[types.GetSegmentPath(&dot1Qforwardalltable.Dot1Qforwardallentry[i])] = types.YChild{"Dot1Qforwardallentry", &dot1Qforwardalltable.Dot1Qforwardallentry[i]}
    }
    dot1Qforwardalltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qforwardalltable.EntityData)
}

// QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry
// Forwarding information for a VLAN, specifying the set
// of ports to which all multicasts should be forwarded,
// configured statically by management or dynamically by
// GMRP.
type QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // The complete set of ports in this VLAN to which all multicast
    // group-addressed frames are to be forwarded. This includes ports for which
    // this need has been determined dynamically by GMRP, or configured statically
    // by management. The type is string.
    Dot1Qforwardallports interface{}

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
    Dot1Qforwardallstaticports interface{}

    // The set of ports configured by management in this VLAN for which the
    // Service Requirement attribute Forward All Multicast Groups may not be
    // dynamically registered by GMRP.  This value will be restored after the
    // device is reset.  A port may not be added in this set if it is already a
    // member of the set of ports in dot1qForwardAllStaticPorts.  The default
    // value is a string of zeros of appropriate length.  The value of this object
    // MUST be retained across reinitializations of the management system. The
    // type is string.
    Dot1Qforwardallforbiddenports interface{}
}

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetEntityData() *types.CommonEntityData {
    dot1Qforwardallentry.EntityData.YFilter = dot1Qforwardallentry.YFilter
    dot1Qforwardallentry.EntityData.YangName = "dot1qForwardAllEntry"
    dot1Qforwardallentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qforwardallentry.EntityData.ParentYangName = "dot1qForwardAllTable"
    dot1Qforwardallentry.EntityData.SegmentPath = "dot1qForwardAllEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qforwardallentry.Dot1Qvlanindex) + "']"
    dot1Qforwardallentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qforwardallentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qforwardallentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qforwardallentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qforwardallentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qforwardallentry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qforwardallentry.Dot1Qvlanindex}
    dot1Qforwardallentry.EntityData.Leafs["dot1qForwardAllPorts"] = types.YLeaf{"Dot1Qforwardallports", dot1Qforwardallentry.Dot1Qforwardallports}
    dot1Qforwardallentry.EntityData.Leafs["dot1qForwardAllStaticPorts"] = types.YLeaf{"Dot1Qforwardallstaticports", dot1Qforwardallentry.Dot1Qforwardallstaticports}
    dot1Qforwardallentry.EntityData.Leafs["dot1qForwardAllForbiddenPorts"] = types.YLeaf{"Dot1Qforwardallforbiddenports", dot1Qforwardallentry.Dot1Qforwardallforbiddenports}
    return &(dot1Qforwardallentry.EntityData)
}

// QBRIDGEMIB_Dot1Qforwardunregisteredtable
// A table containing forwarding information for each
// VLAN, specifying the set of ports to which forwarding of
// multicast group-addressed frames for which no
// more specific forwarding information applies.  This is
// configured statically by management and determined
// dynamically by GMRP.  An entry appears in this table for
// all VLANs that are currently instantiated.
type QBRIDGEMIB_Dot1Qforwardunregisteredtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding information for a VLAN, specifying the set of ports to which all
    // multicasts for which there is no more specific forwarding information shall
    // be forwarded. This is configured statically by management or dynamically by
    // GMRP. The type is slice of
    // QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry.
    Dot1Qforwardunregisteredentry []QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry
}

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetEntityData() *types.CommonEntityData {
    dot1Qforwardunregisteredtable.EntityData.YFilter = dot1Qforwardunregisteredtable.YFilter
    dot1Qforwardunregisteredtable.EntityData.YangName = "dot1qForwardUnregisteredTable"
    dot1Qforwardunregisteredtable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qforwardunregisteredtable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qforwardunregisteredtable.EntityData.SegmentPath = "dot1qForwardUnregisteredTable"
    dot1Qforwardunregisteredtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qforwardunregisteredtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qforwardunregisteredtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qforwardunregisteredtable.EntityData.Children = make(map[string]types.YChild)
    dot1Qforwardunregisteredtable.EntityData.Children["dot1qForwardUnregisteredEntry"] = types.YChild{"Dot1Qforwardunregisteredentry", nil}
    for i := range dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry {
        dot1Qforwardunregisteredtable.EntityData.Children[types.GetSegmentPath(&dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry[i])] = types.YChild{"Dot1Qforwardunregisteredentry", &dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry[i]}
    }
    dot1Qforwardunregisteredtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qforwardunregisteredtable.EntityData)
}

// QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry
// Forwarding information for a VLAN, specifying the set
// of ports to which all multicasts for which there is no
// more specific forwarding information shall be forwarded.
// This is configured statically by management or
// dynamically by GMRP.
type QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // The complete set of ports in this VLAN to which multicast group-addressed
    // frames for which there is no more specific forwarding information will be
    // forwarded. This includes ports for which this need has been determined
    // dynamically by GMRP, or configured statically by management. The type is
    // string.
    Dot1Qforwardunregisteredports interface{}

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
    Dot1Qforwardunregisteredstaticports interface{}

    // The set of ports configured by management in this VLAN for which the
    // Service Requirement attribute Forward Unregistered Multicast Groups may not
    // be dynamically registered by GMRP.  This value will be restored after the
    // device is reset.  A port may not be added in this set if it is already a
    // member of the set of ports in dot1qForwardUnregisteredStaticPorts.  The
    // default value is a string of zeros of appropriate length.  The value of
    // this object MUST be retained across reinitializations of the management
    // system. The type is string.
    Dot1Qforwardunregisteredforbiddenports interface{}
}

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetEntityData() *types.CommonEntityData {
    dot1Qforwardunregisteredentry.EntityData.YFilter = dot1Qforwardunregisteredentry.YFilter
    dot1Qforwardunregisteredentry.EntityData.YangName = "dot1qForwardUnregisteredEntry"
    dot1Qforwardunregisteredentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qforwardunregisteredentry.EntityData.ParentYangName = "dot1qForwardUnregisteredTable"
    dot1Qforwardunregisteredentry.EntityData.SegmentPath = "dot1qForwardUnregisteredEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qforwardunregisteredentry.Dot1Qvlanindex) + "']"
    dot1Qforwardunregisteredentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qforwardunregisteredentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qforwardunregisteredentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qforwardunregisteredentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qforwardunregisteredentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qforwardunregisteredentry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qforwardunregisteredentry.Dot1Qvlanindex}
    dot1Qforwardunregisteredentry.EntityData.Leafs["dot1qForwardUnregisteredPorts"] = types.YLeaf{"Dot1Qforwardunregisteredports", dot1Qforwardunregisteredentry.Dot1Qforwardunregisteredports}
    dot1Qforwardunregisteredentry.EntityData.Leafs["dot1qForwardUnregisteredStaticPorts"] = types.YLeaf{"Dot1Qforwardunregisteredstaticports", dot1Qforwardunregisteredentry.Dot1Qforwardunregisteredstaticports}
    dot1Qforwardunregisteredentry.EntityData.Leafs["dot1qForwardUnregisteredForbiddenPorts"] = types.YLeaf{"Dot1Qforwardunregisteredforbiddenports", dot1Qforwardunregisteredentry.Dot1Qforwardunregisteredforbiddenports}
    return &(dot1Qforwardunregisteredentry.EntityData)
}

// QBRIDGEMIB_Dot1Qstaticunicasttable
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
type QBRIDGEMIB_Dot1Qstaticunicasttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the device by (local or network)
    // management specifying the set of ports to which frames received from a
    // specific port and containing a specific unicast destination address are
    // allowed to be forwarded. The type is slice of
    // QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry.
    Dot1Qstaticunicastentry []QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry
}

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetEntityData() *types.CommonEntityData {
    dot1Qstaticunicasttable.EntityData.YFilter = dot1Qstaticunicasttable.YFilter
    dot1Qstaticunicasttable.EntityData.YangName = "dot1qStaticUnicastTable"
    dot1Qstaticunicasttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qstaticunicasttable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qstaticunicasttable.EntityData.SegmentPath = "dot1qStaticUnicastTable"
    dot1Qstaticunicasttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qstaticunicasttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qstaticunicasttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qstaticunicasttable.EntityData.Children = make(map[string]types.YChild)
    dot1Qstaticunicasttable.EntityData.Children["dot1qStaticUnicastEntry"] = types.YChild{"Dot1Qstaticunicastentry", nil}
    for i := range dot1Qstaticunicasttable.Dot1Qstaticunicastentry {
        dot1Qstaticunicasttable.EntityData.Children[types.GetSegmentPath(&dot1Qstaticunicasttable.Dot1Qstaticunicastentry[i])] = types.YChild{"Dot1Qstaticunicastentry", &dot1Qstaticunicasttable.Dot1Qstaticunicastentry[i]}
    }
    dot1Qstaticunicasttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qstaticunicasttable.EntityData)
}

// QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry
// Filtering information configured into the device by
// (local or network) management specifying the set of
// ports to which frames received from a specific port and
// containing a specific unicast destination address are
// allowed to be forwarded.
type QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to q_bridge_mib.QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry_Dot1Qfdbid
    Dot1Qfdbid interface{}

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object must take the
    // value of a unicast address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Qstaticunicastaddress interface{}

    // This attribute is a key. Either the value '0' or the port number of the
    // port from which a frame must be received in order for this entry's
    // filtering information to apply.  A value of zero indicates that this entry
    // applies on all ports of the device for which there is no other applicable
    // entry. The type is interface{} with range: 0..65535.
    Dot1Qstaticunicastreceiveport interface{}

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
    Dot1Qstaticunicastallowedtogoto interface{}

    // This object indicates the status of this entry. other(1) - this entry is
    // currently in use, but      the conditions under which it will remain     so
    // differ from the following values. invalid(2) - writing this value to the
    // object     removes the corresponding entry. permanent(3) - this entry is
    // currently in use     and will remain so after the next reset of     the
    // bridge. deleteOnReset(4) - this entry is currently in     use and will
    // remain so until the next     reset of the bridge. deleteOnTimeout(5) - this
    // entry is currently in     use and will remain so until it is aged out.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is Dot1Qstaticunicaststatus.
    Dot1Qstaticunicaststatus interface{}
}

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetEntityData() *types.CommonEntityData {
    dot1Qstaticunicastentry.EntityData.YFilter = dot1Qstaticunicastentry.YFilter
    dot1Qstaticunicastentry.EntityData.YangName = "dot1qStaticUnicastEntry"
    dot1Qstaticunicastentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qstaticunicastentry.EntityData.ParentYangName = "dot1qStaticUnicastTable"
    dot1Qstaticunicastentry.EntityData.SegmentPath = "dot1qStaticUnicastEntry" + "[dot1qFdbId='" + fmt.Sprintf("%v", dot1Qstaticunicastentry.Dot1Qfdbid) + "']" + "[dot1qStaticUnicastAddress='" + fmt.Sprintf("%v", dot1Qstaticunicastentry.Dot1Qstaticunicastaddress) + "']" + "[dot1qStaticUnicastReceivePort='" + fmt.Sprintf("%v", dot1Qstaticunicastentry.Dot1Qstaticunicastreceiveport) + "']"
    dot1Qstaticunicastentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qstaticunicastentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qstaticunicastentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qstaticunicastentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qstaticunicastentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qstaticunicastentry.EntityData.Leafs["dot1qFdbId"] = types.YLeaf{"Dot1Qfdbid", dot1Qstaticunicastentry.Dot1Qfdbid}
    dot1Qstaticunicastentry.EntityData.Leafs["dot1qStaticUnicastAddress"] = types.YLeaf{"Dot1Qstaticunicastaddress", dot1Qstaticunicastentry.Dot1Qstaticunicastaddress}
    dot1Qstaticunicastentry.EntityData.Leafs["dot1qStaticUnicastReceivePort"] = types.YLeaf{"Dot1Qstaticunicastreceiveport", dot1Qstaticunicastentry.Dot1Qstaticunicastreceiveport}
    dot1Qstaticunicastentry.EntityData.Leafs["dot1qStaticUnicastAllowedToGoTo"] = types.YLeaf{"Dot1Qstaticunicastallowedtogoto", dot1Qstaticunicastentry.Dot1Qstaticunicastallowedtogoto}
    dot1Qstaticunicastentry.EntityData.Leafs["dot1qStaticUnicastStatus"] = types.YLeaf{"Dot1Qstaticunicaststatus", dot1Qstaticunicastentry.Dot1Qstaticunicaststatus}
    return &(dot1Qstaticunicastentry.EntityData)
}

// QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus represents reinitializations of the management system.
type QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus string

const (
    QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus_other QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus = "other"

    QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus_invalid QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus = "invalid"

    QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus_permanent QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus = "permanent"

    QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus_deleteOnReset QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus = "deleteOnReset"

    QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus_deleteOnTimeout QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry_Dot1Qstaticunicaststatus = "deleteOnTimeout"
)

// QBRIDGEMIB_Dot1Qstaticmulticasttable
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
type QBRIDGEMIB_Dot1Qstaticmulticasttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering information configured into the device by (local or network)
    // management specifying the set of ports to which frames received from this
    // specific port  for this VLAN and containing this Multicast or Broadcast
    // destination address are allowed to be forwarded. The type is slice of
    // QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry.
    Dot1Qstaticmulticastentry []QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry
}

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetEntityData() *types.CommonEntityData {
    dot1Qstaticmulticasttable.EntityData.YFilter = dot1Qstaticmulticasttable.YFilter
    dot1Qstaticmulticasttable.EntityData.YangName = "dot1qStaticMulticastTable"
    dot1Qstaticmulticasttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qstaticmulticasttable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qstaticmulticasttable.EntityData.SegmentPath = "dot1qStaticMulticastTable"
    dot1Qstaticmulticasttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qstaticmulticasttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qstaticmulticasttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qstaticmulticasttable.EntityData.Children = make(map[string]types.YChild)
    dot1Qstaticmulticasttable.EntityData.Children["dot1qStaticMulticastEntry"] = types.YChild{"Dot1Qstaticmulticastentry", nil}
    for i := range dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry {
        dot1Qstaticmulticasttable.EntityData.Children[types.GetSegmentPath(&dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry[i])] = types.YChild{"Dot1Qstaticmulticastentry", &dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry[i]}
    }
    dot1Qstaticmulticasttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qstaticmulticasttable.EntityData)
}

// QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry
// Filtering information configured into the device by
// (local or network) management specifying the set of
// ports to which frames received from this specific port
// 
// for this VLAN and containing this Multicast or Broadcast
// destination address are allowed to be forwarded.
type QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object must take the
    // value of a Multicast or Broadcast address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Dot1Qstaticmulticastaddress interface{}

    // This attribute is a key. Either the value '0' or the port number of the
    // port from which a frame must be received in order for this entry's
    // filtering information to apply.  A value of zero indicates that this entry
    // applies on all ports of the device for which there is no other applicable
    // entry. The type is interface{} with range: 0..65535.
    Dot1Qstaticmulticastreceiveport interface{}

    // The set of ports to which frames received from a specific port and destined
    // for a specific Multicast or Broadcast MAC address must be forwarded,
    // regardless of any dynamic information, e.g., from GMRP.  A port may not be
    // added in this set if it is already a member of the set of ports in
    // dot1qStaticMulticastForbiddenEgressPorts. The default value of this object
    // is a string of ones of appropriate length.  The value of this object MUST
    // be retained across reinitializations of the management system. The type is
    // string.
    Dot1Qstaticmulticaststaticegressports interface{}

    // The set of ports to which frames received from a specific port and destined
    // for a specific Multicast or Broadcast MAC address must not be forwarded,
    // regardless of any dynamic information, e.g., from GMRP.  A port may not be
    // added in this set if it is already a member of the set of ports in
    // dot1qStaticMulticastStaticEgressPorts. The default value of this object is
    // a string of zeros of appropriate length.  The value of this object MUST be
    // retained across reinitializations of the management system. The type is
    // string.
    Dot1Qstaticmulticastforbiddenegressports interface{}

    // This object indicates the status of this entry. other(1) - this entry is
    // currently in use, but     the conditions under which it will remain     so
    // differ from the following values.  invalid(2) - writing this value to the
    // object     removes the corresponding entry. permanent(3) - this entry is
    // currently in use     and will remain so after the next reset of     the
    // bridge. deleteOnReset(4) - this entry is currently in     use and will
    // remain so until the next     reset of the bridge. deleteOnTimeout(5) - this
    // entry is currently in     use and will remain so until it is aged out.  The
    // value of this object MUST be retained across reinitializations of the
    // management system. The type is Dot1Qstaticmulticaststatus.
    Dot1Qstaticmulticaststatus interface{}
}

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetEntityData() *types.CommonEntityData {
    dot1Qstaticmulticastentry.EntityData.YFilter = dot1Qstaticmulticastentry.YFilter
    dot1Qstaticmulticastentry.EntityData.YangName = "dot1qStaticMulticastEntry"
    dot1Qstaticmulticastentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qstaticmulticastentry.EntityData.ParentYangName = "dot1qStaticMulticastTable"
    dot1Qstaticmulticastentry.EntityData.SegmentPath = "dot1qStaticMulticastEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qstaticmulticastentry.Dot1Qvlanindex) + "']" + "[dot1qStaticMulticastAddress='" + fmt.Sprintf("%v", dot1Qstaticmulticastentry.Dot1Qstaticmulticastaddress) + "']" + "[dot1qStaticMulticastReceivePort='" + fmt.Sprintf("%v", dot1Qstaticmulticastentry.Dot1Qstaticmulticastreceiveport) + "']"
    dot1Qstaticmulticastentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qstaticmulticastentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qstaticmulticastentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qstaticmulticastentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qstaticmulticastentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qstaticmulticastentry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qstaticmulticastentry.Dot1Qvlanindex}
    dot1Qstaticmulticastentry.EntityData.Leafs["dot1qStaticMulticastAddress"] = types.YLeaf{"Dot1Qstaticmulticastaddress", dot1Qstaticmulticastentry.Dot1Qstaticmulticastaddress}
    dot1Qstaticmulticastentry.EntityData.Leafs["dot1qStaticMulticastReceivePort"] = types.YLeaf{"Dot1Qstaticmulticastreceiveport", dot1Qstaticmulticastentry.Dot1Qstaticmulticastreceiveport}
    dot1Qstaticmulticastentry.EntityData.Leafs["dot1qStaticMulticastStaticEgressPorts"] = types.YLeaf{"Dot1Qstaticmulticaststaticegressports", dot1Qstaticmulticastentry.Dot1Qstaticmulticaststaticegressports}
    dot1Qstaticmulticastentry.EntityData.Leafs["dot1qStaticMulticastForbiddenEgressPorts"] = types.YLeaf{"Dot1Qstaticmulticastforbiddenegressports", dot1Qstaticmulticastentry.Dot1Qstaticmulticastforbiddenegressports}
    dot1Qstaticmulticastentry.EntityData.Leafs["dot1qStaticMulticastStatus"] = types.YLeaf{"Dot1Qstaticmulticaststatus", dot1Qstaticmulticastentry.Dot1Qstaticmulticaststatus}
    return &(dot1Qstaticmulticastentry.EntityData)
}

// QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus represents reinitializations of the management system.
type QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus string

const (
    QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus_other QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus = "other"

    QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus_invalid QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus = "invalid"

    QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus_permanent QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus = "permanent"

    QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus_deleteOnReset QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus = "deleteOnReset"

    QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus_deleteOnTimeout QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry_Dot1Qstaticmulticaststatus = "deleteOnTimeout"
)

// QBRIDGEMIB_Dot1Qvlancurrenttable
// A table containing current configuration information
// for each VLAN currently configured into the device by
// (local or network) management, or dynamically created
// as a result of GVRP requests received.
type QBRIDGEMIB_Dot1Qvlancurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information for a VLAN configured into the device by  (local or network)
    // management, or dynamically created as a result of GVRP requests received.
    // The type is slice of
    // QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry.
    Dot1Qvlancurrententry []QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry
}

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetEntityData() *types.CommonEntityData {
    dot1Qvlancurrenttable.EntityData.YFilter = dot1Qvlancurrenttable.YFilter
    dot1Qvlancurrenttable.EntityData.YangName = "dot1qVlanCurrentTable"
    dot1Qvlancurrenttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qvlancurrenttable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qvlancurrenttable.EntityData.SegmentPath = "dot1qVlanCurrentTable"
    dot1Qvlancurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qvlancurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qvlancurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qvlancurrenttable.EntityData.Children = make(map[string]types.YChild)
    dot1Qvlancurrenttable.EntityData.Children["dot1qVlanCurrentEntry"] = types.YChild{"Dot1Qvlancurrententry", nil}
    for i := range dot1Qvlancurrenttable.Dot1Qvlancurrententry {
        dot1Qvlancurrenttable.EntityData.Children[types.GetSegmentPath(&dot1Qvlancurrenttable.Dot1Qvlancurrententry[i])] = types.YChild{"Dot1Qvlancurrententry", &dot1Qvlancurrenttable.Dot1Qvlancurrententry[i]}
    }
    dot1Qvlancurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qvlancurrenttable.EntityData)
}

// QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry
// Information for a VLAN configured into the device by
// 
// (local or network) management, or dynamically created
// as a result of GVRP requests received.
type QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention to see how this works. The type is interface{} with
    // range: 0..4294967295.
    Dot1Qvlantimemark interface{}

    // This attribute is a key. The VLAN-ID or other identifier referring to this
    // VLAN. The type is interface{} with range: 0..4294967295.
    Dot1Qvlanindex interface{}

    // The Filtering Database used by this VLAN.  This is one of the dot1qFdbId
    // values in the dot1qFdbTable.  This value is allocated automatically by the
    // device whenever  the VLAN is created: either dynamically by GVRP, or by
    // management, in dot1qVlanStaticTable.  Allocation of this value follows the
    // learning constraints defined for this VLAN in
    // dot1qLearningConstraintsTable. The type is interface{} with range:
    // 0..4294967295.
    Dot1Qvlanfdbid interface{}

    // The set of ports that are transmitting traffic for this VLAN as either
    // tagged or untagged frames. The type is string.
    Dot1Qvlancurrentegressports interface{}

    // The set of ports that are transmitting traffic for this VLAN as untagged
    // frames. The type is string.
    Dot1Qvlancurrentuntaggedports interface{}

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
    // Dot1Qvlanstatus.
    Dot1Qvlanstatus interface{}

    // The value of sysUpTime when this VLAN was created. The type is interface{}
    // with range: 0..4294967295.
    Dot1Qvlancreationtime interface{}
}

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetEntityData() *types.CommonEntityData {
    dot1Qvlancurrententry.EntityData.YFilter = dot1Qvlancurrententry.YFilter
    dot1Qvlancurrententry.EntityData.YangName = "dot1qVlanCurrentEntry"
    dot1Qvlancurrententry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qvlancurrententry.EntityData.ParentYangName = "dot1qVlanCurrentTable"
    dot1Qvlancurrententry.EntityData.SegmentPath = "dot1qVlanCurrentEntry" + "[dot1qVlanTimeMark='" + fmt.Sprintf("%v", dot1Qvlancurrententry.Dot1Qvlantimemark) + "']" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qvlancurrententry.Dot1Qvlanindex) + "']"
    dot1Qvlancurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qvlancurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qvlancurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qvlancurrententry.EntityData.Children = make(map[string]types.YChild)
    dot1Qvlancurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qvlancurrententry.EntityData.Leafs["dot1qVlanTimeMark"] = types.YLeaf{"Dot1Qvlantimemark", dot1Qvlancurrententry.Dot1Qvlantimemark}
    dot1Qvlancurrententry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qvlancurrententry.Dot1Qvlanindex}
    dot1Qvlancurrententry.EntityData.Leafs["dot1qVlanFdbId"] = types.YLeaf{"Dot1Qvlanfdbid", dot1Qvlancurrententry.Dot1Qvlanfdbid}
    dot1Qvlancurrententry.EntityData.Leafs["dot1qVlanCurrentEgressPorts"] = types.YLeaf{"Dot1Qvlancurrentegressports", dot1Qvlancurrententry.Dot1Qvlancurrentegressports}
    dot1Qvlancurrententry.EntityData.Leafs["dot1qVlanCurrentUntaggedPorts"] = types.YLeaf{"Dot1Qvlancurrentuntaggedports", dot1Qvlancurrententry.Dot1Qvlancurrentuntaggedports}
    dot1Qvlancurrententry.EntityData.Leafs["dot1qVlanStatus"] = types.YLeaf{"Dot1Qvlanstatus", dot1Qvlancurrententry.Dot1Qvlanstatus}
    dot1Qvlancurrententry.EntityData.Leafs["dot1qVlanCreationTime"] = types.YLeaf{"Dot1Qvlancreationtime", dot1Qvlancurrententry.Dot1Qvlancreationtime}
    return &(dot1Qvlancurrententry.EntityData)
}

// QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus represents     removed when the last port leaves the VLAN.
type QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus string

const (
    QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus_other QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus = "other"

    QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus_permanent QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus = "permanent"

    QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus_dynamicGvrp QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanstatus = "dynamicGvrp"
)

// QBRIDGEMIB_Dot1Qvlanstatictable
// A table containing static configuration information for
// each VLAN configured into the device by (local or
// network) management.  All entries are permanent and will
// be restored after the device is reset.
type QBRIDGEMIB_Dot1Qvlanstatictable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static information for a VLAN configured into the device by (local or
    // network) management. The type is slice of
    // QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry.
    Dot1Qvlanstaticentry []QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry
}

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetEntityData() *types.CommonEntityData {
    dot1Qvlanstatictable.EntityData.YFilter = dot1Qvlanstatictable.YFilter
    dot1Qvlanstatictable.EntityData.YangName = "dot1qVlanStaticTable"
    dot1Qvlanstatictable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qvlanstatictable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qvlanstatictable.EntityData.SegmentPath = "dot1qVlanStaticTable"
    dot1Qvlanstatictable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qvlanstatictable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qvlanstatictable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qvlanstatictable.EntityData.Children = make(map[string]types.YChild)
    dot1Qvlanstatictable.EntityData.Children["dot1qVlanStaticEntry"] = types.YChild{"Dot1Qvlanstaticentry", nil}
    for i := range dot1Qvlanstatictable.Dot1Qvlanstaticentry {
        dot1Qvlanstatictable.EntityData.Children[types.GetSegmentPath(&dot1Qvlanstatictable.Dot1Qvlanstaticentry[i])] = types.YChild{"Dot1Qvlanstaticentry", &dot1Qvlanstatictable.Dot1Qvlanstaticentry[i]}
    }
    dot1Qvlanstatictable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qvlanstatictable.EntityData)
}

// QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry
// Static information for a VLAN configured into the
// device by (local or network) management.
type QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // An administratively assigned string, which may be used to identify the
    // VLAN. The type is string with length: 0..32.
    Dot1Qvlanstaticname interface{}

    // The set of ports that are permanently assigned to the egress list for this
    // VLAN by management.  Changes to a bit in this object affect the per-port,
    // per-VLAN Registrar control for Registration Fixed for the relevant GVRP
    // state machine on each port.  A port may not be added in this set if it is
    // already a member of the set of ports in dot1qVlanForbiddenEgressPorts.  The
    // default value of this object is a string of zeros of appropriate length,
    // indicating not fixed. The type is string.
    Dot1Qvlanstaticegressports interface{}

    // The set of ports that are prohibited by management from being included in
    // the egress list for this VLAN. Changes to this object that cause a port to
    // be included or excluded affect the per-port, per-VLAN Registrar control for
    // Registration Forbidden for the relevant GVRP state machine on each port.  A
    // port may not be added in this set if it is already a member of the set of
    // ports in dot1qVlanStaticEgressPorts.  The default value of this object is a
    // string of zeros of appropriate length, excluding all ports from the
    // forbidden set. The type is string.
    Dot1Qvlanforbiddenegressports interface{}

    // The set of ports that should transmit egress packets for this VLAN as
    // untagged.  The default value of this object for the default VLAN
    // (dot1qVlanIndex = 1) is a string of appropriate length including all ports.
    // There is no specified default for other VLANs.  If a device agent cannot
    // support the set of ports being set, then it will reject the set operation
    // with an error.  For example, a manager might attempt to set more than one
    // VLAN to be untagged on egress where the device does not support this IEEE
    // 802.1Q option. The type is string.
    Dot1Qvlanstaticuntaggedports interface{}

    // This object indicates the status of this entry. The type is RowStatus.
    Dot1Qvlanstaticrowstatus interface{}
}

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetEntityData() *types.CommonEntityData {
    dot1Qvlanstaticentry.EntityData.YFilter = dot1Qvlanstaticentry.YFilter
    dot1Qvlanstaticentry.EntityData.YangName = "dot1qVlanStaticEntry"
    dot1Qvlanstaticentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qvlanstaticentry.EntityData.ParentYangName = "dot1qVlanStaticTable"
    dot1Qvlanstaticentry.EntityData.SegmentPath = "dot1qVlanStaticEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qvlanstaticentry.Dot1Qvlanindex) + "']"
    dot1Qvlanstaticentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qvlanstaticentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qvlanstaticentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qvlanstaticentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qvlanstaticentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qvlanstaticentry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qvlanstaticentry.Dot1Qvlanindex}
    dot1Qvlanstaticentry.EntityData.Leafs["dot1qVlanStaticName"] = types.YLeaf{"Dot1Qvlanstaticname", dot1Qvlanstaticentry.Dot1Qvlanstaticname}
    dot1Qvlanstaticentry.EntityData.Leafs["dot1qVlanStaticEgressPorts"] = types.YLeaf{"Dot1Qvlanstaticegressports", dot1Qvlanstaticentry.Dot1Qvlanstaticegressports}
    dot1Qvlanstaticentry.EntityData.Leafs["dot1qVlanForbiddenEgressPorts"] = types.YLeaf{"Dot1Qvlanforbiddenegressports", dot1Qvlanstaticentry.Dot1Qvlanforbiddenegressports}
    dot1Qvlanstaticentry.EntityData.Leafs["dot1qVlanStaticUntaggedPorts"] = types.YLeaf{"Dot1Qvlanstaticuntaggedports", dot1Qvlanstaticentry.Dot1Qvlanstaticuntaggedports}
    dot1Qvlanstaticentry.EntityData.Leafs["dot1qVlanStaticRowStatus"] = types.YLeaf{"Dot1Qvlanstaticrowstatus", dot1Qvlanstaticentry.Dot1Qvlanstaticrowstatus}
    return &(dot1Qvlanstaticentry.EntityData)
}

// QBRIDGEMIB_Dot1Qportvlanstatisticstable
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
type QBRIDGEMIB_Dot1Qportvlanstatisticstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic statistics for a VLAN on an interface. The type is slice of
    // QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry.
    Dot1Qportvlanstatisticsentry []QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry
}

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetEntityData() *types.CommonEntityData {
    dot1Qportvlanstatisticstable.EntityData.YFilter = dot1Qportvlanstatisticstable.YFilter
    dot1Qportvlanstatisticstable.EntityData.YangName = "dot1qPortVlanStatisticsTable"
    dot1Qportvlanstatisticstable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qportvlanstatisticstable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qportvlanstatisticstable.EntityData.SegmentPath = "dot1qPortVlanStatisticsTable"
    dot1Qportvlanstatisticstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qportvlanstatisticstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qportvlanstatisticstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qportvlanstatisticstable.EntityData.Children = make(map[string]types.YChild)
    dot1Qportvlanstatisticstable.EntityData.Children["dot1qPortVlanStatisticsEntry"] = types.YChild{"Dot1Qportvlanstatisticsentry", nil}
    for i := range dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry {
        dot1Qportvlanstatisticstable.EntityData.Children[types.GetSegmentPath(&dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry[i])] = types.YChild{"Dot1Qportvlanstatisticsentry", &dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry[i]}
    }
    dot1Qportvlanstatisticstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qportvlanstatisticstable.EntityData)
}

// QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry
// Traffic statistics for a VLAN on an interface.
type QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Dbaseport
    Dot1Dbaseport interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN.  Note that a frame received on this
    // port is counted by this object if and only if it is for a protocol being
    // processed by the local forwarding process for this VLAN.  This object
    // includes received bridge management frames classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP. The type is interface{} with range:
    // 0..4294967295.
    Dot1Qtpvlanportinframes interface{}

    // The number of valid frames transmitted by this port to its segment from the
    // local forwarding process for this VLAN.  This includes bridge management
    // frames originated by this device that are classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP). The type is interface{} with range:
    // 0..4294967295.
    Dot1Qtpvlanportoutframes interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN and that were discarded due to
    // VLAN-related reasons. Specifically, the IEEE 802.1Q counters for Discard
    // Inbound and Discard on Ingress Filtering. The type is interface{} with
    // range: 0..4294967295.
    Dot1Qtpvlanportindiscards interface{}

    // The number of times the associated dot1qTpVlanPortInFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1Qtpvlanportinoverflowframes interface{}

    // The number of times the associated dot1qTpVlanPortOutFrames counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1Qtpvlanportoutoverflowframes interface{}

    // The number of times the associated dot1qTpVlanPortInDiscards counter has
    // overflowed. The type is interface{} with range: 0..4294967295.
    Dot1Qtpvlanportinoverflowdiscards interface{}
}

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetEntityData() *types.CommonEntityData {
    dot1Qportvlanstatisticsentry.EntityData.YFilter = dot1Qportvlanstatisticsentry.YFilter
    dot1Qportvlanstatisticsentry.EntityData.YangName = "dot1qPortVlanStatisticsEntry"
    dot1Qportvlanstatisticsentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qportvlanstatisticsentry.EntityData.ParentYangName = "dot1qPortVlanStatisticsTable"
    dot1Qportvlanstatisticsentry.EntityData.SegmentPath = "dot1qPortVlanStatisticsEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Qportvlanstatisticsentry.Dot1Dbaseport) + "']" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qportvlanstatisticsentry.Dot1Qvlanindex) + "']"
    dot1Qportvlanstatisticsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qportvlanstatisticsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qportvlanstatisticsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qportvlanstatisticsentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qportvlanstatisticsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1dBasePort"] = types.YLeaf{"Dot1Dbaseport", dot1Qportvlanstatisticsentry.Dot1Dbaseport}
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qportvlanstatisticsentry.Dot1Qvlanindex}
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1qTpVlanPortInFrames"] = types.YLeaf{"Dot1Qtpvlanportinframes", dot1Qportvlanstatisticsentry.Dot1Qtpvlanportinframes}
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1qTpVlanPortOutFrames"] = types.YLeaf{"Dot1Qtpvlanportoutframes", dot1Qportvlanstatisticsentry.Dot1Qtpvlanportoutframes}
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1qTpVlanPortInDiscards"] = types.YLeaf{"Dot1Qtpvlanportindiscards", dot1Qportvlanstatisticsentry.Dot1Qtpvlanportindiscards}
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1qTpVlanPortInOverflowFrames"] = types.YLeaf{"Dot1Qtpvlanportinoverflowframes", dot1Qportvlanstatisticsentry.Dot1Qtpvlanportinoverflowframes}
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1qTpVlanPortOutOverflowFrames"] = types.YLeaf{"Dot1Qtpvlanportoutoverflowframes", dot1Qportvlanstatisticsentry.Dot1Qtpvlanportoutoverflowframes}
    dot1Qportvlanstatisticsentry.EntityData.Leafs["dot1qTpVlanPortInOverflowDiscards"] = types.YLeaf{"Dot1Qtpvlanportinoverflowdiscards", dot1Qportvlanstatisticsentry.Dot1Qtpvlanportinoverflowdiscards}
    return &(dot1Qportvlanstatisticsentry.EntityData)
}

// QBRIDGEMIB_Dot1Qportvlanhcstatisticstable
// A table containing per-port, per-VLAN statistics for
// traffic on high-capacity interfaces.
type QBRIDGEMIB_Dot1Qportvlanhcstatisticstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic statistics for a VLAN on a high-capacity interface. The type is
    // slice of
    // QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry.
    Dot1Qportvlanhcstatisticsentry []QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry
}

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetEntityData() *types.CommonEntityData {
    dot1Qportvlanhcstatisticstable.EntityData.YFilter = dot1Qportvlanhcstatisticstable.YFilter
    dot1Qportvlanhcstatisticstable.EntityData.YangName = "dot1qPortVlanHCStatisticsTable"
    dot1Qportvlanhcstatisticstable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qportvlanhcstatisticstable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qportvlanhcstatisticstable.EntityData.SegmentPath = "dot1qPortVlanHCStatisticsTable"
    dot1Qportvlanhcstatisticstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qportvlanhcstatisticstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qportvlanhcstatisticstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qportvlanhcstatisticstable.EntityData.Children = make(map[string]types.YChild)
    dot1Qportvlanhcstatisticstable.EntityData.Children["dot1qPortVlanHCStatisticsEntry"] = types.YChild{"Dot1Qportvlanhcstatisticsentry", nil}
    for i := range dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry {
        dot1Qportvlanhcstatisticstable.EntityData.Children[types.GetSegmentPath(&dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry[i])] = types.YChild{"Dot1Qportvlanhcstatisticsentry", &dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry[i]}
    }
    dot1Qportvlanhcstatisticstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qportvlanhcstatisticstable.EntityData)
}

// QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry
// Traffic statistics for a VLAN on a high-capacity
// interface.
type QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Dbaseport
    Dot1Dbaseport interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN.  Note that a frame received on this
    // port is counted by this object if and only if it is for a  protocol being
    // processed by the local forwarding process for this VLAN.  This object
    // includes received bridge management frames classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP). The type is interface{} with range:
    // 0..18446744073709551615.
    Dot1Qtpvlanporthcinframes interface{}

    // The number of valid frames transmitted by this port to its segment from the
    // local forwarding process for this VLAN.  This includes bridge management
    // frames originated by this device that are classified as belonging to this
    // VLAN (e.g., GMRP, but not GVRP or STP). The type is interface{} with range:
    // 0..18446744073709551615.
    Dot1Qtpvlanporthcoutframes interface{}

    // The number of valid frames received by this port from its segment that were
    // classified as belonging to this VLAN and that were discarded due to
    // VLAN-related reasons. Specifically, the IEEE 802.1Q counters for Discard
    // Inbound and Discard on Ingress Filtering. The type is interface{} with
    // range: 0..18446744073709551615.
    Dot1Qtpvlanporthcindiscards interface{}
}

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetEntityData() *types.CommonEntityData {
    dot1Qportvlanhcstatisticsentry.EntityData.YFilter = dot1Qportvlanhcstatisticsentry.YFilter
    dot1Qportvlanhcstatisticsentry.EntityData.YangName = "dot1qPortVlanHCStatisticsEntry"
    dot1Qportvlanhcstatisticsentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qportvlanhcstatisticsentry.EntityData.ParentYangName = "dot1qPortVlanHCStatisticsTable"
    dot1Qportvlanhcstatisticsentry.EntityData.SegmentPath = "dot1qPortVlanHCStatisticsEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Qportvlanhcstatisticsentry.Dot1Dbaseport) + "']" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qportvlanhcstatisticsentry.Dot1Qvlanindex) + "']"
    dot1Qportvlanhcstatisticsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qportvlanhcstatisticsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qportvlanhcstatisticsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qportvlanhcstatisticsentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qportvlanhcstatisticsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qportvlanhcstatisticsentry.EntityData.Leafs["dot1dBasePort"] = types.YLeaf{"Dot1Dbaseport", dot1Qportvlanhcstatisticsentry.Dot1Dbaseport}
    dot1Qportvlanhcstatisticsentry.EntityData.Leafs["dot1qVlanIndex"] = types.YLeaf{"Dot1Qvlanindex", dot1Qportvlanhcstatisticsentry.Dot1Qvlanindex}
    dot1Qportvlanhcstatisticsentry.EntityData.Leafs["dot1qTpVlanPortHCInFrames"] = types.YLeaf{"Dot1Qtpvlanporthcinframes", dot1Qportvlanhcstatisticsentry.Dot1Qtpvlanporthcinframes}
    dot1Qportvlanhcstatisticsentry.EntityData.Leafs["dot1qTpVlanPortHCOutFrames"] = types.YLeaf{"Dot1Qtpvlanporthcoutframes", dot1Qportvlanhcstatisticsentry.Dot1Qtpvlanporthcoutframes}
    dot1Qportvlanhcstatisticsentry.EntityData.Leafs["dot1qTpVlanPortHCInDiscards"] = types.YLeaf{"Dot1Qtpvlanporthcindiscards", dot1Qportvlanhcstatisticsentry.Dot1Qtpvlanporthcindiscards}
    return &(dot1Qportvlanhcstatisticsentry.EntityData)
}

// QBRIDGEMIB_Dot1Qlearningconstraintstable
// A table containing learning constraints for sets of
// Shared and Independent VLANs.
type QBRIDGEMIB_Dot1Qlearningconstraintstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A learning constraint defined for a VLAN. The type is slice of
    // QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry.
    Dot1Qlearningconstraintsentry []QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry
}

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetEntityData() *types.CommonEntityData {
    dot1Qlearningconstraintstable.EntityData.YFilter = dot1Qlearningconstraintstable.YFilter
    dot1Qlearningconstraintstable.EntityData.YangName = "dot1qLearningConstraintsTable"
    dot1Qlearningconstraintstable.EntityData.BundleName = "cisco_ios_xe"
    dot1Qlearningconstraintstable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Qlearningconstraintstable.EntityData.SegmentPath = "dot1qLearningConstraintsTable"
    dot1Qlearningconstraintstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qlearningconstraintstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qlearningconstraintstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qlearningconstraintstable.EntityData.Children = make(map[string]types.YChild)
    dot1Qlearningconstraintstable.EntityData.Children["dot1qLearningConstraintsEntry"] = types.YChild{"Dot1Qlearningconstraintsentry", nil}
    for i := range dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry {
        dot1Qlearningconstraintstable.EntityData.Children[types.GetSegmentPath(&dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry[i])] = types.YChild{"Dot1Qlearningconstraintsentry", &dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry[i]}
    }
    dot1Qlearningconstraintstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Qlearningconstraintstable.EntityData)
}

// QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry
// A learning constraint defined for a VLAN.
type QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the row in dot1qVlanCurrentTable for
    // the VLAN constrained by this entry. The type is interface{} with range:
    // 0..4294967295.
    Dot1Qconstraintvlan interface{}

    // This attribute is a key. The identity of the constraint set to which
    // dot1qConstraintVlan belongs.  These values may be chosen by the management
    // station. The type is interface{} with range: 0..65535.
    Dot1Qconstraintset interface{}

    // The type of constraint this entry defines. independent(1) - the VLAN,
    // dot1qConstraintVlan,     uses a filtering database independent from all    
    // other VLANs in the same set, defined by     dot1qConstraintSet. shared(2) -
    // the VLAN, dot1qConstraintVlan, shares     the same filtering database as
    // all other VLANs     in the same set, defined by dot1qConstraintSet. The
    // type is Dot1Qconstrainttype.
    Dot1Qconstrainttype interface{}

    // The status of this entry. The type is RowStatus.
    Dot1Qconstraintstatus interface{}
}

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetEntityData() *types.CommonEntityData {
    dot1Qlearningconstraintsentry.EntityData.YFilter = dot1Qlearningconstraintsentry.YFilter
    dot1Qlearningconstraintsentry.EntityData.YangName = "dot1qLearningConstraintsEntry"
    dot1Qlearningconstraintsentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Qlearningconstraintsentry.EntityData.ParentYangName = "dot1qLearningConstraintsTable"
    dot1Qlearningconstraintsentry.EntityData.SegmentPath = "dot1qLearningConstraintsEntry" + "[dot1qConstraintVlan='" + fmt.Sprintf("%v", dot1Qlearningconstraintsentry.Dot1Qconstraintvlan) + "']" + "[dot1qConstraintSet='" + fmt.Sprintf("%v", dot1Qlearningconstraintsentry.Dot1Qconstraintset) + "']"
    dot1Qlearningconstraintsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Qlearningconstraintsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Qlearningconstraintsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Qlearningconstraintsentry.EntityData.Children = make(map[string]types.YChild)
    dot1Qlearningconstraintsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Qlearningconstraintsentry.EntityData.Leafs["dot1qConstraintVlan"] = types.YLeaf{"Dot1Qconstraintvlan", dot1Qlearningconstraintsentry.Dot1Qconstraintvlan}
    dot1Qlearningconstraintsentry.EntityData.Leafs["dot1qConstraintSet"] = types.YLeaf{"Dot1Qconstraintset", dot1Qlearningconstraintsentry.Dot1Qconstraintset}
    dot1Qlearningconstraintsentry.EntityData.Leafs["dot1qConstraintType"] = types.YLeaf{"Dot1Qconstrainttype", dot1Qlearningconstraintsentry.Dot1Qconstrainttype}
    dot1Qlearningconstraintsentry.EntityData.Leafs["dot1qConstraintStatus"] = types.YLeaf{"Dot1Qconstraintstatus", dot1Qlearningconstraintsentry.Dot1Qconstraintstatus}
    return &(dot1Qlearningconstraintsentry.EntityData)
}

// QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry_Dot1Qconstrainttype represents     in the same set, defined by dot1qConstraintSet.
type QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry_Dot1Qconstrainttype string

const (
    QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry_Dot1Qconstrainttype_independent QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry_Dot1Qconstrainttype = "independent"

    QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry_Dot1Qconstrainttype_shared QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry_Dot1Qconstrainttype = "shared"
)

// QBRIDGEMIB_Dot1Vprotocolgrouptable
// A table that contains mappings from Protocol
// Templates to Protocol Group Identifiers used for
// Port-and-Protocol-based VLAN Classification.
type QBRIDGEMIB_Dot1Vprotocolgrouptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A mapping from a Protocol Template to a Protocol Group Identifier. The type
    // is slice of QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry.
    Dot1Vprotocolgroupentry []QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry
}

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetEntityData() *types.CommonEntityData {
    dot1Vprotocolgrouptable.EntityData.YFilter = dot1Vprotocolgrouptable.YFilter
    dot1Vprotocolgrouptable.EntityData.YangName = "dot1vProtocolGroupTable"
    dot1Vprotocolgrouptable.EntityData.BundleName = "cisco_ios_xe"
    dot1Vprotocolgrouptable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Vprotocolgrouptable.EntityData.SegmentPath = "dot1vProtocolGroupTable"
    dot1Vprotocolgrouptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Vprotocolgrouptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Vprotocolgrouptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Vprotocolgrouptable.EntityData.Children = make(map[string]types.YChild)
    dot1Vprotocolgrouptable.EntityData.Children["dot1vProtocolGroupEntry"] = types.YChild{"Dot1Vprotocolgroupentry", nil}
    for i := range dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry {
        dot1Vprotocolgrouptable.EntityData.Children[types.GetSegmentPath(&dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry[i])] = types.YChild{"Dot1Vprotocolgroupentry", &dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry[i]}
    }
    dot1Vprotocolgrouptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Vprotocolgrouptable.EntityData)
}

// QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry
// A mapping from a Protocol Template to a Protocol
// Group Identifier.
type QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The data-link encapsulation format or the
    // 'detagged_frame_type' in a Protocol Template. The type is
    // Dot1Vprotocoltemplateframetype.
    Dot1Vprotocoltemplateframetype interface{}

    // This attribute is a key. The identification of the protocol above the
    // data-link layer in a Protocol Template.  Depending on the frame type, the
    // octet string will have one of the following values:  For 'ethernet',
    // 'rfc1042' and 'snap8021H',     this is the 16-bit (2-octet) IEEE 802.3 Type
    // Field. For 'snapOther',     this is the 40-bit (5-octet) PID. For
    // 'llcOther',     this is the 2-octet IEEE 802.2 Link Service Access    
    // Point (LSAP) pair: first octet for Destination Service     Access Point
    // (DSAP) and second octet for Source Service     Access Point (SSAP). The
    // type is string with length: 2 | 5.
    Dot1Vprotocoltemplateprotocolvalue interface{}

    // Represents a group of protocols that are associated together when assigning
    // a VID to a frame. The type is interface{} with range: 0..2147483647.
    Dot1Vprotocolgroupid interface{}

    // This object indicates the status of this entry. The type is RowStatus.
    Dot1Vprotocolgrouprowstatus interface{}
}

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetEntityData() *types.CommonEntityData {
    dot1Vprotocolgroupentry.EntityData.YFilter = dot1Vprotocolgroupentry.YFilter
    dot1Vprotocolgroupentry.EntityData.YangName = "dot1vProtocolGroupEntry"
    dot1Vprotocolgroupentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Vprotocolgroupentry.EntityData.ParentYangName = "dot1vProtocolGroupTable"
    dot1Vprotocolgroupentry.EntityData.SegmentPath = "dot1vProtocolGroupEntry" + "[dot1vProtocolTemplateFrameType='" + fmt.Sprintf("%v", dot1Vprotocolgroupentry.Dot1Vprotocoltemplateframetype) + "']" + "[dot1vProtocolTemplateProtocolValue='" + fmt.Sprintf("%v", dot1Vprotocolgroupentry.Dot1Vprotocoltemplateprotocolvalue) + "']"
    dot1Vprotocolgroupentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Vprotocolgroupentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Vprotocolgroupentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Vprotocolgroupentry.EntityData.Children = make(map[string]types.YChild)
    dot1Vprotocolgroupentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Vprotocolgroupentry.EntityData.Leafs["dot1vProtocolTemplateFrameType"] = types.YLeaf{"Dot1Vprotocoltemplateframetype", dot1Vprotocolgroupentry.Dot1Vprotocoltemplateframetype}
    dot1Vprotocolgroupentry.EntityData.Leafs["dot1vProtocolTemplateProtocolValue"] = types.YLeaf{"Dot1Vprotocoltemplateprotocolvalue", dot1Vprotocolgroupentry.Dot1Vprotocoltemplateprotocolvalue}
    dot1Vprotocolgroupentry.EntityData.Leafs["dot1vProtocolGroupId"] = types.YLeaf{"Dot1Vprotocolgroupid", dot1Vprotocolgroupentry.Dot1Vprotocolgroupid}
    dot1Vprotocolgroupentry.EntityData.Leafs["dot1vProtocolGroupRowStatus"] = types.YLeaf{"Dot1Vprotocolgrouprowstatus", dot1Vprotocolgroupentry.Dot1Vprotocolgrouprowstatus}
    return &(dot1Vprotocolgroupentry.EntityData)
}

// QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype represents 'detagged_frame_type' in a Protocol Template.
type QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype string

const (
    QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype_ethernet QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype = "ethernet"

    QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype_rfc1042 QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype = "rfc1042"

    QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype_snap8021H QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype = "snap8021H"

    QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype_snapOther QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype = "snapOther"

    QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype_llcOther QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry_Dot1Vprotocoltemplateframetype = "llcOther"
)

// QBRIDGEMIB_Dot1Vprotocolporttable
// A table that contains VID sets used for
// Port-and-Protocol-based VLAN Classification.
type QBRIDGEMIB_Dot1Vprotocolporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VID set for a port. The type is slice of
    // QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry.
    Dot1Vprotocolportentry []QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry
}

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetEntityData() *types.CommonEntityData {
    dot1Vprotocolporttable.EntityData.YFilter = dot1Vprotocolporttable.YFilter
    dot1Vprotocolporttable.EntityData.YangName = "dot1vProtocolPortTable"
    dot1Vprotocolporttable.EntityData.BundleName = "cisco_ios_xe"
    dot1Vprotocolporttable.EntityData.ParentYangName = "Q-BRIDGE-MIB"
    dot1Vprotocolporttable.EntityData.SegmentPath = "dot1vProtocolPortTable"
    dot1Vprotocolporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Vprotocolporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Vprotocolporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Vprotocolporttable.EntityData.Children = make(map[string]types.YChild)
    dot1Vprotocolporttable.EntityData.Children["dot1vProtocolPortEntry"] = types.YChild{"Dot1Vprotocolportentry", nil}
    for i := range dot1Vprotocolporttable.Dot1Vprotocolportentry {
        dot1Vprotocolporttable.EntityData.Children[types.GetSegmentPath(&dot1Vprotocolporttable.Dot1Vprotocolportentry[i])] = types.YChild{"Dot1Vprotocolportentry", &dot1Vprotocolporttable.Dot1Vprotocolportentry[i]}
    }
    dot1Vprotocolporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1Vprotocolporttable.EntityData)
}

// QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry
// A VID set for a port.
type QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..65535. Refers to
    // bridge_mib.BRIDGEMIB_Dot1Dbaseporttable_Dot1Dbaseportentry_Dot1Dbaseport
    Dot1Dbaseport interface{}

    // This attribute is a key. Designates a group of protocols in the Protocol
    // Group Database. The type is interface{} with range: 1..2147483647.
    Dot1Vprotocolportgroupid interface{}

    // The VID associated with a group of protocols for each port. The type is
    // interface{} with range: 1..4094.
    Dot1Vprotocolportgroupvid interface{}

    // This object indicates the status of this entry. The type is RowStatus.
    Dot1Vprotocolportrowstatus interface{}
}

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetEntityData() *types.CommonEntityData {
    dot1Vprotocolportentry.EntityData.YFilter = dot1Vprotocolportentry.YFilter
    dot1Vprotocolportentry.EntityData.YangName = "dot1vProtocolPortEntry"
    dot1Vprotocolportentry.EntityData.BundleName = "cisco_ios_xe"
    dot1Vprotocolportentry.EntityData.ParentYangName = "dot1vProtocolPortTable"
    dot1Vprotocolportentry.EntityData.SegmentPath = "dot1vProtocolPortEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Vprotocolportentry.Dot1Dbaseport) + "']" + "[dot1vProtocolPortGroupId='" + fmt.Sprintf("%v", dot1Vprotocolportentry.Dot1Vprotocolportgroupid) + "']"
    dot1Vprotocolportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dot1Vprotocolportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dot1Vprotocolportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dot1Vprotocolportentry.EntityData.Children = make(map[string]types.YChild)
    dot1Vprotocolportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1Vprotocolportentry.EntityData.Leafs["dot1dBasePort"] = types.YLeaf{"Dot1Dbaseport", dot1Vprotocolportentry.Dot1Dbaseport}
    dot1Vprotocolportentry.EntityData.Leafs["dot1vProtocolPortGroupId"] = types.YLeaf{"Dot1Vprotocolportgroupid", dot1Vprotocolportentry.Dot1Vprotocolportgroupid}
    dot1Vprotocolportentry.EntityData.Leafs["dot1vProtocolPortGroupVid"] = types.YLeaf{"Dot1Vprotocolportgroupvid", dot1Vprotocolportentry.Dot1Vprotocolportgroupvid}
    dot1Vprotocolportentry.EntityData.Leafs["dot1vProtocolPortRowStatus"] = types.YLeaf{"Dot1Vprotocolportrowstatus", dot1Vprotocolportentry.Dot1Vprotocolportrowstatus}
    return &(dot1Vprotocolportentry.EntityData)
}

