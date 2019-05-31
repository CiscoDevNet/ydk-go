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
    TunnelIfTable TUNNELMIB_TunnelIfTable

    // The (conceptual) table containing information on configured tunnels.  This
    // table can be used to map a set of tunnel endpoints to the associated
    // ifIndex value.  It can also be used for row creation.  Note that every row
    // in the tunnelIfTable with a fixed IPv4 destination address should have a
    // corresponding row in the tunnelConfigTable, regardless of whether it was
    // created via SNMP.  Since this table does not support IPv6, it is deprecated
    // in favor of tunnelInetConfigTable.
    TunnelConfigTable TUNNELMIB_TunnelConfigTable

    // The (conceptual) table containing information on configured tunnels.  This
    // table can be used to map a set of tunnel endpoints to the associated
    // ifIndex value.  It can also be used for row creation.  Note that every row
    // in the tunnelIfTable with a fixed destination address should have a
    // corresponding row in the tunnelInetConfigTable, regardless of whether it
    // was created via SNMP.
    TunnelInetConfigTable TUNNELMIB_TunnelInetConfigTable
}

func (tUNNELMIB *TUNNELMIB) GetEntityData() *types.CommonEntityData {
    tUNNELMIB.EntityData.YFilter = tUNNELMIB.YFilter
    tUNNELMIB.EntityData.YangName = "TUNNEL-MIB"
    tUNNELMIB.EntityData.BundleName = "cisco_ios_xe"
    tUNNELMIB.EntityData.ParentYangName = "TUNNEL-MIB"
    tUNNELMIB.EntityData.SegmentPath = "TUNNEL-MIB:TUNNEL-MIB"
    tUNNELMIB.EntityData.AbsolutePath = tUNNELMIB.EntityData.SegmentPath
    tUNNELMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tUNNELMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tUNNELMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tUNNELMIB.EntityData.Children = types.NewOrderedMap()
    tUNNELMIB.EntityData.Children.Append("tunnelIfTable", types.YChild{"TunnelIfTable", &tUNNELMIB.TunnelIfTable})
    tUNNELMIB.EntityData.Children.Append("tunnelConfigTable", types.YChild{"TunnelConfigTable", &tUNNELMIB.TunnelConfigTable})
    tUNNELMIB.EntityData.Children.Append("tunnelInetConfigTable", types.YChild{"TunnelInetConfigTable", &tUNNELMIB.TunnelInetConfigTable})
    tUNNELMIB.EntityData.Leafs = types.NewOrderedMap()

    tUNNELMIB.EntityData.YListKeys = []string {}

    return &(tUNNELMIB.EntityData)
}

// TUNNELMIB_TunnelIfTable
// The (conceptual) table containing information on
// configured tunnels.
type TUNNELMIB_TunnelIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the information on a particular
    // configured tunnel. The type is slice of
    // TUNNELMIB_TunnelIfTable_TunnelIfEntry.
    TunnelIfEntry []*TUNNELMIB_TunnelIfTable_TunnelIfEntry
}

func (tunnelIfTable *TUNNELMIB_TunnelIfTable) GetEntityData() *types.CommonEntityData {
    tunnelIfTable.EntityData.YFilter = tunnelIfTable.YFilter
    tunnelIfTable.EntityData.YangName = "tunnelIfTable"
    tunnelIfTable.EntityData.BundleName = "cisco_ios_xe"
    tunnelIfTable.EntityData.ParentYangName = "TUNNEL-MIB"
    tunnelIfTable.EntityData.SegmentPath = "tunnelIfTable"
    tunnelIfTable.EntityData.AbsolutePath = "TUNNEL-MIB:TUNNEL-MIB/" + tunnelIfTable.EntityData.SegmentPath
    tunnelIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelIfTable.EntityData.Children = types.NewOrderedMap()
    tunnelIfTable.EntityData.Children.Append("tunnelIfEntry", types.YChild{"TunnelIfEntry", nil})
    for i := range tunnelIfTable.TunnelIfEntry {
        tunnelIfTable.EntityData.Children.Append(types.GetSegmentPath(tunnelIfTable.TunnelIfEntry[i]), types.YChild{"TunnelIfEntry", tunnelIfTable.TunnelIfEntry[i]})
    }
    tunnelIfTable.EntityData.Leafs = types.NewOrderedMap()

    tunnelIfTable.EntityData.YListKeys = []string {}

    return &(tunnelIfTable.EntityData)
}

// TUNNELMIB_TunnelIfTable_TunnelIfEntry
// An entry (conceptual row) containing the information
// on a particular configured tunnel.
type TUNNELMIB_TunnelIfTable_TunnelIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The address of the local endpoint of the tunnel (i.e., the source address
    // used in the outer IP header), or 0.0.0.0 if unknown or if the tunnel is
    // over IPv6.  Since this object does not support IPv6, it is deprecated in
    // favor of tunnelIfLocalInetAddress. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TunnelIfLocalAddress interface{}

    // The address of the remote endpoint of the tunnel (i.e., the destination
    // address used in the outer IP header), or 0.0.0.0 if unknown, or an IPv6
    // address, or  the tunnel is not a point-to-point link (e.g., if it is a 6to4
    // tunnel).  Since this object does not support IPv6, it is deprecated in
    // favor of tunnelIfRemoteInetAddress. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TunnelIfRemoteAddress interface{}

    // The encapsulation method used by the tunnel. The type is IANAtunnelType.
    TunnelIfEncapsMethod interface{}

    // The IPv4 TTL or IPv6 Hop Limit to use in the outer IP header.  A value of 0
    // indicates that the value is copied from the payload's header. The type is
    // interface{} with range: 0..255.
    TunnelIfHopLimit interface{}

    // The method used by the tunnel to secure the outer IP header.  The value
    // ipsec indicates that IPsec is used between the tunnel endpoints for
    // authentication or encryption or both.  More specific security-related
    // information may be available in a MIB module for the security protocol in
    // use. The type is TunnelIfSecurity.
    TunnelIfSecurity interface{}

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
    TunnelIfTOS interface{}

    // The method used to set the IPv6 Flow Label value. This object need not be
    // present in rows where tunnelIfAddressType indicates the tunnel is not over
    // IPv6.  A value of -1 indicates that a traffic conditioner is invoked and
    // more information may be available in a traffic conditioner MIB.  Any other
    // value indicates that the Flow Label field is set to the indicated value.
    // The type is interface{} with range: -1..100.
    TunnelIfFlowLabel interface{}

    // The type of address in the corresponding tunnelIfLocalInetAddress and
    // tunnelIfRemoteInetAddress objects. The type is InetAddressType.
    TunnelIfAddressType interface{}

    // The address of the local endpoint of the tunnel (i.e., the source address
    // used in the outer IP header).  If the address is unknown, the value is 
    // 0.0.0.0 for IPv4 or :: for IPv6.  The type of this object is given by
    // tunnelIfAddressType. The type is string with length: 0..255.
    TunnelIfLocalInetAddress interface{}

    // The address of the remote endpoint of the tunnel (i.e., the destination
    // address used in the outer IP header).  If the address is unknown or the
    // tunnel is not a point-to-point link (e.g., if it is a 6to4 tunnel), the
    // value is 0.0.0.0 for tunnels over IPv4 or :: for tunnels over IPv6.  The
    // type of this object is given by tunnelIfAddressType. The type is string
    // with length: 0..255.
    TunnelIfRemoteInetAddress interface{}

    // The maximum number of additional encapsulations permitted for packets
    // undergoing encapsulation at this node.  A value of -1 indicates that no
    // limit is present (except as a result of the packet size). The type is
    // interface{} with range: -1..255.
    TunnelIfEncapsLimit interface{}
}

func (tunnelIfEntry *TUNNELMIB_TunnelIfTable_TunnelIfEntry) GetEntityData() *types.CommonEntityData {
    tunnelIfEntry.EntityData.YFilter = tunnelIfEntry.YFilter
    tunnelIfEntry.EntityData.YangName = "tunnelIfEntry"
    tunnelIfEntry.EntityData.BundleName = "cisco_ios_xe"
    tunnelIfEntry.EntityData.ParentYangName = "tunnelIfTable"
    tunnelIfEntry.EntityData.SegmentPath = "tunnelIfEntry" + types.AddKeyToken(tunnelIfEntry.IfIndex, "ifIndex")
    tunnelIfEntry.EntityData.AbsolutePath = "TUNNEL-MIB:TUNNEL-MIB/tunnelIfTable/" + tunnelIfEntry.EntityData.SegmentPath
    tunnelIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelIfEntry.EntityData.Children = types.NewOrderedMap()
    tunnelIfEntry.EntityData.Leafs = types.NewOrderedMap()
    tunnelIfEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", tunnelIfEntry.IfIndex})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfLocalAddress", types.YLeaf{"TunnelIfLocalAddress", tunnelIfEntry.TunnelIfLocalAddress})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfRemoteAddress", types.YLeaf{"TunnelIfRemoteAddress", tunnelIfEntry.TunnelIfRemoteAddress})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfEncapsMethod", types.YLeaf{"TunnelIfEncapsMethod", tunnelIfEntry.TunnelIfEncapsMethod})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfHopLimit", types.YLeaf{"TunnelIfHopLimit", tunnelIfEntry.TunnelIfHopLimit})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfSecurity", types.YLeaf{"TunnelIfSecurity", tunnelIfEntry.TunnelIfSecurity})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfTOS", types.YLeaf{"TunnelIfTOS", tunnelIfEntry.TunnelIfTOS})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfFlowLabel", types.YLeaf{"TunnelIfFlowLabel", tunnelIfEntry.TunnelIfFlowLabel})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfAddressType", types.YLeaf{"TunnelIfAddressType", tunnelIfEntry.TunnelIfAddressType})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfLocalInetAddress", types.YLeaf{"TunnelIfLocalInetAddress", tunnelIfEntry.TunnelIfLocalInetAddress})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfRemoteInetAddress", types.YLeaf{"TunnelIfRemoteInetAddress", tunnelIfEntry.TunnelIfRemoteInetAddress})
    tunnelIfEntry.EntityData.Leafs.Append("tunnelIfEncapsLimit", types.YLeaf{"TunnelIfEncapsLimit", tunnelIfEntry.TunnelIfEncapsLimit})

    tunnelIfEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(tunnelIfEntry.EntityData)
}

// TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity represents security protocol in use.
type TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity string

const (
    TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity_none TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity = "none"

    TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity_ipsec TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity = "ipsec"

    TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity_other TUNNELMIB_TunnelIfTable_TunnelIfEntry_TunnelIfSecurity = "other"
)

// TUNNELMIB_TunnelConfigTable
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
type TUNNELMIB_TunnelConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) containing the information on a particular
    // configured tunnel.  Since this entry does not support IPv6, it is
    // deprecated in favor of tunnelInetConfigEntry. The type is slice of
    // TUNNELMIB_TunnelConfigTable_TunnelConfigEntry.
    TunnelConfigEntry []*TUNNELMIB_TunnelConfigTable_TunnelConfigEntry
}

func (tunnelConfigTable *TUNNELMIB_TunnelConfigTable) GetEntityData() *types.CommonEntityData {
    tunnelConfigTable.EntityData.YFilter = tunnelConfigTable.YFilter
    tunnelConfigTable.EntityData.YangName = "tunnelConfigTable"
    tunnelConfigTable.EntityData.BundleName = "cisco_ios_xe"
    tunnelConfigTable.EntityData.ParentYangName = "TUNNEL-MIB"
    tunnelConfigTable.EntityData.SegmentPath = "tunnelConfigTable"
    tunnelConfigTable.EntityData.AbsolutePath = "TUNNEL-MIB:TUNNEL-MIB/" + tunnelConfigTable.EntityData.SegmentPath
    tunnelConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelConfigTable.EntityData.Children = types.NewOrderedMap()
    tunnelConfigTable.EntityData.Children.Append("tunnelConfigEntry", types.YChild{"TunnelConfigEntry", nil})
    for i := range tunnelConfigTable.TunnelConfigEntry {
        tunnelConfigTable.EntityData.Children.Append(types.GetSegmentPath(tunnelConfigTable.TunnelConfigEntry[i]), types.YChild{"TunnelConfigEntry", tunnelConfigTable.TunnelConfigEntry[i]})
    }
    tunnelConfigTable.EntityData.Leafs = types.NewOrderedMap()

    tunnelConfigTable.EntityData.YListKeys = []string {}

    return &(tunnelConfigTable.EntityData)
}

// TUNNELMIB_TunnelConfigTable_TunnelConfigEntry
// An entry (conceptual row) containing the information
// on a particular configured tunnel.
// 
// Since this entry does not support IPv6, it is
// deprecated in favor of tunnelInetConfigEntry.
type TUNNELMIB_TunnelConfigTable_TunnelConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address of the local endpoint of the tunnel,
    // or 0.0.0.0 if the device is free to choose any of its addresses at tunnel
    // establishment time.  Since this object does not support IPv6, it is
    // deprecated in favor of tunnelInetConfigLocalAddress. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TunnelConfigLocalAddress interface{}

    // This attribute is a key. The address of the remote endpoint of the tunnel. 
    // Since this object does not support IPv6, it is deprecated in favor of
    // tunnelInetConfigRemoteAddress. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TunnelConfigRemoteAddress interface{}

    // This attribute is a key. The encapsulation method used by the tunnel. 
    // Since this object does not support IPv6, it is deprecated in favor of
    // tunnelInetConfigEncapsMethod. The type is IANAtunnelType.
    TunnelConfigEncapsMethod interface{}

    // This attribute is a key. An identifier used to distinguish between multiple
    // tunnels of the same encapsulation method, with the same endpoints.  If the
    // encapsulation protocol only allows one tunnel per set of endpoint addresses
    // (such as for GRE or IP-in-IP), the value of this object is 1.  For
    // encapsulation methods (such as L2F) which allow multiple parallel tunnels,
    // the manager is responsible for choosing any ID which does not conflict with
    // an existing row, such as choosing a random number.  Since this object does
    // not support IPv6, it is deprecated in favor of tunnelInetConfigID. The type
    // is interface{} with range: 1..2147483647.
    TunnelConfigID interface{}

    // If the value of tunnelConfigStatus for this row is active, then this object
    // contains the value of ifIndex corresponding to the tunnel interface.  A
    // value of 0 is not legal in the active state, and means that the interface
    // index has not yet been assigned.  Since this object does not support IPv6,
    // it is deprecated in favor of tunnelInetConfigIfIndex. The type is
    // interface{} with range: 0..2147483647.
    TunnelConfigIfIndex interface{}

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
    TunnelConfigStatus interface{}
}

func (tunnelConfigEntry *TUNNELMIB_TunnelConfigTable_TunnelConfigEntry) GetEntityData() *types.CommonEntityData {
    tunnelConfigEntry.EntityData.YFilter = tunnelConfigEntry.YFilter
    tunnelConfigEntry.EntityData.YangName = "tunnelConfigEntry"
    tunnelConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    tunnelConfigEntry.EntityData.ParentYangName = "tunnelConfigTable"
    tunnelConfigEntry.EntityData.SegmentPath = "tunnelConfigEntry" + types.AddKeyToken(tunnelConfigEntry.TunnelConfigLocalAddress, "tunnelConfigLocalAddress") + types.AddKeyToken(tunnelConfigEntry.TunnelConfigRemoteAddress, "tunnelConfigRemoteAddress") + types.AddKeyToken(tunnelConfigEntry.TunnelConfigEncapsMethod, "tunnelConfigEncapsMethod") + types.AddKeyToken(tunnelConfigEntry.TunnelConfigID, "tunnelConfigID")
    tunnelConfigEntry.EntityData.AbsolutePath = "TUNNEL-MIB:TUNNEL-MIB/tunnelConfigTable/" + tunnelConfigEntry.EntityData.SegmentPath
    tunnelConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelConfigEntry.EntityData.Children = types.NewOrderedMap()
    tunnelConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    tunnelConfigEntry.EntityData.Leafs.Append("tunnelConfigLocalAddress", types.YLeaf{"TunnelConfigLocalAddress", tunnelConfigEntry.TunnelConfigLocalAddress})
    tunnelConfigEntry.EntityData.Leafs.Append("tunnelConfigRemoteAddress", types.YLeaf{"TunnelConfigRemoteAddress", tunnelConfigEntry.TunnelConfigRemoteAddress})
    tunnelConfigEntry.EntityData.Leafs.Append("tunnelConfigEncapsMethod", types.YLeaf{"TunnelConfigEncapsMethod", tunnelConfigEntry.TunnelConfigEncapsMethod})
    tunnelConfigEntry.EntityData.Leafs.Append("tunnelConfigID", types.YLeaf{"TunnelConfigID", tunnelConfigEntry.TunnelConfigID})
    tunnelConfigEntry.EntityData.Leafs.Append("tunnelConfigIfIndex", types.YLeaf{"TunnelConfigIfIndex", tunnelConfigEntry.TunnelConfigIfIndex})
    tunnelConfigEntry.EntityData.Leafs.Append("tunnelConfigStatus", types.YLeaf{"TunnelConfigStatus", tunnelConfigEntry.TunnelConfigStatus})

    tunnelConfigEntry.EntityData.YListKeys = []string {"TunnelConfigLocalAddress", "TunnelConfigRemoteAddress", "TunnelConfigEncapsMethod", "TunnelConfigID"}

    return &(tunnelConfigEntry.EntityData)
}

// TUNNELMIB_TunnelInetConfigTable
// The (conceptual) table containing information on
// configured tunnels.  This table can be used to map a
// set of tunnel endpoints to the associated ifIndex
// value.  It can also be used for row creation.  Note
// that every row in the tunnelIfTable with a fixed
// destination address should have a corresponding row in
// the tunnelInetConfigTable, regardless of whether it
// was created via SNMP.
type TUNNELMIB_TunnelInetConfigTable struct {
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
    // TUNNELMIB_TunnelInetConfigTable_TunnelInetConfigEntry.
    TunnelInetConfigEntry []*TUNNELMIB_TunnelInetConfigTable_TunnelInetConfigEntry
}

func (tunnelInetConfigTable *TUNNELMIB_TunnelInetConfigTable) GetEntityData() *types.CommonEntityData {
    tunnelInetConfigTable.EntityData.YFilter = tunnelInetConfigTable.YFilter
    tunnelInetConfigTable.EntityData.YangName = "tunnelInetConfigTable"
    tunnelInetConfigTable.EntityData.BundleName = "cisco_ios_xe"
    tunnelInetConfigTable.EntityData.ParentYangName = "TUNNEL-MIB"
    tunnelInetConfigTable.EntityData.SegmentPath = "tunnelInetConfigTable"
    tunnelInetConfigTable.EntityData.AbsolutePath = "TUNNEL-MIB:TUNNEL-MIB/" + tunnelInetConfigTable.EntityData.SegmentPath
    tunnelInetConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelInetConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelInetConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelInetConfigTable.EntityData.Children = types.NewOrderedMap()
    tunnelInetConfigTable.EntityData.Children.Append("tunnelInetConfigEntry", types.YChild{"TunnelInetConfigEntry", nil})
    for i := range tunnelInetConfigTable.TunnelInetConfigEntry {
        tunnelInetConfigTable.EntityData.Children.Append(types.GetSegmentPath(tunnelInetConfigTable.TunnelInetConfigEntry[i]), types.YChild{"TunnelInetConfigEntry", tunnelInetConfigTable.TunnelInetConfigEntry[i]})
    }
    tunnelInetConfigTable.EntityData.Leafs = types.NewOrderedMap()

    tunnelInetConfigTable.EntityData.YListKeys = []string {}

    return &(tunnelInetConfigTable.EntityData)
}

// TUNNELMIB_TunnelInetConfigTable_TunnelInetConfigEntry
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
type TUNNELMIB_TunnelInetConfigTable_TunnelInetConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The address type over which the tunnel
    // encapsulates packets. The type is InetAddressType.
    TunnelInetConfigAddressType interface{}

    // This attribute is a key. The address of the local endpoint of the tunnel,
    // or 0.0.0.0 (for IPv4) or :: (for IPv6) if the device is free to choose any
    // of its addresses at tunnel establishment time. The type is string with
    // length: 0..255.
    TunnelInetConfigLocalAddress interface{}

    // This attribute is a key. The address of the remote endpoint of the tunnel.
    // The type is string with length: 0..255.
    TunnelInetConfigRemoteAddress interface{}

    // This attribute is a key. The encapsulation method used by the tunnel. The
    // type is IANAtunnelType.
    TunnelInetConfigEncapsMethod interface{}

    // This attribute is a key. An identifier used to distinguish between multiple
    // tunnels of the same encapsulation method, with the same endpoints.  If the
    // encapsulation protocol only allows one tunnel per set of endpoint addresses
    // (such as for GRE or IP-in-IP), the value of this object is 1.  For
    // encapsulation methods (such as L2F) which allow multiple parallel tunnels,
    // the manager is responsible for choosing any ID which does not  conflict
    // with an existing row, such as choosing a random number. The type is
    // interface{} with range: 1..2147483647.
    TunnelInetConfigID interface{}

    // If the value of tunnelInetConfigStatus for this row is active, then this
    // object contains the value of ifIndex corresponding to the tunnel interface.
    // A value of 0 is not legal in the active state, and means that the interface
    // index has not yet been assigned. The type is interface{} with range:
    // 0..2147483647.
    TunnelInetConfigIfIndex interface{}

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
    TunnelInetConfigStatus interface{}

    // The storage type of this row.  If the row is permanent(4), no objects in
    // the row need be writable. The type is StorageType.
    TunnelInetConfigStorageType interface{}
}

func (tunnelInetConfigEntry *TUNNELMIB_TunnelInetConfigTable_TunnelInetConfigEntry) GetEntityData() *types.CommonEntityData {
    tunnelInetConfigEntry.EntityData.YFilter = tunnelInetConfigEntry.YFilter
    tunnelInetConfigEntry.EntityData.YangName = "tunnelInetConfigEntry"
    tunnelInetConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    tunnelInetConfigEntry.EntityData.ParentYangName = "tunnelInetConfigTable"
    tunnelInetConfigEntry.EntityData.SegmentPath = "tunnelInetConfigEntry" + types.AddKeyToken(tunnelInetConfigEntry.TunnelInetConfigAddressType, "tunnelInetConfigAddressType") + types.AddKeyToken(tunnelInetConfigEntry.TunnelInetConfigLocalAddress, "tunnelInetConfigLocalAddress") + types.AddKeyToken(tunnelInetConfigEntry.TunnelInetConfigRemoteAddress, "tunnelInetConfigRemoteAddress") + types.AddKeyToken(tunnelInetConfigEntry.TunnelInetConfigEncapsMethod, "tunnelInetConfigEncapsMethod") + types.AddKeyToken(tunnelInetConfigEntry.TunnelInetConfigID, "tunnelInetConfigID")
    tunnelInetConfigEntry.EntityData.AbsolutePath = "TUNNEL-MIB:TUNNEL-MIB/tunnelInetConfigTable/" + tunnelInetConfigEntry.EntityData.SegmentPath
    tunnelInetConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tunnelInetConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tunnelInetConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tunnelInetConfigEntry.EntityData.Children = types.NewOrderedMap()
    tunnelInetConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigAddressType", types.YLeaf{"TunnelInetConfigAddressType", tunnelInetConfigEntry.TunnelInetConfigAddressType})
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigLocalAddress", types.YLeaf{"TunnelInetConfigLocalAddress", tunnelInetConfigEntry.TunnelInetConfigLocalAddress})
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigRemoteAddress", types.YLeaf{"TunnelInetConfigRemoteAddress", tunnelInetConfigEntry.TunnelInetConfigRemoteAddress})
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigEncapsMethod", types.YLeaf{"TunnelInetConfigEncapsMethod", tunnelInetConfigEntry.TunnelInetConfigEncapsMethod})
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigID", types.YLeaf{"TunnelInetConfigID", tunnelInetConfigEntry.TunnelInetConfigID})
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigIfIndex", types.YLeaf{"TunnelInetConfigIfIndex", tunnelInetConfigEntry.TunnelInetConfigIfIndex})
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigStatus", types.YLeaf{"TunnelInetConfigStatus", tunnelInetConfigEntry.TunnelInetConfigStatus})
    tunnelInetConfigEntry.EntityData.Leafs.Append("tunnelInetConfigStorageType", types.YLeaf{"TunnelInetConfigStorageType", tunnelInetConfigEntry.TunnelInetConfigStorageType})

    tunnelInetConfigEntry.EntityData.YListKeys = []string {"TunnelInetConfigAddressType", "TunnelInetConfigLocalAddress", "TunnelInetConfigRemoteAddress", "TunnelInetConfigEncapsMethod", "TunnelInetConfigID"}

    return &(tunnelInetConfigEntry.EntityData)
}

