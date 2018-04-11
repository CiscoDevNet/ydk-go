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

// RFAction represents When read, the value 'noAction' is always returned.
type RFAction string

const (
    RFAction_noAction RFAction = "noAction"

    RFAction_reloadPeer RFAction = "reloadPeer"

    RFAction_reloadShelf RFAction = "reloadShelf"

    RFAction_switchActivity RFAction = "switchActivity"

    RFAction_forceSwitchActivity RFAction = "forceSwitchActivity"
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

// RFIssuState represents       commitVersion state.
type RFIssuState string

const (
    RFIssuState_unset RFIssuState = "unset"

    RFIssuState_init RFIssuState = "init"

    RFIssuState_loadVersion RFIssuState = "loadVersion"

    RFIssuState_runVersion RFIssuState = "runVersion"

    RFIssuState_commitVersion RFIssuState = "commitVersion"
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

// RFClientStatus represents       loss if there is a switchover.
type RFClientStatus string

const (
    RFClientStatus_noStatus RFClientStatus = "noStatus"

    RFClientStatus_clientNotRedundant RFClientStatus = "clientNotRedundant"

    RFClientStatus_clientRedundancyInProgress RFClientStatus = "clientRedundancyInProgress"

    RFClientStatus_clientRedundant RFClientStatus = "clientRedundant"
)

// CISCORFMIB
type CISCORFMIB struct {
    EntityData types.CommonEntityData
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

func (cISCORFMIB *CISCORFMIB) GetEntityData() *types.CommonEntityData {
    cISCORFMIB.EntityData.YFilter = cISCORFMIB.YFilter
    cISCORFMIB.EntityData.YangName = "CISCO-RF-MIB"
    cISCORFMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCORFMIB.EntityData.ParentYangName = "CISCO-RF-MIB"
    cISCORFMIB.EntityData.SegmentPath = "CISCO-RF-MIB:CISCO-RF-MIB"
    cISCORFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCORFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCORFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCORFMIB.EntityData.Children = make(map[string]types.YChild)
    cISCORFMIB.EntityData.Children["cRFStatus"] = types.YChild{"Crfstatus", &cISCORFMIB.Crfstatus}
    cISCORFMIB.EntityData.Children["cRFCfg"] = types.YChild{"Crfcfg", &cISCORFMIB.Crfcfg}
    cISCORFMIB.EntityData.Children["cRFHistory"] = types.YChild{"Crfhistory", &cISCORFMIB.Crfhistory}
    cISCORFMIB.EntityData.Children["cRFStatusRFModeCapsTable"] = types.YChild{"Crfstatusrfmodecapstable", &cISCORFMIB.Crfstatusrfmodecapstable}
    cISCORFMIB.EntityData.Children["cRFHistorySwitchOverTable"] = types.YChild{"Crfhistoryswitchovertable", &cISCORFMIB.Crfhistoryswitchovertable}
    cISCORFMIB.EntityData.Children["cRFStatusRFClientTable"] = types.YChild{"Crfstatusrfclienttable", &cISCORFMIB.Crfstatusrfclienttable}
    cISCORFMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCORFMIB.EntityData)
}

// CISCORFMIB_Crfstatus
type CISCORFMIB_Crfstatus struct {
    EntityData types.CommonEntityData
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

func (crfstatus *CISCORFMIB_Crfstatus) GetEntityData() *types.CommonEntityData {
    crfstatus.EntityData.YFilter = crfstatus.YFilter
    crfstatus.EntityData.YangName = "cRFStatus"
    crfstatus.EntityData.BundleName = "cisco_ios_xe"
    crfstatus.EntityData.ParentYangName = "CISCO-RF-MIB"
    crfstatus.EntityData.SegmentPath = "cRFStatus"
    crfstatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfstatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfstatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfstatus.EntityData.Children = make(map[string]types.YChild)
    crfstatus.EntityData.Leafs = make(map[string]types.YLeaf)
    crfstatus.EntityData.Leafs["cRFStatusUnitId"] = types.YLeaf{"Crfstatusunitid", crfstatus.Crfstatusunitid}
    crfstatus.EntityData.Leafs["cRFStatusUnitState"] = types.YLeaf{"Crfstatusunitstate", crfstatus.Crfstatusunitstate}
    crfstatus.EntityData.Leafs["cRFStatusPeerUnitId"] = types.YLeaf{"Crfstatuspeerunitid", crfstatus.Crfstatuspeerunitid}
    crfstatus.EntityData.Leafs["cRFStatusPeerUnitState"] = types.YLeaf{"Crfstatuspeerunitstate", crfstatus.Crfstatuspeerunitstate}
    crfstatus.EntityData.Leafs["cRFStatusPrimaryMode"] = types.YLeaf{"Crfstatusprimarymode", crfstatus.Crfstatusprimarymode}
    crfstatus.EntityData.Leafs["cRFStatusDuplexMode"] = types.YLeaf{"Crfstatusduplexmode", crfstatus.Crfstatusduplexmode}
    crfstatus.EntityData.Leafs["cRFStatusManualSwactInhibit"] = types.YLeaf{"Crfstatusmanualswactinhibit", crfstatus.Crfstatusmanualswactinhibit}
    crfstatus.EntityData.Leafs["cRFStatusLastSwactReasonCode"] = types.YLeaf{"Crfstatuslastswactreasoncode", crfstatus.Crfstatuslastswactreasoncode}
    crfstatus.EntityData.Leafs["cRFStatusFailoverTime"] = types.YLeaf{"Crfstatusfailovertime", crfstatus.Crfstatusfailovertime}
    crfstatus.EntityData.Leafs["cRFStatusPeerStandByEntryTime"] = types.YLeaf{"Crfstatuspeerstandbyentrytime", crfstatus.Crfstatuspeerstandbyentrytime}
    crfstatus.EntityData.Leafs["cRFStatusIssuState"] = types.YLeaf{"Crfstatusissustate", crfstatus.Crfstatusissustate}
    crfstatus.EntityData.Leafs["cRFStatusIssuStateRev1"] = types.YLeaf{"Crfstatusissustaterev1", crfstatus.Crfstatusissustaterev1}
    crfstatus.EntityData.Leafs["cRFStatusIssuFromVersion"] = types.YLeaf{"Crfstatusissufromversion", crfstatus.Crfstatusissufromversion}
    crfstatus.EntityData.Leafs["cRFStatusIssuToVersion"] = types.YLeaf{"Crfstatusissutoversion", crfstatus.Crfstatusissutoversion}
    return &(crfstatus.EntityData)
}

// CISCORFMIB_Crfcfg
type CISCORFMIB_Crfcfg struct {
    EntityData types.CommonEntityData
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

func (crfcfg *CISCORFMIB_Crfcfg) GetEntityData() *types.CommonEntityData {
    crfcfg.EntityData.YFilter = crfcfg.YFilter
    crfcfg.EntityData.YangName = "cRFCfg"
    crfcfg.EntityData.BundleName = "cisco_ios_xe"
    crfcfg.EntityData.ParentYangName = "CISCO-RF-MIB"
    crfcfg.EntityData.SegmentPath = "cRFCfg"
    crfcfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfcfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfcfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfcfg.EntityData.Children = make(map[string]types.YChild)
    crfcfg.EntityData.Leafs = make(map[string]types.YLeaf)
    crfcfg.EntityData.Leafs["cRFCfgSplitMode"] = types.YLeaf{"Crfcfgsplitmode", crfcfg.Crfcfgsplitmode}
    crfcfg.EntityData.Leafs["cRFCfgKeepaliveThresh"] = types.YLeaf{"Crfcfgkeepalivethresh", crfcfg.Crfcfgkeepalivethresh}
    crfcfg.EntityData.Leafs["cRFCfgKeepaliveThreshMin"] = types.YLeaf{"Crfcfgkeepalivethreshmin", crfcfg.Crfcfgkeepalivethreshmin}
    crfcfg.EntityData.Leafs["cRFCfgKeepaliveThreshMax"] = types.YLeaf{"Crfcfgkeepalivethreshmax", crfcfg.Crfcfgkeepalivethreshmax}
    crfcfg.EntityData.Leafs["cRFCfgKeepaliveTimer"] = types.YLeaf{"Crfcfgkeepalivetimer", crfcfg.Crfcfgkeepalivetimer}
    crfcfg.EntityData.Leafs["cRFCfgKeepaliveTimerMin"] = types.YLeaf{"Crfcfgkeepalivetimermin", crfcfg.Crfcfgkeepalivetimermin}
    crfcfg.EntityData.Leafs["cRFCfgKeepaliveTimerMax"] = types.YLeaf{"Crfcfgkeepalivetimermax", crfcfg.Crfcfgkeepalivetimermax}
    crfcfg.EntityData.Leafs["cRFCfgNotifTimer"] = types.YLeaf{"Crfcfgnotiftimer", crfcfg.Crfcfgnotiftimer}
    crfcfg.EntityData.Leafs["cRFCfgNotifTimerMin"] = types.YLeaf{"Crfcfgnotiftimermin", crfcfg.Crfcfgnotiftimermin}
    crfcfg.EntityData.Leafs["cRFCfgNotifTimerMax"] = types.YLeaf{"Crfcfgnotiftimermax", crfcfg.Crfcfgnotiftimermax}
    crfcfg.EntityData.Leafs["cRFCfgAdminAction"] = types.YLeaf{"Crfcfgadminaction", crfcfg.Crfcfgadminaction}
    crfcfg.EntityData.Leafs["cRFCfgNotifsEnabled"] = types.YLeaf{"Crfcfgnotifsenabled", crfcfg.Crfcfgnotifsenabled}
    crfcfg.EntityData.Leafs["cRFCfgMaintenanceMode"] = types.YLeaf{"Crfcfgmaintenancemode", crfcfg.Crfcfgmaintenancemode}
    crfcfg.EntityData.Leafs["cRFCfgRedundancyMode"] = types.YLeaf{"Crfcfgredundancymode", crfcfg.Crfcfgredundancymode}
    crfcfg.EntityData.Leafs["cRFCfgRedundancyModeDescr"] = types.YLeaf{"Crfcfgredundancymodedescr", crfcfg.Crfcfgredundancymodedescr}
    crfcfg.EntityData.Leafs["cRFCfgRedundancyOperMode"] = types.YLeaf{"Crfcfgredundancyopermode", crfcfg.Crfcfgredundancyopermode}
    return &(crfcfg.EntityData)
}

// CISCORFMIB_Crfhistory
type CISCORFMIB_Crfhistory struct {
    EntityData types.CommonEntityData
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

func (crfhistory *CISCORFMIB_Crfhistory) GetEntityData() *types.CommonEntityData {
    crfhistory.EntityData.YFilter = crfhistory.YFilter
    crfhistory.EntityData.YangName = "cRFHistory"
    crfhistory.EntityData.BundleName = "cisco_ios_xe"
    crfhistory.EntityData.ParentYangName = "CISCO-RF-MIB"
    crfhistory.EntityData.SegmentPath = "cRFHistory"
    crfhistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfhistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfhistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfhistory.EntityData.Children = make(map[string]types.YChild)
    crfhistory.EntityData.Leafs = make(map[string]types.YLeaf)
    crfhistory.EntityData.Leafs["cRFHistoryTableMaxLength"] = types.YLeaf{"Crfhistorytablemaxlength", crfhistory.Crfhistorytablemaxlength}
    crfhistory.EntityData.Leafs["cRFHistoryColdStarts"] = types.YLeaf{"Crfhistorycoldstarts", crfhistory.Crfhistorycoldstarts}
    crfhistory.EntityData.Leafs["cRFHistoryStandByAvailTime"] = types.YLeaf{"Crfhistorystandbyavailtime", crfhistory.Crfhistorystandbyavailtime}
    return &(crfhistory.EntityData)
}

// CISCORFMIB_Crfstatusrfmodecapstable
// This table containing a list of redundancy modes that can be
// supported on the device.
type CISCORFMIB_Crfstatusrfmodecapstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the device implementation specific terminology
    // associated with the redundancy mode that can be supported on the device.
    // The type is slice of
    // CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry.
    Crfstatusrfmodecapsentry []CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry
}

func (crfstatusrfmodecapstable *CISCORFMIB_Crfstatusrfmodecapstable) GetEntityData() *types.CommonEntityData {
    crfstatusrfmodecapstable.EntityData.YFilter = crfstatusrfmodecapstable.YFilter
    crfstatusrfmodecapstable.EntityData.YangName = "cRFStatusRFModeCapsTable"
    crfstatusrfmodecapstable.EntityData.BundleName = "cisco_ios_xe"
    crfstatusrfmodecapstable.EntityData.ParentYangName = "CISCO-RF-MIB"
    crfstatusrfmodecapstable.EntityData.SegmentPath = "cRFStatusRFModeCapsTable"
    crfstatusrfmodecapstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfstatusrfmodecapstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfstatusrfmodecapstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfstatusrfmodecapstable.EntityData.Children = make(map[string]types.YChild)
    crfstatusrfmodecapstable.EntityData.Children["cRFStatusRFModeCapsEntry"] = types.YChild{"Crfstatusrfmodecapsentry", nil}
    for i := range crfstatusrfmodecapstable.Crfstatusrfmodecapsentry {
        crfstatusrfmodecapstable.EntityData.Children[types.GetSegmentPath(&crfstatusrfmodecapstable.Crfstatusrfmodecapsentry[i])] = types.YChild{"Crfstatusrfmodecapsentry", &crfstatusrfmodecapstable.Crfstatusrfmodecapsentry[i]}
    }
    crfstatusrfmodecapstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(crfstatusrfmodecapstable.EntityData)
}

// CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry
// An entry containing the device implementation specific
// terminology associated with the redundancy mode that can be
// supported on the device.
type CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The redundancy mode that can be supported on the
    // device. The type is RFMode.
    Crfstatusrfmodecapsmode interface{}

    // The description of the device implementation specific terminology
    // associated with its supported redundancy mode. The type is string.
    Crfstatusrfmodecapsmodedescr interface{}
}

func (crfstatusrfmodecapsentry *CISCORFMIB_Crfstatusrfmodecapstable_Crfstatusrfmodecapsentry) GetEntityData() *types.CommonEntityData {
    crfstatusrfmodecapsentry.EntityData.YFilter = crfstatusrfmodecapsentry.YFilter
    crfstatusrfmodecapsentry.EntityData.YangName = "cRFStatusRFModeCapsEntry"
    crfstatusrfmodecapsentry.EntityData.BundleName = "cisco_ios_xe"
    crfstatusrfmodecapsentry.EntityData.ParentYangName = "cRFStatusRFModeCapsTable"
    crfstatusrfmodecapsentry.EntityData.SegmentPath = "cRFStatusRFModeCapsEntry" + "[cRFStatusRFModeCapsMode='" + fmt.Sprintf("%v", crfstatusrfmodecapsentry.Crfstatusrfmodecapsmode) + "']"
    crfstatusrfmodecapsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfstatusrfmodecapsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfstatusrfmodecapsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfstatusrfmodecapsentry.EntityData.Children = make(map[string]types.YChild)
    crfstatusrfmodecapsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    crfstatusrfmodecapsentry.EntityData.Leafs["cRFStatusRFModeCapsMode"] = types.YLeaf{"Crfstatusrfmodecapsmode", crfstatusrfmodecapsentry.Crfstatusrfmodecapsmode}
    crfstatusrfmodecapsentry.EntityData.Leafs["cRFStatusRFModeCapsModeDescr"] = types.YLeaf{"Crfstatusrfmodecapsmodedescr", crfstatusrfmodecapsentry.Crfstatusrfmodecapsmodedescr}
    return &(crfstatusrfmodecapsentry.EntityData)
}

// CISCORFMIB_Crfhistoryswitchovertable
// A table that tracks the history of all switchovers that
// have occurred since system initialization. The maximum
// number of entries permissible in this table is defined by
// cRFHistoryTableMaxLength. When the number of entries in
// the table reaches the maximum limit, the next entry
// would replace the oldest existing entry in the table.
type CISCORFMIB_Crfhistoryswitchovertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The entries in this table contain the switchover information. Each entry in
    // the table is indexed by cRFHistorySwitchOverIndex. The index wraps around
    // to 1 after reaching the maximum value. The type is slice of
    // CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry.
    Crfhistoryswitchoverentry []CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry
}

func (crfhistoryswitchovertable *CISCORFMIB_Crfhistoryswitchovertable) GetEntityData() *types.CommonEntityData {
    crfhistoryswitchovertable.EntityData.YFilter = crfhistoryswitchovertable.YFilter
    crfhistoryswitchovertable.EntityData.YangName = "cRFHistorySwitchOverTable"
    crfhistoryswitchovertable.EntityData.BundleName = "cisco_ios_xe"
    crfhistoryswitchovertable.EntityData.ParentYangName = "CISCO-RF-MIB"
    crfhistoryswitchovertable.EntityData.SegmentPath = "cRFHistorySwitchOverTable"
    crfhistoryswitchovertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfhistoryswitchovertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfhistoryswitchovertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfhistoryswitchovertable.EntityData.Children = make(map[string]types.YChild)
    crfhistoryswitchovertable.EntityData.Children["cRFHistorySwitchOverEntry"] = types.YChild{"Crfhistoryswitchoverentry", nil}
    for i := range crfhistoryswitchovertable.Crfhistoryswitchoverentry {
        crfhistoryswitchovertable.EntityData.Children[types.GetSegmentPath(&crfhistoryswitchovertable.Crfhistoryswitchoverentry[i])] = types.YChild{"Crfhistoryswitchoverentry", &crfhistoryswitchovertable.Crfhistoryswitchoverentry[i]}
    }
    crfhistoryswitchovertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(crfhistoryswitchovertable.EntityData)
}

// CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry
// The entries in this table contain the switchover
// information. Each entry in the table is indexed by
// cRFHistorySwitchOverIndex. The index wraps around to 1
// after reaching the maximum value.
type CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry struct {
    EntityData types.CommonEntityData
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

func (crfhistoryswitchoverentry *CISCORFMIB_Crfhistoryswitchovertable_Crfhistoryswitchoverentry) GetEntityData() *types.CommonEntityData {
    crfhistoryswitchoverentry.EntityData.YFilter = crfhistoryswitchoverentry.YFilter
    crfhistoryswitchoverentry.EntityData.YangName = "cRFHistorySwitchOverEntry"
    crfhistoryswitchoverentry.EntityData.BundleName = "cisco_ios_xe"
    crfhistoryswitchoverentry.EntityData.ParentYangName = "cRFHistorySwitchOverTable"
    crfhistoryswitchoverentry.EntityData.SegmentPath = "cRFHistorySwitchOverEntry" + "[cRFHistorySwitchOverIndex='" + fmt.Sprintf("%v", crfhistoryswitchoverentry.Crfhistoryswitchoverindex) + "']"
    crfhistoryswitchoverentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfhistoryswitchoverentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfhistoryswitchoverentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfhistoryswitchoverentry.EntityData.Children = make(map[string]types.YChild)
    crfhistoryswitchoverentry.EntityData.Leafs = make(map[string]types.YLeaf)
    crfhistoryswitchoverentry.EntityData.Leafs["cRFHistorySwitchOverIndex"] = types.YLeaf{"Crfhistoryswitchoverindex", crfhistoryswitchoverentry.Crfhistoryswitchoverindex}
    crfhistoryswitchoverentry.EntityData.Leafs["cRFHistoryPrevActiveUnitId"] = types.YLeaf{"Crfhistoryprevactiveunitid", crfhistoryswitchoverentry.Crfhistoryprevactiveunitid}
    crfhistoryswitchoverentry.EntityData.Leafs["cRFHistoryCurrActiveUnitId"] = types.YLeaf{"Crfhistorycurractiveunitid", crfhistoryswitchoverentry.Crfhistorycurractiveunitid}
    crfhistoryswitchoverentry.EntityData.Leafs["cRFHistorySwitchOverReason"] = types.YLeaf{"Crfhistoryswitchoverreason", crfhistoryswitchoverentry.Crfhistoryswitchoverreason}
    crfhistoryswitchoverentry.EntityData.Leafs["cRFHistorySwactTime"] = types.YLeaf{"Crfhistoryswacttime", crfhistoryswitchoverentry.Crfhistoryswacttime}
    return &(crfhistoryswitchoverentry.EntityData)
}

// CISCORFMIB_Crfstatusrfclienttable
// This table contains a list of RF clients that are
// registered on the device. 
// 
// RF clients are applications that have registered with 
// the Redundancy Facility (RF) to receive RF events and 
// notifications. The purpose of RF clients is to synchronize 
// any relevant data with the standby unit.
type CISCORFMIB_Crfstatusrfclienttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing information on various clients registered with the
    // Redundancy Facility (RF). Entries in this table are always created by the
    // system.  An entry is created in this table when a redundancy aware 
    // application registers with the Redundancy Facility. The entry  is destroyed
    // when that application deregisters from the  Redundancy Facility. The type
    // is slice of CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry.
    Crfstatusrfcliententry []CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry
}

func (crfstatusrfclienttable *CISCORFMIB_Crfstatusrfclienttable) GetEntityData() *types.CommonEntityData {
    crfstatusrfclienttable.EntityData.YFilter = crfstatusrfclienttable.YFilter
    crfstatusrfclienttable.EntityData.YangName = "cRFStatusRFClientTable"
    crfstatusrfclienttable.EntityData.BundleName = "cisco_ios_xe"
    crfstatusrfclienttable.EntityData.ParentYangName = "CISCO-RF-MIB"
    crfstatusrfclienttable.EntityData.SegmentPath = "cRFStatusRFClientTable"
    crfstatusrfclienttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfstatusrfclienttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfstatusrfclienttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfstatusrfclienttable.EntityData.Children = make(map[string]types.YChild)
    crfstatusrfclienttable.EntityData.Children["cRFStatusRFClientEntry"] = types.YChild{"Crfstatusrfcliententry", nil}
    for i := range crfstatusrfclienttable.Crfstatusrfcliententry {
        crfstatusrfclienttable.EntityData.Children[types.GetSegmentPath(&crfstatusrfclienttable.Crfstatusrfcliententry[i])] = types.YChild{"Crfstatusrfcliententry", &crfstatusrfclienttable.Crfstatusrfcliententry[i]}
    }
    crfstatusrfclienttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(crfstatusrfclienttable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (crfstatusrfcliententry *CISCORFMIB_Crfstatusrfclienttable_Crfstatusrfcliententry) GetEntityData() *types.CommonEntityData {
    crfstatusrfcliententry.EntityData.YFilter = crfstatusrfcliententry.YFilter
    crfstatusrfcliententry.EntityData.YangName = "cRFStatusRFClientEntry"
    crfstatusrfcliententry.EntityData.BundleName = "cisco_ios_xe"
    crfstatusrfcliententry.EntityData.ParentYangName = "cRFStatusRFClientTable"
    crfstatusrfcliententry.EntityData.SegmentPath = "cRFStatusRFClientEntry" + "[cRFStatusRFClientID='" + fmt.Sprintf("%v", crfstatusrfcliententry.Crfstatusrfclientid) + "']"
    crfstatusrfcliententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    crfstatusrfcliententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    crfstatusrfcliententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    crfstatusrfcliententry.EntityData.Children = make(map[string]types.YChild)
    crfstatusrfcliententry.EntityData.Leafs = make(map[string]types.YLeaf)
    crfstatusrfcliententry.EntityData.Leafs["cRFStatusRFClientID"] = types.YLeaf{"Crfstatusrfclientid", crfstatusrfcliententry.Crfstatusrfclientid}
    crfstatusrfcliententry.EntityData.Leafs["cRFStatusRFClientDescr"] = types.YLeaf{"Crfstatusrfclientdescr", crfstatusrfcliententry.Crfstatusrfclientdescr}
    crfstatusrfcliententry.EntityData.Leafs["cRFStatusRFClientSeq"] = types.YLeaf{"Crfstatusrfclientseq", crfstatusrfcliententry.Crfstatusrfclientseq}
    crfstatusrfcliententry.EntityData.Leafs["cRFStatusRFClientRedTime"] = types.YLeaf{"Crfstatusrfclientredtime", crfstatusrfcliententry.Crfstatusrfclientredtime}
    crfstatusrfcliententry.EntityData.Leafs["cRFStatusRFClientStatus"] = types.YLeaf{"Crfstatusrfclientstatus", crfstatusrfcliententry.Crfstatusrfclientstatus}
    return &(crfstatusrfcliententry.EntityData)
}

