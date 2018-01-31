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
    parent types.Entity
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

func (cISCOCEFMIB *CISCOCEFMIB) GetFilter() yfilter.YFilter { return cISCOCEFMIB.YFilter }

func (cISCOCEFMIB *CISCOCEFMIB) SetFilter(yf yfilter.YFilter) { cISCOCEFMIB.YFilter = yf }

func (cISCOCEFMIB *CISCOCEFMIB) GetGoName(yname string) string {
    if yname == "cefFIB" { return "Ceffib" }
    if yname == "cefCC" { return "Cefcc" }
    if yname == "cefNotifCntl" { return "Cefnotifcntl" }
    if yname == "cefFIBSummaryTable" { return "Ceffibsummarytable" }
    if yname == "cefPrefixTable" { return "Cefprefixtable" }
    if yname == "cefLMPrefixTable" { return "Ceflmprefixtable" }
    if yname == "cefPathTable" { return "Cefpathtable" }
    if yname == "cefAdjSummaryTable" { return "Cefadjsummarytable" }
    if yname == "cefAdjTable" { return "Cefadjtable" }
    if yname == "cefFESelectionTable" { return "Ceffeselectiontable" }
    if yname == "cefCfgTable" { return "Cefcfgtable" }
    if yname == "cefResourceTable" { return "Cefresourcetable" }
    if yname == "cefIntTable" { return "Cefinttable" }
    if yname == "cefPeerTable" { return "Cefpeertable" }
    if yname == "cefPeerFIBTable" { return "Cefpeerfibtable" }
    if yname == "cefCCGlobalTable" { return "Cefccglobaltable" }
    if yname == "cefCCTypeTable" { return "Cefcctypetable" }
    if yname == "cefInconsistencyRecordTable" { return "Cefinconsistencyrecordtable" }
    if yname == "cefStatsPrefixLenTable" { return "Cefstatsprefixlentable" }
    if yname == "cefSwitchingStatsTable" { return "Cefswitchingstatstable" }
    return ""
}

func (cISCOCEFMIB *CISCOCEFMIB) GetSegmentPath() string {
    return "CISCO-CEF-MIB:CISCO-CEF-MIB"
}

func (cISCOCEFMIB *CISCOCEFMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefFIB" {
        return &cISCOCEFMIB.Ceffib
    }
    if childYangName == "cefCC" {
        return &cISCOCEFMIB.Cefcc
    }
    if childYangName == "cefNotifCntl" {
        return &cISCOCEFMIB.Cefnotifcntl
    }
    if childYangName == "cefFIBSummaryTable" {
        return &cISCOCEFMIB.Ceffibsummarytable
    }
    if childYangName == "cefPrefixTable" {
        return &cISCOCEFMIB.Cefprefixtable
    }
    if childYangName == "cefLMPrefixTable" {
        return &cISCOCEFMIB.Ceflmprefixtable
    }
    if childYangName == "cefPathTable" {
        return &cISCOCEFMIB.Cefpathtable
    }
    if childYangName == "cefAdjSummaryTable" {
        return &cISCOCEFMIB.Cefadjsummarytable
    }
    if childYangName == "cefAdjTable" {
        return &cISCOCEFMIB.Cefadjtable
    }
    if childYangName == "cefFESelectionTable" {
        return &cISCOCEFMIB.Ceffeselectiontable
    }
    if childYangName == "cefCfgTable" {
        return &cISCOCEFMIB.Cefcfgtable
    }
    if childYangName == "cefResourceTable" {
        return &cISCOCEFMIB.Cefresourcetable
    }
    if childYangName == "cefIntTable" {
        return &cISCOCEFMIB.Cefinttable
    }
    if childYangName == "cefPeerTable" {
        return &cISCOCEFMIB.Cefpeertable
    }
    if childYangName == "cefPeerFIBTable" {
        return &cISCOCEFMIB.Cefpeerfibtable
    }
    if childYangName == "cefCCGlobalTable" {
        return &cISCOCEFMIB.Cefccglobaltable
    }
    if childYangName == "cefCCTypeTable" {
        return &cISCOCEFMIB.Cefcctypetable
    }
    if childYangName == "cefInconsistencyRecordTable" {
        return &cISCOCEFMIB.Cefinconsistencyrecordtable
    }
    if childYangName == "cefStatsPrefixLenTable" {
        return &cISCOCEFMIB.Cefstatsprefixlentable
    }
    if childYangName == "cefSwitchingStatsTable" {
        return &cISCOCEFMIB.Cefswitchingstatstable
    }
    return nil
}

func (cISCOCEFMIB *CISCOCEFMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cefFIB"] = &cISCOCEFMIB.Ceffib
    children["cefCC"] = &cISCOCEFMIB.Cefcc
    children["cefNotifCntl"] = &cISCOCEFMIB.Cefnotifcntl
    children["cefFIBSummaryTable"] = &cISCOCEFMIB.Ceffibsummarytable
    children["cefPrefixTable"] = &cISCOCEFMIB.Cefprefixtable
    children["cefLMPrefixTable"] = &cISCOCEFMIB.Ceflmprefixtable
    children["cefPathTable"] = &cISCOCEFMIB.Cefpathtable
    children["cefAdjSummaryTable"] = &cISCOCEFMIB.Cefadjsummarytable
    children["cefAdjTable"] = &cISCOCEFMIB.Cefadjtable
    children["cefFESelectionTable"] = &cISCOCEFMIB.Ceffeselectiontable
    children["cefCfgTable"] = &cISCOCEFMIB.Cefcfgtable
    children["cefResourceTable"] = &cISCOCEFMIB.Cefresourcetable
    children["cefIntTable"] = &cISCOCEFMIB.Cefinttable
    children["cefPeerTable"] = &cISCOCEFMIB.Cefpeertable
    children["cefPeerFIBTable"] = &cISCOCEFMIB.Cefpeerfibtable
    children["cefCCGlobalTable"] = &cISCOCEFMIB.Cefccglobaltable
    children["cefCCTypeTable"] = &cISCOCEFMIB.Cefcctypetable
    children["cefInconsistencyRecordTable"] = &cISCOCEFMIB.Cefinconsistencyrecordtable
    children["cefStatsPrefixLenTable"] = &cISCOCEFMIB.Cefstatsprefixlentable
    children["cefSwitchingStatsTable"] = &cISCOCEFMIB.Cefswitchingstatstable
    return children
}

func (cISCOCEFMIB *CISCOCEFMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOCEFMIB *CISCOCEFMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOCEFMIB *CISCOCEFMIB) GetYangName() string { return "CISCO-CEF-MIB" }

func (cISCOCEFMIB *CISCOCEFMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOCEFMIB *CISCOCEFMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOCEFMIB *CISCOCEFMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOCEFMIB *CISCOCEFMIB) SetParent(parent types.Entity) { cISCOCEFMIB.parent = parent }

func (cISCOCEFMIB *CISCOCEFMIB) GetParent() types.Entity { return cISCOCEFMIB.parent }

func (cISCOCEFMIB *CISCOCEFMIB) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Ceffib
type CISCOCEFMIB_Ceffib struct {
    parent types.Entity
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

func (ceffib *CISCOCEFMIB_Ceffib) GetFilter() yfilter.YFilter { return ceffib.YFilter }

func (ceffib *CISCOCEFMIB_Ceffib) SetFilter(yf yfilter.YFilter) { ceffib.YFilter = yf }

func (ceffib *CISCOCEFMIB_Ceffib) GetGoName(yname string) string {
    if yname == "cefLMPrefixSpinLock" { return "Ceflmprefixspinlock" }
    return ""
}

func (ceffib *CISCOCEFMIB_Ceffib) GetSegmentPath() string {
    return "cefFIB"
}

func (ceffib *CISCOCEFMIB_Ceffib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceffib *CISCOCEFMIB_Ceffib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceffib *CISCOCEFMIB_Ceffib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cefLMPrefixSpinLock"] = ceffib.Ceflmprefixspinlock
    return leafs
}

func (ceffib *CISCOCEFMIB_Ceffib) GetBundleName() string { return "cisco_ios_xe" }

func (ceffib *CISCOCEFMIB_Ceffib) GetYangName() string { return "cefFIB" }

func (ceffib *CISCOCEFMIB_Ceffib) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceffib *CISCOCEFMIB_Ceffib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceffib *CISCOCEFMIB_Ceffib) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceffib *CISCOCEFMIB_Ceffib) SetParent(parent types.Entity) { ceffib.parent = parent }

func (ceffib *CISCOCEFMIB_Ceffib) GetParent() types.Entity { return ceffib.parent }

func (ceffib *CISCOCEFMIB_Ceffib) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Cefcc
type CISCOCEFMIB_Cefcc struct {
    parent types.Entity
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

func (cefcc *CISCOCEFMIB_Cefcc) GetFilter() yfilter.YFilter { return cefcc.YFilter }

func (cefcc *CISCOCEFMIB_Cefcc) SetFilter(yf yfilter.YFilter) { cefcc.YFilter = yf }

func (cefcc *CISCOCEFMIB_Cefcc) GetGoName(yname string) string {
    if yname == "entLastInconsistencyDetectTime" { return "Entlastinconsistencydetecttime" }
    if yname == "cefInconsistencyReset" { return "Cefinconsistencyreset" }
    if yname == "cefInconsistencyResetStatus" { return "Cefinconsistencyresetstatus" }
    return ""
}

func (cefcc *CISCOCEFMIB_Cefcc) GetSegmentPath() string {
    return "cefCC"
}

func (cefcc *CISCOCEFMIB_Cefcc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcc *CISCOCEFMIB_Cefcc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcc *CISCOCEFMIB_Cefcc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entLastInconsistencyDetectTime"] = cefcc.Entlastinconsistencydetecttime
    leafs["cefInconsistencyReset"] = cefcc.Cefinconsistencyreset
    leafs["cefInconsistencyResetStatus"] = cefcc.Cefinconsistencyresetstatus
    return leafs
}

func (cefcc *CISCOCEFMIB_Cefcc) GetBundleName() string { return "cisco_ios_xe" }

func (cefcc *CISCOCEFMIB_Cefcc) GetYangName() string { return "cefCC" }

func (cefcc *CISCOCEFMIB_Cefcc) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcc *CISCOCEFMIB_Cefcc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcc *CISCOCEFMIB_Cefcc) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcc *CISCOCEFMIB_Cefcc) SetParent(parent types.Entity) { cefcc.parent = parent }

func (cefcc *CISCOCEFMIB_Cefcc) GetParent() types.Entity { return cefcc.parent }

func (cefcc *CISCOCEFMIB_Cefcc) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Cefnotifcntl
type CISCOCEFMIB_Cefnotifcntl struct {
    parent types.Entity
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

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetFilter() yfilter.YFilter { return cefnotifcntl.YFilter }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) SetFilter(yf yfilter.YFilter) { cefnotifcntl.YFilter = yf }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetGoName(yname string) string {
    if yname == "cefResourceFailureNotifEnable" { return "Cefresourcefailurenotifenable" }
    if yname == "cefPeerStateChangeNotifEnable" { return "Cefpeerstatechangenotifenable" }
    if yname == "cefPeerFIBStateChangeNotifEnable" { return "Cefpeerfibstatechangenotifenable" }
    if yname == "cefNotifThrottlingInterval" { return "Cefnotifthrottlinginterval" }
    if yname == "cefInconsistencyNotifEnable" { return "Cefinconsistencynotifenable" }
    return ""
}

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetSegmentPath() string {
    return "cefNotifCntl"
}

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cefResourceFailureNotifEnable"] = cefnotifcntl.Cefresourcefailurenotifenable
    leafs["cefPeerStateChangeNotifEnable"] = cefnotifcntl.Cefpeerstatechangenotifenable
    leafs["cefPeerFIBStateChangeNotifEnable"] = cefnotifcntl.Cefpeerfibstatechangenotifenable
    leafs["cefNotifThrottlingInterval"] = cefnotifcntl.Cefnotifthrottlinginterval
    leafs["cefInconsistencyNotifEnable"] = cefnotifcntl.Cefinconsistencynotifenable
    return leafs
}

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetBundleName() string { return "cisco_ios_xe" }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetYangName() string { return "cefNotifCntl" }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) SetParent(parent types.Entity) { cefnotifcntl.parent = parent }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetParent() types.Entity { return cefnotifcntl.parent }

func (cefnotifcntl *CISCOCEFMIB_Cefnotifcntl) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Ceffibsummarytable
// This table contains the summary information
// for the cefPrefixTable.
type CISCOCEFMIB_Ceffibsummarytable struct {
    parent types.Entity
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

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetFilter() yfilter.YFilter { return ceffibsummarytable.YFilter }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) SetFilter(yf yfilter.YFilter) { ceffibsummarytable.YFilter = yf }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetGoName(yname string) string {
    if yname == "cefFIBSummaryEntry" { return "Ceffibsummaryentry" }
    return ""
}

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetSegmentPath() string {
    return "cefFIBSummaryTable"
}

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefFIBSummaryEntry" {
        for _, c := range ceffibsummarytable.Ceffibsummaryentry {
            if ceffibsummarytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry{}
        ceffibsummarytable.Ceffibsummaryentry = append(ceffibsummarytable.Ceffibsummaryentry, child)
        return &ceffibsummarytable.Ceffibsummaryentry[len(ceffibsummarytable.Ceffibsummaryentry)-1]
    }
    return nil
}

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceffibsummarytable.Ceffibsummaryentry {
        children[ceffibsummarytable.Ceffibsummaryentry[i].GetSegmentPath()] = &ceffibsummarytable.Ceffibsummaryentry[i]
    }
    return children
}

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetBundleName() string { return "cisco_ios_xe" }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetYangName() string { return "cefFIBSummaryTable" }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) SetParent(parent types.Entity) { ceffibsummarytable.parent = parent }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetParent() types.Entity { return ceffibsummarytable.parent }

func (ceffibsummarytable *CISCOCEFMIB_Ceffibsummarytable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetFilter() yfilter.YFilter { return ceffibsummaryentry.YFilter }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) SetFilter(yf yfilter.YFilter) { ceffibsummaryentry.YFilter = yf }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefFIBSummaryFwdPrefixes" { return "Ceffibsummaryfwdprefixes" }
    return ""
}

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetSegmentPath() string {
    return "cefFIBSummaryEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceffibsummaryentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", ceffibsummaryentry.Ceffibipversion) + "']"
}

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceffibsummaryentry.Entphysicalindex
    leafs["cefFIBIpVersion"] = ceffibsummaryentry.Ceffibipversion
    leafs["cefFIBSummaryFwdPrefixes"] = ceffibsummaryentry.Ceffibsummaryfwdprefixes
    return leafs
}

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetYangName() string { return "cefFIBSummaryEntry" }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) SetParent(parent types.Entity) { ceffibsummaryentry.parent = parent }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetParent() types.Entity { return ceffibsummaryentry.parent }

func (ceffibsummaryentry *CISCOCEFMIB_Ceffibsummarytable_Ceffibsummaryentry) GetParentYangName() string { return "cefFIBSummaryTable" }

// CISCOCEFMIB_Cefprefixtable
// A list of CEF forwarding prefixes.
type CISCOCEFMIB_Cefprefixtable struct {
    parent types.Entity
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

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetFilter() yfilter.YFilter { return cefprefixtable.YFilter }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) SetFilter(yf yfilter.YFilter) { cefprefixtable.YFilter = yf }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetGoName(yname string) string {
    if yname == "cefPrefixEntry" { return "Cefprefixentry" }
    return ""
}

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetSegmentPath() string {
    return "cefPrefixTable"
}

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefPrefixEntry" {
        for _, c := range cefprefixtable.Cefprefixentry {
            if cefprefixtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefprefixtable_Cefprefixentry{}
        cefprefixtable.Cefprefixentry = append(cefprefixtable.Cefprefixentry, child)
        return &cefprefixtable.Cefprefixentry[len(cefprefixtable.Cefprefixentry)-1]
    }
    return nil
}

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefprefixtable.Cefprefixentry {
        children[cefprefixtable.Cefprefixentry[i].GetSegmentPath()] = &cefprefixtable.Cefprefixentry[i]
    }
    return children
}

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetYangName() string { return "cefPrefixTable" }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) SetParent(parent types.Entity) { cefprefixtable.parent = parent }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetParent() types.Entity { return cefprefixtable.parent }

func (cefprefixtable *CISCOCEFMIB_Cefprefixtable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetFilter() yfilter.YFilter { return cefprefixentry.YFilter }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) SetFilter(yf yfilter.YFilter) { cefprefixentry.YFilter = yf }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefPrefixType" { return "Cefprefixtype" }
    if yname == "cefPrefixAddr" { return "Cefprefixaddr" }
    if yname == "cefPrefixLen" { return "Cefprefixlen" }
    if yname == "cefPrefixForwardingInfo" { return "Cefprefixforwardinginfo" }
    if yname == "cefPrefixPkts" { return "Cefprefixpkts" }
    if yname == "cefPrefixHCPkts" { return "Cefprefixhcpkts" }
    if yname == "cefPrefixBytes" { return "Cefprefixbytes" }
    if yname == "cefPrefixHCBytes" { return "Cefprefixhcbytes" }
    if yname == "cefPrefixInternalNRPkts" { return "Cefprefixinternalnrpkts" }
    if yname == "cefPrefixInternalNRHCPkts" { return "Cefprefixinternalnrhcpkts" }
    if yname == "cefPrefixInternalNRBytes" { return "Cefprefixinternalnrbytes" }
    if yname == "cefPrefixInternalNRHCBytes" { return "Cefprefixinternalnrhcbytes" }
    if yname == "cefPrefixExternalNRPkts" { return "Cefprefixexternalnrpkts" }
    if yname == "cefPrefixExternalNRHCPkts" { return "Cefprefixexternalnrhcpkts" }
    if yname == "cefPrefixExternalNRBytes" { return "Cefprefixexternalnrbytes" }
    if yname == "cefPrefixExternalNRHCBytes" { return "Cefprefixexternalnrhcbytes" }
    return ""
}

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetSegmentPath() string {
    return "cefPrefixEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefprefixentry.Entphysicalindex) + "']" + "[cefPrefixType='" + fmt.Sprintf("%v", cefprefixentry.Cefprefixtype) + "']" + "[cefPrefixAddr='" + fmt.Sprintf("%v", cefprefixentry.Cefprefixaddr) + "']" + "[cefPrefixLen='" + fmt.Sprintf("%v", cefprefixentry.Cefprefixlen) + "']"
}

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefprefixentry.Entphysicalindex
    leafs["cefPrefixType"] = cefprefixentry.Cefprefixtype
    leafs["cefPrefixAddr"] = cefprefixentry.Cefprefixaddr
    leafs["cefPrefixLen"] = cefprefixentry.Cefprefixlen
    leafs["cefPrefixForwardingInfo"] = cefprefixentry.Cefprefixforwardinginfo
    leafs["cefPrefixPkts"] = cefprefixentry.Cefprefixpkts
    leafs["cefPrefixHCPkts"] = cefprefixentry.Cefprefixhcpkts
    leafs["cefPrefixBytes"] = cefprefixentry.Cefprefixbytes
    leafs["cefPrefixHCBytes"] = cefprefixentry.Cefprefixhcbytes
    leafs["cefPrefixInternalNRPkts"] = cefprefixentry.Cefprefixinternalnrpkts
    leafs["cefPrefixInternalNRHCPkts"] = cefprefixentry.Cefprefixinternalnrhcpkts
    leafs["cefPrefixInternalNRBytes"] = cefprefixentry.Cefprefixinternalnrbytes
    leafs["cefPrefixInternalNRHCBytes"] = cefprefixentry.Cefprefixinternalnrhcbytes
    leafs["cefPrefixExternalNRPkts"] = cefprefixentry.Cefprefixexternalnrpkts
    leafs["cefPrefixExternalNRHCPkts"] = cefprefixentry.Cefprefixexternalnrhcpkts
    leafs["cefPrefixExternalNRBytes"] = cefprefixentry.Cefprefixexternalnrbytes
    leafs["cefPrefixExternalNRHCBytes"] = cefprefixentry.Cefprefixexternalnrhcbytes
    return leafs
}

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetYangName() string { return "cefPrefixEntry" }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) SetParent(parent types.Entity) { cefprefixentry.parent = parent }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetParent() types.Entity { return cefprefixentry.parent }

func (cefprefixentry *CISCOCEFMIB_Cefprefixtable_Cefprefixentry) GetParentYangName() string { return "cefPrefixTable" }

// CISCOCEFMIB_Ceflmprefixtable
// A table of Longest Match Prefix Query requests.
// 
// Generator application should utilize the
// cefLMPrefixSpinLock to try to avoid collisions.
// See DESCRIPTION clause of cefLMPrefixSpinLock.
type CISCOCEFMIB_Ceflmprefixtable struct {
    parent types.Entity
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

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetFilter() yfilter.YFilter { return ceflmprefixtable.YFilter }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) SetFilter(yf yfilter.YFilter) { ceflmprefixtable.YFilter = yf }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetGoName(yname string) string {
    if yname == "cefLMPrefixEntry" { return "Ceflmprefixentry" }
    return ""
}

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetSegmentPath() string {
    return "cefLMPrefixTable"
}

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefLMPrefixEntry" {
        for _, c := range ceflmprefixtable.Ceflmprefixentry {
            if ceflmprefixtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry{}
        ceflmprefixtable.Ceflmprefixentry = append(ceflmprefixtable.Ceflmprefixentry, child)
        return &ceflmprefixtable.Ceflmprefixentry[len(ceflmprefixtable.Ceflmprefixentry)-1]
    }
    return nil
}

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceflmprefixtable.Ceflmprefixentry {
        children[ceflmprefixtable.Ceflmprefixentry[i].GetSegmentPath()] = &ceflmprefixtable.Ceflmprefixentry[i]
    }
    return children
}

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetBundleName() string { return "cisco_ios_xe" }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetYangName() string { return "cefLMPrefixTable" }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) SetParent(parent types.Entity) { ceflmprefixtable.parent = parent }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetParent() types.Entity { return ceflmprefixtable.parent }

func (ceflmprefixtable *CISCOCEFMIB_Ceflmprefixtable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetFilter() yfilter.YFilter { return ceflmprefixentry.YFilter }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) SetFilter(yf yfilter.YFilter) { ceflmprefixentry.YFilter = yf }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefLMPrefixDestAddrType" { return "Ceflmprefixdestaddrtype" }
    if yname == "cefLMPrefixDestAddr" { return "Ceflmprefixdestaddr" }
    if yname == "cefLMPrefixState" { return "Ceflmprefixstate" }
    if yname == "cefLMPrefixAddr" { return "Ceflmprefixaddr" }
    if yname == "cefLMPrefixLen" { return "Ceflmprefixlen" }
    if yname == "cefLMPrefixRowStatus" { return "Ceflmprefixrowstatus" }
    return ""
}

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetSegmentPath() string {
    return "cefLMPrefixEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceflmprefixentry.Entphysicalindex) + "']" + "[cefLMPrefixDestAddrType='" + fmt.Sprintf("%v", ceflmprefixentry.Ceflmprefixdestaddrtype) + "']" + "[cefLMPrefixDestAddr='" + fmt.Sprintf("%v", ceflmprefixentry.Ceflmprefixdestaddr) + "']"
}

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceflmprefixentry.Entphysicalindex
    leafs["cefLMPrefixDestAddrType"] = ceflmprefixentry.Ceflmprefixdestaddrtype
    leafs["cefLMPrefixDestAddr"] = ceflmprefixentry.Ceflmprefixdestaddr
    leafs["cefLMPrefixState"] = ceflmprefixentry.Ceflmprefixstate
    leafs["cefLMPrefixAddr"] = ceflmprefixentry.Ceflmprefixaddr
    leafs["cefLMPrefixLen"] = ceflmprefixentry.Ceflmprefixlen
    leafs["cefLMPrefixRowStatus"] = ceflmprefixentry.Ceflmprefixrowstatus
    return leafs
}

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetYangName() string { return "cefLMPrefixEntry" }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) SetParent(parent types.Entity) { ceflmprefixentry.parent = parent }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetParent() types.Entity { return ceflmprefixentry.parent }

func (ceflmprefixentry *CISCOCEFMIB_Ceflmprefixtable_Ceflmprefixentry) GetParentYangName() string { return "cefLMPrefixTable" }

// CISCOCEFMIB_Cefpathtable
// CEF prefix path is a valid route to reach to a 
// destination IP prefix. Multiple paths may exist 
// out of a router to the same destination prefix. 
// This table specify lists of CEF paths.
type CISCOCEFMIB_Cefpathtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contain a CEF prefix
    // path.  entPhysicalIndex is also an index for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_Cefpathtable_Cefpathentry.
    Cefpathentry []CISCOCEFMIB_Cefpathtable_Cefpathentry
}

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetFilter() yfilter.YFilter { return cefpathtable.YFilter }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) SetFilter(yf yfilter.YFilter) { cefpathtable.YFilter = yf }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetGoName(yname string) string {
    if yname == "cefPathEntry" { return "Cefpathentry" }
    return ""
}

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetSegmentPath() string {
    return "cefPathTable"
}

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefPathEntry" {
        for _, c := range cefpathtable.Cefpathentry {
            if cefpathtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefpathtable_Cefpathentry{}
        cefpathtable.Cefpathentry = append(cefpathtable.Cefpathentry, child)
        return &cefpathtable.Cefpathentry[len(cefpathtable.Cefpathentry)-1]
    }
    return nil
}

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefpathtable.Cefpathentry {
        children[cefpathtable.Cefpathentry[i].GetSegmentPath()] = &cefpathtable.Cefpathentry[i]
    }
    return children
}

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetYangName() string { return "cefPathTable" }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) SetParent(parent types.Entity) { cefpathtable.parent = parent }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetParent() types.Entity { return cefpathtable.parent }

func (cefpathtable *CISCOCEFMIB_Cefpathtable) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Cefpathtable_Cefpathentry
// If CEF is enabled on the Managed device,
// each entry contain a CEF prefix path.
// 
// entPhysicalIndex is also an index for this
// table which represents entities of
// 'module' entPhysicalClass which are capable
// of running CEF.
type CISCOCEFMIB_Cefpathtable_Cefpathentry struct {
    parent types.Entity
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

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetFilter() yfilter.YFilter { return cefpathentry.YFilter }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) SetFilter(yf yfilter.YFilter) { cefpathentry.YFilter = yf }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefPrefixType" { return "Cefprefixtype" }
    if yname == "cefPrefixAddr" { return "Cefprefixaddr" }
    if yname == "cefPrefixLen" { return "Cefprefixlen" }
    if yname == "cefPathId" { return "Cefpathid" }
    if yname == "cefPathType" { return "Cefpathtype" }
    if yname == "cefPathInterface" { return "Cefpathinterface" }
    if yname == "cefPathNextHopAddr" { return "Cefpathnexthopaddr" }
    if yname == "cefPathRecurseVrfName" { return "Cefpathrecursevrfname" }
    return ""
}

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetSegmentPath() string {
    return "cefPathEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefpathentry.Entphysicalindex) + "']" + "[cefPrefixType='" + fmt.Sprintf("%v", cefpathentry.Cefprefixtype) + "']" + "[cefPrefixAddr='" + fmt.Sprintf("%v", cefpathentry.Cefprefixaddr) + "']" + "[cefPrefixLen='" + fmt.Sprintf("%v", cefpathentry.Cefprefixlen) + "']" + "[cefPathId='" + fmt.Sprintf("%v", cefpathentry.Cefpathid) + "']"
}

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefpathentry.Entphysicalindex
    leafs["cefPrefixType"] = cefpathentry.Cefprefixtype
    leafs["cefPrefixAddr"] = cefpathentry.Cefprefixaddr
    leafs["cefPrefixLen"] = cefpathentry.Cefprefixlen
    leafs["cefPathId"] = cefpathentry.Cefpathid
    leafs["cefPathType"] = cefpathentry.Cefpathtype
    leafs["cefPathInterface"] = cefpathentry.Cefpathinterface
    leafs["cefPathNextHopAddr"] = cefpathentry.Cefpathnexthopaddr
    leafs["cefPathRecurseVrfName"] = cefpathentry.Cefpathrecursevrfname
    return leafs
}

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetYangName() string { return "cefPathEntry" }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) SetParent(parent types.Entity) { cefpathentry.parent = parent }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetParent() types.Entity { return cefpathentry.parent }

func (cefpathentry *CISCOCEFMIB_Cefpathtable_Cefpathentry) GetParentYangName() string { return "cefPathTable" }

// CISCOCEFMIB_Cefadjsummarytable
// This table contains the summary information
// for the cefAdjTable.
type CISCOCEFMIB_Cefadjsummarytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF
    // Adjacency   summary related attributes for the Managed entity. A row exists
    // for each adjacency link type.  entPhysicalIndex is also an index for this
    // table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry.
    Cefadjsummaryentry []CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry
}

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetFilter() yfilter.YFilter { return cefadjsummarytable.YFilter }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) SetFilter(yf yfilter.YFilter) { cefadjsummarytable.YFilter = yf }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetGoName(yname string) string {
    if yname == "cefAdjSummaryEntry" { return "Cefadjsummaryentry" }
    return ""
}

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetSegmentPath() string {
    return "cefAdjSummaryTable"
}

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefAdjSummaryEntry" {
        for _, c := range cefadjsummarytable.Cefadjsummaryentry {
            if cefadjsummarytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry{}
        cefadjsummarytable.Cefadjsummaryentry = append(cefadjsummarytable.Cefadjsummaryentry, child)
        return &cefadjsummarytable.Cefadjsummaryentry[len(cefadjsummarytable.Cefadjsummaryentry)-1]
    }
    return nil
}

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefadjsummarytable.Cefadjsummaryentry {
        children[cefadjsummarytable.Cefadjsummaryentry[i].GetSegmentPath()] = &cefadjsummarytable.Cefadjsummaryentry[i]
    }
    return children
}

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetBundleName() string { return "cisco_ios_xe" }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetYangName() string { return "cefAdjSummaryTable" }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) SetParent(parent types.Entity) { cefadjsummarytable.parent = parent }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetParent() types.Entity { return cefadjsummarytable.parent }

func (cefadjsummarytable *CISCOCEFMIB_Cefadjsummarytable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetFilter() yfilter.YFilter { return cefadjsummaryentry.YFilter }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) SetFilter(yf yfilter.YFilter) { cefadjsummaryentry.YFilter = yf }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefAdjSummaryLinkType" { return "Cefadjsummarylinktype" }
    if yname == "cefAdjSummaryComplete" { return "Cefadjsummarycomplete" }
    if yname == "cefAdjSummaryIncomplete" { return "Cefadjsummaryincomplete" }
    if yname == "cefAdjSummaryFixup" { return "Cefadjsummaryfixup" }
    if yname == "cefAdjSummaryRedirect" { return "Cefadjsummaryredirect" }
    return ""
}

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetSegmentPath() string {
    return "cefAdjSummaryEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefadjsummaryentry.Entphysicalindex) + "']" + "[cefAdjSummaryLinkType='" + fmt.Sprintf("%v", cefadjsummaryentry.Cefadjsummarylinktype) + "']"
}

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefadjsummaryentry.Entphysicalindex
    leafs["cefAdjSummaryLinkType"] = cefadjsummaryentry.Cefadjsummarylinktype
    leafs["cefAdjSummaryComplete"] = cefadjsummaryentry.Cefadjsummarycomplete
    leafs["cefAdjSummaryIncomplete"] = cefadjsummaryentry.Cefadjsummaryincomplete
    leafs["cefAdjSummaryFixup"] = cefadjsummaryentry.Cefadjsummaryfixup
    leafs["cefAdjSummaryRedirect"] = cefadjsummaryentry.Cefadjsummaryredirect
    return leafs
}

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetYangName() string { return "cefAdjSummaryEntry" }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) SetParent(parent types.Entity) { cefadjsummaryentry.parent = parent }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetParent() types.Entity { return cefadjsummaryentry.parent }

func (cefadjsummaryentry *CISCOCEFMIB_Cefadjsummarytable_Cefadjsummaryentry) GetParentYangName() string { return "cefAdjSummaryTable" }

// CISCOCEFMIB_Cefadjtable
// A list of CEF adjacencies.
type CISCOCEFMIB_Cefadjtable struct {
    parent types.Entity
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

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetFilter() yfilter.YFilter { return cefadjtable.YFilter }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) SetFilter(yf yfilter.YFilter) { cefadjtable.YFilter = yf }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetGoName(yname string) string {
    if yname == "cefAdjEntry" { return "Cefadjentry" }
    return ""
}

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetSegmentPath() string {
    return "cefAdjTable"
}

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefAdjEntry" {
        for _, c := range cefadjtable.Cefadjentry {
            if cefadjtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefadjtable_Cefadjentry{}
        cefadjtable.Cefadjentry = append(cefadjtable.Cefadjentry, child)
        return &cefadjtable.Cefadjentry[len(cefadjtable.Cefadjentry)-1]
    }
    return nil
}

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefadjtable.Cefadjentry {
        children[cefadjtable.Cefadjentry[i].GetSegmentPath()] = &cefadjtable.Cefadjentry[i]
    }
    return children
}

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetYangName() string { return "cefAdjTable" }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) SetParent(parent types.Entity) { cefadjtable.parent = parent }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetParent() types.Entity { return cefadjtable.parent }

func (cefadjtable *CISCOCEFMIB_Cefadjtable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetFilter() yfilter.YFilter { return cefadjentry.YFilter }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) SetFilter(yf yfilter.YFilter) { cefadjentry.YFilter = yf }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cefAdjNextHopAddrType" { return "Cefadjnexthopaddrtype" }
    if yname == "cefAdjNextHopAddr" { return "Cefadjnexthopaddr" }
    if yname == "cefAdjConnId" { return "Cefadjconnid" }
    if yname == "cefAdjSummaryLinkType" { return "Cefadjsummarylinktype" }
    if yname == "cefAdjSource" { return "Cefadjsource" }
    if yname == "cefAdjEncap" { return "Cefadjencap" }
    if yname == "cefAdjFixup" { return "Cefadjfixup" }
    if yname == "cefAdjMTU" { return "Cefadjmtu" }
    if yname == "cefAdjForwardingInfo" { return "Cefadjforwardinginfo" }
    if yname == "cefAdjPkts" { return "Cefadjpkts" }
    if yname == "cefAdjHCPkts" { return "Cefadjhcpkts" }
    if yname == "cefAdjBytes" { return "Cefadjbytes" }
    if yname == "cefAdjHCBytes" { return "Cefadjhcbytes" }
    return ""
}

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetSegmentPath() string {
    return "cefAdjEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefadjentry.Entphysicalindex) + "']" + "[ifIndex='" + fmt.Sprintf("%v", cefadjentry.Ifindex) + "']" + "[cefAdjNextHopAddrType='" + fmt.Sprintf("%v", cefadjentry.Cefadjnexthopaddrtype) + "']" + "[cefAdjNextHopAddr='" + fmt.Sprintf("%v", cefadjentry.Cefadjnexthopaddr) + "']" + "[cefAdjConnId='" + fmt.Sprintf("%v", cefadjentry.Cefadjconnid) + "']" + "[cefAdjSummaryLinkType='" + fmt.Sprintf("%v", cefadjentry.Cefadjsummarylinktype) + "']"
}

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefadjentry.Entphysicalindex
    leafs["ifIndex"] = cefadjentry.Ifindex
    leafs["cefAdjNextHopAddrType"] = cefadjentry.Cefadjnexthopaddrtype
    leafs["cefAdjNextHopAddr"] = cefadjentry.Cefadjnexthopaddr
    leafs["cefAdjConnId"] = cefadjentry.Cefadjconnid
    leafs["cefAdjSummaryLinkType"] = cefadjentry.Cefadjsummarylinktype
    leafs["cefAdjSource"] = cefadjentry.Cefadjsource
    leafs["cefAdjEncap"] = cefadjentry.Cefadjencap
    leafs["cefAdjFixup"] = cefadjentry.Cefadjfixup
    leafs["cefAdjMTU"] = cefadjentry.Cefadjmtu
    leafs["cefAdjForwardingInfo"] = cefadjentry.Cefadjforwardinginfo
    leafs["cefAdjPkts"] = cefadjentry.Cefadjpkts
    leafs["cefAdjHCPkts"] = cefadjentry.Cefadjhcpkts
    leafs["cefAdjBytes"] = cefadjentry.Cefadjbytes
    leafs["cefAdjHCBytes"] = cefadjentry.Cefadjhcbytes
    return leafs
}

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetYangName() string { return "cefAdjEntry" }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) SetParent(parent types.Entity) { cefadjentry.parent = parent }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetParent() types.Entity { return cefadjentry.parent }

func (cefadjentry *CISCOCEFMIB_Cefadjtable_Cefadjentry) GetParentYangName() string { return "cefAdjTable" }

// CISCOCEFMIB_Ceffeselectiontable
// A list of forwarding element selection entries.
type CISCOCEFMIB_Ceffeselectiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contain a CEF
    // forwarding element selection list.  entPhysicalIndex is also an index for
    // this table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry.
    Ceffeselectionentry []CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry
}

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetFilter() yfilter.YFilter { return ceffeselectiontable.YFilter }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) SetFilter(yf yfilter.YFilter) { ceffeselectiontable.YFilter = yf }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetGoName(yname string) string {
    if yname == "cefFESelectionEntry" { return "Ceffeselectionentry" }
    return ""
}

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetSegmentPath() string {
    return "cefFESelectionTable"
}

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefFESelectionEntry" {
        for _, c := range ceffeselectiontable.Ceffeselectionentry {
            if ceffeselectiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry{}
        ceffeselectiontable.Ceffeselectionentry = append(ceffeselectiontable.Ceffeselectionentry, child)
        return &ceffeselectiontable.Ceffeselectionentry[len(ceffeselectiontable.Ceffeselectionentry)-1]
    }
    return nil
}

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceffeselectiontable.Ceffeselectionentry {
        children[ceffeselectiontable.Ceffeselectionentry[i].GetSegmentPath()] = &ceffeselectiontable.Ceffeselectionentry[i]
    }
    return children
}

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetBundleName() string { return "cisco_ios_xe" }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetYangName() string { return "cefFESelectionTable" }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) SetParent(parent types.Entity) { ceffeselectiontable.parent = parent }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetParent() types.Entity { return ceffeselectiontable.parent }

func (ceffeselectiontable *CISCOCEFMIB_Ceffeselectiontable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetFilter() yfilter.YFilter { return ceffeselectionentry.YFilter }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) SetFilter(yf yfilter.YFilter) { ceffeselectionentry.YFilter = yf }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefFESelectionName" { return "Ceffeselectionname" }
    if yname == "cefFESelectionId" { return "Ceffeselectionid" }
    if yname == "cefFESelectionSpecial" { return "Ceffeselectionspecial" }
    if yname == "cefFESelectionLabels" { return "Ceffeselectionlabels" }
    if yname == "cefFESelectionAdjLinkType" { return "Ceffeselectionadjlinktype" }
    if yname == "cefFESelectionAdjInterface" { return "Ceffeselectionadjinterface" }
    if yname == "cefFESelectionAdjNextHopAddrType" { return "Ceffeselectionadjnexthopaddrtype" }
    if yname == "cefFESelectionAdjNextHopAddr" { return "Ceffeselectionadjnexthopaddr" }
    if yname == "cefFESelectionAdjConnId" { return "Ceffeselectionadjconnid" }
    if yname == "cefFESelectionVrfName" { return "Ceffeselectionvrfname" }
    if yname == "cefFESelectionWeight" { return "Ceffeselectionweight" }
    return ""
}

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetSegmentPath() string {
    return "cefFESelectionEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceffeselectionentry.Entphysicalindex) + "']" + "[cefFESelectionName='" + fmt.Sprintf("%v", ceffeselectionentry.Ceffeselectionname) + "']" + "[cefFESelectionId='" + fmt.Sprintf("%v", ceffeselectionentry.Ceffeselectionid) + "']"
}

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceffeselectionentry.Entphysicalindex
    leafs["cefFESelectionName"] = ceffeselectionentry.Ceffeselectionname
    leafs["cefFESelectionId"] = ceffeselectionentry.Ceffeselectionid
    leafs["cefFESelectionSpecial"] = ceffeselectionentry.Ceffeselectionspecial
    leafs["cefFESelectionLabels"] = ceffeselectionentry.Ceffeselectionlabels
    leafs["cefFESelectionAdjLinkType"] = ceffeselectionentry.Ceffeselectionadjlinktype
    leafs["cefFESelectionAdjInterface"] = ceffeselectionentry.Ceffeselectionadjinterface
    leafs["cefFESelectionAdjNextHopAddrType"] = ceffeselectionentry.Ceffeselectionadjnexthopaddrtype
    leafs["cefFESelectionAdjNextHopAddr"] = ceffeselectionentry.Ceffeselectionadjnexthopaddr
    leafs["cefFESelectionAdjConnId"] = ceffeselectionentry.Ceffeselectionadjconnid
    leafs["cefFESelectionVrfName"] = ceffeselectionentry.Ceffeselectionvrfname
    leafs["cefFESelectionWeight"] = ceffeselectionentry.Ceffeselectionweight
    return leafs
}

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetYangName() string { return "cefFESelectionEntry" }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) SetParent(parent types.Entity) { ceffeselectionentry.parent = parent }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetParent() types.Entity { return ceffeselectionentry.parent }

func (ceffeselectionentry *CISCOCEFMIB_Ceffeselectiontable_Ceffeselectionentry) GetParentYangName() string { return "cefFESelectionTable" }

// CISCOCEFMIB_Cefcfgtable
// This table contains global config parameter 
// of CEF on the Managed device.
type CISCOCEFMIB_Cefcfgtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If the Managed device supports CEF,  each entry contains the CEF config 
    // parameter for the managed entity. A row may exist for each IP version type
    // (v4 and v6) depending upon the IP version supported on the device. 
    // entPhysicalIndex is also an index for this table which represents entities
    // of 'module' entPhysicalClass which are capable of running CEF. The type is
    // slice of CISCOCEFMIB_Cefcfgtable_Cefcfgentry.
    Cefcfgentry []CISCOCEFMIB_Cefcfgtable_Cefcfgentry
}

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetFilter() yfilter.YFilter { return cefcfgtable.YFilter }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) SetFilter(yf yfilter.YFilter) { cefcfgtable.YFilter = yf }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetGoName(yname string) string {
    if yname == "cefCfgEntry" { return "Cefcfgentry" }
    return ""
}

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetSegmentPath() string {
    return "cefCfgTable"
}

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefCfgEntry" {
        for _, c := range cefcfgtable.Cefcfgentry {
            if cefcfgtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefcfgtable_Cefcfgentry{}
        cefcfgtable.Cefcfgentry = append(cefcfgtable.Cefcfgentry, child)
        return &cefcfgtable.Cefcfgentry[len(cefcfgtable.Cefcfgentry)-1]
    }
    return nil
}

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcfgtable.Cefcfgentry {
        children[cefcfgtable.Cefcfgentry[i].GetSegmentPath()] = &cefcfgtable.Cefcfgentry[i]
    }
    return children
}

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetYangName() string { return "cefCfgTable" }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) SetParent(parent types.Entity) { cefcfgtable.parent = parent }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetParent() types.Entity { return cefcfgtable.parent }

func (cefcfgtable *CISCOCEFMIB_Cefcfgtable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetFilter() yfilter.YFilter { return cefcfgentry.YFilter }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) SetFilter(yf yfilter.YFilter) { cefcfgentry.YFilter = yf }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefCfgAdminState" { return "Cefcfgadminstate" }
    if yname == "cefCfgOperState" { return "Cefcfgoperstate" }
    if yname == "cefCfgDistributionAdminState" { return "Cefcfgdistributionadminstate" }
    if yname == "cefCfgDistributionOperState" { return "Cefcfgdistributionoperstate" }
    if yname == "cefCfgAccountingMap" { return "Cefcfgaccountingmap" }
    if yname == "cefCfgLoadSharingAlgorithm" { return "Cefcfgloadsharingalgorithm" }
    if yname == "cefCfgLoadSharingID" { return "Cefcfgloadsharingid" }
    if yname == "cefCfgTrafficStatsLoadInterval" { return "Cefcfgtrafficstatsloadinterval" }
    if yname == "cefCfgTrafficStatsUpdateRate" { return "Cefcfgtrafficstatsupdaterate" }
    return ""
}

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetSegmentPath() string {
    return "cefCfgEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfgentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefcfgentry.Ceffibipversion) + "']"
}

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcfgentry.Entphysicalindex
    leafs["cefFIBIpVersion"] = cefcfgentry.Ceffibipversion
    leafs["cefCfgAdminState"] = cefcfgentry.Cefcfgadminstate
    leafs["cefCfgOperState"] = cefcfgentry.Cefcfgoperstate
    leafs["cefCfgDistributionAdminState"] = cefcfgentry.Cefcfgdistributionadminstate
    leafs["cefCfgDistributionOperState"] = cefcfgentry.Cefcfgdistributionoperstate
    leafs["cefCfgAccountingMap"] = cefcfgentry.Cefcfgaccountingmap
    leafs["cefCfgLoadSharingAlgorithm"] = cefcfgentry.Cefcfgloadsharingalgorithm
    leafs["cefCfgLoadSharingID"] = cefcfgentry.Cefcfgloadsharingid
    leafs["cefCfgTrafficStatsLoadInterval"] = cefcfgentry.Cefcfgtrafficstatsloadinterval
    leafs["cefCfgTrafficStatsUpdateRate"] = cefcfgentry.Cefcfgtrafficstatsupdaterate
    return leafs
}

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetYangName() string { return "cefCfgEntry" }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) SetParent(parent types.Entity) { cefcfgentry.parent = parent }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetParent() types.Entity { return cefcfgentry.parent }

func (cefcfgentry *CISCOCEFMIB_Cefcfgtable_Cefcfgentry) GetParentYangName() string { return "cefCfgTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // If the Managed device supports CEF, each entry contains the CEF Resource 
    // parameters for the managed entity.  entPhysicalIndex is also an index for
    // this table which represents entities of 'module' entPhysicalClass which are
    // capable of running CEF. The type is slice of
    // CISCOCEFMIB_Cefresourcetable_Cefresourceentry.
    Cefresourceentry []CISCOCEFMIB_Cefresourcetable_Cefresourceentry
}

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetFilter() yfilter.YFilter { return cefresourcetable.YFilter }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) SetFilter(yf yfilter.YFilter) { cefresourcetable.YFilter = yf }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetGoName(yname string) string {
    if yname == "cefResourceEntry" { return "Cefresourceentry" }
    return ""
}

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetSegmentPath() string {
    return "cefResourceTable"
}

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefResourceEntry" {
        for _, c := range cefresourcetable.Cefresourceentry {
            if cefresourcetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefresourcetable_Cefresourceentry{}
        cefresourcetable.Cefresourceentry = append(cefresourcetable.Cefresourceentry, child)
        return &cefresourcetable.Cefresourceentry[len(cefresourcetable.Cefresourceentry)-1]
    }
    return nil
}

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefresourcetable.Cefresourceentry {
        children[cefresourcetable.Cefresourceentry[i].GetSegmentPath()] = &cefresourcetable.Cefresourceentry[i]
    }
    return children
}

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetBundleName() string { return "cisco_ios_xe" }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetYangName() string { return "cefResourceTable" }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) SetParent(parent types.Entity) { cefresourcetable.parent = parent }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetParent() types.Entity { return cefresourcetable.parent }

func (cefresourcetable *CISCOCEFMIB_Cefresourcetable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetFilter() yfilter.YFilter { return cefresourceentry.YFilter }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) SetFilter(yf yfilter.YFilter) { cefresourceentry.YFilter = yf }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefResourceMemoryUsed" { return "Cefresourcememoryused" }
    if yname == "cefResourceFailureReason" { return "Cefresourcefailurereason" }
    return ""
}

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetSegmentPath() string {
    return "cefResourceEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefresourceentry.Entphysicalindex) + "']"
}

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefresourceentry.Entphysicalindex
    leafs["cefResourceMemoryUsed"] = cefresourceentry.Cefresourcememoryused
    leafs["cefResourceFailureReason"] = cefresourceentry.Cefresourcefailurereason
    return leafs
}

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetYangName() string { return "cefResourceEntry" }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) SetParent(parent types.Entity) { cefresourceentry.parent = parent }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetParent() types.Entity { return cefresourceentry.parent }

func (cefresourceentry *CISCOCEFMIB_Cefresourcetable_Cefresourceentry) GetParentYangName() string { return "cefResourceTable" }

// CISCOCEFMIB_Cefinttable
// This Table contains interface specific
// information of CEF on the Managed
// device.
type CISCOCEFMIB_Cefinttable struct {
    parent types.Entity
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

func (cefinttable *CISCOCEFMIB_Cefinttable) GetFilter() yfilter.YFilter { return cefinttable.YFilter }

func (cefinttable *CISCOCEFMIB_Cefinttable) SetFilter(yf yfilter.YFilter) { cefinttable.YFilter = yf }

func (cefinttable *CISCOCEFMIB_Cefinttable) GetGoName(yname string) string {
    if yname == "cefIntEntry" { return "Cefintentry" }
    return ""
}

func (cefinttable *CISCOCEFMIB_Cefinttable) GetSegmentPath() string {
    return "cefIntTable"
}

func (cefinttable *CISCOCEFMIB_Cefinttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefIntEntry" {
        for _, c := range cefinttable.Cefintentry {
            if cefinttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefinttable_Cefintentry{}
        cefinttable.Cefintentry = append(cefinttable.Cefintentry, child)
        return &cefinttable.Cefintentry[len(cefinttable.Cefintentry)-1]
    }
    return nil
}

func (cefinttable *CISCOCEFMIB_Cefinttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefinttable.Cefintentry {
        children[cefinttable.Cefintentry[i].GetSegmentPath()] = &cefinttable.Cefintentry[i]
    }
    return children
}

func (cefinttable *CISCOCEFMIB_Cefinttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefinttable *CISCOCEFMIB_Cefinttable) GetBundleName() string { return "cisco_ios_xe" }

func (cefinttable *CISCOCEFMIB_Cefinttable) GetYangName() string { return "cefIntTable" }

func (cefinttable *CISCOCEFMIB_Cefinttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefinttable *CISCOCEFMIB_Cefinttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefinttable *CISCOCEFMIB_Cefinttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefinttable *CISCOCEFMIB_Cefinttable) SetParent(parent types.Entity) { cefinttable.parent = parent }

func (cefinttable *CISCOCEFMIB_Cefinttable) GetParent() types.Entity { return cefinttable.parent }

func (cefinttable *CISCOCEFMIB_Cefinttable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetFilter() yfilter.YFilter { return cefintentry.YFilter }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) SetFilter(yf yfilter.YFilter) { cefintentry.YFilter = yf }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cefIntSwitchingState" { return "Cefintswitchingstate" }
    if yname == "cefIntLoadSharing" { return "Cefintloadsharing" }
    if yname == "cefIntNonrecursiveAccouting" { return "Cefintnonrecursiveaccouting" }
    return ""
}

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetSegmentPath() string {
    return "cefIntEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefintentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefintentry.Ceffibipversion) + "']" + "[ifIndex='" + fmt.Sprintf("%v", cefintentry.Ifindex) + "']"
}

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefintentry.Entphysicalindex
    leafs["cefFIBIpVersion"] = cefintentry.Ceffibipversion
    leafs["ifIndex"] = cefintentry.Ifindex
    leafs["cefIntSwitchingState"] = cefintentry.Cefintswitchingstate
    leafs["cefIntLoadSharing"] = cefintentry.Cefintloadsharing
    leafs["cefIntNonrecursiveAccouting"] = cefintentry.Cefintnonrecursiveaccouting
    return leafs
}

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetYangName() string { return "cefIntEntry" }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) SetParent(parent types.Entity) { cefintentry.parent = parent }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetParent() types.Entity { return cefintentry.parent }

func (cefintentry *CISCOCEFMIB_Cefinttable_Cefintentry) GetParentYangName() string { return "cefIntTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF
    // related attributes  associated with a CEF peer entity.  entPhysicalIndex
    // and entPeerPhysicalIndex are also indexes for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_Cefpeertable_Cefpeerentry.
    Cefpeerentry []CISCOCEFMIB_Cefpeertable_Cefpeerentry
}

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetFilter() yfilter.YFilter { return cefpeertable.YFilter }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) SetFilter(yf yfilter.YFilter) { cefpeertable.YFilter = yf }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetGoName(yname string) string {
    if yname == "cefPeerEntry" { return "Cefpeerentry" }
    return ""
}

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetSegmentPath() string {
    return "cefPeerTable"
}

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefPeerEntry" {
        for _, c := range cefpeertable.Cefpeerentry {
            if cefpeertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefpeertable_Cefpeerentry{}
        cefpeertable.Cefpeerentry = append(cefpeertable.Cefpeerentry, child)
        return &cefpeertable.Cefpeerentry[len(cefpeertable.Cefpeerentry)-1]
    }
    return nil
}

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefpeertable.Cefpeerentry {
        children[cefpeertable.Cefpeerentry[i].GetSegmentPath()] = &cefpeertable.Cefpeerentry[i]
    }
    return children
}

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetBundleName() string { return "cisco_ios_xe" }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetYangName() string { return "cefPeerTable" }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) SetParent(parent types.Entity) { cefpeertable.parent = parent }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetParent() types.Entity { return cefpeertable.parent }

func (cefpeertable *CISCOCEFMIB_Cefpeertable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetFilter() yfilter.YFilter { return cefpeerentry.YFilter }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) SetFilter(yf yfilter.YFilter) { cefpeerentry.YFilter = yf }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entPeerPhysicalIndex" { return "Entpeerphysicalindex" }
    if yname == "cefPeerOperState" { return "Cefpeeroperstate" }
    if yname == "cefPeerNumberOfResets" { return "Cefpeernumberofresets" }
    return ""
}

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetSegmentPath() string {
    return "cefPeerEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefpeerentry.Entphysicalindex) + "']" + "[entPeerPhysicalIndex='" + fmt.Sprintf("%v", cefpeerentry.Entpeerphysicalindex) + "']"
}

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefpeerentry.Entphysicalindex
    leafs["entPeerPhysicalIndex"] = cefpeerentry.Entpeerphysicalindex
    leafs["cefPeerOperState"] = cefpeerentry.Cefpeeroperstate
    leafs["cefPeerNumberOfResets"] = cefpeerentry.Cefpeernumberofresets
    return leafs
}

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetYangName() string { return "cefPeerEntry" }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) SetParent(parent types.Entity) { cefpeerentry.parent = parent }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetParent() types.Entity { return cefpeerentry.parent }

func (cefpeerentry *CISCOCEFMIB_Cefpeertable_Cefpeerentry) GetParentYangName() string { return "cefPeerTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry contains the CEF FIB
    // State  associated a CEF peer entity.  entPhysicalIndex and
    // entPeerPhysicalIndex are also indexes for this table which represents
    // entities of 'module' entPhysicalClass which are capable of running CEF. The
    // type is slice of CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry.
    Cefpeerfibentry []CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry
}

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetFilter() yfilter.YFilter { return cefpeerfibtable.YFilter }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) SetFilter(yf yfilter.YFilter) { cefpeerfibtable.YFilter = yf }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetGoName(yname string) string {
    if yname == "cefPeerFIBEntry" { return "Cefpeerfibentry" }
    return ""
}

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetSegmentPath() string {
    return "cefPeerFIBTable"
}

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefPeerFIBEntry" {
        for _, c := range cefpeerfibtable.Cefpeerfibentry {
            if cefpeerfibtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry{}
        cefpeerfibtable.Cefpeerfibentry = append(cefpeerfibtable.Cefpeerfibentry, child)
        return &cefpeerfibtable.Cefpeerfibentry[len(cefpeerfibtable.Cefpeerfibentry)-1]
    }
    return nil
}

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefpeerfibtable.Cefpeerfibentry {
        children[cefpeerfibtable.Cefpeerfibentry[i].GetSegmentPath()] = &cefpeerfibtable.Cefpeerfibentry[i]
    }
    return children
}

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetYangName() string { return "cefPeerFIBTable" }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) SetParent(parent types.Entity) { cefpeerfibtable.parent = parent }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetParent() types.Entity { return cefpeerfibtable.parent }

func (cefpeerfibtable *CISCOCEFMIB_Cefpeerfibtable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetFilter() yfilter.YFilter { return cefpeerfibentry.YFilter }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) SetFilter(yf yfilter.YFilter) { cefpeerfibentry.YFilter = yf }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entPeerPhysicalIndex" { return "Entpeerphysicalindex" }
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefPeerFIBOperState" { return "Cefpeerfiboperstate" }
    return ""
}

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetSegmentPath() string {
    return "cefPeerFIBEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefpeerfibentry.Entphysicalindex) + "']" + "[entPeerPhysicalIndex='" + fmt.Sprintf("%v", cefpeerfibentry.Entpeerphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefpeerfibentry.Ceffibipversion) + "']"
}

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefpeerfibentry.Entphysicalindex
    leafs["entPeerPhysicalIndex"] = cefpeerfibentry.Entpeerphysicalindex
    leafs["cefFIBIpVersion"] = cefpeerfibentry.Ceffibipversion
    leafs["cefPeerFIBOperState"] = cefpeerfibentry.Cefpeerfiboperstate
    return leafs
}

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetYangName() string { return "cefPeerFIBEntry" }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) SetParent(parent types.Entity) { cefpeerfibentry.parent = parent }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetParent() types.Entity { return cefpeerfibentry.parent }

func (cefpeerfibentry *CISCOCEFMIB_Cefpeerfibtable_Cefpeerfibentry) GetParentYangName() string { return "cefPeerFIBTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the global
    // consistency  checker parameter for the managed device. A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device. The type is slice of
    // CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry.
    Cefccglobalentry []CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry
}

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetFilter() yfilter.YFilter { return cefccglobaltable.YFilter }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) SetFilter(yf yfilter.YFilter) { cefccglobaltable.YFilter = yf }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetGoName(yname string) string {
    if yname == "cefCCGlobalEntry" { return "Cefccglobalentry" }
    return ""
}

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetSegmentPath() string {
    return "cefCCGlobalTable"
}

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefCCGlobalEntry" {
        for _, c := range cefccglobaltable.Cefccglobalentry {
            if cefccglobaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry{}
        cefccglobaltable.Cefccglobalentry = append(cefccglobaltable.Cefccglobalentry, child)
        return &cefccglobaltable.Cefccglobalentry[len(cefccglobaltable.Cefccglobalentry)-1]
    }
    return nil
}

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefccglobaltable.Cefccglobalentry {
        children[cefccglobaltable.Cefccglobalentry[i].GetSegmentPath()] = &cefccglobaltable.Cefccglobalentry[i]
    }
    return children
}

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetYangName() string { return "cefCCGlobalTable" }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) SetParent(parent types.Entity) { cefccglobaltable.parent = parent }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetParent() types.Entity { return cefccglobaltable.parent }

func (cefccglobaltable *CISCOCEFMIB_Cefccglobaltable) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry
// If the managed device supports CEF,
// each entry contains the global consistency 
// checker parameter for the managed device.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
type CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry struct {
    parent types.Entity
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

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetFilter() yfilter.YFilter { return cefccglobalentry.YFilter }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) SetFilter(yf yfilter.YFilter) { cefccglobalentry.YFilter = yf }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetGoName(yname string) string {
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefCCGlobalAutoRepairEnabled" { return "Cefccglobalautorepairenabled" }
    if yname == "cefCCGlobalAutoRepairDelay" { return "Cefccglobalautorepairdelay" }
    if yname == "cefCCGlobalAutoRepairHoldDown" { return "Cefccglobalautorepairholddown" }
    if yname == "cefCCGlobalErrorMsgEnabled" { return "Cefccglobalerrormsgenabled" }
    if yname == "cefCCGlobalFullScanAction" { return "Cefccglobalfullscanaction" }
    if yname == "cefCCGlobalFullScanStatus" { return "Cefccglobalfullscanstatus" }
    return ""
}

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetSegmentPath() string {
    return "cefCCGlobalEntry" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefccglobalentry.Ceffibipversion) + "']"
}

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cefFIBIpVersion"] = cefccglobalentry.Ceffibipversion
    leafs["cefCCGlobalAutoRepairEnabled"] = cefccglobalentry.Cefccglobalautorepairenabled
    leafs["cefCCGlobalAutoRepairDelay"] = cefccglobalentry.Cefccglobalautorepairdelay
    leafs["cefCCGlobalAutoRepairHoldDown"] = cefccglobalentry.Cefccglobalautorepairholddown
    leafs["cefCCGlobalErrorMsgEnabled"] = cefccglobalentry.Cefccglobalerrormsgenabled
    leafs["cefCCGlobalFullScanAction"] = cefccglobalentry.Cefccglobalfullscanaction
    leafs["cefCCGlobalFullScanStatus"] = cefccglobalentry.Cefccglobalfullscanstatus
    return leafs
}

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetYangName() string { return "cefCCGlobalEntry" }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) SetParent(parent types.Entity) { cefccglobalentry.parent = parent }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetParent() types.Entity { return cefccglobalentry.parent }

func (cefccglobalentry *CISCOCEFMIB_Cefccglobaltable_Cefccglobalentry) GetParentYangName() string { return "cefCCGlobalTable" }

// CISCOCEFMIB_Cefcctypetable
// This table contains CEF consistency
// checker types specific parameters on the managed device.
// 
// All detected inconsistency are signaled to the
// Management Station via cefInconsistencyDetection
// notification.
type CISCOCEFMIB_Cefcctypetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the consistency 
    // checker statistics for a consistency  checker type. A row may exist for
    // each IP version type (v4 and v6) depending upon the IP version supported on
    // the device. The type is slice of CISCOCEFMIB_Cefcctypetable_Cefcctypeentry.
    Cefcctypeentry []CISCOCEFMIB_Cefcctypetable_Cefcctypeentry
}

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetFilter() yfilter.YFilter { return cefcctypetable.YFilter }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) SetFilter(yf yfilter.YFilter) { cefcctypetable.YFilter = yf }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetGoName(yname string) string {
    if yname == "cefCCTypeEntry" { return "Cefcctypeentry" }
    return ""
}

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetSegmentPath() string {
    return "cefCCTypeTable"
}

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefCCTypeEntry" {
        for _, c := range cefcctypetable.Cefcctypeentry {
            if cefcctypetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefcctypetable_Cefcctypeentry{}
        cefcctypetable.Cefcctypeentry = append(cefcctypetable.Cefcctypeentry, child)
        return &cefcctypetable.Cefcctypeentry[len(cefcctypetable.Cefcctypeentry)-1]
    }
    return nil
}

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcctypetable.Cefcctypeentry {
        children[cefcctypetable.Cefcctypeentry[i].GetSegmentPath()] = &cefcctypetable.Cefcctypeentry[i]
    }
    return children
}

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetYangName() string { return "cefCCTypeTable" }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) SetParent(parent types.Entity) { cefcctypetable.parent = parent }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetParent() types.Entity { return cefcctypetable.parent }

func (cefcctypetable *CISCOCEFMIB_Cefcctypetable) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Cefcctypetable_Cefcctypeentry
// If the managed device supports CEF,
// each entry contains the consistency 
// checker statistics for a consistency 
// checker type.
// A row may exist for each IP version type
// (v4 and v6) depending upon the IP version
// supported on the device.
type CISCOCEFMIB_Cefcctypetable_Cefcctypeentry struct {
    parent types.Entity
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

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetFilter() yfilter.YFilter { return cefcctypeentry.YFilter }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) SetFilter(yf yfilter.YFilter) { cefcctypeentry.YFilter = yf }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetGoName(yname string) string {
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefCCType" { return "Cefcctype" }
    if yname == "cefCCEnabled" { return "Cefccenabled" }
    if yname == "cefCCCount" { return "Cefcccount" }
    if yname == "cefCCPeriod" { return "Cefccperiod" }
    if yname == "cefCCQueriesSent" { return "Cefccqueriessent" }
    if yname == "cefCCQueriesIgnored" { return "Cefccqueriesignored" }
    if yname == "cefCCQueriesChecked" { return "Cefccquerieschecked" }
    if yname == "cefCCQueriesIterated" { return "Cefccqueriesiterated" }
    return ""
}

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetSegmentPath() string {
    return "cefCCTypeEntry" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefcctypeentry.Ceffibipversion) + "']" + "[cefCCType='" + fmt.Sprintf("%v", cefcctypeentry.Cefcctype) + "']"
}

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cefFIBIpVersion"] = cefcctypeentry.Ceffibipversion
    leafs["cefCCType"] = cefcctypeentry.Cefcctype
    leafs["cefCCEnabled"] = cefcctypeentry.Cefccenabled
    leafs["cefCCCount"] = cefcctypeentry.Cefcccount
    leafs["cefCCPeriod"] = cefcctypeentry.Cefccperiod
    leafs["cefCCQueriesSent"] = cefcctypeentry.Cefccqueriessent
    leafs["cefCCQueriesIgnored"] = cefcctypeentry.Cefccqueriesignored
    leafs["cefCCQueriesChecked"] = cefcctypeentry.Cefccquerieschecked
    leafs["cefCCQueriesIterated"] = cefcctypeentry.Cefccqueriesiterated
    return leafs
}

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetYangName() string { return "cefCCTypeEntry" }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) SetParent(parent types.Entity) { cefcctypeentry.parent = parent }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetParent() types.Entity { return cefcctypeentry.parent }

func (cefcctypeentry *CISCOCEFMIB_Cefcctypetable_Cefcctypeentry) GetParentYangName() string { return "cefCCTypeTable" }

// CISCOCEFMIB_Cefinconsistencyrecordtable
// This table contains CEF inconsistency
// records.
type CISCOCEFMIB_Cefinconsistencyrecordtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If the managed device supports CEF, each entry contains the inconsistency 
    // record. The type is slice of
    // CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry.
    Cefinconsistencyrecordentry []CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry
}

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetFilter() yfilter.YFilter { return cefinconsistencyrecordtable.YFilter }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) SetFilter(yf yfilter.YFilter) { cefinconsistencyrecordtable.YFilter = yf }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetGoName(yname string) string {
    if yname == "cefInconsistencyRecordEntry" { return "Cefinconsistencyrecordentry" }
    return ""
}

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetSegmentPath() string {
    return "cefInconsistencyRecordTable"
}

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefInconsistencyRecordEntry" {
        for _, c := range cefinconsistencyrecordtable.Cefinconsistencyrecordentry {
            if cefinconsistencyrecordtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry{}
        cefinconsistencyrecordtable.Cefinconsistencyrecordentry = append(cefinconsistencyrecordtable.Cefinconsistencyrecordentry, child)
        return &cefinconsistencyrecordtable.Cefinconsistencyrecordentry[len(cefinconsistencyrecordtable.Cefinconsistencyrecordentry)-1]
    }
    return nil
}

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefinconsistencyrecordtable.Cefinconsistencyrecordentry {
        children[cefinconsistencyrecordtable.Cefinconsistencyrecordentry[i].GetSegmentPath()] = &cefinconsistencyrecordtable.Cefinconsistencyrecordentry[i]
    }
    return children
}

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetYangName() string { return "cefInconsistencyRecordTable" }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) SetParent(parent types.Entity) { cefinconsistencyrecordtable.parent = parent }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetParent() types.Entity { return cefinconsistencyrecordtable.parent }

func (cefinconsistencyrecordtable *CISCOCEFMIB_Cefinconsistencyrecordtable) GetParentYangName() string { return "CISCO-CEF-MIB" }

// CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry
// If the managed device supports CEF,
// each entry contains the inconsistency 
// record.
type CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry struct {
    parent types.Entity
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

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetFilter() yfilter.YFilter { return cefinconsistencyrecordentry.YFilter }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) SetFilter(yf yfilter.YFilter) { cefinconsistencyrecordentry.YFilter = yf }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetGoName(yname string) string {
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefInconsistencyRecId" { return "Cefinconsistencyrecid" }
    if yname == "cefInconsistencyPrefixType" { return "Cefinconsistencyprefixtype" }
    if yname == "cefInconsistencyPrefixAddr" { return "Cefinconsistencyprefixaddr" }
    if yname == "cefInconsistencyPrefixLen" { return "Cefinconsistencyprefixlen" }
    if yname == "cefInconsistencyVrfName" { return "Cefinconsistencyvrfname" }
    if yname == "cefInconsistencyCCType" { return "Cefinconsistencycctype" }
    if yname == "cefInconsistencyEntity" { return "Cefinconsistencyentity" }
    if yname == "cefInconsistencyReason" { return "Cefinconsistencyreason" }
    return ""
}

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetSegmentPath() string {
    return "cefInconsistencyRecordEntry" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefinconsistencyrecordentry.Ceffibipversion) + "']" + "[cefInconsistencyRecId='" + fmt.Sprintf("%v", cefinconsistencyrecordentry.Cefinconsistencyrecid) + "']"
}

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cefFIBIpVersion"] = cefinconsistencyrecordentry.Ceffibipversion
    leafs["cefInconsistencyRecId"] = cefinconsistencyrecordentry.Cefinconsistencyrecid
    leafs["cefInconsistencyPrefixType"] = cefinconsistencyrecordentry.Cefinconsistencyprefixtype
    leafs["cefInconsistencyPrefixAddr"] = cefinconsistencyrecordentry.Cefinconsistencyprefixaddr
    leafs["cefInconsistencyPrefixLen"] = cefinconsistencyrecordentry.Cefinconsistencyprefixlen
    leafs["cefInconsistencyVrfName"] = cefinconsistencyrecordentry.Cefinconsistencyvrfname
    leafs["cefInconsistencyCCType"] = cefinconsistencyrecordentry.Cefinconsistencycctype
    leafs["cefInconsistencyEntity"] = cefinconsistencyrecordentry.Cefinconsistencyentity
    leafs["cefInconsistencyReason"] = cefinconsistencyrecordentry.Cefinconsistencyreason
    return leafs
}

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetYangName() string { return "cefInconsistencyRecordEntry" }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) SetParent(parent types.Entity) { cefinconsistencyrecordentry.parent = parent }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetParent() types.Entity { return cefinconsistencyrecordentry.parent }

func (cefinconsistencyrecordentry *CISCOCEFMIB_Cefinconsistencyrecordtable_Cefinconsistencyrecordentry) GetParentYangName() string { return "cefInconsistencyRecordTable" }

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
    parent types.Entity
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

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetFilter() yfilter.YFilter { return cefstatsprefixlentable.YFilter }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) SetFilter(yf yfilter.YFilter) { cefstatsprefixlentable.YFilter = yf }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetGoName(yname string) string {
    if yname == "cefStatsPrefixLenEntry" { return "Cefstatsprefixlenentry" }
    return ""
}

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetSegmentPath() string {
    return "cefStatsPrefixLenTable"
}

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefStatsPrefixLenEntry" {
        for _, c := range cefstatsprefixlentable.Cefstatsprefixlenentry {
            if cefstatsprefixlentable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry{}
        cefstatsprefixlentable.Cefstatsprefixlenentry = append(cefstatsprefixlentable.Cefstatsprefixlenentry, child)
        return &cefstatsprefixlentable.Cefstatsprefixlenentry[len(cefstatsprefixlentable.Cefstatsprefixlenentry)-1]
    }
    return nil
}

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefstatsprefixlentable.Cefstatsprefixlenentry {
        children[cefstatsprefixlentable.Cefstatsprefixlenentry[i].GetSegmentPath()] = &cefstatsprefixlentable.Cefstatsprefixlenentry[i]
    }
    return children
}

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetBundleName() string { return "cisco_ios_xe" }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetYangName() string { return "cefStatsPrefixLenTable" }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) SetParent(parent types.Entity) { cefstatsprefixlentable.parent = parent }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetParent() types.Entity { return cefstatsprefixlentable.parent }

func (cefstatsprefixlentable *CISCOCEFMIB_Cefstatsprefixlentable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetFilter() yfilter.YFilter { return cefstatsprefixlenentry.YFilter }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) SetFilter(yf yfilter.YFilter) { cefstatsprefixlenentry.YFilter = yf }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefStatsPrefixLen" { return "Cefstatsprefixlen" }
    if yname == "cefStatsPrefixQueries" { return "Cefstatsprefixqueries" }
    if yname == "cefStatsPrefixHCQueries" { return "Cefstatsprefixhcqueries" }
    if yname == "cefStatsPrefixInserts" { return "Cefstatsprefixinserts" }
    if yname == "cefStatsPrefixHCInserts" { return "Cefstatsprefixhcinserts" }
    if yname == "cefStatsPrefixDeletes" { return "Cefstatsprefixdeletes" }
    if yname == "cefStatsPrefixHCDeletes" { return "Cefstatsprefixhcdeletes" }
    if yname == "cefStatsPrefixElements" { return "Cefstatsprefixelements" }
    if yname == "cefStatsPrefixHCElements" { return "Cefstatsprefixhcelements" }
    return ""
}

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetSegmentPath() string {
    return "cefStatsPrefixLenEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefstatsprefixlenentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefstatsprefixlenentry.Ceffibipversion) + "']" + "[cefStatsPrefixLen='" + fmt.Sprintf("%v", cefstatsprefixlenentry.Cefstatsprefixlen) + "']"
}

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefstatsprefixlenentry.Entphysicalindex
    leafs["cefFIBIpVersion"] = cefstatsprefixlenentry.Ceffibipversion
    leafs["cefStatsPrefixLen"] = cefstatsprefixlenentry.Cefstatsprefixlen
    leafs["cefStatsPrefixQueries"] = cefstatsprefixlenentry.Cefstatsprefixqueries
    leafs["cefStatsPrefixHCQueries"] = cefstatsprefixlenentry.Cefstatsprefixhcqueries
    leafs["cefStatsPrefixInserts"] = cefstatsprefixlenentry.Cefstatsprefixinserts
    leafs["cefStatsPrefixHCInserts"] = cefstatsprefixlenentry.Cefstatsprefixhcinserts
    leafs["cefStatsPrefixDeletes"] = cefstatsprefixlenentry.Cefstatsprefixdeletes
    leafs["cefStatsPrefixHCDeletes"] = cefstatsprefixlenentry.Cefstatsprefixhcdeletes
    leafs["cefStatsPrefixElements"] = cefstatsprefixlenentry.Cefstatsprefixelements
    leafs["cefStatsPrefixHCElements"] = cefstatsprefixlenentry.Cefstatsprefixhcelements
    return leafs
}

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetYangName() string { return "cefStatsPrefixLenEntry" }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) SetParent(parent types.Entity) { cefstatsprefixlenentry.parent = parent }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetParent() types.Entity { return cefstatsprefixlenentry.parent }

func (cefstatsprefixlenentry *CISCOCEFMIB_Cefstatsprefixlentable_Cefstatsprefixlenentry) GetParentYangName() string { return "cefStatsPrefixLenTable" }

// CISCOCEFMIB_Cefswitchingstatstable
// This table specifies the CEF switch stats.
type CISCOCEFMIB_Cefswitchingstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If CEF is enabled on the Managed device, each entry specifies the switching
    // stats. A row may exist for each IP version type (v4 and v6) depending upon
    // the IP version supported on the device.  entPhysicalIndex is also an index
    // for this table which represents entities of 'module' entPhysicalClass which
    // are capable of running CEF. The type is slice of
    // CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry.
    Cefswitchingstatsentry []CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry
}

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetFilter() yfilter.YFilter { return cefswitchingstatstable.YFilter }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) SetFilter(yf yfilter.YFilter) { cefswitchingstatstable.YFilter = yf }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetGoName(yname string) string {
    if yname == "cefSwitchingStatsEntry" { return "Cefswitchingstatsentry" }
    return ""
}

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetSegmentPath() string {
    return "cefSwitchingStatsTable"
}

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefSwitchingStatsEntry" {
        for _, c := range cefswitchingstatstable.Cefswitchingstatsentry {
            if cefswitchingstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry{}
        cefswitchingstatstable.Cefswitchingstatsentry = append(cefswitchingstatstable.Cefswitchingstatsentry, child)
        return &cefswitchingstatstable.Cefswitchingstatsentry[len(cefswitchingstatstable.Cefswitchingstatsentry)-1]
    }
    return nil
}

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefswitchingstatstable.Cefswitchingstatsentry {
        children[cefswitchingstatstable.Cefswitchingstatsentry[i].GetSegmentPath()] = &cefswitchingstatstable.Cefswitchingstatsentry[i]
    }
    return children
}

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetYangName() string { return "cefSwitchingStatsTable" }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) SetParent(parent types.Entity) { cefswitchingstatstable.parent = parent }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetParent() types.Entity { return cefswitchingstatstable.parent }

func (cefswitchingstatstable *CISCOCEFMIB_Cefswitchingstatstable) GetParentYangName() string { return "CISCO-CEF-MIB" }

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
    parent types.Entity
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

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetFilter() yfilter.YFilter { return cefswitchingstatsentry.YFilter }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) SetFilter(yf yfilter.YFilter) { cefswitchingstatsentry.YFilter = yf }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefFIBIpVersion" { return "Ceffibipversion" }
    if yname == "cefSwitchingIndex" { return "Cefswitchingindex" }
    if yname == "cefSwitchingPath" { return "Cefswitchingpath" }
    if yname == "cefSwitchingDrop" { return "Cefswitchingdrop" }
    if yname == "cefSwitchingHCDrop" { return "Cefswitchinghcdrop" }
    if yname == "cefSwitchingPunt" { return "Cefswitchingpunt" }
    if yname == "cefSwitchingHCPunt" { return "Cefswitchinghcpunt" }
    if yname == "cefSwitchingPunt2Host" { return "Cefswitchingpunt2Host" }
    if yname == "cefSwitchingHCPunt2Host" { return "Cefswitchinghcpunt2Host" }
    return ""
}

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetSegmentPath() string {
    return "cefSwitchingStatsEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefswitchingstatsentry.Entphysicalindex) + "']" + "[cefFIBIpVersion='" + fmt.Sprintf("%v", cefswitchingstatsentry.Ceffibipversion) + "']" + "[cefSwitchingIndex='" + fmt.Sprintf("%v", cefswitchingstatsentry.Cefswitchingindex) + "']"
}

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefswitchingstatsentry.Entphysicalindex
    leafs["cefFIBIpVersion"] = cefswitchingstatsentry.Ceffibipversion
    leafs["cefSwitchingIndex"] = cefswitchingstatsentry.Cefswitchingindex
    leafs["cefSwitchingPath"] = cefswitchingstatsentry.Cefswitchingpath
    leafs["cefSwitchingDrop"] = cefswitchingstatsentry.Cefswitchingdrop
    leafs["cefSwitchingHCDrop"] = cefswitchingstatsentry.Cefswitchinghcdrop
    leafs["cefSwitchingPunt"] = cefswitchingstatsentry.Cefswitchingpunt
    leafs["cefSwitchingHCPunt"] = cefswitchingstatsentry.Cefswitchinghcpunt
    leafs["cefSwitchingPunt2Host"] = cefswitchingstatsentry.Cefswitchingpunt2Host
    leafs["cefSwitchingHCPunt2Host"] = cefswitchingstatsentry.Cefswitchinghcpunt2Host
    return leafs
}

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetYangName() string { return "cefSwitchingStatsEntry" }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) SetParent(parent types.Entity) { cefswitchingstatsentry.parent = parent }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetParent() types.Entity { return cefswitchingstatsentry.parent }

func (cefswitchingstatsentry *CISCOCEFMIB_Cefswitchingstatstable_Cefswitchingstatsentry) GetParentYangName() string { return "cefSwitchingStatsTable" }

