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

func (cISCOSESSBORDERCTRLRSTATSMIB *CISCOSESSBORDERCTRLRSTATSMIB) GetEntityData() *types.CommonEntityData {
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.YFilter = cISCOSESSBORDERCTRLRSTATSMIB.YFilter
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.YangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.SegmentPath = "CISCO-SESS-BORDER-CTRLR-STATS-MIB:CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children["csbRadiusStatsTable"] = types.YChild{"Csbradiusstatstable", &cISCOSESSBORDERCTRLRSTATSMIB.Csbradiusstatstable}
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children["csbRfBillRealmStatsTable"] = types.YChild{"Csbrfbillrealmstatstable", &cISCOSESSBORDERCTRLRSTATSMIB.Csbrfbillrealmstatstable}
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children["csbSIPMthdCurrentStatsTable"] = types.YChild{"Csbsipmthdcurrentstatstable", &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdcurrentstatstable}
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children["csbSIPMthdHistoryStatsTable"] = types.YChild{"Csbsipmthdhistorystatstable", &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdhistorystatstable}
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children["csbSIPMthdRCCurrentStatsTable"] = types.YChild{"Csbsipmthdrccurrentstatstable", &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdrccurrentstatstable}
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Children["csbSIPMthdRCHistoryStatsTable"] = types.YChild{"Csbsipmthdrchistorystatstable", &cISCOSESSBORDERCTRLRSTATSMIB.Csbsipmthdrchistorystatstable}
    cISCOSESSBORDERCTRLRSTATSMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOSESSBORDERCTRLRSTATSMIB.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable
// This table has the reporting statistics of various RADIUS
// messages for RADIUS servers with which the client (SBC) shares a
// secret. Each entry in this table is identified by a 
// value of csbRadiusStatsEntIndex. The other indices of this table
// are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbRadiusStatsTable. There is an entry in this
    // table for each RADIUS server, as identified by a  value of
    // csbRadiusStatsEntIndex. The other indices of this  table are
    // csbCallStatsInstanceIndex defined in  csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry.
    Csbradiusstatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry
}

func (csbradiusstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable) GetEntityData() *types.CommonEntityData {
    csbradiusstatstable.EntityData.YFilter = csbradiusstatstable.YFilter
    csbradiusstatstable.EntityData.YangName = "csbRadiusStatsTable"
    csbradiusstatstable.EntityData.BundleName = "cisco_ios_xe"
    csbradiusstatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbradiusstatstable.EntityData.SegmentPath = "csbRadiusStatsTable"
    csbradiusstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbradiusstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbradiusstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbradiusstatstable.EntityData.Children = make(map[string]types.YChild)
    csbradiusstatstable.EntityData.Children["csbRadiusStatsEntry"] = types.YChild{"Csbradiusstatsentry", nil}
    for i := range csbradiusstatstable.Csbradiusstatsentry {
        csbradiusstatstable.EntityData.Children[types.GetSegmentPath(&csbradiusstatstable.Csbradiusstatsentry[i])] = types.YChild{"Csbradiusstatsentry", &csbradiusstatstable.Csbradiusstatsentry[i]}
    }
    csbradiusstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbradiusstatstable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry
// A conceptual row in the csbRadiusStatsTable. There is an
// entry in this table for each RADIUS server, as identified by a 
// value of csbRadiusStatsEntIndex. The other indices of this 
// table are csbCallStatsInstanceIndex defined in 
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry struct {
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

func (csbradiusstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbradiusstatstable_Csbradiusstatsentry) GetEntityData() *types.CommonEntityData {
    csbradiusstatsentry.EntityData.YFilter = csbradiusstatsentry.YFilter
    csbradiusstatsentry.EntityData.YangName = "csbRadiusStatsEntry"
    csbradiusstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbradiusstatsentry.EntityData.ParentYangName = "csbRadiusStatsTable"
    csbradiusstatsentry.EntityData.SegmentPath = "csbRadiusStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbradiusstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbradiusstatsentry.Csbcallstatsserviceindex) + "']" + "[csbRadiusStatsEntIndex='" + fmt.Sprintf("%v", csbradiusstatsentry.Csbradiusstatsentindex) + "']"
    csbradiusstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbradiusstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbradiusstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbradiusstatsentry.EntityData.Children = make(map[string]types.YChild)
    csbradiusstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbradiusstatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbradiusstatsentry.Csbcallstatsinstanceindex}
    csbradiusstatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbradiusstatsentry.Csbcallstatsserviceindex}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsEntIndex"] = types.YLeaf{"Csbradiusstatsentindex", csbradiusstatsentry.Csbradiusstatsentindex}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsClientName"] = types.YLeaf{"Csbradiusstatsclientname", csbradiusstatsentry.Csbradiusstatsclientname}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsClientType"] = types.YLeaf{"Csbradiusstatsclienttype", csbradiusstatsentry.Csbradiusstatsclienttype}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsSrvrName"] = types.YLeaf{"Csbradiusstatssrvrname", csbradiusstatsentry.Csbradiusstatssrvrname}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsAcsReqs"] = types.YLeaf{"Csbradiusstatsacsreqs", csbradiusstatsentry.Csbradiusstatsacsreqs}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsAcsRtrns"] = types.YLeaf{"Csbradiusstatsacsrtrns", csbradiusstatsentry.Csbradiusstatsacsrtrns}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsAcsAccpts"] = types.YLeaf{"Csbradiusstatsacsaccpts", csbradiusstatsentry.Csbradiusstatsacsaccpts}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsAcsRejects"] = types.YLeaf{"Csbradiusstatsacsrejects", csbradiusstatsentry.Csbradiusstatsacsrejects}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsAcsChalls"] = types.YLeaf{"Csbradiusstatsacschalls", csbradiusstatsentry.Csbradiusstatsacschalls}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsActReqs"] = types.YLeaf{"Csbradiusstatsactreqs", csbradiusstatsentry.Csbradiusstatsactreqs}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsActRetrans"] = types.YLeaf{"Csbradiusstatsactretrans", csbradiusstatsentry.Csbradiusstatsactretrans}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsActRsps"] = types.YLeaf{"Csbradiusstatsactrsps", csbradiusstatsentry.Csbradiusstatsactrsps}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsMalformedRsps"] = types.YLeaf{"Csbradiusstatsmalformedrsps", csbradiusstatsentry.Csbradiusstatsmalformedrsps}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsBadAuths"] = types.YLeaf{"Csbradiusstatsbadauths", csbradiusstatsentry.Csbradiusstatsbadauths}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsPending"] = types.YLeaf{"Csbradiusstatspending", csbradiusstatsentry.Csbradiusstatspending}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsTimeouts"] = types.YLeaf{"Csbradiusstatstimeouts", csbradiusstatsentry.Csbradiusstatstimeouts}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsUnknownType"] = types.YLeaf{"Csbradiusstatsunknowntype", csbradiusstatsentry.Csbradiusstatsunknowntype}
    csbradiusstatsentry.EntityData.Leafs["csbRadiusStatsDropped"] = types.YLeaf{"Csbradiusstatsdropped", csbradiusstatsentry.Csbradiusstatsdropped}
    return &(csbradiusstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the csbRfBillRealmStatsTable. There is an entry in this
    // table for each realm, as identified by a  value of csbRfBillRealmStatsIndex
    // and  csbRfBillRealmStatsRealmName. The other indices of this table are
    // csbCallStatsInstanceIndex defined in csbCallStatsInstanceTable and
    // csbCallStatsServiceIndex defined in csbCallStatsTable. The type is slice of
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry.
    Csbrfbillrealmstatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry
}

func (csbrfbillrealmstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable) GetEntityData() *types.CommonEntityData {
    csbrfbillrealmstatstable.EntityData.YFilter = csbrfbillrealmstatstable.YFilter
    csbrfbillrealmstatstable.EntityData.YangName = "csbRfBillRealmStatsTable"
    csbrfbillrealmstatstable.EntityData.BundleName = "cisco_ios_xe"
    csbrfbillrealmstatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbrfbillrealmstatstable.EntityData.SegmentPath = "csbRfBillRealmStatsTable"
    csbrfbillrealmstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbrfbillrealmstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbrfbillrealmstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbrfbillrealmstatstable.EntityData.Children = make(map[string]types.YChild)
    csbrfbillrealmstatstable.EntityData.Children["csbRfBillRealmStatsEntry"] = types.YChild{"Csbrfbillrealmstatsentry", nil}
    for i := range csbrfbillrealmstatstable.Csbrfbillrealmstatsentry {
        csbrfbillrealmstatstable.EntityData.Children[types.GetSegmentPath(&csbrfbillrealmstatstable.Csbrfbillrealmstatsentry[i])] = types.YChild{"Csbrfbillrealmstatsentry", &csbrfbillrealmstatstable.Csbrfbillrealmstatsentry[i]}
    }
    csbrfbillrealmstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbrfbillrealmstatstable.EntityData)
}

// CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry
// A conceptual row in the csbRfBillRealmStatsTable. There
// is an entry in this table for each realm, as identified by a 
// value of csbRfBillRealmStatsIndex and 
// csbRfBillRealmStatsRealmName. The other indices of this
// table are csbCallStatsInstanceIndex defined in
// csbCallStatsInstanceTable and csbCallStatsServiceIndex defined
// in csbCallStatsTable.
type CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry struct {
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

func (csbrfbillrealmstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbrfbillrealmstatstable_Csbrfbillrealmstatsentry) GetEntityData() *types.CommonEntityData {
    csbrfbillrealmstatsentry.EntityData.YFilter = csbrfbillrealmstatsentry.YFilter
    csbrfbillrealmstatsentry.EntityData.YangName = "csbRfBillRealmStatsEntry"
    csbrfbillrealmstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbrfbillrealmstatsentry.EntityData.ParentYangName = "csbRfBillRealmStatsTable"
    csbrfbillrealmstatsentry.EntityData.SegmentPath = "csbRfBillRealmStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbcallstatsserviceindex) + "']" + "[csbRfBillRealmStatsIndex='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbrfbillrealmstatsindex) + "']" + "[csbRfBillRealmStatsRealmName='" + fmt.Sprintf("%v", csbrfbillrealmstatsentry.Csbrfbillrealmstatsrealmname) + "']"
    csbrfbillrealmstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbrfbillrealmstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbrfbillrealmstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbrfbillrealmstatsentry.EntityData.Children = make(map[string]types.YChild)
    csbrfbillrealmstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbrfbillrealmstatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbrfbillrealmstatsentry.Csbcallstatsinstanceindex}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbrfbillrealmstatsentry.Csbcallstatsserviceindex}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsIndex"] = types.YLeaf{"Csbrfbillrealmstatsindex", csbrfbillrealmstatsentry.Csbrfbillrealmstatsindex}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsRealmName"] = types.YLeaf{"Csbrfbillrealmstatsrealmname", csbrfbillrealmstatsentry.Csbrfbillrealmstatsrealmname}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsTotalStartAcrs"] = types.YLeaf{"Csbrfbillrealmstatstotalstartacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatstotalstartacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsTotalInterimAcrs"] = types.YLeaf{"Csbrfbillrealmstatstotalinterimacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatstotalinterimacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsTotalStopAcrs"] = types.YLeaf{"Csbrfbillrealmstatstotalstopacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatstotalstopacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsTotalEventAcrs"] = types.YLeaf{"Csbrfbillrealmstatstotaleventacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatstotaleventacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsSuccStartAcrs"] = types.YLeaf{"Csbrfbillrealmstatssuccstartacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatssuccstartacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsSuccInterimAcrs"] = types.YLeaf{"Csbrfbillrealmstatssuccinterimacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatssuccinterimacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsSuccStopAcrs"] = types.YLeaf{"Csbrfbillrealmstatssuccstopacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatssuccstopacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsSuccEventAcrs"] = types.YLeaf{"Csbrfbillrealmstatssucceventacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatssucceventacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsFailStartAcrs"] = types.YLeaf{"Csbrfbillrealmstatsfailstartacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatsfailstartacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsFailInterimAcrs"] = types.YLeaf{"Csbrfbillrealmstatsfailinterimacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatsfailinterimacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsFailStopAcrs"] = types.YLeaf{"Csbrfbillrealmstatsfailstopacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatsfailstopacrs}
    csbrfbillrealmstatsentry.EntityData.Leafs["csbRfBillRealmStatsFailEventAcrs"] = types.YLeaf{"Csbrfbillrealmstatsfaileventacrs", csbrfbillrealmstatsentry.Csbrfbillrealmstatsfaileventacrs}
    return &(csbrfbillrealmstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csbsipmthdcurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable) GetEntityData() *types.CommonEntityData {
    csbsipmthdcurrentstatstable.EntityData.YFilter = csbsipmthdcurrentstatstable.YFilter
    csbsipmthdcurrentstatstable.EntityData.YangName = "csbSIPMthdCurrentStatsTable"
    csbsipmthdcurrentstatstable.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdcurrentstatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbsipmthdcurrentstatstable.EntityData.SegmentPath = "csbSIPMthdCurrentStatsTable"
    csbsipmthdcurrentstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdcurrentstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdcurrentstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdcurrentstatstable.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdcurrentstatstable.EntityData.Children["csbSIPMthdCurrentStatsEntry"] = types.YChild{"Csbsipmthdcurrentstatsentry", nil}
    for i := range csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry {
        csbsipmthdcurrentstatstable.EntityData.Children[types.GetSegmentPath(&csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry[i])] = types.YChild{"Csbsipmthdcurrentstatsentry", &csbsipmthdcurrentstatstable.Csbsipmthdcurrentstatsentry[i]}
    }
    csbsipmthdcurrentstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbsipmthdcurrentstatstable.EntityData)
}

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

func (csbsipmthdcurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdcurrentstatstable_Csbsipmthdcurrentstatsentry) GetEntityData() *types.CommonEntityData {
    csbsipmthdcurrentstatsentry.EntityData.YFilter = csbsipmthdcurrentstatsentry.YFilter
    csbsipmthdcurrentstatsentry.EntityData.YangName = "csbSIPMthdCurrentStatsEntry"
    csbsipmthdcurrentstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdcurrentstatsentry.EntityData.ParentYangName = "csbSIPMthdCurrentStatsTable"
    csbsipmthdcurrentstatsentry.EntityData.SegmentPath = "csbSIPMthdCurrentStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdCurrentStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsadjname) + "']" + "[csbSIPMthdCurrentStatsMethod='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsmethod) + "']" + "[csbSIPMthdCurrentStatsInterval='" + fmt.Sprintf("%v", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsinterval) + "']"
    csbsipmthdcurrentstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdcurrentstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdcurrentstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdcurrentstatsentry.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdcurrentstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbsipmthdcurrentstatsentry.Csbcallstatsinstanceindex}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbsipmthdcurrentstatsentry.Csbcallstatsserviceindex}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsAdjName"] = types.YLeaf{"Csbsipmthdcurrentstatsadjname", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsadjname}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsMethod"] = types.YLeaf{"Csbsipmthdcurrentstatsmethod", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsmethod}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsInterval"] = types.YLeaf{"Csbsipmthdcurrentstatsinterval", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsinterval}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsMethodName"] = types.YLeaf{"Csbsipmthdcurrentstatsmethodname", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsmethodname}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsReqIn"] = types.YLeaf{"Csbsipmthdcurrentstatsreqin", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsreqin}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsReqOut"] = types.YLeaf{"Csbsipmthdcurrentstatsreqout", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsreqout}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp1xxIn"] = types.YLeaf{"Csbsipmthdcurrentstatsresp1Xxin", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp1Xxin}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp1xxOut"] = types.YLeaf{"Csbsipmthdcurrentstatsresp1Xxout", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp1Xxout}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp2xxIn"] = types.YLeaf{"Csbsipmthdcurrentstatsresp2Xxin", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp2Xxin}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp2xxOut"] = types.YLeaf{"Csbsipmthdcurrentstatsresp2Xxout", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp2Xxout}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp3xxIn"] = types.YLeaf{"Csbsipmthdcurrentstatsresp3Xxin", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp3Xxin}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp3xxOut"] = types.YLeaf{"Csbsipmthdcurrentstatsresp3Xxout", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp3Xxout}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp4xxIn"] = types.YLeaf{"Csbsipmthdcurrentstatsresp4Xxin", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp4Xxin}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp4xxOut"] = types.YLeaf{"Csbsipmthdcurrentstatsresp4Xxout", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp4Xxout}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp5xxIn"] = types.YLeaf{"Csbsipmthdcurrentstatsresp5Xxin", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp5Xxin}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp5xxOut"] = types.YLeaf{"Csbsipmthdcurrentstatsresp5Xxout", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp5Xxout}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp6xxIn"] = types.YLeaf{"Csbsipmthdcurrentstatsresp6Xxin", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp6Xxin}
    csbsipmthdcurrentstatsentry.EntityData.Leafs["csbSIPMthdCurrentStatsResp6xxOut"] = types.YLeaf{"Csbsipmthdcurrentstatsresp6Xxout", csbsipmthdcurrentstatsentry.Csbsipmthdcurrentstatsresp6Xxout}
    return &(csbsipmthdcurrentstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csbsipmthdhistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable) GetEntityData() *types.CommonEntityData {
    csbsipmthdhistorystatstable.EntityData.YFilter = csbsipmthdhistorystatstable.YFilter
    csbsipmthdhistorystatstable.EntityData.YangName = "csbSIPMthdHistoryStatsTable"
    csbsipmthdhistorystatstable.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdhistorystatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbsipmthdhistorystatstable.EntityData.SegmentPath = "csbSIPMthdHistoryStatsTable"
    csbsipmthdhistorystatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdhistorystatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdhistorystatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdhistorystatstable.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdhistorystatstable.EntityData.Children["csbSIPMthdHistoryStatsEntry"] = types.YChild{"Csbsipmthdhistorystatsentry", nil}
    for i := range csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry {
        csbsipmthdhistorystatstable.EntityData.Children[types.GetSegmentPath(&csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry[i])] = types.YChild{"Csbsipmthdhistorystatsentry", &csbsipmthdhistorystatstable.Csbsipmthdhistorystatsentry[i]}
    }
    csbsipmthdhistorystatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbsipmthdhistorystatstable.EntityData)
}

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

func (csbsipmthdhistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdhistorystatstable_Csbsipmthdhistorystatsentry) GetEntityData() *types.CommonEntityData {
    csbsipmthdhistorystatsentry.EntityData.YFilter = csbsipmthdhistorystatsentry.YFilter
    csbsipmthdhistorystatsentry.EntityData.YangName = "csbSIPMthdHistoryStatsEntry"
    csbsipmthdhistorystatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdhistorystatsentry.EntityData.ParentYangName = "csbSIPMthdHistoryStatsTable"
    csbsipmthdhistorystatsentry.EntityData.SegmentPath = "csbSIPMthdHistoryStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdHistoryStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsadjname) + "']" + "[csbSIPMthdHistoryStatsMethod='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsmethod) + "']" + "[csbSIPMthdHistoryStatsInterval='" + fmt.Sprintf("%v", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsinterval) + "']"
    csbsipmthdhistorystatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdhistorystatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdhistorystatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdhistorystatsentry.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdhistorystatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbsipmthdhistorystatsentry.Csbcallstatsinstanceindex}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbsipmthdhistorystatsentry.Csbcallstatsserviceindex}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsAdjName"] = types.YLeaf{"Csbsipmthdhistorystatsadjname", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsadjname}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsMethod"] = types.YLeaf{"Csbsipmthdhistorystatsmethod", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsmethod}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsInterval"] = types.YLeaf{"Csbsipmthdhistorystatsinterval", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsinterval}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsMethodName"] = types.YLeaf{"Csbsipmthdhistorystatsmethodname", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsmethodname}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsReqIn"] = types.YLeaf{"Csbsipmthdhistorystatsreqin", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsreqin}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsReqOut"] = types.YLeaf{"Csbsipmthdhistorystatsreqout", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsreqout}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp1xxIn"] = types.YLeaf{"Csbsipmthdhistorystatsresp1Xxin", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp1Xxin}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp1xxOut"] = types.YLeaf{"Csbsipmthdhistorystatsresp1Xxout", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp1Xxout}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp2xxIn"] = types.YLeaf{"Csbsipmthdhistorystatsresp2Xxin", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp2Xxin}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp2xxOut"] = types.YLeaf{"Csbsipmthdhistorystatsresp2Xxout", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp2Xxout}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp3xxIn"] = types.YLeaf{"Csbsipmthdhistorystatsresp3Xxin", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp3Xxin}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp3xxOut"] = types.YLeaf{"Csbsipmthdhistorystatsresp3Xxout", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp3Xxout}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp4xxIn"] = types.YLeaf{"Csbsipmthdhistorystatsresp4Xxin", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp4Xxin}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp4xxOut"] = types.YLeaf{"Csbsipmthdhistorystatsresp4Xxout", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp4Xxout}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp5xxIn"] = types.YLeaf{"Csbsipmthdhistorystatsresp5Xxin", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp5Xxin}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp5xxOut"] = types.YLeaf{"Csbsipmthdhistorystatsresp5Xxout", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp5Xxout}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp6xxIn"] = types.YLeaf{"Csbsipmthdhistorystatsresp6Xxin", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp6Xxin}
    csbsipmthdhistorystatsentry.EntityData.Leafs["csbSIPMthdHistoryStatsResp6xxOut"] = types.YLeaf{"Csbsipmthdhistorystatsresp6Xxout", csbsipmthdhistorystatsentry.Csbsipmthdhistorystatsresp6Xxout}
    return &(csbsipmthdhistorystatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csbsipmthdrccurrentstatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable) GetEntityData() *types.CommonEntityData {
    csbsipmthdrccurrentstatstable.EntityData.YFilter = csbsipmthdrccurrentstatstable.YFilter
    csbsipmthdrccurrentstatstable.EntityData.YangName = "csbSIPMthdRCCurrentStatsTable"
    csbsipmthdrccurrentstatstable.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdrccurrentstatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbsipmthdrccurrentstatstable.EntityData.SegmentPath = "csbSIPMthdRCCurrentStatsTable"
    csbsipmthdrccurrentstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdrccurrentstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdrccurrentstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdrccurrentstatstable.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdrccurrentstatstable.EntityData.Children["csbSIPMthdRCCurrentStatsEntry"] = types.YChild{"Csbsipmthdrccurrentstatsentry", nil}
    for i := range csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry {
        csbsipmthdrccurrentstatstable.EntityData.Children[types.GetSegmentPath(&csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry[i])] = types.YChild{"Csbsipmthdrccurrentstatsentry", &csbsipmthdrccurrentstatstable.Csbsipmthdrccurrentstatsentry[i]}
    }
    csbsipmthdrccurrentstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbsipmthdrccurrentstatstable.EntityData)
}

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

func (csbsipmthdrccurrentstatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrccurrentstatstable_Csbsipmthdrccurrentstatsentry) GetEntityData() *types.CommonEntityData {
    csbsipmthdrccurrentstatsentry.EntityData.YFilter = csbsipmthdrccurrentstatsentry.YFilter
    csbsipmthdrccurrentstatsentry.EntityData.YangName = "csbSIPMthdRCCurrentStatsEntry"
    csbsipmthdrccurrentstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdrccurrentstatsentry.EntityData.ParentYangName = "csbSIPMthdRCCurrentStatsTable"
    csbsipmthdrccurrentstatsentry.EntityData.SegmentPath = "csbSIPMthdRCCurrentStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdRCCurrentStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsadjname) + "']" + "[csbSIPMthdRCCurrentStatsMethod='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsmethod) + "']" + "[csbSIPMthdRCCurrentStatsRespCode='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespcode) + "']" + "[csbSIPMthdRCCurrentStatsInterval='" + fmt.Sprintf("%v", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsinterval) + "']"
    csbsipmthdrccurrentstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdrccurrentstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdrccurrentstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdrccurrentstatsentry.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdrccurrentstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbsipmthdrccurrentstatsentry.Csbcallstatsinstanceindex}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbsipmthdrccurrentstatsentry.Csbcallstatsserviceindex}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbSIPMthdRCCurrentStatsAdjName"] = types.YLeaf{"Csbsipmthdrccurrentstatsadjname", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsadjname}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbSIPMthdRCCurrentStatsMethod"] = types.YLeaf{"Csbsipmthdrccurrentstatsmethod", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsmethod}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbSIPMthdRCCurrentStatsRespCode"] = types.YLeaf{"Csbsipmthdrccurrentstatsrespcode", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespcode}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbSIPMthdRCCurrentStatsInterval"] = types.YLeaf{"Csbsipmthdrccurrentstatsinterval", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsinterval}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbSIPMthdRCCurrentStatsMethodName"] = types.YLeaf{"Csbsipmthdrccurrentstatsmethodname", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsmethodname}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbSIPMthdRCCurrentStatsRespIn"] = types.YLeaf{"Csbsipmthdrccurrentstatsrespin", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespin}
    csbsipmthdrccurrentstatsentry.EntityData.Leafs["csbSIPMthdRCCurrentStatsRespOut"] = types.YLeaf{"Csbsipmthdrccurrentstatsrespout", csbsipmthdrccurrentstatsentry.Csbsipmthdrccurrentstatsrespout}
    return &(csbsipmthdrccurrentstatsentry.EntityData)
}

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
    // CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry.
    Csbsipmthdrchistorystatsentry []CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry
}

func (csbsipmthdrchistorystatstable *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable) GetEntityData() *types.CommonEntityData {
    csbsipmthdrchistorystatstable.EntityData.YFilter = csbsipmthdrchistorystatstable.YFilter
    csbsipmthdrchistorystatstable.EntityData.YangName = "csbSIPMthdRCHistoryStatsTable"
    csbsipmthdrchistorystatstable.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdrchistorystatstable.EntityData.ParentYangName = "CISCO-SESS-BORDER-CTRLR-STATS-MIB"
    csbsipmthdrchistorystatstable.EntityData.SegmentPath = "csbSIPMthdRCHistoryStatsTable"
    csbsipmthdrchistorystatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdrchistorystatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdrchistorystatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdrchistorystatstable.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdrchistorystatstable.EntityData.Children["csbSIPMthdRCHistoryStatsEntry"] = types.YChild{"Csbsipmthdrchistorystatsentry", nil}
    for i := range csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry {
        csbsipmthdrchistorystatstable.EntityData.Children[types.GetSegmentPath(&csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry[i])] = types.YChild{"Csbsipmthdrchistorystatsentry", &csbsipmthdrchistorystatstable.Csbsipmthdrchistorystatsentry[i]}
    }
    csbsipmthdrchistorystatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csbsipmthdrchistorystatstable.EntityData)
}

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

func (csbsipmthdrchistorystatsentry *CISCOSESSBORDERCTRLRSTATSMIB_Csbsipmthdrchistorystatstable_Csbsipmthdrchistorystatsentry) GetEntityData() *types.CommonEntityData {
    csbsipmthdrchistorystatsentry.EntityData.YFilter = csbsipmthdrchistorystatsentry.YFilter
    csbsipmthdrchistorystatsentry.EntityData.YangName = "csbSIPMthdRCHistoryStatsEntry"
    csbsipmthdrchistorystatsentry.EntityData.BundleName = "cisco_ios_xe"
    csbsipmthdrchistorystatsentry.EntityData.ParentYangName = "csbSIPMthdRCHistoryStatsTable"
    csbsipmthdrchistorystatsentry.EntityData.SegmentPath = "csbSIPMthdRCHistoryStatsEntry" + "[csbCallStatsInstanceIndex='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbcallstatsinstanceindex) + "']" + "[csbCallStatsServiceIndex='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbcallstatsserviceindex) + "']" + "[csbSIPMthdRCHistoryStatsAdjName='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsadjname) + "']" + "[csbSIPMthdRCHistoryStatsMethod='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsmethod) + "']" + "[csbSIPMthdRCHistoryStatsRespCode='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespcode) + "']" + "[csbSIPMthdRCHistoryStatsInterval='" + fmt.Sprintf("%v", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsinterval) + "']"
    csbsipmthdrchistorystatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csbsipmthdrchistorystatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csbsipmthdrchistorystatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csbsipmthdrchistorystatsentry.EntityData.Children = make(map[string]types.YChild)
    csbsipmthdrchistorystatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbCallStatsInstanceIndex"] = types.YLeaf{"Csbcallstatsinstanceindex", csbsipmthdrchistorystatsentry.Csbcallstatsinstanceindex}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbCallStatsServiceIndex"] = types.YLeaf{"Csbcallstatsserviceindex", csbsipmthdrchistorystatsentry.Csbcallstatsserviceindex}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbSIPMthdRCHistoryStatsAdjName"] = types.YLeaf{"Csbsipmthdrchistorystatsadjname", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsadjname}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbSIPMthdRCHistoryStatsMethod"] = types.YLeaf{"Csbsipmthdrchistorystatsmethod", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsmethod}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbSIPMthdRCHistoryStatsRespCode"] = types.YLeaf{"Csbsipmthdrchistorystatsrespcode", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespcode}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbSIPMthdRCHistoryStatsInterval"] = types.YLeaf{"Csbsipmthdrchistorystatsinterval", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsinterval}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbSIPMthdRCHistoryStatsMethodName"] = types.YLeaf{"Csbsipmthdrchistorystatsmethodname", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsmethodname}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbSIPMthdRCHistoryStatsRespIn"] = types.YLeaf{"Csbsipmthdrchistorystatsrespin", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespin}
    csbsipmthdrchistorystatsentry.EntityData.Leafs["csbSIPMthdRCHistoryStatsRespOut"] = types.YLeaf{"Csbsipmthdrchistorystatsrespout", csbsipmthdrchistorystatsentry.Csbsipmthdrchistorystatsrespout}
    return &(csbsipmthdrchistorystatsentry.EntityData)
}

