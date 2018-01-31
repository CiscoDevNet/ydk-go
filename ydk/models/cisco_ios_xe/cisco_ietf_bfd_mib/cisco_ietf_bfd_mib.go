// This document contains the Management information base for
// Bidirectional Forwarding Detection(BFD) Protocol as defined
// in draft-ietf-bfd-base-06.txt.
// 
// BFD is a protocol intended to detect faults in the
// bidirectional path between two forwarding engines, including
// interfaces, data link(s), and to the extent possible the forwarding
// engines themselves, with potentially very low latency.  It operates
// independently of media, data protocols, and routing protocols.
// 
// This MIB module is based on the Internet Draft
// draft-ietf-bfd-mib-03.txt and draft-ietf-bfd-mib-04.txt
package cisco_ietf_bfd_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_bfd_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-BFD-MIB CISCO-IETF-BFD-MIB}", reflect.TypeOf(CISCOIETFBFDMIB{}))
    ydk.RegisterEntity("CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB", reflect.TypeOf(CISCOIETFBFDMIB{}))
}

// CiscoBfdDiag represents A common BFD diagnostic code.
type CiscoBfdDiag string

const (
    CiscoBfdDiag_noDiagnostic CiscoBfdDiag = "noDiagnostic"

    CiscoBfdDiag_controlDetectionTimeExpired CiscoBfdDiag = "controlDetectionTimeExpired"

    CiscoBfdDiag_echoFunctionFailed CiscoBfdDiag = "echoFunctionFailed"

    CiscoBfdDiag_neighborSignaledSessionDown CiscoBfdDiag = "neighborSignaledSessionDown"

    CiscoBfdDiag_forwardingPlaneReset CiscoBfdDiag = "forwardingPlaneReset"

    CiscoBfdDiag_pathDown CiscoBfdDiag = "pathDown"

    CiscoBfdDiag_concatenatedPathDown CiscoBfdDiag = "concatenatedPathDown"

    CiscoBfdDiag_administrativelyDown CiscoBfdDiag = "administrativelyDown"

    CiscoBfdDiag_reverseConcatenatedPathDown CiscoBfdDiag = "reverseConcatenatedPathDown"
)

// CISCOIETFBFDMIB
type CISCOIETFBFDMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ciscobfdscalarobjects CISCOIETFBFDMIB_Ciscobfdscalarobjects

    // The BFD Session Table describes the BFD sessions.
    Ciscobfdsesstable CISCOIETFBFDMIB_Ciscobfdsesstable

    // The BFD Session Mapping Table maps the complex indexing of the BFD sessions
    // to the flat  CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
    Ciscobfdsessmaptable CISCOIETFBFDMIB_Ciscobfdsessmaptable

    // The BFD Session Discriminator Mapping Table maps a local discriminator
    // value to associated BFD sessions' CiscoBfdSessIndexTC used in the
    // ciscoBfdSessTable.
    Ciscobfdsessdiscmaptable CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable

    // The BFD Session IP Mapping Table maps given ciscoBfdSessInterface,
    // ciscoBfdSessAddrType, and ciscoBbfdSessAddr to an associated BFD sessions'
    // CiscoBfdSessIndexTC used in the ciscoBfdSessTable. This table SHOULD
    // contains those BFD sessions are of IP type: singleHop(1) and multiHop(2).
    Ciscobfdsessipmaptable CISCOIETFBFDMIB_Ciscobfdsessipmaptable
}

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetFilter() yfilter.YFilter { return cISCOIETFBFDMIB.YFilter }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFBFDMIB.YFilter = yf }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetGoName(yname string) string {
    if yname == "ciscoBfdScalarObjects" { return "Ciscobfdscalarobjects" }
    if yname == "ciscoBfdSessTable" { return "Ciscobfdsesstable" }
    if yname == "ciscoBfdSessMapTable" { return "Ciscobfdsessmaptable" }
    if yname == "ciscoBfdSessDiscMapTable" { return "Ciscobfdsessdiscmaptable" }
    if yname == "ciscoBfdSessIpMapTable" { return "Ciscobfdsessipmaptable" }
    return ""
}

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetSegmentPath() string {
    return "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB"
}

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoBfdScalarObjects" {
        return &cISCOIETFBFDMIB.Ciscobfdscalarobjects
    }
    if childYangName == "ciscoBfdSessTable" {
        return &cISCOIETFBFDMIB.Ciscobfdsesstable
    }
    if childYangName == "ciscoBfdSessMapTable" {
        return &cISCOIETFBFDMIB.Ciscobfdsessmaptable
    }
    if childYangName == "ciscoBfdSessDiscMapTable" {
        return &cISCOIETFBFDMIB.Ciscobfdsessdiscmaptable
    }
    if childYangName == "ciscoBfdSessIpMapTable" {
        return &cISCOIETFBFDMIB.Ciscobfdsessipmaptable
    }
    return nil
}

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoBfdScalarObjects"] = &cISCOIETFBFDMIB.Ciscobfdscalarobjects
    children["ciscoBfdSessTable"] = &cISCOIETFBFDMIB.Ciscobfdsesstable
    children["ciscoBfdSessMapTable"] = &cISCOIETFBFDMIB.Ciscobfdsessmaptable
    children["ciscoBfdSessDiscMapTable"] = &cISCOIETFBFDMIB.Ciscobfdsessdiscmaptable
    children["ciscoBfdSessIpMapTable"] = &cISCOIETFBFDMIB.Ciscobfdsessipmaptable
    return children
}

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetYangName() string { return "CISCO-IETF-BFD-MIB" }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) SetParent(parent types.Entity) { cISCOIETFBFDMIB.parent = parent }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetParent() types.Entity { return cISCOIETFBFDMIB.parent }

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetParentYangName() string { return "CISCO-IETF-BFD-MIB" }

// CISCOIETFBFDMIB_Ciscobfdscalarobjects
type CISCOIETFBFDMIB_Ciscobfdscalarobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The global administrative status of BFD in this router. The value 'enabled'
    // denotes that the BFD Process is  active on at least one interface;
    // 'disabled' disables   it on all interfaces. The type is
    // Ciscobfdadminstatus.
    Ciscobfdadminstatus interface{}

    // The current version number of the BFD protocol. The type is interface{}
    // with range: 0..4294967295.
    Ciscobfdversionnumber interface{}

    // If this object is set to true(1), then it enables the emission of
    // ciscoBfdSessUp and ciscoBfdSessDown  notifications; otherwise these
    // notifications are not  emitted. The type is bool.
    Ciscobfdsessnotificationsenable interface{}
}

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetFilter() yfilter.YFilter { return ciscobfdscalarobjects.YFilter }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) SetFilter(yf yfilter.YFilter) { ciscobfdscalarobjects.YFilter = yf }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetGoName(yname string) string {
    if yname == "ciscoBfdAdminStatus" { return "Ciscobfdadminstatus" }
    if yname == "ciscoBfdVersionNumber" { return "Ciscobfdversionnumber" }
    if yname == "ciscoBfdSessNotificationsEnable" { return "Ciscobfdsessnotificationsenable" }
    return ""
}

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetSegmentPath() string {
    return "ciscoBfdScalarObjects"
}

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoBfdAdminStatus"] = ciscobfdscalarobjects.Ciscobfdadminstatus
    leafs["ciscoBfdVersionNumber"] = ciscobfdscalarobjects.Ciscobfdversionnumber
    leafs["ciscoBfdSessNotificationsEnable"] = ciscobfdscalarobjects.Ciscobfdsessnotificationsenable
    return leafs
}

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetYangName() string { return "ciscoBfdScalarObjects" }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) SetParent(parent types.Entity) { ciscobfdscalarobjects.parent = parent }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetParent() types.Entity { return ciscobfdscalarobjects.parent }

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetParentYangName() string { return "CISCO-IETF-BFD-MIB" }

// CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus represents it on all interfaces.
type CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus string

const (
    CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus_enabled CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus = "enabled"

    CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus_disabled CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus = "disabled"
)

// CISCOIETFBFDMIB_Ciscobfdsesstable
// The BFD Session Table describes the BFD sessions.
type CISCOIETFBFDMIB_Ciscobfdsesstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The BFD Session Entry describes BFD session. The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry.
    Ciscobfdsessentry []CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry
}

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetFilter() yfilter.YFilter { return ciscobfdsesstable.YFilter }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) SetFilter(yf yfilter.YFilter) { ciscobfdsesstable.YFilter = yf }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetGoName(yname string) string {
    if yname == "ciscoBfdSessEntry" { return "Ciscobfdsessentry" }
    return ""
}

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetSegmentPath() string {
    return "ciscoBfdSessTable"
}

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoBfdSessEntry" {
        for _, c := range ciscobfdsesstable.Ciscobfdsessentry {
            if ciscobfdsesstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry{}
        ciscobfdsesstable.Ciscobfdsessentry = append(ciscobfdsesstable.Ciscobfdsessentry, child)
        return &ciscobfdsesstable.Ciscobfdsessentry[len(ciscobfdsesstable.Ciscobfdsessentry)-1]
    }
    return nil
}

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscobfdsesstable.Ciscobfdsessentry {
        children[ciscobfdsesstable.Ciscobfdsessentry[i].GetSegmentPath()] = &ciscobfdsesstable.Ciscobfdsessentry[i]
    }
    return children
}

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetYangName() string { return "ciscoBfdSessTable" }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) SetParent(parent types.Entity) { ciscobfdsesstable.parent = parent }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetParent() types.Entity { return ciscobfdsesstable.parent }

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetParentYangName() string { return "CISCO-IETF-BFD-MIB" }

// CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry
// The BFD Session Entry describes BFD session.
type CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object contains an index used to represent a
    // unique BFD session on this device. The type is interface{} with range:
    // 1..4294967295.
    Ciscobfdsessindex interface{}

    // This object contains an index used to indicate a local application which
    // owns or maintains this  BFD session. For instance, the MPLS VPN process may
    // maintain a subset of the total number of BFD  sessions.  This application
    // ID provides a convenient  way to segregate sessions by the applications
    // which  maintain them. The type is interface{} with range: 0..4294967295.
    Ciscobfdsessapplicationid interface{}

    // This object specifies the local discriminator for this BFD session, used to
    // uniquely identify it. The type is interface{} with range: 1..4294967295.
    Ciscobfdsessdiscriminator interface{}

    // This object specifies the session discriminator chosen by the remote system
    // for this BFD session. The type is interface{} with range: 0..4294967295.
    Ciscobfdsessremotediscr interface{}

    // The destination UDP Port for BFD. The default value is the well-known value
    // for this port. BFD State failing(5) is only applicable if this BFD session
    // is running version 0. The type is interface{} with range: 0..65535.
    Ciscobfdsessudpport interface{}

    // The perceived state of the BFD session. The type is Ciscobfdsessstate.
    Ciscobfdsessstate interface{}

    // This object specifies status of BFD packet reception from the remote
    // system. Specifically, it is set to true(1) if  the local system is actively
    // receiving BFD packets from the   remote system, and is set to false(0) if
    // the local system   has not received BFD packets recently (within the
    // detection   time) or if the local system is attempting to tear down  the
    // BFD session. The type is bool.
    Ciscobfdsessremoteheardflag interface{}

    // A diagnostic code specifying the local system's reason for the last
    // transition of the session from up(1)   to some other state. The type is
    // CiscoBfdDiag.
    Ciscobfdsessdiag interface{}

    // This object specifies current operating mode that BFD session is operating
    // in. The type is Ciscobfdsessopermode.
    Ciscobfdsessopermode interface{}

    // This object indicates that the local system's desire to use Demand mode.
    // Specifically, it is set   to true(1) if the local system wishes to use  
    // Demand mode or false(0) if not. The type is bool.
    Ciscobfdsessdemandmodedesiredflag interface{}

    // This object indicates that the local system's desire to use Echo mode.
    // Specifically, it is set   to true(1) if the local system wishes to use  
    // Echo mode or false(0) if not. The type is bool.
    Ciscobfdsessechofuncmodedesiredflag interface{}

    // This object indicates that the local system's ability to continue to
    // function through a disruption of   the control plane. Specifically, it is
    // set   to true(1) if the local system BFD implementation is  independent of
    // the control plane. Otherwise, the   value is set to false(0). The type is
    // bool.
    Ciscobfdsesscontrolplanindepflag interface{}

    // This object specifies IP address type of the neighboring IP address which
    // is being monitored with this BFD session.  Only values unknown(0), ipv4(1)
    // or ipv6(2)  have to be supported.    A value of unknown(0) is allowed only
    // when   the outgoing interface is of type point-to-point, or  when the BFD
    // session is not associated with a specific   interface.   If any other
    // unsupported values are attempted in a set  operation, the agent MUST return
    // an inconsistentValue   error. The type is InetAddressType.
    Ciscobfdsessaddrtype interface{}

    // This object specifies the neighboring IP address which is being monitored
    // with this BFD session. It can also be used to enabled BFD on a specific  
    // interface. The value is set to zero when BFD session is not   associated
    // with a specific interface. The type is string with length: 0..255.
    Ciscobfdsessaddr interface{}

    // This object specifies the minimum interval, in microseconds, that the local
    // system would like to use when       transmitting BFD Control packets. The
    // type is interface{} with range: 0..4294967295.
    Ciscobfdsessdesiredmintxinterval interface{}

    // This object specifies the minimum interval, in microseconds, between
    // received  BFD Control packets the   local system is capable of supporting.
    // The type is interface{} with range: 0..4294967295.
    Ciscobfdsessreqminrxinterval interface{}

    // This object specifies the minimum interval, in microseconds, between
    // received BFD Echo packets that this  system is capable of supporting. The
    // type is interface{} with range: 0..4294967295.
    Ciscobfdsessreqminechorxinterval interface{}

    // This object specifies the Detect time multiplier. The type is interface{}
    // with range: 0..4294967295.
    Ciscobfdsessdetectmult interface{}

    // This variable indicates the storage type for this object. Conceptual rows
    // having the value   'permanent' need not allow write-access to any  
    // columnar objects in the row. The type is StorageType.
    Ciscobfdsessstortype interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this  table has a row in the active(1) state, no   objects in
    // this row can be modified except the  ciscoBfdSessRowStatus and
    // ciscoBfdSessStorageType. The type is RowStatus.
    Ciscobfdsessrowstatus interface{}

    // This object indicates that the local system's desire to use Authentication.
    // Specifically, it is set   to true(1) if the local system wishes the session
    // to be authenticated or false(0) if not. The type is bool.
    Ciscobfdsessauthpresflag interface{}

    // The Authentication Type used for this BFD session. This field is valid only
    // when the Authentication Present bit is set. The type is
    // Ciscobfdsessauthenticationtype.
    Ciscobfdsessauthenticationtype interface{}

    // The version number of the BFD protocol that this session is running in. The
    // type is interface{} with range: 0..4294967295.
    Ciscobfdsessversionnumber interface{}

    // The type of this BFD session. The type is Ciscobfdsesstype.
    Ciscobfdsesstype interface{}

    // This object contains an interface index used to indicate the interface
    // which this BFD session is running on. The type is interface{} with range:
    // 1..2147483647.
    Ciscobfdsessinterface interface{}

    // The total number of BFD messages received for this BFD session. The type is
    // interface{} with range: 0..4294967295.
    Ciscobfdsessperfpktin interface{}

    // The total number of BFD messages sent for this BFD session. The type is
    // interface{} with range: 0..4294967295.
    Ciscobfdsessperfpktout interface{}

    // The value of sysUpTime on the most recent occasion at which the session
    // came up. If no such up event exists this object  contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    Ciscobfdsessuptime interface{}

    // The value of sysUpTime on the most recent occasion at which the last time
    // communication was lost with the neighbor. If   no such down event exist
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    Ciscobfdsessperflastsessdowntime interface{}

    // The BFD diag code for the last time communication was lost with the
    // neighbor. If no such down event exists this object   contains a zero value.
    // The type is CiscoBfdDiag.
    Ciscobfdsessperflastcommlostdiag interface{}

    // The number of times this session has gone into the Up state since the
    // router last rebooted. The type is interface{} with range: 0..4294967295.
    Ciscobfdsessperfsessupcount interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of the session counters suffered  a discontinuity.    The relevant counters
    // are the specific instances associated   with this BFD session of any
    // Counter32 object contained in  the ciscoBfdSessPerfTable. If no such
    // discontinuities have occurred   since the last re-initialization of the
    // local management  subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    Ciscobfdsessperfdisctime interface{}

    // This value represents the total number of BFD messages received for this
    // BFD session. It MUST be equal to the  least significant 32 bits of
    // ciscoBfdSessPerfPktIn  if ciscoBfdSessPerfPktInHC is supported according to
    // the rules spelled out in RFC2863. The type is interface{} with range:
    // 0..18446744073709551615.
    Ciscobfdsessperfpktinhc interface{}

    // This value represents the total number of total number of BFD messages
    // transmitted for this   BFD session. It MUST be equal to the  least
    // significant 32 bits of ciscoBfdSessPerfPktIn  if ciscoBfdSessPerfPktOutHC
    // is supported according to  the rules spelled out in RFC2863. The type is
    // interface{} with range: 0..18446744073709551615.
    Ciscobfdsessperfpktouthc interface{}
}

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetFilter() yfilter.YFilter { return ciscobfdsessentry.YFilter }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) SetFilter(yf yfilter.YFilter) { ciscobfdsessentry.YFilter = yf }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetGoName(yname string) string {
    if yname == "ciscoBfdSessIndex" { return "Ciscobfdsessindex" }
    if yname == "ciscoBfdSessApplicationId" { return "Ciscobfdsessapplicationid" }
    if yname == "ciscoBfdSessDiscriminator" { return "Ciscobfdsessdiscriminator" }
    if yname == "ciscoBfdSessRemoteDiscr" { return "Ciscobfdsessremotediscr" }
    if yname == "ciscoBfdSessUdpPort" { return "Ciscobfdsessudpport" }
    if yname == "ciscoBfdSessState" { return "Ciscobfdsessstate" }
    if yname == "ciscoBfdSessRemoteHeardFlag" { return "Ciscobfdsessremoteheardflag" }
    if yname == "ciscoBfdSessDiag" { return "Ciscobfdsessdiag" }
    if yname == "ciscoBfdSessOperMode" { return "Ciscobfdsessopermode" }
    if yname == "ciscoBfdSessDemandModeDesiredFlag" { return "Ciscobfdsessdemandmodedesiredflag" }
    if yname == "ciscoBfdSessEchoFuncModeDesiredFlag" { return "Ciscobfdsessechofuncmodedesiredflag" }
    if yname == "ciscoBfdSessControlPlanIndepFlag" { return "Ciscobfdsesscontrolplanindepflag" }
    if yname == "ciscoBfdSessAddrType" { return "Ciscobfdsessaddrtype" }
    if yname == "ciscoBfdSessAddr" { return "Ciscobfdsessaddr" }
    if yname == "ciscoBfdSessDesiredMinTxInterval" { return "Ciscobfdsessdesiredmintxinterval" }
    if yname == "ciscoBfdSessReqMinRxInterval" { return "Ciscobfdsessreqminrxinterval" }
    if yname == "ciscoBfdSessReqMinEchoRxInterval" { return "Ciscobfdsessreqminechorxinterval" }
    if yname == "ciscoBfdSessDetectMult" { return "Ciscobfdsessdetectmult" }
    if yname == "ciscoBfdSessStorType" { return "Ciscobfdsessstortype" }
    if yname == "ciscoBfdSessRowStatus" { return "Ciscobfdsessrowstatus" }
    if yname == "ciscoBfdSessAuthPresFlag" { return "Ciscobfdsessauthpresflag" }
    if yname == "ciscoBfdSessAuthenticationType" { return "Ciscobfdsessauthenticationtype" }
    if yname == "ciscoBfdSessVersionNumber" { return "Ciscobfdsessversionnumber" }
    if yname == "ciscoBfdSessType" { return "Ciscobfdsesstype" }
    if yname == "ciscoBfdSessInterface" { return "Ciscobfdsessinterface" }
    if yname == "ciscoBfdSessPerfPktIn" { return "Ciscobfdsessperfpktin" }
    if yname == "ciscoBfdSessPerfPktOut" { return "Ciscobfdsessperfpktout" }
    if yname == "ciscoBfdSessUpTime" { return "Ciscobfdsessuptime" }
    if yname == "ciscoBfdSessPerfLastSessDownTime" { return "Ciscobfdsessperflastsessdowntime" }
    if yname == "ciscoBfdSessPerfLastCommLostDiag" { return "Ciscobfdsessperflastcommlostdiag" }
    if yname == "ciscoBfdSessPerfSessUpCount" { return "Ciscobfdsessperfsessupcount" }
    if yname == "ciscoBfdSessPerfDiscTime" { return "Ciscobfdsessperfdisctime" }
    if yname == "ciscoBfdSessPerfPktInHC" { return "Ciscobfdsessperfpktinhc" }
    if yname == "ciscoBfdSessPerfPktOutHC" { return "Ciscobfdsessperfpktouthc" }
    return ""
}

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetSegmentPath() string {
    return "ciscoBfdSessEntry" + "[ciscoBfdSessIndex='" + fmt.Sprintf("%v", ciscobfdsessentry.Ciscobfdsessindex) + "']"
}

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoBfdSessIndex"] = ciscobfdsessentry.Ciscobfdsessindex
    leafs["ciscoBfdSessApplicationId"] = ciscobfdsessentry.Ciscobfdsessapplicationid
    leafs["ciscoBfdSessDiscriminator"] = ciscobfdsessentry.Ciscobfdsessdiscriminator
    leafs["ciscoBfdSessRemoteDiscr"] = ciscobfdsessentry.Ciscobfdsessremotediscr
    leafs["ciscoBfdSessUdpPort"] = ciscobfdsessentry.Ciscobfdsessudpport
    leafs["ciscoBfdSessState"] = ciscobfdsessentry.Ciscobfdsessstate
    leafs["ciscoBfdSessRemoteHeardFlag"] = ciscobfdsessentry.Ciscobfdsessremoteheardflag
    leafs["ciscoBfdSessDiag"] = ciscobfdsessentry.Ciscobfdsessdiag
    leafs["ciscoBfdSessOperMode"] = ciscobfdsessentry.Ciscobfdsessopermode
    leafs["ciscoBfdSessDemandModeDesiredFlag"] = ciscobfdsessentry.Ciscobfdsessdemandmodedesiredflag
    leafs["ciscoBfdSessEchoFuncModeDesiredFlag"] = ciscobfdsessentry.Ciscobfdsessechofuncmodedesiredflag
    leafs["ciscoBfdSessControlPlanIndepFlag"] = ciscobfdsessentry.Ciscobfdsesscontrolplanindepflag
    leafs["ciscoBfdSessAddrType"] = ciscobfdsessentry.Ciscobfdsessaddrtype
    leafs["ciscoBfdSessAddr"] = ciscobfdsessentry.Ciscobfdsessaddr
    leafs["ciscoBfdSessDesiredMinTxInterval"] = ciscobfdsessentry.Ciscobfdsessdesiredmintxinterval
    leafs["ciscoBfdSessReqMinRxInterval"] = ciscobfdsessentry.Ciscobfdsessreqminrxinterval
    leafs["ciscoBfdSessReqMinEchoRxInterval"] = ciscobfdsessentry.Ciscobfdsessreqminechorxinterval
    leafs["ciscoBfdSessDetectMult"] = ciscobfdsessentry.Ciscobfdsessdetectmult
    leafs["ciscoBfdSessStorType"] = ciscobfdsessentry.Ciscobfdsessstortype
    leafs["ciscoBfdSessRowStatus"] = ciscobfdsessentry.Ciscobfdsessrowstatus
    leafs["ciscoBfdSessAuthPresFlag"] = ciscobfdsessentry.Ciscobfdsessauthpresflag
    leafs["ciscoBfdSessAuthenticationType"] = ciscobfdsessentry.Ciscobfdsessauthenticationtype
    leafs["ciscoBfdSessVersionNumber"] = ciscobfdsessentry.Ciscobfdsessversionnumber
    leafs["ciscoBfdSessType"] = ciscobfdsessentry.Ciscobfdsesstype
    leafs["ciscoBfdSessInterface"] = ciscobfdsessentry.Ciscobfdsessinterface
    leafs["ciscoBfdSessPerfPktIn"] = ciscobfdsessentry.Ciscobfdsessperfpktin
    leafs["ciscoBfdSessPerfPktOut"] = ciscobfdsessentry.Ciscobfdsessperfpktout
    leafs["ciscoBfdSessUpTime"] = ciscobfdsessentry.Ciscobfdsessuptime
    leafs["ciscoBfdSessPerfLastSessDownTime"] = ciscobfdsessentry.Ciscobfdsessperflastsessdowntime
    leafs["ciscoBfdSessPerfLastCommLostDiag"] = ciscobfdsessentry.Ciscobfdsessperflastcommlostdiag
    leafs["ciscoBfdSessPerfSessUpCount"] = ciscobfdsessentry.Ciscobfdsessperfsessupcount
    leafs["ciscoBfdSessPerfDiscTime"] = ciscobfdsessentry.Ciscobfdsessperfdisctime
    leafs["ciscoBfdSessPerfPktInHC"] = ciscobfdsessentry.Ciscobfdsessperfpktinhc
    leafs["ciscoBfdSessPerfPktOutHC"] = ciscobfdsessentry.Ciscobfdsessperfpktouthc
    return leafs
}

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetYangName() string { return "ciscoBfdSessEntry" }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) SetParent(parent types.Entity) { ciscobfdsessentry.parent = parent }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetParent() types.Entity { return ciscobfdsessentry.parent }

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetParentYangName() string { return "ciscoBfdSessTable" }

// CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype represents field is valid only when the Authentication Present bit is set
type CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype string

const (
    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype_simplePassword CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype = "simplePassword"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype_keyedMD5 CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype = "keyedMD5"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype_meticulousKeyedMD5 CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype = "meticulousKeyedMD5"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype_keyedSHA1 CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype = "keyedSHA1"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype_meticulousKeyedSHA1 CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessauthenticationtype = "meticulousKeyedSHA1"
)

// CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode represents session is operating in.
type CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode string

const (
    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode_asyncModeWEchoFun CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode = "asyncModeWEchoFun"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode_asynchModeWOEchoFun CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode = "asynchModeWOEchoFun"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode_demandModeWEchoFunction CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode = "demandModeWEchoFunction"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode_demandModeWOEchoFunction CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessopermode = "demandModeWOEchoFunction"
)

// CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate represents The perceived state of the BFD session.
type CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate string

const (
    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate_adminDown CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate = "adminDown"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate_down CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate = "down"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate_init CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate = "init"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate_up CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate = "up"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate_failing CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessstate = "failing"
)

// CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsesstype represents The type of this BFD session.
type CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsesstype string

const (
    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsesstype_singleHop CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsesstype = "singleHop"

    CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsesstype_multiHop CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsesstype = "multiHop"
)

// CISCOIETFBFDMIB_Ciscobfdsessmaptable
// The BFD Session Mapping Table maps the complex
// indexing of the BFD sessions to the flat 
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
type CISCOIETFBFDMIB_Ciscobfdsessmaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The BFD Session Entry describes BFD session that is mapped to this index.
    // The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry.
    Ciscobfdsessmapentry []CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry
}

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetFilter() yfilter.YFilter { return ciscobfdsessmaptable.YFilter }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) SetFilter(yf yfilter.YFilter) { ciscobfdsessmaptable.YFilter = yf }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetGoName(yname string) string {
    if yname == "ciscoBfdSessMapEntry" { return "Ciscobfdsessmapentry" }
    return ""
}

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetSegmentPath() string {
    return "ciscoBfdSessMapTable"
}

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoBfdSessMapEntry" {
        for _, c := range ciscobfdsessmaptable.Ciscobfdsessmapentry {
            if ciscobfdsessmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry{}
        ciscobfdsessmaptable.Ciscobfdsessmapentry = append(ciscobfdsessmaptable.Ciscobfdsessmapentry, child)
        return &ciscobfdsessmaptable.Ciscobfdsessmapentry[len(ciscobfdsessmaptable.Ciscobfdsessmapentry)-1]
    }
    return nil
}

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscobfdsessmaptable.Ciscobfdsessmapentry {
        children[ciscobfdsessmaptable.Ciscobfdsessmapentry[i].GetSegmentPath()] = &ciscobfdsessmaptable.Ciscobfdsessmapentry[i]
    }
    return children
}

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetYangName() string { return "ciscoBfdSessMapTable" }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) SetParent(parent types.Entity) { ciscobfdsessmaptable.parent = parent }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetParent() types.Entity { return ciscobfdsessmaptable.parent }

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetParentYangName() string { return "CISCO-IETF-BFD-MIB" }

// CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry
// The BFD Session Entry describes BFD session
// that is mapped to this index.
type CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessapplicationid
    Ciscobfdsessapplicationid interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessdiscriminator
    Ciscobfdsessdiscriminator interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessaddrtype
    Ciscobfdsessaddrtype interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessaddr
    Ciscobfdsessaddr interface{}

    // This object indicates the CiscoBfdSessIndexTC referred to by the indices of
    // this row. In essence, a mapping is  provided between these indices and the
    // ciscoBfdSessTable. The type is interface{} with range: 1..4294967295.
    Ciscobfdsessmapbfdindex interface{}
}

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetFilter() yfilter.YFilter { return ciscobfdsessmapentry.YFilter }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) SetFilter(yf yfilter.YFilter) { ciscobfdsessmapentry.YFilter = yf }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetGoName(yname string) string {
    if yname == "ciscoBfdSessApplicationId" { return "Ciscobfdsessapplicationid" }
    if yname == "ciscoBfdSessDiscriminator" { return "Ciscobfdsessdiscriminator" }
    if yname == "ciscoBfdSessAddrType" { return "Ciscobfdsessaddrtype" }
    if yname == "ciscoBfdSessAddr" { return "Ciscobfdsessaddr" }
    if yname == "ciscoBfdSessMapBfdIndex" { return "Ciscobfdsessmapbfdindex" }
    return ""
}

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetSegmentPath() string {
    return "ciscoBfdSessMapEntry" + "[ciscoBfdSessApplicationId='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessapplicationid) + "']" + "[ciscoBfdSessDiscriminator='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessdiscriminator) + "']" + "[ciscoBfdSessAddrType='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessaddrtype) + "']" + "[ciscoBfdSessAddr='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessaddr) + "']"
}

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoBfdSessApplicationId"] = ciscobfdsessmapentry.Ciscobfdsessapplicationid
    leafs["ciscoBfdSessDiscriminator"] = ciscobfdsessmapentry.Ciscobfdsessdiscriminator
    leafs["ciscoBfdSessAddrType"] = ciscobfdsessmapentry.Ciscobfdsessaddrtype
    leafs["ciscoBfdSessAddr"] = ciscobfdsessmapentry.Ciscobfdsessaddr
    leafs["ciscoBfdSessMapBfdIndex"] = ciscobfdsessmapentry.Ciscobfdsessmapbfdindex
    return leafs
}

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetYangName() string { return "ciscoBfdSessMapEntry" }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) SetParent(parent types.Entity) { ciscobfdsessmapentry.parent = parent }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetParent() types.Entity { return ciscobfdsessmapentry.parent }

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetParentYangName() string { return "ciscoBfdSessMapTable" }

// CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable
// The BFD Session Discriminator Mapping Table maps a
// local discriminator value to associated BFD sessions'
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
type CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each row contains a mapping between a local discriminator value to an entry
    // in ciscoBfdSessTable. The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry.
    Ciscobfdsessdiscmapentry []CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry
}

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetFilter() yfilter.YFilter { return ciscobfdsessdiscmaptable.YFilter }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) SetFilter(yf yfilter.YFilter) { ciscobfdsessdiscmaptable.YFilter = yf }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetGoName(yname string) string {
    if yname == "ciscoBfdSessDiscMapEntry" { return "Ciscobfdsessdiscmapentry" }
    return ""
}

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetSegmentPath() string {
    return "ciscoBfdSessDiscMapTable"
}

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoBfdSessDiscMapEntry" {
        for _, c := range ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry {
            if ciscobfdsessdiscmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry{}
        ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry = append(ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry, child)
        return &ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry[len(ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry)-1]
    }
    return nil
}

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry {
        children[ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry[i].GetSegmentPath()] = &ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry[i]
    }
    return children
}

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetYangName() string { return "ciscoBfdSessDiscMapTable" }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) SetParent(parent types.Entity) { ciscobfdsessdiscmaptable.parent = parent }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetParent() types.Entity { return ciscobfdsessdiscmaptable.parent }

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetParentYangName() string { return "CISCO-IETF-BFD-MIB" }

// CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry
// Each row contains a mapping between a local discriminator
// value to an entry in ciscoBfdSessTable.
type CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessdiscriminator
    Ciscobfdsessdiscriminator interface{}

    // This object indicates the CiscoBfdSessIndexTC referred to by the index of
    // this row. In essence, a mapping is  provided between this index and the
    // ciscoBfdSessTable. The type is interface{} with range: 1..4294967295.
    Ciscobfdsessdiscmapindex interface{}
}

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetFilter() yfilter.YFilter { return ciscobfdsessdiscmapentry.YFilter }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) SetFilter(yf yfilter.YFilter) { ciscobfdsessdiscmapentry.YFilter = yf }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetGoName(yname string) string {
    if yname == "ciscoBfdSessDiscriminator" { return "Ciscobfdsessdiscriminator" }
    if yname == "ciscoBfdSessDiscMapIndex" { return "Ciscobfdsessdiscmapindex" }
    return ""
}

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetSegmentPath() string {
    return "ciscoBfdSessDiscMapEntry" + "[ciscoBfdSessDiscriminator='" + fmt.Sprintf("%v", ciscobfdsessdiscmapentry.Ciscobfdsessdiscriminator) + "']"
}

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoBfdSessDiscriminator"] = ciscobfdsessdiscmapentry.Ciscobfdsessdiscriminator
    leafs["ciscoBfdSessDiscMapIndex"] = ciscobfdsessdiscmapentry.Ciscobfdsessdiscmapindex
    return leafs
}

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetYangName() string { return "ciscoBfdSessDiscMapEntry" }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) SetParent(parent types.Entity) { ciscobfdsessdiscmapentry.parent = parent }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetParent() types.Entity { return ciscobfdsessdiscmapentry.parent }

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetParentYangName() string { return "ciscoBfdSessDiscMapTable" }

// CISCOIETFBFDMIB_Ciscobfdsessipmaptable
// The BFD Session IP Mapping Table maps given
// ciscoBfdSessInterface, ciscoBfdSessAddrType, and
// ciscoBbfdSessAddr to an associated BFD sessions'
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
// This table SHOULD contains those BFD sessions are
// of IP type: singleHop(1) and multiHop(2).
type CISCOIETFBFDMIB_Ciscobfdsessipmaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each row contains a mapping between ciscoBfdSessInterface,
    // ciscoBfdSessAddrType and ciscoBfdSessAddr values to an  entry in
    // ciscoBfdSessTable. The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry.
    Ciscobfdsessipmapentry []CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry
}

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetFilter() yfilter.YFilter { return ciscobfdsessipmaptable.YFilter }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) SetFilter(yf yfilter.YFilter) { ciscobfdsessipmaptable.YFilter = yf }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetGoName(yname string) string {
    if yname == "ciscoBfdSessIpMapEntry" { return "Ciscobfdsessipmapentry" }
    return ""
}

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetSegmentPath() string {
    return "ciscoBfdSessIpMapTable"
}

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoBfdSessIpMapEntry" {
        for _, c := range ciscobfdsessipmaptable.Ciscobfdsessipmapentry {
            if ciscobfdsessipmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry{}
        ciscobfdsessipmaptable.Ciscobfdsessipmapentry = append(ciscobfdsessipmaptable.Ciscobfdsessipmapentry, child)
        return &ciscobfdsessipmaptable.Ciscobfdsessipmapentry[len(ciscobfdsessipmaptable.Ciscobfdsessipmapentry)-1]
    }
    return nil
}

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscobfdsessipmaptable.Ciscobfdsessipmapentry {
        children[ciscobfdsessipmaptable.Ciscobfdsessipmapentry[i].GetSegmentPath()] = &ciscobfdsessipmaptable.Ciscobfdsessipmapentry[i]
    }
    return children
}

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetYangName() string { return "ciscoBfdSessIpMapTable" }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) SetParent(parent types.Entity) { ciscobfdsessipmaptable.parent = parent }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetParent() types.Entity { return ciscobfdsessipmaptable.parent }

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetParentYangName() string { return "CISCO-IETF-BFD-MIB" }

// CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry
// Each row contains a mapping between ciscoBfdSessInterface,
// ciscoBfdSessAddrType and ciscoBfdSessAddr values to an 
// entry in ciscoBfdSessTable.
type CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessinterface
    Ciscobfdsessinterface interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessaddrtype
    Ciscobfdsessaddrtype interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry_Ciscobfdsessaddr
    Ciscobfdsessaddr interface{}

    // This object indicates the CiscoBfdSessIndexTC referred to by the indices of
    // this row. In essence, a mapping is  provided between these indices and an
    // entry in ciscoBfdSessTable. The type is interface{} with range:
    // 1..4294967295.
    Ciscobfdsessipmapindex interface{}
}

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetFilter() yfilter.YFilter { return ciscobfdsessipmapentry.YFilter }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) SetFilter(yf yfilter.YFilter) { ciscobfdsessipmapentry.YFilter = yf }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetGoName(yname string) string {
    if yname == "ciscoBfdSessInterface" { return "Ciscobfdsessinterface" }
    if yname == "ciscoBfdSessAddrType" { return "Ciscobfdsessaddrtype" }
    if yname == "ciscoBfdSessAddr" { return "Ciscobfdsessaddr" }
    if yname == "ciscoBfdSessIpMapIndex" { return "Ciscobfdsessipmapindex" }
    return ""
}

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetSegmentPath() string {
    return "ciscoBfdSessIpMapEntry" + "[ciscoBfdSessInterface='" + fmt.Sprintf("%v", ciscobfdsessipmapentry.Ciscobfdsessinterface) + "']" + "[ciscoBfdSessAddrType='" + fmt.Sprintf("%v", ciscobfdsessipmapentry.Ciscobfdsessaddrtype) + "']" + "[ciscoBfdSessAddr='" + fmt.Sprintf("%v", ciscobfdsessipmapentry.Ciscobfdsessaddr) + "']"
}

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoBfdSessInterface"] = ciscobfdsessipmapentry.Ciscobfdsessinterface
    leafs["ciscoBfdSessAddrType"] = ciscobfdsessipmapentry.Ciscobfdsessaddrtype
    leafs["ciscoBfdSessAddr"] = ciscobfdsessipmapentry.Ciscobfdsessaddr
    leafs["ciscoBfdSessIpMapIndex"] = ciscobfdsessipmapentry.Ciscobfdsessipmapindex
    return leafs
}

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetYangName() string { return "ciscoBfdSessIpMapEntry" }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) SetParent(parent types.Entity) { ciscobfdsessipmapentry.parent = parent }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetParent() types.Entity { return ciscobfdsessipmapentry.parent }

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetParentYangName() string { return "ciscoBfdSessIpMapTable" }

