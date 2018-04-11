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

func (lLDPMIB *LLDPMIB) GetEntityData() *types.CommonEntityData {
    lLDPMIB.EntityData.YFilter = lLDPMIB.YFilter
    lLDPMIB.EntityData.YangName = "LLDP-MIB"
    lLDPMIB.EntityData.BundleName = "cisco_ios_xe"
    lLDPMIB.EntityData.ParentYangName = "LLDP-MIB"
    lLDPMIB.EntityData.SegmentPath = "LLDP-MIB:LLDP-MIB"
    lLDPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lLDPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lLDPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lLDPMIB.EntityData.Children = make(map[string]types.YChild)
    lLDPMIB.EntityData.Children["lldpConfiguration"] = types.YChild{"Lldpconfiguration", &lLDPMIB.Lldpconfiguration}
    lLDPMIB.EntityData.Children["lldpStatistics"] = types.YChild{"Lldpstatistics", &lLDPMIB.Lldpstatistics}
    lLDPMIB.EntityData.Children["lldpLocalSystemData"] = types.YChild{"Lldplocalsystemdata", &lLDPMIB.Lldplocalsystemdata}
    lLDPMIB.EntityData.Children["lldpPortConfigTable"] = types.YChild{"Lldpportconfigtable", &lLDPMIB.Lldpportconfigtable}
    lLDPMIB.EntityData.Children["lldpStatsTxPortTable"] = types.YChild{"Lldpstatstxporttable", &lLDPMIB.Lldpstatstxporttable}
    lLDPMIB.EntityData.Children["lldpStatsRxPortTable"] = types.YChild{"Lldpstatsrxporttable", &lLDPMIB.Lldpstatsrxporttable}
    lLDPMIB.EntityData.Children["lldpLocPortTable"] = types.YChild{"Lldplocporttable", &lLDPMIB.Lldplocporttable}
    lLDPMIB.EntityData.Children["lldpLocManAddrTable"] = types.YChild{"Lldplocmanaddrtable", &lLDPMIB.Lldplocmanaddrtable}
    lLDPMIB.EntityData.Children["lldpRemTable"] = types.YChild{"Lldpremtable", &lLDPMIB.Lldpremtable}
    lLDPMIB.EntityData.Children["lldpRemManAddrTable"] = types.YChild{"Lldpremmanaddrtable", &lLDPMIB.Lldpremmanaddrtable}
    lLDPMIB.EntityData.Children["lldpRemUnknownTLVTable"] = types.YChild{"Lldpremunknowntlvtable", &lLDPMIB.Lldpremunknowntlvtable}
    lLDPMIB.EntityData.Children["lldpRemOrgDefInfoTable"] = types.YChild{"Lldpremorgdefinfotable", &lLDPMIB.Lldpremorgdefinfotable}
    lLDPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lLDPMIB.EntityData)
}

// LLDPMIB_Lldpconfiguration
type LLDPMIB_Lldpconfiguration struct {
    EntityData types.CommonEntityData
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

func (lldpconfiguration *LLDPMIB_Lldpconfiguration) GetEntityData() *types.CommonEntityData {
    lldpconfiguration.EntityData.YFilter = lldpconfiguration.YFilter
    lldpconfiguration.EntityData.YangName = "lldpConfiguration"
    lldpconfiguration.EntityData.BundleName = "cisco_ios_xe"
    lldpconfiguration.EntityData.ParentYangName = "LLDP-MIB"
    lldpconfiguration.EntityData.SegmentPath = "lldpConfiguration"
    lldpconfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpconfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpconfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpconfiguration.EntityData.Children = make(map[string]types.YChild)
    lldpconfiguration.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpconfiguration.EntityData.Leafs["lldpMessageTxInterval"] = types.YLeaf{"Lldpmessagetxinterval", lldpconfiguration.Lldpmessagetxinterval}
    lldpconfiguration.EntityData.Leafs["lldpMessageTxHoldMultiplier"] = types.YLeaf{"Lldpmessagetxholdmultiplier", lldpconfiguration.Lldpmessagetxholdmultiplier}
    lldpconfiguration.EntityData.Leafs["lldpReinitDelay"] = types.YLeaf{"Lldpreinitdelay", lldpconfiguration.Lldpreinitdelay}
    lldpconfiguration.EntityData.Leafs["lldpTxDelay"] = types.YLeaf{"Lldptxdelay", lldpconfiguration.Lldptxdelay}
    lldpconfiguration.EntityData.Leafs["lldpNotificationInterval"] = types.YLeaf{"Lldpnotificationinterval", lldpconfiguration.Lldpnotificationinterval}
    return &(lldpconfiguration.EntityData)
}

// LLDPMIB_Lldpstatistics
type LLDPMIB_Lldpstatistics struct {
    EntityData types.CommonEntityData
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

func (lldpstatistics *LLDPMIB_Lldpstatistics) GetEntityData() *types.CommonEntityData {
    lldpstatistics.EntityData.YFilter = lldpstatistics.YFilter
    lldpstatistics.EntityData.YangName = "lldpStatistics"
    lldpstatistics.EntityData.BundleName = "cisco_ios_xe"
    lldpstatistics.EntityData.ParentYangName = "LLDP-MIB"
    lldpstatistics.EntityData.SegmentPath = "lldpStatistics"
    lldpstatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpstatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpstatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpstatistics.EntityData.Children = make(map[string]types.YChild)
    lldpstatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpstatistics.EntityData.Leafs["lldpStatsRemTablesLastChangeTime"] = types.YLeaf{"Lldpstatsremtableslastchangetime", lldpstatistics.Lldpstatsremtableslastchangetime}
    lldpstatistics.EntityData.Leafs["lldpStatsRemTablesInserts"] = types.YLeaf{"Lldpstatsremtablesinserts", lldpstatistics.Lldpstatsremtablesinserts}
    lldpstatistics.EntityData.Leafs["lldpStatsRemTablesDeletes"] = types.YLeaf{"Lldpstatsremtablesdeletes", lldpstatistics.Lldpstatsremtablesdeletes}
    lldpstatistics.EntityData.Leafs["lldpStatsRemTablesDrops"] = types.YLeaf{"Lldpstatsremtablesdrops", lldpstatistics.Lldpstatsremtablesdrops}
    lldpstatistics.EntityData.Leafs["lldpStatsRemTablesAgeouts"] = types.YLeaf{"Lldpstatsremtablesageouts", lldpstatistics.Lldpstatsremtablesageouts}
    return &(lldpstatistics.EntityData)
}

// LLDPMIB_Lldplocalsystemdata
type LLDPMIB_Lldplocalsystemdata struct {
    EntityData types.CommonEntityData
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

func (lldplocalsystemdata *LLDPMIB_Lldplocalsystemdata) GetEntityData() *types.CommonEntityData {
    lldplocalsystemdata.EntityData.YFilter = lldplocalsystemdata.YFilter
    lldplocalsystemdata.EntityData.YangName = "lldpLocalSystemData"
    lldplocalsystemdata.EntityData.BundleName = "cisco_ios_xe"
    lldplocalsystemdata.EntityData.ParentYangName = "LLDP-MIB"
    lldplocalsystemdata.EntityData.SegmentPath = "lldpLocalSystemData"
    lldplocalsystemdata.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldplocalsystemdata.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldplocalsystemdata.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldplocalsystemdata.EntityData.Children = make(map[string]types.YChild)
    lldplocalsystemdata.EntityData.Leafs = make(map[string]types.YLeaf)
    lldplocalsystemdata.EntityData.Leafs["lldpLocChassisIdSubtype"] = types.YLeaf{"Lldplocchassisidsubtype", lldplocalsystemdata.Lldplocchassisidsubtype}
    lldplocalsystemdata.EntityData.Leafs["lldpLocChassisId"] = types.YLeaf{"Lldplocchassisid", lldplocalsystemdata.Lldplocchassisid}
    lldplocalsystemdata.EntityData.Leafs["lldpLocSysName"] = types.YLeaf{"Lldplocsysname", lldplocalsystemdata.Lldplocsysname}
    lldplocalsystemdata.EntityData.Leafs["lldpLocSysDesc"] = types.YLeaf{"Lldplocsysdesc", lldplocalsystemdata.Lldplocsysdesc}
    lldplocalsystemdata.EntityData.Leafs["lldpLocSysCapSupported"] = types.YLeaf{"Lldplocsyscapsupported", lldplocalsystemdata.Lldplocsyscapsupported}
    lldplocalsystemdata.EntityData.Leafs["lldpLocSysCapEnabled"] = types.YLeaf{"Lldplocsyscapenabled", lldplocalsystemdata.Lldplocsyscapenabled}
    return &(lldplocalsystemdata.EntityData)
}

// LLDPMIB_Lldpportconfigtable
// The table that controls LLDP frame transmission on individual
// ports.
type LLDPMIB_Lldpportconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LLDP configuration information for a particular port. This configuration
    // parameter controls the transmission and the reception of LLDP frames on
    // those ports whose rows are created in this table. The type is slice of
    // LLDPMIB_Lldpportconfigtable_Lldpportconfigentry.
    Lldpportconfigentry []LLDPMIB_Lldpportconfigtable_Lldpportconfigentry
}

func (lldpportconfigtable *LLDPMIB_Lldpportconfigtable) GetEntityData() *types.CommonEntityData {
    lldpportconfigtable.EntityData.YFilter = lldpportconfigtable.YFilter
    lldpportconfigtable.EntityData.YangName = "lldpPortConfigTable"
    lldpportconfigtable.EntityData.BundleName = "cisco_ios_xe"
    lldpportconfigtable.EntityData.ParentYangName = "LLDP-MIB"
    lldpportconfigtable.EntityData.SegmentPath = "lldpPortConfigTable"
    lldpportconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpportconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpportconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpportconfigtable.EntityData.Children = make(map[string]types.YChild)
    lldpportconfigtable.EntityData.Children["lldpPortConfigEntry"] = types.YChild{"Lldpportconfigentry", nil}
    for i := range lldpportconfigtable.Lldpportconfigentry {
        lldpportconfigtable.EntityData.Children[types.GetSegmentPath(&lldpportconfigtable.Lldpportconfigentry[i])] = types.YChild{"Lldpportconfigentry", &lldpportconfigtable.Lldpportconfigentry[i]}
    }
    lldpportconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpportconfigtable.EntityData)
}

// LLDPMIB_Lldpportconfigtable_Lldpportconfigentry
// LLDP configuration information for a particular port.
// This configuration parameter controls the transmission and
// the reception of LLDP frames on those ports whose rows are
// created in this table.
type LLDPMIB_Lldpportconfigtable_Lldpportconfigentry struct {
    EntityData types.CommonEntityData
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

func (lldpportconfigentry *LLDPMIB_Lldpportconfigtable_Lldpportconfigentry) GetEntityData() *types.CommonEntityData {
    lldpportconfigentry.EntityData.YFilter = lldpportconfigentry.YFilter
    lldpportconfigentry.EntityData.YangName = "lldpPortConfigEntry"
    lldpportconfigentry.EntityData.BundleName = "cisco_ios_xe"
    lldpportconfigentry.EntityData.ParentYangName = "lldpPortConfigTable"
    lldpportconfigentry.EntityData.SegmentPath = "lldpPortConfigEntry" + "[lldpPortConfigPortNum='" + fmt.Sprintf("%v", lldpportconfigentry.Lldpportconfigportnum) + "']"
    lldpportconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpportconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpportconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpportconfigentry.EntityData.Children = make(map[string]types.YChild)
    lldpportconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpportconfigentry.EntityData.Leafs["lldpPortConfigPortNum"] = types.YLeaf{"Lldpportconfigportnum", lldpportconfigentry.Lldpportconfigportnum}
    lldpportconfigentry.EntityData.Leafs["lldpPortConfigAdminStatus"] = types.YLeaf{"Lldpportconfigadminstatus", lldpportconfigentry.Lldpportconfigadminstatus}
    lldpportconfigentry.EntityData.Leafs["lldpPortConfigNotificationEnable"] = types.YLeaf{"Lldpportconfignotificationenable", lldpportconfigentry.Lldpportconfignotificationenable}
    lldpportconfigentry.EntityData.Leafs["lldpPortConfigTLVsTxEnable"] = types.YLeaf{"Lldpportconfigtlvstxenable", lldpportconfigentry.Lldpportconfigtlvstxenable}
    return &(lldpportconfigentry.EntityData)
}

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
    // LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry.
    Lldpstatstxportentry []LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry
}

func (lldpstatstxporttable *LLDPMIB_Lldpstatstxporttable) GetEntityData() *types.CommonEntityData {
    lldpstatstxporttable.EntityData.YFilter = lldpstatstxporttable.YFilter
    lldpstatstxporttable.EntityData.YangName = "lldpStatsTxPortTable"
    lldpstatstxporttable.EntityData.BundleName = "cisco_ios_xe"
    lldpstatstxporttable.EntityData.ParentYangName = "LLDP-MIB"
    lldpstatstxporttable.EntityData.SegmentPath = "lldpStatsTxPortTable"
    lldpstatstxporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpstatstxporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpstatstxporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpstatstxporttable.EntityData.Children = make(map[string]types.YChild)
    lldpstatstxporttable.EntityData.Children["lldpStatsTxPortEntry"] = types.YChild{"Lldpstatstxportentry", nil}
    for i := range lldpstatstxporttable.Lldpstatstxportentry {
        lldpstatstxporttable.EntityData.Children[types.GetSegmentPath(&lldpstatstxporttable.Lldpstatstxportentry[i])] = types.YChild{"Lldpstatstxportentry", &lldpstatstxporttable.Lldpstatstxportentry[i]}
    }
    lldpstatstxporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpstatstxporttable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (lldpstatstxportentry *LLDPMIB_Lldpstatstxporttable_Lldpstatstxportentry) GetEntityData() *types.CommonEntityData {
    lldpstatstxportentry.EntityData.YFilter = lldpstatstxportentry.YFilter
    lldpstatstxportentry.EntityData.YangName = "lldpStatsTxPortEntry"
    lldpstatstxportentry.EntityData.BundleName = "cisco_ios_xe"
    lldpstatstxportentry.EntityData.ParentYangName = "lldpStatsTxPortTable"
    lldpstatstxportentry.EntityData.SegmentPath = "lldpStatsTxPortEntry" + "[lldpStatsTxPortNum='" + fmt.Sprintf("%v", lldpstatstxportentry.Lldpstatstxportnum) + "']"
    lldpstatstxportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpstatstxportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpstatstxportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpstatstxportentry.EntityData.Children = make(map[string]types.YChild)
    lldpstatstxportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpstatstxportentry.EntityData.Leafs["lldpStatsTxPortNum"] = types.YLeaf{"Lldpstatstxportnum", lldpstatstxportentry.Lldpstatstxportnum}
    lldpstatstxportentry.EntityData.Leafs["lldpStatsTxPortFramesTotal"] = types.YLeaf{"Lldpstatstxportframestotal", lldpstatstxportentry.Lldpstatstxportframestotal}
    return &(lldpstatstxportentry.EntityData)
}

// LLDPMIB_Lldpstatsrxporttable
// A table containing LLDP reception statistics for individual
// ports.  Entries are not required to exist in this table while
// the lldpPortConfigEntry object is equal to 'disabled(4)'.
type LLDPMIB_Lldpstatsrxporttable struct {
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
    // LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry.
    Lldpstatsrxportentry []LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry
}

func (lldpstatsrxporttable *LLDPMIB_Lldpstatsrxporttable) GetEntityData() *types.CommonEntityData {
    lldpstatsrxporttable.EntityData.YFilter = lldpstatsrxporttable.YFilter
    lldpstatsrxporttable.EntityData.YangName = "lldpStatsRxPortTable"
    lldpstatsrxporttable.EntityData.BundleName = "cisco_ios_xe"
    lldpstatsrxporttable.EntityData.ParentYangName = "LLDP-MIB"
    lldpstatsrxporttable.EntityData.SegmentPath = "lldpStatsRxPortTable"
    lldpstatsrxporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpstatsrxporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpstatsrxporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpstatsrxporttable.EntityData.Children = make(map[string]types.YChild)
    lldpstatsrxporttable.EntityData.Children["lldpStatsRxPortEntry"] = types.YChild{"Lldpstatsrxportentry", nil}
    for i := range lldpstatsrxporttable.Lldpstatsrxportentry {
        lldpstatsrxporttable.EntityData.Children[types.GetSegmentPath(&lldpstatsrxporttable.Lldpstatsrxportentry[i])] = types.YChild{"Lldpstatsrxportentry", &lldpstatsrxporttable.Lldpstatsrxportentry[i]}
    }
    lldpstatsrxporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpstatsrxporttable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (lldpstatsrxportentry *LLDPMIB_Lldpstatsrxporttable_Lldpstatsrxportentry) GetEntityData() *types.CommonEntityData {
    lldpstatsrxportentry.EntityData.YFilter = lldpstatsrxportentry.YFilter
    lldpstatsrxportentry.EntityData.YangName = "lldpStatsRxPortEntry"
    lldpstatsrxportentry.EntityData.BundleName = "cisco_ios_xe"
    lldpstatsrxportentry.EntityData.ParentYangName = "lldpStatsRxPortTable"
    lldpstatsrxportentry.EntityData.SegmentPath = "lldpStatsRxPortEntry" + "[lldpStatsRxPortNum='" + fmt.Sprintf("%v", lldpstatsrxportentry.Lldpstatsrxportnum) + "']"
    lldpstatsrxportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpstatsrxportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpstatsrxportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpstatsrxportentry.EntityData.Children = make(map[string]types.YChild)
    lldpstatsrxportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpstatsrxportentry.EntityData.Leafs["lldpStatsRxPortNum"] = types.YLeaf{"Lldpstatsrxportnum", lldpstatsrxportentry.Lldpstatsrxportnum}
    lldpstatsrxportentry.EntityData.Leafs["lldpStatsRxPortFramesDiscardedTotal"] = types.YLeaf{"Lldpstatsrxportframesdiscardedtotal", lldpstatsrxportentry.Lldpstatsrxportframesdiscardedtotal}
    lldpstatsrxportentry.EntityData.Leafs["lldpStatsRxPortFramesErrors"] = types.YLeaf{"Lldpstatsrxportframeserrors", lldpstatsrxportentry.Lldpstatsrxportframeserrors}
    lldpstatsrxportentry.EntityData.Leafs["lldpStatsRxPortFramesTotal"] = types.YLeaf{"Lldpstatsrxportframestotal", lldpstatsrxportentry.Lldpstatsrxportframestotal}
    lldpstatsrxportentry.EntityData.Leafs["lldpStatsRxPortTLVsDiscardedTotal"] = types.YLeaf{"Lldpstatsrxporttlvsdiscardedtotal", lldpstatsrxportentry.Lldpstatsrxporttlvsdiscardedtotal}
    lldpstatsrxportentry.EntityData.Leafs["lldpStatsRxPortTLVsUnrecognizedTotal"] = types.YLeaf{"Lldpstatsrxporttlvsunrecognizedtotal", lldpstatsrxportentry.Lldpstatsrxporttlvsunrecognizedtotal}
    lldpstatsrxportentry.EntityData.Leafs["lldpStatsRxPortAgeoutsTotal"] = types.YLeaf{"Lldpstatsrxportageoutstotal", lldpstatsrxportentry.Lldpstatsrxportageoutstotal}
    return &(lldpstatsrxportentry.EntityData)
}

// LLDPMIB_Lldplocporttable
// This table contains one or more rows per port information
// associated with the local system known to this agent.
type LLDPMIB_Lldplocporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular port component.  Entries may be created and
    // deleted in this table by the agent. The type is slice of
    // LLDPMIB_Lldplocporttable_Lldplocportentry.
    Lldplocportentry []LLDPMIB_Lldplocporttable_Lldplocportentry
}

func (lldplocporttable *LLDPMIB_Lldplocporttable) GetEntityData() *types.CommonEntityData {
    lldplocporttable.EntityData.YFilter = lldplocporttable.YFilter
    lldplocporttable.EntityData.YangName = "lldpLocPortTable"
    lldplocporttable.EntityData.BundleName = "cisco_ios_xe"
    lldplocporttable.EntityData.ParentYangName = "LLDP-MIB"
    lldplocporttable.EntityData.SegmentPath = "lldpLocPortTable"
    lldplocporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldplocporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldplocporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldplocporttable.EntityData.Children = make(map[string]types.YChild)
    lldplocporttable.EntityData.Children["lldpLocPortEntry"] = types.YChild{"Lldplocportentry", nil}
    for i := range lldplocporttable.Lldplocportentry {
        lldplocporttable.EntityData.Children[types.GetSegmentPath(&lldplocporttable.Lldplocportentry[i])] = types.YChild{"Lldplocportentry", &lldplocporttable.Lldplocportentry[i]}
    }
    lldplocporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldplocporttable.EntityData)
}

// LLDPMIB_Lldplocporttable_Lldplocportentry
// Information about a particular port component.
// 
// Entries may be created and deleted in this table by the
// agent.
type LLDPMIB_Lldplocporttable_Lldplocportentry struct {
    EntityData types.CommonEntityData
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

func (lldplocportentry *LLDPMIB_Lldplocporttable_Lldplocportentry) GetEntityData() *types.CommonEntityData {
    lldplocportentry.EntityData.YFilter = lldplocportentry.YFilter
    lldplocportentry.EntityData.YangName = "lldpLocPortEntry"
    lldplocportentry.EntityData.BundleName = "cisco_ios_xe"
    lldplocportentry.EntityData.ParentYangName = "lldpLocPortTable"
    lldplocportentry.EntityData.SegmentPath = "lldpLocPortEntry" + "[lldpLocPortNum='" + fmt.Sprintf("%v", lldplocportentry.Lldplocportnum) + "']"
    lldplocportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldplocportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldplocportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldplocportentry.EntityData.Children = make(map[string]types.YChild)
    lldplocportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldplocportentry.EntityData.Leafs["lldpLocPortNum"] = types.YLeaf{"Lldplocportnum", lldplocportentry.Lldplocportnum}
    lldplocportentry.EntityData.Leafs["lldpLocPortIdSubtype"] = types.YLeaf{"Lldplocportidsubtype", lldplocportentry.Lldplocportidsubtype}
    lldplocportentry.EntityData.Leafs["lldpLocPortId"] = types.YLeaf{"Lldplocportid", lldplocportentry.Lldplocportid}
    lldplocportentry.EntityData.Leafs["lldpLocPortDesc"] = types.YLeaf{"Lldplocportdesc", lldplocportentry.Lldplocportdesc}
    return &(lldplocportentry.EntityData)
}

// LLDPMIB_Lldplocmanaddrtable
// This table contains management address information on the
// local system known to this agent.
type LLDPMIB_Lldplocmanaddrtable struct {
    EntityData types.CommonEntityData
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

func (lldplocmanaddrtable *LLDPMIB_Lldplocmanaddrtable) GetEntityData() *types.CommonEntityData {
    lldplocmanaddrtable.EntityData.YFilter = lldplocmanaddrtable.YFilter
    lldplocmanaddrtable.EntityData.YangName = "lldpLocManAddrTable"
    lldplocmanaddrtable.EntityData.BundleName = "cisco_ios_xe"
    lldplocmanaddrtable.EntityData.ParentYangName = "LLDP-MIB"
    lldplocmanaddrtable.EntityData.SegmentPath = "lldpLocManAddrTable"
    lldplocmanaddrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldplocmanaddrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldplocmanaddrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldplocmanaddrtable.EntityData.Children = make(map[string]types.YChild)
    lldplocmanaddrtable.EntityData.Children["lldpLocManAddrEntry"] = types.YChild{"Lldplocmanaddrentry", nil}
    for i := range lldplocmanaddrtable.Lldplocmanaddrentry {
        lldplocmanaddrtable.EntityData.Children[types.GetSegmentPath(&lldplocmanaddrtable.Lldplocmanaddrentry[i])] = types.YChild{"Lldplocmanaddrentry", &lldplocmanaddrtable.Lldplocmanaddrentry[i]}
    }
    lldplocmanaddrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldplocmanaddrtable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (lldplocmanaddrentry *LLDPMIB_Lldplocmanaddrtable_Lldplocmanaddrentry) GetEntityData() *types.CommonEntityData {
    lldplocmanaddrentry.EntityData.YFilter = lldplocmanaddrentry.YFilter
    lldplocmanaddrentry.EntityData.YangName = "lldpLocManAddrEntry"
    lldplocmanaddrentry.EntityData.BundleName = "cisco_ios_xe"
    lldplocmanaddrentry.EntityData.ParentYangName = "lldpLocManAddrTable"
    lldplocmanaddrentry.EntityData.SegmentPath = "lldpLocManAddrEntry" + "[lldpLocManAddrSubtype='" + fmt.Sprintf("%v", lldplocmanaddrentry.Lldplocmanaddrsubtype) + "']" + "[lldpLocManAddr='" + fmt.Sprintf("%v", lldplocmanaddrentry.Lldplocmanaddr) + "']"
    lldplocmanaddrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldplocmanaddrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldplocmanaddrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldplocmanaddrentry.EntityData.Children = make(map[string]types.YChild)
    lldplocmanaddrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldplocmanaddrentry.EntityData.Leafs["lldpLocManAddrSubtype"] = types.YLeaf{"Lldplocmanaddrsubtype", lldplocmanaddrentry.Lldplocmanaddrsubtype}
    lldplocmanaddrentry.EntityData.Leafs["lldpLocManAddr"] = types.YLeaf{"Lldplocmanaddr", lldplocmanaddrentry.Lldplocmanaddr}
    lldplocmanaddrentry.EntityData.Leafs["lldpLocManAddrLen"] = types.YLeaf{"Lldplocmanaddrlen", lldplocmanaddrentry.Lldplocmanaddrlen}
    lldplocmanaddrentry.EntityData.Leafs["lldpLocManAddrIfSubtype"] = types.YLeaf{"Lldplocmanaddrifsubtype", lldplocmanaddrentry.Lldplocmanaddrifsubtype}
    lldplocmanaddrentry.EntityData.Leafs["lldpLocManAddrIfId"] = types.YLeaf{"Lldplocmanaddrifid", lldplocmanaddrentry.Lldplocmanaddrifid}
    lldplocmanaddrentry.EntityData.Leafs["lldpLocManAddrOID"] = types.YLeaf{"Lldplocmanaddroid", lldplocmanaddrentry.Lldplocmanaddroid}
    lldplocmanaddrentry.EntityData.Leafs["lldpConfigManAddrPortsTxEnable"] = types.YLeaf{"Lldpconfigmanaddrportstxenable", lldplocmanaddrentry.Lldpconfigmanaddrportstxenable}
    return &(lldplocmanaddrentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular physical network connection. Entries may be
    // created and deleted in this table by the agent, if a physical topology
    // discovery process is active. The type is slice of
    // LLDPMIB_Lldpremtable_Lldprementry.
    Lldprementry []LLDPMIB_Lldpremtable_Lldprementry
}

func (lldpremtable *LLDPMIB_Lldpremtable) GetEntityData() *types.CommonEntityData {
    lldpremtable.EntityData.YFilter = lldpremtable.YFilter
    lldpremtable.EntityData.YangName = "lldpRemTable"
    lldpremtable.EntityData.BundleName = "cisco_ios_xe"
    lldpremtable.EntityData.ParentYangName = "LLDP-MIB"
    lldpremtable.EntityData.SegmentPath = "lldpRemTable"
    lldpremtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpremtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpremtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpremtable.EntityData.Children = make(map[string]types.YChild)
    lldpremtable.EntityData.Children["lldpRemEntry"] = types.YChild{"Lldprementry", nil}
    for i := range lldpremtable.Lldprementry {
        lldpremtable.EntityData.Children[types.GetSegmentPath(&lldpremtable.Lldprementry[i])] = types.YChild{"Lldprementry", &lldpremtable.Lldprementry[i]}
    }
    lldpremtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpremtable.EntityData)
}

// LLDPMIB_Lldpremtable_Lldprementry
// Information about a particular physical network connection.
// Entries may be created and deleted in this table by the agent,
// if a physical topology discovery process is active.
type LLDPMIB_Lldpremtable_Lldprementry struct {
    EntityData types.CommonEntityData
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

func (lldprementry *LLDPMIB_Lldpremtable_Lldprementry) GetEntityData() *types.CommonEntityData {
    lldprementry.EntityData.YFilter = lldprementry.YFilter
    lldprementry.EntityData.YangName = "lldpRemEntry"
    lldprementry.EntityData.BundleName = "cisco_ios_xe"
    lldprementry.EntityData.ParentYangName = "lldpRemTable"
    lldprementry.EntityData.SegmentPath = "lldpRemEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldprementry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldprementry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldprementry.Lldpremindex) + "']"
    lldprementry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldprementry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldprementry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldprementry.EntityData.Children = make(map[string]types.YChild)
    lldprementry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldprementry.EntityData.Leafs["lldpRemTimeMark"] = types.YLeaf{"Lldpremtimemark", lldprementry.Lldpremtimemark}
    lldprementry.EntityData.Leafs["lldpRemLocalPortNum"] = types.YLeaf{"Lldpremlocalportnum", lldprementry.Lldpremlocalportnum}
    lldprementry.EntityData.Leafs["lldpRemIndex"] = types.YLeaf{"Lldpremindex", lldprementry.Lldpremindex}
    lldprementry.EntityData.Leafs["lldpRemChassisIdSubtype"] = types.YLeaf{"Lldpremchassisidsubtype", lldprementry.Lldpremchassisidsubtype}
    lldprementry.EntityData.Leafs["lldpRemChassisId"] = types.YLeaf{"Lldpremchassisid", lldprementry.Lldpremchassisid}
    lldprementry.EntityData.Leafs["lldpRemPortIdSubtype"] = types.YLeaf{"Lldpremportidsubtype", lldprementry.Lldpremportidsubtype}
    lldprementry.EntityData.Leafs["lldpRemPortId"] = types.YLeaf{"Lldpremportid", lldprementry.Lldpremportid}
    lldprementry.EntityData.Leafs["lldpRemPortDesc"] = types.YLeaf{"Lldpremportdesc", lldprementry.Lldpremportdesc}
    lldprementry.EntityData.Leafs["lldpRemSysName"] = types.YLeaf{"Lldpremsysname", lldprementry.Lldpremsysname}
    lldprementry.EntityData.Leafs["lldpRemSysDesc"] = types.YLeaf{"Lldpremsysdesc", lldprementry.Lldpremsysdesc}
    lldprementry.EntityData.Leafs["lldpRemSysCapSupported"] = types.YLeaf{"Lldpremsyscapsupported", lldprementry.Lldpremsyscapsupported}
    lldprementry.EntityData.Leafs["lldpRemSysCapEnabled"] = types.YLeaf{"Lldpremsyscapenabled", lldprementry.Lldpremsyscapenabled}
    return &(lldprementry.EntityData)
}

// LLDPMIB_Lldpremmanaddrtable
// This table contains one or more rows per management address
// information on the remote system learned on a particular port
// contained in the local chassis known to this agent.
type LLDPMIB_Lldpremmanaddrtable struct {
    EntityData types.CommonEntityData
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

func (lldpremmanaddrtable *LLDPMIB_Lldpremmanaddrtable) GetEntityData() *types.CommonEntityData {
    lldpremmanaddrtable.EntityData.YFilter = lldpremmanaddrtable.YFilter
    lldpremmanaddrtable.EntityData.YangName = "lldpRemManAddrTable"
    lldpremmanaddrtable.EntityData.BundleName = "cisco_ios_xe"
    lldpremmanaddrtable.EntityData.ParentYangName = "LLDP-MIB"
    lldpremmanaddrtable.EntityData.SegmentPath = "lldpRemManAddrTable"
    lldpremmanaddrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpremmanaddrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpremmanaddrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpremmanaddrtable.EntityData.Children = make(map[string]types.YChild)
    lldpremmanaddrtable.EntityData.Children["lldpRemManAddrEntry"] = types.YChild{"Lldpremmanaddrentry", nil}
    for i := range lldpremmanaddrtable.Lldpremmanaddrentry {
        lldpremmanaddrtable.EntityData.Children[types.GetSegmentPath(&lldpremmanaddrtable.Lldpremmanaddrentry[i])] = types.YChild{"Lldpremmanaddrentry", &lldpremmanaddrtable.Lldpremmanaddrentry[i]}
    }
    lldpremmanaddrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpremmanaddrtable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Lldpremmanaddroid interface{}
}

func (lldpremmanaddrentry *LLDPMIB_Lldpremmanaddrtable_Lldpremmanaddrentry) GetEntityData() *types.CommonEntityData {
    lldpremmanaddrentry.EntityData.YFilter = lldpremmanaddrentry.YFilter
    lldpremmanaddrentry.EntityData.YangName = "lldpRemManAddrEntry"
    lldpremmanaddrentry.EntityData.BundleName = "cisco_ios_xe"
    lldpremmanaddrentry.EntityData.ParentYangName = "lldpRemManAddrTable"
    lldpremmanaddrentry.EntityData.SegmentPath = "lldpRemManAddrEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremindex) + "']" + "[lldpRemManAddrSubtype='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremmanaddrsubtype) + "']" + "[lldpRemManAddr='" + fmt.Sprintf("%v", lldpremmanaddrentry.Lldpremmanaddr) + "']"
    lldpremmanaddrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpremmanaddrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpremmanaddrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpremmanaddrentry.EntityData.Children = make(map[string]types.YChild)
    lldpremmanaddrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpremmanaddrentry.EntityData.Leafs["lldpRemTimeMark"] = types.YLeaf{"Lldpremtimemark", lldpremmanaddrentry.Lldpremtimemark}
    lldpremmanaddrentry.EntityData.Leafs["lldpRemLocalPortNum"] = types.YLeaf{"Lldpremlocalportnum", lldpremmanaddrentry.Lldpremlocalportnum}
    lldpremmanaddrentry.EntityData.Leafs["lldpRemIndex"] = types.YLeaf{"Lldpremindex", lldpremmanaddrentry.Lldpremindex}
    lldpremmanaddrentry.EntityData.Leafs["lldpRemManAddrSubtype"] = types.YLeaf{"Lldpremmanaddrsubtype", lldpremmanaddrentry.Lldpremmanaddrsubtype}
    lldpremmanaddrentry.EntityData.Leafs["lldpRemManAddr"] = types.YLeaf{"Lldpremmanaddr", lldpremmanaddrentry.Lldpremmanaddr}
    lldpremmanaddrentry.EntityData.Leafs["lldpRemManAddrIfSubtype"] = types.YLeaf{"Lldpremmanaddrifsubtype", lldpremmanaddrentry.Lldpremmanaddrifsubtype}
    lldpremmanaddrentry.EntityData.Leafs["lldpRemManAddrIfId"] = types.YLeaf{"Lldpremmanaddrifid", lldpremmanaddrentry.Lldpremmanaddrifid}
    lldpremmanaddrentry.EntityData.Leafs["lldpRemManAddrOID"] = types.YLeaf{"Lldpremmanaddroid", lldpremmanaddrentry.Lldpremmanaddroid}
    return &(lldpremmanaddrentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an unrecognized TLV received from a physical network
    // connection.  Entries may be created and deleted in this table by the agent,
    // if a physical topology discovery process is active. The type is slice of
    // LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry.
    Lldpremunknowntlventry []LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry
}

func (lldpremunknowntlvtable *LLDPMIB_Lldpremunknowntlvtable) GetEntityData() *types.CommonEntityData {
    lldpremunknowntlvtable.EntityData.YFilter = lldpremunknowntlvtable.YFilter
    lldpremunknowntlvtable.EntityData.YangName = "lldpRemUnknownTLVTable"
    lldpremunknowntlvtable.EntityData.BundleName = "cisco_ios_xe"
    lldpremunknowntlvtable.EntityData.ParentYangName = "LLDP-MIB"
    lldpremunknowntlvtable.EntityData.SegmentPath = "lldpRemUnknownTLVTable"
    lldpremunknowntlvtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpremunknowntlvtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpremunknowntlvtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpremunknowntlvtable.EntityData.Children = make(map[string]types.YChild)
    lldpremunknowntlvtable.EntityData.Children["lldpRemUnknownTLVEntry"] = types.YChild{"Lldpremunknowntlventry", nil}
    for i := range lldpremunknowntlvtable.Lldpremunknowntlventry {
        lldpremunknowntlvtable.EntityData.Children[types.GetSegmentPath(&lldpremunknowntlvtable.Lldpremunknowntlventry[i])] = types.YChild{"Lldpremunknowntlventry", &lldpremunknowntlvtable.Lldpremunknowntlventry[i]}
    }
    lldpremunknowntlvtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpremunknowntlvtable.EntityData)
}

// LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry
// Information about an unrecognized TLV received from a
// physical network connection.  Entries may be created and
// deleted in this table by the agent, if a physical topology
// discovery process is active.
type LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry struct {
    EntityData types.CommonEntityData
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

func (lldpremunknowntlventry *LLDPMIB_Lldpremunknowntlvtable_Lldpremunknowntlventry) GetEntityData() *types.CommonEntityData {
    lldpremunknowntlventry.EntityData.YFilter = lldpremunknowntlventry.YFilter
    lldpremunknowntlventry.EntityData.YangName = "lldpRemUnknownTLVEntry"
    lldpremunknowntlventry.EntityData.BundleName = "cisco_ios_xe"
    lldpremunknowntlventry.EntityData.ParentYangName = "lldpRemUnknownTLVTable"
    lldpremunknowntlventry.EntityData.SegmentPath = "lldpRemUnknownTLVEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremindex) + "']" + "[lldpRemUnknownTLVType='" + fmt.Sprintf("%v", lldpremunknowntlventry.Lldpremunknowntlvtype) + "']"
    lldpremunknowntlventry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpremunknowntlventry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpremunknowntlventry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpremunknowntlventry.EntityData.Children = make(map[string]types.YChild)
    lldpremunknowntlventry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpremunknowntlventry.EntityData.Leafs["lldpRemTimeMark"] = types.YLeaf{"Lldpremtimemark", lldpremunknowntlventry.Lldpremtimemark}
    lldpremunknowntlventry.EntityData.Leafs["lldpRemLocalPortNum"] = types.YLeaf{"Lldpremlocalportnum", lldpremunknowntlventry.Lldpremlocalportnum}
    lldpremunknowntlventry.EntityData.Leafs["lldpRemIndex"] = types.YLeaf{"Lldpremindex", lldpremunknowntlventry.Lldpremindex}
    lldpremunknowntlventry.EntityData.Leafs["lldpRemUnknownTLVType"] = types.YLeaf{"Lldpremunknowntlvtype", lldpremunknowntlventry.Lldpremunknowntlvtype}
    lldpremunknowntlventry.EntityData.Leafs["lldpRemUnknownTLVInfo"] = types.YLeaf{"Lldpremunknowntlvinfo", lldpremunknowntlventry.Lldpremunknowntlvinfo}
    return &(lldpremunknowntlventry.EntityData)
}

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
    // LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry.
    Lldpremorgdefinfoentry []LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry
}

func (lldpremorgdefinfotable *LLDPMIB_Lldpremorgdefinfotable) GetEntityData() *types.CommonEntityData {
    lldpremorgdefinfotable.EntityData.YFilter = lldpremorgdefinfotable.YFilter
    lldpremorgdefinfotable.EntityData.YangName = "lldpRemOrgDefInfoTable"
    lldpremorgdefinfotable.EntityData.BundleName = "cisco_ios_xe"
    lldpremorgdefinfotable.EntityData.ParentYangName = "LLDP-MIB"
    lldpremorgdefinfotable.EntityData.SegmentPath = "lldpRemOrgDefInfoTable"
    lldpremorgdefinfotable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpremorgdefinfotable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpremorgdefinfotable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpremorgdefinfotable.EntityData.Children = make(map[string]types.YChild)
    lldpremorgdefinfotable.EntityData.Children["lldpRemOrgDefInfoEntry"] = types.YChild{"Lldpremorgdefinfoentry", nil}
    for i := range lldpremorgdefinfotable.Lldpremorgdefinfoentry {
        lldpremorgdefinfotable.EntityData.Children[types.GetSegmentPath(&lldpremorgdefinfotable.Lldpremorgdefinfoentry[i])] = types.YChild{"Lldpremorgdefinfoentry", &lldpremorgdefinfotable.Lldpremorgdefinfoentry[i]}
    }
    lldpremorgdefinfotable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lldpremorgdefinfotable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (lldpremorgdefinfoentry *LLDPMIB_Lldpremorgdefinfotable_Lldpremorgdefinfoentry) GetEntityData() *types.CommonEntityData {
    lldpremorgdefinfoentry.EntityData.YFilter = lldpremorgdefinfoentry.YFilter
    lldpremorgdefinfoentry.EntityData.YangName = "lldpRemOrgDefInfoEntry"
    lldpremorgdefinfoentry.EntityData.BundleName = "cisco_ios_xe"
    lldpremorgdefinfoentry.EntityData.ParentYangName = "lldpRemOrgDefInfoTable"
    lldpremorgdefinfoentry.EntityData.SegmentPath = "lldpRemOrgDefInfoEntry" + "[lldpRemTimeMark='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremtimemark) + "']" + "[lldpRemLocalPortNum='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremlocalportnum) + "']" + "[lldpRemIndex='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremindex) + "']" + "[lldpRemOrgDefInfoOUI='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremorgdefinfooui) + "']" + "[lldpRemOrgDefInfoSubtype='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremorgdefinfosubtype) + "']" + "[lldpRemOrgDefInfoIndex='" + fmt.Sprintf("%v", lldpremorgdefinfoentry.Lldpremorgdefinfoindex) + "']"
    lldpremorgdefinfoentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lldpremorgdefinfoentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lldpremorgdefinfoentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lldpremorgdefinfoentry.EntityData.Children = make(map[string]types.YChild)
    lldpremorgdefinfoentry.EntityData.Leafs = make(map[string]types.YLeaf)
    lldpremorgdefinfoentry.EntityData.Leafs["lldpRemTimeMark"] = types.YLeaf{"Lldpremtimemark", lldpremorgdefinfoentry.Lldpremtimemark}
    lldpremorgdefinfoentry.EntityData.Leafs["lldpRemLocalPortNum"] = types.YLeaf{"Lldpremlocalportnum", lldpremorgdefinfoentry.Lldpremlocalportnum}
    lldpremorgdefinfoentry.EntityData.Leafs["lldpRemIndex"] = types.YLeaf{"Lldpremindex", lldpremorgdefinfoentry.Lldpremindex}
    lldpremorgdefinfoentry.EntityData.Leafs["lldpRemOrgDefInfoOUI"] = types.YLeaf{"Lldpremorgdefinfooui", lldpremorgdefinfoentry.Lldpremorgdefinfooui}
    lldpremorgdefinfoentry.EntityData.Leafs["lldpRemOrgDefInfoSubtype"] = types.YLeaf{"Lldpremorgdefinfosubtype", lldpremorgdefinfoentry.Lldpremorgdefinfosubtype}
    lldpremorgdefinfoentry.EntityData.Leafs["lldpRemOrgDefInfoIndex"] = types.YLeaf{"Lldpremorgdefinfoindex", lldpremorgdefinfoentry.Lldpremorgdefinfoindex}
    lldpremorgdefinfoentry.EntityData.Leafs["lldpRemOrgDefInfo"] = types.YLeaf{"Lldpremorgdefinfo", lldpremorgdefinfoentry.Lldpremorgdefinfo}
    return &(lldpremorgdefinfoentry.EntityData)
}

