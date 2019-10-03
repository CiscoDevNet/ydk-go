// Management Information Base module for LLDP configuration,
// statistics, local system data and remote systems data
// components.
// 
// Copyright (C) IEEE (2005).  This version of this MIB module
// is published as subclause 12.1 of IEEE Std 802.1AB-2005;
// see the standard itself for full legal notices.
package lldp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lldp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:LLDP-MIB LLDP-MIB}", reflect.TypeOf(LLDPMIB{}))
    ydk.RegisterEntity("LLDP-MIB:LLDP-MIB", reflect.TypeOf(LLDPMIB{}))
}

// LldpPortIdSubtype represents based on a value locally assigned.
type LldpPortIdSubtype string

const (
    LldpPortIdSubtype_interfaceAlias LldpPortIdSubtype = "interfaceAlias"

    LldpPortIdSubtype_portComponent LldpPortIdSubtype = "portComponent"

    LldpPortIdSubtype_macAddress LldpPortIdSubtype = "macAddress"

    LldpPortIdSubtype_networkAddress LldpPortIdSubtype = "networkAddress"

    LldpPortIdSubtype_interfaceName LldpPortIdSubtype = "interfaceName"

    LldpPortIdSubtype_agentCircuitId LldpPortIdSubtype = "agentCircuitId"

    LldpPortIdSubtype_local LldpPortIdSubtype = "local"
)

// LldpChassisIdSubtype represents based on a locally defined value.
type LldpChassisIdSubtype string

const (
    LldpChassisIdSubtype_chassisComponent LldpChassisIdSubtype = "chassisComponent"

    LldpChassisIdSubtype_interfaceAlias LldpChassisIdSubtype = "interfaceAlias"

    LldpChassisIdSubtype_portComponent LldpChassisIdSubtype = "portComponent"

    LldpChassisIdSubtype_macAddress LldpChassisIdSubtype = "macAddress"

    LldpChassisIdSubtype_networkAddress LldpChassisIdSubtype = "networkAddress"

    LldpChassisIdSubtype_interfaceName LldpChassisIdSubtype = "interfaceName"

    LldpChassisIdSubtype_local LldpChassisIdSubtype = "local"
)

// LldpManAddrIfSubtype represents identifier based on the system port numbering convention.
type LldpManAddrIfSubtype string

const (
    LldpManAddrIfSubtype_unknown LldpManAddrIfSubtype = "unknown"

    LldpManAddrIfSubtype_ifIndex LldpManAddrIfSubtype = "ifIndex"

    LldpManAddrIfSubtype_systemPortNumber LldpManAddrIfSubtype = "systemPortNumber"
)

// LLDPMIB
type LLDPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    LldpConfiguration LLDPMIB_LldpConfiguration

    
    LldpStatistics LLDPMIB_LldpStatistics

    
    LldpLocalSystemData LLDPMIB_LldpLocalSystemData

    // The table that controls LLDP frame transmission on individual ports.
    LldpPortConfigTable LLDPMIB_LldpPortConfigTable

    // A table containing LLDP transmission statistics for individual ports. 
    // Entries are not required to exist in this table while the
    // lldpPortConfigEntry object is equal to 'disabled(4)'.
    LldpStatsTxPortTable LLDPMIB_LldpStatsTxPortTable

    // A table containing LLDP reception statistics for individual ports.  Entries
    // are not required to exist in this table while the lldpPortConfigEntry
    // object is equal to 'disabled(4)'.
    LldpStatsRxPortTable LLDPMIB_LldpStatsRxPortTable

    // This table contains one or more rows per port information associated with
    // the local system known to this agent.
    LldpLocPortTable LLDPMIB_LldpLocPortTable

    // This table contains management address information on the local system
    // known to this agent.
    LldpLocManAddrTable LLDPMIB_LldpLocManAddrTable

    // This table contains one or more rows per physical network connection known
    // to this agent.  The agent may wish to ensure that only one lldpRemEntry is
    // present for each local port, or it may choose to maintain multiple
    // lldpRemEntries for the same local port.  The following procedure may be
    // used to retrieve remote systems information updates from an LLDP agent:    
    // 1. NMS polls all tables associated with remote systems       and keeps a
    // local copy of the information retrieved.       NMS polls periodically the
    // values of the following       objects:          a.
    // lldpStatsRemTablesInserts          b. lldpStatsRemTablesDeletes          c.
    // lldpStatsRemTablesDrops          d. lldpStatsRemTablesAgeouts          e.
    // lldpStatsRxPortAgeoutsTotal for all ports.     2. LLDP agent updates remote
    // systems MIB objects, and       sends out notifications to a list of
    // notification       destinations.     3. NMS receives the notifications and
    // compares the new       values of objects listed in step 1.         
    // Periodically, NMS should poll the object      
    // lldpStatsRemTablesLastChangeTime to find out if anything       has changed
    // since the last poll.  if something has       changed, NMS will poll the
    // objects listed in step 1 to       figure out what kind of changes occurred
    // in the tables.        if value of lldpStatsRemTablesInserts has changed,   
    // then NMS will walk all tables by employing TimeFilter       with the
    // last-polled time value.  This request will       return new objects or
    // objects whose values are updated       since the last poll.        if value
    // of lldpStatsRemTablesAgeouts has changed,       then NMS will walk the
    // lldpStatsRxPortAgeoutsTotal and       compare the new values with
    // previously recorded ones.       For ports whose lldpStatsRxPortAgeoutsTotal
    // value is       greater than the recorded value, NMS will have to      
    // retrieve objects associated with those ports from       table(s) without
    // employing a TimeFilter (which is       performed by specifying 0 for the
    // TimeFilter.)        lldpStatsRemTablesDeletes and lldpStatsRemTablesDrops  
    // objects are provided for informational purposes.
    LldpRemTable LLDPMIB_LldpRemTable

    // This table contains one or more rows per management address information on
    // the remote system learned on a particular port contained in the local
    // chassis known to this agent.
    LldpRemManAddrTable LLDPMIB_LldpRemManAddrTable

    // This table contains information about an incoming TLV which is not
    // recognized by the receiving LLDP agent.  The TLV may be from a later
    // version of the basic management set.  This table should only contain TLVs
    // that are found in a single LLDP frame.  Entries in this table, associated
    // with an MAC service access point (MSAP, the access point for MAC services
    // provided to the LCC sublayer, defined in IEEE 100, which is also identified
    // with a particular lldpRemLocalPortNum, lldpRemIndex pair) are overwritten
    // with most recently received unrecognized TLV from the same MSAP, or they
    // will naturally age out when the rxInfoTTL timer (associated with the MSAP)
    // expires.
    LldpRemUnknownTLVTable LLDPMIB_LldpRemUnknownTLVTable

    // This table contains one or more rows per physical network connection which
    // advertises the organizationally defined information.  Note that this table
    // contains one or more rows of organizationally defined information that is
    // not recognized by the local agent.  If the local system is capable of
    // recognizing any organizationally defined information, appropriate extension
    // MIBs from the organization should be used for information retrieval.
    LldpRemOrgDefInfoTable LLDPMIB_LldpRemOrgDefInfoTable
}

func (lLDPMIB *LLDPMIB) GetEntityData() *types.CommonEntityData {
    lLDPMIB.EntityData.YFilter = lLDPMIB.YFilter
    lLDPMIB.EntityData.YangName = "LLDP-MIB"
    lLDPMIB.EntityData.BundleName = "cisco_ios_xe"
    lLDPMIB.EntityData.ParentYangName = "LLDP-MIB"
    lLDPMIB.EntityData.SegmentPath = "LLDP-MIB:LLDP-MIB"
    lLDPMIB.EntityData.AbsolutePath = lLDPMIB.EntityData.SegmentPath
    lLDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lLDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lLDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lLDPMIB.EntityData.Children = types.NewOrderedMap()
    lLDPMIB.EntityData.Children.Append("lldpConfiguration", types.YChild{"LldpConfiguration", &lLDPMIB.LldpConfiguration})
    lLDPMIB.EntityData.Children.Append("lldpStatistics", types.YChild{"LldpStatistics", &lLDPMIB.LldpStatistics})
    lLDPMIB.EntityData.Children.Append("lldpLocalSystemData", types.YChild{"LldpLocalSystemData", &lLDPMIB.LldpLocalSystemData})
    lLDPMIB.EntityData.Children.Append("lldpPortConfigTable", types.YChild{"LldpPortConfigTable", &lLDPMIB.LldpPortConfigTable})
    lLDPMIB.EntityData.Children.Append("lldpStatsTxPortTable", types.YChild{"LldpStatsTxPortTable", &lLDPMIB.LldpStatsTxPortTable})
    lLDPMIB.EntityData.Children.Append("lldpStatsRxPortTable", types.YChild{"LldpStatsRxPortTable", &lLDPMIB.LldpStatsRxPortTable})
    lLDPMIB.EntityData.Children.Append("lldpLocPortTable", types.YChild{"LldpLocPortTable", &lLDPMIB.LldpLocPortTable})
    lLDPMIB.EntityData.Children.Append("lldpLocManAddrTable", types.YChild{"LldpLocManAddrTable", &lLDPMIB.LldpLocManAddrTable})
    lLDPMIB.EntityData.Children.Append("lldpRemTable", types.YChild{"LldpRemTable", &lLDPMIB.LldpRemTable})
    lLDPMIB.EntityData.Children.Append("lldpRemManAddrTable", types.YChild{"LldpRemManAddrTable", &lLDPMIB.LldpRemManAddrTable})
    lLDPMIB.EntityData.Children.Append("lldpRemUnknownTLVTable", types.YChild{"LldpRemUnknownTLVTable", &lLDPMIB.LldpRemUnknownTLVTable})
    lLDPMIB.EntityData.Children.Append("lldpRemOrgDefInfoTable", types.YChild{"LldpRemOrgDefInfoTable", &lLDPMIB.LldpRemOrgDefInfoTable})
    lLDPMIB.EntityData.Leafs = types.NewOrderedMap()

    lLDPMIB.EntityData.YListKeys = []string {}

    return &(lLDPMIB.EntityData)
}

// LLDPMIB_LldpConfiguration
type LLDPMIB_LldpConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interval at which LLDP frames are transmitted on behalf of this LLDP
    // agent.  The default value for lldpMessageTxInterval object is 30 seconds. 
    // The value of this object must be restored from non-volatile storage after a
    // re-initialization of the management system. The type is interface{} with
    // range: 5..32768. Units are seconds.
    LldpMessageTxInterval interface{}

    // The time-to-live value expressed as a multiple of the lldpMessageTxInterval
    // object.  The actual time-to-live value used in LLDP frames, transmitted on
    // behalf of this LLDP agent, can be expressed by the following formula: TTL =
    // min(65535, (lldpMessageTxInterval * lldpMessageTxHoldMultiplier)) For
    // example, if the value of lldpMessageTxInterval is '30', and the value of
    // lldpMessageTxHoldMultiplier is '4', then the value '120' is encoded in the
    // TTL field in the LLDP header.  The default value for
    // lldpMessageTxHoldMultiplier object is 4.  The value of this object must be
    // restored from non-volatile storage after a re-initialization of the
    // management system. The type is interface{} with range: 2..10.
    LldpMessageTxHoldMultiplier interface{}

    // The lldpReinitDelay indicates the delay (in units of seconds) from when
    // lldpPortConfigAdminStatus object of a particular port becomes 'disabled'
    // until re-initialization will be attempted.  The default value for
    // lldpReintDelay object is two seconds.  The value of this object must be
    // restored from non-volatile storage after a re-initialization of the
    // management system. The type is interface{} with range: 1..10. Units are
    // seconds.
    LldpReinitDelay interface{}

    // The lldpTxDelay indicates the delay (in units of seconds) between
    // successive LLDP frame transmissions  initiated by value/status changes in
    // the LLDP local systems MIB.  The recommended value for the lldpTxDelay is
    // set by the following  formula:     1 <= lldpTxDelay <= (0.25 *
    // lldpMessageTxInterval)  The default value for lldpTxDelay object is two
    // seconds.  The value of this object must be restored from non-volatile
    // storage after a re-initialization of the management system. The type is
    // interface{} with range: 1..8192. Units are seconds.
    LldpTxDelay interface{}

    // This object controls the transmission of LLDP notifications.  the agent
    // must not generate more than one lldpRemTablesChange notification-event in
    // the indicated period, where a 'notification-event' is the transmission of a
    // single notification PDU type to a list of notification destinations. If
    // additional changes in lldpRemoteSystemsData object groups occur within the
    // indicated throttling period, then these trap- events must be suppressed by
    // the agent. An NMS should periodically check the value of
    // lldpStatsRemTableLastChangeTime to detect any missed lldpRemTablesChange
    // notification-events, e.g. due to throttling or transmission loss.  If
    // notification transmission is enabled for particular ports, the suggested
    // default throttling period is 5 seconds.  The value of this object must be
    // restored from non-volatile storage after a re-initialization of the
    // management system. The type is interface{} with range: 5..3600. Units are
    // seconds.
    LldpNotificationInterval interface{}
}

func (lldpConfiguration *LLDPMIB_LldpConfiguration) GetEntityData() *types.CommonEntityData {
    lldpConfiguration.EntityData.YFilter = lldpConfiguration.YFilter
    lldpConfiguration.EntityData.YangName = "lldpConfiguration"
    lldpConfiguration.EntityData.BundleName = "cisco_ios_xe"
    lldpConfiguration.EntityData.ParentYangName = "LLDP-MIB"
    lldpConfiguration.EntityData.SegmentPath = "lldpConfiguration"
    lldpConfiguration.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpConfiguration.EntityData.SegmentPath
    lldpConfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpConfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpConfiguration.EntityData.Children = types.NewOrderedMap()
    lldpConfiguration.EntityData.Leafs = types.NewOrderedMap()
    lldpConfiguration.EntityData.Leafs.Append("lldpMessageTxInterval", types.YLeaf{"LldpMessageTxInterval", lldpConfiguration.LldpMessageTxInterval})
    lldpConfiguration.EntityData.Leafs.Append("lldpMessageTxHoldMultiplier", types.YLeaf{"LldpMessageTxHoldMultiplier", lldpConfiguration.LldpMessageTxHoldMultiplier})
    lldpConfiguration.EntityData.Leafs.Append("lldpReinitDelay", types.YLeaf{"LldpReinitDelay", lldpConfiguration.LldpReinitDelay})
    lldpConfiguration.EntityData.Leafs.Append("lldpTxDelay", types.YLeaf{"LldpTxDelay", lldpConfiguration.LldpTxDelay})
    lldpConfiguration.EntityData.Leafs.Append("lldpNotificationInterval", types.YLeaf{"LldpNotificationInterval", lldpConfiguration.LldpNotificationInterval})

    lldpConfiguration.EntityData.YListKeys = []string {}

    return &(lldpConfiguration.EntityData)
}

// LLDPMIB_LldpStatistics
type LLDPMIB_LldpStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime object (defined in IETF RFC 3418) at the time an
    // entry is created, modified, or deleted in the in tables associated with the
    // lldpRemoteSystemsData objects and all LLDP extension objects associated
    // with remote systems.  An NMS can use this object to reduce polling of the
    // lldpRemoteSystemsData objects. The type is interface{} with range:
    // 0..4294967295.
    LldpStatsRemTablesLastChangeTime interface{}

    // The number of times the complete set of information advertised by a
    // particular MSAP has been inserted into tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects.  The complete set of
    // information received from a particular MSAP should be inserted into related
    // tables.  If partial information cannot be inserted for a reason such as
    // lack of resources, all of the complete set of information should be
    // removed.  This counter should be incremented only once after the complete
    // set of information is successfully recorded in all related tables.  Any
    // failures during inserting information set which result in deletion of
    // previously inserted information should not trigger any changes in
    // lldpStatsRemTablesInserts since the insert is not completed yet or or in
    // lldpStatsRemTablesDeletes, since the deletion would only be a partial
    // deletion. If the failure was the result of lack of resources, the
    // lldpStatsRemTablesDrops counter should be incremented once. The type is
    // interface{} with range: 0..4294967295. Units are table entries.
    LldpStatsRemTablesInserts interface{}

    // The number of times the complete set of information advertised by a
    // particular MSAP has been deleted from tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects.  This counter should be
    // incremented only once when the complete set of information is completely
    // deleted from all related tables.  Partial deletions, such as deletion of
    // rows associated with a particular MSAP from some tables, but not from all
    // tables are not allowed, thus should not change the value of this counter.
    // The type is interface{} with range: 0..4294967295. Units are table entries.
    LldpStatsRemTablesDeletes interface{}

    // The number of times the complete set of information advertised by a
    // particular MSAP could not be entered into tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects because of insufficient
    // resources. The type is interface{} with range: 0..4294967295. Units are
    // table entries.
    LldpStatsRemTablesDrops interface{}

    // The number of times the complete set of information advertised by a
    // particular MSAP has been deleted from tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects because the information
    // timeliness interval has expired.  This counter should be incremented only
    // once when the complete set of information is completely invalidated (aged
    // out) from all related tables.  Partial aging, similar to deletion case, is
    // not allowed, and thus, should not change the value of this counter. The
    // type is interface{} with range: 0..4294967295.
    LldpStatsRemTablesAgeouts interface{}
}

func (lldpStatistics *LLDPMIB_LldpStatistics) GetEntityData() *types.CommonEntityData {
    lldpStatistics.EntityData.YFilter = lldpStatistics.YFilter
    lldpStatistics.EntityData.YangName = "lldpStatistics"
    lldpStatistics.EntityData.BundleName = "cisco_ios_xe"
    lldpStatistics.EntityData.ParentYangName = "LLDP-MIB"
    lldpStatistics.EntityData.SegmentPath = "lldpStatistics"
    lldpStatistics.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpStatistics.EntityData.SegmentPath
    lldpStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpStatistics.EntityData.Children = types.NewOrderedMap()
    lldpStatistics.EntityData.Leafs = types.NewOrderedMap()
    lldpStatistics.EntityData.Leafs.Append("lldpStatsRemTablesLastChangeTime", types.YLeaf{"LldpStatsRemTablesLastChangeTime", lldpStatistics.LldpStatsRemTablesLastChangeTime})
    lldpStatistics.EntityData.Leafs.Append("lldpStatsRemTablesInserts", types.YLeaf{"LldpStatsRemTablesInserts", lldpStatistics.LldpStatsRemTablesInserts})
    lldpStatistics.EntityData.Leafs.Append("lldpStatsRemTablesDeletes", types.YLeaf{"LldpStatsRemTablesDeletes", lldpStatistics.LldpStatsRemTablesDeletes})
    lldpStatistics.EntityData.Leafs.Append("lldpStatsRemTablesDrops", types.YLeaf{"LldpStatsRemTablesDrops", lldpStatistics.LldpStatsRemTablesDrops})
    lldpStatistics.EntityData.Leafs.Append("lldpStatsRemTablesAgeouts", types.YLeaf{"LldpStatsRemTablesAgeouts", lldpStatistics.LldpStatsRemTablesAgeouts})

    lldpStatistics.EntityData.YListKeys = []string {}

    return &(lldpStatistics.EntityData)
}

// LLDPMIB_LldpLocalSystemData
type LLDPMIB_LldpLocalSystemData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of encoding used to identify the chassis associated with the local
    // system. The type is LldpChassisIdSubtype.
    LldpLocChassisIdSubtype interface{}

    // The string value used to identify the chassis component associated with the
    // local system. The type is string with length: 1..255.
    LldpLocChassisId interface{}

    // The string value used to identify the system name of the local system.  If
    // the local agent supports IETF RFC 3418, lldpLocSysName object should have
    // the same value of sysName object. The type is string with length: 0..255.
    LldpLocSysName interface{}

    // The string value used to identify the system description of the local
    // system.  If the local agent supports IETF RFC 3418, lldpLocSysDesc object
    // should have the same value of sysDesc object. The type is string with
    // length: 0..255.
    LldpLocSysDesc interface{}

    // The bitmap value used to identify which system capabilities are supported
    // on the local system. The type is map[string]bool.
    LldpLocSysCapSupported interface{}

    // The bitmap value used to identify which system capabilities are enabled on
    // the local system. The type is map[string]bool.
    LldpLocSysCapEnabled interface{}
}

func (lldpLocalSystemData *LLDPMIB_LldpLocalSystemData) GetEntityData() *types.CommonEntityData {
    lldpLocalSystemData.EntityData.YFilter = lldpLocalSystemData.YFilter
    lldpLocalSystemData.EntityData.YangName = "lldpLocalSystemData"
    lldpLocalSystemData.EntityData.BundleName = "cisco_ios_xe"
    lldpLocalSystemData.EntityData.ParentYangName = "LLDP-MIB"
    lldpLocalSystemData.EntityData.SegmentPath = "lldpLocalSystemData"
    lldpLocalSystemData.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpLocalSystemData.EntityData.SegmentPath
    lldpLocalSystemData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpLocalSystemData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpLocalSystemData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpLocalSystemData.EntityData.Children = types.NewOrderedMap()
    lldpLocalSystemData.EntityData.Leafs = types.NewOrderedMap()
    lldpLocalSystemData.EntityData.Leafs.Append("lldpLocChassisIdSubtype", types.YLeaf{"LldpLocChassisIdSubtype", lldpLocalSystemData.LldpLocChassisIdSubtype})
    lldpLocalSystemData.EntityData.Leafs.Append("lldpLocChassisId", types.YLeaf{"LldpLocChassisId", lldpLocalSystemData.LldpLocChassisId})
    lldpLocalSystemData.EntityData.Leafs.Append("lldpLocSysName", types.YLeaf{"LldpLocSysName", lldpLocalSystemData.LldpLocSysName})
    lldpLocalSystemData.EntityData.Leafs.Append("lldpLocSysDesc", types.YLeaf{"LldpLocSysDesc", lldpLocalSystemData.LldpLocSysDesc})
    lldpLocalSystemData.EntityData.Leafs.Append("lldpLocSysCapSupported", types.YLeaf{"LldpLocSysCapSupported", lldpLocalSystemData.LldpLocSysCapSupported})
    lldpLocalSystemData.EntityData.Leafs.Append("lldpLocSysCapEnabled", types.YLeaf{"LldpLocSysCapEnabled", lldpLocalSystemData.LldpLocSysCapEnabled})

    lldpLocalSystemData.EntityData.YListKeys = []string {}

    return &(lldpLocalSystemData.EntityData)
}

// LLDPMIB_LldpPortConfigTable
// The table that controls LLDP frame transmission on individual
// ports.
type LLDPMIB_LldpPortConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LLDP configuration information for a particular port. This configuration
    // parameter controls the transmission and the reception of LLDP frames on
    // those ports whose rows are created in this table. The type is slice of
    // LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry.
    LldpPortConfigEntry []*LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry
}

func (lldpPortConfigTable *LLDPMIB_LldpPortConfigTable) GetEntityData() *types.CommonEntityData {
    lldpPortConfigTable.EntityData.YFilter = lldpPortConfigTable.YFilter
    lldpPortConfigTable.EntityData.YangName = "lldpPortConfigTable"
    lldpPortConfigTable.EntityData.BundleName = "cisco_ios_xe"
    lldpPortConfigTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpPortConfigTable.EntityData.SegmentPath = "lldpPortConfigTable"
    lldpPortConfigTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpPortConfigTable.EntityData.SegmentPath
    lldpPortConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpPortConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpPortConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpPortConfigTable.EntityData.Children = types.NewOrderedMap()
    lldpPortConfigTable.EntityData.Children.Append("lldpPortConfigEntry", types.YChild{"LldpPortConfigEntry", nil})
    for i := range lldpPortConfigTable.LldpPortConfigEntry {
        lldpPortConfigTable.EntityData.Children.Append(types.GetSegmentPath(lldpPortConfigTable.LldpPortConfigEntry[i]), types.YChild{"LldpPortConfigEntry", lldpPortConfigTable.LldpPortConfigEntry[i]})
    }
    lldpPortConfigTable.EntityData.Leafs = types.NewOrderedMap()

    lldpPortConfigTable.EntityData.YListKeys = []string {}

    return &(lldpPortConfigTable.EntityData)
}

// LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry
// LLDP configuration information for a particular port.
// This configuration parameter controls the transmission and
// the reception of LLDP frames on those ports whose rows are
// created in this table.
type LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpPortConfigTable. The type is interface{} with range: 1..4096.
    LldpPortConfigPortNum interface{}

    // The administratively desired status of the local LLDP agent.  If the
    // associated lldpPortConfigAdminStatus object has a value of 'txOnly(1)',
    // then LLDP agent will transmit LLDP frames on this port and it will not
    // store any information about the remote systems connected.  If the
    // associated lldpPortConfigAdminStatus object has a value of 'rxOnly(2)',
    // then the LLDP agent will receive, but it will not transmit LLDP frames on
    // this port.  If the associated lldpPortConfigAdminStatus object has a value
    // of 'txAndRx(3)', then the LLDP agent will transmit and receive LLDP frames
    // on this port.  If the associated lldpPortConfigAdminStatus object has a
    // value of 'disabled(4)', then LLDP agent will not transmit or receive LLDP
    // frames on this port.  If there is remote systems information which is
    // received on this port and stored in other tables, before the port's
    // lldpPortConfigAdminStatus becomes disabled, then the information will
    // naturally age out. The type is LldpPortConfigAdminStatus.
    LldpPortConfigAdminStatus interface{}

    // The lldpPortConfigNotificationEnable controls, on a per port basis, 
    // whether or not notifications from the agent are enabled. The value true(1)
    // means that notifications are enabled; the value false(2) means that they
    // are not. The type is bool.
    LldpPortConfigNotificationEnable interface{}

    // The lldpPortConfigTLVsTxEnable, defined as a bitmap, includes the basic set
    // of LLDP TLVs whose transmission is allowed on the local LLDP agent by the
    // network management. Each bit in the bitmap corresponds to a TLV type
    // associated with a specific optional TLV.  It should be noted that the
    // organizationally-specific TLVs are excluded from the lldpTLVsTxEnable
    // bitmap.  LLDP Organization Specific Information Extension MIBs should have
    // similar configuration object to control transmission of their
    // organizationally defined TLVs.  The bit 'portDesc(0)' indicates that LLDP
    // agent should transmit 'Port Description TLV'.  The bit 'sysName(1)'
    // indicates that LLDP agent should transmit 'System Name TLV'.  The bit
    // 'sysDesc(2)' indicates that LLDP agent should transmit 'System Description
    // TLV'.  The bit 'sysCap(3)' indicates that LLDP agent should transmit
    // 'System Capabilities TLV'.  There is no bit reserved for the management
    // address TLV type since transmission of management address TLVs are
    // controlled by another object, lldpConfigManAddrTable.  The default value
    // for lldpPortConfigTLVsTxEnable object is empty set, which means no
    // enumerated values are set.  The value of this object must be restored from
    // non-volatile storage after a re-initialization of the management system.
    // The type is map[string]bool.
    LldpPortConfigTLVsTxEnable interface{}
}

func (lldpPortConfigEntry *LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry) GetEntityData() *types.CommonEntityData {
    lldpPortConfigEntry.EntityData.YFilter = lldpPortConfigEntry.YFilter
    lldpPortConfigEntry.EntityData.YangName = "lldpPortConfigEntry"
    lldpPortConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpPortConfigEntry.EntityData.ParentYangName = "lldpPortConfigTable"
    lldpPortConfigEntry.EntityData.SegmentPath = "lldpPortConfigEntry" + types.AddKeyToken(lldpPortConfigEntry.LldpPortConfigPortNum, "lldpPortConfigPortNum")
    lldpPortConfigEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpPortConfigTable/" + lldpPortConfigEntry.EntityData.SegmentPath
    lldpPortConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpPortConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpPortConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpPortConfigEntry.EntityData.Children = types.NewOrderedMap()
    lldpPortConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpPortConfigEntry.EntityData.Leafs.Append("lldpPortConfigPortNum", types.YLeaf{"LldpPortConfigPortNum", lldpPortConfigEntry.LldpPortConfigPortNum})
    lldpPortConfigEntry.EntityData.Leafs.Append("lldpPortConfigAdminStatus", types.YLeaf{"LldpPortConfigAdminStatus", lldpPortConfigEntry.LldpPortConfigAdminStatus})
    lldpPortConfigEntry.EntityData.Leafs.Append("lldpPortConfigNotificationEnable", types.YLeaf{"LldpPortConfigNotificationEnable", lldpPortConfigEntry.LldpPortConfigNotificationEnable})
    lldpPortConfigEntry.EntityData.Leafs.Append("lldpPortConfigTLVsTxEnable", types.YLeaf{"LldpPortConfigTLVsTxEnable", lldpPortConfigEntry.LldpPortConfigTLVsTxEnable})

    lldpPortConfigEntry.EntityData.YListKeys = []string {"LldpPortConfigPortNum"}

    return &(lldpPortConfigEntry.EntityData)
}

// LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus represents becomes disabled, then the information will naturally age out.
type LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus string

const (
    LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus_txOnly LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus = "txOnly"

    LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus_rxOnly LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus = "rxOnly"

    LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus_txAndRx LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus = "txAndRx"

    LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus_disabled LLDPMIB_LldpPortConfigTable_LldpPortConfigEntry_LldpPortConfigAdminStatus = "disabled"
)

// LLDPMIB_LldpStatsTxPortTable
// A table containing LLDP transmission statistics for
// individual ports.  Entries are not required to exist in
// this table while the lldpPortConfigEntry object is equal to
// 'disabled(4)'.
type LLDPMIB_LldpStatsTxPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LLDP frame transmission statistics for a particular port. The port must be
    // contained in the same chassis as the LLDP agent.  All counter values in a
    // particular entry shall be maintained on a continuing basis and shall not be
    // deleted upon expiration of rxInfoTTL timing counters in the LLDP remote
    // systems MIB of the receipt of a shutdown frame from a remote LLDP agent. 
    // All statistical counters associated with a particular port on the local
    // LLDP agent become frozen whenever the adminStatus is disabled for the same
    // port. The type is slice of
    // LLDPMIB_LldpStatsTxPortTable_LldpStatsTxPortEntry.
    LldpStatsTxPortEntry []*LLDPMIB_LldpStatsTxPortTable_LldpStatsTxPortEntry
}

func (lldpStatsTxPortTable *LLDPMIB_LldpStatsTxPortTable) GetEntityData() *types.CommonEntityData {
    lldpStatsTxPortTable.EntityData.YFilter = lldpStatsTxPortTable.YFilter
    lldpStatsTxPortTable.EntityData.YangName = "lldpStatsTxPortTable"
    lldpStatsTxPortTable.EntityData.BundleName = "cisco_ios_xe"
    lldpStatsTxPortTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpStatsTxPortTable.EntityData.SegmentPath = "lldpStatsTxPortTable"
    lldpStatsTxPortTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpStatsTxPortTable.EntityData.SegmentPath
    lldpStatsTxPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpStatsTxPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpStatsTxPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpStatsTxPortTable.EntityData.Children = types.NewOrderedMap()
    lldpStatsTxPortTable.EntityData.Children.Append("lldpStatsTxPortEntry", types.YChild{"LldpStatsTxPortEntry", nil})
    for i := range lldpStatsTxPortTable.LldpStatsTxPortEntry {
        lldpStatsTxPortTable.EntityData.Children.Append(types.GetSegmentPath(lldpStatsTxPortTable.LldpStatsTxPortEntry[i]), types.YChild{"LldpStatsTxPortEntry", lldpStatsTxPortTable.LldpStatsTxPortEntry[i]})
    }
    lldpStatsTxPortTable.EntityData.Leafs = types.NewOrderedMap()

    lldpStatsTxPortTable.EntityData.YListKeys = []string {}

    return &(lldpStatsTxPortTable.EntityData)
}

// LLDPMIB_LldpStatsTxPortTable_LldpStatsTxPortEntry
// LLDP frame transmission statistics for a particular port.
// The port must be contained in the same chassis as the
// LLDP agent.
// 
// All counter values in a particular entry shall be
// maintained on a continuing basis and shall not be deleted
// upon expiration of rxInfoTTL timing counters in the LLDP
// remote systems MIB of the receipt of a shutdown frame from
// a remote LLDP agent.
// 
// All statistical counters associated with a particular
// port on the local LLDP agent become frozen whenever the
// adminStatus is disabled for the same port.
type LLDPMIB_LldpStatsTxPortTable_LldpStatsTxPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpStatsTable. The type is interface{} with range: 1..4096.
    LldpStatsTxPortNum interface{}

    // The number of LLDP frames transmitted by this LLDP agent on the indicated
    // port. The type is interface{} with range: 0..4294967295.
    LldpStatsTxPortFramesTotal interface{}
}

func (lldpStatsTxPortEntry *LLDPMIB_LldpStatsTxPortTable_LldpStatsTxPortEntry) GetEntityData() *types.CommonEntityData {
    lldpStatsTxPortEntry.EntityData.YFilter = lldpStatsTxPortEntry.YFilter
    lldpStatsTxPortEntry.EntityData.YangName = "lldpStatsTxPortEntry"
    lldpStatsTxPortEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpStatsTxPortEntry.EntityData.ParentYangName = "lldpStatsTxPortTable"
    lldpStatsTxPortEntry.EntityData.SegmentPath = "lldpStatsTxPortEntry" + types.AddKeyToken(lldpStatsTxPortEntry.LldpStatsTxPortNum, "lldpStatsTxPortNum")
    lldpStatsTxPortEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpStatsTxPortTable/" + lldpStatsTxPortEntry.EntityData.SegmentPath
    lldpStatsTxPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpStatsTxPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpStatsTxPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpStatsTxPortEntry.EntityData.Children = types.NewOrderedMap()
    lldpStatsTxPortEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpStatsTxPortEntry.EntityData.Leafs.Append("lldpStatsTxPortNum", types.YLeaf{"LldpStatsTxPortNum", lldpStatsTxPortEntry.LldpStatsTxPortNum})
    lldpStatsTxPortEntry.EntityData.Leafs.Append("lldpStatsTxPortFramesTotal", types.YLeaf{"LldpStatsTxPortFramesTotal", lldpStatsTxPortEntry.LldpStatsTxPortFramesTotal})

    lldpStatsTxPortEntry.EntityData.YListKeys = []string {"LldpStatsTxPortNum"}

    return &(lldpStatsTxPortEntry.EntityData)
}

// LLDPMIB_LldpStatsRxPortTable
// A table containing LLDP reception statistics for individual
// ports.  Entries are not required to exist in this table while
// the lldpPortConfigEntry object is equal to 'disabled(4)'.
type LLDPMIB_LldpStatsRxPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LLDP frame reception statistics for a particular port. The port must be
    // contained in the same chassis as the LLDP agent.  All counter values in a
    // particular entry shall be maintained on a continuing basis and shall not be
    // deleted upon expiration of rxInfoTTL timing counters in the LLDP remote
    // systems MIB of the receipt of a shutdown frame from a remote LLDP agent. 
    // All statistical counters associated with a particular port on the local
    // LLDP agent become frozen whenever the adminStatus is disabled for the same
    // port. The type is slice of
    // LLDPMIB_LldpStatsRxPortTable_LldpStatsRxPortEntry.
    LldpStatsRxPortEntry []*LLDPMIB_LldpStatsRxPortTable_LldpStatsRxPortEntry
}

func (lldpStatsRxPortTable *LLDPMIB_LldpStatsRxPortTable) GetEntityData() *types.CommonEntityData {
    lldpStatsRxPortTable.EntityData.YFilter = lldpStatsRxPortTable.YFilter
    lldpStatsRxPortTable.EntityData.YangName = "lldpStatsRxPortTable"
    lldpStatsRxPortTable.EntityData.BundleName = "cisco_ios_xe"
    lldpStatsRxPortTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpStatsRxPortTable.EntityData.SegmentPath = "lldpStatsRxPortTable"
    lldpStatsRxPortTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpStatsRxPortTable.EntityData.SegmentPath
    lldpStatsRxPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpStatsRxPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpStatsRxPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpStatsRxPortTable.EntityData.Children = types.NewOrderedMap()
    lldpStatsRxPortTable.EntityData.Children.Append("lldpStatsRxPortEntry", types.YChild{"LldpStatsRxPortEntry", nil})
    for i := range lldpStatsRxPortTable.LldpStatsRxPortEntry {
        lldpStatsRxPortTable.EntityData.Children.Append(types.GetSegmentPath(lldpStatsRxPortTable.LldpStatsRxPortEntry[i]), types.YChild{"LldpStatsRxPortEntry", lldpStatsRxPortTable.LldpStatsRxPortEntry[i]})
    }
    lldpStatsRxPortTable.EntityData.Leafs = types.NewOrderedMap()

    lldpStatsRxPortTable.EntityData.YListKeys = []string {}

    return &(lldpStatsRxPortTable.EntityData)
}

// LLDPMIB_LldpStatsRxPortTable_LldpStatsRxPortEntry
// LLDP frame reception statistics for a particular port.
// The port must be contained in the same chassis as the
// LLDP agent.
// 
// All counter values in a particular entry shall be
// maintained on a continuing basis and shall not be deleted
// upon expiration of rxInfoTTL timing counters in the LLDP
// remote systems MIB of the receipt of a shutdown frame from
// a remote LLDP agent.
// 
// All statistical counters associated with a particular
// port on the local LLDP agent become frozen whenever the
// adminStatus is disabled for the same port.
type LLDPMIB_LldpStatsRxPortTable_LldpStatsRxPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpStatsTable. The type is interface{} with range: 1..4096.
    LldpStatsRxPortNum interface{}

    // The number of LLDP frames received by this LLDP agent on the indicated
    // port, and then discarded for any reason. This counter can provide an
    // indication that LLDP header formating problems may exist with the local
    // LLDP agent in the sending system or that LLDPDU validation problems may
    // exist with the local LLDP agent in the receiving system. The type is
    // interface{} with range: 0..4294967295.
    LldpStatsRxPortFramesDiscardedTotal interface{}

    // The number of invalid LLDP frames received by this LLDP agent on the
    // indicated port, while this LLDP agent is enabled. The type is interface{}
    // with range: 0..4294967295.
    LldpStatsRxPortFramesErrors interface{}

    // The number of valid LLDP frames received by this LLDP agent on the
    // indicated port, while this LLDP agent is enabled. The type is interface{}
    // with range: 0..4294967295.
    LldpStatsRxPortFramesTotal interface{}

    // The number of LLDP TLVs discarded for any reason by this LLDP agent on the
    // indicated port. The type is interface{} with range: 0..4294967295.
    LldpStatsRxPortTLVsDiscardedTotal interface{}

    // The number of LLDP TLVs received on the given port that are not recognized
    // by this LLDP agent on the indicated port.  An unrecognized TLV is referred
    // to as the TLV whose type value is in the range of reserved TLV types (000
    // 1001 - 111 1110) in Table 9.1 of IEEE Std 802.1AB-2005.  An unrecognized
    // TLV may be a basic management TLV from a later LLDP version. The type is
    // interface{} with range: 0..4294967295.
    LldpStatsRxPortTLVsUnrecognizedTotal interface{}

    // The counter that represents the number of age-outs that occurred on a given
    // port.  An age-out is the number of times the complete set of information
    // advertised by a particular MSAP has been deleted from tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects because the information
    // timeliness interval has expired.  This counter is similar to
    // lldpStatsRemTablesAgeouts, except that the counter is on a per port basis. 
    // This enables NMS to poll tables associated with the lldpRemoteSystemsData
    // objects and all LLDP extension objects associated with remote systems on
    // the indicated port only.  This counter should be set to zero during agent
    // initialization and its value should not be saved in non-volatile storage.
    // When a port's admin status changes from 'disabled' to 'rxOnly', 'txOnly' or
    // 'txAndRx', the counter associated with the same port should reset to 0. 
    // The agent should also flush all remote system information associated with
    // the same port.  This counter should be incremented only once when the
    // complete set of information is invalidated (aged out) from all related
    // tables on a particular port.  Partial aging is not allowed, and thus,
    // should not change the value of this counter. The type is interface{} with
    // range: 0..4294967295.
    LldpStatsRxPortAgeoutsTotal interface{}
}

func (lldpStatsRxPortEntry *LLDPMIB_LldpStatsRxPortTable_LldpStatsRxPortEntry) GetEntityData() *types.CommonEntityData {
    lldpStatsRxPortEntry.EntityData.YFilter = lldpStatsRxPortEntry.YFilter
    lldpStatsRxPortEntry.EntityData.YangName = "lldpStatsRxPortEntry"
    lldpStatsRxPortEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpStatsRxPortEntry.EntityData.ParentYangName = "lldpStatsRxPortTable"
    lldpStatsRxPortEntry.EntityData.SegmentPath = "lldpStatsRxPortEntry" + types.AddKeyToken(lldpStatsRxPortEntry.LldpStatsRxPortNum, "lldpStatsRxPortNum")
    lldpStatsRxPortEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpStatsRxPortTable/" + lldpStatsRxPortEntry.EntityData.SegmentPath
    lldpStatsRxPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpStatsRxPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpStatsRxPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpStatsRxPortEntry.EntityData.Children = types.NewOrderedMap()
    lldpStatsRxPortEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpStatsRxPortEntry.EntityData.Leafs.Append("lldpStatsRxPortNum", types.YLeaf{"LldpStatsRxPortNum", lldpStatsRxPortEntry.LldpStatsRxPortNum})
    lldpStatsRxPortEntry.EntityData.Leafs.Append("lldpStatsRxPortFramesDiscardedTotal", types.YLeaf{"LldpStatsRxPortFramesDiscardedTotal", lldpStatsRxPortEntry.LldpStatsRxPortFramesDiscardedTotal})
    lldpStatsRxPortEntry.EntityData.Leafs.Append("lldpStatsRxPortFramesErrors", types.YLeaf{"LldpStatsRxPortFramesErrors", lldpStatsRxPortEntry.LldpStatsRxPortFramesErrors})
    lldpStatsRxPortEntry.EntityData.Leafs.Append("lldpStatsRxPortFramesTotal", types.YLeaf{"LldpStatsRxPortFramesTotal", lldpStatsRxPortEntry.LldpStatsRxPortFramesTotal})
    lldpStatsRxPortEntry.EntityData.Leafs.Append("lldpStatsRxPortTLVsDiscardedTotal", types.YLeaf{"LldpStatsRxPortTLVsDiscardedTotal", lldpStatsRxPortEntry.LldpStatsRxPortTLVsDiscardedTotal})
    lldpStatsRxPortEntry.EntityData.Leafs.Append("lldpStatsRxPortTLVsUnrecognizedTotal", types.YLeaf{"LldpStatsRxPortTLVsUnrecognizedTotal", lldpStatsRxPortEntry.LldpStatsRxPortTLVsUnrecognizedTotal})
    lldpStatsRxPortEntry.EntityData.Leafs.Append("lldpStatsRxPortAgeoutsTotal", types.YLeaf{"LldpStatsRxPortAgeoutsTotal", lldpStatsRxPortEntry.LldpStatsRxPortAgeoutsTotal})

    lldpStatsRxPortEntry.EntityData.YListKeys = []string {"LldpStatsRxPortNum"}

    return &(lldpStatsRxPortEntry.EntityData)
}

// LLDPMIB_LldpLocPortTable
// This table contains one or more rows per port information
// associated with the local system known to this agent.
type LLDPMIB_LldpLocPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular port component.  Entries may be created and
    // deleted in this table by the agent. The type is slice of
    // LLDPMIB_LldpLocPortTable_LldpLocPortEntry.
    LldpLocPortEntry []*LLDPMIB_LldpLocPortTable_LldpLocPortEntry
}

func (lldpLocPortTable *LLDPMIB_LldpLocPortTable) GetEntityData() *types.CommonEntityData {
    lldpLocPortTable.EntityData.YFilter = lldpLocPortTable.YFilter
    lldpLocPortTable.EntityData.YangName = "lldpLocPortTable"
    lldpLocPortTable.EntityData.BundleName = "cisco_ios_xe"
    lldpLocPortTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpLocPortTable.EntityData.SegmentPath = "lldpLocPortTable"
    lldpLocPortTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpLocPortTable.EntityData.SegmentPath
    lldpLocPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpLocPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpLocPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpLocPortTable.EntityData.Children = types.NewOrderedMap()
    lldpLocPortTable.EntityData.Children.Append("lldpLocPortEntry", types.YChild{"LldpLocPortEntry", nil})
    for i := range lldpLocPortTable.LldpLocPortEntry {
        lldpLocPortTable.EntityData.Children.Append(types.GetSegmentPath(lldpLocPortTable.LldpLocPortEntry[i]), types.YChild{"LldpLocPortEntry", lldpLocPortTable.LldpLocPortEntry[i]})
    }
    lldpLocPortTable.EntityData.Leafs = types.NewOrderedMap()

    lldpLocPortTable.EntityData.YListKeys = []string {}

    return &(lldpLocPortTable.EntityData)
}

// LLDPMIB_LldpLocPortTable_LldpLocPortEntry
// Information about a particular port component.
// 
// Entries may be created and deleted in this table by the
// agent.
type LLDPMIB_LldpLocPortTable_LldpLocPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpLocPortTable. The type is interface{} with range: 1..4096.
    LldpLocPortNum interface{}

    // The type of port identifier encoding used in the associated 'lldpLocPortId'
    // object. The type is LldpPortIdSubtype.
    LldpLocPortIdSubtype interface{}

    // The string value used to identify the port component associated with a
    // given port in the local system. The type is string with length: 1..255.
    LldpLocPortId interface{}

    // The string value used to identify the 802 LAN station's port description
    // associated with the local system.  If the local agent supports IETF RFC
    // 2863, lldpLocPortDesc object should have the same value of ifDescr object.
    // The type is string with length: 0..255.
    LldpLocPortDesc interface{}
}

func (lldpLocPortEntry *LLDPMIB_LldpLocPortTable_LldpLocPortEntry) GetEntityData() *types.CommonEntityData {
    lldpLocPortEntry.EntityData.YFilter = lldpLocPortEntry.YFilter
    lldpLocPortEntry.EntityData.YangName = "lldpLocPortEntry"
    lldpLocPortEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpLocPortEntry.EntityData.ParentYangName = "lldpLocPortTable"
    lldpLocPortEntry.EntityData.SegmentPath = "lldpLocPortEntry" + types.AddKeyToken(lldpLocPortEntry.LldpLocPortNum, "lldpLocPortNum")
    lldpLocPortEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpLocPortTable/" + lldpLocPortEntry.EntityData.SegmentPath
    lldpLocPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpLocPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpLocPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpLocPortEntry.EntityData.Children = types.NewOrderedMap()
    lldpLocPortEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpLocPortEntry.EntityData.Leafs.Append("lldpLocPortNum", types.YLeaf{"LldpLocPortNum", lldpLocPortEntry.LldpLocPortNum})
    lldpLocPortEntry.EntityData.Leafs.Append("lldpLocPortIdSubtype", types.YLeaf{"LldpLocPortIdSubtype", lldpLocPortEntry.LldpLocPortIdSubtype})
    lldpLocPortEntry.EntityData.Leafs.Append("lldpLocPortId", types.YLeaf{"LldpLocPortId", lldpLocPortEntry.LldpLocPortId})
    lldpLocPortEntry.EntityData.Leafs.Append("lldpLocPortDesc", types.YLeaf{"LldpLocPortDesc", lldpLocPortEntry.LldpLocPortDesc})

    lldpLocPortEntry.EntityData.YListKeys = []string {"LldpLocPortNum"}

    return &(lldpLocPortEntry.EntityData)
}

// LLDPMIB_LldpLocManAddrTable
// This table contains management address information on the
// local system known to this agent.
type LLDPMIB_LldpLocManAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Management address information about a particular chassis component.  There
    // may be multiple management addresses configured on the system identified by
    // a particular lldpLocChassisId.  Each management address should have
    // distinct 'management address type' (lldpLocManAddrSubtype) and 'management
    // address' (lldpLocManAddr.)  Entries may be created and deleted in this
    // table by the agent. The type is slice of
    // LLDPMIB_LldpLocManAddrTable_LldpLocManAddrEntry.
    LldpLocManAddrEntry []*LLDPMIB_LldpLocManAddrTable_LldpLocManAddrEntry
}

func (lldpLocManAddrTable *LLDPMIB_LldpLocManAddrTable) GetEntityData() *types.CommonEntityData {
    lldpLocManAddrTable.EntityData.YFilter = lldpLocManAddrTable.YFilter
    lldpLocManAddrTable.EntityData.YangName = "lldpLocManAddrTable"
    lldpLocManAddrTable.EntityData.BundleName = "cisco_ios_xe"
    lldpLocManAddrTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpLocManAddrTable.EntityData.SegmentPath = "lldpLocManAddrTable"
    lldpLocManAddrTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpLocManAddrTable.EntityData.SegmentPath
    lldpLocManAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpLocManAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpLocManAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpLocManAddrTable.EntityData.Children = types.NewOrderedMap()
    lldpLocManAddrTable.EntityData.Children.Append("lldpLocManAddrEntry", types.YChild{"LldpLocManAddrEntry", nil})
    for i := range lldpLocManAddrTable.LldpLocManAddrEntry {
        lldpLocManAddrTable.EntityData.Children.Append(types.GetSegmentPath(lldpLocManAddrTable.LldpLocManAddrEntry[i]), types.YChild{"LldpLocManAddrEntry", lldpLocManAddrTable.LldpLocManAddrEntry[i]})
    }
    lldpLocManAddrTable.EntityData.Leafs = types.NewOrderedMap()

    lldpLocManAddrTable.EntityData.YListKeys = []string {}

    return &(lldpLocManAddrTable.EntityData)
}

// LLDPMIB_LldpLocManAddrTable_LldpLocManAddrEntry
// Management address information about a particular chassis
// component.  There may be multiple management addresses
// configured on the system identified by a particular
// lldpLocChassisId.  Each management address should have
// distinct 'management address type' (lldpLocManAddrSubtype) and
// 'management address' (lldpLocManAddr.)
// 
// Entries may be created and deleted in this table by the
// agent.
type LLDPMIB_LldpLocManAddrTable_LldpLocManAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type of management address identifier encoding
    // used in the associated 'lldpLocManagmentAddr' object. The type is
    // AddressFamilyNumbers.
    LldpLocManAddrSubtype interface{}

    // This attribute is a key. The string value used to identify the management
    // address component associated with the local system.  The purpose of this
    // address is to contact the management entity. The type is string with
    // length: 1..31.
    LldpLocManAddr interface{}

    // The total length of the management address subtype and the management
    // address fields in LLDPDUs transmitted by the local LLDP agent.  The
    // management address length field is needed so that the receiving systems
    // that do not implement SNMP will not be required to implement an iana family
    // numbers/address length equivalency table in order to decode the management
    // adress. The type is interface{} with range: -2147483648..2147483647.
    LldpLocManAddrLen interface{}

    // The enumeration value that identifies the interface numbering method used
    // for defining the interface number, associated with the local system. The
    // type is LldpManAddrIfSubtype.
    LldpLocManAddrIfSubtype interface{}

    // The integer value used to identify the interface number regarding the
    // management address component associated with the local system. The type is
    // interface{} with range: -2147483648..2147483647.
    LldpLocManAddrIfId interface{}

    // The OID value used to identify the type of hardware component or protocol
    // entity associated with the management address advertised by the local
    // system agent. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    LldpLocManAddrOID interface{}

    // A set of ports that are identified by a PortList, in which each port is
    // represented as a bit.  The corresponding local system management address
    // instance will be transmitted on the member ports of the
    // lldpManAddrPortsTxEnable.    The default value for
    // lldpConfigManAddrPortsTxEnable object is empty binary string, which means
    // no ports are specified for advertising indicated management address
    // instance. The type is string with length: 0..512.
    LldpConfigManAddrPortsTxEnable interface{}
}

func (lldpLocManAddrEntry *LLDPMIB_LldpLocManAddrTable_LldpLocManAddrEntry) GetEntityData() *types.CommonEntityData {
    lldpLocManAddrEntry.EntityData.YFilter = lldpLocManAddrEntry.YFilter
    lldpLocManAddrEntry.EntityData.YangName = "lldpLocManAddrEntry"
    lldpLocManAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpLocManAddrEntry.EntityData.ParentYangName = "lldpLocManAddrTable"
    lldpLocManAddrEntry.EntityData.SegmentPath = "lldpLocManAddrEntry" + types.AddKeyToken(lldpLocManAddrEntry.LldpLocManAddrSubtype, "lldpLocManAddrSubtype") + types.AddKeyToken(lldpLocManAddrEntry.LldpLocManAddr, "lldpLocManAddr")
    lldpLocManAddrEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpLocManAddrTable/" + lldpLocManAddrEntry.EntityData.SegmentPath
    lldpLocManAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpLocManAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpLocManAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpLocManAddrEntry.EntityData.Children = types.NewOrderedMap()
    lldpLocManAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpLocManAddrEntry.EntityData.Leafs.Append("lldpLocManAddrSubtype", types.YLeaf{"LldpLocManAddrSubtype", lldpLocManAddrEntry.LldpLocManAddrSubtype})
    lldpLocManAddrEntry.EntityData.Leafs.Append("lldpLocManAddr", types.YLeaf{"LldpLocManAddr", lldpLocManAddrEntry.LldpLocManAddr})
    lldpLocManAddrEntry.EntityData.Leafs.Append("lldpLocManAddrLen", types.YLeaf{"LldpLocManAddrLen", lldpLocManAddrEntry.LldpLocManAddrLen})
    lldpLocManAddrEntry.EntityData.Leafs.Append("lldpLocManAddrIfSubtype", types.YLeaf{"LldpLocManAddrIfSubtype", lldpLocManAddrEntry.LldpLocManAddrIfSubtype})
    lldpLocManAddrEntry.EntityData.Leafs.Append("lldpLocManAddrIfId", types.YLeaf{"LldpLocManAddrIfId", lldpLocManAddrEntry.LldpLocManAddrIfId})
    lldpLocManAddrEntry.EntityData.Leafs.Append("lldpLocManAddrOID", types.YLeaf{"LldpLocManAddrOID", lldpLocManAddrEntry.LldpLocManAddrOID})
    lldpLocManAddrEntry.EntityData.Leafs.Append("lldpConfigManAddrPortsTxEnable", types.YLeaf{"LldpConfigManAddrPortsTxEnable", lldpLocManAddrEntry.LldpConfigManAddrPortsTxEnable})

    lldpLocManAddrEntry.EntityData.YListKeys = []string {"LldpLocManAddrSubtype", "LldpLocManAddr"}

    return &(lldpLocManAddrEntry.EntityData)
}

// LLDPMIB_LldpRemTable
// This table contains one or more rows per physical network
// connection known to this agent.  The agent may wish to ensure
// that only one lldpRemEntry is present for each local port,
// or it may choose to maintain multiple lldpRemEntries for
// the same local port.
// 
// The following procedure may be used to retrieve remote
// systems information updates from an LLDP agent:
// 
//    1. NMS polls all tables associated with remote systems
//       and keeps a local copy of the information retrieved.
//       NMS polls periodically the values of the following
//       objects:
//          a. lldpStatsRemTablesInserts
//          b. lldpStatsRemTablesDeletes
//          c. lldpStatsRemTablesDrops
//          d. lldpStatsRemTablesAgeouts
//          e. lldpStatsRxPortAgeoutsTotal for all ports.
// 
//    2. LLDP agent updates remote systems MIB objects, and
//       sends out notifications to a list of notification
//       destinations.
// 
//    3. NMS receives the notifications and compares the new
//       values of objects listed in step 1.  
// 
//       Periodically, NMS should poll the object
//       lldpStatsRemTablesLastChangeTime to find out if anything
//       has changed since the last poll.  if something has
//       changed, NMS will poll the objects listed in step 1 to
//       figure out what kind of changes occurred in the tables.
// 
//       if value of lldpStatsRemTablesInserts has changed,
//       then NMS will walk all tables by employing TimeFilter
//       with the last-polled time value.  This request will
//       return new objects or objects whose values are updated
//       since the last poll.
// 
//       if value of lldpStatsRemTablesAgeouts has changed,
//       then NMS will walk the lldpStatsRxPortAgeoutsTotal and
//       compare the new values with previously recorded ones.
//       For ports whose lldpStatsRxPortAgeoutsTotal value is
//       greater than the recorded value, NMS will have to
//       retrieve objects associated with those ports from
//       table(s) without employing a TimeFilter (which is
//       performed by specifying 0 for the TimeFilter.)
// 
//       lldpStatsRemTablesDeletes and lldpStatsRemTablesDrops
//       objects are provided for informational purposes.
type LLDPMIB_LldpRemTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular physical network connection. Entries may be
    // created and deleted in this table by the agent, if a physical topology
    // discovery process is active. The type is slice of
    // LLDPMIB_LldpRemTable_LldpRemEntry.
    LldpRemEntry []*LLDPMIB_LldpRemTable_LldpRemEntry
}

func (lldpRemTable *LLDPMIB_LldpRemTable) GetEntityData() *types.CommonEntityData {
    lldpRemTable.EntityData.YFilter = lldpRemTable.YFilter
    lldpRemTable.EntityData.YangName = "lldpRemTable"
    lldpRemTable.EntityData.BundleName = "cisco_ios_xe"
    lldpRemTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpRemTable.EntityData.SegmentPath = "lldpRemTable"
    lldpRemTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpRemTable.EntityData.SegmentPath
    lldpRemTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemTable.EntityData.Children = types.NewOrderedMap()
    lldpRemTable.EntityData.Children.Append("lldpRemEntry", types.YChild{"LldpRemEntry", nil})
    for i := range lldpRemTable.LldpRemEntry {
        lldpRemTable.EntityData.Children.Append(types.GetSegmentPath(lldpRemTable.LldpRemEntry[i]), types.YChild{"LldpRemEntry", lldpRemTable.LldpRemEntry[i]})
    }
    lldpRemTable.EntityData.Leafs = types.NewOrderedMap()

    lldpRemTable.EntityData.YListKeys = []string {}

    return &(lldpRemTable.EntityData)
}

// LLDPMIB_LldpRemTable_LldpRemEntry
// Information about a particular physical network connection.
// Entries may be created and deleted in this table by the agent,
// if a physical topology discovery process is active.
type LLDPMIB_LldpRemTable_LldpRemEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention in IETF RFC 2021 and 
    // http://www.ietf.org/IESG/Implementations/RFC2021-Implementation.txt to see
    // how TimeFilter works. The type is interface{} with range: 0..4294967295.
    LldpRemTimeMark interface{}

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The lldpRemLocalPortNum identifies the port on which the
    // remote system information is received.  The value of this object is used as
    // a port index to the lldpRemTable. The type is interface{} with range:
    // 1..4096.
    LldpRemLocalPortNum interface{}

    // This attribute is a key. This object represents an arbitrary local integer
    // value used by this agent to identify a particular connection instance,
    // unique only for the indicated remote system.  An agent is encouraged to
    // assign monotonically increasing index values to new entries, starting with
    // one, after each reboot.  It is considered unlikely that the lldpRemIndex
    // will wrap between reboots. The type is interface{} with range:
    // 1..2147483647.
    LldpRemIndex interface{}

    // The type of encoding used to identify the chassis associated with the
    // remote system. The type is LldpChassisIdSubtype.
    LldpRemChassisIdSubtype interface{}

    // The string value used to identify the chassis component associated with the
    // remote system. The type is string with length: 1..255.
    LldpRemChassisId interface{}

    // The type of port identifier encoding used in the associated 'lldpRemPortId'
    // object. The type is LldpPortIdSubtype.
    LldpRemPortIdSubtype interface{}

    // The string value used to identify the port component associated with the
    // remote system. The type is string with length: 1..255.
    LldpRemPortId interface{}

    // The string value used to identify the description of the given port
    // associated with the remote system. The type is string with length: 0..255.
    LldpRemPortDesc interface{}

    // The string value used to identify the system name of the remote system. The
    // type is string with length: 0..255.
    LldpRemSysName interface{}

    // The string value used to identify the system description of the remote
    // system. The type is string with length: 0..255.
    LldpRemSysDesc interface{}

    // The bitmap value used to identify which system capabilities are supported
    // on the remote system. The type is map[string]bool.
    LldpRemSysCapSupported interface{}

    // The bitmap value used to identify which system capabilities are enabled on
    // the remote system. The type is map[string]bool.
    LldpRemSysCapEnabled interface{}
}

func (lldpRemEntry *LLDPMIB_LldpRemTable_LldpRemEntry) GetEntityData() *types.CommonEntityData {
    lldpRemEntry.EntityData.YFilter = lldpRemEntry.YFilter
    lldpRemEntry.EntityData.YangName = "lldpRemEntry"
    lldpRemEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpRemEntry.EntityData.ParentYangName = "lldpRemTable"
    lldpRemEntry.EntityData.SegmentPath = "lldpRemEntry" + types.AddKeyToken(lldpRemEntry.LldpRemTimeMark, "lldpRemTimeMark") + types.AddKeyToken(lldpRemEntry.LldpRemLocalPortNum, "lldpRemLocalPortNum") + types.AddKeyToken(lldpRemEntry.LldpRemIndex, "lldpRemIndex")
    lldpRemEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpRemTable/" + lldpRemEntry.EntityData.SegmentPath
    lldpRemEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemEntry.EntityData.Children = types.NewOrderedMap()
    lldpRemEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpRemEntry.EntityData.Leafs.Append("lldpRemTimeMark", types.YLeaf{"LldpRemTimeMark", lldpRemEntry.LldpRemTimeMark})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemLocalPortNum", types.YLeaf{"LldpRemLocalPortNum", lldpRemEntry.LldpRemLocalPortNum})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemIndex", types.YLeaf{"LldpRemIndex", lldpRemEntry.LldpRemIndex})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemChassisIdSubtype", types.YLeaf{"LldpRemChassisIdSubtype", lldpRemEntry.LldpRemChassisIdSubtype})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemChassisId", types.YLeaf{"LldpRemChassisId", lldpRemEntry.LldpRemChassisId})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemPortIdSubtype", types.YLeaf{"LldpRemPortIdSubtype", lldpRemEntry.LldpRemPortIdSubtype})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemPortId", types.YLeaf{"LldpRemPortId", lldpRemEntry.LldpRemPortId})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemPortDesc", types.YLeaf{"LldpRemPortDesc", lldpRemEntry.LldpRemPortDesc})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemSysName", types.YLeaf{"LldpRemSysName", lldpRemEntry.LldpRemSysName})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemSysDesc", types.YLeaf{"LldpRemSysDesc", lldpRemEntry.LldpRemSysDesc})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemSysCapSupported", types.YLeaf{"LldpRemSysCapSupported", lldpRemEntry.LldpRemSysCapSupported})
    lldpRemEntry.EntityData.Leafs.Append("lldpRemSysCapEnabled", types.YLeaf{"LldpRemSysCapEnabled", lldpRemEntry.LldpRemSysCapEnabled})

    lldpRemEntry.EntityData.YListKeys = []string {"LldpRemTimeMark", "LldpRemLocalPortNum", "LldpRemIndex"}

    return &(lldpRemEntry.EntityData)
}

// LLDPMIB_LldpRemManAddrTable
// This table contains one or more rows per management address
// information on the remote system learned on a particular port
// contained in the local chassis known to this agent.
type LLDPMIB_LldpRemManAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Management address information about a particular chassis component.  There
    // may be multiple management addresses configured on the remote system
    // identified by a particular lldpRemIndex whose information is received on
    // lldpRemLocalPortNum of the local system.  Each management address should
    // have distinct 'management address type' (lldpRemManAddrSubtype) and
    // 'management address' (lldpRemManAddr.)  Entries may be created and deleted
    // in this table by the agent. The type is slice of
    // LLDPMIB_LldpRemManAddrTable_LldpRemManAddrEntry.
    LldpRemManAddrEntry []*LLDPMIB_LldpRemManAddrTable_LldpRemManAddrEntry
}

func (lldpRemManAddrTable *LLDPMIB_LldpRemManAddrTable) GetEntityData() *types.CommonEntityData {
    lldpRemManAddrTable.EntityData.YFilter = lldpRemManAddrTable.YFilter
    lldpRemManAddrTable.EntityData.YangName = "lldpRemManAddrTable"
    lldpRemManAddrTable.EntityData.BundleName = "cisco_ios_xe"
    lldpRemManAddrTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpRemManAddrTable.EntityData.SegmentPath = "lldpRemManAddrTable"
    lldpRemManAddrTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpRemManAddrTable.EntityData.SegmentPath
    lldpRemManAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemManAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemManAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemManAddrTable.EntityData.Children = types.NewOrderedMap()
    lldpRemManAddrTable.EntityData.Children.Append("lldpRemManAddrEntry", types.YChild{"LldpRemManAddrEntry", nil})
    for i := range lldpRemManAddrTable.LldpRemManAddrEntry {
        lldpRemManAddrTable.EntityData.Children.Append(types.GetSegmentPath(lldpRemManAddrTable.LldpRemManAddrEntry[i]), types.YChild{"LldpRemManAddrEntry", lldpRemManAddrTable.LldpRemManAddrEntry[i]})
    }
    lldpRemManAddrTable.EntityData.Leafs = types.NewOrderedMap()

    lldpRemManAddrTable.EntityData.YListKeys = []string {}

    return &(lldpRemManAddrTable.EntityData)
}

// LLDPMIB_LldpRemManAddrTable_LldpRemManAddrEntry
// Management address information about a particular chassis
// component.  There may be multiple management addresses
// configured on the remote system identified by a particular
// lldpRemIndex whose information is received on
// lldpRemLocalPortNum of the local system.  Each management
// address should have distinct 'management address
// type' (lldpRemManAddrSubtype) and 'management address'
// (lldpRemManAddr.)
// 
// Entries may be created and deleted in this table by the
// agent.
type LLDPMIB_LldpRemManAddrTable_LldpRemManAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemTimeMark
    LldpRemTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..4096. Refers to
    // lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemLocalPortNum
    LldpRemLocalPortNum interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemIndex
    LldpRemIndex interface{}

    // This attribute is a key. The type of management address identifier encoding
    // used in the associated 'lldpRemManagmentAddr' object. The type is
    // AddressFamilyNumbers.
    LldpRemManAddrSubtype interface{}

    // This attribute is a key. The string value used to identify the management
    // address component associated with the remote system.  The purpose of this
    // address is to contact the management entity. The type is string with
    // length: 1..31.
    LldpRemManAddr interface{}

    // The enumeration value that identifies the interface numbering method used
    // for defining the interface number, associated with the remote system. The
    // type is LldpManAddrIfSubtype.
    LldpRemManAddrIfSubtype interface{}

    // The integer value used to identify the interface number regarding the
    // management address component associated with the remote system. The type is
    // interface{} with range: -2147483648..2147483647.
    LldpRemManAddrIfId interface{}

    // The OID value used to identify the type of hardware component or protocol
    // entity associated with the management address advertised by the remote
    // system agent. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    LldpRemManAddrOID interface{}
}

func (lldpRemManAddrEntry *LLDPMIB_LldpRemManAddrTable_LldpRemManAddrEntry) GetEntityData() *types.CommonEntityData {
    lldpRemManAddrEntry.EntityData.YFilter = lldpRemManAddrEntry.YFilter
    lldpRemManAddrEntry.EntityData.YangName = "lldpRemManAddrEntry"
    lldpRemManAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpRemManAddrEntry.EntityData.ParentYangName = "lldpRemManAddrTable"
    lldpRemManAddrEntry.EntityData.SegmentPath = "lldpRemManAddrEntry" + types.AddKeyToken(lldpRemManAddrEntry.LldpRemTimeMark, "lldpRemTimeMark") + types.AddKeyToken(lldpRemManAddrEntry.LldpRemLocalPortNum, "lldpRemLocalPortNum") + types.AddKeyToken(lldpRemManAddrEntry.LldpRemIndex, "lldpRemIndex") + types.AddKeyToken(lldpRemManAddrEntry.LldpRemManAddrSubtype, "lldpRemManAddrSubtype") + types.AddKeyToken(lldpRemManAddrEntry.LldpRemManAddr, "lldpRemManAddr")
    lldpRemManAddrEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpRemManAddrTable/" + lldpRemManAddrEntry.EntityData.SegmentPath
    lldpRemManAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemManAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemManAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemManAddrEntry.EntityData.Children = types.NewOrderedMap()
    lldpRemManAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemTimeMark", types.YLeaf{"LldpRemTimeMark", lldpRemManAddrEntry.LldpRemTimeMark})
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemLocalPortNum", types.YLeaf{"LldpRemLocalPortNum", lldpRemManAddrEntry.LldpRemLocalPortNum})
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemIndex", types.YLeaf{"LldpRemIndex", lldpRemManAddrEntry.LldpRemIndex})
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemManAddrSubtype", types.YLeaf{"LldpRemManAddrSubtype", lldpRemManAddrEntry.LldpRemManAddrSubtype})
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemManAddr", types.YLeaf{"LldpRemManAddr", lldpRemManAddrEntry.LldpRemManAddr})
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemManAddrIfSubtype", types.YLeaf{"LldpRemManAddrIfSubtype", lldpRemManAddrEntry.LldpRemManAddrIfSubtype})
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemManAddrIfId", types.YLeaf{"LldpRemManAddrIfId", lldpRemManAddrEntry.LldpRemManAddrIfId})
    lldpRemManAddrEntry.EntityData.Leafs.Append("lldpRemManAddrOID", types.YLeaf{"LldpRemManAddrOID", lldpRemManAddrEntry.LldpRemManAddrOID})

    lldpRemManAddrEntry.EntityData.YListKeys = []string {"LldpRemTimeMark", "LldpRemLocalPortNum", "LldpRemIndex", "LldpRemManAddrSubtype", "LldpRemManAddr"}

    return &(lldpRemManAddrEntry.EntityData)
}

// LLDPMIB_LldpRemUnknownTLVTable
// This table contains information about an incoming TLV which
// is not recognized by the receiving LLDP agent.  The TLV may
// be from a later version of the basic management set.
// 
// This table should only contain TLVs that are found in
// a single LLDP frame.  Entries in this table, associated
// with an MAC service access point (MSAP, the access point
// for MAC services provided to the LCC sublayer, defined
// in IEEE 100, which is also identified with a particular
// lldpRemLocalPortNum, lldpRemIndex pair) are overwritten with
// most recently received unrecognized TLV from the same MSAP,
// or they will naturally age out when the rxInfoTTL timer
// (associated with the MSAP) expires.
type LLDPMIB_LldpRemUnknownTLVTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an unrecognized TLV received from a physical network
    // connection.  Entries may be created and deleted in this table by the agent,
    // if a physical topology discovery process is active. The type is slice of
    // LLDPMIB_LldpRemUnknownTLVTable_LldpRemUnknownTLVEntry.
    LldpRemUnknownTLVEntry []*LLDPMIB_LldpRemUnknownTLVTable_LldpRemUnknownTLVEntry
}

func (lldpRemUnknownTLVTable *LLDPMIB_LldpRemUnknownTLVTable) GetEntityData() *types.CommonEntityData {
    lldpRemUnknownTLVTable.EntityData.YFilter = lldpRemUnknownTLVTable.YFilter
    lldpRemUnknownTLVTable.EntityData.YangName = "lldpRemUnknownTLVTable"
    lldpRemUnknownTLVTable.EntityData.BundleName = "cisco_ios_xe"
    lldpRemUnknownTLVTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpRemUnknownTLVTable.EntityData.SegmentPath = "lldpRemUnknownTLVTable"
    lldpRemUnknownTLVTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpRemUnknownTLVTable.EntityData.SegmentPath
    lldpRemUnknownTLVTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemUnknownTLVTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemUnknownTLVTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemUnknownTLVTable.EntityData.Children = types.NewOrderedMap()
    lldpRemUnknownTLVTable.EntityData.Children.Append("lldpRemUnknownTLVEntry", types.YChild{"LldpRemUnknownTLVEntry", nil})
    for i := range lldpRemUnknownTLVTable.LldpRemUnknownTLVEntry {
        lldpRemUnknownTLVTable.EntityData.Children.Append(types.GetSegmentPath(lldpRemUnknownTLVTable.LldpRemUnknownTLVEntry[i]), types.YChild{"LldpRemUnknownTLVEntry", lldpRemUnknownTLVTable.LldpRemUnknownTLVEntry[i]})
    }
    lldpRemUnknownTLVTable.EntityData.Leafs = types.NewOrderedMap()

    lldpRemUnknownTLVTable.EntityData.YListKeys = []string {}

    return &(lldpRemUnknownTLVTable.EntityData)
}

// LLDPMIB_LldpRemUnknownTLVTable_LldpRemUnknownTLVEntry
// Information about an unrecognized TLV received from a
// physical network connection.  Entries may be created and
// deleted in this table by the agent, if a physical topology
// discovery process is active.
type LLDPMIB_LldpRemUnknownTLVTable_LldpRemUnknownTLVEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemTimeMark
    LldpRemTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..4096. Refers to
    // lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemLocalPortNum
    LldpRemLocalPortNum interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemIndex
    LldpRemIndex interface{}

    // This attribute is a key. This object represents the value extracted from
    // the type field of the TLV. The type is interface{} with range: 9..126.
    LldpRemUnknownTLVType interface{}

    // This object represents the value extracted from the value field of the TLV.
    // The type is string with length: 0..511.
    LldpRemUnknownTLVInfo interface{}
}

func (lldpRemUnknownTLVEntry *LLDPMIB_LldpRemUnknownTLVTable_LldpRemUnknownTLVEntry) GetEntityData() *types.CommonEntityData {
    lldpRemUnknownTLVEntry.EntityData.YFilter = lldpRemUnknownTLVEntry.YFilter
    lldpRemUnknownTLVEntry.EntityData.YangName = "lldpRemUnknownTLVEntry"
    lldpRemUnknownTLVEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpRemUnknownTLVEntry.EntityData.ParentYangName = "lldpRemUnknownTLVTable"
    lldpRemUnknownTLVEntry.EntityData.SegmentPath = "lldpRemUnknownTLVEntry" + types.AddKeyToken(lldpRemUnknownTLVEntry.LldpRemTimeMark, "lldpRemTimeMark") + types.AddKeyToken(lldpRemUnknownTLVEntry.LldpRemLocalPortNum, "lldpRemLocalPortNum") + types.AddKeyToken(lldpRemUnknownTLVEntry.LldpRemIndex, "lldpRemIndex") + types.AddKeyToken(lldpRemUnknownTLVEntry.LldpRemUnknownTLVType, "lldpRemUnknownTLVType")
    lldpRemUnknownTLVEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpRemUnknownTLVTable/" + lldpRemUnknownTLVEntry.EntityData.SegmentPath
    lldpRemUnknownTLVEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemUnknownTLVEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemUnknownTLVEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemUnknownTLVEntry.EntityData.Children = types.NewOrderedMap()
    lldpRemUnknownTLVEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpRemUnknownTLVEntry.EntityData.Leafs.Append("lldpRemTimeMark", types.YLeaf{"LldpRemTimeMark", lldpRemUnknownTLVEntry.LldpRemTimeMark})
    lldpRemUnknownTLVEntry.EntityData.Leafs.Append("lldpRemLocalPortNum", types.YLeaf{"LldpRemLocalPortNum", lldpRemUnknownTLVEntry.LldpRemLocalPortNum})
    lldpRemUnknownTLVEntry.EntityData.Leafs.Append("lldpRemIndex", types.YLeaf{"LldpRemIndex", lldpRemUnknownTLVEntry.LldpRemIndex})
    lldpRemUnknownTLVEntry.EntityData.Leafs.Append("lldpRemUnknownTLVType", types.YLeaf{"LldpRemUnknownTLVType", lldpRemUnknownTLVEntry.LldpRemUnknownTLVType})
    lldpRemUnknownTLVEntry.EntityData.Leafs.Append("lldpRemUnknownTLVInfo", types.YLeaf{"LldpRemUnknownTLVInfo", lldpRemUnknownTLVEntry.LldpRemUnknownTLVInfo})

    lldpRemUnknownTLVEntry.EntityData.YListKeys = []string {"LldpRemTimeMark", "LldpRemLocalPortNum", "LldpRemIndex", "LldpRemUnknownTLVType"}

    return &(lldpRemUnknownTLVEntry.EntityData)
}

// LLDPMIB_LldpRemOrgDefInfoTable
// This table contains one or more rows per physical network
// connection which advertises the organizationally defined
// information.
// 
// Note that this table contains one or more rows of
// organizationally defined information that is not recognized
// by the local agent.
// 
// If the local system is capable of recognizing any
// organizationally defined information, appropriate extension
// MIBs from the organization should be used for information
// retrieval.
type LLDPMIB_LldpRemOrgDefInfoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the unrecognized organizationally defined information
    // advertised by the remote system. The lldpRemTimeMark, lldpRemLocalPortNum,
    // lldpRemIndex, lldpRemOrgDefInfoOUI, lldpRemOrgDefInfoSubtype, and
    // lldpRemOrgDefInfoIndex are indexes to this table.  If there is an
    // lldpRemOrgDefInfoEntry associated with a particular remote system
    // identified by the lldpRemLocalPortNum and lldpRemIndex, there must be an
    // lldpRemEntry associated with the same instance (i.e, using same indexes.) 
    // When the lldpRemEntry for the same index is removed from the lldpRemTable,
    // the associated lldpRemOrgDefInfoEntry should be removed from the
    // lldpRemOrgDefInfoTable.  Entries may be created and deleted in this table
    // by the agent. The type is slice of
    // LLDPMIB_LldpRemOrgDefInfoTable_LldpRemOrgDefInfoEntry.
    LldpRemOrgDefInfoEntry []*LLDPMIB_LldpRemOrgDefInfoTable_LldpRemOrgDefInfoEntry
}

func (lldpRemOrgDefInfoTable *LLDPMIB_LldpRemOrgDefInfoTable) GetEntityData() *types.CommonEntityData {
    lldpRemOrgDefInfoTable.EntityData.YFilter = lldpRemOrgDefInfoTable.YFilter
    lldpRemOrgDefInfoTable.EntityData.YangName = "lldpRemOrgDefInfoTable"
    lldpRemOrgDefInfoTable.EntityData.BundleName = "cisco_ios_xe"
    lldpRemOrgDefInfoTable.EntityData.ParentYangName = "LLDP-MIB"
    lldpRemOrgDefInfoTable.EntityData.SegmentPath = "lldpRemOrgDefInfoTable"
    lldpRemOrgDefInfoTable.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/" + lldpRemOrgDefInfoTable.EntityData.SegmentPath
    lldpRemOrgDefInfoTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemOrgDefInfoTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemOrgDefInfoTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemOrgDefInfoTable.EntityData.Children = types.NewOrderedMap()
    lldpRemOrgDefInfoTable.EntityData.Children.Append("lldpRemOrgDefInfoEntry", types.YChild{"LldpRemOrgDefInfoEntry", nil})
    for i := range lldpRemOrgDefInfoTable.LldpRemOrgDefInfoEntry {
        lldpRemOrgDefInfoTable.EntityData.Children.Append(types.GetSegmentPath(lldpRemOrgDefInfoTable.LldpRemOrgDefInfoEntry[i]), types.YChild{"LldpRemOrgDefInfoEntry", lldpRemOrgDefInfoTable.LldpRemOrgDefInfoEntry[i]})
    }
    lldpRemOrgDefInfoTable.EntityData.Leafs = types.NewOrderedMap()

    lldpRemOrgDefInfoTable.EntityData.YListKeys = []string {}

    return &(lldpRemOrgDefInfoTable.EntityData)
}

// LLDPMIB_LldpRemOrgDefInfoTable_LldpRemOrgDefInfoEntry
// Information about the unrecognized organizationally
// defined information advertised by the remote system.
// The lldpRemTimeMark, lldpRemLocalPortNum, lldpRemIndex,
// lldpRemOrgDefInfoOUI, lldpRemOrgDefInfoSubtype, and
// lldpRemOrgDefInfoIndex are indexes to this table.  If there is
// an lldpRemOrgDefInfoEntry associated with a particular remote
// system identified by the lldpRemLocalPortNum and lldpRemIndex,
// there must be an lldpRemEntry associated with the same
// instance (i.e, using same indexes.)  When the lldpRemEntry
// for the same index is removed from the lldpRemTable, the
// associated lldpRemOrgDefInfoEntry should be removed from
// the lldpRemOrgDefInfoTable.
// 
// Entries may be created and deleted in this table by the
// agent.
type LLDPMIB_LldpRemOrgDefInfoTable_LldpRemOrgDefInfoEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemTimeMark
    LldpRemTimeMark interface{}

    // This attribute is a key. The type is string with range: 1..4096. Refers to
    // lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemLocalPortNum
    LldpRemLocalPortNum interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to lldp_mib.LLDPMIB_LldpRemTable_LldpRemEntry_LldpRemIndex
    LldpRemIndex interface{}

    // This attribute is a key. The Organizationally Unique Identifier (OUI), as
    // defined in IEEE std 802-2001, is a 24 bit (three octets) globally unique
    // assigned number referenced by various standards, of the information
    // received from the remote system. The type is string with length: 3..3.
    LldpRemOrgDefInfoOUI interface{}

    // This attribute is a key. The integer value used to identify the subtype of
    // the organizationally defined information received from the remote system. 
    // The subtype value is required to identify different instances of
    // organizationally defined information that could not be retrieved without a
    // unique identifier that indicates the particular type of information
    // contained in the information string. The type is interface{} with range:
    // 1..255.
    LldpRemOrgDefInfoSubtype interface{}

    // This attribute is a key. This object represents an arbitrary local integer
    // value used by this agent to identify a particular unrecognized
    // organizationally defined information instance, unique only for the
    // lldpRemOrgDefInfoOUI and lldpRemOrgDefInfoSubtype from the same remote
    // system.  An agent is encouraged to assign monotonically increasing index
    // values to new entries, starting with one, after each reboot.  It is
    // considered unlikely that the lldpRemOrgDefInfoIndex will wrap between
    // reboots. The type is interface{} with range: 1..2147483647.
    LldpRemOrgDefInfoIndex interface{}

    // The string value used to identify the organizationally defined information
    // of the remote system.  The encoding for this object should be as defined
    // for SnmpAdminString TC. The type is string with length: 0..507.
    LldpRemOrgDefInfo interface{}
}

func (lldpRemOrgDefInfoEntry *LLDPMIB_LldpRemOrgDefInfoTable_LldpRemOrgDefInfoEntry) GetEntityData() *types.CommonEntityData {
    lldpRemOrgDefInfoEntry.EntityData.YFilter = lldpRemOrgDefInfoEntry.YFilter
    lldpRemOrgDefInfoEntry.EntityData.YangName = "lldpRemOrgDefInfoEntry"
    lldpRemOrgDefInfoEntry.EntityData.BundleName = "cisco_ios_xe"
    lldpRemOrgDefInfoEntry.EntityData.ParentYangName = "lldpRemOrgDefInfoTable"
    lldpRemOrgDefInfoEntry.EntityData.SegmentPath = "lldpRemOrgDefInfoEntry" + types.AddKeyToken(lldpRemOrgDefInfoEntry.LldpRemTimeMark, "lldpRemTimeMark") + types.AddKeyToken(lldpRemOrgDefInfoEntry.LldpRemLocalPortNum, "lldpRemLocalPortNum") + types.AddKeyToken(lldpRemOrgDefInfoEntry.LldpRemIndex, "lldpRemIndex") + types.AddKeyToken(lldpRemOrgDefInfoEntry.LldpRemOrgDefInfoOUI, "lldpRemOrgDefInfoOUI") + types.AddKeyToken(lldpRemOrgDefInfoEntry.LldpRemOrgDefInfoSubtype, "lldpRemOrgDefInfoSubtype") + types.AddKeyToken(lldpRemOrgDefInfoEntry.LldpRemOrgDefInfoIndex, "lldpRemOrgDefInfoIndex")
    lldpRemOrgDefInfoEntry.EntityData.AbsolutePath = "LLDP-MIB:LLDP-MIB/lldpRemOrgDefInfoTable/" + lldpRemOrgDefInfoEntry.EntityData.SegmentPath
    lldpRemOrgDefInfoEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpRemOrgDefInfoEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpRemOrgDefInfoEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpRemOrgDefInfoEntry.EntityData.Children = types.NewOrderedMap()
    lldpRemOrgDefInfoEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpRemOrgDefInfoEntry.EntityData.Leafs.Append("lldpRemTimeMark", types.YLeaf{"LldpRemTimeMark", lldpRemOrgDefInfoEntry.LldpRemTimeMark})
    lldpRemOrgDefInfoEntry.EntityData.Leafs.Append("lldpRemLocalPortNum", types.YLeaf{"LldpRemLocalPortNum", lldpRemOrgDefInfoEntry.LldpRemLocalPortNum})
    lldpRemOrgDefInfoEntry.EntityData.Leafs.Append("lldpRemIndex", types.YLeaf{"LldpRemIndex", lldpRemOrgDefInfoEntry.LldpRemIndex})
    lldpRemOrgDefInfoEntry.EntityData.Leafs.Append("lldpRemOrgDefInfoOUI", types.YLeaf{"LldpRemOrgDefInfoOUI", lldpRemOrgDefInfoEntry.LldpRemOrgDefInfoOUI})
    lldpRemOrgDefInfoEntry.EntityData.Leafs.Append("lldpRemOrgDefInfoSubtype", types.YLeaf{"LldpRemOrgDefInfoSubtype", lldpRemOrgDefInfoEntry.LldpRemOrgDefInfoSubtype})
    lldpRemOrgDefInfoEntry.EntityData.Leafs.Append("lldpRemOrgDefInfoIndex", types.YLeaf{"LldpRemOrgDefInfoIndex", lldpRemOrgDefInfoEntry.LldpRemOrgDefInfoIndex})
    lldpRemOrgDefInfoEntry.EntityData.Leafs.Append("lldpRemOrgDefInfo", types.YLeaf{"LldpRemOrgDefInfo", lldpRemOrgDefInfoEntry.LldpRemOrgDefInfo})

    lldpRemOrgDefInfoEntry.EntityData.YListKeys = []string {"LldpRemTimeMark", "LldpRemLocalPortNum", "LldpRemIndex", "LldpRemOrgDefInfoOUI", "LldpRemOrgDefInfoSubtype", "LldpRemOrgDefInfoIndex"}

    return &(lldpRemOrgDefInfoEntry.EntityData)
}

