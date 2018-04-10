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

func (cISCOSUBSCRIBERSESSIONMIB *CISCOSUBSCRIBERSESSIONMIB) GetEntityData() *types.CommonEntityData {
    cISCOSUBSCRIBERSESSIONMIB.EntityData.YFilter = cISCOSUBSCRIBERSESSIONMIB.YFilter
    cISCOSUBSCRIBERSESSIONMIB.EntityData.YangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.SegmentPath = "CISCO-SUBSCRIBER-SESSION-MIB:CISCO-SUBSCRIBER-SESSION-MIB"
    cISCOSUBSCRIBERSESSIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSUBSCRIBERSESSIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSUBSCRIBERSESSIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubJob"] = types.YChild{"Csubjob", &cISCOSUBSCRIBERSESSIONMIB.Csubjob}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubAggThresh"] = types.YChild{"Csubaggthresh", &cISCOSUBSCRIBERSESSIONMIB.Csubaggthresh}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubSessionTable"] = types.YChild{"Csubsessiontable", &cISCOSUBSCRIBERSESSIONMIB.Csubsessiontable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubSessionByTypeTable"] = types.YChild{"Csubsessionbytypetable", &cISCOSUBSCRIBERSESSIONMIB.Csubsessionbytypetable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubAggStatsTable"] = types.YChild{"Csubaggstatstable", &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatstable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubAggStatsIntTable"] = types.YChild{"Csubaggstatsinttable", &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatsinttable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubAggStatsThreshTable"] = types.YChild{"Csubaggstatsthreshtable", &cISCOSUBSCRIBERSESSIONMIB.Csubaggstatsthreshtable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubJobTable"] = types.YChild{"Csubjobtable", &cISCOSUBSCRIBERSESSIONMIB.Csubjobtable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubJobMatchParamsTable"] = types.YChild{"Csubjobmatchparamstable", &cISCOSUBSCRIBERSESSIONMIB.Csubjobmatchparamstable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubJobQueryParamsTable"] = types.YChild{"Csubjobqueryparamstable", &cISCOSUBSCRIBERSESSIONMIB.Csubjobqueryparamstable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubJobQueueTable"] = types.YChild{"Csubjobqueuetable", &cISCOSUBSCRIBERSESSIONMIB.Csubjobqueuetable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Children["csubJobReportTable"] = types.YChild{"Csubjobreporttable", &cISCOSUBSCRIBERSESSIONMIB.Csubjobreporttable}
    cISCOSUBSCRIBERSESSIONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOSUBSCRIBERSESSIONMIB.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubjob
type CISCOSUBSCRIBERSESSIONMIB_Csubjob struct {
    EntityData types.CommonEntityData
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

func (csubjob *CISCOSUBSCRIBERSESSIONMIB_Csubjob) GetEntityData() *types.CommonEntityData {
    csubjob.EntityData.YFilter = csubjob.YFilter
    csubjob.EntityData.YangName = "csubJob"
    csubjob.EntityData.BundleName = "cisco_ios_xe"
    csubjob.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubjob.EntityData.SegmentPath = "csubJob"
    csubjob.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjob.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjob.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjob.EntityData.Children = make(map[string]types.YChild)
    csubjob.EntityData.Leafs = make(map[string]types.YLeaf)
    csubjob.EntityData.Leafs["csubJobFinishedNotifyEnable"] = types.YLeaf{"Csubjobfinishednotifyenable", csubjob.Csubjobfinishednotifyenable}
    csubjob.EntityData.Leafs["csubJobIndexedAttributes"] = types.YLeaf{"Csubjobindexedattributes", csubjob.Csubjobindexedattributes}
    csubjob.EntityData.Leafs["csubJobIdNext"] = types.YLeaf{"Csubjobidnext", csubjob.Csubjobidnext}
    csubjob.EntityData.Leafs["csubJobMaxNumber"] = types.YLeaf{"Csubjobmaxnumber", csubjob.Csubjobmaxnumber}
    csubjob.EntityData.Leafs["csubJobMaxLife"] = types.YLeaf{"Csubjobmaxlife", csubjob.Csubjobmaxlife}
    csubjob.EntityData.Leafs["csubJobCount"] = types.YLeaf{"Csubjobcount", csubjob.Csubjobcount}
    return &(csubjob.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh
type CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object enables or disables the generation of any of the csubAggStats*
    // threshold notifications. The type is bool.
    Csubaggstatsthreshnotifenable interface{}
}

func (csubaggthresh *CISCOSUBSCRIBERSESSIONMIB_Csubaggthresh) GetEntityData() *types.CommonEntityData {
    csubaggthresh.EntityData.YFilter = csubaggthresh.YFilter
    csubaggthresh.EntityData.YangName = "csubAggThresh"
    csubaggthresh.EntityData.BundleName = "cisco_ios_xe"
    csubaggthresh.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubaggthresh.EntityData.SegmentPath = "csubAggThresh"
    csubaggthresh.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubaggthresh.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubaggthresh.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubaggthresh.EntityData.Children = make(map[string]types.YChild)
    csubaggthresh.EntityData.Leafs = make(map[string]types.YLeaf)
    csubaggthresh.EntityData.Leafs["csubAggStatsThreshNotifEnable"] = types.YLeaf{"Csubaggstatsthreshnotifenable", csubaggthresh.Csubaggstatsthreshnotifenable}
    return &(csubaggthresh.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable
// This table describes a list of subscriber sessions currently
// maintained by the system.
// 
// This table has a sparse dependent relationship on the ifTable,
// containing a row for each interface having an interface type
// describing a subscriber session.
type CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable struct {
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
    // CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry.
    Csubsessionentry []CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry
}

func (csubsessiontable *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable) GetEntityData() *types.CommonEntityData {
    csubsessiontable.EntityData.YFilter = csubsessiontable.YFilter
    csubsessiontable.EntityData.YangName = "csubSessionTable"
    csubsessiontable.EntityData.BundleName = "cisco_ios_xe"
    csubsessiontable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubsessiontable.EntityData.SegmentPath = "csubSessionTable"
    csubsessiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubsessiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubsessiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubsessiontable.EntityData.Children = make(map[string]types.YChild)
    csubsessiontable.EntityData.Children["csubSessionEntry"] = types.YChild{"Csubsessionentry", nil}
    for i := range csubsessiontable.Csubsessionentry {
        csubsessiontable.EntityData.Children[types.GetSegmentPath(&csubsessiontable.Csubsessionentry[i])] = types.YChild{"Csubsessionentry", &csubsessiontable.Csubsessionentry[i]}
    }
    csubsessiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubsessiontable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

func (csubsessionentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessiontable_Csubsessionentry) GetEntityData() *types.CommonEntityData {
    csubsessionentry.EntityData.YFilter = csubsessionentry.YFilter
    csubsessionentry.EntityData.YangName = "csubSessionEntry"
    csubsessionentry.EntityData.BundleName = "cisco_ios_xe"
    csubsessionentry.EntityData.ParentYangName = "csubSessionTable"
    csubsessionentry.EntityData.SegmentPath = "csubSessionEntry" + "[ifIndex='" + fmt.Sprintf("%v", csubsessionentry.Ifindex) + "']"
    csubsessionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubsessionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubsessionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubsessionentry.EntityData.Children = make(map[string]types.YChild)
    csubsessionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubsessionentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", csubsessionentry.Ifindex}
    csubsessionentry.EntityData.Leafs["csubSessionType"] = types.YLeaf{"Csubsessiontype", csubsessionentry.Csubsessiontype}
    csubsessionentry.EntityData.Leafs["csubSessionIpAddrAssignment"] = types.YLeaf{"Csubsessionipaddrassignment", csubsessionentry.Csubsessionipaddrassignment}
    csubsessionentry.EntityData.Leafs["csubSessionState"] = types.YLeaf{"Csubsessionstate", csubsessionentry.Csubsessionstate}
    csubsessionentry.EntityData.Leafs["csubSessionAuthenticated"] = types.YLeaf{"Csubsessionauthenticated", csubsessionentry.Csubsessionauthenticated}
    csubsessionentry.EntityData.Leafs["csubSessionRedundancyMode"] = types.YLeaf{"Csubsessionredundancymode", csubsessionentry.Csubsessionredundancymode}
    csubsessionentry.EntityData.Leafs["csubSessionCreationTime"] = types.YLeaf{"Csubsessioncreationtime", csubsessionentry.Csubsessioncreationtime}
    csubsessionentry.EntityData.Leafs["csubSessionDerivedCfg"] = types.YLeaf{"Csubsessionderivedcfg", csubsessionentry.Csubsessionderivedcfg}
    csubsessionentry.EntityData.Leafs["csubSessionAvailableIdentities"] = types.YLeaf{"Csubsessionavailableidentities", csubsessionentry.Csubsessionavailableidentities}
    csubsessionentry.EntityData.Leafs["csubSessionSubscriberLabel"] = types.YLeaf{"Csubsessionsubscriberlabel", csubsessionentry.Csubsessionsubscriberlabel}
    csubsessionentry.EntityData.Leafs["csubSessionMacAddress"] = types.YLeaf{"Csubsessionmacaddress", csubsessionentry.Csubsessionmacaddress}
    csubsessionentry.EntityData.Leafs["csubSessionNativeVrf"] = types.YLeaf{"Csubsessionnativevrf", csubsessionentry.Csubsessionnativevrf}
    csubsessionentry.EntityData.Leafs["csubSessionNativeIpAddrType"] = types.YLeaf{"Csubsessionnativeipaddrtype", csubsessionentry.Csubsessionnativeipaddrtype}
    csubsessionentry.EntityData.Leafs["csubSessionNativeIpAddr"] = types.YLeaf{"Csubsessionnativeipaddr", csubsessionentry.Csubsessionnativeipaddr}
    csubsessionentry.EntityData.Leafs["csubSessionNativeIpMask"] = types.YLeaf{"Csubsessionnativeipmask", csubsessionentry.Csubsessionnativeipmask}
    csubsessionentry.EntityData.Leafs["csubSessionDomainVrf"] = types.YLeaf{"Csubsessiondomainvrf", csubsessionentry.Csubsessiondomainvrf}
    csubsessionentry.EntityData.Leafs["csubSessionDomainIpAddrType"] = types.YLeaf{"Csubsessiondomainipaddrtype", csubsessionentry.Csubsessiondomainipaddrtype}
    csubsessionentry.EntityData.Leafs["csubSessionDomainIpAddr"] = types.YLeaf{"Csubsessiondomainipaddr", csubsessionentry.Csubsessiondomainipaddr}
    csubsessionentry.EntityData.Leafs["csubSessionDomainIpMask"] = types.YLeaf{"Csubsessiondomainipmask", csubsessionentry.Csubsessiondomainipmask}
    csubsessionentry.EntityData.Leafs["csubSessionPbhk"] = types.YLeaf{"Csubsessionpbhk", csubsessionentry.Csubsessionpbhk}
    csubsessionentry.EntityData.Leafs["csubSessionRemoteId"] = types.YLeaf{"Csubsessionremoteid", csubsessionentry.Csubsessionremoteid}
    csubsessionentry.EntityData.Leafs["csubSessionCircuitId"] = types.YLeaf{"Csubsessioncircuitid", csubsessionentry.Csubsessioncircuitid}
    csubsessionentry.EntityData.Leafs["csubSessionNasPort"] = types.YLeaf{"Csubsessionnasport", csubsessionentry.Csubsessionnasport}
    csubsessionentry.EntityData.Leafs["csubSessionDomain"] = types.YLeaf{"Csubsessiondomain", csubsessionentry.Csubsessiondomain}
    csubsessionentry.EntityData.Leafs["csubSessionUsername"] = types.YLeaf{"Csubsessionusername", csubsessionentry.Csubsessionusername}
    csubsessionentry.EntityData.Leafs["csubSessionAcctSessionId"] = types.YLeaf{"Csubsessionacctsessionid", csubsessionentry.Csubsessionacctsessionid}
    csubsessionentry.EntityData.Leafs["csubSessionDnis"] = types.YLeaf{"Csubsessiondnis", csubsessionentry.Csubsessiondnis}
    csubsessionentry.EntityData.Leafs["csubSessionMedia"] = types.YLeaf{"Csubsessionmedia", csubsessionentry.Csubsessionmedia}
    csubsessionentry.EntityData.Leafs["csubSessionMlpNegotiated"] = types.YLeaf{"Csubsessionmlpnegotiated", csubsessionentry.Csubsessionmlpnegotiated}
    csubsessionentry.EntityData.Leafs["csubSessionProtocol"] = types.YLeaf{"Csubsessionprotocol", csubsessionentry.Csubsessionprotocol}
    csubsessionentry.EntityData.Leafs["csubSessionDhcpClass"] = types.YLeaf{"Csubsessiondhcpclass", csubsessionentry.Csubsessiondhcpclass}
    csubsessionentry.EntityData.Leafs["csubSessionTunnelName"] = types.YLeaf{"Csubsessiontunnelname", csubsessionentry.Csubsessiontunnelname}
    csubsessionentry.EntityData.Leafs["csubSessionLocationIdentifier"] = types.YLeaf{"Csubsessionlocationidentifier", csubsessionentry.Csubsessionlocationidentifier}
    csubsessionentry.EntityData.Leafs["csubSessionServiceIdentifier"] = types.YLeaf{"Csubsessionserviceidentifier", csubsessionentry.Csubsessionserviceidentifier}
    csubsessionentry.EntityData.Leafs["csubSessionLastChanged"] = types.YLeaf{"Csubsessionlastchanged", csubsessionentry.Csubsessionlastchanged}
    csubsessionentry.EntityData.Leafs["csubSessionNativeIpAddrType2"] = types.YLeaf{"Csubsessionnativeipaddrtype2", csubsessionentry.Csubsessionnativeipaddrtype2}
    csubsessionentry.EntityData.Leafs["csubSessionNativeIpAddr2"] = types.YLeaf{"Csubsessionnativeipaddr2", csubsessionentry.Csubsessionnativeipaddr2}
    csubsessionentry.EntityData.Leafs["csubSessionNativeIpMask2"] = types.YLeaf{"Csubsessionnativeipmask2", csubsessionentry.Csubsessionnativeipmask2}
    return &(csubsessionentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csubsessionbytypetable *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable) GetEntityData() *types.CommonEntityData {
    csubsessionbytypetable.EntityData.YFilter = csubsessionbytypetable.YFilter
    csubsessionbytypetable.EntityData.YangName = "csubSessionByTypeTable"
    csubsessionbytypetable.EntityData.BundleName = "cisco_ios_xe"
    csubsessionbytypetable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubsessionbytypetable.EntityData.SegmentPath = "csubSessionByTypeTable"
    csubsessionbytypetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubsessionbytypetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubsessionbytypetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubsessionbytypetable.EntityData.Children = make(map[string]types.YChild)
    csubsessionbytypetable.EntityData.Children["csubSessionByTypeEntry"] = types.YChild{"Csubsessionbytypeentry", nil}
    for i := range csubsessionbytypetable.Csubsessionbytypeentry {
        csubsessionbytypetable.EntityData.Children[types.GetSegmentPath(&csubsessionbytypetable.Csubsessionbytypeentry[i])] = types.YChild{"Csubsessionbytypeentry", &csubsessionbytypetable.Csubsessionbytypeentry[i]}
    }
    csubsessionbytypetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubsessionbytypetable.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates the type of the subscriber
    // session. The type is SubSessionType.
    Csubsessionbytype interface{}

    // This attribute is a key. This object indicates the ifIndex assigned to the
    // subscriber session. The type is interface{} with range: 1..2147483647.
    Csubsessionifindex interface{}
}

func (csubsessionbytypeentry *CISCOSUBSCRIBERSESSIONMIB_Csubsessionbytypetable_Csubsessionbytypeentry) GetEntityData() *types.CommonEntityData {
    csubsessionbytypeentry.EntityData.YFilter = csubsessionbytypeentry.YFilter
    csubsessionbytypeentry.EntityData.YangName = "csubSessionByTypeEntry"
    csubsessionbytypeentry.EntityData.BundleName = "cisco_ios_xe"
    csubsessionbytypeentry.EntityData.ParentYangName = "csubSessionByTypeTable"
    csubsessionbytypeentry.EntityData.SegmentPath = "csubSessionByTypeEntry" + "[csubSessionByType='" + fmt.Sprintf("%v", csubsessionbytypeentry.Csubsessionbytype) + "']" + "[csubSessionIfIndex='" + fmt.Sprintf("%v", csubsessionbytypeentry.Csubsessionifindex) + "']"
    csubsessionbytypeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubsessionbytypeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubsessionbytypeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubsessionbytypeentry.EntityData.Children = make(map[string]types.YChild)
    csubsessionbytypeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubsessionbytypeentry.EntityData.Leafs["csubSessionByType"] = types.YLeaf{"Csubsessionbytype", csubsessionbytypeentry.Csubsessionbytype}
    csubsessionbytypeentry.EntityData.Leafs["csubSessionIfIndex"] = types.YLeaf{"Csubsessionifindex", csubsessionbytypeentry.Csubsessionifindex}
    return &(csubsessionbytypeentry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable
// This table contains sets of aggregated statistics relating to
// subscriber sessions, where each set has a unique scope of
// aggregation.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable struct {
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
    // slice of CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry.
    Csubaggstatsentry []CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry
}

func (csubaggstatstable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable) GetEntityData() *types.CommonEntityData {
    csubaggstatstable.EntityData.YFilter = csubaggstatstable.YFilter
    csubaggstatstable.EntityData.YangName = "csubAggStatsTable"
    csubaggstatstable.EntityData.BundleName = "cisco_ios_xe"
    csubaggstatstable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubaggstatstable.EntityData.SegmentPath = "csubAggStatsTable"
    csubaggstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubaggstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubaggstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubaggstatstable.EntityData.Children = make(map[string]types.YChild)
    csubaggstatstable.EntityData.Children["csubAggStatsEntry"] = types.YChild{"Csubaggstatsentry", nil}
    for i := range csubaggstatstable.Csubaggstatsentry {
        csubaggstatstable.EntityData.Children[types.GetSegmentPath(&csubaggstatstable.Csubaggstatsentry[i])] = types.YChild{"Csubaggstatsentry", &csubaggstatstable.Csubaggstatsentry[i]}
    }
    csubaggstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubaggstatstable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csubaggstatsentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatstable_Csubaggstatsentry) GetEntityData() *types.CommonEntityData {
    csubaggstatsentry.EntityData.YFilter = csubaggstatsentry.YFilter
    csubaggstatsentry.EntityData.YangName = "csubAggStatsEntry"
    csubaggstatsentry.EntityData.BundleName = "cisco_ios_xe"
    csubaggstatsentry.EntityData.ParentYangName = "csubAggStatsTable"
    csubaggstatsentry.EntityData.SegmentPath = "csubAggStatsEntry" + "[csubAggStatsPointType='" + fmt.Sprintf("%v", csubaggstatsentry.Csubaggstatspointtype) + "']" + "[csubAggStatsPoint='" + fmt.Sprintf("%v", csubaggstatsentry.Csubaggstatspoint) + "']" + "[csubAggStatsSessionType='" + fmt.Sprintf("%v", csubaggstatsentry.Csubaggstatssessiontype) + "']"
    csubaggstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubaggstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubaggstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubaggstatsentry.EntityData.Children = make(map[string]types.YChild)
    csubaggstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubaggstatsentry.EntityData.Leafs["csubAggStatsPointType"] = types.YLeaf{"Csubaggstatspointtype", csubaggstatsentry.Csubaggstatspointtype}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsPoint"] = types.YLeaf{"Csubaggstatspoint", csubaggstatsentry.Csubaggstatspoint}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsSessionType"] = types.YLeaf{"Csubaggstatssessiontype", csubaggstatsentry.Csubaggstatssessiontype}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsPendingSessions"] = types.YLeaf{"Csubaggstatspendingsessions", csubaggstatsentry.Csubaggstatspendingsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsUpSessions"] = types.YLeaf{"Csubaggstatsupsessions", csubaggstatsentry.Csubaggstatsupsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsAuthSessions"] = types.YLeaf{"Csubaggstatsauthsessions", csubaggstatsentry.Csubaggstatsauthsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsUnAuthSessions"] = types.YLeaf{"Csubaggstatsunauthsessions", csubaggstatsentry.Csubaggstatsunauthsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsLightWeightSessions"] = types.YLeaf{"Csubaggstatslightweightsessions", csubaggstatsentry.Csubaggstatslightweightsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsRedSessions"] = types.YLeaf{"Csubaggstatsredsessions", csubaggstatsentry.Csubaggstatsredsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsHighUpSessions"] = types.YLeaf{"Csubaggstatshighupsessions", csubaggstatsentry.Csubaggstatshighupsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsAvgSessionUptime"] = types.YLeaf{"Csubaggstatsavgsessionuptime", csubaggstatsentry.Csubaggstatsavgsessionuptime}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsAvgSessionRPM"] = types.YLeaf{"Csubaggstatsavgsessionrpm", csubaggstatsentry.Csubaggstatsavgsessionrpm}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsAvgSessionRPH"] = types.YLeaf{"Csubaggstatsavgsessionrph", csubaggstatsentry.Csubaggstatsavgsessionrph}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsThrottleEngagements"] = types.YLeaf{"Csubaggstatsthrottleengagements", csubaggstatsentry.Csubaggstatsthrottleengagements}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsTotalCreatedSessions"] = types.YLeaf{"Csubaggstatstotalcreatedsessions", csubaggstatsentry.Csubaggstatstotalcreatedsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsTotalFailedSessions"] = types.YLeaf{"Csubaggstatstotalfailedsessions", csubaggstatsentry.Csubaggstatstotalfailedsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsTotalUpSessions"] = types.YLeaf{"Csubaggstatstotalupsessions", csubaggstatsentry.Csubaggstatstotalupsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsTotalAuthSessions"] = types.YLeaf{"Csubaggstatstotalauthsessions", csubaggstatsentry.Csubaggstatstotalauthsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsTotalDiscSessions"] = types.YLeaf{"Csubaggstatstotaldiscsessions", csubaggstatsentry.Csubaggstatstotaldiscsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsTotalLightWeightSessions"] = types.YLeaf{"Csubaggstatstotallightweightsessions", csubaggstatsentry.Csubaggstatstotallightweightsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsTotalFlowsUp"] = types.YLeaf{"Csubaggstatstotalflowsup", csubaggstatsentry.Csubaggstatstotalflowsup}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsDayCreatedSessions"] = types.YLeaf{"Csubaggstatsdaycreatedsessions", csubaggstatsentry.Csubaggstatsdaycreatedsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsDayFailedSessions"] = types.YLeaf{"Csubaggstatsdayfailedsessions", csubaggstatsentry.Csubaggstatsdayfailedsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsDayUpSessions"] = types.YLeaf{"Csubaggstatsdayupsessions", csubaggstatsentry.Csubaggstatsdayupsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsDayAuthSessions"] = types.YLeaf{"Csubaggstatsdayauthsessions", csubaggstatsentry.Csubaggstatsdayauthsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsDayDiscSessions"] = types.YLeaf{"Csubaggstatsdaydiscsessions", csubaggstatsentry.Csubaggstatsdaydiscsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrTimeElapsed"] = types.YLeaf{"Csubaggstatscurrtimeelapsed", csubaggstatsentry.Csubaggstatscurrtimeelapsed}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrValidIntervals"] = types.YLeaf{"Csubaggstatscurrvalidintervals", csubaggstatsentry.Csubaggstatscurrvalidintervals}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrInvalidIntervals"] = types.YLeaf{"Csubaggstatscurrinvalidintervals", csubaggstatsentry.Csubaggstatscurrinvalidintervals}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrFlowsUp"] = types.YLeaf{"Csubaggstatscurrflowsup", csubaggstatsentry.Csubaggstatscurrflowsup}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrCreatedSessions"] = types.YLeaf{"Csubaggstatscurrcreatedsessions", csubaggstatsentry.Csubaggstatscurrcreatedsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrFailedSessions"] = types.YLeaf{"Csubaggstatscurrfailedsessions", csubaggstatsentry.Csubaggstatscurrfailedsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrUpSessions"] = types.YLeaf{"Csubaggstatscurrupsessions", csubaggstatsentry.Csubaggstatscurrupsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrAuthSessions"] = types.YLeaf{"Csubaggstatscurrauthsessions", csubaggstatsentry.Csubaggstatscurrauthsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsCurrDiscSessions"] = types.YLeaf{"Csubaggstatscurrdiscsessions", csubaggstatsentry.Csubaggstatscurrdiscsessions}
    csubaggstatsentry.EntityData.Leafs["csubAggStatsDiscontinuityTime"] = types.YLeaf{"Csubaggstatsdiscontinuitytime", csubaggstatsentry.Csubaggstatsdiscontinuitytime}
    return &(csubaggstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry contains the aggregated subscriber session performance data
    // collected over a single 15-minute measurement interval within a 'scope of
    // aggregation'.  For further details regarding 'scope of aggregation', see
    // the descriptive text associated with the csubAggStatsEntry. The type is
    // slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry.
    Csubaggstatsintentry []CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry
}

func (csubaggstatsinttable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable) GetEntityData() *types.CommonEntityData {
    csubaggstatsinttable.EntityData.YFilter = csubaggstatsinttable.YFilter
    csubaggstatsinttable.EntityData.YangName = "csubAggStatsIntTable"
    csubaggstatsinttable.EntityData.BundleName = "cisco_ios_xe"
    csubaggstatsinttable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubaggstatsinttable.EntityData.SegmentPath = "csubAggStatsIntTable"
    csubaggstatsinttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubaggstatsinttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubaggstatsinttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubaggstatsinttable.EntityData.Children = make(map[string]types.YChild)
    csubaggstatsinttable.EntityData.Children["csubAggStatsIntEntry"] = types.YChild{"Csubaggstatsintentry", nil}
    for i := range csubaggstatsinttable.Csubaggstatsintentry {
        csubaggstatsinttable.EntityData.Children[types.GetSegmentPath(&csubaggstatsinttable.Csubaggstatsintentry[i])] = types.YChild{"Csubaggstatsintentry", &csubaggstatsinttable.Csubaggstatsintentry[i]}
    }
    csubaggstatsinttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubaggstatsinttable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry
// An entry contains the aggregated subscriber session performance
// data collected over a single 15-minute measurement interval
// within a 'scope of aggregation'.  For further details regarding
// 'scope of aggregation', see the descriptive text associated with
// the csubAggStatsEntry.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry struct {
    EntityData types.CommonEntityData
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

func (csubaggstatsintentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsinttable_Csubaggstatsintentry) GetEntityData() *types.CommonEntityData {
    csubaggstatsintentry.EntityData.YFilter = csubaggstatsintentry.YFilter
    csubaggstatsintentry.EntityData.YangName = "csubAggStatsIntEntry"
    csubaggstatsintentry.EntityData.BundleName = "cisco_ios_xe"
    csubaggstatsintentry.EntityData.ParentYangName = "csubAggStatsIntTable"
    csubaggstatsintentry.EntityData.SegmentPath = "csubAggStatsIntEntry" + "[csubAggStatsPointType='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatspointtype) + "']" + "[csubAggStatsPoint='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatspoint) + "']" + "[csubAggStatsSessionType='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatssessiontype) + "']" + "[csubAggStatsIntNumber='" + fmt.Sprintf("%v", csubaggstatsintentry.Csubaggstatsintnumber) + "']"
    csubaggstatsintentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubaggstatsintentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubaggstatsintentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubaggstatsintentry.EntityData.Children = make(map[string]types.YChild)
    csubaggstatsintentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsPointType"] = types.YLeaf{"Csubaggstatspointtype", csubaggstatsintentry.Csubaggstatspointtype}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsPoint"] = types.YLeaf{"Csubaggstatspoint", csubaggstatsintentry.Csubaggstatspoint}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsSessionType"] = types.YLeaf{"Csubaggstatssessiontype", csubaggstatsintentry.Csubaggstatssessiontype}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsIntNumber"] = types.YLeaf{"Csubaggstatsintnumber", csubaggstatsintentry.Csubaggstatsintnumber}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsIntValid"] = types.YLeaf{"Csubaggstatsintvalid", csubaggstatsintentry.Csubaggstatsintvalid}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsIntCreatedSessions"] = types.YLeaf{"Csubaggstatsintcreatedsessions", csubaggstatsintentry.Csubaggstatsintcreatedsessions}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsIntFailedSessions"] = types.YLeaf{"Csubaggstatsintfailedsessions", csubaggstatsintentry.Csubaggstatsintfailedsessions}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsIntUpSessions"] = types.YLeaf{"Csubaggstatsintupsessions", csubaggstatsintentry.Csubaggstatsintupsessions}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsIntAuthSessions"] = types.YLeaf{"Csubaggstatsintauthsessions", csubaggstatsintentry.Csubaggstatsintauthsessions}
    csubaggstatsintentry.EntityData.Leafs["csubAggStatsIntDiscSessions"] = types.YLeaf{"Csubaggstatsintdiscsessions", csubaggstatsintentry.Csubaggstatsintdiscsessions}
    return &(csubaggstatsintentry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable
// Please enter the Table Description here.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table exists for each row in the csubAggStatsTable. Each row
    // defines the set of thresholds and evaluation attributes for an aggregation
    // point. The type is slice of
    // CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry.
    Csubaggstatsthreshentry []CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry
}

func (csubaggstatsthreshtable *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable) GetEntityData() *types.CommonEntityData {
    csubaggstatsthreshtable.EntityData.YFilter = csubaggstatsthreshtable.YFilter
    csubaggstatsthreshtable.EntityData.YangName = "csubAggStatsThreshTable"
    csubaggstatsthreshtable.EntityData.BundleName = "cisco_ios_xe"
    csubaggstatsthreshtable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubaggstatsthreshtable.EntityData.SegmentPath = "csubAggStatsThreshTable"
    csubaggstatsthreshtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubaggstatsthreshtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubaggstatsthreshtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubaggstatsthreshtable.EntityData.Children = make(map[string]types.YChild)
    csubaggstatsthreshtable.EntityData.Children["csubAggStatsThreshEntry"] = types.YChild{"Csubaggstatsthreshentry", nil}
    for i := range csubaggstatsthreshtable.Csubaggstatsthreshentry {
        csubaggstatsthreshtable.EntityData.Children[types.GetSegmentPath(&csubaggstatsthreshtable.Csubaggstatsthreshentry[i])] = types.YChild{"Csubaggstatsthreshentry", &csubaggstatsthreshtable.Csubaggstatsthreshentry[i]}
    }
    csubaggstatsthreshtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubaggstatsthreshtable.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry
// A row in this table exists for each row in the
// csubAggStatsTable.
// Each row defines the set of thresholds and evaluation attributes
// for an aggregation point.
type CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry struct {
    EntityData types.CommonEntityData
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

func (csubaggstatsthreshentry *CISCOSUBSCRIBERSESSIONMIB_Csubaggstatsthreshtable_Csubaggstatsthreshentry) GetEntityData() *types.CommonEntityData {
    csubaggstatsthreshentry.EntityData.YFilter = csubaggstatsthreshentry.YFilter
    csubaggstatsthreshentry.EntityData.YangName = "csubAggStatsThreshEntry"
    csubaggstatsthreshentry.EntityData.BundleName = "cisco_ios_xe"
    csubaggstatsthreshentry.EntityData.ParentYangName = "csubAggStatsThreshTable"
    csubaggstatsthreshentry.EntityData.SegmentPath = "csubAggStatsThreshEntry" + "[csubSessionRisingThresh='" + fmt.Sprintf("%v", csubaggstatsthreshentry.Csubsessionrisingthresh) + "']"
    csubaggstatsthreshentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubaggstatsthreshentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubaggstatsthreshentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubaggstatsthreshentry.EntityData.Children = make(map[string]types.YChild)
    csubaggstatsthreshentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubaggstatsthreshentry.EntityData.Leafs["csubSessionRisingThresh"] = types.YLeaf{"Csubsessionrisingthresh", csubaggstatsthreshentry.Csubsessionrisingthresh}
    csubaggstatsthreshentry.EntityData.Leafs["csubSessionFallingThresh"] = types.YLeaf{"Csubsessionfallingthresh", csubaggstatsthreshentry.Csubsessionfallingthresh}
    csubaggstatsthreshentry.EntityData.Leafs["csubSessionDeltaPercentFallingThresh"] = types.YLeaf{"Csubsessiondeltapercentfallingthresh", csubaggstatsthreshentry.Csubsessiondeltapercentfallingthresh}
    csubaggstatsthreshentry.EntityData.Leafs["csubSessionThreshEvalInterval"] = types.YLeaf{"Csubsessionthreshevalinterval", csubaggstatsthreshentry.Csubsessionthreshevalinterval}
    return &(csubaggstatsthreshentry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubjobtable
// This table contains the subscriber session jobs submitted by
// the EMS/NMS.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobtable struct {
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
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry.
    Csubjobentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry
}

func (csubjobtable *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable) GetEntityData() *types.CommonEntityData {
    csubjobtable.EntityData.YFilter = csubjobtable.YFilter
    csubjobtable.EntityData.YangName = "csubJobTable"
    csubjobtable.EntityData.BundleName = "cisco_ios_xe"
    csubjobtable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubjobtable.EntityData.SegmentPath = "csubJobTable"
    csubjobtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobtable.EntityData.Children = make(map[string]types.YChild)
    csubjobtable.EntityData.Children["csubJobEntry"] = types.YChild{"Csubjobentry", nil}
    for i := range csubjobtable.Csubjobentry {
        csubjobtable.EntityData.Children[types.GetSegmentPath(&csubjobtable.Csubjobentry[i])] = types.YChild{"Csubjobentry", &csubjobtable.Csubjobentry[i]}
    }
    csubjobtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubjobtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csubjobentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobtable_Csubjobentry) GetEntityData() *types.CommonEntityData {
    csubjobentry.EntityData.YFilter = csubjobentry.YFilter
    csubjobentry.EntityData.YangName = "csubJobEntry"
    csubjobentry.EntityData.BundleName = "cisco_ios_xe"
    csubjobentry.EntityData.ParentYangName = "csubJobTable"
    csubjobentry.EntityData.SegmentPath = "csubJobEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobentry.Csubjobid) + "']"
    csubjobentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobentry.EntityData.Children = make(map[string]types.YChild)
    csubjobentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubjobentry.EntityData.Leafs["csubJobId"] = types.YLeaf{"Csubjobid", csubjobentry.Csubjobid}
    csubjobentry.EntityData.Leafs["csubJobStatus"] = types.YLeaf{"Csubjobstatus", csubjobentry.Csubjobstatus}
    csubjobentry.EntityData.Leafs["csubJobStorage"] = types.YLeaf{"Csubjobstorage", csubjobentry.Csubjobstorage}
    csubjobentry.EntityData.Leafs["csubJobType"] = types.YLeaf{"Csubjobtype", csubjobentry.Csubjobtype}
    csubjobentry.EntityData.Leafs["csubJobControl"] = types.YLeaf{"Csubjobcontrol", csubjobentry.Csubjobcontrol}
    csubjobentry.EntityData.Leafs["csubJobState"] = types.YLeaf{"Csubjobstate", csubjobentry.Csubjobstate}
    csubjobentry.EntityData.Leafs["csubJobStartedTime"] = types.YLeaf{"Csubjobstartedtime", csubjobentry.Csubjobstartedtime}
    csubjobentry.EntityData.Leafs["csubJobFinishedTime"] = types.YLeaf{"Csubjobfinishedtime", csubjobentry.Csubjobfinishedtime}
    csubjobentry.EntityData.Leafs["csubJobFinishedReason"] = types.YLeaf{"Csubjobfinishedreason", csubjobentry.Csubjobfinishedreason}
    return &(csubjobentry.EntityData)
}

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
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry.
    Csubjobmatchparamsentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry
}

func (csubjobmatchparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable) GetEntityData() *types.CommonEntityData {
    csubjobmatchparamstable.EntityData.YFilter = csubjobmatchparamstable.YFilter
    csubjobmatchparamstable.EntityData.YangName = "csubJobMatchParamsTable"
    csubjobmatchparamstable.EntityData.BundleName = "cisco_ios_xe"
    csubjobmatchparamstable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubjobmatchparamstable.EntityData.SegmentPath = "csubJobMatchParamsTable"
    csubjobmatchparamstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobmatchparamstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobmatchparamstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobmatchparamstable.EntityData.Children = make(map[string]types.YChild)
    csubjobmatchparamstable.EntityData.Children["csubJobMatchParamsEntry"] = types.YChild{"Csubjobmatchparamsentry", nil}
    for i := range csubjobmatchparamstable.Csubjobmatchparamsentry {
        csubjobmatchparamstable.EntityData.Children[types.GetSegmentPath(&csubjobmatchparamstable.Csubjobmatchparamsentry[i])] = types.YChild{"Csubjobmatchparamsentry", &csubjobmatchparamstable.Csubjobmatchparamsentry[i]}
    }
    csubjobmatchparamstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubjobmatchparamstable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

func (csubjobmatchparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobmatchparamstable_Csubjobmatchparamsentry) GetEntityData() *types.CommonEntityData {
    csubjobmatchparamsentry.EntityData.YFilter = csubjobmatchparamsentry.YFilter
    csubjobmatchparamsentry.EntityData.YangName = "csubJobMatchParamsEntry"
    csubjobmatchparamsentry.EntityData.BundleName = "cisco_ios_xe"
    csubjobmatchparamsentry.EntityData.ParentYangName = "csubJobMatchParamsTable"
    csubjobmatchparamsentry.EntityData.SegmentPath = "csubJobMatchParamsEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobmatchparamsentry.Csubjobid) + "']"
    csubjobmatchparamsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobmatchparamsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobmatchparamsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobmatchparamsentry.EntityData.Children = make(map[string]types.YChild)
    csubjobmatchparamsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubjobmatchparamsentry.EntityData.Leafs["csubJobId"] = types.YLeaf{"Csubjobid", csubjobmatchparamsentry.Csubjobid}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchIdentities"] = types.YLeaf{"Csubjobmatchidentities", csubjobmatchparamsentry.Csubjobmatchidentities}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchOtherParams"] = types.YLeaf{"Csubjobmatchotherparams", csubjobmatchparamsentry.Csubjobmatchotherparams}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchSubscriberLabel"] = types.YLeaf{"Csubjobmatchsubscriberlabel", csubjobmatchparamsentry.Csubjobmatchsubscriberlabel}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchMacAddress"] = types.YLeaf{"Csubjobmatchmacaddress", csubjobmatchparamsentry.Csubjobmatchmacaddress}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchNativeVrf"] = types.YLeaf{"Csubjobmatchnativevrf", csubjobmatchparamsentry.Csubjobmatchnativevrf}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchNativeIpAddrType"] = types.YLeaf{"Csubjobmatchnativeipaddrtype", csubjobmatchparamsentry.Csubjobmatchnativeipaddrtype}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchNativeIpAddr"] = types.YLeaf{"Csubjobmatchnativeipaddr", csubjobmatchparamsentry.Csubjobmatchnativeipaddr}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchNativeIpMask"] = types.YLeaf{"Csubjobmatchnativeipmask", csubjobmatchparamsentry.Csubjobmatchnativeipmask}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDomainVrf"] = types.YLeaf{"Csubjobmatchdomainvrf", csubjobmatchparamsentry.Csubjobmatchdomainvrf}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDomainIpAddrType"] = types.YLeaf{"Csubjobmatchdomainipaddrtype", csubjobmatchparamsentry.Csubjobmatchdomainipaddrtype}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDomainIpAddr"] = types.YLeaf{"Csubjobmatchdomainipaddr", csubjobmatchparamsentry.Csubjobmatchdomainipaddr}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDomainIpMask"] = types.YLeaf{"Csubjobmatchdomainipmask", csubjobmatchparamsentry.Csubjobmatchdomainipmask}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchPbhk"] = types.YLeaf{"Csubjobmatchpbhk", csubjobmatchparamsentry.Csubjobmatchpbhk}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchRemoteId"] = types.YLeaf{"Csubjobmatchremoteid", csubjobmatchparamsentry.Csubjobmatchremoteid}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchCircuitId"] = types.YLeaf{"Csubjobmatchcircuitid", csubjobmatchparamsentry.Csubjobmatchcircuitid}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchNasPort"] = types.YLeaf{"Csubjobmatchnasport", csubjobmatchparamsentry.Csubjobmatchnasport}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDomain"] = types.YLeaf{"Csubjobmatchdomain", csubjobmatchparamsentry.Csubjobmatchdomain}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchUsername"] = types.YLeaf{"Csubjobmatchusername", csubjobmatchparamsentry.Csubjobmatchusername}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchAcctSessionId"] = types.YLeaf{"Csubjobmatchacctsessionid", csubjobmatchparamsentry.Csubjobmatchacctsessionid}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDnis"] = types.YLeaf{"Csubjobmatchdnis", csubjobmatchparamsentry.Csubjobmatchdnis}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchMedia"] = types.YLeaf{"Csubjobmatchmedia", csubjobmatchparamsentry.Csubjobmatchmedia}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchMlpNegotiated"] = types.YLeaf{"Csubjobmatchmlpnegotiated", csubjobmatchparamsentry.Csubjobmatchmlpnegotiated}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchProtocol"] = types.YLeaf{"Csubjobmatchprotocol", csubjobmatchparamsentry.Csubjobmatchprotocol}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchServiceName"] = types.YLeaf{"Csubjobmatchservicename", csubjobmatchparamsentry.Csubjobmatchservicename}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDhcpClass"] = types.YLeaf{"Csubjobmatchdhcpclass", csubjobmatchparamsentry.Csubjobmatchdhcpclass}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchTunnelName"] = types.YLeaf{"Csubjobmatchtunnelname", csubjobmatchparamsentry.Csubjobmatchtunnelname}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchDanglingDuration"] = types.YLeaf{"Csubjobmatchdanglingduration", csubjobmatchparamsentry.Csubjobmatchdanglingduration}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchState"] = types.YLeaf{"Csubjobmatchstate", csubjobmatchparamsentry.Csubjobmatchstate}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchAuthenticated"] = types.YLeaf{"Csubjobmatchauthenticated", csubjobmatchparamsentry.Csubjobmatchauthenticated}
    csubjobmatchparamsentry.EntityData.Leafs["csubJobMatchRedundancyMode"] = types.YLeaf{"Csubjobmatchredundancymode", csubjobmatchparamsentry.Csubjobmatchredundancymode}
    return &(csubjobmatchparamsentry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable
// This table contains subscriber session job parameters
// describing query parameters.
// 
// This table has a sparse-dependent relationship on the
// csubJobTable, containing a row for each job having a
// csubJobType of 'query'.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable struct {
    EntityData types.CommonEntityData
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

func (csubjobqueryparamstable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable) GetEntityData() *types.CommonEntityData {
    csubjobqueryparamstable.EntityData.YFilter = csubjobqueryparamstable.YFilter
    csubjobqueryparamstable.EntityData.YangName = "csubJobQueryParamsTable"
    csubjobqueryparamstable.EntityData.BundleName = "cisco_ios_xe"
    csubjobqueryparamstable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubjobqueryparamstable.EntityData.SegmentPath = "csubJobQueryParamsTable"
    csubjobqueryparamstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobqueryparamstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobqueryparamstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobqueryparamstable.EntityData.Children = make(map[string]types.YChild)
    csubjobqueryparamstable.EntityData.Children["csubJobQueryParamsEntry"] = types.YChild{"Csubjobqueryparamsentry", nil}
    for i := range csubjobqueryparamstable.Csubjobqueryparamsentry {
        csubjobqueryparamstable.EntityData.Children[types.GetSegmentPath(&csubjobqueryparamstable.Csubjobqueryparamsentry[i])] = types.YChild{"Csubjobqueryparamsentry", &csubjobqueryparamstable.Csubjobqueryparamsentry[i]}
    }
    csubjobqueryparamstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubjobqueryparamstable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csubjobqueryparamsentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueryparamstable_Csubjobqueryparamsentry) GetEntityData() *types.CommonEntityData {
    csubjobqueryparamsentry.EntityData.YFilter = csubjobqueryparamsentry.YFilter
    csubjobqueryparamsentry.EntityData.YangName = "csubJobQueryParamsEntry"
    csubjobqueryparamsentry.EntityData.BundleName = "cisco_ios_xe"
    csubjobqueryparamsentry.EntityData.ParentYangName = "csubJobQueryParamsTable"
    csubjobqueryparamsentry.EntityData.SegmentPath = "csubJobQueryParamsEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobqueryparamsentry.Csubjobid) + "']"
    csubjobqueryparamsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobqueryparamsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobqueryparamsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobqueryparamsentry.EntityData.Children = make(map[string]types.YChild)
    csubjobqueryparamsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubjobqueryparamsentry.EntityData.Leafs["csubJobId"] = types.YLeaf{"Csubjobid", csubjobqueryparamsentry.Csubjobid}
    csubjobqueryparamsentry.EntityData.Leafs["csubJobQuerySortKey1"] = types.YLeaf{"Csubjobquerysortkey1", csubjobqueryparamsentry.Csubjobquerysortkey1}
    csubjobqueryparamsentry.EntityData.Leafs["csubJobQuerySortKey2"] = types.YLeaf{"Csubjobquerysortkey2", csubjobqueryparamsentry.Csubjobquerysortkey2}
    csubjobqueryparamsentry.EntityData.Leafs["csubJobQuerySortKey3"] = types.YLeaf{"Csubjobquerysortkey3", csubjobqueryparamsentry.Csubjobquerysortkey3}
    csubjobqueryparamsentry.EntityData.Leafs["csubJobQueryResultingReportSize"] = types.YLeaf{"Csubjobqueryresultingreportsize", csubjobqueryparamsentry.Csubjobqueryresultingreportsize}
    return &(csubjobqueryparamsentry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable
// This table lists the subscriber session jobs currently pending
// in the subscriber session job queue.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable struct {
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
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry.
    Csubjobqueueentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry
}

func (csubjobqueuetable *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable) GetEntityData() *types.CommonEntityData {
    csubjobqueuetable.EntityData.YFilter = csubjobqueuetable.YFilter
    csubjobqueuetable.EntityData.YangName = "csubJobQueueTable"
    csubjobqueuetable.EntityData.BundleName = "cisco_ios_xe"
    csubjobqueuetable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubjobqueuetable.EntityData.SegmentPath = "csubJobQueueTable"
    csubjobqueuetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobqueuetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobqueuetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobqueuetable.EntityData.Children = make(map[string]types.YChild)
    csubjobqueuetable.EntityData.Children["csubJobQueueEntry"] = types.YChild{"Csubjobqueueentry", nil}
    for i := range csubjobqueuetable.Csubjobqueueentry {
        csubjobqueuetable.EntityData.Children[types.GetSegmentPath(&csubjobqueuetable.Csubjobqueueentry[i])] = types.YChild{"Csubjobqueueentry", &csubjobqueuetable.Csubjobqueueentry[i]}
    }
    csubjobqueuetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubjobqueuetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csubjobqueueentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobqueuetable_Csubjobqueueentry) GetEntityData() *types.CommonEntityData {
    csubjobqueueentry.EntityData.YFilter = csubjobqueueentry.YFilter
    csubjobqueueentry.EntityData.YangName = "csubJobQueueEntry"
    csubjobqueueentry.EntityData.BundleName = "cisco_ios_xe"
    csubjobqueueentry.EntityData.ParentYangName = "csubJobQueueTable"
    csubjobqueueentry.EntityData.SegmentPath = "csubJobQueueEntry" + "[csubJobQueueNumber='" + fmt.Sprintf("%v", csubjobqueueentry.Csubjobqueuenumber) + "']"
    csubjobqueueentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobqueueentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobqueueentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobqueueentry.EntityData.Children = make(map[string]types.YChild)
    csubjobqueueentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubjobqueueentry.EntityData.Leafs["csubJobQueueNumber"] = types.YLeaf{"Csubjobqueuenumber", csubjobqueueentry.Csubjobqueuenumber}
    csubjobqueueentry.EntityData.Leafs["csubJobQueueJobId"] = types.YLeaf{"Csubjobqueuejobid", csubjobqueueentry.Csubjobqueuejobid}
    return &(csubjobqueueentry.EntityData)
}

// CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable
// This table contains the reports corresponding to subscriber
// session jobs that have a csubJobType of 'query' and
// csubJobState of 'finished'.
// 
// This table has an expansion dependent relationship on the
// csubJobTable, containing zero or more rows for each job.
type CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable struct {
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
    // CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry.
    Csubjobreportentry []CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry
}

func (csubjobreporttable *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable) GetEntityData() *types.CommonEntityData {
    csubjobreporttable.EntityData.YFilter = csubjobreporttable.YFilter
    csubjobreporttable.EntityData.YangName = "csubJobReportTable"
    csubjobreporttable.EntityData.BundleName = "cisco_ios_xe"
    csubjobreporttable.EntityData.ParentYangName = "CISCO-SUBSCRIBER-SESSION-MIB"
    csubjobreporttable.EntityData.SegmentPath = "csubJobReportTable"
    csubjobreporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobreporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobreporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobreporttable.EntityData.Children = make(map[string]types.YChild)
    csubjobreporttable.EntityData.Children["csubJobReportEntry"] = types.YChild{"Csubjobreportentry", nil}
    for i := range csubjobreporttable.Csubjobreportentry {
        csubjobreporttable.EntityData.Children[types.GetSegmentPath(&csubjobreporttable.Csubjobreportentry[i])] = types.YChild{"Csubjobreportentry", &csubjobreporttable.Csubjobreportentry[i]}
    }
    csubjobreporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csubjobreporttable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (csubjobreportentry *CISCOSUBSCRIBERSESSIONMIB_Csubjobreporttable_Csubjobreportentry) GetEntityData() *types.CommonEntityData {
    csubjobreportentry.EntityData.YFilter = csubjobreportentry.YFilter
    csubjobreportentry.EntityData.YangName = "csubJobReportEntry"
    csubjobreportentry.EntityData.BundleName = "cisco_ios_xe"
    csubjobreportentry.EntityData.ParentYangName = "csubJobReportTable"
    csubjobreportentry.EntityData.SegmentPath = "csubJobReportEntry" + "[csubJobId='" + fmt.Sprintf("%v", csubjobreportentry.Csubjobid) + "']" + "[csubJobReportId='" + fmt.Sprintf("%v", csubjobreportentry.Csubjobreportid) + "']"
    csubjobreportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    csubjobreportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    csubjobreportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    csubjobreportentry.EntityData.Children = make(map[string]types.YChild)
    csubjobreportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    csubjobreportentry.EntityData.Leafs["csubJobId"] = types.YLeaf{"Csubjobid", csubjobreportentry.Csubjobid}
    csubjobreportentry.EntityData.Leafs["csubJobReportId"] = types.YLeaf{"Csubjobreportid", csubjobreportentry.Csubjobreportid}
    csubjobreportentry.EntityData.Leafs["csubJobReportSession"] = types.YLeaf{"Csubjobreportsession", csubjobreportentry.Csubjobreportsession}
    return &(csubjobreportentry.EntityData)
}

