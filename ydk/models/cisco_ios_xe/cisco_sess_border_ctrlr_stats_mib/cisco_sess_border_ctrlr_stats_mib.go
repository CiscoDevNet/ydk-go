// The main purpose of this MIB is to define the statistics
// information for Session Border Controller application. This MIB
// categorizes the statistics information into following types:
// 1. RADIUS Messages Statistics - Represents statistics of
//    various RADIUS messages for RADIUS servers with which the
//    client (SBC) shares a secret.
// 
// 2. Rf Billing Statistics- Represents Rf billing statistics 
//    information which monitors the messages sent per-realm over
//    IMS Rx interface by Rf billing manager(SBC).
// 
// 3. SIP Statistics - Represents SIP requests and various SIP
//    responses on a SIP adjacency in a given interval.
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
// Periodic Statistics - Represents the SBC call statistics 
// information for a particular time interval. E.g. you can 
// specify that you want to retrieve statistics for a summary 
// period of the current or previous 5 minutes, 15 minutes, hour, 
// or day. The statistics for 5 minutes are divided into five 
// minute intervals past the hour - that is, at 0 minutes, 5 
// minutes, 10 minutes... past the hour.  When you retrieve 
// statistics for the current five minute period, you will be 
// given statistics from the start of the interval to the current
// time. When you retrieve statistics for the previous five 
// minutes, you will be given the statistics for the entirety of 
// the previous interval. For example, if it is currently 12:43
// -  the current 5 minute statistics cover 12:40 - 12:43
// 
// -  the previous 5 minute statistics cover 12:35 - 12:40
// 
// The other intervals work similarly.  15 minute statistics are 
// divided into 15 minute intervals past the hour (0 minutes, 15 
// minutes, 30 minutes, 45 minutes).  Hourly statistics are divided
// into intervals on the hour. Daily statistics are divided into
// intervals at 0:00 each day. Therefore, if you retrieve the 
// statistics at 12:43 for each of these intervals, the periods
// covered are as follows.  
// -     current 15 minutes: 12:30 - 12:43
// 
// -     previous 15 minutes: 12:15 - 12:30
// 
// -     current hour: 12:00 - 12:43
// 
// -     last hour: 11:00 - 12:00
// 
// -     current day: 00:00 - 12:43
// 
// -     last day: 00:00 (the day before) - 00:00.
// 
// 
// GLOSSARY
// SBC: Session Border Controller
// 
// CSB: CISCO Session Border Controller
// 
// Adjacency: An adjacency contains the system information to be
// transmitted to next HOP.
// 
// ACR: Accounting Request
// 
// ACA: Accounting Accept
// 
// AVP: Attribute-Value Pairs
// 
// REFERENCES
// 1. CISCO Session Border Controller Documents and FAQ
// http://zed.cisco.com/confluence/display/SBC/SBC
package cisco_sess_border_ctrlr_stats_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_sess_border_ctrlr_stats_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-SESS-BORDER-CTRLR-STATS-MIB CISCO-SESS-BORDER-CTRLR-STATS-MIB}", reflect.TypeOf(CISCOSESSBORDERCTRLRSTATSMIB{}))
    ydk.RegisterEntity("CISCO-SESS-BORDER-CTRLR-STATS-MIB:CISCO-SESS-BORDER-CTRLR-STATS-MIB", reflect.TypeOf(CISCOSESSBORDERCTRLRSTATSMIB{}))
}

// CiscoSbcSIPMethod represents This textual convention represents the various SIP Methods.
type CiscoSbcSIPMethod string

const (
    CiscoSbcSIPMethod_unknown CiscoSbcSIPMethod = "unknown"

    CiscoSbcSIPMethod_ack CiscoSbcSIPMethod = "ack"

    CiscoSbcSIPMethod_bye CiscoSbcSIPMethod = "bye"

    CiscoSbcSIPMethod_cancel CiscoSbcSIPMethod = "cancel"

    CiscoSbcSIPMethod_info CiscoSbcSIPMethod = "info"

    CiscoSbcSIPMethod_invite CiscoSbcSIPMethod = "invite"

    CiscoSbcSIPMethod_message CiscoSbcSIPMethod = "message"

    CiscoSbcSIPMethod_notify CiscoSbcSIPMethod = "notify"

    CiscoSbcSIPMethod_options CiscoSbcSIPMethod = "options"

    CiscoSbcSIPMethod_prack CiscoSbcSIPMethod = "prack"

    CiscoSbcSIPMethod_refer CiscoSbcSIPMethod = "refer"

    CiscoSbcSIPMethod_register CiscoSbcSIPMethod = "register"

    CiscoSbcSIPMethod_subscribe CiscoSbcSIPMethod = "subscribe"

    CiscoSbcSIPMethod_update CiscoSbcSIPMethod = "update"
)

// CiscoSbcRadiusClientType represents This textual convention represents the type of RADIUS client.
type CiscoSbcRadiusClientType string

const (
    CiscoSbcRadiusClientType_authentication CiscoSbcRadiusClientType = "authentication"

    CiscoSbcRadiusClientType_accounting CiscoSbcRadiusClientType = "accounting"
)

// CISCOSESSBORDERCTRLRSTATSMIB
type CISCOSESSBORDERCTRLRSTATSMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table has the reporting statistics of various RADIUS messages for
    // RADIUS servers with which the client (SBC) shares a secret. Each entry in
    // this table is identified by a  value of csbRadiusStatsEntIndex. The other
    // indices of this table are csbCallStatsInstanceIndex defined in 
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable.
    Csbradiusstatstable CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable

    // This table describes the Rf billing statistics information which monitors
    // the messages sent per-realm by Rf billing  manager(SBC). SBC sends Rf
    // billing data using Diameter as a transport protocol. Rf billing uses only
    // ACR and ACA Diameter messages for the transport of billing data. The
    // Accounting-Record-Type AVP on the ACR message labels the type  of the
    // accounting request. The following types are used by Rf billing. 1.   For
    // session-based charging, the types Start (session begins), Interim (session
    // is modified) and Stop (session ends) are used. 2.   For event-based
    // charging, the type Event is used when a chargeable event occurs outside the
    // scope of a session.  Each row of this table is identified by a value of
    // csbRfBillRealmStatsIndex and csbRfBillRealmStatsRealmName. The other
    // indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and  csbCallStatsServiceIndex defined in
    // csbCallStatsTable.
    Csbrfbillrealmstatstable CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable

    // This table reports count of SIP request and various SIP responses  for each
    // SIP method on a SIP adjacency in a given interval. Each entry in this table
    // represents a SIP method, its incoming and outgoing count, individual
    // incoming and outgoing  count of various SIP responses for this method on a
    // SIP adjacency in a given interval. To understand the meaning of  interval
    // please refer <Periodic Statistics> section in  description of
    // ciscoSbcStatsMIB.   This table is indexed on csbSIPMthdCurrentStatsAdjName,
    // csbSIPMthdCurrentStatsMethod and  csbSIPMthdCurrentStatsInterval. The other
    // indices of this table are csbCallStatsInstanceIndex defined in 
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable.
    Csbsipmthdcurrentstatstable CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable

    // This table provide historical count of SIP request and various SIP
    // responses for each SIP method on a SIP adjacency in various interval length
    // defined by the csbSIPMthdHistoryStatsInterval object. Each entry in this
    // table represents a SIP method, its incoming and outgoing count, individual
    // incoming and outgoing  count of various SIP responses for this method on a
    // SIP adjacency in a given interval. The possible values of interval will be
    // previous 5 minutes, previous 15 minutes, previous 1 hour and previous day.
    // To understand the meaning of interval please refer <Periodic Statistics>
    // description of ciscoSbcStatsMIB. This table is indexed on
    // csbSIPMthdHistoryStatsAdjName, csbSIPMthdHistoryStatsMethod and 
    // csbSIPMthdHistoryStatsInterval. The other indices of this table are
    // csbCallStatsInstanceIndex defined in  csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable.
    Csbsipmthdhistorystatstable CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable

    // This table reports SIP method request and response code statistics for each
    // method and response code combination on given SIP adjacency in a given
    // interval. To understand the  meaning of interval please refer <Periodic
    // Statistics> section in description of ciscoSbcStatsMIB. An exact lookup
    // will return a row only  if - 1) detailed response code statistics are
    // turned on in SBC 2) response code  messages sent or received is non zero
    // for     given SIP adjacency, method and interval. Also an inexact lookup
    // will only return rows for messages with non-zero counts, to protect the
    // user from large numbers of rows  for response codes which have not been
    // received or sent.
    Csbsipmthdrccurrentstatstable CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable

    // This table reports historical data for SIP method request and response code
    // statistics for each method and response code  combination in a given past
    // interval. The possible values of  interval will be previous 5 minutes,
    // previous 15 minutes,  previous 1 hour and previous day. To understand the 
    // meaning of interval please refer <Periodic Statistics> section in
    // description of ciscoSbcStatsMIB. An exact lookup will return a row only  if
    // - 1) detailed response code statistics are turned on in SBC 2) response
    // code  messages sent or received is non zero for     given SIP adjacency,
    // method and interval. Also an inexact lookup will only return rows for
    // messages with non-zero counts, to protect the user from large numbers of
    // rows for response codes which have not been received or sent.
    Csbsipmthdrchistorystatstable CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable
}

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetFilter() yfilter.YFilter { return cISCOSESSBORDERCTRLRSTATSMIB.YFilter }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) SetFilter(yf yfilter.YFilter) { cISCOSESSBORDERCTRLRSTATSMIB.YFilter = yf }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetGoName(yname string) string {
    if yname == "csbRadiusStatsTable" { return "Csbradiusstatstable" }
    if yname == "csbRfBillRealmStatsTable" { return "Csbrfbillrealmstatstable" }
    if yname == "csbSIPMthdCurrentStatsTable" { return "Csbsipmthdcurrentstatstable" }
    if yname == "csbSIPMthdHistoryStatsTable" { return "Csbsipmthdhistorystatstable" }
    if yname == "csbSIPMthdRCCurrentStatsTable" { return "Csbsipmthdrccurrentstatstable" }
    if yname == "csbSIPMthdRCHistoryStatsTable" { return "Csbsipmthdrchistorystatstable" }
    return ""
}

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetSegmentPath() string {
    return "CISCO-SESS-BORDER-CTRLR-STATS-MIB:CISCO-SESS-BORDER-CTRLR-STATS-MIB"
}

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbRadiusStatsTable" {
        return &cISCOSESSBORDERCTRLRSTATSMIB.Csbradiusstatstable
    }
    if childYangName == "csbRfBillRealmStatsTable" {
        return &cISCOSESSBORDERCTRLRSTATSMIB.Csbrfbillrealmstatstable
    }
    if childYangName == "csbSIPMthdCurrentStatsTable" {
        return &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdcurrentstatstable
    }
    if childYangName == "csbSIPMthdHistoryStatsTable" {
        return &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdhistorystatstable
    }
    if childYangName == "csbSIPMthdRCCurrentStatsTable" {
        return &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdrccurrentstatstable
    }
    if childYangName == "csbSIPMthdRCHistoryStatsTable" {
        return &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdrchistorystatstable
    }
    return nil
}

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["csbRadiusStatsTable"] = &cISCOSESSBORDERCTRLRSTATSMIB.Csbradiusstatstable
    children["csbRfBillRealmStatsTable"] = &cISCOSESSBORDERCTRLRSTATSMIB.Csbrfbillrealmstatstable
    children["csbSIPMthdCurrentStatsTable"] = &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdcurrentstatstable
    children["csbSIPMthdHistoryStatsTable"] = &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdhistorystatstable
    children["csbSIPMthdRCCurrentStatsTable"] = &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdrccurrentstatstable
    children["csbSIPMthdRCHistoryStatsTable"] = &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdrchistorystatstable
    return children
}

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) SetParent(parent types.Entity) { cISCOSESSBORDERCTRLRSTATSMIB.parent = parent }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetParent() types.Entity { return cISCOSESSBORDERCTRLRSTATSMIB.parent }

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable
// This table has the reporting statistics of various RADIUS
// messages for RADIUS servers with which the client (SBC) shares a
// secret. Each entry in this table is identified by a 
// value of csbRadiusStatsEntIndex. The other indices of this table
// are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the csbRadiusStatsTable. There is an entry in this
    // table for each RADIUS server, as identified by a  value of
    // csbRadiusStatsEntIndex. The other indices of this  table are
    // csbCallStatsInstanceIndex defined in  csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry.
    Csbradiusstatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry
}

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetFilter() yfilter.YFilter { return csbradiusstatstable.YFilter }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) SetFilter(yf yfilter.YFilter) { csbradiusstatstable.YFilter = yf }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetGoName(yname string) string {
    if yname == "csbRadiusStatsEntry" { return "Csbradiusstatsentry" }
    return ""
}

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetSegmentPath() string {
    return "csbRadiusStatsTable"
}

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbRadiusStatsEntry" {
        for _, c := range csbradiusstatstable.Csbradiusstatsentry {
            if csbradiusstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry{}
        csbradiusstatstable.Csbradiusstatsentry = append(csbradiusstatstable.Csbradiusstatsentry, child)
        return &csbradiusstatstable.Csbradiusstatsentry[len(csbradiusstatstable.Csbradiusstatsentry)-1]
    }
    return nil
}

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbradiusstatstable.Csbradiusstatsentry {
        children[csbradiusstatstable.Csbradiusstatsentry[i].GetSegmentPath()] = &csbradiusstatstable.Csbradiusstatsentry[i]
    }
    return children
}

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetYangName() string { return "csbRadiusStatsTable" }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) SetParent(parent types.Entity) { csbradiusstatstable.parent = parent }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetParent() types.Entity { return csbradiusstatstable.parent }

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry
// A conceptual row in the csbRadiusStatsTable. There is an
// entry in this table for each RADIUS server, as identified by a 
// value of csbRadiusStatsEntIndex. The other indices of this 
// table are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry struct {
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

    // This attribute is a key. This object indicates the index of the RADIUS
    // client entity that this server is configured on. This index is assigned 
    // arbitrarily by the engine and is not saved over reboots. The type is
    // interface{} with range: 0..4294967295.
    Csbradiusstatsentindex interface{}

    // This object indicates the client name of the RADIUS client to which that
    // these statistics apply. The type is string.
    Csbradiusstatsclientname interface{}

    // This object indicates the type(authentication or accounting) of the RADIUS
    // clients configured on SBC. The type is CiscoSbcRadiusClientType.
    Csbradiusstatsclienttype interface{}

    // This object indicates the server name of the RADIUS server to which that
    // these statistics apply. The type is string.
    Csbradiusstatssrvrname interface{}

    // This object indicates the number of RADIUS Access-Request packets sent to
    // this server.  This does not include retransmissions. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    Csbradiusstatsacsreqs interface{}

    // This object indicates the number of RADIUS Access-Request packets
    // retransmitted to this RADIUS server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Csbradiusstatsacsrtrns interface{}

    // This object indicates the number of RADIUS Access-Accept packets (valid or
    // invalid) received from this server. The type is interface{} with range:
    // 0..18446744073709551615.
    Csbradiusstatsacsaccpts interface{}

    // This object indicates the number of RADIUS Access-Reject packets (valid or
    // invalid) received from this server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Csbradiusstatsacsrejects interface{}

    // This object indicates the number of RADIUS Access-Challenge packets (valid
    // or invalid) received from this server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Csbradiusstatsacschalls interface{}

    // This object indicates the number of RADIUS Accounting-Request packets sent
    // to this server. This does not include retransmissions. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    Csbradiusstatsactreqs interface{}

    // This object indicates the number of RADIUS Accounting-Request packets
    // retransmitted to this RADIUS server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Csbradiusstatsactretrans interface{}

    // This object indicates the number of RADIUS Accounting-Response packets
    // (valid or invalid) received from this server. The type is interface{} with
    // range: 0..18446744073709551615. Units are packets.
    Csbradiusstatsactrsps interface{}

    // This object indicates the number of malformed RADIUS response packets
    // received from this server.  Malformed packets include packets with an
    // invalid length. Bad authenticators, Signature attributes and unknown types
    // are not included as malformed access responses. The type is interface{}
    // with range: 0..18446744073709551615. Units are packets.
    Csbradiusstatsmalformedrsps interface{}

    // This object indicates the number of RADIUS response packets containing
    // invalid authenticators or Signature attributes received from this server.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Csbradiusstatsbadauths interface{}

    // This object indicates the number of RADIUS request packets destined for
    // this server that have not yet timed out or received a response. This
    // variable is incremented when a request is sent and decremented on receipt
    // of the response or on a timeout or retransmission. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    Csbradiusstatspending interface{}

    // This object indicates the number of RADIUS request timeouts to this server.
    // After a timeout the client may retry to a different server or give up. A
    // retry to a different server is counted as a request as well as a timeout.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    Csbradiusstatstimeouts interface{}

    // This object indicates the number of RADIUS packets of unknown type which
    // were received from this server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Csbradiusstatsunknowntype interface{}

    // This object indicates the number of RADIUS packets which were received from
    // this server and dropped for some other reason. The type is interface{} with
    // range: 0..18446744073709551615. Units are packets.
    Csbradiusstatsdropped interface{}
}

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetFilter() yfilter.YFilter { return csbradiusstatsentry.YFilter }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) SetFilter(yf yfilter.YFilter) { csbradiusstatsentry.YFilter = yf }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbRadiusStatsEntIndex" { return "Csbradiusstatsentindex" }
    if yname == "csbRadiusStatsClientName" { return "Csbradiusstatsclientname" }
    if yname == "csbRadiusStatsClientType" { return "Csbradiusstatsclienttype" }
    if yname == "csbRadiusStatsSrvrName" { return "Csbradiusstatssrvrname" }
    if yname == "csbRadiusStatsAcsReqs" { return "Csbradiusstatsacsreqs" }
    if yname == "csbRadiusStatsAcsRtrns" { return "Csbradiusstatsacsrtrns" }
    if yname == "csbRadiusStatsAcsAccpts" { return "Csbradiusstatsacsaccpts" }
    if yname == "csbRadiusStatsAcsRejects" { return "Csbradiusstatsacsrejects" }
    if yname == "csbRadiusStatsAcsChalls" { return "Csbradiusstatsacschalls" }
    if yname == "csbRadiusStatsActReqs" { return "Csbradiusstatsactreqs" }
    if yname == "csbRadiusStatsActRetrans" { return "Csbradiusstatsactretrans" }
    if yname == "csbRadiusStatsActRsps" { return "Csbradiusstatsactrsps" }
    if yname == "csbRadiusStatsMalformedRsps" { return "Csbradiusstatsmalformedrsps" }
    if yname == "csbRadiusStatsBadAuths" { return "Csbradiusstatsbadauths" }
    if yname == "csbRadiusStatsPending" { return "Csbradiusstatspending" }
    if yname == "csbRadiusStatsTimeouts" { return "Csbradiusstatstimeouts" }
    if yname == "csbRadiusStatsUnknownType" { return "Csbradiusstatsunknowntype" }
    if yname == "csbRadiusStatsDropped" { return "Csbradiusstatsdropped" }
    return ""
}

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetSegmentPath() string {
    return "csbRadiusStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbradiusstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbradiusstatsentry.Csbcallstatsserviceindex) + "']" + "[csbRadiusStatsEntIndex='" + fmt.Sprintf("%v", csbradiusstatsentry.Csbradiusstatsentindex) + "']"
}

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbradiusstatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbradiusstatsentry.Csbcallstatsserviceindex
    leafs["csbRadiusStatsEntIndex"] = csbradiusstatsentry.Csbradiusstatsentindex
    leafs["csbRadiusStatsClientName"] = csbradiusstatsentry.Csbradiusstatsclientname
    leafs["csbRadiusStatsClientType"] = csbradiusstatsentry.Csbradiusstatsclienttype
    leafs["csbRadiusStatsSrvrName"] = csbradiusstatsentry.Csbradiusstatssrvrname
    leafs["csbRadiusStatsAcsReqs"] = csbradiusstatsentry.Csbradiusstatsacsreqs
    leafs["csbRadiusStatsAcsRtrns"] = csbradiusstatsentry.Csbradiusstatsacsrtrns
    leafs["csbRadiusStatsAcsAccpts"] = csbradiusstatsentry.Csbradiusstatsacsaccpts
    leafs["csbRadiusStatsAcsRejects"] = csbradiusstatsentry.Csbradiusstatsacsrejects
    leafs["csbRadiusStatsAcsChalls"] = csbradiusstatsentry.Csbradiusstatsacschalls
    leafs["csbRadiusStatsActReqs"] = csbradiusstatsentry.Csbradiusstatsactreqs
    leafs["csbRadiusStatsActRetrans"] = csbradiusstatsentry.Csbradiusstatsactretrans
    leafs["csbRadiusStatsActRsps"] = csbradiusstatsentry.Csbradiusstatsactrsps
    leafs["csbRadiusStatsMalformedRsps"] = csbradiusstatsentry.Csbradiusstatsmalformedrsps
    leafs["csbRadiusStatsBadAuths"] = csbradiusstatsentry.Csbradiusstatsbadauths
    leafs["csbRadiusStatsPending"] = csbradiusstatsentry.Csbradiusstatspending
    leafs["csbRadiusStatsTimeouts"] = csbradiusstatsentry.Csbradiusstatstimeouts
    leafs["csbRadiusStatsUnknownType"] = csbradiusstatsentry.Csbradiusstatsunknowntype
    leafs["csbRadiusStatsDropped"] = csbradiusstatsentry.Csbradiusstatsdropped
    return leafs
}

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetYangName() string { return "csbRadiusStatsEntry" }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) SetParent(parent types.Entity) { csbradiusstatsentry.parent = parent }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetParent() types.Entity { return csbradiusstatsentry.parent }

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetParentYangName() string { return "csbRadiusStatsTable" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable
// This table describes the Rf billing statistics information
// which monitors the messages sent per-realm by Rf billing 
// manager(SBC). SBC sends Rf billing data using Diameter as a
// transport protocol. Rf billing uses only ACR and ACA Diameter
// messages for the transport of billing data. The
// Accounting-Record-Type AVP on the ACR message labels the type 
// of the accounting request. The following types are used by Rf
// billing.
// 1.   For session-based charging, the types Start (session
// begins), Interim (session is modified) and Stop (session ends)
// are used.
// 2.   For event-based charging, the type Event is used when a
// chargeable event occurs outside the scope of a session.
// 
// Each row of this table is identified by a value of
// csbRfBillRealmStatsIndex and csbRfBillRealmStatsRealmName.
// The other indices of this table are csbCallStatsInstanceIndex
// defined in csbCallStatsInstanceTable and 
// csbCallStatsServiceIndex defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the csbRfBillRealmStatsTable. There is an entry in this
    // table for each realm, as identified by a  value of csbRfBillRealmStatsIndex
    // and  csbRfBillRealmStatsRealmName. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry.
    Csbrfbillrealmstatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry
}

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetFilter() yfilter.YFilter { return csbrfbillrealmstatstable.YFilter }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) SetFilter(yf yfilter.YFilter) { csbrfbillrealmstatstable.YFilter = yf }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetGoName(yname string) string {
    if yname == "csbRfBillRealmStatsEntry" { return "Csbrfbillrealmstatsentry" }
    return ""
}

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetSegmentPath() string {
    return "csbRfBillRealmStatsTable"
}

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbRfBillRealmStatsEntry" {
        for _, c := range csbrfbillrealmstatstable.Csbrfbillrealmstatsentry {
            if csbrfbillrealmstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry{}
        csbrfbillrealmstatstable.Csbrfbillrealmstatsentry = append(csbrfbillrealmstatstable.Csbrfbillrealmstatsentry, child)
        return &csbrfbillrealmstatstable.Csbrfbillrealmstatsentry[len(csbrfbillrealmstatstable.Csbrfbillrealmstatsentry)-1]
    }
    return nil
}

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbrfbillrealmstatstable.Csbrfbillrealmstatsentry {
        children[csbrfbillrealmstatstable.Csbrfbillrealmstatsentry[i].GetSegmentPath()] = &csbrfbillrealmstatstable.Csbrfbillrealmstatsentry[i]
    }
    return children
}

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetYangName() string { return "csbRfBillRealmStatsTable" }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) SetParent(parent types.Entity) { csbrfbillrealmstatstable.parent = parent }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetParent() types.Entity { return csbrfbillrealmstatstable.parent }

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry
// A conceptual row in the csbRfBillRealmStatsTable. There
// is an entry in this table for each realm, as identified by a 
// value of csbRfBillRealmStatsIndex and 
// csbRfBillRealmStatsRealmName. The other indices of this
// table are csbCallStatsInstanceIndex defined in
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry struct {
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

    // This attribute is a key. This object indicates the billing method instance
    // index. The range of valid values for this field is 0 - 31. The type is
    // interface{} with range: 0..31.
    Csbrfbillrealmstatsindex interface{}

    // This attribute is a key. This object indicates the realm for which these
    // statistics are collected. The length of this object is zero when value is
    // not assigned to it. The type is string.
    Csbrfbillrealmstatsrealmname interface{}

    // This object indicates the combined sum of successful and failed Start ACRs
    // since start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatstotalstartacrs interface{}

    // This object indicates the combined sum of successful and failed Interim
    // ACRs since start of day or the last time the statistics were reset. The
    // type is interface{} with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatstotalinterimacrs interface{}

    // This object indicates the combined sum of successful and failed Stop ACRs
    // since start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatstotalstopacrs interface{}

    // This object indicates the combined sum of successful and failed Event ACRs
    // since start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatstotaleventacrs interface{}

    // This object indicates the total number of successful Start ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatssuccstartacrs interface{}

    // This object indicates the total number of successful Interim ACRs since
    // start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatssuccinterimacrs interface{}

    // This object indicates the total number of successful Stop ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatssuccstopacrs interface{}

    // This object indicates the total number of successful Event ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatssucceventacrs interface{}

    // This object indicates the total number of failed Start ACRs since start of
    // day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatsfailstartacrs interface{}

    // This object indicates the total number of failed Interim ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatsfailinterimacrs interface{}

    // This object indicates the total number of failed Stop ACRs since start of
    // day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatsfailstopacrs interface{}

    // This object indicates the total number of failed Event ACRs since start of
    // day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    Csbrfbillrealmstatsfaileventacrs interface{}
}

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetFilter() yfilter.YFilter { return csbrfbillrealmstatsentry.YFilter }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) SetFilter(yf yfilter.YFilter) { csbrfbillrealmstatsentry.YFilter = yf }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbRfBillRealmStatsIndex" { return "Csbrfbillrealmstatsindex" }
    if yname == "csbRfBillRealmStatsRealmName" { return "Csbrfbillrealmstatsrealmname" }
    if yname == "csbRfBillRealmStatsTotalStartAcrs" { return "Csbrfbillrealmstatstotalstartacrs" }
    if yname == "csbRfBillRealmStatsTotalInterimAcrs" { return "Csbrfbillrealmstatstotalinterimacrs" }
    if yname == "csbRfBillRealmStatsTotalStopAcrs" { return "Csbrfbillrealmstatstotalstopacrs" }
    if yname == "csbRfBillRealmStatsTotalEventAcrs" { return "Csbrfbillrealmstatstotaleventacrs" }
    if yname == "csbRfBillRealmStatsSuccStartAcrs" { return "Csbrfbillrealmstatssuccstartacrs" }
    if yname == "csbRfBillRealmStatsSuccInterimAcrs" { return "Csbrfbillrealmstatssuccinterimacrs" }
    if yname == "csbRfBillRealmStatsSuccStopAcrs" { return "Csbrfbillrealmstatssuccstopacrs" }
    if yname == "csbRfBillRealmStatsSuccEventAcrs" { return "Csbrfbillrealmstatssucceventacrs" }
    if yname == "csbRfBillRealmStatsFailStartAcrs" { return "Csbrfbillrealmstatsfailstartacrs" }
    if yname == "csbRfBillRealmStatsFailInterimAcrs" { return "Csbrfbillrealmstatsfailinterimacrs" }
    if yname == "csbRfBillRealmStatsFailStopAcrs" { return "Csbrfbillrealmstatsfailstopacrs" }
    if yname == "csbRfBillRealmStatsFailEventAcrs" { return "Csbrfbillrealmstatsfaileventacrs" }
    return ""
}

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetSegmentPath() string {
    return "csbRfBillRealmStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbcallstatsserviceindex) + "']" + "[csbRfBillRealmStatsIndex='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbrfbillrealmstatsindex) + "']" + "[csbRfBillRealmStatsRealmName='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbrfbillrealmstatsrealmname) + "']"
}

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbrfbillrealmstatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbrfbillrealmstatsentry.Csbcallstatsserviceindex
    leafs["csbRfBillRealmStatsIndex"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatsindex
    leafs["csbRfBillRealmStatsRealmName"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatsrealmname
    leafs["csbRfBillRealmStatsTotalStartAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatstotalstartacrs
    leafs["csbRfBillRealmStatsTotalInterimAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatstotalinterimacrs
    leafs["csbRfBillRealmStatsTotalStopAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatstotalstopacrs
    leafs["csbRfBillRealmStatsTotalEventAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatstotaleventacrs
    leafs["csbRfBillRealmStatsSuccStartAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatssuccstartacrs
    leafs["csbRfBillRealmStatsSuccInterimAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatssuccinterimacrs
    leafs["csbRfBillRealmStatsSuccStopAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatssuccstopacrs
    leafs["csbRfBillRealmStatsSuccEventAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatssucceventacrs
    leafs["csbRfBillRealmStatsFailStartAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatsfailstartacrs
    leafs["csbRfBillRealmStatsFailInterimAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatsfailinterimacrs
    leafs["csbRfBillRealmStatsFailStopAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatsfailstopacrs
    leafs["csbRfBillRealmStatsFailEventAcrs"] = csbrfbillrealmstatsentry.Csbrfbillrealmstatsfaileventacrs
    return leafs
}

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetYangName() string { return "csbRfBillRealmStatsEntry" }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) SetParent(parent types.Entity) { csbrfbillrealmstatsentry.parent = parent }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetParent() types.Entity { return csbrfbillrealmstatsentry.parent }

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetParentYangName() string { return "csbRfBillRealmStatsTable" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable
// This table reports count of SIP request and various SIP
// responses  for each SIP method on a SIP adjacency in a given
// interval. Each entry in this table represents a SIP method, its
// incoming and outgoing count, individual incoming and outgoing 
// count of various SIP responses for this method on a SIP
// adjacency in a given interval. To understand the meaning of 
// interval please refer <Periodic Statistics> section in 
// description of ciscoSbcStatsMIB.  
// This table is indexed on csbSIPMthdCurrentStatsAdjName,
// csbSIPMthdCurrentStatsMethod and 
// csbSIPMthdCurrentStatsInterval. The other indices of this
// table are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the csbSIPMthdCurrentStatsTable. Each row describes a
    // SIP method and various responses count for this method on a given SIP
    // adjacency and given interval. This table is indexed on
    // csbSIPMthdCurrentStatsAdjName, csbSIPMthdCurrentStatsMethod and 
    // csbSIPMthdCurrentStatsInterval. The other indices of this table are
    // csbCallStatsInstanceIndex defined in  csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry.
    Csbsipmthdcurrentstatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry
}

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetFilter() yfilter.YFilter { return csbsipmthdcurrentstatstable.YFilter }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) SetFilter(yf yfilter.YFilter) { csbsipmthdcurrentstatstable.YFilter = yf }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetGoName(yname string) string {
    if yname == "csbSIPMthdCurrentStatsEntry" { return "Csbsipmthdcurrentstatsentry" }
    return ""
}

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetSegmentPath() string {
    return "csbSIPMthdCurrentStatsTable"
}

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbSIPMthdCurrentStatsEntry" {
        for _, c := range csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry {
            if csbsipmthdcurrentstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry{}
        csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry = append(csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry, child)
        return &csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry[len(csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry)-1]
    }
    return nil
}

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry {
        children[csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry[i].GetSegmentPath()] = &csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry[i]
    }
    return children
}

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetYangName() string { return "csbSIPMthdCurrentStatsTable" }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) SetParent(parent types.Entity) { csbsipmthdcurrentstatstable.parent = parent }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetParent() types.Entity { return csbsipmthdcurrentstatstable.parent }

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry
// A conceptual row in the csbSIPMthdCurrentStatsTable. Each row
// describes a SIP method and various responses count for this
// method on a given SIP adjacency and given interval. This table
// is indexed on csbSIPMthdCurrentStatsAdjName,
// csbSIPMthdCurrentStatsMethod and 
// csbSIPMthdCurrentStatsInterval. The other indices of this
// table are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry struct {
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

    // This attribute is a key. This object indicates the name of the SIP
    // adjacency for which stats related with SIP request and all kind of
    // corresponding SIP responses are reported. The object acts as an index of
    // the table. The type is string.
    Csbsipmthdcurrentstatsadjname interface{}

    // This attribute is a key. This object indicates the SIP method Request. The
    // object acts as an index of the table. The type is CiscoSbcSIPMethod.
    Csbsipmthdcurrentstatsmethod interface{}

    // This attribute is a key. This object indicates the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be 5 minutes, 15 minutes, 1 hour , 1 Day. This  object acts as an index for
    // the table. The type is CiscoSbcPeriodicStatsInterval.
    Csbsipmthdcurrentstatsinterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    Csbsipmthdcurrentstatsmethodname interface{}

    // This object indicates the total incoming SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    Csbsipmthdcurrentstatsreqin interface{}

    // This object indicates the total outgoing SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    Csbsipmthdcurrentstatsreqout interface{}

    // This object indicates the total 1xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdcurrentstatsresp1Xxin interface{}

    // This object indicates the total 1xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdcurrentstatsresp1Xxout interface{}

    // This object indicates the total 2xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdcurrentstatsresp2Xxin interface{}

    // This object indicates the total 2xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdcurrentstatsresp2Xxout interface{}

    // This object indicates the total 3xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdcurrentstatsresp3Xxin interface{}

    // This object indicates the total 3xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdcurrentstatsresp3Xxout interface{}

    // This object indicates the total 4xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdcurrentstatsresp4Xxin interface{}

    // This object indicates the total 4xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdcurrentstatsresp4Xxout interface{}

    // This object indicates the total 5xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdcurrentstatsresp5Xxin interface{}

    // This object indicates the total 5xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdcurrentstatsresp5Xxout interface{}

    // This object indicates the total 6xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdcurrentstatsresp6Xxin interface{}

    // This object indicates the total 6xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdcurrentstatsresp6Xxout interface{}
}

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetFilter() yfilter.YFilter { return csbsipmthdcurrentstatsentry.YFilter }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) SetFilter(yf yfilter.YFilter) { csbsipmthdcurrentstatsentry.YFilter = yf }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbSIPMthdCurrentStatsAdjName" { return "Csbsipmthdcurrentstatsadjname" }
    if yname == "csbSIPMthdCurrentStatsMethod" { return "Csbsipmthdcurrentstatsmethod" }
    if yname == "csbSIPMthdCurrentStatsInterval" { return "Csbsipmthdcurrentstatsinterval" }
    if yname == "csbSIPMthdCurrentStatsMethodName" { return "Csbsipmthdcurrentstatsmethodname" }
    if yname == "csbSIPMthdCurrentStatsReqIn" { return "Csbsipmthdcurrentstatsreqin" }
    if yname == "csbSIPMthdCurrentStatsReqOut" { return "Csbsipmthdcurrentstatsreqout" }
    if yname == "csbSIPMthdCurrentStatsResp1xxIn" { return "Csbsipmthdcurrentstatsresp1Xxin" }
    if yname == "csbSIPMthdCurrentStatsResp1xxOut" { return "Csbsipmthdcurrentstatsresp1Xxout" }
    if yname == "csbSIPMthdCurrentStatsResp2xxIn" { return "Csbsipmthdcurrentstatsresp2Xxin" }
    if yname == "csbSIPMthdCurrentStatsResp2xxOut" { return "Csbsipmthdcurrentstatsresp2Xxout" }
    if yname == "csbSIPMthdCurrentStatsResp3xxIn" { return "Csbsipmthdcurrentstatsresp3Xxin" }
    if yname == "csbSIPMthdCurrentStatsResp3xxOut" { return "Csbsipmthdcurrentstatsresp3Xxout" }
    if yname == "csbSIPMthdCurrentStatsResp4xxIn" { return "Csbsipmthdcurrentstatsresp4Xxin" }
    if yname == "csbSIPMthdCurrentStatsResp4xxOut" { return "Csbsipmthdcurrentstatsresp4Xxout" }
    if yname == "csbSIPMthdCurrentStatsResp5xxIn" { return "Csbsipmthdcurrentstatsresp5Xxin" }
    if yname == "csbSIPMthdCurrentStatsResp5xxOut" { return "Csbsipmthdcurrentstatsresp5Xxout" }
    if yname == "csbSIPMthdCurrentStatsResp6xxIn" { return "Csbsipmthdcurrentstatsresp6Xxin" }
    if yname == "csbSIPMthdCurrentStatsResp6xxOut" { return "Csbsipmthdcurrentstatsresp6Xxout" }
    return ""
}

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetSegmentPath() string {
    return "csbSIPMthdCurrentStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdCurrentStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsadjname) + "']" + "[csbSIPMthdCurrentStatsMethod='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsmethod) + "']" + "[csbSIPMthdCurrentStatsInterval='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsinterval) + "']"
}

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbsipmthdcurrentstatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbsipmthdcurrentstatsentry.Csbcallstatsserviceindex
    leafs["csbSIPMthdCurrentStatsAdjName"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsadjname
    leafs["csbSIPMthdCurrentStatsMethod"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsmethod
    leafs["csbSIPMthdCurrentStatsInterval"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsinterval
    leafs["csbSIPMthdCurrentStatsMethodName"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsmethodname
    leafs["csbSIPMthdCurrentStatsReqIn"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsreqin
    leafs["csbSIPMthdCurrentStatsReqOut"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsreqout
    leafs["csbSIPMthdCurrentStatsResp1xxIn"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp1Xxin
    leafs["csbSIPMthdCurrentStatsResp1xxOut"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp1Xxout
    leafs["csbSIPMthdCurrentStatsResp2xxIn"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp2Xxin
    leafs["csbSIPMthdCurrentStatsResp2xxOut"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp2Xxout
    leafs["csbSIPMthdCurrentStatsResp3xxIn"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp3Xxin
    leafs["csbSIPMthdCurrentStatsResp3xxOut"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp3Xxout
    leafs["csbSIPMthdCurrentStatsResp4xxIn"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp4Xxin
    leafs["csbSIPMthdCurrentStatsResp4xxOut"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp4Xxout
    leafs["csbSIPMthdCurrentStatsResp5xxIn"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp5Xxin
    leafs["csbSIPMthdCurrentStatsResp5xxOut"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp5Xxout
    leafs["csbSIPMthdCurrentStatsResp6xxIn"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp6Xxin
    leafs["csbSIPMthdCurrentStatsResp6xxOut"] = csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp6Xxout
    return leafs
}

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetYangName() string { return "csbSIPMthdCurrentStatsEntry" }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) SetParent(parent types.Entity) { csbsipmthdcurrentstatsentry.parent = parent }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetParent() types.Entity { return csbsipmthdcurrentstatsentry.parent }

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetParentYangName() string { return "csbSIPMthdCurrentStatsTable" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable
// This table provide historical count of SIP request and various
// SIP responses for each SIP method on a SIP adjacency in various
// interval length defined by the csbSIPMthdHistoryStatsInterval
// object. Each entry in this table represents a SIP method, its
// incoming and outgoing count, individual incoming and outgoing 
// count of various SIP responses for this method on a SIP
// adjacency in a given interval. The possible values of interval
// will be previous 5 minutes, previous 15 minutes, previous 1 hour
// and previous day. To understand the meaning of interval please
// refer <Periodic Statistics> description of ciscoSbcStatsMIB.
// This table is indexed on csbSIPMthdHistoryStatsAdjName,
// csbSIPMthdHistoryStatsMethod and 
// csbSIPMthdHistoryStatsInterval. The other indices of this
// table are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the csbSIPMthdHistoryStatsTable. The entries in this
    // table are updated as interval completes in the csbSIPMthdCurrentStatsTable
    // table and the data is  moved from that table to this one. This table is
    // indexed on csbSIPMthdHistoryStatsAdjName, csbSIPMthdHistoryStatsMethod and
    // csbSIPMthdHistoryStatsInterval. The other indices of this  table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry.
    Csbsipmthdhistorystatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry
}

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetFilter() yfilter.YFilter { return csbsipmthdhistorystatstable.YFilter }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) SetFilter(yf yfilter.YFilter) { csbsipmthdhistorystatstable.YFilter = yf }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetGoName(yname string) string {
    if yname == "csbSIPMthdHistoryStatsEntry" { return "Csbsipmthdhistorystatsentry" }
    return ""
}

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetSegmentPath() string {
    return "csbSIPMthdHistoryStatsTable"
}

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbSIPMthdHistoryStatsEntry" {
        for _, c := range csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry {
            if csbsipmthdhistorystatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry{}
        csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry = append(csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry, child)
        return &csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry[len(csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry)-1]
    }
    return nil
}

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry {
        children[csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry[i].GetSegmentPath()] = &csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry[i]
    }
    return children
}

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetYangName() string { return "csbSIPMthdHistoryStatsTable" }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) SetParent(parent types.Entity) { csbsipmthdhistorystatstable.parent = parent }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetParent() types.Entity { return csbsipmthdhistorystatstable.parent }

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry
// A conceptual row in the csbSIPMthdHistoryStatsTable. The
// entries in this table are updated as interval completes in
// the csbSIPMthdCurrentStatsTable table and the data is 
// moved from that table to this one.
// This table is indexed on csbSIPMthdHistoryStatsAdjName,
// csbSIPMthdHistoryStatsMethod and
// csbSIPMthdHistoryStatsInterval. The other indices of this 
// table are csbCallStatsInstanceIndex defined in
// csbCallStatsInstanceTable and csbCallStatsServiceIndex
// defined in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry struct {
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

    // This attribute is a key. This object indicates the name of the SIP
    // adjacency for which stats related with SIP request and all kind of
    // corresponding SIP responses are reported. The object acts as an index of
    // the table. The type is string.
    Csbsipmthdhistorystatsadjname interface{}

    // This attribute is a key. This object indicates the SIP method Request. The
    // object acts as an index of the table. The type is CiscoSbcSIPMethod.
    Csbsipmthdhistorystatsmethod interface{}

    // This attribute is a key. This object indicates the interval for which the
    // historical statistics information is to be displayed. The interval values
    // can be previous 5 minutes, previous 15 minutes,  previous 1 hour and
    // previous 1 Day. This object acts as an  index for the table. The type is
    // CiscoSbcPeriodicStatsInterval.
    Csbsipmthdhistorystatsinterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    Csbsipmthdhistorystatsmethodname interface{}

    // This object indicates the total incoming SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    Csbsipmthdhistorystatsreqin interface{}

    // This object indicates the total outgoing SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    Csbsipmthdhistorystatsreqout interface{}

    // This object indicates the total 1xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdhistorystatsresp1Xxin interface{}

    // This object indicates the total 1xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdhistorystatsresp1Xxout interface{}

    // This object indicates the total 2xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdhistorystatsresp2Xxin interface{}

    // This object indicates the total 2xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdhistorystatsresp2Xxout interface{}

    // This object indicates the total 3xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdhistorystatsresp3Xxin interface{}

    // This object indicates the total 3xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdhistorystatsresp3Xxout interface{}

    // This object indicates the total 4xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdhistorystatsresp4Xxin interface{}

    // This object indicates the total 4xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdhistorystatsresp4Xxout interface{}

    // This object indicates the total 5xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdhistorystatsresp5Xxin interface{}

    // This object indicates the total 5xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdhistorystatsresp5Xxout interface{}

    // This object indicates the total 6xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    Csbsipmthdhistorystatsresp6Xxin interface{}

    // This object indicates the total 6xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    Csbsipmthdhistorystatsresp6Xxout interface{}
}

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetFilter() yfilter.YFilter { return csbsipmthdhistorystatsentry.YFilter }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) SetFilter(yf yfilter.YFilter) { csbsipmthdhistorystatsentry.YFilter = yf }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbSIPMthdHistoryStatsAdjName" { return "Csbsipmthdhistorystatsadjname" }
    if yname == "csbSIPMthdHistoryStatsMethod" { return "Csbsipmthdhistorystatsmethod" }
    if yname == "csbSIPMthdHistoryStatsInterval" { return "Csbsipmthdhistorystatsinterval" }
    if yname == "csbSIPMthdHistoryStatsMethodName" { return "Csbsipmthdhistorystatsmethodname" }
    if yname == "csbSIPMthdHistoryStatsReqIn" { return "Csbsipmthdhistorystatsreqin" }
    if yname == "csbSIPMthdHistoryStatsReqOut" { return "Csbsipmthdhistorystatsreqout" }
    if yname == "csbSIPMthdHistoryStatsResp1xxIn" { return "Csbsipmthdhistorystatsresp1Xxin" }
    if yname == "csbSIPMthdHistoryStatsResp1xxOut" { return "Csbsipmthdhistorystatsresp1Xxout" }
    if yname == "csbSIPMthdHistoryStatsResp2xxIn" { return "Csbsipmthdhistorystatsresp2Xxin" }
    if yname == "csbSIPMthdHistoryStatsResp2xxOut" { return "Csbsipmthdhistorystatsresp2Xxout" }
    if yname == "csbSIPMthdHistoryStatsResp3xxIn" { return "Csbsipmthdhistorystatsresp3Xxin" }
    if yname == "csbSIPMthdHistoryStatsResp3xxOut" { return "Csbsipmthdhistorystatsresp3Xxout" }
    if yname == "csbSIPMthdHistoryStatsResp4xxIn" { return "Csbsipmthdhistorystatsresp4Xxin" }
    if yname == "csbSIPMthdHistoryStatsResp4xxOut" { return "Csbsipmthdhistorystatsresp4Xxout" }
    if yname == "csbSIPMthdHistoryStatsResp5xxIn" { return "Csbsipmthdhistorystatsresp5Xxin" }
    if yname == "csbSIPMthdHistoryStatsResp5xxOut" { return "Csbsipmthdhistorystatsresp5Xxout" }
    if yname == "csbSIPMthdHistoryStatsResp6xxIn" { return "Csbsipmthdhistorystatsresp6Xxin" }
    if yname == "csbSIPMthdHistoryStatsResp6xxOut" { return "Csbsipmthdhistorystatsresp6Xxout" }
    return ""
}

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetSegmentPath() string {
    return "csbSIPMthdHistoryStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdHistoryStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsadjname) + "']" + "[csbSIPMthdHistoryStatsMethod='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsmethod) + "']" + "[csbSIPMthdHistoryStatsInterval='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsinterval) + "']"
}

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbsipmthdhistorystatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbsipmthdhistorystatsentry.Csbcallstatsserviceindex
    leafs["csbSIPMthdHistoryStatsAdjName"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsadjname
    leafs["csbSIPMthdHistoryStatsMethod"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsmethod
    leafs["csbSIPMthdHistoryStatsInterval"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsinterval
    leafs["csbSIPMthdHistoryStatsMethodName"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsmethodname
    leafs["csbSIPMthdHistoryStatsReqIn"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsreqin
    leafs["csbSIPMthdHistoryStatsReqOut"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsreqout
    leafs["csbSIPMthdHistoryStatsResp1xxIn"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp1Xxin
    leafs["csbSIPMthdHistoryStatsResp1xxOut"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp1Xxout
    leafs["csbSIPMthdHistoryStatsResp2xxIn"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp2Xxin
    leafs["csbSIPMthdHistoryStatsResp2xxOut"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp2Xxout
    leafs["csbSIPMthdHistoryStatsResp3xxIn"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp3Xxin
    leafs["csbSIPMthdHistoryStatsResp3xxOut"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp3Xxout
    leafs["csbSIPMthdHistoryStatsResp4xxIn"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp4Xxin
    leafs["csbSIPMthdHistoryStatsResp4xxOut"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp4Xxout
    leafs["csbSIPMthdHistoryStatsResp5xxIn"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp5Xxin
    leafs["csbSIPMthdHistoryStatsResp5xxOut"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp5Xxout
    leafs["csbSIPMthdHistoryStatsResp6xxIn"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp6Xxin
    leafs["csbSIPMthdHistoryStatsResp6xxOut"] = csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp6Xxout
    return leafs
}

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetYangName() string { return "csbSIPMthdHistoryStatsEntry" }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) SetParent(parent types.Entity) { csbsipmthdhistorystatsentry.parent = parent }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetParent() types.Entity { return csbsipmthdhistorystatsentry.parent }

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetParentYangName() string { return "csbSIPMthdHistoryStatsTable" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable
// This table reports SIP method request and response code
// statistics for each method and response code combination on
// given SIP adjacency in a given interval. To understand the 
// meaning of interval please refer <Periodic Statistics> section
// in description of ciscoSbcStatsMIB. An exact lookup will return
// a row only  if -
// 1) detailed response code statistics are turned on in SBC
// 2) response code  messages sent or received is non zero for 
//    given SIP adjacency, method and interval.
// Also an inexact lookup will only return rows for messages with
// non-zero counts, to protect the user from large numbers of rows 
// for response codes which have not been received or sent.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the csbSIPMthdRCCurrentStatsTable. Each entry in this
    // table represents a method and response code combination. Each entry in this
    // table is identified by a value of csbSIPMthdRCCurrentStatsAdjName,
    // csbSIPMthdRCCurrentStatsMethod, csbSIPMthdRCCurrentStatsRespCode and
    // csbSIPMthdRCCurrentStatsInterval. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry.
    Csbsipmthdrccurrentstatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry
}

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetFilter() yfilter.YFilter { return csbsipmthdrccurrentstatstable.YFilter }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) SetFilter(yf yfilter.YFilter) { csbsipmthdrccurrentstatstable.YFilter = yf }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetGoName(yname string) string {
    if yname == "csbSIPMthdRCCurrentStatsEntry" { return "Csbsipmthdrccurrentstatsentry" }
    return ""
}

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetSegmentPath() string {
    return "csbSIPMthdRCCurrentStatsTable"
}

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbSIPMthdRCCurrentStatsEntry" {
        for _, c := range csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry {
            if csbsipmthdrccurrentstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry{}
        csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry = append(csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry, child)
        return &csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry[len(csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry)-1]
    }
    return nil
}

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry {
        children[csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry[i].GetSegmentPath()] = &csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry[i]
    }
    return children
}

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetYangName() string { return "csbSIPMthdRCCurrentStatsTable" }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) SetParent(parent types.Entity) { csbsipmthdrccurrentstatstable.parent = parent }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetParent() types.Entity { return csbsipmthdrccurrentstatstable.parent }

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry
// A conceptual row in the csbSIPMthdRCCurrentStatsTable. Each
// entry in this table represents a method and response code
// combination. Each entry in this table is identified by a value
// of csbSIPMthdRCCurrentStatsAdjName,
// csbSIPMthdRCCurrentStatsMethod,
// csbSIPMthdRCCurrentStatsRespCode and
// csbSIPMthdRCCurrentStatsInterval. The other indices of this
// table are csbCallStatsInstanceIndex defined in
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry struct {
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

    // This attribute is a key. This identifies the name of the adjacency for
    // which statistics are reported. This object acts as an index for the table.
    // The type is string.
    Csbsipmthdrccurrentstatsadjname interface{}

    // This attribute is a key. This object indicates the SIP method request. This
    // object acts as an index for the table. The type is CiscoSbcSIPMethod.
    Csbsipmthdrccurrentstatsmethod interface{}

    // This attribute is a key. This object indicates the response code for the
    // SIP message request. The range of valid values for SIP response codes is
    // 100 - 999. This object acts as an index for the table. The type is
    // interface{} with range: 0..4294967295.
    Csbsipmthdrccurrentstatsrespcode interface{}

    // This attribute is a key. This object identifies the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be 5 min, 15 mins, 1 hour , 1 Day. This object acts as an index for the
    // table. The type is CiscoSbcPeriodicStatsInterval.
    Csbsipmthdrccurrentstatsinterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    Csbsipmthdrccurrentstatsmethodname interface{}

    // This object indicates the total SIP messages with this response code this
    // method received on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    Csbsipmthdrccurrentstatsrespin interface{}

    // This object indicates the total SIP messages with this response code for
    // this method sent on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    Csbsipmthdrccurrentstatsrespout interface{}
}

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetFilter() yfilter.YFilter { return csbsipmthdrccurrentstatsentry.YFilter }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) SetFilter(yf yfilter.YFilter) { csbsipmthdrccurrentstatsentry.YFilter = yf }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbSIPMthdRCCurrentStatsAdjName" { return "Csbsipmthdrccurrentstatsadjname" }
    if yname == "csbSIPMthdRCCurrentStatsMethod" { return "Csbsipmthdrccurrentstatsmethod" }
    if yname == "csbSIPMthdRCCurrentStatsRespCode" { return "Csbsipmthdrccurrentstatsrespcode" }
    if yname == "csbSIPMthdRCCurrentStatsInterval" { return "Csbsipmthdrccurrentstatsinterval" }
    if yname == "csbSIPMthdRCCurrentStatsMethodName" { return "Csbsipmthdrccurrentstatsmethodname" }
    if yname == "csbSIPMthdRCCurrentStatsRespIn" { return "Csbsipmthdrccurrentstatsrespin" }
    if yname == "csbSIPMthdRCCurrentStatsRespOut" { return "Csbsipmthdrccurrentstatsrespout" }
    return ""
}

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetSegmentPath() string {
    return "csbSIPMthdRCCurrentStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdRCCurrentStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsadjname) + "']" + "[csbSIPMthdRCCurrentStatsMethod='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsmethod) + "']" + "[csbSIPMthdRCCurrentStatsRespCode='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespcode) + "']" + "[csbSIPMthdRCCurrentStatsInterval='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsinterval) + "']"
}

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbsipmthdrccurrentstatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbsipmthdrccurrentstatsentry.Csbcallstatsserviceindex
    leafs["csbSIPMthdRCCurrentStatsAdjName"] = csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsadjname
    leafs["csbSIPMthdRCCurrentStatsMethod"] = csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsmethod
    leafs["csbSIPMthdRCCurrentStatsRespCode"] = csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespcode
    leafs["csbSIPMthdRCCurrentStatsInterval"] = csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsinterval
    leafs["csbSIPMthdRCCurrentStatsMethodName"] = csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsmethodname
    leafs["csbSIPMthdRCCurrentStatsRespIn"] = csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespin
    leafs["csbSIPMthdRCCurrentStatsRespOut"] = csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespout
    return leafs
}

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetYangName() string { return "csbSIPMthdRCCurrentStatsEntry" }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) SetParent(parent types.Entity) { csbsipmthdrccurrentstatsentry.parent = parent }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetParent() types.Entity { return csbsipmthdrccurrentstatsentry.parent }

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetParentYangName() string { return "csbSIPMthdRCCurrentStatsTable" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable
// This table reports historical data for SIP method request and
// response code statistics for each method and response code 
// combination in a given past interval. The possible values of 
// interval will be previous 5 minutes, previous 15 minutes, 
// previous 1 hour and previous day. To understand the 
// meaning of interval please refer <Periodic Statistics> section
// in description of ciscoSbcStatsMIB. An exact lookup will return
// a row only  if -
// 1) detailed response code statistics are turned on in SBC
// 2) response code  messages sent or received is non zero for 
//    given SIP adjacency, method and interval.
// Also an inexact lookup will only return rows for messages with
// non-zero counts, to protect the user from large numbers of rows
// for response codes which have not been received or sent.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the csbSIPMthdRCHistoryStatsTable. The entries in this
    // table are updated as interval completes in the
    // csbSIPMthdRCCurrentStatsTable table and the data is  moved from that table
    // to this one. Each entry in this table  is identified by a value of
    // csbSIPMthdRCHistoryStatsAdjName,  csbSIPMthdRCHistoryStatsMethod,
    // csbSIPMthdRCHistoryStatsRespCode and csbSIPMthdRCHistoryStatsInterval. The
    // other indices of this table are csbCallStatsInstanceIndex defined in
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry.
    Csbsipmthdrchistorystatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry
}

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetFilter() yfilter.YFilter { return csbsipmthdrchistorystatstable.YFilter }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) SetFilter(yf yfilter.YFilter) { csbsipmthdrchistorystatstable.YFilter = yf }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetGoName(yname string) string {
    if yname == "csbSIPMthdRCHistoryStatsEntry" { return "Csbsipmthdrchistorystatsentry" }
    return ""
}

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetSegmentPath() string {
    return "csbSIPMthdRCHistoryStatsTable"
}

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csbSIPMthdRCHistoryStatsEntry" {
        for _, c := range csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry {
            if csbsipmthdrchistorystatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry{}
        csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry = append(csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry, child)
        return &csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry[len(csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry)-1]
    }
    return nil
}

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry {
        children[csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry[i].GetSegmentPath()] = &csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry[i]
    }
    return children
}

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetYangName() string { return "csbSIPMthdRCHistoryStatsTable" }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) SetParent(parent types.Entity) { csbsipmthdrchistorystatstable.parent = parent }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetParent() types.Entity { return csbsipmthdrchistorystatstable.parent }

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetParentYangName() string { return "CISCO-SESS-BORDER-CTRLR-STATS-MIB" }

// CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry
// A conceptual row in the csbSIPMthdRCHistoryStatsTable. The
// entries in this table are updated as interval completes in
// the csbSIPMthdRCCurrentStatsTable table and the data is 
// moved from that table to this one. Each entry in this table 
// is identified by a value of csbSIPMthdRCHistoryStatsAdjName, 
// csbSIPMthdRCHistoryStatsMethod,
// csbSIPMthdRCHistoryStatsRespCode and
// csbSIPMthdRCHistoryStatsInterval. The other indices of this
// table are csbCallStatsInstanceIndex defined in
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry struct {
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

    // This attribute is a key. This identifies the name of the adjacency for
    // which statistics are reported. This object acts as an index for the table.
    // The type is string.
    Csbsipmthdrchistorystatsadjname interface{}

    // This attribute is a key. This object indicates the SIP method request. This
    // object acts as an index for the table. The type is CiscoSbcSIPMethod.
    Csbsipmthdrchistorystatsmethod interface{}

    // This attribute is a key. This object indicates the response code for the
    // SIP message request. The range of valid values for SIP response codes is
    // 100 - 999. This object acts as an index for the table. The type is
    // interface{} with range: 0..4294967295.
    Csbsipmthdrchistorystatsrespcode interface{}

    // This attribute is a key. This object identifies the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be previous 5 min, previous 15 mins, previous 1  hour , previous 1 Day.
    // This object acts as an index for the table. The type is
    // CiscoSbcPeriodicStatsInterval.
    Csbsipmthdrchistorystatsinterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    Csbsipmthdrchistorystatsmethodname interface{}

    // This object indicates the total SIP messages with this response code this
    // method received on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    Csbsipmthdrchistorystatsrespin interface{}

    // This object indicates the total SIP messages with this response code for
    // this method sent on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    Csbsipmthdrchistorystatsrespout interface{}
}

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetFilter() yfilter.YFilter { return csbsipmthdrchistorystatsentry.YFilter }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) SetFilter(yf yfilter.YFilter) { csbsipmthdrchistorystatsentry.YFilter = yf }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetGoName(yname string) string {
    if yname == "csbCallStatsInstanceIndex" { return "Csbcallstatsinstanceindex" }
    if yname == "csbCallStatsServiceIndex" { return "Csbcallstatsserviceindex" }
    if yname == "csbSIPMthdRCHistoryStatsAdjName" { return "Csbsipmthdrchistorystatsadjname" }
    if yname == "csbSIPMthdRCHistoryStatsMethod" { return "Csbsipmthdrchistorystatsmethod" }
    if yname == "csbSIPMthdRCHistoryStatsRespCode" { return "Csbsipmthdrchistorystatsrespcode" }
    if yname == "csbSIPMthdRCHistoryStatsInterval" { return "Csbsipmthdrchistorystatsinterval" }
    if yname == "csbSIPMthdRCHistoryStatsMethodName" { return "Csbsipmthdrchistorystatsmethodname" }
    if yname == "csbSIPMthdRCHistoryStatsRespIn" { return "Csbsipmthdrchistorystatsrespin" }
    if yname == "csbSIPMthdRCHistoryStatsRespOut" { return "Csbsipmthdrchistorystatsrespout" }
    return ""
}

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetSegmentPath() string {
    return "csbSIPMthdRCHistoryStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdRCHistoryStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsadjname) + "']" + "[csbSIPMthdRCHistoryStatsMethod='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsmethod) + "']" + "[csbSIPMthdRCHistoryStatsRespCode='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespcode) + "']" + "[csbSIPMthdRCHistoryStatsInterval='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsinterval) + "']"
}

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csbCallStatsInstanceIndex"] = csbsipmthdrchistorystatsentry.Csbcallstatsinstanceindex
    leafs["csbCallStatsServiceIndex"] = csbsipmthdrchistorystatsentry.Csbcallstatsserviceindex
    leafs["csbSIPMthdRCHistoryStatsAdjName"] = csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsadjname
    leafs["csbSIPMthdRCHistoryStatsMethod"] = csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsmethod
    leafs["csbSIPMthdRCHistoryStatsRespCode"] = csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespcode
    leafs["csbSIPMthdRCHistoryStatsInterval"] = csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsinterval
    leafs["csbSIPMthdRCHistoryStatsMethodName"] = csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsmethodname
    leafs["csbSIPMthdRCHistoryStatsRespIn"] = csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespin
    leafs["csbSIPMthdRCHistoryStatsRespOut"] = csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespout
    return leafs
}

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetYangName() string { return "csbSIPMthdRCHistoryStatsEntry" }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) SetParent(parent types.Entity) { csbsipmthdrchistorystatsentry.parent = parent }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetParent() types.Entity { return csbsipmthdrchistorystatsentry.parent }

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetParentYangName() string { return "csbSIPMthdRCHistoryStatsTable" }

