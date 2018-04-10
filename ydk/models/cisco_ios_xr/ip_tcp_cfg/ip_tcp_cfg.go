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
    EntityData types.CommonEntityData
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

func (ipTcp *IpTcp) GetEntityData() *types.CommonEntityData {
    ipTcp.EntityData.YFilter = ipTcp.YFilter
    ipTcp.EntityData.YangName = "ip-tcp"
    ipTcp.EntityData.BundleName = "cisco_ios_xr"
    ipTcp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-tcp-cfg"
    ipTcp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-tcp-cfg:ip-tcp"
    ipTcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipTcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipTcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipTcp.EntityData.Children = make(map[string]types.YChild)
    ipTcp.EntityData.Children["directory"] = types.YChild{"Directory", &ipTcp.Directory}
    ipTcp.EntityData.Children["throttle"] = types.YChild{"Throttle", &ipTcp.Throttle}
    ipTcp.EntityData.Children["num-thread"] = types.YChild{"NumThread", &ipTcp.NumThread}
    ipTcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipTcp.EntityData.Leafs["accept-rate"] = types.YLeaf{"AcceptRate", ipTcp.AcceptRate}
    ipTcp.EntityData.Leafs["selective-ack"] = types.YLeaf{"SelectiveAck", ipTcp.SelectiveAck}
    ipTcp.EntityData.Leafs["window-size"] = types.YLeaf{"WindowSize", ipTcp.WindowSize}
    ipTcp.EntityData.Leafs["receive-q"] = types.YLeaf{"ReceiveQ", ipTcp.ReceiveQ}
    ipTcp.EntityData.Leafs["maximum-segment-size"] = types.YLeaf{"MaximumSegmentSize", ipTcp.MaximumSegmentSize}
    ipTcp.EntityData.Leafs["syn-wait-time"] = types.YLeaf{"SynWaitTime", ipTcp.SynWaitTime}
    ipTcp.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", ipTcp.Timestamp}
    ipTcp.EntityData.Leafs["path-mtu-discovery"] = types.YLeaf{"PathMtuDiscovery", ipTcp.PathMtuDiscovery}
    return &(ipTcp.EntityData)
}

// IpTcp_Directory
// TCP directory details
// This type is a presence type.
type IpTcp_Directory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Directory name . The type is string. This attribute is mandatory.
    Directoryname interface{}

    // Set number of Debug files. The type is interface{} with range: 1..10000.
    // The default value is 256.
    MaxDebugFiles interface{}

    // Set size of debug files in bytes. The type is interface{} with range:
    // 1024..4294967295. Units are byte.
    MaxFileSizeFiles interface{}
}

func (directory *IpTcp_Directory) GetEntityData() *types.CommonEntityData {
    directory.EntityData.YFilter = directory.YFilter
    directory.EntityData.YangName = "directory"
    directory.EntityData.BundleName = "cisco_ios_xr"
    directory.EntityData.ParentYangName = "ip-tcp"
    directory.EntityData.SegmentPath = "directory"
    directory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directory.EntityData.Children = make(map[string]types.YChild)
    directory.EntityData.Leafs = make(map[string]types.YLeaf)
    directory.EntityData.Leafs["directoryname"] = types.YLeaf{"Directoryname", directory.Directoryname}
    directory.EntityData.Leafs["max-debug-files"] = types.YLeaf{"MaxDebugFiles", directory.MaxDebugFiles}
    directory.EntityData.Leafs["max-file-size-files"] = types.YLeaf{"MaxFileSizeFiles", directory.MaxFileSizeFiles}
    return &(directory.EntityData)
}

// IpTcp_Throttle
// Throttle TCP receive buffer (in percentage)
// This type is a presence type.
type IpTcp_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Min throttle. The type is interface{} with range: 0..100. This attribute is
    // mandatory.
    TcpminThrottle interface{}

    // Max throttle. The type is interface{} with range: 0..100. This attribute is
    // mandatory.
    Tcpmaxthrottle interface{}
}

func (throttle *IpTcp_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "ip-tcp"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = make(map[string]types.YChild)
    throttle.EntityData.Leafs = make(map[string]types.YLeaf)
    throttle.EntityData.Leafs["tcpmin-throttle"] = types.YLeaf{"TcpminThrottle", throttle.TcpminThrottle}
    throttle.EntityData.Leafs["tcpmaxthrottle"] = types.YLeaf{"Tcpmaxthrottle", throttle.Tcpmaxthrottle}
    return &(throttle.EntityData)
}

// IpTcp_NumThread
// TCP InQueue and OutQueue threads
// This type is a presence type.
type IpTcp_NumThread struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // InQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    TcpInQThreads interface{}

    // OutQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    TcpOutQThreads interface{}
}

func (numThread *IpTcp_NumThread) GetEntityData() *types.CommonEntityData {
    numThread.EntityData.YFilter = numThread.YFilter
    numThread.EntityData.YangName = "num-thread"
    numThread.EntityData.BundleName = "cisco_ios_xr"
    numThread.EntityData.ParentYangName = "ip-tcp"
    numThread.EntityData.SegmentPath = "num-thread"
    numThread.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numThread.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numThread.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numThread.EntityData.Children = make(map[string]types.YChild)
    numThread.EntityData.Leafs = make(map[string]types.YLeaf)
    numThread.EntityData.Leafs["tcp-in-q-threads"] = types.YLeaf{"TcpInQThreads", numThread.TcpInQThreads}
    numThread.EntityData.Leafs["tcp-out-q-threads"] = types.YLeaf{"TcpOutQThreads", numThread.TcpOutQThreads}
    return &(numThread.EntityData)
}

// Ip
// ip
type Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Cinetd configuration data.
    Cinetd Ip_Cinetd

    // Controls forwarding of physical and directed IP broadcasts.
    ForwardProtocol Ip_ForwardProtocol
}

func (ip *Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "Cisco-IOS-XR-ip-tcp-cfg"
    ip.EntityData.SegmentPath = "Cisco-IOS-XR-ip-tcp-cfg:ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Children["cinetd"] = types.YChild{"Cinetd", &ip.Cinetd}
    ip.EntityData.Children["Cisco-IOS-XR-ip-udp-cfg:forward-protocol"] = types.YChild{"ForwardProtocol", &ip.ForwardProtocol}
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ip.EntityData)
}

// Ip_Cinetd
// Cinetd configuration data
type Ip_Cinetd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of service requests accepted per second. The type is interface{}
    // with range: 1..100. The default value is 1.
    RateLimit interface{}

    // Describing services of cinetd.
    Services Ip_Cinetd_Services
}

func (cinetd *Ip_Cinetd) GetEntityData() *types.CommonEntityData {
    cinetd.EntityData.YFilter = cinetd.YFilter
    cinetd.EntityData.YangName = "cinetd"
    cinetd.EntityData.BundleName = "cisco_ios_xr"
    cinetd.EntityData.ParentYangName = "ip"
    cinetd.EntityData.SegmentPath = "cinetd"
    cinetd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cinetd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cinetd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cinetd.EntityData.Children = make(map[string]types.YChild)
    cinetd.EntityData.Children["services"] = types.YChild{"Services", &cinetd.Services}
    cinetd.EntityData.Leafs = make(map[string]types.YLeaf)
    cinetd.EntityData.Leafs["rate-limit"] = types.YLeaf{"RateLimit", cinetd.RateLimit}
    return &(cinetd.EntityData)
}

// Ip_Cinetd_Services
// Describing services of cinetd
type Ip_Cinetd_Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV4 related services.
    Ipv4 Ip_Cinetd_Services_Ipv4

    // VRF table.
    Vrfs Ip_Cinetd_Services_Vrfs

    // IPV6 related services.
    Ipv6 Ip_Cinetd_Services_Ipv6
}

func (services *Ip_Cinetd_Services) GetEntityData() *types.CommonEntityData {
    services.EntityData.YFilter = services.YFilter
    services.EntityData.YangName = "services"
    services.EntityData.BundleName = "cisco_ios_xr"
    services.EntityData.ParentYangName = "cinetd"
    services.EntityData.SegmentPath = "services"
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = make(map[string]types.YChild)
    services.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &services.Ipv4}
    services.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &services.Vrfs}
    services.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &services.Ipv6}
    services.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(services.EntityData)
}

// Ip_Cinetd_Services_Ipv4
// IPV4 related services
type Ip_Cinetd_Services_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describing IPV4 and IPV6 small servers.
    SmallServers Ip_Cinetd_Services_Ipv4_SmallServers
}

func (ipv4 *Ip_Cinetd_Services_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "services"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["small-servers"] = types.YChild{"SmallServers", &ipv4.SmallServers}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// Ip_Cinetd_Services_Ipv4_SmallServers
// Describing IPV4 and IPV6 small servers
type Ip_Cinetd_Services_Ipv4_SmallServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describing TCP related IPV4 and IPV6 small servers.
    TcpSmallServers Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers

    // UDP small servers configuration.
    UdpSmallServers Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers
}

func (smallServers *Ip_Cinetd_Services_Ipv4_SmallServers) GetEntityData() *types.CommonEntityData {
    smallServers.EntityData.YFilter = smallServers.YFilter
    smallServers.EntityData.YangName = "small-servers"
    smallServers.EntityData.BundleName = "cisco_ios_xr"
    smallServers.EntityData.ParentYangName = "ipv4"
    smallServers.EntityData.SegmentPath = "small-servers"
    smallServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    smallServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    smallServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    smallServers.EntityData.Children = make(map[string]types.YChild)
    smallServers.EntityData.Children["tcp-small-servers"] = types.YChild{"TcpSmallServers", &smallServers.TcpSmallServers}
    smallServers.EntityData.Children["Cisco-IOS-XR-ip-udp-cfg:udp-small-servers"] = types.YChild{"UdpSmallServers", &smallServers.UdpSmallServers}
    smallServers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(smallServers.EntityData)
}

// Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers
// Describing TCP related IPV4 and IPV6 small
// servers
// This type is a presence type.
type Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the access list. The type is string.
    AccessControlListName interface{}

    // Set number of allowable TCP small servers, specify 0 for no-limit. The type
    // is one of the following types: enumeration
    // Ip.Cinetd.Services.Ipv6.SmallServers.TcpSmallServers.SmallServer This
    // attribute is mandatory., or int with range: 0..2147483647 This attribute is
    // mandatory..
    SmallServer interface{}
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers) GetEntityData() *types.CommonEntityData {
    tcpSmallServers.EntityData.YFilter = tcpSmallServers.YFilter
    tcpSmallServers.EntityData.YangName = "tcp-small-servers"
    tcpSmallServers.EntityData.BundleName = "cisco_ios_xr"
    tcpSmallServers.EntityData.ParentYangName = "small-servers"
    tcpSmallServers.EntityData.SegmentPath = "tcp-small-servers"
    tcpSmallServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcpSmallServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcpSmallServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcpSmallServers.EntityData.Children = make(map[string]types.YChild)
    tcpSmallServers.EntityData.Leafs = make(map[string]types.YLeaf)
    tcpSmallServers.EntityData.Leafs["access-control-list-name"] = types.YLeaf{"AccessControlListName", tcpSmallServers.AccessControlListName}
    tcpSmallServers.EntityData.Leafs["small-server"] = types.YLeaf{"SmallServer", tcpSmallServers.SmallServer}
    return &(tcpSmallServers.EntityData)
}

// Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers_SmallServer represents specify 0 for no-limit
type Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers_SmallServer string

const (
    // Unlimited Servers
    Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers_SmallServer_no_limit Ip_Cinetd_Services_Ipv4_SmallServers_TcpSmallServers_SmallServer = "no-limit"
)

// Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers
// UDP small servers configuration
// This type is a presence type.
type Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the access list. The type is string.
    AccessControlListName interface{}

    // Set number of allowable small servers, specify 0 for no-limit. The type is
    // one of the following types: enumeration
    // Ip.Cinetd.Services.Ipv4.SmallServers.UdpSmallServers.SmallServer This
    // attribute is mandatory., or int with range: 0..2147483647 This attribute is
    // mandatory..
    SmallServer interface{}
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers) GetEntityData() *types.CommonEntityData {
    udpSmallServers.EntityData.YFilter = udpSmallServers.YFilter
    udpSmallServers.EntityData.YangName = "udp-small-servers"
    udpSmallServers.EntityData.BundleName = "cisco_ios_xr"
    udpSmallServers.EntityData.ParentYangName = "small-servers"
    udpSmallServers.EntityData.SegmentPath = "Cisco-IOS-XR-ip-udp-cfg:udp-small-servers"
    udpSmallServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpSmallServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpSmallServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpSmallServers.EntityData.Children = make(map[string]types.YChild)
    udpSmallServers.EntityData.Leafs = make(map[string]types.YLeaf)
    udpSmallServers.EntityData.Leafs["access-control-list-name"] = types.YLeaf{"AccessControlListName", udpSmallServers.AccessControlListName}
    udpSmallServers.EntityData.Leafs["small-server"] = types.YLeaf{"SmallServer", udpSmallServers.SmallServer}
    return &(udpSmallServers.EntityData)
}

// Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers_SmallServer represents 0 for no-limit
type Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers_SmallServer string

const (
    // Unlimited Servers
    Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers_SmallServer_no_limit Ip_Cinetd_Services_Ipv4_SmallServers_UdpSmallServers_SmallServer = "no-limit"
)

// Ip_Cinetd_Services_Vrfs
// VRF table
type Ip_Cinetd_Services_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of Ip_Cinetd_Services_Vrfs_Vrf.
    Vrf []Ip_Cinetd_Services_Vrfs_Vrf
}

func (vrfs *Ip_Cinetd_Services_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "services"
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

// Ip_Cinetd_Services_Vrfs_Vrf
// VRF specific data
type Ip_Cinetd_Services_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IPV6 related services.
    Ipv6 Ip_Cinetd_Services_Vrfs_Vrf_Ipv6

    // IPV4 related services.
    Ipv4 Ip_Cinetd_Services_Vrfs_Vrf_Ipv4
}

func (vrf *Ip_Cinetd_Services_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &vrf.Ipv6}
    vrf.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &vrf.Ipv4}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6
// IPV6 related services
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TELNET server configuration commands.
    Telnet Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet

    // TFTP server configuration commands.
    Tftp Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp
}

func (ipv6 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "vrf"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["telnet"] = types.YChild{"Telnet", &ipv6.Telnet}
    ipv6.EntityData.Children["tftp"] = types.YChild{"Tftp", &ipv6.Tftp}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet
// TELNET server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP details.
    Tcp Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet) GetEntityData() *types.CommonEntityData {
    telnet.EntityData.YFilter = telnet.YFilter
    telnet.EntityData.YangName = "telnet"
    telnet.EntityData.BundleName = "cisco_ios_xr"
    telnet.EntityData.ParentYangName = "ipv6"
    telnet.EntityData.SegmentPath = "telnet"
    telnet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnet.EntityData.Children = make(map[string]types.YChild)
    telnet.EntityData.Children["tcp"] = types.YChild{"Tcp", &telnet.Tcp}
    telnet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(telnet.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp
// TCP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access list. The type is string.
    AccessListName interface{}

    // Set number of allowable servers. The type is interface{} with range:
    // 1..100. This attribute is mandatory.
    MaximumServer interface{}
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Telnet_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xr"
    tcp.EntityData.ParentYangName = "telnet"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcp.EntityData.Children = make(map[string]types.YChild)
    tcp.EntityData.Leafs = make(map[string]types.YLeaf)
    tcp.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", tcp.AccessListName}
    tcp.EntityData.Leafs["maximum-server"] = types.YLeaf{"MaximumServer", tcp.MaximumServer}
    return &(tcp.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp
// TFTP server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDP details.
    Udp Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp) GetEntityData() *types.CommonEntityData {
    tftp.EntityData.YFilter = tftp.YFilter
    tftp.EntityData.YangName = "tftp"
    tftp.EntityData.BundleName = "cisco_ios_xr"
    tftp.EntityData.ParentYangName = "ipv6"
    tftp.EntityData.SegmentPath = "tftp"
    tftp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftp.EntityData.Children = make(map[string]types.YChild)
    tftp.EntityData.Children["udp"] = types.YChild{"Udp", &tftp.Udp}
    tftp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tftp.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp
// UDP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp struct {
    EntityData types.CommonEntityData
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

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv6_Tftp_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xr"
    udp.EntityData.ParentYangName = "tftp"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udp.EntityData.Children = make(map[string]types.YChild)
    udp.EntityData.Leafs = make(map[string]types.YLeaf)
    udp.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", udp.AccessListName}
    udp.EntityData.Leafs["maximum-server"] = types.YLeaf{"MaximumServer", udp.MaximumServer}
    udp.EntityData.Leafs["home-directory"] = types.YLeaf{"HomeDirectory", udp.HomeDirectory}
    udp.EntityData.Leafs["dscp-value"] = types.YLeaf{"DscpValue", udp.DscpValue}
    return &(udp.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4
// IPV4 related services
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TELNET server configuration commands.
    Telnet Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet

    // TFTP server configuration commands.
    Tftp Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp
}

func (ipv4 *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "vrf"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["telnet"] = types.YChild{"Telnet", &ipv4.Telnet}
    ipv4.EntityData.Children["tftp"] = types.YChild{"Tftp", &ipv4.Tftp}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet
// TELNET server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP details.
    Tcp Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp
}

func (telnet *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet) GetEntityData() *types.CommonEntityData {
    telnet.EntityData.YFilter = telnet.YFilter
    telnet.EntityData.YangName = "telnet"
    telnet.EntityData.BundleName = "cisco_ios_xr"
    telnet.EntityData.ParentYangName = "ipv4"
    telnet.EntityData.SegmentPath = "telnet"
    telnet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnet.EntityData.Children = make(map[string]types.YChild)
    telnet.EntityData.Children["tcp"] = types.YChild{"Tcp", &telnet.Tcp}
    telnet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(telnet.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp
// TCP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access list. The type is string.
    AccessListName interface{}

    // Set number of allowable servers. The type is interface{} with range:
    // 1..100. This attribute is mandatory.
    MaximumServer interface{}
}

func (tcp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Telnet_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xr"
    tcp.EntityData.ParentYangName = "telnet"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcp.EntityData.Children = make(map[string]types.YChild)
    tcp.EntityData.Leafs = make(map[string]types.YLeaf)
    tcp.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", tcp.AccessListName}
    tcp.EntityData.Leafs["maximum-server"] = types.YLeaf{"MaximumServer", tcp.MaximumServer}
    return &(tcp.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp
// TFTP server configuration commands
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDP details.
    Udp Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp
}

func (tftp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp) GetEntityData() *types.CommonEntityData {
    tftp.EntityData.YFilter = tftp.YFilter
    tftp.EntityData.YangName = "tftp"
    tftp.EntityData.BundleName = "cisco_ios_xr"
    tftp.EntityData.ParentYangName = "ipv4"
    tftp.EntityData.SegmentPath = "tftp"
    tftp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftp.EntityData.Children = make(map[string]types.YChild)
    tftp.EntityData.Children["udp"] = types.YChild{"Udp", &tftp.Udp}
    tftp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tftp.EntityData)
}

// Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp
// UDP details
// This type is a presence type.
type Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp struct {
    EntityData types.CommonEntityData
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

func (udp *Ip_Cinetd_Services_Vrfs_Vrf_Ipv4_Tftp_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xr"
    udp.EntityData.ParentYangName = "tftp"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udp.EntityData.Children = make(map[string]types.YChild)
    udp.EntityData.Leafs = make(map[string]types.YLeaf)
    udp.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", udp.AccessListName}
    udp.EntityData.Leafs["maximum-server"] = types.YLeaf{"MaximumServer", udp.MaximumServer}
    udp.EntityData.Leafs["home-directory"] = types.YLeaf{"HomeDirectory", udp.HomeDirectory}
    udp.EntityData.Leafs["dscp-value"] = types.YLeaf{"DscpValue", udp.DscpValue}
    return &(udp.EntityData)
}

// Ip_Cinetd_Services_Ipv6
// IPV6 related services
type Ip_Cinetd_Services_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describing IPV4 and IPV6 small servers.
    SmallServers Ip_Cinetd_Services_Ipv6_SmallServers
}

func (ipv6 *Ip_Cinetd_Services_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "services"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["small-servers"] = types.YChild{"SmallServers", &ipv6.SmallServers}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// Ip_Cinetd_Services_Ipv6_SmallServers
// Describing IPV4 and IPV6 small servers
type Ip_Cinetd_Services_Ipv6_SmallServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describing TCP related IPV4 and IPV6 small servers.
    TcpSmallServers Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers

    // UDP small servers configuration.
    UdpSmallServers Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers
}

func (smallServers *Ip_Cinetd_Services_Ipv6_SmallServers) GetEntityData() *types.CommonEntityData {
    smallServers.EntityData.YFilter = smallServers.YFilter
    smallServers.EntityData.YangName = "small-servers"
    smallServers.EntityData.BundleName = "cisco_ios_xr"
    smallServers.EntityData.ParentYangName = "ipv6"
    smallServers.EntityData.SegmentPath = "small-servers"
    smallServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    smallServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    smallServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    smallServers.EntityData.Children = make(map[string]types.YChild)
    smallServers.EntityData.Children["tcp-small-servers"] = types.YChild{"TcpSmallServers", &smallServers.TcpSmallServers}
    smallServers.EntityData.Children["Cisco-IOS-XR-ip-udp-cfg:udp-small-servers"] = types.YChild{"UdpSmallServers", &smallServers.UdpSmallServers}
    smallServers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(smallServers.EntityData)
}

// Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers
// Describing TCP related IPV4 and IPV6 small
// servers
// This type is a presence type.
type Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the access list. The type is string.
    AccessControlListName interface{}

    // Set number of allowable TCP small servers, specify 0 for no-limit. The type
    // is one of the following types: enumeration
    // Ip.Cinetd.Services.Ipv6.SmallServers.TcpSmallServers.SmallServer This
    // attribute is mandatory., or int with range: 0..2147483647 This attribute is
    // mandatory..
    SmallServer interface{}
}

func (tcpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers) GetEntityData() *types.CommonEntityData {
    tcpSmallServers.EntityData.YFilter = tcpSmallServers.YFilter
    tcpSmallServers.EntityData.YangName = "tcp-small-servers"
    tcpSmallServers.EntityData.BundleName = "cisco_ios_xr"
    tcpSmallServers.EntityData.ParentYangName = "small-servers"
    tcpSmallServers.EntityData.SegmentPath = "tcp-small-servers"
    tcpSmallServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcpSmallServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcpSmallServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcpSmallServers.EntityData.Children = make(map[string]types.YChild)
    tcpSmallServers.EntityData.Leafs = make(map[string]types.YLeaf)
    tcpSmallServers.EntityData.Leafs["access-control-list-name"] = types.YLeaf{"AccessControlListName", tcpSmallServers.AccessControlListName}
    tcpSmallServers.EntityData.Leafs["small-server"] = types.YLeaf{"SmallServer", tcpSmallServers.SmallServer}
    return &(tcpSmallServers.EntityData)
}

// Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers_SmallServer represents specify 0 for no-limit
type Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers_SmallServer string

const (
    // Unlimited Servers
    Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers_SmallServer_no_limit Ip_Cinetd_Services_Ipv6_SmallServers_TcpSmallServers_SmallServer = "no-limit"
)

// Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers
// UDP small servers configuration
// This type is a presence type.
type Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the access list. The type is string.
    AccessControlListName interface{}

    // Set number of allowable small servers, specify 0 for no-limit. The type is
    // one of the following types: enumeration
    // Ip.Cinetd.Services.Ipv6.SmallServers.UdpSmallServers.SmallServer This
    // attribute is mandatory., or int with range: 0..2147483647 This attribute is
    // mandatory..
    SmallServer interface{}
}

func (udpSmallServers *Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers) GetEntityData() *types.CommonEntityData {
    udpSmallServers.EntityData.YFilter = udpSmallServers.YFilter
    udpSmallServers.EntityData.YangName = "udp-small-servers"
    udpSmallServers.EntityData.BundleName = "cisco_ios_xr"
    udpSmallServers.EntityData.ParentYangName = "small-servers"
    udpSmallServers.EntityData.SegmentPath = "Cisco-IOS-XR-ip-udp-cfg:udp-small-servers"
    udpSmallServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpSmallServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpSmallServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpSmallServers.EntityData.Children = make(map[string]types.YChild)
    udpSmallServers.EntityData.Leafs = make(map[string]types.YLeaf)
    udpSmallServers.EntityData.Leafs["access-control-list-name"] = types.YLeaf{"AccessControlListName", udpSmallServers.AccessControlListName}
    udpSmallServers.EntityData.Leafs["small-server"] = types.YLeaf{"SmallServer", udpSmallServers.SmallServer}
    return &(udpSmallServers.EntityData)
}

// Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers_SmallServer represents 0 for no-limit
type Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers_SmallServer string

const (
    // Unlimited Servers
    Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers_SmallServer_no_limit Ip_Cinetd_Services_Ipv6_SmallServers_UdpSmallServers_SmallServer = "no-limit"
)

// Ip_ForwardProtocol
// Controls forwarding of physical and directed IP
// broadcasts
type Ip_ForwardProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets to a specific UDP port.
    Udp Ip_ForwardProtocol_Udp
}

func (forwardProtocol *Ip_ForwardProtocol) GetEntityData() *types.CommonEntityData {
    forwardProtocol.EntityData.YFilter = forwardProtocol.YFilter
    forwardProtocol.EntityData.YangName = "forward-protocol"
    forwardProtocol.EntityData.BundleName = "cisco_ios_xr"
    forwardProtocol.EntityData.ParentYangName = "ip"
    forwardProtocol.EntityData.SegmentPath = "Cisco-IOS-XR-ip-udp-cfg:forward-protocol"
    forwardProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardProtocol.EntityData.Children = make(map[string]types.YChild)
    forwardProtocol.EntityData.Children["udp"] = types.YChild{"Udp", &forwardProtocol.Udp}
    forwardProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(forwardProtocol.EntityData)
}

// Ip_ForwardProtocol_Udp
// Packets to a specific UDP port
type Ip_ForwardProtocol_Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable IP Forward Protocol UDP. The type is interface{}.
    Disable interface{}

    // Port configuration.
    Ports Ip_ForwardProtocol_Udp_Ports
}

func (udp *Ip_ForwardProtocol_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xr"
    udp.EntityData.ParentYangName = "forward-protocol"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udp.EntityData.Children = make(map[string]types.YChild)
    udp.EntityData.Children["ports"] = types.YChild{"Ports", &udp.Ports}
    udp.EntityData.Leafs = make(map[string]types.YLeaf)
    udp.EntityData.Leafs["disable"] = types.YLeaf{"Disable", udp.Disable}
    return &(udp.EntityData)
}

// Ip_ForwardProtocol_Udp_Ports
// Port configuration
type Ip_ForwardProtocol_Udp_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Well-known ports are enabled by default and non well-known ports are
    // disabled by default. It is not allowed to configure the default. The type
    // is slice of Ip_ForwardProtocol_Udp_Ports_Port.
    Port []Ip_ForwardProtocol_Udp_Ports_Port
}

func (ports *Ip_ForwardProtocol_Udp_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "udp"
    ports.EntityData.SegmentPath = "ports"
    ports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ports.EntityData.Children = make(map[string]types.YChild)
    ports.EntityData.Children["port"] = types.YChild{"Port", nil}
    for i := range ports.Port {
        ports.EntityData.Children[types.GetSegmentPath(&ports.Port[i])] = types.YChild{"Port", &ports.Port[i]}
    }
    ports.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ports.EntityData)
}

// Ip_ForwardProtocol_Udp_Ports_Port
// Well-known ports are enabled by default and
// non well-known ports are disabled by default.
// It is not allowed to configure the default.
type Ip_ForwardProtocol_Udp_Ports_Port struct {
    EntityData types.CommonEntityData
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

func (port *Ip_ForwardProtocol_Udp_Ports_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "ports"
    port.EntityData.SegmentPath = "port" + "[port-id='" + fmt.Sprintf("%v", port.PortId) + "']"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    port.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", port.PortId}
    port.EntityData.Leafs["enable"] = types.YLeaf{"Enable", port.Enable}
    return &(port.EntityData)
}

