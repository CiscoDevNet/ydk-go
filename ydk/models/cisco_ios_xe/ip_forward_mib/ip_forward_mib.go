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

    
    Ipforward IPFORWARDMIB_Ipforward

    // This entity's IP Routing table.
    Ipforwardtable IPFORWARDMIB_Ipforwardtable

    // This entity's IP Routing table.
    Ipcidrroutetable IPFORWARDMIB_Ipcidrroutetable
}

func (iPFORWARDMIB *IPFORWARDMIB) GetEntityData() *types.CommonEntityData {
    iPFORWARDMIB.EntityData.YFilter = iPFORWARDMIB.YFilter
    iPFORWARDMIB.EntityData.YangName = "IP-FORWARD-MIB"
    iPFORWARDMIB.EntityData.BundleName = "cisco_ios_xe"
    iPFORWARDMIB.EntityData.ParentYangName = "IP-FORWARD-MIB"
    iPFORWARDMIB.EntityData.SegmentPath = "IP-FORWARD-MIB:IP-FORWARD-MIB"
    iPFORWARDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iPFORWARDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iPFORWARDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iPFORWARDMIB.EntityData.Children = make(map[string]types.YChild)
    iPFORWARDMIB.EntityData.Children["ipForward"] = types.YChild{"Ipforward", &iPFORWARDMIB.Ipforward}
    iPFORWARDMIB.EntityData.Children["ipForwardTable"] = types.YChild{"Ipforwardtable", &iPFORWARDMIB.Ipforwardtable}
    iPFORWARDMIB.EntityData.Children["ipCidrRouteTable"] = types.YChild{"Ipcidrroutetable", &iPFORWARDMIB.Ipcidrroutetable}
    iPFORWARDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iPFORWARDMIB.EntityData)
}

// IPFORWARDMIB_Ipforward
type IPFORWARDMIB_Ipforward struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of current  ipForwardTable  entries that are not invalid. The
    // type is interface{} with range: 0..4294967295.
    Ipforwardnumber interface{}

    // The number of current ipCidrRouteTable entries that are not invalid. The
    // type is interface{} with range: 0..4294967295.
    Ipcidrroutenumber interface{}
}

func (ipforward *IPFORWARDMIB_Ipforward) GetEntityData() *types.CommonEntityData {
    ipforward.EntityData.YFilter = ipforward.YFilter
    ipforward.EntityData.YangName = "ipForward"
    ipforward.EntityData.BundleName = "cisco_ios_xe"
    ipforward.EntityData.ParentYangName = "IP-FORWARD-MIB"
    ipforward.EntityData.SegmentPath = "ipForward"
    ipforward.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipforward.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipforward.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipforward.EntityData.Children = make(map[string]types.YChild)
    ipforward.EntityData.Leafs = make(map[string]types.YLeaf)
    ipforward.EntityData.Leafs["ipForwardNumber"] = types.YLeaf{"Ipforwardnumber", ipforward.Ipforwardnumber}
    ipforward.EntityData.Leafs["ipCidrRouteNumber"] = types.YLeaf{"Ipcidrroutenumber", ipforward.Ipcidrroutenumber}
    return &(ipforward.EntityData)
}

// IPFORWARDMIB_Ipforwardtable
// This entity's IP Routing table.
type IPFORWARDMIB_Ipforwardtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A particular route to  a  particular  destina- tion, under a particular
    // policy. The type is slice of IPFORWARDMIB_Ipforwardtable_Ipforwardentry.
    Ipforwardentry []IPFORWARDMIB_Ipforwardtable_Ipforwardentry
}

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetEntityData() *types.CommonEntityData {
    ipforwardtable.EntityData.YFilter = ipforwardtable.YFilter
    ipforwardtable.EntityData.YangName = "ipForwardTable"
    ipforwardtable.EntityData.BundleName = "cisco_ios_xe"
    ipforwardtable.EntityData.ParentYangName = "IP-FORWARD-MIB"
    ipforwardtable.EntityData.SegmentPath = "ipForwardTable"
    ipforwardtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipforwardtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipforwardtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipforwardtable.EntityData.Children = make(map[string]types.YChild)
    ipforwardtable.EntityData.Children["ipForwardEntry"] = types.YChild{"Ipforwardentry", nil}
    for i := range ipforwardtable.Ipforwardentry {
        ipforwardtable.EntityData.Children[types.GetSegmentPath(&ipforwardtable.Ipforwardentry[i])] = types.YChild{"Ipforwardentry", &ipforwardtable.Ipforwardentry[i]}
    }
    ipforwardtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipforwardtable.EntityData)
}

// IPFORWARDMIB_Ipforwardtable_Ipforwardentry
// A particular route to  a  particular  destina-
// tion, under a particular policy.
type IPFORWARDMIB_Ipforwardtable_Ipforwardentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The destination IP address of this route.   An
    // entry  with  a value of 0.0.0.0 is considered a default route.  This object
    // may not take a Multicast (Class  D) address value.  Any assignment
    // (implicit or  otherwise)  of  an instance  of  this  object to a value x
    // must be rejected if the bitwise logical-AND of  x  with the  value of the
    // corresponding instance of the ipForwardMask object is not equal to x. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipforwarddest interface{}

    // This attribute is a key. The routing mechanism via which this route was
    // learned.  Inclusion of values for gateway rout- ing protocols is not 
    // intended  to  imply  that hosts should support those protocols. The type is
    // Ipforwardproto.
    Ipforwardproto interface{}

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
    Ipforwardpolicy interface{}

    // This attribute is a key. On remote routes, the address of the next sys- tem
    // en route; Otherwise, 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipforwardnexthop interface{}

    // Indicate the mask to be logical-ANDed with the destination  address  before
    // being compared to the value  in  the  ipForwardDest  field.   For those 
    // systems  that  do  not support arbitrary subnet masks, an agent constructs
    // the value  of the  ipForwardMask  by  reference to the IP Ad- dress Class. 
    // Any assignment (implicit or  otherwise)  of  an instance  of  this  object
    // to a value x must be rejected if the bitwise logical-AND of  x  with the 
    // value of the corresponding instance of the ipForwardDest object is not
    // equal to ipForward- Dest. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipforwardmask interface{}

    // The ifIndex value which identifies  the  local interface  through  which 
    // the next hop of this route should be reached. The type is interface{} with
    // range: -2147483648..2147483647.
    Ipforwardifindex interface{}

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
    // relevant  ip- ForwardType object. The type is Ipforwardtype.
    Ipforwardtype interface{}

    // The number of seconds  since  this  route  was last  updated  or  otherwise
    // determined  to be correct.  Note that no semantics of  `too  old' can  be
    // implied except through knowledge of the routing  protocol  by  which  the  
    // route   was learned. The type is interface{} with range:
    // -2147483648..2147483647.
    Ipforwardage interface{}

    // A reference to MIB definitions specific to the particular  routing protocol
    // which is responsi- ble for this route, as determined by the  value
    // specified  in the route's ipForwardProto value. If this information is not
    // present,  its  value should be set to the OBJECT IDENTIFIER { 0 0 }, which
    // is a syntactically valid object  identif- ier, and any implementation
    // conforming to ASN.1 and the Basic Encoding Rules must  be  able  to
    // generate and recognize this value. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Ipforwardinfo interface{}

    // The Autonomous System Number of the Next  Hop. When  this  is  unknown  or
    // not relevant to the protocol indicated by ipForwardProto, zero. The type is
    // interface{} with range: -2147483648..2147483647.
    Ipforwardnexthopas interface{}

    // The primary routing  metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipforwardmetric1 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipforwardmetric2 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipforwardmetric3 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipforwardmetric4 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipForwardProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipforwardmetric5 interface{}
}

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetEntityData() *types.CommonEntityData {
    ipforwardentry.EntityData.YFilter = ipforwardentry.YFilter
    ipforwardentry.EntityData.YangName = "ipForwardEntry"
    ipforwardentry.EntityData.BundleName = "cisco_ios_xe"
    ipforwardentry.EntityData.ParentYangName = "ipForwardTable"
    ipforwardentry.EntityData.SegmentPath = "ipForwardEntry" + "[ipForwardDest='" + fmt.Sprintf("%v", ipforwardentry.Ipforwarddest) + "']" + "[ipForwardProto='" + fmt.Sprintf("%v", ipforwardentry.Ipforwardproto) + "']" + "[ipForwardPolicy='" + fmt.Sprintf("%v", ipforwardentry.Ipforwardpolicy) + "']" + "[ipForwardNextHop='" + fmt.Sprintf("%v", ipforwardentry.Ipforwardnexthop) + "']"
    ipforwardentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipforwardentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipforwardentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipforwardentry.EntityData.Children = make(map[string]types.YChild)
    ipforwardentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipforwardentry.EntityData.Leafs["ipForwardDest"] = types.YLeaf{"Ipforwarddest", ipforwardentry.Ipforwarddest}
    ipforwardentry.EntityData.Leafs["ipForwardProto"] = types.YLeaf{"Ipforwardproto", ipforwardentry.Ipforwardproto}
    ipforwardentry.EntityData.Leafs["ipForwardPolicy"] = types.YLeaf{"Ipforwardpolicy", ipforwardentry.Ipforwardpolicy}
    ipforwardentry.EntityData.Leafs["ipForwardNextHop"] = types.YLeaf{"Ipforwardnexthop", ipforwardentry.Ipforwardnexthop}
    ipforwardentry.EntityData.Leafs["ipForwardMask"] = types.YLeaf{"Ipforwardmask", ipforwardentry.Ipforwardmask}
    ipforwardentry.EntityData.Leafs["ipForwardIfIndex"] = types.YLeaf{"Ipforwardifindex", ipforwardentry.Ipforwardifindex}
    ipforwardentry.EntityData.Leafs["ipForwardType"] = types.YLeaf{"Ipforwardtype", ipforwardentry.Ipforwardtype}
    ipforwardentry.EntityData.Leafs["ipForwardAge"] = types.YLeaf{"Ipforwardage", ipforwardentry.Ipforwardage}
    ipforwardentry.EntityData.Leafs["ipForwardInfo"] = types.YLeaf{"Ipforwardinfo", ipforwardentry.Ipforwardinfo}
    ipforwardentry.EntityData.Leafs["ipForwardNextHopAS"] = types.YLeaf{"Ipforwardnexthopas", ipforwardentry.Ipforwardnexthopas}
    ipforwardentry.EntityData.Leafs["ipForwardMetric1"] = types.YLeaf{"Ipforwardmetric1", ipforwardentry.Ipforwardmetric1}
    ipforwardentry.EntityData.Leafs["ipForwardMetric2"] = types.YLeaf{"Ipforwardmetric2", ipforwardentry.Ipforwardmetric2}
    ipforwardentry.EntityData.Leafs["ipForwardMetric3"] = types.YLeaf{"Ipforwardmetric3", ipforwardentry.Ipforwardmetric3}
    ipforwardentry.EntityData.Leafs["ipForwardMetric4"] = types.YLeaf{"Ipforwardmetric4", ipforwardentry.Ipforwardmetric4}
    ipforwardentry.EntityData.Leafs["ipForwardMetric5"] = types.YLeaf{"Ipforwardmetric5", ipforwardentry.Ipforwardmetric5}
    return &(ipforwardentry.EntityData)
}

// IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto represents hosts should support those protocols.
type IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto string

const (
    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_other IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "other"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_local IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "local"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_netmgmt IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "netmgmt"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_icmp IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "icmp"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_egp IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "egp"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_ggp IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "ggp"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_hello IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "hello"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_rip IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "rip"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_is_is IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "is-is"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_es_is IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "es-is"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_ciscoIgrp IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "ciscoIgrp"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_bbnSpfIgp IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "bbnSpfIgp"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_ospf IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "ospf"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_bgp IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "bgp"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto_idpr IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardproto = "idpr"
)

// IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype represents ForwardType object.
type IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype string

const (
    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype_other IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype = "other"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype_invalid IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype = "invalid"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype_local IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype = "local"

    IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype_remote IPFORWARDMIB_Ipforwardtable_Ipforwardentry_Ipforwardtype = "remote"
)

// IPFORWARDMIB_Ipcidrroutetable
// This entity's IP Routing table.
type IPFORWARDMIB_Ipcidrroutetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A particular route to  a  particular  destina- tion, under a particular
    // policy. The type is slice of
    // IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry.
    Ipcidrrouteentry []IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry
}

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetEntityData() *types.CommonEntityData {
    ipcidrroutetable.EntityData.YFilter = ipcidrroutetable.YFilter
    ipcidrroutetable.EntityData.YangName = "ipCidrRouteTable"
    ipcidrroutetable.EntityData.BundleName = "cisco_ios_xe"
    ipcidrroutetable.EntityData.ParentYangName = "IP-FORWARD-MIB"
    ipcidrroutetable.EntityData.SegmentPath = "ipCidrRouteTable"
    ipcidrroutetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipcidrroutetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipcidrroutetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipcidrroutetable.EntityData.Children = make(map[string]types.YChild)
    ipcidrroutetable.EntityData.Children["ipCidrRouteEntry"] = types.YChild{"Ipcidrrouteentry", nil}
    for i := range ipcidrroutetable.Ipcidrrouteentry {
        ipcidrroutetable.EntityData.Children[types.GetSegmentPath(&ipcidrroutetable.Ipcidrrouteentry[i])] = types.YChild{"Ipcidrrouteentry", &ipcidrroutetable.Ipcidrrouteentry[i]}
    }
    ipcidrroutetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipcidrroutetable.EntityData)
}

// IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry
// A particular route to  a  particular  destina-
// tion, under a particular policy.
type IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The destination IP address of this route.  This
    // object may not take a Multicast (Class  D) address value.  Any assignment
    // (implicit or  otherwise)  of  an instance  of  this  object to a value x
    // must be rejected if the bitwise logical-AND of  x  with the  value of the
    // corresponding instance of the ipCidrRouteMask object is not equal to x. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipcidrroutedest interface{}

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
    Ipcidrroutemask interface{}

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
    Ipcidrroutetos interface{}

    // This attribute is a key. On remote routes, the address of the next sys- tem
    // en route; Otherwise, 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipcidrroutenexthop interface{}

    // The ifIndex value which identifies  the  local interface  through  which 
    // the next hop of this route should be reached. The type is interface{} with
    // range: -2147483648..2147483647.
    Ipcidrrouteifindex interface{}

    // The type of route.  Note that local(3)  refers to  a route for which the
    // next hop is the final destination; remote(4) refers to  a  route  for which
    // the  next  hop is not the final destina- tion.  Routes which do not result
    // in traffic forwarding or rejection should not be displayed even if the
    // implementation keeps them stored internally.   reject (2) refers to a route
    // which, if matched, discards the message as unreachable. This is used in
    // some protocols as a means of correctly aggregating routes. The type is
    // Ipcidrroutetype.
    Ipcidrroutetype interface{}

    // The routing mechanism via which this route was learned.  Inclusion of
    // values for gateway rout- ing protocols is not  intended  to  imply  that
    // hosts should support those protocols. The type is Ipcidrrouteproto.
    Ipcidrrouteproto interface{}

    // The number of seconds  since  this  route  was last  updated  or  otherwise
    // determined  to be correct.  Note that no semantics of  `too  old' can  be
    // implied except through knowledge of the routing  protocol  by  which  the  
    // route   was learned. The type is interface{} with range:
    // -2147483648..2147483647.
    Ipcidrrouteage interface{}

    // A reference to MIB definitions specific to the particular  routing protocol
    // which is responsi- ble for this route, as determined by the  value
    // specified  in the route's ipCidrRouteProto value. If this information is
    // not present,  its  value should be set to the OBJECT IDENTIFIER { 0 0 },
    // which is a syntactically valid object  identif- ier, and any implementation
    // conforming to ASN.1 and the Basic Encoding Rules must  be  able  to
    // generate and recognize this value. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Ipcidrrouteinfo interface{}

    // The Autonomous System Number of the Next  Hop. The  semantics of this
    // object are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value. When  this object is unknown or not relevant its
    // value should be set to zero. The type is interface{} with range:
    // -2147483648..2147483647.
    Ipcidrroutenexthopas interface{}

    // The primary routing  metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipcidrroutemetric1 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipcidrroutemetric2 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipcidrroutemetric3 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipcidrroutemetric4 interface{}

    // An alternate routing metric  for  this  route. The  semantics of this
    // metric are determined by the routing-protocol specified in  the  route's
    // ipCidrRouteProto  value.   If  this metric is not used, its value should be
    // set to -1. The type is interface{} with range: -2147483648..2147483647.
    Ipcidrroutemetric5 interface{}

    // The row status variable, used according to row installation and removal
    // conventions. The type is RowStatus.
    Ipcidrroutestatus interface{}
}

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetEntityData() *types.CommonEntityData {
    ipcidrrouteentry.EntityData.YFilter = ipcidrrouteentry.YFilter
    ipcidrrouteentry.EntityData.YangName = "ipCidrRouteEntry"
    ipcidrrouteentry.EntityData.BundleName = "cisco_ios_xe"
    ipcidrrouteentry.EntityData.ParentYangName = "ipCidrRouteTable"
    ipcidrrouteentry.EntityData.SegmentPath = "ipCidrRouteEntry" + "[ipCidrRouteDest='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutedest) + "']" + "[ipCidrRouteMask='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutemask) + "']" + "[ipCidrRouteTos='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutetos) + "']" + "[ipCidrRouteNextHop='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutenexthop) + "']"
    ipcidrrouteentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipcidrrouteentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipcidrrouteentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipcidrrouteentry.EntityData.Children = make(map[string]types.YChild)
    ipcidrrouteentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteDest"] = types.YLeaf{"Ipcidrroutedest", ipcidrrouteentry.Ipcidrroutedest}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteMask"] = types.YLeaf{"Ipcidrroutemask", ipcidrrouteentry.Ipcidrroutemask}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteTos"] = types.YLeaf{"Ipcidrroutetos", ipcidrrouteentry.Ipcidrroutetos}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteNextHop"] = types.YLeaf{"Ipcidrroutenexthop", ipcidrrouteentry.Ipcidrroutenexthop}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteIfIndex"] = types.YLeaf{"Ipcidrrouteifindex", ipcidrrouteentry.Ipcidrrouteifindex}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteType"] = types.YLeaf{"Ipcidrroutetype", ipcidrrouteentry.Ipcidrroutetype}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteProto"] = types.YLeaf{"Ipcidrrouteproto", ipcidrrouteentry.Ipcidrrouteproto}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteAge"] = types.YLeaf{"Ipcidrrouteage", ipcidrrouteentry.Ipcidrrouteage}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteInfo"] = types.YLeaf{"Ipcidrrouteinfo", ipcidrrouteentry.Ipcidrrouteinfo}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteNextHopAS"] = types.YLeaf{"Ipcidrroutenexthopas", ipcidrrouteentry.Ipcidrroutenexthopas}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteMetric1"] = types.YLeaf{"Ipcidrroutemetric1", ipcidrrouteentry.Ipcidrroutemetric1}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteMetric2"] = types.YLeaf{"Ipcidrroutemetric2", ipcidrrouteentry.Ipcidrroutemetric2}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteMetric3"] = types.YLeaf{"Ipcidrroutemetric3", ipcidrrouteentry.Ipcidrroutemetric3}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteMetric4"] = types.YLeaf{"Ipcidrroutemetric4", ipcidrrouteentry.Ipcidrroutemetric4}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteMetric5"] = types.YLeaf{"Ipcidrroutemetric5", ipcidrrouteentry.Ipcidrroutemetric5}
    ipcidrrouteentry.EntityData.Leafs["ipCidrRouteStatus"] = types.YLeaf{"Ipcidrroutestatus", ipcidrrouteentry.Ipcidrroutestatus}
    return &(ipcidrrouteentry.EntityData)
}

// IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto represents hosts should support those protocols.
type IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto string

const (
    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_other IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "other"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_local IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "local"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_netmgmt IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "netmgmt"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_icmp IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "icmp"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_egp IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "egp"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_ggp IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "ggp"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_hello IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "hello"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_rip IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "rip"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_isIs IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "isIs"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_esIs IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "esIs"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_ciscoIgrp IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "ciscoIgrp"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_bbnSpfIgp IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "bbnSpfIgp"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_ospf IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "ospf"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_bgp IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "bgp"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_idpr IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "idpr"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto_ciscoEigrp IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrrouteproto = "ciscoEigrp"
)

// IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype represents protocols as a means of correctly aggregating routes.
type IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype string

const (
    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype_other IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype = "other"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype_reject IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype = "reject"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype_local IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype = "local"

    IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype_remote IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry_Ipcidrroutetype = "remote"
)

