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
    parent types.Entity
    YFilter yfilter.YFilter

    // Telnet client configuration.
    Client Ipv6Telnet_Client
}

func (ipv6Telnet *Ipv6Telnet) GetFilter() yfilter.YFilter { return ipv6Telnet.YFilter }

func (ipv6Telnet *Ipv6Telnet) SetFilter(yf yfilter.YFilter) { ipv6Telnet.YFilter = yf }

func (ipv6Telnet *Ipv6Telnet) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (ipv6Telnet *Ipv6Telnet) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-telnet-cfg:ipv6-telnet"
}

func (ipv6Telnet *Ipv6Telnet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        return &ipv6Telnet.Client
    }
    return nil
}

func (ipv6Telnet *Ipv6Telnet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client"] = &ipv6Telnet.Client
    return children
}

func (ipv6Telnet *Ipv6Telnet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Telnet *Ipv6Telnet) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Telnet *Ipv6Telnet) GetYangName() string { return "ipv6-telnet" }

func (ipv6Telnet *Ipv6Telnet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Telnet *Ipv6Telnet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Telnet *Ipv6Telnet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Telnet *Ipv6Telnet) SetParent(parent types.Entity) { ipv6Telnet.parent = parent }

func (ipv6Telnet *Ipv6Telnet) GetParent() types.Entity { return ipv6Telnet.parent }

func (ipv6Telnet *Ipv6Telnet) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-telnet-cfg" }

// Ipv6Telnet_Client
// Telnet client configuration
type Ipv6Telnet_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source interface for telnet sessions. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SourceInterface interface{}
}

func (client *Ipv6Telnet_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Ipv6Telnet_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Ipv6Telnet_Client) GetGoName(yname string) string {
    if yname == "source-interface" { return "SourceInterface" }
    return ""
}

func (client *Ipv6Telnet_Client) GetSegmentPath() string {
    return "client"
}

func (client *Ipv6Telnet_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Ipv6Telnet_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Ipv6Telnet_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-interface"] = client.SourceInterface
    return leafs
}

func (client *Ipv6Telnet_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Ipv6Telnet_Client) GetYangName() string { return "client" }

func (client *Ipv6Telnet_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Ipv6Telnet_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Ipv6Telnet_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Ipv6Telnet_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Ipv6Telnet_Client) GetParent() types.Entity { return client.parent }

func (client *Ipv6Telnet_Client) GetParentYangName() string { return "ipv6-telnet" }

// Ipv4Telnet
// ipv4 telnet
type Ipv4Telnet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Telnet client configuration.
    Client Ipv4Telnet_Client
}

func (ipv4Telnet *Ipv4Telnet) GetFilter() yfilter.YFilter { return ipv4Telnet.YFilter }

func (ipv4Telnet *Ipv4Telnet) SetFilter(yf yfilter.YFilter) { ipv4Telnet.YFilter = yf }

func (ipv4Telnet *Ipv4Telnet) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (ipv4Telnet *Ipv4Telnet) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-telnet-cfg:ipv4-telnet"
}

func (ipv4Telnet *Ipv4Telnet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        return &ipv4Telnet.Client
    }
    return nil
}

func (ipv4Telnet *Ipv4Telnet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client"] = &ipv4Telnet.Client
    return children
}

func (ipv4Telnet *Ipv4Telnet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Telnet *Ipv4Telnet) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Telnet *Ipv4Telnet) GetYangName() string { return "ipv4-telnet" }

func (ipv4Telnet *Ipv4Telnet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Telnet *Ipv4Telnet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Telnet *Ipv4Telnet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Telnet *Ipv4Telnet) SetParent(parent types.Entity) { ipv4Telnet.parent = parent }

func (ipv4Telnet *Ipv4Telnet) GetParent() types.Entity { return ipv4Telnet.parent }

func (ipv4Telnet *Ipv4Telnet) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-telnet-cfg" }

// Ipv4Telnet_Client
// Telnet client configuration
type Ipv4Telnet_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source interface for telnet sessions. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SourceInterface interface{}
}

func (client *Ipv4Telnet_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Ipv4Telnet_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Ipv4Telnet_Client) GetGoName(yname string) string {
    if yname == "source-interface" { return "SourceInterface" }
    return ""
}

func (client *Ipv4Telnet_Client) GetSegmentPath() string {
    return "client"
}

func (client *Ipv4Telnet_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Ipv4Telnet_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Ipv4Telnet_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-interface"] = client.SourceInterface
    return leafs
}

func (client *Ipv4Telnet_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Ipv4Telnet_Client) GetYangName() string { return "client" }

func (client *Ipv4Telnet_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Ipv4Telnet_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Ipv4Telnet_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Ipv4Telnet_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Ipv4Telnet_Client) GetParent() types.Entity { return client.parent }

func (client *Ipv4Telnet_Client) GetParentYangName() string { return "ipv4-telnet" }

