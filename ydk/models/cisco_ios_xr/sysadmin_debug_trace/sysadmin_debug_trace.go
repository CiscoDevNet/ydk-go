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
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["debug"] = types.YChild{"Debug", &config.Debug}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(config.EntityData)
}

// Config_Debug
type Config_Debug struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Debug_Trace.
    Trace []Config_Debug_Trace
}

func (debug *Config_Debug) GetEntityData() *types.CommonEntityData {
    debug.EntityData.YFilter = debug.YFilter
    debug.EntityData.YangName = "debug"
    debug.EntityData.BundleName = "cisco_ios_xr"
    debug.EntityData.ParentYangName = "config"
    debug.EntityData.SegmentPath = "debug"
    debug.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    debug.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    debug.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    debug.EntityData.Children = make(map[string]types.YChild)
    debug.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range debug.Trace {
        debug.EntityData.Children[types.GetSegmentPath(&debug.Trace[i])] = types.YChild{"Trace", &debug.Trace[i]}
    }
    debug.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(debug.EntityData)
}

// Config_Debug_Trace
type Config_Debug_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    trace.EntityData.SegmentPath = "trace" + "[connection_type='" + fmt.Sprintf("%v", trace.ConnectionType) + "']"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    trace.EntityData.Leafs["connection_type"] = types.YLeaf{"ConnectionType", trace.ConnectionType}
    trace.EntityData.Leafs["enable"] = types.YLeaf{"Enable", trace.Enable}
    trace.EntityData.Leafs["disable"] = types.YLeaf{"Disable", trace.Disable}
    return &(trace.EntityData)
}

