// This module contains a collection of YANG definitions
// for Cisco IOS-XR fretta-bcm-dpa-oor package configuration.
// 
// This module contains definitions
// for the following management objects:
//   oor-hw-config: Out-Of-Resource (OOR) configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fretta_bcm_dpa_oor_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fretta_bcm_dpa_oor_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fretta-bcm-dpa-oor-cfg oor-hw-config}", reflect.TypeOf(OorHwConfig{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fretta-bcm-dpa-oor-cfg:oor-hw-config", reflect.TypeOf(OorHwConfig{}))
}

// OorHwConfig
// Out-Of-Resource (OOR) configuration
type OorHwConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hardware OOR configuration.
    Hw OorHwConfig_Hw
}

func (oorHwConfig *OorHwConfig) GetEntityData() *types.CommonEntityData {
    oorHwConfig.EntityData.YFilter = oorHwConfig.YFilter
    oorHwConfig.EntityData.YangName = "oor-hw-config"
    oorHwConfig.EntityData.BundleName = "cisco_ios_xr"
    oorHwConfig.EntityData.ParentYangName = "Cisco-IOS-XR-fretta-bcm-dpa-oor-cfg"
    oorHwConfig.EntityData.SegmentPath = "Cisco-IOS-XR-fretta-bcm-dpa-oor-cfg:oor-hw-config"
    oorHwConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorHwConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorHwConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorHwConfig.EntityData.Children = make(map[string]types.YChild)
    oorHwConfig.EntityData.Children["hw"] = types.YChild{"Hw", &oorHwConfig.Hw}
    oorHwConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oorHwConfig.EntityData)
}

// OorHwConfig_Hw
// Hardware OOR configuration
type OorHwConfig_Hw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure OOR hardware dampening timeouts. The type is interface{} with
    // range: 0..600. Units are second.
    Dampening interface{}

    // Configure OOR hardware thresholds.
    Thresholds OorHwConfig_Hw_Thresholds
}

func (hw *OorHwConfig_Hw) GetEntityData() *types.CommonEntityData {
    hw.EntityData.YFilter = hw.YFilter
    hw.EntityData.YangName = "hw"
    hw.EntityData.BundleName = "cisco_ios_xr"
    hw.EntityData.ParentYangName = "oor-hw-config"
    hw.EntityData.SegmentPath = "hw"
    hw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hw.EntityData.Children = make(map[string]types.YChild)
    hw.EntityData.Children["thresholds"] = types.YChild{"Thresholds", &hw.Thresholds}
    hw.EntityData.Leafs = make(map[string]types.YLeaf)
    hw.EntityData.Leafs["dampening"] = types.YLeaf{"Dampening", hw.Dampening}
    return &(hw.EntityData)
}

// OorHwConfig_Hw_Thresholds
// Configure OOR hardware thresholds
type OorHwConfig_Hw_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of OorHwConfig_Hw_Thresholds_Threshold.
    Threshold []OorHwConfig_Hw_Thresholds_Threshold
}

func (thresholds *OorHwConfig_Hw_Thresholds) GetEntityData() *types.CommonEntityData {
    thresholds.EntityData.YFilter = thresholds.YFilter
    thresholds.EntityData.YangName = "thresholds"
    thresholds.EntityData.BundleName = "cisco_ios_xr"
    thresholds.EntityData.ParentYangName = "hw"
    thresholds.EntityData.SegmentPath = "thresholds"
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholds.EntityData.Children = make(map[string]types.YChild)
    thresholds.EntityData.Children["threshold"] = types.YChild{"Threshold", nil}
    for i := range thresholds.Threshold {
        thresholds.EntityData.Children[types.GetSegmentPath(&thresholds.Threshold[i])] = types.YChild{"Threshold", &thresholds.Threshold[i]}
    }
    thresholds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholds.EntityData)
}

// OorHwConfig_Hw_Thresholds_Threshold
// none
type OorHwConfig_Hw_Thresholds_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Color interface{}

    // value in percent. The type is interface{} with range: 0..100. This
    // attribute is mandatory. Units are percentage.
    Percent interface{}
}

func (threshold *OorHwConfig_Hw_Thresholds_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "thresholds"
    threshold.EntityData.SegmentPath = "threshold" + "[color='" + fmt.Sprintf("%v", threshold.Color) + "']"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["color"] = types.YLeaf{"Color", threshold.Color}
    threshold.EntityData.Leafs["percent"] = types.YLeaf{"Percent", threshold.Percent}
    return &(threshold.EntityData)
}

