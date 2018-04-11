// This module contains a collection of YANG definitions
// for IPv6 operational data.
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-ipv6-oper ipv6-data}", reflect.TypeOf(Ipv6Data{}))
    ydk.RegisterEntity("Cisco-IOS-XE-ipv6-oper:ipv6-data", reflect.TypeOf(Ipv6Data{}))
}

// Ipv6NdTdlState represents the Neighbor-Discovery entry.
type Ipv6NdTdlState string

const (
    // Address resolution is in progress and the 
    // link-layer address of the neighbor has not
    // yet been determined.
    Ipv6NdTdlState_ipv6_nd_incomplete Ipv6NdTdlState = "ipv6-nd-incomplete"

    // The neighbor is known to have been reachable
    // recently (within tens of seconds ago).
    Ipv6NdTdlState_ipv6_nd_reachable Ipv6NdTdlState = "ipv6-nd-reachable"

    // The neighbor is no longer known to be reachable
    // but until traffic is sent to the neighbor, no 
    // attempt should be made to verify its reachability
    Ipv6NdTdlState_ipv6_nd_stale Ipv6NdTdlState = "ipv6-nd-stale"

    // Received updated link-layer information. 
    // Behaviour like STALE.
    Ipv6NdTdlState_ipv6_nd_glean Ipv6NdTdlState = "ipv6-nd-glean"

    // The neighbor is no longer known to be reachable, and traffic
    // has recently been sent to the neighbor. Rather than probe the
    // neighbor immediately, however, delay sending probes for a short
    // while in order to give upper layer protocols a chance to provide
    // reachability confirmation.
    Ipv6NdTdlState_ipv6_nd_delay Ipv6NdTdlState = "ipv6-nd-delay"

    // The neighbor is no longer known to be reachable, and unicast
    // Neighbor Solicitation probes are being sent to verify 
    // reachability.
    Ipv6NdTdlState_ipv6_nd_probe Ipv6NdTdlState = "ipv6-nd-probe"

    // Entry is deleted.
    Ipv6NdTdlState_ipv6_nd_delete Ipv6NdTdlState = "ipv6-nd-delete"
)

// Ipv6Data
// Operational state of IPv6
type Ipv6Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of IPv6 neighbors. The type is slice of Ipv6Data_Nd6Info.
    Nd6Info []Ipv6Data_Nd6Info
}

func (ipv6Data *Ipv6Data) GetEntityData() *types.CommonEntityData {
    ipv6Data.EntityData.YFilter = ipv6Data.YFilter
    ipv6Data.EntityData.YangName = "ipv6-data"
    ipv6Data.EntityData.BundleName = "cisco_ios_xe"
    ipv6Data.EntityData.ParentYangName = "Cisco-IOS-XE-ipv6-oper"
    ipv6Data.EntityData.SegmentPath = "Cisco-IOS-XE-ipv6-oper:ipv6-data"
    ipv6Data.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipv6Data.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipv6Data.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipv6Data.EntityData.Children = make(map[string]types.YChild)
    ipv6Data.EntityData.Children["nd6-info"] = types.YChild{"Nd6Info", nil}
    for i := range ipv6Data.Nd6Info {
        ipv6Data.EntityData.Children[types.GetSegmentPath(&ipv6Data.Nd6Info[i])] = types.YChild{"Nd6Info", &ipv6Data.Nd6Info[i]}
    }
    ipv6Data.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Data.EntityData)
}

// Ipv6Data_Nd6Info
// List of IPv6 neighbors
type Ipv6Data_Nd6Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Virtual Router and Forwarding instance that 
    // this neighbor information is associated with. The type is string.
    VrfName interface{}

    // This attribute is a key. The Interface name. The type is string.
    IfName interface{}

    // This attribute is a key. IPv6 address of this neighbor entry. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // State of the entry. The type is Ipv6NdTdlState.
    State interface{}

    // Time before expiration, in seconds. The type is interface{} with range:
    // 0..4294967295.
    IdleTimer interface{}

    // Indicates how long this neighbor entry has  been active, in seconds. The
    // type is interface{} with range: 0..4294967295.
    Uptime interface{}
}

func (nd6Info *Ipv6Data_Nd6Info) GetEntityData() *types.CommonEntityData {
    nd6Info.EntityData.YFilter = nd6Info.YFilter
    nd6Info.EntityData.YangName = "nd6-info"
    nd6Info.EntityData.BundleName = "cisco_ios_xe"
    nd6Info.EntityData.ParentYangName = "ipv6-data"
    nd6Info.EntityData.SegmentPath = "nd6-info" + "[vrf-name='" + fmt.Sprintf("%v", nd6Info.VrfName) + "']" + "[if-name='" + fmt.Sprintf("%v", nd6Info.IfName) + "']" + "[ip='" + fmt.Sprintf("%v", nd6Info.Ip) + "']"
    nd6Info.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nd6Info.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nd6Info.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nd6Info.EntityData.Children = make(map[string]types.YChild)
    nd6Info.EntityData.Leafs = make(map[string]types.YLeaf)
    nd6Info.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", nd6Info.VrfName}
    nd6Info.EntityData.Leafs["if-name"] = types.YLeaf{"IfName", nd6Info.IfName}
    nd6Info.EntityData.Leafs["ip"] = types.YLeaf{"Ip", nd6Info.Ip}
    nd6Info.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nd6Info.MacAddress}
    nd6Info.EntityData.Leafs["state"] = types.YLeaf{"State", nd6Info.State}
    nd6Info.EntityData.Leafs["idle-timer"] = types.YLeaf{"IdleTimer", nd6Info.IdleTimer}
    nd6Info.EntityData.Leafs["uptime"] = types.YLeaf{"Uptime", nd6Info.Uptime}
    return &(nd6Info.EntityData)
}

