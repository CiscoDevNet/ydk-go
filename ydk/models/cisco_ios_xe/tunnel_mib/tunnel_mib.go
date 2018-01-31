// The MIB module for management of IP Tunnels,
// independent of the specific encapsulation scheme in
// use.
// 
// Copyright (C) The Internet Society (2005).  This
// version of this MIB module is part of RFC 4087;  see
// the RFC itself for full legal notices.
package tunnel_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:TUNNEL-MIB TUNNEL-MIB}", reflect.TypeOf(TUNNELMIB{}))
    ydk.RegisterEntity("TUNNEL-MIB:TUNNEL-MIB", reflect.TypeOf(TUNNELMIB{}))
}

// TUNNELMIB
type TUNNELMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The (conceptual) table containing information on configured tunnels.
    Tunneliftable TUNNELMIB_Tunneliftable

    // The (conceptual) table containing information on configured tunnels.  This
    // table can be used to map a set of tunnel endpoints to the associated
    // ifIndex value.  It can also be used for row creation.  Note that every row
    // in the tunnelIfTable with a fixed IPv4 destination address should have a
    // corresponding row in the tunnelConfigTable, regardless of whether it was
    // created via SNMP.  Since this table does not support IPv6, it is deprecated
    // in favor of tunnelInetConfigTable.
    Tunnelconfigtable TUNNELMIB_Tunnelconfigtable

    // The (conceptual) table containing information on configured tunnels.  This
    // table can be used to map a set of tunnel endpoints to the associated
    // ifIndex value.  It can also be used for row creation.  Note that every row
    // in the tunnelIfTable with a fixed destination address should have a
    // corresponding row in the tunnelInetConfigTable, regardless of whether it
    // was created via SNMP.
    Tunnelinetconfigtable TUNNELMIB_Tunnelinetconfigtable
}

func (tUNNELMIB *TUNNELMIB) GetFilter() yfilter.YFilter { return tUNNELMIB.YFilter }

func (tUNNELMIB *TUNNELMIB) SetFilter(yf yfilter.YFilter) { tUNNELMIB.YFilter = yf }

func (tUNNELMIB *TUNNELMIB) GetGoName(yname string) string {
    if yname == "tunnelIfTable" { return "Tunneliftable" }
    if yname == "tunnelConfigTable" { return "Tunnelconfigtable" }
    if yname == "tunnelInetConfigTable" { return "Tunnelinetconfigtable" }
    return ""
}

func (tUNNELMIB *TUNNELMIB) GetSegmentPath() string {
    return "TUNNEL-MIB:TUNNEL-MIB"
}

func (tUNNELMIB *TUNNELMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnelIfTable" {
        return &tUNNELMIB.Tunneliftable
    }
    if childYangName == "tunnelConfigTable" {
        return &tUNNELMIB.Tunnelconfigtable
    }
    if childYangName == "tunnelInetConfigTable" {
        return &tUNNELMIB.Tunnelinetconfigtable
    }
    return nil
}

func (tUNNELMIB *TUNNELMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tunnelIfTable"] = &tUNNELMIB.Tunneliftable
    children["tunnelConfigTable"] = &tUNNELMIB.Tunnelconfigtable
    children["tunnelInetConfigTable"] = &tUNNELMIB.Tunnelinetconfigtable
    return children
}

func (tUNNELMIB *TUNNELMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tUNNELMIB *TUNNELMIB) GetBundleName() string { return "cisco_ios_xe" }

func (tUNNELMIB *TUNNELMIB) GetYangName() string { return "TUNNEL-MIB" }

func (tUNNELMIB *TUNNELMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tUNNELMIB *TUNNELMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tUNNELMIB *TUNNELMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tUNNELMIB *TUNNELMIB) SetParent(parent types.Entity) { tUNNELMIB.parent = parent }

func (tUNNELMIB *TUNNELMIB) GetParent() types.Entity { return tUNNELMIB.parent }

func (tUNNELMIB *TUNNELMIB) GetParentYangName() string { return "TUNNEL-MIB" }

// TUNNELMIB_Tunneliftable
// The (conceptual) table containing information on
// configured tunnels.
type TUNNELMIB_Tunneliftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the information on a particular
    // configured tunnel. The type is slice of
    // TUNNELMIB_Tunneliftable_Tunnelifentry.
    Tunnelifentry []TUNNELMIB_Tunneliftable_Tunnelifentry
}

func (tunneliftable *TUNNELMIB_Tunneliftable) GetFilter() yfilter.YFilter { return tunneliftable.YFilter }

func (tunneliftable *TUNNELMIB_Tunneliftable) SetFilter(yf yfilter.YFilter) { tunneliftable.YFilter = yf }

func (tunneliftable *TUNNELMIB_Tunneliftable) GetGoName(yname string) string {
    if yname == "tunnelIfEntry" { return "Tunnelifentry" }
    return ""
}

func (tunneliftable *TUNNELMIB_Tunneliftable) GetSegmentPath() string {
    return "tunnelIfTable"
}

func (tunneliftable *TUNNELMIB_Tunneliftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnelIfEntry" {
        for _, c := range tunneliftable.Tunnelifentry {
            if tunneliftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TUNNELMIB_Tunneliftable_Tunnelifentry{}
        tunneliftable.Tunnelifentry = append(tunneliftable.Tunnelifentry, child)
        return &tunneliftable.Tunnelifentry[len(tunneliftable.Tunnelifentry)-1]
    }
    return nil
}

func (tunneliftable *TUNNELMIB_Tunneliftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunneliftable.Tunnelifentry {
        children[tunneliftable.Tunnelifentry[i].GetSegmentPath()] = &tunneliftable.Tunnelifentry[i]
    }
    return children
}

func (tunneliftable *TUNNELMIB_Tunneliftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunneliftable *TUNNELMIB_Tunneliftable) GetBundleName() string { return "cisco_ios_xe" }

func (tunneliftable *TUNNELMIB_Tunneliftable) GetYangName() string { return "tunnelIfTable" }

func (tunneliftable *TUNNELMIB_Tunneliftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tunneliftable *TUNNELMIB_Tunneliftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tunneliftable *TUNNELMIB_Tunneliftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tunneliftable *TUNNELMIB_Tunneliftable) SetParent(parent types.Entity) { tunneliftable.parent = parent }

func (tunneliftable *TUNNELMIB_Tunneliftable) GetParent() types.Entity { return tunneliftable.parent }

func (tunneliftable *TUNNELMIB_Tunneliftable) GetParentYangName() string { return "TUNNEL-MIB" }

// TUNNELMIB_Tunneliftable_Tunnelifentry
// An entry (conceptual row) containing the information
// on a particular configured tunnel.
type TUNNELMIB_Tunneliftable_Tunnelifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The address of the local endpoint of the tunnel (i.e., the source address
    // used in the outer IP header), or 0.0.0.0 if unknown or if the tunnel is
    // over IPv6.  Since this object does not support IPv6, it is deprecated in
    // favor of tunnelIfLocalInetAddress. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Tunneliflocaladdress interface{}

    // The address of the remote endpoint of the tunnel (i.e., the destination
    // address used in the outer IP header), or 0.0.0.0 if unknown, or an IPv6
    // address, or  the tunnel is not a point-to-point link (e.g., if it is a 6to4
    // tunnel).  Since this object does not support IPv6, it is deprecated in
    // favor of tunnelIfRemoteInetAddress. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Tunnelifremoteaddress interface{}

    // The encapsulation method used by the tunnel. The type is IANAtunnelType.
    Tunnelifencapsmethod interface{}

    // The IPv4 TTL or IPv6 Hop Limit to use in the outer IP header.  A value of 0
    // indicates that the value is copied from the payload's header. The type is
    // interface{} with range: 0..255.
    Tunnelifhoplimit interface{}

    // The method used by the tunnel to secure the outer IP header.  The value
    // ipsec indicates that IPsec is used between the tunnel endpoints for
    // authentication or encryption or both.  More specific security-related
    // information may be available in a MIB module for the security protocol in
    // use. The type is Tunnelifsecurity.
    Tunnelifsecurity interface{}

    // The method used to set the high 6 bits (the  differentiated services
    // codepoint) of the IPv4 TOS or IPv6 Traffic Class in the outer IP header.  A
    // value of -1 indicates that the bits are copied from the payload's header. 
    // A value of -2 indicates that a traffic conditioner is invoked and more
    // information may be available in a traffic conditioner MIB module. A value
    // between 0 and 63 inclusive indicates that the bit field is set to the
    // indicated value.  Note: instead of the name tunnelIfTOS, a better name
    // would have been tunnelIfDSCPMethod, but the existing name appeared in RFC
    // 2667 and existing objects cannot be renamed. The type is interface{} with
    // range: -2..63.
    Tunneliftos interface{}

    // The method used to set the IPv6 Flow Label value. This object need not be
    // present in rows where tunnelIfAddressType indicates the tunnel is not over
    // IPv6.  A value of -1 indicates that a traffic conditioner is invoked and
    // more information may be available in a traffic conditioner MIB.  Any other
    // value indicates that the Flow Label field is set to the indicated value.
    // The type is interface{} with range: -1..100.
    Tunnelifflowlabel interface{}

    // The type of address in the corresponding tunnelIfLocalInetAddress and
    // tunnelIfRemoteInetAddress objects. The type is InetAddressType.
    Tunnelifaddresstype interface{}

    // The address of the local endpoint of the tunnel (i.e., the source address
    // used in the outer IP header).  If the address is unknown, the value is 
    // 0.0.0.0 for IPv4 or :: for IPv6.  The type of this object is given by
    // tunnelIfAddressType. The type is string with length: 0..255.
    Tunneliflocalinetaddress interface{}

    // The address of the remote endpoint of the tunnel (i.e., the destination
    // address used in the outer IP header).  If the address is unknown or the
    // tunnel is not a point-to-point link (e.g., if it is a 6to4 tunnel), the
    // value is 0.0.0.0 for tunnels over IPv4 or :: for tunnels over IPv6.  The
    // type of this object is given by tunnelIfAddressType. The type is string
    // with length: 0..255.
    Tunnelifremoteinetaddress interface{}

    // The maximum number of additional encapsulations permitted for packets
    // undergoing encapsulation at this node.  A value of -1 indicates that no
    // limit is present (except as a result of the packet size). The type is
    // interface{} with range: -1..255.
    Tunnelifencapslimit interface{}
}

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetFilter() yfilter.YFilter { return tunnelifentry.YFilter }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) SetFilter(yf yfilter.YFilter) { tunnelifentry.YFilter = yf }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "tunnelIfLocalAddress" { return "Tunneliflocaladdress" }
    if yname == "tunnelIfRemoteAddress" { return "Tunnelifremoteaddress" }
    if yname == "tunnelIfEncapsMethod" { return "Tunnelifencapsmethod" }
    if yname == "tunnelIfHopLimit" { return "Tunnelifhoplimit" }
    if yname == "tunnelIfSecurity" { return "Tunnelifsecurity" }
    if yname == "tunnelIfTOS" { return "Tunneliftos" }
    if yname == "tunnelIfFlowLabel" { return "Tunnelifflowlabel" }
    if yname == "tunnelIfAddressType" { return "Tunnelifaddresstype" }
    if yname == "tunnelIfLocalInetAddress" { return "Tunneliflocalinetaddress" }
    if yname == "tunnelIfRemoteInetAddress" { return "Tunnelifremoteinetaddress" }
    if yname == "tunnelIfEncapsLimit" { return "Tunnelifencapslimit" }
    return ""
}

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetSegmentPath() string {
    return "tunnelIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", tunnelifentry.Ifindex) + "']"
}

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = tunnelifentry.Ifindex
    leafs["tunnelIfLocalAddress"] = tunnelifentry.Tunneliflocaladdress
    leafs["tunnelIfRemoteAddress"] = tunnelifentry.Tunnelifremoteaddress
    leafs["tunnelIfEncapsMethod"] = tunnelifentry.Tunnelifencapsmethod
    leafs["tunnelIfHopLimit"] = tunnelifentry.Tunnelifhoplimit
    leafs["tunnelIfSecurity"] = tunnelifentry.Tunnelifsecurity
    leafs["tunnelIfTOS"] = tunnelifentry.Tunneliftos
    leafs["tunnelIfFlowLabel"] = tunnelifentry.Tunnelifflowlabel
    leafs["tunnelIfAddressType"] = tunnelifentry.Tunnelifaddresstype
    leafs["tunnelIfLocalInetAddress"] = tunnelifentry.Tunneliflocalinetaddress
    leafs["tunnelIfRemoteInetAddress"] = tunnelifentry.Tunnelifremoteinetaddress
    leafs["tunnelIfEncapsLimit"] = tunnelifentry.Tunnelifencapslimit
    return leafs
}

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetBundleName() string { return "cisco_ios_xe" }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetYangName() string { return "tunnelIfEntry" }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) SetParent(parent types.Entity) { tunnelifentry.parent = parent }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetParent() types.Entity { return tunnelifentry.parent }

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetParentYangName() string { return "tunnelIfTable" }

// TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity represents security protocol in use.
type TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity string

const (
    TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity_none TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity = "none"

    TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity_ipsec TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity = "ipsec"

    TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity_other TUNNELMIB_Tunneliftable_Tunnelifentry_Tunnelifsecurity = "other"
)

// TUNNELMIB_Tunnelconfigtable
// The (conceptual) table containing information on
// configured tunnels.  This table can be used to map a
// set of tunnel endpoints to the associated ifIndex
// value.  It can also be used for row creation.  Note
// that every row in the tunnelIfTable with a fixed IPv4
// destination address should have a corresponding row in
// the tunnelConfigTable, regardless of whether it was
// created via SNMP.
// 
// Since this table does not support IPv6, it is
// deprecated in favor of tunnelInetConfigTable.
type TUNNELMIB_Tunnelconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the information on a particular
    // configured tunnel.  Since this entry does not support IPv6, it is
    // deprecated in favor of tunnelInetConfigEntry. The type is slice of
    // TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry.
    Tunnelconfigentry []TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry
}

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetFilter() yfilter.YFilter { return tunnelconfigtable.YFilter }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) SetFilter(yf yfilter.YFilter) { tunnelconfigtable.YFilter = yf }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetGoName(yname string) string {
    if yname == "tunnelConfigEntry" { return "Tunnelconfigentry" }
    return ""
}

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetSegmentPath() string {
    return "tunnelConfigTable"
}

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnelConfigEntry" {
        for _, c := range tunnelconfigtable.Tunnelconfigentry {
            if tunnelconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry{}
        tunnelconfigtable.Tunnelconfigentry = append(tunnelconfigtable.Tunnelconfigentry, child)
        return &tunnelconfigtable.Tunnelconfigentry[len(tunnelconfigtable.Tunnelconfigentry)-1]
    }
    return nil
}

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelconfigtable.Tunnelconfigentry {
        children[tunnelconfigtable.Tunnelconfigentry[i].GetSegmentPath()] = &tunnelconfigtable.Tunnelconfigentry[i]
    }
    return children
}

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetYangName() string { return "tunnelConfigTable" }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) SetParent(parent types.Entity) { tunnelconfigtable.parent = parent }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetParent() types.Entity { return tunnelconfigtable.parent }

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetParentYangName() string { return "TUNNEL-MIB" }

// TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry
// An entry (conceptual row) containing the information
// on a particular configured tunnel.
// 
// Since this entry does not support IPv6, it is
// deprecated in favor of tunnelInetConfigEntry.
type TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The address of the local endpoint of the tunnel,
    // or 0.0.0.0 if the device is free to choose any of its addresses at tunnel
    // establishment time.  Since this object does not support IPv6, it is
    // deprecated in favor of tunnelInetConfigLocalAddress. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Tunnelconfiglocaladdress interface{}

    // This attribute is a key. The address of the remote endpoint of the tunnel. 
    // Since this object does not support IPv6, it is deprecated in favor of
    // tunnelInetConfigRemoteAddress. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Tunnelconfigremoteaddress interface{}

    // This attribute is a key. The encapsulation method used by the tunnel. 
    // Since this object does not support IPv6, it is deprecated in favor of
    // tunnelInetConfigEncapsMethod. The type is IANAtunnelType.
    Tunnelconfigencapsmethod interface{}

    // This attribute is a key. An identifier used to distinguish between multiple
    // tunnels of the same encapsulation method, with the same endpoints.  If the
    // encapsulation protocol only allows one tunnel per set of endpoint addresses
    // (such as for GRE or IP-in-IP), the value of this object is 1.  For
    // encapsulation methods (such as L2F) which allow multiple parallel tunnels,
    // the manager is responsible for choosing any ID which does not conflict with
    // an existing row, such as choosing a random number.  Since this object does
    // not support IPv6, it is deprecated in favor of tunnelInetConfigID. The type
    // is interface{} with range: 1..2147483647.
    Tunnelconfigid interface{}

    // If the value of tunnelConfigStatus for this row is active, then this object
    // contains the value of ifIndex corresponding to the tunnel interface.  A
    // value of 0 is not legal in the active state, and means that the interface
    // index has not yet been assigned.  Since this object does not support IPv6,
    // it is deprecated in favor of tunnelInetConfigIfIndex. The type is
    // interface{} with range: 0..2147483647.
    Tunnelconfigifindex interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table.  The agent need not support setting this object to
    // createAndWait or notInService since there are no other writable objects in
    // this table, and writable objects in rows of corresponding tables such as
    // the tunnelIfTable may be modified while this row is active.  To create a
    // row in this table for an encapsulation method which does not support
    // multiple parallel tunnels with the same endpoints, the management station
    // should simply use a tunnelConfigID of 1, and set tunnelConfigStatus to
    // createAndGo.  For encapsulation methods such as L2F which allow multiple
    // parallel tunnels, the management station may select a pseudo-random number
    // to use as the tunnelConfigID and set tunnelConfigStatus to createAndGo.  In
    // the event that this ID is already in use and an inconsistentValue is
    // returned in response to the set operation, the management station should
    // simply select a new pseudo-random number and retry the operation.  Creating
    // a row in this table will cause an interface index to be assigned by the
    // agent in an implementation-dependent manner, and corresponding rows will be
    // instantiated in the ifTable and the tunnelIfTable.  The status of this row
    // will become active as soon as the agent assigns the interface index,
    // regardless of whether the interface is operationally up.  Deleting a row in
    // this table will likewise delete the corresponding row in the ifTable and in
    // the tunnelIfTable.  Since this object does not support IPv6, it is
    // deprecated in favor of tunnelInetConfigStatus. The type is RowStatus.
    Tunnelconfigstatus interface{}
}

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetFilter() yfilter.YFilter { return tunnelconfigentry.YFilter }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) SetFilter(yf yfilter.YFilter) { tunnelconfigentry.YFilter = yf }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetGoName(yname string) string {
    if yname == "tunnelConfigLocalAddress" { return "Tunnelconfiglocaladdress" }
    if yname == "tunnelConfigRemoteAddress" { return "Tunnelconfigremoteaddress" }
    if yname == "tunnelConfigEncapsMethod" { return "Tunnelconfigencapsmethod" }
    if yname == "tunnelConfigID" { return "Tunnelconfigid" }
    if yname == "tunnelConfigIfIndex" { return "Tunnelconfigifindex" }
    if yname == "tunnelConfigStatus" { return "Tunnelconfigstatus" }
    return ""
}

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetSegmentPath() string {
    return "tunnelConfigEntry" + "[tunnelConfigLocalAddress='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfiglocaladdress) + "']" + "[tunnelConfigRemoteAddress='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfigremoteaddress) + "']" + "[tunnelConfigEncapsMethod='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfigencapsmethod) + "']" + "[tunnelConfigID='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfigid) + "']"
}

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnelConfigLocalAddress"] = tunnelconfigentry.Tunnelconfiglocaladdress
    leafs["tunnelConfigRemoteAddress"] = tunnelconfigentry.Tunnelconfigremoteaddress
    leafs["tunnelConfigEncapsMethod"] = tunnelconfigentry.Tunnelconfigencapsmethod
    leafs["tunnelConfigID"] = tunnelconfigentry.Tunnelconfigid
    leafs["tunnelConfigIfIndex"] = tunnelconfigentry.Tunnelconfigifindex
    leafs["tunnelConfigStatus"] = tunnelconfigentry.Tunnelconfigstatus
    return leafs
}

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetYangName() string { return "tunnelConfigEntry" }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) SetParent(parent types.Entity) { tunnelconfigentry.parent = parent }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetParent() types.Entity { return tunnelconfigentry.parent }

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetParentYangName() string { return "tunnelConfigTable" }

// TUNNELMIB_Tunnelinetconfigtable
// The (conceptual) table containing information on
// configured tunnels.  This table can be used to map a
// set of tunnel endpoints to the associated ifIndex
// value.  It can also be used for row creation.  Note
// that every row in the tunnelIfTable with a fixed
// destination address should have a corresponding row in
// the tunnelInetConfigTable, regardless of whether it
// was created via SNMP.
type TUNNELMIB_Tunnelinetconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the information on a particular
    // configured tunnel.  Note that there is a 128 subid maximum for object OIDs.
    // Implementers need to be aware that if the total number of octets in
    // tunnelInetConfigLocalAddress and tunnelInetConfigRemoteAddress exceeds 110
    // then OIDs of column instances in this table will have more than 128
    // sub-identifiers and cannot be accessed using SNMPv1, SNMPv2c, or SNMPv3. 
    // In practice this is not expected to be a problem since IPv4 and IPv6
    // addresses will not cause the limit to be reached, but if other types are
    // supported by an agent, care must be taken to ensure that the sum of the
    // lengths do not cause the limit to be exceeded. The type is slice of
    // TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry.
    Tunnelinetconfigentry []TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry
}

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetFilter() yfilter.YFilter { return tunnelinetconfigtable.YFilter }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) SetFilter(yf yfilter.YFilter) { tunnelinetconfigtable.YFilter = yf }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetGoName(yname string) string {
    if yname == "tunnelInetConfigEntry" { return "Tunnelinetconfigentry" }
    return ""
}

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetSegmentPath() string {
    return "tunnelInetConfigTable"
}

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnelInetConfigEntry" {
        for _, c := range tunnelinetconfigtable.Tunnelinetconfigentry {
            if tunnelinetconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry{}
        tunnelinetconfigtable.Tunnelinetconfigentry = append(tunnelinetconfigtable.Tunnelinetconfigentry, child)
        return &tunnelinetconfigtable.Tunnelinetconfigentry[len(tunnelinetconfigtable.Tunnelinetconfigentry)-1]
    }
    return nil
}

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelinetconfigtable.Tunnelinetconfigentry {
        children[tunnelinetconfigtable.Tunnelinetconfigentry[i].GetSegmentPath()] = &tunnelinetconfigtable.Tunnelinetconfigentry[i]
    }
    return children
}

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetYangName() string { return "tunnelInetConfigTable" }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) SetParent(parent types.Entity) { tunnelinetconfigtable.parent = parent }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetParent() types.Entity { return tunnelinetconfigtable.parent }

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetParentYangName() string { return "TUNNEL-MIB" }

// TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry
// An entry (conceptual row) containing the information
// on a particular configured tunnel.  Note that there is
// a 128 subid maximum for object OIDs.  Implementers
// need to be aware that if the total number of octets in
// tunnelInetConfigLocalAddress and
// tunnelInetConfigRemoteAddress exceeds 110 then OIDs of
// column instances in this table will have more than 128
// sub-identifiers and cannot be accessed using SNMPv1,
// SNMPv2c, or SNMPv3.  In practice this is not expected
// to be a problem since IPv4 and IPv6 addresses will not
// cause the limit to be reached, but if other types are
// supported by an agent, care must be taken to ensure
// that the sum of the lengths do not cause the limit to
// be exceeded.
type TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The address type over which the tunnel
    // encapsulates packets. The type is InetAddressType.
    Tunnelinetconfigaddresstype interface{}

    // This attribute is a key. The address of the local endpoint of the tunnel,
    // or 0.0.0.0 (for IPv4) or :: (for IPv6) if the device is free to choose any
    // of its addresses at tunnel establishment time. The type is string with
    // length: 0..255.
    Tunnelinetconfiglocaladdress interface{}

    // This attribute is a key. The address of the remote endpoint of the tunnel.
    // The type is string with length: 0..255.
    Tunnelinetconfigremoteaddress interface{}

    // This attribute is a key. The encapsulation method used by the tunnel. The
    // type is IANAtunnelType.
    Tunnelinetconfigencapsmethod interface{}

    // This attribute is a key. An identifier used to distinguish between multiple
    // tunnels of the same encapsulation method, with the same endpoints.  If the
    // encapsulation protocol only allows one tunnel per set of endpoint addresses
    // (such as for GRE or IP-in-IP), the value of this object is 1.  For
    // encapsulation methods (such as L2F) which allow multiple parallel tunnels,
    // the manager is responsible for choosing any ID which does not  conflict
    // with an existing row, such as choosing a random number. The type is
    // interface{} with range: 1..2147483647.
    Tunnelinetconfigid interface{}

    // If the value of tunnelInetConfigStatus for this row is active, then this
    // object contains the value of ifIndex corresponding to the tunnel interface.
    // A value of 0 is not legal in the active state, and means that the interface
    // index has not yet been assigned. The type is interface{} with range:
    // 0..2147483647.
    Tunnelinetconfigifindex interface{}

    // The status of this row, by which new entries may be created, or old entries
    // deleted from this table.  The agent need not support setting this object to
    // createAndWait or notInService since there are no other writable objects in
    // this table, and writable objects in rows of corresponding tables such as
    // the tunnelIfTable may be modified while this row is active.  To create a
    // row in this table for an encapsulation method which does not support
    // multiple parallel tunnels with the same endpoints, the management station
    // should simply use a tunnelInetConfigID of 1, and set tunnelInetConfigStatus
    // to createAndGo.  For encapsulation methods such as L2F which allow multiple
    // parallel tunnels, the management station may select a pseudo-random number
    // to use as the tunnelInetConfigID and set tunnelInetConfigStatus to
    // createAndGo.  In the event that this ID is already in use and an
    // inconsistentValue is returned in response to the set operation, the
    // management station should simply select a new pseudo-random number and
    // retry the operation.  Creating a row in this table will cause an interface
    // index to be assigned by the agent in an implementation-dependent manner,
    // and corresponding rows will be instantiated in the ifTable and the 
    // tunnelIfTable.  The status of this row will become active as soon as the
    // agent assigns the interface index, regardless of whether the interface is
    // operationally up.  Deleting a row in this table will likewise delete the
    // corresponding row in the ifTable and in the tunnelIfTable. The type is
    // RowStatus.
    Tunnelinetconfigstatus interface{}

    // The storage type of this row.  If the row is permanent(4), no objects in
    // the row need be writable. The type is StorageType.
    Tunnelinetconfigstoragetype interface{}
}

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetFilter() yfilter.YFilter { return tunnelinetconfigentry.YFilter }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) SetFilter(yf yfilter.YFilter) { tunnelinetconfigentry.YFilter = yf }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetGoName(yname string) string {
    if yname == "tunnelInetConfigAddressType" { return "Tunnelinetconfigaddresstype" }
    if yname == "tunnelInetConfigLocalAddress" { return "Tunnelinetconfiglocaladdress" }
    if yname == "tunnelInetConfigRemoteAddress" { return "Tunnelinetconfigremoteaddress" }
    if yname == "tunnelInetConfigEncapsMethod" { return "Tunnelinetconfigencapsmethod" }
    if yname == "tunnelInetConfigID" { return "Tunnelinetconfigid" }
    if yname == "tunnelInetConfigIfIndex" { return "Tunnelinetconfigifindex" }
    if yname == "tunnelInetConfigStatus" { return "Tunnelinetconfigstatus" }
    if yname == "tunnelInetConfigStorageType" { return "Tunnelinetconfigstoragetype" }
    return ""
}

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetSegmentPath() string {
    return "tunnelInetConfigEntry" + "[tunnelInetConfigAddressType='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigaddresstype) + "']" + "[tunnelInetConfigLocalAddress='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfiglocaladdress) + "']" + "[tunnelInetConfigRemoteAddress='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigremoteaddress) + "']" + "[tunnelInetConfigEncapsMethod='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigencapsmethod) + "']" + "[tunnelInetConfigID='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigid) + "']"
}

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnelInetConfigAddressType"] = tunnelinetconfigentry.Tunnelinetconfigaddresstype
    leafs["tunnelInetConfigLocalAddress"] = tunnelinetconfigentry.Tunnelinetconfiglocaladdress
    leafs["tunnelInetConfigRemoteAddress"] = tunnelinetconfigentry.Tunnelinetconfigremoteaddress
    leafs["tunnelInetConfigEncapsMethod"] = tunnelinetconfigentry.Tunnelinetconfigencapsmethod
    leafs["tunnelInetConfigID"] = tunnelinetconfigentry.Tunnelinetconfigid
    leafs["tunnelInetConfigIfIndex"] = tunnelinetconfigentry.Tunnelinetconfigifindex
    leafs["tunnelInetConfigStatus"] = tunnelinetconfigentry.Tunnelinetconfigstatus
    leafs["tunnelInetConfigStorageType"] = tunnelinetconfigentry.Tunnelinetconfigstoragetype
    return leafs
}

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetYangName() string { return "tunnelInetConfigEntry" }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) SetParent(parent types.Entity) { tunnelinetconfigentry.parent = parent }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetParent() types.Entity { return tunnelinetconfigentry.parent }

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetParentYangName() string { return "tunnelInetConfigTable" }

