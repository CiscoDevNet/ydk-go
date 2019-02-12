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

    
    CiscoBfdScalarObjects CISCOIETFBFDMIB_CiscoBfdScalarObjects

    // The BFD Session Table describes the BFD sessions.
    CiscoBfdSessTable CISCOIETFBFDMIB_CiscoBfdSessTable

    // The BFD Session Mapping Table maps the complex indexing of the BFD sessions
    // to the flat  CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
    CiscoBfdSessMapTable CISCOIETFBFDMIB_CiscoBfdSessMapTable

    // The BFD Session Discriminator Mapping Table maps a local discriminator
    // value to associated BFD sessions' CiscoBfdSessIndexTC used in the
    // ciscoBfdSessTable.
    CiscoBfdSessDiscMapTable CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable

    // The BFD Session IP Mapping Table maps given ciscoBfdSessInterface,
    // ciscoBfdSessAddrType, and ciscoBbfdSessAddr to an associated BFD sessions'
    // CiscoBfdSessIndexTC used in the ciscoBfdSessTable. This table SHOULD
    // contains those BFD sessions are of IP type: singleHop(1) and multiHop(2).
    CiscoBfdSessIpMapTable CISCOIETFBFDMIB_CiscoBfdSessIpMapTable
}

func (cISCOIETFBFDMIB *CISCOIETFBFDMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFBFDMIB.EntityData.YFilter = cISCOIETFBFDMIB.YFilter
    cISCOIETFBFDMIB.EntityData.YangName = "CISCO-IETF-BFD-MIB"
    cISCOIETFBFDMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFBFDMIB.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    cISCOIETFBFDMIB.EntityData.SegmentPath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB"
    cISCOIETFBFDMIB.EntityData.AbsolutePath = cISCOIETFBFDMIB.EntityData.SegmentPath
    cISCOIETFBFDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFBFDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFBFDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFBFDMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFBFDMIB.EntityData.Children.Append("ciscoBfdScalarObjects", types.YChild{"CiscoBfdScalarObjects", &cISCOIETFBFDMIB.CiscoBfdScalarObjects})
    cISCOIETFBFDMIB.EntityData.Children.Append("ciscoBfdSessTable", types.YChild{"CiscoBfdSessTable", &cISCOIETFBFDMIB.CiscoBfdSessTable})
    cISCOIETFBFDMIB.EntityData.Children.Append("ciscoBfdSessMapTable", types.YChild{"CiscoBfdSessMapTable", &cISCOIETFBFDMIB.CiscoBfdSessMapTable})
    cISCOIETFBFDMIB.EntityData.Children.Append("ciscoBfdSessDiscMapTable", types.YChild{"CiscoBfdSessDiscMapTable", &cISCOIETFBFDMIB.CiscoBfdSessDiscMapTable})
    cISCOIETFBFDMIB.EntityData.Children.Append("ciscoBfdSessIpMapTable", types.YChild{"CiscoBfdSessIpMapTable", &cISCOIETFBFDMIB.CiscoBfdSessIpMapTable})
    cISCOIETFBFDMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFBFDMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFBFDMIB.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdScalarObjects
type CISCOIETFBFDMIB_CiscoBfdScalarObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The global administrative status of BFD in this router. The value 'enabled'
    // denotes that the BFD Process is  active on at least one interface;
    // 'disabled' disables   it on all interfaces. The type is
    // CiscoBfdAdminStatus.
    CiscoBfdAdminStatus interface{}

    // The current version number of the BFD protocol. The type is interface{}
    // with range: 0..4294967295.
    CiscoBfdVersionNumber interface{}

    // If this object is set to true(1), then it enables the emission of
    // ciscoBfdSessUp and ciscoBfdSessDown  notifications; otherwise these
    // notifications are not  emitted. The type is bool.
    CiscoBfdSessNotificationsEnable interface{}
}

func (ciscoBfdScalarObjects *CISCOIETFBFDMIB_CiscoBfdScalarObjects) GetEntityData() *types.CommonEntityData {
    ciscoBfdScalarObjects.EntityData.YFilter = ciscoBfdScalarObjects.YFilter
    ciscoBfdScalarObjects.EntityData.YangName = "ciscoBfdScalarObjects"
    ciscoBfdScalarObjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdScalarObjects.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscoBfdScalarObjects.EntityData.SegmentPath = "ciscoBfdScalarObjects"
    ciscoBfdScalarObjects.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/" + ciscoBfdScalarObjects.EntityData.SegmentPath
    ciscoBfdScalarObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdScalarObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdScalarObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdScalarObjects.EntityData.Children = types.NewOrderedMap()
    ciscoBfdScalarObjects.EntityData.Leafs = types.NewOrderedMap()
    ciscoBfdScalarObjects.EntityData.Leafs.Append("ciscoBfdAdminStatus", types.YLeaf{"CiscoBfdAdminStatus", ciscoBfdScalarObjects.CiscoBfdAdminStatus})
    ciscoBfdScalarObjects.EntityData.Leafs.Append("ciscoBfdVersionNumber", types.YLeaf{"CiscoBfdVersionNumber", ciscoBfdScalarObjects.CiscoBfdVersionNumber})
    ciscoBfdScalarObjects.EntityData.Leafs.Append("ciscoBfdSessNotificationsEnable", types.YLeaf{"CiscoBfdSessNotificationsEnable", ciscoBfdScalarObjects.CiscoBfdSessNotificationsEnable})

    ciscoBfdScalarObjects.EntityData.YListKeys = []string {}

    return &(ciscoBfdScalarObjects.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdScalarObjects_CiscoBfdAdminStatus represents it on all interfaces.
type CISCOIETFBFDMIB_CiscoBfdScalarObjects_CiscoBfdAdminStatus string

const (
    CISCOIETFBFDMIB_CiscoBfdScalarObjects_CiscoBfdAdminStatus_enabled CISCOIETFBFDMIB_CiscoBfdScalarObjects_CiscoBfdAdminStatus = "enabled"

    CISCOIETFBFDMIB_CiscoBfdScalarObjects_CiscoBfdAdminStatus_disabled CISCOIETFBFDMIB_CiscoBfdScalarObjects_CiscoBfdAdminStatus = "disabled"
)

// CISCOIETFBFDMIB_CiscoBfdSessTable
// The BFD Session Table describes the BFD sessions.
type CISCOIETFBFDMIB_CiscoBfdSessTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The BFD Session Entry describes BFD session. The type is slice of
    // CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry.
    CiscoBfdSessEntry []*CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry
}

func (ciscoBfdSessTable *CISCOIETFBFDMIB_CiscoBfdSessTable) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessTable.EntityData.YFilter = ciscoBfdSessTable.YFilter
    ciscoBfdSessTable.EntityData.YangName = "ciscoBfdSessTable"
    ciscoBfdSessTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessTable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscoBfdSessTable.EntityData.SegmentPath = "ciscoBfdSessTable"
    ciscoBfdSessTable.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/" + ciscoBfdSessTable.EntityData.SegmentPath
    ciscoBfdSessTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessTable.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessTable.EntityData.Children.Append("ciscoBfdSessEntry", types.YChild{"CiscoBfdSessEntry", nil})
    for i := range ciscoBfdSessTable.CiscoBfdSessEntry {
        ciscoBfdSessTable.EntityData.Children.Append(types.GetSegmentPath(ciscoBfdSessTable.CiscoBfdSessEntry[i]), types.YChild{"CiscoBfdSessEntry", ciscoBfdSessTable.CiscoBfdSessEntry[i]})
    }
    ciscoBfdSessTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoBfdSessTable.EntityData.YListKeys = []string {}

    return &(ciscoBfdSessTable.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry
// The BFD Session Entry describes BFD session.
type CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object contains an index used to represent a
    // unique BFD session on this device. The type is interface{} with range:
    // 1..4294967295.
    CiscoBfdSessIndex interface{}

    // This object contains an index used to indicate a local application which
    // owns or maintains this  BFD session. For instance, the MPLS VPN process may
    // maintain a subset of the total number of BFD  sessions.  This application
    // ID provides a convenient  way to segregate sessions by the applications
    // which  maintain them. The type is interface{} with range: 0..4294967295.
    CiscoBfdSessApplicationId interface{}

    // This object specifies the local discriminator for this BFD session, used to
    // uniquely identify it. The type is interface{} with range: 1..4294967295.
    CiscoBfdSessDiscriminator interface{}

    // This object specifies the session discriminator chosen by the remote system
    // for this BFD session. The type is interface{} with range: 0..4294967295.
    CiscoBfdSessRemoteDiscr interface{}

    // The destination UDP Port for BFD. The default value is the well-known value
    // for this port. BFD State failing(5) is only applicable if this BFD session
    // is running version 0. The type is interface{} with range: 0..65535.
    CiscoBfdSessUdpPort interface{}

    // The perceived state of the BFD session. The type is CiscoBfdSessState.
    CiscoBfdSessState interface{}

    // This object specifies status of BFD packet reception from the remote
    // system. Specifically, it is set to true(1) if  the local system is actively
    // receiving BFD packets from the   remote system, and is set to false(0) if
    // the local system   has not received BFD packets recently (within the
    // detection   time) or if the local system is attempting to tear down  the
    // BFD session. The type is bool.
    CiscoBfdSessRemoteHeardFlag interface{}

    // A diagnostic code specifying the local system's reason for the last
    // transition of the session from up(1)   to some other state. The type is
    // CiscoBfdDiag.
    CiscoBfdSessDiag interface{}

    // This object specifies current operating mode that BFD session is operating
    // in. The type is CiscoBfdSessOperMode.
    CiscoBfdSessOperMode interface{}

    // This object indicates that the local system's desire to use Demand mode.
    // Specifically, it is set   to true(1) if the local system wishes to use  
    // Demand mode or false(0) if not. The type is bool.
    CiscoBfdSessDemandModeDesiredFlag interface{}

    // This object indicates that the local system's desire to use Echo mode.
    // Specifically, it is set   to true(1) if the local system wishes to use  
    // Echo mode or false(0) if not. The type is bool.
    CiscoBfdSessEchoFuncModeDesiredFlag interface{}

    // This object indicates that the local system's ability to continue to
    // function through a disruption of   the control plane. Specifically, it is
    // set   to true(1) if the local system BFD implementation is  independent of
    // the control plane. Otherwise, the   value is set to false(0). The type is
    // bool.
    CiscoBfdSessControlPlanIndepFlag interface{}

    // This object specifies IP address type of the neighboring IP address which
    // is being monitored with this BFD session.  Only values unknown(0), ipv4(1)
    // or ipv6(2)  have to be supported.    A value of unknown(0) is allowed only
    // when   the outgoing interface is of type point-to-point, or  when the BFD
    // session is not associated with a specific   interface.   If any other
    // unsupported values are attempted in a set  operation, the agent MUST return
    // an inconsistentValue   error. The type is InetAddressType.
    CiscoBfdSessAddrType interface{}

    // This object specifies the neighboring IP address which is being monitored
    // with this BFD session. It can also be used to enabled BFD on a specific  
    // interface. The value is set to zero when BFD session is not   associated
    // with a specific interface. The type is string with length: 0..255.
    CiscoBfdSessAddr interface{}

    // This object specifies the minimum interval, in microseconds, that the local
    // system would like to use when       transmitting BFD Control packets. The
    // type is interface{} with range: 0..4294967295.
    CiscoBfdSessDesiredMinTxInterval interface{}

    // This object specifies the minimum interval, in microseconds, between
    // received  BFD Control packets the   local system is capable of supporting.
    // The type is interface{} with range: 0..4294967295.
    CiscoBfdSessReqMinRxInterval interface{}

    // This object specifies the minimum interval, in microseconds, between
    // received BFD Echo packets that this  system is capable of supporting. The
    // type is interface{} with range: 0..4294967295.
    CiscoBfdSessReqMinEchoRxInterval interface{}

    // This object specifies the Detect time multiplier. The type is interface{}
    // with range: 0..4294967295.
    CiscoBfdSessDetectMult interface{}

    // This variable indicates the storage type for this object. Conceptual rows
    // having the value   'permanent' need not allow write-access to any  
    // columnar objects in the row. The type is StorageType.
    CiscoBfdSessStorType interface{}

    // This variable is used to create, modify, and/or delete a row in this table.
    // When a row in this  table has a row in the active(1) state, no   objects in
    // this row can be modified except the  ciscoBfdSessRowStatus and
    // ciscoBfdSessStorageType. The type is RowStatus.
    CiscoBfdSessRowStatus interface{}

    // This object indicates that the local system's desire to use Authentication.
    // Specifically, it is set   to true(1) if the local system wishes the session
    // to be authenticated or false(0) if not. The type is bool.
    CiscoBfdSessAuthPresFlag interface{}

    // The Authentication Type used for this BFD session. This field is valid only
    // when the Authentication Present bit is set. The type is
    // CiscoBfdSessAuthenticationType.
    CiscoBfdSessAuthenticationType interface{}

    // The version number of the BFD protocol that this session is running in. The
    // type is interface{} with range: 0..4294967295.
    CiscoBfdSessVersionNumber interface{}

    // The type of this BFD session. The type is CiscoBfdSessType.
    CiscoBfdSessType interface{}

    // This object contains an interface index used to indicate the interface
    // which this BFD session is running on. The type is interface{} with range:
    // 1..2147483647.
    CiscoBfdSessInterface interface{}

    // The total number of BFD messages received for this BFD session. The type is
    // interface{} with range: 0..4294967295.
    CiscoBfdSessPerfPktIn interface{}

    // The total number of BFD messages sent for this BFD session. The type is
    // interface{} with range: 0..4294967295.
    CiscoBfdSessPerfPktOut interface{}

    // The value of sysUpTime on the most recent occasion at which the session
    // came up. If no such up event exists this object  contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    CiscoBfdSessUpTime interface{}

    // The value of sysUpTime on the most recent occasion at which the last time
    // communication was lost with the neighbor. If   no such down event exist
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    CiscoBfdSessPerfLastSessDownTime interface{}

    // The BFD diag code for the last time communication was lost with the
    // neighbor. If no such down event exists this object   contains a zero value.
    // The type is CiscoBfdDiag.
    CiscoBfdSessPerfLastCommLostDiag interface{}

    // The number of times this session has gone into the Up state since the
    // router last rebooted. The type is interface{} with range: 0..4294967295.
    CiscoBfdSessPerfSessUpCount interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of the session counters suffered  a discontinuity.    The relevant counters
    // are the specific instances associated   with this BFD session of any
    // Counter32 object contained in  the ciscoBfdSessPerfTable. If no such
    // discontinuities have occurred   since the last re-initialization of the
    // local management  subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    CiscoBfdSessPerfDiscTime interface{}

    // This value represents the total number of BFD messages received for this
    // BFD session. It MUST be equal to the  least significant 32 bits of
    // ciscoBfdSessPerfPktIn  if ciscoBfdSessPerfPktInHC is supported according to
    // the rules spelled out in RFC2863. The type is interface{} with range:
    // 0..18446744073709551615.
    CiscoBfdSessPerfPktInHC interface{}

    // This value represents the total number of total number of BFD messages
    // transmitted for this   BFD session. It MUST be equal to the  least
    // significant 32 bits of ciscoBfdSessPerfPktIn  if ciscoBfdSessPerfPktOutHC
    // is supported according to  the rules spelled out in RFC2863. The type is
    // interface{} with range: 0..18446744073709551615.
    CiscoBfdSessPerfPktOutHC interface{}
}

func (ciscoBfdSessEntry *CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessEntry.EntityData.YFilter = ciscoBfdSessEntry.YFilter
    ciscoBfdSessEntry.EntityData.YangName = "ciscoBfdSessEntry"
    ciscoBfdSessEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessEntry.EntityData.ParentYangName = "ciscoBfdSessTable"
    ciscoBfdSessEntry.EntityData.SegmentPath = "ciscoBfdSessEntry" + types.AddKeyToken(ciscoBfdSessEntry.CiscoBfdSessIndex, "ciscoBfdSessIndex")
    ciscoBfdSessEntry.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/ciscoBfdSessTable/" + ciscoBfdSessEntry.EntityData.SegmentPath
    ciscoBfdSessEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessEntry.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessIndex", types.YLeaf{"CiscoBfdSessIndex", ciscoBfdSessEntry.CiscoBfdSessIndex})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessApplicationId", types.YLeaf{"CiscoBfdSessApplicationId", ciscoBfdSessEntry.CiscoBfdSessApplicationId})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessDiscriminator", types.YLeaf{"CiscoBfdSessDiscriminator", ciscoBfdSessEntry.CiscoBfdSessDiscriminator})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessRemoteDiscr", types.YLeaf{"CiscoBfdSessRemoteDiscr", ciscoBfdSessEntry.CiscoBfdSessRemoteDiscr})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessUdpPort", types.YLeaf{"CiscoBfdSessUdpPort", ciscoBfdSessEntry.CiscoBfdSessUdpPort})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessState", types.YLeaf{"CiscoBfdSessState", ciscoBfdSessEntry.CiscoBfdSessState})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessRemoteHeardFlag", types.YLeaf{"CiscoBfdSessRemoteHeardFlag", ciscoBfdSessEntry.CiscoBfdSessRemoteHeardFlag})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessDiag", types.YLeaf{"CiscoBfdSessDiag", ciscoBfdSessEntry.CiscoBfdSessDiag})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessOperMode", types.YLeaf{"CiscoBfdSessOperMode", ciscoBfdSessEntry.CiscoBfdSessOperMode})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessDemandModeDesiredFlag", types.YLeaf{"CiscoBfdSessDemandModeDesiredFlag", ciscoBfdSessEntry.CiscoBfdSessDemandModeDesiredFlag})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessEchoFuncModeDesiredFlag", types.YLeaf{"CiscoBfdSessEchoFuncModeDesiredFlag", ciscoBfdSessEntry.CiscoBfdSessEchoFuncModeDesiredFlag})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessControlPlanIndepFlag", types.YLeaf{"CiscoBfdSessControlPlanIndepFlag", ciscoBfdSessEntry.CiscoBfdSessControlPlanIndepFlag})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessAddrType", types.YLeaf{"CiscoBfdSessAddrType", ciscoBfdSessEntry.CiscoBfdSessAddrType})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessAddr", types.YLeaf{"CiscoBfdSessAddr", ciscoBfdSessEntry.CiscoBfdSessAddr})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessDesiredMinTxInterval", types.YLeaf{"CiscoBfdSessDesiredMinTxInterval", ciscoBfdSessEntry.CiscoBfdSessDesiredMinTxInterval})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessReqMinRxInterval", types.YLeaf{"CiscoBfdSessReqMinRxInterval", ciscoBfdSessEntry.CiscoBfdSessReqMinRxInterval})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessReqMinEchoRxInterval", types.YLeaf{"CiscoBfdSessReqMinEchoRxInterval", ciscoBfdSessEntry.CiscoBfdSessReqMinEchoRxInterval})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessDetectMult", types.YLeaf{"CiscoBfdSessDetectMult", ciscoBfdSessEntry.CiscoBfdSessDetectMult})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessStorType", types.YLeaf{"CiscoBfdSessStorType", ciscoBfdSessEntry.CiscoBfdSessStorType})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessRowStatus", types.YLeaf{"CiscoBfdSessRowStatus", ciscoBfdSessEntry.CiscoBfdSessRowStatus})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessAuthPresFlag", types.YLeaf{"CiscoBfdSessAuthPresFlag", ciscoBfdSessEntry.CiscoBfdSessAuthPresFlag})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessAuthenticationType", types.YLeaf{"CiscoBfdSessAuthenticationType", ciscoBfdSessEntry.CiscoBfdSessAuthenticationType})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessVersionNumber", types.YLeaf{"CiscoBfdSessVersionNumber", ciscoBfdSessEntry.CiscoBfdSessVersionNumber})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessType", types.YLeaf{"CiscoBfdSessType", ciscoBfdSessEntry.CiscoBfdSessType})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessInterface", types.YLeaf{"CiscoBfdSessInterface", ciscoBfdSessEntry.CiscoBfdSessInterface})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfPktIn", types.YLeaf{"CiscoBfdSessPerfPktIn", ciscoBfdSessEntry.CiscoBfdSessPerfPktIn})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfPktOut", types.YLeaf{"CiscoBfdSessPerfPktOut", ciscoBfdSessEntry.CiscoBfdSessPerfPktOut})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessUpTime", types.YLeaf{"CiscoBfdSessUpTime", ciscoBfdSessEntry.CiscoBfdSessUpTime})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfLastSessDownTime", types.YLeaf{"CiscoBfdSessPerfLastSessDownTime", ciscoBfdSessEntry.CiscoBfdSessPerfLastSessDownTime})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfLastCommLostDiag", types.YLeaf{"CiscoBfdSessPerfLastCommLostDiag", ciscoBfdSessEntry.CiscoBfdSessPerfLastCommLostDiag})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfSessUpCount", types.YLeaf{"CiscoBfdSessPerfSessUpCount", ciscoBfdSessEntry.CiscoBfdSessPerfSessUpCount})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfDiscTime", types.YLeaf{"CiscoBfdSessPerfDiscTime", ciscoBfdSessEntry.CiscoBfdSessPerfDiscTime})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfPktInHC", types.YLeaf{"CiscoBfdSessPerfPktInHC", ciscoBfdSessEntry.CiscoBfdSessPerfPktInHC})
    ciscoBfdSessEntry.EntityData.Leafs.Append("ciscoBfdSessPerfPktOutHC", types.YLeaf{"CiscoBfdSessPerfPktOutHC", ciscoBfdSessEntry.CiscoBfdSessPerfPktOutHC})

    ciscoBfdSessEntry.EntityData.YListKeys = []string {"CiscoBfdSessIndex"}

    return &(ciscoBfdSessEntry.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType represents field is valid only when the Authentication Present bit is set
type CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType string

const (
    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType_simplePassword CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType = "simplePassword"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType_keyedMD5 CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType = "keyedMD5"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType_meticulousKeyedMD5 CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType = "meticulousKeyedMD5"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType_keyedSHA1 CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType = "keyedSHA1"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType_meticulousKeyedSHA1 CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAuthenticationType = "meticulousKeyedSHA1"
)

// CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode represents session is operating in.
type CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode string

const (
    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode_asyncModeWEchoFun CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode = "asyncModeWEchoFun"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode_asynchModeWOEchoFun CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode = "asynchModeWOEchoFun"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode_demandModeWEchoFunction CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode = "demandModeWEchoFunction"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode_demandModeWOEchoFunction CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessOperMode = "demandModeWOEchoFunction"
)

// CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState represents The perceived state of the BFD session.
type CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState string

const (
    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState_adminDown CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState = "adminDown"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState_down CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState = "down"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState_init CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState = "init"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState_up CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState = "up"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState_failing CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessState = "failing"
)

// CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessType represents The type of this BFD session.
type CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessType string

const (
    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessType_singleHop CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessType = "singleHop"

    CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessType_multiHop CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessType = "multiHop"
)

// CISCOIETFBFDMIB_CiscoBfdSessMapTable
// The BFD Session Mapping Table maps the complex
// indexing of the BFD sessions to the flat 
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
type CISCOIETFBFDMIB_CiscoBfdSessMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The BFD Session Entry describes BFD session that is mapped to this index.
    // The type is slice of
    // CISCOIETFBFDMIB_CiscoBfdSessMapTable_CiscoBfdSessMapEntry.
    CiscoBfdSessMapEntry []*CISCOIETFBFDMIB_CiscoBfdSessMapTable_CiscoBfdSessMapEntry
}

func (ciscoBfdSessMapTable *CISCOIETFBFDMIB_CiscoBfdSessMapTable) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessMapTable.EntityData.YFilter = ciscoBfdSessMapTable.YFilter
    ciscoBfdSessMapTable.EntityData.YangName = "ciscoBfdSessMapTable"
    ciscoBfdSessMapTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessMapTable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscoBfdSessMapTable.EntityData.SegmentPath = "ciscoBfdSessMapTable"
    ciscoBfdSessMapTable.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/" + ciscoBfdSessMapTable.EntityData.SegmentPath
    ciscoBfdSessMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessMapTable.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessMapTable.EntityData.Children.Append("ciscoBfdSessMapEntry", types.YChild{"CiscoBfdSessMapEntry", nil})
    for i := range ciscoBfdSessMapTable.CiscoBfdSessMapEntry {
        ciscoBfdSessMapTable.EntityData.Children.Append(types.GetSegmentPath(ciscoBfdSessMapTable.CiscoBfdSessMapEntry[i]), types.YChild{"CiscoBfdSessMapEntry", ciscoBfdSessMapTable.CiscoBfdSessMapEntry[i]})
    }
    ciscoBfdSessMapTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoBfdSessMapTable.EntityData.YListKeys = []string {}

    return &(ciscoBfdSessMapTable.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdSessMapTable_CiscoBfdSessMapEntry
// The BFD Session Entry describes BFD session
// that is mapped to this index.
type CISCOIETFBFDMIB_CiscoBfdSessMapTable_CiscoBfdSessMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessApplicationId
    CiscoBfdSessApplicationId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessDiscriminator
    CiscoBfdSessDiscriminator interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAddrType
    CiscoBfdSessAddrType interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAddr
    CiscoBfdSessAddr interface{}

    // This object indicates the CiscoBfdSessIndexTC referred to by the indices of
    // this row. In essence, a mapping is  provided between these indices and the
    // ciscoBfdSessTable. The type is interface{} with range: 1..4294967295.
    CiscoBfdSessMapBfdIndex interface{}
}

func (ciscoBfdSessMapEntry *CISCOIETFBFDMIB_CiscoBfdSessMapTable_CiscoBfdSessMapEntry) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessMapEntry.EntityData.YFilter = ciscoBfdSessMapEntry.YFilter
    ciscoBfdSessMapEntry.EntityData.YangName = "ciscoBfdSessMapEntry"
    ciscoBfdSessMapEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessMapEntry.EntityData.ParentYangName = "ciscoBfdSessMapTable"
    ciscoBfdSessMapEntry.EntityData.SegmentPath = "ciscoBfdSessMapEntry" + types.AddKeyToken(ciscoBfdSessMapEntry.CiscoBfdSessApplicationId, "ciscoBfdSessApplicationId") + types.AddKeyToken(ciscoBfdSessMapEntry.CiscoBfdSessDiscriminator, "ciscoBfdSessDiscriminator") + types.AddKeyToken(ciscoBfdSessMapEntry.CiscoBfdSessAddrType, "ciscoBfdSessAddrType") + types.AddKeyToken(ciscoBfdSessMapEntry.CiscoBfdSessAddr, "ciscoBfdSessAddr")
    ciscoBfdSessMapEntry.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/ciscoBfdSessMapTable/" + ciscoBfdSessMapEntry.EntityData.SegmentPath
    ciscoBfdSessMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessMapEntry.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessMapEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoBfdSessMapEntry.EntityData.Leafs.Append("ciscoBfdSessApplicationId", types.YLeaf{"CiscoBfdSessApplicationId", ciscoBfdSessMapEntry.CiscoBfdSessApplicationId})
    ciscoBfdSessMapEntry.EntityData.Leafs.Append("ciscoBfdSessDiscriminator", types.YLeaf{"CiscoBfdSessDiscriminator", ciscoBfdSessMapEntry.CiscoBfdSessDiscriminator})
    ciscoBfdSessMapEntry.EntityData.Leafs.Append("ciscoBfdSessAddrType", types.YLeaf{"CiscoBfdSessAddrType", ciscoBfdSessMapEntry.CiscoBfdSessAddrType})
    ciscoBfdSessMapEntry.EntityData.Leafs.Append("ciscoBfdSessAddr", types.YLeaf{"CiscoBfdSessAddr", ciscoBfdSessMapEntry.CiscoBfdSessAddr})
    ciscoBfdSessMapEntry.EntityData.Leafs.Append("ciscoBfdSessMapBfdIndex", types.YLeaf{"CiscoBfdSessMapBfdIndex", ciscoBfdSessMapEntry.CiscoBfdSessMapBfdIndex})

    ciscoBfdSessMapEntry.EntityData.YListKeys = []string {"CiscoBfdSessApplicationId", "CiscoBfdSessDiscriminator", "CiscoBfdSessAddrType", "CiscoBfdSessAddr"}

    return &(ciscoBfdSessMapEntry.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable
// The BFD Session Discriminator Mapping Table maps a
// local discriminator value to associated BFD sessions'
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
type CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row contains a mapping between a local discriminator value to an entry
    // in ciscoBfdSessTable. The type is slice of
    // CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable_CiscoBfdSessDiscMapEntry.
    CiscoBfdSessDiscMapEntry []*CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable_CiscoBfdSessDiscMapEntry
}

func (ciscoBfdSessDiscMapTable *CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessDiscMapTable.EntityData.YFilter = ciscoBfdSessDiscMapTable.YFilter
    ciscoBfdSessDiscMapTable.EntityData.YangName = "ciscoBfdSessDiscMapTable"
    ciscoBfdSessDiscMapTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessDiscMapTable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscoBfdSessDiscMapTable.EntityData.SegmentPath = "ciscoBfdSessDiscMapTable"
    ciscoBfdSessDiscMapTable.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/" + ciscoBfdSessDiscMapTable.EntityData.SegmentPath
    ciscoBfdSessDiscMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessDiscMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessDiscMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessDiscMapTable.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessDiscMapTable.EntityData.Children.Append("ciscoBfdSessDiscMapEntry", types.YChild{"CiscoBfdSessDiscMapEntry", nil})
    for i := range ciscoBfdSessDiscMapTable.CiscoBfdSessDiscMapEntry {
        ciscoBfdSessDiscMapTable.EntityData.Children.Append(types.GetSegmentPath(ciscoBfdSessDiscMapTable.CiscoBfdSessDiscMapEntry[i]), types.YChild{"CiscoBfdSessDiscMapEntry", ciscoBfdSessDiscMapTable.CiscoBfdSessDiscMapEntry[i]})
    }
    ciscoBfdSessDiscMapTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoBfdSessDiscMapTable.EntityData.YListKeys = []string {}

    return &(ciscoBfdSessDiscMapTable.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable_CiscoBfdSessDiscMapEntry
// Each row contains a mapping between a local discriminator
// value to an entry in ciscoBfdSessTable.
type CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable_CiscoBfdSessDiscMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessDiscriminator
    CiscoBfdSessDiscriminator interface{}

    // This object indicates the CiscoBfdSessIndexTC referred to by the index of
    // this row. In essence, a mapping is  provided between this index and the
    // ciscoBfdSessTable. The type is interface{} with range: 1..4294967295.
    CiscoBfdSessDiscMapIndex interface{}
}

func (ciscoBfdSessDiscMapEntry *CISCOIETFBFDMIB_CiscoBfdSessDiscMapTable_CiscoBfdSessDiscMapEntry) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessDiscMapEntry.EntityData.YFilter = ciscoBfdSessDiscMapEntry.YFilter
    ciscoBfdSessDiscMapEntry.EntityData.YangName = "ciscoBfdSessDiscMapEntry"
    ciscoBfdSessDiscMapEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessDiscMapEntry.EntityData.ParentYangName = "ciscoBfdSessDiscMapTable"
    ciscoBfdSessDiscMapEntry.EntityData.SegmentPath = "ciscoBfdSessDiscMapEntry" + types.AddKeyToken(ciscoBfdSessDiscMapEntry.CiscoBfdSessDiscriminator, "ciscoBfdSessDiscriminator")
    ciscoBfdSessDiscMapEntry.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/ciscoBfdSessDiscMapTable/" + ciscoBfdSessDiscMapEntry.EntityData.SegmentPath
    ciscoBfdSessDiscMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessDiscMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessDiscMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessDiscMapEntry.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessDiscMapEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoBfdSessDiscMapEntry.EntityData.Leafs.Append("ciscoBfdSessDiscriminator", types.YLeaf{"CiscoBfdSessDiscriminator", ciscoBfdSessDiscMapEntry.CiscoBfdSessDiscriminator})
    ciscoBfdSessDiscMapEntry.EntityData.Leafs.Append("ciscoBfdSessDiscMapIndex", types.YLeaf{"CiscoBfdSessDiscMapIndex", ciscoBfdSessDiscMapEntry.CiscoBfdSessDiscMapIndex})

    ciscoBfdSessDiscMapEntry.EntityData.YListKeys = []string {"CiscoBfdSessDiscriminator"}

    return &(ciscoBfdSessDiscMapEntry.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdSessIpMapTable
// The BFD Session IP Mapping Table maps given
// ciscoBfdSessInterface, ciscoBfdSessAddrType, and
// ciscoBbfdSessAddr to an associated BFD sessions'
// CiscoBfdSessIndexTC used in the ciscoBfdSessTable.
// This table SHOULD contains those BFD sessions are
// of IP type: singleHop(1) and multiHop(2).
type CISCOIETFBFDMIB_CiscoBfdSessIpMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row contains a mapping between ciscoBfdSessInterface,
    // ciscoBfdSessAddrType and ciscoBfdSessAddr values to an  entry in
    // ciscoBfdSessTable. The type is slice of
    // CISCOIETFBFDMIB_CiscoBfdSessIpMapTable_CiscoBfdSessIpMapEntry.
    CiscoBfdSessIpMapEntry []*CISCOIETFBFDMIB_CiscoBfdSessIpMapTable_CiscoBfdSessIpMapEntry
}

func (ciscoBfdSessIpMapTable *CISCOIETFBFDMIB_CiscoBfdSessIpMapTable) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessIpMapTable.EntityData.YFilter = ciscoBfdSessIpMapTable.YFilter
    ciscoBfdSessIpMapTable.EntityData.YangName = "ciscoBfdSessIpMapTable"
    ciscoBfdSessIpMapTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessIpMapTable.EntityData.ParentYangName = "CISCO-IETF-BFD-MIB"
    ciscoBfdSessIpMapTable.EntityData.SegmentPath = "ciscoBfdSessIpMapTable"
    ciscoBfdSessIpMapTable.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/" + ciscoBfdSessIpMapTable.EntityData.SegmentPath
    ciscoBfdSessIpMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessIpMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessIpMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessIpMapTable.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessIpMapTable.EntityData.Children.Append("ciscoBfdSessIpMapEntry", types.YChild{"CiscoBfdSessIpMapEntry", nil})
    for i := range ciscoBfdSessIpMapTable.CiscoBfdSessIpMapEntry {
        ciscoBfdSessIpMapTable.EntityData.Children.Append(types.GetSegmentPath(ciscoBfdSessIpMapTable.CiscoBfdSessIpMapEntry[i]), types.YChild{"CiscoBfdSessIpMapEntry", ciscoBfdSessIpMapTable.CiscoBfdSessIpMapEntry[i]})
    }
    ciscoBfdSessIpMapTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoBfdSessIpMapTable.EntityData.YListKeys = []string {}

    return &(ciscoBfdSessIpMapTable.EntityData)
}

// CISCOIETFBFDMIB_CiscoBfdSessIpMapTable_CiscoBfdSessIpMapEntry
// Each row contains a mapping between ciscoBfdSessInterface,
// ciscoBfdSessAddrType and ciscoBfdSessAddr values to an 
// entry in ciscoBfdSessTable.
type CISCOIETFBFDMIB_CiscoBfdSessIpMapTable_CiscoBfdSessIpMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessInterface
    CiscoBfdSessInterface interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAddrType
    CiscoBfdSessAddrType interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_ietf_bfd_mib.CISCOIETFBFDMIB_CiscoBfdSessTable_CiscoBfdSessEntry_CiscoBfdSessAddr
    CiscoBfdSessAddr interface{}

    // This object indicates the CiscoBfdSessIndexTC referred to by the indices of
    // this row. In essence, a mapping is  provided between these indices and an
    // entry in ciscoBfdSessTable. The type is interface{} with range:
    // 1..4294967295.
    CiscoBfdSessIpMapIndex interface{}
}

func (ciscoBfdSessIpMapEntry *CISCOIETFBFDMIB_CiscoBfdSessIpMapTable_CiscoBfdSessIpMapEntry) GetEntityData() *types.CommonEntityData {
    ciscoBfdSessIpMapEntry.EntityData.YFilter = ciscoBfdSessIpMapEntry.YFilter
    ciscoBfdSessIpMapEntry.EntityData.YangName = "ciscoBfdSessIpMapEntry"
    ciscoBfdSessIpMapEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoBfdSessIpMapEntry.EntityData.ParentYangName = "ciscoBfdSessIpMapTable"
    ciscoBfdSessIpMapEntry.EntityData.SegmentPath = "ciscoBfdSessIpMapEntry" + types.AddKeyToken(ciscoBfdSessIpMapEntry.CiscoBfdSessInterface, "ciscoBfdSessInterface") + types.AddKeyToken(ciscoBfdSessIpMapEntry.CiscoBfdSessAddrType, "ciscoBfdSessAddrType") + types.AddKeyToken(ciscoBfdSessIpMapEntry.CiscoBfdSessAddr, "ciscoBfdSessAddr")
    ciscoBfdSessIpMapEntry.EntityData.AbsolutePath = "CISCO-IETF-BFD-MIB:CISCO-IETF-BFD-MIB/ciscoBfdSessIpMapTable/" + ciscoBfdSessIpMapEntry.EntityData.SegmentPath
    ciscoBfdSessIpMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoBfdSessIpMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoBfdSessIpMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoBfdSessIpMapEntry.EntityData.Children = types.NewOrderedMap()
    ciscoBfdSessIpMapEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoBfdSessIpMapEntry.EntityData.Leafs.Append("ciscoBfdSessInterface", types.YLeaf{"CiscoBfdSessInterface", ciscoBfdSessIpMapEntry.CiscoBfdSessInterface})
    ciscoBfdSessIpMapEntry.EntityData.Leafs.Append("ciscoBfdSessAddrType", types.YLeaf{"CiscoBfdSessAddrType", ciscoBfdSessIpMapEntry.CiscoBfdSessAddrType})
    ciscoBfdSessIpMapEntry.EntityData.Leafs.Append("ciscoBfdSessAddr", types.YLeaf{"CiscoBfdSessAddr", ciscoBfdSessIpMapEntry.CiscoBfdSessAddr})
    ciscoBfdSessIpMapEntry.EntityData.Leafs.Append("ciscoBfdSessIpMapIndex", types.YLeaf{"CiscoBfdSessIpMapIndex", ciscoBfdSessIpMapEntry.CiscoBfdSessIpMapIndex})

    ciscoBfdSessIpMapEntry.EntityData.YListKeys = []string {"CiscoBfdSessInterface", "CiscoBfdSessAddrType", "CiscoBfdSessAddr"}

    return &(ciscoBfdSessIpMapEntry.EntityData)
}

