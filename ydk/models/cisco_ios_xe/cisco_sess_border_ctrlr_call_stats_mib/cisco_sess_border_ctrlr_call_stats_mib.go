// The main purpose of this MIB is to define the statistics
// information for Session Border Controller application. The 
// statistics are mainly of two types - Call statistics and 
// Media statistics. The calls can further be categorized as SIP 
// calls and H.248 calls.
// 
// This MIB categorizes the statistics information into following
// four types:
// 1. Global call statistics - Represents the global call related
//    statistics like call rates, media flows, signaling flows
//    etc.
// 
// 2. Periodic statistics - Represents the SBC call statistics 
//    information for a particular time interval like current 5 
//    minutes, previous 5 minutes, current 15 minutes, previous 15
//    minutes, current hour and previous hour.
// 
// 3. Per flow statistics - Represents the SBC media flow 
//    statistics. These are media statistics for each of the
//    current ongoing call flow.
// 
// 4. H.248 statistics - Represents the H.248 call related 
//    statistics information when H.248 controller is associated
//    with SBC.
// 
// The Session Border Controller (SBC) enables direct IP-to-IP
// interconnect between multiple administrative domains for
// session-based services providing protocol inter-working,
// security, and admission control and management. The SBC is a
// voice over IP (VoIP) device that sits on the border of a 
// network and controls call admission to that network. 
// 
// The primary purpose of an SBC is to protect the interior of the
// network from excessive call load and malicious traffic.
// Additional functions provided by the SBC include media bridging
// and billing services. 
// 
// 
// GLOSSARY
// SBC: Session Border Controller
// 
// CSB: CISCO Session Border Controller
// 
// CAC: Call Admission Control - protects voice traffic from the
// negative effects of other voice traffic and to keep excess
// voice traffic off the network. It is used to prevent congestion
// in Voice traffic. It is used in the Call Setup phase.
// 
// RTP: Real Time Transport Protocol - defines a standardized
// packet format for delivering audio and video over the Internet.
// 
// RTCP: Real Time Control Protocol - It is a sister protocol of
// the Real-time Transport Protocol (RTP). RTCP provides
// out-of-band control information for an RTP flow. It partners
// RTP in the delivery and packaging of multimedia data, but does
// not transport any data itself. It is used periodically to
// transmit control packets to participants in a streaming
// multimedia session.
// 
// VMG: Virtual Media Gateway - introduced to bridge between
// different transmission technologies and to add service to
// end-user connections. Its architecture separates control and
// connectivity functions into physically separate network layers.
// 
// VPN: Virtual Private Network - It is a communications network
// tunneled through another network, and dedicated for a specific
// network.
// 
// Gate Id - Context Identifiers assigned uniquely to a SIP/H.248
// call flows.
// 
// Flow Pair Id: Stream Identifiers assigned uniquely to a
// SIP/H.248 call flows.
// 
// Adjacency: An adjacency contains the system information to be
// transmitted to next HOP.
// 
// REFERENCES
// 1. CISCO Session Border Controller Documents and FAQ
// http://zed.cisco.com/confluence/display/SBC/SBC
package cisco_sess_border_ctrlr_call_stats_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_sess_border_ctrlr_call_stats_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB}", reflect.TypeOf(CISCOSESSBORDERCTRLRCALLSTATSMIB{}))
    ydk.RegisterEntity("CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB:CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB", reflect.TypeOf(CISCOSESSBORDERCTRLRCALLSTATSMIB{}))
}

// CiscoSbcPeriodicStatsInterval represents statistics information.
type CiscoSbcPeriodicStatsInterval string

const (
    CiscoSbcPeriodicStatsInterval_fiveMinute CiscoSbcPeriodicStatsInterval = "fiveMinute"

    CiscoSbcPeriodicStatsInterval_fifteenMinute CiscoSbcPeriodicStatsInterval = "fifteenMinute"

    CiscoSbcPeriodicStatsInterval_oneHour CiscoSbcPeriodicStatsInterval = "oneHour"

    CiscoSbcPeriodicStatsInterval_oneDay CiscoSbcPeriodicStatsInterval = "oneDay"
)

// CISCOSESSBORDERCTRLRCALLSTATSMIB
type CISCOSESSBORDERCTRLRCALLSTATSMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The call stats instance table contains the physical index for each of the
    // physical entity (line card, primary, secondary cards). The index of the
    // table is instance index which uniquely identifies the physical entity
    // present on the box.
    Csbcallstatsinstancetable CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable

    // This table describes the global statistics information in the form of a
    // table which contains call specific information like call rates, media
    // flows, signaling flows etc. The index of the table is service index which
    // corresponds to a particular  service configured on the SBC and all the rows
    // of the table represents the global information regarding all the call flows
    // related to that particular service. The other index of this table is
    // csbCallStatsInstanceIndex which is defined in csbCallStatsInstanceTable.
    Csbcallstatstable CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable

    // This table is used to collect measurement over several different intervals
    // as defined by the csbCurrPeriodicStatsInterval object. When a new interval
    // starts the objects associated with the interval are reset and count up
    // throughout the interval. The index of the table is the interval for which
    // the stats information is to be displayed. The interval values can be 5 min,
    // 15 mins, 1 hour and one day. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable.  The gauge values
    // are reported as :- 1.If the period being queried is current5mins, this is
    // the value at the instant that the query is issued.  2.Otherwise, for the
    // other intevals, this is an average value during the summary period sampled
    // at 5 minute intervals.
    Csbcurrperiodicstatstable CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable

    // This table provide historical measurement in various interval length
    // defined by the csbHistoryStatsInterval object. Each interval may contain
    // one or more entries to allow for detailed measurment to be collected. It is
    // up to the platform to determine the number of intervals to be supported
    // like  5 minutes, 15 minutes, 1 hour and 1 day. In addition, the number of
    // historical entries is also determined by the platform resources.  The gauge
    // values are reported as: If the period being queried is previous5mins, this
    // is the number of calls that were active at the end of the previous 5 minute
    // period. Otherwise for the other intevals, this is an average value during
    // the summary period, sampled at 5 minute intervals.
    Csbhistorystatstable CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable

    // This table describes statistics table for each call flow. The indices of
    // the table are virtual media gateway id, gate id, falow pair id and side id
    // (indices for side A or side B). The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable.
    Csbperflowstatstable CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable

    // This table describes the H.248 statistics for SBC. The index of the table
    // is service index which corresponds to a particular  service configured on
    // the SBC and the index assigned to a particular H.248 controller. The other
    // index of this table is csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable. This table is replaced by the
    // csbH248StatsRev1Table.
    Csbh248Statstable CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable

    // This table describes the H.248 statistics for SBC. The index of the table
    // is service index which corresponds to a particular  service configured on
    // the SBC and the index assigned to a particular H.248 controller. The other
    // index of this table is csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable.
    Csbh248Statsrev1Table CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table
}

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetEntityData() *types.CommonEntityData {
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.YFilter = cISCOSESSBORDERCTRLRCALLSTATSMIB.YFilter
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.YangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.SegmentPath = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB:CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children["csbCallStatsInstanceTable"] = types.YChild{"Csbcallstatsinstancetable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcallstatsinstancetable}
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children["csbCallStatsTable"] = types.YChild{"Csbcallstatstable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcallstatstable}
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children["csbCurrPeriodicStatsTable"] = types.YChild{"Csbcurrperiodicstatstable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcurrperiodicstatstable}
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children["csbHistoryStatsTable"] = types.YChild{"Csbhistorystatstable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbhistorystatstable}
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children["csbPerFlowStatsTable"] = types.YChild{"Csbperflowstatstable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbperflowstatstable}
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children["csbH248StatsTable"] = types.YChild{"Csbh248Statstable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbh248Statstable}
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children["csbH248StatsRev1Table"] = types.YChild{"Csbh248Statsrev1Table", &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbh248Statsrev1Table}
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable
// The call stats instance table contains the physical index for
// each of the physical entity (line card, primary, secondary
// cards). The index of the table is instance index which uniquely
// identifies the physical entity present on the box.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in csbCallStatsInstanceTable. There is an entry in this
    // table for each SBC instance, as identified by a  value of
    // csbCallStatsInstanceIndex. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry.
    Csbcallstatsinstanceentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry
}

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetEntityData() *types.CommonEntityData {
    csbcallstatsinstancetable.EntityData.YFilter = csbcallstatsinstancetable.YFilter
    csbcallstatsinstancetable.EntityData.YangName = "csbCallStatsInstanceTable"
    csbcallstatsinstancetable.EntityData.BundleName = "cisco_ios_xe"
    csbcallstatsinstancetable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbcallstatsinstancetable.EntityData.SegmentPath = "csbCallStatsInstanceTable"
    csbcallstatsinstancetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbcallstatsinstancetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbcallstatsinstancetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbcallstatsinstancetable.EntityData.Children = make(map[string]types.YChild)
    csbcallstatsinstancetable.EntityData.Children["csbCallStatsInstanceEntry"] = types.YChild{"Csbcallstatsinstanceentry", nil}
    for i := range csbcallstatsinstancetable.Csbcallstatsinstanceentry {
        csbcallstatsinstancetable.EntityData.Children[types.GetSegmentPath(&csbcallstatsinstancetable.Csbcallstatsinstanceentry[i])] = types.YChild{"Csbcallstatsinstanceentry", &csbcallstatsinstancetable.Csbcallstatsinstanceentry[i]}
    }
    csbcallstatsinstancetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbcallstatsinstancetable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry
// A conceptual row in csbCallStatsInstanceTable. There is an
// entry in this table for each SBC instance, as identified by a 
// value of csbCallStatsInstanceIndex.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object uniquely identifies the sequence
    // number of an entity or slot that is configured per device. This index is
    // assigned arbitrarily by the engine and is not saved over reboots. The type
    // is interface{} with range: 0..4294967295.
    Csbcallstatsinstanceindex interface{}

    // This object indicates the physical entity for which all the measurements
    // are maintained. The exact type of this entity is described by its
    // entPhysicalVendorType value. The type is interface{} with range:
    // 0..2147483647.
    Csbcallstatsinstancephysicalindex interface{}
}

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetEntityData() *types.CommonEntityData {
    csbcallstatsinstanceentry.EntityData.YFilter = csbcallstatsinstanceentry.YFilter
    csbcallstatsinstanceentry.EntityData.YangName = "csbCallStatsInstanceEntry"
    csbcallstatsinstanceentry.EntityData.BundleName = "cisco_ios_xe"
    csbcallstatsinstanceentry.EntityData.ParentYangName = "csbCallStatsInstanceTable"
    csbcallstatsinstanceentry.EntityData.SegmentPath = "csbCallStatsInstanceEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbcallstatsinstanceentry.Csbcallstatsinstanceindex) + "']"
    csbcallstatsinstanceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbcallstatsinstanceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbcallstatsinstanceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbcallstatsinstanceentry.EntityData.Children = make(map[string]types.YChild)
    csbcallstatsinstanceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbcallstatsinstanceentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbcallstatsinstanceentry.Csbcallstatsinstanceindex}
    csbcallstatsinstanceentry.EntityData.Leafs["csbCallStatsInstancePhysicalIndex"] = types.YLeaf{"Csbcallstatsinstancephysicalindex", csbcallstatsinstanceentry.Csbcallstatsinstancephysicalindex}
    return &(csbcallstatsinstanceentry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable
// This table describes the global statistics information in the
// form of a table which contains call specific information like
// call rates, media flows, signaling flows etc. The index of the
// table is service index which corresponds to a particular 
// service configured on the SBC and all the rows of the table
// represents the global information regarding all the call flows
// related to that particular service. The other index of this
// table is csbCallStatsInstanceIndex which is defined in
// csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStatsGlobalStatsTable. There is an entry in
    // this table for the particular service configured on SBC identified by a
    // value of csbCallStatsInstanceIndex. The other index of this table is
    // csbCallStatsInstanceIndex which is defined in csbCallStatsInstanceTable.
    // The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry.
    Csbcallstatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry
}

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetEntityData() *types.CommonEntityData {
    csbcallstatstable.EntityData.YFilter = csbcallstatstable.YFilter
    csbcallstatstable.EntityData.YangName = "csbCallStatsTable"
    csbcallstatstable.EntityData.BundleName = "cisco_ios_xe"
    csbcallstatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbcallstatstable.EntityData.SegmentPath = "csbCallStatsTable"
    csbcallstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbcallstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbcallstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbcallstatstable.EntityData.Children = make(map[string]types.YChild)
    csbcallstatstable.EntityData.Children["csbCallStatsEntry"] = types.YChild{"Csbcallstatsentry", nil}
    for i := range csbcallstatstable.Csbcallstatsentry {
        csbcallstatstable.EntityData.Children[types.GetSegmentPath(&csbcallstatstable.Csbcallstatsentry[i])] = types.YChild{"Csbcallstatsentry", &csbcallstatstable.Csbcallstatsentry[i]}
    }
    csbcallstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbcallstatstable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry
// An conceptual row in the csbCallStatsGlobalStatsTable. There is
// an entry in this table for the particular service configured on
// SBC identified by a value of csbCallStatsInstanceIndex. The
// other index of this table is csbCallStatsInstanceIndex which is
// defined in csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry_Csbcallstatsinstanceindex
    Csbcallstatsinstanceindex interface{}

    // This attribute is a key. This object identifies the index of the name of
    // the SBC service configured. This object also acts as an index for the
    // table. The type is interface{} with range: 0..4294967295.
    Csbcallstatsserviceindex interface{}

    // This object indicates the configured name of the SBC service. The length of
    // this object is zero when value is not assigned to it. The type is string.
    Csbcallstatssbcname interface{}

    // This object indicates the maximum number of calls per second processed by
    // the Session Border Controller in past 24 hours. The type is interface{}
    // with range: 0..4294967295. Units are calls per second.
    Csbcallstatscallshigh interface{}

    // This object indicates the average calls per second processed by the Session
    // Border Controller. The type is interface{} with range: 0..4294967295. Units
    // are calls per second.
    Csbcallstatsrate1Sec interface{}

    // This object indicates the minimum calls per second in past 24 hours. The
    // type is interface{} with range: 0..4294967295. Units are calls per second.
    Csbcallstatscallslow interface{}

    // This object indicates the number of media flows which are available. The
    // type is interface{} with range: 0..4294967295. Units are flows.
    Csbcallstatsavailableflows interface{}

    // This object indicates the number of media flows which are used. This is for
    // the flow allocated and connected. The flow allocated but not connected is
    // not counted. The type is interface{} with range: 0..4294967295. Units are
    // flows.
    Csbcallstatsusedflows interface{}

    // This object indicates the number of peak flows in SBC. This is the highest
    // recorded value for the active flows since the box was booted/reseted. The
    // type is interface{} with range: 0..4294967295. Units are flows.
    Csbcallstatspeakflows interface{}

    // This object indicates the total number of media support by this instance of
    // SBC. The total number of flows include the available flows and the active
    // flows. This value is since box was booted/reseted. Total flows include the
    // active flows and the flows allocated but not connected. The type is
    // interface{} with range: 0..4294967295. Units are flows.
    Csbcallstatstotalflows interface{}

    // This object indicates the number of active signaling flows for signaling
    // pinholes. The type is interface{} with range: 0..4294967295. Units are
    // signal flows.
    Csbcallstatsusedsigflows interface{}

    // This object indicates the peak signaling flow in SBC. This is the highest
    // recorded value for the active signaling flows. This object is calculated
    // using csbCallStatsUsedFlows. The type is interface{} with range:
    // 0..4294967295. Units are signal flows.
    Csbcallstatspeaksigflows interface{}

    // This object indicates the maximum number of Signalling Flows that can be
    // supported by this instance of SBC. The type is interface{} with range:
    // 0..4294967295. Units are signal flows.
    Csbcallstatstotalsigflows interface{}

    // This object indicates the remaining capacity that can be supported by SBC.
    // The type is interface{} with range: 0..4294967295. Units are packets per
    // second.
    Csbcallstatsavailablepktrate interface{}

    // This object indicates the number of unclassified packets processed by SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Csbcallstatsunclassifiedpkts interface{}

    // This object indicates the total number of RTP packets sent. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    Csbcallstatsrtppktssent interface{}

    // This object indicates the total number of RTP packets received. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    Csbcallstatsrtppktsrcvd interface{}

    // This object indicates the total number of RTP packets discarded. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    Csbcallstatsrtppktsdiscard interface{}

    // This object indicates the number of RTP octets sent by the SBC. The type is
    // interface{} with range: 0..18446744073709551615. Units are octets.
    Csbcallstatsrtpoctetssent interface{}

    // This object indicates the number of RTP octets received by the SBC. The
    // type is interface{} with range: 0..18446744073709551615. Units are octets.
    Csbcallstatsrtpoctetsrcvd interface{}

    // This object indicates the number of RTP octets discarded by the SBC. The
    // type is interface{} with range: 0..18446744073709551615. Units are octets.
    Csbcallstatsrtpoctetsdiscard interface{}

    // This object indicates the accumulated No media event count. The type is
    // interface{} with range: 0..4294967295. Units are no media events.
    Csbcallstatsnomediacount interface{}

    // This object indicates the accumulated route error event count. This counter
    // is for the route error of media stream. The type is interface{} with range:
    // 0..4294967295. Units are route error events.
    Csbcallstatsrouteerrors interface{}

    // This object indicates the number of additional transcoded flows that this
    // media gateway manager (MGM) entity is currently able to configure. The type
    // is interface{} with range: 0..4294967295. Units are flows.
    Csbcallstatsavailabletranscodeflows interface{}

    // This object indicates the current number of transcoded flows that are
    // actively forwarding media traffic.  In this context, a flow is a media
    // stream passing through the device. So a single voice call will be counted
    // only once. The type is interface{} with range: 0..4294967295. Units are
    // flows.
    Csbcallstatsactivetranscodeflows interface{}

    // This object indicates the peak number of active transcoded flows since the
    // statistics were last reset.  In this context, a flow is a media stream
    // passing through the device, so a single voice call will be counted once.
    // The type is interface{} with range: 0..4294967295. Units are flows.
    Csbcallstatspeaktranscodeflows interface{}

    // This object indicates the accumulated total number of transcoded flows. 
    // This count contains both active flows and flows that were allocated but
    // never connected.  In this context, a flow is a media stream passing through
    // the device, so a single voice call will be counted once. The type is
    // interface{} with range: 0..4294967295. Units are flows.
    Csbcallstatstotaltranscodeflows interface{}
}

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetEntityData() *types.CommonEntityData {
    csbcallstatsentry.EntityData.YFilter = csbcallstatsentry.YFilter
    csbcallstatsentry.EntityData.YangName = "csbCallStatsEntry"
    csbcallstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbcallstatsentry.EntityData.ParentYangName = "csbCallStatsTable"
    csbcallstatsentry.EntityData.SegmentPath = "csbCallStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbcallstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbcallstatsentry.Csbcallstatsserviceindex) + "']"
    csbcallstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbcallstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbcallstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbcallstatsentry.EntityData.Children = make(map[string]types.YChild)
    csbcallstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbcallstatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbcallstatsentry.Csbcallstatsinstanceindex}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbcallstatsentry.Csbcallstatsserviceindex}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsSbcName"] = types.YLeaf{"Csbcallstatssbcname", csbcallstatsentry.Csbcallstatssbcname}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsCallsHigh"] = types.YLeaf{"Csbcallstatscallshigh", csbcallstatsentry.Csbcallstatscallshigh}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRate1Sec"] = types.YLeaf{"Csbcallstatsrate1Sec", csbcallstatsentry.Csbcallstatsrate1Sec}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsCallsLow"] = types.YLeaf{"Csbcallstatscallslow", csbcallstatsentry.Csbcallstatscallslow}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsAvailableFlows"] = types.YLeaf{"Csbcallstatsavailableflows", csbcallstatsentry.Csbcallstatsavailableflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsUsedFlows"] = types.YLeaf{"Csbcallstatsusedflows", csbcallstatsentry.Csbcallstatsusedflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsPeakFlows"] = types.YLeaf{"Csbcallstatspeakflows", csbcallstatsentry.Csbcallstatspeakflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsTotalFlows"] = types.YLeaf{"Csbcallstatstotalflows", csbcallstatsentry.Csbcallstatstotalflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsUsedSigFlows"] = types.YLeaf{"Csbcallstatsusedsigflows", csbcallstatsentry.Csbcallstatsusedsigflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsPeakSigFlows"] = types.YLeaf{"Csbcallstatspeaksigflows", csbcallstatsentry.Csbcallstatspeaksigflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsTotalSigFlows"] = types.YLeaf{"Csbcallstatstotalsigflows", csbcallstatsentry.Csbcallstatstotalsigflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsAvailablePktRate"] = types.YLeaf{"Csbcallstatsavailablepktrate", csbcallstatsentry.Csbcallstatsavailablepktrate}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsUnclassifiedPkts"] = types.YLeaf{"Csbcallstatsunclassifiedpkts", csbcallstatsentry.Csbcallstatsunclassifiedpkts}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRTPPktsSent"] = types.YLeaf{"Csbcallstatsrtppktssent", csbcallstatsentry.Csbcallstatsrtppktssent}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRTPPktsRcvd"] = types.YLeaf{"Csbcallstatsrtppktsrcvd", csbcallstatsentry.Csbcallstatsrtppktsrcvd}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRTPPktsDiscard"] = types.YLeaf{"Csbcallstatsrtppktsdiscard", csbcallstatsentry.Csbcallstatsrtppktsdiscard}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRTPOctetsSent"] = types.YLeaf{"Csbcallstatsrtpoctetssent", csbcallstatsentry.Csbcallstatsrtpoctetssent}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRTPOctetsRcvd"] = types.YLeaf{"Csbcallstatsrtpoctetsrcvd", csbcallstatsentry.Csbcallstatsrtpoctetsrcvd}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRTPOctetsDiscard"] = types.YLeaf{"Csbcallstatsrtpoctetsdiscard", csbcallstatsentry.Csbcallstatsrtpoctetsdiscard}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsNoMediaCount"] = types.YLeaf{"Csbcallstatsnomediacount", csbcallstatsentry.Csbcallstatsnomediacount}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsRouteErrors"] = types.YLeaf{"Csbcallstatsrouteerrors", csbcallstatsentry.Csbcallstatsrouteerrors}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsAvailableTranscodeFlows"] = types.YLeaf{"Csbcallstatsavailabletranscodeflows", csbcallstatsentry.Csbcallstatsavailabletranscodeflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsActiveTranscodeFlows"] = types.YLeaf{"Csbcallstatsactivetranscodeflows", csbcallstatsentry.Csbcallstatsactivetranscodeflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsPeakTranscodeFlows"] = types.YLeaf{"Csbcallstatspeaktranscodeflows", csbcallstatsentry.Csbcallstatspeaktranscodeflows}
    csbcallstatsentry.EntityData.Leafs["csbCallStatsTotalTranscodeFlows"] = types.YLeaf{"Csbcallstatstotaltranscodeflows", csbcallstatsentry.Csbcallstatstotaltranscodeflows}
    return &(csbcallstatsentry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable
// This table is used to collect measurement over several
// different intervals as defined by the
// csbCurrPeriodicStatsInterval object. When a new interval starts
// the objects associated with the interval are reset and count up
// throughout the interval. The index of the table is the interval
// for which the stats information is to be displayed. The interval
// values can be 5 min, 15 mins, 1 hour and one day. The other
// indices of this table are csbCallStatsInstanceIndex
// defined in csbCallStatsInstanceTable and
// csbCallStatsServiceIndex defined in csbCallStatsTable.
// 
// The gauge values are reported as :-
// 1.If the period being queried is current5mins, this is the value
// at the instant that the query is issued. 
// 2.Otherwise, for the other intevals, this is an average value
// during the summary period sampled at 5 minute intervals.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCurrPeriodicStatsTable. There is an entry in
    // this table for the particular controller by a value of
    // csbH248StatsCtrlrIndex. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry.
    Csbcurrperiodicstatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry
}

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetEntityData() *types.CommonEntityData {
    csbcurrperiodicstatstable.EntityData.YFilter = csbcurrperiodicstatstable.YFilter
    csbcurrperiodicstatstable.EntityData.YangName = "csbCurrPeriodicStatsTable"
    csbcurrperiodicstatstable.EntityData.BundleName = "cisco_ios_xe"
    csbcurrperiodicstatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbcurrperiodicstatstable.EntityData.SegmentPath = "csbCurrPeriodicStatsTable"
    csbcurrperiodicstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbcurrperiodicstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbcurrperiodicstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbcurrperiodicstatstable.EntityData.Children = make(map[string]types.YChild)
    csbcurrperiodicstatstable.EntityData.Children["csbCurrPeriodicStatsEntry"] = types.YChild{"Csbcurrperiodicstatsentry", nil}
    for i := range csbcurrperiodicstatstable.Csbcurrperiodicstatsentry {
        csbcurrperiodicstatstable.EntityData.Children[types.GetSegmentPath(&csbcurrperiodicstatstable.Csbcurrperiodicstatsentry[i])] = types.YChild{"Csbcurrperiodicstatsentry", &csbcurrperiodicstatstable.Csbcurrperiodicstatsentry[i]}
    }
    csbcurrperiodicstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbcurrperiodicstatstable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry
// An conceptual row in the csbCurrPeriodicStatsTable. There is
// an entry in this table for the particular controller by a value
// of csbH248StatsCtrlrIndex. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry_Csbcallstatsinstanceindex
    Csbcallstatsinstanceindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry_Csbcallstatsserviceindex
    Csbcallstatsserviceindex interface{}

    // This attribute is a key. This object identifies the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be 5 min, 15 mins, 1 hour , 1 Day. This object acts as index for the table.
    // The type is CiscoSbcPeriodicStatsInterval.
    Csbcurrperiodicstatsinterval interface{}

    // This object indicates the number of calls that have become active during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbcurrperiodicstatsactivecalls interface{}

    // This object indicates the number of calls that have become activing during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbcurrperiodicstatsactivatingcalls interface{}

    // This object indicates the number of calls that have become deactiving
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    Csbcurrperiodicstatsdeactivatingcalls interface{}

    // This object indicates the number of total call attempts during this
    // interval. The type is interface{} with range: 0..4294967295.
    Csbcurrperiodicstatstotalcallattempts interface{}

    // This object indicates the number of failed call attempts during this
    // interval. The type is interface{} with range: 0..4294967295.
    Csbcurrperiodicstatsfailedcallattempts interface{}

    // This object indicates the number of call setup failures due to routing
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallroutingfailure interface{}

    // This object indicates the number of call setup failures due to resource
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallresourcefailure interface{}

    // This object indicates the number of call setup failures due to media
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallmediafailure interface{}

    // This object indicates the number of call setup failures due to signaling
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsigfailure interface{}

    // This object indicates the number of active call failures during this
    // interval. The type is interface{} with range: 0..4294967295.
    Csbcurrperiodicstatsactivecallfailure interface{}

    // This object indicates the number of call setup failures due to congestion
    // during this interval. The type is interface{} with range: 0..4294967295.
    Csbcurrperiodicstatscongestionfailure interface{}

    // This object indicates the number of call setup failures due to policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetuppolicyfailure interface{}

    // This object indicates the number of call setup failures due to NA policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetupnapolicyfailure interface{}

    // This object indicates the number of call setup failures due to routing
    // policy failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetuproutingpolicyfailure interface{}

    // This object indicates the number of call setup failures due to CAC policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetupcacpolicyfailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetupcaccalllimitfailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // rate limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetupcacratelimitfailure interface{}

    // This object indicates the number of call setup failures due to CAC
    // bandwidth limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetupcacbandwidthfailure interface{}

    // This object indicates the number of call setup failures due to CAC media
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetupcacmedialimitfailure interface{}

    // This object indicates the number of call update failure due to policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbcurrperiodicstatscallsetupcacmediaupdatefailure interface{}

    // This object indicates the current time at the start of each interval. The
    // type is string with length: 0..80.
    Csbcurrperiodicstatstimestamp interface{}

    // The object indicates the number of transcoded calls that are active during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbcurrperiodicstatstranscodedcalls interface{}

    // The object indicates the number of transrated calls that are active during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbcurrperiodicstatstransratedcalls interface{}

    // This object indicates the total number of call update failures during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbcurrperiodicstatstotalcallupdatefailure interface{}

    // This Object indicates the number of calls through SBC that use IPv6
    // signaling.  This statistic totals all calls that traverse an IPv6 adjacency
    // on either or both sides of SBC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatsactiveipv6Calls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis) during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatsactiveemergencycalls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis) and have used the e2
    // interface to obtain location information for the caller during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbcurrperiodicstatsactivee2Emergencycalls interface{}

    // This object indicates the total number of active calls which use IMS Rx,
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    Csbcurrperiodicstatsimsrxactivecalls interface{}

    // This object indicates the total call Setup failures owing to IMS Rx failure
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are failures.
    Csbcurrperiodicstatsimsrxcallsetupfaiures interface{}

    // This object indicates the total call renegotiation attempts using IMS Rx
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are attempts.
    Csbcurrperiodicstatsimsrxcallrenegotiationattempts interface{}

    // This object indicates the total call renegotiation failures owing to IMS Rx
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295. Units are failures.
    Csbcurrperiodicstatsimsrxcallrenegotiationfailures interface{}

    // The number of active audio transcoded calls through this adjacency or
    // account during this interval. The type is interface{} with range:
    // 0..4294967295. Units are calls.
    Csbcurrperiodicstatsaudiotranscodedcalls interface{}

    // This object indicates the the number of active fax transcoded calls through
    // this adjacency or account during this interval. The type is interface{}
    // with range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatsfaxtranscodedcalls interface{}

    // This object indicates the total call setup failures due to RTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    Csbcurrperiodicstatsrtpdisallowedfailures interface{}

    // This object indicates the total call setup failures due to SRTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    Csbcurrperiodicstatssrtpdisallowedfailures interface{}

    // This object indicates the number of active calls through this adjacency or
    // account which do not use SRTP on any media channels during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatsnonsrtpcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels which use SRTP. This count
    // does not include media  channels that provide interworking between RTP and
    // SRTP during this interval. The type is interface{} with range:
    // 0..4294967295. Units are calls.
    Csbcurrperiodicstatssrtpnoniwcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels that provide interworking
    // between RTP and SRTP during this interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatssrtpiwcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via RFC 2833 during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatsdtmfiw2833Calls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via  inband DTMF tones during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatsdtmfiwinbandcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in media via
    // RFC 2833 and DTMF in media via inband DTMF tones during this interval. The
    // type is interface{} with range: 0..4294967295. Units are calls.
    Csbcurrperiodicstatsdtmfiw2833Inbandcalls interface{}

    // This object indicates the lawful intercept tap attempts requested within
    // the scope of this query during this interval. The type is interface{} with
    // range: 0..4294967295. Units are attempts.
    Csbcurrperiodicstatstotaltapsrequested interface{}

    // This object indicates the lawful intercept tap attempts that have been
    // successfully implemented within the scope of this query during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // success.
    Csbcurrperiodicstatstotaltapssucceeded interface{}

    // This object indicates the Lawful intercept taps currently in place on calls
    // within the scope of this query during this interval. The type is
    // interface{} with range: 0..4294967295. Units are taps.
    Csbcurrperiodicstatscurrenttaps interface{}

    // The number of active calls on this adjacency or account which are to or
    // from registered subscribers using IPSEC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbcurrperiodicipseccalls interface{}
}

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetEntityData() *types.CommonEntityData {
    csbcurrperiodicstatsentry.EntityData.YFilter = csbcurrperiodicstatsentry.YFilter
    csbcurrperiodicstatsentry.EntityData.YangName = "csbCurrPeriodicStatsEntry"
    csbcurrperiodicstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbcurrperiodicstatsentry.EntityData.ParentYangName = "csbCurrPeriodicStatsTable"
    csbcurrperiodicstatsentry.EntityData.SegmentPath = "csbCurrPeriodicStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbcurrperiodicstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbcurrperiodicstatsentry.Csbcallstatsserviceindex) + "']" + "[csbCurrPeriodicStatsInterval='" + fmt.Sprintf("%v", csbcurrperiodicstatsentry.Csbcurrperiodicstatsinterval) + "']"
    csbcurrperiodicstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbcurrperiodicstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbcurrperiodicstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbcurrperiodicstatsentry.EntityData.Children = make(map[string]types.YChild)
    csbcurrperiodicstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbcurrperiodicstatsentry.Csbcallstatsinstanceindex}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbcurrperiodicstatsentry.Csbcallstatsserviceindex}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsInterval"] = types.YLeaf{"Csbcurrperiodicstatsinterval", csbcurrperiodicstatsentry.Csbcurrperiodicstatsinterval}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsActiveCalls"] = types.YLeaf{"Csbcurrperiodicstatsactivecalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivecalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsActivatingCalls"] = types.YLeaf{"Csbcurrperiodicstatsactivatingcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivatingcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsDeactivatingCalls"] = types.YLeaf{"Csbcurrperiodicstatsdeactivatingcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsdeactivatingcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsTotalCallAttempts"] = types.YLeaf{"Csbcurrperiodicstatstotalcallattempts", csbcurrperiodicstatsentry.Csbcurrperiodicstatstotalcallattempts}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsFailedCallAttempts"] = types.YLeaf{"Csbcurrperiodicstatsfailedcallattempts", csbcurrperiodicstatsentry.Csbcurrperiodicstatsfailedcallattempts}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallRoutingFailure"] = types.YLeaf{"Csbcurrperiodicstatscallroutingfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallroutingfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallResourceFailure"] = types.YLeaf{"Csbcurrperiodicstatscallresourcefailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallresourcefailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallMediaFailure"] = types.YLeaf{"Csbcurrperiodicstatscallmediafailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallmediafailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSigFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsigfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsigfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsActiveCallFailure"] = types.YLeaf{"Csbcurrperiodicstatsactivecallfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivecallfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCongestionFailure"] = types.YLeaf{"Csbcurrperiodicstatscongestionfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscongestionfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupPolicyFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetuppolicyfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetuppolicyfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupNAPolicyFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetupnapolicyfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupnapolicyfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupRoutingPolicyFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetuproutingpolicyfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetuproutingpolicyfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupCACPolicyFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetupcacpolicyfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacpolicyfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupCACCallLimitFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetupcaccalllimitfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcaccalllimitfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupCACRateLimitFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetupcacratelimitfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacratelimitfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupCACBandwidthFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetupcacbandwidthfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacbandwidthfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupCACMediaLimitFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetupcacmedialimitfailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacmedialimitfailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCallSetupCACMediaUpdateFailure"] = types.YLeaf{"Csbcurrperiodicstatscallsetupcacmediaupdatefailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacmediaupdatefailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsTimestamp"] = types.YLeaf{"Csbcurrperiodicstatstimestamp", csbcurrperiodicstatsentry.Csbcurrperiodicstatstimestamp}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsTranscodedCalls"] = types.YLeaf{"Csbcurrperiodicstatstranscodedcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatstranscodedcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsTransratedCalls"] = types.YLeaf{"Csbcurrperiodicstatstransratedcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatstransratedcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsTotalCallUpdateFailure"] = types.YLeaf{"Csbcurrperiodicstatstotalcallupdatefailure", csbcurrperiodicstatsentry.Csbcurrperiodicstatstotalcallupdatefailure}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsActiveIpv6Calls"] = types.YLeaf{"Csbcurrperiodicstatsactiveipv6Calls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsactiveipv6Calls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsActiveEmergencyCalls"] = types.YLeaf{"Csbcurrperiodicstatsactiveemergencycalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsactiveemergencycalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsActiveE2EmergencyCalls"] = types.YLeaf{"Csbcurrperiodicstatsactivee2Emergencycalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivee2Emergencycalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsImsRxActiveCalls"] = types.YLeaf{"Csbcurrperiodicstatsimsrxactivecalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxactivecalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsImsRxCallSetupFaiures"] = types.YLeaf{"Csbcurrperiodicstatsimsrxcallsetupfaiures", csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxcallsetupfaiures}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsImsRxCallRenegotiationAttempts"] = types.YLeaf{"Csbcurrperiodicstatsimsrxcallrenegotiationattempts", csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxcallrenegotiationattempts}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsImsRxCallRenegotiationFailures"] = types.YLeaf{"Csbcurrperiodicstatsimsrxcallrenegotiationfailures", csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxcallrenegotiationfailures}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsAudioTranscodedCalls"] = types.YLeaf{"Csbcurrperiodicstatsaudiotranscodedcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsaudiotranscodedcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsFaxTranscodedCalls"] = types.YLeaf{"Csbcurrperiodicstatsfaxtranscodedcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsfaxtranscodedcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsRtpDisallowedFailures"] = types.YLeaf{"Csbcurrperiodicstatsrtpdisallowedfailures", csbcurrperiodicstatsentry.Csbcurrperiodicstatsrtpdisallowedfailures}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsSrtpDisallowedFailures"] = types.YLeaf{"Csbcurrperiodicstatssrtpdisallowedfailures", csbcurrperiodicstatsentry.Csbcurrperiodicstatssrtpdisallowedfailures}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsNonSrtpCalls"] = types.YLeaf{"Csbcurrperiodicstatsnonsrtpcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsnonsrtpcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsSrtpNonIwCalls"] = types.YLeaf{"Csbcurrperiodicstatssrtpnoniwcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatssrtpnoniwcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsSrtpIwCalls"] = types.YLeaf{"Csbcurrperiodicstatssrtpiwcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatssrtpiwcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsDtmfIw2833Calls"] = types.YLeaf{"Csbcurrperiodicstatsdtmfiw2833Calls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsdtmfiw2833Calls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsDtmfIwInbandCalls"] = types.YLeaf{"Csbcurrperiodicstatsdtmfiwinbandcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsdtmfiwinbandcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsDtmfIw2833InbandCalls"] = types.YLeaf{"Csbcurrperiodicstatsdtmfiw2833Inbandcalls", csbcurrperiodicstatsentry.Csbcurrperiodicstatsdtmfiw2833Inbandcalls}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsTotalTapsRequested"] = types.YLeaf{"Csbcurrperiodicstatstotaltapsrequested", csbcurrperiodicstatsentry.Csbcurrperiodicstatstotaltapsrequested}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsTotalTapsSucceeded"] = types.YLeaf{"Csbcurrperiodicstatstotaltapssucceeded", csbcurrperiodicstatsentry.Csbcurrperiodicstatstotaltapssucceeded}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicStatsCurrentTaps"] = types.YLeaf{"Csbcurrperiodicstatscurrenttaps", csbcurrperiodicstatsentry.Csbcurrperiodicstatscurrenttaps}
    csbcurrperiodicstatsentry.EntityData.Leafs["csbCurrPeriodicIpsecCalls"] = types.YLeaf{"Csbcurrperiodicipseccalls", csbcurrperiodicstatsentry.Csbcurrperiodicipseccalls}
    return &(csbcurrperiodicstatsentry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable
// This table provide historical measurement in various interval
// length defined by the csbHistoryStatsInterval object. Each
// interval may contain one or more entries to allow for detailed
// measurment to be collected. It is up to the platform to
// determine the number of intervals to be supported like 
// 5 minutes, 15 minutes, 1 hour and 1 day. In addition, the number
// of
// historical entries is also determined by the platform
// resources.
// 
// The gauge values are reported as:
// If the period being queried is previous5mins, this is the number
// of calls that were active at the end of the previous 5 minute
// period. Otherwise for the other intevals, this is an average
// value during the summary period, sampled at 5 minute intervals.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbHistoryStatsTable. The entries in this table are
    // updated as interval completes in the csbCurrPeriodicStatsTable table and
    // the data is moved from that table to this one. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry.
    Csbhistorystatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry
}

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetEntityData() *types.CommonEntityData {
    csbhistorystatstable.EntityData.YFilter = csbhistorystatstable.YFilter
    csbhistorystatstable.EntityData.YangName = "csbHistoryStatsTable"
    csbhistorystatstable.EntityData.BundleName = "cisco_ios_xe"
    csbhistorystatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbhistorystatstable.EntityData.SegmentPath = "csbHistoryStatsTable"
    csbhistorystatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbhistorystatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbhistorystatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbhistorystatstable.EntityData.Children = make(map[string]types.YChild)
    csbhistorystatstable.EntityData.Children["csbHistoryStatsEntry"] = types.YChild{"Csbhistorystatsentry", nil}
    for i := range csbhistorystatstable.Csbhistorystatsentry {
        csbhistorystatstable.EntityData.Children[types.GetSegmentPath(&csbhistorystatstable.Csbhistorystatsentry[i])] = types.YChild{"Csbhistorystatsentry", &csbhistorystatstable.Csbhistorystatsentry[i]}
    }
    csbhistorystatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbhistorystatstable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry
// A conceptual row in the csbHistoryStatsTable. The entries in
// this table are updated as interval completes in the
// csbCurrPeriodicStatsTable table and the data is moved from that
// table to this one.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry_Csbcallstatsinstanceindex
    Csbcallstatsinstanceindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry_Csbcallstatsserviceindex
    Csbcallstatsserviceindex interface{}

    // This attribute is a key. This object identifies the interval for which the
    // history statistics information is to be displayed. The interval values can
    // be 5 min, 15 mins, 1 hour , 1 day. This object acts as index for the table.
    // The type is CiscoSbcPeriodicStatsInterval.
    Csbhistorystatsinterval interface{}

    // This attribute is a key. The platform assigns a number starting with one
    // and increments it each for each new row. When the maximum          number
    // of row is reached the oldest rows are deleted. It is up to the platform to
    // determine the number of entries for each interval. It is recommended that
    // each platform support at least one entry for 5 min, 15 mins, 1 hour and 1
    // day intervals. The type is interface{} with range: 0..4294967295.
    Csbhistorystatselements interface{}

    // This object indicates the number of active calls history during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbhistorystatsactivecalls interface{}

    // This object indicates the number of total call attempts history during this
    // interval. The type is interface{} with range: 0..4294967295.
    Csbhistorystatstotalcallattempts interface{}

    // This object indicates the number of failed call attempts during this
    // interval. The type is interface{} with range: 0..4294967295.
    Csbhistorystatsfailedcallattempts interface{}

    // This object indicates the number of call setup failures due to routing
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallroutingfailure interface{}

    // This object indicates the number of call setup failures due to resource
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallresourcefailure interface{}

    // This object indicates the number of call setup failures due to media
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallmediafailure interface{}

    // This object indicates the number of call setup failures due to signaling
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatsfailsigfailure interface{}

    // This object indicates the number of active call failures during this
    // interval. The type is interface{} with range: 0..4294967295.
    Csbhistorystatsactivecallfailure interface{}

    // This object indicates the number of call setup failures due to congestion
    // during this interval. The type is interface{} with range: 0..4294967295.
    Csbhistorystatscongestionfailure interface{}

    // This object indicates the number of call setup failures due to some policy
    // violations during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetuppolicyfailure interface{}

    // This object indicates the number of call setup failures due to NA policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetupnapolicyfailure interface{}

    // This object indicates the number of call setup failures due to routing
    // policy failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetuproutingpolicyfailure interface{}

    // This object indicates the number of call setup failures due to CAC policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetupcacpolicyfailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetupcaccalllimitfailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // rate limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetupcacratelimitfailure interface{}

    // This object indicates the number of call setup failures due to CAC
    // bandwidth limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetupcacbandwidthfailure interface{}

    // This object indicates the number of call setup failures due to CAC media
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetupcacmedialimitfailure interface{}

    // This object indicates the number of call update failure due to policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    Csbhistorystatscallsetupcacmediaupdatefailure interface{}

    // This object indicates the time at the start of the interval when
    // measurements were first collected for this interval in the
    // csbCurrPeriodicStatsTable. The type is string with length: 0..80.
    Csbhistorystatstimestamp interface{}

    // The object indicates the number of active transcoded calls during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbhistroystatstranscodedcalls interface{}

    // The object indicates the number of active transrated calls during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbhistroystatstransratedcalls interface{}

    // This object indicates the total call update failures during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatstotalcallupdatefailure interface{}

    // This Object indicates the number of calls through SBC that use IPv6
    // signaling.  This statistic totals all calls that traverse an IPv6 adjacency
    // on either or both sides of SBC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsactiveipv6Calls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis)  during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsactiveemergencycalls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis) and have used the e2
    // interface to obtain location information for the caller. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsactivee2Emergencycalls interface{}

    // This object indicates the total number of active calls which use IMS Rx,
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    Csbhistorystatsimsrxactivecalls interface{}

    // This object indicates the total call setup failures owing to IMS Rx failure
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are failures.
    Csbhistorystatsimsrxcallsetupfailures interface{}

    // This object indicates the total call renegotiation attempts using IMS Rx
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are attempts.
    Csbhistorystatsimsrxcallrenegotiationattempts interface{}

    // This object indicates the total call renegotiation failures owing to IMS Rx
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295. Units are failures.
    Csbhistorystatsimsrxcallrenegotiationfailures interface{}

    // The number of active audio transcoded calls through this adjacency or
    // account during this interval. The type is interface{} with range:
    // 0..4294967295. Units are calls.
    Csbhistorystatsaudiotranscodedcalls interface{}

    // This object indicates the the number of active fax transcoded calls through
    // this adjacency or account during this interval. The type is interface{}
    // with range: 0..4294967295. Units are calls.
    Csbhistorystatsfaxtranscodedcalls interface{}

    // This object indicates the total call setup failures due to RTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    Csbhistorystatsrtpdisallowedfailures interface{}

    // This object indicates the total call setup failures due to SRTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    Csbhistorystatssrtpdisallowedfailures interface{}

    // This object indicates the number of active calls through this adjacency or
    // account which do not use SRTP on any media channels during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsnonsrtpcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels that use SRTP but no media
    // channels that provide interworking between RTP and SRTP during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    Csbhistorystatssrtpnoniwcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels that provide interworking
    // between RTP and SRTP during this interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    Csbhistorystatssrtpiwcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via RFC 2833 during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsdtmfiw2833Calls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via inband DTMF tones during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsdtmfiwinbandcalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in media via
    // RFC 2833 and DTMF in media via inband DTMF tones during this interval. The
    // type is interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsdtmfiw2833Inbandcalls interface{}

    // This object indicates the lawful intercept tap attempts requested within
    // the scope of this query during this interval. The type is interface{} with
    // range: 0..4294967295. Units are attempts.
    Csbhistorystatstotaltapsrequested interface{}

    // This object indicates the lawful intercept tap attempts that have been
    // successfully implemented within the scope of this query during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // success.
    Csbhistorystatstotaltapssucceeded interface{}

    // This object indicates the Lawful intercept taps currently in place on calls
    // within the scope of this query during this interval. The type is
    // interface{} with range: 0..4294967295. Units are taps.
    Csbhistorystatscurrenttaps interface{}

    // The number of active calls on this adjacency or account which are to or
    // from registered subscribers using IPSEC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    Csbhistorystatsipseccalls interface{}
}

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetEntityData() *types.CommonEntityData {
    csbhistorystatsentry.EntityData.YFilter = csbhistorystatsentry.YFilter
    csbhistorystatsentry.EntityData.YangName = "csbHistoryStatsEntry"
    csbhistorystatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbhistorystatsentry.EntityData.ParentYangName = "csbHistoryStatsTable"
    csbhistorystatsentry.EntityData.SegmentPath = "csbHistoryStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbcallstatsserviceindex) + "']" + "[csbHistoryStatsInterval='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbhistorystatsinterval) + "']" + "[csbHistoryStatsElements='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbhistorystatselements) + "']"
    csbhistorystatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbhistorystatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbhistorystatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbhistorystatsentry.EntityData.Children = make(map[string]types.YChild)
    csbhistorystatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbhistorystatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbhistorystatsentry.Csbcallstatsinstanceindex}
    csbhistorystatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbhistorystatsentry.Csbcallstatsserviceindex}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsInterval"] = types.YLeaf{"Csbhistorystatsinterval", csbhistorystatsentry.Csbhistorystatsinterval}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsElements"] = types.YLeaf{"Csbhistorystatselements", csbhistorystatsentry.Csbhistorystatselements}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsActiveCalls"] = types.YLeaf{"Csbhistorystatsactivecalls", csbhistorystatsentry.Csbhistorystatsactivecalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsTotalCallAttempts"] = types.YLeaf{"Csbhistorystatstotalcallattempts", csbhistorystatsentry.Csbhistorystatstotalcallattempts}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsFailedCallAttempts"] = types.YLeaf{"Csbhistorystatsfailedcallattempts", csbhistorystatsentry.Csbhistorystatsfailedcallattempts}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallRoutingFailure"] = types.YLeaf{"Csbhistorystatscallroutingfailure", csbhistorystatsentry.Csbhistorystatscallroutingfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallResourceFailure"] = types.YLeaf{"Csbhistorystatscallresourcefailure", csbhistorystatsentry.Csbhistorystatscallresourcefailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallMediaFailure"] = types.YLeaf{"Csbhistorystatscallmediafailure", csbhistorystatsentry.Csbhistorystatscallmediafailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsFailSigFailure"] = types.YLeaf{"Csbhistorystatsfailsigfailure", csbhistorystatsentry.Csbhistorystatsfailsigfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsActiveCallFailure"] = types.YLeaf{"Csbhistorystatsactivecallfailure", csbhistorystatsentry.Csbhistorystatsactivecallfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCongestionFailure"] = types.YLeaf{"Csbhistorystatscongestionfailure", csbhistorystatsentry.Csbhistorystatscongestionfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupPolicyFailure"] = types.YLeaf{"Csbhistorystatscallsetuppolicyfailure", csbhistorystatsentry.Csbhistorystatscallsetuppolicyfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupNAPolicyFailure"] = types.YLeaf{"Csbhistorystatscallsetupnapolicyfailure", csbhistorystatsentry.Csbhistorystatscallsetupnapolicyfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupRoutingPolicyFailure"] = types.YLeaf{"Csbhistorystatscallsetuproutingpolicyfailure", csbhistorystatsentry.Csbhistorystatscallsetuproutingpolicyfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupCACPolicyFailure"] = types.YLeaf{"Csbhistorystatscallsetupcacpolicyfailure", csbhistorystatsentry.Csbhistorystatscallsetupcacpolicyfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupCACCallLimitFailure"] = types.YLeaf{"Csbhistorystatscallsetupcaccalllimitfailure", csbhistorystatsentry.Csbhistorystatscallsetupcaccalllimitfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupCACRateLimitFailure"] = types.YLeaf{"Csbhistorystatscallsetupcacratelimitfailure", csbhistorystatsentry.Csbhistorystatscallsetupcacratelimitfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupCACBandwidthFailure"] = types.YLeaf{"Csbhistorystatscallsetupcacbandwidthfailure", csbhistorystatsentry.Csbhistorystatscallsetupcacbandwidthfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupCACMediaLimitFailure"] = types.YLeaf{"Csbhistorystatscallsetupcacmedialimitfailure", csbhistorystatsentry.Csbhistorystatscallsetupcacmedialimitfailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCallSetupCACMediaUpdateFailure"] = types.YLeaf{"Csbhistorystatscallsetupcacmediaupdatefailure", csbhistorystatsentry.Csbhistorystatscallsetupcacmediaupdatefailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsTimestamp"] = types.YLeaf{"Csbhistorystatstimestamp", csbhistorystatsentry.Csbhistorystatstimestamp}
    csbhistorystatsentry.EntityData.Leafs["csbHistroyStatsTranscodedCalls"] = types.YLeaf{"Csbhistroystatstranscodedcalls", csbhistorystatsentry.Csbhistroystatstranscodedcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistroyStatsTransratedCalls"] = types.YLeaf{"Csbhistroystatstransratedcalls", csbhistorystatsentry.Csbhistroystatstransratedcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsTotalCallUpdateFailure"] = types.YLeaf{"Csbhistorystatstotalcallupdatefailure", csbhistorystatsentry.Csbhistorystatstotalcallupdatefailure}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsActiveIpv6Calls"] = types.YLeaf{"Csbhistorystatsactiveipv6Calls", csbhistorystatsentry.Csbhistorystatsactiveipv6Calls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsActiveEmergencyCalls"] = types.YLeaf{"Csbhistorystatsactiveemergencycalls", csbhistorystatsentry.Csbhistorystatsactiveemergencycalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsActiveE2EmergencyCalls"] = types.YLeaf{"Csbhistorystatsactivee2Emergencycalls", csbhistorystatsentry.Csbhistorystatsactivee2Emergencycalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsImsRxActiveCalls"] = types.YLeaf{"Csbhistorystatsimsrxactivecalls", csbhistorystatsentry.Csbhistorystatsimsrxactivecalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsImsRxCallSetupFailures"] = types.YLeaf{"Csbhistorystatsimsrxcallsetupfailures", csbhistorystatsentry.Csbhistorystatsimsrxcallsetupfailures}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsImsRxCallRenegotiationAttempts"] = types.YLeaf{"Csbhistorystatsimsrxcallrenegotiationattempts", csbhistorystatsentry.Csbhistorystatsimsrxcallrenegotiationattempts}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsImsRxCallRenegotiationFailures"] = types.YLeaf{"Csbhistorystatsimsrxcallrenegotiationfailures", csbhistorystatsentry.Csbhistorystatsimsrxcallrenegotiationfailures}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsAudioTranscodedCalls"] = types.YLeaf{"Csbhistorystatsaudiotranscodedcalls", csbhistorystatsentry.Csbhistorystatsaudiotranscodedcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsFaxTranscodedCalls"] = types.YLeaf{"Csbhistorystatsfaxtranscodedcalls", csbhistorystatsentry.Csbhistorystatsfaxtranscodedcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsRtpDisallowedFailures"] = types.YLeaf{"Csbhistorystatsrtpdisallowedfailures", csbhistorystatsentry.Csbhistorystatsrtpdisallowedfailures}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsSrtpDisallowedFailures"] = types.YLeaf{"Csbhistorystatssrtpdisallowedfailures", csbhistorystatsentry.Csbhistorystatssrtpdisallowedfailures}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsNonSrtpCalls"] = types.YLeaf{"Csbhistorystatsnonsrtpcalls", csbhistorystatsentry.Csbhistorystatsnonsrtpcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsSrtpNonIwCalls"] = types.YLeaf{"Csbhistorystatssrtpnoniwcalls", csbhistorystatsentry.Csbhistorystatssrtpnoniwcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsSrtpIwCalls"] = types.YLeaf{"Csbhistorystatssrtpiwcalls", csbhistorystatsentry.Csbhistorystatssrtpiwcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsDtmfIw2833Calls"] = types.YLeaf{"Csbhistorystatsdtmfiw2833Calls", csbhistorystatsentry.Csbhistorystatsdtmfiw2833Calls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsDtmfIwInbandCalls"] = types.YLeaf{"Csbhistorystatsdtmfiwinbandcalls", csbhistorystatsentry.Csbhistorystatsdtmfiwinbandcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsDtmfIw2833InbandCalls"] = types.YLeaf{"Csbhistorystatsdtmfiw2833Inbandcalls", csbhistorystatsentry.Csbhistorystatsdtmfiw2833Inbandcalls}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsTotalTapsRequested"] = types.YLeaf{"Csbhistorystatstotaltapsrequested", csbhistorystatsentry.Csbhistorystatstotaltapsrequested}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsTotalTapsSucceeded"] = types.YLeaf{"Csbhistorystatstotaltapssucceeded", csbhistorystatsentry.Csbhistorystatstotaltapssucceeded}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsCurrentTaps"] = types.YLeaf{"Csbhistorystatscurrenttaps", csbhistorystatsentry.Csbhistorystatscurrenttaps}
    csbhistorystatsentry.EntityData.Leafs["csbHistoryStatsIpsecCalls"] = types.YLeaf{"Csbhistorystatsipseccalls", csbhistorystatsentry.Csbhistorystatsipseccalls}
    return &(csbhistorystatsentry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable
// This table describes statistics table for each call flow. The
// indices of the table are virtual media gateway id, gate id,
// falow pair id and side id (indices for side A or side B). The
// other indices of this table are csbCallStatsInstanceIndex
// defined in csbCallStatsInstanceTable and
// csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbPerFlowStatsTable. There is an entry in this
    // table for vdbe Id, gate id, flow pair id and side id. The other indices of
    // this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry.
    Csbperflowstatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry
}

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetEntityData() *types.CommonEntityData {
    csbperflowstatstable.EntityData.YFilter = csbperflowstatstable.YFilter
    csbperflowstatstable.EntityData.YangName = "csbPerFlowStatsTable"
    csbperflowstatstable.EntityData.BundleName = "cisco_ios_xe"
    csbperflowstatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbperflowstatstable.EntityData.SegmentPath = "csbPerFlowStatsTable"
    csbperflowstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbperflowstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbperflowstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbperflowstatstable.EntityData.Children = make(map[string]types.YChild)
    csbperflowstatstable.EntityData.Children["csbPerFlowStatsEntry"] = types.YChild{"Csbperflowstatsentry", nil}
    for i := range csbperflowstatstable.Csbperflowstatsentry {
        csbperflowstatstable.EntityData.Children[types.GetSegmentPath(&csbperflowstatstable.Csbperflowstatsentry[i])] = types.YChild{"Csbperflowstatsentry", &csbperflowstatstable.Csbperflowstatsentry[i]}
    }
    csbperflowstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbperflowstatstable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry
// An conceptual row in the csbPerFlowStatsTable. There is
// an entry in this table for vdbe Id, gate id, flow pair id and
// side id. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry_Csbcallstatsinstanceindex
    Csbcallstatsinstanceindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry_Csbcallstatsserviceindex
    Csbcallstatsserviceindex interface{}

    // This attribute is a key. This object identifies the virtual media gateway
    // id. This object also acts as an index for the table. The type is
    // interface{} with range: 0..255.
    Csbperflowstatsvdbeid interface{}

    // This attribute is a key. This object identifies the gate id. This object
    // also acts as an index for the table. The type is interface{} with range:
    // 0..65535.
    Csbperflowstatsgateid interface{}

    // This attribute is a key. This object identifies the flow pair id. This
    // object also acts as an index for the table. The type is interface{} with
    // range: 0..65535.
    Csbperflowstatsflowpairid interface{}

    // This attribute is a key. This object identifies the index corresponding to
    // side of flow pair either side A or side B. This object also acts as an
    // index for the table. The type is Csbperflowstatssideid.
    Csbperflowstatssideid interface{}

    // This object indicates the type of the flow, like media flow, signaling flow
    // etc. The type is Csbperflowstatsflowtype.
    Csbperflowstatsflowtype interface{}

    // This object indicates the number of RTP packets sent per flow by the SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Csbperflowstatsrtppktssent interface{}

    // This object indicates the number of RTP packets received per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Csbperflowstatsrtppktsrcvd interface{}

    // This object indicates the number of RTP packets discarded  per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Csbperflowstatsrtppktsdiscard interface{}

    // This object indicates the number of RTP octets sent per flow by the SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // octets.
    Csbperflowstatsrtpoctetssent interface{}

    // This object indicates the number of RTP octets received per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // octets.
    Csbperflowstatsrtpoctetsrcvd interface{}

    // This object indicates the number of RTP octets discarded per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // octets.
    Csbperflowstatsrtpoctetsdiscard interface{}

    // The number of RTP packets sent by the remote end point to this MG on this
    // flow. Comparing this with the local number of RTP packets received from the
    // remote end point gives an indication of how many incoming  packets were
    // dropped on this leg of the call. This information is from RTCP packet. Not
    // all endpoints report this statistic, if it is not available it will be set
    // to zero. This statistic will not be available for signaling flows. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    Csbperflowstatsrtcppktssent interface{}

    // The number of RTP packets received by the remote end point from this MG on
    // this flow. Comparing this with the local number of RTP packets sent from
    // this MG to the remote endpoint gives an indication of how many outgoing
    // packets were dropped on this leg of the call. This information is from RTCP
    // packet. Not all endpoints report this statistic, if it is not available it
    // will be set to zero. This statistic will not be available for signaling
    // flows. The type is interface{} with range: 0..18446744073709551615. Units
    // are packets.
    Csbperflowstatsrtcppktsrcvd interface{}

    // The number of RTP packets reported as lost by the remote end point on this
    // flow. This information is from RTCP packet. Not all endpoints report this
    // statistic, if it is not available it will be set to zero. This statistic
    // will not be available for signaling flows. The type is interface{} with
    // range: 0..18446744073709551615. Units are packets.
    Csbperflowstatsrtcppktslost interface{}

    // This object indicates the End Point jitter per flow in the SBC. The type is
    // interface{} with range: 0..18446744073709551615. Units are milliseconds.
    Csbperflowstatsepjitter interface{}

    // This object indicates the maximum burst size for the current FlowPair. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    Csbperflowstatstmanpermbs interface{}

    // This object indicates the bandwidth reserved for flow in kilobytes/second.
    // The type is interface{} with range: 0..4294967295. Units are kilobytes per
    // second.
    Csbperflowstatstmanpersdr interface{}

    // This object indicates the mark packets sent for the current FlowPair with,
    // or zero if none set. The DSCP is a 6-bit value, which will be present in
    // the top 6 bits of the lowest byte of this field. The type is string.
    Csbperflowstatsdscpsettings interface{}

    // This object indicates whether the flow on the current FlowPair has
    // subscribed for the NAT latch event. The type is string with length: 0..10.
    Csbperflowstatsadrstatus interface{}

    // This object indicates the flow on the current FlowPair has subscribed for
    // the media loss event. The type is string with length: 0..10.
    Csbperflowstatsqasettings interface{}

    // This object indicates the number of RTP packets lost per flow by the SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Csbperflowstatsrtppktslost interface{}
}

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetEntityData() *types.CommonEntityData {
    csbperflowstatsentry.EntityData.YFilter = csbperflowstatsentry.YFilter
    csbperflowstatsentry.EntityData.YangName = "csbPerFlowStatsEntry"
    csbperflowstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbperflowstatsentry.EntityData.ParentYangName = "csbPerFlowStatsTable"
    csbperflowstatsentry.EntityData.SegmentPath = "csbPerFlowStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbcallstatsserviceindex) + "']" + "[csbPerFlowStatsVdbeId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatsvdbeid) + "']" + "[csbPerFlowStatsGateId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatsgateid) + "']" + "[csbPerFlowStatsFlowPairId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatsflowpairid) + "']" + "[csbPerFlowStatsSideId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatssideid) + "']"
    csbperflowstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbperflowstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbperflowstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbperflowstatsentry.EntityData.Children = make(map[string]types.YChild)
    csbperflowstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbperflowstatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbperflowstatsentry.Csbcallstatsinstanceindex}
    csbperflowstatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbperflowstatsentry.Csbcallstatsserviceindex}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsVdbeId"] = types.YLeaf{"Csbperflowstatsvdbeid", csbperflowstatsentry.Csbperflowstatsvdbeid}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsGateId"] = types.YLeaf{"Csbperflowstatsgateid", csbperflowstatsentry.Csbperflowstatsgateid}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsFlowPairId"] = types.YLeaf{"Csbperflowstatsflowpairid", csbperflowstatsentry.Csbperflowstatsflowpairid}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsSideId"] = types.YLeaf{"Csbperflowstatssideid", csbperflowstatsentry.Csbperflowstatssideid}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsFlowType"] = types.YLeaf{"Csbperflowstatsflowtype", csbperflowstatsentry.Csbperflowstatsflowtype}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTPPktsSent"] = types.YLeaf{"Csbperflowstatsrtppktssent", csbperflowstatsentry.Csbperflowstatsrtppktssent}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTPPktsRcvd"] = types.YLeaf{"Csbperflowstatsrtppktsrcvd", csbperflowstatsentry.Csbperflowstatsrtppktsrcvd}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTPPktsDiscard"] = types.YLeaf{"Csbperflowstatsrtppktsdiscard", csbperflowstatsentry.Csbperflowstatsrtppktsdiscard}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTPOctetsSent"] = types.YLeaf{"Csbperflowstatsrtpoctetssent", csbperflowstatsentry.Csbperflowstatsrtpoctetssent}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTPOctetsRcvd"] = types.YLeaf{"Csbperflowstatsrtpoctetsrcvd", csbperflowstatsentry.Csbperflowstatsrtpoctetsrcvd}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTPOctetsDiscard"] = types.YLeaf{"Csbperflowstatsrtpoctetsdiscard", csbperflowstatsentry.Csbperflowstatsrtpoctetsdiscard}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTCPPktsSent"] = types.YLeaf{"Csbperflowstatsrtcppktssent", csbperflowstatsentry.Csbperflowstatsrtcppktssent}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTCPPktsRcvd"] = types.YLeaf{"Csbperflowstatsrtcppktsrcvd", csbperflowstatsentry.Csbperflowstatsrtcppktsrcvd}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTCPPktsLost"] = types.YLeaf{"Csbperflowstatsrtcppktslost", csbperflowstatsentry.Csbperflowstatsrtcppktslost}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsEPJitter"] = types.YLeaf{"Csbperflowstatsepjitter", csbperflowstatsentry.Csbperflowstatsepjitter}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsTmanPerMbs"] = types.YLeaf{"Csbperflowstatstmanpermbs", csbperflowstatsentry.Csbperflowstatstmanpermbs}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsTmanPerSdr"] = types.YLeaf{"Csbperflowstatstmanpersdr", csbperflowstatsentry.Csbperflowstatstmanpersdr}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsDscpSettings"] = types.YLeaf{"Csbperflowstatsdscpsettings", csbperflowstatsentry.Csbperflowstatsdscpsettings}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsAdrStatus"] = types.YLeaf{"Csbperflowstatsadrstatus", csbperflowstatsentry.Csbperflowstatsadrstatus}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsQASettings"] = types.YLeaf{"Csbperflowstatsqasettings", csbperflowstatsentry.Csbperflowstatsqasettings}
    csbperflowstatsentry.EntityData.Leafs["csbPerFlowStatsRTPPktsLost"] = types.YLeaf{"Csbperflowstatsrtppktslost", csbperflowstatsentry.Csbperflowstatsrtppktslost}
    return &(csbperflowstatsentry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatsflowtype represents signaling flow etc.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatsflowtype string

const (
    CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatsflowtype_media CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatsflowtype = "media"

    CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatsflowtype_signalling CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatsflowtype = "signalling"
)

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatssideid represents for the table.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatssideid string

const (
    CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatssideid_sideA CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatssideid = "sideA"

    CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatssideid_sideB CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry_Csbperflowstatssideid = "sideB"
)

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable
// This table describes the H.248 statistics for SBC. The index of
// the table is service index which corresponds to a particular 
// service configured on the SBC and the index assigned to a
// particular H.248 controller. The other index of this table is
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable.
// This table is replaced by the csbH248StatsRev1Table.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStath248Table. There is an entry in this
    // table for the particular controller by a value of csbH248StatsCtrlrIndex.
    // The other indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry.
    Csbh248Statsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry
}

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetEntityData() *types.CommonEntityData {
    csbh248Statstable.EntityData.YFilter = csbh248Statstable.YFilter
    csbh248Statstable.EntityData.YangName = "csbH248StatsTable"
    csbh248Statstable.EntityData.BundleName = "cisco_ios_xe"
    csbh248Statstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbh248Statstable.EntityData.SegmentPath = "csbH248StatsTable"
    csbh248Statstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbh248Statstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbh248Statstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbh248Statstable.EntityData.Children = make(map[string]types.YChild)
    csbh248Statstable.EntityData.Children["csbH248StatsEntry"] = types.YChild{"Csbh248Statsentry", nil}
    for i := range csbh248Statstable.Csbh248Statsentry {
        csbh248Statstable.EntityData.Children[types.GetSegmentPath(&csbh248Statstable.Csbh248Statsentry[i])] = types.YChild{"Csbh248Statsentry", &csbh248Statstable.Csbh248Statsentry[i]}
    }
    csbh248Statstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbh248Statstable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry
// An conceptual row in the csbCallStath248Table. There is
// an entry in this table for the particular controller by a value
// of csbH248StatsCtrlrIndex. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry_Csbcallstatsinstanceindex
    Csbcallstatsinstanceindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry_Csbcallstatsserviceindex
    Csbcallstatsserviceindex interface{}

    // This attribute is a key. This object identifies the controller index of the
    // H.248 server. This is also the index for the table. The type is interface{}
    // with range: 1..50.
    Csbh248Statsctrlrindex interface{}

    // This object indicates the requests sent through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrequestssent interface{}

    // This object indicates the requests received through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrequestsrcvd interface{}

    // This object indicates the requests failed on session Controller Interface
    // to an SBE or DBE. The type is interface{} with range: 0..4294967295.
    Csbh248Statsrequestsfailed interface{}

    // This object indicates the requests retried through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrequestsretried interface{}

    // This object indicates the number of replies sent through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrepliessent interface{}

    // This object indicates the number of replies received from the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrepliesrcvd interface{}

    // This object indicates the number of replies retried through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrepliesretried interface{}

    // This object indicates the number of packets sent through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statssegpktssent interface{}

    // This object indicates the number of packets received from the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statssegpktsrcvd interface{}

    // This object indicates the H.248 Controller established time (the time at
    // which the association became established). The type is string with length:
    // 0..80.
    Csbh248Statsestablishedtime interface{}

    // This object indicates the T-Max timeout value. This field specifies the
    // maximum delay (in milliseconds) for a response from an MGC before deciding
    // that the request has failed. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    Csbh248Statstmaxtimeoutval interface{}

    // This object indicates the calculated RTT value. This field specifies the
    // maximum round trip propagation delay in the  network (in milliseconds). The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Csbh248Statsrtt interface{}

    // This object indicates the LT value calculated from RTT value and Max
    // timeout value. This field specifies the maximum delay (in milliseconds) for
    // a response from an MGC plus the maximum round trip propagation delay in the
    // network (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Csbh248Statslt interface{}
}

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetEntityData() *types.CommonEntityData {
    csbh248Statsentry.EntityData.YFilter = csbh248Statsentry.YFilter
    csbh248Statsentry.EntityData.YangName = "csbH248StatsEntry"
    csbh248Statsentry.EntityData.BundleName = "cisco_ios_xe"
    csbh248Statsentry.EntityData.ParentYangName = "csbH248StatsTable"
    csbh248Statsentry.EntityData.SegmentPath = "csbH248StatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbh248Statsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbh248Statsentry.Csbcallstatsserviceindex) + "']" + "[csbH248StatsCtrlrIndex='" + fmt.Sprintf("%v", csbh248Statsentry.Csbh248Statsctrlrindex) + "']"
    csbh248Statsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbh248Statsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbh248Statsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbh248Statsentry.EntityData.Children = make(map[string]types.YChild)
    csbh248Statsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbh248Statsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbh248Statsentry.Csbcallstatsinstanceindex}
    csbh248Statsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbh248Statsentry.Csbcallstatsserviceindex}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsCtrlrIndex"] = types.YLeaf{"Csbh248Statsctrlrindex", csbh248Statsentry.Csbh248Statsctrlrindex}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRequestsSent"] = types.YLeaf{"Csbh248Statsrequestssent", csbh248Statsentry.Csbh248Statsrequestssent}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRequestsRcvd"] = types.YLeaf{"Csbh248Statsrequestsrcvd", csbh248Statsentry.Csbh248Statsrequestsrcvd}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRequestsFailed"] = types.YLeaf{"Csbh248Statsrequestsfailed", csbh248Statsentry.Csbh248Statsrequestsfailed}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRequestsRetried"] = types.YLeaf{"Csbh248Statsrequestsretried", csbh248Statsentry.Csbh248Statsrequestsretried}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRepliesSent"] = types.YLeaf{"Csbh248Statsrepliessent", csbh248Statsentry.Csbh248Statsrepliessent}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRepliesRcvd"] = types.YLeaf{"Csbh248Statsrepliesrcvd", csbh248Statsentry.Csbh248Statsrepliesrcvd}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRepliesRetried"] = types.YLeaf{"Csbh248Statsrepliesretried", csbh248Statsentry.Csbh248Statsrepliesretried}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsSegPktsSent"] = types.YLeaf{"Csbh248Statssegpktssent", csbh248Statsentry.Csbh248Statssegpktssent}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsSegPktsRcvd"] = types.YLeaf{"Csbh248Statssegpktsrcvd", csbh248Statsentry.Csbh248Statssegpktsrcvd}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsEstablishedTime"] = types.YLeaf{"Csbh248Statsestablishedtime", csbh248Statsentry.Csbh248Statsestablishedtime}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsTMaxTimeoutVal"] = types.YLeaf{"Csbh248Statstmaxtimeoutval", csbh248Statsentry.Csbh248Statstmaxtimeoutval}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsRTT"] = types.YLeaf{"Csbh248Statsrtt", csbh248Statsentry.Csbh248Statsrtt}
    csbh248Statsentry.EntityData.Leafs["csbH248StatsLT"] = types.YLeaf{"Csbh248Statslt", csbh248Statsentry.Csbh248Statslt}
    return &(csbh248Statsentry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table
// This table describes the H.248 statistics for SBC. The index of
// the table is service index which corresponds to a particular 
// service configured on the SBC and the index assigned to a
// particular H.248 controller. The other index of this table is
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStath248Table. There is an entry in this
    // table for the particular Vdbe by a value of csbH248StatsVdbeId. The other
    // indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry.
    Csbh248Statsrev1Entry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry
}

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetEntityData() *types.CommonEntityData {
    csbh248Statsrev1Table.EntityData.YFilter = csbh248Statsrev1Table.YFilter
    csbh248Statsrev1Table.EntityData.YangName = "csbH248StatsRev1Table"
    csbh248Statsrev1Table.EntityData.BundleName = "cisco_ios_xe"
    csbh248Statsrev1Table.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbh248Statsrev1Table.EntityData.SegmentPath = "csbH248StatsRev1Table"
    csbh248Statsrev1Table.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbh248Statsrev1Table.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbh248Statsrev1Table.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbh248Statsrev1Table.EntityData.Children = make(map[string]types.YChild)
    csbh248Statsrev1Table.EntityData.Children["csbH248StatsRev1Entry"] = types.YChild{"Csbh248Statsrev1Entry", nil}
    for i := range csbh248Statsrev1Table.Csbh248Statsrev1Entry {
        csbh248Statsrev1Table.EntityData.Children[types.GetSegmentPath(&csbh248Statsrev1Table.Csbh248Statsrev1Entry[i])] = types.YChild{"Csbh248Statsrev1Entry", &csbh248Statsrev1Table.Csbh248Statsrev1Entry[i]}
    }
    csbh248Statsrev1Table.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbh248Statsrev1Table.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry
// An conceptual row in the csbCallStath248Table. There is
// an entry in this table for the particular Vdbe by a value
// of csbH248StatsVdbeId. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry_Csbcallstatsinstanceindex
    Csbcallstatsinstanceindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry_Csbcallstatsserviceindex
    Csbcallstatsserviceindex interface{}

    // This attribute is a key. This object identifies the virtual media gateway
    // id. This is also the index for the table. The type is interface{} with
    // range: 0..255.
    Csbh248Statsvdbeid interface{}

    // This object indicates the requests sent through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrequestssentrev1 interface{}

    // This object indicates the requests received through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrequestsrcvdrev1 interface{}

    // This object indicates the requests failed on session Controller Interface
    // to an SBE or DBE. The type is interface{} with range: 0..4294967295.
    Csbh248Statsrequestsfailedrev1 interface{}

    // This object indicates the requests retried through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrequestsretriedrev1 interface{}

    // This object indicates the number of replies sent through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrepliessentrev1 interface{}

    // This object indicates the number of replies received from the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrepliesrcvdrev1 interface{}

    // This object indicates the number of replies retried through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    Csbh248Statsrepliesretriedrev1 interface{}

    // This object indicates the number of response segments sent by DBE. This
    // field will only be present if segmentation is enabled on this association.
    // The type is interface{} with range: 0..4294967295.
    Csbh248Statssegpktssentrev1 interface{}

    // This object indicates the number of response segments received by DBE. This
    // field will only be present if segmentation is enabled on this association.
    // The type is interface{} with range: 0..4294967295.
    Csbh248Statssegpktsrcvdrev1 interface{}

    // This object indicates the H.248 Controller established time (the time at
    // which the association became established). The type is string with length:
    // 0..80.
    Csbh248Statsestablishedtimerev1 interface{}

    // This object indicates the T-Max timeout value. This field specifies the
    // maximum delay (in milliseconds) for a response from an MGC before deciding
    // that the request has failed. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    Csbh248Statstmaxtimeoutvalrev1 interface{}

    // This object indicates the calculated RTT value. This field specifies the
    // maximum round trip propagation delay in the  network (in milliseconds). The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Csbh248Statsrttrev1 interface{}

    // This object indicates the LT value calculated from RTT value and Max
    // timeout value. This field specifies the maximum delay (in milliseconds) for
    // a response from an MGC plus the maximum round trip propagation delay in the
    // network (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Csbh248Statsltrev1 interface{}
}

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetEntityData() *types.CommonEntityData {
    csbh248Statsrev1Entry.EntityData.YFilter = csbh248Statsrev1Entry.YFilter
    csbh248Statsrev1Entry.EntityData.YangName = "csbH248StatsRev1Entry"
    csbh248Statsrev1Entry.EntityData.BundleName = "cisco_ios_xe"
    csbh248Statsrev1Entry.EntityData.ParentYangName = "csbH248StatsRev1Table"
    csbh248Statsrev1Entry.EntityData.SegmentPath = "csbH248StatsRev1Entry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbh248Statsrev1Entry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbh248Statsrev1Entry.Csbcallstatsserviceindex) + "']" + "[csbH248StatsVdbeId='" + fmt.Sprintf("%v", csbh248Statsrev1Entry.Csbh248Statsvdbeid) + "']"
    csbh248Statsrev1Entry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbh248Statsrev1Entry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbh248Statsrev1Entry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbh248Statsrev1Entry.EntityData.Children = make(map[string]types.YChild)
    csbh248Statsrev1Entry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbh248Statsrev1Entry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbh248Statsrev1Entry.Csbcallstatsinstanceindex}
    csbh248Statsrev1Entry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbh248Statsrev1Entry.Csbcallstatsserviceindex}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsVdbeId"] = types.YLeaf{"Csbh248Statsvdbeid", csbh248Statsrev1Entry.Csbh248Statsvdbeid}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRequestsSentRev1"] = types.YLeaf{"Csbh248Statsrequestssentrev1", csbh248Statsrev1Entry.Csbh248Statsrequestssentrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRequestsRcvdRev1"] = types.YLeaf{"Csbh248Statsrequestsrcvdrev1", csbh248Statsrev1Entry.Csbh248Statsrequestsrcvdrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRequestsFailedRev1"] = types.YLeaf{"Csbh248Statsrequestsfailedrev1", csbh248Statsrev1Entry.Csbh248Statsrequestsfailedrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRequestsRetriedRev1"] = types.YLeaf{"Csbh248Statsrequestsretriedrev1", csbh248Statsrev1Entry.Csbh248Statsrequestsretriedrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRepliesSentRev1"] = types.YLeaf{"Csbh248Statsrepliessentrev1", csbh248Statsrev1Entry.Csbh248Statsrepliessentrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRepliesRcvdRev1"] = types.YLeaf{"Csbh248Statsrepliesrcvdrev1", csbh248Statsrev1Entry.Csbh248Statsrepliesrcvdrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRepliesRetriedRev1"] = types.YLeaf{"Csbh248Statsrepliesretriedrev1", csbh248Statsrev1Entry.Csbh248Statsrepliesretriedrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsSegPktsSentRev1"] = types.YLeaf{"Csbh248Statssegpktssentrev1", csbh248Statsrev1Entry.Csbh248Statssegpktssentrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsSegPktsRcvdRev1"] = types.YLeaf{"Csbh248Statssegpktsrcvdrev1", csbh248Statsrev1Entry.Csbh248Statssegpktsrcvdrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsEstablishedTimeRev1"] = types.YLeaf{"Csbh248Statsestablishedtimerev1", csbh248Statsrev1Entry.Csbh248Statsestablishedtimerev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsTMaxTimeoutValRev1"] = types.YLeaf{"Csbh248Statstmaxtimeoutvalrev1", csbh248Statsrev1Entry.Csbh248Statstmaxtimeoutvalrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsRTTRev1"] = types.YLeaf{"Csbh248Statsrttrev1", csbh248Statsrev1Entry.Csbh248Statsrttrev1}
    csbh248Statsrev1Entry.EntityData.Leafs["csbH248StatsLTRev1"] = types.YLeaf{"Csbh248Statsltrev1", csbh248Statsrev1Entry.Csbh248Statsltrev1}
    return &(csbh248Statsrev1Entry.EntityData)
}

