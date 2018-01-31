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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ipforward IPFORWARDMIB_Ipforward

    // This entity's IP Routing table.
    Ipforwardtable IPFORWARDMIB_Ipforwardtable

    // This entity's IP Routing table.
    Ipcidrroutetable IPFORWARDMIB_Ipcidrroutetable
}

func (iPFORWARDMIB *IPFORWARDMIB) GetFilter() yfilter.YFilter { return iPFORWARDMIB.YFilter }

func (iPFORWARDMIB *IPFORWARDMIB) SetFilter(yf yfilter.YFilter) { iPFORWARDMIB.YFilter = yf }

func (iPFORWARDMIB *IPFORWARDMIB) GetGoName(yname string) string {
    if yname == "ipForward" { return "Ipforward" }
    if yname == "ipForwardTable" { return "Ipforwardtable" }
    if yname == "ipCidrRouteTable" { return "Ipcidrroutetable" }
    return ""
}

func (iPFORWARDMIB *IPFORWARDMIB) GetSegmentPath() string {
    return "IP-FORWARD-MIB:IP-FORWARD-MIB"
}

func (iPFORWARDMIB *IPFORWARDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipForward" {
        return &iPFORWARDMIB.Ipforward
    }
    if childYangName == "ipForwardTable" {
        return &iPFORWARDMIB.Ipforwardtable
    }
    if childYangName == "ipCidrRouteTable" {
        return &iPFORWARDMIB.Ipcidrroutetable
    }
    return nil
}

func (iPFORWARDMIB *IPFORWARDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipForward"] = &iPFORWARDMIB.Ipforward
    children["ipForwardTable"] = &iPFORWARDMIB.Ipforwardtable
    children["ipCidrRouteTable"] = &iPFORWARDMIB.Ipcidrroutetable
    return children
}

func (iPFORWARDMIB *IPFORWARDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iPFORWARDMIB *IPFORWARDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (iPFORWARDMIB *IPFORWARDMIB) GetYangName() string { return "IP-FORWARD-MIB" }

func (iPFORWARDMIB *IPFORWARDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iPFORWARDMIB *IPFORWARDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iPFORWARDMIB *IPFORWARDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iPFORWARDMIB *IPFORWARDMIB) SetParent(parent types.Entity) { iPFORWARDMIB.parent = parent }

func (iPFORWARDMIB *IPFORWARDMIB) GetParent() types.Entity { return iPFORWARDMIB.parent }

func (iPFORWARDMIB *IPFORWARDMIB) GetParentYangName() string { return "IP-FORWARD-MIB" }

// IPFORWARDMIB_Ipforward
type IPFORWARDMIB_Ipforward struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of current  ipForwardTable  entries that are not invalid. The
    // type is interface{} with range: 0..4294967295.
    Ipforwardnumber interface{}

    // The number of current ipCidrRouteTable entries that are not invalid. The
    // type is interface{} with range: 0..4294967295.
    Ipcidrroutenumber interface{}
}

func (ipforward *IPFORWARDMIB_Ipforward) GetFilter() yfilter.YFilter { return ipforward.YFilter }

func (ipforward *IPFORWARDMIB_Ipforward) SetFilter(yf yfilter.YFilter) { ipforward.YFilter = yf }

func (ipforward *IPFORWARDMIB_Ipforward) GetGoName(yname string) string {
    if yname == "ipForwardNumber" { return "Ipforwardnumber" }
    if yname == "ipCidrRouteNumber" { return "Ipcidrroutenumber" }
    return ""
}

func (ipforward *IPFORWARDMIB_Ipforward) GetSegmentPath() string {
    return "ipForward"
}

func (ipforward *IPFORWARDMIB_Ipforward) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipforward *IPFORWARDMIB_Ipforward) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipforward *IPFORWARDMIB_Ipforward) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipForwardNumber"] = ipforward.Ipforwardnumber
    leafs["ipCidrRouteNumber"] = ipforward.Ipcidrroutenumber
    return leafs
}

func (ipforward *IPFORWARDMIB_Ipforward) GetBundleName() string { return "cisco_ios_xe" }

func (ipforward *IPFORWARDMIB_Ipforward) GetYangName() string { return "ipForward" }

func (ipforward *IPFORWARDMIB_Ipforward) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipforward *IPFORWARDMIB_Ipforward) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipforward *IPFORWARDMIB_Ipforward) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipforward *IPFORWARDMIB_Ipforward) SetParent(parent types.Entity) { ipforward.parent = parent }

func (ipforward *IPFORWARDMIB_Ipforward) GetParent() types.Entity { return ipforward.parent }

func (ipforward *IPFORWARDMIB_Ipforward) GetParentYangName() string { return "IP-FORWARD-MIB" }

// IPFORWARDMIB_Ipforwardtable
// This entity's IP Routing table.
type IPFORWARDMIB_Ipforwardtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A particular route to  a  particular  destina- tion, under a particular
    // policy. The type is slice of IPFORWARDMIB_Ipforwardtable_Ipforwardentry.
    Ipforwardentry []IPFORWARDMIB_Ipforwardtable_Ipforwardentry
}

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetFilter() yfilter.YFilter { return ipforwardtable.YFilter }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) SetFilter(yf yfilter.YFilter) { ipforwardtable.YFilter = yf }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetGoName(yname string) string {
    if yname == "ipForwardEntry" { return "Ipforwardentry" }
    return ""
}

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetSegmentPath() string {
    return "ipForwardTable"
}

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipForwardEntry" {
        for _, c := range ipforwardtable.Ipforwardentry {
            if ipforwardtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPFORWARDMIB_Ipforwardtable_Ipforwardentry{}
        ipforwardtable.Ipforwardentry = append(ipforwardtable.Ipforwardentry, child)
        return &ipforwardtable.Ipforwardentry[len(ipforwardtable.Ipforwardentry)-1]
    }
    return nil
}

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipforwardtable.Ipforwardentry {
        children[ipforwardtable.Ipforwardentry[i].GetSegmentPath()] = &ipforwardtable.Ipforwardentry[i]
    }
    return children
}

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetBundleName() string { return "cisco_ios_xe" }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetYangName() string { return "ipForwardTable" }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) SetParent(parent types.Entity) { ipforwardtable.parent = parent }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetParent() types.Entity { return ipforwardtable.parent }

func (ipforwardtable *IPFORWARDMIB_Ipforwardtable) GetParentYangName() string { return "IP-FORWARD-MIB" }

// IPFORWARDMIB_Ipforwardtable_Ipforwardentry
// A particular route to  a  particular  destina-
// tion, under a particular policy.
type IPFORWARDMIB_Ipforwardtable_Ipforwardentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The destination IP address of this route.   An
    // entry  with  a value of 0.0.0.0 is considered a default route.  This object
    // may not take a Multicast (Class  D) address value.  Any assignment
    // (implicit or  otherwise)  of  an instance  of  this  object to a value x
    // must be rejected if the bitwise logical-AND of  x  with the  value of the
    // corresponding instance of the ipForwardMask object is not equal to x. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipforwardnexthop interface{}

    // Indicate the mask to be logical-ANDed with the destination  address  before
    // being compared to the value  in  the  ipForwardDest  field.   For those 
    // systems  that  do  not support arbitrary subnet masks, an agent constructs
    // the value  of the  ipForwardMask  by  reference to the IP Ad- dress Class. 
    // Any assignment (implicit or  otherwise)  of  an instance  of  this  object
    // to a value x must be rejected if the bitwise logical-AND of  x  with the 
    // value of the corresponding instance of the ipForwardDest object is not
    // equal to ipForward- Dest. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetFilter() yfilter.YFilter { return ipforwardentry.YFilter }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) SetFilter(yf yfilter.YFilter) { ipforwardentry.YFilter = yf }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetGoName(yname string) string {
    if yname == "ipForwardDest" { return "Ipforwarddest" }
    if yname == "ipForwardProto" { return "Ipforwardproto" }
    if yname == "ipForwardPolicy" { return "Ipforwardpolicy" }
    if yname == "ipForwardNextHop" { return "Ipforwardnexthop" }
    if yname == "ipForwardMask" { return "Ipforwardmask" }
    if yname == "ipForwardIfIndex" { return "Ipforwardifindex" }
    if yname == "ipForwardType" { return "Ipforwardtype" }
    if yname == "ipForwardAge" { return "Ipforwardage" }
    if yname == "ipForwardInfo" { return "Ipforwardinfo" }
    if yname == "ipForwardNextHopAS" { return "Ipforwardnexthopas" }
    if yname == "ipForwardMetric1" { return "Ipforwardmetric1" }
    if yname == "ipForwardMetric2" { return "Ipforwardmetric2" }
    if yname == "ipForwardMetric3" { return "Ipforwardmetric3" }
    if yname == "ipForwardMetric4" { return "Ipforwardmetric4" }
    if yname == "ipForwardMetric5" { return "Ipforwardmetric5" }
    return ""
}

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetSegmentPath() string {
    return "ipForwardEntry" + "[ipForwardDest='" + fmt.Sprintf("%v", ipforwardentry.Ipforwarddest) + "']" + "[ipForwardProto='" + fmt.Sprintf("%v", ipforwardentry.Ipforwardproto) + "']" + "[ipForwardPolicy='" + fmt.Sprintf("%v", ipforwardentry.Ipforwardpolicy) + "']" + "[ipForwardNextHop='" + fmt.Sprintf("%v", ipforwardentry.Ipforwardnexthop) + "']"
}

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipForwardDest"] = ipforwardentry.Ipforwarddest
    leafs["ipForwardProto"] = ipforwardentry.Ipforwardproto
    leafs["ipForwardPolicy"] = ipforwardentry.Ipforwardpolicy
    leafs["ipForwardNextHop"] = ipforwardentry.Ipforwardnexthop
    leafs["ipForwardMask"] = ipforwardentry.Ipforwardmask
    leafs["ipForwardIfIndex"] = ipforwardentry.Ipforwardifindex
    leafs["ipForwardType"] = ipforwardentry.Ipforwardtype
    leafs["ipForwardAge"] = ipforwardentry.Ipforwardage
    leafs["ipForwardInfo"] = ipforwardentry.Ipforwardinfo
    leafs["ipForwardNextHopAS"] = ipforwardentry.Ipforwardnexthopas
    leafs["ipForwardMetric1"] = ipforwardentry.Ipforwardmetric1
    leafs["ipForwardMetric2"] = ipforwardentry.Ipforwardmetric2
    leafs["ipForwardMetric3"] = ipforwardentry.Ipforwardmetric3
    leafs["ipForwardMetric4"] = ipforwardentry.Ipforwardmetric4
    leafs["ipForwardMetric5"] = ipforwardentry.Ipforwardmetric5
    return leafs
}

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetYangName() string { return "ipForwardEntry" }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) SetParent(parent types.Entity) { ipforwardentry.parent = parent }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetParent() types.Entity { return ipforwardentry.parent }

func (ipforwardentry *IPFORWARDMIB_Ipforwardtable_Ipforwardentry) GetParentYangName() string { return "ipForwardTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A particular route to  a  particular  destina- tion, under a particular
    // policy. The type is slice of
    // IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry.
    Ipcidrrouteentry []IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry
}

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetFilter() yfilter.YFilter { return ipcidrroutetable.YFilter }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) SetFilter(yf yfilter.YFilter) { ipcidrroutetable.YFilter = yf }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetGoName(yname string) string {
    if yname == "ipCidrRouteEntry" { return "Ipcidrrouteentry" }
    return ""
}

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetSegmentPath() string {
    return "ipCidrRouteTable"
}

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipCidrRouteEntry" {
        for _, c := range ipcidrroutetable.Ipcidrrouteentry {
            if ipcidrroutetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry{}
        ipcidrroutetable.Ipcidrrouteentry = append(ipcidrroutetable.Ipcidrrouteentry, child)
        return &ipcidrroutetable.Ipcidrrouteentry[len(ipcidrroutetable.Ipcidrrouteentry)-1]
    }
    return nil
}

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipcidrroutetable.Ipcidrrouteentry {
        children[ipcidrroutetable.Ipcidrrouteentry[i].GetSegmentPath()] = &ipcidrroutetable.Ipcidrrouteentry[i]
    }
    return children
}

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetBundleName() string { return "cisco_ios_xe" }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetYangName() string { return "ipCidrRouteTable" }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) SetParent(parent types.Entity) { ipcidrroutetable.parent = parent }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetParent() types.Entity { return ipcidrroutetable.parent }

func (ipcidrroutetable *IPFORWARDMIB_Ipcidrroutetable) GetParentYangName() string { return "IP-FORWARD-MIB" }

// IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry
// A particular route to  a  particular  destina-
// tion, under a particular policy.
type IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The destination IP address of this route.  This
    // object may not take a Multicast (Class  D) address value.  Any assignment
    // (implicit or  otherwise)  of  an instance  of  this  object to a value x
    // must be rejected if the bitwise logical-AND of  x  with the  value of the
    // corresponding instance of the ipCidrRouteMask object is not equal to x. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetFilter() yfilter.YFilter { return ipcidrrouteentry.YFilter }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) SetFilter(yf yfilter.YFilter) { ipcidrrouteentry.YFilter = yf }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetGoName(yname string) string {
    if yname == "ipCidrRouteDest" { return "Ipcidrroutedest" }
    if yname == "ipCidrRouteMask" { return "Ipcidrroutemask" }
    if yname == "ipCidrRouteTos" { return "Ipcidrroutetos" }
    if yname == "ipCidrRouteNextHop" { return "Ipcidrroutenexthop" }
    if yname == "ipCidrRouteIfIndex" { return "Ipcidrrouteifindex" }
    if yname == "ipCidrRouteType" { return "Ipcidrroutetype" }
    if yname == "ipCidrRouteProto" { return "Ipcidrrouteproto" }
    if yname == "ipCidrRouteAge" { return "Ipcidrrouteage" }
    if yname == "ipCidrRouteInfo" { return "Ipcidrrouteinfo" }
    if yname == "ipCidrRouteNextHopAS" { return "Ipcidrroutenexthopas" }
    if yname == "ipCidrRouteMetric1" { return "Ipcidrroutemetric1" }
    if yname == "ipCidrRouteMetric2" { return "Ipcidrroutemetric2" }
    if yname == "ipCidrRouteMetric3" { return "Ipcidrroutemetric3" }
    if yname == "ipCidrRouteMetric4" { return "Ipcidrroutemetric4" }
    if yname == "ipCidrRouteMetric5" { return "Ipcidrroutemetric5" }
    if yname == "ipCidrRouteStatus" { return "Ipcidrroutestatus" }
    return ""
}

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetSegmentPath() string {
    return "ipCidrRouteEntry" + "[ipCidrRouteDest='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutedest) + "']" + "[ipCidrRouteMask='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutemask) + "']" + "[ipCidrRouteTos='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutetos) + "']" + "[ipCidrRouteNextHop='" + fmt.Sprintf("%v", ipcidrrouteentry.Ipcidrroutenexthop) + "']"
}

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipCidrRouteDest"] = ipcidrrouteentry.Ipcidrroutedest
    leafs["ipCidrRouteMask"] = ipcidrrouteentry.Ipcidrroutemask
    leafs["ipCidrRouteTos"] = ipcidrrouteentry.Ipcidrroutetos
    leafs["ipCidrRouteNextHop"] = ipcidrrouteentry.Ipcidrroutenexthop
    leafs["ipCidrRouteIfIndex"] = ipcidrrouteentry.Ipcidrrouteifindex
    leafs["ipCidrRouteType"] = ipcidrrouteentry.Ipcidrroutetype
    leafs["ipCidrRouteProto"] = ipcidrrouteentry.Ipcidrrouteproto
    leafs["ipCidrRouteAge"] = ipcidrrouteentry.Ipcidrrouteage
    leafs["ipCidrRouteInfo"] = ipcidrrouteentry.Ipcidrrouteinfo
    leafs["ipCidrRouteNextHopAS"] = ipcidrrouteentry.Ipcidrroutenexthopas
    leafs["ipCidrRouteMetric1"] = ipcidrrouteentry.Ipcidrroutemetric1
    leafs["ipCidrRouteMetric2"] = ipcidrrouteentry.Ipcidrroutemetric2
    leafs["ipCidrRouteMetric3"] = ipcidrrouteentry.Ipcidrroutemetric3
    leafs["ipCidrRouteMetric4"] = ipcidrrouteentry.Ipcidrroutemetric4
    leafs["ipCidrRouteMetric5"] = ipcidrrouteentry.Ipcidrroutemetric5
    leafs["ipCidrRouteStatus"] = ipcidrrouteentry.Ipcidrroutestatus
    return leafs
}

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetYangName() string { return "ipCidrRouteEntry" }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) SetParent(parent types.Entity) { ipcidrrouteentry.parent = parent }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetParent() types.Entity { return ipcidrrouteentry.parent }

func (ipcidrrouteentry *IPFORWARDMIB_Ipcidrroutetable_Ipcidrrouteentry) GetParentYangName() string { return "ipCidrRouteTable" }

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

