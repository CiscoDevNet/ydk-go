// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-tcp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ip-tcp: Global IP TCP configuration
//   ip: ip
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_tcp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_tcp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-tcp-cfg ip-tcp}", reflect.TypeOf(IpTcp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-tcp-cfg:ip-tcp", reflect.TypeOf(IpTcp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-tcp-cfg ip}", reflect.TypeOf(Ip{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-tcp-cfg:ip", reflect.TypeOf(Ip{}))
}

// IpTcp
// Global IP TCP configuration
type IpTcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP connection accept rate. The type is interface{} with range: 1..1000.
    // The default value is 500.
    AcceptRate interface{}

    // Enable TCP selective-ACK. The type is interface{}.
    SelectiveAck interface{}

    // TCP receive window size (bytes). The type is interface{} with range:
    // 2048..65535. Units are byte.
    WindowSize interface{}

    // TCP receive Queue Size. The type is interface{} with range: 40..800.
    ReceiveQ interface{}

    // TCP initial maximum segment size. The type is interface{} with range:
    // 68..10000.
    MaximumSegmentSize interface{}

    // Time to wait on new TCP connections in seconds. The type is interface{}
    // with range: 5..30. Units are second.
    SynWaitTime interface{}

    // Enable TCP timestamp option. The type is interface{}.
    Timestamp interface{}

    // Aging time; 0 for infinite, and range be (10,30). The type is interface{}
    // with range: -2147483648..2147483647. Units are minute. The default value is
    // 10.
    PathMtuDiscovery interface{}

    // TCP directory details.
    Directory IpTcp_Directory

    // Throttle TCP receive buffer (in percentage).
    Throttle IpTcp_Throttle

    // TCP InQueue and OutQueue threads.
    NumThread IpTcp_NumThread
}

func (ipTcp *IpTcp) GetFilter() yfilter.YFilter { return ipTcp.YFilter }

func (ipTcp *IpTcp) SetFilter(yf yfilter.YFilter) { ipTcp.YFilter = yf }

func (ipTcp *IpTcp) GetGoName(yname string) string {
    if yname == "accept-rate" { return "AcceptRate" }
    if yname == "selective-ack" { return "SelectiveAck" }
    if yname == "window-size" { return "WindowSize" }
    if yname == "receive-q" { return "ReceiveQ" }
    if yname == "maximum-segment-size" { return "MaximumSegmentSize" }
    if yname == "syn-wait-time" { return "SynWaitTime" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "path-mtu-discovery" { return "PathMtuDiscovery" }
    if yname == "directory" { return "Directory" }
    if yname == "throttle" { return "Throttle" }
    if yname == "num-thread" { return "NumThread" }
    return ""
}

func (ipTcp *IpTcp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-tcp-cfg:ip-tcp"
}

func (ipTcp *IpTcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "directory" {
        return &ipTcp.Directory
    }
    if childYangName == "throttle" {
        return &ipTcp.Throttle
    }
    if childYangName == "num-thread" {
        return &ipTcp.NumThread
    }
    return nil
}

func (ipTcp *IpTcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["directory"] = &ipTcp.Directory
    children["throttle"] = &ipTcp.Throttle
    children["num-thread"] = &ipTcp.NumThread
    return children
}

func (ipTcp *IpTcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accept-rate"] = ipTcp.AcceptRate
    leafs["selective-ack"] = ipTcp.SelectiveAck
    leafs["window-size"] = ipTcp.WindowSize
    leafs["receive-q"] = ipTcp.ReceiveQ
    leafs["maximum-segment-size"] = ipTcp.MaximumSegmentSize
    leafs["syn-wait-time"] = ipTcp.SynWaitTime
    leafs["timestamp"] = ipTcp.Timestamp
    leafs["path-mtu-discovery"] = ipTcp.PathMtuDiscovery
    return leafs
}

func (ipTcp *IpTcp) GetBundleName() string { return "cisco_ios_xr" }

func (ipTcp *IpTcp) GetYangName() string { return "ip-tcp" }

func (ipTcp *IpTcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipTcp *IpTcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipTcp *IpTcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipTcp *IpTcp) SetParent(parent types.Entity) { ipTcp.parent = parent }

func (ipTcp *IpTcp) GetParent() types.Entity { return ipTcp.parent }

func (ipTcp *IpTcp) GetParentYangName() string { return "Cisco-IOS-XR-ip-tcp-cfg" }

// IpTcp_Directory
// TCP directory details
// This type is a presence type.
type IpTcp_Directory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Directory name . The type is string. This attribute is mandatory.
    Directoryname interface{}

    // Set number of Debug files. The type is interface{} with range: 1..10000.
    // This attribute is mandatory.
    MaxDebugFiles interface{}

    // Set size of debug files in bytes. The type is interface{} with range:
    // 1024..4294967295. This attribute is mandatory. Units are byte.
    MaxFileSizeFiles interface{}
}

func (directory *IpTcp_Directory) GetFilter() yfilter.YFilter { return directory.YFilter }

func (directory *IpTcp_Directory) SetFilter(yf yfilter.YFilter) { directory.YFilter = yf }

func (directory *IpTcp_Directory) GetGoName(yname string) string {
    if yname == "directoryname" { return "Directoryname" }
    if yname == "max-debug-files" { return "MaxDebugFiles" }
    if yname == "max-file-size-files" { return "MaxFileSizeFiles" }
    return ""
}

func (directory *IpTcp_Directory) GetSegmentPath() string {
    return "directory"
}

func (directory *IpTcp_Directory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (directory *IpTcp_Directory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (directory *IpTcp_Directory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["directoryname"] = directory.Directoryname
    leafs["max-debug-files"] = directory.MaxDebugFiles
    leafs["max-file-size-files"] = directory.MaxFileSizeFiles
    return leafs
}

func (directory *IpTcp_Directory) GetBundleName() string { return "cisco_ios_xr" }

func (directory *IpTcp_Directory) GetYangName() string { return "directory" }

func (directory *IpTcp_Directory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (directory *IpTcp_Directory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (directory *IpTcp_Directory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (directory *IpTcp_Directory) SetParent(parent types.Entity) { directory.parent = parent }

func (directory *IpTcp_Directory) GetParent() types.Entity { return directory.parent }

func (directory *IpTcp_Directory) GetParentYangName() string { return "ip-tcp" }

// IpTcp_Throttle
// Throttle TCP receive buffer (in percentage)
// This type is a presence type.
type IpTcp_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Min throttle. The type is interface{} with range: 0..100. This attribute is
    // mandatory.
    TcpminThrottle interface{}

    // Max throttle. The type is interface{} with range: 0..100. This attribute is
    // mandatory.
    Tcpmaxthrottle interface{}
}

func (throttle *IpTcp_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *IpTcp_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *IpTcp_Throttle) GetGoName(yname string) string {
    if yname == "tcpmin-throttle" { return "TcpminThrottle" }
    if yname == "tcpmaxthrottle" { return "Tcpmaxthrottle" }
    return ""
}

func (throttle *IpTcp_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *IpTcp_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *IpTcp_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *IpTcp_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcpmin-throttle"] = throttle.TcpminThrottle
    leafs["tcpmaxthrottle"] = throttle.Tcpmaxthrottle
    return leafs
}

func (throttle *IpTcp_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *IpTcp_Throttle) GetYangName() string { return "throttle" }

func (throttle *IpTcp_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *IpTcp_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *IpTcp_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *IpTcp_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *IpTcp_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *IpTcp_Throttle) GetParentYangName() string { return "ip-tcp" }

// IpTcp_NumThread
// TCP InQueue and OutQueue threads
// This type is a presence type.
type IpTcp_NumThread struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // InQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    TcpInQThreads interface{}

    // OutQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    TcpOutQThreads interface{}
}

func (numThread *IpTcp_NumThread) GetFilter() yfilter.YFilter { return numThread.YFilter }

func (numThread *IpTcp_NumThread) SetFilter(yf yfilter.YFilter) { numThread.YFilter = yf }

func (numThread *IpTcp_NumThread) GetGoName(yname string) string {
    if yname == "tcp-in-q-threads" { return "TcpInQThreads" }
    if yname == "tcp-out-q-threads" { return "TcpOutQThreads" }
    return ""
}

func (numThread *IpTcp_NumThread) GetSegmentPath() string {
    return "num-thread"
}

func (numThread *IpTcp_NumThread) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numThread *IpTcp_NumThread) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numThread *IpTcp_NumThread) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-in-q-threads"] = numThread.TcpInQThreads
    leafs["tcp-out-q-threads"] = numThread.TcpOutQThreads
    return leafs
}

func (numThread *IpTcp_NumThread) GetBundleName() string { return "cisco_ios_xr" }

func (numThread *IpTcp_NumThread) GetYangName() string { return "num-thread" }

func (numThread *IpTcp_NumThread) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numThread *IpTcp_NumThread) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numThread *IpTcp_NumThread) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numThread *IpTcp_NumThread) SetParent(parent types.Entity) { numThread.parent = parent }

func (numThread *IpTcp_NumThread) GetParent() types.Entity { return numThread.parent }

func (numThread *IpTcp_NumThread) GetParentYangName() string { return "ip-tcp" }

// Ip
// ip
type Ip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Cinetd configuration data.
    Cinetd Ip_Cinetd

    // Controls forwarding of physical and directed IP broadcasts.
    ForwardProtocol Ip_ForwardProtocol
}

func (ip *Ip) GetFilter() yfilter.YFilter { return ip.YFilter }

func (ip *Ip) SetFilter(yf yfilter.YFilter) { ip.YFilter = yf }

func (ip *Ip) GetGoName(yname string) string {
    if yname == "cinetd" { return "Cinetd" }
    if yname == "Cisco-IOS-XR-ip-udp-cfg:forward-protocol" { return "ForwardProtocol" }
    return ""
}

func (ip *Ip) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-tcp-cfg:ip"
}

func (ip *Ip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cinetd" {
        return &ip.Cinetd
    }
    if childYangName == "Cisco-IOS-XR-ip-udp-cfg:forward-protocol" {
        return &ip.ForwardProtocol
    }
    return nil
}

func (ip *Ip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cinetd"] = &ip.Cinetd
    children["Cisco-IOS-XR-ip-udp-cfg:forward-protocol"] = &ip.ForwardProtocol
    return children
}

func (ip *Ip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ip *Ip) GetBundleName() string { return "cisco_ios_xr" }

func (ip *Ip) GetYangName() string { return "ip" }

func (ip *Ip) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ip *Ip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ip *Ip) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ip *Ip) SetParent(parent types.Entity) { ip.parent = parent }

func (ip *Ip) GetParent() types.Entity { return ip.parent }

func (ip *Ip) GetParentYangName() string { return "Cisco-IOS-XR-ip-tcp-cfg" }

// Ip_Cinetd
// Cinetd configuration data
type Ip_Cinetd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of service requests accepted per second. The type is interface{}
    // with range: 1..100. The default value is 1.
    RateLimit interface{}

    // Describing services of cinetd.
    Services Ip_Cinetd_Services
}

func (cinetd *Ip_Cinetd) GetFilter() yfilter.YFilter { return cinetd.YFilter }

func (cinetd *Ip_Cinetd) SetFilter(yf yfilter.YFilter) { cinetd.YFilter = yf }

func (cinetd *Ip_Cinetd) GetGoName(yname string) string {
    if yname == "rate-limit" { return "RateLimit" }
    if yname == "services" { return "Services" }
    return ""
}

func (cinetd *Ip_Cinetd) GetSegmentPath() string {
    return "cinetd"
}

func (cinetd *Ip_Cinetd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "services" {
        return &cinetd.Services
    }
    return nil
}

func (cinetd *Ip_Cinetd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["services"] = &cinetd.Services
    return children
}

func (cinetd *Ip_Cinetd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rate-limit"] = cinetd.RateLimit
    return leafs
}

func (cinetd *Ip_Cinetd) GetBundleName() string { return "cisco_ios_xr" }

func (cinetd *Ip_Cinetd) GetYangName() string { return "cinetd" }

func (cinetd *Ip_Cinetd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cinetd *Ip_Cinetd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cinetd *Ip_Cinetd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cinetd *Ip_Cinetd) SetParent(parent types.Entity) { cinetd.parent = parent }

func (cinetd *Ip_Cinetd) GetParent() types.Entity { return cinetd.parent }

func (cinetd *Ip_Cinetd) GetParentYangName() string { return "ip" }

// Ip_Cinetd_Services
// Describing services of cinetd
type Ip_Cinetd_Services struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV4 related services.
    Ipv4 Ip_Cinetd_Services_Ipv4

    // VRF table.
    Vrfs Ip_Cinetd_Services_Vrfs

    // IPV6 related services.
    Ipv6 Ip_Cinetd_Services_Ipv6
}

func (services *Ip_Cinetd_Services) GetFilter() yfilter.YFilter { return services.YFilter }

func (services *Ip_Cinetd_Services) SetFilter(yf yfilter.YFilter) { services.YFilter = yf }

func (services *Ip_Cinetd_Services) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (services *Ip_Cinetd_Services) GetSegmentPath() string {
    return "services"
}

func (services *Ip_Cinetd_Services) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &services.Ipv4
    }
    if childYangName == "vrfs" {
        return &services.Vrfs
    }
    if childYangName == "ipv6" {
        return &services.Ipv6
    }
    return nil
}

func (services *Ip_Cinetd_Services) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &services.Ipv4
    children["vrfs"] = &services.Vrfs
    children["ipv6"] = &services.Ipv6
    return children
}

func (services *Ip_Cinetd_Services) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (services *Ip_Cinetd_Services) GetBundleName() string { return "cisco_ios_xr" }

func (services *Ip_Cinetd_Services) GetYangName() string { return "services" }

func (services *Ip_Cinetd_Services) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (services *Ip_Cinetd_Services) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (services *Ip_Cinetd_Services) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (services *Ip_Cinetd_Services) SetParent(parent types.Entity) { services.parent = parent }

func (services *Ip_Cinetd_Services) GetParent() types.Entity { return services.parent }

func (services *Ip_Cinetd_Services) GetParentYangName() string { return "cinetd" }

// Ip_Cinetd_Services_Ipv4
// IPV4 related services
type Ip_Cinetd_Services_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describing IPV4 and IPV6 small servers.
    SmallServers Ip_Cinetd_Services_Ipv4_SmallServers
}

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Ip_Cinetd_Services_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetGoName(yname string) string {
    if yname == "small-servers" { return "SmallServers" }
    return ""
}

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "small-servers" {
        return &ipv4.SmallServers
    }
    return nil
}

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["small-servers"] = &ipv4.SmallServers
    return children
}

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Ip_Cinetd_Services_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetParentYangName() string { return "services" }

// Ip_Cinetd_Services_Ipv4_SmallServers
// Describing IPV4 and IPV6 small servers
type Ip_Cinetd_Services_Ipv4_SmallServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describing TCP related IPV4 and IPV6 small servers.
    TcpSmallServers Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers

    // UDP small servers configuration.
    UdpSmallServers Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers
}

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetFilter() yfilter.YFilter { return smallServers.YFilter }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) SetFilter(yf yfilter.YFilter) { smallServers.YFilter = yf }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetGoName(yname string) string {
    if yname == "tcp-small-servers" { return "TcpSmallServers" }
    if yname == "Cisco-IOS-XR-ip-udp-cfg:udp-small-servers" { return "UdpSmallServers" }
    return ""
}

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetSegmentPath() string {
    return "small-servers"
}

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp-small-servers" {
        return &smallServers.TcpSmallServers
    }
    if childYangName == "Cisco-IOS-XR-ip-udp-cfg:udp-small-servers" {
        return &smallServers.UdpSmallServers
    }
    return nil
}

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcp-small-servers"] = &smallServers.TcpSmallServers
    children["Cisco-IOS-XR-ip-udp-cfg:udp-small-servers"] = &smallServers.UdpSmallServers
    return children
}

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetBundleName() string { return "cisco_ios_xr" }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetYangName() string { return "small-servers" }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) SetParent(parent types.Entity) { smallServers.parent = parent }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetParent() types.Entity { return smallServers.parent }

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetParentYangName() string { return "ipv4" }

// Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers
// Describing TCP related IPV4 and IPV6 small
// servers
// This type is a presence type.
type Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the access list. The type is string.
    AccessControlListName interface{}

    // Set number of allowable TCP small servers, specify 0 for no-limit. The type
    // is interface{} with range: 1..2147483647. This attribute is mandatory.
    SmallServer interface{}
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetFilter() yfilter.YFilter { return tcpSmallServers.YFilter }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) SetFilter(yf yfilter.YFilter) { tcpSmallServers.YFilter = yf }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetGoName(yname string) string {
    if yname == "access-control-list-name" { return "AccessControlListName" }
    if yname == "small-server" { return "SmallServer" }
    return ""
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetSegmentPath() string {
    return "tcp-small-servers"
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-control-list-name"] = tcpSmallServers.AccessControlListName
    leafs["small-server"] = tcpSmallServers.SmallServer
    return leafs
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetBundleName() string { return "cisco_ios_xr" }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetYangName() string { return "tcp-small-servers" }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) SetParent(parent types.Entity) { tcpSmallServers.parent = parent }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetParent() types.Entity { return tcpSmallServers.parent }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetParentYangName() string { return "small-servers" }

// Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers
// UDP small servers configuration
// This type is a presence type.
type Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the access list. The type is string.
    AccessControlListName interface{}

    // Set number of allowable small servers, specify 0 for no-limit. The type is
    // interface{} with range: 1..2147483647. This attribute is mandatory.
    SmallServer interface{}
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetFilter() yfilter.YFilter { return udpSmallServers.YFilter }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) SetFilter(yf yfilter.YFilter) { udpSmallServers.YFilter = yf }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetGoName(yname string) string {
    if yname == "access-control-list-name" { return "AccessControlListName" }
    if yname == "small-server" { return "SmallServer" }
    return ""
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-udp-cfg:udp-small-servers"
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-control-list-name"] = udpSmallServers.AccessControlListName
    leafs["small-server"] = udpSmallServers.SmallServer
    return leafs
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetBundleName() string { return "cisco_ios_xr" }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetYangName() string { return "udp-small-servers" }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) SetParent(parent types.Entity) { udpSmallServers.parent = parent }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetParent() types.Entity { return udpSmallServers.parent }

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetParentYangName() string { return "small-servers" }

// Ip_Cinetd_Services_Vrfs
// VRF table
type Ip_Cinetd_Services_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of Ip_Cinetd_Services_Vrfs_Vrf.
    Vrf []Ip_Cinetd_Services_Vrfs_Vrf
}

func (vrfs *Ip_Cinetd_Services_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ip_Cinetd_Services_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ip_Cinetd_Services_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ip_Cinetd_Services_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ip_Cinetd_Services_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ip_Cinetd_Services_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ip_Cinetd_Services_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ip_Cinetd_Services_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ip_Cinetd_Services_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ip_Cinetd_Services_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ip_Cinetd_Services_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ip_Cinetd_Services_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ip_Cinetd_Services_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ip_Cinetd_Services_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ip_Cinetd_Services_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ip_Cinetd_Services_Vrfs) GetParentYangName() string { return "services" }

// Ip_Cinetd_Services_Vrfs_Vrf
// VRF specific data
type Ip_Cinetd_Services_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPV6 related services.
    Ipv6 Ip_Cinetd_Services_Vrfs_Vrf_Ipv6

    // IPV4 related services.
    Ipv4 Ip_Cinetd_Services_Vrfs_Vrf_Ipv4
}

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &vrf.Ipv6
    }
    if childYangName == "ipv4" {
        return &vrf.Ipv4
    }
    return nil
}

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &vrf.Ipv6
    children["ipv4"] = &vrf.Ipv4
    return children
}

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6
// IPV6 related services
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TELNET server configuration commands.
    Telnet Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet

    // TFTP server configuration commands.
    Tftp Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp
}

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetGoName(yname string) string {
    if yname == "telnet" { return "Telnet" }
    if yname == "tftp" { return "Tftp" }
    return ""
}

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "telnet" {
        return &ipv6.Telnet
    }
    if childYangName == "tftp" {
        return &ipv6.Tftp
    }
    return nil
}

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["telnet"] = &ipv6.Telnet
    children["tftp"] = &ipv6.Tftp
    return children
}

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetParentYangName() string { return "vrf" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet
// TELNET server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP details.
    Tcp Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetFilter() yfilter.YFilter { return telnet.YFilter }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) SetFilter(yf yfilter.YFilter) { telnet.YFilter = yf }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetGoName(yname string) string {
    if yname == "tcp" { return "Tcp" }
    return ""
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetSegmentPath() string {
    return "telnet"
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp" {
        return &telnet.Tcp
    }
    return nil
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcp"] = &telnet.Tcp
    return children
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetBundleName() string { return "cisco_ios_xr" }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetYangName() string { return "telnet" }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) SetParent(parent types.Entity) { telnet.parent = parent }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetParent() types.Entity { return telnet.parent }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetParentYangName() string { return "ipv6" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp
// TCP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access list. The type is string.
    AccessListName interface{}

    // Set number of allowable servers. The type is interface{} with range:
    // 1..100. This attribute is mandatory.
    MaximumServer interface{}
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "maximum-server" { return "MaximumServer" }
    return ""
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = tcp.AccessListName
    leafs["maximum-server"] = tcp.MaximumServer
    return leafs
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetBundleName() string { return "cisco_ios_xr" }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetYangName() string { return "tcp" }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetParentYangName() string { return "telnet" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp
// TFTP server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDP details.
    Udp Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetFilter() yfilter.YFilter { return tftp.YFilter }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) SetFilter(yf yfilter.YFilter) { tftp.YFilter = yf }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetGoName(yname string) string {
    if yname == "udp" { return "Udp" }
    return ""
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetSegmentPath() string {
    return "tftp"
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udp" {
        return &tftp.Udp
    }
    return nil
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udp"] = &tftp.Udp
    return children
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetBundleName() string { return "cisco_ios_xr" }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetYangName() string { return "tftp" }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) SetParent(parent types.Entity) { tftp.parent = parent }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetParent() types.Entity { return tftp.parent }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetParentYangName() string { return "ipv6" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp
// UDP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access list. The type is string.
    AccessListName interface{}

    // Set number of allowable servers, 0 for no-limit. The type is interface{}
    // with range: 0..2147483647.
    MaximumServer interface{}

    // Specify device name where file is read from (e .g. flash:). The type is
    // string. This attribute is mandatory.
    HomeDirectory interface{}

    // Set IP DSCP (DiffServ CodePoint) for TFTP Server Packets. The type is
    // interface{} with range: -2147483648..2147483647.
    DscpValue interface{}
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "maximum-server" { return "MaximumServer" }
    if yname == "home-directory" { return "HomeDirectory" }
    if yname == "dscp-value" { return "DscpValue" }
    return ""
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetSegmentPath() string {
    return "udp"
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = udp.AccessListName
    leafs["maximum-server"] = udp.MaximumServer
    leafs["home-directory"] = udp.HomeDirectory
    leafs["dscp-value"] = udp.DscpValue
    return leafs
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetBundleName() string { return "cisco_ios_xr" }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetYangName() string { return "udp" }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetParent() types.Entity { return udp.parent }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetParentYangName() string { return "tftp" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4
// IPV4 related services
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TELNET server configuration commands.
    Telnet Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet

    // TFTP server configuration commands.
    Tftp Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp
}

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetGoName(yname string) string {
    if yname == "telnet" { return "Telnet" }
    if yname == "tftp" { return "Tftp" }
    return ""
}

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "telnet" {
        return &ipv4.Telnet
    }
    if childYangName == "tftp" {
        return &ipv4.Tftp
    }
    return nil
}

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["telnet"] = &ipv4.Telnet
    children["tftp"] = &ipv4.Tftp
    return children
}

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetParentYangName() string { return "vrf" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet
// TELNET server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP details.
    Tcp Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetFilter() yfilter.YFilter { return telnet.YFilter }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) SetFilter(yf yfilter.YFilter) { telnet.YFilter = yf }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetGoName(yname string) string {
    if yname == "tcp" { return "Tcp" }
    return ""
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetSegmentPath() string {
    return "telnet"
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp" {
        return &telnet.Tcp
    }
    return nil
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcp"] = &telnet.Tcp
    return children
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetBundleName() string { return "cisco_ios_xr" }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetYangName() string { return "telnet" }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) SetParent(parent types.Entity) { telnet.parent = parent }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetParent() types.Entity { return telnet.parent }

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetParentYangName() string { return "ipv4" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp
// TCP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access list. The type is string.
    AccessListName interface{}

    // Set number of allowable servers. The type is interface{} with range:
    // 1..100. This attribute is mandatory.
    MaximumServer interface{}
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "maximum-server" { return "MaximumServer" }
    return ""
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = tcp.AccessListName
    leafs["maximum-server"] = tcp.MaximumServer
    return leafs
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetBundleName() string { return "cisco_ios_xr" }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetYangName() string { return "tcp" }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetParentYangName() string { return "telnet" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp
// TFTP server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDP details.
    Udp Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetFilter() yfilter.YFilter { return tftp.YFilter }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) SetFilter(yf yfilter.YFilter) { tftp.YFilter = yf }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetGoName(yname string) string {
    if yname == "udp" { return "Udp" }
    return ""
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetSegmentPath() string {
    return "tftp"
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udp" {
        return &tftp.Udp
    }
    return nil
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udp"] = &tftp.Udp
    return children
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetBundleName() string { return "cisco_ios_xr" }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetYangName() string { return "tftp" }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) SetParent(parent types.Entity) { tftp.parent = parent }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetParent() types.Entity { return tftp.parent }

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetParentYangName() string { return "ipv4" }

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp
// UDP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access list. The type is string.
    AccessListName interface{}

    // Set number of allowable servers, 0 for no-limit. The type is interface{}
    // with range: 0..2147483647.
    MaximumServer interface{}

    // Specify device name where file is read from (e .g. flash:). The type is
    // string. This attribute is mandatory.
    HomeDirectory interface{}

    // Set IP DSCP (DiffServ CodePoint) for TFTP Server Packets. The type is
    // interface{} with range: -2147483648..2147483647.
    DscpValue interface{}
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "maximum-server" { return "MaximumServer" }
    if yname == "home-directory" { return "HomeDirectory" }
    if yname == "dscp-value" { return "DscpValue" }
    return ""
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetSegmentPath() string {
    return "udp"
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = udp.AccessListName
    leafs["maximum-server"] = udp.MaximumServer
    leafs["home-directory"] = udp.HomeDirectory
    leafs["dscp-value"] = udp.DscpValue
    return leafs
}

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetBundleName() string { return "cisco_ios_xr" }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetYangName() string { return "udp" }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetParent() types.Entity { return udp.parent }

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetParentYangName() string { return "tftp" }

// Ip_Cinetd_Services_Ipv6
// IPV6 related services
type Ip_Cinetd_Services_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describing IPV4 and IPV6 small servers.
    SmallServers Ip_Cinetd_Services_Ipv6_SmallServers
}

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Ip_Cinetd_Services_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetGoName(yname string) string {
    if yname == "small-servers" { return "SmallServers" }
    return ""
}

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "small-servers" {
        return &ipv6.SmallServers
    }
    return nil
}

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["small-servers"] = &ipv6.SmallServers
    return children
}

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Ip_Cinetd_Services_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetParentYangName() string { return "services" }

// Ip_Cinetd_Services_Ipv6_SmallServers
// Describing IPV4 and IPV6 small servers
type Ip_Cinetd_Services_Ipv6_SmallServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describing TCP related IPV4 and IPV6 small servers.
    TcpSmallServers Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers
}

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetFilter() yfilter.YFilter { return smallServers.YFilter }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) SetFilter(yf yfilter.YFilter) { smallServers.YFilter = yf }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetGoName(yname string) string {
    if yname == "tcp-small-servers" { return "TcpSmallServers" }
    return ""
}

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetSegmentPath() string {
    return "small-servers"
}

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp-small-servers" {
        return &smallServers.TcpSmallServers
    }
    return nil
}

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcp-small-servers"] = &smallServers.TcpSmallServers
    return children
}

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetBundleName() string { return "cisco_ios_xr" }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetYangName() string { return "small-servers" }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) SetParent(parent types.Entity) { smallServers.parent = parent }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetParent() types.Entity { return smallServers.parent }

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetParentYangName() string { return "ipv6" }

// Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers
// Describing TCP related IPV4 and IPV6 small
// servers
// This type is a presence type.
type Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the access list. The type is string.
    AccessControlListName interface{}

    // Set number of allowable TCP small servers, specify 0 for no-limit. The type
    // is interface{} with range: 1..2147483647. This attribute is mandatory.
    SmallServer interface{}
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetFilter() yfilter.YFilter { return tcpSmallServers.YFilter }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) SetFilter(yf yfilter.YFilter) { tcpSmallServers.YFilter = yf }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetGoName(yname string) string {
    if yname == "access-control-list-name" { return "AccessControlListName" }
    if yname == "small-server" { return "SmallServer" }
    return ""
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetSegmentPath() string {
    return "tcp-small-servers"
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-control-list-name"] = tcpSmallServers.AccessControlListName
    leafs["small-server"] = tcpSmallServers.SmallServer
    return leafs
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetBundleName() string { return "cisco_ios_xr" }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetYangName() string { return "tcp-small-servers" }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) SetParent(parent types.Entity) { tcpSmallServers.parent = parent }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetParent() types.Entity { return tcpSmallServers.parent }

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetParentYangName() string { return "small-servers" }

// Ip_ForwardProtocol
// Controls forwarding of physical and directed IP
// broadcasts
type Ip_ForwardProtocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packets to a specific UDP port.
    Udp Ip_ForwardProtocol_Udp
}

func (forwardProtocol *Ip_ForwardProtocol) GetFilter() yfilter.YFilter { return forwardProtocol.YFilter }

func (forwardProtocol *Ip_ForwardProtocol) SetFilter(yf yfilter.YFilter) { forwardProtocol.YFilter = yf }

func (forwardProtocol *Ip_ForwardProtocol) GetGoName(yname string) string {
    if yname == "udp" { return "Udp" }
    return ""
}

func (forwardProtocol *Ip_ForwardProtocol) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-udp-cfg:forward-protocol"
}

func (forwardProtocol *Ip_ForwardProtocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udp" {
        return &forwardProtocol.Udp
    }
    return nil
}

func (forwardProtocol *Ip_ForwardProtocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udp"] = &forwardProtocol.Udp
    return children
}

func (forwardProtocol *Ip_ForwardProtocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwardProtocol *Ip_ForwardProtocol) GetBundleName() string { return "cisco_ios_xr" }

func (forwardProtocol *Ip_ForwardProtocol) GetYangName() string { return "forward-protocol" }

func (forwardProtocol *Ip_ForwardProtocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardProtocol *Ip_ForwardProtocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardProtocol *Ip_ForwardProtocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardProtocol *Ip_ForwardProtocol) SetParent(parent types.Entity) { forwardProtocol.parent = parent }

func (forwardProtocol *Ip_ForwardProtocol) GetParent() types.Entity { return forwardProtocol.parent }

func (forwardProtocol *Ip_ForwardProtocol) GetParentYangName() string { return "ip" }

// Ip_ForwardProtocol_Udp
// Packets to a specific UDP port
type Ip_ForwardProtocol_Udp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable IP Forward Protocol UDP. The type is interface{}.
    Disable interface{}

    // Port configuration.
    Ports Ip_ForwardProtocol_Udp_Ports
}

func (udp *Ip_ForwardProtocol_Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *Ip_ForwardProtocol_Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *Ip_ForwardProtocol_Udp) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "ports" { return "Ports" }
    return ""
}

func (udp *Ip_ForwardProtocol_Udp) GetSegmentPath() string {
    return "udp"
}

func (udp *Ip_ForwardProtocol_Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ports" {
        return &udp.Ports
    }
    return nil
}

func (udp *Ip_ForwardProtocol_Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ports"] = &udp.Ports
    return children
}

func (udp *Ip_ForwardProtocol_Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = udp.Disable
    return leafs
}

func (udp *Ip_ForwardProtocol_Udp) GetBundleName() string { return "cisco_ios_xr" }

func (udp *Ip_ForwardProtocol_Udp) GetYangName() string { return "udp" }

func (udp *Ip_ForwardProtocol_Udp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udp *Ip_ForwardProtocol_Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udp *Ip_ForwardProtocol_Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udp *Ip_ForwardProtocol_Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *Ip_ForwardProtocol_Udp) GetParent() types.Entity { return udp.parent }

func (udp *Ip_ForwardProtocol_Udp) GetParentYangName() string { return "forward-protocol" }

// Ip_ForwardProtocol_Udp_Ports
// Port configuration
type Ip_ForwardProtocol_Udp_Ports struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Well-known ports are enabled by default and non well-known ports are
    // disabled by default. It is not allowed to configure the default. The type
    // is slice of Ip_ForwardProtocol_Udp_Ports_Port.
    Port []Ip_ForwardProtocol_Udp_Ports_Port
}

func (ports *Ip_ForwardProtocol_Udp_Ports) GetFilter() yfilter.YFilter { return ports.YFilter }

func (ports *Ip_ForwardProtocol_Udp_Ports) SetFilter(yf yfilter.YFilter) { ports.YFilter = yf }

func (ports *Ip_ForwardProtocol_Udp_Ports) GetGoName(yname string) string {
    if yname == "port" { return "Port" }
    return ""
}

func (ports *Ip_ForwardProtocol_Udp_Ports) GetSegmentPath() string {
    return "ports"
}

func (ports *Ip_ForwardProtocol_Udp_Ports) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port" {
        for _, c := range ports.Port {
            if ports.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ip_ForwardProtocol_Udp_Ports_Port{}
        ports.Port = append(ports.Port, child)
        return &ports.Port[len(ports.Port)-1]
    }
    return nil
}

func (ports *Ip_ForwardProtocol_Udp_Ports) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ports.Port {
        children[ports.Port[i].GetSegmentPath()] = &ports.Port[i]
    }
    return children
}

func (ports *Ip_ForwardProtocol_Udp_Ports) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ports *Ip_ForwardProtocol_Udp_Ports) GetBundleName() string { return "cisco_ios_xr" }

func (ports *Ip_ForwardProtocol_Udp_Ports) GetYangName() string { return "ports" }

func (ports *Ip_ForwardProtocol_Udp_Ports) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ports *Ip_ForwardProtocol_Udp_Ports) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ports *Ip_ForwardProtocol_Udp_Ports) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ports *Ip_ForwardProtocol_Udp_Ports) SetParent(parent types.Entity) { ports.parent = parent }

func (ports *Ip_ForwardProtocol_Udp_Ports) GetParent() types.Entity { return ports.parent }

func (ports *Ip_ForwardProtocol_Udp_Ports) GetParentYangName() string { return "udp" }

// Ip_ForwardProtocol_Udp_Ports_Port
// Well-known ports are enabled by default and
// non well-known ports are disabled by default.
// It is not allowed to configure the default.
type Ip_ForwardProtocol_Udp_Ports_Port struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port number. The type is interface{} with range:
    // 1..65535.
    PortId interface{}

    // Specify 'false' to disable well-known ports Domain (53), TFTP (69),
    // NameServer (42), TACACS (49), NetBiosNameService (137), or
    // NetBiosDatagramService (138).  Specify 'true' to enable non well-known
    // ports. The type is bool. This attribute is mandatory.
    Enable interface{}
}

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetFilter() yfilter.YFilter { return port.YFilter }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) SetFilter(yf yfilter.YFilter) { port.YFilter = yf }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetGoName(yname string) string {
    if yname == "port-id" { return "PortId" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetSegmentPath() string {
    return "port" + "[port-id='" + fmt.Sprintf("%v", port.PortId) + "']"
}

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-id"] = port.PortId
    leafs["enable"] = port.Enable
    return leafs
}

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetBundleName() string { return "cisco_ios_xr" }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetYangName() string { return "port" }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) SetParent(parent types.Entity) { port.parent = parent }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetParent() types.Entity { return port.parent }

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetParentYangName() string { return "ports" }

