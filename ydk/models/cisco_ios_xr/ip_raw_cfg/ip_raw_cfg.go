// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-raw package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ip-raw: Global IP RAW configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_raw_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_raw_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-raw-cfg ip-raw}", reflect.TypeOf(IpRaw{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-raw-cfg:ip-raw", reflect.TypeOf(IpRaw{}))
}

// IpRaw
// Global IP RAW configuration
type IpRaw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RAW receive Queue Size. The type is interface{} with range: 40..800.
    ReceiveQ interface{}

    // RAW InQueue and OutQueue threads.
    NumThread IpRaw_NumThread

    // RAW directory details.
    Directory IpRaw_Directory
}

func (ipRaw *IpRaw) GetEntityData() *types.CommonEntityData {
    ipRaw.EntityData.YFilter = ipRaw.YFilter
    ipRaw.EntityData.YangName = "ip-raw"
    ipRaw.EntityData.BundleName = "cisco_ios_xr"
    ipRaw.EntityData.ParentYangName = "Cisco-IOS-XR-ip-raw-cfg"
    ipRaw.EntityData.SegmentPath = "Cisco-IOS-XR-ip-raw-cfg:ip-raw"
    ipRaw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipRaw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipRaw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipRaw.EntityData.Children = make(map[string]types.YChild)
    ipRaw.EntityData.Children["num-thread"] = types.YChild{"NumThread", &ipRaw.NumThread}
    ipRaw.EntityData.Children["directory"] = types.YChild{"Directory", &ipRaw.Directory}
    ipRaw.EntityData.Leafs = make(map[string]types.YLeaf)
    ipRaw.EntityData.Leafs["receive-q"] = types.YLeaf{"ReceiveQ", ipRaw.ReceiveQ}
    return &(ipRaw.EntityData)
}

// IpRaw_NumThread
// RAW InQueue and OutQueue threads
// This type is a presence type.
type IpRaw_NumThread struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // InQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    RawInQThreads interface{}

    // OutQ Threads. The type is interface{} with range: 1..16. This attribute is
    // mandatory.
    RawOutQThreads interface{}
}

func (numThread *IpRaw_NumThread) GetEntityData() *types.CommonEntityData {
    numThread.EntityData.YFilter = numThread.YFilter
    numThread.EntityData.YangName = "num-thread"
    numThread.EntityData.BundleName = "cisco_ios_xr"
    numThread.EntityData.ParentYangName = "ip-raw"
    numThread.EntityData.SegmentPath = "num-thread"
    numThread.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numThread.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numThread.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numThread.EntityData.Children = make(map[string]types.YChild)
    numThread.EntityData.Leafs = make(map[string]types.YLeaf)
    numThread.EntityData.Leafs["raw-in-q-threads"] = types.YLeaf{"RawInQThreads", numThread.RawInQThreads}
    numThread.EntityData.Leafs["raw-out-q-threads"] = types.YLeaf{"RawOutQThreads", numThread.RawOutQThreads}
    return &(numThread.EntityData)
}

// IpRaw_Directory
// RAW directory details
// This type is a presence type.
type IpRaw_Directory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Directory name. The type is string. This attribute is mandatory.
    Directoryname interface{}

    // Set number of Debug files. The type is interface{} with range: 1..18000.
    // The default value is 256.
    MaxRawDebugFiles interface{}

    // Set size of debug files in bytes. The type is interface{} with range:
    // 1024..4294967295. Units are byte.
    MaxFileSizeFiles interface{}
}

func (directory *IpRaw_Directory) GetEntityData() *types.CommonEntityData {
    directory.EntityData.YFilter = directory.YFilter
    directory.EntityData.YangName = "directory"
    directory.EntityData.BundleName = "cisco_ios_xr"
    directory.EntityData.ParentYangName = "ip-raw"
    directory.EntityData.SegmentPath = "directory"
    directory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directory.EntityData.Children = make(map[string]types.YChild)
    directory.EntityData.Leafs = make(map[string]types.YLeaf)
    directory.EntityData.Leafs["directoryname"] = types.YLeaf{"Directoryname", directory.Directoryname}
    directory.EntityData.Leafs["max-raw-debug-files"] = types.YLeaf{"MaxRawDebugFiles", directory.MaxRawDebugFiles}
    directory.EntityData.Leafs["max-file-size-files"] = types.YLeaf{"MaxFileSizeFiles", directory.MaxFileSizeFiles}
    return &(directory.EntityData)
}

