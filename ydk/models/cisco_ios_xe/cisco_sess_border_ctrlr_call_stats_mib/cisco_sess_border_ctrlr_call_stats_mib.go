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
    CsbCallStatsInstanceTable CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable

    // This table describes the global statistics information in the form of a
    // table which contains call specific information like call rates, media
    // flows, signaling flows etc. The index of the table is service index which
    // corresponds to a particular  service configured on the SBC and all the rows
    // of the table represents the global information regarding all the call flows
    // related to that particular service. The other index of this table is
    // csbCallStatsInstanceIndex which is defined in csbCallStatsInstanceTable.
    CsbCallStatsTable CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable

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
    CsbCurrPeriodicStatsTable CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable

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
    CsbHistoryStatsTable CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable

    // This table describes statistics table for each call flow. The indices of
    // the table are virtual media gateway id, gate id, falow pair id and side id
    // (indices for side A or side B). The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable.
    CsbPerFlowStatsTable CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable

    // This table describes the H.248 statistics for SBC. The index of the table
    // is service index which corresponds to a particular  service configured on
    // the SBC and the index assigned to a particular H.248 controller. The other
    // index of this table is csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable. This table is replaced by the
    // csbH248StatsRev1Table.
    CsbH248StatsTable CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable

    // This table describes the H.248 statistics for SBC. The index of the table
    // is service index which corresponds to a particular  service configured on
    // the SBC and the index assigned to a particular H.248 controller. The other
    // index of this table is csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable.
    CsbH248StatsRev1Table CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table
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

    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children = types.NewOrderedMap()
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children.Append("csbCallStatsInstanceTable", types.YChild{"CsbCallStatsInstanceTable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.CsbCallStatsInstanceTable})
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children.Append("csbCallStatsTable", types.YChild{"CsbCallStatsTable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.CsbCallStatsTable})
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children.Append("csbCurrPeriodicStatsTable", types.YChild{"CsbCurrPeriodicStatsTable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.CsbCurrPeriodicStatsTable})
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children.Append("csbHistoryStatsTable", types.YChild{"CsbHistoryStatsTable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.CsbHistoryStatsTable})
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children.Append("csbPerFlowStatsTable", types.YChild{"CsbPerFlowStatsTable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.CsbPerFlowStatsTable})
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children.Append("csbH248StatsTable", types.YChild{"CsbH248StatsTable", &cISCOSESSBORDERCTRLRCALLSTATSMIB.CsbH248StatsTable})
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Children.Append("csbH248StatsRev1Table", types.YChild{"CsbH248StatsRev1Table", &cISCOSESSBORDERCTRLRCALLSTATSMIB.CsbH248StatsRev1Table})
    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData.YListKeys = []string {}

    return &(cISCOSESSBORDERCTRLRCALLSTATSMIB.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable
// The call stats instance table contains the physical index for
// each of the physical entity (line card, primary, secondary
// cards). The index of the table is instance index which uniquely
// identifies the physical entity present on the box.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in csbCallStatsInstanceTable. There is an entry in this
    // table for each SBC instance, as identified by a  value of
    // csbCallStatsInstanceIndex. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry.
    CsbCallStatsInstanceEntry []*CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry
}

func (csbCallStatsInstanceTable *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable) GetEntityData() *types.CommonEntityData {
    csbCallStatsInstanceTable.EntityData.YFilter = csbCallStatsInstanceTable.YFilter
    csbCallStatsInstanceTable.EntityData.YangName = "csbCallStatsInstanceTable"
    csbCallStatsInstanceTable.EntityData.BundleName = "cisco_ios_xe"
    csbCallStatsInstanceTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbCallStatsInstanceTable.EntityData.SegmentPath = "csbCallStatsInstanceTable"
    csbCallStatsInstanceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbCallStatsInstanceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbCallStatsInstanceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbCallStatsInstanceTable.EntityData.Children = types.NewOrderedMap()
    csbCallStatsInstanceTable.EntityData.Children.Append("csbCallStatsInstanceEntry", types.YChild{"CsbCallStatsInstanceEntry", nil})
    for i := range csbCallStatsInstanceTable.CsbCallStatsInstanceEntry {
        csbCallStatsInstanceTable.EntityData.Children.Append(types.GetSegmentPath(csbCallStatsInstanceTable.CsbCallStatsInstanceEntry[i]), types.YChild{"CsbCallStatsInstanceEntry", csbCallStatsInstanceTable.CsbCallStatsInstanceEntry[i]})
    }
    csbCallStatsInstanceTable.EntityData.Leafs = types.NewOrderedMap()

    csbCallStatsInstanceTable.EntityData.YListKeys = []string {}

    return &(csbCallStatsInstanceTable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry
// A conceptual row in csbCallStatsInstanceTable. There is an
// entry in this table for each SBC instance, as identified by a 
// value of csbCallStatsInstanceIndex.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object uniquely identifies the sequence
    // number of an entity or slot that is configured per device. This index is
    // assigned arbitrarily by the engine and is not saved over reboots. The type
    // is interface{} with range: 0..4294967295.
    CsbCallStatsInstanceIndex interface{}

    // This object indicates the physical entity for which all the measurements
    // are maintained. The exact type of this entity is described by its
    // entPhysicalVendorType value. The type is interface{} with range:
    // 0..2147483647.
    CsbCallStatsInstancePhysicalIndex interface{}
}

func (csbCallStatsInstanceEntry *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry) GetEntityData() *types.CommonEntityData {
    csbCallStatsInstanceEntry.EntityData.YFilter = csbCallStatsInstanceEntry.YFilter
    csbCallStatsInstanceEntry.EntityData.YangName = "csbCallStatsInstanceEntry"
    csbCallStatsInstanceEntry.EntityData.BundleName = "cisco_ios_xe"
    csbCallStatsInstanceEntry.EntityData.ParentYangName = "csbCallStatsInstanceTable"
    csbCallStatsInstanceEntry.EntityData.SegmentPath = "csbCallStatsInstanceEntry" + types.AddKeyToken(csbCallStatsInstanceEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex")
    csbCallStatsInstanceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbCallStatsInstanceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbCallStatsInstanceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbCallStatsInstanceEntry.EntityData.Children = types.NewOrderedMap()
    csbCallStatsInstanceEntry.EntityData.Leafs = types.NewOrderedMap()
    csbCallStatsInstanceEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbCallStatsInstanceEntry.CsbCallStatsInstanceIndex})
    csbCallStatsInstanceEntry.EntityData.Leafs.Append("csbCallStatsInstancePhysicalIndex", types.YLeaf{"CsbCallStatsInstancePhysicalIndex", csbCallStatsInstanceEntry.CsbCallStatsInstancePhysicalIndex})

    csbCallStatsInstanceEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex"}

    return &(csbCallStatsInstanceEntry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable
// This table describes the global statistics information in the
// form of a table which contains call specific information like
// call rates, media flows, signaling flows etc. The index of the
// table is service index which corresponds to a particular 
// service configured on the SBC and all the rows of the table
// represents the global information regarding all the call flows
// related to that particular service. The other index of this
// table is csbCallStatsInstanceIndex which is defined in
// csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStatsGlobalStatsTable. There is an entry in
    // this table for the particular service configured on SBC identified by a
    // value of csbCallStatsInstanceIndex. The other index of this table is
    // csbCallStatsInstanceIndex which is defined in csbCallStatsInstanceTable.
    // The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry.
    CsbCallStatsEntry []*CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry
}

func (csbCallStatsTable *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable) GetEntityData() *types.CommonEntityData {
    csbCallStatsTable.EntityData.YFilter = csbCallStatsTable.YFilter
    csbCallStatsTable.EntityData.YangName = "csbCallStatsTable"
    csbCallStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbCallStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbCallStatsTable.EntityData.SegmentPath = "csbCallStatsTable"
    csbCallStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbCallStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbCallStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbCallStatsTable.EntityData.Children = types.NewOrderedMap()
    csbCallStatsTable.EntityData.Children.Append("csbCallStatsEntry", types.YChild{"CsbCallStatsEntry", nil})
    for i := range csbCallStatsTable.CsbCallStatsEntry {
        csbCallStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbCallStatsTable.CsbCallStatsEntry[i]), types.YChild{"CsbCallStatsEntry", csbCallStatsTable.CsbCallStatsEntry[i]})
    }
    csbCallStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbCallStatsTable.EntityData.YListKeys = []string {}

    return &(csbCallStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry
// An conceptual row in the csbCallStatsGlobalStatsTable. There is
// an entry in this table for the particular service configured on
// SBC identified by a value of csbCallStatsInstanceIndex. The
// other index of this table is csbCallStatsInstanceIndex which is
// defined in csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry_CsbCallStatsInstanceIndex
    CsbCallStatsInstanceIndex interface{}

    // This attribute is a key. This object identifies the index of the name of
    // the SBC service configured. This object also acts as an index for the
    // table. The type is interface{} with range: 0..4294967295.
    CsbCallStatsServiceIndex interface{}

    // This object indicates the configured name of the SBC service. The length of
    // this object is zero when value is not assigned to it. The type is string.
    CsbCallStatsSbcName interface{}

    // This object indicates the maximum number of calls per second processed by
    // the Session Border Controller in past 24 hours. The type is interface{}
    // with range: 0..4294967295. Units are calls per second.
    CsbCallStatsCallsHigh interface{}

    // This object indicates the average calls per second processed by the Session
    // Border Controller. The type is interface{} with range: 0..4294967295. Units
    // are calls per second.
    CsbCallStatsRate1Sec interface{}

    // This object indicates the minimum calls per second in past 24 hours. The
    // type is interface{} with range: 0..4294967295. Units are calls per second.
    CsbCallStatsCallsLow interface{}

    // This object indicates the number of media flows which are available. The
    // type is interface{} with range: 0..4294967295. Units are flows.
    CsbCallStatsAvailableFlows interface{}

    // This object indicates the number of media flows which are used. This is for
    // the flow allocated and connected. The flow allocated but not connected is
    // not counted. The type is interface{} with range: 0..4294967295. Units are
    // flows.
    CsbCallStatsUsedFlows interface{}

    // This object indicates the number of peak flows in SBC. This is the highest
    // recorded value for the active flows since the box was booted/reseted. The
    // type is interface{} with range: 0..4294967295. Units are flows.
    CsbCallStatsPeakFlows interface{}

    // This object indicates the total number of media support by this instance of
    // SBC. The total number of flows include the available flows and the active
    // flows. This value is since box was booted/reseted. Total flows include the
    // active flows and the flows allocated but not connected. The type is
    // interface{} with range: 0..4294967295. Units are flows.
    CsbCallStatsTotalFlows interface{}

    // This object indicates the number of active signaling flows for signaling
    // pinholes. The type is interface{} with range: 0..4294967295. Units are
    // signal flows.
    CsbCallStatsUsedSigFlows interface{}

    // This object indicates the peak signaling flow in SBC. This is the highest
    // recorded value for the active signaling flows. This object is calculated
    // using csbCallStatsUsedFlows. The type is interface{} with range:
    // 0..4294967295. Units are signal flows.
    CsbCallStatsPeakSigFlows interface{}

    // This object indicates the maximum number of Signalling Flows that can be
    // supported by this instance of SBC. The type is interface{} with range:
    // 0..4294967295. Units are signal flows.
    CsbCallStatsTotalSigFlows interface{}

    // This object indicates the remaining capacity that can be supported by SBC.
    // The type is interface{} with range: 0..4294967295. Units are packets per
    // second.
    CsbCallStatsAvailablePktRate interface{}

    // This object indicates the number of unclassified packets processed by SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CsbCallStatsUnclassifiedPkts interface{}

    // This object indicates the total number of RTP packets sent. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    CsbCallStatsRTPPktsSent interface{}

    // This object indicates the total number of RTP packets received. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    CsbCallStatsRTPPktsRcvd interface{}

    // This object indicates the total number of RTP packets discarded. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    CsbCallStatsRTPPktsDiscard interface{}

    // This object indicates the number of RTP octets sent by the SBC. The type is
    // interface{} with range: 0..18446744073709551615. Units are octets.
    CsbCallStatsRTPOctetsSent interface{}

    // This object indicates the number of RTP octets received by the SBC. The
    // type is interface{} with range: 0..18446744073709551615. Units are octets.
    CsbCallStatsRTPOctetsRcvd interface{}

    // This object indicates the number of RTP octets discarded by the SBC. The
    // type is interface{} with range: 0..18446744073709551615. Units are octets.
    CsbCallStatsRTPOctetsDiscard interface{}

    // This object indicates the accumulated No media event count. The type is
    // interface{} with range: 0..4294967295. Units are no media events.
    CsbCallStatsNoMediaCount interface{}

    // This object indicates the accumulated route error event count. This counter
    // is for the route error of media stream. The type is interface{} with range:
    // 0..4294967295. Units are route error events.
    CsbCallStatsRouteErrors interface{}

    // This object indicates the number of additional transcoded flows that this
    // media gateway manager (MGM) entity is currently able to configure. The type
    // is interface{} with range: 0..4294967295. Units are flows.
    CsbCallStatsAvailableTranscodeFlows interface{}

    // This object indicates the current number of transcoded flows that are
    // actively forwarding media traffic.  In this context, a flow is a media
    // stream passing through the device. So a single voice call will be counted
    // only once. The type is interface{} with range: 0..4294967295. Units are
    // flows.
    CsbCallStatsActiveTranscodeFlows interface{}

    // This object indicates the peak number of active transcoded flows since the
    // statistics were last reset.  In this context, a flow is a media stream
    // passing through the device, so a single voice call will be counted once.
    // The type is interface{} with range: 0..4294967295. Units are flows.
    CsbCallStatsPeakTranscodeFlows interface{}

    // This object indicates the accumulated total number of transcoded flows. 
    // This count contains both active flows and flows that were allocated but
    // never connected.  In this context, a flow is a media stream passing through
    // the device, so a single voice call will be counted once. The type is
    // interface{} with range: 0..4294967295. Units are flows.
    CsbCallStatsTotalTranscodeFlows interface{}
}

func (csbCallStatsEntry *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry) GetEntityData() *types.CommonEntityData {
    csbCallStatsEntry.EntityData.YFilter = csbCallStatsEntry.YFilter
    csbCallStatsEntry.EntityData.YangName = "csbCallStatsEntry"
    csbCallStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbCallStatsEntry.EntityData.ParentYangName = "csbCallStatsTable"
    csbCallStatsEntry.EntityData.SegmentPath = "csbCallStatsEntry" + types.AddKeyToken(csbCallStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbCallStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex")
    csbCallStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbCallStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbCallStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbCallStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbCallStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbCallStatsEntry.CsbCallStatsInstanceIndex})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbCallStatsEntry.CsbCallStatsServiceIndex})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsSbcName", types.YLeaf{"CsbCallStatsSbcName", csbCallStatsEntry.CsbCallStatsSbcName})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsCallsHigh", types.YLeaf{"CsbCallStatsCallsHigh", csbCallStatsEntry.CsbCallStatsCallsHigh})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRate1Sec", types.YLeaf{"CsbCallStatsRate1Sec", csbCallStatsEntry.CsbCallStatsRate1Sec})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsCallsLow", types.YLeaf{"CsbCallStatsCallsLow", csbCallStatsEntry.CsbCallStatsCallsLow})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsAvailableFlows", types.YLeaf{"CsbCallStatsAvailableFlows", csbCallStatsEntry.CsbCallStatsAvailableFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsUsedFlows", types.YLeaf{"CsbCallStatsUsedFlows", csbCallStatsEntry.CsbCallStatsUsedFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsPeakFlows", types.YLeaf{"CsbCallStatsPeakFlows", csbCallStatsEntry.CsbCallStatsPeakFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsTotalFlows", types.YLeaf{"CsbCallStatsTotalFlows", csbCallStatsEntry.CsbCallStatsTotalFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsUsedSigFlows", types.YLeaf{"CsbCallStatsUsedSigFlows", csbCallStatsEntry.CsbCallStatsUsedSigFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsPeakSigFlows", types.YLeaf{"CsbCallStatsPeakSigFlows", csbCallStatsEntry.CsbCallStatsPeakSigFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsTotalSigFlows", types.YLeaf{"CsbCallStatsTotalSigFlows", csbCallStatsEntry.CsbCallStatsTotalSigFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsAvailablePktRate", types.YLeaf{"CsbCallStatsAvailablePktRate", csbCallStatsEntry.CsbCallStatsAvailablePktRate})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsUnclassifiedPkts", types.YLeaf{"CsbCallStatsUnclassifiedPkts", csbCallStatsEntry.CsbCallStatsUnclassifiedPkts})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRTPPktsSent", types.YLeaf{"CsbCallStatsRTPPktsSent", csbCallStatsEntry.CsbCallStatsRTPPktsSent})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRTPPktsRcvd", types.YLeaf{"CsbCallStatsRTPPktsRcvd", csbCallStatsEntry.CsbCallStatsRTPPktsRcvd})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRTPPktsDiscard", types.YLeaf{"CsbCallStatsRTPPktsDiscard", csbCallStatsEntry.CsbCallStatsRTPPktsDiscard})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRTPOctetsSent", types.YLeaf{"CsbCallStatsRTPOctetsSent", csbCallStatsEntry.CsbCallStatsRTPOctetsSent})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRTPOctetsRcvd", types.YLeaf{"CsbCallStatsRTPOctetsRcvd", csbCallStatsEntry.CsbCallStatsRTPOctetsRcvd})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRTPOctetsDiscard", types.YLeaf{"CsbCallStatsRTPOctetsDiscard", csbCallStatsEntry.CsbCallStatsRTPOctetsDiscard})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsNoMediaCount", types.YLeaf{"CsbCallStatsNoMediaCount", csbCallStatsEntry.CsbCallStatsNoMediaCount})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsRouteErrors", types.YLeaf{"CsbCallStatsRouteErrors", csbCallStatsEntry.CsbCallStatsRouteErrors})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsAvailableTranscodeFlows", types.YLeaf{"CsbCallStatsAvailableTranscodeFlows", csbCallStatsEntry.CsbCallStatsAvailableTranscodeFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsActiveTranscodeFlows", types.YLeaf{"CsbCallStatsActiveTranscodeFlows", csbCallStatsEntry.CsbCallStatsActiveTranscodeFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsPeakTranscodeFlows", types.YLeaf{"CsbCallStatsPeakTranscodeFlows", csbCallStatsEntry.CsbCallStatsPeakTranscodeFlows})
    csbCallStatsEntry.EntityData.Leafs.Append("csbCallStatsTotalTranscodeFlows", types.YLeaf{"CsbCallStatsTotalTranscodeFlows", csbCallStatsEntry.CsbCallStatsTotalTranscodeFlows})

    csbCallStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex"}

    return &(csbCallStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable
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
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCurrPeriodicStatsTable. There is an entry in
    // this table for the particular controller by a value of
    // csbH248StatsCtrlrIndex. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable_CsbCurrPeriodicStatsEntry.
    CsbCurrPeriodicStatsEntry []*CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable_CsbCurrPeriodicStatsEntry
}

func (csbCurrPeriodicStatsTable *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable) GetEntityData() *types.CommonEntityData {
    csbCurrPeriodicStatsTable.EntityData.YFilter = csbCurrPeriodicStatsTable.YFilter
    csbCurrPeriodicStatsTable.EntityData.YangName = "csbCurrPeriodicStatsTable"
    csbCurrPeriodicStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbCurrPeriodicStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbCurrPeriodicStatsTable.EntityData.SegmentPath = "csbCurrPeriodicStatsTable"
    csbCurrPeriodicStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbCurrPeriodicStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbCurrPeriodicStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbCurrPeriodicStatsTable.EntityData.Children = types.NewOrderedMap()
    csbCurrPeriodicStatsTable.EntityData.Children.Append("csbCurrPeriodicStatsEntry", types.YChild{"CsbCurrPeriodicStatsEntry", nil})
    for i := range csbCurrPeriodicStatsTable.CsbCurrPeriodicStatsEntry {
        csbCurrPeriodicStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbCurrPeriodicStatsTable.CsbCurrPeriodicStatsEntry[i]), types.YChild{"CsbCurrPeriodicStatsEntry", csbCurrPeriodicStatsTable.CsbCurrPeriodicStatsEntry[i]})
    }
    csbCurrPeriodicStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbCurrPeriodicStatsTable.EntityData.YListKeys = []string {}

    return &(csbCurrPeriodicStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable_CsbCurrPeriodicStatsEntry
// An conceptual row in the csbCurrPeriodicStatsTable. There is
// an entry in this table for the particular controller by a value
// of csbH248StatsCtrlrIndex. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable_CsbCurrPeriodicStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry_CsbCallStatsInstanceIndex
    CsbCallStatsInstanceIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry_CsbCallStatsServiceIndex
    CsbCallStatsServiceIndex interface{}

    // This attribute is a key. This object identifies the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be 5 min, 15 mins, 1 hour , 1 Day. This object acts as index for the table.
    // The type is CiscoSbcPeriodicStatsInterval.
    CsbCurrPeriodicStatsInterval interface{}

    // This object indicates the number of calls that have become active during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbCurrPeriodicStatsActiveCalls interface{}

    // This object indicates the number of calls that have become activing during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbCurrPeriodicStatsActivatingCalls interface{}

    // This object indicates the number of calls that have become deactiving
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    CsbCurrPeriodicStatsDeactivatingCalls interface{}

    // This object indicates the number of total call attempts during this
    // interval. The type is interface{} with range: 0..4294967295.
    CsbCurrPeriodicStatsTotalCallAttempts interface{}

    // This object indicates the number of failed call attempts during this
    // interval. The type is interface{} with range: 0..4294967295.
    CsbCurrPeriodicStatsFailedCallAttempts interface{}

    // This object indicates the number of call setup failures due to routing
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallRoutingFailure interface{}

    // This object indicates the number of call setup failures due to resource
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallResourceFailure interface{}

    // This object indicates the number of call setup failures due to media
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallMediaFailure interface{}

    // This object indicates the number of call setup failures due to signaling
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSigFailure interface{}

    // This object indicates the number of active call failures during this
    // interval. The type is interface{} with range: 0..4294967295.
    CsbCurrPeriodicStatsActiveCallFailure interface{}

    // This object indicates the number of call setup failures due to congestion
    // during this interval. The type is interface{} with range: 0..4294967295.
    CsbCurrPeriodicStatsCongestionFailure interface{}

    // This object indicates the number of call setup failures due to policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupPolicyFailure interface{}

    // This object indicates the number of call setup failures due to NA policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupNAPolicyFailure interface{}

    // This object indicates the number of call setup failures due to routing
    // policy failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupRoutingPolicyFailure interface{}

    // This object indicates the number of call setup failures due to CAC policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupCACPolicyFailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupCACCallLimitFailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // rate limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupCACRateLimitFailure interface{}

    // This object indicates the number of call setup failures due to CAC
    // bandwidth limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupCACBandwidthFailure interface{}

    // This object indicates the number of call setup failures due to CAC media
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupCACMediaLimitFailure interface{}

    // This object indicates the number of call update failure due to policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbCurrPeriodicStatsCallSetupCACMediaUpdateFailure interface{}

    // This object indicates the current time at the start of each interval. The
    // type is string with length: 0..80.
    CsbCurrPeriodicStatsTimestamp interface{}

    // The object indicates the number of transcoded calls that are active during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbCurrPeriodicStatsTranscodedCalls interface{}

    // The object indicates the number of transrated calls that are active during
    // this interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbCurrPeriodicStatsTransratedCalls interface{}

    // This object indicates the total number of call update failures during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbCurrPeriodicStatsTotalCallUpdateFailure interface{}

    // This Object indicates the number of calls through SBC that use IPv6
    // signaling.  This statistic totals all calls that traverse an IPv6 adjacency
    // on either or both sides of SBC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsActiveIpv6Calls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis) during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsActiveEmergencyCalls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis) and have used the e2
    // interface to obtain location information for the caller during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbCurrPeriodicStatsActiveE2EmergencyCalls interface{}

    // This object indicates the total number of active calls which use IMS Rx,
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    CsbCurrPeriodicStatsImsRxActiveCalls interface{}

    // This object indicates the total call Setup failures owing to IMS Rx failure
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are failures.
    CsbCurrPeriodicStatsImsRxCallSetupFaiures interface{}

    // This object indicates the total call renegotiation attempts using IMS Rx
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are attempts.
    CsbCurrPeriodicStatsImsRxCallRenegotiationAttempts interface{}

    // This object indicates the total call renegotiation failures owing to IMS Rx
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295. Units are failures.
    CsbCurrPeriodicStatsImsRxCallRenegotiationFailures interface{}

    // The number of active audio transcoded calls through this adjacency or
    // account during this interval. The type is interface{} with range:
    // 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsAudioTranscodedCalls interface{}

    // This object indicates the the number of active fax transcoded calls through
    // this adjacency or account during this interval. The type is interface{}
    // with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsFaxTranscodedCalls interface{}

    // This object indicates the total call setup failures due to RTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    CsbCurrPeriodicStatsRtpDisallowedFailures interface{}

    // This object indicates the total call setup failures due to SRTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    CsbCurrPeriodicStatsSrtpDisallowedFailures interface{}

    // This object indicates the number of active calls through this adjacency or
    // account which do not use SRTP on any media channels during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsNonSrtpCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels which use SRTP. This count
    // does not include media  channels that provide interworking between RTP and
    // SRTP during this interval. The type is interface{} with range:
    // 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsSrtpNonIwCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels that provide interworking
    // between RTP and SRTP during this interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsSrtpIwCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via RFC 2833 during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsDtmfIw2833Calls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via  inband DTMF tones during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsDtmfIwInbandCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in media via
    // RFC 2833 and DTMF in media via inband DTMF tones during this interval. The
    // type is interface{} with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicStatsDtmfIw2833InbandCalls interface{}

    // This object indicates the lawful intercept tap attempts requested within
    // the scope of this query during this interval. The type is interface{} with
    // range: 0..4294967295. Units are attempts.
    CsbCurrPeriodicStatsTotalTapsRequested interface{}

    // This object indicates the lawful intercept tap attempts that have been
    // successfully implemented within the scope of this query during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // success.
    CsbCurrPeriodicStatsTotalTapsSucceeded interface{}

    // This object indicates the Lawful intercept taps currently in place on calls
    // within the scope of this query during this interval. The type is
    // interface{} with range: 0..4294967295. Units are taps.
    CsbCurrPeriodicStatsCurrentTaps interface{}

    // The number of active calls on this adjacency or account which are to or
    // from registered subscribers using IPSEC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbCurrPeriodicIpsecCalls interface{}
}

func (csbCurrPeriodicStatsEntry *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCurrPeriodicStatsTable_CsbCurrPeriodicStatsEntry) GetEntityData() *types.CommonEntityData {
    csbCurrPeriodicStatsEntry.EntityData.YFilter = csbCurrPeriodicStatsEntry.YFilter
    csbCurrPeriodicStatsEntry.EntityData.YangName = "csbCurrPeriodicStatsEntry"
    csbCurrPeriodicStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbCurrPeriodicStatsEntry.EntityData.ParentYangName = "csbCurrPeriodicStatsTable"
    csbCurrPeriodicStatsEntry.EntityData.SegmentPath = "csbCurrPeriodicStatsEntry" + types.AddKeyToken(csbCurrPeriodicStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbCurrPeriodicStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsInterval, "csbCurrPeriodicStatsInterval")
    csbCurrPeriodicStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbCurrPeriodicStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbCurrPeriodicStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbCurrPeriodicStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbCurrPeriodicStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbCurrPeriodicStatsEntry.CsbCallStatsInstanceIndex})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbCurrPeriodicStatsEntry.CsbCallStatsServiceIndex})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsInterval", types.YLeaf{"CsbCurrPeriodicStatsInterval", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsInterval})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsActiveCalls", types.YLeaf{"CsbCurrPeriodicStatsActiveCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsActiveCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsActivatingCalls", types.YLeaf{"CsbCurrPeriodicStatsActivatingCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsActivatingCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsDeactivatingCalls", types.YLeaf{"CsbCurrPeriodicStatsDeactivatingCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsDeactivatingCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsTotalCallAttempts", types.YLeaf{"CsbCurrPeriodicStatsTotalCallAttempts", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsTotalCallAttempts})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsFailedCallAttempts", types.YLeaf{"CsbCurrPeriodicStatsFailedCallAttempts", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsFailedCallAttempts})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallRoutingFailure", types.YLeaf{"CsbCurrPeriodicStatsCallRoutingFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallRoutingFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallResourceFailure", types.YLeaf{"CsbCurrPeriodicStatsCallResourceFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallResourceFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallMediaFailure", types.YLeaf{"CsbCurrPeriodicStatsCallMediaFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallMediaFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSigFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSigFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSigFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsActiveCallFailure", types.YLeaf{"CsbCurrPeriodicStatsActiveCallFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsActiveCallFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCongestionFailure", types.YLeaf{"CsbCurrPeriodicStatsCongestionFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCongestionFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupPolicyFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupPolicyFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupPolicyFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupNAPolicyFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupNAPolicyFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupNAPolicyFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupRoutingPolicyFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupRoutingPolicyFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupRoutingPolicyFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupCACPolicyFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupCACPolicyFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupCACPolicyFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupCACCallLimitFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupCACCallLimitFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupCACCallLimitFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupCACRateLimitFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupCACRateLimitFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupCACRateLimitFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupCACBandwidthFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupCACBandwidthFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupCACBandwidthFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupCACMediaLimitFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupCACMediaLimitFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupCACMediaLimitFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCallSetupCACMediaUpdateFailure", types.YLeaf{"CsbCurrPeriodicStatsCallSetupCACMediaUpdateFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCallSetupCACMediaUpdateFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsTimestamp", types.YLeaf{"CsbCurrPeriodicStatsTimestamp", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsTimestamp})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsTranscodedCalls", types.YLeaf{"CsbCurrPeriodicStatsTranscodedCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsTranscodedCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsTransratedCalls", types.YLeaf{"CsbCurrPeriodicStatsTransratedCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsTransratedCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsTotalCallUpdateFailure", types.YLeaf{"CsbCurrPeriodicStatsTotalCallUpdateFailure", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsTotalCallUpdateFailure})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsActiveIpv6Calls", types.YLeaf{"CsbCurrPeriodicStatsActiveIpv6Calls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsActiveIpv6Calls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsActiveEmergencyCalls", types.YLeaf{"CsbCurrPeriodicStatsActiveEmergencyCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsActiveEmergencyCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsActiveE2EmergencyCalls", types.YLeaf{"CsbCurrPeriodicStatsActiveE2EmergencyCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsActiveE2EmergencyCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsImsRxActiveCalls", types.YLeaf{"CsbCurrPeriodicStatsImsRxActiveCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsImsRxActiveCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsImsRxCallSetupFaiures", types.YLeaf{"CsbCurrPeriodicStatsImsRxCallSetupFaiures", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsImsRxCallSetupFaiures})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsImsRxCallRenegotiationAttempts", types.YLeaf{"CsbCurrPeriodicStatsImsRxCallRenegotiationAttempts", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsImsRxCallRenegotiationAttempts})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsImsRxCallRenegotiationFailures", types.YLeaf{"CsbCurrPeriodicStatsImsRxCallRenegotiationFailures", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsImsRxCallRenegotiationFailures})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsAudioTranscodedCalls", types.YLeaf{"CsbCurrPeriodicStatsAudioTranscodedCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsAudioTranscodedCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsFaxTranscodedCalls", types.YLeaf{"CsbCurrPeriodicStatsFaxTranscodedCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsFaxTranscodedCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsRtpDisallowedFailures", types.YLeaf{"CsbCurrPeriodicStatsRtpDisallowedFailures", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsRtpDisallowedFailures})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsSrtpDisallowedFailures", types.YLeaf{"CsbCurrPeriodicStatsSrtpDisallowedFailures", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsSrtpDisallowedFailures})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsNonSrtpCalls", types.YLeaf{"CsbCurrPeriodicStatsNonSrtpCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsNonSrtpCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsSrtpNonIwCalls", types.YLeaf{"CsbCurrPeriodicStatsSrtpNonIwCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsSrtpNonIwCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsSrtpIwCalls", types.YLeaf{"CsbCurrPeriodicStatsSrtpIwCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsSrtpIwCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsDtmfIw2833Calls", types.YLeaf{"CsbCurrPeriodicStatsDtmfIw2833Calls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsDtmfIw2833Calls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsDtmfIwInbandCalls", types.YLeaf{"CsbCurrPeriodicStatsDtmfIwInbandCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsDtmfIwInbandCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsDtmfIw2833InbandCalls", types.YLeaf{"CsbCurrPeriodicStatsDtmfIw2833InbandCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsDtmfIw2833InbandCalls})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsTotalTapsRequested", types.YLeaf{"CsbCurrPeriodicStatsTotalTapsRequested", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsTotalTapsRequested})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsTotalTapsSucceeded", types.YLeaf{"CsbCurrPeriodicStatsTotalTapsSucceeded", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsTotalTapsSucceeded})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicStatsCurrentTaps", types.YLeaf{"CsbCurrPeriodicStatsCurrentTaps", csbCurrPeriodicStatsEntry.CsbCurrPeriodicStatsCurrentTaps})
    csbCurrPeriodicStatsEntry.EntityData.Leafs.Append("csbCurrPeriodicIpsecCalls", types.YLeaf{"CsbCurrPeriodicIpsecCalls", csbCurrPeriodicStatsEntry.CsbCurrPeriodicIpsecCalls})

    csbCurrPeriodicStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbCurrPeriodicStatsInterval"}

    return &(csbCurrPeriodicStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable
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
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbHistoryStatsTable. The entries in this table are
    // updated as interval completes in the csbCurrPeriodicStatsTable table and
    // the data is moved from that table to this one. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable_CsbHistoryStatsEntry.
    CsbHistoryStatsEntry []*CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable_CsbHistoryStatsEntry
}

func (csbHistoryStatsTable *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable) GetEntityData() *types.CommonEntityData {
    csbHistoryStatsTable.EntityData.YFilter = csbHistoryStatsTable.YFilter
    csbHistoryStatsTable.EntityData.YangName = "csbHistoryStatsTable"
    csbHistoryStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbHistoryStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbHistoryStatsTable.EntityData.SegmentPath = "csbHistoryStatsTable"
    csbHistoryStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbHistoryStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbHistoryStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbHistoryStatsTable.EntityData.Children = types.NewOrderedMap()
    csbHistoryStatsTable.EntityData.Children.Append("csbHistoryStatsEntry", types.YChild{"CsbHistoryStatsEntry", nil})
    for i := range csbHistoryStatsTable.CsbHistoryStatsEntry {
        csbHistoryStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbHistoryStatsTable.CsbHistoryStatsEntry[i]), types.YChild{"CsbHistoryStatsEntry", csbHistoryStatsTable.CsbHistoryStatsEntry[i]})
    }
    csbHistoryStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbHistoryStatsTable.EntityData.YListKeys = []string {}

    return &(csbHistoryStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable_CsbHistoryStatsEntry
// A conceptual row in the csbHistoryStatsTable. The entries in
// this table are updated as interval completes in the
// csbCurrPeriodicStatsTable table and the data is moved from that
// table to this one.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable_CsbHistoryStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry_CsbCallStatsInstanceIndex
    CsbCallStatsInstanceIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry_CsbCallStatsServiceIndex
    CsbCallStatsServiceIndex interface{}

    // This attribute is a key. This object identifies the interval for which the
    // history statistics information is to be displayed. The interval values can
    // be 5 min, 15 mins, 1 hour , 1 day. This object acts as index for the table.
    // The type is CiscoSbcPeriodicStatsInterval.
    CsbHistoryStatsInterval interface{}

    // This attribute is a key. The platform assigns a number starting with one
    // and increments it each for each new row. When the maximum          number
    // of row is reached the oldest rows are deleted. It is up to the platform to
    // determine the number of entries for each interval. It is recommended that
    // each platform support at least one entry for 5 min, 15 mins, 1 hour and 1
    // day intervals. The type is interface{} with range: 0..4294967295.
    CsbHistoryStatsElements interface{}

    // This object indicates the number of active calls history during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbHistoryStatsActiveCalls interface{}

    // This object indicates the number of total call attempts history during this
    // interval. The type is interface{} with range: 0..4294967295.
    CsbHistoryStatsTotalCallAttempts interface{}

    // This object indicates the number of failed call attempts during this
    // interval. The type is interface{} with range: 0..4294967295.
    CsbHistoryStatsFailedCallAttempts interface{}

    // This object indicates the number of call setup failures due to routing
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallRoutingFailure interface{}

    // This object indicates the number of call setup failures due to resource
    // failures during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallResourceFailure interface{}

    // This object indicates the number of call setup failures due to media
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallMediaFailure interface{}

    // This object indicates the number of call setup failures due to signaling
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsFailSigFailure interface{}

    // This object indicates the number of active call failures during this
    // interval. The type is interface{} with range: 0..4294967295.
    CsbHistoryStatsActiveCallFailure interface{}

    // This object indicates the number of call setup failures due to congestion
    // during this interval. The type is interface{} with range: 0..4294967295.
    CsbHistoryStatsCongestionFailure interface{}

    // This object indicates the number of call setup failures due to some policy
    // violations during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupPolicyFailure interface{}

    // This object indicates the number of call setup failures due to NA policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupNAPolicyFailure interface{}

    // This object indicates the number of call setup failures due to routing
    // policy failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupRoutingPolicyFailure interface{}

    // This object indicates the number of call setup failures due to CAC policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupCACPolicyFailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupCACCallLimitFailure interface{}

    // This object indicates the number of call setup failures due to CAC call
    // rate limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupCACRateLimitFailure interface{}

    // This object indicates the number of call setup failures due to CAC
    // bandwidth limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupCACBandwidthFailure interface{}

    // This object indicates the number of call setup failures due to CAC media
    // limit during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupCACMediaLimitFailure interface{}

    // This object indicates the number of call update failure due to policy
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295.
    CsbHistoryStatsCallSetupCACMediaUpdateFailure interface{}

    // This object indicates the time at the start of the interval when
    // measurements were first collected for this interval in the
    // csbCurrPeriodicStatsTable. The type is string with length: 0..80.
    CsbHistoryStatsTimestamp interface{}

    // The object indicates the number of active transcoded calls during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbHistroyStatsTranscodedCalls interface{}

    // The object indicates the number of active transrated calls during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbHistroyStatsTransratedCalls interface{}

    // This object indicates the total call update failures during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsTotalCallUpdateFailure interface{}

    // This Object indicates the number of calls through SBC that use IPv6
    // signaling.  This statistic totals all calls that traverse an IPv6 adjacency
    // on either or both sides of SBC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsActiveIpv6Calls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis)  during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsActiveEmergencyCalls interface{}

    // This object indicates the number of calls through SBC that have been
    // identified as emergency calls (by Number Analysis) and have used the e2
    // interface to obtain location information for the caller. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsActiveE2EmergencyCalls interface{}

    // This object indicates the total number of active calls which use IMS Rx,
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are calls.
    CsbHistoryStatsImsRxActiveCalls interface{}

    // This object indicates the total call setup failures owing to IMS Rx failure
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are failures.
    CsbHistoryStatsImsRxCallSetupFailures interface{}

    // This object indicates the total call renegotiation attempts using IMS Rx
    // during this interval. The type is interface{} with range: 0..4294967295.
    // Units are attempts.
    CsbHistoryStatsImsRxCallRenegotiationAttempts interface{}

    // This object indicates the total call renegotiation failures owing to IMS Rx
    // failure during this interval. The type is interface{} with range:
    // 0..4294967295. Units are failures.
    CsbHistoryStatsImsRxCallRenegotiationFailures interface{}

    // The number of active audio transcoded calls through this adjacency or
    // account during this interval. The type is interface{} with range:
    // 0..4294967295. Units are calls.
    CsbHistoryStatsAudioTranscodedCalls interface{}

    // This object indicates the the number of active fax transcoded calls through
    // this adjacency or account during this interval. The type is interface{}
    // with range: 0..4294967295. Units are calls.
    CsbHistoryStatsFaxTranscodedCalls interface{}

    // This object indicates the total call setup failures due to RTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    CsbHistoryStatsRtpDisallowedFailures interface{}

    // This object indicates the total call setup failures due to SRTP being
    // proposed when not permitted during this interval. The type is interface{}
    // with range: 0..4294967295. Units are failures.
    CsbHistoryStatsSrtpDisallowedFailures interface{}

    // This object indicates the number of active calls through this adjacency or
    // account which do not use SRTP on any media channels during this interval.
    // The type is interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsNonSrtpCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels that use SRTP but no media
    // channels that provide interworking between RTP and SRTP during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // calls.
    CsbHistoryStatsSrtpNonIwCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account that have one or more media channels that provide interworking
    // between RTP and SRTP during this interval. The type is interface{} with
    // range: 0..4294967295. Units are calls.
    CsbHistoryStatsSrtpIwCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via RFC 2833 during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsDtmfIw2833Calls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in signaling
    // and DTMF in media via inband DTMF tones during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsDtmfIwInbandCalls interface{}

    // This object indicates the number of active calls through this adjacency or
    // account for which DTMF interworking is enabled between DTMF in media via
    // RFC 2833 and DTMF in media via inband DTMF tones during this interval. The
    // type is interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsDtmfIw2833InbandCalls interface{}

    // This object indicates the lawful intercept tap attempts requested within
    // the scope of this query during this interval. The type is interface{} with
    // range: 0..4294967295. Units are attempts.
    CsbHistoryStatsTotalTapsRequested interface{}

    // This object indicates the lawful intercept tap attempts that have been
    // successfully implemented within the scope of this query during this
    // interval. The type is interface{} with range: 0..4294967295. Units are
    // success.
    CsbHistoryStatsTotalTapsSucceeded interface{}

    // This object indicates the Lawful intercept taps currently in place on calls
    // within the scope of this query during this interval. The type is
    // interface{} with range: 0..4294967295. Units are taps.
    CsbHistoryStatsCurrentTaps interface{}

    // The number of active calls on this adjacency or account which are to or
    // from registered subscribers using IPSEC during this interval. The type is
    // interface{} with range: 0..4294967295. Units are calls.
    CsbHistoryStatsIpsecCalls interface{}
}

func (csbHistoryStatsEntry *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbHistoryStatsTable_CsbHistoryStatsEntry) GetEntityData() *types.CommonEntityData {
    csbHistoryStatsEntry.EntityData.YFilter = csbHistoryStatsEntry.YFilter
    csbHistoryStatsEntry.EntityData.YangName = "csbHistoryStatsEntry"
    csbHistoryStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbHistoryStatsEntry.EntityData.ParentYangName = "csbHistoryStatsTable"
    csbHistoryStatsEntry.EntityData.SegmentPath = "csbHistoryStatsEntry" + types.AddKeyToken(csbHistoryStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbHistoryStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbHistoryStatsEntry.CsbHistoryStatsInterval, "csbHistoryStatsInterval") + types.AddKeyToken(csbHistoryStatsEntry.CsbHistoryStatsElements, "csbHistoryStatsElements")
    csbHistoryStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbHistoryStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbHistoryStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbHistoryStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbHistoryStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbHistoryStatsEntry.CsbCallStatsInstanceIndex})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbHistoryStatsEntry.CsbCallStatsServiceIndex})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsInterval", types.YLeaf{"CsbHistoryStatsInterval", csbHistoryStatsEntry.CsbHistoryStatsInterval})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsElements", types.YLeaf{"CsbHistoryStatsElements", csbHistoryStatsEntry.CsbHistoryStatsElements})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsActiveCalls", types.YLeaf{"CsbHistoryStatsActiveCalls", csbHistoryStatsEntry.CsbHistoryStatsActiveCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsTotalCallAttempts", types.YLeaf{"CsbHistoryStatsTotalCallAttempts", csbHistoryStatsEntry.CsbHistoryStatsTotalCallAttempts})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsFailedCallAttempts", types.YLeaf{"CsbHistoryStatsFailedCallAttempts", csbHistoryStatsEntry.CsbHistoryStatsFailedCallAttempts})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallRoutingFailure", types.YLeaf{"CsbHistoryStatsCallRoutingFailure", csbHistoryStatsEntry.CsbHistoryStatsCallRoutingFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallResourceFailure", types.YLeaf{"CsbHistoryStatsCallResourceFailure", csbHistoryStatsEntry.CsbHistoryStatsCallResourceFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallMediaFailure", types.YLeaf{"CsbHistoryStatsCallMediaFailure", csbHistoryStatsEntry.CsbHistoryStatsCallMediaFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsFailSigFailure", types.YLeaf{"CsbHistoryStatsFailSigFailure", csbHistoryStatsEntry.CsbHistoryStatsFailSigFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsActiveCallFailure", types.YLeaf{"CsbHistoryStatsActiveCallFailure", csbHistoryStatsEntry.CsbHistoryStatsActiveCallFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCongestionFailure", types.YLeaf{"CsbHistoryStatsCongestionFailure", csbHistoryStatsEntry.CsbHistoryStatsCongestionFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupPolicyFailure", types.YLeaf{"CsbHistoryStatsCallSetupPolicyFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupPolicyFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupNAPolicyFailure", types.YLeaf{"CsbHistoryStatsCallSetupNAPolicyFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupNAPolicyFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupRoutingPolicyFailure", types.YLeaf{"CsbHistoryStatsCallSetupRoutingPolicyFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupRoutingPolicyFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupCACPolicyFailure", types.YLeaf{"CsbHistoryStatsCallSetupCACPolicyFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupCACPolicyFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupCACCallLimitFailure", types.YLeaf{"CsbHistoryStatsCallSetupCACCallLimitFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupCACCallLimitFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupCACRateLimitFailure", types.YLeaf{"CsbHistoryStatsCallSetupCACRateLimitFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupCACRateLimitFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupCACBandwidthFailure", types.YLeaf{"CsbHistoryStatsCallSetupCACBandwidthFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupCACBandwidthFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupCACMediaLimitFailure", types.YLeaf{"CsbHistoryStatsCallSetupCACMediaLimitFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupCACMediaLimitFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCallSetupCACMediaUpdateFailure", types.YLeaf{"CsbHistoryStatsCallSetupCACMediaUpdateFailure", csbHistoryStatsEntry.CsbHistoryStatsCallSetupCACMediaUpdateFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsTimestamp", types.YLeaf{"CsbHistoryStatsTimestamp", csbHistoryStatsEntry.CsbHistoryStatsTimestamp})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistroyStatsTranscodedCalls", types.YLeaf{"CsbHistroyStatsTranscodedCalls", csbHistoryStatsEntry.CsbHistroyStatsTranscodedCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistroyStatsTransratedCalls", types.YLeaf{"CsbHistroyStatsTransratedCalls", csbHistoryStatsEntry.CsbHistroyStatsTransratedCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsTotalCallUpdateFailure", types.YLeaf{"CsbHistoryStatsTotalCallUpdateFailure", csbHistoryStatsEntry.CsbHistoryStatsTotalCallUpdateFailure})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsActiveIpv6Calls", types.YLeaf{"CsbHistoryStatsActiveIpv6Calls", csbHistoryStatsEntry.CsbHistoryStatsActiveIpv6Calls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsActiveEmergencyCalls", types.YLeaf{"CsbHistoryStatsActiveEmergencyCalls", csbHistoryStatsEntry.CsbHistoryStatsActiveEmergencyCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsActiveE2EmergencyCalls", types.YLeaf{"CsbHistoryStatsActiveE2EmergencyCalls", csbHistoryStatsEntry.CsbHistoryStatsActiveE2EmergencyCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsImsRxActiveCalls", types.YLeaf{"CsbHistoryStatsImsRxActiveCalls", csbHistoryStatsEntry.CsbHistoryStatsImsRxActiveCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsImsRxCallSetupFailures", types.YLeaf{"CsbHistoryStatsImsRxCallSetupFailures", csbHistoryStatsEntry.CsbHistoryStatsImsRxCallSetupFailures})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsImsRxCallRenegotiationAttempts", types.YLeaf{"CsbHistoryStatsImsRxCallRenegotiationAttempts", csbHistoryStatsEntry.CsbHistoryStatsImsRxCallRenegotiationAttempts})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsImsRxCallRenegotiationFailures", types.YLeaf{"CsbHistoryStatsImsRxCallRenegotiationFailures", csbHistoryStatsEntry.CsbHistoryStatsImsRxCallRenegotiationFailures})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsAudioTranscodedCalls", types.YLeaf{"CsbHistoryStatsAudioTranscodedCalls", csbHistoryStatsEntry.CsbHistoryStatsAudioTranscodedCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsFaxTranscodedCalls", types.YLeaf{"CsbHistoryStatsFaxTranscodedCalls", csbHistoryStatsEntry.CsbHistoryStatsFaxTranscodedCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsRtpDisallowedFailures", types.YLeaf{"CsbHistoryStatsRtpDisallowedFailures", csbHistoryStatsEntry.CsbHistoryStatsRtpDisallowedFailures})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsSrtpDisallowedFailures", types.YLeaf{"CsbHistoryStatsSrtpDisallowedFailures", csbHistoryStatsEntry.CsbHistoryStatsSrtpDisallowedFailures})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsNonSrtpCalls", types.YLeaf{"CsbHistoryStatsNonSrtpCalls", csbHistoryStatsEntry.CsbHistoryStatsNonSrtpCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsSrtpNonIwCalls", types.YLeaf{"CsbHistoryStatsSrtpNonIwCalls", csbHistoryStatsEntry.CsbHistoryStatsSrtpNonIwCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsSrtpIwCalls", types.YLeaf{"CsbHistoryStatsSrtpIwCalls", csbHistoryStatsEntry.CsbHistoryStatsSrtpIwCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsDtmfIw2833Calls", types.YLeaf{"CsbHistoryStatsDtmfIw2833Calls", csbHistoryStatsEntry.CsbHistoryStatsDtmfIw2833Calls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsDtmfIwInbandCalls", types.YLeaf{"CsbHistoryStatsDtmfIwInbandCalls", csbHistoryStatsEntry.CsbHistoryStatsDtmfIwInbandCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsDtmfIw2833InbandCalls", types.YLeaf{"CsbHistoryStatsDtmfIw2833InbandCalls", csbHistoryStatsEntry.CsbHistoryStatsDtmfIw2833InbandCalls})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsTotalTapsRequested", types.YLeaf{"CsbHistoryStatsTotalTapsRequested", csbHistoryStatsEntry.CsbHistoryStatsTotalTapsRequested})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsTotalTapsSucceeded", types.YLeaf{"CsbHistoryStatsTotalTapsSucceeded", csbHistoryStatsEntry.CsbHistoryStatsTotalTapsSucceeded})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsCurrentTaps", types.YLeaf{"CsbHistoryStatsCurrentTaps", csbHistoryStatsEntry.CsbHistoryStatsCurrentTaps})
    csbHistoryStatsEntry.EntityData.Leafs.Append("csbHistoryStatsIpsecCalls", types.YLeaf{"CsbHistoryStatsIpsecCalls", csbHistoryStatsEntry.CsbHistoryStatsIpsecCalls})

    csbHistoryStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbHistoryStatsInterval", "CsbHistoryStatsElements"}

    return &(csbHistoryStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable
// This table describes statistics table for each call flow. The
// indices of the table are virtual media gateway id, gate id,
// falow pair id and side id (indices for side A or side B). The
// other indices of this table are csbCallStatsInstanceIndex
// defined in csbCallStatsInstanceTable and
// csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbPerFlowStatsTable. There is an entry in this
    // table for vdbe Id, gate id, flow pair id and side id. The other indices of
    // this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry.
    CsbPerFlowStatsEntry []*CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry
}

func (csbPerFlowStatsTable *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable) GetEntityData() *types.CommonEntityData {
    csbPerFlowStatsTable.EntityData.YFilter = csbPerFlowStatsTable.YFilter
    csbPerFlowStatsTable.EntityData.YangName = "csbPerFlowStatsTable"
    csbPerFlowStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbPerFlowStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbPerFlowStatsTable.EntityData.SegmentPath = "csbPerFlowStatsTable"
    csbPerFlowStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbPerFlowStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbPerFlowStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbPerFlowStatsTable.EntityData.Children = types.NewOrderedMap()
    csbPerFlowStatsTable.EntityData.Children.Append("csbPerFlowStatsEntry", types.YChild{"CsbPerFlowStatsEntry", nil})
    for i := range csbPerFlowStatsTable.CsbPerFlowStatsEntry {
        csbPerFlowStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbPerFlowStatsTable.CsbPerFlowStatsEntry[i]), types.YChild{"CsbPerFlowStatsEntry", csbPerFlowStatsTable.CsbPerFlowStatsEntry[i]})
    }
    csbPerFlowStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbPerFlowStatsTable.EntityData.YListKeys = []string {}

    return &(csbPerFlowStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry
// An conceptual row in the csbPerFlowStatsTable. There is
// an entry in this table for vdbe Id, gate id, flow pair id and
// side id. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry_CsbCallStatsInstanceIndex
    CsbCallStatsInstanceIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry_CsbCallStatsServiceIndex
    CsbCallStatsServiceIndex interface{}

    // This attribute is a key. This object identifies the virtual media gateway
    // id. This object also acts as an index for the table. The type is
    // interface{} with range: 0..255.
    CsbPerFlowStatsVdbeId interface{}

    // This attribute is a key. This object identifies the gate id. This object
    // also acts as an index for the table. The type is interface{} with range:
    // 0..65535.
    CsbPerFlowStatsGateId interface{}

    // This attribute is a key. This object identifies the flow pair id. This
    // object also acts as an index for the table. The type is interface{} with
    // range: 0..65535.
    CsbPerFlowStatsFlowPairId interface{}

    // This attribute is a key. This object identifies the index corresponding to
    // side of flow pair either side A or side B. This object also acts as an
    // index for the table. The type is CsbPerFlowStatsSideId.
    CsbPerFlowStatsSideId interface{}

    // This object indicates the type of the flow, like media flow, signaling flow
    // etc. The type is CsbPerFlowStatsFlowType.
    CsbPerFlowStatsFlowType interface{}

    // This object indicates the number of RTP packets sent per flow by the SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CsbPerFlowStatsRTPPktsSent interface{}

    // This object indicates the number of RTP packets received per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CsbPerFlowStatsRTPPktsRcvd interface{}

    // This object indicates the number of RTP packets discarded  per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CsbPerFlowStatsRTPPktsDiscard interface{}

    // This object indicates the number of RTP octets sent per flow by the SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // octets.
    CsbPerFlowStatsRTPOctetsSent interface{}

    // This object indicates the number of RTP octets received per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // octets.
    CsbPerFlowStatsRTPOctetsRcvd interface{}

    // This object indicates the number of RTP octets discarded per flow by the
    // SBC. The type is interface{} with range: 0..18446744073709551615. Units are
    // octets.
    CsbPerFlowStatsRTPOctetsDiscard interface{}

    // The number of RTP packets sent by the remote end point to this MG on this
    // flow. Comparing this with the local number of RTP packets received from the
    // remote end point gives an indication of how many incoming  packets were
    // dropped on this leg of the call. This information is from RTCP packet. Not
    // all endpoints report this statistic, if it is not available it will be set
    // to zero. This statistic will not be available for signaling flows. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    CsbPerFlowStatsRTCPPktsSent interface{}

    // The number of RTP packets received by the remote end point from this MG on
    // this flow. Comparing this with the local number of RTP packets sent from
    // this MG to the remote endpoint gives an indication of how many outgoing
    // packets were dropped on this leg of the call. This information is from RTCP
    // packet. Not all endpoints report this statistic, if it is not available it
    // will be set to zero. This statistic will not be available for signaling
    // flows. The type is interface{} with range: 0..18446744073709551615. Units
    // are packets.
    CsbPerFlowStatsRTCPPktsRcvd interface{}

    // The number of RTP packets reported as lost by the remote end point on this
    // flow. This information is from RTCP packet. Not all endpoints report this
    // statistic, if it is not available it will be set to zero. This statistic
    // will not be available for signaling flows. The type is interface{} with
    // range: 0..18446744073709551615. Units are packets.
    CsbPerFlowStatsRTCPPktsLost interface{}

    // This object indicates the End Point jitter per flow in the SBC. The type is
    // interface{} with range: 0..18446744073709551615. Units are milliseconds.
    CsbPerFlowStatsEPJitter interface{}

    // This object indicates the maximum burst size for the current FlowPair. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    CsbPerFlowStatsTmanPerMbs interface{}

    // This object indicates the bandwidth reserved for flow in kilobytes/second.
    // The type is interface{} with range: 0..4294967295. Units are kilobytes per
    // second.
    CsbPerFlowStatsTmanPerSdr interface{}

    // This object indicates the mark packets sent for the current FlowPair with,
    // or zero if none set. The DSCP is a 6-bit value, which will be present in
    // the top 6 bits of the lowest byte of this field. The type is string.
    CsbPerFlowStatsDscpSettings interface{}

    // This object indicates whether the flow on the current FlowPair has
    // subscribed for the NAT latch event. The type is string with length: 0..10.
    CsbPerFlowStatsAdrStatus interface{}

    // This object indicates the flow on the current FlowPair has subscribed for
    // the media loss event. The type is string with length: 0..10.
    CsbPerFlowStatsQASettings interface{}

    // This object indicates the number of RTP packets lost per flow by the SBC.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CsbPerFlowStatsRTPPktsLost interface{}
}

func (csbPerFlowStatsEntry *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry) GetEntityData() *types.CommonEntityData {
    csbPerFlowStatsEntry.EntityData.YFilter = csbPerFlowStatsEntry.YFilter
    csbPerFlowStatsEntry.EntityData.YangName = "csbPerFlowStatsEntry"
    csbPerFlowStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbPerFlowStatsEntry.EntityData.ParentYangName = "csbPerFlowStatsTable"
    csbPerFlowStatsEntry.EntityData.SegmentPath = "csbPerFlowStatsEntry" + types.AddKeyToken(csbPerFlowStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbPerFlowStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbPerFlowStatsEntry.CsbPerFlowStatsVdbeId, "csbPerFlowStatsVdbeId") + types.AddKeyToken(csbPerFlowStatsEntry.CsbPerFlowStatsGateId, "csbPerFlowStatsGateId") + types.AddKeyToken(csbPerFlowStatsEntry.CsbPerFlowStatsFlowPairId, "csbPerFlowStatsFlowPairId") + types.AddKeyToken(csbPerFlowStatsEntry.CsbPerFlowStatsSideId, "csbPerFlowStatsSideId")
    csbPerFlowStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbPerFlowStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbPerFlowStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbPerFlowStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbPerFlowStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbPerFlowStatsEntry.CsbCallStatsInstanceIndex})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbPerFlowStatsEntry.CsbCallStatsServiceIndex})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsVdbeId", types.YLeaf{"CsbPerFlowStatsVdbeId", csbPerFlowStatsEntry.CsbPerFlowStatsVdbeId})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsGateId", types.YLeaf{"CsbPerFlowStatsGateId", csbPerFlowStatsEntry.CsbPerFlowStatsGateId})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsFlowPairId", types.YLeaf{"CsbPerFlowStatsFlowPairId", csbPerFlowStatsEntry.CsbPerFlowStatsFlowPairId})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsSideId", types.YLeaf{"CsbPerFlowStatsSideId", csbPerFlowStatsEntry.CsbPerFlowStatsSideId})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsFlowType", types.YLeaf{"CsbPerFlowStatsFlowType", csbPerFlowStatsEntry.CsbPerFlowStatsFlowType})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTPPktsSent", types.YLeaf{"CsbPerFlowStatsRTPPktsSent", csbPerFlowStatsEntry.CsbPerFlowStatsRTPPktsSent})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTPPktsRcvd", types.YLeaf{"CsbPerFlowStatsRTPPktsRcvd", csbPerFlowStatsEntry.CsbPerFlowStatsRTPPktsRcvd})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTPPktsDiscard", types.YLeaf{"CsbPerFlowStatsRTPPktsDiscard", csbPerFlowStatsEntry.CsbPerFlowStatsRTPPktsDiscard})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTPOctetsSent", types.YLeaf{"CsbPerFlowStatsRTPOctetsSent", csbPerFlowStatsEntry.CsbPerFlowStatsRTPOctetsSent})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTPOctetsRcvd", types.YLeaf{"CsbPerFlowStatsRTPOctetsRcvd", csbPerFlowStatsEntry.CsbPerFlowStatsRTPOctetsRcvd})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTPOctetsDiscard", types.YLeaf{"CsbPerFlowStatsRTPOctetsDiscard", csbPerFlowStatsEntry.CsbPerFlowStatsRTPOctetsDiscard})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTCPPktsSent", types.YLeaf{"CsbPerFlowStatsRTCPPktsSent", csbPerFlowStatsEntry.CsbPerFlowStatsRTCPPktsSent})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTCPPktsRcvd", types.YLeaf{"CsbPerFlowStatsRTCPPktsRcvd", csbPerFlowStatsEntry.CsbPerFlowStatsRTCPPktsRcvd})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTCPPktsLost", types.YLeaf{"CsbPerFlowStatsRTCPPktsLost", csbPerFlowStatsEntry.CsbPerFlowStatsRTCPPktsLost})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsEPJitter", types.YLeaf{"CsbPerFlowStatsEPJitter", csbPerFlowStatsEntry.CsbPerFlowStatsEPJitter})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsTmanPerMbs", types.YLeaf{"CsbPerFlowStatsTmanPerMbs", csbPerFlowStatsEntry.CsbPerFlowStatsTmanPerMbs})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsTmanPerSdr", types.YLeaf{"CsbPerFlowStatsTmanPerSdr", csbPerFlowStatsEntry.CsbPerFlowStatsTmanPerSdr})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsDscpSettings", types.YLeaf{"CsbPerFlowStatsDscpSettings", csbPerFlowStatsEntry.CsbPerFlowStatsDscpSettings})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsAdrStatus", types.YLeaf{"CsbPerFlowStatsAdrStatus", csbPerFlowStatsEntry.CsbPerFlowStatsAdrStatus})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsQASettings", types.YLeaf{"CsbPerFlowStatsQASettings", csbPerFlowStatsEntry.CsbPerFlowStatsQASettings})
    csbPerFlowStatsEntry.EntityData.Leafs.Append("csbPerFlowStatsRTPPktsLost", types.YLeaf{"CsbPerFlowStatsRTPPktsLost", csbPerFlowStatsEntry.CsbPerFlowStatsRTPPktsLost})

    csbPerFlowStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbPerFlowStatsVdbeId", "CsbPerFlowStatsGateId", "CsbPerFlowStatsFlowPairId", "CsbPerFlowStatsSideId"}

    return &(csbPerFlowStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsFlowType represents signaling flow etc.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsFlowType string

const (
    CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsFlowType_media CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsFlowType = "media"

    CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsFlowType_signalling CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsFlowType = "signalling"
)

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsSideId represents for the table.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsSideId string

const (
    CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsSideId_sideA CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsSideId = "sideA"

    CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsSideId_sideB CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbPerFlowStatsTable_CsbPerFlowStatsEntry_CsbPerFlowStatsSideId = "sideB"
)

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable
// This table describes the H.248 statistics for SBC. The index of
// the table is service index which corresponds to a particular 
// service configured on the SBC and the index assigned to a
// particular H.248 controller. The other index of this table is
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable.
// This table is replaced by the csbH248StatsRev1Table.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStath248Table. There is an entry in this
    // table for the particular controller by a value of csbH248StatsCtrlrIndex.
    // The other indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable_CsbH248StatsEntry.
    CsbH248StatsEntry []*CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable_CsbH248StatsEntry
}

func (csbH248StatsTable *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable) GetEntityData() *types.CommonEntityData {
    csbH248StatsTable.EntityData.YFilter = csbH248StatsTable.YFilter
    csbH248StatsTable.EntityData.YangName = "csbH248StatsTable"
    csbH248StatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbH248StatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbH248StatsTable.EntityData.SegmentPath = "csbH248StatsTable"
    csbH248StatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbH248StatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbH248StatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbH248StatsTable.EntityData.Children = types.NewOrderedMap()
    csbH248StatsTable.EntityData.Children.Append("csbH248StatsEntry", types.YChild{"CsbH248StatsEntry", nil})
    for i := range csbH248StatsTable.CsbH248StatsEntry {
        csbH248StatsTable.EntityData.Children.Append(types.GetSegmentPath(csbH248StatsTable.CsbH248StatsEntry[i]), types.YChild{"CsbH248StatsEntry", csbH248StatsTable.CsbH248StatsEntry[i]})
    }
    csbH248StatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbH248StatsTable.EntityData.YListKeys = []string {}

    return &(csbH248StatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable_CsbH248StatsEntry
// An conceptual row in the csbCallStath248Table. There is
// an entry in this table for the particular controller by a value
// of csbH248StatsCtrlrIndex. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable_CsbH248StatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry_CsbCallStatsInstanceIndex
    CsbCallStatsInstanceIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry_CsbCallStatsServiceIndex
    CsbCallStatsServiceIndex interface{}

    // This attribute is a key. This object identifies the controller index of the
    // H.248 server. This is also the index for the table. The type is interface{}
    // with range: 1..50.
    CsbH248StatsCtrlrIndex interface{}

    // This object indicates the requests sent through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRequestsSent interface{}

    // This object indicates the requests received through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRequestsRcvd interface{}

    // This object indicates the requests failed on session Controller Interface
    // to an SBE or DBE. The type is interface{} with range: 0..4294967295.
    CsbH248StatsRequestsFailed interface{}

    // This object indicates the requests retried through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRequestsRetried interface{}

    // This object indicates the number of replies sent through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRepliesSent interface{}

    // This object indicates the number of replies received from the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRepliesRcvd interface{}

    // This object indicates the number of replies retried through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRepliesRetried interface{}

    // This object indicates the number of packets sent through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsSegPktsSent interface{}

    // This object indicates the number of packets received from the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsSegPktsRcvd interface{}

    // This object indicates the H.248 Controller established time (the time at
    // which the association became established). The type is string with length:
    // 0..80.
    CsbH248StatsEstablishedTime interface{}

    // This object indicates the T-Max timeout value. This field specifies the
    // maximum delay (in milliseconds) for a response from an MGC before deciding
    // that the request has failed. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    CsbH248StatsTMaxTimeoutVal interface{}

    // This object indicates the calculated RTT value. This field specifies the
    // maximum round trip propagation delay in the  network (in milliseconds). The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    CsbH248StatsRTT interface{}

    // This object indicates the LT value calculated from RTT value and Max
    // timeout value. This field specifies the maximum delay (in milliseconds) for
    // a response from an MGC plus the maximum round trip propagation delay in the
    // network (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CsbH248StatsLT interface{}
}

func (csbH248StatsEntry *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsTable_CsbH248StatsEntry) GetEntityData() *types.CommonEntityData {
    csbH248StatsEntry.EntityData.YFilter = csbH248StatsEntry.YFilter
    csbH248StatsEntry.EntityData.YangName = "csbH248StatsEntry"
    csbH248StatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbH248StatsEntry.EntityData.ParentYangName = "csbH248StatsTable"
    csbH248StatsEntry.EntityData.SegmentPath = "csbH248StatsEntry" + types.AddKeyToken(csbH248StatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbH248StatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbH248StatsEntry.CsbH248StatsCtrlrIndex, "csbH248StatsCtrlrIndex")
    csbH248StatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbH248StatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbH248StatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbH248StatsEntry.EntityData.Children = types.NewOrderedMap()
    csbH248StatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbH248StatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbH248StatsEntry.CsbCallStatsInstanceIndex})
    csbH248StatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbH248StatsEntry.CsbCallStatsServiceIndex})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsCtrlrIndex", types.YLeaf{"CsbH248StatsCtrlrIndex", csbH248StatsEntry.CsbH248StatsCtrlrIndex})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRequestsSent", types.YLeaf{"CsbH248StatsRequestsSent", csbH248StatsEntry.CsbH248StatsRequestsSent})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRequestsRcvd", types.YLeaf{"CsbH248StatsRequestsRcvd", csbH248StatsEntry.CsbH248StatsRequestsRcvd})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRequestsFailed", types.YLeaf{"CsbH248StatsRequestsFailed", csbH248StatsEntry.CsbH248StatsRequestsFailed})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRequestsRetried", types.YLeaf{"CsbH248StatsRequestsRetried", csbH248StatsEntry.CsbH248StatsRequestsRetried})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRepliesSent", types.YLeaf{"CsbH248StatsRepliesSent", csbH248StatsEntry.CsbH248StatsRepliesSent})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRepliesRcvd", types.YLeaf{"CsbH248StatsRepliesRcvd", csbH248StatsEntry.CsbH248StatsRepliesRcvd})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRepliesRetried", types.YLeaf{"CsbH248StatsRepliesRetried", csbH248StatsEntry.CsbH248StatsRepliesRetried})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsSegPktsSent", types.YLeaf{"CsbH248StatsSegPktsSent", csbH248StatsEntry.CsbH248StatsSegPktsSent})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsSegPktsRcvd", types.YLeaf{"CsbH248StatsSegPktsRcvd", csbH248StatsEntry.CsbH248StatsSegPktsRcvd})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsEstablishedTime", types.YLeaf{"CsbH248StatsEstablishedTime", csbH248StatsEntry.CsbH248StatsEstablishedTime})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsTMaxTimeoutVal", types.YLeaf{"CsbH248StatsTMaxTimeoutVal", csbH248StatsEntry.CsbH248StatsTMaxTimeoutVal})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsRTT", types.YLeaf{"CsbH248StatsRTT", csbH248StatsEntry.CsbH248StatsRTT})
    csbH248StatsEntry.EntityData.Leafs.Append("csbH248StatsLT", types.YLeaf{"CsbH248StatsLT", csbH248StatsEntry.CsbH248StatsLT})

    csbH248StatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbH248StatsCtrlrIndex"}

    return &(csbH248StatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table
// This table describes the H.248 statistics for SBC. The index of
// the table is service index which corresponds to a particular 
// service configured on the SBC and the index assigned to a
// particular H.248 controller. The other index of this table is
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStath248Table. There is an entry in this
    // table for the particular Vdbe by a value of csbH248StatsVdbeId. The other
    // indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table_CsbH248StatsRev1Entry.
    CsbH248StatsRev1Entry []*CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table_CsbH248StatsRev1Entry
}

func (csbH248StatsRev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table) GetEntityData() *types.CommonEntityData {
    csbH248StatsRev1Table.EntityData.YFilter = csbH248StatsRev1Table.YFilter
    csbH248StatsRev1Table.EntityData.YangName = "csbH248StatsRev1Table"
    csbH248StatsRev1Table.EntityData.BundleName = "cisco_ios_xe"
    csbH248StatsRev1Table.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
    csbH248StatsRev1Table.EntityData.SegmentPath = "csbH248StatsRev1Table"
    csbH248StatsRev1Table.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbH248StatsRev1Table.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbH248StatsRev1Table.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbH248StatsRev1Table.EntityData.Children = types.NewOrderedMap()
    csbH248StatsRev1Table.EntityData.Children.Append("csbH248StatsRev1Entry", types.YChild{"CsbH248StatsRev1Entry", nil})
    for i := range csbH248StatsRev1Table.CsbH248StatsRev1Entry {
        csbH248StatsRev1Table.EntityData.Children.Append(types.GetSegmentPath(csbH248StatsRev1Table.CsbH248StatsRev1Entry[i]), types.YChild{"CsbH248StatsRev1Entry", csbH248StatsRev1Table.CsbH248StatsRev1Entry[i]})
    }
    csbH248StatsRev1Table.EntityData.Leafs = types.NewOrderedMap()

    csbH248StatsRev1Table.EntityData.YListKeys = []string {}

    return &(csbH248StatsRev1Table.EntityData)
}

// CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table_CsbH248StatsRev1Entry
// An conceptual row in the csbCallStath248Table. There is
// an entry in this table for the particular Vdbe by a value
// of csbH248StatsVdbeId. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table_CsbH248StatsRev1Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsInstanceTable_CsbCallStatsInstanceEntry_CsbCallStatsInstanceIndex
    CsbCallStatsInstanceIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_sess_border_ctrlr_call_stats_mib.CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbCallStatsTable_CsbCallStatsEntry_CsbCallStatsServiceIndex
    CsbCallStatsServiceIndex interface{}

    // This attribute is a key. This object identifies the virtual media gateway
    // id. This is also the index for the table. The type is interface{} with
    // range: 0..255.
    CsbH248StatsVdbeId interface{}

    // This object indicates the requests sent through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRequestsSentRev1 interface{}

    // This object indicates the requests received through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRequestsRcvdRev1 interface{}

    // This object indicates the requests failed on session Controller Interface
    // to an SBE or DBE. The type is interface{} with range: 0..4294967295.
    CsbH248StatsRequestsFailedRev1 interface{}

    // This object indicates the requests retried through the Session Controller
    // Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRequestsRetriedRev1 interface{}

    // This object indicates the number of replies sent through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRepliesSentRev1 interface{}

    // This object indicates the number of replies received from the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRepliesRcvdRev1 interface{}

    // This object indicates the number of replies retried through the Session
    // Controller Interface to an SBE or DBE. The type is interface{} with range:
    // 0..4294967295.
    CsbH248StatsRepliesRetriedRev1 interface{}

    // This object indicates the number of response segments sent by DBE. This
    // field will only be present if segmentation is enabled on this association.
    // The type is interface{} with range: 0..4294967295.
    CsbH248StatsSegPktsSentRev1 interface{}

    // This object indicates the number of response segments received by DBE. This
    // field will only be present if segmentation is enabled on this association.
    // The type is interface{} with range: 0..4294967295.
    CsbH248StatsSegPktsRcvdRev1 interface{}

    // This object indicates the H.248 Controller established time (the time at
    // which the association became established). The type is string with length:
    // 0..80.
    CsbH248StatsEstablishedTimeRev1 interface{}

    // This object indicates the T-Max timeout value. This field specifies the
    // maximum delay (in milliseconds) for a response from an MGC before deciding
    // that the request has failed. The type is interface{} with range:
    // -2147483648..2147483647. Units are milliseconds.
    CsbH248StatsTMaxTimeoutValRev1 interface{}

    // This object indicates the calculated RTT value. This field specifies the
    // maximum round trip propagation delay in the  network (in milliseconds). The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    CsbH248StatsRTTRev1 interface{}

    // This object indicates the LT value calculated from RTT value and Max
    // timeout value. This field specifies the maximum delay (in milliseconds) for
    // a response from an MGC plus the maximum round trip propagation delay in the
    // network (in milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CsbH248StatsLTRev1 interface{}
}

func (csbH248StatsRev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_CsbH248StatsRev1Table_CsbH248StatsRev1Entry) GetEntityData() *types.CommonEntityData {
    csbH248StatsRev1Entry.EntityData.YFilter = csbH248StatsRev1Entry.YFilter
    csbH248StatsRev1Entry.EntityData.YangName = "csbH248StatsRev1Entry"
    csbH248StatsRev1Entry.EntityData.BundleName = "cisco_ios_xe"
    csbH248StatsRev1Entry.EntityData.ParentYangName = "csbH248StatsRev1Table"
    csbH248StatsRev1Entry.EntityData.SegmentPath = "csbH248StatsRev1Entry" + types.AddKeyToken(csbH248StatsRev1Entry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbH248StatsRev1Entry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbH248StatsRev1Entry.CsbH248StatsVdbeId, "csbH248StatsVdbeId")
    csbH248StatsRev1Entry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbH248StatsRev1Entry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbH248StatsRev1Entry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbH248StatsRev1Entry.EntityData.Children = types.NewOrderedMap()
    csbH248StatsRev1Entry.EntityData.Leafs = types.NewOrderedMap()
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbH248StatsRev1Entry.CsbCallStatsInstanceIndex})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbH248StatsRev1Entry.CsbCallStatsServiceIndex})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsVdbeId", types.YLeaf{"CsbH248StatsVdbeId", csbH248StatsRev1Entry.CsbH248StatsVdbeId})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRequestsSentRev1", types.YLeaf{"CsbH248StatsRequestsSentRev1", csbH248StatsRev1Entry.CsbH248StatsRequestsSentRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRequestsRcvdRev1", types.YLeaf{"CsbH248StatsRequestsRcvdRev1", csbH248StatsRev1Entry.CsbH248StatsRequestsRcvdRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRequestsFailedRev1", types.YLeaf{"CsbH248StatsRequestsFailedRev1", csbH248StatsRev1Entry.CsbH248StatsRequestsFailedRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRequestsRetriedRev1", types.YLeaf{"CsbH248StatsRequestsRetriedRev1", csbH248StatsRev1Entry.CsbH248StatsRequestsRetriedRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRepliesSentRev1", types.YLeaf{"CsbH248StatsRepliesSentRev1", csbH248StatsRev1Entry.CsbH248StatsRepliesSentRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRepliesRcvdRev1", types.YLeaf{"CsbH248StatsRepliesRcvdRev1", csbH248StatsRev1Entry.CsbH248StatsRepliesRcvdRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRepliesRetriedRev1", types.YLeaf{"CsbH248StatsRepliesRetriedRev1", csbH248StatsRev1Entry.CsbH248StatsRepliesRetriedRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsSegPktsSentRev1", types.YLeaf{"CsbH248StatsSegPktsSentRev1", csbH248StatsRev1Entry.CsbH248StatsSegPktsSentRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsSegPktsRcvdRev1", types.YLeaf{"CsbH248StatsSegPktsRcvdRev1", csbH248StatsRev1Entry.CsbH248StatsSegPktsRcvdRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsEstablishedTimeRev1", types.YLeaf{"CsbH248StatsEstablishedTimeRev1", csbH248StatsRev1Entry.CsbH248StatsEstablishedTimeRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsTMaxTimeoutValRev1", types.YLeaf{"CsbH248StatsTMaxTimeoutValRev1", csbH248StatsRev1Entry.CsbH248StatsTMaxTimeoutValRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsRTTRev1", types.YLeaf{"CsbH248StatsRTTRev1", csbH248StatsRev1Entry.CsbH248StatsRTTRev1})
    csbH248StatsRev1Entry.EntityData.Leafs.Append("csbH248StatsLTRev1", types.YLeaf{"CsbH248StatsLTRev1", csbH248StatsRev1Entry.CsbH248StatsLTRev1})

    csbH248StatsRev1Entry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbH248StatsVdbeId"}

    return &(csbH248StatsRev1Entry.EntityData)
}

