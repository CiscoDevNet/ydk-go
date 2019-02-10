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
    PimInterfaceTable PIMMIB_PimInterfaceTable

    // The (conceptual) table listing the router's PIM neighbors.
    PimNeighborTable PIMMIB_PimNeighborTable

    // The (conceptual) table listing PIM-specific information on a subset of the
    // rows of the ipMRouteTable defined in the IP Multicast MIB.
    PimIpMRouteTable PIMMIB_PimIpMRouteTable

    // The (conceptual) table listing PIM version 1 information for the Rendezvous
    // Points (RPs) for IP multicast groups. This table is deprecated since its
    // function is replaced by the pimRPSetTable for PIM version 2.
    PimRPTable PIMMIB_PimRPTable

    // The (conceptual) table listing PIM information for candidate Rendezvous
    // Points (RPs) for IP multicast groups. When the local router is the BSR,
    // this information is obtained from received Candidate-RP-Advertisements. 
    // When the local router is not the BSR, this information is obtained from
    // received RP-Set messages.
    PimRPSetTable PIMMIB_PimRPSetTable

    // The (conceptual) table listing PIM-specific information on a subset of the
    // rows of the ipMRouteNextHopTable defined in the IP Multicast MIB.
    PimIpMRouteNextHopTable PIMMIB_PimIpMRouteNextHopTable

    // The (conceptual) table listing the IP multicast groups for which the local
    // router is to advertise itself as a Candidate-RP when the value of
    // pimComponentCRPHoldTime is non-zero.  If this table is empty, then the
    // local router      will advertise itself as a Candidate-RP for all groups
    // (providing the value of pimComponentCRPHoldTime is non- zero).
    PimCandidateRPTable PIMMIB_PimCandidateRPTable

    // The (conceptual) table containing objects specific to a PIM domain.  One
    // row exists for each domain to which the router is connected.  A PIM-SM
    // domain is defined as an area of the network over which Bootstrap messages
    // are forwarded. Typically, a PIM-SM router will be a member of exactly one
    // domain.  This table also supports, however, routers which may form a border
    // between two PIM-SM domains and do not forward Bootstrap messages between
    // them.
    PimComponentTable PIMMIB_PimComponentTable
}

func (pIMMIB *PIMMIB) GetEntityData() *types.CommonEntityData {
    pIMMIB.EntityData.YFilter = pIMMIB.YFilter
    pIMMIB.EntityData.YangName = "PIM-MIB"
    pIMMIB.EntityData.BundleName = "cisco_ios_xe"
    pIMMIB.EntityData.ParentYangName = "PIM-MIB"
    pIMMIB.EntityData.SegmentPath = "PIM-MIB:PIM-MIB"
    pIMMIB.EntityData.AbsolutePath = pIMMIB.EntityData.SegmentPath
    pIMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pIMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pIMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pIMMIB.EntityData.Children = types.NewOrderedMap()
    pIMMIB.EntityData.Children.Append("pim", types.YChild{"Pim", &pIMMIB.Pim})
    pIMMIB.EntityData.Children.Append("pimInterfaceTable", types.YChild{"PimInterfaceTable", &pIMMIB.PimInterfaceTable})
    pIMMIB.EntityData.Children.Append("pimNeighborTable", types.YChild{"PimNeighborTable", &pIMMIB.PimNeighborTable})
    pIMMIB.EntityData.Children.Append("pimIpMRouteTable", types.YChild{"PimIpMRouteTable", &pIMMIB.PimIpMRouteTable})
    pIMMIB.EntityData.Children.Append("pimRPTable", types.YChild{"PimRPTable", &pIMMIB.PimRPTable})
    pIMMIB.EntityData.Children.Append("pimRPSetTable", types.YChild{"PimRPSetTable", &pIMMIB.PimRPSetTable})
    pIMMIB.EntityData.Children.Append("pimIpMRouteNextHopTable", types.YChild{"PimIpMRouteNextHopTable", &pIMMIB.PimIpMRouteNextHopTable})
    pIMMIB.EntityData.Children.Append("pimCandidateRPTable", types.YChild{"PimCandidateRPTable", &pIMMIB.PimCandidateRPTable})
    pIMMIB.EntityData.Children.Append("pimComponentTable", types.YChild{"PimComponentTable", &pIMMIB.PimComponentTable})
    pIMMIB.EntityData.Leafs = types.NewOrderedMap()

    pIMMIB.EntityData.YListKeys = []string {}

    return &(pIMMIB.EntityData)
}

// PIMMIB_Pim
type PIMMIB_Pim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The default interval at which periodic PIM-SM Join/Prune messages are to be
    // sent. The type is interface{} with range: -2147483648..2147483647. Units
    // are seconds.
    PimJoinPruneInterval interface{}
}

func (pim *PIMMIB_Pim) GetEntityData() *types.CommonEntityData {
    pim.EntityData.YFilter = pim.YFilter
    pim.EntityData.YangName = "pim"
    pim.EntityData.BundleName = "cisco_ios_xe"
    pim.EntityData.ParentYangName = "PIM-MIB"
    pim.EntityData.SegmentPath = "pim"
    pim.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pim.EntityData.SegmentPath
    pim.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pim.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pim.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pim.EntityData.Children = types.NewOrderedMap()
    pim.EntityData.Leafs = types.NewOrderedMap()
    pim.EntityData.Leafs.Append("pimJoinPruneInterval", types.YLeaf{"PimJoinPruneInterval", pim.PimJoinPruneInterval})

    pim.EntityData.YListKeys = []string {}

    return &(pim.EntityData)
}

// PIMMIB_PimInterfaceTable
// The (conceptual) table listing the router's PIM interfaces.
// IGMP and PIM are enabled on all interfaces listed in this
// table.
type PIMMIB_PimInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimInterfaceTable. The type is slice of
    // PIMMIB_PimInterfaceTable_PimInterfaceEntry.
    PimInterfaceEntry []*PIMMIB_PimInterfaceTable_PimInterfaceEntry
}

func (pimInterfaceTable *PIMMIB_PimInterfaceTable) GetEntityData() *types.CommonEntityData {
    pimInterfaceTable.EntityData.YFilter = pimInterfaceTable.YFilter
    pimInterfaceTable.EntityData.YangName = "pimInterfaceTable"
    pimInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    pimInterfaceTable.EntityData.ParentYangName = "PIM-MIB"
    pimInterfaceTable.EntityData.SegmentPath = "pimInterfaceTable"
    pimInterfaceTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimInterfaceTable.EntityData.SegmentPath
    pimInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimInterfaceTable.EntityData.Children = types.NewOrderedMap()
    pimInterfaceTable.EntityData.Children.Append("pimInterfaceEntry", types.YChild{"PimInterfaceEntry", nil})
    for i := range pimInterfaceTable.PimInterfaceEntry {
        pimInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(pimInterfaceTable.PimInterfaceEntry[i]), types.YChild{"PimInterfaceEntry", pimInterfaceTable.PimInterfaceEntry[i]})
    }
    pimInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    pimInterfaceTable.EntityData.YListKeys = []string {}

    return &(pimInterfaceTable.EntityData)
}

// PIMMIB_PimInterfaceTable_PimInterfaceEntry
// An entry (conceptual row) in the pimInterfaceTable.
type PIMMIB_PimInterfaceTable_PimInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The ifIndex value of this PIM interface. The type
    // is interface{} with range: 1..2147483647.
    PimInterfaceIfIndex interface{}

    // The IP address of the PIM interface. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimInterfaceAddress interface{}

    // The network mask for the IP address of the PIM interface. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimInterfaceNetMask interface{}

    // The configured mode of this PIM interface.  A value of sparseDense is only
    // valid for PIMv1. The type is PimInterfaceMode.
    PimInterfaceMode interface{}

    // The Designated Router on this PIM interface.  For point-to- point
    // interfaces, this object has the value 0.0.0.0. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimInterfaceDR interface{}

    // The frequency at which PIM Hello messages are transmitted on this
    // interface. The type is interface{} with range: -2147483648..2147483647.
    // Units are seconds.
    PimInterfaceHelloInterval interface{}

    // The status of this entry.  Creating the entry enables PIM on the interface;
    // destroying the entry disables PIM on the interface. The type is RowStatus.
    PimInterfaceStatus interface{}

    // The frequency at which PIM Join/Prune messages are transmitted on this PIM
    // interface.  The default value of this object is the pimJoinPruneInterval.
    // The type is interface{} with range: -2147483648..2147483647. Units are
    // seconds.
    PimInterfaceJoinPruneInterval interface{}

    // The preference value for the local interface as a candidate bootstrap
    // router.  The value of -1 is used to indicate that the local interface is
    // not a candidate BSR interface. The type is interface{} with range: -1..255.
    PimInterfaceCBSRPreference interface{}
}

func (pimInterfaceEntry *PIMMIB_PimInterfaceTable_PimInterfaceEntry) GetEntityData() *types.CommonEntityData {
    pimInterfaceEntry.EntityData.YFilter = pimInterfaceEntry.YFilter
    pimInterfaceEntry.EntityData.YangName = "pimInterfaceEntry"
    pimInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    pimInterfaceEntry.EntityData.ParentYangName = "pimInterfaceTable"
    pimInterfaceEntry.EntityData.SegmentPath = "pimInterfaceEntry" + types.AddKeyToken(pimInterfaceEntry.PimInterfaceIfIndex, "pimInterfaceIfIndex")
    pimInterfaceEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimInterfaceTable/" + pimInterfaceEntry.EntityData.SegmentPath
    pimInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    pimInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceIfIndex", types.YLeaf{"PimInterfaceIfIndex", pimInterfaceEntry.PimInterfaceIfIndex})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceAddress", types.YLeaf{"PimInterfaceAddress", pimInterfaceEntry.PimInterfaceAddress})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceNetMask", types.YLeaf{"PimInterfaceNetMask", pimInterfaceEntry.PimInterfaceNetMask})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceMode", types.YLeaf{"PimInterfaceMode", pimInterfaceEntry.PimInterfaceMode})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceDR", types.YLeaf{"PimInterfaceDR", pimInterfaceEntry.PimInterfaceDR})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceHelloInterval", types.YLeaf{"PimInterfaceHelloInterval", pimInterfaceEntry.PimInterfaceHelloInterval})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceStatus", types.YLeaf{"PimInterfaceStatus", pimInterfaceEntry.PimInterfaceStatus})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceJoinPruneInterval", types.YLeaf{"PimInterfaceJoinPruneInterval", pimInterfaceEntry.PimInterfaceJoinPruneInterval})
    pimInterfaceEntry.EntityData.Leafs.Append("pimInterfaceCBSRPreference", types.YLeaf{"PimInterfaceCBSRPreference", pimInterfaceEntry.PimInterfaceCBSRPreference})

    pimInterfaceEntry.EntityData.YListKeys = []string {"PimInterfaceIfIndex"}

    return &(pimInterfaceEntry.EntityData)
}

// PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode represents sparseDense is only valid for PIMv1.
type PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode string

const (
    PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode_dense PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode = "dense"

    PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode_sparse PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode = "sparse"

    PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode_sparseDense PIMMIB_PimInterfaceTable_PimInterfaceEntry_PimInterfaceMode = "sparseDense"
)

// PIMMIB_PimNeighborTable
// The (conceptual) table listing the router's PIM neighbors.
type PIMMIB_PimNeighborTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimNeighborTable. The type is slice of
    // PIMMIB_PimNeighborTable_PimNeighborEntry.
    PimNeighborEntry []*PIMMIB_PimNeighborTable_PimNeighborEntry
}

func (pimNeighborTable *PIMMIB_PimNeighborTable) GetEntityData() *types.CommonEntityData {
    pimNeighborTable.EntityData.YFilter = pimNeighborTable.YFilter
    pimNeighborTable.EntityData.YangName = "pimNeighborTable"
    pimNeighborTable.EntityData.BundleName = "cisco_ios_xe"
    pimNeighborTable.EntityData.ParentYangName = "PIM-MIB"
    pimNeighborTable.EntityData.SegmentPath = "pimNeighborTable"
    pimNeighborTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimNeighborTable.EntityData.SegmentPath
    pimNeighborTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimNeighborTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimNeighborTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimNeighborTable.EntityData.Children = types.NewOrderedMap()
    pimNeighborTable.EntityData.Children.Append("pimNeighborEntry", types.YChild{"PimNeighborEntry", nil})
    for i := range pimNeighborTable.PimNeighborEntry {
        pimNeighborTable.EntityData.Children.Append(types.GetSegmentPath(pimNeighborTable.PimNeighborEntry[i]), types.YChild{"PimNeighborEntry", pimNeighborTable.PimNeighborEntry[i]})
    }
    pimNeighborTable.EntityData.Leafs = types.NewOrderedMap()

    pimNeighborTable.EntityData.YListKeys = []string {}

    return &(pimNeighborTable.EntityData)
}

// PIMMIB_PimNeighborTable_PimNeighborEntry
// An entry (conceptual row) in the pimNeighborTable.
type PIMMIB_PimNeighborTable_PimNeighborEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP address of the PIM neighbor for which this
    // entry contains information. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimNeighborAddress interface{}

    // The value of ifIndex for the interface used to reach this PIM neighbor. The
    // type is interface{} with range: 1..2147483647.
    PimNeighborIfIndex interface{}

    // The time since this PIM neighbor (last) became a neighbor of the local
    // router. The type is interface{} with range: 0..4294967295.
    PimNeighborUpTime interface{}

    // The minimum time remaining before this PIM neighbor will be aged out. The
    // type is interface{} with range: 0..4294967295.
    PimNeighborExpiryTime interface{}

    // The active PIM mode of this neighbor.  This object is deprecated for PIMv2
    // routers since all neighbors on the interface must be either dense or sparse
    // as determined by the protocol running on the interface. The type is
    // PimNeighborMode.
    PimNeighborMode interface{}
}

func (pimNeighborEntry *PIMMIB_PimNeighborTable_PimNeighborEntry) GetEntityData() *types.CommonEntityData {
    pimNeighborEntry.EntityData.YFilter = pimNeighborEntry.YFilter
    pimNeighborEntry.EntityData.YangName = "pimNeighborEntry"
    pimNeighborEntry.EntityData.BundleName = "cisco_ios_xe"
    pimNeighborEntry.EntityData.ParentYangName = "pimNeighborTable"
    pimNeighborEntry.EntityData.SegmentPath = "pimNeighborEntry" + types.AddKeyToken(pimNeighborEntry.PimNeighborAddress, "pimNeighborAddress")
    pimNeighborEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimNeighborTable/" + pimNeighborEntry.EntityData.SegmentPath
    pimNeighborEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimNeighborEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimNeighborEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimNeighborEntry.EntityData.Children = types.NewOrderedMap()
    pimNeighborEntry.EntityData.Leafs = types.NewOrderedMap()
    pimNeighborEntry.EntityData.Leafs.Append("pimNeighborAddress", types.YLeaf{"PimNeighborAddress", pimNeighborEntry.PimNeighborAddress})
    pimNeighborEntry.EntityData.Leafs.Append("pimNeighborIfIndex", types.YLeaf{"PimNeighborIfIndex", pimNeighborEntry.PimNeighborIfIndex})
    pimNeighborEntry.EntityData.Leafs.Append("pimNeighborUpTime", types.YLeaf{"PimNeighborUpTime", pimNeighborEntry.PimNeighborUpTime})
    pimNeighborEntry.EntityData.Leafs.Append("pimNeighborExpiryTime", types.YLeaf{"PimNeighborExpiryTime", pimNeighborEntry.PimNeighborExpiryTime})
    pimNeighborEntry.EntityData.Leafs.Append("pimNeighborMode", types.YLeaf{"PimNeighborMode", pimNeighborEntry.PimNeighborMode})

    pimNeighborEntry.EntityData.YListKeys = []string {"PimNeighborAddress"}

    return &(pimNeighborEntry.EntityData)
}

// PIMMIB_PimNeighborTable_PimNeighborEntry_PimNeighborMode represents the protocol running on the interface.
type PIMMIB_PimNeighborTable_PimNeighborEntry_PimNeighborMode string

const (
    PIMMIB_PimNeighborTable_PimNeighborEntry_PimNeighborMode_dense PIMMIB_PimNeighborTable_PimNeighborEntry_PimNeighborMode = "dense"

    PIMMIB_PimNeighborTable_PimNeighborEntry_PimNeighborMode_sparse PIMMIB_PimNeighborTable_PimNeighborEntry_PimNeighborMode = "sparse"
)

// PIMMIB_PimIpMRouteTable
// The (conceptual) table listing PIM-specific information on
// a subset of the rows of the ipMRouteTable defined in the IP
// Multicast MIB.
type PIMMIB_PimIpMRouteTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimIpMRouteTable.  There is one entry per
    // entry in the ipMRouteTable whose incoming interface is running PIM. The
    // type is slice of PIMMIB_PimIpMRouteTable_PimIpMRouteEntry.
    PimIpMRouteEntry []*PIMMIB_PimIpMRouteTable_PimIpMRouteEntry
}

func (pimIpMRouteTable *PIMMIB_PimIpMRouteTable) GetEntityData() *types.CommonEntityData {
    pimIpMRouteTable.EntityData.YFilter = pimIpMRouteTable.YFilter
    pimIpMRouteTable.EntityData.YangName = "pimIpMRouteTable"
    pimIpMRouteTable.EntityData.BundleName = "cisco_ios_xe"
    pimIpMRouteTable.EntityData.ParentYangName = "PIM-MIB"
    pimIpMRouteTable.EntityData.SegmentPath = "pimIpMRouteTable"
    pimIpMRouteTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimIpMRouteTable.EntityData.SegmentPath
    pimIpMRouteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimIpMRouteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimIpMRouteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimIpMRouteTable.EntityData.Children = types.NewOrderedMap()
    pimIpMRouteTable.EntityData.Children.Append("pimIpMRouteEntry", types.YChild{"PimIpMRouteEntry", nil})
    for i := range pimIpMRouteTable.PimIpMRouteEntry {
        pimIpMRouteTable.EntityData.Children.Append(types.GetSegmentPath(pimIpMRouteTable.PimIpMRouteEntry[i]), types.YChild{"PimIpMRouteEntry", pimIpMRouteTable.PimIpMRouteEntry[i]})
    }
    pimIpMRouteTable.EntityData.Leafs = types.NewOrderedMap()

    pimIpMRouteTable.EntityData.YListKeys = []string {}

    return &(pimIpMRouteTable.EntityData)
}

// PIMMIB_PimIpMRouteTable_PimIpMRouteEntry
// An entry (conceptual row) in the pimIpMRouteTable.  There
// is one entry per entry in the ipMRouteTable whose incoming
// interface is running PIM.
type PIMMIB_PimIpMRouteTable_PimIpMRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteGroup
    IpMRouteGroup interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteSource
    IpMRouteSource interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteTable_IpMRouteEntry_IpMRouteSourceMask
    IpMRouteSourceMask interface{}

    // The time remaining before the router changes its upstream neighbor back to
    // its RPF neighbor.  This timer is called the Assert timer in the PIM Sparse
    // and Dense mode specification.      A value of 0 indicates that no Assert
    // has changed the upstream neighbor away from the RPF neighbor. The type is
    // interface{} with range: 0..4294967295.
    PimIpMRouteUpstreamAssertTimer interface{}

    // The metric advertised by the assert winner on the upstream interface, or 0
    // if no such assert is in received. The type is interface{} with range:
    // -2147483648..2147483647.
    PimIpMRouteAssertMetric interface{}

    // The preference advertised by the assert winner on the upstream interface,
    // or 0 if no such assert is in effect. The type is interface{} with range:
    // -2147483648..2147483647.
    PimIpMRouteAssertMetricPref interface{}

    // The value of the RPT-bit advertised by the assert winner on the upstream
    // interface, or false if no such assert is in effect. The type is bool.
    PimIpMRouteAssertRPTBit interface{}

    // This object describes PIM-specific flags related to a multicast state
    // entry.  See the PIM Sparse Mode specification for the meaning of the RPT
    // and SPT bits. The type is string with length: 1.
    PimIpMRouteFlags interface{}
}

func (pimIpMRouteEntry *PIMMIB_PimIpMRouteTable_PimIpMRouteEntry) GetEntityData() *types.CommonEntityData {
    pimIpMRouteEntry.EntityData.YFilter = pimIpMRouteEntry.YFilter
    pimIpMRouteEntry.EntityData.YangName = "pimIpMRouteEntry"
    pimIpMRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    pimIpMRouteEntry.EntityData.ParentYangName = "pimIpMRouteTable"
    pimIpMRouteEntry.EntityData.SegmentPath = "pimIpMRouteEntry" + types.AddKeyToken(pimIpMRouteEntry.IpMRouteGroup, "ipMRouteGroup") + types.AddKeyToken(pimIpMRouteEntry.IpMRouteSource, "ipMRouteSource") + types.AddKeyToken(pimIpMRouteEntry.IpMRouteSourceMask, "ipMRouteSourceMask")
    pimIpMRouteEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimIpMRouteTable/" + pimIpMRouteEntry.EntityData.SegmentPath
    pimIpMRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimIpMRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimIpMRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimIpMRouteEntry.EntityData.Children = types.NewOrderedMap()
    pimIpMRouteEntry.EntityData.Leafs = types.NewOrderedMap()
    pimIpMRouteEntry.EntityData.Leafs.Append("ipMRouteGroup", types.YLeaf{"IpMRouteGroup", pimIpMRouteEntry.IpMRouteGroup})
    pimIpMRouteEntry.EntityData.Leafs.Append("ipMRouteSource", types.YLeaf{"IpMRouteSource", pimIpMRouteEntry.IpMRouteSource})
    pimIpMRouteEntry.EntityData.Leafs.Append("ipMRouteSourceMask", types.YLeaf{"IpMRouteSourceMask", pimIpMRouteEntry.IpMRouteSourceMask})
    pimIpMRouteEntry.EntityData.Leafs.Append("pimIpMRouteUpstreamAssertTimer", types.YLeaf{"PimIpMRouteUpstreamAssertTimer", pimIpMRouteEntry.PimIpMRouteUpstreamAssertTimer})
    pimIpMRouteEntry.EntityData.Leafs.Append("pimIpMRouteAssertMetric", types.YLeaf{"PimIpMRouteAssertMetric", pimIpMRouteEntry.PimIpMRouteAssertMetric})
    pimIpMRouteEntry.EntityData.Leafs.Append("pimIpMRouteAssertMetricPref", types.YLeaf{"PimIpMRouteAssertMetricPref", pimIpMRouteEntry.PimIpMRouteAssertMetricPref})
    pimIpMRouteEntry.EntityData.Leafs.Append("pimIpMRouteAssertRPTBit", types.YLeaf{"PimIpMRouteAssertRPTBit", pimIpMRouteEntry.PimIpMRouteAssertRPTBit})
    pimIpMRouteEntry.EntityData.Leafs.Append("pimIpMRouteFlags", types.YLeaf{"PimIpMRouteFlags", pimIpMRouteEntry.PimIpMRouteFlags})

    pimIpMRouteEntry.EntityData.YListKeys = []string {"IpMRouteGroup", "IpMRouteSource", "IpMRouteSourceMask"}

    return &(pimIpMRouteEntry.EntityData)
}

// PIMMIB_PimRPTable
// The (conceptual) table listing PIM version 1 information
// for the Rendezvous Points (RPs) for IP multicast groups.
// This table is deprecated since its function is replaced by
// the pimRPSetTable for PIM version 2.
type PIMMIB_PimRPTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimRPTable.  There is one entry per RP
    // address for each IP multicast group. The type is slice of
    // PIMMIB_PimRPTable_PimRPEntry.
    PimRPEntry []*PIMMIB_PimRPTable_PimRPEntry
}

func (pimRPTable *PIMMIB_PimRPTable) GetEntityData() *types.CommonEntityData {
    pimRPTable.EntityData.YFilter = pimRPTable.YFilter
    pimRPTable.EntityData.YangName = "pimRPTable"
    pimRPTable.EntityData.BundleName = "cisco_ios_xe"
    pimRPTable.EntityData.ParentYangName = "PIM-MIB"
    pimRPTable.EntityData.SegmentPath = "pimRPTable"
    pimRPTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimRPTable.EntityData.SegmentPath
    pimRPTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimRPTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimRPTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimRPTable.EntityData.Children = types.NewOrderedMap()
    pimRPTable.EntityData.Children.Append("pimRPEntry", types.YChild{"PimRPEntry", nil})
    for i := range pimRPTable.PimRPEntry {
        pimRPTable.EntityData.Children.Append(types.GetSegmentPath(pimRPTable.PimRPEntry[i]), types.YChild{"PimRPEntry", pimRPTable.PimRPEntry[i]})
    }
    pimRPTable.EntityData.Leafs = types.NewOrderedMap()

    pimRPTable.EntityData.YListKeys = []string {}

    return &(pimRPTable.EntityData)
}

// PIMMIB_PimRPTable_PimRPEntry
// An entry (conceptual row) in the pimRPTable.  There is one
// entry per RP address for each IP multicast group.
type PIMMIB_PimRPTable_PimRPEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP multicast group address for which this
    // entry contains information about an RP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimRPGroupAddress interface{}

    // This attribute is a key. The unicast address of the RP. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimRPAddress interface{}

    // The state of the RP. The type is PimRPState.
    PimRPState interface{}

    // The minimum time remaining before the next state change. When pimRPState is
    // up, this is the minimum time which must expire until it can be declared
    // down.  When pimRPState is down, this is the time until it will be declared
    // up (in order to retry). The type is interface{} with range: 0..4294967295.
    PimRPStateTimer interface{}

    // The value of sysUpTime at the time when the corresponding instance of
    // pimRPState last changed its value. The type is interface{} with range:
    // 0..4294967295.
    PimRPLastChange interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    PimRPRowStatus interface{}
}

func (pimRPEntry *PIMMIB_PimRPTable_PimRPEntry) GetEntityData() *types.CommonEntityData {
    pimRPEntry.EntityData.YFilter = pimRPEntry.YFilter
    pimRPEntry.EntityData.YangName = "pimRPEntry"
    pimRPEntry.EntityData.BundleName = "cisco_ios_xe"
    pimRPEntry.EntityData.ParentYangName = "pimRPTable"
    pimRPEntry.EntityData.SegmentPath = "pimRPEntry" + types.AddKeyToken(pimRPEntry.PimRPGroupAddress, "pimRPGroupAddress") + types.AddKeyToken(pimRPEntry.PimRPAddress, "pimRPAddress")
    pimRPEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimRPTable/" + pimRPEntry.EntityData.SegmentPath
    pimRPEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimRPEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimRPEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimRPEntry.EntityData.Children = types.NewOrderedMap()
    pimRPEntry.EntityData.Leafs = types.NewOrderedMap()
    pimRPEntry.EntityData.Leafs.Append("pimRPGroupAddress", types.YLeaf{"PimRPGroupAddress", pimRPEntry.PimRPGroupAddress})
    pimRPEntry.EntityData.Leafs.Append("pimRPAddress", types.YLeaf{"PimRPAddress", pimRPEntry.PimRPAddress})
    pimRPEntry.EntityData.Leafs.Append("pimRPState", types.YLeaf{"PimRPState", pimRPEntry.PimRPState})
    pimRPEntry.EntityData.Leafs.Append("pimRPStateTimer", types.YLeaf{"PimRPStateTimer", pimRPEntry.PimRPStateTimer})
    pimRPEntry.EntityData.Leafs.Append("pimRPLastChange", types.YLeaf{"PimRPLastChange", pimRPEntry.PimRPLastChange})
    pimRPEntry.EntityData.Leafs.Append("pimRPRowStatus", types.YLeaf{"PimRPRowStatus", pimRPEntry.PimRPRowStatus})

    pimRPEntry.EntityData.YListKeys = []string {"PimRPGroupAddress", "PimRPAddress"}

    return &(pimRPEntry.EntityData)
}

// PIMMIB_PimRPTable_PimRPEntry_PimRPState represents The state of the RP.
type PIMMIB_PimRPTable_PimRPEntry_PimRPState string

const (
    PIMMIB_PimRPTable_PimRPEntry_PimRPState_up PIMMIB_PimRPTable_PimRPEntry_PimRPState = "up"

    PIMMIB_PimRPTable_PimRPEntry_PimRPState_down PIMMIB_PimRPTable_PimRPEntry_PimRPState = "down"
)

// PIMMIB_PimRPSetTable
// The (conceptual) table listing PIM information for
// candidate Rendezvous Points (RPs) for IP multicast groups.
// When the local router is the BSR, this information is
// obtained from received Candidate-RP-Advertisements.  When
// the local router is not the BSR, this information is
// obtained from received RP-Set messages.
type PIMMIB_PimRPSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimRPSetTable. The type is slice of
    // PIMMIB_PimRPSetTable_PimRPSetEntry.
    PimRPSetEntry []*PIMMIB_PimRPSetTable_PimRPSetEntry
}

func (pimRPSetTable *PIMMIB_PimRPSetTable) GetEntityData() *types.CommonEntityData {
    pimRPSetTable.EntityData.YFilter = pimRPSetTable.YFilter
    pimRPSetTable.EntityData.YangName = "pimRPSetTable"
    pimRPSetTable.EntityData.BundleName = "cisco_ios_xe"
    pimRPSetTable.EntityData.ParentYangName = "PIM-MIB"
    pimRPSetTable.EntityData.SegmentPath = "pimRPSetTable"
    pimRPSetTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimRPSetTable.EntityData.SegmentPath
    pimRPSetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimRPSetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimRPSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimRPSetTable.EntityData.Children = types.NewOrderedMap()
    pimRPSetTable.EntityData.Children.Append("pimRPSetEntry", types.YChild{"PimRPSetEntry", nil})
    for i := range pimRPSetTable.PimRPSetEntry {
        pimRPSetTable.EntityData.Children.Append(types.GetSegmentPath(pimRPSetTable.PimRPSetEntry[i]), types.YChild{"PimRPSetEntry", pimRPSetTable.PimRPSetEntry[i]})
    }
    pimRPSetTable.EntityData.Leafs = types.NewOrderedMap()

    pimRPSetTable.EntityData.YListKeys = []string {}

    return &(pimRPSetTable.EntityData)
}

// PIMMIB_PimRPSetTable_PimRPSetEntry
// An entry (conceptual row) in the pimRPSetTable.
type PIMMIB_PimRPSetTable_PimRPSetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key.  A number uniquely identifying the component. 
    // Each protocol instance connected to a separate domain should have a
    // different index value. The type is interface{} with range: 1..255.
    PimRPSetComponent interface{}

    // This attribute is a key. The IP multicast group address which, when
    // combined with pimRPSetGroupMask, gives the group prefix for which this
    // entry contains information about the Candidate-RP. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimRPSetGroupAddress interface{}

    // This attribute is a key. The multicast group address mask which, when
    // combined with pimRPSetGroupAddress, gives the group prefix for which this
    // entry contains information about the Candidate-RP. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimRPSetGroupMask interface{}

    // This attribute is a key. The IP address of the Candidate-RP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimRPSetAddress interface{}

    // The holdtime of a Candidate-RP.  If the local router is not the BSR, this
    // value is 0. The type is interface{} with range: 0..255. Units are seconds.
    PimRPSetHoldTime interface{}

    // The minimum time remaining before the Candidate-RP will be declared down. 
    // If the local router is not the BSR, this value is 0. The type is
    // interface{} with range: 0..4294967295.
    PimRPSetExpiryTime interface{}
}

func (pimRPSetEntry *PIMMIB_PimRPSetTable_PimRPSetEntry) GetEntityData() *types.CommonEntityData {
    pimRPSetEntry.EntityData.YFilter = pimRPSetEntry.YFilter
    pimRPSetEntry.EntityData.YangName = "pimRPSetEntry"
    pimRPSetEntry.EntityData.BundleName = "cisco_ios_xe"
    pimRPSetEntry.EntityData.ParentYangName = "pimRPSetTable"
    pimRPSetEntry.EntityData.SegmentPath = "pimRPSetEntry" + types.AddKeyToken(pimRPSetEntry.PimRPSetComponent, "pimRPSetComponent") + types.AddKeyToken(pimRPSetEntry.PimRPSetGroupAddress, "pimRPSetGroupAddress") + types.AddKeyToken(pimRPSetEntry.PimRPSetGroupMask, "pimRPSetGroupMask") + types.AddKeyToken(pimRPSetEntry.PimRPSetAddress, "pimRPSetAddress")
    pimRPSetEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimRPSetTable/" + pimRPSetEntry.EntityData.SegmentPath
    pimRPSetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimRPSetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimRPSetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimRPSetEntry.EntityData.Children = types.NewOrderedMap()
    pimRPSetEntry.EntityData.Leafs = types.NewOrderedMap()
    pimRPSetEntry.EntityData.Leafs.Append("pimRPSetComponent", types.YLeaf{"PimRPSetComponent", pimRPSetEntry.PimRPSetComponent})
    pimRPSetEntry.EntityData.Leafs.Append("pimRPSetGroupAddress", types.YLeaf{"PimRPSetGroupAddress", pimRPSetEntry.PimRPSetGroupAddress})
    pimRPSetEntry.EntityData.Leafs.Append("pimRPSetGroupMask", types.YLeaf{"PimRPSetGroupMask", pimRPSetEntry.PimRPSetGroupMask})
    pimRPSetEntry.EntityData.Leafs.Append("pimRPSetAddress", types.YLeaf{"PimRPSetAddress", pimRPSetEntry.PimRPSetAddress})
    pimRPSetEntry.EntityData.Leafs.Append("pimRPSetHoldTime", types.YLeaf{"PimRPSetHoldTime", pimRPSetEntry.PimRPSetHoldTime})
    pimRPSetEntry.EntityData.Leafs.Append("pimRPSetExpiryTime", types.YLeaf{"PimRPSetExpiryTime", pimRPSetEntry.PimRPSetExpiryTime})

    pimRPSetEntry.EntityData.YListKeys = []string {"PimRPSetComponent", "PimRPSetGroupAddress", "PimRPSetGroupMask", "PimRPSetAddress"}

    return &(pimRPSetEntry.EntityData)
}

// PIMMIB_PimIpMRouteNextHopTable
// The (conceptual) table listing PIM-specific information on
// a subset of the rows of the ipMRouteNextHopTable defined in
// the IP Multicast MIB.
type PIMMIB_PimIpMRouteNextHopTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimIpMRouteNextHopTable. There is one
    // entry per entry in the ipMRouteNextHopTable whose interface is running PIM
    // and whose ipMRouteNextHopState is pruned(1). The type is slice of
    // PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry.
    PimIpMRouteNextHopEntry []*PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry
}

func (pimIpMRouteNextHopTable *PIMMIB_PimIpMRouteNextHopTable) GetEntityData() *types.CommonEntityData {
    pimIpMRouteNextHopTable.EntityData.YFilter = pimIpMRouteNextHopTable.YFilter
    pimIpMRouteNextHopTable.EntityData.YangName = "pimIpMRouteNextHopTable"
    pimIpMRouteNextHopTable.EntityData.BundleName = "cisco_ios_xe"
    pimIpMRouteNextHopTable.EntityData.ParentYangName = "PIM-MIB"
    pimIpMRouteNextHopTable.EntityData.SegmentPath = "pimIpMRouteNextHopTable"
    pimIpMRouteNextHopTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimIpMRouteNextHopTable.EntityData.SegmentPath
    pimIpMRouteNextHopTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimIpMRouteNextHopTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimIpMRouteNextHopTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimIpMRouteNextHopTable.EntityData.Children = types.NewOrderedMap()
    pimIpMRouteNextHopTable.EntityData.Children.Append("pimIpMRouteNextHopEntry", types.YChild{"PimIpMRouteNextHopEntry", nil})
    for i := range pimIpMRouteNextHopTable.PimIpMRouteNextHopEntry {
        pimIpMRouteNextHopTable.EntityData.Children.Append(types.GetSegmentPath(pimIpMRouteNextHopTable.PimIpMRouteNextHopEntry[i]), types.YChild{"PimIpMRouteNextHopEntry", pimIpMRouteNextHopTable.PimIpMRouteNextHopEntry[i]})
    }
    pimIpMRouteNextHopTable.EntityData.Leafs = types.NewOrderedMap()

    pimIpMRouteNextHopTable.EntityData.YListKeys = []string {}

    return &(pimIpMRouteNextHopTable.EntityData)
}

// PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry
// An entry (conceptual row) in the pimIpMRouteNextHopTable.
// There is one entry per entry in the ipMRouteNextHopTable
// whose interface is running PIM and whose
// ipMRouteNextHopState is pruned(1).
type PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopGroup
    IpMRouteNextHopGroup interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopSource
    IpMRouteNextHopSource interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopSourceMask
    IpMRouteNextHopSourceMask interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopIfIndex
    IpMRouteNextHopIfIndex interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // ipmroute_std_mib.IPMROUTESTDMIB_IpMRouteNextHopTable_IpMRouteNextHopEntry_IpMRouteNextHopAddress
    IpMRouteNextHopAddress interface{}

    // This object indicates why the downstream interface was pruned, whether in
    // response to a PIM prune message or due to PIM Assert processing. The type
    // is PimIpMRouteNextHopPruneReason.
    PimIpMRouteNextHopPruneReason interface{}
}

func (pimIpMRouteNextHopEntry *PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry) GetEntityData() *types.CommonEntityData {
    pimIpMRouteNextHopEntry.EntityData.YFilter = pimIpMRouteNextHopEntry.YFilter
    pimIpMRouteNextHopEntry.EntityData.YangName = "pimIpMRouteNextHopEntry"
    pimIpMRouteNextHopEntry.EntityData.BundleName = "cisco_ios_xe"
    pimIpMRouteNextHopEntry.EntityData.ParentYangName = "pimIpMRouteNextHopTable"
    pimIpMRouteNextHopEntry.EntityData.SegmentPath = "pimIpMRouteNextHopEntry" + types.AddKeyToken(pimIpMRouteNextHopEntry.IpMRouteNextHopGroup, "ipMRouteNextHopGroup") + types.AddKeyToken(pimIpMRouteNextHopEntry.IpMRouteNextHopSource, "ipMRouteNextHopSource") + types.AddKeyToken(pimIpMRouteNextHopEntry.IpMRouteNextHopSourceMask, "ipMRouteNextHopSourceMask") + types.AddKeyToken(pimIpMRouteNextHopEntry.IpMRouteNextHopIfIndex, "ipMRouteNextHopIfIndex") + types.AddKeyToken(pimIpMRouteNextHopEntry.IpMRouteNextHopAddress, "ipMRouteNextHopAddress")
    pimIpMRouteNextHopEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimIpMRouteNextHopTable/" + pimIpMRouteNextHopEntry.EntityData.SegmentPath
    pimIpMRouteNextHopEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimIpMRouteNextHopEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimIpMRouteNextHopEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimIpMRouteNextHopEntry.EntityData.Children = types.NewOrderedMap()
    pimIpMRouteNextHopEntry.EntityData.Leafs = types.NewOrderedMap()
    pimIpMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopGroup", types.YLeaf{"IpMRouteNextHopGroup", pimIpMRouteNextHopEntry.IpMRouteNextHopGroup})
    pimIpMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopSource", types.YLeaf{"IpMRouteNextHopSource", pimIpMRouteNextHopEntry.IpMRouteNextHopSource})
    pimIpMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopSourceMask", types.YLeaf{"IpMRouteNextHopSourceMask", pimIpMRouteNextHopEntry.IpMRouteNextHopSourceMask})
    pimIpMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopIfIndex", types.YLeaf{"IpMRouteNextHopIfIndex", pimIpMRouteNextHopEntry.IpMRouteNextHopIfIndex})
    pimIpMRouteNextHopEntry.EntityData.Leafs.Append("ipMRouteNextHopAddress", types.YLeaf{"IpMRouteNextHopAddress", pimIpMRouteNextHopEntry.IpMRouteNextHopAddress})
    pimIpMRouteNextHopEntry.EntityData.Leafs.Append("pimIpMRouteNextHopPruneReason", types.YLeaf{"PimIpMRouteNextHopPruneReason", pimIpMRouteNextHopEntry.PimIpMRouteNextHopPruneReason})

    pimIpMRouteNextHopEntry.EntityData.YListKeys = []string {"IpMRouteNextHopGroup", "IpMRouteNextHopSource", "IpMRouteNextHopSourceMask", "IpMRouteNextHopIfIndex", "IpMRouteNextHopAddress"}

    return &(pimIpMRouteNextHopEntry.EntityData)
}

// PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason represents PIM Assert processing.
type PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason string

const (
    PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason_other PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason = "other"

    PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason_prune PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason = "prune"

    PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason_assert PIMMIB_PimIpMRouteNextHopTable_PimIpMRouteNextHopEntry_PimIpMRouteNextHopPruneReason = "assert"
)

// PIMMIB_PimCandidateRPTable
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
type PIMMIB_PimCandidateRPTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimCandidateRPTable. The type is slice of
    // PIMMIB_PimCandidateRPTable_PimCandidateRPEntry.
    PimCandidateRPEntry []*PIMMIB_PimCandidateRPTable_PimCandidateRPEntry
}

func (pimCandidateRPTable *PIMMIB_PimCandidateRPTable) GetEntityData() *types.CommonEntityData {
    pimCandidateRPTable.EntityData.YFilter = pimCandidateRPTable.YFilter
    pimCandidateRPTable.EntityData.YangName = "pimCandidateRPTable"
    pimCandidateRPTable.EntityData.BundleName = "cisco_ios_xe"
    pimCandidateRPTable.EntityData.ParentYangName = "PIM-MIB"
    pimCandidateRPTable.EntityData.SegmentPath = "pimCandidateRPTable"
    pimCandidateRPTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimCandidateRPTable.EntityData.SegmentPath
    pimCandidateRPTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimCandidateRPTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimCandidateRPTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimCandidateRPTable.EntityData.Children = types.NewOrderedMap()
    pimCandidateRPTable.EntityData.Children.Append("pimCandidateRPEntry", types.YChild{"PimCandidateRPEntry", nil})
    for i := range pimCandidateRPTable.PimCandidateRPEntry {
        pimCandidateRPTable.EntityData.Children.Append(types.GetSegmentPath(pimCandidateRPTable.PimCandidateRPEntry[i]), types.YChild{"PimCandidateRPEntry", pimCandidateRPTable.PimCandidateRPEntry[i]})
    }
    pimCandidateRPTable.EntityData.Leafs = types.NewOrderedMap()

    pimCandidateRPTable.EntityData.YListKeys = []string {}

    return &(pimCandidateRPTable.EntityData)
}

// PIMMIB_PimCandidateRPTable_PimCandidateRPEntry
// An entry (conceptual row) in the pimCandidateRPTable.
type PIMMIB_PimCandidateRPTable_PimCandidateRPEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP multicast group address which, when
    // combined with pimCandidateRPGroupMask, identifies a group prefix for which
    // the local router will advertise itself as a Candidate-RP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimCandidateRPGroupAddress interface{}

    // This attribute is a key. The multicast group address mask which, when
    // combined with pimCandidateRPGroupMask, identifies a group prefix for which
    // the local router will advertise itself as a Candidate-RP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimCandidateRPGroupMask interface{}

    // The (unicast) address of the interface which will be      advertised as a
    // Candidate-RP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimCandidateRPAddress interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table. The type is RowStatus.
    PimCandidateRPRowStatus interface{}
}

func (pimCandidateRPEntry *PIMMIB_PimCandidateRPTable_PimCandidateRPEntry) GetEntityData() *types.CommonEntityData {
    pimCandidateRPEntry.EntityData.YFilter = pimCandidateRPEntry.YFilter
    pimCandidateRPEntry.EntityData.YangName = "pimCandidateRPEntry"
    pimCandidateRPEntry.EntityData.BundleName = "cisco_ios_xe"
    pimCandidateRPEntry.EntityData.ParentYangName = "pimCandidateRPTable"
    pimCandidateRPEntry.EntityData.SegmentPath = "pimCandidateRPEntry" + types.AddKeyToken(pimCandidateRPEntry.PimCandidateRPGroupAddress, "pimCandidateRPGroupAddress") + types.AddKeyToken(pimCandidateRPEntry.PimCandidateRPGroupMask, "pimCandidateRPGroupMask")
    pimCandidateRPEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimCandidateRPTable/" + pimCandidateRPEntry.EntityData.SegmentPath
    pimCandidateRPEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimCandidateRPEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimCandidateRPEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimCandidateRPEntry.EntityData.Children = types.NewOrderedMap()
    pimCandidateRPEntry.EntityData.Leafs = types.NewOrderedMap()
    pimCandidateRPEntry.EntityData.Leafs.Append("pimCandidateRPGroupAddress", types.YLeaf{"PimCandidateRPGroupAddress", pimCandidateRPEntry.PimCandidateRPGroupAddress})
    pimCandidateRPEntry.EntityData.Leafs.Append("pimCandidateRPGroupMask", types.YLeaf{"PimCandidateRPGroupMask", pimCandidateRPEntry.PimCandidateRPGroupMask})
    pimCandidateRPEntry.EntityData.Leafs.Append("pimCandidateRPAddress", types.YLeaf{"PimCandidateRPAddress", pimCandidateRPEntry.PimCandidateRPAddress})
    pimCandidateRPEntry.EntityData.Leafs.Append("pimCandidateRPRowStatus", types.YLeaf{"PimCandidateRPRowStatus", pimCandidateRPEntry.PimCandidateRPRowStatus})

    pimCandidateRPEntry.EntityData.YListKeys = []string {"PimCandidateRPGroupAddress", "PimCandidateRPGroupMask"}

    return &(pimCandidateRPEntry.EntityData)
}

// PIMMIB_PimComponentTable
// The (conceptual) table containing objects specific to a PIM
// domain.  One row exists for each domain to which the router
// is connected.  A PIM-SM domain is defined as an area of the
// network over which Bootstrap messages are forwarded.
// Typically, a PIM-SM router will be a member of exactly one
// domain.  This table also supports, however, routers which
// may form a border between two PIM-SM domains and do not
// forward Bootstrap messages between them.
type PIMMIB_PimComponentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the pimComponentTable. The type is slice of
    // PIMMIB_PimComponentTable_PimComponentEntry.
    PimComponentEntry []*PIMMIB_PimComponentTable_PimComponentEntry
}

func (pimComponentTable *PIMMIB_PimComponentTable) GetEntityData() *types.CommonEntityData {
    pimComponentTable.EntityData.YFilter = pimComponentTable.YFilter
    pimComponentTable.EntityData.YangName = "pimComponentTable"
    pimComponentTable.EntityData.BundleName = "cisco_ios_xe"
    pimComponentTable.EntityData.ParentYangName = "PIM-MIB"
    pimComponentTable.EntityData.SegmentPath = "pimComponentTable"
    pimComponentTable.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/" + pimComponentTable.EntityData.SegmentPath
    pimComponentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimComponentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimComponentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimComponentTable.EntityData.Children = types.NewOrderedMap()
    pimComponentTable.EntityData.Children.Append("pimComponentEntry", types.YChild{"PimComponentEntry", nil})
    for i := range pimComponentTable.PimComponentEntry {
        pimComponentTable.EntityData.Children.Append(types.GetSegmentPath(pimComponentTable.PimComponentEntry[i]), types.YChild{"PimComponentEntry", pimComponentTable.PimComponentEntry[i]})
    }
    pimComponentTable.EntityData.Leafs = types.NewOrderedMap()

    pimComponentTable.EntityData.YListKeys = []string {}

    return &(pimComponentTable.EntityData)
}

// PIMMIB_PimComponentTable_PimComponentEntry
// An entry (conceptual row) in the pimComponentTable.
type PIMMIB_PimComponentTable_PimComponentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A number uniquely identifying the component.  Each
    // protocol instance connected to a separate domain should have a different
    // index value.  Routers that only support membership in a single PIM-SM
    // domain should use a pimComponentIndex value of 1. The type is interface{}
    // with range: 1..255.
    PimComponentIndex interface{}

    // The IP address of the bootstrap router (BSR) for the local PIM region. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PimComponentBSRAddress interface{}

    // The minimum time remaining before the bootstrap router in the local domain
    // will be declared down.  For candidate BSRs, this is the time until the
    // component sends an RP-Set message.  For other routers, this is the time
    // until it may accept an RP-Set message from a lower candidate BSR. The type
    // is interface{} with range: 0..4294967295.
    PimComponentBSRExpiryTime interface{}

    // The holdtime of the component when it is a candidate RP in the local
    // domain.  The value of 0 is used to indicate that the local system is not a
    // Candidate-RP. The type is interface{} with range: 0..255. Units are
    // seconds.
    PimComponentCRPHoldTime interface{}

    // The status of this entry.  Creating the entry creates another protocol
    // instance; destroying the entry disables a protocol instance. The type is
    // RowStatus.
    PimComponentStatus interface{}
}

func (pimComponentEntry *PIMMIB_PimComponentTable_PimComponentEntry) GetEntityData() *types.CommonEntityData {
    pimComponentEntry.EntityData.YFilter = pimComponentEntry.YFilter
    pimComponentEntry.EntityData.YangName = "pimComponentEntry"
    pimComponentEntry.EntityData.BundleName = "cisco_ios_xe"
    pimComponentEntry.EntityData.ParentYangName = "pimComponentTable"
    pimComponentEntry.EntityData.SegmentPath = "pimComponentEntry" + types.AddKeyToken(pimComponentEntry.PimComponentIndex, "pimComponentIndex")
    pimComponentEntry.EntityData.AbsolutePath = "PIM-MIB:PIM-MIB/pimComponentTable/" + pimComponentEntry.EntityData.SegmentPath
    pimComponentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pimComponentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pimComponentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pimComponentEntry.EntityData.Children = types.NewOrderedMap()
    pimComponentEntry.EntityData.Leafs = types.NewOrderedMap()
    pimComponentEntry.EntityData.Leafs.Append("pimComponentIndex", types.YLeaf{"PimComponentIndex", pimComponentEntry.PimComponentIndex})
    pimComponentEntry.EntityData.Leafs.Append("pimComponentBSRAddress", types.YLeaf{"PimComponentBSRAddress", pimComponentEntry.PimComponentBSRAddress})
    pimComponentEntry.EntityData.Leafs.Append("pimComponentBSRExpiryTime", types.YLeaf{"PimComponentBSRExpiryTime", pimComponentEntry.PimComponentBSRExpiryTime})
    pimComponentEntry.EntityData.Leafs.Append("pimComponentCRPHoldTime", types.YLeaf{"PimComponentCRPHoldTime", pimComponentEntry.PimComponentCRPHoldTime})
    pimComponentEntry.EntityData.Leafs.Append("pimComponentStatus", types.YLeaf{"PimComponentStatus", pimComponentEntry.PimComponentStatus})

    pimComponentEntry.EntityData.YListKeys = []string {"PimComponentIndex"}

    return &(pimComponentEntry.EntityData)
}

