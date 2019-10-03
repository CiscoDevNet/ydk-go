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

    
    MplsL3VpnScalars MPLSL3VPNSTDMIB_MplsL3VpnScalars

    // This table specifies per-interface MPLS capability and associated
    // information.
    MplsL3VpnIfConfTable MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable

    // This table specifies per-interface MPLS L3VPN VRF Table capability and
    // associated information. Entries in this table define VRF routing instances
    // associated with MPLS/VPN interfaces.  Note that multiple interfaces can
    // belong to the same VRF instance.  The collection of all VRF instances
    // comprises an actual VPN.
    MplsL3VpnVrfTable MPLSL3VPNSTDMIB_MplsL3VpnVrfTable

    // This table specifies per-VRF route target association. Each entry
    // identifies a connectivity policy supported as part of a VPN.
    MplsL3VpnVrfRTTable MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable

    // This table specifies per-interface MPLS L3VPN VRF Table routing
    // information.  Entries in this table define VRF routing entries associated
    // with the specified MPLS/VPN interfaces.  Note  that this table contains
    // both BGP and Interior Gateway Protocol IGP routes, as both may appear in
    // the same VRF.
    MplsL3VpnVrfRteTable MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable
}

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetEntityData() *types.CommonEntityData {
    mPLSL3VPNSTDMIB.EntityData.YFilter = mPLSL3VPNSTDMIB.YFilter
    mPLSL3VPNSTDMIB.EntityData.YangName = "MPLS-L3VPN-STD-MIB"
    mPLSL3VPNSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSL3VPNSTDMIB.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mPLSL3VPNSTDMIB.EntityData.SegmentPath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB"
    mPLSL3VPNSTDMIB.EntityData.AbsolutePath = mPLSL3VPNSTDMIB.EntityData.SegmentPath
    mPLSL3VPNSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSL3VPNSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSL3VPNSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSL3VPNSTDMIB.EntityData.Children = types.NewOrderedMap()
    mPLSL3VPNSTDMIB.EntityData.Children.Append("mplsL3VpnScalars", types.YChild{"MplsL3VpnScalars", &mPLSL3VPNSTDMIB.MplsL3VpnScalars})
    mPLSL3VPNSTDMIB.EntityData.Children.Append("mplsL3VpnIfConfTable", types.YChild{"MplsL3VpnIfConfTable", &mPLSL3VPNSTDMIB.MplsL3VpnIfConfTable})
    mPLSL3VPNSTDMIB.EntityData.Children.Append("mplsL3VpnVrfTable", types.YChild{"MplsL3VpnVrfTable", &mPLSL3VPNSTDMIB.MplsL3VpnVrfTable})
    mPLSL3VPNSTDMIB.EntityData.Children.Append("mplsL3VpnVrfRTTable", types.YChild{"MplsL3VpnVrfRTTable", &mPLSL3VPNSTDMIB.MplsL3VpnVrfRTTable})
    mPLSL3VPNSTDMIB.EntityData.Children.Append("mplsL3VpnVrfRteTable", types.YChild{"MplsL3VpnVrfRteTable", &mPLSL3VPNSTDMIB.MplsL3VpnVrfRteTable})
    mPLSL3VPNSTDMIB.EntityData.Leafs = types.NewOrderedMap()

    mPLSL3VPNSTDMIB.EntityData.YListKeys = []string {}

    return &(mPLSL3VPNSTDMIB.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnScalars
type MPLSL3VPNSTDMIB_MplsL3VpnScalars struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of VRFs that are configured on this node. The type is
    // interface{} with range: 0..4294967295.
    MplsL3VpnConfiguredVrfs interface{}

    // The number of VRFs that are active on this node. That is, those VRFs whose
    // corresponding mplsL3VpnVrfOperStatus object value is equal to operational
    // (1). The type is interface{} with range: 0..4294967295.
    MplsL3VpnActiveVrfs interface{}

    // Total number of interfaces connected to a VRF. The type is interface{} with
    // range: 0..4294967295.
    MplsL3VpnConnectedInterfaces interface{}

    // If this object is true, then it enables the generation of all notifications
    // defined in this MIB.  This object's value should be preserved across agent
    // reboots. The type is bool.
    MplsL3VpnNotificationEnable interface{}

    // Denotes maximum number of routes that the device will allow all VRFs
    // jointly to hold.  If this value is set to 0, this indicates that the device
    // is unable to determine the absolute maximum.  In this case, the configured
    // maximum MAY not actually be allowed by the device. The type is interface{}
    // with range: 0..4294967295.
    MplsL3VpnVrfConfMaxPossRts interface{}

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
    MplsL3VpnVrfConfRteMxThrshTime interface{}

    // The number of illegally received labels above which the
    // mplsNumVrfSecIllglLblThrshExcd notification is issued.  The persistence of
    // this value mimics that of the device's configuration. The type is
    // interface{} with range: 0..4294967295.
    MplsL3VpnIllLblRcvThrsh interface{}
}

func (mplsL3VpnScalars *MPLSL3VPNSTDMIB_MplsL3VpnScalars) GetEntityData() *types.CommonEntityData {
    mplsL3VpnScalars.EntityData.YFilter = mplsL3VpnScalars.YFilter
    mplsL3VpnScalars.EntityData.YangName = "mplsL3VpnScalars"
    mplsL3VpnScalars.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnScalars.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsL3VpnScalars.EntityData.SegmentPath = "mplsL3VpnScalars"
    mplsL3VpnScalars.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/" + mplsL3VpnScalars.EntityData.SegmentPath
    mplsL3VpnScalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnScalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnScalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnScalars.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnScalars.EntityData.Leafs = types.NewOrderedMap()
    mplsL3VpnScalars.EntityData.Leafs.Append("mplsL3VpnConfiguredVrfs", types.YLeaf{"MplsL3VpnConfiguredVrfs", mplsL3VpnScalars.MplsL3VpnConfiguredVrfs})
    mplsL3VpnScalars.EntityData.Leafs.Append("mplsL3VpnActiveVrfs", types.YLeaf{"MplsL3VpnActiveVrfs", mplsL3VpnScalars.MplsL3VpnActiveVrfs})
    mplsL3VpnScalars.EntityData.Leafs.Append("mplsL3VpnConnectedInterfaces", types.YLeaf{"MplsL3VpnConnectedInterfaces", mplsL3VpnScalars.MplsL3VpnConnectedInterfaces})
    mplsL3VpnScalars.EntityData.Leafs.Append("mplsL3VpnNotificationEnable", types.YLeaf{"MplsL3VpnNotificationEnable", mplsL3VpnScalars.MplsL3VpnNotificationEnable})
    mplsL3VpnScalars.EntityData.Leafs.Append("mplsL3VpnVrfConfMaxPossRts", types.YLeaf{"MplsL3VpnVrfConfMaxPossRts", mplsL3VpnScalars.MplsL3VpnVrfConfMaxPossRts})
    mplsL3VpnScalars.EntityData.Leafs.Append("mplsL3VpnVrfConfRteMxThrshTime", types.YLeaf{"MplsL3VpnVrfConfRteMxThrshTime", mplsL3VpnScalars.MplsL3VpnVrfConfRteMxThrshTime})
    mplsL3VpnScalars.EntityData.Leafs.Append("mplsL3VpnIllLblRcvThrsh", types.YLeaf{"MplsL3VpnIllLblRcvThrsh", mplsL3VpnScalars.MplsL3VpnIllLblRcvThrsh})

    mplsL3VpnScalars.EntityData.YListKeys = []string {}

    return &(mplsL3VpnScalars.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every interface capable of
    // supporting MPLS L3VPN. Each entry in this table is meant to correspond to
    // an entry in the Interfaces Table. The type is slice of
    // MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry.
    MplsL3VpnIfConfEntry []*MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry
}

func (mplsL3VpnIfConfTable *MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable) GetEntityData() *types.CommonEntityData {
    mplsL3VpnIfConfTable.EntityData.YFilter = mplsL3VpnIfConfTable.YFilter
    mplsL3VpnIfConfTable.EntityData.YangName = "mplsL3VpnIfConfTable"
    mplsL3VpnIfConfTable.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnIfConfTable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsL3VpnIfConfTable.EntityData.SegmentPath = "mplsL3VpnIfConfTable"
    mplsL3VpnIfConfTable.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/" + mplsL3VpnIfConfTable.EntityData.SegmentPath
    mplsL3VpnIfConfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnIfConfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnIfConfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnIfConfTable.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnIfConfTable.EntityData.Children.Append("mplsL3VpnIfConfEntry", types.YChild{"MplsL3VpnIfConfEntry", nil})
    for i := range mplsL3VpnIfConfTable.MplsL3VpnIfConfEntry {
        mplsL3VpnIfConfTable.EntityData.Children.Append(types.GetSegmentPath(mplsL3VpnIfConfTable.MplsL3VpnIfConfEntry[i]), types.YChild{"MplsL3VpnIfConfEntry", mplsL3VpnIfConfTable.MplsL3VpnIfConfEntry[i]})
    }
    mplsL3VpnIfConfTable.EntityData.Leafs = types.NewOrderedMap()

    mplsL3VpnIfConfTable.EntityData.YListKeys = []string {}

    return &(mplsL3VpnIfConfTable.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry
// An entry in this table is created by an LSR for
// every interface capable of supporting MPLS L3VPN.
// Each entry in this table is meant to correspond to
// an entry in the Interfaces Table.
type MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_l3vpn_std_mib.MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfName
    MplsL3VpnVrfName interface{}

    // This attribute is a key. This is a unique index for an entry in the
    // mplsL3VpnIfConfTable.  A non-zero index for an entry indicates the ifIndex
    // for the corresponding interface entry in the MPLS-VPN-layer in the ifTable.
    // Note that this table does not necessarily correspond one-to-one with all
    // entries in the Interface MIB having an ifType of MPLS-layer; rather, only
    // those that are enabled for MPLS L3VPN functionality. The type is
    // interface{} with range: 1..2147483647.
    MplsL3VpnIfConfIndex interface{}

    // Denotes whether this link participates in a carrier's carrier, enterprise,
    // or inter-provider scenario. The type is MplsL3VpnIfVpnClassification.
    MplsL3VpnIfVpnClassification interface{}

    // Denotes the route distribution protocol across the PE-CE link.  Note that
    // more than one routing protocol may be enabled at the same time; thus, this
    // object is specified as a bitmask.  For example, static(5) and ospf(2) are a
    // typical configuration. The type is map[string]bool.
    MplsL3VpnIfVpnRouteDistProtocol interface{}

    // The storage type for this VPN If entry. Conceptual rows having the value
    // 'permanent' need not allow write access to any columnar objects in the row.
    // The type is StorageType.
    MplsL3VpnIfConfStorageType interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // Rows in this table signify that the specified interface is associated with
    // this VRF.  If the row creation operation succeeds, the interface will have
    // been associated with the specified VRF, otherwise the agent MUST not allow
    // the association.  If the agent only allows read-only operations on this
    // table, it MUST create entries in this table as they are created on the
    // device.  When a row in this table is in active(1) state, no objects in that
    // row can be modified except mplsL3VpnIfConfStorageType and
    // mplsL3VpnIfConfRowStatus. The type is RowStatus.
    MplsL3VpnIfConfRowStatus interface{}
}

func (mplsL3VpnIfConfEntry *MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry) GetEntityData() *types.CommonEntityData {
    mplsL3VpnIfConfEntry.EntityData.YFilter = mplsL3VpnIfConfEntry.YFilter
    mplsL3VpnIfConfEntry.EntityData.YangName = "mplsL3VpnIfConfEntry"
    mplsL3VpnIfConfEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnIfConfEntry.EntityData.ParentYangName = "mplsL3VpnIfConfTable"
    mplsL3VpnIfConfEntry.EntityData.SegmentPath = "mplsL3VpnIfConfEntry" + types.AddKeyToken(mplsL3VpnIfConfEntry.MplsL3VpnVrfName, "mplsL3VpnVrfName") + types.AddKeyToken(mplsL3VpnIfConfEntry.MplsL3VpnIfConfIndex, "mplsL3VpnIfConfIndex")
    mplsL3VpnIfConfEntry.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/mplsL3VpnIfConfTable/" + mplsL3VpnIfConfEntry.EntityData.SegmentPath
    mplsL3VpnIfConfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnIfConfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnIfConfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnIfConfEntry.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnIfConfEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsL3VpnIfConfEntry.EntityData.Leafs.Append("mplsL3VpnVrfName", types.YLeaf{"MplsL3VpnVrfName", mplsL3VpnIfConfEntry.MplsL3VpnVrfName})
    mplsL3VpnIfConfEntry.EntityData.Leafs.Append("mplsL3VpnIfConfIndex", types.YLeaf{"MplsL3VpnIfConfIndex", mplsL3VpnIfConfEntry.MplsL3VpnIfConfIndex})
    mplsL3VpnIfConfEntry.EntityData.Leafs.Append("mplsL3VpnIfVpnClassification", types.YLeaf{"MplsL3VpnIfVpnClassification", mplsL3VpnIfConfEntry.MplsL3VpnIfVpnClassification})
    mplsL3VpnIfConfEntry.EntityData.Leafs.Append("mplsL3VpnIfVpnRouteDistProtocol", types.YLeaf{"MplsL3VpnIfVpnRouteDistProtocol", mplsL3VpnIfConfEntry.MplsL3VpnIfVpnRouteDistProtocol})
    mplsL3VpnIfConfEntry.EntityData.Leafs.Append("mplsL3VpnIfConfStorageType", types.YLeaf{"MplsL3VpnIfConfStorageType", mplsL3VpnIfConfEntry.MplsL3VpnIfConfStorageType})
    mplsL3VpnIfConfEntry.EntityData.Leafs.Append("mplsL3VpnIfConfRowStatus", types.YLeaf{"MplsL3VpnIfConfRowStatus", mplsL3VpnIfConfEntry.MplsL3VpnIfConfRowStatus})

    mplsL3VpnIfConfEntry.EntityData.YListKeys = []string {"MplsL3VpnVrfName", "MplsL3VpnIfConfIndex"}

    return &(mplsL3VpnIfConfEntry.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification represents scenario.
type MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification string

const (
    MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification_carrierOfCarrier MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification = "carrierOfCarrier"

    MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification_enterprise MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification = "enterprise"

    MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification_interProvider MPLSL3VPNSTDMIB_MplsL3VpnIfConfTable_MplsL3VpnIfConfEntry_MplsL3VpnIfVpnClassification = "interProvider"
)

// MPLSL3VPNSTDMIB_MplsL3VpnVrfTable
// This table specifies per-interface MPLS L3VPN
// VRF Table capability and associated information.
// Entries in this table define VRF routing instances
// associated with MPLS/VPN interfaces.  Note that
// multiple interfaces can belong to the same VRF
// instance.  The collection of all VRF instances
// comprises an actual VPN.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS L3VPN.  The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of
    // MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry.
    MplsL3VpnVrfEntry []*MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry
}

func (mplsL3VpnVrfTable *MPLSL3VPNSTDMIB_MplsL3VpnVrfTable) GetEntityData() *types.CommonEntityData {
    mplsL3VpnVrfTable.EntityData.YFilter = mplsL3VpnVrfTable.YFilter
    mplsL3VpnVrfTable.EntityData.YangName = "mplsL3VpnVrfTable"
    mplsL3VpnVrfTable.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnVrfTable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsL3VpnVrfTable.EntityData.SegmentPath = "mplsL3VpnVrfTable"
    mplsL3VpnVrfTable.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/" + mplsL3VpnVrfTable.EntityData.SegmentPath
    mplsL3VpnVrfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnVrfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnVrfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnVrfTable.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnVrfTable.EntityData.Children.Append("mplsL3VpnVrfEntry", types.YChild{"MplsL3VpnVrfEntry", nil})
    for i := range mplsL3VpnVrfTable.MplsL3VpnVrfEntry {
        mplsL3VpnVrfTable.EntityData.Children.Append(types.GetSegmentPath(mplsL3VpnVrfTable.MplsL3VpnVrfEntry[i]), types.YChild{"MplsL3VpnVrfEntry", mplsL3VpnVrfTable.MplsL3VpnVrfEntry[i]})
    }
    mplsL3VpnVrfTable.EntityData.Leafs = types.NewOrderedMap()

    mplsL3VpnVrfTable.EntityData.YListKeys = []string {}

    return &(mplsL3VpnVrfTable.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS L3VPN.  The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The human-readable name of this VPN.  This MAY be
    // equivalent to the [RFC2685] VPN-ID, but may also vary.  If it is set to the
    // VPN ID, it MUST be equivalent to the value of mplsL3VpnVrfVpnId. It is
    // strongly recommended that all sites supporting VRFs that are part of the
    // same VPN use the same naming convention for VRFs as well as the same VPN
    // ID. The type is string with length: 0..31.
    MplsL3VpnVrfName interface{}

    // The VPN ID as specified in [RFC2685].  If a VPN ID has not been specified
    // for this VRF, then this variable SHOULD be set to a zero-length OCTET
    // STRING. The type is string with length: 0..0 | 7..7.
    MplsL3VpnVrfVpnId interface{}

    // The human-readable description of this VRF. The type is string.
    MplsL3VpnVrfDescription interface{}

    // The route distinguisher for this VRF. The type is string with length:
    // 0..256.
    MplsL3VpnVrfRD interface{}

    // The time at which this VRF entry was created. The type is interface{} with
    // range: 0..4294967295.
    MplsL3VpnVrfCreationTime interface{}

    // Denotes whether or not a VRF is operational.  A VRF is up(1) when there is
    // at least one interface associated with the VRF whose ifOperStatus is up(1).
    // A VRF is down(2) when: a. There does not exist at least one interface whose
    // ifOperStatus is up(1). b. There are no interfaces associated with the VRF.
    // The type is MplsL3VpnVrfOperStatus.
    MplsL3VpnVrfOperStatus interface{}

    // Total number of interfaces connected to this VRF with ifOperStatus = up(1).
    // This value should increase when an interface is associated with the
    // corresponding VRF and its corresponding ifOperStatus is equal to up(1).  If
    // an interface is associated whose ifOperStatus is not up(1), then the value
    // is not incremented until such time as it transitions to this state.  This
    // value should be decremented when an interface is disassociated with a VRF
    // or the corresponding ifOperStatus transitions out of the up(1) state to any
    // other state. The type is interface{} with range: 0..4294967295.
    MplsL3VpnVrfActiveInterfaces interface{}

    // Total number of interfaces connected to this VRF (independent of
    // ifOperStatus type). The type is interface{} with range: 0..4294967295.
    MplsL3VpnVrfAssociatedInterfaces interface{}

    // Denotes mid-level water marker for the number of routes that this VRF may
    // hold. The type is interface{} with range: 0..4294967295.
    MplsL3VpnVrfConfMidRteThresh interface{}

    // Denotes high-level water marker for the number of routes that this VRF may
    // hold. The type is interface{} with range: 0..4294967295.
    MplsL3VpnVrfConfHighRteThresh interface{}

    // Denotes maximum number of routes that this VRF is configured to hold.  This
    // value MUST be less than or equal to mplsL3VpnVrfConfMaxPossRts unless it is
    // set to 0. The type is interface{} with range: 0..4294967295.
    MplsL3VpnVrfConfMaxRoutes interface{}

    // The value of sysUpTime at the time of the last change of this table entry,
    // which includes changes of VRF parameters defined in this table or addition
    // or deletion of interfaces associated with this VRF. The type is interface{}
    // with range: 0..4294967295.
    MplsL3VpnVrfConfLastChanged interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified except mplsL3VpnVrfConfAdminStatus, mplsL3VpnVrfConfRowStatus,
    // and mplsL3VpnVrfConfStorageType. The type is RowStatus.
    MplsL3VpnVrfConfRowStatus interface{}

    // Indicates the desired operational status of this VRF. The type is
    // MplsL3VpnVrfConfAdminStatus.
    MplsL3VpnVrfConfAdminStatus interface{}

    // The storage type for this VPN VRF entry. Conceptual rows having the value
    // 'permanent' need not allow write access to any columnar objects in the row.
    // The type is StorageType.
    MplsL3VpnVrfConfStorageType interface{}

    // Indicates the number of illegally received labels on this VPN/VRF. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // mplsL3VpnVrfSecDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    MplsL3VpnVrfSecIllegalLblVltns interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    MplsL3VpnVrfSecDiscontinuityTime interface{}

    // Indicates the number of routes added to this VPN/VRF since the last
    // discontinuity.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of mplsL3VpnVrfPerfDiscTime. The type is interface{} with
    // range: 0..4294967295.
    MplsL3VpnVrfPerfRoutesAdded interface{}

    // Indicates the number of routes removed from this VPN/VRF.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // mplsL3VpnVrfPerfDiscTime. The type is interface{} with range:
    // 0..4294967295.
    MplsL3VpnVrfPerfRoutesDeleted interface{}

    // Indicates the number of routes currently used by this VRF. The type is
    // interface{} with range: 0..4294967295.
    MplsL3VpnVrfPerfCurrNumRoutes interface{}

    // This counter should be incremented when the number of routes contained by
    // the specified VRF exceeds or attempts to exceed the maximum allowed value
    // as indicated by mplsL3VpnVrfMaxRouteThreshold.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // mplsL3VpnVrfPerfDiscTime. The type is interface{} with range:
    // 0..4294967295.
    MplsL3VpnVrfPerfRoutesDropped interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entry's counters suffered a discontinuity.  If no such
    // discontinuities have occurred since the last re-initialization of the local
    // management subsystem, then this object contains a zero value. The type is
    // interface{} with range: 0..4294967295.
    MplsL3VpnVrfPerfDiscTime interface{}
}

func (mplsL3VpnVrfEntry *MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry) GetEntityData() *types.CommonEntityData {
    mplsL3VpnVrfEntry.EntityData.YFilter = mplsL3VpnVrfEntry.YFilter
    mplsL3VpnVrfEntry.EntityData.YangName = "mplsL3VpnVrfEntry"
    mplsL3VpnVrfEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnVrfEntry.EntityData.ParentYangName = "mplsL3VpnVrfTable"
    mplsL3VpnVrfEntry.EntityData.SegmentPath = "mplsL3VpnVrfEntry" + types.AddKeyToken(mplsL3VpnVrfEntry.MplsL3VpnVrfName, "mplsL3VpnVrfName")
    mplsL3VpnVrfEntry.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/mplsL3VpnVrfTable/" + mplsL3VpnVrfEntry.EntityData.SegmentPath
    mplsL3VpnVrfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnVrfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnVrfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnVrfEntry.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnVrfEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfName", types.YLeaf{"MplsL3VpnVrfName", mplsL3VpnVrfEntry.MplsL3VpnVrfName})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfVpnId", types.YLeaf{"MplsL3VpnVrfVpnId", mplsL3VpnVrfEntry.MplsL3VpnVrfVpnId})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfDescription", types.YLeaf{"MplsL3VpnVrfDescription", mplsL3VpnVrfEntry.MplsL3VpnVrfDescription})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfRD", types.YLeaf{"MplsL3VpnVrfRD", mplsL3VpnVrfEntry.MplsL3VpnVrfRD})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfCreationTime", types.YLeaf{"MplsL3VpnVrfCreationTime", mplsL3VpnVrfEntry.MplsL3VpnVrfCreationTime})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfOperStatus", types.YLeaf{"MplsL3VpnVrfOperStatus", mplsL3VpnVrfEntry.MplsL3VpnVrfOperStatus})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfActiveInterfaces", types.YLeaf{"MplsL3VpnVrfActiveInterfaces", mplsL3VpnVrfEntry.MplsL3VpnVrfActiveInterfaces})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfAssociatedInterfaces", types.YLeaf{"MplsL3VpnVrfAssociatedInterfaces", mplsL3VpnVrfEntry.MplsL3VpnVrfAssociatedInterfaces})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfConfMidRteThresh", types.YLeaf{"MplsL3VpnVrfConfMidRteThresh", mplsL3VpnVrfEntry.MplsL3VpnVrfConfMidRteThresh})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfConfHighRteThresh", types.YLeaf{"MplsL3VpnVrfConfHighRteThresh", mplsL3VpnVrfEntry.MplsL3VpnVrfConfHighRteThresh})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfConfMaxRoutes", types.YLeaf{"MplsL3VpnVrfConfMaxRoutes", mplsL3VpnVrfEntry.MplsL3VpnVrfConfMaxRoutes})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfConfLastChanged", types.YLeaf{"MplsL3VpnVrfConfLastChanged", mplsL3VpnVrfEntry.MplsL3VpnVrfConfLastChanged})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfConfRowStatus", types.YLeaf{"MplsL3VpnVrfConfRowStatus", mplsL3VpnVrfEntry.MplsL3VpnVrfConfRowStatus})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfConfAdminStatus", types.YLeaf{"MplsL3VpnVrfConfAdminStatus", mplsL3VpnVrfEntry.MplsL3VpnVrfConfAdminStatus})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfConfStorageType", types.YLeaf{"MplsL3VpnVrfConfStorageType", mplsL3VpnVrfEntry.MplsL3VpnVrfConfStorageType})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfSecIllegalLblVltns", types.YLeaf{"MplsL3VpnVrfSecIllegalLblVltns", mplsL3VpnVrfEntry.MplsL3VpnVrfSecIllegalLblVltns})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfSecDiscontinuityTime", types.YLeaf{"MplsL3VpnVrfSecDiscontinuityTime", mplsL3VpnVrfEntry.MplsL3VpnVrfSecDiscontinuityTime})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfPerfRoutesAdded", types.YLeaf{"MplsL3VpnVrfPerfRoutesAdded", mplsL3VpnVrfEntry.MplsL3VpnVrfPerfRoutesAdded})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfPerfRoutesDeleted", types.YLeaf{"MplsL3VpnVrfPerfRoutesDeleted", mplsL3VpnVrfEntry.MplsL3VpnVrfPerfRoutesDeleted})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfPerfCurrNumRoutes", types.YLeaf{"MplsL3VpnVrfPerfCurrNumRoutes", mplsL3VpnVrfEntry.MplsL3VpnVrfPerfCurrNumRoutes})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfPerfRoutesDropped", types.YLeaf{"MplsL3VpnVrfPerfRoutesDropped", mplsL3VpnVrfEntry.MplsL3VpnVrfPerfRoutesDropped})
    mplsL3VpnVrfEntry.EntityData.Leafs.Append("mplsL3VpnVrfPerfDiscTime", types.YLeaf{"MplsL3VpnVrfPerfDiscTime", mplsL3VpnVrfEntry.MplsL3VpnVrfPerfDiscTime})

    mplsL3VpnVrfEntry.EntityData.YListKeys = []string {"MplsL3VpnVrfName"}

    return &(mplsL3VpnVrfEntry.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus represents VRF.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus string

const (
    MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus_up MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus = "up"

    MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus_down MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus = "down"

    MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus_testing MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfConfAdminStatus = "testing"
)

// MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfOperStatus represents b. There are no interfaces associated with the VRF.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfOperStatus string

const (
    MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfOperStatus_up MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfOperStatus = "up"

    MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfOperStatus_down MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfOperStatus = "down"
)

// MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable
// This table specifies per-VRF route target association.
// Each entry identifies a connectivity policy supported
// as part of a VPN.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for each route target
    // configured for a VRF supporting a MPLS L3VPN instance.  The indexing
    // provides an ordering per-VRF instance.  See [RFC4364] for a complete
    // definition of a route target. The type is slice of
    // MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable_MplsL3VpnVrfRTEntry.
    MplsL3VpnVrfRTEntry []*MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable_MplsL3VpnVrfRTEntry
}

func (mplsL3VpnVrfRTTable *MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable) GetEntityData() *types.CommonEntityData {
    mplsL3VpnVrfRTTable.EntityData.YFilter = mplsL3VpnVrfRTTable.YFilter
    mplsL3VpnVrfRTTable.EntityData.YangName = "mplsL3VpnVrfRTTable"
    mplsL3VpnVrfRTTable.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnVrfRTTable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsL3VpnVrfRTTable.EntityData.SegmentPath = "mplsL3VpnVrfRTTable"
    mplsL3VpnVrfRTTable.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/" + mplsL3VpnVrfRTTable.EntityData.SegmentPath
    mplsL3VpnVrfRTTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnVrfRTTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnVrfRTTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnVrfRTTable.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnVrfRTTable.EntityData.Children.Append("mplsL3VpnVrfRTEntry", types.YChild{"MplsL3VpnVrfRTEntry", nil})
    for i := range mplsL3VpnVrfRTTable.MplsL3VpnVrfRTEntry {
        mplsL3VpnVrfRTTable.EntityData.Children.Append(types.GetSegmentPath(mplsL3VpnVrfRTTable.MplsL3VpnVrfRTEntry[i]), types.YChild{"MplsL3VpnVrfRTEntry", mplsL3VpnVrfRTTable.MplsL3VpnVrfRTEntry[i]})
    }
    mplsL3VpnVrfRTTable.EntityData.Leafs = types.NewOrderedMap()

    mplsL3VpnVrfRTTable.EntityData.YListKeys = []string {}

    return &(mplsL3VpnVrfRTTable.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable_MplsL3VpnVrfRTEntry
// An entry in this table is created by an LSR for
// each route target configured for a VRF supporting
// a MPLS L3VPN instance.  The indexing provides an
// ordering per-VRF instance.  See [RFC4364] for a
// complete definition of a route target.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable_MplsL3VpnVrfRTEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_l3vpn_std_mib.MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfName
    MplsL3VpnVrfName interface{}

    // This attribute is a key. Auxiliary index for route targets configured for a
    // particular VRF. The type is interface{} with range: 1..4294967295.
    MplsL3VpnVrfRTIndex interface{}

    // This attribute is a key. The route target distribution type. The type is
    // MplsL3VpnRtType.
    MplsL3VpnVrfRTType interface{}

    // The route target distribution policy. The type is string with length:
    // 0..256.
    MplsL3VpnVrfRT interface{}

    // Description of the route target. The type is string.
    MplsL3VpnVrfRTDescr interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this table is in active(1) state, no objects in that row can
    // be modified except mplsL3VpnVrfRTRowStatus. The type is RowStatus.
    MplsL3VpnVrfRTRowStatus interface{}

    // The storage type for this VPN route target (RT) entry. Conceptual rows
    // having the value 'permanent' need not allow write access to any columnar
    // objects in the row. The type is StorageType.
    MplsL3VpnVrfRTStorageType interface{}
}

func (mplsL3VpnVrfRTEntry *MPLSL3VPNSTDMIB_MplsL3VpnVrfRTTable_MplsL3VpnVrfRTEntry) GetEntityData() *types.CommonEntityData {
    mplsL3VpnVrfRTEntry.EntityData.YFilter = mplsL3VpnVrfRTEntry.YFilter
    mplsL3VpnVrfRTEntry.EntityData.YangName = "mplsL3VpnVrfRTEntry"
    mplsL3VpnVrfRTEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnVrfRTEntry.EntityData.ParentYangName = "mplsL3VpnVrfRTTable"
    mplsL3VpnVrfRTEntry.EntityData.SegmentPath = "mplsL3VpnVrfRTEntry" + types.AddKeyToken(mplsL3VpnVrfRTEntry.MplsL3VpnVrfName, "mplsL3VpnVrfName") + types.AddKeyToken(mplsL3VpnVrfRTEntry.MplsL3VpnVrfRTIndex, "mplsL3VpnVrfRTIndex") + types.AddKeyToken(mplsL3VpnVrfRTEntry.MplsL3VpnVrfRTType, "mplsL3VpnVrfRTType")
    mplsL3VpnVrfRTEntry.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/mplsL3VpnVrfRTTable/" + mplsL3VpnVrfRTEntry.EntityData.SegmentPath
    mplsL3VpnVrfRTEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnVrfRTEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnVrfRTEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnVrfRTEntry.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnVrfRTEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsL3VpnVrfRTEntry.EntityData.Leafs.Append("mplsL3VpnVrfName", types.YLeaf{"MplsL3VpnVrfName", mplsL3VpnVrfRTEntry.MplsL3VpnVrfName})
    mplsL3VpnVrfRTEntry.EntityData.Leafs.Append("mplsL3VpnVrfRTIndex", types.YLeaf{"MplsL3VpnVrfRTIndex", mplsL3VpnVrfRTEntry.MplsL3VpnVrfRTIndex})
    mplsL3VpnVrfRTEntry.EntityData.Leafs.Append("mplsL3VpnVrfRTType", types.YLeaf{"MplsL3VpnVrfRTType", mplsL3VpnVrfRTEntry.MplsL3VpnVrfRTType})
    mplsL3VpnVrfRTEntry.EntityData.Leafs.Append("mplsL3VpnVrfRT", types.YLeaf{"MplsL3VpnVrfRT", mplsL3VpnVrfRTEntry.MplsL3VpnVrfRT})
    mplsL3VpnVrfRTEntry.EntityData.Leafs.Append("mplsL3VpnVrfRTDescr", types.YLeaf{"MplsL3VpnVrfRTDescr", mplsL3VpnVrfRTEntry.MplsL3VpnVrfRTDescr})
    mplsL3VpnVrfRTEntry.EntityData.Leafs.Append("mplsL3VpnVrfRTRowStatus", types.YLeaf{"MplsL3VpnVrfRTRowStatus", mplsL3VpnVrfRTEntry.MplsL3VpnVrfRTRowStatus})
    mplsL3VpnVrfRTEntry.EntityData.Leafs.Append("mplsL3VpnVrfRTStorageType", types.YLeaf{"MplsL3VpnVrfRTStorageType", mplsL3VpnVrfRTEntry.MplsL3VpnVrfRTStorageType})

    mplsL3VpnVrfRTEntry.EntityData.YListKeys = []string {"MplsL3VpnVrfName", "MplsL3VpnVrfRTIndex", "MplsL3VpnVrfRTType"}

    return &(mplsL3VpnVrfRTEntry.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable
// This table specifies per-interface MPLS L3VPN VRF Table
// routing information.  Entries in this table define VRF routing
// entries associated with the specified MPLS/VPN interfaces.  Note
// 
// that this table contains both BGP and Interior Gateway Protocol
// IGP routes, as both may appear in the same VRF.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable struct {
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
    // MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry.
    MplsL3VpnVrfRteEntry []*MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry
}

func (mplsL3VpnVrfRteTable *MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable) GetEntityData() *types.CommonEntityData {
    mplsL3VpnVrfRteTable.EntityData.YFilter = mplsL3VpnVrfRteTable.YFilter
    mplsL3VpnVrfRteTable.EntityData.YangName = "mplsL3VpnVrfRteTable"
    mplsL3VpnVrfRteTable.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnVrfRteTable.EntityData.ParentYangName = "MPLS-L3VPN-STD-MIB"
    mplsL3VpnVrfRteTable.EntityData.SegmentPath = "mplsL3VpnVrfRteTable"
    mplsL3VpnVrfRteTable.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/" + mplsL3VpnVrfRteTable.EntityData.SegmentPath
    mplsL3VpnVrfRteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnVrfRteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnVrfRteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnVrfRteTable.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnVrfRteTable.EntityData.Children.Append("mplsL3VpnVrfRteEntry", types.YChild{"MplsL3VpnVrfRteEntry", nil})
    for i := range mplsL3VpnVrfRteTable.MplsL3VpnVrfRteEntry {
        mplsL3VpnVrfRteTable.EntityData.Children.Append(types.GetSegmentPath(mplsL3VpnVrfRteTable.MplsL3VpnVrfRteEntry[i]), types.YChild{"MplsL3VpnVrfRteEntry", mplsL3VpnVrfRteTable.MplsL3VpnVrfRteEntry[i]})
    }
    mplsL3VpnVrfRteTable.EntityData.Leafs = types.NewOrderedMap()

    mplsL3VpnVrfRteTable.EntityData.YListKeys = []string {}

    return &(mplsL3VpnVrfRteTable.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry
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
type MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_l3vpn_std_mib.MPLSL3VPNSTDMIB_MplsL3VpnVrfTable_MplsL3VpnVrfEntry_MplsL3VpnVrfName
    MplsL3VpnVrfName interface{}

    // This attribute is a key. The type of the mplsL3VpnVrfRteInetCidrDest
    // address, as defined in the InetAddress MIB.  Only those address types that
    // may appear in an actual routing table are allowed as values of this object.
    // The type is InetAddressType.
    MplsL3VpnVrfRteInetCidrDestType interface{}

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
    MplsL3VpnVrfRteInetCidrDest interface{}

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
    MplsL3VpnVrfRteInetCidrPfxLen interface{}

    // This attribute is a key. This object is an opaque object without any
    // defined semantics.  Its purpose is to serve as an additional index that may
    // delineate between multiple entries to the same destination.  The value { 0
    // 0 } shall be used as the default value for this object. The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    MplsL3VpnVrfRteInetCidrPolicy interface{}

    // This attribute is a key. The type of the mplsL3VpnVrfRteInetCidrNextHop
    // address, as defined in the InetAddress MIB.  Value should be set to
    // unknown(0) for non-remote routes.  Only those address types that may appear
    // in an actual routing table are allowed as values of this object. The type
    // is InetAddressType.
    MplsL3VpnVrfRteInetCidrNHopType interface{}

    // This attribute is a key. On remote routes, the address of the next system
    // en route.  For non-remote routes, a zero-length string. The type of this
    // address is determined by the value of the mplsL3VpnVrfRteInetCidrNHopType
    // object. The type is string with length: 0..255.
    MplsL3VpnVrfRteInetCidrNextHop interface{}

    // The ifIndex value that identifies the local interface through which the
    // next hop of this route should be reached.  A value of 0 is valid and
    // represents the scenario where no interface is specified. The type is
    // interface{} with range: 0..2147483647.
    MplsL3VpnVrfRteInetCidrIfIndex interface{}

    // The type of route.  Note that local(3) refers to a route for which the next
    // hop is the final destination; remote(4) refers to a route for which the
    // next hop is not the final destination.  Routes that do not result in
    // traffic forwarding or rejection should not be displayed even if the
    // implementation keeps them stored internally.  reject(2) refers to a route
    // that, if matched, discards the message as unreachable and returns a
    // notification (e.g., ICMP error) to the message sender.  This is used in
    // some protocols as a means of correctly aggregating routes.  blackhole(5)
    // refers to a route that, if matched, discards the message silently. The type
    // is MplsL3VpnVrfRteInetCidrType.
    MplsL3VpnVrfRteInetCidrType interface{}

    // The routing mechanism via which this route was learned. Inclusion of values
    // for gateway routing protocols is not intended to imply that hosts should
    // support those protocols. The type is IANAipRouteProtocol.
    MplsL3VpnVrfRteInetCidrProto interface{}

    // The number of seconds since this route was last updated or otherwise
    // determined to be correct.  Note that no semantics of 'too old' can be
    // implied except through knowledge of the routing protocol by which the route
    // was learned. The type is interface{} with range: 0..4294967295.
    MplsL3VpnVrfRteInetCidrAge interface{}

    // The Autonomous System Number of the next hop.  The semantics of this object
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  When this object is unknown or not
    // relevant, its value should be set to zero. The type is interface{} with
    // range: 0..4294967295.
    MplsL3VpnVrfRteInetCidrNextHopAS interface{}

    // The primary routing metric for this route.  The semantics of this metric
    // are determined by the  routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    MplsL3VpnVrfRteInetCidrMetric1 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    MplsL3VpnVrfRteInetCidrMetric2 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    MplsL3VpnVrfRteInetCidrMetric3 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    MplsL3VpnVrfRteInetCidrMetric4 interface{}

    // An alternate routing metric for this route.  The semantics of this metric
    // are determined by the routing protocol specified in the route's
    // mplsL3VpnVrfRteInetCidrProto value.  If this metric is not used, its value
    // should be set to -1. The type is interface{} with range: -1..2147483647.
    MplsL3VpnVrfRteInetCidrMetric5 interface{}

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
    MplsL3VpnVrfRteXCPointer interface{}

    // The row status variable, used according to row installation and removal
    // conventions.  A row entry cannot be modified when the status is marked as
    // active(1). The type is RowStatus.
    MplsL3VpnVrfRteInetCidrStatus interface{}
}

func (mplsL3VpnVrfRteEntry *MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry) GetEntityData() *types.CommonEntityData {
    mplsL3VpnVrfRteEntry.EntityData.YFilter = mplsL3VpnVrfRteEntry.YFilter
    mplsL3VpnVrfRteEntry.EntityData.YangName = "mplsL3VpnVrfRteEntry"
    mplsL3VpnVrfRteEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsL3VpnVrfRteEntry.EntityData.ParentYangName = "mplsL3VpnVrfRteTable"
    mplsL3VpnVrfRteEntry.EntityData.SegmentPath = "mplsL3VpnVrfRteEntry" + types.AddKeyToken(mplsL3VpnVrfRteEntry.MplsL3VpnVrfName, "mplsL3VpnVrfName") + types.AddKeyToken(mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrDestType, "mplsL3VpnVrfRteInetCidrDestType") + types.AddKeyToken(mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrDest, "mplsL3VpnVrfRteInetCidrDest") + types.AddKeyToken(mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrPfxLen, "mplsL3VpnVrfRteInetCidrPfxLen") + types.AddKeyToken(mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrPolicy, "mplsL3VpnVrfRteInetCidrPolicy") + types.AddKeyToken(mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrNHopType, "mplsL3VpnVrfRteInetCidrNHopType") + types.AddKeyToken(mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrNextHop, "mplsL3VpnVrfRteInetCidrNextHop")
    mplsL3VpnVrfRteEntry.EntityData.AbsolutePath = "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB/mplsL3VpnVrfRteTable/" + mplsL3VpnVrfRteEntry.EntityData.SegmentPath
    mplsL3VpnVrfRteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsL3VpnVrfRteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsL3VpnVrfRteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsL3VpnVrfRteEntry.EntityData.Children = types.NewOrderedMap()
    mplsL3VpnVrfRteEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfName", types.YLeaf{"MplsL3VpnVrfName", mplsL3VpnVrfRteEntry.MplsL3VpnVrfName})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrDestType", types.YLeaf{"MplsL3VpnVrfRteInetCidrDestType", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrDestType})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrDest", types.YLeaf{"MplsL3VpnVrfRteInetCidrDest", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrDest})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrPfxLen", types.YLeaf{"MplsL3VpnVrfRteInetCidrPfxLen", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrPfxLen})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrPolicy", types.YLeaf{"MplsL3VpnVrfRteInetCidrPolicy", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrPolicy})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrNHopType", types.YLeaf{"MplsL3VpnVrfRteInetCidrNHopType", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrNHopType})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrNextHop", types.YLeaf{"MplsL3VpnVrfRteInetCidrNextHop", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrNextHop})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrIfIndex", types.YLeaf{"MplsL3VpnVrfRteInetCidrIfIndex", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrIfIndex})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrType", types.YLeaf{"MplsL3VpnVrfRteInetCidrType", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrType})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrProto", types.YLeaf{"MplsL3VpnVrfRteInetCidrProto", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrProto})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrAge", types.YLeaf{"MplsL3VpnVrfRteInetCidrAge", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrAge})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrNextHopAS", types.YLeaf{"MplsL3VpnVrfRteInetCidrNextHopAS", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrNextHopAS})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrMetric1", types.YLeaf{"MplsL3VpnVrfRteInetCidrMetric1", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrMetric1})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrMetric2", types.YLeaf{"MplsL3VpnVrfRteInetCidrMetric2", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrMetric2})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrMetric3", types.YLeaf{"MplsL3VpnVrfRteInetCidrMetric3", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrMetric3})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrMetric4", types.YLeaf{"MplsL3VpnVrfRteInetCidrMetric4", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrMetric4})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrMetric5", types.YLeaf{"MplsL3VpnVrfRteInetCidrMetric5", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrMetric5})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteXCPointer", types.YLeaf{"MplsL3VpnVrfRteXCPointer", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteXCPointer})
    mplsL3VpnVrfRteEntry.EntityData.Leafs.Append("mplsL3VpnVrfRteInetCidrStatus", types.YLeaf{"MplsL3VpnVrfRteInetCidrStatus", mplsL3VpnVrfRteEntry.MplsL3VpnVrfRteInetCidrStatus})

    mplsL3VpnVrfRteEntry.EntityData.YListKeys = []string {"MplsL3VpnVrfName", "MplsL3VpnVrfRteInetCidrDestType", "MplsL3VpnVrfRteInetCidrDest", "MplsL3VpnVrfRteInetCidrPfxLen", "MplsL3VpnVrfRteInetCidrPolicy", "MplsL3VpnVrfRteInetCidrNHopType", "MplsL3VpnVrfRteInetCidrNextHop"}

    return &(mplsL3VpnVrfRteEntry.EntityData)
}

// MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType represents discards the message silently.
type MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType string

const (
    MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType_other MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType = "other"

    MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType_reject MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType = "reject"

    MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType_local MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType = "local"

    MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType_remote MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType = "remote"

    MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType_blackhole MPLSL3VPNSTDMIB_MplsL3VpnVrfRteTable_MplsL3VpnVrfRteEntry_MplsL3VpnVrfRteInetCidrType = "blackhole"
)

