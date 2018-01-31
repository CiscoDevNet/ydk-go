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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // UDP receive Queue Size. The type is interface{} with range: 40..800.
    ReceiveQ interface{}

    // UDP InQueue and OutQueue threads.
    NumThread IpUdp_NumThread

    // UDP directory details.
    Directory IpUdp_Directory
}

func (ipUdp *IpUdp) GetFilter() yfilter.YFilter { return ipUdp.YFilter }

func (ipUdp *IpUdp) SetFilter(yf yfilter.YFilter) { ipUdp.YFilter = yf }

func (ipUdp *IpUdp) GetGoName(yname string) string {
    if yname == "receive-q" { return "ReceiveQ" }
    if yname == "num-thread" { return "NumThread" }
    if yname == "directory" { return "Directory" }
    return ""
}

func (ipUdp *IpUdp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-udp-cfg:ip-udp"
}

func (ipUdp *IpUdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "num-thread" {
        return &ipUdp.NumThread
    }
    if childYangName == "directory" {
        return &ipUdp.Directory
    }
    return nil
}

func (ipUdp *IpUdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["num-thread"] = &ipUdp.NumThread
    children["directory"] = &ipUdp.Directory
    return children
}

func (ipUdp *IpUdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receive-q"] = ipUdp.ReceiveQ
    return leafs
}

func (ipUdp *IpUdp) GetBundleName() string { return "cisco_ios_xr" }

func (ipUdp *IpUdp) GetYangName() string { return "ip-udp" }

func (ipUdp *IpUdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipUdp *IpUdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipUdp *IpUdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipUdp *IpUdp) SetParent(parent types.Entity) { ipUdp.parent = parent }

func (ipUdp *IpUdp) GetParent() types.Entity { return ipUdp.parent }

func (ipUdp *IpUdp) GetParentYangName() string { return "Cisco-IOS-XR-ip-udp-cfg" }

// IpUdp_NumThread
// UDP InQueue and OutQueue threads
// This type is a presence type.
type IpUdp_NumThread struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // InQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    UdpInQThreads interface{}

    // OutQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    UdpOutQThreads interface{}
}

func (numThread *IpUdp_NumThread) GetFilter() yfilter.YFilter { return numThread.YFilter }

func (numThread *IpUdp_NumThread) SetFilter(yf yfilter.YFilter) { numThread.YFilter = yf }

func (numThread *IpUdp_NumThread) GetGoName(yname string) string {
    if yname == "udp-in-q-threads" { return "UdpInQThreads" }
    if yname == "udp-out-q-threads" { return "UdpOutQThreads" }
    return ""
}

func (numThread *IpUdp_NumThread) GetSegmentPath() string {
    return "num-thread"
}

func (numThread *IpUdp_NumThread) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numThread *IpUdp_NumThread) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numThread *IpUdp_NumThread) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udp-in-q-threads"] = numThread.UdpInQThreads
    leafs["udp-out-q-threads"] = numThread.UdpOutQThreads
    return leafs
}

func (numThread *IpUdp_NumThread) GetBundleName() string { return "cisco_ios_xr" }

func (numThread *IpUdp_NumThread) GetYangName() string { return "num-thread" }

func (numThread *IpUdp_NumThread) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numThread *IpUdp_NumThread) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numThread *IpUdp_NumThread) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numThread *IpUdp_NumThread) SetParent(parent types.Entity) { numThread.parent = parent }

func (numThread *IpUdp_NumThread) GetParent() types.Entity { return numThread.parent }

func (numThread *IpUdp_NumThread) GetParentYangName() string { return "ip-udp" }

// IpUdp_Directory
// UDP directory details
// This type is a presence type.
type IpUdp_Directory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Directory name. The type is string. This attribute is mandatory.
    Directoryname interface{}

    // Set number of Debug files. The type is interface{} with range: 1..5000.
    // This attribute is mandatory.
    MaxUdpDebugFiles interface{}

    // Set size of debug files in bytes. The type is interface{} with range:
    // 1024..4294967295. This attribute is mandatory. Units are byte.
    MaxFileSizeFiles interface{}
}

func (directory *IpUdp_Directory) GetFilter() yfilter.YFilter { return directory.YFilter }

func (directory *IpUdp_Directory) SetFilter(yf yfilter.YFilter) { directory.YFilter = yf }

func (directory *IpUdp_Directory) GetGoName(yname string) string {
    if yname == "directoryname" { return "Directoryname" }
    if yname == "max-udp-debug-files" { return "MaxUdpDebugFiles" }
    if yname == "max-file-size-files" { return "MaxFileSizeFiles" }
    return ""
}

func (directory *IpUdp_Directory) GetSegmentPath() string {
    return "directory"
}

func (directory *IpUdp_Directory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (directory *IpUdp_Directory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (directory *IpUdp_Directory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["directoryname"] = directory.Directoryname
    leafs["max-udp-debug-files"] = directory.MaxUdpDebugFiles
    leafs["max-file-size-files"] = directory.MaxFileSizeFiles
    return leafs
}

func (directory *IpUdp_Directory) GetBundleName() string { return "cisco_ios_xr" }

func (directory *IpUdp_Directory) GetYangName() string { return "directory" }

func (directory *IpUdp_Directory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (directory *IpUdp_Directory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (directory *IpUdp_Directory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (directory *IpUdp_Directory) SetParent(parent types.Entity) { directory.parent = parent }

func (directory *IpUdp_Directory) GetParent() types.Entity { return directory.parent }

func (directory *IpUdp_Directory) GetParentYangName() string { return "ip-udp" }

