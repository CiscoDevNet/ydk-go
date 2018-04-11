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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs IpDomain_Vrfs
}

func (ipDomain *IpDomain) GetEntityData() *types.CommonEntityData {
    ipDomain.EntityData.YFilter = ipDomain.YFilter
    ipDomain.EntityData.YangName = "ip-domain"
    ipDomain.EntityData.BundleName = "cisco_ios_xr"
    ipDomain.EntityData.ParentYangName = "Cisco-IOS-XR-ip-domain-cfg"
    ipDomain.EntityData.SegmentPath = "Cisco-IOS-XR-ip-domain-cfg:ip-domain"
    ipDomain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDomain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDomain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDomain.EntityData.Children = make(map[string]types.YChild)
    ipDomain.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &ipDomain.Vrfs}
    ipDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipDomain.EntityData)
}

// IpDomain_Vrfs
// VRF table
type IpDomain_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of IpDomain_Vrfs_Vrf.
    Vrf []IpDomain_Vrfs_Vrf
}

func (vrfs *IpDomain_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "ip-domain"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// IpDomain_Vrfs_Vrf
// VRF specific data
type IpDomain_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Disable Domain Name System hostname translation. The type is interface{}.
    Lookup interface{}

    // Default multicast domain name. The type is string.
    MulticastDomain interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: b'[a-zA-Z0-9./-]+'.
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

func (vrf *IpDomain_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["ipv6-hosts"] = types.YChild{"Ipv6Hosts", &vrf.Ipv6Hosts}
    vrf.EntityData.Children["servers"] = types.YChild{"Servers", &vrf.Servers}
    vrf.EntityData.Children["lists"] = types.YChild{"Lists", &vrf.Lists}
    vrf.EntityData.Children["ipv4-hosts"] = types.YChild{"Ipv4Hosts", &vrf.Ipv4Hosts}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["lookup"] = types.YLeaf{"Lookup", vrf.Lookup}
    vrf.EntityData.Leafs["multicast-domain"] = types.YLeaf{"MulticastDomain", vrf.MulticastDomain}
    vrf.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", vrf.SourceInterface}
    vrf.EntityData.Leafs["name"] = types.YLeaf{"Name", vrf.Name}
    return &(vrf.EntityData)
}

// IpDomain_Vrfs_Vrf_Ipv6Hosts
// IPv6 host
type IpDomain_Vrfs_Vrf_Ipv6Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Host name and up to 4 host IPv6 addresses. The type is slice of
    // IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host.
    Ipv6Host []IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host
}

func (ipv6Hosts *IpDomain_Vrfs_Vrf_Ipv6Hosts) GetEntityData() *types.CommonEntityData {
    ipv6Hosts.EntityData.YFilter = ipv6Hosts.YFilter
    ipv6Hosts.EntityData.YangName = "ipv6-hosts"
    ipv6Hosts.EntityData.BundleName = "cisco_ios_xr"
    ipv6Hosts.EntityData.ParentYangName = "vrf"
    ipv6Hosts.EntityData.SegmentPath = "ipv6-hosts"
    ipv6Hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Hosts.EntityData.Children = make(map[string]types.YChild)
    ipv6Hosts.EntityData.Children["ipv6-host"] = types.YChild{"Ipv6Host", nil}
    for i := range ipv6Hosts.Ipv6Host {
        ipv6Hosts.EntityData.Children[types.GetSegmentPath(&ipv6Hosts.Ipv6Host[i])] = types.YChild{"Ipv6Host", &ipv6Hosts.Ipv6Host[i]}
    }
    ipv6Hosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Hosts.EntityData)
}

// IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host
// Host name and up to 4 host IPv6 addresses
type IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A hostname. The type is string.
    HostName interface{}

    // Host IPv6 addresses. The type is slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address []interface{}
}

func (ipv6Host *IpDomain_Vrfs_Vrf_Ipv6Hosts_Ipv6Host) GetEntityData() *types.CommonEntityData {
    ipv6Host.EntityData.YFilter = ipv6Host.YFilter
    ipv6Host.EntityData.YangName = "ipv6-host"
    ipv6Host.EntityData.BundleName = "cisco_ios_xr"
    ipv6Host.EntityData.ParentYangName = "ipv6-hosts"
    ipv6Host.EntityData.SegmentPath = "ipv6-host" + "[host-name='" + fmt.Sprintf("%v", ipv6Host.HostName) + "']"
    ipv6Host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Host.EntityData.Children = make(map[string]types.YChild)
    ipv6Host.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Host.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", ipv6Host.HostName}
    ipv6Host.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv6Host.Address}
    return &(ipv6Host.EntityData)
}

// IpDomain_Vrfs_Vrf_Servers
// Name server addresses
type IpDomain_Vrfs_Vrf_Servers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name server address. The type is slice of IpDomain_Vrfs_Vrf_Servers_Server.
    Server []IpDomain_Vrfs_Vrf_Servers_Server
}

func (servers *IpDomain_Vrfs_Vrf_Servers) GetEntityData() *types.CommonEntityData {
    servers.EntityData.YFilter = servers.YFilter
    servers.EntityData.YangName = "servers"
    servers.EntityData.BundleName = "cisco_ios_xr"
    servers.EntityData.ParentYangName = "vrf"
    servers.EntityData.SegmentPath = "servers"
    servers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servers.EntityData.Children = make(map[string]types.YChild)
    servers.EntityData.Children["server"] = types.YChild{"Server", nil}
    for i := range servers.Server {
        servers.EntityData.Children[types.GetSegmentPath(&servers.Server[i])] = types.YChild{"Server", &servers.Server[i]}
    }
    servers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(servers.EntityData)
}

// IpDomain_Vrfs_Vrf_Servers_Server
// Name server address
type IpDomain_Vrfs_Vrf_Servers_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    Order interface{}

    // This attribute is a key. A name server address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ServerAddress interface{}
}

func (server *IpDomain_Vrfs_Vrf_Servers_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "servers"
    server.EntityData.SegmentPath = "server" + "[order='" + fmt.Sprintf("%v", server.Order) + "']" + "[server-address='" + fmt.Sprintf("%v", server.ServerAddress) + "']"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    server.EntityData.Leafs["order"] = types.YLeaf{"Order", server.Order}
    server.EntityData.Leafs["server-address"] = types.YLeaf{"ServerAddress", server.ServerAddress}
    return &(server.EntityData)
}

// IpDomain_Vrfs_Vrf_Lists
// Domain names to complete unqualified host
// names
type IpDomain_Vrfs_Vrf_Lists struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain name to complete unqualified host names. The type is slice of
    // IpDomain_Vrfs_Vrf_Lists_List.
    List []IpDomain_Vrfs_Vrf_Lists_List
}

func (lists *IpDomain_Vrfs_Vrf_Lists) GetEntityData() *types.CommonEntityData {
    lists.EntityData.YFilter = lists.YFilter
    lists.EntityData.YangName = "lists"
    lists.EntityData.BundleName = "cisco_ios_xr"
    lists.EntityData.ParentYangName = "vrf"
    lists.EntityData.SegmentPath = "lists"
    lists.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lists.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lists.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lists.EntityData.Children = make(map[string]types.YChild)
    lists.EntityData.Children["list"] = types.YChild{"List", nil}
    for i := range lists.List {
        lists.EntityData.Children[types.GetSegmentPath(&lists.List[i])] = types.YChild{"List", &lists.List[i]}
    }
    lists.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lists.EntityData)
}

// IpDomain_Vrfs_Vrf_Lists_List
// Domain name to complete unqualified host
// names
type IpDomain_Vrfs_Vrf_Lists_List struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the names in the order of
    // precedence. The type is interface{} with range: -2147483648..2147483647.
    Order interface{}

    // This attribute is a key. A domain name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ListName interface{}
}

func (list *IpDomain_Vrfs_Vrf_Lists_List) GetEntityData() *types.CommonEntityData {
    list.EntityData.YFilter = list.YFilter
    list.EntityData.YangName = "list"
    list.EntityData.BundleName = "cisco_ios_xr"
    list.EntityData.ParentYangName = "lists"
    list.EntityData.SegmentPath = "list" + "[order='" + fmt.Sprintf("%v", list.Order) + "']" + "[list-name='" + fmt.Sprintf("%v", list.ListName) + "']"
    list.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    list.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    list.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    list.EntityData.Children = make(map[string]types.YChild)
    list.EntityData.Leafs = make(map[string]types.YLeaf)
    list.EntityData.Leafs["order"] = types.YLeaf{"Order", list.Order}
    list.EntityData.Leafs["list-name"] = types.YLeaf{"ListName", list.ListName}
    return &(list.EntityData)
}

// IpDomain_Vrfs_Vrf_Ipv4Hosts
// IPv4 host
type IpDomain_Vrfs_Vrf_Ipv4Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Host name and up to 8 host IPv4 addresses. The type is slice of
    // IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host.
    Ipv4Host []IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host
}

func (ipv4Hosts *IpDomain_Vrfs_Vrf_Ipv4Hosts) GetEntityData() *types.CommonEntityData {
    ipv4Hosts.EntityData.YFilter = ipv4Hosts.YFilter
    ipv4Hosts.EntityData.YangName = "ipv4-hosts"
    ipv4Hosts.EntityData.BundleName = "cisco_ios_xr"
    ipv4Hosts.EntityData.ParentYangName = "vrf"
    ipv4Hosts.EntityData.SegmentPath = "ipv4-hosts"
    ipv4Hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Hosts.EntityData.Children = make(map[string]types.YChild)
    ipv4Hosts.EntityData.Children["ipv4-host"] = types.YChild{"Ipv4Host", nil}
    for i := range ipv4Hosts.Ipv4Host {
        ipv4Hosts.EntityData.Children[types.GetSegmentPath(&ipv4Hosts.Ipv4Host[i])] = types.YChild{"Ipv4Host", &ipv4Hosts.Ipv4Host[i]}
    }
    ipv4Hosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Hosts.EntityData)
}

// IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host
// Host name and up to 8 host IPv4 addresses
type IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A hostname. The type is string.
    HostName interface{}

    // Host IPv4 addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address []interface{}
}

func (ipv4Host *IpDomain_Vrfs_Vrf_Ipv4Hosts_Ipv4Host) GetEntityData() *types.CommonEntityData {
    ipv4Host.EntityData.YFilter = ipv4Host.YFilter
    ipv4Host.EntityData.YangName = "ipv4-host"
    ipv4Host.EntityData.BundleName = "cisco_ios_xr"
    ipv4Host.EntityData.ParentYangName = "ipv4-hosts"
    ipv4Host.EntityData.SegmentPath = "ipv4-host" + "[host-name='" + fmt.Sprintf("%v", ipv4Host.HostName) + "']"
    ipv4Host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Host.EntityData.Children = make(map[string]types.YChild)
    ipv4Host.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Host.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", ipv4Host.HostName}
    ipv4Host.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4Host.Address}
    return &(ipv4Host.EntityData)
}

