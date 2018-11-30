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

    
    RttMonAppl CISCORTTMONMIB_RttMonAppl

    // A table of which contains the supported Rtt Monitor Types.  See the
    // RttMonRttType textual convention for the definition of each type.
    RttMonApplSupportedRttTypesTable CISCORTTMONMIB_RttMonApplSupportedRttTypesTable

    // A table of which contains the supported Rtt Monitor Protocols.  See the
    // RttMonProtocol textual convention  for the definition of each protocol.
    RttMonApplSupportedProtocolsTable CISCORTTMONMIB_RttMonApplSupportedProtocolsTable

    // A table of which contains the previously configured Script Names and File
    // IO targets.  These Script Names and File IO targets are installed via a
    // different mechanism than this application, and are specific to each
    // platform.
    RttMonApplPreConfigedTable CISCORTTMONMIB_RttMonApplPreConfigedTable

    // A table which contains the definitions for key-strings that will be used in
    // authenticating RTR Control Protocol.
    RttMonApplAuthTable CISCORTTMONMIB_RttMonApplAuthTable

    // A table of Round Trip Time (RTT) monitoring definitions.  The RTT
    // administration control is in multiple tables.   This first table, is used
    // to create a conceptual RTT  control row.  The following tables contain
    // objects which  configure scheduling, information gathering, and 
    // notification/trigger generation.  All of these tables  will create the same
    // conceptual RTT control row as this  table using this tables' index as their
    // own index.   This table is limited in size by the agent  implementation. 
    // The object rttMonApplNumCtrlAdminEntry will reflect this tables maximum
    // number of entries.
    RttMonCtrlAdminTable CISCORTTMONMIB_RttMonCtrlAdminTable

    // A table that contains Round Trip Time (RTT) specific definitions.  This
    // table is controlled via the  rttMonCtrlAdminTable.  Entries in this table
    // are created via the rttMonCtrlAdminStatus object.
    RttMonEchoAdminTable CISCORTTMONMIB_RttMonEchoAdminTable

    // A table of Round Trip Time (RTT) monitoring 'fileIO' specific definitions. 
    // When the RttMonRttType is not 'fileIO' this table is not valid.  This table
    // is controlled via the  rttMonCtrlAdminTable.  Entries in this table are
    // created via the rttMonCtrlAdminStatus object.
    RttMonFileIOAdminTable CISCORTTMONMIB_RttMonFileIOAdminTable

    // A table of Round Trip Time (RTT) monitoring 'script' specific definitions. 
    // When the RttMonRttType is not 'script' this table is not valid.  This table
    // is controlled via the rttMonCtrlAdminTable.  Entries in this table are
    // created via the rttMonCtrlAdminStatus object.
    RttMonScriptAdminTable CISCORTTMONMIB_RttMonScriptAdminTable

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
    RttMonReactTriggerAdminTable CISCORTTMONMIB_RttMonReactTriggerAdminTable

    // A table to store the hop addresses in a Loose Source Routing path. Response
    // times are computed along the specified path using ping.  This maximum table
    // size is limited by the size of the  maximum number of hop addresses that
    // can fit in an IP header, which is 8. The object rttMonEchoPathAdminEntry
    // will reflect  this tables maximum number of entries.  This table is coupled
    // with rttMonCtrlAdminStatus.
    RttMonEchoPathAdminTable CISCORTTMONMIB_RttMonEchoPathAdminTable

    // A table of Round Trip Time (RTT) monitoring group scheduling specific
    // definitions. This table is used to create a conceptual group scheduling
    // control row. The entries in this control row contain objects used to define
    // group schedule configuration parameters.  The objects of this table will be
    // used to schedule a group of probes identified by the conceptual rows of the
    // rttMonCtrlAdminTable.
    RttMonGrpScheduleAdminTable CISCORTTMONMIB_RttMonGrpScheduleAdminTable

    // A table of Auto SAA L3 MPLS VPN definitions.  The Auto SAA L3 MPLS VPN
    // administration control is in multiple tables.  This first table, is used to
    // create a conceptual Auto SAA L3 MPLS VPN control row.  The following tables
    // contain objects which used in type specific configurations, scheduling and
    // reaction configurations. All of these tables will create the same
    // conceptual control row as this table using this table's index as their own
    // index.  In order to a row in this table to become active the following
    // objects must be defined.   rttMplsVpnMonCtrlRttType,  
    // rttMplsVpnMonCtrlVrfName and   rttMplsVpnMonSchedulePeriod.
    RttMplsVpnMonCtrlTable CISCORTTMONMIB_RttMplsVpnMonCtrlTable

    // A table that contains the reaction configurations. Each conceptual row in
    // rttMonReactTable corresponds to a reaction configured for the probe defined
    // in rttMonCtrlAdminTable.  For each reaction configured for a probe there is
    // an entry in the table.  Each Probe can have multiple reactions and hence
    // there can be multiple rows for a particular probe.  This table is coupled
    // with rttMonCtrlAdminTable.
    RttMonReactTable CISCORTTMONMIB_RttMonReactTable

    // This table contains information about the generated operation id as part of
    // a parent IP SLA operation. The parent operation id is pseudo-random number,
    // selected by the management  station based on an operation started by the
    // management  station,when creating a row via the rttMonCtrlAdminStatus
    // object in the rttMonCtrlAdminTable table.
    RttMonGeneratedOperTable CISCORTTMONMIB_RttMonGeneratedOperTable

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
    RttMonStatsCaptureTable CISCORTTMONMIB_RttMonStatsCaptureTable

    // The statistics collection database.  This table has the exact same behavior
    // as the rttMonStatsCaptureTable, except it does not keep statistical
    // distribution information.  For a complete table description see the
    // rttMonStatsCaptureTable object.
    RttMonStatsCollectTable CISCORTTMONMIB_RttMonStatsCollectTable

    // The statistics totals database.  This table has the exact same behavior as
    // the rttMonStatsCaptureTable, except it only keeps 60 minute group values. 
    // For a complete table description see the rttMonStatsCaptureTable object.
    RttMonStatsTotalsTable CISCORTTMONMIB_RttMonStatsTotalsTable

    // The HTTP statistics collection database.  The HTTP statistics table
    // contains summarized information of the results for a conceptual RTT control
    // row. A rolling accumulated history of this information is maintained in a 
    // series of hourly 'group(s)'.  The operation of this table is same as that
    // of  rttMonStatsCaptureTable, except that this table can only  store a
    // maximum of 2 hours of data.
    RttMonHTTPStatsTable CISCORTTMONMIB_RttMonHTTPStatsTable

    // The Jitter statistics collection database.  The Jitter statistics table
    // contains summarized information of the results for a conceptual RTT control
    // row. A rolling accumulated history of this information is maintained in a 
    // series of hourly 'group(s)'.  The operation of this table is same as that
    // of  rttMonStatsCaptureTable, except that this table will store  2 hours of
    // data.
    RttMonJitterStatsTable CISCORTTMONMIB_RttMonJitterStatsTable

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
    RttMonLpdGrpStatsTable CISCORTTMONMIB_RttMonLpdGrpStatsTable

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
    RttMonHistoryCollectionTable CISCORTTMONMIB_RttMonHistoryCollectionTable

    // A table which contains the status of latest HTTP RTT operation.
    RttMonLatestHTTPOperTable CISCORTTMONMIB_RttMonLatestHTTPOperTable

    // A table which contains the status of latest Jitter operation.
    RttMonLatestJitterOperTable CISCORTTMONMIB_RttMonLatestJitterOperTable
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

    cISCORTTMONMIB.EntityData.Children = types.NewOrderedMap()
    cISCORTTMONMIB.EntityData.Children.Append("rttMonAppl", types.YChild{"RttMonAppl", &cISCORTTMONMIB.RttMonAppl})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonApplSupportedRttTypesTable", types.YChild{"RttMonApplSupportedRttTypesTable", &cISCORTTMONMIB.RttMonApplSupportedRttTypesTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonApplSupportedProtocolsTable", types.YChild{"RttMonApplSupportedProtocolsTable", &cISCORTTMONMIB.RttMonApplSupportedProtocolsTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonApplPreConfigedTable", types.YChild{"RttMonApplPreConfigedTable", &cISCORTTMONMIB.RttMonApplPreConfigedTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonApplAuthTable", types.YChild{"RttMonApplAuthTable", &cISCORTTMONMIB.RttMonApplAuthTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonCtrlAdminTable", types.YChild{"RttMonCtrlAdminTable", &cISCORTTMONMIB.RttMonCtrlAdminTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonEchoAdminTable", types.YChild{"RttMonEchoAdminTable", &cISCORTTMONMIB.RttMonEchoAdminTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonFileIOAdminTable", types.YChild{"RttMonFileIOAdminTable", &cISCORTTMONMIB.RttMonFileIOAdminTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonScriptAdminTable", types.YChild{"RttMonScriptAdminTable", &cISCORTTMONMIB.RttMonScriptAdminTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonReactTriggerAdminTable", types.YChild{"RttMonReactTriggerAdminTable", &cISCORTTMONMIB.RttMonReactTriggerAdminTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonEchoPathAdminTable", types.YChild{"RttMonEchoPathAdminTable", &cISCORTTMONMIB.RttMonEchoPathAdminTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonGrpScheduleAdminTable", types.YChild{"RttMonGrpScheduleAdminTable", &cISCORTTMONMIB.RttMonGrpScheduleAdminTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMplsVpnMonCtrlTable", types.YChild{"RttMplsVpnMonCtrlTable", &cISCORTTMONMIB.RttMplsVpnMonCtrlTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonReactTable", types.YChild{"RttMonReactTable", &cISCORTTMONMIB.RttMonReactTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonGeneratedOperTable", types.YChild{"RttMonGeneratedOperTable", &cISCORTTMONMIB.RttMonGeneratedOperTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonStatsCaptureTable", types.YChild{"RttMonStatsCaptureTable", &cISCORTTMONMIB.RttMonStatsCaptureTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonStatsCollectTable", types.YChild{"RttMonStatsCollectTable", &cISCORTTMONMIB.RttMonStatsCollectTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonStatsTotalsTable", types.YChild{"RttMonStatsTotalsTable", &cISCORTTMONMIB.RttMonStatsTotalsTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonHTTPStatsTable", types.YChild{"RttMonHTTPStatsTable", &cISCORTTMONMIB.RttMonHTTPStatsTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonJitterStatsTable", types.YChild{"RttMonJitterStatsTable", &cISCORTTMONMIB.RttMonJitterStatsTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonLpdGrpStatsTable", types.YChild{"RttMonLpdGrpStatsTable", &cISCORTTMONMIB.RttMonLpdGrpStatsTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonHistoryCollectionTable", types.YChild{"RttMonHistoryCollectionTable", &cISCORTTMONMIB.RttMonHistoryCollectionTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonLatestHTTPOperTable", types.YChild{"RttMonLatestHTTPOperTable", &cISCORTTMONMIB.RttMonLatestHTTPOperTable})
    cISCORTTMONMIB.EntityData.Children.Append("rttMonLatestJitterOperTable", types.YChild{"RttMonLatestJitterOperTable", &cISCORTTMONMIB.RttMonLatestJitterOperTable})
    cISCORTTMONMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCORTTMONMIB.EntityData.YListKeys = []string {}

    return &(cISCORTTMONMIB.EntityData)
}

// CISCORTTMONMIB_RttMonAppl
type CISCORTTMONMIB_RttMonAppl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Round Trip Time monitoring application version string.  The format will be:
    // 'Version.Release.Patch-Level: Textual-Description'  For example:  '1.0.0:
    // Initial RTT Application'. The type is string.
    RttMonApplVersion interface{}

    // The maximum size of the data portion an echo packet supported by this RTT
    // application.  This is the maximum value that can be specified by
    // (rttMonEchoAdminPktDataRequestSize + ARR Header) or 
    // (rttMonEchoAdminPktDataResponseSize + ARR Header) in the
    // rttMonCtrlAdminTable.  This object is undefined for conceptual RTT  control
    // rows when the RttMonRttType object is set to 'fileIO' or 'script'. The type
    // is interface{} with range: 0..16384. Units are octets.
    RttMonApplMaxPacketDataSize interface{}

    // The last time at which a set operation occurred on any of the objects in
    // this MIB.  The managing  application can inspect this value in order to 
    // determine whether changes have been made without  retrieving the entire
    // Administration portion of this MIB.  This object applies to all settable
    // objects in this MIB, including the 'Reset' objects that could clear saved
    // history/statistics. The type is interface{} with range: 0..4294967295.
    RttMonApplTimeOfLastSet interface{}

    // This object defines the maximum number of entries that can be added to the
    // rttMonCtrlAdminTable. It is calculated at the system init time. The value
    // is impacted when rttMonApplFreeMemLowWaterMark is changed. The type is
    // interface{} with range: 1..2147483647.
    RttMonApplNumCtrlAdminEntry interface{}

    // When set to 'reset' the entire RTT application goes through a reset
    // sequence, making a best  effort to revert to its startup condition.  Any 
    // and all rows in the Overall Control Group will be immediately deleted,
    // together with any associated rows in the Statistics Collection Group, and 
    // History Collection Group.  All open connections  will also be closed. 
    // Finally the  rttMonApplPreConfigedTable will reset (see 
    // rttMonApplPreConfigedReset). The type is RttReset.
    RttMonApplReset interface{}

    // When set to 'reset' the RTT application will reset the Application
    // Preconfigured MIB section.  This will force the RTT application to delete
    // all entries in the rttMonApplPreConfigedTable and then to repopulate the
    // table with the current configuration.  This provides a mechanism to load
    // and unload user scripts and file paths. The type is RttReset.
    RttMonApplPreConfigedReset interface{}

    // This object defines the number of new probes that can be configured on a
    // router. The number depends on the value  of rttMonApplFreeMemLowWaterMark,
    // free bytes available on the router and the system configured
    // rttMonCtrlAdminEntry number. Equation: rttMonApplProbeCapacity = 
    // MIN(((Free_Bytes_on_the_Router - rttMonApplFreeMemLowWaterMark)/
    // Memory_required_by_each_probe), rttMonApplNumCtrlAdminEntry - 
    // Num_of_Probes_already_configured)). The type is interface{} with range:
    // 1..2147483647.
    RttMonApplProbeCapacity interface{}

    // This object defines the amount of free memory a router must have in order
    // to configure RTR. If RTR found out that the memory is falling below this
    // mark, it will not allow new probes to be configured.  This value should not
    // be set higher (or very close to) than  the free bytes available on the
    // router. The type is interface{} with range: 0..2147483647.
    RttMonApplFreeMemLowWaterMark interface{}

    // An error description for the last error message caused by set.  Currently,
    // it includes set error caused due to setting rttMonApplFreeMemLowWaterMark
    // greater than the available free memory on the router or not enough memory
    // left to create new probes. The type is string.
    RttMonApplLatestSetError interface{}

    // Enable or disable RTR responder on the router. The type is bool.
    RttMonApplResponder interface{}

    // This object is used to reset certain objects within the
    // rttMonLpdGrpStatsTable.  When the object is set to value of an active LPD
    // Group identifier the associated objects will be reset. The reset objects
    // will be set to a value as specified in the object's description.  The
    // following objects will not be reset. - rttMonLpdGrpStatsTargetPE -
    // rttMonLpdGrpStatsGroupProbeIndex - rttMonLpdGrpStatsGroupIndex -
    // rttMonLpdGrpStatsStartTimeIndex. The type is interface{} with range:
    // 0..2147483647.
    RttMonApplLpdGrpStatsReset interface{}
}

func (rttMonAppl *CISCORTTMONMIB_RttMonAppl) GetEntityData() *types.CommonEntityData {
    rttMonAppl.EntityData.YFilter = rttMonAppl.YFilter
    rttMonAppl.EntityData.YangName = "rttMonAppl"
    rttMonAppl.EntityData.BundleName = "cisco_ios_xe"
    rttMonAppl.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonAppl.EntityData.SegmentPath = "rttMonAppl"
    rttMonAppl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonAppl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonAppl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonAppl.EntityData.Children = types.NewOrderedMap()
    rttMonAppl.EntityData.Leafs = types.NewOrderedMap()
    rttMonAppl.EntityData.Leafs.Append("rttMonApplVersion", types.YLeaf{"RttMonApplVersion", rttMonAppl.RttMonApplVersion})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplMaxPacketDataSize", types.YLeaf{"RttMonApplMaxPacketDataSize", rttMonAppl.RttMonApplMaxPacketDataSize})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplTimeOfLastSet", types.YLeaf{"RttMonApplTimeOfLastSet", rttMonAppl.RttMonApplTimeOfLastSet})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplNumCtrlAdminEntry", types.YLeaf{"RttMonApplNumCtrlAdminEntry", rttMonAppl.RttMonApplNumCtrlAdminEntry})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplReset", types.YLeaf{"RttMonApplReset", rttMonAppl.RttMonApplReset})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplPreConfigedReset", types.YLeaf{"RttMonApplPreConfigedReset", rttMonAppl.RttMonApplPreConfigedReset})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplProbeCapacity", types.YLeaf{"RttMonApplProbeCapacity", rttMonAppl.RttMonApplProbeCapacity})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplFreeMemLowWaterMark", types.YLeaf{"RttMonApplFreeMemLowWaterMark", rttMonAppl.RttMonApplFreeMemLowWaterMark})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplLatestSetError", types.YLeaf{"RttMonApplLatestSetError", rttMonAppl.RttMonApplLatestSetError})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplResponder", types.YLeaf{"RttMonApplResponder", rttMonAppl.RttMonApplResponder})
    rttMonAppl.EntityData.Leafs.Append("rttMonApplLpdGrpStatsReset", types.YLeaf{"RttMonApplLpdGrpStatsReset", rttMonAppl.RttMonApplLpdGrpStatsReset})

    rttMonAppl.EntityData.YListKeys = []string {}

    return &(rttMonAppl.EntityData)
}

// CISCORTTMONMIB_RttMonApplSupportedRttTypesTable
// A table of which contains the supported Rtt
// Monitor Types.
// 
// See the RttMonRttType textual convention for
// the definition of each type.
type CISCORTTMONMIB_RttMonApplSupportedRttTypesTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list that presents the valid Rtt Monitor Types. The type is slice of
    // CISCORTTMONMIB_RttMonApplSupportedRttTypesTable_RttMonApplSupportedRttTypesEntry.
    RttMonApplSupportedRttTypesEntry []*CISCORTTMONMIB_RttMonApplSupportedRttTypesTable_RttMonApplSupportedRttTypesEntry
}

func (rttMonApplSupportedRttTypesTable *CISCORTTMONMIB_RttMonApplSupportedRttTypesTable) GetEntityData() *types.CommonEntityData {
    rttMonApplSupportedRttTypesTable.EntityData.YFilter = rttMonApplSupportedRttTypesTable.YFilter
    rttMonApplSupportedRttTypesTable.EntityData.YangName = "rttMonApplSupportedRttTypesTable"
    rttMonApplSupportedRttTypesTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplSupportedRttTypesTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonApplSupportedRttTypesTable.EntityData.SegmentPath = "rttMonApplSupportedRttTypesTable"
    rttMonApplSupportedRttTypesTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplSupportedRttTypesTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplSupportedRttTypesTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplSupportedRttTypesTable.EntityData.Children = types.NewOrderedMap()
    rttMonApplSupportedRttTypesTable.EntityData.Children.Append("rttMonApplSupportedRttTypesEntry", types.YChild{"RttMonApplSupportedRttTypesEntry", nil})
    for i := range rttMonApplSupportedRttTypesTable.RttMonApplSupportedRttTypesEntry {
        rttMonApplSupportedRttTypesTable.EntityData.Children.Append(types.GetSegmentPath(rttMonApplSupportedRttTypesTable.RttMonApplSupportedRttTypesEntry[i]), types.YChild{"RttMonApplSupportedRttTypesEntry", rttMonApplSupportedRttTypesTable.RttMonApplSupportedRttTypesEntry[i]})
    }
    rttMonApplSupportedRttTypesTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonApplSupportedRttTypesTable.EntityData.YListKeys = []string {}

    return &(rttMonApplSupportedRttTypesTable.EntityData)
}

// CISCORTTMONMIB_RttMonApplSupportedRttTypesTable_RttMonApplSupportedRttTypesEntry
// A list that presents the valid Rtt Monitor
// Types.
type CISCORTTMONMIB_RttMonApplSupportedRttTypesTable_RttMonApplSupportedRttTypesEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object indexes the supported 'RttMonRttType'
    // types. The type is RttMonRttType.
    RttMonApplSupportedRttTypes interface{}

    // This object defines the supported 'RttMonRttType' types. The type is bool.
    RttMonApplSupportedRttTypesValid interface{}
}

func (rttMonApplSupportedRttTypesEntry *CISCORTTMONMIB_RttMonApplSupportedRttTypesTable_RttMonApplSupportedRttTypesEntry) GetEntityData() *types.CommonEntityData {
    rttMonApplSupportedRttTypesEntry.EntityData.YFilter = rttMonApplSupportedRttTypesEntry.YFilter
    rttMonApplSupportedRttTypesEntry.EntityData.YangName = "rttMonApplSupportedRttTypesEntry"
    rttMonApplSupportedRttTypesEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplSupportedRttTypesEntry.EntityData.ParentYangName = "rttMonApplSupportedRttTypesTable"
    rttMonApplSupportedRttTypesEntry.EntityData.SegmentPath = "rttMonApplSupportedRttTypesEntry" + types.AddKeyToken(rttMonApplSupportedRttTypesEntry.RttMonApplSupportedRttTypes, "rttMonApplSupportedRttTypes")
    rttMonApplSupportedRttTypesEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplSupportedRttTypesEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplSupportedRttTypesEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplSupportedRttTypesEntry.EntityData.Children = types.NewOrderedMap()
    rttMonApplSupportedRttTypesEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonApplSupportedRttTypesEntry.EntityData.Leafs.Append("rttMonApplSupportedRttTypes", types.YLeaf{"RttMonApplSupportedRttTypes", rttMonApplSupportedRttTypesEntry.RttMonApplSupportedRttTypes})
    rttMonApplSupportedRttTypesEntry.EntityData.Leafs.Append("rttMonApplSupportedRttTypesValid", types.YLeaf{"RttMonApplSupportedRttTypesValid", rttMonApplSupportedRttTypesEntry.RttMonApplSupportedRttTypesValid})

    rttMonApplSupportedRttTypesEntry.EntityData.YListKeys = []string {"RttMonApplSupportedRttTypes"}

    return &(rttMonApplSupportedRttTypesEntry.EntityData)
}

// CISCORTTMONMIB_RttMonApplSupportedProtocolsTable
// A table of which contains the supported Rtt
// Monitor Protocols.
// 
// See the RttMonProtocol textual convention 
// for the definition of each protocol.
type CISCORTTMONMIB_RttMonApplSupportedProtocolsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list that presents the valid Rtt Monitor Protocols. The type is slice of
    // CISCORTTMONMIB_RttMonApplSupportedProtocolsTable_RttMonApplSupportedProtocolsEntry.
    RttMonApplSupportedProtocolsEntry []*CISCORTTMONMIB_RttMonApplSupportedProtocolsTable_RttMonApplSupportedProtocolsEntry
}

func (rttMonApplSupportedProtocolsTable *CISCORTTMONMIB_RttMonApplSupportedProtocolsTable) GetEntityData() *types.CommonEntityData {
    rttMonApplSupportedProtocolsTable.EntityData.YFilter = rttMonApplSupportedProtocolsTable.YFilter
    rttMonApplSupportedProtocolsTable.EntityData.YangName = "rttMonApplSupportedProtocolsTable"
    rttMonApplSupportedProtocolsTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplSupportedProtocolsTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonApplSupportedProtocolsTable.EntityData.SegmentPath = "rttMonApplSupportedProtocolsTable"
    rttMonApplSupportedProtocolsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplSupportedProtocolsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplSupportedProtocolsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplSupportedProtocolsTable.EntityData.Children = types.NewOrderedMap()
    rttMonApplSupportedProtocolsTable.EntityData.Children.Append("rttMonApplSupportedProtocolsEntry", types.YChild{"RttMonApplSupportedProtocolsEntry", nil})
    for i := range rttMonApplSupportedProtocolsTable.RttMonApplSupportedProtocolsEntry {
        rttMonApplSupportedProtocolsTable.EntityData.Children.Append(types.GetSegmentPath(rttMonApplSupportedProtocolsTable.RttMonApplSupportedProtocolsEntry[i]), types.YChild{"RttMonApplSupportedProtocolsEntry", rttMonApplSupportedProtocolsTable.RttMonApplSupportedProtocolsEntry[i]})
    }
    rttMonApplSupportedProtocolsTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonApplSupportedProtocolsTable.EntityData.YListKeys = []string {}

    return &(rttMonApplSupportedProtocolsTable.EntityData)
}

// CISCORTTMONMIB_RttMonApplSupportedProtocolsTable_RttMonApplSupportedProtocolsEntry
// A list that presents the valid Rtt Monitor
// Protocols.
type CISCORTTMONMIB_RttMonApplSupportedProtocolsTable_RttMonApplSupportedProtocolsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object indexes the supported 'RttMonProtocol'
    // protocols. The type is RttMonProtocol.
    RttMonApplSupportedProtocols interface{}

    // This object defines the supported 'RttMonProtocol' protocols. The type is
    // bool.
    RttMonApplSupportedProtocolsValid interface{}
}

func (rttMonApplSupportedProtocolsEntry *CISCORTTMONMIB_RttMonApplSupportedProtocolsTable_RttMonApplSupportedProtocolsEntry) GetEntityData() *types.CommonEntityData {
    rttMonApplSupportedProtocolsEntry.EntityData.YFilter = rttMonApplSupportedProtocolsEntry.YFilter
    rttMonApplSupportedProtocolsEntry.EntityData.YangName = "rttMonApplSupportedProtocolsEntry"
    rttMonApplSupportedProtocolsEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplSupportedProtocolsEntry.EntityData.ParentYangName = "rttMonApplSupportedProtocolsTable"
    rttMonApplSupportedProtocolsEntry.EntityData.SegmentPath = "rttMonApplSupportedProtocolsEntry" + types.AddKeyToken(rttMonApplSupportedProtocolsEntry.RttMonApplSupportedProtocols, "rttMonApplSupportedProtocols")
    rttMonApplSupportedProtocolsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplSupportedProtocolsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplSupportedProtocolsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplSupportedProtocolsEntry.EntityData.Children = types.NewOrderedMap()
    rttMonApplSupportedProtocolsEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonApplSupportedProtocolsEntry.EntityData.Leafs.Append("rttMonApplSupportedProtocols", types.YLeaf{"RttMonApplSupportedProtocols", rttMonApplSupportedProtocolsEntry.RttMonApplSupportedProtocols})
    rttMonApplSupportedProtocolsEntry.EntityData.Leafs.Append("rttMonApplSupportedProtocolsValid", types.YLeaf{"RttMonApplSupportedProtocolsValid", rttMonApplSupportedProtocolsEntry.RttMonApplSupportedProtocolsValid})

    rttMonApplSupportedProtocolsEntry.EntityData.YListKeys = []string {"RttMonApplSupportedProtocols"}

    return &(rttMonApplSupportedProtocolsEntry.EntityData)
}

// CISCORTTMONMIB_RttMonApplPreConfigedTable
// A table of which contains the previously
// configured Script Names and File IO targets.
// 
// These Script Names and File IO targets are installed
// via a different mechanism than this application, and
// are specific to each platform.
type CISCORTTMONMIB_RttMonApplPreConfigedTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that describe the previously configured Script Names and
    // File IO targets. The type is slice of
    // CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry.
    RttMonApplPreConfigedEntry []*CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry
}

func (rttMonApplPreConfigedTable *CISCORTTMONMIB_RttMonApplPreConfigedTable) GetEntityData() *types.CommonEntityData {
    rttMonApplPreConfigedTable.EntityData.YFilter = rttMonApplPreConfigedTable.YFilter
    rttMonApplPreConfigedTable.EntityData.YangName = "rttMonApplPreConfigedTable"
    rttMonApplPreConfigedTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplPreConfigedTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonApplPreConfigedTable.EntityData.SegmentPath = "rttMonApplPreConfigedTable"
    rttMonApplPreConfigedTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplPreConfigedTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplPreConfigedTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplPreConfigedTable.EntityData.Children = types.NewOrderedMap()
    rttMonApplPreConfigedTable.EntityData.Children.Append("rttMonApplPreConfigedEntry", types.YChild{"RttMonApplPreConfigedEntry", nil})
    for i := range rttMonApplPreConfigedTable.RttMonApplPreConfigedEntry {
        rttMonApplPreConfigedTable.EntityData.Children.Append(types.GetSegmentPath(rttMonApplPreConfigedTable.RttMonApplPreConfigedEntry[i]), types.YChild{"RttMonApplPreConfigedEntry", rttMonApplPreConfigedTable.RttMonApplPreConfigedEntry[i]})
    }
    rttMonApplPreConfigedTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonApplPreConfigedTable.EntityData.YListKeys = []string {}

    return &(rttMonApplPreConfigedTable.EntityData)
}

// CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry
// A list of objects that describe the previously
// configured Script Names and File IO targets.
type CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is the type of value being stored in the
    // rttMonApplPreConfigedName object. The type is RttMonApplPreConfigedType.
    RttMonApplPreConfigedType interface{}

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
    RttMonApplPreConfigedName interface{}

    // When this row exists, this value will be 'true'. This object exists only to
    // create a valid row in this  table. The type is bool.
    RttMonApplPreConfigedValid interface{}
}

func (rttMonApplPreConfigedEntry *CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry) GetEntityData() *types.CommonEntityData {
    rttMonApplPreConfigedEntry.EntityData.YFilter = rttMonApplPreConfigedEntry.YFilter
    rttMonApplPreConfigedEntry.EntityData.YangName = "rttMonApplPreConfigedEntry"
    rttMonApplPreConfigedEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplPreConfigedEntry.EntityData.ParentYangName = "rttMonApplPreConfigedTable"
    rttMonApplPreConfigedEntry.EntityData.SegmentPath = "rttMonApplPreConfigedEntry" + types.AddKeyToken(rttMonApplPreConfigedEntry.RttMonApplPreConfigedType, "rttMonApplPreConfigedType") + types.AddKeyToken(rttMonApplPreConfigedEntry.RttMonApplPreConfigedName, "rttMonApplPreConfigedName")
    rttMonApplPreConfigedEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplPreConfigedEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplPreConfigedEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplPreConfigedEntry.EntityData.Children = types.NewOrderedMap()
    rttMonApplPreConfigedEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonApplPreConfigedEntry.EntityData.Leafs.Append("rttMonApplPreConfigedType", types.YLeaf{"RttMonApplPreConfigedType", rttMonApplPreConfigedEntry.RttMonApplPreConfigedType})
    rttMonApplPreConfigedEntry.EntityData.Leafs.Append("rttMonApplPreConfigedName", types.YLeaf{"RttMonApplPreConfigedName", rttMonApplPreConfigedEntry.RttMonApplPreConfigedName})
    rttMonApplPreConfigedEntry.EntityData.Leafs.Append("rttMonApplPreConfigedValid", types.YLeaf{"RttMonApplPreConfigedValid", rttMonApplPreConfigedEntry.RttMonApplPreConfigedValid})

    rttMonApplPreConfigedEntry.EntityData.YListKeys = []string {"RttMonApplPreConfigedType", "RttMonApplPreConfigedName"}

    return &(rttMonApplPreConfigedEntry.EntityData)
}

// CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry_RttMonApplPreConfigedType represents rttMonApplPreConfigedName object.
type CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry_RttMonApplPreConfigedType string

const (
    CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry_RttMonApplPreConfigedType_filePath CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry_RttMonApplPreConfigedType = "filePath"

    CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry_RttMonApplPreConfigedType_scriptName CISCORTTMONMIB_RttMonApplPreConfigedTable_RttMonApplPreConfigedEntry_RttMonApplPreConfigedType = "scriptName"
)

// CISCORTTMONMIB_RttMonApplAuthTable
// A table which contains the definitions for key-strings
// that will be used in authenticating RTR Control Protocol.
type CISCORTTMONMIB_RttMonApplAuthTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list that presents the valid parameters for Authenticating RTR Control
    // Protocol. The type is slice of
    // CISCORTTMONMIB_RttMonApplAuthTable_RttMonApplAuthEntry.
    RttMonApplAuthEntry []*CISCORTTMONMIB_RttMonApplAuthTable_RttMonApplAuthEntry
}

func (rttMonApplAuthTable *CISCORTTMONMIB_RttMonApplAuthTable) GetEntityData() *types.CommonEntityData {
    rttMonApplAuthTable.EntityData.YFilter = rttMonApplAuthTable.YFilter
    rttMonApplAuthTable.EntityData.YangName = "rttMonApplAuthTable"
    rttMonApplAuthTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplAuthTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonApplAuthTable.EntityData.SegmentPath = "rttMonApplAuthTable"
    rttMonApplAuthTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplAuthTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplAuthTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplAuthTable.EntityData.Children = types.NewOrderedMap()
    rttMonApplAuthTable.EntityData.Children.Append("rttMonApplAuthEntry", types.YChild{"RttMonApplAuthEntry", nil})
    for i := range rttMonApplAuthTable.RttMonApplAuthEntry {
        rttMonApplAuthTable.EntityData.Children.Append(types.GetSegmentPath(rttMonApplAuthTable.RttMonApplAuthEntry[i]), types.YChild{"RttMonApplAuthEntry", rttMonApplAuthTable.RttMonApplAuthEntry[i]})
    }
    rttMonApplAuthTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonApplAuthTable.EntityData.YListKeys = []string {}

    return &(rttMonApplAuthTable.EntityData)
}

// CISCORTTMONMIB_RttMonApplAuthTable_RttMonApplAuthEntry
// A list that presents the valid parameters for Authenticating
// RTR Control Protocol.
type CISCORTTMONMIB_RttMonApplAuthTable_RttMonApplAuthEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonApplAuthTable. This is a pseudo-random number selected by the
    // management station when creating a row via the rttMonApplAuthStatus 
    // object. If the pseudo-random number is already in use, an 
    // 'inconsistentValue' is returned. Currently, only one row  can be created.
    // The type is interface{} with range: 0..2147483647.
    RttMonApplAuthIndex interface{}

    // A string which represents the key-chain name. If multiple key-strings are
    // specified, then the authenticator will  alternate between the specified
    // strings. The type is string with length: 1..48.
    RttMonApplAuthKeyChain interface{}

    // A string which represents a key-string name whose id is 1. The type is
    // string with length: 1..48.
    RttMonApplAuthKeyString1 interface{}

    // A string which represents a key-string name whose id is 2. The type is
    // string with length: 1..48.
    RttMonApplAuthKeyString2 interface{}

    // A string which represents a key-string name whose id is 3. The type is
    // string with length: 1..48.
    RttMonApplAuthKeyString3 interface{}

    // A string which represents a key-string name whose id is 4. The type is
    // string with length: 1..48.
    RttMonApplAuthKeyString4 interface{}

    // A string which represents a key-string name whose id is 5. The type is
    // string with length: 1..48.
    RttMonApplAuthKeyString5 interface{}

    // The status of the Authentication row. The type is RowStatus.
    RttMonApplAuthStatus interface{}
}

func (rttMonApplAuthEntry *CISCORTTMONMIB_RttMonApplAuthTable_RttMonApplAuthEntry) GetEntityData() *types.CommonEntityData {
    rttMonApplAuthEntry.EntityData.YFilter = rttMonApplAuthEntry.YFilter
    rttMonApplAuthEntry.EntityData.YangName = "rttMonApplAuthEntry"
    rttMonApplAuthEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonApplAuthEntry.EntityData.ParentYangName = "rttMonApplAuthTable"
    rttMonApplAuthEntry.EntityData.SegmentPath = "rttMonApplAuthEntry" + types.AddKeyToken(rttMonApplAuthEntry.RttMonApplAuthIndex, "rttMonApplAuthIndex")
    rttMonApplAuthEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonApplAuthEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonApplAuthEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonApplAuthEntry.EntityData.Children = types.NewOrderedMap()
    rttMonApplAuthEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthIndex", types.YLeaf{"RttMonApplAuthIndex", rttMonApplAuthEntry.RttMonApplAuthIndex})
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthKeyChain", types.YLeaf{"RttMonApplAuthKeyChain", rttMonApplAuthEntry.RttMonApplAuthKeyChain})
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthKeyString1", types.YLeaf{"RttMonApplAuthKeyString1", rttMonApplAuthEntry.RttMonApplAuthKeyString1})
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthKeyString2", types.YLeaf{"RttMonApplAuthKeyString2", rttMonApplAuthEntry.RttMonApplAuthKeyString2})
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthKeyString3", types.YLeaf{"RttMonApplAuthKeyString3", rttMonApplAuthEntry.RttMonApplAuthKeyString3})
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthKeyString4", types.YLeaf{"RttMonApplAuthKeyString4", rttMonApplAuthEntry.RttMonApplAuthKeyString4})
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthKeyString5", types.YLeaf{"RttMonApplAuthKeyString5", rttMonApplAuthEntry.RttMonApplAuthKeyString5})
    rttMonApplAuthEntry.EntityData.Leafs.Append("rttMonApplAuthStatus", types.YLeaf{"RttMonApplAuthStatus", rttMonApplAuthEntry.RttMonApplAuthStatus})

    rttMonApplAuthEntry.EntityData.YListKeys = []string {"RttMonApplAuthIndex"}

    return &(rttMonApplAuthEntry.EntityData)
}

// CISCORTTMONMIB_RttMonCtrlAdminTable
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
type CISCORTTMONMIB_RttMonCtrlAdminTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual RTT control row. The type
    // is slice of CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry.
    RttMonCtrlAdminEntry []*CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry
}

func (rttMonCtrlAdminTable *CISCORTTMONMIB_RttMonCtrlAdminTable) GetEntityData() *types.CommonEntityData {
    rttMonCtrlAdminTable.EntityData.YFilter = rttMonCtrlAdminTable.YFilter
    rttMonCtrlAdminTable.EntityData.YangName = "rttMonCtrlAdminTable"
    rttMonCtrlAdminTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonCtrlAdminTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonCtrlAdminTable.EntityData.SegmentPath = "rttMonCtrlAdminTable"
    rttMonCtrlAdminTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonCtrlAdminTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonCtrlAdminTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonCtrlAdminTable.EntityData.Children = types.NewOrderedMap()
    rttMonCtrlAdminTable.EntityData.Children.Append("rttMonCtrlAdminEntry", types.YChild{"RttMonCtrlAdminEntry", nil})
    for i := range rttMonCtrlAdminTable.RttMonCtrlAdminEntry {
        rttMonCtrlAdminTable.EntityData.Children.Append(types.GetSegmentPath(rttMonCtrlAdminTable.RttMonCtrlAdminEntry[i]), types.YChild{"RttMonCtrlAdminEntry", rttMonCtrlAdminTable.RttMonCtrlAdminEntry[i]})
    }
    rttMonCtrlAdminTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonCtrlAdminTable.EntityData.YListKeys = []string {}

    return &(rttMonCtrlAdminTable.EntityData)
}

// CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry
// A base list of objects that define a conceptual RTT
// control row.
type CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonCtrlAdminTable. This is a pseudo-random number, selected by the
    // management  station or auto-generated based on  operation started by the 
    // management station,when creating a row via  the rttMonCtrlAdminStatus
    // object.  If the pseudo-random   number is already in use an
    // 'inconsistentValue' return code   will be returned when set operation is
    // attempted. The type is interface{} with range: 1..2147483647.
    RttMonCtrlAdminIndex interface{}

    // Identifies the entity that created this table row. The type is string with
    // length: 0..255.
    RttMonCtrlAdminOwner interface{}

    // A string which is used by a managing application to identify the RTT
    // target.  This string is inserted into trap notifications, but has no other
    // significance to the  agent. The type is string with length: 0..16.
    RttMonCtrlAdminTag interface{}

    // The type of RTT operation to be performed.  This value must be set in the
    // same PDU or before setting any type specific configuration.  Note: The RTT
    // operation 'lspGroup' cannot be created via this control row. It will be
    // created automatically by Auto SAA L3 MPLS VPN when rttMplsVpnMonCtrlLpd is
    // 'true'. The type is RttMonRttType.
    RttMonCtrlAdminRttType interface{}

    // This object defines an administrative threshold limit. If the RTT operation
    // time exceeds this limit and if the  conditions specified in
    // rttMonReactAdminThresholdType or  rttMonHistoryAdminFilter are satisfied, a
    // threshold is generated. The type is interface{} with range: 0..2147483647.
    // Units are milliseconds.
    RttMonCtrlAdminThreshold interface{}

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
    RttMonCtrlAdminFrequency interface{}

    // Specifies the duration to wait for a RTT operation completion.  The value
    // of this object cannot be set to  a value which would specify a duration
    // exceeding  rttMonCtrlAdminFrequency.  For connection oriented protocols,
    // this may cause the connection to be closed by the probe.  Once closed, it
    // will be assumed that the connection reestablishment will be performed.  To
    // prevent unwanted closure of connections, be sure to set this value to a
    // realistic connection timeout. The type is interface{} with range:
    // 0..604800000. Units are milliseconds.
    RttMonCtrlAdminTimeout interface{}

    // When set to true, the resulting data in each RTT operation is compared with
    // the expected data.  This includes checking header information (if possible)
    // and exact packet size.  Any mismatch will be recorded in the
    // rttMonStatsCollectVerifyErrors object.  Some RttMonRttTypes may not support
    // this option.  When a type does not support this option, the agent will 
    // transition this object to false.  It is the management applications
    // responsibility to check for this  transition. The type is bool.
    RttMonCtrlAdminVerifyData interface{}

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
    RttMonCtrlAdminStatus interface{}

    // When set to true, this entry will be shown in 'show running' command and
    // can be saved into Non-volatile memory. The type is bool.
    RttMonCtrlAdminNvgen interface{}

    // If the operation is created through auto measure group creation, then this
    // string will specify the group name to which this operation is associated.
    // The type is string with length: 0..64.
    RttMonCtrlAdminGroupName interface{}

    // This object value will be placed into the rttMonCtrlOperRttLife object when
    // the rttMonCtrlOperState object transitions to 'active' or 'pending'.  The
    // value 2147483647 has a special meaning.  When this object is set to
    // 2147483647, the  rttMonCtrlOperRttLife object will not decrement.   And
    // thus the life time will never end. The type is interface{} with range:
    // 0..2147483647. Units are seconds.
    RttMonScheduleAdminRttLife interface{}

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
    RttMonScheduleAdminRttStartTime interface{}

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
    RttMonScheduleAdminConceptRowAgeout interface{}

    // When set to true, this entry will be scheduled to run automatically for the
    // specified duration equal to the life configured, at the same time daily. 
    // This value cannot be set to true  (a) if rttMonScheduleAdminRttLife object
    // has value greater or    equal to 86400 seconds. (b) if sum of values of
    // rttMonScheduleAdminRttLife and    rttMonScheduleAdminConceptRowAgeout is
    // less or equal to    86400 seconds. The type is bool.
    RttMonScheduleAdminRttRecurring interface{}

    // The amount of time this conceptual Rtt control row will exist when not in
    // an 'active' rttMonCtrlOperState.  When this conceptual Rtt control row
    // enters an 'active' state, this timer will be reset and suspended.  When
    // this conceptual RTT control row enters a state other than 'active', the
    // timer will be restarted.  NOTE:  It is the same as
    // rttMonScheduleAdminConceptRowAgeout        except DEFVAL is 0 to be
    // consistent with CLI ageout        default.  When this value is set to zero,
    // this entry will never be aged out. The type is interface{} with range:
    // 0..2073600. Units are seconds.
    RttMonScheduleAdminConceptRowAgeoutV2 interface{}

    // If true, a reaction is generated when a RTT operation to a
    // rttMonEchoAdminTargetAddress (echo type) causes 
    // rttMonCtrlOperConnectionLostOccurred to change its  value.  Thus
    // connections to intermediate hops will  not cause this value to change.
    // rttMonReactAdminConnectionEnable object is superseded by rttMonReactVar.
    // The type is bool.
    RttMonReactAdminConnectionEnable interface{}

    // If true, a reaction is generated when a RTT operation causes
    // rttMonCtrlOperTimeoutOccurred  to change its value.    When the
    // RttMonRttType is 'pathEcho' timeouts to  intermediate hops will not cause 
    // rttMonCtrlOperTimeoutOccurred to change its value.
    // rttMonReactAdminTimeoutEnable object is superseded by rttMonReactVar. The
    // type is bool.
    RttMonReactAdminTimeoutEnable interface{}

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
    // RttMonReactAdminThresholdType.
    RttMonReactAdminThresholdType interface{}

    // This object defines a threshold limit. If the RTT operation time falls
    // below this limit and if the conditions specified in
    // rttMonReactAdminThresholdType are satisfied, an  threshold is generated.
    // rttMonReactAdminThresholdFalling object is superseded by
    // rttMonReactThresholdFalling. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    RttMonReactAdminThresholdFalling interface{}

    // This object defines the 'x' value of the xOfy condition specified in
    // rttMonReactAdminThresholdType. rttMonReactAdminThresholdCount object is
    // superseded by rttMonReactThresholdCountX. The type is interface{} with
    // range: 1..16.
    RttMonReactAdminThresholdCount interface{}

    // This object defines the 'y' value of the xOfy condition specified in
    // rttMonReactAdminThresholdType. rttMonReactAdminThresholdCount2 object is
    // superseded by rttMonReactThresholdCountyY. The type is interface{} with
    // range: 1..16.
    RttMonReactAdminThresholdCount2 interface{}

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
    // RttMonReactAdminActionType.
    RttMonReactAdminActionType interface{}

    // If true, a reaction is generated when a RTT operation causes
    // rttMonCtrlOperVerifyErrorOccurred  to change its value.
    // rttMonReactAdminVerifyErrorEnable object is superseded by rttMonReactVar.
    // The type is bool.
    RttMonReactAdminVerifyErrorEnable interface{}

    // The maximum number of groups of paths to record. Specifically this is the
    // number of hourly groups  to keep before rolling over.    The value of one
    // is not advisable because the  group will close and immediately be deleted
    // before the network management station will have the  opportunity to
    // retrieve the statistics.   The value used in the rttMonStatsCaptureTable to
    // uniquely identify this group is the  rttMonStatsCaptureStartTimeIndex. 
    // HTTP and Jitter probes store only two hours of data.  When this object is
    // set to the value of zero all  rttMonStatsCaptureTable data capturing will
    // be shut off. The type is interface{} with range: 0..25.
    RttMonStatisticsAdminNumHourGroups interface{}

    // When RttMonRttType is 'pathEcho' this is the maximum number of statistics
    // paths to record per hourly group.   This value directly represents the path
    // to a target.   For all other RttMonRttTypes this value will be  forced to
    // one by the agent.  NOTE: For 'pathEcho' a source to target path will be    
    // created to to hold all errors that occur when a        specific path or
    // connection has not be found/setup.        Thus, it is advised to set this
    // value greater       than one.  Since this index does not rollover, only the
    // first rttMonStatisticsAdminNumPaths will be kept. The type is interface{}
    // with range: 1..128.
    RttMonStatisticsAdminNumPaths interface{}

    // When RttMonRttType is 'pathEcho' this is the maximum number of statistics
    // hops to record per path group.   This value directly represents the number
    // of hops along  a path to a target, thus we can only support 30 hops.   For
    // all other RttMonRttTypes this value will be  forced to one by the agent. 
    // Since this index does not rollover, only the first
    // rttMonStatisticsAdminNumHops will be kept. This object  is applicable to
    // pathEcho probes only. The type is interface{} with range: 1..30.
    RttMonStatisticsAdminNumHops interface{}

    // The maximum number of statistical distribution Buckets to accumulate. 
    // Since this index does not rollover, only the first
    // rttMonStatisticsAdminNumDistBuckets will be kept.  The last
    // rttMonStatisticsAdminNumDistBucket will contain all entries from its
    // distribution interval start point to infinity. This object is not
    // applicable  to http and jitter probes. The type is interface{} with range:
    // 1..20.
    RttMonStatisticsAdminNumDistBuckets interface{}

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
    RttMonStatisticsAdminDistInterval interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the rttMonCtrlOperRttLife object.  A
    // new life is created when the same conceptual RTT control row is restarted
    // via the transition of the  rttMonCtrlOperRttLife object and its subsequent 
    // countdown.  The value of zero will shut off all  rttMonHistoryAdminTable
    // data collection. The type is interface{} with range: 0..2.
    RttMonHistoryAdminNumLives interface{}

    // The maximum number of history buckets to record.  When the RttMonRttType is
    // 'pathEcho'  this value directly  represents a path to a target.  For all
    // other  RttMonRttTypes this value should be set to the number  of operations
    // to keep per lifetime.  After rttMonHistoryAdminNumBuckets are filled, the 
    // and the oldest entries are deleted and the most recent
    // rttMonHistoryAdminNumBuckets buckets are retained. The type is interface{}
    // with range: 1..60.
    RttMonHistoryAdminNumBuckets interface{}

    // The maximum number of history samples to record per bucket.  When the
    // RttMonRttType is 'pathEcho' this  value directly represents the number of
    // hops along a  path to a target, thus we can only support 30 hops. For all
    // other RttMonRttTypes this value will be  forced to one by the agent. The
    // type is interface{} with range: 1..30.
    RttMonHistoryAdminNumSamples interface{}

    // Defines a filter for adding RTT results to the history buffer:  none       
    // - no history is recorded all           - the results of all completion
    // times                   and failed completions are recorded overThreshold -
    // the results of completion times                  over
    // rttMonCtrlAdminThreshold are                   recorded. failures      -
    // the results of failed operations (only)                   are recorded. The
    // type is RttMonHistoryAdminFilter.
    RttMonHistoryAdminFilter interface{}

    // This object is updated whenever an object in the conceptual RTT control row
    // is changed or updated. The type is interface{} with range: 0..4294967295.
    RttMonCtrlOperModificationTime interface{}

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
    RttMonCtrlOperDiagText interface{}

    // This object is set when the rttMonCtrlOperState is set to reset. The type
    // is interface{} with range: 0..4294967295.
    RttMonCtrlOperResetTime interface{}

    // This object is the number of octets currently in use by this composite
    // conceptual RTT row.  A composite conceptual row include the control,
    // statistics, and  history conceptual rows combined.  (All octets that are
    // addressed via the rttMonCtrlAdminIndex in this mib.). The type is
    // interface{} with range: 0..4294967295.
    RttMonCtrlOperOctetsInUse interface{}

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
    RttMonCtrlOperConnectionLostOccurred interface{}

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
    RttMonCtrlOperTimeoutOccurred interface{}

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
    RttMonCtrlOperOverThresholdOccurred interface{}

    // This is the total number of probe operations that have been attempted.    
    // This value is incremented for each start of an RTT  operation.  Thus when
    // rttMonCtrlAdminRttType is set to  'pathEcho' this value will be incremented
    // by one and  not for very every hop along the path.  This object has the
    // special behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object.  This value is not effected by the rollover of a
    // statistics hourly group. The type is interface{} with range: 0..2147483647.
    RttMonCtrlOperNumRtts interface{}

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
    RttMonCtrlOperRttLife interface{}

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
    // multi-bound PDU. The type is RttMonCtrlOperState.
    RttMonCtrlOperState interface{}

    // This object is true if rttMonCtrlAdminVerifyData is set to true and data
    // corruption occurs. The type is bool.
    RttMonCtrlOperVerifyErrorOccurred interface{}

    // The completion time of the latest RTT operation successfully completed. 
    // The unit of this object will be microsecond when rttMonCtrlAdminRttType is
    // set to 'jitter' and  rttMonEchoAdminPrecision is set to 'microsecond'.
    // Otherwise, the unit of this object will be millisecond. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds/microseconds.
    RttMonLatestRttOperCompletionTime interface{}

    // A sense code for the completion status of the latest RTT operation. The
    // type is RttResponseSense.
    RttMonLatestRttOperSense interface{}

    // An application specific sense code for the completion status of the latest
    // RTT operation.  This  object will only be valid when the 
    // rttMonLatestRttOperSense object is set to  'applicationSpecific'. 
    // Otherwise, this object's  value is not valid. The type is interface{} with
    // range: 0..2147483647.
    RttMonLatestRttOperApplSpecificSense interface{}

    // A sense description for the completion status of the latest RTT operation
    // when the  rttMonLatestRttOperSense object is set to  'applicationSpecific'.
    // The type is string.
    RttMonLatestRttOperSenseDescription interface{}

    // The value of the agent system time at the time of the latest RTT operation.
    // The type is interface{} with range: 0..4294967295.
    RttMonLatestRttOperTime interface{}

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
    RttMonLatestRttOperAddress interface{}
}

func (rttMonCtrlAdminEntry *CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry) GetEntityData() *types.CommonEntityData {
    rttMonCtrlAdminEntry.EntityData.YFilter = rttMonCtrlAdminEntry.YFilter
    rttMonCtrlAdminEntry.EntityData.YangName = "rttMonCtrlAdminEntry"
    rttMonCtrlAdminEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonCtrlAdminEntry.EntityData.ParentYangName = "rttMonCtrlAdminTable"
    rttMonCtrlAdminEntry.EntityData.SegmentPath = "rttMonCtrlAdminEntry" + types.AddKeyToken(rttMonCtrlAdminEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex")
    rttMonCtrlAdminEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonCtrlAdminEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonCtrlAdminEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonCtrlAdminEntry.EntityData.Children = types.NewOrderedMap()
    rttMonCtrlAdminEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonCtrlAdminEntry.RttMonCtrlAdminIndex})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminOwner", types.YLeaf{"RttMonCtrlAdminOwner", rttMonCtrlAdminEntry.RttMonCtrlAdminOwner})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminTag", types.YLeaf{"RttMonCtrlAdminTag", rttMonCtrlAdminEntry.RttMonCtrlAdminTag})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminRttType", types.YLeaf{"RttMonCtrlAdminRttType", rttMonCtrlAdminEntry.RttMonCtrlAdminRttType})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminThreshold", types.YLeaf{"RttMonCtrlAdminThreshold", rttMonCtrlAdminEntry.RttMonCtrlAdminThreshold})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminFrequency", types.YLeaf{"RttMonCtrlAdminFrequency", rttMonCtrlAdminEntry.RttMonCtrlAdminFrequency})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminTimeout", types.YLeaf{"RttMonCtrlAdminTimeout", rttMonCtrlAdminEntry.RttMonCtrlAdminTimeout})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminVerifyData", types.YLeaf{"RttMonCtrlAdminVerifyData", rttMonCtrlAdminEntry.RttMonCtrlAdminVerifyData})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminStatus", types.YLeaf{"RttMonCtrlAdminStatus", rttMonCtrlAdminEntry.RttMonCtrlAdminStatus})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminNvgen", types.YLeaf{"RttMonCtrlAdminNvgen", rttMonCtrlAdminEntry.RttMonCtrlAdminNvgen})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminGroupName", types.YLeaf{"RttMonCtrlAdminGroupName", rttMonCtrlAdminEntry.RttMonCtrlAdminGroupName})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonScheduleAdminRttLife", types.YLeaf{"RttMonScheduleAdminRttLife", rttMonCtrlAdminEntry.RttMonScheduleAdminRttLife})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonScheduleAdminRttStartTime", types.YLeaf{"RttMonScheduleAdminRttStartTime", rttMonCtrlAdminEntry.RttMonScheduleAdminRttStartTime})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonScheduleAdminConceptRowAgeout", types.YLeaf{"RttMonScheduleAdminConceptRowAgeout", rttMonCtrlAdminEntry.RttMonScheduleAdminConceptRowAgeout})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonScheduleAdminRttRecurring", types.YLeaf{"RttMonScheduleAdminRttRecurring", rttMonCtrlAdminEntry.RttMonScheduleAdminRttRecurring})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonScheduleAdminConceptRowAgeoutV2", types.YLeaf{"RttMonScheduleAdminConceptRowAgeoutV2", rttMonCtrlAdminEntry.RttMonScheduleAdminConceptRowAgeoutV2})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminConnectionEnable", types.YLeaf{"RttMonReactAdminConnectionEnable", rttMonCtrlAdminEntry.RttMonReactAdminConnectionEnable})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminTimeoutEnable", types.YLeaf{"RttMonReactAdminTimeoutEnable", rttMonCtrlAdminEntry.RttMonReactAdminTimeoutEnable})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminThresholdType", types.YLeaf{"RttMonReactAdminThresholdType", rttMonCtrlAdminEntry.RttMonReactAdminThresholdType})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminThresholdFalling", types.YLeaf{"RttMonReactAdminThresholdFalling", rttMonCtrlAdminEntry.RttMonReactAdminThresholdFalling})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminThresholdCount", types.YLeaf{"RttMonReactAdminThresholdCount", rttMonCtrlAdminEntry.RttMonReactAdminThresholdCount})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminThresholdCount2", types.YLeaf{"RttMonReactAdminThresholdCount2", rttMonCtrlAdminEntry.RttMonReactAdminThresholdCount2})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminActionType", types.YLeaf{"RttMonReactAdminActionType", rttMonCtrlAdminEntry.RttMonReactAdminActionType})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonReactAdminVerifyErrorEnable", types.YLeaf{"RttMonReactAdminVerifyErrorEnable", rttMonCtrlAdminEntry.RttMonReactAdminVerifyErrorEnable})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonStatisticsAdminNumHourGroups", types.YLeaf{"RttMonStatisticsAdminNumHourGroups", rttMonCtrlAdminEntry.RttMonStatisticsAdminNumHourGroups})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonStatisticsAdminNumPaths", types.YLeaf{"RttMonStatisticsAdminNumPaths", rttMonCtrlAdminEntry.RttMonStatisticsAdminNumPaths})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonStatisticsAdminNumHops", types.YLeaf{"RttMonStatisticsAdminNumHops", rttMonCtrlAdminEntry.RttMonStatisticsAdminNumHops})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonStatisticsAdminNumDistBuckets", types.YLeaf{"RttMonStatisticsAdminNumDistBuckets", rttMonCtrlAdminEntry.RttMonStatisticsAdminNumDistBuckets})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonStatisticsAdminDistInterval", types.YLeaf{"RttMonStatisticsAdminDistInterval", rttMonCtrlAdminEntry.RttMonStatisticsAdminDistInterval})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonHistoryAdminNumLives", types.YLeaf{"RttMonHistoryAdminNumLives", rttMonCtrlAdminEntry.RttMonHistoryAdminNumLives})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonHistoryAdminNumBuckets", types.YLeaf{"RttMonHistoryAdminNumBuckets", rttMonCtrlAdminEntry.RttMonHistoryAdminNumBuckets})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonHistoryAdminNumSamples", types.YLeaf{"RttMonHistoryAdminNumSamples", rttMonCtrlAdminEntry.RttMonHistoryAdminNumSamples})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonHistoryAdminFilter", types.YLeaf{"RttMonHistoryAdminFilter", rttMonCtrlAdminEntry.RttMonHistoryAdminFilter})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperModificationTime", types.YLeaf{"RttMonCtrlOperModificationTime", rttMonCtrlAdminEntry.RttMonCtrlOperModificationTime})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperDiagText", types.YLeaf{"RttMonCtrlOperDiagText", rttMonCtrlAdminEntry.RttMonCtrlOperDiagText})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperResetTime", types.YLeaf{"RttMonCtrlOperResetTime", rttMonCtrlAdminEntry.RttMonCtrlOperResetTime})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperOctetsInUse", types.YLeaf{"RttMonCtrlOperOctetsInUse", rttMonCtrlAdminEntry.RttMonCtrlOperOctetsInUse})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperConnectionLostOccurred", types.YLeaf{"RttMonCtrlOperConnectionLostOccurred", rttMonCtrlAdminEntry.RttMonCtrlOperConnectionLostOccurred})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperTimeoutOccurred", types.YLeaf{"RttMonCtrlOperTimeoutOccurred", rttMonCtrlAdminEntry.RttMonCtrlOperTimeoutOccurred})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperOverThresholdOccurred", types.YLeaf{"RttMonCtrlOperOverThresholdOccurred", rttMonCtrlAdminEntry.RttMonCtrlOperOverThresholdOccurred})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperNumRtts", types.YLeaf{"RttMonCtrlOperNumRtts", rttMonCtrlAdminEntry.RttMonCtrlOperNumRtts})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperRttLife", types.YLeaf{"RttMonCtrlOperRttLife", rttMonCtrlAdminEntry.RttMonCtrlOperRttLife})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperState", types.YLeaf{"RttMonCtrlOperState", rttMonCtrlAdminEntry.RttMonCtrlOperState})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonCtrlOperVerifyErrorOccurred", types.YLeaf{"RttMonCtrlOperVerifyErrorOccurred", rttMonCtrlAdminEntry.RttMonCtrlOperVerifyErrorOccurred})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonLatestRttOperCompletionTime", types.YLeaf{"RttMonLatestRttOperCompletionTime", rttMonCtrlAdminEntry.RttMonLatestRttOperCompletionTime})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonLatestRttOperSense", types.YLeaf{"RttMonLatestRttOperSense", rttMonCtrlAdminEntry.RttMonLatestRttOperSense})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonLatestRttOperApplSpecificSense", types.YLeaf{"RttMonLatestRttOperApplSpecificSense", rttMonCtrlAdminEntry.RttMonLatestRttOperApplSpecificSense})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonLatestRttOperSenseDescription", types.YLeaf{"RttMonLatestRttOperSenseDescription", rttMonCtrlAdminEntry.RttMonLatestRttOperSenseDescription})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonLatestRttOperTime", types.YLeaf{"RttMonLatestRttOperTime", rttMonCtrlAdminEntry.RttMonLatestRttOperTime})
    rttMonCtrlAdminEntry.EntityData.Leafs.Append("rttMonLatestRttOperAddress", types.YLeaf{"RttMonLatestRttOperAddress", rttMonCtrlAdminEntry.RttMonLatestRttOperAddress})

    rttMonCtrlAdminEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex"}

    return &(rttMonCtrlAdminEntry.EntityData)
}

// CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState represents        part of a multi-bound PDU.
type CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState string

const (
    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState_reset CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState = "reset"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState_orderlyStop CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState = "orderlyStop"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState_immediateStop CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState = "immediateStop"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState_pending CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState = "pending"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState_inactive CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState = "inactive"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState_active CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState = "active"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState_restart CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlOperState = "restart"
)

// CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter represents                  are recorded.
type CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter string

const (
    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter_none CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter = "none"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter_all CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter = "all"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter_overThreshold CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter = "overThreshold"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter_failures CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonHistoryAdminFilter = "failures"
)

// CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType represents rttMonReactActionType.
type CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType string

const (
    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_none CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "none"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_trapOnly CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "trapOnly"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_nmvtOnly CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "nmvtOnly"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_triggerOnly CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "triggerOnly"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_trapAndNmvt CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "trapAndNmvt"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_trapAndTrigger CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "trapAndTrigger"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_nmvtAndTrigger CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "nmvtAndTrigger"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType_trapNmvtAndTrigger CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminActionType = "trapNmvtAndTrigger"
)

// CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType represents rttMonReactThresholdType.
type CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType string

const (
    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType_never CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType = "never"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType_immediate CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType = "immediate"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType_consecutive CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType = "consecutive"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType_xOfy CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType = "xOfy"

    CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType_average CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonReactAdminThresholdType = "average"
)

// CISCORTTMONMIB_RttMonEchoAdminTable
// A table that contains Round Trip Time (RTT) specific
// definitions.
// 
// This table is controlled via the 
// rttMonCtrlAdminTable.  Entries in this table are
// created via the rttMonCtrlAdminStatus object.
type CISCORTTMONMIB_RttMonEchoAdminTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for RttMonRttType
    // conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry.
    RttMonEchoAdminEntry []*CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry
}

func (rttMonEchoAdminTable *CISCORTTMONMIB_RttMonEchoAdminTable) GetEntityData() *types.CommonEntityData {
    rttMonEchoAdminTable.EntityData.YFilter = rttMonEchoAdminTable.YFilter
    rttMonEchoAdminTable.EntityData.YangName = "rttMonEchoAdminTable"
    rttMonEchoAdminTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonEchoAdminTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonEchoAdminTable.EntityData.SegmentPath = "rttMonEchoAdminTable"
    rttMonEchoAdminTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonEchoAdminTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonEchoAdminTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonEchoAdminTable.EntityData.Children = types.NewOrderedMap()
    rttMonEchoAdminTable.EntityData.Children.Append("rttMonEchoAdminEntry", types.YChild{"RttMonEchoAdminEntry", nil})
    for i := range rttMonEchoAdminTable.RttMonEchoAdminEntry {
        rttMonEchoAdminTable.EntityData.Children.Append(types.GetSegmentPath(rttMonEchoAdminTable.RttMonEchoAdminEntry[i]), types.YChild{"RttMonEchoAdminEntry", rttMonEchoAdminTable.RttMonEchoAdminEntry[i]})
    }
    rttMonEchoAdminTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonEchoAdminTable.EntityData.YListKeys = []string {}

    return &(rttMonEchoAdminTable.EntityData)
}

// CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry
// A list of objects that define specific configuration for
// RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // Specifies the protocol to be used to perform the RTT operation. The
    // following list defines what protocol  should be used for each probe type: 
    // echo, pathEcho   - ipIcmpEcho / mplsLspPingAppl udpEcho          -
    // ipUdpEchoAppl tcpConnect       - ipTcpConn http             - httpAppl
    // jitter           - jitterAppl dlsw             - dlswAppl dhcp            
    // - dhcpAppl ftp              - ftpAppl mplsLspPing      - mplsLspPingAppl
    // voip             - voipAppl video            - videoAppl  When this
    // protocol does not support the type, a 'badValue' error will be returned.
    // The type is RttMonProtocol.
    RttMonEchoAdminProtocol interface{}

    // A string which specifies the address of the target. The type is string.
    RttMonEchoAdminTargetAddress interface{}

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
    RttMonEchoAdminPktDataRequestSize interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the response message. This value is passed to the RTT Echo
    // Server via a field in the ARR Header.  For non-ARR RTT request/response
    // (i.e. ipIcmpecho) this value will be set by the agent to match the size of
    // rttMonEchoAdminPktDataRequestSize, when native payloads are supported. 
    // REMEMBER:  The ARR Header overhead is not included             in this
    // value.  This object is only supported by SNA protocols. The type is
    // interface{} with range: 0..16384.
    RttMonEchoAdminPktDataResponseSize interface{}

    // This object represents the target's port number. This object is applicable
    // to udpEcho, tcpConnect and jitter probes. The type is interface{} with
    // range: 0..65536.
    RttMonEchoAdminTargetPort interface{}

    // A string which specifies the IP address of the source. This object is
    // applicable to all probes except dns, dlsw  and sna. The type is string.
    RttMonEchoAdminSourceAddress interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a  port allocated by the system. This
    // object is applicable  to all probes except dns, dlsw and sna. The type is
    // interface{} with range: 0..65536.
    RttMonEchoAdminSourcePort interface{}

    // If this object is enabled, then the RTR application will send control
    // messages to a responder, residing on the  target router to respond to the
    // data request packets being  sent by the source router. This object is not
    // applicable to  echo, pathEcho, dns and http probes. The type is bool.
    RttMonEchoAdminControlEnable interface{}

    // This object represents the type of service octet in an IP header. This
    // object is not applicable to dhcp, dns,  ethernetPing and ethernetJitter.
    // The type is interface{} with range: 0..255.
    RttMonEchoAdminTOS interface{}

    // If this object is enabled then it means that the application calculates
    // response time for a specific path, defined in rttMonEchoPathAdminEntry.
    // This object is applicable to echo  probe only. The type is bool.
    RttMonEchoAdminLSREnable interface{}

    // A string which specifies the address of the target. This string can be in
    // IP address format or a hostname. This object is applicable to dns probe
    // only. The type is string.
    RttMonEchoAdminTargetAddressString interface{}

    // A string which specifies the ip address of the name-server. This object is
    // applicable to dns probe only. The type is string.
    RttMonEchoAdminNameServer interface{}

    // A code that represents the specific type of RTT operation. This object is
    // applicable to http and ftp probe only. The type is RttMonOperation.
    RttMonEchoAdminOperation interface{}

    // A string which specifies the version number of the HTTP Server.  The syntax
    // for the version string is  <major number>.<minor number> An example would
    // be 1.0,  1.1 etc.,.  This object is applicable to http probe only. The type
    // is string with length: 3..10.
    RttMonEchoAdminHTTPVersion interface{}

    // A string which represents the URL to which a HTTP probe should communicate
    // with. This object is applicable to http probe only. The type is string.
    RttMonEchoAdminURL interface{}

    // If this object is false then it means that HTTP request should not download
    // cached pages. This means that the request should  be forwarded to the
    // origin server. This object is applicable to http probe only. The type is
    // bool.
    RttMonEchoAdminCache interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. This value is currently used for  Jitter probe. This object
    // is applicable to jitter probe only. The type is interface{} with range:
    // 0..60000. Units are milliseconds.
    RttMonEchoAdminInterval interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is currently used for Jitter probe.  This object is applicable
    // to jitter probe only. The type is interface{} with range:
    // -2147483648..2147483647.
    RttMonEchoAdminNumPackets interface{}

    // This string represents the proxy server information. This object is
    // applicable to http probe only. The type is string.
    RttMonEchoAdminProxy interface{}

    // This string stores the content of HTTP raw request. If the request cannot
    // fit into String1 then it should  be split and put in Strings 1 through 5. 
    // This string stores the content of the DHCP raw option data.  The raw DHCP
    // option data must be in HEX. If an odd number of characters are specified, a
    // 0 will be appended to the end of the string.  Only DHCP option 82 (decimal)
    // is allowed. Here is an example of a valid string: 5208010610005A6F1234 Only
    // rttMonEchoAdminString1 is used for dhcp, Strings 1 through 5 are not used. 
    // This object is applicable to http and dhcp probe  types only. The type is
    // string.
    RttMonEchoAdminString1 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    RttMonEchoAdminString2 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    RttMonEchoAdminString3 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    RttMonEchoAdminString4 interface{}

    // This string stores the content of HTTP raw request.
    // rttMonEchoAdminString1-5 are concatenated to  form the HTTP raw request
    // used in the RTT operation. This object is applicable to http probe only.
    // The type is string.
    RttMonEchoAdminString5 interface{}

    // A code that represents the specific type of RTT operation. This object is
    // applicable to ftp probe only. The type is RttMonOperation.
    RttMonEchoAdminMode interface{}

    // This field is used to specify the VPN name in which the RTT operation will
    // be used. For regular RTT operation this field should not be configured. The
    // agent  will use this field to identify the VPN routing Table for this
    // operation. The type is string with length: 0..32.
    RttMonEchoAdminVrfName interface{}

    // Specifies the codec type to be used with jitter probe. This is applicable
    // only for the jitter probe.  If codec-type is configured the following
    // parameters cannot be  configured. rttMonEchoAdminPktDataRequestSize
    // rttMonEchoAdminInterval rttMonEchoAdminNumPackets. The type is
    // RttMonCodecType.
    RttMonEchoAdminCodecType interface{}

    // This field represents the inter-packet delay between packets and is in
    // milliseconds. This object is applicable only to jitter probe which uses
    // codec type. The type is interface{} with range: 0..60000. Units are
    // milliseconds.
    RttMonEchoAdminCodecInterval interface{}

    // This object represents the number of octets that needs to be placed into
    // the Data portion of the message. This value is used only for jitter probe
    // which uses codec type. The type is interface{} with range: 0..16384. Units
    // are octets.
    RttMonEchoAdminCodecPayload interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is used only for jitter probe which uses codec type. The type is
    // interface{} with range: 0..60000.
    RttMonEchoAdminCodecNumPackets interface{}

    // The advantage factor is dependant on the type of access and how the service
    // is to be used. Conventional Wire-line     0 Mobility within Building    5
    // Mobility within geographic area  10 Access to hard-to-reach location   20 
    // This will be used while calculating the ICPIF values This valid only for
    // Jitter while calculating the ICPIF value. The type is interface{} with
    // range: 0..20.
    RttMonEchoAdminICPIFAdvFactor interface{}

    // The type of the target FEC for the RTT 'echo' and 'pathEcho' operations
    // based on 'mplsLspPingAppl' RttMonProtocol.  ldpIpv4Prefix   - LDP IPv4
    // prefix. The type is RttMonEchoAdminLSPFECType.
    RttMonEchoAdminLSPFECType interface{}

    // A string which specifies a valid 127/8 address. This address is of the form
    // 127.x.y.z. This address is not used to route the MPLS echo packet to the
    // destination but is used for load balancing in cases where the IP payload's
    // destination address is used for load balancing. The type is string.
    RttMonEchoAdminLSPSelector interface{}

    // This object specifies the reply mode for the LSP Echo requests. The type is
    // RttMonLSPPingReplyMode.
    RttMonEchoAdminLSPReplyMode interface{}

    // This object represents the TTL setting for MPLS echo request packets. For
    // ping operation this represents the TTL value to be set in the echo request
    // packet. For trace operation it represent the maximum ttl value that can be
    // set in the echo request packets starting with TTL=1.  For 'echo' based on
    // mplsLspPingAppl the default TTL will be set to 255, and for 'pathEcho'
    // based on mplsLspPingAppl the default will be set to 30.  Note: This object
    // cannot be set to the value of 0. The default value of 0 signifies the
    // default TTL values to be used for 'echo' and 'pathEcho' based on
    // 'mplsLspPingAppl'. The type is interface{} with range: 0..255.
    RttMonEchoAdminLSPTTL interface{}

    // This object represents the EXP value that needs to be put as precedence bit
    // in the MPLS echo request IP header. The type is interface{} with range:
    // 0..7.
    RttMonEchoAdminLSPExp interface{}

    // This object specifies the accuracy of statistics that needs to be
    // calculated milliseconds - The accuracy of stats will be of milliseconds
    // microseconds - The accuracy of stats will be in microseconds. This value
    // can be set only for jitter operation. The type is RttMonEchoAdminPrecision.
    RttMonEchoAdminPrecision interface{}

    // This object specifies the priority that will be assigned to probe packet. 
    // This value can be set only for jitter  operation. The type is
    // RttMonEchoAdminProbePakPriority.
    RttMonEchoAdminProbePakPriority interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered acceptable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of NTP offsets on source and responder. The value specified is
    // microseconds. This value can be set only for jitter operation  with
    // precision of microsecond. The type is interface{} with range:
    // -2147483648..2147483647. Units are microseconds.
    RttMonEchoAdminOWNTPSyncTolAbs interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered acceptable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of  NTP offsets on source and responder. The value is
    // expressed  as the percentage of actual oneway latency that is measured. 
    // This value can be set only for jitter operation with precision  of
    // microsecond. The type is interface{} with range: 0..100.
    RttMonEchoAdminOWNTPSyncTolPct interface{}

    // This object specifies whether the value in specified for oneway NTP sync
    // tolerance is absolute value or percent value. The type is
    // RttMonEchoAdminOWNTPSyncTolType.
    RttMonEchoAdminOWNTPSyncTolType interface{}

    // This string stores the called number of post dial delay. This object is
    // applicable to voip post dial delay probe only. The number will be like the
    // one actualy the user could dial. It has the number required by the local
    // country dial plan, plus E.164 number. The maximum length is 24 digits. Only
    // digit (0-9) is allowed. The type is string with length: 0..24.
    RttMonEchoAdminCalledNumber interface{}

    // A code that represents the detect point of post dial delay. This object is
    // applicable to SAA post dial delay probe only. The type is RttMonOperation.
    RttMonEchoAdminDetectPoint interface{}

    // A boolean that represents VoIP GK registration delay. This object is
    // applicable to SAA GK registration delay  probe only. The type is bool.
    RttMonEchoAdminGKRegistration interface{}

    // A string which specifies the voice-port on the source gateway. This object
    // is applicable to RTP probe only. The type is string.
    RttMonEchoAdminSourceVoicePort interface{}

    // Duration of RTP/Video Probe session. This object is applicable to RTP and
    // Video probe. The type is interface{} with range: 1..600.
    RttMonEchoAdminCallDuration interface{}

    // This object specifies the DSCP value to be set in the IP header of the LSP
    // echo reply packet. The value of this object will be in range of DiffServ
    // codepoint values between 0 to 63.  Note: This object cannot be set to value
    // of 255. This default value specifies that DSCP is not set for this row. The
    // type is interface{} with range: 0..63 | 255..None.
    RttMonEchoAdminLSPReplyDscp interface{}

    // This object specifies if the explicit-null label is to be added to LSP echo
    // requests which are sent while performing RTT operation. The type is bool.
    RttMonEchoAdminLSPNullShim interface{}

    // This object specifies the destination maintenance point ID. It is only
    // applicable to ethernetPing and ethernetJitter  operation. It will be set to
    // 0 for other types of  operations. The type is interface{} with range:
    // 0..8191.
    RttMonEchoAdminTargetMPID interface{}

    // This object specifies the name of the domain in which the destination
    // maintenance point lies. It is only applicable to  ethernetPing and
    // ethernetJitter operation. The type is string.
    RttMonEchoAdminTargetDomainName interface{}

    // This object specifies the ID of the VLAN in which the destination
    // maintenance point lies. It is only applicable to  ethernetPing and
    // ethernetJitter operation.  It will be set to 0 for other types of
    // operations. The type is interface{} with range: 1..4094.
    RttMonEchoAdminTargetVLAN interface{}

    // This object specifies the class of service in an Ethernet packet header. It
    // is only applicable to ethernetPing and  ethernetJitter operation. The type
    // is interface{} with range: 0..7.
    RttMonEchoAdminEthernetCOS interface{}

    // This object specifies MPLS LSP pseudowire VCCV ID values between 1 to
    // 2147483647.  Note: This object cannot be set to value of 0. This default
    // value specifies that VCCV is not set for this row. The type is interface{}
    // with range: 0..2147483647.
    RttMonEchoAdminLSPVccvID interface{}

    // This object specifies the Ethernet Virtual Connection in which the
    // destination maintenance point lies. It is only  applicable to ethernetPing
    // and ethernetJitter operation.  It will be set to NULL for other types of
    // operations. The type is string with length: 0..100.
    RttMonEchoAdminTargetEVC interface{}

    // This object specifies that Port Level CFM testing towards an Outward/Down
    // MEP will be used. It is only applicable to  ethernetPing and ethernetJitter
    // operation.  It will be set to NULL for other types of operations. The type
    // is bool.
    RttMonEchoAdminTargetMEPPort interface{}

    // A string which represents the profile name to which a video probe should
    // use. This object is applicable to video probe only. The type is string with
    // length: 0..255.
    RttMonEchoAdminVideoTrafficProfile interface{}

    // This object represents the Differentiated Service Code Point (DSCP) QoS
    // marking in the generated synthetic packets.  Value - DiffServ Class     0 -
    // BE (default)    10 - AF11    12 - AF12    14 - AF13    18 - AF21    20 -
    // AF22    22 - AF23    26 - AF31    28 - AF32    30 - AF33    34 - AF41    36
    // - AF42    38 - AF43     8 - CS1    16 - CS2    24 - CS3    32 - CS4    40 -
    // CS5    48 - CS6    56 - CS7    46 - EF. The type is interface{} with range:
    // 0..63.
    RttMonEchoAdminDscp interface{}

    // This object represents the video traffic generation source.  be : best
    // effort using DSP but without reservation gs : guaranteed service using DSP
    // with reservation na : not applicable for not using DSP. The type is
    // RttMonEchoAdminReserveDsp.
    RttMonEchoAdminReserveDsp interface{}

    // This object represents the network input interface on the sender router
    // where the synthetic packets are received from the emulated endpoint source.
    // This is used for path congruence with correct feature processing at the
    // sender router.  The user can get the InterfaceIndex number from ifIndex
    // object by looking up in ifTable. In fact, it should be useful to first get
    // the entry by the augmented table ifXTable which has ifName object which
    // matches the interface name used on the router or switch equipment console.
    // The type is interface{} with range: 0..2147483647.
    RttMonEchoAdminInputInterface interface{}

    // This object specifies the IP address of the emulated source from which the
    // synthetic packets would be generated. If this object is not specified, the
    // emulated source IP address will by default be the same as
    // rttMonEchoAdminSourceAddress. This object is applicable to video probes.
    // The type is string.
    RttMonEchoAdminEmulateSourceAddress interface{}

    // This object represents the port number of the emulated source from which
    // the synthetic packets would be generated. If this object is not specified,
    // the emulated source port number will by default be the same as
    // rttMonEchoAdminSourcePort. This object is applicable to video probes. The
    // type is interface{} with range: 0..65536.
    RttMonEchoAdminEmulateSourcePort interface{}

    // This object specifies the IP address of the emulated target by which the
    // synthetic packets would be received. If this object is not specified, the
    // emulated target IP address will by default be the same as
    // rttMonEchoAdminTargetAddress. This object is applicable to video probes.
    // The type is string.
    RttMonEchoAdminEmulateTargetAddress interface{}

    // This object represents the port number of the emulated target by which the
    // synthetic packets would be received. If this object is not specified, the
    // emulated target port number will by default be the same as
    // rttMonEchoAdminTargetPort. This object is applicable to video probes. The
    // type is interface{} with range: 0..65536.
    RttMonEchoAdminEmulateTargetPort interface{}

    // This object indicates the MAC address of the target device. This object is
    // only applicable for Y.1731 operations.  rttMonEchoAdminTargetMacAddress and
    // rttMonEchoAdminTargetMPID may not be used in conjunction. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RttMonEchoAdminTargetMacAddress interface{}

    // This object indicates the MAC address of the source device. This object is
    // only applicable for Y.1731 operations.  rttMonEchoAdminSourceMacAddress and
    // rttMonEchoAdminSourceMPID may not be used in conjunction. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RttMonEchoAdminSourceMacAddress interface{}

    // This object indicates the source maintenance point ID.  It is only
    // applicable to Y.1731 operation.  It will be set to zero for other types of
    // opearations.  rttMonEchoAdminSourceMPID and rttMonEchoAdminSourceMacAddress
    // may not be used in conjunction. The type is interface{} with range:
    // 0..8191.
    RttMonEchoAdminSourceMPID interface{}

    // This object specifies the name of endpoint list which a probe uses to
    // generate operations. The type is string with length: 1..64.
    RttMonEchoAdminEndPointListName interface{}

    // This object specifies if Source Specific Multicast is to be added. This
    // object is applicable to multicast probe only. The type is bool.
    RttMonEchoAdminSSM interface{}

    // This object specifies the maximum number of retries for control message.
    // The type is interface{} with range: 1..5.
    RttMonEchoAdminControlRetry interface{}

    // This object specifies the wait duration before control message timeout. The
    // type is interface{} with range: 1..10000. Units are milliseconds.
    RttMonEchoAdminControlTimeout interface{}

    // This object specifies number of packets to be sent for multicast tree
    // setup. This object is applicable to multicast probe only. The type is
    // interface{} with range: 0..10.
    RttMonEchoAdminIgmpTreeInit interface{}

    // This object indicates that packets will be sent in burst. The type is bool.
    RttMonEchoAdminEnableBurst interface{}

    // This object indicates the number of burst cycles to be sent during the
    // aggregate interval. This value is currently used for Y1731 SLM(Synthetic
    // Loss Measurment) probe. This object is applicable to Y1731 SLM probe only.
    // The type is interface{} with range: -2147483648..2147483647.
    RttMonEchoAdminAggBurstCycles interface{}

    // This object indicates the number of frames over which to calculate the
    // frame loss ratio. This object is applicable  to Y1731 SLM probe only. The
    // type is interface{} with range: -2147483648..2147483647.
    RttMonEchoAdminLossRatioNumFrames interface{}

    // This object indicates the number of frames over which to calculate the
    // availability. This object is applicable to Y1731 SLM probe only. The type
    // is interface{} with range: -2147483648..2147483647.
    RttMonEchoAdminAvailNumFrames interface{}

    // This object specifies whether timestamp optimization is enabled.  When the
    // value is 'true' then timestamp optimization is enabled.  The probe will
    // utilize lower layer (Hardware/Packet Processor) timestamping values to
    // improve accuracy of statistics.  This value can be set only for udp jitter
    // operation with precision of microsecond. The type is bool.
    RttMonEchoAdminTstampOptimization interface{}
}

func (rttMonEchoAdminEntry *CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry) GetEntityData() *types.CommonEntityData {
    rttMonEchoAdminEntry.EntityData.YFilter = rttMonEchoAdminEntry.YFilter
    rttMonEchoAdminEntry.EntityData.YangName = "rttMonEchoAdminEntry"
    rttMonEchoAdminEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonEchoAdminEntry.EntityData.ParentYangName = "rttMonEchoAdminTable"
    rttMonEchoAdminEntry.EntityData.SegmentPath = "rttMonEchoAdminEntry" + types.AddKeyToken(rttMonEchoAdminEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex")
    rttMonEchoAdminEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonEchoAdminEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonEchoAdminEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonEchoAdminEntry.EntityData.Children = types.NewOrderedMap()
    rttMonEchoAdminEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonEchoAdminEntry.RttMonCtrlAdminIndex})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminProtocol", types.YLeaf{"RttMonEchoAdminProtocol", rttMonEchoAdminEntry.RttMonEchoAdminProtocol})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetAddress", types.YLeaf{"RttMonEchoAdminTargetAddress", rttMonEchoAdminEntry.RttMonEchoAdminTargetAddress})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminPktDataRequestSize", types.YLeaf{"RttMonEchoAdminPktDataRequestSize", rttMonEchoAdminEntry.RttMonEchoAdminPktDataRequestSize})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminPktDataResponseSize", types.YLeaf{"RttMonEchoAdminPktDataResponseSize", rttMonEchoAdminEntry.RttMonEchoAdminPktDataResponseSize})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetPort", types.YLeaf{"RttMonEchoAdminTargetPort", rttMonEchoAdminEntry.RttMonEchoAdminTargetPort})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminSourceAddress", types.YLeaf{"RttMonEchoAdminSourceAddress", rttMonEchoAdminEntry.RttMonEchoAdminSourceAddress})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminSourcePort", types.YLeaf{"RttMonEchoAdminSourcePort", rttMonEchoAdminEntry.RttMonEchoAdminSourcePort})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminControlEnable", types.YLeaf{"RttMonEchoAdminControlEnable", rttMonEchoAdminEntry.RttMonEchoAdminControlEnable})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTOS", types.YLeaf{"RttMonEchoAdminTOS", rttMonEchoAdminEntry.RttMonEchoAdminTOS})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSREnable", types.YLeaf{"RttMonEchoAdminLSREnable", rttMonEchoAdminEntry.RttMonEchoAdminLSREnable})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetAddressString", types.YLeaf{"RttMonEchoAdminTargetAddressString", rttMonEchoAdminEntry.RttMonEchoAdminTargetAddressString})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminNameServer", types.YLeaf{"RttMonEchoAdminNameServer", rttMonEchoAdminEntry.RttMonEchoAdminNameServer})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminOperation", types.YLeaf{"RttMonEchoAdminOperation", rttMonEchoAdminEntry.RttMonEchoAdminOperation})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminHTTPVersion", types.YLeaf{"RttMonEchoAdminHTTPVersion", rttMonEchoAdminEntry.RttMonEchoAdminHTTPVersion})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminURL", types.YLeaf{"RttMonEchoAdminURL", rttMonEchoAdminEntry.RttMonEchoAdminURL})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminCache", types.YLeaf{"RttMonEchoAdminCache", rttMonEchoAdminEntry.RttMonEchoAdminCache})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminInterval", types.YLeaf{"RttMonEchoAdminInterval", rttMonEchoAdminEntry.RttMonEchoAdminInterval})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminNumPackets", types.YLeaf{"RttMonEchoAdminNumPackets", rttMonEchoAdminEntry.RttMonEchoAdminNumPackets})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminProxy", types.YLeaf{"RttMonEchoAdminProxy", rttMonEchoAdminEntry.RttMonEchoAdminProxy})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminString1", types.YLeaf{"RttMonEchoAdminString1", rttMonEchoAdminEntry.RttMonEchoAdminString1})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminString2", types.YLeaf{"RttMonEchoAdminString2", rttMonEchoAdminEntry.RttMonEchoAdminString2})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminString3", types.YLeaf{"RttMonEchoAdminString3", rttMonEchoAdminEntry.RttMonEchoAdminString3})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminString4", types.YLeaf{"RttMonEchoAdminString4", rttMonEchoAdminEntry.RttMonEchoAdminString4})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminString5", types.YLeaf{"RttMonEchoAdminString5", rttMonEchoAdminEntry.RttMonEchoAdminString5})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminMode", types.YLeaf{"RttMonEchoAdminMode", rttMonEchoAdminEntry.RttMonEchoAdminMode})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminVrfName", types.YLeaf{"RttMonEchoAdminVrfName", rttMonEchoAdminEntry.RttMonEchoAdminVrfName})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminCodecType", types.YLeaf{"RttMonEchoAdminCodecType", rttMonEchoAdminEntry.RttMonEchoAdminCodecType})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminCodecInterval", types.YLeaf{"RttMonEchoAdminCodecInterval", rttMonEchoAdminEntry.RttMonEchoAdminCodecInterval})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminCodecPayload", types.YLeaf{"RttMonEchoAdminCodecPayload", rttMonEchoAdminEntry.RttMonEchoAdminCodecPayload})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminCodecNumPackets", types.YLeaf{"RttMonEchoAdminCodecNumPackets", rttMonEchoAdminEntry.RttMonEchoAdminCodecNumPackets})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminICPIFAdvFactor", types.YLeaf{"RttMonEchoAdminICPIFAdvFactor", rttMonEchoAdminEntry.RttMonEchoAdminICPIFAdvFactor})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPFECType", types.YLeaf{"RttMonEchoAdminLSPFECType", rttMonEchoAdminEntry.RttMonEchoAdminLSPFECType})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPSelector", types.YLeaf{"RttMonEchoAdminLSPSelector", rttMonEchoAdminEntry.RttMonEchoAdminLSPSelector})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPReplyMode", types.YLeaf{"RttMonEchoAdminLSPReplyMode", rttMonEchoAdminEntry.RttMonEchoAdminLSPReplyMode})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPTTL", types.YLeaf{"RttMonEchoAdminLSPTTL", rttMonEchoAdminEntry.RttMonEchoAdminLSPTTL})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPExp", types.YLeaf{"RttMonEchoAdminLSPExp", rttMonEchoAdminEntry.RttMonEchoAdminLSPExp})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminPrecision", types.YLeaf{"RttMonEchoAdminPrecision", rttMonEchoAdminEntry.RttMonEchoAdminPrecision})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminProbePakPriority", types.YLeaf{"RttMonEchoAdminProbePakPriority", rttMonEchoAdminEntry.RttMonEchoAdminProbePakPriority})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminOWNTPSyncTolAbs", types.YLeaf{"RttMonEchoAdminOWNTPSyncTolAbs", rttMonEchoAdminEntry.RttMonEchoAdminOWNTPSyncTolAbs})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminOWNTPSyncTolPct", types.YLeaf{"RttMonEchoAdminOWNTPSyncTolPct", rttMonEchoAdminEntry.RttMonEchoAdminOWNTPSyncTolPct})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminOWNTPSyncTolType", types.YLeaf{"RttMonEchoAdminOWNTPSyncTolType", rttMonEchoAdminEntry.RttMonEchoAdminOWNTPSyncTolType})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminCalledNumber", types.YLeaf{"RttMonEchoAdminCalledNumber", rttMonEchoAdminEntry.RttMonEchoAdminCalledNumber})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminDetectPoint", types.YLeaf{"RttMonEchoAdminDetectPoint", rttMonEchoAdminEntry.RttMonEchoAdminDetectPoint})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminGKRegistration", types.YLeaf{"RttMonEchoAdminGKRegistration", rttMonEchoAdminEntry.RttMonEchoAdminGKRegistration})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminSourceVoicePort", types.YLeaf{"RttMonEchoAdminSourceVoicePort", rttMonEchoAdminEntry.RttMonEchoAdminSourceVoicePort})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminCallDuration", types.YLeaf{"RttMonEchoAdminCallDuration", rttMonEchoAdminEntry.RttMonEchoAdminCallDuration})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPReplyDscp", types.YLeaf{"RttMonEchoAdminLSPReplyDscp", rttMonEchoAdminEntry.RttMonEchoAdminLSPReplyDscp})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPNullShim", types.YLeaf{"RttMonEchoAdminLSPNullShim", rttMonEchoAdminEntry.RttMonEchoAdminLSPNullShim})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetMPID", types.YLeaf{"RttMonEchoAdminTargetMPID", rttMonEchoAdminEntry.RttMonEchoAdminTargetMPID})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetDomainName", types.YLeaf{"RttMonEchoAdminTargetDomainName", rttMonEchoAdminEntry.RttMonEchoAdminTargetDomainName})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetVLAN", types.YLeaf{"RttMonEchoAdminTargetVLAN", rttMonEchoAdminEntry.RttMonEchoAdminTargetVLAN})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminEthernetCOS", types.YLeaf{"RttMonEchoAdminEthernetCOS", rttMonEchoAdminEntry.RttMonEchoAdminEthernetCOS})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLSPVccvID", types.YLeaf{"RttMonEchoAdminLSPVccvID", rttMonEchoAdminEntry.RttMonEchoAdminLSPVccvID})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetEVC", types.YLeaf{"RttMonEchoAdminTargetEVC", rttMonEchoAdminEntry.RttMonEchoAdminTargetEVC})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetMEPPort", types.YLeaf{"RttMonEchoAdminTargetMEPPort", rttMonEchoAdminEntry.RttMonEchoAdminTargetMEPPort})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminVideoTrafficProfile", types.YLeaf{"RttMonEchoAdminVideoTrafficProfile", rttMonEchoAdminEntry.RttMonEchoAdminVideoTrafficProfile})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminDscp", types.YLeaf{"RttMonEchoAdminDscp", rttMonEchoAdminEntry.RttMonEchoAdminDscp})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminReserveDsp", types.YLeaf{"RttMonEchoAdminReserveDsp", rttMonEchoAdminEntry.RttMonEchoAdminReserveDsp})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminInputInterface", types.YLeaf{"RttMonEchoAdminInputInterface", rttMonEchoAdminEntry.RttMonEchoAdminInputInterface})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminEmulateSourceAddress", types.YLeaf{"RttMonEchoAdminEmulateSourceAddress", rttMonEchoAdminEntry.RttMonEchoAdminEmulateSourceAddress})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminEmulateSourcePort", types.YLeaf{"RttMonEchoAdminEmulateSourcePort", rttMonEchoAdminEntry.RttMonEchoAdminEmulateSourcePort})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminEmulateTargetAddress", types.YLeaf{"RttMonEchoAdminEmulateTargetAddress", rttMonEchoAdminEntry.RttMonEchoAdminEmulateTargetAddress})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminEmulateTargetPort", types.YLeaf{"RttMonEchoAdminEmulateTargetPort", rttMonEchoAdminEntry.RttMonEchoAdminEmulateTargetPort})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTargetMacAddress", types.YLeaf{"RttMonEchoAdminTargetMacAddress", rttMonEchoAdminEntry.RttMonEchoAdminTargetMacAddress})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminSourceMacAddress", types.YLeaf{"RttMonEchoAdminSourceMacAddress", rttMonEchoAdminEntry.RttMonEchoAdminSourceMacAddress})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminSourceMPID", types.YLeaf{"RttMonEchoAdminSourceMPID", rttMonEchoAdminEntry.RttMonEchoAdminSourceMPID})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminEndPointListName", types.YLeaf{"RttMonEchoAdminEndPointListName", rttMonEchoAdminEntry.RttMonEchoAdminEndPointListName})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminSSM", types.YLeaf{"RttMonEchoAdminSSM", rttMonEchoAdminEntry.RttMonEchoAdminSSM})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminControlRetry", types.YLeaf{"RttMonEchoAdminControlRetry", rttMonEchoAdminEntry.RttMonEchoAdminControlRetry})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminControlTimeout", types.YLeaf{"RttMonEchoAdminControlTimeout", rttMonEchoAdminEntry.RttMonEchoAdminControlTimeout})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminIgmpTreeInit", types.YLeaf{"RttMonEchoAdminIgmpTreeInit", rttMonEchoAdminEntry.RttMonEchoAdminIgmpTreeInit})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminEnableBurst", types.YLeaf{"RttMonEchoAdminEnableBurst", rttMonEchoAdminEntry.RttMonEchoAdminEnableBurst})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminAggBurstCycles", types.YLeaf{"RttMonEchoAdminAggBurstCycles", rttMonEchoAdminEntry.RttMonEchoAdminAggBurstCycles})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminLossRatioNumFrames", types.YLeaf{"RttMonEchoAdminLossRatioNumFrames", rttMonEchoAdminEntry.RttMonEchoAdminLossRatioNumFrames})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminAvailNumFrames", types.YLeaf{"RttMonEchoAdminAvailNumFrames", rttMonEchoAdminEntry.RttMonEchoAdminAvailNumFrames})
    rttMonEchoAdminEntry.EntityData.Leafs.Append("rttMonEchoAdminTstampOptimization", types.YLeaf{"RttMonEchoAdminTstampOptimization", rttMonEchoAdminEntry.RttMonEchoAdminTstampOptimization})

    rttMonEchoAdminEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex"}

    return &(rttMonEchoAdminEntry.EntityData)
}

// CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminLSPFECType represents ldpIpv4Prefix   - LDP IPv4 prefix.
type CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminLSPFECType string

const (
    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminLSPFECType_ldpIpv4Prefix CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminLSPFECType = "ldpIpv4Prefix"
)

// CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminOWNTPSyncTolType represents NTP sync tolerance is absolute value or percent value
type CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminOWNTPSyncTolType string

const (
    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminOWNTPSyncTolType_percent CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminOWNTPSyncTolType = "percent"

    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminOWNTPSyncTolType_absolute CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminOWNTPSyncTolType = "absolute"
)

// CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminPrecision represents This value can be set only for jitter operation
type CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminPrecision string

const (
    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminPrecision_milliseconds CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminPrecision = "milliseconds"

    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminPrecision_microseconds CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminPrecision = "microseconds"
)

// CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminProbePakPriority represents operation
type CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminProbePakPriority string

const (
    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminProbePakPriority_normal CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminProbePakPriority = "normal"

    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminProbePakPriority_high CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminProbePakPriority = "high"
)

// CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp represents na : not applicable for not using DSP
type CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp string

const (
    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp_be CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp = "be"

    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp_gs CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp = "gs"

    CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp_na CISCORTTMONMIB_RttMonEchoAdminTable_RttMonEchoAdminEntry_RttMonEchoAdminReserveDsp = "na"
)

// CISCORTTMONMIB_RttMonFileIOAdminTable
// A table of Round Trip Time (RTT) monitoring 'fileIO'
// specific definitions.
// 
// When the RttMonRttType is not 'fileIO' this table is
// not valid.
// 
// This table is controlled via the 
// rttMonCtrlAdminTable.  Entries in this table are
// created via the rttMonCtrlAdminStatus object.
type CISCORTTMONMIB_RttMonFileIOAdminTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for 'fileIO'
    // RttMonRttType conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry.
    RttMonFileIOAdminEntry []*CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry
}

func (rttMonFileIOAdminTable *CISCORTTMONMIB_RttMonFileIOAdminTable) GetEntityData() *types.CommonEntityData {
    rttMonFileIOAdminTable.EntityData.YFilter = rttMonFileIOAdminTable.YFilter
    rttMonFileIOAdminTable.EntityData.YangName = "rttMonFileIOAdminTable"
    rttMonFileIOAdminTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonFileIOAdminTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonFileIOAdminTable.EntityData.SegmentPath = "rttMonFileIOAdminTable"
    rttMonFileIOAdminTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonFileIOAdminTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonFileIOAdminTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonFileIOAdminTable.EntityData.Children = types.NewOrderedMap()
    rttMonFileIOAdminTable.EntityData.Children.Append("rttMonFileIOAdminEntry", types.YChild{"RttMonFileIOAdminEntry", nil})
    for i := range rttMonFileIOAdminTable.RttMonFileIOAdminEntry {
        rttMonFileIOAdminTable.EntityData.Children.Append(types.GetSegmentPath(rttMonFileIOAdminTable.RttMonFileIOAdminEntry[i]), types.YChild{"RttMonFileIOAdminEntry", rttMonFileIOAdminTable.RttMonFileIOAdminEntry[i]})
    }
    rttMonFileIOAdminTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonFileIOAdminTable.EntityData.YListKeys = []string {}

    return &(rttMonFileIOAdminTable.EntityData)
}

// CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry
// A list of objects that define specific configuration for
// 'fileIO' RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // The fully qualified file path that will be the target of the RTT operation.
    // This value must match one of the rttMonApplPreConfigedName entries. The
    // type is string.
    RttMonFileIOAdminFilePath interface{}

    // The size of the file to write/read from the File Server. The type is
    // RttMonFileIOAdminSize. Units are bytes.
    RttMonFileIOAdminSize interface{}

    // The File I/O action to be performed. The type is RttMonFileIOAdminAction.
    RttMonFileIOAdminAction interface{}
}

func (rttMonFileIOAdminEntry *CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry) GetEntityData() *types.CommonEntityData {
    rttMonFileIOAdminEntry.EntityData.YFilter = rttMonFileIOAdminEntry.YFilter
    rttMonFileIOAdminEntry.EntityData.YangName = "rttMonFileIOAdminEntry"
    rttMonFileIOAdminEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonFileIOAdminEntry.EntityData.ParentYangName = "rttMonFileIOAdminTable"
    rttMonFileIOAdminEntry.EntityData.SegmentPath = "rttMonFileIOAdminEntry" + types.AddKeyToken(rttMonFileIOAdminEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex")
    rttMonFileIOAdminEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonFileIOAdminEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonFileIOAdminEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonFileIOAdminEntry.EntityData.Children = types.NewOrderedMap()
    rttMonFileIOAdminEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonFileIOAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonFileIOAdminEntry.RttMonCtrlAdminIndex})
    rttMonFileIOAdminEntry.EntityData.Leafs.Append("rttMonFileIOAdminFilePath", types.YLeaf{"RttMonFileIOAdminFilePath", rttMonFileIOAdminEntry.RttMonFileIOAdminFilePath})
    rttMonFileIOAdminEntry.EntityData.Leafs.Append("rttMonFileIOAdminSize", types.YLeaf{"RttMonFileIOAdminSize", rttMonFileIOAdminEntry.RttMonFileIOAdminSize})
    rttMonFileIOAdminEntry.EntityData.Leafs.Append("rttMonFileIOAdminAction", types.YLeaf{"RttMonFileIOAdminAction", rttMonFileIOAdminEntry.RttMonFileIOAdminAction})

    rttMonFileIOAdminEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex"}

    return &(rttMonFileIOAdminEntry.EntityData)
}

// CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction represents The File I/O action to be performed.
type CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction string

const (
    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction_write CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction = "write"

    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction_read CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction = "read"

    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction_writeRead CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminAction = "writeRead"
)

// CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize represents Server.
type CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize string

const (
    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize_n256 CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize = "n256"

    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize_n1k CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize = "n1k"

    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize_n64k CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize = "n64k"

    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize_n128k CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize = "n128k"

    CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize_n256k CISCORTTMONMIB_RttMonFileIOAdminTable_RttMonFileIOAdminEntry_RttMonFileIOAdminSize = "n256k"
)

// CISCORTTMONMIB_RttMonScriptAdminTable
// A table of Round Trip Time (RTT) monitoring 'script'
// specific definitions.
// 
// When the RttMonRttType is not 'script' this table is
// not valid.
// 
// This table is controlled via the
// rttMonCtrlAdminTable.  Entries in this table are
// created via the rttMonCtrlAdminStatus object.
type CISCORTTMONMIB_RttMonScriptAdminTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for 'script'
    // RttMonRttType conceptual Rtt control rows. The type is slice of
    // CISCORTTMONMIB_RttMonScriptAdminTable_RttMonScriptAdminEntry.
    RttMonScriptAdminEntry []*CISCORTTMONMIB_RttMonScriptAdminTable_RttMonScriptAdminEntry
}

func (rttMonScriptAdminTable *CISCORTTMONMIB_RttMonScriptAdminTable) GetEntityData() *types.CommonEntityData {
    rttMonScriptAdminTable.EntityData.YFilter = rttMonScriptAdminTable.YFilter
    rttMonScriptAdminTable.EntityData.YangName = "rttMonScriptAdminTable"
    rttMonScriptAdminTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonScriptAdminTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonScriptAdminTable.EntityData.SegmentPath = "rttMonScriptAdminTable"
    rttMonScriptAdminTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonScriptAdminTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonScriptAdminTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonScriptAdminTable.EntityData.Children = types.NewOrderedMap()
    rttMonScriptAdminTable.EntityData.Children.Append("rttMonScriptAdminEntry", types.YChild{"RttMonScriptAdminEntry", nil})
    for i := range rttMonScriptAdminTable.RttMonScriptAdminEntry {
        rttMonScriptAdminTable.EntityData.Children.Append(types.GetSegmentPath(rttMonScriptAdminTable.RttMonScriptAdminEntry[i]), types.YChild{"RttMonScriptAdminEntry", rttMonScriptAdminTable.RttMonScriptAdminEntry[i]})
    }
    rttMonScriptAdminTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonScriptAdminTable.EntityData.YListKeys = []string {}

    return &(rttMonScriptAdminTable.EntityData)
}

// CISCORTTMONMIB_RttMonScriptAdminTable_RttMonScriptAdminEntry
// A list of objects that define specific configuration for
// 'script' RttMonRttType conceptual Rtt control rows.
type CISCORTTMONMIB_RttMonScriptAdminTable_RttMonScriptAdminEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This will be the Name of the Script that will be used to generate RTT
    // operations.    This object must match one of the  rttMonApplPreConfigedName
    // entries. The type is string.
    RttMonScriptAdminName interface{}

    // This will be the actual command line parameters passed to the
    // rttMonScriptAdminName when being executed. The type is string.
    RttMonScriptAdminCmdLineParams interface{}
}

func (rttMonScriptAdminEntry *CISCORTTMONMIB_RttMonScriptAdminTable_RttMonScriptAdminEntry) GetEntityData() *types.CommonEntityData {
    rttMonScriptAdminEntry.EntityData.YFilter = rttMonScriptAdminEntry.YFilter
    rttMonScriptAdminEntry.EntityData.YangName = "rttMonScriptAdminEntry"
    rttMonScriptAdminEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonScriptAdminEntry.EntityData.ParentYangName = "rttMonScriptAdminTable"
    rttMonScriptAdminEntry.EntityData.SegmentPath = "rttMonScriptAdminEntry" + types.AddKeyToken(rttMonScriptAdminEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex")
    rttMonScriptAdminEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonScriptAdminEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonScriptAdminEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonScriptAdminEntry.EntityData.Children = types.NewOrderedMap()
    rttMonScriptAdminEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonScriptAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonScriptAdminEntry.RttMonCtrlAdminIndex})
    rttMonScriptAdminEntry.EntityData.Leafs.Append("rttMonScriptAdminName", types.YLeaf{"RttMonScriptAdminName", rttMonScriptAdminEntry.RttMonScriptAdminName})
    rttMonScriptAdminEntry.EntityData.Leafs.Append("rttMonScriptAdminCmdLineParams", types.YLeaf{"RttMonScriptAdminCmdLineParams", rttMonScriptAdminEntry.RttMonScriptAdminCmdLineParams})

    rttMonScriptAdminEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex"}

    return &(rttMonScriptAdminEntry.EntityData)
}

// CISCORTTMONMIB_RttMonReactTriggerAdminTable
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
type CISCORTTMONMIB_RttMonReactTriggerAdminTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that will be triggered when a reaction condition is
    // violated. The type is slice of
    // CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry.
    RttMonReactTriggerAdminEntry []*CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry
}

func (rttMonReactTriggerAdminTable *CISCORTTMONMIB_RttMonReactTriggerAdminTable) GetEntityData() *types.CommonEntityData {
    rttMonReactTriggerAdminTable.EntityData.YFilter = rttMonReactTriggerAdminTable.YFilter
    rttMonReactTriggerAdminTable.EntityData.YangName = "rttMonReactTriggerAdminTable"
    rttMonReactTriggerAdminTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonReactTriggerAdminTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonReactTriggerAdminTable.EntityData.SegmentPath = "rttMonReactTriggerAdminTable"
    rttMonReactTriggerAdminTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonReactTriggerAdminTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonReactTriggerAdminTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonReactTriggerAdminTable.EntityData.Children = types.NewOrderedMap()
    rttMonReactTriggerAdminTable.EntityData.Children.Append("rttMonReactTriggerAdminEntry", types.YChild{"RttMonReactTriggerAdminEntry", nil})
    for i := range rttMonReactTriggerAdminTable.RttMonReactTriggerAdminEntry {
        rttMonReactTriggerAdminTable.EntityData.Children.Append(types.GetSegmentPath(rttMonReactTriggerAdminTable.RttMonReactTriggerAdminEntry[i]), types.YChild{"RttMonReactTriggerAdminEntry", rttMonReactTriggerAdminTable.RttMonReactTriggerAdminEntry[i]})
    }
    rttMonReactTriggerAdminTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonReactTriggerAdminTable.EntityData.YListKeys = []string {}

    return &(rttMonReactTriggerAdminTable.EntityData)
}

// CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry
// A list of objects that will be triggered when
// a reaction condition is violated.
type CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. This object points to a single conceptual Rtt
    // control row.  If this row does not exist and this value is  triggered no
    // action will result.  The conceptual Rtt control row will be triggered for
    // the  rttMonCtrlOperRttLife length.  If this conceptual Rtt control row is
    // already active, rttMonCtrlOperRttLife will not be updated, and its life
    // will continue as previously  defined. The type is interface{} with range:
    // 1..2147483647.
    RttMonReactTriggerAdminRttMonCtrlAdminIndex interface{}

    // This object is used to create Trigger entries. The type is RowStatus.
    RttMonReactTriggerAdminStatus interface{}

    // This object takes on the value active when its associated entry in the 
    // rttMonReactTriggerAdminTable has been triggered.  When the associated entry
    // in the rttMonReactTriggerAdminTable is not under a trigger state, this
    // object will be pending.  When this object is in the active state this entry
    // can not be retriggered. The type is RttMonReactTriggerOperState.
    RttMonReactTriggerOperState interface{}
}

func (rttMonReactTriggerAdminEntry *CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry) GetEntityData() *types.CommonEntityData {
    rttMonReactTriggerAdminEntry.EntityData.YFilter = rttMonReactTriggerAdminEntry.YFilter
    rttMonReactTriggerAdminEntry.EntityData.YangName = "rttMonReactTriggerAdminEntry"
    rttMonReactTriggerAdminEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonReactTriggerAdminEntry.EntityData.ParentYangName = "rttMonReactTriggerAdminTable"
    rttMonReactTriggerAdminEntry.EntityData.SegmentPath = "rttMonReactTriggerAdminEntry" + types.AddKeyToken(rttMonReactTriggerAdminEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonReactTriggerAdminEntry.RttMonReactTriggerAdminRttMonCtrlAdminIndex, "rttMonReactTriggerAdminRttMonCtrlAdminIndex")
    rttMonReactTriggerAdminEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonReactTriggerAdminEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonReactTriggerAdminEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonReactTriggerAdminEntry.EntityData.Children = types.NewOrderedMap()
    rttMonReactTriggerAdminEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonReactTriggerAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonReactTriggerAdminEntry.RttMonCtrlAdminIndex})
    rttMonReactTriggerAdminEntry.EntityData.Leafs.Append("rttMonReactTriggerAdminRttMonCtrlAdminIndex", types.YLeaf{"RttMonReactTriggerAdminRttMonCtrlAdminIndex", rttMonReactTriggerAdminEntry.RttMonReactTriggerAdminRttMonCtrlAdminIndex})
    rttMonReactTriggerAdminEntry.EntityData.Leafs.Append("rttMonReactTriggerAdminStatus", types.YLeaf{"RttMonReactTriggerAdminStatus", rttMonReactTriggerAdminEntry.RttMonReactTriggerAdminStatus})
    rttMonReactTriggerAdminEntry.EntityData.Leafs.Append("rttMonReactTriggerOperState", types.YLeaf{"RttMonReactTriggerOperState", rttMonReactTriggerAdminEntry.RttMonReactTriggerOperState})

    rttMonReactTriggerAdminEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonReactTriggerAdminRttMonCtrlAdminIndex"}

    return &(rttMonReactTriggerAdminEntry.EntityData)
}

// CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry_RttMonReactTriggerOperState represents this entry can not be retriggered.
type CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry_RttMonReactTriggerOperState string

const (
    CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry_RttMonReactTriggerOperState_active CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry_RttMonReactTriggerOperState = "active"

    CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry_RttMonReactTriggerOperState_pending CISCORTTMONMIB_RttMonReactTriggerAdminTable_RttMonReactTriggerAdminEntry_RttMonReactTriggerOperState = "pending"
)

// CISCORTTMONMIB_RttMonEchoPathAdminTable
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
type CISCORTTMONMIB_RttMonEchoPathAdminTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define intermediate hop's IP Address.  This entry
    // can be added only if the rttMonCtrlAdminRttType is 'echo'. The entry gets
    // deleted when the corresponding RTR entry, which has an index of
    // rttMonCtrlAdminIndex, is deleted. The type is slice of
    // CISCORTTMONMIB_RttMonEchoPathAdminTable_RttMonEchoPathAdminEntry.
    RttMonEchoPathAdminEntry []*CISCORTTMONMIB_RttMonEchoPathAdminTable_RttMonEchoPathAdminEntry
}

func (rttMonEchoPathAdminTable *CISCORTTMONMIB_RttMonEchoPathAdminTable) GetEntityData() *types.CommonEntityData {
    rttMonEchoPathAdminTable.EntityData.YFilter = rttMonEchoPathAdminTable.YFilter
    rttMonEchoPathAdminTable.EntityData.YangName = "rttMonEchoPathAdminTable"
    rttMonEchoPathAdminTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonEchoPathAdminTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonEchoPathAdminTable.EntityData.SegmentPath = "rttMonEchoPathAdminTable"
    rttMonEchoPathAdminTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonEchoPathAdminTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonEchoPathAdminTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonEchoPathAdminTable.EntityData.Children = types.NewOrderedMap()
    rttMonEchoPathAdminTable.EntityData.Children.Append("rttMonEchoPathAdminEntry", types.YChild{"RttMonEchoPathAdminEntry", nil})
    for i := range rttMonEchoPathAdminTable.RttMonEchoPathAdminEntry {
        rttMonEchoPathAdminTable.EntityData.Children.Append(types.GetSegmentPath(rttMonEchoPathAdminTable.RttMonEchoPathAdminEntry[i]), types.YChild{"RttMonEchoPathAdminEntry", rttMonEchoPathAdminTable.RttMonEchoPathAdminEntry[i]})
    }
    rttMonEchoPathAdminTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonEchoPathAdminTable.EntityData.YListKeys = []string {}

    return &(rttMonEchoPathAdminTable.EntityData)
}

// CISCORTTMONMIB_RttMonEchoPathAdminTable_RttMonEchoPathAdminEntry
// A list of objects that define intermediate hop's IP Address.
// 
// This entry can be added only if the rttMonCtrlAdminRttType is
// 'echo'. The entry gets deleted when the corresponding RTR entry,
// which has an index of rttMonCtrlAdminIndex, is deleted.
type CISCORTTMONMIB_RttMonEchoPathAdminTable_RttMonEchoPathAdminEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonEchoPathAdminTable. This number represents the hop address number in
    // a specific  ping path. All the indicies should start from 1 and must be 
    // contiguous ie., entries should be  (say rttMonCtrlAdminIndex = 1)  1.1,
    // 1.2, 1.3, they cannot be 1.1, 1.2, 1.4. The type is interface{} with range:
    // 1..8.
    RttMonEchoPathAdminHopIndex interface{}

    // A string which specifies the address of an intermediate hop's IP Address
    // for a RTT 'echo' operation. The type is string.
    RttMonEchoPathAdminHopAddress interface{}
}

func (rttMonEchoPathAdminEntry *CISCORTTMONMIB_RttMonEchoPathAdminTable_RttMonEchoPathAdminEntry) GetEntityData() *types.CommonEntityData {
    rttMonEchoPathAdminEntry.EntityData.YFilter = rttMonEchoPathAdminEntry.YFilter
    rttMonEchoPathAdminEntry.EntityData.YangName = "rttMonEchoPathAdminEntry"
    rttMonEchoPathAdminEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonEchoPathAdminEntry.EntityData.ParentYangName = "rttMonEchoPathAdminTable"
    rttMonEchoPathAdminEntry.EntityData.SegmentPath = "rttMonEchoPathAdminEntry" + types.AddKeyToken(rttMonEchoPathAdminEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonEchoPathAdminEntry.RttMonEchoPathAdminHopIndex, "rttMonEchoPathAdminHopIndex")
    rttMonEchoPathAdminEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonEchoPathAdminEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonEchoPathAdminEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonEchoPathAdminEntry.EntityData.Children = types.NewOrderedMap()
    rttMonEchoPathAdminEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonEchoPathAdminEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonEchoPathAdminEntry.RttMonCtrlAdminIndex})
    rttMonEchoPathAdminEntry.EntityData.Leafs.Append("rttMonEchoPathAdminHopIndex", types.YLeaf{"RttMonEchoPathAdminHopIndex", rttMonEchoPathAdminEntry.RttMonEchoPathAdminHopIndex})
    rttMonEchoPathAdminEntry.EntityData.Leafs.Append("rttMonEchoPathAdminHopAddress", types.YLeaf{"RttMonEchoPathAdminHopAddress", rttMonEchoPathAdminEntry.RttMonEchoPathAdminHopAddress})

    rttMonEchoPathAdminEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonEchoPathAdminHopIndex"}

    return &(rttMonEchoPathAdminEntry.EntityData)
}

// CISCORTTMONMIB_RttMonGrpScheduleAdminTable
// A table of Round Trip Time (RTT) monitoring group scheduling
// specific definitions.
// This table is used to create a conceptual group scheduling
// control row. The entries in this control row contain objects
// used to define group schedule configuration parameters.
// 
// The objects of this table will be used to schedule a group of
// probes identified by the conceptual rows of the
// rttMonCtrlAdminTable.
type CISCORTTMONMIB_RttMonGrpScheduleAdminTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define a conceptual group scheduling control row.
    // The type is slice of
    // CISCORTTMONMIB_RttMonGrpScheduleAdminTable_RttMonGrpScheduleAdminEntry.
    RttMonGrpScheduleAdminEntry []*CISCORTTMONMIB_RttMonGrpScheduleAdminTable_RttMonGrpScheduleAdminEntry
}

func (rttMonGrpScheduleAdminTable *CISCORTTMONMIB_RttMonGrpScheduleAdminTable) GetEntityData() *types.CommonEntityData {
    rttMonGrpScheduleAdminTable.EntityData.YFilter = rttMonGrpScheduleAdminTable.YFilter
    rttMonGrpScheduleAdminTable.EntityData.YangName = "rttMonGrpScheduleAdminTable"
    rttMonGrpScheduleAdminTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonGrpScheduleAdminTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonGrpScheduleAdminTable.EntityData.SegmentPath = "rttMonGrpScheduleAdminTable"
    rttMonGrpScheduleAdminTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonGrpScheduleAdminTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonGrpScheduleAdminTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonGrpScheduleAdminTable.EntityData.Children = types.NewOrderedMap()
    rttMonGrpScheduleAdminTable.EntityData.Children.Append("rttMonGrpScheduleAdminEntry", types.YChild{"RttMonGrpScheduleAdminEntry", nil})
    for i := range rttMonGrpScheduleAdminTable.RttMonGrpScheduleAdminEntry {
        rttMonGrpScheduleAdminTable.EntityData.Children.Append(types.GetSegmentPath(rttMonGrpScheduleAdminTable.RttMonGrpScheduleAdminEntry[i]), types.YChild{"RttMonGrpScheduleAdminEntry", rttMonGrpScheduleAdminTable.RttMonGrpScheduleAdminEntry[i]})
    }
    rttMonGrpScheduleAdminTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonGrpScheduleAdminTable.EntityData.YListKeys = []string {}

    return &(rttMonGrpScheduleAdminTable.EntityData)
}

// CISCORTTMONMIB_RttMonGrpScheduleAdminTable_RttMonGrpScheduleAdminEntry
// A list of objects that define a conceptual group scheduling
// control row.
type CISCORTTMONMIB_RttMonGrpScheduleAdminTable_RttMonGrpScheduleAdminEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMonGrpScheduleAdminTable.  This is a pseudo-random number selected by
    // the management station when creating a row via the
    // rttMonGrpScheduleAdminStatus object. If the pseudo-random number is already
    // in use an 'inconsistentValue' return code will be returned when set
    // operation is attempted. The type is interface{} with range: 1..2147483647.
    RttMonGrpScheduleAdminIndex interface{}

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
    RttMonGrpScheduleAdminProbes interface{}

    // Specifies the time duration over which all the probes have to be scheduled.
    // The type is interface{} with range: 0..604800. Units are seconds.
    RttMonGrpScheduleAdminPeriod interface{}

    // Specifies the duration between initiating each RTT operation for all the
    // probes specified in the group.  The value of this object is only effective
    // when both rttMonGrpScheduleAdminFreqMax and rttMonGrpScheduleAdminFreqMin 
    // have zero values. The type is interface{} with range: 0..604800. Units are
    // seconds.
    RttMonGrpScheduleAdminFrequency interface{}

    // This object specifies the life of all the probes included in the object
    // rttMonGrpScheduleAdminProbes, that are getting group scheduled. This value
    // will be placed into rttMonScheduleAdminRttLife object for each of the
    // probes listed in rttMonGrpScheduleAdminProbes when this conceptual control
    // row becomes 'active'.  The value 2147483647 has a special meaning. When
    // this object is set to 2147483647, the rttMonCtrlOperRttLife object for all
    // the probes listed in rttMonGrpScheduleAdminProbes,  will not decrement. And
    // thus the life time of the probes will never end. The type is interface{}
    // with range: 0..2147483647. Units are seconds.
    RttMonGrpScheduleAdminLife interface{}

    // This object specifies the ageout value of all the probes included in the
    // object rttMonGrpScheduleAdminProbes, that are getting group scheduled. This
    // value will be placed into rttMonScheduleAdminConceptRowAgeout object for
    // each of the probes listed in rttMonGrpScheduleAdminProbes when this
    // conceptual control row becomes 'active'.  When this value is set to zero,
    // the probes listed in rttMonGrpScheduleAdminProbes, will never ageout. The
    // type is interface{} with range: 0..2073600. Units are seconds.
    RttMonGrpScheduleAdminAgeout interface{}

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
    RttMonGrpScheduleAdminStatus interface{}

    // Specifies the max duration between initiating each RTT operation for all
    // the probes specified in the group.  If this is 0 and
    // rttMonGrpScheduleAdminFreqMin is also 0 then
    // rttMonGrpScheduleAdminFrequency becomes the fixed frequency. The type is
    // interface{} with range: 0..604800. Units are seconds.
    RttMonGrpScheduleAdminFreqMax interface{}

    // Specifies the min duration between initiating each RTT operation for all
    // the probes specified in the group.  The value of this object cannot be
    // greater than the value of rttMonGrpScheduleAdminFreqMax.  If this is 0 and
    // rttMonGrpScheduleAdminFreqMax is 0 then rttMonGrpScheduleAdminFrequency
    // becomes the fixed frequency. The type is interface{} with range: 0..604800.
    // Units are seconds.
    RttMonGrpScheduleAdminFreqMin interface{}

    // This is the time in seconds after which the member probes of this group
    // specified in rttMonGrpScheduleAdminProbes will transition to active state.
    // The type is interface{} with range: 0..604800. Units are seconds.
    RttMonGrpScheduleAdminStartTime interface{}

    // Addition of members to an existing group will be allowed if this object is
    // set to TRUE (1). The members, IDs of which are mentioned in
    // rttMonGrpScheduleAdminProbes object are added to the existing group. The
    // type is bool.
    RttMonGrpScheduleAdminAdd interface{}

    // Removal of members from an existing group will be allowed if this object is
    // set to TRUE (1). The members, IDs of which are mentioned in
    // rttMonGrpScheduleAdminProbes object are removed from the existing group.
    // The type is bool.
    RttMonGrpScheduleAdminDelete interface{}

    // When this is set to true then all members of this group will be stopped and
    // rescheduled using the previously set values of this group. The type is
    // bool.
    RttMonGrpScheduleAdminReset interface{}
}

func (rttMonGrpScheduleAdminEntry *CISCORTTMONMIB_RttMonGrpScheduleAdminTable_RttMonGrpScheduleAdminEntry) GetEntityData() *types.CommonEntityData {
    rttMonGrpScheduleAdminEntry.EntityData.YFilter = rttMonGrpScheduleAdminEntry.YFilter
    rttMonGrpScheduleAdminEntry.EntityData.YangName = "rttMonGrpScheduleAdminEntry"
    rttMonGrpScheduleAdminEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonGrpScheduleAdminEntry.EntityData.ParentYangName = "rttMonGrpScheduleAdminTable"
    rttMonGrpScheduleAdminEntry.EntityData.SegmentPath = "rttMonGrpScheduleAdminEntry" + types.AddKeyToken(rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminIndex, "rttMonGrpScheduleAdminIndex")
    rttMonGrpScheduleAdminEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonGrpScheduleAdminEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonGrpScheduleAdminEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonGrpScheduleAdminEntry.EntityData.Children = types.NewOrderedMap()
    rttMonGrpScheduleAdminEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminIndex", types.YLeaf{"RttMonGrpScheduleAdminIndex", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminIndex})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminProbes", types.YLeaf{"RttMonGrpScheduleAdminProbes", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminProbes})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminPeriod", types.YLeaf{"RttMonGrpScheduleAdminPeriod", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminPeriod})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminFrequency", types.YLeaf{"RttMonGrpScheduleAdminFrequency", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminFrequency})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminLife", types.YLeaf{"RttMonGrpScheduleAdminLife", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminLife})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminAgeout", types.YLeaf{"RttMonGrpScheduleAdminAgeout", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminAgeout})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminStatus", types.YLeaf{"RttMonGrpScheduleAdminStatus", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminStatus})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminFreqMax", types.YLeaf{"RttMonGrpScheduleAdminFreqMax", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminFreqMax})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminFreqMin", types.YLeaf{"RttMonGrpScheduleAdminFreqMin", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminFreqMin})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminStartTime", types.YLeaf{"RttMonGrpScheduleAdminStartTime", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminStartTime})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminAdd", types.YLeaf{"RttMonGrpScheduleAdminAdd", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminAdd})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminDelete", types.YLeaf{"RttMonGrpScheduleAdminDelete", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminDelete})
    rttMonGrpScheduleAdminEntry.EntityData.Leafs.Append("rttMonGrpScheduleAdminReset", types.YLeaf{"RttMonGrpScheduleAdminReset", rttMonGrpScheduleAdminEntry.RttMonGrpScheduleAdminReset})

    rttMonGrpScheduleAdminEntry.EntityData.YListKeys = []string {"RttMonGrpScheduleAdminIndex"}

    return &(rttMonGrpScheduleAdminEntry.EntityData)
}

// CISCORTTMONMIB_RttMplsVpnMonCtrlTable
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
type CISCORTTMONMIB_RttMplsVpnMonCtrlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual Auto SAA L3 MPLS VPN
    // control row. The type is slice of
    // CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry.
    RttMplsVpnMonCtrlEntry []*CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry
}

func (rttMplsVpnMonCtrlTable *CISCORTTMONMIB_RttMplsVpnMonCtrlTable) GetEntityData() *types.CommonEntityData {
    rttMplsVpnMonCtrlTable.EntityData.YFilter = rttMplsVpnMonCtrlTable.YFilter
    rttMplsVpnMonCtrlTable.EntityData.YangName = "rttMplsVpnMonCtrlTable"
    rttMplsVpnMonCtrlTable.EntityData.BundleName = "cisco_ios_xe"
    rttMplsVpnMonCtrlTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMplsVpnMonCtrlTable.EntityData.SegmentPath = "rttMplsVpnMonCtrlTable"
    rttMplsVpnMonCtrlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMplsVpnMonCtrlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMplsVpnMonCtrlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMplsVpnMonCtrlTable.EntityData.Children = types.NewOrderedMap()
    rttMplsVpnMonCtrlTable.EntityData.Children.Append("rttMplsVpnMonCtrlEntry", types.YChild{"RttMplsVpnMonCtrlEntry", nil})
    for i := range rttMplsVpnMonCtrlTable.RttMplsVpnMonCtrlEntry {
        rttMplsVpnMonCtrlTable.EntityData.Children.Append(types.GetSegmentPath(rttMplsVpnMonCtrlTable.RttMplsVpnMonCtrlEntry[i]), types.YChild{"RttMplsVpnMonCtrlEntry", rttMplsVpnMonCtrlTable.RttMplsVpnMonCtrlEntry[i]})
    }
    rttMplsVpnMonCtrlTable.EntityData.Leafs = types.NewOrderedMap()

    rttMplsVpnMonCtrlTable.EntityData.YListKeys = []string {}

    return &(rttMplsVpnMonCtrlTable.EntityData)
}

// CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry
// A base list of objects that define a conceptual Auto SAA L3
// MPLS VPN control row.
type CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in the
    // rttMplsVpnMonCtrlTable.  This is a pseudo-random number selected by the
    // management station when creating a row via the rttMplsVpnMonCtrlStatus
    // object.  If the pseudo-random number is already in use an
    // 'inconsistentValue' return code will be returned when set operation is
    // attempted. The type is interface{} with range: 1..2147483647.
    RttMplsVpnMonCtrlIndex interface{}

    // The type of RTT operation to be performed for Auto SAA L3 MPLS VPN.  This
    // value must be set in the same PDU of rttMplsVpnMonCtrlStatus.  This value
    // must be set before setting any other parameter configuration of an Auto SAA
    // L3 MPLS VPN. The type is RttMplsVpnMonRttType.
    RttMplsVpnMonCtrlRttType interface{}

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
    RttMplsVpnMonCtrlVrfName interface{}

    // A string which is used by a managing application to identify the RTT
    // target.  This string will be configured as rttMonCtrlAdminTag for all the
    // operations configured by this Auto SAA L3 MPLS VPN.  The usage of this
    // value in Auto SAA L3 MPLS VPN is same as rttMonCtrlAdminTag in RTT
    // operation. The type is string with length: 0..255.
    RttMplsVpnMonCtrlTag interface{}

    // This object defines an administrative threshold limit.  This value will be
    // configured as rttMonCtrlAdminThreshold for all the operations that will be
    // configured by the current Auto SAA L3 MPLS VPN.  The usage of this value in
    // Auto SAA L3 MPLS VPN is same as rttMonCtrlAdminThreshold. The type is
    // interface{} with range: 0..2147483647. Units are milliseconds.
    RttMplsVpnMonCtrlThreshold interface{}

    // Specifies the duration to wait for a RTT operation configured automatically
    // by the Auto SAA L3 MPLS VPN to complete.   The value of this object cannot
    // be set to a value which would specify a duration exceeding
    // rttMplsVpnMonScheduleFrequency.  The usage of this value in Auto SAA L3
    // MPLS VPN is similar to rttMonCtrlAdminTimeout. The type is interface{} with
    // range: 0..604800000. Units are milliseconds.
    RttMplsVpnMonCtrlTimeout interface{}

    // Specifies the frequency at which the automatic PE addition should take
    // place if there is any for an Auto SAA L3 MPLS VPN.  New RTT operations
    // corresponding to the new PEs discovered will be created and scheduled.  The
    // default value for this object is 4 hours. The maximum value supported is 49
    // days. The type is interface{} with range: 1..70560. Units are minutes.
    RttMplsVpnMonCtrlScanInterval interface{}

    // Specifies the frequency at which the automatic PE deletion should take
    // place.  This object specifies the number of times of
    // rttMonMplslmCtrlScanInterval (rttMplsVpnMonCtrlDelScanFactor *
    // rttMplsVpnMonCtrlScanInterval) to wait before removing the PEs. This object
    // doesn't directly specify the explicit value to wait before removing the PEs
    // that were down.  If this object set 0 the entries will never removed. The
    // type is interface{} with range: 0..2147483647.
    RttMplsVpnMonCtrlDelScanFactor interface{}

    // This object represents the EXP value that needs to be put as precedence bit
    // of an IP header. The type is interface{} with range: 0..7.
    RttMplsVpnMonCtrlEXP interface{}

    // This object represents the native payload size that needs to be put on the
    // packet.  This value will be configured as rttMonEchoAdminPktDataRequestSize
    // for all the RTT operations configured by the current Auto SAA L3 MPLS VPN. 
    // The minimum request size for jitter probe is 16. The maximum for jitter
    // probe is 1500. The default request size is 32 for jitter probe.  For echo
    // and pathEcho default request size is 28. The minimum request size for echo
    // and pathEcho is 28 bytes. The type is interface{} with range: 0..16384.
    // Units are octets.
    RttMplsVpnMonCtrlRequestSize interface{}

    // When set to true, the resulting data in each RTT operation created by the
    // current Auto SAA L3 MPLS VPN is compared with the expected data. This
    // includes checking header information (if possible) and exact packet size. 
    // Any mismatch will be recorded in the rttMonStatsCollectVerifyErrors object
    // of each RTT operation created by the current Auto SAA L3 MPLS VPN. The type
    // is bool.
    RttMplsVpnMonCtrlVerifyData interface{}

    // The storage type of this conceptual row. When set to 'nonVolatile', this
    // entry will be shown in 'show running' command and can be saved into
    // Non-volatile memory.  By Default the entry will not be saved into
    // Non-volatile memory.  This object can be set to either 'volatile' or
    // 'nonVolatile'. Other values are not applicable for this conceptual row and
    // are not supported. The type is StorageType.
    RttMplsVpnMonCtrlStorageType interface{}

    // This object holds the list of probes ID's that are created by the Auto SAA
    // L3 MPLS VPN.  The probes will be specified in the following form. (a)
    // Individual ID's with comma separated as 1,5,3. (b) Range form including
    // hyphens with multiple ranges being     separated by comma as 1-10,12-34.
    // (c) Mix of the above two forms as 1,2,4-10,12,15,19-25. The type is string.
    RttMplsVpnMonCtrlProbeList interface{}

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
    RttMplsVpnMonCtrlStatus interface{}

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
    RttMplsVpnMonCtrlLpd interface{}

    // This object holds the list of LPD Group IDs that are created for this Auto
    // SAA L3 MPLS VPN row.  This object will be applicable only when LSP Path
    // Discovery is enabled for this row.  The LPD Groups will be specified in the
    // following form. (a) Individual ID's with comma separated as 1,5,3. (b)
    // Range form including hyphens with multiple ranges being     separated by
    // comma as 1-10,12-34. (c) Mix of the above two forms as
    // 1,2,4-10,12,15,19-25. The type is string.
    RttMplsVpnMonCtrlLpdGrpList interface{}

    // The completion time of the LSP Path Discovery for the entire set of PEs
    // which are discovered for this Auto SAA.  This object will be applicable
    // only when LSP Path Discovery is enabled for this row. The type is
    // interface{} with range: 1..65535. Units are minutes.
    RttMplsVpnMonCtrlLpdCompTime interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. This value is currently used for Jitter probe. This object is
    // applicable to jitter probe only.  The usage of this value in RTT operation
    // is same as rttMonEchoAdminInterval. The type is interface{} with range:
    // 1..60000. Units are milliseconds.
    RttMplsVpnMonTypeInterval interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is currently used for Jitter probe. This object is applicable to
    // jitter probe only.  The usage of this value in RTT operation is same as
    // rttMonEchoAdminNumPackets. The type is interface{} with range: 1..60000.
    RttMplsVpnMonTypeNumPackets interface{}

    // This object represents the target's port number to which the packets need
    // to be sent.  This value will be configured as target port for all the
    // operations that is going to be configured   The usage of this value is same
    // as rttMonEchoAdminTargetPort in RTT operation. This object is applicable to
    // jitter type.  If this object is not being set random port will be used as
    // destination port. The type is interface{} with range: 1..65536.
    RttMplsVpnMonTypeDestPort interface{}

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
    // trap is sent. The type is RttMplsVpnMonTypeSecFreqType.
    RttMplsVpnMonTypeSecFreqType interface{}

    // This object represents the value that needs to be applied to secondary
    // frequency of individual RTT operations configured by Auto SAA L3 MPLS VPN. 
    // Setting rttMplsVpnMonTypeSecFreqValue without setting
    // rttMplsVpnMonTypeSecFreqType will not have any effect. The type is
    // interface{} with range: 1..604800.
    RttMplsVpnMonTypeSecFreqValue interface{}

    // A string which specifies the address of the local host (127.X.X.X).  This
    // object will be used as lsp-selector in MPLS RTT operations configured by
    // the Auto SAA L3 MPLS VPN.  When LSP Path Discovery is enabled for the row,
    // this object will be used to indicate the base LSP selector value to be used
    // in the LSP Path Discovery.  This value of this object is significant in
    // MPLS load balancing scenario. This value will be used as one of the
    // parameter in that load balancing. The type is string.
    RttMplsVpnMonTypeLspSelector interface{}

    // This object specifies the reply mode for the LSP Echo requests originated
    // by the operations configured by the Auto SAA L3 MPLS VPN.  This object is
    // currently used by echo and pathEcho. The type is RttMonLSPPingReplyMode.
    RttMplsVpnMonTypeLSPReplyMode interface{}

    // This object represents the TTL setting for MPLS echo request packets
    // originated by the operations configured by the Auto SAA L3 MPLS VPN.  This
    // object is currently used by echo and pathEcho.  For 'echo' the default TTL
    // will be set to 255. For 'pathEcho' the default will be set to 30.  Note:
    // This object cannot be set to the value of 0. The default value of 0
    // signifies the default TTL values will be used for 'echo' and 'pathEcho'.
    // The type is interface{} with range: 0..255.
    RttMplsVpnMonTypeLSPTTL interface{}

    // This object specifies the DSCP value to be set in the IP header of the LSP
    // echo reply packet. The value of this object will be in range of DiffServ
    // codepoint values between 0 to 63.  Note: This object cannot be set to value
    // of 255. This default value specifies that DSCP is not set for this row. The
    // type is interface{} with range: 0..63 | 255..None.
    RttMplsVpnMonTypeLSPReplyDscp interface{}

    // This object represents the number of concurrent path discovery requests
    // that will be active at one time per MPLS VPN control row. This object is
    // meant for reducing the time for discovery of all the paths to the target in
    // a large customer network. However its value should be chosen such that it
    // does not cause any performance impact.  Note: If the customer network has
    // low end routers in the Core it is recommended to keep this value low. The
    // type is interface{} with range: 1..15.
    RttMplsVpnMonTypeLpdMaxSessions interface{}

    // This object specifies the maximum allowed duration of a particular tree
    // trace request.  If no response is received in configured time the request
    // will be considered a failure. The type is interface{} with range: 1..900.
    // Units are seconds.
    RttMplsVpnMonTypeLpdSessTimeout interface{}

    // This object specifies the timeout value for the LSP echo requests which are
    // sent while performing the LSP Path  Discovery. The type is interface{} with
    // range: 0..604800000. Units are milliseconds.
    RttMplsVpnMonTypeLpdEchoTimeout interface{}

    // This object specifies the send interval between LSP echo requests which are
    // sent while performing the LSP Path  Discovery. The type is interface{} with
    // range: 0..3600000. Units are milliseconds.
    RttMplsVpnMonTypeLpdEchoInterval interface{}

    // This object specifies if the explicit-null label is added to LSP echo
    // requests which are sent while performing the LSP Path Discovery.  If set to
    // TRUE all the probes configured as part of this control row will send the
    // LSP echo requests with the explicit-null label added. The type is bool.
    RttMplsVpnMonTypeLpdEchoNullShim interface{}

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
    RttMplsVpnMonTypeLpdScanPeriod interface{}

    // The maximum number of hours of data to be kept per LPD group. The LPD group
    // statistics will be kept in an hourly bucket. At the maximum there can be
    // two buckets. The value of 'one' is not advisable because the group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value used in the
    // rttMplsVpnLpdGroupStatsTable to uniquely identify this group is the
    // rttMonStatsCaptureStartTimeIndex.  Note: When this object is set to the
    // value of '0' all rttMplsVpnLpdGroupStatsTable data capturing will be shut
    // off. The type is interface{} with range: 0..2.
    RttMplsVpnMonTypeLpdStatHours interface{}

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
    RttMplsVpnMonScheduleRttStartTime interface{}

    // Specifies the time duration over which all the probes created by the
    // current Auto SAA L3 MPLS VPN have to be scheduled.  This object must be set
    // first before setting rttMplsVpnMonScheduleRttStartTime. The type is
    // interface{} with range: 1..604800. Units are seconds.
    RttMplsVpnMonSchedulePeriod interface{}

    // Specifies the duration between initiating each RTT operation configured by
    // the Auto SAA L3 MPLS VPN.  This object cannot be set to a value which would
    // be a shorter duration than rttMplsVpnMonCtrlTimeout.  The usage of this
    // value in RTT operation is same as rttMonCtrlAdminFrequency. The type is
    // interface{} with range: 1..604800. Units are seconds.
    RttMplsVpnMonScheduleFrequency interface{}

    // The value set for this will be applied as rttMonReactAdminConnectionEnable
    // for individual probes created by the Auto SAA L3 MPLS VPN.  When this
    // object is set to true, rttMonReactVar for individual probes created by the
    // Auto SAA L3 MPLS VPN will be set to 'connectionLoss(8)'. The type is bool.
    RttMplsVpnMonReactConnectionEnable interface{}

    // The value set for this will be applied as rttMonReactAdminTimeoutEnable for
    // individual probes created by the Auto SAA L3 MPLS VPN.  When this object is
    // set to true, rttMonReactVar for individual probes created by the Auto SAA
    // L3 MPLS VPN will be set to 'timeout(7)'. The type is bool.
    RttMplsVpnMonReactTimeoutEnable interface{}

    // The value corresponding to this object will be applied as
    // rttMonReactAdminThresholdType for individual probes created by the Auto SAA
    // L3 MPLS VPN.  The value corresponding to this object will be applied as
    // rttMonReactThresholdType for individual probes created by the Auto SAA L3
    // MPLS VPN. The type is RttMplsVpnMonReactThresholdType.
    RttMplsVpnMonReactThresholdType interface{}

    // This object value will be applied as rttMonReactAdminThresholdCount for
    // individual probes created by the Auto SAA L3 MPLS VPN.  This object value
    // will be applied as rttMonReactThresholdCountX for individual probes created
    // by the Auto SAA L3 MPLS VPN. The type is interface{} with range: 1..16.
    RttMplsVpnMonReactThresholdCount interface{}

    // The value corresponding to this object will be applied as
    // rttMonReactAdminActionType of individual probes created by this Auto SAA L3
    // MPLS VPN.  The value corresponding to this object will be applied as
    // rttMonReactActionType of individual probes created by this Auto SAA L3 MPLS
    // VPN. The type is RttMplsVpnMonReactActionType.
    RttMplsVpnMonReactActionType interface{}

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
    // type is RttMplsVpnMonReactLpdNotifyType.
    RttMplsVpnMonReactLpdNotifyType interface{}

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
    RttMplsVpnMonReactLpdRetryCount interface{}
}

func (rttMplsVpnMonCtrlEntry *CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry) GetEntityData() *types.CommonEntityData {
    rttMplsVpnMonCtrlEntry.EntityData.YFilter = rttMplsVpnMonCtrlEntry.YFilter
    rttMplsVpnMonCtrlEntry.EntityData.YangName = "rttMplsVpnMonCtrlEntry"
    rttMplsVpnMonCtrlEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMplsVpnMonCtrlEntry.EntityData.ParentYangName = "rttMplsVpnMonCtrlTable"
    rttMplsVpnMonCtrlEntry.EntityData.SegmentPath = "rttMplsVpnMonCtrlEntry" + types.AddKeyToken(rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlIndex, "rttMplsVpnMonCtrlIndex")
    rttMplsVpnMonCtrlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMplsVpnMonCtrlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMplsVpnMonCtrlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMplsVpnMonCtrlEntry.EntityData.Children = types.NewOrderedMap()
    rttMplsVpnMonCtrlEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlIndex", types.YLeaf{"RttMplsVpnMonCtrlIndex", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlIndex})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlRttType", types.YLeaf{"RttMplsVpnMonCtrlRttType", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlRttType})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlVrfName", types.YLeaf{"RttMplsVpnMonCtrlVrfName", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlVrfName})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlTag", types.YLeaf{"RttMplsVpnMonCtrlTag", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlTag})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlThreshold", types.YLeaf{"RttMplsVpnMonCtrlThreshold", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlThreshold})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlTimeout", types.YLeaf{"RttMplsVpnMonCtrlTimeout", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlTimeout})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlScanInterval", types.YLeaf{"RttMplsVpnMonCtrlScanInterval", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlScanInterval})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlDelScanFactor", types.YLeaf{"RttMplsVpnMonCtrlDelScanFactor", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlDelScanFactor})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlEXP", types.YLeaf{"RttMplsVpnMonCtrlEXP", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlEXP})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlRequestSize", types.YLeaf{"RttMplsVpnMonCtrlRequestSize", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlRequestSize})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlVerifyData", types.YLeaf{"RttMplsVpnMonCtrlVerifyData", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlVerifyData})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlStorageType", types.YLeaf{"RttMplsVpnMonCtrlStorageType", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlStorageType})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlProbeList", types.YLeaf{"RttMplsVpnMonCtrlProbeList", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlProbeList})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlStatus", types.YLeaf{"RttMplsVpnMonCtrlStatus", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlStatus})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlLpd", types.YLeaf{"RttMplsVpnMonCtrlLpd", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlLpd})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlLpdGrpList", types.YLeaf{"RttMplsVpnMonCtrlLpdGrpList", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlLpdGrpList})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonCtrlLpdCompTime", types.YLeaf{"RttMplsVpnMonCtrlLpdCompTime", rttMplsVpnMonCtrlEntry.RttMplsVpnMonCtrlLpdCompTime})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeInterval", types.YLeaf{"RttMplsVpnMonTypeInterval", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeInterval})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeNumPackets", types.YLeaf{"RttMplsVpnMonTypeNumPackets", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeNumPackets})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeDestPort", types.YLeaf{"RttMplsVpnMonTypeDestPort", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeDestPort})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeSecFreqType", types.YLeaf{"RttMplsVpnMonTypeSecFreqType", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeSecFreqType})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeSecFreqValue", types.YLeaf{"RttMplsVpnMonTypeSecFreqValue", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeSecFreqValue})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLspSelector", types.YLeaf{"RttMplsVpnMonTypeLspSelector", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLspSelector})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLSPReplyMode", types.YLeaf{"RttMplsVpnMonTypeLSPReplyMode", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLSPReplyMode})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLSPTTL", types.YLeaf{"RttMplsVpnMonTypeLSPTTL", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLSPTTL})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLSPReplyDscp", types.YLeaf{"RttMplsVpnMonTypeLSPReplyDscp", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLSPReplyDscp})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLpdMaxSessions", types.YLeaf{"RttMplsVpnMonTypeLpdMaxSessions", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLpdMaxSessions})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLpdSessTimeout", types.YLeaf{"RttMplsVpnMonTypeLpdSessTimeout", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLpdSessTimeout})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLpdEchoTimeout", types.YLeaf{"RttMplsVpnMonTypeLpdEchoTimeout", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLpdEchoTimeout})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLpdEchoInterval", types.YLeaf{"RttMplsVpnMonTypeLpdEchoInterval", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLpdEchoInterval})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLpdEchoNullShim", types.YLeaf{"RttMplsVpnMonTypeLpdEchoNullShim", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLpdEchoNullShim})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLpdScanPeriod", types.YLeaf{"RttMplsVpnMonTypeLpdScanPeriod", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLpdScanPeriod})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonTypeLpdStatHours", types.YLeaf{"RttMplsVpnMonTypeLpdStatHours", rttMplsVpnMonCtrlEntry.RttMplsVpnMonTypeLpdStatHours})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonScheduleRttStartTime", types.YLeaf{"RttMplsVpnMonScheduleRttStartTime", rttMplsVpnMonCtrlEntry.RttMplsVpnMonScheduleRttStartTime})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonSchedulePeriod", types.YLeaf{"RttMplsVpnMonSchedulePeriod", rttMplsVpnMonCtrlEntry.RttMplsVpnMonSchedulePeriod})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonScheduleFrequency", types.YLeaf{"RttMplsVpnMonScheduleFrequency", rttMplsVpnMonCtrlEntry.RttMplsVpnMonScheduleFrequency})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonReactConnectionEnable", types.YLeaf{"RttMplsVpnMonReactConnectionEnable", rttMplsVpnMonCtrlEntry.RttMplsVpnMonReactConnectionEnable})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonReactTimeoutEnable", types.YLeaf{"RttMplsVpnMonReactTimeoutEnable", rttMplsVpnMonCtrlEntry.RttMplsVpnMonReactTimeoutEnable})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonReactThresholdType", types.YLeaf{"RttMplsVpnMonReactThresholdType", rttMplsVpnMonCtrlEntry.RttMplsVpnMonReactThresholdType})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonReactThresholdCount", types.YLeaf{"RttMplsVpnMonReactThresholdCount", rttMplsVpnMonCtrlEntry.RttMplsVpnMonReactThresholdCount})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonReactActionType", types.YLeaf{"RttMplsVpnMonReactActionType", rttMplsVpnMonCtrlEntry.RttMplsVpnMonReactActionType})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonReactLpdNotifyType", types.YLeaf{"RttMplsVpnMonReactLpdNotifyType", rttMplsVpnMonCtrlEntry.RttMplsVpnMonReactLpdNotifyType})
    rttMplsVpnMonCtrlEntry.EntityData.Leafs.Append("rttMplsVpnMonReactLpdRetryCount", types.YLeaf{"RttMplsVpnMonReactLpdRetryCount", rttMplsVpnMonCtrlEntry.RttMplsVpnMonReactLpdRetryCount})

    rttMplsVpnMonCtrlEntry.EntityData.YListKeys = []string {"RttMplsVpnMonCtrlIndex"}

    return &(rttMplsVpnMonCtrlEntry.EntityData)
}

// CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactActionType represents this Auto SAA L3 MPLS VPN.
type CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactActionType string

const (
    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactActionType_none CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactActionType = "none"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactActionType_trapOnly CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactActionType = "trapOnly"
)

// CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType represents depending on the failure conditions.
type CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType string

const (
    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType_none CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType = "none"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType_lpdPathDiscovery CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType = "lpdPathDiscovery"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType_lpdGroupStatus CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType = "lpdGroupStatus"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType_lpdAll CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactLpdNotifyType = "lpdAll"
)

// CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType represents the Auto SAA L3 MPLS VPN.
type CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType string

const (
    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType_never CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType = "never"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType_immediate CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType = "immediate"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType_consecutive CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonReactThresholdType = "consecutive"
)

// CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType represents original frequency once the trap is sent.
type CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType string

const (
    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType_none CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType = "none"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType_timeout CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType = "timeout"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType_connectionLoss CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType = "connectionLoss"

    CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType_both CISCORTTMONMIB_RttMplsVpnMonCtrlTable_RttMplsVpnMonCtrlEntry_RttMplsVpnMonTypeSecFreqType = "both"
)

// CISCORTTMONMIB_RttMonReactTable
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
type CISCORTTMONMIB_RttMonReactTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual reaction configuration
    // control row. The type is slice of
    // CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry.
    RttMonReactEntry []*CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry
}

func (rttMonReactTable *CISCORTTMONMIB_RttMonReactTable) GetEntityData() *types.CommonEntityData {
    rttMonReactTable.EntityData.YFilter = rttMonReactTable.YFilter
    rttMonReactTable.EntityData.YangName = "rttMonReactTable"
    rttMonReactTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonReactTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonReactTable.EntityData.SegmentPath = "rttMonReactTable"
    rttMonReactTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonReactTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonReactTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonReactTable.EntityData.Children = types.NewOrderedMap()
    rttMonReactTable.EntityData.Children.Append("rttMonReactEntry", types.YChild{"RttMonReactEntry", nil})
    for i := range rttMonReactTable.RttMonReactEntry {
        rttMonReactTable.EntityData.Children.Append(types.GetSegmentPath(rttMonReactTable.RttMonReactEntry[i]), types.YChild{"RttMonReactEntry", rttMonReactTable.RttMonReactEntry[i]})
    }
    rttMonReactTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonReactTable.EntityData.YListKeys = []string {}

    return &(rttMonReactTable.EntityData)
}

// CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry
// A base list of objects that define a conceptual reaction
// configuration control row.
type CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. This object along with rttMonCtrlAdminIndex
    // identifies a particular reaction-configuration for a particular probe. This
    // is a pseudo-random number selected by the management station when creating
    // a row via the rttMonReactStatus. If the pseudo-random number is already in
    // use an 'inconsistentValue' return code will be returned when set operation
    // is attempted. The type is interface{} with range: 1..2147483647.
    RttMonReactConfigIndex interface{}

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
    RttMonReactVar interface{}

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
    // true. The type is RttMonReactThresholdType.
    RttMonReactThresholdType interface{}

    // Specifies what type(s), if any, of reaction(s) to generate if an operation
    // violates one of the watched ( reaction-configuration ) conditions:  none   
    // - no reaction is generated trapOnly           - a trap is generated
    // triggerOnly        - all trigger actions defined for this                  
    // entry are initiated trapAndTrigger     - both a trap and all trigger
    // actions                      are initiated A trigger action is defined via
    // the rttMonReactTriggerAdminTable. The type is RttMonReactActionType.
    RttMonReactActionType interface{}

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
    RttMonReactThresholdRising interface{}

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
    RttMonReactThresholdFalling interface{}

    // If rttMonReactThresholdType value is 'xOfy', this object defines the 'x'
    // value.  If rttMonReactThresholdType value is 'consecutive' this object
    // defines the number of consecutive occurrences that needs threshold
    // violation before setting  rttMonReactOccurred as true.  If
    // rttMonReactThresholdType value is 'average' this object defines the number
    // of samples that needs be considered for calculating average.  This object
    // has no meaning if rttMonReactThresholdType has value of 'never' and
    // 'immediate'. The type is interface{} with range: 1..16.
    RttMonReactThresholdCountX interface{}

    // This object defines the 'y' value of the xOfy condition if
    // rttMonReactThresholdType is 'xOfy'.  For other values of
    // rttMonReactThresholdType, this object is not applicable. The type is
    // interface{} with range: 1..16.
    RttMonReactThresholdCountY interface{}

    // This object will be set when the configured threshold condition is violated
    // as defined by rttMonReactThresholdType and holds the actual value that
    // violated the configured threshold values.  This object is not valid for the
    // following values of rttMonReactVar and It will be always 0:   - timeout   -
    // connectionLoss   - verifyError. The type is interface{} with range:
    // 0..2147483647.
    RttMonReactValue interface{}

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
    RttMonReactOccurred interface{}

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
    RttMonReactStatus interface{}
}

func (rttMonReactEntry *CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry) GetEntityData() *types.CommonEntityData {
    rttMonReactEntry.EntityData.YFilter = rttMonReactEntry.YFilter
    rttMonReactEntry.EntityData.YangName = "rttMonReactEntry"
    rttMonReactEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonReactEntry.EntityData.ParentYangName = "rttMonReactTable"
    rttMonReactEntry.EntityData.SegmentPath = "rttMonReactEntry" + types.AddKeyToken(rttMonReactEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonReactEntry.RttMonReactConfigIndex, "rttMonReactConfigIndex")
    rttMonReactEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonReactEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonReactEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonReactEntry.EntityData.Children = types.NewOrderedMap()
    rttMonReactEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonReactEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonReactEntry.RttMonCtrlAdminIndex})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactConfigIndex", types.YLeaf{"RttMonReactConfigIndex", rttMonReactEntry.RttMonReactConfigIndex})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactVar", types.YLeaf{"RttMonReactVar", rttMonReactEntry.RttMonReactVar})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactThresholdType", types.YLeaf{"RttMonReactThresholdType", rttMonReactEntry.RttMonReactThresholdType})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactActionType", types.YLeaf{"RttMonReactActionType", rttMonReactEntry.RttMonReactActionType})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactThresholdRising", types.YLeaf{"RttMonReactThresholdRising", rttMonReactEntry.RttMonReactThresholdRising})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactThresholdFalling", types.YLeaf{"RttMonReactThresholdFalling", rttMonReactEntry.RttMonReactThresholdFalling})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactThresholdCountX", types.YLeaf{"RttMonReactThresholdCountX", rttMonReactEntry.RttMonReactThresholdCountX})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactThresholdCountY", types.YLeaf{"RttMonReactThresholdCountY", rttMonReactEntry.RttMonReactThresholdCountY})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactValue", types.YLeaf{"RttMonReactValue", rttMonReactEntry.RttMonReactValue})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactOccurred", types.YLeaf{"RttMonReactOccurred", rttMonReactEntry.RttMonReactOccurred})
    rttMonReactEntry.EntityData.Leafs.Append("rttMonReactStatus", types.YLeaf{"RttMonReactStatus", rttMonReactEntry.RttMonReactStatus})

    rttMonReactEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonReactConfigIndex"}

    return &(rttMonReactEntry.EntityData)
}

// CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType represents rttMonReactTriggerAdminTable.
type CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType string

const (
    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType_none CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType = "none"

    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType_trapOnly CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType = "trapOnly"

    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType_triggerOnly CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType = "triggerOnly"

    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType_trapAndTrigger CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactActionType = "trapAndTrigger"
)

// CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType represents rttMonReactOccurred was true.
type CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType string

const (
    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType_never CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType = "never"

    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType_immediate CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType = "immediate"

    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType_consecutive CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType = "consecutive"

    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType_xOfy CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType = "xOfy"

    CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType_average CISCORTTMONMIB_RttMonReactTable_RttMonReactEntry_RttMonReactThresholdType = "average"
)

// CISCORTTMONMIB_RttMonGeneratedOperTable
// This table contains information about the generated
// operation id as part of a parent IP SLA operation. The parent
// operation id is pseudo-random number, selected by the management 
// station based on an operation started by the management 
// station,when creating a row via the rttMonCtrlAdminStatus
// object in the rttMonCtrlAdminTable table.
type CISCORTTMONMIB_RttMonGeneratedOperTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Generated Oper table corresponding to a child or generated
    // operation as part of a parent IP SLA operation. The type is slice of
    // CISCORTTMONMIB_RttMonGeneratedOperTable_RttMonGeneratedOperEntry.
    RttMonGeneratedOperEntry []*CISCORTTMONMIB_RttMonGeneratedOperTable_RttMonGeneratedOperEntry
}

func (rttMonGeneratedOperTable *CISCORTTMONMIB_RttMonGeneratedOperTable) GetEntityData() *types.CommonEntityData {
    rttMonGeneratedOperTable.EntityData.YFilter = rttMonGeneratedOperTable.YFilter
    rttMonGeneratedOperTable.EntityData.YangName = "rttMonGeneratedOperTable"
    rttMonGeneratedOperTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonGeneratedOperTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonGeneratedOperTable.EntityData.SegmentPath = "rttMonGeneratedOperTable"
    rttMonGeneratedOperTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonGeneratedOperTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonGeneratedOperTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonGeneratedOperTable.EntityData.Children = types.NewOrderedMap()
    rttMonGeneratedOperTable.EntityData.Children.Append("rttMonGeneratedOperEntry", types.YChild{"RttMonGeneratedOperEntry", nil})
    for i := range rttMonGeneratedOperTable.RttMonGeneratedOperEntry {
        rttMonGeneratedOperTable.EntityData.Children.Append(types.GetSegmentPath(rttMonGeneratedOperTable.RttMonGeneratedOperEntry[i]), types.YChild{"RttMonGeneratedOperEntry", rttMonGeneratedOperTable.RttMonGeneratedOperEntry[i]})
    }
    rttMonGeneratedOperTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonGeneratedOperTable.EntityData.YListKeys = []string {}

    return &(rttMonGeneratedOperTable.EntityData)
}

// CISCORTTMONMIB_RttMonGeneratedOperTable_RttMonGeneratedOperEntry
// An entry in the Generated Oper table corresponding to
// a child or generated operation as part of a parent
// IP SLA operation.
type CISCORTTMONMIB_RttMonGeneratedOperTable_RttMonGeneratedOperEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. The type of Internet address, IPv4 or IPv6, of a
    // responder for an IP SLA operation. The type is InetAddressType.
    RttMonGeneratedOperRespIpAddrType interface{}

    // This attribute is a key. The internet address of a responder for IP SLA
    // operation. The type of this address is determined by the value of
    // rttMonGeneratedOperRespIpAddrType. The type is string with length: 0..255.
    RttMonGeneratedOperRespIpAddr interface{}

    // This is a pseudo-random number, auto-generated based to identify a child
    // operation based on a parent  operation started by the management
    // station,when  creating a row via the rttMonCtrlAdminStatus object. The type
    // is interface{} with range: 1..2147483647.
    RttMonGeneratedOperCtrlAdminIndex interface{}
}

func (rttMonGeneratedOperEntry *CISCORTTMONMIB_RttMonGeneratedOperTable_RttMonGeneratedOperEntry) GetEntityData() *types.CommonEntityData {
    rttMonGeneratedOperEntry.EntityData.YFilter = rttMonGeneratedOperEntry.YFilter
    rttMonGeneratedOperEntry.EntityData.YangName = "rttMonGeneratedOperEntry"
    rttMonGeneratedOperEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonGeneratedOperEntry.EntityData.ParentYangName = "rttMonGeneratedOperTable"
    rttMonGeneratedOperEntry.EntityData.SegmentPath = "rttMonGeneratedOperEntry" + types.AddKeyToken(rttMonGeneratedOperEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonGeneratedOperEntry.RttMonGeneratedOperRespIpAddrType, "rttMonGeneratedOperRespIpAddrType") + types.AddKeyToken(rttMonGeneratedOperEntry.RttMonGeneratedOperRespIpAddr, "rttMonGeneratedOperRespIpAddr")
    rttMonGeneratedOperEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonGeneratedOperEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonGeneratedOperEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonGeneratedOperEntry.EntityData.Children = types.NewOrderedMap()
    rttMonGeneratedOperEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonGeneratedOperEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonGeneratedOperEntry.RttMonCtrlAdminIndex})
    rttMonGeneratedOperEntry.EntityData.Leafs.Append("rttMonGeneratedOperRespIpAddrType", types.YLeaf{"RttMonGeneratedOperRespIpAddrType", rttMonGeneratedOperEntry.RttMonGeneratedOperRespIpAddrType})
    rttMonGeneratedOperEntry.EntityData.Leafs.Append("rttMonGeneratedOperRespIpAddr", types.YLeaf{"RttMonGeneratedOperRespIpAddr", rttMonGeneratedOperEntry.RttMonGeneratedOperRespIpAddr})
    rttMonGeneratedOperEntry.EntityData.Leafs.Append("rttMonGeneratedOperCtrlAdminIndex", types.YLeaf{"RttMonGeneratedOperCtrlAdminIndex", rttMonGeneratedOperEntry.RttMonGeneratedOperCtrlAdminIndex})

    rttMonGeneratedOperEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonGeneratedOperRespIpAddrType", "RttMonGeneratedOperRespIpAddr"}

    return &(rttMonGeneratedOperEntry.EntityData)
}

// CISCORTTMONMIB_RttMonStatsCaptureTable
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
type CISCORTTMONMIB_RttMonStatsCaptureTable struct {
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
    // of CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry.
    RttMonStatsCaptureEntry []*CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry
}

func (rttMonStatsCaptureTable *CISCORTTMONMIB_RttMonStatsCaptureTable) GetEntityData() *types.CommonEntityData {
    rttMonStatsCaptureTable.EntityData.YFilter = rttMonStatsCaptureTable.YFilter
    rttMonStatsCaptureTable.EntityData.YangName = "rttMonStatsCaptureTable"
    rttMonStatsCaptureTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonStatsCaptureTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonStatsCaptureTable.EntityData.SegmentPath = "rttMonStatsCaptureTable"
    rttMonStatsCaptureTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonStatsCaptureTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonStatsCaptureTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonStatsCaptureTable.EntityData.Children = types.NewOrderedMap()
    rttMonStatsCaptureTable.EntityData.Children.Append("rttMonStatsCaptureEntry", types.YChild{"RttMonStatsCaptureEntry", nil})
    for i := range rttMonStatsCaptureTable.RttMonStatsCaptureEntry {
        rttMonStatsCaptureTable.EntityData.Children.Append(types.GetSegmentPath(rttMonStatsCaptureTable.RttMonStatsCaptureEntry[i]), types.YChild{"RttMonStatsCaptureEntry", rttMonStatsCaptureTable.RttMonStatsCaptureEntry[i]})
    }
    rttMonStatsCaptureTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonStatsCaptureTable.EntityData.YListKeys = []string {}

    return &(rttMonStatsCaptureTable.EntityData)
}

// CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry
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
type CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. The time when this row was created.  This object
    // is the second index of the  rttMonStatsCaptureTable Table.  The the number
    // of rttMonStatsCaptureStartTimeIndex   groups exceeds the
    // rttMonStatisticsAdminNumHourGroups value, the oldest
    // rttMonStatsCaptureStartTimeIndex  group will be removed and replaced with
    // the new entry.  When the RttMonRttType is 'pathEcho', this object also 
    // uniquely defines a group of paths.  See the  rttMonStatsCaptureEntry
    // object. The type is interface{} with range: 0..4294967295.
    RttMonStatsCaptureStartTimeIndex interface{}

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
    RttMonStatsCapturePathIndex interface{}

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
    RttMonStatsCaptureHopIndex interface{}

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
    RttMonStatsCaptureDistIndex interface{}

    // The number of RTT operations that have completed without an error and
    // without timing out.  This object has the special behavior as defined by the
    // ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB object. The type is
    // interface{} with range: 0..2147483647.
    RttMonStatsCaptureCompletions interface{}

    // The number of RTT operations successfully completed, but in excess of
    // rttMonCtrlAdminThreshold.  This number is a subset of the accumulation of
    // all  rttMonStatsCaptureCompletions.  The operation time  of these completed
    // operations will be accumulated.  This object has the special behavior as
    // defined by the ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB
    // object. The type is interface{} with range: 0..2147483647.
    RttMonStatsCaptureOverThresholds interface{}

    // The accumulated completion time of RTT operations which complete
    // successfully. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    RttMonStatsCaptureSumCompletionTime interface{}

    // The low order 32 bits of the accumulated squares of completion times (in
    // milliseconds) of RTT  operations which complete successfully.  Low/High
    // order is defined where the binary number will look as follows:
    // ------------------------------------------------- | High order 32 bits    |
    // Low order 32 bits     | -------------------------------------------------
    // For example the number 4294967296 would have all Low order bits as '0' and
    // the rightmost High order bit will be 1 (zeros,1). The type is interface{}
    // with range: 0..4294967295.
    RttMonStatsCaptureSumCompletionTime2Low interface{}

    // The high order 32 bits of the accumulated squares of completion times (in
    // milliseconds) of RTT  operations which complete successfully.  See the
    // rttMonStatsCaptureSumCompletionTime2Low object for a definition of Low/High
    // Order. The type is interface{} with range: 0..4294967295.
    RttMonStatsCaptureSumCompletionTime2High interface{}

    // The maximum completion time of any RTT operation which completes
    // successfully. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    RttMonStatsCaptureCompletionTimeMax interface{}

    // The minimum completion time of any RTT operation which completes
    // successfully. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    RttMonStatsCaptureCompletionTimeMin interface{}
}

func (rttMonStatsCaptureEntry *CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry) GetEntityData() *types.CommonEntityData {
    rttMonStatsCaptureEntry.EntityData.YFilter = rttMonStatsCaptureEntry.YFilter
    rttMonStatsCaptureEntry.EntityData.YangName = "rttMonStatsCaptureEntry"
    rttMonStatsCaptureEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonStatsCaptureEntry.EntityData.ParentYangName = "rttMonStatsCaptureTable"
    rttMonStatsCaptureEntry.EntityData.SegmentPath = "rttMonStatsCaptureEntry" + types.AddKeyToken(rttMonStatsCaptureEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonStatsCaptureEntry.RttMonStatsCaptureStartTimeIndex, "rttMonStatsCaptureStartTimeIndex") + types.AddKeyToken(rttMonStatsCaptureEntry.RttMonStatsCapturePathIndex, "rttMonStatsCapturePathIndex") + types.AddKeyToken(rttMonStatsCaptureEntry.RttMonStatsCaptureHopIndex, "rttMonStatsCaptureHopIndex") + types.AddKeyToken(rttMonStatsCaptureEntry.RttMonStatsCaptureDistIndex, "rttMonStatsCaptureDistIndex")
    rttMonStatsCaptureEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonStatsCaptureEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonStatsCaptureEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonStatsCaptureEntry.EntityData.Children = types.NewOrderedMap()
    rttMonStatsCaptureEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonStatsCaptureEntry.RttMonCtrlAdminIndex})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureStartTimeIndex", types.YLeaf{"RttMonStatsCaptureStartTimeIndex", rttMonStatsCaptureEntry.RttMonStatsCaptureStartTimeIndex})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCapturePathIndex", types.YLeaf{"RttMonStatsCapturePathIndex", rttMonStatsCaptureEntry.RttMonStatsCapturePathIndex})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureHopIndex", types.YLeaf{"RttMonStatsCaptureHopIndex", rttMonStatsCaptureEntry.RttMonStatsCaptureHopIndex})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureDistIndex", types.YLeaf{"RttMonStatsCaptureDistIndex", rttMonStatsCaptureEntry.RttMonStatsCaptureDistIndex})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureCompletions", types.YLeaf{"RttMonStatsCaptureCompletions", rttMonStatsCaptureEntry.RttMonStatsCaptureCompletions})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureOverThresholds", types.YLeaf{"RttMonStatsCaptureOverThresholds", rttMonStatsCaptureEntry.RttMonStatsCaptureOverThresholds})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureSumCompletionTime", types.YLeaf{"RttMonStatsCaptureSumCompletionTime", rttMonStatsCaptureEntry.RttMonStatsCaptureSumCompletionTime})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureSumCompletionTime2Low", types.YLeaf{"RttMonStatsCaptureSumCompletionTime2Low", rttMonStatsCaptureEntry.RttMonStatsCaptureSumCompletionTime2Low})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureSumCompletionTime2High", types.YLeaf{"RttMonStatsCaptureSumCompletionTime2High", rttMonStatsCaptureEntry.RttMonStatsCaptureSumCompletionTime2High})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureCompletionTimeMax", types.YLeaf{"RttMonStatsCaptureCompletionTimeMax", rttMonStatsCaptureEntry.RttMonStatsCaptureCompletionTimeMax})
    rttMonStatsCaptureEntry.EntityData.Leafs.Append("rttMonStatsCaptureCompletionTimeMin", types.YLeaf{"RttMonStatsCaptureCompletionTimeMin", rttMonStatsCaptureEntry.RttMonStatsCaptureCompletionTimeMin})

    rttMonStatsCaptureEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonStatsCaptureStartTimeIndex", "RttMonStatsCapturePathIndex", "RttMonStatsCaptureHopIndex", "RttMonStatsCaptureDistIndex"}

    return &(rttMonStatsCaptureEntry.EntityData)
}

// CISCORTTMONMIB_RttMonStatsCollectTable
// The statistics collection database.
// 
// This table has the exact same behavior as the
// rttMonStatsCaptureTable, except it does not keep
// statistical distribution information.
// 
// For a complete table description see
// the rttMonStatsCaptureTable object.
type CISCORTTMONMIB_RttMonStatsCollectTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry has the exact same
    // behavior as the  rttMonStatsCaptureEntry, except it does not keep
    // statistical distribution information.  For a complete entry description see
    // the rttMonStatsCaptureEntry object. The type is slice of
    // CISCORTTMONMIB_RttMonStatsCollectTable_RttMonStatsCollectEntry.
    RttMonStatsCollectEntry []*CISCORTTMONMIB_RttMonStatsCollectTable_RttMonStatsCollectEntry
}

func (rttMonStatsCollectTable *CISCORTTMONMIB_RttMonStatsCollectTable) GetEntityData() *types.CommonEntityData {
    rttMonStatsCollectTable.EntityData.YFilter = rttMonStatsCollectTable.YFilter
    rttMonStatsCollectTable.EntityData.YangName = "rttMonStatsCollectTable"
    rttMonStatsCollectTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonStatsCollectTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonStatsCollectTable.EntityData.SegmentPath = "rttMonStatsCollectTable"
    rttMonStatsCollectTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonStatsCollectTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonStatsCollectTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonStatsCollectTable.EntityData.Children = types.NewOrderedMap()
    rttMonStatsCollectTable.EntityData.Children.Append("rttMonStatsCollectEntry", types.YChild{"RttMonStatsCollectEntry", nil})
    for i := range rttMonStatsCollectTable.RttMonStatsCollectEntry {
        rttMonStatsCollectTable.EntityData.Children.Append(types.GetSegmentPath(rttMonStatsCollectTable.RttMonStatsCollectEntry[i]), types.YChild{"RttMonStatsCollectEntry", rttMonStatsCollectTable.RttMonStatsCollectEntry[i]})
    }
    rttMonStatsCollectTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonStatsCollectTable.EntityData.YListKeys = []string {}

    return &(rttMonStatsCollectTable.EntityData)
}

// CISCORTTMONMIB_RttMonStatsCollectTable_RttMonStatsCollectEntry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry has the exact same behavior as the 
// rttMonStatsCaptureEntry, except it does not keep
// statistical distribution information.
// 
// For a complete entry description see
// the rttMonStatsCaptureEntry object.
type CISCORTTMONMIB_RttMonStatsCollectTable_RttMonStatsCollectEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry_RttMonStatsCaptureStartTimeIndex
    RttMonStatsCaptureStartTimeIndex interface{}

    // This attribute is a key. The type is string with range: 1..128. Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry_RttMonStatsCapturePathIndex
    RttMonStatsCapturePathIndex interface{}

    // This attribute is a key. The type is string with range: 1..30. Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry_RttMonStatsCaptureHopIndex
    RttMonStatsCaptureHopIndex interface{}

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
    RttMonStatsCollectNumDisconnects interface{}

    // The number of occasions when a RTT operation was not completed before a
    // timeout occurred, i.e. rttMonCtrlAdminTimeout was exceeded.  Since the RTT
    // operation was never completed, the  completion time of these operations are
    // not accumulated, nor do they increment rttMonStatsCaptureCompletions (in 
    // any of the statistics distribution buckets).  This object has the special
    // behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    RttMonStatsCollectTimeouts interface{}

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
    RttMonStatsCollectBusies interface{}

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
    RttMonStatsCollectNoConnections interface{}

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
    RttMonStatsCollectDrops interface{}

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
    RttMonStatsCollectSequenceErrors interface{}

    // The number of RTT operation completions received with data that does not
    // compare with the expected data.  The  completion time of these operations
    // are not accumulated,  nor do they increment rttMonStatsCaptureCompletions
    // (in the expected Distribution Bucket).  This object has the special
    // behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    RttMonStatsCollectVerifyErrors interface{}

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
    RttMonStatsCollectAddress interface{}

    // The number of occasions when control enable request failed. Currently it is
    // used for multicast operation type.  This object has the special behavior as
    // defined by the ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB
    // object. rttMonControlEnableErrors object is superseded by
    // rttMonStatsCollectCtrlEnErrors. The type is interface{} with range:
    // 0..2147483647.
    RttMonControlEnableErrors interface{}

    // The number of occasions when stats retrieval request failed. Currently it
    // is used for multicast operation type.  This object has the special behavior
    // as defined by the ROLLOVER NOTE in the DESCRIPTION of the ciscoRttMonMIB
    // object. rttMonStatsRetrieveErrors object is superseded by
    // rttMonStatsCollectRetrieveErrors. The type is interface{} with range:
    // 0..2147483647.
    RttMonStatsRetrieveErrors interface{}

    // The object is same as rttMonControlEnableErrors, with corrected name for
    // consistency.  The number of occasions when control enable request failed.
    // Currently it is used for multicast operation type.  This object has the
    // special behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    RttMonStatsCollectCtrlEnErrors interface{}

    // The object is same as rttMonStatsRetrieveErrors, with corrected name for
    // consistency.  The number of occasions when stats retrieval request failed.
    // Currently it is used for multicast operation type.  This object has the
    // special behavior as defined by the ROLLOVER NOTE in the DESCRIPTION of the
    // ciscoRttMonMIB object. The type is interface{} with range: 0..2147483647.
    RttMonStatsCollectRetrieveErrors interface{}
}

func (rttMonStatsCollectEntry *CISCORTTMONMIB_RttMonStatsCollectTable_RttMonStatsCollectEntry) GetEntityData() *types.CommonEntityData {
    rttMonStatsCollectEntry.EntityData.YFilter = rttMonStatsCollectEntry.YFilter
    rttMonStatsCollectEntry.EntityData.YangName = "rttMonStatsCollectEntry"
    rttMonStatsCollectEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonStatsCollectEntry.EntityData.ParentYangName = "rttMonStatsCollectTable"
    rttMonStatsCollectEntry.EntityData.SegmentPath = "rttMonStatsCollectEntry" + types.AddKeyToken(rttMonStatsCollectEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonStatsCollectEntry.RttMonStatsCaptureStartTimeIndex, "rttMonStatsCaptureStartTimeIndex") + types.AddKeyToken(rttMonStatsCollectEntry.RttMonStatsCapturePathIndex, "rttMonStatsCapturePathIndex") + types.AddKeyToken(rttMonStatsCollectEntry.RttMonStatsCaptureHopIndex, "rttMonStatsCaptureHopIndex")
    rttMonStatsCollectEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonStatsCollectEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonStatsCollectEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonStatsCollectEntry.EntityData.Children = types.NewOrderedMap()
    rttMonStatsCollectEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonStatsCollectEntry.RttMonCtrlAdminIndex})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCaptureStartTimeIndex", types.YLeaf{"RttMonStatsCaptureStartTimeIndex", rttMonStatsCollectEntry.RttMonStatsCaptureStartTimeIndex})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCapturePathIndex", types.YLeaf{"RttMonStatsCapturePathIndex", rttMonStatsCollectEntry.RttMonStatsCapturePathIndex})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCaptureHopIndex", types.YLeaf{"RttMonStatsCaptureHopIndex", rttMonStatsCollectEntry.RttMonStatsCaptureHopIndex})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectNumDisconnects", types.YLeaf{"RttMonStatsCollectNumDisconnects", rttMonStatsCollectEntry.RttMonStatsCollectNumDisconnects})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectTimeouts", types.YLeaf{"RttMonStatsCollectTimeouts", rttMonStatsCollectEntry.RttMonStatsCollectTimeouts})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectBusies", types.YLeaf{"RttMonStatsCollectBusies", rttMonStatsCollectEntry.RttMonStatsCollectBusies})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectNoConnections", types.YLeaf{"RttMonStatsCollectNoConnections", rttMonStatsCollectEntry.RttMonStatsCollectNoConnections})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectDrops", types.YLeaf{"RttMonStatsCollectDrops", rttMonStatsCollectEntry.RttMonStatsCollectDrops})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectSequenceErrors", types.YLeaf{"RttMonStatsCollectSequenceErrors", rttMonStatsCollectEntry.RttMonStatsCollectSequenceErrors})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectVerifyErrors", types.YLeaf{"RttMonStatsCollectVerifyErrors", rttMonStatsCollectEntry.RttMonStatsCollectVerifyErrors})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectAddress", types.YLeaf{"RttMonStatsCollectAddress", rttMonStatsCollectEntry.RttMonStatsCollectAddress})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonControlEnableErrors", types.YLeaf{"RttMonControlEnableErrors", rttMonStatsCollectEntry.RttMonControlEnableErrors})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsRetrieveErrors", types.YLeaf{"RttMonStatsRetrieveErrors", rttMonStatsCollectEntry.RttMonStatsRetrieveErrors})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectCtrlEnErrors", types.YLeaf{"RttMonStatsCollectCtrlEnErrors", rttMonStatsCollectEntry.RttMonStatsCollectCtrlEnErrors})
    rttMonStatsCollectEntry.EntityData.Leafs.Append("rttMonStatsCollectRetrieveErrors", types.YLeaf{"RttMonStatsCollectRetrieveErrors", rttMonStatsCollectEntry.RttMonStatsCollectRetrieveErrors})

    rttMonStatsCollectEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonStatsCaptureStartTimeIndex", "RttMonStatsCapturePathIndex", "RttMonStatsCaptureHopIndex"}

    return &(rttMonStatsCollectEntry.EntityData)
}

// CISCORTTMONMIB_RttMonStatsTotalsTable
// The statistics totals database.
// 
// This table has the exact same behavior as the
// rttMonStatsCaptureTable, except it only keeps
// 60 minute group values.
// 
// For a complete table description see
// the rttMonStatsCaptureTable object.
type CISCORTTMONMIB_RttMonStatsTotalsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry has the exact same
    // behavior as the  rttMonStatsCaptureEntry, except it only keeps 60 minute
    // group values.  For a complete entry description see the
    // rttMonStatsCaptureEntry object. The type is slice of
    // CISCORTTMONMIB_RttMonStatsTotalsTable_RttMonStatsTotalsEntry.
    RttMonStatsTotalsEntry []*CISCORTTMONMIB_RttMonStatsTotalsTable_RttMonStatsTotalsEntry
}

func (rttMonStatsTotalsTable *CISCORTTMONMIB_RttMonStatsTotalsTable) GetEntityData() *types.CommonEntityData {
    rttMonStatsTotalsTable.EntityData.YFilter = rttMonStatsTotalsTable.YFilter
    rttMonStatsTotalsTable.EntityData.YangName = "rttMonStatsTotalsTable"
    rttMonStatsTotalsTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonStatsTotalsTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonStatsTotalsTable.EntityData.SegmentPath = "rttMonStatsTotalsTable"
    rttMonStatsTotalsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonStatsTotalsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonStatsTotalsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonStatsTotalsTable.EntityData.Children = types.NewOrderedMap()
    rttMonStatsTotalsTable.EntityData.Children.Append("rttMonStatsTotalsEntry", types.YChild{"RttMonStatsTotalsEntry", nil})
    for i := range rttMonStatsTotalsTable.RttMonStatsTotalsEntry {
        rttMonStatsTotalsTable.EntityData.Children.Append(types.GetSegmentPath(rttMonStatsTotalsTable.RttMonStatsTotalsEntry[i]), types.YChild{"RttMonStatsTotalsEntry", rttMonStatsTotalsTable.RttMonStatsTotalsEntry[i]})
    }
    rttMonStatsTotalsTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonStatsTotalsTable.EntityData.YListKeys = []string {}

    return &(rttMonStatsTotalsTable.EntityData)
}

// CISCORTTMONMIB_RttMonStatsTotalsTable_RttMonStatsTotalsEntry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry has the exact same behavior as the 
// rttMonStatsCaptureEntry, except it only keeps
// 60 minute group values.
// 
// For a complete entry description see
// the rttMonStatsCaptureEntry object.
type CISCORTTMONMIB_RttMonStatsTotalsTable_RttMonStatsTotalsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonStatsCaptureTable_RttMonStatsCaptureEntry_RttMonStatsCaptureStartTimeIndex
    RttMonStatsCaptureStartTimeIndex interface{}

    // The length of time since this conceptual statistics row was created. The
    // type is interface{} with range: 0..2147483647.
    RttMonStatsTotalsElapsedTime interface{}

    // The number of RTT operations that have been initiated.  This number
    // includes all RTT operations which succeed  or fail for whatever reason. 
    // This object has the special behavior as defined by the ROLLOVER NOTE in the
    // DESCRIPTION of the ciscoRttMonMIB object. The type is interface{} with
    // range: 0..2147483647.
    RttMonStatsTotalsInitiations interface{}
}

func (rttMonStatsTotalsEntry *CISCORTTMONMIB_RttMonStatsTotalsTable_RttMonStatsTotalsEntry) GetEntityData() *types.CommonEntityData {
    rttMonStatsTotalsEntry.EntityData.YFilter = rttMonStatsTotalsEntry.YFilter
    rttMonStatsTotalsEntry.EntityData.YangName = "rttMonStatsTotalsEntry"
    rttMonStatsTotalsEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonStatsTotalsEntry.EntityData.ParentYangName = "rttMonStatsTotalsTable"
    rttMonStatsTotalsEntry.EntityData.SegmentPath = "rttMonStatsTotalsEntry" + types.AddKeyToken(rttMonStatsTotalsEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonStatsTotalsEntry.RttMonStatsCaptureStartTimeIndex, "rttMonStatsCaptureStartTimeIndex")
    rttMonStatsTotalsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonStatsTotalsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonStatsTotalsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonStatsTotalsEntry.EntityData.Children = types.NewOrderedMap()
    rttMonStatsTotalsEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonStatsTotalsEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonStatsTotalsEntry.RttMonCtrlAdminIndex})
    rttMonStatsTotalsEntry.EntityData.Leafs.Append("rttMonStatsCaptureStartTimeIndex", types.YLeaf{"RttMonStatsCaptureStartTimeIndex", rttMonStatsTotalsEntry.RttMonStatsCaptureStartTimeIndex})
    rttMonStatsTotalsEntry.EntityData.Leafs.Append("rttMonStatsTotalsElapsedTime", types.YLeaf{"RttMonStatsTotalsElapsedTime", rttMonStatsTotalsEntry.RttMonStatsTotalsElapsedTime})
    rttMonStatsTotalsEntry.EntityData.Leafs.Append("rttMonStatsTotalsInitiations", types.YLeaf{"RttMonStatsTotalsInitiations", rttMonStatsTotalsEntry.RttMonStatsTotalsInitiations})

    rttMonStatsTotalsEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonStatsCaptureStartTimeIndex"}

    return &(rttMonStatsTotalsEntry.EntityData)
}

// CISCORTTMONMIB_RttMonHTTPStatsTable
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
type CISCORTTMONMIB_RttMonHTTPStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry is created only if the
    // rttMonCtrlAdminRttType  is http. The operation of this table is same as
    // that of rttMonStatsCaptureTable. The type is slice of
    // CISCORTTMONMIB_RttMonHTTPStatsTable_RttMonHTTPStatsEntry.
    RttMonHTTPStatsEntry []*CISCORTTMONMIB_RttMonHTTPStatsTable_RttMonHTTPStatsEntry
}

func (rttMonHTTPStatsTable *CISCORTTMONMIB_RttMonHTTPStatsTable) GetEntityData() *types.CommonEntityData {
    rttMonHTTPStatsTable.EntityData.YFilter = rttMonHTTPStatsTable.YFilter
    rttMonHTTPStatsTable.EntityData.YangName = "rttMonHTTPStatsTable"
    rttMonHTTPStatsTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonHTTPStatsTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonHTTPStatsTable.EntityData.SegmentPath = "rttMonHTTPStatsTable"
    rttMonHTTPStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonHTTPStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonHTTPStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonHTTPStatsTable.EntityData.Children = types.NewOrderedMap()
    rttMonHTTPStatsTable.EntityData.Children.Append("rttMonHTTPStatsEntry", types.YChild{"RttMonHTTPStatsEntry", nil})
    for i := range rttMonHTTPStatsTable.RttMonHTTPStatsEntry {
        rttMonHTTPStatsTable.EntityData.Children.Append(types.GetSegmentPath(rttMonHTTPStatsTable.RttMonHTTPStatsEntry[i]), types.YChild{"RttMonHTTPStatsEntry", rttMonHTTPStatsTable.RttMonHTTPStatsEntry[i]})
    }
    rttMonHTTPStatsTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonHTTPStatsTable.EntityData.YListKeys = []string {}

    return &(rttMonHTTPStatsTable.EntityData)
}

// CISCORTTMONMIB_RttMonHTTPStatsTable_RttMonHTTPStatsEntry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry is created only if the rttMonCtrlAdminRttType 
// is http. The operation of this table is same as that of
// rttMonStatsCaptureTable.
type CISCORTTMONMIB_RttMonHTTPStatsTable_RttMonHTTPStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. This is the time when this row was created. This
    // index uniquely identifies a HTTP Stats row in the  rttMonHTTPStatsTable.
    // The type is interface{} with range: 0..4294967295.
    RttMonHTTPStatsStartTimeIndex interface{}

    // The number of HTTP operations that have completed successfully. The type is
    // interface{} with range: 0..4294967295.
    RttMonHTTPStatsCompletions interface{}

    // The number of HTTP operations that violate threshold. The type is
    // interface{} with range: 0..4294967295.
    RttMonHTTPStatsOverThresholds interface{}

    // The sum of HTTP operations that are successfully measured. The type is
    // interface{} with range: 0..4294967295.
    RttMonHTTPStatsRTTSum interface{}

    // The sum of squares of the RTT's that are successfully measured (low order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    RttMonHTTPStatsRTTSum2Low interface{}

    // The sum of squares of the RTT's that are successfully measured (high order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    RttMonHTTPStatsRTTSum2High interface{}

    // The minimum RTT taken to perform HTTP operation. The type is interface{}
    // with range: 0..4294967295.
    RttMonHTTPStatsRTTMin interface{}

    // The maximum RTT taken to perform HTTP operation. The type is interface{}
    // with range: 0..4294967295. Units are milliseconds.
    RttMonHTTPStatsRTTMax interface{}

    // The sum of RTT taken to perform DNS query within the HTTP operation. The
    // type is interface{} with range: 0..4294967295.
    RttMonHTTPStatsDNSRTTSum interface{}

    // The sum of RTT taken to connect to the HTTP server. The type is interface{}
    // with range: 0..4294967295.
    RttMonHTTPStatsTCPConnectRTTSum interface{}

    // The sum of RTT taken to download the object specified by URL. The type is
    // interface{} with range: 0..4294967295.
    RttMonHTTPStatsTransactionRTTSum interface{}

    // The sum of the size of the message body received as a response to the HTTP
    // request. The type is interface{} with range: 0..4294967295.
    RttMonHTTPStatsMessageBodyOctetsSum interface{}

    // The number of requests that could not connect to the DNS Server. The type
    // is interface{} with range: 0..4294967295.
    RttMonHTTPStatsDNSServerTimeout interface{}

    // The number of requests that could not connect to the the HTTP Server. The
    // type is interface{} with range: 0..4294967295.
    RttMonHTTPStatsTCPConnectTimeout interface{}

    // The number of requests that timed out during HTTP transaction. The type is
    // interface{} with range: 0..4294967295.
    RttMonHTTPStatsTransactionTimeout interface{}

    // The number of requests that had DNS Query errors. The type is interface{}
    // with range: 0..4294967295.
    RttMonHTTPStatsDNSQueryError interface{}

    // The number of requests that had HTTP errors while downloading the base
    // page. The type is interface{} with range: 0..4294967295.
    RttMonHTTPStatsHTTPError interface{}

    // The number of occasions when a HTTP operation could not be initiated
    // because an internal error. The type is interface{} with range:
    // 0..4294967295.
    RttMonHTTPStatsError interface{}

    // The number of occasions when an HTTP operation could not be initiated
    // because a previous HTTP operation has not been completed. The type is
    // interface{} with range: 0..4294967295.
    RttMonHTTPStatsBusies interface{}
}

func (rttMonHTTPStatsEntry *CISCORTTMONMIB_RttMonHTTPStatsTable_RttMonHTTPStatsEntry) GetEntityData() *types.CommonEntityData {
    rttMonHTTPStatsEntry.EntityData.YFilter = rttMonHTTPStatsEntry.YFilter
    rttMonHTTPStatsEntry.EntityData.YangName = "rttMonHTTPStatsEntry"
    rttMonHTTPStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonHTTPStatsEntry.EntityData.ParentYangName = "rttMonHTTPStatsTable"
    rttMonHTTPStatsEntry.EntityData.SegmentPath = "rttMonHTTPStatsEntry" + types.AddKeyToken(rttMonHTTPStatsEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonHTTPStatsEntry.RttMonHTTPStatsStartTimeIndex, "rttMonHTTPStatsStartTimeIndex")
    rttMonHTTPStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonHTTPStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonHTTPStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonHTTPStatsEntry.EntityData.Children = types.NewOrderedMap()
    rttMonHTTPStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonHTTPStatsEntry.RttMonCtrlAdminIndex})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsStartTimeIndex", types.YLeaf{"RttMonHTTPStatsStartTimeIndex", rttMonHTTPStatsEntry.RttMonHTTPStatsStartTimeIndex})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsCompletions", types.YLeaf{"RttMonHTTPStatsCompletions", rttMonHTTPStatsEntry.RttMonHTTPStatsCompletions})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsOverThresholds", types.YLeaf{"RttMonHTTPStatsOverThresholds", rttMonHTTPStatsEntry.RttMonHTTPStatsOverThresholds})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsRTTSum", types.YLeaf{"RttMonHTTPStatsRTTSum", rttMonHTTPStatsEntry.RttMonHTTPStatsRTTSum})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsRTTSum2Low", types.YLeaf{"RttMonHTTPStatsRTTSum2Low", rttMonHTTPStatsEntry.RttMonHTTPStatsRTTSum2Low})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsRTTSum2High", types.YLeaf{"RttMonHTTPStatsRTTSum2High", rttMonHTTPStatsEntry.RttMonHTTPStatsRTTSum2High})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsRTTMin", types.YLeaf{"RttMonHTTPStatsRTTMin", rttMonHTTPStatsEntry.RttMonHTTPStatsRTTMin})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsRTTMax", types.YLeaf{"RttMonHTTPStatsRTTMax", rttMonHTTPStatsEntry.RttMonHTTPStatsRTTMax})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsDNSRTTSum", types.YLeaf{"RttMonHTTPStatsDNSRTTSum", rttMonHTTPStatsEntry.RttMonHTTPStatsDNSRTTSum})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsTCPConnectRTTSum", types.YLeaf{"RttMonHTTPStatsTCPConnectRTTSum", rttMonHTTPStatsEntry.RttMonHTTPStatsTCPConnectRTTSum})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsTransactionRTTSum", types.YLeaf{"RttMonHTTPStatsTransactionRTTSum", rttMonHTTPStatsEntry.RttMonHTTPStatsTransactionRTTSum})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsMessageBodyOctetsSum", types.YLeaf{"RttMonHTTPStatsMessageBodyOctetsSum", rttMonHTTPStatsEntry.RttMonHTTPStatsMessageBodyOctetsSum})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsDNSServerTimeout", types.YLeaf{"RttMonHTTPStatsDNSServerTimeout", rttMonHTTPStatsEntry.RttMonHTTPStatsDNSServerTimeout})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsTCPConnectTimeout", types.YLeaf{"RttMonHTTPStatsTCPConnectTimeout", rttMonHTTPStatsEntry.RttMonHTTPStatsTCPConnectTimeout})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsTransactionTimeout", types.YLeaf{"RttMonHTTPStatsTransactionTimeout", rttMonHTTPStatsEntry.RttMonHTTPStatsTransactionTimeout})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsDNSQueryError", types.YLeaf{"RttMonHTTPStatsDNSQueryError", rttMonHTTPStatsEntry.RttMonHTTPStatsDNSQueryError})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsHTTPError", types.YLeaf{"RttMonHTTPStatsHTTPError", rttMonHTTPStatsEntry.RttMonHTTPStatsHTTPError})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsError", types.YLeaf{"RttMonHTTPStatsError", rttMonHTTPStatsEntry.RttMonHTTPStatsError})
    rttMonHTTPStatsEntry.EntityData.Leafs.Append("rttMonHTTPStatsBusies", types.YLeaf{"RttMonHTTPStatsBusies", rttMonHTTPStatsEntry.RttMonHTTPStatsBusies})

    rttMonHTTPStatsEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonHTTPStatsStartTimeIndex"}

    return &(rttMonHTTPStatsEntry.EntityData)
}

// CISCORTTMONMIB_RttMonJitterStatsTable
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
type CISCORTTMONMIB_RttMonJitterStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects which accumulate the results of a series of RTT
    // operations over a 60 minute time period.  This entry is created only if the
    // rttMonCtrlAdminRttType  is jitter. The operation of this table is same as
    // that of rttMonStatsCaptureTable. The type is slice of
    // CISCORTTMONMIB_RttMonJitterStatsTable_RttMonJitterStatsEntry.
    RttMonJitterStatsEntry []*CISCORTTMONMIB_RttMonJitterStatsTable_RttMonJitterStatsEntry
}

func (rttMonJitterStatsTable *CISCORTTMONMIB_RttMonJitterStatsTable) GetEntityData() *types.CommonEntityData {
    rttMonJitterStatsTable.EntityData.YFilter = rttMonJitterStatsTable.YFilter
    rttMonJitterStatsTable.EntityData.YangName = "rttMonJitterStatsTable"
    rttMonJitterStatsTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonJitterStatsTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonJitterStatsTable.EntityData.SegmentPath = "rttMonJitterStatsTable"
    rttMonJitterStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonJitterStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonJitterStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonJitterStatsTable.EntityData.Children = types.NewOrderedMap()
    rttMonJitterStatsTable.EntityData.Children.Append("rttMonJitterStatsEntry", types.YChild{"RttMonJitterStatsEntry", nil})
    for i := range rttMonJitterStatsTable.RttMonJitterStatsEntry {
        rttMonJitterStatsTable.EntityData.Children.Append(types.GetSegmentPath(rttMonJitterStatsTable.RttMonJitterStatsEntry[i]), types.YChild{"RttMonJitterStatsEntry", rttMonJitterStatsTable.RttMonJitterStatsEntry[i]})
    }
    rttMonJitterStatsTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonJitterStatsTable.EntityData.YListKeys = []string {}

    return &(rttMonJitterStatsTable.EntityData)
}

// CISCORTTMONMIB_RttMonJitterStatsTable_RttMonJitterStatsEntry
// A list of objects which accumulate the results of a
// series of RTT operations over a 60 minute time period.
// 
// This entry is created only if the rttMonCtrlAdminRttType 
// is jitter. The operation of this table is same as that of
// rttMonStatsCaptureTable.
type CISCORTTMONMIB_RttMonJitterStatsTable_RttMonJitterStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. The time when this row was created. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsStartTimeIndex interface{}

    // The number of jitter operation that have completed successfully. The type
    // is interface{} with range: 0..4294967295.
    RttMonJitterStatsCompletions interface{}

    // The number of jitter operations that violate threshold. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsOverThresholds interface{}

    // The number of RTT's that are successfully measured. The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsNumOfRTT interface{}

    // The sum of RTT's that are successfully measured (low order 32 bits). The
    // high order 32 bits are stored in rttMonJitterStatsRTTSumHigh. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsRTTSum interface{}

    // The sum of squares of RTT's that are successfully measured (low order 32
    // bits). The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsRTTSum2Low interface{}

    // The sum of squares of RTT's that are successfully measured (high order 32
    // bits). The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsRTTSum2High interface{}

    // The minimum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsRTTMin interface{}

    // The maximum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsRTTMax interface{}

    // The minimum of absolute values of all positive jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsMinOfPositivesSD interface{}

    // The maximum of absolute values of all positive jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsMaxOfPositivesSD interface{}

    // The sum of number of all positive jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsNumOfPositivesSD interface{}

    // The sum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsSumOfPositivesSD interface{}

    // The sum of square of RTT's of all positive jitter values from packets sent
    // from source to destination (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2PositivesSDLow interface{}

    // The sum of square of RTT's of all positive jitter values from packets sent
    // from source to destination (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2PositivesSDHigh interface{}

    // The minimum of all negative jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsMinOfNegativesSD interface{}

    // The maximum of all negative jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsMaxOfNegativesSD interface{}

    // The sum of number of all negative jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsNumOfNegativesSD interface{}

    // The sum of RTT's of all negative jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsSumOfNegativesSD interface{}

    // The sum of square of RTT's of all negative jitter values from packets sent
    // from source to destination (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2NegativesSDLow interface{}

    // The sum of square of RTT's of all negative jitter values from packets sent
    // from source to destination (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2NegativesSDHigh interface{}

    // The minimum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsMinOfPositivesDS interface{}

    // The maximum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsMaxOfPositivesDS interface{}

    // The sum of number of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsNumOfPositivesDS interface{}

    // The sum of RTT's of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsSumOfPositivesDS interface{}

    // The sum of squares of RTT's of all positive jitter values from packets sent
    // from destination to source (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2PositivesDSLow interface{}

    // The sum of squares of RTT's of all positive jitter values from packets sent
    // from destination to source (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2PositivesDSHigh interface{}

    // The minimum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsMinOfNegativesDS interface{}

    // The maximum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsMaxOfNegativesDS interface{}

    // The sum of number of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsNumOfNegativesDS interface{}

    // The sum of RTT's of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsSumOfNegativesDS interface{}

    // The sum of squares of RTT's of all negative jitter values from packets sent
    // from destination to source (low order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2NegativesDSLow interface{}

    // The sum of squares of RTT's of all negative jitter values from packets sent
    // from destination to source (high order 32 bits). The type is interface{}
    // with range: 0..4294967295.
    RttMonJitterStatsSum2NegativesDSHigh interface{}

    // The number of packets lost when sent from source to destination. The type
    // is interface{} with range: 0..4294967295.
    RttMonJitterStatsPacketLossSD interface{}

    // The number of packets lost when sent from destination to source. The type
    // is interface{} with range: 0..4294967295.
    RttMonJitterStatsPacketLossDS interface{}

    // The number of packets arrived out of sequence. The type is interface{} with
    // range: 0..4294967295.
    RttMonJitterStatsPacketOutOfSequence interface{}

    // The number of packets that are lost for which we cannot determine the
    // direction. The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsPacketMIA interface{}

    // The number of packets that arrived after the timeout. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsPacketLateArrival interface{}

    // The number of occasions when a jitter operation could not be initiated
    // because an internal error. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsError interface{}

    // The number of occasions when a jitter operation could not be initiated
    // because a previous jitter operation has not been completed. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsBusies interface{}

    // The sum of one way times from source to destination (low order 32 bits).
    // The high order 32 bits are stored in rttMonJitterStatsOWSumSDHigh. The type
    // is interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSumSD interface{}

    // The sum of squares of one way times from source to destination (low order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSum2SDLow interface{}

    // The sum of squares of one way times from source to destination (high order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSum2SDHigh interface{}

    // The minimum of all one way times from source to destination.
    // rttMonJitterStatsOWMinSD object is superseded by
    // rttMonJitterStatsOWMinSDNew. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMinSD interface{}

    // The maximum of all one way times from source to destination.
    // rttMonJitterStatsOWMaxSD object is superseded by
    // rttMonJitterStatsOWMaxSDNew. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMaxSD interface{}

    // The sum of one way times from destination to source (low order 32 bits).
    // The high order 32 bits are stored in rttMonJitterStatsOWSumDSHigh. The type
    // is interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSumDS interface{}

    // The sum of squares of one way times from destination to source (low order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSum2DSLow interface{}

    // The sum of squares of one way times from destination to source (high order
    // 32 bits). The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSum2DSHigh interface{}

    // The minimum of all one way times from destination to source.
    // rttMonJitterStatsOWMinDS object is superseded by
    // rttMonJitterStatsOWMinDSNew. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMinDS interface{}

    // The maximum of all one way times from destination to source.
    // rttMonJitterStatsOWMaxDS object is superseded by
    // rttMonJitterStatsOWMaxDSNew. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMaxDS interface{}

    // The number of one way times that are successfully measured. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsNumOfOW interface{}

    // The minimum of all one way times from source to destination. Replaces
    // deprecated rttMonJitterStatsOWMinSD. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMinSDNew interface{}

    // The maximum of all one way times from source to destination. Replaces
    // deprecated rttMonJitterStatsOWMaxSD. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMaxSDNew interface{}

    // The minimum of all one way times from destination to source. Replaces
    // deprecated rttMonJitterStatsOWMinDS. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMinDSNew interface{}

    // The maximum of all one way times from destination to source. Replaces
    // deprecated rttMonJitterStatsOWMaxDS. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsOWMaxDSNew interface{}

    // The minimum of all MOS values for the jitter operations in hundreds.  This
    // value will be 0 if   - rttMonEchoAdminCodecType of the operation is
    // notApplicable   - the operation is not started   - the operation is started
    // but failed This value will be 1 for packet loss of 10% or more. The type is
    // interface{} with range: 0..None | 100..500.
    RttMonJitterStatsMinOfMOS interface{}

    // The maximum of all MOS values for the jitter operations in hunderds.  This
    // value will be 0 if   - rttMonEchoAdminCodecType of the operation is
    // notApplicable   - the operation is not started   - the operation is started
    // but failed This value will be 1 for packet loss of 10% or more. The type is
    // interface{} with range: 0..None | 100..500.
    RttMonJitterStatsMaxOfMOS interface{}

    // The minimum of all ICPIF values for the jitter operations.  This value will
    // be 93 for packet loss of 10% or more. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsMinOfICPIF interface{}

    // The maximum of all ICPIF values for the jitter operations.  This value will
    // be 93 for packet loss of 10% or more. The type is interface{} with range:
    // 0..4294967295.
    RttMonJitterStatsMaxOfICPIF interface{}

    // Interarrival Jitter (RFC 1889) at responder. The type is interface{} with
    // range: 0..4294967295.
    RttMonJitterStatsIAJOut interface{}

    // Interarrival Jitter (RFC 1889) at sender. The type is interface{} with
    // range: 0..4294967295.
    RttMonJitterStatsIAJIn interface{}

    // The average of positive and negative jitter values for SD and DS direction.
    // The type is interface{} with range: 0..4294967295.
    RttMonJitterStatsAvgJitter interface{}

    // The average of positive and negative jitter values in SD direction. The
    // type is interface{} with range: 0..4294967295.
    RttMonJitterStatsAvgJitterSD interface{}

    // The average of positive and negative jitter values in DS direction. The
    // type is interface{} with range: 0..4294967295.
    RttMonJitterStatsAvgJitterDS interface{}

    // The number of RTT operations that have completed with sender and responder
    // out of sync with NTP. The NTP sync means  the total of NTP offset on sender
    // and responder is within  configured tolerance level. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsUnSyncRTs interface{}

    // The sum of RTT's that are successfully measured (high order 32 bits). The
    // low order 32 bits are  stored in rttMonJitterStatsRTTSum. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsRTTSumHigh interface{}

    // The sum of one way times from source to destination (high order 32 bits).
    // The low order 32 bits are  stored in rttMonJitterStatsOWSumSD. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSumSDHigh interface{}

    // The sum of one way times from destination to source (high order 32 bits).
    // The low order 32 bits are stored in rttMonJitterStatsOWSumDS. The type is
    // interface{} with range: 0..4294967295.
    RttMonJitterStatsOWSumDSHigh interface{}
}

func (rttMonJitterStatsEntry *CISCORTTMONMIB_RttMonJitterStatsTable_RttMonJitterStatsEntry) GetEntityData() *types.CommonEntityData {
    rttMonJitterStatsEntry.EntityData.YFilter = rttMonJitterStatsEntry.YFilter
    rttMonJitterStatsEntry.EntityData.YangName = "rttMonJitterStatsEntry"
    rttMonJitterStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonJitterStatsEntry.EntityData.ParentYangName = "rttMonJitterStatsTable"
    rttMonJitterStatsEntry.EntityData.SegmentPath = "rttMonJitterStatsEntry" + types.AddKeyToken(rttMonJitterStatsEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonJitterStatsEntry.RttMonJitterStatsStartTimeIndex, "rttMonJitterStatsStartTimeIndex")
    rttMonJitterStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonJitterStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonJitterStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonJitterStatsEntry.EntityData.Children = types.NewOrderedMap()
    rttMonJitterStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonJitterStatsEntry.RttMonCtrlAdminIndex})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsStartTimeIndex", types.YLeaf{"RttMonJitterStatsStartTimeIndex", rttMonJitterStatsEntry.RttMonJitterStatsStartTimeIndex})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsCompletions", types.YLeaf{"RttMonJitterStatsCompletions", rttMonJitterStatsEntry.RttMonJitterStatsCompletions})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOverThresholds", types.YLeaf{"RttMonJitterStatsOverThresholds", rttMonJitterStatsEntry.RttMonJitterStatsOverThresholds})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsNumOfRTT", types.YLeaf{"RttMonJitterStatsNumOfRTT", rttMonJitterStatsEntry.RttMonJitterStatsNumOfRTT})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsRTTSum", types.YLeaf{"RttMonJitterStatsRTTSum", rttMonJitterStatsEntry.RttMonJitterStatsRTTSum})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsRTTSum2Low", types.YLeaf{"RttMonJitterStatsRTTSum2Low", rttMonJitterStatsEntry.RttMonJitterStatsRTTSum2Low})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsRTTSum2High", types.YLeaf{"RttMonJitterStatsRTTSum2High", rttMonJitterStatsEntry.RttMonJitterStatsRTTSum2High})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsRTTMin", types.YLeaf{"RttMonJitterStatsRTTMin", rttMonJitterStatsEntry.RttMonJitterStatsRTTMin})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsRTTMax", types.YLeaf{"RttMonJitterStatsRTTMax", rttMonJitterStatsEntry.RttMonJitterStatsRTTMax})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMinOfPositivesSD", types.YLeaf{"RttMonJitterStatsMinOfPositivesSD", rttMonJitterStatsEntry.RttMonJitterStatsMinOfPositivesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMaxOfPositivesSD", types.YLeaf{"RttMonJitterStatsMaxOfPositivesSD", rttMonJitterStatsEntry.RttMonJitterStatsMaxOfPositivesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsNumOfPositivesSD", types.YLeaf{"RttMonJitterStatsNumOfPositivesSD", rttMonJitterStatsEntry.RttMonJitterStatsNumOfPositivesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSumOfPositivesSD", types.YLeaf{"RttMonJitterStatsSumOfPositivesSD", rttMonJitterStatsEntry.RttMonJitterStatsSumOfPositivesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2PositivesSDLow", types.YLeaf{"RttMonJitterStatsSum2PositivesSDLow", rttMonJitterStatsEntry.RttMonJitterStatsSum2PositivesSDLow})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2PositivesSDHigh", types.YLeaf{"RttMonJitterStatsSum2PositivesSDHigh", rttMonJitterStatsEntry.RttMonJitterStatsSum2PositivesSDHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMinOfNegativesSD", types.YLeaf{"RttMonJitterStatsMinOfNegativesSD", rttMonJitterStatsEntry.RttMonJitterStatsMinOfNegativesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMaxOfNegativesSD", types.YLeaf{"RttMonJitterStatsMaxOfNegativesSD", rttMonJitterStatsEntry.RttMonJitterStatsMaxOfNegativesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsNumOfNegativesSD", types.YLeaf{"RttMonJitterStatsNumOfNegativesSD", rttMonJitterStatsEntry.RttMonJitterStatsNumOfNegativesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSumOfNegativesSD", types.YLeaf{"RttMonJitterStatsSumOfNegativesSD", rttMonJitterStatsEntry.RttMonJitterStatsSumOfNegativesSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2NegativesSDLow", types.YLeaf{"RttMonJitterStatsSum2NegativesSDLow", rttMonJitterStatsEntry.RttMonJitterStatsSum2NegativesSDLow})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2NegativesSDHigh", types.YLeaf{"RttMonJitterStatsSum2NegativesSDHigh", rttMonJitterStatsEntry.RttMonJitterStatsSum2NegativesSDHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMinOfPositivesDS", types.YLeaf{"RttMonJitterStatsMinOfPositivesDS", rttMonJitterStatsEntry.RttMonJitterStatsMinOfPositivesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMaxOfPositivesDS", types.YLeaf{"RttMonJitterStatsMaxOfPositivesDS", rttMonJitterStatsEntry.RttMonJitterStatsMaxOfPositivesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsNumOfPositivesDS", types.YLeaf{"RttMonJitterStatsNumOfPositivesDS", rttMonJitterStatsEntry.RttMonJitterStatsNumOfPositivesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSumOfPositivesDS", types.YLeaf{"RttMonJitterStatsSumOfPositivesDS", rttMonJitterStatsEntry.RttMonJitterStatsSumOfPositivesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2PositivesDSLow", types.YLeaf{"RttMonJitterStatsSum2PositivesDSLow", rttMonJitterStatsEntry.RttMonJitterStatsSum2PositivesDSLow})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2PositivesDSHigh", types.YLeaf{"RttMonJitterStatsSum2PositivesDSHigh", rttMonJitterStatsEntry.RttMonJitterStatsSum2PositivesDSHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMinOfNegativesDS", types.YLeaf{"RttMonJitterStatsMinOfNegativesDS", rttMonJitterStatsEntry.RttMonJitterStatsMinOfNegativesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMaxOfNegativesDS", types.YLeaf{"RttMonJitterStatsMaxOfNegativesDS", rttMonJitterStatsEntry.RttMonJitterStatsMaxOfNegativesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsNumOfNegativesDS", types.YLeaf{"RttMonJitterStatsNumOfNegativesDS", rttMonJitterStatsEntry.RttMonJitterStatsNumOfNegativesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSumOfNegativesDS", types.YLeaf{"RttMonJitterStatsSumOfNegativesDS", rttMonJitterStatsEntry.RttMonJitterStatsSumOfNegativesDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2NegativesDSLow", types.YLeaf{"RttMonJitterStatsSum2NegativesDSLow", rttMonJitterStatsEntry.RttMonJitterStatsSum2NegativesDSLow})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsSum2NegativesDSHigh", types.YLeaf{"RttMonJitterStatsSum2NegativesDSHigh", rttMonJitterStatsEntry.RttMonJitterStatsSum2NegativesDSHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsPacketLossSD", types.YLeaf{"RttMonJitterStatsPacketLossSD", rttMonJitterStatsEntry.RttMonJitterStatsPacketLossSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsPacketLossDS", types.YLeaf{"RttMonJitterStatsPacketLossDS", rttMonJitterStatsEntry.RttMonJitterStatsPacketLossDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsPacketOutOfSequence", types.YLeaf{"RttMonJitterStatsPacketOutOfSequence", rttMonJitterStatsEntry.RttMonJitterStatsPacketOutOfSequence})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsPacketMIA", types.YLeaf{"RttMonJitterStatsPacketMIA", rttMonJitterStatsEntry.RttMonJitterStatsPacketMIA})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsPacketLateArrival", types.YLeaf{"RttMonJitterStatsPacketLateArrival", rttMonJitterStatsEntry.RttMonJitterStatsPacketLateArrival})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsError", types.YLeaf{"RttMonJitterStatsError", rttMonJitterStatsEntry.RttMonJitterStatsError})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsBusies", types.YLeaf{"RttMonJitterStatsBusies", rttMonJitterStatsEntry.RttMonJitterStatsBusies})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSumSD", types.YLeaf{"RttMonJitterStatsOWSumSD", rttMonJitterStatsEntry.RttMonJitterStatsOWSumSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSum2SDLow", types.YLeaf{"RttMonJitterStatsOWSum2SDLow", rttMonJitterStatsEntry.RttMonJitterStatsOWSum2SDLow})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSum2SDHigh", types.YLeaf{"RttMonJitterStatsOWSum2SDHigh", rttMonJitterStatsEntry.RttMonJitterStatsOWSum2SDHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMinSD", types.YLeaf{"RttMonJitterStatsOWMinSD", rttMonJitterStatsEntry.RttMonJitterStatsOWMinSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMaxSD", types.YLeaf{"RttMonJitterStatsOWMaxSD", rttMonJitterStatsEntry.RttMonJitterStatsOWMaxSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSumDS", types.YLeaf{"RttMonJitterStatsOWSumDS", rttMonJitterStatsEntry.RttMonJitterStatsOWSumDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSum2DSLow", types.YLeaf{"RttMonJitterStatsOWSum2DSLow", rttMonJitterStatsEntry.RttMonJitterStatsOWSum2DSLow})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSum2DSHigh", types.YLeaf{"RttMonJitterStatsOWSum2DSHigh", rttMonJitterStatsEntry.RttMonJitterStatsOWSum2DSHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMinDS", types.YLeaf{"RttMonJitterStatsOWMinDS", rttMonJitterStatsEntry.RttMonJitterStatsOWMinDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMaxDS", types.YLeaf{"RttMonJitterStatsOWMaxDS", rttMonJitterStatsEntry.RttMonJitterStatsOWMaxDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsNumOfOW", types.YLeaf{"RttMonJitterStatsNumOfOW", rttMonJitterStatsEntry.RttMonJitterStatsNumOfOW})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMinSDNew", types.YLeaf{"RttMonJitterStatsOWMinSDNew", rttMonJitterStatsEntry.RttMonJitterStatsOWMinSDNew})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMaxSDNew", types.YLeaf{"RttMonJitterStatsOWMaxSDNew", rttMonJitterStatsEntry.RttMonJitterStatsOWMaxSDNew})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMinDSNew", types.YLeaf{"RttMonJitterStatsOWMinDSNew", rttMonJitterStatsEntry.RttMonJitterStatsOWMinDSNew})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWMaxDSNew", types.YLeaf{"RttMonJitterStatsOWMaxDSNew", rttMonJitterStatsEntry.RttMonJitterStatsOWMaxDSNew})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMinOfMOS", types.YLeaf{"RttMonJitterStatsMinOfMOS", rttMonJitterStatsEntry.RttMonJitterStatsMinOfMOS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMaxOfMOS", types.YLeaf{"RttMonJitterStatsMaxOfMOS", rttMonJitterStatsEntry.RttMonJitterStatsMaxOfMOS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMinOfICPIF", types.YLeaf{"RttMonJitterStatsMinOfICPIF", rttMonJitterStatsEntry.RttMonJitterStatsMinOfICPIF})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsMaxOfICPIF", types.YLeaf{"RttMonJitterStatsMaxOfICPIF", rttMonJitterStatsEntry.RttMonJitterStatsMaxOfICPIF})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsIAJOut", types.YLeaf{"RttMonJitterStatsIAJOut", rttMonJitterStatsEntry.RttMonJitterStatsIAJOut})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsIAJIn", types.YLeaf{"RttMonJitterStatsIAJIn", rttMonJitterStatsEntry.RttMonJitterStatsIAJIn})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsAvgJitter", types.YLeaf{"RttMonJitterStatsAvgJitter", rttMonJitterStatsEntry.RttMonJitterStatsAvgJitter})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsAvgJitterSD", types.YLeaf{"RttMonJitterStatsAvgJitterSD", rttMonJitterStatsEntry.RttMonJitterStatsAvgJitterSD})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsAvgJitterDS", types.YLeaf{"RttMonJitterStatsAvgJitterDS", rttMonJitterStatsEntry.RttMonJitterStatsAvgJitterDS})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsUnSyncRTs", types.YLeaf{"RttMonJitterStatsUnSyncRTs", rttMonJitterStatsEntry.RttMonJitterStatsUnSyncRTs})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsRTTSumHigh", types.YLeaf{"RttMonJitterStatsRTTSumHigh", rttMonJitterStatsEntry.RttMonJitterStatsRTTSumHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSumSDHigh", types.YLeaf{"RttMonJitterStatsOWSumSDHigh", rttMonJitterStatsEntry.RttMonJitterStatsOWSumSDHigh})
    rttMonJitterStatsEntry.EntityData.Leafs.Append("rttMonJitterStatsOWSumDSHigh", types.YLeaf{"RttMonJitterStatsOWSumDSHigh", rttMonJitterStatsEntry.RttMonJitterStatsOWSumDSHigh})

    rttMonJitterStatsEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonJitterStatsStartTimeIndex"}

    return &(rttMonJitterStatsEntry.EntityData)
}

// CISCORTTMONMIB_RttMonLpdGrpStatsTable
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
type CISCORTTMONMIB_RttMonLpdGrpStatsTable struct {
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
    // CISCORTTMONMIB_RttMonLpdGrpStatsTable_RttMonLpdGrpStatsEntry.
    RttMonLpdGrpStatsEntry []*CISCORTTMONMIB_RttMonLpdGrpStatsTable_RttMonLpdGrpStatsEntry
}

func (rttMonLpdGrpStatsTable *CISCORTTMONMIB_RttMonLpdGrpStatsTable) GetEntityData() *types.CommonEntityData {
    rttMonLpdGrpStatsTable.EntityData.YFilter = rttMonLpdGrpStatsTable.YFilter
    rttMonLpdGrpStatsTable.EntityData.YangName = "rttMonLpdGrpStatsTable"
    rttMonLpdGrpStatsTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonLpdGrpStatsTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonLpdGrpStatsTable.EntityData.SegmentPath = "rttMonLpdGrpStatsTable"
    rttMonLpdGrpStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonLpdGrpStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonLpdGrpStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonLpdGrpStatsTable.EntityData.Children = types.NewOrderedMap()
    rttMonLpdGrpStatsTable.EntityData.Children.Append("rttMonLpdGrpStatsEntry", types.YChild{"RttMonLpdGrpStatsEntry", nil})
    for i := range rttMonLpdGrpStatsTable.RttMonLpdGrpStatsEntry {
        rttMonLpdGrpStatsTable.EntityData.Children.Append(types.GetSegmentPath(rttMonLpdGrpStatsTable.RttMonLpdGrpStatsEntry[i]), types.YChild{"RttMonLpdGrpStatsEntry", rttMonLpdGrpStatsTable.RttMonLpdGrpStatsEntry[i]})
    }
    rttMonLpdGrpStatsTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonLpdGrpStatsTable.EntityData.YListKeys = []string {}

    return &(rttMonLpdGrpStatsTable.EntityData)
}

// CISCORTTMONMIB_RttMonLpdGrpStatsTable_RttMonLpdGrpStatsEntry
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
type CISCORTTMONMIB_RttMonLpdGrpStatsTable_RttMonLpdGrpStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Uniquely identifies a row in
    // rttMonLpdGrpStatsTable.  This is a pseudo-random number which identifies a
    // particular LPD group. The type is interface{} with range: 1..2147483647.
    RttMonLpdGrpStatsGroupIndex interface{}

    // This attribute is a key. The time when this row was created.  This object
    // is the second index of the rttMonLpdGrpStatsTable. When the number of
    // rttMonLpdGrpStatsStartTimeIndex groups exceeds the
    // rttMplsVpnMonTypeLpdStatHours value, the oldest
    // rttMonLpdGrpStatsStartTimeIndex group will be removed and replaced with the
    // new entry. The type is interface{} with range: 0..4294967295.
    RttMonLpdGrpStatsStartTimeIndex interface{}

    // The object is a string that specifies the address of the target PE for this
    // LPD group. The type is string.
    RttMonLpdGrpStatsTargetPE interface{}

    // This object represents the number of successfull completions of 'single
    // probes' for all the set of paths in the LPD group.  Whenever the
    // rttMonLatestRttOperSense value is 'ok' for a particular probe in the LPD
    // Group this object will be incremented.  This object will be set to '0' on
    // reset. The type is interface{} with range: 0..2147483647. Units are passes.
    RttMonLpdGrpStatsNumOfPass interface{}

    // This object represents the number of failed operations of 'single probes'
    // for all the set of paths in the LPD group.  Whenever the
    // rttMonLatestRttOperSense has a value other than 'ok' or 'timeout' for a
    // particular probe in the LPD Group this object will be incremented.  This
    // object will be set to '0' on reset. The type is interface{} with range:
    // 0..2147483647. Units are failures.
    RttMonLpdGrpStatsNumOfFail interface{}

    // This object represents the number of timed out operations of 'single
    // probes' for all the set of paths in the LPD group.  Whenever the
    // rttMonLatestRttOperSense has a value of 'timeout' for a particular probe in
    // the LPD Group this object will be incremented.  This object will be set to
    // '0' on reset. The type is interface{} with range: 0..2147483647. Units are
    // timeouts.
    RttMonLpdGrpStatsNumOfTimeout interface{}

    // The average RTT across all set of probes in the LPD group.  This object
    // will be set to '0' on reset. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    RttMonLpdGrpStatsAvgRTT interface{}

    // The minimum of RTT's for all set of probes in the LPD group that were
    // successfully measured.  This object will be set to '0' on reset. The type
    // is interface{} with range: 0..2147483647. Units are milliseconds.
    RttMonLpdGrpStatsMinRTT interface{}

    // The maximum of RTT's for all set of probes in the LPD group that were
    // successfully measured.  This object will be set to '0' on reset. The type
    // is interface{} with range: 0..2147483647. Units are milliseconds.
    RttMonLpdGrpStatsMaxRTT interface{}

    // The minimum number of active paths discovered to the
    // rttMonLpdGrpStatsTargetPE target.  This object will be set to '0' on reset.
    // The type is interface{} with range: 0..2147483647. Units are paths.
    RttMonLpdGrpStatsMinNumPaths interface{}

    // The maximum number of active paths discovered to the
    // rttMonLpdGrpStatsTargetPE target.  This object will be set to '0' on reset.
    // The type is interface{} with range: 0..2147483647. Units are paths.
    RttMonLpdGrpStatsMaxNumPaths interface{}

    // The time when the last LSP Path Discovery to the group was attempted.  This
    // object will be set to '0' on reset. The type is interface{} with range:
    // 0..4294967295. Units are tenths of milliseconds.
    RttMonLpdGrpStatsLPDStartTime interface{}

    // This object is set to true when the LSP Path Discovery to the target PE
    // i.e. rttMonLpdGrpStatsTargetPE fails, and set to false when the LSP Path
    // Discovery succeeds.  When this value changes and
    // rttMplsVpnMonReactLpdNotifyType is set to 'lpdPathDiscovery' or 'lpdAll' a
    // rttMonLpdDiscoveryNotification will be generated.  This object will be set
    // to 'FALSE' on reset. The type is bool.
    RttMonLpdGrpStatsLPDFailOccurred interface{}

    // This object identifies the cause of failure for the LSP Path Discovery last
    // attempted. It will be only valid if rttMonLpdGrpStatsLPDFailOccurred is set
    // to true.  This object will be set to 'unknown' on reset. The type is
    // RttMplsVpnMonLpdFailureSense.
    RttMonLpdGrpStatsLPDFailCause interface{}

    // The completion time of the last successfull LSP Path Discovery to the
    // target PE.  This object will be set to '0' on reset. The type is
    // interface{} with range: 0..65535. Units are seconds.
    RttMonLpdGrpStatsLPDCompTime interface{}

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
    RttMonLpdGrpStatsGroupStatus interface{}

    // This object identifies 'lspGroup' probe uniquely created for this
    // particular LPD Group. The type is interface{} with range: 1..2147483647.
    // Units are identifier.
    RttMonLpdGrpStatsGroupProbeIndex interface{}

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
    RttMonLpdGrpStatsPathIds interface{}

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
    RttMonLpdGrpStatsProbeStatus interface{}

    // This object specifies the time when this statistics row was last reset
    // using the rttMonApplLpdGrpStatsReset object. The type is interface{} with
    // range: 0..4294967295.
    RttMonLpdGrpStatsResetTime interface{}
}

func (rttMonLpdGrpStatsEntry *CISCORTTMONMIB_RttMonLpdGrpStatsTable_RttMonLpdGrpStatsEntry) GetEntityData() *types.CommonEntityData {
    rttMonLpdGrpStatsEntry.EntityData.YFilter = rttMonLpdGrpStatsEntry.YFilter
    rttMonLpdGrpStatsEntry.EntityData.YangName = "rttMonLpdGrpStatsEntry"
    rttMonLpdGrpStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonLpdGrpStatsEntry.EntityData.ParentYangName = "rttMonLpdGrpStatsTable"
    rttMonLpdGrpStatsEntry.EntityData.SegmentPath = "rttMonLpdGrpStatsEntry" + types.AddKeyToken(rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsGroupIndex, "rttMonLpdGrpStatsGroupIndex") + types.AddKeyToken(rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsStartTimeIndex, "rttMonLpdGrpStatsStartTimeIndex")
    rttMonLpdGrpStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonLpdGrpStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonLpdGrpStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonLpdGrpStatsEntry.EntityData.Children = types.NewOrderedMap()
    rttMonLpdGrpStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsGroupIndex", types.YLeaf{"RttMonLpdGrpStatsGroupIndex", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsGroupIndex})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsStartTimeIndex", types.YLeaf{"RttMonLpdGrpStatsStartTimeIndex", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsStartTimeIndex})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsTargetPE", types.YLeaf{"RttMonLpdGrpStatsTargetPE", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsTargetPE})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsNumOfPass", types.YLeaf{"RttMonLpdGrpStatsNumOfPass", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsNumOfPass})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsNumOfFail", types.YLeaf{"RttMonLpdGrpStatsNumOfFail", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsNumOfFail})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsNumOfTimeout", types.YLeaf{"RttMonLpdGrpStatsNumOfTimeout", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsNumOfTimeout})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsAvgRTT", types.YLeaf{"RttMonLpdGrpStatsAvgRTT", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsAvgRTT})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsMinRTT", types.YLeaf{"RttMonLpdGrpStatsMinRTT", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsMinRTT})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsMaxRTT", types.YLeaf{"RttMonLpdGrpStatsMaxRTT", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsMaxRTT})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsMinNumPaths", types.YLeaf{"RttMonLpdGrpStatsMinNumPaths", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsMinNumPaths})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsMaxNumPaths", types.YLeaf{"RttMonLpdGrpStatsMaxNumPaths", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsMaxNumPaths})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsLPDStartTime", types.YLeaf{"RttMonLpdGrpStatsLPDStartTime", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsLPDStartTime})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsLPDFailOccurred", types.YLeaf{"RttMonLpdGrpStatsLPDFailOccurred", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsLPDFailOccurred})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsLPDFailCause", types.YLeaf{"RttMonLpdGrpStatsLPDFailCause", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsLPDFailCause})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsLPDCompTime", types.YLeaf{"RttMonLpdGrpStatsLPDCompTime", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsLPDCompTime})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsGroupStatus", types.YLeaf{"RttMonLpdGrpStatsGroupStatus", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsGroupStatus})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsGroupProbeIndex", types.YLeaf{"RttMonLpdGrpStatsGroupProbeIndex", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsGroupProbeIndex})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsPathIds", types.YLeaf{"RttMonLpdGrpStatsPathIds", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsPathIds})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsProbeStatus", types.YLeaf{"RttMonLpdGrpStatsProbeStatus", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsProbeStatus})
    rttMonLpdGrpStatsEntry.EntityData.Leafs.Append("rttMonLpdGrpStatsResetTime", types.YLeaf{"RttMonLpdGrpStatsResetTime", rttMonLpdGrpStatsEntry.RttMonLpdGrpStatsResetTime})

    rttMonLpdGrpStatsEntry.EntityData.YListKeys = []string {"RttMonLpdGrpStatsGroupIndex", "RttMonLpdGrpStatsStartTimeIndex"}

    return &(rttMonLpdGrpStatsEntry.EntityData)
}

// CISCORTTMONMIB_RttMonHistoryCollectionTable
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
type CISCORTTMONMIB_RttMonHistoryCollectionTable struct {
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
    // CISCORTTMONMIB_RttMonHistoryCollectionTable_RttMonHistoryCollectionEntry.
    RttMonHistoryCollectionEntry []*CISCORTTMONMIB_RttMonHistoryCollectionTable_RttMonHistoryCollectionEntry
}

func (rttMonHistoryCollectionTable *CISCORTTMONMIB_RttMonHistoryCollectionTable) GetEntityData() *types.CommonEntityData {
    rttMonHistoryCollectionTable.EntityData.YFilter = rttMonHistoryCollectionTable.YFilter
    rttMonHistoryCollectionTable.EntityData.YangName = "rttMonHistoryCollectionTable"
    rttMonHistoryCollectionTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonHistoryCollectionTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonHistoryCollectionTable.EntityData.SegmentPath = "rttMonHistoryCollectionTable"
    rttMonHistoryCollectionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonHistoryCollectionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonHistoryCollectionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonHistoryCollectionTable.EntityData.Children = types.NewOrderedMap()
    rttMonHistoryCollectionTable.EntityData.Children.Append("rttMonHistoryCollectionEntry", types.YChild{"RttMonHistoryCollectionEntry", nil})
    for i := range rttMonHistoryCollectionTable.RttMonHistoryCollectionEntry {
        rttMonHistoryCollectionTable.EntityData.Children.Append(types.GetSegmentPath(rttMonHistoryCollectionTable.RttMonHistoryCollectionEntry[i]), types.YChild{"RttMonHistoryCollectionEntry", rttMonHistoryCollectionTable.RttMonHistoryCollectionEntry[i]})
    }
    rttMonHistoryCollectionTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonHistoryCollectionTable.EntityData.YListKeys = []string {}

    return &(rttMonHistoryCollectionTable.EntityData)
}

// CISCORTTMONMIB_RttMonHistoryCollectionTable_RttMonHistoryCollectionEntry
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
type CISCORTTMONMIB_RttMonHistoryCollectionTable_RttMonHistoryCollectionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // This attribute is a key. This uniquely defines a life for a conceptual
    // history row.  For a particular value of rttMonHistoryCollectionLifeIndex,
    // the agent assigns the first value of 1, the second value  of 2, and so on. 
    // The sequence keeps incrementing,  despite older (lower) values being
    // removed from the  table. The type is interface{} with range: 1..2147483647.
    RttMonHistoryCollectionLifeIndex interface{}

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
    RttMonHistoryCollectionBucketIndex interface{}

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
    RttMonHistoryCollectionSampleIndex interface{}

    // The time that the RTT operation was initiated. The type is interface{} with
    // range: 0..4294967295.
    RttMonHistoryCollectionSampleTime interface{}

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
    RttMonHistoryCollectionAddress interface{}

    // This is the operation completion time of the RTT operation.  If the RTT
    // operation fails  (rttMonHistoryCollectionSense is any  value other than
    // ok), this has a value of 0. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    RttMonHistoryCollectionCompletionTime interface{}

    // A sense code for the completion status of the RTT operation. The type is
    // RttResponseSense.
    RttMonHistoryCollectionSense interface{}

    // An application specific sense code for the completion status of the last
    // RTT operation.  This  object will only be valid when the 
    // rttMonHistoryCollectionSense object is set to  'applicationSpecific'. 
    // Otherwise, this object's  value is not valid. The type is interface{} with
    // range: 0..2147483647.
    RttMonHistoryCollectionApplSpecificSense interface{}

    // A sense description for the completion status of the last RTT operation
    // when the  rttMonHistoryCollectionSense object is set to 
    // 'applicationSpecific'. The type is string.
    RttMonHistoryCollectionSenseDescription interface{}
}

func (rttMonHistoryCollectionEntry *CISCORTTMONMIB_RttMonHistoryCollectionTable_RttMonHistoryCollectionEntry) GetEntityData() *types.CommonEntityData {
    rttMonHistoryCollectionEntry.EntityData.YFilter = rttMonHistoryCollectionEntry.YFilter
    rttMonHistoryCollectionEntry.EntityData.YangName = "rttMonHistoryCollectionEntry"
    rttMonHistoryCollectionEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonHistoryCollectionEntry.EntityData.ParentYangName = "rttMonHistoryCollectionTable"
    rttMonHistoryCollectionEntry.EntityData.SegmentPath = "rttMonHistoryCollectionEntry" + types.AddKeyToken(rttMonHistoryCollectionEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex") + types.AddKeyToken(rttMonHistoryCollectionEntry.RttMonHistoryCollectionLifeIndex, "rttMonHistoryCollectionLifeIndex") + types.AddKeyToken(rttMonHistoryCollectionEntry.RttMonHistoryCollectionBucketIndex, "rttMonHistoryCollectionBucketIndex") + types.AddKeyToken(rttMonHistoryCollectionEntry.RttMonHistoryCollectionSampleIndex, "rttMonHistoryCollectionSampleIndex")
    rttMonHistoryCollectionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonHistoryCollectionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonHistoryCollectionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonHistoryCollectionEntry.EntityData.Children = types.NewOrderedMap()
    rttMonHistoryCollectionEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonHistoryCollectionEntry.RttMonCtrlAdminIndex})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionLifeIndex", types.YLeaf{"RttMonHistoryCollectionLifeIndex", rttMonHistoryCollectionEntry.RttMonHistoryCollectionLifeIndex})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionBucketIndex", types.YLeaf{"RttMonHistoryCollectionBucketIndex", rttMonHistoryCollectionEntry.RttMonHistoryCollectionBucketIndex})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionSampleIndex", types.YLeaf{"RttMonHistoryCollectionSampleIndex", rttMonHistoryCollectionEntry.RttMonHistoryCollectionSampleIndex})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionSampleTime", types.YLeaf{"RttMonHistoryCollectionSampleTime", rttMonHistoryCollectionEntry.RttMonHistoryCollectionSampleTime})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionAddress", types.YLeaf{"RttMonHistoryCollectionAddress", rttMonHistoryCollectionEntry.RttMonHistoryCollectionAddress})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionCompletionTime", types.YLeaf{"RttMonHistoryCollectionCompletionTime", rttMonHistoryCollectionEntry.RttMonHistoryCollectionCompletionTime})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionSense", types.YLeaf{"RttMonHistoryCollectionSense", rttMonHistoryCollectionEntry.RttMonHistoryCollectionSense})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionApplSpecificSense", types.YLeaf{"RttMonHistoryCollectionApplSpecificSense", rttMonHistoryCollectionEntry.RttMonHistoryCollectionApplSpecificSense})
    rttMonHistoryCollectionEntry.EntityData.Leafs.Append("rttMonHistoryCollectionSenseDescription", types.YLeaf{"RttMonHistoryCollectionSenseDescription", rttMonHistoryCollectionEntry.RttMonHistoryCollectionSenseDescription})

    rttMonHistoryCollectionEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex", "RttMonHistoryCollectionLifeIndex", "RttMonHistoryCollectionBucketIndex", "RttMonHistoryCollectionSampleIndex"}

    return &(rttMonHistoryCollectionEntry.EntityData)
}

// CISCORTTMONMIB_RttMonLatestHTTPOperTable
// A table which contains the status of latest HTTP RTT
// operation.
type CISCORTTMONMIB_RttMonLatestHTTPOperTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that record the latest HTTP RTT operation. This entry is
    // created automatically after the  rttMonCtrlAdminEntry is created. Also the
    // entry is  automatically deleted when rttMonCtrlAdminEntry is deleted. The
    // type is slice of
    // CISCORTTMONMIB_RttMonLatestHTTPOperTable_RttMonLatestHTTPOperEntry.
    RttMonLatestHTTPOperEntry []*CISCORTTMONMIB_RttMonLatestHTTPOperTable_RttMonLatestHTTPOperEntry
}

func (rttMonLatestHTTPOperTable *CISCORTTMONMIB_RttMonLatestHTTPOperTable) GetEntityData() *types.CommonEntityData {
    rttMonLatestHTTPOperTable.EntityData.YFilter = rttMonLatestHTTPOperTable.YFilter
    rttMonLatestHTTPOperTable.EntityData.YangName = "rttMonLatestHTTPOperTable"
    rttMonLatestHTTPOperTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonLatestHTTPOperTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonLatestHTTPOperTable.EntityData.SegmentPath = "rttMonLatestHTTPOperTable"
    rttMonLatestHTTPOperTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonLatestHTTPOperTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonLatestHTTPOperTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonLatestHTTPOperTable.EntityData.Children = types.NewOrderedMap()
    rttMonLatestHTTPOperTable.EntityData.Children.Append("rttMonLatestHTTPOperEntry", types.YChild{"RttMonLatestHTTPOperEntry", nil})
    for i := range rttMonLatestHTTPOperTable.RttMonLatestHTTPOperEntry {
        rttMonLatestHTTPOperTable.EntityData.Children.Append(types.GetSegmentPath(rttMonLatestHTTPOperTable.RttMonLatestHTTPOperEntry[i]), types.YChild{"RttMonLatestHTTPOperEntry", rttMonLatestHTTPOperTable.RttMonLatestHTTPOperEntry[i]})
    }
    rttMonLatestHTTPOperTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonLatestHTTPOperTable.EntityData.YListKeys = []string {}

    return &(rttMonLatestHTTPOperTable.EntityData)
}

// CISCORTTMONMIB_RttMonLatestHTTPOperTable_RttMonLatestHTTPOperEntry
// A list of objects that record the latest HTTP RTT
// operation. This entry is created automatically after the 
// rttMonCtrlAdminEntry is created. Also the entry is 
// automatically deleted when rttMonCtrlAdminEntry is deleted.
type CISCORTTMONMIB_RttMonLatestHTTPOperTable_RttMonLatestHTTPOperEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // Round Trip Time taken to perform HTTP operation. This value is the sum of
    // DNSRTT, TCPConnectRTT and TransactionRTT. The type is interface{} with
    // range: 0..4294967295.
    RttMonLatestHTTPOperRTT interface{}

    // Round Trip Time taken to perform DNS query within the HTTP operation. If an
    // IP Address is specified in the URL,  then DNSRTT is 0. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestHTTPOperDNSRTT interface{}

    // Round Trip Time taken to connect to the HTTP server. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestHTTPOperTCPConnectRTT interface{}

    // Round Trip Time taken to download the object specified by the URL. The type
    // is interface{} with range: 0..4294967295.
    RttMonLatestHTTPOperTransactionRTT interface{}

    // The size of the message body received as a response to the HTTP request.
    // The type is interface{} with range: 0..4294967295.
    RttMonLatestHTTPOperMessageBodyOctets interface{}

    // An application specific sense code for the completion status of the latest
    // RTT operation. The type is RttResponseSense.
    RttMonLatestHTTPOperSense interface{}

    // An sense description for the completion status of the latest RTT operation.
    // The type is string.
    RttMonLatestHTTPErrorSenseDescription interface{}
}

func (rttMonLatestHTTPOperEntry *CISCORTTMONMIB_RttMonLatestHTTPOperTable_RttMonLatestHTTPOperEntry) GetEntityData() *types.CommonEntityData {
    rttMonLatestHTTPOperEntry.EntityData.YFilter = rttMonLatestHTTPOperEntry.YFilter
    rttMonLatestHTTPOperEntry.EntityData.YangName = "rttMonLatestHTTPOperEntry"
    rttMonLatestHTTPOperEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonLatestHTTPOperEntry.EntityData.ParentYangName = "rttMonLatestHTTPOperTable"
    rttMonLatestHTTPOperEntry.EntityData.SegmentPath = "rttMonLatestHTTPOperEntry" + types.AddKeyToken(rttMonLatestHTTPOperEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex")
    rttMonLatestHTTPOperEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonLatestHTTPOperEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonLatestHTTPOperEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonLatestHTTPOperEntry.EntityData.Children = types.NewOrderedMap()
    rttMonLatestHTTPOperEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonLatestHTTPOperEntry.RttMonCtrlAdminIndex})
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonLatestHTTPOperRTT", types.YLeaf{"RttMonLatestHTTPOperRTT", rttMonLatestHTTPOperEntry.RttMonLatestHTTPOperRTT})
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonLatestHTTPOperDNSRTT", types.YLeaf{"RttMonLatestHTTPOperDNSRTT", rttMonLatestHTTPOperEntry.RttMonLatestHTTPOperDNSRTT})
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonLatestHTTPOperTCPConnectRTT", types.YLeaf{"RttMonLatestHTTPOperTCPConnectRTT", rttMonLatestHTTPOperEntry.RttMonLatestHTTPOperTCPConnectRTT})
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonLatestHTTPOperTransactionRTT", types.YLeaf{"RttMonLatestHTTPOperTransactionRTT", rttMonLatestHTTPOperEntry.RttMonLatestHTTPOperTransactionRTT})
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonLatestHTTPOperMessageBodyOctets", types.YLeaf{"RttMonLatestHTTPOperMessageBodyOctets", rttMonLatestHTTPOperEntry.RttMonLatestHTTPOperMessageBodyOctets})
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonLatestHTTPOperSense", types.YLeaf{"RttMonLatestHTTPOperSense", rttMonLatestHTTPOperEntry.RttMonLatestHTTPOperSense})
    rttMonLatestHTTPOperEntry.EntityData.Leafs.Append("rttMonLatestHTTPErrorSenseDescription", types.YLeaf{"RttMonLatestHTTPErrorSenseDescription", rttMonLatestHTTPOperEntry.RttMonLatestHTTPErrorSenseDescription})

    rttMonLatestHTTPOperEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex"}

    return &(rttMonLatestHTTPOperEntry.EntityData)
}

// CISCORTTMONMIB_RttMonLatestJitterOperTable
// A table which contains the status of latest Jitter
// operation.
type CISCORTTMONMIB_RttMonLatestJitterOperTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that record the latest Jitter operation. The type is
    // slice of
    // CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry.
    RttMonLatestJitterOperEntry []*CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry
}

func (rttMonLatestJitterOperTable *CISCORTTMONMIB_RttMonLatestJitterOperTable) GetEntityData() *types.CommonEntityData {
    rttMonLatestJitterOperTable.EntityData.YFilter = rttMonLatestJitterOperTable.YFilter
    rttMonLatestJitterOperTable.EntityData.YangName = "rttMonLatestJitterOperTable"
    rttMonLatestJitterOperTable.EntityData.BundleName = "cisco_ios_xe"
    rttMonLatestJitterOperTable.EntityData.ParentYangName = "CISCO-RTTMON-MIB"
    rttMonLatestJitterOperTable.EntityData.SegmentPath = "rttMonLatestJitterOperTable"
    rttMonLatestJitterOperTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonLatestJitterOperTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonLatestJitterOperTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonLatestJitterOperTable.EntityData.Children = types.NewOrderedMap()
    rttMonLatestJitterOperTable.EntityData.Children.Append("rttMonLatestJitterOperEntry", types.YChild{"RttMonLatestJitterOperEntry", nil})
    for i := range rttMonLatestJitterOperTable.RttMonLatestJitterOperEntry {
        rttMonLatestJitterOperTable.EntityData.Children.Append(types.GetSegmentPath(rttMonLatestJitterOperTable.RttMonLatestJitterOperEntry[i]), types.YChild{"RttMonLatestJitterOperEntry", rttMonLatestJitterOperTable.RttMonLatestJitterOperEntry[i]})
    }
    rttMonLatestJitterOperTable.EntityData.Leafs = types.NewOrderedMap()

    rttMonLatestJitterOperTable.EntityData.YListKeys = []string {}

    return &(rttMonLatestJitterOperTable.EntityData)
}

// CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry
// A list of objects that record the latest Jitter
// operation.
type CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_rttmon_mib.CISCORTTMONMIB_RttMonCtrlAdminTable_RttMonCtrlAdminEntry_RttMonCtrlAdminIndex
    RttMonCtrlAdminIndex interface{}

    // The number of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperNumOfRTT interface{}

    // The sum of Jitter RTT's that are successfully measured (low order 32 bits).
    // The high order 32 bits are stored in rttMonLatestJitterOperRTTSumHigh. The
    // type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperRTTSum interface{}

    // The sum of squares of RTT's that are successfully measured (low order 32
    // bits). The high order 32 bits are stored in
    // rttMonLatestJitterOperRTTSum2High. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperRTTSum2 interface{}

    // The minimum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperRTTMin interface{}

    // The maximum of RTT's that were successfully measured. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperRTTMax interface{}

    // The minimum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperMinOfPositivesSD interface{}

    // The maximum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperMaxOfPositivesSD interface{}

    // The sum of all positive jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperNumOfPositivesSD interface{}

    // The sum of RTT's of all positive jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperSumOfPositivesSD interface{}

    // The sum of square of RTT's of all positive jitter values from packets sent
    // from source to destination. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperSum2PositivesSD interface{}

    // The minimum of absolute values of all negative jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperMinOfNegativesSD interface{}

    // The maximum of absolute values of all negative jitter values from packets
    // sent from source to destination. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperMaxOfNegativesSD interface{}

    // The sum of number of all negative jitter values from packets sent from
    // source to destination. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperNumOfNegativesSD interface{}

    // The sum of all negative jitter values from packets sent from source to
    // destination. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperSumOfNegativesSD interface{}

    // The sum of square of RTT's of all negative jitter values from packets sent
    // from source to destination. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperSum2NegativesSD interface{}

    // The minimum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperMinOfPositivesDS interface{}

    // The maximum of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperMaxOfPositivesDS interface{}

    // The sum of number of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperNumOfPositivesDS interface{}

    // The sum of RTT's of all positive jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperSumOfPositivesDS interface{}

    // The sum of squares of RTT's of all positive jitter values from packets sent
    // from destination to source. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperSum2PositivesDS interface{}

    // The minimum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperMinOfNegativesDS interface{}

    // The maximum of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperMaxOfNegativesDS interface{}

    // The sum of number of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperNumOfNegativesDS interface{}

    // The sum of RTT's of all negative jitter values from packets sent from
    // destination to source. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperSumOfNegativesDS interface{}

    // The sum of squares of RTT's of all negative jitter values from packets sent
    // from destination to source. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperSum2NegativesDS interface{}

    // The number of packets lost when sent from source to destination. The type
    // is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperPacketLossSD interface{}

    // The number of packets lost when sent from destination to source. The type
    // is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperPacketLossDS interface{}

    // The number of packets arrived out of sequence. The type is interface{} with
    // range: 0..4294967295.
    RttMonLatestJitterOperPacketOutOfSequence interface{}

    // The number of packets that are lost for which we cannot determine the
    // direction. The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperPacketMIA interface{}

    // The number of packets that arrived after the timeout. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperPacketLateArrival interface{}

    // An application specific sense code for the completion status of the latest
    // Jitter RTT operation. The type is RttResponseSense.
    RttMonLatestJitterOperSense interface{}

    // An sense description for the completion status of the latest Jitter RTT
    // operation. The type is string.
    RttMonLatestJitterErrorSenseDescription interface{}

    // The sum of one way latency from source to destination (low order 32 bits).
    // The high order 32 bits are stored in rttMonLatestJitterOperOWSumSDHigh. The
    // type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWSumSD interface{}

    // The sum of squares of one way latency from source to destination (low order
    // 32 bits). The high order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2SDHigh. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperOWSum2SD interface{}

    // The minimum of all one way latency from source to destination. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWMinSD interface{}

    // The maximum of all one way latency from source to destination. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWMaxSD interface{}

    // The sum of one way latency from destination to source (low order 32 bits).
    // The high order 32 bits are stored in rttMonLatestJitterOperOWSumDSHigh. The
    // type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWSumDS interface{}

    // The sum of squares of one way latency from destination to source (low order
    // 32 bits). The high order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2DSHigh. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperOWSum2DS interface{}

    // The minimum of all one way latency from destination to source. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWMinDS interface{}

    // The maximum of all one way latency from destination to source. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWMaxDS interface{}

    // The number of successful one way latency measurements. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperNumOfOW interface{}

    // The MOS value for the latest jitter operation in hundreds. This value will
    // be 0 if   - rttMonEchoAdminCodecType of the operation is notApplicable   -
    // the operation is not started   - the operation is started but failed This
    // value will be 1 for packet loss of 10% or more. The type is interface{}
    // with range: 0..None | 100..500.
    RttMonLatestJitterOperMOS interface{}

    // Represents ICPIF value for the latest jitter operation.  This value will be
    // 93 for packet loss of 10% or more. The type is interface{} with range:
    // 0..2147483647.
    RttMonLatestJitterOperICPIF interface{}

    // Interarrival Jitter (RC1889) at responder. The type is interface{} with
    // range: 0..2147483647.
    RttMonLatestJitterOperIAJOut interface{}

    // Interarrival Jitter (RFC1889) at source. The type is interface{} with
    // range: 0..2147483647.
    RttMonLatestJitterOperIAJIn interface{}

    // The average of positive and negative jitter values in SD and DS direction
    // for latest operation. The type is interface{} with range: 0..2147483647.
    RttMonLatestJitterOperAvgJitter interface{}

    // The average of positive and negative jitter values from source to
    // destination for latest operation. The type is interface{} with range:
    // 0..2147483647.
    RttMonLatestJitterOperAvgSDJ interface{}

    // The average of positive and negative jitter values from destination to
    // source for latest operation. The type is interface{} with range:
    // 0..2147483647.
    RttMonLatestJitterOperAvgDSJ interface{}

    // The average latency value from source to destination. The type is
    // interface{} with range: 0..2147483647.
    RttMonLatestJitterOperOWAvgSD interface{}

    // The average latency value from destination to source. The type is
    // interface{} with range: 0..2147483647.
    RttMonLatestJitterOperOWAvgDS interface{}

    // A value of sync(1) means sender and responder was in sync with NTP. The NTP
    // sync means the total of NTP offset  on sender and responder is within
    // configured tolerance level. The type is RttMonLatestJitterOperNTPState.
    RttMonLatestJitterOperNTPState interface{}

    // The number of RTT operations that have completed with sender and responder
    // out of sync with NTP. The NTP sync means  the total of NTP offset on sender
    // and responder is within  configured tolerance level. The type is
    // interface{} with range: 0..4294967295.
    RttMonLatestJitterOperUnSyncRTs interface{}

    // The sum of Jitter RTT's that are successfully measured. (high order 32
    // bits). The low order 32 bits are stored in rttMonLatestJitterOperRTTSum.
    // The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperRTTSumHigh interface{}

    // The sum of squares of RTT's that are successfully measured (high order 32
    // bits). The low order 32 bits are stored in rttMonLatestJitterOperRTTSum2.
    // The type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperRTTSum2High interface{}

    // The sum of one way latency from source to destination (high order 32 bits).
    // The low order 32 bits are stored in rttMonLatestJitterOperOWSumSD. The type
    // is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWSumSDHigh interface{}

    // The sum of squares of one way latency from source to destination (high
    // order 32 bits). The low order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2SD. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperOWSum2SDHigh interface{}

    // The sum of one way latency from destination to source (high order 32 bits).
    // The low order 32 bits are stored  in rttMonLatestJitterOperOWSumDS. The
    // type is interface{} with range: 0..4294967295.
    RttMonLatestJitterOperOWSumDSHigh interface{}

    // The sum of squares of one way latency from destination to source (high
    // order 32 bits). The low order 32 bits are stored in
    // rttMonLatestJitterOperOWSum2DS. The type is interface{} with range:
    // 0..4294967295.
    RttMonLatestJitterOperOWSum2DSHigh interface{}
}

func (rttMonLatestJitterOperEntry *CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry) GetEntityData() *types.CommonEntityData {
    rttMonLatestJitterOperEntry.EntityData.YFilter = rttMonLatestJitterOperEntry.YFilter
    rttMonLatestJitterOperEntry.EntityData.YangName = "rttMonLatestJitterOperEntry"
    rttMonLatestJitterOperEntry.EntityData.BundleName = "cisco_ios_xe"
    rttMonLatestJitterOperEntry.EntityData.ParentYangName = "rttMonLatestJitterOperTable"
    rttMonLatestJitterOperEntry.EntityData.SegmentPath = "rttMonLatestJitterOperEntry" + types.AddKeyToken(rttMonLatestJitterOperEntry.RttMonCtrlAdminIndex, "rttMonCtrlAdminIndex")
    rttMonLatestJitterOperEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttMonLatestJitterOperEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttMonLatestJitterOperEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttMonLatestJitterOperEntry.EntityData.Children = types.NewOrderedMap()
    rttMonLatestJitterOperEntry.EntityData.Leafs = types.NewOrderedMap()
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonCtrlAdminIndex", types.YLeaf{"RttMonCtrlAdminIndex", rttMonLatestJitterOperEntry.RttMonCtrlAdminIndex})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperNumOfRTT", types.YLeaf{"RttMonLatestJitterOperNumOfRTT", rttMonLatestJitterOperEntry.RttMonLatestJitterOperNumOfRTT})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperRTTSum", types.YLeaf{"RttMonLatestJitterOperRTTSum", rttMonLatestJitterOperEntry.RttMonLatestJitterOperRTTSum})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperRTTSum2", types.YLeaf{"RttMonLatestJitterOperRTTSum2", rttMonLatestJitterOperEntry.RttMonLatestJitterOperRTTSum2})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperRTTMin", types.YLeaf{"RttMonLatestJitterOperRTTMin", rttMonLatestJitterOperEntry.RttMonLatestJitterOperRTTMin})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperRTTMax", types.YLeaf{"RttMonLatestJitterOperRTTMax", rttMonLatestJitterOperEntry.RttMonLatestJitterOperRTTMax})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMinOfPositivesSD", types.YLeaf{"RttMonLatestJitterOperMinOfPositivesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMinOfPositivesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMaxOfPositivesSD", types.YLeaf{"RttMonLatestJitterOperMaxOfPositivesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMaxOfPositivesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperNumOfPositivesSD", types.YLeaf{"RttMonLatestJitterOperNumOfPositivesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperNumOfPositivesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSumOfPositivesSD", types.YLeaf{"RttMonLatestJitterOperSumOfPositivesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSumOfPositivesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSum2PositivesSD", types.YLeaf{"RttMonLatestJitterOperSum2PositivesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSum2PositivesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMinOfNegativesSD", types.YLeaf{"RttMonLatestJitterOperMinOfNegativesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMinOfNegativesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMaxOfNegativesSD", types.YLeaf{"RttMonLatestJitterOperMaxOfNegativesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMaxOfNegativesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperNumOfNegativesSD", types.YLeaf{"RttMonLatestJitterOperNumOfNegativesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperNumOfNegativesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSumOfNegativesSD", types.YLeaf{"RttMonLatestJitterOperSumOfNegativesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSumOfNegativesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSum2NegativesSD", types.YLeaf{"RttMonLatestJitterOperSum2NegativesSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSum2NegativesSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMinOfPositivesDS", types.YLeaf{"RttMonLatestJitterOperMinOfPositivesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMinOfPositivesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMaxOfPositivesDS", types.YLeaf{"RttMonLatestJitterOperMaxOfPositivesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMaxOfPositivesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperNumOfPositivesDS", types.YLeaf{"RttMonLatestJitterOperNumOfPositivesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperNumOfPositivesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSumOfPositivesDS", types.YLeaf{"RttMonLatestJitterOperSumOfPositivesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSumOfPositivesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSum2PositivesDS", types.YLeaf{"RttMonLatestJitterOperSum2PositivesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSum2PositivesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMinOfNegativesDS", types.YLeaf{"RttMonLatestJitterOperMinOfNegativesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMinOfNegativesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMaxOfNegativesDS", types.YLeaf{"RttMonLatestJitterOperMaxOfNegativesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMaxOfNegativesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperNumOfNegativesDS", types.YLeaf{"RttMonLatestJitterOperNumOfNegativesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperNumOfNegativesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSumOfNegativesDS", types.YLeaf{"RttMonLatestJitterOperSumOfNegativesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSumOfNegativesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSum2NegativesDS", types.YLeaf{"RttMonLatestJitterOperSum2NegativesDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSum2NegativesDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperPacketLossSD", types.YLeaf{"RttMonLatestJitterOperPacketLossSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperPacketLossSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperPacketLossDS", types.YLeaf{"RttMonLatestJitterOperPacketLossDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperPacketLossDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperPacketOutOfSequence", types.YLeaf{"RttMonLatestJitterOperPacketOutOfSequence", rttMonLatestJitterOperEntry.RttMonLatestJitterOperPacketOutOfSequence})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperPacketMIA", types.YLeaf{"RttMonLatestJitterOperPacketMIA", rttMonLatestJitterOperEntry.RttMonLatestJitterOperPacketMIA})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperPacketLateArrival", types.YLeaf{"RttMonLatestJitterOperPacketLateArrival", rttMonLatestJitterOperEntry.RttMonLatestJitterOperPacketLateArrival})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperSense", types.YLeaf{"RttMonLatestJitterOperSense", rttMonLatestJitterOperEntry.RttMonLatestJitterOperSense})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterErrorSenseDescription", types.YLeaf{"RttMonLatestJitterErrorSenseDescription", rttMonLatestJitterOperEntry.RttMonLatestJitterErrorSenseDescription})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSumSD", types.YLeaf{"RttMonLatestJitterOperOWSumSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSumSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSum2SD", types.YLeaf{"RttMonLatestJitterOperOWSum2SD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSum2SD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWMinSD", types.YLeaf{"RttMonLatestJitterOperOWMinSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWMinSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWMaxSD", types.YLeaf{"RttMonLatestJitterOperOWMaxSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWMaxSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSumDS", types.YLeaf{"RttMonLatestJitterOperOWSumDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSumDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSum2DS", types.YLeaf{"RttMonLatestJitterOperOWSum2DS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSum2DS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWMinDS", types.YLeaf{"RttMonLatestJitterOperOWMinDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWMinDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWMaxDS", types.YLeaf{"RttMonLatestJitterOperOWMaxDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWMaxDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperNumOfOW", types.YLeaf{"RttMonLatestJitterOperNumOfOW", rttMonLatestJitterOperEntry.RttMonLatestJitterOperNumOfOW})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperMOS", types.YLeaf{"RttMonLatestJitterOperMOS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperMOS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperICPIF", types.YLeaf{"RttMonLatestJitterOperICPIF", rttMonLatestJitterOperEntry.RttMonLatestJitterOperICPIF})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperIAJOut", types.YLeaf{"RttMonLatestJitterOperIAJOut", rttMonLatestJitterOperEntry.RttMonLatestJitterOperIAJOut})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperIAJIn", types.YLeaf{"RttMonLatestJitterOperIAJIn", rttMonLatestJitterOperEntry.RttMonLatestJitterOperIAJIn})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperAvgJitter", types.YLeaf{"RttMonLatestJitterOperAvgJitter", rttMonLatestJitterOperEntry.RttMonLatestJitterOperAvgJitter})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperAvgSDJ", types.YLeaf{"RttMonLatestJitterOperAvgSDJ", rttMonLatestJitterOperEntry.RttMonLatestJitterOperAvgSDJ})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperAvgDSJ", types.YLeaf{"RttMonLatestJitterOperAvgDSJ", rttMonLatestJitterOperEntry.RttMonLatestJitterOperAvgDSJ})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWAvgSD", types.YLeaf{"RttMonLatestJitterOperOWAvgSD", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWAvgSD})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWAvgDS", types.YLeaf{"RttMonLatestJitterOperOWAvgDS", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWAvgDS})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperNTPState", types.YLeaf{"RttMonLatestJitterOperNTPState", rttMonLatestJitterOperEntry.RttMonLatestJitterOperNTPState})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperUnSyncRTs", types.YLeaf{"RttMonLatestJitterOperUnSyncRTs", rttMonLatestJitterOperEntry.RttMonLatestJitterOperUnSyncRTs})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperRTTSumHigh", types.YLeaf{"RttMonLatestJitterOperRTTSumHigh", rttMonLatestJitterOperEntry.RttMonLatestJitterOperRTTSumHigh})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperRTTSum2High", types.YLeaf{"RttMonLatestJitterOperRTTSum2High", rttMonLatestJitterOperEntry.RttMonLatestJitterOperRTTSum2High})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSumSDHigh", types.YLeaf{"RttMonLatestJitterOperOWSumSDHigh", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSumSDHigh})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSum2SDHigh", types.YLeaf{"RttMonLatestJitterOperOWSum2SDHigh", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSum2SDHigh})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSumDSHigh", types.YLeaf{"RttMonLatestJitterOperOWSumDSHigh", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSumDSHigh})
    rttMonLatestJitterOperEntry.EntityData.Leafs.Append("rttMonLatestJitterOperOWSum2DSHigh", types.YLeaf{"RttMonLatestJitterOperOWSum2DSHigh", rttMonLatestJitterOperEntry.RttMonLatestJitterOperOWSum2DSHigh})

    rttMonLatestJitterOperEntry.EntityData.YListKeys = []string {"RttMonCtrlAdminIndex"}

    return &(rttMonLatestJitterOperEntry.EntityData)
}

// CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry_RttMonLatestJitterOperNTPState represents on sender and responder is within configured tolerance level.
type CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry_RttMonLatestJitterOperNTPState string

const (
    CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry_RttMonLatestJitterOperNTPState_sync CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry_RttMonLatestJitterOperNTPState = "sync"

    CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry_RttMonLatestJitterOperNTPState_outOfSync CISCORTTMONMIB_RttMonLatestJitterOperTable_RttMonLatestJitterOperEntry_RttMonLatestJitterOperNTPState = "outOfSync"
)

