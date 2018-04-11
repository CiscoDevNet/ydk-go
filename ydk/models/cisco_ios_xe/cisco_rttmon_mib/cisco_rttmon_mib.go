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
    EntityData types.CommonEntityData
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

func (cISCORTTMONMIB *CISCORTTMONMIB) GetEntityData() *types.CommonEntityData {
    cISCORTTMONMIB.EntityData.YFilter = cISCORTTMONMIB.YFilter
    cISCORTTMONMIB.EntityData.YangName = "CISCO-RTTMON-MIB"
    cISCORTTMONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCORTTMONMIB.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    cISCORTTMONMIB.EntityData.SegmentPath = "CISCO-RTTMON-MIB:CISCO-RTTMON-MIB"
    cISCORTTMONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCORTTMONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCORTTMONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCORTTMONMIB.EntityData.Children = make(map[string]types.YChild)
    cISCORTTMONMIB.EntityData.Children["rttMonAppl"] = types.YChild{"Rttmonappl", &cISCORTTMONMIB.Rttmonappl}
    cISCORTTMONMIB.EntityData.Children["rttMonApplSupportedRttTypesTable"] = types.YChild{"Rttmonapplsupportedrtttypestable", &cISCORTTMONMIB.Rttmonapplsupportedrtttypestable}
    cISCORTTMONMIB.EntityData.Children["rttMonApplSupportedProtocolsTable"] = types.YChild{"Rttmonapplsupportedprotocolstable", &cISCORTTMONMIB.Rttmonapplsupportedprotocolstable}
    cISCORTTMONMIB.EntityData.Children["rttMonApplPreConfigedTable"] = types.YChild{"Rttmonapplpreconfigedtable", &cISCORTTMONMIB.Rttmonapplpreconfigedtable}
    cISCORTTMONMIB.EntityData.Children["rttMonApplAuthTable"] = types.YChild{"Rttmonapplauthtable", &cISCORTTMONMIB.Rttmonapplauthtable}
    cISCORTTMONMIB.EntityData.Children["rttMonCtrlAdminTable"] = types.YChild{"Rttmonctrladmintable", &cISCORTTMONMIB.Rttmonctrladmintable}
    cISCORTTMONMIB.EntityData.Children["rttMonEchoAdminTable"] = types.YChild{"Rttmonechoadmintable", &cISCORTTMONMIB.Rttmonechoadmintable}
    cISCORTTMONMIB.EntityData.Children["rttMonFileIOAdminTable"] = types.YChild{"Rttmonfileioadmintable", &cISCORTTMONMIB.Rttmonfileioadmintable}
    cISCORTTMONMIB.EntityData.Children["rttMonScriptAdminTable"] = types.YChild{"Rttmonscriptadmintable", &cISCORTTMONMIB.Rttmonscriptadmintable}
    cISCORTTMONMIB.EntityData.Children["rttMonReactTriggerAdminTable"] = types.YChild{"Rttmonreacttriggeradmintable", &cISCORTTMONMIB.Rttmonreacttriggeradmintable}
    cISCORTTMONMIB.EntityData.Children["rttMonEchoPathAdminTable"] = types.YChild{"Rttmonechopathadmintable", &cISCORTTMONMIB.Rttmonechopathadmintable}
    cISCORTTMONMIB.EntityData.Children["rttMonGrpScheduleAdminTable"] = types.YChild{"Rttmongrpscheduleadmintable", &cISCORTTMONMIB.Rttmongrpscheduleadmintable}
    cISCORTTMONMIB.EntityData.Children["rttMplsVpnMonCtrlTable"] = types.YChild{"Rttmplsvpnmonctrltable", &cISCORTTMONMIB.Rttmplsvpnmonctrltable}
    cISCORTTMONMIB.EntityData.Children["rttMonReactTable"] = types.YChild{"Rttmonreacttable", &cISCORTTMONMIB.Rttmonreacttable}
    cISCORTTMONMIB.EntityData.Children["rttMonGeneratedOperTable"] = types.YChild{"Rttmongeneratedopertable", &cISCORTTMONMIB.Rttmongeneratedopertable}
    cISCORTTMONMIB.EntityData.Children["rttMonStatsCaptureTable"] = types.YChild{"Rttmonstatscapturetable", &cISCORTTMONMIB.Rttmonstatscapturetable}
    cISCORTTMONMIB.EntityData.Children["rttMonStatsCollectTable"] = types.YChild{"Rttmonstatscollecttable", &cISCORTTMONMIB.Rttmonstatscollecttable}
    cISCORTTMONMIB.EntityData.Children["rttMonStatsTotalsTable"] = types.YChild{"Rttmonstatstotalstable", &cISCORTTMONMIB.Rttmonstatstotalstable}
    cISCORTTMONMIB.EntityData.Children["rttMonHTTPStatsTable"] = types.YChild{"Rttmonhttpstatstable", &cISCORTTMONMIB.Rttmonhttpstatstable}
    cISCORTTMONMIB.EntityData.Children["rttMonJitterStatsTable"] = types.YChild{"Rttmonjitterstatstable", &cISCORTTMONMIB.Rttmonjitterstatstable}
    cISCORTTMONMIB.EntityData.Children["rttMonLpdGrpStatsTable"] = types.YChild{"Rttmonlpdgrpstatstable", &cISCORTTMONMIB.Rttmonlpdgrpstatstable}
    cISCORTTMONMIB.EntityData.Children["rttMonHistoryCollectionTable"] = types.YChild{"Rttmonhistorycollectiontable", &cISCORTTMONMIB.Rttmonhistorycollectiontable}
    cISCORTTMONMIB.EntityData.Children["rttMonLatestHTTPOperTable"] = types.YChild{"Rttmonlatesthttpopertable", &cISCORTTMONMIB.Rttmonlatesthttpopertable}
    cISCORTTMONMIB.EntityData.Children["rttMonLatestJitterOperTable"] = types.YChild{"Rttmonlatestjitteropertable", &cISCORTTMONMIB.Rttmonlatestjitteropertable}
    cISCORTTMONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCORTTMONMIB.EntityData)
}

// CISCORTTMONMIB_Rttmonappl
type CISCORTTMONMIB_Rttmonappl struct {
    EntityData types.CommonEntityData
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

func (rttmonappl *CISCORTTMONMIB_Rttmonappl) GetEntityData() *types.CommonEntityData {
    rttmonappl.EntityData.YFilter = rttmonappl.YFilter
    rttmonappl.EntityData.YangName = "rttMonAppl"
    rttmonappl.EntityData.BundleName = "cisco_ios_xe"
    rttmonappl.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonappl.EntityData.SegmentPath = "rttMonAppl"
    rttmonappl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonappl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonappl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonappl.EntityData.Children = make(map[string]types.YChild)
    rttmonappl.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonappl.EntityData.Leafs["rttMonApplVersion"] = types.YLeaf{"Rttmonapplversion", rttmonappl.Rttmonapplversion}
    rttmonappl.EntityData.Leafs["rttMonApplMaxPacketDataSize"] = types.YLeaf{"Rttmonapplmaxpacketdatasize", rttmonappl.Rttmonapplmaxpacketdatasize}
    rttmonappl.EntityData.Leafs["rttMonApplTimeOfLastSet"] = types.YLeaf{"Rttmonappltimeoflastset", rttmonappl.Rttmonappltimeoflastset}
    rttmonappl.EntityData.Leafs["rttMonApplNumCtrlAdminEntry"] = types.YLeaf{"Rttmonapplnumctrladminentry", rttmonappl.Rttmonapplnumctrladminentry}
    rttmonappl.EntityData.Leafs["rttMonApplReset"] = types.YLeaf{"Rttmonapplreset", rttmonappl.Rttmonapplreset}
    rttmonappl.EntityData.Leafs["rttMonApplPreConfigedReset"] = types.YLeaf{"Rttmonapplpreconfigedreset", rttmonappl.Rttmonapplpreconfigedreset}
    rttmonappl.EntityData.Leafs["rttMonApplProbeCapacity"] = types.YLeaf{"Rttmonapplprobecapacity", rttmonappl.Rttmonapplprobecapacity}
    rttmonappl.EntityData.Leafs["rttMonApplFreeMemLowWaterMark"] = types.YLeaf{"Rttmonapplfreememlowwatermark", rttmonappl.Rttmonapplfreememlowwatermark}
    rttmonappl.EntityData.Leafs["rttMonApplLatestSetError"] = types.YLeaf{"Rttmonappllatestseterror", rttmonappl.Rttmonappllatestseterror}
    rttmonappl.EntityData.Leafs["rttMonApplResponder"] = types.YLeaf{"Rttmonapplresponder", rttmonappl.Rttmonapplresponder}
    rttmonappl.EntityData.Leafs["rttMonApplLpdGrpStatsReset"] = types.YLeaf{"Rttmonappllpdgrpstatsreset", rttmonappl.Rttmonappllpdgrpstatsreset}
    return &(rttmonappl.EntityData)
}

// CISCORTTMONMIB_Rttmonapplsupportedrtttypestable
// A table of which contains the supported Rtt
// Monitor Types.
// 
// See the RttMonRttType textual convention for
// the definition of each type.
type CISCORTTMONMIB_Rttmonapplsupportedrtttypestable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list that presents the valid Rtt Monitor Types. The type is slice of
    // CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry.
    Rttmonapplsupportedrtttypesentry []CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry
}

func (rttmonapplsupportedrtttypestable *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable) GetEntityData() *types.CommonEntityData {
    rttmonapplsupportedrtttypestable.EntityData.YFilter = rttmonapplsupportedrtttypestable.YFilter
    rttmonapplsupportedrtttypestable.EntityData.YangName = "rttMonApplSupportedRttTypesTable"
    rttmonapplsupportedrtttypestable.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplsupportedrtttypestable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonapplsupportedrtttypestable.EntityData.SegmentPath = "rttMonApplSupportedRttTypesTable"
    rttmonapplsupportedrtttypestable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplsupportedrtttypestable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplsupportedrtttypestable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplsupportedrtttypestable.EntityData.Children = make(map[string]types.YChild)
    rttmonapplsupportedrtttypestable.EntityData.Children["rttMonApplSupportedRttTypesEntry"] = types.YChild{"Rttmonapplsupportedrtttypesentry", nil}
    for i := range rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry {
        rttmonapplsupportedrtttypestable.EntityData.Children[types.GetSegmentPath(&rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry[i])] = types.YChild{"Rttmonapplsupportedrtttypesentry", &rttmonapplsupportedrtttypestable.Rttmonapplsupportedrtttypesentry[i]}
    }
    rttmonapplsupportedrtttypestable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonapplsupportedrtttypestable.EntityData)
}

// CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry
// A list that presents the valid Rtt Monitor
// Types.
type CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object indexes the supported 'RttMonRttType'
    // types. The type is RttMonRttType.
    Rttmonapplsupportedrtttypes interface{}

    // This object defines the supported 'RttMonRttType' types. The type is bool.
    Rttmonapplsupportedrtttypesvalid interface{}
}

func (rttmonapplsupportedrtttypesentry *CISCORTTMONMIB_Rttmonapplsupportedrtttypestable_Rttmonapplsupportedrtttypesentry) GetEntityData() *types.CommonEntityData {
    rttmonapplsupportedrtttypesentry.EntityData.YFilter = rttmonapplsupportedrtttypesentry.YFilter
    rttmonapplsupportedrtttypesentry.EntityData.YangName = "rttMonApplSupportedRttTypesEntry"
    rttmonapplsupportedrtttypesentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplsupportedrtttypesentry.EntityData.ParentYangName = "rttMonApplSupportedRttTypesTable"
    rttmonapplsupportedrtttypesentry.EntityData.SegmentPath = "rttMonApplSupportedRttTypesEntry" + "[rttMonApplSupportedRttTypes='" + fmt.Sprintf("%v", rttmonapplsupportedrtttypesentry.Rttmonapplsupportedrtttypes) + "']"
    rttmonapplsupportedrtttypesentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplsupportedrtttypesentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplsupportedrtttypesentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplsupportedrtttypesentry.EntityData.Children = make(map[string]types.YChild)
    rttmonapplsupportedrtttypesentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonapplsupportedrtttypesentry.EntityData.Leafs["rttMonApplSupportedRttTypes"] = types.YLeaf{"Rttmonapplsupportedrtttypes", rttmonapplsupportedrtttypesentry.Rttmonapplsupportedrtttypes}
    rttmonapplsupportedrtttypesentry.EntityData.Leafs["rttMonApplSupportedRttTypesValid"] = types.YLeaf{"Rttmonapplsupportedrtttypesvalid", rttmonapplsupportedrtttypesentry.Rttmonapplsupportedrtttypesvalid}
    return &(rttmonapplsupportedrtttypesentry.EntityData)
}

// CISCORTTMONMIB_Rttmonapplsupportedprotocolstable
// A table of which contains the supported Rtt
// Monitor Protocols.
// 
// See the RttMonProtocol textual convention 
// for the definition of each protocol.
type CISCORTTMONMIB_Rttmonapplsupportedprotocolstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list that presents the valid Rtt Monitor Protocols. The type is slice of
    // CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry.
    Rttmonapplsupportedprotocolsentry []CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry
}

func (rttmonapplsupportedprotocolstable *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable) GetEntityData() *types.CommonEntityData {
    rttmonapplsupportedprotocolstable.EntityData.YFilter = rttmonapplsupportedprotocolstable.YFilter
    rttmonapplsupportedprotocolstable.EntityData.YangName = "rttMonApplSupportedProtocolsTable"
    rttmonapplsupportedprotocolstable.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplsupportedprotocolstable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonapplsupportedprotocolstable.EntityData.SegmentPath = "rttMonApplSupportedProtocolsTable"
    rttmonapplsupportedprotocolstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplsupportedprotocolstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplsupportedprotocolstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplsupportedprotocolstable.EntityData.Children = make(map[string]types.YChild)
    rttmonapplsupportedprotocolstable.EntityData.Children["rttMonApplSupportedProtocolsEntry"] = types.YChild{"Rttmonapplsupportedprotocolsentry", nil}
    for i := range rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry {
        rttmonapplsupportedprotocolstable.EntityData.Children[types.GetSegmentPath(&rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry[i])] = types.YChild{"Rttmonapplsupportedprotocolsentry", &rttmonapplsupportedprotocolstable.Rttmonapplsupportedprotocolsentry[i]}
    }
    rttmonapplsupportedprotocolstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonapplsupportedprotocolstable.EntityData)
}

// CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry
// A list that presents the valid Rtt Monitor
// Protocols.
type CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object indexes the supported 'RttMonProtocol'
    // protocols. The type is RttMonProtocol.
    Rttmonapplsupportedprotocols interface{}

    // This object defines the supported 'RttMonProtocol' protocols. The type is
    // bool.
    Rttmonapplsupportedprotocolsvalid interface{}
}

func (rttmonapplsupportedprotocolsentry *CISCORTTMONMIB_Rttmonapplsupportedprotocolstable_Rttmonapplsupportedprotocolsentry) GetEntityData() *types.CommonEntityData {
    rttmonapplsupportedprotocolsentry.EntityData.YFilter = rttmonapplsupportedprotocolsentry.YFilter
    rttmonapplsupportedprotocolsentry.EntityData.YangName = "rttMonApplSupportedProtocolsEntry"
    rttmonapplsupportedprotocolsentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplsupportedprotocolsentry.EntityData.ParentYangName = "rttMonApplSupportedProtocolsTable"
    rttmonapplsupportedprotocolsentry.EntityData.SegmentPath = "rttMonApplSupportedProtocolsEntry" + "[rttMonApplSupportedProtocols='" + fmt.Sprintf("%v", rttmonapplsupportedprotocolsentry.Rttmonapplsupportedprotocols) + "']"
    rttmonapplsupportedprotocolsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplsupportedprotocolsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplsupportedprotocolsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplsupportedprotocolsentry.EntityData.Children = make(map[string]types.YChild)
    rttmonapplsupportedprotocolsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonapplsupportedprotocolsentry.EntityData.Leafs["rttMonApplSupportedProtocols"] = types.YLeaf{"Rttmonapplsupportedprotocols", rttmonapplsupportedprotocolsentry.Rttmonapplsupportedprotocols}
    rttmonapplsupportedprotocolsentry.EntityData.Leafs["rttMonApplSupportedProtocolsValid"] = types.YLeaf{"Rttmonapplsupportedprotocolsvalid", rttmonapplsupportedprotocolsentry.Rttmonapplsupportedprotocolsvalid}
    return &(rttmonapplsupportedprotocolsentry.EntityData)
}

// CISCORTTMONMIB_Rttmonapplpreconfigedtable
// A table of which contains the previously
// configured Script Names and File IO targets.
// 
// These Script Names and File IO targets are installed
// via a different mechanism than this application, and
// are specific to each platform.
type CISCORTTMONMIB_Rttmonapplpreconfigedtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that describe the previously configured Script Names and
    // File IO targets. The type is slice of
    // CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry.
    Rttmonapplpreconfigedentry []CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry
}

func (rttmonapplpreconfigedtable *CISCORTTMONMIB_Rttmonapplpreconfigedtable) GetEntityData() *types.CommonEntityData {
    rttmonapplpreconfigedtable.EntityData.YFilter = rttmonapplpreconfigedtable.YFilter
    rttmonapplpreconfigedtable.EntityData.YangName = "rttMonApplPreConfigedTable"
    rttmonapplpreconfigedtable.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplpreconfigedtable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonapplpreconfigedtable.EntityData.SegmentPath = "rttMonApplPreConfigedTable"
    rttmonapplpreconfigedtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplpreconfigedtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplpreconfigedtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplpreconfigedtable.EntityData.Children = make(map[string]types.YChild)
    rttmonapplpreconfigedtable.EntityData.Children["rttMonApplPreConfigedEntry"] = types.YChild{"Rttmonapplpreconfigedentry", nil}
    for i := range rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry {
        rttmonapplpreconfigedtable.EntityData.Children[types.GetSegmentPath(&rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry[i])] = types.YChild{"Rttmonapplpreconfigedentry", &rttmonapplpreconfigedtable.Rttmonapplpreconfigedentry[i]}
    }
    rttmonapplpreconfigedtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonapplpreconfigedtable.EntityData)
}

// CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry
// A list of objects that describe the previously
// configured Script Names and File IO targets.
type CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry struct {
    EntityData types.CommonEntityData
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

func (rttmonapplpreconfigedentry *CISCORTTMONMIB_Rttmonapplpreconfigedtable_Rttmonapplpreconfigedentry) GetEntityData() *types.CommonEntityData {
    rttmonapplpreconfigedentry.EntityData.YFilter = rttmonapplpreconfigedentry.YFilter
    rttmonapplpreconfigedentry.EntityData.YangName = "rttMonApplPreConfigedEntry"
    rttmonapplpreconfigedentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplpreconfigedentry.EntityData.ParentYangName = "rttMonApplPreConfigedTable"
    rttmonapplpreconfigedentry.EntityData.SegmentPath = "rttMonApplPreConfigedEntry" + "[rttMonApplPreConfigedType='" + fmt.Sprintf("%v", rttmonapplpreconfigedentry.Rttmonapplpreconfigedtype) + "']" + "[rttMonApplPreConfigedName='" + fmt.Sprintf("%v", rttmonapplpreconfigedentry.Rttmonapplpreconfigedname) + "']"
    rttmonapplpreconfigedentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplpreconfigedentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplpreconfigedentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplpreconfigedentry.EntityData.Children = make(map[string]types.YChild)
    rttmonapplpreconfigedentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonapplpreconfigedentry.EntityData.Leafs["rttMonApplPreConfigedType"] = types.YLeaf{"Rttmonapplpreconfigedtype", rttmonapplpreconfigedentry.Rttmonapplpreconfigedtype}
    rttmonapplpreconfigedentry.EntityData.Leafs["rttMonApplPreConfigedName"] = types.YLeaf{"Rttmonapplpreconfigedname", rttmonapplpreconfigedentry.Rttmonapplpreconfigedname}
    rttmonapplpreconfigedentry.EntityData.Leafs["rttMonApplPreConfigedValid"] = types.YLeaf{"Rttmonapplpreconfigedvalid", rttmonapplpreconfigedentry.Rttmonapplpreconfigedvalid}
    return &(rttmonapplpreconfigedentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list that presents the valid parameters for Authenticating RTR Control
    // Protocol. The type is slice of
    // CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry.
    Rttmonapplauthentry []CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry
}

func (rttmonapplauthtable *CISCORTTMONMIB_Rttmonapplauthtable) GetEntityData() *types.CommonEntityData {
    rttmonapplauthtable.EntityData.YFilter = rttmonapplauthtable.YFilter
    rttmonapplauthtable.EntityData.YangName = "rttMonApplAuthTable"
    rttmonapplauthtable.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplauthtable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonapplauthtable.EntityData.SegmentPath = "rttMonApplAuthTable"
    rttmonapplauthtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplauthtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplauthtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplauthtable.EntityData.Children = make(map[string]types.YChild)
    rttmonapplauthtable.EntityData.Children["rttMonApplAuthEntry"] = types.YChild{"Rttmonapplauthentry", nil}
    for i := range rttmonapplauthtable.Rttmonapplauthentry {
        rttmonapplauthtable.EntityData.Children[types.GetSegmentPath(&rttmonapplauthtable.Rttmonapplauthentry[i])] = types.YChild{"Rttmonapplauthentry", &rttmonapplauthtable.Rttmonapplauthentry[i]}
    }
    rttmonapplauthtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonapplauthtable.EntityData)
}

// CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry
// A list that presents the valid parameters for Authenticating
// RTR Control Protocol.
type CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry struct {
    EntityData types.CommonEntityData
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

func (rttmonapplauthentry *CISCORTTMONMIB_Rttmonapplauthtable_Rttmonapplauthentry) GetEntityData() *types.CommonEntityData {
    rttmonapplauthentry.EntityData.YFilter = rttmonapplauthentry.YFilter
    rttmonapplauthentry.EntityData.YangName = "rttMonApplAuthEntry"
    rttmonapplauthentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonapplauthentry.EntityData.ParentYangName = "rttMonApplAuthTable"
    rttmonapplauthentry.EntityData.SegmentPath = "rttMonApplAuthEntry" + "[rttMonApplAuthIndex='" + fmt.Sprintf("%v", rttmonapplauthentry.Rttmonapplauthindex) + "']"
    rttmonapplauthentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonapplauthentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonapplauthentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonapplauthentry.EntityData.Children = make(map[string]types.YChild)
    rttmonapplauthentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthIndex"] = types.YLeaf{"Rttmonapplauthindex", rttmonapplauthentry.Rttmonapplauthindex}
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthKeyChain"] = types.YLeaf{"Rttmonapplauthkeychain", rttmonapplauthentry.Rttmonapplauthkeychain}
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthKeyString1"] = types.YLeaf{"Rttmonapplauthkeystring1", rttmonapplauthentry.Rttmonapplauthkeystring1}
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthKeyString2"] = types.YLeaf{"Rttmonapplauthkeystring2", rttmonapplauthentry.Rttmonapplauthkeystring2}
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthKeyString3"] = types.YLeaf{"Rttmonapplauthkeystring3", rttmonapplauthentry.Rttmonapplauthkeystring3}
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthKeyString4"] = types.YLeaf{"Rttmonapplauthkeystring4", rttmonapplauthentry.Rttmonapplauthkeystring4}
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthKeyString5"] = types.YLeaf{"Rttmonapplauthkeystring5", rttmonapplauthentry.Rttmonapplauthkeystring5}
    rttmonapplauthentry.EntityData.Leafs["rttMonApplAuthStatus"] = types.YLeaf{"Rttmonapplauthstatus", rttmonapplauthentry.Rttmonapplauthstatus}
    return &(rttmonapplauthentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual RTT control row. The type
    // is slice of CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry.
    Rttmonctrladminentry []CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry
}

func (rttmonctrladmintable *CISCORTTMONMIB_Rttmonctrladmintable) GetEntityData() *types.CommonEntityData {
    rttmonctrladmintable.EntityData.YFilter = rttmonctrladmintable.YFilter
    rttmonctrladmintable.EntityData.YangName = "rttMonCtrlAdminTable"
    rttmonctrladmintable.EntityData.BundleName = "cisco_ios_xe"
    rttmonctrladmintable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonctrladmintable.EntityData.SegmentPath = "rttMonCtrlAdminTable"
    rttmonctrladmintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonctrladmintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonctrladmintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonctrladmintable.EntityData.Children = make(map[string]types.YChild)
    rttmonctrladmintable.EntityData.Children["rttMonCtrlAdminEntry"] = types.YChild{"Rttmonctrladminentry", nil}
    for i := range rttmonctrladmintable.Rttmonctrladminentry {
        rttmonctrladmintable.EntityData.Children[types.GetSegmentPath(&rttmonctrladmintable.Rttmonctrladminentry[i])] = types.YChild{"Rttmonctrladminentry", &rttmonctrladmintable.Rttmonctrladminentry[i]}
    }
    rttmonctrladmintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonctrladmintable.EntityData)
}

// CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry
// A base list of objects that define a conceptual RTT
// control row.
type CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry struct {
    EntityData types.CommonEntityData
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

func (rttmonctrladminentry *CISCORTTMONMIB_Rttmonctrladmintable_Rttmonctrladminentry) GetEntityData() *types.CommonEntityData {
    rttmonctrladminentry.EntityData.YFilter = rttmonctrladminentry.YFilter
    rttmonctrladminentry.EntityData.YangName = "rttMonCtrlAdminEntry"
    rttmonctrladminentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonctrladminentry.EntityData.ParentYangName = "rttMonCtrlAdminTable"
    rttmonctrladminentry.EntityData.SegmentPath = "rttMonCtrlAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonctrladminentry.Rttmonctrladminindex) + "']"
    rttmonctrladminentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonctrladminentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonctrladminentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonctrladminentry.EntityData.Children = make(map[string]types.YChild)
    rttmonctrladminentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonctrladminentry.Rttmonctrladminindex}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminOwner"] = types.YLeaf{"Rttmonctrladminowner", rttmonctrladminentry.Rttmonctrladminowner}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminTag"] = types.YLeaf{"Rttmonctrladmintag", rttmonctrladminentry.Rttmonctrladmintag}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminRttType"] = types.YLeaf{"Rttmonctrladminrtttype", rttmonctrladminentry.Rttmonctrladminrtttype}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminThreshold"] = types.YLeaf{"Rttmonctrladminthreshold", rttmonctrladminentry.Rttmonctrladminthreshold}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminFrequency"] = types.YLeaf{"Rttmonctrladminfrequency", rttmonctrladminentry.Rttmonctrladminfrequency}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminTimeout"] = types.YLeaf{"Rttmonctrladmintimeout", rttmonctrladminentry.Rttmonctrladmintimeout}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminVerifyData"] = types.YLeaf{"Rttmonctrladminverifydata", rttmonctrladminentry.Rttmonctrladminverifydata}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminStatus"] = types.YLeaf{"Rttmonctrladminstatus", rttmonctrladminentry.Rttmonctrladminstatus}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminNvgen"] = types.YLeaf{"Rttmonctrladminnvgen", rttmonctrladminentry.Rttmonctrladminnvgen}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlAdminGroupName"] = types.YLeaf{"Rttmonctrladmingroupname", rttmonctrladminentry.Rttmonctrladmingroupname}
    rttmonctrladminentry.EntityData.Leafs["rttMonScheduleAdminRttLife"] = types.YLeaf{"Rttmonscheduleadminrttlife", rttmonctrladminentry.Rttmonscheduleadminrttlife}
    rttmonctrladminentry.EntityData.Leafs["rttMonScheduleAdminRttStartTime"] = types.YLeaf{"Rttmonscheduleadminrttstarttime", rttmonctrladminentry.Rttmonscheduleadminrttstarttime}
    rttmonctrladminentry.EntityData.Leafs["rttMonScheduleAdminConceptRowAgeout"] = types.YLeaf{"Rttmonscheduleadminconceptrowageout", rttmonctrladminentry.Rttmonscheduleadminconceptrowageout}
    rttmonctrladminentry.EntityData.Leafs["rttMonScheduleAdminRttRecurring"] = types.YLeaf{"Rttmonscheduleadminrttrecurring", rttmonctrladminentry.Rttmonscheduleadminrttrecurring}
    rttmonctrladminentry.EntityData.Leafs["rttMonScheduleAdminConceptRowAgeoutV2"] = types.YLeaf{"Rttmonscheduleadminconceptrowageoutv2", rttmonctrladminentry.Rttmonscheduleadminconceptrowageoutv2}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminConnectionEnable"] = types.YLeaf{"Rttmonreactadminconnectionenable", rttmonctrladminentry.Rttmonreactadminconnectionenable}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminTimeoutEnable"] = types.YLeaf{"Rttmonreactadmintimeoutenable", rttmonctrladminentry.Rttmonreactadmintimeoutenable}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminThresholdType"] = types.YLeaf{"Rttmonreactadminthresholdtype", rttmonctrladminentry.Rttmonreactadminthresholdtype}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminThresholdFalling"] = types.YLeaf{"Rttmonreactadminthresholdfalling", rttmonctrladminentry.Rttmonreactadminthresholdfalling}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminThresholdCount"] = types.YLeaf{"Rttmonreactadminthresholdcount", rttmonctrladminentry.Rttmonreactadminthresholdcount}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminThresholdCount2"] = types.YLeaf{"Rttmonreactadminthresholdcount2", rttmonctrladminentry.Rttmonreactadminthresholdcount2}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminActionType"] = types.YLeaf{"Rttmonreactadminactiontype", rttmonctrladminentry.Rttmonreactadminactiontype}
    rttmonctrladminentry.EntityData.Leafs["rttMonReactAdminVerifyErrorEnable"] = types.YLeaf{"Rttmonreactadminverifyerrorenable", rttmonctrladminentry.Rttmonreactadminverifyerrorenable}
    rttmonctrladminentry.EntityData.Leafs["rttMonStatisticsAdminNumHourGroups"] = types.YLeaf{"Rttmonstatisticsadminnumhourgroups", rttmonctrladminentry.Rttmonstatisticsadminnumhourgroups}
    rttmonctrladminentry.EntityData.Leafs["rttMonStatisticsAdminNumPaths"] = types.YLeaf{"Rttmonstatisticsadminnumpaths", rttmonctrladminentry.Rttmonstatisticsadminnumpaths}
    rttmonctrladminentry.EntityData.Leafs["rttMonStatisticsAdminNumHops"] = types.YLeaf{"Rttmonstatisticsadminnumhops", rttmonctrladminentry.Rttmonstatisticsadminnumhops}
    rttmonctrladminentry.EntityData.Leafs["rttMonStatisticsAdminNumDistBuckets"] = types.YLeaf{"Rttmonstatisticsadminnumdistbuckets", rttmonctrladminentry.Rttmonstatisticsadminnumdistbuckets}
    rttmonctrladminentry.EntityData.Leafs["rttMonStatisticsAdminDistInterval"] = types.YLeaf{"Rttmonstatisticsadmindistinterval", rttmonctrladminentry.Rttmonstatisticsadmindistinterval}
    rttmonctrladminentry.EntityData.Leafs["rttMonHistoryAdminNumLives"] = types.YLeaf{"Rttmonhistoryadminnumlives", rttmonctrladminentry.Rttmonhistoryadminnumlives}
    rttmonctrladminentry.EntityData.Leafs["rttMonHistoryAdminNumBuckets"] = types.YLeaf{"Rttmonhistoryadminnumbuckets", rttmonctrladminentry.Rttmonhistoryadminnumbuckets}
    rttmonctrladminentry.EntityData.Leafs["rttMonHistoryAdminNumSamples"] = types.YLeaf{"Rttmonhistoryadminnumsamples", rttmonctrladminentry.Rttmonhistoryadminnumsamples}
    rttmonctrladminentry.EntityData.Leafs["rttMonHistoryAdminFilter"] = types.YLeaf{"Rttmonhistoryadminfilter", rttmonctrladminentry.Rttmonhistoryadminfilter}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperModificationTime"] = types.YLeaf{"Rttmonctrlopermodificationtime", rttmonctrladminentry.Rttmonctrlopermodificationtime}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperDiagText"] = types.YLeaf{"Rttmonctrloperdiagtext", rttmonctrladminentry.Rttmonctrloperdiagtext}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperResetTime"] = types.YLeaf{"Rttmonctrloperresettime", rttmonctrladminentry.Rttmonctrloperresettime}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperOctetsInUse"] = types.YLeaf{"Rttmonctrloperoctetsinuse", rttmonctrladminentry.Rttmonctrloperoctetsinuse}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperConnectionLostOccurred"] = types.YLeaf{"Rttmonctrloperconnectionlostoccurred", rttmonctrladminentry.Rttmonctrloperconnectionlostoccurred}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperTimeoutOccurred"] = types.YLeaf{"Rttmonctrlopertimeoutoccurred", rttmonctrladminentry.Rttmonctrlopertimeoutoccurred}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperOverThresholdOccurred"] = types.YLeaf{"Rttmonctrloperoverthresholdoccurred", rttmonctrladminentry.Rttmonctrloperoverthresholdoccurred}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperNumRtts"] = types.YLeaf{"Rttmonctrlopernumrtts", rttmonctrladminentry.Rttmonctrlopernumrtts}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperRttLife"] = types.YLeaf{"Rttmonctrloperrttlife", rttmonctrladminentry.Rttmonctrloperrttlife}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperState"] = types.YLeaf{"Rttmonctrloperstate", rttmonctrladminentry.Rttmonctrloperstate}
    rttmonctrladminentry.EntityData.Leafs["rttMonCtrlOperVerifyErrorOccurred"] = types.YLeaf{"Rttmonctrloperverifyerroroccurred", rttmonctrladminentry.Rttmonctrloperverifyerroroccurred}
    rttmonctrladminentry.EntityData.Leafs["rttMonLatestRttOperCompletionTime"] = types.YLeaf{"Rttmonlatestrttopercompletiontime", rttmonctrladminentry.Rttmonlatestrttopercompletiontime}
    rttmonctrladminentry.EntityData.Leafs["rttMonLatestRttOperSense"] = types.YLeaf{"Rttmonlatestrttopersense", rttmonctrladminentry.Rttmonlatestrttopersense}
    rttmonctrladminentry.EntityData.Leafs["rttMonLatestRttOperApplSpecificSense"] = types.YLeaf{"Rttmonlatestrttoperapplspecificsense", rttmonctrladminentry.Rttmonlatestrttoperapplspecificsense}
    rttmonctrladminentry.EntityData.Leafs["rttMonLatestRttOperSenseDescription"] = types.YLeaf{"Rttmonlatestrttopersensedescription", rttmonctrladminentry.Rttmonlatestrttopersensedescription}
    rttmonctrladminentry.EntityData.Leafs["rttMonLatestRttOperTime"] = types.YLeaf{"Rttmonlatestrttopertime", rttmonctrladminentry.Rttmonlatestrttopertime}
    rttmonctrladminentry.EntityData.Leafs["rttMonLatestRttOperAddress"] = types.YLeaf{"Rttmonlatestrttoperaddress", rttmonctrladminentry.Rttmonlatestrttoperaddress}
    return &(rttmonctrladminentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for RttMonRttType
    // conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry.
    Rttmonechoadminentry []CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry
}

func (rttmonechoadmintable *CISCORTTMONMIB_Rttmonechoadmintable) GetEntityData() *types.CommonEntityData {
    rttmonechoadmintable.EntityData.YFilter = rttmonechoadmintable.YFilter
    rttmonechoadmintable.EntityData.YangName = "rttMonEchoAdminTable"
    rttmonechoadmintable.EntityData.BundleName = "cisco_ios_xe"
    rttmonechoadmintable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonechoadmintable.EntityData.SegmentPath = "rttMonEchoAdminTable"
    rttmonechoadmintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonechoadmintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonechoadmintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonechoadmintable.EntityData.Children = make(map[string]types.YChild)
    rttmonechoadmintable.EntityData.Children["rttMonEchoAdminEntry"] = types.YChild{"Rttmonechoadminentry", nil}
    for i := range rttmonechoadmintable.Rttmonechoadminentry {
        rttmonechoadmintable.EntityData.Children[types.GetSegmentPath(&rttmonechoadmintable.Rttmonechoadminentry[i])] = types.YChild{"Rttmonechoadminentry", &rttmonechoadmintable.Rttmonechoadminentry[i]}
    }
    rttmonechoadmintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonechoadmintable.EntityData)
}

// CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry
// A list of objects that define specific configuration for
// RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry struct {
    EntityData types.CommonEntityData
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
    // string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Rttmonechoadmintargetmacaddress interface{}

    // This object indicates the MAC address of the source device. This object is
    // only applicable for Y.1731 operations.  rttMonEchoAdminSourceMacAddress and
    // rttMonEchoAdminSourceMPID may not be used in conjunction. The type is
    // string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

func (rttmonechoadminentry *CISCORTTMONMIB_Rttmonechoadmintable_Rttmonechoadminentry) GetEntityData() *types.CommonEntityData {
    rttmonechoadminentry.EntityData.YFilter = rttmonechoadminentry.YFilter
    rttmonechoadminentry.EntityData.YangName = "rttMonEchoAdminEntry"
    rttmonechoadminentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonechoadminentry.EntityData.ParentYangName = "rttMonEchoAdminTable"
    rttmonechoadminentry.EntityData.SegmentPath = "rttMonEchoAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonechoadminentry.Rttmonctrladminindex) + "']"
    rttmonechoadminentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonechoadminentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonechoadminentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonechoadminentry.EntityData.Children = make(map[string]types.YChild)
    rttmonechoadminentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonechoadminentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonechoadminentry.Rttmonctrladminindex}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminProtocol"] = types.YLeaf{"Rttmonechoadminprotocol", rttmonechoadminentry.Rttmonechoadminprotocol}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetAddress"] = types.YLeaf{"Rttmonechoadmintargetaddress", rttmonechoadminentry.Rttmonechoadmintargetaddress}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminPktDataRequestSize"] = types.YLeaf{"Rttmonechoadminpktdatarequestsize", rttmonechoadminentry.Rttmonechoadminpktdatarequestsize}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminPktDataResponseSize"] = types.YLeaf{"Rttmonechoadminpktdataresponsesize", rttmonechoadminentry.Rttmonechoadminpktdataresponsesize}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetPort"] = types.YLeaf{"Rttmonechoadmintargetport", rttmonechoadminentry.Rttmonechoadmintargetport}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminSourceAddress"] = types.YLeaf{"Rttmonechoadminsourceaddress", rttmonechoadminentry.Rttmonechoadminsourceaddress}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminSourcePort"] = types.YLeaf{"Rttmonechoadminsourceport", rttmonechoadminentry.Rttmonechoadminsourceport}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminControlEnable"] = types.YLeaf{"Rttmonechoadmincontrolenable", rttmonechoadminentry.Rttmonechoadmincontrolenable}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTOS"] = types.YLeaf{"Rttmonechoadmintos", rttmonechoadminentry.Rttmonechoadmintos}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSREnable"] = types.YLeaf{"Rttmonechoadminlsrenable", rttmonechoadminentry.Rttmonechoadminlsrenable}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetAddressString"] = types.YLeaf{"Rttmonechoadmintargetaddressstring", rttmonechoadminentry.Rttmonechoadmintargetaddressstring}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminNameServer"] = types.YLeaf{"Rttmonechoadminnameserver", rttmonechoadminentry.Rttmonechoadminnameserver}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminOperation"] = types.YLeaf{"Rttmonechoadminoperation", rttmonechoadminentry.Rttmonechoadminoperation}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminHTTPVersion"] = types.YLeaf{"Rttmonechoadminhttpversion", rttmonechoadminentry.Rttmonechoadminhttpversion}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminURL"] = types.YLeaf{"Rttmonechoadminurl", rttmonechoadminentry.Rttmonechoadminurl}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminCache"] = types.YLeaf{"Rttmonechoadmincache", rttmonechoadminentry.Rttmonechoadmincache}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminInterval"] = types.YLeaf{"Rttmonechoadmininterval", rttmonechoadminentry.Rttmonechoadmininterval}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminNumPackets"] = types.YLeaf{"Rttmonechoadminnumpackets", rttmonechoadminentry.Rttmonechoadminnumpackets}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminProxy"] = types.YLeaf{"Rttmonechoadminproxy", rttmonechoadminentry.Rttmonechoadminproxy}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminString1"] = types.YLeaf{"Rttmonechoadminstring1", rttmonechoadminentry.Rttmonechoadminstring1}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminString2"] = types.YLeaf{"Rttmonechoadminstring2", rttmonechoadminentry.Rttmonechoadminstring2}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminString3"] = types.YLeaf{"Rttmonechoadminstring3", rttmonechoadminentry.Rttmonechoadminstring3}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminString4"] = types.YLeaf{"Rttmonechoadminstring4", rttmonechoadminentry.Rttmonechoadminstring4}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminString5"] = types.YLeaf{"Rttmonechoadminstring5", rttmonechoadminentry.Rttmonechoadminstring5}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminMode"] = types.YLeaf{"Rttmonechoadminmode", rttmonechoadminentry.Rttmonechoadminmode}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminVrfName"] = types.YLeaf{"Rttmonechoadminvrfname", rttmonechoadminentry.Rttmonechoadminvrfname}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminCodecType"] = types.YLeaf{"Rttmonechoadmincodectype", rttmonechoadminentry.Rttmonechoadmincodectype}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminCodecInterval"] = types.YLeaf{"Rttmonechoadmincodecinterval", rttmonechoadminentry.Rttmonechoadmincodecinterval}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminCodecPayload"] = types.YLeaf{"Rttmonechoadmincodecpayload", rttmonechoadminentry.Rttmonechoadmincodecpayload}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminCodecNumPackets"] = types.YLeaf{"Rttmonechoadmincodecnumpackets", rttmonechoadminentry.Rttmonechoadmincodecnumpackets}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminICPIFAdvFactor"] = types.YLeaf{"Rttmonechoadminicpifadvfactor", rttmonechoadminentry.Rttmonechoadminicpifadvfactor}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPFECType"] = types.YLeaf{"Rttmonechoadminlspfectype", rttmonechoadminentry.Rttmonechoadminlspfectype}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPSelector"] = types.YLeaf{"Rttmonechoadminlspselector", rttmonechoadminentry.Rttmonechoadminlspselector}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPReplyMode"] = types.YLeaf{"Rttmonechoadminlspreplymode", rttmonechoadminentry.Rttmonechoadminlspreplymode}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPTTL"] = types.YLeaf{"Rttmonechoadminlspttl", rttmonechoadminentry.Rttmonechoadminlspttl}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPExp"] = types.YLeaf{"Rttmonechoadminlspexp", rttmonechoadminentry.Rttmonechoadminlspexp}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminPrecision"] = types.YLeaf{"Rttmonechoadminprecision", rttmonechoadminentry.Rttmonechoadminprecision}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminProbePakPriority"] = types.YLeaf{"Rttmonechoadminprobepakpriority", rttmonechoadminentry.Rttmonechoadminprobepakpriority}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminOWNTPSyncTolAbs"] = types.YLeaf{"Rttmonechoadminowntpsynctolabs", rttmonechoadminentry.Rttmonechoadminowntpsynctolabs}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminOWNTPSyncTolPct"] = types.YLeaf{"Rttmonechoadminowntpsynctolpct", rttmonechoadminentry.Rttmonechoadminowntpsynctolpct}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminOWNTPSyncTolType"] = types.YLeaf{"Rttmonechoadminowntpsynctoltype", rttmonechoadminentry.Rttmonechoadminowntpsynctoltype}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminCalledNumber"] = types.YLeaf{"Rttmonechoadmincallednumber", rttmonechoadminentry.Rttmonechoadmincallednumber}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminDetectPoint"] = types.YLeaf{"Rttmonechoadmindetectpoint", rttmonechoadminentry.Rttmonechoadmindetectpoint}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminGKRegistration"] = types.YLeaf{"Rttmonechoadmingkregistration", rttmonechoadminentry.Rttmonechoadmingkregistration}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminSourceVoicePort"] = types.YLeaf{"Rttmonechoadminsourcevoiceport", rttmonechoadminentry.Rttmonechoadminsourcevoiceport}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminCallDuration"] = types.YLeaf{"Rttmonechoadmincallduration", rttmonechoadminentry.Rttmonechoadmincallduration}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPReplyDscp"] = types.YLeaf{"Rttmonechoadminlspreplydscp", rttmonechoadminentry.Rttmonechoadminlspreplydscp}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPNullShim"] = types.YLeaf{"Rttmonechoadminlspnullshim", rttmonechoadminentry.Rttmonechoadminlspnullshim}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetMPID"] = types.YLeaf{"Rttmonechoadmintargetmpid", rttmonechoadminentry.Rttmonechoadmintargetmpid}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetDomainName"] = types.YLeaf{"Rttmonechoadmintargetdomainname", rttmonechoadminentry.Rttmonechoadmintargetdomainname}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetVLAN"] = types.YLeaf{"Rttmonechoadmintargetvlan", rttmonechoadminentry.Rttmonechoadmintargetvlan}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminEthernetCOS"] = types.YLeaf{"Rttmonechoadminethernetcos", rttmonechoadminentry.Rttmonechoadminethernetcos}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLSPVccvID"] = types.YLeaf{"Rttmonechoadminlspvccvid", rttmonechoadminentry.Rttmonechoadminlspvccvid}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetEVC"] = types.YLeaf{"Rttmonechoadmintargetevc", rttmonechoadminentry.Rttmonechoadmintargetevc}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetMEPPort"] = types.YLeaf{"Rttmonechoadmintargetmepport", rttmonechoadminentry.Rttmonechoadmintargetmepport}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminVideoTrafficProfile"] = types.YLeaf{"Rttmonechoadminvideotrafficprofile", rttmonechoadminentry.Rttmonechoadminvideotrafficprofile}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminDscp"] = types.YLeaf{"Rttmonechoadmindscp", rttmonechoadminentry.Rttmonechoadmindscp}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminReserveDsp"] = types.YLeaf{"Rttmonechoadminreservedsp", rttmonechoadminentry.Rttmonechoadminreservedsp}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminInputInterface"] = types.YLeaf{"Rttmonechoadmininputinterface", rttmonechoadminentry.Rttmonechoadmininputinterface}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminEmulateSourceAddress"] = types.YLeaf{"Rttmonechoadminemulatesourceaddress", rttmonechoadminentry.Rttmonechoadminemulatesourceaddress}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminEmulateSourcePort"] = types.YLeaf{"Rttmonechoadminemulatesourceport", rttmonechoadminentry.Rttmonechoadminemulatesourceport}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminEmulateTargetAddress"] = types.YLeaf{"Rttmonechoadminemulatetargetaddress", rttmonechoadminentry.Rttmonechoadminemulatetargetaddress}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminEmulateTargetPort"] = types.YLeaf{"Rttmonechoadminemulatetargetport", rttmonechoadminentry.Rttmonechoadminemulatetargetport}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTargetMacAddress"] = types.YLeaf{"Rttmonechoadmintargetmacaddress", rttmonechoadminentry.Rttmonechoadmintargetmacaddress}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminSourceMacAddress"] = types.YLeaf{"Rttmonechoadminsourcemacaddress", rttmonechoadminentry.Rttmonechoadminsourcemacaddress}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminSourceMPID"] = types.YLeaf{"Rttmonechoadminsourcempid", rttmonechoadminentry.Rttmonechoadminsourcempid}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminEndPointListName"] = types.YLeaf{"Rttmonechoadminendpointlistname", rttmonechoadminentry.Rttmonechoadminendpointlistname}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminSSM"] = types.YLeaf{"Rttmonechoadminssm", rttmonechoadminentry.Rttmonechoadminssm}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminControlRetry"] = types.YLeaf{"Rttmonechoadmincontrolretry", rttmonechoadminentry.Rttmonechoadmincontrolretry}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminControlTimeout"] = types.YLeaf{"Rttmonechoadmincontroltimeout", rttmonechoadminentry.Rttmonechoadmincontroltimeout}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminIgmpTreeInit"] = types.YLeaf{"Rttmonechoadminigmptreeinit", rttmonechoadminentry.Rttmonechoadminigmptreeinit}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminEnableBurst"] = types.YLeaf{"Rttmonechoadminenableburst", rttmonechoadminentry.Rttmonechoadminenableburst}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminAggBurstCycles"] = types.YLeaf{"Rttmonechoadminaggburstcycles", rttmonechoadminentry.Rttmonechoadminaggburstcycles}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminLossRatioNumFrames"] = types.YLeaf{"Rttmonechoadminlossrationumframes", rttmonechoadminentry.Rttmonechoadminlossrationumframes}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminAvailNumFrames"] = types.YLeaf{"Rttmonechoadminavailnumframes", rttmonechoadminentry.Rttmonechoadminavailnumframes}
    rttmonechoadminentry.EntityData.Leafs["rttMonEchoAdminTstampOptimization"] = types.YLeaf{"Rttmonechoadmintstampoptimization", rttmonechoadminentry.Rttmonechoadmintstampoptimization}
    return &(rttmonechoadminentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for 'fileIO'
    // RttMonRttType conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry.
    Rttmonfileioadminentry []CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry
}

func (rttmonfileioadmintable *CISCORTTMONMIB_Rttmonfileioadmintable) GetEntityData() *types.CommonEntityData {
    rttmonfileioadmintable.EntityData.YFilter = rttmonfileioadmintable.YFilter
    rttmonfileioadmintable.EntityData.YangName = "rttMonFileIOAdminTable"
    rttmonfileioadmintable.EntityData.BundleName = "cisco_ios_xe"
    rttmonfileioadmintable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonfileioadmintable.EntityData.SegmentPath = "rttMonFileIOAdminTable"
    rttmonfileioadmintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonfileioadmintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonfileioadmintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonfileioadmintable.EntityData.Children = make(map[string]types.YChild)
    rttmonfileioadmintable.EntityData.Children["rttMonFileIOAdminEntry"] = types.YChild{"Rttmonfileioadminentry", nil}
    for i := range rttmonfileioadmintable.Rttmonfileioadminentry {
        rttmonfileioadmintable.EntityData.Children[types.GetSegmentPath(&rttmonfileioadmintable.Rttmonfileioadminentry[i])] = types.YChild{"Rttmonfileioadminentry", &rttmonfileioadmintable.Rttmonfileioadminentry[i]}
    }
    rttmonfileioadmintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonfileioadmintable.EntityData)
}

// CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry
// A list of objects that define specific configuration for
// 'fileIO' RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry struct {
    EntityData types.CommonEntityData
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

func (rttmonfileioadminentry *CISCORTTMONMIB_Rttmonfileioadmintable_Rttmonfileioadminentry) GetEntityData() *types.CommonEntityData {
    rttmonfileioadminentry.EntityData.YFilter = rttmonfileioadminentry.YFilter
    rttmonfileioadminentry.EntityData.YangName = "rttMonFileIOAdminEntry"
    rttmonfileioadminentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonfileioadminentry.EntityData.ParentYangName = "rttMonFileIOAdminTable"
    rttmonfileioadminentry.EntityData.SegmentPath = "rttMonFileIOAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonfileioadminentry.Rttmonctrladminindex) + "']"
    rttmonfileioadminentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonfileioadminentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonfileioadminentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonfileioadminentry.EntityData.Children = make(map[string]types.YChild)
    rttmonfileioadminentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonfileioadminentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonfileioadminentry.Rttmonctrladminindex}
    rttmonfileioadminentry.EntityData.Leafs["rttMonFileIOAdminFilePath"] = types.YLeaf{"Rttmonfileioadminfilepath", rttmonfileioadminentry.Rttmonfileioadminfilepath}
    rttmonfileioadminentry.EntityData.Leafs["rttMonFileIOAdminSize"] = types.YLeaf{"Rttmonfileioadminsize", rttmonfileioadminentry.Rttmonfileioadminsize}
    rttmonfileioadminentry.EntityData.Leafs["rttMonFileIOAdminAction"] = types.YLeaf{"Rttmonfileioadminaction", rttmonfileioadminentry.Rttmonfileioadminaction}
    return &(rttmonfileioadminentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for 'script'
    // RttMonRttType conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry.
    Rttmonscriptadminentry []CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry
}

func (rttmonscriptadmintable *CISCORTTMONMIB_Rttmonscriptadmintable) GetEntityData() *types.CommonEntityData {
    rttmonscriptadmintable.EntityData.YFilter = rttmonscriptadmintable.YFilter
    rttmonscriptadmintable.EntityData.YangName = "rttMonScriptAdminTable"
    rttmonscriptadmintable.EntityData.BundleName = "cisco_ios_xe"
    rttmonscriptadmintable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonscriptadmintable.EntityData.SegmentPath = "rttMonScriptAdminTable"
    rttmonscriptadmintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonscriptadmintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonscriptadmintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonscriptadmintable.EntityData.Children = make(map[string]types.YChild)
    rttmonscriptadmintable.EntityData.Children["rttMonScriptAdminEntry"] = types.YChild{"Rttmonscriptadminentry", nil}
    for i := range rttmonscriptadmintable.Rttmonscriptadminentry {
        rttmonscriptadmintable.EntityData.Children[types.GetSegmentPath(&rttmonscriptadmintable.Rttmonscriptadminentry[i])] = types.YChild{"Rttmonscriptadminentry", &rttmonscriptadmintable.Rttmonscriptadminentry[i]}
    }
    rttmonscriptadmintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonscriptadmintable.EntityData)
}

// CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry
// A list of objects that define specific configuration for
// 'script' RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry struct {
    EntityData types.CommonEntityData
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

func (rttmonscriptadminentry *CISCORTTMONMIB_Rttmonscriptadmintable_Rttmonscriptadminentry) GetEntityData() *types.CommonEntityData {
    rttmonscriptadminentry.EntityData.YFilter = rttmonscriptadminentry.YFilter
    rttmonscriptadminentry.EntityData.YangName = "rttMonScriptAdminEntry"
    rttmonscriptadminentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonscriptadminentry.EntityData.ParentYangName = "rttMonScriptAdminTable"
    rttmonscriptadminentry.EntityData.SegmentPath = "rttMonScriptAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonscriptadminentry.Rttmonctrladminindex) + "']"
    rttmonscriptadminentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonscriptadminentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonscriptadminentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonscriptadminentry.EntityData.Children = make(map[string]types.YChild)
    rttmonscriptadminentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonscriptadminentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonscriptadminentry.Rttmonctrladminindex}
    rttmonscriptadminentry.EntityData.Leafs["rttMonScriptAdminName"] = types.YLeaf{"Rttmonscriptadminname", rttmonscriptadminentry.Rttmonscriptadminname}
    rttmonscriptadminentry.EntityData.Leafs["rttMonScriptAdminCmdLineParams"] = types.YLeaf{"Rttmonscriptadmincmdlineparams", rttmonscriptadminentry.Rttmonscriptadmincmdlineparams}
    return &(rttmonscriptadminentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that will be triggered when a reaction condition is
    // violated. The type is slice of
    // CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry.
    Rttmonreacttriggeradminentry []CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry
}

func (rttmonreacttriggeradmintable *CISCORTTMONMIB_Rttmonreacttriggeradmintable) GetEntityData() *types.CommonEntityData {
    rttmonreacttriggeradmintable.EntityData.YFilter = rttmonreacttriggeradmintable.YFilter
    rttmonreacttriggeradmintable.EntityData.YangName = "rttMonReactTriggerAdminTable"
    rttmonreacttriggeradmintable.EntityData.BundleName = "cisco_ios_xe"
    rttmonreacttriggeradmintable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonreacttriggeradmintable.EntityData.SegmentPath = "rttMonReactTriggerAdminTable"
    rttmonreacttriggeradmintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonreacttriggeradmintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonreacttriggeradmintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonreacttriggeradmintable.EntityData.Children = make(map[string]types.YChild)
    rttmonreacttriggeradmintable.EntityData.Children["rttMonReactTriggerAdminEntry"] = types.YChild{"Rttmonreacttriggeradminentry", nil}
    for i := range rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry {
        rttmonreacttriggeradmintable.EntityData.Children[types.GetSegmentPath(&rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry[i])] = types.YChild{"Rttmonreacttriggeradminentry", &rttmonreacttriggeradmintable.Rttmonreacttriggeradminentry[i]}
    }
    rttmonreacttriggeradmintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonreacttriggeradmintable.EntityData)
}

// CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry
// A list of objects that will be triggered when
// a reaction condition is violated.
type CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry struct {
    EntityData types.CommonEntityData
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

func (rttmonreacttriggeradminentry *CISCORTTMONMIB_Rttmonreacttriggeradmintable_Rttmonreacttriggeradminentry) GetEntityData() *types.CommonEntityData {
    rttmonreacttriggeradminentry.EntityData.YFilter = rttmonreacttriggeradminentry.YFilter
    rttmonreacttriggeradminentry.EntityData.YangName = "rttMonReactTriggerAdminEntry"
    rttmonreacttriggeradminentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonreacttriggeradminentry.EntityData.ParentYangName = "rttMonReactTriggerAdminTable"
    rttmonreacttriggeradminentry.EntityData.SegmentPath = "rttMonReactTriggerAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonreacttriggeradminentry.Rttmonctrladminindex) + "']" + "[rttMonReactTriggerAdminRttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonreacttriggeradminentry.Rttmonreacttriggeradminrttmonctrladminindex) + "']"
    rttmonreacttriggeradminentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonreacttriggeradminentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonreacttriggeradminentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonreacttriggeradminentry.EntityData.Children = make(map[string]types.YChild)
    rttmonreacttriggeradminentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonreacttriggeradminentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonreacttriggeradminentry.Rttmonctrladminindex}
    rttmonreacttriggeradminentry.EntityData.Leafs["rttMonReactTriggerAdminRttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonreacttriggeradminrttmonctrladminindex", rttmonreacttriggeradminentry.Rttmonreacttriggeradminrttmonctrladminindex}
    rttmonreacttriggeradminentry.EntityData.Leafs["rttMonReactTriggerAdminStatus"] = types.YLeaf{"Rttmonreacttriggeradminstatus", rttmonreacttriggeradminentry.Rttmonreacttriggeradminstatus}
    rttmonreacttriggeradminentry.EntityData.Leafs["rttMonReactTriggerOperState"] = types.YLeaf{"Rttmonreacttriggeroperstate", rttmonreacttriggeradminentry.Rttmonreacttriggeroperstate}
    return &(rttmonreacttriggeradminentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define intermediate hop's IP Address.  This entry
    // can be added only if the rttMonCtrlAdminRttType is 'echo'. The entry gets
    // deleted when the corresponding RTR entry, which has an index of
    // rttMonCtrlAdminIndex, is deleted. The type is slice of
    // CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry.
    Rttmonechopathadminentry []CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry
}

func (rttmonechopathadmintable *CISCORTTMONMIB_Rttmonechopathadmintable) GetEntityData() *types.CommonEntityData {
    rttmonechopathadmintable.EntityData.YFilter = rttmonechopathadmintable.YFilter
    rttmonechopathadmintable.EntityData.YangName = "rttMonEchoPathAdminTable"
    rttmonechopathadmintable.EntityData.BundleName = "cisco_ios_xe"
    rttmonechopathadmintable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonechopathadmintable.EntityData.SegmentPath = "rttMonEchoPathAdminTable"
    rttmonechopathadmintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonechopathadmintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonechopathadmintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonechopathadmintable.EntityData.Children = make(map[string]types.YChild)
    rttmonechopathadmintable.EntityData.Children["rttMonEchoPathAdminEntry"] = types.YChild{"Rttmonechopathadminentry", nil}
    for i := range rttmonechopathadmintable.Rttmonechopathadminentry {
        rttmonechopathadmintable.EntityData.Children[types.GetSegmentPath(&rttmonechopathadmintable.Rttmonechopathadminentry[i])] = types.YChild{"Rttmonechopathadminentry", &rttmonechopathadmintable.Rttmonechopathadminentry[i]}
    }
    rttmonechopathadmintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonechopathadmintable.EntityData)
}

// CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry
// A list of objects that define intermediate hop's IP Address.
// 
// This entry can be added only if the rttMonCtrlAdminRttType is
// 'echo'. The entry gets deleted when the corresponding RTR entry,
// which has an index of rttMonCtrlAdminIndex, is deleted.
type CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry struct {
    EntityData types.CommonEntityData
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

func (rttmonechopathadminentry *CISCORTTMONMIB_Rttmonechopathadmintable_Rttmonechopathadminentry) GetEntityData() *types.CommonEntityData {
    rttmonechopathadminentry.EntityData.YFilter = rttmonechopathadminentry.YFilter
    rttmonechopathadminentry.EntityData.YangName = "rttMonEchoPathAdminEntry"
    rttmonechopathadminentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonechopathadminentry.EntityData.ParentYangName = "rttMonEchoPathAdminTable"
    rttmonechopathadminentry.EntityData.SegmentPath = "rttMonEchoPathAdminEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonechopathadminentry.Rttmonctrladminindex) + "']" + "[rttMonEchoPathAdminHopIndex='" + fmt.Sprintf("%v", rttmonechopathadminentry.Rttmonechopathadminhopindex) + "']"
    rttmonechopathadminentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonechopathadminentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonechopathadminentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonechopathadminentry.EntityData.Children = make(map[string]types.YChild)
    rttmonechopathadminentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonechopathadminentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonechopathadminentry.Rttmonctrladminindex}
    rttmonechopathadminentry.EntityData.Leafs["rttMonEchoPathAdminHopIndex"] = types.YLeaf{"Rttmonechopathadminhopindex", rttmonechopathadminentry.Rttmonechopathadminhopindex}
    rttmonechopathadminentry.EntityData.Leafs["rttMonEchoPathAdminHopAddress"] = types.YLeaf{"Rttmonechopathadminhopaddress", rttmonechopathadminentry.Rttmonechopathadminhopaddress}
    return &(rttmonechopathadminentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define a conceptual group scheduling control row.
    // The type is slice of
    // CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry.
    Rttmongrpscheduleadminentry []CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry
}

func (rttmongrpscheduleadmintable *CISCORTTMONMIB_Rttmongrpscheduleadmintable) GetEntityData() *types.CommonEntityData {
    rttmongrpscheduleadmintable.EntityData.YFilter = rttmongrpscheduleadmintable.YFilter
    rttmongrpscheduleadmintable.EntityData.YangName = "rttMonGrpScheduleAdminTable"
    rttmongrpscheduleadmintable.EntityData.BundleName = "cisco_ios_xe"
    rttmongrpscheduleadmintable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmongrpscheduleadmintable.EntityData.SegmentPath = "rttMonGrpScheduleAdminTable"
    rttmongrpscheduleadmintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmongrpscheduleadmintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmongrpscheduleadmintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmongrpscheduleadmintable.EntityData.Children = make(map[string]types.YChild)
    rttmongrpscheduleadmintable.EntityData.Children["rttMonGrpScheduleAdminEntry"] = types.YChild{"Rttmongrpscheduleadminentry", nil}
    for i := range rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry {
        rttmongrpscheduleadmintable.EntityData.Children[types.GetSegmentPath(&rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry[i])] = types.YChild{"Rttmongrpscheduleadminentry", &rttmongrpscheduleadmintable.Rttmongrpscheduleadminentry[i]}
    }
    rttmongrpscheduleadmintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmongrpscheduleadmintable.EntityData)
}

// CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry
// A list of objects that define a conceptual group scheduling
// control row.
type CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry struct {
    EntityData types.CommonEntityData
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

func (rttmongrpscheduleadminentry *CISCORTTMONMIB_Rttmongrpscheduleadmintable_Rttmongrpscheduleadminentry) GetEntityData() *types.CommonEntityData {
    rttmongrpscheduleadminentry.EntityData.YFilter = rttmongrpscheduleadminentry.YFilter
    rttmongrpscheduleadminentry.EntityData.YangName = "rttMonGrpScheduleAdminEntry"
    rttmongrpscheduleadminentry.EntityData.BundleName = "cisco_ios_xe"
    rttmongrpscheduleadminentry.EntityData.ParentYangName = "rttMonGrpScheduleAdminTable"
    rttmongrpscheduleadminentry.EntityData.SegmentPath = "rttMonGrpScheduleAdminEntry" + "[rttMonGrpScheduleAdminIndex='" + fmt.Sprintf("%v", rttmongrpscheduleadminentry.Rttmongrpscheduleadminindex) + "']"
    rttmongrpscheduleadminentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmongrpscheduleadminentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmongrpscheduleadminentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmongrpscheduleadminentry.EntityData.Children = make(map[string]types.YChild)
    rttmongrpscheduleadminentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminIndex"] = types.YLeaf{"Rttmongrpscheduleadminindex", rttmongrpscheduleadminentry.Rttmongrpscheduleadminindex}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminProbes"] = types.YLeaf{"Rttmongrpscheduleadminprobes", rttmongrpscheduleadminentry.Rttmongrpscheduleadminprobes}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminPeriod"] = types.YLeaf{"Rttmongrpscheduleadminperiod", rttmongrpscheduleadminentry.Rttmongrpscheduleadminperiod}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminFrequency"] = types.YLeaf{"Rttmongrpscheduleadminfrequency", rttmongrpscheduleadminentry.Rttmongrpscheduleadminfrequency}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminLife"] = types.YLeaf{"Rttmongrpscheduleadminlife", rttmongrpscheduleadminentry.Rttmongrpscheduleadminlife}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminAgeout"] = types.YLeaf{"Rttmongrpscheduleadminageout", rttmongrpscheduleadminentry.Rttmongrpscheduleadminageout}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminStatus"] = types.YLeaf{"Rttmongrpscheduleadminstatus", rttmongrpscheduleadminentry.Rttmongrpscheduleadminstatus}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminFreqMax"] = types.YLeaf{"Rttmongrpscheduleadminfreqmax", rttmongrpscheduleadminentry.Rttmongrpscheduleadminfreqmax}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminFreqMin"] = types.YLeaf{"Rttmongrpscheduleadminfreqmin", rttmongrpscheduleadminentry.Rttmongrpscheduleadminfreqmin}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminStartTime"] = types.YLeaf{"Rttmongrpscheduleadminstarttime", rttmongrpscheduleadminentry.Rttmongrpscheduleadminstarttime}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminAdd"] = types.YLeaf{"Rttmongrpscheduleadminadd", rttmongrpscheduleadminentry.Rttmongrpscheduleadminadd}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminDelete"] = types.YLeaf{"Rttmongrpscheduleadmindelete", rttmongrpscheduleadminentry.Rttmongrpscheduleadmindelete}
    rttmongrpscheduleadminentry.EntityData.Leafs["rttMonGrpScheduleAdminReset"] = types.YLeaf{"Rttmongrpscheduleadminreset", rttmongrpscheduleadminentry.Rttmongrpscheduleadminreset}
    return &(rttmongrpscheduleadminentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual Auto SAA L3 MPLS VPN
    // control row. The type is slice of
    // CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry.
    Rttmplsvpnmonctrlentry []CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry
}

func (rttmplsvpnmonctrltable *CISCORTTMONMIB_Rttmplsvpnmonctrltable) GetEntityData() *types.CommonEntityData {
    rttmplsvpnmonctrltable.EntityData.YFilter = rttmplsvpnmonctrltable.YFilter
    rttmplsvpnmonctrltable.EntityData.YangName = "rttMplsVpnMonCtrlTable"
    rttmplsvpnmonctrltable.EntityData.BundleName = "cisco_ios_xe"
    rttmplsvpnmonctrltable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmplsvpnmonctrltable.EntityData.SegmentPath = "rttMplsVpnMonCtrlTable"
    rttmplsvpnmonctrltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmplsvpnmonctrltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmplsvpnmonctrltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmplsvpnmonctrltable.EntityData.Children = make(map[string]types.YChild)
    rttmplsvpnmonctrltable.EntityData.Children["rttMplsVpnMonCtrlEntry"] = types.YChild{"Rttmplsvpnmonctrlentry", nil}
    for i := range rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry {
        rttmplsvpnmonctrltable.EntityData.Children[types.GetSegmentPath(&rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry[i])] = types.YChild{"Rttmplsvpnmonctrlentry", &rttmplsvpnmonctrltable.Rttmplsvpnmonctrlentry[i]}
    }
    rttmplsvpnmonctrltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmplsvpnmonctrltable.EntityData)
}

// CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry
// A base list of objects that define a conceptual Auto SAA L3
// MPLS VPN control row.
type CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry struct {
    EntityData types.CommonEntityData
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

func (rttmplsvpnmonctrlentry *CISCORTTMONMIB_Rttmplsvpnmonctrltable_Rttmplsvpnmonctrlentry) GetEntityData() *types.CommonEntityData {
    rttmplsvpnmonctrlentry.EntityData.YFilter = rttmplsvpnmonctrlentry.YFilter
    rttmplsvpnmonctrlentry.EntityData.YangName = "rttMplsVpnMonCtrlEntry"
    rttmplsvpnmonctrlentry.EntityData.BundleName = "cisco_ios_xe"
    rttmplsvpnmonctrlentry.EntityData.ParentYangName = "rttMplsVpnMonCtrlTable"
    rttmplsvpnmonctrlentry.EntityData.SegmentPath = "rttMplsVpnMonCtrlEntry" + "[rttMplsVpnMonCtrlIndex='" + fmt.Sprintf("%v", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlindex) + "']"
    rttmplsvpnmonctrlentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmplsvpnmonctrlentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmplsvpnmonctrlentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmplsvpnmonctrlentry.EntityData.Children = make(map[string]types.YChild)
    rttmplsvpnmonctrlentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlIndex"] = types.YLeaf{"Rttmplsvpnmonctrlindex", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlindex}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlRttType"] = types.YLeaf{"Rttmplsvpnmonctrlrtttype", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlrtttype}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlVrfName"] = types.YLeaf{"Rttmplsvpnmonctrlvrfname", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlvrfname}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlTag"] = types.YLeaf{"Rttmplsvpnmonctrltag", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrltag}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlThreshold"] = types.YLeaf{"Rttmplsvpnmonctrlthreshold", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlthreshold}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlTimeout"] = types.YLeaf{"Rttmplsvpnmonctrltimeout", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrltimeout}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlScanInterval"] = types.YLeaf{"Rttmplsvpnmonctrlscaninterval", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlscaninterval}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlDelScanFactor"] = types.YLeaf{"Rttmplsvpnmonctrldelscanfactor", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrldelscanfactor}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlEXP"] = types.YLeaf{"Rttmplsvpnmonctrlexp", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlexp}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlRequestSize"] = types.YLeaf{"Rttmplsvpnmonctrlrequestsize", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlrequestsize}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlVerifyData"] = types.YLeaf{"Rttmplsvpnmonctrlverifydata", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlverifydata}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlStorageType"] = types.YLeaf{"Rttmplsvpnmonctrlstoragetype", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlstoragetype}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlProbeList"] = types.YLeaf{"Rttmplsvpnmonctrlprobelist", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlprobelist}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlStatus"] = types.YLeaf{"Rttmplsvpnmonctrlstatus", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrlstatus}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlLpd"] = types.YLeaf{"Rttmplsvpnmonctrllpd", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrllpd}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlLpdGrpList"] = types.YLeaf{"Rttmplsvpnmonctrllpdgrplist", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrllpdgrplist}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonCtrlLpdCompTime"] = types.YLeaf{"Rttmplsvpnmonctrllpdcomptime", rttmplsvpnmonctrlentry.Rttmplsvpnmonctrllpdcomptime}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeInterval"] = types.YLeaf{"Rttmplsvpnmontypeinterval", rttmplsvpnmonctrlentry.Rttmplsvpnmontypeinterval}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeNumPackets"] = types.YLeaf{"Rttmplsvpnmontypenumpackets", rttmplsvpnmonctrlentry.Rttmplsvpnmontypenumpackets}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeDestPort"] = types.YLeaf{"Rttmplsvpnmontypedestport", rttmplsvpnmonctrlentry.Rttmplsvpnmontypedestport}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeSecFreqType"] = types.YLeaf{"Rttmplsvpnmontypesecfreqtype", rttmplsvpnmonctrlentry.Rttmplsvpnmontypesecfreqtype}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeSecFreqValue"] = types.YLeaf{"Rttmplsvpnmontypesecfreqvalue", rttmplsvpnmonctrlentry.Rttmplsvpnmontypesecfreqvalue}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLspSelector"] = types.YLeaf{"Rttmplsvpnmontypelspselector", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspselector}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLSPReplyMode"] = types.YLeaf{"Rttmplsvpnmontypelspreplymode", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspreplymode}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLSPTTL"] = types.YLeaf{"Rttmplsvpnmontypelspttl", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspttl}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLSPReplyDscp"] = types.YLeaf{"Rttmplsvpnmontypelspreplydscp", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelspreplydscp}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLpdMaxSessions"] = types.YLeaf{"Rttmplsvpnmontypelpdmaxsessions", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdmaxsessions}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLpdSessTimeout"] = types.YLeaf{"Rttmplsvpnmontypelpdsesstimeout", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdsesstimeout}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLpdEchoTimeout"] = types.YLeaf{"Rttmplsvpnmontypelpdechotimeout", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdechotimeout}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLpdEchoInterval"] = types.YLeaf{"Rttmplsvpnmontypelpdechointerval", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdechointerval}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLpdEchoNullShim"] = types.YLeaf{"Rttmplsvpnmontypelpdechonullshim", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdechonullshim}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLpdScanPeriod"] = types.YLeaf{"Rttmplsvpnmontypelpdscanperiod", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdscanperiod}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonTypeLpdStatHours"] = types.YLeaf{"Rttmplsvpnmontypelpdstathours", rttmplsvpnmonctrlentry.Rttmplsvpnmontypelpdstathours}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonScheduleRttStartTime"] = types.YLeaf{"Rttmplsvpnmonschedulerttstarttime", rttmplsvpnmonctrlentry.Rttmplsvpnmonschedulerttstarttime}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonSchedulePeriod"] = types.YLeaf{"Rttmplsvpnmonscheduleperiod", rttmplsvpnmonctrlentry.Rttmplsvpnmonscheduleperiod}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonScheduleFrequency"] = types.YLeaf{"Rttmplsvpnmonschedulefrequency", rttmplsvpnmonctrlentry.Rttmplsvpnmonschedulefrequency}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonReactConnectionEnable"] = types.YLeaf{"Rttmplsvpnmonreactconnectionenable", rttmplsvpnmonctrlentry.Rttmplsvpnmonreactconnectionenable}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonReactTimeoutEnable"] = types.YLeaf{"Rttmplsvpnmonreacttimeoutenable", rttmplsvpnmonctrlentry.Rttmplsvpnmonreacttimeoutenable}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonReactThresholdType"] = types.YLeaf{"Rttmplsvpnmonreactthresholdtype", rttmplsvpnmonctrlentry.Rttmplsvpnmonreactthresholdtype}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonReactThresholdCount"] = types.YLeaf{"Rttmplsvpnmonreactthresholdcount", rttmplsvpnmonctrlentry.Rttmplsvpnmonreactthresholdcount}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonReactActionType"] = types.YLeaf{"Rttmplsvpnmonreactactiontype", rttmplsvpnmonctrlentry.Rttmplsvpnmonreactactiontype}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonReactLpdNotifyType"] = types.YLeaf{"Rttmplsvpnmonreactlpdnotifytype", rttmplsvpnmonctrlentry.Rttmplsvpnmonreactlpdnotifytype}
    rttmplsvpnmonctrlentry.EntityData.Leafs["rttMplsVpnMonReactLpdRetryCount"] = types.YLeaf{"Rttmplsvpnmonreactlpdretrycount", rttmplsvpnmonctrlentry.Rttmplsvpnmonreactlpdretrycount}
    return &(rttmplsvpnmonctrlentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual reaction configuration
    // control row. The type is slice of
    // CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry.
    Rttmonreactentry []CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry
}

func (rttmonreacttable *CISCORTTMONMIB_Rttmonreacttable) GetEntityData() *types.CommonEntityData {
    rttmonreacttable.EntityData.YFilter = rttmonreacttable.YFilter
    rttmonreacttable.EntityData.YangName = "rttMonReactTable"
    rttmonreacttable.EntityData.BundleName = "cisco_ios_xe"
    rttmonreacttable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonreacttable.EntityData.SegmentPath = "rttMonReactTable"
    rttmonreacttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonreacttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonreacttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonreacttable.EntityData.Children = make(map[string]types.YChild)
    rttmonreacttable.EntityData.Children["rttMonReactEntry"] = types.YChild{"Rttmonreactentry", nil}
    for i := range rttmonreacttable.Rttmonreactentry {
        rttmonreacttable.EntityData.Children[types.GetSegmentPath(&rttmonreacttable.Rttmonreactentry[i])] = types.YChild{"Rttmonreactentry", &rttmonreacttable.Rttmonreactentry[i]}
    }
    rttmonreacttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonreacttable.EntityData)
}

// CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry
// A base list of objects that define a conceptual reaction
// configuration control row.
type CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry struct {
    EntityData types.CommonEntityData
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

func (rttmonreactentry *CISCORTTMONMIB_Rttmonreacttable_Rttmonreactentry) GetEntityData() *types.CommonEntityData {
    rttmonreactentry.EntityData.YFilter = rttmonreactentry.YFilter
    rttmonreactentry.EntityData.YangName = "rttMonReactEntry"
    rttmonreactentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonreactentry.EntityData.ParentYangName = "rttMonReactTable"
    rttmonreactentry.EntityData.SegmentPath = "rttMonReactEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonreactentry.Rttmonctrladminindex) + "']" + "[rttMonReactConfigIndex='" + fmt.Sprintf("%v", rttmonreactentry.Rttmonreactconfigindex) + "']"
    rttmonreactentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonreactentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonreactentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonreactentry.EntityData.Children = make(map[string]types.YChild)
    rttmonreactentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonreactentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonreactentry.Rttmonctrladminindex}
    rttmonreactentry.EntityData.Leafs["rttMonReactConfigIndex"] = types.YLeaf{"Rttmonreactconfigindex", rttmonreactentry.Rttmonreactconfigindex}
    rttmonreactentry.EntityData.Leafs["rttMonReactVar"] = types.YLeaf{"Rttmonreactvar", rttmonreactentry.Rttmonreactvar}
    rttmonreactentry.EntityData.Leafs["rttMonReactThresholdType"] = types.YLeaf{"Rttmonreactthresholdtype", rttmonreactentry.Rttmonreactthresholdtype}
    rttmonreactentry.EntityData.Leafs["rttMonReactActionType"] = types.YLeaf{"Rttmonreactactiontype", rttmonreactentry.Rttmonreactactiontype}
    rttmonreactentry.EntityData.Leafs["rttMonReactThresholdRising"] = types.YLeaf{"Rttmonreactthresholdrising", rttmonreactentry.Rttmonreactthresholdrising}
    rttmonreactentry.EntityData.Leafs["rttMonReactThresholdFalling"] = types.YLeaf{"Rttmonreactthresholdfalling", rttmonreactentry.Rttmonreactthresholdfalling}
    rttmonreactentry.EntityData.Leafs["rttMonReactThresholdCountX"] = types.YLeaf{"Rttmonreactthresholdcountx", rttmonreactentry.Rttmonreactthresholdcountx}
    rttmonreactentry.EntityData.Leafs["rttMonReactThresholdCountY"] = types.YLeaf{"Rttmonreactthresholdcounty", rttmonreactentry.Rttmonreactthresholdcounty}
    rttmonreactentry.EntityData.Leafs["rttMonReactValue"] = types.YLeaf{"Rttmonreactvalue", rttmonreactentry.Rttmonreactvalue}
    rttmonreactentry.EntityData.Leafs["rttMonReactOccurred"] = types.YLeaf{"Rttmonreactoccurred", rttmonreactentry.Rttmonreactoccurred}
    rttmonreactentry.EntityData.Leafs["rttMonReactStatus"] = types.YLeaf{"Rttmonreactstatus", rttmonreactentry.Rttmonreactstatus}
    return &(rttmonreactentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Generated Oper table corresponding to a child or generated
    // operation as part of a parent IP SLA operation. The type is slice of
    // CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry.
    Rttmongeneratedoperentry []CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry
}

func (rttmongeneratedopertable *CISCORTTMONMIB_Rttmongeneratedopertable) GetEntityData() *types.CommonEntityData {
    rttmongeneratedopertable.EntityData.YFilter = rttmongeneratedopertable.YFilter
    rttmongeneratedopertable.EntityData.YangName = "rttMonGeneratedOperTable"
    rttmongeneratedopertable.EntityData.BundleName = "cisco_ios_xe"
    rttmongeneratedopertable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmongeneratedopertable.EntityData.SegmentPath = "rttMonGeneratedOperTable"
    rttmongeneratedopertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmongeneratedopertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmongeneratedopertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmongeneratedopertable.EntityData.Children = make(map[string]types.YChild)
    rttmongeneratedopertable.EntityData.Children["rttMonGeneratedOperEntry"] = types.YChild{"Rttmongeneratedoperentry", nil}
    for i := range rttmongeneratedopertable.Rttmongeneratedoperentry {
        rttmongeneratedopertable.EntityData.Children[types.GetSegmentPath(&rttmongeneratedopertable.Rttmongeneratedoperentry[i])] = types.YChild{"Rttmongeneratedoperentry", &rttmongeneratedopertable.Rttmongeneratedoperentry[i]}
    }
    rttmongeneratedopertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmongeneratedopertable.EntityData)
}

// CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry
// An entry in the Generated Oper table corresponding to
// a child or generated operation as part of a parent
// IP SLA operation.
type CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry struct {
    EntityData types.CommonEntityData
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

func (rttmongeneratedoperentry *CISCORTTMONMIB_Rttmongeneratedopertable_Rttmongeneratedoperentry) GetEntityData() *types.CommonEntityData {
    rttmongeneratedoperentry.EntityData.YFilter = rttmongeneratedoperentry.YFilter
    rttmongeneratedoperentry.EntityData.YangName = "rttMonGeneratedOperEntry"
    rttmongeneratedoperentry.EntityData.BundleName = "cisco_ios_xe"
    rttmongeneratedoperentry.EntityData.ParentYangName = "rttMonGeneratedOperTable"
    rttmongeneratedoperentry.EntityData.SegmentPath = "rttMonGeneratedOperEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmongeneratedoperentry.Rttmonctrladminindex) + "']" + "[rttMonGeneratedOperRespIpAddrType='" + fmt.Sprintf("%v", rttmongeneratedoperentry.Rttmongeneratedoperrespipaddrtype) + "']" + "[rttMonGeneratedOperRespIpAddr='" + fmt.Sprintf("%v", rttmongeneratedoperentry.Rttmongeneratedoperrespipaddr) + "']"
    rttmongeneratedoperentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmongeneratedoperentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmongeneratedoperentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmongeneratedoperentry.EntityData.Children = make(map[string]types.YChild)
    rttmongeneratedoperentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmongeneratedoperentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmongeneratedoperentry.Rttmonctrladminindex}
    rttmongeneratedoperentry.EntityData.Leafs["rttMonGeneratedOperRespIpAddrType"] = types.YLeaf{"Rttmongeneratedoperrespipaddrtype", rttmongeneratedoperentry.Rttmongeneratedoperrespipaddrtype}
    rttmongeneratedoperentry.EntityData.Leafs["rttMonGeneratedOperRespIpAddr"] = types.YLeaf{"Rttmongeneratedoperrespipaddr", rttmongeneratedoperentry.Rttmongeneratedoperrespipaddr}
    rttmongeneratedoperentry.EntityData.Leafs["rttMonGeneratedOperCtrlAdminIndex"] = types.YLeaf{"Rttmongeneratedoperctrladminindex", rttmongeneratedoperentry.Rttmongeneratedoperctrladminindex}
    return &(rttmongeneratedoperentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonstatscapturetable *CISCORTTMONMIB_Rttmonstatscapturetable) GetEntityData() *types.CommonEntityData {
    rttmonstatscapturetable.EntityData.YFilter = rttmonstatscapturetable.YFilter
    rttmonstatscapturetable.EntityData.YangName = "rttMonStatsCaptureTable"
    rttmonstatscapturetable.EntityData.BundleName = "cisco_ios_xe"
    rttmonstatscapturetable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonstatscapturetable.EntityData.SegmentPath = "rttMonStatsCaptureTable"
    rttmonstatscapturetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonstatscapturetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonstatscapturetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonstatscapturetable.EntityData.Children = make(map[string]types.YChild)
    rttmonstatscapturetable.EntityData.Children["rttMonStatsCaptureEntry"] = types.YChild{"Rttmonstatscaptureentry", nil}
    for i := range rttmonstatscapturetable.Rttmonstatscaptureentry {
        rttmonstatscapturetable.EntityData.Children[types.GetSegmentPath(&rttmonstatscapturetable.Rttmonstatscaptureentry[i])] = types.YChild{"Rttmonstatscaptureentry", &rttmonstatscapturetable.Rttmonstatscaptureentry[i]}
    }
    rttmonstatscapturetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonstatscapturetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonstatscaptureentry *CISCORTTMONMIB_Rttmonstatscapturetable_Rttmonstatscaptureentry) GetEntityData() *types.CommonEntityData {
    rttmonstatscaptureentry.EntityData.YFilter = rttmonstatscaptureentry.YFilter
    rttmonstatscaptureentry.EntityData.YangName = "rttMonStatsCaptureEntry"
    rttmonstatscaptureentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonstatscaptureentry.EntityData.ParentYangName = "rttMonStatsCaptureTable"
    rttmonstatscaptureentry.EntityData.SegmentPath = "rttMonStatsCaptureEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonctrladminindex) + "']" + "[rttMonStatsCaptureStartTimeIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturestarttimeindex) + "']" + "[rttMonStatsCapturePathIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturepathindex) + "']" + "[rttMonStatsCaptureHopIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturehopindex) + "']" + "[rttMonStatsCaptureDistIndex='" + fmt.Sprintf("%v", rttmonstatscaptureentry.Rttmonstatscapturedistindex) + "']"
    rttmonstatscaptureentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonstatscaptureentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonstatscaptureentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonstatscaptureentry.EntityData.Children = make(map[string]types.YChild)
    rttmonstatscaptureentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonstatscaptureentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonstatscaptureentry.Rttmonctrladminindex}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureStartTimeIndex"] = types.YLeaf{"Rttmonstatscapturestarttimeindex", rttmonstatscaptureentry.Rttmonstatscapturestarttimeindex}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCapturePathIndex"] = types.YLeaf{"Rttmonstatscapturepathindex", rttmonstatscaptureentry.Rttmonstatscapturepathindex}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureHopIndex"] = types.YLeaf{"Rttmonstatscapturehopindex", rttmonstatscaptureentry.Rttmonstatscapturehopindex}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureDistIndex"] = types.YLeaf{"Rttmonstatscapturedistindex", rttmonstatscaptureentry.Rttmonstatscapturedistindex}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureCompletions"] = types.YLeaf{"Rttmonstatscapturecompletions", rttmonstatscaptureentry.Rttmonstatscapturecompletions}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureOverThresholds"] = types.YLeaf{"Rttmonstatscaptureoverthresholds", rttmonstatscaptureentry.Rttmonstatscaptureoverthresholds}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureSumCompletionTime"] = types.YLeaf{"Rttmonstatscapturesumcompletiontime", rttmonstatscaptureentry.Rttmonstatscapturesumcompletiontime}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureSumCompletionTime2Low"] = types.YLeaf{"Rttmonstatscapturesumcompletiontime2Low", rttmonstatscaptureentry.Rttmonstatscapturesumcompletiontime2Low}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureSumCompletionTime2High"] = types.YLeaf{"Rttmonstatscapturesumcompletiontime2High", rttmonstatscaptureentry.Rttmonstatscapturesumcompletiontime2High}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureCompletionTimeMax"] = types.YLeaf{"Rttmonstatscapturecompletiontimemax", rttmonstatscaptureentry.Rttmonstatscapturecompletiontimemax}
    rttmonstatscaptureentry.EntityData.Leafs["rttMonStatsCaptureCompletionTimeMin"] = types.YLeaf{"Rttmonstatscapturecompletiontimemin", rttmonstatscaptureentry.Rttmonstatscapturecompletiontimemin}
    return &(rttmonstatscaptureentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry has the exact same
    // behavior as the  rttMonStatsCaptureEntry, except it does not keep
    // statistical distribution information.  For a complete entry description see
    // the rttMonStatsCaptureEntry object. The type is slice of
    // CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry.
    Rttmonstatscollectentry []CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry
}

func (rttmonstatscollecttable *CISCORTTMONMIB_Rttmonstatscollecttable) GetEntityData() *types.CommonEntityData {
    rttmonstatscollecttable.EntityData.YFilter = rttmonstatscollecttable.YFilter
    rttmonstatscollecttable.EntityData.YangName = "rttMonStatsCollectTable"
    rttmonstatscollecttable.EntityData.BundleName = "cisco_ios_xe"
    rttmonstatscollecttable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonstatscollecttable.EntityData.SegmentPath = "rttMonStatsCollectTable"
    rttmonstatscollecttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonstatscollecttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonstatscollecttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonstatscollecttable.EntityData.Children = make(map[string]types.YChild)
    rttmonstatscollecttable.EntityData.Children["rttMonStatsCollectEntry"] = types.YChild{"Rttmonstatscollectentry", nil}
    for i := range rttmonstatscollecttable.Rttmonstatscollectentry {
        rttmonstatscollecttable.EntityData.Children[types.GetSegmentPath(&rttmonstatscollecttable.Rttmonstatscollectentry[i])] = types.YChild{"Rttmonstatscollectentry", &rttmonstatscollecttable.Rttmonstatscollectentry[i]}
    }
    rttmonstatscollecttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonstatscollecttable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonstatscollectentry *CISCORTTMONMIB_Rttmonstatscollecttable_Rttmonstatscollectentry) GetEntityData() *types.CommonEntityData {
    rttmonstatscollectentry.EntityData.YFilter = rttmonstatscollectentry.YFilter
    rttmonstatscollectentry.EntityData.YangName = "rttMonStatsCollectEntry"
    rttmonstatscollectentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonstatscollectentry.EntityData.ParentYangName = "rttMonStatsCollectTable"
    rttmonstatscollectentry.EntityData.SegmentPath = "rttMonStatsCollectEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonctrladminindex) + "']" + "[rttMonStatsCaptureStartTimeIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonstatscapturestarttimeindex) + "']" + "[rttMonStatsCapturePathIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonstatscapturepathindex) + "']" + "[rttMonStatsCaptureHopIndex='" + fmt.Sprintf("%v", rttmonstatscollectentry.Rttmonstatscapturehopindex) + "']"
    rttmonstatscollectentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonstatscollectentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonstatscollectentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonstatscollectentry.EntityData.Children = make(map[string]types.YChild)
    rttmonstatscollectentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonstatscollectentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonstatscollectentry.Rttmonctrladminindex}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCaptureStartTimeIndex"] = types.YLeaf{"Rttmonstatscapturestarttimeindex", rttmonstatscollectentry.Rttmonstatscapturestarttimeindex}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCapturePathIndex"] = types.YLeaf{"Rttmonstatscapturepathindex", rttmonstatscollectentry.Rttmonstatscapturepathindex}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCaptureHopIndex"] = types.YLeaf{"Rttmonstatscapturehopindex", rttmonstatscollectentry.Rttmonstatscapturehopindex}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectNumDisconnects"] = types.YLeaf{"Rttmonstatscollectnumdisconnects", rttmonstatscollectentry.Rttmonstatscollectnumdisconnects}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectTimeouts"] = types.YLeaf{"Rttmonstatscollecttimeouts", rttmonstatscollectentry.Rttmonstatscollecttimeouts}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectBusies"] = types.YLeaf{"Rttmonstatscollectbusies", rttmonstatscollectentry.Rttmonstatscollectbusies}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectNoConnections"] = types.YLeaf{"Rttmonstatscollectnoconnections", rttmonstatscollectentry.Rttmonstatscollectnoconnections}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectDrops"] = types.YLeaf{"Rttmonstatscollectdrops", rttmonstatscollectentry.Rttmonstatscollectdrops}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectSequenceErrors"] = types.YLeaf{"Rttmonstatscollectsequenceerrors", rttmonstatscollectentry.Rttmonstatscollectsequenceerrors}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectVerifyErrors"] = types.YLeaf{"Rttmonstatscollectverifyerrors", rttmonstatscollectentry.Rttmonstatscollectverifyerrors}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectAddress"] = types.YLeaf{"Rttmonstatscollectaddress", rttmonstatscollectentry.Rttmonstatscollectaddress}
    rttmonstatscollectentry.EntityData.Leafs["rttMonControlEnableErrors"] = types.YLeaf{"Rttmoncontrolenableerrors", rttmonstatscollectentry.Rttmoncontrolenableerrors}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsRetrieveErrors"] = types.YLeaf{"Rttmonstatsretrieveerrors", rttmonstatscollectentry.Rttmonstatsretrieveerrors}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectCtrlEnErrors"] = types.YLeaf{"Rttmonstatscollectctrlenerrors", rttmonstatscollectentry.Rttmonstatscollectctrlenerrors}
    rttmonstatscollectentry.EntityData.Leafs["rttMonStatsCollectRetrieveErrors"] = types.YLeaf{"Rttmonstatscollectretrieveerrors", rttmonstatscollectentry.Rttmonstatscollectretrieveerrors}
    return &(rttmonstatscollectentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry has the exact same
    // behavior as the  rttMonStatsCaptureEntry, except it only keeps 60 minute
    // group values.  For a complete entry description see the
    // rttMonStatsCaptureEntry object. The type is slice of
    // CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry.
    Rttmonstatstotalsentry []CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry
}

func (rttmonstatstotalstable *CISCORTTMONMIB_Rttmonstatstotalstable) GetEntityData() *types.CommonEntityData {
    rttmonstatstotalstable.EntityData.YFilter = rttmonstatstotalstable.YFilter
    rttmonstatstotalstable.EntityData.YangName = "rttMonStatsTotalsTable"
    rttmonstatstotalstable.EntityData.BundleName = "cisco_ios_xe"
    rttmonstatstotalstable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonstatstotalstable.EntityData.SegmentPath = "rttMonStatsTotalsTable"
    rttmonstatstotalstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonstatstotalstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonstatstotalstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonstatstotalstable.EntityData.Children = make(map[string]types.YChild)
    rttmonstatstotalstable.EntityData.Children["rttMonStatsTotalsEntry"] = types.YChild{"Rttmonstatstotalsentry", nil}
    for i := range rttmonstatstotalstable.Rttmonstatstotalsentry {
        rttmonstatstotalstable.EntityData.Children[types.GetSegmentPath(&rttmonstatstotalstable.Rttmonstatstotalsentry[i])] = types.YChild{"Rttmonstatstotalsentry", &rttmonstatstotalstable.Rttmonstatstotalsentry[i]}
    }
    rttmonstatstotalstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonstatstotalstable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonstatstotalsentry *CISCORTTMONMIB_Rttmonstatstotalstable_Rttmonstatstotalsentry) GetEntityData() *types.CommonEntityData {
    rttmonstatstotalsentry.EntityData.YFilter = rttmonstatstotalsentry.YFilter
    rttmonstatstotalsentry.EntityData.YangName = "rttMonStatsTotalsEntry"
    rttmonstatstotalsentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonstatstotalsentry.EntityData.ParentYangName = "rttMonStatsTotalsTable"
    rttmonstatstotalsentry.EntityData.SegmentPath = "rttMonStatsTotalsEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonstatstotalsentry.Rttmonctrladminindex) + "']" + "[rttMonStatsCaptureStartTimeIndex='" + fmt.Sprintf("%v", rttmonstatstotalsentry.Rttmonstatscapturestarttimeindex) + "']"
    rttmonstatstotalsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonstatstotalsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonstatstotalsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonstatstotalsentry.EntityData.Children = make(map[string]types.YChild)
    rttmonstatstotalsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonstatstotalsentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonstatstotalsentry.Rttmonctrladminindex}
    rttmonstatstotalsentry.EntityData.Leafs["rttMonStatsCaptureStartTimeIndex"] = types.YLeaf{"Rttmonstatscapturestarttimeindex", rttmonstatstotalsentry.Rttmonstatscapturestarttimeindex}
    rttmonstatstotalsentry.EntityData.Leafs["rttMonStatsTotalsElapsedTime"] = types.YLeaf{"Rttmonstatstotalselapsedtime", rttmonstatstotalsentry.Rttmonstatstotalselapsedtime}
    rttmonstatstotalsentry.EntityData.Leafs["rttMonStatsTotalsInitiations"] = types.YLeaf{"Rttmonstatstotalsinitiations", rttmonstatstotalsentry.Rttmonstatstotalsinitiations}
    return &(rttmonstatstotalsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry is created only if the
    // rttMonCtrlAdminRttType  is http. The operation of this table is same as
    // that of rttMonStatsCaptureTable. The type is slice of
    // CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry.
    Rttmonhttpstatsentry []CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry
}

func (rttmonhttpstatstable *CISCORTTMONMIB_Rttmonhttpstatstable) GetEntityData() *types.CommonEntityData {
    rttmonhttpstatstable.EntityData.YFilter = rttmonhttpstatstable.YFilter
    rttmonhttpstatstable.EntityData.YangName = "rttMonHTTPStatsTable"
    rttmonhttpstatstable.EntityData.BundleName = "cisco_ios_xe"
    rttmonhttpstatstable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonhttpstatstable.EntityData.SegmentPath = "rttMonHTTPStatsTable"
    rttmonhttpstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonhttpstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonhttpstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonhttpstatstable.EntityData.Children = make(map[string]types.YChild)
    rttmonhttpstatstable.EntityData.Children["rttMonHTTPStatsEntry"] = types.YChild{"Rttmonhttpstatsentry", nil}
    for i := range rttmonhttpstatstable.Rttmonhttpstatsentry {
        rttmonhttpstatstable.EntityData.Children[types.GetSegmentPath(&rttmonhttpstatstable.Rttmonhttpstatsentry[i])] = types.YChild{"Rttmonhttpstatsentry", &rttmonhttpstatstable.Rttmonhttpstatsentry[i]}
    }
    rttmonhttpstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonhttpstatstable.EntityData)
}

// CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry is created only if the rttMonCtrlAdminRttType 
// is http. The operation of this table is same as that of
// rttMonStatsCaptureTable.
type CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry struct {
    EntityData types.CommonEntityData
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

func (rttmonhttpstatsentry *CISCORTTMONMIB_Rttmonhttpstatstable_Rttmonhttpstatsentry) GetEntityData() *types.CommonEntityData {
    rttmonhttpstatsentry.EntityData.YFilter = rttmonhttpstatsentry.YFilter
    rttmonhttpstatsentry.EntityData.YangName = "rttMonHTTPStatsEntry"
    rttmonhttpstatsentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonhttpstatsentry.EntityData.ParentYangName = "rttMonHTTPStatsTable"
    rttmonhttpstatsentry.EntityData.SegmentPath = "rttMonHTTPStatsEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonhttpstatsentry.Rttmonctrladminindex) + "']" + "[rttMonHTTPStatsStartTimeIndex='" + fmt.Sprintf("%v", rttmonhttpstatsentry.Rttmonhttpstatsstarttimeindex) + "']"
    rttmonhttpstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonhttpstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonhttpstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonhttpstatsentry.EntityData.Children = make(map[string]types.YChild)
    rttmonhttpstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonhttpstatsentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonhttpstatsentry.Rttmonctrladminindex}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsStartTimeIndex"] = types.YLeaf{"Rttmonhttpstatsstarttimeindex", rttmonhttpstatsentry.Rttmonhttpstatsstarttimeindex}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsCompletions"] = types.YLeaf{"Rttmonhttpstatscompletions", rttmonhttpstatsentry.Rttmonhttpstatscompletions}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsOverThresholds"] = types.YLeaf{"Rttmonhttpstatsoverthresholds", rttmonhttpstatsentry.Rttmonhttpstatsoverthresholds}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsRTTSum"] = types.YLeaf{"Rttmonhttpstatsrttsum", rttmonhttpstatsentry.Rttmonhttpstatsrttsum}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsRTTSum2Low"] = types.YLeaf{"Rttmonhttpstatsrttsum2Low", rttmonhttpstatsentry.Rttmonhttpstatsrttsum2Low}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsRTTSum2High"] = types.YLeaf{"Rttmonhttpstatsrttsum2High", rttmonhttpstatsentry.Rttmonhttpstatsrttsum2High}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsRTTMin"] = types.YLeaf{"Rttmonhttpstatsrttmin", rttmonhttpstatsentry.Rttmonhttpstatsrttmin}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsRTTMax"] = types.YLeaf{"Rttmonhttpstatsrttmax", rttmonhttpstatsentry.Rttmonhttpstatsrttmax}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsDNSRTTSum"] = types.YLeaf{"Rttmonhttpstatsdnsrttsum", rttmonhttpstatsentry.Rttmonhttpstatsdnsrttsum}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsTCPConnectRTTSum"] = types.YLeaf{"Rttmonhttpstatstcpconnectrttsum", rttmonhttpstatsentry.Rttmonhttpstatstcpconnectrttsum}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsTransactionRTTSum"] = types.YLeaf{"Rttmonhttpstatstransactionrttsum", rttmonhttpstatsentry.Rttmonhttpstatstransactionrttsum}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsMessageBodyOctetsSum"] = types.YLeaf{"Rttmonhttpstatsmessagebodyoctetssum", rttmonhttpstatsentry.Rttmonhttpstatsmessagebodyoctetssum}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsDNSServerTimeout"] = types.YLeaf{"Rttmonhttpstatsdnsservertimeout", rttmonhttpstatsentry.Rttmonhttpstatsdnsservertimeout}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsTCPConnectTimeout"] = types.YLeaf{"Rttmonhttpstatstcpconnecttimeout", rttmonhttpstatsentry.Rttmonhttpstatstcpconnecttimeout}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsTransactionTimeout"] = types.YLeaf{"Rttmonhttpstatstransactiontimeout", rttmonhttpstatsentry.Rttmonhttpstatstransactiontimeout}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsDNSQueryError"] = types.YLeaf{"Rttmonhttpstatsdnsqueryerror", rttmonhttpstatsentry.Rttmonhttpstatsdnsqueryerror}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsHTTPError"] = types.YLeaf{"Rttmonhttpstatshttperror", rttmonhttpstatsentry.Rttmonhttpstatshttperror}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsError"] = types.YLeaf{"Rttmonhttpstatserror", rttmonhttpstatsentry.Rttmonhttpstatserror}
    rttmonhttpstatsentry.EntityData.Leafs["rttMonHTTPStatsBusies"] = types.YLeaf{"Rttmonhttpstatsbusies", rttmonhttpstatsentry.Rttmonhttpstatsbusies}
    return &(rttmonhttpstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry is created only if the
    // rttMonCtrlAdminRttType  is jitter. The operation of this table is same as
    // that of rttMonStatsCaptureTable. The type is slice of
    // CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry.
    Rttmonjitterstatsentry []CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry
}

func (rttmonjitterstatstable *CISCORTTMONMIB_Rttmonjitterstatstable) GetEntityData() *types.CommonEntityData {
    rttmonjitterstatstable.EntityData.YFilter = rttmonjitterstatstable.YFilter
    rttmonjitterstatstable.EntityData.YangName = "rttMonJitterStatsTable"
    rttmonjitterstatstable.EntityData.BundleName = "cisco_ios_xe"
    rttmonjitterstatstable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonjitterstatstable.EntityData.SegmentPath = "rttMonJitterStatsTable"
    rttmonjitterstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonjitterstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonjitterstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonjitterstatstable.EntityData.Children = make(map[string]types.YChild)
    rttmonjitterstatstable.EntityData.Children["rttMonJitterStatsEntry"] = types.YChild{"Rttmonjitterstatsentry", nil}
    for i := range rttmonjitterstatstable.Rttmonjitterstatsentry {
        rttmonjitterstatstable.EntityData.Children[types.GetSegmentPath(&rttmonjitterstatstable.Rttmonjitterstatsentry[i])] = types.YChild{"Rttmonjitterstatsentry", &rttmonjitterstatstable.Rttmonjitterstatsentry[i]}
    }
    rttmonjitterstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonjitterstatstable.EntityData)
}

// CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry is created only if the rttMonCtrlAdminRttType 
// is jitter. The operation of this table is same as that of
// rttMonStatsCaptureTable.
type CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry struct {
    EntityData types.CommonEntityData
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

func (rttmonjitterstatsentry *CISCORTTMONMIB_Rttmonjitterstatstable_Rttmonjitterstatsentry) GetEntityData() *types.CommonEntityData {
    rttmonjitterstatsentry.EntityData.YFilter = rttmonjitterstatsentry.YFilter
    rttmonjitterstatsentry.EntityData.YangName = "rttMonJitterStatsEntry"
    rttmonjitterstatsentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonjitterstatsentry.EntityData.ParentYangName = "rttMonJitterStatsTable"
    rttmonjitterstatsentry.EntityData.SegmentPath = "rttMonJitterStatsEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonjitterstatsentry.Rttmonctrladminindex) + "']" + "[rttMonJitterStatsStartTimeIndex='" + fmt.Sprintf("%v", rttmonjitterstatsentry.Rttmonjitterstatsstarttimeindex) + "']"
    rttmonjitterstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonjitterstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonjitterstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonjitterstatsentry.EntityData.Children = make(map[string]types.YChild)
    rttmonjitterstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonjitterstatsentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonjitterstatsentry.Rttmonctrladminindex}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsStartTimeIndex"] = types.YLeaf{"Rttmonjitterstatsstarttimeindex", rttmonjitterstatsentry.Rttmonjitterstatsstarttimeindex}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsCompletions"] = types.YLeaf{"Rttmonjitterstatscompletions", rttmonjitterstatsentry.Rttmonjitterstatscompletions}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOverThresholds"] = types.YLeaf{"Rttmonjitterstatsoverthresholds", rttmonjitterstatsentry.Rttmonjitterstatsoverthresholds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsNumOfRTT"] = types.YLeaf{"Rttmonjitterstatsnumofrtt", rttmonjitterstatsentry.Rttmonjitterstatsnumofrtt}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsRTTSum"] = types.YLeaf{"Rttmonjitterstatsrttsum", rttmonjitterstatsentry.Rttmonjitterstatsrttsum}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsRTTSum2Low"] = types.YLeaf{"Rttmonjitterstatsrttsum2Low", rttmonjitterstatsentry.Rttmonjitterstatsrttsum2Low}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsRTTSum2High"] = types.YLeaf{"Rttmonjitterstatsrttsum2High", rttmonjitterstatsentry.Rttmonjitterstatsrttsum2High}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsRTTMin"] = types.YLeaf{"Rttmonjitterstatsrttmin", rttmonjitterstatsentry.Rttmonjitterstatsrttmin}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsRTTMax"] = types.YLeaf{"Rttmonjitterstatsrttmax", rttmonjitterstatsentry.Rttmonjitterstatsrttmax}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMinOfPositivesSD"] = types.YLeaf{"Rttmonjitterstatsminofpositivessd", rttmonjitterstatsentry.Rttmonjitterstatsminofpositivessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMaxOfPositivesSD"] = types.YLeaf{"Rttmonjitterstatsmaxofpositivessd", rttmonjitterstatsentry.Rttmonjitterstatsmaxofpositivessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsNumOfPositivesSD"] = types.YLeaf{"Rttmonjitterstatsnumofpositivessd", rttmonjitterstatsentry.Rttmonjitterstatsnumofpositivessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSumOfPositivesSD"] = types.YLeaf{"Rttmonjitterstatssumofpositivessd", rttmonjitterstatsentry.Rttmonjitterstatssumofpositivessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2PositivesSDLow"] = types.YLeaf{"Rttmonjitterstatssum2Positivessdlow", rttmonjitterstatsentry.Rttmonjitterstatssum2Positivessdlow}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2PositivesSDHigh"] = types.YLeaf{"Rttmonjitterstatssum2Positivessdhigh", rttmonjitterstatsentry.Rttmonjitterstatssum2Positivessdhigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMinOfNegativesSD"] = types.YLeaf{"Rttmonjitterstatsminofnegativessd", rttmonjitterstatsentry.Rttmonjitterstatsminofnegativessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMaxOfNegativesSD"] = types.YLeaf{"Rttmonjitterstatsmaxofnegativessd", rttmonjitterstatsentry.Rttmonjitterstatsmaxofnegativessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsNumOfNegativesSD"] = types.YLeaf{"Rttmonjitterstatsnumofnegativessd", rttmonjitterstatsentry.Rttmonjitterstatsnumofnegativessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSumOfNegativesSD"] = types.YLeaf{"Rttmonjitterstatssumofnegativessd", rttmonjitterstatsentry.Rttmonjitterstatssumofnegativessd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2NegativesSDLow"] = types.YLeaf{"Rttmonjitterstatssum2Negativessdlow", rttmonjitterstatsentry.Rttmonjitterstatssum2Negativessdlow}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2NegativesSDHigh"] = types.YLeaf{"Rttmonjitterstatssum2Negativessdhigh", rttmonjitterstatsentry.Rttmonjitterstatssum2Negativessdhigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMinOfPositivesDS"] = types.YLeaf{"Rttmonjitterstatsminofpositivesds", rttmonjitterstatsentry.Rttmonjitterstatsminofpositivesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMaxOfPositivesDS"] = types.YLeaf{"Rttmonjitterstatsmaxofpositivesds", rttmonjitterstatsentry.Rttmonjitterstatsmaxofpositivesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsNumOfPositivesDS"] = types.YLeaf{"Rttmonjitterstatsnumofpositivesds", rttmonjitterstatsentry.Rttmonjitterstatsnumofpositivesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSumOfPositivesDS"] = types.YLeaf{"Rttmonjitterstatssumofpositivesds", rttmonjitterstatsentry.Rttmonjitterstatssumofpositivesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2PositivesDSLow"] = types.YLeaf{"Rttmonjitterstatssum2Positivesdslow", rttmonjitterstatsentry.Rttmonjitterstatssum2Positivesdslow}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2PositivesDSHigh"] = types.YLeaf{"Rttmonjitterstatssum2Positivesdshigh", rttmonjitterstatsentry.Rttmonjitterstatssum2Positivesdshigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMinOfNegativesDS"] = types.YLeaf{"Rttmonjitterstatsminofnegativesds", rttmonjitterstatsentry.Rttmonjitterstatsminofnegativesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMaxOfNegativesDS"] = types.YLeaf{"Rttmonjitterstatsmaxofnegativesds", rttmonjitterstatsentry.Rttmonjitterstatsmaxofnegativesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsNumOfNegativesDS"] = types.YLeaf{"Rttmonjitterstatsnumofnegativesds", rttmonjitterstatsentry.Rttmonjitterstatsnumofnegativesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSumOfNegativesDS"] = types.YLeaf{"Rttmonjitterstatssumofnegativesds", rttmonjitterstatsentry.Rttmonjitterstatssumofnegativesds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2NegativesDSLow"] = types.YLeaf{"Rttmonjitterstatssum2Negativesdslow", rttmonjitterstatsentry.Rttmonjitterstatssum2Negativesdslow}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsSum2NegativesDSHigh"] = types.YLeaf{"Rttmonjitterstatssum2Negativesdshigh", rttmonjitterstatsentry.Rttmonjitterstatssum2Negativesdshigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsPacketLossSD"] = types.YLeaf{"Rttmonjitterstatspacketlosssd", rttmonjitterstatsentry.Rttmonjitterstatspacketlosssd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsPacketLossDS"] = types.YLeaf{"Rttmonjitterstatspacketlossds", rttmonjitterstatsentry.Rttmonjitterstatspacketlossds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsPacketOutOfSequence"] = types.YLeaf{"Rttmonjitterstatspacketoutofsequence", rttmonjitterstatsentry.Rttmonjitterstatspacketoutofsequence}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsPacketMIA"] = types.YLeaf{"Rttmonjitterstatspacketmia", rttmonjitterstatsentry.Rttmonjitterstatspacketmia}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsPacketLateArrival"] = types.YLeaf{"Rttmonjitterstatspacketlatearrival", rttmonjitterstatsentry.Rttmonjitterstatspacketlatearrival}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsError"] = types.YLeaf{"Rttmonjitterstatserror", rttmonjitterstatsentry.Rttmonjitterstatserror}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsBusies"] = types.YLeaf{"Rttmonjitterstatsbusies", rttmonjitterstatsentry.Rttmonjitterstatsbusies}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSumSD"] = types.YLeaf{"Rttmonjitterstatsowsumsd", rttmonjitterstatsentry.Rttmonjitterstatsowsumsd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSum2SDLow"] = types.YLeaf{"Rttmonjitterstatsowsum2Sdlow", rttmonjitterstatsentry.Rttmonjitterstatsowsum2Sdlow}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSum2SDHigh"] = types.YLeaf{"Rttmonjitterstatsowsum2Sdhigh", rttmonjitterstatsentry.Rttmonjitterstatsowsum2Sdhigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMinSD"] = types.YLeaf{"Rttmonjitterstatsowminsd", rttmonjitterstatsentry.Rttmonjitterstatsowminsd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMaxSD"] = types.YLeaf{"Rttmonjitterstatsowmaxsd", rttmonjitterstatsentry.Rttmonjitterstatsowmaxsd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSumDS"] = types.YLeaf{"Rttmonjitterstatsowsumds", rttmonjitterstatsentry.Rttmonjitterstatsowsumds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSum2DSLow"] = types.YLeaf{"Rttmonjitterstatsowsum2Dslow", rttmonjitterstatsentry.Rttmonjitterstatsowsum2Dslow}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSum2DSHigh"] = types.YLeaf{"Rttmonjitterstatsowsum2Dshigh", rttmonjitterstatsentry.Rttmonjitterstatsowsum2Dshigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMinDS"] = types.YLeaf{"Rttmonjitterstatsowminds", rttmonjitterstatsentry.Rttmonjitterstatsowminds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMaxDS"] = types.YLeaf{"Rttmonjitterstatsowmaxds", rttmonjitterstatsentry.Rttmonjitterstatsowmaxds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsNumOfOW"] = types.YLeaf{"Rttmonjitterstatsnumofow", rttmonjitterstatsentry.Rttmonjitterstatsnumofow}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMinSDNew"] = types.YLeaf{"Rttmonjitterstatsowminsdnew", rttmonjitterstatsentry.Rttmonjitterstatsowminsdnew}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMaxSDNew"] = types.YLeaf{"Rttmonjitterstatsowmaxsdnew", rttmonjitterstatsentry.Rttmonjitterstatsowmaxsdnew}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMinDSNew"] = types.YLeaf{"Rttmonjitterstatsowmindsnew", rttmonjitterstatsentry.Rttmonjitterstatsowmindsnew}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWMaxDSNew"] = types.YLeaf{"Rttmonjitterstatsowmaxdsnew", rttmonjitterstatsentry.Rttmonjitterstatsowmaxdsnew}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMinOfMOS"] = types.YLeaf{"Rttmonjitterstatsminofmos", rttmonjitterstatsentry.Rttmonjitterstatsminofmos}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMaxOfMOS"] = types.YLeaf{"Rttmonjitterstatsmaxofmos", rttmonjitterstatsentry.Rttmonjitterstatsmaxofmos}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMinOfICPIF"] = types.YLeaf{"Rttmonjitterstatsminoficpif", rttmonjitterstatsentry.Rttmonjitterstatsminoficpif}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsMaxOfICPIF"] = types.YLeaf{"Rttmonjitterstatsmaxoficpif", rttmonjitterstatsentry.Rttmonjitterstatsmaxoficpif}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsIAJOut"] = types.YLeaf{"Rttmonjitterstatsiajout", rttmonjitterstatsentry.Rttmonjitterstatsiajout}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsIAJIn"] = types.YLeaf{"Rttmonjitterstatsiajin", rttmonjitterstatsentry.Rttmonjitterstatsiajin}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsAvgJitter"] = types.YLeaf{"Rttmonjitterstatsavgjitter", rttmonjitterstatsentry.Rttmonjitterstatsavgjitter}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsAvgJitterSD"] = types.YLeaf{"Rttmonjitterstatsavgjittersd", rttmonjitterstatsentry.Rttmonjitterstatsavgjittersd}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsAvgJitterDS"] = types.YLeaf{"Rttmonjitterstatsavgjitterds", rttmonjitterstatsentry.Rttmonjitterstatsavgjitterds}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsUnSyncRTs"] = types.YLeaf{"Rttmonjitterstatsunsyncrts", rttmonjitterstatsentry.Rttmonjitterstatsunsyncrts}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsRTTSumHigh"] = types.YLeaf{"Rttmonjitterstatsrttsumhigh", rttmonjitterstatsentry.Rttmonjitterstatsrttsumhigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSumSDHigh"] = types.YLeaf{"Rttmonjitterstatsowsumsdhigh", rttmonjitterstatsentry.Rttmonjitterstatsowsumsdhigh}
    rttmonjitterstatsentry.EntityData.Leafs["rttMonJitterStatsOWSumDSHigh"] = types.YLeaf{"Rttmonjitterstatsowsumdshigh", rttmonjitterstatsentry.Rttmonjitterstatsowsumdshigh}
    return &(rttmonjitterstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonlpdgrpstatstable *CISCORTTMONMIB_Rttmonlpdgrpstatstable) GetEntityData() *types.CommonEntityData {
    rttmonlpdgrpstatstable.EntityData.YFilter = rttmonlpdgrpstatstable.YFilter
    rttmonlpdgrpstatstable.EntityData.YangName = "rttMonLpdGrpStatsTable"
    rttmonlpdgrpstatstable.EntityData.BundleName = "cisco_ios_xe"
    rttmonlpdgrpstatstable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonlpdgrpstatstable.EntityData.SegmentPath = "rttMonLpdGrpStatsTable"
    rttmonlpdgrpstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonlpdgrpstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonlpdgrpstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonlpdgrpstatstable.EntityData.Children = make(map[string]types.YChild)
    rttmonlpdgrpstatstable.EntityData.Children["rttMonLpdGrpStatsEntry"] = types.YChild{"Rttmonlpdgrpstatsentry", nil}
    for i := range rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry {
        rttmonlpdgrpstatstable.EntityData.Children[types.GetSegmentPath(&rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry[i])] = types.YChild{"Rttmonlpdgrpstatsentry", &rttmonlpdgrpstatstable.Rttmonlpdgrpstatsentry[i]}
    }
    rttmonlpdgrpstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonlpdgrpstatstable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonlpdgrpstatsentry *CISCORTTMONMIB_Rttmonlpdgrpstatstable_Rttmonlpdgrpstatsentry) GetEntityData() *types.CommonEntityData {
    rttmonlpdgrpstatsentry.EntityData.YFilter = rttmonlpdgrpstatsentry.YFilter
    rttmonlpdgrpstatsentry.EntityData.YangName = "rttMonLpdGrpStatsEntry"
    rttmonlpdgrpstatsentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonlpdgrpstatsentry.EntityData.ParentYangName = "rttMonLpdGrpStatsTable"
    rttmonlpdgrpstatsentry.EntityData.SegmentPath = "rttMonLpdGrpStatsEntry" + "[rttMonLpdGrpStatsGroupIndex='" + fmt.Sprintf("%v", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupindex) + "']" + "[rttMonLpdGrpStatsStartTimeIndex='" + fmt.Sprintf("%v", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsstarttimeindex) + "']"
    rttmonlpdgrpstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonlpdgrpstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonlpdgrpstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonlpdgrpstatsentry.EntityData.Children = make(map[string]types.YChild)
    rttmonlpdgrpstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsGroupIndex"] = types.YLeaf{"Rttmonlpdgrpstatsgroupindex", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupindex}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsStartTimeIndex"] = types.YLeaf{"Rttmonlpdgrpstatsstarttimeindex", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsstarttimeindex}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsTargetPE"] = types.YLeaf{"Rttmonlpdgrpstatstargetpe", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatstargetpe}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsNumOfPass"] = types.YLeaf{"Rttmonlpdgrpstatsnumofpass", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsnumofpass}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsNumOfFail"] = types.YLeaf{"Rttmonlpdgrpstatsnumoffail", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsnumoffail}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsNumOfTimeout"] = types.YLeaf{"Rttmonlpdgrpstatsnumoftimeout", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsnumoftimeout}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsAvgRTT"] = types.YLeaf{"Rttmonlpdgrpstatsavgrtt", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsavgrtt}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsMinRTT"] = types.YLeaf{"Rttmonlpdgrpstatsminrtt", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsminrtt}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsMaxRTT"] = types.YLeaf{"Rttmonlpdgrpstatsmaxrtt", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsmaxrtt}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsMinNumPaths"] = types.YLeaf{"Rttmonlpdgrpstatsminnumpaths", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsminnumpaths}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsMaxNumPaths"] = types.YLeaf{"Rttmonlpdgrpstatsmaxnumpaths", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsmaxnumpaths}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsLPDStartTime"] = types.YLeaf{"Rttmonlpdgrpstatslpdstarttime", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdstarttime}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsLPDFailOccurred"] = types.YLeaf{"Rttmonlpdgrpstatslpdfailoccurred", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdfailoccurred}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsLPDFailCause"] = types.YLeaf{"Rttmonlpdgrpstatslpdfailcause", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdfailcause}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsLPDCompTime"] = types.YLeaf{"Rttmonlpdgrpstatslpdcomptime", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatslpdcomptime}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsGroupStatus"] = types.YLeaf{"Rttmonlpdgrpstatsgroupstatus", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupstatus}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsGroupProbeIndex"] = types.YLeaf{"Rttmonlpdgrpstatsgroupprobeindex", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsgroupprobeindex}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsPathIds"] = types.YLeaf{"Rttmonlpdgrpstatspathids", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatspathids}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsProbeStatus"] = types.YLeaf{"Rttmonlpdgrpstatsprobestatus", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsprobestatus}
    rttmonlpdgrpstatsentry.EntityData.Leafs["rttMonLpdGrpStatsResetTime"] = types.YLeaf{"Rttmonlpdgrpstatsresettime", rttmonlpdgrpstatsentry.Rttmonlpdgrpstatsresettime}
    return &(rttmonlpdgrpstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonhistorycollectiontable *CISCORTTMONMIB_Rttmonhistorycollectiontable) GetEntityData() *types.CommonEntityData {
    rttmonhistorycollectiontable.EntityData.YFilter = rttmonhistorycollectiontable.YFilter
    rttmonhistorycollectiontable.EntityData.YangName = "rttMonHistoryCollectionTable"
    rttmonhistorycollectiontable.EntityData.BundleName = "cisco_ios_xe"
    rttmonhistorycollectiontable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonhistorycollectiontable.EntityData.SegmentPath = "rttMonHistoryCollectionTable"
    rttmonhistorycollectiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonhistorycollectiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonhistorycollectiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonhistorycollectiontable.EntityData.Children = make(map[string]types.YChild)
    rttmonhistorycollectiontable.EntityData.Children["rttMonHistoryCollectionEntry"] = types.YChild{"Rttmonhistorycollectionentry", nil}
    for i := range rttmonhistorycollectiontable.Rttmonhistorycollectionentry {
        rttmonhistorycollectiontable.EntityData.Children[types.GetSegmentPath(&rttmonhistorycollectiontable.Rttmonhistorycollectionentry[i])] = types.YChild{"Rttmonhistorycollectionentry", &rttmonhistorycollectiontable.Rttmonhistorycollectionentry[i]}
    }
    rttmonhistorycollectiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonhistorycollectiontable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (rttmonhistorycollectionentry *CISCORTTMONMIB_Rttmonhistorycollectiontable_Rttmonhistorycollectionentry) GetEntityData() *types.CommonEntityData {
    rttmonhistorycollectionentry.EntityData.YFilter = rttmonhistorycollectionentry.YFilter
    rttmonhistorycollectionentry.EntityData.YangName = "rttMonHistoryCollectionEntry"
    rttmonhistorycollectionentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonhistorycollectionentry.EntityData.ParentYangName = "rttMonHistoryCollectionTable"
    rttmonhistorycollectionentry.EntityData.SegmentPath = "rttMonHistoryCollectionEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonctrladminindex) + "']" + "[rttMonHistoryCollectionLifeIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonhistorycollectionlifeindex) + "']" + "[rttMonHistoryCollectionBucketIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonhistorycollectionbucketindex) + "']" + "[rttMonHistoryCollectionSampleIndex='" + fmt.Sprintf("%v", rttmonhistorycollectionentry.Rttmonhistorycollectionsampleindex) + "']"
    rttmonhistorycollectionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonhistorycollectionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonhistorycollectionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonhistorycollectionentry.EntityData.Children = make(map[string]types.YChild)
    rttmonhistorycollectionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonhistorycollectionentry.Rttmonctrladminindex}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionLifeIndex"] = types.YLeaf{"Rttmonhistorycollectionlifeindex", rttmonhistorycollectionentry.Rttmonhistorycollectionlifeindex}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionBucketIndex"] = types.YLeaf{"Rttmonhistorycollectionbucketindex", rttmonhistorycollectionentry.Rttmonhistorycollectionbucketindex}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionSampleIndex"] = types.YLeaf{"Rttmonhistorycollectionsampleindex", rttmonhistorycollectionentry.Rttmonhistorycollectionsampleindex}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionSampleTime"] = types.YLeaf{"Rttmonhistorycollectionsampletime", rttmonhistorycollectionentry.Rttmonhistorycollectionsampletime}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionAddress"] = types.YLeaf{"Rttmonhistorycollectionaddress", rttmonhistorycollectionentry.Rttmonhistorycollectionaddress}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionCompletionTime"] = types.YLeaf{"Rttmonhistorycollectioncompletiontime", rttmonhistorycollectionentry.Rttmonhistorycollectioncompletiontime}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionSense"] = types.YLeaf{"Rttmonhistorycollectionsense", rttmonhistorycollectionentry.Rttmonhistorycollectionsense}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionApplSpecificSense"] = types.YLeaf{"Rttmonhistorycollectionapplspecificsense", rttmonhistorycollectionentry.Rttmonhistorycollectionapplspecificsense}
    rttmonhistorycollectionentry.EntityData.Leafs["rttMonHistoryCollectionSenseDescription"] = types.YLeaf{"Rttmonhistorycollectionsensedescription", rttmonhistorycollectionentry.Rttmonhistorycollectionsensedescription}
    return &(rttmonhistorycollectionentry.EntityData)
}

// CISCORTTMONMIB_Rttmonlatesthttpopertable
// A table which contains the status of latest HTTP RTT
// operation.
type CISCORTTMONMIB_Rttmonlatesthttpopertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that record the latest HTTP RTT operation. This entry is
    // created automatically after the  rttMonCtrlAdminEntry is created. Also the
    // entry is  automatically deleted when rttMonCtrlAdminEntry is deleted. The
    // type is slice of
    // CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry.
    Rttmonlatesthttpoperentry []CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry
}

func (rttmonlatesthttpopertable *CISCORTTMONMIB_Rttmonlatesthttpopertable) GetEntityData() *types.CommonEntityData {
    rttmonlatesthttpopertable.EntityData.YFilter = rttmonlatesthttpopertable.YFilter
    rttmonlatesthttpopertable.EntityData.YangName = "rttMonLatestHTTPOperTable"
    rttmonlatesthttpopertable.EntityData.BundleName = "cisco_ios_xe"
    rttmonlatesthttpopertable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonlatesthttpopertable.EntityData.SegmentPath = "rttMonLatestHTTPOperTable"
    rttmonlatesthttpopertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonlatesthttpopertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonlatesthttpopertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonlatesthttpopertable.EntityData.Children = make(map[string]types.YChild)
    rttmonlatesthttpopertable.EntityData.Children["rttMonLatestHTTPOperEntry"] = types.YChild{"Rttmonlatesthttpoperentry", nil}
    for i := range rttmonlatesthttpopertable.Rttmonlatesthttpoperentry {
        rttmonlatesthttpopertable.EntityData.Children[types.GetSegmentPath(&rttmonlatesthttpopertable.Rttmonlatesthttpoperentry[i])] = types.YChild{"Rttmonlatesthttpoperentry", &rttmonlatesthttpopertable.Rttmonlatesthttpoperentry[i]}
    }
    rttmonlatesthttpopertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonlatesthttpopertable.EntityData)
}

// CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry
// A list of objects that record the latest HTTP RTT
// operation. This entry is created automatically after the 
// rttMonCtrlAdminEntry is created. Also the entry is 
// automatically deleted when rttMonCtrlAdminEntry is deleted.
type CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry struct {
    EntityData types.CommonEntityData
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

func (rttmonlatesthttpoperentry *CISCORTTMONMIB_Rttmonlatesthttpopertable_Rttmonlatesthttpoperentry) GetEntityData() *types.CommonEntityData {
    rttmonlatesthttpoperentry.EntityData.YFilter = rttmonlatesthttpoperentry.YFilter
    rttmonlatesthttpoperentry.EntityData.YangName = "rttMonLatestHTTPOperEntry"
    rttmonlatesthttpoperentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonlatesthttpoperentry.EntityData.ParentYangName = "rttMonLatestHTTPOperTable"
    rttmonlatesthttpoperentry.EntityData.SegmentPath = "rttMonLatestHTTPOperEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonlatesthttpoperentry.Rttmonctrladminindex) + "']"
    rttmonlatesthttpoperentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonlatesthttpoperentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonlatesthttpoperentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonlatesthttpoperentry.EntityData.Children = make(map[string]types.YChild)
    rttmonlatesthttpoperentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonlatesthttpoperentry.Rttmonctrladminindex}
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonLatestHTTPOperRTT"] = types.YLeaf{"Rttmonlatesthttpoperrtt", rttmonlatesthttpoperentry.Rttmonlatesthttpoperrtt}
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonLatestHTTPOperDNSRTT"] = types.YLeaf{"Rttmonlatesthttpoperdnsrtt", rttmonlatesthttpoperentry.Rttmonlatesthttpoperdnsrtt}
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonLatestHTTPOperTCPConnectRTT"] = types.YLeaf{"Rttmonlatesthttpopertcpconnectrtt", rttmonlatesthttpoperentry.Rttmonlatesthttpopertcpconnectrtt}
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonLatestHTTPOperTransactionRTT"] = types.YLeaf{"Rttmonlatesthttpopertransactionrtt", rttmonlatesthttpoperentry.Rttmonlatesthttpopertransactionrtt}
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonLatestHTTPOperMessageBodyOctets"] = types.YLeaf{"Rttmonlatesthttpopermessagebodyoctets", rttmonlatesthttpoperentry.Rttmonlatesthttpopermessagebodyoctets}
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonLatestHTTPOperSense"] = types.YLeaf{"Rttmonlatesthttpopersense", rttmonlatesthttpoperentry.Rttmonlatesthttpopersense}
    rttmonlatesthttpoperentry.EntityData.Leafs["rttMonLatestHTTPErrorSenseDescription"] = types.YLeaf{"Rttmonlatesthttperrorsensedescription", rttmonlatesthttpoperentry.Rttmonlatesthttperrorsensedescription}
    return &(rttmonlatesthttpoperentry.EntityData)
}

// CISCORTTMONMIB_Rttmonlatestjitteropertable
// A table which contains the status of latest Jitter
// operation.
type CISCORTTMONMIB_Rttmonlatestjitteropertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that record the latest Jitter operation. The type is
    // slice of
    // CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry.
    Rttmonlatestjitteroperentry []CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry
}

func (rttmonlatestjitteropertable *CISCORTTMONMIB_Rttmonlatestjitteropertable) GetEntityData() *types.CommonEntityData {
    rttmonlatestjitteropertable.EntityData.YFilter = rttmonlatestjitteropertable.YFilter
    rttmonlatestjitteropertable.EntityData.YangName = "rttMonLatestJitterOperTable"
    rttmonlatestjitteropertable.EntityData.BundleName = "cisco_ios_xe"
    rttmonlatestjitteropertable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttmonlatestjitteropertable.EntityData.SegmentPath = "rttMonLatestJitterOperTable"
    rttmonlatestjitteropertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonlatestjitteropertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonlatestjitteropertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonlatestjitteropertable.EntityData.Children = make(map[string]types.YChild)
    rttmonlatestjitteropertable.EntityData.Children["rttMonLatestJitterOperEntry"] = types.YChild{"Rttmonlatestjitteroperentry", nil}
    for i := range rttmonlatestjitteropertable.Rttmonlatestjitteroperentry {
        rttmonlatestjitteropertable.EntityData.Children[types.GetSegmentPath(&rttmonlatestjitteropertable.Rttmonlatestjitteroperentry[i])] = types.YChild{"Rttmonlatestjitteroperentry", &rttmonlatestjitteropertable.Rttmonlatestjitteroperentry[i]}
    }
    rttmonlatestjitteropertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rttmonlatestjitteropertable.EntityData)
}

// CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry
// A list of objects that record the latest Jitter
// operation.
type CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry struct {
    EntityData types.CommonEntityData
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

func (rttmonlatestjitteroperentry *CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry) GetEntityData() *types.CommonEntityData {
    rttmonlatestjitteroperentry.EntityData.YFilter = rttmonlatestjitteroperentry.YFilter
    rttmonlatestjitteroperentry.EntityData.YangName = "rttMonLatestJitterOperEntry"
    rttmonlatestjitteroperentry.EntityData.BundleName = "cisco_ios_xe"
    rttmonlatestjitteroperentry.EntityData.ParentYangName = "rttMonLatestJitterOperTable"
    rttmonlatestjitteroperentry.EntityData.SegmentPath = "rttMonLatestJitterOperEntry" + "[rttMonCtrlAdminIndex='" + fmt.Sprintf("%v", rttmonlatestjitteroperentry.Rttmonctrladminindex) + "']"
    rttmonlatestjitteroperentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttmonlatestjitteroperentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttmonlatestjitteroperentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttmonlatestjitteroperentry.EntityData.Children = make(map[string]types.YChild)
    rttmonlatestjitteroperentry.EntityData.Leafs = make(map[string]types.YLeaf)
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonCtrlAdminIndex"] = types.YLeaf{"Rttmonctrladminindex", rttmonlatestjitteroperentry.Rttmonctrladminindex}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperNumOfRTT"] = types.YLeaf{"Rttmonlatestjitteropernumofrtt", rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofrtt}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperRTTSum"] = types.YLeaf{"Rttmonlatestjitteroperrttsum", rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsum}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperRTTSum2"] = types.YLeaf{"Rttmonlatestjitteroperrttsum2", rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsum2}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperRTTMin"] = types.YLeaf{"Rttmonlatestjitteroperrttmin", rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttmin}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperRTTMax"] = types.YLeaf{"Rttmonlatestjitteroperrttmax", rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttmax}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMinOfPositivesSD"] = types.YLeaf{"Rttmonlatestjitteroperminofpositivessd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofpositivessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMaxOfPositivesSD"] = types.YLeaf{"Rttmonlatestjitteropermaxofpositivessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofpositivessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperNumOfPositivesSD"] = types.YLeaf{"Rttmonlatestjitteropernumofpositivessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofpositivessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSumOfPositivesSD"] = types.YLeaf{"Rttmonlatestjitteropersumofpositivessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofpositivessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSum2PositivesSD"] = types.YLeaf{"Rttmonlatestjitteropersum2Positivessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Positivessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMinOfNegativesSD"] = types.YLeaf{"Rttmonlatestjitteroperminofnegativessd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofnegativessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMaxOfNegativesSD"] = types.YLeaf{"Rttmonlatestjitteropermaxofnegativessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofnegativessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperNumOfNegativesSD"] = types.YLeaf{"Rttmonlatestjitteropernumofnegativessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofnegativessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSumOfNegativesSD"] = types.YLeaf{"Rttmonlatestjitteropersumofnegativessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofnegativessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSum2NegativesSD"] = types.YLeaf{"Rttmonlatestjitteropersum2Negativessd", rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Negativessd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMinOfPositivesDS"] = types.YLeaf{"Rttmonlatestjitteroperminofpositivesds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofpositivesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMaxOfPositivesDS"] = types.YLeaf{"Rttmonlatestjitteropermaxofpositivesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofpositivesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperNumOfPositivesDS"] = types.YLeaf{"Rttmonlatestjitteropernumofpositivesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofpositivesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSumOfPositivesDS"] = types.YLeaf{"Rttmonlatestjitteropersumofpositivesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofpositivesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSum2PositivesDS"] = types.YLeaf{"Rttmonlatestjitteropersum2Positivesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Positivesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMinOfNegativesDS"] = types.YLeaf{"Rttmonlatestjitteroperminofnegativesds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperminofnegativesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMaxOfNegativesDS"] = types.YLeaf{"Rttmonlatestjitteropermaxofnegativesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropermaxofnegativesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperNumOfNegativesDS"] = types.YLeaf{"Rttmonlatestjitteropernumofnegativesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofnegativesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSumOfNegativesDS"] = types.YLeaf{"Rttmonlatestjitteropersumofnegativesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropersumofnegativesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSum2NegativesDS"] = types.YLeaf{"Rttmonlatestjitteropersum2Negativesds", rttmonlatestjitteroperentry.Rttmonlatestjitteropersum2Negativesds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperPacketLossSD"] = types.YLeaf{"Rttmonlatestjitteroperpacketlosssd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketlosssd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperPacketLossDS"] = types.YLeaf{"Rttmonlatestjitteroperpacketlossds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketlossds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperPacketOutOfSequence"] = types.YLeaf{"Rttmonlatestjitteroperpacketoutofsequence", rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketoutofsequence}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperPacketMIA"] = types.YLeaf{"Rttmonlatestjitteroperpacketmia", rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketmia}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperPacketLateArrival"] = types.YLeaf{"Rttmonlatestjitteroperpacketlatearrival", rttmonlatestjitteroperentry.Rttmonlatestjitteroperpacketlatearrival}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperSense"] = types.YLeaf{"Rttmonlatestjitteropersense", rttmonlatestjitteroperentry.Rttmonlatestjitteropersense}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterErrorSenseDescription"] = types.YLeaf{"Rttmonlatestjittererrorsensedescription", rttmonlatestjitteroperentry.Rttmonlatestjittererrorsensedescription}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSumSD"] = types.YLeaf{"Rttmonlatestjitteroperowsumsd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumsd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSum2SD"] = types.YLeaf{"Rttmonlatestjitteroperowsum2Sd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Sd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWMinSD"] = types.YLeaf{"Rttmonlatestjitteroperowminsd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowminsd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWMaxSD"] = types.YLeaf{"Rttmonlatestjitteroperowmaxsd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowmaxsd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSumDS"] = types.YLeaf{"Rttmonlatestjitteroperowsumds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSum2DS"] = types.YLeaf{"Rttmonlatestjitteroperowsum2Ds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Ds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWMinDS"] = types.YLeaf{"Rttmonlatestjitteroperowminds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowminds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWMaxDS"] = types.YLeaf{"Rttmonlatestjitteroperowmaxds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowmaxds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperNumOfOW"] = types.YLeaf{"Rttmonlatestjitteropernumofow", rttmonlatestjitteroperentry.Rttmonlatestjitteropernumofow}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperMOS"] = types.YLeaf{"Rttmonlatestjitteropermos", rttmonlatestjitteroperentry.Rttmonlatestjitteropermos}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperICPIF"] = types.YLeaf{"Rttmonlatestjitteropericpif", rttmonlatestjitteroperentry.Rttmonlatestjitteropericpif}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperIAJOut"] = types.YLeaf{"Rttmonlatestjitteroperiajout", rttmonlatestjitteroperentry.Rttmonlatestjitteroperiajout}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperIAJIn"] = types.YLeaf{"Rttmonlatestjitteroperiajin", rttmonlatestjitteroperentry.Rttmonlatestjitteroperiajin}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperAvgJitter"] = types.YLeaf{"Rttmonlatestjitteroperavgjitter", rttmonlatestjitteroperentry.Rttmonlatestjitteroperavgjitter}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperAvgSDJ"] = types.YLeaf{"Rttmonlatestjitteroperavgsdj", rttmonlatestjitteroperentry.Rttmonlatestjitteroperavgsdj}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperAvgDSJ"] = types.YLeaf{"Rttmonlatestjitteroperavgdsj", rttmonlatestjitteroperentry.Rttmonlatestjitteroperavgdsj}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWAvgSD"] = types.YLeaf{"Rttmonlatestjitteroperowavgsd", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowavgsd}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWAvgDS"] = types.YLeaf{"Rttmonlatestjitteroperowavgds", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowavgds}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperNTPState"] = types.YLeaf{"Rttmonlatestjitteroperntpstate", rttmonlatestjitteroperentry.Rttmonlatestjitteroperntpstate}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperUnSyncRTs"] = types.YLeaf{"Rttmonlatestjitteroperunsyncrts", rttmonlatestjitteroperentry.Rttmonlatestjitteroperunsyncrts}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperRTTSumHigh"] = types.YLeaf{"Rttmonlatestjitteroperrttsumhigh", rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsumhigh}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperRTTSum2High"] = types.YLeaf{"Rttmonlatestjitteroperrttsum2High", rttmonlatestjitteroperentry.Rttmonlatestjitteroperrttsum2High}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSumSDHigh"] = types.YLeaf{"Rttmonlatestjitteroperowsumsdhigh", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumsdhigh}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSum2SDHigh"] = types.YLeaf{"Rttmonlatestjitteroperowsum2Sdhigh", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Sdhigh}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSumDSHigh"] = types.YLeaf{"Rttmonlatestjitteroperowsumdshigh", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsumdshigh}
    rttmonlatestjitteroperentry.EntityData.Leafs["rttMonLatestJitterOperOWSum2DSHigh"] = types.YLeaf{"Rttmonlatestjitteroperowsum2Dshigh", rttmonlatestjitteroperentry.Rttmonlatestjitteroperowsum2Dshigh}
    return &(rttmonlatestjitteroperentry.EntityData)
}

// CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate represents on sender and responder is within configured tolerance level.
type CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate string

const (
    CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate_sync CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate = "sync"

    CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate_outOfSync CISCORTTMONMIB_Rttmonlatestjitteropertable_Rttmonlatestjitteroperentry_Rttmonlatestjitteroperntpstate = "outOfSync"
)

