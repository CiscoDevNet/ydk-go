// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-udp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ip-udp: Global IP UDP configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ip-tcp-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_udp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_udp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-udp-cfg ip-udp}", reflect.TypeOf(IpUdp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-udp-cfg:ip-udp", reflect.TypeOf(IpUdp{}))
}

// IpUdp
// Global IP UDP configuration
type IpUdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDP receive Queue Size. The type is interface{} with range: 40..800.
    ReceiveQ interface{}

    // UDP InQueue and OutQueue threads.
    NumThread IpUdp_NumThread

    // UDP directory details.
    Directory IpUdp_Directory
}

func (ipUdp *IpUdp) GetEntityData() *types.CommonEntityData {
    ipUdp.EntityData.YFilter = ipUdp.YFilter
    ipUdp.EntityData.YangName = "ip-udp"
    ipUdp.EntityData.BundleName = "cisco_ios_xr"
    ipUdp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-udp-cfg"
    ipUdp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-udp-cfg:ip-udp"
    ipUdp.EntityData.AbsolutePath = ipUdp.EntityData.SegmentPath
    ipUdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipUdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipUdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipUdp.EntityData.Children = types.NewOrderedMap()
    ipUdp.EntityData.Children.Append("num-thread", types.YChild{"NumThread", &ipUdp.NumThread})
    ipUdp.EntityData.Children.Append("directory", types.YChild{"Directory", &ipUdp.Directory})
    ipUdp.EntityData.Leafs = types.NewOrderedMap()
    ipUdp.EntityData.Leafs.Append("receive-q", types.YLeaf{"ReceiveQ", ipUdp.ReceiveQ})

    ipUdp.EntityData.YListKeys = []string {}

    return &(ipUdp.EntityData)
}

// IpUdp_NumThread
// UDP InQueue and OutQueue threads
// This type is a presence type.
type IpUdp_NumThread struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // InQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    UdpInQThreads interface{}

    // OutQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    UdpOutQThreads interface{}
}

func (numThread *IpUdp_NumThread) GetEntityData() *types.CommonEntityData {
    numThread.EntityData.YFilter = numThread.YFilter
    numThread.EntityData.YangName = "num-thread"
    numThread.EntityData.BundleName = "cisco_ios_xr"
    numThread.EntityData.ParentYangName = "ip-udp"
    numThread.EntityData.SegmentPath = "num-thread"
    numThread.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-udp-cfg:ip-udp/" + numThread.EntityData.SegmentPath
    numThread.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numThread.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numThread.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numThread.EntityData.Children = types.NewOrderedMap()
    numThread.EntityData.Leafs = types.NewOrderedMap()
    numThread.EntityData.Leafs.Append("udp-in-q-threads", types.YLeaf{"UdpInQThreads", numThread.UdpInQThreads})
    numThread.EntityData.Leafs.Append("udp-out-q-threads", types.YLeaf{"UdpOutQThreads", numThread.UdpOutQThreads})

    numThread.EntityData.YListKeys = []string {}

    return &(numThread.EntityData)
}

// IpUdp_Directory
// UDP directory details
// This type is a presence type.
type IpUdp_Directory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Directory name. The type is string. This attribute is mandatory.
    Directoryname interface{}

    // Set number of Debug files. The type is interface{} with range: 1..5000. The
    // default value is 256.
    MaxUdpDebugFiles interface{}

    // Set size of debug files in bytes. The type is interface{} with range:
    // 1024..4294967295. Units are byte.
    MaxFileSizeFiles interface{}
}

func (directory *IpUdp_Directory) GetEntityData() *types.CommonEntityData {
    directory.EntityData.YFilter = directory.YFilter
    directory.EntityData.YangName = "directory"
    directory.EntityData.BundleName = "cisco_ios_xr"
    directory.EntityData.ParentYangName = "ip-udp"
    directory.EntityData.SegmentPath = "directory"
    directory.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-udp-cfg:ip-udp/" + directory.EntityData.SegmentPath
    directory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directory.EntityData.Children = types.NewOrderedMap()
    directory.EntityData.Leafs = types.NewOrderedMap()
    directory.EntityData.Leafs.Append("directoryname", types.YLeaf{"Directoryname", directory.Directoryname})
    directory.EntityData.Leafs.Append("max-udp-debug-files", types.YLeaf{"MaxUdpDebugFiles", directory.MaxUdpDebugFiles})
    directory.EntityData.Leafs.Append("max-file-size-files", types.YLeaf{"MaxFileSizeFiles", directory.MaxFileSizeFiles})

    directory.EntityData.YListKeys = []string {}

    return &(directory.EntityData)
}

