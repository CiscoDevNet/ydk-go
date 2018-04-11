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

    
    Ceffib CISCOCEFMIB_Ceffib

    
    Cefcc CISCOCEFMIB_Cefcc

    
    Cefnotifcntl CISCOCEFMIB_Cefnotifcntl

    // This table contains the summary information for the cefPrefixTable.
    Ceffibsummarytable CISCOCEFMIB_Ceffibsummarytable

    // A list of CEF forwarding prefixes.
    Cefprefixtable CISCOCEFMIB_Cefprefixtable

    // A table of Longest Match Prefix Query requests.  Generator application
    // should utilize the cefLMPrefixSpinLock to try to avoid collisions. See
    // DESCRIPTION clause of cefLMPrefixSpinLock.
    Ceflmprefixtable CISCOCEFMIB_Ceflmprefixtable

    // CEF prefix path is a valid route to reach to a  destination IP prefix.
    // Multiple paths may exist  out of a router to the same destination prefix. 
    // This table specify lists of CEF paths.
    Cefpathtable CISCOCEFMIB_Cefpathtable

    // This table contains the summary information for the cefAdjTable.
    Cefadjsummarytable CISCOCEFMIB_Cefadjsummarytable

    // A list of CEF adjacencies.
    Cefadjtable CISCOCEFMIB_Cefadjtable

    // A list of forwarding element selection entries.
    Ceffeselectiontable CISCOCEFMIB_Ceffeselectiontable

    // This table contains global config parameter  of CEF on the Managed device.
    Cefcfgtable CISCOCEFMIB_Cefcfgtable

    // This table contains global resource  information of CEF on the Managed
    // device.
    Cefresourcetable CISCOCEFMIB_Cefresourcetable

    // This Table contains interface specific information of CEF on the Managed
    // device.
    Cefinttable CISCOCEFMIB_Cefinttable

    // Entity acting as RP (Routing Processor) keeps the CEF states for the line
    // card entities and communicates with the line card entities using XDR. This
    // Table contains the CEF information  related to peer entities on the managed
    // device.
    Cefpeertable CISCOCEFMIB_Cefpeertable

    // Entity acting as RP (Routing Processor) keep the CEF FIB states for the
    // line card entities and communicate with the line card entities using XDR.
    // This Table contains the CEF FIB State  related to peer entities on the
    // managed device.
    Cefpeerfibtable CISCOCEFMIB_Cefpeerfibtable

    // This table contains CEF consistency checker (CC) global parameters for the
    // managed device.
    Cefccglobaltable CISCOCEFMIB_Cefccglobaltable

    // This table contains CEF consistency checker types specific parameters on
    // the managed device.  All detected inconsistency are signaled to the
    // Management Station via cefInconsistencyDetection notification.
    Cefcctypetable CISCOCEFMIB_Cefcctypetable

    // This table contains CEF inconsistency records.
    Cefinconsistencyrecordtable CISCOCEFMIB_Cefinconsistencyrecordtable

    // This table specifies the CEF stats based on the Prefix Length.
    Cefstatsprefixlentable CISCOCEFMIB_Cefstatsprefixlentable

    // This table specifies the CEF switch stats.
    Cefswitchingstatstable CISCOCEFMIB_Cefswitchingstatstable
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

    cISCOCEFMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOCEFMIB.EntityData.Children["cefFIB"] = types.YChild{"Ceffib", &cISCOCEFMIB.Ceffib}
    cISCOCEFMIB.EntityData.Children["cefCC"] = types.YChild{"Cefcc", &cISCOCEFMIB.Cefcc}
    cISCOCEFMIB.EntityData.Children["cefNotifCntl"] = types.YChild{"Cefnotifcntl", &cISCOCEFMIB.Cefnotifcntl}
    cISCOCEFMIB.EntityData.Children["cefFIBSummaryTable"] = types.YChild{"Ceffibsummarytable", &cISCOCEFMIB.Ceffibsummarytable}
    cISCOCEFMIB.EntityData.Children["cefPrefixTable"] = types.YChild{"Cefprefixtable", &cISCOCEFMIB.Cefprefixtable}
    cISCOCEFMIB.EntityData.Children["cefLMPrefixTable"] = types.YChild{"Ceflmprefixtable", &cISCOCEFMIB.Ceflmprefixtable}
    cISCOCEFMIB.EntityData.Children["cefPathTable"] = types.YChild{"Cefpathtable", &cISCOCEFMIB.Cefpathtable}
    cISCOCEFMIB.EntityData.Children["cefAdjSummaryTable"] = types.YChild{"Cefadjsummarytable", &cISCOCEFMIB.Cefadjsummarytable}
    cISCOCEFMIB.EntityData.Children["cefAdjTable"] = types.YChild{"Cefadjtable", &cISCOCEFMIB.Cefadjtable}
    cISCOCEFMIB.EntityData.Children["cefFESelectionTable"] = types.YChild{"Ceffeselectiontable", &cISCOCEFMIB.Ceffeselectiontable}
    cISCOCEFMIB.EntityData.Children["cefCfgTable"] = types.YChild{"Cefcfgtable", &cISCOCEFMIB.Cefcfgtable}
    cISCOCEFMIB.EntityData.Children["cefResourceTable"] = types.YChild{"Cefresourcetable", &cISCOCEFMIB.Cefresourcetable}
    cISCOCEFMIB.EntityData.Children["cefIntTable"] = types.YChild{"Cefinttable", &cISCOCEFMIB.Cefinttable}
    cISCOCEFMIB.EntityData.Children["cefPeerTable"] = types.YChild{"Cefpeertable", &cISCOCEFMIB.Cefpeertable}
    cISCOCEFMIB.EntityData.Children["cefPeerFIBTable"] = types.YChild{"Cefpeerfibtable", &cISCOCEFMIB.Cefpeerfibtable}
    cISCOCEFMIB.EntityData.Children["cefCCGlobalTable"] = types.YChild{"Cefccglobaltable", &cISCOCEFMIB.Cefccglobaltable}
    cISCOCEFMIB.EntityData.Children["cefCCTypeTable"] = types.YChild{"Cefcctypetable", &cISCOCEFMIB.Cefcctypetable}
    cISCOCEFMIB.EntityData.Children["cefInconsistencyRecordTable"] = types.YChild{"Cefinconsistencyrecordtable", &cISCOCEFMIB.Cefinconsistencyrecordtable}
    cISCOCEFMIB.EntityData.Children["cefStatsPrefixLenTable"] = types.YChild{"Cefstatsprefixlentable", &cISCOCEFMIB.Cefstatsprefixlentable}
    cISCOCEFMIB.EntityData.Children["cefSwitchingStatsTable"] = types.YChild{"Cefswitchingstatstable", &cISCOCEFMIB.Cefswitchingstatstable}
    cISCOCEFMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOCEFMIB.EntityData)
}

// CISCOCEFMIB_Ceffib
type CISCOCEFMIB_Ceffib struct {
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
    Ceflmprefixspinlock interface{}
}

func (ceffib *CISCOCEFMIB_Ceffib) GetEntityData() *types.CommonEntityData {
    ceffib.EntityData.YFilter = ceffib.YFilter
    ceffib.EntityData.YangName = "cefFIB"
    ceffib.EntityData.BundleName = "cisco_ios_xe"
    ceffib.EntityData.ParentYangName = "CISCO-CEF-MIB"
    ceffib.EntityData.SegmentPath = "cefFIB"
    ceffib.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceffib.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceffib.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceffib.EntityData.Children = make(map[string]types.YChild)
    ceffib.EntityData.Leafs = make(map[string]types.YLeaf)
    ceffib.EntityData.Leafs["cefLMPrefixSpinLock"] = types.YLeaf{"Ceflmprefixspinlock", ceffib.Ceflmprefixspinlock}
    return &(ceffib.EntityData)
}

// CISCOCEFMIB_Cefcc
type CISCOCEFMIB_Cefcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time an inconsistency is detecetd. The type
    // is interface{} with range: 0..4294967295.
    Entlastinconsistencydetecttime interface{}

    // Setting the value of this object to ccActionStart(1) will reset all the
    // active consistency checkers.  The Management station should poll the 
    // cefInconsistencyResetStatus object to get the  state of inconsistency reset
    // operation.  This operation once started, cannot be aborted. Hence, the
    // value of this object cannot be set to ccActionAbort(2).  The value of this
    // object can't be set to ccActionStart(1),  if the value of object
    // cefInconsistencyResetStatus is ccStatusRunning(2). The type is CefCCAction.
    Cefinconsistencyreset interface{}

    // Indicates the status of the consistency reset request. The type is
    // CefCCStatus.
    Cefinconsistencyresetstatus interface{}
}

func (cefcc *CISCOCEFMIB_Cefcc) GetEntityData() *types.CommonEntityData {
    cefcc.EntityData.YFilter = cefcc.YFilter
    cefcc.EntityData.YangName = "cefCC"
    cefcc.EntityData.BundleName = "cisco_ios_xe"
    cefcc.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefcc.EntityData.SegmentPath = "cefCC"
    cefcc.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcc.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcc.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcc.EntityData.Children = make(map[string]types.YChild)
    cefcc.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcc.EntityData.Leafs["entLastInconsistencyDetectTime"] = types.YLeaf{"Entlastinconsistencydetecttime", cefcc.Entlastinconsistencydetecttime}
    cefcc.EntityData.Leafs["cefInconsistencyReset"] = types.YLeaf{"Cefinconsistencyreset", cefcc.Cefinconsistencyreset}
    cefcc.EntityData.Leafs["cefInconsistencyResetStatus"] = types.YLeaf{"Cefinconsistencyresetstatus", cefcc.Cefinconsistencyresetstatus}
    return &(cefcc.EntityData)
}

// CISCOCEFMIB_Cefnotifcntl
type CISCOCEFMIB_Cefnotifcntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether or not a notification should be generated on the
    // detection of CEF resource Failure. The type is bool.
    Cefresourcefailurenotifenable interface{}

    // Indicates whether or not a notification should be generated on the
    // detection of CEF peer state change. The type is bool.
    Cefpeerstatechangenotifenable interface{}

    // Indicates whether or not a notification should be generated on the
    // detection of CEF FIB peer state change. The type is bool.
    Cefpeerfibstatechangenotifenable interface{}

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
    Cefnotifthrottlinginterval interface{}

    // Indicates whether cefInconsistencyDetection notification should be
    // generated for this managed device. The type is bool.
    Cefinconsistencynotifenable interface{}
}

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetEntityData() *types.CommonEntityData {
    cefnotifcntl.EntityData.YFilter = cefnotifcntl.YFilter
    cefnotifcntl.EntityData.YangName = "cefNotifCntl"
    cefnotifcntl.EntityData.BundleName = "cisco_ios_xe"
    cefnotifcntl.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefnotifcntl.EntityData.SegmentPath = "cefNotifCntl"
    cefnotifcntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefnotifcntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefnotifcntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefnotifcntl.EntityData.Children = make(map[string]types.YChild)
    cefnotifcntl.EntityData.Leafs = make(map[string]types.YLeaf)
    cefnotifcntl.EntityData.Leafs["cefResourceFailureNotifEnable"] = types.YLeaf{"Cefresourcefailurenotifenable", cefnotifcntl.Cefresourcefailurenotifenable}
    cefnotifcntl.EntityData.Leafs["cefPeerStateChangeNotifEnable"] = types.YLeaf{"Cefpeerstatechangenotifenable", cefnotifcntl.Cefpeerstatechangenotifenable}
    cefnotifcntl.EntityData.Leafs["cefPeerFIBStateChangeNotifEnable"] = types.YLeaf{"Cefpeerfibstatechangenotifenable", cefnotifcntl.Cefpeerfibstatechangenotifenable}
    cefnotifcntl.EntityData.Leafs["cefNotifThrottlingInterval"] = types.YLeaf{"Cefnotifthrottlinginterval", cefnotifcntl.Cefnotifthrottlinginterval}
    cefnotifcntl.EntityData.Leafs["cefInconsistencyNotifEnable"] = types.YLeaf{"Cefinconsistencynotifenable", cefnotifcntl.Cefinconsistencynotifenable}
    return &(cefnotifcntl.EntityData)
}

// CISCOCEFMIB_Ceffibsummarytable
// This table contains the summary information
// for the cefPrefixTable.
type CISCOCEFMIB_Ceffibsummarytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the FIB
    // summary related attributes for the managed entity.  A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device.  entPhysicalIndex is also an index for this table which
    // represents entities of 'module' entPhysicalClass which are capable of
    // running CEF. The type is slice of
    // CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry.
    Ceffibsummaryentry []CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry
}

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetEntityData() *types.CommonEntityData {
    ceffibsummarytable.EntityData.YFilter = ceffibsummarytable.YFilter
    ceffibsummarytable.EntityData.YangName = "cefFIBSummaryTable"
    ceffibsummarytable.EntityData.BundleName = "cisco_ios_xe"
    ceffibsummarytable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    ceffibsummarytable.EntityData.SegmentPath = "cefFIBSummaryTable"
    ceffibsummarytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceffibsummarytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceffibsummarytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceffibsummarytable.EntityData.Children = make(map[string]types.YChild)
    ceffibsummarytable.EntityData.Children["cefFIBSummaryEntry"] = types.YChild{"Ceffibsummaryentry", nil}
    for i := range ceffibsummarytable.Ceffibsummaryentry {
        ceffibsummarytable.EntityData.Children[types.GetSegmentPath(&ceffibsummarytable.Ceffibsummaryentry[i])] = types.YChild{"Ceffibsummaryentry", &ceffibsummarytable.Ceffibsummaryentry[i]}
    }
    ceffibsummarytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceffibsummarytable.EntityData)
}

// CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry
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
type CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The version of IP forwarding. The type is
    // CefIpVersion.
    Ceffibipversion interface{}

    // Total number of forwarding Prefixes in FIB for the IP version specified by
    // cefFIBIpVersion object. The type is interface{} with range: 0..4294967295.
    Ceffibsummaryfwdprefixes interface{}
}

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetEntityData() *types.CommonEntityData {
    ceffibsummaryentry.EntityData.YFilter = ceffibsummaryentry.YFilter
    ceffibsummaryentry.EntityData.YangName = "cefFIBSummaryEntry"
    ceffibsummaryentry.EntityData.BundleName = "cisco_ios_xe"
    ceffibsummaryentry.EntityData.ParentYangName = "cefFIBSummaryTable"
    ceffibsummaryentry.EntityData.SegmentPath = "cefFIBSummaryEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceffibsummaryentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", ceffibsummaryentry.Ceffibipversion) + "']"
    ceffibsummaryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceffibsummaryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceffibsummaryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceffibsummaryentry.EntityData.Children = make(map[string]types.YChild)
    ceffibsummaryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceffibsummaryentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceffibsummaryentry.Entphysicalindex}
    ceffibsummaryentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", ceffibsummaryentry.Ceffibipversion}
    ceffibsummaryentry.EntityData.Leafs["cefFIBSummaryFwdPrefixes"] = types.YLeaf{"Ceffibsummaryfwdprefixes", ceffibsummaryentry.Ceffibsummaryfwdprefixes}
    return &(ceffibsummaryentry.EntityData)
}

// CISCOCEFMIB_Cefprefixtable
// A list of CEF forwarding prefixes.
type CISCOCEFMIB_Cefprefixtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the forwarding
    // prefix attributes.   CEF prefix based non-recursive stats are maintained in
    // internal and external buckets (depending upon the  value of
    // cefIntNonrecursiveAccouting object in the  CefIntEntry).  entPhysicalIndex
    // is also an index for this table which represents entities of 'module'
    // entPhysicalClass which are capable of running CEF. The type is slice of
    // CISCOCEFMIB_Cefprefixtable_Cefprefixentry.
    Cefprefixentry []CISCOCEFMIB_Cefprefixtable_Cefprefixentry
}

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetEntityData() *types.CommonEntityData {
    cefprefixtable.EntityData.YFilter = cefprefixtable.YFilter
    cefprefixtable.EntityData.YangName = "cefPrefixTable"
    cefprefixtable.EntityData.BundleName = "cisco_ios_xe"
    cefprefixtable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefprefixtable.EntityData.SegmentPath = "cefPrefixTable"
    cefprefixtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefprefixtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefprefixtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefprefixtable.EntityData.Children = make(map[string]types.YChild)
    cefprefixtable.EntityData.Children["cefPrefixEntry"] = types.YChild{"Cefprefixentry", nil}
    for i := range cefprefixtable.Cefprefixentry {
        cefprefixtable.EntityData.Children[types.GetSegmentPath(&cefprefixtable.Cefprefixentry[i])] = types.YChild{"Cefprefixentry", &cefprefixtable.Cefprefixentry[i]}
    }
    cefprefixtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefprefixtable.EntityData)
}

// CISCOCEFMIB_Cefprefixtable_Cefprefixentry
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
type CISCOCEFMIB_Cefprefixtable_Cefprefixentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The Network Prefix Type. This object specifies the
    // address type used for cefPrefixAddr.  Prefix entries are only valid for the
    // address type of ipv4(1) and ipv6(2). The type is InetAddressType.
    Cefprefixtype interface{}

    // This attribute is a key. The Network Prefix Address. The type of this
    // address is determined by the value of the cefPrefixType object. This object
    // is a Prefix Address containing the  prefix with length specified by
    // cefPrefixLen.  Any bits beyond the length specified by cefPrefixLen are
    // zeroed. The type is string with length: 0..255.
    Cefprefixaddr interface{}

    // This attribute is a key. Length in bits of the FIB Address prefix. The type
    // is interface{} with range: 0..2040.
    Cefprefixlen interface{}

    // This object indicates the associated forwarding element selection entries
    // in cefFESelectionTable. The value of this object is index value
    // (cefFESelectionName) of cefFESelectionTable. The type is string.
    Cefprefixforwardinginfo interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of packets switched to
    // this prefix. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    Cefprefixpkts interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of packets switched to
    // this prefix.   This object is a 64-bit version of  cefPrefixPkts. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    Cefprefixhcpkts interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of bytes switched to
    // this prefix. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    Cefprefixbytes interface{}

    // If CEF accounting is set to enable per prefix accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable 'perPrefix'
    // accounting), then this object represents the  number of bytes switched to
    // this prefix.  This object is a 64-bit version of  cefPrefixBytes. The type
    // is interface{} with range: 0..18446744073709551615. Units are bytes.
    Cefprefixhcbytes interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the internal bucket switched using this prefix.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    Cefprefixinternalnrpkts interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the internal bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixInternalNRPkts. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    Cefprefixinternalnrhcpkts interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the internal bucket switched using this prefix. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    Cefprefixinternalnrbytes interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the internal bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixInternalNRBytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    Cefprefixinternalnrhcbytes interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the external bucket switched using this prefix.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    Cefprefixexternalnrpkts interface{}

    // If CEF accounting is set to enable non-recursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents the number of
    // non-recursive packets in the external bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixExternalNRPkts. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    Cefprefixexternalnrhcpkts interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the external bucket switched using this prefix. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    Cefprefixexternalnrbytes interface{}

    // If CEF accounting is set to enable nonRecursive accounting (value of
    // cefCfgAccountingMap object in  the cefCfgEntry is set to enable
    // 'nonRecursive'  accounting), then this object represents  the number of
    // non-recursive bytes in the external bucket switched using this prefix. 
    // This object is a 64-bit version of  cefPrefixExternalNRBytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    Cefprefixexternalnrhcbytes interface{}
}

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetEntityData() *types.CommonEntityData {
    cefprefixentry.EntityData.YFilter = cefprefixentry.YFilter
    cefprefixentry.EntityData.YangName = "cefPrefixEntry"
    cefprefixentry.EntityData.BundleName = "cisco_ios_xe"
    cefprefixentry.EntityData.ParentYangName = "cefPrefixTable"
    cefprefixentry.EntityData.SegmentPath = "cefPrefixEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefprefixentry.Entphysicalindex) + "']" + "[cefPrefixType='" + fmt.Sprintf("%v", cefprefixentry.Cefprefixtype) + "']" + "[cefPrefixAddr='" + fmt.Sprintf("%v", cefprefixentry.Cefprefixaddr) + "']" + "[cefPrefixLen='" + fmt.Sprintf("%v", cefprefixentry.Cefprefixlen) + "']"
    cefprefixentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefprefixentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefprefixentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefprefixentry.EntityData.Children = make(map[string]types.YChild)
    cefprefixentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefprefixentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefprefixentry.Entphysicalindex}
    cefprefixentry.EntityData.Leafs["cefPrefixType"] = types.YLeaf{"Cefprefixtype", cefprefixentry.Cefprefixtype}
    cefprefixentry.EntityData.Leafs["cefPrefixAddr"] = types.YLeaf{"Cefprefixaddr", cefprefixentry.Cefprefixaddr}
    cefprefixentry.EntityData.Leafs["cefPrefixLen"] = types.YLeaf{"Cefprefixlen", cefprefixentry.Cefprefixlen}
    cefprefixentry.EntityData.Leafs["cefPrefixForwardingInfo"] = types.YLeaf{"Cefprefixforwardinginfo", cefprefixentry.Cefprefixforwardinginfo}
    cefprefixentry.EntityData.Leafs["cefPrefixPkts"] = types.YLeaf{"Cefprefixpkts", cefprefixentry.Cefprefixpkts}
    cefprefixentry.EntityData.Leafs["cefPrefixHCPkts"] = types.YLeaf{"Cefprefixhcpkts", cefprefixentry.Cefprefixhcpkts}
    cefprefixentry.EntityData.Leafs["cefPrefixBytes"] = types.YLeaf{"Cefprefixbytes", cefprefixentry.Cefprefixbytes}
    cefprefixentry.EntityData.Leafs["cefPrefixHCBytes"] = types.YLeaf{"Cefprefixhcbytes", cefprefixentry.Cefprefixhcbytes}
    cefprefixentry.EntityData.Leafs["cefPrefixInternalNRPkts"] = types.YLeaf{"Cefprefixinternalnrpkts", cefprefixentry.Cefprefixinternalnrpkts}
    cefprefixentry.EntityData.Leafs["cefPrefixInternalNRHCPkts"] = types.YLeaf{"Cefprefixinternalnrhcpkts", cefprefixentry.Cefprefixinternalnrhcpkts}
    cefprefixentry.EntityData.Leafs["cefPrefixInternalNRBytes"] = types.YLeaf{"Cefprefixinternalnrbytes", cefprefixentry.Cefprefixinternalnrbytes}
    cefprefixentry.EntityData.Leafs["cefPrefixInternalNRHCBytes"] = types.YLeaf{"Cefprefixinternalnrhcbytes", cefprefixentry.Cefprefixinternalnrhcbytes}
    cefprefixentry.EntityData.Leafs["cefPrefixExternalNRPkts"] = types.YLeaf{"Cefprefixexternalnrpkts", cefprefixentry.Cefprefixexternalnrpkts}
    cefprefixentry.EntityData.Leafs["cefPrefixExternalNRHCPkts"] = types.YLeaf{"Cefprefixexternalnrhcpkts", cefprefixentry.Cefprefixexternalnrhcpkts}
    cefprefixentry.EntityData.Leafs["cefPrefixExternalNRBytes"] = types.YLeaf{"Cefprefixexternalnrbytes", cefprefixentry.Cefprefixexternalnrbytes}
    cefprefixentry.EntityData.Leafs["cefPrefixExternalNRHCBytes"] = types.YLeaf{"Cefprefixexternalnrhcbytes", cefprefixentry.Cefprefixexternalnrhcbytes}
    return &(cefprefixentry.EntityData)
}

// CISCOCEFMIB_Ceflmprefixtable
// A table of Longest Match Prefix Query requests.
// 
// Generator application should utilize the
// cefLMPrefixSpinLock to try to avoid collisions.
// See DESCRIPTION clause of cefLMPrefixSpinLock.
type CISCOCEFMIB_Ceflmprefixtable struct {
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
    // CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry.
    Ceflmprefixentry []CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry
}

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetEntityData() *types.CommonEntityData {
    ceflmprefixtable.EntityData.YFilter = ceflmprefixtable.YFilter
    ceflmprefixtable.EntityData.YangName = "cefLMPrefixTable"
    ceflmprefixtable.EntityData.BundleName = "cisco_ios_xe"
    ceflmprefixtable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    ceflmprefixtable.EntityData.SegmentPath = "cefLMPrefixTable"
    ceflmprefixtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceflmprefixtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceflmprefixtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceflmprefixtable.EntityData.Children = make(map[string]types.YChild)
    ceflmprefixtable.EntityData.Children["cefLMPrefixEntry"] = types.YChild{"Ceflmprefixentry", nil}
    for i := range ceflmprefixtable.Ceflmprefixentry {
        ceflmprefixtable.EntityData.Children[types.GetSegmentPath(&ceflmprefixtable.Ceflmprefixentry[i])] = types.YChild{"Ceflmprefixentry", &ceflmprefixtable.Ceflmprefixentry[i]}
    }
    ceflmprefixtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceflmprefixtable.EntityData)
}

// CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry
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
type CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The Destination Address Type. This object
    // specifies the address type used for cefLMPrefixDestAddr.  Longest Match
    // Prefix entries are only valid  for the address type of ipv4(1) and ipv6(2).
    // The type is InetAddressType.
    Ceflmprefixdestaddrtype interface{}

    // This attribute is a key. The Destination Address. The type of this address
    // is determined by the value of the cefLMPrefixDestAddrType object. The type
    // is string with length: 0..255.
    Ceflmprefixdestaddr interface{}

    // Indicates the state of this prefix search request. The type is
    // CefPrefixSearchState.
    Ceflmprefixstate interface{}

    // The Network Prefix Address. Index to the cefPrefixTable. The type of this
    // address is determined by the value of the cefLMPrefixDestAddrType object.
    // The type is string with length: 0..255.
    Ceflmprefixaddr interface{}

    // The Network Prefix Length. Index to the cefPrefixTable. The type is
    // interface{} with range: 0..2040.
    Ceflmprefixlen interface{}

    // The status of this table entry.  Once the entry  status is set to
    // active(1), the associated entry  cannot be modified until the request
    // completes (cefLMPrefixState transitions to matchFound(2)  or
    // noMatchFound(3)).  Once the longest match request has been created (i.e.
    // the cefLMPrefixRowStatus has been made active), the entry cannot be
    // modified - the only operation possible after this is to delete the row. The
    // type is RowStatus.
    Ceflmprefixrowstatus interface{}
}

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetEntityData() *types.CommonEntityData {
    ceflmprefixentry.EntityData.YFilter = ceflmprefixentry.YFilter
    ceflmprefixentry.EntityData.YangName = "cefLMPrefixEntry"
    ceflmprefixentry.EntityData.BundleName = "cisco_ios_xe"
    ceflmprefixentry.EntityData.ParentYangName = "cefLMPrefixTable"
    ceflmprefixentry.EntityData.SegmentPath = "cefLMPrefixEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceflmprefixentry.Entphysicalindex) + "']" + "[cefLMPrefixDestAddrType='" + fmt.Sprintf("%v", ceflmprefixentry.Ceflmprefixdestaddrtype) + "']" + "[cefLMPrefixDestAddr='" + fmt.Sprintf("%v", ceflmprefixentry.Ceflmprefixdestaddr) + "']"
    ceflmprefixentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceflmprefixentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceflmprefixentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceflmprefixentry.EntityData.Children = make(map[string]types.YChild)
    ceflmprefixentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceflmprefixentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceflmprefixentry.Entphysicalindex}
    ceflmprefixentry.EntityData.Leafs["cefLMPrefixDestAddrType"] = types.YLeaf{"Ceflmprefixdestaddrtype", ceflmprefixentry.Ceflmprefixdestaddrtype}
    ceflmprefixentry.EntityData.Leafs["cefLMPrefixDestAddr"] = types.YLeaf{"Ceflmprefixdestaddr", ceflmprefixentry.Ceflmprefixdestaddr}
    ceflmprefixentry.EntityData.Leafs["cefLMPrefixState"] = types.YLeaf{"Ceflmprefixstate", ceflmprefixentry.Ceflmprefixstate}
    ceflmprefixentry.EntityData.Leafs["cefLMPrefixAddr"] = types.YLeaf{"Ceflmprefixaddr", ceflmprefixentry.Ceflmprefixaddr}
    ceflmprefixentry.EntityData.Leafs["cefLMPrefixLen"] = types.YLeaf{"Ceflmprefixlen", ceflmprefixentry.Ceflmprefixlen}
    ceflmprefixentry.EntityData.Leafs["cefLMPrefixRowStatus"] = types.YLeaf{"Ceflmprefixrowstatus", ceflmprefixentry.Ceflmprefixrowstatus}
    return &(ceflmprefixentry.EntityData)
}

// CISCOCEFMIB_Cefpathtable
// CEF prefix path is a valid route to reach to a 
// destination IP prefix. Multiple paths may exist 
// out of a router to the same destination prefix. 
// This table specify lists of CEF paths.
type CISCOCEFMIB_Cefpathtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contain a CEF prefix
    // path.  entPhysicalIndex is also an index for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_Cefpathtable_Cefpathentry.
    Cefpathentry []CISCOCEFMIB_Cefpathtable_Cefpathentry
}

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetEntityData() *types.CommonEntityData {
    cefpathtable.EntityData.YFilter = cefpathtable.YFilter
    cefpathtable.EntityData.YangName = "cefPathTable"
    cefpathtable.EntityData.BundleName = "cisco_ios_xe"
    cefpathtable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefpathtable.EntityData.SegmentPath = "cefPathTable"
    cefpathtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefpathtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefpathtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefpathtable.EntityData.Children = make(map[string]types.YChild)
    cefpathtable.EntityData.Children["cefPathEntry"] = types.YChild{"Cefpathentry", nil}
    for i := range cefpathtable.Cefpathentry {
        cefpathtable.EntityData.Children[types.GetSegmentPath(&cefpathtable.Cefpathentry[i])] = types.YChild{"Cefpathentry", &cefpathtable.Cefpathentry[i]}
    }
    cefpathtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefpathtable.EntityData)
}

// CISCOCEFMIB_Cefpathtable_Cefpathentry
// If CEF is enabled on the Managed device,
// each entry contain a CEF prefix path.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_Cefpathtable_Cefpathentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Cefprefixtable_Cefprefixentry_Cefprefixtype
    Cefprefixtype interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Cefprefixtable_Cefprefixentry_Cefprefixaddr
    Cefprefixaddr interface{}

    // This attribute is a key. The type is string with range: 0..2040. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Cefprefixtable_Cefprefixentry_Cefprefixlen
    Cefprefixlen interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this prefix path entry. The type is interface{} with range:
    // 1..2147483647.
    Cefpathid interface{}

    // Type for this CEF Path. The type is CefPathType.
    Cefpathtype interface{}

    // Interface associated with this CEF path.  A value of zero for this object
    // will indicate that no interface is associated with this path  entry. The
    // type is interface{} with range: 0..2147483647.
    Cefpathinterface interface{}

    // Next hop address associated with this CEF path.  The value of this object
    // is only relevant for attached next hop and recursive next hop   path types
    // (when the object cefPathType is set to attachedNexthop(4) or
    // recursiveNexthop(5)). and will be set to zero for other path types.  The
    // type of this address is determined by the value of the cefPrefixType
    // object. The type is string with length: 0..255.
    Cefpathnexthopaddr interface{}

    // The recursive vrf name associated with this path.  The value of this object
    // is only relevant for recursive next hop path types (when the  object
    // cefPathType is set to recursiveNexthop(5)), and '0x00' will be returned for
    // other path types. The type is string with length: 0..31.
    Cefpathrecursevrfname interface{}
}

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetEntityData() *types.CommonEntityData {
    cefpathentry.EntityData.YFilter = cefpathentry.YFilter
    cefpathentry.EntityData.YangName = "cefPathEntry"
    cefpathentry.EntityData.BundleName = "cisco_ios_xe"
    cefpathentry.EntityData.ParentYangName = "cefPathTable"
    cefpathentry.EntityData.SegmentPath = "cefPathEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefpathentry.Entphysicalindex) + "']" + "[cefPrefixType='" + fmt.Sprintf("%v", cefpathentry.Cefprefixtype) + "']" + "[cefPrefixAddr='" + fmt.Sprintf("%v", cefpathentry.Cefprefixaddr) + "']" + "[cefPrefixLen='" + fmt.Sprintf("%v", cefpathentry.Cefprefixlen) + "']" + "[cefPathId='" + fmt.Sprintf("%v", cefpathentry.Cefpathid) + "']"
    cefpathentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefpathentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefpathentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefpathentry.EntityData.Children = make(map[string]types.YChild)
    cefpathentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefpathentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefpathentry.Entphysicalindex}
    cefpathentry.EntityData.Leafs["cefPrefixType"] = types.YLeaf{"Cefprefixtype", cefpathentry.Cefprefixtype}
    cefpathentry.EntityData.Leafs["cefPrefixAddr"] = types.YLeaf{"Cefprefixaddr", cefpathentry.Cefprefixaddr}
    cefpathentry.EntityData.Leafs["cefPrefixLen"] = types.YLeaf{"Cefprefixlen", cefpathentry.Cefprefixlen}
    cefpathentry.EntityData.Leafs["cefPathId"] = types.YLeaf{"Cefpathid", cefpathentry.Cefpathid}
    cefpathentry.EntityData.Leafs["cefPathType"] = types.YLeaf{"Cefpathtype", cefpathentry.Cefpathtype}
    cefpathentry.EntityData.Leafs["cefPathInterface"] = types.YLeaf{"Cefpathinterface", cefpathentry.Cefpathinterface}
    cefpathentry.EntityData.Leafs["cefPathNextHopAddr"] = types.YLeaf{"Cefpathnexthopaddr", cefpathentry.Cefpathnexthopaddr}
    cefpathentry.EntityData.Leafs["cefPathRecurseVrfName"] = types.YLeaf{"Cefpathrecursevrfname", cefpathentry.Cefpathrecursevrfname}
    return &(cefpathentry.EntityData)
}

// CISCOCEFMIB_Cefadjsummarytable
// This table contains the summary information
// for the cefAdjTable.
type CISCOCEFMIB_Cefadjsummarytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF
    // Adjacency   summary related attributes for the Managed entity. A row exists
    // for each adjacency link type.  entPhysicalIndex is also an index for this
    // table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry.
    Cefadjsummaryentry []CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry
}

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetEntityData() *types.CommonEntityData {
    cefadjsummarytable.EntityData.YFilter = cefadjsummarytable.YFilter
    cefadjsummarytable.EntityData.YangName = "cefAdjSummaryTable"
    cefadjsummarytable.EntityData.BundleName = "cisco_ios_xe"
    cefadjsummarytable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefadjsummarytable.EntityData.SegmentPath = "cefAdjSummaryTable"
    cefadjsummarytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefadjsummarytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefadjsummarytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefadjsummarytable.EntityData.Children = make(map[string]types.YChild)
    cefadjsummarytable.EntityData.Children["cefAdjSummaryEntry"] = types.YChild{"Cefadjsummaryentry", nil}
    for i := range cefadjsummarytable.Cefadjsummaryentry {
        cefadjsummarytable.EntityData.Children[types.GetSegmentPath(&cefadjsummarytable.Cefadjsummaryentry[i])] = types.YChild{"Cefadjsummaryentry", &cefadjsummarytable.Cefadjsummaryentry[i]}
    }
    cefadjsummarytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefadjsummarytable.EntityData)
}

// CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry
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
type CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The link type of the adjacency. The type is
    // CefAdjLinkType.
    Cefadjsummarylinktype interface{}

    // The total number of complete adjacencies.  The total number of adjacencies
    // which can be used  to switch traffic to a neighbour. The type is
    // interface{} with range: 0..4294967295.
    Cefadjsummarycomplete interface{}

    // The total number of incomplete adjacencies.  The total number of
    // adjacencies which cannot be  used to switch traffic in their current state.
    // The type is interface{} with range: 0..4294967295.
    Cefadjsummaryincomplete interface{}

    // The total number of adjacencies for which the Layer 2 encapsulation string
    // (header) may be  updated (fixed up) at packet switch time. The type is
    // interface{} with range: 0..4294967295.
    Cefadjsummaryfixup interface{}

    // The total number of adjacencies for which  ip redirect (or icmp
    // redirection) is enabled. The value of this object is only relevant for ipv4
    // and ipv6 link type (when the index object  cefAdjSummaryLinkType value is
    // ipv4(1) or ipv6(2)) and will be set to zero for other link types. The type
    // is interface{} with range: 0..4294967295.
    Cefadjsummaryredirect interface{}
}

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetEntityData() *types.CommonEntityData {
    cefadjsummaryentry.EntityData.YFilter = cefadjsummaryentry.YFilter
    cefadjsummaryentry.EntityData.YangName = "cefAdjSummaryEntry"
    cefadjsummaryentry.EntityData.BundleName = "cisco_ios_xe"
    cefadjsummaryentry.EntityData.ParentYangName = "cefAdjSummaryTable"
    cefadjsummaryentry.EntityData.SegmentPath = "cefAdjSummaryEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefadjsummaryentry.Entphysicalindex) + "']" + "[cefAdjSummaryLinkType='" + fmt.Sprintf("%v", cefadjsummaryentry.Cefadjsummarylinktype) + "']"
    cefadjsummaryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefadjsummaryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefadjsummaryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefadjsummaryentry.EntityData.Children = make(map[string]types.YChild)
    cefadjsummaryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefadjsummaryentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefadjsummaryentry.Entphysicalindex}
    cefadjsummaryentry.EntityData.Leafs["cefAdjSummaryLinkType"] = types.YLeaf{"Cefadjsummarylinktype", cefadjsummaryentry.Cefadjsummarylinktype}
    cefadjsummaryentry.EntityData.Leafs["cefAdjSummaryComplete"] = types.YLeaf{"Cefadjsummarycomplete", cefadjsummaryentry.Cefadjsummarycomplete}
    cefadjsummaryentry.EntityData.Leafs["cefAdjSummaryIncomplete"] = types.YLeaf{"Cefadjsummaryincomplete", cefadjsummaryentry.Cefadjsummaryincomplete}
    cefadjsummaryentry.EntityData.Leafs["cefAdjSummaryFixup"] = types.YLeaf{"Cefadjsummaryfixup", cefadjsummaryentry.Cefadjsummaryfixup}
    cefadjsummaryentry.EntityData.Leafs["cefAdjSummaryRedirect"] = types.YLeaf{"Cefadjsummaryredirect", cefadjsummaryentry.Cefadjsummaryredirect}
    return &(cefadjsummaryentry.EntityData)
}

// CISCOCEFMIB_Cefadjtable
// A list of CEF adjacencies.
type CISCOCEFMIB_Cefadjtable struct {
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
    // CISCOCEFMIB_Cefadjtable_Cefadjentry.
    Cefadjentry []CISCOCEFMIB_Cefadjtable_Cefadjentry
}

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetEntityData() *types.CommonEntityData {
    cefadjtable.EntityData.YFilter = cefadjtable.YFilter
    cefadjtable.EntityData.YangName = "cefAdjTable"
    cefadjtable.EntityData.BundleName = "cisco_ios_xe"
    cefadjtable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefadjtable.EntityData.SegmentPath = "cefAdjTable"
    cefadjtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefadjtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefadjtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefadjtable.EntityData.Children = make(map[string]types.YChild)
    cefadjtable.EntityData.Children["cefAdjEntry"] = types.YChild{"Cefadjentry", nil}
    for i := range cefadjtable.Cefadjentry {
        cefadjtable.EntityData.Children[types.GetSegmentPath(&cefadjtable.Cefadjentry[i])] = types.YChild{"Cefadjentry", &cefadjtable.Cefadjentry[i]}
    }
    cefadjtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefadjtable.EntityData)
}

// CISCOCEFMIB_Cefadjtable_Cefadjentry
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
type CISCOCEFMIB_Cefadjtable_Cefadjentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. Address type for the cefAdjNextHopAddr. This
    // object specifies the address type used for cefAdjNextHopAddr.   Adjacency
    // entries are only valid for the  address type of ipv4(1) and ipv6(2). The
    // type is InetAddressType.
    Cefadjnexthopaddrtype interface{}

    // This attribute is a key. The next Hop address for this adjacency. The type
    // of this address is determined by the value of the cefAdjNextHopAddrType
    // object. The type is string with length: 0..255.
    Cefadjnexthopaddr interface{}

    // This attribute is a key. In cases where cefLinkType, interface and the next
    // hop address are not able to uniquely define an adjacency entry (e.g. ATM
    // and Frame Relay Bundles), this object is a unique identifier to
    // differentiate between these adjacency entries.   In all the other cases the
    // value of this  index object will be 0. The type is interface{} with range:
    // 0..4294967295.
    Cefadjconnid interface{}

    // This attribute is a key. The type is CefAdjLinkType. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry_Cefadjsummarylinktype
    Cefadjsummarylinktype interface{}

    // If the adjacency is created because some neighbour discovery mechanism has
    // discovered a neighbour and all the information required to build a frame
    // header to encapsulate traffic to the neighbour is available then the source
    // of adjacency is set to the mechanism by which the adjacency is learned. The
    // type is map[string]bool.
    Cefadjsource interface{}

    // The layer 2 encapsulation string to be used for sending the packet out
    // using this adjacency. The type is string.
    Cefadjencap interface{}

    // For the cases, where the encapsulation string is decided at packet switch
    // time, the adjacency  encapsulation string specified by object cefAdjEncap 
    // require a fixup. I.e. for the adjacencies out of IP  Tunnels, the string
    // prepended is an IP header which has  fields which can only be setup at
    // packet switch time.  The value of this object represent the kind of fixup
    // applied to the packet.  If the encapsulation string doesn't require any
    // fixup, then the value of this object will be of zero length. The type is
    // string.
    Cefadjfixup interface{}

    // The Layer 3 MTU which can be transmitted using  this adjacency. The type is
    // interface{} with range: 0..65535. Units are bytes.
    Cefadjmtu interface{}

    // This object selects a forwarding info entry  defined in the
    // cefFESelectionTable. The  selected target is defined by an entry in the
    // cefFESelectionTable whose index value (cefFESelectionName)  is equal to
    // this object.  The value of this object will be of zero length if this
    // adjacency entry is the last forwarding  element in the forwarding path. The
    // type is string.
    Cefadjforwardinginfo interface{}

    // Number of pkts transmitted using this adjacency. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    Cefadjpkts interface{}

    // Number of pkts transmitted using this adjacency. This object is a 64-bit
    // version of cefAdjPkts. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Cefadjhcpkts interface{}

    // Number of bytes transmitted using this adjacency. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    Cefadjbytes interface{}

    // Number of bytes transmitted using this adjacency. This object is a 64-bit
    // version of cefAdjBytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Cefadjhcbytes interface{}
}

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetEntityData() *types.CommonEntityData {
    cefadjentry.EntityData.YFilter = cefadjentry.YFilter
    cefadjentry.EntityData.YangName = "cefAdjEntry"
    cefadjentry.EntityData.BundleName = "cisco_ios_xe"
    cefadjentry.EntityData.ParentYangName = "cefAdjTable"
    cefadjentry.EntityData.SegmentPath = "cefAdjEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefadjentry.Entphysicalindex) + "']" + "[ifIndex='" + fmt.Sprintf("%v", cefadjentry.Ifindex) + "']" + "[cefAdjNextHopAddrType='" + fmt.Sprintf("%v", cefadjentry.Cefadjnexthopaddrtype) + "']" + "[cefAdjNextHopAddr='" + fmt.Sprintf("%v", cefadjentry.Cefadjnexthopaddr) + "']" + "[cefAdjConnId='" + fmt.Sprintf("%v", cefadjentry.Cefadjconnid) + "']" + "[cefAdjSummaryLinkType='" + fmt.Sprintf("%v", cefadjentry.Cefadjsummarylinktype) + "']"
    cefadjentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefadjentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefadjentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefadjentry.EntityData.Children = make(map[string]types.YChild)
    cefadjentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefadjentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefadjentry.Entphysicalindex}
    cefadjentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cefadjentry.Ifindex}
    cefadjentry.EntityData.Leafs["cefAdjNextHopAddrType"] = types.YLeaf{"Cefadjnexthopaddrtype", cefadjentry.Cefadjnexthopaddrtype}
    cefadjentry.EntityData.Leafs["cefAdjNextHopAddr"] = types.YLeaf{"Cefadjnexthopaddr", cefadjentry.Cefadjnexthopaddr}
    cefadjentry.EntityData.Leafs["cefAdjConnId"] = types.YLeaf{"Cefadjconnid", cefadjentry.Cefadjconnid}
    cefadjentry.EntityData.Leafs["cefAdjSummaryLinkType"] = types.YLeaf{"Cefadjsummarylinktype", cefadjentry.Cefadjsummarylinktype}
    cefadjentry.EntityData.Leafs["cefAdjSource"] = types.YLeaf{"Cefadjsource", cefadjentry.Cefadjsource}
    cefadjentry.EntityData.Leafs["cefAdjEncap"] = types.YLeaf{"Cefadjencap", cefadjentry.Cefadjencap}
    cefadjentry.EntityData.Leafs["cefAdjFixup"] = types.YLeaf{"Cefadjfixup", cefadjentry.Cefadjfixup}
    cefadjentry.EntityData.Leafs["cefAdjMTU"] = types.YLeaf{"Cefadjmtu", cefadjentry.Cefadjmtu}
    cefadjentry.EntityData.Leafs["cefAdjForwardingInfo"] = types.YLeaf{"Cefadjforwardinginfo", cefadjentry.Cefadjforwardinginfo}
    cefadjentry.EntityData.Leafs["cefAdjPkts"] = types.YLeaf{"Cefadjpkts", cefadjentry.Cefadjpkts}
    cefadjentry.EntityData.Leafs["cefAdjHCPkts"] = types.YLeaf{"Cefadjhcpkts", cefadjentry.Cefadjhcpkts}
    cefadjentry.EntityData.Leafs["cefAdjBytes"] = types.YLeaf{"Cefadjbytes", cefadjentry.Cefadjbytes}
    cefadjentry.EntityData.Leafs["cefAdjHCBytes"] = types.YLeaf{"Cefadjhcbytes", cefadjentry.Cefadjhcbytes}
    return &(cefadjentry.EntityData)
}

// CISCOCEFMIB_Ceffeselectiontable
// A list of forwarding element selection entries.
type CISCOCEFMIB_Ceffeselectiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contain a CEF
    // forwarding element selection list.  entPhysicalIndex is also an index for
    // this table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry.
    Ceffeselectionentry []CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry
}

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetEntityData() *types.CommonEntityData {
    ceffeselectiontable.EntityData.YFilter = ceffeselectiontable.YFilter
    ceffeselectiontable.EntityData.YangName = "cefFESelectionTable"
    ceffeselectiontable.EntityData.BundleName = "cisco_ios_xe"
    ceffeselectiontable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    ceffeselectiontable.EntityData.SegmentPath = "cefFESelectionTable"
    ceffeselectiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceffeselectiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceffeselectiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceffeselectiontable.EntityData.Children = make(map[string]types.YChild)
    ceffeselectiontable.EntityData.Children["cefFESelectionEntry"] = types.YChild{"Ceffeselectionentry", nil}
    for i := range ceffeselectiontable.Ceffeselectionentry {
        ceffeselectiontable.EntityData.Children[types.GetSegmentPath(&ceffeselectiontable.Ceffeselectionentry[i])] = types.YChild{"Ceffeselectionentry", &ceffeselectiontable.Ceffeselectionentry[i]}
    }
    ceffeselectiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceffeselectiontable.EntityData)
}

// CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry
// If CEF is enabled on the Managed device,
// each entry contain a CEF forwarding element
// selection list.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier used
    // to select a set of forwarding element lists. The type is string with
    // length: 1..32.
    Ceffeselectionname interface{}

    // This attribute is a key. Secondary index to identify a forwarding elements
    // List  in this Table. The type is interface{} with range: 1..2147483647.
    Ceffeselectionid interface{}

    // Special processing for a destination is indicated through the use of
    // special  forwarding element.   If the forwarding element list contains the 
    // special forwarding element, then this object  represents the type of
    // special forwarding element. The type is CefForwardingElementSpecialType.
    Ceffeselectionspecial interface{}

    // This object represent the MPLS Labels  associated with this forwarding
    // Element List.  The value of this object will be irrelevant and will be set
    // to zero length if the forwarding element list  doesn't contain a label
    // forwarding element. A zero  length label list will indicate that there is
    // no label forwarding element associated with this selection entry. The type
    // is string with length: 0..255.
    Ceffeselectionlabels interface{}

    // This object represent the link type for the adjacency associated with this
    // forwarding  Element List.  The value of this object will be irrelevant and
    // will be set to unknown(5) if the forwarding element list  doesn't contain
    // an adjacency forwarding element. The type is CefAdjLinkType.
    Ceffeselectionadjlinktype interface{}

    // This object represent the interface for the adjacency associated with this
    // forwarding  Element List.  The value of this object will be irrelevant and
    // will be set to zero if the forwarding element list doesn't  contain an
    // adjacency forwarding element. The type is interface{} with range:
    // 0..2147483647.
    Ceffeselectionadjinterface interface{}

    // This object represent the next hop address type for the adjacency
    // associated with this forwarding  Element List.  The value of this object
    // will be irrelevant and will be set to unknown(0) if the forwarding element
    // list  doesn't contain an adjacency forwarding element. The type is
    // InetAddressType.
    Ceffeselectionadjnexthopaddrtype interface{}

    // This object represent the next hop address for the adjacency associated
    // with this forwarding  Element List.  The value of this object will be
    // irrelevant and will be set to zero if the forwarding element list doesn't 
    // contain an adjacency forwarding element. The type is string with length:
    // 0..255.
    Ceffeselectionadjnexthopaddr interface{}

    // This object represent the connection id for the adjacency associated with
    // this forwarding  Element List.  The value of this object will be irrelevant
    // and will be set to zero if the forwarding element list doesn't  contain an
    // adjacency forwarding element.   In cases where cefFESelectionAdjLinkType,
    // interface  and the next hop address are not able to uniquely  define an
    // adjacency entry (e.g. ATM and Frame Relay Bundles), this object is a unique
    // identifier to differentiate between these adjacency entries. The type is
    // interface{} with range: 0..4294967295.
    Ceffeselectionadjconnid interface{}

    // This object represent the Vrf name for the lookup associated with this
    // forwarding  Element List.  The value of this object will be irrelevant and
    // will be set to a string containing the single octet 0x00 if the forwarding
    // element list  doesn't contain a lookup forwarding element. The type is
    // string with length: 0..31.
    Ceffeselectionvrfname interface{}

    // This object represent the weighting for  load balancing between multiple
    // Forwarding Element Lists. The value of this object will be zero if load
    // balancing is associated with this selection entry. The type is interface{}
    // with range: 0..4294967295.
    Ceffeselectionweight interface{}
}

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetEntityData() *types.CommonEntityData {
    ceffeselectionentry.EntityData.YFilter = ceffeselectionentry.YFilter
    ceffeselectionentry.EntityData.YangName = "cefFESelectionEntry"
    ceffeselectionentry.EntityData.BundleName = "cisco_ios_xe"
    ceffeselectionentry.EntityData.ParentYangName = "cefFESelectionTable"
    ceffeselectionentry.EntityData.SegmentPath = "cefFESelectionEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceffeselectionentry.Entphysicalindex) + "']" + "[cefFESelectionName='" + fmt.Sprintf("%v", ceffeselectionentry.Ceffeselectionname) + "']" + "[cefFESelectionId='" + fmt.Sprintf("%v", ceffeselectionentry.Ceffeselectionid) + "']"
    ceffeselectionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceffeselectionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceffeselectionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceffeselectionentry.EntityData.Children = make(map[string]types.YChild)
    ceffeselectionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceffeselectionentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceffeselectionentry.Entphysicalindex}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionName"] = types.YLeaf{"Ceffeselectionname", ceffeselectionentry.Ceffeselectionname}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionId"] = types.YLeaf{"Ceffeselectionid", ceffeselectionentry.Ceffeselectionid}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionSpecial"] = types.YLeaf{"Ceffeselectionspecial", ceffeselectionentry.Ceffeselectionspecial}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionLabels"] = types.YLeaf{"Ceffeselectionlabels", ceffeselectionentry.Ceffeselectionlabels}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionAdjLinkType"] = types.YLeaf{"Ceffeselectionadjlinktype", ceffeselectionentry.Ceffeselectionadjlinktype}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionAdjInterface"] = types.YLeaf{"Ceffeselectionadjinterface", ceffeselectionentry.Ceffeselectionadjinterface}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionAdjNextHopAddrType"] = types.YLeaf{"Ceffeselectionadjnexthopaddrtype", ceffeselectionentry.Ceffeselectionadjnexthopaddrtype}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionAdjNextHopAddr"] = types.YLeaf{"Ceffeselectionadjnexthopaddr", ceffeselectionentry.Ceffeselectionadjnexthopaddr}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionAdjConnId"] = types.YLeaf{"Ceffeselectionadjconnid", ceffeselectionentry.Ceffeselectionadjconnid}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionVrfName"] = types.YLeaf{"Ceffeselectionvrfname", ceffeselectionentry.Ceffeselectionvrfname}
    ceffeselectionentry.EntityData.Leafs["cefFESelectionWeight"] = types.YLeaf{"Ceffeselectionweight", ceffeselectionentry.Ceffeselectionweight}
    return &(ceffeselectionentry.EntityData)
}

// CISCOCEFMIB_Cefcfgtable
// This table contains global config parameter 
// of CEF on the Managed device.
type CISCOCEFMIB_Cefcfgtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the Managed device supports CEF,  each entry contains the CEF config 
    // parameter for the managed entity. A row may exist for each IP version type
    // (v4 and v6) depending upon the IP version supported on the device. 
    // entPhysicalIndex is also an index for this table which represents entities
    // of 'module' entPhysicalClass which are capable of running CEF. The type is
    // slice of CISCOCEFMIB_Cefcfgtable_Cefcfgentry.
    Cefcfgentry []CISCOCEFMIB_Cefcfgtable_Cefcfgentry
}

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetEntityData() *types.CommonEntityData {
    cefcfgtable.EntityData.YFilter = cefcfgtable.YFilter
    cefcfgtable.EntityData.YangName = "cefCfgTable"
    cefcfgtable.EntityData.BundleName = "cisco_ios_xe"
    cefcfgtable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefcfgtable.EntityData.SegmentPath = "cefCfgTable"
    cefcfgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfgtable.EntityData.Children = make(map[string]types.YChild)
    cefcfgtable.EntityData.Children["cefCfgEntry"] = types.YChild{"Cefcfgentry", nil}
    for i := range cefcfgtable.Cefcfgentry {
        cefcfgtable.EntityData.Children[types.GetSegmentPath(&cefcfgtable.Cefcfgentry[i])] = types.YChild{"Cefcfgentry", &cefcfgtable.Cefcfgentry[i]}
    }
    cefcfgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcfgtable.EntityData)
}

// CISCOCEFMIB_Cefcfgtable_Cefcfgentry
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
type CISCOCEFMIB_Cefcfgtable_Cefcfgentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // The desired state of CEF. The type is CefAdminStatus.
    Cefcfgadminstate interface{}

    // The current operational state of CEF.  If the cefCfgAdminState is
    // disabled(2), then cefOperState will eventually go to the down(2) state
    // unless some error has occurred.   If cefCfgAdminState is changed to
    // enabled(1) then  cefCfgOperState should change to up(1) only if the  CEF
    // entity is ready to forward the packets using  Cisco Express Forwarding
    // (CEF) else it should remain  in the down(2) state. The up(1) state for this
    // object  indicates that CEF entity is forwarding the packet using Cisco
    // Express Forwarding. The type is CefOperStatus.
    Cefcfgoperstate interface{}

    // The desired state of CEF distribution. The type is CefAdminStatus.
    Cefcfgdistributionadminstate interface{}

    // The current operational state of CEF distribution.  If the
    // cefCfgDistributionAdminState is disabled(2), then cefDistributionOperState
    // will eventually go to the down(2) state unless some error has occurred.   
    // If cefCfgDistributionAdminState is changed to enabled(1)  then
    // cefCfgDistributionOperState should change to up(1)  only if the CEF entity
    // is ready to forward the packets  using Distributed Cisco Express Forwarding
    // (dCEF) else  it should remain in the down(2) state. The up(1) state  for
    // this object indicates that CEF entity is forwarding the packet using
    // Distributed Cisco Express Forwarding. The type is CefOperStatus.
    Cefcfgdistributionoperstate interface{}

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
    Cefcfgaccountingmap interface{}

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
    // type is Cefcfgloadsharingalgorithm.
    Cefcfgloadsharingalgorithm interface{}

    // The Fixed ID associated with the managed object cefCfgLoadSharingAlgorithm.
    // The hash of this object value may be used by the Load Sharing Algorithm. 
    // The value of this object is not relevant and will be set to zero if the
    // value of managed object  cefCfgLoadSharingAlgorithm is set to none(1) or
    // original(2). The default value of this object is calculated by the device
    // at the time of initialization. The type is interface{} with range:
    // 0..4294967295.
    Cefcfgloadsharingid interface{}

    // The interval time over which the CEF traffic statistics are collected. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    Cefcfgtrafficstatsloadinterval interface{}

    // The frequency with which the line card sends the traffic load statistics to
    // the Router Processor.  Setting the value of this object to 0 will disable
    // the CEF traffic statistics collection. The type is interface{} with range:
    // 0..65535. Units are seconds.
    Cefcfgtrafficstatsupdaterate interface{}
}

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetEntityData() *types.CommonEntityData {
    cefcfgentry.EntityData.YFilter = cefcfgentry.YFilter
    cefcfgentry.EntityData.YangName = "cefCfgEntry"
    cefcfgentry.EntityData.BundleName = "cisco_ios_xe"
    cefcfgentry.EntityData.ParentYangName = "cefCfgTable"
    cefcfgentry.EntityData.SegmentPath = "cefCfgEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfgentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefcfgentry.Ceffibipversion) + "']"
    cefcfgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfgentry.EntityData.Children = make(map[string]types.YChild)
    cefcfgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfgentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcfgentry.Entphysicalindex}
    cefcfgentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefcfgentry.Ceffibipversion}
    cefcfgentry.EntityData.Leafs["cefCfgAdminState"] = types.YLeaf{"Cefcfgadminstate", cefcfgentry.Cefcfgadminstate}
    cefcfgentry.EntityData.Leafs["cefCfgOperState"] = types.YLeaf{"Cefcfgoperstate", cefcfgentry.Cefcfgoperstate}
    cefcfgentry.EntityData.Leafs["cefCfgDistributionAdminState"] = types.YLeaf{"Cefcfgdistributionadminstate", cefcfgentry.Cefcfgdistributionadminstate}
    cefcfgentry.EntityData.Leafs["cefCfgDistributionOperState"] = types.YLeaf{"Cefcfgdistributionoperstate", cefcfgentry.Cefcfgdistributionoperstate}
    cefcfgentry.EntityData.Leafs["cefCfgAccountingMap"] = types.YLeaf{"Cefcfgaccountingmap", cefcfgentry.Cefcfgaccountingmap}
    cefcfgentry.EntityData.Leafs["cefCfgLoadSharingAlgorithm"] = types.YLeaf{"Cefcfgloadsharingalgorithm", cefcfgentry.Cefcfgloadsharingalgorithm}
    cefcfgentry.EntityData.Leafs["cefCfgLoadSharingID"] = types.YLeaf{"Cefcfgloadsharingid", cefcfgentry.Cefcfgloadsharingid}
    cefcfgentry.EntityData.Leafs["cefCfgTrafficStatsLoadInterval"] = types.YLeaf{"Cefcfgtrafficstatsloadinterval", cefcfgentry.Cefcfgtrafficstatsloadinterval}
    cefcfgentry.EntityData.Leafs["cefCfgTrafficStatsUpdateRate"] = types.YLeaf{"Cefcfgtrafficstatsupdaterate", cefcfgentry.Cefcfgtrafficstatsupdaterate}
    return &(cefcfgentry.EntityData)
}

// CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm represents object cefLoadSharingID. 
type CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm string

const (
    CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm_none CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm = "none"

    CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm_original CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm = "original"

    CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm_tunnel CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm = "tunnel"

    CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm_universal CISCOCEFMIB_Cefcfgtable_Cefcfgentry_Cefcfgloadsharingalgorithm = "universal"
)

// CISCOCEFMIB_Cefresourcetable
// This table contains global resource 
// information of CEF on the Managed device.
type CISCOCEFMIB_Cefresourcetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the Managed device supports CEF, each entry contains the CEF Resource 
    // parameters for the managed entity.  entPhysicalIndex is also an index for
    // this table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_Cefresourcetable_Cefresourceentry.
    Cefresourceentry []CISCOCEFMIB_Cefresourcetable_Cefresourceentry
}

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetEntityData() *types.CommonEntityData {
    cefresourcetable.EntityData.YFilter = cefresourcetable.YFilter
    cefresourcetable.EntityData.YangName = "cefResourceTable"
    cefresourcetable.EntityData.BundleName = "cisco_ios_xe"
    cefresourcetable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefresourcetable.EntityData.SegmentPath = "cefResourceTable"
    cefresourcetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefresourcetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefresourcetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefresourcetable.EntityData.Children = make(map[string]types.YChild)
    cefresourcetable.EntityData.Children["cefResourceEntry"] = types.YChild{"Cefresourceentry", nil}
    for i := range cefresourcetable.Cefresourceentry {
        cefresourcetable.EntityData.Children[types.GetSegmentPath(&cefresourcetable.Cefresourceentry[i])] = types.YChild{"Cefresourceentry", &cefresourcetable.Cefresourceentry[i]}
    }
    cefresourcetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefresourcetable.EntityData)
}

// CISCOCEFMIB_Cefresourcetable_Cefresourceentry
// If the Managed device supports CEF,
// each entry contains the CEF Resource 
// parameters for the managed entity.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_Cefresourcetable_Cefresourceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // Indicates the number of bytes from the Processor Memory Pool that are
    // currently in use by CEF on the managed entity. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    Cefresourcememoryused interface{}

    // The CEF resource failure reason which may lead to CEF being disabled on the
    // managed entity. The type is CefFailureReason.
    Cefresourcefailurereason interface{}
}

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetEntityData() *types.CommonEntityData {
    cefresourceentry.EntityData.YFilter = cefresourceentry.YFilter
    cefresourceentry.EntityData.YangName = "cefResourceEntry"
    cefresourceentry.EntityData.BundleName = "cisco_ios_xe"
    cefresourceentry.EntityData.ParentYangName = "cefResourceTable"
    cefresourceentry.EntityData.SegmentPath = "cefResourceEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefresourceentry.Entphysicalindex) + "']"
    cefresourceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefresourceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefresourceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefresourceentry.EntityData.Children = make(map[string]types.YChild)
    cefresourceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefresourceentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefresourceentry.Entphysicalindex}
    cefresourceentry.EntityData.Leafs["cefResourceMemoryUsed"] = types.YLeaf{"Cefresourcememoryused", cefresourceentry.Cefresourcememoryused}
    cefresourceentry.EntityData.Leafs["cefResourceFailureReason"] = types.YLeaf{"Cefresourcefailurereason", cefresourceentry.Cefresourcefailurereason}
    return &(cefresourceentry.EntityData)
}

// CISCOCEFMIB_Cefinttable
// This Table contains interface specific
// information of CEF on the Managed
// device.
type CISCOCEFMIB_Cefinttable struct {
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
    // running CEF. The type is slice of CISCOCEFMIB_Cefinttable_Cefintentry.
    Cefintentry []CISCOCEFMIB_Cefinttable_Cefintentry
}

func (cefinttable *CISCOCEFMIB_Cefinttable) GetEntityData() *types.CommonEntityData {
    cefinttable.EntityData.YFilter = cefinttable.YFilter
    cefinttable.EntityData.YangName = "cefIntTable"
    cefinttable.EntityData.BundleName = "cisco_ios_xe"
    cefinttable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefinttable.EntityData.SegmentPath = "cefIntTable"
    cefinttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefinttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefinttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefinttable.EntityData.Children = make(map[string]types.YChild)
    cefinttable.EntityData.Children["cefIntEntry"] = types.YChild{"Cefintentry", nil}
    for i := range cefinttable.Cefintentry {
        cefinttable.EntityData.Children[types.GetSegmentPath(&cefinttable.Cefintentry[i])] = types.YChild{"Cefintentry", &cefinttable.Cefintentry[i]}
    }
    cefinttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefinttable.EntityData)
}

// CISCOCEFMIB_Cefinttable_Cefintentry
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
type CISCOCEFMIB_Cefinttable_Cefintentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The CEF switching State for the interface.  If CEF is enabled but
    // distributed CEF(dCEF) is disabled then CEF is in cefEnabled(1) state.  If
    // distributed CEF is enabled, then CEF is in  distCefEnabled(2) state. The
    // cefDisabled(3) state indicates that CEF is disabled.  The CEF switching
    // state is only applicable to the received packet on the interface. The type
    // is Cefintswitchingstate.
    Cefintswitchingstate interface{}

    // The status of load sharing on the interface.  perPacket(1) : Router to send
    // data packets                over successive equal-cost paths               
    // without regard to individual hosts                or user sessions. 
    // perDestination(2) : Router to use multiple, equal-cost                    
    // paths to achieve load sharing  Load sharing is enabled by default  for an
    // interface when CEF is enabled. The type is Cefintloadsharing.
    Cefintloadsharing interface{}

    // The CEF accounting mode for the interface. CEF prefix based non-recursive
    // accounting  on an interface can be configured to store  the stats for
    // non-recursive prefixes in a internal  or external bucket.  internal(1)  : 
    // Count input traffic in the nonrecursive                 internal bucket 
    // external(2)  :  Count input traffic in the nonrecursive                
    // external bucket  The value of this object will only be effective if  value
    // of the object cefAccountingMap is set to enable nonRecursive(1) accounting.
    // The type is Cefintnonrecursiveaccouting.
    Cefintnonrecursiveaccouting interface{}
}

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetEntityData() *types.CommonEntityData {
    cefintentry.EntityData.YFilter = cefintentry.YFilter
    cefintentry.EntityData.YangName = "cefIntEntry"
    cefintentry.EntityData.BundleName = "cisco_ios_xe"
    cefintentry.EntityData.ParentYangName = "cefIntTable"
    cefintentry.EntityData.SegmentPath = "cefIntEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefintentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefintentry.Ceffibipversion) + "']" + "[ifIndex='" + fmt.Sprintf("%v", cefintentry.Ifindex) + "']"
    cefintentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefintentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefintentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefintentry.EntityData.Children = make(map[string]types.YChild)
    cefintentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefintentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefintentry.Entphysicalindex}
    cefintentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefintentry.Ceffibipversion}
    cefintentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cefintentry.Ifindex}
    cefintentry.EntityData.Leafs["cefIntSwitchingState"] = types.YLeaf{"Cefintswitchingstate", cefintentry.Cefintswitchingstate}
    cefintentry.EntityData.Leafs["cefIntLoadSharing"] = types.YLeaf{"Cefintloadsharing", cefintentry.Cefintloadsharing}
    cefintentry.EntityData.Leafs["cefIntNonrecursiveAccouting"] = types.YLeaf{"Cefintnonrecursiveaccouting", cefintentry.Cefintnonrecursiveaccouting}
    return &(cefintentry.EntityData)
}

// CISCOCEFMIB_Cefinttable_Cefintentry_Cefintloadsharing represents for an interface when CEF is enabled.
type CISCOCEFMIB_Cefinttable_Cefintentry_Cefintloadsharing string

const (
    CISCOCEFMIB_Cefinttable_Cefintentry_Cefintloadsharing_perPacket CISCOCEFMIB_Cefinttable_Cefintentry_Cefintloadsharing = "perPacket"

    CISCOCEFMIB_Cefinttable_Cefintentry_Cefintloadsharing_perDestination CISCOCEFMIB_Cefinttable_Cefintentry_Cefintloadsharing = "perDestination"
)

// CISCOCEFMIB_Cefinttable_Cefintentry_Cefintnonrecursiveaccouting represents nonRecursive(1) accounting.
type CISCOCEFMIB_Cefinttable_Cefintentry_Cefintnonrecursiveaccouting string

const (
    CISCOCEFMIB_Cefinttable_Cefintentry_Cefintnonrecursiveaccouting_internal CISCOCEFMIB_Cefinttable_Cefintentry_Cefintnonrecursiveaccouting = "internal"

    CISCOCEFMIB_Cefinttable_Cefintentry_Cefintnonrecursiveaccouting_external CISCOCEFMIB_Cefinttable_Cefintentry_Cefintnonrecursiveaccouting = "external"
)

// CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate represents received packet on the interface.
type CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate string

const (
    CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate_cefEnabled CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate = "cefEnabled"

    CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate_distCefEnabled CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate = "distCefEnabled"

    CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate_cefDisabled CISCOCEFMIB_Cefinttable_Cefintentry_Cefintswitchingstate = "cefDisabled"
)

// CISCOCEFMIB_Cefpeertable
// Entity acting as RP (Routing Processor) keeps
// the CEF states for the line card entities and
// communicates with the line card entities using
// XDR. This Table contains the CEF information 
// related to peer entities on the managed device.
type CISCOCEFMIB_Cefpeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF
    // related attributes  associated with a CEF peer entity.  entPhysicalIndex
    // and entPeerPhysicalIndex are also indexes for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_Cefpeertable_Cefpeerentry.
    Cefpeerentry []CISCOCEFMIB_Cefpeertable_Cefpeerentry
}

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetEntityData() *types.CommonEntityData {
    cefpeertable.EntityData.YFilter = cefpeertable.YFilter
    cefpeertable.EntityData.YangName = "cefPeerTable"
    cefpeertable.EntityData.BundleName = "cisco_ios_xe"
    cefpeertable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefpeertable.EntityData.SegmentPath = "cefPeerTable"
    cefpeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefpeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefpeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefpeertable.EntityData.Children = make(map[string]types.YChild)
    cefpeertable.EntityData.Children["cefPeerEntry"] = types.YChild{"Cefpeerentry", nil}
    for i := range cefpeertable.Cefpeerentry {
        cefpeertable.EntityData.Children[types.GetSegmentPath(&cefpeertable.Cefpeerentry[i])] = types.YChild{"Cefpeerentry", &cefpeertable.Cefpeerentry[i]}
    }
    cefpeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefpeertable.EntityData)
}

// CISCOCEFMIB_Cefpeertable_Cefpeerentry
// If CEF is enabled on the Managed device,
// each entry contains the CEF related attributes 
// associated with a CEF peer entity.
// 
// entPhysicalIndex and entPeerPhysicalIndex are
// also indexes for this table which represents
// entities of 'module' entPhysicalClass which are
// capable of running CEF.
type CISCOCEFMIB_Cefpeertable_Cefpeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The entity index for the CEF peer entity. Only the
    // entities of 'module'  entPhysicalClass are included here. The type is
    // interface{} with range: 1..2147483647.
    Entpeerphysicalindex interface{}

    // The current CEF operational state of the CEF peer entity.  Cef peer entity
    // oper state will be peerDisabled(1) in  the following condition:     : Cef
    // Peer entity encounters fatal error i.e. resource      allocation failure,
    // ipc failure etc     : When a reload/delete request is received from the Cef
    // Peer Entity  Once the peer entity is up and no fatal error is encountered,
    // then the value of this object will transits to the peerUp(3)  state.  If
    // the Cef Peer entity is in held stage, then the value of this object will be
    // peerHold(3). Cef peer entity can only transit to peerDisabled(1) state from
    // the peerHold(3) state. The type is Cefpeeroperstate.
    Cefpeeroperstate interface{}

    // Number of times the session with CEF peer entity  has been reset. The type
    // is interface{} with range: 0..4294967295.
    Cefpeernumberofresets interface{}
}

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetEntityData() *types.CommonEntityData {
    cefpeerentry.EntityData.YFilter = cefpeerentry.YFilter
    cefpeerentry.EntityData.YangName = "cefPeerEntry"
    cefpeerentry.EntityData.BundleName = "cisco_ios_xe"
    cefpeerentry.EntityData.ParentYangName = "cefPeerTable"
    cefpeerentry.EntityData.SegmentPath = "cefPeerEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefpeerentry.Entphysicalindex) + "']" + "[entPeerPhysicalIndex='" + fmt.Sprintf("%v", cefpeerentry.Entpeerphysicalindex) + "']"
    cefpeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefpeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefpeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefpeerentry.EntityData.Children = make(map[string]types.YChild)
    cefpeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefpeerentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefpeerentry.Entphysicalindex}
    cefpeerentry.EntityData.Leafs["entPeerPhysicalIndex"] = types.YLeaf{"Entpeerphysicalindex", cefpeerentry.Entpeerphysicalindex}
    cefpeerentry.EntityData.Leafs["cefPeerOperState"] = types.YLeaf{"Cefpeeroperstate", cefpeerentry.Cefpeeroperstate}
    cefpeerentry.EntityData.Leafs["cefPeerNumberOfResets"] = types.YLeaf{"Cefpeernumberofresets", cefpeerentry.Cefpeernumberofresets}
    return &(cefpeerentry.EntityData)
}

// CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate represents transit to peerDisabled(1) state from the peerHold(3) state.
type CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate string

const (
    CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate_peerDisabled CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate = "peerDisabled"

    CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate_peerUp CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate = "peerUp"

    CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate_peerHold CISCOCEFMIB_Cefpeertable_Cefpeerentry_Cefpeeroperstate = "peerHold"
)

// CISCOCEFMIB_Cefpeerfibtable
// Entity acting as RP (Routing Processor) keep
// the CEF FIB states for the line card entities and
// communicate with the line card entities using
// XDR. This Table contains the CEF FIB State 
// related to peer entities on the managed device.
type CISCOCEFMIB_Cefpeerfibtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF FIB
    // State  associated a CEF peer entity.  entPhysicalIndex and
    // entPeerPhysicalIndex are also indexes for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry.
    Cefpeerfibentry []CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry
}

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetEntityData() *types.CommonEntityData {
    cefpeerfibtable.EntityData.YFilter = cefpeerfibtable.YFilter
    cefpeerfibtable.EntityData.YangName = "cefPeerFIBTable"
    cefpeerfibtable.EntityData.BundleName = "cisco_ios_xe"
    cefpeerfibtable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefpeerfibtable.EntityData.SegmentPath = "cefPeerFIBTable"
    cefpeerfibtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefpeerfibtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefpeerfibtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefpeerfibtable.EntityData.Children = make(map[string]types.YChild)
    cefpeerfibtable.EntityData.Children["cefPeerFIBEntry"] = types.YChild{"Cefpeerfibentry", nil}
    for i := range cefpeerfibtable.Cefpeerfibentry {
        cefpeerfibtable.EntityData.Children[types.GetSegmentPath(&cefpeerfibtable.Cefpeerfibentry[i])] = types.YChild{"Cefpeerfibentry", &cefpeerfibtable.Cefpeerfibentry[i]}
    }
    cefpeerfibtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefpeerfibtable.EntityData)
}

// CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry
// If CEF is enabled on the Managed device,
// each entry contains the CEF FIB State 
// associated a CEF peer entity.
// 
// entPhysicalIndex and entPeerPhysicalIndex are
// also indexes for this table which represents
// entities of 'module' entPhysicalClass which are
// capable of running CEF.
type CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_cef_mib.CISCOCEFMIB_Cefpeertable_Cefpeerentry_Entpeerphysicalindex
    Entpeerphysicalindex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // The current CEF FIB Operational State for the  CEF peer entity. The type is
    // Cefpeerfiboperstate.
    Cefpeerfiboperstate interface{}
}

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetEntityData() *types.CommonEntityData {
    cefpeerfibentry.EntityData.YFilter = cefpeerfibentry.YFilter
    cefpeerfibentry.EntityData.YangName = "cefPeerFIBEntry"
    cefpeerfibentry.EntityData.BundleName = "cisco_ios_xe"
    cefpeerfibentry.EntityData.ParentYangName = "cefPeerFIBTable"
    cefpeerfibentry.EntityData.SegmentPath = "cefPeerFIBEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefpeerfibentry.Entphysicalindex) + "']" + "[entPeerPhysicalIndex='" + fmt.Sprintf("%v", cefpeerfibentry.Entpeerphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefpeerfibentry.Ceffibipversion) + "']"
    cefpeerfibentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefpeerfibentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefpeerfibentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefpeerfibentry.EntityData.Children = make(map[string]types.YChild)
    cefpeerfibentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefpeerfibentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefpeerfibentry.Entphysicalindex}
    cefpeerfibentry.EntityData.Leafs["entPeerPhysicalIndex"] = types.YLeaf{"Entpeerphysicalindex", cefpeerfibentry.Entpeerphysicalindex}
    cefpeerfibentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefpeerfibentry.Ceffibipversion}
    cefpeerfibentry.EntityData.Leafs["cefPeerFIBOperState"] = types.YLeaf{"Cefpeerfiboperstate", cefpeerfibentry.Cefpeerfiboperstate}
    return &(cefpeerfibentry.EntityData)
}

// CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate represents CEF peer entity.
type CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate string

const (
    CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate_peerFIBDown CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate = "peerFIBDown"

    CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate_peerFIBUp CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate = "peerFIBUp"

    CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate_peerFIBReloadRequest CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate = "peerFIBReloadRequest"

    CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate_peerFIBReloading CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate = "peerFIBReloading"

    CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate_peerFIBSynced CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry_Cefpeerfiboperstate = "peerFIBSynced"
)

// CISCOCEFMIB_Cefccglobaltable
// This table contains CEF consistency checker
// (CC) global parameters for the managed device.
type CISCOCEFMIB_Cefccglobaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the global
    // consistency  checker parameter for the managed device. A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device. The type is slice of
    // CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry.
    Cefccglobalentry []CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry
}

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetEntityData() *types.CommonEntityData {
    cefccglobaltable.EntityData.YFilter = cefccglobaltable.YFilter
    cefccglobaltable.EntityData.YangName = "cefCCGlobalTable"
    cefccglobaltable.EntityData.BundleName = "cisco_ios_xe"
    cefccglobaltable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefccglobaltable.EntityData.SegmentPath = "cefCCGlobalTable"
    cefccglobaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefccglobaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefccglobaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefccglobaltable.EntityData.Children = make(map[string]types.YChild)
    cefccglobaltable.EntityData.Children["cefCCGlobalEntry"] = types.YChild{"Cefccglobalentry", nil}
    for i := range cefccglobaltable.Cefccglobalentry {
        cefccglobaltable.EntityData.Children[types.GetSegmentPath(&cefccglobaltable.Cefccglobalentry[i])] = types.YChild{"Cefccglobalentry", &cefccglobaltable.Cefccglobalentry[i]}
    }
    cefccglobaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefccglobaltable.EntityData)
}

// CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry
// If the managed device supports CEF,
// each entry contains the global consistency 
// checker parameter for the managed device.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
type CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // Once an inconsistency has been detected,  CEF has the ability to repair the
    // problem.  This object indicates the status of auto-repair  function for the
    // consistency checkers. The type is bool.
    Cefccglobalautorepairenabled interface{}

    // Indiactes how long the consistency checker  waits to fix an inconsistency. 
    // The value of this object has no effect when the value of object
    // cefCCGlobalAutoRepairEnabled is 'false'. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    Cefccglobalautorepairdelay interface{}

    // Indicates how long the consistency checker waits to re-enable auto-repair
    // after  auto-repair runs.  The value of this object has no effect when the
    // value of object cefCCGlobalAutoRepairEnabled is 'false'. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Cefccglobalautorepairholddown interface{}

    // Enables the consistency checker to generate  an error message when it
    // detects an inconsistency. The type is bool.
    Cefccglobalerrormsgenabled interface{}

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
    Cefccglobalfullscanaction interface{}

    // Indicates the status of the full scan consistency checker request. The type
    // is CefCCStatus.
    Cefccglobalfullscanstatus interface{}
}

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetEntityData() *types.CommonEntityData {
    cefccglobalentry.EntityData.YFilter = cefccglobalentry.YFilter
    cefccglobalentry.EntityData.YangName = "cefCCGlobalEntry"
    cefccglobalentry.EntityData.BundleName = "cisco_ios_xe"
    cefccglobalentry.EntityData.ParentYangName = "cefCCGlobalTable"
    cefccglobalentry.EntityData.SegmentPath = "cefCCGlobalEntry" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefccglobalentry.Ceffibipversion) + "']"
    cefccglobalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefccglobalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefccglobalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefccglobalentry.EntityData.Children = make(map[string]types.YChild)
    cefccglobalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefccglobalentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefccglobalentry.Ceffibipversion}
    cefccglobalentry.EntityData.Leafs["cefCCGlobalAutoRepairEnabled"] = types.YLeaf{"Cefccglobalautorepairenabled", cefccglobalentry.Cefccglobalautorepairenabled}
    cefccglobalentry.EntityData.Leafs["cefCCGlobalAutoRepairDelay"] = types.YLeaf{"Cefccglobalautorepairdelay", cefccglobalentry.Cefccglobalautorepairdelay}
    cefccglobalentry.EntityData.Leafs["cefCCGlobalAutoRepairHoldDown"] = types.YLeaf{"Cefccglobalautorepairholddown", cefccglobalentry.Cefccglobalautorepairholddown}
    cefccglobalentry.EntityData.Leafs["cefCCGlobalErrorMsgEnabled"] = types.YLeaf{"Cefccglobalerrormsgenabled", cefccglobalentry.Cefccglobalerrormsgenabled}
    cefccglobalentry.EntityData.Leafs["cefCCGlobalFullScanAction"] = types.YLeaf{"Cefccglobalfullscanaction", cefccglobalentry.Cefccglobalfullscanaction}
    cefccglobalentry.EntityData.Leafs["cefCCGlobalFullScanStatus"] = types.YLeaf{"Cefccglobalfullscanstatus", cefccglobalentry.Cefccglobalfullscanstatus}
    return &(cefccglobalentry.EntityData)
}

// CISCOCEFMIB_Cefcctypetable
// This table contains CEF consistency
// checker types specific parameters on the managed device.
// 
// All detected inconsistency are signaled to the
// Management Station via cefInconsistencyDetection
// notification.
type CISCOCEFMIB_Cefcctypetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the consistency 
    // checker statistics for a consistency  checker type. A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device. The type is slice of CISCOCEFMIB_Cefcctypetable_Cefcctypeentry.
    Cefcctypeentry []CISCOCEFMIB_Cefcctypetable_Cefcctypeentry
}

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetEntityData() *types.CommonEntityData {
    cefcctypetable.EntityData.YFilter = cefcctypetable.YFilter
    cefcctypetable.EntityData.YangName = "cefCCTypeTable"
    cefcctypetable.EntityData.BundleName = "cisco_ios_xe"
    cefcctypetable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefcctypetable.EntityData.SegmentPath = "cefCCTypeTable"
    cefcctypetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcctypetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcctypetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcctypetable.EntityData.Children = make(map[string]types.YChild)
    cefcctypetable.EntityData.Children["cefCCTypeEntry"] = types.YChild{"Cefcctypeentry", nil}
    for i := range cefcctypetable.Cefcctypeentry {
        cefcctypetable.EntityData.Children[types.GetSegmentPath(&cefcctypetable.Cefcctypeentry[i])] = types.YChild{"Cefcctypeentry", &cefcctypetable.Cefcctypeentry[i]}
    }
    cefcctypetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcctypetable.EntityData)
}

// CISCOCEFMIB_Cefcctypetable_Cefcctypeentry
// If the managed device supports CEF,
// each entry contains the consistency 
// checker statistics for a consistency 
// checker type.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
type CISCOCEFMIB_Cefcctypetable_Cefcctypeentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // This attribute is a key. Type of the consistency checker. The type is
    // CefCCType.
    Cefcctype interface{}

    // Enables the passive consistency checker. Passive consistency checkers are
    // disabled by default.  Full-scan consistency checkers are always enabled. An
    // attempt to set this object to 'false' for an active consistency checker
    // will result in 'wrongValue' error. The type is bool.
    Cefccenabled interface{}

    // The maximum number of prefixes to check per scan.  The default value for
    // this object  depends upon the consistency checker type.  The value of this
    // object will be irrelevant  for some of the consistency checkers and will be
    // set to 0.  A Management Station cannot set the value of this object to 0.
    // The type is interface{} with range: 0..4294967295.
    Cefcccount interface{}

    // The period between scans for the consistency checker. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Cefccperiod interface{}

    // Number of prefix consistency queries sent to CEF forwarding databases by
    // this consistency checker. The type is interface{} with range:
    // 0..4294967295.
    Cefccqueriessent interface{}

    // Number of prefix consistency queries for which the consistency checks were
    // not performed by this  consistency checker. This may be because of some
    // internal error or resource failure. The type is interface{} with range:
    // 0..4294967295.
    Cefccqueriesignored interface{}

    // Number of prefix consistency queries processed by this  consistency
    // checker. The type is interface{} with range: 0..4294967295.
    Cefccquerieschecked interface{}

    // Number of prefix consistency queries iterated back to the master database
    // by this consistency checker. The type is interface{} with range:
    // 0..4294967295.
    Cefccqueriesiterated interface{}
}

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetEntityData() *types.CommonEntityData {
    cefcctypeentry.EntityData.YFilter = cefcctypeentry.YFilter
    cefcctypeentry.EntityData.YangName = "cefCCTypeEntry"
    cefcctypeentry.EntityData.BundleName = "cisco_ios_xe"
    cefcctypeentry.EntityData.ParentYangName = "cefCCTypeTable"
    cefcctypeentry.EntityData.SegmentPath = "cefCCTypeEntry" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefcctypeentry.Ceffibipversion) + "']" + "[cefCCType='" + fmt.Sprintf("%v", cefcctypeentry.Cefcctype) + "']"
    cefcctypeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcctypeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcctypeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcctypeentry.EntityData.Children = make(map[string]types.YChild)
    cefcctypeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcctypeentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefcctypeentry.Ceffibipversion}
    cefcctypeentry.EntityData.Leafs["cefCCType"] = types.YLeaf{"Cefcctype", cefcctypeentry.Cefcctype}
    cefcctypeentry.EntityData.Leafs["cefCCEnabled"] = types.YLeaf{"Cefccenabled", cefcctypeentry.Cefccenabled}
    cefcctypeentry.EntityData.Leafs["cefCCCount"] = types.YLeaf{"Cefcccount", cefcctypeentry.Cefcccount}
    cefcctypeentry.EntityData.Leafs["cefCCPeriod"] = types.YLeaf{"Cefccperiod", cefcctypeentry.Cefccperiod}
    cefcctypeentry.EntityData.Leafs["cefCCQueriesSent"] = types.YLeaf{"Cefccqueriessent", cefcctypeentry.Cefccqueriessent}
    cefcctypeentry.EntityData.Leafs["cefCCQueriesIgnored"] = types.YLeaf{"Cefccqueriesignored", cefcctypeentry.Cefccqueriesignored}
    cefcctypeentry.EntityData.Leafs["cefCCQueriesChecked"] = types.YLeaf{"Cefccquerieschecked", cefcctypeentry.Cefccquerieschecked}
    cefcctypeentry.EntityData.Leafs["cefCCQueriesIterated"] = types.YLeaf{"Cefccqueriesiterated", cefcctypeentry.Cefccqueriesiterated}
    return &(cefcctypeentry.EntityData)
}

// CISCOCEFMIB_Cefinconsistencyrecordtable
// This table contains CEF inconsistency
// records.
type CISCOCEFMIB_Cefinconsistencyrecordtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the inconsistency 
    // record. The type is slice of
    // CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry.
    Cefinconsistencyrecordentry []CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry
}

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetEntityData() *types.CommonEntityData {
    cefinconsistencyrecordtable.EntityData.YFilter = cefinconsistencyrecordtable.YFilter
    cefinconsistencyrecordtable.EntityData.YangName = "cefInconsistencyRecordTable"
    cefinconsistencyrecordtable.EntityData.BundleName = "cisco_ios_xe"
    cefinconsistencyrecordtable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefinconsistencyrecordtable.EntityData.SegmentPath = "cefInconsistencyRecordTable"
    cefinconsistencyrecordtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefinconsistencyrecordtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefinconsistencyrecordtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefinconsistencyrecordtable.EntityData.Children = make(map[string]types.YChild)
    cefinconsistencyrecordtable.EntityData.Children["cefInconsistencyRecordEntry"] = types.YChild{"Cefinconsistencyrecordentry", nil}
    for i := range cefinconsistencyrecordtable.Cefinconsistencyrecordentry {
        cefinconsistencyrecordtable.EntityData.Children[types.GetSegmentPath(&cefinconsistencyrecordtable.Cefinconsistencyrecordentry[i])] = types.YChild{"Cefinconsistencyrecordentry", &cefinconsistencyrecordtable.Cefinconsistencyrecordentry[i]}
    }
    cefinconsistencyrecordtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefinconsistencyrecordtable.EntityData)
}

// CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry
// If the managed device supports CEF,
// each entry contains the inconsistency 
// record.
type CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this inconsistency record entry. The type is interface{}
    // with range: 1..2147483647.
    Cefinconsistencyrecid interface{}

    // The network prefix type associated with this inconsistency record. The type
    // is InetAddressType.
    Cefinconsistencyprefixtype interface{}

    // The network prefix address associated with this  inconsistency record.  The
    // type of this address is determined by the value of the
    // cefInconsistencyPrefixType object. The type is string with length: 0..255.
    Cefinconsistencyprefixaddr interface{}

    // Length in bits of the inconsistency address prefix. The type is interface{}
    // with range: 0..2040.
    Cefinconsistencyprefixlen interface{}

    // Vrf name associated with this inconsistency record. The type is string with
    // length: 0..31.
    Cefinconsistencyvrfname interface{}

    // The type of consistency checker who generated this inconsistency record.
    // The type is CefCCType.
    Cefinconsistencycctype interface{}

    // The entity for which this inconsistency record was  generated. The value of
    // this object will be  irrelevant and will be set to 0 when the inconsisency 
    // record is applicable for all the entities. The type is interface{} with
    // range: 0..2147483647.
    Cefinconsistencyentity interface{}

    // The reason for generating this inconsistency record.   missing(1):       
    // the prefix is missing  checksumErr(2):    checksum error was found 
    // unknown(3):        reason is unknown. The type is Cefinconsistencyreason.
    Cefinconsistencyreason interface{}
}

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetEntityData() *types.CommonEntityData {
    cefinconsistencyrecordentry.EntityData.YFilter = cefinconsistencyrecordentry.YFilter
    cefinconsistencyrecordentry.EntityData.YangName = "cefInconsistencyRecordEntry"
    cefinconsistencyrecordentry.EntityData.BundleName = "cisco_ios_xe"
    cefinconsistencyrecordentry.EntityData.ParentYangName = "cefInconsistencyRecordTable"
    cefinconsistencyrecordentry.EntityData.SegmentPath = "cefInconsistencyRecordEntry" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefinconsistencyrecordentry.Ceffibipversion) + "']" + "[cefInconsistencyRecId='" + fmt.Sprintf("%v", cefinconsistencyrecordentry.Cefinconsistencyrecid) + "']"
    cefinconsistencyrecordentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefinconsistencyrecordentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefinconsistencyrecordentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefinconsistencyrecordentry.EntityData.Children = make(map[string]types.YChild)
    cefinconsistencyrecordentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefinconsistencyrecordentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefinconsistencyrecordentry.Ceffibipversion}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyRecId"] = types.YLeaf{"Cefinconsistencyrecid", cefinconsistencyrecordentry.Cefinconsistencyrecid}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyPrefixType"] = types.YLeaf{"Cefinconsistencyprefixtype", cefinconsistencyrecordentry.Cefinconsistencyprefixtype}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyPrefixAddr"] = types.YLeaf{"Cefinconsistencyprefixaddr", cefinconsistencyrecordentry.Cefinconsistencyprefixaddr}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyPrefixLen"] = types.YLeaf{"Cefinconsistencyprefixlen", cefinconsistencyrecordentry.Cefinconsistencyprefixlen}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyVrfName"] = types.YLeaf{"Cefinconsistencyvrfname", cefinconsistencyrecordentry.Cefinconsistencyvrfname}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyCCType"] = types.YLeaf{"Cefinconsistencycctype", cefinconsistencyrecordentry.Cefinconsistencycctype}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyEntity"] = types.YLeaf{"Cefinconsistencyentity", cefinconsistencyrecordentry.Cefinconsistencyentity}
    cefinconsistencyrecordentry.EntityData.Leafs["cefInconsistencyReason"] = types.YLeaf{"Cefinconsistencyreason", cefinconsistencyrecordentry.Cefinconsistencyreason}
    return &(cefinconsistencyrecordentry.EntityData)
}

// CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason represents unknown(3):        reason is unknown
type CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason string

const (
    CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason_missing CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason = "missing"

    CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason_checksumErr CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason = "checksumErr"

    CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason_unknown CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry_Cefinconsistencyreason = "unknown"
)

// CISCOCEFMIB_Cefstatsprefixlentable
// This table specifies the CEF stats based
// on the Prefix Length.
type CISCOCEFMIB_Cefstatsprefixlentable struct {
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
    // CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry.
    Cefstatsprefixlenentry []CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry
}

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetEntityData() *types.CommonEntityData {
    cefstatsprefixlentable.EntityData.YFilter = cefstatsprefixlentable.YFilter
    cefstatsprefixlentable.EntityData.YangName = "cefStatsPrefixLenTable"
    cefstatsprefixlentable.EntityData.BundleName = "cisco_ios_xe"
    cefstatsprefixlentable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefstatsprefixlentable.EntityData.SegmentPath = "cefStatsPrefixLenTable"
    cefstatsprefixlentable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefstatsprefixlentable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefstatsprefixlentable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefstatsprefixlentable.EntityData.Children = make(map[string]types.YChild)
    cefstatsprefixlentable.EntityData.Children["cefStatsPrefixLenEntry"] = types.YChild{"Cefstatsprefixlenentry", nil}
    for i := range cefstatsprefixlentable.Cefstatsprefixlenentry {
        cefstatsprefixlentable.EntityData.Children[types.GetSegmentPath(&cefstatsprefixlentable.Cefstatsprefixlenentry[i])] = types.YChild{"Cefstatsprefixlenentry", &cefstatsprefixlentable.Cefstatsprefixlenentry[i]}
    }
    cefstatsprefixlentable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefstatsprefixlentable.EntityData)
}

// CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry
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
type CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // This attribute is a key. Length in bits of the Destination IP prefix. As
    // 0.0.0.0/0 is a valid prefix, hence  0 is a valid prefix length. The type is
    // interface{} with range: 0..2040.
    Cefstatsprefixlen interface{}

    // Number of queries received in the FIB database  for the specified IP prefix
    // length. The type is interface{} with range: 0..4294967295.
    Cefstatsprefixqueries interface{}

    // Number of queries received in the FIB database for the specified IP prefix
    // length. This object is a 64-bit version of  cefStatsPrefixQueries. The type
    // is interface{} with range: 0..18446744073709551615.
    Cefstatsprefixhcqueries interface{}

    // Number of insert operations performed to the FIB  database for the
    // specified IP prefix length. The type is interface{} with range:
    // 0..4294967295.
    Cefstatsprefixinserts interface{}

    // Number of insert operations performed to the FIB  database for the
    // specified IP prefix length. This object is a 64-bit version of 
    // cefStatsPrefixInsert. The type is interface{} with range:
    // 0..18446744073709551615.
    Cefstatsprefixhcinserts interface{}

    // Number of delete operations performed to the FIB  database for the
    // specified IP prefix length. The type is interface{} with range:
    // 0..4294967295.
    Cefstatsprefixdeletes interface{}

    // Number of delete operations performed to the FIB  database for the
    // specified IP prefix length. This object is a 64-bit version of 
    // cefStatsPrefixDelete. The type is interface{} with range:
    // 0..18446744073709551615.
    Cefstatsprefixhcdeletes interface{}

    // Total number of elements in the FIB database for the specified IP prefix
    // length. The type is interface{} with range: 0..4294967295.
    Cefstatsprefixelements interface{}

    // Total number of elements in the FIB database for the specified IP prefix
    // length. This object is a 64-bit version of  cefStatsPrefixElements. The
    // type is interface{} with range: 0..18446744073709551615.
    Cefstatsprefixhcelements interface{}
}

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetEntityData() *types.CommonEntityData {
    cefstatsprefixlenentry.EntityData.YFilter = cefstatsprefixlenentry.YFilter
    cefstatsprefixlenentry.EntityData.YangName = "cefStatsPrefixLenEntry"
    cefstatsprefixlenentry.EntityData.BundleName = "cisco_ios_xe"
    cefstatsprefixlenentry.EntityData.ParentYangName = "cefStatsPrefixLenTable"
    cefstatsprefixlenentry.EntityData.SegmentPath = "cefStatsPrefixLenEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefstatsprefixlenentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefstatsprefixlenentry.Ceffibipversion) + "']" + "[cefStatsPrefixLen='" + fmt.Sprintf("%v", cefstatsprefixlenentry.Cefstatsprefixlen) + "']"
    cefstatsprefixlenentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefstatsprefixlenentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefstatsprefixlenentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefstatsprefixlenentry.EntityData.Children = make(map[string]types.YChild)
    cefstatsprefixlenentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefstatsprefixlenentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefstatsprefixlenentry.Entphysicalindex}
    cefstatsprefixlenentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefstatsprefixlenentry.Ceffibipversion}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixLen"] = types.YLeaf{"Cefstatsprefixlen", cefstatsprefixlenentry.Cefstatsprefixlen}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixQueries"] = types.YLeaf{"Cefstatsprefixqueries", cefstatsprefixlenentry.Cefstatsprefixqueries}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixHCQueries"] = types.YLeaf{"Cefstatsprefixhcqueries", cefstatsprefixlenentry.Cefstatsprefixhcqueries}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixInserts"] = types.YLeaf{"Cefstatsprefixinserts", cefstatsprefixlenentry.Cefstatsprefixinserts}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixHCInserts"] = types.YLeaf{"Cefstatsprefixhcinserts", cefstatsprefixlenentry.Cefstatsprefixhcinserts}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixDeletes"] = types.YLeaf{"Cefstatsprefixdeletes", cefstatsprefixlenentry.Cefstatsprefixdeletes}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixHCDeletes"] = types.YLeaf{"Cefstatsprefixhcdeletes", cefstatsprefixlenentry.Cefstatsprefixhcdeletes}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixElements"] = types.YLeaf{"Cefstatsprefixelements", cefstatsprefixlenentry.Cefstatsprefixelements}
    cefstatsprefixlenentry.EntityData.Leafs["cefStatsPrefixHCElements"] = types.YLeaf{"Cefstatsprefixhcelements", cefstatsprefixlenentry.Cefstatsprefixhcelements}
    return &(cefstatsprefixlenentry.EntityData)
}

// CISCOCEFMIB_Cefswitchingstatstable
// This table specifies the CEF switch stats.
type CISCOCEFMIB_Cefswitchingstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry specifies the switching
    // stats. A row may exist for each IP version type (v4 and v6) depending upon
    // the IP version supported on the device.  entPhysicalIndex is also an index
    // for this table which represents entities of 'module' entPhysicalClass which
    // are capable of running CEF. The type is slice of
    // CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry.
    Cefswitchingstatsentry []CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry
}

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetEntityData() *types.CommonEntityData {
    cefswitchingstatstable.EntityData.YFilter = cefswitchingstatstable.YFilter
    cefswitchingstatstable.EntityData.YangName = "cefSwitchingStatsTable"
    cefswitchingstatstable.EntityData.BundleName = "cisco_ios_xe"
    cefswitchingstatstable.EntityData.ParentYangName = "CISCO-CEF-MIB"
    cefswitchingstatstable.EntityData.SegmentPath = "cefSwitchingStatsTable"
    cefswitchingstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefswitchingstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefswitchingstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefswitchingstatstable.EntityData.Children = make(map[string]types.YChild)
    cefswitchingstatstable.EntityData.Children["cefSwitchingStatsEntry"] = types.YChild{"Cefswitchingstatsentry", nil}
    for i := range cefswitchingstatstable.Cefswitchingstatsentry {
        cefswitchingstatstable.EntityData.Children[types.GetSegmentPath(&cefswitchingstatstable.Cefswitchingstatsentry[i])] = types.YChild{"Cefswitchingstatsentry", &cefswitchingstatstable.Cefswitchingstatsentry[i]}
    }
    cefswitchingstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefswitchingstatstable.EntityData)
}

// CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry
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
type CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is CefIpVersion. Refers to
    // cisco_cef_mib.CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry_Ceffibipversion
    Ceffibipversion interface{}

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this switching stats entry. The type is interface{} with
    // range: 1..2147483647.
    Cefswitchingindex interface{}

    // Switch path where the feature was executed. Available switch paths are
    // platform-dependent. Following are the examples of switching paths:     RIB
    // : switching with CEF assistance     Low-end switching (LES) : CEF switch
    // path     PAS : CEF turbo switch path. The type is string with length:
    // 1..32.
    Cefswitchingpath interface{}

    // Number of packets dropped by CEF. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Cefswitchingdrop interface{}

    // Number of packets dropped by CEF. This object is a 64-bit version of
    // cefSwitchingDrop. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Cefswitchinghcdrop interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the next-fastest switching vector. The type is interface{} with
    // range: 0..4294967295. Units are packets.
    Cefswitchingpunt interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the next-fastest switching vector. This object is a 64-bit
    // version of cefSwitchingPunt. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets.
    Cefswitchinghcpunt interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the host (process switching path).  For most of the switching
    // paths, the value of this object may be similar to cefSwitchingPunt. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cefswitchingpunt2Host interface{}

    // Number of packets that could not be switched in the normal path and were
    // punted to the host (process switching path).  For most of the switching
    // paths, the value of this object may be similar to cefSwitchingPunt. This
    // object is a 64-bit version of cefSwitchingPunt2Host. The type is
    // interface{} with range: 0..18446744073709551615. Units are packets.
    Cefswitchinghcpunt2Host interface{}
}

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetEntityData() *types.CommonEntityData {
    cefswitchingstatsentry.EntityData.YFilter = cefswitchingstatsentry.YFilter
    cefswitchingstatsentry.EntityData.YangName = "cefSwitchingStatsEntry"
    cefswitchingstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cefswitchingstatsentry.EntityData.ParentYangName = "cefSwitchingStatsTable"
    cefswitchingstatsentry.EntityData.SegmentPath = "cefSwitchingStatsEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefswitchingstatsentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefswitchingstatsentry.Ceffibipversion) + "']" + "[cefSwitchingIndex='" + fmt.Sprintf("%v", cefswitchingstatsentry.Cefswitchingindex) + "']"
    cefswitchingstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefswitchingstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefswitchingstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefswitchingstatsentry.EntityData.Children = make(map[string]types.YChild)
    cefswitchingstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefswitchingstatsentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefswitchingstatsentry.Entphysicalindex}
    cefswitchingstatsentry.EntityData.Leafs["cefFIBIpVersion"] = types.YLeaf{"Ceffibipversion", cefswitchingstatsentry.Ceffibipversion}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingIndex"] = types.YLeaf{"Cefswitchingindex", cefswitchingstatsentry.Cefswitchingindex}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingPath"] = types.YLeaf{"Cefswitchingpath", cefswitchingstatsentry.Cefswitchingpath}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingDrop"] = types.YLeaf{"Cefswitchingdrop", cefswitchingstatsentry.Cefswitchingdrop}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingHCDrop"] = types.YLeaf{"Cefswitchinghcdrop", cefswitchingstatsentry.Cefswitchinghcdrop}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingPunt"] = types.YLeaf{"Cefswitchingpunt", cefswitchingstatsentry.Cefswitchingpunt}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingHCPunt"] = types.YLeaf{"Cefswitchinghcpunt", cefswitchingstatsentry.Cefswitchinghcpunt}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingPunt2Host"] = types.YLeaf{"Cefswitchingpunt2Host", cefswitchingstatsentry.Cefswitchingpunt2Host}
    cefswitchingstatsentry.EntityData.Leafs["cefSwitchingHCPunt2Host"] = types.YLeaf{"Cefswitchinghcpunt2Host", cefswitchingstatsentry.Cefswitchinghcpunt2Host}
    return &(cefswitchingstatsentry.EntityData)
}

