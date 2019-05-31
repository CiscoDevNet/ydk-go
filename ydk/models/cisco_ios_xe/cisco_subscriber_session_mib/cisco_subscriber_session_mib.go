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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CsubJob CISCOSUBSCRIBERSESSIONMIB_CsubJob

    
    CsubAggThresh CISCOSUBSCRIBERSESSIONMIB_CsubAggThresh

    // This table describes a list of subscriber sessions currently maintained by
    // the system.  This table has a sparse dependent relationship on the ifTable,
    // containing a row for each interface having an interface type describing a
    // subscriber session.
    CsubSessionTable CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable

    // This table describes a list of subscriber sessions currently maintained by
    // the system.  The tables sorts the subscriber sessions first by the
    // subscriber session's type and second by the ifIndex assigned to the
    // subscriber session.
    CsubSessionByTypeTable CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable

    // This table contains sets of aggregated statistics relating to subscriber
    // sessions, where each set has a unique scope of aggregation.
    CsubAggStatsTable CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable

    // This table contains aggregated subscriber session performance data
    // collected for as much as a day's worth of 15-minute measurement intervals. 
    // This table has an expansion dependent relationship on the
    // csubAggStatsTable, containing zero or more rows for each row contained by
    // the csubAggStatsTable.  Observe that the collection and maintenance of
    // aggregated subscriber performance data is OPTIONAL for all scopes of
    // aggregation.  However, an implementation should maintain at least one
    // interval for the 'scope of aggregation' that contains all subscriber
    // sessions maintained by the system.
    CsubAggStatsIntTable CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable

    // Please enter the Table Description here.
    CsubAggStatsThreshTable CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable

    // This table contains the subscriber session jobs submitted by the EMS/NMS.
    CsubJobTable CISCOSUBSCRIBERSESSIONMIB_CsubJobTable

    // This table contains subscriber session job parameters describing match
    // criteria.  This table has a sparse-dependent relationship on the
    // csubJobTable, containing a row for each job having a csubJobType of 'query'
    // or 'clear'.
    CsubJobMatchParamsTable CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable

    // This table contains subscriber session job parameters describing query
    // parameters.  This table has a sparse-dependent relationship on the
    // csubJobTable, containing a row for each job having a csubJobType of
    // 'query'.
    CsubJobQueryParamsTable CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable

    // This table lists the subscriber session jobs currently pending in the
    // subscriber session job queue.
    CsubJobQueueTable CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable

    // This table contains the reports corresponding to subscriber session jobs
    // that have a csubJobType of 'query' and csubJobState of 'finished'.  This
    // table has an expansion dependent relationship on the csubJobTable,
    // containing zero or more rows for each job.
    CsubJobReportTable CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable
}

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetEntityData() *types.CommonEntityData {
    cISCOSUBSCRIBERSESSIONMIB.EntityData.YFilter = cISCOSUBSCRIBERSESSIONMIB.YFilter
    cISCOSUBSCRIBERSESSIONMIB.EntityData.YangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.SegmentPath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.AbsolutePath = cISCOSUBSCRIBERSESSIONMIB.EntityData.SegmentPath
    cISCOSUBSCRIBERSESSIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSUBSCRIBERSESSIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSUBSCRIBERSESSIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children = types.NewOrderedMap()
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubJob", types.YChild{"CsubJob", &cISCOSUBSCRIBERSESSIONMIB.CsubJob})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubAggThresh", types.YChild{"CsubAggThresh", &cISCOSUBSCRIBERSESSIONMIB.CsubAggThresh})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubSessionTable", types.YChild{"CsubSessionTable", &cISCOSUBSCRIBERSESSIONMIB.CsubSessionTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubSessionByTypeTable", types.YChild{"CsubSessionByTypeTable", &cISCOSUBSCRIBERSESSIONMIB.CsubSessionByTypeTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubAggStatsTable", types.YChild{"CsubAggStatsTable", &cISCOSUBSCRIBERSESSIONMIB.CsubAggStatsTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubAggStatsIntTable", types.YChild{"CsubAggStatsIntTable", &cISCOSUBSCRIBERSESSIONMIB.CsubAggStatsIntTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubAggStatsThreshTable", types.YChild{"CsubAggStatsThreshTable", &cISCOSUBSCRIBERSESSIONMIB.CsubAggStatsThreshTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubJobTable", types.YChild{"CsubJobTable", &cISCOSUBSCRIBERSESSIONMIB.CsubJobTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubJobMatchParamsTable", types.YChild{"CsubJobMatchParamsTable", &cISCOSUBSCRIBERSESSIONMIB.CsubJobMatchParamsTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubJobQueryParamsTable", types.YChild{"CsubJobQueryParamsTable", &cISCOSUBSCRIBERSESSIONMIB.CsubJobQueryParamsTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubJobQueueTable", types.YChild{"CsubJobQueueTable", &cISCOSUBSCRIBERSESSIONMIB.CsubJobQueueTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children.Append("csubJobReportTable", types.YChild{"CsubJobReportTable", &cISCOSUBSCRIBERSESSIONMIB.CsubJobReportTable})
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOSUBSCRIBERSESSIONMIB.EntityData.YListKeys = []string {}

    return &(cISCOSUBSCRIBERSESSIONMIB.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJob
type CISCOSUBSCRIBERSESSIONMIB_CsubJob struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies whether the system generates a csubJobFinishedNotify
    // notification when the system finishes processing a subscriber session job.
    // The type is bool.
    CsubJobFinishedNotifyEnable interface{}

    // This object indicates which subscriber session identities the system
    // maintains as searchable keys.  This value serves the EMS/NMS in configuring
    // a subscriber session query, as at least one match criteria must be an
    // 'indexed attribute'. The type is map[string]bool.
    CsubJobIndexedAttributes interface{}

    // This object indicates the next available identifier for the creation of a
    // new row in the csubJobTable.  If no available identifier exists, then this
    // object has the value '0'. The type is interface{} with range:
    // 0..4294967295.
    CsubJobIdNext interface{}

    // This object indicates the maximum number of outstanding subscriber session
    // jobs the system can support.  If the csubJobTable contains a number of rows
    // (i.e., the value of csubJobCount) equal to this value, then any attempt to
    // create a new row will result in a response with an error-status of
    // 'resourceUnavailable'. The type is interface{} with range: 1..4294967295.
    // Units are jobs.
    CsubJobMaxNumber interface{}

    // This object specifies the maximum life a subscriber session report
    // corresponding to a subscriber session job having a csubJobType of 'query'. 
    // The system tracks the time elapsed after it finishes processing a query. 
    // When the time elapsed reaches the value specified by this column, the
    // system automatically  destroys the report.  A value of '0' disables the
    // automatic destruction of reports. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    CsubJobMaxLife interface{}

    // This object indicates the number of subscriber session jobs currently
    // maintained by the csubJobTable. The type is interface{} with range:
    // 0..4294967295. Units are jobs.
    CsubJobCount interface{}
}

func (csubJob *CISCOSUBSCRIBERSESSIONMIB_CsubJob) GetEntityData() *types.CommonEntityData {
    csubJob.EntityData.YFilter = csubJob.YFilter
    csubJob.EntityData.YangName = "csubJob"
    csubJob.EntityData.BundleName = "cisco_ios_xe"
    csubJob.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubJob.EntityData.SegmentPath = "csubJob"
    csubJob.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubJob.EntityData.SegmentPath
    csubJob.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJob.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJob.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJob.EntityData.Children = types.NewOrderedMap()
    csubJob.EntityData.Leafs = types.NewOrderedMap()
    csubJob.EntityData.Leafs.Append("csubJobFinishedNotifyEnable", types.YLeaf{"CsubJobFinishedNotifyEnable", csubJob.CsubJobFinishedNotifyEnable})
    csubJob.EntityData.Leafs.Append("csubJobIndexedAttributes", types.YLeaf{"CsubJobIndexedAttributes", csubJob.CsubJobIndexedAttributes})
    csubJob.EntityData.Leafs.Append("csubJobIdNext", types.YLeaf{"CsubJobIdNext", csubJob.CsubJobIdNext})
    csubJob.EntityData.Leafs.Append("csubJobMaxNumber", types.YLeaf{"CsubJobMaxNumber", csubJob.CsubJobMaxNumber})
    csubJob.EntityData.Leafs.Append("csubJobMaxLife", types.YLeaf{"CsubJobMaxLife", csubJob.CsubJobMaxLife})
    csubJob.EntityData.Leafs.Append("csubJobCount", types.YLeaf{"CsubJobCount", csubJob.CsubJobCount})

    csubJob.EntityData.YListKeys = []string {}

    return &(csubJob.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubAggThresh
type CISCOSUBSCRIBERSESSIONMIB_CsubAggThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object enables or disables the generation of any of the csubAggStats*
    // threshold notifications. The type is bool.
    CsubAggStatsThreshNotifEnable interface{}
}

func (csubAggThresh *CISCOSUBSCRIBERSESSIONMIB_CsubAggThresh) GetEntityData() *types.CommonEntityData {
    csubAggThresh.EntityData.YFilter = csubAggThresh.YFilter
    csubAggThresh.EntityData.YangName = "csubAggThresh"
    csubAggThresh.EntityData.BundleName = "cisco_ios_xe"
    csubAggThresh.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubAggThresh.EntityData.SegmentPath = "csubAggThresh"
    csubAggThresh.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubAggThresh.EntityData.SegmentPath
    csubAggThresh.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubAggThresh.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubAggThresh.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubAggThresh.EntityData.Children = types.NewOrderedMap()
    csubAggThresh.EntityData.Leafs = types.NewOrderedMap()
    csubAggThresh.EntityData.Leafs.Append("csubAggStatsThreshNotifEnable", types.YLeaf{"CsubAggStatsThreshNotifEnable", csubAggThresh.CsubAggStatsThreshNotifEnable})

    csubAggThresh.EntityData.YListKeys = []string {}

    return &(csubAggThresh.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable
// This table describes a list of subscriber sessions currently
// maintained by the system.
// 
// This table has a sparse dependent relationship on the ifTable,
// containing a row for each interface having an interface type
// describing a subscriber session.
type CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry contains data describing a subscriber sessions, including its
    // state, configuration, and collected identities.  An entry exists for a
    // corresponding entry in the ifTable describing a subscriber session. 
    // Currently, subscriber sessions must have one of the following ifType
    // values:      'ppp'     'ipSubscriberSession'     'l2SubscriberSession'  The
    // system creates an entry when it establishes a subscriber session. 
    // Likewise, the system destroys an entry when it terminates the corresponding
    // subscriber session. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry.
    CsubSessionEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry
}

func (csubSessionTable *CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable) GetEntityData() *types.CommonEntityData {
    csubSessionTable.EntityData.YFilter = csubSessionTable.YFilter
    csubSessionTable.EntityData.YangName = "csubSessionTable"
    csubSessionTable.EntityData.BundleName = "cisco_ios_xe"
    csubSessionTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubSessionTable.EntityData.SegmentPath = "csubSessionTable"
    csubSessionTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubSessionTable.EntityData.SegmentPath
    csubSessionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubSessionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubSessionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubSessionTable.EntityData.Children = types.NewOrderedMap()
    csubSessionTable.EntityData.Children.Append("csubSessionEntry", types.YChild{"CsubSessionEntry", nil})
    for i := range csubSessionTable.CsubSessionEntry {
        csubSessionTable.EntityData.Children.Append(types.GetSegmentPath(csubSessionTable.CsubSessionEntry[i]), types.YChild{"CsubSessionEntry", csubSessionTable.CsubSessionEntry[i]})
    }
    csubSessionTable.EntityData.Leafs = types.NewOrderedMap()

    csubSessionTable.EntityData.YListKeys = []string {}

    return &(csubSessionTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object indicates the type of subscriber session. The type is
    // SubSessionType.
    CsubSessionType interface{}

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
    // CsubSessionIpAddrAssignment.
    CsubSessionIpAddrAssignment interface{}

    // This object indicates the operational state of the subscriber session. The
    // type is SubSessionState.
    CsubSessionState interface{}

    // This object indicates whether the system has successfully authenticated the
    // subscriber session.      'false'         The subscriber session is
    // operationally up, but in the         UNAUTHENTICATED state.      'true'    
    // The subscriber session is operationally up, but in the        
    // AUTHENTICATED state. The type is bool.
    CsubSessionAuthenticated interface{}

    // This object indicates the redundancy mode in which the subscriber session
    // is operating. The type is SubSessionRedundancyMode.
    CsubSessionRedundancyMode interface{}

    // This object indicates when the subscriber session was created (i.e., when
    // the subscriber session was initiated). The type is string.
    CsubSessionCreationTime interface{}

    // This object indicates the row in the cdtTemplateTable (defined by the
    // CISCO-DYNAMIC-TEMPLATE-MIB) describing the derived configuration for the
    // subscriber session.  Observe that the value of cdtTemplateType
    // corresponding to the referenced row must be 'derived'. The type is string
    // with length: 1..64.
    CsubSessionDerivedCfg interface{}

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
    CsubSessionAvailableIdentities interface{}

    // This object indicates a positive integer-value uniquely identifying the
    // subscriber session within the scope of the system.  This column is valid
    // only if the 'subscriberLabel' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is interface{} with range:
    // 0..4294967295.
    CsubSessionSubscriberLabel interface{}

    // This object indicates the MAC address of the subscriber.  This column is
    // valid only if the 'macAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    CsubSessionMacAddress interface{}

    // This object indicates the VRF originating the subscriber session.  This
    // column is valid only if the 'nativeVrf' bit of the corresponding instance
    // of csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionNativeVrf interface{}

    // This object indicates the type of IP address assigned to the subscriber on
    // the customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is InetAddressType.
    CsubSessionNativeIpAddrType interface{}

    // This object indicates the IP address assigned to the subscriber on the
    // customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    CsubSessionNativeIpAddr interface{}

    // This object indicates the corresponding mask for the IP address assigned to
    // the subscriber on the customer-facing side of the system.  This column is
    // valid only if the 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    CsubSessionNativeIpMask interface{}

    // This object indicates the VRF to which the system transfers the subscriber
    // session's traffic.  This column is valid only if the 'domainVrf' bit of the
    // corresponding instance of csubSessionAvailableIdentities is '1'. The type
    // is string.
    CsubSessionDomainVrf interface{}

    // This object indicates the type of IP address assigned to the subscriber on
    // the service-facing side of the system.  This column is valid only if the
    // 'domainIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is InetAddressType.
    CsubSessionDomainIpAddrType interface{}

    // This object indicates the IP address assigned to the subscriber on the
    // service-facing side of the system.  This column is valid only if the
    // 'domainIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    CsubSessionDomainIpAddr interface{}

    // This object indicates the corresponding mask for the IP address assigned to
    // the subscriber on the service-facing side of the system.  This column is
    // valid only if the 'domainIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    CsubSessionDomainIpMask interface{}

    // This object indicates the PBHK identifying the subscriber.  This column is
    // valid only if the 'pbhk' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionPbhk interface{}

    // This object indicates the Remote-Id identifying the 'calling station' or AN
    // supporting the circuit that provides access to the subscriber.  This column
    // is valid only if the 'remoteId' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionRemoteId interface{}

    // This object indicates the Circuit-Id identifying the circuit supported by
    // the 'calling station' or AN providing access to the subscriber.  This
    // column is valid only if the 'circuitId' bit of the corresponding instance
    // of csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionCircuitId interface{}

    // This object indicates the NAS port-identifier identifying the port on the
    // NAS providing access to the subscriber.  This column is valid only if the
    // 'nasPort' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionNasPort interface{}

    // This object indicates the domain associated with the subscriber.  This
    // column is valid only if the 'domain' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionDomain interface{}

    // This object indicates the username identifying the subscriber.  This column
    // is valid only if the 'username' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionUsername interface{}

    // This object indicates the accounting session identifier assigned to the
    // subscriber session.  This column is valid only if the 'accountingSid' bit
    // of the corresponding instance of csubSessionAvailableIdentities is '1'. The
    // type is interface{} with range: 0..4294967295.
    CsubSessionAcctSessionId interface{}

    // This object indicates the DNIS number dialed by the subscriber to access
    // the 'calling station' or AN.  This column is valid only if the 'dnis' bit
    // of the corresponding instance of csubSessionAvailableIdentities is '1'. The
    // type is string.
    CsubSessionDnis interface{}

    // This object indicates the type of media providing access to the subscriber.
    // This column is valid only if the 'media' bit of the corresponding instance
    // of csubSessionAvailableIdentities is '1'. The type is SubscriberMediaType.
    CsubSessionMedia interface{}

    // This object indicates whether the subscriber session was established using
    // multi-link PPP negotiation.  This column is valid only if the
    // 'mlpNegotiated' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is bool.
    CsubSessionMlpNegotiated interface{}

    // This object indicates the type of protocol providing access to the
    // subscriber.  This column is valid only if the 'protocol' bit of the
    // corresponding instance of csubSessionAvailableIdentities is '1'. The type
    // is SubscriberProtocolType.
    CsubSessionProtocol interface{}

    // This object indicates the name of the class matching the DHCP DISCOVER
    // message received from the subscriber.  This column is valid only if the
    // 'dhcpClass' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string.
    CsubSessionDhcpClass interface{}

    // This object indicates the name of the VPDN used to carry the subscriber
    // session to the system.  This column is valid only if the 'tunnelName' bit
    // of the corresponding instance of csubSessionAvailableIdentities is '1'. The
    // type is string.
    CsubSessionTunnelName interface{}

    // This object indicates the location of the subscriber. The type is string.
    CsubSessionLocationIdentifier interface{}

    // This object indicates the name used to identify the services subscribed by
    // a particular session. The type is string.
    CsubSessionServiceIdentifier interface{}

    // This object indicates when the subscriber session is updated with new
    // policy information. The type is string.
    CsubSessionLastChanged interface{}

    // This object indicates the type of IP address assigned to the subscriber on
    // the customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'.  In Dual stack scenarios both
    // 'csubSessionNativeIpAddrType' and  'csubSessionNativeIpAddrType2' will be
    // valid. The type is InetAddressType.
    CsubSessionNativeIpAddrType2 interface{}

    // This object indicates the IP address assigned to the subscriber on the
    // customer-facing side of the system.  This column is valid only if the
    // 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    CsubSessionNativeIpAddr2 interface{}

    // This object indicates the corresponding mask for the IP address assigned to
    // the subscriber on the customer-facing side of the system.  This column is
    // valid only if the 'nativeIpAddress' bit of the corresponding instance of
    // csubSessionAvailableIdentities is '1'. The type is string with length:
    // 0..255.
    CsubSessionNativeIpMask2 interface{}
}

func (csubSessionEntry *CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry) GetEntityData() *types.CommonEntityData {
    csubSessionEntry.EntityData.YFilter = csubSessionEntry.YFilter
    csubSessionEntry.EntityData.YangName = "csubSessionEntry"
    csubSessionEntry.EntityData.BundleName = "cisco_ios_xe"
    csubSessionEntry.EntityData.ParentYangName = "csubSessionTable"
    csubSessionEntry.EntityData.SegmentPath = "csubSessionEntry" + types.AddKeyToken(csubSessionEntry.IfIndex, "ifIndex")
    csubSessionEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubSessionTable/" + csubSessionEntry.EntityData.SegmentPath
    csubSessionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubSessionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubSessionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubSessionEntry.EntityData.Children = types.NewOrderedMap()
    csubSessionEntry.EntityData.Leafs = types.NewOrderedMap()
    csubSessionEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", csubSessionEntry.IfIndex})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionType", types.YLeaf{"CsubSessionType", csubSessionEntry.CsubSessionType})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionIpAddrAssignment", types.YLeaf{"CsubSessionIpAddrAssignment", csubSessionEntry.CsubSessionIpAddrAssignment})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionState", types.YLeaf{"CsubSessionState", csubSessionEntry.CsubSessionState})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionAuthenticated", types.YLeaf{"CsubSessionAuthenticated", csubSessionEntry.CsubSessionAuthenticated})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionRedundancyMode", types.YLeaf{"CsubSessionRedundancyMode", csubSessionEntry.CsubSessionRedundancyMode})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionCreationTime", types.YLeaf{"CsubSessionCreationTime", csubSessionEntry.CsubSessionCreationTime})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDerivedCfg", types.YLeaf{"CsubSessionDerivedCfg", csubSessionEntry.CsubSessionDerivedCfg})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionAvailableIdentities", types.YLeaf{"CsubSessionAvailableIdentities", csubSessionEntry.CsubSessionAvailableIdentities})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionSubscriberLabel", types.YLeaf{"CsubSessionSubscriberLabel", csubSessionEntry.CsubSessionSubscriberLabel})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionMacAddress", types.YLeaf{"CsubSessionMacAddress", csubSessionEntry.CsubSessionMacAddress})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNativeVrf", types.YLeaf{"CsubSessionNativeVrf", csubSessionEntry.CsubSessionNativeVrf})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNativeIpAddrType", types.YLeaf{"CsubSessionNativeIpAddrType", csubSessionEntry.CsubSessionNativeIpAddrType})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNativeIpAddr", types.YLeaf{"CsubSessionNativeIpAddr", csubSessionEntry.CsubSessionNativeIpAddr})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNativeIpMask", types.YLeaf{"CsubSessionNativeIpMask", csubSessionEntry.CsubSessionNativeIpMask})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDomainVrf", types.YLeaf{"CsubSessionDomainVrf", csubSessionEntry.CsubSessionDomainVrf})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDomainIpAddrType", types.YLeaf{"CsubSessionDomainIpAddrType", csubSessionEntry.CsubSessionDomainIpAddrType})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDomainIpAddr", types.YLeaf{"CsubSessionDomainIpAddr", csubSessionEntry.CsubSessionDomainIpAddr})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDomainIpMask", types.YLeaf{"CsubSessionDomainIpMask", csubSessionEntry.CsubSessionDomainIpMask})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionPbhk", types.YLeaf{"CsubSessionPbhk", csubSessionEntry.CsubSessionPbhk})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionRemoteId", types.YLeaf{"CsubSessionRemoteId", csubSessionEntry.CsubSessionRemoteId})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionCircuitId", types.YLeaf{"CsubSessionCircuitId", csubSessionEntry.CsubSessionCircuitId})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNasPort", types.YLeaf{"CsubSessionNasPort", csubSessionEntry.CsubSessionNasPort})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDomain", types.YLeaf{"CsubSessionDomain", csubSessionEntry.CsubSessionDomain})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionUsername", types.YLeaf{"CsubSessionUsername", csubSessionEntry.CsubSessionUsername})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionAcctSessionId", types.YLeaf{"CsubSessionAcctSessionId", csubSessionEntry.CsubSessionAcctSessionId})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDnis", types.YLeaf{"CsubSessionDnis", csubSessionEntry.CsubSessionDnis})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionMedia", types.YLeaf{"CsubSessionMedia", csubSessionEntry.CsubSessionMedia})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionMlpNegotiated", types.YLeaf{"CsubSessionMlpNegotiated", csubSessionEntry.CsubSessionMlpNegotiated})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionProtocol", types.YLeaf{"CsubSessionProtocol", csubSessionEntry.CsubSessionProtocol})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionDhcpClass", types.YLeaf{"CsubSessionDhcpClass", csubSessionEntry.CsubSessionDhcpClass})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionTunnelName", types.YLeaf{"CsubSessionTunnelName", csubSessionEntry.CsubSessionTunnelName})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionLocationIdentifier", types.YLeaf{"CsubSessionLocationIdentifier", csubSessionEntry.CsubSessionLocationIdentifier})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionServiceIdentifier", types.YLeaf{"CsubSessionServiceIdentifier", csubSessionEntry.CsubSessionServiceIdentifier})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionLastChanged", types.YLeaf{"CsubSessionLastChanged", csubSessionEntry.CsubSessionLastChanged})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNativeIpAddrType2", types.YLeaf{"CsubSessionNativeIpAddrType2", csubSessionEntry.CsubSessionNativeIpAddrType2})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNativeIpAddr2", types.YLeaf{"CsubSessionNativeIpAddr2", csubSessionEntry.CsubSessionNativeIpAddr2})
    csubSessionEntry.EntityData.Leafs.Append("csubSessionNativeIpMask2", types.YLeaf{"CsubSessionNativeIpMask2", csubSessionEntry.CsubSessionNativeIpMask2})

    csubSessionEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(csubSessionEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment represents         profile.
type CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment string

const (
    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_none CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "none"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_other CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "other"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_static CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "static"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_localPool CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "localPool"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_dhcpv4 CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "dhcpv4"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_dhcpv6 CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "dhcpv6"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_userProfileIpAddr CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "userProfileIpAddr"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_userProfileIpSubnet CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "userProfileIpSubnet"

    CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment_userProfileNamedPool CISCOSUBSCRIBERSESSIONMIB_CsubSessionTable_CsubSessionEntry_CsubSessionIpAddrAssignment = "userProfileNamedPool"
)

// CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable
// This table describes a list of subscriber sessions currently
// maintained by the system.  The tables sorts the subscriber
// sessions first by the subscriber session's type and second by
// the ifIndex assigned to the subscriber session.
type CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry identifies a subscriber session.  An entry exists for a
    // corresponding entry in the ifTable describing a subscriber session. 
    // Currently, subscriber sessions must have one of the following ifType
    // values:      'ppp'     'ipSubscriberSession'     'l2SubscriberSession'  The
    // system creates an entry when it establishes a subscriber session. 
    // Likewise, the system destroys an entry when it terminates the corresponding
    // subscriber session. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable_CsubSessionByTypeEntry.
    CsubSessionByTypeEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable_CsubSessionByTypeEntry
}

func (csubSessionByTypeTable *CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable) GetEntityData() *types.CommonEntityData {
    csubSessionByTypeTable.EntityData.YFilter = csubSessionByTypeTable.YFilter
    csubSessionByTypeTable.EntityData.YangName = "csubSessionByTypeTable"
    csubSessionByTypeTable.EntityData.BundleName = "cisco_ios_xe"
    csubSessionByTypeTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubSessionByTypeTable.EntityData.SegmentPath = "csubSessionByTypeTable"
    csubSessionByTypeTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubSessionByTypeTable.EntityData.SegmentPath
    csubSessionByTypeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubSessionByTypeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubSessionByTypeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubSessionByTypeTable.EntityData.Children = types.NewOrderedMap()
    csubSessionByTypeTable.EntityData.Children.Append("csubSessionByTypeEntry", types.YChild{"CsubSessionByTypeEntry", nil})
    for i := range csubSessionByTypeTable.CsubSessionByTypeEntry {
        csubSessionByTypeTable.EntityData.Children.Append(types.GetSegmentPath(csubSessionByTypeTable.CsubSessionByTypeEntry[i]), types.YChild{"CsubSessionByTypeEntry", csubSessionByTypeTable.CsubSessionByTypeEntry[i]})
    }
    csubSessionByTypeTable.EntityData.Leafs = types.NewOrderedMap()

    csubSessionByTypeTable.EntityData.YListKeys = []string {}

    return &(csubSessionByTypeTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable_CsubSessionByTypeEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable_CsubSessionByTypeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object indicates the type of the subscriber
    // session. The type is SubSessionType.
    CsubSessionByType interface{}

    // This attribute is a key. This object indicates the ifIndex assigned to the
    // subscriber session. The type is interface{} with range: 1..2147483647.
    CsubSessionIfIndex interface{}
}

func (csubSessionByTypeEntry *CISCOSUBSCRIBERSESSIONMIB_CsubSessionByTypeTable_CsubSessionByTypeEntry) GetEntityData() *types.CommonEntityData {
    csubSessionByTypeEntry.EntityData.YFilter = csubSessionByTypeEntry.YFilter
    csubSessionByTypeEntry.EntityData.YangName = "csubSessionByTypeEntry"
    csubSessionByTypeEntry.EntityData.BundleName = "cisco_ios_xe"
    csubSessionByTypeEntry.EntityData.ParentYangName = "csubSessionByTypeTable"
    csubSessionByTypeEntry.EntityData.SegmentPath = "csubSessionByTypeEntry" + types.AddKeyToken(csubSessionByTypeEntry.CsubSessionByType, "csubSessionByType") + types.AddKeyToken(csubSessionByTypeEntry.CsubSessionIfIndex, "csubSessionIfIndex")
    csubSessionByTypeEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubSessionByTypeTable/" + csubSessionByTypeEntry.EntityData.SegmentPath
    csubSessionByTypeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubSessionByTypeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubSessionByTypeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubSessionByTypeEntry.EntityData.Children = types.NewOrderedMap()
    csubSessionByTypeEntry.EntityData.Leafs = types.NewOrderedMap()
    csubSessionByTypeEntry.EntityData.Leafs.Append("csubSessionByType", types.YLeaf{"CsubSessionByType", csubSessionByTypeEntry.CsubSessionByType})
    csubSessionByTypeEntry.EntityData.Leafs.Append("csubSessionIfIndex", types.YLeaf{"CsubSessionIfIndex", csubSessionByTypeEntry.CsubSessionIfIndex})

    csubSessionByTypeEntry.EntityData.YListKeys = []string {"CsubSessionByType", "CsubSessionIfIndex"}

    return &(csubSessionByTypeEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable
// This table contains sets of aggregated statistics relating to
// subscriber sessions, where each set has a unique scope of
// aggregation.
type CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable struct {
    EntityData types.CommonEntityData
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
    // slice of CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry.
    CsubAggStatsEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry
}

func (csubAggStatsTable *CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable) GetEntityData() *types.CommonEntityData {
    csubAggStatsTable.EntityData.YFilter = csubAggStatsTable.YFilter
    csubAggStatsTable.EntityData.YangName = "csubAggStatsTable"
    csubAggStatsTable.EntityData.BundleName = "cisco_ios_xe"
    csubAggStatsTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubAggStatsTable.EntityData.SegmentPath = "csubAggStatsTable"
    csubAggStatsTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubAggStatsTable.EntityData.SegmentPath
    csubAggStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubAggStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubAggStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubAggStatsTable.EntityData.Children = types.NewOrderedMap()
    csubAggStatsTable.EntityData.Children.Append("csubAggStatsEntry", types.YChild{"CsubAggStatsEntry", nil})
    for i := range csubAggStatsTable.CsubAggStatsEntry {
        csubAggStatsTable.EntityData.Children.Append(types.GetSegmentPath(csubAggStatsTable.CsubAggStatsEntry[i]), types.YChild{"CsubAggStatsEntry", csubAggStatsTable.CsubAggStatsEntry[i]})
    }
    csubAggStatsTable.EntityData.Leafs = types.NewOrderedMap()

    csubAggStatsTable.EntityData.YListKeys = []string {}

    return &(csubAggStatsTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object indicates format of the
    // csubAggStatsPoint for this entry.   The format for the csubAggStatsPoint is
    // as follows:   csubAggStatsPointType      csubAggStatsPoint 
    // ----------------------     ------------------      'physical'              
    // PhysicalIndex      'interface'                InterfaceIndex. The type is
    // CsubAggStatsPointType.
    CsubAggStatsPointType interface{}

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
    CsubAggStatsPoint interface{}

    // This attribute is a key. This object indicates one of the determining
    // factors affecting the 'scope of aggregation' for the statistics contained
    // by the row.  If the value of this column is 'all', then the 'scope of
    // aggregation' may include all subscriber sessions, regardless of type.  If
    // the value of this column is not 'all', then the 'scope of aggregation' may
    // include subscriber sessions of the indicated subscriber session type. The
    // type is SubSessionType.
    CsubAggStatsSessionType interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are in the PENDING state. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsPendingSessions interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are in the UP state. The type is interface{}
    // with range: 0..4294967295. Units are sessions.
    CsubAggStatsUpSessions interface{}

    // This object indicates the current number of subscriber session within the
    // 'scope of aggregation' that have been authenticated. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsAuthSessions interface{}

    // This object indicates the current number of subscriber session within the
    // 'scope of aggregation' that have not been authenticated. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsUnAuthSessions interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are less resource intensive. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsLightWeightSessions interface{}

    // This object indicates the current number of subscriber sessions within the
    // 'scope of aggregation' that are redundant (i.e.,  sessions with a
    // csubSessionRedundancyMode of 'standby'). The type is interface{} with
    // range: 0..4294967295. Units are sessions.
    CsubAggStatsRedSessions interface{}

    // This object indicates the highest number of subscriber sessions within the
    // 'scope of aggregation' observed simultaneously in the UP state since the
    // last discontinuity time. The type is interface{} with range: 0..4294967295.
    // Units are sessions.
    CsubAggStatsHighUpSessions interface{}

    // This object indicates the average time subscriber sessions within the
    // 'scope of aggregation' spent in the UP state.  The system calculates this
    // average over all subscriber sessions maintained by all nodes contained by
    // the 'scope of aggregation' since the last discontinuity time. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    CsubAggStatsAvgSessionUptime interface{}

    // This object indicates the average rate (per minute) at which the nodes
    // contained by the 'scope of aggregation' have established new subscriber
    // sessions. The type is interface{} with range: 0..4294967295. Units are
    // sessions per minute.
    CsubAggStatsAvgSessionRPM interface{}

    // This object indicates the average rate (per hour) at which the nodes
    // contained by the 'scope of aggregation' have established new subscriber
    // sessions. The type is interface{} with range: 0..4294967295. Units are
    // sessions per hour.
    CsubAggStatsAvgSessionRPH interface{}

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
    CsubAggStatsThrottleEngagements interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created since the discontinuity time. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    CsubAggStatsTotalCreatedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect since the last discontinuity time. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    CsubAggStatsTotalFailedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state since the last
    // discontinuity time.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of csubAggStatsDiscontinuityTime. The type is
    // interface{} with range: 0..18446744073709551615. Units are sessions.
    CsubAggStatsTotalUpSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state since the last discontinuity time.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    CsubAggStatsTotalAuthSessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event since the last discontinuity time. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // csubAggStatsDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    CsubAggStatsTotalDiscSessions interface{}

    // This object indicates the total number of subscriber sessions that are less
    // resource intensive. The type is interface{} with range:
    // 0..18446744073709551615. Units are sessions.
    CsubAggStatsTotalLightWeightSessions interface{}

    // This object indicates the total number of differential traffic classes on
    // subscriber sessions. IP ACLs are used to create differential flows(Traffic
    // Classes).  Each Traffic Class can have a different set of features applied.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // sessions.
    CsubAggStatsTotalFlowsUp interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created during the last 24 hours.  The system
    // calculates the value of this column by summing the values of all instances
    // of csubAggStatsIntCreatedSessions that expand this row and have a
    // corresponding csubAggStatsIntValid of 'true'. The type is interface{} with
    // range: 0..4294967295. Units are sessions.
    CsubAggStatsDayCreatedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect during the last 24 hours.  The system
    // calculates the value of this column by summing the values of all instances
    // of csubAggStatsIntFailedSessions that expand this row and have a
    // corresponding csubAggStatsIntValid of 'true'. The type is interface{} with
    // range: 0..4294967295. Units are sessions.
    CsubAggStatsDayFailedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state during the last 24
    // hours.  The system calculates the value of this column by summing the
    // values of all instances of csubAggStatsIntUpSessions that expand this row
    // and have a corresponding csubAggStatsIntValid of 'true'. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsDayUpSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state during the last 24 hours.  The system calculates the
    // value of this column by summing the values of all instances of
    // csubAggStatsIntAuthSessions that expand this row and have a corresponding
    // csubAggStatsIntValid of 'true'. The type is interface{} with range:
    // 0..4294967295. Units are sessions.
    CsubAggStatsDayAuthSessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event during the last 24 hours.  The system calculates
    // the value of this column by summing the values of all instances of
    // csubAggStatsIntDiscSessions that expand this row and have a corresponding
    // csubAggStatsIntValid of 'true'. The type is interface{} with range:
    // 0..4294967295. Units are sessions.
    CsubAggStatsDayDiscSessions interface{}

    // This object indicates the time that has elapsed since the beginning of the
    // current 15-minute measurement interval.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, then the value of this column will be the maximum value.
    // The type is interface{} with range: 0..899. Units are seconds.
    CsubAggStatsCurrTimeElapsed interface{}

    // This object indicates the number of intervals for which data was collected.
    // The value of this column will be '96' unless the measurement was started
    // (or restarted) within 1,440 minutes, in which case the value will be the
    // number of complete 15-minute intervals for which the system has at least
    // some data.  In certain cases it is possible that some intervals are
    // unavailable, in which case the value of this column will be maximum
    // interval number for which data is available. The type is interface{} with
    // range: 0..96. Units are intervals.
    CsubAggStatsCurrValidIntervals interface{}

    // This object indicates the number of intervals in the range from 0 to
    // csubCurrValidIntervals for which no data is available.  This object will
    // typically be '0' except in certain circumstances when some intervals are
    // unavailable. The type is interface{} with range: 0..96. Units are
    // intervals.
    CsubAggStatsCurrInvalidIntervals interface{}

    // This object indicates the current number of differential traffic classes on
    // subscriber sessions currently UP. IP ACLs are used to create differential
    // flows (Traffic Classes).Each Traffic Class can have a different set of
    // features applied. The type is interface{} with range: 0..96. Units are
    // intervals.
    CsubAggStatsCurrFlowsUp interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created during the current 15-minute interval. The
    // type is interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsCurrCreatedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect during the current 15-minute interval. The
    // type is interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsCurrFailedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state during the current
    // 15-minute interval. The type is interface{} with range: 0..4294967295.
    // Units are sessions.
    CsubAggStatsCurrUpSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state during the current 15-minute interval. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsCurrAuthSessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event during the current 15-minute interval. The type
    // is interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsCurrDiscSessions interface{}

    // The date and time (as determined by the system's clock) of the most recent
    // occurrence of an event affecting the  continuity of the aggregation
    // statistics for this aggregation  point. The type is string.
    CsubAggStatsDiscontinuityTime interface{}
}

func (csubAggStatsEntry *CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry) GetEntityData() *types.CommonEntityData {
    csubAggStatsEntry.EntityData.YFilter = csubAggStatsEntry.YFilter
    csubAggStatsEntry.EntityData.YangName = "csubAggStatsEntry"
    csubAggStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    csubAggStatsEntry.EntityData.ParentYangName = "csubAggStatsTable"
    csubAggStatsEntry.EntityData.SegmentPath = "csubAggStatsEntry" + types.AddKeyToken(csubAggStatsEntry.CsubAggStatsPointType, "csubAggStatsPointType") + types.AddKeyToken(csubAggStatsEntry.CsubAggStatsPoint, "csubAggStatsPoint") + types.AddKeyToken(csubAggStatsEntry.CsubAggStatsSessionType, "csubAggStatsSessionType")
    csubAggStatsEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubAggStatsTable/" + csubAggStatsEntry.EntityData.SegmentPath
    csubAggStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubAggStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubAggStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubAggStatsEntry.EntityData.Children = types.NewOrderedMap()
    csubAggStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsPointType", types.YLeaf{"CsubAggStatsPointType", csubAggStatsEntry.CsubAggStatsPointType})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsPoint", types.YLeaf{"CsubAggStatsPoint", csubAggStatsEntry.CsubAggStatsPoint})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsSessionType", types.YLeaf{"CsubAggStatsSessionType", csubAggStatsEntry.CsubAggStatsSessionType})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsPendingSessions", types.YLeaf{"CsubAggStatsPendingSessions", csubAggStatsEntry.CsubAggStatsPendingSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsUpSessions", types.YLeaf{"CsubAggStatsUpSessions", csubAggStatsEntry.CsubAggStatsUpSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsAuthSessions", types.YLeaf{"CsubAggStatsAuthSessions", csubAggStatsEntry.CsubAggStatsAuthSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsUnAuthSessions", types.YLeaf{"CsubAggStatsUnAuthSessions", csubAggStatsEntry.CsubAggStatsUnAuthSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsLightWeightSessions", types.YLeaf{"CsubAggStatsLightWeightSessions", csubAggStatsEntry.CsubAggStatsLightWeightSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsRedSessions", types.YLeaf{"CsubAggStatsRedSessions", csubAggStatsEntry.CsubAggStatsRedSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsHighUpSessions", types.YLeaf{"CsubAggStatsHighUpSessions", csubAggStatsEntry.CsubAggStatsHighUpSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsAvgSessionUptime", types.YLeaf{"CsubAggStatsAvgSessionUptime", csubAggStatsEntry.CsubAggStatsAvgSessionUptime})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsAvgSessionRPM", types.YLeaf{"CsubAggStatsAvgSessionRPM", csubAggStatsEntry.CsubAggStatsAvgSessionRPM})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsAvgSessionRPH", types.YLeaf{"CsubAggStatsAvgSessionRPH", csubAggStatsEntry.CsubAggStatsAvgSessionRPH})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsThrottleEngagements", types.YLeaf{"CsubAggStatsThrottleEngagements", csubAggStatsEntry.CsubAggStatsThrottleEngagements})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsTotalCreatedSessions", types.YLeaf{"CsubAggStatsTotalCreatedSessions", csubAggStatsEntry.CsubAggStatsTotalCreatedSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsTotalFailedSessions", types.YLeaf{"CsubAggStatsTotalFailedSessions", csubAggStatsEntry.CsubAggStatsTotalFailedSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsTotalUpSessions", types.YLeaf{"CsubAggStatsTotalUpSessions", csubAggStatsEntry.CsubAggStatsTotalUpSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsTotalAuthSessions", types.YLeaf{"CsubAggStatsTotalAuthSessions", csubAggStatsEntry.CsubAggStatsTotalAuthSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsTotalDiscSessions", types.YLeaf{"CsubAggStatsTotalDiscSessions", csubAggStatsEntry.CsubAggStatsTotalDiscSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsTotalLightWeightSessions", types.YLeaf{"CsubAggStatsTotalLightWeightSessions", csubAggStatsEntry.CsubAggStatsTotalLightWeightSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsTotalFlowsUp", types.YLeaf{"CsubAggStatsTotalFlowsUp", csubAggStatsEntry.CsubAggStatsTotalFlowsUp})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsDayCreatedSessions", types.YLeaf{"CsubAggStatsDayCreatedSessions", csubAggStatsEntry.CsubAggStatsDayCreatedSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsDayFailedSessions", types.YLeaf{"CsubAggStatsDayFailedSessions", csubAggStatsEntry.CsubAggStatsDayFailedSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsDayUpSessions", types.YLeaf{"CsubAggStatsDayUpSessions", csubAggStatsEntry.CsubAggStatsDayUpSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsDayAuthSessions", types.YLeaf{"CsubAggStatsDayAuthSessions", csubAggStatsEntry.CsubAggStatsDayAuthSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsDayDiscSessions", types.YLeaf{"CsubAggStatsDayDiscSessions", csubAggStatsEntry.CsubAggStatsDayDiscSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrTimeElapsed", types.YLeaf{"CsubAggStatsCurrTimeElapsed", csubAggStatsEntry.CsubAggStatsCurrTimeElapsed})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrValidIntervals", types.YLeaf{"CsubAggStatsCurrValidIntervals", csubAggStatsEntry.CsubAggStatsCurrValidIntervals})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrInvalidIntervals", types.YLeaf{"CsubAggStatsCurrInvalidIntervals", csubAggStatsEntry.CsubAggStatsCurrInvalidIntervals})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrFlowsUp", types.YLeaf{"CsubAggStatsCurrFlowsUp", csubAggStatsEntry.CsubAggStatsCurrFlowsUp})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrCreatedSessions", types.YLeaf{"CsubAggStatsCurrCreatedSessions", csubAggStatsEntry.CsubAggStatsCurrCreatedSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrFailedSessions", types.YLeaf{"CsubAggStatsCurrFailedSessions", csubAggStatsEntry.CsubAggStatsCurrFailedSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrUpSessions", types.YLeaf{"CsubAggStatsCurrUpSessions", csubAggStatsEntry.CsubAggStatsCurrUpSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrAuthSessions", types.YLeaf{"CsubAggStatsCurrAuthSessions", csubAggStatsEntry.CsubAggStatsCurrAuthSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsCurrDiscSessions", types.YLeaf{"CsubAggStatsCurrDiscSessions", csubAggStatsEntry.CsubAggStatsCurrDiscSessions})
    csubAggStatsEntry.EntityData.Leafs.Append("csubAggStatsDiscontinuityTime", types.YLeaf{"CsubAggStatsDiscontinuityTime", csubAggStatsEntry.CsubAggStatsDiscontinuityTime})

    csubAggStatsEntry.EntityData.YListKeys = []string {"CsubAggStatsPointType", "CsubAggStatsPoint", "CsubAggStatsSessionType"}

    return &(csubAggStatsEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPointType represents     'interface'                InterfaceIndex
type CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPointType string

const (
    CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPointType_physical CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPointType = "physical"

    CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPointType_interface_ CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPointType = "interface"
)

// CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable
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
type CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry contains the aggregated subscriber session performance data
    // collected over a single 15-minute measurement interval within a 'scope of
    // aggregation'.  For further details regarding 'scope of aggregation', see
    // the descriptive text associated with the csubAggStatsEntry. The type is
    // slice of
    // CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable_CsubAggStatsIntEntry.
    CsubAggStatsIntEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable_CsubAggStatsIntEntry
}

func (csubAggStatsIntTable *CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable) GetEntityData() *types.CommonEntityData {
    csubAggStatsIntTable.EntityData.YFilter = csubAggStatsIntTable.YFilter
    csubAggStatsIntTable.EntityData.YangName = "csubAggStatsIntTable"
    csubAggStatsIntTable.EntityData.BundleName = "cisco_ios_xe"
    csubAggStatsIntTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubAggStatsIntTable.EntityData.SegmentPath = "csubAggStatsIntTable"
    csubAggStatsIntTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubAggStatsIntTable.EntityData.SegmentPath
    csubAggStatsIntTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubAggStatsIntTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubAggStatsIntTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubAggStatsIntTable.EntityData.Children = types.NewOrderedMap()
    csubAggStatsIntTable.EntityData.Children.Append("csubAggStatsIntEntry", types.YChild{"CsubAggStatsIntEntry", nil})
    for i := range csubAggStatsIntTable.CsubAggStatsIntEntry {
        csubAggStatsIntTable.EntityData.Children.Append(types.GetSegmentPath(csubAggStatsIntTable.CsubAggStatsIntEntry[i]), types.YChild{"CsubAggStatsIntEntry", csubAggStatsIntTable.CsubAggStatsIntEntry[i]})
    }
    csubAggStatsIntTable.EntityData.Leafs = types.NewOrderedMap()

    csubAggStatsIntTable.EntityData.YListKeys = []string {}

    return &(csubAggStatsIntTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable_CsubAggStatsIntEntry
// An entry contains the aggregated subscriber session performance
// data collected over a single 15-minute measurement interval
// within a 'scope of aggregation'.  For further details regarding
// 'scope of aggregation', see the descriptive text associated with
// the csubAggStatsEntry.
type CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable_CsubAggStatsIntEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is CsubAggStatsPointType. Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPointType
    CsubAggStatsPointType interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsPoint
    CsubAggStatsPoint interface{}

    // This attribute is a key. The type is SubSessionType. Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsTable_CsubAggStatsEntry_CsubAggStatsSessionType
    CsubAggStatsSessionType interface{}

    // This attribute is a key. This object indicates the interval number
    // identifying the 15-minute measurement interval for which aggregated
    // subscriber session performance data was successfully collected by the
    // system.  The interval identified by the value '1' represents the most
    // recent 15-minute measurement interval, and the interval identified by the
    // value (n) represents the interval immediately preceding the interval
    // identified by the value (n-1). The type is interface{} with range: 1..96.
    CsubAggStatsIntNumber interface{}

    // This object indicates whether the data for the 15-minute measurement
    // interval is valid. The type is bool.
    CsubAggStatsIntValid interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' created during the 15-minute measurement interval.
    // The type is interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsIntCreatedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that were in the PENDING state and terminated for
    // reasons other than disconnect during the 15-minute measurement interval.
    // The type is interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsIntFailedSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned to the UP state during the
    // 15-minute measurement interval. The type is interface{} with range:
    // 0..4294967295. Units are sessions.
    CsubAggStatsIntUpSessions interface{}

    // This object indicates the total number of subscriber sessions within the
    // 'scope of aggregation' that transitioned from the UNAUTHENTICATED to the
    // AUTHENTICATED state during the 15-minute measurement interval. The type is
    // interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsIntAuthSessions interface{}

    // This object indicates the total number of subscriber sessions terminated
    // due to a disconnect event during the 15-minute measurement interval. The
    // type is interface{} with range: 0..4294967295. Units are sessions.
    CsubAggStatsIntDiscSessions interface{}
}

func (csubAggStatsIntEntry *CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsIntTable_CsubAggStatsIntEntry) GetEntityData() *types.CommonEntityData {
    csubAggStatsIntEntry.EntityData.YFilter = csubAggStatsIntEntry.YFilter
    csubAggStatsIntEntry.EntityData.YangName = "csubAggStatsIntEntry"
    csubAggStatsIntEntry.EntityData.BundleName = "cisco_ios_xe"
    csubAggStatsIntEntry.EntityData.ParentYangName = "csubAggStatsIntTable"
    csubAggStatsIntEntry.EntityData.SegmentPath = "csubAggStatsIntEntry" + types.AddKeyToken(csubAggStatsIntEntry.CsubAggStatsPointType, "csubAggStatsPointType") + types.AddKeyToken(csubAggStatsIntEntry.CsubAggStatsPoint, "csubAggStatsPoint") + types.AddKeyToken(csubAggStatsIntEntry.CsubAggStatsSessionType, "csubAggStatsSessionType") + types.AddKeyToken(csubAggStatsIntEntry.CsubAggStatsIntNumber, "csubAggStatsIntNumber")
    csubAggStatsIntEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubAggStatsIntTable/" + csubAggStatsIntEntry.EntityData.SegmentPath
    csubAggStatsIntEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubAggStatsIntEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubAggStatsIntEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubAggStatsIntEntry.EntityData.Children = types.NewOrderedMap()
    csubAggStatsIntEntry.EntityData.Leafs = types.NewOrderedMap()
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsPointType", types.YLeaf{"CsubAggStatsPointType", csubAggStatsIntEntry.CsubAggStatsPointType})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsPoint", types.YLeaf{"CsubAggStatsPoint", csubAggStatsIntEntry.CsubAggStatsPoint})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsSessionType", types.YLeaf{"CsubAggStatsSessionType", csubAggStatsIntEntry.CsubAggStatsSessionType})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsIntNumber", types.YLeaf{"CsubAggStatsIntNumber", csubAggStatsIntEntry.CsubAggStatsIntNumber})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsIntValid", types.YLeaf{"CsubAggStatsIntValid", csubAggStatsIntEntry.CsubAggStatsIntValid})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsIntCreatedSessions", types.YLeaf{"CsubAggStatsIntCreatedSessions", csubAggStatsIntEntry.CsubAggStatsIntCreatedSessions})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsIntFailedSessions", types.YLeaf{"CsubAggStatsIntFailedSessions", csubAggStatsIntEntry.CsubAggStatsIntFailedSessions})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsIntUpSessions", types.YLeaf{"CsubAggStatsIntUpSessions", csubAggStatsIntEntry.CsubAggStatsIntUpSessions})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsIntAuthSessions", types.YLeaf{"CsubAggStatsIntAuthSessions", csubAggStatsIntEntry.CsubAggStatsIntAuthSessions})
    csubAggStatsIntEntry.EntityData.Leafs.Append("csubAggStatsIntDiscSessions", types.YLeaf{"CsubAggStatsIntDiscSessions", csubAggStatsIntEntry.CsubAggStatsIntDiscSessions})

    csubAggStatsIntEntry.EntityData.YListKeys = []string {"CsubAggStatsPointType", "CsubAggStatsPoint", "CsubAggStatsSessionType", "CsubAggStatsIntNumber"}

    return &(csubAggStatsIntEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable
// Please enter the Table Description here.
type CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table exists for each row in the csubAggStatsTable. Each row
    // defines the set of thresholds and evaluation attributes for an aggregation
    // point. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable_CsubAggStatsThreshEntry.
    CsubAggStatsThreshEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable_CsubAggStatsThreshEntry
}

func (csubAggStatsThreshTable *CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable) GetEntityData() *types.CommonEntityData {
    csubAggStatsThreshTable.EntityData.YFilter = csubAggStatsThreshTable.YFilter
    csubAggStatsThreshTable.EntityData.YangName = "csubAggStatsThreshTable"
    csubAggStatsThreshTable.EntityData.BundleName = "cisco_ios_xe"
    csubAggStatsThreshTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubAggStatsThreshTable.EntityData.SegmentPath = "csubAggStatsThreshTable"
    csubAggStatsThreshTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubAggStatsThreshTable.EntityData.SegmentPath
    csubAggStatsThreshTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubAggStatsThreshTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubAggStatsThreshTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubAggStatsThreshTable.EntityData.Children = types.NewOrderedMap()
    csubAggStatsThreshTable.EntityData.Children.Append("csubAggStatsThreshEntry", types.YChild{"CsubAggStatsThreshEntry", nil})
    for i := range csubAggStatsThreshTable.CsubAggStatsThreshEntry {
        csubAggStatsThreshTable.EntityData.Children.Append(types.GetSegmentPath(csubAggStatsThreshTable.CsubAggStatsThreshEntry[i]), types.YChild{"CsubAggStatsThreshEntry", csubAggStatsThreshTable.CsubAggStatsThreshEntry[i]})
    }
    csubAggStatsThreshTable.EntityData.Leafs = types.NewOrderedMap()

    csubAggStatsThreshTable.EntityData.YListKeys = []string {}

    return &(csubAggStatsThreshTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable_CsubAggStatsThreshEntry
// A row in this table exists for each row in the
// csubAggStatsTable.
// Each row defines the set of thresholds and evaluation attributes
// for an aggregation point.
type CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable_CsubAggStatsThreshEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This threshold, if non-zero, indicates the rising
    // threshold for the value of csubAggStatsUpSessions for the aggregation
    // point, When the current sample of csubAggStatsUpSessions is greater than or
    // equal to this threshold, and the value of csubAggStatsUpSessions for the
    // last sample interval was less than this thershold, the
    // csubSessionRisingNotif is triggered.           If the value of this
    // threshold is 0, the threshold is not evaluated. The type is interface{}
    // with range: 0..4294967295.
    CsubSessionRisingThresh interface{}

    // This threshold, if non-zero, indicates the falling threshold for the value
    // of csubAggStatsUpSessions for the aggregation point, When the current
    // sample of csubAggStatsUpSessions is less than or equal to this threshold,
    // and the value of csubAggStatsUpSessions for the last sample interval was
    // greater than this thershold, the csubSessionFallingNotif is triggered.     
    // If the value of this threshold is 0, the threshold is not evaluated. The
    // type is interface{} with range: 0..4294967295.
    CsubSessionFallingThresh interface{}

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
    CsubSessionDeltaPercentFallingThresh interface{}

    // The value of this object sets the number of seconds between samples for
    // threshold evaluation. For implementations capable of per-session event
    // evaluation of thresholds this object represents the maximum number of
    // seconds between samples. The type is interface{} with range: 0..900.
    CsubSessionThreshEvalInterval interface{}
}

func (csubAggStatsThreshEntry *CISCOSUBSCRIBERSESSIONMIB_CsubAggStatsThreshTable_CsubAggStatsThreshEntry) GetEntityData() *types.CommonEntityData {
    csubAggStatsThreshEntry.EntityData.YFilter = csubAggStatsThreshEntry.YFilter
    csubAggStatsThreshEntry.EntityData.YangName = "csubAggStatsThreshEntry"
    csubAggStatsThreshEntry.EntityData.BundleName = "cisco_ios_xe"
    csubAggStatsThreshEntry.EntityData.ParentYangName = "csubAggStatsThreshTable"
    csubAggStatsThreshEntry.EntityData.SegmentPath = "csubAggStatsThreshEntry" + types.AddKeyToken(csubAggStatsThreshEntry.CsubSessionRisingThresh, "csubSessionRisingThresh")
    csubAggStatsThreshEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubAggStatsThreshTable/" + csubAggStatsThreshEntry.EntityData.SegmentPath
    csubAggStatsThreshEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubAggStatsThreshEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubAggStatsThreshEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubAggStatsThreshEntry.EntityData.Children = types.NewOrderedMap()
    csubAggStatsThreshEntry.EntityData.Leafs = types.NewOrderedMap()
    csubAggStatsThreshEntry.EntityData.Leafs.Append("csubSessionRisingThresh", types.YLeaf{"CsubSessionRisingThresh", csubAggStatsThreshEntry.CsubSessionRisingThresh})
    csubAggStatsThreshEntry.EntityData.Leafs.Append("csubSessionFallingThresh", types.YLeaf{"CsubSessionFallingThresh", csubAggStatsThreshEntry.CsubSessionFallingThresh})
    csubAggStatsThreshEntry.EntityData.Leafs.Append("csubSessionDeltaPercentFallingThresh", types.YLeaf{"CsubSessionDeltaPercentFallingThresh", csubAggStatsThreshEntry.CsubSessionDeltaPercentFallingThresh})
    csubAggStatsThreshEntry.EntityData.Leafs.Append("csubSessionThreshEvalInterval", types.YLeaf{"CsubSessionThreshEvalInterval", csubAggStatsThreshEntry.CsubSessionThreshEvalInterval})

    csubAggStatsThreshEntry.EntityData.YListKeys = []string {"CsubSessionRisingThresh"}

    return &(csubAggStatsThreshEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobTable
// This table contains the subscriber session jobs submitted by
// the EMS/NMS.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobTable struct {
    EntityData types.CommonEntityData
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
    // CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry.
    CsubJobEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry
}

func (csubJobTable *CISCOSUBSCRIBERSESSIONMIB_CsubJobTable) GetEntityData() *types.CommonEntityData {
    csubJobTable.EntityData.YFilter = csubJobTable.YFilter
    csubJobTable.EntityData.YangName = "csubJobTable"
    csubJobTable.EntityData.BundleName = "cisco_ios_xe"
    csubJobTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubJobTable.EntityData.SegmentPath = "csubJobTable"
    csubJobTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubJobTable.EntityData.SegmentPath
    csubJobTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobTable.EntityData.Children = types.NewOrderedMap()
    csubJobTable.EntityData.Children.Append("csubJobEntry", types.YChild{"CsubJobEntry", nil})
    for i := range csubJobTable.CsubJobEntry {
        csubJobTable.EntityData.Children.Append(types.GetSegmentPath(csubJobTable.CsubJobEntry[i]), types.YChild{"CsubJobEntry", csubJobTable.CsubJobEntry[i]})
    }
    csubJobTable.EntityData.Leafs = types.NewOrderedMap()

    csubJobTable.EntityData.YListKeys = []string {}

    return &(csubJobTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object indicates an arbitrary, positive
    // integer-value uniquely identifying the subscriber session job. The type is
    // interface{} with range: 1..4294967295.
    CsubJobId interface{}

    // This object specifies the status of the subscriber session job.  The
    // following columns must be valid before activating a subscriber session job:
    // - csubJobStorage     - csubJobType     - csubJobControl  However, these
    // objects specify a default value.  Thus, it is possible to use create-and-go
    // semantics without setting any additional columns.  An implementation must
    // allow the EMS/NMS to modify any column when this column is 'active',
    // including columns defined in tables that have a one-to-one or sparse
    // dependent relationship on this table. The type is RowStatus.
    CsubJobStatus interface{}

    // This object specifies what happens to the subscriber session job upon
    // restart. The type is StorageType.
    CsubJobStorage interface{}

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
    // in     the csubJobMatchParamsTable. The type is CsubJobType.
    CsubJobType interface{}

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
    // this column is always 'noop'. The type is CsubJobControl.
    CsubJobControl interface{}

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
    // finished. The type is CsubJobState.
    CsubJobState interface{}

    // This object indicates the value of sysUpTime when the system started
    // executing the subscriber session job.  This value will be '0' when the
    // corresponding instance of csubJobState is 'idle' or 'pending'. The type is
    // interface{} with range: 0..4294967295.
    CsubJobStartedTime interface{}

    // This object indicates the value of sysUpTime when the system finished
    // executing the subscriber session job, for whatever reason.  This value will
    // be '0' when the corresponding instance of csubJobState is 'idle',
    // 'pending', or 'inProgress'. The type is interface{} with range:
    // 0..4294967295.
    CsubJobFinishedTime interface{}

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
    // from completing the job. The type is CsubJobFinishedReason.
    CsubJobFinishedReason interface{}
}

func (csubJobEntry *CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry) GetEntityData() *types.CommonEntityData {
    csubJobEntry.EntityData.YFilter = csubJobEntry.YFilter
    csubJobEntry.EntityData.YangName = "csubJobEntry"
    csubJobEntry.EntityData.BundleName = "cisco_ios_xe"
    csubJobEntry.EntityData.ParentYangName = "csubJobTable"
    csubJobEntry.EntityData.SegmentPath = "csubJobEntry" + types.AddKeyToken(csubJobEntry.CsubJobId, "csubJobId")
    csubJobEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubJobTable/" + csubJobEntry.EntityData.SegmentPath
    csubJobEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobEntry.EntityData.Children = types.NewOrderedMap()
    csubJobEntry.EntityData.Leafs = types.NewOrderedMap()
    csubJobEntry.EntityData.Leafs.Append("csubJobId", types.YLeaf{"CsubJobId", csubJobEntry.CsubJobId})
    csubJobEntry.EntityData.Leafs.Append("csubJobStatus", types.YLeaf{"CsubJobStatus", csubJobEntry.CsubJobStatus})
    csubJobEntry.EntityData.Leafs.Append("csubJobStorage", types.YLeaf{"CsubJobStorage", csubJobEntry.CsubJobStorage})
    csubJobEntry.EntityData.Leafs.Append("csubJobType", types.YLeaf{"CsubJobType", csubJobEntry.CsubJobType})
    csubJobEntry.EntityData.Leafs.Append("csubJobControl", types.YLeaf{"CsubJobControl", csubJobEntry.CsubJobControl})
    csubJobEntry.EntityData.Leafs.Append("csubJobState", types.YLeaf{"CsubJobState", csubJobEntry.CsubJobState})
    csubJobEntry.EntityData.Leafs.Append("csubJobStartedTime", types.YLeaf{"CsubJobStartedTime", csubJobEntry.CsubJobStartedTime})
    csubJobEntry.EntityData.Leafs.Append("csubJobFinishedTime", types.YLeaf{"CsubJobFinishedTime", csubJobEntry.CsubJobFinishedTime})
    csubJobEntry.EntityData.Leafs.Append("csubJobFinishedReason", types.YLeaf{"CsubJobFinishedReason", csubJobEntry.CsubJobFinishedReason})

    csubJobEntry.EntityData.YListKeys = []string {"CsubJobId"}

    return &(csubJobEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl represents When read, this column is always 'noop'.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl string

const (
    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl_noop CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl = "noop"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl_start CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl = "start"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl_abort CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl = "abort"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl_release CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobControl = "release"
)

// CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason represents         prevented it from completing the job.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason string

const (
    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason_invalid CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason = "invalid"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason_other CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason = "other"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason_normal CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason = "normal"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason_aborted CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason = "aborted"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason_insufficientResources CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason = "insufficientResources"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason_error_ CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobFinishedReason = "error"
)

// CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState represents         job finished.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState string

const (
    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState_idle CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState = "idle"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState_pending CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState = "pending"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState_inProgress CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState = "inProgress"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState_finished CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobState = "finished"
)

// CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType represents     the csubJobMatchParamsTable.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType string

const (
    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType_noop CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType = "noop"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType_query CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType = "query"

    CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType_clear CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobType = "clear"
)

// CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable
// This table contains subscriber session job parameters
// describing match criteria.
// 
// This table has a sparse-dependent relationship on the
// csubJobTable, containing a row for each job having a
// csubJobType of 'query' or 'clear'.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable struct {
    EntityData types.CommonEntityData
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
    // CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable_CsubJobMatchParamsEntry.
    CsubJobMatchParamsEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable_CsubJobMatchParamsEntry
}

func (csubJobMatchParamsTable *CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable) GetEntityData() *types.CommonEntityData {
    csubJobMatchParamsTable.EntityData.YFilter = csubJobMatchParamsTable.YFilter
    csubJobMatchParamsTable.EntityData.YangName = "csubJobMatchParamsTable"
    csubJobMatchParamsTable.EntityData.BundleName = "cisco_ios_xe"
    csubJobMatchParamsTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubJobMatchParamsTable.EntityData.SegmentPath = "csubJobMatchParamsTable"
    csubJobMatchParamsTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubJobMatchParamsTable.EntityData.SegmentPath
    csubJobMatchParamsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobMatchParamsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobMatchParamsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobMatchParamsTable.EntityData.Children = types.NewOrderedMap()
    csubJobMatchParamsTable.EntityData.Children.Append("csubJobMatchParamsEntry", types.YChild{"CsubJobMatchParamsEntry", nil})
    for i := range csubJobMatchParamsTable.CsubJobMatchParamsEntry {
        csubJobMatchParamsTable.EntityData.Children.Append(types.GetSegmentPath(csubJobMatchParamsTable.CsubJobMatchParamsEntry[i]), types.YChild{"CsubJobMatchParamsEntry", csubJobMatchParamsTable.CsubJobMatchParamsEntry[i]})
    }
    csubJobMatchParamsTable.EntityData.Leafs = types.NewOrderedMap()

    csubJobMatchParamsTable.EntityData.YListKeys = []string {}

    return &(csubJobMatchParamsTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable_CsubJobMatchParamsEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable_CsubJobMatchParamsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobId
    CsubJobId interface{}

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
    CsubJobMatchIdentities interface{}

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
    CsubJobMatchOtherParams interface{}

    // This object specifies the subscriber label that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'subscriberLabel' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is interface{} with
    // range: 0..4294967295.
    CsubJobMatchSubscriberLabel interface{}

    // This object specifies the MAC address that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'macAddress' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    CsubJobMatchMacAddress interface{}

    // This object specifies the native VRF the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'nativeVrf' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchNativeVrf interface{}

    // This object specifies the type of Internet address specified by
    // csubJobMatchNativeIpAddr and csubJobMatchNativeIpMask.  This value is valid
    // only if the 'nativeIpAddress' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is InetAddressType.
    CsubJobMatchNativeIpAddrType interface{}

    // This object specifies the native IP address that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'nativeIpAddress' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string with length:
    // 0..255.
    CsubJobMatchNativeIpAddr interface{}

    // This object specifies the mask used when matching the native IP address
    // against subscriber sessions in the process of executing a subscriber
    // session job.  This value is valid only if the 'nativeIpAddress' bit of the
    // corresponding instance of csubJobMatchIdentities is '1'. The type is string
    // with length: 0..255.
    CsubJobMatchNativeIpMask interface{}

    // This object specifies the domain VRF the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'domainVrf' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchDomainVrf interface{}

    // This object specifies the type of Internet address specified by
    // csubJobMatchDomainIpAddr and csubJobMatchDomainIpMask.  This value is valid
    // only if the 'domainIpAddress' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is InetAddressType.
    CsubJobMatchDomainIpAddrType interface{}

    // This object specifies the domain IP address that the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'domainIpAddress' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string with length:
    // 0..255.
    CsubJobMatchDomainIpAddr interface{}

    // This object specifies the mask used when matching the domain IP address
    // against subscriber sessions in the process of executing a subscriber
    // session job.  This value is valid only if the 'domainIpAddress' bit of the
    // corresponding instance of csubJobMatchIdentities is '1'. The type is string
    // with length: 0..255.
    CsubJobMatchDomainIpMask interface{}

    // This object specifies the PBHK that the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'pbhk' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchPbhk interface{}

    // This object specifies the Remote-Id the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'remoteId' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchRemoteId interface{}

    // This object specifies the Circuit-Id the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'circuitId' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchCircuitId interface{}

    // This object specifies the NAS port-identifier the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'nasPort' bit of the corresponding instance
    // of csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchNasPort interface{}

    // This object specifies the domain the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'domain' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchDomain interface{}

    // This object specifies the username the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'username' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchUsername interface{}

    // This object specifies the accounting session identifier the system matches
    // against subscriber sessions in the process of executing a subscriber
    // session job.  This value is valid only if the 'accountingSid' bit of the
    // corresponding instance of csubJobMatchIdentities is '1'. The type is
    // interface{} with range: 0..4294967295.
    CsubJobMatchAcctSessionId interface{}

    // This object specifies the DNIS number the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'dnis' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchDnis interface{}

    // This object specifies the media type the system matches against subscriber
    // sessions in the process of executing a subscriber session job.  This value
    // is valid only if the 'media' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is SubscriberMediaType.
    CsubJobMatchMedia interface{}

    // This object specifies the MLP negotiated flag the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'mlpNegotiated' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is bool.
    CsubJobMatchMlpNegotiated interface{}

    // This object specifies the protocol type the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'protocol' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is
    // SubscriberProtocolType.
    CsubJobMatchProtocol interface{}

    // This object specifies the service name the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'serviceName' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchServiceName interface{}

    // This object specifies the DHCP class name the system matches against
    // subscriber sessions in the process of executing a subscriber session job. 
    // This value is valid only if the 'dhcpClass' bit of the corresponding
    // instance of csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchDhcpClass interface{}

    // This object specifies the tunnel name the system matches against subscriber
    // session in the process of executing a subscriber session job.  This value
    // is valid only if the 'tunnelName' bit of the corresponding instance of
    // csubJobMatchIdentities is '1'. The type is string.
    CsubJobMatchTunnelName interface{}

    // This object specifies the minimum interval of time a subscriber session can
    // remain dangling in order for the system to consider it a match in the
    // process of executing a subscriber session job. A 'dangling' subscriber
    // session is one in the PENDING state.  The value '0' cannot be written to an
    // instance of this object. However, it serves as a convenient value when the
    // column is not valid.  This value is valid only if the 'danglingDuration'
    // bit of the corresponding instance of csubJobMatchOtherParams is '1'. The
    // type is interface{} with range: 0..3600.
    CsubJobMatchDanglingDuration interface{}

    // This object specifies the state of a subscriber session in order for the
    // system to consider a match in the process of executing a subscriber session
    // job.  The value 'other' is not valid and an implementation should not allow
    // it to be written to this column.  This value is valid only if the 'state'
    // bit of the corresponding instance of csubJobMatchOtherParams is '1'. The
    // type is SubSessionState.
    CsubJobMatchState interface{}

    // This object specifies whether a subscriber session should be
    // unauthenticated for the system to consider a match in the process of
    // executing a subscriber session job.  If this column is 'false', then the
    // subscriber session job matches subscriber sessions that are
    // unauthenticated.  If this column is 'true', then the subscriber session job
    // matches subscriber session that are authenticated.  This value is valid
    // only if the 'authenticated' bit of the corresponding instance of
    // csubJobMatchParams is '1'. The type is bool.
    CsubJobMatchAuthenticated interface{}

    // This object specifies the redudancy mode of the subscriber session in order
    // for the system to consider a match in the process of executing a subscriber
    // session job.  The value 'other' is not valid and an implementation should
    // not allow it to be written to this column.  This value is valid only if the
    // 'redundancyMode' bit of the corresponding instance of
    // csubJobMatchOtherParams is '1'. The type is SubSessionRedundancyMode.
    CsubJobMatchRedundancyMode interface{}
}

func (csubJobMatchParamsEntry *CISCOSUBSCRIBERSESSIONMIB_CsubJobMatchParamsTable_CsubJobMatchParamsEntry) GetEntityData() *types.CommonEntityData {
    csubJobMatchParamsEntry.EntityData.YFilter = csubJobMatchParamsEntry.YFilter
    csubJobMatchParamsEntry.EntityData.YangName = "csubJobMatchParamsEntry"
    csubJobMatchParamsEntry.EntityData.BundleName = "cisco_ios_xe"
    csubJobMatchParamsEntry.EntityData.ParentYangName = "csubJobMatchParamsTable"
    csubJobMatchParamsEntry.EntityData.SegmentPath = "csubJobMatchParamsEntry" + types.AddKeyToken(csubJobMatchParamsEntry.CsubJobId, "csubJobId")
    csubJobMatchParamsEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubJobMatchParamsTable/" + csubJobMatchParamsEntry.EntityData.SegmentPath
    csubJobMatchParamsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobMatchParamsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobMatchParamsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobMatchParamsEntry.EntityData.Children = types.NewOrderedMap()
    csubJobMatchParamsEntry.EntityData.Leafs = types.NewOrderedMap()
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobId", types.YLeaf{"CsubJobId", csubJobMatchParamsEntry.CsubJobId})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchIdentities", types.YLeaf{"CsubJobMatchIdentities", csubJobMatchParamsEntry.CsubJobMatchIdentities})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchOtherParams", types.YLeaf{"CsubJobMatchOtherParams", csubJobMatchParamsEntry.CsubJobMatchOtherParams})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchSubscriberLabel", types.YLeaf{"CsubJobMatchSubscriberLabel", csubJobMatchParamsEntry.CsubJobMatchSubscriberLabel})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchMacAddress", types.YLeaf{"CsubJobMatchMacAddress", csubJobMatchParamsEntry.CsubJobMatchMacAddress})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchNativeVrf", types.YLeaf{"CsubJobMatchNativeVrf", csubJobMatchParamsEntry.CsubJobMatchNativeVrf})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchNativeIpAddrType", types.YLeaf{"CsubJobMatchNativeIpAddrType", csubJobMatchParamsEntry.CsubJobMatchNativeIpAddrType})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchNativeIpAddr", types.YLeaf{"CsubJobMatchNativeIpAddr", csubJobMatchParamsEntry.CsubJobMatchNativeIpAddr})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchNativeIpMask", types.YLeaf{"CsubJobMatchNativeIpMask", csubJobMatchParamsEntry.CsubJobMatchNativeIpMask})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDomainVrf", types.YLeaf{"CsubJobMatchDomainVrf", csubJobMatchParamsEntry.CsubJobMatchDomainVrf})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDomainIpAddrType", types.YLeaf{"CsubJobMatchDomainIpAddrType", csubJobMatchParamsEntry.CsubJobMatchDomainIpAddrType})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDomainIpAddr", types.YLeaf{"CsubJobMatchDomainIpAddr", csubJobMatchParamsEntry.CsubJobMatchDomainIpAddr})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDomainIpMask", types.YLeaf{"CsubJobMatchDomainIpMask", csubJobMatchParamsEntry.CsubJobMatchDomainIpMask})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchPbhk", types.YLeaf{"CsubJobMatchPbhk", csubJobMatchParamsEntry.CsubJobMatchPbhk})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchRemoteId", types.YLeaf{"CsubJobMatchRemoteId", csubJobMatchParamsEntry.CsubJobMatchRemoteId})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchCircuitId", types.YLeaf{"CsubJobMatchCircuitId", csubJobMatchParamsEntry.CsubJobMatchCircuitId})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchNasPort", types.YLeaf{"CsubJobMatchNasPort", csubJobMatchParamsEntry.CsubJobMatchNasPort})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDomain", types.YLeaf{"CsubJobMatchDomain", csubJobMatchParamsEntry.CsubJobMatchDomain})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchUsername", types.YLeaf{"CsubJobMatchUsername", csubJobMatchParamsEntry.CsubJobMatchUsername})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchAcctSessionId", types.YLeaf{"CsubJobMatchAcctSessionId", csubJobMatchParamsEntry.CsubJobMatchAcctSessionId})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDnis", types.YLeaf{"CsubJobMatchDnis", csubJobMatchParamsEntry.CsubJobMatchDnis})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchMedia", types.YLeaf{"CsubJobMatchMedia", csubJobMatchParamsEntry.CsubJobMatchMedia})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchMlpNegotiated", types.YLeaf{"CsubJobMatchMlpNegotiated", csubJobMatchParamsEntry.CsubJobMatchMlpNegotiated})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchProtocol", types.YLeaf{"CsubJobMatchProtocol", csubJobMatchParamsEntry.CsubJobMatchProtocol})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchServiceName", types.YLeaf{"CsubJobMatchServiceName", csubJobMatchParamsEntry.CsubJobMatchServiceName})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDhcpClass", types.YLeaf{"CsubJobMatchDhcpClass", csubJobMatchParamsEntry.CsubJobMatchDhcpClass})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchTunnelName", types.YLeaf{"CsubJobMatchTunnelName", csubJobMatchParamsEntry.CsubJobMatchTunnelName})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchDanglingDuration", types.YLeaf{"CsubJobMatchDanglingDuration", csubJobMatchParamsEntry.CsubJobMatchDanglingDuration})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchState", types.YLeaf{"CsubJobMatchState", csubJobMatchParamsEntry.CsubJobMatchState})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchAuthenticated", types.YLeaf{"CsubJobMatchAuthenticated", csubJobMatchParamsEntry.CsubJobMatchAuthenticated})
    csubJobMatchParamsEntry.EntityData.Leafs.Append("csubJobMatchRedundancyMode", types.YLeaf{"CsubJobMatchRedundancyMode", csubJobMatchParamsEntry.CsubJobMatchRedundancyMode})

    csubJobMatchParamsEntry.EntityData.YListKeys = []string {"CsubJobId"}

    return &(csubJobMatchParamsEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable
// This table contains subscriber session job parameters
// describing query parameters.
// 
// This table has a sparse-dependent relationship on the
// csubJobTable, containing a row for each job having a
// csubJobType of 'query'.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describes a set of subscriber session query parameters.  The
    // system automatically creates an entry when the EMS/NMS sets the
    // corresponding instance of csubJobType to 'query'.  Likewise, the system
    // automatically destroys an entry under the following circumstances:  1)  The
    // EMS/NMS destroys the corresponding row in the csubJobTable.  2)  The
    // EMS/NMS sets the corresponding instance of csubJobType to     'noop' or
    // 'clear'. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable_CsubJobQueryParamsEntry.
    CsubJobQueryParamsEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable_CsubJobQueryParamsEntry
}

func (csubJobQueryParamsTable *CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable) GetEntityData() *types.CommonEntityData {
    csubJobQueryParamsTable.EntityData.YFilter = csubJobQueryParamsTable.YFilter
    csubJobQueryParamsTable.EntityData.YangName = "csubJobQueryParamsTable"
    csubJobQueryParamsTable.EntityData.BundleName = "cisco_ios_xe"
    csubJobQueryParamsTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubJobQueryParamsTable.EntityData.SegmentPath = "csubJobQueryParamsTable"
    csubJobQueryParamsTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubJobQueryParamsTable.EntityData.SegmentPath
    csubJobQueryParamsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobQueryParamsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobQueryParamsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobQueryParamsTable.EntityData.Children = types.NewOrderedMap()
    csubJobQueryParamsTable.EntityData.Children.Append("csubJobQueryParamsEntry", types.YChild{"CsubJobQueryParamsEntry", nil})
    for i := range csubJobQueryParamsTable.CsubJobQueryParamsEntry {
        csubJobQueryParamsTable.EntityData.Children.Append(types.GetSegmentPath(csubJobQueryParamsTable.CsubJobQueryParamsEntry[i]), types.YChild{"CsubJobQueryParamsEntry", csubJobQueryParamsTable.CsubJobQueryParamsEntry[i]})
    }
    csubJobQueryParamsTable.EntityData.Leafs = types.NewOrderedMap()

    csubJobQueryParamsTable.EntityData.YListKeys = []string {}

    return &(csubJobQueryParamsTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable_CsubJobQueryParamsEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable_CsubJobQueryParamsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobId
    CsubJobId interface{}

    // This object specifies the first subscriber identity that the system uses
    // when sorting subscriber sessions into the final report corresponding to a
    // subscriber session query.  It is not valid to set this column to 'other' or
    // 'ifIndex'. The type is SubSessionIdentity.
    CsubJobQuerySortKey1 interface{}

    // This object specifies the second subscriber identity that the system uses
    // when sorting subscriber sessions into the final report corresponding to a
    // subscriber session query.  If it is the desire to have the final report
    // sorted on a single subscriber identity, then this column should be 'other'.
    // The type is SubSessionIdentity.
    CsubJobQuerySortKey2 interface{}

    // This object specifies the third subscriber identity that the system uses
    // when sorting subscriber sessions into the final report corresponding to a
    // subscriber session query.  If it is the desire to have the final report
    // sorted on one or two subscriber identities, then this column should be
    // 'other'. The type is SubSessionIdentity.
    CsubJobQuerySortKey3 interface{}

    // This object indicates the number of subscriber sessions that matched the
    // corresponding subscriber session query.  The value of this column should be
    // '0' unless the corresponding value of csubJobState is 'finished'.  The
    // value of this column should be '0' after the EMS/NMS sets the corresponding
    // csubJobControl to 'release'. The type is interface{} with range:
    // 0..4294967295.
    CsubJobQueryResultingReportSize interface{}
}

func (csubJobQueryParamsEntry *CISCOSUBSCRIBERSESSIONMIB_CsubJobQueryParamsTable_CsubJobQueryParamsEntry) GetEntityData() *types.CommonEntityData {
    csubJobQueryParamsEntry.EntityData.YFilter = csubJobQueryParamsEntry.YFilter
    csubJobQueryParamsEntry.EntityData.YangName = "csubJobQueryParamsEntry"
    csubJobQueryParamsEntry.EntityData.BundleName = "cisco_ios_xe"
    csubJobQueryParamsEntry.EntityData.ParentYangName = "csubJobQueryParamsTable"
    csubJobQueryParamsEntry.EntityData.SegmentPath = "csubJobQueryParamsEntry" + types.AddKeyToken(csubJobQueryParamsEntry.CsubJobId, "csubJobId")
    csubJobQueryParamsEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubJobQueryParamsTable/" + csubJobQueryParamsEntry.EntityData.SegmentPath
    csubJobQueryParamsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobQueryParamsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobQueryParamsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobQueryParamsEntry.EntityData.Children = types.NewOrderedMap()
    csubJobQueryParamsEntry.EntityData.Leafs = types.NewOrderedMap()
    csubJobQueryParamsEntry.EntityData.Leafs.Append("csubJobId", types.YLeaf{"CsubJobId", csubJobQueryParamsEntry.CsubJobId})
    csubJobQueryParamsEntry.EntityData.Leafs.Append("csubJobQuerySortKey1", types.YLeaf{"CsubJobQuerySortKey1", csubJobQueryParamsEntry.CsubJobQuerySortKey1})
    csubJobQueryParamsEntry.EntityData.Leafs.Append("csubJobQuerySortKey2", types.YLeaf{"CsubJobQuerySortKey2", csubJobQueryParamsEntry.CsubJobQuerySortKey2})
    csubJobQueryParamsEntry.EntityData.Leafs.Append("csubJobQuerySortKey3", types.YLeaf{"CsubJobQuerySortKey3", csubJobQueryParamsEntry.CsubJobQuerySortKey3})
    csubJobQueryParamsEntry.EntityData.Leafs.Append("csubJobQueryResultingReportSize", types.YLeaf{"CsubJobQueryResultingReportSize", csubJobQueryParamsEntry.CsubJobQueryResultingReportSize})

    csubJobQueryParamsEntry.EntityData.YListKeys = []string {"CsubJobId"}

    return &(csubJobQueryParamsEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable
// This table lists the subscriber session jobs currently pending
// in the subscriber session job queue.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable struct {
    EntityData types.CommonEntityData
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
    // CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable_CsubJobQueueEntry.
    CsubJobQueueEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable_CsubJobQueueEntry
}

func (csubJobQueueTable *CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable) GetEntityData() *types.CommonEntityData {
    csubJobQueueTable.EntityData.YFilter = csubJobQueueTable.YFilter
    csubJobQueueTable.EntityData.YangName = "csubJobQueueTable"
    csubJobQueueTable.EntityData.BundleName = "cisco_ios_xe"
    csubJobQueueTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubJobQueueTable.EntityData.SegmentPath = "csubJobQueueTable"
    csubJobQueueTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubJobQueueTable.EntityData.SegmentPath
    csubJobQueueTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobQueueTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobQueueTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobQueueTable.EntityData.Children = types.NewOrderedMap()
    csubJobQueueTable.EntityData.Children.Append("csubJobQueueEntry", types.YChild{"CsubJobQueueEntry", nil})
    for i := range csubJobQueueTable.CsubJobQueueEntry {
        csubJobQueueTable.EntityData.Children.Append(types.GetSegmentPath(csubJobQueueTable.CsubJobQueueEntry[i]), types.YChild{"CsubJobQueueEntry", csubJobQueueTable.CsubJobQueueEntry[i]})
    }
    csubJobQueueTable.EntityData.Leafs = types.NewOrderedMap()

    csubJobQueueTable.EntityData.YListKeys = []string {}

    return &(csubJobQueueTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable_CsubJobQueueEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable_CsubJobQueueEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object indicates an positive, integer-value
    // that uniquely identifies the entry in the table. The value of this object
    // starts at '1' and monotonically increases for each subscriber session job
    // inserted into the subscriber session job queue.  If the value of this
    // object is '4294967295', the system will reset it to '1' when it inserts the
    // next subscriber session job into the subscriber session job queue. The type
    // is interface{} with range: 1..4294967295.
    CsubJobQueueNumber interface{}

    // This object indicates the identifier associated with the subscriber session
    // job. The type is interface{} with range: 1..4294967295.
    CsubJobQueueJobId interface{}
}

func (csubJobQueueEntry *CISCOSUBSCRIBERSESSIONMIB_CsubJobQueueTable_CsubJobQueueEntry) GetEntityData() *types.CommonEntityData {
    csubJobQueueEntry.EntityData.YFilter = csubJobQueueEntry.YFilter
    csubJobQueueEntry.EntityData.YangName = "csubJobQueueEntry"
    csubJobQueueEntry.EntityData.BundleName = "cisco_ios_xe"
    csubJobQueueEntry.EntityData.ParentYangName = "csubJobQueueTable"
    csubJobQueueEntry.EntityData.SegmentPath = "csubJobQueueEntry" + types.AddKeyToken(csubJobQueueEntry.CsubJobQueueNumber, "csubJobQueueNumber")
    csubJobQueueEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubJobQueueTable/" + csubJobQueueEntry.EntityData.SegmentPath
    csubJobQueueEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobQueueEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobQueueEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobQueueEntry.EntityData.Children = types.NewOrderedMap()
    csubJobQueueEntry.EntityData.Leafs = types.NewOrderedMap()
    csubJobQueueEntry.EntityData.Leafs.Append("csubJobQueueNumber", types.YLeaf{"CsubJobQueueNumber", csubJobQueueEntry.CsubJobQueueNumber})
    csubJobQueueEntry.EntityData.Leafs.Append("csubJobQueueJobId", types.YLeaf{"CsubJobQueueJobId", csubJobQueueEntry.CsubJobQueueJobId})

    csubJobQueueEntry.EntityData.YListKeys = []string {"CsubJobQueueNumber"}

    return &(csubJobQueueEntry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable
// This table contains the reports corresponding to subscriber
// session jobs that have a csubJobType of 'query' and
// csubJobState of 'finished'.
// 
// This table has an expansion dependent relationship on the
// csubJobTable, containing zero or more rows for each job.
type CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable struct {
    EntityData types.CommonEntityData
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
    // CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable_CsubJobReportEntry.
    CsubJobReportEntry []*CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable_CsubJobReportEntry
}

func (csubJobReportTable *CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable) GetEntityData() *types.CommonEntityData {
    csubJobReportTable.EntityData.YFilter = csubJobReportTable.YFilter
    csubJobReportTable.EntityData.YangName = "csubJobReportTable"
    csubJobReportTable.EntityData.BundleName = "cisco_ios_xe"
    csubJobReportTable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubJobReportTable.EntityData.SegmentPath = "csubJobReportTable"
    csubJobReportTable.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/" + csubJobReportTable.EntityData.SegmentPath
    csubJobReportTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobReportTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobReportTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobReportTable.EntityData.Children = types.NewOrderedMap()
    csubJobReportTable.EntityData.Children.Append("csubJobReportEntry", types.YChild{"CsubJobReportEntry", nil})
    for i := range csubJobReportTable.CsubJobReportEntry {
        csubJobReportTable.EntityData.Children.Append(types.GetSegmentPath(csubJobReportTable.CsubJobReportEntry[i]), types.YChild{"CsubJobReportEntry", csubJobReportTable.CsubJobReportEntry[i]})
    }
    csubJobReportTable.EntityData.Leafs = types.NewOrderedMap()

    csubJobReportTable.EntityData.YListKeys = []string {}

    return &(csubJobReportTable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable_CsubJobReportEntry
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
type CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable_CsubJobReportEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_subscriber_session_mib.CISCOSUBSCRIBERSESSIONMIB_CsubJobTable_CsubJobEntry_CsubJobId
    CsubJobId interface{}

    // This attribute is a key. This object indicates an arbitrary, positive,
    // integer-value that uniquely identifies this entry in a report.  This
    // auxiliary value is necessary, as the corresponding subscriber session job
    // can specify up to three subscriber identities on which to sort the
    // subscriber sessions that end up in the final report. The type is
    // interface{} with range: 1..4294967295.
    CsubJobReportId interface{}

    // This object indicates the ifIndex-value assigned to the subscriber session
    // that satisfied the match criteria specified by the corresponding subscriber
    // session job having a csubJobType of 'query'. The type is interface{} with
    // range: 1..2147483647.
    CsubJobReportSession interface{}
}

func (csubJobReportEntry *CISCOSUBSCRIBERSESSIONMIB_CsubJobReportTable_CsubJobReportEntry) GetEntityData() *types.CommonEntityData {
    csubJobReportEntry.EntityData.YFilter = csubJobReportEntry.YFilter
    csubJobReportEntry.EntityData.YangName = "csubJobReportEntry"
    csubJobReportEntry.EntityData.BundleName = "cisco_ios_xe"
    csubJobReportEntry.EntityData.ParentYangName = "csubJobReportTable"
    csubJobReportEntry.EntityData.SegmentPath = "csubJobReportEntry" + types.AddKeyToken(csubJobReportEntry.CsubJobId, "csubJobId") + types.AddKeyToken(csubJobReportEntry.CsubJobReportId, "csubJobReportId")
    csubJobReportEntry.EntityData.AbsolutePath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB/csubJobReportTable/" + csubJobReportEntry.EntityData.SegmentPath
    csubJobReportEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubJobReportEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubJobReportEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubJobReportEntry.EntityData.Children = types.NewOrderedMap()
    csubJobReportEntry.EntityData.Leafs = types.NewOrderedMap()
    csubJobReportEntry.EntityData.Leafs.Append("csubJobId", types.YLeaf{"CsubJobId", csubJobReportEntry.CsubJobId})
    csubJobReportEntry.EntityData.Leafs.Append("csubJobReportId", types.YLeaf{"CsubJobReportId", csubJobReportEntry.CsubJobReportId})
    csubJobReportEntry.EntityData.Leafs.Append("csubJobReportSession", types.YLeaf{"CsubJobReportSession", csubJobReportEntry.CsubJobReportSession})

    csubJobReportEntry.EntityData.YListKeys = []string {"CsubJobId", "CsubJobReportId"}

    return &(csubJobReportEntry.EntityData)
}

