// The MIB module for the display of CIDR multipath IP Routes.
package ip_forward_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_forward_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:IP-FORWARD-MIB IP-FORWARD-MIB}", reflect.TypeOf(IPFORWARDMIB{}))
    ydk.RegisterEntity("IP-FORWARD-MIB:IP-FORWARD-MIB", reflect.TypeOf(IPFORWARDMIB{}))
}

// IPFORWARDMIB
type IPFORWARDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    IpForward IPFORWARDMIB_IpForward

    // This entity's IP Routing table.
    IpForwardTable IPFORWARDMIB_IpForwardTable

    // This entity's IP Routing table.
    IpCidrRouteTable IPFORWARDMIB_IpCidrRouteTable
}

func (iPFORWARDMIB *IPFORWARDMIB) GetEntityData() *types.CommonEntityData {
    iPFORWARDMIB.EntityData.YFilter = iPFORWARDMIB.YFilter
    iPFORWARDMIB.EntityData.YangName = "IP-FORWARD-MIB"
    iPFORWARDMIB.EntityData.BundleName = "cisco_ios_xe"
    iPFORWARDMIB.EntityData.ParentYangName = "IP-FORWARD-MIB"
    iPFORWARDMIB.EntityData.SegmentPath = "IP-FORWARD-MIB:IP-FORWARD-MIB"
    iPFORWARDMIB.EntityData.AbsolutePath = iPFORWARDMIB.EntityData.SegmentPath
    iPFORWARDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iPFORWARDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iPFORWARDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iPFORWARDMIB.EntityData.Children = types.NewOrderedMap()
    iPFORWARDMIB.EntityData.Children.Append("ipForward", types.YChild{"IpForward", &iPFORWARDMIB.IpForward})
    iPFORWARDMIB.EntityData.Children.Append("ipForwardTable", types.YChild{"IpForwardTable", &iPFORWARDMIB.IpForwardTable})
    iPFORWARDMIB.EntityData.Children.Append("ipCidrRouteTable", types.YChild{"IpCidrRouteTable", &iPFORWARDMIB.IpCidrRouteTable})
    iPFORWARDMIB.EntityData.Leafs = types.NewOrderedMap()

    iPFORWARDMIB.EntityData.YListKeys = []string {}

    return &(iPFORWARDMIB.EntityData)
}

// IPFORWARDMIB_IpForward
type IPFORWARDMIB_IpForward struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of current  ipForwardTable  entries that are not invalid. The
    // type is interface{} with range: 0..4294967295.
    IpForwardNumber interface{}

    // The number of current ipCidrRouteTable entries that are not invalid. The
    // type is interface{} with range: 0..4294967295.
    IpCidrRouteNumber interface{}
}

func (ipForward *IPFORWARDMIB_IpForward) GetEntityData() *types.CommonEntityData {
    ipForward.EntityData.YFilter = ipForward.YFilter
    ipForward.EntityData.YangName = "ipForward"
    ipForward.EntityData.BundleName = "cisco_ios_xe"
    ipForward.EntityData.ParentYangName = "IP-FORWARD-MIB"
    ipForward.EntityData.SegmentPath = "ipForward"
    ipForward.EntityData.AbsolutePath = "IP-FORWARD-MIB:IP-FORWARD-MIB/" + ipForward.EntityData.SegmentPath
    ipForward.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipForward.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipForward.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipForward.EntityData.Children = types.NewOrderedMap()
    ipForward.EntityData.Leafs = types.NewOrderedMap()
    ipForward.EntityData.Leafs.Append("ipForwardNumber", types.YLeaf{"IpForwardNumber", ipForward.IpForwardNumber})
    ipForward.EntityData.Leafs.Append("ipCidrRouteNumber", types.YLeaf{"IpCidrRouteNumber", ipForward.IpCidrRouteNumber})

    ipForward.EntityData.YListKeys = []string {}

    return &(ipForward.EntityData)
}

// IPFORWARDMIB_IpForwardTable
// This entity's IP Routing table.
type IPFORWARDMIB_IpForwardTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A particular route to  a  particular  destina- tion, under a particular
    // policy. The type is slice of IPFORWARDMIB_IpForwardTable_IpForwardEntry.
    IpForwardEntry []*IPFORWARDMIB_IpForwardTable_IpForwardEntry
}

func (ipForwardTable *IPFORWARDMIB_IpForwardTable) GetEntityData() *types.CommonEntityData {
    ipForwardTable.EntityData.YFilter = ipForwardTable.YFilter
    ipForwardTable.EntityData.YangName = "ipForwardTable"
    ipForwardTable.EntityData.BundleName = "cisco_ios_xe"
    ipForwardTable.EntityData.ParentYangName = "IP-FORWARD-MIB"
    ipForwardTable.EntityData.SegmentPath = "ipForwardTable"
    ipForwardTable.EntityData.AbsolutePath = "IP-FORWARD-MIB:IP-FORWARD-MIB/" + ipForwardTable.EntityData.SegmentPath
    ipForwardTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipForwardTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipForwardTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipForwardTable.EntityData.Children = types.NewOrderedMap()
    ipForwardTable.EntityData.Children.Append("ipForwardEntry", types.YChild{"IpForwardEntry", nil})
    for i := range ipForwardTable.IpForwardEntry {
        ipForwardTable.EntityData.Children.Append(types.GetSegmentPath(ipForwardTable.IpForwardEntry[i]), types.YChild{"IpForwardEntry", ipForwardTable.IpForwardEntry[i]})
    }
    ipForwardTable.EntityData.Leafs = types.NewOrderedMap()

    ipForwardTable.EntityData.YListKeys = []string {}

    return &(ipForwardTable.EntityData)
}

// IPFORWARDMIB_IpForwardTable_IpForwardEntry
// A particular route to  a  particular  destina-
// tion, under a particular policy.
type IPFORWARDMIB_IpForwardTable_IpForwardEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The destination IP address of this route.   An
    // entry  with  a value of 0.0.0.0 is considered a default route.  This object
    // may not take a Multicast (Class  D) address value.  Any assignment
    // (implicit or  otherwise)  of  an instance  of  this  object to a value x
    // must be rejected if the bitwise logical-AND of  x  with the  value of the
    // corresponding instance of the ipForwardMask object is not equal to x. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpForwardDest interface{}

    // This attribute is a key. The routing mechanism via which this route was
    // learned.  Inclusion of values for gateway rout- ing protocols is not 
    // intended  to  imply  that hosts should support those protocols. The type is
    // IpForwardProto.
    IpForwardProto interface{}

    // This attribute is a key. The general set of conditions that would cause the
    // selection  of  one multipath route (set of next hops for a given
    // destination) is  referred to as 'policy'.  Unless the mechanism indicated
    // by ipForwardPro- to specifies otherwise, the policy specifier is the IP TOS
    // Field.  The encoding of IP TOS is as  specified  by  the  following
    // convention.  Zero indicates the default path if no more  specific policy
    // applies.  +-----+-----+-----+-----+-----+-----+-----+-----+ |              
    // |                       |     | |   PRECEDENCE    |    TYPE OF SERVICE    |
    // 0  | |                 |                       |     |
    // +-----+-----+-----+-----+-----+-----+-----+-----+           IP TOS         
    // IP TOS     Field     Policy      Field     Policy     Contents    Code     
    // Contents    Code     0 0 0 0  ==>   0      0 0 0 1  ==>   2     0 0 1 0 
    // ==>   4      0 0 1 1  ==>   6     0 1 0 0  ==>   8      0 1 0 1  ==>  10   
    // 0 1 1 0  ==>  12      0 1 1 1  ==>  14     1 0 0 0  ==>  16      1 0 0 1 
    // ==>  18     1 0 1 0  ==>  20      1 0 1 1  ==>  22     1 1 0 0  ==>  24    
    // 1 1 0 1  ==>  26     1 1 1 0  ==>  28      1 1 1 1  ==>  30  Protocols
    // defining 'policy' otherwise must  ei- ther define a set of values which are
    // valid for this  object  or  must  implement  an  integer- instanced  policy
    // table for which this object's value acts as an index. The type is
    // interface{} with range: -2147483648..2147483647.
    IpForwardPolicy interface{}

    // This attribute is a key. On remote routes, the address of the next sys- tem
    // en route; Otherwise, 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpForwardNextHop interface{}

    // Indicate the mask to be logical-ANDed with the destination  address  before
    // being compared to the value  in  the  ipForwardDest  field.   For those 
    // systems  that  do  not support arbitrary subnet masks, an agent constructs
    // the value  of the  ipForwardMask  by  reference to the IP Ad- dress Class. 
    // Any assignment (implicit or  otherwise)  of  an instance  of  this  object
    // to a value x must be rejected if the bitwise logical-AND of  x  with the 
    // value of the corresponding instance of the ipForwardDest object is not
    // equal to ipForward- Dest. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpForwardMask interface{}

    // The ifIndex value which identifies  the  local interface  through  which 
    // the next hop of this route should be reached. The type is interface{} with
    // range: -2147483648..2147483647.
    IpForwardIfIndex interface{}

    // The type of route.  Note that local(3)  refers to  a route for which the
    // next hop is the final destination; remote(4) refers to  a  route  for which
    // the  next  hop is not the final destina- tion.  Setting this object to the
    // value invalid(2) has the  effect  of  invalidating the corresponding entry
    // in the ipForwardTable object.   That  is, it  effectively  disassociates 
    // the destination identified with said entry from the route iden- tified   
    // with    said   entry.    It   is   an implementation-specific matter  as 
    // to  whether the agent removes an invalidated entry from the table. 
    // Accordingly, management  stations  must be prepared to receive tabular
    // information from agents that corresponds to entries not current- ly  in 
    // use.  Proper interpretation of such en- tries requires examination of the
    // relevant  ip- ForwardType object. The type is IpForwardType.
    IpForwardType interface{}

    // The number of seconds  since  this  route  was last  updated  or  otherwise
    // determined  to be correct.  Note that no semantics of  `too  old' can  be
    // implied except through knowledge of the routing  protocol  by  which  the  
    // route   was learned. The type is interface{} with range:
    // -2147483648..2147483647.
    IpForwardAge interface{}

    // A reference to MIB definitions specific to the particular  routing protocol
    // which is responsi- ble for this route, as determined by the  value
    // specified  in the route's ipForwardProto value. If this information is not
    // present,  its  value should be set to the OBJECT IDENTIFIER { 0 0 }, which
    // is a syntactically valid object  identif- ier, and any implementation
    // conforming to ASN.1 and the Basic Encoding Rules must  be  able  to
    // generate and recognize this value. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    IpForwardInfo interface{}

    // The Autonomous System Number of the Next  Hop. When  this  is  unknown  or
    // not relevant to the protocol indicated by ipForwardProto, zero. The type is
    // interface{} with range: -2147483648..2147483647.
    IpForwardNextHopAS interface{}

    // The primary routing  metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpForwardMetric1 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpForwardMetric2 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpForwardMetric3 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpForwardMetric4 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpForwardMetric5 interface{}
}

func (ipForwardEntry *IPFORWARDMIB_IpForwardTable_IpForwardEntry) GetEntityData() *types.CommonEntityData {
    ipForwardEntry.EntityData.YFilter = ipForwardEntry.YFilter
    ipForwardEntry.EntityData.YangName = "ipForwardEntry"
    ipForwardEntry.EntityData.BundleName = "cisco_ios_xe"
    ipForwardEntry.EntityData.ParentYangName = "ipForwardTable"
    ipForwardEntry.EntityData.SegmentPath = "ipForwardEntry" + types.AddKeyToken(ipForwardEntry.IpForwardDest, "ipForwardDest") + types.AddKeyToken(ipForwardEntry.IpForwardProto, "ipForwardProto") + types.AddKeyToken(ipForwardEntry.IpForwardPolicy, "ipForwardPolicy") + types.AddKeyToken(ipForwardEntry.IpForwardNextHop, "ipForwardNextHop")
    ipForwardEntry.EntityData.AbsolutePath = "IP-FORWARD-MIB:IP-FORWARD-MIB/ipForwardTable/" + ipForwardEntry.EntityData.SegmentPath
    ipForwardEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipForwardEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipForwardEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipForwardEntry.EntityData.Children = types.NewOrderedMap()
    ipForwardEntry.EntityData.Leafs = types.NewOrderedMap()
    ipForwardEntry.EntityData.Leafs.Append("ipForwardDest", types.YLeaf{"IpForwardDest", ipForwardEntry.IpForwardDest})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardProto", types.YLeaf{"IpForwardProto", ipForwardEntry.IpForwardProto})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardPolicy", types.YLeaf{"IpForwardPolicy", ipForwardEntry.IpForwardPolicy})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardNextHop", types.YLeaf{"IpForwardNextHop", ipForwardEntry.IpForwardNextHop})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardMask", types.YLeaf{"IpForwardMask", ipForwardEntry.IpForwardMask})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardIfIndex", types.YLeaf{"IpForwardIfIndex", ipForwardEntry.IpForwardIfIndex})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardType", types.YLeaf{"IpForwardType", ipForwardEntry.IpForwardType})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardAge", types.YLeaf{"IpForwardAge", ipForwardEntry.IpForwardAge})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardInfo", types.YLeaf{"IpForwardInfo", ipForwardEntry.IpForwardInfo})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardNextHopAS", types.YLeaf{"IpForwardNextHopAS", ipForwardEntry.IpForwardNextHopAS})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardMetric1", types.YLeaf{"IpForwardMetric1", ipForwardEntry.IpForwardMetric1})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardMetric2", types.YLeaf{"IpForwardMetric2", ipForwardEntry.IpForwardMetric2})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardMetric3", types.YLeaf{"IpForwardMetric3", ipForwardEntry.IpForwardMetric3})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardMetric4", types.YLeaf{"IpForwardMetric4", ipForwardEntry.IpForwardMetric4})
    ipForwardEntry.EntityData.Leafs.Append("ipForwardMetric5", types.YLeaf{"IpForwardMetric5", ipForwardEntry.IpForwardMetric5})

    ipForwardEntry.EntityData.YListKeys = []string {"IpForwardDest", "IpForwardProto", "IpForwardPolicy", "IpForwardNextHop"}

    return &(ipForwardEntry.EntityData)
}

// IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto represents hosts should support those protocols.
type IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto string

const (
    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_other IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "other"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_local IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "local"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_netmgmt IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "netmgmt"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_icmp IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "icmp"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_egp IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "egp"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_ggp IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "ggp"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_hello IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "hello"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_rip IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "rip"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_is_is IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "is-is"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_es_is IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "es-is"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_ciscoIgrp IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "ciscoIgrp"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_bbnSpfIgp IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "bbnSpfIgp"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_ospf IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "ospf"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_bgp IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "bgp"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto_idpr IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardProto = "idpr"
)

// IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType represents ForwardType object.
type IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType string

const (
    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType_other IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType = "other"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType_invalid IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType = "invalid"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType_local IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType = "local"

    IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType_remote IPFORWARDMIB_IpForwardTable_IpForwardEntry_IpForwardType = "remote"
)

// IPFORWARDMIB_IpCidrRouteTable
// This entity's IP Routing table.
type IPFORWARDMIB_IpCidrRouteTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A particular route to  a  particular  destina- tion, under a particular
    // policy. The type is slice of
    // IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry.
    IpCidrRouteEntry []*IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry
}

func (ipCidrRouteTable *IPFORWARDMIB_IpCidrRouteTable) GetEntityData() *types.CommonEntityData {
    ipCidrRouteTable.EntityData.YFilter = ipCidrRouteTable.YFilter
    ipCidrRouteTable.EntityData.YangName = "ipCidrRouteTable"
    ipCidrRouteTable.EntityData.BundleName = "cisco_ios_xe"
    ipCidrRouteTable.EntityData.ParentYangName = "IP-FORWARD-MIB"
    ipCidrRouteTable.EntityData.SegmentPath = "ipCidrRouteTable"
    ipCidrRouteTable.EntityData.AbsolutePath = "IP-FORWARD-MIB:IP-FORWARD-MIB/" + ipCidrRouteTable.EntityData.SegmentPath
    ipCidrRouteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipCidrRouteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipCidrRouteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipCidrRouteTable.EntityData.Children = types.NewOrderedMap()
    ipCidrRouteTable.EntityData.Children.Append("ipCidrRouteEntry", types.YChild{"IpCidrRouteEntry", nil})
    for i := range ipCidrRouteTable.IpCidrRouteEntry {
        ipCidrRouteTable.EntityData.Children.Append(types.GetSegmentPath(ipCidrRouteTable.IpCidrRouteEntry[i]), types.YChild{"IpCidrRouteEntry", ipCidrRouteTable.IpCidrRouteEntry[i]})
    }
    ipCidrRouteTable.EntityData.Leafs = types.NewOrderedMap()

    ipCidrRouteTable.EntityData.YListKeys = []string {}

    return &(ipCidrRouteTable.EntityData)
}

// IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry
// A particular route to  a  particular  destina-
// tion, under a particular policy.
type IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The destination IP address of this route.  This
    // object may not take a Multicast (Class  D) address value.  Any assignment
    // (implicit or  otherwise)  of  an instance  of  this  object to a value x
    // must be rejected if the bitwise logical-AND of  x  with the  value of the
    // corresponding instance of the ipCidrRouteMask object is not equal to x. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpCidrRouteDest interface{}

    // This attribute is a key. Indicate the mask to be logical-ANDed with the
    // destination  address  before  being compared to the value  in  the 
    // ipCidrRouteDest  field.   For those  systems  that  do  not support
    // arbitrary subnet masks, an agent constructs the value  of the 
    // ipCidrRouteMask  by  reference to the IP Ad- dress Class.  Any assignment
    // (implicit or  otherwise)  of  an instance  of  this  object to a value x
    // must be rejected if the bitwise logical-AND of  x  with the  value of the
    // corresponding instance of the ipCidrRouteDest object is not equal to
    // ipCidrRoute- Dest. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpCidrRouteMask interface{}

    // This attribute is a key. The policy specifier is the IP TOS Field.  The
    // encoding of IP TOS is as specified  by  the  following convention. Zero
    // indicates the default path if no more  specific policy applies. 
    // +-----+-----+-----+-----+-----+-----+-----+-----+ |                 |      
    // |     | |   PRECEDENCE    |    TYPE OF SERVICE    |  0  | |                
    // |                       |     |
    // +-----+-----+-----+-----+-----+-----+-----+-----+           IP TOS         
    // IP TOS     Field     Policy      Field     Policy     Contents    Code     
    // Contents    Code     0 0 0 0  ==>   0      0 0 0 1  ==>   2     0 0 1 0 
    // ==>   4      0 0 1 1  ==>   6     0 1 0 0  ==>   8      0 1 0 1  ==>  10   
    // 0 1 1 0  ==>  12      0 1 1 1  ==>  14     1 0 0 0  ==>  16      1 0 0 1 
    // ==>  18     1 0 1 0  ==>  20      1 0 1 1  ==>  22     1 1 0 0  ==>  24    
    // 1 1 0 1  ==>  26     1 1 1 0  ==>  28      1 1 1 1  ==>  30. The type is
    // interface{} with range: -2147483648..2147483647.
    IpCidrRouteTos interface{}

    // This attribute is a key. On remote routes, the address of the next sys- tem
    // en route; Otherwise, 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpCidrRouteNextHop interface{}

    // The ifIndex value which identifies  the  local interface  through  which 
    // the next hop of this route should be reached. The type is interface{} with
    // range: -2147483648..2147483647.
    IpCidrRouteIfIndex interface{}

    // The type of route.  Note that local(3)  refers to  a route for which the
    // next hop is the final destination; remote(4) refers to  a  route  for which
    // the  next  hop is not the final destina- tion.  Routes which do not result
    // in traffic forwarding or rejection should not be displayed even if the
    // implementation keeps them stored internally.   reject (2) refers to a route
    // which, if matched, discards the message as unreachable. This is used in
    // some protocols as a means of correctly aggregating routes. The type is
    // IpCidrRouteType.
    IpCidrRouteType interface{}

    // The routing mechanism via which this route was learned.  Inclusion of
    // values for gateway rout- ing protocols is not  intended  to  imply  that
    // hosts should support those protocols. The type is IpCidrRouteProto.
    IpCidrRouteProto interface{}

    // The number of seconds  since  this  route  was last  updated  or  otherwise
    // determined  to be correct.  Note that no semantics of  `too  old' can  be
    // implied except through knowledge of the routing  protocol  by  which  the  
    // route   was learned. The type is interface{} with range:
    // -2147483648..2147483647.
    IpCidrRouteAge interface{}

    // A reference to MIB definitions specific to the particular  routing protocol
    // which is responsi- ble for this route, as determined by the  value
    // specified  in the route's ipCidrRouteProto value. If this information is
    // not present,  its  value should be set to the OBJECT IDENTIFIER { 0 0 },
    // which is a syntactically valid object  identif- ier, and any implementation
    // conforming to ASN.1 and the Basic Encoding Rules must  be  able  to
    // generate and recognize this value. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    IpCidrRouteInfo interface{}

    // The Autonomous System Number of the Next  Hop. The  semantics of this
    // object are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value. When  this object is unknown or not relevant its
    // value should be set to zero. The type is interface{} with range:
    // -2147483648..2147483647.
    IpCidrRouteNextHopAS interface{}

    // The primary routing  metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpCidrRouteMetric1 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpCidrRouteMetric2 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpCidrRouteMetric3 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpCidrRouteMetric4 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    IpCidrRouteMetric5 interface{}

    // The row status variable, used according to row installation and removal
    // conventions. The type is RowStatus.
    IpCidrRouteStatus interface{}
}

func (ipCidrRouteEntry *IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry) GetEntityData() *types.CommonEntityData {
    ipCidrRouteEntry.EntityData.YFilter = ipCidrRouteEntry.YFilter
    ipCidrRouteEntry.EntityData.YangName = "ipCidrRouteEntry"
    ipCidrRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    ipCidrRouteEntry.EntityData.ParentYangName = "ipCidrRouteTable"
    ipCidrRouteEntry.EntityData.SegmentPath = "ipCidrRouteEntry" + types.AddKeyToken(ipCidrRouteEntry.IpCidrRouteDest, "ipCidrRouteDest") + types.AddKeyToken(ipCidrRouteEntry.IpCidrRouteMask, "ipCidrRouteMask") + types.AddKeyToken(ipCidrRouteEntry.IpCidrRouteTos, "ipCidrRouteTos") + types.AddKeyToken(ipCidrRouteEntry.IpCidrRouteNextHop, "ipCidrRouteNextHop")
    ipCidrRouteEntry.EntityData.AbsolutePath = "IP-FORWARD-MIB:IP-FORWARD-MIB/ipCidrRouteTable/" + ipCidrRouteEntry.EntityData.SegmentPath
    ipCidrRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipCidrRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipCidrRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipCidrRouteEntry.EntityData.Children = types.NewOrderedMap()
    ipCidrRouteEntry.EntityData.Leafs = types.NewOrderedMap()
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteDest", types.YLeaf{"IpCidrRouteDest", ipCidrRouteEntry.IpCidrRouteDest})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteMask", types.YLeaf{"IpCidrRouteMask", ipCidrRouteEntry.IpCidrRouteMask})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteTos", types.YLeaf{"IpCidrRouteTos", ipCidrRouteEntry.IpCidrRouteTos})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteNextHop", types.YLeaf{"IpCidrRouteNextHop", ipCidrRouteEntry.IpCidrRouteNextHop})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteIfIndex", types.YLeaf{"IpCidrRouteIfIndex", ipCidrRouteEntry.IpCidrRouteIfIndex})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteType", types.YLeaf{"IpCidrRouteType", ipCidrRouteEntry.IpCidrRouteType})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteProto", types.YLeaf{"IpCidrRouteProto", ipCidrRouteEntry.IpCidrRouteProto})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteAge", types.YLeaf{"IpCidrRouteAge", ipCidrRouteEntry.IpCidrRouteAge})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteInfo", types.YLeaf{"IpCidrRouteInfo", ipCidrRouteEntry.IpCidrRouteInfo})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteNextHopAS", types.YLeaf{"IpCidrRouteNextHopAS", ipCidrRouteEntry.IpCidrRouteNextHopAS})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteMetric1", types.YLeaf{"IpCidrRouteMetric1", ipCidrRouteEntry.IpCidrRouteMetric1})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteMetric2", types.YLeaf{"IpCidrRouteMetric2", ipCidrRouteEntry.IpCidrRouteMetric2})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteMetric3", types.YLeaf{"IpCidrRouteMetric3", ipCidrRouteEntry.IpCidrRouteMetric3})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteMetric4", types.YLeaf{"IpCidrRouteMetric4", ipCidrRouteEntry.IpCidrRouteMetric4})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteMetric5", types.YLeaf{"IpCidrRouteMetric5", ipCidrRouteEntry.IpCidrRouteMetric5})
    ipCidrRouteEntry.EntityData.Leafs.Append("ipCidrRouteStatus", types.YLeaf{"IpCidrRouteStatus", ipCidrRouteEntry.IpCidrRouteStatus})

    ipCidrRouteEntry.EntityData.YListKeys = []string {"IpCidrRouteDest", "IpCidrRouteMask", "IpCidrRouteTos", "IpCidrRouteNextHop"}

    return &(ipCidrRouteEntry.EntityData)
}

// IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto represents hosts should support those protocols.
type IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto string

const (
    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_other IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "other"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_local IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "local"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_netmgmt IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "netmgmt"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_icmp IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "icmp"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_egp IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "egp"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_ggp IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "ggp"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_hello IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "hello"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_rip IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "rip"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_isIs IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "isIs"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_esIs IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "esIs"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_ciscoIgrp IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "ciscoIgrp"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_bbnSpfIgp IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "bbnSpfIgp"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_ospf IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "ospf"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_bgp IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "bgp"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_idpr IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "idpr"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto_ciscoEigrp IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteProto = "ciscoEigrp"
)

// IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType represents protocols as a means of correctly aggregating routes.
type IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType string

const (
    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType_other IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType = "other"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType_reject IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType = "reject"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType_local IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType = "local"

    IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType_remote IPFORWARDMIB_IpCidrRouteTable_IpCidrRouteEntry_IpCidrRouteType = "remote"
)

