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

    
    CmplsFrrScalars CISCOIETFFRRMIB_CmplsFrrScalars

    // This table shows detour setup constraints.
    CmplsFrrConstTable CISCOIETFFRRMIB_CmplsFrrConstTable

    // The fast reroute log table records fast reroute events such as protected
    // links going up or down or the FRR feature kicking in.
    CmplsFrrLogTable CISCOIETFFRRMIB_CmplsFrrLogTable

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
    CmplsFrrFacRouteDBTable CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable
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

    cISCOIETFFRRMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFFRRMIB.EntityData.Children.Append("cmplsFrrScalars", types.YChild{"CmplsFrrScalars", &cISCOIETFFRRMIB.CmplsFrrScalars})
    cISCOIETFFRRMIB.EntityData.Children.Append("cmplsFrrConstTable", types.YChild{"CmplsFrrConstTable", &cISCOIETFFRRMIB.CmplsFrrConstTable})
    cISCOIETFFRRMIB.EntityData.Children.Append("cmplsFrrLogTable", types.YChild{"CmplsFrrLogTable", &cISCOIETFFRRMIB.CmplsFrrLogTable})
    cISCOIETFFRRMIB.EntityData.Children.Append("cmplsFrrFacRouteDBTable", types.YChild{"CmplsFrrFacRouteDBTable", &cISCOIETFFRRMIB.CmplsFrrFacRouteDBTable})
    cISCOIETFFRRMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFFRRMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFFRRMIB.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrScalars
type CISCOIETFFRRMIB_CmplsFrrScalars struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of detour LSPs entering the device if
    // mplsFrrConstProtectionMethod is set to oneToOneBackup(0), or or 0 if
    // mplsFrrConstProtectionMethod is set to facilityBackup(1). The type is
    // interface{} with range: 0..4294967295.
    CmplsFrrDetourIncoming interface{}

    // The number of detour LSPs leaving the device if
    // mplsFrrConstProtectionMethod is set to oneToOneBackup(0), or 0 if
    // mplsFrrConstProtectionMethod is set to  to facilityBackup(1). The type is
    // interface{} with range: 0..4294967295.
    CmplsFrrDetourOutgoing interface{}

    // The number of detour LSPs originating at this PLR if
    // mplsFrrConstProtectionMethod is set to oneToOneBackup(0). This object MUST
    // return 0 if the mplsFrrConstProtectionMethod  is set to facilityBackup(1).
    // The type is interface{} with range: 0..4294967295.
    CmplsFrrDetourOriginating interface{}

    // The number of tunnel instances that are switched over to their
    // corresponding detour LSP if mplsFrrConstProtectionMethod is set to
    // oneToOneBackup(0), or tunnels being switched over if
    // mplsFrrConstProtectionMethod is set to facilityBackup(1). The type is
    // interface{} with range: 0..4294967295.
    CmplsFrrSwitchover interface{}

    // Indicates the number of MPLS interfaces configured for  protection by the
    // FRR feature, otherwise this value MUST return 0 to indicate that LSPs
    // traversing any  interface may be protected. The type is interface{} with
    // range: 0..4294967295.
    CmplsFrrNumOfConfIfs interface{}

    // Indicates the number of interfaces currently being protected  by the FRR
    // feature if mplsFrrConstProtectionMethod is set to facilityBackup(1),
    // otherwise this value should return 0 to indicate that LSPs traversing any
    // interface may be protected. This value MUST be less than or equal to
    // mplsFrrConfIfs. The type is interface{} with range: 0..4294967295.
    CmplsFrrActProtectedIfs interface{}

    // Indicates the number of bypass tunnels configured to  protect facilities on
    // this LSR using the FRR feature  if mplsFrrConstProtectionMethod is set to 
    // facilityBackup(1), otherwise this value MUST return  0. The type is
    // interface{} with range: 0..4294967295.
    CmplsFrrConfProtectingTuns interface{}

    // Indicates the number of bypass tunnels indicated in
    // mplsFrrConfProtectingTuns whose operStatus is up(1) indicating that they
    // are currently protecting facilities on this LSR using the FRR feature. This
    // object MUST return 0 if mplsFrrConstProtectionMethod  is set to
    // facilityBackup(1). The type is interface{} with range: 0..4294967295.
    CmplsFrrActProtectedTuns interface{}

    // Indicates the number of LSPs currently protected by  the FRR feature. If
    // mplsFrrConstProtectionMethod is set  to facilityBackup(1)this object MUST
    // return 0. The type is interface{} with range: 0..4294967295.
    CmplsFrrActProtectedLSPs interface{}

    // Indicates which protection method is to be used for fast reroute. Some
    // devices may require a reboot of their routing processors if this variable
    // is changed. An agent which does not wish to reboot or modify its FRR mode 
    // MUST return an inconsistentValue error. Please  consult the device's agent
    // capability statement  for more details. The type is
    // CmplsFrrConstProtectionMethod.
    CmplsFrrConstProtectionMethod interface{}

    // Enables or disables FRR notifications defined in this MIB module.
    // Notifications are disabled by default. The type is bool.
    CmplsFrrNotifsEnabled interface{}

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
    CmplsFrrLogTableMaxEntries interface{}

    // Indicates the current number of entries in the FRR log table. The type is
    // interface{} with range: 0..4294967295.
    CmplsFrrLogTableCurrEntries interface{}

    // This variable indicates the number of milliseconds that must elapse between
    // notification emissions. If events occur more rapidly, the implementation
    // may simply fail to emit these notifications during that period, or may
    // queue them until an appropriate time in the future. A value of 0 means no
    // minimum  elapsed period is specified. The type is interface{} with range:
    // 0..4294967295.
    CmplsFrrNotifMaxRate interface{}
}

func (cmplsFrrScalars *CISCOIETFFRRMIB_CmplsFrrScalars) GetEntityData() *types.CommonEntityData {
    cmplsFrrScalars.EntityData.YFilter = cmplsFrrScalars.YFilter
    cmplsFrrScalars.EntityData.YangName = "cmplsFrrScalars"
    cmplsFrrScalars.EntityData.BundleName = "cisco_ios_xe"
    cmplsFrrScalars.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsFrrScalars.EntityData.SegmentPath = "cmplsFrrScalars"
    cmplsFrrScalars.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsFrrScalars.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsFrrScalars.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsFrrScalars.EntityData.Children = types.NewOrderedMap()
    cmplsFrrScalars.EntityData.Leafs = types.NewOrderedMap()
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrDetourIncoming", types.YLeaf{"CmplsFrrDetourIncoming", cmplsFrrScalars.CmplsFrrDetourIncoming})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrDetourOutgoing", types.YLeaf{"CmplsFrrDetourOutgoing", cmplsFrrScalars.CmplsFrrDetourOutgoing})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrDetourOriginating", types.YLeaf{"CmplsFrrDetourOriginating", cmplsFrrScalars.CmplsFrrDetourOriginating})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrSwitchover", types.YLeaf{"CmplsFrrSwitchover", cmplsFrrScalars.CmplsFrrSwitchover})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrNumOfConfIfs", types.YLeaf{"CmplsFrrNumOfConfIfs", cmplsFrrScalars.CmplsFrrNumOfConfIfs})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrActProtectedIfs", types.YLeaf{"CmplsFrrActProtectedIfs", cmplsFrrScalars.CmplsFrrActProtectedIfs})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrConfProtectingTuns", types.YLeaf{"CmplsFrrConfProtectingTuns", cmplsFrrScalars.CmplsFrrConfProtectingTuns})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrActProtectedTuns", types.YLeaf{"CmplsFrrActProtectedTuns", cmplsFrrScalars.CmplsFrrActProtectedTuns})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrActProtectedLSPs", types.YLeaf{"CmplsFrrActProtectedLSPs", cmplsFrrScalars.CmplsFrrActProtectedLSPs})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrConstProtectionMethod", types.YLeaf{"CmplsFrrConstProtectionMethod", cmplsFrrScalars.CmplsFrrConstProtectionMethod})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrNotifsEnabled", types.YLeaf{"CmplsFrrNotifsEnabled", cmplsFrrScalars.CmplsFrrNotifsEnabled})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrLogTableMaxEntries", types.YLeaf{"CmplsFrrLogTableMaxEntries", cmplsFrrScalars.CmplsFrrLogTableMaxEntries})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrLogTableCurrEntries", types.YLeaf{"CmplsFrrLogTableCurrEntries", cmplsFrrScalars.CmplsFrrLogTableCurrEntries})
    cmplsFrrScalars.EntityData.Leafs.Append("cmplsFrrNotifMaxRate", types.YLeaf{"CmplsFrrNotifMaxRate", cmplsFrrScalars.CmplsFrrNotifMaxRate})

    cmplsFrrScalars.EntityData.YListKeys = []string {}

    return &(cmplsFrrScalars.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrScalars_CmplsFrrConstProtectionMethod represents for more details.
type CISCOIETFFRRMIB_CmplsFrrScalars_CmplsFrrConstProtectionMethod string

const (
    CISCOIETFFRRMIB_CmplsFrrScalars_CmplsFrrConstProtectionMethod_oneToOneBackup CISCOIETFFRRMIB_CmplsFrrScalars_CmplsFrrConstProtectionMethod = "oneToOneBackup"

    CISCOIETFFRRMIB_CmplsFrrScalars_CmplsFrrConstProtectionMethod_facilityBackup CISCOIETFFRRMIB_CmplsFrrScalars_CmplsFrrConstProtectionMethod = "facilityBackup"
)

// CISCOIETFFRRMIB_CmplsFrrConstTable
// This table shows detour setup constraints.
type CISCOIETFFRRMIB_CmplsFrrConstTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents detour LSP or bypass tunnel  setup
    // constraints for a tunnel instance to be protected by  detour LSPs or a
    // tunnel. Agents must allow entries in this table  to be created only for
    // tunnel instances that require fast-reroute. Entries indexed with
    // mplsFrrConstIfIndex set to 0 apply to all interfaces on this device for
    // which the FRR feature can operate on. The type is slice of
    // CISCOIETFFRRMIB_CmplsFrrConstTable_CmplsFrrConstEntry.
    CmplsFrrConstEntry []*CISCOIETFFRRMIB_CmplsFrrConstTable_CmplsFrrConstEntry
}

func (cmplsFrrConstTable *CISCOIETFFRRMIB_CmplsFrrConstTable) GetEntityData() *types.CommonEntityData {
    cmplsFrrConstTable.EntityData.YFilter = cmplsFrrConstTable.YFilter
    cmplsFrrConstTable.EntityData.YangName = "cmplsFrrConstTable"
    cmplsFrrConstTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsFrrConstTable.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsFrrConstTable.EntityData.SegmentPath = "cmplsFrrConstTable"
    cmplsFrrConstTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsFrrConstTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsFrrConstTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsFrrConstTable.EntityData.Children = types.NewOrderedMap()
    cmplsFrrConstTable.EntityData.Children.Append("cmplsFrrConstEntry", types.YChild{"CmplsFrrConstEntry", nil})
    for i := range cmplsFrrConstTable.CmplsFrrConstEntry {
        cmplsFrrConstTable.EntityData.Children.Append(types.GetSegmentPath(cmplsFrrConstTable.CmplsFrrConstEntry[i]), types.YChild{"CmplsFrrConstEntry", cmplsFrrConstTable.CmplsFrrConstEntry[i]})
    }
    cmplsFrrConstTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsFrrConstTable.EntityData.YListKeys = []string {}

    return &(cmplsFrrConstTable.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrConstTable_CmplsFrrConstEntry
// An entry in this table represents detour LSP or bypass tunnel 
// setup constraints for a tunnel instance to be protected by 
// detour LSPs or a tunnel. Agents must allow entries in this table 
// to be created only for tunnel instances that require fast-reroute.
// Entries indexed with mplsFrrConstIfIndex set to 0 apply to all
// interfaces on this device for which the FRR feature can operate
// on.
type CISCOIETFFRRMIB_CmplsFrrConstTable_CmplsFrrConstEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies an interface for which fast
    // reroute is configured. Tabular entries indexed with a 0 value apply to all
    // interfaces on this device for which the FRR feature can operate on. The
    // type is interface{} with range: 0..2147483647.
    CmplsFrrConstIfIndex interface{}

    // This attribute is a key. Uniquely identifies a tunnel for which fast
    // reroute is requested. The type is interface{} with range: 0..65535.
    CmplsFrrConstTunnelIndex interface{}

    // This attribute is a key. Uniquely identifies an instance of this tunnel for
    // which fast reroute is requested. The type is interface{} with range:
    // 0..4294967295.
    CmplsFrrConstTunnelInstance interface{}

    // Indicates the setup priority of detour LSP. The type is interface{} with
    // range: 0..7.
    CmplsFrrConstSetupPrio interface{}

    // Indicates the holding priority for detour LSP. The type is interface{} with
    // range: 0..7.
    CmplsFrrConstHoldingPrio interface{}

    // A link satisfies the include-any constraint if and only if the constraint
    // is zero, or the link and the constraint have a resource class in common.
    // The type is interface{} with range: 0..4294967295.
    CmplsFrrConstInclAnyAffinity interface{}

    // A link satisfies the include-all constraint if and only if the link
    // contains all of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    CmplsFrrConstInclAllAffinity interface{}

    // A link satisfies the exclude-all constraint if and only if the link
    // contains none of the administrative groups specified in the constraint. The
    // type is interface{} with range: 0..4294967295.
    CmplsFrrConstExclAllAffinity interface{}

    // The maximum number of hops that the detour LSP may traverse. The type is
    // interface{} with range: 1..65535.
    CmplsFrrConstHopLimit interface{}

    // This variable represents the bandwidth for detour LSPs of this tunnel, in
    // units of thousands of bits per second (Kbps). The type is interface{} with
    // range: 0..4294967295.
    CmplsFrrConstBandwidth interface{}

    // This object is used to create, modify, and/or delete a row in this table.
    // The type is RowStatus.
    CmplsFrrConstRowStatus interface{}

    // The number of backup tunnels protecting the specified interface. The type
    // is interface{} with range: 0..4294967295.
    CmplsFrrConstNumProtectingTunOnIf interface{}

    // The number of tunnels protected on this interface. The type is interface{}
    // with range: 0..4294967295.
    CmplsFrrConstNumProtectedTunOnIf interface{}
}

func (cmplsFrrConstEntry *CISCOIETFFRRMIB_CmplsFrrConstTable_CmplsFrrConstEntry) GetEntityData() *types.CommonEntityData {
    cmplsFrrConstEntry.EntityData.YFilter = cmplsFrrConstEntry.YFilter
    cmplsFrrConstEntry.EntityData.YangName = "cmplsFrrConstEntry"
    cmplsFrrConstEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsFrrConstEntry.EntityData.ParentYangName = "cmplsFrrConstTable"
    cmplsFrrConstEntry.EntityData.SegmentPath = "cmplsFrrConstEntry" + types.AddKeyToken(cmplsFrrConstEntry.CmplsFrrConstIfIndex, "cmplsFrrConstIfIndex") + types.AddKeyToken(cmplsFrrConstEntry.CmplsFrrConstTunnelIndex, "cmplsFrrConstTunnelIndex") + types.AddKeyToken(cmplsFrrConstEntry.CmplsFrrConstTunnelInstance, "cmplsFrrConstTunnelInstance")
    cmplsFrrConstEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsFrrConstEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsFrrConstEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsFrrConstEntry.EntityData.Children = types.NewOrderedMap()
    cmplsFrrConstEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstIfIndex", types.YLeaf{"CmplsFrrConstIfIndex", cmplsFrrConstEntry.CmplsFrrConstIfIndex})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstTunnelIndex", types.YLeaf{"CmplsFrrConstTunnelIndex", cmplsFrrConstEntry.CmplsFrrConstTunnelIndex})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstTunnelInstance", types.YLeaf{"CmplsFrrConstTunnelInstance", cmplsFrrConstEntry.CmplsFrrConstTunnelInstance})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstSetupPrio", types.YLeaf{"CmplsFrrConstSetupPrio", cmplsFrrConstEntry.CmplsFrrConstSetupPrio})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstHoldingPrio", types.YLeaf{"CmplsFrrConstHoldingPrio", cmplsFrrConstEntry.CmplsFrrConstHoldingPrio})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstInclAnyAffinity", types.YLeaf{"CmplsFrrConstInclAnyAffinity", cmplsFrrConstEntry.CmplsFrrConstInclAnyAffinity})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstInclAllAffinity", types.YLeaf{"CmplsFrrConstInclAllAffinity", cmplsFrrConstEntry.CmplsFrrConstInclAllAffinity})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstExclAllAffinity", types.YLeaf{"CmplsFrrConstExclAllAffinity", cmplsFrrConstEntry.CmplsFrrConstExclAllAffinity})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstHopLimit", types.YLeaf{"CmplsFrrConstHopLimit", cmplsFrrConstEntry.CmplsFrrConstHopLimit})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstBandwidth", types.YLeaf{"CmplsFrrConstBandwidth", cmplsFrrConstEntry.CmplsFrrConstBandwidth})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstRowStatus", types.YLeaf{"CmplsFrrConstRowStatus", cmplsFrrConstEntry.CmplsFrrConstRowStatus})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstNumProtectingTunOnIf", types.YLeaf{"CmplsFrrConstNumProtectingTunOnIf", cmplsFrrConstEntry.CmplsFrrConstNumProtectingTunOnIf})
    cmplsFrrConstEntry.EntityData.Leafs.Append("cmplsFrrConstNumProtectedTunOnIf", types.YLeaf{"CmplsFrrConstNumProtectedTunOnIf", cmplsFrrConstEntry.CmplsFrrConstNumProtectedTunOnIf})

    cmplsFrrConstEntry.EntityData.YListKeys = []string {"CmplsFrrConstIfIndex", "CmplsFrrConstTunnelIndex", "CmplsFrrConstTunnelInstance"}

    return &(cmplsFrrConstEntry.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrLogTable
// The fast reroute log table records fast reroute events such
// as protected links going up or down or the FRR feature
// kicking in.
type CISCOIETFFRRMIB_CmplsFrrLogTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created to describe one fast reroute event. 
    // Entries in this table are only created and destroyed by the agent
    // implementation. The maximum number  of entries in this log is governed by
    // the scalar. The type is slice of
    // CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry.
    CmplsFrrLogEntry []*CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry
}

func (cmplsFrrLogTable *CISCOIETFFRRMIB_CmplsFrrLogTable) GetEntityData() *types.CommonEntityData {
    cmplsFrrLogTable.EntityData.YFilter = cmplsFrrLogTable.YFilter
    cmplsFrrLogTable.EntityData.YangName = "cmplsFrrLogTable"
    cmplsFrrLogTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsFrrLogTable.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsFrrLogTable.EntityData.SegmentPath = "cmplsFrrLogTable"
    cmplsFrrLogTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsFrrLogTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsFrrLogTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsFrrLogTable.EntityData.Children = types.NewOrderedMap()
    cmplsFrrLogTable.EntityData.Children.Append("cmplsFrrLogEntry", types.YChild{"CmplsFrrLogEntry", nil})
    for i := range cmplsFrrLogTable.CmplsFrrLogEntry {
        cmplsFrrLogTable.EntityData.Children.Append(types.GetSegmentPath(cmplsFrrLogTable.CmplsFrrLogEntry[i]), types.YChild{"CmplsFrrLogEntry", cmplsFrrLogTable.CmplsFrrLogEntry[i]})
    }
    cmplsFrrLogTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsFrrLogTable.EntityData.YListKeys = []string {}

    return &(cmplsFrrLogTable.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry
// An entry in this table is created to describe one fast
// reroute event.  Entries in this table are only created and
// destroyed by the agent implementation. The maximum number 
// of entries in this log is governed by the scalar.
type CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a fast reroute event entry.
    // The type is interface{} with range: 0..4294967295.
    CmplsFrrLogIndex interface{}

    // This object provides the amount of time ticks since this event occured. The
    // type is interface{} with range: 0..4294967295.
    CmplsFrrLogEventTime interface{}

    // This object indicates which interface was affected by this FRR event. This
    // value may be set to 0 if mplsFrrConstProtectionMethod is set to
    // oneToOneBackup(0). The type is interface{} with range: 0..2147483647.
    CmplsFrrLogInterface interface{}

    // This object describes what type of fast reroute event occured. The type is
    // CmplsFrrLogEventType.
    CmplsFrrLogEventType interface{}

    // This object describes the duration of this event. The type is interface{}
    // with range: 0..4294967295.
    CmplsFrrLogEventDuration interface{}

    // This object contains an implementation-specific explanation of the event.
    // The type is string with length: 128.
    CmplsFrrLogEventReasonString interface{}
}

func (cmplsFrrLogEntry *CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry) GetEntityData() *types.CommonEntityData {
    cmplsFrrLogEntry.EntityData.YFilter = cmplsFrrLogEntry.YFilter
    cmplsFrrLogEntry.EntityData.YangName = "cmplsFrrLogEntry"
    cmplsFrrLogEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsFrrLogEntry.EntityData.ParentYangName = "cmplsFrrLogTable"
    cmplsFrrLogEntry.EntityData.SegmentPath = "cmplsFrrLogEntry" + types.AddKeyToken(cmplsFrrLogEntry.CmplsFrrLogIndex, "cmplsFrrLogIndex")
    cmplsFrrLogEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsFrrLogEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsFrrLogEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsFrrLogEntry.EntityData.Children = types.NewOrderedMap()
    cmplsFrrLogEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsFrrLogEntry.EntityData.Leafs.Append("cmplsFrrLogIndex", types.YLeaf{"CmplsFrrLogIndex", cmplsFrrLogEntry.CmplsFrrLogIndex})
    cmplsFrrLogEntry.EntityData.Leafs.Append("cmplsFrrLogEventTime", types.YLeaf{"CmplsFrrLogEventTime", cmplsFrrLogEntry.CmplsFrrLogEventTime})
    cmplsFrrLogEntry.EntityData.Leafs.Append("cmplsFrrLogInterface", types.YLeaf{"CmplsFrrLogInterface", cmplsFrrLogEntry.CmplsFrrLogInterface})
    cmplsFrrLogEntry.EntityData.Leafs.Append("cmplsFrrLogEventType", types.YLeaf{"CmplsFrrLogEventType", cmplsFrrLogEntry.CmplsFrrLogEventType})
    cmplsFrrLogEntry.EntityData.Leafs.Append("cmplsFrrLogEventDuration", types.YLeaf{"CmplsFrrLogEventDuration", cmplsFrrLogEntry.CmplsFrrLogEventDuration})
    cmplsFrrLogEntry.EntityData.Leafs.Append("cmplsFrrLogEventReasonString", types.YLeaf{"CmplsFrrLogEventReasonString", cmplsFrrLogEntry.CmplsFrrLogEventReasonString})

    cmplsFrrLogEntry.EntityData.YListKeys = []string {"CmplsFrrLogIndex"}

    return &(cmplsFrrLogEntry.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry_CmplsFrrLogEventType represents occured.
type CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry_CmplsFrrLogEventType string

const (
    CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry_CmplsFrrLogEventType_other CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry_CmplsFrrLogEventType = "other"

    CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry_CmplsFrrLogEventType_protected CISCOIETFFRRMIB_CmplsFrrLogTable_CmplsFrrLogEntry_CmplsFrrLogEventType = "protected"
)

// CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable
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
type CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the mplsFrrDBTable represents a single protected LSP, protected
    // by a backup tunnel and defined for a specific protected interface. Note
    // that for brevity, managers should consult the mplsTunnelTable present in
    // the MPLS-TE MIB for additional information about the protecting and
    // protected tunnels, and the ifEntry in the IF-MIB for the protected
    // interface. The type is slice of
    // CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry.
    CmplsFrrFacRouteDBEntry []*CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry
}

func (cmplsFrrFacRouteDBTable *CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable) GetEntityData() *types.CommonEntityData {
    cmplsFrrFacRouteDBTable.EntityData.YFilter = cmplsFrrFacRouteDBTable.YFilter
    cmplsFrrFacRouteDBTable.EntityData.YangName = "cmplsFrrFacRouteDBTable"
    cmplsFrrFacRouteDBTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsFrrFacRouteDBTable.EntityData.ParentYangName = "CISCO-IETF-FRR-MIB"
    cmplsFrrFacRouteDBTable.EntityData.SegmentPath = "cmplsFrrFacRouteDBTable"
    cmplsFrrFacRouteDBTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsFrrFacRouteDBTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsFrrFacRouteDBTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsFrrFacRouteDBTable.EntityData.Children = types.NewOrderedMap()
    cmplsFrrFacRouteDBTable.EntityData.Children.Append("cmplsFrrFacRouteDBEntry", types.YChild{"CmplsFrrFacRouteDBEntry", nil})
    for i := range cmplsFrrFacRouteDBTable.CmplsFrrFacRouteDBEntry {
        cmplsFrrFacRouteDBTable.EntityData.Children.Append(types.GetSegmentPath(cmplsFrrFacRouteDBTable.CmplsFrrFacRouteDBEntry[i]), types.YChild{"CmplsFrrFacRouteDBEntry", cmplsFrrFacRouteDBTable.CmplsFrrFacRouteDBEntry[i]})
    }
    cmplsFrrFacRouteDBTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsFrrFacRouteDBTable.EntityData.YListKeys = []string {}

    return &(cmplsFrrFacRouteDBTable.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry
// An entry in the mplsFrrDBTable represents a single protected
// LSP, protected by a backup tunnel and defined for a specific
// protected interface. Note that for brevity, managers should
// consult the mplsTunnelTable present in the MPLS-TE MIB for
// additional information about the protecting and protected
// tunnels, and the ifEntry in the IF-MIB for the protected
// interface.
type CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies the interface configured for
    // FRR protection. The type is interface{} with range: 1..2147483647.
    CmplsFrrFacRouteProtectedIfIndex interface{}

    // This attribute is a key. Uniquely identifies the mplsTunnelEntry primary
    // index for the tunnel head interface designated to protect the  interface as
    // specified in the mplsFrrFacRouteIfProtectedIndex (and all of the tunnels
    // using this interface). The type is interface{} with range: 0..65535.
    CmplsFrrFacRouteProtectingTunIndex interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is interface{} with range: 0..65535.
    CmplsFrrFacRouteProtectedTunIndex interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is interface{} with range: 0..4294967295.
    CmplsFrrFacRouteProtectedTunInstance interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is string with length: 4.
    CmplsFrrFacRouteProtectedTunIngressLSRId interface{}

    // This attribute is a key. Uniquely identifies an mplsTunnelEntry that is
    // being protected by FRR. The type is string with length: 4.
    CmplsFrrFacRouteProtectedTunEgressLSRId interface{}

    // Specifies the state of the protected tunnel.  active  This tunnel's label
    // has been placed in the          LFIB and is ready to be applied to incoming
    // packets.           ready -  This tunnel's label entry has been created but
    // is          not yet in the LFIB.           partial - This tunnel's label
    // entry as not been fully           created. The type is
    // CmplsFrrFacRouteProtectedTunStatus.
    CmplsFrrFacRouteProtectedTunStatus interface{}

    // Specifies the amount of bandwidth in megabytes per second that is actually
    // reserved by the backup tunnel for facility backup. This value is repeated
    // here from the MPLS- TE MIB because the tunnel entry will reveal the
    // bandwidth reserved by the signaling protocol, which is typically 0 for
    // backup tunnels so as to not over-book bandwidth. However, internal
    // reservations are typically made on the PLR, thus this value should be
    // revealed here. The type is interface{} with range: 0..4294967295.
    CmplsFrrFacRouteProtectingTunResvBw interface{}

    // Indicates type of the resource protection. The type is
    // CmplsFrrFacRouteProtectingTunProtectionType.
    CmplsFrrFacRouteProtectingTunProtectionType interface{}
}

func (cmplsFrrFacRouteDBEntry *CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry) GetEntityData() *types.CommonEntityData {
    cmplsFrrFacRouteDBEntry.EntityData.YFilter = cmplsFrrFacRouteDBEntry.YFilter
    cmplsFrrFacRouteDBEntry.EntityData.YangName = "cmplsFrrFacRouteDBEntry"
    cmplsFrrFacRouteDBEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsFrrFacRouteDBEntry.EntityData.ParentYangName = "cmplsFrrFacRouteDBTable"
    cmplsFrrFacRouteDBEntry.EntityData.SegmentPath = "cmplsFrrFacRouteDBEntry" + types.AddKeyToken(cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedIfIndex, "cmplsFrrFacRouteProtectedIfIndex") + types.AddKeyToken(cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectingTunIndex, "cmplsFrrFacRouteProtectingTunIndex") + types.AddKeyToken(cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunIndex, "cmplsFrrFacRouteProtectedTunIndex") + types.AddKeyToken(cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunInstance, "cmplsFrrFacRouteProtectedTunInstance") + types.AddKeyToken(cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunIngressLSRId, "cmplsFrrFacRouteProtectedTunIngressLSRId") + types.AddKeyToken(cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunEgressLSRId, "cmplsFrrFacRouteProtectedTunEgressLSRId")
    cmplsFrrFacRouteDBEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsFrrFacRouteDBEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsFrrFacRouteDBEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsFrrFacRouteDBEntry.EntityData.Children = types.NewOrderedMap()
    cmplsFrrFacRouteDBEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectedIfIndex", types.YLeaf{"CmplsFrrFacRouteProtectedIfIndex", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedIfIndex})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectingTunIndex", types.YLeaf{"CmplsFrrFacRouteProtectingTunIndex", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectingTunIndex})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectedTunIndex", types.YLeaf{"CmplsFrrFacRouteProtectedTunIndex", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunIndex})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectedTunInstance", types.YLeaf{"CmplsFrrFacRouteProtectedTunInstance", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunInstance})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectedTunIngressLSRId", types.YLeaf{"CmplsFrrFacRouteProtectedTunIngressLSRId", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunIngressLSRId})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectedTunEgressLSRId", types.YLeaf{"CmplsFrrFacRouteProtectedTunEgressLSRId", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunEgressLSRId})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectedTunStatus", types.YLeaf{"CmplsFrrFacRouteProtectedTunStatus", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectedTunStatus})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectingTunResvBw", types.YLeaf{"CmplsFrrFacRouteProtectingTunResvBw", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectingTunResvBw})
    cmplsFrrFacRouteDBEntry.EntityData.Leafs.Append("cmplsFrrFacRouteProtectingTunProtectionType", types.YLeaf{"CmplsFrrFacRouteProtectingTunProtectionType", cmplsFrrFacRouteDBEntry.CmplsFrrFacRouteProtectingTunProtectionType})

    cmplsFrrFacRouteDBEntry.EntityData.YListKeys = []string {"CmplsFrrFacRouteProtectedIfIndex", "CmplsFrrFacRouteProtectingTunIndex", "CmplsFrrFacRouteProtectedTunIndex", "CmplsFrrFacRouteProtectedTunInstance", "CmplsFrrFacRouteProtectedTunIngressLSRId", "CmplsFrrFacRouteProtectedTunEgressLSRId"}

    return &(cmplsFrrFacRouteDBEntry.EntityData)
}

// CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus represents           created.
type CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus string

const (
    CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus_active CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus = "active"

    CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus_ready CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus = "ready"

    CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus_partial CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectedTunStatus = "partial"
)

// CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectingTunProtectionType represents Indicates type of the resource protection.
type CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectingTunProtectionType string

const (
    CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectingTunProtectionType_linkProtection CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectingTunProtectionType = "linkProtection"

    CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectingTunProtectionType_nodeProtection CISCOIETFFRRMIB_CmplsFrrFacRouteDBTable_CmplsFrrFacRouteDBEntry_CmplsFrrFacRouteProtectingTunProtectionType = "nodeProtection"
)

