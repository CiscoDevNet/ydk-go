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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Crfstatus CISCORFMIB_Crfstatus

    
    Crfcfg CISCORFMIB_Crfcfg

    
    Crfhistory CISCORFMIB_Crfhistory

    // This table containing a list of redundancy modes that can be supported on
    // the device.
    Crfstatusrfmodecapstable CISCORFMIB_Crfstatusrfmodecapstable

    // A table that tracks the history of all switchovers that have occurred since
    // system initialization. The maximum number of entries permissible in this
    // table is defined by cRFHistoryTableMaxLength. When the number of entries in
    // the table reaches the maximum limit, the next entry would replace the
    // oldest existing entry in the table.
    Crfhistoryswitchovertable CISCORFMIB_Crfhistoryswitchovertable

    // This table contains a list of RF clients that are registered on the device.
    // RF clients are applications that have registered with  the Redundancy
    // Facility (RF) to receive RF events and  notifications. The purpose of RF
    // clients is to synchronize  any relevant data with the standby unit.
    Crfstatusrfclienttable CISCORFMIB_Crfstatusrfclienttable
}

func (cISCORFMIB *CISCORFMIB) GetFilter() yfilter.YFilter { return cISCORFMIB.YFilter }

func (cISCORFMIB *CISCORFMIB) SetFilter(yf yfilter.YFilter) { cISCORFMIB.YFilter = yf }

func (cISCORFMIB *CISCORFMIB) GetGoName(yname string) string {
    if yname == "cRFStatus" { return "Crfstatus" }
    if yname == "cRFCfg" { return "Crfcfg" }
    if yname == "cRFHistory" { return "Crfhistory" }
    if yname == "cRFStatusRFModeCapsTable" { return "Crfstatusrfmodecapstable" }
    if yname == "cRFHistorySwitchOverTable" { return "Crfhistoryswitchovertable" }
    if yname == "cRFStatusRFClientTable" { return "Crfstatusrfclienttable" }
    return ""
}

func (cISCORFMIB *CISCORFMIB) GetSegmentPath() string {
    return "CISCO-RF-MIB:CISCO-RF-MIB"
}

func (cISCORFMIB *CISCORFMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cRFStatus" {
        return &cISCORFMIB.Crfstatus
    }
    if childYangName == "cRFCfg" {
        return &cISCORFMIB.Crfcfg
    }
    if childYangName == "cRFHistory" {
        return &cISCORFMIB.Crfhistory
    }
    if childYangName == "cRFStatusRFModeCapsTable" {
        return &cISCORFMIB.Crfstatusrfmodecapstable
    }
    if childYangName == "cRFHistorySwitchOverTable" {
        return &cISCORFMIB.Crfhistoryswitchovertable
    }
    if childYangName == "cRFStatusRFClientTable" {
        return &cISCORFMIB.Crfstatusrfclienttable
    }
    return nil
}

func (cISCORFMIB *CISCORFMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cRFStatus"] = &cISCORFMIB.Crfstatus
    children["cRFCfg"] = &cISCORFMIB.Crfcfg
    children["cRFHistory"] = &cISCORFMIB.Crfhistory
    children["cRFStatusRFModeCapsTable"] = &cISCORFMIB.Crfstatusrfmodecapstable
    children["cRFHistorySwitchOverTable"] = &cISCORFMIB.Crfhistoryswitchovertable
    children["cRFStatusRFClientTable"] = &cISCORFMIB.Crfstatusrfclienttable
    return children
}

func (cISCORFMIB *CISCORFMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCORFMIB *CISCORFMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCORFMIB *CISCORFMIB) GetYangName() string { return "CISCO-RF-MIB" }

func (cISCORFMIB *CISCORFMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCORFMIB *CISCORFMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCORFMIB *CISCORFMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCORFMIB *CISCORFMIB) SetParent(parent types.Entity) { cISCORFMIB.parent = parent }

func (cISCORFMIB *CISCORFMIB) GetParent() types.Entity { return cISCORFMIB.parent }

func (cISCORFMIB *CISCORFMIB) GetParentYangName() string { return "CISCO-RF-MIB" }

// CISCORFMIB_Crfstatus
type CISCORFMIB_Crfstatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A unique identifier for this redundant unit. This identifier is
    // implementation-specific but the method for selecting the id must remain
    // consistent throughout the redundant system.  Some example identifiers
    // include: slot id, physical or logical entity id, or a unique id assigned
    // internally by the RF subsystem. The type is interface{} with range:
    // 0..2147483647.
    Crfstatusunitid interface{}

    // The current state of RF on this unit. The type is RFState.
    Crfstatusunitstate interface{}

    // A unique identifier for the redundant peer unit. This identifier is
    // implementation-specific but the method for selecting the id must remain
    // consistent throughout the redundant system.  Some example identifiers
    // include: slot id, physical or logical entity id, or a unique id assigned
    // internally by the RF subsystem. The type is interface{} with range:
    // 0..2147483647.
    Crfstatuspeerunitid interface{}

    // The current state of RF on the peer unit. The type is RFState.
    Crfstatuspeerunitstate interface{}

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
    Crfstatusprimarymode interface{}

    // Indicates whether the redundant peer unit has been detected or not. If the
    // redundant peer unit is detected, this object is true. If the redundant peer
    // unit is not detected, this object is false. The type is bool.
    Crfstatusduplexmode interface{}

    // Indicates whether a manual switch of activity is permitted. If a manual
    // switch of activity is allowed, this object is false. If a manual switch of
    // activity is not allowed, this object is true. Note that the value of this
    // object is the inverse of the status of manual SWACTs.  This object does not
    // indicate whether a switch of activity is or has occurred. This object only
    // indicates if the user-controllable capability is enabled or not.  A switch
    // of activity is the event in which the standby redundant unit becomes active
    // and the previously active unit becomes standby. The type is bool.
    Crfstatusmanualswactinhibit interface{}

    // The reason for the last switch of activity. The type is RFSwactReasonType.
    Crfstatuslastswactreasoncode interface{}

    // The value of sysUpTime when the primary redundant unit took over as active.
    // The value of this object will be 0 till the first switchover. The type is
    // interface{} with range: 0..4294967295.
    Crfstatusfailovertime interface{}

    // The value of sysUpTime when the peer redundant unit entered the standbyHot
    // state. The value will be 0 on system initialization. The type is
    // interface{} with range: 0..4294967295.
    Crfstatuspeerstandbyentrytime interface{}

    // The current ISSU state of the system. The type is RFIssuState.
    Crfstatusissustate interface{}

    // The current ISSU state of the system. The type is RFIssuStateRev1.
    Crfstatusissustaterev1 interface{}

    // The IOS version from with the user is upgrading. The type is string.
    Crfstatusissufromversion interface{}

    // The IOS version to with the user is upgrading. The type is string.
    Crfstatusissutoversion interface{}
}

func (crfstatus *CISCORFMIB_Crfstatus) GetFilter() yfilter.YFilter { return crfstatus.YFilter }

func (crfstatus *CISCORFMIB_Crfstatus) SetFilter(yf yfilter.YFilter) { crfstatus.YFilter = yf }

func (crfstatus *CISCORFMIB_Crfstatus) GetGoName(yname string) string {
    if yname == "cRFStatusUnitId" { return "Crfstatusunitid" }
    if yname == "cRFStatusUnitState" { return "Crfstatusunitstate" }
    if yname == "cRFStatusPeerUnitId" { return "Crfstatuspeerunitid" }
    if yname == "cRFStatusPeerUnitState" { return "Crfstatuspeerunitstate" }
    if yname == "cRFStatusPrimaryMode" { return "Crfstatusprimarymode" }
    if yname == "cRFStatusDuplexMode" { return "Crfstatusduplexmode" }
    if yname == "cRFStatusManualSwactInhibit" { return "Crfstatusmanualswactinhibit" }
    if yname == "cRFStatusLastSwactReasonCode" { return "Crfstatuslastswactreasoncode" }
    if yname == "cRFStatusFailoverTime" { return "Crfstatusfailovertime" }
    if yname == "cRFStatusPeerStandByEntryTime" { return "Crfstatuspeerstandbyentrytime" }
    if yname == "cRFStatusIssuState" { return "Crfstatusissustate" }
    if yname == "cRFStatusIssuStateRev1" { return "Crfstatusissustaterev1" }
    if yname == "cRFStatusIssuFromVersion" { return "Crfstatusissufromversion" }
    if yname == "cRFStatusIssuToVersion" { return "Crfstatusissutoversion" }
    return ""
}

func (crfstatus *CISCORFMIB_Crfstatus) GetSegmentPath() string {
    return "cRFStatus"
}

func (crfstatus *CISCORFMIB_Crfstatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crfstatus *CISCORFMIB_Crfstatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crfstatus *CISCORFMIB_Crfstatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cRFStatusUnitId"] = crfstatus.Crfstatusunitid
    leafs["cRFStatusUnitState"] = crfstatus.Crfstatusunitstate
    leafs["cRFStatusPeerUnitId"] = crfstatus.Crfstatuspeerunitid
    leafs["cRFStatusPeerUnitState"] = crfstatus.Crfstatuspeerunitstate
    leafs["cRFStatusPrimaryMode"] = crfstatus.Crfstatusprimarymode
    leafs["cRFStatusDuplexMode"] = crfstatus.Crfstatusduplexmode
    leafs["cRFStatusManualSwactInhibit"] = crfstatus.Crfstatusmanualswactinhibit
    leafs["cRFStatusLastSwactReasonCode"] = crfstatus.Crfstatuslastswactreasoncode
    leafs["cRFStatusFailoverTime"] = crfstatus.Crfstatusfailovertime
    leafs["cRFStatusPeerStandByEntryTime"] = crfstatus.Crfstatuspeerstandbyentrytime
    leafs["cRFStatusIssuState"] = crfstatus.Crfstatusissustate
    leafs["cRFStatusIssuStateRev1"] = crfstatus.Crfstatusissustaterev1
    leafs["cRFStatusIssuFromVersion"] = crfstatus.Crfstatusissufromversion
    leafs["cRFStatusIssuToVersion"] = crfstatus.Crfstatusissutoversion
    return leafs
}

func (crfstatus *CISCORFMIB_Crfstatus) GetBundleName() string { return "cisco_ios_xe" }

func (crfstatus *CISCORFMIB_Crfstatus) GetYangName() string { return "cRFStatus" }

func (crfstatus *CISCORFMIB_Crfstatus) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfstatus *CISCORFMIB_Crfstatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfstatus *CISCORFMIB_Crfstatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfstatus *CISCORFMIB_Crfstatus) SetParent(parent types.Entity) { crfstatus.parent = parent }

func (crfstatus *CISCORFMIB_Crfstatus) GetParent() types.Entity { return crfstatus.parent }

func (crfstatus *CISCORFMIB_Crfstatus) GetParentYangName() string { return "CISCO-RF-MIB" }

// CISCORFMIB_Crfcfg
type CISCORFMIB_Crfcfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates whether redundant units may communicate synchronization messages
    // with each other. If communication is not permitted, this object is set to
    // true. If communication is permitted, this object is set to false.  In split
    // mode (true), the active unit will not communicate with the standby unit.
    // The standby unit progression will not occur. When split mode is disabled
    // (false), the standby unit is reset to recover.  Split mode (true) is useful
    // for maintenance operations. The type is bool.
    Crfcfgsplitmode interface{}

    // On platforms that support keep-alives, the keep-alive threshold value
    // designates the number of lost keep-alives tolerated before a failure
    // condition is declared. If this occurs, a SWACT notification is sent.  On
    // platforms that do not support keep-alives, this object has no purpose or
    // effect. The type is interface{} with range: 0..4294967295.
    Crfcfgkeepalivethresh interface{}

    // The minimum acceptable value for the cRFCfgKeepaliveThresh object. The type
    // is interface{} with range: 0..4294967295.
    Crfcfgkeepalivethreshmin interface{}

    // The maximum acceptable value for the cRFCfgKeepaliveThresh object. The type
    // is interface{} with range: 0..4294967295.
    Crfcfgkeepalivethreshmax interface{}

    // On platforms that support keep-alives, the keep-alive timer value is used
    // to guard against lost keep-alives. The RF subsystem expects to receive a
    // keep-alive within this period. If a keep-alive is not received within this
    // time period, a SWACT notification is sent.  On platforms that do not
    // support keep-alives, this object has no purpose or effect. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Crfcfgkeepalivetimer interface{}

    // The minimum acceptable value for the cRFCfgKeepaliveTimer object. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    Crfcfgkeepalivetimermin interface{}

    // The maximum acceptable value for the cRFCfgKeepaliveTimer object. The type
    // is interface{} with range: 0..4294967295. Units are milliseconds.
    Crfcfgkeepalivetimermax interface{}

    // Note that the term 'notification' here refers to an RF notification and not
    // an SNMP notification.  As the standby unit progresses to the 'standbyHot'
    // state, asynchronous messages are sent from the active unit to the standby
    // unit which must then be acknowledged by the standby unit. If the active
    // unit receives the acknowledgement during the time period specified by this
    // object, progression proceeds as normal. If the timer expires and an
    // acknowledgement was not received by the active unit, a switch of activity
    // occurs. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Crfcfgnotiftimer interface{}

    // The minimum acceptable value for the cRFCfgNotifTimer object. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Crfcfgnotiftimermin interface{}

    // The maximum acceptable value for the cRFCfgNotifTimer object. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Crfcfgnotiftimermax interface{}

    // This variable is set to invoke RF subsystem action commands. The commands
    // are useful for maintenance and software upgrade activities. The type is
    // RFAction.
    Crfcfgadminaction interface{}

    // Allows enabling/disabling of RF subsystem notifications. The type is bool.
    Crfcfgnotifsenabled interface{}

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
    Crfcfgmaintenancemode interface{}

    // Indicates the redundancy mode configured on the device. The type is RFMode.
    Crfcfgredundancymode interface{}

    // Further clarifies or describes the redundancy mode indicated by
    // cRFCfgRedundancyMode. Implementation-specific terminology associated with
    // the current redundancy mode may be presented here. The type is string.
    Crfcfgredundancymodedescr interface{}

    // Indicate the operational redundancy mode of the device. The type is RFMode.
    Crfcfgredundancyopermode interface{}
}

func (crfcfg *CISCORFMIB_Crfcfg) GetFilter() yfilter.YFilter { return crfcfg.YFilter }

func (crfcfg *CISCORFMIB_Crfcfg) SetFilter(yf yfilter.YFilter) { crfcfg.YFilter = yf }

func (crfcfg *CISCORFMIB_Crfcfg) GetGoName(yname string) string {
    if yname == "cRFCfgSplitMode" { return "Crfcfgsplitmode" }
    if yname == "cRFCfgKeepaliveThresh" { return "Crfcfgkeepalivethresh" }
    if yname == "cRFCfgKeepaliveThreshMin" { return "Crfcfgkeepalivethreshmin" }
    if yname == "cRFCfgKeepaliveThreshMax" { return "Crfcfgkeepalivethreshmax" }
    if yname == "cRFCfgKeepaliveTimer" { return "Crfcfgkeepalivetimer" }
    if yname == "cRFCfgKeepaliveTimerMin" { return "Crfcfgkeepalivetimermin" }
    if yname == "cRFCfgKeepaliveTimerMax" { return "Crfcfgkeepalivetimermax" }
    if yname == "cRFCfgNotifTimer" { return "Crfcfgnotiftimer" }
    if yname == "cRFCfgNotifTimerMin" { return "Crfcfgnotiftimermin" }
    if yname == "cRFCfgNotifTimerMax" { return "Crfcfgnotiftimermax" }
    if yname == "cRFCfgAdminAction" { return "Crfcfgadminaction" }
    if yname == "cRFCfgNotifsEnabled" { return "Crfcfgnotifsenabled" }
    if yname == "cRFCfgMaintenanceMode" { return "Crfcfgmaintenancemode" }
    if yname == "cRFCfgRedundancyMode" { return "Crfcfgredundancymode" }
    if yname == "cRFCfgRedundancyModeDescr" { return "Crfcfgredundancymodedescr" }
    if yname == "cRFCfgRedundancyOperMode" { return "Crfcfgredundancyopermode" }
    return ""
}

func (crfcfg *CISCORFMIB_Crfcfg) GetSegmentPath() string {
    return "cRFCfg"
}

func (crfcfg *CISCORFMIB_Crfcfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crfcfg *CISCORFMIB_Crfcfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crfcfg *CISCORFMIB_Crfcfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cRFCfgSplitMode"] = crfcfg.Crfcfgsplitmode
    leafs["cRFCfgKeepaliveThresh"] = crfcfg.Crfcfgkeepalivethresh
    leafs["cRFCfgKeepaliveThreshMin"] = crfcfg.Crfcfgkeepalivethreshmin
    leafs["cRFCfgKeepaliveThreshMax"] = crfcfg.Crfcfgkeepalivethreshmax
    leafs["cRFCfgKeepaliveTimer"] = crfcfg.Crfcfgkeepalivetimer
    leafs["cRFCfgKeepaliveTimerMin"] = crfcfg.Crfcfgkeepalivetimermin
    leafs["cRFCfgKeepaliveTimerMax"] = crfcfg.Crfcfgkeepalivetimermax
    leafs["cRFCfgNotifTimer"] = crfcfg.Crfcfgnotiftimer
    leafs["cRFCfgNotifTimerMin"] = crfcfg.Crfcfgnotiftimermin
    leafs["cRFCfgNotifTimerMax"] = crfcfg.Crfcfgnotiftimermax
    leafs["cRFCfgAdminAction"] = crfcfg.Crfcfgadminaction
    leafs["cRFCfgNotifsEnabled"] = crfcfg.Crfcfgnotifsenabled
    leafs["cRFCfgMaintenanceMode"] = crfcfg.Crfcfgmaintenancemode
    leafs["cRFCfgRedundancyMode"] = crfcfg.Crfcfgredundancymode
    leafs["cRFCfgRedundancyModeDescr"] = crfcfg.Crfcfgredundancymodedescr
    leafs["cRFCfgRedundancyOperMode"] = crfcfg.Crfcfgredundancyopermode
    return leafs
}

func (crfcfg *CISCORFMIB_Crfcfg) GetBundleName() string { return "cisco_ios_xe" }

func (crfcfg *CISCORFMIB_Crfcfg) GetYangName() string { return "cRFCfg" }

func (crfcfg *CISCORFMIB_Crfcfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfcfg *CISCORFMIB_Crfcfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfcfg *CISCORFMIB_Crfcfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfcfg *CISCORFMIB_Crfcfg) SetParent(parent types.Entity) { crfcfg.parent = parent }

func (crfcfg *CISCORFMIB_Crfcfg) GetParent() types.Entity { return crfcfg.parent }

func (crfcfg *CISCORFMIB_Crfcfg) GetParentYangName() string { return "CISCO-RF-MIB" }

// CISCORFMIB_Crfhistory
type CISCORFMIB_Crfhistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of entries permissible in the history table. A value of 0
    // will result in no history being maintained. The type is interface{} with
    // range: 0..50.
    Crfhistorytablemaxlength interface{}

    // Indicates the number of system cold starts. This includes the number of
    // system cold starts due to switchover failure and the number of manual
    // restarts. The type is interface{} with range: 0..4294967295.
    Crfhistorycoldstarts interface{}

    // Indicates the cumulative time that a standby redundant unit has been
    // available since last system initialization. The type is interface{} with
    // range: 0..2147483647.
    Crfhistorystandbyavailtime interface{}
}

func (crfhistory *CISCORFMIB_Crfhistory) GetFilter() yfilter.YFilter { return crfhistory.YFilter }

func (crfhistory *CISCORFMIB_Crfhistory) SetFilter(yf yfilter.YFilter) { crfhistory.YFilter = yf }

func (crfhistory *CISCORFMIB_Crfhistory) GetGoName(yname string) string {
    if yname == "cRFHistoryTableMaxLength" { return "Crfhistorytablemaxlength" }
    if yname == "cRFHistoryColdStarts" { return "Crfhistorycoldstarts" }
    if yname == "cRFHistoryStandByAvailTime" { return "Crfhistorystandbyavailtime" }
    return ""
}

func (crfhistory *CISCORFMIB_Crfhistory) GetSegmentPath() string {
    return "cRFHistory"
}

func (crfhistory *CISCORFMIB_Crfhistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crfhistory *CISCORFMIB_Crfhistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crfhistory *CISCORFMIB_Crfhistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cRFHistoryTableMaxLength"] = crfhistory.Crfhistorytablemaxlength
    leafs["cRFHistoryColdStarts"] = crfhistory.Crfhistorycoldstarts
    leafs["cRFHistoryStandByAvailTime"] = crfhistory.Crfhistorystandbyavailtime
    return leafs
}

func (crfhistory *CISCORFMIB_Crfhistory) GetBundleName() string { return "cisco_ios_xe" }

func (crfhistory *CISCORFMIB_Crfhistory) GetYangName() string { return "cRFHistory" }

func (crfhistory *CISCORFMIB_Crfhistory) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfhistory *CISCORFMIB_Crfhistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfhistory *CISCORFMIB_Crfhistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfhistory *CISCORFMIB_Crfhistory) SetParent(parent types.Entity) { crfhistory.parent = parent }

func (crfhistory *CISCORFMIB_Crfhistory) GetParent() types.Entity { return crfhistory.parent }

func (crfhistory *CISCORFMIB_Crfhistory) GetParentYangName() string { return "CISCO-RF-MIB" }

// CISCORFMIB_Crfstatusrfmodecapstable
// This table containing a list of redundancy modes that can be
// supported on the device.
type CISCORFMIB_Crfstatusrfmodecapstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the device implementation specific terminology
    // associated with the redundancy mode that can be supported on the device.
    // The type is slice of
    // CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry.
    Crfstatusrfmodecapsentry []CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry
}

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetFilter() yfilter.YFilter { return crfstatusrfmodecapstable.YFilter }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) SetFilter(yf yfilter.YFilter) { crfstatusrfmodecapstable.YFilter = yf }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetGoName(yname string) string {
    if yname == "cRFStatusRFModeCapsEntry" { return "Crfstatusrfmodecapsentry" }
    return ""
}

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetSegmentPath() string {
    return "cRFStatusRFModeCapsTable"
}

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cRFStatusRFModeCapsEntry" {
        for _, c := range crfstatusrfmodecapstable.Crfstatusrfmodecapsentry {
            if crfstatusrfmodecapstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry{}
        crfstatusrfmodecapstable.Crfstatusrfmodecapsentry = append(crfstatusrfmodecapstable.Crfstatusrfmodecapsentry, child)
        return &crfstatusrfmodecapstable.Crfstatusrfmodecapsentry[len(crfstatusrfmodecapstable.Crfstatusrfmodecapsentry)-1]
    }
    return nil
}

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crfstatusrfmodecapstable.Crfstatusrfmodecapsentry {
        children[crfstatusrfmodecapstable.Crfstatusrfmodecapsentry[i].GetSegmentPath()] = &crfstatusrfmodecapstable.Crfstatusrfmodecapsentry[i]
    }
    return children
}

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetBundleName() string { return "cisco_ios_xe" }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetYangName() string { return "cRFStatusRFModeCapsTable" }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) SetParent(parent types.Entity) { crfstatusrfmodecapstable.parent = parent }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetParent() types.Entity { return crfstatusrfmodecapstable.parent }

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetParentYangName() string { return "CISCO-RF-MIB" }

// CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry
// An entry containing the device implementation specific
// terminology associated with the redundancy mode that can be
// supported on the device.
type CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The redundancy mode that can be supported on the
    // device. The type is RFMode.
    Crfstatusrfmodecapsmode interface{}

    // The description of the device implementation specific terminology
    // associated with its supported redundancy mode. The type is string.
    Crfstatusrfmodecapsmodedescr interface{}
}

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetFilter() yfilter.YFilter { return crfstatusrfmodecapsentry.YFilter }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) SetFilter(yf yfilter.YFilter) { crfstatusrfmodecapsentry.YFilter = yf }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetGoName(yname string) string {
    if yname == "cRFStatusRFModeCapsMode" { return "Crfstatusrfmodecapsmode" }
    if yname == "cRFStatusRFModeCapsModeDescr" { return "Crfstatusrfmodecapsmodedescr" }
    return ""
}

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetSegmentPath() string {
    return "cRFStatusRFModeCapsEntry" + "[cRFStatusRFModeCapsMode='" + fmt.Sprintf("%v", crfstatusrfmodecapsentry.Crfstatusrfmodecapsmode) + "']"
}

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cRFStatusRFModeCapsMode"] = crfstatusrfmodecapsentry.Crfstatusrfmodecapsmode
    leafs["cRFStatusRFModeCapsModeDescr"] = crfstatusrfmodecapsentry.Crfstatusrfmodecapsmodedescr
    return leafs
}

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetBundleName() string { return "cisco_ios_xe" }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetYangName() string { return "cRFStatusRFModeCapsEntry" }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) SetParent(parent types.Entity) { crfstatusrfmodecapsentry.parent = parent }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetParent() types.Entity { return crfstatusrfmodecapsentry.parent }

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetParentYangName() string { return "cRFStatusRFModeCapsTable" }

// CISCORFMIB_Crfhistoryswitchovertable
// A table that tracks the history of all switchovers that
// have occurred since system initialization. The maximum
// number of entries permissible in this table is defined by
// cRFHistoryTableMaxLength. When the number of entries in
// the table reaches the maximum limit, the next entry
// would replace the oldest existing entry in the table.
type CISCORFMIB_Crfhistoryswitchovertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The entries in this table contain the switchover information. Each entry in
    // the table is indexed by cRFHistorySwitchOverIndex. The index wraps around
    // to 1 after reaching the maximum value. The type is slice of
    // CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry.
    Crfhistoryswitchoverentry []CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry
}

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetFilter() yfilter.YFilter { return crfhistoryswitchovertable.YFilter }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) SetFilter(yf yfilter.YFilter) { crfhistoryswitchovertable.YFilter = yf }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetGoName(yname string) string {
    if yname == "cRFHistorySwitchOverEntry" { return "Crfhistoryswitchoverentry" }
    return ""
}

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetSegmentPath() string {
    return "cRFHistorySwitchOverTable"
}

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cRFHistorySwitchOverEntry" {
        for _, c := range crfhistoryswitchovertable.Crfhistoryswitchoverentry {
            if crfhistoryswitchovertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry{}
        crfhistoryswitchovertable.Crfhistoryswitchoverentry = append(crfhistoryswitchovertable.Crfhistoryswitchoverentry, child)
        return &crfhistoryswitchovertable.Crfhistoryswitchoverentry[len(crfhistoryswitchovertable.Crfhistoryswitchoverentry)-1]
    }
    return nil
}

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crfhistoryswitchovertable.Crfhistoryswitchoverentry {
        children[crfhistoryswitchovertable.Crfhistoryswitchoverentry[i].GetSegmentPath()] = &crfhistoryswitchovertable.Crfhistoryswitchoverentry[i]
    }
    return children
}

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetBundleName() string { return "cisco_ios_xe" }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetYangName() string { return "cRFHistorySwitchOverTable" }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) SetParent(parent types.Entity) { crfhistoryswitchovertable.parent = parent }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetParent() types.Entity { return crfhistoryswitchovertable.parent }

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetParentYangName() string { return "CISCO-RF-MIB" }

// CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry
// The entries in this table contain the switchover
// information. Each entry in the table is indexed by
// cRFHistorySwitchOverIndex. The index wraps around to 1
// after reaching the maximum value.
type CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A monotonically increasing integer for the purpose
    // of indexing history table. After reaching maximum value, it wraps around to
    // 1. The type is interface{} with range: 1..4294967295.
    Crfhistoryswitchoverindex interface{}

    // Indicates the primary redundant unit that went down. The type is
    // interface{} with range: 0..2147483647.
    Crfhistoryprevactiveunitid interface{}

    // Indicates the secondary redundant unit that took over as active. The type
    // is interface{} with range: 0..2147483647.
    Crfhistorycurractiveunitid interface{}

    // Indicates the reason for the switchover. The type is RFSwactReasonType.
    Crfhistoryswitchoverreason interface{}

    // Indicates the Date & Time when switchover occurred. The type is string.
    Crfhistoryswacttime interface{}
}

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetFilter() yfilter.YFilter { return crfhistoryswitchoverentry.YFilter }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) SetFilter(yf yfilter.YFilter) { crfhistoryswitchoverentry.YFilter = yf }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetGoName(yname string) string {
    if yname == "cRFHistorySwitchOverIndex" { return "Crfhistoryswitchoverindex" }
    if yname == "cRFHistoryPrevActiveUnitId" { return "Crfhistoryprevactiveunitid" }
    if yname == "cRFHistoryCurrActiveUnitId" { return "Crfhistorycurractiveunitid" }
    if yname == "cRFHistorySwitchOverReason" { return "Crfhistoryswitchoverreason" }
    if yname == "cRFHistorySwactTime" { return "Crfhistoryswacttime" }
    return ""
}

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetSegmentPath() string {
    return "cRFHistorySwitchOverEntry" + "[cRFHistorySwitchOverIndex='" + fmt.Sprintf("%v", crfhistoryswitchoverentry.Crfhistoryswitchoverindex) + "']"
}

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cRFHistorySwitchOverIndex"] = crfhistoryswitchoverentry.Crfhistoryswitchoverindex
    leafs["cRFHistoryPrevActiveUnitId"] = crfhistoryswitchoverentry.Crfhistoryprevactiveunitid
    leafs["cRFHistoryCurrActiveUnitId"] = crfhistoryswitchoverentry.Crfhistorycurractiveunitid
    leafs["cRFHistorySwitchOverReason"] = crfhistoryswitchoverentry.Crfhistoryswitchoverreason
    leafs["cRFHistorySwactTime"] = crfhistoryswitchoverentry.Crfhistoryswacttime
    return leafs
}

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetBundleName() string { return "cisco_ios_xe" }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetYangName() string { return "cRFHistorySwitchOverEntry" }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) SetParent(parent types.Entity) { crfhistoryswitchoverentry.parent = parent }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetParent() types.Entity { return crfhistoryswitchoverentry.parent }

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetParentYangName() string { return "cRFHistorySwitchOverTable" }

// CISCORFMIB_Crfstatusrfclienttable
// This table contains a list of RF clients that are
// registered on the device. 
// 
// RF clients are applications that have registered with 
// the Redundancy Facility (RF) to receive RF events and 
// notifications. The purpose of RF clients is to synchronize 
// any relevant data with the standby unit.
type CISCORFMIB_Crfstatusrfclienttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing information on various clients registered with the
    // Redundancy Facility (RF). Entries in this table are always created by the
    // system.  An entry is created in this table when a redundancy aware 
    // application registers with the Redundancy Facility. The entry  is destroyed
    // when that application deregisters from the  Redundancy Facility. The type
    // is slice of CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry.
    Crfstatusrfcliententry []CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry
}

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetFilter() yfilter.YFilter { return crfstatusrfclienttable.YFilter }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) SetFilter(yf yfilter.YFilter) { crfstatusrfclienttable.YFilter = yf }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetGoName(yname string) string {
    if yname == "cRFStatusRFClientEntry" { return "Crfstatusrfcliententry" }
    return ""
}

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetSegmentPath() string {
    return "cRFStatusRFClientTable"
}

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cRFStatusRFClientEntry" {
        for _, c := range crfstatusrfclienttable.Crfstatusrfcliententry {
            if crfstatusrfclienttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry{}
        crfstatusrfclienttable.Crfstatusrfcliententry = append(crfstatusrfclienttable.Crfstatusrfcliententry, child)
        return &crfstatusrfclienttable.Crfstatusrfcliententry[len(crfstatusrfclienttable.Crfstatusrfcliententry)-1]
    }
    return nil
}

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crfstatusrfclienttable.Crfstatusrfcliententry {
        children[crfstatusrfclienttable.Crfstatusrfcliententry[i].GetSegmentPath()] = &crfstatusrfclienttable.Crfstatusrfcliententry[i]
    }
    return children
}

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetBundleName() string { return "cisco_ios_xe" }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetYangName() string { return "cRFStatusRFClientTable" }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) SetParent(parent types.Entity) { crfstatusrfclienttable.parent = parent }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetParent() types.Entity { return crfstatusrfclienttable.parent }

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetParentYangName() string { return "CISCO-RF-MIB" }

// CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry
// An entry containing information on various clients
// registered with the Redundancy Facility (RF). Entries in
// this table are always created by the system.
// 
// An entry is created in this table when a redundancy aware 
// application registers with the Redundancy Facility. The entry 
// is destroyed when that application deregisters from the 
// Redundancy Facility.
type CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A unique identifier for the client which
    // registered with the Redundancy Facility. The type is interface{} with
    // range: 1..4294967295.
    Crfstatusrfclientid interface{}

    // The description of the client which has registered with the Redundancy
    // Facility. The type is string.
    Crfstatusrfclientdescr interface{}

    // The sequence number of the client. The system assigns the sequence numbers
    // based on the order of registration of the Redundancy Facility clients. 
    // This is used for deciding order of RF events sent to clients. The type is
    // interface{} with range: 0..4294967295.
    Crfstatusrfclientseq interface{}

    // Time taken for this client to become Redundant. This value is meaningful
    // when the value of cRFStatusRFClientStatus is not 'noStatus'. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Crfstatusrfclientredtime interface{}

    // This object provides the status of the Redundancy Facility client. The type
    // is RFClientStatus.
    Crfstatusrfclientstatus interface{}
}

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetFilter() yfilter.YFilter { return crfstatusrfcliententry.YFilter }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) SetFilter(yf yfilter.YFilter) { crfstatusrfcliententry.YFilter = yf }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetGoName(yname string) string {
    if yname == "cRFStatusRFClientID" { return "Crfstatusrfclientid" }
    if yname == "cRFStatusRFClientDescr" { return "Crfstatusrfclientdescr" }
    if yname == "cRFStatusRFClientSeq" { return "Crfstatusrfclientseq" }
    if yname == "cRFStatusRFClientRedTime" { return "Crfstatusrfclientredtime" }
    if yname == "cRFStatusRFClientStatus" { return "Crfstatusrfclientstatus" }
    return ""
}

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetSegmentPath() string {
    return "cRFStatusRFClientEntry" + "[cRFStatusRFClientID='" + fmt.Sprintf("%v", crfstatusrfcliententry.Crfstatusrfclientid) + "']"
}

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cRFStatusRFClientID"] = crfstatusrfcliententry.Crfstatusrfclientid
    leafs["cRFStatusRFClientDescr"] = crfstatusrfcliententry.Crfstatusrfclientdescr
    leafs["cRFStatusRFClientSeq"] = crfstatusrfcliententry.Crfstatusrfclientseq
    leafs["cRFStatusRFClientRedTime"] = crfstatusrfcliententry.Crfstatusrfclientredtime
    leafs["cRFStatusRFClientStatus"] = crfstatusrfcliententry.Crfstatusrfclientstatus
    return leafs
}

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetBundleName() string { return "cisco_ios_xe" }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetYangName() string { return "cRFStatusRFClientEntry" }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) SetParent(parent types.Entity) { crfstatusrfcliententry.parent = parent }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetParent() types.Entity { return crfstatusrfcliententry.parent }

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetParentYangName() string { return "cRFStatusRFClientTable" }

