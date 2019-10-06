// This MIB provides configuration control and status for the
// Redundancy Framework (RF) subsystem. RF provides a mechanism
// for logical redundancy of software functionality and is
// designed to support 1:1 redundancy on processor cards. RF is
// not intended to solve all redundancy schemes. Nor is RF
// designed to support redundant hardware, such as power
// supplies.
// 
// Redundancy is concerned with the duplication of data elements
// and software functions to provide an alternative in case of
// failure. It is a key component to meeting 99.999% availability
// requirements for Class 5 carrier solutions.
// 
// In the scope of this MIB definition, peer software elements
// are redundant and redundant software elements are peers.
package cisco_rf_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_rf_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-RF-MIB CISCO-RF-MIB}", reflect.TypeOf(CISCORFMIB{}))
    ydk.RegisterEntity("CISCO-RF-MIB:CISCO-RF-MIB", reflect.TypeOf(CISCORFMIB{}))
}

// RFState represents * Sub-state of 'standbyCold'
type RFState string

const (
    RFState_notKnown RFState = "notKnown"

    RFState_disabled RFState = "disabled"

    RFState_initialization RFState = "initialization"

    RFState_negotiation RFState = "negotiation"

    RFState_standbyCold RFState = "standbyCold"

    RFState_standbyColdConfig RFState = "standbyColdConfig"

    RFState_standbyColdFileSys RFState = "standbyColdFileSys"

    RFState_standbyColdBulk RFState = "standbyColdBulk"

    RFState_standbyHot RFState = "standbyHot"

    RFState_activeFast RFState = "activeFast"

    RFState_activeDrain RFState = "activeDrain"

    RFState_activePreconfig RFState = "activePreconfig"

    RFState_activePostconfig RFState = "activePostconfig"

    RFState_active RFState = "active"

    RFState_activeExtraload RFState = "activeExtraload"

    RFState_activeHandback RFState = "activeHandback"
)

// RFIssuState represents       commitVersion state.
type RFIssuState string

const (
    RFIssuState_unset RFIssuState = "unset"

    RFIssuState_init RFIssuState = "init"

    RFIssuState_loadVersion RFIssuState = "loadVersion"

    RFIssuState_runVersion RFIssuState = "runVersion"

    RFIssuState_commitVersion RFIssuState = "commitVersion"
)

// RFAction represents When read, the value 'noAction' is always returned.
type RFAction string

const (
    RFAction_noAction RFAction = "noAction"

    RFAction_reloadPeer RFAction = "reloadPeer"

    RFAction_reloadShelf RFAction = "reloadShelf"

    RFAction_switchActivity RFAction = "switchActivity"

    RFAction_forceSwitchActivity RFAction = "forceSwitchActivity"
)

// RFMode represents       immediately able to handle new calls.
type RFMode string

const (
    RFMode_nonRedundant RFMode = "nonRedundant"

    RFMode_staticLoadShareNonRedundant RFMode = "staticLoadShareNonRedundant"

    RFMode_dynamicLoadShareNonRedundant RFMode = "dynamicLoadShareNonRedundant"

    RFMode_staticLoadShareRedundant RFMode = "staticLoadShareRedundant"

    RFMode_dynamicLoadShareRedundant RFMode = "dynamicLoadShareRedundant"

    RFMode_coldStandbyRedundant RFMode = "coldStandbyRedundant"

    RFMode_warmStandbyRedundant RFMode = "warmStandbyRedundant"

    RFMode_hotStandbyRedundant RFMode = "hotStandbyRedundant"
)

// RFClientStatus represents       loss if there is a switchover.
type RFClientStatus string

const (
    RFClientStatus_noStatus RFClientStatus = "noStatus"

    RFClientStatus_clientNotRedundant RFClientStatus = "clientNotRedundant"

    RFClientStatus_clientRedundancyInProgress RFClientStatus = "clientRedundancyInProgress"

    RFClientStatus_clientRedundant RFClientStatus = "clientRedundant"
)

// RFSwactReasonType represents     - active unit removal caused an auto SWACT
type RFSwactReasonType string

const (
    RFSwactReasonType_unsupported RFSwactReasonType = "unsupported"

    RFSwactReasonType_none RFSwactReasonType = "none"

    RFSwactReasonType_notKnown RFSwactReasonType = "notKnown"

    RFSwactReasonType_userInitiated RFSwactReasonType = "userInitiated"

    RFSwactReasonType_userForced RFSwactReasonType = "userForced"

    RFSwactReasonType_activeUnitFailed RFSwactReasonType = "activeUnitFailed"

    RFSwactReasonType_activeUnitRemoved RFSwactReasonType = "activeUnitRemoved"
)

// RFIssuStateRev1 represents       commitVersion.
type RFIssuStateRev1 string

const (
    RFIssuStateRev1_init RFIssuStateRev1 = "init"

    RFIssuStateRev1_systemReset RFIssuStateRev1 = "systemReset"

    RFIssuStateRev1_loadVersion RFIssuStateRev1 = "loadVersion"

    RFIssuStateRev1_loadVersionSwitchover RFIssuStateRev1 = "loadVersionSwitchover"

    RFIssuStateRev1_runVersion RFIssuStateRev1 = "runVersion"

    RFIssuStateRev1_runVersionSwitchover RFIssuStateRev1 = "runVersionSwitchover"

    RFIssuStateRev1_commitVersion RFIssuStateRev1 = "commitVersion"
)

// CISCORFMIB
type CISCORFMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CRFStatus CISCORFMIB_CRFStatus

    
    CRFCfg CISCORFMIB_CRFCfg

    
    CRFHistory CISCORFMIB_CRFHistory

    // This table containing a list of redundancy modes that can be supported on
    // the device.
    CRFStatusRFModeCapsTable CISCORFMIB_CRFStatusRFModeCapsTable

    // A table that tracks the history of all switchovers that have occurred since
    // system initialization. The maximum number of entries permissible in this
    // table is defined by cRFHistoryTableMaxLength. When the number of entries in
    // the table reaches the maximum limit, the next entry would replace the
    // oldest existing entry in the table.
    CRFHistorySwitchOverTable CISCORFMIB_CRFHistorySwitchOverTable

    // This table contains a list of RF clients that are registered on the device.
    // RF clients are applications that have registered with  the Redundancy
    // Facility (RF) to receive RF events and  notifications. The purpose of RF
    // clients is to synchronize  any relevant data with the standby unit.
    CRFStatusRFClientTable CISCORFMIB_CRFStatusRFClientTable
}

func (cISCORFMIB *CISCORFMIB) GetEntityData() *types.CommonEntityData {
    cISCORFMIB.EntityData.YFilter = cISCORFMIB.YFilter
    cISCORFMIB.EntityData.YangName = "CISCO-RF-MIB"
    cISCORFMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCORFMIB.EntityData.ParentYangName = "CISCO-RF-MIB"
    cISCORFMIB.EntityData.SegmentPath = "CISCO-RF-MIB:CISCO-RF-MIB"
    cISCORFMIB.EntityData.AbsolutePath = cISCORFMIB.EntityData.SegmentPath
    cISCORFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCORFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCORFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCORFMIB.EntityData.Children = types.NewOrderedMap()
    cISCORFMIB.EntityData.Children.Append("cRFStatus", types.YChild{"CRFStatus", &cISCORFMIB.CRFStatus})
    cISCORFMIB.EntityData.Children.Append("cRFCfg", types.YChild{"CRFCfg", &cISCORFMIB.CRFCfg})
    cISCORFMIB.EntityData.Children.Append("cRFHistory", types.YChild{"CRFHistory", &cISCORFMIB.CRFHistory})
    cISCORFMIB.EntityData.Children.Append("cRFStatusRFModeCapsTable", types.YChild{"CRFStatusRFModeCapsTable", &cISCORFMIB.CRFStatusRFModeCapsTable})
    cISCORFMIB.EntityData.Children.Append("cRFHistorySwitchOverTable", types.YChild{"CRFHistorySwitchOverTable", &cISCORFMIB.CRFHistorySwitchOverTable})
    cISCORFMIB.EntityData.Children.Append("cRFStatusRFClientTable", types.YChild{"CRFStatusRFClientTable", &cISCORFMIB.CRFStatusRFClientTable})
    cISCORFMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCORFMIB.EntityData.YListKeys = []string {}

    return &(cISCORFMIB.EntityData)
}

// CISCORFMIB_CRFStatus
type CISCORFMIB_CRFStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A unique identifier for this redundant unit. This identifier is
    // implementation-specific but the method for selecting the id must remain
    // consistent throughout the redundant system.  Some example identifiers
    // include: slot id, physical or logical entity id, or a unique id assigned
    // internally by the RF subsystem. The type is interface{} with range:
    // 0..2147483647.
    CRFStatusUnitId interface{}

    // The current state of RF on this unit. The type is RFState.
    CRFStatusUnitState interface{}

    // A unique identifier for the redundant peer unit. This identifier is
    // implementation-specific but the method for selecting the id must remain
    // consistent throughout the redundant system.  Some example identifiers
    // include: slot id, physical or logical entity id, or a unique id assigned
    // internally by the RF subsystem. The type is interface{} with range:
    // 0..2147483647.
    CRFStatusPeerUnitId interface{}

    // The current state of RF on the peer unit. The type is RFState.
    CRFStatusPeerUnitState interface{}

    // Indicates whether this is the primary redundant unit or not. If this unit
    // is the primary unit, this object is true. If this unit is the secondary
    // unit, this object is false.  Note that the terms 'primary/secondary' are
    // not synonymous with the terms 'active/standby'. At any given time, the
    // primary unit may be the active unit, or the primary unit may be the standby
    // unit. Likewise,   the secondary unit, at any given time, may be the active
    // unit, or the secondary unit may be the standby unit.  The primary unit is
    // given a higher priority or precedence over the secondary unit. In a race
    // condition (usually at initialization time) or any situation where the
    // redundant units are unable to successfully negotiate activity between
    // themselves, the primary unit will always become the active unit and the
    // secondary unit will fall back to standby. Only one redundant unit can be
    // the primary unit at any given time.  The algorithm for determining the
    // primary unit is system dependent, such as 'the redundant unit with the
    // lower numeric unit id is always the primary unit.'. The type is bool.
    CRFStatusPrimaryMode interface{}

    // Indicates whether the redundant peer unit has been detected or not. If the
    // redundant peer unit is detected, this object is true. If the redundant peer
    // unit is not detected, this object is false. The type is bool.
    CRFStatusDuplexMode interface{}

    // Indicates whether a manual switch of activity is permitted. If a manual
    // switch of activity is allowed, this object is false. If a manual switch of
    // activity is not allowed, this object is true. Note that the value of this
    // object is the inverse of the status of manual SWACTs.  This object does not
    // indicate whether a switch of activity is or has occurred. This object only
    // indicates if the user-controllable capability is enabled or not.  A switch
    // of activity is the event in which the standby redundant unit becomes active
    // and the previously active unit becomes standby. The type is bool.
    CRFStatusManualSwactInhibit interface{}

    // The reason for the last switch of activity. The type is RFSwactReasonType.
    CRFStatusLastSwactReasonCode interface{}

    // The value of sysUpTime when the primary redundant unit took over as active.
    // The value of this object will be 0 till the first switchover. The type is
    // interface{} with range: 0..4294967295.
    CRFStatusFailoverTime interface{}

    // The value of sysUpTime when the peer redundant unit entered the standbyHot
    // state. The value will be 0 on system initialization. The type is
    // interface{} with range: 0..4294967295.
    CRFStatusPeerStandByEntryTime interface{}

    // The current ISSU state of the system. The type is RFIssuState.
    CRFStatusIssuState interface{}

    // The current ISSU state of the system. The type is RFIssuStateRev1.
    CRFStatusIssuStateRev1 interface{}

    // The IOS version from with the user is upgrading. The type is string.
    CRFStatusIssuFromVersion interface{}

    // The IOS version to with the user is upgrading. The type is string.
    CRFStatusIssuToVersion interface{}
}

func (cRFStatus *CISCORFMIB_CRFStatus) GetEntityData() *types.CommonEntityData {
    cRFStatus.EntityData.YFilter = cRFStatus.YFilter
    cRFStatus.EntityData.YangName = "cRFStatus"
    cRFStatus.EntityData.BundleName = "cisco_ios_xe"
    cRFStatus.EntityData.ParentYangName = "CISCO-RF-MIB"
    cRFStatus.EntityData.SegmentPath = "cRFStatus"
    cRFStatus.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/" + cRFStatus.EntityData.SegmentPath
    cRFStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFStatus.EntityData.Children = types.NewOrderedMap()
    cRFStatus.EntityData.Leafs = types.NewOrderedMap()
    cRFStatus.EntityData.Leafs.Append("cRFStatusUnitId", types.YLeaf{"CRFStatusUnitId", cRFStatus.CRFStatusUnitId})
    cRFStatus.EntityData.Leafs.Append("cRFStatusUnitState", types.YLeaf{"CRFStatusUnitState", cRFStatus.CRFStatusUnitState})
    cRFStatus.EntityData.Leafs.Append("cRFStatusPeerUnitId", types.YLeaf{"CRFStatusPeerUnitId", cRFStatus.CRFStatusPeerUnitId})
    cRFStatus.EntityData.Leafs.Append("cRFStatusPeerUnitState", types.YLeaf{"CRFStatusPeerUnitState", cRFStatus.CRFStatusPeerUnitState})
    cRFStatus.EntityData.Leafs.Append("cRFStatusPrimaryMode", types.YLeaf{"CRFStatusPrimaryMode", cRFStatus.CRFStatusPrimaryMode})
    cRFStatus.EntityData.Leafs.Append("cRFStatusDuplexMode", types.YLeaf{"CRFStatusDuplexMode", cRFStatus.CRFStatusDuplexMode})
    cRFStatus.EntityData.Leafs.Append("cRFStatusManualSwactInhibit", types.YLeaf{"CRFStatusManualSwactInhibit", cRFStatus.CRFStatusManualSwactInhibit})
    cRFStatus.EntityData.Leafs.Append("cRFStatusLastSwactReasonCode", types.YLeaf{"CRFStatusLastSwactReasonCode", cRFStatus.CRFStatusLastSwactReasonCode})
    cRFStatus.EntityData.Leafs.Append("cRFStatusFailoverTime", types.YLeaf{"CRFStatusFailoverTime", cRFStatus.CRFStatusFailoverTime})
    cRFStatus.EntityData.Leafs.Append("cRFStatusPeerStandByEntryTime", types.YLeaf{"CRFStatusPeerStandByEntryTime", cRFStatus.CRFStatusPeerStandByEntryTime})
    cRFStatus.EntityData.Leafs.Append("cRFStatusIssuState", types.YLeaf{"CRFStatusIssuState", cRFStatus.CRFStatusIssuState})
    cRFStatus.EntityData.Leafs.Append("cRFStatusIssuStateRev1", types.YLeaf{"CRFStatusIssuStateRev1", cRFStatus.CRFStatusIssuStateRev1})
    cRFStatus.EntityData.Leafs.Append("cRFStatusIssuFromVersion", types.YLeaf{"CRFStatusIssuFromVersion", cRFStatus.CRFStatusIssuFromVersion})
    cRFStatus.EntityData.Leafs.Append("cRFStatusIssuToVersion", types.YLeaf{"CRFStatusIssuToVersion", cRFStatus.CRFStatusIssuToVersion})

    cRFStatus.EntityData.YListKeys = []string {}

    return &(cRFStatus.EntityData)
}

// CISCORFMIB_CRFCfg
type CISCORFMIB_CRFCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether redundant units may communicate synchronization messages
    // with each other. If communication is not permitted, this object is set to
    // true. If communication is permitted, this object is set to false.  In split
    // mode (true), the active unit will not communicate with the standby unit.
    // The standby unit progression will not occur. When split mode is disabled
    // (false), the standby unit is reset to recover.  Split mode (true) is useful
    // for maintenance operations. The type is bool.
    CRFCfgSplitMode interface{}

    // On platforms that support keep-alives, the keep-alive threshold value
    // designates the number of lost keep-alives tolerated before a failure
    // condition is declared. If this occurs, a SWACT notification is sent.  On
    // platforms that do not support keep-alives, this object has no purpose or
    // effect. The type is interface{} with range: 0..4294967295.
    CRFCfgKeepaliveThresh interface{}

    // The minimum acceptable value for the cRFCfgKeepaliveThresh object. The type
    // is interface{} with range: 0..4294967295.
    CRFCfgKeepaliveThreshMin interface{}

    // The maximum acceptable value for the cRFCfgKeepaliveThresh object. The type
    // is interface{} with range: 0..4294967295.
    CRFCfgKeepaliveThreshMax interface{}

    // On platforms that support keep-alives, the keep-alive timer value is used
    // to guard against lost keep-alives. The RF subsystem expects to receive a
    // keep-alive within this period. If a keep-alive is not received within this
    // time period, a SWACT notification is sent.  On platforms that do not
    // support keep-alives, this object has no purpose or effect. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CRFCfgKeepaliveTimer interface{}

    // The minimum acceptable value for the cRFCfgKeepaliveTimer object. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    CRFCfgKeepaliveTimerMin interface{}

    // The maximum acceptable value for the cRFCfgKeepaliveTimer object. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    CRFCfgKeepaliveTimerMax interface{}

    // Note that the term 'notification' here refers to an RF notification and not
    // an SNMP notification.  As the standby unit progresses to the 'standbyHot'
    // state, asynchronous messages are sent from the active unit to the standby
    // unit which must then be acknowledged by the standby unit. If the active
    // unit receives the acknowledgement during the time period specified by this
    // object, progression proceeds as normal. If the timer expires and an
    // acknowledgement was not received by the active unit, a switch of activity
    // occurs. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    CRFCfgNotifTimer interface{}

    // The minimum acceptable value for the cRFCfgNotifTimer object. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CRFCfgNotifTimerMin interface{}

    // The maximum acceptable value for the cRFCfgNotifTimer object. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CRFCfgNotifTimerMax interface{}

    // This variable is set to invoke RF subsystem action commands. The commands
    // are useful for maintenance and software upgrade activities. The type is
    // RFAction.
    CRFCfgAdminAction interface{}

    // Allows enabling/disabling of RF subsystem notifications. The type is bool.
    CRFCfgNotifsEnabled interface{}

    // Indicates whether redundant units may communicate synchronization messages
    // with each other. If communication is not permitted, this object is set to
    // 'true'. If communication is permitted, this object is set to 'false'.  If
    // the value of this object is 'true', the redundant system is considered to
    // be in a maintenance mode of operation. If the value of this object is
    // 'false', the redundant system is considered to be in a normal
    // (non-maintenance) mode of operation.  In maintenance mode (true), the
    // active unit will not communicate with the standby unit. The standby unit
    // progression will not occur. When maintenance mode is disabled (false), the
    // standby unit is reset to recover.  Maintenance mode (true) is useful for
    // maintenance-type operations. The type is bool.
    CRFCfgMaintenanceMode interface{}

    // Indicates the redundancy mode configured on the device. The type is RFMode.
    CRFCfgRedundancyMode interface{}

    // Further clarifies or describes the redundancy mode indicated by
    // cRFCfgRedundancyMode. Implementation-specific terminology associated with
    // the current redundancy mode may be presented here. The type is string.
    CRFCfgRedundancyModeDescr interface{}

    // Indicate the operational redundancy mode of the device. The type is RFMode.
    CRFCfgRedundancyOperMode interface{}
}

func (cRFCfg *CISCORFMIB_CRFCfg) GetEntityData() *types.CommonEntityData {
    cRFCfg.EntityData.YFilter = cRFCfg.YFilter
    cRFCfg.EntityData.YangName = "cRFCfg"
    cRFCfg.EntityData.BundleName = "cisco_ios_xe"
    cRFCfg.EntityData.ParentYangName = "CISCO-RF-MIB"
    cRFCfg.EntityData.SegmentPath = "cRFCfg"
    cRFCfg.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/" + cRFCfg.EntityData.SegmentPath
    cRFCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFCfg.EntityData.Children = types.NewOrderedMap()
    cRFCfg.EntityData.Leafs = types.NewOrderedMap()
    cRFCfg.EntityData.Leafs.Append("cRFCfgSplitMode", types.YLeaf{"CRFCfgSplitMode", cRFCfg.CRFCfgSplitMode})
    cRFCfg.EntityData.Leafs.Append("cRFCfgKeepaliveThresh", types.YLeaf{"CRFCfgKeepaliveThresh", cRFCfg.CRFCfgKeepaliveThresh})
    cRFCfg.EntityData.Leafs.Append("cRFCfgKeepaliveThreshMin", types.YLeaf{"CRFCfgKeepaliveThreshMin", cRFCfg.CRFCfgKeepaliveThreshMin})
    cRFCfg.EntityData.Leafs.Append("cRFCfgKeepaliveThreshMax", types.YLeaf{"CRFCfgKeepaliveThreshMax", cRFCfg.CRFCfgKeepaliveThreshMax})
    cRFCfg.EntityData.Leafs.Append("cRFCfgKeepaliveTimer", types.YLeaf{"CRFCfgKeepaliveTimer", cRFCfg.CRFCfgKeepaliveTimer})
    cRFCfg.EntityData.Leafs.Append("cRFCfgKeepaliveTimerMin", types.YLeaf{"CRFCfgKeepaliveTimerMin", cRFCfg.CRFCfgKeepaliveTimerMin})
    cRFCfg.EntityData.Leafs.Append("cRFCfgKeepaliveTimerMax", types.YLeaf{"CRFCfgKeepaliveTimerMax", cRFCfg.CRFCfgKeepaliveTimerMax})
    cRFCfg.EntityData.Leafs.Append("cRFCfgNotifTimer", types.YLeaf{"CRFCfgNotifTimer", cRFCfg.CRFCfgNotifTimer})
    cRFCfg.EntityData.Leafs.Append("cRFCfgNotifTimerMin", types.YLeaf{"CRFCfgNotifTimerMin", cRFCfg.CRFCfgNotifTimerMin})
    cRFCfg.EntityData.Leafs.Append("cRFCfgNotifTimerMax", types.YLeaf{"CRFCfgNotifTimerMax", cRFCfg.CRFCfgNotifTimerMax})
    cRFCfg.EntityData.Leafs.Append("cRFCfgAdminAction", types.YLeaf{"CRFCfgAdminAction", cRFCfg.CRFCfgAdminAction})
    cRFCfg.EntityData.Leafs.Append("cRFCfgNotifsEnabled", types.YLeaf{"CRFCfgNotifsEnabled", cRFCfg.CRFCfgNotifsEnabled})
    cRFCfg.EntityData.Leafs.Append("cRFCfgMaintenanceMode", types.YLeaf{"CRFCfgMaintenanceMode", cRFCfg.CRFCfgMaintenanceMode})
    cRFCfg.EntityData.Leafs.Append("cRFCfgRedundancyMode", types.YLeaf{"CRFCfgRedundancyMode", cRFCfg.CRFCfgRedundancyMode})
    cRFCfg.EntityData.Leafs.Append("cRFCfgRedundancyModeDescr", types.YLeaf{"CRFCfgRedundancyModeDescr", cRFCfg.CRFCfgRedundancyModeDescr})
    cRFCfg.EntityData.Leafs.Append("cRFCfgRedundancyOperMode", types.YLeaf{"CRFCfgRedundancyOperMode", cRFCfg.CRFCfgRedundancyOperMode})

    cRFCfg.EntityData.YListKeys = []string {}

    return &(cRFCfg.EntityData)
}

// CISCORFMIB_CRFHistory
type CISCORFMIB_CRFHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of entries permissible in the history table. A value of 0
    // will result in no history being maintained. The type is interface{} with
    // range: 0..50.
    CRFHistoryTableMaxLength interface{}

    // Indicates the number of system cold starts. This includes the number of
    // system cold starts due to switchover failure and the number of manual
    // restarts. The type is interface{} with range: 0..4294967295.
    CRFHistoryColdStarts interface{}

    // Indicates the cumulative time that a standby redundant unit has been
    // available since last system initialization. The type is interface{} with
    // range: 0..2147483647.
    CRFHistoryStandByAvailTime interface{}
}

func (cRFHistory *CISCORFMIB_CRFHistory) GetEntityData() *types.CommonEntityData {
    cRFHistory.EntityData.YFilter = cRFHistory.YFilter
    cRFHistory.EntityData.YangName = "cRFHistory"
    cRFHistory.EntityData.BundleName = "cisco_ios_xe"
    cRFHistory.EntityData.ParentYangName = "CISCO-RF-MIB"
    cRFHistory.EntityData.SegmentPath = "cRFHistory"
    cRFHistory.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/" + cRFHistory.EntityData.SegmentPath
    cRFHistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFHistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFHistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFHistory.EntityData.Children = types.NewOrderedMap()
    cRFHistory.EntityData.Leafs = types.NewOrderedMap()
    cRFHistory.EntityData.Leafs.Append("cRFHistoryTableMaxLength", types.YLeaf{"CRFHistoryTableMaxLength", cRFHistory.CRFHistoryTableMaxLength})
    cRFHistory.EntityData.Leafs.Append("cRFHistoryColdStarts", types.YLeaf{"CRFHistoryColdStarts", cRFHistory.CRFHistoryColdStarts})
    cRFHistory.EntityData.Leafs.Append("cRFHistoryStandByAvailTime", types.YLeaf{"CRFHistoryStandByAvailTime", cRFHistory.CRFHistoryStandByAvailTime})

    cRFHistory.EntityData.YListKeys = []string {}

    return &(cRFHistory.EntityData)
}

// CISCORFMIB_CRFStatusRFModeCapsTable
// This table containing a list of redundancy modes that can be
// supported on the device.
type CISCORFMIB_CRFStatusRFModeCapsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the device implementation specific terminology
    // associated with the redundancy mode that can be supported on the device.
    // The type is slice of
    // CISCORFMIB_CRFStatusRFModeCapsTable_CRFStatusRFModeCapsEntry.
    CRFStatusRFModeCapsEntry []*CISCORFMIB_CRFStatusRFModeCapsTable_CRFStatusRFModeCapsEntry
}

func (cRFStatusRFModeCapsTable *CISCORFMIB_CRFStatusRFModeCapsTable) GetEntityData() *types.CommonEntityData {
    cRFStatusRFModeCapsTable.EntityData.YFilter = cRFStatusRFModeCapsTable.YFilter
    cRFStatusRFModeCapsTable.EntityData.YangName = "cRFStatusRFModeCapsTable"
    cRFStatusRFModeCapsTable.EntityData.BundleName = "cisco_ios_xe"
    cRFStatusRFModeCapsTable.EntityData.ParentYangName = "CISCO-RF-MIB"
    cRFStatusRFModeCapsTable.EntityData.SegmentPath = "cRFStatusRFModeCapsTable"
    cRFStatusRFModeCapsTable.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/" + cRFStatusRFModeCapsTable.EntityData.SegmentPath
    cRFStatusRFModeCapsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFStatusRFModeCapsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFStatusRFModeCapsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFStatusRFModeCapsTable.EntityData.Children = types.NewOrderedMap()
    cRFStatusRFModeCapsTable.EntityData.Children.Append("cRFStatusRFModeCapsEntry", types.YChild{"CRFStatusRFModeCapsEntry", nil})
    for i := range cRFStatusRFModeCapsTable.CRFStatusRFModeCapsEntry {
        cRFStatusRFModeCapsTable.EntityData.Children.Append(types.GetSegmentPath(cRFStatusRFModeCapsTable.CRFStatusRFModeCapsEntry[i]), types.YChild{"CRFStatusRFModeCapsEntry", cRFStatusRFModeCapsTable.CRFStatusRFModeCapsEntry[i]})
    }
    cRFStatusRFModeCapsTable.EntityData.Leafs = types.NewOrderedMap()

    cRFStatusRFModeCapsTable.EntityData.YListKeys = []string {}

    return &(cRFStatusRFModeCapsTable.EntityData)
}

// CISCORFMIB_CRFStatusRFModeCapsTable_CRFStatusRFModeCapsEntry
// An entry containing the device implementation specific
// terminology associated with the redundancy mode that can be
// supported on the device.
type CISCORFMIB_CRFStatusRFModeCapsTable_CRFStatusRFModeCapsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The redundancy mode that can be supported on the
    // device. The type is RFMode.
    CRFStatusRFModeCapsMode interface{}

    // The description of the device implementation specific terminology
    // associated with its supported redundancy mode. The type is string.
    CRFStatusRFModeCapsModeDescr interface{}
}

func (cRFStatusRFModeCapsEntry *CISCORFMIB_CRFStatusRFModeCapsTable_CRFStatusRFModeCapsEntry) GetEntityData() *types.CommonEntityData {
    cRFStatusRFModeCapsEntry.EntityData.YFilter = cRFStatusRFModeCapsEntry.YFilter
    cRFStatusRFModeCapsEntry.EntityData.YangName = "cRFStatusRFModeCapsEntry"
    cRFStatusRFModeCapsEntry.EntityData.BundleName = "cisco_ios_xe"
    cRFStatusRFModeCapsEntry.EntityData.ParentYangName = "cRFStatusRFModeCapsTable"
    cRFStatusRFModeCapsEntry.EntityData.SegmentPath = "cRFStatusRFModeCapsEntry" + types.AddKeyToken(cRFStatusRFModeCapsEntry.CRFStatusRFModeCapsMode, "cRFStatusRFModeCapsMode")
    cRFStatusRFModeCapsEntry.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/cRFStatusRFModeCapsTable/" + cRFStatusRFModeCapsEntry.EntityData.SegmentPath
    cRFStatusRFModeCapsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFStatusRFModeCapsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFStatusRFModeCapsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFStatusRFModeCapsEntry.EntityData.Children = types.NewOrderedMap()
    cRFStatusRFModeCapsEntry.EntityData.Leafs = types.NewOrderedMap()
    cRFStatusRFModeCapsEntry.EntityData.Leafs.Append("cRFStatusRFModeCapsMode", types.YLeaf{"CRFStatusRFModeCapsMode", cRFStatusRFModeCapsEntry.CRFStatusRFModeCapsMode})
    cRFStatusRFModeCapsEntry.EntityData.Leafs.Append("cRFStatusRFModeCapsModeDescr", types.YLeaf{"CRFStatusRFModeCapsModeDescr", cRFStatusRFModeCapsEntry.CRFStatusRFModeCapsModeDescr})

    cRFStatusRFModeCapsEntry.EntityData.YListKeys = []string {"CRFStatusRFModeCapsMode"}

    return &(cRFStatusRFModeCapsEntry.EntityData)
}

// CISCORFMIB_CRFHistorySwitchOverTable
// A table that tracks the history of all switchovers that
// have occurred since system initialization. The maximum
// number of entries permissible in this table is defined by
// cRFHistoryTableMaxLength. When the number of entries in
// the table reaches the maximum limit, the next entry
// would replace the oldest existing entry in the table.
type CISCORFMIB_CRFHistorySwitchOverTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The entries in this table contain the switchover information. Each entry in
    // the table is indexed by cRFHistorySwitchOverIndex. The index wraps around
    // to 1 after reaching the maximum value. The type is slice of
    // CISCORFMIB_CRFHistorySwitchOverTable_CRFHistorySwitchOverEntry.
    CRFHistorySwitchOverEntry []*CISCORFMIB_CRFHistorySwitchOverTable_CRFHistorySwitchOverEntry
}

func (cRFHistorySwitchOverTable *CISCORFMIB_CRFHistorySwitchOverTable) GetEntityData() *types.CommonEntityData {
    cRFHistorySwitchOverTable.EntityData.YFilter = cRFHistorySwitchOverTable.YFilter
    cRFHistorySwitchOverTable.EntityData.YangName = "cRFHistorySwitchOverTable"
    cRFHistorySwitchOverTable.EntityData.BundleName = "cisco_ios_xe"
    cRFHistorySwitchOverTable.EntityData.ParentYangName = "CISCO-RF-MIB"
    cRFHistorySwitchOverTable.EntityData.SegmentPath = "cRFHistorySwitchOverTable"
    cRFHistorySwitchOverTable.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/" + cRFHistorySwitchOverTable.EntityData.SegmentPath
    cRFHistorySwitchOverTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFHistorySwitchOverTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFHistorySwitchOverTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFHistorySwitchOverTable.EntityData.Children = types.NewOrderedMap()
    cRFHistorySwitchOverTable.EntityData.Children.Append("cRFHistorySwitchOverEntry", types.YChild{"CRFHistorySwitchOverEntry", nil})
    for i := range cRFHistorySwitchOverTable.CRFHistorySwitchOverEntry {
        cRFHistorySwitchOverTable.EntityData.Children.Append(types.GetSegmentPath(cRFHistorySwitchOverTable.CRFHistorySwitchOverEntry[i]), types.YChild{"CRFHistorySwitchOverEntry", cRFHistorySwitchOverTable.CRFHistorySwitchOverEntry[i]})
    }
    cRFHistorySwitchOverTable.EntityData.Leafs = types.NewOrderedMap()

    cRFHistorySwitchOverTable.EntityData.YListKeys = []string {}

    return &(cRFHistorySwitchOverTable.EntityData)
}

// CISCORFMIB_CRFHistorySwitchOverTable_CRFHistorySwitchOverEntry
// The entries in this table contain the switchover
// information. Each entry in the table is indexed by
// cRFHistorySwitchOverIndex. The index wraps around to 1
// after reaching the maximum value.
type CISCORFMIB_CRFHistorySwitchOverTable_CRFHistorySwitchOverEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A monotonically increasing integer for the purpose
    // of indexing history table. After reaching maximum value, it wraps around to
    // 1. The type is interface{} with range: 1..4294967295.
    CRFHistorySwitchOverIndex interface{}

    // Indicates the primary redundant unit that went down. The type is
    // interface{} with range: 0..2147483647.
    CRFHistoryPrevActiveUnitId interface{}

    // Indicates the secondary redundant unit that took over as active. The type
    // is interface{} with range: 0..2147483647.
    CRFHistoryCurrActiveUnitId interface{}

    // Indicates the reason for the switchover. The type is RFSwactReasonType.
    CRFHistorySwitchOverReason interface{}

    // Indicates the Date & Time when switchover occurred. The type is string.
    CRFHistorySwactTime interface{}
}

func (cRFHistorySwitchOverEntry *CISCORFMIB_CRFHistorySwitchOverTable_CRFHistorySwitchOverEntry) GetEntityData() *types.CommonEntityData {
    cRFHistorySwitchOverEntry.EntityData.YFilter = cRFHistorySwitchOverEntry.YFilter
    cRFHistorySwitchOverEntry.EntityData.YangName = "cRFHistorySwitchOverEntry"
    cRFHistorySwitchOverEntry.EntityData.BundleName = "cisco_ios_xe"
    cRFHistorySwitchOverEntry.EntityData.ParentYangName = "cRFHistorySwitchOverTable"
    cRFHistorySwitchOverEntry.EntityData.SegmentPath = "cRFHistorySwitchOverEntry" + types.AddKeyToken(cRFHistorySwitchOverEntry.CRFHistorySwitchOverIndex, "cRFHistorySwitchOverIndex")
    cRFHistorySwitchOverEntry.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/cRFHistorySwitchOverTable/" + cRFHistorySwitchOverEntry.EntityData.SegmentPath
    cRFHistorySwitchOverEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFHistorySwitchOverEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFHistorySwitchOverEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFHistorySwitchOverEntry.EntityData.Children = types.NewOrderedMap()
    cRFHistorySwitchOverEntry.EntityData.Leafs = types.NewOrderedMap()
    cRFHistorySwitchOverEntry.EntityData.Leafs.Append("cRFHistorySwitchOverIndex", types.YLeaf{"CRFHistorySwitchOverIndex", cRFHistorySwitchOverEntry.CRFHistorySwitchOverIndex})
    cRFHistorySwitchOverEntry.EntityData.Leafs.Append("cRFHistoryPrevActiveUnitId", types.YLeaf{"CRFHistoryPrevActiveUnitId", cRFHistorySwitchOverEntry.CRFHistoryPrevActiveUnitId})
    cRFHistorySwitchOverEntry.EntityData.Leafs.Append("cRFHistoryCurrActiveUnitId", types.YLeaf{"CRFHistoryCurrActiveUnitId", cRFHistorySwitchOverEntry.CRFHistoryCurrActiveUnitId})
    cRFHistorySwitchOverEntry.EntityData.Leafs.Append("cRFHistorySwitchOverReason", types.YLeaf{"CRFHistorySwitchOverReason", cRFHistorySwitchOverEntry.CRFHistorySwitchOverReason})
    cRFHistorySwitchOverEntry.EntityData.Leafs.Append("cRFHistorySwactTime", types.YLeaf{"CRFHistorySwactTime", cRFHistorySwitchOverEntry.CRFHistorySwactTime})

    cRFHistorySwitchOverEntry.EntityData.YListKeys = []string {"CRFHistorySwitchOverIndex"}

    return &(cRFHistorySwitchOverEntry.EntityData)
}

// CISCORFMIB_CRFStatusRFClientTable
// This table contains a list of RF clients that are
// registered on the device. 
// 
// RF clients are applications that have registered with 
// the Redundancy Facility (RF) to receive RF events and 
// notifications. The purpose of RF clients is to synchronize 
// any relevant data with the standby unit.
type CISCORFMIB_CRFStatusRFClientTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing information on various clients registered with the
    // Redundancy Facility (RF). Entries in this table are always created by the
    // system.  An entry is created in this table when a redundancy aware 
    // application registers with the Redundancy Facility. The entry  is destroyed
    // when that application deregisters from the  Redundancy Facility. The type
    // is slice of CISCORFMIB_CRFStatusRFClientTable_CRFStatusRFClientEntry.
    CRFStatusRFClientEntry []*CISCORFMIB_CRFStatusRFClientTable_CRFStatusRFClientEntry
}

func (cRFStatusRFClientTable *CISCORFMIB_CRFStatusRFClientTable) GetEntityData() *types.CommonEntityData {
    cRFStatusRFClientTable.EntityData.YFilter = cRFStatusRFClientTable.YFilter
    cRFStatusRFClientTable.EntityData.YangName = "cRFStatusRFClientTable"
    cRFStatusRFClientTable.EntityData.BundleName = "cisco_ios_xe"
    cRFStatusRFClientTable.EntityData.ParentYangName = "CISCO-RF-MIB"
    cRFStatusRFClientTable.EntityData.SegmentPath = "cRFStatusRFClientTable"
    cRFStatusRFClientTable.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/" + cRFStatusRFClientTable.EntityData.SegmentPath
    cRFStatusRFClientTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFStatusRFClientTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFStatusRFClientTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFStatusRFClientTable.EntityData.Children = types.NewOrderedMap()
    cRFStatusRFClientTable.EntityData.Children.Append("cRFStatusRFClientEntry", types.YChild{"CRFStatusRFClientEntry", nil})
    for i := range cRFStatusRFClientTable.CRFStatusRFClientEntry {
        cRFStatusRFClientTable.EntityData.Children.Append(types.GetSegmentPath(cRFStatusRFClientTable.CRFStatusRFClientEntry[i]), types.YChild{"CRFStatusRFClientEntry", cRFStatusRFClientTable.CRFStatusRFClientEntry[i]})
    }
    cRFStatusRFClientTable.EntityData.Leafs = types.NewOrderedMap()

    cRFStatusRFClientTable.EntityData.YListKeys = []string {}

    return &(cRFStatusRFClientTable.EntityData)
}

// CISCORFMIB_CRFStatusRFClientTable_CRFStatusRFClientEntry
// An entry containing information on various clients
// registered with the Redundancy Facility (RF). Entries in
// this table are always created by the system.
// 
// An entry is created in this table when a redundancy aware 
// application registers with the Redundancy Facility. The entry 
// is destroyed when that application deregisters from the 
// Redundancy Facility.
type CISCORFMIB_CRFStatusRFClientTable_CRFStatusRFClientEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique identifier for the client which
    // registered with the Redundancy Facility. The type is interface{} with
    // range: 1..4294967295.
    CRFStatusRFClientID interface{}

    // The description of the client which has registered with the Redundancy
    // Facility. The type is string.
    CRFStatusRFClientDescr interface{}

    // The sequence number of the client. The system assigns the sequence numbers
    // based on the order of registration of the Redundancy Facility clients. 
    // This is used for deciding order of RF events sent to clients. The type is
    // interface{} with range: 0..4294967295.
    CRFStatusRFClientSeq interface{}

    // Time taken for this client to become Redundant. This value is meaningful
    // when the value of cRFStatusRFClientStatus is not 'noStatus'. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CRFStatusRFClientRedTime interface{}

    // This object provides the status of the Redundancy Facility client. The type
    // is RFClientStatus.
    CRFStatusRFClientStatus interface{}
}

func (cRFStatusRFClientEntry *CISCORFMIB_CRFStatusRFClientTable_CRFStatusRFClientEntry) GetEntityData() *types.CommonEntityData {
    cRFStatusRFClientEntry.EntityData.YFilter = cRFStatusRFClientEntry.YFilter
    cRFStatusRFClientEntry.EntityData.YangName = "cRFStatusRFClientEntry"
    cRFStatusRFClientEntry.EntityData.BundleName = "cisco_ios_xe"
    cRFStatusRFClientEntry.EntityData.ParentYangName = "cRFStatusRFClientTable"
    cRFStatusRFClientEntry.EntityData.SegmentPath = "cRFStatusRFClientEntry" + types.AddKeyToken(cRFStatusRFClientEntry.CRFStatusRFClientID, "cRFStatusRFClientID")
    cRFStatusRFClientEntry.EntityData.AbsolutePath = "CISCO-RF-MIB:CISCO-RF-MIB/cRFStatusRFClientTable/" + cRFStatusRFClientEntry.EntityData.SegmentPath
    cRFStatusRFClientEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cRFStatusRFClientEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cRFStatusRFClientEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cRFStatusRFClientEntry.EntityData.Children = types.NewOrderedMap()
    cRFStatusRFClientEntry.EntityData.Leafs = types.NewOrderedMap()
    cRFStatusRFClientEntry.EntityData.Leafs.Append("cRFStatusRFClientID", types.YLeaf{"CRFStatusRFClientID", cRFStatusRFClientEntry.CRFStatusRFClientID})
    cRFStatusRFClientEntry.EntityData.Leafs.Append("cRFStatusRFClientDescr", types.YLeaf{"CRFStatusRFClientDescr", cRFStatusRFClientEntry.CRFStatusRFClientDescr})
    cRFStatusRFClientEntry.EntityData.Leafs.Append("cRFStatusRFClientSeq", types.YLeaf{"CRFStatusRFClientSeq", cRFStatusRFClientEntry.CRFStatusRFClientSeq})
    cRFStatusRFClientEntry.EntityData.Leafs.Append("cRFStatusRFClientRedTime", types.YLeaf{"CRFStatusRFClientRedTime", cRFStatusRFClientEntry.CRFStatusRFClientRedTime})
    cRFStatusRFClientEntry.EntityData.Leafs.Append("cRFStatusRFClientStatus", types.YLeaf{"CRFStatusRFClientStatus", cRFStatusRFClientEntry.CRFStatusRFClientStatus})

    cRFStatusRFClientEntry.EntityData.YListKeys = []string {"CRFStatusRFClientID"}

    return &(cRFStatusRFClientEntry.EntityData)
}

