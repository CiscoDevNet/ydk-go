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
    EntityData types.CommonEntityData
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

func (tUNNELMIB *TUNNELMIB) GetEntityData() *types.CommonEntityData {
    tUNNELMIB.EntityData.YFilter = tUNNELMIB.YFilter
    tUNNELMIB.EntityData.YangName = "TUNNEL-MIB"
    tUNNELMIB.EntityData.BundleName = "cisco_ios_xe"
    tUNNELMIB.EntityData.ParentYangName = "TUNNEL-MIB"
    tUNNELMIB.EntityData.SegmentPath = "TUNNEL-MIB:TUNNEL-MIB"
    tUNNELMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tUNNELMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tUNNELMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tUNNELMIB.EntityData.Children = make(map[string]types.YChild)
    tUNNELMIB.EntityData.Children["tunnelIfTable"] = types.YChild{"Tunneliftable", &tUNNELMIB.Tunneliftable}
    tUNNELMIB.EntityData.Children["tunnelConfigTable"] = types.YChild{"Tunnelconfigtable", &tUNNELMIB.Tunnelconfigtable}
    tUNNELMIB.EntityData.Children["tunnelInetConfigTable"] = types.YChild{"Tunnelinetconfigtable", &tUNNELMIB.Tunnelinetconfigtable}
    tUNNELMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tUNNELMIB.EntityData)
}

// TUNNELMIB_Tunneliftable
// The (conceptual) table containing information on
// configured tunnels.
type TUNNELMIB_Tunneliftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the information on a particular
    // configured tunnel. The type is slice of
    // TUNNELMIB_Tunneliftable_Tunnelifentry.
    Tunnelifentry []TUNNELMIB_Tunneliftable_Tunnelifentry
}

func (tunneliftable *TUNNELMIB_Tunneliftable) GetEntityData() *types.CommonEntityData {
    tunneliftable.EntityData.YFilter = tunneliftable.YFilter
    tunneliftable.EntityData.YangName = "tunnelIfTable"
    tunneliftable.EntityData.BundleName = "cisco_ios_xe"
    tunneliftable.EntityData.ParentYangName = "TUNNEL-MIB"
    tunneliftable.EntityData.SegmentPath = "tunnelIfTable"
    tunneliftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunneliftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunneliftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunneliftable.EntityData.Children = make(map[string]types.YChild)
    tunneliftable.EntityData.Children["tunnelIfEntry"] = types.YChild{"Tunnelifentry", nil}
    for i := range tunneliftable.Tunnelifentry {
        tunneliftable.EntityData.Children[types.GetSegmentPath(&tunneliftable.Tunnelifentry[i])] = types.YChild{"Tunnelifentry", &tunneliftable.Tunnelifentry[i]}
    }
    tunneliftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunneliftable.EntityData)
}

// TUNNELMIB_Tunneliftable_Tunnelifentry
// An entry (conceptual row) containing the information
// on a particular configured tunnel.
type TUNNELMIB_Tunneliftable_Tunnelifentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The address of the local endpoint of the tunnel (i.e., the source address
    // used in the outer IP header), or 0.0.0.0 if unknown or if the tunnel is
    // over IPv6.  Since this object does not support IPv6, it is deprecated in
    // favor of tunnelIfLocalInetAddress. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Tunneliflocaladdress interface{}

    // The address of the remote endpoint of the tunnel (i.e., the destination
    // address used in the outer IP header), or 0.0.0.0 if unknown, or an IPv6
    // address, or  the tunnel is not a point-to-point link (e.g., if it is a 6to4
    // tunnel).  Since this object does not support IPv6, it is deprecated in
    // favor of tunnelIfRemoteInetAddress. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (tunnelifentry *TUNNELMIB_Tunneliftable_Tunnelifentry) GetEntityData() *types.CommonEntityData {
    tunnelifentry.EntityData.YFilter = tunnelifentry.YFilter
    tunnelifentry.EntityData.YangName = "tunnelIfEntry"
    tunnelifentry.EntityData.BundleName = "cisco_ios_xe"
    tunnelifentry.EntityData.ParentYangName = "tunnelIfTable"
    tunnelifentry.EntityData.SegmentPath = "tunnelIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", tunnelifentry.Ifindex) + "']"
    tunnelifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelifentry.EntityData.Children = make(map[string]types.YChild)
    tunnelifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelifentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", tunnelifentry.Ifindex}
    tunnelifentry.EntityData.Leafs["tunnelIfLocalAddress"] = types.YLeaf{"Tunneliflocaladdress", tunnelifentry.Tunneliflocaladdress}
    tunnelifentry.EntityData.Leafs["tunnelIfRemoteAddress"] = types.YLeaf{"Tunnelifremoteaddress", tunnelifentry.Tunnelifremoteaddress}
    tunnelifentry.EntityData.Leafs["tunnelIfEncapsMethod"] = types.YLeaf{"Tunnelifencapsmethod", tunnelifentry.Tunnelifencapsmethod}
    tunnelifentry.EntityData.Leafs["tunnelIfHopLimit"] = types.YLeaf{"Tunnelifhoplimit", tunnelifentry.Tunnelifhoplimit}
    tunnelifentry.EntityData.Leafs["tunnelIfSecurity"] = types.YLeaf{"Tunnelifsecurity", tunnelifentry.Tunnelifsecurity}
    tunnelifentry.EntityData.Leafs["tunnelIfTOS"] = types.YLeaf{"Tunneliftos", tunnelifentry.Tunneliftos}
    tunnelifentry.EntityData.Leafs["tunnelIfFlowLabel"] = types.YLeaf{"Tunnelifflowlabel", tunnelifentry.Tunnelifflowlabel}
    tunnelifentry.EntityData.Leafs["tunnelIfAddressType"] = types.YLeaf{"Tunnelifaddresstype", tunnelifentry.Tunnelifaddresstype}
    tunnelifentry.EntityData.Leafs["tunnelIfLocalInetAddress"] = types.YLeaf{"Tunneliflocalinetaddress", tunnelifentry.Tunneliflocalinetaddress}
    tunnelifentry.EntityData.Leafs["tunnelIfRemoteInetAddress"] = types.YLeaf{"Tunnelifremoteinetaddress", tunnelifentry.Tunnelifremoteinetaddress}
    tunnelifentry.EntityData.Leafs["tunnelIfEncapsLimit"] = types.YLeaf{"Tunnelifencapslimit", tunnelifentry.Tunnelifencapslimit}
    return &(tunnelifentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the information on a particular
    // configured tunnel.  Since this entry does not support IPv6, it is
    // deprecated in favor of tunnelInetConfigEntry. The type is slice of
    // TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry.
    Tunnelconfigentry []TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry
}

func (tunnelconfigtable *TUNNELMIB_Tunnelconfigtable) GetEntityData() *types.CommonEntityData {
    tunnelconfigtable.EntityData.YFilter = tunnelconfigtable.YFilter
    tunnelconfigtable.EntityData.YangName = "tunnelConfigTable"
    tunnelconfigtable.EntityData.BundleName = "cisco_ios_xe"
    tunnelconfigtable.EntityData.ParentYangName = "TUNNEL-MIB"
    tunnelconfigtable.EntityData.SegmentPath = "tunnelConfigTable"
    tunnelconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelconfigtable.EntityData.Children = make(map[string]types.YChild)
    tunnelconfigtable.EntityData.Children["tunnelConfigEntry"] = types.YChild{"Tunnelconfigentry", nil}
    for i := range tunnelconfigtable.Tunnelconfigentry {
        tunnelconfigtable.EntityData.Children[types.GetSegmentPath(&tunnelconfigtable.Tunnelconfigentry[i])] = types.YChild{"Tunnelconfigentry", &tunnelconfigtable.Tunnelconfigentry[i]}
    }
    tunnelconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelconfigtable.EntityData)
}

// TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry
// An entry (conceptual row) containing the information
// on a particular configured tunnel.
// 
// Since this entry does not support IPv6, it is
// deprecated in favor of tunnelInetConfigEntry.
type TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The address of the local endpoint of the tunnel,
    // or 0.0.0.0 if the device is free to choose any of its addresses at tunnel
    // establishment time.  Since this object does not support IPv6, it is
    // deprecated in favor of tunnelInetConfigLocalAddress. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Tunnelconfiglocaladdress interface{}

    // This attribute is a key. The address of the remote endpoint of the tunnel. 
    // Since this object does not support IPv6, it is deprecated in favor of
    // tunnelInetConfigRemoteAddress. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (tunnelconfigentry *TUNNELMIB_Tunnelconfigtable_Tunnelconfigentry) GetEntityData() *types.CommonEntityData {
    tunnelconfigentry.EntityData.YFilter = tunnelconfigentry.YFilter
    tunnelconfigentry.EntityData.YangName = "tunnelConfigEntry"
    tunnelconfigentry.EntityData.BundleName = "cisco_ios_xe"
    tunnelconfigentry.EntityData.ParentYangName = "tunnelConfigTable"
    tunnelconfigentry.EntityData.SegmentPath = "tunnelConfigEntry" + "[tunnelConfigLocalAddress='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfiglocaladdress) + "']" + "[tunnelConfigRemoteAddress='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfigremoteaddress) + "']" + "[tunnelConfigEncapsMethod='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfigencapsmethod) + "']" + "[tunnelConfigID='" + fmt.Sprintf("%v", tunnelconfigentry.Tunnelconfigid) + "']"
    tunnelconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelconfigentry.EntityData.Children = make(map[string]types.YChild)
    tunnelconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelconfigentry.EntityData.Leafs["tunnelConfigLocalAddress"] = types.YLeaf{"Tunnelconfiglocaladdress", tunnelconfigentry.Tunnelconfiglocaladdress}
    tunnelconfigentry.EntityData.Leafs["tunnelConfigRemoteAddress"] = types.YLeaf{"Tunnelconfigremoteaddress", tunnelconfigentry.Tunnelconfigremoteaddress}
    tunnelconfigentry.EntityData.Leafs["tunnelConfigEncapsMethod"] = types.YLeaf{"Tunnelconfigencapsmethod", tunnelconfigentry.Tunnelconfigencapsmethod}
    tunnelconfigentry.EntityData.Leafs["tunnelConfigID"] = types.YLeaf{"Tunnelconfigid", tunnelconfigentry.Tunnelconfigid}
    tunnelconfigentry.EntityData.Leafs["tunnelConfigIfIndex"] = types.YLeaf{"Tunnelconfigifindex", tunnelconfigentry.Tunnelconfigifindex}
    tunnelconfigentry.EntityData.Leafs["tunnelConfigStatus"] = types.YLeaf{"Tunnelconfigstatus", tunnelconfigentry.Tunnelconfigstatus}
    return &(tunnelconfigentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (tunnelinetconfigtable *TUNNELMIB_Tunnelinetconfigtable) GetEntityData() *types.CommonEntityData {
    tunnelinetconfigtable.EntityData.YFilter = tunnelinetconfigtable.YFilter
    tunnelinetconfigtable.EntityData.YangName = "tunnelInetConfigTable"
    tunnelinetconfigtable.EntityData.BundleName = "cisco_ios_xe"
    tunnelinetconfigtable.EntityData.ParentYangName = "TUNNEL-MIB"
    tunnelinetconfigtable.EntityData.SegmentPath = "tunnelInetConfigTable"
    tunnelinetconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelinetconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelinetconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelinetconfigtable.EntityData.Children = make(map[string]types.YChild)
    tunnelinetconfigtable.EntityData.Children["tunnelInetConfigEntry"] = types.YChild{"Tunnelinetconfigentry", nil}
    for i := range tunnelinetconfigtable.Tunnelinetconfigentry {
        tunnelinetconfigtable.EntityData.Children[types.GetSegmentPath(&tunnelinetconfigtable.Tunnelinetconfigentry[i])] = types.YChild{"Tunnelinetconfigentry", &tunnelinetconfigtable.Tunnelinetconfigentry[i]}
    }
    tunnelinetconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelinetconfigtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (tunnelinetconfigentry *TUNNELMIB_Tunnelinetconfigtable_Tunnelinetconfigentry) GetEntityData() *types.CommonEntityData {
    tunnelinetconfigentry.EntityData.YFilter = tunnelinetconfigentry.YFilter
    tunnelinetconfigentry.EntityData.YangName = "tunnelInetConfigEntry"
    tunnelinetconfigentry.EntityData.BundleName = "cisco_ios_xe"
    tunnelinetconfigentry.EntityData.ParentYangName = "tunnelInetConfigTable"
    tunnelinetconfigentry.EntityData.SegmentPath = "tunnelInetConfigEntry" + "[tunnelInetConfigAddressType='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigaddresstype) + "']" + "[tunnelInetConfigLocalAddress='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfiglocaladdress) + "']" + "[tunnelInetConfigRemoteAddress='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigremoteaddress) + "']" + "[tunnelInetConfigEncapsMethod='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigencapsmethod) + "']" + "[tunnelInetConfigID='" + fmt.Sprintf("%v", tunnelinetconfigentry.Tunnelinetconfigid) + "']"
    tunnelinetconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelinetconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelinetconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelinetconfigentry.EntityData.Children = make(map[string]types.YChild)
    tunnelinetconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigAddressType"] = types.YLeaf{"Tunnelinetconfigaddresstype", tunnelinetconfigentry.Tunnelinetconfigaddresstype}
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigLocalAddress"] = types.YLeaf{"Tunnelinetconfiglocaladdress", tunnelinetconfigentry.Tunnelinetconfiglocaladdress}
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigRemoteAddress"] = types.YLeaf{"Tunnelinetconfigremoteaddress", tunnelinetconfigentry.Tunnelinetconfigremoteaddress}
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigEncapsMethod"] = types.YLeaf{"Tunnelinetconfigencapsmethod", tunnelinetconfigentry.Tunnelinetconfigencapsmethod}
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigID"] = types.YLeaf{"Tunnelinetconfigid", tunnelinetconfigentry.Tunnelinetconfigid}
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigIfIndex"] = types.YLeaf{"Tunnelinetconfigifindex", tunnelinetconfigentry.Tunnelinetconfigifindex}
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigStatus"] = types.YLeaf{"Tunnelinetconfigstatus", tunnelinetconfigentry.Tunnelinetconfigstatus}
    tunnelinetconfigentry.EntityData.Leafs["tunnelInetConfigStorageType"] = types.YLeaf{"Tunnelinetconfigstoragetype", tunnelinetconfigentry.Tunnelinetconfigstoragetype}
    return &(tunnelinetconfigentry.EntityData)
}

