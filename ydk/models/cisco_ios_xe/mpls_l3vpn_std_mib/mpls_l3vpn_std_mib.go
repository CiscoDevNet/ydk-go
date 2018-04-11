// This MIB contains managed object definitions for the
// Layer-3 Multiprotocol Label Switching Virtual
// Private Networks.
// 
// Copyright (C) The Internet Society (2006).  This
// version of this MIB module is part of RFC4382; see
// the RFC itself for full legal notices.
package mpls_l3vpn_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_l3vpn_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:MPLS-L3VPN-STD-MIB MPLS-L3VPN-STD-MIB}", reflect.TypeOf(MPLSL3VPNSTDMIB{}))
    ydk.RegisterEntity("MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB", reflect.TypeOf(MPLSL3VPNSTDMIB{}))
}

// MplsL3VpnRtType represents route target, see [RFC4364].
type MplsL3VpnRtType string

const (
    MplsL3VpnRtType_import_ MplsL3VpnRtType = "import"

    MplsL3VpnRtType_export MplsL3VpnRtType = "export"

    MplsL3VpnRtType_both MplsL3VpnRtType = "both"
)

// MPLSL3VPNSTDMIB
type MPLSL3VPNSTDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Mplsl3Vpnscalars MPLSL3VPNSTDMIB_Mplsl3Vpnscalars

    // This table specifies per-interface MPLS capability and associated
    // information.
    Mplsl3Vpnifconftable MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable

    // This table specifies per-interface MPLS L3VPN VRF Table capability and
    // associated information. Entries in this table define VRF routing instances
    // associated with MPLS/VPN interfaces.  Note that multiple interfaces can
    // belong to the same VRF instance.  The collection of all VRF instances
    // comprises an actual VPN.
    Mplsl3Vpnvrftable MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable

    // This table specifies per-VRF route target association. Each entry
    // identifies a connectivity policy supported as part of a VPN.
    Mplsl3Vpnvrfrttable MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable

    // This table specifies per-interface MPLS L3VPN VRF Table routing
    // information.  Entries in this table define VRF routing entries associated
    // with the specified MPLS/VPN interfaces.  Note  that this table contains
    // both BGP and Interior Gateway Protocol IGP routes, as both may appear in
    // the same VRF.
    Mplsl3Vpnvrfrtetable MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable
}

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSL3VPNSTDMIB.EntityData.YFilter = mPLSL3VPNSTDMIB.YFilter
    mPLSL3VPNSTDMIB.EntityData.YangName = "MPLS-L3VPN-STD-MIB"
    mPLSL3VPNSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSL3VPNSTDMIB.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mPLSL3VPNSTDMIB.EntityData.SegmentPath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB"
    mPLSL3VPNSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSL3VPNSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSL3VPNSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSL3VPNSTDMIB.EntityData.Children = make(map[string]types.YChild)
    mPLSL3VPNSTDMIB.EntityData.Children["mplsL3VpnScalars"] = types.YChild{"Mplsl3Vpnscalars", &mPLSL3VPNSTDMIB.Mplsl3Vpnscalars}
    mPLSL3VPNSTDMIB.EntityData.Children["mplsL3VpnIfConfTable"] = types.YChild{"Mplsl3Vpnifconftable", &mPLSL3VPNSTDMIB.Mplsl3Vpnifconftable}
    mPLSL3VPNSTDMIB.EntityData.Children["mplsL3VpnVrfTable"] = types.YChild{"Mplsl3Vpnvrftable", &mPLSL3VPNSTDMIB.Mplsl3Vpnvrftable}
    mPLSL3VPNSTDMIB.EntityData.Children["mplsL3VpnVrfRTTable"] = types.YChild{"Mplsl3Vpnvrfrttable", &mPLSL3VPNSTDMIB.Mplsl3Vpnvrfrttable}
    mPLSL3VPNSTDMIB.EntityData.Children["mplsL3VpnVrfRteTable"] = types.YChild{"Mplsl3Vpnvrfrtetable", &mPLSL3VPNSTDMIB.Mplsl3Vpnvrfrtetable}
    mPLSL3VPNSTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mPLSL3VPNSTDMIB.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnscalars
type MPLSL3VPNSTDMIB_Mplsl3Vpnscalars struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of VRFs that are configured on this node. The type is
    // interface{} with range: 0..4294967295.
    Mplsl3Vpnconfiguredvrfs interface{}

    // The number of VRFs that are active on this node. That is, those VRFs whose
    // corresponding mplsL3VpnVrfOperStatus object value is equal to operational
    // (1). The type is interface{} with range: 0..4294967295.
    Mplsl3Vpnactivevrfs interface{}

    // Total number of interfaces connected to a VRF. The type is interface{} with
    // range: 0..4294967295.
    Mplsl3Vpnconnectedinterfaces interface{}

    // If this object is true, then it enables the generation of all notifications
    // defined in this MIB.  This object's value should be preserved across agent
    // reboots. The type is bool.
    Mplsl3Vpnnotificationenable interface{}

    // Denotes maximum number of routes that the device will allow all VRFs
    // jointly to hold.  If this value is set to 0, this indicates that the device
    // is unable to determine the absolute maximum.  In this case, the configured
    // maximum MAY not actually be allowed by the device. The type is interface{}
    // with range: 0..4294967295.
    Mplsl3Vpnvrfconfmaxpossrts interface{}

    // Denotes the interval in seconds, at which the route max threshold
    // notification may be reissued after the maximum value has been exceeded (or
    // has been reached if mplsL3VpnVrfConfMaxRoutes and
    // mplsL3VpnVrfConfHighRteThresh are equal) and the initial notification has
    // been issued.  This value is intended to prevent continuous generation of
    // notifications by an agent in the event that routes are continually added to
    // a VRF after it has reached its maximum value.  If this value is set to 0,
    // the agent should only issue a single notification at the time that the
    // maximum threshold has been reached, and should not issue any more
    // notifications until the value of routes has fallen below the configured
    // threshold value.  This is the recommended default behavior. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Mplsl3Vpnvrfconfrtemxthrshtime interface{}

    // The number of illegally received labels above which the
    // mplsNumVrfSecIllglLblThrshExcd notification is issued.  The persistence of
    // this value mimics that of the device's configuration. The type is
    // interface{} with range: 0..4294967295.
    Mplsl3Vpnilllblrcvthrsh interface{}
}

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnscalars.EntityData.YFilter = mplsl3Vpnscalars.YFilter
    mplsl3Vpnscalars.EntityData.YangName = "mplsL3VpnScalars"
    mplsl3Vpnscalars.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnscalars.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsl3Vpnscalars.EntityData.SegmentPath = "mplsL3VpnScalars"
    mplsl3Vpnscalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnscalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnscalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnscalars.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnscalars.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsl3Vpnscalars.EntityData.Leafs["mplsL3VpnConfiguredVrfs"] = types.YLeaf{"Mplsl3Vpnconfiguredvrfs", mplsl3Vpnscalars.Mplsl3Vpnconfiguredvrfs}
    mplsl3Vpnscalars.EntityData.Leafs["mplsL3VpnActiveVrfs"] = types.YLeaf{"Mplsl3Vpnactivevrfs", mplsl3Vpnscalars.Mplsl3Vpnactivevrfs}
    mplsl3Vpnscalars.EntityData.Leafs["mplsL3VpnConnectedInterfaces"] = types.YLeaf{"Mplsl3Vpnconnectedinterfaces", mplsl3Vpnscalars.Mplsl3Vpnconnectedinterfaces}
    mplsl3Vpnscalars.EntityData.Leafs["mplsL3VpnNotificationEnable"] = types.YLeaf{"Mplsl3Vpnnotificationenable", mplsl3Vpnscalars.Mplsl3Vpnnotificationenable}
    mplsl3Vpnscalars.EntityData.Leafs["mplsL3VpnVrfConfMaxPossRts"] = types.YLeaf{"Mplsl3Vpnvrfconfmaxpossrts", mplsl3Vpnscalars.Mplsl3Vpnvrfconfmaxpossrts}
    mplsl3Vpnscalars.EntityData.Leafs["mplsL3VpnVrfConfRteMxThrshTime"] = types.YLeaf{"Mplsl3Vpnvrfconfrtemxthrshtime", mplsl3Vpnscalars.Mplsl3Vpnvrfconfrtemxthrshtime}
    mplsl3Vpnscalars.EntityData.Leafs["mplsL3VpnIllLblRcvThrsh"] = types.YLeaf{"Mplsl3Vpnilllblrcvthrsh", mplsl3Vpnscalars.Mplsl3Vpnilllblrcvthrsh}
    return &(mplsl3Vpnscalars.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every interface capable of
    // supporting MPLS L3VPN. Each entry in this table is meant to correspond to
    // an entry in the Interfaces Table. The type is slice of
    // MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry.
    Mplsl3Vpnifconfentry []MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry
}

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnifconftable.EntityData.YFilter = mplsl3Vpnifconftable.YFilter
    mplsl3Vpnifconftable.EntityData.YangName = "mplsL3VpnIfConfTable"
    mplsl3Vpnifconftable.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnifconftable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsl3Vpnifconftable.EntityData.SegmentPath = "mplsL3VpnIfConfTable"
    mplsl3Vpnifconftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnifconftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnifconftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnifconftable.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnifconftable.EntityData.Children["mplsL3VpnIfConfEntry"] = types.YChild{"Mplsl3Vpnifconfentry", nil}
    for i := range mplsl3Vpnifconftable.Mplsl3Vpnifconfentry {
        mplsl3Vpnifconftable.EntityData.Children[types.GetSegmentPath(&mplsl3Vpnifconftable.Mplsl3Vpnifconfentry[i])] = types.YChild{"Mplsl3Vpnifconfentry", &mplsl3Vpnifconftable.Mplsl3Vpnifconfentry[i]}
    }
    mplsl3Vpnifconftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsl3Vpnifconftable.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry
// An entry in this table is created by an LSR for
// every interface capable of supporting MPLS L3VPN.
// Each entry in this table is meant to correspond to
// an entry in the Interfaces Table.
type MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_l3vpn_std_mib.MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfname
    Mplsl3Vpnvrfname interface{}

    // This attribute is a key. This is a unique index for an entry in the
    // mplsL3VpnIfConfTable.  A non-zero index for an entry indicates the ifIndex
    // for the corresponding interface entry in the MPLS-VPN-layer in the ifTable.
    // Note that this table does not necessarily correspond one-to-one with all
    // entries in the Interface MIB having an ifType of MPLS-layer; rather, only
    // those that are enabled for MPLS L3VPN functionality. The type is
    // interface{} with range: 1..2147483647.
    Mplsl3Vpnifconfindex interface{}

    // Denotes whether this link participates in a carrier's carrier, enterprise,
    // or inter-provider scenario. The type is Mplsl3Vpnifvpnclassification.
    Mplsl3Vpnifvpnclassification interface{}

    // Denotes the route distribution protocol across the PE-CE link.  Note that
    // more than one routing protocol may be enabled at the same time; thus, this
    // object is specified as a bitmask.  For example, static(5) and ospf(2) are a
    // typical configuration. The type is map[string]bool.
    Mplsl3Vpnifvpnroutedistprotocol interface{}

    // The storage type for this VPN If entry. Conceptual rows having the value
    // 'permanent' need not allow write access to any columnar objects in the row.
    // The type is StorageType.
    Mplsl3Vpnifconfstoragetype interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // Rows in this table signify that the specified interface is associated with
    // this VRF.  If the row creation operation succeeds, the interface will have
    // been associated with the specified VRF, otherwise the agent MUST not allow
    // the association.  If the agent only allows read-only operations on this
    // table, it MUST create entries in this table as they are created on the
    // device.  When a row in this table is in active(1) state, no objects in that
    // row can be modified except mplsL3VpnIfConfStorageType and
    // mplsL3VpnIfConfRowStatus. The type is RowStatus.
    Mplsl3Vpnifconfrowstatus interface{}
}

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnifconfentry.EntityData.YFilter = mplsl3Vpnifconfentry.YFilter
    mplsl3Vpnifconfentry.EntityData.YangName = "mplsL3VpnIfConfEntry"
    mplsl3Vpnifconfentry.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnifconfentry.EntityData.ParentYangName = "mplsL3VpnIfConfTable"
    mplsl3Vpnifconfentry.EntityData.SegmentPath = "mplsL3VpnIfConfEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnifconfentry.Mplsl3Vpnvrfname) + "']" + "[mplsL3VpnIfConfIndex='" + fmt.Sprintf("%v", mplsl3Vpnifconfentry.Mplsl3Vpnifconfindex) + "']"
    mplsl3Vpnifconfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnifconfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnifconfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnifconfentry.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnifconfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsl3Vpnifconfentry.EntityData.Leafs["mplsL3VpnVrfName"] = types.YLeaf{"Mplsl3Vpnvrfname", mplsl3Vpnifconfentry.Mplsl3Vpnvrfname}
    mplsl3Vpnifconfentry.EntityData.Leafs["mplsL3VpnIfConfIndex"] = types.YLeaf{"Mplsl3Vpnifconfindex", mplsl3Vpnifconfentry.Mplsl3Vpnifconfindex}
    mplsl3Vpnifconfentry.EntityData.Leafs["mplsL3VpnIfVpnClassification"] = types.YLeaf{"Mplsl3Vpnifvpnclassification", mplsl3Vpnifconfentry.Mplsl3Vpnifvpnclassification}
    mplsl3Vpnifconfentry.EntityData.Leafs["mplsL3VpnIfVpnRouteDistProtocol"] = types.YLeaf{"Mplsl3Vpnifvpnroutedistprotocol", mplsl3Vpnifconfentry.Mplsl3Vpnifvpnroutedistprotocol}
    mplsl3Vpnifconfentry.EntityData.Leafs["mplsL3VpnIfConfStorageType"] = types.YLeaf{"Mplsl3Vpnifconfstoragetype", mplsl3Vpnifconfentry.Mplsl3Vpnifconfstoragetype}
    mplsl3Vpnifconfentry.EntityData.Leafs["mplsL3VpnIfConfRowStatus"] = types.YLeaf{"Mplsl3Vpnifconfrowstatus", mplsl3Vpnifconfentry.Mplsl3Vpnifconfrowstatus}
    return &(mplsl3Vpnifconfentry.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification represents scenario.
type MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification string

const (
    MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification_carrierOfCarrier MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification = "carrierOfCarrier"

    MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification_enterprise MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification = "enterprise"

    MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification_interProvider MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry_Mplsl3Vpnifvpnclassification = "interProvider"
)

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable
// This table specifies per-interface MPLS L3VPN
// VRF Table capability and associated information.
// Entries in this table define VRF routing instances
// associated with MPLS/VPN interfaces.  Note that
// multiple interfaces can belong to the same VRF
// instance.  The collection of all VRF instances
// comprises an actual VPN.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS L3VPN.  The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of
    // MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry.
    Mplsl3Vpnvrfentry []MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry
}

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnvrftable.EntityData.YFilter = mplsl3Vpnvrftable.YFilter
    mplsl3Vpnvrftable.EntityData.YangName = "mplsL3VpnVrfTable"
    mplsl3Vpnvrftable.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnvrftable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsl3Vpnvrftable.EntityData.SegmentPath = "mplsL3VpnVrfTable"
    mplsl3Vpnvrftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnvrftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnvrftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnvrftable.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnvrftable.EntityData.Children["mplsL3VpnVrfEntry"] = types.YChild{"Mplsl3Vpnvrfentry", nil}
    for i := range mplsl3Vpnvrftable.Mplsl3Vpnvrfentry {
        mplsl3Vpnvrftable.EntityData.Children[types.GetSegmentPath(&mplsl3Vpnvrftable.Mplsl3Vpnvrfentry[i])] = types.YChild{"Mplsl3Vpnvrfentry", &mplsl3Vpnvrftable.Mplsl3Vpnvrfentry[i]}
    }
    mplsl3Vpnvrftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsl3Vpnvrftable.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS L3VPN.  The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The human-readable name of this VPN.  This MAY be
    // equivalent to the [RFC2685] VPN-ID, but may also vary.  If it is set to the
    // VPN ID, it MUST be equivalent to the value of mplsL3VpnVrfVpnId. It is
    // strongly recommended that all sites supporting VRFs that are part of the
    // same VPN use the same naming convention for VRFs as well as the same VPN
    // ID. The type is string with length: 0..31.
    Mplsl3Vpnvrfname interface{}

    // The VPN ID as specified in [RFC2685].  If a VPN ID has not been specified
    // for this VRF, then this variable SHOULD be set to a zero-length OCTET
    // STRING. The type is string with length: 0 | 7.
    Mplsl3Vpnvrfvpnid interface{}

    // The human-readable description of this VRF. The type is string.
    Mplsl3Vpnvrfdescription interface{}

    // The route distinguisher for this VRF. The type is string with length:
    // 0..256.
    Mplsl3Vpnvrfrd interface{}

    // The time at which this VRF entry was created. The type is interface{} with
    // range: 0..4294967295.
    Mplsl3Vpnvrfcreationtime interface{}

    // Denotes whether or not a VRF is operational.  A VRF is up(1) when there is
    // at least one interface associated with the VRF whose ifOperStatus is up(1).
    // A VRF is down(2) when: a. There does not exist at least one interface whose
    // ifOperStatus is up(1). b. There are no interfaces associated with the VRF.
    // The type is Mplsl3Vpnvrfoperstatus.
    Mplsl3Vpnvrfoperstatus interface{}

    // Total number of interfaces connected to this VRF with ifOperStatus = up(1).
    // This value should increase when an interface is associated with the
    // corresponding VRF and its corresponding ifOperStatus is equal to up(1).  If
    // an interface is associated whose ifOperStatus is not up(1), then the value
    // is not incremented until such time as it transitions to this state.  This
    // value should be decremented when an interface is disassociated with a VRF
    // or the corresponding ifOperStatus transitions out of the up(1) state to any
    // other state. The type is interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfactiveinterfaces interface{}

    // Total number of interfaces connected to this VRF (independent of
    // ifOperStatus type). The type is interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfassociatedinterfaces interface{}

    // Denotes mid-level water marker for the number of routes that this VRF may
    // hold. The type is interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfconfmidrtethresh interface{}

    // Denotes high-level water marker for the number of routes that this VRF may
    // hold. The type is interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfconfhighrtethresh interface{}

    // Denotes maximum number of routes that this VRF is configured to hold.  This
    // value MUST be less than or equal to mplsL3VpnVrfConfMaxPossRts unless it is
    // set to 0. The type is interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfconfmaxroutes interface{}

    // The value of sysUpTime at the time of the last change of this table entry,
    // which includes changes of VRF parameters defined in this table or addition
    // or deletion of interfaces associated with this VRF. The type is interface{}
    // with range: 0..4294967295.
    Mplsl3Vpnvrfconflastchanged interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified except mplsL3VpnVrfConfAdminStatus, mplsL3VpnVrfConfRowStatus,
    // and mplsL3VpnVrfConfStorageType. The type is RowStatus.
    Mplsl3Vpnvrfconfrowstatus interface{}

    // Indicates the desired operational status of this VRF. The type is
    // Mplsl3Vpnvrfconfadminstatus.
    Mplsl3Vpnvrfconfadminstatus interface{}

    // The storage type for this VPN VRF entry. Conceptual rows having the value
    // 'permanent' need not allow write access to any columnar objects in the row.
    // The type is StorageType.
    Mplsl3Vpnvrfconfstoragetype interface{}

    // Indicates the number of illegally received labels on this VPN/VRF. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsL3VpnVrfSecDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsl3Vpnvrfsecillegallblvltns interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfsecdiscontinuitytime interface{}

    // Indicates the number of routes added to this VPN/VRF since the last
    // discontinuity.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of mplsL3VpnVrfPerfDiscTime. The type is interface{} with
    // range: 0..4294967295.
    Mplsl3Vpnvrfperfroutesadded interface{}

    // Indicates the number of routes removed from this VPN/VRF.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // mplsL3VpnVrfPerfDiscTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsl3Vpnvrfperfroutesdeleted interface{}

    // Indicates the number of routes currently used by this VRF. The type is
    // interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfperfcurrnumroutes interface{}

    // This counter should be incremented when the number of routes contained by
    // the specified VRF exceeds or attempts to exceed the maximum allowed value
    // as indicated by mplsL3VpnVrfMaxRouteThreshold.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // mplsL3VpnVrfPerfDiscTime. The type is interface{} with range:
    // 0..4294967295.
    Mplsl3Vpnvrfperfroutesdropped interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfperfdisctime interface{}
}

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnvrfentry.EntityData.YFilter = mplsl3Vpnvrfentry.YFilter
    mplsl3Vpnvrfentry.EntityData.YangName = "mplsL3VpnVrfEntry"
    mplsl3Vpnvrfentry.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnvrfentry.EntityData.ParentYangName = "mplsL3VpnVrfTable"
    mplsl3Vpnvrfentry.EntityData.SegmentPath = "mplsL3VpnVrfEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnvrfentry.Mplsl3Vpnvrfname) + "']"
    mplsl3Vpnvrfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnvrfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnvrfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnvrfentry.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnvrfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfName"] = types.YLeaf{"Mplsl3Vpnvrfname", mplsl3Vpnvrfentry.Mplsl3Vpnvrfname}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfVpnId"] = types.YLeaf{"Mplsl3Vpnvrfvpnid", mplsl3Vpnvrfentry.Mplsl3Vpnvrfvpnid}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfDescription"] = types.YLeaf{"Mplsl3Vpnvrfdescription", mplsl3Vpnvrfentry.Mplsl3Vpnvrfdescription}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfRD"] = types.YLeaf{"Mplsl3Vpnvrfrd", mplsl3Vpnvrfentry.Mplsl3Vpnvrfrd}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfCreationTime"] = types.YLeaf{"Mplsl3Vpnvrfcreationtime", mplsl3Vpnvrfentry.Mplsl3Vpnvrfcreationtime}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfOperStatus"] = types.YLeaf{"Mplsl3Vpnvrfoperstatus", mplsl3Vpnvrfentry.Mplsl3Vpnvrfoperstatus}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfActiveInterfaces"] = types.YLeaf{"Mplsl3Vpnvrfactiveinterfaces", mplsl3Vpnvrfentry.Mplsl3Vpnvrfactiveinterfaces}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfAssociatedInterfaces"] = types.YLeaf{"Mplsl3Vpnvrfassociatedinterfaces", mplsl3Vpnvrfentry.Mplsl3Vpnvrfassociatedinterfaces}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfConfMidRteThresh"] = types.YLeaf{"Mplsl3Vpnvrfconfmidrtethresh", mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfmidrtethresh}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfConfHighRteThresh"] = types.YLeaf{"Mplsl3Vpnvrfconfhighrtethresh", mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfhighrtethresh}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfConfMaxRoutes"] = types.YLeaf{"Mplsl3Vpnvrfconfmaxroutes", mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfmaxroutes}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfConfLastChanged"] = types.YLeaf{"Mplsl3Vpnvrfconflastchanged", mplsl3Vpnvrfentry.Mplsl3Vpnvrfconflastchanged}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfConfRowStatus"] = types.YLeaf{"Mplsl3Vpnvrfconfrowstatus", mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfrowstatus}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfConfAdminStatus"] = types.YLeaf{"Mplsl3Vpnvrfconfadminstatus", mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfadminstatus}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfConfStorageType"] = types.YLeaf{"Mplsl3Vpnvrfconfstoragetype", mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfstoragetype}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfSecIllegalLblVltns"] = types.YLeaf{"Mplsl3Vpnvrfsecillegallblvltns", mplsl3Vpnvrfentry.Mplsl3Vpnvrfsecillegallblvltns}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfSecDiscontinuityTime"] = types.YLeaf{"Mplsl3Vpnvrfsecdiscontinuitytime", mplsl3Vpnvrfentry.Mplsl3Vpnvrfsecdiscontinuitytime}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfPerfRoutesAdded"] = types.YLeaf{"Mplsl3Vpnvrfperfroutesadded", mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfroutesadded}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfPerfRoutesDeleted"] = types.YLeaf{"Mplsl3Vpnvrfperfroutesdeleted", mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfroutesdeleted}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfPerfCurrNumRoutes"] = types.YLeaf{"Mplsl3Vpnvrfperfcurrnumroutes", mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfcurrnumroutes}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfPerfRoutesDropped"] = types.YLeaf{"Mplsl3Vpnvrfperfroutesdropped", mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfroutesdropped}
    mplsl3Vpnvrfentry.EntityData.Leafs["mplsL3VpnVrfPerfDiscTime"] = types.YLeaf{"Mplsl3Vpnvrfperfdisctime", mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfdisctime}
    return &(mplsl3Vpnvrfentry.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus represents VRF.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus string

const (
    MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus_up MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus = "up"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus_down MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus = "down"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus_testing MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfconfadminstatus = "testing"
)

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfoperstatus represents b. There are no interfaces associated with the VRF.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfoperstatus string

const (
    MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfoperstatus_up MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfoperstatus = "up"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfoperstatus_down MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfoperstatus = "down"
)

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable
// This table specifies per-VRF route target association.
// Each entry identifies a connectivity policy supported
// as part of a VPN.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for each route target
    // configured for a VRF supporting a MPLS L3VPN instance.  The indexing
    // provides an ordering per-VRF instance.  See [RFC4364] for a complete
    // definition of a route target. The type is slice of
    // MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry.
    Mplsl3Vpnvrfrtentry []MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry
}

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnvrfrttable.EntityData.YFilter = mplsl3Vpnvrfrttable.YFilter
    mplsl3Vpnvrfrttable.EntityData.YangName = "mplsL3VpnVrfRTTable"
    mplsl3Vpnvrfrttable.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnvrfrttable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsl3Vpnvrfrttable.EntityData.SegmentPath = "mplsL3VpnVrfRTTable"
    mplsl3Vpnvrfrttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnvrfrttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnvrfrttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnvrfrttable.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnvrfrttable.EntityData.Children["mplsL3VpnVrfRTEntry"] = types.YChild{"Mplsl3Vpnvrfrtentry", nil}
    for i := range mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry {
        mplsl3Vpnvrfrttable.EntityData.Children[types.GetSegmentPath(&mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry[i])] = types.YChild{"Mplsl3Vpnvrfrtentry", &mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry[i]}
    }
    mplsl3Vpnvrfrttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsl3Vpnvrfrttable.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry
// An entry in this table is created by an LSR for
// each route target configured for a VRF supporting
// a MPLS L3VPN instance.  The indexing provides an
// ordering per-VRF instance.  See [RFC4364] for a
// complete definition of a route target.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_l3vpn_std_mib.MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfname
    Mplsl3Vpnvrfname interface{}

    // This attribute is a key. Auxiliary index for route targets configured for a
    // particular VRF. The type is interface{} with range: 1..4294967295.
    Mplsl3Vpnvrfrtindex interface{}

    // This attribute is a key. The route target distribution type. The type is
    // MplsL3VpnRtType.
    Mplsl3Vpnvrfrttype interface{}

    // The route target distribution policy. The type is string with length:
    // 0..256.
    Mplsl3Vpnvrfrt interface{}

    // Description of the route target. The type is string.
    Mplsl3Vpnvrfrtdescr interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified except mplsL3VpnVrfRTRowStatus. The type is RowStatus.
    Mplsl3Vpnvrfrtrowstatus interface{}

    // The storage type for this VPN route target (RT) entry. Conceptual rows
    // having the value 'permanent' need not allow write access to any columnar
    // objects in the row. The type is StorageType.
    Mplsl3Vpnvrfrtstoragetype interface{}
}

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnvrfrtentry.EntityData.YFilter = mplsl3Vpnvrfrtentry.YFilter
    mplsl3Vpnvrfrtentry.EntityData.YangName = "mplsL3VpnVrfRTEntry"
    mplsl3Vpnvrfrtentry.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnvrfrtentry.EntityData.ParentYangName = "mplsL3VpnVrfRTTable"
    mplsl3Vpnvrfrtentry.EntityData.SegmentPath = "mplsL3VpnVrfRTEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfname) + "']" + "[mplsL3VpnVrfRTIndex='" + fmt.Sprintf("%v", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtindex) + "']" + "[mplsL3VpnVrfRTType='" + fmt.Sprintf("%v", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrttype) + "']"
    mplsl3Vpnvrfrtentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnvrfrtentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnvrfrtentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnvrfrtentry.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnvrfrtentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsl3Vpnvrfrtentry.EntityData.Leafs["mplsL3VpnVrfName"] = types.YLeaf{"Mplsl3Vpnvrfname", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfname}
    mplsl3Vpnvrfrtentry.EntityData.Leafs["mplsL3VpnVrfRTIndex"] = types.YLeaf{"Mplsl3Vpnvrfrtindex", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtindex}
    mplsl3Vpnvrfrtentry.EntityData.Leafs["mplsL3VpnVrfRTType"] = types.YLeaf{"Mplsl3Vpnvrfrttype", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrttype}
    mplsl3Vpnvrfrtentry.EntityData.Leafs["mplsL3VpnVrfRT"] = types.YLeaf{"Mplsl3Vpnvrfrt", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrt}
    mplsl3Vpnvrfrtentry.EntityData.Leafs["mplsL3VpnVrfRTDescr"] = types.YLeaf{"Mplsl3Vpnvrfrtdescr", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtdescr}
    mplsl3Vpnvrfrtentry.EntityData.Leafs["mplsL3VpnVrfRTRowStatus"] = types.YLeaf{"Mplsl3Vpnvrfrtrowstatus", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtrowstatus}
    mplsl3Vpnvrfrtentry.EntityData.Leafs["mplsL3VpnVrfRTStorageType"] = types.YLeaf{"Mplsl3Vpnvrfrtstoragetype", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtstoragetype}
    return &(mplsl3Vpnvrfrtentry.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable
// This table specifies per-interface MPLS L3VPN VRF Table
// routing information.  Entries in this table define VRF routing
// entries associated with the specified MPLS/VPN interfaces.  Note
// 
// that this table contains both BGP and Interior Gateway Protocol
// IGP routes, as both may appear in the same VRF.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every route present
    // configured (either dynamically or statically) within the context of a
    // specific VRF capable of supporting MPLS/BGP VPN.  The indexing provides an
    // ordering of VRFs per-VPN interface.  Implementers need to be aware that
    // there are quite a few index objects that together can exceed the size
    // allowed for an Object Identifier (OID).  So implementers must make sure
    // that OIDs of column instances in this table will have no more than 128
    // sub-identifiers, otherwise they cannot be accessed using SNMPv1, SNMPv2c,
    // or SNMPv3. The type is slice of
    // MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry.
    Mplsl3Vpnvrfrteentry []MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry
}

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnvrfrtetable.EntityData.YFilter = mplsl3Vpnvrfrtetable.YFilter
    mplsl3Vpnvrfrtetable.EntityData.YangName = "mplsL3VpnVrfRteTable"
    mplsl3Vpnvrfrtetable.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnvrfrtetable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsl3Vpnvrfrtetable.EntityData.SegmentPath = "mplsL3VpnVrfRteTable"
    mplsl3Vpnvrfrtetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnvrfrtetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnvrfrtetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnvrfrtetable.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnvrfrtetable.EntityData.Children["mplsL3VpnVrfRteEntry"] = types.YChild{"Mplsl3Vpnvrfrteentry", nil}
    for i := range mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry {
        mplsl3Vpnvrfrtetable.EntityData.Children[types.GetSegmentPath(&mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry[i])] = types.YChild{"Mplsl3Vpnvrfrteentry", &mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry[i]}
    }
    mplsl3Vpnvrfrtetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsl3Vpnvrfrtetable.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry
// An entry in this table is created by an LSR for every route
// present configured (either dynamically or statically) within
// the context of a specific VRF capable of supporting MPLS/BGP
// VPN.  The indexing provides an ordering of VRFs per-VPN
// interface.
// 
// Implementers need to be aware that there are quite a few
// index objects that together can exceed the size allowed
// for an Object Identifier (OID).  So implementers must make
// sure that OIDs of column instances in this table will have
// no more than 128 sub-identifiers, otherwise they cannot be
// accessed using SNMPv1, SNMPv2c, or SNMPv3.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_l3vpn_std_mib.MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry_Mplsl3Vpnvrfname
    Mplsl3Vpnvrfname interface{}

    // This attribute is a key. The type of the mplsL3VpnVrfRteInetCidrDest
    // address, as defined in the InetAddress MIB.  Only those address types that
    // may appear in an actual routing table are allowed as values of this object.
    // The type is InetAddressType.
    Mplsl3Vpnvrfrteinetcidrdesttype interface{}

    // This attribute is a key. The destination IP address of this route.  The
    // type of this address is determined by the value of the
    // mplsL3VpnVrfRteInetCidrDestType object.  The values for the index objects
    // mplsL3VpnVrfRteInetCidrDest and mplsL3VpnVrfRteInetCidrPfxLen must be
    // consistent.  When the value of mplsL3VpnVrfRteInetCidrDest is x, then the
    // bitwise logical-AND of x with the value of the mask formed from the
    // corresponding index object mplsL3VpnVrfRteInetCidrPfxLen MUST be equal to
    // x.  If not, then the index pair is not consistent and an inconsistentName
    // error must be returned on SET or CREATE requests. The type is string with
    // length: 0..255.
    Mplsl3Vpnvrfrteinetcidrdest interface{}

    // This attribute is a key. Indicates the number of leading one bits that form
    // the  mask to be logical-ANDed with the destination address before being
    // compared to the value in the mplsL3VpnVrfRteInetCidrDest field.  The values
    // for the index objects mplsL3VpnVrfRteInetCidrDest and
    // mplsL3VpnVrfRteInetCidrPfxLen must be consistent.  When the value of
    // mplsL3VpnVrfRteInetCidrDest is x, then the bitwise logical-AND of x with
    // the value of the mask formed from the corresponding index object
    // mplsL3VpnVrfRteInetCidrPfxLen MUST be equal to x.  If not, then the index
    // pair is not consistent and an inconsistentName error must be returned on
    // SET or CREATE requests. The type is interface{} with range: 0..128.
    Mplsl3Vpnvrfrteinetcidrpfxlen interface{}

    // This attribute is a key. This object is an opaque object without any
    // defined semantics.  Its purpose is to serve as an additional index that may
    // delineate between multiple entries to the same destination.  The value { 0
    // 0 } shall be used as the default value for this object. The type is string
    // with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplsl3Vpnvrfrteinetcidrpolicy interface{}

    // This attribute is a key. The type of the mplsL3VpnVrfRteInetCidrNextHop
    // address, as defined in the InetAddress MIB.  Value should be set to
    // unknown(0) for non-remote routes.  Only those address types that may appear
    // in an actual routing table are allowed as values of this object. The type
    // is InetAddressType.
    Mplsl3Vpnvrfrteinetcidrnhoptype interface{}

    // This attribute is a key. On remote routes, the address of the next system
    // en route.  For non-remote routes, a zero-length string. The type of this
    // address is determined by the value of the mplsL3VpnVrfRteInetCidrNHopType
    // object. The type is string with length: 0..255.
    Mplsl3Vpnvrfrteinetcidrnexthop interface{}

    // The ifIndex value that identifies the local interface through which the
    // next hop of this route should be reached.  A value of 0 is valid and
    // represents the scenario where no interface is specified. The type is
    // interface{} with range: 0..2147483647.
    Mplsl3Vpnvrfrteinetcidrifindex interface{}

    // The type of route.  Note that local(3) refers to a route for which the next
    // hop is the final destination; remote(4) refers to a route for which the
    // next hop is not the final destination.  Routes that do not result in
    // traffic forwarding or rejection should not be displayed even if the
    // implementation keeps them stored internally.  reject(2) refers to a route
    // that, if matched, discards the message as unreachable and returns a
    // notification (e.g., ICMP error) to the message sender.  This is used in
    // some protocols as a means of correctly aggregating routes.  blackhole(5)
    // refers to a route that, if matched, discards the message silently. The type
    // is Mplsl3Vpnvrfrteinetcidrtype.
    Mplsl3Vpnvrfrteinetcidrtype interface{}

    // The routing mechanism via which this route was learned. Inclusion of values
    // for gateway routing protocols is not intended to imply that hosts should
    // support those protocols. The type is IANAipRouteProtocol.
    Mplsl3Vpnvrfrteinetcidrproto interface{}

    // The number of seconds since this route was last updated or otherwise
    // determined to be correct.  Note that no semantics of 'too old' can be
    // implied except through knowledge of the routing protocol by which the route
    // was learned. The type is interface{} with range: 0..4294967295.
    Mplsl3Vpnvrfrteinetcidrage interface{}

    // The Autonomous System Number of the next hop.  The semantics of this object
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  When this object is unknown or not
    // relevant, its value should be set to zero. The type is interface{} with
    // range: 0..4294967295.
    Mplsl3Vpnvrfrteinetcidrnexthopas interface{}

    // The primary routing metric for this route.  The semantics of this metric
    // are determined by the  routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    Mplsl3Vpnvrfrteinetcidrmetric1 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    Mplsl3Vpnvrfrteinetcidrmetric2 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    Mplsl3Vpnvrfrteinetcidrmetric3 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    Mplsl3Vpnvrfrteinetcidrmetric4 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    Mplsl3Vpnvrfrteinetcidrmetric5 interface{}

    // Index into mplsXCTable that identifies which cross- connect entry is
    // associated with this VRF route entry by containing the mplsXCIndex of that
    // cross-connect entry. The string containing the single-octet 0x00 indicates
    // that a label stack is not associated with this route entry.  This can be
    // the case because the label bindings have not yet been established, or
    // because some change in the agent has removed them.  When the label stack
    // associated with this VRF route is created, it MUST establish the associated
    // cross-connect entry in the mplsXCTable and then set that index to the value
    // of this object.  Changes to the cross-connect object in the mplsXCTable
    // MUST automatically be reflected in the value of this object.  If this
    // object represents a static routing entry, then the manager must ensure that
    // this entry is maintained consistently in the corresponding mplsXCTable as
    // well. The type is string with length: 1..24.
    Mplsl3Vpnvrfrtexcpointer interface{}

    // The row status variable, used according to row installation and removal
    // conventions.  A row entry cannot be modified when the status is marked as
    // active(1). The type is RowStatus.
    Mplsl3Vpnvrfrteinetcidrstatus interface{}
}

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetEntityData() *types.CommonEntityData {
    mplsl3Vpnvrfrteentry.EntityData.YFilter = mplsl3Vpnvrfrteentry.YFilter
    mplsl3Vpnvrfrteentry.EntityData.YangName = "mplsL3VpnVrfRteEntry"
    mplsl3Vpnvrfrteentry.EntityData.BundleName = "cisco_ios_xe"
    mplsl3Vpnvrfrteentry.EntityData.ParentYangName = "mplsL3VpnVrfRteTable"
    mplsl3Vpnvrfrteentry.EntityData.SegmentPath = "mplsL3VpnVrfRteEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfname) + "']" + "[mplsL3VpnVrfRteInetCidrDestType='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdesttype) + "']" + "[mplsL3VpnVrfRteInetCidrDest='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdest) + "']" + "[mplsL3VpnVrfRteInetCidrPfxLen='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpfxlen) + "']" + "[mplsL3VpnVrfRteInetCidrPolicy='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpolicy) + "']" + "[mplsL3VpnVrfRteInetCidrNHopType='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnhoptype) + "']" + "[mplsL3VpnVrfRteInetCidrNextHop='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnexthop) + "']"
    mplsl3Vpnvrfrteentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsl3Vpnvrfrteentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsl3Vpnvrfrteentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsl3Vpnvrfrteentry.EntityData.Children = make(map[string]types.YChild)
    mplsl3Vpnvrfrteentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfName"] = types.YLeaf{"Mplsl3Vpnvrfname", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfname}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrDestType"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrdesttype", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdesttype}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrDest"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrdest", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdest}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrPfxLen"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrpfxlen", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpfxlen}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrPolicy"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrpolicy", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpolicy}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrNHopType"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrnhoptype", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnhoptype}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrNextHop"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrnexthop", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnexthop}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrIfIndex"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrifindex", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrifindex}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrType"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrtype", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrtype}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrProto"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrproto", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrproto}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrAge"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrage", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrage}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrNextHopAS"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrnexthopas", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnexthopas}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrMetric1"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrmetric1", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric1}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrMetric2"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrmetric2", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric2}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrMetric3"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrmetric3", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric3}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrMetric4"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrmetric4", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric4}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrMetric5"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrmetric5", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric5}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteXCPointer"] = types.YLeaf{"Mplsl3Vpnvrfrtexcpointer", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrtexcpointer}
    mplsl3Vpnvrfrteentry.EntityData.Leafs["mplsL3VpnVrfRteInetCidrStatus"] = types.YLeaf{"Mplsl3Vpnvrfrteinetcidrstatus", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrstatus}
    return &(mplsl3Vpnvrfrteentry.EntityData)
}

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype represents discards the message silently.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype string

const (
    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_other MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "other"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_reject MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "reject"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_local MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "local"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_remote MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "remote"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_blackhole MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "blackhole"
)

