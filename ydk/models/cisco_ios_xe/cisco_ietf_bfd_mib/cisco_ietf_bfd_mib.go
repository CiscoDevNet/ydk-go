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
    EntityData types.CommonEntityData
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

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFBFDMIB.EntityData.YFilter = cISCOIETFBFDMIB.YFilter
    cISCOIETFBFDMIB.EntityData.YangName = "CISCO-IETF-BFD-MIB"
    cISCOIETFBFDMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFBFDMIB.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    cISCOIETFBFDMIB.EntityData.SegmentPath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB"
    cISCOIETFBFDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFBFDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFBFDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFBFDMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFBFDMIB.EntityData.Children["ciscoBfdScalarObjects"] = types.YChild{"Ciscobfdscalarobjects", &cISCOIETFBFDMIB.Ciscobfdscalarobjects}
    cISCOIETFBFDMIB.EntityData.Children["ciscoBfdSessTable"] = types.YChild{"Ciscobfdsesstable", &cISCOIETFBFDMIB.Ciscobfdsesstable}
    cISCOIETFBFDMIB.EntityData.Children["ciscoBfdSessMapTable"] = types.YChild{"Ciscobfdsessmaptable", &cISCOIETFBFDMIB.Ciscobfdsessmaptable}
    cISCOIETFBFDMIB.EntityData.Children["ciscoBfdSessDiscMapTable"] = types.YChild{"Ciscobfdsessdiscmaptable", &cISCOIETFBFDMIB.Ciscobfdsessdiscmaptable}
    cISCOIETFBFDMIB.EntityData.Children["ciscoBfdSessIpMapTable"] = types.YChild{"Ciscobfdsessipmaptable", &cISCOIETFBFDMIB.Ciscobfdsessipmaptable}
    cISCOIETFBFDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFBFDMIB.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdscalarobjects
type CISCOIETFBFDMIB_Ciscobfdscalarobjects struct {
    EntityData types.CommonEntityData
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

func (ciscobfdscalarobjects *CISCOIETFBFDMIB_Ciscobfdscalarobjects) GetEntityData() *types.CommonEntityData {
    ciscobfdscalarobjects.EntityData.YFilter = ciscobfdscalarobjects.YFilter
    ciscobfdscalarobjects.EntityData.YangName = "ciscoBfdScalarObjects"
    ciscobfdscalarobjects.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdscalarobjects.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscobfdscalarobjects.EntityData.SegmentPath = "ciscoBfdScalarObjects"
    ciscobfdscalarobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdscalarobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdscalarobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdscalarobjects.EntityData.Children = make(map[string]types.YChild)
    ciscobfdscalarobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscobfdscalarobjects.EntityData.Leafs["ciscoBfdAdminStatus"] = types.YLeaf{"Ciscobfdadminstatus", ciscobfdscalarobjects.Ciscobfdadminstatus}
    ciscobfdscalarobjects.EntityData.Leafs["ciscoBfdVersionNumber"] = types.YLeaf{"Ciscobfdversionnumber", ciscobfdscalarobjects.Ciscobfdversionnumber}
    ciscobfdscalarobjects.EntityData.Leafs["ciscoBfdSessNotificationsEnable"] = types.YLeaf{"Ciscobfdsessnotificationsenable", ciscobfdscalarobjects.Ciscobfdsessnotificationsenable}
    return &(ciscobfdscalarobjects.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus represents it on all interfaces.
type CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus string

const (
    CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus_enabled CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus = "enabled"

    CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus_disabled CISCOIETFBFDMIB_Ciscobfdscalarobjects_Ciscobfdadminstatus = "disabled"
)

// CISCOIETFBFDMIB_Ciscobfdsesstable
// The BFD Session Table describes the BFD sessions.
type CISCOIETFBFDMIB_Ciscobfdsesstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The BFD Session Entry describes BFD session. The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry.
    Ciscobfdsessentry []CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry
}

func (ciscobfdsesstable *CISCOIETFBFDMIB_Ciscobfdsesstable) GetEntityData() *types.CommonEntityData {
    ciscobfdsesstable.EntityData.YFilter = ciscobfdsesstable.YFilter
    ciscobfdsesstable.EntityData.YangName = "ciscoBfdSessTable"
    ciscobfdsesstable.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsesstable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscobfdsesstable.EntityData.SegmentPath = "ciscoBfdSessTable"
    ciscobfdsesstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsesstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsesstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsesstable.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsesstable.EntityData.Children["ciscoBfdSessEntry"] = types.YChild{"Ciscobfdsessentry", nil}
    for i := range ciscobfdsesstable.Ciscobfdsessentry {
        ciscobfdsesstable.EntityData.Children[types.GetSegmentPath(&ciscobfdsesstable.Ciscobfdsessentry[i])] = types.YChild{"Ciscobfdsessentry", &ciscobfdsesstable.Ciscobfdsessentry[i]}
    }
    ciscobfdsesstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscobfdsesstable.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry
// The BFD Session Entry describes BFD session.
type CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry struct {
    EntityData types.CommonEntityData
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

func (ciscobfdsessentry *CISCOIETFBFDMIB_Ciscobfdsesstable_Ciscobfdsessentry) GetEntityData() *types.CommonEntityData {
    ciscobfdsessentry.EntityData.YFilter = ciscobfdsessentry.YFilter
    ciscobfdsessentry.EntityData.YangName = "ciscoBfdSessEntry"
    ciscobfdsessentry.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsessentry.EntityData.ParentYangName = "ciscoBfdSessTable"
    ciscobfdsessentry.EntityData.SegmentPath = "ciscoBfdSessEntry" + "[ciscoBfdSessIndex='" + fmt.Sprintf("%v", ciscobfdsessentry.Ciscobfdsessindex) + "']"
    ciscobfdsessentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsessentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsessentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsessentry.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsessentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessIndex"] = types.YLeaf{"Ciscobfdsessindex", ciscobfdsessentry.Ciscobfdsessindex}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessApplicationId"] = types.YLeaf{"Ciscobfdsessapplicationid", ciscobfdsessentry.Ciscobfdsessapplicationid}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessDiscriminator"] = types.YLeaf{"Ciscobfdsessdiscriminator", ciscobfdsessentry.Ciscobfdsessdiscriminator}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessRemoteDiscr"] = types.YLeaf{"Ciscobfdsessremotediscr", ciscobfdsessentry.Ciscobfdsessremotediscr}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessUdpPort"] = types.YLeaf{"Ciscobfdsessudpport", ciscobfdsessentry.Ciscobfdsessudpport}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessState"] = types.YLeaf{"Ciscobfdsessstate", ciscobfdsessentry.Ciscobfdsessstate}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessRemoteHeardFlag"] = types.YLeaf{"Ciscobfdsessremoteheardflag", ciscobfdsessentry.Ciscobfdsessremoteheardflag}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessDiag"] = types.YLeaf{"Ciscobfdsessdiag", ciscobfdsessentry.Ciscobfdsessdiag}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessOperMode"] = types.YLeaf{"Ciscobfdsessopermode", ciscobfdsessentry.Ciscobfdsessopermode}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessDemandModeDesiredFlag"] = types.YLeaf{"Ciscobfdsessdemandmodedesiredflag", ciscobfdsessentry.Ciscobfdsessdemandmodedesiredflag}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessEchoFuncModeDesiredFlag"] = types.YLeaf{"Ciscobfdsessechofuncmodedesiredflag", ciscobfdsessentry.Ciscobfdsessechofuncmodedesiredflag}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessControlPlanIndepFlag"] = types.YLeaf{"Ciscobfdsesscontrolplanindepflag", ciscobfdsessentry.Ciscobfdsesscontrolplanindepflag}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessAddrType"] = types.YLeaf{"Ciscobfdsessaddrtype", ciscobfdsessentry.Ciscobfdsessaddrtype}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessAddr"] = types.YLeaf{"Ciscobfdsessaddr", ciscobfdsessentry.Ciscobfdsessaddr}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessDesiredMinTxInterval"] = types.YLeaf{"Ciscobfdsessdesiredmintxinterval", ciscobfdsessentry.Ciscobfdsessdesiredmintxinterval}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessReqMinRxInterval"] = types.YLeaf{"Ciscobfdsessreqminrxinterval", ciscobfdsessentry.Ciscobfdsessreqminrxinterval}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessReqMinEchoRxInterval"] = types.YLeaf{"Ciscobfdsessreqminechorxinterval", ciscobfdsessentry.Ciscobfdsessreqminechorxinterval}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessDetectMult"] = types.YLeaf{"Ciscobfdsessdetectmult", ciscobfdsessentry.Ciscobfdsessdetectmult}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessStorType"] = types.YLeaf{"Ciscobfdsessstortype", ciscobfdsessentry.Ciscobfdsessstortype}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessRowStatus"] = types.YLeaf{"Ciscobfdsessrowstatus", ciscobfdsessentry.Ciscobfdsessrowstatus}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessAuthPresFlag"] = types.YLeaf{"Ciscobfdsessauthpresflag", ciscobfdsessentry.Ciscobfdsessauthpresflag}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessAuthenticationType"] = types.YLeaf{"Ciscobfdsessauthenticationtype", ciscobfdsessentry.Ciscobfdsessauthenticationtype}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessVersionNumber"] = types.YLeaf{"Ciscobfdsessversionnumber", ciscobfdsessentry.Ciscobfdsessversionnumber}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessType"] = types.YLeaf{"Ciscobfdsesstype", ciscobfdsessentry.Ciscobfdsesstype}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessInterface"] = types.YLeaf{"Ciscobfdsessinterface", ciscobfdsessentry.Ciscobfdsessinterface}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfPktIn"] = types.YLeaf{"Ciscobfdsessperfpktin", ciscobfdsessentry.Ciscobfdsessperfpktin}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfPktOut"] = types.YLeaf{"Ciscobfdsessperfpktout", ciscobfdsessentry.Ciscobfdsessperfpktout}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessUpTime"] = types.YLeaf{"Ciscobfdsessuptime", ciscobfdsessentry.Ciscobfdsessuptime}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfLastSessDownTime"] = types.YLeaf{"Ciscobfdsessperflastsessdowntime", ciscobfdsessentry.Ciscobfdsessperflastsessdowntime}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfLastCommLostDiag"] = types.YLeaf{"Ciscobfdsessperflastcommlostdiag", ciscobfdsessentry.Ciscobfdsessperflastcommlostdiag}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfSessUpCount"] = types.YLeaf{"Ciscobfdsessperfsessupcount", ciscobfdsessentry.Ciscobfdsessperfsessupcount}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfDiscTime"] = types.YLeaf{"Ciscobfdsessperfdisctime", ciscobfdsessentry.Ciscobfdsessperfdisctime}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfPktInHC"] = types.YLeaf{"Ciscobfdsessperfpktinhc", ciscobfdsessentry.Ciscobfdsessperfpktinhc}
    ciscobfdsessentry.EntityData.Leafs["ciscoBfdSessPerfPktOutHC"] = types.YLeaf{"Ciscobfdsessperfpktouthc", ciscobfdsessentry.Ciscobfdsessperfpktouthc}
    return &(ciscobfdsessentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The BFD Session Entry describes BFD session that is mapped to this index.
    // The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry.
    Ciscobfdsessmapentry []CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry
}

func (ciscobfdsessmaptable *CISCOIETFBFDMIB_Ciscobfdsessmaptable) GetEntityData() *types.CommonEntityData {
    ciscobfdsessmaptable.EntityData.YFilter = ciscobfdsessmaptable.YFilter
    ciscobfdsessmaptable.EntityData.YangName = "ciscoBfdSessMapTable"
    ciscobfdsessmaptable.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsessmaptable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscobfdsessmaptable.EntityData.SegmentPath = "ciscoBfdSessMapTable"
    ciscobfdsessmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsessmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsessmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsessmaptable.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsessmaptable.EntityData.Children["ciscoBfdSessMapEntry"] = types.YChild{"Ciscobfdsessmapentry", nil}
    for i := range ciscobfdsessmaptable.Ciscobfdsessmapentry {
        ciscobfdsessmaptable.EntityData.Children[types.GetSegmentPath(&ciscobfdsessmaptable.Ciscobfdsessmapentry[i])] = types.YChild{"Ciscobfdsessmapentry", &ciscobfdsessmaptable.Ciscobfdsessmapentry[i]}
    }
    ciscobfdsessmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscobfdsessmaptable.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry
// The BFD Session Entry describes BFD session
// that is mapped to this index.
type CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry struct {
    EntityData types.CommonEntityData
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

func (ciscobfdsessmapentry *CISCOIETFBFDMIB_Ciscobfdsessmaptable_Ciscobfdsessmapentry) GetEntityData() *types.CommonEntityData {
    ciscobfdsessmapentry.EntityData.YFilter = ciscobfdsessmapentry.YFilter
    ciscobfdsessmapentry.EntityData.YangName = "ciscoBfdSessMapEntry"
    ciscobfdsessmapentry.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsessmapentry.EntityData.ParentYangName = "ciscoBfdSessMapTable"
    ciscobfdsessmapentry.EntityData.SegmentPath = "ciscoBfdSessMapEntry" + "[ciscoBfdSessApplicationId='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessapplicationid) + "']" + "[ciscoBfdSessDiscriminator='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessdiscriminator) + "']" + "[ciscoBfdSessAddrType='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessaddrtype) + "']" + "[ciscoBfdSessAddr='" + fmt.Sprintf("%v", ciscobfdsessmapentry.Ciscobfdsessaddr) + "']"
    ciscobfdsessmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsessmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsessmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsessmapentry.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsessmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscobfdsessmapentry.EntityData.Leafs["ciscoBfdSessApplicationId"] = types.YLeaf{"Ciscobfdsessapplicationid", ciscobfdsessmapentry.Ciscobfdsessapplicationid}
    ciscobfdsessmapentry.EntityData.Leafs["ciscoBfdSessDiscriminator"] = types.YLeaf{"Ciscobfdsessdiscriminator", ciscobfdsessmapentry.Ciscobfdsessdiscriminator}
    ciscobfdsessmapentry.EntityData.Leafs["ciscoBfdSessAddrType"] = types.YLeaf{"Ciscobfdsessaddrtype", ciscobfdsessmapentry.Ciscobfdsessaddrtype}
    ciscobfdsessmapentry.EntityData.Leafs["ciscoBfdSessAddr"] = types.YLeaf{"Ciscobfdsessaddr", ciscobfdsessmapentry.Ciscobfdsessaddr}
    ciscobfdsessmapentry.EntityData.Leafs["ciscoBfdSessMapBfdIndex"] = types.YLeaf{"Ciscobfdsessmapbfdindex", ciscobfdsessmapentry.Ciscobfdsessmapbfdindex}
    return &(ciscobfdsessmapentry.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable
// The BFD Session Discriminator Mapping Table maps a
// local discriminator value to associated BFD sessions'
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
type CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row contains a mapping between a local discriminator value to an entry
    // in ciscoBfdSessTable. The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry.
    Ciscobfdsessdiscmapentry []CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry
}

func (ciscobfdsessdiscmaptable *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable) GetEntityData() *types.CommonEntityData {
    ciscobfdsessdiscmaptable.EntityData.YFilter = ciscobfdsessdiscmaptable.YFilter
    ciscobfdsessdiscmaptable.EntityData.YangName = "ciscoBfdSessDiscMapTable"
    ciscobfdsessdiscmaptable.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsessdiscmaptable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscobfdsessdiscmaptable.EntityData.SegmentPath = "ciscoBfdSessDiscMapTable"
    ciscobfdsessdiscmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsessdiscmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsessdiscmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsessdiscmaptable.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsessdiscmaptable.EntityData.Children["ciscoBfdSessDiscMapEntry"] = types.YChild{"Ciscobfdsessdiscmapentry", nil}
    for i := range ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry {
        ciscobfdsessdiscmaptable.EntityData.Children[types.GetSegmentPath(&ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry[i])] = types.YChild{"Ciscobfdsessdiscmapentry", &ciscobfdsessdiscmaptable.Ciscobfdsessdiscmapentry[i]}
    }
    ciscobfdsessdiscmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscobfdsessdiscmaptable.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry
// Each row contains a mapping between a local discriminator
// value to an entry in ciscoBfdSessTable.
type CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry struct {
    EntityData types.CommonEntityData
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

func (ciscobfdsessdiscmapentry *CISCOIETFBFDMIB_Ciscobfdsessdiscmaptable_Ciscobfdsessdiscmapentry) GetEntityData() *types.CommonEntityData {
    ciscobfdsessdiscmapentry.EntityData.YFilter = ciscobfdsessdiscmapentry.YFilter
    ciscobfdsessdiscmapentry.EntityData.YangName = "ciscoBfdSessDiscMapEntry"
    ciscobfdsessdiscmapentry.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsessdiscmapentry.EntityData.ParentYangName = "ciscoBfdSessDiscMapTable"
    ciscobfdsessdiscmapentry.EntityData.SegmentPath = "ciscoBfdSessDiscMapEntry" + "[ciscoBfdSessDiscriminator='" + fmt.Sprintf("%v", ciscobfdsessdiscmapentry.Ciscobfdsessdiscriminator) + "']"
    ciscobfdsessdiscmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsessdiscmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsessdiscmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsessdiscmapentry.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsessdiscmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscobfdsessdiscmapentry.EntityData.Leafs["ciscoBfdSessDiscriminator"] = types.YLeaf{"Ciscobfdsessdiscriminator", ciscobfdsessdiscmapentry.Ciscobfdsessdiscriminator}
    ciscobfdsessdiscmapentry.EntityData.Leafs["ciscoBfdSessDiscMapIndex"] = types.YLeaf{"Ciscobfdsessdiscmapindex", ciscobfdsessdiscmapentry.Ciscobfdsessdiscmapindex}
    return &(ciscobfdsessdiscmapentry.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdsessipmaptable
// The BFD Session IP Mapping Table maps given
// ciscoBfdSessInterface, ciscoBfdSessAddrType, and
// ciscoBbfdSessAddr to an associated BFD sessions'
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
// This table SHOULD contains those BFD sessions are
// of IP type: singleHop(1) and multiHop(2).
type CISCOIETFBFDMIB_Ciscobfdsessipmaptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row contains a mapping between ciscoBfdSessInterface,
    // ciscoBfdSessAddrType and ciscoBfdSessAddr values to an  entry in
    // ciscoBfdSessTable. The type is slice of
    // CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry.
    Ciscobfdsessipmapentry []CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry
}

func (ciscobfdsessipmaptable *CISCOIETFBFDMIB_Ciscobfdsessipmaptable) GetEntityData() *types.CommonEntityData {
    ciscobfdsessipmaptable.EntityData.YFilter = ciscobfdsessipmaptable.YFilter
    ciscobfdsessipmaptable.EntityData.YangName = "ciscoBfdSessIpMapTable"
    ciscobfdsessipmaptable.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsessipmaptable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscobfdsessipmaptable.EntityData.SegmentPath = "ciscoBfdSessIpMapTable"
    ciscobfdsessipmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsessipmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsessipmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsessipmaptable.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsessipmaptable.EntityData.Children["ciscoBfdSessIpMapEntry"] = types.YChild{"Ciscobfdsessipmapentry", nil}
    for i := range ciscobfdsessipmaptable.Ciscobfdsessipmapentry {
        ciscobfdsessipmaptable.EntityData.Children[types.GetSegmentPath(&ciscobfdsessipmaptable.Ciscobfdsessipmapentry[i])] = types.YChild{"Ciscobfdsessipmapentry", &ciscobfdsessipmaptable.Ciscobfdsessipmapentry[i]}
    }
    ciscobfdsessipmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscobfdsessipmaptable.EntityData)
}

// CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry
// Each row contains a mapping between ciscoBfdSessInterface,
// ciscoBfdSessAddrType and ciscoBfdSessAddr values to an 
// entry in ciscoBfdSessTable.
type CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry struct {
    EntityData types.CommonEntityData
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

func (ciscobfdsessipmapentry *CISCOIETFBFDMIB_Ciscobfdsessipmaptable_Ciscobfdsessipmapentry) GetEntityData() *types.CommonEntityData {
    ciscobfdsessipmapentry.EntityData.YFilter = ciscobfdsessipmapentry.YFilter
    ciscobfdsessipmapentry.EntityData.YangName = "ciscoBfdSessIpMapEntry"
    ciscobfdsessipmapentry.EntityData.BundleName = "cisco_ios_xe"
    ciscobfdsessipmapentry.EntityData.ParentYangName = "ciscoBfdSessIpMapTable"
    ciscobfdsessipmapentry.EntityData.SegmentPath = "ciscoBfdSessIpMapEntry" + "[ciscoBfdSessInterface='" + fmt.Sprintf("%v", ciscobfdsessipmapentry.Ciscobfdsessinterface) + "']" + "[ciscoBfdSessAddrType='" + fmt.Sprintf("%v", ciscobfdsessipmapentry.Ciscobfdsessaddrtype) + "']" + "[ciscoBfdSessAddr='" + fmt.Sprintf("%v", ciscobfdsessipmapentry.Ciscobfdsessaddr) + "']"
    ciscobfdsessipmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscobfdsessipmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscobfdsessipmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscobfdsessipmapentry.EntityData.Children = make(map[string]types.YChild)
    ciscobfdsessipmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscobfdsessipmapentry.EntityData.Leafs["ciscoBfdSessInterface"] = types.YLeaf{"Ciscobfdsessinterface", ciscobfdsessipmapentry.Ciscobfdsessinterface}
    ciscobfdsessipmapentry.EntityData.Leafs["ciscoBfdSessAddrType"] = types.YLeaf{"Ciscobfdsessaddrtype", ciscobfdsessipmapentry.Ciscobfdsessaddrtype}
    ciscobfdsessipmapentry.EntityData.Leafs["ciscoBfdSessAddr"] = types.YLeaf{"Ciscobfdsessaddr", ciscobfdsessipmapentry.Ciscobfdsessaddr}
    ciscobfdsessipmapentry.EntityData.Leafs["ciscoBfdSessIpMapIndex"] = types.YLeaf{"Ciscobfdsessipmapindex", ciscobfdsessipmapentry.Ciscobfdsessipmapindex}
    return &(ciscobfdsessipmapentry.EntityData)
}

