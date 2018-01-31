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
    parent types.Entity
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

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetFilter() yfilter.YFilter { return mPLSL3VPNSTDMIB.YFilter }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) SetFilter(yf yfilter.YFilter) { mPLSL3VPNSTDMIB.YFilter = yf }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetGoName(yname string) string {
    if yname == "mplsL3VpnScalars" { return "Mplsl3Vpnscalars" }
    if yname == "mplsL3VpnIfConfTable" { return "Mplsl3Vpnifconftable" }
    if yname == "mplsL3VpnVrfTable" { return "Mplsl3Vpnvrftable" }
    if yname == "mplsL3VpnVrfRTTable" { return "Mplsl3Vpnvrfrttable" }
    if yname == "mplsL3VpnVrfRteTable" { return "Mplsl3Vpnvrfrtetable" }
    return ""
}

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetSegmentPath() string {
    return "MPLS-L3VPN-STD-MIB:MPLS-L3VPN-STD-MIB"
}

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsL3VpnScalars" {
        return &mPLSL3VPNSTDMIB.Mplsl3Vpnscalars
    }
    if childYangName == "mplsL3VpnIfConfTable" {
        return &mPLSL3VPNSTDMIB.Mplsl3Vpnifconftable
    }
    if childYangName == "mplsL3VpnVrfTable" {
        return &mPLSL3VPNSTDMIB.Mplsl3Vpnvrftable
    }
    if childYangName == "mplsL3VpnVrfRTTable" {
        return &mPLSL3VPNSTDMIB.Mplsl3Vpnvrfrttable
    }
    if childYangName == "mplsL3VpnVrfRteTable" {
        return &mPLSL3VPNSTDMIB.Mplsl3Vpnvrfrtetable
    }
    return nil
}

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mplsL3VpnScalars"] = &mPLSL3VPNSTDMIB.Mplsl3Vpnscalars
    children["mplsL3VpnIfConfTable"] = &mPLSL3VPNSTDMIB.Mplsl3Vpnifconftable
    children["mplsL3VpnVrfTable"] = &mPLSL3VPNSTDMIB.Mplsl3Vpnvrftable
    children["mplsL3VpnVrfRTTable"] = &mPLSL3VPNSTDMIB.Mplsl3Vpnvrfrttable
    children["mplsL3VpnVrfRteTable"] = &mPLSL3VPNSTDMIB.Mplsl3Vpnvrfrtetable
    return children
}

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetYangName() string { return "MPLS-L3VPN-STD-MIB" }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) SetParent(parent types.Entity) { mPLSL3VPNSTDMIB.parent = parent }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetParent() types.Entity { return mPLSL3VPNSTDMIB.parent }

func (mPLSL3VPNSTDMIB *MPLSL3VPNSTDMIB) GetParentYangName() string { return "MPLS-L3VPN-STD-MIB" }

// MPLSL3VPNSTDMIB_Mplsl3Vpnscalars
type MPLSL3VPNSTDMIB_Mplsl3Vpnscalars struct {
    parent types.Entity
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

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetFilter() yfilter.YFilter { return mplsl3Vpnscalars.YFilter }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) SetFilter(yf yfilter.YFilter) { mplsl3Vpnscalars.YFilter = yf }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetGoName(yname string) string {
    if yname == "mplsL3VpnConfiguredVrfs" { return "Mplsl3Vpnconfiguredvrfs" }
    if yname == "mplsL3VpnActiveVrfs" { return "Mplsl3Vpnactivevrfs" }
    if yname == "mplsL3VpnConnectedInterfaces" { return "Mplsl3Vpnconnectedinterfaces" }
    if yname == "mplsL3VpnNotificationEnable" { return "Mplsl3Vpnnotificationenable" }
    if yname == "mplsL3VpnVrfConfMaxPossRts" { return "Mplsl3Vpnvrfconfmaxpossrts" }
    if yname == "mplsL3VpnVrfConfRteMxThrshTime" { return "Mplsl3Vpnvrfconfrtemxthrshtime" }
    if yname == "mplsL3VpnIllLblRcvThrsh" { return "Mplsl3Vpnilllblrcvthrsh" }
    return ""
}

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetSegmentPath() string {
    return "mplsL3VpnScalars"
}

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsL3VpnConfiguredVrfs"] = mplsl3Vpnscalars.Mplsl3Vpnconfiguredvrfs
    leafs["mplsL3VpnActiveVrfs"] = mplsl3Vpnscalars.Mplsl3Vpnactivevrfs
    leafs["mplsL3VpnConnectedInterfaces"] = mplsl3Vpnscalars.Mplsl3Vpnconnectedinterfaces
    leafs["mplsL3VpnNotificationEnable"] = mplsl3Vpnscalars.Mplsl3Vpnnotificationenable
    leafs["mplsL3VpnVrfConfMaxPossRts"] = mplsl3Vpnscalars.Mplsl3Vpnvrfconfmaxpossrts
    leafs["mplsL3VpnVrfConfRteMxThrshTime"] = mplsl3Vpnscalars.Mplsl3Vpnvrfconfrtemxthrshtime
    leafs["mplsL3VpnIllLblRcvThrsh"] = mplsl3Vpnscalars.Mplsl3Vpnilllblrcvthrsh
    return leafs
}

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetYangName() string { return "mplsL3VpnScalars" }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) SetParent(parent types.Entity) { mplsl3Vpnscalars.parent = parent }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetParent() types.Entity { return mplsl3Vpnscalars.parent }

func (mplsl3Vpnscalars *MPLSL3VPNSTDMIB_Mplsl3Vpnscalars) GetParentYangName() string { return "MPLS-L3VPN-STD-MIB" }

// MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable
// This table specifies per-interface MPLS capability
// and associated information.
type MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every interface capable of
    // supporting MPLS L3VPN. Each entry in this table is meant to correspond to
    // an entry in the Interfaces Table. The type is slice of
    // MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry.
    Mplsl3Vpnifconfentry []MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry
}

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetFilter() yfilter.YFilter { return mplsl3Vpnifconftable.YFilter }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) SetFilter(yf yfilter.YFilter) { mplsl3Vpnifconftable.YFilter = yf }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetGoName(yname string) string {
    if yname == "mplsL3VpnIfConfEntry" { return "Mplsl3Vpnifconfentry" }
    return ""
}

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetSegmentPath() string {
    return "mplsL3VpnIfConfTable"
}

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsL3VpnIfConfEntry" {
        for _, c := range mplsl3Vpnifconftable.Mplsl3Vpnifconfentry {
            if mplsl3Vpnifconftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry{}
        mplsl3Vpnifconftable.Mplsl3Vpnifconfentry = append(mplsl3Vpnifconftable.Mplsl3Vpnifconfentry, child)
        return &mplsl3Vpnifconftable.Mplsl3Vpnifconfentry[len(mplsl3Vpnifconftable.Mplsl3Vpnifconfentry)-1]
    }
    return nil
}

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsl3Vpnifconftable.Mplsl3Vpnifconfentry {
        children[mplsl3Vpnifconftable.Mplsl3Vpnifconfentry[i].GetSegmentPath()] = &mplsl3Vpnifconftable.Mplsl3Vpnifconfentry[i]
    }
    return children
}

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetYangName() string { return "mplsL3VpnIfConfTable" }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) SetParent(parent types.Entity) { mplsl3Vpnifconftable.parent = parent }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetParent() types.Entity { return mplsl3Vpnifconftable.parent }

func (mplsl3Vpnifconftable *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable) GetParentYangName() string { return "MPLS-L3VPN-STD-MIB" }

// MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry
// An entry in this table is created by an LSR for
// every interface capable of supporting MPLS L3VPN.
// Each entry in this table is meant to correspond to
// an entry in the Interfaces Table.
type MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry struct {
    parent types.Entity
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

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetFilter() yfilter.YFilter { return mplsl3Vpnifconfentry.YFilter }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) SetFilter(yf yfilter.YFilter) { mplsl3Vpnifconfentry.YFilter = yf }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetGoName(yname string) string {
    if yname == "mplsL3VpnVrfName" { return "Mplsl3Vpnvrfname" }
    if yname == "mplsL3VpnIfConfIndex" { return "Mplsl3Vpnifconfindex" }
    if yname == "mplsL3VpnIfVpnClassification" { return "Mplsl3Vpnifvpnclassification" }
    if yname == "mplsL3VpnIfVpnRouteDistProtocol" { return "Mplsl3Vpnifvpnroutedistprotocol" }
    if yname == "mplsL3VpnIfConfStorageType" { return "Mplsl3Vpnifconfstoragetype" }
    if yname == "mplsL3VpnIfConfRowStatus" { return "Mplsl3Vpnifconfrowstatus" }
    return ""
}

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetSegmentPath() string {
    return "mplsL3VpnIfConfEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnifconfentry.Mplsl3Vpnvrfname) + "']" + "[mplsL3VpnIfConfIndex='" + fmt.Sprintf("%v", mplsl3Vpnifconfentry.Mplsl3Vpnifconfindex) + "']"
}

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsL3VpnVrfName"] = mplsl3Vpnifconfentry.Mplsl3Vpnvrfname
    leafs["mplsL3VpnIfConfIndex"] = mplsl3Vpnifconfentry.Mplsl3Vpnifconfindex
    leafs["mplsL3VpnIfVpnClassification"] = mplsl3Vpnifconfentry.Mplsl3Vpnifvpnclassification
    leafs["mplsL3VpnIfVpnRouteDistProtocol"] = mplsl3Vpnifconfentry.Mplsl3Vpnifvpnroutedistprotocol
    leafs["mplsL3VpnIfConfStorageType"] = mplsl3Vpnifconfentry.Mplsl3Vpnifconfstoragetype
    leafs["mplsL3VpnIfConfRowStatus"] = mplsl3Vpnifconfentry.Mplsl3Vpnifconfrowstatus
    return leafs
}

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetYangName() string { return "mplsL3VpnIfConfEntry" }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) SetParent(parent types.Entity) { mplsl3Vpnifconfentry.parent = parent }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetParent() types.Entity { return mplsl3Vpnifconfentry.parent }

func (mplsl3Vpnifconfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnifconftable_Mplsl3Vpnifconfentry) GetParentYangName() string { return "mplsL3VpnIfConfTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for every VRF capable of
    // supporting MPLS L3VPN.  The indexing provides an ordering of VRFs per-VPN
    // interface. The type is slice of
    // MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry.
    Mplsl3Vpnvrfentry []MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry
}

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetFilter() yfilter.YFilter { return mplsl3Vpnvrftable.YFilter }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) SetFilter(yf yfilter.YFilter) { mplsl3Vpnvrftable.YFilter = yf }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetGoName(yname string) string {
    if yname == "mplsL3VpnVrfEntry" { return "Mplsl3Vpnvrfentry" }
    return ""
}

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetSegmentPath() string {
    return "mplsL3VpnVrfTable"
}

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsL3VpnVrfEntry" {
        for _, c := range mplsl3Vpnvrftable.Mplsl3Vpnvrfentry {
            if mplsl3Vpnvrftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry{}
        mplsl3Vpnvrftable.Mplsl3Vpnvrfentry = append(mplsl3Vpnvrftable.Mplsl3Vpnvrfentry, child)
        return &mplsl3Vpnvrftable.Mplsl3Vpnvrfentry[len(mplsl3Vpnvrftable.Mplsl3Vpnvrfentry)-1]
    }
    return nil
}

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsl3Vpnvrftable.Mplsl3Vpnvrfentry {
        children[mplsl3Vpnvrftable.Mplsl3Vpnvrfentry[i].GetSegmentPath()] = &mplsl3Vpnvrftable.Mplsl3Vpnvrfentry[i]
    }
    return children
}

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetYangName() string { return "mplsL3VpnVrfTable" }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) SetParent(parent types.Entity) { mplsl3Vpnvrftable.parent = parent }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetParent() types.Entity { return mplsl3Vpnvrftable.parent }

func (mplsl3Vpnvrftable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable) GetParentYangName() string { return "MPLS-L3VPN-STD-MIB" }

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry
// An entry in this table is created by an LSR for
// every VRF capable of supporting MPLS L3VPN.  The
// indexing provides an ordering of VRFs per-VPN
// interface.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry struct {
    parent types.Entity
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

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetFilter() yfilter.YFilter { return mplsl3Vpnvrfentry.YFilter }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) SetFilter(yf yfilter.YFilter) { mplsl3Vpnvrfentry.YFilter = yf }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetGoName(yname string) string {
    if yname == "mplsL3VpnVrfName" { return "Mplsl3Vpnvrfname" }
    if yname == "mplsL3VpnVrfVpnId" { return "Mplsl3Vpnvrfvpnid" }
    if yname == "mplsL3VpnVrfDescription" { return "Mplsl3Vpnvrfdescription" }
    if yname == "mplsL3VpnVrfRD" { return "Mplsl3Vpnvrfrd" }
    if yname == "mplsL3VpnVrfCreationTime" { return "Mplsl3Vpnvrfcreationtime" }
    if yname == "mplsL3VpnVrfOperStatus" { return "Mplsl3Vpnvrfoperstatus" }
    if yname == "mplsL3VpnVrfActiveInterfaces" { return "Mplsl3Vpnvrfactiveinterfaces" }
    if yname == "mplsL3VpnVrfAssociatedInterfaces" { return "Mplsl3Vpnvrfassociatedinterfaces" }
    if yname == "mplsL3VpnVrfConfMidRteThresh" { return "Mplsl3Vpnvrfconfmidrtethresh" }
    if yname == "mplsL3VpnVrfConfHighRteThresh" { return "Mplsl3Vpnvrfconfhighrtethresh" }
    if yname == "mplsL3VpnVrfConfMaxRoutes" { return "Mplsl3Vpnvrfconfmaxroutes" }
    if yname == "mplsL3VpnVrfConfLastChanged" { return "Mplsl3Vpnvrfconflastchanged" }
    if yname == "mplsL3VpnVrfConfRowStatus" { return "Mplsl3Vpnvrfconfrowstatus" }
    if yname == "mplsL3VpnVrfConfAdminStatus" { return "Mplsl3Vpnvrfconfadminstatus" }
    if yname == "mplsL3VpnVrfConfStorageType" { return "Mplsl3Vpnvrfconfstoragetype" }
    if yname == "mplsL3VpnVrfSecIllegalLblVltns" { return "Mplsl3Vpnvrfsecillegallblvltns" }
    if yname == "mplsL3VpnVrfSecDiscontinuityTime" { return "Mplsl3Vpnvrfsecdiscontinuitytime" }
    if yname == "mplsL3VpnVrfPerfRoutesAdded" { return "Mplsl3Vpnvrfperfroutesadded" }
    if yname == "mplsL3VpnVrfPerfRoutesDeleted" { return "Mplsl3Vpnvrfperfroutesdeleted" }
    if yname == "mplsL3VpnVrfPerfCurrNumRoutes" { return "Mplsl3Vpnvrfperfcurrnumroutes" }
    if yname == "mplsL3VpnVrfPerfRoutesDropped" { return "Mplsl3Vpnvrfperfroutesdropped" }
    if yname == "mplsL3VpnVrfPerfDiscTime" { return "Mplsl3Vpnvrfperfdisctime" }
    return ""
}

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetSegmentPath() string {
    return "mplsL3VpnVrfEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnvrfentry.Mplsl3Vpnvrfname) + "']"
}

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsL3VpnVrfName"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfname
    leafs["mplsL3VpnVrfVpnId"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfvpnid
    leafs["mplsL3VpnVrfDescription"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfdescription
    leafs["mplsL3VpnVrfRD"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfrd
    leafs["mplsL3VpnVrfCreationTime"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfcreationtime
    leafs["mplsL3VpnVrfOperStatus"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfoperstatus
    leafs["mplsL3VpnVrfActiveInterfaces"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfactiveinterfaces
    leafs["mplsL3VpnVrfAssociatedInterfaces"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfassociatedinterfaces
    leafs["mplsL3VpnVrfConfMidRteThresh"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfmidrtethresh
    leafs["mplsL3VpnVrfConfHighRteThresh"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfhighrtethresh
    leafs["mplsL3VpnVrfConfMaxRoutes"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfmaxroutes
    leafs["mplsL3VpnVrfConfLastChanged"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfconflastchanged
    leafs["mplsL3VpnVrfConfRowStatus"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfrowstatus
    leafs["mplsL3VpnVrfConfAdminStatus"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfadminstatus
    leafs["mplsL3VpnVrfConfStorageType"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfconfstoragetype
    leafs["mplsL3VpnVrfSecIllegalLblVltns"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfsecillegallblvltns
    leafs["mplsL3VpnVrfSecDiscontinuityTime"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfsecdiscontinuitytime
    leafs["mplsL3VpnVrfPerfRoutesAdded"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfroutesadded
    leafs["mplsL3VpnVrfPerfRoutesDeleted"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfroutesdeleted
    leafs["mplsL3VpnVrfPerfCurrNumRoutes"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfcurrnumroutes
    leafs["mplsL3VpnVrfPerfRoutesDropped"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfroutesdropped
    leafs["mplsL3VpnVrfPerfDiscTime"] = mplsl3Vpnvrfentry.Mplsl3Vpnvrfperfdisctime
    return leafs
}

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetYangName() string { return "mplsL3VpnVrfEntry" }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) SetParent(parent types.Entity) { mplsl3Vpnvrfentry.parent = parent }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetParent() types.Entity { return mplsl3Vpnvrfentry.parent }

func (mplsl3Vpnvrfentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrftable_Mplsl3Vpnvrfentry) GetParentYangName() string { return "mplsL3VpnVrfTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by an LSR for each route target
    // configured for a VRF supporting a MPLS L3VPN instance.  The indexing
    // provides an ordering per-VRF instance.  See [RFC4364] for a complete
    // definition of a route target. The type is slice of
    // MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry.
    Mplsl3Vpnvrfrtentry []MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry
}

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetFilter() yfilter.YFilter { return mplsl3Vpnvrfrttable.YFilter }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) SetFilter(yf yfilter.YFilter) { mplsl3Vpnvrfrttable.YFilter = yf }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetGoName(yname string) string {
    if yname == "mplsL3VpnVrfRTEntry" { return "Mplsl3Vpnvrfrtentry" }
    return ""
}

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetSegmentPath() string {
    return "mplsL3VpnVrfRTTable"
}

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsL3VpnVrfRTEntry" {
        for _, c := range mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry {
            if mplsl3Vpnvrfrttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry{}
        mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry = append(mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry, child)
        return &mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry[len(mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry)-1]
    }
    return nil
}

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry {
        children[mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry[i].GetSegmentPath()] = &mplsl3Vpnvrfrttable.Mplsl3Vpnvrfrtentry[i]
    }
    return children
}

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetYangName() string { return "mplsL3VpnVrfRTTable" }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) SetParent(parent types.Entity) { mplsl3Vpnvrfrttable.parent = parent }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetParent() types.Entity { return mplsl3Vpnvrfrttable.parent }

func (mplsl3Vpnvrfrttable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable) GetParentYangName() string { return "MPLS-L3VPN-STD-MIB" }

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry
// An entry in this table is created by an LSR for
// each route target configured for a VRF supporting
// a MPLS L3VPN instance.  The indexing provides an
// ordering per-VRF instance.  See [RFC4364] for a
// complete definition of a route target.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry struct {
    parent types.Entity
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

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetFilter() yfilter.YFilter { return mplsl3Vpnvrfrtentry.YFilter }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) SetFilter(yf yfilter.YFilter) { mplsl3Vpnvrfrtentry.YFilter = yf }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetGoName(yname string) string {
    if yname == "mplsL3VpnVrfName" { return "Mplsl3Vpnvrfname" }
    if yname == "mplsL3VpnVrfRTIndex" { return "Mplsl3Vpnvrfrtindex" }
    if yname == "mplsL3VpnVrfRTType" { return "Mplsl3Vpnvrfrttype" }
    if yname == "mplsL3VpnVrfRT" { return "Mplsl3Vpnvrfrt" }
    if yname == "mplsL3VpnVrfRTDescr" { return "Mplsl3Vpnvrfrtdescr" }
    if yname == "mplsL3VpnVrfRTRowStatus" { return "Mplsl3Vpnvrfrtrowstatus" }
    if yname == "mplsL3VpnVrfRTStorageType" { return "Mplsl3Vpnvrfrtstoragetype" }
    return ""
}

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetSegmentPath() string {
    return "mplsL3VpnVrfRTEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfname) + "']" + "[mplsL3VpnVrfRTIndex='" + fmt.Sprintf("%v", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtindex) + "']" + "[mplsL3VpnVrfRTType='" + fmt.Sprintf("%v", mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrttype) + "']"
}

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsL3VpnVrfName"] = mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfname
    leafs["mplsL3VpnVrfRTIndex"] = mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtindex
    leafs["mplsL3VpnVrfRTType"] = mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrttype
    leafs["mplsL3VpnVrfRT"] = mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrt
    leafs["mplsL3VpnVrfRTDescr"] = mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtdescr
    leafs["mplsL3VpnVrfRTRowStatus"] = mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtrowstatus
    leafs["mplsL3VpnVrfRTStorageType"] = mplsl3Vpnvrfrtentry.Mplsl3Vpnvrfrtstoragetype
    return leafs
}

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetYangName() string { return "mplsL3VpnVrfRTEntry" }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) SetParent(parent types.Entity) { mplsl3Vpnvrfrtentry.parent = parent }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetParent() types.Entity { return mplsl3Vpnvrfrtentry.parent }

func (mplsl3Vpnvrfrtentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrttable_Mplsl3Vpnvrfrtentry) GetParentYangName() string { return "mplsL3VpnVrfRTTable" }

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable
// This table specifies per-interface MPLS L3VPN VRF Table
// routing information.  Entries in this table define VRF routing
// entries associated with the specified MPLS/VPN interfaces.  Note
// 
// that this table contains both BGP and Interior Gateway Protocol
// IGP routes, as both may appear in the same VRF.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable struct {
    parent types.Entity
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

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetFilter() yfilter.YFilter { return mplsl3Vpnvrfrtetable.YFilter }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) SetFilter(yf yfilter.YFilter) { mplsl3Vpnvrfrtetable.YFilter = yf }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetGoName(yname string) string {
    if yname == "mplsL3VpnVrfRteEntry" { return "Mplsl3Vpnvrfrteentry" }
    return ""
}

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetSegmentPath() string {
    return "mplsL3VpnVrfRteTable"
}

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mplsL3VpnVrfRteEntry" {
        for _, c := range mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry {
            if mplsl3Vpnvrfrtetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry{}
        mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry = append(mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry, child)
        return &mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry[len(mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry)-1]
    }
    return nil
}

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry {
        children[mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry[i].GetSegmentPath()] = &mplsl3Vpnvrfrtetable.Mplsl3Vpnvrfrteentry[i]
    }
    return children
}

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetYangName() string { return "mplsL3VpnVrfRteTable" }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) SetParent(parent types.Entity) { mplsl3Vpnvrfrtetable.parent = parent }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetParent() types.Entity { return mplsl3Vpnvrfrtetable.parent }

func (mplsl3Vpnvrfrtetable *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable) GetParentYangName() string { return "MPLS-L3VPN-STD-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetFilter() yfilter.YFilter { return mplsl3Vpnvrfrteentry.YFilter }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) SetFilter(yf yfilter.YFilter) { mplsl3Vpnvrfrteentry.YFilter = yf }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetGoName(yname string) string {
    if yname == "mplsL3VpnVrfName" { return "Mplsl3Vpnvrfname" }
    if yname == "mplsL3VpnVrfRteInetCidrDestType" { return "Mplsl3Vpnvrfrteinetcidrdesttype" }
    if yname == "mplsL3VpnVrfRteInetCidrDest" { return "Mplsl3Vpnvrfrteinetcidrdest" }
    if yname == "mplsL3VpnVrfRteInetCidrPfxLen" { return "Mplsl3Vpnvrfrteinetcidrpfxlen" }
    if yname == "mplsL3VpnVrfRteInetCidrPolicy" { return "Mplsl3Vpnvrfrteinetcidrpolicy" }
    if yname == "mplsL3VpnVrfRteInetCidrNHopType" { return "Mplsl3Vpnvrfrteinetcidrnhoptype" }
    if yname == "mplsL3VpnVrfRteInetCidrNextHop" { return "Mplsl3Vpnvrfrteinetcidrnexthop" }
    if yname == "mplsL3VpnVrfRteInetCidrIfIndex" { return "Mplsl3Vpnvrfrteinetcidrifindex" }
    if yname == "mplsL3VpnVrfRteInetCidrType" { return "Mplsl3Vpnvrfrteinetcidrtype" }
    if yname == "mplsL3VpnVrfRteInetCidrProto" { return "Mplsl3Vpnvrfrteinetcidrproto" }
    if yname == "mplsL3VpnVrfRteInetCidrAge" { return "Mplsl3Vpnvrfrteinetcidrage" }
    if yname == "mplsL3VpnVrfRteInetCidrNextHopAS" { return "Mplsl3Vpnvrfrteinetcidrnexthopas" }
    if yname == "mplsL3VpnVrfRteInetCidrMetric1" { return "Mplsl3Vpnvrfrteinetcidrmetric1" }
    if yname == "mplsL3VpnVrfRteInetCidrMetric2" { return "Mplsl3Vpnvrfrteinetcidrmetric2" }
    if yname == "mplsL3VpnVrfRteInetCidrMetric3" { return "Mplsl3Vpnvrfrteinetcidrmetric3" }
    if yname == "mplsL3VpnVrfRteInetCidrMetric4" { return "Mplsl3Vpnvrfrteinetcidrmetric4" }
    if yname == "mplsL3VpnVrfRteInetCidrMetric5" { return "Mplsl3Vpnvrfrteinetcidrmetric5" }
    if yname == "mplsL3VpnVrfRteXCPointer" { return "Mplsl3Vpnvrfrtexcpointer" }
    if yname == "mplsL3VpnVrfRteInetCidrStatus" { return "Mplsl3Vpnvrfrteinetcidrstatus" }
    return ""
}

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetSegmentPath() string {
    return "mplsL3VpnVrfRteEntry" + "[mplsL3VpnVrfName='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfname) + "']" + "[mplsL3VpnVrfRteInetCidrDestType='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdesttype) + "']" + "[mplsL3VpnVrfRteInetCidrDest='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdest) + "']" + "[mplsL3VpnVrfRteInetCidrPfxLen='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpfxlen) + "']" + "[mplsL3VpnVrfRteInetCidrPolicy='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpolicy) + "']" + "[mplsL3VpnVrfRteInetCidrNHopType='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnhoptype) + "']" + "[mplsL3VpnVrfRteInetCidrNextHop='" + fmt.Sprintf("%v", mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnexthop) + "']"
}

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsL3VpnVrfName"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfname
    leafs["mplsL3VpnVrfRteInetCidrDestType"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdesttype
    leafs["mplsL3VpnVrfRteInetCidrDest"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrdest
    leafs["mplsL3VpnVrfRteInetCidrPfxLen"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpfxlen
    leafs["mplsL3VpnVrfRteInetCidrPolicy"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrpolicy
    leafs["mplsL3VpnVrfRteInetCidrNHopType"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnhoptype
    leafs["mplsL3VpnVrfRteInetCidrNextHop"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnexthop
    leafs["mplsL3VpnVrfRteInetCidrIfIndex"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrifindex
    leafs["mplsL3VpnVrfRteInetCidrType"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrtype
    leafs["mplsL3VpnVrfRteInetCidrProto"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrproto
    leafs["mplsL3VpnVrfRteInetCidrAge"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrage
    leafs["mplsL3VpnVrfRteInetCidrNextHopAS"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrnexthopas
    leafs["mplsL3VpnVrfRteInetCidrMetric1"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric1
    leafs["mplsL3VpnVrfRteInetCidrMetric2"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric2
    leafs["mplsL3VpnVrfRteInetCidrMetric3"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric3
    leafs["mplsL3VpnVrfRteInetCidrMetric4"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric4
    leafs["mplsL3VpnVrfRteInetCidrMetric5"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrmetric5
    leafs["mplsL3VpnVrfRteXCPointer"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrtexcpointer
    leafs["mplsL3VpnVrfRteInetCidrStatus"] = mplsl3Vpnvrfrteentry.Mplsl3Vpnvrfrteinetcidrstatus
    return leafs
}

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetBundleName() string { return "cisco_ios_xe" }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetYangName() string { return "mplsL3VpnVrfRteEntry" }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) SetParent(parent types.Entity) { mplsl3Vpnvrfrteentry.parent = parent }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetParent() types.Entity { return mplsl3Vpnvrfrteentry.parent }

func (mplsl3Vpnvrfrteentry *MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry) GetParentYangName() string { return "mplsL3VpnVrfRteTable" }

// MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype represents discards the message silently.
type MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype string

const (
    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_other MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "other"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_reject MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "reject"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_local MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "local"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_remote MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "remote"

    MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype_blackhole MPLSL3VPNSTDMIB_Mplsl3Vpnvrfrtetable_Mplsl3Vpnvrfrteentry_Mplsl3Vpnvrfrteinetcidrtype = "blackhole"
)

