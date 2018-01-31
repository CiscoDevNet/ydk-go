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
    parent types.Entity
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

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetFilter() yfilter.YFilter { return cISCOSESSBORDERCTRLRCALLSTATSMIB.YFilter }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) SetFilter(yf yfilter.YFilter) { cISCOSESSBORDERCTRLRCALLSTATSMIB.YFilter = yf }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceTable" { return "Csbcallstatsinstancetable" }
    if yname == "csbCallStatsTable" { return "Csbcallstatstable" }
    if yname == "csbCurrPeriodicStatsTable" { return "Csbcurrperiodicstatstable" }
    if yname == "csbHistoryStatsTable" { return "Csbhistorystatstable" }
    if yname == "csbPerFlowStatsTable" { return "Csbperflowstatstable" }
    if yname == "csbH248StatsTable" { return "Csbh248Statstable" }
    if yname == "csbH248StatsRev1Table" { return "Csbh248Statsrev1Table" }
    return ""
}

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetSegmentPath() string {
    return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB:CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB"
}

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbCallStatsInstanceTable" {
        return &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcallstatsinstancetable
    }
    if childYangName == "csbCallStatsTable" {
        return &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcallstatstable
    }
    if childYangName == "csbCurrPeriodicStatsTable" {
        return &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcurrperiodicstatstable
    }
    if childYangName == "csbHistoryStatsTable" {
        return &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbhistorystatstable
    }
    if childYangName == "csbPerFlowStatsTable" {
        return &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbperflowstatstable
    }
    if childYangName == "csbH248StatsTable" {
        return &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbh248Statstable
    }
    if childYangName == "csbH248StatsRev1Table" {
        return &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbh248Statsrev1Table
    }
    return nil
}

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["csbCallStatsInstanceTable"] = &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcallstatsinstancetable
    children["csbCallStatsTable"] = &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcallstatstable
    children["csbCurrPeriodicStatsTable"] = &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbcurrperiodicstatstable
    children["csbHistoryStatsTable"] = &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbhistorystatstable
    children["csbPerFlowStatsTable"] = &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbperflowstatstable
    children["csbH248StatsTable"] = &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbh248Statstable
    children["csbH248StatsRev1Table"] = &cISCOSESSBORDERCTRLRCALLSTATSMIB.Csbh248Statsrev1Table
    return children
}

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) SetParent(parent types.Entity) { cISCOSESSBORDERCTRLRCALLSTATSMIB.parent = parent }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetParent() types.Entity { return cISCOSESSBORDERCTRLRCALLSTATSMIB.parent }

func (cISCOSESSBORDERCTRLRCALLSTATSMIB *CISCOSESSBORDERCTRLRCALLSTATSMIB) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable
// The call stats instance table contains the physical index for
// each of the physical entity (line card, primary, secondary
// cards). The index of the table is instance index which uniquely
// identifies the physical entity present on the box.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in csbCallStatsInstanceTable. There is an entry in this
    // table for each SBC instance, as identified by a  value of
    // csbCallStatsInstanceIndex. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry.
    Csbcallstatsinstanceentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry
}

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetFilter() yfilter.YFilter { return csbcallstatsinstancetable.YFilter }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) SetFilter(yf yfilter.YFilter) { csbcallstatsinstancetable.YFilter = yf }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceEntry" { return "Csbcallstatsinstanceentry" }
    return ""
}

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetSegmentPath() string {
    return "csbCallStatsInstanceTable"
}

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbCallStatsInstanceEntry" {
        for _, c := range csbcallstatsinstancetable.Csbcallstatsinstanceentry {
            if csbcallstatsinstancetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry{}
        csbcallstatsinstancetable.Csbcallstatsinstanceentry = append(csbcallstatsinstancetable.Csbcallstatsinstanceentry, child)
        return &csbcallstatsinstancetable.Csbcallstatsinstanceentry[len(csbcallstatsinstancetable.Csbcallstatsinstanceentry)-1]
    }
    return nil
}

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbcallstatsinstancetable.Csbcallstatsinstanceentry {
        children[csbcallstatsinstancetable.Csbcallstatsinstanceentry[i].GetSegmentPath()] = &csbcallstatsinstancetable.Csbcallstatsinstanceentry[i]
    }
    return children
}

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetBundleName() string { return "cisco_ios_xe" }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetYangName() string { return "csbCallStatsInstanceTable" }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) SetParent(parent types.Entity) { csbcallstatsinstancetable.parent = parent }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetParent() types.Entity { return csbcallstatsinstancetable.parent }

func (csbcallstatsinstancetable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry
// A conceptual row in csbCallStatsInstanceTable. There is an
// entry in this table for each SBC instance, as identified by a 
// value of csbCallStatsInstanceIndex.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry struct {
    parent types.Entity
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

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetFilter() yfilter.YFilter { return csbcallstatsinstanceentry.YFilter }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) SetFilter(yf yfilter.YFilter) { csbcallstatsinstanceentry.YFilter = yf }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsInstancePhysicalIndex" { return "Csbcallstatsinstancephysicalindex" }
    return ""
}

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetSegmentPath() string {
    return "csbCallStatsInstanceEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbcallstatsinstanceentry.Csbcallstatsinstanceindex) + "']"
}

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbcallstatsinstanceentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsInstancePhysicalIndex"] = csbcallstatsinstanceentry.Csbcallstatsinstancephysicalindex
    return leafs
}

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetYangName() string { return "csbCallStatsInstanceEntry" }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) SetParent(parent types.Entity) { csbcallstatsinstanceentry.parent = parent }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetParent() types.Entity { return csbcallstatsinstanceentry.parent }

func (csbcallstatsinstanceentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatsinstancetable_Csbcallstatsinstanceentry) GetParentYangName() string { return "csbCallStatsInstanceTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStatsGlobalStatsTable. There is an entry in
    // this table for the particular service configured on SBC identified by a
    // value of csbCallStatsInstanceIndex. The other index of this table is
    // csbCallStatsInstanceIndex which is defined in csbCallStatsInstanceTable.
    // The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry.
    Csbcallstatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry
}

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetFilter() yfilter.YFilter { return csbcallstatstable.YFilter }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) SetFilter(yf yfilter.YFilter) { csbcallstatstable.YFilter = yf }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetGoName(yname string) string {
    if yname == "csbCallStatsEntry" { return "Csbcallstatsentry" }
    return ""
}

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetSegmentPath() string {
    return "csbCallStatsTable"
}

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbCallStatsEntry" {
        for _, c := range csbcallstatstable.Csbcallstatsentry {
            if csbcallstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry{}
        csbcallstatstable.Csbcallstatsentry = append(csbcallstatstable.Csbcallstatsentry, child)
        return &csbcallstatstable.Csbcallstatsentry[len(csbcallstatstable.Csbcallstatsentry)-1]
    }
    return nil
}

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbcallstatstable.Csbcallstatsentry {
        children[csbcallstatstable.Csbcallstatsentry[i].GetSegmentPath()] = &csbcallstatstable.Csbcallstatsentry[i]
    }
    return children
}

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetYangName() string { return "csbCallStatsTable" }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) SetParent(parent types.Entity) { csbcallstatstable.parent = parent }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetParent() types.Entity { return csbcallstatstable.parent }

func (csbcallstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry
// An conceptual row in the csbCallStatsGlobalStatsTable. There is
// an entry in this table for the particular service configured on
// SBC identified by a value of csbCallStatsInstanceIndex. The
// other index of this table is csbCallStatsInstanceIndex which is
// defined in csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry struct {
    parent types.Entity
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

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetFilter() yfilter.YFilter { return csbcallstatsentry.YFilter }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) SetFilter(yf yfilter.YFilter) { csbcallstatsentry.YFilter = yf }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbCallStatsSbcName" { return "Csbcallstatssbcname" }
    if yname == "csbCallStatsCallsHigh" { return "Csbcallstatscallshigh" }
    if yname == "csbCallStatsRate1Sec" { return "Csbcallstatsrate1Sec" }
    if yname == "csbCallStatsCallsLow" { return "Csbcallstatscallslow" }
    if yname == "csbCallStatsAvailableFlows" { return "Csbcallstatsavailableflows" }
    if yname == "csbCallStatsUsedFlows" { return "Csbcallstatsusedflows" }
    if yname == "csbCallStatsPeakFlows" { return "Csbcallstatspeakflows" }
    if yname == "csbCallStatsTotalFlows" { return "Csbcallstatstotalflows" }
    if yname == "csbCallStatsUsedSigFlows" { return "Csbcallstatsusedsigflows" }
    if yname == "csbCallStatsPeakSigFlows" { return "Csbcallstatspeaksigflows" }
    if yname == "csbCallStatsTotalSigFlows" { return "Csbcallstatstotalsigflows" }
    if yname == "csbCallStatsAvailablePktRate" { return "Csbcallstatsavailablepktrate" }
    if yname == "csbCallStatsUnclassifiedPkts" { return "Csbcallstatsunclassifiedpkts" }
    if yname == "csbCallStatsRTPPktsSent" { return "Csbcallstatsrtppktssent" }
    if yname == "csbCallStatsRTPPktsRcvd" { return "Csbcallstatsrtppktsrcvd" }
    if yname == "csbCallStatsRTPPktsDiscard" { return "Csbcallstatsrtppktsdiscard" }
    if yname == "csbCallStatsRTPOctetsSent" { return "Csbcallstatsrtpoctetssent" }
    if yname == "csbCallStatsRTPOctetsRcvd" { return "Csbcallstatsrtpoctetsrcvd" }
    if yname == "csbCallStatsRTPOctetsDiscard" { return "Csbcallstatsrtpoctetsdiscard" }
    if yname == "csbCallStatsNoMediaCount" { return "Csbcallstatsnomediacount" }
    if yname == "csbCallStatsRouteErrors" { return "Csbcallstatsrouteerrors" }
    if yname == "csbCallStatsAvailableTranscodeFlows" { return "Csbcallstatsavailabletranscodeflows" }
    if yname == "csbCallStatsActiveTranscodeFlows" { return "Csbcallstatsactivetranscodeflows" }
    if yname == "csbCallStatsPeakTranscodeFlows" { return "Csbcallstatspeaktranscodeflows" }
    if yname == "csbCallStatsTotalTranscodeFlows" { return "Csbcallstatstotaltranscodeflows" }
    return ""
}

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetSegmentPath() string {
    return "csbCallStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbcallstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbcallstatsentry.Csbcallstatsserviceindex) + "']"
}

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbcallstatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbcallstatsentry.Csbcallstatsserviceindex
    leafs["csbCallStatsSbcName"] = csbcallstatsentry.Csbcallstatssbcname
    leafs["csbCallStatsCallsHigh"] = csbcallstatsentry.Csbcallstatscallshigh
    leafs["csbCallStatsRate1Sec"] = csbcallstatsentry.Csbcallstatsrate1Sec
    leafs["csbCallStatsCallsLow"] = csbcallstatsentry.Csbcallstatscallslow
    leafs["csbCallStatsAvailableFlows"] = csbcallstatsentry.Csbcallstatsavailableflows
    leafs["csbCallStatsUsedFlows"] = csbcallstatsentry.Csbcallstatsusedflows
    leafs["csbCallStatsPeakFlows"] = csbcallstatsentry.Csbcallstatspeakflows
    leafs["csbCallStatsTotalFlows"] = csbcallstatsentry.Csbcallstatstotalflows
    leafs["csbCallStatsUsedSigFlows"] = csbcallstatsentry.Csbcallstatsusedsigflows
    leafs["csbCallStatsPeakSigFlows"] = csbcallstatsentry.Csbcallstatspeaksigflows
    leafs["csbCallStatsTotalSigFlows"] = csbcallstatsentry.Csbcallstatstotalsigflows
    leafs["csbCallStatsAvailablePktRate"] = csbcallstatsentry.Csbcallstatsavailablepktrate
    leafs["csbCallStatsUnclassifiedPkts"] = csbcallstatsentry.Csbcallstatsunclassifiedpkts
    leafs["csbCallStatsRTPPktsSent"] = csbcallstatsentry.Csbcallstatsrtppktssent
    leafs["csbCallStatsRTPPktsRcvd"] = csbcallstatsentry.Csbcallstatsrtppktsrcvd
    leafs["csbCallStatsRTPPktsDiscard"] = csbcallstatsentry.Csbcallstatsrtppktsdiscard
    leafs["csbCallStatsRTPOctetsSent"] = csbcallstatsentry.Csbcallstatsrtpoctetssent
    leafs["csbCallStatsRTPOctetsRcvd"] = csbcallstatsentry.Csbcallstatsrtpoctetsrcvd
    leafs["csbCallStatsRTPOctetsDiscard"] = csbcallstatsentry.Csbcallstatsrtpoctetsdiscard
    leafs["csbCallStatsNoMediaCount"] = csbcallstatsentry.Csbcallstatsnomediacount
    leafs["csbCallStatsRouteErrors"] = csbcallstatsentry.Csbcallstatsrouteerrors
    leafs["csbCallStatsAvailableTranscodeFlows"] = csbcallstatsentry.Csbcallstatsavailabletranscodeflows
    leafs["csbCallStatsActiveTranscodeFlows"] = csbcallstatsentry.Csbcallstatsactivetranscodeflows
    leafs["csbCallStatsPeakTranscodeFlows"] = csbcallstatsentry.Csbcallstatspeaktranscodeflows
    leafs["csbCallStatsTotalTranscodeFlows"] = csbcallstatsentry.Csbcallstatstotaltranscodeflows
    return leafs
}

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetYangName() string { return "csbCallStatsEntry" }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) SetParent(parent types.Entity) { csbcallstatsentry.parent = parent }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetParent() types.Entity { return csbcallstatsentry.parent }

func (csbcallstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcallstatstable_Csbcallstatsentry) GetParentYangName() string { return "csbCallStatsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An conceptual row in the csbCurrPeriodicStatsTable. There is an entry in
    // this table for the particular controller by a value of
    // csbH248StatsCtrlrIndex. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry.
    Csbcurrperiodicstatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry
}

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetFilter() yfilter.YFilter { return csbcurrperiodicstatstable.YFilter }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) SetFilter(yf yfilter.YFilter) { csbcurrperiodicstatstable.YFilter = yf }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetGoName(yname string) string {
    if yname == "csbCurrPeriodicStatsEntry" { return "Csbcurrperiodicstatsentry" }
    return ""
}

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetSegmentPath() string {
    return "csbCurrPeriodicStatsTable"
}

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbCurrPeriodicStatsEntry" {
        for _, c := range csbcurrperiodicstatstable.Csbcurrperiodicstatsentry {
            if csbcurrperiodicstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry{}
        csbcurrperiodicstatstable.Csbcurrperiodicstatsentry = append(csbcurrperiodicstatstable.Csbcurrperiodicstatsentry, child)
        return &csbcurrperiodicstatstable.Csbcurrperiodicstatsentry[len(csbcurrperiodicstatstable.Csbcurrperiodicstatsentry)-1]
    }
    return nil
}

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbcurrperiodicstatstable.Csbcurrperiodicstatsentry {
        children[csbcurrperiodicstatstable.Csbcurrperiodicstatsentry[i].GetSegmentPath()] = &csbcurrperiodicstatstable.Csbcurrperiodicstatsentry[i]
    }
    return children
}

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetYangName() string { return "csbCurrPeriodicStatsTable" }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) SetParent(parent types.Entity) { csbcurrperiodicstatstable.parent = parent }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetParent() types.Entity { return csbcurrperiodicstatstable.parent }

func (csbcurrperiodicstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry
// An conceptual row in the csbCurrPeriodicStatsTable. There is
// an entry in this table for the particular controller by a value
// of csbH248StatsCtrlrIndex. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry struct {
    parent types.Entity
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

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetFilter() yfilter.YFilter { return csbcurrperiodicstatsentry.YFilter }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) SetFilter(yf yfilter.YFilter) { csbcurrperiodicstatsentry.YFilter = yf }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbCurrPeriodicStatsInterval" { return "Csbcurrperiodicstatsinterval" }
    if yname == "csbCurrPeriodicStatsActiveCalls" { return "Csbcurrperiodicstatsactivecalls" }
    if yname == "csbCurrPeriodicStatsActivatingCalls" { return "Csbcurrperiodicstatsactivatingcalls" }
    if yname == "csbCurrPeriodicStatsDeactivatingCalls" { return "Csbcurrperiodicstatsdeactivatingcalls" }
    if yname == "csbCurrPeriodicStatsTotalCallAttempts" { return "Csbcurrperiodicstatstotalcallattempts" }
    if yname == "csbCurrPeriodicStatsFailedCallAttempts" { return "Csbcurrperiodicstatsfailedcallattempts" }
    if yname == "csbCurrPeriodicStatsCallRoutingFailure" { return "Csbcurrperiodicstatscallroutingfailure" }
    if yname == "csbCurrPeriodicStatsCallResourceFailure" { return "Csbcurrperiodicstatscallresourcefailure" }
    if yname == "csbCurrPeriodicStatsCallMediaFailure" { return "Csbcurrperiodicstatscallmediafailure" }
    if yname == "csbCurrPeriodicStatsCallSigFailure" { return "Csbcurrperiodicstatscallsigfailure" }
    if yname == "csbCurrPeriodicStatsActiveCallFailure" { return "Csbcurrperiodicstatsactivecallfailure" }
    if yname == "csbCurrPeriodicStatsCongestionFailure" { return "Csbcurrperiodicstatscongestionfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupPolicyFailure" { return "Csbcurrperiodicstatscallsetuppolicyfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupNAPolicyFailure" { return "Csbcurrperiodicstatscallsetupnapolicyfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupRoutingPolicyFailure" { return "Csbcurrperiodicstatscallsetuproutingpolicyfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupCACPolicyFailure" { return "Csbcurrperiodicstatscallsetupcacpolicyfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupCACCallLimitFailure" { return "Csbcurrperiodicstatscallsetupcaccalllimitfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupCACRateLimitFailure" { return "Csbcurrperiodicstatscallsetupcacratelimitfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupCACBandwidthFailure" { return "Csbcurrperiodicstatscallsetupcacbandwidthfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupCACMediaLimitFailure" { return "Csbcurrperiodicstatscallsetupcacmedialimitfailure" }
    if yname == "csbCurrPeriodicStatsCallSetupCACMediaUpdateFailure" { return "Csbcurrperiodicstatscallsetupcacmediaupdatefailure" }
    if yname == "csbCurrPeriodicStatsTimestamp" { return "Csbcurrperiodicstatstimestamp" }
    if yname == "csbCurrPeriodicStatsTranscodedCalls" { return "Csbcurrperiodicstatstranscodedcalls" }
    if yname == "csbCurrPeriodicStatsTransratedCalls" { return "Csbcurrperiodicstatstransratedcalls" }
    if yname == "csbCurrPeriodicStatsTotalCallUpdateFailure" { return "Csbcurrperiodicstatstotalcallupdatefailure" }
    if yname == "csbCurrPeriodicStatsActiveIpv6Calls" { return "Csbcurrperiodicstatsactiveipv6Calls" }
    if yname == "csbCurrPeriodicStatsActiveEmergencyCalls" { return "Csbcurrperiodicstatsactiveemergencycalls" }
    if yname == "csbCurrPeriodicStatsActiveE2EmergencyCalls" { return "Csbcurrperiodicstatsactivee2Emergencycalls" }
    if yname == "csbCurrPeriodicStatsImsRxActiveCalls" { return "Csbcurrperiodicstatsimsrxactivecalls" }
    if yname == "csbCurrPeriodicStatsImsRxCallSetupFaiures" { return "Csbcurrperiodicstatsimsrxcallsetupfaiures" }
    if yname == "csbCurrPeriodicStatsImsRxCallRenegotiationAttempts" { return "Csbcurrperiodicstatsimsrxcallrenegotiationattempts" }
    if yname == "csbCurrPeriodicStatsImsRxCallRenegotiationFailures" { return "Csbcurrperiodicstatsimsrxcallrenegotiationfailures" }
    if yname == "csbCurrPeriodicStatsAudioTranscodedCalls" { return "Csbcurrperiodicstatsaudiotranscodedcalls" }
    if yname == "csbCurrPeriodicStatsFaxTranscodedCalls" { return "Csbcurrperiodicstatsfaxtranscodedcalls" }
    if yname == "csbCurrPeriodicStatsRtpDisallowedFailures" { return "Csbcurrperiodicstatsrtpdisallowedfailures" }
    if yname == "csbCurrPeriodicStatsSrtpDisallowedFailures" { return "Csbcurrperiodicstatssrtpdisallowedfailures" }
    if yname == "csbCurrPeriodicStatsNonSrtpCalls" { return "Csbcurrperiodicstatsnonsrtpcalls" }
    if yname == "csbCurrPeriodicStatsSrtpNonIwCalls" { return "Csbcurrperiodicstatssrtpnoniwcalls" }
    if yname == "csbCurrPeriodicStatsSrtpIwCalls" { return "Csbcurrperiodicstatssrtpiwcalls" }
    if yname == "csbCurrPeriodicStatsDtmfIw2833Calls" { return "Csbcurrperiodicstatsdtmfiw2833Calls" }
    if yname == "csbCurrPeriodicStatsDtmfIwInbandCalls" { return "Csbcurrperiodicstatsdtmfiwinbandcalls" }
    if yname == "csbCurrPeriodicStatsDtmfIw2833InbandCalls" { return "Csbcurrperiodicstatsdtmfiw2833Inbandcalls" }
    if yname == "csbCurrPeriodicStatsTotalTapsRequested" { return "Csbcurrperiodicstatstotaltapsrequested" }
    if yname == "csbCurrPeriodicStatsTotalTapsSucceeded" { return "Csbcurrperiodicstatstotaltapssucceeded" }
    if yname == "csbCurrPeriodicStatsCurrentTaps" { return "Csbcurrperiodicstatscurrenttaps" }
    if yname == "csbCurrPeriodicIpsecCalls" { return "Csbcurrperiodicipseccalls" }
    return ""
}

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetSegmentPath() string {
    return "csbCurrPeriodicStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbcurrperiodicstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbcurrperiodicstatsentry.Csbcallstatsserviceindex) + "']" + "[csbCurrPeriodicStatsInterval='" + fmt.Sprintf("%v", csbcurrperiodicstatsentry.Csbcurrperiodicstatsinterval) + "']"
}

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbcurrperiodicstatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbcurrperiodicstatsentry.Csbcallstatsserviceindex
    leafs["csbCurrPeriodicStatsInterval"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsinterval
    leafs["csbCurrPeriodicStatsActiveCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivecalls
    leafs["csbCurrPeriodicStatsActivatingCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivatingcalls
    leafs["csbCurrPeriodicStatsDeactivatingCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsdeactivatingcalls
    leafs["csbCurrPeriodicStatsTotalCallAttempts"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatstotalcallattempts
    leafs["csbCurrPeriodicStatsFailedCallAttempts"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsfailedcallattempts
    leafs["csbCurrPeriodicStatsCallRoutingFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallroutingfailure
    leafs["csbCurrPeriodicStatsCallResourceFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallresourcefailure
    leafs["csbCurrPeriodicStatsCallMediaFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallmediafailure
    leafs["csbCurrPeriodicStatsCallSigFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsigfailure
    leafs["csbCurrPeriodicStatsActiveCallFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivecallfailure
    leafs["csbCurrPeriodicStatsCongestionFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscongestionfailure
    leafs["csbCurrPeriodicStatsCallSetupPolicyFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetuppolicyfailure
    leafs["csbCurrPeriodicStatsCallSetupNAPolicyFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupnapolicyfailure
    leafs["csbCurrPeriodicStatsCallSetupRoutingPolicyFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetuproutingpolicyfailure
    leafs["csbCurrPeriodicStatsCallSetupCACPolicyFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacpolicyfailure
    leafs["csbCurrPeriodicStatsCallSetupCACCallLimitFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcaccalllimitfailure
    leafs["csbCurrPeriodicStatsCallSetupCACRateLimitFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacratelimitfailure
    leafs["csbCurrPeriodicStatsCallSetupCACBandwidthFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacbandwidthfailure
    leafs["csbCurrPeriodicStatsCallSetupCACMediaLimitFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacmedialimitfailure
    leafs["csbCurrPeriodicStatsCallSetupCACMediaUpdateFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscallsetupcacmediaupdatefailure
    leafs["csbCurrPeriodicStatsTimestamp"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatstimestamp
    leafs["csbCurrPeriodicStatsTranscodedCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatstranscodedcalls
    leafs["csbCurrPeriodicStatsTransratedCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatstransratedcalls
    leafs["csbCurrPeriodicStatsTotalCallUpdateFailure"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatstotalcallupdatefailure
    leafs["csbCurrPeriodicStatsActiveIpv6Calls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsactiveipv6Calls
    leafs["csbCurrPeriodicStatsActiveEmergencyCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsactiveemergencycalls
    leafs["csbCurrPeriodicStatsActiveE2EmergencyCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsactivee2Emergencycalls
    leafs["csbCurrPeriodicStatsImsRxActiveCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxactivecalls
    leafs["csbCurrPeriodicStatsImsRxCallSetupFaiures"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxcallsetupfaiures
    leafs["csbCurrPeriodicStatsImsRxCallRenegotiationAttempts"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxcallrenegotiationattempts
    leafs["csbCurrPeriodicStatsImsRxCallRenegotiationFailures"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsimsrxcallrenegotiationfailures
    leafs["csbCurrPeriodicStatsAudioTranscodedCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsaudiotranscodedcalls
    leafs["csbCurrPeriodicStatsFaxTranscodedCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsfaxtranscodedcalls
    leafs["csbCurrPeriodicStatsRtpDisallowedFailures"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsrtpdisallowedfailures
    leafs["csbCurrPeriodicStatsSrtpDisallowedFailures"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatssrtpdisallowedfailures
    leafs["csbCurrPeriodicStatsNonSrtpCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsnonsrtpcalls
    leafs["csbCurrPeriodicStatsSrtpNonIwCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatssrtpnoniwcalls
    leafs["csbCurrPeriodicStatsSrtpIwCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatssrtpiwcalls
    leafs["csbCurrPeriodicStatsDtmfIw2833Calls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsdtmfiw2833Calls
    leafs["csbCurrPeriodicStatsDtmfIwInbandCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsdtmfiwinbandcalls
    leafs["csbCurrPeriodicStatsDtmfIw2833InbandCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatsdtmfiw2833Inbandcalls
    leafs["csbCurrPeriodicStatsTotalTapsRequested"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatstotaltapsrequested
    leafs["csbCurrPeriodicStatsTotalTapsSucceeded"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatstotaltapssucceeded
    leafs["csbCurrPeriodicStatsCurrentTaps"] = csbcurrperiodicstatsentry.Csbcurrperiodicstatscurrenttaps
    leafs["csbCurrPeriodicIpsecCalls"] = csbcurrperiodicstatsentry.Csbcurrperiodicipseccalls
    return leafs
}

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetYangName() string { return "csbCurrPeriodicStatsEntry" }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) SetParent(parent types.Entity) { csbcurrperiodicstatsentry.parent = parent }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetParent() types.Entity { return csbcurrperiodicstatsentry.parent }

func (csbcurrperiodicstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbcurrperiodicstatstable_Csbcurrperiodicstatsentry) GetParentYangName() string { return "csbCurrPeriodicStatsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the csbHistoryStatsTable. The entries in this table are
    // updated as interval completes in the csbCurrPeriodicStatsTable table and
    // the data is moved from that table to this one. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry.
    Csbhistorystatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry
}

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetFilter() yfilter.YFilter { return csbhistorystatstable.YFilter }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) SetFilter(yf yfilter.YFilter) { csbhistorystatstable.YFilter = yf }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetGoName(yname string) string {
    if yname == "csbHistoryStatsEntry" { return "Csbhistorystatsentry" }
    return ""
}

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetSegmentPath() string {
    return "csbHistoryStatsTable"
}

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbHistoryStatsEntry" {
        for _, c := range csbhistorystatstable.Csbhistorystatsentry {
            if csbhistorystatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry{}
        csbhistorystatstable.Csbhistorystatsentry = append(csbhistorystatstable.Csbhistorystatsentry, child)
        return &csbhistorystatstable.Csbhistorystatsentry[len(csbhistorystatstable.Csbhistorystatsentry)-1]
    }
    return nil
}

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbhistorystatstable.Csbhistorystatsentry {
        children[csbhistorystatstable.Csbhistorystatsentry[i].GetSegmentPath()] = &csbhistorystatstable.Csbhistorystatsentry[i]
    }
    return children
}

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetYangName() string { return "csbHistoryStatsTable" }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) SetParent(parent types.Entity) { csbhistorystatstable.parent = parent }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetParent() types.Entity { return csbhistorystatstable.parent }

func (csbhistorystatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry
// A conceptual row in the csbHistoryStatsTable. The entries in
// this table are updated as interval completes in the
// csbCurrPeriodicStatsTable table and the data is moved from that
// table to this one.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry struct {
    parent types.Entity
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

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetFilter() yfilter.YFilter { return csbhistorystatsentry.YFilter }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) SetFilter(yf yfilter.YFilter) { csbhistorystatsentry.YFilter = yf }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbHistoryStatsInterval" { return "Csbhistorystatsinterval" }
    if yname == "csbHistoryStatsElements" { return "Csbhistorystatselements" }
    if yname == "csbHistoryStatsActiveCalls" { return "Csbhistorystatsactivecalls" }
    if yname == "csbHistoryStatsTotalCallAttempts" { return "Csbhistorystatstotalcallattempts" }
    if yname == "csbHistoryStatsFailedCallAttempts" { return "Csbhistorystatsfailedcallattempts" }
    if yname == "csbHistoryStatsCallRoutingFailure" { return "Csbhistorystatscallroutingfailure" }
    if yname == "csbHistoryStatsCallResourceFailure" { return "Csbhistorystatscallresourcefailure" }
    if yname == "csbHistoryStatsCallMediaFailure" { return "Csbhistorystatscallmediafailure" }
    if yname == "csbHistoryStatsFailSigFailure" { return "Csbhistorystatsfailsigfailure" }
    if yname == "csbHistoryStatsActiveCallFailure" { return "Csbhistorystatsactivecallfailure" }
    if yname == "csbHistoryStatsCongestionFailure" { return "Csbhistorystatscongestionfailure" }
    if yname == "csbHistoryStatsCallSetupPolicyFailure" { return "Csbhistorystatscallsetuppolicyfailure" }
    if yname == "csbHistoryStatsCallSetupNAPolicyFailure" { return "Csbhistorystatscallsetupnapolicyfailure" }
    if yname == "csbHistoryStatsCallSetupRoutingPolicyFailure" { return "Csbhistorystatscallsetuproutingpolicyfailure" }
    if yname == "csbHistoryStatsCallSetupCACPolicyFailure" { return "Csbhistorystatscallsetupcacpolicyfailure" }
    if yname == "csbHistoryStatsCallSetupCACCallLimitFailure" { return "Csbhistorystatscallsetupcaccalllimitfailure" }
    if yname == "csbHistoryStatsCallSetupCACRateLimitFailure" { return "Csbhistorystatscallsetupcacratelimitfailure" }
    if yname == "csbHistoryStatsCallSetupCACBandwidthFailure" { return "Csbhistorystatscallsetupcacbandwidthfailure" }
    if yname == "csbHistoryStatsCallSetupCACMediaLimitFailure" { return "Csbhistorystatscallsetupcacmedialimitfailure" }
    if yname == "csbHistoryStatsCallSetupCACMediaUpdateFailure" { return "Csbhistorystatscallsetupcacmediaupdatefailure" }
    if yname == "csbHistoryStatsTimestamp" { return "Csbhistorystatstimestamp" }
    if yname == "csbHistroyStatsTranscodedCalls" { return "Csbhistroystatstranscodedcalls" }
    if yname == "csbHistroyStatsTransratedCalls" { return "Csbhistroystatstransratedcalls" }
    if yname == "csbHistoryStatsTotalCallUpdateFailure" { return "Csbhistorystatstotalcallupdatefailure" }
    if yname == "csbHistoryStatsActiveIpv6Calls" { return "Csbhistorystatsactiveipv6Calls" }
    if yname == "csbHistoryStatsActiveEmergencyCalls" { return "Csbhistorystatsactiveemergencycalls" }
    if yname == "csbHistoryStatsActiveE2EmergencyCalls" { return "Csbhistorystatsactivee2Emergencycalls" }
    if yname == "csbHistoryStatsImsRxActiveCalls" { return "Csbhistorystatsimsrxactivecalls" }
    if yname == "csbHistoryStatsImsRxCallSetupFailures" { return "Csbhistorystatsimsrxcallsetupfailures" }
    if yname == "csbHistoryStatsImsRxCallRenegotiationAttempts" { return "Csbhistorystatsimsrxcallrenegotiationattempts" }
    if yname == "csbHistoryStatsImsRxCallRenegotiationFailures" { return "Csbhistorystatsimsrxcallrenegotiationfailures" }
    if yname == "csbHistoryStatsAudioTranscodedCalls" { return "Csbhistorystatsaudiotranscodedcalls" }
    if yname == "csbHistoryStatsFaxTranscodedCalls" { return "Csbhistorystatsfaxtranscodedcalls" }
    if yname == "csbHistoryStatsRtpDisallowedFailures" { return "Csbhistorystatsrtpdisallowedfailures" }
    if yname == "csbHistoryStatsSrtpDisallowedFailures" { return "Csbhistorystatssrtpdisallowedfailures" }
    if yname == "csbHistoryStatsNonSrtpCalls" { return "Csbhistorystatsnonsrtpcalls" }
    if yname == "csbHistoryStatsSrtpNonIwCalls" { return "Csbhistorystatssrtpnoniwcalls" }
    if yname == "csbHistoryStatsSrtpIwCalls" { return "Csbhistorystatssrtpiwcalls" }
    if yname == "csbHistoryStatsDtmfIw2833Calls" { return "Csbhistorystatsdtmfiw2833Calls" }
    if yname == "csbHistoryStatsDtmfIwInbandCalls" { return "Csbhistorystatsdtmfiwinbandcalls" }
    if yname == "csbHistoryStatsDtmfIw2833InbandCalls" { return "Csbhistorystatsdtmfiw2833Inbandcalls" }
    if yname == "csbHistoryStatsTotalTapsRequested" { return "Csbhistorystatstotaltapsrequested" }
    if yname == "csbHistoryStatsTotalTapsSucceeded" { return "Csbhistorystatstotaltapssucceeded" }
    if yname == "csbHistoryStatsCurrentTaps" { return "Csbhistorystatscurrenttaps" }
    if yname == "csbHistoryStatsIpsecCalls" { return "Csbhistorystatsipseccalls" }
    return ""
}

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetSegmentPath() string {
    return "csbHistoryStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbcallstatsserviceindex) + "']" + "[csbHistoryStatsInterval='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbhistorystatsinterval) + "']" + "[csbHistoryStatsElements='" + fmt.Sprintf("%v", csbhistorystatsentry.Csbhistorystatselements) + "']"
}

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbhistorystatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbhistorystatsentry.Csbcallstatsserviceindex
    leafs["csbHistoryStatsInterval"] = csbhistorystatsentry.Csbhistorystatsinterval
    leafs["csbHistoryStatsElements"] = csbhistorystatsentry.Csbhistorystatselements
    leafs["csbHistoryStatsActiveCalls"] = csbhistorystatsentry.Csbhistorystatsactivecalls
    leafs["csbHistoryStatsTotalCallAttempts"] = csbhistorystatsentry.Csbhistorystatstotalcallattempts
    leafs["csbHistoryStatsFailedCallAttempts"] = csbhistorystatsentry.Csbhistorystatsfailedcallattempts
    leafs["csbHistoryStatsCallRoutingFailure"] = csbhistorystatsentry.Csbhistorystatscallroutingfailure
    leafs["csbHistoryStatsCallResourceFailure"] = csbhistorystatsentry.Csbhistorystatscallresourcefailure
    leafs["csbHistoryStatsCallMediaFailure"] = csbhistorystatsentry.Csbhistorystatscallmediafailure
    leafs["csbHistoryStatsFailSigFailure"] = csbhistorystatsentry.Csbhistorystatsfailsigfailure
    leafs["csbHistoryStatsActiveCallFailure"] = csbhistorystatsentry.Csbhistorystatsactivecallfailure
    leafs["csbHistoryStatsCongestionFailure"] = csbhistorystatsentry.Csbhistorystatscongestionfailure
    leafs["csbHistoryStatsCallSetupPolicyFailure"] = csbhistorystatsentry.Csbhistorystatscallsetuppolicyfailure
    leafs["csbHistoryStatsCallSetupNAPolicyFailure"] = csbhistorystatsentry.Csbhistorystatscallsetupnapolicyfailure
    leafs["csbHistoryStatsCallSetupRoutingPolicyFailure"] = csbhistorystatsentry.Csbhistorystatscallsetuproutingpolicyfailure
    leafs["csbHistoryStatsCallSetupCACPolicyFailure"] = csbhistorystatsentry.Csbhistorystatscallsetupcacpolicyfailure
    leafs["csbHistoryStatsCallSetupCACCallLimitFailure"] = csbhistorystatsentry.Csbhistorystatscallsetupcaccalllimitfailure
    leafs["csbHistoryStatsCallSetupCACRateLimitFailure"] = csbhistorystatsentry.Csbhistorystatscallsetupcacratelimitfailure
    leafs["csbHistoryStatsCallSetupCACBandwidthFailure"] = csbhistorystatsentry.Csbhistorystatscallsetupcacbandwidthfailure
    leafs["csbHistoryStatsCallSetupCACMediaLimitFailure"] = csbhistorystatsentry.Csbhistorystatscallsetupcacmedialimitfailure
    leafs["csbHistoryStatsCallSetupCACMediaUpdateFailure"] = csbhistorystatsentry.Csbhistorystatscallsetupcacmediaupdatefailure
    leafs["csbHistoryStatsTimestamp"] = csbhistorystatsentry.Csbhistorystatstimestamp
    leafs["csbHistroyStatsTranscodedCalls"] = csbhistorystatsentry.Csbhistroystatstranscodedcalls
    leafs["csbHistroyStatsTransratedCalls"] = csbhistorystatsentry.Csbhistroystatstransratedcalls
    leafs["csbHistoryStatsTotalCallUpdateFailure"] = csbhistorystatsentry.Csbhistorystatstotalcallupdatefailure
    leafs["csbHistoryStatsActiveIpv6Calls"] = csbhistorystatsentry.Csbhistorystatsactiveipv6Calls
    leafs["csbHistoryStatsActiveEmergencyCalls"] = csbhistorystatsentry.Csbhistorystatsactiveemergencycalls
    leafs["csbHistoryStatsActiveE2EmergencyCalls"] = csbhistorystatsentry.Csbhistorystatsactivee2Emergencycalls
    leafs["csbHistoryStatsImsRxActiveCalls"] = csbhistorystatsentry.Csbhistorystatsimsrxactivecalls
    leafs["csbHistoryStatsImsRxCallSetupFailures"] = csbhistorystatsentry.Csbhistorystatsimsrxcallsetupfailures
    leafs["csbHistoryStatsImsRxCallRenegotiationAttempts"] = csbhistorystatsentry.Csbhistorystatsimsrxcallrenegotiationattempts
    leafs["csbHistoryStatsImsRxCallRenegotiationFailures"] = csbhistorystatsentry.Csbhistorystatsimsrxcallrenegotiationfailures
    leafs["csbHistoryStatsAudioTranscodedCalls"] = csbhistorystatsentry.Csbhistorystatsaudiotranscodedcalls
    leafs["csbHistoryStatsFaxTranscodedCalls"] = csbhistorystatsentry.Csbhistorystatsfaxtranscodedcalls
    leafs["csbHistoryStatsRtpDisallowedFailures"] = csbhistorystatsentry.Csbhistorystatsrtpdisallowedfailures
    leafs["csbHistoryStatsSrtpDisallowedFailures"] = csbhistorystatsentry.Csbhistorystatssrtpdisallowedfailures
    leafs["csbHistoryStatsNonSrtpCalls"] = csbhistorystatsentry.Csbhistorystatsnonsrtpcalls
    leafs["csbHistoryStatsSrtpNonIwCalls"] = csbhistorystatsentry.Csbhistorystatssrtpnoniwcalls
    leafs["csbHistoryStatsSrtpIwCalls"] = csbhistorystatsentry.Csbhistorystatssrtpiwcalls
    leafs["csbHistoryStatsDtmfIw2833Calls"] = csbhistorystatsentry.Csbhistorystatsdtmfiw2833Calls
    leafs["csbHistoryStatsDtmfIwInbandCalls"] = csbhistorystatsentry.Csbhistorystatsdtmfiwinbandcalls
    leafs["csbHistoryStatsDtmfIw2833InbandCalls"] = csbhistorystatsentry.Csbhistorystatsdtmfiw2833Inbandcalls
    leafs["csbHistoryStatsTotalTapsRequested"] = csbhistorystatsentry.Csbhistorystatstotaltapsrequested
    leafs["csbHistoryStatsTotalTapsSucceeded"] = csbhistorystatsentry.Csbhistorystatstotaltapssucceeded
    leafs["csbHistoryStatsCurrentTaps"] = csbhistorystatsentry.Csbhistorystatscurrenttaps
    leafs["csbHistoryStatsIpsecCalls"] = csbhistorystatsentry.Csbhistorystatsipseccalls
    return leafs
}

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetYangName() string { return "csbHistoryStatsEntry" }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) SetParent(parent types.Entity) { csbhistorystatsentry.parent = parent }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetParent() types.Entity { return csbhistorystatsentry.parent }

func (csbhistorystatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbhistorystatstable_Csbhistorystatsentry) GetParentYangName() string { return "csbHistoryStatsTable" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable
// This table describes statistics table for each call flow. The
// indices of the table are virtual media gateway id, gate id,
// falow pair id and side id (indices for side A or side B). The
// other indices of this table are csbCallStatsInstanceIndex
// defined in csbCallStatsInstanceTable and
// csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An conceptual row in the csbPerFlowStatsTable. There is an entry in this
    // table for vdbe Id, gate id, flow pair id and side id. The other indices of
    // this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry.
    Csbperflowstatsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry
}

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetFilter() yfilter.YFilter { return csbperflowstatstable.YFilter }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) SetFilter(yf yfilter.YFilter) { csbperflowstatstable.YFilter = yf }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetGoName(yname string) string {
    if yname == "csbPerFlowStatsEntry" { return "Csbperflowstatsentry" }
    return ""
}

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetSegmentPath() string {
    return "csbPerFlowStatsTable"
}

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbPerFlowStatsEntry" {
        for _, c := range csbperflowstatstable.Csbperflowstatsentry {
            if csbperflowstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry{}
        csbperflowstatstable.Csbperflowstatsentry = append(csbperflowstatstable.Csbperflowstatsentry, child)
        return &csbperflowstatstable.Csbperflowstatsentry[len(csbperflowstatstable.Csbperflowstatsentry)-1]
    }
    return nil
}

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbperflowstatstable.Csbperflowstatsentry {
        children[csbperflowstatstable.Csbperflowstatsentry[i].GetSegmentPath()] = &csbperflowstatstable.Csbperflowstatsentry[i]
    }
    return children
}

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetYangName() string { return "csbPerFlowStatsTable" }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) SetParent(parent types.Entity) { csbperflowstatstable.parent = parent }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetParent() types.Entity { return csbperflowstatstable.parent }

func (csbperflowstatstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry
// An conceptual row in the csbPerFlowStatsTable. There is
// an entry in this table for vdbe Id, gate id, flow pair id and
// side id. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry struct {
    parent types.Entity
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

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetFilter() yfilter.YFilter { return csbperflowstatsentry.YFilter }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) SetFilter(yf yfilter.YFilter) { csbperflowstatsentry.YFilter = yf }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbPerFlowStatsVdbeId" { return "Csbperflowstatsvdbeid" }
    if yname == "csbPerFlowStatsGateId" { return "Csbperflowstatsgateid" }
    if yname == "csbPerFlowStatsFlowPairId" { return "Csbperflowstatsflowpairid" }
    if yname == "csbPerFlowStatsSideId" { return "Csbperflowstatssideid" }
    if yname == "csbPerFlowStatsFlowType" { return "Csbperflowstatsflowtype" }
    if yname == "csbPerFlowStatsRTPPktsSent" { return "Csbperflowstatsrtppktssent" }
    if yname == "csbPerFlowStatsRTPPktsRcvd" { return "Csbperflowstatsrtppktsrcvd" }
    if yname == "csbPerFlowStatsRTPPktsDiscard" { return "Csbperflowstatsrtppktsdiscard" }
    if yname == "csbPerFlowStatsRTPOctetsSent" { return "Csbperflowstatsrtpoctetssent" }
    if yname == "csbPerFlowStatsRTPOctetsRcvd" { return "Csbperflowstatsrtpoctetsrcvd" }
    if yname == "csbPerFlowStatsRTPOctetsDiscard" { return "Csbperflowstatsrtpoctetsdiscard" }
    if yname == "csbPerFlowStatsRTCPPktsSent" { return "Csbperflowstatsrtcppktssent" }
    if yname == "csbPerFlowStatsRTCPPktsRcvd" { return "Csbperflowstatsrtcppktsrcvd" }
    if yname == "csbPerFlowStatsRTCPPktsLost" { return "Csbperflowstatsrtcppktslost" }
    if yname == "csbPerFlowStatsEPJitter" { return "Csbperflowstatsepjitter" }
    if yname == "csbPerFlowStatsTmanPerMbs" { return "Csbperflowstatstmanpermbs" }
    if yname == "csbPerFlowStatsTmanPerSdr" { return "Csbperflowstatstmanpersdr" }
    if yname == "csbPerFlowStatsDscpSettings" { return "Csbperflowstatsdscpsettings" }
    if yname == "csbPerFlowStatsAdrStatus" { return "Csbperflowstatsadrstatus" }
    if yname == "csbPerFlowStatsQASettings" { return "Csbperflowstatsqasettings" }
    if yname == "csbPerFlowStatsRTPPktsLost" { return "Csbperflowstatsrtppktslost" }
    return ""
}

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetSegmentPath() string {
    return "csbPerFlowStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbcallstatsserviceindex) + "']" + "[csbPerFlowStatsVdbeId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatsvdbeid) + "']" + "[csbPerFlowStatsGateId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatsgateid) + "']" + "[csbPerFlowStatsFlowPairId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatsflowpairid) + "']" + "[csbPerFlowStatsSideId='" + fmt.Sprintf("%v", csbperflowstatsentry.Csbperflowstatssideid) + "']"
}

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbperflowstatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbperflowstatsentry.Csbcallstatsserviceindex
    leafs["csbPerFlowStatsVdbeId"] = csbperflowstatsentry.Csbperflowstatsvdbeid
    leafs["csbPerFlowStatsGateId"] = csbperflowstatsentry.Csbperflowstatsgateid
    leafs["csbPerFlowStatsFlowPairId"] = csbperflowstatsentry.Csbperflowstatsflowpairid
    leafs["csbPerFlowStatsSideId"] = csbperflowstatsentry.Csbperflowstatssideid
    leafs["csbPerFlowStatsFlowType"] = csbperflowstatsentry.Csbperflowstatsflowtype
    leafs["csbPerFlowStatsRTPPktsSent"] = csbperflowstatsentry.Csbperflowstatsrtppktssent
    leafs["csbPerFlowStatsRTPPktsRcvd"] = csbperflowstatsentry.Csbperflowstatsrtppktsrcvd
    leafs["csbPerFlowStatsRTPPktsDiscard"] = csbperflowstatsentry.Csbperflowstatsrtppktsdiscard
    leafs["csbPerFlowStatsRTPOctetsSent"] = csbperflowstatsentry.Csbperflowstatsrtpoctetssent
    leafs["csbPerFlowStatsRTPOctetsRcvd"] = csbperflowstatsentry.Csbperflowstatsrtpoctetsrcvd
    leafs["csbPerFlowStatsRTPOctetsDiscard"] = csbperflowstatsentry.Csbperflowstatsrtpoctetsdiscard
    leafs["csbPerFlowStatsRTCPPktsSent"] = csbperflowstatsentry.Csbperflowstatsrtcppktssent
    leafs["csbPerFlowStatsRTCPPktsRcvd"] = csbperflowstatsentry.Csbperflowstatsrtcppktsrcvd
    leafs["csbPerFlowStatsRTCPPktsLost"] = csbperflowstatsentry.Csbperflowstatsrtcppktslost
    leafs["csbPerFlowStatsEPJitter"] = csbperflowstatsentry.Csbperflowstatsepjitter
    leafs["csbPerFlowStatsTmanPerMbs"] = csbperflowstatsentry.Csbperflowstatstmanpermbs
    leafs["csbPerFlowStatsTmanPerSdr"] = csbperflowstatsentry.Csbperflowstatstmanpersdr
    leafs["csbPerFlowStatsDscpSettings"] = csbperflowstatsentry.Csbperflowstatsdscpsettings
    leafs["csbPerFlowStatsAdrStatus"] = csbperflowstatsentry.Csbperflowstatsadrstatus
    leafs["csbPerFlowStatsQASettings"] = csbperflowstatsentry.Csbperflowstatsqasettings
    leafs["csbPerFlowStatsRTPPktsLost"] = csbperflowstatsentry.Csbperflowstatsrtppktslost
    return leafs
}

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetYangName() string { return "csbPerFlowStatsEntry" }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) SetParent(parent types.Entity) { csbperflowstatsentry.parent = parent }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetParent() types.Entity { return csbperflowstatsentry.parent }

func (csbperflowstatsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbperflowstatstable_Csbperflowstatsentry) GetParentYangName() string { return "csbPerFlowStatsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStath248Table. There is an entry in this
    // table for the particular controller by a value of csbH248StatsCtrlrIndex.
    // The other indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry.
    Csbh248Statsentry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry
}

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetFilter() yfilter.YFilter { return csbh248Statstable.YFilter }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) SetFilter(yf yfilter.YFilter) { csbh248Statstable.YFilter = yf }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetGoName(yname string) string {
    if yname == "csbH248StatsEntry" { return "Csbh248Statsentry" }
    return ""
}

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetSegmentPath() string {
    return "csbH248StatsTable"
}

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbH248StatsEntry" {
        for _, c := range csbh248Statstable.Csbh248Statsentry {
            if csbh248Statstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry{}
        csbh248Statstable.Csbh248Statsentry = append(csbh248Statstable.Csbh248Statsentry, child)
        return &csbh248Statstable.Csbh248Statsentry[len(csbh248Statstable.Csbh248Statsentry)-1]
    }
    return nil
}

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbh248Statstable.Csbh248Statsentry {
        children[csbh248Statstable.Csbh248Statsentry[i].GetSegmentPath()] = &csbh248Statstable.Csbh248Statsentry[i]
    }
    return children
}

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetYangName() string { return "csbH248StatsTable" }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) SetParent(parent types.Entity) { csbh248Statstable.parent = parent }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetParent() types.Entity { return csbh248Statstable.parent }

func (csbh248Statstable *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry
// An conceptual row in the csbCallStath248Table. There is
// an entry in this table for the particular controller by a value
// of csbH248StatsCtrlrIndex. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry struct {
    parent types.Entity
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

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetFilter() yfilter.YFilter { return csbh248Statsentry.YFilter }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) SetFilter(yf yfilter.YFilter) { csbh248Statsentry.YFilter = yf }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbH248StatsCtrlrIndex" { return "Csbh248Statsctrlrindex" }
    if yname == "csbH248StatsRequestsSent" { return "Csbh248Statsrequestssent" }
    if yname == "csbH248StatsRequestsRcvd" { return "Csbh248Statsrequestsrcvd" }
    if yname == "csbH248StatsRequestsFailed" { return "Csbh248Statsrequestsfailed" }
    if yname == "csbH248StatsRequestsRetried" { return "Csbh248Statsrequestsretried" }
    if yname == "csbH248StatsRepliesSent" { return "Csbh248Statsrepliessent" }
    if yname == "csbH248StatsRepliesRcvd" { return "Csbh248Statsrepliesrcvd" }
    if yname == "csbH248StatsRepliesRetried" { return "Csbh248Statsrepliesretried" }
    if yname == "csbH248StatsSegPktsSent" { return "Csbh248Statssegpktssent" }
    if yname == "csbH248StatsSegPktsRcvd" { return "Csbh248Statssegpktsrcvd" }
    if yname == "csbH248StatsEstablishedTime" { return "Csbh248Statsestablishedtime" }
    if yname == "csbH248StatsTMaxTimeoutVal" { return "Csbh248Statstmaxtimeoutval" }
    if yname == "csbH248StatsRTT" { return "Csbh248Statsrtt" }
    if yname == "csbH248StatsLT" { return "Csbh248Statslt" }
    return ""
}

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetSegmentPath() string {
    return "csbH248StatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbh248Statsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbh248Statsentry.Csbcallstatsserviceindex) + "']" + "[csbH248StatsCtrlrIndex='" + fmt.Sprintf("%v", csbh248Statsentry.Csbh248Statsctrlrindex) + "']"
}

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbh248Statsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbh248Statsentry.Csbcallstatsserviceindex
    leafs["csbH248StatsCtrlrIndex"] = csbh248Statsentry.Csbh248Statsctrlrindex
    leafs["csbH248StatsRequestsSent"] = csbh248Statsentry.Csbh248Statsrequestssent
    leafs["csbH248StatsRequestsRcvd"] = csbh248Statsentry.Csbh248Statsrequestsrcvd
    leafs["csbH248StatsRequestsFailed"] = csbh248Statsentry.Csbh248Statsrequestsfailed
    leafs["csbH248StatsRequestsRetried"] = csbh248Statsentry.Csbh248Statsrequestsretried
    leafs["csbH248StatsRepliesSent"] = csbh248Statsentry.Csbh248Statsrepliessent
    leafs["csbH248StatsRepliesRcvd"] = csbh248Statsentry.Csbh248Statsrepliesrcvd
    leafs["csbH248StatsRepliesRetried"] = csbh248Statsentry.Csbh248Statsrepliesretried
    leafs["csbH248StatsSegPktsSent"] = csbh248Statsentry.Csbh248Statssegpktssent
    leafs["csbH248StatsSegPktsRcvd"] = csbh248Statsentry.Csbh248Statssegpktsrcvd
    leafs["csbH248StatsEstablishedTime"] = csbh248Statsentry.Csbh248Statsestablishedtime
    leafs["csbH248StatsTMaxTimeoutVal"] = csbh248Statsentry.Csbh248Statstmaxtimeoutval
    leafs["csbH248StatsRTT"] = csbh248Statsentry.Csbh248Statsrtt
    leafs["csbH248StatsLT"] = csbh248Statsentry.Csbh248Statslt
    return leafs
}

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetYangName() string { return "csbH248StatsEntry" }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) SetParent(parent types.Entity) { csbh248Statsentry.parent = parent }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetParent() types.Entity { return csbh248Statsentry.parent }

func (csbh248Statsentry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statstable_Csbh248Statsentry) GetParentYangName() string { return "csbH248StatsTable" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table
// This table describes the H.248 statistics for SBC. The index of
// the table is service index which corresponds to a particular 
// service configured on the SBC and the index assigned to a
// particular H.248 controller. The other index of this table is
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An conceptual row in the csbCallStath248Table. There is an entry in this
    // table for the particular Vdbe by a value of csbH248StatsVdbeId. The other
    // indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry.
    Csbh248Statsrev1Entry []CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry
}

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetFilter() yfilter.YFilter { return csbh248Statsrev1Table.YFilter }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) SetFilter(yf yfilter.YFilter) { csbh248Statsrev1Table.YFilter = yf }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetGoName(yname string) string {
    if yname == "csbH248StatsRev1Entry" { return "Csbh248Statsrev1Entry" }
    return ""
}

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetSegmentPath() string {
    return "csbH248StatsRev1Table"
}

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbH248StatsRev1Entry" {
        for _, c := range csbh248Statsrev1Table.Csbh248Statsrev1Entry {
            if csbh248Statsrev1Table.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry{}
        csbh248Statsrev1Table.Csbh248Statsrev1Entry = append(csbh248Statsrev1Table.Csbh248Statsrev1Entry, child)
        return &csbh248Statsrev1Table.Csbh248Statsrev1Entry[len(csbh248Statsrev1Table.Csbh248Statsrev1Entry)-1]
    }
    return nil
}

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbh248Statsrev1Table.Csbh248Statsrev1Entry {
        children[csbh248Statsrev1Table.Csbh248Statsrev1Entry[i].GetSegmentPath()] = &csbh248Statsrev1Table.Csbh248Statsrev1Entry[i]
    }
    return children
}

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetBundleName() string { return "cisco_ios_xe" }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetYangName() string { return "csbH248StatsRev1Table" }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) SetParent(parent types.Entity) { csbh248Statsrev1Table.parent = parent }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetParent() types.Entity { return csbh248Statsrev1Table.parent }

func (csbh248Statsrev1Table *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-CALL-STATS-MIB" }

// CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry
// An conceptual row in the csbCallStath248Table. There is
// an entry in this table for the particular Vdbe by a value
// of csbH248StatsVdbeId. The other indices of this table are
// csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable
// and csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry struct {
    parent types.Entity
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

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetFilter() yfilter.YFilter { return csbh248Statsrev1Entry.YFilter }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) SetFilter(yf yfilter.YFilter) { csbh248Statsrev1Entry.YFilter = yf }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbH248StatsVdbeId" { return "Csbh248Statsvdbeid" }
    if yname == "csbH248StatsRequestsSentRev1" { return "Csbh248Statsrequestssentrev1" }
    if yname == "csbH248StatsRequestsRcvdRev1" { return "Csbh248Statsrequestsrcvdrev1" }
    if yname == "csbH248StatsRequestsFailedRev1" { return "Csbh248Statsrequestsfailedrev1" }
    if yname == "csbH248StatsRequestsRetriedRev1" { return "Csbh248Statsrequestsretriedrev1" }
    if yname == "csbH248StatsRepliesSentRev1" { return "Csbh248Statsrepliessentrev1" }
    if yname == "csbH248StatsRepliesRcvdRev1" { return "Csbh248Statsrepliesrcvdrev1" }
    if yname == "csbH248StatsRepliesRetriedRev1" { return "Csbh248Statsrepliesretriedrev1" }
    if yname == "csbH248StatsSegPktsSentRev1" { return "Csbh248Statssegpktssentrev1" }
    if yname == "csbH248StatsSegPktsRcvdRev1" { return "Csbh248Statssegpktsrcvdrev1" }
    if yname == "csbH248StatsEstablishedTimeRev1" { return "Csbh248Statsestablishedtimerev1" }
    if yname == "csbH248StatsTMaxTimeoutValRev1" { return "Csbh248Statstmaxtimeoutvalrev1" }
    if yname == "csbH248StatsRTTRev1" { return "Csbh248Statsrttrev1" }
    if yname == "csbH248StatsLTRev1" { return "Csbh248Statsltrev1" }
    return ""
}

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetSegmentPath() string {
    return "csbH248StatsRev1Entry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbh248Statsrev1Entry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbh248Statsrev1Entry.Csbcallstatsserviceindex) + "']" + "[csbH248StatsVdbeId='" + fmt.Sprintf("%v", csbh248Statsrev1Entry.Csbh248Statsvdbeid) + "']"
}

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbh248Statsrev1Entry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbh248Statsrev1Entry.Csbcallstatsserviceindex
    leafs["csbH248StatsVdbeId"] = csbh248Statsrev1Entry.Csbh248Statsvdbeid
    leafs["csbH248StatsRequestsSentRev1"] = csbh248Statsrev1Entry.Csbh248Statsrequestssentrev1
    leafs["csbH248StatsRequestsRcvdRev1"] = csbh248Statsrev1Entry.Csbh248Statsrequestsrcvdrev1
    leafs["csbH248StatsRequestsFailedRev1"] = csbh248Statsrev1Entry.Csbh248Statsrequestsfailedrev1
    leafs["csbH248StatsRequestsRetriedRev1"] = csbh248Statsrev1Entry.Csbh248Statsrequestsretriedrev1
    leafs["csbH248StatsRepliesSentRev1"] = csbh248Statsrev1Entry.Csbh248Statsrepliessentrev1
    leafs["csbH248StatsRepliesRcvdRev1"] = csbh248Statsrev1Entry.Csbh248Statsrepliesrcvdrev1
    leafs["csbH248StatsRepliesRetriedRev1"] = csbh248Statsrev1Entry.Csbh248Statsrepliesretriedrev1
    leafs["csbH248StatsSegPktsSentRev1"] = csbh248Statsrev1Entry.Csbh248Statssegpktssentrev1
    leafs["csbH248StatsSegPktsRcvdRev1"] = csbh248Statsrev1Entry.Csbh248Statssegpktsrcvdrev1
    leafs["csbH248StatsEstablishedTimeRev1"] = csbh248Statsrev1Entry.Csbh248Statsestablishedtimerev1
    leafs["csbH248StatsTMaxTimeoutValRev1"] = csbh248Statsrev1Entry.Csbh248Statstmaxtimeoutvalrev1
    leafs["csbH248StatsRTTRev1"] = csbh248Statsrev1Entry.Csbh248Statsrttrev1
    leafs["csbH248StatsLTRev1"] = csbh248Statsrev1Entry.Csbh248Statsltrev1
    return leafs
}

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetBundleName() string { return "cisco_ios_xe" }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetYangName() string { return "csbH248StatsRev1Entry" }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) SetParent(parent types.Entity) { csbh248Statsrev1Entry.parent = parent }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetParent() types.Entity { return csbh248Statsrev1Entry.parent }

func (csbh248Statsrev1Entry *CISCOSESSBORDERCTRLRCALLSTATSMIB_Csbh248Statsrev1Table_Csbh248Statsrev1Entry) GetParentYangName() string { return "csbH248StatsRev1Table" }

