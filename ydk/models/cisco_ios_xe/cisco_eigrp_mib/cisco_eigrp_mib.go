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
    parent types.Entity
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

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetFilter() yfilter.YFilter { return cISCOEIGRPMIB.YFilter }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) SetFilter(yf yfilter.YFilter) { cISCOEIGRPMIB.YFilter = yf }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetGoName(yname string) string {
    if yname == "cEigrpVpnTable" { return "Ceigrpvpntable" }
    if yname == "cEigrpTraffStatsTable" { return "Ceigrptraffstatstable" }
    if yname == "cEigrpTopoTable" { return "Ceigrptopotable" }
    if yname == "cEigrpPeerTable" { return "Ceigrppeertable" }
    if yname == "cEigrpInterfaceTable" { return "Ceigrpinterfacetable" }
    return ""
}

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetSegmentPath() string {
    return "CISCO-EIGRP-MIB:CISCO-EIGRP-MIB"
}

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cEigrpVpnTable" {
        return &cISCOEIGRPMIB.Ceigrpvpntable
    }
    if childYangName == "cEigrpTraffStatsTable" {
        return &cISCOEIGRPMIB.Ceigrptraffstatstable
    }
    if childYangName == "cEigrpTopoTable" {
        return &cISCOEIGRPMIB.Ceigrptopotable
    }
    if childYangName == "cEigrpPeerTable" {
        return &cISCOEIGRPMIB.Ceigrppeertable
    }
    if childYangName == "cEigrpInterfaceTable" {
        return &cISCOEIGRPMIB.Ceigrpinterfacetable
    }
    return nil
}

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cEigrpVpnTable"] = &cISCOEIGRPMIB.Ceigrpvpntable
    children["cEigrpTraffStatsTable"] = &cISCOEIGRPMIB.Ceigrptraffstatstable
    children["cEigrpTopoTable"] = &cISCOEIGRPMIB.Ceigrptopotable
    children["cEigrpPeerTable"] = &cISCOEIGRPMIB.Ceigrppeertable
    children["cEigrpInterfaceTable"] = &cISCOEIGRPMIB.Ceigrpinterfacetable
    return children
}

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetYangName() string { return "CISCO-EIGRP-MIB" }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) SetParent(parent types.Entity) { cISCOEIGRPMIB.parent = parent }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetParent() types.Entity { return cISCOEIGRPMIB.parent }

func (cISCOEIGRPMIB *CISCOEIGRPMIB) GetParentYangName() string { return "CISCO-EIGRP-MIB" }

// CISCOEIGRPMIB_Ceigrpvpntable
// This table contains information on those VPN's configured
// to run EIGRP.  The VPN creation on a router is independent
// of the routing protocol to be used over it.   A VPN is
// given a name and has a dedicated routing table associated
// with it.  This routing table is identified internally
// by a unique integer value.
type CISCOEIGRPMIB_Ceigrpvpntable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information relating to a single VPN which is configured to run EIGRP. The
    // type is slice of CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry.
    Ceigrpvpnentry []CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry
}

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetFilter() yfilter.YFilter { return ceigrpvpntable.YFilter }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) SetFilter(yf yfilter.YFilter) { ceigrpvpntable.YFilter = yf }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetGoName(yname string) string {
    if yname == "cEigrpVpnEntry" { return "Ceigrpvpnentry" }
    return ""
}

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetSegmentPath() string {
    return "cEigrpVpnTable"
}

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cEigrpVpnEntry" {
        for _, c := range ceigrpvpntable.Ceigrpvpnentry {
            if ceigrpvpntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry{}
        ceigrpvpntable.Ceigrpvpnentry = append(ceigrpvpntable.Ceigrpvpnentry, child)
        return &ceigrpvpntable.Ceigrpvpnentry[len(ceigrpvpntable.Ceigrpvpnentry)-1]
    }
    return nil
}

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceigrpvpntable.Ceigrpvpnentry {
        children[ceigrpvpntable.Ceigrpvpnentry[i].GetSegmentPath()] = &ceigrpvpntable.Ceigrpvpnentry[i]
    }
    return children
}

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetYangName() string { return "cEigrpVpnTable" }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) SetParent(parent types.Entity) { ceigrpvpntable.parent = parent }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetParent() types.Entity { return ceigrpvpntable.parent }

func (ceigrpvpntable *CISCOEIGRPMIB_Ceigrpvpntable) GetParentYangName() string { return "CISCO-EIGRP-MIB" }

// CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry
// Information relating to a single VPN which is configured
// to run EIGRP.
type CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The unique VPN identifier.  This is a unique
    // integer relative to all other VPN's defined on the router.  It also
    // identifies internally the routing table instance. The type is interface{}
    // with range: 0..4294967295.
    Ceigrpvpnid interface{}

    // The name given to the VPN. The type is string.
    Ceigrpvpnname interface{}
}

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetFilter() yfilter.YFilter { return ceigrpvpnentry.YFilter }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) SetFilter(yf yfilter.YFilter) { ceigrpvpnentry.YFilter = yf }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetGoName(yname string) string {
    if yname == "cEigrpVpnId" { return "Ceigrpvpnid" }
    if yname == "cEigrpVpnName" { return "Ceigrpvpnname" }
    return ""
}

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetSegmentPath() string {
    return "cEigrpVpnEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrpvpnentry.Ceigrpvpnid) + "']"
}

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cEigrpVpnId"] = ceigrpvpnentry.Ceigrpvpnid
    leafs["cEigrpVpnName"] = ceigrpvpnentry.Ceigrpvpnname
    return leafs
}

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetYangName() string { return "cEigrpVpnEntry" }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) SetParent(parent types.Entity) { ceigrpvpnentry.parent = parent }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetParent() types.Entity { return ceigrpvpnentry.parent }

func (ceigrpvpnentry *CISCOEIGRPMIB_Ceigrpvpntable_Ceigrpvpnentry) GetParentYangName() string { return "cEigrpVpnTable" }

// CISCOEIGRPMIB_Ceigrptraffstatstable
// Table of EIGRP traffic statistics and information
// associated with all EIGRP autonomous systems.
type CISCOEIGRPMIB_Ceigrptraffstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The set of statistics and information for a single EIGRP Autonomous System.
    // The type is slice of
    // CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry.
    Ceigrptraffstatsentry []CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry
}

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetFilter() yfilter.YFilter { return ceigrptraffstatstable.YFilter }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) SetFilter(yf yfilter.YFilter) { ceigrptraffstatstable.YFilter = yf }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetGoName(yname string) string {
    if yname == "cEigrpTraffStatsEntry" { return "Ceigrptraffstatsentry" }
    return ""
}

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetSegmentPath() string {
    return "cEigrpTraffStatsTable"
}

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cEigrpTraffStatsEntry" {
        for _, c := range ceigrptraffstatstable.Ceigrptraffstatsentry {
            if ceigrptraffstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry{}
        ceigrptraffstatstable.Ceigrptraffstatsentry = append(ceigrptraffstatstable.Ceigrptraffstatsentry, child)
        return &ceigrptraffstatstable.Ceigrptraffstatsentry[len(ceigrptraffstatstable.Ceigrptraffstatsentry)-1]
    }
    return nil
}

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceigrptraffstatstable.Ceigrptraffstatsentry {
        children[ceigrptraffstatstable.Ceigrptraffstatsentry[i].GetSegmentPath()] = &ceigrptraffstatstable.Ceigrptraffstatsentry[i]
    }
    return children
}

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetYangName() string { return "cEigrpTraffStatsTable" }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) SetParent(parent types.Entity) { ceigrptraffstatstable.parent = parent }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetParent() types.Entity { return ceigrptraffstatstable.parent }

func (ceigrptraffstatstable *CISCOEIGRPMIB_Ceigrptraffstatstable) GetParentYangName() string { return "CISCO-EIGRP-MIB" }

// CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry
// The set of statistics and information for a single EIGRP
// Autonomous System.
type CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry struct {
    parent types.Entity
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

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetFilter() yfilter.YFilter { return ceigrptraffstatsentry.YFilter }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) SetFilter(yf yfilter.YFilter) { ceigrptraffstatsentry.YFilter = yf }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetGoName(yname string) string {
    if yname == "cEigrpVpnId" { return "Ceigrpvpnid" }
    if yname == "cEigrpAsNumber" { return "Ceigrpasnumber" }
    if yname == "cEigrpNbrCount" { return "Ceigrpnbrcount" }
    if yname == "cEigrpHellosSent" { return "Ceigrphellossent" }
    if yname == "cEigrpHellosRcvd" { return "Ceigrphellosrcvd" }
    if yname == "cEigrpUpdatesSent" { return "Ceigrpupdatessent" }
    if yname == "cEigrpUpdatesRcvd" { return "Ceigrpupdatesrcvd" }
    if yname == "cEigrpQueriesSent" { return "Ceigrpqueriessent" }
    if yname == "cEigrpQueriesRcvd" { return "Ceigrpqueriesrcvd" }
    if yname == "cEigrpRepliesSent" { return "Ceigrprepliessent" }
    if yname == "cEigrpRepliesRcvd" { return "Ceigrprepliesrcvd" }
    if yname == "cEigrpAcksSent" { return "Ceigrpackssent" }
    if yname == "cEigrpAcksRcvd" { return "Ceigrpacksrcvd" }
    if yname == "cEigrpInputQHighMark" { return "Ceigrpinputqhighmark" }
    if yname == "cEigrpInputQDrops" { return "Ceigrpinputqdrops" }
    if yname == "cEigrpSiaQueriesSent" { return "Ceigrpsiaqueriessent" }
    if yname == "cEigrpSiaQueriesRcvd" { return "Ceigrpsiaqueriesrcvd" }
    if yname == "cEigrpAsRouterIdType" { return "Ceigrpasrouteridtype" }
    if yname == "cEigrpAsRouterId" { return "Ceigrpasrouterid" }
    if yname == "cEigrpTopoRoutes" { return "Ceigrptoporoutes" }
    if yname == "cEigrpHeadSerial" { return "Ceigrpheadserial" }
    if yname == "cEigrpNextSerial" { return "Ceigrpnextserial" }
    if yname == "cEigrpXmitPendReplies" { return "Ceigrpxmitpendreplies" }
    if yname == "cEigrpXmitDummies" { return "Ceigrpxmitdummies" }
    return ""
}

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetSegmentPath() string {
    return "cEigrpTraffStatsEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrptraffstatsentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrptraffstatsentry.Ceigrpasnumber) + "']"
}

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cEigrpVpnId"] = ceigrptraffstatsentry.Ceigrpvpnid
    leafs["cEigrpAsNumber"] = ceigrptraffstatsentry.Ceigrpasnumber
    leafs["cEigrpNbrCount"] = ceigrptraffstatsentry.Ceigrpnbrcount
    leafs["cEigrpHellosSent"] = ceigrptraffstatsentry.Ceigrphellossent
    leafs["cEigrpHellosRcvd"] = ceigrptraffstatsentry.Ceigrphellosrcvd
    leafs["cEigrpUpdatesSent"] = ceigrptraffstatsentry.Ceigrpupdatessent
    leafs["cEigrpUpdatesRcvd"] = ceigrptraffstatsentry.Ceigrpupdatesrcvd
    leafs["cEigrpQueriesSent"] = ceigrptraffstatsentry.Ceigrpqueriessent
    leafs["cEigrpQueriesRcvd"] = ceigrptraffstatsentry.Ceigrpqueriesrcvd
    leafs["cEigrpRepliesSent"] = ceigrptraffstatsentry.Ceigrprepliessent
    leafs["cEigrpRepliesRcvd"] = ceigrptraffstatsentry.Ceigrprepliesrcvd
    leafs["cEigrpAcksSent"] = ceigrptraffstatsentry.Ceigrpackssent
    leafs["cEigrpAcksRcvd"] = ceigrptraffstatsentry.Ceigrpacksrcvd
    leafs["cEigrpInputQHighMark"] = ceigrptraffstatsentry.Ceigrpinputqhighmark
    leafs["cEigrpInputQDrops"] = ceigrptraffstatsentry.Ceigrpinputqdrops
    leafs["cEigrpSiaQueriesSent"] = ceigrptraffstatsentry.Ceigrpsiaqueriessent
    leafs["cEigrpSiaQueriesRcvd"] = ceigrptraffstatsentry.Ceigrpsiaqueriesrcvd
    leafs["cEigrpAsRouterIdType"] = ceigrptraffstatsentry.Ceigrpasrouteridtype
    leafs["cEigrpAsRouterId"] = ceigrptraffstatsentry.Ceigrpasrouterid
    leafs["cEigrpTopoRoutes"] = ceigrptraffstatsentry.Ceigrptoporoutes
    leafs["cEigrpHeadSerial"] = ceigrptraffstatsentry.Ceigrpheadserial
    leafs["cEigrpNextSerial"] = ceigrptraffstatsentry.Ceigrpnextserial
    leafs["cEigrpXmitPendReplies"] = ceigrptraffstatsentry.Ceigrpxmitpendreplies
    leafs["cEigrpXmitDummies"] = ceigrptraffstatsentry.Ceigrpxmitdummies
    return leafs
}

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetYangName() string { return "cEigrpTraffStatsEntry" }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) SetParent(parent types.Entity) { ceigrptraffstatsentry.parent = parent }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetParent() types.Entity { return ceigrptraffstatsentry.parent }

func (ceigrptraffstatsentry *CISCOEIGRPMIB_Ceigrptraffstatstable_Ceigrptraffstatsentry) GetParentYangName() string { return "cEigrpTraffStatsTable" }

// CISCOEIGRPMIB_Ceigrptopotable
// The table of EIGRP routes and their associated
// attributes for an Autonomous System (AS) configured
// in a VPN is called a topology table.  All route entries in
// the topology table will be indexed by IP network type,
// IP network number and network mask (prefix) size.
type CISCOEIGRPMIB_Ceigrptopotable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The entry for a single EIGRP topology table in the given AS. The type is
    // slice of CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry.
    Ceigrptopoentry []CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry
}

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetFilter() yfilter.YFilter { return ceigrptopotable.YFilter }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) SetFilter(yf yfilter.YFilter) { ceigrptopotable.YFilter = yf }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetGoName(yname string) string {
    if yname == "cEigrpTopoEntry" { return "Ceigrptopoentry" }
    return ""
}

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetSegmentPath() string {
    return "cEigrpTopoTable"
}

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cEigrpTopoEntry" {
        for _, c := range ceigrptopotable.Ceigrptopoentry {
            if ceigrptopotable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry{}
        ceigrptopotable.Ceigrptopoentry = append(ceigrptopotable.Ceigrptopoentry, child)
        return &ceigrptopotable.Ceigrptopoentry[len(ceigrptopotable.Ceigrptopoentry)-1]
    }
    return nil
}

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceigrptopotable.Ceigrptopoentry {
        children[ceigrptopotable.Ceigrptopoentry[i].GetSegmentPath()] = &ceigrptopotable.Ceigrptopoentry[i]
    }
    return children
}

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetYangName() string { return "cEigrpTopoTable" }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) SetParent(parent types.Entity) { ceigrptopotable.parent = parent }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetParent() types.Entity { return ceigrptopotable.parent }

func (ceigrptopotable *CISCOEIGRPMIB_Ceigrptopotable) GetParentYangName() string { return "CISCO-EIGRP-MIB" }

// CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry
// The entry for a single EIGRP topology table in the given
// AS.
type CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry struct {
    parent types.Entity
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

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetFilter() yfilter.YFilter { return ceigrptopoentry.YFilter }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) SetFilter(yf yfilter.YFilter) { ceigrptopoentry.YFilter = yf }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetGoName(yname string) string {
    if yname == "cEigrpVpnId" { return "Ceigrpvpnid" }
    if yname == "cEigrpAsNumber" { return "Ceigrpasnumber" }
    if yname == "cEigrpDestNetType" { return "Ceigrpdestnettype" }
    if yname == "cEigrpDestNet" { return "Ceigrpdestnet" }
    if yname == "cEigrpDestNetPrefixLen" { return "Ceigrpdestnetprefixlen" }
    if yname == "cEigrpActive" { return "Ceigrpactive" }
    if yname == "cEigrpStuckInActive" { return "Ceigrpstuckinactive" }
    if yname == "cEigrpDestSuccessors" { return "Ceigrpdestsuccessors" }
    if yname == "cEigrpFdistance" { return "Ceigrpfdistance" }
    if yname == "cEigrpRouteOriginType" { return "Ceigrprouteorigintype" }
    if yname == "cEigrpRouteOriginAddrType" { return "Ceigrprouteoriginaddrtype" }
    if yname == "cEigrpRouteOriginAddr" { return "Ceigrprouteoriginaddr" }
    if yname == "cEigrpNextHopAddressType" { return "Ceigrpnexthopaddresstype" }
    if yname == "cEigrpNextHopAddress" { return "Ceigrpnexthopaddress" }
    if yname == "cEigrpNextHopInterface" { return "Ceigrpnexthopinterface" }
    if yname == "cEigrpDistance" { return "Ceigrpdistance" }
    if yname == "cEigrpReportDistance" { return "Ceigrpreportdistance" }
    return ""
}

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetSegmentPath() string {
    return "cEigrpTopoEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpasnumber) + "']" + "[cEigrpDestNetType='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpdestnettype) + "']" + "[cEigrpDestNet='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpdestnet) + "']" + "[cEigrpDestNetPrefixLen='" + fmt.Sprintf("%v", ceigrptopoentry.Ceigrpdestnetprefixlen) + "']"
}

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cEigrpVpnId"] = ceigrptopoentry.Ceigrpvpnid
    leafs["cEigrpAsNumber"] = ceigrptopoentry.Ceigrpasnumber
    leafs["cEigrpDestNetType"] = ceigrptopoentry.Ceigrpdestnettype
    leafs["cEigrpDestNet"] = ceigrptopoentry.Ceigrpdestnet
    leafs["cEigrpDestNetPrefixLen"] = ceigrptopoentry.Ceigrpdestnetprefixlen
    leafs["cEigrpActive"] = ceigrptopoentry.Ceigrpactive
    leafs["cEigrpStuckInActive"] = ceigrptopoentry.Ceigrpstuckinactive
    leafs["cEigrpDestSuccessors"] = ceigrptopoentry.Ceigrpdestsuccessors
    leafs["cEigrpFdistance"] = ceigrptopoentry.Ceigrpfdistance
    leafs["cEigrpRouteOriginType"] = ceigrptopoentry.Ceigrprouteorigintype
    leafs["cEigrpRouteOriginAddrType"] = ceigrptopoentry.Ceigrprouteoriginaddrtype
    leafs["cEigrpRouteOriginAddr"] = ceigrptopoentry.Ceigrprouteoriginaddr
    leafs["cEigrpNextHopAddressType"] = ceigrptopoentry.Ceigrpnexthopaddresstype
    leafs["cEigrpNextHopAddress"] = ceigrptopoentry.Ceigrpnexthopaddress
    leafs["cEigrpNextHopInterface"] = ceigrptopoentry.Ceigrpnexthopinterface
    leafs["cEigrpDistance"] = ceigrptopoentry.Ceigrpdistance
    leafs["cEigrpReportDistance"] = ceigrptopoentry.Ceigrpreportdistance
    return leafs
}

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetYangName() string { return "cEigrpTopoEntry" }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) SetParent(parent types.Entity) { ceigrptopoentry.parent = parent }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetParent() types.Entity { return ceigrptopoentry.parent }

func (ceigrptopoentry *CISCOEIGRPMIB_Ceigrptopotable_Ceigrptopoentry) GetParentYangName() string { return "cEigrpTopoTable" }

// CISCOEIGRPMIB_Ceigrppeertable
// The table of established EIGRP peers (neighbors) in the
// selected autonomous system.   Peers are indexed by their
// unique internal handle id, as well as the AS number and
// VPN id.   The peer entry is removed from the table if
// the peer is declared down.
type CISCOEIGRPMIB_Ceigrppeertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics and operational parameters for a single peer in the AS. The type
    // is slice of CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry.
    Ceigrppeerentry []CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry
}

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetFilter() yfilter.YFilter { return ceigrppeertable.YFilter }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) SetFilter(yf yfilter.YFilter) { ceigrppeertable.YFilter = yf }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetGoName(yname string) string {
    if yname == "cEigrpPeerEntry" { return "Ceigrppeerentry" }
    return ""
}

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetSegmentPath() string {
    return "cEigrpPeerTable"
}

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cEigrpPeerEntry" {
        for _, c := range ceigrppeertable.Ceigrppeerentry {
            if ceigrppeertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry{}
        ceigrppeertable.Ceigrppeerentry = append(ceigrppeertable.Ceigrppeerentry, child)
        return &ceigrppeertable.Ceigrppeerentry[len(ceigrppeertable.Ceigrppeerentry)-1]
    }
    return nil
}

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceigrppeertable.Ceigrppeerentry {
        children[ceigrppeertable.Ceigrppeerentry[i].GetSegmentPath()] = &ceigrppeertable.Ceigrppeerentry[i]
    }
    return children
}

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetYangName() string { return "cEigrpPeerTable" }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) SetParent(parent types.Entity) { ceigrppeertable.parent = parent }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetParent() types.Entity { return ceigrppeertable.parent }

func (ceigrppeertable *CISCOEIGRPMIB_Ceigrppeertable) GetParentYangName() string { return "CISCO-EIGRP-MIB" }

// CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry
// Statistics and operational parameters for a single peer
// in the AS.
type CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry struct {
    parent types.Entity
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

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetFilter() yfilter.YFilter { return ceigrppeerentry.YFilter }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) SetFilter(yf yfilter.YFilter) { ceigrppeerentry.YFilter = yf }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetGoName(yname string) string {
    if yname == "cEigrpVpnId" { return "Ceigrpvpnid" }
    if yname == "cEigrpAsNumber" { return "Ceigrpasnumber" }
    if yname == "cEigrpHandle" { return "Ceigrphandle" }
    if yname == "cEigrpPeerAddrType" { return "Ceigrppeeraddrtype" }
    if yname == "cEigrpPeerAddr" { return "Ceigrppeeraddr" }
    if yname == "cEigrpPeerIfIndex" { return "Ceigrppeerifindex" }
    if yname == "cEigrpHoldTime" { return "Ceigrpholdtime" }
    if yname == "cEigrpUpTime" { return "Ceigrpuptime" }
    if yname == "cEigrpSrtt" { return "Ceigrpsrtt" }
    if yname == "cEigrpRto" { return "Ceigrprto" }
    if yname == "cEigrpPktsEnqueued" { return "Ceigrppktsenqueued" }
    if yname == "cEigrpLastSeq" { return "Ceigrplastseq" }
    if yname == "cEigrpVersion" { return "Ceigrpversion" }
    if yname == "cEigrpRetrans" { return "Ceigrpretrans" }
    if yname == "cEigrpRetries" { return "Ceigrpretries" }
    return ""
}

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetSegmentPath() string {
    return "cEigrpPeerEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrppeerentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrppeerentry.Ceigrpasnumber) + "']" + "[cEigrpHandle='" + fmt.Sprintf("%v", ceigrppeerentry.Ceigrphandle) + "']"
}

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cEigrpVpnId"] = ceigrppeerentry.Ceigrpvpnid
    leafs["cEigrpAsNumber"] = ceigrppeerentry.Ceigrpasnumber
    leafs["cEigrpHandle"] = ceigrppeerentry.Ceigrphandle
    leafs["cEigrpPeerAddrType"] = ceigrppeerentry.Ceigrppeeraddrtype
    leafs["cEigrpPeerAddr"] = ceigrppeerentry.Ceigrppeeraddr
    leafs["cEigrpPeerIfIndex"] = ceigrppeerentry.Ceigrppeerifindex
    leafs["cEigrpHoldTime"] = ceigrppeerentry.Ceigrpholdtime
    leafs["cEigrpUpTime"] = ceigrppeerentry.Ceigrpuptime
    leafs["cEigrpSrtt"] = ceigrppeerentry.Ceigrpsrtt
    leafs["cEigrpRto"] = ceigrppeerentry.Ceigrprto
    leafs["cEigrpPktsEnqueued"] = ceigrppeerentry.Ceigrppktsenqueued
    leafs["cEigrpLastSeq"] = ceigrppeerentry.Ceigrplastseq
    leafs["cEigrpVersion"] = ceigrppeerentry.Ceigrpversion
    leafs["cEigrpRetrans"] = ceigrppeerentry.Ceigrpretrans
    leafs["cEigrpRetries"] = ceigrppeerentry.Ceigrpretries
    return leafs
}

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetYangName() string { return "cEigrpPeerEntry" }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) SetParent(parent types.Entity) { ceigrppeerentry.parent = parent }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetParent() types.Entity { return ceigrppeerentry.parent }

func (ceigrppeerentry *CISCOEIGRPMIB_Ceigrppeertable_Ceigrppeerentry) GetParentYangName() string { return "cEigrpPeerTable" }

// CISCOEIGRPMIB_Ceigrpinterfacetable
// The table of interfaces over which EIGRP is running, and
// their associated statistics.   This table is independent
// of whether any peer adjacencies have been formed over
// the interfaces or not.   Interfaces running EIGRP are
// determined by whether their assigned IP addresses fall
// within configured EIGRP network statements.
type CISCOEIGRPMIB_Ceigrpinterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information for a single interface running EIGRP in the AS and VPN. The
    // type is slice of CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry.
    Ceigrpinterfaceentry []CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry
}

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetFilter() yfilter.YFilter { return ceigrpinterfacetable.YFilter }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) SetFilter(yf yfilter.YFilter) { ceigrpinterfacetable.YFilter = yf }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetGoName(yname string) string {
    if yname == "cEigrpInterfaceEntry" { return "Ceigrpinterfaceentry" }
    return ""
}

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetSegmentPath() string {
    return "cEigrpInterfaceTable"
}

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cEigrpInterfaceEntry" {
        for _, c := range ceigrpinterfacetable.Ceigrpinterfaceentry {
            if ceigrpinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry{}
        ceigrpinterfacetable.Ceigrpinterfaceentry = append(ceigrpinterfacetable.Ceigrpinterfaceentry, child)
        return &ceigrpinterfacetable.Ceigrpinterfaceentry[len(ceigrpinterfacetable.Ceigrpinterfaceentry)-1]
    }
    return nil
}

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceigrpinterfacetable.Ceigrpinterfaceentry {
        children[ceigrpinterfacetable.Ceigrpinterfaceentry[i].GetSegmentPath()] = &ceigrpinterfacetable.Ceigrpinterfaceentry[i]
    }
    return children
}

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetYangName() string { return "cEigrpInterfaceTable" }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) SetParent(parent types.Entity) { ceigrpinterfacetable.parent = parent }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetParent() types.Entity { return ceigrpinterfacetable.parent }

func (ceigrpinterfacetable *CISCOEIGRPMIB_Ceigrpinterfacetable) GetParentYangName() string { return "CISCO-EIGRP-MIB" }

// CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry
// Information for a single interface running EIGRP in the
// AS and VPN.
type CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry struct {
    parent types.Entity
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

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetFilter() yfilter.YFilter { return ceigrpinterfaceentry.YFilter }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) SetFilter(yf yfilter.YFilter) { ceigrpinterfaceentry.YFilter = yf }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetGoName(yname string) string {
    if yname == "cEigrpVpnId" { return "Ceigrpvpnid" }
    if yname == "cEigrpAsNumber" { return "Ceigrpasnumber" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cEigrpPeerCount" { return "Ceigrppeercount" }
    if yname == "cEigrpXmitReliableQ" { return "Ceigrpxmitreliableq" }
    if yname == "cEigrpXmitUnreliableQ" { return "Ceigrpxmitunreliableq" }
    if yname == "cEigrpMeanSrtt" { return "Ceigrpmeansrtt" }
    if yname == "cEigrpPacingReliable" { return "Ceigrppacingreliable" }
    if yname == "cEigrpPacingUnreliable" { return "Ceigrppacingunreliable" }
    if yname == "cEigrpMFlowTimer" { return "Ceigrpmflowtimer" }
    if yname == "cEigrpPendingRoutes" { return "Ceigrppendingroutes" }
    if yname == "cEigrpHelloInterval" { return "Ceigrphellointerval" }
    if yname == "cEigrpXmitNextSerial" { return "Ceigrpxmitnextserial" }
    if yname == "cEigrpUMcasts" { return "Ceigrpumcasts" }
    if yname == "cEigrpRMcasts" { return "Ceigrprmcasts" }
    if yname == "cEigrpUUcasts" { return "Ceigrpuucasts" }
    if yname == "cEigrpRUcasts" { return "Ceigrprucasts" }
    if yname == "cEigrpMcastExcepts" { return "Ceigrpmcastexcepts" }
    if yname == "cEigrpCRpkts" { return "Ceigrpcrpkts" }
    if yname == "cEigrpAcksSuppressed" { return "Ceigrpackssuppressed" }
    if yname == "cEigrpRetransSent" { return "Ceigrpretranssent" }
    if yname == "cEigrpOOSrvcd" { return "Ceigrpoosrvcd" }
    if yname == "cEigrpAuthMode" { return "Ceigrpauthmode" }
    if yname == "cEigrpAuthKeyChain" { return "Ceigrpauthkeychain" }
    return ""
}

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetSegmentPath() string {
    return "cEigrpInterfaceEntry" + "[cEigrpVpnId='" + fmt.Sprintf("%v", ceigrpinterfaceentry.Ceigrpvpnid) + "']" + "[cEigrpAsNumber='" + fmt.Sprintf("%v", ceigrpinterfaceentry.Ceigrpasnumber) + "']" + "[ifIndex='" + fmt.Sprintf("%v", ceigrpinterfaceentry.Ifindex) + "']"
}

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cEigrpVpnId"] = ceigrpinterfaceentry.Ceigrpvpnid
    leafs["cEigrpAsNumber"] = ceigrpinterfaceentry.Ceigrpasnumber
    leafs["ifIndex"] = ceigrpinterfaceentry.Ifindex
    leafs["cEigrpPeerCount"] = ceigrpinterfaceentry.Ceigrppeercount
    leafs["cEigrpXmitReliableQ"] = ceigrpinterfaceentry.Ceigrpxmitreliableq
    leafs["cEigrpXmitUnreliableQ"] = ceigrpinterfaceentry.Ceigrpxmitunreliableq
    leafs["cEigrpMeanSrtt"] = ceigrpinterfaceentry.Ceigrpmeansrtt
    leafs["cEigrpPacingReliable"] = ceigrpinterfaceentry.Ceigrppacingreliable
    leafs["cEigrpPacingUnreliable"] = ceigrpinterfaceentry.Ceigrppacingunreliable
    leafs["cEigrpMFlowTimer"] = ceigrpinterfaceentry.Ceigrpmflowtimer
    leafs["cEigrpPendingRoutes"] = ceigrpinterfaceentry.Ceigrppendingroutes
    leafs["cEigrpHelloInterval"] = ceigrpinterfaceentry.Ceigrphellointerval
    leafs["cEigrpXmitNextSerial"] = ceigrpinterfaceentry.Ceigrpxmitnextserial
    leafs["cEigrpUMcasts"] = ceigrpinterfaceentry.Ceigrpumcasts
    leafs["cEigrpRMcasts"] = ceigrpinterfaceentry.Ceigrprmcasts
    leafs["cEigrpUUcasts"] = ceigrpinterfaceentry.Ceigrpuucasts
    leafs["cEigrpRUcasts"] = ceigrpinterfaceentry.Ceigrprucasts
    leafs["cEigrpMcastExcepts"] = ceigrpinterfaceentry.Ceigrpmcastexcepts
    leafs["cEigrpCRpkts"] = ceigrpinterfaceentry.Ceigrpcrpkts
    leafs["cEigrpAcksSuppressed"] = ceigrpinterfaceentry.Ceigrpackssuppressed
    leafs["cEigrpRetransSent"] = ceigrpinterfaceentry.Ceigrpretranssent
    leafs["cEigrpOOSrvcd"] = ceigrpinterfaceentry.Ceigrpoosrvcd
    leafs["cEigrpAuthMode"] = ceigrpinterfaceentry.Ceigrpauthmode
    leafs["cEigrpAuthKeyChain"] = ceigrpinterfaceentry.Ceigrpauthkeychain
    return leafs
}

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetYangName() string { return "cEigrpInterfaceEntry" }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) SetParent(parent types.Entity) { ceigrpinterfaceentry.parent = parent }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetParent() types.Entity { return ceigrpinterfaceentry.parent }

func (ceigrpinterfaceentry *CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry) GetParentYangName() string { return "cEigrpInterfaceTable" }

// CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode represents md5   :  MD5 authentication enabled on the interface
type CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode string

const (
    CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode_none CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode = "none"

    CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode_md5 CISCOEIGRPMIB_Ceigrpinterfacetable_Ceigrpinterfaceentry_Ceigrpauthmode = "md5"
)

