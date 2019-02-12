// Enhanced Interior Gateway Protocol (EIGRP) is a Cisco
// proprietary distance vector routing protocol.   It is based on
// the Diffusing Update Algorithm (DUAL), which is a method of
// finding loop-free paths through a network.   Directly
// connected routers running EIGRP form neighbor adjacencies in
// order to propagate best-path and alternate-path routing
// information for configured and learned routes.
// 
// The tables defined within the MIB are closely aligned with how
// the router command-line interface for EIGRP displays
// information on EIGRP configurations, i.e., the topology table
// contains objects associated with the EIGRP topology commands,
// and the peer table contains objects associated withe EIGRP
// neighbor commands, etc.
// 
// There are five main tables within this mib:
// 
//    EIGRP VPN table
//       Contains information regarding which virtual private
//       networks (VPN) are configured with EIGRP.
// 
//    EIGRP traffic statistics table
//       Contains counter & statistcs regarding specific types of
//       EIGRP packets sent and related collective information
//       per VPN and per autonomous system (AS).
// 
//    EIGRP topology table
//       Contains information regarding EIGRP routes received in
//       updates and originated locally.   EIGRP sends and
//       receives routing updates from adjacent routers running
//       EIGRP with which it formed a peer relationship.
// 
//    EIGRP peer (neighbor) table
//       Contains information about neighbor EIGRP routers with
//       which peer adjacencies have been established.   EIGRP
//       uses a Hello protocol to form neighbor relationships
//       with directly connected routers also running EIGRP.
// 
//    EIGRP interfaces table
//       Contains information and statistics on each of the
//       interfaces on the router over which EIGRP has been
//       configured to run.
package cisco_eigrp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_eigrp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-EIGRP-MIB CISCO-EIGRP-MIB}", reflect.TypeOf(CISCOEIGRPMIB{}))
    ydk.RegisterEntity("CISCO-EIGRP-MIB:CISCO-EIGRP-MIB", reflect.TypeOf(CISCOEIGRPMIB{}))
}

// CISCOEIGRPMIB
type CISCOEIGRPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains information on those VPN's configured to run EIGRP. 
    // The VPN creation on a router is independent of the routing protocol to be
    // used over it.   A VPN is given a name and has a dedicated routing table
    // associated with it.  This routing table is identified internally by a
    // unique integer value.
    CEigrpVpnTable CISCOEIGRPMIB_CEigrpVpnTable

    // Table of EIGRP traffic statistics and information associated with all EIGRP
    // autonomous systems.
    CEigrpTraffStatsTable CISCOEIGRPMIB_CEigrpTraffStatsTable

    // The table of EIGRP routes and their associated attributes for an Autonomous
    // System (AS) configured in a VPN is called a topology table.  All route
    // entries in the topology table will be indexed by IP network type, IP
    // network number and network mask (prefix) size.
    CEigrpTopoTable CISCOEIGRPMIB_CEigrpTopoTable

    // The table of established EIGRP peers (neighbors) in the selected autonomous
    // system.   Peers are indexed by their unique internal handle id, as well as
    // the AS number and VPN id.   The peer entry is removed from the table if the
    // peer is declared down.
    CEigrpPeerTable CISCOEIGRPMIB_CEigrpPeerTable

    // The table of interfaces over which EIGRP is running, and their associated
    // statistics.   This table is independent of whether any peer adjacencies
    // have been formed over the interfaces or not.   Interfaces running EIGRP are
    // determined by whether their assigned IP addresses fall within configured
    // EIGRP network statements.
    CEigrpInterfaceTable CISCOEIGRPMIB_CEigrpInterfaceTable
}

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetEntityData() *types.CommonEntityData {
    cISCOEIGRPMIB.EntityData.YFilter = cISCOEIGRPMIB.YFilter
    cISCOEIGRPMIB.EntityData.YangName = "CISCO-EIGRP-MIB"
    cISCOEIGRPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOEIGRPMIB.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    cISCOEIGRPMIB.EntityData.SegmentPath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB"
    cISCOEIGRPMIB.EntityData.AbsolutePath = cISCOEIGRPMIB.EntityData.SegmentPath
    cISCOEIGRPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOEIGRPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOEIGRPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOEIGRPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOEIGRPMIB.EntityData.Children.Append("cEigrpVpnTable", types.YChild{"CEigrpVpnTable", &cISCOEIGRPMIB.CEigrpVpnTable})
    cISCOEIGRPMIB.EntityData.Children.Append("cEigrpTraffStatsTable", types.YChild{"CEigrpTraffStatsTable", &cISCOEIGRPMIB.CEigrpTraffStatsTable})
    cISCOEIGRPMIB.EntityData.Children.Append("cEigrpTopoTable", types.YChild{"CEigrpTopoTable", &cISCOEIGRPMIB.CEigrpTopoTable})
    cISCOEIGRPMIB.EntityData.Children.Append("cEigrpPeerTable", types.YChild{"CEigrpPeerTable", &cISCOEIGRPMIB.CEigrpPeerTable})
    cISCOEIGRPMIB.EntityData.Children.Append("cEigrpInterfaceTable", types.YChild{"CEigrpInterfaceTable", &cISCOEIGRPMIB.CEigrpInterfaceTable})
    cISCOEIGRPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOEIGRPMIB.EntityData.YListKeys = []string {}

    return &(cISCOEIGRPMIB.EntityData)
}

// CISCOEIGRPMIB_CEigrpVpnTable
// This table contains information on those VPN's configured
// to run EIGRP.  The VPN creation on a router is independent
// of the routing protocol to be used over it.   A VPN is
// given a name and has a dedicated routing table associated
// with it.  This routing table is identified internally
// by a unique integer value.
type CISCOEIGRPMIB_CEigrpVpnTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single VPN which is configured to run EIGRP. The
    // type is slice of CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry.
    CEigrpVpnEntry []*CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry
}

func (cEigrpVpnTable *CISCOEIGRPMIB_CEigrpVpnTable) GetEntityData() *types.CommonEntityData {
    cEigrpVpnTable.EntityData.YFilter = cEigrpVpnTable.YFilter
    cEigrpVpnTable.EntityData.YangName = "cEigrpVpnTable"
    cEigrpVpnTable.EntityData.BundleName = "cisco_ios_xe"
    cEigrpVpnTable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    cEigrpVpnTable.EntityData.SegmentPath = "cEigrpVpnTable"
    cEigrpVpnTable.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/" + cEigrpVpnTable.EntityData.SegmentPath
    cEigrpVpnTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpVpnTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpVpnTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpVpnTable.EntityData.Children = types.NewOrderedMap()
    cEigrpVpnTable.EntityData.Children.Append("cEigrpVpnEntry", types.YChild{"CEigrpVpnEntry", nil})
    for i := range cEigrpVpnTable.CEigrpVpnEntry {
        cEigrpVpnTable.EntityData.Children.Append(types.GetSegmentPath(cEigrpVpnTable.CEigrpVpnEntry[i]), types.YChild{"CEigrpVpnEntry", cEigrpVpnTable.CEigrpVpnEntry[i]})
    }
    cEigrpVpnTable.EntityData.Leafs = types.NewOrderedMap()

    cEigrpVpnTable.EntityData.YListKeys = []string {}

    return &(cEigrpVpnTable.EntityData)
}

// CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry
// Information relating to a single VPN which is configured
// to run EIGRP.
type CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The unique VPN identifier.  This is a unique
    // integer relative to all other VPN's defined on the router.  It also
    // identifies internally the routing table instance. The type is interface{}
    // with range: 0..4294967295.
    CEigrpVpnId interface{}

    // The name given to the VPN. The type is string.
    CEigrpVpnName interface{}
}

func (cEigrpVpnEntry *CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry) GetEntityData() *types.CommonEntityData {
    cEigrpVpnEntry.EntityData.YFilter = cEigrpVpnEntry.YFilter
    cEigrpVpnEntry.EntityData.YangName = "cEigrpVpnEntry"
    cEigrpVpnEntry.EntityData.BundleName = "cisco_ios_xe"
    cEigrpVpnEntry.EntityData.ParentYangName = "cEigrpVpnTable"
    cEigrpVpnEntry.EntityData.SegmentPath = "cEigrpVpnEntry" + types.AddKeyToken(cEigrpVpnEntry.CEigrpVpnId, "cEigrpVpnId")
    cEigrpVpnEntry.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/cEigrpVpnTable/" + cEigrpVpnEntry.EntityData.SegmentPath
    cEigrpVpnEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpVpnEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpVpnEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpVpnEntry.EntityData.Children = types.NewOrderedMap()
    cEigrpVpnEntry.EntityData.Leafs = types.NewOrderedMap()
    cEigrpVpnEntry.EntityData.Leafs.Append("cEigrpVpnId", types.YLeaf{"CEigrpVpnId", cEigrpVpnEntry.CEigrpVpnId})
    cEigrpVpnEntry.EntityData.Leafs.Append("cEigrpVpnName", types.YLeaf{"CEigrpVpnName", cEigrpVpnEntry.CEigrpVpnName})

    cEigrpVpnEntry.EntityData.YListKeys = []string {"CEigrpVpnId"}

    return &(cEigrpVpnEntry.EntityData)
}

// CISCOEIGRPMIB_CEigrpTraffStatsTable
// Table of EIGRP traffic statistics and information
// associated with all EIGRP autonomous systems.
type CISCOEIGRPMIB_CEigrpTraffStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The set of statistics and information for a single EIGRP Autonomous System.
    // The type is slice of
    // CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry.
    CEigrpTraffStatsEntry []*CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry
}

func (cEigrpTraffStatsTable *CISCOEIGRPMIB_CEigrpTraffStatsTable) GetEntityData() *types.CommonEntityData {
    cEigrpTraffStatsTable.EntityData.YFilter = cEigrpTraffStatsTable.YFilter
    cEigrpTraffStatsTable.EntityData.YangName = "cEigrpTraffStatsTable"
    cEigrpTraffStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cEigrpTraffStatsTable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    cEigrpTraffStatsTable.EntityData.SegmentPath = "cEigrpTraffStatsTable"
    cEigrpTraffStatsTable.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/" + cEigrpTraffStatsTable.EntityData.SegmentPath
    cEigrpTraffStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpTraffStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpTraffStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpTraffStatsTable.EntityData.Children = types.NewOrderedMap()
    cEigrpTraffStatsTable.EntityData.Children.Append("cEigrpTraffStatsEntry", types.YChild{"CEigrpTraffStatsEntry", nil})
    for i := range cEigrpTraffStatsTable.CEigrpTraffStatsEntry {
        cEigrpTraffStatsTable.EntityData.Children.Append(types.GetSegmentPath(cEigrpTraffStatsTable.CEigrpTraffStatsEntry[i]), types.YChild{"CEigrpTraffStatsEntry", cEigrpTraffStatsTable.CEigrpTraffStatsEntry[i]})
    }
    cEigrpTraffStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cEigrpTraffStatsTable.EntityData.YListKeys = []string {}

    return &(cEigrpTraffStatsTable.EntityData)
}

// CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry
// The set of statistics and information for a single EIGRP
// Autonomous System.
type CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry_CEigrpVpnId
    CEigrpVpnId interface{}

    // This attribute is a key. The Autonomous System number which is unique
    // integer per VPN. The type is interface{} with range: 0..4294967295.
    CEigrpAsNumber interface{}

    // The total number of live EIGRP neighbors formed on all interfaces whose IP
    // addresses fall under networks configured in the EIGRP AS. The type is
    // interface{} with range: 0..4294967295.
    CEigrpNbrCount interface{}

    // The total number Hello packets that have been sent to all EIGRP neighbors
    // formed on all interfaces whose IP addresses fall under networks configured
    // for the EIGRP AS. The type is interface{} with range: 0..4294967295.
    CEigrpHellosSent interface{}

    // The total number Hello packets that have been received from all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpHellosRcvd interface{}

    // The total number routing update packets that have been sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpUpdatesSent interface{}

    // The total number routing update packets that have been received from all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpUpdatesRcvd interface{}

    // The total number alternate route query packets that have been sent to all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpQueriesSent interface{}

    // The total number alternate route query packets that have been received from
    // all EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpQueriesRcvd interface{}

    // The total number query reply packets that have been sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpRepliesSent interface{}

    // The total number query reply packets that have been received from all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpRepliesRcvd interface{}

    // The total number packet acknowledgements that have been sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpAcksSent interface{}

    // The total number packet acknowledgements that have been received from all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpAcksRcvd interface{}

    // The highest number of EIGRP packets in the input queue waiting to be
    // processed internally addressed to this AS. The type is interface{} with
    // range: 0..4294967295.
    CEigrpInputQHighMark interface{}

    // The number of EIGRP packets dropped from the input queue due to it being
    // full within the AS. The type is interface{} with range: 0..4294967295.
    CEigrpInputQDrops interface{}

    // The total number of Stuck-In-Active (SIA) query packets sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpSiaQueriesSent interface{}

    // The total number of Stuck-In-Active (SIA) query packets received from all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    CEigrpSiaQueriesRcvd interface{}

    // The format of the router-id configured or automatically selected for the
    // EIGRP AS. The type is InetAddressType.
    CEigrpAsRouterIdType interface{}

    // The router-id configured or automatically selected for the EIGRP AS.   Each
    // EIGRP routing process has a unique router-id selected from each autonomous
    // system configured. The format is governed by object cEigrpAsRouterIdType.
    // The type is string with length: 0..255.
    CEigrpAsRouterId interface{}

    // The total number of EIGRP derived routes currently existing in the topology
    // table for the AS. The type is interface{} with range: 0..4294967295.
    CEigrpTopoRoutes interface{}

    // Routes in a topology table for an AS are assigned serial numbers and are
    // sequenced internally as they are inserted and deleted.   The serial number
    // of the first route in that internal sequence is called the head serial
    // number. Each AS has its own topology table, and its own serial number
    // space, each of which begins with the value 1. A serial number of zero
    // implies that there are no routes in the topology. The type is interface{}
    // with range: 0..18446744073709551615.
    CEigrpHeadSerial interface{}

    // The serial number that would be assigned to the next new or changed route
    // in the topology table for the AS. The type is interface{} with range:
    // 0..18446744073709551615.
    CEigrpNextSerial interface{}

    // When alternate route query packets are sent to adjacent EIGRP peers in an
    // AS, replies are expected.   This object is the total number of outstanding
    // replies expected to queries that have been sent to peers in the current AS.
    // It remains at zero most of the time until an EIGRP route becomes active.
    // The type is interface{} with range: 0..4294967295.
    CEigrpXmitPendReplies interface{}

    // A dummy is a temporary internal entity used as a place holder in the
    // topology table for an AS. They are not transmitted in routing updates. 
    // This is the total number currently in existence associated with the AS. The
    // type is interface{} with range: 0..4294967295.
    CEigrpXmitDummies interface{}
}

func (cEigrpTraffStatsEntry *CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry) GetEntityData() *types.CommonEntityData {
    cEigrpTraffStatsEntry.EntityData.YFilter = cEigrpTraffStatsEntry.YFilter
    cEigrpTraffStatsEntry.EntityData.YangName = "cEigrpTraffStatsEntry"
    cEigrpTraffStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cEigrpTraffStatsEntry.EntityData.ParentYangName = "cEigrpTraffStatsTable"
    cEigrpTraffStatsEntry.EntityData.SegmentPath = "cEigrpTraffStatsEntry" + types.AddKeyToken(cEigrpTraffStatsEntry.CEigrpVpnId, "cEigrpVpnId") + types.AddKeyToken(cEigrpTraffStatsEntry.CEigrpAsNumber, "cEigrpAsNumber")
    cEigrpTraffStatsEntry.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/cEigrpTraffStatsTable/" + cEigrpTraffStatsEntry.EntityData.SegmentPath
    cEigrpTraffStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpTraffStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpTraffStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpTraffStatsEntry.EntityData.Children = types.NewOrderedMap()
    cEigrpTraffStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpVpnId", types.YLeaf{"CEigrpVpnId", cEigrpTraffStatsEntry.CEigrpVpnId})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpAsNumber", types.YLeaf{"CEigrpAsNumber", cEigrpTraffStatsEntry.CEigrpAsNumber})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpNbrCount", types.YLeaf{"CEigrpNbrCount", cEigrpTraffStatsEntry.CEigrpNbrCount})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpHellosSent", types.YLeaf{"CEigrpHellosSent", cEigrpTraffStatsEntry.CEigrpHellosSent})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpHellosRcvd", types.YLeaf{"CEigrpHellosRcvd", cEigrpTraffStatsEntry.CEigrpHellosRcvd})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpUpdatesSent", types.YLeaf{"CEigrpUpdatesSent", cEigrpTraffStatsEntry.CEigrpUpdatesSent})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpUpdatesRcvd", types.YLeaf{"CEigrpUpdatesRcvd", cEigrpTraffStatsEntry.CEigrpUpdatesRcvd})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpQueriesSent", types.YLeaf{"CEigrpQueriesSent", cEigrpTraffStatsEntry.CEigrpQueriesSent})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpQueriesRcvd", types.YLeaf{"CEigrpQueriesRcvd", cEigrpTraffStatsEntry.CEigrpQueriesRcvd})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpRepliesSent", types.YLeaf{"CEigrpRepliesSent", cEigrpTraffStatsEntry.CEigrpRepliesSent})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpRepliesRcvd", types.YLeaf{"CEigrpRepliesRcvd", cEigrpTraffStatsEntry.CEigrpRepliesRcvd})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpAcksSent", types.YLeaf{"CEigrpAcksSent", cEigrpTraffStatsEntry.CEigrpAcksSent})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpAcksRcvd", types.YLeaf{"CEigrpAcksRcvd", cEigrpTraffStatsEntry.CEigrpAcksRcvd})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpInputQHighMark", types.YLeaf{"CEigrpInputQHighMark", cEigrpTraffStatsEntry.CEigrpInputQHighMark})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpInputQDrops", types.YLeaf{"CEigrpInputQDrops", cEigrpTraffStatsEntry.CEigrpInputQDrops})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpSiaQueriesSent", types.YLeaf{"CEigrpSiaQueriesSent", cEigrpTraffStatsEntry.CEigrpSiaQueriesSent})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpSiaQueriesRcvd", types.YLeaf{"CEigrpSiaQueriesRcvd", cEigrpTraffStatsEntry.CEigrpSiaQueriesRcvd})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpAsRouterIdType", types.YLeaf{"CEigrpAsRouterIdType", cEigrpTraffStatsEntry.CEigrpAsRouterIdType})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpAsRouterId", types.YLeaf{"CEigrpAsRouterId", cEigrpTraffStatsEntry.CEigrpAsRouterId})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpTopoRoutes", types.YLeaf{"CEigrpTopoRoutes", cEigrpTraffStatsEntry.CEigrpTopoRoutes})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpHeadSerial", types.YLeaf{"CEigrpHeadSerial", cEigrpTraffStatsEntry.CEigrpHeadSerial})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpNextSerial", types.YLeaf{"CEigrpNextSerial", cEigrpTraffStatsEntry.CEigrpNextSerial})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpXmitPendReplies", types.YLeaf{"CEigrpXmitPendReplies", cEigrpTraffStatsEntry.CEigrpXmitPendReplies})
    cEigrpTraffStatsEntry.EntityData.Leafs.Append("cEigrpXmitDummies", types.YLeaf{"CEigrpXmitDummies", cEigrpTraffStatsEntry.CEigrpXmitDummies})

    cEigrpTraffStatsEntry.EntityData.YListKeys = []string {"CEigrpVpnId", "CEigrpAsNumber"}

    return &(cEigrpTraffStatsEntry.EntityData)
}

// CISCOEIGRPMIB_CEigrpTopoTable
// The table of EIGRP routes and their associated
// attributes for an Autonomous System (AS) configured
// in a VPN is called a topology table.  All route entries in
// the topology table will be indexed by IP network type,
// IP network number and network mask (prefix) size.
type CISCOEIGRPMIB_CEigrpTopoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The entry for a single EIGRP topology table in the given AS. The type is
    // slice of CISCOEIGRPMIB_CEigrpTopoTable_CEigrpTopoEntry.
    CEigrpTopoEntry []*CISCOEIGRPMIB_CEigrpTopoTable_CEigrpTopoEntry
}

func (cEigrpTopoTable *CISCOEIGRPMIB_CEigrpTopoTable) GetEntityData() *types.CommonEntityData {
    cEigrpTopoTable.EntityData.YFilter = cEigrpTopoTable.YFilter
    cEigrpTopoTable.EntityData.YangName = "cEigrpTopoTable"
    cEigrpTopoTable.EntityData.BundleName = "cisco_ios_xe"
    cEigrpTopoTable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    cEigrpTopoTable.EntityData.SegmentPath = "cEigrpTopoTable"
    cEigrpTopoTable.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/" + cEigrpTopoTable.EntityData.SegmentPath
    cEigrpTopoTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpTopoTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpTopoTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpTopoTable.EntityData.Children = types.NewOrderedMap()
    cEigrpTopoTable.EntityData.Children.Append("cEigrpTopoEntry", types.YChild{"CEigrpTopoEntry", nil})
    for i := range cEigrpTopoTable.CEigrpTopoEntry {
        cEigrpTopoTable.EntityData.Children.Append(types.GetSegmentPath(cEigrpTopoTable.CEigrpTopoEntry[i]), types.YChild{"CEigrpTopoEntry", cEigrpTopoTable.CEigrpTopoEntry[i]})
    }
    cEigrpTopoTable.EntityData.Leafs = types.NewOrderedMap()

    cEigrpTopoTable.EntityData.YListKeys = []string {}

    return &(cEigrpTopoTable.EntityData)
}

// CISCOEIGRPMIB_CEigrpTopoTable_CEigrpTopoEntry
// The entry for a single EIGRP topology table in the given
// AS.
type CISCOEIGRPMIB_CEigrpTopoTable_CEigrpTopoEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry_CEigrpVpnId
    CEigrpVpnId interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry_CEigrpAsNumber
    CEigrpAsNumber interface{}

    // This attribute is a key. The format of the destination IP network number
    // for a single route in the topology table in the AS specified in
    // cEigrpDestNet. The type is InetAddressType.
    CEigrpDestNetType interface{}

    // This attribute is a key. The destination IP network number for a single
    // route in the topology table in the AS.  The format is governed by object
    // cEigrpDestNetType. The type is string with length: 0..255.
    CEigrpDestNet interface{}

    // This attribute is a key. The prefix length associated with the destination
    // IP network address for a single route in the topology table in the AS.  The
    // format is governed by the object cEigrpDestNetType. The type is interface{}
    // with range: 0..2040.
    CEigrpDestNetPrefixLen interface{}

    // A value of true(1) indicates the route to the destination network has
    // failed and an active (query) search for an alternative path is in progress.
    // A value of false(2) indicates the route is stable (passive). The type is
    // bool.
    CEigrpActive interface{}

    // A value of true(1) indicates that that this route which is in active state
    // (cEigrpActive = true(1)) has not received any replies to queries for
    // alternate paths, and a second EIGRP route query, called a stuck-in-active
    // query, has now been sent. The type is bool.
    CEigrpStuckInActive interface{}

    // A successor is the next routing hop for a path to the destination IP
    // network number for a single route in the topology table in the AS.  There
    // can be several potential successors if there are multiple paths to the
    // destination.   This is the total number of successors for a topology entry.
    // The type is interface{} with range: 0..4294967295.
    CEigrpDestSuccessors interface{}

    // The feasibility (best) distance is the minimum distance from this router to
    // the destination IP network in this topology entry.  The feasibility
    // distance is used in determining the best successor for a path to the
    // destination network. The type is interface{} with range: 0..4294967295.
    CEigrpFdistance interface{}

    // This is a text string describing the internal origin of the EIGRP route
    // represented by the topology entry. The type is string.
    CEigrpRouteOriginType interface{}

    // The format of the IP address defined as the origin of this topology route
    // entry. The type is InetAddressType.
    CEigrpRouteOriginAddrType interface{}

    // If the origin of the topology route entry is external to this router, then
    // this object is the IP address of the router from which it originated.  The
    // format  is governed by object cEigrpRouteOriginAddrType. The type is string
    // with length: 0..255.
    CEigrpRouteOriginAddr interface{}

    // The format of the next hop IP address for the route represented by the
    // topology entry. The type is InetAddressType.
    CEigrpNextHopAddressType interface{}

    // This is the next hop IP address for the route represented by the topology
    // entry.   The next hop is where network traffic will be routed to in order
    // to reach the destination network for this topology entry.  The format is
    // governed by cEigrpNextHopAddressType. The type is string with length:
    // 0..255.
    CEigrpNextHopAddress interface{}

    // The interface through which the next hop IP address is reached to send
    // network traffic to the destination network represented by the topology
    // entry. The type is string.
    CEigrpNextHopInterface interface{}

    // The computed distance to the destination network entry from this router.
    // The type is interface{} with range: 0..4294967295.
    CEigrpDistance interface{}

    // The computed distance to the destination network in the topology entry
    // reported to this router by the originator of this route. The type is
    // interface{} with range: 0..4294967295.
    CEigrpReportDistance interface{}
}

func (cEigrpTopoEntry *CISCOEIGRPMIB_CEigrpTopoTable_CEigrpTopoEntry) GetEntityData() *types.CommonEntityData {
    cEigrpTopoEntry.EntityData.YFilter = cEigrpTopoEntry.YFilter
    cEigrpTopoEntry.EntityData.YangName = "cEigrpTopoEntry"
    cEigrpTopoEntry.EntityData.BundleName = "cisco_ios_xe"
    cEigrpTopoEntry.EntityData.ParentYangName = "cEigrpTopoTable"
    cEigrpTopoEntry.EntityData.SegmentPath = "cEigrpTopoEntry" + types.AddKeyToken(cEigrpTopoEntry.CEigrpVpnId, "cEigrpVpnId") + types.AddKeyToken(cEigrpTopoEntry.CEigrpAsNumber, "cEigrpAsNumber") + types.AddKeyToken(cEigrpTopoEntry.CEigrpDestNetType, "cEigrpDestNetType") + types.AddKeyToken(cEigrpTopoEntry.CEigrpDestNet, "cEigrpDestNet") + types.AddKeyToken(cEigrpTopoEntry.CEigrpDestNetPrefixLen, "cEigrpDestNetPrefixLen")
    cEigrpTopoEntry.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/cEigrpTopoTable/" + cEigrpTopoEntry.EntityData.SegmentPath
    cEigrpTopoEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpTopoEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpTopoEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpTopoEntry.EntityData.Children = types.NewOrderedMap()
    cEigrpTopoEntry.EntityData.Leafs = types.NewOrderedMap()
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpVpnId", types.YLeaf{"CEigrpVpnId", cEigrpTopoEntry.CEigrpVpnId})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpAsNumber", types.YLeaf{"CEigrpAsNumber", cEigrpTopoEntry.CEigrpAsNumber})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpDestNetType", types.YLeaf{"CEigrpDestNetType", cEigrpTopoEntry.CEigrpDestNetType})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpDestNet", types.YLeaf{"CEigrpDestNet", cEigrpTopoEntry.CEigrpDestNet})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpDestNetPrefixLen", types.YLeaf{"CEigrpDestNetPrefixLen", cEigrpTopoEntry.CEigrpDestNetPrefixLen})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpActive", types.YLeaf{"CEigrpActive", cEigrpTopoEntry.CEigrpActive})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpStuckInActive", types.YLeaf{"CEigrpStuckInActive", cEigrpTopoEntry.CEigrpStuckInActive})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpDestSuccessors", types.YLeaf{"CEigrpDestSuccessors", cEigrpTopoEntry.CEigrpDestSuccessors})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpFdistance", types.YLeaf{"CEigrpFdistance", cEigrpTopoEntry.CEigrpFdistance})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpRouteOriginType", types.YLeaf{"CEigrpRouteOriginType", cEigrpTopoEntry.CEigrpRouteOriginType})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpRouteOriginAddrType", types.YLeaf{"CEigrpRouteOriginAddrType", cEigrpTopoEntry.CEigrpRouteOriginAddrType})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpRouteOriginAddr", types.YLeaf{"CEigrpRouteOriginAddr", cEigrpTopoEntry.CEigrpRouteOriginAddr})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpNextHopAddressType", types.YLeaf{"CEigrpNextHopAddressType", cEigrpTopoEntry.CEigrpNextHopAddressType})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpNextHopAddress", types.YLeaf{"CEigrpNextHopAddress", cEigrpTopoEntry.CEigrpNextHopAddress})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpNextHopInterface", types.YLeaf{"CEigrpNextHopInterface", cEigrpTopoEntry.CEigrpNextHopInterface})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpDistance", types.YLeaf{"CEigrpDistance", cEigrpTopoEntry.CEigrpDistance})
    cEigrpTopoEntry.EntityData.Leafs.Append("cEigrpReportDistance", types.YLeaf{"CEigrpReportDistance", cEigrpTopoEntry.CEigrpReportDistance})

    cEigrpTopoEntry.EntityData.YListKeys = []string {"CEigrpVpnId", "CEigrpAsNumber", "CEigrpDestNetType", "CEigrpDestNet", "CEigrpDestNetPrefixLen"}

    return &(cEigrpTopoEntry.EntityData)
}

// CISCOEIGRPMIB_CEigrpPeerTable
// The table of established EIGRP peers (neighbors) in the
// selected autonomous system.   Peers are indexed by their
// unique internal handle id, as well as the AS number and
// VPN id.   The peer entry is removed from the table if
// the peer is declared down.
type CISCOEIGRPMIB_CEigrpPeerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics and operational parameters for a single peer in the AS. The type
    // is slice of CISCOEIGRPMIB_CEigrpPeerTable_CEigrpPeerEntry.
    CEigrpPeerEntry []*CISCOEIGRPMIB_CEigrpPeerTable_CEigrpPeerEntry
}

func (cEigrpPeerTable *CISCOEIGRPMIB_CEigrpPeerTable) GetEntityData() *types.CommonEntityData {
    cEigrpPeerTable.EntityData.YFilter = cEigrpPeerTable.YFilter
    cEigrpPeerTable.EntityData.YangName = "cEigrpPeerTable"
    cEigrpPeerTable.EntityData.BundleName = "cisco_ios_xe"
    cEigrpPeerTable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    cEigrpPeerTable.EntityData.SegmentPath = "cEigrpPeerTable"
    cEigrpPeerTable.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/" + cEigrpPeerTable.EntityData.SegmentPath
    cEigrpPeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpPeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpPeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpPeerTable.EntityData.Children = types.NewOrderedMap()
    cEigrpPeerTable.EntityData.Children.Append("cEigrpPeerEntry", types.YChild{"CEigrpPeerEntry", nil})
    for i := range cEigrpPeerTable.CEigrpPeerEntry {
        cEigrpPeerTable.EntityData.Children.Append(types.GetSegmentPath(cEigrpPeerTable.CEigrpPeerEntry[i]), types.YChild{"CEigrpPeerEntry", cEigrpPeerTable.CEigrpPeerEntry[i]})
    }
    cEigrpPeerTable.EntityData.Leafs = types.NewOrderedMap()

    cEigrpPeerTable.EntityData.YListKeys = []string {}

    return &(cEigrpPeerTable.EntityData)
}

// CISCOEIGRPMIB_CEigrpPeerTable_CEigrpPeerEntry
// Statistics and operational parameters for a single peer
// in the AS.
type CISCOEIGRPMIB_CEigrpPeerTable_CEigrpPeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry_CEigrpVpnId
    CEigrpVpnId interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry_CEigrpAsNumber
    CEigrpAsNumber interface{}

    // This attribute is a key. The unique internal identifier for the peer in the
    // AS. This is a unique value among peer entries in a selected table. The type
    // is interface{} with range: 0..4294967295.
    CEigrpHandle interface{}

    // The format of the remote source IP address used by the peer to establish
    // the EIGRP adjacency with this router. The type is InetAddressType.
    CEigrpPeerAddrType interface{}

    // The source IP address used by the peer to establish the EIGRP adjacency
    // with this router.  The format is governed by object cEigrpPeerAddrType. The
    // type is string with length: 0..255.
    CEigrpPeerAddr interface{}

    // The ifIndex of the interface on this router through which this peer can be
    // reached. The type is interface{} with range: 0..2147483647.
    CEigrpPeerIfIndex interface{}

    // The count-down timer indicating how much time must pass without receiving a
    // hello packet from this EIGRP peer before this router declares the peer
    // down. A peer declared as down is removed from the table and is no longer
    // visible. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CEigrpHoldTime interface{}

    // The elapsed time since the EIGRP adjacency was first established with the
    // peer. The type is string.
    CEigrpUpTime interface{}

    // The computed smooth round trip time for packets to and from the peer. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    CEigrpSrtt interface{}

    // The computed retransmission timeout for the peer. This value is computed
    // over time as packets are sent to the peer and acknowledgements are received
    // from it, and is the amount of time to wait before resending a packet from
    // the retransmission queue to the peer when an expected acknowledgement has
    // not been received. The type is interface{} with range: 0..4294967295. Units
    // are milliseconds.
    CEigrpRto interface{}

    // The number of any EIGRP packets currently enqueued waiting to be sent to
    // this peer. The type is interface{} with range: 0..4294967295.
    CEigrpPktsEnqueued interface{}

    // All transmitted EIGRP packets have a sequence number assigned. This is the
    // sequence number of the last EIGRP packet sent to this peer. The type is
    // interface{} with range: 0..4294967295.
    CEigrpLastSeq interface{}

    // The EIGRP version information reported by the remote peer. The type is
    // string.
    CEigrpVersion interface{}

    // The cumulative number of retransmissions to this peer during the period
    // that the peer adjacency has remained up. The type is interface{} with
    // range: 0..4294967295.
    CEigrpRetrans interface{}

    // The number of times the current unacknowledged packet has been retried,
    // i.e. resent to this peer to be acknowledged. The type is interface{} with
    // range: 0..4294967295.
    CEigrpRetries interface{}
}

func (cEigrpPeerEntry *CISCOEIGRPMIB_CEigrpPeerTable_CEigrpPeerEntry) GetEntityData() *types.CommonEntityData {
    cEigrpPeerEntry.EntityData.YFilter = cEigrpPeerEntry.YFilter
    cEigrpPeerEntry.EntityData.YangName = "cEigrpPeerEntry"
    cEigrpPeerEntry.EntityData.BundleName = "cisco_ios_xe"
    cEigrpPeerEntry.EntityData.ParentYangName = "cEigrpPeerTable"
    cEigrpPeerEntry.EntityData.SegmentPath = "cEigrpPeerEntry" + types.AddKeyToken(cEigrpPeerEntry.CEigrpVpnId, "cEigrpVpnId") + types.AddKeyToken(cEigrpPeerEntry.CEigrpAsNumber, "cEigrpAsNumber") + types.AddKeyToken(cEigrpPeerEntry.CEigrpHandle, "cEigrpHandle")
    cEigrpPeerEntry.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/cEigrpPeerTable/" + cEigrpPeerEntry.EntityData.SegmentPath
    cEigrpPeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpPeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpPeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpPeerEntry.EntityData.Children = types.NewOrderedMap()
    cEigrpPeerEntry.EntityData.Leafs = types.NewOrderedMap()
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpVpnId", types.YLeaf{"CEigrpVpnId", cEigrpPeerEntry.CEigrpVpnId})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpAsNumber", types.YLeaf{"CEigrpAsNumber", cEigrpPeerEntry.CEigrpAsNumber})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpHandle", types.YLeaf{"CEigrpHandle", cEigrpPeerEntry.CEigrpHandle})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpPeerAddrType", types.YLeaf{"CEigrpPeerAddrType", cEigrpPeerEntry.CEigrpPeerAddrType})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpPeerAddr", types.YLeaf{"CEigrpPeerAddr", cEigrpPeerEntry.CEigrpPeerAddr})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpPeerIfIndex", types.YLeaf{"CEigrpPeerIfIndex", cEigrpPeerEntry.CEigrpPeerIfIndex})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpHoldTime", types.YLeaf{"CEigrpHoldTime", cEigrpPeerEntry.CEigrpHoldTime})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpUpTime", types.YLeaf{"CEigrpUpTime", cEigrpPeerEntry.CEigrpUpTime})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpSrtt", types.YLeaf{"CEigrpSrtt", cEigrpPeerEntry.CEigrpSrtt})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpRto", types.YLeaf{"CEigrpRto", cEigrpPeerEntry.CEigrpRto})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpPktsEnqueued", types.YLeaf{"CEigrpPktsEnqueued", cEigrpPeerEntry.CEigrpPktsEnqueued})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpLastSeq", types.YLeaf{"CEigrpLastSeq", cEigrpPeerEntry.CEigrpLastSeq})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpVersion", types.YLeaf{"CEigrpVersion", cEigrpPeerEntry.CEigrpVersion})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpRetrans", types.YLeaf{"CEigrpRetrans", cEigrpPeerEntry.CEigrpRetrans})
    cEigrpPeerEntry.EntityData.Leafs.Append("cEigrpRetries", types.YLeaf{"CEigrpRetries", cEigrpPeerEntry.CEigrpRetries})

    cEigrpPeerEntry.EntityData.YListKeys = []string {"CEigrpVpnId", "CEigrpAsNumber", "CEigrpHandle"}

    return &(cEigrpPeerEntry.EntityData)
}

// CISCOEIGRPMIB_CEigrpInterfaceTable
// The table of interfaces over which EIGRP is running, and
// their associated statistics.   This table is independent
// of whether any peer adjacencies have been formed over
// the interfaces or not.   Interfaces running EIGRP are
// determined by whether their assigned IP addresses fall
// within configured EIGRP network statements.
type CISCOEIGRPMIB_CEigrpInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information for a single interface running EIGRP in the AS and VPN. The
    // type is slice of CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry.
    CEigrpInterfaceEntry []*CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry
}

func (cEigrpInterfaceTable *CISCOEIGRPMIB_CEigrpInterfaceTable) GetEntityData() *types.CommonEntityData {
    cEigrpInterfaceTable.EntityData.YFilter = cEigrpInterfaceTable.YFilter
    cEigrpInterfaceTable.EntityData.YangName = "cEigrpInterfaceTable"
    cEigrpInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    cEigrpInterfaceTable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    cEigrpInterfaceTable.EntityData.SegmentPath = "cEigrpInterfaceTable"
    cEigrpInterfaceTable.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/" + cEigrpInterfaceTable.EntityData.SegmentPath
    cEigrpInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpInterfaceTable.EntityData.Children = types.NewOrderedMap()
    cEigrpInterfaceTable.EntityData.Children.Append("cEigrpInterfaceEntry", types.YChild{"CEigrpInterfaceEntry", nil})
    for i := range cEigrpInterfaceTable.CEigrpInterfaceEntry {
        cEigrpInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(cEigrpInterfaceTable.CEigrpInterfaceEntry[i]), types.YChild{"CEigrpInterfaceEntry", cEigrpInterfaceTable.CEigrpInterfaceEntry[i]})
    }
    cEigrpInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    cEigrpInterfaceTable.EntityData.YListKeys = []string {}

    return &(cEigrpInterfaceTable.EntityData)
}

// CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry
// Information for a single interface running EIGRP in the
// AS and VPN.
type CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_CEigrpVpnTable_CEigrpVpnEntry_CEigrpVpnId
    CEigrpVpnId interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_CEigrpTraffStatsTable_CEigrpTraffStatsEntry_CEigrpAsNumber
    CEigrpAsNumber interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of EIGRP adjacencies currently formed with peers reached through
    // this interface. The type is interface{} with range: 0..4294967295.
    CEigrpPeerCount interface{}

    // The number of EIGRP packets currently waiting in the reliable transport
    // (acknowledgement-required)  transmission queue to be sent to a peer. The
    // type is interface{} with range: 0..4294967295.
    CEigrpXmitReliableQ interface{}

    // The number EIGRP of packets currently waiting in the unreliable transport
    // (no acknowledgement required) transmission queue. The type is interface{}
    // with range: 0..4294967295.
    CEigrpXmitUnreliableQ interface{}

    // The average of all the computed smooth round trip time values for a packet
    // to and from all peers established on this interface. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CEigrpMeanSrtt interface{}

    // The configured time interval between EIGRP packet transmissions on the
    // interface when the reliable transport method is used. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CEigrpPacingReliable interface{}

    // The configured time interval between EIGRP packet transmissions on the
    // interface when the unreliable transport method is used. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CEigrpPacingUnreliable interface{}

    // The configured multicast flow control timer value for this interface. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    CEigrpMFlowTimer interface{}

    // The number of queued EIGRP routing updates awaiting transmission on this
    // interface. The type is interface{} with range: 0..4294967295.
    CEigrpPendingRoutes interface{}

    // The configured time interval between Hello packet transmissions for this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CEigrpHelloInterval interface{}

    // The serial number of the next EIGRP packet that is to be queued for
    // transmission on this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CEigrpXmitNextSerial interface{}

    // The total number of unreliable (no acknowledgement required) EIGRP
    // multicast packets sent on this interface. The type is interface{} with
    // range: 0..4294967295.
    CEigrpUMcasts interface{}

    // The total number of reliable (acknowledgement required) EIGRP multicast
    // packets sent on this interface. The type is interface{} with range:
    // 0..4294967295.
    CEigrpRMcasts interface{}

    // The total number of unreliable (no acknowledgement required) EIGRP unicast
    // packets sent on this interface. The type is interface{} with range:
    // 0..4294967295.
    CEigrpUUcasts interface{}

    // The total number of reliable (acknowledgement required) unicast packets
    // sent on this interface. The type is interface{} with range: 0..4294967295.
    CEigrpRUcasts interface{}

    // The total number of EIGRP multicast exception transmissions that have
    // occurred on this interface. The type is interface{} with range:
    // 0..4294967295.
    CEigrpMcastExcepts interface{}

    // The total number EIGRP Conditional-Receive packets sent on this interface.
    // The type is interface{} with range: 0..4294967295.
    CEigrpCRpkts interface{}

    // The total number of individual EIGRP acknowledgement packets that have been
    // suppressed and combined in an already enqueued outbound reliable packet on
    // this interface. The type is interface{} with range: 0..4294967295.
    CEigrpAcksSuppressed interface{}

    // The total number EIGRP packet retransmissions sent on the interface. The
    // type is interface{} with range: 0..4294967295.
    CEigrpRetransSent interface{}

    // The total number of out-of-sequence EIGRP packets received. The type is
    // interface{} with range: 0..4294967295.
    CEigrpOOSrvcd interface{}

    // The EIGRP authentication mode of the interface. none  :  no authentication
    // enabled on the interface md5   :  MD5 authentication enabled on the
    // interface. The type is CEigrpAuthMode.
    CEigrpAuthMode interface{}

    // The name of the authentication key-chain configured on this interface.  
    // The key-chain is a reference to which set of secret keys are to be accessed
    // in order to determine which secret key string to use.  The key chain name
    // is not the secret key string password and can also be used in other routing
    // protocols, such as RIP and ISIS. The type is string.
    CEigrpAuthKeyChain interface{}
}

func (cEigrpInterfaceEntry *CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry) GetEntityData() *types.CommonEntityData {
    cEigrpInterfaceEntry.EntityData.YFilter = cEigrpInterfaceEntry.YFilter
    cEigrpInterfaceEntry.EntityData.YangName = "cEigrpInterfaceEntry"
    cEigrpInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    cEigrpInterfaceEntry.EntityData.ParentYangName = "cEigrpInterfaceTable"
    cEigrpInterfaceEntry.EntityData.SegmentPath = "cEigrpInterfaceEntry" + types.AddKeyToken(cEigrpInterfaceEntry.CEigrpVpnId, "cEigrpVpnId") + types.AddKeyToken(cEigrpInterfaceEntry.CEigrpAsNumber, "cEigrpAsNumber") + types.AddKeyToken(cEigrpInterfaceEntry.IfIndex, "ifIndex")
    cEigrpInterfaceEntry.EntityData.AbsolutePath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB/cEigrpInterfaceTable/" + cEigrpInterfaceEntry.EntityData.SegmentPath
    cEigrpInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cEigrpInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cEigrpInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cEigrpInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    cEigrpInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpVpnId", types.YLeaf{"CEigrpVpnId", cEigrpInterfaceEntry.CEigrpVpnId})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpAsNumber", types.YLeaf{"CEigrpAsNumber", cEigrpInterfaceEntry.CEigrpAsNumber})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cEigrpInterfaceEntry.IfIndex})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpPeerCount", types.YLeaf{"CEigrpPeerCount", cEigrpInterfaceEntry.CEigrpPeerCount})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpXmitReliableQ", types.YLeaf{"CEigrpXmitReliableQ", cEigrpInterfaceEntry.CEigrpXmitReliableQ})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpXmitUnreliableQ", types.YLeaf{"CEigrpXmitUnreliableQ", cEigrpInterfaceEntry.CEigrpXmitUnreliableQ})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpMeanSrtt", types.YLeaf{"CEigrpMeanSrtt", cEigrpInterfaceEntry.CEigrpMeanSrtt})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpPacingReliable", types.YLeaf{"CEigrpPacingReliable", cEigrpInterfaceEntry.CEigrpPacingReliable})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpPacingUnreliable", types.YLeaf{"CEigrpPacingUnreliable", cEigrpInterfaceEntry.CEigrpPacingUnreliable})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpMFlowTimer", types.YLeaf{"CEigrpMFlowTimer", cEigrpInterfaceEntry.CEigrpMFlowTimer})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpPendingRoutes", types.YLeaf{"CEigrpPendingRoutes", cEigrpInterfaceEntry.CEigrpPendingRoutes})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpHelloInterval", types.YLeaf{"CEigrpHelloInterval", cEigrpInterfaceEntry.CEigrpHelloInterval})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpXmitNextSerial", types.YLeaf{"CEigrpXmitNextSerial", cEigrpInterfaceEntry.CEigrpXmitNextSerial})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpUMcasts", types.YLeaf{"CEigrpUMcasts", cEigrpInterfaceEntry.CEigrpUMcasts})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpRMcasts", types.YLeaf{"CEigrpRMcasts", cEigrpInterfaceEntry.CEigrpRMcasts})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpUUcasts", types.YLeaf{"CEigrpUUcasts", cEigrpInterfaceEntry.CEigrpUUcasts})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpRUcasts", types.YLeaf{"CEigrpRUcasts", cEigrpInterfaceEntry.CEigrpRUcasts})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpMcastExcepts", types.YLeaf{"CEigrpMcastExcepts", cEigrpInterfaceEntry.CEigrpMcastExcepts})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpCRpkts", types.YLeaf{"CEigrpCRpkts", cEigrpInterfaceEntry.CEigrpCRpkts})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpAcksSuppressed", types.YLeaf{"CEigrpAcksSuppressed", cEigrpInterfaceEntry.CEigrpAcksSuppressed})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpRetransSent", types.YLeaf{"CEigrpRetransSent", cEigrpInterfaceEntry.CEigrpRetransSent})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpOOSrvcd", types.YLeaf{"CEigrpOOSrvcd", cEigrpInterfaceEntry.CEigrpOOSrvcd})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpAuthMode", types.YLeaf{"CEigrpAuthMode", cEigrpInterfaceEntry.CEigrpAuthMode})
    cEigrpInterfaceEntry.EntityData.Leafs.Append("cEigrpAuthKeyChain", types.YLeaf{"CEigrpAuthKeyChain", cEigrpInterfaceEntry.CEigrpAuthKeyChain})

    cEigrpInterfaceEntry.EntityData.YListKeys = []string {"CEigrpVpnId", "CEigrpAsNumber", "IfIndex"}

    return &(cEigrpInterfaceEntry.EntityData)
}

// CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry_CEigrpAuthMode represents md5   :  MD5 authentication enabled on the interface
type CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry_CEigrpAuthMode string

const (
    CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry_CEigrpAuthMode_none CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry_CEigrpAuthMode = "none"

    CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry_CEigrpAuthMode_md5 CISCOEIGRPMIB_CEigrpInterfaceTable_CEigrpInterfaceEntry_CEigrpAuthMode = "md5"
)

