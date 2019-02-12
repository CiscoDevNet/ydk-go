// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// This module contains definitions
// for the following management objects:
// debug_trace: Calvados debug trace.
// 
// Copyright (c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_debug_trace

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_debug_trace"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-debug-trace config}", reflect.TypeOf(Config{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-debug-trace:config", reflect.TypeOf(Config{}))
}

// Config
type Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Debug Config_Debug
}

func (config *Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-debug-trace"
    config.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-debug-trace:config"
    config.EntityData.AbsolutePath = config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("debug", types.YChild{"Debug", &config.Debug})
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Config_Debug
type Config_Debug struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Debug_Trace.
    Trace []*Config_Debug_Trace
}

func (debug *Config_Debug) GetEntityData() *types.CommonEntityData {
    debug.EntityData.YFilter = debug.YFilter
    debug.EntityData.YangName = "debug"
    debug.EntityData.BundleName = "cisco_ios_xr"
    debug.EntityData.ParentYangName = "config"
    debug.EntityData.SegmentPath = "debug"
    debug.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-debug-trace:config/" + debug.EntityData.SegmentPath
    debug.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    debug.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    debug.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    debug.EntityData.Children = types.NewOrderedMap()
    debug.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range debug.Trace {
        debug.EntityData.Children.Append(types.GetSegmentPath(debug.Trace[i]), types.YChild{"Trace", debug.Trace[i]})
    }
    debug.EntityData.Leafs = types.NewOrderedMap()

    debug.EntityData.YListKeys = []string {}

    return &(debug.EntityData)
}

// Config_Debug_Trace
type Config_Debug_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    ConnectionType interface{}

    // The type is interface{}.
    Enable interface{}

    // The type is interface{}.
    Disable interface{}
}

func (trace *Config_Debug_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "debug"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.ConnectionType, "connection_type")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-debug-trace:config/debug/" + trace.EntityData.SegmentPath
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("connection_type", types.YLeaf{"ConnectionType", trace.ConnectionType})
    trace.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", trace.Enable})
    trace.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", trace.Disable})

    trace.EntityData.YListKeys = []string {"ConnectionType"}

    return &(trace.EntityData)
}

