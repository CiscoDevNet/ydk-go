// This module defines a MIB for Round Trip Time
// (RTT) monitoring of a list of targets, using a 
// variety of protocols.
// 
// The table structure overview is a follows (t: 
//  indicates a table, at:  indicates an augmented 
//  table, and it:  indicates table with the same 
//  indices/control as parent table):
// 
// RTTMON MIB
// |--- Application Group
// |    |--- Application Identity
// |    |--- Application Capabilities
// |    |--- Application Reset
// |    |t-- Supported RTT Types
// |         |--- Truth Value
// |    |t-- Supported Protocols
// |         |--- Truth Value
// |    |t-- Application Preconfigured
// |         |--- Script Names
// |         |--- File Paths
// |    |--- Responder control
// |    |t-- Control Protocol Authentication
// |
// |--- Overall Control Group
// |    |t-- Master Definitions Table
// |    |    |--- Global Configuration Definitions
// |    |         |--- Config for a single RTT Life
// |    |    |it- Echo Specific Configuration
// |    |    |it- Echo Path Hop Address Configuration
// |    |    |it- File I/O Specific Configuration
// |    |    |it- Script Specific Configuration
// |    |    |at- Schedule Configuration
// |    |    |at- Reaction Specific Config
// |    |    |at- Statistics Capture Configuration
// |    |    |at- History Collection Configuration
// |    |    |at- Monitoring Operational State
// |    |    |at- Last RTT operation
// |    |
// |    |t-- Reaction Trigger Table
// |         |at- Reaction Trigger Operational State
// |
// |--- Statistics Collection Group
// |    |t-- Statistics Capture Table
// |         |--- Captured Statistics
// |              |--- Path Information
// |              |--- Distribution Capture 
// |              |--- Mean and Deviation Capture
// |         |it- Statistics Collection Table
// |    |it- Statistics Totals Table
// |    |t-- HTTP Stats Table
// |    |t-- Jitter Stats Table
// |
// |--- History Collection Group
// |    |t-- History Collection Table
// |         |-- Path Information
// |         |-- Completion Information per operation
// |
// |--- Latest Operation Group
// |    |t-- Latest HTTP Oper Table
// |    |t-- Latest Jitter Oper Table
// 
// DEFINITIONS:
//   conceptual RTT control row - 
//           This is a row in the 'Overall Control 
//           Group'.  This row is indexed via the 
//           rttMonCtrlAdminIndex object.  This row 
//           is spread across multiple real tables 
//           in the 'Overall Control Group'.
//   probe -
//           This is the entity that executes via a 
//           conceptual RTT control row and populates
//           a conceptual statistics row and a 
//           conceptual history row.
//   Rtt operation -
//           This is a single operation performed by
//           a probe.  This operation can be a single
//           Rtt attempt/completion or a group of Rtt
//           attempts/completions that produce one
//           operation table entry.
// 
// ARR Protocol Definition:
// 
// The format of the RTT Asymmetric Request/Responses 
//  (ARR) protocol is as follows:
// 
//   The ARR Header (total of 12 octets): 
// 
//   4 octet -> eyecatcher: 'WxYz'
//   1 octet -> version   : 0x01 - protocol version
//   1 octet -> command   : 0x01 - logoff request
//                          0x02 - echo request
//                          0x03 - echo response
//                          0x04 - software version request
//                          0x05 - software version response
//   2 octet -> sequence number (Network Byte Order)
//   4 octet -> response data size (Network Byte Order)
// 
//   The ARR Data:
// 
//   n octets -> request/response data
//                         : 'AB..ZAB..ZAB..' 
// 
//   For software version request/response the 
//    protocol version octet will contain the version
//    number of the responder.  Thus the sequence 
//    number, etc will not be included.
// 
//   For snaLU0EchoAppl and snaLU2EchoAppl all character 
//    fields will be in EBCDIC.
// 
//   The response data should be appended to the 
//    origin request data.  This allows data  
//    verification to check the data that flows in 
//    both directions.  If the response data size is
//    smaller than the request data size the original
//    request data will be truncated.  
// 
//   An example would be:
//     Request:        /       Response:
//     'WxYz'          /       'WxYz'
//     0x01            /       0x01
//     0x02            /       0x03
//     0x0001          /       0x0001
//     0x00000008      /       0x00000008
//     'ABCDEF'        /       'ABCDEFGH'
// 
//   NOTE: We requested 8 bytes in the response and 
//         the response had 8 bytes.  The size of the
//         request data has no correlation to the
//         size of the response data.
// 
// NOTE:  For native RTT request/response (i.e. 
//        ipIcmpecho) operations both the 'Header' 
//        and 'Data' will be included.  Only the 
//        'sequence number' in the Header will be 
//        valid.
// 
// NOTE:  For non-connection oriented protocol the 
//        initial RTT request/response operation will
//        be preceded with an RTT request/response 
//        operation to the target address to force 
//        path exploration and to prove 
//        connectivity.  The History collection table
//        will contain these responses, but the 
//        Statistics capture table will omit them to
//        prevent skewed results.
package cisco_rttmon_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_rttmon_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-RTTMON-MIB CISCO-RTTMON-MIB}", reflect.TypeOf(CISCORTTMONMIB{}))
    ydk.RegisterEntity("CISCO-RTTMON-MIB:CISCO-RTTMON-MIB", reflect.TypeOf(CISCORTTMONMIB{}))
}

// CISCORTTMONMIB
type CISCORTTMONMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Rttmonappl CISCORTTMONMIB_Rttmonappl

    // A table of which contains the supported Rtt Monitor Types.  See the
    // RttMonRttType textual convention for the definition of each type.
    Rttmonapplsupportedrtttypestable CISCORTTMONMIB_Rttmonapplsupportedrtttypestable

    // A table of which contains the supported Rtt Monitor Protocols.  See the
    // RttMonProtocol textual convention  for the definition of each protocol.
    Rttmonapplsupportedprotocolstable CISCORTTMONMIB_Rttmonapplsupportedprotocolstable

    // A table of which contains the previously configured Script Names and File
    // IO targets.  These Script Names and File IO targets are installed via a
    // different mechanism than this application, and are specific to each
    // platform.
    Rttmonapplpreconfigedtable CISCORTTMONMIB_Rttmonapplpreconfigedtable

    // A table which contains the definitions for key-strings that will be used in
    // authenticating RTR Control Protocol.
    Rttmonapplauthtable CISCORTTMONMIB_Rttmonapplauthtable

    // A table of Round Trip Time (RTT) monitoring definitions.  The RTT
    // administration control is in multiple tables.   This first table, is used
    // to create a conceptual RTT  control row.  The following tables contain
    // objects which  configure scheduling, information gathering, and 
    // notification/trigger generation.  All of these tables  will create the same
    // conceptual RTT control row as this  table using this tables' index as their
    // own index.   This table is limited in size by the agent  implementation. 
    // The object rttMonApplNumCtrlAdminEntry will reflect this tables maximum
    // number of entries.
    Rttmonctrladmintable CISCORTTMONMIB_Rttmonctrladmintable

    // A table that contains Round Trip Time (RTT) specific definitions.  This
    // table is controlled via the  rttMonCtrlAdminTable.  Entries in this table
    // are created via the rttMonCtrlAdminStatus object.
    Rttmonechoadmintable CISCORTTMONMIB_Rttmonechoadmintable

    // A table of Round Trip Time (RTT) monitoring 'fileIO' specific definitions. 
    // When the RttMonRttType is not 'fileIO' this table is not valid.  This table
    // is controlled via the  rttMonCtrlAdminTable.  Entries in this table are
    // created via the rttMonCtrlAdminStatus object.
    Rttmonfileioadmintable CISCORTTMONMIB_Rttmonfileioadmintable

    // A table of Round Trip Time (RTT) monitoring 'script' specific definitions. 
    // When the RttMonRttType is not 'script' this table is not valid.  This table
    // is controlled via the rttMonCtrlAdminTable.  Entries in this table are
    // created via the rttMonCtrlAdminStatus object.
    Rttmonscriptadmintable CISCORTTMONMIB_Rttmonscriptadmintable

    // A table of which contains the list of conceptual RTT control rows that will
    // start to collect data when a  reaction condition is violated and when 
    // rttMonReactAdminActionType is set to one of the  following:   - 
    // triggerOnly   -  trapAndTrigger   -  nmvtAndTrigger   -  trapNmvtAndTrigger
    // or when a reaction condition is violated and when any of the row in
    // rttMonReactTable has rttMonReactActionType as one of the following:   -
    // triggerOnly   - trapAndTrigger  The goal of this table is to define one or
    // more  additional conceptual RTT control rows that will become active and
    // start to collect additional history and statistics (depending on the rows
    // configuration values), when a problem has been detected.  If the conceptual
    // RTT control row is undefined, and a  trigger occurs, no action will take
    // place.    If the conceptual RTT control row is scheduled to start  at a
    // later time, triggering that row will have no effect.  If the conceptual RTT
    // control row is currently active,  triggering that row will have no effect
    // on that row, but  the rttMonReactTriggerOperState object will transition to
    // 'active'.  An entry in this table can only be triggered when it is not
    // currently in a triggered state.  The object rttMonReactTriggerOperState
    // will  reflect the state of each entry in this table.
    Rttmonreacttriggeradmintable CISCORTTMONMIB_Rttmonreacttriggeradmintable

    // A table to store the hop addresses in a Loose Source Routing path. Response
    // times are computed along the specified path using ping.  This maximum table
    // size is limited by the size of the  maximum number of hop addresses that
    // can fit in an IP header, which is 8. The object rttMonEchoPathAdminEntry
    // will reflect  this tables maximum number of entries.  This table is coupled
    // with rttMonCtrlAdminStatus.
    Rttmonechopathadmintable CISCORTTMONMIB_Rttmonechopathadmintable

    // A table of Round Trip Time (RTT) monitoring group scheduling specific
    // definitions. This table is used to create a conceptual group scheduling
    // control row. The entries in this control row contain objects used to define
    // group schedule configuration parameters.  The objects of this table will be
    // used to schedule a group of probes identified by the conceptual rows of the
    // rttMonCtrlAdminTable.
    Rttmongrpscheduleadmintable CISCORTTMONMIB_Rttmongrpscheduleadmintable

    // A table of Auto SAA L3 MPLS VPN definitions.  The Auto SAA L3 MPLS VPN
    // administration control is in multiple tables.  This first table, is used to
    // create a conceptual Auto SAA L3 MPLS VPN control row.  The following tables
    // contain objects which used in type specific configurations, scheduling and
    // reaction configurations. All of these tables will create the same
    // conceptual control row as this table using this table's index as their own
    // index.  In order to a row in this table to become active the following
    // objects must be defined.   rttMplsVpnMonCtrlRttType,  
    // rttMplsVpnMonCtrlVrfName and   rttMplsVpnMonSchedulePeriod.
    Rttmplsvpnmonctrltable CISCORTTMONMIB_Rttmplsvpnmonctrltable

    // A table that contains the reaction configurations. Each conceptual row in
    // rttMonReactTable corresponds to a reaction configured for the probe defined
    // in rttMonCtrlAdminTable.  For each reaction configured for a probe there is
    // an entry in the table.  Each Probe can have multiple reactions and hence
    // there can be multiple rows for a particular probe.  This table is coupled
    // with rttMonCtrlAdminTable.
    Rttmonreacttable CISCORTTMONMIB_Rttmonreacttable

    // This table contains information about the generated operation id as part of
    // a parent IP SLA operation. The parent operation id is pseudo-random number,
    // selected by the management  station based on an operation started by the
    // management  station,when creating a row via the rttMonCtrlAdminStatus
    // object in the rttMonCtrlAdminTable table.
    Rttmongeneratedopertable CISCORTTMONMIB_Rttmongeneratedopertable

    // The statistics capture database.  The statistics capture table contains
    // summarized  information of the results for a conceptual RTT control  row. 
    // A rolling accumulated history of this information  is maintained in a
    // series of hourly 'group(s)'.  Each  'group' contains a series of 'path(s)',
    // each 'path'  contains a series of 'hop(s)', each 'hop' contains a  series
    // of 'statistics distribution bucket(s)'.  Each conceptual statistics row has
    // a current hourly  group, into which RTT results are accumulated.  At the 
    // end of each hour a new hourly group is created which  then becomes current.
    // The counters and accumulators in  the new group are initialized to zero. 
    // The previous  group(s) is kept in the table until the table contains 
    // rttMonStatisticsAdminNumHourGroups groups for the  conceptual statistics
    // row;  at this point, the oldest  group is discarded and is replaced by the
    // newly created  one.  The hourly group is uniquely identified by the 
    // rttMonStatsCaptureStartTimeIndex object.  If the activity for a conceptual
    // RTT control row ceases  because the rttMonCtrlOperState object transitions
    // to  'inactive', the corresponding current hourly group in  this table is
    // 'frozen', and a new hourly group is  created when activity is resumed.  If
    // the activity for a conceptual RTT control row ceases  because the
    // rttMonCtrlOperState object transitions to  'pending' this whole table will
    // be cleared and reset to  its initial state.  When the RttMonRttType is
    // 'pathEcho', the path  exploration RTT requests' statistics will not be 
    // accumulated in this table.  NOTE: When the RttMonRttType is 'pathEcho', a
    // source to        target rttMonStatsCapturePathIndex path will be       
    // created for each rttMonStatsCaptureStartTimeIndex        to hold all errors
    // that occur when a specific path       had not been found or connection has
    // not be setup.  Using this rttMonStatsCaptureTable, a managing  application
    // can retrieve summarized data from accurately  measured periods, which is
    // synchronized across multiple  conceptual RTT control rows.  With the new
    // hourly group creation being performed on a 60 minute period, the  managing
    // station has plenty of time to collect the data,  and need not be concerned
    // with the vagaries of network  delays and lost PDU's when trying to get
    // matching data.   Also, the managing station can spread the data gathering 
    // over a longer period, which removes the need for a flood  of get requests
    // in a short period which otherwise would  occur.
    Rttmonstatscapturetable CISCORTTMONMIB_Rttmonstatscapturetable

    // The statistics collection database.  This table has the exact same behavior
    // as the rttMonStatsCaptureTable, except it does not keep statistical
    // distribution information.  For a complete table description see the
    // rttMonStatsCaptureTable object.
    Rttmonstatscollecttable CISCORTTMONMIB_Rttmonstatscollecttable

    // The statistics totals database.  This table has the exact same behavior as
    // the rttMonStatsCaptureTable, except it only keeps 60 minute group values. 
    // For a complete table description see the rttMonStatsCaptureTable object.
    Rttmonstatstotalstable CISCORTTMONMIB_Rttmonstatstotalstable

    // The HTTP statistics collection database.  The HTTP statistics table
    // contains summarized information of the results for a conceptual RTT control
    // row. A rolling accumulated history of this information is maintained in a 
    // series of hourly 'group(s)'.  The operation of this table is same as that
    // of  rttMonStatsCaptureTable, except that this table can only  store a
    // maximum of 2 hours of data.
    Rttmonhttpstatstable CISCORTTMONMIB_Rttmonhttpstatstable

    // The Jitter statistics collection database.  The Jitter statistics table
    // contains summarized information of the results for a conceptual RTT control
    // row. A rolling accumulated history of this information is maintained in a 
    // series of hourly 'group(s)'.  The operation of this table is same as that
    // of  rttMonStatsCaptureTable, except that this table will store  2 hours of
    // data.
    Rttmonjitterstatstable CISCORTTMONMIB_Rttmonjitterstatstable

    // The Auto SAA L3 MPLS VPN LPD Group Database.  The LPD Group statistics
    // table contains summarized performance statistics for the LPD group.  LPD
    // Group - The set of 'single probes' which are subset of the 'lspGroup' probe
    // traversing set of paths between two PE end points are grouped together and
    // called as the LPD group. The LPD group will be uniquely referenced by the
    // LPD Group ID.  A rolling accumulated history of this information is
    // maintained in a series of hourly 'group(s)'.  Each conceptual statistics
    // row has a current hourly group, into which RTT results are accumulated. At
    // the end of each hour a new hourly group is created which then becomes
    // current. The counters and accumulators in the new group are initialized to
    // zero. The previous group(s) is kept in the table until the table contains
    // rttMplsVpnMonTypeLpdStatHours groups for the conceptual statistics row;  at
    // this point, the oldest group is discarded and is replaced by the newly
    // created one. The hourly group is uniquely identified by the
    // rttMonLpdGrpStatsStartTimeIndex object.
    Rttmonlpdgrpstatstable CISCORTTMONMIB_Rttmonlpdgrpstatstable

    // The history collection database.  The history table contains a point by
    // point rolling  history of the most recent RTT operations for each 
    // conceptual RTT control row.  The rolling history of this  information is
    // maintained in a series of 'live(s)', each containing a series of
    // 'bucket(s)', each 'bucket'  contains a series of 'sample(s)'.  Each
    // conceptual history row can have lives.  A life is  defined by the
    // rttMonCtrlOperRttLife object.  A new life  will be created when
    // rttMonCtrlOperState transitions 'active'.  When the number of lives become
    // greater  than rttMonHistoryAdminNumLives the oldest life will be  discarded
    // and a new life will be created by incrementing the index.  The path
    // exploration RTT operation will be kept as an entry in this table.
    Rttmonhistorycollectiontable CISCORTTMONMIB_Rttmonhistorycollectiontable

    // A table which contains the status of latest HTTP RTT operation.
    Rttmonlatesthttpopertable CISCORTTMONMIB_Rttmonlatesthttpopertable

    // A table which contains the status of latest Jitter operation.
    Rttmonlatestjitteropertable CISCORTTMONMIB_Rttmonlatestjitteropertable
}

func (cISCORTTMONMIB *CISCORTTMONMIB) GetFilter() yfilter.YFilter { return cISCORTTMONMIB.YFilter }

func (cISCORTTMONMIB *CISCORTTMONMIB) SetFilter(yf yfilter.YFilter) { cISCORTTMONMIB.YFilter = yf }

func (cISCORTTMONMIB *CISCORTTMONMIB) GetGoName(yname string) string {
    if yname == "rttMonAppl" { return "Rttmonappl" }
    if yname == "rttMonApplSupportedRttTypesTable" { return "Rttmonapplsupportedrtttypestable" }
    if yname == "rttMonApplSupportedProtocolsTable" { return "Rttmonapplsupportedprotocolstable" }
    if yname == "rttMonApplPreConfigedTable" { return "Rttmonapplpreconfigedtable" }
    if yname == "rttMonApplAuthTable" { return "Rttmonapplauthtable" }
    if yname == "rttMonCtrlAdminTable" { return "Rttmonctrladmintable" }
    if yname == "rttMonEchoAdminTable" { return "Rttmonechoadmintable" }
    if yname == "rttMonFileIOAdminTable" { return "Rttmonfileioadmintable" }
    if yname == "rttMonScriptAdminTable" { return "Rttmonscriptadmintable" }
    if yname == "rttMonReactTriggerAdminTable" { return "Rttmonreacttriggeradmintable" }
    if yname == "rttMonEchoPathAdminTable" { return "Rttmonechopathadmintable" }
    if yname == "rttMonGrpScheduleAdminTable" { return "Rttmongrpscheduleadmintable" }
    if yname == "rttMplsVpnMonCtrlTable" { return "Rttmplsvpnmonctrltable" }
    if yname == "rttMonReactTable" { return "Rttmonreacttable" }
    if yname == "rttMonGeneratedOperTable" { return "Rttmongeneratedopertable" }
    if yname == "rttMonStatsCaptureTable" { return "Rttmonstatscapturetable" }
    if yname == "rttMonStatsCollectTable" { return "Rttmonstatscollecttable" }
    if yname == "rttMonStatsTotalsTable" { return "Rttmonstatstotalstable" }
    if yname == "rttMonHTTPStatsTable" { return "Rttmonhttpstatstable" }
    if yname == "rttMonJitterStatsTable" { return "Rttmonjitterstatstable" }
    if yname == "rttMonLpdGrpStatsTable" { return "Rttmonlpdgrpstatstable" }
    if yname == "rttMonHistoryCollectionTable" { return "Rttmonhistorycollectiontable" }
    if yname == "rttMonLatestHTTPOperTable" { return "Rttmonlatesthttpopertable" }
    if yname == "rttMonLatestJitterOperTable" { return "Rttmonlatestjitteropertable" }
    return ""
}

func (cISCORTTMONMIB *CISCORTTMONMIB) GetSegmentPath() string {
    return "CISCO-RTTMON-MIB:CISCO-RTTMON-MIB"
}

func (cISCORTTMONMIB *CISCORTTMONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonAppl" {
        return &cISCORTTMONMIB.Rttmonappl
    }
    if childYangName == "rttMonApplSupportedRttTypesTable" {
        return &cISCORTTMONMIB.Rttmonapplsupportedrtttypestable
    }
    if childYangName == "rttMonApplSupportedProtocolsTable" {
        return &cISCORTTMONMIB.Rttmonapplsupportedprotocolstable
    }
    if childYangName == "rttMonApplPreConfigedTable" {
        return &cISCORTTMONMIB.Rttmonapplpreconfigedtable
    }
    if childYangName == "rttMonApplAuthTable" {
        return &cISCORTTMONMIB.Rttmonapplauthtable
    }
    if childYangName == "rttMonCtrlAdminTable" {
        return &cISCORTTMONMIB.Rttmonctrladmintable
    }
    if childYangName == "rttMonEchoAdminTable" {
        return &cISCORTTMONMIB.Rttmonechoadmintable
    }
    if childYangName == "rttMonFileIOAdminTable" {
        return &cISCORTTMONMIB.Rttmonfileioadmintable
    }
    if childYangName == "rttMonScriptAdminTable" {
        return &cISCORTTMONMIB.Rttmonscriptadmintable
    }
    if childYangName == "rttMonReactTriggerAdminTable" {
        return &cISCORTTMONMIB.Rttmonreacttriggeradmintable
    }
    if childYangName == "rttMonEchoPathAdminTable" {
        return &cISCORTTMONMIB.Rttmonechopathadmintable
    }
    if childYangName == "rttMonGrpScheduleAdminTable" {
        return &cISCORTTMONMIB.Rttmongrpscheduleadmintable
    }
    if childYangName == "rttMplsVpnMonCtrlTable" {
        return &cISCORTTMONMIB.Rttmplsvpnmonctrltable
    }
    if childYangName == "rttMonReactTable" {
        return &cISCORTTMONMIB.Rttmonreacttable
    }
    if childYangName == "rttMonGeneratedOperTable" {
        return &cISCORTTMONMIB.Rttmongeneratedopertable
    }
    if childYangName == "rttMonStatsCaptureTable" {
        return &cISCORTTMONMIB.Rttmonstatscapturetable
    }
    if childYangName == "rttMonStatsCollectTable" {
        return &cISCORTTMONMIB.Rttmonstatscollecttable
    }
    if childYangName == "rttMonStatsTotalsTable" {
        return &cISCORTTMONMIB.Rttmonstatstotalstable
    }
    if childYangName == "rttMonHTTPStatsTable" {
        return &cISCORTTMONMIB.Rttmonhttpstatstable
    }
    if childYangName == "rttMonJitterStatsTable" {
        return &cISCORTTMONMIB.Rttmonjitterstatstable
    }
    if childYangName == "rttMonLpdGrpStatsTable" {
        return &cISCORTTMONMIB.Rttmonlpdgrpstatstable
    }
    if childYangName == "rttMonHistoryCollectionTable" {
        return &cISCORTTMONMIB.Rttmonhistorycollectiontable
    }
    if childYangName == "rttMonLatestHTTPOperTable" {
        return &cISCORTTMONMIB.Rttmonlatesthttpopertable
    }
    if childYangName == "rttMonLatestJitterOperTable" {
        return &cISCORTTMONMIB.Rttmonlatestjitteropertable
    }
    return nil
}

func (cISCORTTMONMIB *CISCORTTMONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rttMonAppl"] = &cISCORTTMONMIB.Rttmonappl
    children["rttMonApplSupportedRttTypesTable"] = &cISCORTTMONMIB.Rttmonapplsupportedrtttypestable
    children["rttMonApplSupportedProtocolsTable"] = &cISCORTTMONMIB.Rttmonapplsupportedprotocolstable
    children["rttMonApplPreConfigedTable"] = &cISCORTTMONMIB.Rttmonapplpreconfigedtable
    children["rttMonApplAuthTable"] = &cISCORTTMONMIB.Rttmonapplauthtable
    children["rttMonCtrlAdminTable"] = &cISCORTTMONMIB.Rttmonctrladmintable
    children["rttMonEchoAdminTable"] = &cISCORTTMONMIB.Rttmonechoadmintable
    children["rttMonFileIOAdminTable"] = &cISCORTTMONMIB.Rttmonfileioadmintable
    children["rttMonScriptAdminTable"] = &cISCORTTMONMIB.Rttmonscriptadmintable
    children["rttMonReactTriggerAdminTable"] = &cISCORTTMONMIB.Rttmonreacttriggeradmintable
    children["rttMonEchoPathAdminTable"] = &cISCORTTMONMIB.Rttmonechopathadmintable
    children["rttMonGrpScheduleAdminTable"] = &cISCORTTMONMIB.Rttmongrpscheduleadmintable
    children["rttMplsVpnMonCtrlTable"] = &cISCORTTMONMIB.Rttmplsvpnmonctrltable
    children["rttMonReactTable"] = &cISCORTTMONMIB.Rttmonreacttable
    children["rttMonGeneratedOperTable"] = &cISCORTTMONMIB.Rttmongeneratedopertable
    children["rttMonStatsCaptureTable"] = &cISCORTTMONMIB.Rttmonstatscapturetable
    children["rttMonStatsCollectTable"] = &cISCORTTMONMIB.Rttmonstatscollecttable
    children["rttMonStatsTotalsTable"] = &cISCORTTMONMIB.Rttmonstatstotalstable
    children["rttMonHTTPStatsTable"] = &cISCORTTMONMIB.Rttmonhttpstatstable
    children["rttMonJitterStatsTable"] = &cISCORTTMONMIB.Rttmonjitterstatstable
    children["rttMonLpdGrpStatsTable"] = &cISCORTTMONMIB.Rttmonlpdgrpstatstable
    children["rttMonHistoryCollectionTable"] = &cISCORTTMONMIB.Rttmonhistorycollectiontable
    children["rttMonLatestHTTPOperTable"] = &cISCORTTMONMIB.Rttmonlatesthttpopertable
    children["rttMonLatestJitterOperTable"] = &cISCORTTMONMIB.Rttmonlatestjitteropertable
    return children
}

func (cISCORTTMONMIB *CISCORTTMONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCORTTMONMIB *CISCORTTMONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCORTTMONMIB *CISCORTTMONMIB) GetYangName() string { return "CISCO-RTTMON-MIB" }

func (cISCORTTMONMIB *CISCORTTMONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCORTTMONMIB *CISCORTTMONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCORTTMONMIB *CISCORTTMONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCORTTMONMIB *CISCORTTMONMIB) SetParent(parent types.Entity) { cISCORTTMONMIB.parent = parent }

func (cISCORTTMONMIB *CISCORTTMONMIB) GetParent() types.Entity { return cISCORTTMONMIB.parent }

func (cISCORTTMONMIB *CISCORTTMONMIB) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonappl
type CISCORTTMONMIB_Rttmonappl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Round Trip Time monitoring application version string.  The format will be:
    // 'Version.Release.Patch-Level: Textual-Description'  For example:  '1.0.0:
    // Initial RTT Application'. The type is string.
    Rttmonapplversion interface{}

    // The maximum size of the data portion an echo packet supported by this RTT
    // application.  This is the maximum value that can be specified by
    // (rttMonEchoAdminPktDataRequestSize + ARR Header) or 
    // (rttMonEchoAdminPktDataResponseSize + ARR Header) in the
    // rttMonCtrlAdminTable.  This object is undefined for conceptual RTT  control
    // rows when the RttMonRttType object is set to 'fileIO' or 'script'. The type
    // is interface{} with range: 0..16384. Units are octets.
    Rttmonapplmaxpacketdatasize interface{}

    // The last time at which a set operation occurred on any of the objects in
    // this MIB.  The managing  application can inspect this value in order to 
    // determine whether changes have been made without  retrieving the entire
    // Administration portion of this MIB.  This object applies to all settable
    // objects in this MIB, including the 'Reset' objects that could clear saved
    // history/statistics. The type is interface{} with range: 0..4294967295.
    Rttmonappltimeoflastset interface{}

    // This object defines the maximum number of entries that can be added to the
    // rttMonCtrlAdminTable. It is calculated at the system init time. The value
    // is impacted when rttMonApplFreeMemLowWaterMark is changed. The type is
    // interface{} with range: 1..2147483647.
    Rttmonapplnumctrladminentry interface{}

    // When set to 'reset' the entire RTT application goes through a reset
    // sequence, making a best  effort to revert to its startup condition.  Any 
    // and all rows in the Overall Control Group will be immediately deleted,
    // together with any associated rows in the Statistics Collection Group, and 
    // History Collection Group.  All open connections  will also be closed. 
    // Finally the  rttMonApplPreConfigedTable will reset (see 
    // rttMonApplPreConfigedReset). The type is RttReset.
    Rttmonapplreset interface{}

    // When set to 'reset' the RTT application will reset the Application
    // Preconfigured MIB section.  This will force the RTT application to delete
    // all entries in the rttMonApplPreConfigedTable and then to repopulate the
    // table with the current configuration.  This provides a mechanism to load
    // and unload user scripts and file paths. The type is RttReset.
    Rttmonapplpreconfigedreset interface{}

    // This object defines the number of new probes that can be configured on a
    // router. The number depends on the value  of rttMonApplFreeMemLowWaterMark,
    // free bytes available on the router and the system configured
    // rttMonCtrlAdminEntry number. Equation: rttMonApplProbeCapacity = 
    // MIN(((Free_Bytes_on_the_Router - rttMonApplFreeMemLowWaterMark)/
    // Memory_required_by_each_probe), rttMonApplNumCtrlAdminEntry - 
    // Num_of_Probes_already_configured)). The type is interface{} with range:
    // 1..2147483647.
    Rttmonapplprobecapacity interface{}

    // This object defines the amount of free memory a router must have in order
    // to configure RTR. If RTR found out that the memory is falling below this
    // mark, it will not allow new probes to be configured.  This value should not
    // be set higher (or very close to) than  the free bytes available on the
    // router. The type is interface{} with range: 0..2147483647.
    Rttmonapplfreememlowwatermark interface{}

    // An error description for the last error message caused by set.  Currently,
    // it includes set error caused due to setting rttMonApplFreeMemLowWaterMark
    // greater than the available free memory on the router or not enough memory
    // left to create new probes. The type is string.
    Rttmonappllatestseterror interface{}

    // Enable or disable RTR responder on the router. The type is bool.
    Rttmonapplresponder interface{}

    // This object is used to reset certain objects within the
    // rttMonLpdGrpStatsTable.  When the object is set to value of an active LPD
    // Group identifier the associated objects will be reset. The reset objects
    // will be set to a value as specified in the object's description.  The
    // following objects will not be reset. - rttMonLpdGrpStatsTargetPE -
    // rttMonLpdGrpStatsGroupProbeIndex - rttMonLpdGrpStatsGroupIndex -
    // rttMonLpdGrpStatsStartTimeIndex. The type is interface{} with range:
    // 0..2147483647.
    Rttmonappllpdgrpstatsreset interface{}
}

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetFilter() yfilter.YFilter { return rttmonappl.YFilter }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) SetFilter(yf yfilter.YFilter) { rttmonappl.YFilter = yf }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetGoName(yname string) string {
    if yname == "rttMonApplVersion" { return "Rttmonapplversion" }
    if yname == "rttMonApplMaxPacketDataSize" { return "Rttmonapplmaxpacketdatasize" }
    if yname == "rttMonApplTimeOfLastSet" { return "Rttmonappltimeoflastset" }
    if yname == "rttMonApplNumCtrlAdminEntry" { return "Rttmonapplnumctrladminentry" }
    if yname == "rttMonApplReset" { return "Rttmonapplreset" }
    if yname == "rttMonApplPreConfigedReset" { return "Rttmonapplpreconfigedreset" }
    if yname == "rttMonApplProbeCapacity" { return "Rttmonapplprobecapacity" }
    if yname == "rttMonApplFreeMemLowWaterMark" { return "Rttmonapplfreememlowwatermark" }
    if yname == "rttMonApplLatestSetError" { return "Rttmonappllatestseterror" }
    if yname == "rttMonApplResponder" { return "Rttmonapplresponder" }
    if yname == "rttMonApplLpdGrpStatsReset" { return "Rttmonappllpdgrpstatsreset" }
    return ""
}

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetSegmentPath() string {
    return "rttMonAppl"
}

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonApplVersion"] = rttmonappl.Rttmonapplversion
    leafs["rttMonApplMaxPacketDataSize"] = rttmonappl.Rttmonapplmaxpacketdatasize
    leafs["rttMonApplTimeOfLastSet"] = rttmonappl.Rttmonappltimeoflastset
    leafs["rttMonApplNumCtrlAdminEntry"] = rttmonappl.Rttmonapplnumctrladminentry
    leafs["rttMonApplReset"] = rttmonappl.Rttmonapplreset
    leafs["rttMonApplPreConfigedReset"] = rttmonappl.Rttmonapplpreconfigedreset
    leafs["rttMonApplProbeCapacity"] = rttmonappl.Rttmonapplprobecapacity
    leafs["rttMonApplFreeMemLowWaterMark"] = rttmonappl.Rttmonapplfreememlowwatermark
    leafs["rttMonApplLatestSetError"] = rttmonappl.Rttmonappllatestseterror
    leafs["rttMonApplResponder"] = rttmonappl.Rttmonapplresponder
    leafs["rttMonApplLpdGrpStatsReset"] = rttmonappl.Rttmonappllpdgrpstatsreset
    return leafs
}

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetYangName() string { return "rttMonAppl" }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) SetParent(parent types.Entity) { rttmonappl.parent = parent }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetParent() types.Entity { return rttmonappl.parent }

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonapplsupportedrtttypestable
// A table of which contains the supported Rtt
// Monitor Types.
// 
// See the RttMonRttType textual convention for
// the definition of each type.
type CISCORTTMONMIB_Rttmonapplsupportedrtttypestable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list that presents the valid Rtt Monitor Types. The type is slice of
    // CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry.
    Rttmonapplsupportedrtttypesentry []CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry
}

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetFilter() yfilter.YFilter { return rttmonapplsupportedrtttypestable.YFilter }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) SetFilter(yf yfilter.YFilter) { rttmonapplsupportedrtttypestable.YFilter = yf }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetGoName(yname string) string {
    if yname == "rttMonApplSupportedRttTypesEntry" { return "Rttmonapplsupportedrtttypesentry" }
    return ""
}

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetSegmentPath() string {
    return "rttMonApplSupportedRttTypesTable"
}

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonApplSupportedRttTypesEntry" {
        for _, c := range rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry {
            if rttmonapplsupportedrtttypestable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry{}
        rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry = append(rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry, child)
        return &rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry[len(rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry)-1]
    }
    return nil
}

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry {
        children[rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry[i].GetSegmentPath()] = &rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry[i]
    }
    return children
}

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetYangName() string { return "rttMonApplSupportedRttTypesTable" }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) SetParent(parent types.Entity) { rttmonapplsupportedrtttypestable.parent = parent }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetParent() types.Entity { return rttmonapplsupportedrtttypestable.parent }

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry
// A list that presents the valid Rtt Monitor
// Types.
type CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indexes the supported 'RttMonRttType'
    // types. The type is RttMonRttType.
    Rttmonapplsupportedrtttypes interface{}

    // This object defines the supported 'RttMonRttType' types. The type is bool.
    Rttmonapplsupportedrtttypesvalid interface{}
}

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetFilter() yfilter.YFilter { return rttmonapplsupportedrtttypesentry.YFilter }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) SetFilter(yf yfilter.YFilter) { rttmonapplsupportedrtttypesentry.YFilter = yf }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetGoName(yname string) string {
    if yname == "rttMonApplSupportedRttTypes" { return "Rttmonapplsupportedrtttypes" }
    if yname == "rttMonApplSupportedRttTypesValid" { return "Rttmonapplsupportedrtttypesvalid" }
    return ""
}

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetSegmentPath() string {
    return "rttMonApplSupportedRttTypesEntry" + "[rttMonApplSupportedRttTypes='" + fmt.Sprintf("%v", rttmonapplsupportedrtttypesentry.Rttmonapplsupportedrtttypes) + "']"
}

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonApplSupportedRttTypes"] = rttmonapplsupportedrtttypesentry.Rttmonapplsupportedrtttypes
    leafs["rttMonApplSupportedRttTypesValid"] = rttmonapplsupportedrtttypesentry.Rttmonapplsupportedrtttypesvalid
    return leafs
}

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetYangName() string { return "rttMonApplSupportedRttTypesEntry" }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) SetParent(parent types.Entity) { rttmonapplsupportedrtttypesentry.parent = parent }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetParent() types.Entity { return rttmonapplsupportedrtttypesentry.parent }

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetParentYangName() string { return "rttMonApplSupportedRttTypesTable" }

// CISCORTTMONMIB_Rttmonapplsupportedprotocolstable
// A table of which contains the supported Rtt
// Monitor Protocols.
// 
// See the RttMonProtocol textual convention 
// for the definition of each protocol.
type CISCORTTMONMIB_Rttmonapplsupportedprotocolstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list that presents the valid Rtt Monitor Protocols. The type is slice of
    // CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry.
    Rttmonapplsupportedprotocolsentry []CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry
}

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetFilter() yfilter.YFilter { return rttmonapplsupportedprotocolstable.YFilter }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) SetFilter(yf yfilter.YFilter) { rttmonapplsupportedprotocolstable.YFilter = yf }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetGoName(yname string) string {
    if yname == "rttMonApplSupportedProtocolsEntry" { return "Rttmonapplsupportedprotocolsentry" }
    return ""
}

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetSegmentPath() string {
    return "rttMonApplSupportedProtocolsTable"
}

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonApplSupportedProtocolsEntry" {
        for _, c := range rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry {
            if rttmonapplsupportedprotocolstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry{}
        rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry = append(rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry, child)
        return &rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry[len(rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry)-1]
    }
    return nil
}

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry {
        children[rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry[i].GetSegmentPath()] = &rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry[i]
    }
    return children
}

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetYangName() string { return "rttMonApplSupportedProtocolsTable" }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) SetParent(parent types.Entity) { rttmonapplsupportedprotocolstable.parent = parent }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetParent() types.Entity { return rttmonapplsupportedprotocolstable.parent }

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry
// A list that presents the valid Rtt Monitor
// Protocols.
type CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indexes the supported 'RttMonProtocol'
    // protocols. The type is RttMonProtocol.
    Rttmonapplsupportedprotocols interface{}

    // This object defines the supported 'RttMonProtocol' protocols. The type is
    // bool.
    Rttmonapplsupportedprotocolsvalid interface{}
}

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetFilter() yfilter.YFilter { return rttmonapplsupportedprotocolsentry.YFilter }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) SetFilter(yf yfilter.YFilter) { rttmonapplsupportedprotocolsentry.YFilter = yf }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetGoName(yname string) string {
    if yname == "rttMonApplSupportedProtocols" { return "Rttmonapplsupportedprotocols" }
    if yname == "rttMonApplSupportedProtocolsValid" { return "Rttmonapplsupportedprotocolsvalid" }
    return ""
}

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetSegmentPath() string {
    return "rttMonApplSupportedProtocolsEntry" + "[rttMonApplSupportedProtocols='" + fmt.Sprintf("%v", rttmonapplsupportedprotocolsentry.Rttmonapplsupportedprotocols) + "']"
}

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonApplSupportedProtocols"] = rttmonapplsupportedprotocolsentry.Rttmonapplsupportedprotocols
    leafs["rttMonApplSupportedProtocolsValid"] = rttmonapplsupportedprotocolsentry.Rttmonapplsupportedprotocolsvalid
    return leafs
}

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetYangName() string { return "rttMonApplSupportedProtocolsEntry" }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) SetParent(parent types.Entity) { rttmonapplsupportedprotocolsentry.parent = parent }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetParent() types.Entity { return rttmonapplsupportedprotocolsentry.parent }

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetParentYangName() string { return "rttMonApplSupportedProtocolsTable" }

// CISCORTTMONMIB_Rttmonapplpreconfigedtable
// A table of which contains the previously
// configured Script Names and File IO targets.
// 
// These Script Names and File IO targets are installed
// via a different mechanism than this application, and
// are specific to each platform.
type CISCORTTMONMIB_Rttmonapplpreconfigedtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that describe the previously configured Script Names and
    // File IO targets. The type is slice of
    // CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry.
    Rttmonapplpreconfigedentry []CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry
}

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetFilter() yfilter.YFilter { return rttmonapplpreconfigedtable.YFilter }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) SetFilter(yf yfilter.YFilter) { rttmonapplpreconfigedtable.YFilter = yf }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetGoName(yname string) string {
    if yname == "rttMonApplPreConfigedEntry" { return "Rttmonapplpreconfigedentry" }
    return ""
}

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetSegmentPath() string {
    return "rttMonApplPreConfigedTable"
}

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonApplPreConfigedEntry" {
        for _, c := range rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry {
            if rttmonapplpreconfigedtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry{}
        rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry = append(rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry, child)
        return &rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry[len(rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry)-1]
    }
    return nil
}

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry {
        children[rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry[i].GetSegmentPath()] = &rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry[i]
    }
    return children
}

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetYangName() string { return "rttMonApplPreConfigedTable" }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) SetParent(parent types.Entity) { rttmonapplpreconfigedtable.parent = parent }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetParent() types.Entity { return rttmonapplpreconfigedtable.parent }

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry
// A list of objects that describe the previously
// configured Script Names and File IO targets.
type CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is the type of value being stored in the
    // rttMonApplPreConfigedName object. The type is Rttmonapplpreconfigedtype.
    Rttmonapplpreconfigedtype interface{}

    // This attribute is a key. This is either one of the following depending on
    // the value of the rttMonApplPreConfigedType object:   - The file path to a
    // server.  One of these file paths     must be used when defining an entry in
    // the     rttMonFileIOAdminTable table with 'fileIO' as the     value of the
    // rttMonCtrlAdminRttType object.   - The script name to be used when
    // generating RTT     operations.  One of these script names must be used    
    // when defining an entry in the rttMonScriptAdminTable     table with
    // 'script' as the value of the     rttMonCtrlAdminRttType object.  NOTE:  For
    // script names, command line parameters         can follow these names in the
    // rttMonScriptAdminTable table. The type is string.
    Rttmonapplpreconfigedname interface{}

    // When this row exists, this value will be 'true'. This object exists only to
    // create a valid row in this  table. The type is bool.
    Rttmonapplpreconfigedvalid interface{}
}

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetFilter() yfilter.YFilter { return rttmonapplpreconfigedentry.YFilter }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) SetFilter(yf yfilter.YFilter) { rttmonapplpreconfigedentry.YFilter = yf }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetGoName(yname string) string {
    if yname == "rttMonApplPreConfigedType" { return "Rttmonapplpreconfigedtype" }
    if yname == "rttMonApplPreConfigedName" { return "Rttmonapplpreconfigedname" }
    if yname == "rttMonApplPreConfigedValid" { return "Rttmonapplpreconfigedvalid" }
    return ""
}

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetSegmentPath() string {
    return "rttMonApplPreConfigedEntry" + "[rttMonApplPreConfigedType='" + fmt.Sprintf("%v", rttmonapplpreconfigedentry.Rttmonapplpreconfigedtype) + "']" + "[rttMonApplPreConfigedName='" + fmt.Sprintf("%v", rttmonapplpreconfigedentry.Rttmonapplpreconfigedname) + "']"
}

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonApplPreConfigedType"] = rttmonapplpreconfigedentry.Rttmonapplpreconfigedtype
    leafs["rttMonApplPreConfigedName"] = rttmonapplpreconfigedentry.Rttmonapplpreconfigedname
    leafs["rttMonApplPreConfigedValid"] = rttmonapplpreconfigedentry.Rttmonapplpreconfigedvalid
    return leafs
}

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetYangName() string { return "rttMonApplPreConfigedEntry" }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) SetParent(parent types.Entity) { rttmonapplpreconfigedentry.parent = parent }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetParent() types.Entity { return rttmonapplpreconfigedentry.parent }

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetParentYangName() string { return "rttMonApplPreConfigedTable" }

// CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry_Rttmonapplpreconfigedtype represents rttMonApplPreConfigedName object.
type CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry_Rttmonapplpreconfigedtype string

const (
    CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry_Rttmonapplpreconfigedtype_filePath CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry_Rttmonapplpreconfigedtype = "filePath"

    CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry_Rttmonapplpreconfigedtype_scriptName CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry_Rttmonapplpreconfigedtype = "scriptName"
)

// CISCORTTMONMIB_Rttmonapplauthtable
// A table which contains the definitions for key-strings
// that will be used in authenticating RTR Control Protocol.
type CISCORTTMONMIB_Rttmonapplauthtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list that presents the valid parameters for Authenticating RTR Control
    // Protocol. The type is slice of
    // CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry.
    Rttmonapplauthentry []CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry
}

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetFilter() yfilter.YFilter { return rttmonapplauthtable.YFilter }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) SetFilter(yf yfilter.YFilter) { rttmonapplauthtable.YFilter = yf }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetGoName(yname string) string {
    if yname == "rttMonApplAuthEntry" { return "Rttmonapplauthentry" }
    return ""
}

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetSegmentPath() string {
    return "rttMonApplAuthTable"
}

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonApplAuthEntry" {
        for _, c := range rttmonapplauthtable.Rttmonapplauthentry {
            if rttmonapplauthtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry{}
        rttmonapplauthtable.Rttmonapplauthentry = append(rttmonapplauthtable.Rttmonapplauthentry, child)
        return &rttmonapplauthtable.Rttmonapplauthentry[len(rttmonapplauthtable.Rttmonapplauthentry)-1]
    }
    return nil
}

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonapplauthtable.Rttmonapplauthentry {
        children[rttmonapplauthtable.Rttmonapplauthentry[i].GetSegmentPath()] = &rttmonapplauthtable.Rttmonapplauthentry[i]
    }
    return children
}

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetYangName() string { return "rttMonApplAuthTable" }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) SetParent(parent types.Entity) { rttmonapplauthtable.parent = parent }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetParent() types.Entity { return rttmonapplauthtable.parent }

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry
// A list that presents the valid parameters for Authenticating
// RTR Control Protocol.
type CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonApplAuthTable. This is a pseudo-random number selected by the
    // management station when creating a row via the rttMonApplAuthStatus 
    // object. If the pseudo-random number is already in use, an 
    // 'inconsistentValue' is returned. Currently, only one row  can be created.
    // The type is interface{} with range: 0..2147483647.
    Rttmonapplauthindex interface{}

    // A string which represents the key-chain name. If multiple key-strings are
    // specified, then the authenticator will  alternate between the specified
    // strings. The type is string with length: 1..48.
    Rttmonapplauthkeychain interface{}

    // A string which represents a key-string name whose id is 1. The type is
    // string with length: 1..48.
    Rttmonapplauthkeystring1 interface{}

    // A string which represents a key-string name whose id is 2. The type is
    // string with length: 1..48.
    Rttmonapplauthkeystring2 interface{}

    // A string which represents a key-string name whose id is 3. The type is
    // string with length: 1..48.
    Rttmonapplauthkeystring3 interface{}

    // A string which represents a key-string name whose id is 4. The type is
    // string with length: 1..48.
    Rttmonapplauthkeystring4 interface{}

    // A string which represents a key-string name whose id is 5. The type is
    // string with length: 1..48.
    Rttmonapplauthkeystring5 interface{}

    // The status of the Authentication row. The type is RowStatus.
    Rttmonapplauthstatus interface{}
}

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetFilter() yfilter.YFilter { return rttmonapplauthentry.YFilter }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) SetFilter(yf yfilter.YFilter) { rttmonapplauthentry.YFilter = yf }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetGoName(yname string) string {
    if yname == "rttMonApplAuthIndex" { return "Rttmonapplauthindex" }
    if yname == "rttMonApplAuthKeyChain" { return "Rttmonapplauthkeychain" }
    if yname == "rttMonApplAuthKeyString1" { return "Rttmonapplauthkeystring1" }
    if yname == "rttMonApplAuthKeyString2" { return "Rttmonapplauthkeystring2" }
    if yname == "rttMonApplAuthKeyString3" { return "Rttmonapplauthkeystring3" }
    if yname == "rttMonApplAuthKeyString4" { return "Rttmonapplauthkeystring4" }
    if yname == "rttMonApplAuthKeyString5" { return "Rttmonapplauthkeystring5" }
    if yname == "rttMonApplAuthStatus" { return "Rttmonapplauthstatus" }
    return ""
}

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetSegmentPath() string {
    return "rttMonApplAuthEntry" + "[rttMonApplAuthIndex='" + fmt.Sprintf("%v", rttmonapplauthentry.Rttmonapplauthindex) + "']"
}

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonApplAuthIndex"] = rttmonapplauthentry.Rttmonapplauthindex
    leafs["rttMonApplAuthKeyChain"] = rttmonapplauthentry.Rttmonapplauthkeychain
    leafs["rttMonApplAuthKeyString1"] = rttmonapplauthentry.Rttmonapplauthkeystring1
    leafs["rttMonApplAuthKeyString2"] = rttmonapplauthentry.Rttmonapplauthkeystring2
    leafs["rttMonApplAuthKeyString3"] = rttmonapplauthentry.Rttmonapplauthkeystring3
    leafs["rttMonApplAuthKeyString4"] = rttmonapplauthentry.Rttmonapplauthkeystring4
    leafs["rttMonApplAuthKeyString5"] = rttmonapplauthentry.Rttmonapplauthkeystring5
    leafs["rttMonApplAuthStatus"] = rttmonapplauthentry.Rttmonapplauthstatus
    return leafs
}

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetYangName() string { return "rttMonApplAuthEntry" }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) SetParent(parent types.Entity) { rttmonapplauthentry.parent = parent }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetParent() types.Entity { return rttmonapplauthentry.parent }

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetParentYangName() string { return "rttMonApplAuthTable" }

// CISCORTTMONMIB_Rttmonctrladmintable
// A table of Round Trip Time (RTT) monitoring definitions.
// 
// The RTT administration control is in multiple tables.  
// This first table, is used to create a conceptual RTT 
// control row.  The following tables contain objects which 
// configure scheduling, information gathering, and 
// notification/trigger generation.  All of these tables 
// will create the same conceptual RTT control row as this 
// table using this tables' index as their own index. 
// 
// This table is limited in size by the agent 
// implementation.  The object rttMonApplNumCtrlAdminEntry
// will reflect this tables maximum number of entries.
type CISCORTTMONMIB_Rttmonctrladmintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual RTT control row. The type
    // is slice of CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry.
    Rttmonctrladminentry []CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry
}

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetFilter() yfilter.YFilter { return rttmonctrladmintable.YFilter }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) SetFilter(yf yfilter.YFilter) { rttmonctrladmintable.YFilter = yf }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminEntry" { return "Rttmonctrladminentry" }
    return ""
}

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetSegmentPath() string {
    return "rttMonCtrlAdminTable"
}

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonCtrlAdminEntry" {
        for _, c := range rttmonctrladmintable.Rttmonctrladminentry {
            if rttmonctrladmintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry{}
        rttmonctrladmintable.Rttmonctrladminentry = append(rttmonctrladmintable.Rttmonctrladminentry, child)
        return &rttmonctrladmintable.Rttmonctrladminentry[len(rttmonctrladmintable.Rttmonctrladminentry)-1]
    }
    return nil
}

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonctrladmintable.Rttmonctrladminentry {
        children[rttmonctrladmintable.Rttmonctrladminentry[i].GetSegmentPath()] = &rttmonctrladmintable.Rttmonctrladminentry[i]
    }
    return children
}

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetYangName() string { return "rttMonCtrlAdminTable" }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) SetParent(parent types.Entity) { rttmonctrladmintable.parent = parent }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetParent() types.Entity { return rttmonctrladmintable.parent }

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry
// A base list of objects that define a conceptual RTT
// control row.
type CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonCtrlAdminTable. This is a pseudo-random number, selected by the
    // management  station or auto-generated based on  operation started by the 
    // management station,when creating a row via  the rttMonCtrlAdminStatus
    // object.  If the pseudo-random   number is already in use an
    // 'inconsistentValue' return code   will be returned when set operation is
    // attempted. The type is interface{} with range: 1..2147483647.
    Rttmonctrladminindex interface{}

    // Identifies the entity that created this table row. The type is string with
    // length: 0..255.
    Rttmonctrladminowner interface{}

    // A string which is used by a managing application to identify the RTT
    // target.  This string is inserted into trap notifications, but has no other
    // significance to the  agent. The type is string with length: 0..16.
    Rttmonctrladmintag interface{}

    // The type of RTT operation to be performed.  This value must be set in the
    // same PDU or before setting any type specific configuration.  Note: The RTT
    // operation 'lspGroup' cannot be created via this control row. It will be
    // created automatically by Auto SAA L3 MPLS VPN when rttMplsVpnMonCtrlLpd is
    // 'true'. The type is RttMonRttType.
    Rttmonctrladminrtttype interface{}

    // This object defines an administrative threshold limit. If the RTT operation
    // time exceeds this limit and if the  conditions specified in
    // rttMonReactAdminThresholdType or  rttMonHistoryAdminFilter are satisfied, a
    // threshold is generated. The type is interface{} with range: 0..2147483647.
    // Units are milliseconds.
    Rttmonctrladminthreshold interface{}

    // Specifies the duration between initiating each RTT operation.   This object
    // cannot be set to a value which would be a  shorter duration than
    // rttMonCtrlAdminTimeout.  When the RttMonRttType specifies an operation that
    // is synchronous in nature, it may happen that the next RTT  operation is
    // blocked by a RTT operation which has not yet completed.  In this case, the
    // value of a counter (rttMonStatsCollectBusies) in rttMonStatsCaptureTable is
    // incremented in lieu of initiating a RTT operation, and  the next attempt
    // will occur at the next rttMonCtrlAdminFrequency expiration.   NOTE:  When
    // the rttMonCtrlAdminRttType object is defined         to be 'pathEcho',
    // setting this value to a small        value for your network size may cause
    // an operation        attempt (or multiple attempts) to be started        
    // before the previous operation has finished.  In         this situation the
    // rttMonStatsCollectBusies object        will be incremented in lieu of
    // initiating a new         RTT operation, and the next attempt will occur at 
    // the next rttMonCtrlAdminFrequency expiration.  When the
    // rttMonCtrlAdminRttType object is defined to be 'pathEcho', the suggested
    // value for this object  is greater than rttMonCtrlAdminTimeout times the 
    // maximum number of expected hops to the target.  NOTE:  When the
    // rttMonCtrlAdminRttType object is defined         to be 'dhcp', the minimum
    // allowed value for this        object is 10 seconds.  This restriction is
    // due to        protocol limitations described in RFC 2131. The type is
    // interface{} with range: 0..604800. Units are seconds.
    Rttmonctrladminfrequency interface{}

    // Specifies the duration to wait for a RTT operation completion.  The value
    // of this object cannot be set to  a value which would specify a duration
    // exceeding  rttMonCtrlAdminFrequency.  For connection oriented protocols,
    // this may cause the connection to be closed by the probe.  Once closed, it
    // will be assumed that the connection reestablishment will be performed.  To
    // prevent unwanted closure of connections, be sure to set this value to a
    // realistic connection timeout. The type is interface{} with range:
    // 0..604800000. Units are milliseconds.
    Rttmonctrladmintimeout interface{}

    // When set to true, the resulting data in each RTT operation is compared with
    // the expected data.  This includes checking header information (if possible)
    // and exact packet size.  Any mismatch will be recorded in the
    // rttMonStatsCollectVerifyErrors object.  Some RttMonRttTypes may not support
    // this option.  When a type does not support this option, the agent will 
    // transition this object to false.  It is the management applications
    // responsibility to check for this  transition. The type is bool.
    Rttmonctrladminverifydata interface{}

    // The status of the conceptual RTT control row.  In order for this object to
    // become active, the following  row objects must be defined:    -
    // rttMonCtrlAdminRttType Additionally:  - for echo, pathEcho based on
    // 'ipIcmpEcho' and dlsw probes     rttMonEchoAdminProtocol and     
    // rttMonEchoAdminTargetAddress;  - for echo, pathEcho based on
    // 'mplsLspPingAppl'     rttMonEchoAdminProtocol, rttMonEchoAdminTargetAddress
    // and rttMonEchoAdminLSPFECType  - for udpEcho, tcpConnect and jitter probes 
    // rttMonEchoAdminTargetAddress and     rttMonEchoAdminTargetPort  - for http
    // and ftp probe     rttMonEchoAdminURL   - for dns probe    
    // rttMonEchoAdminTargetAddressString      rttMonEchoAdminNameServer   - dhcp
    // probe doesn't require any additional objects  All other objects can assume
    // default values. The  conceptual Rtt control row will be placed into a 
    // 'pending' state (via the rttMonCtrlOperState object) if
    // rttMonScheduleAdminRttStartTime is not specified.  Most conceptual Rtt
    // control row objects cannot be  modified once this conceptual Rtt control
    // row has been  created.  The objects that can change are the following:   -
    // Objects in the rttMonReactAdminTable can be modified    as needed without
    // setting this object to     'notInService'.  - Objects in the
    // rttMonScheduleAdminTable can be     modified only when this object has the
    // value of    'notInService'.  - The rttMonCtrlOperState can be modified to
    // control    the state of the probe.  Once this object is in 'active' status,
    // it cannot be  set to 'notInService' while the rttMonCtrlOperState is in
    // 'active' state.  Thus the rttMonCtrlOperState  object must be transitioned
    // first.   This object can be set to 'destroy' from any value at any time.
    // The type is RowStatus.
    Rttmonctrladminstatus interface{}

    // When set to true, this entry will be shown in 'show running' command and
    // can be saved into Non-volatile memory. The type is bool.
    Rttmonctrladminnvgen interface{}

    // If the operation is created through auto measure group creation, then this
    // string will specify the group name to which this operation is associated.
    // The type is string with length: 0..64.
    Rttmonctrladmingroupname interface{}

    // This object value will be placed into the rttMonCtrlOperRttLife object when
    // the rttMonCtrlOperState object transitions to 'active' or 'pending'.  The
    // value 2147483647 has a special meaning.  When this object is set to
    // 2147483647, the  rttMonCtrlOperRttLife object will not decrement.   And
    // thus the life time will never end. The type is interface{} with range:
    // 0..2147483647. Units are seconds.
    Rttmonscheduleadminrttlife interface{}

    // This is the time when this conceptional row will activate.    This is the
    // value of MIB-II's sysUpTime in the future. When sysUpTime equals this value
    // this object will  cause the activation of a conceptual Rtt row.  When an
    // agent has the capability to determine date and  time, the agent should
    // store this object as DateAndTime. This allows the agent to completely reset
    // (restart) and still be able to start conceptual Rtt rows at the  intended
    // time.  If the agent cannot keep date and time and the agent resets, all
    // entries should take on one of the special value defined below.  The first
    // special value allows this conceptual Rtt  control row to immediately
    // transition the  rttMonCtrlOperState object into 'active' state when the
    // rttMonCtrlAdminStatus  object transitions to active. This special value is
    // defined to be a value of this object that, when initially set, is 1.  The
    // second special value allows this conceptual Rtt  control row to immediately
    // transition the  rttMonCtrlOperState object into 'pending' state when  the
    // rttMonCtrlAdminStatus object transitions to active.   Also, when the
    // rttMonCtrlOperRttLife counts down to zero  (and not when set to zero), this
    // special value causes  this conceptual Rtt control row to  retransition the 
    // rttMonCtrlOperState object into 'pending' state.  This  special value is
    // defined to be a value of this object  that, when initially set, is smaller
    // than the current sysUpTime. (With the exception of one, as defined in the
    // previous paragraph). The type is interface{} with range: 0..4294967295.
    Rttmonscheduleadminrttstarttime interface{}

    // The amount of time this conceptual Rtt control row will exist when not in
    // an 'active' rttMonCtrlOperState.  When this conceptual Rtt control row
    // enters an 'active'  state, this timer will be reset and suspended.  When 
    // this conceptual RTT control row enters a state other  than 'active', the
    // timer will be restarted.  NOTE:  When a conceptual Rtt control row ages
    // out, the         agent needs to remove the associated entries in        
    // the rttMonReactTriggerAdminTable and         rttMonReactTriggerOperTable. 
    // When this value is set to zero, this entry will never be aged out.
    // rttMonScheduleAdminConceptRowAgeout object is superseded by
    // rttMonScheduleAdminConceptRowAgeoutV2. The type is interface{} with range:
    // 0..2073600. Units are seconds.
    Rttmonscheduleadminconceptrowageout interface{}

    // When set to true, this entry will be scheduled to run automatically for the
    // specified duration equal to the life configured, at the same time daily. 
    // This value cannot be set to true  (a) if rttMonScheduleAdminRttLife object
    // has value greater or    equal to 86400 seconds. (b) if sum of values of
    // rttMonScheduleAdminRttLife and    rttMonScheduleAdminConceptRowAgeout is
    // less or equal to    86400 seconds. The type is bool.
    Rttmonscheduleadminrttrecurring interface{}

    // The amount of time this conceptual Rtt control row will exist when not in
    // an 'active' rttMonCtrlOperState.  When this conceptual Rtt control row
    // enters an 'active' state, this timer will be reset and suspended.  When
    // this conceptual RTT control row enters a state other than 'active', the
    // timer will be restarted.  NOTE:  It is the same as
    // rttMonScheduleAdminConceptRowAgeout        except DEFVAL is 0 to be
    // consistent with CLI ageout        default.  When this value is set to zero,
    // this entry will never be aged out. The type is interface{} with range:
    // 0..2073600. Units are seconds.
    Rttmonscheduleadminconceptrowageoutv2 interface{}

    // If true, a reaction is generated when a RTT operation to a
    // rttMonEchoAdminTargetAddress (echo type) causes 
    // rttMonCtrlOperConnectionLostOccurred to change its  value.  Thus
    // connections to intermediate hops will  not cause this value to change.
    // rttMonReactAdminConnectionEnable object is superseded by rttMonReactVar.
    // The type is bool.
    Rttmonreactadminconnectionenable interface{}

    // If true, a reaction is generated when a RTT operation causes
    // rttMonCtrlOperTimeoutOccurred  to change its value.    When the
    // RttMonRttType is 'pathEcho' timeouts to  intermediate hops will not cause 
    // rttMonCtrlOperTimeoutOccurred to change its value.
    // rttMonReactAdminTimeoutEnable object is superseded by rttMonReactVar. The
    // type is bool.
    Rttmonreactadmintimeoutenable interface{}

    // This object specifies the conditions under which
    // rttMonCtrlOperOverThresholdOccurred is changed:  NOTE:  When the
    // RttMonRttType is 'pathEcho' this         objects' value and all associated 
    // object values are only valid when RTT         'echo' operations are to the 
    // rttMonEchoAdminTargetAddress object address.  Thus        'pathEcho'
    // operations to intermediate        hops will not cause this object to
    // change.  never       - rttMonCtrlOperOverThresholdOccurred is              
    // never set immediate   - rttMonCtrlOperOverThresholdOccurred is set         
    // to true when an operation completion time                 exceeds
    // rttMonCtrlAdminThreshold;                 conversely                
    // rttMonCtrlOperOverThresholdOccurred is set                 to false when an
    // operation completion time                 falls below                
    // rttMonReactAdminThresholdFalling  consecutive -
    // rttMonCtrlOperOverThresholdOccurred is set                 to true when an
    // operation completion time                 exceeds rttMonCtrlAdminThreshold
    // on                 rttMonReactAdminThresholdCount consecutive              
    // RTT operations; conversely,                
    // rttMonCtrlOperOverThresholdOccurred is set                 to false when an
    // operation completion time                falls under the                
    // rttMonReactAdminThresholdFalling                 for the same number of
    // consecutive                 operations  xOfy        -
    // rttMonCtrlOperOverThresholdOccurred is set                 to true when x
    // (as specified by                 rttMonReactAdminThresholdCount) out of the
    // last y (as specified by                 rttMonReactAdminThresholdCount2)   
    // operation completion time exceeds                 rttMonCtrlAdminThreshold;
    // conversely, it is set to false when x,                 out of the last y
    // operation completion                time fall below               
    // rttMonReactAdminThresholdFalling                NOTE: When x > y, the probe
    // will never                      generate a reaction. average     -
    // rttMonCtrlOperOverThresholdOccurred is set                 to true when the
    // running average of the                 previous
    // rttMonReactAdminThresholdCount                 operation completion times
    // exceed                 rttMonCtrlAdminThreshold; conversely, it            
    // is set to false when the running average                 falls below the   
    // rttMonReactAdminThresholdFalling  If this value is changed by a management
    // station,  rttMonCtrlOperOverThresholdOccurred is set to false, but  no
    // reaction is generated if the prior value of 
    // rttMonCtrlOperOverThresholdOccurred was true. rttMonReactAdminThresholdType
    // object is superseded by rttMonReactThresholdType. The type is
    // Rttmonreactadminthresholdtype.
    Rttmonreactadminthresholdtype interface{}

    // This object defines a threshold limit. If the RTT operation time falls
    // below this limit and if the conditions specified in
    // rttMonReactAdminThresholdType are satisfied, an  threshold is generated.
    // rttMonReactAdminThresholdFalling object is superseded by
    // rttMonReactThresholdFalling. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    Rttmonreactadminthresholdfalling interface{}

    // This object defines the 'x' value of the xOfy condition specified in
    // rttMonReactAdminThresholdType. rttMonReactAdminThresholdCount object is
    // superseded by rttMonReactThresholdCountX. The type is interface{} with
    // range: 1..16.
    Rttmonreactadminthresholdcount interface{}

    // This object defines the 'y' value of the xOfy condition specified in
    // rttMonReactAdminThresholdType. rttMonReactAdminThresholdCount2 object is
    // superseded by rttMonReactThresholdCountyY. The type is interface{} with
    // range: 1..16.
    Rttmonreactadminthresholdcount2 interface{}

    // Specifies what type(s), if any, of reaction(s) to generate if an operation
    // violates one of the watched  conditions:  none               - no reaction
    // is generated trapOnly           - a trap is generated nmvtOnly           -
    // an SNA NMVT is generated triggerOnly        - all trigger actions defined
    // for this                        entry are initiated trapAndNmvt        -
    // both a trap and an SNA NMVT are                        generated
    // trapAndTrigger     - both a trap and all trigger actions                   
    // are initiated  nmvtAndTrigger     - both a NMVT and all trigger actions    
    // are initiated trapNmvtAndTrigger - a NMVT, trap, and all trigger actions   
    // are initiated  A trigger action is defined via the 
    // rttMonReactTriggerAdminTable. rttMonReactAdminActionType object is
    // superseded by rttMonReactActionType. The type is
    // Rttmonreactadminactiontype.
    Rttmonreactadminactiontype interface{}

    // If true, a reaction is generated when a RTT operation causes
    // rttMonCtrlOperVerifyErrorOccurred  to change its value.
    // rttMonReactAdminVerifyErrorEnable object is superseded by rttMonReactVar.
    // The type is bool.
    Rttmonreactadminverifyerrorenable interface{}

    // The maximum number of groups of paths to record. Specifically this is the
    // number of hourly groups  to keep before rolling over.    The value of one
    // is not advisable because the  group will close and immediately be deleted
    // before the network management station will have the  opportunity to
    // retrieve the statistics.   The value used in the rttMonStatsCaptureTable to
    // uniquely identify this group is the  rttMonStatsCaptureStartTimeIndex. 
    // HTTP and Jitter probes store only two hours of data.  When this object is
    // set to the value of zero all  rttMonStatsCaptureTable data capturing will
    // be shut off. The type is interface{} with range: 0..25.
    Rttmonstatisticsadminnumhourgroups interface{}

    // When RttMonRttType is 'pathEcho' this is the maximum number of statistics
    // paths to record per hourly group.   This value directly represents the path
    // to a target.   For all other RttMonRttTypes this value will be  forced to
    // one by the agent.  NOTE: For 'pathEcho' a source to target path will be    
    // created to to hold all errors that occur when a        specific path or
    // connection has not be found/setup.        Thus, it is advised to set this
    // value greater       than one.  Since this index does not rollover, only the
    // first rttMonStatisticsAdminNumPaths will be kept. The type is interface{}
    // with range: 1..128.
    Rttmonstatisticsadminnumpaths interface{}

    // When RttMonRttType is 'pathEcho' this is the maximum number of statistics
    // hops to record per path group.   This value directly represents the number
    // of hops along  a path to a target, thus we can only support 30 hops.   For
    // all other RttMonRttTypes this value will be  forced to one by the agent. 
    // Since this index does not rollover, only the first
    // rttMonStatisticsAdminNumHops will be kept. This object  is applicable to
    // pathEcho probes only. The type is interface{} with range: 1..30.
    Rttmonstatisticsadminnumhops interface{}

    // The maximum number of statistical distribution Buckets to accumulate. 
    // Since this index does not rollover, only the first
    // rttMonStatisticsAdminNumDistBuckets will be kept.  The last
    // rttMonStatisticsAdminNumDistBucket will contain all entries from its
    // distribution interval start point to infinity. This object is not
    // applicable  to http and jitter probes. The type is interface{} with range:
    // 1..20.
    Rttmonstatisticsadminnumdistbuckets interface{}

    // The statistical distribution buckets interval.  Distribution Bucket
    // Example:  rttMonStatisticsAdminNumDistBuckets = 5 buckets
    // rttMonStatisticsAdminDistInterval = 10 milliseconds  | Bucket 1 | Bucket 2
    // | Bucket 3 | Bucket 4 | Bucket 5  | |  0-9 ms  | 10-19 ms | 20-29 ms |
    // 30-39 ms | 40-Inf ms |  Odd Example:  rttMonStatisticsAdminNumDistBuckets =
    // 1 buckets rttMonStatisticsAdminDistInterval = 10 milliseconds  | Bucket 1 
    // | |  0-Inf ms |  Thus, this odd example shows that the value of 
    // rttMonStatisticsAdminDistInterval does not apply when
    // rttMonStatisticsAdminNumDistBuckets is one. This object is not applicable
    // to http and jitter probes. The type is interface{} with range: 1..100.
    // Units are milliseconds.
    Rttmonstatisticsadmindistinterval interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the rttMonCtrlOperRttLife object.  A
    // new life is created when the same conceptual RTT control row is restarted
    // via the transition of the  rttMonCtrlOperRttLife object and its subsequent 
    // countdown.  The value of zero will shut off all  rttMonHistoryAdminTable
    // data collection. The type is interface{} with range: 0..2.
    Rttmonhistoryadminnumlives interface{}

    // The maximum number of history buckets to record.  When the RttMonRttType is
    // 'pathEcho'  this value directly  represents a path to a target.  For all
    // other  RttMonRttTypes this value should be set to the number  of operations
    // to keep per lifetime.  After rttMonHistoryAdminNumBuckets are filled, the 
    // and the oldest entries are deleted and the most recent
    // rttMonHistoryAdminNumBuckets buckets are retained. The type is interface{}
    // with range: 1..60.
    Rttmonhistoryadminnumbuckets interface{}

    // The maximum number of history samples to record per bucket.  When the
    // RttMonRttType is 'pathEcho' this  value directly represents the number of
    // hops along a  path to a target, thus we can only support 30 hops. For all
    // other RttMonRttTypes this value will be  forced to one by the agent. The
    // type is interface{} with range: 1..30.
    Rttmonhistoryadminnumsamples interface{}

    // Defines a filter for adding RTT results to the history buffer:  none       
    // - no history is recorded all           - the results of all completion
    // times                   and failed completions are recorded overThreshold -
    // the results of completion times                  over
    // rttMonCtrlAdminThreshold are                   recorded. failures      -
    // the results of failed operations (only)                   are recorded. The
    // type is Rttmonhistoryadminfilter.
    Rttmonhistoryadminfilter interface{}

    // This object is updated whenever an object in the conceptual RTT control row
    // is changed or updated. The type is interface{} with range: 0..4294967295.
    Rttmonctrlopermodificationtime interface{}

    // A string which can be used as an aid in tracing problems. The content of
    // this field will depend on the type of  target (rttMonEchoAdminProtocol).  
    // When rttMonEchoAdminProtocol is one of snaLU0EchoAppl, or  snaLU2EchoAppl
    // this object contains the name of the  Logical Unit (LU) being used for this
    // RTT session (from the HOST's point of view), once the session has been 
    // established; this can then be used to correlate this  name to the
    // connection information stored in the  Mainframe Host.  When
    // rttMonEchoAdminProtocol is snaLU62EchoAppl, this  object contains the
    // Logical Unit (LU) name being used for this RTT session, once the session
    // has been established.   This name can be used by the management application
    // to  correlate this objects value to the connection  information stored at
    // this SNMP Agent via the APPC or  APPN mib.  When rttMonEchoAdminProtocol is
    // not one of the  previously mentioned values, this value will be null.  It
    // is primarily intended that this object contains  information which has
    // significance to a human operator. The type is string with length: 0..51.
    Rttmonctrloperdiagtext interface{}

    // This object is set when the rttMonCtrlOperState is set to reset. The type
    // is interface{} with range: 0..4294967295.
    Rttmonctrloperresettime interface{}

    // This object is the number of octets currently in use by this composite
    // conceptual RTT row.  A composite conceptual row include the control,
    // statistics, and  history conceptual rows combined.  (All octets that are
    // addressed via the rttMonCtrlAdminIndex in this mib.). The type is
    // interface{} with range: 0..4294967295.
    Rttmonctrloperoctetsinuse interface{}

    // This object will only change its value when the RttMonRttType is 'echo' or
    // 'pathEcho'.  This object is set to true when the RTT connection fails  to
    // be established or is lost, and set to false when a  connection is
    // reestablished.  When the RttMonRttType is 'pathEcho', connection loss
    // applies only to the rttMonEchoAdminTargetAddress and not to intermediate
    // hops to the Target.  When this value changes and 
    // rttMonReactAdminConnectionEnable is true, a reaction  will occur.   If a
    // trap is sent it is a  rttMonConnectionChangeNotification.  When this value
    // changes and any one of the rttMonReactTable row has rttMonReactVar object
    // value as 'connectionLoss(8)', a reaction may occur.  If a trap is sent it
    // is rttMonNotification with rttMonReactVar value of 'connectionLoss'. The
    // type is bool.
    Rttmonctrloperconnectionlostoccurred interface{}

    // This object will change its value for all RttMonRttTypes.  This object is
    // set to true when an operation times out,  and set to false when an
    // operation completes under  rttMonCtrlAdminTimeout.  When this value
    // changes, a  reaction may occur, as defined by 
    // rttMonReactAdminTimeoutEnable.   When the RttMonRttType is 'pathEcho', this
    // timeout applies only to the rttMonEchoAdminTargetAddress and not to
    // intermediate hops to the Target.  If a trap is sent it is a
    // rttMonTimeoutNotification.  When this value changes and any one of the
    // rttMonReactTable row has rttMonReactVar object value as 'timeout(7)', a
    // reaction may occur.  If a trap is sent it is rttMonNotification with
    // rttMonReactVar value of 'timeout'. The type is bool.
    Rttmonctrlopertimeoutoccurred interface{}

    // This object will change its value for all RttMonRttTypes.  This object is
    // changed by operation completion times over threshold, as defined by
    // rttMonReactAdminThresholdType.   When this value changes, a reaction may
    // occur, as defined  by rttMonReactAdminThresholdType.   If a trap is sent it
    // is a rttMonThresholdNotification.  This object is set to true if the
    // operation completion time exceeds the rttMonCtrlAdminThreshold and set to
    // false when an operation completes under rttMonCtrlAdminThreshold. When this
    // value changes, a reaction may occur, as defined by
    // rttMonReactThresholdType.  If a trap is sent it is rttMonNotification with
    // rttMonReactVar value of 'rtt'. The type is bool.
    Rttmonctrloperoverthresholdoccurred interface{}

    // This is the total number of probe operations that have been attempted.    
    // This value is incremented for each start of an RTT  operation.  Thus when
    // rttMonCtrlAdminRttType is set to  'pathEcho' this value will be incremented
    // by one and  not for very every hop along the path.  This object has the
    // special behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object.  This value is not effected by the rollover of a
    // statistics hourly group. The type is interface{} with range: 0..2147483647.
    Rttmonctrlopernumrtts interface{}

    // This object is decremented every second, until it reaches zero.  When the
    // value of this object is zero RTT operations for this row are suspended. 
    // This  object will either reach zero by a countdown or  it will transition
    // to zero via setting the rttMonCtrlOperState.  When this object reaches zero
    // the agent needs to  transition the rttMonCtrlOperState to 'inactive'. 
    // REMEMBER:  The value 2147483647 has a special             meaning.  When
    // this object has the            value 2147483647, this object will          
    // not decrement.  And thus the life             time will never.  When the
    // rttMonCtrlOperState object is 'active' and  the rttMonReactTriggerOperState
    // object transitions to  'active' this object will not be updated with the 
    // current value of rttMonCrtlAdminRttLife object. The type is interface{}
    // with range: 0..2147483647. Units are seconds.
    Rttmonctrloperrttlife interface{}

    // The RttMonOperStatus object is used to manage the 'state' of the probe that
    // is implementing  conceptual RTT control row.  This status object has six
    // defined values:  reset(1)          - reset this entry, transition          
    // to 'pending' orderlyStop(2)    - shutdown this entry at the end            
    // of the next RTT operation attempt,                       transition to
    // 'inactive' immediateStop(3)  - shutdown this entry immediately             
    // (if possible), transition to                       'inactive' pending(4)   
    // - this value is not settable and                      this conceptual RTT
    // control row is                       waiting for further control either    
    // via the rttMonScheduleAdminTable                       or the
    // rttMonReactAdminTable/                      rttMonReactTriggerAdminTable;  
    // This object can transition to this                      value via two
    // mechanisms, first by                      reseting this object, and second 
    // by creating a conceptual Rtt control                      row with the     
    // rttMonScheduleAdminRttStartTime                      object with the its
    // special value inactive(5)       - this value is not settable and           
    // this conceptual RTT control row is                       waiting for
    // further control via                      the rttMonScheduleAdminTable;     
    // This object can transition to this                      value via two
    // mechanisms, first by                      setting this object to
    // 'orderlyStop'                      or 'immediateStop', second by           
    // the rttMonCtrlOperRttLife object                      reaching zero
    // active(6)         - this value is not settable and                     
    // this conceptual RTT control row is                      currently active
    // restart(7)        - this value is only settable when the                   
    // state is active. It clears the data                      of this entry and
    // remain on active state.  The probes action when this object is set to
    // 'reset':   -  all rows in rttMonStatsCaptureTable that relate to       
    // this conceptual RTT control row are destroyed and        the indices are
    // set to 1   -  if rttMonStatisticsAdminNumHourGroups is not zero, a       
    // single new rttMonStatsCaptureTable row is created   -  all rows in
    // rttMonHistoryCaptureTable that relate        to this RTT definition are
    // destroyed and the indices       are set to 1   -  implied history used for
    // timeout or threshold       notification (see rttMonReactAdminThresholdType
    // or       rttMonReactThresholdType)       is purged   - 
    // rttMonCtrlOperRttLife is set to        rttMonScheduleAdminRttLife   - 
    // rttMonCtrlOperNumRtts is set to zero   -  rttMonCtrlOperTimeoutOccurred,   
    // rttMonCtrlOperOverThresholdOccurred, and       
    // rttMonCtrlOperConnectionLostOccurred are set to        false; if this
    // causes a change in the value of        either of these objects, resolution
    // notifications        will not occur   -  the next RTT operation is
    // controlled by the objects       in the rttMonScheduleAdminTable or the     
    // rttMonReactAdminTable/rttMonReactTriggerAdminTable   -  if the
    // rttMonReactTriggerOperState is 'active', it        will transition to
    // 'pending'   -  all rttMonReactTriggerAdminEntries pointing to       this
    // conceptual entry with their        rttMonReactTriggerOperState object
    // 'active',        will transition their OperState to 'pending'   -  all open
    // connections must be maintained  This can be used to synchronize various RTT
    // definitions, so that the RTT requests occur  simultaneously, or as
    // simultaneously as possible.  The probes action when this object transitions
    // to    'inactive' (via setting this object to 'orderlyStop'    or
    // 'immediateStop' or by rttMonCtrlOperRttLife    reaching zero):   -  all
    // statistics and history collection information       table entries will be
    // closed and kept   -  implied history used for timeout or threshold      
    // notification (see rttMonReactAdminThresholdType or      
    // rttMonReactThresholdType)       is purged   - 
    // rttMonCtrlOperTimeoutOccurred,        rttMonCtrlOperOverThresholdOccurred,
    // and        rttMonCtrlOperConnectionLostOccurred are set to        false; if
    // this causes a change in the value of        either of these objects,
    // resolution notifications        will not occur.   -  the next RTT request
    // is controlled by the objects       in the rttMonScheduleAdminTable   -  if
    // the rttMonReactTriggerOperState is 'active', it        will transition to
    // 'pending' (this denotes that       the Trigger will be ready the next time
    // this       object goes active)   -  all rttMonReactTriggerAdminEntries
    // pointing to       this conceptual entry with their       
    // rttMonReactTriggerOperState object 'active',        will transition their
    // OperState to 'pending'   -  all open connections are to be closed and
    // cleanup.               rttMonCtrlOperState                     STATE       
    // +-------------------------------------------+           |      A       |   
    // B      |      C      | ACTION       |  'pending'   |  'inactive'  |  
    // 'active'  | +----------------+--------------+--------------+-------------+
    // | OperState set  |    noError   |inconsistent- |   noError   | |  to
    // 'reset'    |              | Value        |             | |                |
    // -> A      |              |   -> A      |
    // +----------------+--------------+--------------+-------------+ | OperState
    // set  |    noError   |    noError   |   noError   | |to 'orderlyStop'|    ->
    // B      |    -> B      |   -> B      | |     or to      |              |    
    // |             | |'immediateStop' |              |              |           
    // | +----------------+--------------+--------------+-------------+ |  Event
    // causes  |    -> C      |    -> B      |   -> C      | | Trigger State  |   
    // |              |   see (3)   | | to transition  |              |           
    // |             | | to 'active'    |              |              |           
    // | +----------------+--------------+--------------+-------------+ |
    // AdminStatus    |    -> C      |    -> C      |   see (1)   | | transitions
    // to |              |              |             | | 'active' &     |        
    // |              |             | | RttStartTime is|              |           
    // |             | | special value  |              |              |           
    // | | of one.        |              |              |             |
    // +----------------+--------------+--------------+-------------+ |
    // AdminStatus    |    -> A      |    -> A      |   see (1)   | | transitions
    // to |              |              |             | | 'active' &     |        
    // |              |             | | RttStartTime is|              |           
    // |             | | special value  |              |              |           
    // | | of less than   |              |              |             | | current
    // time,  |              |              |             | | excluding one. |    
    // |              |             |
    // +----------------+--------------+--------------+-------------+ |
    // AdminStatus    |    -> A      |    -> B      |   see (2)   | | transitions
    // to |              |              |             | | 'notInService' |        
    // |              |             |
    // +----------------+--------------+--------------+-------------+ |
    // AdminStatus    |    -> B      |    -> B      |   -> B      | | transitions
    // to |              |              |             | | 'delete'       |        
    // |              |             |
    // +----------------+--------------+--------------+-------------+ |
    // AdminStatus is |    -> C      |    -> C      |   -> C      | | 'active' &
    // the |              |              |   see (3)   | | RttStartTime   |       
    // |              |             | | arrives        |              |           
    // |             |
    // +----------------+--------------+--------------+-------------+ |  
    // RowAgeout    |    -> B      |    -> B      |   -> B      | |    expires    
    // |              |              |             |
    // +----------------+--------------+--------------+-------------+ | 
    // OperRttLife   |    N/A       |    N/A       |   -> B      | | counts down
    // to |              |              |             | | zero           |        
    // |              |             |
    // +----------------+--------------+--------------+-------------+  (1) -
    // rttMonCtrlOperState must have transitioned to 'inactive' or 'pending'
    // before the rttMonCtrlAdminStatus can transition to 'active'.  See (2). (2)
    // - rttMonCtrlAdminStatus cannot transition to 'notInService' unless
    // rttMonCtrlOperState has been previously forced to 'inactive' or 'pending'.
    // (3) - when this happens the rttMonCtrlOperRttLife will not be updated with
    // the rttMonCtrlAdminRttLife.  NOTE:  In order for all objects in a PDU to be
    // set        at the same time, this object can not be        part of a
    // multi-bound PDU. The type is Rttmonctrloperstate.
    Rttmonctrloperstate interface{}

    // This object is true if rttMonCtrlAdminVerifyData is set to true and data
    // corruption occurs. The type is bool.
    Rttmonctrloperverifyerroroccurred interface{}

    // The completion time of the latest RTT operation successfully completed. 
    // The unit of this object will be microsecond when rttMonCtrlAdminRttType is
    // set to 'jitter' and  rttMonEchoAdminPrecision is set to 'microsecond'.
    // Otherwise, the unit of this object will be millisecond. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds/microseconds.
    Rttmonlatestrttopercompletiontime interface{}

    // A sense code for the completion status of the latest RTT operation. The
    // type is RttResponseSense.
    Rttmonlatestrttopersense interface{}

    // An application specific sense code for the completion status of the latest
    // RTT operation.  This  object will only be valid when the 
    // rttMonLatestRttOperSense object is set to  'applicationSpecific'. 
    // Otherwise, this object's  value is not valid. The type is interface{} with
    // range: 0..2147483647.
    Rttmonlatestrttoperapplspecificsense interface{}

    // A sense description for the completion status of the latest RTT operation
    // when the  rttMonLatestRttOperSense object is set to  'applicationSpecific'.
    // The type is string.
    Rttmonlatestrttopersensedescription interface{}

    // The value of the agent system time at the time of the latest RTT operation.
    // The type is interface{} with range: 0..4294967295.
    Rttmonlatestrttopertime interface{}

    // When the RttMonRttType is 'echo', 'pathEcho', 'udpEcho', 'tcpConnect',
    // 'dns' and 'dlsw' this is a string which specifies  the address of the
    // target for this RTT operation.  When the  RttMonRttType is not one of these
    // types this object will  be null.  This address will be the address of the
    // hop along the path to the rttMonEchoAdminTargetAddress address, including
    // rttMonEchoAdminTargetAddress address, or just the
    // rttMonEchoAdminTargetAddress address, when the path information is not
    // collected.  This behavior is defined by the rttMonCtrlAdminRttType object. 
    // The interpretation of this string depends on the type of RTT operation
    // selected, as specified by the rttMonEchoAdminProtocol object.  See
    // rttMonEchoAdminTargetAddress for a complete description. The type is
    // string.
    Rttmonlatestrttoperaddress interface{}
}

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetFilter() yfilter.YFilter { return rttmonctrladminentry.YFilter }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) SetFilter(yf yfilter.YFilter) { rttmonctrladminentry.YFilter = yf }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonCtrlAdminOwner" { return "Rttmonctrladminowner" }
    if yname == "rttMonCtrlAdminTag" { return "Rttmonctrladmintag" }
    if yname == "rttMonCtrlAdminRttType" { return "Rttmonctrladminrtttype" }
    if yname == "rttMonCtrlAdminThreshold" { return "Rttmonctrladminthreshold" }
    if yname == "rttMonCtrlAdminFrequency" { return "Rttmonctrladminfrequency" }
    if yname == "rttMonCtrlAdminTimeout" { return "Rttmonctrladmintimeout" }
    if yname == "rttMonCtrlAdminVerifyData" { return "Rttmonctrladminverifydata" }
    if yname == "rttMonCtrlAdminStatus" { return "Rttmonctrladminstatus" }
    if yname == "rttMonCtrlAdminNvgen" { return "Rttmonctrladminnvgen" }
    if yname == "rttMonCtrlAdminGroupName" { return "Rttmonctrladmingroupname" }
    if yname == "rttMonScheduleAdminRttLife" { return "Rttmonscheduleadminrttlife" }
    if yname == "rttMonScheduleAdminRttStartTime" { return "Rttmonscheduleadminrttstarttime" }
    if yname == "rttMonScheduleAdminConceptRowAgeout" { return "Rttmonscheduleadminconceptrowageout" }
    if yname == "rttMonScheduleAdminRttRecurring" { return "Rttmonscheduleadminrttrecurring" }
    if yname == "rttMonScheduleAdminConceptRowAgeoutV2" { return "Rttmonscheduleadminconceptrowageoutv2" }
    if yname == "rttMonReactAdminConnectionEnable" { return "Rttmonreactadminconnectionenable" }
    if yname == "rttMonReactAdminTimeoutEnable" { return "Rttmonreactadmintimeoutenable" }
    if yname == "rttMonReactAdminThresholdType" { return "Rttmonreactadminthresholdtype" }
    if yname == "rttMonReactAdminThresholdFalling" { return "Rttmonreactadminthresholdfalling" }
    if yname == "rttMonReactAdminThresholdCount" { return "Rttmonreactadminthresholdcount" }
    if yname == "rttMonReactAdminThresholdCount2" { return "Rttmonreactadminthresholdcount2" }
    if yname == "rttMonReactAdminActionType" { return "Rttmonreactadminactiontype" }
    if yname == "rttMonReactAdminVerifyErrorEnable" { return "Rttmonreactadminverifyerrorenable" }
    if yname == "rttMonStatisticsAdminNumHourGroups" { return "Rttmonstatisticsadminnumhourgroups" }
    if yname == "rttMonStatisticsAdminNumPaths" { return "Rttmonstatisticsadminnumpaths" }
    if yname == "rttMonStatisticsAdminNumHops" { return "Rttmonstatisticsadminnumhops" }
    if yname == "rttMonStatisticsAdminNumDistBuckets" { return "Rttmonstatisticsadminnumdistbuckets" }
    if yname == "rttMonStatisticsAdminDistInterval" { return "Rttmonstatisticsadmindistinterval" }
    if yname == "rttMonHistoryAdminNumLives" { return "Rttmonhistoryadminnumlives" }
    if yname == "rttMonHistoryAdminNumBuckets" { return "Rttmonhistoryadminnumbuckets" }
    if yname == "rttMonHistoryAdminNumSamples" { return "Rttmonhistoryadminnumsamples" }
    if yname == "rttMonHistoryAdminFilter" { return "Rttmonhistoryadminfilter" }
    if yname == "rttMonCtrlOperModificationTime" { return "Rttmonctrlopermodificationtime" }
    if yname == "rttMonCtrlOperDiagText" { return "Rttmonctrloperdiagtext" }
    if yname == "rttMonCtrlOperResetTime" { return "Rttmonctrloperresettime" }
    if yname == "rttMonCtrlOperOctetsInUse" { return "Rttmonctrloperoctetsinuse" }
    if yname == "rttMonCtrlOperConnectionLostOccurred" { return "Rttmonctrloperconnectionlostoccurred" }
    if yname == "rttMonCtrlOperTimeoutOccurred" { return "Rttmonctrlopertimeoutoccurred" }
    if yname == "rttMonCtrlOperOverThresholdOccurred" { return "Rttmonctrloperoverthresholdoccurred" }
    if yname == "rttMonCtrlOperNumRtts" { return "Rttmonctrlopernumrtts" }
    if yname == "rttMonCtrlOperRttLife" { return "Rttmonctrloperrttlife" }
    if yname == "rttMonCtrlOperState" { return "Rttmonctrloperstate" }
    if yname == "rttMonCtrlOperVerifyErrorOccurred" { return "Rttmonctrloperverifyerroroccurred" }
    if yname == "rttMonLatestRttOperCompletionTime" { return "Rttmonlatestrttopercompletiontime" }
    if yname == "rttMonLatestRttOperSense" { return "Rttmonlatestrttopersense" }
    if yname == "rttMonLatestRttOperApplSpecificSense" { return "Rttmonlatestrttoperapplspecificsense" }
    if yname == "rttMonLatestRttOperSenseDescription" { return "Rttmonlatestrttopersensedescription" }
    if yname == "rttMonLatestRttOperTime" { return "Rttmonlatestrttopertime" }
    if yname == "rttMonLatestRttOperAddress" { return "Rttmonlatestrttoperaddress" }
    return ""
}

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetSegmentPath() string {
    return "rttMonCtrlAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonctrladminentry.Rttmonctrladminindex) + "']"
}

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonctrladminentry.Rttmonctrladminindex
    leafs["rttMonCtrlAdminOwner"] = rttmonctrladminentry.Rttmonctrladminowner
    leafs["rttMonCtrlAdminTag"] = rttmonctrladminentry.Rttmonctrladmintag
    leafs["rttMonCtrlAdminRttType"] = rttmonctrladminentry.Rttmonctrladminrtttype
    leafs["rttMonCtrlAdminThreshold"] = rttmonctrladminentry.Rttmonctrladminthreshold
    leafs["rttMonCtrlAdminFrequency"] = rttmonctrladminentry.Rttmonctrladminfrequency
    leafs["rttMonCtrlAdminTimeout"] = rttmonctrladminentry.Rttmonctrladmintimeout
    leafs["rttMonCtrlAdminVerifyData"] = rttmonctrladminentry.Rttmonctrladminverifydata
    leafs["rttMonCtrlAdminStatus"] = rttmonctrladminentry.Rttmonctrladminstatus
    leafs["rttMonCtrlAdminNvgen"] = rttmonctrladminentry.Rttmonctrladminnvgen
    leafs["rttMonCtrlAdminGroupName"] = rttmonctrladminentry.Rttmonctrladmingroupname
    leafs["rttMonScheduleAdminRttLife"] = rttmonctrladminentry.Rttmonscheduleadminrttlife
    leafs["rttMonScheduleAdminRttStartTime"] = rttmonctrladminentry.Rttmonscheduleadminrttstarttime
    leafs["rttMonScheduleAdminConceptRowAgeout"] = rttmonctrladminentry.Rttmonscheduleadminconceptrowageout
    leafs["rttMonScheduleAdminRttRecurring"] = rttmonctrladminentry.Rttmonscheduleadminrttrecurring
    leafs["rttMonScheduleAdminConceptRowAgeoutV2"] = rttmonctrladminentry.Rttmonscheduleadminconceptrowageoutv2
    leafs["rttMonReactAdminConnectionEnable"] = rttmonctrladminentry.Rttmonreactadminconnectionenable
    leafs["rttMonReactAdminTimeoutEnable"] = rttmonctrladminentry.Rttmonreactadmintimeoutenable
    leafs["rttMonReactAdminThresholdType"] = rttmonctrladminentry.Rttmonreactadminthresholdtype
    leafs["rttMonReactAdminThresholdFalling"] = rttmonctrladminentry.Rttmonreactadminthresholdfalling
    leafs["rttMonReactAdminThresholdCount"] = rttmonctrladminentry.Rttmonreactadminthresholdcount
    leafs["rttMonReactAdminThresholdCount2"] = rttmonctrladminentry.Rttmonreactadminthresholdcount2
    leafs["rttMonReactAdminActionType"] = rttmonctrladminentry.Rttmonreactadminactiontype
    leafs["rttMonReactAdminVerifyErrorEnable"] = rttmonctrladminentry.Rttmonreactadminverifyerrorenable
    leafs["rttMonStatisticsAdminNumHourGroups"] = rttmonctrladminentry.Rttmonstatisticsadminnumhourgroups
    leafs["rttMonStatisticsAdminNumPaths"] = rttmonctrladminentry.Rttmonstatisticsadminnumpaths
    leafs["rttMonStatisticsAdminNumHops"] = rttmonctrladminentry.Rttmonstatisticsadminnumhops
    leafs["rttMonStatisticsAdminNumDistBuckets"] = rttmonctrladminentry.Rttmonstatisticsadminnumdistbuckets
    leafs["rttMonStatisticsAdminDistInterval"] = rttmonctrladminentry.Rttmonstatisticsadmindistinterval
    leafs["rttMonHistoryAdminNumLives"] = rttmonctrladminentry.Rttmonhistoryadminnumlives
    leafs["rttMonHistoryAdminNumBuckets"] = rttmonctrladminentry.Rttmonhistoryadminnumbuckets
    leafs["rttMonHistoryAdminNumSamples"] = rttmonctrladminentry.Rttmonhistoryadminnumsamples
    leafs["rttMonHistoryAdminFilter"] = rttmonctrladminentry.Rttmonhistoryadminfilter
    leafs["rttMonCtrlOperModificationTime"] = rttmonctrladminentry.Rttmonctrlopermodificationtime
    leafs["rttMonCtrlOperDiagText"] = rttmonctrladminentry.Rttmonctrloperdiagtext
    leafs["rttMonCtrlOperResetTime"] = rttmonctrladminentry.Rttmonctrloperresettime
    leafs["rttMonCtrlOperOctetsInUse"] = rttmonctrladminentry.Rttmonctrloperoctetsinuse
    leafs["rttMonCtrlOperConnectionLostOccurred"] = rttmonctrladminentry.Rttmonctrloperconnectionlostoccurred
    leafs["rttMonCtrlOperTimeoutOccurred"] = rttmonctrladminentry.Rttmonctrlopertimeoutoccurred
    leafs["rttMonCtrlOperOverThresholdOccurred"] = rttmonctrladminentry.Rttmonctrloperoverthresholdoccurred
    leafs["rttMonCtrlOperNumRtts"] = rttmonctrladminentry.Rttmonctrlopernumrtts
    leafs["rttMonCtrlOperRttLife"] = rttmonctrladminentry.Rttmonctrloperrttlife
    leafs["rttMonCtrlOperState"] = rttmonctrladminentry.Rttmonctrloperstate
    leafs["rttMonCtrlOperVerifyErrorOccurred"] = rttmonctrladminentry.Rttmonctrloperverifyerroroccurred
    leafs["rttMonLatestRttOperCompletionTime"] = rttmonctrladminentry.Rttmonlatestrttopercompletiontime
    leafs["rttMonLatestRttOperSense"] = rttmonctrladminentry.Rttmonlatestrttopersense
    leafs["rttMonLatestRttOperApplSpecificSense"] = rttmonctrladminentry.Rttmonlatestrttoperapplspecificsense
    leafs["rttMonLatestRttOperSenseDescription"] = rttmonctrladminentry.Rttmonlatestrttopersensedescription
    leafs["rttMonLatestRttOperTime"] = rttmonctrladminentry.Rttmonlatestrttopertime
    leafs["rttMonLatestRttOperAddress"] = rttmonctrladminentry.Rttmonlatestrttoperaddress
    return leafs
}

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetYangName() string { return "rttMonCtrlAdminEntry" }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) SetParent(parent types.Entity) { rttmonctrladminentry.parent = parent }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetParent() types.Entity { return rttmonctrladminentry.parent }

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetParentYangName() string { return "rttMonCtrlAdminTable" }

// CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate represents        part of a multi-bound PDU.
type CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate string

const (
    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate_reset CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate = "reset"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate_orderlyStop CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate = "orderlyStop"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate_immediateStop CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate = "immediateStop"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate_pending CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate = "pending"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate_inactive CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate = "inactive"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate_active CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate = "active"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate_restart CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrloperstate = "restart"
)

// CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter represents                  are recorded.
type CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter string

const (
    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter_none CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter = "none"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter_all CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter = "all"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter_overThreshold CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter = "overThreshold"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter_failures CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonhistoryadminfilter = "failures"
)

// CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype represents rttMonReactActionType.
type CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype string

const (
    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_none CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "none"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_trapOnly CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "trapOnly"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_nmvtOnly CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "nmvtOnly"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_triggerOnly CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "triggerOnly"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_trapAndNmvt CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "trapAndNmvt"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_trapAndTrigger CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "trapAndTrigger"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_nmvtAndTrigger CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "nmvtAndTrigger"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype_trapNmvtAndTrigger CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminactiontype = "trapNmvtAndTrigger"
)

// CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype represents rttMonReactThresholdType.
type CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype string

const (
    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype_never CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype = "never"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype_immediate CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype = "immediate"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype_consecutive CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype = "consecutive"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype_xOfy CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype = "xOfy"

    CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype_average CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonreactadminthresholdtype = "average"
)

// CISCORTTMONMIB_Rttmonechoadmintable
// A table that contains Round Trip Time (RTT) specific
// definitions.
// 
// This table is controlled via the 
// rttMonCtrlAdminTable.  Entries in this table are
// created via the rttMonCtrlAdminStatus object.
type CISCORTTMONMIB_Rttmonechoadmintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for RttMonRttType
    // conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry.
    Rttmonechoadminentry []CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry
}

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetFilter() yfilter.YFilter { return rttmonechoadmintable.YFilter }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) SetFilter(yf yfilter.YFilter) { rttmonechoadmintable.YFilter = yf }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetGoName(yname string) string {
    if yname == "rttMonEchoAdminEntry" { return "Rttmonechoadminentry" }
    return ""
}

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetSegmentPath() string {
    return "rttMonEchoAdminTable"
}

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonEchoAdminEntry" {
        for _, c := range rttmonechoadmintable.Rttmonechoadminentry {
            if rttmonechoadmintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry{}
        rttmonechoadmintable.Rttmonechoadminentry = append(rttmonechoadmintable.Rttmonechoadminentry, child)
        return &rttmonechoadmintable.Rttmonechoadminentry[len(rttmonechoadmintable.Rttmonechoadminentry)-1]
    }
    return nil
}

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonechoadmintable.Rttmonechoadminentry {
        children[rttmonechoadmintable.Rttmonechoadminentry[i].GetSegmentPath()] = &rttmonechoadmintable.Rttmonechoadminentry[i]
    }
    return children
}

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetYangName() string { return "rttMonEchoAdminTable" }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) SetParent(parent types.Entity) { rttmonechoadmintable.parent = parent }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetParent() types.Entity { return rttmonechoadmintable.parent }

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry
// A list of objects that define specific configuration for
// RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // Specifies the protocol to be used to perform the RTT operation. The
    // following list defines what protocol  should be used for each probe type: 
    // echo, pathEcho   - ipIcmpEcho / mplsLspPingAppl udpEcho          -
    // ipUdpEchoAppl tcpConnect       - ipTcpConn http             - httpAppl
    // jitter           - jitterAppl dlsw             - dlswAppl dhcp            
    // - dhcpAppl ftp              - ftpAppl mplsLspPing      - mplsLspPingAppl
    // voip             - voipAppl video            - videoAppl  When this
    // protocol does not support the type, a 'badValue' error will be returned.
    // The type is RttMonProtocol.
    Rttmonechoadminprotocol interface{}

    // A string which specifies the address of the target. The type is string.
    Rttmonechoadmintargetaddress interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the request  message, when using SNA protocols.  For non-ARR
    // protocols' RTT request/responses,  this value represents the native payload
    // size.  REMEMBER:  The ARR Header overhead is not included             in
    // this value.  For echo probes the total packet size = (IP header(20) +  ICMP
    // header(8) + 8 (internal timestamps) + request size).  For echo and pathEcho
    // default request size is 28. For udp probe, default request size is 16 and
    // for jitter  probe it is 32. For dlsw probes default request size is 0.  The
    // minimum request size for echo and pathEcho is 28 bytes, for udp it is 4 and
    // for jitter it is 16. For udp and jitter probes the maximum request size is
    // 1500.  For ethernetPing the default request size is 66. For ethernetJitter
    // the default request size is 51. The type is interface{} with range:
    // 0..16384. Units are octets.
    Rttmonechoadminpktdatarequestsize interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the response message. This value is passed to the RTT Echo
    // Server via a field in the ARR Header.  For non-ARR RTT request/response
    // (i.e. ipIcmpecho) this value will be set by the agent to match the size of
    // rttMonEchoAdminPktDataRequestSize, when native payloads are supported. 
    // REMEMBER:  The ARR Header overhead is not included             in this
    // value.  This object is only supported by SNA protocols. The type is
    // interface{} with range: 0..16384.
    Rttmonechoadminpktdataresponsesize interface{}

    // This object represents the target's port number. This object is applicable
    // to udpEcho, tcpConnect and jitter probes. The type is interface{} with
    // range: 0..65536.
    Rttmonechoadmintargetport interface{}

    // A string which specifies the IP address of the source. This object is
    // applicable to all probes except dns, dlsw  and sna. The type is string.
    Rttmonechoadminsourceaddress interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a  port allocated by the system. This
    // object is applicable  to all probes except dns, dlsw and sna. The type is
    // interface{} with range: 0..65536.
    Rttmonechoadminsourceport interface{}

    // If this object is enabled, then the RTR application will send control
    // messages to a responder, residing on the  target router to respond to the
    // data request packets being  sent by the source router. This object is not
    // applicable to  echo, pathEcho, dns and http probes. The type is bool.
    Rttmonechoadmincontrolenable interface{}

    // This object represents the type of service octet in an IP header. This
    // object is not applicable to dhcp, dns,  ethernetPing and ethernetJitter.
    // The type is interface{} with range: 0..255.
    Rttmonechoadmintos interface{}

    // If this object is enabled then it means that the application calculates
    // response time for a specific path, defined in rttMonEchoPathAdminEntry.
    // This object is applicable to echo  probe only. The type is bool.
    Rttmonechoadminlsrenable interface{}

    // A string which specifies the address of the target. This string can be in
    // IP address format or a hostname. This object is applicable to dns probe
    // only. The type is string.
    Rttmonechoadmintargetaddressstring interface{}

    // A string which specifies the ip address of the name-server. This object is
    // applicable to dns probe only. The type is string.
    Rttmonechoadminnameserver interface{}

    // A code that represents the specific type of RTT operation. This object is
    // applicable to http and ftp probe only. The type is RttMonOperation.
    Rttmonechoadminoperation interface{}

    // A string which specifies the version number of the HTTP Server.  The syntax
    // for the version string is  <major number>.<minor number> An example would
    // be 1.0,  1.1 etc.,.  This object is applicable to http probe only. The type
    // is string with length: 3..10.
    Rttmonechoadminhttpversion interface{}

    // A string which represents the URL to which a HTTP probe should communicate
    // with. This object is applicable to http probe only. The type is string.
    Rttmonechoadminurl interface{}

    // If this object is false then it means that HTTP request should not download
    // cached pages. This means that the request should  be forwarded to the
    // origin server. This object is applicable to http probe only. The type is
    // bool.
    Rttmonechoadmincache interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. This value is currently used for  Jitter probe. This object
    // is applicable to jitter probe only. The type is interface{} with range:
    // 0..60000. Units are milliseconds.
    Rttmonechoadmininterval interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is currently used for Jitter probe.  This object is applicable
    // to jitter probe only. The type is interface{} with range:
    // -2147483648..2147483647.
    Rttmonechoadminnumpackets interface{}

    // This string represents the proxy server information. This object is
    // applicable to http probe only. The type is string.
    Rttmonechoadminproxy interface{}

    // This string stores the content of HTTP raw request. If the request cannot
    // fit into String1 then it should  be split and put in Strings 1 through 5. 
    // This string stores the content of the DHCP raw option data.  The raw DHCP
    // option data must be in HEX. If an odd number of characters are specified, a
    // 0 will be appended to the end of the string.  Only DHCP option 82 (decimal)
    // is allowed. Here is an example of a valid string: 5208010610005A6F1234 Only
    // rttMonEchoAdminString1 is used for dhcp, Strings 1 through 5 are not used. 
    // This object is applicable to http and dhcp probe  types only. The type is
    // string.
    Rttmonechoadminstring1 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    Rttmonechoadminstring2 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    Rttmonechoadminstring3 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    Rttmonechoadminstring4 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    Rttmonechoadminstring5 interface{}

    // A code that represents the specific type of RTT operation. This object is
    // applicable to ftp probe only. The type is RttMonOperation.
    Rttmonechoadminmode interface{}

    // This field is used to specify the VPN name in which the RTT operation will
    // be used. For regular RTT operation this field should not be configured. The
    // agent  will use this field to identify the VPN routing Table for this
    // operation. The type is string with length: 0..32.
    Rttmonechoadminvrfname interface{}

    // Specifies the codec type to be used with jitter probe. This is applicable
    // only for the jitter probe.  If codec-type is configured the following
    // parameters cannot be  configured. rttMonEchoAdminPktDataRequestSize
    // rttMonEchoAdminInterval rttMonEchoAdminNumPackets. The type is
    // RttMonCodecType.
    Rttmonechoadmincodectype interface{}

    // This field represents the inter-packet delay between packets and is in
    // milliseconds. This object is applicable only to jitter probe which uses
    // codec type. The type is interface{} with range: 0..60000. Units are
    // milliseconds.
    Rttmonechoadmincodecinterval interface{}

    // This object represents the number of octets that needs to be placed into
    // the Data portion of the message. This value is used only for jitter probe
    // which uses codec type. The type is interface{} with range: 0..16384. Units
    // are octets.
    Rttmonechoadmincodecpayload interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is used only for jitter probe which uses codec type. The type is
    // interface{} with range: 0..60000.
    Rttmonechoadmincodecnumpackets interface{}

    // The advantage factor is dependant on the type of access and how the service
    // is to be used. Conventional Wire-line     0 Mobility within Building    5
    // Mobility within geographic area  10 Access to hard-to-reach location   20 
    // This will be used while calculating the ICPIF values This valid only for
    // Jitter while calculating the ICPIF value. The type is interface{} with
    // range: 0..20.
    Rttmonechoadminicpifadvfactor interface{}

    // The type of the target FEC for the RTT 'echo' and 'pathEcho' operations
    // based on 'mplsLspPingAppl' RttMonProtocol.  ldpIpv4Prefix   - LDP IPv4
    // prefix. The type is Rttmonechoadminlspfectype.
    Rttmonechoadminlspfectype interface{}

    // A string which specifies a valid 127/8 address. This address is of the form
    // 127.x.y.z. This address is not used to route the MPLS echo packet to the
    // destination but is used for load balancing in cases where the IP payload's
    // destination address is used for load balancing. The type is string.
    Rttmonechoadminlspselector interface{}

    // This object specifies the reply mode for the LSP Echo requests. The type is
    // RttMonLSPPingReplyMode.
    Rttmonechoadminlspreplymode interface{}

    // This object represents the TTL setting for MPLS echo request packets. For
    // ping operation this represents the TTL value to be set in the echo request
    // packet. For trace operation it represent the maximum ttl value that can be
    // set in the echo request packets starting with TTL=1.  For 'echo' based on
    // mplsLspPingAppl the default TTL will be set to 255, and for 'pathEcho'
    // based on mplsLspPingAppl the default will be set to 30.  Note: This object
    // cannot be set to the value of 0. The default value of 0 signifies the
    // default TTL values to be used for 'echo' and 'pathEcho' based on
    // 'mplsLspPingAppl'. The type is interface{} with range: 0..255.
    Rttmonechoadminlspttl interface{}

    // This object represents the EXP value that needs to be put as precedence bit
    // in the MPLS echo request IP header. The type is interface{} with range:
    // 0..7.
    Rttmonechoadminlspexp interface{}

    // This object specifies the accuracy of statistics that needs to be
    // calculated milliseconds - The accuracy of stats will be of milliseconds
    // microseconds - The accuracy of stats will be in microseconds. This value
    // can be set only for jitter operation. The type is Rttmonechoadminprecision.
    Rttmonechoadminprecision interface{}

    // This object specifies the priority that will be assigned to probe packet. 
    // This value can be set only for jitter  operation. The type is
    // Rttmonechoadminprobepakpriority.
    Rttmonechoadminprobepakpriority interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered acceptable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of NTP offsets on source and responder. The value specified is
    // microseconds. This value can be set only for jitter operation  with
    // precision of microsecond. The type is interface{} with range:
    // -2147483648..2147483647. Units are microseconds.
    Rttmonechoadminowntpsynctolabs interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered acceptable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of  NTP offsets on source and responder. The value is
    // expressed  as the percentage of actual oneway latency that is measured. 
    // This value can be set only for jitter operation with precision  of
    // microsecond. The type is interface{} with range: 0..100.
    Rttmonechoadminowntpsynctolpct interface{}

    // This object specifies whether the value in specified for oneway NTP sync
    // tolerance is absolute value or percent value. The type is
    // Rttmonechoadminowntpsynctoltype.
    Rttmonechoadminowntpsynctoltype interface{}

    // This string stores the called number of post dial delay. This object is
    // applicable to voip post dial delay probe only. The number will be like the
    // one actualy the user could dial. It has the number required by the local
    // country dial plan, plus E.164 number. The maximum length is 24 digits. Only
    // digit (0-9) is allowed. The type is string with length: 0..24.
    Rttmonechoadmincallednumber interface{}

    // A code that represents the detect point of post dial delay. This object is
    // applicable to SAA post dial delay probe only. The type is RttMonOperation.
    Rttmonechoadmindetectpoint interface{}

    // A boolean that represents VoIP GK registration delay. This object is
    // applicable to SAA GK registration delay  probe only. The type is bool.
    Rttmonechoadmingkregistration interface{}

    // A string which specifies the voice-port on the source gateway. This object
    // is applicable to RTP probe only. The type is string.
    Rttmonechoadminsourcevoiceport interface{}

    // Duration of RTP/Video Probe session. This object is applicable to RTP and
    // Video probe. The type is interface{} with range: 1..600.
    Rttmonechoadmincallduration interface{}

    // This object specifies the DSCP value to be set in the IP header of the LSP
    // echo reply packet. The value of this object will be in range of DiffServ
    // codepoint values between 0 to 63.  Note: This object cannot be set to value
    // of 255. This default value specifies that DSCP is not set for this row. The
    // type is interface{} with range: 0..63 | 255..None.
    Rttmonechoadminlspreplydscp interface{}

    // This object specifies if the explicit-null label is to be added to LSP echo
    // requests which are sent while performing RTT operation. The type is bool.
    Rttmonechoadminlspnullshim interface{}

    // This object specifies the destination maintenance point ID. It is only
    // applicable to ethernetPing and ethernetJitter  operation. It will be set to
    // 0 for other types of  operations. The type is interface{} with range:
    // 0..8191.
    Rttmonechoadmintargetmpid interface{}

    // This object specifies the name of the domain in which the destination
    // maintenance point lies. It is only applicable to  ethernetPing and
    // ethernetJitter operation. The type is string.
    Rttmonechoadmintargetdomainname interface{}

    // This object specifies the ID of the VLAN in which the destination
    // maintenance point lies. It is only applicable to  ethernetPing and
    // ethernetJitter operation.  It will be set to 0 for other types of
    // operations. The type is interface{} with range: 1..4094.
    Rttmonechoadmintargetvlan interface{}

    // This object specifies the class of service in an Ethernet packet header. It
    // is only applicable to ethernetPing and  ethernetJitter operation. The type
    // is interface{} with range: 0..7.
    Rttmonechoadminethernetcos interface{}

    // This object specifies MPLS LSP pseudowire VCCV ID values between 1 to
    // 2147483647.  Note: This object cannot be set to value of 0. This default
    // value specifies that VCCV is not set for this row. The type is interface{}
    // with range: 0..2147483647.
    Rttmonechoadminlspvccvid interface{}

    // This object specifies the Ethernet Virtual Connection in which the
    // destination maintenance point lies. It is only  applicable to ethernetPing
    // and ethernetJitter operation.  It will be set to NULL for other types of
    // operations. The type is string with length: 0..100.
    Rttmonechoadmintargetevc interface{}

    // This object specifies that Port Level CFM testing towards an Outward/Down
    // MEP will be used. It is only applicable to  ethernetPing and ethernetJitter
    // operation.  It will be set to NULL for other types of operations. The type
    // is bool.
    Rttmonechoadmintargetmepport interface{}

    // A string which represents the profile name to which a video probe should
    // use. This object is applicable to video probe only. The type is string with
    // length: 0..255.
    Rttmonechoadminvideotrafficprofile interface{}

    // This object represents the Differentiated Service Code Point (DSCP) QoS
    // marking in the generated synthetic packets.  Value - DiffServ Class     0 -
    // BE (default)    10 - AF11    12 - AF12    14 - AF13    18 - AF21    20 -
    // AF22    22 - AF23    26 - AF31    28 - AF32    30 - AF33    34 - AF41    36
    // - AF42    38 - AF43     8 - CS1    16 - CS2    24 - CS3    32 - CS4    40 -
    // CS5    48 - CS6    56 - CS7    46 - EF. The type is interface{} with range:
    // 0..63.
    Rttmonechoadmindscp interface{}

    // This object represents the video traffic generation source.  be : best
    // effort using DSP but without reservation gs : guaranteed service using DSP
    // with reservation na : not applicable for not using DSP. The type is
    // Rttmonechoadminreservedsp.
    Rttmonechoadminreservedsp interface{}

    // This object represents the network input interface on the sender router
    // where the synthetic packets are received from the emulated endpoint source.
    // This is used for path congruence with correct feature processing at the
    // sender router.  The user can get the InterfaceIndex number from ifIndex
    // object by looking up in ifTable. In fact, it should be useful to first get
    // the entry by the augmented table ifXTable which has ifName object which
    // matches the interface name used on the router or switch equipment console.
    // The type is interface{} with range: 0..2147483647.
    Rttmonechoadmininputinterface interface{}

    // This object specifies the IP address of the emulated source from which the
    // synthetic packets would be generated. If this object is not specified, the
    // emulated source IP address will by default be the same as
    // rttMonEchoAdminSourceAddress. This object is applicable to video probes.
    // The type is string.
    Rttmonechoadminemulatesourceaddress interface{}

    // This object represents the port number of the emulated source from which
    // the synthetic packets would be generated. If this object is not specified,
    // the emulated source port number will by default be the same as
    // rttMonEchoAdminSourcePort. This object is applicable to video probes. The
    // type is interface{} with range: 0..65536.
    Rttmonechoadminemulatesourceport interface{}

    // This object specifies the IP address of the emulated target by which the
    // synthetic packets would be received. If this object is not specified, the
    // emulated target IP address will by default be the same as
    // rttMonEchoAdminTargetAddress. This object is applicable to video probes.
    // The type is string.
    Rttmonechoadminemulatetargetaddress interface{}

    // This object represents the port number of the emulated target by which the
    // synthetic packets would be received. If this object is not specified, the
    // emulated target port number will by default be the same as
    // rttMonEchoAdminTargetPort. This object is applicable to video probes. The
    // type is interface{} with range: 0..65536.
    Rttmonechoadminemulatetargetport interface{}

    // This object indicates the MAC address of the target device. This object is
    // only applicable for Y.1731 operations.  rttMonEchoAdminTargetMacAddress and
    // rttMonEchoAdminTargetMPID may not be used in conjunction. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Rttmonechoadmintargetmacaddress interface{}

    // This object indicates the MAC address of the source device. This object is
    // only applicable for Y.1731 operations.  rttMonEchoAdminSourceMacAddress and
    // rttMonEchoAdminSourceMPID may not be used in conjunction. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Rttmonechoadminsourcemacaddress interface{}

    // This object indicates the source maintenance point ID.  It is only
    // applicable to Y.1731 operation.  It will be set to zero for other types of
    // opearations.  rttMonEchoAdminSourceMPID and rttMonEchoAdminSourceMacAddress
    // may not be used in conjunction. The type is interface{} with range:
    // 0..8191.
    Rttmonechoadminsourcempid interface{}

    // This object specifies the name of endpoint list which a probe uses to
    // generate operations. The type is string with length: 1..64.
    Rttmonechoadminendpointlistname interface{}

    // This object specifies if Source Specific Multicast is to be added. This
    // object is applicable to multicast probe only. The type is bool.
    Rttmonechoadminssm interface{}

    // This object specifies the maximum number of retries for control message.
    // The type is interface{} with range: 1..5.
    Rttmonechoadmincontrolretry interface{}

    // This object specifies the wait duration before control message timeout. The
    // type is interface{} with range: 1..10000. Units are milliseconds.
    Rttmonechoadmincontroltimeout interface{}

    // This object specifies number of packets to be sent for multicast tree
    // setup. This object is applicable to multicast probe only. The type is
    // interface{} with range: 0..10.
    Rttmonechoadminigmptreeinit interface{}

    // This object indicates that packets will be sent in burst. The type is bool.
    Rttmonechoadminenableburst interface{}

    // This object indicates the number of burst cycles to be sent during the
    // aggregate interval. This value is currently used for Y1731 SLM(Synthetic
    // Loss Measurment) probe. This object is applicable to Y1731 SLM probe only.
    // The type is interface{} with range: -2147483648..2147483647.
    Rttmonechoadminaggburstcycles interface{}

    // This object indicates the number of frames over which to calculate the
    // frame loss ratio. This object is applicable  to Y1731 SLM probe only. The
    // type is interface{} with range: -2147483648..2147483647.
    Rttmonechoadminlossrationumframes interface{}

    // This object indicates the number of frames over which to calculate the
    // availability. This object is applicable to Y1731 SLM probe only. The type
    // is interface{} with range: -2147483648..2147483647.
    Rttmonechoadminavailnumframes interface{}

    // This object specifies whether timestamp optimization is enabled.  When the
    // value is 'true' then timestamp optimization is enabled.  The probe will
    // utilize lower layer (Hardware/Packet Processor) timestamping values to
    // improve accuracy of statistics.  This value can be set only for udp jitter
    // operation with precision of microsecond. The type is bool.
    Rttmonechoadmintstampoptimization interface{}
}

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetFilter() yfilter.YFilter { return rttmonechoadminentry.YFilter }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) SetFilter(yf yfilter.YFilter) { rttmonechoadminentry.YFilter = yf }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonEchoAdminProtocol" { return "Rttmonechoadminprotocol" }
    if yname == "rttMonEchoAdminTargetAddress" { return "Rttmonechoadmintargetaddress" }
    if yname == "rttMonEchoAdminPktDataRequestSize" { return "Rttmonechoadminpktdatarequestsize" }
    if yname == "rttMonEchoAdminPktDataResponseSize" { return "Rttmonechoadminpktdataresponsesize" }
    if yname == "rttMonEchoAdminTargetPort" { return "Rttmonechoadmintargetport" }
    if yname == "rttMonEchoAdminSourceAddress" { return "Rttmonechoadminsourceaddress" }
    if yname == "rttMonEchoAdminSourcePort" { return "Rttmonechoadminsourceport" }
    if yname == "rttMonEchoAdminControlEnable" { return "Rttmonechoadmincontrolenable" }
    if yname == "rttMonEchoAdminTOS" { return "Rttmonechoadmintos" }
    if yname == "rttMonEchoAdminLSREnable" { return "Rttmonechoadminlsrenable" }
    if yname == "rttMonEchoAdminTargetAddressString" { return "Rttmonechoadmintargetaddressstring" }
    if yname == "rttMonEchoAdminNameServer" { return "Rttmonechoadminnameserver" }
    if yname == "rttMonEchoAdminOperation" { return "Rttmonechoadminoperation" }
    if yname == "rttMonEchoAdminHTTPVersion" { return "Rttmonechoadminhttpversion" }
    if yname == "rttMonEchoAdminURL" { return "Rttmonechoadminurl" }
    if yname == "rttMonEchoAdminCache" { return "Rttmonechoadmincache" }
    if yname == "rttMonEchoAdminInterval" { return "Rttmonechoadmininterval" }
    if yname == "rttMonEchoAdminNumPackets" { return "Rttmonechoadminnumpackets" }
    if yname == "rttMonEchoAdminProxy" { return "Rttmonechoadminproxy" }
    if yname == "rttMonEchoAdminString1" { return "Rttmonechoadminstring1" }
    if yname == "rttMonEchoAdminString2" { return "Rttmonechoadminstring2" }
    if yname == "rttMonEchoAdminString3" { return "Rttmonechoadminstring3" }
    if yname == "rttMonEchoAdminString4" { return "Rttmonechoadminstring4" }
    if yname == "rttMonEchoAdminString5" { return "Rttmonechoadminstring5" }
    if yname == "rttMonEchoAdminMode" { return "Rttmonechoadminmode" }
    if yname == "rttMonEchoAdminVrfName" { return "Rttmonechoadminvrfname" }
    if yname == "rttMonEchoAdminCodecType" { return "Rttmonechoadmincodectype" }
    if yname == "rttMonEchoAdminCodecInterval" { return "Rttmonechoadmincodecinterval" }
    if yname == "rttMonEchoAdminCodecPayload" { return "Rttmonechoadmincodecpayload" }
    if yname == "rttMonEchoAdminCodecNumPackets" { return "Rttmonechoadmincodecnumpackets" }
    if yname == "rttMonEchoAdminICPIFAdvFactor" { return "Rttmonechoadminicpifadvfactor" }
    if yname == "rttMonEchoAdminLSPFECType" { return "Rttmonechoadminlspfectype" }
    if yname == "rttMonEchoAdminLSPSelector" { return "Rttmonechoadminlspselector" }
    if yname == "rttMonEchoAdminLSPReplyMode" { return "Rttmonechoadminlspreplymode" }
    if yname == "rttMonEchoAdminLSPTTL" { return "Rttmonechoadminlspttl" }
    if yname == "rttMonEchoAdminLSPExp" { return "Rttmonechoadminlspexp" }
    if yname == "rttMonEchoAdminPrecision" { return "Rttmonechoadminprecision" }
    if yname == "rttMonEchoAdminProbePakPriority" { return "Rttmonechoadminprobepakpriority" }
    if yname == "rttMonEchoAdminOWNTPSyncTolAbs" { return "Rttmonechoadminowntpsynctolabs" }
    if yname == "rttMonEchoAdminOWNTPSyncTolPct" { return "Rttmonechoadminowntpsynctolpct" }
    if yname == "rttMonEchoAdminOWNTPSyncTolType" { return "Rttmonechoadminowntpsynctoltype" }
    if yname == "rttMonEchoAdminCalledNumber" { return "Rttmonechoadmincallednumber" }
    if yname == "rttMonEchoAdminDetectPoint" { return "Rttmonechoadmindetectpoint" }
    if yname == "rttMonEchoAdminGKRegistration" { return "Rttmonechoadmingkregistration" }
    if yname == "rttMonEchoAdminSourceVoicePort" { return "Rttmonechoadminsourcevoiceport" }
    if yname == "rttMonEchoAdminCallDuration" { return "Rttmonechoadmincallduration" }
    if yname == "rttMonEchoAdminLSPReplyDscp" { return "Rttmonechoadminlspreplydscp" }
    if yname == "rttMonEchoAdminLSPNullShim" { return "Rttmonechoadminlspnullshim" }
    if yname == "rttMonEchoAdminTargetMPID" { return "Rttmonechoadmintargetmpid" }
    if yname == "rttMonEchoAdminTargetDomainName" { return "Rttmonechoadmintargetdomainname" }
    if yname == "rttMonEchoAdminTargetVLAN" { return "Rttmonechoadmintargetvlan" }
    if yname == "rttMonEchoAdminEthernetCOS" { return "Rttmonechoadminethernetcos" }
    if yname == "rttMonEchoAdminLSPVccvID" { return "Rttmonechoadminlspvccvid" }
    if yname == "rttMonEchoAdminTargetEVC" { return "Rttmonechoadmintargetevc" }
    if yname == "rttMonEchoAdminTargetMEPPort" { return "Rttmonechoadmintargetmepport" }
    if yname == "rttMonEchoAdminVideoTrafficProfile" { return "Rttmonechoadminvideotrafficprofile" }
    if yname == "rttMonEchoAdminDscp" { return "Rttmonechoadmindscp" }
    if yname == "rttMonEchoAdminReserveDsp" { return "Rttmonechoadminreservedsp" }
    if yname == "rttMonEchoAdminInputInterface" { return "Rttmonechoadmininputinterface" }
    if yname == "rttMonEchoAdminEmulateSourceAddress" { return "Rttmonechoadminemulatesourceaddress" }
    if yname == "rttMonEchoAdminEmulateSourcePort" { return "Rttmonechoadminemulatesourceport" }
    if yname == "rttMonEchoAdminEmulateTargetAddress" { return "Rttmonechoadminemulatetargetaddress" }
    if yname == "rttMonEchoAdminEmulateTargetPort" { return "Rttmonechoadminemulatetargetport" }
    if yname == "rttMonEchoAdminTargetMacAddress" { return "Rttmonechoadmintargetmacaddress" }
    if yname == "rttMonEchoAdminSourceMacAddress" { return "Rttmonechoadminsourcemacaddress" }
    if yname == "rttMonEchoAdminSourceMPID" { return "Rttmonechoadminsourcempid" }
    if yname == "rttMonEchoAdminEndPointListName" { return "Rttmonechoadminendpointlistname" }
    if yname == "rttMonEchoAdminSSM" { return "Rttmonechoadminssm" }
    if yname == "rttMonEchoAdminControlRetry" { return "Rttmonechoadmincontrolretry" }
    if yname == "rttMonEchoAdminControlTimeout" { return "Rttmonechoadmincontroltimeout" }
    if yname == "rttMonEchoAdminIgmpTreeInit" { return "Rttmonechoadminigmptreeinit" }
    if yname == "rttMonEchoAdminEnableBurst" { return "Rttmonechoadminenableburst" }
    if yname == "rttMonEchoAdminAggBurstCycles" { return "Rttmonechoadminaggburstcycles" }
    if yname == "rttMonEchoAdminLossRatioNumFrames" { return "Rttmonechoadminlossrationumframes" }
    if yname == "rttMonEchoAdminAvailNumFrames" { return "Rttmonechoadminavailnumframes" }
    if yname == "rttMonEchoAdminTstampOptimization" { return "Rttmonechoadmintstampoptimization" }
    return ""
}

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetSegmentPath() string {
    return "rttMonEchoAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonechoadminentry.Rttmonctrladminindex) + "']"
}

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonechoadminentry.Rttmonctrladminindex
    leafs["rttMonEchoAdminProtocol"] = rttmonechoadminentry.Rttmonechoadminprotocol
    leafs["rttMonEchoAdminTargetAddress"] = rttmonechoadminentry.Rttmonechoadmintargetaddress
    leafs["rttMonEchoAdminPktDataRequestSize"] = rttmonechoadminentry.Rttmonechoadminpktdatarequestsize
    leafs["rttMonEchoAdminPktDataResponseSize"] = rttmonechoadminentry.Rttmonechoadminpktdataresponsesize
    leafs["rttMonEchoAdminTargetPort"] = rttmonechoadminentry.Rttmonechoadmintargetport
    leafs["rttMonEchoAdminSourceAddress"] = rttmonechoadminentry.Rttmonechoadminsourceaddress
    leafs["rttMonEchoAdminSourcePort"] = rttmonechoadminentry.Rttmonechoadminsourceport
    leafs["rttMonEchoAdminControlEnable"] = rttmonechoadminentry.Rttmonechoadmincontrolenable
    leafs["rttMonEchoAdminTOS"] = rttmonechoadminentry.Rttmonechoadmintos
    leafs["rttMonEchoAdminLSREnable"] = rttmonechoadminentry.Rttmonechoadminlsrenable
    leafs["rttMonEchoAdminTargetAddressString"] = rttmonechoadminentry.Rttmonechoadmintargetaddressstring
    leafs["rttMonEchoAdminNameServer"] = rttmonechoadminentry.Rttmonechoadminnameserver
    leafs["rttMonEchoAdminOperation"] = rttmonechoadminentry.Rttmonechoadminoperation
    leafs["rttMonEchoAdminHTTPVersion"] = rttmonechoadminentry.Rttmonechoadminhttpversion
    leafs["rttMonEchoAdminURL"] = rttmonechoadminentry.Rttmonechoadminurl
    leafs["rttMonEchoAdminCache"] = rttmonechoadminentry.Rttmonechoadmincache
    leafs["rttMonEchoAdminInterval"] = rttmonechoadminentry.Rttmonechoadmininterval
    leafs["rttMonEchoAdminNumPackets"] = rttmonechoadminentry.Rttmonechoadminnumpackets
    leafs["rttMonEchoAdminProxy"] = rttmonechoadminentry.Rttmonechoadminproxy
    leafs["rttMonEchoAdminString1"] = rttmonechoadminentry.Rttmonechoadminstring1
    leafs["rttMonEchoAdminString2"] = rttmonechoadminentry.Rttmonechoadminstring2
    leafs["rttMonEchoAdminString3"] = rttmonechoadminentry.Rttmonechoadminstring3
    leafs["rttMonEchoAdminString4"] = rttmonechoadminentry.Rttmonechoadminstring4
    leafs["rttMonEchoAdminString5"] = rttmonechoadminentry.Rttmonechoadminstring5
    leafs["rttMonEchoAdminMode"] = rttmonechoadminentry.Rttmonechoadminmode
    leafs["rttMonEchoAdminVrfName"] = rttmonechoadminentry.Rttmonechoadminvrfname
    leafs["rttMonEchoAdminCodecType"] = rttmonechoadminentry.Rttmonechoadmincodectype
    leafs["rttMonEchoAdminCodecInterval"] = rttmonechoadminentry.Rttmonechoadmincodecinterval
    leafs["rttMonEchoAdminCodecPayload"] = rttmonechoadminentry.Rttmonechoadmincodecpayload
    leafs["rttMonEchoAdminCodecNumPackets"] = rttmonechoadminentry.Rttmonechoadmincodecnumpackets
    leafs["rttMonEchoAdminICPIFAdvFactor"] = rttmonechoadminentry.Rttmonechoadminicpifadvfactor
    leafs["rttMonEchoAdminLSPFECType"] = rttmonechoadminentry.Rttmonechoadminlspfectype
    leafs["rttMonEchoAdminLSPSelector"] = rttmonechoadminentry.Rttmonechoadminlspselector
    leafs["rttMonEchoAdminLSPReplyMode"] = rttmonechoadminentry.Rttmonechoadminlspreplymode
    leafs["rttMonEchoAdminLSPTTL"] = rttmonechoadminentry.Rttmonechoadminlspttl
    leafs["rttMonEchoAdminLSPExp"] = rttmonechoadminentry.Rttmonechoadminlspexp
    leafs["rttMonEchoAdminPrecision"] = rttmonechoadminentry.Rttmonechoadminprecision
    leafs["rttMonEchoAdminProbePakPriority"] = rttmonechoadminentry.Rttmonechoadminprobepakpriority
    leafs["rttMonEchoAdminOWNTPSyncTolAbs"] = rttmonechoadminentry.Rttmonechoadminowntpsynctolabs
    leafs["rttMonEchoAdminOWNTPSyncTolPct"] = rttmonechoadminentry.Rttmonechoadminowntpsynctolpct
    leafs["rttMonEchoAdminOWNTPSyncTolType"] = rttmonechoadminentry.Rttmonechoadminowntpsynctoltype
    leafs["rttMonEchoAdminCalledNumber"] = rttmonechoadminentry.Rttmonechoadmincallednumber
    leafs["rttMonEchoAdminDetectPoint"] = rttmonechoadminentry.Rttmonechoadmindetectpoint
    leafs["rttMonEchoAdminGKRegistration"] = rttmonechoadminentry.Rttmonechoadmingkregistration
    leafs["rttMonEchoAdminSourceVoicePort"] = rttmonechoadminentry.Rttmonechoadminsourcevoiceport
    leafs["rttMonEchoAdminCallDuration"] = rttmonechoadminentry.Rttmonechoadmincallduration
    leafs["rttMonEchoAdminLSPReplyDscp"] = rttmonechoadminentry.Rttmonechoadminlspreplydscp
    leafs["rttMonEchoAdminLSPNullShim"] = rttmonechoadminentry.Rttmonechoadminlspnullshim
    leafs["rttMonEchoAdminTargetMPID"] = rttmonechoadminentry.Rttmonechoadmintargetmpid
    leafs["rttMonEchoAdminTargetDomainName"] = rttmonechoadminentry.Rttmonechoadmintargetdomainname
    leafs["rttMonEchoAdminTargetVLAN"] = rttmonechoadminentry.Rttmonechoadmintargetvlan
    leafs["rttMonEchoAdminEthernetCOS"] = rttmonechoadminentry.Rttmonechoadminethernetcos
    leafs["rttMonEchoAdminLSPVccvID"] = rttmonechoadminentry.Rttmonechoadminlspvccvid
    leafs["rttMonEchoAdminTargetEVC"] = rttmonechoadminentry.Rttmonechoadmintargetevc
    leafs["rttMonEchoAdminTargetMEPPort"] = rttmonechoadminentry.Rttmonechoadmintargetmepport
    leafs["rttMonEchoAdminVideoTrafficProfile"] = rttmonechoadminentry.Rttmonechoadminvideotrafficprofile
    leafs["rttMonEchoAdminDscp"] = rttmonechoadminentry.Rttmonechoadmindscp
    leafs["rttMonEchoAdminReserveDsp"] = rttmonechoadminentry.Rttmonechoadminreservedsp
    leafs["rttMonEchoAdminInputInterface"] = rttmonechoadminentry.Rttmonechoadmininputinterface
    leafs["rttMonEchoAdminEmulateSourceAddress"] = rttmonechoadminentry.Rttmonechoadminemulatesourceaddress
    leafs["rttMonEchoAdminEmulateSourcePort"] = rttmonechoadminentry.Rttmonechoadminemulatesourceport
    leafs["rttMonEchoAdminEmulateTargetAddress"] = rttmonechoadminentry.Rttmonechoadminemulatetargetaddress
    leafs["rttMonEchoAdminEmulateTargetPort"] = rttmonechoadminentry.Rttmonechoadminemulatetargetport
    leafs["rttMonEchoAdminTargetMacAddress"] = rttmonechoadminentry.Rttmonechoadmintargetmacaddress
    leafs["rttMonEchoAdminSourceMacAddress"] = rttmonechoadminentry.Rttmonechoadminsourcemacaddress
    leafs["rttMonEchoAdminSourceMPID"] = rttmonechoadminentry.Rttmonechoadminsourcempid
    leafs["rttMonEchoAdminEndPointListName"] = rttmonechoadminentry.Rttmonechoadminendpointlistname
    leafs["rttMonEchoAdminSSM"] = rttmonechoadminentry.Rttmonechoadminssm
    leafs["rttMonEchoAdminControlRetry"] = rttmonechoadminentry.Rttmonechoadmincontrolretry
    leafs["rttMonEchoAdminControlTimeout"] = rttmonechoadminentry.Rttmonechoadmincontroltimeout
    leafs["rttMonEchoAdminIgmpTreeInit"] = rttmonechoadminentry.Rttmonechoadminigmptreeinit
    leafs["rttMonEchoAdminEnableBurst"] = rttmonechoadminentry.Rttmonechoadminenableburst
    leafs["rttMonEchoAdminAggBurstCycles"] = rttmonechoadminentry.Rttmonechoadminaggburstcycles
    leafs["rttMonEchoAdminLossRatioNumFrames"] = rttmonechoadminentry.Rttmonechoadminlossrationumframes
    leafs["rttMonEchoAdminAvailNumFrames"] = rttmonechoadminentry.Rttmonechoadminavailnumframes
    leafs["rttMonEchoAdminTstampOptimization"] = rttmonechoadminentry.Rttmonechoadmintstampoptimization
    return leafs
}

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetYangName() string { return "rttMonEchoAdminEntry" }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) SetParent(parent types.Entity) { rttmonechoadminentry.parent = parent }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetParent() types.Entity { return rttmonechoadminentry.parent }

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetParentYangName() string { return "rttMonEchoAdminTable" }

// CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminlspfectype represents ldpIpv4Prefix   - LDP IPv4 prefix.
type CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminlspfectype string

const (
    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminlspfectype_ldpIpv4Prefix CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminlspfectype = "ldpIpv4Prefix"
)

// CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminowntpsynctoltype represents NTP sync tolerance is absolute value or percent value
type CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminowntpsynctoltype string

const (
    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminowntpsynctoltype_percent CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminowntpsynctoltype = "percent"

    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminowntpsynctoltype_absolute CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminowntpsynctoltype = "absolute"
)

// CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprecision represents This value can be set only for jitter operation
type CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprecision string

const (
    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprecision_milliseconds CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprecision = "milliseconds"

    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprecision_microseconds CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprecision = "microseconds"
)

// CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprobepakpriority represents operation
type CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprobepakpriority string

const (
    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprobepakpriority_normal CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprobepakpriority = "normal"

    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprobepakpriority_high CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminprobepakpriority = "high"
)

// CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp represents na : not applicable for not using DSP
type CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp string

const (
    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp_be CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp = "be"

    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp_gs CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp = "gs"

    CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp_na CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry_Rttmonechoadminreservedsp = "na"
)

// CISCORTTMONMIB_Rttmonfileioadmintable
// A table of Round Trip Time (RTT) monitoring 'fileIO'
// specific definitions.
// 
// When the RttMonRttType is not 'fileIO' this table is
// not valid.
// 
// This table is controlled via the 
// rttMonCtrlAdminTable.  Entries in this table are
// created via the rttMonCtrlAdminStatus object.
type CISCORTTMONMIB_Rttmonfileioadmintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for 'fileIO'
    // RttMonRttType conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry.
    Rttmonfileioadminentry []CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry
}

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetFilter() yfilter.YFilter { return rttmonfileioadmintable.YFilter }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) SetFilter(yf yfilter.YFilter) { rttmonfileioadmintable.YFilter = yf }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetGoName(yname string) string {
    if yname == "rttMonFileIOAdminEntry" { return "Rttmonfileioadminentry" }
    return ""
}

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetSegmentPath() string {
    return "rttMonFileIOAdminTable"
}

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonFileIOAdminEntry" {
        for _, c := range rttmonfileioadmintable.Rttmonfileioadminentry {
            if rttmonfileioadmintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry{}
        rttmonfileioadmintable.Rttmonfileioadminentry = append(rttmonfileioadmintable.Rttmonfileioadminentry, child)
        return &rttmonfileioadmintable.Rttmonfileioadminentry[len(rttmonfileioadmintable.Rttmonfileioadminentry)-1]
    }
    return nil
}

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonfileioadmintable.Rttmonfileioadminentry {
        children[rttmonfileioadmintable.Rttmonfileioadminentry[i].GetSegmentPath()] = &rttmonfileioadmintable.Rttmonfileioadminentry[i]
    }
    return children
}

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetYangName() string { return "rttMonFileIOAdminTable" }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) SetParent(parent types.Entity) { rttmonfileioadmintable.parent = parent }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetParent() types.Entity { return rttmonfileioadmintable.parent }

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry
// A list of objects that define specific configuration for
// 'fileIO' RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // The fully qualified file path that will be the target of the RTT operation.
    // This value must match one of the rttMonApplPreConfigedName entries. The
    // type is string.
    Rttmonfileioadminfilepath interface{}

    // The size of the file to write/read from the File Server. The type is
    // Rttmonfileioadminsize. Units are bytes.
    Rttmonfileioadminsize interface{}

    // The File I/O action to be performed. The type is Rttmonfileioadminaction.
    Rttmonfileioadminaction interface{}
}

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetFilter() yfilter.YFilter { return rttmonfileioadminentry.YFilter }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) SetFilter(yf yfilter.YFilter) { rttmonfileioadminentry.YFilter = yf }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonFileIOAdminFilePath" { return "Rttmonfileioadminfilepath" }
    if yname == "rttMonFileIOAdminSize" { return "Rttmonfileioadminsize" }
    if yname == "rttMonFileIOAdminAction" { return "Rttmonfileioadminaction" }
    return ""
}

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetSegmentPath() string {
    return "rttMonFileIOAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonfileioadminentry.Rttmonctrladminindex) + "']"
}

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonfileioadminentry.Rttmonctrladminindex
    leafs["rttMonFileIOAdminFilePath"] = rttmonfileioadminentry.Rttmonfileioadminfilepath
    leafs["rttMonFileIOAdminSize"] = rttmonfileioadminentry.Rttmonfileioadminsize
    leafs["rttMonFileIOAdminAction"] = rttmonfileioadminentry.Rttmonfileioadminaction
    return leafs
}

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetYangName() string { return "rttMonFileIOAdminEntry" }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) SetParent(parent types.Entity) { rttmonfileioadminentry.parent = parent }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetParent() types.Entity { return rttmonfileioadminentry.parent }

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetParentYangName() string { return "rttMonFileIOAdminTable" }

// CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction represents The File I/O action to be performed.
type CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction string

const (
    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction_write CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction = "write"

    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction_read CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction = "read"

    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction_writeRead CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminaction = "writeRead"
)

// CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize represents Server.
type CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize string

const (
    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize_n256 CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize = "n256"

    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize_n1k CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize = "n1k"

    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize_n64k CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize = "n64k"

    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize_n128k CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize = "n128k"

    CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize_n256k CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry_Rttmonfileioadminsize = "n256k"
)

// CISCORTTMONMIB_Rttmonscriptadmintable
// A table of Round Trip Time (RTT) monitoring 'script'
// specific definitions.
// 
// When the RttMonRttType is not 'script' this table is
// not valid.
// 
// This table is controlled via the
// rttMonCtrlAdminTable.  Entries in this table are
// created via the rttMonCtrlAdminStatus object.
type CISCORTTMONMIB_Rttmonscriptadmintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for 'script'
    // RttMonRttType conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry.
    Rttmonscriptadminentry []CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry
}

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetFilter() yfilter.YFilter { return rttmonscriptadmintable.YFilter }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) SetFilter(yf yfilter.YFilter) { rttmonscriptadmintable.YFilter = yf }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetGoName(yname string) string {
    if yname == "rttMonScriptAdminEntry" { return "Rttmonscriptadminentry" }
    return ""
}

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetSegmentPath() string {
    return "rttMonScriptAdminTable"
}

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonScriptAdminEntry" {
        for _, c := range rttmonscriptadmintable.Rttmonscriptadminentry {
            if rttmonscriptadmintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry{}
        rttmonscriptadmintable.Rttmonscriptadminentry = append(rttmonscriptadmintable.Rttmonscriptadminentry, child)
        return &rttmonscriptadmintable.Rttmonscriptadminentry[len(rttmonscriptadmintable.Rttmonscriptadminentry)-1]
    }
    return nil
}

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonscriptadmintable.Rttmonscriptadminentry {
        children[rttmonscriptadmintable.Rttmonscriptadminentry[i].GetSegmentPath()] = &rttmonscriptadmintable.Rttmonscriptadminentry[i]
    }
    return children
}

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetYangName() string { return "rttMonScriptAdminTable" }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) SetParent(parent types.Entity) { rttmonscriptadmintable.parent = parent }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetParent() types.Entity { return rttmonscriptadmintable.parent }

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry
// A list of objects that define specific configuration for
// 'script' RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This will be the Name of the Script that will be used to generate RTT
    // operations.    This object must match one of the  rttMonApplPreConfigedName
    // entries. The type is string.
    Rttmonscriptadminname interface{}

    // This will be the actual command line parameters passed to the
    // rttMonScriptAdminName when being executed. The type is string.
    Rttmonscriptadmincmdlineparams interface{}
}

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetFilter() yfilter.YFilter { return rttmonscriptadminentry.YFilter }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) SetFilter(yf yfilter.YFilter) { rttmonscriptadminentry.YFilter = yf }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonScriptAdminName" { return "Rttmonscriptadminname" }
    if yname == "rttMonScriptAdminCmdLineParams" { return "Rttmonscriptadmincmdlineparams" }
    return ""
}

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetSegmentPath() string {
    return "rttMonScriptAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonscriptadminentry.Rttmonctrladminindex) + "']"
}

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonscriptadminentry.Rttmonctrladminindex
    leafs["rttMonScriptAdminName"] = rttmonscriptadminentry.Rttmonscriptadminname
    leafs["rttMonScriptAdminCmdLineParams"] = rttmonscriptadminentry.Rttmonscriptadmincmdlineparams
    return leafs
}

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetYangName() string { return "rttMonScriptAdminEntry" }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) SetParent(parent types.Entity) { rttmonscriptadminentry.parent = parent }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetParent() types.Entity { return rttmonscriptadminentry.parent }

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetParentYangName() string { return "rttMonScriptAdminTable" }

// CISCORTTMONMIB_Rttmonreacttriggeradmintable
// A table of which contains the list of conceptual RTT
// control rows that will start to collect data when a 
// reaction condition is violated and when 
// rttMonReactAdminActionType is set to one of the 
// following:
//   -  triggerOnly
//   -  trapAndTrigger
//   -  nmvtAndTrigger
//   -  trapNmvtAndTrigger
// or when a reaction condition is violated and when any of the
// row in rttMonReactTable has rttMonReactActionType as one of
// the following:
//   - triggerOnly
//   - trapAndTrigger
// 
// The goal of this table is to define one or more 
// additional conceptual RTT control rows that will become
// active and start to collect additional history and
// statistics (depending on the rows configuration values),
// when a problem has been detected.
// 
// If the conceptual RTT control row is undefined, and a 
// trigger occurs, no action will take place.  
// 
// If the conceptual RTT control row is scheduled to start 
// at a later time, triggering that row will have no effect.
// 
// If the conceptual RTT control row is currently active, 
// triggering that row will have no effect on that row, but 
// the rttMonReactTriggerOperState object will transition to 
// 'active'.
// 
// An entry in this table can only be triggered when
// it is not currently in a triggered state.  The
// object rttMonReactTriggerOperState will 
// reflect the state of each entry in this table.
type CISCORTTMONMIB_Rttmonreacttriggeradmintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that will be triggered when a reaction condition is
    // violated. The type is slice of
    // CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry.
    Rttmonreacttriggeradminentry []CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry
}

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetFilter() yfilter.YFilter { return rttmonreacttriggeradmintable.YFilter }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) SetFilter(yf yfilter.YFilter) { rttmonreacttriggeradmintable.YFilter = yf }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetGoName(yname string) string {
    if yname == "rttMonReactTriggerAdminEntry" { return "Rttmonreacttriggeradminentry" }
    return ""
}

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetSegmentPath() string {
    return "rttMonReactTriggerAdminTable"
}

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonReactTriggerAdminEntry" {
        for _, c := range rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry {
            if rttmonreacttriggeradmintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry{}
        rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry = append(rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry, child)
        return &rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry[len(rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry)-1]
    }
    return nil
}

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry {
        children[rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry[i].GetSegmentPath()] = &rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry[i]
    }
    return children
}

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetYangName() string { return "rttMonReactTriggerAdminTable" }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) SetParent(parent types.Entity) { rttmonreacttriggeradmintable.parent = parent }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetParent() types.Entity { return rttmonreacttriggeradmintable.parent }

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry
// A list of objects that will be triggered when
// a reaction condition is violated.
type CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. This object points to a single conceptual Rtt
    // control row.  If this row does not exist and this value is  triggered no
    // action will result.  The conceptual Rtt control row will be triggered for
    // the  rttMonCtrlOperRttLife length.  If this conceptual Rtt control row is
    // already active, rttMonCtrlOperRttLife will not be updated, and its life
    // will continue as previously  defined. The type is interface{} with range:
    // 1..2147483647.
    Rttmonreacttriggeradminrttmonctrladminindex interface{}

    // This object is used to create Trigger entries. The type is RowStatus.
    Rttmonreacttriggeradminstatus interface{}

    // This object takes on the value active when its associated entry in the 
    // rttMonReactTriggerAdminTable has been triggered.  When the associated entry
    // in the rttMonReactTriggerAdminTable is not under a trigger state, this
    // object will be pending.  When this object is in the active state this entry
    // can not be retriggered. The type is Rttmonreacttriggeroperstate.
    Rttmonreacttriggeroperstate interface{}
}

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetFilter() yfilter.YFilter { return rttmonreacttriggeradminentry.YFilter }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) SetFilter(yf yfilter.YFilter) { rttmonreacttriggeradminentry.YFilter = yf }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonReactTriggerAdminRttMonCtrlAdminIndex" { return "Rttmonreacttriggeradminrttmonctrladminindex" }
    if yname == "rttMonReactTriggerAdminStatus" { return "Rttmonreacttriggeradminstatus" }
    if yname == "rttMonReactTriggerOperState" { return "Rttmonreacttriggeroperstate" }
    return ""
}

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetSegmentPath() string {
    return "rttMonReactTriggerAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonreacttriggeradminentry.Rttmonctrladminindex) + "']" + "[rttMonReactTriggerAdminRttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonreacttriggeradminentry.Rttmonreacttriggeradminrttmonctrladminindex) + "']"
}

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonreacttriggeradminentry.Rttmonctrladminindex
    leafs["rttMonReactTriggerAdminRttMonCtrlAdminIndex"] = rttmonreacttriggeradminentry.Rttmonreacttriggeradminrttmonctrladminindex
    leafs["rttMonReactTriggerAdminStatus"] = rttmonreacttriggeradminentry.Rttmonreacttriggeradminstatus
    leafs["rttMonReactTriggerOperState"] = rttmonreacttriggeradminentry.Rttmonreacttriggeroperstate
    return leafs
}

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetYangName() string { return "rttMonReactTriggerAdminEntry" }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) SetParent(parent types.Entity) { rttmonreacttriggeradminentry.parent = parent }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetParent() types.Entity { return rttmonreacttriggeradminentry.parent }

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetParentYangName() string { return "rttMonReactTriggerAdminTable" }

// CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry_Rttmonreacttriggeroperstate represents this entry can not be retriggered.
type CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry_Rttmonreacttriggeroperstate string

const (
    CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry_Rttmonreacttriggeroperstate_active CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry_Rttmonreacttriggeroperstate = "active"

    CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry_Rttmonreacttriggeroperstate_pending CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry_Rttmonreacttriggeroperstate = "pending"
)

// CISCORTTMONMIB_Rttmonechopathadmintable
// A table to store the hop addresses in a Loose Source Routing
// path. Response times are computed along the specified path
// using ping.
// 
// This maximum table size is limited by the size of the 
// maximum number of hop addresses that can fit in an IP header,
// which is 8. The object rttMonEchoPathAdminEntry will reflect 
// this tables maximum number of entries.
// 
// This table is coupled with rttMonCtrlAdminStatus.
type CISCORTTMONMIB_Rttmonechopathadmintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that define intermediate hop's IP Address.  This entry
    // can be added only if the rttMonCtrlAdminRttType is 'echo'. The entry gets
    // deleted when the corresponding RTR entry, which has an index of
    // rttMonCtrlAdminIndex, is deleted. The type is slice of
    // CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry.
    Rttmonechopathadminentry []CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry
}

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetFilter() yfilter.YFilter { return rttmonechopathadmintable.YFilter }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) SetFilter(yf yfilter.YFilter) { rttmonechopathadmintable.YFilter = yf }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetGoName(yname string) string {
    if yname == "rttMonEchoPathAdminEntry" { return "Rttmonechopathadminentry" }
    return ""
}

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetSegmentPath() string {
    return "rttMonEchoPathAdminTable"
}

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonEchoPathAdminEntry" {
        for _, c := range rttmonechopathadmintable.Rttmonechopathadminentry {
            if rttmonechopathadmintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry{}
        rttmonechopathadmintable.Rttmonechopathadminentry = append(rttmonechopathadmintable.Rttmonechopathadminentry, child)
        return &rttmonechopathadmintable.Rttmonechopathadminentry[len(rttmonechopathadmintable.Rttmonechopathadminentry)-1]
    }
    return nil
}

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonechopathadmintable.Rttmonechopathadminentry {
        children[rttmonechopathadmintable.Rttmonechopathadminentry[i].GetSegmentPath()] = &rttmonechopathadmintable.Rttmonechopathadminentry[i]
    }
    return children
}

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetYangName() string { return "rttMonEchoPathAdminTable" }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) SetParent(parent types.Entity) { rttmonechopathadmintable.parent = parent }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetParent() types.Entity { return rttmonechopathadmintable.parent }

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry
// A list of objects that define intermediate hop's IP Address.
// 
// This entry can be added only if the rttMonCtrlAdminRttType is
// 'echo'. The entry gets deleted when the corresponding RTR entry,
// which has an index of rttMonCtrlAdminIndex, is deleted.
type CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonEchoPathAdminTable. This number represents the hop address number in
    // a specific  ping path. All the indicies should start from 1 and must be 
    // contiguous ie., entries should be  (say rttMonCtrlAdminIndex = 1)  1.1,
    // 1.2, 1.3, they cannot be 1.1, 1.2, 1.4. The type is interface{} with range:
    // 1..8.
    Rttmonechopathadminhopindex interface{}

    // A string which specifies the address of an intermediate hop's IP Address
    // for a RTT 'echo' operation. The type is string.
    Rttmonechopathadminhopaddress interface{}
}

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetFilter() yfilter.YFilter { return rttmonechopathadminentry.YFilter }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) SetFilter(yf yfilter.YFilter) { rttmonechopathadminentry.YFilter = yf }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonEchoPathAdminHopIndex" { return "Rttmonechopathadminhopindex" }
    if yname == "rttMonEchoPathAdminHopAddress" { return "Rttmonechopathadminhopaddress" }
    return ""
}

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetSegmentPath() string {
    return "rttMonEchoPathAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonechopathadminentry.Rttmonctrladminindex) + "']" + "[rttMonEchoPathAdminHopIndex='" + fmt.Sprintf("%v", rttmonechopathadminentry.Rttmonechopathadminhopindex) + "']"
}

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonechopathadminentry.Rttmonctrladminindex
    leafs["rttMonEchoPathAdminHopIndex"] = rttmonechopathadminentry.Rttmonechopathadminhopindex
    leafs["rttMonEchoPathAdminHopAddress"] = rttmonechopathadminentry.Rttmonechopathadminhopaddress
    return leafs
}

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetYangName() string { return "rttMonEchoPathAdminEntry" }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) SetParent(parent types.Entity) { rttmonechopathadminentry.parent = parent }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetParent() types.Entity { return rttmonechopathadminentry.parent }

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetParentYangName() string { return "rttMonEchoPathAdminTable" }

// CISCORTTMONMIB_Rttmongrpscheduleadmintable
// A table of Round Trip Time (RTT) monitoring group scheduling
// specific definitions.
// This table is used to create a conceptual group scheduling
// control row. The entries in this control row contain objects
// used to define group schedule configuration parameters.
// 
// The objects of this table will be used to schedule a group of
// probes identified by the conceptual rows of the
// rttMonCtrlAdminTable.
type CISCORTTMONMIB_Rttmongrpscheduleadmintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that define a conceptual group scheduling control row.
    // The type is slice of
    // CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry.
    Rttmongrpscheduleadminentry []CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry
}

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetFilter() yfilter.YFilter { return rttmongrpscheduleadmintable.YFilter }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) SetFilter(yf yfilter.YFilter) { rttmongrpscheduleadmintable.YFilter = yf }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetGoName(yname string) string {
    if yname == "rttMonGrpScheduleAdminEntry" { return "Rttmongrpscheduleadminentry" }
    return ""
}

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetSegmentPath() string {
    return "rttMonGrpScheduleAdminTable"
}

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonGrpScheduleAdminEntry" {
        for _, c := range rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry {
            if rttmongrpscheduleadmintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry{}
        rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry = append(rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry, child)
        return &rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry[len(rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry)-1]
    }
    return nil
}

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry {
        children[rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry[i].GetSegmentPath()] = &rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry[i]
    }
    return children
}

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetYangName() string { return "rttMonGrpScheduleAdminTable" }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) SetParent(parent types.Entity) { rttmongrpscheduleadmintable.parent = parent }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetParent() types.Entity { return rttmongrpscheduleadmintable.parent }

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry
// A list of objects that define a conceptual group scheduling
// control row.
type CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonGrpScheduleAdminTable.  This is a pseudo-random number selected by
    // the management station when creating a row via the
    // rttMonGrpScheduleAdminStatus object. If the pseudo-random number is already
    // in use an 'inconsistentValue' return code will be returned when set
    // operation is attempted. The type is interface{} with range: 1..2147483647.
    Rttmongrpscheduleadminindex interface{}

    // A string which holds the different probes which are to be group scheduled.
    // The probes can be specified in the following forms. (a) Individual ID's
    // with comma separated as 23,45,34. (b) Range form including hyphens with
    // multiple ranges being     separated by a comma as 1-10,12-34. (c) Mix of
    // the above two forms as 1,2,4-10,12,15,19-25.  Any whitespace in the string
    // is considered an error. Duplicates and overlapping ranges as an example
    // 1,2,3,2-10 are considered fine. For a single range like 1-20 the upper
    // value (in this example 20) must be greater than lower value (1), otherwise
    // it's treated as an error. The agent will not normalize the list e.g., it
    // will not change 1,2,1-10 or even 1,2,3,4,5,6.. to 1-10. The type is string
    // with length: 0..200.
    Rttmongrpscheduleadminprobes interface{}

    // Specifies the time duration over which all the probes have to be scheduled.
    // The type is interface{} with range: 0..604800. Units are seconds.
    Rttmongrpscheduleadminperiod interface{}

    // Specifies the duration between initiating each RTT operation for all the
    // probes specified in the group.  The value of this object is only effective
    // when both rttMonGrpScheduleAdminFreqMax and rttMonGrpScheduleAdminFreqMin 
    // have zero values. The type is interface{} with range: 0..604800. Units are
    // seconds.
    Rttmongrpscheduleadminfrequency interface{}

    // This object specifies the life of all the probes included in the object
    // rttMonGrpScheduleAdminProbes, that are getting group scheduled. This value
    // will be placed into rttMonScheduleAdminRttLife object for each of the
    // probes listed in rttMonGrpScheduleAdminProbes when this conceptual control
    // row becomes 'active'.  The value 2147483647 has a special meaning. When
    // this object is set to 2147483647, the rttMonCtrlOperRttLife object for all
    // the probes listed in rttMonGrpScheduleAdminProbes,  will not decrement. And
    // thus the life time of the probes will never end. The type is interface{}
    // with range: 0..2147483647. Units are seconds.
    Rttmongrpscheduleadminlife interface{}

    // This object specifies the ageout value of all the probes included in the
    // object rttMonGrpScheduleAdminProbes, that are getting group scheduled. This
    // value will be placed into rttMonScheduleAdminConceptRowAgeout object for
    // each of the probes listed in rttMonGrpScheduleAdminProbes when this
    // conceptual control row becomes 'active'.  When this value is set to zero,
    // the probes listed in rttMonGrpScheduleAdminProbes, will never ageout. The
    // type is interface{} with range: 0..2073600. Units are seconds.
    Rttmongrpscheduleadminageout interface{}

    // The status of the conceptual RTT group schedule control row.  In order for
    // this object to become active, the following row objects must be defined:  -
    // rttMonGrpScheduleAdminProbes  - rttMonGrpScheduleAdminPeriod All other
    // objects can assume default values.  The conceptual RTT group schedule
    // control row objects cannot be modified once this conceptual RTT group
    // schedule control row has been created. Once this object is in 'active'
    // status, it cannot be set to 'notInService'. When this object moves to
    // 'active' state it will schedule the probes of the rttMonCtrlAdminTable
    // which had been created using 'createAndWait'.  This object can be set to
    // 'destroy' from any value at any time. When this object is set to 'destroy'
    // it will stop all the probes of the rttMonCtrlAdminTable, which had been
    // group scheduled by it earlier, before destroying the RTT group schedule
    // control row. The type is RowStatus.
    Rttmongrpscheduleadminstatus interface{}

    // Specifies the max duration between initiating each RTT operation for all
    // the probes specified in the group.  If this is 0 and
    // rttMonGrpScheduleAdminFreqMin is also 0 then
    // rttMonGrpScheduleAdminFrequency becomes the fixed frequency. The type is
    // interface{} with range: 0..604800. Units are seconds.
    Rttmongrpscheduleadminfreqmax interface{}

    // Specifies the min duration between initiating each RTT operation for all
    // the probes specified in the group.  The value of this object cannot be
    // greater than the value of rttMonGrpScheduleAdminFreqMax.  If this is 0 and
    // rttMonGrpScheduleAdminFreqMax is 0 then rttMonGrpScheduleAdminFrequency
    // becomes the fixed frequency. The type is interface{} with range: 0..604800.
    // Units are seconds.
    Rttmongrpscheduleadminfreqmin interface{}

    // This is the time in seconds after which the member probes of this group
    // specified in rttMonGrpScheduleAdminProbes will transition to active state.
    // The type is interface{} with range: 0..604800. Units are seconds.
    Rttmongrpscheduleadminstarttime interface{}

    // Addition of members to an existing group will be allowed if this object is
    // set to TRUE (1). The members, IDs of which are mentioned in
    // rttMonGrpScheduleAdminProbes object are added to the existing group. The
    // type is bool.
    Rttmongrpscheduleadminadd interface{}

    // Removal of members from an existing group will be allowed if this object is
    // set to TRUE (1). The members, IDs of which are mentioned in
    // rttMonGrpScheduleAdminProbes object are removed from the existing group.
    // The type is bool.
    Rttmongrpscheduleadmindelete interface{}

    // When this is set to true then all members of this group will be stopped and
    // rescheduled using the previously set values of this group. The type is
    // bool.
    Rttmongrpscheduleadminreset interface{}
}

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetFilter() yfilter.YFilter { return rttmongrpscheduleadminentry.YFilter }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) SetFilter(yf yfilter.YFilter) { rttmongrpscheduleadminentry.YFilter = yf }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetGoName(yname string) string {
    if yname == "rttMonGrpScheduleAdminIndex" { return "Rttmongrpscheduleadminindex" }
    if yname == "rttMonGrpScheduleAdminProbes" { return "Rttmongrpscheduleadminprobes" }
    if yname == "rttMonGrpScheduleAdminPeriod" { return "Rttmongrpscheduleadminperiod" }
    if yname == "rttMonGrpScheduleAdminFrequency" { return "Rttmongrpscheduleadminfrequency" }
    if yname == "rttMonGrpScheduleAdminLife" { return "Rttmongrpscheduleadminlife" }
    if yname == "rttMonGrpScheduleAdminAgeout" { return "Rttmongrpscheduleadminageout" }
    if yname == "rttMonGrpScheduleAdminStatus" { return "Rttmongrpscheduleadminstatus" }
    if yname == "rttMonGrpScheduleAdminFreqMax" { return "Rttmongrpscheduleadminfreqmax" }
    if yname == "rttMonGrpScheduleAdminFreqMin" { return "Rttmongrpscheduleadminfreqmin" }
    if yname == "rttMonGrpScheduleAdminStartTime" { return "Rttmongrpscheduleadminstarttime" }
    if yname == "rttMonGrpScheduleAdminAdd" { return "Rttmongrpscheduleadminadd" }
    if yname == "rttMonGrpScheduleAdminDelete" { return "Rttmongrpscheduleadmindelete" }
    if yname == "rttMonGrpScheduleAdminReset" { return "Rttmongrpscheduleadminreset" }
    return ""
}

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetSegmentPath() string {
    return "rttMonGrpScheduleAdminEntry" + "[rttMonGrpScheduleAdminIndex='" + fmt.Sprintf("%v", rttmongrpscheduleadminentry.Rttmongrpscheduleadminindex) + "']"
}

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonGrpScheduleAdminIndex"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminindex
    leafs["rttMonGrpScheduleAdminProbes"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminprobes
    leafs["rttMonGrpScheduleAdminPeriod"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminperiod
    leafs["rttMonGrpScheduleAdminFrequency"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminfrequency
    leafs["rttMonGrpScheduleAdminLife"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminlife
    leafs["rttMonGrpScheduleAdminAgeout"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminageout
    leafs["rttMonGrpScheduleAdminStatus"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminstatus
    leafs["rttMonGrpScheduleAdminFreqMax"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminfreqmax
    leafs["rttMonGrpScheduleAdminFreqMin"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminfreqmin
    leafs["rttMonGrpScheduleAdminStartTime"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminstarttime
    leafs["rttMonGrpScheduleAdminAdd"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminadd
    leafs["rttMonGrpScheduleAdminDelete"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadmindelete
    leafs["rttMonGrpScheduleAdminReset"] = rttmongrpscheduleadminentry.Rttmongrpscheduleadminreset
    return leafs
}

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetYangName() string { return "rttMonGrpScheduleAdminEntry" }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) SetParent(parent types.Entity) { rttmongrpscheduleadminentry.parent = parent }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetParent() types.Entity { return rttmongrpscheduleadminentry.parent }

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetParentYangName() string { return "rttMonGrpScheduleAdminTable" }

// CISCORTTMONMIB_Rttmplsvpnmonctrltable
// A table of Auto SAA L3 MPLS VPN definitions.
// 
// The Auto SAA L3 MPLS VPN administration control is in multiple
// tables.
// 
// This first table, is used to create a conceptual Auto SAA L3
// MPLS VPN control row.  The following tables contain objects
// which used in type specific configurations, scheduling and
// reaction configurations. All of these tables will create the
// same conceptual control row as this table using this table's
// index as their own index.
// 
// In order to a row in this table to become active the following
// objects must be defined.
//   rttMplsVpnMonCtrlRttType,
//   rttMplsVpnMonCtrlVrfName and
//   rttMplsVpnMonSchedulePeriod.
type CISCORTTMONMIB_Rttmplsvpnmonctrltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual Auto SAA L3 MPLS VPN
    // control row. The type is slice of
    // CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry.
    Rttmplsvpnmonctrlentry []CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry
}

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetFilter() yfilter.YFilter { return rttmplsvpnmonctrltable.YFilter }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) SetFilter(yf yfilter.YFilter) { rttmplsvpnmonctrltable.YFilter = yf }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetGoName(yname string) string {
    if yname == "rttMplsVpnMonCtrlEntry" { return "Rttmplsvpnmonctrlentry" }
    return ""
}

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetSegmentPath() string {
    return "rttMplsVpnMonCtrlTable"
}

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMplsVpnMonCtrlEntry" {
        for _, c := range rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry {
            if rttmplsvpnmonctrltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry{}
        rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry = append(rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry, child)
        return &rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry[len(rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry)-1]
    }
    return nil
}

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry {
        children[rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry[i].GetSegmentPath()] = &rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry[i]
    }
    return children
}

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetYangName() string { return "rttMplsVpnMonCtrlTable" }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) SetParent(parent types.Entity) { rttmplsvpnmonctrltable.parent = parent }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetParent() types.Entity { return rttmplsvpnmonctrltable.parent }

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry
// A base list of objects that define a conceptual Auto SAA L3
// MPLS VPN control row.
type CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMplsVpnMonCtrlTable.  This is a pseudo-random number selected by the
    // management station when creating a row via the rttMplsVpnMonCtrlStatus
    // object.  If the pseudo-random number is already in use an
    // 'inconsistentValue' return code will be returned when set operation is
    // attempted. The type is interface{} with range: 1..2147483647.
    Rttmplsvpnmonctrlindex interface{}

    // The type of RTT operation to be performed for Auto SAA L3 MPLS VPN.  This
    // value must be set in the same PDU of rttMplsVpnMonCtrlStatus.  This value
    // must be set before setting any other parameter configuration of an Auto SAA
    // L3 MPLS VPN. The type is RttMplsVpnMonRttType.
    Rttmplsvpnmonctrlrtttype interface{}

    // This field is used to specify the VPN name for which the Auto SAA L3 MPLS
    // VPN RTT operation will be used.  This value must be set in the same PDU of
    // rttMplsVpnMonCtrlStatus.  The Auto SAA L3 MPLS VPN will find the PEs
    // participating in this VPN and configure RTT operation corresponding to
    // value specified in rttMplsVpnMonCtrlRttType.  If the VPN corresponds to the
    // value configured for this object doesn't exist 'inconsistentValue' error
    // will be returned.  The value 'saa-vrf-all' has a special meaning. When this
    // object is set to 'saa-vrf-all', all the VPNs in the PE will be discovered
    // and Auto SAA L3 MPLS VPN will configure RTT operations corresponding to all
    // these PEs with the value specified in rttMplsVpnMonCtrlRttType as type for
    // those operations.  So, the user should avoid using this string for a
    // particular VPN name when using this feature in order to avoid ambiguity.
    // The type is string with length: 0..32.
    Rttmplsvpnmonctrlvrfname interface{}

    // A string which is used by a managing application to identify the RTT
    // target.  This string will be configured as rttMonCtrlAdminTag for all the
    // operations configured by this Auto SAA L3 MPLS VPN.  The usage of this
    // value in Auto SAA L3 MPLS VPN is same as rttMonCtrlAdminTag in RTT
    // operation. The type is string with length: 0..255.
    Rttmplsvpnmonctrltag interface{}

    // This object defines an administrative threshold limit.  This value will be
    // configured as rttMonCtrlAdminThreshold for all the operations that will be
    // configured by the current Auto SAA L3 MPLS VPN.  The usage of this value in
    // Auto SAA L3 MPLS VPN is same as rttMonCtrlAdminThreshold. The type is
    // interface{} with range: 0..2147483647. Units are milliseconds.
    Rttmplsvpnmonctrlthreshold interface{}

    // Specifies the duration to wait for a RTT operation configured automatically
    // by the Auto SAA L3 MPLS VPN to complete.   The value of this object cannot
    // be set to a value which would specify a duration exceeding
    // rttMplsVpnMonScheduleFrequency.  The usage of this value in Auto SAA L3
    // MPLS VPN is similar to rttMonCtrlAdminTimeout. The type is interface{} with
    // range: 0..604800000. Units are milliseconds.
    Rttmplsvpnmonctrltimeout interface{}

    // Specifies the frequency at which the automatic PE addition should take
    // place if there is any for an Auto SAA L3 MPLS VPN.  New RTT operations
    // corresponding to the new PEs discovered will be created and scheduled.  The
    // default value for this object is 4 hours. The maximum value supported is 49
    // days. The type is interface{} with range: 1..70560. Units are minutes.
    Rttmplsvpnmonctrlscaninterval interface{}

    // Specifies the frequency at which the automatic PE deletion should take
    // place.  This object specifies the number of times of
    // rttMonMplslmCtrlScanInterval (rttMplsVpnMonCtrlDelScanFactor *
    // rttMplsVpnMonCtrlScanInterval) to wait before removing the PEs. This object
    // doesn't directly specify the explicit value to wait before removing the PEs
    // that were down.  If this object set 0 the entries will never removed. The
    // type is interface{} with range: 0..2147483647.
    Rttmplsvpnmonctrldelscanfactor interface{}

    // This object represents the EXP value that needs to be put as precedence bit
    // of an IP header. The type is interface{} with range: 0..7.
    Rttmplsvpnmonctrlexp interface{}

    // This object represents the native payload size that needs to be put on the
    // packet.  This value will be configured as rttMonEchoAdminPktDataRequestSize
    // for all the RTT operations configured by the current Auto SAA L3 MPLS VPN. 
    // The minimum request size for jitter probe is 16. The maximum for jitter
    // probe is 1500. The default request size is 32 for jitter probe.  For echo
    // and pathEcho default request size is 28. The minimum request size for echo
    // and pathEcho is 28 bytes. The type is interface{} with range: 0..16384.
    // Units are octets.
    Rttmplsvpnmonctrlrequestsize interface{}

    // When set to true, the resulting data in each RTT operation created by the
    // current Auto SAA L3 MPLS VPN is compared with the expected data. This
    // includes checking header information (if possible) and exact packet size. 
    // Any mismatch will be recorded in the rttMonStatsCollectVerifyErrors object
    // of each RTT operation created by the current Auto SAA L3 MPLS VPN. The type
    // is bool.
    Rttmplsvpnmonctrlverifydata interface{}

    // The storage type of this conceptual row. When set to 'nonVolatile', this
    // entry will be shown in 'show running' command and can be saved into
    // Non-volatile memory.  By Default the entry will not be saved into
    // Non-volatile memory.  This object can be set to either 'volatile' or
    // 'nonVolatile'. Other values are not applicable for this conceptual row and
    // are not supported. The type is StorageType.
    Rttmplsvpnmonctrlstoragetype interface{}

    // This object holds the list of probes ID's that are created by the Auto SAA
    // L3 MPLS VPN.  The probes will be specified in the following form. (a)
    // Individual ID's with comma separated as 1,5,3. (b) Range form including
    // hyphens with multiple ranges being     separated by comma as 1-10,12-34.
    // (c) Mix of the above two forms as 1,2,4-10,12,15,19-25. The type is string.
    Rttmplsvpnmonctrlprobelist interface{}

    // The status of the conceptual Auto SAA L3 MPLS VPN control row.  In order
    // for this object to become active rttMplsVpnMonCtrlRttType, 
    // rttMplsVpnMonCtrlVrfName and  rttMplsVpnMonSchedulePeriod objects must be
    // defined. All other objects can assume default values.  If the object is set
    // to 'createAndGo' rttMplsVpnMonCtrlRttType, rttMplsVpnMonCtrlVrfName and
    // rttMplsVpnMonSchedulePeriod needs to be set along with
    // rttMplsVpnMonCtrlStatus.  If the object is set to 'createAndWait'
    // rttMplsVpnMonCtrlRttType and rttMplsVpnMonCtrlVrfName needs to be set along
    // with rttMplsVpnMonCtrlStatus. rttMplsVpnMonSchedulePeriod needs to be
    // specified before setting rttMplsVpnMonCtrlStatus to 'active'.  The
    // following objects cannot be modified after creating the Auto SAA L3 MPLS
    // VPN conceptual row.   - rttMplsVpnMonCtrlRttType  -
    // rttMplsVpnMonCtrlVrfName  The following objects can be modified even after
    // creating the Auto SAA L3 MPLS VPN conceptual row by setting this object to
    // 'notInService'   - All other writable objects in rttMplsVpnMonCtrlTable
    // except    rttMplsVpnMonCtrlRttType and rttMplsVpnMonCtrlVrfName.  - Objects
    // in the rttMplsVpnMonTypeTable.  - Objects in the
    // rttMplsVpnMonScheduleTable.  The following objects can be modified as
    // needed without setting this object to 'notInService' even after creating
    // the Auto SAA L3 MPLS VPN conceptual row.   - Objects in
    // rttMplsVpnMonReactTable.  This object can be set to 'destroy' from any
    // value at any time. When this object is set to 'destroy' it will stop and
    // destroy all the probes created by this Auto SAA L3 MPLS VPN before
    // destroying Auto SAA L3 MPLS VPN control row. The type is RowStatus.
    Rttmplsvpnmonctrlstatus interface{}

    // When set to true, this implies that LPD (LSP Path Discovery) is enabled for
    // this row.  The Auto SAA L3 MPLS VPN will find all the paths to each of the
    // PE's and configure RTT operation with rttMonCtrlAdminRttType value as
    // 'lspGroup'. The 'lspGroup' probe will walk through the list of set of
    // information that uniquely identifies a path and send the LSP echo requests
    // across them. All these LSP echo requests sent for 1st path, 2nd path etc.
    // can be thought of as 'single probes' sent as a part of 'lspGroup'. These
    // single probes will of type 'rttMplsVpnMonCtrlRttType'.  'lspGroup' probe is
    // a superset of individual probes that will test multiple paths. For example
    // Suppose there are 10 paths to the target. One 'lspGroup' probe will be
    // created which will store all the information related to uniquely identify
    // the 10 paths. When the 'lspGroup' probe will run it will sweep through the
    // set of information for 1st path, 2nd path, 3rd path and so on till it has
    // tested all the paths. The type is bool.
    Rttmplsvpnmonctrllpd interface{}

    // This object holds the list of LPD Group IDs that are created for this Auto
    // SAA L3 MPLS VPN row.  This object will be applicable only when LSP Path
    // Discovery is enabled for this row.  The LPD Groups will be specified in the
    // following form. (a) Individual ID's with comma separated as 1,5,3. (b)
    // Range form including hyphens with multiple ranges being     separated by
    // comma as 1-10,12-34. (c) Mix of the above two forms as
    // 1,2,4-10,12,15,19-25. The type is string.
    Rttmplsvpnmonctrllpdgrplist interface{}

    // The completion time of the LSP Path Discovery for the entire set of PEs
    // which are discovered for this Auto SAA.  This object will be applicable
    // only when LSP Path Discovery is enabled for this row. The type is
    // interface{} with range: 1..65535. Units are minutes.
    Rttmplsvpnmonctrllpdcomptime interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. This value is currently used for Jitter probe. This object is
    // applicable to jitter probe only.  The usage of this value in RTT operation
    // is same as rttMonEchoAdminInterval. The type is interface{} with range:
    // 1..60000. Units are milliseconds.
    Rttmplsvpnmontypeinterval interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is currently used for Jitter probe. This object is applicable to
    // jitter probe only.  The usage of this value in RTT operation is same as
    // rttMonEchoAdminNumPackets. The type is interface{} with range: 1..60000.
    Rttmplsvpnmontypenumpackets interface{}

    // This object represents the target's port number to which the packets need
    // to be sent.  This value will be configured as target port for all the
    // operations that is going to be configured   The usage of this value is same
    // as rttMonEchoAdminTargetPort in RTT operation. This object is applicable to
    // jitter type.  If this object is not being set random port will be used as
    // destination port. The type is interface{} with range: 1..65536.
    Rttmplsvpnmontypedestport interface{}

    // This object specifies the reaction type for which the
    // rttMplsVpnMonTypeSecFreqValue should be applied.  The Value 'timeout' will
    // cause secondary frequency to be set for frequency on timeout condition. 
    // The Value 'connectionLoss' will cause secondary frequency to be set for
    // frequency on connectionloss condition.  The Value 'both' will cause
    // secondary frequency to be set for frequency on either of
    // timeout/connectionloss condition.  Notifications must be configured on
    // corresponding reaction type in order to rttMplsVpnMonTypeSecFreqValue get
    // effect.  When LSP Path Discovery is enabled for this row the following
    // rttMplsVpnMonReactLpdNotifyType notifications must be configured in order
    // to rttMplsVpnMonTypeSecFreqValue get effect.   - 'lpdGroupStatus' or
    // 'lpdAll'.  Since the Frequency of the operation changes the stats will be
    // collected in new bucket.  If any of the reaction type
    // (timeout/connectionLoss) occurred for an operation configured by this Auto
    // SAA L3 MPLS VPN and the following conditions are satisfied, the frequency
    // of the operation will be changed to rttMplsVpnMonTypeSecFreqValue.    1)
    // rttMplsVpnMonTypeSecFreqType is set for a reaction type  
    // (timeout/connectionLoss).   2) A notification is configured for the same
    // reaction type   (timeout/connectionLoss).  When LSP Path Discovery is
    // enabled for this row, if any of the reaction type (timeout/connectionLoss)
    // occurred for 'single probes' configured by this Auto SAA L3 MPLS VPN and
    // the following conditions are satisfied, the secondary frequency
    // rttMplsVpnMonTypeSecFreqValue will be applied to the 'lspGroup' probe.   
    // 1) rttMplsVpnMonTypeSecFreqType is set for a reaction type  
    // (timeout/connectionLoss/both).   2) rttMplsVpnMonReactLpdNotifyType object
    // must be set to   value of 'lpdGroupStatus' or 'lpdAll'.  The frequency of
    // the individual operations will be restored to original frequency once the
    // trap is sent. The type is Rttmplsvpnmontypesecfreqtype.
    Rttmplsvpnmontypesecfreqtype interface{}

    // This object represents the value that needs to be applied to secondary
    // frequency of individual RTT operations configured by Auto SAA L3 MPLS VPN. 
    // Setting rttMplsVpnMonTypeSecFreqValue without setting
    // rttMplsVpnMonTypeSecFreqType will not have any effect. The type is
    // interface{} with range: 1..604800.
    Rttmplsvpnmontypesecfreqvalue interface{}

    // A string which specifies the address of the local host (127.X.X.X).  This
    // object will be used as lsp-selector in MPLS RTT operations configured by
    // the Auto SAA L3 MPLS VPN.  When LSP Path Discovery is enabled for the row,
    // this object will be used to indicate the base LSP selector value to be used
    // in the LSP Path Discovery.  This value of this object is significant in
    // MPLS load balancing scenario. This value will be used as one of the
    // parameter in that load balancing. The type is string.
    Rttmplsvpnmontypelspselector interface{}

    // This object specifies the reply mode for the LSP Echo requests originated
    // by the operations configured by the Auto SAA L3 MPLS VPN.  This object is
    // currently used by echo and pathEcho. The type is RttMonLSPPingReplyMode.
    Rttmplsvpnmontypelspreplymode interface{}

    // This object represents the TTL setting for MPLS echo request packets
    // originated by the operations configured by the Auto SAA L3 MPLS VPN.  This
    // object is currently used by echo and pathEcho.  For 'echo' the default TTL
    // will be set to 255. For 'pathEcho' the default will be set to 30.  Note:
    // This object cannot be set to the value of 0. The default value of 0
    // signifies the default TTL values will be used for 'echo' and 'pathEcho'.
    // The type is interface{} with range: 0..255.
    Rttmplsvpnmontypelspttl interface{}

    // This object specifies the DSCP value to be set in the IP header of the LSP
    // echo reply packet. The value of this object will be in range of DiffServ
    // codepoint values between 0 to 63.  Note: This object cannot be set to value
    // of 255. This default value specifies that DSCP is not set for this row. The
    // type is interface{} with range: 0..63 | 255..None.
    Rttmplsvpnmontypelspreplydscp interface{}

    // This object represents the number of concurrent path discovery requests
    // that will be active at one time per MPLS VPN control row. This object is
    // meant for reducing the time for discovery of all the paths to the target in
    // a large customer network. However its value should be chosen such that it
    // does not cause any performance impact.  Note: If the customer network has
    // low end routers in the Core it is recommended to keep this value low. The
    // type is interface{} with range: 1..15.
    Rttmplsvpnmontypelpdmaxsessions interface{}

    // This object specifies the maximum allowed duration of a particular tree
    // trace request.  If no response is received in configured time the request
    // will be considered a failure. The type is interface{} with range: 1..900.
    // Units are seconds.
    Rttmplsvpnmontypelpdsesstimeout interface{}

    // This object specifies the timeout value for the LSP echo requests which are
    // sent while performing the LSP Path  Discovery. The type is interface{} with
    // range: 0..604800000. Units are milliseconds.
    Rttmplsvpnmontypelpdechotimeout interface{}

    // This object specifies the send interval between LSP echo requests which are
    // sent while performing the LSP Path  Discovery. The type is interface{} with
    // range: 0..3600000. Units are milliseconds.
    Rttmplsvpnmontypelpdechointerval interface{}

    // This object specifies if the explicit-null label is added to LSP echo
    // requests which are sent while performing the LSP Path Discovery.  If set to
    // TRUE all the probes configured as part of this control row will send the
    // LSP echo requests with the explicit-null label added. The type is bool.
    Rttmplsvpnmontypelpdechonullshim interface{}

    // This object specifies the scan time for the completion of LSP Path
    // Discovery for all the PEs discovered for this control row. If the scan
    // period is exceeded on completion of the LSP Path Discovery for all the PEs,
    // the next discovery will start immediately else it will wait till expiry of
    // scan period.  For example: If the value is set to 30 minutes then on start
    // of the LSP Path Discovery a timestamp will be taken say T1. At the end of
    // the tree trace discovery one more timestamp will be taken again say T2. If
    // (T2-T1) is greater than 30, the next discovery will start immediately else
    // next discovery  will wait for [30 - (T2-T1)].  Note: If the object is set
    // to a special value of '0', it will force immediate start of the next
    // discovery on all neighbours without any delay. The type is interface{} with
    // range: 0..7200. Units are minutes.
    Rttmplsvpnmontypelpdscanperiod interface{}

    // The maximum number of hours of data to be kept per LPD group. The LPD group
    // statistics will be kept in an hourly bucket. At the maximum there can be
    // two buckets. The value of 'one' is not advisable because the group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value used in the
    // rttMplsVpnLpdGroupStatsTable to uniquely identify this group is the
    // rttMonStatsCaptureStartTimeIndex.  Note: When this object is set to the
    // value of '0' all rttMplsVpnLpdGroupStatsTable data capturing will be shut
    // off. The type is interface{} with range: 0..2.
    Rttmplsvpnmontypelpdstathours interface{}

    // This is the time when this conceptual row will activate.
    // rttMplsVpnMonSchedulePeriod object must be specified before setting this
    // object.  This is the value of MIB-II's sysUpTime in the future. When
    // sysUpTime equals this value this object will cause the activation of a
    // conceptual Auto SAA L3 MPLS VPN row.  When an agent has the capability to
    // determine date and time, the agent should store this object as DateAndTime.
    // This allows the agent to be able to activate conceptual Auto SAA L3 MPLS
    // VPN row at the intended time.  If this object has value as 1, this means
    // start the operation now itself. Value of 0 puts the operation in pending
    // state. The type is interface{} with range: 0..4294967295.
    Rttmplsvpnmonschedulerttstarttime interface{}

    // Specifies the time duration over which all the probes created by the
    // current Auto SAA L3 MPLS VPN have to be scheduled.  This object must be set
    // first before setting rttMplsVpnMonScheduleRttStartTime. The type is
    // interface{} with range: 1..604800. Units are seconds.
    Rttmplsvpnmonscheduleperiod interface{}

    // Specifies the duration between initiating each RTT operation configured by
    // the Auto SAA L3 MPLS VPN.  This object cannot be set to a value which would
    // be a shorter duration than rttMplsVpnMonCtrlTimeout.  The usage of this
    // value in RTT operation is same as rttMonCtrlAdminFrequency. The type is
    // interface{} with range: 1..604800. Units are seconds.
    Rttmplsvpnmonschedulefrequency interface{}

    // The value set for this will be applied as rttMonReactAdminConnectionEnable
    // for individual probes created by the Auto SAA L3 MPLS VPN.  When this
    // object is set to true, rttMonReactVar for individual probes created by the
    // Auto SAA L3 MPLS VPN will be set to 'connectionLoss(8)'. The type is bool.
    Rttmplsvpnmonreactconnectionenable interface{}

    // The value set for this will be applied as rttMonReactAdminTimeoutEnable for
    // individual probes created by the Auto SAA L3 MPLS VPN.  When this object is
    // set to true, rttMonReactVar for individual probes created by the Auto SAA
    // L3 MPLS VPN will be set to 'timeout(7)'. The type is bool.
    Rttmplsvpnmonreacttimeoutenable interface{}

    // The value corresponding to this object will be applied as
    // rttMonReactAdminThresholdType for individual probes created by the Auto SAA
    // L3 MPLS VPN.  The value corresponding to this object will be applied as
    // rttMonReactThresholdType for individual probes created by the Auto SAA L3
    // MPLS VPN. The type is Rttmplsvpnmonreactthresholdtype.
    Rttmplsvpnmonreactthresholdtype interface{}

    // This object value will be applied as rttMonReactAdminThresholdCount for
    // individual probes created by the Auto SAA L3 MPLS VPN.  This object value
    // will be applied as rttMonReactThresholdCountX for individual probes created
    // by the Auto SAA L3 MPLS VPN. The type is interface{} with range: 1..16.
    Rttmplsvpnmonreactthresholdcount interface{}

    // The value corresponding to this object will be applied as
    // rttMonReactAdminActionType of individual probes created by this Auto SAA L3
    // MPLS VPN.  The value corresponding to this object will be applied as
    // rttMonReactActionType of individual probes created by this Auto SAA L3 MPLS
    // VPN. The type is Rttmplsvpnmonreactactiontype.
    Rttmplsvpnmonreactactiontype interface{}

    // This object specifies the type of LPD notifications to be generated for the
    // current Auto SAA L3 MPLS VPN control row.  This object will be applicable
    // only when LSP Path Discovery is enabled for this row.  There are two types
    // of notifications supported for the LPD - (a) rttMonLpdDiscoveryNotification
    // - This notification will     be sent on the failure of LSP Path Discovery
    // to the     particular PE. Reversal of the failure will also result in    
    // sending the notification. (b) rttMonLpdGrpStatusNotification - Individual
    // probes in an LPD     group will not generate notifications independently
    // but will     be generating dependent on the state of the group. Any    
    // individual probe can initiate the generation of a     notification,
    // dependent on the state of the group.     Notifications are only generated
    // if the failure/restoration     of an individual probe causes the state of
    // the group to     change.  The Value 'none' will not cause any notifications
    // to be sent.  The Value 'lpdPathDiscovery' will cause (a) to be sent.  The
    // Value 'lpdGroupStatus' will cause (b) to be sent.  The Value 'lpdAll' will
    // cause both (a) and (b) to sent depending on the failure conditions. The
    // type is Rttmplsvpnmonreactlpdnotifytype.
    Rttmplsvpnmonreactlpdnotifytype interface{}

    // This object value specifies the number of attempts to be performed before
    // declaring the path as 'down'. Each 'single probe' which is part of
    // 'lspGroup' probe will be retried these many times before marking it as
    // 'down'.  This object will be applicable only when LSP Path Discovery is
    // enabled for this row.    - When rttMplsVpnMonTypeSecFreqType is not
    // configured, the     failure count will be incremented at the next cycle of 
    // 'lspGroup' probe at interval's of     rttMplsVpnMonScheduleFrequency value.
    // For example: Assume there are 10 paths discovered and on     the first run
    // of the 'lspGroup' probe first two paths failed     and rest passed. On the
    // second run all the probes will be      run again. The probes 1 and 2 will
    // be retried till the     rttMplsVpnMonReactLpdRetryCount value, and     then
    // marked as 'down' and rttMonLpdGrpStatusNotification      will be sent if
    // configured.    - When rttMplsVpnMonTypeSecFreqType value is anything other 
    // than 'none', the retry will happen for the failed probes at     the
    // rttMplsVpnMonTypeSecFreqValue and only the failed     probes will be
    // retried.      For example: Assume there are 10 paths discovered and on the 
    // first run of the 'lspGroup' probe first two paths failed and     rest
    // passed. The secondary frequency will be applied to the     failed probes.
    // At secondary frequency interval the first two     probes will be run again.
    // The probes 1 and 2 will be retried     till the
    // rttMplsVpnMonReactLpdRetryCount value, and     then marked as 'down' and
    // rttMonLpdGrpStatusNotification      will be sent if configured. The type is
    // interface{} with range: 1..16. Units are attempts.
    Rttmplsvpnmonreactlpdretrycount interface{}
}

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetFilter() yfilter.YFilter { return rttmplsvpnmonctrlentry.YFilter }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) SetFilter(yf yfilter.YFilter) { rttmplsvpnmonctrlentry.YFilter = yf }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetGoName(yname string) string {
    if yname == "rttMplsVpnMonCtrlIndex" { return "Rttmplsvpnmonctrlindex" }
    if yname == "rttMplsVpnMonCtrlRttType" { return "Rttmplsvpnmonctrlrtttype" }
    if yname == "rttMplsVpnMonCtrlVrfName" { return "Rttmplsvpnmonctrlvrfname" }
    if yname == "rttMplsVpnMonCtrlTag" { return "Rttmplsvpnmonctrltag" }
    if yname == "rttMplsVpnMonCtrlThreshold" { return "Rttmplsvpnmonctrlthreshold" }
    if yname == "rttMplsVpnMonCtrlTimeout" { return "Rttmplsvpnmonctrltimeout" }
    if yname == "rttMplsVpnMonCtrlScanInterval" { return "Rttmplsvpnmonctrlscaninterval" }
    if yname == "rttMplsVpnMonCtrlDelScanFactor" { return "Rttmplsvpnmonctrldelscanfactor" }
    if yname == "rttMplsVpnMonCtrlEXP" { return "Rttmplsvpnmonctrlexp" }
    if yname == "rttMplsVpnMonCtrlRequestSize" { return "Rttmplsvpnmonctrlrequestsize" }
    if yname == "rttMplsVpnMonCtrlVerifyData" { return "Rttmplsvpnmonctrlverifydata" }
    if yname == "rttMplsVpnMonCtrlStorageType" { return "Rttmplsvpnmonctrlstoragetype" }
    if yname == "rttMplsVpnMonCtrlProbeList" { return "Rttmplsvpnmonctrlprobelist" }
    if yname == "rttMplsVpnMonCtrlStatus" { return "Rttmplsvpnmonctrlstatus" }
    if yname == "rttMplsVpnMonCtrlLpd" { return "Rttmplsvpnmonctrllpd" }
    if yname == "rttMplsVpnMonCtrlLpdGrpList" { return "Rttmplsvpnmonctrllpdgrplist" }
    if yname == "rttMplsVpnMonCtrlLpdCompTime" { return "Rttmplsvpnmonctrllpdcomptime" }
    if yname == "rttMplsVpnMonTypeInterval" { return "Rttmplsvpnmontypeinterval" }
    if yname == "rttMplsVpnMonTypeNumPackets" { return "Rttmplsvpnmontypenumpackets" }
    if yname == "rttMplsVpnMonTypeDestPort" { return "Rttmplsvpnmontypedestport" }
    if yname == "rttMplsVpnMonTypeSecFreqType" { return "Rttmplsvpnmontypesecfreqtype" }
    if yname == "rttMplsVpnMonTypeSecFreqValue" { return "Rttmplsvpnmontypesecfreqvalue" }
    if yname == "rttMplsVpnMonTypeLspSelector" { return "Rttmplsvpnmontypelspselector" }
    if yname == "rttMplsVpnMonTypeLSPReplyMode" { return "Rttmplsvpnmontypelspreplymode" }
    if yname == "rttMplsVpnMonTypeLSPTTL" { return "Rttmplsvpnmontypelspttl" }
    if yname == "rttMplsVpnMonTypeLSPReplyDscp" { return "Rttmplsvpnmontypelspreplydscp" }
    if yname == "rttMplsVpnMonTypeLpdMaxSessions" { return "Rttmplsvpnmontypelpdmaxsessions" }
    if yname == "rttMplsVpnMonTypeLpdSessTimeout" { return "Rttmplsvpnmontypelpdsesstimeout" }
    if yname == "rttMplsVpnMonTypeLpdEchoTimeout" { return "Rttmplsvpnmontypelpdechotimeout" }
    if yname == "rttMplsVpnMonTypeLpdEchoInterval" { return "Rttmplsvpnmontypelpdechointerval" }
    if yname == "rttMplsVpnMonTypeLpdEchoNullShim" { return "Rttmplsvpnmontypelpdechonullshim" }
    if yname == "rttMplsVpnMonTypeLpdScanPeriod" { return "Rttmplsvpnmontypelpdscanperiod" }
    if yname == "rttMplsVpnMonTypeLpdStatHours" { return "Rttmplsvpnmontypelpdstathours" }
    if yname == "rttMplsVpnMonScheduleRttStartTime" { return "Rttmplsvpnmonschedulerttstarttime" }
    if yname == "rttMplsVpnMonSchedulePeriod" { return "Rttmplsvpnmonscheduleperiod" }
    if yname == "rttMplsVpnMonScheduleFrequency" { return "Rttmplsvpnmonschedulefrequency" }
    if yname == "rttMplsVpnMonReactConnectionEnable" { return "Rttmplsvpnmonreactconnectionenable" }
    if yname == "rttMplsVpnMonReactTimeoutEnable" { return "Rttmplsvpnmonreacttimeoutenable" }
    if yname == "rttMplsVpnMonReactThresholdType" { return "Rttmplsvpnmonreactthresholdtype" }
    if yname == "rttMplsVpnMonReactThresholdCount" { return "Rttmplsvpnmonreactthresholdcount" }
    if yname == "rttMplsVpnMonReactActionType" { return "Rttmplsvpnmonreactactiontype" }
    if yname == "rttMplsVpnMonReactLpdNotifyType" { return "Rttmplsvpnmonreactlpdnotifytype" }
    if yname == "rttMplsVpnMonReactLpdRetryCount" { return "Rttmplsvpnmonreactlpdretrycount" }
    return ""
}

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetSegmentPath() string {
    return "rttMplsVpnMonCtrlEntry" + "[rttMplsVpnMonCtrlIndex='" + fmt.Sprintf("%v", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlindex) + "']"
}

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMplsVpnMonCtrlIndex"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlindex
    leafs["rttMplsVpnMonCtrlRttType"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlrtttype
    leafs["rttMplsVpnMonCtrlVrfName"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlvrfname
    leafs["rttMplsVpnMonCtrlTag"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrltag
    leafs["rttMplsVpnMonCtrlThreshold"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlthreshold
    leafs["rttMplsVpnMonCtrlTimeout"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrltimeout
    leafs["rttMplsVpnMonCtrlScanInterval"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlscaninterval
    leafs["rttMplsVpnMonCtrlDelScanFactor"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrldelscanfactor
    leafs["rttMplsVpnMonCtrlEXP"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlexp
    leafs["rttMplsVpnMonCtrlRequestSize"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlrequestsize
    leafs["rttMplsVpnMonCtrlVerifyData"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlverifydata
    leafs["rttMplsVpnMonCtrlStorageType"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlstoragetype
    leafs["rttMplsVpnMonCtrlProbeList"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlprobelist
    leafs["rttMplsVpnMonCtrlStatus"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlstatus
    leafs["rttMplsVpnMonCtrlLpd"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrllpd
    leafs["rttMplsVpnMonCtrlLpdGrpList"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrllpdgrplist
    leafs["rttMplsVpnMonCtrlLpdCompTime"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonctrllpdcomptime
    leafs["rttMplsVpnMonTypeInterval"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypeinterval
    leafs["rttMplsVpnMonTypeNumPackets"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypenumpackets
    leafs["rttMplsVpnMonTypeDestPort"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypedestport
    leafs["rttMplsVpnMonTypeSecFreqType"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypesecfreqtype
    leafs["rttMplsVpnMonTypeSecFreqValue"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypesecfreqvalue
    leafs["rttMplsVpnMonTypeLspSelector"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspselector
    leafs["rttMplsVpnMonTypeLSPReplyMode"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspreplymode
    leafs["rttMplsVpnMonTypeLSPTTL"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspttl
    leafs["rttMplsVpnMonTypeLSPReplyDscp"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspreplydscp
    leafs["rttMplsVpnMonTypeLpdMaxSessions"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdmaxsessions
    leafs["rttMplsVpnMonTypeLpdSessTimeout"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdsesstimeout
    leafs["rttMplsVpnMonTypeLpdEchoTimeout"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdechotimeout
    leafs["rttMplsVpnMonTypeLpdEchoInterval"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdechointerval
    leafs["rttMplsVpnMonTypeLpdEchoNullShim"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdechonullshim
    leafs["rttMplsVpnMonTypeLpdScanPeriod"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdscanperiod
    leafs["rttMplsVpnMonTypeLpdStatHours"] = rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdstathours
    leafs["rttMplsVpnMonScheduleRttStartTime"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonschedulerttstarttime
    leafs["rttMplsVpnMonSchedulePeriod"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonscheduleperiod
    leafs["rttMplsVpnMonScheduleFrequency"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonschedulefrequency
    leafs["rttMplsVpnMonReactConnectionEnable"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonreactconnectionenable
    leafs["rttMplsVpnMonReactTimeoutEnable"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonreacttimeoutenable
    leafs["rttMplsVpnMonReactThresholdType"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonreactthresholdtype
    leafs["rttMplsVpnMonReactThresholdCount"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonreactthresholdcount
    leafs["rttMplsVpnMonReactActionType"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonreactactiontype
    leafs["rttMplsVpnMonReactLpdNotifyType"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonreactlpdnotifytype
    leafs["rttMplsVpnMonReactLpdRetryCount"] = rttmplsvpnmonctrlentry.Rttmplsvpnmonreactlpdretrycount
    return leafs
}

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetYangName() string { return "rttMplsVpnMonCtrlEntry" }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) SetParent(parent types.Entity) { rttmplsvpnmonctrlentry.parent = parent }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetParent() types.Entity { return rttmplsvpnmonctrlentry.parent }

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetParentYangName() string { return "rttMplsVpnMonCtrlTable" }

// CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactactiontype represents this Auto SAA L3 MPLS VPN.
type CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactactiontype string

const (
    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactactiontype_none CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactactiontype = "none"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactactiontype_trapOnly CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactactiontype = "trapOnly"
)

// CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype represents depending on the failure conditions.
type CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype string

const (
    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype_none CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype = "none"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype_lpdPathDiscovery CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype = "lpdPathDiscovery"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype_lpdGroupStatus CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype = "lpdGroupStatus"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype_lpdAll CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactlpdnotifytype = "lpdAll"
)

// CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype represents the Auto SAA L3 MPLS VPN.
type CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype string

const (
    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype_never CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype = "never"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype_immediate CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype = "immediate"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype_consecutive CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmonreactthresholdtype = "consecutive"
)

// CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype represents original frequency once the trap is sent.
type CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype string

const (
    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype_none CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype = "none"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype_timeout CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype = "timeout"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype_connectionLoss CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype = "connectionLoss"

    CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype_both CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry_Rttmplsvpnmontypesecfreqtype = "both"
)

// CISCORTTMONMIB_Rttmonreacttable
// A table that contains the reaction configurations. Each
// conceptual row in rttMonReactTable corresponds to a reaction
// configured for the probe defined in rttMonCtrlAdminTable.
// 
// For each reaction configured for a probe there is an entry in
// the table.
// 
// Each Probe can have multiple reactions and hence there can be
// multiple rows for a particular probe.
// 
// This table is coupled with rttMonCtrlAdminTable.
type CISCORTTMONMIB_Rttmonreacttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual reaction configuration
    // control row. The type is slice of
    // CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry.
    Rttmonreactentry []CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry
}

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetFilter() yfilter.YFilter { return rttmonreacttable.YFilter }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) SetFilter(yf yfilter.YFilter) { rttmonreacttable.YFilter = yf }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetGoName(yname string) string {
    if yname == "rttMonReactEntry" { return "Rttmonreactentry" }
    return ""
}

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetSegmentPath() string {
    return "rttMonReactTable"
}

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonReactEntry" {
        for _, c := range rttmonreacttable.Rttmonreactentry {
            if rttmonreacttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry{}
        rttmonreacttable.Rttmonreactentry = append(rttmonreacttable.Rttmonreactentry, child)
        return &rttmonreacttable.Rttmonreactentry[len(rttmonreacttable.Rttmonreactentry)-1]
    }
    return nil
}

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonreacttable.Rttmonreactentry {
        children[rttmonreacttable.Rttmonreactentry[i].GetSegmentPath()] = &rttmonreacttable.Rttmonreactentry[i]
    }
    return children
}

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetYangName() string { return "rttMonReactTable" }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) SetParent(parent types.Entity) { rttmonreacttable.parent = parent }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetParent() types.Entity { return rttmonreacttable.parent }

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry
// A base list of objects that define a conceptual reaction
// configuration control row.
type CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. This object along with rttMonCtrlAdminIndex
    // identifies a particular reaction-configuration for a particular probe. This
    // is a pseudo-random number selected by the management station when creating
    // a row via the rttMonReactStatus. If the pseudo-random number is already in
    // use an 'inconsistentValue' return code will be returned when set operation
    // is attempted. The type is interface{} with range: 1..2147483647.
    Rttmonreactconfigindex interface{}

    // This object specifies the type of reaction configured for a probe.  The
    // reaction types 'rtt', 'timeout', and 'connectionLoss'  can be configured
    // for all probe types.  The reaction type 'verifyError' can be configured for
    // all  probe types except RTP probe type.  The reaction types 'jitterSDAvg',
    // 'jitterDSAvg', 'jitterAvg',  'packetLateArrival', 'packetOutOfSequence', 
    // 'maxOfPositiveSD', 'maxOfNegativeSD', 'maxOfPositiveDS' and
    // 'maxOfNegativeDS' can be configured for UDP jitter  and ICMP jitter probe
    // types only.  The reaction types 'mos' and 'icpif' can be configured  for
    // UDP jitter and ICMP jitter probe types only.  The reaction types
    // 'packetLossDS', 'packetLossSD' and  'packetMIA' can be configured for UDP
    // jitter, and  RTP probe types only.  The reaction types 'iaJitterDS',
    // 'frameLossDS', 'mosLQDS',  'mosCQDS', 'rFactorDS', 'iaJitterSD',
    // 'rFactorSD', 'mosCQSD'  can be configured for RTP probe type only.  The
    // reaction types 'successivePacketLoss', 'maxOfLatencyDS',  'maxOfLatencySD',
    // 'latencyDSAvg', 'latencySDAvg' and  'packetLoss' can be configured for ICMP
    // jitter probe  type only. The type is RttMonReactVar.
    Rttmonreactvar interface{}

    // This object specifies the conditions under which the notification ( trap )
    // is sent.  never       - rttMonReactOccurred is never set  immediate   -
    // rttMonReactOccurred is set to 'true' when the               value of
    // parameter for which reaction is               configured ( e.g rtt,
    // jitterAvg, packetLossSD,               mos etc ) violates the threshold.   
    // Conversely, rttMonReactOccurred is set to 'false'               when the
    // parameter ( e.g rtt, jitterAvg,               packetLossSD, mos etc ) is
    // below the threshold               limits.  consecutive -
    // rttMonReactOccurred is set to true when the value               of
    // parameter for which reaction is configured               ( e.g rtt,
    // jitterAvg, packetLossSD, mos etc )               violates the threshold for
    // configured consecutive               times.               Conversely,
    // rttMonReactOccurred is set to false               when the value of
    // parameter ( e.g rtt, jitterAvg               packetLossSD, mos etc ) is
    // below the threshold               limits for the same number of consecutive
    // operations.  xOfy        - rttMonReactOccurred is set to true when x       
    // ( as specified by rttMonReactThresholdCountX )               out of the
    // last y ( as specified by               rttMonReacthresholdCountY ) times
    // the value of               parameter for which the reaction is configured  
    // ( e.g rtt, jitterAvg, packetLossSD, mos etc )               violates the
    // threshold.               Conversely, it is set to false when x, out of the 
    // last y times the value of parameter               ( e.g rtt, jitterAvg,
    // packetLossSD, mos ) is               below the threshold limits.           
    // NOTE: When x > y, the probe will never                     generate a
    // reaction.  average    - rttMonReactOccurred is set to true when the        
    // average ( rttMonReactThresholdCountX times )              value of
    // parameter for which reaction is               configured ( e.g rtt,
    // jitterAvg, packetLossSD,              mos etc ) violates the threshold
    // condition.              Conversely, it is set to false when the            
    // average value of parameter ( e.g rtt, jitterAvg,              packetLossSD,
    // mos etc ) is below the threshold              limits.  If this value is
    // changed by a management station, rttMonReactOccurred is set to false, but
    // no reaction is generated if the prior value of rttMonReactOccurred was
    // true. The type is Rttmonreactthresholdtype.
    Rttmonreactthresholdtype interface{}

    // Specifies what type(s), if any, of reaction(s) to generate if an operation
    // violates one of the watched ( reaction-configuration ) conditions:  none   
    // - no reaction is generated trapOnly           - a trap is generated
    // triggerOnly        - all trigger actions defined for this                  
    // entry are initiated trapAndTrigger     - both a trap and all trigger
    // actions                      are initiated A trigger action is defined via
    // the rttMonReactTriggerAdminTable. The type is Rttmonreactactiontype.
    Rttmonreactactiontype interface{}

    // This object defines the higher threshold limit. If the value ( e.g rtt,
    // jitterAvg, packetLossSD etc ) rises above this limit and if the condition
    // specified in rttMonReactThresholdType are satisfied, a trap is generated. 
    // Default value of rttMonReactThresholdRising for    'rtt' is 5000   
    // 'jitterAvg' is 100.    'jitterSDAvg' is 100.    'jitterDSAvg' 100.   
    // 'packetLossSD' is 10000.    'packetLossDS' is 10000.    'mos' is 500.   
    // 'icpif' is 93.    'packetMIA' is 10000.    'packetLateArrival' is 10000.   
    // 'packetOutOfSequence' is 10000.    'maxOfPositiveSD' is 10000.   
    // 'maxOfNegativeSD' is 10000.    'maxOfPositiveDS' is 10000.   
    // 'maxOfNegativeDS' is 10000.    'iaJitterDS' is 20.    'frameLossDS' is
    // 10000.    'mosLQDS' is 400.    'mosCQDS' is 400.    'rFactorDS' is 80.   
    // 'successivePacketLoss' is 1000.    'maxOfLatencyDS' is 5000.   
    // 'maxOfLatencySD' is 5000.    'latencyDSAvg' is 5000.    'latencySDAvg' is
    // 5000.    'packetLoss' is 10000.  This object is not applicable if the
    // rttMonReactVar is 'timeout', 'connectionLoss' or 'verifyError'. For
    // 'timeout', 'connectionLoss' and 'verifyError' default value of 
    // rttMonReactThresholdRising will be 0. The type is interface{} with range:
    // 0..2147483647.
    Rttmonreactthresholdrising interface{}

    // This object defines a lower threshold limit. If the value ( e.g rtt,
    // jitterAvg, packetLossSD etc ) falls below this limit and if the conditions
    // specified in rttMonReactThresholdType are satisfied, a trap is generated. 
    // Default value of rttMonReactThresholdFalling    'rtt' is 3000   
    // 'jitterAvg' is 100.    'jitterSDAvg' is 100.    'jitterDSAvg' 100.   
    // 'packetLossSD' is 10000.    'packetLossDS' is 10000.    'mos' is 500.   
    // 'icpif' is 93.    'packetMIA' is 10000.    'packetLateArrival' is 10000.   
    // 'packetOutOfSequence' is 10000.    'maxOfPositiveSD' is 10000.   
    // 'maxOfNegativeSD' is 10000.    'maxOfPositiveDS' is 10000.   
    // 'maxOfNegativeDS' is 10000.    'iaJitterDS' is 20.    'frameLossDS' is
    // 10000.    'mosLQDS' is 310.    'mosCQDS' is 310.    'rFactorDS' is 60.   
    // 'successivePacketLoss' is 1000.    'maxOfLatencyDS' is 3000.   
    // 'maxOfLatencySD' is 3000.    'latencyDSAvg' is 3000.    'latencySDAvg' is
    // 3000.    'packetLoss' is 10000.    'iaJitterSD' is 20.    'mosCQSD' is 310.
    // 'rFactorSD' is 60.  This object is not applicable if the rttMonReactVar is
    // 'timeout', 'connectionLoss' or 'verifyError'. For 'timeout',
    // 'connectionLoss' and 'verifyError' default value of
    // rttMonReactThresholdFalling will be 0. The type is interface{} with range:
    // 0..2147483647.
    Rttmonreactthresholdfalling interface{}

    // If rttMonReactThresholdType value is 'xOfy', this object defines the 'x'
    // value.  If rttMonReactThresholdType value is 'consecutive' this object
    // defines the number of consecutive occurrences that needs threshold
    // violation before setting  rttMonReactOccurred as true.  If
    // rttMonReactThresholdType value is 'average' this object defines the number
    // of samples that needs be considered for calculating average.  This object
    // has no meaning if rttMonReactThresholdType has value of 'never' and
    // 'immediate'. The type is interface{} with range: 1..16.
    Rttmonreactthresholdcountx interface{}

    // This object defines the 'y' value of the xOfy condition if
    // rttMonReactThresholdType is 'xOfy'.  For other values of
    // rttMonReactThresholdType, this object is not applicable. The type is
    // interface{} with range: 1..16.
    Rttmonreactthresholdcounty interface{}

    // This object will be set when the configured threshold condition is violated
    // as defined by rttMonReactThresholdType and holds the actual value that
    // violated the configured threshold values.  This object is not valid for the
    // following values of rttMonReactVar and It will be always 0:   - timeout   -
    // connectionLoss   - verifyError. The type is interface{} with range:
    // 0..2147483647.
    Rttmonreactvalue interface{}

    // This object is set to true when the configured threshold condition is
    // violated as defined by rttMonReactThresholdType. It will be again set to
    // 'false' if the condition reverses.  This object is set to true in the
    // following conditions:  - rttMonReactVar is set to timeout and   
    // rttMonCtrlOperTimeoutOccurred set to true.  - rttMonReactVar is set to
    // connectionLoss and    rttMonCtrlOperConnectionLostOccurred set to true.  -
    // rttMonReactVar is set to verifyError and   
    // rttMonCtrlOperVerifyErrorOccurred is set to true.  - For all other values
    // of rttMonReactVar, if the    corresponding value exceeds the configured   
    // rttMonReactThresholdRising.   This object is set to false in the following
    // conditions:  - rttMonReactVar is set to timeout and   
    // rttMonCtrlOperTimeoutOccurred set to false.  - rttMonReactVar is set to
    // connectionLoss and     rttMonCtrlOperConnectionLostOccurred set to false. 
    // - rttMonReactVar is set to verifyError and   
    // rttMonCtrlOperVerifyErrorOccurred is set to false.  - For all other values
    // of rttMonReactVar, if the    corresponding value fall below the configured 
    // rttMonReactThresholdFalling.  When the RttMonRttType is 'pathEcho' or
    // 'pathJitter', this object is applied only to the 
    // rttMonEchoAdminTargetAddress and not to intermediate hops to the Target.
    // The type is bool.
    Rttmonreactoccurred interface{}

    // This objects indicates the status of the conceptual RTT Reaction Control
    // Row.Only CreateAndGo and destroy  operations are permitted on the row. 
    // When this object moves to active state, the conceptual row having the
    // Reaction configuration for the probe is monitored and the notifications are
    // generated when the threshold violation takes place.  In order for this
    // object to become active rttMonReactVar must be defined. All other objects
    // assume the default value.  This object can be set to 'destroy' from any
    // value at any time. When this object is set to 'destroy' no reaction
    // configuration for the probes would exist. The reaction configuration for
    // the probe is removed. The type is RowStatus.
    Rttmonreactstatus interface{}
}

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetFilter() yfilter.YFilter { return rttmonreactentry.YFilter }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) SetFilter(yf yfilter.YFilter) { rttmonreactentry.YFilter = yf }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonReactConfigIndex" { return "Rttmonreactconfigindex" }
    if yname == "rttMonReactVar" { return "Rttmonreactvar" }
    if yname == "rttMonReactThresholdType" { return "Rttmonreactthresholdtype" }
    if yname == "rttMonReactActionType" { return "Rttmonreactactiontype" }
    if yname == "rttMonReactThresholdRising" { return "Rttmonreactthresholdrising" }
    if yname == "rttMonReactThresholdFalling" { return "Rttmonreactthresholdfalling" }
    if yname == "rttMonReactThresholdCountX" { return "Rttmonreactthresholdcountx" }
    if yname == "rttMonReactThresholdCountY" { return "Rttmonreactthresholdcounty" }
    if yname == "rttMonReactValue" { return "Rttmonreactvalue" }
    if yname == "rttMonReactOccurred" { return "Rttmonreactoccurred" }
    if yname == "rttMonReactStatus" { return "Rttmonreactstatus" }
    return ""
}

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetSegmentPath() string {
    return "rttMonReactEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonreactentry.Rttmonctrladminindex) + "']" + "[rttMonReactConfigIndex='" + fmt.Sprintf("%v", rttmonreactentry.Rttmonreactconfigindex) + "']"
}

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonreactentry.Rttmonctrladminindex
    leafs["rttMonReactConfigIndex"] = rttmonreactentry.Rttmonreactconfigindex
    leafs["rttMonReactVar"] = rttmonreactentry.Rttmonreactvar
    leafs["rttMonReactThresholdType"] = rttmonreactentry.Rttmonreactthresholdtype
    leafs["rttMonReactActionType"] = rttmonreactentry.Rttmonreactactiontype
    leafs["rttMonReactThresholdRising"] = rttmonreactentry.Rttmonreactthresholdrising
    leafs["rttMonReactThresholdFalling"] = rttmonreactentry.Rttmonreactthresholdfalling
    leafs["rttMonReactThresholdCountX"] = rttmonreactentry.Rttmonreactthresholdcountx
    leafs["rttMonReactThresholdCountY"] = rttmonreactentry.Rttmonreactthresholdcounty
    leafs["rttMonReactValue"] = rttmonreactentry.Rttmonreactvalue
    leafs["rttMonReactOccurred"] = rttmonreactentry.Rttmonreactoccurred
    leafs["rttMonReactStatus"] = rttmonreactentry.Rttmonreactstatus
    return leafs
}

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetYangName() string { return "rttMonReactEntry" }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) SetParent(parent types.Entity) { rttmonreactentry.parent = parent }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetParent() types.Entity { return rttmonreactentry.parent }

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetParentYangName() string { return "rttMonReactTable" }

// CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype represents rttMonReactTriggerAdminTable.
type CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype string

const (
    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype_none CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype = "none"

    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype_trapOnly CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype = "trapOnly"

    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype_triggerOnly CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype = "triggerOnly"

    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype_trapAndTrigger CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactactiontype = "trapAndTrigger"
)

// CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype represents rttMonReactOccurred was true.
type CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype string

const (
    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype_never CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype = "never"

    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype_immediate CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype = "immediate"

    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype_consecutive CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype = "consecutive"

    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype_xOfy CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype = "xOfy"

    CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype_average CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry_Rttmonreactthresholdtype = "average"
)

// CISCORTTMONMIB_Rttmongeneratedopertable
// This table contains information about the generated
// operation id as part of a parent IP SLA operation. The parent
// operation id is pseudo-random number, selected by the management 
// station based on an operation started by the management 
// station,when creating a row via the rttMonCtrlAdminStatus
// object in the rttMonCtrlAdminTable table.
type CISCORTTMONMIB_Rttmongeneratedopertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the Generated Oper table corresponding to a child or generated
    // operation as part of a parent IP SLA operation. The type is slice of
    // CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry.
    Rttmongeneratedoperentry []CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry
}

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetFilter() yfilter.YFilter { return rttmongeneratedopertable.YFilter }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) SetFilter(yf yfilter.YFilter) { rttmongeneratedopertable.YFilter = yf }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetGoName(yname string) string {
    if yname == "rttMonGeneratedOperEntry" { return "Rttmongeneratedoperentry" }
    return ""
}

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetSegmentPath() string {
    return "rttMonGeneratedOperTable"
}

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonGeneratedOperEntry" {
        for _, c := range rttmongeneratedopertable.Rttmongeneratedoperentry {
            if rttmongeneratedopertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry{}
        rttmongeneratedopertable.Rttmongeneratedoperentry = append(rttmongeneratedopertable.Rttmongeneratedoperentry, child)
        return &rttmongeneratedopertable.Rttmongeneratedoperentry[len(rttmongeneratedopertable.Rttmongeneratedoperentry)-1]
    }
    return nil
}

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmongeneratedopertable.Rttmongeneratedoperentry {
        children[rttmongeneratedopertable.Rttmongeneratedoperentry[i].GetSegmentPath()] = &rttmongeneratedopertable.Rttmongeneratedoperentry[i]
    }
    return children
}

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetYangName() string { return "rttMonGeneratedOperTable" }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) SetParent(parent types.Entity) { rttmongeneratedopertable.parent = parent }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetParent() types.Entity { return rttmongeneratedopertable.parent }

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry
// An entry in the Generated Oper table corresponding to
// a child or generated operation as part of a parent
// IP SLA operation.
type CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. The type of Internet address, IPv4 or IPv6, of a
    // responder for an IP SLA operation. The type is InetAddressType.
    Rttmongeneratedoperrespipaddrtype interface{}

    // This attribute is a key. The internet address of a responder for IP SLA
    // operation. The type of this address is determined by the value of
    // rttMonGeneratedOperRespIpAddrType. The type is string with length: 0..255.
    Rttmongeneratedoperrespipaddr interface{}

    // This is a pseudo-random number, auto-generated based to identify a child
    // operation based on a parent  operation started by the management
    // station,when  creating a row via the rttMonCtrlAdminStatus object. The type
    // is interface{} with range: 1..2147483647.
    Rttmongeneratedoperctrladminindex interface{}
}

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetFilter() yfilter.YFilter { return rttmongeneratedoperentry.YFilter }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) SetFilter(yf yfilter.YFilter) { rttmongeneratedoperentry.YFilter = yf }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonGeneratedOperRespIpAddrType" { return "Rttmongeneratedoperrespipaddrtype" }
    if yname == "rttMonGeneratedOperRespIpAddr" { return "Rttmongeneratedoperrespipaddr" }
    if yname == "rttMonGeneratedOperCtrlAdminIndex" { return "Rttmongeneratedoperctrladminindex" }
    return ""
}

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetSegmentPath() string {
    return "rttMonGeneratedOperEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmongeneratedoperentry.Rttmonctrladminindex) + "']" + "[rttMonGeneratedOperRespIpAddrType='" + fmt.Sprintf("%v", rttmongeneratedoperentry.Rttmongeneratedoperrespipaddrtype) + "']" + "[rttMonGeneratedOperRespIpAddr='" + fmt.Sprintf("%v", rttmongeneratedoperentry.Rttmongeneratedoperrespipaddr) + "']"
}

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmongeneratedoperentry.Rttmonctrladminindex
    leafs["rttMonGeneratedOperRespIpAddrType"] = rttmongeneratedoperentry.Rttmongeneratedoperrespipaddrtype
    leafs["rttMonGeneratedOperRespIpAddr"] = rttmongeneratedoperentry.Rttmongeneratedoperrespipaddr
    leafs["rttMonGeneratedOperCtrlAdminIndex"] = rttmongeneratedoperentry.Rttmongeneratedoperctrladminindex
    return leafs
}

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetYangName() string { return "rttMonGeneratedOperEntry" }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) SetParent(parent types.Entity) { rttmongeneratedoperentry.parent = parent }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetParent() types.Entity { return rttmongeneratedoperentry.parent }

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetParentYangName() string { return "rttMonGeneratedOperTable" }

// CISCORTTMONMIB_Rttmonstatscapturetable
// The statistics capture database.
// 
// The statistics capture table contains summarized 
// information of the results for a conceptual RTT control 
// row.  A rolling accumulated history of this information 
// is maintained in a series of hourly 'group(s)'.  Each 
// 'group' contains a series of 'path(s)', each 'path' 
// contains a series of 'hop(s)', each 'hop' contains a 
// series of 'statistics distribution bucket(s)'.
// 
// Each conceptual statistics row has a current hourly 
// group, into which RTT results are accumulated.  At the 
// end of each hour a new hourly group is created which 
// then becomes current.  The counters and accumulators in 
// the new group are initialized to zero.  The previous 
// group(s) is kept in the table until the table contains 
// rttMonStatisticsAdminNumHourGroups groups for the 
// conceptual statistics row;  at this point, the oldest 
// group is discarded and is replaced by the newly created 
// one.  The hourly group is uniquely identified by the 
// rttMonStatsCaptureStartTimeIndex object.
// 
// If the activity for a conceptual RTT control row ceases 
// because the rttMonCtrlOperState object transitions to 
// 'inactive', the corresponding current hourly group in 
// this table is 'frozen', and a new hourly group is 
// created when activity is resumed.
// 
// If the activity for a conceptual RTT control row ceases 
// because the rttMonCtrlOperState object transitions to 
// 'pending' this whole table will be cleared and reset to 
// its initial state.
// 
// When the RttMonRttType is 'pathEcho', the path 
// exploration RTT requests' statistics will not be 
// accumulated in this table.
// 
// NOTE: When the RttMonRttType is 'pathEcho', a source to 
//       target rttMonStatsCapturePathIndex path will be 
//       created for each rttMonStatsCaptureStartTimeIndex 
//       to hold all errors that occur when a specific path
//       had not been found or connection has not be setup.
// 
// Using this rttMonStatsCaptureTable, a managing 
// application can retrieve summarized data from accurately 
// measured periods, which is synchronized across multiple 
// conceptual RTT control rows.  With the new hourly group
// creation being performed on a 60 minute period, the 
// managing station has plenty of time to collect the data, 
// and need not be concerned with the vagaries of network 
// delays and lost PDU's when trying to get matching data.  
// Also, the managing station can spread the data gathering 
// over a longer period, which removes the need for a flood 
// of get requests in a short period which otherwise would 
// occur.
type CISCORTTMONMIB_Rttmonstatscapturetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  The statistics capture table is a
    // rollover table.  When rttMonStatsCaptureStartTimeIndex groups exceeds the 
    // rttMonStatisticsAdminNumHourGroups value, the oldest  corresponding hourly
    // group will be deleted and will be  replaced with the new
    // rttMonStatsCaptureStartTimeIndex hourly group.    All other indices will
    // fill to there maximum size.   The statistics capture table has five
    // indices.  Each described as follows:    -  The first index correlates its
    // entries to a       conceptual RTT control row via the       
    // rttMonCtrlAdminIndex object.   -  The second index is a rollover group and
    // it        uniquely identifies a 60 minute group. (The       
    // rttMonStatsCaptureStartTimeIndex object       is used to make this value
    // unique.)   -  When the RttMonRttType is 'pathEcho', the third        index
    // uniquely identifies the paths in a        statistics period.  (The period
    // is 60       minutes.)  A path will be created for each       unique path
    // through the network.  Note:  A       path that does not contain the target
    // is       considered a different path than one which       uses the exact
    // same path, but does contain the       target.  For all other values of
    // RttMonRttType       this index will be one.   -  When the RttMonRttType is
    // 'pathEcho', the fourth        index uniquely identifies the hops in each
    // path,        as grouped by the third index.  This index does        imply
    // the order of the hops along the path to a        target.  For all other
    // values of RttMonRttType       this index will be one.   -  The fifth index
    // uniquely creates a statistical       distribution bucket. The type is slice
    // of CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry.
    Rttmonstatscaptureentry []CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry
}

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetFilter() yfilter.YFilter { return rttmonstatscapturetable.YFilter }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) SetFilter(yf yfilter.YFilter) { rttmonstatscapturetable.YFilter = yf }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetGoName(yname string) string {
    if yname == "rttMonStatsCaptureEntry" { return "Rttmonstatscaptureentry" }
    return ""
}

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetSegmentPath() string {
    return "rttMonStatsCaptureTable"
}

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonStatsCaptureEntry" {
        for _, c := range rttmonstatscapturetable.Rttmonstatscaptureentry {
            if rttmonstatscapturetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry{}
        rttmonstatscapturetable.Rttmonstatscaptureentry = append(rttmonstatscapturetable.Rttmonstatscaptureentry, child)
        return &rttmonstatscapturetable.Rttmonstatscaptureentry[len(rttmonstatscapturetable.Rttmonstatscaptureentry)-1]
    }
    return nil
}

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonstatscapturetable.Rttmonstatscaptureentry {
        children[rttmonstatscapturetable.Rttmonstatscaptureentry[i].GetSegmentPath()] = &rttmonstatscapturetable.Rttmonstatscaptureentry[i]
    }
    return children
}

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetYangName() string { return "rttMonStatsCaptureTable" }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) SetParent(parent types.Entity) { rttmonstatscapturetable.parent = parent }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetParent() types.Entity { return rttmonstatscapturetable.parent }

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// The statistics capture table is a rollover table.  When
// rttMonStatsCaptureStartTimeIndex groups exceeds the 
// rttMonStatisticsAdminNumHourGroups value, the oldest 
// corresponding hourly group will be deleted and will be 
// replaced with the new rttMonStatsCaptureStartTimeIndex
// hourly group.  
// 
// All other indices will fill to there maximum size. 
// 
// The statistics capture table has five indices.  Each
// described as follows:
// 
//   -  The first index correlates its entries to a
//       conceptual RTT control row via the 
//       rttMonCtrlAdminIndex object.
//   -  The second index is a rollover group and it 
//       uniquely identifies a 60 minute group. (The 
//       rttMonStatsCaptureStartTimeIndex object
//       is used to make this value unique.)
//   -  When the RttMonRttType is 'pathEcho', the third 
//       index uniquely identifies the paths in a 
//       statistics period.  (The period is 60
//       minutes.)  A path will be created for each
//       unique path through the network.  Note:  A
//       path that does not contain the target is
//       considered a different path than one which
//       uses the exact same path, but does contain the
//       target.  For all other values of RttMonRttType
//       this index will be one.
//   -  When the RttMonRttType is 'pathEcho', the fourth 
//       index uniquely identifies the hops in each path, 
//       as grouped by the third index.  This index does 
//       imply the order of the hops along the path to a 
//       target.  For all other values of RttMonRttType
//       this index will be one.
//   -  The fifth index uniquely creates a statistical
//       distribution bucket.
type CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. The time when this row was created.  This object
    // is the second index of the  rttMonStatsCaptureTable Table.  The the number
    // of rttMonStatsCaptureStartTimeIndex   groups exceeds the
    // rttMonStatisticsAdminNumHourGroups value, the oldest
    // rttMonStatsCaptureStartTimeIndex  group will be removed and replaced with
    // the new entry.  When the RttMonRttType is 'pathEcho', this object also 
    // uniquely defines a group of paths.  See the  rttMonStatsCaptureEntry
    // object. The type is interface{} with range: 0..4294967295.
    Rttmonstatscapturestarttimeindex interface{}

    // This attribute is a key. When the RttMonRttType is 'pathEcho', this object
    // uniquely defines a path for a given value of 
    // rttMonStatsCaptureStartTimeIndex.  For all other values of RttMonRttType,
    // this object will be one.  For a particular value of 
    // rttMonStatsCaptureStartTimeIndex, the agent assigns the first instance of a
    // path a value of 1, then second  instance a value of 2, and so on.  The
    // sequence keeps  incrementing until the number of paths equals 
    // rttMonStatisticsAdminNumPaths value, then no new paths  are kept for the
    // current rttMonStatsCaptureStartTimeIndex  group.  NOTE: A source to target
    // rttMonStatsCapturePathIndex       path will be created for each       
    // rttMonStatsCaptureStartTimeIndex to hold all        errors that occur when
    // a specific path or        connection has not be setup.  This value directly
    // represents the path to a target. We can only support 128 paths. The type is
    // interface{} with range: 1..128.
    Rttmonstatscapturepathindex interface{}

    // This attribute is a key. When the RttMonRttType is 'pathEcho', this object
    // uniquely defines a hop for a given value of  rttMonStatsCapturePathIndex. 
    // For all other values of RttMonRttType, this object will be one.  For a
    // particular value of rttMonStatsCapturePathIndex, the agent assigns the
    // first instance of a hop a value of 1, then second instance a value of 2,
    // and so on.  The sequence keeps incrementing until the number of  hops
    // equals rttMonStatisticsAdminNumHops value, then no new hops are kept for
    // the current rttMonStatsCapturePathIndex.  This value directly represents a
    // hop along the path to a target, thus we can only support 30 hops.  This
    // value shows the order along the path to a target. The type is interface{}
    // with range: 1..30.
    Rttmonstatscapturehopindex interface{}

    // This attribute is a key. This object uniquely defines a statistical
    // distribution bucket for a given value of rttMonStatsCaptureHopIndex.  For a
    // particular value of rttMonStatsCaptureHopIndex, the agent assigns the first
    // instance of a distribution a value of 1, then second instance a value of 2,
    // and so on.  The sequence keeps incrementing until the number of  statistics
    // distribution intervals equals  rttMonStatisticsAdminNumDistBuckets value,
    // then all values that fall above the last interval will be placed into the
    // last interval.  Each of these Statistics Distribution Buckets contain  the
    // results of each completion as defined by  rttMonStatisticsAdminDistInterval
    // object. The type is interface{} with range: 1..20.
    Rttmonstatscapturedistindex interface{}

    // The number of RTT operations that have completed without an error and
    // without timing out.  This object has the special behavior as defined by the
    // ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB object. The type is
    // interface{} with range: 0..2147483647.
    Rttmonstatscapturecompletions interface{}

    // The number of RTT operations successfully completed, but in excess of
    // rttMonCtrlAdminThreshold.  This number is a subset of the accumulation of
    // all  rttMonStatsCaptureCompletions.  The operation time  of these completed
    // operations will be accumulated.  This object has the special behavior as
    // defined by the ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB
    // object. The type is interface{} with range: 0..2147483647.
    Rttmonstatscaptureoverthresholds interface{}

    // The accumulated completion time of RTT operations which complete
    // successfully. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Rttmonstatscapturesumcompletiontime interface{}

    // The low order 32 bits of the accumulated squares of completion times (in
    // milliseconds) of RTT  operations which complete successfully.  Low/High
    // order is defined where the binary number will look as follows:
    // ------------------------------------------------- | High order 32 bits    |
    // Low order 32 bits     | -------------------------------------------------
    // For example the number 4294967296 would have all Low order bits as '0' and
    // the rightmost High order bit will be 1 (zeros,1). The type is interface{}
    // with range: 0..4294967295.
    Rttmonstatscapturesumcompletiontime2Low interface{}

    // The high order 32 bits of the accumulated squares of completion times (in
    // milliseconds) of RTT  operations which complete successfully.  See the
    // rttMonStatsCaptureSumCompletionTime2Low object for a definition of Low/High
    // Order. The type is interface{} with range: 0..4294967295.
    Rttmonstatscapturesumcompletiontime2High interface{}

    // The maximum completion time of any RTT operation which completes
    // successfully. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Rttmonstatscapturecompletiontimemax interface{}

    // The minimum completion time of any RTT operation which completes
    // successfully. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Rttmonstatscapturecompletiontimemin interface{}
}

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetFilter() yfilter.YFilter { return rttmonstatscaptureentry.YFilter }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) SetFilter(yf yfilter.YFilter) { rttmonstatscaptureentry.YFilter = yf }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonStatsCaptureStartTimeIndex" { return "Rttmonstatscapturestarttimeindex" }
    if yname == "rttMonStatsCapturePathIndex" { return "Rttmonstatscapturepathindex" }
    if yname == "rttMonStatsCaptureHopIndex" { return "Rttmonstatscapturehopindex" }
    if yname == "rttMonStatsCaptureDistIndex" { return "Rttmonstatscapturedistindex" }
    if yname == "rttMonStatsCaptureCompletions" { return "Rttmonstatscapturecompletions" }
    if yname == "rttMonStatsCaptureOverThresholds" { return "Rttmonstatscaptureoverthresholds" }
    if yname == "rttMonStatsCaptureSumCompletionTime" { return "Rttmonstatscapturesumcompletiontime" }
    if yname == "rttMonStatsCaptureSumCompletionTime2Low" { return "Rttmonstatscapturesumcompletiontime2Low" }
    if yname == "rttMonStatsCaptureSumCompletionTime2High" { return "Rttmonstatscapturesumcompletiontime2High" }
    if yname == "rttMonStatsCaptureCompletionTimeMax" { return "Rttmonstatscapturecompletiontimemax" }
    if yname == "rttMonStatsCaptureCompletionTimeMin" { return "Rttmonstatscapturecompletiontimemin" }
    return ""
}

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetSegmentPath() string {
    return "rttMonStatsCaptureEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonctrladminindex) + "']" + "[rttMonStatsCaptureStartTimeIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturestarttimeindex) + "']" + "[rttMonStatsCapturePathIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturepathindex) + "']" + "[rttMonStatsCaptureHopIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturehopindex) + "']" + "[rttMonStatsCaptureDistIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturedistindex) + "']"
}

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonstatscaptureentry.Rttmonctrladminindex
    leafs["rttMonStatsCaptureStartTimeIndex"] = rttmonstatscaptureentry.Rttmonstatscapturestarttimeindex
    leafs["rttMonStatsCapturePathIndex"] = rttmonstatscaptureentry.Rttmonstatscapturepathindex
    leafs["rttMonStatsCaptureHopIndex"] = rttmonstatscaptureentry.Rttmonstatscapturehopindex
    leafs["rttMonStatsCaptureDistIndex"] = rttmonstatscaptureentry.Rttmonstatscapturedistindex
    leafs["rttMonStatsCaptureCompletions"] = rttmonstatscaptureentry.Rttmonstatscapturecompletions
    leafs["rttMonStatsCaptureOverThresholds"] = rttmonstatscaptureentry.Rttmonstatscaptureoverthresholds
    leafs["rttMonStatsCaptureSumCompletionTime"] = rttmonstatscaptureentry.Rttmonstatscapturesumcompletiontime
    leafs["rttMonStatsCaptureSumCompletionTime2Low"] = rttmonstatscaptureentry.Rttmonstatscapturesumcompletiontime2Low
    leafs["rttMonStatsCaptureSumCompletionTime2High"] = rttmonstatscaptureentry.Rttmonstatscapturesumcompletiontime2High
    leafs["rttMonStatsCaptureCompletionTimeMax"] = rttmonstatscaptureentry.Rttmonstatscapturecompletiontimemax
    leafs["rttMonStatsCaptureCompletionTimeMin"] = rttmonstatscaptureentry.Rttmonstatscapturecompletiontimemin
    return leafs
}

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetYangName() string { return "rttMonStatsCaptureEntry" }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) SetParent(parent types.Entity) { rttmonstatscaptureentry.parent = parent }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetParent() types.Entity { return rttmonstatscaptureentry.parent }

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetParentYangName() string { return "rttMonStatsCaptureTable" }

// CISCORTTMONMIB_Rttmonstatscollecttable
// The statistics collection database.
// 
// This table has the exact same behavior as the
// rttMonStatsCaptureTable, except it does not keep
// statistical distribution information.
// 
// For a complete table description see
// the rttMonStatsCaptureTable object.
type CISCORTTMONMIB_Rttmonstatscollecttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry has the exact same
    // behavior as the  rttMonStatsCaptureEntry, except it does not keep
    // statistical distribution information.  For a complete entry description see
    // the rttMonStatsCaptureEntry object. The type is slice of
    // CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry.
    Rttmonstatscollectentry []CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry
}

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetFilter() yfilter.YFilter { return rttmonstatscollecttable.YFilter }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) SetFilter(yf yfilter.YFilter) { rttmonstatscollecttable.YFilter = yf }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetGoName(yname string) string {
    if yname == "rttMonStatsCollectEntry" { return "Rttmonstatscollectentry" }
    return ""
}

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetSegmentPath() string {
    return "rttMonStatsCollectTable"
}

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonStatsCollectEntry" {
        for _, c := range rttmonstatscollecttable.Rttmonstatscollectentry {
            if rttmonstatscollecttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry{}
        rttmonstatscollecttable.Rttmonstatscollectentry = append(rttmonstatscollecttable.Rttmonstatscollectentry, child)
        return &rttmonstatscollecttable.Rttmonstatscollectentry[len(rttmonstatscollecttable.Rttmonstatscollectentry)-1]
    }
    return nil
}

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonstatscollecttable.Rttmonstatscollectentry {
        children[rttmonstatscollecttable.Rttmonstatscollectentry[i].GetSegmentPath()] = &rttmonstatscollecttable.Rttmonstatscollectentry[i]
    }
    return children
}

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetYangName() string { return "rttMonStatsCollectTable" }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) SetParent(parent types.Entity) { rttmonstatscollecttable.parent = parent }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetParent() types.Entity { return rttmonstatscollecttable.parent }

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry has the exact same behavior as the 
// rttMonStatsCaptureEntry, except it does not keep
// statistical distribution information.
// 
// For a complete entry description see
// the rttMonStatsCaptureEntry object.
type CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry_Rttmonstatscapturestarttimeindex
    Rttmonstatscapturestarttimeindex interface{}

    // This attribute is a key. The type is string with range: 1..128. Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry_Rttmonstatscapturepathindex
    Rttmonstatscapturepathindex interface{}

    // This attribute is a key. The type is string with range: 1..30. Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry_Rttmonstatscapturehopindex
    Rttmonstatscapturehopindex interface{}

    // When the RttMonRttType is 'echo' or pathEcho', this object represents the
    // number of times that the target or  hop along the path to a target became
    // disconnected.  For all other values of RttMonRttType, this object will
    // remain zero.  For connectionless protocols this has no meaning, and will
    // consequently remain 0.  When rttMonEchoAdminProtocol is one of snaRUEcho,
    // this is the number of times that an LU-SSCP session was lost,  for
    // snaLU0EchoAppl, snaLU2EchoAppl, snaLu62Echo, and for  snaLU62EchoAppl, this
    // is the number of times that LU-LU  session was lost.  Since this error does
    // not indicate any information about the failure of an RTT operation, no
    // response time  information for this instance will be recorded in the 
    // appropriate objects.  If this error occurs and the
    // rttMonStatsCapturePathIndex  cannot be determined, this error will be
    // accumulated in  the source to target path, that will always exist.  This
    // object has the special behavior as defined by the ROLLOVER NOTE in the
    // DESCRIPTION of the ciscoRttMonMIB object. The type is interface{} with
    // range: 0..2147483647.
    Rttmonstatscollectnumdisconnects interface{}

    // The number of occasions when a RTT operation was not completed before a
    // timeout occurred, i.e. rttMonCtrlAdminTimeout was exceeded.  Since the RTT
    // operation was never completed, the  completion time of these operations are
    // not accumulated, nor do they increment rttMonStatsCaptureCompletions (in 
    // any of the statistics distribution buckets).  This object has the special
    // behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    Rttmonstatscollecttimeouts interface{}

    // The number of occasions when a RTT operation could not be initiated because
    // a previous RTT operation has not  been completed.  When the RttMonRttType
    // is 'pathEcho' this can occur for both connection oriented protocols and
    // connectionless protocols.  When the RttMonRttType is 'echo' this can only
    // occur for connection oriented protocols such as SNA.   When the initiation
    // of a new operation cannot be started, this object will be incremented and
    // the operation will be omitted.  (The next operation will start at the next 
    // Frequency).  Since, a RTT operation was never initiated,  the completion
    // time of these operations is not  accumulated, nor do they increment 
    // rttMonStatsCaptureCompletions.  When the RttMonRttType is 'pathEcho', and
    // this error  occurs and the rttMonStatsCapturePathIndex cannot be 
    // determined, this error will be accumulated in the source  to target path,
    // that will always exist.  This object has the special behavior as defined by
    // the ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB object. The type
    // is interface{} with range: 0..2147483647.
    Rttmonstatscollectbusies interface{}

    // When the RttMonRttType is 'echo' or 'pathEcho' this is the number of
    // occasions when a RTT operation could not be initiated because the
    // connection to the target has not  been established.  For all other
    // RttMonRttTypes this object will remain zero.  This cannot occur for
    // connectionless protocols, but may occur for connection oriented protocols,
    // such as SNA.  Since a RTT operation was never initiated, the completion
    // time of these operations are not accumulated, nor do they increment
    // rttMonStatsCaptureCompletions.   If this error occurs and the
    // rttMonStatsCapturePathIndex cannot be determined, this error will be
    // accumulated in the source to target path, that will always exist.  This
    // object has the special behavior as defined by the ROLLOVER NOTE in the
    // DESCRIPTION of the ciscoRttMonMIB object. The type is interface{} with
    // range: 0..2147483647.
    Rttmonstatscollectnoconnections interface{}

    // The number of occasions when a RTT operation could not be initiated because
    // some necessary internal resource  (for example memory, or SNA subsystem)
    // was not available, or the operation completion could not be recognized. 
    // Since a RTT operation was never initiated or was not recognized, the
    // completion time of these operations  are not accumulated, nor do they
    // increment  rttMonStatsCaptureCompletions (in the expected  Distribution
    // Bucket).  When the RttMonRttType is 'pathEcho', and this error  occurs and
    // the rttMonStatsCapturePathIndex cannot be  determined, this error will be
    // accumulated in the  source to target path, that will always exist.  This
    // object has the special behavior as defined by the ROLLOVER NOTE in the
    // DESCRIPTION of the ciscoRttMonMIB object. The type is interface{} with
    // range: 0..2147483647.
    Rttmonstatscollectdrops interface{}

    // When the RttMonRttType is 'echo' of 'pathEcho' this is the number of RTT
    // operation completions received with  an unexpected sequence identifier. 
    // For all other values of RttMonRttType this object will remain zero.  When
    // this has occurred some of the possible reasons may be:      - a duplicate
    // packet was received    - a response was received after it had timed-out   
    // - a corrupted packet was received and was not detected  The completion time
    // of these operations are not  accumulated, nor do they increment 
    // rttMonStatsCaptureCompletions (in the expected Distribution Bucket).  This
    // object has the special behavior as defined by the ROLLOVER NOTE in the
    // DESCRIPTION of the ciscoRttMonMIB object. The type is interface{} with
    // range: 0..2147483647.
    Rttmonstatscollectsequenceerrors interface{}

    // The number of RTT operation completions received with data that does not
    // compare with the expected data.  The  completion time of these operations
    // are not accumulated,  nor do they increment rttMonStatsCaptureCompletions
    // (in the expected Distribution Bucket).  This object has the special
    // behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    Rttmonstatscollectverifyerrors interface{}

    // This object only applies when the RttMonRttType is 'echo', 'pathEcho',
    // 'dlsw', 'udpEcho', 'tcpConnect'.   For all other values of the
    // RttMonRttType, this will be  null.   The object is a string which specifies
    // the address of  the target for the this RTT operation.  This address will
    // be the address of the hop along the  path to the
    // rttMonEchoAdminTargetAddress address,  including
    // rttMonEchoAdminTargetAddress address, or just  the
    // rttMonEchoAdminTargetAddress address, when the  path information is not
    // collected.  This behavior is defined by the rttMonCtrlAdminRttType object. 
    // The interpretation of this string depends on the type  of RTT operation
    // selected, as specified by the  rttMonEchoAdminProtocol object. The type is
    // string.
    Rttmonstatscollectaddress interface{}

    // The number of occasions when control enable request failed. Currently it is
    // used for multicast operation type.  This object has the special behavior as
    // defined by the ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB
    // object. rttMonControlEnableErrors object is superseded by
    // rttMonStatsCollectCtrlEnErrors. The type is interface{} with range:
    // 0..2147483647.
    Rttmoncontrolenableerrors interface{}

    // The number of occasions when stats retrieval request failed. Currently it
    // is used for multicast operation type.  This object has the special behavior
    // as defined by the ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB
    // object. rttMonStatsRetrieveErrors object is superseded by
    // rttMonStatsCollectRetrieveErrors. The type is interface{} with range:
    // 0..2147483647.
    Rttmonstatsretrieveerrors interface{}

    // The object is same as rttMonControlEnableErrors, with corrected name for
    // consistency.  The number of occasions when control enable request failed.
    // Currently it is used for multicast operation type.  This object has the
    // special behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    Rttmonstatscollectctrlenerrors interface{}

    // The object is same as rttMonStatsRetrieveErrors, with corrected name for
    // consistency.  The number of occasions when stats retrieval request failed.
    // Currently it is used for multicast operation type.  This object has the
    // special behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    Rttmonstatscollectretrieveerrors interface{}
}

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetFilter() yfilter.YFilter { return rttmonstatscollectentry.YFilter }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) SetFilter(yf yfilter.YFilter) { rttmonstatscollectentry.YFilter = yf }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonStatsCaptureStartTimeIndex" { return "Rttmonstatscapturestarttimeindex" }
    if yname == "rttMonStatsCapturePathIndex" { return "Rttmonstatscapturepathindex" }
    if yname == "rttMonStatsCaptureHopIndex" { return "Rttmonstatscapturehopindex" }
    if yname == "rttMonStatsCollectNumDisconnects" { return "Rttmonstatscollectnumdisconnects" }
    if yname == "rttMonStatsCollectTimeouts" { return "Rttmonstatscollecttimeouts" }
    if yname == "rttMonStatsCollectBusies" { return "Rttmonstatscollectbusies" }
    if yname == "rttMonStatsCollectNoConnections" { return "Rttmonstatscollectnoconnections" }
    if yname == "rttMonStatsCollectDrops" { return "Rttmonstatscollectdrops" }
    if yname == "rttMonStatsCollectSequenceErrors" { return "Rttmonstatscollectsequenceerrors" }
    if yname == "rttMonStatsCollectVerifyErrors" { return "Rttmonstatscollectverifyerrors" }
    if yname == "rttMonStatsCollectAddress" { return "Rttmonstatscollectaddress" }
    if yname == "rttMonControlEnableErrors" { return "Rttmoncontrolenableerrors" }
    if yname == "rttMonStatsRetrieveErrors" { return "Rttmonstatsretrieveerrors" }
    if yname == "rttMonStatsCollectCtrlEnErrors" { return "Rttmonstatscollectctrlenerrors" }
    if yname == "rttMonStatsCollectRetrieveErrors" { return "Rttmonstatscollectretrieveerrors" }
    return ""
}

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetSegmentPath() string {
    return "rttMonStatsCollectEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonctrladminindex) + "']" + "[rttMonStatsCaptureStartTimeIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonstatscapturestarttimeindex) + "']" + "[rttMonStatsCapturePathIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonstatscapturepathindex) + "']" + "[rttMonStatsCaptureHopIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonstatscapturehopindex) + "']"
}

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonstatscollectentry.Rttmonctrladminindex
    leafs["rttMonStatsCaptureStartTimeIndex"] = rttmonstatscollectentry.Rttmonstatscapturestarttimeindex
    leafs["rttMonStatsCapturePathIndex"] = rttmonstatscollectentry.Rttmonstatscapturepathindex
    leafs["rttMonStatsCaptureHopIndex"] = rttmonstatscollectentry.Rttmonstatscapturehopindex
    leafs["rttMonStatsCollectNumDisconnects"] = rttmonstatscollectentry.Rttmonstatscollectnumdisconnects
    leafs["rttMonStatsCollectTimeouts"] = rttmonstatscollectentry.Rttmonstatscollecttimeouts
    leafs["rttMonStatsCollectBusies"] = rttmonstatscollectentry.Rttmonstatscollectbusies
    leafs["rttMonStatsCollectNoConnections"] = rttmonstatscollectentry.Rttmonstatscollectnoconnections
    leafs["rttMonStatsCollectDrops"] = rttmonstatscollectentry.Rttmonstatscollectdrops
    leafs["rttMonStatsCollectSequenceErrors"] = rttmonstatscollectentry.Rttmonstatscollectsequenceerrors
    leafs["rttMonStatsCollectVerifyErrors"] = rttmonstatscollectentry.Rttmonstatscollectverifyerrors
    leafs["rttMonStatsCollectAddress"] = rttmonstatscollectentry.Rttmonstatscollectaddress
    leafs["rttMonControlEnableErrors"] = rttmonstatscollectentry.Rttmoncontrolenableerrors
    leafs["rttMonStatsRetrieveErrors"] = rttmonstatscollectentry.Rttmonstatsretrieveerrors
    leafs["rttMonStatsCollectCtrlEnErrors"] = rttmonstatscollectentry.Rttmonstatscollectctrlenerrors
    leafs["rttMonStatsCollectRetrieveErrors"] = rttmonstatscollectentry.Rttmonstatscollectretrieveerrors
    return leafs
}

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetYangName() string { return "rttMonStatsCollectEntry" }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) SetParent(parent types.Entity) { rttmonstatscollectentry.parent = parent }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetParent() types.Entity { return rttmonstatscollectentry.parent }

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetParentYangName() string { return "rttMonStatsCollectTable" }

// CISCORTTMONMIB_Rttmonstatstotalstable
// The statistics totals database.
// 
// This table has the exact same behavior as the
// rttMonStatsCaptureTable, except it only keeps
// 60 minute group values.
// 
// For a complete table description see
// the rttMonStatsCaptureTable object.
type CISCORTTMONMIB_Rttmonstatstotalstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry has the exact same
    // behavior as the  rttMonStatsCaptureEntry, except it only keeps 60 minute
    // group values.  For a complete entry description see the
    // rttMonStatsCaptureEntry object. The type is slice of
    // CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry.
    Rttmonstatstotalsentry []CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry
}

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetFilter() yfilter.YFilter { return rttmonstatstotalstable.YFilter }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) SetFilter(yf yfilter.YFilter) { rttmonstatstotalstable.YFilter = yf }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetGoName(yname string) string {
    if yname == "rttMonStatsTotalsEntry" { return "Rttmonstatstotalsentry" }
    return ""
}

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetSegmentPath() string {
    return "rttMonStatsTotalsTable"
}

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonStatsTotalsEntry" {
        for _, c := range rttmonstatstotalstable.Rttmonstatstotalsentry {
            if rttmonstatstotalstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry{}
        rttmonstatstotalstable.Rttmonstatstotalsentry = append(rttmonstatstotalstable.Rttmonstatstotalsentry, child)
        return &rttmonstatstotalstable.Rttmonstatstotalsentry[len(rttmonstatstotalstable.Rttmonstatstotalsentry)-1]
    }
    return nil
}

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonstatstotalstable.Rttmonstatstotalsentry {
        children[rttmonstatstotalstable.Rttmonstatstotalsentry[i].GetSegmentPath()] = &rttmonstatstotalstable.Rttmonstatstotalsentry[i]
    }
    return children
}

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetYangName() string { return "rttMonStatsTotalsTable" }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) SetParent(parent types.Entity) { rttmonstatstotalstable.parent = parent }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetParent() types.Entity { return rttmonstatstotalstable.parent }

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry has the exact same behavior as the 
// rttMonStatsCaptureEntry, except it only keeps
// 60 minute group values.
// 
// For a complete entry description see
// the rttMonStatsCaptureEntry object.
type CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry_Rttmonstatscapturestarttimeindex
    Rttmonstatscapturestarttimeindex interface{}

    // The length of time since this conceptual statistics row was created. The
    // type is interface{} with range: 0..2147483647.
    Rttmonstatstotalselapsedtime interface{}

    // The number of RTT operations that have been initiated.  This number
    // includes all RTT operations which succeed  or fail for whatever reason. 
    // This object has the special behavior as defined by the ROLLOVER NOTE in the
    // DESCRIPTION of the ciscoRttMonMIB object. The type is interface{} with
    // range: 0..2147483647.
    Rttmonstatstotalsinitiations interface{}
}

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetFilter() yfilter.YFilter { return rttmonstatstotalsentry.YFilter }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) SetFilter(yf yfilter.YFilter) { rttmonstatstotalsentry.YFilter = yf }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonStatsCaptureStartTimeIndex" { return "Rttmonstatscapturestarttimeindex" }
    if yname == "rttMonStatsTotalsElapsedTime" { return "Rttmonstatstotalselapsedtime" }
    if yname == "rttMonStatsTotalsInitiations" { return "Rttmonstatstotalsinitiations" }
    return ""
}

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetSegmentPath() string {
    return "rttMonStatsTotalsEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonstatstotalsentry.Rttmonctrladminindex) + "']" + "[rttMonStatsCaptureStartTimeIndex='" + fmt.Sprintf("%v", rttmonstatstotalsentry.Rttmonstatscapturestarttimeindex) + "']"
}

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonstatstotalsentry.Rttmonctrladminindex
    leafs["rttMonStatsCaptureStartTimeIndex"] = rttmonstatstotalsentry.Rttmonstatscapturestarttimeindex
    leafs["rttMonStatsTotalsElapsedTime"] = rttmonstatstotalsentry.Rttmonstatstotalselapsedtime
    leafs["rttMonStatsTotalsInitiations"] = rttmonstatstotalsentry.Rttmonstatstotalsinitiations
    return leafs
}

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetYangName() string { return "rttMonStatsTotalsEntry" }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) SetParent(parent types.Entity) { rttmonstatstotalsentry.parent = parent }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetParent() types.Entity { return rttmonstatstotalsentry.parent }

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetParentYangName() string { return "rttMonStatsTotalsTable" }

// CISCORTTMONMIB_Rttmonhttpstatstable
// The HTTP statistics collection database.
// 
// The HTTP statistics table contains summarized information of
// the results for a conceptual RTT control row. A rolling
// accumulated history of this information is maintained in a 
// series of hourly 'group(s)'.
// 
// The operation of this table is same as that of 
// rttMonStatsCaptureTable, except that this table can only 
// store a maximum of 2 hours of data.
type CISCORTTMONMIB_Rttmonhttpstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry is created only if the
    // rttMonCtrlAdminRttType  is http. The operation of this table is same as
    // that of rttMonStatsCaptureTable. The type is slice of
    // CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry.
    Rttmonhttpstatsentry []CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry
}

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetFilter() yfilter.YFilter { return rttmonhttpstatstable.YFilter }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) SetFilter(yf yfilter.YFilter) { rttmonhttpstatstable.YFilter = yf }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetGoName(yname string) string {
    if yname == "rttMonHTTPStatsEntry" { return "Rttmonhttpstatsentry" }
    return ""
}

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetSegmentPath() string {
    return "rttMonHTTPStatsTable"
}

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonHTTPStatsEntry" {
        for _, c := range rttmonhttpstatstable.Rttmonhttpstatsentry {
            if rttmonhttpstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry{}
        rttmonhttpstatstable.Rttmonhttpstatsentry = append(rttmonhttpstatstable.Rttmonhttpstatsentry, child)
        return &rttmonhttpstatstable.Rttmonhttpstatsentry[len(rttmonhttpstatstable.Rttmonhttpstatsentry)-1]
    }
    return nil
}

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonhttpstatstable.Rttmonhttpstatsentry {
        children[rttmonhttpstatstable.Rttmonhttpstatsentry[i].GetSegmentPath()] = &rttmonhttpstatstable.Rttmonhttpstatsentry[i]
    }
    return children
}

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetYangName() string { return "rttMonHTTPStatsTable" }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) SetParent(parent types.Entity) { rttmonhttpstatstable.parent = parent }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetParent() types.Entity { return rttmonhttpstatstable.parent }

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry is created only if the rttMonCtrlAdminRttType 
// is http. The operation of this table is same as that of
// rttMonStatsCaptureTable.
type CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. This is the time when this row was created. This
    // index uniquely identifies a HTTP Stats row in the  rttMonHTTPStatsTable.
    // The type is interface{} with range: 0..4294967295.
    Rttmonhttpstatsstarttimeindex interface{}

    // The number of HTTP operations that have completed successfully. The type is
    // interface{} with range: 0..4294967295.
    Rttmonhttpstatscompletions interface{}

    // The number of HTTP operations that violate threshold. The type is
    // interface{} with range: 0..4294967295.
    Rttmonhttpstatsoverthresholds interface{}

    // The sum of HTTP operations that are successfully measured. The type is
    // interface{} with range: 0..4294967295.
    Rttmonhttpstatsrttsum interface{}

    // The sum of squares of the RTT's that are successfully measured (low order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    Rttmonhttpstatsrttsum2Low interface{}

    // The sum of squares of the RTT's that are successfully measured (high order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    Rttmonhttpstatsrttsum2High interface{}

    // The minimum RTT taken to perform HTTP operation. The type is interface{}
    // with range: 0..4294967295.
    Rttmonhttpstatsrttmin interface{}

    // The maximum RTT taken to perform HTTP operation. The type is interface{}
    // with range: 0..4294967295. Units are milliseconds.
    Rttmonhttpstatsrttmax interface{}

    // The sum of RTT taken to perform DNS query within the HTTP operation. The
    // type is interface{} with range: 0..4294967295.
    Rttmonhttpstatsdnsrttsum interface{}

    // The sum of RTT taken to connect to the HTTP server. The type is interface{}
    // with range: 0..4294967295.
    Rttmonhttpstatstcpconnectrttsum interface{}

    // The sum of RTT taken to download the object specified by URL. The type is
    // interface{} with range: 0..4294967295.
    Rttmonhttpstatstransactionrttsum interface{}

    // The sum of the size of the message body received as a response to the HTTP
    // request. The type is interface{} with range: 0..4294967295.
    Rttmonhttpstatsmessagebodyoctetssum interface{}

    // The number of requests that could not connect to the DNS Server. The type
    // is interface{} with range: 0..4294967295.
    Rttmonhttpstatsdnsservertimeout interface{}

    // The number of requests that could not connect to the the HTTP Server. The
    // type is interface{} with range: 0..4294967295.
    Rttmonhttpstatstcpconnecttimeout interface{}

    // The number of requests that timed out during HTTP transaction. The type is
    // interface{} with range: 0..4294967295.
    Rttmonhttpstatstransactiontimeout interface{}

    // The number of requests that had DNS Query errors. The type is interface{}
    // with range: 0..4294967295.
    Rttmonhttpstatsdnsqueryerror interface{}

    // The number of requests that had HTTP errors while downloading the base
    // page. The type is interface{} with range: 0..4294967295.
    Rttmonhttpstatshttperror interface{}

    // The number of occasions when a HTTP operation could not be initiated
    // because an internal error. The type is interface{} with range:
    // 0..4294967295.
    Rttmonhttpstatserror interface{}

    // The number of occasions when an HTTP operation could not be initiated
    // because a previous HTTP operation has not been completed. The type is
    // interface{} with range: 0..4294967295.
    Rttmonhttpstatsbusies interface{}
}

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetFilter() yfilter.YFilter { return rttmonhttpstatsentry.YFilter }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) SetFilter(yf yfilter.YFilter) { rttmonhttpstatsentry.YFilter = yf }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonHTTPStatsStartTimeIndex" { return "Rttmonhttpstatsstarttimeindex" }
    if yname == "rttMonHTTPStatsCompletions" { return "Rttmonhttpstatscompletions" }
    if yname == "rttMonHTTPStatsOverThresholds" { return "Rttmonhttpstatsoverthresholds" }
    if yname == "rttMonHTTPStatsRTTSum" { return "Rttmonhttpstatsrttsum" }
    if yname == "rttMonHTTPStatsRTTSum2Low" { return "Rttmonhttpstatsrttsum2Low" }
    if yname == "rttMonHTTPStatsRTTSum2High" { return "Rttmonhttpstatsrttsum2High" }
    if yname == "rttMonHTTPStatsRTTMin" { return "Rttmonhttpstatsrttmin" }
    if yname == "rttMonHTTPStatsRTTMax" { return "Rttmonhttpstatsrttmax" }
    if yname == "rttMonHTTPStatsDNSRTTSum" { return "Rttmonhttpstatsdnsrttsum" }
    if yname == "rttMonHTTPStatsTCPConnectRTTSum" { return "Rttmonhttpstatstcpconnectrttsum" }
    if yname == "rttMonHTTPStatsTransactionRTTSum" { return "Rttmonhttpstatstransactionrttsum" }
    if yname == "rttMonHTTPStatsMessageBodyOctetsSum" { return "Rttmonhttpstatsmessagebodyoctetssum" }
    if yname == "rttMonHTTPStatsDNSServerTimeout" { return "Rttmonhttpstatsdnsservertimeout" }
    if yname == "rttMonHTTPStatsTCPConnectTimeout" { return "Rttmonhttpstatstcpconnecttimeout" }
    if yname == "rttMonHTTPStatsTransactionTimeout" { return "Rttmonhttpstatstransactiontimeout" }
    if yname == "rttMonHTTPStatsDNSQueryError" { return "Rttmonhttpstatsdnsqueryerror" }
    if yname == "rttMonHTTPStatsHTTPError" { return "Rttmonhttpstatshttperror" }
    if yname == "rttMonHTTPStatsError" { return "Rttmonhttpstatserror" }
    if yname == "rttMonHTTPStatsBusies" { return "Rttmonhttpstatsbusies" }
    return ""
}

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetSegmentPath() string {
    return "rttMonHTTPStatsEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonhttpstatsentry.Rttmonctrladminindex) + "']" + "[rttMonHTTPStatsStartTimeIndex='" + fmt.Sprintf("%v", rttmonhttpstatsentry.Rttmonhttpstatsstarttimeindex) + "']"
}

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonhttpstatsentry.Rttmonctrladminindex
    leafs["rttMonHTTPStatsStartTimeIndex"] = rttmonhttpstatsentry.Rttmonhttpstatsstarttimeindex
    leafs["rttMonHTTPStatsCompletions"] = rttmonhttpstatsentry.Rttmonhttpstatscompletions
    leafs["rttMonHTTPStatsOverThresholds"] = rttmonhttpstatsentry.Rttmonhttpstatsoverthresholds
    leafs["rttMonHTTPStatsRTTSum"] = rttmonhttpstatsentry.Rttmonhttpstatsrttsum
    leafs["rttMonHTTPStatsRTTSum2Low"] = rttmonhttpstatsentry.Rttmonhttpstatsrttsum2Low
    leafs["rttMonHTTPStatsRTTSum2High"] = rttmonhttpstatsentry.Rttmonhttpstatsrttsum2High
    leafs["rttMonHTTPStatsRTTMin"] = rttmonhttpstatsentry.Rttmonhttpstatsrttmin
    leafs["rttMonHTTPStatsRTTMax"] = rttmonhttpstatsentry.Rttmonhttpstatsrttmax
    leafs["rttMonHTTPStatsDNSRTTSum"] = rttmonhttpstatsentry.Rttmonhttpstatsdnsrttsum
    leafs["rttMonHTTPStatsTCPConnectRTTSum"] = rttmonhttpstatsentry.Rttmonhttpstatstcpconnectrttsum
    leafs["rttMonHTTPStatsTransactionRTTSum"] = rttmonhttpstatsentry.Rttmonhttpstatstransactionrttsum
    leafs["rttMonHTTPStatsMessageBodyOctetsSum"] = rttmonhttpstatsentry.Rttmonhttpstatsmessagebodyoctetssum
    leafs["rttMonHTTPStatsDNSServerTimeout"] = rttmonhttpstatsentry.Rttmonhttpstatsdnsservertimeout
    leafs["rttMonHTTPStatsTCPConnectTimeout"] = rttmonhttpstatsentry.Rttmonhttpstatstcpconnecttimeout
    leafs["rttMonHTTPStatsTransactionTimeout"] = rttmonhttpstatsentry.Rttmonhttpstatstransactiontimeout
    leafs["rttMonHTTPStatsDNSQueryError"] = rttmonhttpstatsentry.Rttmonhttpstatsdnsqueryerror
    leafs["rttMonHTTPStatsHTTPError"] = rttmonhttpstatsentry.Rttmonhttpstatshttperror
    leafs["rttMonHTTPStatsError"] = rttmonhttpstatsentry.Rttmonhttpstatserror
    leafs["rttMonHTTPStatsBusies"] = rttmonhttpstatsentry.Rttmonhttpstatsbusies
    return leafs
}

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetYangName() string { return "rttMonHTTPStatsEntry" }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) SetParent(parent types.Entity) { rttmonhttpstatsentry.parent = parent }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetParent() types.Entity { return rttmonhttpstatsentry.parent }

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetParentYangName() string { return "rttMonHTTPStatsTable" }

// CISCORTTMONMIB_Rttmonjitterstatstable
// The Jitter statistics collection database.
// 
// The Jitter statistics table contains summarized information of
// the results for a conceptual RTT control row. A rolling
// accumulated history of this information is maintained in a 
// series of hourly 'group(s)'.
// 
// The operation of this table is same as that of 
// rttMonStatsCaptureTable, except that this table will store 
// 2 hours of data.
type CISCORTTMONMIB_Rttmonjitterstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry is created only if the
    // rttMonCtrlAdminRttType  is jitter. The operation of this table is same as
    // that of rttMonStatsCaptureTable. The type is slice of
    // CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry.
    Rttmonjitterstatsentry []CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry
}

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetFilter() yfilter.YFilter { return rttmonjitterstatstable.YFilter }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) SetFilter(yf yfilter.YFilter) { rttmonjitterstatstable.YFilter = yf }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetGoName(yname string) string {
    if yname == "rttMonJitterStatsEntry" { return "Rttmonjitterstatsentry" }
    return ""
}

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetSegmentPath() string {
    return "rttMonJitterStatsTable"
}

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonJitterStatsEntry" {
        for _, c := range rttmonjitterstatstable.Rttmonjitterstatsentry {
            if rttmonjitterstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry{}
        rttmonjitterstatstable.Rttmonjitterstatsentry = append(rttmonjitterstatstable.Rttmonjitterstatsentry, child)
        return &rttmonjitterstatstable.Rttmonjitterstatsentry[len(rttmonjitterstatstable.Rttmonjitterstatsentry)-1]
    }
    return nil
}

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonjitterstatstable.Rttmonjitterstatsentry {
        children[rttmonjitterstatstable.Rttmonjitterstatsentry[i].GetSegmentPath()] = &rttmonjitterstatstable.Rttmonjitterstatsentry[i]
    }
    return children
}

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetYangName() string { return "rttMonJitterStatsTable" }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) SetParent(parent types.Entity) { rttmonjitterstatstable.parent = parent }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetParent() types.Entity { return rttmonjitterstatstable.parent }

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry is created only if the rttMonCtrlAdminRttType 
// is jitter. The operation of this table is same as that of
// rttMonStatsCaptureTable.
type CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. The time when this row was created. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsstarttimeindex interface{}

    // The number of jitter operation that have completed successfully. The type
    // is interface{} with range: 0..4294967295.
    Rttmonjitterstatscompletions interface{}

    // The number of jitter operations that violate threshold. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsoverthresholds interface{}

    // The number of RTT's that are successfully measured. The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatsnumofrtt interface{}

    // The sum of RTT's that are successfully measured (low order 32 bits). The
    // high order 32 bits are stored in rttMonJitterStatsRTTSumHigh. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsrttsum interface{}

    // The sum of squares of RTT's that are successfully measured (low order 32
    // bits). The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsrttsum2Low interface{}

    // The sum of squares of RTT's that are successfully measured (high order 32
    // bits). The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsrttsum2High interface{}

    // The minimum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsrttmin interface{}

    // The maximum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsrttmax interface{}

    // The minimum of absolute values of all positive jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsminofpositivessd interface{}

    // The maximum of absolute values of all positive jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsmaxofpositivessd interface{}

    // The sum of number of all positive jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsnumofpositivessd interface{}

    // The sum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatssumofpositivessd interface{}

    // The sum of square of RTT's of all positive jitter values from packets sent
    // from source to destination (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Positivessdlow interface{}

    // The sum of square of RTT's of all positive jitter values from packets sent
    // from source to destination (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Positivessdhigh interface{}

    // The minimum of all negative jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsminofnegativessd interface{}

    // The maximum of all negative jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsmaxofnegativessd interface{}

    // The sum of number of all negative jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsnumofnegativessd interface{}

    // The sum of RTT's of all negative jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatssumofnegativessd interface{}

    // The sum of square of RTT's of all negative jitter values from packets sent
    // from source to destination (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Negativessdlow interface{}

    // The sum of square of RTT's of all negative jitter values from packets sent
    // from source to destination (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Negativessdhigh interface{}

    // The minimum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsminofpositivesds interface{}

    // The maximum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsmaxofpositivesds interface{}

    // The sum of number of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsnumofpositivesds interface{}

    // The sum of RTT's of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatssumofpositivesds interface{}

    // The sum of squares of RTT's of all positive jitter values from packets sent
    // from destination to source (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Positivesdslow interface{}

    // The sum of squares of RTT's of all positive jitter values from packets sent
    // from destination to source (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Positivesdshigh interface{}

    // The minimum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsminofnegativesds interface{}

    // The maximum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsmaxofnegativesds interface{}

    // The sum of number of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsnumofnegativesds interface{}

    // The sum of RTT's of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatssumofnegativesds interface{}

    // The sum of squares of RTT's of all negative jitter values from packets sent
    // from destination to source (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Negativesdslow interface{}

    // The sum of squares of RTT's of all negative jitter values from packets sent
    // from destination to source (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    Rttmonjitterstatssum2Negativesdshigh interface{}

    // The number of packets lost when sent from source to destination. The type
    // is interface{} with range: 0..4294967295.
    Rttmonjitterstatspacketlosssd interface{}

    // The number of packets lost when sent from destination to source. The type
    // is interface{} with range: 0..4294967295.
    Rttmonjitterstatspacketlossds interface{}

    // The number of packets arrived out of sequence. The type is interface{} with
    // range: 0..4294967295.
    Rttmonjitterstatspacketoutofsequence interface{}

    // The number of packets that are lost for which we cannot determine the
    // direction. The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatspacketmia interface{}

    // The number of packets that arrived after the timeout. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatspacketlatearrival interface{}

    // The number of occasions when a jitter operation could not be initiated
    // because an internal error. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatserror interface{}

    // The number of occasions when a jitter operation could not be initiated
    // because a previous jitter operation has not been completed. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsbusies interface{}

    // The sum of one way times from source to destination (low order 32 bits).
    // The high order 32 bits are stored in rttMonJitterStatsOWSumSDHigh. The type
    // is interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsumsd interface{}

    // The sum of squares of one way times from source to destination (low order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsum2Sdlow interface{}

    // The sum of squares of one way times from source to destination (high order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsum2Sdhigh interface{}

    // The minimum of all one way times from source to destination.
    // rttMonJitterStatsOWMinSD object is superseded by
    // rttMonJitterStatsOWMinSDNew. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowminsd interface{}

    // The maximum of all one way times from source to destination.
    // rttMonJitterStatsOWMaxSD object is superseded by
    // rttMonJitterStatsOWMaxSDNew. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowmaxsd interface{}

    // The sum of one way times from destination to source (low order 32 bits).
    // The high order 32 bits are stored in rttMonJitterStatsOWSumDSHigh. The type
    // is interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsumds interface{}

    // The sum of squares of one way times from destination to source (low order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsum2Dslow interface{}

    // The sum of squares of one way times from destination to source (high order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsum2Dshigh interface{}

    // The minimum of all one way times from destination to source.
    // rttMonJitterStatsOWMinDS object is superseded by
    // rttMonJitterStatsOWMinDSNew. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowminds interface{}

    // The maximum of all one way times from destination to source.
    // rttMonJitterStatsOWMaxDS object is superseded by
    // rttMonJitterStatsOWMaxDSNew. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowmaxds interface{}

    // The number of one way times that are successfully measured. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsnumofow interface{}

    // The minimum of all one way times from source to destination. Replaces
    // deprecated rttMonJitterStatsOWMinSD. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowminsdnew interface{}

    // The maximum of all one way times from source to destination. Replaces
    // deprecated rttMonJitterStatsOWMaxSD. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowmaxsdnew interface{}

    // The minimum of all one way times from destination to source. Replaces
    // deprecated rttMonJitterStatsOWMinDS. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowmindsnew interface{}

    // The maximum of all one way times from destination to source. Replaces
    // deprecated rttMonJitterStatsOWMaxDS. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsowmaxdsnew interface{}

    // The minimum of all MOS values for the jitter operations in hundreds.  This
    // value will be 0 if   - rttMonEchoAdminCodecType of the operation is
    // notApplicable   - the operation is not started   - the operation is started
    // but failed This value will be 1 for packet loss of 10% or more. The type is
    // interface{} with range: 0..None | 100..500.
    Rttmonjitterstatsminofmos interface{}

    // The maximum of all MOS values for the jitter operations in hunderds.  This
    // value will be 0 if   - rttMonEchoAdminCodecType of the operation is
    // notApplicable   - the operation is not started   - the operation is started
    // but failed This value will be 1 for packet loss of 10% or more. The type is
    // interface{} with range: 0..None | 100..500.
    Rttmonjitterstatsmaxofmos interface{}

    // The minimum of all ICPIF values for the jitter operations.  This value will
    // be 93 for packet loss of 10% or more. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsminoficpif interface{}

    // The maximum of all ICPIF values for the jitter operations.  This value will
    // be 93 for packet loss of 10% or more. The type is interface{} with range:
    // 0..4294967295.
    Rttmonjitterstatsmaxoficpif interface{}

    // Interarrival Jitter (RFC 1889) at responder. The type is interface{} with
    // range: 0..4294967295.
    Rttmonjitterstatsiajout interface{}

    // Interarrival Jitter (RFC 1889) at sender. The type is interface{} with
    // range: 0..4294967295.
    Rttmonjitterstatsiajin interface{}

    // The average of positive and negative jitter values for SD and DS direction.
    // The type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsavgjitter interface{}

    // The average of positive and negative jitter values in SD direction. The
    // type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsavgjittersd interface{}

    // The average of positive and negative jitter values in DS direction. The
    // type is interface{} with range: 0..4294967295.
    Rttmonjitterstatsavgjitterds interface{}

    // The number of RTT operations that have completed with sender and responder
    // out of sync with NTP. The NTP sync means  the total of NTP offset on sender
    // and responder is within  configured tolerance level. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsunsyncrts interface{}

    // The sum of RTT's that are successfully measured (high order 32 bits). The
    // low order 32 bits are  stored in rttMonJitterStatsRTTSum. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsrttsumhigh interface{}

    // The sum of one way times from source to destination (high order 32 bits).
    // The low order 32 bits are  stored in rttMonJitterStatsOWSumSD. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsumsdhigh interface{}

    // The sum of one way times from destination to source (high order 32 bits).
    // The low order 32 bits are stored in rttMonJitterStatsOWSumDS. The type is
    // interface{} with range: 0..4294967295.
    Rttmonjitterstatsowsumdshigh interface{}
}

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetFilter() yfilter.YFilter { return rttmonjitterstatsentry.YFilter }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) SetFilter(yf yfilter.YFilter) { rttmonjitterstatsentry.YFilter = yf }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonJitterStatsStartTimeIndex" { return "Rttmonjitterstatsstarttimeindex" }
    if yname == "rttMonJitterStatsCompletions" { return "Rttmonjitterstatscompletions" }
    if yname == "rttMonJitterStatsOverThresholds" { return "Rttmonjitterstatsoverthresholds" }
    if yname == "rttMonJitterStatsNumOfRTT" { return "Rttmonjitterstatsnumofrtt" }
    if yname == "rttMonJitterStatsRTTSum" { return "Rttmonjitterstatsrttsum" }
    if yname == "rttMonJitterStatsRTTSum2Low" { return "Rttmonjitterstatsrttsum2Low" }
    if yname == "rttMonJitterStatsRTTSum2High" { return "Rttmonjitterstatsrttsum2High" }
    if yname == "rttMonJitterStatsRTTMin" { return "Rttmonjitterstatsrttmin" }
    if yname == "rttMonJitterStatsRTTMax" { return "Rttmonjitterstatsrttmax" }
    if yname == "rttMonJitterStatsMinOfPositivesSD" { return "Rttmonjitterstatsminofpositivessd" }
    if yname == "rttMonJitterStatsMaxOfPositivesSD" { return "Rttmonjitterstatsmaxofpositivessd" }
    if yname == "rttMonJitterStatsNumOfPositivesSD" { return "Rttmonjitterstatsnumofpositivessd" }
    if yname == "rttMonJitterStatsSumOfPositivesSD" { return "Rttmonjitterstatssumofpositivessd" }
    if yname == "rttMonJitterStatsSum2PositivesSDLow" { return "Rttmonjitterstatssum2Positivessdlow" }
    if yname == "rttMonJitterStatsSum2PositivesSDHigh" { return "Rttmonjitterstatssum2Positivessdhigh" }
    if yname == "rttMonJitterStatsMinOfNegativesSD" { return "Rttmonjitterstatsminofnegativessd" }
    if yname == "rttMonJitterStatsMaxOfNegativesSD" { return "Rttmonjitterstatsmaxofnegativessd" }
    if yname == "rttMonJitterStatsNumOfNegativesSD" { return "Rttmonjitterstatsnumofnegativessd" }
    if yname == "rttMonJitterStatsSumOfNegativesSD" { return "Rttmonjitterstatssumofnegativessd" }
    if yname == "rttMonJitterStatsSum2NegativesSDLow" { return "Rttmonjitterstatssum2Negativessdlow" }
    if yname == "rttMonJitterStatsSum2NegativesSDHigh" { return "Rttmonjitterstatssum2Negativessdhigh" }
    if yname == "rttMonJitterStatsMinOfPositivesDS" { return "Rttmonjitterstatsminofpositivesds" }
    if yname == "rttMonJitterStatsMaxOfPositivesDS" { return "Rttmonjitterstatsmaxofpositivesds" }
    if yname == "rttMonJitterStatsNumOfPositivesDS" { return "Rttmonjitterstatsnumofpositivesds" }
    if yname == "rttMonJitterStatsSumOfPositivesDS" { return "Rttmonjitterstatssumofpositivesds" }
    if yname == "rttMonJitterStatsSum2PositivesDSLow" { return "Rttmonjitterstatssum2Positivesdslow" }
    if yname == "rttMonJitterStatsSum2PositivesDSHigh" { return "Rttmonjitterstatssum2Positivesdshigh" }
    if yname == "rttMonJitterStatsMinOfNegativesDS" { return "Rttmonjitterstatsminofnegativesds" }
    if yname == "rttMonJitterStatsMaxOfNegativesDS" { return "Rttmonjitterstatsmaxofnegativesds" }
    if yname == "rttMonJitterStatsNumOfNegativesDS" { return "Rttmonjitterstatsnumofnegativesds" }
    if yname == "rttMonJitterStatsSumOfNegativesDS" { return "Rttmonjitterstatssumofnegativesds" }
    if yname == "rttMonJitterStatsSum2NegativesDSLow" { return "Rttmonjitterstatssum2Negativesdslow" }
    if yname == "rttMonJitterStatsSum2NegativesDSHigh" { return "Rttmonjitterstatssum2Negativesdshigh" }
    if yname == "rttMonJitterStatsPacketLossSD" { return "Rttmonjitterstatspacketlosssd" }
    if yname == "rttMonJitterStatsPacketLossDS" { return "Rttmonjitterstatspacketlossds" }
    if yname == "rttMonJitterStatsPacketOutOfSequence" { return "Rttmonjitterstatspacketoutofsequence" }
    if yname == "rttMonJitterStatsPacketMIA" { return "Rttmonjitterstatspacketmia" }
    if yname == "rttMonJitterStatsPacketLateArrival" { return "Rttmonjitterstatspacketlatearrival" }
    if yname == "rttMonJitterStatsError" { return "Rttmonjitterstatserror" }
    if yname == "rttMonJitterStatsBusies" { return "Rttmonjitterstatsbusies" }
    if yname == "rttMonJitterStatsOWSumSD" { return "Rttmonjitterstatsowsumsd" }
    if yname == "rttMonJitterStatsOWSum2SDLow" { return "Rttmonjitterstatsowsum2Sdlow" }
    if yname == "rttMonJitterStatsOWSum2SDHigh" { return "Rttmonjitterstatsowsum2Sdhigh" }
    if yname == "rttMonJitterStatsOWMinSD" { return "Rttmonjitterstatsowminsd" }
    if yname == "rttMonJitterStatsOWMaxSD" { return "Rttmonjitterstatsowmaxsd" }
    if yname == "rttMonJitterStatsOWSumDS" { return "Rttmonjitterstatsowsumds" }
    if yname == "rttMonJitterStatsOWSum2DSLow" { return "Rttmonjitterstatsowsum2Dslow" }
    if yname == "rttMonJitterStatsOWSum2DSHigh" { return "Rttmonjitterstatsowsum2Dshigh" }
    if yname == "rttMonJitterStatsOWMinDS" { return "Rttmonjitterstatsowminds" }
    if yname == "rttMonJitterStatsOWMaxDS" { return "Rttmonjitterstatsowmaxds" }
    if yname == "rttMonJitterStatsNumOfOW" { return "Rttmonjitterstatsnumofow" }
    if yname == "rttMonJitterStatsOWMinSDNew" { return "Rttmonjitterstatsowminsdnew" }
    if yname == "rttMonJitterStatsOWMaxSDNew" { return "Rttmonjitterstatsowmaxsdnew" }
    if yname == "rttMonJitterStatsOWMinDSNew" { return "Rttmonjitterstatsowmindsnew" }
    if yname == "rttMonJitterStatsOWMaxDSNew" { return "Rttmonjitterstatsowmaxdsnew" }
    if yname == "rttMonJitterStatsMinOfMOS" { return "Rttmonjitterstatsminofmos" }
    if yname == "rttMonJitterStatsMaxOfMOS" { return "Rttmonjitterstatsmaxofmos" }
    if yname == "rttMonJitterStatsMinOfICPIF" { return "Rttmonjitterstatsminoficpif" }
    if yname == "rttMonJitterStatsMaxOfICPIF" { return "Rttmonjitterstatsmaxoficpif" }
    if yname == "rttMonJitterStatsIAJOut" { return "Rttmonjitterstatsiajout" }
    if yname == "rttMonJitterStatsIAJIn" { return "Rttmonjitterstatsiajin" }
    if yname == "rttMonJitterStatsAvgJitter" { return "Rttmonjitterstatsavgjitter" }
    if yname == "rttMonJitterStatsAvgJitterSD" { return "Rttmonjitterstatsavgjittersd" }
    if yname == "rttMonJitterStatsAvgJitterDS" { return "Rttmonjitterstatsavgjitterds" }
    if yname == "rttMonJitterStatsUnSyncRTs" { return "Rttmonjitterstatsunsyncrts" }
    if yname == "rttMonJitterStatsRTTSumHigh" { return "Rttmonjitterstatsrttsumhigh" }
    if yname == "rttMonJitterStatsOWSumSDHigh" { return "Rttmonjitterstatsowsumsdhigh" }
    if yname == "rttMonJitterStatsOWSumDSHigh" { return "Rttmonjitterstatsowsumdshigh" }
    return ""
}

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetSegmentPath() string {
    return "rttMonJitterStatsEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonjitterstatsentry.Rttmonctrladminindex) + "']" + "[rttMonJitterStatsStartTimeIndex='" + fmt.Sprintf("%v", rttmonjitterstatsentry.Rttmonjitterstatsstarttimeindex) + "']"
}

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonjitterstatsentry.Rttmonctrladminindex
    leafs["rttMonJitterStatsStartTimeIndex"] = rttmonjitterstatsentry.Rttmonjitterstatsstarttimeindex
    leafs["rttMonJitterStatsCompletions"] = rttmonjitterstatsentry.Rttmonjitterstatscompletions
    leafs["rttMonJitterStatsOverThresholds"] = rttmonjitterstatsentry.Rttmonjitterstatsoverthresholds
    leafs["rttMonJitterStatsNumOfRTT"] = rttmonjitterstatsentry.Rttmonjitterstatsnumofrtt
    leafs["rttMonJitterStatsRTTSum"] = rttmonjitterstatsentry.Rttmonjitterstatsrttsum
    leafs["rttMonJitterStatsRTTSum2Low"] = rttmonjitterstatsentry.Rttmonjitterstatsrttsum2Low
    leafs["rttMonJitterStatsRTTSum2High"] = rttmonjitterstatsentry.Rttmonjitterstatsrttsum2High
    leafs["rttMonJitterStatsRTTMin"] = rttmonjitterstatsentry.Rttmonjitterstatsrttmin
    leafs["rttMonJitterStatsRTTMax"] = rttmonjitterstatsentry.Rttmonjitterstatsrttmax
    leafs["rttMonJitterStatsMinOfPositivesSD"] = rttmonjitterstatsentry.Rttmonjitterstatsminofpositivessd
    leafs["rttMonJitterStatsMaxOfPositivesSD"] = rttmonjitterstatsentry.Rttmonjitterstatsmaxofpositivessd
    leafs["rttMonJitterStatsNumOfPositivesSD"] = rttmonjitterstatsentry.Rttmonjitterstatsnumofpositivessd
    leafs["rttMonJitterStatsSumOfPositivesSD"] = rttmonjitterstatsentry.Rttmonjitterstatssumofpositivessd
    leafs["rttMonJitterStatsSum2PositivesSDLow"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Positivessdlow
    leafs["rttMonJitterStatsSum2PositivesSDHigh"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Positivessdhigh
    leafs["rttMonJitterStatsMinOfNegativesSD"] = rttmonjitterstatsentry.Rttmonjitterstatsminofnegativessd
    leafs["rttMonJitterStatsMaxOfNegativesSD"] = rttmonjitterstatsentry.Rttmonjitterstatsmaxofnegativessd
    leafs["rttMonJitterStatsNumOfNegativesSD"] = rttmonjitterstatsentry.Rttmonjitterstatsnumofnegativessd
    leafs["rttMonJitterStatsSumOfNegativesSD"] = rttmonjitterstatsentry.Rttmonjitterstatssumofnegativessd
    leafs["rttMonJitterStatsSum2NegativesSDLow"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Negativessdlow
    leafs["rttMonJitterStatsSum2NegativesSDHigh"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Negativessdhigh
    leafs["rttMonJitterStatsMinOfPositivesDS"] = rttmonjitterstatsentry.Rttmonjitterstatsminofpositivesds
    leafs["rttMonJitterStatsMaxOfPositivesDS"] = rttmonjitterstatsentry.Rttmonjitterstatsmaxofpositivesds
    leafs["rttMonJitterStatsNumOfPositivesDS"] = rttmonjitterstatsentry.Rttmonjitterstatsnumofpositivesds
    leafs["rttMonJitterStatsSumOfPositivesDS"] = rttmonjitterstatsentry.Rttmonjitterstatssumofpositivesds
    leafs["rttMonJitterStatsSum2PositivesDSLow"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Positivesdslow
    leafs["rttMonJitterStatsSum2PositivesDSHigh"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Positivesdshigh
    leafs["rttMonJitterStatsMinOfNegativesDS"] = rttmonjitterstatsentry.Rttmonjitterstatsminofnegativesds
    leafs["rttMonJitterStatsMaxOfNegativesDS"] = rttmonjitterstatsentry.Rttmonjitterstatsmaxofnegativesds
    leafs["rttMonJitterStatsNumOfNegativesDS"] = rttmonjitterstatsentry.Rttmonjitterstatsnumofnegativesds
    leafs["rttMonJitterStatsSumOfNegativesDS"] = rttmonjitterstatsentry.Rttmonjitterstatssumofnegativesds
    leafs["rttMonJitterStatsSum2NegativesDSLow"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Negativesdslow
    leafs["rttMonJitterStatsSum2NegativesDSHigh"] = rttmonjitterstatsentry.Rttmonjitterstatssum2Negativesdshigh
    leafs["rttMonJitterStatsPacketLossSD"] = rttmonjitterstatsentry.Rttmonjitterstatspacketlosssd
    leafs["rttMonJitterStatsPacketLossDS"] = rttmonjitterstatsentry.Rttmonjitterstatspacketlossds
    leafs["rttMonJitterStatsPacketOutOfSequence"] = rttmonjitterstatsentry.Rttmonjitterstatspacketoutofsequence
    leafs["rttMonJitterStatsPacketMIA"] = rttmonjitterstatsentry.Rttmonjitterstatspacketmia
    leafs["rttMonJitterStatsPacketLateArrival"] = rttmonjitterstatsentry.Rttmonjitterstatspacketlatearrival
    leafs["rttMonJitterStatsError"] = rttmonjitterstatsentry.Rttmonjitterstatserror
    leafs["rttMonJitterStatsBusies"] = rttmonjitterstatsentry.Rttmonjitterstatsbusies
    leafs["rttMonJitterStatsOWSumSD"] = rttmonjitterstatsentry.Rttmonjitterstatsowsumsd
    leafs["rttMonJitterStatsOWSum2SDLow"] = rttmonjitterstatsentry.Rttmonjitterstatsowsum2Sdlow
    leafs["rttMonJitterStatsOWSum2SDHigh"] = rttmonjitterstatsentry.Rttmonjitterstatsowsum2Sdhigh
    leafs["rttMonJitterStatsOWMinSD"] = rttmonjitterstatsentry.Rttmonjitterstatsowminsd
    leafs["rttMonJitterStatsOWMaxSD"] = rttmonjitterstatsentry.Rttmonjitterstatsowmaxsd
    leafs["rttMonJitterStatsOWSumDS"] = rttmonjitterstatsentry.Rttmonjitterstatsowsumds
    leafs["rttMonJitterStatsOWSum2DSLow"] = rttmonjitterstatsentry.Rttmonjitterstatsowsum2Dslow
    leafs["rttMonJitterStatsOWSum2DSHigh"] = rttmonjitterstatsentry.Rttmonjitterstatsowsum2Dshigh
    leafs["rttMonJitterStatsOWMinDS"] = rttmonjitterstatsentry.Rttmonjitterstatsowminds
    leafs["rttMonJitterStatsOWMaxDS"] = rttmonjitterstatsentry.Rttmonjitterstatsowmaxds
    leafs["rttMonJitterStatsNumOfOW"] = rttmonjitterstatsentry.Rttmonjitterstatsnumofow
    leafs["rttMonJitterStatsOWMinSDNew"] = rttmonjitterstatsentry.Rttmonjitterstatsowminsdnew
    leafs["rttMonJitterStatsOWMaxSDNew"] = rttmonjitterstatsentry.Rttmonjitterstatsowmaxsdnew
    leafs["rttMonJitterStatsOWMinDSNew"] = rttmonjitterstatsentry.Rttmonjitterstatsowmindsnew
    leafs["rttMonJitterStatsOWMaxDSNew"] = rttmonjitterstatsentry.Rttmonjitterstatsowmaxdsnew
    leafs["rttMonJitterStatsMinOfMOS"] = rttmonjitterstatsentry.Rttmonjitterstatsminofmos
    leafs["rttMonJitterStatsMaxOfMOS"] = rttmonjitterstatsentry.Rttmonjitterstatsmaxofmos
    leafs["rttMonJitterStatsMinOfICPIF"] = rttmonjitterstatsentry.Rttmonjitterstatsminoficpif
    leafs["rttMonJitterStatsMaxOfICPIF"] = rttmonjitterstatsentry.Rttmonjitterstatsmaxoficpif
    leafs["rttMonJitterStatsIAJOut"] = rttmonjitterstatsentry.Rttmonjitterstatsiajout
    leafs["rttMonJitterStatsIAJIn"] = rttmonjitterstatsentry.Rttmonjitterstatsiajin
    leafs["rttMonJitterStatsAvgJitter"] = rttmonjitterstatsentry.Rttmonjitterstatsavgjitter
    leafs["rttMonJitterStatsAvgJitterSD"] = rttmonjitterstatsentry.Rttmonjitterstatsavgjittersd
    leafs["rttMonJitterStatsAvgJitterDS"] = rttmonjitterstatsentry.Rttmonjitterstatsavgjitterds
    leafs["rttMonJitterStatsUnSyncRTs"] = rttmonjitterstatsentry.Rttmonjitterstatsunsyncrts
    leafs["rttMonJitterStatsRTTSumHigh"] = rttmonjitterstatsentry.Rttmonjitterstatsrttsumhigh
    leafs["rttMonJitterStatsOWSumSDHigh"] = rttmonjitterstatsentry.Rttmonjitterstatsowsumsdhigh
    leafs["rttMonJitterStatsOWSumDSHigh"] = rttmonjitterstatsentry.Rttmonjitterstatsowsumdshigh
    return leafs
}

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetYangName() string { return "rttMonJitterStatsEntry" }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) SetParent(parent types.Entity) { rttmonjitterstatsentry.parent = parent }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetParent() types.Entity { return rttmonjitterstatsentry.parent }

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetParentYangName() string { return "rttMonJitterStatsTable" }

// CISCORTTMONMIB_Rttmonlpdgrpstatstable
// The Auto SAA L3 MPLS VPN LPD Group Database.
// 
// The LPD Group statistics table contains summarized performance
// statistics for the LPD group.
// 
// LPD Group - The set of 'single probes' which are subset of the
// 'lspGroup' probe traversing set of paths between two PE end
// points are grouped together and called as the LPD group. The
// LPD group will be uniquely referenced by the LPD Group ID.
// 
// A rolling accumulated history of this information is maintained
// in a series of hourly 'group(s)'.
// 
// Each conceptual statistics row has a current hourly group, into
// which RTT results are accumulated. At the end of each hour a new
// hourly group is created which then becomes current. The
// counters and accumulators in the new group are initialized to
// zero. The previous group(s) is kept in the table until the table
// contains rttMplsVpnMonTypeLpdStatHours groups for the
// conceptual statistics row;  at this point, the oldest group is
// discarded and is replaced by the newly created one. The hourly
// group is uniquely identified by the
// rttMonLpdGrpStatsStartTimeIndex object.
type CISCORTTMONMIB_Rttmonlpdgrpstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a set of RTT operations
    // over a 60 minute time period.  The LPD group statistics table is a rollover
    // table. When rttMonLpdGrpStatsStartTimeIndex groups exceeds the
    // rttMplsVpnMonTypeLpdStatHours value, the oldest corresponding hourly group
    // will be deleted and will be replaced with the new
    // rttMonLpdGrpStatsStartTimeIndex hourly group.  The LPD group statistics
    // table has two indices. Each described as follows:  - The first index
    // correlates its entries to a LPD group via the   
    // rttMonLpdGrpStatsGroupIndex object. - The second index is a rollover group
    // and it uniquely     identifies a 60 minute group. (The    
    // rttMonLpdGrpStatsStartTimeIndex is used to make this value     unique.).
    // The type is slice of
    // CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry.
    Rttmonlpdgrpstatsentry []CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry
}

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetFilter() yfilter.YFilter { return rttmonlpdgrpstatstable.YFilter }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) SetFilter(yf yfilter.YFilter) { rttmonlpdgrpstatstable.YFilter = yf }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetGoName(yname string) string {
    if yname == "rttMonLpdGrpStatsEntry" { return "Rttmonlpdgrpstatsentry" }
    return ""
}

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetSegmentPath() string {
    return "rttMonLpdGrpStatsTable"
}

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonLpdGrpStatsEntry" {
        for _, c := range rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry {
            if rttmonlpdgrpstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry{}
        rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry = append(rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry, child)
        return &rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry[len(rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry)-1]
    }
    return nil
}

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry {
        children[rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry[i].GetSegmentPath()] = &rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry[i]
    }
    return children
}

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetYangName() string { return "rttMonLpdGrpStatsTable" }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) SetParent(parent types.Entity) { rttmonlpdgrpstatstable.parent = parent }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetParent() types.Entity { return rttmonlpdgrpstatstable.parent }

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry
// A list of objects which accumulate the results of a set of RTT
// operations over a 60 minute time period.
// 
// The LPD group statistics table is a rollover table. When
// rttMonLpdGrpStatsStartTimeIndex groups exceeds the
// rttMplsVpnMonTypeLpdStatHours value, the oldest corresponding
// hourly group will be deleted and will be replaced with the new
// rttMonLpdGrpStatsStartTimeIndex hourly group.
// 
// The LPD group statistics table has two indices. Each described
// as follows:
// 
// - The first index correlates its entries to a LPD group via the
//    rttMonLpdGrpStatsGroupIndex object.
// - The second index is a rollover group and it uniquely 
//    identifies a 60 minute group. (The 
//    rttMonLpdGrpStatsStartTimeIndex is used to make this value 
//    unique.)
type CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in
    // rttMonLpdGrpStatsTable.  This is a pseudo-random number which identifies a
    // particular LPD group. The type is interface{} with range: 1..2147483647.
    Rttmonlpdgrpstatsgroupindex interface{}

    // This attribute is a key. The time when this row was created.  This object
    // is the second index of the rttMonLpdGrpStatsTable. When the number of
    // rttMonLpdGrpStatsStartTimeIndex groups exceeds the
    // rttMplsVpnMonTypeLpdStatHours value, the oldest
    // rttMonLpdGrpStatsStartTimeIndex group will be removed and replaced with the
    // new entry. The type is interface{} with range: 0..4294967295.
    Rttmonlpdgrpstatsstarttimeindex interface{}

    // The object is a string that specifies the address of the target PE for this
    // LPD group. The type is string.
    Rttmonlpdgrpstatstargetpe interface{}

    // This object represents the number of successfull completions of 'single
    // probes' for all the set of paths in the LPD group.  Whenever the
    // rttMonLatestRttOperSense value is 'ok' for a particular probe in the LPD
    // Group this object will be incremented.  This object will be set to '0' on
    // reset. The type is interface{} with range: 0..2147483647. Units are passes.
    Rttmonlpdgrpstatsnumofpass interface{}

    // This object represents the number of failed operations of 'single probes'
    // for all the set of paths in the LPD group.  Whenever the
    // rttMonLatestRttOperSense has a value other than 'ok' or 'timeout' for a
    // particular probe in the LPD Group this object will be incremented.  This
    // object will be set to '0' on reset. The type is interface{} with range:
    // 0..2147483647. Units are failures.
    Rttmonlpdgrpstatsnumoffail interface{}

    // This object represents the number of timed out operations of 'single
    // probes' for all the set of paths in the LPD group.  Whenever the
    // rttMonLatestRttOperSense has a value of 'timeout' for a particular probe in
    // the LPD Group this object will be incremented.  This object will be set to
    // '0' on reset. The type is interface{} with range: 0..2147483647. Units are
    // timeouts.
    Rttmonlpdgrpstatsnumoftimeout interface{}

    // The average RTT across all set of probes in the LPD group.  This object
    // will be set to '0' on reset. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    Rttmonlpdgrpstatsavgrtt interface{}

    // The minimum of RTT's for all set of probes in the LPD group that were
    // successfully measured.  This object will be set to '0' on reset. The type
    // is interface{} with range: 0..2147483647. Units are milliseconds.
    Rttmonlpdgrpstatsminrtt interface{}

    // The maximum of RTT's for all set of probes in the LPD group that were
    // successfully measured.  This object will be set to '0' on reset. The type
    // is interface{} with range: 0..2147483647. Units are milliseconds.
    Rttmonlpdgrpstatsmaxrtt interface{}

    // The minimum number of active paths discovered to the
    // rttMonLpdGrpStatsTargetPE target.  This object will be set to '0' on reset.
    // The type is interface{} with range: 0..2147483647. Units are paths.
    Rttmonlpdgrpstatsminnumpaths interface{}

    // The maximum number of active paths discovered to the
    // rttMonLpdGrpStatsTargetPE target.  This object will be set to '0' on reset.
    // The type is interface{} with range: 0..2147483647. Units are paths.
    Rttmonlpdgrpstatsmaxnumpaths interface{}

    // The time when the last LSP Path Discovery to the group was attempted.  This
    // object will be set to '0' on reset. The type is interface{} with range:
    // 0..4294967295. Units are tenths of milliseconds.
    Rttmonlpdgrpstatslpdstarttime interface{}

    // This object is set to true when the LSP Path Discovery to the target PE
    // i.e. rttMonLpdGrpStatsTargetPE fails, and set to false when the LSP Path
    // Discovery succeeds.  When this value changes and
    // rttMplsVpnMonReactLpdNotifyType is set to 'lpdPathDiscovery' or 'lpdAll' a
    // rttMonLpdDiscoveryNotification will be generated.  This object will be set
    // to 'FALSE' on reset. The type is bool.
    Rttmonlpdgrpstatslpdfailoccurred interface{}

    // This object identifies the cause of failure for the LSP Path Discovery last
    // attempted. It will be only valid if rttMonLpdGrpStatsLPDFailOccurred is set
    // to true.  This object will be set to 'unknown' on reset. The type is
    // RttMplsVpnMonLpdFailureSense.
    Rttmonlpdgrpstatslpdfailcause interface{}

    // The completion time of the last successfull LSP Path Discovery to the
    // target PE.  This object will be set to '0' on reset. The type is
    // interface{} with range: 0..65535. Units are seconds.
    Rttmonlpdgrpstatslpdcomptime interface{}

    // This object identifies the LPD Group status.  When the LPD Group status
    // changes and rttMplsVpnMonReactLpdNotifyType is set to 'lpdGroupStatus' or
    // 'lpdAll' a rttMonLpdGrpStatusNotification will be generated.  When the LPD
    // Group status value is 'unknown' or changes to 'unknown' this notification
    // will not be generated.  When LSP Path Discovery is enabled for a particular
    // row in rttMplsVpnMonCtrlTable, 'single probes' in the 'lspGroup' probe
    // cannot generate notifications independently but will be generating
    // depending on the state of the group. Notifications  are only generated if
    // the failure/restoration of an individual probe causes the state of the LPD
    // Group to change.  This object will be set to 'unknown' on reset. The type
    // is RttMplsVpnMonLpdGrpStatus.
    Rttmonlpdgrpstatsgroupstatus interface{}

    // This object identifies 'lspGroup' probe uniquely created for this
    // particular LPD Group. The type is interface{} with range: 1..2147483647.
    // Units are identifier.
    Rttmonlpdgrpstatsgroupprobeindex interface{}

    // A string which holds the list of information to uniquely identify the paths
    // to the target PE. This information is used by the 'single probes' when
    // testing the paths.  Following three parameters are needed to uniquely
    // identify a  path   - lsp-selector (127.x.x.x)   - outgoing-interface (i/f) 
    // - label-stack (s), if mutiple labels they will be colon (:)     separated. 
    // These parameters will be hyphen (-) separated for a particular path. This
    // set of information will be comma (,) separated for all the paths discovered
    // as part of this LPD Group.  For example: If there are 5 paths in the LPD
    // group then this object will return all the identifier's to uniquely
    // identify the path.  The output will look like '127.0.0.1-Se3/0.1-20:18,
    // 127.0.0.2-Se3/0.1-20,127.0.0.3-Se3/0.1-20,127.0.0.4-Se3/0.1-20,
    // 127.0.0.5-Se3/0.1-20'.  This object will be set to '0' on reset. The type
    // is string.
    Rttmonlpdgrpstatspathids interface{}

    // A string which holds the latest operation return code for all the set of
    // 'single probes' which are part of the LPD group. The return codes will be
    // comma separated and will follow the same sequence of probes as followed in
    // 'rttMonLpdGrpStatsPathIds'. The latest operation return code will be mapped
    // to 'up','down' or 'unkwown'.  'up' - Probe state is up when the
    // rttMonLatestRttOperSense value is 'ok'. 'down' - Probe state is down when
    // the rttMonLatestRttOperSense has value other then 'ok' and 'other'.
    // 'unknown' - Probe state is unkown when the rttMonLatestRttOperSense value
    // is 'other'.  For example: If there are 5 paths in the LPD group then this
    // object output will look like 'ok,ok,ok,down,down'.  This object will be set
    // to '0' on reset. The type is string.
    Rttmonlpdgrpstatsprobestatus interface{}

    // This object specifies the time when this statistics row was last reset
    // using the rttMonApplLpdGrpStatsReset object. The type is interface{} with
    // range: 0..4294967295.
    Rttmonlpdgrpstatsresettime interface{}
}

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetFilter() yfilter.YFilter { return rttmonlpdgrpstatsentry.YFilter }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) SetFilter(yf yfilter.YFilter) { rttmonlpdgrpstatsentry.YFilter = yf }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetGoName(yname string) string {
    if yname == "rttMonLpdGrpStatsGroupIndex" { return "Rttmonlpdgrpstatsgroupindex" }
    if yname == "rttMonLpdGrpStatsStartTimeIndex" { return "Rttmonlpdgrpstatsstarttimeindex" }
    if yname == "rttMonLpdGrpStatsTargetPE" { return "Rttmonlpdgrpstatstargetpe" }
    if yname == "rttMonLpdGrpStatsNumOfPass" { return "Rttmonlpdgrpstatsnumofpass" }
    if yname == "rttMonLpdGrpStatsNumOfFail" { return "Rttmonlpdgrpstatsnumoffail" }
    if yname == "rttMonLpdGrpStatsNumOfTimeout" { return "Rttmonlpdgrpstatsnumoftimeout" }
    if yname == "rttMonLpdGrpStatsAvgRTT" { return "Rttmonlpdgrpstatsavgrtt" }
    if yname == "rttMonLpdGrpStatsMinRTT" { return "Rttmonlpdgrpstatsminrtt" }
    if yname == "rttMonLpdGrpStatsMaxRTT" { return "Rttmonlpdgrpstatsmaxrtt" }
    if yname == "rttMonLpdGrpStatsMinNumPaths" { return "Rttmonlpdgrpstatsminnumpaths" }
    if yname == "rttMonLpdGrpStatsMaxNumPaths" { return "Rttmonlpdgrpstatsmaxnumpaths" }
    if yname == "rttMonLpdGrpStatsLPDStartTime" { return "Rttmonlpdgrpstatslpdstarttime" }
    if yname == "rttMonLpdGrpStatsLPDFailOccurred" { return "Rttmonlpdgrpstatslpdfailoccurred" }
    if yname == "rttMonLpdGrpStatsLPDFailCause" { return "Rttmonlpdgrpstatslpdfailcause" }
    if yname == "rttMonLpdGrpStatsLPDCompTime" { return "Rttmonlpdgrpstatslpdcomptime" }
    if yname == "rttMonLpdGrpStatsGroupStatus" { return "Rttmonlpdgrpstatsgroupstatus" }
    if yname == "rttMonLpdGrpStatsGroupProbeIndex" { return "Rttmonlpdgrpstatsgroupprobeindex" }
    if yname == "rttMonLpdGrpStatsPathIds" { return "Rttmonlpdgrpstatspathids" }
    if yname == "rttMonLpdGrpStatsProbeStatus" { return "Rttmonlpdgrpstatsprobestatus" }
    if yname == "rttMonLpdGrpStatsResetTime" { return "Rttmonlpdgrpstatsresettime" }
    return ""
}

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetSegmentPath() string {
    return "rttMonLpdGrpStatsEntry" + "[rttMonLpdGrpStatsGroupIndex='" + fmt.Sprintf("%v", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupindex) + "']" + "[rttMonLpdGrpStatsStartTimeIndex='" + fmt.Sprintf("%v", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsstarttimeindex) + "']"
}

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonLpdGrpStatsGroupIndex"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupindex
    leafs["rttMonLpdGrpStatsStartTimeIndex"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsstarttimeindex
    leafs["rttMonLpdGrpStatsTargetPE"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatstargetpe
    leafs["rttMonLpdGrpStatsNumOfPass"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsnumofpass
    leafs["rttMonLpdGrpStatsNumOfFail"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsnumoffail
    leafs["rttMonLpdGrpStatsNumOfTimeout"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsnumoftimeout
    leafs["rttMonLpdGrpStatsAvgRTT"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsavgrtt
    leafs["rttMonLpdGrpStatsMinRTT"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsminrtt
    leafs["rttMonLpdGrpStatsMaxRTT"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsmaxrtt
    leafs["rttMonLpdGrpStatsMinNumPaths"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsminnumpaths
    leafs["rttMonLpdGrpStatsMaxNumPaths"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsmaxnumpaths
    leafs["rttMonLpdGrpStatsLPDStartTime"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdstarttime
    leafs["rttMonLpdGrpStatsLPDFailOccurred"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdfailoccurred
    leafs["rttMonLpdGrpStatsLPDFailCause"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdfailcause
    leafs["rttMonLpdGrpStatsLPDCompTime"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdcomptime
    leafs["rttMonLpdGrpStatsGroupStatus"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupstatus
    leafs["rttMonLpdGrpStatsGroupProbeIndex"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupprobeindex
    leafs["rttMonLpdGrpStatsPathIds"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatspathids
    leafs["rttMonLpdGrpStatsProbeStatus"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsprobestatus
    leafs["rttMonLpdGrpStatsResetTime"] = rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsresettime
    return leafs
}

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetYangName() string { return "rttMonLpdGrpStatsEntry" }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) SetParent(parent types.Entity) { rttmonlpdgrpstatsentry.parent = parent }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetParent() types.Entity { return rttmonlpdgrpstatsentry.parent }

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetParentYangName() string { return "rttMonLpdGrpStatsTable" }

// CISCORTTMONMIB_Rttmonhistorycollectiontable
// The history collection database.
// 
// The history table contains a point by point rolling 
// history of the most recent RTT operations for each 
// conceptual RTT control row.  The rolling history of this 
// information is maintained in a series of 'live(s)', each
// containing a series of 'bucket(s)', each 'bucket' 
// contains a series of 'sample(s)'.
// 
// Each conceptual history row can have lives.  A life is 
// defined by the rttMonCtrlOperRttLife object.  A new life 
// will be created when rttMonCtrlOperState transitions
// 'active'.  When the number of lives become greater 
// than rttMonHistoryAdminNumLives the oldest life will be 
// discarded and a new life will be created by incrementing
// the index.
// 
// The path exploration RTT operation will be kept as an
// entry in this table.
type CISCORTTMONMIB_Rttmonhistorycollectiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of history objects that are recorded for each RTT operation.  The
    // history collection table has four indices.  Each  described as follows:   -
    // The first index correlates its entries to a        conceptual RTT control
    // row via the        rttMonCtrlAdminIndex object.     -  The second index
    // uniquely identifies the results        of each 'life' as defined by the    
    // rttMonCtrlOperRttLife object.     -  The third index uniquely identifies
    // the number of        buckets in a life.  A bucket will contain one       
    // sample per bucket if the rttMonCtrlAdminRttType        object is set to any
    // value       other than 'pathEcho'.  If the        rttMonCtrlAdminRttType
    // object is set to        'pathEcho', a bucket will contain one sample per   
    // hop along a path to the target (including the        target).     -  The
    // fourth index uniquely identifies the number of        samples in a bucket. 
    // Again, if the        rttMonCtrlAdminRttType object is set to       
    // 'pathEcho', this value is associated with each        hop in an ascending
    // order, thus for the        first hop on a path, this index will be 1, the  
    // second will be 2 and so on.   For all other values       of
    // rttMonCtrlAdminRttType this will be 1. The type is slice of
    // CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry.
    Rttmonhistorycollectionentry []CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry
}

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetFilter() yfilter.YFilter { return rttmonhistorycollectiontable.YFilter }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) SetFilter(yf yfilter.YFilter) { rttmonhistorycollectiontable.YFilter = yf }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetGoName(yname string) string {
    if yname == "rttMonHistoryCollectionEntry" { return "Rttmonhistorycollectionentry" }
    return ""
}

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetSegmentPath() string {
    return "rttMonHistoryCollectionTable"
}

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonHistoryCollectionEntry" {
        for _, c := range rttmonhistorycollectiontable.Rttmonhistorycollectionentry {
            if rttmonhistorycollectiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry{}
        rttmonhistorycollectiontable.Rttmonhistorycollectionentry = append(rttmonhistorycollectiontable.Rttmonhistorycollectionentry, child)
        return &rttmonhistorycollectiontable.Rttmonhistorycollectionentry[len(rttmonhistorycollectiontable.Rttmonhistorycollectionentry)-1]
    }
    return nil
}

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonhistorycollectiontable.Rttmonhistorycollectionentry {
        children[rttmonhistorycollectiontable.Rttmonhistorycollectionentry[i].GetSegmentPath()] = &rttmonhistorycollectiontable.Rttmonhistorycollectionentry[i]
    }
    return children
}

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetYangName() string { return "rttMonHistoryCollectionTable" }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) SetParent(parent types.Entity) { rttmonhistorycollectiontable.parent = parent }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetParent() types.Entity { return rttmonhistorycollectiontable.parent }

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry
// A list of history objects that are recorded for each
// RTT operation.
// 
// The history collection table has four indices.  Each 
// described as follows:
//   -  The first index correlates its entries to a 
//       conceptual RTT control row via the 
//       rttMonCtrlAdminIndex object.  
//   -  The second index uniquely identifies the results 
//       of each 'life' as defined by the 
//       rttMonCtrlOperRttLife object.  
//   -  The third index uniquely identifies the number of 
//       buckets in a life.  A bucket will contain one 
//       sample per bucket if the rttMonCtrlAdminRttType 
//       object is set to any value
//       other than 'pathEcho'.  If the 
//       rttMonCtrlAdminRttType object is set to 
//       'pathEcho', a bucket will contain one sample per 
//       hop along a path to the target (including the 
//       target).  
//   -  The fourth index uniquely identifies the number of 
//       samples in a bucket.   Again, if the 
//       rttMonCtrlAdminRttType object is set to 
//       'pathEcho', this value is associated with each 
//       hop in an ascending order, thus for the 
//       first hop on a path, this index will be 1, the 
//       second will be 2 and so on.   For all other values
//       of rttMonCtrlAdminRttType this will be 1.
type CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // This attribute is a key. This uniquely defines a life for a conceptual
    // history row.  For a particular value of rttMonHistoryCollectionLifeIndex,
    // the agent assigns the first value of 1, the second value  of 2, and so on. 
    // The sequence keeps incrementing,  despite older (lower) values being
    // removed from the  table. The type is interface{} with range: 1..2147483647.
    Rttmonhistorycollectionlifeindex interface{}

    // This attribute is a key. When the RttMonRttType is 'pathEcho', this
    // uniquely defines a bucket for a given value of 
    // rttMonHistoryCollectionLifeIndex.  For all other  RttMonRttType this value
    // will be the number of operations per a lifetime.  Thus, this object 
    // increments on each operation attempt.  For a particular value of 
    // rttMonHistoryCollectionLifeIndex, the agent assigns  the first value of 1,
    // the second value of 2, and so  on.  The sequence keeps incrementing until
    // the number of buckets equals rttMonHistoryAdminNumBuckets, after which the
    // most recent rttMonHistoryAdminNumBuckets  buckets are retained (the index
    // is incremented though). The type is interface{} with range: 1..2147483647.
    Rttmonhistorycollectionbucketindex interface{}

    // This attribute is a key. This uniquely defines a row for a given value of
    // rttMonHistoryCollectionBucketIndex.  This object represents a hop along a
    // path to the Target.  For a particular value of 
    // rttMonHistoryCollectionBucketIndex, the agent assigns  the first value of
    // 1, the second value of 2, and so on. The sequence keeps incrementing until
    // the number of  samples equals rttMonHistoryAdminNumSamples, then no  new
    // samples are created for the current  rttMonHistoryCollectionBucketIndex. 
    // When the RttMonRttType is 'pathEcho', this value  directly represents the
    // number of hops along a  path to a target, thus we can only support 512
    // hops. For all other values of RttMonRttType this object will be one. The
    // type is interface{} with range: 1..512.
    Rttmonhistorycollectionsampleindex interface{}

    // The time that the RTT operation was initiated. The type is interface{} with
    // range: 0..4294967295.
    Rttmonhistorycollectionsampletime interface{}

    // When the RttMonRttType is 'echo' or 'pathEcho' this is a string which
    // specifies the address of the target for the this RTT operation.  For all
    // other values of RttMonRttType this string will be null.  This address will
    // be the address of the hop along the path to the
    // rttMonEchoAdminTargetAddress address, including
    // rttMonEchoAdminTargetAddress address, or just the
    // rttMonEchoAdminTargetAddress address, when the path information is not
    // collected.  This behavior is defined by the rttMonCtrlAdminRttType object. 
    // The interpretation of this string depends on the type of RTT operation
    // selected, as specified by the rttMonEchoAdminProtocol object.  See
    // rttMonEchoAdminTargetAddress for a complete description. The type is
    // string.
    Rttmonhistorycollectionaddress interface{}

    // This is the operation completion time of the RTT operation.  If the RTT
    // operation fails  (rttMonHistoryCollectionSense is any  value other than
    // ok), this has a value of 0. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Rttmonhistorycollectioncompletiontime interface{}

    // A sense code for the completion status of the RTT operation. The type is
    // RttResponseSense.
    Rttmonhistorycollectionsense interface{}

    // An application specific sense code for the completion status of the last
    // RTT operation.  This  object will only be valid when the 
    // rttMonHistoryCollectionSense object is set to  'applicationSpecific'. 
    // Otherwise, this object's  value is not valid. The type is interface{} with
    // range: 0..2147483647.
    Rttmonhistorycollectionapplspecificsense interface{}

    // A sense description for the completion status of the last RTT operation
    // when the  rttMonHistoryCollectionSense object is set to 
    // 'applicationSpecific'. The type is string.
    Rttmonhistorycollectionsensedescription interface{}
}

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetFilter() yfilter.YFilter { return rttmonhistorycollectionentry.YFilter }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) SetFilter(yf yfilter.YFilter) { rttmonhistorycollectionentry.YFilter = yf }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonHistoryCollectionLifeIndex" { return "Rttmonhistorycollectionlifeindex" }
    if yname == "rttMonHistoryCollectionBucketIndex" { return "Rttmonhistorycollectionbucketindex" }
    if yname == "rttMonHistoryCollectionSampleIndex" { return "Rttmonhistorycollectionsampleindex" }
    if yname == "rttMonHistoryCollectionSampleTime" { return "Rttmonhistorycollectionsampletime" }
    if yname == "rttMonHistoryCollectionAddress" { return "Rttmonhistorycollectionaddress" }
    if yname == "rttMonHistoryCollectionCompletionTime" { return "Rttmonhistorycollectioncompletiontime" }
    if yname == "rttMonHistoryCollectionSense" { return "Rttmonhistorycollectionsense" }
    if yname == "rttMonHistoryCollectionApplSpecificSense" { return "Rttmonhistorycollectionapplspecificsense" }
    if yname == "rttMonHistoryCollectionSenseDescription" { return "Rttmonhistorycollectionsensedescription" }
    return ""
}

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetSegmentPath() string {
    return "rttMonHistoryCollectionEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonctrladminindex) + "']" + "[rttMonHistoryCollectionLifeIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonhistorycollectionlifeindex) + "']" + "[rttMonHistoryCollectionBucketIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonhistorycollectionbucketindex) + "']" + "[rttMonHistoryCollectionSampleIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonhistorycollectionsampleindex) + "']"
}

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonhistorycollectionentry.Rttmonctrladminindex
    leafs["rttMonHistoryCollectionLifeIndex"] = rttmonhistorycollectionentry.Rttmonhistorycollectionlifeindex
    leafs["rttMonHistoryCollectionBucketIndex"] = rttmonhistorycollectionentry.Rttmonhistorycollectionbucketindex
    leafs["rttMonHistoryCollectionSampleIndex"] = rttmonhistorycollectionentry.Rttmonhistorycollectionsampleindex
    leafs["rttMonHistoryCollectionSampleTime"] = rttmonhistorycollectionentry.Rttmonhistorycollectionsampletime
    leafs["rttMonHistoryCollectionAddress"] = rttmonhistorycollectionentry.Rttmonhistorycollectionaddress
    leafs["rttMonHistoryCollectionCompletionTime"] = rttmonhistorycollectionentry.Rttmonhistorycollectioncompletiontime
    leafs["rttMonHistoryCollectionSense"] = rttmonhistorycollectionentry.Rttmonhistorycollectionsense
    leafs["rttMonHistoryCollectionApplSpecificSense"] = rttmonhistorycollectionentry.Rttmonhistorycollectionapplspecificsense
    leafs["rttMonHistoryCollectionSenseDescription"] = rttmonhistorycollectionentry.Rttmonhistorycollectionsensedescription
    return leafs
}

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetYangName() string { return "rttMonHistoryCollectionEntry" }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) SetParent(parent types.Entity) { rttmonhistorycollectionentry.parent = parent }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetParent() types.Entity { return rttmonhistorycollectionentry.parent }

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetParentYangName() string { return "rttMonHistoryCollectionTable" }

// CISCORTTMONMIB_Rttmonlatesthttpopertable
// A table which contains the status of latest HTTP RTT
// operation.
type CISCORTTMONMIB_Rttmonlatesthttpopertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that record the latest HTTP RTT operation. This entry is
    // created automatically after the  rttMonCtrlAdminEntry is created. Also the
    // entry is  automatically deleted when rttMonCtrlAdminEntry is deleted. The
    // type is slice of
    // CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry.
    Rttmonlatesthttpoperentry []CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry
}

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetFilter() yfilter.YFilter { return rttmonlatesthttpopertable.YFilter }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) SetFilter(yf yfilter.YFilter) { rttmonlatesthttpopertable.YFilter = yf }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetGoName(yname string) string {
    if yname == "rttMonLatestHTTPOperEntry" { return "Rttmonlatesthttpoperentry" }
    return ""
}

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetSegmentPath() string {
    return "rttMonLatestHTTPOperTable"
}

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonLatestHTTPOperEntry" {
        for _, c := range rttmonlatesthttpopertable.Rttmonlatesthttpoperentry {
            if rttmonlatesthttpopertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry{}
        rttmonlatesthttpopertable.Rttmonlatesthttpoperentry = append(rttmonlatesthttpopertable.Rttmonlatesthttpoperentry, child)
        return &rttmonlatesthttpopertable.Rttmonlatesthttpoperentry[len(rttmonlatesthttpopertable.Rttmonlatesthttpoperentry)-1]
    }
    return nil
}

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonlatesthttpopertable.Rttmonlatesthttpoperentry {
        children[rttmonlatesthttpopertable.Rttmonlatesthttpoperentry[i].GetSegmentPath()] = &rttmonlatesthttpopertable.Rttmonlatesthttpoperentry[i]
    }
    return children
}

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetYangName() string { return "rttMonLatestHTTPOperTable" }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) SetParent(parent types.Entity) { rttmonlatesthttpopertable.parent = parent }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetParent() types.Entity { return rttmonlatesthttpopertable.parent }

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry
// A list of objects that record the latest HTTP RTT
// operation. This entry is created automatically after the 
// rttMonCtrlAdminEntry is created. Also the entry is 
// automatically deleted when rttMonCtrlAdminEntry is deleted.
type CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // Round Trip Time taken to perform HTTP operation. This value is the sum of
    // DNSRTT, TCPConnectRTT and TransactionRTT. The type is interface{} with
    // range: 0..4294967295.
    Rttmonlatesthttpoperrtt interface{}

    // Round Trip Time taken to perform DNS query within the HTTP operation. If an
    // IP Address is specified in the URL,  then DNSRTT is 0. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatesthttpoperdnsrtt interface{}

    // Round Trip Time taken to connect to the HTTP server. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatesthttpopertcpconnectrtt interface{}

    // Round Trip Time taken to download the object specified by the URL. The type
    // is interface{} with range: 0..4294967295.
    Rttmonlatesthttpopertransactionrtt interface{}

    // The size of the message body received as a response to the HTTP request.
    // The type is interface{} with range: 0..4294967295.
    Rttmonlatesthttpopermessagebodyoctets interface{}

    // An application specific sense code for the completion status of the latest
    // RTT operation. The type is RttResponseSense.
    Rttmonlatesthttpopersense interface{}

    // An sense description for the completion status of the latest RTT operation.
    // The type is string.
    Rttmonlatesthttperrorsensedescription interface{}
}

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetFilter() yfilter.YFilter { return rttmonlatesthttpoperentry.YFilter }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) SetFilter(yf yfilter.YFilter) { rttmonlatesthttpoperentry.YFilter = yf }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonLatestHTTPOperRTT" { return "Rttmonlatesthttpoperrtt" }
    if yname == "rttMonLatestHTTPOperDNSRTT" { return "Rttmonlatesthttpoperdnsrtt" }
    if yname == "rttMonLatestHTTPOperTCPConnectRTT" { return "Rttmonlatesthttpopertcpconnectrtt" }
    if yname == "rttMonLatestHTTPOperTransactionRTT" { return "Rttmonlatesthttpopertransactionrtt" }
    if yname == "rttMonLatestHTTPOperMessageBodyOctets" { return "Rttmonlatesthttpopermessagebodyoctets" }
    if yname == "rttMonLatestHTTPOperSense" { return "Rttmonlatesthttpopersense" }
    if yname == "rttMonLatestHTTPErrorSenseDescription" { return "Rttmonlatesthttperrorsensedescription" }
    return ""
}

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetSegmentPath() string {
    return "rttMonLatestHTTPOperEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonlatesthttpoperentry.Rttmonctrladminindex) + "']"
}

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonlatesthttpoperentry.Rttmonctrladminindex
    leafs["rttMonLatestHTTPOperRTT"] = rttmonlatesthttpoperentry.Rttmonlatesthttpoperrtt
    leafs["rttMonLatestHTTPOperDNSRTT"] = rttmonlatesthttpoperentry.Rttmonlatesthttpoperdnsrtt
    leafs["rttMonLatestHTTPOperTCPConnectRTT"] = rttmonlatesthttpoperentry.Rttmonlatesthttpopertcpconnectrtt
    leafs["rttMonLatestHTTPOperTransactionRTT"] = rttmonlatesthttpoperentry.Rttmonlatesthttpopertransactionrtt
    leafs["rttMonLatestHTTPOperMessageBodyOctets"] = rttmonlatesthttpoperentry.Rttmonlatesthttpopermessagebodyoctets
    leafs["rttMonLatestHTTPOperSense"] = rttmonlatesthttpoperentry.Rttmonlatesthttpopersense
    leafs["rttMonLatestHTTPErrorSenseDescription"] = rttmonlatesthttpoperentry.Rttmonlatesthttperrorsensedescription
    return leafs
}

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetYangName() string { return "rttMonLatestHTTPOperEntry" }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) SetParent(parent types.Entity) { rttmonlatesthttpoperentry.parent = parent }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetParent() types.Entity { return rttmonlatesthttpoperentry.parent }

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetParentYangName() string { return "rttMonLatestHTTPOperTable" }

// CISCORTTMONMIB_Rttmonlatestjitteropertable
// A table which contains the status of latest Jitter
// operation.
type CISCORTTMONMIB_Rttmonlatestjitteropertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that record the latest Jitter operation. The type is
    // slice of
    // CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry.
    Rttmonlatestjitteroperentry []CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry
}

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetFilter() yfilter.YFilter { return rttmonlatestjitteropertable.YFilter }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) SetFilter(yf yfilter.YFilter) { rttmonlatestjitteropertable.YFilter = yf }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetGoName(yname string) string {
    if yname == "rttMonLatestJitterOperEntry" { return "Rttmonlatestjitteroperentry" }
    return ""
}

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetSegmentPath() string {
    return "rttMonLatestJitterOperTable"
}

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rttMonLatestJitterOperEntry" {
        for _, c := range rttmonlatestjitteropertable.Rttmonlatestjitteroperentry {
            if rttmonlatestjitteropertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry{}
        rttmonlatestjitteropertable.Rttmonlatestjitteroperentry = append(rttmonlatestjitteropertable.Rttmonlatestjitteroperentry, child)
        return &rttmonlatestjitteropertable.Rttmonlatestjitteroperentry[len(rttmonlatestjitteropertable.Rttmonlatestjitteroperentry)-1]
    }
    return nil
}

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rttmonlatestjitteropertable.Rttmonlatestjitteroperentry {
        children[rttmonlatestjitteropertable.Rttmonlatestjitteroperentry[i].GetSegmentPath()] = &rttmonlatestjitteropertable.Rttmonlatestjitteroperentry[i]
    }
    return children
}

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetYangName() string { return "rttMonLatestJitterOperTable" }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) SetParent(parent types.Entity) { rttmonlatestjitteropertable.parent = parent }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetParent() types.Entity { return rttmonlatestjitteropertable.parent }

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetParentYangName() string { return "CISCO-RTTMON-MIB" }

// CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry
// A list of objects that record the latest Jitter
// operation.
type CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry_Rttmonctrladminindex
    Rttmonctrladminindex interface{}

    // The number of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteropernumofrtt interface{}

    // The sum of Jitter RTT's that are successfully measured (low order 32 bits).
    // The high order 32 bits are stored in rttMonLatestJitterOperRTTSumHigh. The
    // type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperrttsum interface{}

    // The sum of squares of RTT's that are successfully measured (low order 32
    // bits). The high order 32 bits are stored in
    // rttMonLatestJitterOperRTTSum2High. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteroperrttsum2 interface{}

    // The minimum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperrttmin interface{}

    // The maximum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperrttmax interface{}

    // The minimum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperminofpositivessd interface{}

    // The maximum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropermaxofpositivessd interface{}

    // The sum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropernumofpositivessd interface{}

    // The sum of RTT's of all positive jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropersumofpositivessd interface{}

    // The sum of square of RTT's of all positive jitter values from packets sent
    // from source to destination. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteropersum2Positivessd interface{}

    // The minimum of absolute values of all negative jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteroperminofnegativessd interface{}

    // The maximum of absolute values of all negative jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteropermaxofnegativessd interface{}

    // The sum of number of all negative jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropernumofnegativessd interface{}

    // The sum of all negative jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropersumofnegativessd interface{}

    // The sum of square of RTT's of all negative jitter values from packets sent
    // from source to destination. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteropersum2Negativessd interface{}

    // The minimum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperminofpositivesds interface{}

    // The maximum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropermaxofpositivesds interface{}

    // The sum of number of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropernumofpositivesds interface{}

    // The sum of RTT's of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropersumofpositivesds interface{}

    // The sum of squares of RTT's of all positive jitter values from packets sent
    // from destination to source. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteropersum2Positivesds interface{}

    // The minimum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperminofnegativesds interface{}

    // The maximum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropermaxofnegativesds interface{}

    // The sum of number of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropernumofnegativesds interface{}

    // The sum of RTT's of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteropersumofnegativesds interface{}

    // The sum of squares of RTT's of all negative jitter values from packets sent
    // from destination to source. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteropersum2Negativesds interface{}

    // The number of packets lost when sent from source to destination. The type
    // is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperpacketlosssd interface{}

    // The number of packets lost when sent from destination to source. The type
    // is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperpacketlossds interface{}

    // The number of packets arrived out of sequence. The type is interface{} with
    // range: 0..4294967295.
    Rttmonlatestjitteroperpacketoutofsequence interface{}

    // The number of packets that are lost for which we cannot determine the
    // direction. The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperpacketmia interface{}

    // The number of packets that arrived after the timeout. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperpacketlatearrival interface{}

    // An application specific sense code for the completion status of the latest
    // Jitter RTT operation. The type is RttResponseSense.
    Rttmonlatestjitteropersense interface{}

    // An sense description for the completion status of the latest Jitter RTT
    // operation. The type is string.
    Rttmonlatestjittererrorsensedescription interface{}

    // The sum of one way latency from source to destination (low order 32 bits).
    // The high order 32 bits are stored in rttMonLatestJitterOperOWSumSDHigh. The
    // type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowsumsd interface{}

    // The sum of squares of one way latency from source to destination (low order
    // 32 bits). The high order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2SDHigh. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteroperowsum2Sd interface{}

    // The minimum of all one way latency from source to destination. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowminsd interface{}

    // The maximum of all one way latency from source to destination. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowmaxsd interface{}

    // The sum of one way latency from destination to source (low order 32 bits).
    // The high order 32 bits are stored in rttMonLatestJitterOperOWSumDSHigh. The
    // type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowsumds interface{}

    // The sum of squares of one way latency from destination to source (low order
    // 32 bits). The high order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2DSHigh. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteroperowsum2Ds interface{}

    // The minimum of all one way latency from destination to source. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowminds interface{}

    // The maximum of all one way latency from destination to source. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowmaxds interface{}

    // The number of successful one way latency measurements. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteropernumofow interface{}

    // The MOS value for the latest jitter operation in hundreds. This value will
    // be 0 if   - rttMonEchoAdminCodecType of the operation is notApplicable   -
    // the operation is not started   - the operation is started but failed This
    // value will be 1 for packet loss of 10% or more. The type is interface{}
    // with range: 0..None | 100..500.
    Rttmonlatestjitteropermos interface{}

    // Represents ICPIF value for the latest jitter operation.  This value will be
    // 93 for packet loss of 10% or more. The type is interface{} with range:
    // 0..2147483647.
    Rttmonlatestjitteropericpif interface{}

    // Interarrival Jitter (RC1889) at responder. The type is interface{} with
    // range: 0..2147483647.
    Rttmonlatestjitteroperiajout interface{}

    // Interarrival Jitter (RFC1889) at source. The type is interface{} with
    // range: 0..2147483647.
    Rttmonlatestjitteroperiajin interface{}

    // The average of positive and negative jitter values in SD and DS direction
    // for latest operation. The type is interface{} with range: 0..2147483647.
    Rttmonlatestjitteroperavgjitter interface{}

    // The average of positive and negative jitter values from source to
    // destination for latest operation. The type is interface{} with range:
    // 0..2147483647.
    Rttmonlatestjitteroperavgsdj interface{}

    // The average of positive and negative jitter values from destination to
    // source for latest operation. The type is interface{} with range:
    // 0..2147483647.
    Rttmonlatestjitteroperavgdsj interface{}

    // The average latency value from source to destination. The type is
    // interface{} with range: 0..2147483647.
    Rttmonlatestjitteroperowavgsd interface{}

    // The average latency value from destination to source. The type is
    // interface{} with range: 0..2147483647.
    Rttmonlatestjitteroperowavgds interface{}

    // A value of sync(1) means sender and responder was in sync with NTP. The NTP
    // sync means the total of NTP offset  on sender and responder is within
    // configured tolerance level. The type is Rttmonlatestjitteroperntpstate.
    Rttmonlatestjitteroperntpstate interface{}

    // The number of RTT operations that have completed with sender and responder
    // out of sync with NTP. The NTP sync means  the total of NTP offset on sender
    // and responder is within  configured tolerance level. The type is
    // interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperunsyncrts interface{}

    // The sum of Jitter RTT's that are successfully measured. (high order 32
    // bits). The low order 32 bits are stored in rttMonLatestJitterOperRTTSum.
    // The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperrttsumhigh interface{}

    // The sum of squares of RTT's that are successfully measured (high order 32
    // bits). The low order 32 bits are stored in rttMonLatestJitterOperRTTSum2.
    // The type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperrttsum2High interface{}

    // The sum of one way latency from source to destination (high order 32 bits).
    // The low order 32 bits are stored in rttMonLatestJitterOperOWSumSD. The type
    // is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowsumsdhigh interface{}

    // The sum of squares of one way latency from source to destination (high
    // order 32 bits). The low order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2SD. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteroperowsum2Sdhigh interface{}

    // The sum of one way latency from destination to source (high order 32 bits).
    // The low order 32 bits are stored  in rttMonLatestJitterOperOWSumDS. The
    // type is interface{} with range: 0..4294967295.
    Rttmonlatestjitteroperowsumdshigh interface{}

    // The sum of squares of one way latency from destination to source (high
    // order 32 bits). The low order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2DS. The type is interface{} with range:
    // 0..4294967295.
    Rttmonlatestjitteroperowsum2Dshigh interface{}
}

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetFilter() yfilter.YFilter { return rttmonlatestjitteroperentry.YFilter }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) SetFilter(yf yfilter.YFilter) { rttmonlatestjitteroperentry.YFilter = yf }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetGoName(yname string) string {
    if yname == "rttMonCtrlAdminIndex" { return "Rttmonctrladminindex" }
    if yname == "rttMonLatestJitterOperNumOfRTT" { return "Rttmonlatestjitteropernumofrtt" }
    if yname == "rttMonLatestJitterOperRTTSum" { return "Rttmonlatestjitteroperrttsum" }
    if yname == "rttMonLatestJitterOperRTTSum2" { return "Rttmonlatestjitteroperrttsum2" }
    if yname == "rttMonLatestJitterOperRTTMin" { return "Rttmonlatestjitteroperrttmin" }
    if yname == "rttMonLatestJitterOperRTTMax" { return "Rttmonlatestjitteroperrttmax" }
    if yname == "rttMonLatestJitterOperMinOfPositivesSD" { return "Rttmonlatestjitteroperminofpositivessd" }
    if yname == "rttMonLatestJitterOperMaxOfPositivesSD" { return "Rttmonlatestjitteropermaxofpositivessd" }
    if yname == "rttMonLatestJitterOperNumOfPositivesSD" { return "Rttmonlatestjitteropernumofpositivessd" }
    if yname == "rttMonLatestJitterOperSumOfPositivesSD" { return "Rttmonlatestjitteropersumofpositivessd" }
    if yname == "rttMonLatestJitterOperSum2PositivesSD" { return "Rttmonlatestjitteropersum2Positivessd" }
    if yname == "rttMonLatestJitterOperMinOfNegativesSD" { return "Rttmonlatestjitteroperminofnegativessd" }
    if yname == "rttMonLatestJitterOperMaxOfNegativesSD" { return "Rttmonlatestjitteropermaxofnegativessd" }
    if yname == "rttMonLatestJitterOperNumOfNegativesSD" { return "Rttmonlatestjitteropernumofnegativessd" }
    if yname == "rttMonLatestJitterOperSumOfNegativesSD" { return "Rttmonlatestjitteropersumofnegativessd" }
    if yname == "rttMonLatestJitterOperSum2NegativesSD" { return "Rttmonlatestjitteropersum2Negativessd" }
    if yname == "rttMonLatestJitterOperMinOfPositivesDS" { return "Rttmonlatestjitteroperminofpositivesds" }
    if yname == "rttMonLatestJitterOperMaxOfPositivesDS" { return "Rttmonlatestjitteropermaxofpositivesds" }
    if yname == "rttMonLatestJitterOperNumOfPositivesDS" { return "Rttmonlatestjitteropernumofpositivesds" }
    if yname == "rttMonLatestJitterOperSumOfPositivesDS" { return "Rttmonlatestjitteropersumofpositivesds" }
    if yname == "rttMonLatestJitterOperSum2PositivesDS" { return "Rttmonlatestjitteropersum2Positivesds" }
    if yname == "rttMonLatestJitterOperMinOfNegativesDS" { return "Rttmonlatestjitteroperminofnegativesds" }
    if yname == "rttMonLatestJitterOperMaxOfNegativesDS" { return "Rttmonlatestjitteropermaxofnegativesds" }
    if yname == "rttMonLatestJitterOperNumOfNegativesDS" { return "Rttmonlatestjitteropernumofnegativesds" }
    if yname == "rttMonLatestJitterOperSumOfNegativesDS" { return "Rttmonlatestjitteropersumofnegativesds" }
    if yname == "rttMonLatestJitterOperSum2NegativesDS" { return "Rttmonlatestjitteropersum2Negativesds" }
    if yname == "rttMonLatestJitterOperPacketLossSD" { return "Rttmonlatestjitteroperpacketlosssd" }
    if yname == "rttMonLatestJitterOperPacketLossDS" { return "Rttmonlatestjitteroperpacketlossds" }
    if yname == "rttMonLatestJitterOperPacketOutOfSequence" { return "Rttmonlatestjitteroperpacketoutofsequence" }
    if yname == "rttMonLatestJitterOperPacketMIA" { return "Rttmonlatestjitteroperpacketmia" }
    if yname == "rttMonLatestJitterOperPacketLateArrival" { return "Rttmonlatestjitteroperpacketlatearrival" }
    if yname == "rttMonLatestJitterOperSense" { return "Rttmonlatestjitteropersense" }
    if yname == "rttMonLatestJitterErrorSenseDescription" { return "Rttmonlatestjittererrorsensedescription" }
    if yname == "rttMonLatestJitterOperOWSumSD" { return "Rttmonlatestjitteroperowsumsd" }
    if yname == "rttMonLatestJitterOperOWSum2SD" { return "Rttmonlatestjitteroperowsum2Sd" }
    if yname == "rttMonLatestJitterOperOWMinSD" { return "Rttmonlatestjitteroperowminsd" }
    if yname == "rttMonLatestJitterOperOWMaxSD" { return "Rttmonlatestjitteroperowmaxsd" }
    if yname == "rttMonLatestJitterOperOWSumDS" { return "Rttmonlatestjitteroperowsumds" }
    if yname == "rttMonLatestJitterOperOWSum2DS" { return "Rttmonlatestjitteroperowsum2Ds" }
    if yname == "rttMonLatestJitterOperOWMinDS" { return "Rttmonlatestjitteroperowminds" }
    if yname == "rttMonLatestJitterOperOWMaxDS" { return "Rttmonlatestjitteroperowmaxds" }
    if yname == "rttMonLatestJitterOperNumOfOW" { return "Rttmonlatestjitteropernumofow" }
    if yname == "rttMonLatestJitterOperMOS" { return "Rttmonlatestjitteropermos" }
    if yname == "rttMonLatestJitterOperICPIF" { return "Rttmonlatestjitteropericpif" }
    if yname == "rttMonLatestJitterOperIAJOut" { return "Rttmonlatestjitteroperiajout" }
    if yname == "rttMonLatestJitterOperIAJIn" { return "Rttmonlatestjitteroperiajin" }
    if yname == "rttMonLatestJitterOperAvgJitter" { return "Rttmonlatestjitteroperavgjitter" }
    if yname == "rttMonLatestJitterOperAvgSDJ" { return "Rttmonlatestjitteroperavgsdj" }
    if yname == "rttMonLatestJitterOperAvgDSJ" { return "Rttmonlatestjitteroperavgdsj" }
    if yname == "rttMonLatestJitterOperOWAvgSD" { return "Rttmonlatestjitteroperowavgsd" }
    if yname == "rttMonLatestJitterOperOWAvgDS" { return "Rttmonlatestjitteroperowavgds" }
    if yname == "rttMonLatestJitterOperNTPState" { return "Rttmonlatestjitteroperntpstate" }
    if yname == "rttMonLatestJitterOperUnSyncRTs" { return "Rttmonlatestjitteroperunsyncrts" }
    if yname == "rttMonLatestJitterOperRTTSumHigh" { return "Rttmonlatestjitteroperrttsumhigh" }
    if yname == "rttMonLatestJitterOperRTTSum2High" { return "Rttmonlatestjitteroperrttsum2High" }
    if yname == "rttMonLatestJitterOperOWSumSDHigh" { return "Rttmonlatestjitteroperowsumsdhigh" }
    if yname == "rttMonLatestJitterOperOWSum2SDHigh" { return "Rttmonlatestjitteroperowsum2Sdhigh" }
    if yname == "rttMonLatestJitterOperOWSumDSHigh" { return "Rttmonlatestjitteroperowsumdshigh" }
    if yname == "rttMonLatestJitterOperOWSum2DSHigh" { return "Rttmonlatestjitteroperowsum2Dshigh" }
    return ""
}

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetSegmentPath() string {
    return "rttMonLatestJitterOperEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonlatestjitteroperentry.Rttmonctrladminindex) + "']"
}

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rttMonCtrlAdminIndex"] = rttmonlatestjitteroperentry.Rttmonctrladminindex
    leafs["rttMonLatestJitterOperNumOfRTT"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofrtt
    leafs["rttMonLatestJitterOperRTTSum"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsum
    leafs["rttMonLatestJitterOperRTTSum2"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsum2
    leafs["rttMonLatestJitterOperRTTMin"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttmin
    leafs["rttMonLatestJitterOperRTTMax"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttmax
    leafs["rttMonLatestJitterOperMinOfPositivesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofpositivessd
    leafs["rttMonLatestJitterOperMaxOfPositivesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofpositivessd
    leafs["rttMonLatestJitterOperNumOfPositivesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofpositivessd
    leafs["rttMonLatestJitterOperSumOfPositivesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofpositivessd
    leafs["rttMonLatestJitterOperSum2PositivesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Positivessd
    leafs["rttMonLatestJitterOperMinOfNegativesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofnegativessd
    leafs["rttMonLatestJitterOperMaxOfNegativesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofnegativessd
    leafs["rttMonLatestJitterOperNumOfNegativesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofnegativessd
    leafs["rttMonLatestJitterOperSumOfNegativesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofnegativessd
    leafs["rttMonLatestJitterOperSum2NegativesSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Negativessd
    leafs["rttMonLatestJitterOperMinOfPositivesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofpositivesds
    leafs["rttMonLatestJitterOperMaxOfPositivesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofpositivesds
    leafs["rttMonLatestJitterOperNumOfPositivesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofpositivesds
    leafs["rttMonLatestJitterOperSumOfPositivesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofpositivesds
    leafs["rttMonLatestJitterOperSum2PositivesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Positivesds
    leafs["rttMonLatestJitterOperMinOfNegativesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofnegativesds
    leafs["rttMonLatestJitterOperMaxOfNegativesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofnegativesds
    leafs["rttMonLatestJitterOperNumOfNegativesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofnegativesds
    leafs["rttMonLatestJitterOperSumOfNegativesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofnegativesds
    leafs["rttMonLatestJitterOperSum2NegativesDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Negativesds
    leafs["rttMonLatestJitterOperPacketLossSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketlosssd
    leafs["rttMonLatestJitterOperPacketLossDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketlossds
    leafs["rttMonLatestJitterOperPacketOutOfSequence"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketoutofsequence
    leafs["rttMonLatestJitterOperPacketMIA"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketmia
    leafs["rttMonLatestJitterOperPacketLateArrival"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketlatearrival
    leafs["rttMonLatestJitterOperSense"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropersense
    leafs["rttMonLatestJitterErrorSenseDescription"] = rttmonlatestjitteroperentry.Rttmonlatestjittererrorsensedescription
    leafs["rttMonLatestJitterOperOWSumSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumsd
    leafs["rttMonLatestJitterOperOWSum2SD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Sd
    leafs["rttMonLatestJitterOperOWMinSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowminsd
    leafs["rttMonLatestJitterOperOWMaxSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowmaxsd
    leafs["rttMonLatestJitterOperOWSumDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumds
    leafs["rttMonLatestJitterOperOWSum2DS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Ds
    leafs["rttMonLatestJitterOperOWMinDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowminds
    leafs["rttMonLatestJitterOperOWMaxDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowmaxds
    leafs["rttMonLatestJitterOperNumOfOW"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofow
    leafs["rttMonLatestJitterOperMOS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropermos
    leafs["rttMonLatestJitterOperICPIF"] = rttmonlatestjitteroperentry.Rttmonlatestjitteropericpif
    leafs["rttMonLatestJitterOperIAJOut"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperiajout
    leafs["rttMonLatestJitterOperIAJIn"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperiajin
    leafs["rttMonLatestJitterOperAvgJitter"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperavgjitter
    leafs["rttMonLatestJitterOperAvgSDJ"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperavgsdj
    leafs["rttMonLatestJitterOperAvgDSJ"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperavgdsj
    leafs["rttMonLatestJitterOperOWAvgSD"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowavgsd
    leafs["rttMonLatestJitterOperOWAvgDS"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowavgds
    leafs["rttMonLatestJitterOperNTPState"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperntpstate
    leafs["rttMonLatestJitterOperUnSyncRTs"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperunsyncrts
    leafs["rttMonLatestJitterOperRTTSumHigh"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsumhigh
    leafs["rttMonLatestJitterOperRTTSum2High"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsum2High
    leafs["rttMonLatestJitterOperOWSumSDHigh"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumsdhigh
    leafs["rttMonLatestJitterOperOWSum2SDHigh"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Sdhigh
    leafs["rttMonLatestJitterOperOWSumDSHigh"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumdshigh
    leafs["rttMonLatestJitterOperOWSum2DSHigh"] = rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Dshigh
    return leafs
}

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetBundleName() string { return "cisco_ios_xe" }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetYangName() string { return "rttMonLatestJitterOperEntry" }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) SetParent(parent types.Entity) { rttmonlatestjitteroperentry.parent = parent }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetParent() types.Entity { return rttmonlatestjitteroperentry.parent }

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetParentYangName() string { return "rttMonLatestJitterOperTable" }

// CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate represents on sender and responder is within configured tolerance level.
type CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate string

const (
    CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate_sync CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate = "sync"

    CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate_outOfSync CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate = "outOfSync"
)

