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
    parent types.Entity
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

func (qBRIDGEMIB *QBRIDGEMIB) GetFilter() yfilter.YFilter { return qBRIDGEMIB.YFilter }

func (qBRIDGEMIB *QBRIDGEMIB) SetFilter(yf yfilter.YFilter) { qBRIDGEMIB.YFilter = yf }

func (qBRIDGEMIB *QBRIDGEMIB) GetGoName(yname string) string {
    if yname == "dot1qBase" { return "Dot1Qbase" }
    if yname == "dot1qVlan" { return "Dot1Qvlan" }
    if yname == "dot1qFdbTable" { return "Dot1Qfdbtable" }
    if yname == "dot1qTpFdbTable" { return "Dot1Qtpfdbtable" }
    if yname == "dot1qTpGroupTable" { return "Dot1Qtpgrouptable" }
    if yname == "dot1qForwardAllTable" { return "Dot1Qforwardalltable" }
    if yname == "dot1qForwardUnregisteredTable" { return "Dot1Qforwardunregisteredtable" }
    if yname == "dot1qStaticUnicastTable" { return "Dot1Qstaticunicasttable" }
    if yname == "dot1qStaticMulticastTable" { return "Dot1Qstaticmulticasttable" }
    if yname == "dot1qVlanCurrentTable" { return "Dot1Qvlancurrenttable" }
    if yname == "dot1qVlanStaticTable" { return "Dot1Qvlanstatictable" }
    if yname == "dot1qPortVlanStatisticsTable" { return "Dot1Qportvlanstatisticstable" }
    if yname == "dot1qPortVlanHCStatisticsTable" { return "Dot1Qportvlanhcstatisticstable" }
    if yname == "dot1qLearningConstraintsTable" { return "Dot1Qlearningconstraintstable" }
    if yname == "dot1vProtocolGroupTable" { return "Dot1Vprotocolgrouptable" }
    if yname == "dot1vProtocolPortTable" { return "Dot1Vprotocolporttable" }
    return ""
}

func (qBRIDGEMIB *QBRIDGEMIB) GetSegmentPath() string {
    return "Q-BRIDGE-MIB:Q-BRIDGE-MIB"
}

func (qBRIDGEMIB *QBRIDGEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qBase" {
        return &qBRIDGEMIB.Dot1Qbase
    }
    if childYangName == "dot1qVlan" {
        return &qBRIDGEMIB.Dot1Qvlan
    }
    if childYangName == "dot1qFdbTable" {
        return &qBRIDGEMIB.Dot1Qfdbtable
    }
    if childYangName == "dot1qTpFdbTable" {
        return &qBRIDGEMIB.Dot1Qtpfdbtable
    }
    if childYangName == "dot1qTpGroupTable" {
        return &qBRIDGEMIB.Dot1Qtpgrouptable
    }
    if childYangName == "dot1qForwardAllTable" {
        return &qBRIDGEMIB.Dot1Qforwardalltable
    }
    if childYangName == "dot1qForwardUnregisteredTable" {
        return &qBRIDGEMIB.Dot1Qforwardunregisteredtable
    }
    if childYangName == "dot1qStaticUnicastTable" {
        return &qBRIDGEMIB.Dot1Qstaticunicasttable
    }
    if childYangName == "dot1qStaticMulticastTable" {
        return &qBRIDGEMIB.Dot1Qstaticmulticasttable
    }
    if childYangName == "dot1qVlanCurrentTable" {
        return &qBRIDGEMIB.Dot1Qvlancurrenttable
    }
    if childYangName == "dot1qVlanStaticTable" {
        return &qBRIDGEMIB.Dot1Qvlanstatictable
    }
    if childYangName == "dot1qPortVlanStatisticsTable" {
        return &qBRIDGEMIB.Dot1Qportvlanstatisticstable
    }
    if childYangName == "dot1qPortVlanHCStatisticsTable" {
        return &qBRIDGEMIB.Dot1Qportvlanhcstatisticstable
    }
    if childYangName == "dot1qLearningConstraintsTable" {
        return &qBRIDGEMIB.Dot1Qlearningconstraintstable
    }
    if childYangName == "dot1vProtocolGroupTable" {
        return &qBRIDGEMIB.Dot1Vprotocolgrouptable
    }
    if childYangName == "dot1vProtocolPortTable" {
        return &qBRIDGEMIB.Dot1Vprotocolporttable
    }
    return nil
}

func (qBRIDGEMIB *QBRIDGEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dot1qBase"] = &qBRIDGEMIB.Dot1Qbase
    children["dot1qVlan"] = &qBRIDGEMIB.Dot1Qvlan
    children["dot1qFdbTable"] = &qBRIDGEMIB.Dot1Qfdbtable
    children["dot1qTpFdbTable"] = &qBRIDGEMIB.Dot1Qtpfdbtable
    children["dot1qTpGroupTable"] = &qBRIDGEMIB.Dot1Qtpgrouptable
    children["dot1qForwardAllTable"] = &qBRIDGEMIB.Dot1Qforwardalltable
    children["dot1qForwardUnregisteredTable"] = &qBRIDGEMIB.Dot1Qforwardunregisteredtable
    children["dot1qStaticUnicastTable"] = &qBRIDGEMIB.Dot1Qstaticunicasttable
    children["dot1qStaticMulticastTable"] = &qBRIDGEMIB.Dot1Qstaticmulticasttable
    children["dot1qVlanCurrentTable"] = &qBRIDGEMIB.Dot1Qvlancurrenttable
    children["dot1qVlanStaticTable"] = &qBRIDGEMIB.Dot1Qvlanstatictable
    children["dot1qPortVlanStatisticsTable"] = &qBRIDGEMIB.Dot1Qportvlanstatisticstable
    children["dot1qPortVlanHCStatisticsTable"] = &qBRIDGEMIB.Dot1Qportvlanhcstatisticstable
    children["dot1qLearningConstraintsTable"] = &qBRIDGEMIB.Dot1Qlearningconstraintstable
    children["dot1vProtocolGroupTable"] = &qBRIDGEMIB.Dot1Vprotocolgrouptable
    children["dot1vProtocolPortTable"] = &qBRIDGEMIB.Dot1Vprotocolporttable
    return children
}

func (qBRIDGEMIB *QBRIDGEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qBRIDGEMIB *QBRIDGEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (qBRIDGEMIB *QBRIDGEMIB) GetYangName() string { return "Q-BRIDGE-MIB" }

func (qBRIDGEMIB *QBRIDGEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qBRIDGEMIB *QBRIDGEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qBRIDGEMIB *QBRIDGEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qBRIDGEMIB *QBRIDGEMIB) SetParent(parent types.Entity) { qBRIDGEMIB.parent = parent }

func (qBRIDGEMIB *QBRIDGEMIB) GetParent() types.Entity { return qBRIDGEMIB.parent }

func (qBRIDGEMIB *QBRIDGEMIB) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qbase
type QBRIDGEMIB_Dot1Qbase struct {
    parent types.Entity
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

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetFilter() yfilter.YFilter { return dot1Qbase.YFilter }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) SetFilter(yf yfilter.YFilter) { dot1Qbase.YFilter = yf }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetGoName(yname string) string {
    if yname == "dot1qVlanVersionNumber" { return "Dot1Qvlanversionnumber" }
    if yname == "dot1qMaxVlanId" { return "Dot1Qmaxvlanid" }
    if yname == "dot1qMaxSupportedVlans" { return "Dot1Qmaxsupportedvlans" }
    if yname == "dot1qNumVlans" { return "Dot1Qnumvlans" }
    if yname == "dot1qGvrpStatus" { return "Dot1Qgvrpstatus" }
    return ""
}

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetSegmentPath() string {
    return "dot1qBase"
}

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanVersionNumber"] = dot1Qbase.Dot1Qvlanversionnumber
    leafs["dot1qMaxVlanId"] = dot1Qbase.Dot1Qmaxvlanid
    leafs["dot1qMaxSupportedVlans"] = dot1Qbase.Dot1Qmaxsupportedvlans
    leafs["dot1qNumVlans"] = dot1Qbase.Dot1Qnumvlans
    leafs["dot1qGvrpStatus"] = dot1Qbase.Dot1Qgvrpstatus
    return leafs
}

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetYangName() string { return "dot1qBase" }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) SetParent(parent types.Entity) { dot1Qbase.parent = parent }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetParent() types.Entity { return dot1Qbase.parent }

func (dot1Qbase *QBRIDGEMIB_Dot1Qbase) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber represents supports.
type QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber string

const (
    QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber_version1 QBRIDGEMIB_Dot1Qbase_Dot1Qvlanversionnumber = "version1"
)

// QBRIDGEMIB_Dot1Qvlan
type QBRIDGEMIB_Dot1Qvlan struct {
    parent types.Entity
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

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetFilter() yfilter.YFilter { return dot1Qvlan.YFilter }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) SetFilter(yf yfilter.YFilter) { dot1Qvlan.YFilter = yf }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetGoName(yname string) string {
    if yname == "dot1qVlanNumDeletes" { return "Dot1Qvlannumdeletes" }
    if yname == "dot1qNextFreeLocalVlanIndex" { return "Dot1Qnextfreelocalvlanindex" }
    if yname == "dot1qConstraintSetDefault" { return "Dot1Qconstraintsetdefault" }
    if yname == "dot1qConstraintTypeDefault" { return "Dot1Qconstrainttypedefault" }
    return ""
}

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetSegmentPath() string {
    return "dot1qVlan"
}

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanNumDeletes"] = dot1Qvlan.Dot1Qvlannumdeletes
    leafs["dot1qNextFreeLocalVlanIndex"] = dot1Qvlan.Dot1Qnextfreelocalvlanindex
    leafs["dot1qConstraintSetDefault"] = dot1Qvlan.Dot1Qconstraintsetdefault
    leafs["dot1qConstraintTypeDefault"] = dot1Qvlan.Dot1Qconstrainttypedefault
    return leafs
}

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetYangName() string { return "dot1qVlan" }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) SetParent(parent types.Entity) { dot1Qvlan.parent = parent }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetParent() types.Entity { return dot1Qvlan.parent }

func (dot1Qvlan *QBRIDGEMIB_Dot1Qvlan) GetParentYangName() string { return "Q-BRIDGE-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a specific Filtering Database. The type is slice of
    // QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry.
    Dot1Qfdbentry []QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry
}

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetFilter() yfilter.YFilter { return dot1Qfdbtable.YFilter }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) SetFilter(yf yfilter.YFilter) { dot1Qfdbtable.YFilter = yf }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetGoName(yname string) string {
    if yname == "dot1qFdbEntry" { return "Dot1Qfdbentry" }
    return ""
}

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetSegmentPath() string {
    return "dot1qFdbTable"
}

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qFdbEntry" {
        for _, c := range dot1Qfdbtable.Dot1Qfdbentry {
            if dot1Qfdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry{}
        dot1Qfdbtable.Dot1Qfdbentry = append(dot1Qfdbtable.Dot1Qfdbentry, child)
        return &dot1Qfdbtable.Dot1Qfdbentry[len(dot1Qfdbtable.Dot1Qfdbentry)-1]
    }
    return nil
}

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qfdbtable.Dot1Qfdbentry {
        children[dot1Qfdbtable.Dot1Qfdbentry[i].GetSegmentPath()] = &dot1Qfdbtable.Dot1Qfdbentry[i]
    }
    return children
}

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetYangName() string { return "dot1qFdbTable" }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) SetParent(parent types.Entity) { dot1Qfdbtable.parent = parent }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetParent() types.Entity { return dot1Qfdbtable.parent }

func (dot1Qfdbtable *QBRIDGEMIB_Dot1Qfdbtable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry
// Information about a specific Filtering Database.
type QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identity of this Filtering Database. The type
    // is interface{} with range: 0..4294967295.
    Dot1Qfdbid interface{}

    // The current number of dynamic entries in this Filtering Database. The type
    // is interface{} with range: 0..4294967295.
    Dot1Qfdbdynamiccount interface{}
}

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetFilter() yfilter.YFilter { return dot1Qfdbentry.YFilter }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) SetFilter(yf yfilter.YFilter) { dot1Qfdbentry.YFilter = yf }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetGoName(yname string) string {
    if yname == "dot1qFdbId" { return "Dot1Qfdbid" }
    if yname == "dot1qFdbDynamicCount" { return "Dot1Qfdbdynamiccount" }
    return ""
}

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetSegmentPath() string {
    return "dot1qFdbEntry" + "[dot1qFdbId='" + fmt.Sprintf("%v", dot1Qfdbentry.Dot1Qfdbid) + "']"
}

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qFdbId"] = dot1Qfdbentry.Dot1Qfdbid
    leafs["dot1qFdbDynamicCount"] = dot1Qfdbentry.Dot1Qfdbdynamiccount
    return leafs
}

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetYangName() string { return "dot1qFdbEntry" }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) SetParent(parent types.Entity) { dot1Qfdbentry.parent = parent }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetParent() types.Entity { return dot1Qfdbentry.parent }

func (dot1Qfdbentry *QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry) GetParentYangName() string { return "dot1qFdbTable" }

// QBRIDGEMIB_Dot1Qtpfdbtable
// A table that contains information about unicast entries
// for which the device has forwarding and/or filtering
// information.  This information is used by the
// transparent bridging function in determining how to
// propagate a received frame.
type QBRIDGEMIB_Dot1Qtpfdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a specific unicast MAC address for which the device has
    // some forwarding and/or filtering information. The type is slice of
    // QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry.
    Dot1Qtpfdbentry []QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry
}

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetFilter() yfilter.YFilter { return dot1Qtpfdbtable.YFilter }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) SetFilter(yf yfilter.YFilter) { dot1Qtpfdbtable.YFilter = yf }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetGoName(yname string) string {
    if yname == "dot1qTpFdbEntry" { return "Dot1Qtpfdbentry" }
    return ""
}

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetSegmentPath() string {
    return "dot1qTpFdbTable"
}

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qTpFdbEntry" {
        for _, c := range dot1Qtpfdbtable.Dot1Qtpfdbentry {
            if dot1Qtpfdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry{}
        dot1Qtpfdbtable.Dot1Qtpfdbentry = append(dot1Qtpfdbtable.Dot1Qtpfdbentry, child)
        return &dot1Qtpfdbtable.Dot1Qtpfdbentry[len(dot1Qtpfdbtable.Dot1Qtpfdbentry)-1]
    }
    return nil
}

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qtpfdbtable.Dot1Qtpfdbentry {
        children[dot1Qtpfdbtable.Dot1Qtpfdbentry[i].GetSegmentPath()] = &dot1Qtpfdbtable.Dot1Qtpfdbentry[i]
    }
    return children
}

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetYangName() string { return "dot1qTpFdbTable" }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) SetParent(parent types.Entity) { dot1Qtpfdbtable.parent = parent }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetParent() types.Entity { return dot1Qtpfdbtable.parent }

func (dot1Qtpfdbtable *QBRIDGEMIB_Dot1Qtpfdbtable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry
// Information about a specific unicast MAC address for
// which the device has some forwarding and/or filtering
// information.
type QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to q_bridge_mib.QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry_Dot1Qfdbid
    Dot1Qfdbid interface{}

    // This attribute is a key. A unicast MAC address for which the device has
    // forwarding and/or filtering information. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetFilter() yfilter.YFilter { return dot1Qtpfdbentry.YFilter }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) SetFilter(yf yfilter.YFilter) { dot1Qtpfdbentry.YFilter = yf }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetGoName(yname string) string {
    if yname == "dot1qFdbId" { return "Dot1Qfdbid" }
    if yname == "dot1qTpFdbAddress" { return "Dot1Qtpfdbaddress" }
    if yname == "dot1qTpFdbPort" { return "Dot1Qtpfdbport" }
    if yname == "dot1qTpFdbStatus" { return "Dot1Qtpfdbstatus" }
    return ""
}

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetSegmentPath() string {
    return "dot1qTpFdbEntry" + "[dot1qFdbId='" + fmt.Sprintf("%v", dot1Qtpfdbentry.Dot1Qfdbid) + "']" + "[dot1qTpFdbAddress='" + fmt.Sprintf("%v", dot1Qtpfdbentry.Dot1Qtpfdbaddress) + "']"
}

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qFdbId"] = dot1Qtpfdbentry.Dot1Qfdbid
    leafs["dot1qTpFdbAddress"] = dot1Qtpfdbentry.Dot1Qtpfdbaddress
    leafs["dot1qTpFdbPort"] = dot1Qtpfdbentry.Dot1Qtpfdbport
    leafs["dot1qTpFdbStatus"] = dot1Qtpfdbentry.Dot1Qtpfdbstatus
    return leafs
}

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetYangName() string { return "dot1qTpFdbEntry" }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) SetParent(parent types.Entity) { dot1Qtpfdbentry.parent = parent }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetParent() types.Entity { return dot1Qtpfdbentry.parent }

func (dot1Qtpfdbentry *QBRIDGEMIB_Dot1Qtpfdbtable_Dot1Qtpfdbentry) GetParentYangName() string { return "dot1qTpFdbTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Filtering information configured into the bridge by management, or learned
    // dynamically, specifying the set of ports to which frames received on a VLAN
    // and containing a specific Group destination address are allowed to be
    // forwarded.  The subset of these ports learned dynamically is also provided.
    // The type is slice of QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry.
    Dot1Qtpgroupentry []QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry
}

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetFilter() yfilter.YFilter { return dot1Qtpgrouptable.YFilter }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) SetFilter(yf yfilter.YFilter) { dot1Qtpgrouptable.YFilter = yf }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetGoName(yname string) string {
    if yname == "dot1qTpGroupEntry" { return "Dot1Qtpgroupentry" }
    return ""
}

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetSegmentPath() string {
    return "dot1qTpGroupTable"
}

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qTpGroupEntry" {
        for _, c := range dot1Qtpgrouptable.Dot1Qtpgroupentry {
            if dot1Qtpgrouptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry{}
        dot1Qtpgrouptable.Dot1Qtpgroupentry = append(dot1Qtpgrouptable.Dot1Qtpgroupentry, child)
        return &dot1Qtpgrouptable.Dot1Qtpgroupentry[len(dot1Qtpgrouptable.Dot1Qtpgroupentry)-1]
    }
    return nil
}

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qtpgrouptable.Dot1Qtpgroupentry {
        children[dot1Qtpgrouptable.Dot1Qtpgroupentry[i].GetSegmentPath()] = &dot1Qtpgrouptable.Dot1Qtpgroupentry[i]
    }
    return children
}

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetYangName() string { return "dot1qTpGroupTable" }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) SetParent(parent types.Entity) { dot1Qtpgrouptable.parent = parent }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetParent() types.Entity { return dot1Qtpgrouptable.parent }

func (dot1Qtpgrouptable *QBRIDGEMIB_Dot1Qtpgrouptable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry
// Filtering information configured into the bridge by
// management, or learned dynamically, specifying the set of
// ports to which frames received on a VLAN and containing
// a specific Group destination address are allowed to be
// forwarded.  The subset of these ports learned dynamically
// is also provided.
type QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // This attribute is a key. The destination Group MAC address in a frame to
    // which this entry's filtering information applies. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetFilter() yfilter.YFilter { return dot1Qtpgroupentry.YFilter }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) SetFilter(yf yfilter.YFilter) { dot1Qtpgroupentry.YFilter = yf }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetGoName(yname string) string {
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qTpGroupAddress" { return "Dot1Qtpgroupaddress" }
    if yname == "dot1qTpGroupEgressPorts" { return "Dot1Qtpgroupegressports" }
    if yname == "dot1qTpGroupLearnt" { return "Dot1Qtpgrouplearnt" }
    return ""
}

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetSegmentPath() string {
    return "dot1qTpGroupEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qtpgroupentry.Dot1Qvlanindex) + "']" + "[dot1qTpGroupAddress='" + fmt.Sprintf("%v", dot1Qtpgroupentry.Dot1Qtpgroupaddress) + "']"
}

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanIndex"] = dot1Qtpgroupentry.Dot1Qvlanindex
    leafs["dot1qTpGroupAddress"] = dot1Qtpgroupentry.Dot1Qtpgroupaddress
    leafs["dot1qTpGroupEgressPorts"] = dot1Qtpgroupentry.Dot1Qtpgroupegressports
    leafs["dot1qTpGroupLearnt"] = dot1Qtpgroupentry.Dot1Qtpgrouplearnt
    return leafs
}

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetYangName() string { return "dot1qTpGroupEntry" }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) SetParent(parent types.Entity) { dot1Qtpgroupentry.parent = parent }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetParent() types.Entity { return dot1Qtpgroupentry.parent }

func (dot1Qtpgroupentry *QBRIDGEMIB_Dot1Qtpgrouptable_Dot1Qtpgroupentry) GetParentYangName() string { return "dot1qTpGroupTable" }

// QBRIDGEMIB_Dot1Qforwardalltable
// A table containing forwarding information for each
// 
// VLAN, specifying the set of ports to which forwarding of
// all multicasts applies, configured statically by
// management or dynamically by GMRP.  An entry appears in
// this table for all VLANs that are currently
// instantiated.
type QBRIDGEMIB_Dot1Qforwardalltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Forwarding information for a VLAN, specifying the set of ports to which all
    // multicasts should be forwarded, configured statically by management or
    // dynamically by GMRP. The type is slice of
    // QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry.
    Dot1Qforwardallentry []QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry
}

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetFilter() yfilter.YFilter { return dot1Qforwardalltable.YFilter }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) SetFilter(yf yfilter.YFilter) { dot1Qforwardalltable.YFilter = yf }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetGoName(yname string) string {
    if yname == "dot1qForwardAllEntry" { return "Dot1Qforwardallentry" }
    return ""
}

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetSegmentPath() string {
    return "dot1qForwardAllTable"
}

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qForwardAllEntry" {
        for _, c := range dot1Qforwardalltable.Dot1Qforwardallentry {
            if dot1Qforwardalltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry{}
        dot1Qforwardalltable.Dot1Qforwardallentry = append(dot1Qforwardalltable.Dot1Qforwardallentry, child)
        return &dot1Qforwardalltable.Dot1Qforwardallentry[len(dot1Qforwardalltable.Dot1Qforwardallentry)-1]
    }
    return nil
}

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qforwardalltable.Dot1Qforwardallentry {
        children[dot1Qforwardalltable.Dot1Qforwardallentry[i].GetSegmentPath()] = &dot1Qforwardalltable.Dot1Qforwardallentry[i]
    }
    return children
}

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetYangName() string { return "dot1qForwardAllTable" }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) SetParent(parent types.Entity) { dot1Qforwardalltable.parent = parent }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetParent() types.Entity { return dot1Qforwardalltable.parent }

func (dot1Qforwardalltable *QBRIDGEMIB_Dot1Qforwardalltable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry
// Forwarding information for a VLAN, specifying the set
// of ports to which all multicasts should be forwarded,
// configured statically by management or dynamically by
// GMRP.
type QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry struct {
    parent types.Entity
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

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetFilter() yfilter.YFilter { return dot1Qforwardallentry.YFilter }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) SetFilter(yf yfilter.YFilter) { dot1Qforwardallentry.YFilter = yf }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetGoName(yname string) string {
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qForwardAllPorts" { return "Dot1Qforwardallports" }
    if yname == "dot1qForwardAllStaticPorts" { return "Dot1Qforwardallstaticports" }
    if yname == "dot1qForwardAllForbiddenPorts" { return "Dot1Qforwardallforbiddenports" }
    return ""
}

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetSegmentPath() string {
    return "dot1qForwardAllEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qforwardallentry.Dot1Qvlanindex) + "']"
}

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanIndex"] = dot1Qforwardallentry.Dot1Qvlanindex
    leafs["dot1qForwardAllPorts"] = dot1Qforwardallentry.Dot1Qforwardallports
    leafs["dot1qForwardAllStaticPorts"] = dot1Qforwardallentry.Dot1Qforwardallstaticports
    leafs["dot1qForwardAllForbiddenPorts"] = dot1Qforwardallentry.Dot1Qforwardallforbiddenports
    return leafs
}

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetYangName() string { return "dot1qForwardAllEntry" }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) SetParent(parent types.Entity) { dot1Qforwardallentry.parent = parent }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetParent() types.Entity { return dot1Qforwardallentry.parent }

func (dot1Qforwardallentry *QBRIDGEMIB_Dot1Qforwardalltable_Dot1Qforwardallentry) GetParentYangName() string { return "dot1qForwardAllTable" }

// QBRIDGEMIB_Dot1Qforwardunregisteredtable
// A table containing forwarding information for each
// VLAN, specifying the set of ports to which forwarding of
// multicast group-addressed frames for which no
// more specific forwarding information applies.  This is
// configured statically by management and determined
// dynamically by GMRP.  An entry appears in this table for
// all VLANs that are currently instantiated.
type QBRIDGEMIB_Dot1Qforwardunregisteredtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Forwarding information for a VLAN, specifying the set of ports to which all
    // multicasts for which there is no more specific forwarding information shall
    // be forwarded. This is configured statically by management or dynamically by
    // GMRP. The type is slice of
    // QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry.
    Dot1Qforwardunregisteredentry []QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry
}

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetFilter() yfilter.YFilter { return dot1Qforwardunregisteredtable.YFilter }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) SetFilter(yf yfilter.YFilter) { dot1Qforwardunregisteredtable.YFilter = yf }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetGoName(yname string) string {
    if yname == "dot1qForwardUnregisteredEntry" { return "Dot1Qforwardunregisteredentry" }
    return ""
}

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetSegmentPath() string {
    return "dot1qForwardUnregisteredTable"
}

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qForwardUnregisteredEntry" {
        for _, c := range dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry {
            if dot1Qforwardunregisteredtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry{}
        dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry = append(dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry, child)
        return &dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry[len(dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry)-1]
    }
    return nil
}

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry {
        children[dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry[i].GetSegmentPath()] = &dot1Qforwardunregisteredtable.Dot1Qforwardunregisteredentry[i]
    }
    return children
}

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetYangName() string { return "dot1qForwardUnregisteredTable" }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) SetParent(parent types.Entity) { dot1Qforwardunregisteredtable.parent = parent }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetParent() types.Entity { return dot1Qforwardunregisteredtable.parent }

func (dot1Qforwardunregisteredtable *QBRIDGEMIB_Dot1Qforwardunregisteredtable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry
// Forwarding information for a VLAN, specifying the set
// of ports to which all multicasts for which there is no
// more specific forwarding information shall be forwarded.
// This is configured statically by management or
// dynamically by GMRP.
type QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry struct {
    parent types.Entity
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

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetFilter() yfilter.YFilter { return dot1Qforwardunregisteredentry.YFilter }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) SetFilter(yf yfilter.YFilter) { dot1Qforwardunregisteredentry.YFilter = yf }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetGoName(yname string) string {
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qForwardUnregisteredPorts" { return "Dot1Qforwardunregisteredports" }
    if yname == "dot1qForwardUnregisteredStaticPorts" { return "Dot1Qforwardunregisteredstaticports" }
    if yname == "dot1qForwardUnregisteredForbiddenPorts" { return "Dot1Qforwardunregisteredforbiddenports" }
    return ""
}

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetSegmentPath() string {
    return "dot1qForwardUnregisteredEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qforwardunregisteredentry.Dot1Qvlanindex) + "']"
}

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanIndex"] = dot1Qforwardunregisteredentry.Dot1Qvlanindex
    leafs["dot1qForwardUnregisteredPorts"] = dot1Qforwardunregisteredentry.Dot1Qforwardunregisteredports
    leafs["dot1qForwardUnregisteredStaticPorts"] = dot1Qforwardunregisteredentry.Dot1Qforwardunregisteredstaticports
    leafs["dot1qForwardUnregisteredForbiddenPorts"] = dot1Qforwardunregisteredentry.Dot1Qforwardunregisteredforbiddenports
    return leafs
}

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetYangName() string { return "dot1qForwardUnregisteredEntry" }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) SetParent(parent types.Entity) { dot1Qforwardunregisteredentry.parent = parent }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetParent() types.Entity { return dot1Qforwardunregisteredentry.parent }

func (dot1Qforwardunregisteredentry *QBRIDGEMIB_Dot1Qforwardunregisteredtable_Dot1Qforwardunregisteredentry) GetParentYangName() string { return "dot1qForwardUnregisteredTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Filtering information configured into the device by (local or network)
    // management specifying the set of ports to which frames received from a
    // specific port and containing a specific unicast destination address are
    // allowed to be forwarded. The type is slice of
    // QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry.
    Dot1Qstaticunicastentry []QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry
}

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetFilter() yfilter.YFilter { return dot1Qstaticunicasttable.YFilter }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) SetFilter(yf yfilter.YFilter) { dot1Qstaticunicasttable.YFilter = yf }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetGoName(yname string) string {
    if yname == "dot1qStaticUnicastEntry" { return "Dot1Qstaticunicastentry" }
    return ""
}

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetSegmentPath() string {
    return "dot1qStaticUnicastTable"
}

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qStaticUnicastEntry" {
        for _, c := range dot1Qstaticunicasttable.Dot1Qstaticunicastentry {
            if dot1Qstaticunicasttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry{}
        dot1Qstaticunicasttable.Dot1Qstaticunicastentry = append(dot1Qstaticunicasttable.Dot1Qstaticunicastentry, child)
        return &dot1Qstaticunicasttable.Dot1Qstaticunicastentry[len(dot1Qstaticunicasttable.Dot1Qstaticunicastentry)-1]
    }
    return nil
}

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qstaticunicasttable.Dot1Qstaticunicastentry {
        children[dot1Qstaticunicasttable.Dot1Qstaticunicastentry[i].GetSegmentPath()] = &dot1Qstaticunicasttable.Dot1Qstaticunicastentry[i]
    }
    return children
}

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetYangName() string { return "dot1qStaticUnicastTable" }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) SetParent(parent types.Entity) { dot1Qstaticunicasttable.parent = parent }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetParent() types.Entity { return dot1Qstaticunicasttable.parent }

func (dot1Qstaticunicasttable *QBRIDGEMIB_Dot1Qstaticunicasttable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry
// Filtering information configured into the device by
// (local or network) management specifying the set of
// ports to which frames received from a specific port and
// containing a specific unicast destination address are
// allowed to be forwarded.
type QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to q_bridge_mib.QBRIDGEMIB_Dot1Qfdbtable_Dot1Qfdbentry_Dot1Qfdbid
    Dot1Qfdbid interface{}

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object must take the
    // value of a unicast address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetFilter() yfilter.YFilter { return dot1Qstaticunicastentry.YFilter }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) SetFilter(yf yfilter.YFilter) { dot1Qstaticunicastentry.YFilter = yf }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetGoName(yname string) string {
    if yname == "dot1qFdbId" { return "Dot1Qfdbid" }
    if yname == "dot1qStaticUnicastAddress" { return "Dot1Qstaticunicastaddress" }
    if yname == "dot1qStaticUnicastReceivePort" { return "Dot1Qstaticunicastreceiveport" }
    if yname == "dot1qStaticUnicastAllowedToGoTo" { return "Dot1Qstaticunicastallowedtogoto" }
    if yname == "dot1qStaticUnicastStatus" { return "Dot1Qstaticunicaststatus" }
    return ""
}

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetSegmentPath() string {
    return "dot1qStaticUnicastEntry" + "[dot1qFdbId='" + fmt.Sprintf("%v", dot1Qstaticunicastentry.Dot1Qfdbid) + "']" + "[dot1qStaticUnicastAddress='" + fmt.Sprintf("%v", dot1Qstaticunicastentry.Dot1Qstaticunicastaddress) + "']" + "[dot1qStaticUnicastReceivePort='" + fmt.Sprintf("%v", dot1Qstaticunicastentry.Dot1Qstaticunicastreceiveport) + "']"
}

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qFdbId"] = dot1Qstaticunicastentry.Dot1Qfdbid
    leafs["dot1qStaticUnicastAddress"] = dot1Qstaticunicastentry.Dot1Qstaticunicastaddress
    leafs["dot1qStaticUnicastReceivePort"] = dot1Qstaticunicastentry.Dot1Qstaticunicastreceiveport
    leafs["dot1qStaticUnicastAllowedToGoTo"] = dot1Qstaticunicastentry.Dot1Qstaticunicastallowedtogoto
    leafs["dot1qStaticUnicastStatus"] = dot1Qstaticunicastentry.Dot1Qstaticunicaststatus
    return leafs
}

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetYangName() string { return "dot1qStaticUnicastEntry" }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) SetParent(parent types.Entity) { dot1Qstaticunicastentry.parent = parent }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetParent() types.Entity { return dot1Qstaticunicastentry.parent }

func (dot1Qstaticunicastentry *QBRIDGEMIB_Dot1Qstaticunicasttable_Dot1Qstaticunicastentry) GetParentYangName() string { return "dot1qStaticUnicastTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Filtering information configured into the device by (local or network)
    // management specifying the set of ports to which frames received from this
    // specific port  for this VLAN and containing this Multicast or Broadcast
    // destination address are allowed to be forwarded. The type is slice of
    // QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry.
    Dot1Qstaticmulticastentry []QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry
}

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetFilter() yfilter.YFilter { return dot1Qstaticmulticasttable.YFilter }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) SetFilter(yf yfilter.YFilter) { dot1Qstaticmulticasttable.YFilter = yf }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetGoName(yname string) string {
    if yname == "dot1qStaticMulticastEntry" { return "Dot1Qstaticmulticastentry" }
    return ""
}

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetSegmentPath() string {
    return "dot1qStaticMulticastTable"
}

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qStaticMulticastEntry" {
        for _, c := range dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry {
            if dot1Qstaticmulticasttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry{}
        dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry = append(dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry, child)
        return &dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry[len(dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry)-1]
    }
    return nil
}

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry {
        children[dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry[i].GetSegmentPath()] = &dot1Qstaticmulticasttable.Dot1Qstaticmulticastentry[i]
    }
    return children
}

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetYangName() string { return "dot1qStaticMulticastTable" }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) SetParent(parent types.Entity) { dot1Qstaticmulticasttable.parent = parent }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetParent() types.Entity { return dot1Qstaticmulticasttable.parent }

func (dot1Qstaticmulticasttable *QBRIDGEMIB_Dot1Qstaticmulticasttable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry
// Filtering information configured into the device by
// (local or network) management specifying the set of
// ports to which frames received from this specific port
// 
// for this VLAN and containing this Multicast or Broadcast
// destination address are allowed to be forwarded.
type QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // q_bridge_mib.QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry_Dot1Qvlanindex
    Dot1Qvlanindex interface{}

    // This attribute is a key. The destination MAC address in a frame to which
    // this entry's filtering information applies.  This object must take the
    // value of a Multicast or Broadcast address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetFilter() yfilter.YFilter { return dot1Qstaticmulticastentry.YFilter }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) SetFilter(yf yfilter.YFilter) { dot1Qstaticmulticastentry.YFilter = yf }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetGoName(yname string) string {
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qStaticMulticastAddress" { return "Dot1Qstaticmulticastaddress" }
    if yname == "dot1qStaticMulticastReceivePort" { return "Dot1Qstaticmulticastreceiveport" }
    if yname == "dot1qStaticMulticastStaticEgressPorts" { return "Dot1Qstaticmulticaststaticegressports" }
    if yname == "dot1qStaticMulticastForbiddenEgressPorts" { return "Dot1Qstaticmulticastforbiddenegressports" }
    if yname == "dot1qStaticMulticastStatus" { return "Dot1Qstaticmulticaststatus" }
    return ""
}

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetSegmentPath() string {
    return "dot1qStaticMulticastEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qstaticmulticastentry.Dot1Qvlanindex) + "']" + "[dot1qStaticMulticastAddress='" + fmt.Sprintf("%v", dot1Qstaticmulticastentry.Dot1Qstaticmulticastaddress) + "']" + "[dot1qStaticMulticastReceivePort='" + fmt.Sprintf("%v", dot1Qstaticmulticastentry.Dot1Qstaticmulticastreceiveport) + "']"
}

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanIndex"] = dot1Qstaticmulticastentry.Dot1Qvlanindex
    leafs["dot1qStaticMulticastAddress"] = dot1Qstaticmulticastentry.Dot1Qstaticmulticastaddress
    leafs["dot1qStaticMulticastReceivePort"] = dot1Qstaticmulticastentry.Dot1Qstaticmulticastreceiveport
    leafs["dot1qStaticMulticastStaticEgressPorts"] = dot1Qstaticmulticastentry.Dot1Qstaticmulticaststaticegressports
    leafs["dot1qStaticMulticastForbiddenEgressPorts"] = dot1Qstaticmulticastentry.Dot1Qstaticmulticastforbiddenegressports
    leafs["dot1qStaticMulticastStatus"] = dot1Qstaticmulticastentry.Dot1Qstaticmulticaststatus
    return leafs
}

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetYangName() string { return "dot1qStaticMulticastEntry" }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) SetParent(parent types.Entity) { dot1Qstaticmulticastentry.parent = parent }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetParent() types.Entity { return dot1Qstaticmulticastentry.parent }

func (dot1Qstaticmulticastentry *QBRIDGEMIB_Dot1Qstaticmulticasttable_Dot1Qstaticmulticastentry) GetParentYangName() string { return "dot1qStaticMulticastTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information for a VLAN configured into the device by  (local or network)
    // management, or dynamically created as a result of GVRP requests received.
    // The type is slice of
    // QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry.
    Dot1Qvlancurrententry []QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry
}

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetFilter() yfilter.YFilter { return dot1Qvlancurrenttable.YFilter }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) SetFilter(yf yfilter.YFilter) { dot1Qvlancurrenttable.YFilter = yf }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetGoName(yname string) string {
    if yname == "dot1qVlanCurrentEntry" { return "Dot1Qvlancurrententry" }
    return ""
}

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetSegmentPath() string {
    return "dot1qVlanCurrentTable"
}

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qVlanCurrentEntry" {
        for _, c := range dot1Qvlancurrenttable.Dot1Qvlancurrententry {
            if dot1Qvlancurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry{}
        dot1Qvlancurrenttable.Dot1Qvlancurrententry = append(dot1Qvlancurrenttable.Dot1Qvlancurrententry, child)
        return &dot1Qvlancurrenttable.Dot1Qvlancurrententry[len(dot1Qvlancurrenttable.Dot1Qvlancurrententry)-1]
    }
    return nil
}

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qvlancurrenttable.Dot1Qvlancurrententry {
        children[dot1Qvlancurrenttable.Dot1Qvlancurrententry[i].GetSegmentPath()] = &dot1Qvlancurrenttable.Dot1Qvlancurrententry[i]
    }
    return children
}

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetYangName() string { return "dot1qVlanCurrentTable" }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) SetParent(parent types.Entity) { dot1Qvlancurrenttable.parent = parent }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetParent() types.Entity { return dot1Qvlancurrenttable.parent }

func (dot1Qvlancurrenttable *QBRIDGEMIB_Dot1Qvlancurrenttable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry
// Information for a VLAN configured into the device by
// 
// (local or network) management, or dynamically created
// as a result of GVRP requests received.
type QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry struct {
    parent types.Entity
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

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetFilter() yfilter.YFilter { return dot1Qvlancurrententry.YFilter }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) SetFilter(yf yfilter.YFilter) { dot1Qvlancurrententry.YFilter = yf }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetGoName(yname string) string {
    if yname == "dot1qVlanTimeMark" { return "Dot1Qvlantimemark" }
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qVlanFdbId" { return "Dot1Qvlanfdbid" }
    if yname == "dot1qVlanCurrentEgressPorts" { return "Dot1Qvlancurrentegressports" }
    if yname == "dot1qVlanCurrentUntaggedPorts" { return "Dot1Qvlancurrentuntaggedports" }
    if yname == "dot1qVlanStatus" { return "Dot1Qvlanstatus" }
    if yname == "dot1qVlanCreationTime" { return "Dot1Qvlancreationtime" }
    return ""
}

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetSegmentPath() string {
    return "dot1qVlanCurrentEntry" + "[dot1qVlanTimeMark='" + fmt.Sprintf("%v", dot1Qvlancurrententry.Dot1Qvlantimemark) + "']" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qvlancurrententry.Dot1Qvlanindex) + "']"
}

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanTimeMark"] = dot1Qvlancurrententry.Dot1Qvlantimemark
    leafs["dot1qVlanIndex"] = dot1Qvlancurrententry.Dot1Qvlanindex
    leafs["dot1qVlanFdbId"] = dot1Qvlancurrententry.Dot1Qvlanfdbid
    leafs["dot1qVlanCurrentEgressPorts"] = dot1Qvlancurrententry.Dot1Qvlancurrentegressports
    leafs["dot1qVlanCurrentUntaggedPorts"] = dot1Qvlancurrententry.Dot1Qvlancurrentuntaggedports
    leafs["dot1qVlanStatus"] = dot1Qvlancurrententry.Dot1Qvlanstatus
    leafs["dot1qVlanCreationTime"] = dot1Qvlancurrententry.Dot1Qvlancreationtime
    return leafs
}

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetYangName() string { return "dot1qVlanCurrentEntry" }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) SetParent(parent types.Entity) { dot1Qvlancurrententry.parent = parent }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetParent() types.Entity { return dot1Qvlancurrententry.parent }

func (dot1Qvlancurrententry *QBRIDGEMIB_Dot1Qvlancurrenttable_Dot1Qvlancurrententry) GetParentYangName() string { return "dot1qVlanCurrentTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Static information for a VLAN configured into the device by (local or
    // network) management. The type is slice of
    // QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry.
    Dot1Qvlanstaticentry []QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry
}

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetFilter() yfilter.YFilter { return dot1Qvlanstatictable.YFilter }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) SetFilter(yf yfilter.YFilter) { dot1Qvlanstatictable.YFilter = yf }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetGoName(yname string) string {
    if yname == "dot1qVlanStaticEntry" { return "Dot1Qvlanstaticentry" }
    return ""
}

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetSegmentPath() string {
    return "dot1qVlanStaticTable"
}

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qVlanStaticEntry" {
        for _, c := range dot1Qvlanstatictable.Dot1Qvlanstaticentry {
            if dot1Qvlanstatictable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry{}
        dot1Qvlanstatictable.Dot1Qvlanstaticentry = append(dot1Qvlanstatictable.Dot1Qvlanstaticentry, child)
        return &dot1Qvlanstatictable.Dot1Qvlanstaticentry[len(dot1Qvlanstatictable.Dot1Qvlanstaticentry)-1]
    }
    return nil
}

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qvlanstatictable.Dot1Qvlanstaticentry {
        children[dot1Qvlanstatictable.Dot1Qvlanstaticentry[i].GetSegmentPath()] = &dot1Qvlanstatictable.Dot1Qvlanstaticentry[i]
    }
    return children
}

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetYangName() string { return "dot1qVlanStaticTable" }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) SetParent(parent types.Entity) { dot1Qvlanstatictable.parent = parent }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetParent() types.Entity { return dot1Qvlanstatictable.parent }

func (dot1Qvlanstatictable *QBRIDGEMIB_Dot1Qvlanstatictable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry
// Static information for a VLAN configured into the
// device by (local or network) management.
type QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry struct {
    parent types.Entity
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

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetFilter() yfilter.YFilter { return dot1Qvlanstaticentry.YFilter }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) SetFilter(yf yfilter.YFilter) { dot1Qvlanstaticentry.YFilter = yf }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetGoName(yname string) string {
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qVlanStaticName" { return "Dot1Qvlanstaticname" }
    if yname == "dot1qVlanStaticEgressPorts" { return "Dot1Qvlanstaticegressports" }
    if yname == "dot1qVlanForbiddenEgressPorts" { return "Dot1Qvlanforbiddenegressports" }
    if yname == "dot1qVlanStaticUntaggedPorts" { return "Dot1Qvlanstaticuntaggedports" }
    if yname == "dot1qVlanStaticRowStatus" { return "Dot1Qvlanstaticrowstatus" }
    return ""
}

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetSegmentPath() string {
    return "dot1qVlanStaticEntry" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qvlanstaticentry.Dot1Qvlanindex) + "']"
}

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qVlanIndex"] = dot1Qvlanstaticentry.Dot1Qvlanindex
    leafs["dot1qVlanStaticName"] = dot1Qvlanstaticentry.Dot1Qvlanstaticname
    leafs["dot1qVlanStaticEgressPorts"] = dot1Qvlanstaticentry.Dot1Qvlanstaticegressports
    leafs["dot1qVlanForbiddenEgressPorts"] = dot1Qvlanstaticentry.Dot1Qvlanforbiddenegressports
    leafs["dot1qVlanStaticUntaggedPorts"] = dot1Qvlanstaticentry.Dot1Qvlanstaticuntaggedports
    leafs["dot1qVlanStaticRowStatus"] = dot1Qvlanstaticentry.Dot1Qvlanstaticrowstatus
    return leafs
}

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetYangName() string { return "dot1qVlanStaticEntry" }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) SetParent(parent types.Entity) { dot1Qvlanstaticentry.parent = parent }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetParent() types.Entity { return dot1Qvlanstaticentry.parent }

func (dot1Qvlanstaticentry *QBRIDGEMIB_Dot1Qvlanstatictable_Dot1Qvlanstaticentry) GetParentYangName() string { return "dot1qVlanStaticTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Traffic statistics for a VLAN on an interface. The type is slice of
    // QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry.
    Dot1Qportvlanstatisticsentry []QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry
}

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetFilter() yfilter.YFilter { return dot1Qportvlanstatisticstable.YFilter }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) SetFilter(yf yfilter.YFilter) { dot1Qportvlanstatisticstable.YFilter = yf }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetGoName(yname string) string {
    if yname == "dot1qPortVlanStatisticsEntry" { return "Dot1Qportvlanstatisticsentry" }
    return ""
}

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetSegmentPath() string {
    return "dot1qPortVlanStatisticsTable"
}

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qPortVlanStatisticsEntry" {
        for _, c := range dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry {
            if dot1Qportvlanstatisticstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry{}
        dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry = append(dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry, child)
        return &dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry[len(dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry)-1]
    }
    return nil
}

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry {
        children[dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry[i].GetSegmentPath()] = &dot1Qportvlanstatisticstable.Dot1Qportvlanstatisticsentry[i]
    }
    return children
}

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetYangName() string { return "dot1qPortVlanStatisticsTable" }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) SetParent(parent types.Entity) { dot1Qportvlanstatisticstable.parent = parent }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetParent() types.Entity { return dot1Qportvlanstatisticstable.parent }

func (dot1Qportvlanstatisticstable *QBRIDGEMIB_Dot1Qportvlanstatisticstable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry
// Traffic statistics for a VLAN on an interface.
type QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry struct {
    parent types.Entity
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

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetFilter() yfilter.YFilter { return dot1Qportvlanstatisticsentry.YFilter }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) SetFilter(yf yfilter.YFilter) { dot1Qportvlanstatisticsentry.YFilter = yf }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetGoName(yname string) string {
    if yname == "dot1dBasePort" { return "Dot1Dbaseport" }
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qTpVlanPortInFrames" { return "Dot1Qtpvlanportinframes" }
    if yname == "dot1qTpVlanPortOutFrames" { return "Dot1Qtpvlanportoutframes" }
    if yname == "dot1qTpVlanPortInDiscards" { return "Dot1Qtpvlanportindiscards" }
    if yname == "dot1qTpVlanPortInOverflowFrames" { return "Dot1Qtpvlanportinoverflowframes" }
    if yname == "dot1qTpVlanPortOutOverflowFrames" { return "Dot1Qtpvlanportoutoverflowframes" }
    if yname == "dot1qTpVlanPortInOverflowDiscards" { return "Dot1Qtpvlanportinoverflowdiscards" }
    return ""
}

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetSegmentPath() string {
    return "dot1qPortVlanStatisticsEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Qportvlanstatisticsentry.Dot1Dbaseport) + "']" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qportvlanstatisticsentry.Dot1Qvlanindex) + "']"
}

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBasePort"] = dot1Qportvlanstatisticsentry.Dot1Dbaseport
    leafs["dot1qVlanIndex"] = dot1Qportvlanstatisticsentry.Dot1Qvlanindex
    leafs["dot1qTpVlanPortInFrames"] = dot1Qportvlanstatisticsentry.Dot1Qtpvlanportinframes
    leafs["dot1qTpVlanPortOutFrames"] = dot1Qportvlanstatisticsentry.Dot1Qtpvlanportoutframes
    leafs["dot1qTpVlanPortInDiscards"] = dot1Qportvlanstatisticsentry.Dot1Qtpvlanportindiscards
    leafs["dot1qTpVlanPortInOverflowFrames"] = dot1Qportvlanstatisticsentry.Dot1Qtpvlanportinoverflowframes
    leafs["dot1qTpVlanPortOutOverflowFrames"] = dot1Qportvlanstatisticsentry.Dot1Qtpvlanportoutoverflowframes
    leafs["dot1qTpVlanPortInOverflowDiscards"] = dot1Qportvlanstatisticsentry.Dot1Qtpvlanportinoverflowdiscards
    return leafs
}

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetYangName() string { return "dot1qPortVlanStatisticsEntry" }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) SetParent(parent types.Entity) { dot1Qportvlanstatisticsentry.parent = parent }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetParent() types.Entity { return dot1Qportvlanstatisticsentry.parent }

func (dot1Qportvlanstatisticsentry *QBRIDGEMIB_Dot1Qportvlanstatisticstable_Dot1Qportvlanstatisticsentry) GetParentYangName() string { return "dot1qPortVlanStatisticsTable" }

// QBRIDGEMIB_Dot1Qportvlanhcstatisticstable
// A table containing per-port, per-VLAN statistics for
// traffic on high-capacity interfaces.
type QBRIDGEMIB_Dot1Qportvlanhcstatisticstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Traffic statistics for a VLAN on a high-capacity interface. The type is
    // slice of
    // QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry.
    Dot1Qportvlanhcstatisticsentry []QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry
}

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetFilter() yfilter.YFilter { return dot1Qportvlanhcstatisticstable.YFilter }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) SetFilter(yf yfilter.YFilter) { dot1Qportvlanhcstatisticstable.YFilter = yf }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetGoName(yname string) string {
    if yname == "dot1qPortVlanHCStatisticsEntry" { return "Dot1Qportvlanhcstatisticsentry" }
    return ""
}

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetSegmentPath() string {
    return "dot1qPortVlanHCStatisticsTable"
}

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qPortVlanHCStatisticsEntry" {
        for _, c := range dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry {
            if dot1Qportvlanhcstatisticstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry{}
        dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry = append(dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry, child)
        return &dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry[len(dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry)-1]
    }
    return nil
}

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry {
        children[dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry[i].GetSegmentPath()] = &dot1Qportvlanhcstatisticstable.Dot1Qportvlanhcstatisticsentry[i]
    }
    return children
}

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetYangName() string { return "dot1qPortVlanHCStatisticsTable" }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) SetParent(parent types.Entity) { dot1Qportvlanhcstatisticstable.parent = parent }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetParent() types.Entity { return dot1Qportvlanhcstatisticstable.parent }

func (dot1Qportvlanhcstatisticstable *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry
// Traffic statistics for a VLAN on a high-capacity
// interface.
type QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry struct {
    parent types.Entity
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

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetFilter() yfilter.YFilter { return dot1Qportvlanhcstatisticsentry.YFilter }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) SetFilter(yf yfilter.YFilter) { dot1Qportvlanhcstatisticsentry.YFilter = yf }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetGoName(yname string) string {
    if yname == "dot1dBasePort" { return "Dot1Dbaseport" }
    if yname == "dot1qVlanIndex" { return "Dot1Qvlanindex" }
    if yname == "dot1qTpVlanPortHCInFrames" { return "Dot1Qtpvlanporthcinframes" }
    if yname == "dot1qTpVlanPortHCOutFrames" { return "Dot1Qtpvlanporthcoutframes" }
    if yname == "dot1qTpVlanPortHCInDiscards" { return "Dot1Qtpvlanporthcindiscards" }
    return ""
}

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetSegmentPath() string {
    return "dot1qPortVlanHCStatisticsEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Qportvlanhcstatisticsentry.Dot1Dbaseport) + "']" + "[dot1qVlanIndex='" + fmt.Sprintf("%v", dot1Qportvlanhcstatisticsentry.Dot1Qvlanindex) + "']"
}

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBasePort"] = dot1Qportvlanhcstatisticsentry.Dot1Dbaseport
    leafs["dot1qVlanIndex"] = dot1Qportvlanhcstatisticsentry.Dot1Qvlanindex
    leafs["dot1qTpVlanPortHCInFrames"] = dot1Qportvlanhcstatisticsentry.Dot1Qtpvlanporthcinframes
    leafs["dot1qTpVlanPortHCOutFrames"] = dot1Qportvlanhcstatisticsentry.Dot1Qtpvlanporthcoutframes
    leafs["dot1qTpVlanPortHCInDiscards"] = dot1Qportvlanhcstatisticsentry.Dot1Qtpvlanporthcindiscards
    return leafs
}

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetYangName() string { return "dot1qPortVlanHCStatisticsEntry" }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) SetParent(parent types.Entity) { dot1Qportvlanhcstatisticsentry.parent = parent }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetParent() types.Entity { return dot1Qportvlanhcstatisticsentry.parent }

func (dot1Qportvlanhcstatisticsentry *QBRIDGEMIB_Dot1Qportvlanhcstatisticstable_Dot1Qportvlanhcstatisticsentry) GetParentYangName() string { return "dot1qPortVlanHCStatisticsTable" }

// QBRIDGEMIB_Dot1Qlearningconstraintstable
// A table containing learning constraints for sets of
// Shared and Independent VLANs.
type QBRIDGEMIB_Dot1Qlearningconstraintstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A learning constraint defined for a VLAN. The type is slice of
    // QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry.
    Dot1Qlearningconstraintsentry []QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry
}

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetFilter() yfilter.YFilter { return dot1Qlearningconstraintstable.YFilter }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) SetFilter(yf yfilter.YFilter) { dot1Qlearningconstraintstable.YFilter = yf }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetGoName(yname string) string {
    if yname == "dot1qLearningConstraintsEntry" { return "Dot1Qlearningconstraintsentry" }
    return ""
}

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetSegmentPath() string {
    return "dot1qLearningConstraintsTable"
}

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1qLearningConstraintsEntry" {
        for _, c := range dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry {
            if dot1Qlearningconstraintstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry{}
        dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry = append(dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry, child)
        return &dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry[len(dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry)-1]
    }
    return nil
}

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry {
        children[dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry[i].GetSegmentPath()] = &dot1Qlearningconstraintstable.Dot1Qlearningconstraintsentry[i]
    }
    return children
}

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetYangName() string { return "dot1qLearningConstraintsTable" }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) SetParent(parent types.Entity) { dot1Qlearningconstraintstable.parent = parent }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetParent() types.Entity { return dot1Qlearningconstraintstable.parent }

func (dot1Qlearningconstraintstable *QBRIDGEMIB_Dot1Qlearningconstraintstable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry
// A learning constraint defined for a VLAN.
type QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry struct {
    parent types.Entity
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

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetFilter() yfilter.YFilter { return dot1Qlearningconstraintsentry.YFilter }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) SetFilter(yf yfilter.YFilter) { dot1Qlearningconstraintsentry.YFilter = yf }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetGoName(yname string) string {
    if yname == "dot1qConstraintVlan" { return "Dot1Qconstraintvlan" }
    if yname == "dot1qConstraintSet" { return "Dot1Qconstraintset" }
    if yname == "dot1qConstraintType" { return "Dot1Qconstrainttype" }
    if yname == "dot1qConstraintStatus" { return "Dot1Qconstraintstatus" }
    return ""
}

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetSegmentPath() string {
    return "dot1qLearningConstraintsEntry" + "[dot1qConstraintVlan='" + fmt.Sprintf("%v", dot1Qlearningconstraintsentry.Dot1Qconstraintvlan) + "']" + "[dot1qConstraintSet='" + fmt.Sprintf("%v", dot1Qlearningconstraintsentry.Dot1Qconstraintset) + "']"
}

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1qConstraintVlan"] = dot1Qlearningconstraintsentry.Dot1Qconstraintvlan
    leafs["dot1qConstraintSet"] = dot1Qlearningconstraintsentry.Dot1Qconstraintset
    leafs["dot1qConstraintType"] = dot1Qlearningconstraintsentry.Dot1Qconstrainttype
    leafs["dot1qConstraintStatus"] = dot1Qlearningconstraintsentry.Dot1Qconstraintstatus
    return leafs
}

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetYangName() string { return "dot1qLearningConstraintsEntry" }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) SetParent(parent types.Entity) { dot1Qlearningconstraintsentry.parent = parent }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetParent() types.Entity { return dot1Qlearningconstraintsentry.parent }

func (dot1Qlearningconstraintsentry *QBRIDGEMIB_Dot1Qlearningconstraintstable_Dot1Qlearningconstraintsentry) GetParentYangName() string { return "dot1qLearningConstraintsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A mapping from a Protocol Template to a Protocol Group Identifier. The type
    // is slice of QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry.
    Dot1Vprotocolgroupentry []QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry
}

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetFilter() yfilter.YFilter { return dot1Vprotocolgrouptable.YFilter }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) SetFilter(yf yfilter.YFilter) { dot1Vprotocolgrouptable.YFilter = yf }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetGoName(yname string) string {
    if yname == "dot1vProtocolGroupEntry" { return "Dot1Vprotocolgroupentry" }
    return ""
}

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetSegmentPath() string {
    return "dot1vProtocolGroupTable"
}

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1vProtocolGroupEntry" {
        for _, c := range dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry {
            if dot1Vprotocolgrouptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry{}
        dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry = append(dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry, child)
        return &dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry[len(dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry)-1]
    }
    return nil
}

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry {
        children[dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry[i].GetSegmentPath()] = &dot1Vprotocolgrouptable.Dot1Vprotocolgroupentry[i]
    }
    return children
}

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetYangName() string { return "dot1vProtocolGroupTable" }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) SetParent(parent types.Entity) { dot1Vprotocolgrouptable.parent = parent }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetParent() types.Entity { return dot1Vprotocolgrouptable.parent }

func (dot1Vprotocolgrouptable *QBRIDGEMIB_Dot1Vprotocolgrouptable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry
// A mapping from a Protocol Template to a Protocol
// Group Identifier.
type QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry struct {
    parent types.Entity
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

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetFilter() yfilter.YFilter { return dot1Vprotocolgroupentry.YFilter }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) SetFilter(yf yfilter.YFilter) { dot1Vprotocolgroupentry.YFilter = yf }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetGoName(yname string) string {
    if yname == "dot1vProtocolTemplateFrameType" { return "Dot1Vprotocoltemplateframetype" }
    if yname == "dot1vProtocolTemplateProtocolValue" { return "Dot1Vprotocoltemplateprotocolvalue" }
    if yname == "dot1vProtocolGroupId" { return "Dot1Vprotocolgroupid" }
    if yname == "dot1vProtocolGroupRowStatus" { return "Dot1Vprotocolgrouprowstatus" }
    return ""
}

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetSegmentPath() string {
    return "dot1vProtocolGroupEntry" + "[dot1vProtocolTemplateFrameType='" + fmt.Sprintf("%v", dot1Vprotocolgroupentry.Dot1Vprotocoltemplateframetype) + "']" + "[dot1vProtocolTemplateProtocolValue='" + fmt.Sprintf("%v", dot1Vprotocolgroupentry.Dot1Vprotocoltemplateprotocolvalue) + "']"
}

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1vProtocolTemplateFrameType"] = dot1Vprotocolgroupentry.Dot1Vprotocoltemplateframetype
    leafs["dot1vProtocolTemplateProtocolValue"] = dot1Vprotocolgroupentry.Dot1Vprotocoltemplateprotocolvalue
    leafs["dot1vProtocolGroupId"] = dot1Vprotocolgroupentry.Dot1Vprotocolgroupid
    leafs["dot1vProtocolGroupRowStatus"] = dot1Vprotocolgroupentry.Dot1Vprotocolgrouprowstatus
    return leafs
}

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetYangName() string { return "dot1vProtocolGroupEntry" }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) SetParent(parent types.Entity) { dot1Vprotocolgroupentry.parent = parent }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetParent() types.Entity { return dot1Vprotocolgroupentry.parent }

func (dot1Vprotocolgroupentry *QBRIDGEMIB_Dot1Vprotocolgrouptable_Dot1Vprotocolgroupentry) GetParentYangName() string { return "dot1vProtocolGroupTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A VID set for a port. The type is slice of
    // QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry.
    Dot1Vprotocolportentry []QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry
}

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetFilter() yfilter.YFilter { return dot1Vprotocolporttable.YFilter }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) SetFilter(yf yfilter.YFilter) { dot1Vprotocolporttable.YFilter = yf }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetGoName(yname string) string {
    if yname == "dot1vProtocolPortEntry" { return "Dot1Vprotocolportentry" }
    return ""
}

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetSegmentPath() string {
    return "dot1vProtocolPortTable"
}

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1vProtocolPortEntry" {
        for _, c := range dot1Vprotocolporttable.Dot1Vprotocolportentry {
            if dot1Vprotocolporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry{}
        dot1Vprotocolporttable.Dot1Vprotocolportentry = append(dot1Vprotocolporttable.Dot1Vprotocolportentry, child)
        return &dot1Vprotocolporttable.Dot1Vprotocolportentry[len(dot1Vprotocolporttable.Dot1Vprotocolportentry)-1]
    }
    return nil
}

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dot1Vprotocolporttable.Dot1Vprotocolportentry {
        children[dot1Vprotocolporttable.Dot1Vprotocolportentry[i].GetSegmentPath()] = &dot1Vprotocolporttable.Dot1Vprotocolportentry[i]
    }
    return children
}

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetYangName() string { return "dot1vProtocolPortTable" }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) SetParent(parent types.Entity) { dot1Vprotocolporttable.parent = parent }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetParent() types.Entity { return dot1Vprotocolporttable.parent }

func (dot1Vprotocolporttable *QBRIDGEMIB_Dot1Vprotocolporttable) GetParentYangName() string { return "Q-BRIDGE-MIB" }

// QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry
// A VID set for a port.
type QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry struct {
    parent types.Entity
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

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetFilter() yfilter.YFilter { return dot1Vprotocolportentry.YFilter }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) SetFilter(yf yfilter.YFilter) { dot1Vprotocolportentry.YFilter = yf }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetGoName(yname string) string {
    if yname == "dot1dBasePort" { return "Dot1Dbaseport" }
    if yname == "dot1vProtocolPortGroupId" { return "Dot1Vprotocolportgroupid" }
    if yname == "dot1vProtocolPortGroupVid" { return "Dot1Vprotocolportgroupvid" }
    if yname == "dot1vProtocolPortRowStatus" { return "Dot1Vprotocolportrowstatus" }
    return ""
}

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetSegmentPath() string {
    return "dot1vProtocolPortEntry" + "[dot1dBasePort='" + fmt.Sprintf("%v", dot1Vprotocolportentry.Dot1Dbaseport) + "']" + "[dot1vProtocolPortGroupId='" + fmt.Sprintf("%v", dot1Vprotocolportentry.Dot1Vprotocolportgroupid) + "']"
}

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot1dBasePort"] = dot1Vprotocolportentry.Dot1Dbaseport
    leafs["dot1vProtocolPortGroupId"] = dot1Vprotocolportentry.Dot1Vprotocolportgroupid
    leafs["dot1vProtocolPortGroupVid"] = dot1Vprotocolportentry.Dot1Vprotocolportgroupvid
    leafs["dot1vProtocolPortRowStatus"] = dot1Vprotocolportentry.Dot1Vprotocolportrowstatus
    return leafs
}

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetBundleName() string { return "cisco_ios_xe" }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetYangName() string { return "dot1vProtocolPortEntry" }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) SetParent(parent types.Entity) { dot1Vprotocolportentry.parent = parent }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetParent() types.Entity { return dot1Vprotocolportentry.parent }

func (dot1Vprotocolportentry *QBRIDGEMIB_Dot1Vprotocolporttable_Dot1Vprotocolportentry) GetParentYangName() string { return "dot1vProtocolPortTable" }

