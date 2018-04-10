// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-telnet package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-telnet: IPv6 telnet configuration
//   ipv4-telnet: ipv4 telnet
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_telnet_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_telnet_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-telnet-cfg ipv6-telnet}", reflect.TypeOf(Ipv6Telnet{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-telnet-cfg:ipv6-telnet", reflect.TypeOf(Ipv6Telnet{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-telnet-cfg ipv4-telnet}", reflect.TypeOf(Ipv4Telnet{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-telnet-cfg:ipv4-telnet", reflect.TypeOf(Ipv4Telnet{}))
}

// Ipv6Telnet
// IPv6 telnet configuration
type Ipv6Telnet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Telnet client configuration.
    Client Ipv6Telnet_Client
}

func (ipv6Telnet *Ipv6Telnet) GetEntityData() *types.CommonEntityData {
    ipv6Telnet.EntityData.YFilter = ipv6Telnet.YFilter
    ipv6Telnet.EntityData.YangName = "ipv6-telnet"
    ipv6Telnet.EntityData.BundleName = "cisco_ios_xr"
    ipv6Telnet.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-telnet-cfg"
    ipv6Telnet.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-telnet-cfg:ipv6-telnet"
    ipv6Telnet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Telnet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Telnet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Telnet.EntityData.Children = make(map[string]types.YChild)
    ipv6Telnet.EntityData.Children["client"] = types.YChild{"Client", &ipv6Telnet.Client}
    ipv6Telnet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Telnet.EntityData)
}

// Ipv6Telnet_Client
// Telnet client configuration
type Ipv6Telnet_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source interface for telnet sessions. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}
}

func (client *Ipv6Telnet_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "ipv6-telnet"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", client.SourceInterface}
    return &(client.EntityData)
}

// Ipv4Telnet
// ipv4 telnet
type Ipv4Telnet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Telnet client configuration.
    Client Ipv4Telnet_Client
}

func (ipv4Telnet *Ipv4Telnet) GetEntityData() *types.CommonEntityData {
    ipv4Telnet.EntityData.YFilter = ipv4Telnet.YFilter
    ipv4Telnet.EntityData.YangName = "ipv4-telnet"
    ipv4Telnet.EntityData.BundleName = "cisco_ios_xr"
    ipv4Telnet.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-telnet-cfg"
    ipv4Telnet.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-telnet-cfg:ipv4-telnet"
    ipv4Telnet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Telnet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Telnet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Telnet.EntityData.Children = make(map[string]types.YChild)
    ipv4Telnet.EntityData.Children["client"] = types.YChild{"Client", &ipv4Telnet.Client}
    ipv4Telnet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Telnet.EntityData)
}

// Ipv4Telnet_Client
// Telnet client configuration
type Ipv4Telnet_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source interface for telnet sessions. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}
}

func (client *Ipv4Telnet_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "ipv4-telnet"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", client.SourceInterface}
    return &(client.EntityData)
}

