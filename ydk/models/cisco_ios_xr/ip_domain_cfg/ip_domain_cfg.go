// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-domain package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ip-domain: IP domain configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_domain_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_domain_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-domain-cfg ip-domain}", reflect.TypeOf(IpDomain{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-domain-cfg:ip-domain", reflect.TypeOf(IpDomain{}))
}

// IpDomain
// IP domain configuration
type IpDomain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs IpDomain_Vrfs
}

func (ipDomain *IpDomain) GetFilter() yfilter.YFilter { return ipDomain.YFilter }

func (ipDomain *IpDomain) SetFilter(yf yfilter.YFilter) { ipDomain.YFilter = yf }

func (ipDomain *IpDomain) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (ipDomain *IpDomain) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-domain-cfg:ip-domain"
}

func (ipDomain *IpDomain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &ipDomain.Vrfs
    }
    return nil
}

func (ipDomain *IpDomain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &ipDomain.Vrfs
    return children
}

func (ipDomain *IpDomain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipDomain *IpDomain) GetBundleName() string { return "cisco_ios_xr" }

func (ipDomain *IpDomain) GetYangName() string { return "ip-domain" }

func (ipDomain *IpDomain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipDomain *IpDomain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipDomain *IpDomain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipDomain *IpDomain) SetParent(parent types.Entity) { ipDomain.parent = parent }

func (ipDomain *IpDomain) GetParent() types.Entity { return ipDomain.parent }

func (ipDomain *IpDomain) GetParentYangName() string { return "Cisco-IOS-XR-ip-domain-cfg" }

// IpDomain_Vrfs
// VRF table
type IpDomain_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of IpDomain_Vrfs_Vrf.
    Vrf []IpDomain_Vrfs_Vrf
}

func (vrfs *IpDomain_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *IpDomain_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *IpDomain_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *IpDomain_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *IpDomain_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *IpDomain_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *IpDomain_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *IpDomain_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *IpDomain_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *IpDomain_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *IpDomain_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *IpDomain_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *IpDomain_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *IpDomain_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *IpDomain_Vrfs) GetParentYangName() string { return "ip-domain" }

// IpDomain_Vrfs_Vrf
// VRF specific data
type IpDomain_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Disable Domain Name System hostname translation. The type is interface{}.
    Lookup interface{}

    // Default multicast domain name. The type is string.
    MulticastDomain interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Default domain name. The type is string.
    Name interface{}

    // IPv6 host.
    Ipv6Hosts IpDomain_Vrfs_Vrf_Ipv6Hosts

    // Name server addresses.
    Servers IpDomain_Vrfs_Vrf_Servers

    // Domain names to complete unqualified host names.
    Lists IpDomain_Vrfs_Vrf_Lists

    // IPv4 host.
    Ipv4Hosts IpDomain_Vrfs_Vrf_Ipv4Hosts
}

func (vrf *IpDomain_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *IpDomain_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *IpDomain_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "lookup" { return "Lookup" }
    if yname == "multicast-domain" { return "MulticastDomain" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "name" { return "Name" }
    if yname == "ipv6-hosts" { return "Ipv6Hosts" }
    if yname == "servers" { return "Servers" }
    if yname == "lists" { return "Lists" }
    if yname == "ipv4-hosts" { return "Ipv4Hosts" }
    return ""
}

func (vrf *IpDomain_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *IpDomain_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-hosts" {
        return &vrf.Ipv6Hosts
    }
    if childYangName == "servers" {
        return &vrf.Servers
    }
    if childYangName == "lists" {
        return &vrf.Lists
    }
    if childYangName == "ipv4-hosts" {
        return &vrf.Ipv4Hosts
    }
    return nil
}

func (vrf *IpDomain_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6-hosts"] = &vrf.Ipv6Hosts
    children["servers"] = &vrf.Servers
    children["lists"] = &vrf.Lists
    children["ipv4-hosts"] = &vrf.Ipv4Hosts
    return children
}

func (vrf *IpDomain_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["lookup"] = vrf.Lookup
    leafs["multicast-domain"] = vrf.MulticastDomain
    leafs["source-interface"] = vrf.SourceInterface
    leafs["name"] = vrf.Name
    return leafs
}

func (vrf *IpDomain_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *IpDomain_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *IpDomain_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *IpDomain_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *IpDomain_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *IpDomain_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *IpDomain_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *IpDomain_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// IpDomain_Vrfs_Vrf_Ipv6Hosts
// IPv6 host
type IpDomain_Vrfs_Vrf_Ipv6Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Host name and up to 4 host IPv6 addresses. The type is slice of
    // IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host.
    Ipv6Host []IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host
}

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetFilter() yfilter.YFilter { return ipv6Hosts.YFilter }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) SetFilter(yf yfilter.YFilter) { ipv6Hosts.YFilter = yf }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetGoName(yname string) string {
    if yname == "ipv6-host" { return "Ipv6Host" }
    return ""
}

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetSegmentPath() string {
    return "ipv6-hosts"
}

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-host" {
        for _, c := range ipv6Hosts.Ipv6Host {
            if ipv6Hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host{}
        ipv6Hosts.Ipv6Host = append(ipv6Hosts.Ipv6Host, child)
        return &ipv6Hosts.Ipv6Host[len(ipv6Hosts.Ipv6Host)-1]
    }
    return nil
}

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6Hosts.Ipv6Host {
        children[ipv6Hosts.Ipv6Host[i].GetSegmentPath()] = &ipv6Hosts.Ipv6Host[i]
    }
    return children
}

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetYangName() string { return "ipv6-hosts" }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) SetParent(parent types.Entity) { ipv6Hosts.parent = parent }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetParent() types.Entity { return ipv6Hosts.parent }

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetParentYangName() string { return "vrf" }

// IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host
// Host name and up to 4 host IPv6 addresses
type IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A hostname. The type is string.
    HostName interface{}

    // Host IPv6 addresses. The type is slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address []interface{}
}

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetFilter() yfilter.YFilter { return ipv6Host.YFilter }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) SetFilter(yf yfilter.YFilter) { ipv6Host.YFilter = yf }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetGoName(yname string) string {
    if yname == "host-name" { return "HostName" }
    if yname == "address" { return "Address" }
    return ""
}

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetSegmentPath() string {
    return "ipv6-host" + "[host-name='" + fmt.Sprintf("%v", ipv6Host.HostName) + "']"
}

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name"] = ipv6Host.HostName
    leafs["address"] = ipv6Host.Address
    return leafs
}

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetYangName() string { return "ipv6-host" }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) SetParent(parent types.Entity) { ipv6Host.parent = parent }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetParent() types.Entity { return ipv6Host.parent }

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetParentYangName() string { return "ipv6-hosts" }

// IpDomain_Vrfs_Vrf_Servers
// Name server addresses
type IpDomain_Vrfs_Vrf_Servers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name server address. The type is slice of IpDomain_Vrfs_Vrf_Servers_Server.
    Server []IpDomain_Vrfs_Vrf_Servers_Server
}

func (servers *IpDomain_Vrfs_Vrf_Servers) GetFilter() yfilter.YFilter { return servers.YFilter }

func (servers *IpDomain_Vrfs_Vrf_Servers) SetFilter(yf yfilter.YFilter) { servers.YFilter = yf }

func (servers *IpDomain_Vrfs_Vrf_Servers) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    return ""
}

func (servers *IpDomain_Vrfs_Vrf_Servers) GetSegmentPath() string {
    return "servers"
}

func (servers *IpDomain_Vrfs_Vrf_Servers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        for _, c := range servers.Server {
            if servers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf_Servers_Server{}
        servers.Server = append(servers.Server, child)
        return &servers.Server[len(servers.Server)-1]
    }
    return nil
}

func (servers *IpDomain_Vrfs_Vrf_Servers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range servers.Server {
        children[servers.Server[i].GetSegmentPath()] = &servers.Server[i]
    }
    return children
}

func (servers *IpDomain_Vrfs_Vrf_Servers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servers *IpDomain_Vrfs_Vrf_Servers) GetBundleName() string { return "cisco_ios_xr" }

func (servers *IpDomain_Vrfs_Vrf_Servers) GetYangName() string { return "servers" }

func (servers *IpDomain_Vrfs_Vrf_Servers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servers *IpDomain_Vrfs_Vrf_Servers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servers *IpDomain_Vrfs_Vrf_Servers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servers *IpDomain_Vrfs_Vrf_Servers) SetParent(parent types.Entity) { servers.parent = parent }

func (servers *IpDomain_Vrfs_Vrf_Servers) GetParent() types.Entity { return servers.parent }

func (servers *IpDomain_Vrfs_Vrf_Servers) GetParentYangName() string { return "vrf" }

// IpDomain_Vrfs_Vrf_Servers_Server
// Name server address
type IpDomain_Vrfs_Vrf_Servers_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    Order interface{}

    // This attribute is a key. A name server address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}
}

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetGoName(yname string) string {
    if yname == "order" { return "Order" }
    if yname == "server-address" { return "ServerAddress" }
    return ""
}

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetSegmentPath() string {
    return "server" + "[order='" + fmt.Sprintf("%v", server.Order) + "']" + "[server-address='" + fmt.Sprintf("%v", server.ServerAddress) + "']"
}

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["order"] = server.Order
    leafs["server-address"] = server.ServerAddress
    return leafs
}

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetYangName() string { return "server" }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetParent() types.Entity { return server.parent }

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetParentYangName() string { return "servers" }

// IpDomain_Vrfs_Vrf_Lists
// Domain names to complete unqualified host
// names
type IpDomain_Vrfs_Vrf_Lists struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain name to complete unqualified host names. The type is slice of
    // IpDomain_Vrfs_Vrf_Lists_List.
    List []IpDomain_Vrfs_Vrf_Lists_List
}

func (lists *IpDomain_Vrfs_Vrf_Lists) GetFilter() yfilter.YFilter { return lists.YFilter }

func (lists *IpDomain_Vrfs_Vrf_Lists) SetFilter(yf yfilter.YFilter) { lists.YFilter = yf }

func (lists *IpDomain_Vrfs_Vrf_Lists) GetGoName(yname string) string {
    if yname == "list" { return "List" }
    return ""
}

func (lists *IpDomain_Vrfs_Vrf_Lists) GetSegmentPath() string {
    return "lists"
}

func (lists *IpDomain_Vrfs_Vrf_Lists) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "list" {
        for _, c := range lists.List {
            if lists.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf_Lists_List{}
        lists.List = append(lists.List, child)
        return &lists.List[len(lists.List)-1]
    }
    return nil
}

func (lists *IpDomain_Vrfs_Vrf_Lists) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lists.List {
        children[lists.List[i].GetSegmentPath()] = &lists.List[i]
    }
    return children
}

func (lists *IpDomain_Vrfs_Vrf_Lists) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lists *IpDomain_Vrfs_Vrf_Lists) GetBundleName() string { return "cisco_ios_xr" }

func (lists *IpDomain_Vrfs_Vrf_Lists) GetYangName() string { return "lists" }

func (lists *IpDomain_Vrfs_Vrf_Lists) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lists *IpDomain_Vrfs_Vrf_Lists) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lists *IpDomain_Vrfs_Vrf_Lists) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lists *IpDomain_Vrfs_Vrf_Lists) SetParent(parent types.Entity) { lists.parent = parent }

func (lists *IpDomain_Vrfs_Vrf_Lists) GetParent() types.Entity { return lists.parent }

func (lists *IpDomain_Vrfs_Vrf_Lists) GetParentYangName() string { return "vrf" }

// IpDomain_Vrfs_Vrf_Lists_List
// Domain name to complete unqualified host
// names
type IpDomain_Vrfs_Vrf_Lists_List struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the names in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    Order interface{}

    // This attribute is a key. A domain name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ListName interface{}
}

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetFilter() yfilter.YFilter { return list.YFilter }

func (list *IpDomain_Vrfs_Vrf_Lists_List) SetFilter(yf yfilter.YFilter) { list.YFilter = yf }

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetGoName(yname string) string {
    if yname == "order" { return "Order" }
    if yname == "list-name" { return "ListName" }
    return ""
}

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetSegmentPath() string {
    return "list" + "[order='" + fmt.Sprintf("%v", list.Order) + "']" + "[list-name='" + fmt.Sprintf("%v", list.ListName) + "']"
}

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["order"] = list.Order
    leafs["list-name"] = list.ListName
    return leafs
}

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetBundleName() string { return "cisco_ios_xr" }

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetYangName() string { return "list" }

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (list *IpDomain_Vrfs_Vrf_Lists_List) SetParent(parent types.Entity) { list.parent = parent }

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetParent() types.Entity { return list.parent }

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetParentYangName() string { return "lists" }

// IpDomain_Vrfs_Vrf_Ipv4Hosts
// IPv4 host
type IpDomain_Vrfs_Vrf_Ipv4Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Host name and up to 8 host IPv4 addresses. The type is slice of
    // IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host.
    Ipv4Host []IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host
}

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetFilter() yfilter.YFilter { return ipv4Hosts.YFilter }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) SetFilter(yf yfilter.YFilter) { ipv4Hosts.YFilter = yf }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetGoName(yname string) string {
    if yname == "ipv4-host" { return "Ipv4Host" }
    return ""
}

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetSegmentPath() string {
    return "ipv4-hosts"
}

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-host" {
        for _, c := range ipv4Hosts.Ipv4Host {
            if ipv4Hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host{}
        ipv4Hosts.Ipv4Host = append(ipv4Hosts.Ipv4Host, child)
        return &ipv4Hosts.Ipv4Host[len(ipv4Hosts.Ipv4Host)-1]
    }
    return nil
}

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Hosts.Ipv4Host {
        children[ipv4Hosts.Ipv4Host[i].GetSegmentPath()] = &ipv4Hosts.Ipv4Host[i]
    }
    return children
}

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetYangName() string { return "ipv4-hosts" }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) SetParent(parent types.Entity) { ipv4Hosts.parent = parent }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetParent() types.Entity { return ipv4Hosts.parent }

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetParentYangName() string { return "vrf" }

// IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host
// Host name and up to 8 host IPv4 addresses
type IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A hostname. The type is string.
    HostName interface{}

    // Host IPv4 addresses. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address []interface{}
}

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetFilter() yfilter.YFilter { return ipv4Host.YFilter }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) SetFilter(yf yfilter.YFilter) { ipv4Host.YFilter = yf }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetGoName(yname string) string {
    if yname == "host-name" { return "HostName" }
    if yname == "address" { return "Address" }
    return ""
}

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetSegmentPath() string {
    return "ipv4-host" + "[host-name='" + fmt.Sprintf("%v", ipv4Host.HostName) + "']"
}

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name"] = ipv4Host.HostName
    leafs["address"] = ipv4Host.Address
    return leafs
}

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetYangName() string { return "ipv4-host" }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) SetParent(parent types.Entity) { ipv4Host.parent = parent }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetParent() types.Entity { return ipv4Host.parent }

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetParentYangName() string { return "ipv4-hosts" }

