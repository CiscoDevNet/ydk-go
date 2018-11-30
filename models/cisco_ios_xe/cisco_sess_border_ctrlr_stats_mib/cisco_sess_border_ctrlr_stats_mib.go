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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table has the reporting statistics of various RADIUS messages for
    // RADIUS servers with which the client (SBC) shares a secret. Each entry in
    // this table is identified by a  value of csbRadiusStatsEntIndex. The other
    // indices of this table are csbCallStatsInstanceIndex defined in 
    // csbCallStatsInstanceTable and csbCallStatsServiceIndex defined in
    // csbCallStatsTable.
    CsbRadiusStatsTable CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable

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
    CsbRfBillRealmStatsTable CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable

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
    CsbSIPMthdCurrentStatsTable CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable

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
    CsbSIPMthdHistoryStatsTable CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable

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
    CsbSIPMthdRCCurrentStatsTable CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable

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
    CsbSIPMthdRCHistoryStatsTable CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable
}

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetEntityData() *types.CommonEntityData {
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.YFilter = cISCOSESSBORDERCTRLRSTATSMIB.YFilter
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.YangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.SegmentPath = "CISCO-SESS-BORDER-CTRLR-STATS-MIB:CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children = types.NewOrderedMap()
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children.Append("csbRadiusStatsTable", types.YChild{"CsbRadiusStatsTable", &cISCOSESSBORDERCTRLRSTATSMIB.CsbRadiusStatsTable})
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children.Append("csbRfBillRealmStatsTable", types.YChild{"CsbRfBillRealmStatsTable", &cISCOSESSBORDERCTRLRSTATSMIB.CsbRfBillRealmStatsTable})
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children.Append("csbSIPMthdCurrentStatsTable", types.YChild{"CsbSIPMthdCurrentStatsTable", &cISCOSESSBORDERCTRLRSTATSMIB.CsbSIPMthdCurrentStatsTable})
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children.Append("csbSIPMthdHistoryStatsTable", types.YChild{"CsbSIPMthdHistoryStatsTable", &cISCOSESSBORDERCTRLRSTATSMIB.CsbSIPMthdHistoryStatsTable})
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children.Append("csbSIPMthdRCCurrentStatsTable", types.YChild{"CsbSIPMthdRCCurrentStatsTable", &cISCOSESSBORDERCTRLRSTATSMIB.CsbSIPMthdRCCurrentStatsTable})
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children.Append("csbSIPMthdRCHistoryStatsTable", types.YChild{"CsbSIPMthdRCHistoryStatsTable", &cISCOSESSBORDERCTRLRSTATSMIB.CsbSIPMthdRCHistoryStatsTable})
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.YListKeys = []string {}

    return &(cISCOSESSBORDERCTRLRSTATSMIB.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable
// This table has the reporting statistics of various RADIUS
// messages for RADIUS servers with which the client (SBC) shares a
// secret. Each entry in this table is identified by a 
// value of csbRadiusStatsEntIndex. The other indices of this table
// are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbRadiusStatsTable. There is an entry in this
    // table for each RADIUS server, as identified by a  value of
    // csbRadiusStatsEntIndex. The other indices of this  table are
    // csbCallStatsInstanceIndex defined in  csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable_CsbRadiusStatsEntry.
    CsbRadiusStatsEntry []*CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable_CsbRadiusStatsEntry
}

func (csbRadiusStatsTable *CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable) GetEntityData() *types.CommonEntityData {
    csbRadiusStatsTable.EntityData.YFilter = csbRadiusStatsTable.YFilter
    csbRadiusStatsTable.EntityData.YangName = "csbRadiusStatsTable"
    csbRadiusStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbRadiusStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbRadiusStatsTable.EntityData.SegmentPath = "csbRadiusStatsTable"
    csbRadiusStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbRadiusStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbRadiusStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbRadiusStatsTable.EntityData.Children = types.NewOrderedMap()
    csbRadiusStatsTable.EntityData.Children.Append("csbRadiusStatsEntry", types.YChild{"CsbRadiusStatsEntry", nil})
    for i := range csbRadiusStatsTable.CsbRadiusStatsEntry {
        csbRadiusStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbRadiusStatsTable.CsbRadiusStatsEntry[i]), types.YChild{"CsbRadiusStatsEntry", csbRadiusStatsTable.CsbRadiusStatsEntry[i]})
    }
    csbRadiusStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbRadiusStatsTable.EntityData.YListKeys = []string {}

    return &(csbRadiusStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable_CsbRadiusStatsEntry
// A conceptual row in the csbRadiusStatsTable. There is an
// entry in this table for each RADIUS server, as identified by a 
// value of csbRadiusStatsEntIndex. The other indices of this 
// table are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable_CsbRadiusStatsEntry struct {
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

    // This attribute is a key. This object indicates the index of the RADIUS
    // client entity that this server is configured on. This index is assigned 
    // arbitrarily by the engine and is not saved over reboots. The type is
    // interface{} with range: 0..4294967295.
    CsbRadiusStatsEntIndex interface{}

    // This object indicates the client name of the RADIUS client to which that
    // these statistics apply. The type is string.
    CsbRadiusStatsClientName interface{}

    // This object indicates the type(authentication or accounting) of the RADIUS
    // clients configured on SBC. The type is CiscoSbcRadiusClientType.
    CsbRadiusStatsClientType interface{}

    // This object indicates the server name of the RADIUS server to which that
    // these statistics apply. The type is string.
    CsbRadiusStatsSrvrName interface{}

    // This object indicates the number of RADIUS Access-Request packets sent to
    // this server.  This does not include retransmissions. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    CsbRadiusStatsAcsReqs interface{}

    // This object indicates the number of RADIUS Access-Request packets
    // retransmitted to this RADIUS server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CsbRadiusStatsAcsRtrns interface{}

    // This object indicates the number of RADIUS Access-Accept packets (valid or
    // invalid) received from this server. The type is interface{} with range:
    // 0..18446744073709551615.
    CsbRadiusStatsAcsAccpts interface{}

    // This object indicates the number of RADIUS Access-Reject packets (valid or
    // invalid) received from this server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CsbRadiusStatsAcsRejects interface{}

    // This object indicates the number of RADIUS Access-Challenge packets (valid
    // or invalid) received from this server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CsbRadiusStatsAcsChalls interface{}

    // This object indicates the number of RADIUS Accounting-Request packets sent
    // to this server. This does not include retransmissions. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    CsbRadiusStatsActReqs interface{}

    // This object indicates the number of RADIUS Accounting-Request packets
    // retransmitted to this RADIUS server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CsbRadiusStatsActRetrans interface{}

    // This object indicates the number of RADIUS Accounting-Response packets
    // (valid or invalid) received from this server. The type is interface{} with
    // range: 0..18446744073709551615. Units are packets.
    CsbRadiusStatsActRsps interface{}

    // This object indicates the number of malformed RADIUS response packets
    // received from this server.  Malformed packets include packets with an
    // invalid length. Bad authenticators, Signature attributes and unknown types
    // are not included as malformed access responses. The type is interface{}
    // with range: 0..18446744073709551615. Units are packets.
    CsbRadiusStatsMalformedRsps interface{}

    // This object indicates the number of RADIUS response packets containing
    // invalid authenticators or Signature attributes received from this server.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CsbRadiusStatsBadAuths interface{}

    // This object indicates the number of RADIUS request packets destined for
    // this server that have not yet timed out or received a response. This
    // variable is incremented when a request is sent and decremented on receipt
    // of the response or on a timeout or retransmission. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    CsbRadiusStatsPending interface{}

    // This object indicates the number of RADIUS request timeouts to this server.
    // After a timeout the client may retry to a different server or give up. A
    // retry to a different server is counted as a request as well as a timeout.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets.
    CsbRadiusStatsTimeouts interface{}

    // This object indicates the number of RADIUS packets of unknown type which
    // were received from this server. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CsbRadiusStatsUnknownType interface{}

    // This object indicates the number of RADIUS packets which were received from
    // this server and dropped for some other reason. The type is interface{} with
    // range: 0..18446744073709551615. Units are packets.
    CsbRadiusStatsDropped interface{}
}

func (csbRadiusStatsEntry *CISCOSESSBORDERCTRLRSTATSMIB_CsbRadiusStatsTable_CsbRadiusStatsEntry) GetEntityData() *types.CommonEntityData {
    csbRadiusStatsEntry.EntityData.YFilter = csbRadiusStatsEntry.YFilter
    csbRadiusStatsEntry.EntityData.YangName = "csbRadiusStatsEntry"
    csbRadiusStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbRadiusStatsEntry.EntityData.ParentYangName = "csbRadiusStatsTable"
    csbRadiusStatsEntry.EntityData.SegmentPath = "csbRadiusStatsEntry" + types.AddKeyToken(csbRadiusStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbRadiusStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbRadiusStatsEntry.CsbRadiusStatsEntIndex, "csbRadiusStatsEntIndex")
    csbRadiusStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbRadiusStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbRadiusStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbRadiusStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbRadiusStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbRadiusStatsEntry.CsbCallStatsInstanceIndex})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbRadiusStatsEntry.CsbCallStatsServiceIndex})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsEntIndex", types.YLeaf{"CsbRadiusStatsEntIndex", csbRadiusStatsEntry.CsbRadiusStatsEntIndex})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsClientName", types.YLeaf{"CsbRadiusStatsClientName", csbRadiusStatsEntry.CsbRadiusStatsClientName})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsClientType", types.YLeaf{"CsbRadiusStatsClientType", csbRadiusStatsEntry.CsbRadiusStatsClientType})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsSrvrName", types.YLeaf{"CsbRadiusStatsSrvrName", csbRadiusStatsEntry.CsbRadiusStatsSrvrName})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsAcsReqs", types.YLeaf{"CsbRadiusStatsAcsReqs", csbRadiusStatsEntry.CsbRadiusStatsAcsReqs})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsAcsRtrns", types.YLeaf{"CsbRadiusStatsAcsRtrns", csbRadiusStatsEntry.CsbRadiusStatsAcsRtrns})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsAcsAccpts", types.YLeaf{"CsbRadiusStatsAcsAccpts", csbRadiusStatsEntry.CsbRadiusStatsAcsAccpts})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsAcsRejects", types.YLeaf{"CsbRadiusStatsAcsRejects", csbRadiusStatsEntry.CsbRadiusStatsAcsRejects})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsAcsChalls", types.YLeaf{"CsbRadiusStatsAcsChalls", csbRadiusStatsEntry.CsbRadiusStatsAcsChalls})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsActReqs", types.YLeaf{"CsbRadiusStatsActReqs", csbRadiusStatsEntry.CsbRadiusStatsActReqs})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsActRetrans", types.YLeaf{"CsbRadiusStatsActRetrans", csbRadiusStatsEntry.CsbRadiusStatsActRetrans})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsActRsps", types.YLeaf{"CsbRadiusStatsActRsps", csbRadiusStatsEntry.CsbRadiusStatsActRsps})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsMalformedRsps", types.YLeaf{"CsbRadiusStatsMalformedRsps", csbRadiusStatsEntry.CsbRadiusStatsMalformedRsps})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsBadAuths", types.YLeaf{"CsbRadiusStatsBadAuths", csbRadiusStatsEntry.CsbRadiusStatsBadAuths})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsPending", types.YLeaf{"CsbRadiusStatsPending", csbRadiusStatsEntry.CsbRadiusStatsPending})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsTimeouts", types.YLeaf{"CsbRadiusStatsTimeouts", csbRadiusStatsEntry.CsbRadiusStatsTimeouts})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsUnknownType", types.YLeaf{"CsbRadiusStatsUnknownType", csbRadiusStatsEntry.CsbRadiusStatsUnknownType})
    csbRadiusStatsEntry.EntityData.Leafs.Append("csbRadiusStatsDropped", types.YLeaf{"CsbRadiusStatsDropped", csbRadiusStatsEntry.CsbRadiusStatsDropped})

    csbRadiusStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbRadiusStatsEntIndex"}

    return &(csbRadiusStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbRfBillRealmStatsTable. There is an entry in this
    // table for each realm, as identified by a  value of csbRfBillRealmStatsIndex
    // and  csbRfBillRealmStatsRealmName. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable_CsbRfBillRealmStatsEntry.
    CsbRfBillRealmStatsEntry []*CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable_CsbRfBillRealmStatsEntry
}

func (csbRfBillRealmStatsTable *CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable) GetEntityData() *types.CommonEntityData {
    csbRfBillRealmStatsTable.EntityData.YFilter = csbRfBillRealmStatsTable.YFilter
    csbRfBillRealmStatsTable.EntityData.YangName = "csbRfBillRealmStatsTable"
    csbRfBillRealmStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbRfBillRealmStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbRfBillRealmStatsTable.EntityData.SegmentPath = "csbRfBillRealmStatsTable"
    csbRfBillRealmStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbRfBillRealmStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbRfBillRealmStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbRfBillRealmStatsTable.EntityData.Children = types.NewOrderedMap()
    csbRfBillRealmStatsTable.EntityData.Children.Append("csbRfBillRealmStatsEntry", types.YChild{"CsbRfBillRealmStatsEntry", nil})
    for i := range csbRfBillRealmStatsTable.CsbRfBillRealmStatsEntry {
        csbRfBillRealmStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbRfBillRealmStatsTable.CsbRfBillRealmStatsEntry[i]), types.YChild{"CsbRfBillRealmStatsEntry", csbRfBillRealmStatsTable.CsbRfBillRealmStatsEntry[i]})
    }
    csbRfBillRealmStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbRfBillRealmStatsTable.EntityData.YListKeys = []string {}

    return &(csbRfBillRealmStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable_CsbRfBillRealmStatsEntry
// A conceptual row in the csbRfBillRealmStatsTable. There
// is an entry in this table for each realm, as identified by a 
// value of csbRfBillRealmStatsIndex and 
// csbRfBillRealmStatsRealmName. The other indices of this
// table are csbCallStatsInstanceIndex defined in
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable_CsbRfBillRealmStatsEntry struct {
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

    // This attribute is a key. This object indicates the billing method instance
    // index. The range of valid values for this field is 0 - 31. The type is
    // interface{} with range: 0..31.
    CsbRfBillRealmStatsIndex interface{}

    // This attribute is a key. This object indicates the realm for which these
    // statistics are collected. The length of this object is zero when value is
    // not assigned to it. The type is string.
    CsbRfBillRealmStatsRealmName interface{}

    // This object indicates the combined sum of successful and failed Start ACRs
    // since start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsTotalStartAcrs interface{}

    // This object indicates the combined sum of successful and failed Interim
    // ACRs since start of day or the last time the statistics were reset. The
    // type is interface{} with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsTotalInterimAcrs interface{}

    // This object indicates the combined sum of successful and failed Stop ACRs
    // since start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsTotalStopAcrs interface{}

    // This object indicates the combined sum of successful and failed Event ACRs
    // since start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsTotalEventAcrs interface{}

    // This object indicates the total number of successful Start ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsSuccStartAcrs interface{}

    // This object indicates the total number of successful Interim ACRs since
    // start of day or the last time the statistics were reset. The type is
    // interface{} with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsSuccInterimAcrs interface{}

    // This object indicates the total number of successful Stop ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsSuccStopAcrs interface{}

    // This object indicates the total number of successful Event ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsSuccEventAcrs interface{}

    // This object indicates the total number of failed Start ACRs since start of
    // day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsFailStartAcrs interface{}

    // This object indicates the total number of failed Interim ACRs since start
    // of day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsFailInterimAcrs interface{}

    // This object indicates the total number of failed Stop ACRs since start of
    // day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsFailStopAcrs interface{}

    // This object indicates the total number of failed Event ACRs since start of
    // day or the last time the statistics were reset. The type is interface{}
    // with range: 0..4294967295. Units are ACRs.
    CsbRfBillRealmStatsFailEventAcrs interface{}
}

func (csbRfBillRealmStatsEntry *CISCOSESSBORDERCTRLRSTATSMIB_CsbRfBillRealmStatsTable_CsbRfBillRealmStatsEntry) GetEntityData() *types.CommonEntityData {
    csbRfBillRealmStatsEntry.EntityData.YFilter = csbRfBillRealmStatsEntry.YFilter
    csbRfBillRealmStatsEntry.EntityData.YangName = "csbRfBillRealmStatsEntry"
    csbRfBillRealmStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbRfBillRealmStatsEntry.EntityData.ParentYangName = "csbRfBillRealmStatsTable"
    csbRfBillRealmStatsEntry.EntityData.SegmentPath = "csbRfBillRealmStatsEntry" + types.AddKeyToken(csbRfBillRealmStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbRfBillRealmStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbRfBillRealmStatsEntry.CsbRfBillRealmStatsIndex, "csbRfBillRealmStatsIndex") + types.AddKeyToken(csbRfBillRealmStatsEntry.CsbRfBillRealmStatsRealmName, "csbRfBillRealmStatsRealmName")
    csbRfBillRealmStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbRfBillRealmStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbRfBillRealmStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbRfBillRealmStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbRfBillRealmStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbRfBillRealmStatsEntry.CsbCallStatsInstanceIndex})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbRfBillRealmStatsEntry.CsbCallStatsServiceIndex})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsIndex", types.YLeaf{"CsbRfBillRealmStatsIndex", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsIndex})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsRealmName", types.YLeaf{"CsbRfBillRealmStatsRealmName", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsRealmName})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsTotalStartAcrs", types.YLeaf{"CsbRfBillRealmStatsTotalStartAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsTotalStartAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsTotalInterimAcrs", types.YLeaf{"CsbRfBillRealmStatsTotalInterimAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsTotalInterimAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsTotalStopAcrs", types.YLeaf{"CsbRfBillRealmStatsTotalStopAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsTotalStopAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsTotalEventAcrs", types.YLeaf{"CsbRfBillRealmStatsTotalEventAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsTotalEventAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsSuccStartAcrs", types.YLeaf{"CsbRfBillRealmStatsSuccStartAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsSuccStartAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsSuccInterimAcrs", types.YLeaf{"CsbRfBillRealmStatsSuccInterimAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsSuccInterimAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsSuccStopAcrs", types.YLeaf{"CsbRfBillRealmStatsSuccStopAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsSuccStopAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsSuccEventAcrs", types.YLeaf{"CsbRfBillRealmStatsSuccEventAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsSuccEventAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsFailStartAcrs", types.YLeaf{"CsbRfBillRealmStatsFailStartAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsFailStartAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsFailInterimAcrs", types.YLeaf{"CsbRfBillRealmStatsFailInterimAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsFailInterimAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsFailStopAcrs", types.YLeaf{"CsbRfBillRealmStatsFailStopAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsFailStopAcrs})
    csbRfBillRealmStatsEntry.EntityData.Leafs.Append("csbRfBillRealmStatsFailEventAcrs", types.YLeaf{"CsbRfBillRealmStatsFailEventAcrs", csbRfBillRealmStatsEntry.CsbRfBillRealmStatsFailEventAcrs})

    csbRfBillRealmStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbRfBillRealmStatsIndex", "CsbRfBillRealmStatsRealmName"}

    return &(csbRfBillRealmStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbSIPMthdCurrentStatsTable. Each row describes a
    // SIP method and various responses count for this method on a given SIP
    // adjacency and given interval. This table is indexed on
    // csbSIPMthdCurrentStatsAdjName, csbSIPMthdCurrentStatsMethod and 
    // csbSIPMthdCurrentStatsInterval. The other indices of this table are
    // csbCallStatsInstanceIndex defined in  csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable_CsbSIPMthdCurrentStatsEntry.
    CsbSIPMthdCurrentStatsEntry []*CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable_CsbSIPMthdCurrentStatsEntry
}

func (csbSIPMthdCurrentStatsTable *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable) GetEntityData() *types.CommonEntityData {
    csbSIPMthdCurrentStatsTable.EntityData.YFilter = csbSIPMthdCurrentStatsTable.YFilter
    csbSIPMthdCurrentStatsTable.EntityData.YangName = "csbSIPMthdCurrentStatsTable"
    csbSIPMthdCurrentStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdCurrentStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbSIPMthdCurrentStatsTable.EntityData.SegmentPath = "csbSIPMthdCurrentStatsTable"
    csbSIPMthdCurrentStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdCurrentStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdCurrentStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdCurrentStatsTable.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdCurrentStatsTable.EntityData.Children.Append("csbSIPMthdCurrentStatsEntry", types.YChild{"CsbSIPMthdCurrentStatsEntry", nil})
    for i := range csbSIPMthdCurrentStatsTable.CsbSIPMthdCurrentStatsEntry {
        csbSIPMthdCurrentStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbSIPMthdCurrentStatsTable.CsbSIPMthdCurrentStatsEntry[i]), types.YChild{"CsbSIPMthdCurrentStatsEntry", csbSIPMthdCurrentStatsTable.CsbSIPMthdCurrentStatsEntry[i]})
    }
    csbSIPMthdCurrentStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbSIPMthdCurrentStatsTable.EntityData.YListKeys = []string {}

    return &(csbSIPMthdCurrentStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable_CsbSIPMthdCurrentStatsEntry
// A conceptual row in the csbSIPMthdCurrentStatsTable. Each row
// describes a SIP method and various responses count for this
// method on a given SIP adjacency and given interval. This table
// is indexed on csbSIPMthdCurrentStatsAdjName,
// csbSIPMthdCurrentStatsMethod and 
// csbSIPMthdCurrentStatsInterval. The other indices of this
// table are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable_CsbSIPMthdCurrentStatsEntry struct {
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

    // This attribute is a key. This object indicates the name of the SIP
    // adjacency for which stats related with SIP request and all kind of
    // corresponding SIP responses are reported. The object acts as an index of
    // the table. The type is string.
    CsbSIPMthdCurrentStatsAdjName interface{}

    // This attribute is a key. This object indicates the SIP method Request. The
    // object acts as an index of the table. The type is CiscoSbcSIPMethod.
    CsbSIPMthdCurrentStatsMethod interface{}

    // This attribute is a key. This object indicates the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be 5 minutes, 15 minutes, 1 hour , 1 Day. This  object acts as an index for
    // the table. The type is CiscoSbcPeriodicStatsInterval.
    CsbSIPMthdCurrentStatsInterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    CsbSIPMthdCurrentStatsMethodName interface{}

    // This object indicates the total incoming SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    CsbSIPMthdCurrentStatsReqIn interface{}

    // This object indicates the total outgoing SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    CsbSIPMthdCurrentStatsReqOut interface{}

    // This object indicates the total 1xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdCurrentStatsResp1xxIn interface{}

    // This object indicates the total 1xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdCurrentStatsResp1xxOut interface{}

    // This object indicates the total 2xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdCurrentStatsResp2xxIn interface{}

    // This object indicates the total 2xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdCurrentStatsResp2xxOut interface{}

    // This object indicates the total 3xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdCurrentStatsResp3xxIn interface{}

    // This object indicates the total 3xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdCurrentStatsResp3xxOut interface{}

    // This object indicates the total 4xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdCurrentStatsResp4xxIn interface{}

    // This object indicates the total 4xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdCurrentStatsResp4xxOut interface{}

    // This object indicates the total 5xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdCurrentStatsResp5xxIn interface{}

    // This object indicates the total 5xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdCurrentStatsResp5xxOut interface{}

    // This object indicates the total 6xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdCurrentStatsResp6xxIn interface{}

    // This object indicates the total 6xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdCurrentStatsResp6xxOut interface{}
}

func (csbSIPMthdCurrentStatsEntry *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdCurrentStatsTable_CsbSIPMthdCurrentStatsEntry) GetEntityData() *types.CommonEntityData {
    csbSIPMthdCurrentStatsEntry.EntityData.YFilter = csbSIPMthdCurrentStatsEntry.YFilter
    csbSIPMthdCurrentStatsEntry.EntityData.YangName = "csbSIPMthdCurrentStatsEntry"
    csbSIPMthdCurrentStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdCurrentStatsEntry.EntityData.ParentYangName = "csbSIPMthdCurrentStatsTable"
    csbSIPMthdCurrentStatsEntry.EntityData.SegmentPath = "csbSIPMthdCurrentStatsEntry" + types.AddKeyToken(csbSIPMthdCurrentStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbSIPMthdCurrentStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsAdjName, "csbSIPMthdCurrentStatsAdjName") + types.AddKeyToken(csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsMethod, "csbSIPMthdCurrentStatsMethod") + types.AddKeyToken(csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsInterval, "csbSIPMthdCurrentStatsInterval")
    csbSIPMthdCurrentStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdCurrentStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdCurrentStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdCurrentStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbSIPMthdCurrentStatsEntry.CsbCallStatsInstanceIndex})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbSIPMthdCurrentStatsEntry.CsbCallStatsServiceIndex})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsAdjName", types.YLeaf{"CsbSIPMthdCurrentStatsAdjName", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsAdjName})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsMethod", types.YLeaf{"CsbSIPMthdCurrentStatsMethod", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsMethod})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsInterval", types.YLeaf{"CsbSIPMthdCurrentStatsInterval", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsInterval})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsMethodName", types.YLeaf{"CsbSIPMthdCurrentStatsMethodName", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsMethodName})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsReqIn", types.YLeaf{"CsbSIPMthdCurrentStatsReqIn", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsReqIn})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsReqOut", types.YLeaf{"CsbSIPMthdCurrentStatsReqOut", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsReqOut})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp1xxIn", types.YLeaf{"CsbSIPMthdCurrentStatsResp1xxIn", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp1xxIn})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp1xxOut", types.YLeaf{"CsbSIPMthdCurrentStatsResp1xxOut", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp1xxOut})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp2xxIn", types.YLeaf{"CsbSIPMthdCurrentStatsResp2xxIn", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp2xxIn})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp2xxOut", types.YLeaf{"CsbSIPMthdCurrentStatsResp2xxOut", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp2xxOut})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp3xxIn", types.YLeaf{"CsbSIPMthdCurrentStatsResp3xxIn", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp3xxIn})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp3xxOut", types.YLeaf{"CsbSIPMthdCurrentStatsResp3xxOut", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp3xxOut})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp4xxIn", types.YLeaf{"CsbSIPMthdCurrentStatsResp4xxIn", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp4xxIn})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp4xxOut", types.YLeaf{"CsbSIPMthdCurrentStatsResp4xxOut", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp4xxOut})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp5xxIn", types.YLeaf{"CsbSIPMthdCurrentStatsResp5xxIn", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp5xxIn})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp5xxOut", types.YLeaf{"CsbSIPMthdCurrentStatsResp5xxOut", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp5xxOut})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp6xxIn", types.YLeaf{"CsbSIPMthdCurrentStatsResp6xxIn", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp6xxIn})
    csbSIPMthdCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdCurrentStatsResp6xxOut", types.YLeaf{"CsbSIPMthdCurrentStatsResp6xxOut", csbSIPMthdCurrentStatsEntry.CsbSIPMthdCurrentStatsResp6xxOut})

    csbSIPMthdCurrentStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbSIPMthdCurrentStatsAdjName", "CsbSIPMthdCurrentStatsMethod", "CsbSIPMthdCurrentStatsInterval"}

    return &(csbSIPMthdCurrentStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbSIPMthdHistoryStatsTable. The entries in this
    // table are updated as interval completes in the csbSIPMthdCurrentStatsTable
    // table and the data is  moved from that table to this one. This table is
    // indexed on csbSIPMthdHistoryStatsAdjName, csbSIPMthdHistoryStatsMethod and
    // csbSIPMthdHistoryStatsInterval. The other indices of this  table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable_CsbSIPMthdHistoryStatsEntry.
    CsbSIPMthdHistoryStatsEntry []*CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable_CsbSIPMthdHistoryStatsEntry
}

func (csbSIPMthdHistoryStatsTable *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable) GetEntityData() *types.CommonEntityData {
    csbSIPMthdHistoryStatsTable.EntityData.YFilter = csbSIPMthdHistoryStatsTable.YFilter
    csbSIPMthdHistoryStatsTable.EntityData.YangName = "csbSIPMthdHistoryStatsTable"
    csbSIPMthdHistoryStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdHistoryStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbSIPMthdHistoryStatsTable.EntityData.SegmentPath = "csbSIPMthdHistoryStatsTable"
    csbSIPMthdHistoryStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdHistoryStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdHistoryStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdHistoryStatsTable.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdHistoryStatsTable.EntityData.Children.Append("csbSIPMthdHistoryStatsEntry", types.YChild{"CsbSIPMthdHistoryStatsEntry", nil})
    for i := range csbSIPMthdHistoryStatsTable.CsbSIPMthdHistoryStatsEntry {
        csbSIPMthdHistoryStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbSIPMthdHistoryStatsTable.CsbSIPMthdHistoryStatsEntry[i]), types.YChild{"CsbSIPMthdHistoryStatsEntry", csbSIPMthdHistoryStatsTable.CsbSIPMthdHistoryStatsEntry[i]})
    }
    csbSIPMthdHistoryStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbSIPMthdHistoryStatsTable.EntityData.YListKeys = []string {}

    return &(csbSIPMthdHistoryStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable_CsbSIPMthdHistoryStatsEntry
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable_CsbSIPMthdHistoryStatsEntry struct {
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

    // This attribute is a key. This object indicates the name of the SIP
    // adjacency for which stats related with SIP request and all kind of
    // corresponding SIP responses are reported. The object acts as an index of
    // the table. The type is string.
    CsbSIPMthdHistoryStatsAdjName interface{}

    // This attribute is a key. This object indicates the SIP method Request. The
    // object acts as an index of the table. The type is CiscoSbcSIPMethod.
    CsbSIPMthdHistoryStatsMethod interface{}

    // This attribute is a key. This object indicates the interval for which the
    // historical statistics information is to be displayed. The interval values
    // can be previous 5 minutes, previous 15 minutes,  previous 1 hour and
    // previous 1 Day. This object acts as an  index for the table. The type is
    // CiscoSbcPeriodicStatsInterval.
    CsbSIPMthdHistoryStatsInterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    CsbSIPMthdHistoryStatsMethodName interface{}

    // This object indicates the total incoming SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    CsbSIPMthdHistoryStatsReqIn interface{}

    // This object indicates the total outgoing SIP message requests of this type
    // on this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are requests.
    CsbSIPMthdHistoryStatsReqOut interface{}

    // This object indicates the total 1xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdHistoryStatsResp1xxIn interface{}

    // This object indicates the total 1xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdHistoryStatsResp1xxOut interface{}

    // This object indicates the total 2xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdHistoryStatsResp2xxIn interface{}

    // This object indicates the total 2xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdHistoryStatsResp2xxOut interface{}

    // This object indicates the total 3xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdHistoryStatsResp3xxIn interface{}

    // This object indicates the total 3xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdHistoryStatsResp3xxOut interface{}

    // This object indicates the total 4xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdHistoryStatsResp4xxIn interface{}

    // This object indicates the total 4xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdHistoryStatsResp4xxOut interface{}

    // This object indicates the total 5xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdHistoryStatsResp5xxIn interface{}

    // This object indicates the total 5xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdHistoryStatsResp5xxOut interface{}

    // This object indicates the total 6xx responses for this method received on
    // this SIP adjacency. The type is interface{} with range: 0..4294967295.
    // Units are responses.
    CsbSIPMthdHistoryStatsResp6xxIn interface{}

    // This object indicates the total 6xx responses for this method sent on this
    // SIP adjacency. The type is interface{} with range: 0..4294967295. Units are
    // responses.
    CsbSIPMthdHistoryStatsResp6xxOut interface{}
}

func (csbSIPMthdHistoryStatsEntry *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdHistoryStatsTable_CsbSIPMthdHistoryStatsEntry) GetEntityData() *types.CommonEntityData {
    csbSIPMthdHistoryStatsEntry.EntityData.YFilter = csbSIPMthdHistoryStatsEntry.YFilter
    csbSIPMthdHistoryStatsEntry.EntityData.YangName = "csbSIPMthdHistoryStatsEntry"
    csbSIPMthdHistoryStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdHistoryStatsEntry.EntityData.ParentYangName = "csbSIPMthdHistoryStatsTable"
    csbSIPMthdHistoryStatsEntry.EntityData.SegmentPath = "csbSIPMthdHistoryStatsEntry" + types.AddKeyToken(csbSIPMthdHistoryStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbSIPMthdHistoryStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsAdjName, "csbSIPMthdHistoryStatsAdjName") + types.AddKeyToken(csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsMethod, "csbSIPMthdHistoryStatsMethod") + types.AddKeyToken(csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsInterval, "csbSIPMthdHistoryStatsInterval")
    csbSIPMthdHistoryStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdHistoryStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdHistoryStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdHistoryStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbSIPMthdHistoryStatsEntry.CsbCallStatsInstanceIndex})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbSIPMthdHistoryStatsEntry.CsbCallStatsServiceIndex})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsAdjName", types.YLeaf{"CsbSIPMthdHistoryStatsAdjName", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsAdjName})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsMethod", types.YLeaf{"CsbSIPMthdHistoryStatsMethod", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsMethod})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsInterval", types.YLeaf{"CsbSIPMthdHistoryStatsInterval", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsInterval})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsMethodName", types.YLeaf{"CsbSIPMthdHistoryStatsMethodName", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsMethodName})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsReqIn", types.YLeaf{"CsbSIPMthdHistoryStatsReqIn", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsReqIn})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsReqOut", types.YLeaf{"CsbSIPMthdHistoryStatsReqOut", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsReqOut})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp1xxIn", types.YLeaf{"CsbSIPMthdHistoryStatsResp1xxIn", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp1xxIn})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp1xxOut", types.YLeaf{"CsbSIPMthdHistoryStatsResp1xxOut", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp1xxOut})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp2xxIn", types.YLeaf{"CsbSIPMthdHistoryStatsResp2xxIn", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp2xxIn})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp2xxOut", types.YLeaf{"CsbSIPMthdHistoryStatsResp2xxOut", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp2xxOut})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp3xxIn", types.YLeaf{"CsbSIPMthdHistoryStatsResp3xxIn", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp3xxIn})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp3xxOut", types.YLeaf{"CsbSIPMthdHistoryStatsResp3xxOut", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp3xxOut})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp4xxIn", types.YLeaf{"CsbSIPMthdHistoryStatsResp4xxIn", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp4xxIn})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp4xxOut", types.YLeaf{"CsbSIPMthdHistoryStatsResp4xxOut", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp4xxOut})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp5xxIn", types.YLeaf{"CsbSIPMthdHistoryStatsResp5xxIn", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp5xxIn})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp5xxOut", types.YLeaf{"CsbSIPMthdHistoryStatsResp5xxOut", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp5xxOut})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp6xxIn", types.YLeaf{"CsbSIPMthdHistoryStatsResp6xxIn", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp6xxIn})
    csbSIPMthdHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdHistoryStatsResp6xxOut", types.YLeaf{"CsbSIPMthdHistoryStatsResp6xxOut", csbSIPMthdHistoryStatsEntry.CsbSIPMthdHistoryStatsResp6xxOut})

    csbSIPMthdHistoryStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbSIPMthdHistoryStatsAdjName", "CsbSIPMthdHistoryStatsMethod", "CsbSIPMthdHistoryStatsInterval"}

    return &(csbSIPMthdHistoryStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbSIPMthdRCCurrentStatsTable. Each entry in this
    // table represents a method and response code combination. Each entry in this
    // table is identified by a value of csbSIPMthdRCCurrentStatsAdjName,
    // csbSIPMthdRCCurrentStatsMethod, csbSIPMthdRCCurrentStatsRespCode and
    // csbSIPMthdRCCurrentStatsInterval. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable_CsbSIPMthdRCCurrentStatsEntry.
    CsbSIPMthdRCCurrentStatsEntry []*CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable_CsbSIPMthdRCCurrentStatsEntry
}

func (csbSIPMthdRCCurrentStatsTable *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable) GetEntityData() *types.CommonEntityData {
    csbSIPMthdRCCurrentStatsTable.EntityData.YFilter = csbSIPMthdRCCurrentStatsTable.YFilter
    csbSIPMthdRCCurrentStatsTable.EntityData.YangName = "csbSIPMthdRCCurrentStatsTable"
    csbSIPMthdRCCurrentStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdRCCurrentStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbSIPMthdRCCurrentStatsTable.EntityData.SegmentPath = "csbSIPMthdRCCurrentStatsTable"
    csbSIPMthdRCCurrentStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdRCCurrentStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdRCCurrentStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdRCCurrentStatsTable.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdRCCurrentStatsTable.EntityData.Children.Append("csbSIPMthdRCCurrentStatsEntry", types.YChild{"CsbSIPMthdRCCurrentStatsEntry", nil})
    for i := range csbSIPMthdRCCurrentStatsTable.CsbSIPMthdRCCurrentStatsEntry {
        csbSIPMthdRCCurrentStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbSIPMthdRCCurrentStatsTable.CsbSIPMthdRCCurrentStatsEntry[i]), types.YChild{"CsbSIPMthdRCCurrentStatsEntry", csbSIPMthdRCCurrentStatsTable.CsbSIPMthdRCCurrentStatsEntry[i]})
    }
    csbSIPMthdRCCurrentStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbSIPMthdRCCurrentStatsTable.EntityData.YListKeys = []string {}

    return &(csbSIPMthdRCCurrentStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable_CsbSIPMthdRCCurrentStatsEntry
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable_CsbSIPMthdRCCurrentStatsEntry struct {
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

    // This attribute is a key. This identifies the name of the adjacency for
    // which statistics are reported. This object acts as an index for the table.
    // The type is string.
    CsbSIPMthdRCCurrentStatsAdjName interface{}

    // This attribute is a key. This object indicates the SIP method request. This
    // object acts as an index for the table. The type is CiscoSbcSIPMethod.
    CsbSIPMthdRCCurrentStatsMethod interface{}

    // This attribute is a key. This object indicates the response code for the
    // SIP message request. The range of valid values for SIP response codes is
    // 100 - 999. This object acts as an index for the table. The type is
    // interface{} with range: 0..4294967295.
    CsbSIPMthdRCCurrentStatsRespCode interface{}

    // This attribute is a key. This object identifies the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be 5 min, 15 mins, 1 hour , 1 Day. This object acts as an index for the
    // table. The type is CiscoSbcPeriodicStatsInterval.
    CsbSIPMthdRCCurrentStatsInterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    CsbSIPMthdRCCurrentStatsMethodName interface{}

    // This object indicates the total SIP messages with this response code this
    // method received on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    CsbSIPMthdRCCurrentStatsRespIn interface{}

    // This object indicates the total SIP messages with this response code for
    // this method sent on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    CsbSIPMthdRCCurrentStatsRespOut interface{}
}

func (csbSIPMthdRCCurrentStatsEntry *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCCurrentStatsTable_CsbSIPMthdRCCurrentStatsEntry) GetEntityData() *types.CommonEntityData {
    csbSIPMthdRCCurrentStatsEntry.EntityData.YFilter = csbSIPMthdRCCurrentStatsEntry.YFilter
    csbSIPMthdRCCurrentStatsEntry.EntityData.YangName = "csbSIPMthdRCCurrentStatsEntry"
    csbSIPMthdRCCurrentStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdRCCurrentStatsEntry.EntityData.ParentYangName = "csbSIPMthdRCCurrentStatsTable"
    csbSIPMthdRCCurrentStatsEntry.EntityData.SegmentPath = "csbSIPMthdRCCurrentStatsEntry" + types.AddKeyToken(csbSIPMthdRCCurrentStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbSIPMthdRCCurrentStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsAdjName, "csbSIPMthdRCCurrentStatsAdjName") + types.AddKeyToken(csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsMethod, "csbSIPMthdRCCurrentStatsMethod") + types.AddKeyToken(csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsRespCode, "csbSIPMthdRCCurrentStatsRespCode") + types.AddKeyToken(csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsInterval, "csbSIPMthdRCCurrentStatsInterval")
    csbSIPMthdRCCurrentStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdRCCurrentStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdRCCurrentStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdRCCurrentStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbSIPMthdRCCurrentStatsEntry.CsbCallStatsInstanceIndex})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbSIPMthdRCCurrentStatsEntry.CsbCallStatsServiceIndex})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCCurrentStatsAdjName", types.YLeaf{"CsbSIPMthdRCCurrentStatsAdjName", csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsAdjName})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCCurrentStatsMethod", types.YLeaf{"CsbSIPMthdRCCurrentStatsMethod", csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsMethod})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCCurrentStatsRespCode", types.YLeaf{"CsbSIPMthdRCCurrentStatsRespCode", csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsRespCode})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCCurrentStatsInterval", types.YLeaf{"CsbSIPMthdRCCurrentStatsInterval", csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsInterval})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCCurrentStatsMethodName", types.YLeaf{"CsbSIPMthdRCCurrentStatsMethodName", csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsMethodName})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCCurrentStatsRespIn", types.YLeaf{"CsbSIPMthdRCCurrentStatsRespIn", csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsRespIn})
    csbSIPMthdRCCurrentStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCCurrentStatsRespOut", types.YLeaf{"CsbSIPMthdRCCurrentStatsRespOut", csbSIPMthdRCCurrentStatsEntry.CsbSIPMthdRCCurrentStatsRespOut})

    csbSIPMthdRCCurrentStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbSIPMthdRCCurrentStatsAdjName", "CsbSIPMthdRCCurrentStatsMethod", "CsbSIPMthdRCCurrentStatsRespCode", "CsbSIPMthdRCCurrentStatsInterval"}

    return &(csbSIPMthdRCCurrentStatsEntry.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable struct {
    EntityData types.CommonEntityData
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
    // CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable_CsbSIPMthdRCHistoryStatsEntry.
    CsbSIPMthdRCHistoryStatsEntry []*CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable_CsbSIPMthdRCHistoryStatsEntry
}

func (csbSIPMthdRCHistoryStatsTable *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable) GetEntityData() *types.CommonEntityData {
    csbSIPMthdRCHistoryStatsTable.EntityData.YFilter = csbSIPMthdRCHistoryStatsTable.YFilter
    csbSIPMthdRCHistoryStatsTable.EntityData.YangName = "csbSIPMthdRCHistoryStatsTable"
    csbSIPMthdRCHistoryStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdRCHistoryStatsTable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbSIPMthdRCHistoryStatsTable.EntityData.SegmentPath = "csbSIPMthdRCHistoryStatsTable"
    csbSIPMthdRCHistoryStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdRCHistoryStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdRCHistoryStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdRCHistoryStatsTable.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdRCHistoryStatsTable.EntityData.Children.Append("csbSIPMthdRCHistoryStatsEntry", types.YChild{"CsbSIPMthdRCHistoryStatsEntry", nil})
    for i := range csbSIPMthdRCHistoryStatsTable.CsbSIPMthdRCHistoryStatsEntry {
        csbSIPMthdRCHistoryStatsTable.EntityData.Children.Append(types.GetSegmentPath(csbSIPMthdRCHistoryStatsTable.CsbSIPMthdRCHistoryStatsEntry[i]), types.YChild{"CsbSIPMthdRCHistoryStatsEntry", csbSIPMthdRCHistoryStatsTable.CsbSIPMthdRCHistoryStatsEntry[i]})
    }
    csbSIPMthdRCHistoryStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csbSIPMthdRCHistoryStatsTable.EntityData.YListKeys = []string {}

    return &(csbSIPMthdRCHistoryStatsTable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable_CsbSIPMthdRCHistoryStatsEntry
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
type CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable_CsbSIPMthdRCHistoryStatsEntry struct {
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

    // This attribute is a key. This identifies the name of the adjacency for
    // which statistics are reported. This object acts as an index for the table.
    // The type is string.
    CsbSIPMthdRCHistoryStatsAdjName interface{}

    // This attribute is a key. This object indicates the SIP method request. This
    // object acts as an index for the table. The type is CiscoSbcSIPMethod.
    CsbSIPMthdRCHistoryStatsMethod interface{}

    // This attribute is a key. This object indicates the response code for the
    // SIP message request. The range of valid values for SIP response codes is
    // 100 - 999. This object acts as an index for the table. The type is
    // interface{} with range: 0..4294967295.
    CsbSIPMthdRCHistoryStatsRespCode interface{}

    // This attribute is a key. This object identifies the interval for which the
    // periodic statistics information is to be displayed. The interval values can
    // be previous 5 min, previous 15 mins, previous 1  hour , previous 1 Day.
    // This object acts as an index for the table. The type is
    // CiscoSbcPeriodicStatsInterval.
    CsbSIPMthdRCHistoryStatsInterval interface{}

    // This object indicates the text representation of the SIP method request.
    // E.g. INVITE, ACK, BYE etc. The type is string.
    CsbSIPMthdRCHistoryStatsMethodName interface{}

    // This object indicates the total SIP messages with this response code this
    // method received on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    CsbSIPMthdRCHistoryStatsRespIn interface{}

    // This object indicates the total SIP messages with this response code for
    // this method sent on this SIP adjacency. The type is interface{} with range:
    // 0..4294967295. Units are responses.
    CsbSIPMthdRCHistoryStatsRespOut interface{}
}

func (csbSIPMthdRCHistoryStatsEntry *CISCOSESSBORDERCTRLRSTATSMIB_CsbSIPMthdRCHistoryStatsTable_CsbSIPMthdRCHistoryStatsEntry) GetEntityData() *types.CommonEntityData {
    csbSIPMthdRCHistoryStatsEntry.EntityData.YFilter = csbSIPMthdRCHistoryStatsEntry.YFilter
    csbSIPMthdRCHistoryStatsEntry.EntityData.YangName = "csbSIPMthdRCHistoryStatsEntry"
    csbSIPMthdRCHistoryStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csbSIPMthdRCHistoryStatsEntry.EntityData.ParentYangName = "csbSIPMthdRCHistoryStatsTable"
    csbSIPMthdRCHistoryStatsEntry.EntityData.SegmentPath = "csbSIPMthdRCHistoryStatsEntry" + types.AddKeyToken(csbSIPMthdRCHistoryStatsEntry.CsbCallStatsInstanceIndex, "csbCallStatsInstanceIndex") + types.AddKeyToken(csbSIPMthdRCHistoryStatsEntry.CsbCallStatsServiceIndex, "csbCallStatsServiceIndex") + types.AddKeyToken(csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsAdjName, "csbSIPMthdRCHistoryStatsAdjName") + types.AddKeyToken(csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsMethod, "csbSIPMthdRCHistoryStatsMethod") + types.AddKeyToken(csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsRespCode, "csbSIPMthdRCHistoryStatsRespCode") + types.AddKeyToken(csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsInterval, "csbSIPMthdRCHistoryStatsInterval")
    csbSIPMthdRCHistoryStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbSIPMthdRCHistoryStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbSIPMthdRCHistoryStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbSIPMthdRCHistoryStatsEntry.EntityData.Children = types.NewOrderedMap()
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbCallStatsInstanceIndex", types.YLeaf{"CsbCallStatsInstanceIndex", csbSIPMthdRCHistoryStatsEntry.CsbCallStatsInstanceIndex})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbCallStatsServiceIndex", types.YLeaf{"CsbCallStatsServiceIndex", csbSIPMthdRCHistoryStatsEntry.CsbCallStatsServiceIndex})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCHistoryStatsAdjName", types.YLeaf{"CsbSIPMthdRCHistoryStatsAdjName", csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsAdjName})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCHistoryStatsMethod", types.YLeaf{"CsbSIPMthdRCHistoryStatsMethod", csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsMethod})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCHistoryStatsRespCode", types.YLeaf{"CsbSIPMthdRCHistoryStatsRespCode", csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsRespCode})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCHistoryStatsInterval", types.YLeaf{"CsbSIPMthdRCHistoryStatsInterval", csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsInterval})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCHistoryStatsMethodName", types.YLeaf{"CsbSIPMthdRCHistoryStatsMethodName", csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsMethodName})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCHistoryStatsRespIn", types.YLeaf{"CsbSIPMthdRCHistoryStatsRespIn", csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsRespIn})
    csbSIPMthdRCHistoryStatsEntry.EntityData.Leafs.Append("csbSIPMthdRCHistoryStatsRespOut", types.YLeaf{"CsbSIPMthdRCHistoryStatsRespOut", csbSIPMthdRCHistoryStatsEntry.CsbSIPMthdRCHistoryStatsRespOut})

    csbSIPMthdRCHistoryStatsEntry.EntityData.YListKeys = []string {"CsbCallStatsInstanceIndex", "CsbCallStatsServiceIndex", "CsbSIPMthdRCHistoryStatsAdjName", "CsbSIPMthdRCHistoryStatsMethod", "CsbSIPMthdRCHistoryStatsRespCode", "CsbSIPMthdRCHistoryStatsInterval"}

    return &(csbSIPMthdRCHistoryStatsEntry.EntityData)
}

