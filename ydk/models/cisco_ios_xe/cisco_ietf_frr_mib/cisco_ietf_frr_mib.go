// This MIB module contains managed object definitions for MPLS
// Fast Reroute (FRR) as defined in:Pan, P., Gan, D., Swallow, G.,
// Vasseur, J.Ph., Cooper, D., Atlas, A., Jork, M., Fast Reroute
// Techniques in RSVP-TE, draft-ietf-mpls-rsvp-lsp-fastreroute-
// 00.txt, January 2002.
package cisco_ietf_frr_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_frr_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-FRR-MIB CISCO-IETF-FRR-MIB}", reflect.TypeOf(CISCOIETFFRRMIB{}))
    ydk.RegisterEntity("CISCO-IETF-FRR-MIB:CISCO-IETF-FRR-MIB", reflect.TypeOf(CISCOIETFFRRMIB{}))
}

// CISCOIETFFRRMIB
type CISCOIETFFRRMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cmplsfrrscalars CISCOIETFFRRMIB_Cmplsfrrscalars

    // This table shows detour setup constraints.
    Cmplsfrrconsttable CISCOIETFFRRMIB_Cmplsfrrconsttable

    // The fast reroute log table records fast reroute events such as protected
    // links going up or down or the FRR feature kicking in.
    Cmplsfrrlogtable CISCOIETFFRRMIB_Cmplsfrrlogtable

    // The mplsFrrFacRouteDBTable provides information about the  fast reroute
    // database.  Each entry belongs to an interface, protecting backup tunnel and
    // protected tunnel. MPLS  interfaces defined on this node are protected by
    // backup tunnels and are indexed by mplsFrrFacRouteProtectedIndex. Backup
    // tunnels defined to protect the tunnels traversing an interface, and are
    // indexed by  mplsFrrFacRouteProtectingTunIndex.  Note that the tunnel 
    // instance index is not required, since it is implied to be 0,  which
    // indicates the tunnel head interface for the protecting  tunnel. The
    // protecting tunnel is defined to exist on the PLR  in the FRR specification.
    // Protected tunnels are the LSPs that  traverse the protected link.  These
    // LSPs are uniquely  identified by mplsFrrFacRouteProtectedTunIndex,
    // mplsFrrFacRouteProtectedTunInstance, 
    // mplsFrrFacRouteProtectedTunIngressLSRId, and 
    // mplsFrrFacRouteProtectedTunEgressLSRId.
    Cmplsfrrfacroutedbtable CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable
}

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetFilter() yfilter.YFilter { return cISCOIETFFRRMIB.YFilter }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFFRRMIB.YFilter = yf }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetGoName(yname string) string {
    if yname == "cmplsFrrScalars" { return "Cmplsfrrscalars" }
    if yname == "cmplsFrrConstTable" { return "Cmplsfrrconsttable" }
    if yname == "cmplsFrrLogTable" { return "Cmplsfrrlogtable" }
    if yname == "cmplsFrrFacRouteDBTable" { return "Cmplsfrrfacroutedbtable" }
    return ""
}

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetSegmentPath() string {
    return "CISCO-IETF-FRR-MIB:CISCO-IETF-FRR-MIB"
}

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsFrrScalars" {
        return &cISCOIETFFRRMIB.Cmplsfrrscalars
    }
    if childYangName == "cmplsFrrConstTable" {
        return &cISCOIETFFRRMIB.Cmplsfrrconsttable
    }
    if childYangName == "cmplsFrrLogTable" {
        return &cISCOIETFFRRMIB.Cmplsfrrlogtable
    }
    if childYangName == "cmplsFrrFacRouteDBTable" {
        return &cISCOIETFFRRMIB.Cmplsfrrfacroutedbtable
    }
    return nil
}

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cmplsFrrScalars"] = &cISCOIETFFRRMIB.Cmplsfrrscalars
    children["cmplsFrrConstTable"] = &cISCOIETFFRRMIB.Cmplsfrrconsttable
    children["cmplsFrrLogTable"] = &cISCOIETFFRRMIB.Cmplsfrrlogtable
    children["cmplsFrrFacRouteDBTable"] = &cISCOIETFFRRMIB.Cmplsfrrfacroutedbtable
    return children
}

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetYangName() string { return "CISCO-IETF-FRR-MIB" }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) SetParent(parent types.Entity) { cISCOIETFFRRMIB.parent = parent }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetParent() types.Entity { return cISCOIETFFRRMIB.parent }

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetParentYangName() string { return "CISCO-IETF-FRR-MIB" }

// CISCOIETFFRRMIB_Cmplsfrrscalars
type CISCOIETFFRRMIB_Cmplsfrrscalars struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of detour LSPs entering the device if
    // mplsFrrConstProtectionMethod is set to oneToOneBackup(0), or or 0 if
    // mplsFrrConstProtectionMethod is set to facilityBackup(1). The type is
    // interface{} with range: 0..4294967295.
    Cmplsfrrdetourincoming interface{}

    // The number of detour LSPs leaving the device if
    // mplsFrrConstProtectionMethod is set to oneToOneBackup(0), or 0 if
    // mplsFrrConstProtectionMethod is set to  to facilityBackup(1). The type is
    // interface{} with range: 0..4294967295.
    Cmplsfrrdetouroutgoing interface{}

    // The number of detour LSPs originating at this PLR if
    // mplsFrrConstProtectionMethod is set to oneToOneBackup(0). This object MUST
    // return 0 if the mplsFrrConstProtectionMethod  is set to facilityBackup(1).
    // The type is interface{} with range: 0..4294967295.
    Cmplsfrrdetouroriginating interface{}

    // The number of tunnel instances that are switched over to their
    // corresponding detour LSP if mplsFrrConstProtectionMethod is set to
    // oneToOneBackup(0), or tunnels being switched over if
    // mplsFrrConstProtectionMethod is set to facilityBackup(1). The type is
    // interface{} with range: 0..4294967295.
    Cmplsfrrswitchover interface{}

    // Indicates the number of MPLS interfaces configured for  protection by the
    // FRR feature, otherwise this value MUST return 0 to indicate that LSPs
    // traversing any  interface may be protected. The type is interface{} with
    // range: 0..4294967295.
    Cmplsfrrnumofconfifs interface{}

    // Indicates the number of interfaces currently being protected  by the FRR
    // feature if mplsFrrConstProtectionMethod is set to facilityBackup(1),
    // otherwise this value should return 0 to indicate that LSPs traversing any
    // interface may be protected. This value MUST be less than or equal to
    // mplsFrrConfIfs. The type is interface{} with range: 0..4294967295.
    Cmplsfrractprotectedifs interface{}

    // Indicates the number of bypass tunnels configured to  protect facilities on
    // this LSR using the FRR feature  if mplsFrrConstProtectionMethod is set to 
    // facilityBackup(1), otherwise this value MUST return  0. The type is
    // interface{} with range: 0..4294967295.
    Cmplsfrrconfprotectingtuns interface{}

    // Indicates the number of bypass tunnels indicated in
    // mplsFrrConfProtectingTuns whose operStatus is up(1) indicating that they
    // are currently protecting facilities on this LSR using the FRR feature. This
    // object MUST return 0 if mplsFrrConstProtectionMethod  is set to
    // facilityBackup(1). The type is interface{} with range: 0..4294967295.
    Cmplsfrractprotectedtuns interface{}

    // Indicates the number of LSPs currently protected by  the FRR feature. If
    // mplsFrrConstProtectionMethod is set  to facilityBackup(1)this object MUST
    // return 0. The type is interface{} with range: 0..4294967295.
    Cmplsfrractprotectedlsps interface{}

    // Indicates which protection method is to be used for fast reroute. Some
    // devices may require a reboot of their routing processors if this variable
    // is changed. An agent which does not wish to reboot or modify its FRR mode 
    // MUST return an inconsistentValue error. Please  consult the device's agent
    // capability statement  for more details. The type is
    // Cmplsfrrconstprotectionmethod.
    Cmplsfrrconstprotectionmethod interface{}

    // Enables or disables FRR notifications defined in this MIB module.
    // Notifications are disabled by default. The type is bool.
    Cmplsfrrnotifsenabled interface{}

    // Indicates the maximum number of entries allowed in the FRR Log table.
    // Agents receiving SETs for values that cannot be used must return an
    // inconsistent value error. If a manager sets this value to 0, this indicates
    // that no logging should take place by the agent.    If this value is
    // returned as 0, this indicates that no additional log entries will be added
    // to the current table either because the table has been completely filled or
    // logging has been disabled. However, agents may wish to not delete existing
    // entries in the log table so that managers may review them in the future.  
    // It is implied that when mplsFrrLogTableCurrEntries  has reached the value
    // of this variable, that logging  entries may not continue to be added to the
    // table,  although existing ones may remain.  Furthermore, an agent may begin
    // to delete existing (perhaps the oldest entries) entries to make room for
    // new ones. The type is interface{} with range: 0..4294967295.
    Cmplsfrrlogtablemaxentries interface{}

    // Indicates the current number of entries in the FRR log table. The type is
    // interface{} with range: 0..4294967295.
    Cmplsfrrlogtablecurrentries interface{}

    // This variable indicates the number of milliseconds that must elapse between
    // notification emissions. If events occur more rapidly, the implementation
    // may simply fail to emit these notifications during that period, or may
    // queue them until an appropriate time in the future. A value of 0 means no
    // minimum  elapsed period is specified. The type is interface{} with range:
    // 0..4294967295.
    Cmplsfrrnotifmaxrate interface{}
}

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetFilter() yfilter.YFilter { return cmplsfrrscalars.YFilter }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) SetFilter(yf yfilter.YFilter) { cmplsfrrscalars.YFilter = yf }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetGoName(yname string) string {
    if yname == "cmplsFrrDetourIncoming" { return "Cmplsfrrdetourincoming" }
    if yname == "cmplsFrrDetourOutgoing" { return "Cmplsfrrdetouroutgoing" }
    if yname == "cmplsFrrDetourOriginating" { return "Cmplsfrrdetouroriginating" }
    if yname == "cmplsFrrSwitchover" { return "Cmplsfrrswitchover" }
    if yname == "cmplsFrrNumOfConfIfs" { return "Cmplsfrrnumofconfifs" }
    if yname == "cmplsFrrActProtectedIfs" { return "Cmplsfrractprotectedifs" }
    if yname == "cmplsFrrConfProtectingTuns" { return "Cmplsfrrconfprotectingtuns" }
    if yname == "cmplsFrrActProtectedTuns" { return "Cmplsfrractprotectedtuns" }
    if yname == "cmplsFrrActProtectedLSPs" { return "Cmplsfrractprotectedlsps" }
    if yname == "cmplsFrrConstProtectionMethod" { return "Cmplsfrrconstprotectionmethod" }
    if yname == "cmplsFrrNotifsEnabled" { return "Cmplsfrrnotifsenabled" }
    if yname == "cmplsFrrLogTableMaxEntries" { return "Cmplsfrrlogtablemaxentries" }
    if yname == "cmplsFrrLogTableCurrEntries" { return "Cmplsfrrlogtablecurrentries" }
    if yname == "cmplsFrrNotifMaxRate" { return "Cmplsfrrnotifmaxrate" }
    return ""
}

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetSegmentPath() string {
    return "cmplsFrrScalars"
}

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsFrrDetourIncoming"] = cmplsfrrscalars.Cmplsfrrdetourincoming
    leafs["cmplsFrrDetourOutgoing"] = cmplsfrrscalars.Cmplsfrrdetouroutgoing
    leafs["cmplsFrrDetourOriginating"] = cmplsfrrscalars.Cmplsfrrdetouroriginating
    leafs["cmplsFrrSwitchover"] = cmplsfrrscalars.Cmplsfrrswitchover
    leafs["cmplsFrrNumOfConfIfs"] = cmplsfrrscalars.Cmplsfrrnumofconfifs
    leafs["cmplsFrrActProtectedIfs"] = cmplsfrrscalars.Cmplsfrractprotectedifs
    leafs["cmplsFrrConfProtectingTuns"] = cmplsfrrscalars.Cmplsfrrconfprotectingtuns
    leafs["cmplsFrrActProtectedTuns"] = cmplsfrrscalars.Cmplsfrractprotectedtuns
    leafs["cmplsFrrActProtectedLSPs"] = cmplsfrrscalars.Cmplsfrractprotectedlsps
    leafs["cmplsFrrConstProtectionMethod"] = cmplsfrrscalars.Cmplsfrrconstprotectionmethod
    leafs["cmplsFrrNotifsEnabled"] = cmplsfrrscalars.Cmplsfrrnotifsenabled
    leafs["cmplsFrrLogTableMaxEntries"] = cmplsfrrscalars.Cmplsfrrlogtablemaxentries
    leafs["cmplsFrrLogTableCurrEntries"] = cmplsfrrscalars.Cmplsfrrlogtablecurrentries
    leafs["cmplsFrrNotifMaxRate"] = cmplsfrrscalars.Cmplsfrrnotifmaxrate
    return leafs
}

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetYangName() string { return "cmplsFrrScalars" }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) SetParent(parent types.Entity) { cmplsfrrscalars.parent = parent }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetParent() types.Entity { return cmplsfrrscalars.parent }

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetParentYangName() string { return "CISCO-IETF-FRR-MIB" }

// CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod represents for more details.
type CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod string

const (
    CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod_oneToOneBackup CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod = "oneToOneBackup"

    CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod_facilityBackup CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod = "facilityBackup"
)

// CISCOIETFFRRMIB_Cmplsfrrconsttable
// This table shows detour setup constraints.
type CISCOIETFFRRMIB_Cmplsfrrconsttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents detour LSP or bypass tunnel  setup
    // constraints for a tunnel instance to be protected by  detour LSPs or a
    // tunnel. Agents must allow entries in this table  to be created only for
    // tunnel instances that require fast-reroute. Entries indexed with
    // mplsFrrConstIfIndex set to 0 apply to all interfaces on this device for
    // which the FRR feature can operate on. The type is slice of
    // CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry.
    Cmplsfrrconstentry []CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry
}

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetFilter() yfilter.YFilter { return cmplsfrrconsttable.YFilter }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) SetFilter(yf yfilter.YFilter) { cmplsfrrconsttable.YFilter = yf }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetGoName(yname string) string {
    if yname == "cmplsFrrConstEntry" { return "Cmplsfrrconstentry" }
    return ""
}

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetSegmentPath() string {
    return "cmplsFrrConstTable"
}

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsFrrConstEntry" {
        for _, c := range cmplsfrrconsttable.Cmplsfrrconstentry {
            if cmplsfrrconsttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry{}
        cmplsfrrconsttable.Cmplsfrrconstentry = append(cmplsfrrconsttable.Cmplsfrrconstentry, child)
        return &cmplsfrrconsttable.Cmplsfrrconstentry[len(cmplsfrrconsttable.Cmplsfrrconstentry)-1]
    }
    return nil
}

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplsfrrconsttable.Cmplsfrrconstentry {
        children[cmplsfrrconsttable.Cmplsfrrconstentry[i].GetSegmentPath()] = &cmplsfrrconsttable.Cmplsfrrconstentry[i]
    }
    return children
}

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetYangName() string { return "cmplsFrrConstTable" }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) SetParent(parent types.Entity) { cmplsfrrconsttable.parent = parent }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetParent() types.Entity { return cmplsfrrconsttable.parent }

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetParentYangName() string { return "CISCO-IETF-FRR-MIB" }

// CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry
// An entry in this table represents detour LSP or bypass tunnel 
// setup constraints for a tunnel instance to be protected by 
// detour LSPs or a tunnel. Agents must allow entries in this table 
// to be created only for tunnel instances that require fast-reroute.
// Entries indexed with mplsFrrConstIfIndex set to 0 apply to all
// interfaces on this device for which the FRR feature can operate
// on.
type CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies an interface for which fast
    // reroute is configured. Tabular entries indexed with a 0 value apply to all
    // interfaces on this device for which the FRR feature can operate on. The
    // type is interface{} with range: 0..2147483647.
    Cmplsfrrconstifindex interface{}

    // This attribute is a key. Uniquely identifies a tunnel for which fast
    // reroute is requested. The type is interface{} with range: 0..65535.
    Cmplsfrrconsttunnelindex interface{}

    // This attribute is a key. Uniquely identifies an instance of this tunnel for
    // which fast reroute is requested. The type is interface{} with range:
    // 0..4294967295.
    Cmplsfrrconsttunnelinstance interface{}

    // Indicates the setup priority of detour LSP. The type is interface{} with
    // range: 0..7.
    Cmplsfrrconstsetupprio interface{}

    // Indicates the holding priority for detour LSP. The type is interface{} with
    // range: 0..7.
    Cmplsfrrconstholdingprio interface{}

    // A link satisfies the include-any constraint if and only if the constraint
    // is zero, or the link and the constraint have a resource class in common.
    // The type is interface{} with range: 0..4294967295.
    Cmplsfrrconstinclanyaffinity interface{}

    // A link satisfies the include-all constraint if and only if the link
    // contains all of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    Cmplsfrrconstinclallaffinity interface{}

    // A link satisfies the exclude-all constraint if and only if the link
    // contains none of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    Cmplsfrrconstexclallaffinity interface{}

    // The maximum number of hops that the detour LSP may traverse. The type is
    // interface{} with range: 1..65535.
    Cmplsfrrconsthoplimit interface{}

    // This variable represents the bandwidth for detour LSPs of this tunnel, in
    // units of thousands of bits per second (Kbps). The type is interface{} with
    // range: 0..4294967295.
    Cmplsfrrconstbandwidth interface{}

    // This object is used to create, modify, and/or delete a row in this table.
    // The type is RowStatus.
    Cmplsfrrconstrowstatus interface{}

    // The number of backup tunnels protecting the specified interface. The type
    // is interface{} with range: 0..4294967295.
    Cmplsfrrconstnumprotectingtunonif interface{}

    // The number of tunnels protected on this interface. The type is interface{}
    // with range: 0..4294967295.
    Cmplsfrrconstnumprotectedtunonif interface{}
}

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetFilter() yfilter.YFilter { return cmplsfrrconstentry.YFilter }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) SetFilter(yf yfilter.YFilter) { cmplsfrrconstentry.YFilter = yf }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetGoName(yname string) string {
    if yname == "cmplsFrrConstIfIndex" { return "Cmplsfrrconstifindex" }
    if yname == "cmplsFrrConstTunnelIndex" { return "Cmplsfrrconsttunnelindex" }
    if yname == "cmplsFrrConstTunnelInstance" { return "Cmplsfrrconsttunnelinstance" }
    if yname == "cmplsFrrConstSetupPrio" { return "Cmplsfrrconstsetupprio" }
    if yname == "cmplsFrrConstHoldingPrio" { return "Cmplsfrrconstholdingprio" }
    if yname == "cmplsFrrConstInclAnyAffinity" { return "Cmplsfrrconstinclanyaffinity" }
    if yname == "cmplsFrrConstInclAllAffinity" { return "Cmplsfrrconstinclallaffinity" }
    if yname == "cmplsFrrConstExclAllAffinity" { return "Cmplsfrrconstexclallaffinity" }
    if yname == "cmplsFrrConstHopLimit" { return "Cmplsfrrconsthoplimit" }
    if yname == "cmplsFrrConstBandwidth" { return "Cmplsfrrconstbandwidth" }
    if yname == "cmplsFrrConstRowStatus" { return "Cmplsfrrconstrowstatus" }
    if yname == "cmplsFrrConstNumProtectingTunOnIf" { return "Cmplsfrrconstnumprotectingtunonif" }
    if yname == "cmplsFrrConstNumProtectedTunOnIf" { return "Cmplsfrrconstnumprotectedtunonif" }
    return ""
}

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetSegmentPath() string {
    return "cmplsFrrConstEntry" + "[cmplsFrrConstIfIndex='" + fmt.Sprintf("%v", cmplsfrrconstentry.Cmplsfrrconstifindex) + "']" + "[cmplsFrrConstTunnelIndex='" + fmt.Sprintf("%v", cmplsfrrconstentry.Cmplsfrrconsttunnelindex) + "']" + "[cmplsFrrConstTunnelInstance='" + fmt.Sprintf("%v", cmplsfrrconstentry.Cmplsfrrconsttunnelinstance) + "']"
}

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsFrrConstIfIndex"] = cmplsfrrconstentry.Cmplsfrrconstifindex
    leafs["cmplsFrrConstTunnelIndex"] = cmplsfrrconstentry.Cmplsfrrconsttunnelindex
    leafs["cmplsFrrConstTunnelInstance"] = cmplsfrrconstentry.Cmplsfrrconsttunnelinstance
    leafs["cmplsFrrConstSetupPrio"] = cmplsfrrconstentry.Cmplsfrrconstsetupprio
    leafs["cmplsFrrConstHoldingPrio"] = cmplsfrrconstentry.Cmplsfrrconstholdingprio
    leafs["cmplsFrrConstInclAnyAffinity"] = cmplsfrrconstentry.Cmplsfrrconstinclanyaffinity
    leafs["cmplsFrrConstInclAllAffinity"] = cmplsfrrconstentry.Cmplsfrrconstinclallaffinity
    leafs["cmplsFrrConstExclAllAffinity"] = cmplsfrrconstentry.Cmplsfrrconstexclallaffinity
    leafs["cmplsFrrConstHopLimit"] = cmplsfrrconstentry.Cmplsfrrconsthoplimit
    leafs["cmplsFrrConstBandwidth"] = cmplsfrrconstentry.Cmplsfrrconstbandwidth
    leafs["cmplsFrrConstRowStatus"] = cmplsfrrconstentry.Cmplsfrrconstrowstatus
    leafs["cmplsFrrConstNumProtectingTunOnIf"] = cmplsfrrconstentry.Cmplsfrrconstnumprotectingtunonif
    leafs["cmplsFrrConstNumProtectedTunOnIf"] = cmplsfrrconstentry.Cmplsfrrconstnumprotectedtunonif
    return leafs
}

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetYangName() string { return "cmplsFrrConstEntry" }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) SetParent(parent types.Entity) { cmplsfrrconstentry.parent = parent }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetParent() types.Entity { return cmplsfrrconstentry.parent }

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetParentYangName() string { return "cmplsFrrConstTable" }

// CISCOIETFFRRMIB_Cmplsfrrlogtable
// The fast reroute log table records fast reroute events such
// as protected links going up or down or the FRR feature
// kicking in.
type CISCOIETFFRRMIB_Cmplsfrrlogtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created to describe one fast reroute event. 
    // Entries in this table are only created and destroyed by the agent
    // implementation. The maximum number  of entries in this log is governed by
    // the scalar. The type is slice of
    // CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry.
    Cmplsfrrlogentry []CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry
}

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetFilter() yfilter.YFilter { return cmplsfrrlogtable.YFilter }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) SetFilter(yf yfilter.YFilter) { cmplsfrrlogtable.YFilter = yf }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetGoName(yname string) string {
    if yname == "cmplsFrrLogEntry" { return "Cmplsfrrlogentry" }
    return ""
}

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetSegmentPath() string {
    return "cmplsFrrLogTable"
}

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsFrrLogEntry" {
        for _, c := range cmplsfrrlogtable.Cmplsfrrlogentry {
            if cmplsfrrlogtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry{}
        cmplsfrrlogtable.Cmplsfrrlogentry = append(cmplsfrrlogtable.Cmplsfrrlogentry, child)
        return &cmplsfrrlogtable.Cmplsfrrlogentry[len(cmplsfrrlogtable.Cmplsfrrlogentry)-1]
    }
    return nil
}

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplsfrrlogtable.Cmplsfrrlogentry {
        children[cmplsfrrlogtable.Cmplsfrrlogentry[i].GetSegmentPath()] = &cmplsfrrlogtable.Cmplsfrrlogentry[i]
    }
    return children
}

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetYangName() string { return "cmplsFrrLogTable" }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) SetParent(parent types.Entity) { cmplsfrrlogtable.parent = parent }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetParent() types.Entity { return cmplsfrrlogtable.parent }

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetParentYangName() string { return "CISCO-IETF-FRR-MIB" }

// CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry
// An entry in this table is created to describe one fast
// reroute event.  Entries in this table are only created and
// destroyed by the agent implementation. The maximum number 
// of entries in this log is governed by the scalar.
type CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a fast reroute event entry.
    // The type is interface{} with range: 0..4294967295.
    Cmplsfrrlogindex interface{}

    // This object provides the amount of time ticks since this event occured. The
    // type is interface{} with range: 0..4294967295.
    Cmplsfrrlogeventtime interface{}

    // This object indicates which interface was affected by this FRR event. This
    // value may be set to 0 if mplsFrrConstProtectionMethod is set to
    // oneToOneBackup(0). The type is interface{} with range: 0..2147483647.
    Cmplsfrrloginterface interface{}

    // This object describes what type of fast reroute event occured. The type is
    // Cmplsfrrlogeventtype.
    Cmplsfrrlogeventtype interface{}

    // This object describes the duration of this event. The type is interface{}
    // with range: 0..4294967295.
    Cmplsfrrlogeventduration interface{}

    // This object contains an implementation-specific explanation of the event.
    // The type is string with length: 128.
    Cmplsfrrlogeventreasonstring interface{}
}

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetFilter() yfilter.YFilter { return cmplsfrrlogentry.YFilter }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) SetFilter(yf yfilter.YFilter) { cmplsfrrlogentry.YFilter = yf }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetGoName(yname string) string {
    if yname == "cmplsFrrLogIndex" { return "Cmplsfrrlogindex" }
    if yname == "cmplsFrrLogEventTime" { return "Cmplsfrrlogeventtime" }
    if yname == "cmplsFrrLogInterface" { return "Cmplsfrrloginterface" }
    if yname == "cmplsFrrLogEventType" { return "Cmplsfrrlogeventtype" }
    if yname == "cmplsFrrLogEventDuration" { return "Cmplsfrrlogeventduration" }
    if yname == "cmplsFrrLogEventReasonString" { return "Cmplsfrrlogeventreasonstring" }
    return ""
}

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetSegmentPath() string {
    return "cmplsFrrLogEntry" + "[cmplsFrrLogIndex='" + fmt.Sprintf("%v", cmplsfrrlogentry.Cmplsfrrlogindex) + "']"
}

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsFrrLogIndex"] = cmplsfrrlogentry.Cmplsfrrlogindex
    leafs["cmplsFrrLogEventTime"] = cmplsfrrlogentry.Cmplsfrrlogeventtime
    leafs["cmplsFrrLogInterface"] = cmplsfrrlogentry.Cmplsfrrloginterface
    leafs["cmplsFrrLogEventType"] = cmplsfrrlogentry.Cmplsfrrlogeventtype
    leafs["cmplsFrrLogEventDuration"] = cmplsfrrlogentry.Cmplsfrrlogeventduration
    leafs["cmplsFrrLogEventReasonString"] = cmplsfrrlogentry.Cmplsfrrlogeventreasonstring
    return leafs
}

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetYangName() string { return "cmplsFrrLogEntry" }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) SetParent(parent types.Entity) { cmplsfrrlogentry.parent = parent }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetParent() types.Entity { return cmplsfrrlogentry.parent }

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetParentYangName() string { return "cmplsFrrLogTable" }

// CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry_Cmplsfrrlogeventtype represents occured.
type CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry_Cmplsfrrlogeventtype string

const (
    CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry_Cmplsfrrlogeventtype_other CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry_Cmplsfrrlogeventtype = "other"

    CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry_Cmplsfrrlogeventtype_protected CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry_Cmplsfrrlogeventtype = "protected"
)

// CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable
// The mplsFrrFacRouteDBTable provides information about the 
// fast reroute database.  Each entry belongs to an interface,
// protecting backup tunnel and protected tunnel. MPLS 
// interfaces defined on this node are protected by backup
// tunnels and are indexed by mplsFrrFacRouteProtectedIndex.
// Backup tunnels defined to protect the tunnels traversing an
// interface, and are indexed by 
// mplsFrrFacRouteProtectingTunIndex.  Note that the tunnel 
// instance index is not required, since it is implied to be 0, 
// which indicates the tunnel head interface for the protecting 
// tunnel. The protecting tunnel is defined to exist on the PLR 
// in the FRR specification.  Protected tunnels are the LSPs that 
// traverse the protected link.  These LSPs are uniquely 
// identified by mplsFrrFacRouteProtectedTunIndex,
// mplsFrrFacRouteProtectedTunInstance, 
// mplsFrrFacRouteProtectedTunIngressLSRId, and 
// mplsFrrFacRouteProtectedTunEgressLSRId.
type CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the mplsFrrDBTable represents a single protected LSP, protected
    // by a backup tunnel and defined for a specific protected interface. Note
    // that for brevity, managers should consult the mplsTunnelTable present in
    // the MPLS-TE MIB for additional information about the protecting and
    // protected tunnels, and the ifEntry in the IF-MIB for the protected
    // interface. The type is slice of
    // CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry.
    Cmplsfrrfacroutedbentry []CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry
}

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetFilter() yfilter.YFilter { return cmplsfrrfacroutedbtable.YFilter }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) SetFilter(yf yfilter.YFilter) { cmplsfrrfacroutedbtable.YFilter = yf }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetGoName(yname string) string {
    if yname == "cmplsFrrFacRouteDBEntry" { return "Cmplsfrrfacroutedbentry" }
    return ""
}

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetSegmentPath() string {
    return "cmplsFrrFacRouteDBTable"
}

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsFrrFacRouteDBEntry" {
        for _, c := range cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry {
            if cmplsfrrfacroutedbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry{}
        cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry = append(cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry, child)
        return &cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry[len(cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry)-1]
    }
    return nil
}

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry {
        children[cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry[i].GetSegmentPath()] = &cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry[i]
    }
    return children
}

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetYangName() string { return "cmplsFrrFacRouteDBTable" }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) SetParent(parent types.Entity) { cmplsfrrfacroutedbtable.parent = parent }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetParent() types.Entity { return cmplsfrrfacroutedbtable.parent }

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetParentYangName() string { return "CISCO-IETF-FRR-MIB" }

// CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry
// An entry in the mplsFrrDBTable represents a single protected
// LSP, protected by a backup tunnel and defined for a specific
// protected interface. Note that for brevity, managers should
// consult the mplsTunnelTable present in the MPLS-TE MIB for
// additional information about the protecting and protected
// tunnels, and the ifEntry in the IF-MIB for the protected
// interface.
type CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies the interface configured for
    // FRR protection. The type is interface{} with range: 1..2147483647.
    Cmplsfrrfacrouteprotectedifindex interface{}

    // This attribute is a key. Uniquely identifies the mplsTunnelEntry primary
    // index for the tunnel head interface designated to protect the  interface as
    // specified in the mplsFrrFacRouteIfProtectedIndex (and all of the tunnels
    // using this interface). The type is interface{} with range: 0..65535.
    Cmplsfrrfacrouteprotectingtunindex interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is interface{} with range: 0..65535.
    Cmplsfrrfacrouteprotectedtunindex interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is interface{} with range: 0..4294967295.
    Cmplsfrrfacrouteprotectedtuninstance interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is string with length: 4.
    Cmplsfrrfacrouteprotectedtuningresslsrid interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is string with length: 4.
    Cmplsfrrfacrouteprotectedtunegresslsrid interface{}

    // Specifies the state of the protected tunnel.  active  This tunnel's label
    // has been placed in the          LFIB and is ready to be applied to incoming
    // packets.           ready -  This tunnel's label entry has been created but
    // is          not yet in the LFIB.           partial - This tunnel's label
    // entry as not been fully           created. The type is
    // Cmplsfrrfacrouteprotectedtunstatus.
    Cmplsfrrfacrouteprotectedtunstatus interface{}

    // Specifies the amount of bandwidth in megabytes per second that is actually
    // reserved by the backup tunnel for facility backup. This value is repeated
    // here from the MPLS- TE MIB because the tunnel entry will reveal the
    // bandwidth reserved by the signaling protocol, which is typically 0 for
    // backup tunnels so as to not over-book bandwidth. However, internal
    // reservations are typically made on the PLR, thus this value should be
    // revealed here. The type is interface{} with range: 0..4294967295.
    Cmplsfrrfacrouteprotectingtunresvbw interface{}

    // Indicates type of the resource protection. The type is
    // Cmplsfrrfacrouteprotectingtunprotectiontype.
    Cmplsfrrfacrouteprotectingtunprotectiontype interface{}
}

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetFilter() yfilter.YFilter { return cmplsfrrfacroutedbentry.YFilter }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) SetFilter(yf yfilter.YFilter) { cmplsfrrfacroutedbentry.YFilter = yf }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetGoName(yname string) string {
    if yname == "cmplsFrrFacRouteProtectedIfIndex" { return "Cmplsfrrfacrouteprotectedifindex" }
    if yname == "cmplsFrrFacRouteProtectingTunIndex" { return "Cmplsfrrfacrouteprotectingtunindex" }
    if yname == "cmplsFrrFacRouteProtectedTunIndex" { return "Cmplsfrrfacrouteprotectedtunindex" }
    if yname == "cmplsFrrFacRouteProtectedTunInstance" { return "Cmplsfrrfacrouteprotectedtuninstance" }
    if yname == "cmplsFrrFacRouteProtectedTunIngressLSRId" { return "Cmplsfrrfacrouteprotectedtuningresslsrid" }
    if yname == "cmplsFrrFacRouteProtectedTunEgressLSRId" { return "Cmplsfrrfacrouteprotectedtunegresslsrid" }
    if yname == "cmplsFrrFacRouteProtectedTunStatus" { return "Cmplsfrrfacrouteprotectedtunstatus" }
    if yname == "cmplsFrrFacRouteProtectingTunResvBw" { return "Cmplsfrrfacrouteprotectingtunresvbw" }
    if yname == "cmplsFrrFacRouteProtectingTunProtectionType" { return "Cmplsfrrfacrouteprotectingtunprotectiontype" }
    return ""
}

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetSegmentPath() string {
    return "cmplsFrrFacRouteDBEntry" + "[cmplsFrrFacRouteProtectedIfIndex='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedifindex) + "']" + "[cmplsFrrFacRouteProtectingTunIndex='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunindex) + "']" + "[cmplsFrrFacRouteProtectedTunIndex='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunindex) + "']" + "[cmplsFrrFacRouteProtectedTunInstance='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuninstance) + "']" + "[cmplsFrrFacRouteProtectedTunIngressLSRId='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuningresslsrid) + "']" + "[cmplsFrrFacRouteProtectedTunEgressLSRId='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunegresslsrid) + "']"
}

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsFrrFacRouteProtectedIfIndex"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedifindex
    leafs["cmplsFrrFacRouteProtectingTunIndex"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunindex
    leafs["cmplsFrrFacRouteProtectedTunIndex"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunindex
    leafs["cmplsFrrFacRouteProtectedTunInstance"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuninstance
    leafs["cmplsFrrFacRouteProtectedTunIngressLSRId"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuningresslsrid
    leafs["cmplsFrrFacRouteProtectedTunEgressLSRId"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunegresslsrid
    leafs["cmplsFrrFacRouteProtectedTunStatus"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunstatus
    leafs["cmplsFrrFacRouteProtectingTunResvBw"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunresvbw
    leafs["cmplsFrrFacRouteProtectingTunProtectionType"] = cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunprotectiontype
    return leafs
}

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetYangName() string { return "cmplsFrrFacRouteDBEntry" }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) SetParent(parent types.Entity) { cmplsfrrfacroutedbentry.parent = parent }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetParent() types.Entity { return cmplsfrrfacroutedbentry.parent }

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetParentYangName() string { return "cmplsFrrFacRouteDBTable" }

// CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus represents           created.
type CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus string

const (
    CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus_active CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus = "active"

    CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus_ready CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus = "ready"

    CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus_partial CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectedtunstatus = "partial"
)

// CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectingtunprotectiontype represents Indicates type of the resource protection.
type CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectingtunprotectiontype string

const (
    CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectingtunprotectiontype_linkProtection CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectingtunprotectiontype = "linkProtection"

    CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectingtunprotectiontype_nodeProtection CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry_Cmplsfrrfacrouteprotectingtunprotectiontype = "nodeProtection"
)

