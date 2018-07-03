// Cisco Express Forwarding (CEF) describes a high speed 
// switching mechanism that a router uses to forward packets
// from the inbound to the outbound interface. 
// 
// CEF uses two sets of data structures
// or tables, which it stores in router memory:
// 
// Forwarding information base (FIB) - Describes a database
// of information used to make forwarding decisions. It is 
// conceptually similar to a routing table or route-cache, 
// although its implementation is different.
// 
// Adjacency - Two nodes in the network are said to be 
// adjacent if they can reach each other via a single hop 
// across a link layer.           
// 
// CEF path is a valid route to reach to a destination 
// IP prefix. Multiple paths may exist out of a router to the 
// same destination prefix. CEF Load balancing capability 
// share the traffic to the destination IP prefix over all 
// the active paths. 
// 
// After obtaining the prefix in the CEF table with the
// longest match, output forwarding follows the chain of 
// forwarding elements. 
// 
// Forwarding element (FE) may process the packet, forward 
// the packet, drop or punt the packet or it may also
// pass the packet to the next forwarding element in the 
// chain for further processing.
// 
// Forwarding Elements are of various types
// but this MIB only represents the forwarding elements of
// adjacency and label types. Hence a forwarding element 
// chain will be represented as a list of labels and
// adjacency. The adjacency may point to a forwarding element
// list again, if it is not the last forwarding element in this
// chain. 
// 
// For the simplest IP forwarding case, the prefix entry will 
// point at an adjacency forwarding element.
// The IP adjacency processing function will apply the output
// features, add the encapsulation (performing any required 
// fixups), and may send the packet out.
// 
// If loadbalancing is configured, the prefix entry will point 
// to lists of forwarding elements. One of these lists will be 
// selected to forward the packet. 
// 
// Each forwarding element list dictates which of a set of 
// possible packet transformations to apply on the way to 
// the same neighbour. 
// 
// The following diagram represents relationship
// between three of the core tables in this MIB module.
// 
//  cefPrefixTable             cefFESelectionTable
// 
//  +---------------+  points           +--------------+   
//  |   |     |     |  a set     +----> |   |   |   |  | 
//  |---------------|  of FE     |      |--------------|   
//  |   |     |     |  Selection |      |   |   |   |  |
//  |---------------|  Entries   |      |--------------|    
//  |   |     |     |------------+      |              |<----+ 
//  |---------------|                   |--------------|     |
//  |               |    +--------------|   |   |   |  |     |
//  +---------------+    |              +--------------+     |
//                       |                                   |
//                 points to an                              |
//                 adjacency entry                           |
//                       |                                   |
//                       |   cefAdjTable                     |
//                       |  +---------------+  may point     |
//                       +->|   |     |     |  to a set      |
//                          |---------------|  of FE         |
//                          |   |     |     |  Selection     |
//                          |---------------|  Entries       | 
//                          |   |     |     |----------------+
//                          |---------------| 
//                          |               | 
//                          +---------------+ 
// 
// Some of the Cisco series routers (e.g. 7500 & 12000) 
// support distributed CEF (dCEF), in which the line cards 
// (LCs) make the packet forwarding decisions using locally 
// stored copies of the same Forwarding information base (FIB)
// and adjacency tables as the Routing Processor (RP).
//           
// Inter-Process Communication (IPC) is the protocol used 
// by routers that support distributed packet forwarding. 
// CEF updates are encoded as external Data Representation 
// (XDR) information elements inside IPC messages. 
//          
// This MIB reflects the distributed nature of CEF, e.g. CEF
// has different instances running on the RP and the line cards.
// 
// There may be instances of inconsistency between the
// CEF forwarding databases(i.e between CEF forwarding 
// database on line cards and the CEF forwarding database
// on the RP). CEF consistency checkers (CC) detects 
// this inconsistency.
// 
// When two databases are compared by a consistency checker, 
// a set of records from the first (master) database is 
// looked up in the second (slave).
// 
// There are two types of consistency checkers, 
// active and passive. Active consistency checkers 
// are invoked in response to some stimulus, i.e. 
// when a packet cannot be forwarded because the 
// prefix is not in the forwarding table or 
// in response to a Management Station request.
// 
// Passive consistency checkers operate in the background, 
// scanning portions of the databases on a periodic basis.
// 
// The full-scan checkers are active consistency checkers
// which are invoked in response to a Management Station
// Request.
// 
// If 64-bit counter objects in this MIB are supported,
// then their associated 32-bit counter objects 
// must also be supported. The 32-bit counters will
// report the low 32-bits of the associated 64-bit 
// counter count (e.g., cefPrefixPkts will report the 
// least significant 32 bits of cefPrefixHCPkts).
// The same rule should be applied for the 64-bit gauge
// objects and their assocaited 32-bit gauge objects.
// 
// If 64-bit counters in this MIB are not supported,
// then an agent MUST NOT instantiate the corresponding
// objects with an incorrect value; rather, it MUST 
// respond with the appropriate error/exception 
// condition (e.g., noSuchInstance or noSuchName). 
// 
// Counters related to CEF accounting (e.g.,
// cefPrefixPkts) MUST NOT be instantiated if the
// corresponding accounting method has been disabled.  
//  
// This MIB allows configuration and monitoring of CEF 
// related objects.
package cisco_cef_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cef_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CEF-MIB CISCO-CEF-MIB}", reflect.TypeOf(CISCOCEFMIB{}))
    ydk.RegisterEntity("CISCO-CEF-MIB:CISCO-CEF-MIB", reflect.TypeOf(CISCOCEFMIB{}))
}

// CISCOCEFMIB
type CISCOCEFMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CefFIB CISCOCEFMIB_CefFIB

    
    CefCC CISCOCEFMIB_CefCC

    
    CefNotifCntl CISCOCEFMIB_CefNotifCntl

    // This table contains the summary information for the cefPrefixTable.
    CefFIBSummaryTable CISCOCEFMIB_CefFIBSummaryTable

    // A list of CEF forwarding prefixes.
    CefPrefixTable CISCOCEFMIB_CefPrefixTable

    // A table of Longest Match Prefix Query requests.  Generator application
    // should utilize the cefLMPrefixSpinLock to try to avoid collisions. See
    // DESCRIPTION clause of cefLMPrefixSpinLock.
    CefLMPrefixTable CISCOCEFMIB_CefLMPrefixTable

    // CEF prefix path is a valid route to reach to a  destination IP prefix.
    // Multiple paths may exist  out of a router to the same destination prefix. 
    // This table specify lists of CEF paths.
    CefPathTable CISCOCEFMIB_CefPathTable

    // This table contains the summary information for the cefAdjTable.
    CefAdjSummaryTable CISCOCEFMIB_CefAdjSummaryTable

    // A list of CEF adjacencies.
    CefAdjTable CISCOCEFMIB_CefAdjTable

    // A list of forwarding element selection entries.
    CefFESelectionTable CISCOCEFMIB_CefFESelectionTable

    // This table contains global config parameter  of CEF on the Managed device.
    CefCfgTable CISCOCEFMIB_CefCfgTable

    // This table contains global resource  information of CEF on the Managed
    // device.
    CefResourceTable CISCOCEFMIB_CefResourceTable

    // This Table contains interface specific information of CEF on the Managed
    // device.
    CefIntTable CISCOCEFMIB_CefIntTable

    // Entity acting as RP (Routing Processor) keeps the CEF states for the line
    // card entities and communicates with the line card entities using XDR. This
    // Table contains the CEF information  related to peer entities on the managed
    // device.
    CefPeerTable CISCOCEFMIB_CefPeerTable

    // Entity acting as RP (Routing Processor) keep the CEF FIB states for the
    // line card entities and communicate with the line card entities using XDR.
    // This Table contains the CEF FIB State  related to peer entities on the
    // managed device.
    CefPeerFIBTable CISCOCEFMIB_CefPeerFIBTable

    // This table contains CEF consistency checker (CC) global parameters for the
    // managed device.
    CefCCGlobalTable CISCOCEFMIB_CefCCGlobalTable

    // This table contains CEF consistency checker types specific parameters on
    // the managed device.  All detected inconsistency are signaled to the
    // Management Station via cefInconsistencyDetection notification.
    CefCCTypeTable CISCOCEFMIB_CefCCTypeTable

    // This table contains CEF inconsistency records.
    CefInconsistencyRecordTable CISCOCEFMIB_CefInconsistencyRecordTable

    // This table specifies the CEF stats based on the Prefix Length.
    CefStatsPrefixLenTable CISCOCEFMIB_CefStatsPrefixLenTable

    // This table specifies the CEF switch stats.
    CefSwitchingStatsTable CISCOCEFMIB_CefSwitchingStatsTable
}

func (cISCOCEFMIB *CISCOCEFMIB) GetEntityData() *types.CommonEntityData {
    cISCOCEFMIB.EntityData.YFilter = cISCOCEFMIB.YFilter
    cISCOCEFMIB.EntityData.YangName = "CISCO-CEF-MIB"
    cISCOCEFMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCEFMIB.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cISCOCEFMIB.EntityData.SegmentPath = "CISCO-CEF-MIB:CISCO-CEF-MIB"
    cISCOCEFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCEFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCEFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCEFMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCEFMIB.EntityData.Children.Append("cefFIB", types.YChild{"CefFIB", &cISCOCEFMIB.CefFIB})
    cISCOCEFMIB.EntityData.Children.Append("cefCC", types.YChild{"CefCC", &cISCOCEFMIB.CefCC})
    cISCOCEFMIB.EntityData.Children.Append("cefNotifCntl", types.YChild{"CefNotifCntl", &cISCOCEFMIB.CefNotifCntl})
    cISCOCEFMIB.EntityData.Children.Append("cefFIBSummaryTable", types.YChild{"CefFIBSummaryTable", &cISCOCEFMIB.CefFIBSummaryTable})
    cISCOCEFMIB.EntityData.Children.Append("cefPrefixTable", types.YChild{"CefPrefixTable", &cISCOCEFMIB.CefPrefixTable})
    cISCOCEFMIB.EntityData.Children.Append("cefLMPrefixTable", types.YChild{"CefLMPrefixTable", &cISCOCEFMIB.CefLMPrefixTable})
    cISCOCEFMIB.EntityData.Children.Append("cefPathTable", types.YChild{"CefPathTable", &cISCOCEFMIB.CefPathTable})
    cISCOCEFMIB.EntityData.Children.Append("cefAdjSummaryTable", types.YChild{"CefAdjSummaryTable", &cISCOCEFMIB.CefAdjSummaryTable})
    cISCOCEFMIB.EntityData.Children.Append("cefAdjTable", types.YChild{"CefAdjTable", &cISCOCEFMIB.CefAdjTable})
    cISCOCEFMIB.EntityData.Children.Append("cefFESelectionTable", types.YChild{"CefFESelectionTable", &cISCOCEFMIB.CefFESelectionTable})
    cISCOCEFMIB.EntityData.Children.Append("cefCfgTable", types.YChild{"CefCfgTable", &cISCOCEFMIB.CefCfgTable})
    cISCOCEFMIB.EntityData.Children.Append("cefResourceTable", types.YChild{"CefResourceTable", &cISCOCEFMIB.CefResourceTable})
    cISCOCEFMIB.EntityData.Children.Append("cefIntTable", types.YChild{"CefIntTable", &cISCOCEFMIB.CefIntTable})
    cISCOCEFMIB.EntityData.Children.Append("cefPeerTable", types.YChild{"CefPeerTable", &cISCOCEFMIB.CefPeerTable})
    cISCOCEFMIB.EntityData.Children.Append("cefPeerFIBTable", types.YChild{"CefPeerFIBTable", &cISCOCEFMIB.CefPeerFIBTable})
    cISCOCEFMIB.EntityData.Children.Append("cefCCGlobalTable", types.YChild{"CefCCGlobalTable", &cISCOCEFMIB.CefCCGlobalTable})
    cISCOCEFMIB.EntityData.Children.Append("cefCCTypeTable", types.YChild{"CefCCTypeTable", &cISCOCEFMIB.CefCCTypeTable})
    cISCOCEFMIB.EntityData.Children.Append("cefInconsistencyRecordTable", types.YChild{"CefInconsistencyRecordTable", &cISCOCEFMIB.CefInconsistencyRecordTable})
    cISCOCEFMIB.EntityData.Children.Append("cefStatsPrefixLenTable", types.YChild{"CefStatsPrefixLenTable", &cISCOCEFMIB.CefStatsPrefixLenTable})
    cISCOCEFMIB.EntityData.Children.Append("cefSwitchingStatsTable", types.YChild{"CefSwitchingStatsTable", &cISCOCEFMIB.CefSwitchingStatsTable})
    cISCOCEFMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCEFMIB.EntityData.YListKeys = []string {}

    return &(cISCOCEFMIB.EntityData)
}

// CISCOCEFMIB_CefFIB
type CISCOCEFMIB_CefFIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An advisory lock used to allow cooperating SNMP Command Generator
    // applications to coordinate their use of the Set operation in creating
    // Longest Match Prefix Entries in cefLMPrefixTable.  When creating a new
    // longest prefix match entry, the value of cefLMPrefixSpinLock should be
    // retrieved.   The destination address should be determined to be unique by
    // the SNMP Command Generator application by consulting the cefLMPrefixTable.
    // Finally, the longest  prefix entry may be created (Set), including the
    // advisory lock.         If another SNMP Command Generator application has
    // altered the longest prefix entry in the meantime,  then the spin lock's
    // value will have changed,  and so this creation will fail because it will
    // specify the wrong value for the spin lock.  Since this is an advisory lock,
    // the use of this lock is not enforced, but not using this lock may lead to
    // conflict with the another SNMP command responder  application which may
    // also be acting on the cefLMPrefixTable. The type is interface{} with range:
    // 0..2147483647.
    CefLMPrefixSpinLock interface{}
}

func (cefFIB *CISCOCEFMIB_CefFIB) GetEntityData() *types.CommonEntityData {
    cefFIB.EntityData.YFilter = cefFIB.YFilter
    cefFIB.EntityData.YangName = "cefFIB"
    cefFIB.EntityData.BundleName = "cisco_ios_xe"
    cefFIB.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefFIB.EntityData.SegmentPath = "cefFIB"
    cefFIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefFIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefFIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefFIB.EntityData.Children = types.NewOrderedMap()
    cefFIB.EntityData.Leafs = types.NewOrderedMap()
    cefFIB.EntityData.Leafs.Append("cefLMPrefixSpinLock", types.YLeaf{"CefLMPrefixSpinLock", cefFIB.CefLMPrefixSpinLock})

    cefFIB.EntityData.YListKeys = []string {}

    return &(cefFIB.EntityData)
}

// CISCOCEFMIB_CefCC
type CISCOCEFMIB_CefCC struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time an inconsistency is detecetd. The type
    // is interface{} with range: 0..4294967295.
    EntLastInconsistencyDetectTime interface{}

    // Setting the value of this object to ccActionStart(1) will reset all the
    // active consistency checkers.  The Management station should poll the 
    // cefInconsistencyResetStatus object to get the  state of inconsistency reset
    // operation.  This operation once started, cannot be aborted. Hence, the
    // value of this object cannot be set to ccActionAbort(2).  The value of this
    // object can't be set to ccActionStart(1),  if the value of object
    // cefInconsistencyResetStatus is ccStatusRunning(2). The type is CefCCAction.
    CefInconsistencyReset interface{}

    // Indicates the status of the consistency reset request. The type is
    // CefCCStatus.
    CefInconsistencyResetStatus interface{}
}

func (cefCC *CISCOCEFMIB_CefCC) GetEntityData() *types.CommonEntityData {
    cefCC.EntityData.YFilter = cefCC.YFilter
    cefCC.EntityData.YangName = "cefCC"
    cefCC.EntityData.BundleName = "cisco_ios_xe"
    cefCC.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefCC.EntityData.SegmentPath = "cefCC"
    cefCC.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefCC.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefCC.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefCC.EntityData.Children = types.NewOrderedMap()
    cefCC.EntityData.Leafs = types.NewOrderedMap()
    cefCC.EntityData.Leafs.Append("entLastInconsistencyDetectTime", types.YLeaf{"EntLastInconsistencyDetectTime", cefCC.EntLastInconsistencyDetectTime})
    cefCC.EntityData.Leafs.Append("cefInconsistencyReset", types.YLeaf{"CefInconsistencyReset", cefCC.CefInconsistencyReset})
    cefCC.EntityData.Leafs.Append("cefInconsistencyResetStatus", types.YLeaf{"CefInconsistencyResetStatus", cefCC.CefInconsistencyResetStatus})

    cefCC.EntityData.YListKeys = []string {}

    return &(cefCC.EntityData)
}

// CISCOCEFMIB_CefNotifCntl
type CISCOCEFMIB_CefNotifCntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether or not a notification should be generated on the
    // detection of CEF resource Failure. The type is bool.
    CefResourceFailureNotifEnable interface{}

    // Indicates whether or not a notification should be generated on the
    // detection of CEF peer state change. The type is bool.
    CefPeerStateChangeNotifEnable interface{}

    // Indicates whether or not a notification should be generated on the
    // detection of CEF FIB peer state change. The type is bool.
    CefPeerFIBStateChangeNotifEnable interface{}

    // This object controls the generation of the cefInconsistencyDetection
    // notification.  If this object has a value of zero, then the throttle
    // control is disabled.  If this object has a non-zero value, then the agent
    // must not generate more than one  cefInconsistencyDetection
    // 'notification-event' in the  indicated period, where a 'notification-event'
    // is the transmission of a single trap or inform PDU to a list of
    // notification destinations.  If additional inconsistency is detected within
    // the  throttling period, then notification-events for these inconsistencies
    // should be suppressed by the agent until the current throttling period
    // expires.  At the end of a throttling period, one notification-event should
    // be generated if any inconsistency was detected since the start of the 
    // throttling period. In such a case,  another throttling period is started
    // right away.  An NMS should periodically poll cefInconsistencyRecordTable to
    // detect any missed cefInconsistencyDetection notification-events, e.g., due
    // to throttling or transmission loss.   If cefNotifThrottlingInterval
    // notification generation is enabled, the suggested default throttling period
    // is 60 seconds, but generation of the cefInconsistencyDetection notification
    // should be disabled by default.  If the agent is capable of storing
    // non-volatile configuration, then the value of this object must be restored
    // after a re-initialization of the management system.  The actual
    // transmission of notifications is controlled via the MIB modules in RFC
    // 3413. The type is interface{} with range: 0..3600. Units are seconds.
    CefNotifThrottlingInterval interface{}

    // Indicates whether cefInconsistencyDetection notification should be
    // generated for this managed device. The type is bool.
    CefInconsistencyNotifEnable interface{}
}

func (cefNotifCntl *CISCOCEFMIB_CefNotifCntl) GetEntityData() *types.CommonEntityData {
    cefNotifCntl.EntityData.YFilter = cefNotifCntl.YFilter
    cefNotifCntl.EntityData.YangName = "cefNotifCntl"
    cefNotifCntl.EntityData.BundleName = "cisco_ios_xe"
    cefNotifCntl.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefNotifCntl.EntityData.SegmentPath = "cefNotifCntl"
    cefNotifCntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefNotifCntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefNotifCntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefNotifCntl.EntityData.Children = types.NewOrderedMap()
    cefNotifCntl.EntityData.Leafs = types.NewOrderedMap()
    cefNotifCntl.EntityData.Leafs.Append("cefResourceFailureNotifEnable", types.YLeaf{"CefResourceFailureNotifEnable", cefNotifCntl.CefResourceFailureNotifEnable})
    cefNotifCntl.EntityData.Leafs.Append("cefPeerStateChangeNotifEnable", types.YLeaf{"CefPeerStateChangeNotifEnable", cefNotifCntl.CefPeerStateChangeNotifEnable})
    cefNotifCntl.EntityData.Leafs.Append("cefPeerFIBStateChangeNotifEnable", types.YLeaf{"CefPeerFIBStateChangeNotifEnable", cefNotifCntl.CefPeerFIBStateChangeNotifEnable})
    cefNotifCntl.EntityData.Leafs.Append("cefNotifThrottlingInterval", types.YLeaf{"CefNotifThrottlingInterval", cefNotifCntl.CefNotifThrottlingInterval})
    cefNotifCntl.EntityData.Leafs.Append("cefInconsistencyNotifEnable", types.YLeaf{"CefInconsistencyNotifEnable", cefNotifCntl.CefInconsistencyNotifEnable})

    cefNotifCntl.EntityData.YListKeys = []string {}

    return &(cefNotifCntl.EntityData)
}

// CISCOCEFMIB_CefFIBSummaryTable
// This table contains the summary information
// for the cefPrefixTable.
type CISCOCEFMIB_CefFIBSummaryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the FIB
    // summary related attributes for the managed entity.  A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device.  entPhysicalIndex is also an index for this table which
    // represents entities of 'module' entPhysicalClass which are capable of
    // running CEF. The type is slice of
    // CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry.
    CefFIBSummaryEntry []*CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry
}

func (cefFIBSummaryTable *CISCOCEFMIB_CefFIBSummaryTable) GetEntityData() *types.CommonEntityData {
    cefFIBSummaryTable.EntityData.YFilter = cefFIBSummaryTable.YFilter
    cefFIBSummaryTable.EntityData.YangName = "cefFIBSummaryTable"
    cefFIBSummaryTable.EntityData.BundleName = "cisco_ios_xe"
    cefFIBSummaryTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefFIBSummaryTable.EntityData.SegmentPath = "cefFIBSummaryTable"
    cefFIBSummaryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefFIBSummaryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefFIBSummaryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefFIBSummaryTable.EntityData.Children = types.NewOrderedMap()
    cefFIBSummaryTable.EntityData.Children.Append("cefFIBSummaryEntry", types.YChild{"CefFIBSummaryEntry", nil})
    for i := range cefFIBSummaryTable.CefFIBSummaryEntry {
        cefFIBSummaryTable.EntityData.Children.Append(types.GetSegmentPath(cefFIBSummaryTable.CefFIBSummaryEntry[i]), types.YChild{"CefFIBSummaryEntry", cefFIBSummaryTable.CefFIBSummaryEntry[i]})
    }
    cefFIBSummaryTable.EntityData.Leafs = types.NewOrderedMap()

    cefFIBSummaryTable.EntityData.YListKeys = []string {}

    return &(cefFIBSummaryTable.EntityData)
}

// CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry
// If CEF is enabled on the Managed device,
// each entry contains the FIB summary related
// attributes for the managed entity.
// 
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The version of IP forwarding. The type is
    // CefIpVersion.
    CefFIBIpVersion interface{}

    // Total number of forwarding Prefixes in FIB for the IP version specified by
    // cefFIBIpVersion object. The type is interface{} with range: 0..4294967295.
    CefFIBSummaryFwdPrefixes interface{}
}

func (cefFIBSummaryEntry *CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry) GetEntityData() *types.CommonEntityData {
    cefFIBSummaryEntry.EntityData.YFilter = cefFIBSummaryEntry.YFilter
    cefFIBSummaryEntry.EntityData.YangName = "cefFIBSummaryEntry"
    cefFIBSummaryEntry.EntityData.BundleName = "cisco_ios_xe"
    cefFIBSummaryEntry.EntityData.ParentYangName = "cefFIBSummaryTable"
    cefFIBSummaryEntry.EntityData.SegmentPath = "cefFIBSummaryEntry" + types.AddKeyToken(cefFIBSummaryEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefFIBSummaryEntry.CefFIBIpVersion, "cefFIBIpVersion")
    cefFIBSummaryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefFIBSummaryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefFIBSummaryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefFIBSummaryEntry.EntityData.Children = types.NewOrderedMap()
    cefFIBSummaryEntry.EntityData.Leafs = types.NewOrderedMap()
    cefFIBSummaryEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefFIBSummaryEntry.EntPhysicalIndex})
    cefFIBSummaryEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefFIBSummaryEntry.CefFIBIpVersion})
    cefFIBSummaryEntry.EntityData.Leafs.Append("cefFIBSummaryFwdPrefixes", types.YLeaf{"CefFIBSummaryFwdPrefixes", cefFIBSummaryEntry.CefFIBSummaryFwdPrefixes})

    cefFIBSummaryEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefFIBIpVersion"}

    return &(cefFIBSummaryEntry.EntityData)
}

// CISCOCEFMIB_CefPrefixTable
// A list of CEF forwarding prefixes.
type CISCOCEFMIB_CefPrefixTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the forwarding
    // prefix attributes.   CEF prefix based non-recursive stats are maintained in
    // internal and external buckets (depending upon the  value of
    // cefIntNonrecursiveAccouting object in the  CefIntEntry).  entPhysicalIndex
    // is also an index for this table which represents entities of 'module'
    // entPhysicalClass which are capable of running CEF. The type is slice of
    // CISCOCEFMIB_CefPrefixTable_CefPrefixEntry.
    CefPrefixEntry []*CISCOCEFMIB_CefPrefixTable_CefPrefixEntry
}

func (cefPrefixTable *CISCOCEFMIB_CefPrefixTable) GetEntityData() *types.CommonEntityData {
    cefPrefixTable.EntityData.YFilter = cefPrefixTable.YFilter
    cefPrefixTable.EntityData.YangName = "cefPrefixTable"
    cefPrefixTable.EntityData.BundleName = "cisco_ios_xe"
    cefPrefixTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefPrefixTable.EntityData.SegmentPath = "cefPrefixTable"
    cefPrefixTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPrefixTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPrefixTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPrefixTable.EntityData.Children = types.NewOrderedMap()
    cefPrefixTable.EntityData.Children.Append("cefPrefixEntry", types.YChild{"CefPrefixEntry", nil})
    for i := range cefPrefixTable.CefPrefixEntry {
        cefPrefixTable.EntityData.Children.Append(types.GetSegmentPath(cefPrefixTable.CefPrefixEntry[i]), types.YChild{"CefPrefixEntry", cefPrefixTable.CefPrefixEntry[i]})
    }
    cefPrefixTable.EntityData.Leafs = types.NewOrderedMap()

    cefPrefixTable.EntityData.YListKeys = []string {}

    return &(cefPrefixTable.EntityData)
}

// CISCOCEFMIB_CefPrefixTable_CefPrefixEntry
// If CEF is enabled on the Managed device,
// each entry contains the forwarding 
// prefix attributes. 
// 
// CEF prefix based non-recursive stats are maintained
// in internal and external buckets (depending upon the 
// value of cefIntNonrecursiveAccouting object in the 
// CefIntEntry).
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefPrefixTable_CefPrefixEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The Network Prefix Type. This object specifies the
    // address type used for cefPrefixAddr.  Prefix entries are only valid for the
    // address type of ipv4(1) and ipv6(2). The type is InetAddressType.
    CefPrefixType interface{}

    // This attribute is a key. The Network Prefix Address. The type of this
    // address is determined by the value of the cefPrefixType object. This object
    // is a Prefix Address containing the  prefix with length specified by
    // cefPrefixLen.  Any bits beyond the length specified by cefPrefixLen are
    // zeroed. The type is string with length: 0..255.
    CefPrefixAddr interface{}

    // This attribute is a key. Length in bits of the FIB Address prefix. The type
    // is interface{} with range: 0..2040.
    CefPrefixLen interface{}

    // This object indicates the associated forwarding element selection entries
    // in cefFESelectionTable. The value of this object is index value
    // (cefFESelectionName) of cefFESelectionTable. The type is string.
    CefPrefixForwardingInfo interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of packets switched to
    // this prefix. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    CefPrefixPkts interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of packets switched to
    // this prefix.   This object is a 64-bit version of  cefPrefixPkts. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    CefPrefixHCPkts interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of bytes switched to
    // this prefix. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    CefPrefixBytes interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of bytes switched to
    // this prefix.  This object is a 64-bit version of  cefPrefixBytes. The type
    // is interface{} with range: 0..18446744073709551615. Units are bytes.
    CefPrefixHCBytes interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the internal bucket switched using this prefix.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    CefPrefixInternalNRPkts interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the internal bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixInternalNRPkts. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    CefPrefixInternalNRHCPkts interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the internal bucket switched using this prefix. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    CefPrefixInternalNRBytes interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the internal bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixInternalNRBytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    CefPrefixInternalNRHCBytes interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the external bucket switched using this prefix.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    CefPrefixExternalNRPkts interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the external bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixExternalNRPkts. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    CefPrefixExternalNRHCPkts interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the external bucket switched using this prefix. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    CefPrefixExternalNRBytes interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the external bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixExternalNRBytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    CefPrefixExternalNRHCBytes interface{}
}

func (cefPrefixEntry *CISCOCEFMIB_CefPrefixTable_CefPrefixEntry) GetEntityData() *types.CommonEntityData {
    cefPrefixEntry.EntityData.YFilter = cefPrefixEntry.YFilter
    cefPrefixEntry.EntityData.YangName = "cefPrefixEntry"
    cefPrefixEntry.EntityData.BundleName = "cisco_ios_xe"
    cefPrefixEntry.EntityData.ParentYangName = "cefPrefixTable"
    cefPrefixEntry.EntityData.SegmentPath = "cefPrefixEntry" + types.AddKeyToken(cefPrefixEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefPrefixEntry.CefPrefixType, "cefPrefixType") + types.AddKeyToken(cefPrefixEntry.CefPrefixAddr, "cefPrefixAddr") + types.AddKeyToken(cefPrefixEntry.CefPrefixLen, "cefPrefixLen")
    cefPrefixEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPrefixEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPrefixEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPrefixEntry.EntityData.Children = types.NewOrderedMap()
    cefPrefixEntry.EntityData.Leafs = types.NewOrderedMap()
    cefPrefixEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefPrefixEntry.EntPhysicalIndex})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixType", types.YLeaf{"CefPrefixType", cefPrefixEntry.CefPrefixType})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixAddr", types.YLeaf{"CefPrefixAddr", cefPrefixEntry.CefPrefixAddr})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixLen", types.YLeaf{"CefPrefixLen", cefPrefixEntry.CefPrefixLen})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixForwardingInfo", types.YLeaf{"CefPrefixForwardingInfo", cefPrefixEntry.CefPrefixForwardingInfo})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixPkts", types.YLeaf{"CefPrefixPkts", cefPrefixEntry.CefPrefixPkts})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixHCPkts", types.YLeaf{"CefPrefixHCPkts", cefPrefixEntry.CefPrefixHCPkts})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixBytes", types.YLeaf{"CefPrefixBytes", cefPrefixEntry.CefPrefixBytes})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixHCBytes", types.YLeaf{"CefPrefixHCBytes", cefPrefixEntry.CefPrefixHCBytes})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixInternalNRPkts", types.YLeaf{"CefPrefixInternalNRPkts", cefPrefixEntry.CefPrefixInternalNRPkts})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixInternalNRHCPkts", types.YLeaf{"CefPrefixInternalNRHCPkts", cefPrefixEntry.CefPrefixInternalNRHCPkts})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixInternalNRBytes", types.YLeaf{"CefPrefixInternalNRBytes", cefPrefixEntry.CefPrefixInternalNRBytes})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixInternalNRHCBytes", types.YLeaf{"CefPrefixInternalNRHCBytes", cefPrefixEntry.CefPrefixInternalNRHCBytes})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixExternalNRPkts", types.YLeaf{"CefPrefixExternalNRPkts", cefPrefixEntry.CefPrefixExternalNRPkts})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixExternalNRHCPkts", types.YLeaf{"CefPrefixExternalNRHCPkts", cefPrefixEntry.CefPrefixExternalNRHCPkts})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixExternalNRBytes", types.YLeaf{"CefPrefixExternalNRBytes", cefPrefixEntry.CefPrefixExternalNRBytes})
    cefPrefixEntry.EntityData.Leafs.Append("cefPrefixExternalNRHCBytes", types.YLeaf{"CefPrefixExternalNRHCBytes", cefPrefixEntry.CefPrefixExternalNRHCBytes})

    cefPrefixEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefPrefixType", "CefPrefixAddr", "CefPrefixLen"}

    return &(cefPrefixEntry.EntityData)
}

// CISCOCEFMIB_CefLMPrefixTable
// A table of Longest Match Prefix Query requests.
// 
// Generator application should utilize the
// cefLMPrefixSpinLock to try to avoid collisions.
// See DESCRIPTION clause of cefLMPrefixSpinLock.
type CISCOCEFMIB_CefLMPrefixTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the managed device, then each entry represents a
    // longest Match Prefix request.  A management station wishing to get the
    // longest Match prefix for a given destination address should create the
    // associate instance of the row status. The row status should be set to
    // active(1) to initiate the request. Note that  this entire procedure may be
    // initiated via a  single set request which specifies a row status  of
    // createAndGo(4).  Once the request completes, the management station  should
    // retrieve the values of the objects of  interest, and should then delete the
    // entry.  In order  to prevent old entries from clogging the table,  entries
    // will be aged out, but an entry will never be  deleted within 5 minutes of
    // completion. Entries are lost after an agent restart.  I.e. to find out the
    // longest prefix match for  destination address of A.B.C.D on entity whose
    // entityPhysicalIndex is 1, the Management station will create an entry in
    // cefLMPrefixTable with
    // cefLMPrefixRowStatus.1(entPhysicalIndex).1(ipv4).A.B.C.D set to
    // createAndGo(4). Management Station may query the value of objects
    // cefLMPrefix and cefLMPrefixLen to find out the corresponding prefix entry
    // from the cefPrefixTable once the value of cefLMPrefixState is set to
    // matchFound(2).  entPhysicalIndex is also an index for this table which
    // represents entities of 'module' entPhysicalClass which are capable of
    // running CEF. The type is slice of
    // CISCOCEFMIB_CefLMPrefixTable_CefLMPrefixEntry.
    CefLMPrefixEntry []*CISCOCEFMIB_CefLMPrefixTable_CefLMPrefixEntry
}

func (cefLMPrefixTable *CISCOCEFMIB_CefLMPrefixTable) GetEntityData() *types.CommonEntityData {
    cefLMPrefixTable.EntityData.YFilter = cefLMPrefixTable.YFilter
    cefLMPrefixTable.EntityData.YangName = "cefLMPrefixTable"
    cefLMPrefixTable.EntityData.BundleName = "cisco_ios_xe"
    cefLMPrefixTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefLMPrefixTable.EntityData.SegmentPath = "cefLMPrefixTable"
    cefLMPrefixTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefLMPrefixTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefLMPrefixTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefLMPrefixTable.EntityData.Children = types.NewOrderedMap()
    cefLMPrefixTable.EntityData.Children.Append("cefLMPrefixEntry", types.YChild{"CefLMPrefixEntry", nil})
    for i := range cefLMPrefixTable.CefLMPrefixEntry {
        cefLMPrefixTable.EntityData.Children.Append(types.GetSegmentPath(cefLMPrefixTable.CefLMPrefixEntry[i]), types.YChild{"CefLMPrefixEntry", cefLMPrefixTable.CefLMPrefixEntry[i]})
    }
    cefLMPrefixTable.EntityData.Leafs = types.NewOrderedMap()

    cefLMPrefixTable.EntityData.YListKeys = []string {}

    return &(cefLMPrefixTable.EntityData)
}

// CISCOCEFMIB_CefLMPrefixTable_CefLMPrefixEntry
// If CEF is enabled on the managed device, then each
// entry represents a longest Match Prefix request.
// 
// A management station wishing to get the longest
// Match prefix for a given destination address
// should create the associate instance of the
// row status. The row status should be set to
// active(1) to initiate the request. Note that 
// this entire procedure may be initiated via a 
// single set request which specifies a row status 
// of createAndGo(4).
// 
// Once the request completes, the management station 
// should retrieve the values of the objects of 
// interest, and should then delete the entry.  In order 
// to prevent old entries from clogging the table, 
// entries will be aged out, but an entry will never be 
// deleted within 5 minutes of completion.
// Entries are lost after an agent restart.
// 
// I.e. to find out the longest prefix match for 
// destination address of A.B.C.D on entity whose
// entityPhysicalIndex is 1, the Management station
// will create an entry in cefLMPrefixTable with
// cefLMPrefixRowStatus.1(entPhysicalIndex).1(ipv4).A.B.C.D
// set to createAndGo(4). Management Station may query the
// value of objects cefLMPrefix and cefLMPrefixLen
// to find out the corresponding prefix entry from the
// cefPrefixTable once the value of cefLMPrefixState
// is set to matchFound(2).
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefLMPrefixTable_CefLMPrefixEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The Destination Address Type. This object
    // specifies the address type used for cefLMPrefixDestAddr.  Longest Match
    // Prefix entries are only valid  for the address type of ipv4(1) and ipv6(2).
    // The type is InetAddressType.
    CefLMPrefixDestAddrType interface{}

    // This attribute is a key. The Destination Address. The type of this address
    // is determined by the value of the cefLMPrefixDestAddrType object. The type
    // is string with length: 0..255.
    CefLMPrefixDestAddr interface{}

    // Indicates the state of this prefix search request. The type is
    // CefPrefixSearchState.
    CefLMPrefixState interface{}

    // The Network Prefix Address. Index to the cefPrefixTable. The type of this
    // address is determined by the value of the cefLMPrefixDestAddrType object.
    // The type is string with length: 0..255.
    CefLMPrefixAddr interface{}

    // The Network Prefix Length. Index to the cefPrefixTable. The type is
    // interface{} with range: 0..2040.
    CefLMPrefixLen interface{}

    // The status of this table entry.  Once the entry  status is set to
    // active(1), the associated entry  cannot be modified until the request
    // completes (cefLMPrefixState transitions to matchFound(2)  or
    // noMatchFound(3)).  Once the longest match request has been created (i.e.
    // the cefLMPrefixRowStatus has been made active), the entry cannot be
    // modified - the only operation possible after this is to delete the row. The
    // type is RowStatus.
    CefLMPrefixRowStatus interface{}
}

func (cefLMPrefixEntry *CISCOCEFMIB_CefLMPrefixTable_CefLMPrefixEntry) GetEntityData() *types.CommonEntityData {
    cefLMPrefixEntry.EntityData.YFilter = cefLMPrefixEntry.YFilter
    cefLMPrefixEntry.EntityData.YangName = "cefLMPrefixEntry"
    cefLMPrefixEntry.EntityData.BundleName = "cisco_ios_xe"
    cefLMPrefixEntry.EntityData.ParentYangName = "cefLMPrefixTable"
    cefLMPrefixEntry.EntityData.SegmentPath = "cefLMPrefixEntry" + types.AddKeyToken(cefLMPrefixEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefLMPrefixEntry.CefLMPrefixDestAddrType, "cefLMPrefixDestAddrType") + types.AddKeyToken(cefLMPrefixEntry.CefLMPrefixDestAddr, "cefLMPrefixDestAddr")
    cefLMPrefixEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefLMPrefixEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefLMPrefixEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefLMPrefixEntry.EntityData.Children = types.NewOrderedMap()
    cefLMPrefixEntry.EntityData.Leafs = types.NewOrderedMap()
    cefLMPrefixEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefLMPrefixEntry.EntPhysicalIndex})
    cefLMPrefixEntry.EntityData.Leafs.Append("cefLMPrefixDestAddrType", types.YLeaf{"CefLMPrefixDestAddrType", cefLMPrefixEntry.CefLMPrefixDestAddrType})
    cefLMPrefixEntry.EntityData.Leafs.Append("cefLMPrefixDestAddr", types.YLeaf{"CefLMPrefixDestAddr", cefLMPrefixEntry.CefLMPrefixDestAddr})
    cefLMPrefixEntry.EntityData.Leafs.Append("cefLMPrefixState", types.YLeaf{"CefLMPrefixState", cefLMPrefixEntry.CefLMPrefixState})
    cefLMPrefixEntry.EntityData.Leafs.Append("cefLMPrefixAddr", types.YLeaf{"CefLMPrefixAddr", cefLMPrefixEntry.CefLMPrefixAddr})
    cefLMPrefixEntry.EntityData.Leafs.Append("cefLMPrefixLen", types.YLeaf{"CefLMPrefixLen", cefLMPrefixEntry.CefLMPrefixLen})
    cefLMPrefixEntry.EntityData.Leafs.Append("cefLMPrefixRowStatus", types.YLeaf{"CefLMPrefixRowStatus", cefLMPrefixEntry.CefLMPrefixRowStatus})

    cefLMPrefixEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefLMPrefixDestAddrType", "CefLMPrefixDestAddr"}

    return &(cefLMPrefixEntry.EntityData)
}

// CISCOCEFMIB_CefPathTable
// CEF prefix path is a valid route to reach to a 
// destination IP prefix. Multiple paths may exist 
// out of a router to the same destination prefix. 
// This table specify lists of CEF paths.
type CISCOCEFMIB_CefPathTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contain a CEF prefix
    // path.  entPhysicalIndex is also an index for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_CefPathTable_CefPathEntry.
    CefPathEntry []*CISCOCEFMIB_CefPathTable_CefPathEntry
}

func (cefPathTable *CISCOCEFMIB_CefPathTable) GetEntityData() *types.CommonEntityData {
    cefPathTable.EntityData.YFilter = cefPathTable.YFilter
    cefPathTable.EntityData.YangName = "cefPathTable"
    cefPathTable.EntityData.BundleName = "cisco_ios_xe"
    cefPathTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefPathTable.EntityData.SegmentPath = "cefPathTable"
    cefPathTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPathTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPathTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPathTable.EntityData.Children = types.NewOrderedMap()
    cefPathTable.EntityData.Children.Append("cefPathEntry", types.YChild{"CefPathEntry", nil})
    for i := range cefPathTable.CefPathEntry {
        cefPathTable.EntityData.Children.Append(types.GetSegmentPath(cefPathTable.CefPathEntry[i]), types.YChild{"CefPathEntry", cefPathTable.CefPathEntry[i]})
    }
    cefPathTable.EntityData.Leafs = types.NewOrderedMap()

    cefPathTable.EntityData.YListKeys = []string {}

    return &(cefPathTable.EntityData)
}

// CISCOCEFMIB_CefPathTable_CefPathEntry
// If CEF is enabled on the Managed device,
// each entry contain a CEF prefix path.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefPathTable_CefPathEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefPrefixTable_CefPrefixEntry_CefPrefixType
    CefPrefixType interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefPrefixTable_CefPrefixEntry_CefPrefixAddr
    CefPrefixAddr interface{}

    // This attribute is a key. The type is string with range: 0..2040. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefPrefixTable_CefPrefixEntry_CefPrefixLen
    CefPrefixLen interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this prefix path entry. The type is interface{} with range:
    // 1..2147483647.
    CefPathId interface{}

    // Type for this CEF Path. The type is CefPathType.
    CefPathType interface{}

    // Interface associated with this CEF path.  A value of zero for this object
    // will indicate that no interface is associated with this path  entry. The
    // type is interface{} with range: 0..2147483647.
    CefPathInterface interface{}

    // Next hop address associated with this CEF path.  The value of this object
    // is only relevant for attached next hop and recursive next hop   path types
    // (when the object cefPathType is set to attachedNexthop(4) or
    // recursiveNexthop(5)). and will be set to zero for other path types.  The
    // type of this address is determined by the value of the cefPrefixType
    // object. The type is string with length: 0..255.
    CefPathNextHopAddr interface{}

    // The recursive vrf name associated with this path.  The value of this object
    // is only relevant for recursive next hop path types (when the  object
    // cefPathType is set to recursiveNexthop(5)), and '0x00' will be returned for
    // other path types. The type is string with length: 0..31.
    CefPathRecurseVrfName interface{}
}

func (cefPathEntry *CISCOCEFMIB_CefPathTable_CefPathEntry) GetEntityData() *types.CommonEntityData {
    cefPathEntry.EntityData.YFilter = cefPathEntry.YFilter
    cefPathEntry.EntityData.YangName = "cefPathEntry"
    cefPathEntry.EntityData.BundleName = "cisco_ios_xe"
    cefPathEntry.EntityData.ParentYangName = "cefPathTable"
    cefPathEntry.EntityData.SegmentPath = "cefPathEntry" + types.AddKeyToken(cefPathEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefPathEntry.CefPrefixType, "cefPrefixType") + types.AddKeyToken(cefPathEntry.CefPrefixAddr, "cefPrefixAddr") + types.AddKeyToken(cefPathEntry.CefPrefixLen, "cefPrefixLen") + types.AddKeyToken(cefPathEntry.CefPathId, "cefPathId")
    cefPathEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPathEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPathEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPathEntry.EntityData.Children = types.NewOrderedMap()
    cefPathEntry.EntityData.Leafs = types.NewOrderedMap()
    cefPathEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefPathEntry.EntPhysicalIndex})
    cefPathEntry.EntityData.Leafs.Append("cefPrefixType", types.YLeaf{"CefPrefixType", cefPathEntry.CefPrefixType})
    cefPathEntry.EntityData.Leafs.Append("cefPrefixAddr", types.YLeaf{"CefPrefixAddr", cefPathEntry.CefPrefixAddr})
    cefPathEntry.EntityData.Leafs.Append("cefPrefixLen", types.YLeaf{"CefPrefixLen", cefPathEntry.CefPrefixLen})
    cefPathEntry.EntityData.Leafs.Append("cefPathId", types.YLeaf{"CefPathId", cefPathEntry.CefPathId})
    cefPathEntry.EntityData.Leafs.Append("cefPathType", types.YLeaf{"CefPathType", cefPathEntry.CefPathType})
    cefPathEntry.EntityData.Leafs.Append("cefPathInterface", types.YLeaf{"CefPathInterface", cefPathEntry.CefPathInterface})
    cefPathEntry.EntityData.Leafs.Append("cefPathNextHopAddr", types.YLeaf{"CefPathNextHopAddr", cefPathEntry.CefPathNextHopAddr})
    cefPathEntry.EntityData.Leafs.Append("cefPathRecurseVrfName", types.YLeaf{"CefPathRecurseVrfName", cefPathEntry.CefPathRecurseVrfName})

    cefPathEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefPrefixType", "CefPrefixAddr", "CefPrefixLen", "CefPathId"}

    return &(cefPathEntry.EntityData)
}

// CISCOCEFMIB_CefAdjSummaryTable
// This table contains the summary information
// for the cefAdjTable.
type CISCOCEFMIB_CefAdjSummaryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF
    // Adjacency   summary related attributes for the Managed entity. A row exists
    // for each adjacency link type.  entPhysicalIndex is also an index for this
    // table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_CefAdjSummaryTable_CefAdjSummaryEntry.
    CefAdjSummaryEntry []*CISCOCEFMIB_CefAdjSummaryTable_CefAdjSummaryEntry
}

func (cefAdjSummaryTable *CISCOCEFMIB_CefAdjSummaryTable) GetEntityData() *types.CommonEntityData {
    cefAdjSummaryTable.EntityData.YFilter = cefAdjSummaryTable.YFilter
    cefAdjSummaryTable.EntityData.YangName = "cefAdjSummaryTable"
    cefAdjSummaryTable.EntityData.BundleName = "cisco_ios_xe"
    cefAdjSummaryTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefAdjSummaryTable.EntityData.SegmentPath = "cefAdjSummaryTable"
    cefAdjSummaryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefAdjSummaryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefAdjSummaryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefAdjSummaryTable.EntityData.Children = types.NewOrderedMap()
    cefAdjSummaryTable.EntityData.Children.Append("cefAdjSummaryEntry", types.YChild{"CefAdjSummaryEntry", nil})
    for i := range cefAdjSummaryTable.CefAdjSummaryEntry {
        cefAdjSummaryTable.EntityData.Children.Append(types.GetSegmentPath(cefAdjSummaryTable.CefAdjSummaryEntry[i]), types.YChild{"CefAdjSummaryEntry", cefAdjSummaryTable.CefAdjSummaryEntry[i]})
    }
    cefAdjSummaryTable.EntityData.Leafs = types.NewOrderedMap()

    cefAdjSummaryTable.EntityData.YListKeys = []string {}

    return &(cefAdjSummaryTable.EntityData)
}

// CISCOCEFMIB_CefAdjSummaryTable_CefAdjSummaryEntry
// If CEF is enabled on the Managed device,
// each entry contains the CEF Adjacency  
// summary related attributes for the
// Managed entity. A row exists for
// each adjacency link type.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefAdjSummaryTable_CefAdjSummaryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The link type of the adjacency. The type is
    // CefAdjLinkType.
    CefAdjSummaryLinkType interface{}

    // The total number of complete adjacencies.  The total number of adjacencies
    // which can be used  to switch traffic to a neighbour. The type is
    // interface{} with range: 0..4294967295.
    CefAdjSummaryComplete interface{}

    // The total number of incomplete adjacencies.  The total number of
    // adjacencies which cannot be  used to switch traffic in their current state.
    // The type is interface{} with range: 0..4294967295.
    CefAdjSummaryIncomplete interface{}

    // The total number of adjacencies for which the Layer 2 encapsulation string
    // (header) may be  updated (fixed up) at packet switch time. The type is
    // interface{} with range: 0..4294967295.
    CefAdjSummaryFixup interface{}

    // The total number of adjacencies for which  ip redirect (or icmp
    // redirection) is enabled. The value of this object is only relevant for ipv4
    // and ipv6 link type (when the index object  cefAdjSummaryLinkType value is
    // ipv4(1) or ipv6(2)) and will be set to zero for other link types. The type
    // is interface{} with range: 0..4294967295.
    CefAdjSummaryRedirect interface{}
}

func (cefAdjSummaryEntry *CISCOCEFMIB_CefAdjSummaryTable_CefAdjSummaryEntry) GetEntityData() *types.CommonEntityData {
    cefAdjSummaryEntry.EntityData.YFilter = cefAdjSummaryEntry.YFilter
    cefAdjSummaryEntry.EntityData.YangName = "cefAdjSummaryEntry"
    cefAdjSummaryEntry.EntityData.BundleName = "cisco_ios_xe"
    cefAdjSummaryEntry.EntityData.ParentYangName = "cefAdjSummaryTable"
    cefAdjSummaryEntry.EntityData.SegmentPath = "cefAdjSummaryEntry" + types.AddKeyToken(cefAdjSummaryEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefAdjSummaryEntry.CefAdjSummaryLinkType, "cefAdjSummaryLinkType")
    cefAdjSummaryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefAdjSummaryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefAdjSummaryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefAdjSummaryEntry.EntityData.Children = types.NewOrderedMap()
    cefAdjSummaryEntry.EntityData.Leafs = types.NewOrderedMap()
    cefAdjSummaryEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefAdjSummaryEntry.EntPhysicalIndex})
    cefAdjSummaryEntry.EntityData.Leafs.Append("cefAdjSummaryLinkType", types.YLeaf{"CefAdjSummaryLinkType", cefAdjSummaryEntry.CefAdjSummaryLinkType})
    cefAdjSummaryEntry.EntityData.Leafs.Append("cefAdjSummaryComplete", types.YLeaf{"CefAdjSummaryComplete", cefAdjSummaryEntry.CefAdjSummaryComplete})
    cefAdjSummaryEntry.EntityData.Leafs.Append("cefAdjSummaryIncomplete", types.YLeaf{"CefAdjSummaryIncomplete", cefAdjSummaryEntry.CefAdjSummaryIncomplete})
    cefAdjSummaryEntry.EntityData.Leafs.Append("cefAdjSummaryFixup", types.YLeaf{"CefAdjSummaryFixup", cefAdjSummaryEntry.CefAdjSummaryFixup})
    cefAdjSummaryEntry.EntityData.Leafs.Append("cefAdjSummaryRedirect", types.YLeaf{"CefAdjSummaryRedirect", cefAdjSummaryEntry.CefAdjSummaryRedirect})

    cefAdjSummaryEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefAdjSummaryLinkType"}

    return &(cefAdjSummaryEntry.EntityData)
}

// CISCOCEFMIB_CefAdjTable
// A list of CEF adjacencies.
type CISCOCEFMIB_CefAdjTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the adjacency 
    // attributes. Adjacency entries may exist for all the interfaces on which
    // packets can be switched out of the device. The interface is instantiated by
    // ifIndex.   Therefore, the interface index must have been assigned,
    // according to the applicable procedures, before it can be meaningfully used.
    // Generally, this means that the interface must exist.  entPhysicalIndex is
    // also an index for this table which represents entities of 'module'
    // entPhysicalClass which are capable of running CEF. The type is slice of
    // CISCOCEFMIB_CefAdjTable_CefAdjEntry.
    CefAdjEntry []*CISCOCEFMIB_CefAdjTable_CefAdjEntry
}

func (cefAdjTable *CISCOCEFMIB_CefAdjTable) GetEntityData() *types.CommonEntityData {
    cefAdjTable.EntityData.YFilter = cefAdjTable.YFilter
    cefAdjTable.EntityData.YangName = "cefAdjTable"
    cefAdjTable.EntityData.BundleName = "cisco_ios_xe"
    cefAdjTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefAdjTable.EntityData.SegmentPath = "cefAdjTable"
    cefAdjTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefAdjTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefAdjTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefAdjTable.EntityData.Children = types.NewOrderedMap()
    cefAdjTable.EntityData.Children.Append("cefAdjEntry", types.YChild{"CefAdjEntry", nil})
    for i := range cefAdjTable.CefAdjEntry {
        cefAdjTable.EntityData.Children.Append(types.GetSegmentPath(cefAdjTable.CefAdjEntry[i]), types.YChild{"CefAdjEntry", cefAdjTable.CefAdjEntry[i]})
    }
    cefAdjTable.EntityData.Leafs = types.NewOrderedMap()

    cefAdjTable.EntityData.YListKeys = []string {}

    return &(cefAdjTable.EntityData)
}

// CISCOCEFMIB_CefAdjTable_CefAdjEntry
// If CEF is enabled on the Managed device,
// each entry contains the adjacency 
// attributes. Adjacency entries may exist
// for all the interfaces on which packets
// can be switched out of the device.
// The interface is instantiated by ifIndex.  
// Therefore, the interface index must have been
// assigned, according to the applicable procedures,
// before it can be meaningfully used.
// Generally, this means that the interface must exist.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefAdjTable_CefAdjEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. Address type for the cefAdjNextHopAddr. This
    // object specifies the address type used for cefAdjNextHopAddr.   Adjacency
    // entries are only valid for the  address type of ipv4(1) and ipv6(2). The
    // type is InetAddressType.
    CefAdjNextHopAddrType interface{}

    // This attribute is a key. The next Hop address for this adjacency. The type
    // of this address is determined by the value of the cefAdjNextHopAddrType
    // object. The type is string with length: 0..255.
    CefAdjNextHopAddr interface{}

    // This attribute is a key. In cases where cefLinkType, interface and the next
    // hop address are not able to uniquely define an adjacency entry (e.g. ATM
    // and Frame Relay Bundles), this object is a unique identifier to
    // differentiate between these adjacency entries.   In all the other cases the
    // value of this  index object will be 0. The type is interface{} with range:
    // 0..4294967295.
    CefAdjConnId interface{}

    // This attribute is a key. The type is CefAdjLinkType. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefAdjSummaryTable_CefAdjSummaryEntry_CefAdjSummaryLinkType
    CefAdjSummaryLinkType interface{}

    // If the adjacency is created because some neighbour discovery mechanism has
    // discovered a neighbour and all the information required to build a frame
    // header to encapsulate traffic to the neighbour is available then the source
    // of adjacency is set to the mechanism by which the adjacency is learned. The
    // type is map[string]bool.
    CefAdjSource interface{}

    // The layer 2 encapsulation string to be used for sending the packet out
    // using this adjacency. The type is string.
    CefAdjEncap interface{}

    // For the cases, where the encapsulation string is decided at packet switch
    // time, the adjacency  encapsulation string specified by object cefAdjEncap 
    // require a fixup. I.e. for the adjacencies out of IP  Tunnels, the string
    // prepended is an IP header which has  fields which can only be setup at
    // packet switch time.  The value of this object represent the kind of fixup
    // applied to the packet.  If the encapsulation string doesn't require any
    // fixup, then the value of this object will be of zero length. The type is
    // string.
    CefAdjFixup interface{}

    // The Layer 3 MTU which can be transmitted using  this adjacency. The type is
    // interface{} with range: 0..65535. Units are bytes.
    CefAdjMTU interface{}

    // This object selects a forwarding info entry  defined in the
    // cefFESelectionTable. The  selected target is defined by an entry in the
    // cefFESelectionTable whose index value (cefFESelectionName)  is equal to
    // this object.  The value of this object will be of zero length if this
    // adjacency entry is the last forwarding  element in the forwarding path. The
    // type is string.
    CefAdjForwardingInfo interface{}

    // Number of pkts transmitted using this adjacency. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    CefAdjPkts interface{}

    // Number of pkts transmitted using this adjacency. This object is a 64-bit
    // version of cefAdjPkts. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CefAdjHCPkts interface{}

    // Number of bytes transmitted using this adjacency. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    CefAdjBytes interface{}

    // Number of bytes transmitted using this adjacency. This object is a 64-bit
    // version of cefAdjBytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CefAdjHCBytes interface{}
}

func (cefAdjEntry *CISCOCEFMIB_CefAdjTable_CefAdjEntry) GetEntityData() *types.CommonEntityData {
    cefAdjEntry.EntityData.YFilter = cefAdjEntry.YFilter
    cefAdjEntry.EntityData.YangName = "cefAdjEntry"
    cefAdjEntry.EntityData.BundleName = "cisco_ios_xe"
    cefAdjEntry.EntityData.ParentYangName = "cefAdjTable"
    cefAdjEntry.EntityData.SegmentPath = "cefAdjEntry" + types.AddKeyToken(cefAdjEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefAdjEntry.IfIndex, "ifIndex") + types.AddKeyToken(cefAdjEntry.CefAdjNextHopAddrType, "cefAdjNextHopAddrType") + types.AddKeyToken(cefAdjEntry.CefAdjNextHopAddr, "cefAdjNextHopAddr") + types.AddKeyToken(cefAdjEntry.CefAdjConnId, "cefAdjConnId") + types.AddKeyToken(cefAdjEntry.CefAdjSummaryLinkType, "cefAdjSummaryLinkType")
    cefAdjEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefAdjEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefAdjEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefAdjEntry.EntityData.Children = types.NewOrderedMap()
    cefAdjEntry.EntityData.Leafs = types.NewOrderedMap()
    cefAdjEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefAdjEntry.EntPhysicalIndex})
    cefAdjEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cefAdjEntry.IfIndex})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjNextHopAddrType", types.YLeaf{"CefAdjNextHopAddrType", cefAdjEntry.CefAdjNextHopAddrType})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjNextHopAddr", types.YLeaf{"CefAdjNextHopAddr", cefAdjEntry.CefAdjNextHopAddr})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjConnId", types.YLeaf{"CefAdjConnId", cefAdjEntry.CefAdjConnId})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjSummaryLinkType", types.YLeaf{"CefAdjSummaryLinkType", cefAdjEntry.CefAdjSummaryLinkType})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjSource", types.YLeaf{"CefAdjSource", cefAdjEntry.CefAdjSource})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjEncap", types.YLeaf{"CefAdjEncap", cefAdjEntry.CefAdjEncap})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjFixup", types.YLeaf{"CefAdjFixup", cefAdjEntry.CefAdjFixup})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjMTU", types.YLeaf{"CefAdjMTU", cefAdjEntry.CefAdjMTU})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjForwardingInfo", types.YLeaf{"CefAdjForwardingInfo", cefAdjEntry.CefAdjForwardingInfo})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjPkts", types.YLeaf{"CefAdjPkts", cefAdjEntry.CefAdjPkts})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjHCPkts", types.YLeaf{"CefAdjHCPkts", cefAdjEntry.CefAdjHCPkts})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjBytes", types.YLeaf{"CefAdjBytes", cefAdjEntry.CefAdjBytes})
    cefAdjEntry.EntityData.Leafs.Append("cefAdjHCBytes", types.YLeaf{"CefAdjHCBytes", cefAdjEntry.CefAdjHCBytes})

    cefAdjEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "IfIndex", "CefAdjNextHopAddrType", "CefAdjNextHopAddr", "CefAdjConnId", "CefAdjSummaryLinkType"}

    return &(cefAdjEntry.EntityData)
}

// CISCOCEFMIB_CefFESelectionTable
// A list of forwarding element selection entries.
type CISCOCEFMIB_CefFESelectionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contain a CEF
    // forwarding element selection list.  entPhysicalIndex is also an index for
    // this table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_CefFESelectionTable_CefFESelectionEntry.
    CefFESelectionEntry []*CISCOCEFMIB_CefFESelectionTable_CefFESelectionEntry
}

func (cefFESelectionTable *CISCOCEFMIB_CefFESelectionTable) GetEntityData() *types.CommonEntityData {
    cefFESelectionTable.EntityData.YFilter = cefFESelectionTable.YFilter
    cefFESelectionTable.EntityData.YangName = "cefFESelectionTable"
    cefFESelectionTable.EntityData.BundleName = "cisco_ios_xe"
    cefFESelectionTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefFESelectionTable.EntityData.SegmentPath = "cefFESelectionTable"
    cefFESelectionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefFESelectionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefFESelectionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefFESelectionTable.EntityData.Children = types.NewOrderedMap()
    cefFESelectionTable.EntityData.Children.Append("cefFESelectionEntry", types.YChild{"CefFESelectionEntry", nil})
    for i := range cefFESelectionTable.CefFESelectionEntry {
        cefFESelectionTable.EntityData.Children.Append(types.GetSegmentPath(cefFESelectionTable.CefFESelectionEntry[i]), types.YChild{"CefFESelectionEntry", cefFESelectionTable.CefFESelectionEntry[i]})
    }
    cefFESelectionTable.EntityData.Leafs = types.NewOrderedMap()

    cefFESelectionTable.EntityData.YListKeys = []string {}

    return &(cefFESelectionTable.EntityData)
}

// CISCOCEFMIB_CefFESelectionTable_CefFESelectionEntry
// If CEF is enabled on the Managed device,
// each entry contain a CEF forwarding element
// selection list.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefFESelectionTable_CefFESelectionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier used
    // to select a set of forwarding element lists. The type is string with
    // length: 1..32.
    CefFESelectionName interface{}

    // This attribute is a key. Secondary index to identify a forwarding elements
    // List  in this Table. The type is interface{} with range: 1..2147483647.
    CefFESelectionId interface{}

    // Special processing for a destination is indicated through the use of
    // special  forwarding element.   If the forwarding element list contains the 
    // special forwarding element, then this object  represents the type of
    // special forwarding element. The type is CefForwardingElementSpecialType.
    CefFESelectionSpecial interface{}

    // This object represent the MPLS Labels  associated with this forwarding
    // Element List.  The value of this object will be irrelevant and will be set
    // to zero length if the forwarding element list  doesn't contain a label
    // forwarding element. A zero  length label list will indicate that there is
    // no label forwarding element associated with this selection entry. The type
    // is string with length: 0..255.
    CefFESelectionLabels interface{}

    // This object represent the link type for the adjacency associated with this
    // forwarding  Element List.  The value of this object will be irrelevant and
    // will be set to unknown(5) if the forwarding element list  doesn't contain
    // an adjacency forwarding element. The type is CefAdjLinkType.
    CefFESelectionAdjLinkType interface{}

    // This object represent the interface for the adjacency associated with this
    // forwarding  Element List.  The value of this object will be irrelevant and
    // will be set to zero if the forwarding element list doesn't  contain an
    // adjacency forwarding element. The type is interface{} with range:
    // 0..2147483647.
    CefFESelectionAdjInterface interface{}

    // This object represent the next hop address type for the adjacency
    // associated with this forwarding  Element List.  The value of this object
    // will be irrelevant and will be set to unknown(0) if the forwarding element
    // list  doesn't contain an adjacency forwarding element. The type is
    // InetAddressType.
    CefFESelectionAdjNextHopAddrType interface{}

    // This object represent the next hop address for the adjacency associated
    // with this forwarding  Element List.  The value of this object will be
    // irrelevant and will be set to zero if the forwarding element list doesn't 
    // contain an adjacency forwarding element. The type is string with length:
    // 0..255.
    CefFESelectionAdjNextHopAddr interface{}

    // This object represent the connection id for the adjacency associated with
    // this forwarding  Element List.  The value of this object will be irrelevant
    // and will be set to zero if the forwarding element list doesn't  contain an
    // adjacency forwarding element.   In cases where cefFESelectionAdjLinkType,
    // interface  and the next hop address are not able to uniquely  define an
    // adjacency entry (e.g. ATM and Frame Relay Bundles), this object is a unique
    // identifier to differentiate between these adjacency entries. The type is
    // interface{} with range: 0..4294967295.
    CefFESelectionAdjConnId interface{}

    // This object represent the Vrf name for the lookup associated with this
    // forwarding  Element List.  The value of this object will be irrelevant and
    // will be set to a string containing the single octet 0x00 if the forwarding
    // element list  doesn't contain a lookup forwarding element. The type is
    // string with length: 0..31.
    CefFESelectionVrfName interface{}

    // This object represent the weighting for  load balancing between multiple
    // Forwarding Element Lists. The value of this object will be zero if load
    // balancing is associated with this selection entry. The type is interface{}
    // with range: 0..4294967295.
    CefFESelectionWeight interface{}
}

func (cefFESelectionEntry *CISCOCEFMIB_CefFESelectionTable_CefFESelectionEntry) GetEntityData() *types.CommonEntityData {
    cefFESelectionEntry.EntityData.YFilter = cefFESelectionEntry.YFilter
    cefFESelectionEntry.EntityData.YangName = "cefFESelectionEntry"
    cefFESelectionEntry.EntityData.BundleName = "cisco_ios_xe"
    cefFESelectionEntry.EntityData.ParentYangName = "cefFESelectionTable"
    cefFESelectionEntry.EntityData.SegmentPath = "cefFESelectionEntry" + types.AddKeyToken(cefFESelectionEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefFESelectionEntry.CefFESelectionName, "cefFESelectionName") + types.AddKeyToken(cefFESelectionEntry.CefFESelectionId, "cefFESelectionId")
    cefFESelectionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefFESelectionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefFESelectionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefFESelectionEntry.EntityData.Children = types.NewOrderedMap()
    cefFESelectionEntry.EntityData.Leafs = types.NewOrderedMap()
    cefFESelectionEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefFESelectionEntry.EntPhysicalIndex})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionName", types.YLeaf{"CefFESelectionName", cefFESelectionEntry.CefFESelectionName})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionId", types.YLeaf{"CefFESelectionId", cefFESelectionEntry.CefFESelectionId})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionSpecial", types.YLeaf{"CefFESelectionSpecial", cefFESelectionEntry.CefFESelectionSpecial})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionLabels", types.YLeaf{"CefFESelectionLabels", cefFESelectionEntry.CefFESelectionLabels})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionAdjLinkType", types.YLeaf{"CefFESelectionAdjLinkType", cefFESelectionEntry.CefFESelectionAdjLinkType})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionAdjInterface", types.YLeaf{"CefFESelectionAdjInterface", cefFESelectionEntry.CefFESelectionAdjInterface})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionAdjNextHopAddrType", types.YLeaf{"CefFESelectionAdjNextHopAddrType", cefFESelectionEntry.CefFESelectionAdjNextHopAddrType})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionAdjNextHopAddr", types.YLeaf{"CefFESelectionAdjNextHopAddr", cefFESelectionEntry.CefFESelectionAdjNextHopAddr})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionAdjConnId", types.YLeaf{"CefFESelectionAdjConnId", cefFESelectionEntry.CefFESelectionAdjConnId})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionVrfName", types.YLeaf{"CefFESelectionVrfName", cefFESelectionEntry.CefFESelectionVrfName})
    cefFESelectionEntry.EntityData.Leafs.Append("cefFESelectionWeight", types.YLeaf{"CefFESelectionWeight", cefFESelectionEntry.CefFESelectionWeight})

    cefFESelectionEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefFESelectionName", "CefFESelectionId"}

    return &(cefFESelectionEntry.EntityData)
}

// CISCOCEFMIB_CefCfgTable
// This table contains global config parameter 
// of CEF on the Managed device.
type CISCOCEFMIB_CefCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the Managed device supports CEF,  each entry contains the CEF config 
    // parameter for the managed entity. A row may exist for each IP version type
    // (v4 and v6) depending upon the IP version supported on the device. 
    // entPhysicalIndex is also an index for this table which represents entities
    // of 'module' entPhysicalClass which are capable of running CEF. The type is
    // slice of CISCOCEFMIB_CefCfgTable_CefCfgEntry.
    CefCfgEntry []*CISCOCEFMIB_CefCfgTable_CefCfgEntry
}

func (cefCfgTable *CISCOCEFMIB_CefCfgTable) GetEntityData() *types.CommonEntityData {
    cefCfgTable.EntityData.YFilter = cefCfgTable.YFilter
    cefCfgTable.EntityData.YangName = "cefCfgTable"
    cefCfgTable.EntityData.BundleName = "cisco_ios_xe"
    cefCfgTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefCfgTable.EntityData.SegmentPath = "cefCfgTable"
    cefCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefCfgTable.EntityData.Children = types.NewOrderedMap()
    cefCfgTable.EntityData.Children.Append("cefCfgEntry", types.YChild{"CefCfgEntry", nil})
    for i := range cefCfgTable.CefCfgEntry {
        cefCfgTable.EntityData.Children.Append(types.GetSegmentPath(cefCfgTable.CefCfgEntry[i]), types.YChild{"CefCfgEntry", cefCfgTable.CefCfgEntry[i]})
    }
    cefCfgTable.EntityData.Leafs = types.NewOrderedMap()

    cefCfgTable.EntityData.YListKeys = []string {}

    return &(cefCfgTable.EntityData)
}

// CISCOCEFMIB_CefCfgTable_CefCfgEntry
// If the Managed device supports CEF, 
// each entry contains the CEF config 
// parameter for the managed entity.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefCfgTable_CefCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // The desired state of CEF. The type is CefAdminStatus.
    CefCfgAdminState interface{}

    // The current operational state of CEF.  If the cefCfgAdminState is
    // disabled(2), then cefOperState will eventually go to the down(2) state
    // unless some error has occurred.   If cefCfgAdminState is changed to
    // enabled(1) then  cefCfgOperState should change to up(1) only if the  CEF
    // entity is ready to forward the packets using  Cisco Express Forwarding
    // (CEF) else it should remain  in the down(2) state. The up(1) state for this
    // object  indicates that CEF entity is forwarding the packet using Cisco
    // Express Forwarding. The type is CefOperStatus.
    CefCfgOperState interface{}

    // The desired state of CEF distribution. The type is CefAdminStatus.
    CefCfgDistributionAdminState interface{}

    // The current operational state of CEF distribution.  If the
    // cefCfgDistributionAdminState is disabled(2), then cefDistributionOperState
    // will eventually go to the down(2) state unless some error has occurred.   
    // If cefCfgDistributionAdminState is changed to enabled(1)  then
    // cefCfgDistributionOperState should change to up(1)  only if the CEF entity
    // is ready to forward the packets  using Distributed Cisco Express Forwarding
    // (dCEF) else  it should remain in the down(2) state. The up(1) state  for
    // this object indicates that CEF entity is forwarding the packet using
    // Distributed Cisco Express Forwarding. The type is CefOperStatus.
    CefCfgDistributionOperState interface{}

    // This object represents a bitmap of network accounting options.  CEF network
    // accounting is disabled by default.  CEF network accounting can be enabled 
    // by selecting one or more of the following CEF accounting option for the
    // value of this object.    nonRecursive(0):  enables accounting through      
    // nonrecursive prefixes.   perPrefix(1):     enables the collection of the
    // numbers                     of pkts and bytes express forwarded            
    // to a destination (prefix)   prefixLength(2):  enables accounting through   
    // prefixlength.           Once the accounting is enabled, the corresponding
    // stats  can be retrieved from the cefPrefixTable and 
    // cefStatsPrefixLenTable.  . The type is map[string]bool.
    CefCfgAccountingMap interface{}

    // Indicates the CEF Load balancing algorithm.  Setting this object to none(1)
    // will disable the Load sharing for the specified entry.  CEF load balancing
    // can be enabled by setting  this object to one of following Algorithms:  
    // original(2)  : This algorithm is based on a                  source and
    // destination hash    tunnel(3)    : This algorithm is used in               
    // tunnels environments or in                 environments where there are    
    // only a few source                      universal(4)  : This algorithm uses
    // a source and                   destination and ID hash  If the value of
    // this object is set to 'tunnel' or 'universal', then the FIXED ID for these
    // algorithms may be specified by the managed  object cefLoadSharingID. . The
    // type is CefCfgLoadSharingAlgorithm.
    CefCfgLoadSharingAlgorithm interface{}

    // The Fixed ID associated with the managed object cefCfgLoadSharingAlgorithm.
    // The hash of this object value may be used by the Load Sharing Algorithm. 
    // The value of this object is not relevant and will be set to zero if the
    // value of managed object  cefCfgLoadSharingAlgorithm is set to none(1) or
    // original(2). The default value of this object is calculated by the device
    // at the time of initialization. The type is interface{} with range:
    // 0..4294967295.
    CefCfgLoadSharingID interface{}

    // The interval time over which the CEF traffic statistics are collected. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    CefCfgTrafficStatsLoadInterval interface{}

    // The frequency with which the line card sends the traffic load statistics to
    // the Router Processor.  Setting the value of this object to 0 will disable
    // the CEF traffic statistics collection. The type is interface{} with range:
    // 0..65535. Units are seconds.
    CefCfgTrafficStatsUpdateRate interface{}
}

func (cefCfgEntry *CISCOCEFMIB_CefCfgTable_CefCfgEntry) GetEntityData() *types.CommonEntityData {
    cefCfgEntry.EntityData.YFilter = cefCfgEntry.YFilter
    cefCfgEntry.EntityData.YangName = "cefCfgEntry"
    cefCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    cefCfgEntry.EntityData.ParentYangName = "cefCfgTable"
    cefCfgEntry.EntityData.SegmentPath = "cefCfgEntry" + types.AddKeyToken(cefCfgEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefCfgEntry.CefFIBIpVersion, "cefFIBIpVersion")
    cefCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefCfgEntry.EntityData.Children = types.NewOrderedMap()
    cefCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    cefCfgEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefCfgEntry.EntPhysicalIndex})
    cefCfgEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefCfgEntry.CefFIBIpVersion})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgAdminState", types.YLeaf{"CefCfgAdminState", cefCfgEntry.CefCfgAdminState})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgOperState", types.YLeaf{"CefCfgOperState", cefCfgEntry.CefCfgOperState})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgDistributionAdminState", types.YLeaf{"CefCfgDistributionAdminState", cefCfgEntry.CefCfgDistributionAdminState})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgDistributionOperState", types.YLeaf{"CefCfgDistributionOperState", cefCfgEntry.CefCfgDistributionOperState})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgAccountingMap", types.YLeaf{"CefCfgAccountingMap", cefCfgEntry.CefCfgAccountingMap})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgLoadSharingAlgorithm", types.YLeaf{"CefCfgLoadSharingAlgorithm", cefCfgEntry.CefCfgLoadSharingAlgorithm})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgLoadSharingID", types.YLeaf{"CefCfgLoadSharingID", cefCfgEntry.CefCfgLoadSharingID})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgTrafficStatsLoadInterval", types.YLeaf{"CefCfgTrafficStatsLoadInterval", cefCfgEntry.CefCfgTrafficStatsLoadInterval})
    cefCfgEntry.EntityData.Leafs.Append("cefCfgTrafficStatsUpdateRate", types.YLeaf{"CefCfgTrafficStatsUpdateRate", cefCfgEntry.CefCfgTrafficStatsUpdateRate})

    cefCfgEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefFIBIpVersion"}

    return &(cefCfgEntry.EntityData)
}

// CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm represents object cefLoadSharingID. 
type CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm string

const (
    CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm_none CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm = "none"

    CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm_original CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm = "original"

    CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm_tunnel CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm = "tunnel"

    CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm_universal CISCOCEFMIB_CefCfgTable_CefCfgEntry_CefCfgLoadSharingAlgorithm = "universal"
)

// CISCOCEFMIB_CefResourceTable
// This table contains global resource 
// information of CEF on the Managed device.
type CISCOCEFMIB_CefResourceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the Managed device supports CEF, each entry contains the CEF Resource 
    // parameters for the managed entity.  entPhysicalIndex is also an index for
    // this table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_CefResourceTable_CefResourceEntry.
    CefResourceEntry []*CISCOCEFMIB_CefResourceTable_CefResourceEntry
}

func (cefResourceTable *CISCOCEFMIB_CefResourceTable) GetEntityData() *types.CommonEntityData {
    cefResourceTable.EntityData.YFilter = cefResourceTable.YFilter
    cefResourceTable.EntityData.YangName = "cefResourceTable"
    cefResourceTable.EntityData.BundleName = "cisco_ios_xe"
    cefResourceTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefResourceTable.EntityData.SegmentPath = "cefResourceTable"
    cefResourceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefResourceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefResourceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefResourceTable.EntityData.Children = types.NewOrderedMap()
    cefResourceTable.EntityData.Children.Append("cefResourceEntry", types.YChild{"CefResourceEntry", nil})
    for i := range cefResourceTable.CefResourceEntry {
        cefResourceTable.EntityData.Children.Append(types.GetSegmentPath(cefResourceTable.CefResourceEntry[i]), types.YChild{"CefResourceEntry", cefResourceTable.CefResourceEntry[i]})
    }
    cefResourceTable.EntityData.Leafs = types.NewOrderedMap()

    cefResourceTable.EntityData.YListKeys = []string {}

    return &(cefResourceTable.EntityData)
}

// CISCOCEFMIB_CefResourceTable_CefResourceEntry
// If the Managed device supports CEF,
// each entry contains the CEF Resource 
// parameters for the managed entity.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefResourceTable_CefResourceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // Indicates the number of bytes from the Processor Memory Pool that are
    // currently in use by CEF on the managed entity. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    CefResourceMemoryUsed interface{}

    // The CEF resource failure reason which may lead to CEF being disabled on the
    // managed entity. The type is CefFailureReason.
    CefResourceFailureReason interface{}
}

func (cefResourceEntry *CISCOCEFMIB_CefResourceTable_CefResourceEntry) GetEntityData() *types.CommonEntityData {
    cefResourceEntry.EntityData.YFilter = cefResourceEntry.YFilter
    cefResourceEntry.EntityData.YangName = "cefResourceEntry"
    cefResourceEntry.EntityData.BundleName = "cisco_ios_xe"
    cefResourceEntry.EntityData.ParentYangName = "cefResourceTable"
    cefResourceEntry.EntityData.SegmentPath = "cefResourceEntry" + types.AddKeyToken(cefResourceEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefResourceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefResourceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefResourceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefResourceEntry.EntityData.Children = types.NewOrderedMap()
    cefResourceEntry.EntityData.Leafs = types.NewOrderedMap()
    cefResourceEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefResourceEntry.EntPhysicalIndex})
    cefResourceEntry.EntityData.Leafs.Append("cefResourceMemoryUsed", types.YLeaf{"CefResourceMemoryUsed", cefResourceEntry.CefResourceMemoryUsed})
    cefResourceEntry.EntityData.Leafs.Append("cefResourceFailureReason", types.YLeaf{"CefResourceFailureReason", cefResourceEntry.CefResourceFailureReason})

    cefResourceEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefResourceEntry.EntityData)
}

// CISCOCEFMIB_CefIntTable
// This Table contains interface specific
// information of CEF on the Managed
// device.
type CISCOCEFMIB_CefIntTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device,  each entry contains the CEF
    // attributes  associated with an interface. The interface is instantiated by
    // ifIndex.   Therefore, the interface index must have been assigned,
    // according to the applicable procedures, before it can be meaningfully used.
    // Generally, this means that the interface must exist.  A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device.  entPhysicalIndex is also an index for this table which
    // represents entities of 'module' entPhysicalClass which are capable of
    // running CEF. The type is slice of CISCOCEFMIB_CefIntTable_CefIntEntry.
    CefIntEntry []*CISCOCEFMIB_CefIntTable_CefIntEntry
}

func (cefIntTable *CISCOCEFMIB_CefIntTable) GetEntityData() *types.CommonEntityData {
    cefIntTable.EntityData.YFilter = cefIntTable.YFilter
    cefIntTable.EntityData.YangName = "cefIntTable"
    cefIntTable.EntityData.BundleName = "cisco_ios_xe"
    cefIntTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefIntTable.EntityData.SegmentPath = "cefIntTable"
    cefIntTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefIntTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefIntTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefIntTable.EntityData.Children = types.NewOrderedMap()
    cefIntTable.EntityData.Children.Append("cefIntEntry", types.YChild{"CefIntEntry", nil})
    for i := range cefIntTable.CefIntEntry {
        cefIntTable.EntityData.Children.Append(types.GetSegmentPath(cefIntTable.CefIntEntry[i]), types.YChild{"CefIntEntry", cefIntTable.CefIntEntry[i]})
    }
    cefIntTable.EntityData.Leafs = types.NewOrderedMap()

    cefIntTable.EntityData.YListKeys = []string {}

    return &(cefIntTable.EntityData)
}

// CISCOCEFMIB_CefIntTable_CefIntEntry
// If CEF is enabled on the Managed device, 
// each entry contains the CEF attributes 
// associated with an interface.
// The interface is instantiated by ifIndex.  
// Therefore, the interface index must have been
// assigned, according to the applicable procedures,
// before it can be meaningfully used.
// Generally, this means that the interface must exist.
// 
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefIntTable_CefIntEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The CEF switching State for the interface.  If CEF is enabled but
    // distributed CEF(dCEF) is disabled then CEF is in cefEnabled(1) state.  If
    // distributed CEF is enabled, then CEF is in  distCefEnabled(2) state. The
    // cefDisabled(3) state indicates that CEF is disabled.  The CEF switching
    // state is only applicable to the received packet on the interface. The type
    // is CefIntSwitchingState.
    CefIntSwitchingState interface{}

    // The status of load sharing on the interface.  perPacket(1) : Router to send
    // data packets                over successive equal-cost paths               
    // without regard to individual hosts                or user sessions. 
    // perDestination(2) : Router to use multiple, equal-cost                    
    // paths to achieve load sharing  Load sharing is enabled by default  for an
    // interface when CEF is enabled. The type is CefIntLoadSharing.
    CefIntLoadSharing interface{}

    // The CEF accounting mode for the interface. CEF prefix based non-recursive
    // accounting  on an interface can be configured to store  the stats for
    // non-recursive prefixes in a internal  or external bucket.  internal(1)  : 
    // Count input traffic in the nonrecursive                 internal bucket 
    // external(2)  :  Count input traffic in the nonrecursive                
    // external bucket  The value of this object will only be effective if  value
    // of the object cefAccountingMap is set to enable nonRecursive(1) accounting.
    // The type is CefIntNonrecursiveAccouting.
    CefIntNonrecursiveAccouting interface{}
}

func (cefIntEntry *CISCOCEFMIB_CefIntTable_CefIntEntry) GetEntityData() *types.CommonEntityData {
    cefIntEntry.EntityData.YFilter = cefIntEntry.YFilter
    cefIntEntry.EntityData.YangName = "cefIntEntry"
    cefIntEntry.EntityData.BundleName = "cisco_ios_xe"
    cefIntEntry.EntityData.ParentYangName = "cefIntTable"
    cefIntEntry.EntityData.SegmentPath = "cefIntEntry" + types.AddKeyToken(cefIntEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefIntEntry.CefFIBIpVersion, "cefFIBIpVersion") + types.AddKeyToken(cefIntEntry.IfIndex, "ifIndex")
    cefIntEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefIntEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefIntEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefIntEntry.EntityData.Children = types.NewOrderedMap()
    cefIntEntry.EntityData.Leafs = types.NewOrderedMap()
    cefIntEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefIntEntry.EntPhysicalIndex})
    cefIntEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefIntEntry.CefFIBIpVersion})
    cefIntEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cefIntEntry.IfIndex})
    cefIntEntry.EntityData.Leafs.Append("cefIntSwitchingState", types.YLeaf{"CefIntSwitchingState", cefIntEntry.CefIntSwitchingState})
    cefIntEntry.EntityData.Leafs.Append("cefIntLoadSharing", types.YLeaf{"CefIntLoadSharing", cefIntEntry.CefIntLoadSharing})
    cefIntEntry.EntityData.Leafs.Append("cefIntNonrecursiveAccouting", types.YLeaf{"CefIntNonrecursiveAccouting", cefIntEntry.CefIntNonrecursiveAccouting})

    cefIntEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefFIBIpVersion", "IfIndex"}

    return &(cefIntEntry.EntityData)
}

// CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntLoadSharing represents for an interface when CEF is enabled.
type CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntLoadSharing string

const (
    CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntLoadSharing_perPacket CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntLoadSharing = "perPacket"

    CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntLoadSharing_perDestination CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntLoadSharing = "perDestination"
)

// CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntNonrecursiveAccouting represents nonRecursive(1) accounting.
type CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntNonrecursiveAccouting string

const (
    CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntNonrecursiveAccouting_internal CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntNonrecursiveAccouting = "internal"

    CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntNonrecursiveAccouting_external CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntNonrecursiveAccouting = "external"
)

// CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState represents received packet on the interface.
type CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState string

const (
    CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState_cefEnabled CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState = "cefEnabled"

    CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState_distCefEnabled CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState = "distCefEnabled"

    CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState_cefDisabled CISCOCEFMIB_CefIntTable_CefIntEntry_CefIntSwitchingState = "cefDisabled"
)

// CISCOCEFMIB_CefPeerTable
// Entity acting as RP (Routing Processor) keeps
// the CEF states for the line card entities and
// communicates with the line card entities using
// XDR. This Table contains the CEF information 
// related to peer entities on the managed device.
type CISCOCEFMIB_CefPeerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF
    // related attributes  associated with a CEF peer entity.  entPhysicalIndex
    // and entPeerPhysicalIndex are also indexes for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_CefPeerTable_CefPeerEntry.
    CefPeerEntry []*CISCOCEFMIB_CefPeerTable_CefPeerEntry
}

func (cefPeerTable *CISCOCEFMIB_CefPeerTable) GetEntityData() *types.CommonEntityData {
    cefPeerTable.EntityData.YFilter = cefPeerTable.YFilter
    cefPeerTable.EntityData.YangName = "cefPeerTable"
    cefPeerTable.EntityData.BundleName = "cisco_ios_xe"
    cefPeerTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefPeerTable.EntityData.SegmentPath = "cefPeerTable"
    cefPeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPeerTable.EntityData.Children = types.NewOrderedMap()
    cefPeerTable.EntityData.Children.Append("cefPeerEntry", types.YChild{"CefPeerEntry", nil})
    for i := range cefPeerTable.CefPeerEntry {
        cefPeerTable.EntityData.Children.Append(types.GetSegmentPath(cefPeerTable.CefPeerEntry[i]), types.YChild{"CefPeerEntry", cefPeerTable.CefPeerEntry[i]})
    }
    cefPeerTable.EntityData.Leafs = types.NewOrderedMap()

    cefPeerTable.EntityData.YListKeys = []string {}

    return &(cefPeerTable.EntityData)
}

// CISCOCEFMIB_CefPeerTable_CefPeerEntry
// If CEF is enabled on the Managed device,
// each entry contains the CEF related attributes 
// associated with a CEF peer entity.
// 
// entPhysicalIndex and entPeerPhysicalIndex are
// also indexes for this table which represents
// entities of 'module' entPhysicalClass which are
// capable of running CEF.
type CISCOCEFMIB_CefPeerTable_CefPeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The entity index for the CEF peer entity. Only the
    // entities of 'module'  entPhysicalClass are included here. The type is
    // interface{} with range: 1..2147483647.
    EntPeerPhysicalIndex interface{}

    // The current CEF operational state of the CEF peer entity.  Cef peer entity
    // oper state will be peerDisabled(1) in  the following condition:     : Cef
    // Peer entity encounters fatal error i.e. resource      allocation failure,
    // ipc failure etc     : When a reload/delete request is received from the Cef
    // Peer Entity  Once the peer entity is up and no fatal error is encountered,
    // then the value of this object will transits to the peerUp(3)  state.  If
    // the Cef Peer entity is in held stage, then the value of this object will be
    // peerHold(3). Cef peer entity can only transit to peerDisabled(1) state from
    // the peerHold(3) state. The type is CefPeerOperState.
    CefPeerOperState interface{}

    // Number of times the session with CEF peer entity  has been reset. The type
    // is interface{} with range: 0..4294967295.
    CefPeerNumberOfResets interface{}
}

func (cefPeerEntry *CISCOCEFMIB_CefPeerTable_CefPeerEntry) GetEntityData() *types.CommonEntityData {
    cefPeerEntry.EntityData.YFilter = cefPeerEntry.YFilter
    cefPeerEntry.EntityData.YangName = "cefPeerEntry"
    cefPeerEntry.EntityData.BundleName = "cisco_ios_xe"
    cefPeerEntry.EntityData.ParentYangName = "cefPeerTable"
    cefPeerEntry.EntityData.SegmentPath = "cefPeerEntry" + types.AddKeyToken(cefPeerEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefPeerEntry.EntPeerPhysicalIndex, "entPeerPhysicalIndex")
    cefPeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPeerEntry.EntityData.Children = types.NewOrderedMap()
    cefPeerEntry.EntityData.Leafs = types.NewOrderedMap()
    cefPeerEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefPeerEntry.EntPhysicalIndex})
    cefPeerEntry.EntityData.Leafs.Append("entPeerPhysicalIndex", types.YLeaf{"EntPeerPhysicalIndex", cefPeerEntry.EntPeerPhysicalIndex})
    cefPeerEntry.EntityData.Leafs.Append("cefPeerOperState", types.YLeaf{"CefPeerOperState", cefPeerEntry.CefPeerOperState})
    cefPeerEntry.EntityData.Leafs.Append("cefPeerNumberOfResets", types.YLeaf{"CefPeerNumberOfResets", cefPeerEntry.CefPeerNumberOfResets})

    cefPeerEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "EntPeerPhysicalIndex"}

    return &(cefPeerEntry.EntityData)
}

// CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState represents transit to peerDisabled(1) state from the peerHold(3) state.
type CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState string

const (
    CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState_peerDisabled CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState = "peerDisabled"

    CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState_peerUp CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState = "peerUp"

    CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState_peerHold CISCOCEFMIB_CefPeerTable_CefPeerEntry_CefPeerOperState = "peerHold"
)

// CISCOCEFMIB_CefPeerFIBTable
// Entity acting as RP (Routing Processor) keep
// the CEF FIB states for the line card entities and
// communicate with the line card entities using
// XDR. This Table contains the CEF FIB State 
// related to peer entities on the managed device.
type CISCOCEFMIB_CefPeerFIBTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF FIB
    // State  associated a CEF peer entity.  entPhysicalIndex and
    // entPeerPhysicalIndex are also indexes for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry.
    CefPeerFIBEntry []*CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry
}

func (cefPeerFIBTable *CISCOCEFMIB_CefPeerFIBTable) GetEntityData() *types.CommonEntityData {
    cefPeerFIBTable.EntityData.YFilter = cefPeerFIBTable.YFilter
    cefPeerFIBTable.EntityData.YangName = "cefPeerFIBTable"
    cefPeerFIBTable.EntityData.BundleName = "cisco_ios_xe"
    cefPeerFIBTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefPeerFIBTable.EntityData.SegmentPath = "cefPeerFIBTable"
    cefPeerFIBTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPeerFIBTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPeerFIBTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPeerFIBTable.EntityData.Children = types.NewOrderedMap()
    cefPeerFIBTable.EntityData.Children.Append("cefPeerFIBEntry", types.YChild{"CefPeerFIBEntry", nil})
    for i := range cefPeerFIBTable.CefPeerFIBEntry {
        cefPeerFIBTable.EntityData.Children.Append(types.GetSegmentPath(cefPeerFIBTable.CefPeerFIBEntry[i]), types.YChild{"CefPeerFIBEntry", cefPeerFIBTable.CefPeerFIBEntry[i]})
    }
    cefPeerFIBTable.EntityData.Leafs = types.NewOrderedMap()

    cefPeerFIBTable.EntityData.YListKeys = []string {}

    return &(cefPeerFIBTable.EntityData)
}

// CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry
// If CEF is enabled on the Managed device,
// each entry contains the CEF FIB State 
// associated a CEF peer entity.
// 
// entPhysicalIndex and entPeerPhysicalIndex are
// also indexes for this table which represents
// entities of 'module' entPhysicalClass which are
// capable of running CEF.
type CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefPeerTable_CefPeerEntry_EntPeerPhysicalIndex
    EntPeerPhysicalIndex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // The current CEF FIB Operational State for the  CEF peer entity. The type is
    // CefPeerFIBOperState.
    CefPeerFIBOperState interface{}
}

func (cefPeerFIBEntry *CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry) GetEntityData() *types.CommonEntityData {
    cefPeerFIBEntry.EntityData.YFilter = cefPeerFIBEntry.YFilter
    cefPeerFIBEntry.EntityData.YangName = "cefPeerFIBEntry"
    cefPeerFIBEntry.EntityData.BundleName = "cisco_ios_xe"
    cefPeerFIBEntry.EntityData.ParentYangName = "cefPeerFIBTable"
    cefPeerFIBEntry.EntityData.SegmentPath = "cefPeerFIBEntry" + types.AddKeyToken(cefPeerFIBEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefPeerFIBEntry.EntPeerPhysicalIndex, "entPeerPhysicalIndex") + types.AddKeyToken(cefPeerFIBEntry.CefFIBIpVersion, "cefFIBIpVersion")
    cefPeerFIBEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefPeerFIBEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefPeerFIBEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefPeerFIBEntry.EntityData.Children = types.NewOrderedMap()
    cefPeerFIBEntry.EntityData.Leafs = types.NewOrderedMap()
    cefPeerFIBEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefPeerFIBEntry.EntPhysicalIndex})
    cefPeerFIBEntry.EntityData.Leafs.Append("entPeerPhysicalIndex", types.YLeaf{"EntPeerPhysicalIndex", cefPeerFIBEntry.EntPeerPhysicalIndex})
    cefPeerFIBEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefPeerFIBEntry.CefFIBIpVersion})
    cefPeerFIBEntry.EntityData.Leafs.Append("cefPeerFIBOperState", types.YLeaf{"CefPeerFIBOperState", cefPeerFIBEntry.CefPeerFIBOperState})

    cefPeerFIBEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "EntPeerPhysicalIndex", "CefFIBIpVersion"}

    return &(cefPeerFIBEntry.EntityData)
}

// CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState represents CEF peer entity.
type CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState string

const (
    CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState_peerFIBDown CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState = "peerFIBDown"

    CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState_peerFIBUp CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState = "peerFIBUp"

    CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState_peerFIBReloadRequest CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState = "peerFIBReloadRequest"

    CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState_peerFIBReloading CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState = "peerFIBReloading"

    CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState_peerFIBSynced CISCOCEFMIB_CefPeerFIBTable_CefPeerFIBEntry_CefPeerFIBOperState = "peerFIBSynced"
)

// CISCOCEFMIB_CefCCGlobalTable
// This table contains CEF consistency checker
// (CC) global parameters for the managed device.
type CISCOCEFMIB_CefCCGlobalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the global
    // consistency  checker parameter for the managed device. A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device. The type is slice of
    // CISCOCEFMIB_CefCCGlobalTable_CefCCGlobalEntry.
    CefCCGlobalEntry []*CISCOCEFMIB_CefCCGlobalTable_CefCCGlobalEntry
}

func (cefCCGlobalTable *CISCOCEFMIB_CefCCGlobalTable) GetEntityData() *types.CommonEntityData {
    cefCCGlobalTable.EntityData.YFilter = cefCCGlobalTable.YFilter
    cefCCGlobalTable.EntityData.YangName = "cefCCGlobalTable"
    cefCCGlobalTable.EntityData.BundleName = "cisco_ios_xe"
    cefCCGlobalTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefCCGlobalTable.EntityData.SegmentPath = "cefCCGlobalTable"
    cefCCGlobalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefCCGlobalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefCCGlobalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefCCGlobalTable.EntityData.Children = types.NewOrderedMap()
    cefCCGlobalTable.EntityData.Children.Append("cefCCGlobalEntry", types.YChild{"CefCCGlobalEntry", nil})
    for i := range cefCCGlobalTable.CefCCGlobalEntry {
        cefCCGlobalTable.EntityData.Children.Append(types.GetSegmentPath(cefCCGlobalTable.CefCCGlobalEntry[i]), types.YChild{"CefCCGlobalEntry", cefCCGlobalTable.CefCCGlobalEntry[i]})
    }
    cefCCGlobalTable.EntityData.Leafs = types.NewOrderedMap()

    cefCCGlobalTable.EntityData.YListKeys = []string {}

    return &(cefCCGlobalTable.EntityData)
}

// CISCOCEFMIB_CefCCGlobalTable_CefCCGlobalEntry
// If the managed device supports CEF,
// each entry contains the global consistency 
// checker parameter for the managed device.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
type CISCOCEFMIB_CefCCGlobalTable_CefCCGlobalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // Once an inconsistency has been detected,  CEF has the ability to repair the
    // problem.  This object indicates the status of auto-repair  function for the
    // consistency checkers. The type is bool.
    CefCCGlobalAutoRepairEnabled interface{}

    // Indiactes how long the consistency checker  waits to fix an inconsistency. 
    // The value of this object has no effect when the value of object
    // cefCCGlobalAutoRepairEnabled is 'false'. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    CefCCGlobalAutoRepairDelay interface{}

    // Indicates how long the consistency checker waits to re-enable auto-repair
    // after  auto-repair runs.  The value of this object has no effect when the
    // value of object cefCCGlobalAutoRepairEnabled is 'false'. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    CefCCGlobalAutoRepairHoldDown interface{}

    // Enables the consistency checker to generate  an error message when it
    // detects an inconsistency. The type is bool.
    CefCCGlobalErrorMsgEnabled interface{}

    // Setting the value of this object to ccActionStart(1) will start the full
    // scan consistency checkers.  The Management station should poll the 
    // cefCCGlobalFullScanStatus object to get the  state of full-scan operation. 
    // Once the full-scan operation completes (value of cefCCGlobalFullScanStatus
    // object is ccStatusDone(3)),  the Management station should retrieve the
    // values of the related stats object from the cefCCTypeTable.  Setting the
    // value of this object to ccActionAbort(2) will  abort the full-scan
    // operation.  The value of this object can't be set to ccActionStart(1),  if
    // the value of object cefCCGlobalFullScanStatus is ccStatusRunning(2).  The
    // value of this object will be set to cefActionNone(1) when the full scan
    // consistency checkers have never been activated.  A Management Station
    // cannot set the value of this object to cefActionNone(1). The type is
    // CefCCAction.
    CefCCGlobalFullScanAction interface{}

    // Indicates the status of the full scan consistency checker request. The type
    // is CefCCStatus.
    CefCCGlobalFullScanStatus interface{}
}

func (cefCCGlobalEntry *CISCOCEFMIB_CefCCGlobalTable_CefCCGlobalEntry) GetEntityData() *types.CommonEntityData {
    cefCCGlobalEntry.EntityData.YFilter = cefCCGlobalEntry.YFilter
    cefCCGlobalEntry.EntityData.YangName = "cefCCGlobalEntry"
    cefCCGlobalEntry.EntityData.BundleName = "cisco_ios_xe"
    cefCCGlobalEntry.EntityData.ParentYangName = "cefCCGlobalTable"
    cefCCGlobalEntry.EntityData.SegmentPath = "cefCCGlobalEntry" + types.AddKeyToken(cefCCGlobalEntry.CefFIBIpVersion, "cefFIBIpVersion")
    cefCCGlobalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefCCGlobalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefCCGlobalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefCCGlobalEntry.EntityData.Children = types.NewOrderedMap()
    cefCCGlobalEntry.EntityData.Leafs = types.NewOrderedMap()
    cefCCGlobalEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefCCGlobalEntry.CefFIBIpVersion})
    cefCCGlobalEntry.EntityData.Leafs.Append("cefCCGlobalAutoRepairEnabled", types.YLeaf{"CefCCGlobalAutoRepairEnabled", cefCCGlobalEntry.CefCCGlobalAutoRepairEnabled})
    cefCCGlobalEntry.EntityData.Leafs.Append("cefCCGlobalAutoRepairDelay", types.YLeaf{"CefCCGlobalAutoRepairDelay", cefCCGlobalEntry.CefCCGlobalAutoRepairDelay})
    cefCCGlobalEntry.EntityData.Leafs.Append("cefCCGlobalAutoRepairHoldDown", types.YLeaf{"CefCCGlobalAutoRepairHoldDown", cefCCGlobalEntry.CefCCGlobalAutoRepairHoldDown})
    cefCCGlobalEntry.EntityData.Leafs.Append("cefCCGlobalErrorMsgEnabled", types.YLeaf{"CefCCGlobalErrorMsgEnabled", cefCCGlobalEntry.CefCCGlobalErrorMsgEnabled})
    cefCCGlobalEntry.EntityData.Leafs.Append("cefCCGlobalFullScanAction", types.YLeaf{"CefCCGlobalFullScanAction", cefCCGlobalEntry.CefCCGlobalFullScanAction})
    cefCCGlobalEntry.EntityData.Leafs.Append("cefCCGlobalFullScanStatus", types.YLeaf{"CefCCGlobalFullScanStatus", cefCCGlobalEntry.CefCCGlobalFullScanStatus})

    cefCCGlobalEntry.EntityData.YListKeys = []string {"CefFIBIpVersion"}

    return &(cefCCGlobalEntry.EntityData)
}

// CISCOCEFMIB_CefCCTypeTable
// This table contains CEF consistency
// checker types specific parameters on the managed device.
// 
// All detected inconsistency are signaled to the
// Management Station via cefInconsistencyDetection
// notification.
type CISCOCEFMIB_CefCCTypeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the consistency 
    // checker statistics for a consistency  checker type. A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device. The type is slice of CISCOCEFMIB_CefCCTypeTable_CefCCTypeEntry.
    CefCCTypeEntry []*CISCOCEFMIB_CefCCTypeTable_CefCCTypeEntry
}

func (cefCCTypeTable *CISCOCEFMIB_CefCCTypeTable) GetEntityData() *types.CommonEntityData {
    cefCCTypeTable.EntityData.YFilter = cefCCTypeTable.YFilter
    cefCCTypeTable.EntityData.YangName = "cefCCTypeTable"
    cefCCTypeTable.EntityData.BundleName = "cisco_ios_xe"
    cefCCTypeTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefCCTypeTable.EntityData.SegmentPath = "cefCCTypeTable"
    cefCCTypeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefCCTypeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefCCTypeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefCCTypeTable.EntityData.Children = types.NewOrderedMap()
    cefCCTypeTable.EntityData.Children.Append("cefCCTypeEntry", types.YChild{"CefCCTypeEntry", nil})
    for i := range cefCCTypeTable.CefCCTypeEntry {
        cefCCTypeTable.EntityData.Children.Append(types.GetSegmentPath(cefCCTypeTable.CefCCTypeEntry[i]), types.YChild{"CefCCTypeEntry", cefCCTypeTable.CefCCTypeEntry[i]})
    }
    cefCCTypeTable.EntityData.Leafs = types.NewOrderedMap()

    cefCCTypeTable.EntityData.YListKeys = []string {}

    return &(cefCCTypeTable.EntityData)
}

// CISCOCEFMIB_CefCCTypeTable_CefCCTypeEntry
// If the managed device supports CEF,
// each entry contains the consistency 
// checker statistics for a consistency 
// checker type.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
type CISCOCEFMIB_CefCCTypeTable_CefCCTypeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // This attribute is a key. Type of the consistency checker. The type is
    // CefCCType.
    CefCCType interface{}

    // Enables the passive consistency checker. Passive consistency checkers are
    // disabled by default.  Full-scan consistency checkers are always enabled. An
    // attempt to set this object to 'false' for an active consistency checker
    // will result in 'wrongValue' error. The type is bool.
    CefCCEnabled interface{}

    // The maximum number of prefixes to check per scan.  The default value for
    // this object  depends upon the consistency checker type.  The value of this
    // object will be irrelevant  for some of the consistency checkers and will be
    // set to 0.  A Management Station cannot set the value of this object to 0.
    // The type is interface{} with range: 0..4294967295.
    CefCCCount interface{}

    // The period between scans for the consistency checker. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    CefCCPeriod interface{}

    // Number of prefix consistency queries sent to CEF forwarding databases by
    // this consistency checker. The type is interface{} with range:
    // 0..4294967295.
    CefCCQueriesSent interface{}

    // Number of prefix consistency queries for which the consistency checks were
    // not performed by this  consistency checker. This may be because of some
    // internal error or resource failure. The type is interface{} with range:
    // 0..4294967295.
    CefCCQueriesIgnored interface{}

    // Number of prefix consistency queries processed by this  consistency
    // checker. The type is interface{} with range: 0..4294967295.
    CefCCQueriesChecked interface{}

    // Number of prefix consistency queries iterated back to the master database
    // by this consistency checker. The type is interface{} with range:
    // 0..4294967295.
    CefCCQueriesIterated interface{}
}

func (cefCCTypeEntry *CISCOCEFMIB_CefCCTypeTable_CefCCTypeEntry) GetEntityData() *types.CommonEntityData {
    cefCCTypeEntry.EntityData.YFilter = cefCCTypeEntry.YFilter
    cefCCTypeEntry.EntityData.YangName = "cefCCTypeEntry"
    cefCCTypeEntry.EntityData.BundleName = "cisco_ios_xe"
    cefCCTypeEntry.EntityData.ParentYangName = "cefCCTypeTable"
    cefCCTypeEntry.EntityData.SegmentPath = "cefCCTypeEntry" + types.AddKeyToken(cefCCTypeEntry.CefFIBIpVersion, "cefFIBIpVersion") + types.AddKeyToken(cefCCTypeEntry.CefCCType, "cefCCType")
    cefCCTypeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefCCTypeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefCCTypeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefCCTypeEntry.EntityData.Children = types.NewOrderedMap()
    cefCCTypeEntry.EntityData.Leafs = types.NewOrderedMap()
    cefCCTypeEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefCCTypeEntry.CefFIBIpVersion})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCType", types.YLeaf{"CefCCType", cefCCTypeEntry.CefCCType})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCEnabled", types.YLeaf{"CefCCEnabled", cefCCTypeEntry.CefCCEnabled})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCCount", types.YLeaf{"CefCCCount", cefCCTypeEntry.CefCCCount})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCPeriod", types.YLeaf{"CefCCPeriod", cefCCTypeEntry.CefCCPeriod})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCQueriesSent", types.YLeaf{"CefCCQueriesSent", cefCCTypeEntry.CefCCQueriesSent})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCQueriesIgnored", types.YLeaf{"CefCCQueriesIgnored", cefCCTypeEntry.CefCCQueriesIgnored})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCQueriesChecked", types.YLeaf{"CefCCQueriesChecked", cefCCTypeEntry.CefCCQueriesChecked})
    cefCCTypeEntry.EntityData.Leafs.Append("cefCCQueriesIterated", types.YLeaf{"CefCCQueriesIterated", cefCCTypeEntry.CefCCQueriesIterated})

    cefCCTypeEntry.EntityData.YListKeys = []string {"CefFIBIpVersion", "CefCCType"}

    return &(cefCCTypeEntry.EntityData)
}

// CISCOCEFMIB_CefInconsistencyRecordTable
// This table contains CEF inconsistency
// records.
type CISCOCEFMIB_CefInconsistencyRecordTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the inconsistency 
    // record. The type is slice of
    // CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry.
    CefInconsistencyRecordEntry []*CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry
}

func (cefInconsistencyRecordTable *CISCOCEFMIB_CefInconsistencyRecordTable) GetEntityData() *types.CommonEntityData {
    cefInconsistencyRecordTable.EntityData.YFilter = cefInconsistencyRecordTable.YFilter
    cefInconsistencyRecordTable.EntityData.YangName = "cefInconsistencyRecordTable"
    cefInconsistencyRecordTable.EntityData.BundleName = "cisco_ios_xe"
    cefInconsistencyRecordTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefInconsistencyRecordTable.EntityData.SegmentPath = "cefInconsistencyRecordTable"
    cefInconsistencyRecordTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefInconsistencyRecordTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefInconsistencyRecordTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefInconsistencyRecordTable.EntityData.Children = types.NewOrderedMap()
    cefInconsistencyRecordTable.EntityData.Children.Append("cefInconsistencyRecordEntry", types.YChild{"CefInconsistencyRecordEntry", nil})
    for i := range cefInconsistencyRecordTable.CefInconsistencyRecordEntry {
        cefInconsistencyRecordTable.EntityData.Children.Append(types.GetSegmentPath(cefInconsistencyRecordTable.CefInconsistencyRecordEntry[i]), types.YChild{"CefInconsistencyRecordEntry", cefInconsistencyRecordTable.CefInconsistencyRecordEntry[i]})
    }
    cefInconsistencyRecordTable.EntityData.Leafs = types.NewOrderedMap()

    cefInconsistencyRecordTable.EntityData.YListKeys = []string {}

    return &(cefInconsistencyRecordTable.EntityData)
}

// CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry
// If the managed device supports CEF,
// each entry contains the inconsistency 
// record.
type CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this inconsistency record entry. The type is interface{}
    // with range: 1..2147483647.
    CefInconsistencyRecId interface{}

    // The network prefix type associated with this inconsistency record. The type
    // is InetAddressType.
    CefInconsistencyPrefixType interface{}

    // The network prefix address associated with this  inconsistency record.  The
    // type of this address is determined by the value of the
    // cefInconsistencyPrefixType object. The type is string with length: 0..255.
    CefInconsistencyPrefixAddr interface{}

    // Length in bits of the inconsistency address prefix. The type is interface{}
    // with range: 0..2040.
    CefInconsistencyPrefixLen interface{}

    // Vrf name associated with this inconsistency record. The type is string with
    // length: 0..31.
    CefInconsistencyVrfName interface{}

    // The type of consistency checker who generated this inconsistency record.
    // The type is CefCCType.
    CefInconsistencyCCType interface{}

    // The entity for which this inconsistency record was  generated. The value of
    // this object will be  irrelevant and will be set to 0 when the inconsisency 
    // record is applicable for all the entities. The type is interface{} with
    // range: 0..2147483647.
    CefInconsistencyEntity interface{}

    // The reason for generating this inconsistency record.   missing(1):       
    // the prefix is missing  checksumErr(2):    checksum error was found 
    // unknown(3):        reason is unknown. The type is CefInconsistencyReason.
    CefInconsistencyReason interface{}
}

func (cefInconsistencyRecordEntry *CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry) GetEntityData() *types.CommonEntityData {
    cefInconsistencyRecordEntry.EntityData.YFilter = cefInconsistencyRecordEntry.YFilter
    cefInconsistencyRecordEntry.EntityData.YangName = "cefInconsistencyRecordEntry"
    cefInconsistencyRecordEntry.EntityData.BundleName = "cisco_ios_xe"
    cefInconsistencyRecordEntry.EntityData.ParentYangName = "cefInconsistencyRecordTable"
    cefInconsistencyRecordEntry.EntityData.SegmentPath = "cefInconsistencyRecordEntry" + types.AddKeyToken(cefInconsistencyRecordEntry.CefFIBIpVersion, "cefFIBIpVersion") + types.AddKeyToken(cefInconsistencyRecordEntry.CefInconsistencyRecId, "cefInconsistencyRecId")
    cefInconsistencyRecordEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefInconsistencyRecordEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefInconsistencyRecordEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefInconsistencyRecordEntry.EntityData.Children = types.NewOrderedMap()
    cefInconsistencyRecordEntry.EntityData.Leafs = types.NewOrderedMap()
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefInconsistencyRecordEntry.CefFIBIpVersion})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyRecId", types.YLeaf{"CefInconsistencyRecId", cefInconsistencyRecordEntry.CefInconsistencyRecId})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyPrefixType", types.YLeaf{"CefInconsistencyPrefixType", cefInconsistencyRecordEntry.CefInconsistencyPrefixType})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyPrefixAddr", types.YLeaf{"CefInconsistencyPrefixAddr", cefInconsistencyRecordEntry.CefInconsistencyPrefixAddr})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyPrefixLen", types.YLeaf{"CefInconsistencyPrefixLen", cefInconsistencyRecordEntry.CefInconsistencyPrefixLen})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyVrfName", types.YLeaf{"CefInconsistencyVrfName", cefInconsistencyRecordEntry.CefInconsistencyVrfName})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyCCType", types.YLeaf{"CefInconsistencyCCType", cefInconsistencyRecordEntry.CefInconsistencyCCType})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyEntity", types.YLeaf{"CefInconsistencyEntity", cefInconsistencyRecordEntry.CefInconsistencyEntity})
    cefInconsistencyRecordEntry.EntityData.Leafs.Append("cefInconsistencyReason", types.YLeaf{"CefInconsistencyReason", cefInconsistencyRecordEntry.CefInconsistencyReason})

    cefInconsistencyRecordEntry.EntityData.YListKeys = []string {"CefFIBIpVersion", "CefInconsistencyRecId"}

    return &(cefInconsistencyRecordEntry.EntityData)
}

// CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason represents unknown(3):        reason is unknown
type CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason string

const (
    CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason_missing CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason = "missing"

    CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason_checksumErr CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason = "checksumErr"

    CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason_unknown CISCOCEFMIB_CefInconsistencyRecordTable_CefInconsistencyRecordEntry_CefInconsistencyReason = "unknown"
)

// CISCOCEFMIB_CefStatsPrefixLenTable
// This table specifies the CEF stats based
// on the Prefix Length.
type CISCOCEFMIB_CefStatsPrefixLenTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device and if CEF accounting is set to
    // enable  prefix length based accounting (value of  cefCfgAccountingMap
    // object in the  cefCfgEntry is set to enable 'prefixLength'  accounting),
    // each entry contains the traffic  statistics for a prefix length. A row may
    // exist for each IP version type (v4 and v6) depending upon the IP version
    // supported on the device.  entPhysicalIndex is also an index for this table
    // which represents entities of 'module' entPhysicalClass which are capable of
    // running CEF. The type is slice of
    // CISCOCEFMIB_CefStatsPrefixLenTable_CefStatsPrefixLenEntry.
    CefStatsPrefixLenEntry []*CISCOCEFMIB_CefStatsPrefixLenTable_CefStatsPrefixLenEntry
}

func (cefStatsPrefixLenTable *CISCOCEFMIB_CefStatsPrefixLenTable) GetEntityData() *types.CommonEntityData {
    cefStatsPrefixLenTable.EntityData.YFilter = cefStatsPrefixLenTable.YFilter
    cefStatsPrefixLenTable.EntityData.YangName = "cefStatsPrefixLenTable"
    cefStatsPrefixLenTable.EntityData.BundleName = "cisco_ios_xe"
    cefStatsPrefixLenTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefStatsPrefixLenTable.EntityData.SegmentPath = "cefStatsPrefixLenTable"
    cefStatsPrefixLenTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefStatsPrefixLenTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefStatsPrefixLenTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefStatsPrefixLenTable.EntityData.Children = types.NewOrderedMap()
    cefStatsPrefixLenTable.EntityData.Children.Append("cefStatsPrefixLenEntry", types.YChild{"CefStatsPrefixLenEntry", nil})
    for i := range cefStatsPrefixLenTable.CefStatsPrefixLenEntry {
        cefStatsPrefixLenTable.EntityData.Children.Append(types.GetSegmentPath(cefStatsPrefixLenTable.CefStatsPrefixLenEntry[i]), types.YChild{"CefStatsPrefixLenEntry", cefStatsPrefixLenTable.CefStatsPrefixLenEntry[i]})
    }
    cefStatsPrefixLenTable.EntityData.Leafs = types.NewOrderedMap()

    cefStatsPrefixLenTable.EntityData.YListKeys = []string {}

    return &(cefStatsPrefixLenTable.EntityData)
}

// CISCOCEFMIB_CefStatsPrefixLenTable_CefStatsPrefixLenEntry
// If CEF is enabled on the Managed device
// and if CEF accounting is set to enable 
// prefix length based accounting (value of 
// cefCfgAccountingMap object in the 
// cefCfgEntry is set to enable 'prefixLength' 
// accounting), each entry contains the traffic 
// statistics for a prefix length.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefStatsPrefixLenTable_CefStatsPrefixLenEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // This attribute is a key. Length in bits of the Destination IP prefix. As
    // 0.0.0.0/0 is a valid prefix, hence  0 is a valid prefix length. The type is
    // interface{} with range: 0..2040.
    CefStatsPrefixLen interface{}

    // Number of queries received in the FIB database  for the specified IP prefix
    // length. The type is interface{} with range: 0..4294967295.
    CefStatsPrefixQueries interface{}

    // Number of queries received in the FIB database for the specified IP prefix
    // length. This object is a 64-bit version of  cefStatsPrefixQueries. The type
    // is interface{} with range: 0..18446744073709551615.
    CefStatsPrefixHCQueries interface{}

    // Number of insert operations performed to the FIB  database for the
    // specified IP prefix length. The type is interface{} with range:
    // 0..4294967295.
    CefStatsPrefixInserts interface{}

    // Number of insert operations performed to the FIB  database for the
    // specified IP prefix length. This object is a 64-bit version of 
    // cefStatsPrefixInsert. The type is interface{} with range:
    // 0..18446744073709551615.
    CefStatsPrefixHCInserts interface{}

    // Number of delete operations performed to the FIB  database for the
    // specified IP prefix length. The type is interface{} with range:
    // 0..4294967295.
    CefStatsPrefixDeletes interface{}

    // Number of delete operations performed to the FIB  database for the
    // specified IP prefix length. This object is a 64-bit version of 
    // cefStatsPrefixDelete. The type is interface{} with range:
    // 0..18446744073709551615.
    CefStatsPrefixHCDeletes interface{}

    // Total number of elements in the FIB database for the specified IP prefix
    // length. The type is interface{} with range: 0..4294967295.
    CefStatsPrefixElements interface{}

    // Total number of elements in the FIB database for the specified IP prefix
    // length. This object is a 64-bit version of  cefStatsPrefixElements. The
    // type is interface{} with range: 0..18446744073709551615.
    CefStatsPrefixHCElements interface{}
}

func (cefStatsPrefixLenEntry *CISCOCEFMIB_CefStatsPrefixLenTable_CefStatsPrefixLenEntry) GetEntityData() *types.CommonEntityData {
    cefStatsPrefixLenEntry.EntityData.YFilter = cefStatsPrefixLenEntry.YFilter
    cefStatsPrefixLenEntry.EntityData.YangName = "cefStatsPrefixLenEntry"
    cefStatsPrefixLenEntry.EntityData.BundleName = "cisco_ios_xe"
    cefStatsPrefixLenEntry.EntityData.ParentYangName = "cefStatsPrefixLenTable"
    cefStatsPrefixLenEntry.EntityData.SegmentPath = "cefStatsPrefixLenEntry" + types.AddKeyToken(cefStatsPrefixLenEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefStatsPrefixLenEntry.CefFIBIpVersion, "cefFIBIpVersion") + types.AddKeyToken(cefStatsPrefixLenEntry.CefStatsPrefixLen, "cefStatsPrefixLen")
    cefStatsPrefixLenEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefStatsPrefixLenEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefStatsPrefixLenEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefStatsPrefixLenEntry.EntityData.Children = types.NewOrderedMap()
    cefStatsPrefixLenEntry.EntityData.Leafs = types.NewOrderedMap()
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefStatsPrefixLenEntry.EntPhysicalIndex})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefStatsPrefixLenEntry.CefFIBIpVersion})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixLen", types.YLeaf{"CefStatsPrefixLen", cefStatsPrefixLenEntry.CefStatsPrefixLen})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixQueries", types.YLeaf{"CefStatsPrefixQueries", cefStatsPrefixLenEntry.CefStatsPrefixQueries})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixHCQueries", types.YLeaf{"CefStatsPrefixHCQueries", cefStatsPrefixLenEntry.CefStatsPrefixHCQueries})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixInserts", types.YLeaf{"CefStatsPrefixInserts", cefStatsPrefixLenEntry.CefStatsPrefixInserts})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixHCInserts", types.YLeaf{"CefStatsPrefixHCInserts", cefStatsPrefixLenEntry.CefStatsPrefixHCInserts})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixDeletes", types.YLeaf{"CefStatsPrefixDeletes", cefStatsPrefixLenEntry.CefStatsPrefixDeletes})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixHCDeletes", types.YLeaf{"CefStatsPrefixHCDeletes", cefStatsPrefixLenEntry.CefStatsPrefixHCDeletes})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixElements", types.YLeaf{"CefStatsPrefixElements", cefStatsPrefixLenEntry.CefStatsPrefixElements})
    cefStatsPrefixLenEntry.EntityData.Leafs.Append("cefStatsPrefixHCElements", types.YLeaf{"CefStatsPrefixHCElements", cefStatsPrefixLenEntry.CefStatsPrefixHCElements})

    cefStatsPrefixLenEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefFIBIpVersion", "CefStatsPrefixLen"}

    return &(cefStatsPrefixLenEntry.EntityData)
}

// CISCOCEFMIB_CefSwitchingStatsTable
// This table specifies the CEF switch stats.
type CISCOCEFMIB_CefSwitchingStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry specifies the switching
    // stats. A row may exist for each IP version type (v4 and v6) depending upon
    // the IP version supported on the device.  entPhysicalIndex is also an index
    // for this table which represents entities of 'module' entPhysicalClass which
    // are capable of running CEF. The type is slice of
    // CISCOCEFMIB_CefSwitchingStatsTable_CefSwitchingStatsEntry.
    CefSwitchingStatsEntry []*CISCOCEFMIB_CefSwitchingStatsTable_CefSwitchingStatsEntry
}

func (cefSwitchingStatsTable *CISCOCEFMIB_CefSwitchingStatsTable) GetEntityData() *types.CommonEntityData {
    cefSwitchingStatsTable.EntityData.YFilter = cefSwitchingStatsTable.YFilter
    cefSwitchingStatsTable.EntityData.YangName = "cefSwitchingStatsTable"
    cefSwitchingStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cefSwitchingStatsTable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefSwitchingStatsTable.EntityData.SegmentPath = "cefSwitchingStatsTable"
    cefSwitchingStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefSwitchingStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefSwitchingStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefSwitchingStatsTable.EntityData.Children = types.NewOrderedMap()
    cefSwitchingStatsTable.EntityData.Children.Append("cefSwitchingStatsEntry", types.YChild{"CefSwitchingStatsEntry", nil})
    for i := range cefSwitchingStatsTable.CefSwitchingStatsEntry {
        cefSwitchingStatsTable.EntityData.Children.Append(types.GetSegmentPath(cefSwitchingStatsTable.CefSwitchingStatsEntry[i]), types.YChild{"CefSwitchingStatsEntry", cefSwitchingStatsTable.CefSwitchingStatsEntry[i]})
    }
    cefSwitchingStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cefSwitchingStatsTable.EntityData.YListKeys = []string {}

    return &(cefSwitchingStatsTable.EntityData)
}

// CISCOCEFMIB_CefSwitchingStatsTable_CefSwitchingStatsEntry
// If CEF is enabled on the Managed device,
// each entry specifies the switching stats.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_CefSwitchingStatsTable_CefSwitchingStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_CefFIBSummaryTable_CefFIBSummaryEntry_CefFIBIpVersion
    CefFIBIpVersion interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this switching stats entry. The type is interface{} with
    // range: 1..2147483647.
    CefSwitchingIndex interface{}

    // Switch path where the feature was executed. Available switch paths are
    // platform-dependent. Following are the examples of switching paths:     RIB
    // : switching with CEF assistance     Low-end switching (LES) : CEF switch
    // path     PAS : CEF turbo switch path. The type is string with length:
    // 1..32.
    CefSwitchingPath interface{}

    // Number of packets dropped by CEF. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CefSwitchingDrop interface{}

    // Number of packets dropped by CEF. This object is a 64-bit version of
    // cefSwitchingDrop. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CefSwitchingHCDrop interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the next-fastest switching vector. The type is interface{} with
    // range: 0..4294967295. Units are packets.
    CefSwitchingPunt interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the next-fastest switching vector. This object is a 64-bit
    // version of cefSwitchingPunt. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    CefSwitchingHCPunt interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the host (process switching path).  For most of the switching
    // paths, the value of this object may be similar to cefSwitchingPunt. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CefSwitchingPunt2Host interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the host (process switching path).  For most of the switching
    // paths, the value of this object may be similar to cefSwitchingPunt. This
    // object is a 64-bit version of cefSwitchingPunt2Host. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    CefSwitchingHCPunt2Host interface{}
}

func (cefSwitchingStatsEntry *CISCOCEFMIB_CefSwitchingStatsTable_CefSwitchingStatsEntry) GetEntityData() *types.CommonEntityData {
    cefSwitchingStatsEntry.EntityData.YFilter = cefSwitchingStatsEntry.YFilter
    cefSwitchingStatsEntry.EntityData.YangName = "cefSwitchingStatsEntry"
    cefSwitchingStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cefSwitchingStatsEntry.EntityData.ParentYangName = "cefSwitchingStatsTable"
    cefSwitchingStatsEntry.EntityData.SegmentPath = "cefSwitchingStatsEntry" + types.AddKeyToken(cefSwitchingStatsEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefSwitchingStatsEntry.CefFIBIpVersion, "cefFIBIpVersion") + types.AddKeyToken(cefSwitchingStatsEntry.CefSwitchingIndex, "cefSwitchingIndex")
    cefSwitchingStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefSwitchingStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefSwitchingStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefSwitchingStatsEntry.EntityData.Children = types.NewOrderedMap()
    cefSwitchingStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cefSwitchingStatsEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefSwitchingStatsEntry.EntPhysicalIndex})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefFIBIpVersion", types.YLeaf{"CefFIBIpVersion", cefSwitchingStatsEntry.CefFIBIpVersion})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingIndex", types.YLeaf{"CefSwitchingIndex", cefSwitchingStatsEntry.CefSwitchingIndex})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingPath", types.YLeaf{"CefSwitchingPath", cefSwitchingStatsEntry.CefSwitchingPath})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingDrop", types.YLeaf{"CefSwitchingDrop", cefSwitchingStatsEntry.CefSwitchingDrop})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingHCDrop", types.YLeaf{"CefSwitchingHCDrop", cefSwitchingStatsEntry.CefSwitchingHCDrop})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingPunt", types.YLeaf{"CefSwitchingPunt", cefSwitchingStatsEntry.CefSwitchingPunt})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingHCPunt", types.YLeaf{"CefSwitchingHCPunt", cefSwitchingStatsEntry.CefSwitchingHCPunt})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingPunt2Host", types.YLeaf{"CefSwitchingPunt2Host", cefSwitchingStatsEntry.CefSwitchingPunt2Host})
    cefSwitchingStatsEntry.EntityData.Leafs.Append("cefSwitchingHCPunt2Host", types.YLeaf{"CefSwitchingHCPunt2Host", cefSwitchingStatsEntry.CefSwitchingHCPunt2Host})

    cefSwitchingStatsEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefFIBIpVersion", "CefSwitchingIndex"}

    return &(cefSwitchingStatsEntry.EntityData)
}

