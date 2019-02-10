// This MIB contains managed object definitions for the
// Multiprotocol Label Switching (MPLS)/Border Gateway
// 
// 
// Protocol (BGP) Virtual Private Networks (VPNs) as
// defined in : Rosen, E., Viswanathan, A., and R.
// Callon, Multiprotocol Label Switching Architecture,
// RFC3031, January 2001.
package mpls_vpn_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_vpn_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:MPLS-VPN-MIB MPLS-VPN-MIB}", reflect.TypeOf(MPLSVPNMIB{}))
    ydk.RegisterEntity("MPLS-VPN-MIB:MPLS-VPN-MIB", reflect.TypeOf(MPLSVPNMIB{}))
}

// MPLSVPNMIB
type MPLSVPNMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    MplsVpnScalars MPLSVPNMIB_MplsVpnScalars

    // This table specifies per-interface MPLS capability and associated
    // information.
    MplsVpnInterfaceConfTable MPLSVPNMIB_MplsVpnInterfaceConfTable

    // This table specifies per-interface MPLS/BGP VPN VRF Table capability and
    // associated information. Entries in this table define VRF routing instances
    // associated with MPLS/VPN interfaces. Note that multiple interfaces can
    // belong to the same VRF instance. The collection of all VRF instances
    // comprises an actual VPN.
    MplsVpnVrfTable MPLSVPNMIB_MplsVpnVrfTable

    // This table specifies per-VRF route target association. Each entry
    // identifies a connectivity policy supported as part of a VPN.
    MplsVpnVrfRouteTargetTable MPLSVPNMIB_MplsVpnVrfRouteTargetTable

    // Each entry in this table specifies a per-interface  MPLS/EBGP neighbor.
    MplsVpnVrfBgpNbrAddrTable MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable

    // This table specifies per-VRF vpnv4 multi-protocol prefixes supported by
    // BGP.
    MplsVpnVrfBgpNbrPrefixTable MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable

    // This table specifies per-interface MPLS/BGP VPN VRF Table routing
    // information. Entries in this table define VRF routing entries associated
    // with the specified MPLS/VPN interfaces. Note that this table contains both
    // BGP and IGP routes, as both may appear in the same VRF.
    MplsVpnVrfRouteTable MPLSVPNMIB_MplsVpnVrfRouteTable
}

func (mPLSVPNMIB *MPLSVPNMIB) GetEntityData() *types.CommonEntityData {
    mPLSVPNMIB.EntityData.YFilter = mPLSVPNMIB.YFilter
    mPLSVPNMIB.EntityData.YangName = "MPLS-VPN-MIB"
    mPLSVPNMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSVPNMIB.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mPLSVPNMIB.EntityData.SegmentPath = "MPLS-VPN-MIB:MPLS-VPN-MIB"
    mPLSVPNMIB.EntityData.AbsolutePath = mPLSVPNMIB.EntityData.SegmentPath
    mPLSVPNMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSVPNMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSVPNMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSVPNMIB.EntityData.Children = types.NewOrderedMap()
    mPLSVPNMIB.EntityData.Children.Append("mplsVpnScalars", types.YChild{"MplsVpnScalars", &mPLSVPNMIB.MplsVpnScalars})
    mPLSVPNMIB.EntityData.Children.Append("mplsVpnInterfaceConfTable", types.YChild{"MplsVpnInterfaceConfTable", &mPLSVPNMIB.MplsVpnInterfaceConfTable})
    mPLSVPNMIB.EntityData.Children.Append("mplsVpnVrfTable", types.YChild{"MplsVpnVrfTable", &mPLSVPNMIB.MplsVpnVrfTable})
    mPLSVPNMIB.EntityData.Children.Append("mplsVpnVrfRouteTargetTable", types.YChild{"MplsVpnVrfRouteTargetTable", &mPLSVPNMIB.MplsVpnVrfRouteTargetTable})
    mPLSVPNMIB.EntityData.Children.Append("mplsVpnVrfBgpNbrAddrTable", types.YChild{"MplsVpnVrfBgpNbrAddrTable", &mPLSVPNMIB.MplsVpnVrfBgpNbrAddrTable})
    mPLSVPNMIB.EntityData.Children.Append("mplsVpnVrfBgpNbrPrefixTable", types.YChild{"MplsVpnVrfBgpNbrPrefixTable", &mPLSVPNMIB.MplsVpnVrfBgpNbrPrefixTable})
    mPLSVPNMIB.EntityData.Children.Append("mplsVpnVrfRouteTable", types.YChild{"MplsVpnVrfRouteTable", &mPLSVPNMIB.MplsVpnVrfRouteTable})
    mPLSVPNMIB.EntityData.Leafs = types.NewOrderedMap()

    mPLSVPNMIB.EntityData.YListKeys = []string {}

    return &(mPLSVPNMIB.EntityData)
}

// MPLSVPNMIB_MplsVpnScalars
type MPLSVPNMIB_MplsVpnScalars struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of VRFs which are configured on this node. The type is
    // interface{} with range: 0..4294967295.
    MplsVpnConfiguredVrfs interface{}

    // The number of VRFs which are active on this node. That is, those VRFs whose
    // corresponding mplsVpnVrfOperStatus  object value is equal to operational
    // (1). The type is interface{} with range: 0..4294967295.
    MplsVpnActiveVrfs interface{}

    // Total number of interfaces connected to a VRF. The type is interface{} with
    // range: 0..4294967295.
    MplsVpnConnectedInterfaces interface{}

    // If this object is true, then it enables the generation of all notifications
    // defined in  this MIB. The type is bool.
    MplsVpnNotificationEnable interface{}

    // Denotes maximum number of routes which the device will allow all VRFs
    // jointly to hold. If this value is set to 0, this indicates that the device
    // is  unable to determine the absolute maximum. In this case, the configured
    // maximum MAY not actually be allowed by the device. The type is interface{}
    // with range: 0..4294967295.
    MplsVpnVrfConfMaxPossibleRoutes interface{}
}

func (mplsVpnScalars *MPLSVPNMIB_MplsVpnScalars) GetEntityData() *types.CommonEntityData {
    mplsVpnScalars.EntityData.YFilter = mplsVpnScalars.YFilter
    mplsVpnScalars.EntityData.YangName = "mplsVpnScalars"
    mplsVpnScalars.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnScalars.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsVpnScalars.EntityData.SegmentPath = "mplsVpnScalars"
    mplsVpnScalars.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/" + mplsVpnScalars.EntityData.SegmentPath
    mplsVpnScalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnScalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnScalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnScalars.EntityData.Children = types.NewOrderedMap()
    mplsVpnScalars.EntityData.Leafs = types.NewOrderedMap()
    mplsVpnScalars.EntityData.Leafs.Append("mplsVpnConfiguredVrfs", types.YLeaf{"MplsVpnConfiguredVrfs", mplsVpnScalars.MplsVpnConfiguredVrfs})
    mplsVpnScalars.EntityData.Leafs.Append("mplsVpnActiveVrfs", types.YLeaf{"MplsVpnActiveVrfs", mplsVpnScalars.MplsVpnActiveVrfs})
    mplsVpnScalars.EntityData.Leafs.Append("mplsVpnConnectedInterfaces", types.YLeaf{"MplsVpnConnectedInterfaces", mplsVpnScalars.MplsVpnConnectedInterfaces})
    mplsVpnScalars.EntityData.Leafs.Append("mplsVpnNotificationEnable", types.YLeaf{"MplsVpnNotificationEnable", mplsVpnScalars.MplsVpnNotificationEnable})
    mplsVpnScalars.EntityData.Leafs.Append("mplsVpnVrfConfMaxPossibleRoutes", types.YLeaf{"MplsVpnVrfConfMaxPossibleRoutes", mplsVpnScalars.MplsVpnVrfConfMaxPossibleRoutes})

    mplsVpnScalars.EntityData.YListKeys = []string {}

    return &(mplsVpnScalars.EntityData)
}

// MPLSVPNMIB_MplsVpnInterfaceConfTable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSVPNMIB_MplsVpnInterfaceConfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every interface capable of
    // supporting MPLS/BGP VPN.   Each entry in this table is meant to correspond
    // to an entry in the Interfaces Table. The type is slice of
    // MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry.
    MplsVpnInterfaceConfEntry []*MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry
}

func (mplsVpnInterfaceConfTable *MPLSVPNMIB_MplsVpnInterfaceConfTable) GetEntityData() *types.CommonEntityData {
    mplsVpnInterfaceConfTable.EntityData.YFilter = mplsVpnInterfaceConfTable.YFilter
    mplsVpnInterfaceConfTable.EntityData.YangName = "mplsVpnInterfaceConfTable"
    mplsVpnInterfaceConfTable.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnInterfaceConfTable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsVpnInterfaceConfTable.EntityData.SegmentPath = "mplsVpnInterfaceConfTable"
    mplsVpnInterfaceConfTable.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/" + mplsVpnInterfaceConfTable.EntityData.SegmentPath
    mplsVpnInterfaceConfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnInterfaceConfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnInterfaceConfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnInterfaceConfTable.EntityData.Children = types.NewOrderedMap()
    mplsVpnInterfaceConfTable.EntityData.Children.Append("mplsVpnInterfaceConfEntry", types.YChild{"MplsVpnInterfaceConfEntry", nil})
    for i := range mplsVpnInterfaceConfTable.MplsVpnInterfaceConfEntry {
        mplsVpnInterfaceConfTable.EntityData.Children.Append(types.GetSegmentPath(mplsVpnInterfaceConfTable.MplsVpnInterfaceConfEntry[i]), types.YChild{"MplsVpnInterfaceConfEntry", mplsVpnInterfaceConfTable.MplsVpnInterfaceConfEntry[i]})
    }
    mplsVpnInterfaceConfTable.EntityData.Leafs = types.NewOrderedMap()

    mplsVpnInterfaceConfTable.EntityData.YListKeys = []string {}

    return &(mplsVpnInterfaceConfTable.EntityData)
}

// MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry
// An entry in this table is created by an LSR for
// every interface capable of supporting MPLS/BGP VPN.
// 
// 
// Each entry in this table is meant to correspond to
// an entry in the Interfaces Table.
type MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfName
    MplsVpnVrfName interface{}

    // This attribute is a key. This is a unique index for an entry in the
    // MplsVPNInterfaceConfTable. A non-zero index for an entry indicates the
    // ifIndex for the corresponding interface entry in the MPLS-VPN-layer in the
    // ifTable. Note that this table does not necessarily correspond one-to-one
    // with all entries in the Interface MIB having an ifType of MPLS-layer;
    // rather, only those which are enabled for MPLS/BGP VPN functionality. The
    // type is interface{} with range: 1..2147483647.
    MplsVpnInterfaceConfIndex interface{}

    // Either the providerEdge(0) (PE) or customerEdge(1) (CE) bit MUST be set.
    // The type is MplsVpnInterfaceLabelEdgeType.
    MplsVpnInterfaceLabelEdgeType interface{}

    // Denotes whether this link participates in a carrier-of-carrier's,
    // enterprise, or inter-provider scenario. The type is
    // MplsVpnInterfaceVpnClassification.
    MplsVpnInterfaceVpnClassification interface{}

    // Denotes the route distribution protocol across the PE-CE link. Note that
    // more than one routing protocol may be enabled at the same time. The type is
    // map[string]bool.
    MplsVpnInterfaceVpnRouteDistProtocol interface{}

    // The storage type for this entry. The type is StorageType.
    MplsVpnInterfaceConfStorageType interface{}

    // The row status for this entry. This value is used to create a row in this
    // table, signifying that the specified interface is to be associated with the
    // specified interface. If this operation succeeds, the interface will have
    // been associated, otherwise the agent would not allow the association.  If
    // the agent only allows read-only operations on this table, it will create
    // entries in this table as they are created. The type is RowStatus.
    MplsVpnInterfaceConfRowStatus interface{}
}

func (mplsVpnInterfaceConfEntry *MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry) GetEntityData() *types.CommonEntityData {
    mplsVpnInterfaceConfEntry.EntityData.YFilter = mplsVpnInterfaceConfEntry.YFilter
    mplsVpnInterfaceConfEntry.EntityData.YangName = "mplsVpnInterfaceConfEntry"
    mplsVpnInterfaceConfEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnInterfaceConfEntry.EntityData.ParentYangName = "mplsVpnInterfaceConfTable"
    mplsVpnInterfaceConfEntry.EntityData.SegmentPath = "mplsVpnInterfaceConfEntry" + types.AddKeyToken(mplsVpnInterfaceConfEntry.MplsVpnVrfName, "mplsVpnVrfName") + types.AddKeyToken(mplsVpnInterfaceConfEntry.MplsVpnInterfaceConfIndex, "mplsVpnInterfaceConfIndex")
    mplsVpnInterfaceConfEntry.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/mplsVpnInterfaceConfTable/" + mplsVpnInterfaceConfEntry.EntityData.SegmentPath
    mplsVpnInterfaceConfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnInterfaceConfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnInterfaceConfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnInterfaceConfEntry.EntityData.Children = types.NewOrderedMap()
    mplsVpnInterfaceConfEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsVpnInterfaceConfEntry.EntityData.Leafs.Append("mplsVpnVrfName", types.YLeaf{"MplsVpnVrfName", mplsVpnInterfaceConfEntry.MplsVpnVrfName})
    mplsVpnInterfaceConfEntry.EntityData.Leafs.Append("mplsVpnInterfaceConfIndex", types.YLeaf{"MplsVpnInterfaceConfIndex", mplsVpnInterfaceConfEntry.MplsVpnInterfaceConfIndex})
    mplsVpnInterfaceConfEntry.EntityData.Leafs.Append("mplsVpnInterfaceLabelEdgeType", types.YLeaf{"MplsVpnInterfaceLabelEdgeType", mplsVpnInterfaceConfEntry.MplsVpnInterfaceLabelEdgeType})
    mplsVpnInterfaceConfEntry.EntityData.Leafs.Append("mplsVpnInterfaceVpnClassification", types.YLeaf{"MplsVpnInterfaceVpnClassification", mplsVpnInterfaceConfEntry.MplsVpnInterfaceVpnClassification})
    mplsVpnInterfaceConfEntry.EntityData.Leafs.Append("mplsVpnInterfaceVpnRouteDistProtocol", types.YLeaf{"MplsVpnInterfaceVpnRouteDistProtocol", mplsVpnInterfaceConfEntry.MplsVpnInterfaceVpnRouteDistProtocol})
    mplsVpnInterfaceConfEntry.EntityData.Leafs.Append("mplsVpnInterfaceConfStorageType", types.YLeaf{"MplsVpnInterfaceConfStorageType", mplsVpnInterfaceConfEntry.MplsVpnInterfaceConfStorageType})
    mplsVpnInterfaceConfEntry.EntityData.Leafs.Append("mplsVpnInterfaceConfRowStatus", types.YLeaf{"MplsVpnInterfaceConfRowStatus", mplsVpnInterfaceConfEntry.MplsVpnInterfaceConfRowStatus})

    mplsVpnInterfaceConfEntry.EntityData.YListKeys = []string {"MplsVpnVrfName", "MplsVpnInterfaceConfIndex"}

    return &(mplsVpnInterfaceConfEntry.EntityData)
}

// MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceLabelEdgeType represents (CE) bit MUST be set.
type MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceLabelEdgeType string

const (
    MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceLabelEdgeType_providerEdge MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceLabelEdgeType = "providerEdge"

    MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceLabelEdgeType_customerEdge MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceLabelEdgeType = "customerEdge"
)

// MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification represents scenario.
type MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification string

const (
    MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification_carrierOfCarrier MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification = "carrierOfCarrier"

    MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification_enterprise MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification = "enterprise"

    MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification_interProvider MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceVpnClassification = "interProvider"
)

// MPLSVPNMIB_MplsVpnVrfTable
// This table specifies per-interface MPLS/BGP VPN
// VRF Table capability and associated information.
// Entries in this table define VRF routing instances
// associated with MPLS/VPN interfaces. Note that
// multiple interfaces can belong to the same VRF
// instance. The collection of all VRF instances
// comprises an actual VPN.
type MPLSVPNMIB_MplsVpnVrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS/BGP VPN. The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry.
    MplsVpnVrfEntry []*MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry
}

func (mplsVpnVrfTable *MPLSVPNMIB_MplsVpnVrfTable) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfTable.EntityData.YFilter = mplsVpnVrfTable.YFilter
    mplsVpnVrfTable.EntityData.YangName = "mplsVpnVrfTable"
    mplsVpnVrfTable.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfTable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsVpnVrfTable.EntityData.SegmentPath = "mplsVpnVrfTable"
    mplsVpnVrfTable.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/" + mplsVpnVrfTable.EntityData.SegmentPath
    mplsVpnVrfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfTable.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfTable.EntityData.Children.Append("mplsVpnVrfEntry", types.YChild{"MplsVpnVrfEntry", nil})
    for i := range mplsVpnVrfTable.MplsVpnVrfEntry {
        mplsVpnVrfTable.EntityData.Children.Append(types.GetSegmentPath(mplsVpnVrfTable.MplsVpnVrfEntry[i]), types.YChild{"MplsVpnVrfEntry", mplsVpnVrfTable.MplsVpnVrfEntry[i]})
    }
    mplsVpnVrfTable.EntityData.Leafs = types.NewOrderedMap()

    mplsVpnVrfTable.EntityData.YListKeys = []string {}

    return &(mplsVpnVrfTable.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS/BGP VPN. The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The human-readable name of this VPN. This MAY be
    // equivalent to the RFC2685 VPN-ID. The type is string with length: 0..31.
    MplsVpnVrfName interface{}

    // The human-readable description of this VRF. The type is string.
    MplsVpnVrfDescription interface{}

    // The route distinguisher for this VRF. The type is string with length:
    // 0..256.
    MplsVpnVrfRouteDistinguisher interface{}

    // The time at which this VRF entry was created. The type is interface{} with
    // range: 0..4294967295.
    MplsVpnVrfCreationTime interface{}

    // Denotes whether a VRF is operational or not. A VRF is  up(1) when at least
    // one interface associated with the VRF, which ifOperStatus is up(1). A VRF
    // is down(2) when:  a. There does not exist at least one interface whose   
    // ifOperStatus is up(1).  b. There are no interfaces associated with the VRF.
    // The type is MplsVpnVrfOperStatus.
    MplsVpnVrfOperStatus interface{}

    // Total number of interfaces connected to this VRF with    ifOperStatus =
    // up(1).   This counter should be incremented when:  a. When the ifOperStatus
    // of one of the connected interfaces     changes from down(2) to up(1).  b.
    // When an interface with ifOperStatus = up(1) is connected    to this VRF. 
    // This counter should be decremented when:  a. When the ifOperStatus of one
    // of the connected interfaces     changes from up(1) to down(2).  b. When one
    // of the connected interfaces with     ifOperStatus = up(1) gets disconnected
    // from this VRF. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfActiveInterfaces interface{}

    // Total number of interfaces connected to this VRF  (independent of
    // ifOperStatus type). The type is interface{} with range: 0..4294967295.
    MplsVpnVrfAssociatedInterfaces interface{}

    // Denotes mid-level water marker for the number of routes which  this VRF may
    // hold. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfConfMidRouteThreshold interface{}

    // Denotes high-level water marker for the number of routes which  this VRF
    // may hold. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfConfHighRouteThreshold interface{}

    // Denotes maximum number of routes which this VRF is configured to hold. This
    // value MUST be less than or equal to mplsVrfMaxPossibleRoutes unless it is
    // set to 0. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfConfMaxRoutes interface{}

    // The value of sysUpTime at the time of the last change of this table entry,
    // which includes changes of VRF parameters defined in this table or addition
    // or deletion of interfaces associated with this VRF. The type is interface{}
    // with range: 0..4294967295.
    MplsVpnVrfConfLastChanged interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // The type is RowStatus.
    MplsVpnVrfConfRowStatus interface{}

    // The storage type for this entry. The type is StorageType.
    MplsVpnVrfConfStorageType interface{}

    // Indicates the number of illegally received labels on this VPN/VRF. The type
    // is interface{} with range: 0..4294967295.
    MplsVpnVrfSecIllegalLabelViolations interface{}

    // The number of illegally received labels above which this  notification is
    // issued. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfSecIllegalLabelRcvThresh interface{}

    // Indicates the number of routes added to this VPN/VRF over the coarse of its
    // lifetime. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfPerfRoutesAdded interface{}

    // Indicates the number of routes removed from this VPN/VRF. The type is
    // interface{} with range: 0..4294967295.
    MplsVpnVrfPerfRoutesDeleted interface{}

    // Indicates the number of routes currently used by this VRF. The type is
    // interface{} with range: 0..4294967295.
    MplsVpnVrfPerfCurrNumRoutes interface{}
}

func (mplsVpnVrfEntry *MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfEntry.EntityData.YFilter = mplsVpnVrfEntry.YFilter
    mplsVpnVrfEntry.EntityData.YangName = "mplsVpnVrfEntry"
    mplsVpnVrfEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfEntry.EntityData.ParentYangName = "mplsVpnVrfTable"
    mplsVpnVrfEntry.EntityData.SegmentPath = "mplsVpnVrfEntry" + types.AddKeyToken(mplsVpnVrfEntry.MplsVpnVrfName, "mplsVpnVrfName")
    mplsVpnVrfEntry.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/mplsVpnVrfTable/" + mplsVpnVrfEntry.EntityData.SegmentPath
    mplsVpnVrfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfEntry.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfName", types.YLeaf{"MplsVpnVrfName", mplsVpnVrfEntry.MplsVpnVrfName})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfDescription", types.YLeaf{"MplsVpnVrfDescription", mplsVpnVrfEntry.MplsVpnVrfDescription})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfRouteDistinguisher", types.YLeaf{"MplsVpnVrfRouteDistinguisher", mplsVpnVrfEntry.MplsVpnVrfRouteDistinguisher})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfCreationTime", types.YLeaf{"MplsVpnVrfCreationTime", mplsVpnVrfEntry.MplsVpnVrfCreationTime})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfOperStatus", types.YLeaf{"MplsVpnVrfOperStatus", mplsVpnVrfEntry.MplsVpnVrfOperStatus})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfActiveInterfaces", types.YLeaf{"MplsVpnVrfActiveInterfaces", mplsVpnVrfEntry.MplsVpnVrfActiveInterfaces})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfAssociatedInterfaces", types.YLeaf{"MplsVpnVrfAssociatedInterfaces", mplsVpnVrfEntry.MplsVpnVrfAssociatedInterfaces})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfConfMidRouteThreshold", types.YLeaf{"MplsVpnVrfConfMidRouteThreshold", mplsVpnVrfEntry.MplsVpnVrfConfMidRouteThreshold})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfConfHighRouteThreshold", types.YLeaf{"MplsVpnVrfConfHighRouteThreshold", mplsVpnVrfEntry.MplsVpnVrfConfHighRouteThreshold})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfConfMaxRoutes", types.YLeaf{"MplsVpnVrfConfMaxRoutes", mplsVpnVrfEntry.MplsVpnVrfConfMaxRoutes})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfConfLastChanged", types.YLeaf{"MplsVpnVrfConfLastChanged", mplsVpnVrfEntry.MplsVpnVrfConfLastChanged})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfConfRowStatus", types.YLeaf{"MplsVpnVrfConfRowStatus", mplsVpnVrfEntry.MplsVpnVrfConfRowStatus})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfConfStorageType", types.YLeaf{"MplsVpnVrfConfStorageType", mplsVpnVrfEntry.MplsVpnVrfConfStorageType})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfSecIllegalLabelViolations", types.YLeaf{"MplsVpnVrfSecIllegalLabelViolations", mplsVpnVrfEntry.MplsVpnVrfSecIllegalLabelViolations})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfSecIllegalLabelRcvThresh", types.YLeaf{"MplsVpnVrfSecIllegalLabelRcvThresh", mplsVpnVrfEntry.MplsVpnVrfSecIllegalLabelRcvThresh})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfPerfRoutesAdded", types.YLeaf{"MplsVpnVrfPerfRoutesAdded", mplsVpnVrfEntry.MplsVpnVrfPerfRoutesAdded})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfPerfRoutesDeleted", types.YLeaf{"MplsVpnVrfPerfRoutesDeleted", mplsVpnVrfEntry.MplsVpnVrfPerfRoutesDeleted})
    mplsVpnVrfEntry.EntityData.Leafs.Append("mplsVpnVrfPerfCurrNumRoutes", types.YLeaf{"MplsVpnVrfPerfCurrNumRoutes", mplsVpnVrfEntry.MplsVpnVrfPerfCurrNumRoutes})

    mplsVpnVrfEntry.EntityData.YListKeys = []string {"MplsVpnVrfName"}

    return &(mplsVpnVrfEntry.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfOperStatus represents b. There are no interfaces associated with the VRF.
type MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfOperStatus string

const (
    MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfOperStatus_up MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfOperStatus = "up"

    MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfOperStatus_down MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfOperStatus = "down"
)

// MPLSVPNMIB_MplsVpnVrfRouteTargetTable
// This table specifies per-VRF route target association.
// Each entry identifies a connectivity policy supported
// as part of a VPN.
type MPLSVPNMIB_MplsVpnVrfRouteTargetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for each route target
    // configured for a VRF supporting a MPLS/BGP VPN instance. The indexing
    // provides an ordering per-VRF instance. The type is slice of
    // MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry.
    MplsVpnVrfRouteTargetEntry []*MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry
}

func (mplsVpnVrfRouteTargetTable *MPLSVPNMIB_MplsVpnVrfRouteTargetTable) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfRouteTargetTable.EntityData.YFilter = mplsVpnVrfRouteTargetTable.YFilter
    mplsVpnVrfRouteTargetTable.EntityData.YangName = "mplsVpnVrfRouteTargetTable"
    mplsVpnVrfRouteTargetTable.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfRouteTargetTable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsVpnVrfRouteTargetTable.EntityData.SegmentPath = "mplsVpnVrfRouteTargetTable"
    mplsVpnVrfRouteTargetTable.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/" + mplsVpnVrfRouteTargetTable.EntityData.SegmentPath
    mplsVpnVrfRouteTargetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfRouteTargetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfRouteTargetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfRouteTargetTable.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfRouteTargetTable.EntityData.Children.Append("mplsVpnVrfRouteTargetEntry", types.YChild{"MplsVpnVrfRouteTargetEntry", nil})
    for i := range mplsVpnVrfRouteTargetTable.MplsVpnVrfRouteTargetEntry {
        mplsVpnVrfRouteTargetTable.EntityData.Children.Append(types.GetSegmentPath(mplsVpnVrfRouteTargetTable.MplsVpnVrfRouteTargetEntry[i]), types.YChild{"MplsVpnVrfRouteTargetEntry", mplsVpnVrfRouteTargetTable.MplsVpnVrfRouteTargetEntry[i]})
    }
    mplsVpnVrfRouteTargetTable.EntityData.Leafs = types.NewOrderedMap()

    mplsVpnVrfRouteTargetTable.EntityData.YListKeys = []string {}

    return &(mplsVpnVrfRouteTargetTable.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry
//  An entry in this table is created by an LSR for
// each route target configured for a VRF supporting
// a MPLS/BGP VPN instance. The indexing provides an
// ordering per-VRF instance.
type MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfName
    MplsVpnVrfName interface{}

    // This attribute is a key. Auxiliary index for route-targets configured for a
    // particular VRF. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfRouteTargetIndex interface{}

    // This attribute is a key. The route target export distribution type. The
    // type is MplsVpnVrfRouteTargetType.
    MplsVpnVrfRouteTargetType interface{}

    // The route target distribution policy. The type is string with length:
    // 0..256.
    MplsVpnVrfRouteTarget interface{}

    // Description of the route target. The type is string.
    MplsVpnVrfRouteTargetDescr interface{}

    // Row status for this entry. The type is RowStatus.
    MplsVpnVrfRouteTargetRowStatus interface{}
}

func (mplsVpnVrfRouteTargetEntry *MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfRouteTargetEntry.EntityData.YFilter = mplsVpnVrfRouteTargetEntry.YFilter
    mplsVpnVrfRouteTargetEntry.EntityData.YangName = "mplsVpnVrfRouteTargetEntry"
    mplsVpnVrfRouteTargetEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfRouteTargetEntry.EntityData.ParentYangName = "mplsVpnVrfRouteTargetTable"
    mplsVpnVrfRouteTargetEntry.EntityData.SegmentPath = "mplsVpnVrfRouteTargetEntry" + types.AddKeyToken(mplsVpnVrfRouteTargetEntry.MplsVpnVrfName, "mplsVpnVrfName") + types.AddKeyToken(mplsVpnVrfRouteTargetEntry.MplsVpnVrfRouteTargetIndex, "mplsVpnVrfRouteTargetIndex") + types.AddKeyToken(mplsVpnVrfRouteTargetEntry.MplsVpnVrfRouteTargetType, "mplsVpnVrfRouteTargetType")
    mplsVpnVrfRouteTargetEntry.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/mplsVpnVrfRouteTargetTable/" + mplsVpnVrfRouteTargetEntry.EntityData.SegmentPath
    mplsVpnVrfRouteTargetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfRouteTargetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfRouteTargetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfRouteTargetEntry.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfRouteTargetEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsVpnVrfRouteTargetEntry.EntityData.Leafs.Append("mplsVpnVrfName", types.YLeaf{"MplsVpnVrfName", mplsVpnVrfRouteTargetEntry.MplsVpnVrfName})
    mplsVpnVrfRouteTargetEntry.EntityData.Leafs.Append("mplsVpnVrfRouteTargetIndex", types.YLeaf{"MplsVpnVrfRouteTargetIndex", mplsVpnVrfRouteTargetEntry.MplsVpnVrfRouteTargetIndex})
    mplsVpnVrfRouteTargetEntry.EntityData.Leafs.Append("mplsVpnVrfRouteTargetType", types.YLeaf{"MplsVpnVrfRouteTargetType", mplsVpnVrfRouteTargetEntry.MplsVpnVrfRouteTargetType})
    mplsVpnVrfRouteTargetEntry.EntityData.Leafs.Append("mplsVpnVrfRouteTarget", types.YLeaf{"MplsVpnVrfRouteTarget", mplsVpnVrfRouteTargetEntry.MplsVpnVrfRouteTarget})
    mplsVpnVrfRouteTargetEntry.EntityData.Leafs.Append("mplsVpnVrfRouteTargetDescr", types.YLeaf{"MplsVpnVrfRouteTargetDescr", mplsVpnVrfRouteTargetEntry.MplsVpnVrfRouteTargetDescr})
    mplsVpnVrfRouteTargetEntry.EntityData.Leafs.Append("mplsVpnVrfRouteTargetRowStatus", types.YLeaf{"MplsVpnVrfRouteTargetRowStatus", mplsVpnVrfRouteTargetEntry.MplsVpnVrfRouteTargetRowStatus})

    mplsVpnVrfRouteTargetEntry.EntityData.YListKeys = []string {"MplsVpnVrfName", "MplsVpnVrfRouteTargetIndex", "MplsVpnVrfRouteTargetType"}

    return &(mplsVpnVrfRouteTargetEntry.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType represents The route target export distribution type.
type MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType string

const (
    MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType_import_ MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType = "import"

    MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType_export MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType = "export"

    MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType_both MPLSVPNMIB_MplsVpnVrfRouteTargetTable_MplsVpnVrfRouteTargetEntry_MplsVpnVrfRouteTargetType = "both"
)

// MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable
// Each entry in this table specifies a per-interface 
// MPLS/EBGP neighbor.
type MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS/BGP VPN. The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of
    // MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry.
    MplsVpnVrfBgpNbrAddrEntry []*MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry
}

func (mplsVpnVrfBgpNbrAddrTable *MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfBgpNbrAddrTable.EntityData.YFilter = mplsVpnVrfBgpNbrAddrTable.YFilter
    mplsVpnVrfBgpNbrAddrTable.EntityData.YangName = "mplsVpnVrfBgpNbrAddrTable"
    mplsVpnVrfBgpNbrAddrTable.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfBgpNbrAddrTable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsVpnVrfBgpNbrAddrTable.EntityData.SegmentPath = "mplsVpnVrfBgpNbrAddrTable"
    mplsVpnVrfBgpNbrAddrTable.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/" + mplsVpnVrfBgpNbrAddrTable.EntityData.SegmentPath
    mplsVpnVrfBgpNbrAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfBgpNbrAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfBgpNbrAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfBgpNbrAddrTable.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfBgpNbrAddrTable.EntityData.Children.Append("mplsVpnVrfBgpNbrAddrEntry", types.YChild{"MplsVpnVrfBgpNbrAddrEntry", nil})
    for i := range mplsVpnVrfBgpNbrAddrTable.MplsVpnVrfBgpNbrAddrEntry {
        mplsVpnVrfBgpNbrAddrTable.EntityData.Children.Append(types.GetSegmentPath(mplsVpnVrfBgpNbrAddrTable.MplsVpnVrfBgpNbrAddrEntry[i]), types.YChild{"MplsVpnVrfBgpNbrAddrEntry", mplsVpnVrfBgpNbrAddrTable.MplsVpnVrfBgpNbrAddrEntry[i]})
    }
    mplsVpnVrfBgpNbrAddrTable.EntityData.Leafs = types.NewOrderedMap()

    mplsVpnVrfBgpNbrAddrTable.EntityData.YListKeys = []string {}

    return &(mplsVpnVrfBgpNbrAddrTable.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS/BGP VPN. The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfName
    MplsVpnVrfName interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // mpls_vpn_mib.MPLSVPNMIB_MplsVpnInterfaceConfTable_MplsVpnInterfaceConfEntry_MplsVpnInterfaceConfIndex
    MplsVpnInterfaceConfIndex interface{}

    // This attribute is a key. This is a unique tertiary index for an entry in
    // the MplsVpnVrfBgpNbrAddrEntry Table. The type is interface{} with range:
    // 0..4294967295.
    MplsVpnVrfBgpNbrIndex interface{}

    // Denotes the role played by this EBGP neighbor with respect to this VRF. The
    // type is MplsVpnVrfBgpNbrRole.
    MplsVpnVrfBgpNbrRole interface{}

    // Denotes the address family of the PE address. The type is InetAddressType.
    MplsVpnVrfBgpNbrType interface{}

    // Denotes the EBGP neighbor address. The type is string with length: 0..255.
    MplsVpnVrfBgpNbrAddr interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // The type is RowStatus.
    MplsVpnVrfBgpNbrRowStatus interface{}

    // The storage type for this entry. The type is StorageType.
    MplsVpnVrfBgpNbrStorageType interface{}
}

func (mplsVpnVrfBgpNbrAddrEntry *MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfBgpNbrAddrEntry.EntityData.YFilter = mplsVpnVrfBgpNbrAddrEntry.YFilter
    mplsVpnVrfBgpNbrAddrEntry.EntityData.YangName = "mplsVpnVrfBgpNbrAddrEntry"
    mplsVpnVrfBgpNbrAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfBgpNbrAddrEntry.EntityData.ParentYangName = "mplsVpnVrfBgpNbrAddrTable"
    mplsVpnVrfBgpNbrAddrEntry.EntityData.SegmentPath = "mplsVpnVrfBgpNbrAddrEntry" + types.AddKeyToken(mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfName, "mplsVpnVrfName") + types.AddKeyToken(mplsVpnVrfBgpNbrAddrEntry.MplsVpnInterfaceConfIndex, "mplsVpnInterfaceConfIndex") + types.AddKeyToken(mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfBgpNbrIndex, "mplsVpnVrfBgpNbrIndex")
    mplsVpnVrfBgpNbrAddrEntry.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/mplsVpnVrfBgpNbrAddrTable/" + mplsVpnVrfBgpNbrAddrEntry.EntityData.SegmentPath
    mplsVpnVrfBgpNbrAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfBgpNbrAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfBgpNbrAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfBgpNbrAddrEntry.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnVrfName", types.YLeaf{"MplsVpnVrfName", mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfName})
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnInterfaceConfIndex", types.YLeaf{"MplsVpnInterfaceConfIndex", mplsVpnVrfBgpNbrAddrEntry.MplsVpnInterfaceConfIndex})
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnVrfBgpNbrIndex", types.YLeaf{"MplsVpnVrfBgpNbrIndex", mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfBgpNbrIndex})
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnVrfBgpNbrRole", types.YLeaf{"MplsVpnVrfBgpNbrRole", mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfBgpNbrRole})
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnVrfBgpNbrType", types.YLeaf{"MplsVpnVrfBgpNbrType", mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfBgpNbrType})
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnVrfBgpNbrAddr", types.YLeaf{"MplsVpnVrfBgpNbrAddr", mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfBgpNbrAddr})
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnVrfBgpNbrRowStatus", types.YLeaf{"MplsVpnVrfBgpNbrRowStatus", mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfBgpNbrRowStatus})
    mplsVpnVrfBgpNbrAddrEntry.EntityData.Leafs.Append("mplsVpnVrfBgpNbrStorageType", types.YLeaf{"MplsVpnVrfBgpNbrStorageType", mplsVpnVrfBgpNbrAddrEntry.MplsVpnVrfBgpNbrStorageType})

    mplsVpnVrfBgpNbrAddrEntry.EntityData.YListKeys = []string {"MplsVpnVrfName", "MplsVpnInterfaceConfIndex", "MplsVpnVrfBgpNbrIndex"}

    return &(mplsVpnVrfBgpNbrAddrEntry.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry_MplsVpnVrfBgpNbrRole represents with respect to this VRF.
type MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry_MplsVpnVrfBgpNbrRole string

const (
    MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry_MplsVpnVrfBgpNbrRole_ce MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry_MplsVpnVrfBgpNbrRole = "ce"

    MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry_MplsVpnVrfBgpNbrRole_pe MPLSVPNMIB_MplsVpnVrfBgpNbrAddrTable_MplsVpnVrfBgpNbrAddrEntry_MplsVpnVrfBgpNbrRole = "pe"
)

// MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable
// This table specifies per-VRF vpnv4 multi-protocol
// prefixes supported by BGP.
type MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every BGP prefix associated
    // with a VRF supporting a  MPLS/BGP VPN. The indexing provides an ordering of
    // BGP prefixes per VRF. The type is slice of
    // MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry.
    MplsVpnVrfBgpNbrPrefixEntry []*MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry
}

func (mplsVpnVrfBgpNbrPrefixTable *MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfBgpNbrPrefixTable.EntityData.YFilter = mplsVpnVrfBgpNbrPrefixTable.YFilter
    mplsVpnVrfBgpNbrPrefixTable.EntityData.YangName = "mplsVpnVrfBgpNbrPrefixTable"
    mplsVpnVrfBgpNbrPrefixTable.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfBgpNbrPrefixTable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsVpnVrfBgpNbrPrefixTable.EntityData.SegmentPath = "mplsVpnVrfBgpNbrPrefixTable"
    mplsVpnVrfBgpNbrPrefixTable.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/" + mplsVpnVrfBgpNbrPrefixTable.EntityData.SegmentPath
    mplsVpnVrfBgpNbrPrefixTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfBgpNbrPrefixTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfBgpNbrPrefixTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfBgpNbrPrefixTable.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfBgpNbrPrefixTable.EntityData.Children.Append("mplsVpnVrfBgpNbrPrefixEntry", types.YChild{"MplsVpnVrfBgpNbrPrefixEntry", nil})
    for i := range mplsVpnVrfBgpNbrPrefixTable.MplsVpnVrfBgpNbrPrefixEntry {
        mplsVpnVrfBgpNbrPrefixTable.EntityData.Children.Append(types.GetSegmentPath(mplsVpnVrfBgpNbrPrefixTable.MplsVpnVrfBgpNbrPrefixEntry[i]), types.YChild{"MplsVpnVrfBgpNbrPrefixEntry", mplsVpnVrfBgpNbrPrefixTable.MplsVpnVrfBgpNbrPrefixEntry[i]})
    }
    mplsVpnVrfBgpNbrPrefixTable.EntityData.Leafs = types.NewOrderedMap()

    mplsVpnVrfBgpNbrPrefixTable.EntityData.YListKeys = []string {}

    return &(mplsVpnVrfBgpNbrPrefixTable.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry
// An entry in this table is created by an LSR for
// every BGP prefix associated with a VRF supporting a 
// MPLS/BGP VPN. The indexing provides an ordering of 
// BGP prefixes per VRF.
type MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfName
    MplsVpnVrfName interface{}

    // This attribute is a key. An IP address prefix in the Network Layer
    // Reachability Information field.  This object is an IP address containing
    // the prefix with length specified by mplsVpnVrfBgpPathAttrIpAddrPrefixLen.
    // Any bits beyond the length specified by
    // mplsVpnVrfBgpPathAttrIpAddrPrefixLen are zeroed. The type is string with
    // length: 0..255.
    MplsVpnVrfBgpPathAttrIpAddrPrefix interface{}

    // This attribute is a key. Length in bits of the IP address prefix in the
    // Network Layer Reachability Information field. The type is interface{} with
    // range: 0..32.
    MplsVpnVrfBgpPathAttrIpAddrPrefixLen interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with length: 0..255.
    MplsVpnVrfBgpPathAttrPeer interface{}

    // The ultimate origin of the path information. The type is
    // MplsVpnVrfBgpPathAttrOrigin.
    MplsVpnVrfBgpPathAttrOrigin interface{}

    // The sequence of AS path segments.  Each AS path segment is represented by a
    // triple <type, length, value>.   The type is a 1-octet field which has two 
    // possible values:      1      AS_SET: unordered set of ASs a            
    // route in the UPDATE             message has traversed      2     
    // AS_SEQUENCE: ordered set of ASs             a route in the UPDATE          
    // message has traversed.             The length is a 1-octet field containing
    // the               number of ASs in the value field.              The value
    // field contains one or more AS             numbers, each AS is represented
    // in the octet             string as a pair of octets according to the       
    // following algorithm:              first-byte-of-pair = ASNumber / 256;     
    // second-byte-of-pair = ASNumber & 255;. The type is string with length:
    // 2..255.
    MplsVpnVrfBgpPathAttrASPathSegment interface{}

    // The address of the border router that should be used for the destination
    // network. The type is string with length: 0..255.
    MplsVpnVrfBgpPathAttrNextHop interface{}

    // This metric is used to discriminate between multiple exit points to an
    // adjacent autonomous system.  A value of -1 indicates the absence of this
    // attribute. The type is interface{} with range: -1..2147483647.
    MplsVpnVrfBgpPathAttrMultiExitDisc interface{}

    // The originating BGP4 speaker's degree of preference for an advertised
    // route.  A value of -1 indicates the absence of this attribute. The type is
    // interface{} with range: -1..2147483647.
    MplsVpnVrfBgpPathAttrLocalPref interface{}

    // Whether or not the local system has selected a less specific route without
    // selecting a more specific route. The type is
    // MplsVpnVrfBgpPathAttrAtomicAggregate.
    MplsVpnVrfBgpPathAttrAtomicAggregate interface{}

    // The AS number of the last BGP4 speaker that performed route aggregation.  A
    // value of zero (0) indicates the absence of this attribute. The type is
    // interface{} with range: 0..65535.
    MplsVpnVrfBgpPathAttrAggregatorAS interface{}

    // The IP address of the last BGP4 speaker that performed route aggregation. 
    // A value of 0.0.0.0 indicates the absence of this attribute. The type is
    // string with length: 0..255.
    MplsVpnVrfBgpPathAttrAggregatorAddr interface{}

    // The degree of preference calculated by the receiving BGP4 speaker for an
    // advertised route.  A value of -1 indicates the absence of this attribute.
    // The type is interface{} with range: -1..2147483647.
    MplsVpnVrfBgpPathAttrCalcLocalPref interface{}

    // An indication of whether or not this route was chosen as the best BGP4
    // route. The type is MplsVpnVrfBgpPathAttrBest.
    MplsVpnVrfBgpPathAttrBest interface{}

    // One or more path attributes not understood by this BGP4 speaker.  Size zero
    // (0) indicates the absence of such attribute(s).  Octets beyond the maximum
    // size, if any, are not recorded by this object. The type is string with
    // length: 0..255.
    MplsVpnVrfBgpPathAttrUnknown interface{}
}

func (mplsVpnVrfBgpNbrPrefixEntry *MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.YFilter = mplsVpnVrfBgpNbrPrefixEntry.YFilter
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.YangName = "mplsVpnVrfBgpNbrPrefixEntry"
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.ParentYangName = "mplsVpnVrfBgpNbrPrefixTable"
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.SegmentPath = "mplsVpnVrfBgpNbrPrefixEntry" + types.AddKeyToken(mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfName, "mplsVpnVrfName") + types.AddKeyToken(mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrIpAddrPrefix, "mplsVpnVrfBgpPathAttrIpAddrPrefix") + types.AddKeyToken(mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrIpAddrPrefixLen, "mplsVpnVrfBgpPathAttrIpAddrPrefixLen") + types.AddKeyToken(mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrPeer, "mplsVpnVrfBgpPathAttrPeer")
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/mplsVpnVrfBgpNbrPrefixTable/" + mplsVpnVrfBgpNbrPrefixEntry.EntityData.SegmentPath
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfName", types.YLeaf{"MplsVpnVrfName", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfName})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrIpAddrPrefix", types.YLeaf{"MplsVpnVrfBgpPathAttrIpAddrPrefix", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrIpAddrPrefix})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrIpAddrPrefixLen", types.YLeaf{"MplsVpnVrfBgpPathAttrIpAddrPrefixLen", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrIpAddrPrefixLen})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrPeer", types.YLeaf{"MplsVpnVrfBgpPathAttrPeer", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrPeer})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrOrigin", types.YLeaf{"MplsVpnVrfBgpPathAttrOrigin", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrOrigin})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrASPathSegment", types.YLeaf{"MplsVpnVrfBgpPathAttrASPathSegment", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrASPathSegment})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrNextHop", types.YLeaf{"MplsVpnVrfBgpPathAttrNextHop", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrNextHop})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrMultiExitDisc", types.YLeaf{"MplsVpnVrfBgpPathAttrMultiExitDisc", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrMultiExitDisc})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrLocalPref", types.YLeaf{"MplsVpnVrfBgpPathAttrLocalPref", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrLocalPref})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrAtomicAggregate", types.YLeaf{"MplsVpnVrfBgpPathAttrAtomicAggregate", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrAtomicAggregate})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrAggregatorAS", types.YLeaf{"MplsVpnVrfBgpPathAttrAggregatorAS", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrAggregatorAS})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrAggregatorAddr", types.YLeaf{"MplsVpnVrfBgpPathAttrAggregatorAddr", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrAggregatorAddr})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrCalcLocalPref", types.YLeaf{"MplsVpnVrfBgpPathAttrCalcLocalPref", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrCalcLocalPref})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrBest", types.YLeaf{"MplsVpnVrfBgpPathAttrBest", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrBest})
    mplsVpnVrfBgpNbrPrefixEntry.EntityData.Leafs.Append("mplsVpnVrfBgpPathAttrUnknown", types.YLeaf{"MplsVpnVrfBgpPathAttrUnknown", mplsVpnVrfBgpNbrPrefixEntry.MplsVpnVrfBgpPathAttrUnknown})

    mplsVpnVrfBgpNbrPrefixEntry.EntityData.YListKeys = []string {"MplsVpnVrfName", "MplsVpnVrfBgpPathAttrIpAddrPrefix", "MplsVpnVrfBgpPathAttrIpAddrPrefixLen", "MplsVpnVrfBgpPathAttrPeer"}

    return &(mplsVpnVrfBgpNbrPrefixEntry.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrAtomicAggregate represents selecting a more specific route.
type MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrAtomicAggregate string

const (
    MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrAtomicAggregate_lessSpecificRrouteNotSelected MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrAtomicAggregate = "lessSpecificRrouteNotSelected"

    MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrAtomicAggregate_lessSpecificRouteSelected MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrAtomicAggregate = "lessSpecificRouteSelected"
)

// MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrBest represents was chosen as the best BGP4 route.
type MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrBest string

const (
    MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrBest_false_ MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrBest = "false"

    MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrBest_true_ MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrBest = "true"
)

// MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin represents information.
type MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin string

const (
    MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin_igp MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin = "igp"

    MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin_egp MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin = "egp"

    MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin_incomplete MPLSVPNMIB_MplsVpnVrfBgpNbrPrefixTable_MplsVpnVrfBgpNbrPrefixEntry_MplsVpnVrfBgpPathAttrOrigin = "incomplete"
)

// MPLSVPNMIB_MplsVpnVrfRouteTable
// This table specifies per-interface MPLS/BGP VPN VRF Table
// routing information. Entries in this table define VRF routing
// entries associated with the specified MPLS/VPN interfaces. Note
// that this table contains both BGP and IGP routes, as both may
// appear in the same VRF.
type MPLSVPNMIB_MplsVpnVrfRouteTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every route present
    // configured (either dynamically or statically) within the context of a
    // specific VRF capable of supporting MPLS/BGP VPN. The indexing provides an
    // ordering of VRFs per-VPN interface. The type is slice of
    // MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry.
    MplsVpnVrfRouteEntry []*MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry
}

func (mplsVpnVrfRouteTable *MPLSVPNMIB_MplsVpnVrfRouteTable) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfRouteTable.EntityData.YFilter = mplsVpnVrfRouteTable.YFilter
    mplsVpnVrfRouteTable.EntityData.YangName = "mplsVpnVrfRouteTable"
    mplsVpnVrfRouteTable.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfRouteTable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsVpnVrfRouteTable.EntityData.SegmentPath = "mplsVpnVrfRouteTable"
    mplsVpnVrfRouteTable.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/" + mplsVpnVrfRouteTable.EntityData.SegmentPath
    mplsVpnVrfRouteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfRouteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfRouteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfRouteTable.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfRouteTable.EntityData.Children.Append("mplsVpnVrfRouteEntry", types.YChild{"MplsVpnVrfRouteEntry", nil})
    for i := range mplsVpnVrfRouteTable.MplsVpnVrfRouteEntry {
        mplsVpnVrfRouteTable.EntityData.Children.Append(types.GetSegmentPath(mplsVpnVrfRouteTable.MplsVpnVrfRouteEntry[i]), types.YChild{"MplsVpnVrfRouteEntry", mplsVpnVrfRouteTable.MplsVpnVrfRouteEntry[i]})
    }
    mplsVpnVrfRouteTable.EntityData.Leafs = types.NewOrderedMap()

    mplsVpnVrfRouteTable.EntityData.YListKeys = []string {}

    return &(mplsVpnVrfRouteTable.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry
// An entry in this table is created by an LSR for every route
// present configured (either dynamically or statically) within
// the context of a specific VRF capable of supporting MPLS/BGP
// VPN. The indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_MplsVpnVrfTable_MplsVpnVrfEntry_MplsVpnVrfName
    MplsVpnVrfName interface{}

    // This attribute is a key. The destination IP address of this route. This
    // object may not take a Multicast (Class D) address value.  Any assignment
    // (implicit or otherwise) of an instance of this object to a value x must be
    // rejected if the bit-wise logical-AND of x with the value of the
    // corresponding instance of the mplsVpnVrfRouteMask object is not equal to x.
    // The type is string with length: 0..255.
    MplsVpnVrfRouteDest interface{}

    // This attribute is a key. Indicate the mask to be logical-ANDed with the
    // destination  address  before  being compared to the value  in  the 
    // mplsVpnVrfRouteDest field. For those  systems  that  do  not support
    // arbitrary subnet masks, an agent constructs the value of the
    // mplsVpnVrfRouteMask by reference   to the IP Address Class.  Any assignment
    // (implicit or otherwise) of an instance of this object to a value x must be
    // rejected if the bit-wise logical-AND of x with the value of the
    // corresponding instance of the mplsVpnVrfRouteDest object is not equal to
    // mplsVpnVrfRouteDest. The type is string with length: 0..255.
    MplsVpnVrfRouteMask interface{}

    // This attribute is a key. The IP TOS Field is used to specify the policy to
    // be applied to this route.  The encoding of IP TOS is as specified  by  the 
    // following convention. Zero indicates the default path if no more specific
    // policy applies.  +-----+-----+-----+-----+-----+-----+-----+-----+ |       
    // |                       |     | |   PRECEDENCE    |    TYPE OF SERVICE    |
    // 0  | |                 |                       |     |
    // +-----+-----+-----+-----+-----+-----+-----+-----+             IP TOS       
    // IP TOS       Field     Policy      Field     Policy       Contents    Code 
    // Contents    Code       0 0 0 0  ==>   0      0 0 0 1  ==>   2       0 0 1 0
    // ==>   4      0 0 1 1  ==>   6       0 1 0 0  ==>   8      0 1 0 1  ==>  10 
    // 0 1 1 0  ==>  12      0 1 1 1  ==>  14       1 0 0 0  ==>  16      1 0 0 1 
    // ==>  18       1 0 1 0  ==>  20      1 0 1 1  ==>  22       1 1 0 0  ==>  24
    // 1 1 0 1  ==>  26       1 1 1 0  ==>  28      1 1 1 1  ==>  30. The type is
    // interface{} with range: 0..4294967295.
    MplsVpnVrfRouteTos interface{}

    // This attribute is a key. On remote routes, the address of the next system
    // en route; Otherwise, 0.0.0.0. . The type is string with length: 0..255.
    MplsVpnVrfRouteNextHop interface{}

    // The address type of the mplsVpnVrfRouteDest entry. The type is
    // InetAddressType.
    MplsVpnVrfRouteDestAddrType interface{}

    // The address type of mplsVpnVrfRouteMask. The type is InetAddressType.
    MplsVpnVrfRouteMaskAddrType interface{}

    // The address type of the mplsVpnVrfRouteNextHopAddrType object. The type is
    // InetAddressType.
    MplsVpnVrfRouteNextHopAddrType interface{}

    // The ifIndex value that identifies the local interface  through  which  the
    // next hop of this route should be reached. The type is interface{} with
    // range: 0..2147483647.
    MplsVpnVrfRouteIfIndex interface{}

    // The type of route.  Note that local(3)  refers to a route for which the
    // next hop is the final destination; remote(4) refers to a route for that the
    // next  hop is not the final destination. Routes which do not result in
    // traffic forwarding or rejection should not be displayed even if the
    // implementation keeps them stored internally.  reject (2) refers to a route
    // which, if matched, discards the message as unreachable. This is used in
    // some protocols as a means of correctly aggregating routes. The type is
    // MplsVpnVrfRouteType.
    MplsVpnVrfRouteType interface{}

    // The routing mechanism via which this route was learned.  Inclusion of
    // values for gateway rout- ing protocols is not  intended  to  imply  that
    // hosts should support those protocols. The type is MplsVpnVrfRouteProto.
    MplsVpnVrfRouteProto interface{}

    // The number of seconds since this route was last updated or otherwise
    // determined to be correct. Note that no semantics of `too old' can be
    // implied except through knowledge of the routing protocol by which the route
    // was learned. The type is interface{} with range: 0..4294967295.
    MplsVpnVrfRouteAge interface{}

    // A reference to MIB definitions specific to the particular routing protocol
    // which is responsi-   ble for this route, as determined by the  value
    // specified  in the route's mplsVpnVrfRouteProto value. If this information
    // is not present, its value SHOULD be set to the OBJECT IDENTIFIER { 0 0 },
    // which is a syntactically valid object identif-ier, and any implementation
    // conforming to ASN.1 and the Basic Encoding Rules must be able to generate
    // and recognize this value. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    MplsVpnVrfRouteInfo interface{}

    // The Autonomous System Number of the Next Hop. The semantics of this object
    // are determined by the routing-protocol specified in the route's
    // mplsVpnVrfRouteProto value. When this object is unknown or not relevant its
    // value should be set to zero. The type is interface{} with range:
    // 0..4294967295.
    MplsVpnVrfRouteNextHopAS interface{}

    // The primary routing metric for this route. The semantics of this metric are
    // determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    MplsVpnVrfRouteMetric1 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    MplsVpnVrfRouteMetric2 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    MplsVpnVrfRouteMetric3 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    MplsVpnVrfRouteMetric4 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    MplsVpnVrfRouteMetric5 interface{}

    // Row status for this table. It is used according to row installation and
    // removal conventions. The type is RowStatus.
    MplsVpnVrfRouteRowStatus interface{}

    // Storage type value. The type is StorageType.
    MplsVpnVrfRouteStorageType interface{}
}

func (mplsVpnVrfRouteEntry *MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry) GetEntityData() *types.CommonEntityData {
    mplsVpnVrfRouteEntry.EntityData.YFilter = mplsVpnVrfRouteEntry.YFilter
    mplsVpnVrfRouteEntry.EntityData.YangName = "mplsVpnVrfRouteEntry"
    mplsVpnVrfRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    mplsVpnVrfRouteEntry.EntityData.ParentYangName = "mplsVpnVrfRouteTable"
    mplsVpnVrfRouteEntry.EntityData.SegmentPath = "mplsVpnVrfRouteEntry" + types.AddKeyToken(mplsVpnVrfRouteEntry.MplsVpnVrfName, "mplsVpnVrfName") + types.AddKeyToken(mplsVpnVrfRouteEntry.MplsVpnVrfRouteDest, "mplsVpnVrfRouteDest") + types.AddKeyToken(mplsVpnVrfRouteEntry.MplsVpnVrfRouteMask, "mplsVpnVrfRouteMask") + types.AddKeyToken(mplsVpnVrfRouteEntry.MplsVpnVrfRouteTos, "mplsVpnVrfRouteTos") + types.AddKeyToken(mplsVpnVrfRouteEntry.MplsVpnVrfRouteNextHop, "mplsVpnVrfRouteNextHop")
    mplsVpnVrfRouteEntry.EntityData.AbsolutePath = "MPLS-VPN-MIB:MPLS-VPN-MIB/mplsVpnVrfRouteTable/" + mplsVpnVrfRouteEntry.EntityData.SegmentPath
    mplsVpnVrfRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsVpnVrfRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsVpnVrfRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsVpnVrfRouteEntry.EntityData.Children = types.NewOrderedMap()
    mplsVpnVrfRouteEntry.EntityData.Leafs = types.NewOrderedMap()
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfName", types.YLeaf{"MplsVpnVrfName", mplsVpnVrfRouteEntry.MplsVpnVrfName})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteDest", types.YLeaf{"MplsVpnVrfRouteDest", mplsVpnVrfRouteEntry.MplsVpnVrfRouteDest})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteMask", types.YLeaf{"MplsVpnVrfRouteMask", mplsVpnVrfRouteEntry.MplsVpnVrfRouteMask})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteTos", types.YLeaf{"MplsVpnVrfRouteTos", mplsVpnVrfRouteEntry.MplsVpnVrfRouteTos})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteNextHop", types.YLeaf{"MplsVpnVrfRouteNextHop", mplsVpnVrfRouteEntry.MplsVpnVrfRouteNextHop})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteDestAddrType", types.YLeaf{"MplsVpnVrfRouteDestAddrType", mplsVpnVrfRouteEntry.MplsVpnVrfRouteDestAddrType})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteMaskAddrType", types.YLeaf{"MplsVpnVrfRouteMaskAddrType", mplsVpnVrfRouteEntry.MplsVpnVrfRouteMaskAddrType})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteNextHopAddrType", types.YLeaf{"MplsVpnVrfRouteNextHopAddrType", mplsVpnVrfRouteEntry.MplsVpnVrfRouteNextHopAddrType})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteIfIndex", types.YLeaf{"MplsVpnVrfRouteIfIndex", mplsVpnVrfRouteEntry.MplsVpnVrfRouteIfIndex})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteType", types.YLeaf{"MplsVpnVrfRouteType", mplsVpnVrfRouteEntry.MplsVpnVrfRouteType})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteProto", types.YLeaf{"MplsVpnVrfRouteProto", mplsVpnVrfRouteEntry.MplsVpnVrfRouteProto})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteAge", types.YLeaf{"MplsVpnVrfRouteAge", mplsVpnVrfRouteEntry.MplsVpnVrfRouteAge})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteInfo", types.YLeaf{"MplsVpnVrfRouteInfo", mplsVpnVrfRouteEntry.MplsVpnVrfRouteInfo})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteNextHopAS", types.YLeaf{"MplsVpnVrfRouteNextHopAS", mplsVpnVrfRouteEntry.MplsVpnVrfRouteNextHopAS})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteMetric1", types.YLeaf{"MplsVpnVrfRouteMetric1", mplsVpnVrfRouteEntry.MplsVpnVrfRouteMetric1})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteMetric2", types.YLeaf{"MplsVpnVrfRouteMetric2", mplsVpnVrfRouteEntry.MplsVpnVrfRouteMetric2})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteMetric3", types.YLeaf{"MplsVpnVrfRouteMetric3", mplsVpnVrfRouteEntry.MplsVpnVrfRouteMetric3})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteMetric4", types.YLeaf{"MplsVpnVrfRouteMetric4", mplsVpnVrfRouteEntry.MplsVpnVrfRouteMetric4})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteMetric5", types.YLeaf{"MplsVpnVrfRouteMetric5", mplsVpnVrfRouteEntry.MplsVpnVrfRouteMetric5})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteRowStatus", types.YLeaf{"MplsVpnVrfRouteRowStatus", mplsVpnVrfRouteEntry.MplsVpnVrfRouteRowStatus})
    mplsVpnVrfRouteEntry.EntityData.Leafs.Append("mplsVpnVrfRouteStorageType", types.YLeaf{"MplsVpnVrfRouteStorageType", mplsVpnVrfRouteEntry.MplsVpnVrfRouteStorageType})

    mplsVpnVrfRouteEntry.EntityData.YListKeys = []string {"MplsVpnVrfName", "MplsVpnVrfRouteDest", "MplsVpnVrfRouteMask", "MplsVpnVrfRouteTos", "MplsVpnVrfRouteNextHop"}

    return &(mplsVpnVrfRouteEntry.EntityData)
}

// MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto represents hosts should support those protocols.
type MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto string

const (
    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_other MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "other"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_local MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "local"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_netmgmt MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "netmgmt"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_icmp MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "icmp"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_egp MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "egp"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_ggp MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "ggp"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_hello MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "hello"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_rip MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "rip"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_isIs MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "isIs"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_esIs MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "esIs"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_ciscoIgrp MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "ciscoIgrp"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_bbnSpfIgp MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "bbnSpfIgp"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_ospf MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "ospf"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_bgp MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "bgp"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_idpr MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "idpr"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto_ciscoEigrp MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteProto = "ciscoEigrp"
)

// MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType represents routes.
type MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType string

const (
    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType_other MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType = "other"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType_reject MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType = "reject"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType_local MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType = "local"

    MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType_remote MPLSVPNMIB_MplsVpnVrfRouteTable_MplsVpnVrfRouteEntry_MplsVpnVrfRouteType = "remote"
)

