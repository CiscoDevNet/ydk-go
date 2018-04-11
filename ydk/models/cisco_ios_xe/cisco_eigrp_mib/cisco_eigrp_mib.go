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
    Ceigrpvpntable CISCOEIGRPMIB_Ceigrpvpntable

    // Table of EIGRP traffic statistics and information associated with all EIGRP
    // autonomous systems.
    Ceigrptraffstatstable CISCOEIGRPMIB_Ceigrptraffstatstable

    // The table of EIGRP routes and their associated attributes for an Autonomous
    // System (AS) configured in a VPN is called a topology table.  All route
    // entries in the topology table will be indexed by IP network type, IP
    // network number and network mask (prefix) size.
    Ceigrptopotable CISCOEIGRPMIB_Ceigrptopotable

    // The table of established EIGRP peers (neighbors) in the selected autonomous
    // system.   Peers are indexed by their unique internal handle id, as well as
    // the AS number and VPN id.   The peer entry is removed from the table if the
    // peer is declared down.
    Ceigrppeertable CISCOEIGRPMIB_Ceigrppeertable

    // The table of interfaces over which EIGRP is running, and their associated
    // statistics.   This table is independent of whether any peer adjacencies
    // have been formed over the interfaces or not.   Interfaces running EIGRP are
    // determined by whether their assigned IP addresses fall within configured
    // EIGRP network statements.
    Ceigrpinterfacetable CISCOEIGRPMIB_Ceigrpinterfacetable
}

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetEntityData() *types.CommonEntityData {
    cISCOEIGRPMIB.EntityData.YFilter = cISCOEIGRPMIB.YFilter
    cISCOEIGRPMIB.EntityData.YangName = "CISCO-EIGRP-MIB"
    cISCOEIGRPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOEIGRPMIB.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    cISCOEIGRPMIB.EntityData.SegmentPath = "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB"
    cISCOEIGRPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOEIGRPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOEIGRPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOEIGRPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOEIGRPMIB.EntityData.Children["cEigrpVpnTable"] = types.YChild{"Ceigrpvpntable", &cISCOEIGRPMIB.Ceigrpvpntable}
    cISCOEIGRPMIB.EntityData.Children["cEigrpTraffStatsTable"] = types.YChild{"Ceigrptraffstatstable", &cISCOEIGRPMIB.Ceigrptraffstatstable}
    cISCOEIGRPMIB.EntityData.Children["cEigrpTopoTable"] = types.YChild{"Ceigrptopotable", &cISCOEIGRPMIB.Ceigrptopotable}
    cISCOEIGRPMIB.EntityData.Children["cEigrpPeerTable"] = types.YChild{"Ceigrppeertable", &cISCOEIGRPMIB.Ceigrppeertable}
    cISCOEIGRPMIB.EntityData.Children["cEigrpInterfaceTable"] = types.YChild{"Ceigrpinterfacetable", &cISCOEIGRPMIB.Ceigrpinterfacetable}
    cISCOEIGRPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOEIGRPMIB.EntityData)
}

// CISCOEIGRPMIB_Ceigrpvpntable
// This table contains information on those VPN's configured
// to run EIGRP.  The VPN creation on a router is independent
// of the routing protocol to be used over it.   A VPN is
// given a name and has a dedicated routing table associated
// with it.  This routing table is identified internally
// by a unique integer value.
type CISCOEIGRPMIB_Ceigrpvpntable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single VPN which is configured to run EIGRP. The
    // type is slice of CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry.
    Ceigrpvpnentry []CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry
}

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetEntityData() *types.CommonEntityData {
    ceigrpvpntable.EntityData.YFilter = ceigrpvpntable.YFilter
    ceigrpvpntable.EntityData.YangName = "cEigrpVpnTable"
    ceigrpvpntable.EntityData.BundleName = "cisco_ios_xe"
    ceigrpvpntable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    ceigrpvpntable.EntityData.SegmentPath = "cEigrpVpnTable"
    ceigrpvpntable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrpvpntable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrpvpntable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrpvpntable.EntityData.Children = make(map[string]types.YChild)
    ceigrpvpntable.EntityData.Children["cEigrpVpnEntry"] = types.YChild{"Ceigrpvpnentry", nil}
    for i := range ceigrpvpntable.Ceigrpvpnentry {
        ceigrpvpntable.EntityData.Children[types.GetSegmentPath(&ceigrpvpntable.Ceigrpvpnentry[i])] = types.YChild{"Ceigrpvpnentry", &ceigrpvpntable.Ceigrpvpnentry[i]}
    }
    ceigrpvpntable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceigrpvpntable.EntityData)
}

// CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry
// Information relating to a single VPN which is configured
// to run EIGRP.
type CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The unique VPN identifier.  This is a unique
    // integer relative to all other VPN's defined on the router.  It also
    // identifies internally the routing table instance. The type is interface{}
    // with range: 0..4294967295.
    Ceigrpvpnid interface{}

    // The name given to the VPN. The type is string.
    Ceigrpvpnname interface{}
}

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetEntityData() *types.CommonEntityData {
    ceigrpvpnentry.EntityData.YFilter = ceigrpvpnentry.YFilter
    ceigrpvpnentry.EntityData.YangName = "cEigrpVpnEntry"
    ceigrpvpnentry.EntityData.BundleName = "cisco_ios_xe"
    ceigrpvpnentry.EntityData.ParentYangName = "cEigrpVpnTable"
    ceigrpvpnentry.EntityData.SegmentPath = "cEigrpVpnEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrpvpnentry.Ceigrpvpnid) + "']"
    ceigrpvpnentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrpvpnentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrpvpnentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrpvpnentry.EntityData.Children = make(map[string]types.YChild)
    ceigrpvpnentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceigrpvpnentry.EntityData.Leafs["cEigrpVpnId"] = types.YLeaf{"Ceigrpvpnid", ceigrpvpnentry.Ceigrpvpnid}
    ceigrpvpnentry.EntityData.Leafs["cEigrpVpnName"] = types.YLeaf{"Ceigrpvpnname", ceigrpvpnentry.Ceigrpvpnname}
    return &(ceigrpvpnentry.EntityData)
}

// CISCOEIGRPMIB_Ceigrptraffstatstable
// Table of EIGRP traffic statistics and information
// associated with all EIGRP autonomous systems.
type CISCOEIGRPMIB_Ceigrptraffstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The set of statistics and information for a single EIGRP Autonomous System.
    // The type is slice of
    // CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry.
    Ceigrptraffstatsentry []CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry
}

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetEntityData() *types.CommonEntityData {
    ceigrptraffstatstable.EntityData.YFilter = ceigrptraffstatstable.YFilter
    ceigrptraffstatstable.EntityData.YangName = "cEigrpTraffStatsTable"
    ceigrptraffstatstable.EntityData.BundleName = "cisco_ios_xe"
    ceigrptraffstatstable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    ceigrptraffstatstable.EntityData.SegmentPath = "cEigrpTraffStatsTable"
    ceigrptraffstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrptraffstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrptraffstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrptraffstatstable.EntityData.Children = make(map[string]types.YChild)
    ceigrptraffstatstable.EntityData.Children["cEigrpTraffStatsEntry"] = types.YChild{"Ceigrptraffstatsentry", nil}
    for i := range ceigrptraffstatstable.Ceigrptraffstatsentry {
        ceigrptraffstatstable.EntityData.Children[types.GetSegmentPath(&ceigrptraffstatstable.Ceigrptraffstatsentry[i])] = types.YChild{"Ceigrptraffstatsentry", &ceigrptraffstatstable.Ceigrptraffstatsentry[i]}
    }
    ceigrptraffstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceigrptraffstatstable.EntityData)
}

// CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry
// The set of statistics and information for a single EIGRP
// Autonomous System.
type CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry_Ceigrpvpnid
    Ceigrpvpnid interface{}

    // This attribute is a key. The Autonomous System number which is unique
    // integer per VPN. The type is interface{} with range: 0..4294967295.
    Ceigrpasnumber interface{}

    // The total number of live EIGRP neighbors formed on all interfaces whose IP
    // addresses fall under networks configured in the EIGRP AS. The type is
    // interface{} with range: 0..4294967295.
    Ceigrpnbrcount interface{}

    // The total number Hello packets that have been sent to all EIGRP neighbors
    // formed on all interfaces whose IP addresses fall under networks configured
    // for the EIGRP AS. The type is interface{} with range: 0..4294967295.
    Ceigrphellossent interface{}

    // The total number Hello packets that have been received from all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrphellosrcvd interface{}

    // The total number routing update packets that have been sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpupdatessent interface{}

    // The total number routing update packets that have been received from all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpupdatesrcvd interface{}

    // The total number alternate route query packets that have been sent to all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpqueriessent interface{}

    // The total number alternate route query packets that have been received from
    // all EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpqueriesrcvd interface{}

    // The total number query reply packets that have been sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrprepliessent interface{}

    // The total number query reply packets that have been received from all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrprepliesrcvd interface{}

    // The total number packet acknowledgements that have been sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpackssent interface{}

    // The total number packet acknowledgements that have been received from all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpacksrcvd interface{}

    // The highest number of EIGRP packets in the input queue waiting to be
    // processed internally addressed to this AS. The type is interface{} with
    // range: 0..4294967295.
    Ceigrpinputqhighmark interface{}

    // The number of EIGRP packets dropped from the input queue due to it being
    // full within the AS. The type is interface{} with range: 0..4294967295.
    Ceigrpinputqdrops interface{}

    // The total number of Stuck-In-Active (SIA) query packets sent to all EIGRP
    // neighbors formed on all interfaces whose IP addresses fall under networks
    // configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpsiaqueriessent interface{}

    // The total number of Stuck-In-Active (SIA) query packets received from all
    // EIGRP neighbors formed on all interfaces whose IP addresses fall under
    // networks configured for the EIGRP AS. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpsiaqueriesrcvd interface{}

    // The format of the router-id configured or automatically selected for the
    // EIGRP AS. The type is InetAddressType.
    Ceigrpasrouteridtype interface{}

    // The router-id configured or automatically selected for the EIGRP AS.   Each
    // EIGRP routing process has a unique router-id selected from each autonomous
    // system configured. The format is governed by object cEigrpAsRouterIdType.
    // The type is string with length: 0..255.
    Ceigrpasrouterid interface{}

    // The total number of EIGRP derived routes currently existing in the topology
    // table for the AS. The type is interface{} with range: 0..4294967295.
    Ceigrptoporoutes interface{}

    // Routes in a topology table for an AS are assigned serial numbers and are
    // sequenced internally as they are inserted and deleted.   The serial number
    // of the first route in that internal sequence is called the head serial
    // number. Each AS has its own topology table, and its own serial number
    // space, each of which begins with the value 1. A serial number of zero
    // implies that there are no routes in the topology. The type is interface{}
    // with range: 0..18446744073709551615.
    Ceigrpheadserial interface{}

    // The serial number that would be assigned to the next new or changed route
    // in the topology table for the AS. The type is interface{} with range:
    // 0..18446744073709551615.
    Ceigrpnextserial interface{}

    // When alternate route query packets are sent to adjacent EIGRP peers in an
    // AS, replies are expected.   This object is the total number of outstanding
    // replies expected to queries that have been sent to peers in the current AS.
    // It remains at zero most of the time until an EIGRP route becomes active.
    // The type is interface{} with range: 0..4294967295.
    Ceigrpxmitpendreplies interface{}

    // A dummy is a temporary internal entity used as a place holder in the
    // topology table for an AS. They are not transmitted in routing updates. 
    // This is the total number currently in existence associated with the AS. The
    // type is interface{} with range: 0..4294967295.
    Ceigrpxmitdummies interface{}
}

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetEntityData() *types.CommonEntityData {
    ceigrptraffstatsentry.EntityData.YFilter = ceigrptraffstatsentry.YFilter
    ceigrptraffstatsentry.EntityData.YangName = "cEigrpTraffStatsEntry"
    ceigrptraffstatsentry.EntityData.BundleName = "cisco_ios_xe"
    ceigrptraffstatsentry.EntityData.ParentYangName = "cEigrpTraffStatsTable"
    ceigrptraffstatsentry.EntityData.SegmentPath = "cEigrpTraffStatsEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrptraffstatsentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrptraffstatsentry.Ceigrpasnumber) + "']"
    ceigrptraffstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrptraffstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrptraffstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrptraffstatsentry.EntityData.Children = make(map[string]types.YChild)
    ceigrptraffstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpVpnId"] = types.YLeaf{"Ceigrpvpnid", ceigrptraffstatsentry.Ceigrpvpnid}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpAsNumber"] = types.YLeaf{"Ceigrpasnumber", ceigrptraffstatsentry.Ceigrpasnumber}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpNbrCount"] = types.YLeaf{"Ceigrpnbrcount", ceigrptraffstatsentry.Ceigrpnbrcount}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpHellosSent"] = types.YLeaf{"Ceigrphellossent", ceigrptraffstatsentry.Ceigrphellossent}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpHellosRcvd"] = types.YLeaf{"Ceigrphellosrcvd", ceigrptraffstatsentry.Ceigrphellosrcvd}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpUpdatesSent"] = types.YLeaf{"Ceigrpupdatessent", ceigrptraffstatsentry.Ceigrpupdatessent}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpUpdatesRcvd"] = types.YLeaf{"Ceigrpupdatesrcvd", ceigrptraffstatsentry.Ceigrpupdatesrcvd}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpQueriesSent"] = types.YLeaf{"Ceigrpqueriessent", ceigrptraffstatsentry.Ceigrpqueriessent}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpQueriesRcvd"] = types.YLeaf{"Ceigrpqueriesrcvd", ceigrptraffstatsentry.Ceigrpqueriesrcvd}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpRepliesSent"] = types.YLeaf{"Ceigrprepliessent", ceigrptraffstatsentry.Ceigrprepliessent}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpRepliesRcvd"] = types.YLeaf{"Ceigrprepliesrcvd", ceigrptraffstatsentry.Ceigrprepliesrcvd}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpAcksSent"] = types.YLeaf{"Ceigrpackssent", ceigrptraffstatsentry.Ceigrpackssent}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpAcksRcvd"] = types.YLeaf{"Ceigrpacksrcvd", ceigrptraffstatsentry.Ceigrpacksrcvd}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpInputQHighMark"] = types.YLeaf{"Ceigrpinputqhighmark", ceigrptraffstatsentry.Ceigrpinputqhighmark}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpInputQDrops"] = types.YLeaf{"Ceigrpinputqdrops", ceigrptraffstatsentry.Ceigrpinputqdrops}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpSiaQueriesSent"] = types.YLeaf{"Ceigrpsiaqueriessent", ceigrptraffstatsentry.Ceigrpsiaqueriessent}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpSiaQueriesRcvd"] = types.YLeaf{"Ceigrpsiaqueriesrcvd", ceigrptraffstatsentry.Ceigrpsiaqueriesrcvd}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpAsRouterIdType"] = types.YLeaf{"Ceigrpasrouteridtype", ceigrptraffstatsentry.Ceigrpasrouteridtype}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpAsRouterId"] = types.YLeaf{"Ceigrpasrouterid", ceigrptraffstatsentry.Ceigrpasrouterid}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpTopoRoutes"] = types.YLeaf{"Ceigrptoporoutes", ceigrptraffstatsentry.Ceigrptoporoutes}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpHeadSerial"] = types.YLeaf{"Ceigrpheadserial", ceigrptraffstatsentry.Ceigrpheadserial}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpNextSerial"] = types.YLeaf{"Ceigrpnextserial", ceigrptraffstatsentry.Ceigrpnextserial}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpXmitPendReplies"] = types.YLeaf{"Ceigrpxmitpendreplies", ceigrptraffstatsentry.Ceigrpxmitpendreplies}
    ceigrptraffstatsentry.EntityData.Leafs["cEigrpXmitDummies"] = types.YLeaf{"Ceigrpxmitdummies", ceigrptraffstatsentry.Ceigrpxmitdummies}
    return &(ceigrptraffstatsentry.EntityData)
}

// CISCOEIGRPMIB_Ceigrptopotable
// The table of EIGRP routes and their associated
// attributes for an Autonomous System (AS) configured
// in a VPN is called a topology table.  All route entries in
// the topology table will be indexed by IP network type,
// IP network number and network mask (prefix) size.
type CISCOEIGRPMIB_Ceigrptopotable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The entry for a single EIGRP topology table in the given AS. The type is
    // slice of CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry.
    Ceigrptopoentry []CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry
}

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetEntityData() *types.CommonEntityData {
    ceigrptopotable.EntityData.YFilter = ceigrptopotable.YFilter
    ceigrptopotable.EntityData.YangName = "cEigrpTopoTable"
    ceigrptopotable.EntityData.BundleName = "cisco_ios_xe"
    ceigrptopotable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    ceigrptopotable.EntityData.SegmentPath = "cEigrpTopoTable"
    ceigrptopotable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrptopotable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrptopotable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrptopotable.EntityData.Children = make(map[string]types.YChild)
    ceigrptopotable.EntityData.Children["cEigrpTopoEntry"] = types.YChild{"Ceigrptopoentry", nil}
    for i := range ceigrptopotable.Ceigrptopoentry {
        ceigrptopotable.EntityData.Children[types.GetSegmentPath(&ceigrptopotable.Ceigrptopoentry[i])] = types.YChild{"Ceigrptopoentry", &ceigrptopotable.Ceigrptopoentry[i]}
    }
    ceigrptopotable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceigrptopotable.EntityData)
}

// CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry
// The entry for a single EIGRP topology table in the given
// AS.
type CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry_Ceigrpvpnid
    Ceigrpvpnid interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry_Ceigrpasnumber
    Ceigrpasnumber interface{}

    // This attribute is a key. The format of the destination IP network number
    // for a single route in the topology table in the AS specified in
    // cEigrpDestNet. The type is InetAddressType.
    Ceigrpdestnettype interface{}

    // This attribute is a key. The destination IP network number for a single
    // route in the topology table in the AS.  The format is governed by object
    // cEigrpDestNetType. The type is string with length: 0..255.
    Ceigrpdestnet interface{}

    // This attribute is a key. The prefix length associated with the destination
    // IP network address for a single route in the topology table in the AS.  The
    // format is governed by the object cEigrpDestNetType. The type is interface{}
    // with range: 0..2040.
    Ceigrpdestnetprefixlen interface{}

    // A value of true(1) indicates the route to the destination network has
    // failed and an active (query) search for an alternative path is in progress.
    // A value of false(2) indicates the route is stable (passive). The type is
    // bool.
    Ceigrpactive interface{}

    // A value of true(1) indicates that that this route which is in active state
    // (cEigrpActive = true(1)) has not received any replies to queries for
    // alternate paths, and a second EIGRP route query, called a stuck-in-active
    // query, has now been sent. The type is bool.
    Ceigrpstuckinactive interface{}

    // A successor is the next routing hop for a path to the destination IP
    // network number for a single route in the topology table in the AS.  There
    // can be several potential successors if there are multiple paths to the
    // destination.   This is the total number of successors for a topology entry.
    // The type is interface{} with range: 0..4294967295.
    Ceigrpdestsuccessors interface{}

    // The feasibility (best) distance is the minimum distance from this router to
    // the destination IP network in this topology entry.  The feasibility
    // distance is used in determining the best successor for a path to the
    // destination network. The type is interface{} with range: 0..4294967295.
    Ceigrpfdistance interface{}

    // This is a text string describing the internal origin of the EIGRP route
    // represented by the topology entry. The type is string.
    Ceigrprouteorigintype interface{}

    // The format of the IP address defined as the origin of this topology route
    // entry. The type is InetAddressType.
    Ceigrprouteoriginaddrtype interface{}

    // If the origin of the topology route entry is external to this router, then
    // this object is the IP address of the router from which it originated.  The
    // format  is governed by object cEigrpRouteOriginAddrType. The type is string
    // with length: 0..255.
    Ceigrprouteoriginaddr interface{}

    // The format of the next hop IP address for the route represented by the
    // topology entry. The type is InetAddressType.
    Ceigrpnexthopaddresstype interface{}

    // This is the next hop IP address for the route represented by the topology
    // entry.   The next hop is where network traffic will be routed to in order
    // to reach the destination network for this topology entry.  The format is
    // governed by cEigrpNextHopAddressType. The type is string with length:
    // 0..255.
    Ceigrpnexthopaddress interface{}

    // The interface through which the next hop IP address is reached to send
    // network traffic to the destination network represented by the topology
    // entry. The type is string.
    Ceigrpnexthopinterface interface{}

    // The computed distance to the destination network entry from this router.
    // The type is interface{} with range: 0..4294967295.
    Ceigrpdistance interface{}

    // The computed distance to the destination network in the topology entry
    // reported to this router by the originator of this route. The type is
    // interface{} with range: 0..4294967295.
    Ceigrpreportdistance interface{}
}

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetEntityData() *types.CommonEntityData {
    ceigrptopoentry.EntityData.YFilter = ceigrptopoentry.YFilter
    ceigrptopoentry.EntityData.YangName = "cEigrpTopoEntry"
    ceigrptopoentry.EntityData.BundleName = "cisco_ios_xe"
    ceigrptopoentry.EntityData.ParentYangName = "cEigrpTopoTable"
    ceigrptopoentry.EntityData.SegmentPath = "cEigrpTopoEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpasnumber) + "']" + "[cEigrpDestNetType='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpdestnettype) + "']" + "[cEigrpDestNet='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpdestnet) + "']" + "[cEigrpDestNetPrefixLen='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpdestnetprefixlen) + "']"
    ceigrptopoentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrptopoentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrptopoentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrptopoentry.EntityData.Children = make(map[string]types.YChild)
    ceigrptopoentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceigrptopoentry.EntityData.Leafs["cEigrpVpnId"] = types.YLeaf{"Ceigrpvpnid", ceigrptopoentry.Ceigrpvpnid}
    ceigrptopoentry.EntityData.Leafs["cEigrpAsNumber"] = types.YLeaf{"Ceigrpasnumber", ceigrptopoentry.Ceigrpasnumber}
    ceigrptopoentry.EntityData.Leafs["cEigrpDestNetType"] = types.YLeaf{"Ceigrpdestnettype", ceigrptopoentry.Ceigrpdestnettype}
    ceigrptopoentry.EntityData.Leafs["cEigrpDestNet"] = types.YLeaf{"Ceigrpdestnet", ceigrptopoentry.Ceigrpdestnet}
    ceigrptopoentry.EntityData.Leafs["cEigrpDestNetPrefixLen"] = types.YLeaf{"Ceigrpdestnetprefixlen", ceigrptopoentry.Ceigrpdestnetprefixlen}
    ceigrptopoentry.EntityData.Leafs["cEigrpActive"] = types.YLeaf{"Ceigrpactive", ceigrptopoentry.Ceigrpactive}
    ceigrptopoentry.EntityData.Leafs["cEigrpStuckInActive"] = types.YLeaf{"Ceigrpstuckinactive", ceigrptopoentry.Ceigrpstuckinactive}
    ceigrptopoentry.EntityData.Leafs["cEigrpDestSuccessors"] = types.YLeaf{"Ceigrpdestsuccessors", ceigrptopoentry.Ceigrpdestsuccessors}
    ceigrptopoentry.EntityData.Leafs["cEigrpFdistance"] = types.YLeaf{"Ceigrpfdistance", ceigrptopoentry.Ceigrpfdistance}
    ceigrptopoentry.EntityData.Leafs["cEigrpRouteOriginType"] = types.YLeaf{"Ceigrprouteorigintype", ceigrptopoentry.Ceigrprouteorigintype}
    ceigrptopoentry.EntityData.Leafs["cEigrpRouteOriginAddrType"] = types.YLeaf{"Ceigrprouteoriginaddrtype", ceigrptopoentry.Ceigrprouteoriginaddrtype}
    ceigrptopoentry.EntityData.Leafs["cEigrpRouteOriginAddr"] = types.YLeaf{"Ceigrprouteoriginaddr", ceigrptopoentry.Ceigrprouteoriginaddr}
    ceigrptopoentry.EntityData.Leafs["cEigrpNextHopAddressType"] = types.YLeaf{"Ceigrpnexthopaddresstype", ceigrptopoentry.Ceigrpnexthopaddresstype}
    ceigrptopoentry.EntityData.Leafs["cEigrpNextHopAddress"] = types.YLeaf{"Ceigrpnexthopaddress", ceigrptopoentry.Ceigrpnexthopaddress}
    ceigrptopoentry.EntityData.Leafs["cEigrpNextHopInterface"] = types.YLeaf{"Ceigrpnexthopinterface", ceigrptopoentry.Ceigrpnexthopinterface}
    ceigrptopoentry.EntityData.Leafs["cEigrpDistance"] = types.YLeaf{"Ceigrpdistance", ceigrptopoentry.Ceigrpdistance}
    ceigrptopoentry.EntityData.Leafs["cEigrpReportDistance"] = types.YLeaf{"Ceigrpreportdistance", ceigrptopoentry.Ceigrpreportdistance}
    return &(ceigrptopoentry.EntityData)
}

// CISCOEIGRPMIB_Ceigrppeertable
// The table of established EIGRP peers (neighbors) in the
// selected autonomous system.   Peers are indexed by their
// unique internal handle id, as well as the AS number and
// VPN id.   The peer entry is removed from the table if
// the peer is declared down.
type CISCOEIGRPMIB_Ceigrppeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics and operational parameters for a single peer in the AS. The type
    // is slice of CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry.
    Ceigrppeerentry []CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry
}

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetEntityData() *types.CommonEntityData {
    ceigrppeertable.EntityData.YFilter = ceigrppeertable.YFilter
    ceigrppeertable.EntityData.YangName = "cEigrpPeerTable"
    ceigrppeertable.EntityData.BundleName = "cisco_ios_xe"
    ceigrppeertable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    ceigrppeertable.EntityData.SegmentPath = "cEigrpPeerTable"
    ceigrppeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrppeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrppeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrppeertable.EntityData.Children = make(map[string]types.YChild)
    ceigrppeertable.EntityData.Children["cEigrpPeerEntry"] = types.YChild{"Ceigrppeerentry", nil}
    for i := range ceigrppeertable.Ceigrppeerentry {
        ceigrppeertable.EntityData.Children[types.GetSegmentPath(&ceigrppeertable.Ceigrppeerentry[i])] = types.YChild{"Ceigrppeerentry", &ceigrppeertable.Ceigrppeerentry[i]}
    }
    ceigrppeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceigrppeertable.EntityData)
}

// CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry
// Statistics and operational parameters for a single peer
// in the AS.
type CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry_Ceigrpvpnid
    Ceigrpvpnid interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry_Ceigrpasnumber
    Ceigrpasnumber interface{}

    // This attribute is a key. The unique internal identifier for the peer in the
    // AS. This is a unique value among peer entries in a selected table. The type
    // is interface{} with range: 0..4294967295.
    Ceigrphandle interface{}

    // The format of the remote source IP address used by the peer to establish
    // the EIGRP adjacency with this router. The type is InetAddressType.
    Ceigrppeeraddrtype interface{}

    // The source IP address used by the peer to establish the EIGRP adjacency
    // with this router.  The format is governed by object cEigrpPeerAddrType. The
    // type is string with length: 0..255.
    Ceigrppeeraddr interface{}

    // The ifIndex of the interface on this router through which this peer can be
    // reached. The type is interface{} with range: 0..2147483647.
    Ceigrppeerifindex interface{}

    // The count-down timer indicating how much time must pass without receiving a
    // hello packet from this EIGRP peer before this router declares the peer
    // down. A peer declared as down is removed from the table and is no longer
    // visible. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Ceigrpholdtime interface{}

    // The elapsed time since the EIGRP adjacency was first established with the
    // peer. The type is string.
    Ceigrpuptime interface{}

    // The computed smooth round trip time for packets to and from the peer. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Ceigrpsrtt interface{}

    // The computed retransmission timeout for the peer. This value is computed
    // over time as packets are sent to the peer and acknowledgements are received
    // from it, and is the amount of time to wait before resending a packet from
    // the retransmission queue to the peer when an expected acknowledgement has
    // not been received. The type is interface{} with range: 0..4294967295. Units
    // are milliseconds.
    Ceigrprto interface{}

    // The number of any EIGRP packets currently enqueued waiting to be sent to
    // this peer. The type is interface{} with range: 0..4294967295.
    Ceigrppktsenqueued interface{}

    // All transmitted EIGRP packets have a sequence number assigned. This is the
    // sequence number of the last EIGRP packet sent to this peer. The type is
    // interface{} with range: 0..4294967295.
    Ceigrplastseq interface{}

    // The EIGRP version information reported by the remote peer. The type is
    // string.
    Ceigrpversion interface{}

    // The cumulative number of retransmissions to this peer during the period
    // that the peer adjacency has remained up. The type is interface{} with
    // range: 0..4294967295.
    Ceigrpretrans interface{}

    // The number of times the current unacknowledged packet has been retried,
    // i.e. resent to this peer to be acknowledged. The type is interface{} with
    // range: 0..4294967295.
    Ceigrpretries interface{}
}

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetEntityData() *types.CommonEntityData {
    ceigrppeerentry.EntityData.YFilter = ceigrppeerentry.YFilter
    ceigrppeerentry.EntityData.YangName = "cEigrpPeerEntry"
    ceigrppeerentry.EntityData.BundleName = "cisco_ios_xe"
    ceigrppeerentry.EntityData.ParentYangName = "cEigrpPeerTable"
    ceigrppeerentry.EntityData.SegmentPath = "cEigrpPeerEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrppeerentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrppeerentry.Ceigrpasnumber) + "']" + "[cEigrpHandle='" + fmt.Sprintf("%v", ceigrppeerentry.Ceigrphandle) + "']"
    ceigrppeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrppeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrppeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrppeerentry.EntityData.Children = make(map[string]types.YChild)
    ceigrppeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceigrppeerentry.EntityData.Leafs["cEigrpVpnId"] = types.YLeaf{"Ceigrpvpnid", ceigrppeerentry.Ceigrpvpnid}
    ceigrppeerentry.EntityData.Leafs["cEigrpAsNumber"] = types.YLeaf{"Ceigrpasnumber", ceigrppeerentry.Ceigrpasnumber}
    ceigrppeerentry.EntityData.Leafs["cEigrpHandle"] = types.YLeaf{"Ceigrphandle", ceigrppeerentry.Ceigrphandle}
    ceigrppeerentry.EntityData.Leafs["cEigrpPeerAddrType"] = types.YLeaf{"Ceigrppeeraddrtype", ceigrppeerentry.Ceigrppeeraddrtype}
    ceigrppeerentry.EntityData.Leafs["cEigrpPeerAddr"] = types.YLeaf{"Ceigrppeeraddr", ceigrppeerentry.Ceigrppeeraddr}
    ceigrppeerentry.EntityData.Leafs["cEigrpPeerIfIndex"] = types.YLeaf{"Ceigrppeerifindex", ceigrppeerentry.Ceigrppeerifindex}
    ceigrppeerentry.EntityData.Leafs["cEigrpHoldTime"] = types.YLeaf{"Ceigrpholdtime", ceigrppeerentry.Ceigrpholdtime}
    ceigrppeerentry.EntityData.Leafs["cEigrpUpTime"] = types.YLeaf{"Ceigrpuptime", ceigrppeerentry.Ceigrpuptime}
    ceigrppeerentry.EntityData.Leafs["cEigrpSrtt"] = types.YLeaf{"Ceigrpsrtt", ceigrppeerentry.Ceigrpsrtt}
    ceigrppeerentry.EntityData.Leafs["cEigrpRto"] = types.YLeaf{"Ceigrprto", ceigrppeerentry.Ceigrprto}
    ceigrppeerentry.EntityData.Leafs["cEigrpPktsEnqueued"] = types.YLeaf{"Ceigrppktsenqueued", ceigrppeerentry.Ceigrppktsenqueued}
    ceigrppeerentry.EntityData.Leafs["cEigrpLastSeq"] = types.YLeaf{"Ceigrplastseq", ceigrppeerentry.Ceigrplastseq}
    ceigrppeerentry.EntityData.Leafs["cEigrpVersion"] = types.YLeaf{"Ceigrpversion", ceigrppeerentry.Ceigrpversion}
    ceigrppeerentry.EntityData.Leafs["cEigrpRetrans"] = types.YLeaf{"Ceigrpretrans", ceigrppeerentry.Ceigrpretrans}
    ceigrppeerentry.EntityData.Leafs["cEigrpRetries"] = types.YLeaf{"Ceigrpretries", ceigrppeerentry.Ceigrpretries}
    return &(ceigrppeerentry.EntityData)
}

// CISCOEIGRPMIB_Ceigrpinterfacetable
// The table of interfaces over which EIGRP is running, and
// their associated statistics.   This table is independent
// of whether any peer adjacencies have been formed over
// the interfaces or not.   Interfaces running EIGRP are
// determined by whether their assigned IP addresses fall
// within configured EIGRP network statements.
type CISCOEIGRPMIB_Ceigrpinterfacetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information for a single interface running EIGRP in the AS and VPN. The
    // type is slice of CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry.
    Ceigrpinterfaceentry []CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry
}

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetEntityData() *types.CommonEntityData {
    ceigrpinterfacetable.EntityData.YFilter = ceigrpinterfacetable.YFilter
    ceigrpinterfacetable.EntityData.YangName = "cEigrpInterfaceTable"
    ceigrpinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    ceigrpinterfacetable.EntityData.ParentYangName = "CISCO-EIGRP-MIB"
    ceigrpinterfacetable.EntityData.SegmentPath = "cEigrpInterfaceTable"
    ceigrpinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrpinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrpinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrpinterfacetable.EntityData.Children = make(map[string]types.YChild)
    ceigrpinterfacetable.EntityData.Children["cEigrpInterfaceEntry"] = types.YChild{"Ceigrpinterfaceentry", nil}
    for i := range ceigrpinterfacetable.Ceigrpinterfaceentry {
        ceigrpinterfacetable.EntityData.Children[types.GetSegmentPath(&ceigrpinterfacetable.Ceigrpinterfaceentry[i])] = types.YChild{"Ceigrpinterfaceentry", &ceigrpinterfacetable.Ceigrpinterfaceentry[i]}
    }
    ceigrpinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceigrpinterfacetable.EntityData)
}

// CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry
// Information for a single interface running EIGRP in the
// AS and VPN.
type CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry_Ceigrpvpnid
    Ceigrpvpnid interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_eigrp_mib.CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry_Ceigrpasnumber
    Ceigrpasnumber interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of EIGRP adjacencies currently formed with peers reached through
    // this interface. The type is interface{} with range: 0..4294967295.
    Ceigrppeercount interface{}

    // The number of EIGRP packets currently waiting in the reliable transport
    // (acknowledgement-required)  transmission queue to be sent to a peer. The
    // type is interface{} with range: 0..4294967295.
    Ceigrpxmitreliableq interface{}

    // The number EIGRP of packets currently waiting in the unreliable transport
    // (no acknowledgement required) transmission queue. The type is interface{}
    // with range: 0..4294967295.
    Ceigrpxmitunreliableq interface{}

    // The average of all the computed smooth round trip time values for a packet
    // to and from all peers established on this interface. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Ceigrpmeansrtt interface{}

    // The configured time interval between EIGRP packet transmissions on the
    // interface when the reliable transport method is used. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Ceigrppacingreliable interface{}

    // The configured time interval between EIGRP packet transmissions on the
    // interface when the unreliable transport method is used. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Ceigrppacingunreliable interface{}

    // The configured multicast flow control timer value for this interface. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Ceigrpmflowtimer interface{}

    // The number of queued EIGRP routing updates awaiting transmission on this
    // interface. The type is interface{} with range: 0..4294967295.
    Ceigrppendingroutes interface{}

    // The configured time interval between Hello packet transmissions for this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Ceigrphellointerval interface{}

    // The serial number of the next EIGRP packet that is to be queued for
    // transmission on this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Ceigrpxmitnextserial interface{}

    // The total number of unreliable (no acknowledgement required) EIGRP
    // multicast packets sent on this interface. The type is interface{} with
    // range: 0..4294967295.
    Ceigrpumcasts interface{}

    // The total number of reliable (acknowledgement required) EIGRP multicast
    // packets sent on this interface. The type is interface{} with range:
    // 0..4294967295.
    Ceigrprmcasts interface{}

    // The total number of unreliable (no acknowledgement required) EIGRP unicast
    // packets sent on this interface. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpuucasts interface{}

    // The total number of reliable (acknowledgement required) unicast packets
    // sent on this interface. The type is interface{} with range: 0..4294967295.
    Ceigrprucasts interface{}

    // The total number of EIGRP multicast exception transmissions that have
    // occurred on this interface. The type is interface{} with range:
    // 0..4294967295.
    Ceigrpmcastexcepts interface{}

    // The total number EIGRP Conditional-Receive packets sent on this interface.
    // The type is interface{} with range: 0..4294967295.
    Ceigrpcrpkts interface{}

    // The total number of individual EIGRP acknowledgement packets that have been
    // suppressed and combined in an already enqueued outbound reliable packet on
    // this interface. The type is interface{} with range: 0..4294967295.
    Ceigrpackssuppressed interface{}

    // The total number EIGRP packet retransmissions sent on the interface. The
    // type is interface{} with range: 0..4294967295.
    Ceigrpretranssent interface{}

    // The total number of out-of-sequence EIGRP packets received. The type is
    // interface{} with range: 0..4294967295.
    Ceigrpoosrvcd interface{}

    // The EIGRP authentication mode of the interface. none  :  no authentication
    // enabled on the interface md5   :  MD5 authentication enabled on the
    // interface. The type is Ceigrpauthmode.
    Ceigrpauthmode interface{}

    // The name of the authentication key-chain configured on this interface.  
    // The key-chain is a reference to which set of secret keys are to be accessed
    // in order to determine which secret key string to use.  The key chain name
    // is not the secret key string password and can also be used in other routing
    // protocols, such as RIP and ISIS. The type is string.
    Ceigrpauthkeychain interface{}
}

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetEntityData() *types.CommonEntityData {
    ceigrpinterfaceentry.EntityData.YFilter = ceigrpinterfaceentry.YFilter
    ceigrpinterfaceentry.EntityData.YangName = "cEigrpInterfaceEntry"
    ceigrpinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    ceigrpinterfaceentry.EntityData.ParentYangName = "cEigrpInterfaceTable"
    ceigrpinterfaceentry.EntityData.SegmentPath = "cEigrpInterfaceEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrpinterfaceentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrpinterfaceentry.Ceigrpasnumber) + "']" + "[ifIndex='" + fmt.Sprintf("%v", ceigrpinterfaceentry.Ifindex) + "']"
    ceigrpinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceigrpinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceigrpinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceigrpinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    ceigrpinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpVpnId"] = types.YLeaf{"Ceigrpvpnid", ceigrpinterfaceentry.Ceigrpvpnid}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpAsNumber"] = types.YLeaf{"Ceigrpasnumber", ceigrpinterfaceentry.Ceigrpasnumber}
    ceigrpinterfaceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", ceigrpinterfaceentry.Ifindex}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpPeerCount"] = types.YLeaf{"Ceigrppeercount", ceigrpinterfaceentry.Ceigrppeercount}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpXmitReliableQ"] = types.YLeaf{"Ceigrpxmitreliableq", ceigrpinterfaceentry.Ceigrpxmitreliableq}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpXmitUnreliableQ"] = types.YLeaf{"Ceigrpxmitunreliableq", ceigrpinterfaceentry.Ceigrpxmitunreliableq}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpMeanSrtt"] = types.YLeaf{"Ceigrpmeansrtt", ceigrpinterfaceentry.Ceigrpmeansrtt}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpPacingReliable"] = types.YLeaf{"Ceigrppacingreliable", ceigrpinterfaceentry.Ceigrppacingreliable}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpPacingUnreliable"] = types.YLeaf{"Ceigrppacingunreliable", ceigrpinterfaceentry.Ceigrppacingunreliable}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpMFlowTimer"] = types.YLeaf{"Ceigrpmflowtimer", ceigrpinterfaceentry.Ceigrpmflowtimer}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpPendingRoutes"] = types.YLeaf{"Ceigrppendingroutes", ceigrpinterfaceentry.Ceigrppendingroutes}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpHelloInterval"] = types.YLeaf{"Ceigrphellointerval", ceigrpinterfaceentry.Ceigrphellointerval}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpXmitNextSerial"] = types.YLeaf{"Ceigrpxmitnextserial", ceigrpinterfaceentry.Ceigrpxmitnextserial}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpUMcasts"] = types.YLeaf{"Ceigrpumcasts", ceigrpinterfaceentry.Ceigrpumcasts}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpRMcasts"] = types.YLeaf{"Ceigrprmcasts", ceigrpinterfaceentry.Ceigrprmcasts}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpUUcasts"] = types.YLeaf{"Ceigrpuucasts", ceigrpinterfaceentry.Ceigrpuucasts}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpRUcasts"] = types.YLeaf{"Ceigrprucasts", ceigrpinterfaceentry.Ceigrprucasts}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpMcastExcepts"] = types.YLeaf{"Ceigrpmcastexcepts", ceigrpinterfaceentry.Ceigrpmcastexcepts}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpCRpkts"] = types.YLeaf{"Ceigrpcrpkts", ceigrpinterfaceentry.Ceigrpcrpkts}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpAcksSuppressed"] = types.YLeaf{"Ceigrpackssuppressed", ceigrpinterfaceentry.Ceigrpackssuppressed}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpRetransSent"] = types.YLeaf{"Ceigrpretranssent", ceigrpinterfaceentry.Ceigrpretranssent}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpOOSrvcd"] = types.YLeaf{"Ceigrpoosrvcd", ceigrpinterfaceentry.Ceigrpoosrvcd}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpAuthMode"] = types.YLeaf{"Ceigrpauthmode", ceigrpinterfaceentry.Ceigrpauthmode}
    ceigrpinterfaceentry.EntityData.Leafs["cEigrpAuthKeyChain"] = types.YLeaf{"Ceigrpauthkeychain", ceigrpinterfaceentry.Ceigrpauthkeychain}
    return &(ceigrpinterfaceentry.EntityData)
}

// CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode represents md5   :  MD5 authentication enabled on the interface
type CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode string

const (
    CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode_none CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode = "none"

    CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode_md5 CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode = "md5"
)

