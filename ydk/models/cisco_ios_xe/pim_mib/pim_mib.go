// The MIB module for management of PIM routers.
package pim_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pim_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:PIM-MIB PIM-MIB}", reflect.TypeOf(PIMMIB{}))
    ydk.RegisterEntity("PIM-MIB:PIM-MIB", reflect.TypeOf(PIMMIB{}))
}

// PIMMIB
type PIMMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Pim PIMMIB_Pim

    // The (conceptual) table listing the router's PIM interfaces. IGMP and PIM
    // are enabled on all interfaces listed in this table.
    Piminterfacetable PIMMIB_Piminterfacetable

    // The (conceptual) table listing the router's PIM neighbors.
    Pimneighbortable PIMMIB_Pimneighbortable

    // The (conceptual) table listing PIM-specific information on a subset of the
    // rows of the ipMRouteTable defined in the IP Multicast MIB.
    Pimipmroutetable PIMMIB_Pimipmroutetable

    // The (conceptual) table listing PIM version 1 information for the Rendezvous
    // Points (RPs) for IP multicast groups. This table is deprecated since its
    // function is replaced by the pimRPSetTable for PIM version 2.
    Pimrptable PIMMIB_Pimrptable

    // The (conceptual) table listing PIM information for candidate Rendezvous
    // Points (RPs) for IP multicast groups. When the local router is the BSR,
    // this information is obtained from received Candidate-RP-Advertisements. 
    // When the local router is not the BSR, this information is obtained from
    // received RP-Set messages.
    Pimrpsettable PIMMIB_Pimrpsettable

    // The (conceptual) table listing PIM-specific information on a subset of the
    // rows of the ipMRouteNextHopTable defined in the IP Multicast MIB.
    Pimipmroutenexthoptable PIMMIB_Pimipmroutenexthoptable

    // The (conceptual) table listing the IP multicast groups for which the local
    // router is to advertise itself as a Candidate-RP when the value of
    // pimComponentCRPHoldTime is non-zero.  If this table is empty, then the
    // local router      will advertise itself as a Candidate-RP for all groups
    // (providing the value of pimComponentCRPHoldTime is non- zero).
    Pimcandidaterptable PIMMIB_Pimcandidaterptable

    // The (conceptual) table containing objects specific to a PIM domain.  One
    // row exists for each domain to which the router is connected.  A PIM-SM
    // domain is defined as an area of the network over which Bootstrap messages
    // are forwarded. Typically, a PIM-SM router will be a member of exactly one
    // domain.  This table also supports, however, routers which may form a border
    // between two PIM-SM domains and do not forward Bootstrap messages between
    // them.
    Pimcomponenttable PIMMIB_Pimcomponenttable
}

func (pIMMIB *PIMMIB) GetFilter() yfilter.YFilter { return pIMMIB.YFilter }

func (pIMMIB *PIMMIB) SetFilter(yf yfilter.YFilter) { pIMMIB.YFilter = yf }

func (pIMMIB *PIMMIB) GetGoName(yname string) string {
    if yname == "pim" { return "Pim" }
    if yname == "pimInterfaceTable" { return "Piminterfacetable" }
    if yname == "pimNeighborTable" { return "Pimneighbortable" }
    if yname == "pimIpMRouteTable" { return "Pimipmroutetable" }
    if yname == "pimRPTable" { return "Pimrptable" }
    if yname == "pimRPSetTable" { return "Pimrpsettable" }
    if yname == "pimIpMRouteNextHopTable" { return "Pimipmroutenexthoptable" }
    if yname == "pimCandidateRPTable" { return "Pimcandidaterptable" }
    if yname == "pimComponentTable" { return "Pimcomponenttable" }
    return ""
}

func (pIMMIB *PIMMIB) GetSegmentPath() string {
    return "PIM-MIB:PIM-MIB"
}

func (pIMMIB *PIMMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pim" {
        return &pIMMIB.Pim
    }
    if childYangName == "pimInterfaceTable" {
        return &pIMMIB.Piminterfacetable
    }
    if childYangName == "pimNeighborTable" {
        return &pIMMIB.Pimneighbortable
    }
    if childYangName == "pimIpMRouteTable" {
        return &pIMMIB.Pimipmroutetable
    }
    if childYangName == "pimRPTable" {
        return &pIMMIB.Pimrptable
    }
    if childYangName == "pimRPSetTable" {
        return &pIMMIB.Pimrpsettable
    }
    if childYangName == "pimIpMRouteNextHopTable" {
        return &pIMMIB.Pimipmroutenexthoptable
    }
    if childYangName == "pimCandidateRPTable" {
        return &pIMMIB.Pimcandidaterptable
    }
    if childYangName == "pimComponentTable" {
        return &pIMMIB.Pimcomponenttable
    }
    return nil
}

func (pIMMIB *PIMMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pim"] = &pIMMIB.Pim
    children["pimInterfaceTable"] = &pIMMIB.Piminterfacetable
    children["pimNeighborTable"] = &pIMMIB.Pimneighbortable
    children["pimIpMRouteTable"] = &pIMMIB.Pimipmroutetable
    children["pimRPTable"] = &pIMMIB.Pimrptable
    children["pimRPSetTable"] = &pIMMIB.Pimrpsettable
    children["pimIpMRouteNextHopTable"] = &pIMMIB.Pimipmroutenexthoptable
    children["pimCandidateRPTable"] = &pIMMIB.Pimcandidaterptable
    children["pimComponentTable"] = &pIMMIB.Pimcomponenttable
    return children
}

func (pIMMIB *PIMMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pIMMIB *PIMMIB) GetBundleName() string { return "cisco_ios_xe" }

func (pIMMIB *PIMMIB) GetYangName() string { return "PIM-MIB" }

func (pIMMIB *PIMMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pIMMIB *PIMMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pIMMIB *PIMMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pIMMIB *PIMMIB) SetParent(parent types.Entity) { pIMMIB.parent = parent }

func (pIMMIB *PIMMIB) GetParent() types.Entity { return pIMMIB.parent }

func (pIMMIB *PIMMIB) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pim
type PIMMIB_Pim struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The default interval at which periodic PIM-SM Join/Prune messages are to be
    // sent. The type is interface{} with range: -2147483648..2147483647. Units
    // are seconds.
    Pimjoinpruneinterval interface{}
}

func (pim *PIMMIB_Pim) GetFilter() yfilter.YFilter { return pim.YFilter }

func (pim *PIMMIB_Pim) SetFilter(yf yfilter.YFilter) { pim.YFilter = yf }

func (pim *PIMMIB_Pim) GetGoName(yname string) string {
    if yname == "pimJoinPruneInterval" { return "Pimjoinpruneinterval" }
    return ""
}

func (pim *PIMMIB_Pim) GetSegmentPath() string {
    return "pim"
}

func (pim *PIMMIB_Pim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pim *PIMMIB_Pim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pim *PIMMIB_Pim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pimJoinPruneInterval"] = pim.Pimjoinpruneinterval
    return leafs
}

func (pim *PIMMIB_Pim) GetBundleName() string { return "cisco_ios_xe" }

func (pim *PIMMIB_Pim) GetYangName() string { return "pim" }

func (pim *PIMMIB_Pim) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pim *PIMMIB_Pim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pim *PIMMIB_Pim) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pim *PIMMIB_Pim) SetParent(parent types.Entity) { pim.parent = parent }

func (pim *PIMMIB_Pim) GetParent() types.Entity { return pim.parent }

func (pim *PIMMIB_Pim) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Piminterfacetable
// The (conceptual) table listing the router's PIM interfaces.
// IGMP and PIM are enabled on all interfaces listed in this
// table.
type PIMMIB_Piminterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimInterfaceTable. The type is slice of
    // PIMMIB_Piminterfacetable_Piminterfaceentry.
    Piminterfaceentry []PIMMIB_Piminterfacetable_Piminterfaceentry
}

func (piminterfacetable *PIMMIB_Piminterfacetable) GetFilter() yfilter.YFilter { return piminterfacetable.YFilter }

func (piminterfacetable *PIMMIB_Piminterfacetable) SetFilter(yf yfilter.YFilter) { piminterfacetable.YFilter = yf }

func (piminterfacetable *PIMMIB_Piminterfacetable) GetGoName(yname string) string {
    if yname == "pimInterfaceEntry" { return "Piminterfaceentry" }
    return ""
}

func (piminterfacetable *PIMMIB_Piminterfacetable) GetSegmentPath() string {
    return "pimInterfaceTable"
}

func (piminterfacetable *PIMMIB_Piminterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimInterfaceEntry" {
        for _, c := range piminterfacetable.Piminterfaceentry {
            if piminterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Piminterfacetable_Piminterfaceentry{}
        piminterfacetable.Piminterfaceentry = append(piminterfacetable.Piminterfaceentry, child)
        return &piminterfacetable.Piminterfaceentry[len(piminterfacetable.Piminterfaceentry)-1]
    }
    return nil
}

func (piminterfacetable *PIMMIB_Piminterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range piminterfacetable.Piminterfaceentry {
        children[piminterfacetable.Piminterfaceentry[i].GetSegmentPath()] = &piminterfacetable.Piminterfaceentry[i]
    }
    return children
}

func (piminterfacetable *PIMMIB_Piminterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (piminterfacetable *PIMMIB_Piminterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (piminterfacetable *PIMMIB_Piminterfacetable) GetYangName() string { return "pimInterfaceTable" }

func (piminterfacetable *PIMMIB_Piminterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (piminterfacetable *PIMMIB_Piminterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (piminterfacetable *PIMMIB_Piminterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (piminterfacetable *PIMMIB_Piminterfacetable) SetParent(parent types.Entity) { piminterfacetable.parent = parent }

func (piminterfacetable *PIMMIB_Piminterfacetable) GetParent() types.Entity { return piminterfacetable.parent }

func (piminterfacetable *PIMMIB_Piminterfacetable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Piminterfacetable_Piminterfaceentry
// An entry (conceptual row) in the pimInterfaceTable.
type PIMMIB_Piminterfacetable_Piminterfaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of this PIM interface. The type
    // is interface{} with range: 1..2147483647.
    Piminterfaceifindex interface{}

    // The IP address of the PIM interface. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Piminterfaceaddress interface{}

    // The network mask for the IP address of the PIM interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Piminterfacenetmask interface{}

    // The configured mode of this PIM interface.  A value of sparseDense is only
    // valid for PIMv1. The type is Piminterfacemode.
    Piminterfacemode interface{}

    // The Designated Router on this PIM interface.  For point-to- point
    // interfaces, this object has the value 0.0.0.0. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Piminterfacedr interface{}

    // The frequency at which PIM Hello messages are transmitted on this
    // interface. The type is interface{} with range: -2147483648..2147483647.
    // Units are seconds.
    Piminterfacehellointerval interface{}

    // The status of this entry.  Creating the entry enables PIM on the interface;
    // destroying the entry disables PIM on the interface. The type is RowStatus.
    Piminterfacestatus interface{}

    // The frequency at which PIM Join/Prune messages are transmitted on this PIM
    // interface.  The default value of this object is the pimJoinPruneInterval.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // seconds.
    Piminterfacejoinpruneinterval interface{}

    // The preference value for the local interface as a candidate bootstrap
    // router.  The value of -1 is used to indicate that the local interface is
    // not a candidate BSR interface. The type is interface{} with range: -1..255.
    Piminterfacecbsrpreference interface{}
}

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetFilter() yfilter.YFilter { return piminterfaceentry.YFilter }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) SetFilter(yf yfilter.YFilter) { piminterfaceentry.YFilter = yf }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetGoName(yname string) string {
    if yname == "pimInterfaceIfIndex" { return "Piminterfaceifindex" }
    if yname == "pimInterfaceAddress" { return "Piminterfaceaddress" }
    if yname == "pimInterfaceNetMask" { return "Piminterfacenetmask" }
    if yname == "pimInterfaceMode" { return "Piminterfacemode" }
    if yname == "pimInterfaceDR" { return "Piminterfacedr" }
    if yname == "pimInterfaceHelloInterval" { return "Piminterfacehellointerval" }
    if yname == "pimInterfaceStatus" { return "Piminterfacestatus" }
    if yname == "pimInterfaceJoinPruneInterval" { return "Piminterfacejoinpruneinterval" }
    if yname == "pimInterfaceCBSRPreference" { return "Piminterfacecbsrpreference" }
    return ""
}

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetSegmentPath() string {
    return "pimInterfaceEntry" + "[pimInterfaceIfIndex='" + fmt.Sprintf("%v", piminterfaceentry.Piminterfaceifindex) + "']"
}

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pimInterfaceIfIndex"] = piminterfaceentry.Piminterfaceifindex
    leafs["pimInterfaceAddress"] = piminterfaceentry.Piminterfaceaddress
    leafs["pimInterfaceNetMask"] = piminterfaceentry.Piminterfacenetmask
    leafs["pimInterfaceMode"] = piminterfaceentry.Piminterfacemode
    leafs["pimInterfaceDR"] = piminterfaceentry.Piminterfacedr
    leafs["pimInterfaceHelloInterval"] = piminterfaceentry.Piminterfacehellointerval
    leafs["pimInterfaceStatus"] = piminterfaceentry.Piminterfacestatus
    leafs["pimInterfaceJoinPruneInterval"] = piminterfaceentry.Piminterfacejoinpruneinterval
    leafs["pimInterfaceCBSRPreference"] = piminterfaceentry.Piminterfacecbsrpreference
    return leafs
}

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetYangName() string { return "pimInterfaceEntry" }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) SetParent(parent types.Entity) { piminterfaceentry.parent = parent }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetParent() types.Entity { return piminterfaceentry.parent }

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetParentYangName() string { return "pimInterfaceTable" }

// PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode represents sparseDense is only valid for PIMv1.
type PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode string

const (
    PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode_dense PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode = "dense"

    PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode_sparse PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode = "sparse"

    PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode_sparseDense PIMMIB_Piminterfacetable_Piminterfaceentry_Piminterfacemode = "sparseDense"
)

// PIMMIB_Pimneighbortable
// The (conceptual) table listing the router's PIM neighbors.
type PIMMIB_Pimneighbortable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimNeighborTable. The type is slice of
    // PIMMIB_Pimneighbortable_Pimneighborentry.
    Pimneighborentry []PIMMIB_Pimneighbortable_Pimneighborentry
}

func (pimneighbortable *PIMMIB_Pimneighbortable) GetFilter() yfilter.YFilter { return pimneighbortable.YFilter }

func (pimneighbortable *PIMMIB_Pimneighbortable) SetFilter(yf yfilter.YFilter) { pimneighbortable.YFilter = yf }

func (pimneighbortable *PIMMIB_Pimneighbortable) GetGoName(yname string) string {
    if yname == "pimNeighborEntry" { return "Pimneighborentry" }
    return ""
}

func (pimneighbortable *PIMMIB_Pimneighbortable) GetSegmentPath() string {
    return "pimNeighborTable"
}

func (pimneighbortable *PIMMIB_Pimneighbortable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimNeighborEntry" {
        for _, c := range pimneighbortable.Pimneighborentry {
            if pimneighbortable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Pimneighbortable_Pimneighborentry{}
        pimneighbortable.Pimneighborentry = append(pimneighbortable.Pimneighborentry, child)
        return &pimneighbortable.Pimneighborentry[len(pimneighbortable.Pimneighborentry)-1]
    }
    return nil
}

func (pimneighbortable *PIMMIB_Pimneighbortable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pimneighbortable.Pimneighborentry {
        children[pimneighbortable.Pimneighborentry[i].GetSegmentPath()] = &pimneighbortable.Pimneighborentry[i]
    }
    return children
}

func (pimneighbortable *PIMMIB_Pimneighbortable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pimneighbortable *PIMMIB_Pimneighbortable) GetBundleName() string { return "cisco_ios_xe" }

func (pimneighbortable *PIMMIB_Pimneighbortable) GetYangName() string { return "pimNeighborTable" }

func (pimneighbortable *PIMMIB_Pimneighbortable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimneighbortable *PIMMIB_Pimneighbortable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimneighbortable *PIMMIB_Pimneighbortable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimneighbortable *PIMMIB_Pimneighbortable) SetParent(parent types.Entity) { pimneighbortable.parent = parent }

func (pimneighbortable *PIMMIB_Pimneighbortable) GetParent() types.Entity { return pimneighbortable.parent }

func (pimneighbortable *PIMMIB_Pimneighbortable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pimneighbortable_Pimneighborentry
// An entry (conceptual row) in the pimNeighborTable.
type PIMMIB_Pimneighbortable_Pimneighborentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the PIM neighbor for which this
    // entry contains information. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimneighboraddress interface{}

    // The value of ifIndex for the interface used to reach this PIM neighbor. The
    // type is interface{} with range: 1..2147483647.
    Pimneighborifindex interface{}

    // The time since this PIM neighbor (last) became a neighbor of the local
    // router. The type is interface{} with range: 0..4294967295.
    Pimneighboruptime interface{}

    // The minimum time remaining before this PIM neighbor will be aged out. The
    // type is interface{} with range: 0..4294967295.
    Pimneighborexpirytime interface{}

    // The active PIM mode of this neighbor.  This object is deprecated for PIMv2
    // routers since all neighbors on the interface must be either dense or sparse
    // as determined by the protocol running on the interface. The type is
    // Pimneighbormode.
    Pimneighbormode interface{}
}

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetFilter() yfilter.YFilter { return pimneighborentry.YFilter }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) SetFilter(yf yfilter.YFilter) { pimneighborentry.YFilter = yf }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetGoName(yname string) string {
    if yname == "pimNeighborAddress" { return "Pimneighboraddress" }
    if yname == "pimNeighborIfIndex" { return "Pimneighborifindex" }
    if yname == "pimNeighborUpTime" { return "Pimneighboruptime" }
    if yname == "pimNeighborExpiryTime" { return "Pimneighborexpirytime" }
    if yname == "pimNeighborMode" { return "Pimneighbormode" }
    return ""
}

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetSegmentPath() string {
    return "pimNeighborEntry" + "[pimNeighborAddress='" + fmt.Sprintf("%v", pimneighborentry.Pimneighboraddress) + "']"
}

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pimNeighborAddress"] = pimneighborentry.Pimneighboraddress
    leafs["pimNeighborIfIndex"] = pimneighborentry.Pimneighborifindex
    leafs["pimNeighborUpTime"] = pimneighborentry.Pimneighboruptime
    leafs["pimNeighborExpiryTime"] = pimneighborentry.Pimneighborexpirytime
    leafs["pimNeighborMode"] = pimneighborentry.Pimneighbormode
    return leafs
}

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetBundleName() string { return "cisco_ios_xe" }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetYangName() string { return "pimNeighborEntry" }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) SetParent(parent types.Entity) { pimneighborentry.parent = parent }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetParent() types.Entity { return pimneighborentry.parent }

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetParentYangName() string { return "pimNeighborTable" }

// PIMMIB_Pimneighbortable_Pimneighborentry_Pimneighbormode represents the protocol running on the interface.
type PIMMIB_Pimneighbortable_Pimneighborentry_Pimneighbormode string

const (
    PIMMIB_Pimneighbortable_Pimneighborentry_Pimneighbormode_dense PIMMIB_Pimneighbortable_Pimneighborentry_Pimneighbormode = "dense"

    PIMMIB_Pimneighbortable_Pimneighborentry_Pimneighbormode_sparse PIMMIB_Pimneighbortable_Pimneighborentry_Pimneighbormode = "sparse"
)

// PIMMIB_Pimipmroutetable
// The (conceptual) table listing PIM-specific information on
// a subset of the rows of the ipMRouteTable defined in the IP
// Multicast MIB.
type PIMMIB_Pimipmroutetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimIpMRouteTable.  There is one entry per
    // entry in the ipMRouteTable whose incoming interface is running PIM. The
    // type is slice of PIMMIB_Pimipmroutetable_Pimipmrouteentry.
    Pimipmrouteentry []PIMMIB_Pimipmroutetable_Pimipmrouteentry
}

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetFilter() yfilter.YFilter { return pimipmroutetable.YFilter }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) SetFilter(yf yfilter.YFilter) { pimipmroutetable.YFilter = yf }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetGoName(yname string) string {
    if yname == "pimIpMRouteEntry" { return "Pimipmrouteentry" }
    return ""
}

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetSegmentPath() string {
    return "pimIpMRouteTable"
}

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimIpMRouteEntry" {
        for _, c := range pimipmroutetable.Pimipmrouteentry {
            if pimipmroutetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Pimipmroutetable_Pimipmrouteentry{}
        pimipmroutetable.Pimipmrouteentry = append(pimipmroutetable.Pimipmrouteentry, child)
        return &pimipmroutetable.Pimipmrouteentry[len(pimipmroutetable.Pimipmrouteentry)-1]
    }
    return nil
}

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pimipmroutetable.Pimipmrouteentry {
        children[pimipmroutetable.Pimipmrouteentry[i].GetSegmentPath()] = &pimipmroutetable.Pimipmrouteentry[i]
    }
    return children
}

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetBundleName() string { return "cisco_ios_xe" }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetYangName() string { return "pimIpMRouteTable" }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) SetParent(parent types.Entity) { pimipmroutetable.parent = parent }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetParent() types.Entity { return pimipmroutetable.parent }

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pimipmroutetable_Pimipmrouteentry
// An entry (conceptual row) in the pimIpMRouteTable.  There
// is one entry per entry in the ipMRouteTable whose incoming
// interface is running PIM.
type PIMMIB_Pimipmroutetable_Pimipmrouteentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmroutegroup
    Ipmroutegroup interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmroutesource
    Ipmroutesource interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmroutesourcemask
    Ipmroutesourcemask interface{}

    // The time remaining before the router changes its upstream neighbor back to
    // its RPF neighbor.  This timer is called the Assert timer in the PIM Sparse
    // and Dense mode specification.      A value of 0 indicates that no Assert
    // has changed the upstream neighbor away from the RPF neighbor. The type is
    // interface{} with range: 0..4294967295.
    Pimipmrouteupstreamasserttimer interface{}

    // The metric advertised by the assert winner on the upstream interface, or 0
    // if no such assert is in received. The type is interface{} with range:
    // -2147483648..2147483647.
    Pimipmrouteassertmetric interface{}

    // The preference advertised by the assert winner on the upstream interface,
    // or 0 if no such assert is in effect. The type is interface{} with range:
    // -2147483648..2147483647.
    Pimipmrouteassertmetricpref interface{}

    // The value of the RPT-bit advertised by the assert winner on the upstream
    // interface, or false if no such assert is in effect. The type is bool.
    Pimipmrouteassertrptbit interface{}

    // This object describes PIM-specific flags related to a multicast state
    // entry.  See the PIM Sparse Mode specification for the meaning of the RPT
    // and SPT bits. The type is string with length: 1.
    Pimipmrouteflags interface{}
}

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetFilter() yfilter.YFilter { return pimipmrouteentry.YFilter }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) SetFilter(yf yfilter.YFilter) { pimipmrouteentry.YFilter = yf }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetGoName(yname string) string {
    if yname == "ipMRouteGroup" { return "Ipmroutegroup" }
    if yname == "ipMRouteSource" { return "Ipmroutesource" }
    if yname == "ipMRouteSourceMask" { return "Ipmroutesourcemask" }
    if yname == "pimIpMRouteUpstreamAssertTimer" { return "Pimipmrouteupstreamasserttimer" }
    if yname == "pimIpMRouteAssertMetric" { return "Pimipmrouteassertmetric" }
    if yname == "pimIpMRouteAssertMetricPref" { return "Pimipmrouteassertmetricpref" }
    if yname == "pimIpMRouteAssertRPTBit" { return "Pimipmrouteassertrptbit" }
    if yname == "pimIpMRouteFlags" { return "Pimipmrouteflags" }
    return ""
}

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetSegmentPath() string {
    return "pimIpMRouteEntry" + "[ipMRouteGroup='" + fmt.Sprintf("%v", pimipmrouteentry.Ipmroutegroup) + "']" + "[ipMRouteSource='" + fmt.Sprintf("%v", pimipmrouteentry.Ipmroutesource) + "']" + "[ipMRouteSourceMask='" + fmt.Sprintf("%v", pimipmrouteentry.Ipmroutesourcemask) + "']"
}

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteGroup"] = pimipmrouteentry.Ipmroutegroup
    leafs["ipMRouteSource"] = pimipmrouteentry.Ipmroutesource
    leafs["ipMRouteSourceMask"] = pimipmrouteentry.Ipmroutesourcemask
    leafs["pimIpMRouteUpstreamAssertTimer"] = pimipmrouteentry.Pimipmrouteupstreamasserttimer
    leafs["pimIpMRouteAssertMetric"] = pimipmrouteentry.Pimipmrouteassertmetric
    leafs["pimIpMRouteAssertMetricPref"] = pimipmrouteentry.Pimipmrouteassertmetricpref
    leafs["pimIpMRouteAssertRPTBit"] = pimipmrouteentry.Pimipmrouteassertrptbit
    leafs["pimIpMRouteFlags"] = pimipmrouteentry.Pimipmrouteflags
    return leafs
}

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetBundleName() string { return "cisco_ios_xe" }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetYangName() string { return "pimIpMRouteEntry" }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) SetParent(parent types.Entity) { pimipmrouteentry.parent = parent }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetParent() types.Entity { return pimipmrouteentry.parent }

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetParentYangName() string { return "pimIpMRouteTable" }

// PIMMIB_Pimrptable
// The (conceptual) table listing PIM version 1 information
// for the Rendezvous Points (RPs) for IP multicast groups.
// This table is deprecated since its function is replaced by
// the pimRPSetTable for PIM version 2.
type PIMMIB_Pimrptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimRPTable.  There is one entry per RP
    // address for each IP multicast group. The type is slice of
    // PIMMIB_Pimrptable_Pimrpentry.
    Pimrpentry []PIMMIB_Pimrptable_Pimrpentry
}

func (pimrptable *PIMMIB_Pimrptable) GetFilter() yfilter.YFilter { return pimrptable.YFilter }

func (pimrptable *PIMMIB_Pimrptable) SetFilter(yf yfilter.YFilter) { pimrptable.YFilter = yf }

func (pimrptable *PIMMIB_Pimrptable) GetGoName(yname string) string {
    if yname == "pimRPEntry" { return "Pimrpentry" }
    return ""
}

func (pimrptable *PIMMIB_Pimrptable) GetSegmentPath() string {
    return "pimRPTable"
}

func (pimrptable *PIMMIB_Pimrptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimRPEntry" {
        for _, c := range pimrptable.Pimrpentry {
            if pimrptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Pimrptable_Pimrpentry{}
        pimrptable.Pimrpentry = append(pimrptable.Pimrpentry, child)
        return &pimrptable.Pimrpentry[len(pimrptable.Pimrpentry)-1]
    }
    return nil
}

func (pimrptable *PIMMIB_Pimrptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pimrptable.Pimrpentry {
        children[pimrptable.Pimrpentry[i].GetSegmentPath()] = &pimrptable.Pimrpentry[i]
    }
    return children
}

func (pimrptable *PIMMIB_Pimrptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pimrptable *PIMMIB_Pimrptable) GetBundleName() string { return "cisco_ios_xe" }

func (pimrptable *PIMMIB_Pimrptable) GetYangName() string { return "pimRPTable" }

func (pimrptable *PIMMIB_Pimrptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimrptable *PIMMIB_Pimrptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimrptable *PIMMIB_Pimrptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimrptable *PIMMIB_Pimrptable) SetParent(parent types.Entity) { pimrptable.parent = parent }

func (pimrptable *PIMMIB_Pimrptable) GetParent() types.Entity { return pimrptable.parent }

func (pimrptable *PIMMIB_Pimrptable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pimrptable_Pimrpentry
// An entry (conceptual row) in the pimRPTable.  There is one
// entry per RP address for each IP multicast group.
type PIMMIB_Pimrptable_Pimrpentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains information about an RP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimrpgroupaddress interface{}

    // This attribute is a key. The unicast address of the RP. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimrpaddress interface{}

    // The state of the RP. The type is Pimrpstate.
    Pimrpstate interface{}

    // The minimum time remaining before the next state change. When pimRPState is
    // up, this is the minimum time which must expire until it can be declared
    // down.  When pimRPState is down, this is the time until it will be declared
    // up (in order to retry). The type is interface{} with range: 0..4294967295.
    Pimrpstatetimer interface{}

    // The value of sysUpTime at the time when the corresponding instance of
    // pimRPState last changed its value. The type is interface{} with range:
    // 0..4294967295.
    Pimrplastchange interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    Pimrprowstatus interface{}
}

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetFilter() yfilter.YFilter { return pimrpentry.YFilter }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) SetFilter(yf yfilter.YFilter) { pimrpentry.YFilter = yf }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetGoName(yname string) string {
    if yname == "pimRPGroupAddress" { return "Pimrpgroupaddress" }
    if yname == "pimRPAddress" { return "Pimrpaddress" }
    if yname == "pimRPState" { return "Pimrpstate" }
    if yname == "pimRPStateTimer" { return "Pimrpstatetimer" }
    if yname == "pimRPLastChange" { return "Pimrplastchange" }
    if yname == "pimRPRowStatus" { return "Pimrprowstatus" }
    return ""
}

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetSegmentPath() string {
    return "pimRPEntry" + "[pimRPGroupAddress='" + fmt.Sprintf("%v", pimrpentry.Pimrpgroupaddress) + "']" + "[pimRPAddress='" + fmt.Sprintf("%v", pimrpentry.Pimrpaddress) + "']"
}

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pimRPGroupAddress"] = pimrpentry.Pimrpgroupaddress
    leafs["pimRPAddress"] = pimrpentry.Pimrpaddress
    leafs["pimRPState"] = pimrpentry.Pimrpstate
    leafs["pimRPStateTimer"] = pimrpentry.Pimrpstatetimer
    leafs["pimRPLastChange"] = pimrpentry.Pimrplastchange
    leafs["pimRPRowStatus"] = pimrpentry.Pimrprowstatus
    return leafs
}

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetBundleName() string { return "cisco_ios_xe" }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetYangName() string { return "pimRPEntry" }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) SetParent(parent types.Entity) { pimrpentry.parent = parent }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetParent() types.Entity { return pimrpentry.parent }

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetParentYangName() string { return "pimRPTable" }

// PIMMIB_Pimrptable_Pimrpentry_Pimrpstate represents The state of the RP.
type PIMMIB_Pimrptable_Pimrpentry_Pimrpstate string

const (
    PIMMIB_Pimrptable_Pimrpentry_Pimrpstate_up PIMMIB_Pimrptable_Pimrpentry_Pimrpstate = "up"

    PIMMIB_Pimrptable_Pimrpentry_Pimrpstate_down PIMMIB_Pimrptable_Pimrpentry_Pimrpstate = "down"
)

// PIMMIB_Pimrpsettable
// The (conceptual) table listing PIM information for
// candidate Rendezvous Points (RPs) for IP multicast groups.
// When the local router is the BSR, this information is
// obtained from received Candidate-RP-Advertisements.  When
// the local router is not the BSR, this information is
// obtained from received RP-Set messages.
type PIMMIB_Pimrpsettable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimRPSetTable. The type is slice of
    // PIMMIB_Pimrpsettable_Pimrpsetentry.
    Pimrpsetentry []PIMMIB_Pimrpsettable_Pimrpsetentry
}

func (pimrpsettable *PIMMIB_Pimrpsettable) GetFilter() yfilter.YFilter { return pimrpsettable.YFilter }

func (pimrpsettable *PIMMIB_Pimrpsettable) SetFilter(yf yfilter.YFilter) { pimrpsettable.YFilter = yf }

func (pimrpsettable *PIMMIB_Pimrpsettable) GetGoName(yname string) string {
    if yname == "pimRPSetEntry" { return "Pimrpsetentry" }
    return ""
}

func (pimrpsettable *PIMMIB_Pimrpsettable) GetSegmentPath() string {
    return "pimRPSetTable"
}

func (pimrpsettable *PIMMIB_Pimrpsettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimRPSetEntry" {
        for _, c := range pimrpsettable.Pimrpsetentry {
            if pimrpsettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Pimrpsettable_Pimrpsetentry{}
        pimrpsettable.Pimrpsetentry = append(pimrpsettable.Pimrpsetentry, child)
        return &pimrpsettable.Pimrpsetentry[len(pimrpsettable.Pimrpsetentry)-1]
    }
    return nil
}

func (pimrpsettable *PIMMIB_Pimrpsettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pimrpsettable.Pimrpsetentry {
        children[pimrpsettable.Pimrpsetentry[i].GetSegmentPath()] = &pimrpsettable.Pimrpsetentry[i]
    }
    return children
}

func (pimrpsettable *PIMMIB_Pimrpsettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pimrpsettable *PIMMIB_Pimrpsettable) GetBundleName() string { return "cisco_ios_xe" }

func (pimrpsettable *PIMMIB_Pimrpsettable) GetYangName() string { return "pimRPSetTable" }

func (pimrpsettable *PIMMIB_Pimrpsettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimrpsettable *PIMMIB_Pimrpsettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimrpsettable *PIMMIB_Pimrpsettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimrpsettable *PIMMIB_Pimrpsettable) SetParent(parent types.Entity) { pimrpsettable.parent = parent }

func (pimrpsettable *PIMMIB_Pimrpsettable) GetParent() types.Entity { return pimrpsettable.parent }

func (pimrpsettable *PIMMIB_Pimrpsettable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pimrpsettable_Pimrpsetentry
// An entry (conceptual row) in the pimRPSetTable.
type PIMMIB_Pimrpsettable_Pimrpsetentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key.  A number uniquely identifying the component. 
    // Each protocol instance connected to a separate domain should have a
    // different index value. The type is interface{} with range: 1..255.
    Pimrpsetcomponent interface{}

    // This attribute is a key. The IP multicast group address which, when
    // combined with pimRPSetGroupMask, gives the group prefix for which this
    // entry contains information about the Candidate-RP. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimrpsetgroupaddress interface{}

    // This attribute is a key. The multicast group address mask which, when
    // combined with pimRPSetGroupAddress, gives the group prefix for which this
    // entry contains information about the Candidate-RP. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimrpsetgroupmask interface{}

    // This attribute is a key. The IP address of the Candidate-RP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimrpsetaddress interface{}

    // The holdtime of a Candidate-RP.  If the local router is not the BSR, this
    // value is 0. The type is interface{} with range: 0..255. Units are seconds.
    Pimrpsetholdtime interface{}

    // The minimum time remaining before the Candidate-RP will be declared down. 
    // If the local router is not the BSR, this value is 0. The type is
    // interface{} with range: 0..4294967295.
    Pimrpsetexpirytime interface{}
}

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetFilter() yfilter.YFilter { return pimrpsetentry.YFilter }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) SetFilter(yf yfilter.YFilter) { pimrpsetentry.YFilter = yf }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetGoName(yname string) string {
    if yname == "pimRPSetComponent" { return "Pimrpsetcomponent" }
    if yname == "pimRPSetGroupAddress" { return "Pimrpsetgroupaddress" }
    if yname == "pimRPSetGroupMask" { return "Pimrpsetgroupmask" }
    if yname == "pimRPSetAddress" { return "Pimrpsetaddress" }
    if yname == "pimRPSetHoldTime" { return "Pimrpsetholdtime" }
    if yname == "pimRPSetExpiryTime" { return "Pimrpsetexpirytime" }
    return ""
}

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetSegmentPath() string {
    return "pimRPSetEntry" + "[pimRPSetComponent='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetcomponent) + "']" + "[pimRPSetGroupAddress='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetgroupaddress) + "']" + "[pimRPSetGroupMask='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetgroupmask) + "']" + "[pimRPSetAddress='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetaddress) + "']"
}

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pimRPSetComponent"] = pimrpsetentry.Pimrpsetcomponent
    leafs["pimRPSetGroupAddress"] = pimrpsetentry.Pimrpsetgroupaddress
    leafs["pimRPSetGroupMask"] = pimrpsetentry.Pimrpsetgroupmask
    leafs["pimRPSetAddress"] = pimrpsetentry.Pimrpsetaddress
    leafs["pimRPSetHoldTime"] = pimrpsetentry.Pimrpsetholdtime
    leafs["pimRPSetExpiryTime"] = pimrpsetentry.Pimrpsetexpirytime
    return leafs
}

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetBundleName() string { return "cisco_ios_xe" }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetYangName() string { return "pimRPSetEntry" }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) SetParent(parent types.Entity) { pimrpsetentry.parent = parent }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetParent() types.Entity { return pimrpsetentry.parent }

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetParentYangName() string { return "pimRPSetTable" }

// PIMMIB_Pimipmroutenexthoptable
// The (conceptual) table listing PIM-specific information on
// a subset of the rows of the ipMRouteNextHopTable defined in
// the IP Multicast MIB.
type PIMMIB_Pimipmroutenexthoptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimIpMRouteNextHopTable. There is one
    // entry per entry in the ipMRouteNextHopTable whose interface is running PIM
    // and whose ipMRouteNextHopState is pruned(1). The type is slice of
    // PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry.
    Pimipmroutenexthopentry []PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry
}

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetFilter() yfilter.YFilter { return pimipmroutenexthoptable.YFilter }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) SetFilter(yf yfilter.YFilter) { pimipmroutenexthoptable.YFilter = yf }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetGoName(yname string) string {
    if yname == "pimIpMRouteNextHopEntry" { return "Pimipmroutenexthopentry" }
    return ""
}

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetSegmentPath() string {
    return "pimIpMRouteNextHopTable"
}

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimIpMRouteNextHopEntry" {
        for _, c := range pimipmroutenexthoptable.Pimipmroutenexthopentry {
            if pimipmroutenexthoptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry{}
        pimipmroutenexthoptable.Pimipmroutenexthopentry = append(pimipmroutenexthoptable.Pimipmroutenexthopentry, child)
        return &pimipmroutenexthoptable.Pimipmroutenexthopentry[len(pimipmroutenexthoptable.Pimipmroutenexthopentry)-1]
    }
    return nil
}

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pimipmroutenexthoptable.Pimipmroutenexthopentry {
        children[pimipmroutenexthoptable.Pimipmroutenexthopentry[i].GetSegmentPath()] = &pimipmroutenexthoptable.Pimipmroutenexthopentry[i]
    }
    return children
}

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetBundleName() string { return "cisco_ios_xe" }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetYangName() string { return "pimIpMRouteNextHopTable" }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) SetParent(parent types.Entity) { pimipmroutenexthoptable.parent = parent }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetParent() types.Entity { return pimipmroutenexthoptable.parent }

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry
// An entry (conceptual row) in the pimIpMRouteNextHopTable.
// There is one entry per entry in the ipMRouteNextHopTable
// whose interface is running PIM and whose
// ipMRouteNextHopState is pruned(1).
type PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopgroup
    Ipmroutenexthopgroup interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopsource
    Ipmroutenexthopsource interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopsourcemask
    Ipmroutenexthopsourcemask interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopifindex
    Ipmroutenexthopifindex interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopaddress
    Ipmroutenexthopaddress interface{}

    // This object indicates why the downstream interface was pruned, whether in
    // response to a PIM prune message or due to PIM Assert processing. The type
    // is Pimipmroutenexthopprunereason.
    Pimipmroutenexthopprunereason interface{}
}

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetFilter() yfilter.YFilter { return pimipmroutenexthopentry.YFilter }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) SetFilter(yf yfilter.YFilter) { pimipmroutenexthopentry.YFilter = yf }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetGoName(yname string) string {
    if yname == "ipMRouteNextHopGroup" { return "Ipmroutenexthopgroup" }
    if yname == "ipMRouteNextHopSource" { return "Ipmroutenexthopsource" }
    if yname == "ipMRouteNextHopSourceMask" { return "Ipmroutenexthopsourcemask" }
    if yname == "ipMRouteNextHopIfIndex" { return "Ipmroutenexthopifindex" }
    if yname == "ipMRouteNextHopAddress" { return "Ipmroutenexthopaddress" }
    if yname == "pimIpMRouteNextHopPruneReason" { return "Pimipmroutenexthopprunereason" }
    return ""
}

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetSegmentPath() string {
    return "pimIpMRouteNextHopEntry" + "[ipMRouteNextHopGroup='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopgroup) + "']" + "[ipMRouteNextHopSource='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopsource) + "']" + "[ipMRouteNextHopSourceMask='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopsourcemask) + "']" + "[ipMRouteNextHopIfIndex='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopifindex) + "']" + "[ipMRouteNextHopAddress='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopaddress) + "']"
}

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipMRouteNextHopGroup"] = pimipmroutenexthopentry.Ipmroutenexthopgroup
    leafs["ipMRouteNextHopSource"] = pimipmroutenexthopentry.Ipmroutenexthopsource
    leafs["ipMRouteNextHopSourceMask"] = pimipmroutenexthopentry.Ipmroutenexthopsourcemask
    leafs["ipMRouteNextHopIfIndex"] = pimipmroutenexthopentry.Ipmroutenexthopifindex
    leafs["ipMRouteNextHopAddress"] = pimipmroutenexthopentry.Ipmroutenexthopaddress
    leafs["pimIpMRouteNextHopPruneReason"] = pimipmroutenexthopentry.Pimipmroutenexthopprunereason
    return leafs
}

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetBundleName() string { return "cisco_ios_xe" }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetYangName() string { return "pimIpMRouteNextHopEntry" }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) SetParent(parent types.Entity) { pimipmroutenexthopentry.parent = parent }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetParent() types.Entity { return pimipmroutenexthopentry.parent }

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetParentYangName() string { return "pimIpMRouteNextHopTable" }

// PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason represents PIM Assert processing.
type PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason string

const (
    PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason_other PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason = "other"

    PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason_prune PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason = "prune"

    PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason_assert PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry_Pimipmroutenexthopprunereason = "assert"
)

// PIMMIB_Pimcandidaterptable
// The (conceptual) table listing the IP multicast groups for
// which the local router is to advertise itself as a
// Candidate-RP when the value of pimComponentCRPHoldTime is
// non-zero.  If this table is empty, then the local router
// 
// 
// 
// 
// 
// will advertise itself as a Candidate-RP for all groups
// (providing the value of pimComponentCRPHoldTime is non-
// zero).
type PIMMIB_Pimcandidaterptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimCandidateRPTable. The type is slice of
    // PIMMIB_Pimcandidaterptable_Pimcandidaterpentry.
    Pimcandidaterpentry []PIMMIB_Pimcandidaterptable_Pimcandidaterpentry
}

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetFilter() yfilter.YFilter { return pimcandidaterptable.YFilter }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) SetFilter(yf yfilter.YFilter) { pimcandidaterptable.YFilter = yf }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetGoName(yname string) string {
    if yname == "pimCandidateRPEntry" { return "Pimcandidaterpentry" }
    return ""
}

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetSegmentPath() string {
    return "pimCandidateRPTable"
}

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimCandidateRPEntry" {
        for _, c := range pimcandidaterptable.Pimcandidaterpentry {
            if pimcandidaterptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Pimcandidaterptable_Pimcandidaterpentry{}
        pimcandidaterptable.Pimcandidaterpentry = append(pimcandidaterptable.Pimcandidaterpentry, child)
        return &pimcandidaterptable.Pimcandidaterpentry[len(pimcandidaterptable.Pimcandidaterpentry)-1]
    }
    return nil
}

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pimcandidaterptable.Pimcandidaterpentry {
        children[pimcandidaterptable.Pimcandidaterpentry[i].GetSegmentPath()] = &pimcandidaterptable.Pimcandidaterpentry[i]
    }
    return children
}

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetBundleName() string { return "cisco_ios_xe" }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetYangName() string { return "pimCandidateRPTable" }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) SetParent(parent types.Entity) { pimcandidaterptable.parent = parent }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetParent() types.Entity { return pimcandidaterptable.parent }

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pimcandidaterptable_Pimcandidaterpentry
// An entry (conceptual row) in the pimCandidateRPTable.
type PIMMIB_Pimcandidaterptable_Pimcandidaterpentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address which, when
    // combined with pimCandidateRPGroupMask, identifies a group prefix for which
    // the local router will advertise itself as a Candidate-RP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimcandidaterpgroupaddress interface{}

    // This attribute is a key. The multicast group address mask which, when
    // combined with pimCandidateRPGroupMask, identifies a group prefix for which
    // the local router will advertise itself as a Candidate-RP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimcandidaterpgroupmask interface{}

    // The (unicast) address of the interface which will be      advertised as a
    // Candidate-RP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimcandidaterpaddress interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    Pimcandidaterprowstatus interface{}
}

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetFilter() yfilter.YFilter { return pimcandidaterpentry.YFilter }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) SetFilter(yf yfilter.YFilter) { pimcandidaterpentry.YFilter = yf }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetGoName(yname string) string {
    if yname == "pimCandidateRPGroupAddress" { return "Pimcandidaterpgroupaddress" }
    if yname == "pimCandidateRPGroupMask" { return "Pimcandidaterpgroupmask" }
    if yname == "pimCandidateRPAddress" { return "Pimcandidaterpaddress" }
    if yname == "pimCandidateRPRowStatus" { return "Pimcandidaterprowstatus" }
    return ""
}

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetSegmentPath() string {
    return "pimCandidateRPEntry" + "[pimCandidateRPGroupAddress='" + fmt.Sprintf("%v", pimcandidaterpentry.Pimcandidaterpgroupaddress) + "']" + "[pimCandidateRPGroupMask='" + fmt.Sprintf("%v", pimcandidaterpentry.Pimcandidaterpgroupmask) + "']"
}

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pimCandidateRPGroupAddress"] = pimcandidaterpentry.Pimcandidaterpgroupaddress
    leafs["pimCandidateRPGroupMask"] = pimcandidaterpentry.Pimcandidaterpgroupmask
    leafs["pimCandidateRPAddress"] = pimcandidaterpentry.Pimcandidaterpaddress
    leafs["pimCandidateRPRowStatus"] = pimcandidaterpentry.Pimcandidaterprowstatus
    return leafs
}

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetBundleName() string { return "cisco_ios_xe" }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetYangName() string { return "pimCandidateRPEntry" }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) SetParent(parent types.Entity) { pimcandidaterpentry.parent = parent }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetParent() types.Entity { return pimcandidaterpentry.parent }

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetParentYangName() string { return "pimCandidateRPTable" }

// PIMMIB_Pimcomponenttable
// The (conceptual) table containing objects specific to a PIM
// domain.  One row exists for each domain to which the router
// is connected.  A PIM-SM domain is defined as an area of the
// network over which Bootstrap messages are forwarded.
// Typically, a PIM-SM router will be a member of exactly one
// domain.  This table also supports, however, routers which
// may form a border between two PIM-SM domains and do not
// forward Bootstrap messages between them.
type PIMMIB_Pimcomponenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimComponentTable. The type is slice of
    // PIMMIB_Pimcomponenttable_Pimcomponententry.
    Pimcomponententry []PIMMIB_Pimcomponenttable_Pimcomponententry
}

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetFilter() yfilter.YFilter { return pimcomponenttable.YFilter }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) SetFilter(yf yfilter.YFilter) { pimcomponenttable.YFilter = yf }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetGoName(yname string) string {
    if yname == "pimComponentEntry" { return "Pimcomponententry" }
    return ""
}

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetSegmentPath() string {
    return "pimComponentTable"
}

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pimComponentEntry" {
        for _, c := range pimcomponenttable.Pimcomponententry {
            if pimcomponenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PIMMIB_Pimcomponenttable_Pimcomponententry{}
        pimcomponenttable.Pimcomponententry = append(pimcomponenttable.Pimcomponententry, child)
        return &pimcomponenttable.Pimcomponententry[len(pimcomponenttable.Pimcomponententry)-1]
    }
    return nil
}

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pimcomponenttable.Pimcomponententry {
        children[pimcomponenttable.Pimcomponententry[i].GetSegmentPath()] = &pimcomponenttable.Pimcomponententry[i]
    }
    return children
}

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetBundleName() string { return "cisco_ios_xe" }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetYangName() string { return "pimComponentTable" }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) SetParent(parent types.Entity) { pimcomponenttable.parent = parent }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetParent() types.Entity { return pimcomponenttable.parent }

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetParentYangName() string { return "PIM-MIB" }

// PIMMIB_Pimcomponenttable_Pimcomponententry
// An entry (conceptual row) in the pimComponentTable.
type PIMMIB_Pimcomponenttable_Pimcomponententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A number uniquely identifying the component.  Each
    // protocol instance connected to a separate domain should have a different
    // index value.  Routers that only support membership in a single PIM-SM
    // domain should use a pimComponentIndex value of 1. The type is interface{}
    // with range: 1..255.
    Pimcomponentindex interface{}

    // The IP address of the bootstrap router (BSR) for the local PIM region. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Pimcomponentbsraddress interface{}

    // The minimum time remaining before the bootstrap router in the local domain
    // will be declared down.  For candidate BSRs, this is the time until the
    // component sends an RP-Set message.  For other routers, this is the time
    // until it may accept an RP-Set message from a lower candidate BSR. The type
    // is interface{} with range: 0..4294967295.
    Pimcomponentbsrexpirytime interface{}

    // The holdtime of the component when it is a candidate RP in the local
    // domain.  The value of 0 is used to indicate that the local system is not a
    // Candidate-RP. The type is interface{} with range: 0..255. Units are
    // seconds.
    Pimcomponentcrpholdtime interface{}

    // The status of this entry.  Creating the entry creates another protocol
    // instance; destroying the entry disables a protocol instance. The type is
    // RowStatus.
    Pimcomponentstatus interface{}
}

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetFilter() yfilter.YFilter { return pimcomponententry.YFilter }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) SetFilter(yf yfilter.YFilter) { pimcomponententry.YFilter = yf }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetGoName(yname string) string {
    if yname == "pimComponentIndex" { return "Pimcomponentindex" }
    if yname == "pimComponentBSRAddress" { return "Pimcomponentbsraddress" }
    if yname == "pimComponentBSRExpiryTime" { return "Pimcomponentbsrexpirytime" }
    if yname == "pimComponentCRPHoldTime" { return "Pimcomponentcrpholdtime" }
    if yname == "pimComponentStatus" { return "Pimcomponentstatus" }
    return ""
}

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetSegmentPath() string {
    return "pimComponentEntry" + "[pimComponentIndex='" + fmt.Sprintf("%v", pimcomponententry.Pimcomponentindex) + "']"
}

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pimComponentIndex"] = pimcomponententry.Pimcomponentindex
    leafs["pimComponentBSRAddress"] = pimcomponententry.Pimcomponentbsraddress
    leafs["pimComponentBSRExpiryTime"] = pimcomponententry.Pimcomponentbsrexpirytime
    leafs["pimComponentCRPHoldTime"] = pimcomponententry.Pimcomponentcrpholdtime
    leafs["pimComponentStatus"] = pimcomponententry.Pimcomponentstatus
    return leafs
}

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetBundleName() string { return "cisco_ios_xe" }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetYangName() string { return "pimComponentEntry" }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) SetParent(parent types.Entity) { pimcomponententry.parent = parent }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetParent() types.Entity { return pimcomponententry.parent }

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetParentYangName() string { return "pimComponentTable" }

