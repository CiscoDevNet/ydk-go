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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Lldpconfiguration LLDPMIB_Lldpconfiguration

    
    Lldpstatistics LLDPMIB_Lldpstatistics

    
    Lldplocalsystemdata LLDPMIB_Lldplocalsystemdata

    // The table that controls LLDP frame transmission on individual ports.
    Lldpportconfigtable LLDPMIB_Lldpportconfigtable

    // A table containing LLDP transmission statistics for individual ports. 
    // Entries are not required to exist in this table while the
    // lldpPortConfigEntry object is equal to 'disabled(4)'.
    Lldpstatstxporttable LLDPMIB_Lldpstatstxporttable

    // A table containing LLDP reception statistics for individual ports.  Entries
    // are not required to exist in this table while the lldpPortConfigEntry
    // object is equal to 'disabled(4)'.
    Lldpstatsrxporttable LLDPMIB_Lldpstatsrxporttable

    // This table contains one or more rows per port information associated with
    // the local system known to this agent.
    Lldplocporttable LLDPMIB_Lldplocporttable

    // This table contains management address information on the local system
    // known to this agent.
    Lldplocmanaddrtable LLDPMIB_Lldplocmanaddrtable

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
    Lldpremtable LLDPMIB_Lldpremtable

    // This table contains one or more rows per management address information on
    // the remote system learned on a particular port contained in the local
    // chassis known to this agent.
    Lldpremmanaddrtable LLDPMIB_Lldpremmanaddrtable

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
    Lldpremunknowntlvtable LLDPMIB_Lldpremunknowntlvtable

    // This table contains one or more rows per physical network connection which
    // advertises the organizationally defined information.  Note that this table
    // contains one or more rows of organizationally defined information that is
    // not recognized by the local agent.  If the local system is capable of
    // recognizing any organizationally defined information, appropriate extension
    // MIBs from the organization should be used for information retrieval.
    Lldpremorgdefinfotable LLDPMIB_Lldpremorgdefinfotable
}

func (lLDPMIB *LLDPMIB) GetFilter() yfilter.YFilter { return lLDPMIB.YFilter }

func (lLDPMIB *LLDPMIB) SetFilter(yf yfilter.YFilter) { lLDPMIB.YFilter = yf }

func (lLDPMIB *LLDPMIB) GetGoName(yname string) string {
    if yname == "lldpConfiguration" { return "Lldpconfiguration" }
    if yname == "lldpStatistics" { return "Lldpstatistics" }
    if yname == "lldpLocalSystemData" { return "Lldplocalsystemdata" }
    if yname == "lldpPortConfigTable" { return "Lldpportconfigtable" }
    if yname == "lldpStatsTxPortTable" { return "Lldpstatstxporttable" }
    if yname == "lldpStatsRxPortTable" { return "Lldpstatsrxporttable" }
    if yname == "lldpLocPortTable" { return "Lldplocporttable" }
    if yname == "lldpLocManAddrTable" { return "Lldplocmanaddrtable" }
    if yname == "lldpRemTable" { return "Lldpremtable" }
    if yname == "lldpRemManAddrTable" { return "Lldpremmanaddrtable" }
    if yname == "lldpRemUnknownTLVTable" { return "Lldpremunknowntlvtable" }
    if yname == "lldpRemOrgDefInfoTable" { return "Lldpremorgdefinfotable" }
    return ""
}

func (lLDPMIB *LLDPMIB) GetSegmentPath() string {
    return "LLDP-MIB:LLDP-MIB"
}

func (lLDPMIB *LLDPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpConfiguration" {
        return &lLDPMIB.Lldpconfiguration
    }
    if childYangName == "lldpStatistics" {
        return &lLDPMIB.Lldpstatistics
    }
    if childYangName == "lldpLocalSystemData" {
        return &lLDPMIB.Lldplocalsystemdata
    }
    if childYangName == "lldpPortConfigTable" {
        return &lLDPMIB.Lldpportconfigtable
    }
    if childYangName == "lldpStatsTxPortTable" {
        return &lLDPMIB.Lldpstatstxporttable
    }
    if childYangName == "lldpStatsRxPortTable" {
        return &lLDPMIB.Lldpstatsrxporttable
    }
    if childYangName == "lldpLocPortTable" {
        return &lLDPMIB.Lldplocporttable
    }
    if childYangName == "lldpLocManAddrTable" {
        return &lLDPMIB.Lldplocmanaddrtable
    }
    if childYangName == "lldpRemTable" {
        return &lLDPMIB.Lldpremtable
    }
    if childYangName == "lldpRemManAddrTable" {
        return &lLDPMIB.Lldpremmanaddrtable
    }
    if childYangName == "lldpRemUnknownTLVTable" {
        return &lLDPMIB.Lldpremunknowntlvtable
    }
    if childYangName == "lldpRemOrgDefInfoTable" {
        return &lLDPMIB.Lldpremorgdefinfotable
    }
    return nil
}

func (lLDPMIB *LLDPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lldpConfiguration"] = &lLDPMIB.Lldpconfiguration
    children["lldpStatistics"] = &lLDPMIB.Lldpstatistics
    children["lldpLocalSystemData"] = &lLDPMIB.Lldplocalsystemdata
    children["lldpPortConfigTable"] = &lLDPMIB.Lldpportconfigtable
    children["lldpStatsTxPortTable"] = &lLDPMIB.Lldpstatstxporttable
    children["lldpStatsRxPortTable"] = &lLDPMIB.Lldpstatsrxporttable
    children["lldpLocPortTable"] = &lLDPMIB.Lldplocporttable
    children["lldpLocManAddrTable"] = &lLDPMIB.Lldplocmanaddrtable
    children["lldpRemTable"] = &lLDPMIB.Lldpremtable
    children["lldpRemManAddrTable"] = &lLDPMIB.Lldpremmanaddrtable
    children["lldpRemUnknownTLVTable"] = &lLDPMIB.Lldpremunknowntlvtable
    children["lldpRemOrgDefInfoTable"] = &lLDPMIB.Lldpremorgdefinfotable
    return children
}

func (lLDPMIB *LLDPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lLDPMIB *LLDPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (lLDPMIB *LLDPMIB) GetYangName() string { return "LLDP-MIB" }

func (lLDPMIB *LLDPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lLDPMIB *LLDPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lLDPMIB *LLDPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lLDPMIB *LLDPMIB) SetParent(parent types.Entity) { lLDPMIB.parent = parent }

func (lLDPMIB *LLDPMIB) GetParent() types.Entity { return lLDPMIB.parent }

func (lLDPMIB *LLDPMIB) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpconfiguration
type LLDPMIB_Lldpconfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interval at which LLDP frames are transmitted on behalf of this LLDP
    // agent.  The default value for lldpMessageTxInterval object is 30 seconds. 
    // The value of this object must be restored from non-volatile storage after a
    // re-initialization of the management system. The type is interface{} with
    // range: 5..32768. Units are seconds.
    Lldpmessagetxinterval interface{}

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
    Lldpmessagetxholdmultiplier interface{}

    // The lldpReinitDelay indicates the delay (in units of seconds) from when
    // lldpPortConfigAdminStatus object of a particular port becomes 'disabled'
    // until re-initialization will be attempted.  The default value for
    // lldpReintDelay object is two seconds.  The value of this object must be
    // restored from non-volatile storage after a re-initialization of the
    // management system. The type is interface{} with range: 1..10. Units are
    // seconds.
    Lldpreinitdelay interface{}

    // The lldpTxDelay indicates the delay (in units of seconds) between
    // successive LLDP frame transmissions  initiated by value/status changes in
    // the LLDP local systems MIB.  The recommended value for the lldpTxDelay is
    // set by the following  formula:     1 <= lldpTxDelay <= (0.25 *
    // lldpMessageTxInterval)  The default value for lldpTxDelay object is two
    // seconds.  The value of this object must be restored from non-volatile
    // storage after a re-initialization of the management system. The type is
    // interface{} with range: 1..8192. Units are seconds.
    Lldptxdelay interface{}

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
    Lldpnotificationinterval interface{}
}

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetFilter() yfilter.YFilter { return lldpconfiguration.YFilter }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) SetFilter(yf yfilter.YFilter) { lldpconfiguration.YFilter = yf }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetGoName(yname string) string {
    if yname == "lldpMessageTxInterval" { return "Lldpmessagetxinterval" }
    if yname == "lldpMessageTxHoldMultiplier" { return "Lldpmessagetxholdmultiplier" }
    if yname == "lldpReinitDelay" { return "Lldpreinitdelay" }
    if yname == "lldpTxDelay" { return "Lldptxdelay" }
    if yname == "lldpNotificationInterval" { return "Lldpnotificationinterval" }
    return ""
}

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetSegmentPath() string {
    return "lldpConfiguration"
}

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpMessageTxInterval"] = lldpconfiguration.Lldpmessagetxinterval
    leafs["lldpMessageTxHoldMultiplier"] = lldpconfiguration.Lldpmessagetxholdmultiplier
    leafs["lldpReinitDelay"] = lldpconfiguration.Lldpreinitdelay
    leafs["lldpTxDelay"] = lldpconfiguration.Lldptxdelay
    leafs["lldpNotificationInterval"] = lldpconfiguration.Lldpnotificationinterval
    return leafs
}

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetBundleName() string { return "cisco_ios_xe" }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetYangName() string { return "lldpConfiguration" }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) SetParent(parent types.Entity) { lldpconfiguration.parent = parent }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetParent() types.Entity { return lldpconfiguration.parent }

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpstatistics
type LLDPMIB_Lldpstatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of sysUpTime object (defined in IETF RFC 3418) at the time an
    // entry is created, modified, or deleted in the in tables associated with the
    // lldpRemoteSystemsData objects and all LLDP extension objects associated
    // with remote systems.  An NMS can use this object to reduce polling of the
    // lldpRemoteSystemsData objects. The type is interface{} with range:
    // 0..4294967295.
    Lldpstatsremtableslastchangetime interface{}

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
    Lldpstatsremtablesinserts interface{}

    // The number of times the complete set of information advertised by a
    // particular MSAP has been deleted from tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects.  This counter should be
    // incremented only once when the complete set of information is completely
    // deleted from all related tables.  Partial deletions, such as deletion of
    // rows associated with a particular MSAP from some tables, but not from all
    // tables are not allowed, thus should not change the value of this counter.
    // The type is interface{} with range: 0..4294967295. Units are table entries.
    Lldpstatsremtablesdeletes interface{}

    // The number of times the complete set of information advertised by a
    // particular MSAP could not be entered into tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects because of insufficient
    // resources. The type is interface{} with range: 0..4294967295. Units are
    // table entries.
    Lldpstatsremtablesdrops interface{}

    // The number of times the complete set of information advertised by a
    // particular MSAP has been deleted from tables contained in
    // lldpRemoteSystemsData and lldpExtensions objects because the information
    // timeliness interval has expired.  This counter should be incremented only
    // once when the complete set of information is completely invalidated (aged
    // out) from all related tables.  Partial aging, similar to deletion case, is
    // not allowed, and thus, should not change the value of this counter. The
    // type is interface{} with range: 0..4294967295.
    Lldpstatsremtablesageouts interface{}
}

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetFilter() yfilter.YFilter { return lldpstatistics.YFilter }

func (lldpstatistics *LLDPMIB_Lldpstatistics) SetFilter(yf yfilter.YFilter) { lldpstatistics.YFilter = yf }

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetGoName(yname string) string {
    if yname == "lldpStatsRemTablesLastChangeTime" { return "Lldpstatsremtableslastchangetime" }
    if yname == "lldpStatsRemTablesInserts" { return "Lldpstatsremtablesinserts" }
    if yname == "lldpStatsRemTablesDeletes" { return "Lldpstatsremtablesdeletes" }
    if yname == "lldpStatsRemTablesDrops" { return "Lldpstatsremtablesdrops" }
    if yname == "lldpStatsRemTablesAgeouts" { return "Lldpstatsremtablesageouts" }
    return ""
}

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetSegmentPath() string {
    return "lldpStatistics"
}

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpStatsRemTablesLastChangeTime"] = lldpstatistics.Lldpstatsremtableslastchangetime
    leafs["lldpStatsRemTablesInserts"] = lldpstatistics.Lldpstatsremtablesinserts
    leafs["lldpStatsRemTablesDeletes"] = lldpstatistics.Lldpstatsremtablesdeletes
    leafs["lldpStatsRemTablesDrops"] = lldpstatistics.Lldpstatsremtablesdrops
    leafs["lldpStatsRemTablesAgeouts"] = lldpstatistics.Lldpstatsremtablesageouts
    return leafs
}

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetBundleName() string { return "cisco_ios_xe" }

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetYangName() string { return "lldpStatistics" }

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpstatistics *LLDPMIB_Lldpstatistics) SetParent(parent types.Entity) { lldpstatistics.parent = parent }

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetParent() types.Entity { return lldpstatistics.parent }

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldplocalsystemdata
type LLDPMIB_Lldplocalsystemdata struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type of encoding used to identify the chassis associated with the local
    // system. The type is LldpChassisIdSubtype.
    Lldplocchassisidsubtype interface{}

    // The string value used to identify the chassis component associated with the
    // local system. The type is string with length: 1..255.
    Lldplocchassisid interface{}

    // The string value used to identify the system name of the local system.  If
    // the local agent supports IETF RFC 3418, lldpLocSysName object should have
    // the same value of sysName object. The type is string with length: 0..255.
    Lldplocsysname interface{}

    // The string value used to identify the system description of the local
    // system.  If the local agent supports IETF RFC 3418, lldpLocSysDesc object
    // should have the same value of sysDesc object. The type is string with
    // length: 0..255.
    Lldplocsysdesc interface{}

    // The bitmap value used to identify which system capabilities are supported
    // on the local system. The type is map[string]bool.
    Lldplocsyscapsupported interface{}

    // The bitmap value used to identify which system capabilities are enabled on
    // the local system. The type is map[string]bool.
    Lldplocsyscapenabled interface{}
}

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetFilter() yfilter.YFilter { return lldplocalsystemdata.YFilter }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) SetFilter(yf yfilter.YFilter) { lldplocalsystemdata.YFilter = yf }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetGoName(yname string) string {
    if yname == "lldpLocChassisIdSubtype" { return "Lldplocchassisidsubtype" }
    if yname == "lldpLocChassisId" { return "Lldplocchassisid" }
    if yname == "lldpLocSysName" { return "Lldplocsysname" }
    if yname == "lldpLocSysDesc" { return "Lldplocsysdesc" }
    if yname == "lldpLocSysCapSupported" { return "Lldplocsyscapsupported" }
    if yname == "lldpLocSysCapEnabled" { return "Lldplocsyscapenabled" }
    return ""
}

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetSegmentPath() string {
    return "lldpLocalSystemData"
}

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpLocChassisIdSubtype"] = lldplocalsystemdata.Lldplocchassisidsubtype
    leafs["lldpLocChassisId"] = lldplocalsystemdata.Lldplocchassisid
    leafs["lldpLocSysName"] = lldplocalsystemdata.Lldplocsysname
    leafs["lldpLocSysDesc"] = lldplocalsystemdata.Lldplocsysdesc
    leafs["lldpLocSysCapSupported"] = lldplocalsystemdata.Lldplocsyscapsupported
    leafs["lldpLocSysCapEnabled"] = lldplocalsystemdata.Lldplocsyscapenabled
    return leafs
}

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetBundleName() string { return "cisco_ios_xe" }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetYangName() string { return "lldpLocalSystemData" }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) SetParent(parent types.Entity) { lldplocalsystemdata.parent = parent }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetParent() types.Entity { return lldplocalsystemdata.parent }

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpportconfigtable
// The table that controls LLDP frame transmission on individual
// ports.
type LLDPMIB_Lldpportconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LLDP configuration information for a particular port. This configuration
    // parameter controls the transmission and the reception of LLDP frames on
    // those ports whose rows are created in this table. The type is slice of
    // LLDPMIB_Lldpportconfigtable_Lldpportconfigentry.
    Lldpportconfigentry []LLDPMIB_Lldpportconfigtable_Lldpportconfigentry
}

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetFilter() yfilter.YFilter { return lldpportconfigtable.YFilter }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) SetFilter(yf yfilter.YFilter) { lldpportconfigtable.YFilter = yf }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetGoName(yname string) string {
    if yname == "lldpPortConfigEntry" { return "Lldpportconfigentry" }
    return ""
}

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetSegmentPath() string {
    return "lldpPortConfigTable"
}

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpPortConfigEntry" {
        for _, c := range lldpportconfigtable.Lldpportconfigentry {
            if lldpportconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldpportconfigtable_Lldpportconfigentry{}
        lldpportconfigtable.Lldpportconfigentry = append(lldpportconfigtable.Lldpportconfigentry, child)
        return &lldpportconfigtable.Lldpportconfigentry[len(lldpportconfigtable.Lldpportconfigentry)-1]
    }
    return nil
}

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpportconfigtable.Lldpportconfigentry {
        children[lldpportconfigtable.Lldpportconfigentry[i].GetSegmentPath()] = &lldpportconfigtable.Lldpportconfigentry[i]
    }
    return children
}

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetYangName() string { return "lldpPortConfigTable" }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) SetParent(parent types.Entity) { lldpportconfigtable.parent = parent }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetParent() types.Entity { return lldpportconfigtable.parent }

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpportconfigtable_Lldpportconfigentry
// LLDP configuration information for a particular port.
// This configuration parameter controls the transmission and
// the reception of LLDP frames on those ports whose rows are
// created in this table.
type LLDPMIB_Lldpportconfigtable_Lldpportconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpPortConfigTable. The type is interface{} with range: 1..4096.
    Lldpportconfigportnum interface{}

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
    // naturally age out. The type is Lldpportconfigadminstatus.
    Lldpportconfigadminstatus interface{}

    // The lldpPortConfigNotificationEnable controls, on a per port basis, 
    // whether or not notifications from the agent are enabled. The value true(1)
    // means that notifications are enabled; the value false(2) means that they
    // are not. The type is bool.
    Lldpportconfignotificationenable interface{}

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
    Lldpportconfigtlvstxenable interface{}
}

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetFilter() yfilter.YFilter { return lldpportconfigentry.YFilter }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) SetFilter(yf yfilter.YFilter) { lldpportconfigentry.YFilter = yf }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetGoName(yname string) string {
    if yname == "lldpPortConfigPortNum" { return "Lldpportconfigportnum" }
    if yname == "lldpPortConfigAdminStatus" { return "Lldpportconfigadminstatus" }
    if yname == "lldpPortConfigNotificationEnable" { return "Lldpportconfignotificationenable" }
    if yname == "lldpPortConfigTLVsTxEnable" { return "Lldpportconfigtlvstxenable" }
    return ""
}

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetSegmentPath() string {
    return "lldpPortConfigEntry" + "[lldpPortConfigPortNum='" + fmt.Sprintf("%v", lldpportconfigentry.Lldpportconfigportnum) + "']"
}

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpPortConfigPortNum"] = lldpportconfigentry.Lldpportconfigportnum
    leafs["lldpPortConfigAdminStatus"] = lldpportconfigentry.Lldpportconfigadminstatus
    leafs["lldpPortConfigNotificationEnable"] = lldpportconfigentry.Lldpportconfignotificationenable
    leafs["lldpPortConfigTLVsTxEnable"] = lldpportconfigentry.Lldpportconfigtlvstxenable
    return leafs
}

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetYangName() string { return "lldpPortConfigEntry" }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) SetParent(parent types.Entity) { lldpportconfigentry.parent = parent }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetParent() types.Entity { return lldpportconfigentry.parent }

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetParentYangName() string { return "lldpPortConfigTable" }

// LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus represents becomes disabled, then the information will naturally age out.
type LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus string

const (
    LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus_txOnly LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus = "txOnly"

    LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus_rxOnly LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus = "rxOnly"

    LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus_txAndRx LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus = "txAndRx"

    LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus_disabled LLDPMIB_Lldpportconfigtable_Lldpportconfigentry_Lldpportconfigadminstatus = "disabled"
)

// LLDPMIB_Lldpstatstxporttable
// A table containing LLDP transmission statistics for
// individual ports.  Entries are not required to exist in
// this table while the lldpPortConfigEntry object is equal to
// 'disabled(4)'.
type LLDPMIB_Lldpstatstxporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LLDP frame transmission statistics for a particular port. The port must be
    // contained in the same chassis as the LLDP agent.  All counter values in a
    // particular entry shall be maintained on a continuing basis and shall not be
    // deleted upon expiration of rxInfoTTL timing counters in the LLDP remote
    // systems MIB of the receipt of a shutdown frame from a remote LLDP agent. 
    // All statistical counters associated with a particular port on the local
    // LLDP agent become frozen whenever the adminStatus is disabled for the same
    // port. The type is slice of
    // LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry.
    Lldpstatstxportentry []LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry
}

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetFilter() yfilter.YFilter { return lldpstatstxporttable.YFilter }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) SetFilter(yf yfilter.YFilter) { lldpstatstxporttable.YFilter = yf }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetGoName(yname string) string {
    if yname == "lldpStatsTxPortEntry" { return "Lldpstatstxportentry" }
    return ""
}

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetSegmentPath() string {
    return "lldpStatsTxPortTable"
}

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpStatsTxPortEntry" {
        for _, c := range lldpstatstxporttable.Lldpstatstxportentry {
            if lldpstatstxporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry{}
        lldpstatstxporttable.Lldpstatstxportentry = append(lldpstatstxporttable.Lldpstatstxportentry, child)
        return &lldpstatstxporttable.Lldpstatstxportentry[len(lldpstatstxporttable.Lldpstatstxportentry)-1]
    }
    return nil
}

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpstatstxporttable.Lldpstatstxportentry {
        children[lldpstatstxporttable.Lldpstatstxportentry[i].GetSegmentPath()] = &lldpstatstxporttable.Lldpstatstxportentry[i]
    }
    return children
}

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetBundleName() string { return "cisco_ios_xe" }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetYangName() string { return "lldpStatsTxPortTable" }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) SetParent(parent types.Entity) { lldpstatstxporttable.parent = parent }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetParent() types.Entity { return lldpstatstxporttable.parent }

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry
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
type LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpStatsTable. The type is interface{} with range: 1..4096.
    Lldpstatstxportnum interface{}

    // The number of LLDP frames transmitted by this LLDP agent on the indicated
    // port. The type is interface{} with range: 0..4294967295.
    Lldpstatstxportframestotal interface{}
}

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetFilter() yfilter.YFilter { return lldpstatstxportentry.YFilter }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) SetFilter(yf yfilter.YFilter) { lldpstatstxportentry.YFilter = yf }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetGoName(yname string) string {
    if yname == "lldpStatsTxPortNum" { return "Lldpstatstxportnum" }
    if yname == "lldpStatsTxPortFramesTotal" { return "Lldpstatstxportframestotal" }
    return ""
}

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetSegmentPath() string {
    return "lldpStatsTxPortEntry" + "[lldpStatsTxPortNum='" + fmt.Sprintf("%v", lldpstatstxportentry.Lldpstatstxportnum) + "']"
}

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpStatsTxPortNum"] = lldpstatstxportentry.Lldpstatstxportnum
    leafs["lldpStatsTxPortFramesTotal"] = lldpstatstxportentry.Lldpstatstxportframestotal
    return leafs
}

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetBundleName() string { return "cisco_ios_xe" }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetYangName() string { return "lldpStatsTxPortEntry" }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) SetParent(parent types.Entity) { lldpstatstxportentry.parent = parent }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetParent() types.Entity { return lldpstatstxportentry.parent }

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetParentYangName() string { return "lldpStatsTxPortTable" }

// LLDPMIB_Lldpstatsrxporttable
// A table containing LLDP reception statistics for individual
// ports.  Entries are not required to exist in this table while
// the lldpPortConfigEntry object is equal to 'disabled(4)'.
type LLDPMIB_Lldpstatsrxporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LLDP frame reception statistics for a particular port. The port must be
    // contained in the same chassis as the LLDP agent.  All counter values in a
    // particular entry shall be maintained on a continuing basis and shall not be
    // deleted upon expiration of rxInfoTTL timing counters in the LLDP remote
    // systems MIB of the receipt of a shutdown frame from a remote LLDP agent. 
    // All statistical counters associated with a particular port on the local
    // LLDP agent become frozen whenever the adminStatus is disabled for the same
    // port. The type is slice of
    // LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry.
    Lldpstatsrxportentry []LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry
}

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetFilter() yfilter.YFilter { return lldpstatsrxporttable.YFilter }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) SetFilter(yf yfilter.YFilter) { lldpstatsrxporttable.YFilter = yf }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetGoName(yname string) string {
    if yname == "lldpStatsRxPortEntry" { return "Lldpstatsrxportentry" }
    return ""
}

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetSegmentPath() string {
    return "lldpStatsRxPortTable"
}

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpStatsRxPortEntry" {
        for _, c := range lldpstatsrxporttable.Lldpstatsrxportentry {
            if lldpstatsrxporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry{}
        lldpstatsrxporttable.Lldpstatsrxportentry = append(lldpstatsrxporttable.Lldpstatsrxportentry, child)
        return &lldpstatsrxporttable.Lldpstatsrxportentry[len(lldpstatsrxporttable.Lldpstatsrxportentry)-1]
    }
    return nil
}

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpstatsrxporttable.Lldpstatsrxportentry {
        children[lldpstatsrxporttable.Lldpstatsrxportentry[i].GetSegmentPath()] = &lldpstatsrxporttable.Lldpstatsrxportentry[i]
    }
    return children
}

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetBundleName() string { return "cisco_ios_xe" }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetYangName() string { return "lldpStatsRxPortTable" }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) SetParent(parent types.Entity) { lldpstatsrxporttable.parent = parent }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetParent() types.Entity { return lldpstatsrxporttable.parent }

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry
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
type LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpStatsTable. The type is interface{} with range: 1..4096.
    Lldpstatsrxportnum interface{}

    // The number of LLDP frames received by this LLDP agent on the indicated
    // port, and then discarded for any reason. This counter can provide an
    // indication that LLDP header formating problems may exist with the local
    // LLDP agent in the sending system or that LLDPDU validation problems may
    // exist with the local LLDP agent in the receiving system. The type is
    // interface{} with range: 0..4294967295.
    Lldpstatsrxportframesdiscardedtotal interface{}

    // The number of invalid LLDP frames received by this LLDP agent on the
    // indicated port, while this LLDP agent is enabled. The type is interface{}
    // with range: 0..4294967295.
    Lldpstatsrxportframeserrors interface{}

    // The number of valid LLDP frames received by this LLDP agent on the
    // indicated port, while this LLDP agent is enabled. The type is interface{}
    // with range: 0..4294967295.
    Lldpstatsrxportframestotal interface{}

    // The number of LLDP TLVs discarded for any reason by this LLDP agent on the
    // indicated port. The type is interface{} with range: 0..4294967295.
    Lldpstatsrxporttlvsdiscardedtotal interface{}

    // The number of LLDP TLVs received on the given port that are not recognized
    // by this LLDP agent on the indicated port.  An unrecognized TLV is referred
    // to as the TLV whose type value is in the range of reserved TLV types (000
    // 1001 - 111 1110) in Table 9.1 of IEEE Std 802.1AB-2005.  An unrecognized
    // TLV may be a basic management TLV from a later LLDP version. The type is
    // interface{} with range: 0..4294967295.
    Lldpstatsrxporttlvsunrecognizedtotal interface{}

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
    Lldpstatsrxportageoutstotal interface{}
}

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetFilter() yfilter.YFilter { return lldpstatsrxportentry.YFilter }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) SetFilter(yf yfilter.YFilter) { lldpstatsrxportentry.YFilter = yf }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetGoName(yname string) string {
    if yname == "lldpStatsRxPortNum" { return "Lldpstatsrxportnum" }
    if yname == "lldpStatsRxPortFramesDiscardedTotal" { return "Lldpstatsrxportframesdiscardedtotal" }
    if yname == "lldpStatsRxPortFramesErrors" { return "Lldpstatsrxportframeserrors" }
    if yname == "lldpStatsRxPortFramesTotal" { return "Lldpstatsrxportframestotal" }
    if yname == "lldpStatsRxPortTLVsDiscardedTotal" { return "Lldpstatsrxporttlvsdiscardedtotal" }
    if yname == "lldpStatsRxPortTLVsUnrecognizedTotal" { return "Lldpstatsrxporttlvsunrecognizedtotal" }
    if yname == "lldpStatsRxPortAgeoutsTotal" { return "Lldpstatsrxportageoutstotal" }
    return ""
}

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetSegmentPath() string {
    return "lldpStatsRxPortEntry" + "[lldpStatsRxPortNum='" + fmt.Sprintf("%v", lldpstatsrxportentry.Lldpstatsrxportnum) + "']"
}

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpStatsRxPortNum"] = lldpstatsrxportentry.Lldpstatsrxportnum
    leafs["lldpStatsRxPortFramesDiscardedTotal"] = lldpstatsrxportentry.Lldpstatsrxportframesdiscardedtotal
    leafs["lldpStatsRxPortFramesErrors"] = lldpstatsrxportentry.Lldpstatsrxportframeserrors
    leafs["lldpStatsRxPortFramesTotal"] = lldpstatsrxportentry.Lldpstatsrxportframestotal
    leafs["lldpStatsRxPortTLVsDiscardedTotal"] = lldpstatsrxportentry.Lldpstatsrxporttlvsdiscardedtotal
    leafs["lldpStatsRxPortTLVsUnrecognizedTotal"] = lldpstatsrxportentry.Lldpstatsrxporttlvsunrecognizedtotal
    leafs["lldpStatsRxPortAgeoutsTotal"] = lldpstatsrxportentry.Lldpstatsrxportageoutstotal
    return leafs
}

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetBundleName() string { return "cisco_ios_xe" }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetYangName() string { return "lldpStatsRxPortEntry" }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) SetParent(parent types.Entity) { lldpstatsrxportentry.parent = parent }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetParent() types.Entity { return lldpstatsrxportentry.parent }

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetParentYangName() string { return "lldpStatsRxPortTable" }

// LLDPMIB_Lldplocporttable
// This table contains one or more rows per port information
// associated with the local system known to this agent.
type LLDPMIB_Lldplocporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular port component.  Entries may be created and
    // deleted in this table by the agent. The type is slice of
    // LLDPMIB_Lldplocporttable_Lldplocportentry.
    Lldplocportentry []LLDPMIB_Lldplocporttable_Lldplocportentry
}

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetFilter() yfilter.YFilter { return lldplocporttable.YFilter }

func (lldplocporttable *LLDPMIB_Lldplocporttable) SetFilter(yf yfilter.YFilter) { lldplocporttable.YFilter = yf }

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetGoName(yname string) string {
    if yname == "lldpLocPortEntry" { return "Lldplocportentry" }
    return ""
}

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetSegmentPath() string {
    return "lldpLocPortTable"
}

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpLocPortEntry" {
        for _, c := range lldplocporttable.Lldplocportentry {
            if lldplocporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldplocporttable_Lldplocportentry{}
        lldplocporttable.Lldplocportentry = append(lldplocporttable.Lldplocportentry, child)
        return &lldplocporttable.Lldplocportentry[len(lldplocporttable.Lldplocportentry)-1]
    }
    return nil
}

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldplocporttable.Lldplocportentry {
        children[lldplocporttable.Lldplocportentry[i].GetSegmentPath()] = &lldplocporttable.Lldplocportentry[i]
    }
    return children
}

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetBundleName() string { return "cisco_ios_xe" }

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetYangName() string { return "lldpLocPortTable" }

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldplocporttable *LLDPMIB_Lldplocporttable) SetParent(parent types.Entity) { lldplocporttable.parent = parent }

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetParent() types.Entity { return lldplocporttable.parent }

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldplocporttable_Lldplocportentry
// Information about a particular port component.
// 
// Entries may be created and deleted in this table by the
// agent.
type LLDPMIB_Lldplocporttable_Lldplocportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The value of this object is used as a port index to the
    // lldpLocPortTable. The type is interface{} with range: 1..4096.
    Lldplocportnum interface{}

    // The type of port identifier encoding used in the associated 'lldpLocPortId'
    // object. The type is LldpPortIdSubtype.
    Lldplocportidsubtype interface{}

    // The string value used to identify the port component associated with a
    // given port in the local system. The type is string with length: 1..255.
    Lldplocportid interface{}

    // The string value used to identify the 802 LAN station's port description
    // associated with the local system.  If the local agent supports IETF RFC
    // 2863, lldpLocPortDesc object should have the same value of ifDescr object.
    // The type is string with length: 0..255.
    Lldplocportdesc interface{}
}

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetFilter() yfilter.YFilter { return lldplocportentry.YFilter }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) SetFilter(yf yfilter.YFilter) { lldplocportentry.YFilter = yf }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetGoName(yname string) string {
    if yname == "lldpLocPortNum" { return "Lldplocportnum" }
    if yname == "lldpLocPortIdSubtype" { return "Lldplocportidsubtype" }
    if yname == "lldpLocPortId" { return "Lldplocportid" }
    if yname == "lldpLocPortDesc" { return "Lldplocportdesc" }
    return ""
}

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetSegmentPath() string {
    return "lldpLocPortEntry" + "[lldpLocPortNum='" + fmt.Sprintf("%v", lldplocportentry.Lldplocportnum) + "']"
}

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpLocPortNum"] = lldplocportentry.Lldplocportnum
    leafs["lldpLocPortIdSubtype"] = lldplocportentry.Lldplocportidsubtype
    leafs["lldpLocPortId"] = lldplocportentry.Lldplocportid
    leafs["lldpLocPortDesc"] = lldplocportentry.Lldplocportdesc
    return leafs
}

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetBundleName() string { return "cisco_ios_xe" }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetYangName() string { return "lldpLocPortEntry" }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) SetParent(parent types.Entity) { lldplocportentry.parent = parent }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetParent() types.Entity { return lldplocportentry.parent }

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetParentYangName() string { return "lldpLocPortTable" }

// LLDPMIB_Lldplocmanaddrtable
// This table contains management address information on the
// local system known to this agent.
type LLDPMIB_Lldplocmanaddrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Management address information about a particular chassis component.  There
    // may be multiple management addresses configured on the system identified by
    // a particular lldpLocChassisId.  Each management address should have
    // distinct 'management address type' (lldpLocManAddrSubtype) and 'management
    // address' (lldpLocManAddr.)  Entries may be created and deleted in this
    // table by the agent. The type is slice of
    // LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry.
    Lldplocmanaddrentry []LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry
}

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetFilter() yfilter.YFilter { return lldplocmanaddrtable.YFilter }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) SetFilter(yf yfilter.YFilter) { lldplocmanaddrtable.YFilter = yf }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetGoName(yname string) string {
    if yname == "lldpLocManAddrEntry" { return "Lldplocmanaddrentry" }
    return ""
}

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetSegmentPath() string {
    return "lldpLocManAddrTable"
}

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpLocManAddrEntry" {
        for _, c := range lldplocmanaddrtable.Lldplocmanaddrentry {
            if lldplocmanaddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry{}
        lldplocmanaddrtable.Lldplocmanaddrentry = append(lldplocmanaddrtable.Lldplocmanaddrentry, child)
        return &lldplocmanaddrtable.Lldplocmanaddrentry[len(lldplocmanaddrtable.Lldplocmanaddrentry)-1]
    }
    return nil
}

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldplocmanaddrtable.Lldplocmanaddrentry {
        children[lldplocmanaddrtable.Lldplocmanaddrentry[i].GetSegmentPath()] = &lldplocmanaddrtable.Lldplocmanaddrentry[i]
    }
    return children
}

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetYangName() string { return "lldpLocManAddrTable" }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) SetParent(parent types.Entity) { lldplocmanaddrtable.parent = parent }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetParent() types.Entity { return lldplocmanaddrtable.parent }

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry
// Management address information about a particular chassis
// component.  There may be multiple management addresses
// configured on the system identified by a particular
// lldpLocChassisId.  Each management address should have
// distinct 'management address type' (lldpLocManAddrSubtype) and
// 'management address' (lldpLocManAddr.)
// 
// Entries may be created and deleted in this table by the
// agent.
type LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type of management address identifier encoding
    // used in the associated 'lldpLocManagmentAddr' object. The type is
    // AddressFamilyNumbers.
    Lldplocmanaddrsubtype interface{}

    // This attribute is a key. The string value used to identify the management
    // address component associated with the local system.  The purpose of this
    // address is to contact the management entity. The type is string with
    // length: 1..31.
    Lldplocmanaddr interface{}

    // The total length of the management address subtype and the management
    // address fields in LLDPDUs transmitted by the local LLDP agent.  The
    // management address length field is needed so that the receiving systems
    // that do not implement SNMP will not be required to implement an iana family
    // numbers/address length equivalency table in order to decode the management
    // adress. The type is interface{} with range: -2147483648..2147483647.
    Lldplocmanaddrlen interface{}

    // The enumeration value that identifies the interface numbering method used
    // for defining the interface number, associated with the local system. The
    // type is LldpManAddrIfSubtype.
    Lldplocmanaddrifsubtype interface{}

    // The integer value used to identify the interface number regarding the
    // management address component associated with the local system. The type is
    // interface{} with range: -2147483648..2147483647.
    Lldplocmanaddrifid interface{}

    // The OID value used to identify the type of hardware component or protocol
    // entity associated with the management address advertised by the local
    // system agent. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Lldplocmanaddroid interface{}

    // A set of ports that are identified by a PortList, in which each port is
    // represented as a bit.  The corresponding local system management address
    // instance will be transmitted on the member ports of the
    // lldpManAddrPortsTxEnable.    The default value for
    // lldpConfigManAddrPortsTxEnable object is empty binary string, which means
    // no ports are specified for advertising indicated management address
    // instance. The type is string with length: 0..512.
    Lldpconfigmanaddrportstxenable interface{}
}

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetFilter() yfilter.YFilter { return lldplocmanaddrentry.YFilter }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) SetFilter(yf yfilter.YFilter) { lldplocmanaddrentry.YFilter = yf }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetGoName(yname string) string {
    if yname == "lldpLocManAddrSubtype" { return "Lldplocmanaddrsubtype" }
    if yname == "lldpLocManAddr" { return "Lldplocmanaddr" }
    if yname == "lldpLocManAddrLen" { return "Lldplocmanaddrlen" }
    if yname == "lldpLocManAddrIfSubtype" { return "Lldplocmanaddrifsubtype" }
    if yname == "lldpLocManAddrIfId" { return "Lldplocmanaddrifid" }
    if yname == "lldpLocManAddrOID" { return "Lldplocmanaddroid" }
    if yname == "lldpConfigManAddrPortsTxEnable" { return "Lldpconfigmanaddrportstxenable" }
    return ""
}

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetSegmentPath() string {
    return "lldpLocManAddrEntry" + "[lldpLocManAddrSubtype='" + fmt.Sprintf("%v", lldplocmanaddrentry.Lldplocmanaddrsubtype) + "']" + "[lldpLocManAddr='" + fmt.Sprintf("%v", lldplocmanaddrentry.Lldplocmanaddr) + "']"
}

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpLocManAddrSubtype"] = lldplocmanaddrentry.Lldplocmanaddrsubtype
    leafs["lldpLocManAddr"] = lldplocmanaddrentry.Lldplocmanaddr
    leafs["lldpLocManAddrLen"] = lldplocmanaddrentry.Lldplocmanaddrlen
    leafs["lldpLocManAddrIfSubtype"] = lldplocmanaddrentry.Lldplocmanaddrifsubtype
    leafs["lldpLocManAddrIfId"] = lldplocmanaddrentry.Lldplocmanaddrifid
    leafs["lldpLocManAddrOID"] = lldplocmanaddrentry.Lldplocmanaddroid
    leafs["lldpConfigManAddrPortsTxEnable"] = lldplocmanaddrentry.Lldpconfigmanaddrportstxenable
    return leafs
}

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetYangName() string { return "lldpLocManAddrEntry" }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) SetParent(parent types.Entity) { lldplocmanaddrentry.parent = parent }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetParent() types.Entity { return lldplocmanaddrentry.parent }

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetParentYangName() string { return "lldpLocManAddrTable" }

// LLDPMIB_Lldpremtable
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
type LLDPMIB_Lldpremtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular physical network connection. Entries may be
    // created and deleted in this table by the agent, if a physical topology
    // discovery process is active. The type is slice of
    // LLDPMIB_Lldpremtable_Lldprementry.
    Lldprementry []LLDPMIB_Lldpremtable_Lldprementry
}

func (lldpremtable *LLDPMIB_Lldpremtable) GetFilter() yfilter.YFilter { return lldpremtable.YFilter }

func (lldpremtable *LLDPMIB_Lldpremtable) SetFilter(yf yfilter.YFilter) { lldpremtable.YFilter = yf }

func (lldpremtable *LLDPMIB_Lldpremtable) GetGoName(yname string) string {
    if yname == "lldpRemEntry" { return "Lldprementry" }
    return ""
}

func (lldpremtable *LLDPMIB_Lldpremtable) GetSegmentPath() string {
    return "lldpRemTable"
}

func (lldpremtable *LLDPMIB_Lldpremtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpRemEntry" {
        for _, c := range lldpremtable.Lldprementry {
            if lldpremtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldpremtable_Lldprementry{}
        lldpremtable.Lldprementry = append(lldpremtable.Lldprementry, child)
        return &lldpremtable.Lldprementry[len(lldpremtable.Lldprementry)-1]
    }
    return nil
}

func (lldpremtable *LLDPMIB_Lldpremtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpremtable.Lldprementry {
        children[lldpremtable.Lldprementry[i].GetSegmentPath()] = &lldpremtable.Lldprementry[i]
    }
    return children
}

func (lldpremtable *LLDPMIB_Lldpremtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpremtable *LLDPMIB_Lldpremtable) GetBundleName() string { return "cisco_ios_xe" }

func (lldpremtable *LLDPMIB_Lldpremtable) GetYangName() string { return "lldpRemTable" }

func (lldpremtable *LLDPMIB_Lldpremtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpremtable *LLDPMIB_Lldpremtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpremtable *LLDPMIB_Lldpremtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpremtable *LLDPMIB_Lldpremtable) SetParent(parent types.Entity) { lldpremtable.parent = parent }

func (lldpremtable *LLDPMIB_Lldpremtable) GetParent() types.Entity { return lldpremtable.parent }

func (lldpremtable *LLDPMIB_Lldpremtable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpremtable_Lldprementry
// Information about a particular physical network connection.
// Entries may be created and deleted in this table by the agent,
// if a physical topology discovery process is active.
type LLDPMIB_Lldpremtable_Lldprementry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A TimeFilter for this entry.  See the TimeFilter
    // textual convention in IETF RFC 2021 and 
    // http://www.ietf.org/IESG/Implementations/RFC2021-Implementation.txt to see
    // how TimeFilter works. The type is interface{} with range: 0..4294967295.
    Lldpremtimemark interface{}

    // This attribute is a key. The index value used to identify the port
    // component (contained in the local chassis with the LLDP agent) associated
    // with this entry.  The lldpRemLocalPortNum identifies the port on which the
    // remote system information is received.  The value of this object is used as
    // a port index to the lldpRemTable. The type is interface{} with range:
    // 1..4096.
    Lldpremlocalportnum interface{}

    // This attribute is a key. This object represents an arbitrary local integer
    // value used by this agent to identify a particular connection instance,
    // unique only for the indicated remote system.  An agent is encouraged to
    // assign monotonically increasing index values to new entries, starting with
    // one, after each reboot.  It is considered unlikely that the lldpRemIndex
    // will wrap between reboots. The type is interface{} with range:
    // 1..2147483647.
    Lldpremindex interface{}

    // The type of encoding used to identify the chassis associated with the
    // remote system. The type is LldpChassisIdSubtype.
    Lldpremchassisidsubtype interface{}

    // The string value used to identify the chassis component associated with the
    // remote system. The type is string with length: 1..255.
    Lldpremchassisid interface{}

    // The type of port identifier encoding used in the associated 'lldpRemPortId'
    // object. The type is LldpPortIdSubtype.
    Lldpremportidsubtype interface{}

    // The string value used to identify the port component associated with the
    // remote system. The type is string with length: 1..255.
    Lldpremportid interface{}

    // The string value used to identify the description of the given port
    // associated with the remote system. The type is string with length: 0..255.
    Lldpremportdesc interface{}

    // The string value used to identify the system name of the remote system. The
    // type is string with length: 0..255.
    Lldpremsysname interface{}

    // The string value used to identify the system description of the remote
    // system. The type is string with length: 0..255.
    Lldpremsysdesc interface{}

    // The bitmap value used to identify which system capabilities are supported
    // on the remote system. The type is map[string]bool.
    Lldpremsyscapsupported interface{}

    // The bitmap value used to identify which system capabilities are enabled on
    // the remote system. The type is map[string]bool.
    Lldpremsyscapenabled interface{}
}

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetFilter() yfilter.YFilter { return lldprementry.YFilter }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) SetFilter(yf yfilter.YFilter) { lldprementry.YFilter = yf }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetGoName(yname string) string {
    if yname == "lldpRemTimeMark" { return "Lldpremtimemark" }
    if yname == "lldpRemLocalPortNum" { return "Lldpremlocalportnum" }
    if yname == "lldpRemIndex" { return "Lldpremindex" }
    if yname == "lldpRemChassisIdSubtype" { return "Lldpremchassisidsubtype" }
    if yname == "lldpRemChassisId" { return "Lldpremchassisid" }
    if yname == "lldpRemPortIdSubtype" { return "Lldpremportidsubtype" }
    if yname == "lldpRemPortId" { return "Lldpremportid" }
    if yname == "lldpRemPortDesc" { return "Lldpremportdesc" }
    if yname == "lldpRemSysName" { return "Lldpremsysname" }
    if yname == "lldpRemSysDesc" { return "Lldpremsysdesc" }
    if yname == "lldpRemSysCapSupported" { return "Lldpremsyscapsupported" }
    if yname == "lldpRemSysCapEnabled" { return "Lldpremsyscapenabled" }
    return ""
}

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetSegmentPath() string {
    return "lldpRemEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldprementry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldprementry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldprementry.Lldpremindex) + "']"
}

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpRemTimeMark"] = lldprementry.Lldpremtimemark
    leafs["lldpRemLocalPortNum"] = lldprementry.Lldpremlocalportnum
    leafs["lldpRemIndex"] = lldprementry.Lldpremindex
    leafs["lldpRemChassisIdSubtype"] = lldprementry.Lldpremchassisidsubtype
    leafs["lldpRemChassisId"] = lldprementry.Lldpremchassisid
    leafs["lldpRemPortIdSubtype"] = lldprementry.Lldpremportidsubtype
    leafs["lldpRemPortId"] = lldprementry.Lldpremportid
    leafs["lldpRemPortDesc"] = lldprementry.Lldpremportdesc
    leafs["lldpRemSysName"] = lldprementry.Lldpremsysname
    leafs["lldpRemSysDesc"] = lldprementry.Lldpremsysdesc
    leafs["lldpRemSysCapSupported"] = lldprementry.Lldpremsyscapsupported
    leafs["lldpRemSysCapEnabled"] = lldprementry.Lldpremsyscapenabled
    return leafs
}

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetBundleName() string { return "cisco_ios_xe" }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetYangName() string { return "lldpRemEntry" }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) SetParent(parent types.Entity) { lldprementry.parent = parent }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetParent() types.Entity { return lldprementry.parent }

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetParentYangName() string { return "lldpRemTable" }

// LLDPMIB_Lldpremmanaddrtable
// This table contains one or more rows per management address
// information on the remote system learned on a particular port
// contained in the local chassis known to this agent.
type LLDPMIB_Lldpremmanaddrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Management address information about a particular chassis component.  There
    // may be multiple management addresses configured on the remote system
    // identified by a particular lldpRemIndex whose information is received on
    // lldpRemLocalPortNum of the local system.  Each management address should
    // have distinct 'management address type' (lldpRemManAddrSubtype) and
    // 'management address' (lldpRemManAddr.)  Entries may be created and deleted
    // in this table by the agent. The type is slice of
    // LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry.
    Lldpremmanaddrentry []LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry
}

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetFilter() yfilter.YFilter { return lldpremmanaddrtable.YFilter }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) SetFilter(yf yfilter.YFilter) { lldpremmanaddrtable.YFilter = yf }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetGoName(yname string) string {
    if yname == "lldpRemManAddrEntry" { return "Lldpremmanaddrentry" }
    return ""
}

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetSegmentPath() string {
    return "lldpRemManAddrTable"
}

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpRemManAddrEntry" {
        for _, c := range lldpremmanaddrtable.Lldpremmanaddrentry {
            if lldpremmanaddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry{}
        lldpremmanaddrtable.Lldpremmanaddrentry = append(lldpremmanaddrtable.Lldpremmanaddrentry, child)
        return &lldpremmanaddrtable.Lldpremmanaddrentry[len(lldpremmanaddrtable.Lldpremmanaddrentry)-1]
    }
    return nil
}

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpremmanaddrtable.Lldpremmanaddrentry {
        children[lldpremmanaddrtable.Lldpremmanaddrentry[i].GetSegmentPath()] = &lldpremmanaddrtable.Lldpremmanaddrentry[i]
    }
    return children
}

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetYangName() string { return "lldpRemManAddrTable" }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) SetParent(parent types.Entity) { lldpremmanaddrtable.parent = parent }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetParent() types.Entity { return lldpremmanaddrtable.parent }

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry
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
type LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremtimemark
    Lldpremtimemark interface{}

    // This attribute is a key. The type is string with range: 1..4096. Refers to
    // lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremlocalportnum
    Lldpremlocalportnum interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremindex
    Lldpremindex interface{}

    // This attribute is a key. The type of management address identifier encoding
    // used in the associated 'lldpRemManagmentAddr' object. The type is
    // AddressFamilyNumbers.
    Lldpremmanaddrsubtype interface{}

    // This attribute is a key. The string value used to identify the management
    // address component associated with the remote system.  The purpose of this
    // address is to contact the management entity. The type is string with
    // length: 1..31.
    Lldpremmanaddr interface{}

    // The enumeration value that identifies the interface numbering method used
    // for defining the interface number, associated with the remote system. The
    // type is LldpManAddrIfSubtype.
    Lldpremmanaddrifsubtype interface{}

    // The integer value used to identify the interface number regarding the
    // management address component associated with the remote system. The type is
    // interface{} with range: -2147483648..2147483647.
    Lldpremmanaddrifid interface{}

    // The OID value used to identify the type of hardware component or protocol
    // entity associated with the management address advertised by the remote
    // system agent. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Lldpremmanaddroid interface{}
}

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetFilter() yfilter.YFilter { return lldpremmanaddrentry.YFilter }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) SetFilter(yf yfilter.YFilter) { lldpremmanaddrentry.YFilter = yf }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetGoName(yname string) string {
    if yname == "lldpRemTimeMark" { return "Lldpremtimemark" }
    if yname == "lldpRemLocalPortNum" { return "Lldpremlocalportnum" }
    if yname == "lldpRemIndex" { return "Lldpremindex" }
    if yname == "lldpRemManAddrSubtype" { return "Lldpremmanaddrsubtype" }
    if yname == "lldpRemManAddr" { return "Lldpremmanaddr" }
    if yname == "lldpRemManAddrIfSubtype" { return "Lldpremmanaddrifsubtype" }
    if yname == "lldpRemManAddrIfId" { return "Lldpremmanaddrifid" }
    if yname == "lldpRemManAddrOID" { return "Lldpremmanaddroid" }
    return ""
}

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetSegmentPath() string {
    return "lldpRemManAddrEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremindex) + "']" + "[lldpRemManAddrSubtype='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremmanaddrsubtype) + "']" + "[lldpRemManAddr='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremmanaddr) + "']"
}

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpRemTimeMark"] = lldpremmanaddrentry.Lldpremtimemark
    leafs["lldpRemLocalPortNum"] = lldpremmanaddrentry.Lldpremlocalportnum
    leafs["lldpRemIndex"] = lldpremmanaddrentry.Lldpremindex
    leafs["lldpRemManAddrSubtype"] = lldpremmanaddrentry.Lldpremmanaddrsubtype
    leafs["lldpRemManAddr"] = lldpremmanaddrentry.Lldpremmanaddr
    leafs["lldpRemManAddrIfSubtype"] = lldpremmanaddrentry.Lldpremmanaddrifsubtype
    leafs["lldpRemManAddrIfId"] = lldpremmanaddrentry.Lldpremmanaddrifid
    leafs["lldpRemManAddrOID"] = lldpremmanaddrentry.Lldpremmanaddroid
    return leafs
}

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetYangName() string { return "lldpRemManAddrEntry" }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) SetParent(parent types.Entity) { lldpremmanaddrentry.parent = parent }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetParent() types.Entity { return lldpremmanaddrentry.parent }

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetParentYangName() string { return "lldpRemManAddrTable" }

// LLDPMIB_Lldpremunknowntlvtable
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
type LLDPMIB_Lldpremunknowntlvtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an unrecognized TLV received from a physical network
    // connection.  Entries may be created and deleted in this table by the agent,
    // if a physical topology discovery process is active. The type is slice of
    // LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry.
    Lldpremunknowntlventry []LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry
}

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetFilter() yfilter.YFilter { return lldpremunknowntlvtable.YFilter }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) SetFilter(yf yfilter.YFilter) { lldpremunknowntlvtable.YFilter = yf }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetGoName(yname string) string {
    if yname == "lldpRemUnknownTLVEntry" { return "Lldpremunknowntlventry" }
    return ""
}

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetSegmentPath() string {
    return "lldpRemUnknownTLVTable"
}

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpRemUnknownTLVEntry" {
        for _, c := range lldpremunknowntlvtable.Lldpremunknowntlventry {
            if lldpremunknowntlvtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry{}
        lldpremunknowntlvtable.Lldpremunknowntlventry = append(lldpremunknowntlvtable.Lldpremunknowntlventry, child)
        return &lldpremunknowntlvtable.Lldpremunknowntlventry[len(lldpremunknowntlvtable.Lldpremunknowntlventry)-1]
    }
    return nil
}

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpremunknowntlvtable.Lldpremunknowntlventry {
        children[lldpremunknowntlvtable.Lldpremunknowntlventry[i].GetSegmentPath()] = &lldpremunknowntlvtable.Lldpremunknowntlventry[i]
    }
    return children
}

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetBundleName() string { return "cisco_ios_xe" }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetYangName() string { return "lldpRemUnknownTLVTable" }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) SetParent(parent types.Entity) { lldpremunknowntlvtable.parent = parent }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetParent() types.Entity { return lldpremunknowntlvtable.parent }

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry
// Information about an unrecognized TLV received from a
// physical network connection.  Entries may be created and
// deleted in this table by the agent, if a physical topology
// discovery process is active.
type LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremtimemark
    Lldpremtimemark interface{}

    // This attribute is a key. The type is string with range: 1..4096. Refers to
    // lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremlocalportnum
    Lldpremlocalportnum interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremindex
    Lldpremindex interface{}

    // This attribute is a key. This object represents the value extracted from
    // the type field of the TLV. The type is interface{} with range: 9..126.
    Lldpremunknowntlvtype interface{}

    // This object represents the value extracted from the value field of the TLV.
    // The type is string with length: 0..511.
    Lldpremunknowntlvinfo interface{}
}

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetFilter() yfilter.YFilter { return lldpremunknowntlventry.YFilter }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) SetFilter(yf yfilter.YFilter) { lldpremunknowntlventry.YFilter = yf }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetGoName(yname string) string {
    if yname == "lldpRemTimeMark" { return "Lldpremtimemark" }
    if yname == "lldpRemLocalPortNum" { return "Lldpremlocalportnum" }
    if yname == "lldpRemIndex" { return "Lldpremindex" }
    if yname == "lldpRemUnknownTLVType" { return "Lldpremunknowntlvtype" }
    if yname == "lldpRemUnknownTLVInfo" { return "Lldpremunknowntlvinfo" }
    return ""
}

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetSegmentPath() string {
    return "lldpRemUnknownTLVEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremindex) + "']" + "[lldpRemUnknownTLVType='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremunknowntlvtype) + "']"
}

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpRemTimeMark"] = lldpremunknowntlventry.Lldpremtimemark
    leafs["lldpRemLocalPortNum"] = lldpremunknowntlventry.Lldpremlocalportnum
    leafs["lldpRemIndex"] = lldpremunknowntlventry.Lldpremindex
    leafs["lldpRemUnknownTLVType"] = lldpremunknowntlventry.Lldpremunknowntlvtype
    leafs["lldpRemUnknownTLVInfo"] = lldpremunknowntlventry.Lldpremunknowntlvinfo
    return leafs
}

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetBundleName() string { return "cisco_ios_xe" }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetYangName() string { return "lldpRemUnknownTLVEntry" }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) SetParent(parent types.Entity) { lldpremunknowntlventry.parent = parent }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetParent() types.Entity { return lldpremunknowntlventry.parent }

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetParentYangName() string { return "lldpRemUnknownTLVTable" }

// LLDPMIB_Lldpremorgdefinfotable
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
type LLDPMIB_Lldpremorgdefinfotable struct {
    parent types.Entity
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
    // LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry.
    Lldpremorgdefinfoentry []LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry
}

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetFilter() yfilter.YFilter { return lldpremorgdefinfotable.YFilter }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) SetFilter(yf yfilter.YFilter) { lldpremorgdefinfotable.YFilter = yf }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetGoName(yname string) string {
    if yname == "lldpRemOrgDefInfoEntry" { return "Lldpremorgdefinfoentry" }
    return ""
}

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetSegmentPath() string {
    return "lldpRemOrgDefInfoTable"
}

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldpRemOrgDefInfoEntry" {
        for _, c := range lldpremorgdefinfotable.Lldpremorgdefinfoentry {
            if lldpremorgdefinfotable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry{}
        lldpremorgdefinfotable.Lldpremorgdefinfoentry = append(lldpremorgdefinfotable.Lldpremorgdefinfoentry, child)
        return &lldpremorgdefinfotable.Lldpremorgdefinfoentry[len(lldpremorgdefinfotable.Lldpremorgdefinfoentry)-1]
    }
    return nil
}

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lldpremorgdefinfotable.Lldpremorgdefinfoentry {
        children[lldpremorgdefinfotable.Lldpremorgdefinfoentry[i].GetSegmentPath()] = &lldpremorgdefinfotable.Lldpremorgdefinfoentry[i]
    }
    return children
}

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetBundleName() string { return "cisco_ios_xe" }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetYangName() string { return "lldpRemOrgDefInfoTable" }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) SetParent(parent types.Entity) { lldpremorgdefinfotable.parent = parent }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetParent() types.Entity { return lldpremorgdefinfotable.parent }

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetParentYangName() string { return "LLDP-MIB" }

// LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry
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
type LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremtimemark
    Lldpremtimemark interface{}

    // This attribute is a key. The type is string with range: 1..4096. Refers to
    // lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremlocalportnum
    Lldpremlocalportnum interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to lldp_mib.LLDPMIB_Lldpremtable_Lldprementry_Lldpremindex
    Lldpremindex interface{}

    // This attribute is a key. The Organizationally Unique Identifier (OUI), as
    // defined in IEEE std 802-2001, is a 24 bit (three octets) globally unique
    // assigned number referenced by various standards, of the information
    // received from the remote system. The type is string with length: 3.
    Lldpremorgdefinfooui interface{}

    // This attribute is a key. The integer value used to identify the subtype of
    // the organizationally defined information received from the remote system. 
    // The subtype value is required to identify different instances of
    // organizationally defined information that could not be retrieved without a
    // unique identifier that indicates the particular type of information
    // contained in the information string. The type is interface{} with range:
    // 1..255.
    Lldpremorgdefinfosubtype interface{}

    // This attribute is a key. This object represents an arbitrary local integer
    // value used by this agent to identify a particular unrecognized
    // organizationally defined information instance, unique only for the
    // lldpRemOrgDefInfoOUI and lldpRemOrgDefInfoSubtype from the same remote
    // system.  An agent is encouraged to assign monotonically increasing index
    // values to new entries, starting with one, after each reboot.  It is
    // considered unlikely that the lldpRemOrgDefInfoIndex will wrap between
    // reboots. The type is interface{} with range: 1..2147483647.
    Lldpremorgdefinfoindex interface{}

    // The string value used to identify the organizationally defined information
    // of the remote system.  The encoding for this object should be as defined
    // for SnmpAdminString TC. The type is string with length: 0..507.
    Lldpremorgdefinfo interface{}
}

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetFilter() yfilter.YFilter { return lldpremorgdefinfoentry.YFilter }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) SetFilter(yf yfilter.YFilter) { lldpremorgdefinfoentry.YFilter = yf }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetGoName(yname string) string {
    if yname == "lldpRemTimeMark" { return "Lldpremtimemark" }
    if yname == "lldpRemLocalPortNum" { return "Lldpremlocalportnum" }
    if yname == "lldpRemIndex" { return "Lldpremindex" }
    if yname == "lldpRemOrgDefInfoOUI" { return "Lldpremorgdefinfooui" }
    if yname == "lldpRemOrgDefInfoSubtype" { return "Lldpremorgdefinfosubtype" }
    if yname == "lldpRemOrgDefInfoIndex" { return "Lldpremorgdefinfoindex" }
    if yname == "lldpRemOrgDefInfo" { return "Lldpremorgdefinfo" }
    return ""
}

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetSegmentPath() string {
    return "lldpRemOrgDefInfoEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremindex) + "']" + "[lldpRemOrgDefInfoOUI='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremorgdefinfooui) + "']" + "[lldpRemOrgDefInfoSubtype='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremorgdefinfosubtype) + "']" + "[lldpRemOrgDefInfoIndex='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremorgdefinfoindex) + "']"
}

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lldpRemTimeMark"] = lldpremorgdefinfoentry.Lldpremtimemark
    leafs["lldpRemLocalPortNum"] = lldpremorgdefinfoentry.Lldpremlocalportnum
    leafs["lldpRemIndex"] = lldpremorgdefinfoentry.Lldpremindex
    leafs["lldpRemOrgDefInfoOUI"] = lldpremorgdefinfoentry.Lldpremorgdefinfooui
    leafs["lldpRemOrgDefInfoSubtype"] = lldpremorgdefinfoentry.Lldpremorgdefinfosubtype
    leafs["lldpRemOrgDefInfoIndex"] = lldpremorgdefinfoentry.Lldpremorgdefinfoindex
    leafs["lldpRemOrgDefInfo"] = lldpremorgdefinfoentry.Lldpremorgdefinfo
    return leafs
}

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetBundleName() string { return "cisco_ios_xe" }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetYangName() string { return "lldpRemOrgDefInfoEntry" }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) SetParent(parent types.Entity) { lldpremorgdefinfoentry.parent = parent }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetParent() types.Entity { return lldpremorgdefinfoentry.parent }

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetParentYangName() string { return "lldpRemOrgDefInfoTable" }

