// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-domain package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ip-domain: Domain server and host data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of VRFs.
    Vrfs IpDomain_Vrfs
}

func (ipDomain *IpDomain) GetFilter() yfilter.YFilter { return ipDomain.YFilter }

func (ipDomain *IpDomain) SetFilter(yf yfilter.YFilter) { ipDomain.YFilter = yf }

func (ipDomain *IpDomain) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (ipDomain *IpDomain) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-domain-oper:ip-domain"
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

func (ipDomain *IpDomain) GetParentYangName() string { return "Cisco-IOS-XR-ip-domain-oper" }

// IpDomain_Vrfs
// List of VRFs
type IpDomain_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF instance. The type is slice of IpDomain_Vrfs_Vrf.
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
// VRF instance
type IpDomain_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Domain server data.
    Server IpDomain_Vrfs_Vrf_Server

    // List of domain hosts.
    Hosts IpDomain_Vrfs_Vrf_Hosts
}

func (vrf *IpDomain_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *IpDomain_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *IpDomain_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "server" { return "Server" }
    if yname == "hosts" { return "Hosts" }
    return ""
}

func (vrf *IpDomain_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *IpDomain_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        return &vrf.Server
    }
    if childYangName == "hosts" {
        return &vrf.Hosts
    }
    return nil
}

func (vrf *IpDomain_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["server"] = &vrf.Server
    children["hosts"] = &vrf.Hosts
    return children
}

func (vrf *IpDomain_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
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

// IpDomain_Vrfs_Vrf_Server
// Domain server data
type IpDomain_Vrfs_Vrf_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain lookup. The type is ServerDomainLkup.
    DomainLookup interface{}

    // Domain name. The type is string with length: 0..256.
    DomainName interface{}

    // Domain list. The type is slice of string with length: 0..256.
    Domain []interface{}

    // Server address list. The type is slice of
    // IpDomain_Vrfs_Vrf_Server_ServerAddress.
    ServerAddress []IpDomain_Vrfs_Vrf_Server_ServerAddress
}

func (server *IpDomain_Vrfs_Vrf_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *IpDomain_Vrfs_Vrf_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *IpDomain_Vrfs_Vrf_Server) GetGoName(yname string) string {
    if yname == "domain-lookup" { return "DomainLookup" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "domain" { return "Domain" }
    if yname == "server-address" { return "ServerAddress" }
    return ""
}

func (server *IpDomain_Vrfs_Vrf_Server) GetSegmentPath() string {
    return "server"
}

func (server *IpDomain_Vrfs_Vrf_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server-address" {
        for _, c := range server.ServerAddress {
            if server.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf_Server_ServerAddress{}
        server.ServerAddress = append(server.ServerAddress, child)
        return &server.ServerAddress[len(server.ServerAddress)-1]
    }
    return nil
}

func (server *IpDomain_Vrfs_Vrf_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range server.ServerAddress {
        children[server.ServerAddress[i].GetSegmentPath()] = &server.ServerAddress[i]
    }
    return children
}

func (server *IpDomain_Vrfs_Vrf_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-lookup"] = server.DomainLookup
    leafs["domain-name"] = server.DomainName
    leafs["domain"] = server.Domain
    return leafs
}

func (server *IpDomain_Vrfs_Vrf_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *IpDomain_Vrfs_Vrf_Server) GetYangName() string { return "server" }

func (server *IpDomain_Vrfs_Vrf_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *IpDomain_Vrfs_Vrf_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *IpDomain_Vrfs_Vrf_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *IpDomain_Vrfs_Vrf_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *IpDomain_Vrfs_Vrf_Server) GetParent() types.Entity { return server.parent }

func (server *IpDomain_Vrfs_Vrf_Server) GetParentYangName() string { return "vrf" }

// IpDomain_Vrfs_Vrf_Server_ServerAddress
// Server address list
type IpDomain_Vrfs_Vrf_Server_ServerAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetFilter() yfilter.YFilter { return serverAddress.YFilter }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) SetFilter(yf yfilter.YFilter) { serverAddress.YFilter = yf }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetSegmentPath() string {
    return "server-address"
}

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = serverAddress.AfName
    leafs["ipv4-address"] = serverAddress.Ipv4Address
    leafs["ipv6-address"] = serverAddress.Ipv6Address
    return leafs
}

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetBundleName() string { return "cisco_ios_xr" }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetYangName() string { return "server-address" }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) SetParent(parent types.Entity) { serverAddress.parent = parent }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetParent() types.Entity { return serverAddress.parent }

func (serverAddress *IpDomain_Vrfs_Vrf_Server_ServerAddress) GetParentYangName() string { return "server" }

// IpDomain_Vrfs_Vrf_Hosts
// List of domain hosts
type IpDomain_Vrfs_Vrf_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP domain-name, lookup style, nameservers for specific host. The type is
    // slice of IpDomain_Vrfs_Vrf_Hosts_Host.
    Host []IpDomain_Vrfs_Vrf_Hosts_Host
}

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetYangName() string { return "hosts" }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *IpDomain_Vrfs_Vrf_Hosts) GetParentYangName() string { return "vrf" }

// IpDomain_Vrfs_Vrf_Hosts_Host
// IP domain-name, lookup style, nameservers for
// specific host
type IpDomain_Vrfs_Vrf_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    HostAddress []IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress
}

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetGoName(yname string) string {
    if yname == "host-name" { return "HostName" }
    if yname == "af-name" { return "AfName" }
    if yname == "age" { return "Age" }
    if yname == "host-alias-list" { return "HostAliasList" }
    if yname == "host-address" { return "HostAddress" }
    return ""
}

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetSegmentPath() string {
    return "host" + "[host-name='" + fmt.Sprintf("%v", host.HostName) + "']"
}

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host-alias-list" {
        return &host.HostAliasList
    }
    if childYangName == "host-address" {
        for _, c := range host.HostAddress {
            if host.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress{}
        host.HostAddress = append(host.HostAddress, child)
        return &host.HostAddress[len(host.HostAddress)-1]
    }
    return nil
}

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["host-alias-list"] = &host.HostAliasList
    for i := range host.HostAddress {
        children[host.HostAddress[i].GetSegmentPath()] = &host.HostAddress[i]
    }
    return children
}

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name"] = host.HostName
    leafs["af-name"] = host.AfName
    leafs["age"] = host.Age
    return leafs
}

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetYangName() string { return "host" }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *IpDomain_Vrfs_Vrf_Hosts_Host) GetParentYangName() string { return "hosts" }

// IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList
// Host alias
type IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Host alias list. The type is slice of string with length: 0..256.
    HostAlias []interface{}
}

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetFilter() yfilter.YFilter { return hostAliasList.YFilter }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) SetFilter(yf yfilter.YFilter) { hostAliasList.YFilter = yf }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetGoName(yname string) string {
    if yname == "host-alias" { return "HostAlias" }
    return ""
}

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetSegmentPath() string {
    return "host-alias-list"
}

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-alias"] = hostAliasList.HostAlias
    return leafs
}

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetBundleName() string { return "cisco_ios_xr" }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetYangName() string { return "host-alias-list" }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) SetParent(parent types.Entity) { hostAliasList.parent = parent }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetParent() types.Entity { return hostAliasList.parent }

func (hostAliasList *IpDomain_Vrfs_Vrf_Hosts_Host_HostAliasList) GetParentYangName() string { return "host" }

// IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress
// Host address list
type IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetFilter() yfilter.YFilter { return hostAddress.YFilter }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) SetFilter(yf yfilter.YFilter) { hostAddress.YFilter = yf }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetSegmentPath() string {
    return "host-address"
}

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = hostAddress.AfName
    leafs["ipv4-address"] = hostAddress.Ipv4Address
    leafs["ipv6-address"] = hostAddress.Ipv6Address
    return leafs
}

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetBundleName() string { return "cisco_ios_xr" }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetYangName() string { return "host-address" }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) SetParent(parent types.Entity) { hostAddress.parent = parent }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetParent() types.Entity { return hostAddress.parent }

func (hostAddress *IpDomain_Vrfs_Vrf_Hosts_Host_HostAddress) GetParentYangName() string { return "host" }

