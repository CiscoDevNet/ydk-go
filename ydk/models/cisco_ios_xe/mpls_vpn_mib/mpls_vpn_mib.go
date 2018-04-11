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

    
    Mplsvpnscalars MPLSVPNMIB_Mplsvpnscalars

    // This table specifies per-interface MPLS capability and associated
    // information.
    Mplsvpninterfaceconftable MPLSVPNMIB_Mplsvpninterfaceconftable

    // This table specifies per-interface MPLS/BGP VPN VRF Table capability and
    // associated information. Entries in this table define VRF routing instances
    // associated with MPLS/VPN interfaces. Note that multiple interfaces can
    // belong to the same VRF instance. The collection of all VRF instances
    // comprises an actual VPN.
    Mplsvpnvrftable MPLSVPNMIB_Mplsvpnvrftable

    // This table specifies per-VRF route target association. Each entry
    // identifies a connectivity policy supported as part of a VPN.
    Mplsvpnvrfroutetargettable MPLSVPNMIB_Mplsvpnvrfroutetargettable

    // Each entry in this table specifies a per-interface  MPLS/EBGP neighbor.
    Mplsvpnvrfbgpnbraddrtable MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable

    // This table specifies per-VRF vpnv4 multi-protocol prefixes supported by
    // BGP.
    Mplsvpnvrfbgpnbrprefixtable MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable

    // This table specifies per-interface MPLS/BGP VPN VRF Table routing
    // information. Entries in this table define VRF routing entries associated
    // with the specified MPLS/VPN interfaces. Note that this table contains both
    // BGP and IGP routes, as both may appear in the same VRF.
    Mplsvpnvrfroutetable MPLSVPNMIB_Mplsvpnvrfroutetable
}

func (mPLSVPNMIB *MPLSVPNMIB) GetEntityData() *types.CommonEntityData {
    mPLSVPNMIB.EntityData.YFilter = mPLSVPNMIB.YFilter
    mPLSVPNMIB.EntityData.YangName = "MPLS-VPN-MIB"
    mPLSVPNMIB.EntityData.BundleName = "cisco_ios_xe"
    mPLSVPNMIB.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mPLSVPNMIB.EntityData.SegmentPath = "MPLS-VPN-MIB:MPLS-VPN-MIB"
    mPLSVPNMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mPLSVPNMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mPLSVPNMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mPLSVPNMIB.EntityData.Children = make(map[string]types.YChild)
    mPLSVPNMIB.EntityData.Children["mplsVpnScalars"] = types.YChild{"Mplsvpnscalars", &mPLSVPNMIB.Mplsvpnscalars}
    mPLSVPNMIB.EntityData.Children["mplsVpnInterfaceConfTable"] = types.YChild{"Mplsvpninterfaceconftable", &mPLSVPNMIB.Mplsvpninterfaceconftable}
    mPLSVPNMIB.EntityData.Children["mplsVpnVrfTable"] = types.YChild{"Mplsvpnvrftable", &mPLSVPNMIB.Mplsvpnvrftable}
    mPLSVPNMIB.EntityData.Children["mplsVpnVrfRouteTargetTable"] = types.YChild{"Mplsvpnvrfroutetargettable", &mPLSVPNMIB.Mplsvpnvrfroutetargettable}
    mPLSVPNMIB.EntityData.Children["mplsVpnVrfBgpNbrAddrTable"] = types.YChild{"Mplsvpnvrfbgpnbraddrtable", &mPLSVPNMIB.Mplsvpnvrfbgpnbraddrtable}
    mPLSVPNMIB.EntityData.Children["mplsVpnVrfBgpNbrPrefixTable"] = types.YChild{"Mplsvpnvrfbgpnbrprefixtable", &mPLSVPNMIB.Mplsvpnvrfbgpnbrprefixtable}
    mPLSVPNMIB.EntityData.Children["mplsVpnVrfRouteTable"] = types.YChild{"Mplsvpnvrfroutetable", &mPLSVPNMIB.Mplsvpnvrfroutetable}
    mPLSVPNMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mPLSVPNMIB.EntityData)
}

// MPLSVPNMIB_Mplsvpnscalars
type MPLSVPNMIB_Mplsvpnscalars struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of VRFs which are configured on this node. The type is
    // interface{} with range: 0..4294967295.
    Mplsvpnconfiguredvrfs interface{}

    // The number of VRFs which are active on this node. That is, those VRFs whose
    // corresponding mplsVpnVrfOperStatus  object value is equal to operational
    // (1). The type is interface{} with range: 0..4294967295.
    Mplsvpnactivevrfs interface{}

    // Total number of interfaces connected to a VRF. The type is interface{} with
    // range: 0..4294967295.
    Mplsvpnconnectedinterfaces interface{}

    // If this object is true, then it enables the generation of all notifications
    // defined in  this MIB. The type is bool.
    Mplsvpnnotificationenable interface{}

    // Denotes maximum number of routes which the device will allow all VRFs
    // jointly to hold. If this value is set to 0, this indicates that the device
    // is  unable to determine the absolute maximum. In this case, the configured
    // maximum MAY not actually be allowed by the device. The type is interface{}
    // with range: 0..4294967295.
    Mplsvpnvrfconfmaxpossibleroutes interface{}
}

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetEntityData() *types.CommonEntityData {
    mplsvpnscalars.EntityData.YFilter = mplsvpnscalars.YFilter
    mplsvpnscalars.EntityData.YangName = "mplsVpnScalars"
    mplsvpnscalars.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnscalars.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsvpnscalars.EntityData.SegmentPath = "mplsVpnScalars"
    mplsvpnscalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnscalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnscalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnscalars.EntityData.Children = make(map[string]types.YChild)
    mplsvpnscalars.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsvpnscalars.EntityData.Leafs["mplsVpnConfiguredVrfs"] = types.YLeaf{"Mplsvpnconfiguredvrfs", mplsvpnscalars.Mplsvpnconfiguredvrfs}
    mplsvpnscalars.EntityData.Leafs["mplsVpnActiveVrfs"] = types.YLeaf{"Mplsvpnactivevrfs", mplsvpnscalars.Mplsvpnactivevrfs}
    mplsvpnscalars.EntityData.Leafs["mplsVpnConnectedInterfaces"] = types.YLeaf{"Mplsvpnconnectedinterfaces", mplsvpnscalars.Mplsvpnconnectedinterfaces}
    mplsvpnscalars.EntityData.Leafs["mplsVpnNotificationEnable"] = types.YLeaf{"Mplsvpnnotificationenable", mplsvpnscalars.Mplsvpnnotificationenable}
    mplsvpnscalars.EntityData.Leafs["mplsVpnVrfConfMaxPossibleRoutes"] = types.YLeaf{"Mplsvpnvrfconfmaxpossibleroutes", mplsvpnscalars.Mplsvpnvrfconfmaxpossibleroutes}
    return &(mplsvpnscalars.EntityData)
}

// MPLSVPNMIB_Mplsvpninterfaceconftable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSVPNMIB_Mplsvpninterfaceconftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every interface capable of
    // supporting MPLS/BGP VPN.   Each entry in this table is meant to correspond
    // to an entry in the Interfaces Table. The type is slice of
    // MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry.
    Mplsvpninterfaceconfentry []MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry
}

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetEntityData() *types.CommonEntityData {
    mplsvpninterfaceconftable.EntityData.YFilter = mplsvpninterfaceconftable.YFilter
    mplsvpninterfaceconftable.EntityData.YangName = "mplsVpnInterfaceConfTable"
    mplsvpninterfaceconftable.EntityData.BundleName = "cisco_ios_xe"
    mplsvpninterfaceconftable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsvpninterfaceconftable.EntityData.SegmentPath = "mplsVpnInterfaceConfTable"
    mplsvpninterfaceconftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpninterfaceconftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpninterfaceconftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpninterfaceconftable.EntityData.Children = make(map[string]types.YChild)
    mplsvpninterfaceconftable.EntityData.Children["mplsVpnInterfaceConfEntry"] = types.YChild{"Mplsvpninterfaceconfentry", nil}
    for i := range mplsvpninterfaceconftable.Mplsvpninterfaceconfentry {
        mplsvpninterfaceconftable.EntityData.Children[types.GetSegmentPath(&mplsvpninterfaceconftable.Mplsvpninterfaceconfentry[i])] = types.YChild{"Mplsvpninterfaceconfentry", &mplsvpninterfaceconftable.Mplsvpninterfaceconfentry[i]}
    }
    mplsvpninterfaceconftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsvpninterfaceconftable.EntityData)
}

// MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry
// An entry in this table is created by an LSR for
// every interface capable of supporting MPLS/BGP VPN.
// 
// 
// Each entry in this table is meant to correspond to
// an entry in the Interfaces Table.
type MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfname
    Mplsvpnvrfname interface{}

    // This attribute is a key. This is a unique index for an entry in the
    // MplsVPNInterfaceConfTable. A non-zero index for an entry indicates the
    // ifIndex for the corresponding interface entry in the MPLS-VPN-layer in the
    // ifTable. Note that this table does not necessarily correspond one-to-one
    // with all entries in the Interface MIB having an ifType of MPLS-layer;
    // rather, only those which are enabled for MPLS/BGP VPN functionality. The
    // type is interface{} with range: 1..2147483647.
    Mplsvpninterfaceconfindex interface{}

    // Either the providerEdge(0) (PE) or customerEdge(1) (CE) bit MUST be set.
    // The type is Mplsvpninterfacelabeledgetype.
    Mplsvpninterfacelabeledgetype interface{}

    // Denotes whether this link participates in a carrier-of-carrier's,
    // enterprise, or inter-provider scenario. The type is
    // Mplsvpninterfacevpnclassification.
    Mplsvpninterfacevpnclassification interface{}

    // Denotes the route distribution protocol across the PE-CE link. Note that
    // more than one routing protocol may be enabled at the same time. The type is
    // map[string]bool.
    Mplsvpninterfacevpnroutedistprotocol interface{}

    // The storage type for this entry. The type is StorageType.
    Mplsvpninterfaceconfstoragetype interface{}

    // The row status for this entry. This value is used to create a row in this
    // table, signifying that the specified interface is to be associated with the
    // specified interface. If this operation succeeds, the interface will have
    // been associated, otherwise the agent would not allow the association.  If
    // the agent only allows read-only operations on this table, it will create
    // entries in this table as they are created. The type is RowStatus.
    Mplsvpninterfaceconfrowstatus interface{}
}

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetEntityData() *types.CommonEntityData {
    mplsvpninterfaceconfentry.EntityData.YFilter = mplsvpninterfaceconfentry.YFilter
    mplsvpninterfaceconfentry.EntityData.YangName = "mplsVpnInterfaceConfEntry"
    mplsvpninterfaceconfentry.EntityData.BundleName = "cisco_ios_xe"
    mplsvpninterfaceconfentry.EntityData.ParentYangName = "mplsVpnInterfaceConfTable"
    mplsvpninterfaceconfentry.EntityData.SegmentPath = "mplsVpnInterfaceConfEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpninterfaceconfentry.Mplsvpnvrfname) + "']" + "[mplsVpnInterfaceConfIndex='" + fmt.Sprintf("%v", mplsvpninterfaceconfentry.Mplsvpninterfaceconfindex) + "']"
    mplsvpninterfaceconfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpninterfaceconfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpninterfaceconfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpninterfaceconfentry.EntityData.Children = make(map[string]types.YChild)
    mplsvpninterfaceconfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsvpninterfaceconfentry.EntityData.Leafs["mplsVpnVrfName"] = types.YLeaf{"Mplsvpnvrfname", mplsvpninterfaceconfentry.Mplsvpnvrfname}
    mplsvpninterfaceconfentry.EntityData.Leafs["mplsVpnInterfaceConfIndex"] = types.YLeaf{"Mplsvpninterfaceconfindex", mplsvpninterfaceconfentry.Mplsvpninterfaceconfindex}
    mplsvpninterfaceconfentry.EntityData.Leafs["mplsVpnInterfaceLabelEdgeType"] = types.YLeaf{"Mplsvpninterfacelabeledgetype", mplsvpninterfaceconfentry.Mplsvpninterfacelabeledgetype}
    mplsvpninterfaceconfentry.EntityData.Leafs["mplsVpnInterfaceVpnClassification"] = types.YLeaf{"Mplsvpninterfacevpnclassification", mplsvpninterfaceconfentry.Mplsvpninterfacevpnclassification}
    mplsvpninterfaceconfentry.EntityData.Leafs["mplsVpnInterfaceVpnRouteDistProtocol"] = types.YLeaf{"Mplsvpninterfacevpnroutedistprotocol", mplsvpninterfaceconfentry.Mplsvpninterfacevpnroutedistprotocol}
    mplsvpninterfaceconfentry.EntityData.Leafs["mplsVpnInterfaceConfStorageType"] = types.YLeaf{"Mplsvpninterfaceconfstoragetype", mplsvpninterfaceconfentry.Mplsvpninterfaceconfstoragetype}
    mplsvpninterfaceconfentry.EntityData.Leafs["mplsVpnInterfaceConfRowStatus"] = types.YLeaf{"Mplsvpninterfaceconfrowstatus", mplsvpninterfaceconfentry.Mplsvpninterfaceconfrowstatus}
    return &(mplsvpninterfaceconfentry.EntityData)
}

// MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacelabeledgetype represents (CE) bit MUST be set.
type MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacelabeledgetype string

const (
    MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacelabeledgetype_providerEdge MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacelabeledgetype = "providerEdge"

    MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacelabeledgetype_customerEdge MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacelabeledgetype = "customerEdge"
)

// MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification represents scenario.
type MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification string

const (
    MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification_carrierOfCarrier MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification = "carrierOfCarrier"

    MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification_enterprise MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification = "enterprise"

    MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification_interProvider MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfacevpnclassification = "interProvider"
)

// MPLSVPNMIB_Mplsvpnvrftable
// This table specifies per-interface MPLS/BGP VPN
// VRF Table capability and associated information.
// Entries in this table define VRF routing instances
// associated with MPLS/VPN interfaces. Note that
// multiple interfaces can belong to the same VRF
// instance. The collection of all VRF instances
// comprises an actual VPN.
type MPLSVPNMIB_Mplsvpnvrftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS/BGP VPN. The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry.
    Mplsvpnvrfentry []MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry
}

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetEntityData() *types.CommonEntityData {
    mplsvpnvrftable.EntityData.YFilter = mplsvpnvrftable.YFilter
    mplsvpnvrftable.EntityData.YangName = "mplsVpnVrfTable"
    mplsvpnvrftable.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrftable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsvpnvrftable.EntityData.SegmentPath = "mplsVpnVrfTable"
    mplsvpnvrftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrftable.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrftable.EntityData.Children["mplsVpnVrfEntry"] = types.YChild{"Mplsvpnvrfentry", nil}
    for i := range mplsvpnvrftable.Mplsvpnvrfentry {
        mplsvpnvrftable.EntityData.Children[types.GetSegmentPath(&mplsvpnvrftable.Mplsvpnvrfentry[i])] = types.YChild{"Mplsvpnvrfentry", &mplsvpnvrftable.Mplsvpnvrfentry[i]}
    }
    mplsvpnvrftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsvpnvrftable.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS/BGP VPN. The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The human-readable name of this VPN. This MAY be
    // equivalent to the RFC2685 VPN-ID. The type is string with length: 0..31.
    Mplsvpnvrfname interface{}

    // The human-readable description of this VRF. The type is string.
    Mplsvpnvrfdescription interface{}

    // The route distinguisher for this VRF. The type is string with length:
    // 0..256.
    Mplsvpnvrfroutedistinguisher interface{}

    // The time at which this VRF entry was created. The type is interface{} with
    // range: 0..4294967295.
    Mplsvpnvrfcreationtime interface{}

    // Denotes whether a VRF is operational or not. A VRF is  up(1) when at least
    // one interface associated with the VRF, which ifOperStatus is up(1). A VRF
    // is down(2) when:  a. There does not exist at least one interface whose   
    // ifOperStatus is up(1).  b. There are no interfaces associated with the VRF.
    // The type is Mplsvpnvrfoperstatus.
    Mplsvpnvrfoperstatus interface{}

    // Total number of interfaces connected to this VRF with    ifOperStatus =
    // up(1).   This counter should be incremented when:  a. When the ifOperStatus
    // of one of the connected interfaces     changes from down(2) to up(1).  b.
    // When an interface with ifOperStatus = up(1) is connected    to this VRF. 
    // This counter should be decremented when:  a. When the ifOperStatus of one
    // of the connected interfaces     changes from up(1) to down(2).  b. When one
    // of the connected interfaces with     ifOperStatus = up(1) gets disconnected
    // from this VRF. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfactiveinterfaces interface{}

    // Total number of interfaces connected to this VRF  (independent of
    // ifOperStatus type). The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfassociatedinterfaces interface{}

    // Denotes mid-level water marker for the number of routes which  this VRF may
    // hold. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfconfmidroutethreshold interface{}

    // Denotes high-level water marker for the number of routes which  this VRF
    // may hold. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfconfhighroutethreshold interface{}

    // Denotes maximum number of routes which this VRF is configured to hold. This
    // value MUST be less than or equal to mplsVrfMaxPossibleRoutes unless it is
    // set to 0. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfconfmaxroutes interface{}

    // The value of sysUpTime at the time of the last change of this table entry,
    // which includes changes of VRF parameters defined in this table or addition
    // or deletion of interfaces associated with this VRF. The type is interface{}
    // with range: 0..4294967295.
    Mplsvpnvrfconflastchanged interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // The type is RowStatus.
    Mplsvpnvrfconfrowstatus interface{}

    // The storage type for this entry. The type is StorageType.
    Mplsvpnvrfconfstoragetype interface{}

    // Indicates the number of illegally received labels on this VPN/VRF. The type
    // is interface{} with range: 0..4294967295.
    Mplsvpnvrfsecillegallabelviolations interface{}

    // The number of illegally received labels above which this  notification is
    // issued. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfsecillegallabelrcvthresh interface{}

    // Indicates the number of routes added to this VPN/VRF over the coarse of its
    // lifetime. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfperfroutesadded interface{}

    // Indicates the number of routes removed from this VPN/VRF. The type is
    // interface{} with range: 0..4294967295.
    Mplsvpnvrfperfroutesdeleted interface{}

    // Indicates the number of routes currently used by this VRF. The type is
    // interface{} with range: 0..4294967295.
    Mplsvpnvrfperfcurrnumroutes interface{}
}

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfentry.EntityData.YFilter = mplsvpnvrfentry.YFilter
    mplsvpnvrfentry.EntityData.YangName = "mplsVpnVrfEntry"
    mplsvpnvrfentry.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfentry.EntityData.ParentYangName = "mplsVpnVrfTable"
    mplsvpnvrfentry.EntityData.SegmentPath = "mplsVpnVrfEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfentry.Mplsvpnvrfname) + "']"
    mplsvpnvrfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfentry.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfName"] = types.YLeaf{"Mplsvpnvrfname", mplsvpnvrfentry.Mplsvpnvrfname}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfDescription"] = types.YLeaf{"Mplsvpnvrfdescription", mplsvpnvrfentry.Mplsvpnvrfdescription}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfRouteDistinguisher"] = types.YLeaf{"Mplsvpnvrfroutedistinguisher", mplsvpnvrfentry.Mplsvpnvrfroutedistinguisher}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfCreationTime"] = types.YLeaf{"Mplsvpnvrfcreationtime", mplsvpnvrfentry.Mplsvpnvrfcreationtime}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfOperStatus"] = types.YLeaf{"Mplsvpnvrfoperstatus", mplsvpnvrfentry.Mplsvpnvrfoperstatus}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfActiveInterfaces"] = types.YLeaf{"Mplsvpnvrfactiveinterfaces", mplsvpnvrfentry.Mplsvpnvrfactiveinterfaces}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfAssociatedInterfaces"] = types.YLeaf{"Mplsvpnvrfassociatedinterfaces", mplsvpnvrfentry.Mplsvpnvrfassociatedinterfaces}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfConfMidRouteThreshold"] = types.YLeaf{"Mplsvpnvrfconfmidroutethreshold", mplsvpnvrfentry.Mplsvpnvrfconfmidroutethreshold}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfConfHighRouteThreshold"] = types.YLeaf{"Mplsvpnvrfconfhighroutethreshold", mplsvpnvrfentry.Mplsvpnvrfconfhighroutethreshold}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfConfMaxRoutes"] = types.YLeaf{"Mplsvpnvrfconfmaxroutes", mplsvpnvrfentry.Mplsvpnvrfconfmaxroutes}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfConfLastChanged"] = types.YLeaf{"Mplsvpnvrfconflastchanged", mplsvpnvrfentry.Mplsvpnvrfconflastchanged}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfConfRowStatus"] = types.YLeaf{"Mplsvpnvrfconfrowstatus", mplsvpnvrfentry.Mplsvpnvrfconfrowstatus}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfConfStorageType"] = types.YLeaf{"Mplsvpnvrfconfstoragetype", mplsvpnvrfentry.Mplsvpnvrfconfstoragetype}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfSecIllegalLabelViolations"] = types.YLeaf{"Mplsvpnvrfsecillegallabelviolations", mplsvpnvrfentry.Mplsvpnvrfsecillegallabelviolations}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfSecIllegalLabelRcvThresh"] = types.YLeaf{"Mplsvpnvrfsecillegallabelrcvthresh", mplsvpnvrfentry.Mplsvpnvrfsecillegallabelrcvthresh}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfPerfRoutesAdded"] = types.YLeaf{"Mplsvpnvrfperfroutesadded", mplsvpnvrfentry.Mplsvpnvrfperfroutesadded}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfPerfRoutesDeleted"] = types.YLeaf{"Mplsvpnvrfperfroutesdeleted", mplsvpnvrfentry.Mplsvpnvrfperfroutesdeleted}
    mplsvpnvrfentry.EntityData.Leafs["mplsVpnVrfPerfCurrNumRoutes"] = types.YLeaf{"Mplsvpnvrfperfcurrnumroutes", mplsvpnvrfentry.Mplsvpnvrfperfcurrnumroutes}
    return &(mplsvpnvrfentry.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfoperstatus represents b. There are no interfaces associated with the VRF.
type MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfoperstatus string

const (
    MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfoperstatus_up MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfoperstatus = "up"

    MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfoperstatus_down MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfoperstatus = "down"
)

// MPLSVPNMIB_Mplsvpnvrfroutetargettable
// This table specifies per-VRF route target association.
// Each entry identifies a connectivity policy supported
// as part of a VPN.
type MPLSVPNMIB_Mplsvpnvrfroutetargettable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for each route target
    // configured for a VRF supporting a MPLS/BGP VPN instance. The indexing
    // provides an ordering per-VRF instance. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry.
    Mplsvpnvrfroutetargetentry []MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry
}

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfroutetargettable.EntityData.YFilter = mplsvpnvrfroutetargettable.YFilter
    mplsvpnvrfroutetargettable.EntityData.YangName = "mplsVpnVrfRouteTargetTable"
    mplsvpnvrfroutetargettable.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfroutetargettable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsvpnvrfroutetargettable.EntityData.SegmentPath = "mplsVpnVrfRouteTargetTable"
    mplsvpnvrfroutetargettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfroutetargettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfroutetargettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfroutetargettable.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfroutetargettable.EntityData.Children["mplsVpnVrfRouteTargetEntry"] = types.YChild{"Mplsvpnvrfroutetargetentry", nil}
    for i := range mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry {
        mplsvpnvrfroutetargettable.EntityData.Children[types.GetSegmentPath(&mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry[i])] = types.YChild{"Mplsvpnvrfroutetargetentry", &mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry[i]}
    }
    mplsvpnvrfroutetargettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsvpnvrfroutetargettable.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry
//  An entry in this table is created by an LSR for
// each route target configured for a VRF supporting
// a MPLS/BGP VPN instance. The indexing provides an
// ordering per-VRF instance.
type MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfname
    Mplsvpnvrfname interface{}

    // This attribute is a key. Auxiliary index for route-targets configured for a
    // particular VRF. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfroutetargetindex interface{}

    // This attribute is a key. The route target export distribution type. The
    // type is Mplsvpnvrfroutetargettype.
    Mplsvpnvrfroutetargettype interface{}

    // The route target distribution policy. The type is string with length:
    // 0..256.
    Mplsvpnvrfroutetarget interface{}

    // Description of the route target. The type is string.
    Mplsvpnvrfroutetargetdescr interface{}

    // Row status for this entry. The type is RowStatus.
    Mplsvpnvrfroutetargetrowstatus interface{}
}

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfroutetargetentry.EntityData.YFilter = mplsvpnvrfroutetargetentry.YFilter
    mplsvpnvrfroutetargetentry.EntityData.YangName = "mplsVpnVrfRouteTargetEntry"
    mplsvpnvrfroutetargetentry.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfroutetargetentry.EntityData.ParentYangName = "mplsVpnVrfRouteTargetTable"
    mplsvpnvrfroutetargetentry.EntityData.SegmentPath = "mplsVpnVrfRouteTargetEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfroutetargetentry.Mplsvpnvrfname) + "']" + "[mplsVpnVrfRouteTargetIndex='" + fmt.Sprintf("%v", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetindex) + "']" + "[mplsVpnVrfRouteTargetType='" + fmt.Sprintf("%v", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargettype) + "']"
    mplsvpnvrfroutetargetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfroutetargetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfroutetargetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfroutetargetentry.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfroutetargetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsvpnvrfroutetargetentry.EntityData.Leafs["mplsVpnVrfName"] = types.YLeaf{"Mplsvpnvrfname", mplsvpnvrfroutetargetentry.Mplsvpnvrfname}
    mplsvpnvrfroutetargetentry.EntityData.Leafs["mplsVpnVrfRouteTargetIndex"] = types.YLeaf{"Mplsvpnvrfroutetargetindex", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetindex}
    mplsvpnvrfroutetargetentry.EntityData.Leafs["mplsVpnVrfRouteTargetType"] = types.YLeaf{"Mplsvpnvrfroutetargettype", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargettype}
    mplsvpnvrfroutetargetentry.EntityData.Leafs["mplsVpnVrfRouteTarget"] = types.YLeaf{"Mplsvpnvrfroutetarget", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetarget}
    mplsvpnvrfroutetargetentry.EntityData.Leafs["mplsVpnVrfRouteTargetDescr"] = types.YLeaf{"Mplsvpnvrfroutetargetdescr", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetdescr}
    mplsvpnvrfroutetargetentry.EntityData.Leafs["mplsVpnVrfRouteTargetRowStatus"] = types.YLeaf{"Mplsvpnvrfroutetargetrowstatus", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetrowstatus}
    return &(mplsvpnvrfroutetargetentry.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype represents The route target export distribution type.
type MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype string

const (
    MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype_import_ MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype = "import"

    MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype_export MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype = "export"

    MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype_both MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry_Mplsvpnvrfroutetargettype = "both"
)

// MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable
// Each entry in this table specifies a per-interface 
// MPLS/EBGP neighbor.
type MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS/BGP VPN. The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry.
    Mplsvpnvrfbgpnbraddrentry []MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry
}

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfbgpnbraddrtable.EntityData.YFilter = mplsvpnvrfbgpnbraddrtable.YFilter
    mplsvpnvrfbgpnbraddrtable.EntityData.YangName = "mplsVpnVrfBgpNbrAddrTable"
    mplsvpnvrfbgpnbraddrtable.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfbgpnbraddrtable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsvpnvrfbgpnbraddrtable.EntityData.SegmentPath = "mplsVpnVrfBgpNbrAddrTable"
    mplsvpnvrfbgpnbraddrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfbgpnbraddrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfbgpnbraddrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfbgpnbraddrtable.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfbgpnbraddrtable.EntityData.Children["mplsVpnVrfBgpNbrAddrEntry"] = types.YChild{"Mplsvpnvrfbgpnbraddrentry", nil}
    for i := range mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry {
        mplsvpnvrfbgpnbraddrtable.EntityData.Children[types.GetSegmentPath(&mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry[i])] = types.YChild{"Mplsvpnvrfbgpnbraddrentry", &mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry[i]}
    }
    mplsvpnvrfbgpnbraddrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsvpnvrfbgpnbraddrtable.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS/BGP VPN. The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfname
    Mplsvpnvrfname interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // mpls_vpn_mib.MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry_Mplsvpninterfaceconfindex
    Mplsvpninterfaceconfindex interface{}

    // This attribute is a key. This is a unique tertiary index for an entry in
    // the MplsVpnVrfBgpNbrAddrEntry Table. The type is interface{} with range:
    // 0..4294967295.
    Mplsvpnvrfbgpnbrindex interface{}

    // Denotes the role played by this EBGP neighbor with respect to this VRF. The
    // type is Mplsvpnvrfbgpnbrrole.
    Mplsvpnvrfbgpnbrrole interface{}

    // Denotes the address family of the PE address. The type is InetAddressType.
    Mplsvpnvrfbgpnbrtype interface{}

    // Denotes the EBGP neighbor address. The type is string with length: 0..255.
    Mplsvpnvrfbgpnbraddr interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // The type is RowStatus.
    Mplsvpnvrfbgpnbrrowstatus interface{}

    // The storage type for this entry. The type is StorageType.
    Mplsvpnvrfbgpnbrstoragetype interface{}
}

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfbgpnbraddrentry.EntityData.YFilter = mplsvpnvrfbgpnbraddrentry.YFilter
    mplsvpnvrfbgpnbraddrentry.EntityData.YangName = "mplsVpnVrfBgpNbrAddrEntry"
    mplsvpnvrfbgpnbraddrentry.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfbgpnbraddrentry.EntityData.ParentYangName = "mplsVpnVrfBgpNbrAddrTable"
    mplsvpnvrfbgpnbraddrentry.EntityData.SegmentPath = "mplsVpnVrfBgpNbrAddrEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfname) + "']" + "[mplsVpnInterfaceConfIndex='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbraddrentry.Mplsvpninterfaceconfindex) + "']" + "[mplsVpnVrfBgpNbrIndex='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrindex) + "']"
    mplsvpnvrfbgpnbraddrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfbgpnbraddrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfbgpnbraddrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfbgpnbraddrentry.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnVrfName"] = types.YLeaf{"Mplsvpnvrfname", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfname}
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnInterfaceConfIndex"] = types.YLeaf{"Mplsvpninterfaceconfindex", mplsvpnvrfbgpnbraddrentry.Mplsvpninterfaceconfindex}
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnVrfBgpNbrIndex"] = types.YLeaf{"Mplsvpnvrfbgpnbrindex", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrindex}
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnVrfBgpNbrRole"] = types.YLeaf{"Mplsvpnvrfbgpnbrrole", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrrole}
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnVrfBgpNbrType"] = types.YLeaf{"Mplsvpnvrfbgpnbrtype", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrtype}
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnVrfBgpNbrAddr"] = types.YLeaf{"Mplsvpnvrfbgpnbraddr", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbraddr}
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnVrfBgpNbrRowStatus"] = types.YLeaf{"Mplsvpnvrfbgpnbrrowstatus", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrrowstatus}
    mplsvpnvrfbgpnbraddrentry.EntityData.Leafs["mplsVpnVrfBgpNbrStorageType"] = types.YLeaf{"Mplsvpnvrfbgpnbrstoragetype", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrstoragetype}
    return &(mplsvpnvrfbgpnbraddrentry.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry_Mplsvpnvrfbgpnbrrole represents with respect to this VRF.
type MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry_Mplsvpnvrfbgpnbrrole string

const (
    MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry_Mplsvpnvrfbgpnbrrole_ce MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry_Mplsvpnvrfbgpnbrrole = "ce"

    MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry_Mplsvpnvrfbgpnbrrole_pe MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry_Mplsvpnvrfbgpnbrrole = "pe"
)

// MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable
// This table specifies per-VRF vpnv4 multi-protocol
// prefixes supported by BGP.
type MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every BGP prefix associated
    // with a VRF supporting a  MPLS/BGP VPN. The indexing provides an ordering of
    // BGP prefixes per VRF. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry.
    Mplsvpnvrfbgpnbrprefixentry []MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry
}

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfbgpnbrprefixtable.EntityData.YFilter = mplsvpnvrfbgpnbrprefixtable.YFilter
    mplsvpnvrfbgpnbrprefixtable.EntityData.YangName = "mplsVpnVrfBgpNbrPrefixTable"
    mplsvpnvrfbgpnbrprefixtable.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfbgpnbrprefixtable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsvpnvrfbgpnbrprefixtable.EntityData.SegmentPath = "mplsVpnVrfBgpNbrPrefixTable"
    mplsvpnvrfbgpnbrprefixtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfbgpnbrprefixtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfbgpnbrprefixtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfbgpnbrprefixtable.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfbgpnbrprefixtable.EntityData.Children["mplsVpnVrfBgpNbrPrefixEntry"] = types.YChild{"Mplsvpnvrfbgpnbrprefixentry", nil}
    for i := range mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry {
        mplsvpnvrfbgpnbrprefixtable.EntityData.Children[types.GetSegmentPath(&mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry[i])] = types.YChild{"Mplsvpnvrfbgpnbrprefixentry", &mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry[i]}
    }
    mplsvpnvrfbgpnbrprefixtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsvpnvrfbgpnbrprefixtable.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry
// An entry in this table is created by an LSR for
// every BGP prefix associated with a VRF supporting a 
// MPLS/BGP VPN. The indexing provides an ordering of 
// BGP prefixes per VRF.
type MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfname
    Mplsvpnvrfname interface{}

    // This attribute is a key. An IP address prefix in the Network Layer
    // Reachability Information field.  This object is an IP address containing
    // the prefix with length specified by mplsVpnVrfBgpPathAttrIpAddrPrefixLen.
    // Any bits beyond the length specified by
    // mplsVpnVrfBgpPathAttrIpAddrPrefixLen are zeroed. The type is string with
    // length: 0..255.
    Mplsvpnvrfbgppathattripaddrprefix interface{}

    // This attribute is a key. Length in bits of the IP address prefix in the
    // Network Layer Reachability Information field. The type is interface{} with
    // range: 0..32.
    Mplsvpnvrfbgppathattripaddrprefixlen interface{}

    // This attribute is a key. The IP address of the peer where the path
    // information was learned. The type is string with length: 0..255.
    Mplsvpnvrfbgppathattrpeer interface{}

    // The ultimate origin of the path information. The type is
    // Mplsvpnvrfbgppathattrorigin.
    Mplsvpnvrfbgppathattrorigin interface{}

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
    Mplsvpnvrfbgppathattraspathsegment interface{}

    // The address of the border router that should be used for the destination
    // network. The type is string with length: 0..255.
    Mplsvpnvrfbgppathattrnexthop interface{}

    // This metric is used to discriminate between multiple exit points to an
    // adjacent autonomous system.  A value of -1 indicates the absence of this
    // attribute. The type is interface{} with range: -1..2147483647.
    Mplsvpnvrfbgppathattrmultiexitdisc interface{}

    // The originating BGP4 speaker's degree of preference for an advertised
    // route.  A value of -1 indicates the absence of this attribute. The type is
    // interface{} with range: -1..2147483647.
    Mplsvpnvrfbgppathattrlocalpref interface{}

    // Whether or not the local system has selected a less specific route without
    // selecting a more specific route. The type is
    // Mplsvpnvrfbgppathattratomicaggregate.
    Mplsvpnvrfbgppathattratomicaggregate interface{}

    // The AS number of the last BGP4 speaker that performed route aggregation.  A
    // value of zero (0) indicates the absence of this attribute. The type is
    // interface{} with range: 0..65535.
    Mplsvpnvrfbgppathattraggregatoras interface{}

    // The IP address of the last BGP4 speaker that performed route aggregation. 
    // A value of 0.0.0.0 indicates the absence of this attribute. The type is
    // string with length: 0..255.
    Mplsvpnvrfbgppathattraggregatoraddr interface{}

    // The degree of preference calculated by the receiving BGP4 speaker for an
    // advertised route.  A value of -1 indicates the absence of this attribute.
    // The type is interface{} with range: -1..2147483647.
    Mplsvpnvrfbgppathattrcalclocalpref interface{}

    // An indication of whether or not this route was chosen as the best BGP4
    // route. The type is Mplsvpnvrfbgppathattrbest.
    Mplsvpnvrfbgppathattrbest interface{}

    // One or more path attributes not understood by this BGP4 speaker.  Size zero
    // (0) indicates the absence of such attribute(s).  Octets beyond the maximum
    // size, if any, are not recorded by this object. The type is string with
    // length: 0..255.
    Mplsvpnvrfbgppathattrunknown interface{}
}

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfbgpnbrprefixentry.EntityData.YFilter = mplsvpnvrfbgpnbrprefixentry.YFilter
    mplsvpnvrfbgpnbrprefixentry.EntityData.YangName = "mplsVpnVrfBgpNbrPrefixEntry"
    mplsvpnvrfbgpnbrprefixentry.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfbgpnbrprefixentry.EntityData.ParentYangName = "mplsVpnVrfBgpNbrPrefixTable"
    mplsvpnvrfbgpnbrprefixentry.EntityData.SegmentPath = "mplsVpnVrfBgpNbrPrefixEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfname) + "']" + "[mplsVpnVrfBgpPathAttrIpAddrPrefix='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefix) + "']" + "[mplsVpnVrfBgpPathAttrIpAddrPrefixLen='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefixlen) + "']" + "[mplsVpnVrfBgpPathAttrPeer='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrpeer) + "']"
    mplsvpnvrfbgpnbrprefixentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfbgpnbrprefixentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfbgpnbrprefixentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfbgpnbrprefixentry.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfName"] = types.YLeaf{"Mplsvpnvrfname", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfname}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrIpAddrPrefix"] = types.YLeaf{"Mplsvpnvrfbgppathattripaddrprefix", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefix}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrIpAddrPrefixLen"] = types.YLeaf{"Mplsvpnvrfbgppathattripaddrprefixlen", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefixlen}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrPeer"] = types.YLeaf{"Mplsvpnvrfbgppathattrpeer", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrpeer}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrOrigin"] = types.YLeaf{"Mplsvpnvrfbgppathattrorigin", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrorigin}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrASPathSegment"] = types.YLeaf{"Mplsvpnvrfbgppathattraspathsegment", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattraspathsegment}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrNextHop"] = types.YLeaf{"Mplsvpnvrfbgppathattrnexthop", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrnexthop}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrMultiExitDisc"] = types.YLeaf{"Mplsvpnvrfbgppathattrmultiexitdisc", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrmultiexitdisc}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrLocalPref"] = types.YLeaf{"Mplsvpnvrfbgppathattrlocalpref", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrlocalpref}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrAtomicAggregate"] = types.YLeaf{"Mplsvpnvrfbgppathattratomicaggregate", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattratomicaggregate}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrAggregatorAS"] = types.YLeaf{"Mplsvpnvrfbgppathattraggregatoras", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattraggregatoras}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrAggregatorAddr"] = types.YLeaf{"Mplsvpnvrfbgppathattraggregatoraddr", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattraggregatoraddr}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrCalcLocalPref"] = types.YLeaf{"Mplsvpnvrfbgppathattrcalclocalpref", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrcalclocalpref}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrBest"] = types.YLeaf{"Mplsvpnvrfbgppathattrbest", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrbest}
    mplsvpnvrfbgpnbrprefixentry.EntityData.Leafs["mplsVpnVrfBgpPathAttrUnknown"] = types.YLeaf{"Mplsvpnvrfbgppathattrunknown", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrunknown}
    return &(mplsvpnvrfbgpnbrprefixentry.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattratomicaggregate represents selecting a more specific route.
type MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattratomicaggregate string

const (
    MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattratomicaggregate_lessSpecificRrouteNotSelected MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattratomicaggregate = "lessSpecificRrouteNotSelected"

    MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattratomicaggregate_lessSpecificRouteSelected MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattratomicaggregate = "lessSpecificRouteSelected"
)

// MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrbest represents was chosen as the best BGP4 route.
type MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrbest string

const (
    MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrbest_false MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrbest = "false"

    MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrbest_true MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrbest = "true"
)

// MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin represents information.
type MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin string

const (
    MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin_igp MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin = "igp"

    MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin_egp MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin = "egp"

    MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin_incomplete MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry_Mplsvpnvrfbgppathattrorigin = "incomplete"
)

// MPLSVPNMIB_Mplsvpnvrfroutetable
// This table specifies per-interface MPLS/BGP VPN VRF Table
// routing information. Entries in this table define VRF routing
// entries associated with the specified MPLS/VPN interfaces. Note
// that this table contains both BGP and IGP routes, as both may
// appear in the same VRF.
type MPLSVPNMIB_Mplsvpnvrfroutetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every route present
    // configured (either dynamically or statically) within the context of a
    // specific VRF capable of supporting MPLS/BGP VPN. The indexing provides an
    // ordering of VRFs per-VPN interface. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry.
    Mplsvpnvrfrouteentry []MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry
}

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfroutetable.EntityData.YFilter = mplsvpnvrfroutetable.YFilter
    mplsvpnvrfroutetable.EntityData.YangName = "mplsVpnVrfRouteTable"
    mplsvpnvrfroutetable.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfroutetable.EntityData.ParentYangName = "MPLS-VPN-MIB"
    mplsvpnvrfroutetable.EntityData.SegmentPath = "mplsVpnVrfRouteTable"
    mplsvpnvrfroutetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfroutetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfroutetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfroutetable.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfroutetable.EntityData.Children["mplsVpnVrfRouteEntry"] = types.YChild{"Mplsvpnvrfrouteentry", nil}
    for i := range mplsvpnvrfroutetable.Mplsvpnvrfrouteentry {
        mplsvpnvrfroutetable.EntityData.Children[types.GetSegmentPath(&mplsvpnvrfroutetable.Mplsvpnvrfrouteentry[i])] = types.YChild{"Mplsvpnvrfrouteentry", &mplsvpnvrfroutetable.Mplsvpnvrfrouteentry[i]}
    }
    mplsvpnvrfroutetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsvpnvrfroutetable.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry
// An entry in this table is created by an LSR for every route
// present configured (either dynamically or statically) within
// the context of a specific VRF capable of supporting MPLS/BGP
// VPN. The indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..31. Refers to
    // mpls_vpn_mib.MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry_Mplsvpnvrfname
    Mplsvpnvrfname interface{}

    // This attribute is a key. The destination IP address of this route. This
    // object may not take a Multicast (Class D) address value.  Any assignment
    // (implicit or otherwise) of an instance of this object to a value x must be
    // rejected if the bit-wise logical-AND of x with the value of the
    // corresponding instance of the mplsVpnVrfRouteMask object is not equal to x.
    // The type is string with length: 0..255.
    Mplsvpnvrfroutedest interface{}

    // This attribute is a key. Indicate the mask to be logical-ANDed with the
    // destination  address  before  being compared to the value  in  the 
    // mplsVpnVrfRouteDest field. For those  systems  that  do  not support
    // arbitrary subnet masks, an agent constructs the value of the
    // mplsVpnVrfRouteMask by reference   to the IP Address Class.  Any assignment
    // (implicit or otherwise) of an instance of this object to a value x must be
    // rejected if the bit-wise logical-AND of x with the value of the
    // corresponding instance of the mplsVpnVrfRouteDest object is not equal to
    // mplsVpnVrfRouteDest. The type is string with length: 0..255.
    Mplsvpnvrfroutemask interface{}

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
    Mplsvpnvrfroutetos interface{}

    // This attribute is a key. On remote routes, the address of the next system
    // en route; Otherwise, 0.0.0.0. . The type is string with length: 0..255.
    Mplsvpnvrfroutenexthop interface{}

    // The address type of the mplsVpnVrfRouteDest entry. The type is
    // InetAddressType.
    Mplsvpnvrfroutedestaddrtype interface{}

    // The address type of mplsVpnVrfRouteMask. The type is InetAddressType.
    Mplsvpnvrfroutemaskaddrtype interface{}

    // The address type of the mplsVpnVrfRouteNextHopAddrType object. The type is
    // InetAddressType.
    Mplsvpnvrfroutenexthopaddrtype interface{}

    // The ifIndex value that identifies the local interface  through  which  the
    // next hop of this route should be reached. The type is interface{} with
    // range: 0..2147483647.
    Mplsvpnvrfrouteifindex interface{}

    // The type of route.  Note that local(3)  refers to a route for which the
    // next hop is the final destination; remote(4) refers to a route for that the
    // next  hop is not the final destination. Routes which do not result in
    // traffic forwarding or rejection should not be displayed even if the
    // implementation keeps them stored internally.  reject (2) refers to a route
    // which, if matched, discards the message as unreachable. This is used in
    // some protocols as a means of correctly aggregating routes. The type is
    // Mplsvpnvrfroutetype.
    Mplsvpnvrfroutetype interface{}

    // The routing mechanism via which this route was learned.  Inclusion of
    // values for gateway rout- ing protocols is not  intended  to  imply  that
    // hosts should support those protocols. The type is Mplsvpnvrfrouteproto.
    Mplsvpnvrfrouteproto interface{}

    // The number of seconds since this route was last updated or otherwise
    // determined to be correct. Note that no semantics of `too old' can be
    // implied except through knowledge of the routing protocol by which the route
    // was learned. The type is interface{} with range: 0..4294967295.
    Mplsvpnvrfrouteage interface{}

    // A reference to MIB definitions specific to the particular routing protocol
    // which is responsi-   ble for this route, as determined by the  value
    // specified  in the route's mplsVpnVrfRouteProto value. If this information
    // is not present, its value SHOULD be set to the OBJECT IDENTIFIER { 0 0 },
    // which is a syntactically valid object identif-ier, and any implementation
    // conforming to ASN.1 and the Basic Encoding Rules must be able to generate
    // and recognize this value. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mplsvpnvrfrouteinfo interface{}

    // The Autonomous System Number of the Next Hop. The semantics of this object
    // are determined by the routing-protocol specified in the route's
    // mplsVpnVrfRouteProto value. When this object is unknown or not relevant its
    // value should be set to zero. The type is interface{} with range:
    // 0..4294967295.
    Mplsvpnvrfroutenexthopas interface{}

    // The primary routing metric for this route. The semantics of this metric are
    // determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Mplsvpnvrfroutemetric1 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Mplsvpnvrfroutemetric2 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Mplsvpnvrfroutemetric3 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Mplsvpnvrfroutemetric4 interface{}

    // An alternate routing metric for this route. The semantics of this metric
    // are determined by the routing-protocol specified in  the  route's
    // mplsVpnVrfRouteProto value. If this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Mplsvpnvrfroutemetric5 interface{}

    // Row status for this table. It is used according to row installation and
    // removal conventions. The type is RowStatus.
    Mplsvpnvrfrouterowstatus interface{}

    // Storage type value. The type is StorageType.
    Mplsvpnvrfroutestoragetype interface{}
}

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetEntityData() *types.CommonEntityData {
    mplsvpnvrfrouteentry.EntityData.YFilter = mplsvpnvrfrouteentry.YFilter
    mplsvpnvrfrouteentry.EntityData.YangName = "mplsVpnVrfRouteEntry"
    mplsvpnvrfrouteentry.EntityData.BundleName = "cisco_ios_xe"
    mplsvpnvrfrouteentry.EntityData.ParentYangName = "mplsVpnVrfRouteTable"
    mplsvpnvrfrouteentry.EntityData.SegmentPath = "mplsVpnVrfRouteEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfname) + "']" + "[mplsVpnVrfRouteDest='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutedest) + "']" + "[mplsVpnVrfRouteMask='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutemask) + "']" + "[mplsVpnVrfRouteTos='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutetos) + "']" + "[mplsVpnVrfRouteNextHop='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthop) + "']"
    mplsvpnvrfrouteentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsvpnvrfrouteentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsvpnvrfrouteentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsvpnvrfrouteentry.EntityData.Children = make(map[string]types.YChild)
    mplsvpnvrfrouteentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfName"] = types.YLeaf{"Mplsvpnvrfname", mplsvpnvrfrouteentry.Mplsvpnvrfname}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteDest"] = types.YLeaf{"Mplsvpnvrfroutedest", mplsvpnvrfrouteentry.Mplsvpnvrfroutedest}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteMask"] = types.YLeaf{"Mplsvpnvrfroutemask", mplsvpnvrfrouteentry.Mplsvpnvrfroutemask}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteTos"] = types.YLeaf{"Mplsvpnvrfroutetos", mplsvpnvrfrouteentry.Mplsvpnvrfroutetos}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteNextHop"] = types.YLeaf{"Mplsvpnvrfroutenexthop", mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthop}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteDestAddrType"] = types.YLeaf{"Mplsvpnvrfroutedestaddrtype", mplsvpnvrfrouteentry.Mplsvpnvrfroutedestaddrtype}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteMaskAddrType"] = types.YLeaf{"Mplsvpnvrfroutemaskaddrtype", mplsvpnvrfrouteentry.Mplsvpnvrfroutemaskaddrtype}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteNextHopAddrType"] = types.YLeaf{"Mplsvpnvrfroutenexthopaddrtype", mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthopaddrtype}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteIfIndex"] = types.YLeaf{"Mplsvpnvrfrouteifindex", mplsvpnvrfrouteentry.Mplsvpnvrfrouteifindex}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteType"] = types.YLeaf{"Mplsvpnvrfroutetype", mplsvpnvrfrouteentry.Mplsvpnvrfroutetype}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteProto"] = types.YLeaf{"Mplsvpnvrfrouteproto", mplsvpnvrfrouteentry.Mplsvpnvrfrouteproto}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteAge"] = types.YLeaf{"Mplsvpnvrfrouteage", mplsvpnvrfrouteentry.Mplsvpnvrfrouteage}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteInfo"] = types.YLeaf{"Mplsvpnvrfrouteinfo", mplsvpnvrfrouteentry.Mplsvpnvrfrouteinfo}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteNextHopAS"] = types.YLeaf{"Mplsvpnvrfroutenexthopas", mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthopas}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteMetric1"] = types.YLeaf{"Mplsvpnvrfroutemetric1", mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric1}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteMetric2"] = types.YLeaf{"Mplsvpnvrfroutemetric2", mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric2}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteMetric3"] = types.YLeaf{"Mplsvpnvrfroutemetric3", mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric3}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteMetric4"] = types.YLeaf{"Mplsvpnvrfroutemetric4", mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric4}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteMetric5"] = types.YLeaf{"Mplsvpnvrfroutemetric5", mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric5}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteRowStatus"] = types.YLeaf{"Mplsvpnvrfrouterowstatus", mplsvpnvrfrouteentry.Mplsvpnvrfrouterowstatus}
    mplsvpnvrfrouteentry.EntityData.Leafs["mplsVpnVrfRouteStorageType"] = types.YLeaf{"Mplsvpnvrfroutestoragetype", mplsvpnvrfrouteentry.Mplsvpnvrfroutestoragetype}
    return &(mplsvpnvrfrouteentry.EntityData)
}

// MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto represents hosts should support those protocols.
type MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto string

const (
    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_other MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "other"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_local MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "local"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_netmgmt MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "netmgmt"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_icmp MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "icmp"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_egp MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "egp"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_ggp MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "ggp"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_hello MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "hello"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_rip MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "rip"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_isIs MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "isIs"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_esIs MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "esIs"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_ciscoIgrp MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "ciscoIgrp"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_bbnSpfIgp MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "bbnSpfIgp"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_ospf MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "ospf"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_bgp MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "bgp"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_idpr MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "idpr"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto_ciscoEigrp MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfrouteproto = "ciscoEigrp"
)

// MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype represents routes.
type MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype string

const (
    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype_other MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype = "other"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype_reject MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype = "reject"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype_local MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype = "local"

    MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype_remote MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry_Mplsvpnvrfroutetype = "remote"
)

