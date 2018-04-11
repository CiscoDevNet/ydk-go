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
    EntityData types.CommonEntityData
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

func (pIMMIB *PIMMIB) GetEntityData() *types.CommonEntityData {
    pIMMIB.EntityData.YFilter = pIMMIB.YFilter
    pIMMIB.EntityData.YangName = "PIM-MIB"
    pIMMIB.EntityData.BundleName = "cisco_ios_xe"
    pIMMIB.EntityData.ParentYangName = "PIM-MIB"
    pIMMIB.EntityData.SegmentPath = "PIM-MIB:PIM-MIB"
    pIMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pIMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pIMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pIMMIB.EntityData.Children = make(map[string]types.YChild)
    pIMMIB.EntityData.Children["pim"] = types.YChild{"Pim", &pIMMIB.Pim}
    pIMMIB.EntityData.Children["pimInterfaceTable"] = types.YChild{"Piminterfacetable", &pIMMIB.Piminterfacetable}
    pIMMIB.EntityData.Children["pimNeighborTable"] = types.YChild{"Pimneighbortable", &pIMMIB.Pimneighbortable}
    pIMMIB.EntityData.Children["pimIpMRouteTable"] = types.YChild{"Pimipmroutetable", &pIMMIB.Pimipmroutetable}
    pIMMIB.EntityData.Children["pimRPTable"] = types.YChild{"Pimrptable", &pIMMIB.Pimrptable}
    pIMMIB.EntityData.Children["pimRPSetTable"] = types.YChild{"Pimrpsettable", &pIMMIB.Pimrpsettable}
    pIMMIB.EntityData.Children["pimIpMRouteNextHopTable"] = types.YChild{"Pimipmroutenexthoptable", &pIMMIB.Pimipmroutenexthoptable}
    pIMMIB.EntityData.Children["pimCandidateRPTable"] = types.YChild{"Pimcandidaterptable", &pIMMIB.Pimcandidaterptable}
    pIMMIB.EntityData.Children["pimComponentTable"] = types.YChild{"Pimcomponenttable", &pIMMIB.Pimcomponenttable}
    pIMMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pIMMIB.EntityData)
}

// PIMMIB_Pim
type PIMMIB_Pim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The default interval at which periodic PIM-SM Join/Prune messages are to be
    // sent. The type is interface{} with range: -2147483648..2147483647. Units
    // are seconds.
    Pimjoinpruneinterval interface{}
}

func (pim *PIMMIB_Pim) GetEntityData() *types.CommonEntityData {
    pim.EntityData.YFilter = pim.YFilter
    pim.EntityData.YangName = "pim"
    pim.EntityData.BundleName = "cisco_ios_xe"
    pim.EntityData.ParentYangName = "PIM-MIB"
    pim.EntityData.SegmentPath = "pim"
    pim.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pim.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pim.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pim.EntityData.Children = make(map[string]types.YChild)
    pim.EntityData.Leafs = make(map[string]types.YLeaf)
    pim.EntityData.Leafs["pimJoinPruneInterval"] = types.YLeaf{"Pimjoinpruneinterval", pim.Pimjoinpruneinterval}
    return &(pim.EntityData)
}

// PIMMIB_Piminterfacetable
// The (conceptual) table listing the router's PIM interfaces.
// IGMP and PIM are enabled on all interfaces listed in this
// table.
type PIMMIB_Piminterfacetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimInterfaceTable. The type is slice of
    // PIMMIB_Piminterfacetable_Piminterfaceentry.
    Piminterfaceentry []PIMMIB_Piminterfacetable_Piminterfaceentry
}

func (piminterfacetable *PIMMIB_Piminterfacetable) GetEntityData() *types.CommonEntityData {
    piminterfacetable.EntityData.YFilter = piminterfacetable.YFilter
    piminterfacetable.EntityData.YangName = "pimInterfaceTable"
    piminterfacetable.EntityData.BundleName = "cisco_ios_xe"
    piminterfacetable.EntityData.ParentYangName = "PIM-MIB"
    piminterfacetable.EntityData.SegmentPath = "pimInterfaceTable"
    piminterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    piminterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    piminterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    piminterfacetable.EntityData.Children = make(map[string]types.YChild)
    piminterfacetable.EntityData.Children["pimInterfaceEntry"] = types.YChild{"Piminterfaceentry", nil}
    for i := range piminterfacetable.Piminterfaceentry {
        piminterfacetable.EntityData.Children[types.GetSegmentPath(&piminterfacetable.Piminterfaceentry[i])] = types.YChild{"Piminterfaceentry", &piminterfacetable.Piminterfaceentry[i]}
    }
    piminterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(piminterfacetable.EntityData)
}

// PIMMIB_Piminterfacetable_Piminterfaceentry
// An entry (conceptual row) in the pimInterfaceTable.
type PIMMIB_Piminterfacetable_Piminterfaceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of this PIM interface. The type
    // is interface{} with range: 1..2147483647.
    Piminterfaceifindex interface{}

    // The IP address of the PIM interface. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Piminterfaceaddress interface{}

    // The network mask for the IP address of the PIM interface. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Piminterfacenetmask interface{}

    // The configured mode of this PIM interface.  A value of sparseDense is only
    // valid for PIMv1. The type is Piminterfacemode.
    Piminterfacemode interface{}

    // The Designated Router on this PIM interface.  For point-to- point
    // interfaces, this object has the value 0.0.0.0. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (piminterfaceentry *PIMMIB_Piminterfacetable_Piminterfaceentry) GetEntityData() *types.CommonEntityData {
    piminterfaceentry.EntityData.YFilter = piminterfaceentry.YFilter
    piminterfaceentry.EntityData.YangName = "pimInterfaceEntry"
    piminterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    piminterfaceentry.EntityData.ParentYangName = "pimInterfaceTable"
    piminterfaceentry.EntityData.SegmentPath = "pimInterfaceEntry" + "[pimInterfaceIfIndex='" + fmt.Sprintf("%v", piminterfaceentry.Piminterfaceifindex) + "']"
    piminterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    piminterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    piminterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    piminterfaceentry.EntityData.Children = make(map[string]types.YChild)
    piminterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    piminterfaceentry.EntityData.Leafs["pimInterfaceIfIndex"] = types.YLeaf{"Piminterfaceifindex", piminterfaceentry.Piminterfaceifindex}
    piminterfaceentry.EntityData.Leafs["pimInterfaceAddress"] = types.YLeaf{"Piminterfaceaddress", piminterfaceentry.Piminterfaceaddress}
    piminterfaceentry.EntityData.Leafs["pimInterfaceNetMask"] = types.YLeaf{"Piminterfacenetmask", piminterfaceentry.Piminterfacenetmask}
    piminterfaceentry.EntityData.Leafs["pimInterfaceMode"] = types.YLeaf{"Piminterfacemode", piminterfaceentry.Piminterfacemode}
    piminterfaceentry.EntityData.Leafs["pimInterfaceDR"] = types.YLeaf{"Piminterfacedr", piminterfaceentry.Piminterfacedr}
    piminterfaceentry.EntityData.Leafs["pimInterfaceHelloInterval"] = types.YLeaf{"Piminterfacehellointerval", piminterfaceentry.Piminterfacehellointerval}
    piminterfaceentry.EntityData.Leafs["pimInterfaceStatus"] = types.YLeaf{"Piminterfacestatus", piminterfaceentry.Piminterfacestatus}
    piminterfaceentry.EntityData.Leafs["pimInterfaceJoinPruneInterval"] = types.YLeaf{"Piminterfacejoinpruneinterval", piminterfaceentry.Piminterfacejoinpruneinterval}
    piminterfaceentry.EntityData.Leafs["pimInterfaceCBSRPreference"] = types.YLeaf{"Piminterfacecbsrpreference", piminterfaceentry.Piminterfacecbsrpreference}
    return &(piminterfaceentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimNeighborTable. The type is slice of
    // PIMMIB_Pimneighbortable_Pimneighborentry.
    Pimneighborentry []PIMMIB_Pimneighbortable_Pimneighborentry
}

func (pimneighbortable *PIMMIB_Pimneighbortable) GetEntityData() *types.CommonEntityData {
    pimneighbortable.EntityData.YFilter = pimneighbortable.YFilter
    pimneighbortable.EntityData.YangName = "pimNeighborTable"
    pimneighbortable.EntityData.BundleName = "cisco_ios_xe"
    pimneighbortable.EntityData.ParentYangName = "PIM-MIB"
    pimneighbortable.EntityData.SegmentPath = "pimNeighborTable"
    pimneighbortable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimneighbortable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimneighbortable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimneighbortable.EntityData.Children = make(map[string]types.YChild)
    pimneighbortable.EntityData.Children["pimNeighborEntry"] = types.YChild{"Pimneighborentry", nil}
    for i := range pimneighbortable.Pimneighborentry {
        pimneighbortable.EntityData.Children[types.GetSegmentPath(&pimneighbortable.Pimneighborentry[i])] = types.YChild{"Pimneighborentry", &pimneighbortable.Pimneighborentry[i]}
    }
    pimneighbortable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pimneighbortable.EntityData)
}

// PIMMIB_Pimneighbortable_Pimneighborentry
// An entry (conceptual row) in the pimNeighborTable.
type PIMMIB_Pimneighbortable_Pimneighborentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP address of the PIM neighbor for which this
    // entry contains information. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (pimneighborentry *PIMMIB_Pimneighbortable_Pimneighborentry) GetEntityData() *types.CommonEntityData {
    pimneighborentry.EntityData.YFilter = pimneighborentry.YFilter
    pimneighborentry.EntityData.YangName = "pimNeighborEntry"
    pimneighborentry.EntityData.BundleName = "cisco_ios_xe"
    pimneighborentry.EntityData.ParentYangName = "pimNeighborTable"
    pimneighborentry.EntityData.SegmentPath = "pimNeighborEntry" + "[pimNeighborAddress='" + fmt.Sprintf("%v", pimneighborentry.Pimneighboraddress) + "']"
    pimneighborentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimneighborentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimneighborentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimneighborentry.EntityData.Children = make(map[string]types.YChild)
    pimneighborentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pimneighborentry.EntityData.Leafs["pimNeighborAddress"] = types.YLeaf{"Pimneighboraddress", pimneighborentry.Pimneighboraddress}
    pimneighborentry.EntityData.Leafs["pimNeighborIfIndex"] = types.YLeaf{"Pimneighborifindex", pimneighborentry.Pimneighborifindex}
    pimneighborentry.EntityData.Leafs["pimNeighborUpTime"] = types.YLeaf{"Pimneighboruptime", pimneighborentry.Pimneighboruptime}
    pimneighborentry.EntityData.Leafs["pimNeighborExpiryTime"] = types.YLeaf{"Pimneighborexpirytime", pimneighborentry.Pimneighborexpirytime}
    pimneighborentry.EntityData.Leafs["pimNeighborMode"] = types.YLeaf{"Pimneighbormode", pimneighborentry.Pimneighbormode}
    return &(pimneighborentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimIpMRouteTable.  There is one entry per
    // entry in the ipMRouteTable whose incoming interface is running PIM. The
    // type is slice of PIMMIB_Pimipmroutetable_Pimipmrouteentry.
    Pimipmrouteentry []PIMMIB_Pimipmroutetable_Pimipmrouteentry
}

func (pimipmroutetable *PIMMIB_Pimipmroutetable) GetEntityData() *types.CommonEntityData {
    pimipmroutetable.EntityData.YFilter = pimipmroutetable.YFilter
    pimipmroutetable.EntityData.YangName = "pimIpMRouteTable"
    pimipmroutetable.EntityData.BundleName = "cisco_ios_xe"
    pimipmroutetable.EntityData.ParentYangName = "PIM-MIB"
    pimipmroutetable.EntityData.SegmentPath = "pimIpMRouteTable"
    pimipmroutetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimipmroutetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimipmroutetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimipmroutetable.EntityData.Children = make(map[string]types.YChild)
    pimipmroutetable.EntityData.Children["pimIpMRouteEntry"] = types.YChild{"Pimipmrouteentry", nil}
    for i := range pimipmroutetable.Pimipmrouteentry {
        pimipmroutetable.EntityData.Children[types.GetSegmentPath(&pimipmroutetable.Pimipmrouteentry[i])] = types.YChild{"Pimipmrouteentry", &pimipmroutetable.Pimipmrouteentry[i]}
    }
    pimipmroutetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pimipmroutetable.EntityData)
}

// PIMMIB_Pimipmroutetable_Pimipmrouteentry
// An entry (conceptual row) in the pimIpMRouteTable.  There
// is one entry per entry in the ipMRouteTable whose incoming
// interface is running PIM.
type PIMMIB_Pimipmroutetable_Pimipmrouteentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmroutegroup
    Ipmroutegroup interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutetable_Ipmrouteentry_Ipmroutesource
    Ipmroutesource interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (pimipmrouteentry *PIMMIB_Pimipmroutetable_Pimipmrouteentry) GetEntityData() *types.CommonEntityData {
    pimipmrouteentry.EntityData.YFilter = pimipmrouteentry.YFilter
    pimipmrouteentry.EntityData.YangName = "pimIpMRouteEntry"
    pimipmrouteentry.EntityData.BundleName = "cisco_ios_xe"
    pimipmrouteentry.EntityData.ParentYangName = "pimIpMRouteTable"
    pimipmrouteentry.EntityData.SegmentPath = "pimIpMRouteEntry" + "[ipMRouteGroup='" + fmt.Sprintf("%v", pimipmrouteentry.Ipmroutegroup) + "']" + "[ipMRouteSource='" + fmt.Sprintf("%v", pimipmrouteentry.Ipmroutesource) + "']" + "[ipMRouteSourceMask='" + fmt.Sprintf("%v", pimipmrouteentry.Ipmroutesourcemask) + "']"
    pimipmrouteentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimipmrouteentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimipmrouteentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimipmrouteentry.EntityData.Children = make(map[string]types.YChild)
    pimipmrouteentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pimipmrouteentry.EntityData.Leafs["ipMRouteGroup"] = types.YLeaf{"Ipmroutegroup", pimipmrouteentry.Ipmroutegroup}
    pimipmrouteentry.EntityData.Leafs["ipMRouteSource"] = types.YLeaf{"Ipmroutesource", pimipmrouteentry.Ipmroutesource}
    pimipmrouteentry.EntityData.Leafs["ipMRouteSourceMask"] = types.YLeaf{"Ipmroutesourcemask", pimipmrouteentry.Ipmroutesourcemask}
    pimipmrouteentry.EntityData.Leafs["pimIpMRouteUpstreamAssertTimer"] = types.YLeaf{"Pimipmrouteupstreamasserttimer", pimipmrouteentry.Pimipmrouteupstreamasserttimer}
    pimipmrouteentry.EntityData.Leafs["pimIpMRouteAssertMetric"] = types.YLeaf{"Pimipmrouteassertmetric", pimipmrouteentry.Pimipmrouteassertmetric}
    pimipmrouteentry.EntityData.Leafs["pimIpMRouteAssertMetricPref"] = types.YLeaf{"Pimipmrouteassertmetricpref", pimipmrouteentry.Pimipmrouteassertmetricpref}
    pimipmrouteentry.EntityData.Leafs["pimIpMRouteAssertRPTBit"] = types.YLeaf{"Pimipmrouteassertrptbit", pimipmrouteentry.Pimipmrouteassertrptbit}
    pimipmrouteentry.EntityData.Leafs["pimIpMRouteFlags"] = types.YLeaf{"Pimipmrouteflags", pimipmrouteentry.Pimipmrouteflags}
    return &(pimipmrouteentry.EntityData)
}

// PIMMIB_Pimrptable
// The (conceptual) table listing PIM version 1 information
// for the Rendezvous Points (RPs) for IP multicast groups.
// This table is deprecated since its function is replaced by
// the pimRPSetTable for PIM version 2.
type PIMMIB_Pimrptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimRPTable.  There is one entry per RP
    // address for each IP multicast group. The type is slice of
    // PIMMIB_Pimrptable_Pimrpentry.
    Pimrpentry []PIMMIB_Pimrptable_Pimrpentry
}

func (pimrptable *PIMMIB_Pimrptable) GetEntityData() *types.CommonEntityData {
    pimrptable.EntityData.YFilter = pimrptable.YFilter
    pimrptable.EntityData.YangName = "pimRPTable"
    pimrptable.EntityData.BundleName = "cisco_ios_xe"
    pimrptable.EntityData.ParentYangName = "PIM-MIB"
    pimrptable.EntityData.SegmentPath = "pimRPTable"
    pimrptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimrptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimrptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimrptable.EntityData.Children = make(map[string]types.YChild)
    pimrptable.EntityData.Children["pimRPEntry"] = types.YChild{"Pimrpentry", nil}
    for i := range pimrptable.Pimrpentry {
        pimrptable.EntityData.Children[types.GetSegmentPath(&pimrptable.Pimrpentry[i])] = types.YChild{"Pimrpentry", &pimrptable.Pimrpentry[i]}
    }
    pimrptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pimrptable.EntityData)
}

// PIMMIB_Pimrptable_Pimrpentry
// An entry (conceptual row) in the pimRPTable.  There is one
// entry per RP address for each IP multicast group.
type PIMMIB_Pimrptable_Pimrpentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains information about an RP. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Pimrpgroupaddress interface{}

    // This attribute is a key. The unicast address of the RP. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (pimrpentry *PIMMIB_Pimrptable_Pimrpentry) GetEntityData() *types.CommonEntityData {
    pimrpentry.EntityData.YFilter = pimrpentry.YFilter
    pimrpentry.EntityData.YangName = "pimRPEntry"
    pimrpentry.EntityData.BundleName = "cisco_ios_xe"
    pimrpentry.EntityData.ParentYangName = "pimRPTable"
    pimrpentry.EntityData.SegmentPath = "pimRPEntry" + "[pimRPGroupAddress='" + fmt.Sprintf("%v", pimrpentry.Pimrpgroupaddress) + "']" + "[pimRPAddress='" + fmt.Sprintf("%v", pimrpentry.Pimrpaddress) + "']"
    pimrpentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimrpentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimrpentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimrpentry.EntityData.Children = make(map[string]types.YChild)
    pimrpentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pimrpentry.EntityData.Leafs["pimRPGroupAddress"] = types.YLeaf{"Pimrpgroupaddress", pimrpentry.Pimrpgroupaddress}
    pimrpentry.EntityData.Leafs["pimRPAddress"] = types.YLeaf{"Pimrpaddress", pimrpentry.Pimrpaddress}
    pimrpentry.EntityData.Leafs["pimRPState"] = types.YLeaf{"Pimrpstate", pimrpentry.Pimrpstate}
    pimrpentry.EntityData.Leafs["pimRPStateTimer"] = types.YLeaf{"Pimrpstatetimer", pimrpentry.Pimrpstatetimer}
    pimrpentry.EntityData.Leafs["pimRPLastChange"] = types.YLeaf{"Pimrplastchange", pimrpentry.Pimrplastchange}
    pimrpentry.EntityData.Leafs["pimRPRowStatus"] = types.YLeaf{"Pimrprowstatus", pimrpentry.Pimrprowstatus}
    return &(pimrpentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimRPSetTable. The type is slice of
    // PIMMIB_Pimrpsettable_Pimrpsetentry.
    Pimrpsetentry []PIMMIB_Pimrpsettable_Pimrpsetentry
}

func (pimrpsettable *PIMMIB_Pimrpsettable) GetEntityData() *types.CommonEntityData {
    pimrpsettable.EntityData.YFilter = pimrpsettable.YFilter
    pimrpsettable.EntityData.YangName = "pimRPSetTable"
    pimrpsettable.EntityData.BundleName = "cisco_ios_xe"
    pimrpsettable.EntityData.ParentYangName = "PIM-MIB"
    pimrpsettable.EntityData.SegmentPath = "pimRPSetTable"
    pimrpsettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimrpsettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimrpsettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimrpsettable.EntityData.Children = make(map[string]types.YChild)
    pimrpsettable.EntityData.Children["pimRPSetEntry"] = types.YChild{"Pimrpsetentry", nil}
    for i := range pimrpsettable.Pimrpsetentry {
        pimrpsettable.EntityData.Children[types.GetSegmentPath(&pimrpsettable.Pimrpsetentry[i])] = types.YChild{"Pimrpsetentry", &pimrpsettable.Pimrpsetentry[i]}
    }
    pimrpsettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pimrpsettable.EntityData)
}

// PIMMIB_Pimrpsettable_Pimrpsetentry
// An entry (conceptual row) in the pimRPSetTable.
type PIMMIB_Pimrpsettable_Pimrpsetentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key.  A number uniquely identifying the component. 
    // Each protocol instance connected to a separate domain should have a
    // different index value. The type is interface{} with range: 1..255.
    Pimrpsetcomponent interface{}

    // This attribute is a key. The IP multicast group address which, when
    // combined with pimRPSetGroupMask, gives the group prefix for which this
    // entry contains information about the Candidate-RP. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Pimrpsetgroupaddress interface{}

    // This attribute is a key. The multicast group address mask which, when
    // combined with pimRPSetGroupAddress, gives the group prefix for which this
    // entry contains information about the Candidate-RP. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Pimrpsetgroupmask interface{}

    // This attribute is a key. The IP address of the Candidate-RP. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Pimrpsetaddress interface{}

    // The holdtime of a Candidate-RP.  If the local router is not the BSR, this
    // value is 0. The type is interface{} with range: 0..255. Units are seconds.
    Pimrpsetholdtime interface{}

    // The minimum time remaining before the Candidate-RP will be declared down. 
    // If the local router is not the BSR, this value is 0. The type is
    // interface{} with range: 0..4294967295.
    Pimrpsetexpirytime interface{}
}

func (pimrpsetentry *PIMMIB_Pimrpsettable_Pimrpsetentry) GetEntityData() *types.CommonEntityData {
    pimrpsetentry.EntityData.YFilter = pimrpsetentry.YFilter
    pimrpsetentry.EntityData.YangName = "pimRPSetEntry"
    pimrpsetentry.EntityData.BundleName = "cisco_ios_xe"
    pimrpsetentry.EntityData.ParentYangName = "pimRPSetTable"
    pimrpsetentry.EntityData.SegmentPath = "pimRPSetEntry" + "[pimRPSetComponent='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetcomponent) + "']" + "[pimRPSetGroupAddress='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetgroupaddress) + "']" + "[pimRPSetGroupMask='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetgroupmask) + "']" + "[pimRPSetAddress='" + fmt.Sprintf("%v", pimrpsetentry.Pimrpsetaddress) + "']"
    pimrpsetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimrpsetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimrpsetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimrpsetentry.EntityData.Children = make(map[string]types.YChild)
    pimrpsetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pimrpsetentry.EntityData.Leafs["pimRPSetComponent"] = types.YLeaf{"Pimrpsetcomponent", pimrpsetentry.Pimrpsetcomponent}
    pimrpsetentry.EntityData.Leafs["pimRPSetGroupAddress"] = types.YLeaf{"Pimrpsetgroupaddress", pimrpsetentry.Pimrpsetgroupaddress}
    pimrpsetentry.EntityData.Leafs["pimRPSetGroupMask"] = types.YLeaf{"Pimrpsetgroupmask", pimrpsetentry.Pimrpsetgroupmask}
    pimrpsetentry.EntityData.Leafs["pimRPSetAddress"] = types.YLeaf{"Pimrpsetaddress", pimrpsetentry.Pimrpsetaddress}
    pimrpsetentry.EntityData.Leafs["pimRPSetHoldTime"] = types.YLeaf{"Pimrpsetholdtime", pimrpsetentry.Pimrpsetholdtime}
    pimrpsetentry.EntityData.Leafs["pimRPSetExpiryTime"] = types.YLeaf{"Pimrpsetexpirytime", pimrpsetentry.Pimrpsetexpirytime}
    return &(pimrpsetentry.EntityData)
}

// PIMMIB_Pimipmroutenexthoptable
// The (conceptual) table listing PIM-specific information on
// a subset of the rows of the ipMRouteNextHopTable defined in
// the IP Multicast MIB.
type PIMMIB_Pimipmroutenexthoptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimIpMRouteNextHopTable. There is one
    // entry per entry in the ipMRouteNextHopTable whose interface is running PIM
    // and whose ipMRouteNextHopState is pruned(1). The type is slice of
    // PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry.
    Pimipmroutenexthopentry []PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry
}

func (pimipmroutenexthoptable *PIMMIB_Pimipmroutenexthoptable) GetEntityData() *types.CommonEntityData {
    pimipmroutenexthoptable.EntityData.YFilter = pimipmroutenexthoptable.YFilter
    pimipmroutenexthoptable.EntityData.YangName = "pimIpMRouteNextHopTable"
    pimipmroutenexthoptable.EntityData.BundleName = "cisco_ios_xe"
    pimipmroutenexthoptable.EntityData.ParentYangName = "PIM-MIB"
    pimipmroutenexthoptable.EntityData.SegmentPath = "pimIpMRouteNextHopTable"
    pimipmroutenexthoptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimipmroutenexthoptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimipmroutenexthoptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimipmroutenexthoptable.EntityData.Children = make(map[string]types.YChild)
    pimipmroutenexthoptable.EntityData.Children["pimIpMRouteNextHopEntry"] = types.YChild{"Pimipmroutenexthopentry", nil}
    for i := range pimipmroutenexthoptable.Pimipmroutenexthopentry {
        pimipmroutenexthoptable.EntityData.Children[types.GetSegmentPath(&pimipmroutenexthoptable.Pimipmroutenexthopentry[i])] = types.YChild{"Pimipmroutenexthopentry", &pimipmroutenexthoptable.Pimipmroutenexthopentry[i]}
    }
    pimipmroutenexthoptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pimipmroutenexthoptable.EntityData)
}

// PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry
// An entry (conceptual row) in the pimIpMRouteNextHopTable.
// There is one entry per entry in the ipMRouteNextHopTable
// whose interface is running PIM and whose
// ipMRouteNextHopState is pruned(1).
type PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopgroup
    Ipmroutenexthopgroup interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopsource
    Ipmroutenexthopsource interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopsourcemask
    Ipmroutenexthopsourcemask interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopifindex
    Ipmroutenexthopifindex interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_Ipmroutenexthoptable_Ipmroutenexthopentry_Ipmroutenexthopaddress
    Ipmroutenexthopaddress interface{}

    // This object indicates why the downstream interface was pruned, whether in
    // response to a PIM prune message or due to PIM Assert processing. The type
    // is Pimipmroutenexthopprunereason.
    Pimipmroutenexthopprunereason interface{}
}

func (pimipmroutenexthopentry *PIMMIB_Pimipmroutenexthoptable_Pimipmroutenexthopentry) GetEntityData() *types.CommonEntityData {
    pimipmroutenexthopentry.EntityData.YFilter = pimipmroutenexthopentry.YFilter
    pimipmroutenexthopentry.EntityData.YangName = "pimIpMRouteNextHopEntry"
    pimipmroutenexthopentry.EntityData.BundleName = "cisco_ios_xe"
    pimipmroutenexthopentry.EntityData.ParentYangName = "pimIpMRouteNextHopTable"
    pimipmroutenexthopentry.EntityData.SegmentPath = "pimIpMRouteNextHopEntry" + "[ipMRouteNextHopGroup='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopgroup) + "']" + "[ipMRouteNextHopSource='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopsource) + "']" + "[ipMRouteNextHopSourceMask='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopsourcemask) + "']" + "[ipMRouteNextHopIfIndex='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopifindex) + "']" + "[ipMRouteNextHopAddress='" + fmt.Sprintf("%v", pimipmroutenexthopentry.Ipmroutenexthopaddress) + "']"
    pimipmroutenexthopentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimipmroutenexthopentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimipmroutenexthopentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimipmroutenexthopentry.EntityData.Children = make(map[string]types.YChild)
    pimipmroutenexthopentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pimipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopGroup"] = types.YLeaf{"Ipmroutenexthopgroup", pimipmroutenexthopentry.Ipmroutenexthopgroup}
    pimipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopSource"] = types.YLeaf{"Ipmroutenexthopsource", pimipmroutenexthopentry.Ipmroutenexthopsource}
    pimipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopSourceMask"] = types.YLeaf{"Ipmroutenexthopsourcemask", pimipmroutenexthopentry.Ipmroutenexthopsourcemask}
    pimipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopIfIndex"] = types.YLeaf{"Ipmroutenexthopifindex", pimipmroutenexthopentry.Ipmroutenexthopifindex}
    pimipmroutenexthopentry.EntityData.Leafs["ipMRouteNextHopAddress"] = types.YLeaf{"Ipmroutenexthopaddress", pimipmroutenexthopentry.Ipmroutenexthopaddress}
    pimipmroutenexthopentry.EntityData.Leafs["pimIpMRouteNextHopPruneReason"] = types.YLeaf{"Pimipmroutenexthopprunereason", pimipmroutenexthopentry.Pimipmroutenexthopprunereason}
    return &(pimipmroutenexthopentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimCandidateRPTable. The type is slice of
    // PIMMIB_Pimcandidaterptable_Pimcandidaterpentry.
    Pimcandidaterpentry []PIMMIB_Pimcandidaterptable_Pimcandidaterpentry
}

func (pimcandidaterptable *PIMMIB_Pimcandidaterptable) GetEntityData() *types.CommonEntityData {
    pimcandidaterptable.EntityData.YFilter = pimcandidaterptable.YFilter
    pimcandidaterptable.EntityData.YangName = "pimCandidateRPTable"
    pimcandidaterptable.EntityData.BundleName = "cisco_ios_xe"
    pimcandidaterptable.EntityData.ParentYangName = "PIM-MIB"
    pimcandidaterptable.EntityData.SegmentPath = "pimCandidateRPTable"
    pimcandidaterptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimcandidaterptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimcandidaterptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimcandidaterptable.EntityData.Children = make(map[string]types.YChild)
    pimcandidaterptable.EntityData.Children["pimCandidateRPEntry"] = types.YChild{"Pimcandidaterpentry", nil}
    for i := range pimcandidaterptable.Pimcandidaterpentry {
        pimcandidaterptable.EntityData.Children[types.GetSegmentPath(&pimcandidaterptable.Pimcandidaterpentry[i])] = types.YChild{"Pimcandidaterpentry", &pimcandidaterptable.Pimcandidaterpentry[i]}
    }
    pimcandidaterptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pimcandidaterptable.EntityData)
}

// PIMMIB_Pimcandidaterptable_Pimcandidaterpentry
// An entry (conceptual row) in the pimCandidateRPTable.
type PIMMIB_Pimcandidaterptable_Pimcandidaterpentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address which, when
    // combined with pimCandidateRPGroupMask, identifies a group prefix for which
    // the local router will advertise itself as a Candidate-RP. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Pimcandidaterpgroupaddress interface{}

    // This attribute is a key. The multicast group address mask which, when
    // combined with pimCandidateRPGroupMask, identifies a group prefix for which
    // the local router will advertise itself as a Candidate-RP. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Pimcandidaterpgroupmask interface{}

    // The (unicast) address of the interface which will be      advertised as a
    // Candidate-RP. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Pimcandidaterpaddress interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    Pimcandidaterprowstatus interface{}
}

func (pimcandidaterpentry *PIMMIB_Pimcandidaterptable_Pimcandidaterpentry) GetEntityData() *types.CommonEntityData {
    pimcandidaterpentry.EntityData.YFilter = pimcandidaterpentry.YFilter
    pimcandidaterpentry.EntityData.YangName = "pimCandidateRPEntry"
    pimcandidaterpentry.EntityData.BundleName = "cisco_ios_xe"
    pimcandidaterpentry.EntityData.ParentYangName = "pimCandidateRPTable"
    pimcandidaterpentry.EntityData.SegmentPath = "pimCandidateRPEntry" + "[pimCandidateRPGroupAddress='" + fmt.Sprintf("%v", pimcandidaterpentry.Pimcandidaterpgroupaddress) + "']" + "[pimCandidateRPGroupMask='" + fmt.Sprintf("%v", pimcandidaterpentry.Pimcandidaterpgroupmask) + "']"
    pimcandidaterpentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimcandidaterpentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimcandidaterpentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimcandidaterpentry.EntityData.Children = make(map[string]types.YChild)
    pimcandidaterpentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pimcandidaterpentry.EntityData.Leafs["pimCandidateRPGroupAddress"] = types.YLeaf{"Pimcandidaterpgroupaddress", pimcandidaterpentry.Pimcandidaterpgroupaddress}
    pimcandidaterpentry.EntityData.Leafs["pimCandidateRPGroupMask"] = types.YLeaf{"Pimcandidaterpgroupmask", pimcandidaterpentry.Pimcandidaterpgroupmask}
    pimcandidaterpentry.EntityData.Leafs["pimCandidateRPAddress"] = types.YLeaf{"Pimcandidaterpaddress", pimcandidaterpentry.Pimcandidaterpaddress}
    pimcandidaterpentry.EntityData.Leafs["pimCandidateRPRowStatus"] = types.YLeaf{"Pimcandidaterprowstatus", pimcandidaterpentry.Pimcandidaterprowstatus}
    return &(pimcandidaterpentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimComponentTable. The type is slice of
    // PIMMIB_Pimcomponenttable_Pimcomponententry.
    Pimcomponententry []PIMMIB_Pimcomponenttable_Pimcomponententry
}

func (pimcomponenttable *PIMMIB_Pimcomponenttable) GetEntityData() *types.CommonEntityData {
    pimcomponenttable.EntityData.YFilter = pimcomponenttable.YFilter
    pimcomponenttable.EntityData.YangName = "pimComponentTable"
    pimcomponenttable.EntityData.BundleName = "cisco_ios_xe"
    pimcomponenttable.EntityData.ParentYangName = "PIM-MIB"
    pimcomponenttable.EntityData.SegmentPath = "pimComponentTable"
    pimcomponenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimcomponenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimcomponenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimcomponenttable.EntityData.Children = make(map[string]types.YChild)
    pimcomponenttable.EntityData.Children["pimComponentEntry"] = types.YChild{"Pimcomponententry", nil}
    for i := range pimcomponenttable.Pimcomponententry {
        pimcomponenttable.EntityData.Children[types.GetSegmentPath(&pimcomponenttable.Pimcomponententry[i])] = types.YChild{"Pimcomponententry", &pimcomponenttable.Pimcomponententry[i]}
    }
    pimcomponenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pimcomponenttable.EntityData)
}

// PIMMIB_Pimcomponenttable_Pimcomponententry
// An entry (conceptual row) in the pimComponentTable.
type PIMMIB_Pimcomponenttable_Pimcomponententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A number uniquely identifying the component.  Each
    // protocol instance connected to a separate domain should have a different
    // index value.  Routers that only support membership in a single PIM-SM
    // domain should use a pimComponentIndex value of 1. The type is interface{}
    // with range: 1..255.
    Pimcomponentindex interface{}

    // The IP address of the bootstrap router (BSR) for the local PIM region. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (pimcomponententry *PIMMIB_Pimcomponenttable_Pimcomponententry) GetEntityData() *types.CommonEntityData {
    pimcomponententry.EntityData.YFilter = pimcomponententry.YFilter
    pimcomponententry.EntityData.YangName = "pimComponentEntry"
    pimcomponententry.EntityData.BundleName = "cisco_ios_xe"
    pimcomponententry.EntityData.ParentYangName = "pimComponentTable"
    pimcomponententry.EntityData.SegmentPath = "pimComponentEntry" + "[pimComponentIndex='" + fmt.Sprintf("%v", pimcomponententry.Pimcomponentindex) + "']"
    pimcomponententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimcomponententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimcomponententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimcomponententry.EntityData.Children = make(map[string]types.YChild)
    pimcomponententry.EntityData.Leafs = make(map[string]types.YLeaf)
    pimcomponententry.EntityData.Leafs["pimComponentIndex"] = types.YLeaf{"Pimcomponentindex", pimcomponententry.Pimcomponentindex}
    pimcomponententry.EntityData.Leafs["pimComponentBSRAddress"] = types.YLeaf{"Pimcomponentbsraddress", pimcomponententry.Pimcomponentbsraddress}
    pimcomponententry.EntityData.Leafs["pimComponentBSRExpiryTime"] = types.YLeaf{"Pimcomponentbsrexpirytime", pimcomponententry.Pimcomponentbsrexpirytime}
    pimcomponententry.EntityData.Leafs["pimComponentCRPHoldTime"] = types.YLeaf{"Pimcomponentcrpholdtime", pimcomponententry.Pimcomponentcrpholdtime}
    pimcomponententry.EntityData.Leafs["pimComponentStatus"] = types.YLeaf{"Pimcomponentstatus", pimcomponententry.Pimcomponentstatus}
    return &(pimcomponententry.EntityData)
}

