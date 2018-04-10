// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all 'hw-module' commands for Sysadmin.
// 
// Copyright(c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_hw_module

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_hw_module"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-hw-module hw-module}", reflect.TypeOf(HwModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-hw-module:hw-module", reflect.TypeOf(HwModule{}))
}

// HwModule
type HwModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Config HwModule_Config

    
    Oper HwModule_Oper

    
    Shhwfpd HwModule_Shhwfpd
}

func (hwModule *HwModule) GetEntityData() *types.CommonEntityData {
    hwModule.EntityData.YFilter = hwModule.YFilter
    hwModule.EntityData.YangName = "hw-module"
    hwModule.EntityData.BundleName = "cisco_ios_xr"
    hwModule.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-hw-module"
    hwModule.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-hw-module:hw-module"
    hwModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModule.EntityData.Children = make(map[string]types.YChild)
    hwModule.EntityData.Children["config"] = types.YChild{"Config", &hwModule.Config}
    hwModule.EntityData.Children["oper"] = types.YChild{"Oper", &hwModule.Oper}
    hwModule.EntityData.Children["shhwfpd"] = types.YChild{"Shhwfpd", &hwModule.Shhwfpd}
    hwModule.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwModule.EntityData)
}

// HwModule_Config
type HwModule_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Shutdown HwModule_Config_Shutdown

    
    Reset HwModule_Config_Reset

    
    Offline HwModule_Config_Offline

    
    AttentionLed HwModule_Config_AttentionLed

    // The type is slice of HwModule_Config_Location.
    Location []HwModule_Config_Location
}

func (config *HwModule_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "hw-module"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["shutdown"] = types.YChild{"Shutdown", &config.Shutdown}
    config.EntityData.Children["reset"] = types.YChild{"Reset", &config.Reset}
    config.EntityData.Children["offline"] = types.YChild{"Offline", &config.Offline}
    config.EntityData.Children["attention-led"] = types.YChild{"AttentionLed", &config.AttentionLed}
    config.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range config.Location {
        config.EntityData.Children[types.GetSegmentPath(&config.Location[i])] = types.YChild{"Location", &config.Location[i]}
    }
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(config.EntityData)
}

// HwModule_Config_Shutdown
type HwModule_Config_Shutdown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of HwModule_Config_Shutdown_Location.
    Location []HwModule_Config_Shutdown_Location
}

func (shutdown *HwModule_Config_Shutdown) GetEntityData() *types.CommonEntityData {
    shutdown.EntityData.YFilter = shutdown.YFilter
    shutdown.EntityData.YangName = "shutdown"
    shutdown.EntityData.BundleName = "cisco_ios_xr"
    shutdown.EntityData.ParentYangName = "config"
    shutdown.EntityData.SegmentPath = "shutdown"
    shutdown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shutdown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shutdown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shutdown.EntityData.Children = make(map[string]types.YChild)
    shutdown.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range shutdown.Location {
        shutdown.EntityData.Children[types.GetSegmentPath(&shutdown.Location[i])] = types.YChild{"Location", &shutdown.Location[i]}
    }
    shutdown.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shutdown.EntityData)
}

// HwModule_Config_Shutdown_Location
type HwModule_Config_Shutdown_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    // This attribute is mandatory.
    Location interface{}
}

func (location *HwModule_Config_Shutdown_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "shutdown"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// HwModule_Config_Reset
type HwModule_Config_Reset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Auto HwModule_Config_Reset_Auto
}

func (reset *HwModule_Config_Reset) GetEntityData() *types.CommonEntityData {
    reset.EntityData.YFilter = reset.YFilter
    reset.EntityData.YangName = "reset"
    reset.EntityData.BundleName = "cisco_ios_xr"
    reset.EntityData.ParentYangName = "config"
    reset.EntityData.SegmentPath = "reset"
    reset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reset.EntityData.Children = make(map[string]types.YChild)
    reset.EntityData.Children["auto"] = types.YChild{"Auto", &reset.Auto}
    reset.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reset.EntityData)
}

// HwModule_Config_Reset_Auto
type HwModule_Config_Reset_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Disable HwModule_Config_Reset_Auto_Disable
}

func (auto *HwModule_Config_Reset_Auto) GetEntityData() *types.CommonEntityData {
    auto.EntityData.YFilter = auto.YFilter
    auto.EntityData.YangName = "auto"
    auto.EntityData.BundleName = "cisco_ios_xr"
    auto.EntityData.ParentYangName = "reset"
    auto.EntityData.SegmentPath = "auto"
    auto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auto.EntityData.Children = make(map[string]types.YChild)
    auto.EntityData.Children["disable"] = types.YChild{"Disable", &auto.Disable}
    auto.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(auto.EntityData)
}

// HwModule_Config_Reset_Auto_Disable
type HwModule_Config_Reset_Auto_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of HwModule_Config_Reset_Auto_Disable_Location.
    Location []HwModule_Config_Reset_Auto_Disable_Location
}

func (disable *HwModule_Config_Reset_Auto_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "auto"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = make(map[string]types.YChild)
    disable.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range disable.Location {
        disable.EntityData.Children[types.GetSegmentPath(&disable.Location[i])] = types.YChild{"Location", &disable.Location[i]}
    }
    disable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(disable.EntityData)
}

// HwModule_Config_Reset_Auto_Disable_Location
type HwModule_Config_Reset_Auto_Disable_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    // This attribute is mandatory.
    Location interface{}
}

func (location *HwModule_Config_Reset_Auto_Disable_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "disable"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// HwModule_Config_Offline
type HwModule_Config_Offline struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of HwModule_Config_Offline_Location.
    Location []HwModule_Config_Offline_Location
}

func (offline *HwModule_Config_Offline) GetEntityData() *types.CommonEntityData {
    offline.EntityData.YFilter = offline.YFilter
    offline.EntityData.YangName = "offline"
    offline.EntityData.BundleName = "cisco_ios_xr"
    offline.EntityData.ParentYangName = "config"
    offline.EntityData.SegmentPath = "offline"
    offline.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    offline.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    offline.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    offline.EntityData.Children = make(map[string]types.YChild)
    offline.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range offline.Location {
        offline.EntityData.Children[types.GetSegmentPath(&offline.Location[i])] = types.YChild{"Location", &offline.Location[i]}
    }
    offline.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(offline.EntityData)
}

// HwModule_Config_Offline_Location
type HwModule_Config_Offline_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    // This attribute is mandatory.
    Location interface{}
}

func (location *HwModule_Config_Offline_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "offline"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// HwModule_Config_AttentionLed
type HwModule_Config_AttentionLed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of HwModule_Config_AttentionLed_Location.
    Location []HwModule_Config_AttentionLed_Location
}

func (attentionLed *HwModule_Config_AttentionLed) GetEntityData() *types.CommonEntityData {
    attentionLed.EntityData.YFilter = attentionLed.YFilter
    attentionLed.EntityData.YangName = "attention-led"
    attentionLed.EntityData.BundleName = "cisco_ios_xr"
    attentionLed.EntityData.ParentYangName = "config"
    attentionLed.EntityData.SegmentPath = "attention-led"
    attentionLed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attentionLed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attentionLed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attentionLed.EntityData.Children = make(map[string]types.YChild)
    attentionLed.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range attentionLed.Location {
        attentionLed.EntityData.Children[types.GetSegmentPath(&attentionLed.Location[i])] = types.YChild{"Location", &attentionLed.Location[i]}
    }
    attentionLed.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attentionLed.EntityData)
}

// HwModule_Config_AttentionLed_Location
type HwModule_Config_AttentionLed_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'([0-255])|((0?[0-9]|1[0-5])/(FC(0?[0-5])))|((0?[0-9]|1[0-5])/(RP(0?[0-1])))|((0?[0-9]|1[0-5])/(0?[0-9]|1[0-5]))|((F[0-3])/(FC(0?[0-9]|1[0-1])))|((F[0-3])/(SC(0?[0-1])))|(0?[0-9]/(SC(0?[0-1])))|(0?[0-9]/(FT(0?[0-2])))|(0?[0-9]/(PM(0?[0-9])))'.
    // This attribute is mandatory.
    Location interface{}
}

func (location *HwModule_Config_AttentionLed_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "attention-led"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// HwModule_Config_Location
type HwModule_Config_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    Logging HwModule_Config_Location_Logging
}

func (location *HwModule_Config_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "config"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["logging"] = types.YChild{"Logging", &location.Logging}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// HwModule_Config_Location_Logging
type HwModule_Config_Location_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Onboard HwModule_Config_Location_Logging_Onboard
}

func (logging *HwModule_Config_Location_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "location"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Children["onboard"] = types.YChild{"Onboard", &logging.Onboard}
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(logging.EntityData)
}

// HwModule_Config_Location_Logging_Onboard
type HwModule_Config_Location_Logging_Onboard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Disable interface{}
}

func (onboard *HwModule_Config_Location_Logging_Onboard) GetEntityData() *types.CommonEntityData {
    onboard.EntityData.YFilter = onboard.YFilter
    onboard.EntityData.YangName = "onboard"
    onboard.EntityData.BundleName = "cisco_ios_xr"
    onboard.EntityData.ParentYangName = "logging"
    onboard.EntityData.SegmentPath = "onboard"
    onboard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onboard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onboard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onboard.EntityData.Children = make(map[string]types.YChild)
    onboard.EntityData.Leafs = make(map[string]types.YLeaf)
    onboard.EntityData.Leafs["disable"] = types.YLeaf{"Disable", onboard.Disable}
    return &(onboard.EntityData)
}

// HwModule_Oper
type HwModule_Oper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of HwModule_Oper_Location.
    Location []HwModule_Oper_Location
}

func (oper *HwModule_Oper) GetEntityData() *types.CommonEntityData {
    oper.EntityData.YFilter = oper.YFilter
    oper.EntityData.YangName = "oper"
    oper.EntityData.BundleName = "cisco_ios_xr"
    oper.EntityData.ParentYangName = "hw-module"
    oper.EntityData.SegmentPath = "oper"
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = make(map[string]types.YChild)
    oper.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range oper.Location {
        oper.EntityData.Children[types.GetSegmentPath(&oper.Location[i])] = types.YChild{"Location", &oper.Location[i]}
    }
    oper.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oper.EntityData)
}

// HwModule_Oper_Location
type HwModule_Oper_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    
    Actions HwModule_Oper_Location_Actions

    
    Show HwModule_Oper_Location_Show
}

func (location *HwModule_Oper_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "oper"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["actions"] = types.YChild{"Actions", &location.Actions}
    location.EntityData.Children["show"] = types.YChild{"Show", &location.Show}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// HwModule_Oper_Location_Actions
type HwModule_Oper_Location_Actions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cbootmedia HwModule_Oper_Location_Actions_Cbootmedia
}

func (actions *HwModule_Oper_Location_Actions) GetEntityData() *types.CommonEntityData {
    actions.EntityData.YFilter = actions.YFilter
    actions.EntityData.YangName = "actions"
    actions.EntityData.BundleName = "cisco_ios_xr"
    actions.EntityData.ParentYangName = "location"
    actions.EntityData.SegmentPath = "actions"
    actions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actions.EntityData.Children = make(map[string]types.YChild)
    actions.EntityData.Children["cbootmedia"] = types.YChild{"Cbootmedia", &actions.Cbootmedia}
    actions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(actions.EntityData)
}

// HwModule_Oper_Location_Actions_Cbootmedia
type HwModule_Oper_Location_Actions_Cbootmedia struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of HwModule_Oper_Location_Actions_Cbootmedia_Bootmedia.
    Bootmedia []HwModule_Oper_Location_Actions_Cbootmedia_Bootmedia
}

func (cbootmedia *HwModule_Oper_Location_Actions_Cbootmedia) GetEntityData() *types.CommonEntityData {
    cbootmedia.EntityData.YFilter = cbootmedia.YFilter
    cbootmedia.EntityData.YangName = "cbootmedia"
    cbootmedia.EntityData.BundleName = "cisco_ios_xr"
    cbootmedia.EntityData.ParentYangName = "actions"
    cbootmedia.EntityData.SegmentPath = "cbootmedia"
    cbootmedia.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbootmedia.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbootmedia.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbootmedia.EntityData.Children = make(map[string]types.YChild)
    cbootmedia.EntityData.Children["bootmedia"] = types.YChild{"Bootmedia", nil}
    for i := range cbootmedia.Bootmedia {
        cbootmedia.EntityData.Children[types.GetSegmentPath(&cbootmedia.Bootmedia[i])] = types.YChild{"Bootmedia", &cbootmedia.Bootmedia[i]}
    }
    cbootmedia.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbootmedia.EntityData)
}

// HwModule_Oper_Location_Actions_Cbootmedia_Bootmedia
type HwModule_Oper_Location_Actions_Cbootmedia_Bootmedia struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Bootmedium interface{}
}

func (bootmedia *HwModule_Oper_Location_Actions_Cbootmedia_Bootmedia) GetEntityData() *types.CommonEntityData {
    bootmedia.EntityData.YFilter = bootmedia.YFilter
    bootmedia.EntityData.YangName = "bootmedia"
    bootmedia.EntityData.BundleName = "cisco_ios_xr"
    bootmedia.EntityData.ParentYangName = "cbootmedia"
    bootmedia.EntityData.SegmentPath = "bootmedia" + "[bootmedium='" + fmt.Sprintf("%v", bootmedia.Bootmedium) + "']"
    bootmedia.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootmedia.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootmedia.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootmedia.EntityData.Children = make(map[string]types.YChild)
    bootmedia.EntityData.Leafs = make(map[string]types.YLeaf)
    bootmedia.EntityData.Leafs["bootmedium"] = types.YLeaf{"Bootmedium", bootmedia.Bootmedium}
    return &(bootmedia.EntityData)
}

// HwModule_Oper_Location_Show
type HwModule_Oper_Location_Show struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (show *HwModule_Oper_Location_Show) GetEntityData() *types.CommonEntityData {
    show.EntityData.YFilter = show.YFilter
    show.EntityData.YangName = "show"
    show.EntityData.BundleName = "cisco_ios_xr"
    show.EntityData.ParentYangName = "location"
    show.EntityData.SegmentPath = "show"
    show.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    show.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    show.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    show.EntityData.Children = make(map[string]types.YChild)
    show.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(show.EntityData)
}

// HwModule_Shhwfpd
type HwModule_Shhwfpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of HwModule_Shhwfpd_Alocation.
    Alocation []HwModule_Shhwfpd_Alocation

    // The type is slice of HwModule_Shhwfpd_Fpd.
    Fpd []HwModule_Shhwfpd_Fpd
}

func (shhwfpd *HwModule_Shhwfpd) GetEntityData() *types.CommonEntityData {
    shhwfpd.EntityData.YFilter = shhwfpd.YFilter
    shhwfpd.EntityData.YangName = "shhwfpd"
    shhwfpd.EntityData.BundleName = "cisco_ios_xr"
    shhwfpd.EntityData.ParentYangName = "hw-module"
    shhwfpd.EntityData.SegmentPath = "shhwfpd"
    shhwfpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shhwfpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shhwfpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shhwfpd.EntityData.Children = make(map[string]types.YChild)
    shhwfpd.EntityData.Children["alocation"] = types.YChild{"Alocation", nil}
    for i := range shhwfpd.Alocation {
        shhwfpd.EntityData.Children[types.GetSegmentPath(&shhwfpd.Alocation[i])] = types.YChild{"Alocation", &shhwfpd.Alocation[i]}
    }
    shhwfpd.EntityData.Children["fpd"] = types.YChild{"Fpd", nil}
    for i := range shhwfpd.Fpd {
        shhwfpd.EntityData.Children[types.GetSegmentPath(&shhwfpd.Fpd[i])] = types.YChild{"Fpd", &shhwfpd.Fpd[i]}
    }
    shhwfpd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shhwfpd.EntityData)
}

// HwModule_Shhwfpd_Alocation
type HwModule_Shhwfpd_Alocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Locs interface{}

    // The type is slice of HwModule_Shhwfpd_Alocation_Fpd.
    Fpd []HwModule_Shhwfpd_Alocation_Fpd
}

func (alocation *HwModule_Shhwfpd_Alocation) GetEntityData() *types.CommonEntityData {
    alocation.EntityData.YFilter = alocation.YFilter
    alocation.EntityData.YangName = "alocation"
    alocation.EntityData.BundleName = "cisco_ios_xr"
    alocation.EntityData.ParentYangName = "shhwfpd"
    alocation.EntityData.SegmentPath = "alocation" + "[locs='" + fmt.Sprintf("%v", alocation.Locs) + "']"
    alocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alocation.EntityData.Children = make(map[string]types.YChild)
    alocation.EntityData.Children["fpd"] = types.YChild{"Fpd", nil}
    for i := range alocation.Fpd {
        alocation.EntityData.Children[types.GetSegmentPath(&alocation.Fpd[i])] = types.YChild{"Fpd", &alocation.Fpd[i]}
    }
    alocation.EntityData.Leafs = make(map[string]types.YLeaf)
    alocation.EntityData.Leafs["locs"] = types.YLeaf{"Locs", alocation.Locs}
    return &(alocation.EntityData)
}

// HwModule_Shhwfpd_Alocation_Fpd
type HwModule_Shhwfpd_Alocation_Fpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Fpdname interface{}

    // Status. The type is string.
    State interface{}

    // HWver. The type is string.
    Hwver interface{}

    // FPD ver. The type is string.
    Fpdver interface{}

    // Card type. The type is string.
    Cardname interface{}

    // Attribute. The type is string.
    Attr interface{}

    // FPD Programed. The type is string.
    Fpddnld interface{}
}

func (fpd *HwModule_Shhwfpd_Alocation_Fpd) GetEntityData() *types.CommonEntityData {
    fpd.EntityData.YFilter = fpd.YFilter
    fpd.EntityData.YangName = "fpd"
    fpd.EntityData.BundleName = "cisco_ios_xr"
    fpd.EntityData.ParentYangName = "alocation"
    fpd.EntityData.SegmentPath = "fpd" + "[fpdname='" + fmt.Sprintf("%v", fpd.Fpdname) + "']"
    fpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd.EntityData.Children = make(map[string]types.YChild)
    fpd.EntityData.Leafs = make(map[string]types.YLeaf)
    fpd.EntityData.Leafs["fpdname"] = types.YLeaf{"Fpdname", fpd.Fpdname}
    fpd.EntityData.Leafs["state"] = types.YLeaf{"State", fpd.State}
    fpd.EntityData.Leafs["hwver"] = types.YLeaf{"Hwver", fpd.Hwver}
    fpd.EntityData.Leafs["fpdver"] = types.YLeaf{"Fpdver", fpd.Fpdver}
    fpd.EntityData.Leafs["cardname"] = types.YLeaf{"Cardname", fpd.Cardname}
    fpd.EntityData.Leafs["attr"] = types.YLeaf{"Attr", fpd.Attr}
    fpd.EntityData.Leafs["fpddnld"] = types.YLeaf{"Fpddnld", fpd.Fpddnld}
    return &(fpd.EntityData)
}

// HwModule_Shhwfpd_Fpd
type HwModule_Shhwfpd_Fpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Fpdname interface{}
}

func (fpd *HwModule_Shhwfpd_Fpd) GetEntityData() *types.CommonEntityData {
    fpd.EntityData.YFilter = fpd.YFilter
    fpd.EntityData.YangName = "fpd"
    fpd.EntityData.BundleName = "cisco_ios_xr"
    fpd.EntityData.ParentYangName = "shhwfpd"
    fpd.EntityData.SegmentPath = "fpd" + "[fpdname='" + fmt.Sprintf("%v", fpd.Fpdname) + "']"
    fpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd.EntityData.Children = make(map[string]types.YChild)
    fpd.EntityData.Leafs = make(map[string]types.YLeaf)
    fpd.EntityData.Leafs["fpdname"] = types.YLeaf{"Fpdname", fpd.Fpdname}
    return &(fpd.EntityData)
}

