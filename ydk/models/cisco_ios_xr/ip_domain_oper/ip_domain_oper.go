// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-domain package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ip-domain: Domain server and host data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_domain_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_domain_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-domain-oper ip-domain}", reflect.TypeOf(IpDomain{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-domain-oper:ip-domain", reflect.TypeOf(IpDomain{}))
}

type HostAddressBase struct {
}

func (id HostAddressBase) String() string {
	return "Cisco-IOS-XR-ip-domain-oper-sub1:Host-address-base"
}

type Ipv4 struct {
}

func (id Ipv4) String() string {
	return "Cisco-IOS-XR-ip-domain-oper-sub1:ipv4"
}

type Ipv6 struct {
}

func (id Ipv6) String() string {
	return "Cisco-IOS-XR-ip-domain-oper-sub1:ipv6"
}

// ServerDomainLkup represents Domain look up
type ServerDomainLkup string

const (
    // Static mapping
    ServerDomainLkup_static_mapping ServerDomainLkup = "static-mapping"

    // Domain service
    ServerDomainLkup_domain_service ServerDomainLkup = "domain-service"
)

// IpDomain
// Domain server and host data
type IpDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of VRFs.
    Vrfs IpDomain_Vrfs
}

func (ipDomain *IpDomain) GetEntityData() *types.CommonEntityData {
    ipDomain.EntityData.YFilter = ipDomain.YFilter
    ipDomain.EntityData.YangName = "ip-domain"
    ipDomain.EntityData.BundleName = "cisco_ios_xr"
    ipDomain.EntityData.ParentYangName = "Cisco-IOS-XR-ip-domain-oper"
    ipDomain.EntityData.SegmentPath = "Cisco-IOS-XR-ip-domain-oper:ip-domain"
    ipDomain.EntityData.AbsolutePath = ipDomain.EntityData.SegmentPath
    ipDomain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDomain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDomain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDomain.EntityData.Children = types.NewOrderedMap()
    ipDomain.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &ipDomain.Vrfs})
    ipDomain.EntityData.Leafs = types.NewOrderedMap()

    ipDomain.EntityData.YListKeys = []string {}

    return &(ipDomain.EntityData)
}

// IpDomain_Vrfs
// List of VRFs
type IpDomain_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF instance. The type is slice of IpDomain_Vrfs_Vrf.
    Vrf []*IpDomain_Vrfs_Vrf
}

func (vrfs *IpDomain_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "ip-domain"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// IpDomain_Vrfs_Vrf
// VRF instance
type IpDomain_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Domain server data.
    Server IpDomain_Vrfs_Vrf_Server

    // List of domain hosts.
    Hosts IpDomain_Vrfs_Vrf_Hosts
}

func (vrf *IpDomain_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("server", types.YChild{"Server", &vrf.Server})
    vrf.EntityData.Children.Append("hosts", types.YChild{"Hosts", &vrf.Hosts})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// IpDomain_Vrfs_Vrf_Server
// Domain server data
type IpDomain_Vrfs_Vrf_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain lookup. The type is ServerDomainLkup.
    DomainLookup interface{}

    // Domain name. The type is string with length: 0..256.
    DomainName interface{}

    // Domain list. The type is slice of string with length: 0..256.
    Domain []interface{}

    // Server address list. The type is slice of
    // IpDomain_Vrfs_Vrf_Server_ServerAddress.
    ServerAddress []*IpDomain_Vrfs_Vrf_Server_ServerAddress
}

func (server *IpDomain_Vrfs_Vrf_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "vrf"
    server.EntityData.SegmentPath = "server"
    server.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/vrfs/vrf/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Children.Append("server-address", types.YChild{"ServerAddress", nil})
    for i := range server.ServerAddress {
        types.SetYListKey(server.ServerAddress[i], i)
        server.EntityData.Children.Append(types.GetSegmentPath(server.ServerAddress[i]), types.YChild{"ServerAddress", server.ServerAddress[i]})
    }
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("domain-lookup", types.YLeaf{"DomainLookup", server.DomainLookup})
    server.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", server.DomainName})
    server.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", server.Domain})

    server.EntityData.YListKeys = []string {}

    return &(server.EntityData)
}

// IpDomain_Vrfs_Vrf_Server_ServerAddress
// Server address list
type IpDomain_Vrfs_Vrf_Server_ServerAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AFName. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetEntityData() *types.CommonEntityData {
    serverAddress.EntityData.YFilter = serverAddress.YFilter
    serverAddress.EntityData.YangName = "server-address"
    serverAddress.EntityData.BundleName = "cisco_ios_xr"
    serverAddress.EntityData.ParentYangName = "server"
    serverAddress.EntityData.SegmentPath = "server-address" + types.AddNoKeyToken(serverAddress)
    serverAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/vrfs/vrf/server/" + serverAddress.EntityData.SegmentPath
    serverAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverAddress.EntityData.Children = types.NewOrderedMap()
    serverAddress.EntityData.Leafs = types.NewOrderedMap()
    serverAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", serverAddress.AfName})
    serverAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", serverAddress.Ipv4Address})
    serverAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", serverAddress.Ipv6Address})

    serverAddress.EntityData.YListKeys = []string {}

    return &(serverAddress.EntityData)
}

// IpDomain_Vrfs_Vrf_Hosts
// List of domain hosts
type IpDomain_Vrfs_Vrf_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP domain-name, lookup style, nameservers for specific host. The type is
    // slice of IpDomain_Vrfs_Vrf_Hosts_Host.
    Host []*IpDomain_Vrfs_Vrf_Hosts_Host
}

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "vrf"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/vrfs/vrf/" + hosts.EntityData.SegmentPath
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// IpDomain_Vrfs_Vrf_Hosts_Host
// IP domain-name, lookup style, nameservers for
// specific host
type IpDomain_Vrfs_Vrf_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Hostname. The type is string.
    HostName interface{}

    // Address type. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // Age in hours. The type is interface{} with range: 0..65535. Units are hour.
    Age interface{}

    // Host alias.
    HostAliasList IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList

    // Host address list. The type is slice of
    // IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress.
    HostAddress []*IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress
}

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.HostName, "host-name")
    host.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/vrfs/vrf/hosts/" + host.EntityData.SegmentPath
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Children.Append("host-alias-list", types.YChild{"HostAliasList", &host.HostAliasList})
    host.EntityData.Children.Append("host-address", types.YChild{"HostAddress", nil})
    for i := range host.HostAddress {
        types.SetYListKey(host.HostAddress[i], i)
        host.EntityData.Children.Append(types.GetSegmentPath(host.HostAddress[i]), types.YChild{"HostAddress", host.HostAddress[i]})
    }
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", host.HostName})
    host.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", host.AfName})
    host.EntityData.Leafs.Append("age", types.YLeaf{"Age", host.Age})

    host.EntityData.YListKeys = []string {"HostName"}

    return &(host.EntityData)
}

// IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList
// Host alias
type IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Host alias list. The type is slice of string with length: 0..256.
    HostAlias []interface{}
}

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetEntityData() *types.CommonEntityData {
    hostAliasList.EntityData.YFilter = hostAliasList.YFilter
    hostAliasList.EntityData.YangName = "host-alias-list"
    hostAliasList.EntityData.BundleName = "cisco_ios_xr"
    hostAliasList.EntityData.ParentYangName = "host"
    hostAliasList.EntityData.SegmentPath = "host-alias-list"
    hostAliasList.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/vrfs/vrf/hosts/host/" + hostAliasList.EntityData.SegmentPath
    hostAliasList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostAliasList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostAliasList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostAliasList.EntityData.Children = types.NewOrderedMap()
    hostAliasList.EntityData.Leafs = types.NewOrderedMap()
    hostAliasList.EntityData.Leafs.Append("host-alias", types.YLeaf{"HostAlias", hostAliasList.HostAlias})

    hostAliasList.EntityData.YListKeys = []string {}

    return &(hostAliasList.EntityData)
}

// IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress
// Host address list
type IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AFName. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetEntityData() *types.CommonEntityData {
    hostAddress.EntityData.YFilter = hostAddress.YFilter
    hostAddress.EntityData.YangName = "host-address"
    hostAddress.EntityData.BundleName = "cisco_ios_xr"
    hostAddress.EntityData.ParentYangName = "host"
    hostAddress.EntityData.SegmentPath = "host-address" + types.AddNoKeyToken(hostAddress)
    hostAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-domain-oper:ip-domain/vrfs/vrf/hosts/host/" + hostAddress.EntityData.SegmentPath
    hostAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostAddress.EntityData.Children = types.NewOrderedMap()
    hostAddress.EntityData.Leafs = types.NewOrderedMap()
    hostAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", hostAddress.AfName})
    hostAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", hostAddress.Ipv4Address})
    hostAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", hostAddress.Ipv6Address})

    hostAddress.EntityData.YListKeys = []string {}

    return &(hostAddress.EntityData)
}

