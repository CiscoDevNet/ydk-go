// This MIB defines objects describing subscriber sessions, or
// more specifically, subscriber sessions terminated by a RAS.  A
// subscriber session consists of the traffic between a CPE and a
// NAS, as illustrated in the diagram below.
// 
//                                    Service
//                    Access          Provider
//                    Network         Network
// +--+  +---+  +--+   {   }   +---+   {   }
// |PC+--+CPE+--+AN+--{     }--+NAS+--{     }
// +--+  +---+  +--+   {   }   +---+   {   }
//           |                 |
//           |<--------------->|
//                subscriber
//                 session
// 
// A subscriber session behaves according to the FSM illustrated
// below.
// 
//       +-----------------+
//  +--->|  DISCONNECTED   |<-------+
//  |    +--------+--------+        |
//  |             |                 |
//  | failed      | initiated       | disconnect
//  |             V                 |
//  |    +-----------------+        |
//  +----+     PENDING     +--------+
//       +--------+--------+        |
//                |                 |
//                | established     |
//                V                 |
// +----------------------------+   |
// |             UP             |   |
// |                            +---+
// |  +-----------------+       |
// |  | UNAUTHENTICATED |       |
// |  +--------+--------+       |
// |           |                |
// |           | authenticated  |
// |           V                |
// |  +-----------------+       |
// |  |  AUTHENTICATED  |       |
// |  +-----------------+       |
// |                            |
// +----------------------------+
// 
// A subscriber session in the DISCONNECTED state technically
// doesn't exist; that is, the system does not maintain a context
// to describe a disconnected subscriber session.
// 
// Once the system detects the initiation of a subscriber session,
// then it creates a context and places the subscriber session in
// the PENDING state.  The initiation of a subscriber session can
// occur either through provisioning or the reception of a packet.
// In the PENDING state, a system does not forward subscriber
// traffic.
// 
// A pending subscriber session can become DISCONNECTED if
// it fails to come up (e.g., a timeout) or if the system or the
// subscriber explicitly terminates the subscriber session.
// 
// A pending subscriber session can become UP if the system
// successfully configures and applies any relevant policies.
// Once in the UP state, a system forwards subscriber traffic.
// 
// A operationally UP subscriber session can become DISCONNECTED if
// either system or the subscriber terminates it.
// 
// A operationally UP subscriber session can either be
// UNAUTHENTICATED or AUTHENTICATED.  When the system is in the
// process of checking a the credentials associated with a
// subscriber session, it is in the UNAUTHENTICATED state.  When
// the system successfully completes this process, it transitions
// the subscriber session to the AUTHENTICATED state.  If the
// process fails, then the system terminates the subscriber
// session.
// 
// Besides describing individual subscriber sessions, this MIB
// module provides an EMS/NMS with the means to perform the
// following functions:
// 
// 1)  Enumerate subscriber sessions by ifIndex.
// 
// 2)  Enumerate subscriber sessions by subscriber session type and
//     ifIndex.
// 
// 3)  Monitor aggregated statistics relating to subscriber
//     sessions:
// 
//     a.  System-wide
//     b.  System-wide by subscriber session type
//     c.  Per node
//     d.  Per node by subscriber session type
// 
// 4)  Collect 15-minute aggregated performance data history
//     relating to subscriber sessions:
// 
//     a.  System-wide
//     b.  System-wide by subscriber session type
//     c.  Per node
//     d.  Per node by subscriber session type
// 
// 5)  Submit a query for a report containing those subscriber
//     sessions that match a specified identity match criteria.
// 
// 6)  Terminate those subscriber session that match a
//     specified identify match criteria.
// 
// GLOSSARY
// ========
// 
// Access Concentrator
//     See NAS.
// 
// Access Network
//     The network that provides connectivity between a AN and NAS.
//     An access network may provide L2 or L3 connectivity. If the
//     access network provide L2 connectivity, it may switch
//     traffic or tunnel it through a MPLS or IP network.
// 
// AN (Access Node)
//     A device (e.g., a DSLAM) that multiplexes or switches
//     traffic between many CPEs and an access network.
// 
// BRAS (Broadband Remote Access Server)
//     See NAS.
// 
// CPE (Customer Premise Equipment)
//     A device (e.g., a DSL modem) that connects a customer's
//     network to an AN.
// 
// DHCP (Dynamic Host Configuration Protocol)
//     The protocol that provides a framework for transacting
//     configuration information to devices on an IP network, as
//     specified by RFC-2131.
// 
// NAS (Network Access Server)
//     A device that switches or routes traffic between subscriber
//     sessions and service provider networks.
// 
// Network Service
//     Access to the Internet backbone, voice, video, or other
//     content.
// 
// Node
//     A physical entity capable of maintaining a subscriber
//     session within a distributed system.  The notion of a node
//     is not applicable to a centralized system.
// 
// PADI (PPPoE Active Discovery Initiation)
//     A subscriber broadcasts a PADI packet to start the process
//     of discovering access concentrators capable of serving it.
// 
// PADO (PPPoE Active Discovery Offer)
//     The packet sent by an access concentrator to a subscriber
//     indicating that it can serve the subscriber.
// 
// PADR (PPPoE Active Discovery Request)
//     The packet sent by a subscriber to an access concentrator
//     requesting a PPPoE connection.
// 
// PADS (PPPoE Active Discovery Session-confirmation)
//     The packet sent by an access concentrator to a subscriber
//     confirming the request for a PPPoE connection.  Once this
//     packet has been sent, then the PPP can proceed as specified
//     by RFC-1661.
// 
// PADT (PPPoE Active Discovery Terminate)
//     The packet indicating that a PPPoE connection has been
//     terminated.  Either the subscriber or the access
//     concentrator can send this packet.
// 
// PPP (Point-to-Point Protocol)
//     The standard method for transporting multi-protocol
//     datagrams over point-to-point links, as defined by RFC-1661.
//     The PPP specifies the encapsulation for these datagrams and
//     the protocols necessary for establishing, configuring, and
//     maintaining connectivity.
// 
// PPPoE (Point-to-Point Protocol over Ethernet)
//     The protocol and encapsulation necessary to support a PPP
//     connection over an Ethernet connection, as defined by IETF
//     RFC-2516.
// 
// Service Provider Network
//     The network that provides connectivity between a NAS and a
//     network service.
// 
// Subscriber
//     A customer of a network service.
// 
// Subscriber Session
//     A context maintained by a NAS for the purpose of classifying
//     a subscriber's traffic, maintaining a subscriber's identity,
//     applying configuration and policies, and maintaining
//     statistics.  For more information on the types of subscriber
//     sessions, see the CISCO-SUBSCRIBER-SESSION-TC-MIB.
package cisco_subscriber_session_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_subscriber_session_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-SUBSCRIBER-SESSION-MIB CISCO-SUBSCRIBER-SESSION-MIB}", reflect.TypeOf(CISCOSUBSCRIBERSESSIONMIB{}))
    ydk.RegisterEntity("CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB", reflect.TypeOf(CISCOSUBSCRIBERSESSIONMIB{}))
}

// CISCOSUBSCRIBERSESSIONMIB
type CISCOSUBSCRIBERSESSIONMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Csubjob CISCOSUBSCRIBERSESSIONMIB_Csubjob

    
    Csubaggthresh CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh

    // This table describes a list of subscriber sessions currently maintained by
    // the system.  This table has a sparse dependent relationship on the ifTable,
    // containing a row for each interface having an interface type describing a
    // subscriber session.
    Csubsessiontable CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable

    // This table describes a list of subscriber sessions currently maintained by
    // the system.  The tables sorts the subscriber sessions first by the
    // subscriber session's type and second by the ifIndex assigned to the
    // subscriber session.
    Csubsessionbytypetable CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable

    // This table contains sets of aggregated statistics relating to subscriber
    // sessions, where each set has a unique scope of aggregation.
    Csubaggstatstable CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable

    // This table contains aggregated subscriber session performance data
    // collected for as much as a day's worth of 15-minute measurement intervals. 
    // This table has an expansion dependent relationship on the
    // csubAggStatsTable, containing zero or more rows for each row contained by
    // the csubAggStatsTable.  Observe that the collection and maintenance of
    // aggregated subscriber performance data is OPTIONAL for all scopes of
    // aggregation.  However, an implementation should maintain at least one
    // interval for the 'scope of aggregation' that contains all subscriber
    // sessions maintained by the system.
    Csubaggstatsinttable CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable

    // Please enter the Table Description here.
    Csubaggstatsthreshtable CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable

    // This table contains the subscriber session jobs submitted by the EMS/NMS.
    Csubjobtable CISCOSUBSCRIBERSESSIONMIB_Csubjobtable

    // This table contains subscriber session job parameters describing match
    // criteria.  This table has a sparse-dependent relationship on the
    // csubJobTable, containing a row for each job having a csubJobType of 'query'
    // or 'clear'.
    Csubjobmatchparamstable CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable

    // This table contains subscriber session job parameters describing query
    // parameters.  This table has a sparse-dependent relationship on the
    // csubJobTable, containing a row for each job having a csubJobType of
    // 'query'.
    Csubjobqueryparamstable CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable

    // This table lists the subscriber session jobs currently pending in the
    // subscriber session job queue.
    Csubjobqueuetable CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable

    // This table contains the reports corresponding to subscriber session jobs
    // that have a csubJobType of 'query' and csubJobState of 'finished'.  This
    // table has an expansion dependent relationship on the csubJobTable,
    // containing zero or more rows for each job.
    Csubjobreporttable CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable
}

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetFilter() yfilter.YFilter { return cISCOSUBSCRIBERSESSIONMIB.YFilter }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) SetFilter(yf yfilter.YFilter) { cISCOSUBSCRIBERSESSIONMIB.YFilter = yf }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetGoName(yname string) string {
    if yname == "csubJob" { return "Csubjob" }
    if yname == "csubAggThresh" { return "Csubaggthresh" }
    if yname == "csubSessionTable" { return "Csubsessiontable" }
    if yname == "csubSessionByTypeTable" { return "Csubsessionbytypetable" }
    if yname == "csubAggStatsTable" { return "Csubaggstatstable" }
    if yname == "csubAggStatsIntTable" { return "Csubaggstatsinttable" }
    if yname == "csubAggStatsThreshTable" { return "Csubaggstatsthreshtable" }
    if yname == "csubJobTable" { return "Csubjobtable" }
    if yname == "csubJobMatchParamsTable" { return "Csubjobmatchparamstable" }
    if yname == "csubJobQueryParamsTable" { return "Csubjobqueryparamstable" }
    if yname == "csubJobQueueTable" { return "Csubjobqueuetable" }
    if yname == "csubJobReportTable" { return "Csubjobreporttable" }
    return ""
}

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetSegmentPath() string {
    return "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB"
}

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubJob" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubjob
    }
    if childYangName == "csubAggThresh" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubaggthresh
    }
    if childYangName == "csubSessionTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubsessiontable
    }
    if childYangName == "csubSessionByTypeTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubsessionbytypetable
    }
    if childYangName == "csubAggStatsTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatstable
    }
    if childYangName == "csubAggStatsIntTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatsinttable
    }
    if childYangName == "csubAggStatsThreshTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatsthreshtable
    }
    if childYangName == "csubJobTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubjobtable
    }
    if childYangName == "csubJobMatchParamsTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubjobmatchparamstable
    }
    if childYangName == "csubJobQueryParamsTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubjobqueryparamstable
    }
    if childYangName == "csubJobQueueTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubjobqueuetable
    }
    if childYangName == "csubJobReportTable" {
        return &cISCOSUBSCRIBERSESSIONMIB.Csubjobreporttable
    }
    return nil
}

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["csubJob"] = &cISCOSUBSCRIBERSESSIONMIB.Csubjob
    children["csubAggThresh"] = &cISCOSUBSCRIBERSESSIONMIB.Csubaggthresh
    children["csubSessionTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubsessiontable
    children["csubSessionByTypeTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubsessionbytypetable
    children["csubAggStatsTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatstable
    children["csubAggStatsIntTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatsinttable
    children["csubAggStatsThreshTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatsthreshtable
    children["csubJobTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubjobtable
    children["csubJobMatchParamsTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubjobmatchparamstable
    children["csubJobQueryParamsTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubjobqueryparamstable
    children["csubJobQueueTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubjobqueuetable
    children["csubJobReportTable"] = &cISCOSUBSCRIBERSESSIONMIB.Csubjobreporttable
    return children
}

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) SetParent(parent types.Entity) { cISCOSUBSCRIBERSESSIONMIB.parent = parent }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetParent() types.Entity { return cISCOSUBSCRIBERSESSIONMIB.parent }

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjob
type CISCOSUBSCRIBERSESSIONMIB_Csubjob struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object specifies whether the system generates a csubJobFinishedNotify
    // notification when the system finishes processing a subscriber session job.
    // The type is bool.
    Csubjobfinishednotifyenable interface{}

    // This object indicates which subscriber session identities the system
    // maintains as searchable keys.  This value serves the EMS/NMS in configuring
    // a subscriber session query, as at least one match criteria must be an
    // 'indexed attribute'. The type is map[string]bool.
    Csubjobindexedattributes interface{}

    // This object indicates the next available identifier for the creation of a
    // new row in the csubJobTable.  If no available identifier exists, then this
    // object has the value '0'. The type is interface{} with range:
    // 0..4294967295.
    Csubjobidnext interface{}

    // This object indicates the maximum number of outstanding subscriber session
    // jobs the system can support.  If the csubJobTable contains a number of rows
    // (i.e., the value of csubJobCount) equal to this value, then any attempt to
    // create a new row will result in a response with an error-status of
    // 'resourceUnavailable'. The type is interface{} with range: 1..4294967295.
    // Units are jobs.
    Csubjobmaxnumber interface{}

    // This object specifies the maximum life a subscriber session report
    // corresponding to a subscriber session job having a csubJobType of 'query'. 
    // The system tracks the time elapsed after it finishes processing a query. 
    // When the time elapsed reaches the value specified by this column, the
    // system automatically  destroys the report.  A value of '0' disables the
    // automatic destruction of reports. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    Csubjobmaxlife interface{}

    // This object indicates the number of subscriber session jobs currently
    // maintained by the csubJobTable. The type is interface{} with range:
    // 0..4294967295. Units are jobs.
    Csubjobcount interface{}
}

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetFilter() yfilter.YFilter { return csubjob.YFilter }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) SetFilter(yf yfilter.YFilter) { csubjob.YFilter = yf }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetGoName(yname string) string {
    if yname == "csubJobFinishedNotifyEnable" { return "Csubjobfinishednotifyenable" }
    if yname == "csubJobIndexedAttributes" { return "Csubjobindexedattributes" }
    if yname == "csubJobIdNext" { return "Csubjobidnext" }
    if yname == "csubJobMaxNumber" { return "Csubjobmaxnumber" }
    if yname == "csubJobMaxLife" { return "Csubjobmaxlife" }
    if yname == "csubJobCount" { return "Csubjobcount" }
    return ""
}

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetSegmentPath() string {
    return "csubJob"
}

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubJobFinishedNotifyEnable"] = csubjob.Csubjobfinishednotifyenable
    leafs["csubJobIndexedAttributes"] = csubjob.Csubjobindexedattributes
    leafs["csubJobIdNext"] = csubjob.Csubjobidnext
    leafs["csubJobMaxNumber"] = csubjob.Csubjobmaxnumber
    leafs["csubJobMaxLife"] = csubjob.Csubjobmaxlife
    leafs["csubJobCount"] = csubjob.Csubjobcount
    return leafs
}

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetBundleName() string { return "cisco_ios_xe" }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetYangName() string { return "csubJob" }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) SetParent(parent types.Entity) { csubjob.parent = parent }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetParent() types.Entity { return csubjob.parent }

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh
type CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object enables or disables the generation of any of the csubAggStats*
    // threshold notifications. The type is bool.
    Csubaggstatsthreshnotifenable interface{}
}

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetFilter() yfilter.YFilter { return csubaggthresh.YFilter }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) SetFilter(yf yfilter.YFilter) { csubaggthresh.YFilter = yf }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetGoName(yname string) string {
    if yname == "csubAggStatsThreshNotifEnable" { return "Csubaggstatsthreshnotifenable" }
    return ""
}

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetSegmentPath() string {
    return "csubAggThresh"
}

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubAggStatsThreshNotifEnable"] = csubaggthresh.Csubaggstatsthreshnotifenable
    return leafs
}

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetBundleName() string { return "cisco_ios_xe" }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetYangName() string { return "csubAggThresh" }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) SetParent(parent types.Entity) { csubaggthresh.parent = parent }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetParent() types.Entity { return csubaggthresh.parent }

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable
// This table describes a list of subscriber sessions currently
// maintained by the system.
// 
// This table has a sparse dependent relationship on the ifTable,
// containing a row for each interface having an interface type
// describing a subscriber session.
type CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry contains data describing a subscriber sessions, including its
    // state, configuration, and collected identities.  An entry exists for a
    // corresponding entry in the ifTable describing a subscriber session. 
    // Currently, subscriber sessions must have one of the following ifType
    // values:      'ppp'     'ipSubscriberSession'     'l2SubscriberSession'  The
    // system creates an entry when it establishes a subscriber session. 
    // Likewise, the system destroys an entry when it terminates the corresponding
    // subscriber session. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry.
    Csubsessionentry []CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry
}

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetFilter() yfilter.YFilter { return csubsessiontable.YFilter }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) SetFilter(yf yfilter.YFilter) { csubsessiontable.YFilter = yf }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetGoName(yname string) string {
    if yname == "csubSessionEntry" { return "Csubsessionentry" }
    return ""
}

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetSegmentPath() string {
    return "csubSessionTable"
}

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubSessionEntry" {
        for _, c := range csubsessiontable.Csubsessionentry {
            if csubsessiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry{}
        csubsessiontable.Csubsessionentry = append(csubsessiontable.Csubsessionentry, child)
        return &csubsessiontable.Csubsessionentry[len(csubsessiontable.Csubsessionentry)-1]
    }
    return nil
}

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubsessiontable.Csubsessionentry {
        children[csubsessiontable.Csubsessionentry[i].GetSegmentPath()] = &csubsessiontable.Csubsessionentry[i]
    }
    return children
}

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetBundleName() string { return "cisco_ios_xe" }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetYangName() string { return "csubSessionTable" }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) SetParent(parent types.Entity) { csubsessiontable.parent = parent }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetParent() types.Entity { return csubsessiontable.parent }

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry
// This entry contains data describing a subscriber sessions,
// including its state, configuration, and collected identities.
// 
// An entry exists for a corresponding entry in the ifTable
// describing a subscriber session.  Currently, subscriber
// sessions must have one of the following ifType values:
// 
//     'ppp'
//     'ipSubscriberSession'
//     'l2SubscriberSession'
// 
// The system creates an entry when it establishes a subscriber
// session.  Likewise, the system destroys an entry when it
// terminates the corresponding subscriber session.
type CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object indicates the type of subscriber session. The type is
    // SubSessionType.
    Csubsessiontype interface{}

    // This object indicates how the system assigns IP addresses to the
    // subscriber:      'none'         The system does not an involvement in (or
    // is it aware         of) the assignment an subscriber IP addresses.  For    
    // example, a system does not have an involvement in the         assignment of
    // subscriber IP addresses for IP interface         subscriber sessions.     
    // 'other'         The system assigns subscriber IP addresses using a        
    // method not recognized by this MIB module.      'static'         Subscriber
    // IP addresses have been configured correctly         for the service domain.
    // The system does not have an         involvement in the assignment of the IP
    // address.      'localPool'         The system assigns subscriber IP
    // addresses from a         locally configured pool of IP addresses.     
    // 'dhcpv4'         The system assigns subscriber IP addresses are using the  
    // DHCPv4.      'dhcpv6'         The system assigns subscriber IP addresses
    // using the         DHCPv6.      'userProfileIpAddr'         The system
    // assigns subscriber IP addresses from a user         profile.     
    // 'userProfileIpSubnet'         The system assigns the subscriber an IP
    // subnet from a         user profile.      'userProfileNamedPool'         The
    // system assigns subscriber IP addresses from a         locally configured
    // named pool specified by a user         profile. The type is
    // Csubsessionipaddrassignment.
    Csubsessionipaddrassignment interface{}

    // This object indicates the operational state of the subscriber session. The
    // type is SubSessionState.
    Csubsessionstate interface{}

    // This object indicates whether the system has successfully authenticated the
    // subscriber session.      'false'         The subscriber session is
    // operationally up, but in the         UNAUTHENTICATED state.      'true'    
    // The subscriber session is operationally up, but in the        
    // AUTHENTICATED state. The type is bool.
    Csubsessionauthenticated interface{}

    // This object indicates the redundancy mode in which the subscriber session
    // is operating. The type is SubSessionRedundancyMode.
    Csubsessionredundancymode interface{}

    // This object indicates when the subscriber session was created (i.e., when
    // the subscriber session was initiated). The type is string.
    Csubsessioncreationtime interface{}

    // This object indicates the row in the cdtTemplateTable (defined by the
    // CISCO-DYNAMIC-TEMPLATE-MIB) describing the derived configuration for the
    // subscriber session.  Observe that the value of cdtTemplateType
    // corresponding to the referenced row must be 'derived'. The type is string
    // with length: 1..64.
    Csubsessionderivedcfg interface{}

    // This object indicates the subscriber identities that the system has
    // successfully collected for the subscriber session.  Each bit in this bit
    // string corresponds to a column in this table.  If the bit is '0', then the
    // value of the corresponding column is invalid.  If the bit is '1', then the
    // value of the corresponding column represents the value of the subscriber
    // identity collected by the system.  The following list specifies the
    // mappings between the bits and the columns:      'subscriberLabel' =>
    // csubSessionSubscriberLabel     'macAddress'      => csubSessionMacAddress  
    // 'nativeVrf'       => csubSessionNativeVrf     'nativeIpAddress' =>
    // csubSessionNativeIpAddrType,                         
    // csubSessionNativeIpAddr,                          csubSessionNativeIpMask  
    // 'nativeIpAddress2'=> csubSessionNativeIpAddrType2,                         
    // csubSessionNativeIpAddr2,                          csubSessionNativeIpMask2
    // 'domainVrf'       => csubSessionDomainVrf     'domainIpAddress' =>
    // csubSessionDomainIpAddrType,                         
    // csubSessionDomainIpAddr,                          csubSessionDomainIpMask  
    // 'pbhk'            => csubSessionPbhk     'remoteId'        =>
    // csubSessionRemoteId     'circuitId'       => csubSessionCircuitId    
    // 'nasPort'         => csubSessionNasPort     'domain'          =>
    // csubSessionDomain     'username'        => csubSessionUsername    
    // 'acctSessionId'   => csubSessionAcctSessionId     'dnis'            =>
    // csubSessionDnis     'media'           => csubSessionMedia    
    // 'mlpNegotiated'   => csubSessionMlpNegotiated     'protocol'        =>
    // csubSessionProtocol     'dhcpClass'       => csubSessionDhcpClass    
    // 'tunnelName'      => csubSessionTunnelName  Observe that the bit 'ifIndex'
    // should always be '1'.  This identity maps to the ifIndex assigned to the
    // subscriber session.  Observe that the bit 'serviceName' maps to one or more
    // instance of ccbptPolicyMap (defined by the CISCO-CBP-TARGET-MIB). The type
    // is map[string]bool.
    Csubsessionavailableidentities interface{}

    // This object indicates a positive integer-value uniquely identifying the
    // subscriber session within the scope of the system.  This column is valid
    // only if the 'subscriberLabel' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is interface{} with range:
    // 0..4294967295.
    Csubsessionsubscriberlabel interface{}

    // This object indicates the MAC address of the subscriber.  This column is
    // valid only if the 'macAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Csubsessionmacaddress interface{}

    // This object indicates the VRF originating the subscriber session.  This
    // column is valid only if the 'nativeVrf' bit of the corresponding instance
    // of csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessionnativevrf interface{}

    // This object indicates the type of IP address assigned to the subscriber on
    // the customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is InetAddressType.
    Csubsessionnativeipaddrtype interface{}

    // This object indicates the IP address assigned to the subscriber on the
    // customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    Csubsessionnativeipaddr interface{}

    // This object indicates the corresponding mask for the IP address assigned to
    // the subscriber on the customer-facing side of the system.  This column is
    // valid only if the 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    Csubsessionnativeipmask interface{}

    // This object indicates the VRF to which the system transfers the subscriber
    // session's traffic.  This column is valid only if the 'domainVrf' bit of the
    // corresponding instance of csubSessionAvailableIdentities is '1'. The type
    // is string.
    Csubsessiondomainvrf interface{}

    // This object indicates the type of IP address assigned to the subscriber on
    // the service-facing side of the system.  This column is valid only if the
    // 'domainIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is InetAddressType.
    Csubsessiondomainipaddrtype interface{}

    // This object indicates the IP address assigned to the subscriber on the
    // service-facing side of the system.  This column is valid only if the
    // 'domainIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    Csubsessiondomainipaddr interface{}

    // This object indicates the corresponding mask for the IP address assigned to
    // the subscriber on the service-facing side of the system.  This column is
    // valid only if the 'domainIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    Csubsessiondomainipmask interface{}

    // This object indicates the PBHK identifying the subscriber.  This column is
    // valid only if the 'pbhk' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessionpbhk interface{}

    // This object indicates the Remote-Id identifying the 'calling station' or AN
    // supporting the circuit that provides access to the subscriber.  This column
    // is valid only if the 'remoteId' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessionremoteid interface{}

    // This object indicates the Circuit-Id identifying the circuit supported by
    // the 'calling station' or AN providing access to the subscriber.  This
    // column is valid only if the 'circuitId' bit of the corresponding instance
    // of csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessioncircuitid interface{}

    // This object indicates the NAS port-identifier identifying the port on the
    // NAS providing access to the subscriber.  This column is valid only if the
    // 'nasPort' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessionnasport interface{}

    // This object indicates the domain associated with the subscriber.  This
    // column is valid only if the 'domain' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessiondomain interface{}

    // This object indicates the username identifying the subscriber.  This column
    // is valid only if the 'username' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessionusername interface{}

    // This object indicates the accounting session identifier assigned to the
    // subscriber session.  This column is valid only if the 'accountingSid' bit
    // of the corresponding instance of csubSessionAvailableIdentities is '1'. The
    // type is interface{} with range: 0..4294967295.
    Csubsessionacctsessionid interface{}

    // This object indicates the DNIS number dialed by the subscriber to access
    // the 'calling station' or AN.  This column is valid only if the 'dnis' bit
    // of the corresponding instance of csubSessionAvailableIdentities is '1'. The
    // type is string.
    Csubsessiondnis interface{}

    // This object indicates the type of media providing access to the subscriber.
    // This column is valid only if the 'media' bit of the corresponding instance
    // of csubSessionAvailableIdentities is '1'. The type is SubscriberMediaType.
    Csubsessionmedia interface{}

    // This object indicates whether the subscriber session was established using
    // multi-link PPP negotiation.  This column is valid only if the
    // 'mlpNegotiated' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is bool.
    Csubsessionmlpnegotiated interface{}

    // This object indicates the type of protocol providing access to the
    // subscriber.  This column is valid only if the 'protocol' bit of the
    // corresponding instance of csubSessionAvailableIdentities is '1'. The type
    // is SubscriberProtocolType.
    Csubsessionprotocol interface{}

    // This object indicates the name of the class matching the DHCP DISCOVER
    // message received from the subscriber.  This column is valid only if the
    // 'dhcpClass' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    Csubsessiondhcpclass interface{}

    // This object indicates the name of the VPDN used to carry the subscriber
    // session to the system.  This column is valid only if the 'tunnelName' bit
    // of the corresponding instance of csubSessionAvailableIdentities is '1'. The
    // type is string.
    Csubsessiontunnelname interface{}

    // This object indicates the location of the subscriber. The type is string.
    Csubsessionlocationidentifier interface{}

    // This object indicates the name used to identify the services subscribed by
    // a particular session. The type is string.
    Csubsessionserviceidentifier interface{}

    // This object indicates when the subscriber session is updated with new
    // policy information. The type is string.
    Csubsessionlastchanged interface{}

    // This object indicates the type of IP address assigned to the subscriber on
    // the customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'.  In Dual stack scenarios both
    // 'csubSessionNativeIpAddrType' and  'csubSessionNativeIpAddrType2' will be
    // valid. The type is InetAddressType.
    Csubsessionnativeipaddrtype2 interface{}

    // This object indicates the IP address assigned to the subscriber on the
    // customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    Csubsessionnativeipaddr2 interface{}

    // This object indicates the corresponding mask for the IP address assigned to
    // the subscriber on the customer-facing side of the system.  This column is
    // valid only if the 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    Csubsessionnativeipmask2 interface{}
}

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetFilter() yfilter.YFilter { return csubsessionentry.YFilter }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) SetFilter(yf yfilter.YFilter) { csubsessionentry.YFilter = yf }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "csubSessionType" { return "Csubsessiontype" }
    if yname == "csubSessionIpAddrAssignment" { return "Csubsessionipaddrassignment" }
    if yname == "csubSessionState" { return "Csubsessionstate" }
    if yname == "csubSessionAuthenticated" { return "Csubsessionauthenticated" }
    if yname == "csubSessionRedundancyMode" { return "Csubsessionredundancymode" }
    if yname == "csubSessionCreationTime" { return "Csubsessioncreationtime" }
    if yname == "csubSessionDerivedCfg" { return "Csubsessionderivedcfg" }
    if yname == "csubSessionAvailableIdentities" { return "Csubsessionavailableidentities" }
    if yname == "csubSessionSubscriberLabel" { return "Csubsessionsubscriberlabel" }
    if yname == "csubSessionMacAddress" { return "Csubsessionmacaddress" }
    if yname == "csubSessionNativeVrf" { return "Csubsessionnativevrf" }
    if yname == "csubSessionNativeIpAddrType" { return "Csubsessionnativeipaddrtype" }
    if yname == "csubSessionNativeIpAddr" { return "Csubsessionnativeipaddr" }
    if yname == "csubSessionNativeIpMask" { return "Csubsessionnativeipmask" }
    if yname == "csubSessionDomainVrf" { return "Csubsessiondomainvrf" }
    if yname == "csubSessionDomainIpAddrType" { return "Csubsessiondomainipaddrtype" }
    if yname == "csubSessionDomainIpAddr" { return "Csubsessiondomainipaddr" }
    if yname == "csubSessionDomainIpMask" { return "Csubsessiondomainipmask" }
    if yname == "csubSessionPbhk" { return "Csubsessionpbhk" }
    if yname == "csubSessionRemoteId" { return "Csubsessionremoteid" }
    if yname == "csubSessionCircuitId" { return "Csubsessioncircuitid" }
    if yname == "csubSessionNasPort" { return "Csubsessionnasport" }
    if yname == "csubSessionDomain" { return "Csubsessiondomain" }
    if yname == "csubSessionUsername" { return "Csubsessionusername" }
    if yname == "csubSessionAcctSessionId" { return "Csubsessionacctsessionid" }
    if yname == "csubSessionDnis" { return "Csubsessiondnis" }
    if yname == "csubSessionMedia" { return "Csubsessionmedia" }
    if yname == "csubSessionMlpNegotiated" { return "Csubsessionmlpnegotiated" }
    if yname == "csubSessionProtocol" { return "Csubsessionprotocol" }
    if yname == "csubSessionDhcpClass" { return "Csubsessiondhcpclass" }
    if yname == "csubSessionTunnelName" { return "Csubsessiontunnelname" }
    if yname == "csubSessionLocationIdentifier" { return "Csubsessionlocationidentifier" }
    if yname == "csubSessionServiceIdentifier" { return "Csubsessionserviceidentifier" }
    if yname == "csubSessionLastChanged" { return "Csubsessionlastchanged" }
    if yname == "csubSessionNativeIpAddrType2" { return "Csubsessionnativeipaddrtype2" }
    if yname == "csubSessionNativeIpAddr2" { return "Csubsessionnativeipaddr2" }
    if yname == "csubSessionNativeIpMask2" { return "Csubsessionnativeipmask2" }
    return ""
}

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetSegmentPath() string {
    return "csubSessionEntry" + "[ifIndex='" + fmt.Sprintf("%v", csubsessionentry.Ifindex) + "']"
}

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = csubsessionentry.Ifindex
    leafs["csubSessionType"] = csubsessionentry.Csubsessiontype
    leafs["csubSessionIpAddrAssignment"] = csubsessionentry.Csubsessionipaddrassignment
    leafs["csubSessionState"] = csubsessionentry.Csubsessionstate
    leafs["csubSessionAuthenticated"] = csubsessionentry.Csubsessionauthenticated
    leafs["csubSessionRedundancyMode"] = csubsessionentry.Csubsessionredundancymode
    leafs["csubSessionCreationTime"] = csubsessionentry.Csubsessioncreationtime
    leafs["csubSessionDerivedCfg"] = csubsessionentry.Csubsessionderivedcfg
    leafs["csubSessionAvailableIdentities"] = csubsessionentry.Csubsessionavailableidentities
    leafs["csubSessionSubscriberLabel"] = csubsessionentry.Csubsessionsubscriberlabel
    leafs["csubSessionMacAddress"] = csubsessionentry.Csubsessionmacaddress
    leafs["csubSessionNativeVrf"] = csubsessionentry.Csubsessionnativevrf
    leafs["csubSessionNativeIpAddrType"] = csubsessionentry.Csubsessionnativeipaddrtype
    leafs["csubSessionNativeIpAddr"] = csubsessionentry.Csubsessionnativeipaddr
    leafs["csubSessionNativeIpMask"] = csubsessionentry.Csubsessionnativeipmask
    leafs["csubSessionDomainVrf"] = csubsessionentry.Csubsessiondomainvrf
    leafs["csubSessionDomainIpAddrType"] = csubsessionentry.Csubsessiondomainipaddrtype
    leafs["csubSessionDomainIpAddr"] = csubsessionentry.Csubsessiondomainipaddr
    leafs["csubSessionDomainIpMask"] = csubsessionentry.Csubsessiondomainipmask
    leafs["csubSessionPbhk"] = csubsessionentry.Csubsessionpbhk
    leafs["csubSessionRemoteId"] = csubsessionentry.Csubsessionremoteid
    leafs["csubSessionCircuitId"] = csubsessionentry.Csubsessioncircuitid
    leafs["csubSessionNasPort"] = csubsessionentry.Csubsessionnasport
    leafs["csubSessionDomain"] = csubsessionentry.Csubsessiondomain
    leafs["csubSessionUsername"] = csubsessionentry.Csubsessionusername
    leafs["csubSessionAcctSessionId"] = csubsessionentry.Csubsessionacctsessionid
    leafs["csubSessionDnis"] = csubsessionentry.Csubsessiondnis
    leafs["csubSessionMedia"] = csubsessionentry.Csubsessionmedia
    leafs["csubSessionMlpNegotiated"] = csubsessionentry.Csubsessionmlpnegotiated
    leafs["csubSessionProtocol"] = csubsessionentry.Csubsessionprotocol
    leafs["csubSessionDhcpClass"] = csubsessionentry.Csubsessiondhcpclass
    leafs["csubSessionTunnelName"] = csubsessionentry.Csubsessiontunnelname
    leafs["csubSessionLocationIdentifier"] = csubsessionentry.Csubsessionlocationidentifier
    leafs["csubSessionServiceIdentifier"] = csubsessionentry.Csubsessionserviceidentifier
    leafs["csubSessionLastChanged"] = csubsessionentry.Csubsessionlastchanged
    leafs["csubSessionNativeIpAddrType2"] = csubsessionentry.Csubsessionnativeipaddrtype2
    leafs["csubSessionNativeIpAddr2"] = csubsessionentry.Csubsessionnativeipaddr2
    leafs["csubSessionNativeIpMask2"] = csubsessionentry.Csubsessionnativeipmask2
    return leafs
}

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetYangName() string { return "csubSessionEntry" }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) SetParent(parent types.Entity) { csubsessionentry.parent = parent }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetParent() types.Entity { return csubsessionentry.parent }

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetParentYangName() string { return "csubSessionTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment represents         profile.
type CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment string

const (
    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_none CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "none"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_other CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "other"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_static CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "static"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_localPool CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "localPool"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_dhcpv4 CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "dhcpv4"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_dhcpv6 CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "dhcpv6"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_userProfileIpAddr CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "userProfileIpAddr"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_userProfileIpSubnet CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "userProfileIpSubnet"

    CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment_userProfileNamedPool CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry_Csubsessionipaddrassignment = "userProfileNamedPool"
)

// CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable
// This table describes a list of subscriber sessions currently
// maintained by the system.  The tables sorts the subscriber
// sessions first by the subscriber session's type and second by
// the ifIndex assigned to the subscriber session.
type CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry identifies a subscriber session.  An entry exists for a
    // corresponding entry in the ifTable describing a subscriber session. 
    // Currently, subscriber sessions must have one of the following ifType
    // values:      'ppp'     'ipSubscriberSession'     'l2SubscriberSession'  The
    // system creates an entry when it establishes a subscriber session. 
    // Likewise, the system destroys an entry when it terminates the corresponding
    // subscriber session. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry.
    Csubsessionbytypeentry []CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry
}

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetFilter() yfilter.YFilter { return csubsessionbytypetable.YFilter }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) SetFilter(yf yfilter.YFilter) { csubsessionbytypetable.YFilter = yf }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetGoName(yname string) string {
    if yname == "csubSessionByTypeEntry" { return "Csubsessionbytypeentry" }
    return ""
}

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetSegmentPath() string {
    return "csubSessionByTypeTable"
}

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubSessionByTypeEntry" {
        for _, c := range csubsessionbytypetable.Csubsessionbytypeentry {
            if csubsessionbytypetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry{}
        csubsessionbytypetable.Csubsessionbytypeentry = append(csubsessionbytypetable.Csubsessionbytypeentry, child)
        return &csubsessionbytypetable.Csubsessionbytypeentry[len(csubsessionbytypetable.Csubsessionbytypeentry)-1]
    }
    return nil
}

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubsessionbytypetable.Csubsessionbytypeentry {
        children[csubsessionbytypetable.Csubsessionbytypeentry[i].GetSegmentPath()] = &csubsessionbytypetable.Csubsessionbytypeentry[i]
    }
    return children
}

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetBundleName() string { return "cisco_ios_xe" }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetYangName() string { return "csubSessionByTypeTable" }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) SetParent(parent types.Entity) { csubsessionbytypetable.parent = parent }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetParent() types.Entity { return csubsessionbytypetable.parent }

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry
// This entry identifies a subscriber session.
// 
// An entry exists for a corresponding entry in the ifTable
// describing a subscriber session.  Currently, subscriber
// sessions must have one of the following ifType values:
// 
//     'ppp'
//     'ipSubscriberSession'
//     'l2SubscriberSession'
// 
// The system creates an entry when it establishes a subscriber
// session.  Likewise, the system destroys an entry when it
// terminates the corresponding subscriber session.
type CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates the type of the subscriber
    // session. The type is SubSessionType.
    Csubsessionbytype interface{}

    // This attribute is a key. This object indicates the ifIndex assigned to the
    // subscriber session. The type is interface{} with range: 1..2147483647.
    Csubsessionifindex interface{}
}

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetFilter() yfilter.YFilter { return csubsessionbytypeentry.YFilter }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) SetFilter(yf yfilter.YFilter) { csubsessionbytypeentry.YFilter = yf }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetGoName(yname string) string {
    if yname == "csubSessionByType" { return "Csubsessionbytype" }
    if yname == "csubSessionIfIndex" { return "Csubsessionifindex" }
    return ""
}

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetSegmentPath() string {
    return "csubSessionByTypeEntry" + "[csubSessionByType='" + fmt.Sprintf("%v", csubsessionbytypeentry.Csubsessionbytype) + "']" + "[csubSessionIfIndex='" + fmt.Sprintf("%v", csubsessionbytypeentry.Csubsessionifindex) + "']"
}

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubSessionByType"] = csubsessionbytypeentry.Csubsessionbytype
    leafs["csubSessionIfIndex"] = csubsessionbytypeentry.Csubsessionifindex
    return leafs
}

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetYangName() string { return "csubSessionByTypeEntry" }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) SetParent(parent types.Entity) { csubsessionbytypeentry.parent = parent }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetParent() types.Entity { return csubsessionbytypeentry.parent }

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetParentYangName() string { return "csubSessionByTypeTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable
// This table contains sets of aggregated statistics relating to
// subscriber sessions, where each set has a unique scope of
// aggregation.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry contains a set of aggregated statistics relating to those
    // subscriber sessions that fall into a 'scope of  aggregation'.   A 'scope of
    // aggregation' is the set of subscriber sessions  that meet specified
    // criteria.  For example, a 'scope of  aggregation' may be the set of all
    // PPPoE subscriber sessions  maintained by the system.  The following
    // criteria define the  'scope of aggregation':   1)  Aggregation Point type  
    // Aggregation point type identifies the format of the         
    // csubAggStatsPoint for this entry.   2)  Aggregation Point         
    // 'Physical' Aggregation Point type case:          In a distributed system, a
    // 'node' represents a physical          entity capable of maintaining the
    // context representing          a subscriber session.           If the 'scope
    // of aggregation' specifies a physical          entity having an
    // entPhysicalClass of 'chassis', then          the set of subscriber sessions
    // in the 'scope of          aggregation' may contain the subscriber sessions
    // maintained by all          the nodes contained in the system.           If
    // the 'scope of aggregation' specifies a physical          entity having an
    // entPhysicalClass of 'module' (e.g., a          line card), then the set of
    // subscriber sessions in the          'scope of aggregation' may contain the
    // subscriber          sessions maintained by the nodes contained by the      
    // module.           If the 'scope of aggregation' specifies a physical       
    // entity having an entPhysicalClass of 'cpu', then the          set of
    // subscriber sessions in the 'scope of aggregation'          may contain the
    // subscriber sessions maintained by the node          running on that
    // processor.           Observe that a centralized system (i.e., a system     
    // that essentially contains a single node) can only          support a 'scope
    // of aggregation' that specifies a          physical entity classified as a
    // 'chassis'.           If the scope of aggregation specifies 'interface',    
    // then the scope is the set of subscriber sessions carried          by the
    // interface identified the ifIndex value          represented in the
    // csubAggStatsPoint value.   2)  Subscriber Session Type          If the
    // 'scope of aggregation' specifies the value 'all'          for the
    // subscriber session type, then the set of          subscriber sessions in
    // the 'scope of aggregation' may          contain all subscriber sessions,
    // regardless of type.           If the 'scope of aggregation' specifies a
    // value other          than 'all' for the subscriber session type, then the  
    // set of subscriber sessions in the 'scope of aggregation may         
    // contain only those subscriber sessions of the specified          type.  
    // Implementation Guidance  =======================  A system MUST maintain a
    // set of statistics with a 'scope of  aggregation' that contains all
    // subscriber sessions maintained  by the system.  The system creates this
    // entry during the  initialization of the SNMP entity.   A system SHOULD
    // maintain a set of statistics for each 'scope of  aggregation' containing
    // subscriber sessions of each subscriber  session type the system is capable
    // of providing access.  If the  system supports these sets of statistics,
    // then it creates these  entries during the initialization of the SNMP
    // entity.   A system MAY maintain sets of node-specific statistics.  if the 
    // system supports sets of node-specific statistics, then it  creates the
    // appropriate entries upon detection of a physical  entity (resulting from
    // system restart or insertion) containing  those nodes.  Likewise, the system
    // destroys these entries  upon removal of the physical entity. The type is
    // slice of CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry.
    Csubaggstatsentry []CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry
}

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetFilter() yfilter.YFilter { return csubaggstatstable.YFilter }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) SetFilter(yf yfilter.YFilter) { csubaggstatstable.YFilter = yf }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetGoName(yname string) string {
    if yname == "csubAggStatsEntry" { return "Csubaggstatsentry" }
    return ""
}

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetSegmentPath() string {
    return "csubAggStatsTable"
}

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubAggStatsEntry" {
        for _, c := range csubaggstatstable.Csubaggstatsentry {
            if csubaggstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry{}
        csubaggstatstable.Csubaggstatsentry = append(csubaggstatstable.Csubaggstatsentry, child)
        return &csubaggstatstable.Csubaggstatsentry[len(csubaggstatstable.Csubaggstatsentry)-1]
    }
    return nil
}

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubaggstatstable.Csubaggstatsentry {
        children[csubaggstatstable.Csubaggstatsentry[i].GetSegmentPath()] = &csubaggstatstable.Csubaggstatsentry[i]
    }
    return children
}

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetYangName() string { return "csubAggStatsTable" }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) SetParent(parent types.Entity) { csubaggstatstable.parent = parent }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetParent() types.Entity { return csubaggstatstable.parent }

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry
// An entry contains a set of aggregated statistics relating to
// those subscriber sessions that fall into a 'scope of 
// aggregation'. 
// 
// A 'scope of aggregation' is the set of subscriber sessions 
// that meet specified criteria.  For example, a 'scope of 
// aggregation' may be the set of all PPPoE subscriber sessions 
// maintained by the system.  The following criteria define the 
// 'scope of aggregation': 
// 
// 1)  Aggregation Point type 
//         Aggregation point type identifies the format of the 
//         csubAggStatsPoint for this entry. 
// 
// 2)  Aggregation Point 
//         'Physical' Aggregation Point type case: 
//         In a distributed system, a 'node' represents a physical 
//         entity capable of maintaining the context representing 
//         a subscriber session. 
// 
//         If the 'scope of aggregation' specifies a physical 
//         entity having an entPhysicalClass of 'chassis', then 
//         the set of subscriber sessions in the 'scope of 
//         aggregation' may contain the subscriber sessions maintained by all 
//         the nodes contained in the system. 
// 
//         If the 'scope of aggregation' specifies a physical 
//         entity having an entPhysicalClass of 'module' (e.g., a 
//         line card), then the set of subscriber sessions in the 
//         'scope of aggregation' may contain the subscriber 
//         sessions maintained by the nodes contained by the 
//         module. 
// 
//         If the 'scope of aggregation' specifies a physical 
//         entity having an entPhysicalClass of 'cpu', then the 
//         set of subscriber sessions in the 'scope of aggregation' 
//         may contain the subscriber sessions maintained by the node 
//         running on that processor. 
// 
//         Observe that a centralized system (i.e., a system 
//         that essentially contains a single node) can only 
//         support a 'scope of aggregation' that specifies a 
//         physical entity classified as a 'chassis'. 
// 
//         If the scope of aggregation specifies 'interface', 
//         then the scope is the set of subscriber sessions carried 
//         by the interface identified the ifIndex value 
//         represented in the csubAggStatsPoint value. 
// 
// 2)  Subscriber Session Type 
//         If the 'scope of aggregation' specifies the value 'all' 
//         for the subscriber session type, then the set of 
//         subscriber sessions in the 'scope of aggregation' may 
//         contain all subscriber sessions, regardless of type. 
// 
//         If the 'scope of aggregation' specifies a value other 
//         than 'all' for the subscriber session type, then the 
//         set of subscriber sessions in the 'scope of aggregation may 
//         contain only those subscriber sessions of the specified 
//         type. 
// 
// Implementation Guidance 
// ======================= 
// A system MUST maintain a set of statistics with a 'scope of 
// aggregation' that contains all subscriber sessions maintained 
// by the system.  The system creates this entry during the 
// initialization of the SNMP entity. 
// 
// A system SHOULD maintain a set of statistics for each 'scope of 
// aggregation' containing subscriber sessions of each subscriber 
// session type the system is capable of providing access.  If the 
// system supports these sets of statistics, then it creates these 
// entries during the initialization of the SNMP entity. 
// 
// A system MAY maintain sets of node-specific statistics.  if the 
// system supports sets of node-specific statistics, then it 
// creates the appropriate entries upon detection of a physical 
// entity (resulting from system restart or insertion) containing 
// those nodes.  Likewise, the system destroys these entries 
// upon removal of the physical entity.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates format of the
    // csubAggStatsPoint for this entry.   The format for the csubAggStatsPoint is
    // as follows:   csubAggStatsPointType      csubAggStatsPoint 
    // ----------------------     ------------------      'physical'              
    // PhysicalIndex      'interface'                InterfaceIndex. The type is
    // Csubaggstatspointtype.
    Csubaggstatspointtype interface{}

    // This attribute is a key. This object should be read with
    // csubAggStatsPointType always.  This object indicates one of the determining
    // factors affecting  the 'scope of aggregation' for the set of statistics
    // contained  by the row.   The value indicated by this object should be
    // interpreted as the  identifier for the point type specific base table.  
    // For point types of 'physical', the type specific base table is  the
    // entPhysicalTable and this value is a PhysicalIndex.  For  point types of
    // 'interface', the type specific base table is  the ifTable and this value is
    // an InterfaceIndex.   If this column indicates a physical entity which has
    // an  entPhysicalClass of 'chassis', then the 'scope of aggregation'  may
    // includes those subscriber sessions maintained by all nodes  contained by
    // the system.   If this column indicates a physical entity which has an 
    // entPhysicalClass of  'module' (e.g., a line card), then the  'scope of
    // aggregation' may include those subscriber sessions  maintained by the nodes
    // contained by the module.   If this column indicates a physical entity which
    // has an  entPhysicalClass of 'cpu', then the 'scope of aggregation' may 
    // include those subscriber sessions maintained by the node  running on the
    // processor.   Aggregation points with entPhysicalTable / ifTable overlap: 
    // For interfaces which map directly to physical 'port' class  entities in the
    // entPhysicalTable, the preferred representation  as aggregation points is
    // the 'physical' point type and  PhysicalIndex identifier. The type is
    // interface{} with range: 1..4294967295.
    Csubaggstatspoint interface{}

    // This attribute is a key. This object indicates one of the determining
    // factors affecting the 'scope of aggregation' for the statistics contained
    // by the row.  If the value of this column is 'all', then the 'scope of
    // aggregation' may include all subscriber sessions, regardless of type.  If
    // the value of this column is not 'all', then the 'scope of aggregation' may
    // include subscriber sessions of the indicated subscriber session type. The
    // type is SubSessionType.
    Csubaggstatssessiontype interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are in the PENDING state. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatspendingsessions interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are in the UP state. The type is interface{}
    // with range: 0..4294967295. Units are sessions.
    Csubaggstatsupsessions interface{}

    // This object indicates the current number of subscriber session within the
    // 'scope of aggregation' that have been authenticated. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatsauthsessions interface{}

    // This object indicates the current number of subscriber session within the
    // 'scope of aggregation' that have not been authenticated. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatsunauthsessions interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are less resource intensive. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatslightweightsessions interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are redundant (i.e.,  sessions with a
    // csubSessionRedundancyMode of 'standby'). The type is interface{} with
    // range: 0..4294967295. Units are sessions.
    Csubaggstatsredsessions interface{}

    // This object indicates the highest number of subscriber sessions within the
    // 'scope of aggregation' observed simultaneously in the UP state since the
    // last discontinuity time. The type is interface{} with range: 0..4294967295.
    // Units are sessions.
    Csubaggstatshighupsessions interface{}

    // This object indicates the average time subscriber sessions within the
    // 'scope of aggregation' spent in the UP state.  The system calculates this
    // average over all subscriber sessions maintained by all nodes contained by
    // the 'scope of aggregation' since the last discontinuity time. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Csubaggstatsavgsessionuptime interface{}

    // This object indicates the average rate (per minute) at which the nodes
    // contained by the 'scope of aggregation' have established new subscriber
    // sessions. The type is interface{} with range: 0..4294967295. Units are
    // sessions per minute.
    Csubaggstatsavgsessionrpm interface{}

    // This object indicates the average rate (per hour) at which the nodes
    // contained by the 'scope of aggregation' have established new subscriber
    // sessions. The type is interface{} with range: 0..4294967295. Units are
    // sessions per hour.
    Csubaggstatsavgsessionrph interface{}

    // This object indicates the number of times that nodes contained within the
    // 'scope of aggregation' have engaged the subscriber session throttle since
    // the last discontinuity time.  The mechanics of a subscriber session
    // throttle vary with subscriber session type and implementation.  However,
    // the general concept of the throttle prevents a node from having to deal
    // with more than a configured number of requests to establish subscriber
    // sessions from the same CPE within the a configured interval of time.  When
    // the number of requests exceeds the configured threshold within the
    // configured interval, then the node processing the requests engages the
    // throttle. Typically, when a node engages a throttle, it drops requests from
    // the CPE for some period of time, after which the node disengages the
    // throttle.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of csubAggStatsDiscontinuityTime. The type is interface{} with
    // range: 0..18446744073709551615. Units are engagements.
    Csubaggstatsthrottleengagements interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created since the discontinuity time. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    Csubaggstatstotalcreatedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect since the last discontinuity time. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    Csubaggstatstotalfailedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state since the last
    // discontinuity time.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of csubAggStatsDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615. Units are sessions.
    Csubaggstatstotalupsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state since the last discontinuity time.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    Csubaggstatstotalauthsessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event since the last discontinuity time. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    Csubaggstatstotaldiscsessions interface{}

    // This object indicates the total number of subscriber sessions that are less
    // resource intensive. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    Csubaggstatstotallightweightsessions interface{}

    // This object indicates the total number of differential traffic classes on
    // subscriber sessions. IP ACLs are used to create differential flows(Traffic
    // Classes).  Each Traffic Class can have a different set of features applied.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // sessions.
    Csubaggstatstotalflowsup interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created during the last 24 hours.  The system
    // calculates the value of this column by summing the values of all instances
    // of csubAggStatsIntCreatedSessions that expand this row and have a
    // corresponding csubAggStatsIntValid of 'true'. The type is interface{} with
    // range: 0..4294967295. Units are sessions.
    Csubaggstatsdaycreatedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect during the last 24 hours.  The system
    // calculates the value of this column by summing the values of all instances
    // of csubAggStatsIntFailedSessions that expand this row and have a
    // corresponding csubAggStatsIntValid of 'true'. The type is interface{} with
    // range: 0..4294967295. Units are sessions.
    Csubaggstatsdayfailedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state during the last 24
    // hours.  The system calculates the value of this column by summing the
    // values of all instances of csubAggStatsIntUpSessions that expand this row
    // and have a corresponding csubAggStatsIntValid of 'true'. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatsdayupsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state during the last 24 hours.  The system calculates the
    // value of this column by summing the values of all instances of
    // csubAggStatsIntAuthSessions that expand this row and have a corresponding
    // csubAggStatsIntValid of 'true'. The type is interface{} with range:
    // 0..4294967295. Units are sessions.
    Csubaggstatsdayauthsessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event during the last 24 hours.  The system calculates
    // the value of this column by summing the values of all instances of
    // csubAggStatsIntDiscSessions that expand this row and have a corresponding
    // csubAggStatsIntValid of 'true'. The type is interface{} with range:
    // 0..4294967295. Units are sessions.
    Csubaggstatsdaydiscsessions interface{}

    // This object indicates the time that has elapsed since the beginning of the
    // current 15-minute measurement interval.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, then the value of this column will be the maximum value.
    // The type is interface{} with range: 0..899. Units are seconds.
    Csubaggstatscurrtimeelapsed interface{}

    // This object indicates the number of intervals for which data was collected.
    // The value of this column will be '96' unless the measurement was started
    // (or restarted) within 1,440 minutes, in which case the value will be the
    // number of complete 15-minute intervals for which the system has at least
    // some data.  In certain cases it is possible that some intervals are
    // unavailable, in which case the value of this column will be maximum
    // interval number for which data is available. The type is interface{} with
    // range: 0..96. Units are intervals.
    Csubaggstatscurrvalidintervals interface{}

    // This object indicates the number of intervals in the range from 0 to
    // csubCurrValidIntervals for which no data is available.  This object will
    // typically be '0' except in certain circumstances when some intervals are
    // unavailable. The type is interface{} with range: 0..96. Units are
    // intervals.
    Csubaggstatscurrinvalidintervals interface{}

    // This object indicates the current number of differential traffic classes on
    // subscriber sessions currently UP. IP ACLs are used to create differential
    // flows (Traffic Classes).Each Traffic Class can have a different set of
    // features applied. The type is interface{} with range: 0..96. Units are
    // intervals.
    Csubaggstatscurrflowsup interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created during the current 15-minute interval. The
    // type is interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatscurrcreatedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect during the current 15-minute interval. The
    // type is interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatscurrfailedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state during the current
    // 15-minute interval. The type is interface{} with range: 0..4294967295.
    // Units are sessions.
    Csubaggstatscurrupsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state during the current 15-minute interval. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatscurrauthsessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event during the current 15-minute interval. The type
    // is interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatscurrdiscsessions interface{}

    // The date and time (as determined by the system's clock) of the most recent
    // occurrence of an event affecting the  continuity of the aggregation
    // statistics for this aggregation  point. The type is string.
    Csubaggstatsdiscontinuitytime interface{}
}

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetFilter() yfilter.YFilter { return csubaggstatsentry.YFilter }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) SetFilter(yf yfilter.YFilter) { csubaggstatsentry.YFilter = yf }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetGoName(yname string) string {
    if yname == "csubAggStatsPointType" { return "Csubaggstatspointtype" }
    if yname == "csubAggStatsPoint" { return "Csubaggstatspoint" }
    if yname == "csubAggStatsSessionType" { return "Csubaggstatssessiontype" }
    if yname == "csubAggStatsPendingSessions" { return "Csubaggstatspendingsessions" }
    if yname == "csubAggStatsUpSessions" { return "Csubaggstatsupsessions" }
    if yname == "csubAggStatsAuthSessions" { return "Csubaggstatsauthsessions" }
    if yname == "csubAggStatsUnAuthSessions" { return "Csubaggstatsunauthsessions" }
    if yname == "csubAggStatsLightWeightSessions" { return "Csubaggstatslightweightsessions" }
    if yname == "csubAggStatsRedSessions" { return "Csubaggstatsredsessions" }
    if yname == "csubAggStatsHighUpSessions" { return "Csubaggstatshighupsessions" }
    if yname == "csubAggStatsAvgSessionUptime" { return "Csubaggstatsavgsessionuptime" }
    if yname == "csubAggStatsAvgSessionRPM" { return "Csubaggstatsavgsessionrpm" }
    if yname == "csubAggStatsAvgSessionRPH" { return "Csubaggstatsavgsessionrph" }
    if yname == "csubAggStatsThrottleEngagements" { return "Csubaggstatsthrottleengagements" }
    if yname == "csubAggStatsTotalCreatedSessions" { return "Csubaggstatstotalcreatedsessions" }
    if yname == "csubAggStatsTotalFailedSessions" { return "Csubaggstatstotalfailedsessions" }
    if yname == "csubAggStatsTotalUpSessions" { return "Csubaggstatstotalupsessions" }
    if yname == "csubAggStatsTotalAuthSessions" { return "Csubaggstatstotalauthsessions" }
    if yname == "csubAggStatsTotalDiscSessions" { return "Csubaggstatstotaldiscsessions" }
    if yname == "csubAggStatsTotalLightWeightSessions" { return "Csubaggstatstotallightweightsessions" }
    if yname == "csubAggStatsTotalFlowsUp" { return "Csubaggstatstotalflowsup" }
    if yname == "csubAggStatsDayCreatedSessions" { return "Csubaggstatsdaycreatedsessions" }
    if yname == "csubAggStatsDayFailedSessions" { return "Csubaggstatsdayfailedsessions" }
    if yname == "csubAggStatsDayUpSessions" { return "Csubaggstatsdayupsessions" }
    if yname == "csubAggStatsDayAuthSessions" { return "Csubaggstatsdayauthsessions" }
    if yname == "csubAggStatsDayDiscSessions" { return "Csubaggstatsdaydiscsessions" }
    if yname == "csubAggStatsCurrTimeElapsed" { return "Csubaggstatscurrtimeelapsed" }
    if yname == "csubAggStatsCurrValidIntervals" { return "Csubaggstatscurrvalidintervals" }
    if yname == "csubAggStatsCurrInvalidIntervals" { return "Csubaggstatscurrinvalidintervals" }
    if yname == "csubAggStatsCurrFlowsUp" { return "Csubaggstatscurrflowsup" }
    if yname == "csubAggStatsCurrCreatedSessions" { return "Csubaggstatscurrcreatedsessions" }
    if yname == "csubAggStatsCurrFailedSessions" { return "Csubaggstatscurrfailedsessions" }
    if yname == "csubAggStatsCurrUpSessions" { return "Csubaggstatscurrupsessions" }
    if yname == "csubAggStatsCurrAuthSessions" { return "Csubaggstatscurrauthsessions" }
    if yname == "csubAggStatsCurrDiscSessions" { return "Csubaggstatscurrdiscsessions" }
    if yname == "csubAggStatsDiscontinuityTime" { return "Csubaggstatsdiscontinuitytime" }
    return ""
}

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetSegmentPath() string {
    return "csubAggStatsEntry" + "[csubAggStatsPointType='" + fmt.Sprintf("%v", csubaggstatsentry.Csubaggstatspointtype) + "']" + "[csubAggStatsPoint='" + fmt.Sprintf("%v", csubaggstatsentry.Csubaggstatspoint) + "']" + "[csubAggStatsSessionType='" + fmt.Sprintf("%v", csubaggstatsentry.Csubaggstatssessiontype) + "']"
}

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubAggStatsPointType"] = csubaggstatsentry.Csubaggstatspointtype
    leafs["csubAggStatsPoint"] = csubaggstatsentry.Csubaggstatspoint
    leafs["csubAggStatsSessionType"] = csubaggstatsentry.Csubaggstatssessiontype
    leafs["csubAggStatsPendingSessions"] = csubaggstatsentry.Csubaggstatspendingsessions
    leafs["csubAggStatsUpSessions"] = csubaggstatsentry.Csubaggstatsupsessions
    leafs["csubAggStatsAuthSessions"] = csubaggstatsentry.Csubaggstatsauthsessions
    leafs["csubAggStatsUnAuthSessions"] = csubaggstatsentry.Csubaggstatsunauthsessions
    leafs["csubAggStatsLightWeightSessions"] = csubaggstatsentry.Csubaggstatslightweightsessions
    leafs["csubAggStatsRedSessions"] = csubaggstatsentry.Csubaggstatsredsessions
    leafs["csubAggStatsHighUpSessions"] = csubaggstatsentry.Csubaggstatshighupsessions
    leafs["csubAggStatsAvgSessionUptime"] = csubaggstatsentry.Csubaggstatsavgsessionuptime
    leafs["csubAggStatsAvgSessionRPM"] = csubaggstatsentry.Csubaggstatsavgsessionrpm
    leafs["csubAggStatsAvgSessionRPH"] = csubaggstatsentry.Csubaggstatsavgsessionrph
    leafs["csubAggStatsThrottleEngagements"] = csubaggstatsentry.Csubaggstatsthrottleengagements
    leafs["csubAggStatsTotalCreatedSessions"] = csubaggstatsentry.Csubaggstatstotalcreatedsessions
    leafs["csubAggStatsTotalFailedSessions"] = csubaggstatsentry.Csubaggstatstotalfailedsessions
    leafs["csubAggStatsTotalUpSessions"] = csubaggstatsentry.Csubaggstatstotalupsessions
    leafs["csubAggStatsTotalAuthSessions"] = csubaggstatsentry.Csubaggstatstotalauthsessions
    leafs["csubAggStatsTotalDiscSessions"] = csubaggstatsentry.Csubaggstatstotaldiscsessions
    leafs["csubAggStatsTotalLightWeightSessions"] = csubaggstatsentry.Csubaggstatstotallightweightsessions
    leafs["csubAggStatsTotalFlowsUp"] = csubaggstatsentry.Csubaggstatstotalflowsup
    leafs["csubAggStatsDayCreatedSessions"] = csubaggstatsentry.Csubaggstatsdaycreatedsessions
    leafs["csubAggStatsDayFailedSessions"] = csubaggstatsentry.Csubaggstatsdayfailedsessions
    leafs["csubAggStatsDayUpSessions"] = csubaggstatsentry.Csubaggstatsdayupsessions
    leafs["csubAggStatsDayAuthSessions"] = csubaggstatsentry.Csubaggstatsdayauthsessions
    leafs["csubAggStatsDayDiscSessions"] = csubaggstatsentry.Csubaggstatsdaydiscsessions
    leafs["csubAggStatsCurrTimeElapsed"] = csubaggstatsentry.Csubaggstatscurrtimeelapsed
    leafs["csubAggStatsCurrValidIntervals"] = csubaggstatsentry.Csubaggstatscurrvalidintervals
    leafs["csubAggStatsCurrInvalidIntervals"] = csubaggstatsentry.Csubaggstatscurrinvalidintervals
    leafs["csubAggStatsCurrFlowsUp"] = csubaggstatsentry.Csubaggstatscurrflowsup
    leafs["csubAggStatsCurrCreatedSessions"] = csubaggstatsentry.Csubaggstatscurrcreatedsessions
    leafs["csubAggStatsCurrFailedSessions"] = csubaggstatsentry.Csubaggstatscurrfailedsessions
    leafs["csubAggStatsCurrUpSessions"] = csubaggstatsentry.Csubaggstatscurrupsessions
    leafs["csubAggStatsCurrAuthSessions"] = csubaggstatsentry.Csubaggstatscurrauthsessions
    leafs["csubAggStatsCurrDiscSessions"] = csubaggstatsentry.Csubaggstatscurrdiscsessions
    leafs["csubAggStatsDiscontinuityTime"] = csubaggstatsentry.Csubaggstatsdiscontinuitytime
    return leafs
}

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetYangName() string { return "csubAggStatsEntry" }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) SetParent(parent types.Entity) { csubaggstatsentry.parent = parent }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetParent() types.Entity { return csubaggstatsentry.parent }

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetParentYangName() string { return "csubAggStatsTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspointtype represents     'interface'                InterfaceIndex
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspointtype string

const (
    CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspointtype_physical CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspointtype = "physical"

    CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspointtype_interface_ CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspointtype = "interface"
)

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable
// This table contains aggregated subscriber session performance
// data collected for as much as a day's worth of 15-minute
// measurement intervals.
// 
// This table has an expansion dependent relationship on the
// csubAggStatsTable, containing zero or more rows for each row
// contained by the csubAggStatsTable.
// 
// Observe that the collection and maintenance of aggregated
// subscriber performance data is OPTIONAL for all scopes of
// aggregation.  However, an implementation should maintain at
// least one interval for the 'scope of aggregation' that contains
// all subscriber sessions maintained by the system.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry contains the aggregated subscriber session performance data
    // collected over a single 15-minute measurement interval within a 'scope of
    // aggregation'.  For further details regarding 'scope of aggregation', see
    // the descriptive text associated with the csubAggStatsEntry. The type is
    // slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry.
    Csubaggstatsintentry []CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry
}

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetFilter() yfilter.YFilter { return csubaggstatsinttable.YFilter }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) SetFilter(yf yfilter.YFilter) { csubaggstatsinttable.YFilter = yf }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetGoName(yname string) string {
    if yname == "csubAggStatsIntEntry" { return "Csubaggstatsintentry" }
    return ""
}

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetSegmentPath() string {
    return "csubAggStatsIntTable"
}

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubAggStatsIntEntry" {
        for _, c := range csubaggstatsinttable.Csubaggstatsintentry {
            if csubaggstatsinttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry{}
        csubaggstatsinttable.Csubaggstatsintentry = append(csubaggstatsinttable.Csubaggstatsintentry, child)
        return &csubaggstatsinttable.Csubaggstatsintentry[len(csubaggstatsinttable.Csubaggstatsintentry)-1]
    }
    return nil
}

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubaggstatsinttable.Csubaggstatsintentry {
        children[csubaggstatsinttable.Csubaggstatsintentry[i].GetSegmentPath()] = &csubaggstatsinttable.Csubaggstatsintentry[i]
    }
    return children
}

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetBundleName() string { return "cisco_ios_xe" }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetYangName() string { return "csubAggStatsIntTable" }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) SetParent(parent types.Entity) { csubaggstatsinttable.parent = parent }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetParent() types.Entity { return csubaggstatsinttable.parent }

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry
// An entry contains the aggregated subscriber session performance
// data collected over a single 15-minute measurement interval
// within a 'scope of aggregation'.  For further details regarding
// 'scope of aggregation', see the descriptive text associated with
// the csubAggStatsEntry.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is Csubaggstatspointtype. Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspointtype
    Csubaggstatspointtype interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatspoint
    Csubaggstatspoint interface{}

    // This attribute is a key. The type is SubSessionType. Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry_Csubaggstatssessiontype
    Csubaggstatssessiontype interface{}

    // This attribute is a key. This object indicates the interval number
    // identifying the 15-minute measurement interval for which aggregated
    // subscriber session performance data was successfully collected by the
    // system.  The interval identified by the value '1' represents the most
    // recent 15-minute measurement interval, and the interval identified by the
    // value (n) represents the interval immediately preceding the interval
    // identified by the value (n-1). The type is interface{} with range: 1..96.
    Csubaggstatsintnumber interface{}

    // This object indicates whether the data for the 15-minute measurement
    // interval is valid. The type is bool.
    Csubaggstatsintvalid interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created during the 15-minute measurement interval.
    // The type is interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatsintcreatedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect during the 15-minute measurement interval.
    // The type is interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatsintfailedsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state during the
    // 15-minute measurement interval. The type is interface{} with range:
    // 0..4294967295. Units are sessions.
    Csubaggstatsintupsessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state during the 15-minute measurement interval. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatsintauthsessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event during the 15-minute measurement interval. The
    // type is interface{} with range: 0..4294967295. Units are sessions.
    Csubaggstatsintdiscsessions interface{}
}

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetFilter() yfilter.YFilter { return csubaggstatsintentry.YFilter }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) SetFilter(yf yfilter.YFilter) { csubaggstatsintentry.YFilter = yf }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetGoName(yname string) string {
    if yname == "csubAggStatsPointType" { return "Csubaggstatspointtype" }
    if yname == "csubAggStatsPoint" { return "Csubaggstatspoint" }
    if yname == "csubAggStatsSessionType" { return "Csubaggstatssessiontype" }
    if yname == "csubAggStatsIntNumber" { return "Csubaggstatsintnumber" }
    if yname == "csubAggStatsIntValid" { return "Csubaggstatsintvalid" }
    if yname == "csubAggStatsIntCreatedSessions" { return "Csubaggstatsintcreatedsessions" }
    if yname == "csubAggStatsIntFailedSessions" { return "Csubaggstatsintfailedsessions" }
    if yname == "csubAggStatsIntUpSessions" { return "Csubaggstatsintupsessions" }
    if yname == "csubAggStatsIntAuthSessions" { return "Csubaggstatsintauthsessions" }
    if yname == "csubAggStatsIntDiscSessions" { return "Csubaggstatsintdiscsessions" }
    return ""
}

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetSegmentPath() string {
    return "csubAggStatsIntEntry" + "[csubAggStatsPointType='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatspointtype) + "']" + "[csubAggStatsPoint='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatspoint) + "']" + "[csubAggStatsSessionType='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatssessiontype) + "']" + "[csubAggStatsIntNumber='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatsintnumber) + "']"
}

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubAggStatsPointType"] = csubaggstatsintentry.Csubaggstatspointtype
    leafs["csubAggStatsPoint"] = csubaggstatsintentry.Csubaggstatspoint
    leafs["csubAggStatsSessionType"] = csubaggstatsintentry.Csubaggstatssessiontype
    leafs["csubAggStatsIntNumber"] = csubaggstatsintentry.Csubaggstatsintnumber
    leafs["csubAggStatsIntValid"] = csubaggstatsintentry.Csubaggstatsintvalid
    leafs["csubAggStatsIntCreatedSessions"] = csubaggstatsintentry.Csubaggstatsintcreatedsessions
    leafs["csubAggStatsIntFailedSessions"] = csubaggstatsintentry.Csubaggstatsintfailedsessions
    leafs["csubAggStatsIntUpSessions"] = csubaggstatsintentry.Csubaggstatsintupsessions
    leafs["csubAggStatsIntAuthSessions"] = csubaggstatsintentry.Csubaggstatsintauthsessions
    leafs["csubAggStatsIntDiscSessions"] = csubaggstatsintentry.Csubaggstatsintdiscsessions
    return leafs
}

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetYangName() string { return "csubAggStatsIntEntry" }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) SetParent(parent types.Entity) { csubaggstatsintentry.parent = parent }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetParent() types.Entity { return csubaggstatsintentry.parent }

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetParentYangName() string { return "csubAggStatsIntTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable
// Please enter the Table Description here.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in this table exists for each row in the csubAggStatsTable. Each row
    // defines the set of thresholds and evaluation attributes for an aggregation
    // point. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry.
    Csubaggstatsthreshentry []CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry
}

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetFilter() yfilter.YFilter { return csubaggstatsthreshtable.YFilter }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) SetFilter(yf yfilter.YFilter) { csubaggstatsthreshtable.YFilter = yf }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetGoName(yname string) string {
    if yname == "csubAggStatsThreshEntry" { return "Csubaggstatsthreshentry" }
    return ""
}

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetSegmentPath() string {
    return "csubAggStatsThreshTable"
}

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubAggStatsThreshEntry" {
        for _, c := range csubaggstatsthreshtable.Csubaggstatsthreshentry {
            if csubaggstatsthreshtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry{}
        csubaggstatsthreshtable.Csubaggstatsthreshentry = append(csubaggstatsthreshtable.Csubaggstatsthreshentry, child)
        return &csubaggstatsthreshtable.Csubaggstatsthreshentry[len(csubaggstatsthreshtable.Csubaggstatsthreshentry)-1]
    }
    return nil
}

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubaggstatsthreshtable.Csubaggstatsthreshentry {
        children[csubaggstatsthreshtable.Csubaggstatsthreshentry[i].GetSegmentPath()] = &csubaggstatsthreshtable.Csubaggstatsthreshentry[i]
    }
    return children
}

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetBundleName() string { return "cisco_ios_xe" }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetYangName() string { return "csubAggStatsThreshTable" }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) SetParent(parent types.Entity) { csubaggstatsthreshtable.parent = parent }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetParent() types.Entity { return csubaggstatsthreshtable.parent }

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry
// A row in this table exists for each row in the
// csubAggStatsTable.
// Each row defines the set of thresholds and evaluation attributes
// for an aggregation point.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This threshold, if non-zero, indicates the rising
    // threshold for the value of csubAggStatsUpSessions for the aggregation
    // point, When the current sample of csubAggStatsUpSessions is greater than or
    // equal to this threshold, and the value of csubAggStatsUpSessions for the
    // last sample interval was less than this thershold, the
    // csubSessionRisingNotif is triggered.           If the value of this
    // threshold is 0, the threshold is not evaluated. The type is interface{}
    // with range: 0..4294967295.
    Csubsessionrisingthresh interface{}

    // This threshold, if non-zero, indicates the falling threshold for the value
    // of csubAggStatsUpSessions for the aggregation point, When the current
    // sample of csubAggStatsUpSessions is less than or equal to this threshold,
    // and the value of csubAggStatsUpSessions for the last sample interval was
    // greater than this thershold, the csubSessionFallingNotif is triggered.     
    // If the value of this threshold is 0, the threshold is not evaluated. The
    // type is interface{} with range: 0..4294967295.
    Csubsessionfallingthresh interface{}

    // This threshold, if non-zero, indicates the delta falling threshold as a
    // percentage of the value of csubAggStatsUpSessions for the aggregation
    // point, The delta as a percentage of csubAggStatsUpSessions (delta_percent)
    // is calculated as follows:       current value of csubAggStatsUpSessions =
    // value(n)              previous sampled value of csubAggStatsUpSessions =
    // value(n-1)               delta_percent = ((value(n-1) - value(n)) /
    // value(n-1)) * 100           If the delta_percent value of the current
    // evaluation interval is          greater than the value of this threshold, a
    // csubSessionDeltaPercentLossNotif is triggered.           If the value of
    // this threshold is 0, the threshold is not evaluated. The type is
    // interface{} with range: 0..100.
    Csubsessiondeltapercentfallingthresh interface{}

    // The value of this object sets the number of seconds between samples for
    // threshold evaluation. For implementations capable of per-session event
    // evaluation of thresholds this object represents the maximum number of
    // seconds between samples. The type is interface{} with range: 0..900.
    Csubsessionthreshevalinterval interface{}
}

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetFilter() yfilter.YFilter { return csubaggstatsthreshentry.YFilter }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) SetFilter(yf yfilter.YFilter) { csubaggstatsthreshentry.YFilter = yf }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetGoName(yname string) string {
    if yname == "csubSessionRisingThresh" { return "Csubsessionrisingthresh" }
    if yname == "csubSessionFallingThresh" { return "Csubsessionfallingthresh" }
    if yname == "csubSessionDeltaPercentFallingThresh" { return "Csubsessiondeltapercentfallingthresh" }
    if yname == "csubSessionThreshEvalInterval" { return "Csubsessionthreshevalinterval" }
    return ""
}

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetSegmentPath() string {
    return "csubAggStatsThreshEntry" + "[csubSessionRisingThresh='" + fmt.Sprintf("%v", csubaggstatsthreshentry.Csubsessionrisingthresh) + "']"
}

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubSessionRisingThresh"] = csubaggstatsthreshentry.Csubsessionrisingthresh
    leafs["csubSessionFallingThresh"] = csubaggstatsthreshentry.Csubsessionfallingthresh
    leafs["csubSessionDeltaPercentFallingThresh"] = csubaggstatsthreshentry.Csubsessiondeltapercentfallingthresh
    leafs["csubSessionThreshEvalInterval"] = csubaggstatsthreshentry.Csubsessionthreshevalinterval
    return leafs
}

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetYangName() string { return "csubAggStatsThreshEntry" }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) SetParent(parent types.Entity) { csubaggstatsthreshentry.parent = parent }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetParent() types.Entity { return csubaggstatsthreshentry.parent }

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetParentYangName() string { return "csubAggStatsThreshTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobtable
// This table contains the subscriber session jobs submitted by
// the EMS/NMS.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describing a subscriber session job.  At this time, subscriber
    // session jobs can perform one of two tasks:  - Subscriber Session Query    
    // This type of job invokes the report generator, which builds     a list of
    // subscriber sessions matching criteria specified by     the corresponding
    // row in the csubJobMatchParamsTable.  The     list built by the report
    // generator must conform to     parameters specified by the corresponding row
    // in     csubJobQueryParamsTable, which at this time only affects     sorting
    // order.  - Subscriber Session Clear     This type of job causes the system
    // to terminate those     subscriber sessions matching criteria specified by
    // the     corresponding row in the csubJobMatchParamsTable.  The following
    // procedure summarizes how the EMS/NMS can start and monitor a subscriber
    // session job:  1)  The EMS/NMS must start by reading csubJobIdNext.  If it
    // is     zero, continue polling csubJobIdNext until it is non-zero.  2)  The
    // EMS/NMS creates a row in the csubJobTable using the     instance identifier
    // retrieved in the last step.  Since every     object contained by the entry
    // with a MAX-ACCESS of      'read-create' specifies a default value, it makes
    // little     difference whether the EMS/NMS employs create-and-wait or    
    // create-and-go semantics.  3)  The EMS/NMS sets the type of subscriber
    // session job by     setting the corresponding instance of csubJobType.  4a)
    // If the job is a 'query', then the EMS/NMS must configure     the query
    // before starting it by setting columns contained     by the corresponding
    // rows in the csubJobMatchParamsTable and     csubJobQueryParamsTable.  4b)
    // If job is a 'clear', then the EMS/NMS must configure     the job before
    // starting it by setting columns contained by     the corresponding row in
    // the csubJobMatchParamsTable.  5)  The EMS/NMS can now start the job by
    // setting the      corresponding instance of csubJobControl to 'start'.  6) 
    // The EMS/NMS can monitor the progress of the job by polling     the
    // corresponding instance of csubJobState.  It can also     wait for a
    // csubJobFinishedNotify notification.  When the     state of the job
    // transitions to 'finished', then the system     has finished executing the
    // job.  7)  The EMS/NMS can determine the final status of the job by    
    // reading the corresponding instance of csubJobFinishedReason.     If job is
    // a 'query' and the corresponding instance of     csubJobFinishedReason is
    // 'normal', then the EMS/NMS can     safely read the report by retrieving the
    // corresponding     rows from the csubJobReportTable.  8a) After a job has
    // finished, the EMS/NMS has the option of     destroying it.  It can do this
    // by simply setting the     corresponding instance of  csubJobStatus to
    // 'destroy'.     Alternatively, the EMS/NMS may retain the job and execute it
    // again in the future (by returning to step 5).  Additionally,     nothing
    // would prevent the EMS/NMS from changing the job's     type, which causes
    // the automatic destruction of the     corresponding report.  8b) If the job
    // is a 'query' and the EMS/NMS opts to retain the     job, then it may
    // consider releasing the corresponding report     after reading it.  It can
    // do this by setting the     corresponding instance of csubJobControl to
    // 'release'. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry.
    Csubjobentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry
}

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetFilter() yfilter.YFilter { return csubjobtable.YFilter }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) SetFilter(yf yfilter.YFilter) { csubjobtable.YFilter = yf }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetGoName(yname string) string {
    if yname == "csubJobEntry" { return "Csubjobentry" }
    return ""
}

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetSegmentPath() string {
    return "csubJobTable"
}

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubJobEntry" {
        for _, c := range csubjobtable.Csubjobentry {
            if csubjobtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry{}
        csubjobtable.Csubjobentry = append(csubjobtable.Csubjobentry, child)
        return &csubjobtable.Csubjobentry[len(csubjobtable.Csubjobentry)-1]
    }
    return nil
}

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubjobtable.Csubjobentry {
        children[csubjobtable.Csubjobentry[i].GetSegmentPath()] = &csubjobtable.Csubjobentry[i]
    }
    return children
}

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetYangName() string { return "csubJobTable" }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) SetParent(parent types.Entity) { csubjobtable.parent = parent }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetParent() types.Entity { return csubjobtable.parent }

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry
// An entry describing a subscriber session job.  At this time,
// subscriber session jobs can perform one of two tasks:
// 
// - Subscriber Session Query
//     This type of job invokes the report generator, which builds
//     a list of subscriber sessions matching criteria specified by
//     the corresponding row in the csubJobMatchParamsTable.  The
//     list built by the report generator must conform to
//     parameters specified by the corresponding row in
//     csubJobQueryParamsTable, which at this time only affects
//     sorting order.
// 
// - Subscriber Session Clear
//     This type of job causes the system to terminate those
//     subscriber sessions matching criteria specified by the
//     corresponding row in the csubJobMatchParamsTable.
// 
// The following procedure summarizes how the EMS/NMS can start and
// monitor a subscriber session job:
// 
// 1)  The EMS/NMS must start by reading csubJobIdNext.  If it is
//     zero, continue polling csubJobIdNext until it is non-zero.
// 
// 2)  The EMS/NMS creates a row in the csubJobTable using the
//     instance identifier retrieved in the last step.  Since every
//     object contained by the entry with a MAX-ACCESS of 
//     'read-create' specifies a default value, it makes little
//     difference whether the EMS/NMS employs create-and-wait or
//     create-and-go semantics.
// 
// 3)  The EMS/NMS sets the type of subscriber session job by
//     setting the corresponding instance of csubJobType.
// 
// 4a) If the job is a 'query', then the EMS/NMS must configure
//     the query before starting it by setting columns contained
//     by the corresponding rows in the csubJobMatchParamsTable and
//     csubJobQueryParamsTable.
// 
// 4b) If job is a 'clear', then the EMS/NMS must configure
//     the job before starting it by setting columns contained by
//     the corresponding row in the csubJobMatchParamsTable.
// 
// 5)  The EMS/NMS can now start the job by setting the 
//     corresponding instance of csubJobControl to 'start'.
// 
// 6)  The EMS/NMS can monitor the progress of the job by polling
//     the corresponding instance of csubJobState.  It can also
//     wait for a csubJobFinishedNotify notification.  When the
//     state of the job transitions to 'finished', then the system
//     has finished executing the job.
// 
// 7)  The EMS/NMS can determine the final status of the job by
//     reading the corresponding instance of csubJobFinishedReason.
//     If job is a 'query' and the corresponding instance of
//     csubJobFinishedReason is 'normal', then the EMS/NMS can
//     safely read the report by retrieving the corresponding
//     rows from the csubJobReportTable.
// 
// 8a) After a job has finished, the EMS/NMS has the option of
//     destroying it.  It can do this by simply setting the
//     corresponding instance of  csubJobStatus to 'destroy'.
//     Alternatively, the EMS/NMS may retain the job and execute it
//     again in the future (by returning to step 5).  Additionally,
//     nothing would prevent the EMS/NMS from changing the job's
//     type, which causes the automatic destruction of the
//     corresponding report.
// 
// 8b) If the job is a 'query' and the EMS/NMS opts to retain the
//     job, then it may consider releasing the corresponding report
//     after reading it.  It can do this by setting the
//     corresponding instance of csubJobControl to 'release'.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates an arbitrary, positive
    // integer-value uniquely identifying the subscriber session job. The type is
    // interface{} with range: 1..4294967295.
    Csubjobid interface{}

    // This object specifies the status of the subscriber session job.  The
    // following columns must be valid before activating a subscriber session job:
    // - csubJobStorage     - csubJobType     - csubJobControl  However, these
    // objects specify a default value.  Thus, it is possible to use create-and-go
    // semantics without setting any additional columns.  An implementation must
    // allow the EMS/NMS to modify any column when this column is 'active',
    // including columns defined in tables that have a one-to-one or sparse
    // dependent relationship on this table. The type is RowStatus.
    Csubjobstatus interface{}

    // This object specifies what happens to the subscriber session job upon
    // restart. The type is StorageType.
    Csubjobstorage interface{}

    // This object specifies the type of subscriber session job:  'noop'     This
    // type of job does nothing and simply serves as a     convenient default
    // value for newly created jobs, thereby     allowing create-and-go row
    // creation without having to     specify the type of job.  'query'     This
    // type of job starts a subscriber session query.  The     system searches for
    // any subscriber sessions matching the     configured criteria and sorts them
    // into a resulting     report.      Upon activation of a subscriber session
    // with this value,     the system automatically creates corresponding rows in
    // the csubJobMatchParamsTable and csubQueryParamsTable.  'clear'     This
    // type of job causes the system to terminated all     subscriber sessions
    // matching configured criteria.      Upon activation of a subscriber session
    // with this value,     the system automatically creates a corresponding row
    // in     the csubJobMatchParamsTable. The type is Csubjobtype.
    Csubjobtype interface{}

    // This object specifies an action relating to the subscriber session job:    
    // 'noop'         This action does nothing.      'start'         If the
    // corresponding instance of csubJobType is 'noop',         then this action
    // simply causes the system to set the         corresponding instances of
    // csubJobState and         csubJobFinishedReason to 'finished' and 'normal', 
    // respectively.          If the corresponding instance of csubJobType is not 
    // 'noop' and the system is not executing a subscriber         session job,
    // then this action causes the system         immediately execute the
    // subscriber session job.          If the corresponding instance of
    // csubJobType is not         'noop' and the system is already executing a
    // subscriber         session job, then this action causes the system to put  
    // the job on the subscriber session job queue.      'abort'         If the
    // subscriber session job is in the subscriber         session job queue, then
    // this action causes the system to         remove the job from the queue.    
    // If the system is executing the subscriber session job,         then this
    // action causes the system to stop the job.      'release'         This
    // action causes the system to destroy any         corresponding rows in the
    // csubJobReportTable.          The system only accepts this action for a
    // previously         executed subscriber session job having a corresponding  
    // instance of csubJobType set to 'query'.  Any attempt to         issue this
    // action under other circumstances will result         in a response
    // indicating an  error-status of         'inconsistentValue'.  When read,
    // this column is always 'noop'. The type is Csubjobcontrol.
    Csubjobcontrol interface{}

    // This object indicates the current state of the subscriber session job:     
    // 'idle'         This state indicates that the system has not executed       
    // the subscriber session job since it was created.      'pending'        
    // This state indicates that the subscriber session job is         waiting in
    // the subscriber session job queue.      'inProgress'         This state
    // indicates that the system is executing the         subscriber session job. 
    // Observe that the system may         execute more than one subscriber
    // session job at a time.      'finished'         This state indicates that
    // the system has executed the         subscriber session job and it has
    // finished.  The         corresponding instance of csubJobFinishedReason     
    // indicates further details regarding the reason why the         job
    // finished. The type is Csubjobstate.
    Csubjobstate interface{}

    // This object indicates the value of sysUpTime when the system started
    // executing the subscriber session job.  This value will be '0' when the
    // corresponding instance of csubJobState is 'idle' or 'pending'. The type is
    // interface{} with range: 0..4294967295.
    Csubjobstartedtime interface{}

    // This object indicates the value of sysUpTime when the system finished
    // executing the subscriber session job, for whatever reason.  This value will
    // be '0' when the corresponding instance of csubJobState is 'idle',
    // 'pending', or 'inProgress'. The type is interface{} with range:
    // 0..4294967295.
    Csubjobfinishedtime interface{}

    // This object indicates the reason why the system finished executing the
    // subscriber session job:      'invalid'         Indicates that the
    // corresponding instance of         csubJobState is either 'idle', 'pending',
    // or         'inProgress'.      'other'         Indicates that the system
    // finished executing the         subscriber session job abnormally for a
    // reason not         recognized by this MIB module.      'normal'        
    // Indicates that the system finished executing the         subscriber session
    // job with no problems.      'aborted'         Indicates that the system
    // finished executing the         subscriber session job as the result of the
    // EMS/NMS         writing 'abort' to the corresponding instance of        
    // csubJobControl.      'insufficientResources'         Indicates that the
    // system finished executing the         subscriber session job abnormally due
    // to insufficient         resources to continue.      'error'        
    // Indicates that the system encountered an error that         prevented it
    // from completing the job. The type is Csubjobfinishedreason.
    Csubjobfinishedreason interface{}
}

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetFilter() yfilter.YFilter { return csubjobentry.YFilter }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) SetFilter(yf yfilter.YFilter) { csubjobentry.YFilter = yf }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetGoName(yname string) string {
    if yname == "csubJobId" { return "Csubjobid" }
    if yname == "csubJobStatus" { return "Csubjobstatus" }
    if yname == "csubJobStorage" { return "Csubjobstorage" }
    if yname == "csubJobType" { return "Csubjobtype" }
    if yname == "csubJobControl" { return "Csubjobcontrol" }
    if yname == "csubJobState" { return "Csubjobstate" }
    if yname == "csubJobStartedTime" { return "Csubjobstartedtime" }
    if yname == "csubJobFinishedTime" { return "Csubjobfinishedtime" }
    if yname == "csubJobFinishedReason" { return "Csubjobfinishedreason" }
    return ""
}

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetSegmentPath() string {
    return "csubJobEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobentry.Csubjobid) + "']"
}

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubJobId"] = csubjobentry.Csubjobid
    leafs["csubJobStatus"] = csubjobentry.Csubjobstatus
    leafs["csubJobStorage"] = csubjobentry.Csubjobstorage
    leafs["csubJobType"] = csubjobentry.Csubjobtype
    leafs["csubJobControl"] = csubjobentry.Csubjobcontrol
    leafs["csubJobState"] = csubjobentry.Csubjobstate
    leafs["csubJobStartedTime"] = csubjobentry.Csubjobstartedtime
    leafs["csubJobFinishedTime"] = csubjobentry.Csubjobfinishedtime
    leafs["csubJobFinishedReason"] = csubjobentry.Csubjobfinishedreason
    return leafs
}

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetYangName() string { return "csubJobEntry" }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) SetParent(parent types.Entity) { csubjobentry.parent = parent }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetParent() types.Entity { return csubjobentry.parent }

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetParentYangName() string { return "csubJobTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol represents When read, this column is always 'noop'.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol string

const (
    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol_noop CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol = "noop"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol_start CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol = "start"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol_abort CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol = "abort"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol_release CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobcontrol = "release"
)

// CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason represents         prevented it from completing the job.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason string

const (
    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason_invalid CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason = "invalid"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason_other CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason = "other"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason_normal CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason = "normal"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason_aborted CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason = "aborted"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason_insufficientResources CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason = "insufficientResources"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason_error CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobfinishedreason = "error"
)

// CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate represents         job finished.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate string

const (
    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate_idle CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate = "idle"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate_pending CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate = "pending"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate_inProgress CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate = "inProgress"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate_finished CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobstate = "finished"
)

// CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype represents     the csubJobMatchParamsTable.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype string

const (
    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype_noop CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype = "noop"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype_query CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype = "query"

    CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype_clear CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobtype = "clear"
)

// CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable
// This table contains subscriber session job parameters
// describing match criteria.
// 
// This table has a sparse-dependent relationship on the
// csubJobTable, containing a row for each job having a
// csubJobType of 'query' or 'clear'.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describes a set of subscriber session match criteria. The set
    // contains those subscriber session identities specified by
    // csubJobMatchIdentities.  If the corresponding row in the csubJobTable has a
    // csubJobType of 'query', then the system builds a report containing those
    // subscriber sessions matching these criteria.  If the corresponding row in
    // the csubJobTable has a csubJobType of 'clear', then the system terminates
    // those subscriber sessions matching these criteria.  The system
    // automatically creates an entry when the EMS/NMS sets the corresponding
    // instance of csubJobType to 'query' or 'clear'. Likewise, the system
    // automatically destroys an entry under the following circumstances:  1)  The
    // EMS/NMS destroys the corresponding row in the     csubJobTable.  2)  The
    // EMS/NMS sets the corresponding instance of csubJobType     to 'noop'. The
    // type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry.
    Csubjobmatchparamsentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry
}

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetFilter() yfilter.YFilter { return csubjobmatchparamstable.YFilter }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) SetFilter(yf yfilter.YFilter) { csubjobmatchparamstable.YFilter = yf }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetGoName(yname string) string {
    if yname == "csubJobMatchParamsEntry" { return "Csubjobmatchparamsentry" }
    return ""
}

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetSegmentPath() string {
    return "csubJobMatchParamsTable"
}

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubJobMatchParamsEntry" {
        for _, c := range csubjobmatchparamstable.Csubjobmatchparamsentry {
            if csubjobmatchparamstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry{}
        csubjobmatchparamstable.Csubjobmatchparamsentry = append(csubjobmatchparamstable.Csubjobmatchparamsentry, child)
        return &csubjobmatchparamstable.Csubjobmatchparamsentry[len(csubjobmatchparamstable.Csubjobmatchparamsentry)-1]
    }
    return nil
}

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubjobmatchparamstable.Csubjobmatchparamsentry {
        children[csubjobmatchparamstable.Csubjobmatchparamsentry[i].GetSegmentPath()] = &csubjobmatchparamstable.Csubjobmatchparamsentry[i]
    }
    return children
}

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetYangName() string { return "csubJobMatchParamsTable" }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) SetParent(parent types.Entity) { csubjobmatchparamstable.parent = parent }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetParent() types.Entity { return csubjobmatchparamstable.parent }

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry
// An entry describes a set of subscriber session match criteria.
// The set contains those subscriber session identities specified
// by csubJobMatchIdentities.
// 
// If the corresponding row in the csubJobTable has a csubJobType
// of 'query', then the system builds a report containing those
// subscriber sessions matching these criteria.
// 
// If the corresponding row in the csubJobTable has a csubJobType
// of 'clear', then the system terminates those subscriber
// sessions matching these criteria.
// 
// The system automatically creates an entry when the EMS/NMS sets
// the corresponding instance of csubJobType to 'query' or 'clear'.
// Likewise, the system automatically destroys an entry under the
// following circumstances:
// 
// 1)  The EMS/NMS destroys the corresponding row in the
//     csubJobTable.
// 
// 2)  The EMS/NMS sets the corresponding instance of csubJobType
//     to 'noop'.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobid
    Csubjobid interface{}

    // This object specifies the subscriber identities that the system uses to
    // determine the subscriber sessions the job executes on.  Each bit in this
    // bit string corresponds to one or more columns in this table.  If the bit is
    // '0', then the value of the corresponding columns are invalid.  If the bit
    // is '1', then the value of corresponding columns are valid.  The following
    // list specifies the mappings between the bits and the columns:     
    // 'subscriberLabel' => csubJobMatchSubscriberLabel     'macAddress'      =>
    // csubJobMatchMacAddress     'nativeVrf'       => csubJobMatchNativeVrf    
    // 'nativeIpAddress' => csubJobMatchNativeIpAddrType,                         
    // csubJobMatchNativeIpAddr,                         
    // csubJobMatchNativeIpMask,     'domainVrf'       => csubJobMatchDomainVrf   
    // 'domainIpAddress' => csubJobMatchDomainIpAddrType,                         
    // csubJobMatchDomainIpAddr,                          csubJobMatchDomainIpMask
    // 'pbhk'            => csubJobMatchPbhk     'remoteId'        =>
    // csubJobMatchRemoteId     'circuitId'       => csubJobMatchCircuitId    
    // 'nasPort'         => csubJobMatchNasPort     'domain'          =>
    // csubJobMatchDomain     'username'        => csubJobMatchUsername    
    // 'acctSessionId'   => csubJobMatchAcctSessionId     'dnis'            =>
    // csubJobMatchDnis     'media'           => csubJobMatchMedia    
    // 'mlpNegotiated'   => csubJobMatchMlpNegotiated     'protocol'        =>
    // csubJobMatchProtocol     'serviceName'     => csubJobMatchServiceName    
    // 'dhcpClass'       => csubJobMatchDhcpClass     'tunnelName'      =>
    // csubJobMatchTunnelName  Observe that the bit 'ifIndex' has no meaning, as
    // subscriber session jobs do not match against this subscriber session
    // identity. The type is map[string]bool.
    Csubjobmatchidentities interface{}

    // This object specifies other parameters relating to subscriber sessions a
    // subscriber session job may match against.  Each bit in this bit string
    // corresponds to a column in this table.  If the bit is '0', then the value
    // of the corresponding column is invalid.  If the bit is '1', then the value
    // of the corresponding column represents the value of the parameter identity
    // the system should match against for the corresponding subscriber session
    // job.  The following list specifies the mappings between bits and the
    // columns:      'danglingDuration' => csubJobMatchDanglingDuration    
    // 'state'            => csubJobMatchState     'authenticated'    =>
    // csubJobMatchAuthenticated     'redundancyMode'   =>
    // csubJobMatchRedundancyMode. The type is map[string]bool.
    Csubjobmatchotherparams interface{}

    // This object specifies the subscriber label that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'subscriberLabel' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is interface{} with
    // range: 0..4294967295.
    Csubjobmatchsubscriberlabel interface{}

    // This object specifies the MAC address that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'macAddress' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Csubjobmatchmacaddress interface{}

    // This object specifies the native VRF the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'nativeVrf' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchnativevrf interface{}

    // This object specifies the type of Internet address specified by
    // csubJobMatchNativeIpAddr and csubJobMatchNativeIpMask.  This value is valid
    // only if the 'nativeIpAddress' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is InetAddressType.
    Csubjobmatchnativeipaddrtype interface{}

    // This object specifies the native IP address that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'nativeIpAddress' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string with length:
    // 0..255.
    Csubjobmatchnativeipaddr interface{}

    // This object specifies the mask used when matching the native IP address
    // against subscriber sessions in the process of executing a subscriber
    // session job.  This value is valid only if the 'nativeIpAddress' bit of the
    // corresponding instance of csubJobMatchIdentities is '1'. The type is string
    // with length: 0..255.
    Csubjobmatchnativeipmask interface{}

    // This object specifies the domain VRF the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'domainVrf' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchdomainvrf interface{}

    // This object specifies the type of Internet address specified by
    // csubJobMatchDomainIpAddr and csubJobMatchDomainIpMask.  This value is valid
    // only if the 'domainIpAddress' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is InetAddressType.
    Csubjobmatchdomainipaddrtype interface{}

    // This object specifies the domain IP address that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'domainIpAddress' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string with length:
    // 0..255.
    Csubjobmatchdomainipaddr interface{}

    // This object specifies the mask used when matching the domain IP address
    // against subscriber sessions in the process of executing a subscriber
    // session job.  This value is valid only if the 'domainIpAddress' bit of the
    // corresponding instance of csubJobMatchIdentities is '1'. The type is string
    // with length: 0..255.
    Csubjobmatchdomainipmask interface{}

    // This object specifies the PBHK that the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'pbhk' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchpbhk interface{}

    // This object specifies the Remote-Id the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'remoteId' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchremoteid interface{}

    // This object specifies the Circuit-Id the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'circuitId' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchcircuitid interface{}

    // This object specifies the NAS port-identifier the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'nasPort' bit of the corresponding instance
    // of csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchnasport interface{}

    // This object specifies the domain the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'domain' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchdomain interface{}

    // This object specifies the username the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'username' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchusername interface{}

    // This object specifies the accounting session identifier the system matches
    // against subscriber sessions in the process of executing a subscriber
    // session job.  This value is valid only if the 'accountingSid' bit of the
    // corresponding instance of csubJobMatchIdentities is '1'. The type is
    // interface{} with range: 0..4294967295.
    Csubjobmatchacctsessionid interface{}

    // This object specifies the DNIS number the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'dnis' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchdnis interface{}

    // This object specifies the media type the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'media' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is SubscriberMediaType.
    Csubjobmatchmedia interface{}

    // This object specifies the MLP negotiated flag the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'mlpNegotiated' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is bool.
    Csubjobmatchmlpnegotiated interface{}

    // This object specifies the protocol type the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'protocol' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is
    // SubscriberProtocolType.
    Csubjobmatchprotocol interface{}

    // This object specifies the service name the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'serviceName' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchservicename interface{}

    // This object specifies the DHCP class name the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'dhcpClass' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchdhcpclass interface{}

    // This object specifies the tunnel name the system matches against subscriber
    // session in the process of executing a subscriber session job.  This value
    // is valid only if the 'tunnelName' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    Csubjobmatchtunnelname interface{}

    // This object specifies the minimum interval of time a subscriber session can
    // remain dangling in order for the system to consider it a match in the
    // process of executing a subscriber session job. A 'dangling' subscriber
    // session is one in the PENDING state.  The value '0' cannot be written to an
    // instance of this object. However, it serves as a convenient value when the
    // column is not valid.  This value is valid only if the 'danglingDuration'
    // bit of the corresponding instance of csubJobMatchOtherParams is '1'. The
    // type is interface{} with range: 0..3600.
    Csubjobmatchdanglingduration interface{}

    // This object specifies the state of a subscriber session in order for the
    // system to consider a match in the process of executing a subscriber session
    // job.  The value 'other' is not valid and an implementation should not allow
    // it to be written to this column.  This value is valid only if the 'state'
    // bit of the corresponding instance of csubJobMatchOtherParams is '1'. The
    // type is SubSessionState.
    Csubjobmatchstate interface{}

    // This object specifies whether a subscriber session should be
    // unauthenticated for the system to consider a match in the process of
    // executing a subscriber session job.  If this column is 'false', then the
    // subscriber session job matches subscriber sessions that are
    // unauthenticated.  If this column is 'true', then the subscriber session job
    // matches subscriber session that are authenticated.  This value is valid
    // only if the 'authenticated' bit of the corresponding instance of
    // csubJobMatchParams is '1'. The type is bool.
    Csubjobmatchauthenticated interface{}

    // This object specifies the redudancy mode of the subscriber session in order
    // for the system to consider a match in the process of executing a subscriber
    // session job.  The value 'other' is not valid and an implementation should
    // not allow it to be written to this column.  This value is valid only if the
    // 'redundancyMode' bit of the corresponding instance of
    // csubJobMatchOtherParams is '1'. The type is SubSessionRedundancyMode.
    Csubjobmatchredundancymode interface{}
}

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetFilter() yfilter.YFilter { return csubjobmatchparamsentry.YFilter }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) SetFilter(yf yfilter.YFilter) { csubjobmatchparamsentry.YFilter = yf }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetGoName(yname string) string {
    if yname == "csubJobId" { return "Csubjobid" }
    if yname == "csubJobMatchIdentities" { return "Csubjobmatchidentities" }
    if yname == "csubJobMatchOtherParams" { return "Csubjobmatchotherparams" }
    if yname == "csubJobMatchSubscriberLabel" { return "Csubjobmatchsubscriberlabel" }
    if yname == "csubJobMatchMacAddress" { return "Csubjobmatchmacaddress" }
    if yname == "csubJobMatchNativeVrf" { return "Csubjobmatchnativevrf" }
    if yname == "csubJobMatchNativeIpAddrType" { return "Csubjobmatchnativeipaddrtype" }
    if yname == "csubJobMatchNativeIpAddr" { return "Csubjobmatchnativeipaddr" }
    if yname == "csubJobMatchNativeIpMask" { return "Csubjobmatchnativeipmask" }
    if yname == "csubJobMatchDomainVrf" { return "Csubjobmatchdomainvrf" }
    if yname == "csubJobMatchDomainIpAddrType" { return "Csubjobmatchdomainipaddrtype" }
    if yname == "csubJobMatchDomainIpAddr" { return "Csubjobmatchdomainipaddr" }
    if yname == "csubJobMatchDomainIpMask" { return "Csubjobmatchdomainipmask" }
    if yname == "csubJobMatchPbhk" { return "Csubjobmatchpbhk" }
    if yname == "csubJobMatchRemoteId" { return "Csubjobmatchremoteid" }
    if yname == "csubJobMatchCircuitId" { return "Csubjobmatchcircuitid" }
    if yname == "csubJobMatchNasPort" { return "Csubjobmatchnasport" }
    if yname == "csubJobMatchDomain" { return "Csubjobmatchdomain" }
    if yname == "csubJobMatchUsername" { return "Csubjobmatchusername" }
    if yname == "csubJobMatchAcctSessionId" { return "Csubjobmatchacctsessionid" }
    if yname == "csubJobMatchDnis" { return "Csubjobmatchdnis" }
    if yname == "csubJobMatchMedia" { return "Csubjobmatchmedia" }
    if yname == "csubJobMatchMlpNegotiated" { return "Csubjobmatchmlpnegotiated" }
    if yname == "csubJobMatchProtocol" { return "Csubjobmatchprotocol" }
    if yname == "csubJobMatchServiceName" { return "Csubjobmatchservicename" }
    if yname == "csubJobMatchDhcpClass" { return "Csubjobmatchdhcpclass" }
    if yname == "csubJobMatchTunnelName" { return "Csubjobmatchtunnelname" }
    if yname == "csubJobMatchDanglingDuration" { return "Csubjobmatchdanglingduration" }
    if yname == "csubJobMatchState" { return "Csubjobmatchstate" }
    if yname == "csubJobMatchAuthenticated" { return "Csubjobmatchauthenticated" }
    if yname == "csubJobMatchRedundancyMode" { return "Csubjobmatchredundancymode" }
    return ""
}

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetSegmentPath() string {
    return "csubJobMatchParamsEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobmatchparamsentry.Csubjobid) + "']"
}

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubJobId"] = csubjobmatchparamsentry.Csubjobid
    leafs["csubJobMatchIdentities"] = csubjobmatchparamsentry.Csubjobmatchidentities
    leafs["csubJobMatchOtherParams"] = csubjobmatchparamsentry.Csubjobmatchotherparams
    leafs["csubJobMatchSubscriberLabel"] = csubjobmatchparamsentry.Csubjobmatchsubscriberlabel
    leafs["csubJobMatchMacAddress"] = csubjobmatchparamsentry.Csubjobmatchmacaddress
    leafs["csubJobMatchNativeVrf"] = csubjobmatchparamsentry.Csubjobmatchnativevrf
    leafs["csubJobMatchNativeIpAddrType"] = csubjobmatchparamsentry.Csubjobmatchnativeipaddrtype
    leafs["csubJobMatchNativeIpAddr"] = csubjobmatchparamsentry.Csubjobmatchnativeipaddr
    leafs["csubJobMatchNativeIpMask"] = csubjobmatchparamsentry.Csubjobmatchnativeipmask
    leafs["csubJobMatchDomainVrf"] = csubjobmatchparamsentry.Csubjobmatchdomainvrf
    leafs["csubJobMatchDomainIpAddrType"] = csubjobmatchparamsentry.Csubjobmatchdomainipaddrtype
    leafs["csubJobMatchDomainIpAddr"] = csubjobmatchparamsentry.Csubjobmatchdomainipaddr
    leafs["csubJobMatchDomainIpMask"] = csubjobmatchparamsentry.Csubjobmatchdomainipmask
    leafs["csubJobMatchPbhk"] = csubjobmatchparamsentry.Csubjobmatchpbhk
    leafs["csubJobMatchRemoteId"] = csubjobmatchparamsentry.Csubjobmatchremoteid
    leafs["csubJobMatchCircuitId"] = csubjobmatchparamsentry.Csubjobmatchcircuitid
    leafs["csubJobMatchNasPort"] = csubjobmatchparamsentry.Csubjobmatchnasport
    leafs["csubJobMatchDomain"] = csubjobmatchparamsentry.Csubjobmatchdomain
    leafs["csubJobMatchUsername"] = csubjobmatchparamsentry.Csubjobmatchusername
    leafs["csubJobMatchAcctSessionId"] = csubjobmatchparamsentry.Csubjobmatchacctsessionid
    leafs["csubJobMatchDnis"] = csubjobmatchparamsentry.Csubjobmatchdnis
    leafs["csubJobMatchMedia"] = csubjobmatchparamsentry.Csubjobmatchmedia
    leafs["csubJobMatchMlpNegotiated"] = csubjobmatchparamsentry.Csubjobmatchmlpnegotiated
    leafs["csubJobMatchProtocol"] = csubjobmatchparamsentry.Csubjobmatchprotocol
    leafs["csubJobMatchServiceName"] = csubjobmatchparamsentry.Csubjobmatchservicename
    leafs["csubJobMatchDhcpClass"] = csubjobmatchparamsentry.Csubjobmatchdhcpclass
    leafs["csubJobMatchTunnelName"] = csubjobmatchparamsentry.Csubjobmatchtunnelname
    leafs["csubJobMatchDanglingDuration"] = csubjobmatchparamsentry.Csubjobmatchdanglingduration
    leafs["csubJobMatchState"] = csubjobmatchparamsentry.Csubjobmatchstate
    leafs["csubJobMatchAuthenticated"] = csubjobmatchparamsentry.Csubjobmatchauthenticated
    leafs["csubJobMatchRedundancyMode"] = csubjobmatchparamsentry.Csubjobmatchredundancymode
    return leafs
}

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetYangName() string { return "csubJobMatchParamsEntry" }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) SetParent(parent types.Entity) { csubjobmatchparamsentry.parent = parent }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetParent() types.Entity { return csubjobmatchparamsentry.parent }

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetParentYangName() string { return "csubJobMatchParamsTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable
// This table contains subscriber session job parameters
// describing query parameters.
// 
// This table has a sparse-dependent relationship on the
// csubJobTable, containing a row for each job having a
// csubJobType of 'query'.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describes a set of subscriber session query parameters.  The
    // system automatically creates an entry when the EMS/NMS sets the
    // corresponding instance of csubJobType to 'query'.  Likewise, the system
    // automatically destroys an entry under the following circumstances:  1)  The
    // EMS/NMS destroys the corresponding row in the csubJobTable.  2)  The
    // EMS/NMS sets the corresponding instance of csubJobType to     'noop' or
    // 'clear'. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry.
    Csubjobqueryparamsentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry
}

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetFilter() yfilter.YFilter { return csubjobqueryparamstable.YFilter }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) SetFilter(yf yfilter.YFilter) { csubjobqueryparamstable.YFilter = yf }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetGoName(yname string) string {
    if yname == "csubJobQueryParamsEntry" { return "Csubjobqueryparamsentry" }
    return ""
}

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetSegmentPath() string {
    return "csubJobQueryParamsTable"
}

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubJobQueryParamsEntry" {
        for _, c := range csubjobqueryparamstable.Csubjobqueryparamsentry {
            if csubjobqueryparamstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry{}
        csubjobqueryparamstable.Csubjobqueryparamsentry = append(csubjobqueryparamstable.Csubjobqueryparamsentry, child)
        return &csubjobqueryparamstable.Csubjobqueryparamsentry[len(csubjobqueryparamstable.Csubjobqueryparamsentry)-1]
    }
    return nil
}

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubjobqueryparamstable.Csubjobqueryparamsentry {
        children[csubjobqueryparamstable.Csubjobqueryparamsentry[i].GetSegmentPath()] = &csubjobqueryparamstable.Csubjobqueryparamsentry[i]
    }
    return children
}

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetYangName() string { return "csubJobQueryParamsTable" }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) SetParent(parent types.Entity) { csubjobqueryparamstable.parent = parent }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetParent() types.Entity { return csubjobqueryparamstable.parent }

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry
// An entry describes a set of subscriber session query
// parameters.
// 
// The system automatically creates an entry when the EMS/NMS sets
// the corresponding instance of csubJobType to 'query'.  Likewise,
// the system automatically destroys an entry under the following
// circumstances:
// 
// 1)  The EMS/NMS destroys the corresponding row in the csubJobTable.
// 
// 2)  The EMS/NMS sets the corresponding instance of csubJobType to
//     'noop' or 'clear'.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobid
    Csubjobid interface{}

    // This object specifies the first subscriber identity that the system uses
    // when sorting subscriber sessions into the final report corresponding to a
    // subscriber session query.  It is not valid to set this column to 'other' or
    // 'ifIndex'. The type is SubSessionIdentity.
    Csubjobquerysortkey1 interface{}

    // This object specifies the second subscriber identity that the system uses
    // when sorting subscriber sessions into the final report corresponding to a
    // subscriber session query.  If it is the desire to have the final report
    // sorted on a single subscriber identity, then this column should be 'other'.
    // The type is SubSessionIdentity.
    Csubjobquerysortkey2 interface{}

    // This object specifies the third subscriber identity that the system uses
    // when sorting subscriber sessions into the final report corresponding to a
    // subscriber session query.  If it is the desire to have the final report
    // sorted on one or two subscriber identities, then this column should be
    // 'other'. The type is SubSessionIdentity.
    Csubjobquerysortkey3 interface{}

    // This object indicates the number of subscriber sessions that matched the
    // corresponding subscriber session query.  The value of this column should be
    // '0' unless the corresponding value of csubJobState is 'finished'.  The
    // value of this column should be '0' after the EMS/NMS sets the corresponding
    // csubJobControl to 'release'. The type is interface{} with range:
    // 0..4294967295.
    Csubjobqueryresultingreportsize interface{}
}

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetFilter() yfilter.YFilter { return csubjobqueryparamsentry.YFilter }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) SetFilter(yf yfilter.YFilter) { csubjobqueryparamsentry.YFilter = yf }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetGoName(yname string) string {
    if yname == "csubJobId" { return "Csubjobid" }
    if yname == "csubJobQuerySortKey1" { return "Csubjobquerysortkey1" }
    if yname == "csubJobQuerySortKey2" { return "Csubjobquerysortkey2" }
    if yname == "csubJobQuerySortKey3" { return "Csubjobquerysortkey3" }
    if yname == "csubJobQueryResultingReportSize" { return "Csubjobqueryresultingreportsize" }
    return ""
}

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetSegmentPath() string {
    return "csubJobQueryParamsEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobqueryparamsentry.Csubjobid) + "']"
}

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubJobId"] = csubjobqueryparamsentry.Csubjobid
    leafs["csubJobQuerySortKey1"] = csubjobqueryparamsentry.Csubjobquerysortkey1
    leafs["csubJobQuerySortKey2"] = csubjobqueryparamsentry.Csubjobquerysortkey2
    leafs["csubJobQuerySortKey3"] = csubjobqueryparamsentry.Csubjobquerysortkey3
    leafs["csubJobQueryResultingReportSize"] = csubjobqueryparamsentry.Csubjobqueryresultingreportsize
    return leafs
}

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetYangName() string { return "csubJobQueryParamsEntry" }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) SetParent(parent types.Entity) { csubjobqueryparamsentry.parent = parent }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetParent() types.Entity { return csubjobqueryparamsentry.parent }

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetParentYangName() string { return "csubJobQueryParamsTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable
// This table lists the subscriber session jobs currently pending
// in the subscriber session job queue.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describing an subscriber session job in the subscriber session job
    // queue.  The system creates an entry in this table when it places a
    // subscriber session job on the subscriber session job queue.  It does this
    // when the EMS/NMS sets an instance of csubJobControl to 'start' and the
    // system is already executing a subscriber session job.  Likewise, the system
    // destroys an entry when it removes it from the queue.  This occurs under the
    // following circumstances:  1)  The system has finished executing a job, for
    // whatever     reason, and is ready to start executing the job at the head   
    // of the queue.  2)  The EMS/NMS has set an instance of csubJobControl to
    // 'abort'     for a job that was on the queue. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry.
    Csubjobqueueentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry
}

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetFilter() yfilter.YFilter { return csubjobqueuetable.YFilter }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) SetFilter(yf yfilter.YFilter) { csubjobqueuetable.YFilter = yf }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetGoName(yname string) string {
    if yname == "csubJobQueueEntry" { return "Csubjobqueueentry" }
    return ""
}

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetSegmentPath() string {
    return "csubJobQueueTable"
}

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubJobQueueEntry" {
        for _, c := range csubjobqueuetable.Csubjobqueueentry {
            if csubjobqueuetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry{}
        csubjobqueuetable.Csubjobqueueentry = append(csubjobqueuetable.Csubjobqueueentry, child)
        return &csubjobqueuetable.Csubjobqueueentry[len(csubjobqueuetable.Csubjobqueueentry)-1]
    }
    return nil
}

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubjobqueuetable.Csubjobqueueentry {
        children[csubjobqueuetable.Csubjobqueueentry[i].GetSegmentPath()] = &csubjobqueuetable.Csubjobqueueentry[i]
    }
    return children
}

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetYangName() string { return "csubJobQueueTable" }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) SetParent(parent types.Entity) { csubjobqueuetable.parent = parent }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetParent() types.Entity { return csubjobqueuetable.parent }

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry
// An entry describing an subscriber session job in the
// subscriber session job queue.
// 
// The system creates an entry in this table when it places a
// subscriber session job on the subscriber session job queue.  It
// does this when the EMS/NMS sets an instance of csubJobControl to
// 'start' and the system is already executing a subscriber session
// job.  Likewise, the system destroys an entry when it removes it
// from the queue.  This occurs under the following circumstances:
// 
// 1)  The system has finished executing a job, for whatever
//     reason, and is ready to start executing the job at the head
//     of the queue.
// 
// 2)  The EMS/NMS has set an instance of csubJobControl to 'abort'
//     for a job that was on the queue.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates an positive, integer-value
    // that uniquely identifies the entry in the table. The value of this object
    // starts at '1' and monotonically increases for each subscriber session job
    // inserted into the subscriber session job queue.  If the value of this
    // object is '4294967295', the system will reset it to '1' when it inserts the
    // next subscriber session job into the subscriber session job queue. The type
    // is interface{} with range: 1..4294967295.
    Csubjobqueuenumber interface{}

    // This object indicates the identifier associated with the subscriber session
    // job. The type is interface{} with range: 1..4294967295.
    Csubjobqueuejobid interface{}
}

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetFilter() yfilter.YFilter { return csubjobqueueentry.YFilter }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) SetFilter(yf yfilter.YFilter) { csubjobqueueentry.YFilter = yf }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetGoName(yname string) string {
    if yname == "csubJobQueueNumber" { return "Csubjobqueuenumber" }
    if yname == "csubJobQueueJobId" { return "Csubjobqueuejobid" }
    return ""
}

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetSegmentPath() string {
    return "csubJobQueueEntry" + "[csubJobQueueNumber='" + fmt.Sprintf("%v", csubjobqueueentry.Csubjobqueuenumber) + "']"
}

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubJobQueueNumber"] = csubjobqueueentry.Csubjobqueuenumber
    leafs["csubJobQueueJobId"] = csubjobqueueentry.Csubjobqueuejobid
    return leafs
}

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetYangName() string { return "csubJobQueueEntry" }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) SetParent(parent types.Entity) { csubjobqueueentry.parent = parent }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetParent() types.Entity { return csubjobqueueentry.parent }

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetParentYangName() string { return "csubJobQueueTable" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable
// This table contains the reports corresponding to subscriber
// session jobs that have a csubJobType of 'query' and
// csubJobState of 'finished'.
// 
// This table has an expansion dependent relationship on the
// csubJobTable, containing zero or more rows for each job.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describes a subscriber session that satisfied the match criteria
    // specified by the corresponding job.  The system creates an entry for each
    // subscriber session that satisfied the specified match criteria of a
    // subscriber session job having a csubJobType of 'query'.  However, it does
    // not create these entries until after the system has successfully executed
    // the subscriber session job.  The system destroys an entry under the
    // following circumstances:  1)  The corresponding subscriber session job has
    // been destroyed     by the EMS/NMS.  2)  The value of csubJobMaxLife is
    // non-zero and the age of the     report has reached the specified maximum
    // life.  3)  The EMS/NMS has set the corresponding instance of    
    // csubJobControl to 'release'.  4)  The EMS/NMS has restarted the
    // corresponding subscriber     session job (i.e., has set the corresponding
    // instance of     csubJobControl to 'start'). The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry.
    Csubjobreportentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry
}

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetFilter() yfilter.YFilter { return csubjobreporttable.YFilter }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) SetFilter(yf yfilter.YFilter) { csubjobreporttable.YFilter = yf }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetGoName(yname string) string {
    if yname == "csubJobReportEntry" { return "Csubjobreportentry" }
    return ""
}

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetSegmentPath() string {
    return "csubJobReportTable"
}

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csubJobReportEntry" {
        for _, c := range csubjobreporttable.Csubjobreportentry {
            if csubjobreporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry{}
        csubjobreporttable.Csubjobreportentry = append(csubjobreporttable.Csubjobreportentry, child)
        return &csubjobreporttable.Csubjobreportentry[len(csubjobreporttable.Csubjobreportentry)-1]
    }
    return nil
}

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csubjobreporttable.Csubjobreportentry {
        children[csubjobreporttable.Csubjobreportentry[i].GetSegmentPath()] = &csubjobreporttable.Csubjobreportentry[i]
    }
    return children
}

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetYangName() string { return "csubJobReportTable" }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) SetParent(parent types.Entity) { csubjobreporttable.parent = parent }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetParent() types.Entity { return csubjobreporttable.parent }

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetParentYangName() string { return "CISCO-SUBSCRIBER-SESSION-MIB" }

// CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry
// An entry describes a subscriber session that satisfied the
// match criteria specified by the corresponding job.
// 
// The system creates an entry for each subscriber session that
// satisfied the specified match criteria of a subscriber session
// job having a csubJobType of 'query'.  However, it does not
// create these entries until after the system has successfully
// executed the subscriber session job.
// 
// The system destroys an entry under the following circumstances:
// 
// 1)  The corresponding subscriber session job has been destroyed
//     by the EMS/NMS.
// 
// 2)  The value of csubJobMaxLife is non-zero and the age of the
//     report has reached the specified maximum life.
// 
// 3)  The EMS/NMS has set the corresponding instance of
//     csubJobControl to 'release'.
// 
// 4)  The EMS/NMS has restarted the corresponding subscriber
//     session job (i.e., has set the corresponding instance of
//     csubJobControl to 'start').
type CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry_Csubjobid
    Csubjobid interface{}

    // This attribute is a key. This object indicates an arbitrary, positive,
    // integer-value that uniquely identifies this entry in a report.  This
    // auxiliary value is necessary, as the corresponding subscriber session job
    // can specify up to three subscriber identities on which to sort the
    // subscriber sessions that end up in the final report. The type is
    // interface{} with range: 1..4294967295.
    Csubjobreportid interface{}

    // This object indicates the ifIndex-value assigned to the subscriber session
    // that satisfied the match criteria specified by the corresponding subscriber
    // session job having a csubJobType of 'query'. The type is interface{} with
    // range: 1..2147483647.
    Csubjobreportsession interface{}
}

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetFilter() yfilter.YFilter { return csubjobreportentry.YFilter }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) SetFilter(yf yfilter.YFilter) { csubjobreportentry.YFilter = yf }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetGoName(yname string) string {
    if yname == "csubJobId" { return "Csubjobid" }
    if yname == "csubJobReportId" { return "Csubjobreportid" }
    if yname == "csubJobReportSession" { return "Csubjobreportsession" }
    return ""
}

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetSegmentPath() string {
    return "csubJobReportEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobreportentry.Csubjobid) + "']" + "[csubJobReportId='" + fmt.Sprintf("%v", csubjobreportentry.Csubjobreportid) + "']"
}

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["csubJobId"] = csubjobreportentry.Csubjobid
    leafs["csubJobReportId"] = csubjobreportentry.Csubjobreportid
    leafs["csubJobReportSession"] = csubjobreportentry.Csubjobreportsession
    return leafs
}

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetBundleName() string { return "cisco_ios_xe" }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetYangName() string { return "csubJobReportEntry" }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) SetParent(parent types.Entity) { csubjobreportentry.parent = parent }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetParent() types.Entity { return csubjobreportentry.parent }

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetParentYangName() string { return "csubJobReportTable" }

