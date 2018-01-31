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
    parent types.Entity
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

func (mPLSVPNMIB *MPLSVPNMIB) GetFilter() yfilter.YFilter { return mPLSVPNMIB.YFilter }

func (mPLSVPNMIB *MPLSVPNMIB) SetFilter(yf yfilter.YFilter) { mPLSVPNMIB.YFilter = yf }

func (mPLSVPNMIB *MPLSVPNMIB) GetGoName(yname string) string {
    if yname == "mplsVpnScalars" { return "Mplsvpnscalars" }
    if yname == "mplsVpnInterfaceConfTable" { return "Mplsvpninterfaceconftable" }
    if yname == "mplsVpnVrfTable" { return "Mplsvpnvrftable" }
    if yname == "mplsVpnVrfRouteTargetTable" { return "Mplsvpnvrfroutetargettable" }
    if yname == "mplsVpnVrfBgpNbrAddrTable" { return "Mplsvpnvrfbgpnbraddrtable" }
    if yname == "mplsVpnVrfBgpNbrPrefixTable" { return "Mplsvpnvrfbgpnbrprefixtable" }
    if yname == "mplsVpnVrfRouteTable" { return "Mplsvpnvrfroutetable" }
    return ""
}

func (mPLSVPNMIB *MPLSVPNMIB) GetSegmentPath() string {
    return "MPLS-VPN-MIB:MPLS-VPN-MIB"
}

func (mPLSVPNMIB *MPLSVPNMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsVpnScalars" {
        return &mPLSVPNMIB.Mplsvpnscalars
    }
    if childYangName == "mplsVpnInterfaceConfTable" {
        return &mPLSVPNMIB.Mplsvpninterfaceconftable
    }
    if childYangName == "mplsVpnVrfTable" {
        return &mPLSVPNMIB.Mplsvpnvrftable
    }
    if childYangName == "mplsVpnVrfRouteTargetTable" {
        return &mPLSVPNMIB.Mplsvpnvrfroutetargettable
    }
    if childYangName == "mplsVpnVrfBgpNbrAddrTable" {
        return &mPLSVPNMIB.Mplsvpnvrfbgpnbraddrtable
    }
    if childYangName == "mplsVpnVrfBgpNbrPrefixTable" {
        return &mPLSVPNMIB.Mplsvpnvrfbgpnbrprefixtable
    }
    if childYangName == "mplsVpnVrfRouteTable" {
        return &mPLSVPNMIB.Mplsvpnvrfroutetable
    }
    return nil
}

func (mPLSVPNMIB *MPLSVPNMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mplsVpnScalars"] = &mPLSVPNMIB.Mplsvpnscalars
    children["mplsVpnInterfaceConfTable"] = &mPLSVPNMIB.Mplsvpninterfaceconftable
    children["mplsVpnVrfTable"] = &mPLSVPNMIB.Mplsvpnvrftable
    children["mplsVpnVrfRouteTargetTable"] = &mPLSVPNMIB.Mplsvpnvrfroutetargettable
    children["mplsVpnVrfBgpNbrAddrTable"] = &mPLSVPNMIB.Mplsvpnvrfbgpnbraddrtable
    children["mplsVpnVrfBgpNbrPrefixTable"] = &mPLSVPNMIB.Mplsvpnvrfbgpnbrprefixtable
    children["mplsVpnVrfRouteTable"] = &mPLSVPNMIB.Mplsvpnvrfroutetable
    return children
}

func (mPLSVPNMIB *MPLSVPNMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mPLSVPNMIB *MPLSVPNMIB) GetBundleName() string { return "cisco_ios_xe" }

func (mPLSVPNMIB *MPLSVPNMIB) GetYangName() string { return "MPLS-VPN-MIB" }

func (mPLSVPNMIB *MPLSVPNMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mPLSVPNMIB *MPLSVPNMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mPLSVPNMIB *MPLSVPNMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mPLSVPNMIB *MPLSVPNMIB) SetParent(parent types.Entity) { mPLSVPNMIB.parent = parent }

func (mPLSVPNMIB *MPLSVPNMIB) GetParent() types.Entity { return mPLSVPNMIB.parent }

func (mPLSVPNMIB *MPLSVPNMIB) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpnscalars
type MPLSVPNMIB_Mplsvpnscalars struct {
    parent types.Entity
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

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetFilter() yfilter.YFilter { return mplsvpnscalars.YFilter }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) SetFilter(yf yfilter.YFilter) { mplsvpnscalars.YFilter = yf }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetGoName(yname string) string {
    if yname == "mplsVpnConfiguredVrfs" { return "Mplsvpnconfiguredvrfs" }
    if yname == "mplsVpnActiveVrfs" { return "Mplsvpnactivevrfs" }
    if yname == "mplsVpnConnectedInterfaces" { return "Mplsvpnconnectedinterfaces" }
    if yname == "mplsVpnNotificationEnable" { return "Mplsvpnnotificationenable" }
    if yname == "mplsVpnVrfConfMaxPossibleRoutes" { return "Mplsvpnvrfconfmaxpossibleroutes" }
    return ""
}

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetSegmentPath() string {
    return "mplsVpnScalars"
}

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsVpnConfiguredVrfs"] = mplsvpnscalars.Mplsvpnconfiguredvrfs
    leafs["mplsVpnActiveVrfs"] = mplsvpnscalars.Mplsvpnactivevrfs
    leafs["mplsVpnConnectedInterfaces"] = mplsvpnscalars.Mplsvpnconnectedinterfaces
    leafs["mplsVpnNotificationEnable"] = mplsvpnscalars.Mplsvpnnotificationenable
    leafs["mplsVpnVrfConfMaxPossibleRoutes"] = mplsvpnscalars.Mplsvpnvrfconfmaxpossibleroutes
    return leafs
}

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetYangName() string { return "mplsVpnScalars" }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) SetParent(parent types.Entity) { mplsvpnscalars.parent = parent }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetParent() types.Entity { return mplsvpnscalars.parent }

func (mplsvpnscalars *MPLSVPNMIB_Mplsvpnscalars) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpninterfaceconftable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSVPNMIB_Mplsvpninterfaceconftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every interface capable of
    // supporting MPLS/BGP VPN.   Each entry in this table is meant to correspond
    // to an entry in the Interfaces Table. The type is slice of
    // MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry.
    Mplsvpninterfaceconfentry []MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry
}

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetFilter() yfilter.YFilter { return mplsvpninterfaceconftable.YFilter }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) SetFilter(yf yfilter.YFilter) { mplsvpninterfaceconftable.YFilter = yf }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetGoName(yname string) string {
    if yname == "mplsVpnInterfaceConfEntry" { return "Mplsvpninterfaceconfentry" }
    return ""
}

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetSegmentPath() string {
    return "mplsVpnInterfaceConfTable"
}

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsVpnInterfaceConfEntry" {
        for _, c := range mplsvpninterfaceconftable.Mplsvpninterfaceconfentry {
            if mplsvpninterfaceconftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry{}
        mplsvpninterfaceconftable.Mplsvpninterfaceconfentry = append(mplsvpninterfaceconftable.Mplsvpninterfaceconfentry, child)
        return &mplsvpninterfaceconftable.Mplsvpninterfaceconfentry[len(mplsvpninterfaceconftable.Mplsvpninterfaceconfentry)-1]
    }
    return nil
}

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsvpninterfaceconftable.Mplsvpninterfaceconfentry {
        children[mplsvpninterfaceconftable.Mplsvpninterfaceconfentry[i].GetSegmentPath()] = &mplsvpninterfaceconftable.Mplsvpninterfaceconfentry[i]
    }
    return children
}

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetYangName() string { return "mplsVpnInterfaceConfTable" }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) SetParent(parent types.Entity) { mplsvpninterfaceconftable.parent = parent }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetParent() types.Entity { return mplsvpninterfaceconftable.parent }

func (mplsvpninterfaceconftable *MPLSVPNMIB_Mplsvpninterfaceconftable) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry
// An entry in this table is created by an LSR for
// every interface capable of supporting MPLS/BGP VPN.
// 
// 
// Each entry in this table is meant to correspond to
// an entry in the Interfaces Table.
type MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry struct {
    parent types.Entity
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

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetFilter() yfilter.YFilter { return mplsvpninterfaceconfentry.YFilter }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) SetFilter(yf yfilter.YFilter) { mplsvpninterfaceconfentry.YFilter = yf }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetGoName(yname string) string {
    if yname == "mplsVpnVrfName" { return "Mplsvpnvrfname" }
    if yname == "mplsVpnInterfaceConfIndex" { return "Mplsvpninterfaceconfindex" }
    if yname == "mplsVpnInterfaceLabelEdgeType" { return "Mplsvpninterfacelabeledgetype" }
    if yname == "mplsVpnInterfaceVpnClassification" { return "Mplsvpninterfacevpnclassification" }
    if yname == "mplsVpnInterfaceVpnRouteDistProtocol" { return "Mplsvpninterfacevpnroutedistprotocol" }
    if yname == "mplsVpnInterfaceConfStorageType" { return "Mplsvpninterfaceconfstoragetype" }
    if yname == "mplsVpnInterfaceConfRowStatus" { return "Mplsvpninterfaceconfrowstatus" }
    return ""
}

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetSegmentPath() string {
    return "mplsVpnInterfaceConfEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpninterfaceconfentry.Mplsvpnvrfname) + "']" + "[mplsVpnInterfaceConfIndex='" + fmt.Sprintf("%v", mplsvpninterfaceconfentry.Mplsvpninterfaceconfindex) + "']"
}

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsVpnVrfName"] = mplsvpninterfaceconfentry.Mplsvpnvrfname
    leafs["mplsVpnInterfaceConfIndex"] = mplsvpninterfaceconfentry.Mplsvpninterfaceconfindex
    leafs["mplsVpnInterfaceLabelEdgeType"] = mplsvpninterfaceconfentry.Mplsvpninterfacelabeledgetype
    leafs["mplsVpnInterfaceVpnClassification"] = mplsvpninterfaceconfentry.Mplsvpninterfacevpnclassification
    leafs["mplsVpnInterfaceVpnRouteDistProtocol"] = mplsvpninterfaceconfentry.Mplsvpninterfacevpnroutedistprotocol
    leafs["mplsVpnInterfaceConfStorageType"] = mplsvpninterfaceconfentry.Mplsvpninterfaceconfstoragetype
    leafs["mplsVpnInterfaceConfRowStatus"] = mplsvpninterfaceconfentry.Mplsvpninterfaceconfrowstatus
    return leafs
}

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetYangName() string { return "mplsVpnInterfaceConfEntry" }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) SetParent(parent types.Entity) { mplsvpninterfaceconfentry.parent = parent }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetParent() types.Entity { return mplsvpninterfaceconfentry.parent }

func (mplsvpninterfaceconfentry *MPLSVPNMIB_Mplsvpninterfaceconftable_Mplsvpninterfaceconfentry) GetParentYangName() string { return "mplsVpnInterfaceConfTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS/BGP VPN. The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry.
    Mplsvpnvrfentry []MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry
}

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetFilter() yfilter.YFilter { return mplsvpnvrftable.YFilter }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) SetFilter(yf yfilter.YFilter) { mplsvpnvrftable.YFilter = yf }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetGoName(yname string) string {
    if yname == "mplsVpnVrfEntry" { return "Mplsvpnvrfentry" }
    return ""
}

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetSegmentPath() string {
    return "mplsVpnVrfTable"
}

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsVpnVrfEntry" {
        for _, c := range mplsvpnvrftable.Mplsvpnvrfentry {
            if mplsvpnvrftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry{}
        mplsvpnvrftable.Mplsvpnvrfentry = append(mplsvpnvrftable.Mplsvpnvrfentry, child)
        return &mplsvpnvrftable.Mplsvpnvrfentry[len(mplsvpnvrftable.Mplsvpnvrfentry)-1]
    }
    return nil
}

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsvpnvrftable.Mplsvpnvrfentry {
        children[mplsvpnvrftable.Mplsvpnvrfentry[i].GetSegmentPath()] = &mplsvpnvrftable.Mplsvpnvrfentry[i]
    }
    return children
}

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetYangName() string { return "mplsVpnVrfTable" }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) SetParent(parent types.Entity) { mplsvpnvrftable.parent = parent }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetParent() types.Entity { return mplsvpnvrftable.parent }

func (mplsvpnvrftable *MPLSVPNMIB_Mplsvpnvrftable) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS/BGP VPN. The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry struct {
    parent types.Entity
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

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetFilter() yfilter.YFilter { return mplsvpnvrfentry.YFilter }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) SetFilter(yf yfilter.YFilter) { mplsvpnvrfentry.YFilter = yf }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetGoName(yname string) string {
    if yname == "mplsVpnVrfName" { return "Mplsvpnvrfname" }
    if yname == "mplsVpnVrfDescription" { return "Mplsvpnvrfdescription" }
    if yname == "mplsVpnVrfRouteDistinguisher" { return "Mplsvpnvrfroutedistinguisher" }
    if yname == "mplsVpnVrfCreationTime" { return "Mplsvpnvrfcreationtime" }
    if yname == "mplsVpnVrfOperStatus" { return "Mplsvpnvrfoperstatus" }
    if yname == "mplsVpnVrfActiveInterfaces" { return "Mplsvpnvrfactiveinterfaces" }
    if yname == "mplsVpnVrfAssociatedInterfaces" { return "Mplsvpnvrfassociatedinterfaces" }
    if yname == "mplsVpnVrfConfMidRouteThreshold" { return "Mplsvpnvrfconfmidroutethreshold" }
    if yname == "mplsVpnVrfConfHighRouteThreshold" { return "Mplsvpnvrfconfhighroutethreshold" }
    if yname == "mplsVpnVrfConfMaxRoutes" { return "Mplsvpnvrfconfmaxroutes" }
    if yname == "mplsVpnVrfConfLastChanged" { return "Mplsvpnvrfconflastchanged" }
    if yname == "mplsVpnVrfConfRowStatus" { return "Mplsvpnvrfconfrowstatus" }
    if yname == "mplsVpnVrfConfStorageType" { return "Mplsvpnvrfconfstoragetype" }
    if yname == "mplsVpnVrfSecIllegalLabelViolations" { return "Mplsvpnvrfsecillegallabelviolations" }
    if yname == "mplsVpnVrfSecIllegalLabelRcvThresh" { return "Mplsvpnvrfsecillegallabelrcvthresh" }
    if yname == "mplsVpnVrfPerfRoutesAdded" { return "Mplsvpnvrfperfroutesadded" }
    if yname == "mplsVpnVrfPerfRoutesDeleted" { return "Mplsvpnvrfperfroutesdeleted" }
    if yname == "mplsVpnVrfPerfCurrNumRoutes" { return "Mplsvpnvrfperfcurrnumroutes" }
    return ""
}

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetSegmentPath() string {
    return "mplsVpnVrfEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfentry.Mplsvpnvrfname) + "']"
}

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsVpnVrfName"] = mplsvpnvrfentry.Mplsvpnvrfname
    leafs["mplsVpnVrfDescription"] = mplsvpnvrfentry.Mplsvpnvrfdescription
    leafs["mplsVpnVrfRouteDistinguisher"] = mplsvpnvrfentry.Mplsvpnvrfroutedistinguisher
    leafs["mplsVpnVrfCreationTime"] = mplsvpnvrfentry.Mplsvpnvrfcreationtime
    leafs["mplsVpnVrfOperStatus"] = mplsvpnvrfentry.Mplsvpnvrfoperstatus
    leafs["mplsVpnVrfActiveInterfaces"] = mplsvpnvrfentry.Mplsvpnvrfactiveinterfaces
    leafs["mplsVpnVrfAssociatedInterfaces"] = mplsvpnvrfentry.Mplsvpnvrfassociatedinterfaces
    leafs["mplsVpnVrfConfMidRouteThreshold"] = mplsvpnvrfentry.Mplsvpnvrfconfmidroutethreshold
    leafs["mplsVpnVrfConfHighRouteThreshold"] = mplsvpnvrfentry.Mplsvpnvrfconfhighroutethreshold
    leafs["mplsVpnVrfConfMaxRoutes"] = mplsvpnvrfentry.Mplsvpnvrfconfmaxroutes
    leafs["mplsVpnVrfConfLastChanged"] = mplsvpnvrfentry.Mplsvpnvrfconflastchanged
    leafs["mplsVpnVrfConfRowStatus"] = mplsvpnvrfentry.Mplsvpnvrfconfrowstatus
    leafs["mplsVpnVrfConfStorageType"] = mplsvpnvrfentry.Mplsvpnvrfconfstoragetype
    leafs["mplsVpnVrfSecIllegalLabelViolations"] = mplsvpnvrfentry.Mplsvpnvrfsecillegallabelviolations
    leafs["mplsVpnVrfSecIllegalLabelRcvThresh"] = mplsvpnvrfentry.Mplsvpnvrfsecillegallabelrcvthresh
    leafs["mplsVpnVrfPerfRoutesAdded"] = mplsvpnvrfentry.Mplsvpnvrfperfroutesadded
    leafs["mplsVpnVrfPerfRoutesDeleted"] = mplsvpnvrfentry.Mplsvpnvrfperfroutesdeleted
    leafs["mplsVpnVrfPerfCurrNumRoutes"] = mplsvpnvrfentry.Mplsvpnvrfperfcurrnumroutes
    return leafs
}

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetYangName() string { return "mplsVpnVrfEntry" }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) SetParent(parent types.Entity) { mplsvpnvrfentry.parent = parent }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetParent() types.Entity { return mplsvpnvrfentry.parent }

func (mplsvpnvrfentry *MPLSVPNMIB_Mplsvpnvrftable_Mplsvpnvrfentry) GetParentYangName() string { return "mplsVpnVrfTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for each route target
    // configured for a VRF supporting a MPLS/BGP VPN instance. The indexing
    // provides an ordering per-VRF instance. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry.
    Mplsvpnvrfroutetargetentry []MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry
}

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetFilter() yfilter.YFilter { return mplsvpnvrfroutetargettable.YFilter }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) SetFilter(yf yfilter.YFilter) { mplsvpnvrfroutetargettable.YFilter = yf }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetGoName(yname string) string {
    if yname == "mplsVpnVrfRouteTargetEntry" { return "Mplsvpnvrfroutetargetentry" }
    return ""
}

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetSegmentPath() string {
    return "mplsVpnVrfRouteTargetTable"
}

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsVpnVrfRouteTargetEntry" {
        for _, c := range mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry {
            if mplsvpnvrfroutetargettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry{}
        mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry = append(mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry, child)
        return &mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry[len(mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry)-1]
    }
    return nil
}

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry {
        children[mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry[i].GetSegmentPath()] = &mplsvpnvrfroutetargettable.Mplsvpnvrfroutetargetentry[i]
    }
    return children
}

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetYangName() string { return "mplsVpnVrfRouteTargetTable" }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) SetParent(parent types.Entity) { mplsvpnvrfroutetargettable.parent = parent }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetParent() types.Entity { return mplsvpnvrfroutetargettable.parent }

func (mplsvpnvrfroutetargettable *MPLSVPNMIB_Mplsvpnvrfroutetargettable) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry
//  An entry in this table is created by an LSR for
// each route target configured for a VRF supporting
// a MPLS/BGP VPN instance. The indexing provides an
// ordering per-VRF instance.
type MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry struct {
    parent types.Entity
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

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetFilter() yfilter.YFilter { return mplsvpnvrfroutetargetentry.YFilter }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) SetFilter(yf yfilter.YFilter) { mplsvpnvrfroutetargetentry.YFilter = yf }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetGoName(yname string) string {
    if yname == "mplsVpnVrfName" { return "Mplsvpnvrfname" }
    if yname == "mplsVpnVrfRouteTargetIndex" { return "Mplsvpnvrfroutetargetindex" }
    if yname == "mplsVpnVrfRouteTargetType" { return "Mplsvpnvrfroutetargettype" }
    if yname == "mplsVpnVrfRouteTarget" { return "Mplsvpnvrfroutetarget" }
    if yname == "mplsVpnVrfRouteTargetDescr" { return "Mplsvpnvrfroutetargetdescr" }
    if yname == "mplsVpnVrfRouteTargetRowStatus" { return "Mplsvpnvrfroutetargetrowstatus" }
    return ""
}

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetSegmentPath() string {
    return "mplsVpnVrfRouteTargetEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfroutetargetentry.Mplsvpnvrfname) + "']" + "[mplsVpnVrfRouteTargetIndex='" + fmt.Sprintf("%v", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetindex) + "']" + "[mplsVpnVrfRouteTargetType='" + fmt.Sprintf("%v", mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargettype) + "']"
}

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsVpnVrfName"] = mplsvpnvrfroutetargetentry.Mplsvpnvrfname
    leafs["mplsVpnVrfRouteTargetIndex"] = mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetindex
    leafs["mplsVpnVrfRouteTargetType"] = mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargettype
    leafs["mplsVpnVrfRouteTarget"] = mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetarget
    leafs["mplsVpnVrfRouteTargetDescr"] = mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetdescr
    leafs["mplsVpnVrfRouteTargetRowStatus"] = mplsvpnvrfroutetargetentry.Mplsvpnvrfroutetargetrowstatus
    return leafs
}

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetYangName() string { return "mplsVpnVrfRouteTargetEntry" }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) SetParent(parent types.Entity) { mplsvpnvrfroutetargetentry.parent = parent }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetParent() types.Entity { return mplsvpnvrfroutetargetentry.parent }

func (mplsvpnvrfroutetargetentry *MPLSVPNMIB_Mplsvpnvrfroutetargettable_Mplsvpnvrfroutetargetentry) GetParentYangName() string { return "mplsVpnVrfRouteTargetTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS/BGP VPN. The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry.
    Mplsvpnvrfbgpnbraddrentry []MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry
}

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetFilter() yfilter.YFilter { return mplsvpnvrfbgpnbraddrtable.YFilter }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) SetFilter(yf yfilter.YFilter) { mplsvpnvrfbgpnbraddrtable.YFilter = yf }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetGoName(yname string) string {
    if yname == "mplsVpnVrfBgpNbrAddrEntry" { return "Mplsvpnvrfbgpnbraddrentry" }
    return ""
}

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetSegmentPath() string {
    return "mplsVpnVrfBgpNbrAddrTable"
}

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsVpnVrfBgpNbrAddrEntry" {
        for _, c := range mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry {
            if mplsvpnvrfbgpnbraddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry{}
        mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry = append(mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry, child)
        return &mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry[len(mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry)-1]
    }
    return nil
}

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry {
        children[mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry[i].GetSegmentPath()] = &mplsvpnvrfbgpnbraddrtable.Mplsvpnvrfbgpnbraddrentry[i]
    }
    return children
}

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetYangName() string { return "mplsVpnVrfBgpNbrAddrTable" }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) SetParent(parent types.Entity) { mplsvpnvrfbgpnbraddrtable.parent = parent }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetParent() types.Entity { return mplsvpnvrfbgpnbraddrtable.parent }

func (mplsvpnvrfbgpnbraddrtable *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS/BGP VPN. The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry struct {
    parent types.Entity
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

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetFilter() yfilter.YFilter { return mplsvpnvrfbgpnbraddrentry.YFilter }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) SetFilter(yf yfilter.YFilter) { mplsvpnvrfbgpnbraddrentry.YFilter = yf }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetGoName(yname string) string {
    if yname == "mplsVpnVrfName" { return "Mplsvpnvrfname" }
    if yname == "mplsVpnInterfaceConfIndex" { return "Mplsvpninterfaceconfindex" }
    if yname == "mplsVpnVrfBgpNbrIndex" { return "Mplsvpnvrfbgpnbrindex" }
    if yname == "mplsVpnVrfBgpNbrRole" { return "Mplsvpnvrfbgpnbrrole" }
    if yname == "mplsVpnVrfBgpNbrType" { return "Mplsvpnvrfbgpnbrtype" }
    if yname == "mplsVpnVrfBgpNbrAddr" { return "Mplsvpnvrfbgpnbraddr" }
    if yname == "mplsVpnVrfBgpNbrRowStatus" { return "Mplsvpnvrfbgpnbrrowstatus" }
    if yname == "mplsVpnVrfBgpNbrStorageType" { return "Mplsvpnvrfbgpnbrstoragetype" }
    return ""
}

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetSegmentPath() string {
    return "mplsVpnVrfBgpNbrAddrEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfname) + "']" + "[mplsVpnInterfaceConfIndex='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbraddrentry.Mplsvpninterfaceconfindex) + "']" + "[mplsVpnVrfBgpNbrIndex='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrindex) + "']"
}

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsVpnVrfName"] = mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfname
    leafs["mplsVpnInterfaceConfIndex"] = mplsvpnvrfbgpnbraddrentry.Mplsvpninterfaceconfindex
    leafs["mplsVpnVrfBgpNbrIndex"] = mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrindex
    leafs["mplsVpnVrfBgpNbrRole"] = mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrrole
    leafs["mplsVpnVrfBgpNbrType"] = mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrtype
    leafs["mplsVpnVrfBgpNbrAddr"] = mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbraddr
    leafs["mplsVpnVrfBgpNbrRowStatus"] = mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrrowstatus
    leafs["mplsVpnVrfBgpNbrStorageType"] = mplsvpnvrfbgpnbraddrentry.Mplsvpnvrfbgpnbrstoragetype
    return leafs
}

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetYangName() string { return "mplsVpnVrfBgpNbrAddrEntry" }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) SetParent(parent types.Entity) { mplsvpnvrfbgpnbraddrentry.parent = parent }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetParent() types.Entity { return mplsvpnvrfbgpnbraddrentry.parent }

func (mplsvpnvrfbgpnbraddrentry *MPLSVPNMIB_Mplsvpnvrfbgpnbraddrtable_Mplsvpnvrfbgpnbraddrentry) GetParentYangName() string { return "mplsVpnVrfBgpNbrAddrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every BGP prefix associated
    // with a VRF supporting a  MPLS/BGP VPN. The indexing provides an ordering of
    // BGP prefixes per VRF. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry.
    Mplsvpnvrfbgpnbrprefixentry []MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry
}

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetFilter() yfilter.YFilter { return mplsvpnvrfbgpnbrprefixtable.YFilter }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) SetFilter(yf yfilter.YFilter) { mplsvpnvrfbgpnbrprefixtable.YFilter = yf }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetGoName(yname string) string {
    if yname == "mplsVpnVrfBgpNbrPrefixEntry" { return "Mplsvpnvrfbgpnbrprefixentry" }
    return ""
}

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetSegmentPath() string {
    return "mplsVpnVrfBgpNbrPrefixTable"
}

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsVpnVrfBgpNbrPrefixEntry" {
        for _, c := range mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry {
            if mplsvpnvrfbgpnbrprefixtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry{}
        mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry = append(mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry, child)
        return &mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry[len(mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry)-1]
    }
    return nil
}

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry {
        children[mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry[i].GetSegmentPath()] = &mplsvpnvrfbgpnbrprefixtable.Mplsvpnvrfbgpnbrprefixentry[i]
    }
    return children
}

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetYangName() string { return "mplsVpnVrfBgpNbrPrefixTable" }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) SetParent(parent types.Entity) { mplsvpnvrfbgpnbrprefixtable.parent = parent }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetParent() types.Entity { return mplsvpnvrfbgpnbrprefixtable.parent }

func (mplsvpnvrfbgpnbrprefixtable *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry
// An entry in this table is created by an LSR for
// every BGP prefix associated with a VRF supporting a 
// MPLS/BGP VPN. The indexing provides an ordering of 
// BGP prefixes per VRF.
type MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry struct {
    parent types.Entity
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

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetFilter() yfilter.YFilter { return mplsvpnvrfbgpnbrprefixentry.YFilter }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) SetFilter(yf yfilter.YFilter) { mplsvpnvrfbgpnbrprefixentry.YFilter = yf }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetGoName(yname string) string {
    if yname == "mplsVpnVrfName" { return "Mplsvpnvrfname" }
    if yname == "mplsVpnVrfBgpPathAttrIpAddrPrefix" { return "Mplsvpnvrfbgppathattripaddrprefix" }
    if yname == "mplsVpnVrfBgpPathAttrIpAddrPrefixLen" { return "Mplsvpnvrfbgppathattripaddrprefixlen" }
    if yname == "mplsVpnVrfBgpPathAttrPeer" { return "Mplsvpnvrfbgppathattrpeer" }
    if yname == "mplsVpnVrfBgpPathAttrOrigin" { return "Mplsvpnvrfbgppathattrorigin" }
    if yname == "mplsVpnVrfBgpPathAttrASPathSegment" { return "Mplsvpnvrfbgppathattraspathsegment" }
    if yname == "mplsVpnVrfBgpPathAttrNextHop" { return "Mplsvpnvrfbgppathattrnexthop" }
    if yname == "mplsVpnVrfBgpPathAttrMultiExitDisc" { return "Mplsvpnvrfbgppathattrmultiexitdisc" }
    if yname == "mplsVpnVrfBgpPathAttrLocalPref" { return "Mplsvpnvrfbgppathattrlocalpref" }
    if yname == "mplsVpnVrfBgpPathAttrAtomicAggregate" { return "Mplsvpnvrfbgppathattratomicaggregate" }
    if yname == "mplsVpnVrfBgpPathAttrAggregatorAS" { return "Mplsvpnvrfbgppathattraggregatoras" }
    if yname == "mplsVpnVrfBgpPathAttrAggregatorAddr" { return "Mplsvpnvrfbgppathattraggregatoraddr" }
    if yname == "mplsVpnVrfBgpPathAttrCalcLocalPref" { return "Mplsvpnvrfbgppathattrcalclocalpref" }
    if yname == "mplsVpnVrfBgpPathAttrBest" { return "Mplsvpnvrfbgppathattrbest" }
    if yname == "mplsVpnVrfBgpPathAttrUnknown" { return "Mplsvpnvrfbgppathattrunknown" }
    return ""
}

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetSegmentPath() string {
    return "mplsVpnVrfBgpNbrPrefixEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfname) + "']" + "[mplsVpnVrfBgpPathAttrIpAddrPrefix='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefix) + "']" + "[mplsVpnVrfBgpPathAttrIpAddrPrefixLen='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefixlen) + "']" + "[mplsVpnVrfBgpPathAttrPeer='" + fmt.Sprintf("%v", mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrpeer) + "']"
}

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsVpnVrfName"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfname
    leafs["mplsVpnVrfBgpPathAttrIpAddrPrefix"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefix
    leafs["mplsVpnVrfBgpPathAttrIpAddrPrefixLen"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattripaddrprefixlen
    leafs["mplsVpnVrfBgpPathAttrPeer"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrpeer
    leafs["mplsVpnVrfBgpPathAttrOrigin"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrorigin
    leafs["mplsVpnVrfBgpPathAttrASPathSegment"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattraspathsegment
    leafs["mplsVpnVrfBgpPathAttrNextHop"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrnexthop
    leafs["mplsVpnVrfBgpPathAttrMultiExitDisc"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrmultiexitdisc
    leafs["mplsVpnVrfBgpPathAttrLocalPref"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrlocalpref
    leafs["mplsVpnVrfBgpPathAttrAtomicAggregate"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattratomicaggregate
    leafs["mplsVpnVrfBgpPathAttrAggregatorAS"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattraggregatoras
    leafs["mplsVpnVrfBgpPathAttrAggregatorAddr"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattraggregatoraddr
    leafs["mplsVpnVrfBgpPathAttrCalcLocalPref"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrcalclocalpref
    leafs["mplsVpnVrfBgpPathAttrBest"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrbest
    leafs["mplsVpnVrfBgpPathAttrUnknown"] = mplsvpnvrfbgpnbrprefixentry.Mplsvpnvrfbgppathattrunknown
    return leafs
}

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetYangName() string { return "mplsVpnVrfBgpNbrPrefixEntry" }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) SetParent(parent types.Entity) { mplsvpnvrfbgpnbrprefixentry.parent = parent }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetParent() types.Entity { return mplsvpnvrfbgpnbrprefixentry.parent }

func (mplsvpnvrfbgpnbrprefixentry *MPLSVPNMIB_Mplsvpnvrfbgpnbrprefixtable_Mplsvpnvrfbgpnbrprefixentry) GetParentYangName() string { return "mplsVpnVrfBgpNbrPrefixTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every route present
    // configured (either dynamically or statically) within the context of a
    // specific VRF capable of supporting MPLS/BGP VPN. The indexing provides an
    // ordering of VRFs per-VPN interface. The type is slice of
    // MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry.
    Mplsvpnvrfrouteentry []MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry
}

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetFilter() yfilter.YFilter { return mplsvpnvrfroutetable.YFilter }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) SetFilter(yf yfilter.YFilter) { mplsvpnvrfroutetable.YFilter = yf }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetGoName(yname string) string {
    if yname == "mplsVpnVrfRouteEntry" { return "Mplsvpnvrfrouteentry" }
    return ""
}

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetSegmentPath() string {
    return "mplsVpnVrfRouteTable"
}

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsVpnVrfRouteEntry" {
        for _, c := range mplsvpnvrfroutetable.Mplsvpnvrfrouteentry {
            if mplsvpnvrfroutetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry{}
        mplsvpnvrfroutetable.Mplsvpnvrfrouteentry = append(mplsvpnvrfroutetable.Mplsvpnvrfrouteentry, child)
        return &mplsvpnvrfroutetable.Mplsvpnvrfrouteentry[len(mplsvpnvrfroutetable.Mplsvpnvrfrouteentry)-1]
    }
    return nil
}

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsvpnvrfroutetable.Mplsvpnvrfrouteentry {
        children[mplsvpnvrfroutetable.Mplsvpnvrfrouteentry[i].GetSegmentPath()] = &mplsvpnvrfroutetable.Mplsvpnvrfrouteentry[i]
    }
    return children
}

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetYangName() string { return "mplsVpnVrfRouteTable" }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) SetParent(parent types.Entity) { mplsvpnvrfroutetable.parent = parent }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetParent() types.Entity { return mplsvpnvrfroutetable.parent }

func (mplsvpnvrfroutetable *MPLSVPNMIB_Mplsvpnvrfroutetable) GetParentYangName() string { return "MPLS-VPN-MIB" }

// MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry
// An entry in this table is created by an LSR for every route
// present configured (either dynamically or statically) within
// the context of a specific VRF capable of supporting MPLS/BGP
// VPN. The indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetFilter() yfilter.YFilter { return mplsvpnvrfrouteentry.YFilter }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) SetFilter(yf yfilter.YFilter) { mplsvpnvrfrouteentry.YFilter = yf }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetGoName(yname string) string {
    if yname == "mplsVpnVrfName" { return "Mplsvpnvrfname" }
    if yname == "mplsVpnVrfRouteDest" { return "Mplsvpnvrfroutedest" }
    if yname == "mplsVpnVrfRouteMask" { return "Mplsvpnvrfroutemask" }
    if yname == "mplsVpnVrfRouteTos" { return "Mplsvpnvrfroutetos" }
    if yname == "mplsVpnVrfRouteNextHop" { return "Mplsvpnvrfroutenexthop" }
    if yname == "mplsVpnVrfRouteDestAddrType" { return "Mplsvpnvrfroutedestaddrtype" }
    if yname == "mplsVpnVrfRouteMaskAddrType" { return "Mplsvpnvrfroutemaskaddrtype" }
    if yname == "mplsVpnVrfRouteNextHopAddrType" { return "Mplsvpnvrfroutenexthopaddrtype" }
    if yname == "mplsVpnVrfRouteIfIndex" { return "Mplsvpnvrfrouteifindex" }
    if yname == "mplsVpnVrfRouteType" { return "Mplsvpnvrfroutetype" }
    if yname == "mplsVpnVrfRouteProto" { return "Mplsvpnvrfrouteproto" }
    if yname == "mplsVpnVrfRouteAge" { return "Mplsvpnvrfrouteage" }
    if yname == "mplsVpnVrfRouteInfo" { return "Mplsvpnvrfrouteinfo" }
    if yname == "mplsVpnVrfRouteNextHopAS" { return "Mplsvpnvrfroutenexthopas" }
    if yname == "mplsVpnVrfRouteMetric1" { return "Mplsvpnvrfroutemetric1" }
    if yname == "mplsVpnVrfRouteMetric2" { return "Mplsvpnvrfroutemetric2" }
    if yname == "mplsVpnVrfRouteMetric3" { return "Mplsvpnvrfroutemetric3" }
    if yname == "mplsVpnVrfRouteMetric4" { return "Mplsvpnvrfroutemetric4" }
    if yname == "mplsVpnVrfRouteMetric5" { return "Mplsvpnvrfroutemetric5" }
    if yname == "mplsVpnVrfRouteRowStatus" { return "Mplsvpnvrfrouterowstatus" }
    if yname == "mplsVpnVrfRouteStorageType" { return "Mplsvpnvrfroutestoragetype" }
    return ""
}

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetSegmentPath() string {
    return "mplsVpnVrfRouteEntry" + "[mplsVpnVrfName='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfname) + "']" + "[mplsVpnVrfRouteDest='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutedest) + "']" + "[mplsVpnVrfRouteMask='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutemask) + "']" + "[mplsVpnVrfRouteTos='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutetos) + "']" + "[mplsVpnVrfRouteNextHop='" + fmt.Sprintf("%v", mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthop) + "']"
}

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsVpnVrfName"] = mplsvpnvrfrouteentry.Mplsvpnvrfname
    leafs["mplsVpnVrfRouteDest"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutedest
    leafs["mplsVpnVrfRouteMask"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutemask
    leafs["mplsVpnVrfRouteTos"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutetos
    leafs["mplsVpnVrfRouteNextHop"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthop
    leafs["mplsVpnVrfRouteDestAddrType"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutedestaddrtype
    leafs["mplsVpnVrfRouteMaskAddrType"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutemaskaddrtype
    leafs["mplsVpnVrfRouteNextHopAddrType"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthopaddrtype
    leafs["mplsVpnVrfRouteIfIndex"] = mplsvpnvrfrouteentry.Mplsvpnvrfrouteifindex
    leafs["mplsVpnVrfRouteType"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutetype
    leafs["mplsVpnVrfRouteProto"] = mplsvpnvrfrouteentry.Mplsvpnvrfrouteproto
    leafs["mplsVpnVrfRouteAge"] = mplsvpnvrfrouteentry.Mplsvpnvrfrouteage
    leafs["mplsVpnVrfRouteInfo"] = mplsvpnvrfrouteentry.Mplsvpnvrfrouteinfo
    leafs["mplsVpnVrfRouteNextHopAS"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutenexthopas
    leafs["mplsVpnVrfRouteMetric1"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric1
    leafs["mplsVpnVrfRouteMetric2"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric2
    leafs["mplsVpnVrfRouteMetric3"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric3
    leafs["mplsVpnVrfRouteMetric4"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric4
    leafs["mplsVpnVrfRouteMetric5"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutemetric5
    leafs["mplsVpnVrfRouteRowStatus"] = mplsvpnvrfrouteentry.Mplsvpnvrfrouterowstatus
    leafs["mplsVpnVrfRouteStorageType"] = mplsvpnvrfrouteentry.Mplsvpnvrfroutestoragetype
    return leafs
}

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetYangName() string { return "mplsVpnVrfRouteEntry" }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) SetParent(parent types.Entity) { mplsvpnvrfrouteentry.parent = parent }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetParent() types.Entity { return mplsvpnvrfrouteentry.parent }

func (mplsvpnvrfrouteentry *MPLSVPNMIB_Mplsvpnvrfroutetable_Mplsvpnvrfrouteentry) GetParentYangName() string { return "mplsVpnVrfRouteTable" }

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

