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
    EntityData types.CommonEntityData
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

func (cISCOIETFFRRMIB *CISCOIETFFRRMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFFRRMIB.EntityData.YFilter = cISCOIETFFRRMIB.YFilter
    cISCOIETFFRRMIB.EntityData.YangName = "CISCO-IETF-FRR-MIB"
    cISCOIETFFRRMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFFRRMIB.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cISCOIETFFRRMIB.EntityData.SegmentPath = "CISCO-IETF-FRR-MIB:CISCO-IETF-FRR-MIB"
    cISCOIETFFRRMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFFRRMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFFRRMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFFRRMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFFRRMIB.EntityData.Children["cmplsFrrScalars"] = types.YChild{"Cmplsfrrscalars", &cISCOIETFFRRMIB.Cmplsfrrscalars}
    cISCOIETFFRRMIB.EntityData.Children["cmplsFrrConstTable"] = types.YChild{"Cmplsfrrconsttable", &cISCOIETFFRRMIB.Cmplsfrrconsttable}
    cISCOIETFFRRMIB.EntityData.Children["cmplsFrrLogTable"] = types.YChild{"Cmplsfrrlogtable", &cISCOIETFFRRMIB.Cmplsfrrlogtable}
    cISCOIETFFRRMIB.EntityData.Children["cmplsFrrFacRouteDBTable"] = types.YChild{"Cmplsfrrfacroutedbtable", &cISCOIETFFRRMIB.Cmplsfrrfacroutedbtable}
    cISCOIETFFRRMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFFRRMIB.EntityData)
}

// CISCOIETFFRRMIB_Cmplsfrrscalars
type CISCOIETFFRRMIB_Cmplsfrrscalars struct {
    EntityData types.CommonEntityData
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

func (cmplsfrrscalars *CISCOIETFFRRMIB_Cmplsfrrscalars) GetEntityData() *types.CommonEntityData {
    cmplsfrrscalars.EntityData.YFilter = cmplsfrrscalars.YFilter
    cmplsfrrscalars.EntityData.YangName = "cmplsFrrScalars"
    cmplsfrrscalars.EntityData.BundleName = "cisco_ios_xe"
    cmplsfrrscalars.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsfrrscalars.EntityData.SegmentPath = "cmplsFrrScalars"
    cmplsfrrscalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsfrrscalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsfrrscalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsfrrscalars.EntityData.Children = make(map[string]types.YChild)
    cmplsfrrscalars.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrDetourIncoming"] = types.YLeaf{"Cmplsfrrdetourincoming", cmplsfrrscalars.Cmplsfrrdetourincoming}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrDetourOutgoing"] = types.YLeaf{"Cmplsfrrdetouroutgoing", cmplsfrrscalars.Cmplsfrrdetouroutgoing}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrDetourOriginating"] = types.YLeaf{"Cmplsfrrdetouroriginating", cmplsfrrscalars.Cmplsfrrdetouroriginating}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrSwitchover"] = types.YLeaf{"Cmplsfrrswitchover", cmplsfrrscalars.Cmplsfrrswitchover}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrNumOfConfIfs"] = types.YLeaf{"Cmplsfrrnumofconfifs", cmplsfrrscalars.Cmplsfrrnumofconfifs}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrActProtectedIfs"] = types.YLeaf{"Cmplsfrractprotectedifs", cmplsfrrscalars.Cmplsfrractprotectedifs}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrConfProtectingTuns"] = types.YLeaf{"Cmplsfrrconfprotectingtuns", cmplsfrrscalars.Cmplsfrrconfprotectingtuns}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrActProtectedTuns"] = types.YLeaf{"Cmplsfrractprotectedtuns", cmplsfrrscalars.Cmplsfrractprotectedtuns}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrActProtectedLSPs"] = types.YLeaf{"Cmplsfrractprotectedlsps", cmplsfrrscalars.Cmplsfrractprotectedlsps}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrConstProtectionMethod"] = types.YLeaf{"Cmplsfrrconstprotectionmethod", cmplsfrrscalars.Cmplsfrrconstprotectionmethod}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrNotifsEnabled"] = types.YLeaf{"Cmplsfrrnotifsenabled", cmplsfrrscalars.Cmplsfrrnotifsenabled}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrLogTableMaxEntries"] = types.YLeaf{"Cmplsfrrlogtablemaxentries", cmplsfrrscalars.Cmplsfrrlogtablemaxentries}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrLogTableCurrEntries"] = types.YLeaf{"Cmplsfrrlogtablecurrentries", cmplsfrrscalars.Cmplsfrrlogtablecurrentries}
    cmplsfrrscalars.EntityData.Leafs["cmplsFrrNotifMaxRate"] = types.YLeaf{"Cmplsfrrnotifmaxrate", cmplsfrrscalars.Cmplsfrrnotifmaxrate}
    return &(cmplsfrrscalars.EntityData)
}

// CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod represents for more details.
type CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod string

const (
    CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod_oneToOneBackup CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod = "oneToOneBackup"

    CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod_facilityBackup CISCOIETFFRRMIB_Cmplsfrrscalars_Cmplsfrrconstprotectionmethod = "facilityBackup"
)

// CISCOIETFFRRMIB_Cmplsfrrconsttable
// This table shows detour setup constraints.
type CISCOIETFFRRMIB_Cmplsfrrconsttable struct {
    EntityData types.CommonEntityData
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

func (cmplsfrrconsttable *CISCOIETFFRRMIB_Cmplsfrrconsttable) GetEntityData() *types.CommonEntityData {
    cmplsfrrconsttable.EntityData.YFilter = cmplsfrrconsttable.YFilter
    cmplsfrrconsttable.EntityData.YangName = "cmplsFrrConstTable"
    cmplsfrrconsttable.EntityData.BundleName = "cisco_ios_xe"
    cmplsfrrconsttable.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsfrrconsttable.EntityData.SegmentPath = "cmplsFrrConstTable"
    cmplsfrrconsttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsfrrconsttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsfrrconsttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsfrrconsttable.EntityData.Children = make(map[string]types.YChild)
    cmplsfrrconsttable.EntityData.Children["cmplsFrrConstEntry"] = types.YChild{"Cmplsfrrconstentry", nil}
    for i := range cmplsfrrconsttable.Cmplsfrrconstentry {
        cmplsfrrconsttable.EntityData.Children[types.GetSegmentPath(&cmplsfrrconsttable.Cmplsfrrconstentry[i])] = types.YChild{"Cmplsfrrconstentry", &cmplsfrrconsttable.Cmplsfrrconstentry[i]}
    }
    cmplsfrrconsttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplsfrrconsttable.EntityData)
}

// CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry
// An entry in this table represents detour LSP or bypass tunnel 
// setup constraints for a tunnel instance to be protected by 
// detour LSPs or a tunnel. Agents must allow entries in this table 
// to be created only for tunnel instances that require fast-reroute.
// Entries indexed with mplsFrrConstIfIndex set to 0 apply to all
// interfaces on this device for which the FRR feature can operate
// on.
type CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry struct {
    EntityData types.CommonEntityData
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

func (cmplsfrrconstentry *CISCOIETFFRRMIB_Cmplsfrrconsttable_Cmplsfrrconstentry) GetEntityData() *types.CommonEntityData {
    cmplsfrrconstentry.EntityData.YFilter = cmplsfrrconstentry.YFilter
    cmplsfrrconstentry.EntityData.YangName = "cmplsFrrConstEntry"
    cmplsfrrconstentry.EntityData.BundleName = "cisco_ios_xe"
    cmplsfrrconstentry.EntityData.ParentYangName = "cmplsFrrConstTable"
    cmplsfrrconstentry.EntityData.SegmentPath = "cmplsFrrConstEntry" + "[cmplsFrrConstIfIndex='" + fmt.Sprintf("%v", cmplsfrrconstentry.Cmplsfrrconstifindex) + "']" + "[cmplsFrrConstTunnelIndex='" + fmt.Sprintf("%v", cmplsfrrconstentry.Cmplsfrrconsttunnelindex) + "']" + "[cmplsFrrConstTunnelInstance='" + fmt.Sprintf("%v", cmplsfrrconstentry.Cmplsfrrconsttunnelinstance) + "']"
    cmplsfrrconstentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsfrrconstentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsfrrconstentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsfrrconstentry.EntityData.Children = make(map[string]types.YChild)
    cmplsfrrconstentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstIfIndex"] = types.YLeaf{"Cmplsfrrconstifindex", cmplsfrrconstentry.Cmplsfrrconstifindex}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstTunnelIndex"] = types.YLeaf{"Cmplsfrrconsttunnelindex", cmplsfrrconstentry.Cmplsfrrconsttunnelindex}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstTunnelInstance"] = types.YLeaf{"Cmplsfrrconsttunnelinstance", cmplsfrrconstentry.Cmplsfrrconsttunnelinstance}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstSetupPrio"] = types.YLeaf{"Cmplsfrrconstsetupprio", cmplsfrrconstentry.Cmplsfrrconstsetupprio}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstHoldingPrio"] = types.YLeaf{"Cmplsfrrconstholdingprio", cmplsfrrconstentry.Cmplsfrrconstholdingprio}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstInclAnyAffinity"] = types.YLeaf{"Cmplsfrrconstinclanyaffinity", cmplsfrrconstentry.Cmplsfrrconstinclanyaffinity}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstInclAllAffinity"] = types.YLeaf{"Cmplsfrrconstinclallaffinity", cmplsfrrconstentry.Cmplsfrrconstinclallaffinity}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstExclAllAffinity"] = types.YLeaf{"Cmplsfrrconstexclallaffinity", cmplsfrrconstentry.Cmplsfrrconstexclallaffinity}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstHopLimit"] = types.YLeaf{"Cmplsfrrconsthoplimit", cmplsfrrconstentry.Cmplsfrrconsthoplimit}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstBandwidth"] = types.YLeaf{"Cmplsfrrconstbandwidth", cmplsfrrconstentry.Cmplsfrrconstbandwidth}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstRowStatus"] = types.YLeaf{"Cmplsfrrconstrowstatus", cmplsfrrconstentry.Cmplsfrrconstrowstatus}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstNumProtectingTunOnIf"] = types.YLeaf{"Cmplsfrrconstnumprotectingtunonif", cmplsfrrconstentry.Cmplsfrrconstnumprotectingtunonif}
    cmplsfrrconstentry.EntityData.Leafs["cmplsFrrConstNumProtectedTunOnIf"] = types.YLeaf{"Cmplsfrrconstnumprotectedtunonif", cmplsfrrconstentry.Cmplsfrrconstnumprotectedtunonif}
    return &(cmplsfrrconstentry.EntityData)
}

// CISCOIETFFRRMIB_Cmplsfrrlogtable
// The fast reroute log table records fast reroute events such
// as protected links going up or down or the FRR feature
// kicking in.
type CISCOIETFFRRMIB_Cmplsfrrlogtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created to describe one fast reroute event. 
    // Entries in this table are only created and destroyed by the agent
    // implementation. The maximum number  of entries in this log is governed by
    // the scalar. The type is slice of
    // CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry.
    Cmplsfrrlogentry []CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry
}

func (cmplsfrrlogtable *CISCOIETFFRRMIB_Cmplsfrrlogtable) GetEntityData() *types.CommonEntityData {
    cmplsfrrlogtable.EntityData.YFilter = cmplsfrrlogtable.YFilter
    cmplsfrrlogtable.EntityData.YangName = "cmplsFrrLogTable"
    cmplsfrrlogtable.EntityData.BundleName = "cisco_ios_xe"
    cmplsfrrlogtable.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsfrrlogtable.EntityData.SegmentPath = "cmplsFrrLogTable"
    cmplsfrrlogtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsfrrlogtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsfrrlogtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsfrrlogtable.EntityData.Children = make(map[string]types.YChild)
    cmplsfrrlogtable.EntityData.Children["cmplsFrrLogEntry"] = types.YChild{"Cmplsfrrlogentry", nil}
    for i := range cmplsfrrlogtable.Cmplsfrrlogentry {
        cmplsfrrlogtable.EntityData.Children[types.GetSegmentPath(&cmplsfrrlogtable.Cmplsfrrlogentry[i])] = types.YChild{"Cmplsfrrlogentry", &cmplsfrrlogtable.Cmplsfrrlogentry[i]}
    }
    cmplsfrrlogtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplsfrrlogtable.EntityData)
}

// CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry
// An entry in this table is created to describe one fast
// reroute event.  Entries in this table are only created and
// destroyed by the agent implementation. The maximum number 
// of entries in this log is governed by the scalar.
type CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry struct {
    EntityData types.CommonEntityData
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

func (cmplsfrrlogentry *CISCOIETFFRRMIB_Cmplsfrrlogtable_Cmplsfrrlogentry) GetEntityData() *types.CommonEntityData {
    cmplsfrrlogentry.EntityData.YFilter = cmplsfrrlogentry.YFilter
    cmplsfrrlogentry.EntityData.YangName = "cmplsFrrLogEntry"
    cmplsfrrlogentry.EntityData.BundleName = "cisco_ios_xe"
    cmplsfrrlogentry.EntityData.ParentYangName = "cmplsFrrLogTable"
    cmplsfrrlogentry.EntityData.SegmentPath = "cmplsFrrLogEntry" + "[cmplsFrrLogIndex='" + fmt.Sprintf("%v", cmplsfrrlogentry.Cmplsfrrlogindex) + "']"
    cmplsfrrlogentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsfrrlogentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsfrrlogentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsfrrlogentry.EntityData.Children = make(map[string]types.YChild)
    cmplsfrrlogentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsfrrlogentry.EntityData.Leafs["cmplsFrrLogIndex"] = types.YLeaf{"Cmplsfrrlogindex", cmplsfrrlogentry.Cmplsfrrlogindex}
    cmplsfrrlogentry.EntityData.Leafs["cmplsFrrLogEventTime"] = types.YLeaf{"Cmplsfrrlogeventtime", cmplsfrrlogentry.Cmplsfrrlogeventtime}
    cmplsfrrlogentry.EntityData.Leafs["cmplsFrrLogInterface"] = types.YLeaf{"Cmplsfrrloginterface", cmplsfrrlogentry.Cmplsfrrloginterface}
    cmplsfrrlogentry.EntityData.Leafs["cmplsFrrLogEventType"] = types.YLeaf{"Cmplsfrrlogeventtype", cmplsfrrlogentry.Cmplsfrrlogeventtype}
    cmplsfrrlogentry.EntityData.Leafs["cmplsFrrLogEventDuration"] = types.YLeaf{"Cmplsfrrlogeventduration", cmplsfrrlogentry.Cmplsfrrlogeventduration}
    cmplsfrrlogentry.EntityData.Leafs["cmplsFrrLogEventReasonString"] = types.YLeaf{"Cmplsfrrlogeventreasonstring", cmplsfrrlogentry.Cmplsfrrlogeventreasonstring}
    return &(cmplsfrrlogentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cmplsfrrfacroutedbtable *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable) GetEntityData() *types.CommonEntityData {
    cmplsfrrfacroutedbtable.EntityData.YFilter = cmplsfrrfacroutedbtable.YFilter
    cmplsfrrfacroutedbtable.EntityData.YangName = "cmplsFrrFacRouteDBTable"
    cmplsfrrfacroutedbtable.EntityData.BundleName = "cisco_ios_xe"
    cmplsfrrfacroutedbtable.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsfrrfacroutedbtable.EntityData.SegmentPath = "cmplsFrrFacRouteDBTable"
    cmplsfrrfacroutedbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsfrrfacroutedbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsfrrfacroutedbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsfrrfacroutedbtable.EntityData.Children = make(map[string]types.YChild)
    cmplsfrrfacroutedbtable.EntityData.Children["cmplsFrrFacRouteDBEntry"] = types.YChild{"Cmplsfrrfacroutedbentry", nil}
    for i := range cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry {
        cmplsfrrfacroutedbtable.EntityData.Children[types.GetSegmentPath(&cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry[i])] = types.YChild{"Cmplsfrrfacroutedbentry", &cmplsfrrfacroutedbtable.Cmplsfrrfacroutedbentry[i]}
    }
    cmplsfrrfacroutedbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplsfrrfacroutedbtable.EntityData)
}

// CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry
// An entry in the mplsFrrDBTable represents a single protected
// LSP, protected by a backup tunnel and defined for a specific
// protected interface. Note that for brevity, managers should
// consult the mplsTunnelTable present in the MPLS-TE MIB for
// additional information about the protecting and protected
// tunnels, and the ifEntry in the IF-MIB for the protected
// interface.
type CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry struct {
    EntityData types.CommonEntityData
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

func (cmplsfrrfacroutedbentry *CISCOIETFFRRMIB_Cmplsfrrfacroutedbtable_Cmplsfrrfacroutedbentry) GetEntityData() *types.CommonEntityData {
    cmplsfrrfacroutedbentry.EntityData.YFilter = cmplsfrrfacroutedbentry.YFilter
    cmplsfrrfacroutedbentry.EntityData.YangName = "cmplsFrrFacRouteDBEntry"
    cmplsfrrfacroutedbentry.EntityData.BundleName = "cisco_ios_xe"
    cmplsfrrfacroutedbentry.EntityData.ParentYangName = "cmplsFrrFacRouteDBTable"
    cmplsfrrfacroutedbentry.EntityData.SegmentPath = "cmplsFrrFacRouteDBEntry" + "[cmplsFrrFacRouteProtectedIfIndex='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedifindex) + "']" + "[cmplsFrrFacRouteProtectingTunIndex='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunindex) + "']" + "[cmplsFrrFacRouteProtectedTunIndex='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunindex) + "']" + "[cmplsFrrFacRouteProtectedTunInstance='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuninstance) + "']" + "[cmplsFrrFacRouteProtectedTunIngressLSRId='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuningresslsrid) + "']" + "[cmplsFrrFacRouteProtectedTunEgressLSRId='" + fmt.Sprintf("%v", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunegresslsrid) + "']"
    cmplsfrrfacroutedbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsfrrfacroutedbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsfrrfacroutedbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsfrrfacroutedbentry.EntityData.Children = make(map[string]types.YChild)
    cmplsfrrfacroutedbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectedIfIndex"] = types.YLeaf{"Cmplsfrrfacrouteprotectedifindex", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedifindex}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectingTunIndex"] = types.YLeaf{"Cmplsfrrfacrouteprotectingtunindex", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunindex}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectedTunIndex"] = types.YLeaf{"Cmplsfrrfacrouteprotectedtunindex", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunindex}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectedTunInstance"] = types.YLeaf{"Cmplsfrrfacrouteprotectedtuninstance", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuninstance}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectedTunIngressLSRId"] = types.YLeaf{"Cmplsfrrfacrouteprotectedtuningresslsrid", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtuningresslsrid}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectedTunEgressLSRId"] = types.YLeaf{"Cmplsfrrfacrouteprotectedtunegresslsrid", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunegresslsrid}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectedTunStatus"] = types.YLeaf{"Cmplsfrrfacrouteprotectedtunstatus", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectedtunstatus}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectingTunResvBw"] = types.YLeaf{"Cmplsfrrfacrouteprotectingtunresvbw", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunresvbw}
    cmplsfrrfacroutedbentry.EntityData.Leafs["cmplsFrrFacRouteProtectingTunProtectionType"] = types.YLeaf{"Cmplsfrrfacrouteprotectingtunprotectiontype", cmplsfrrfacroutedbentry.Cmplsfrrfacrouteprotectingtunprotectiontype}
    return &(cmplsfrrfacroutedbentry.EntityData)
}

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

