// This module contains definitions
// for the Calvados model objects.
// 
// FPD CLI support for both oper and config
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_fpd_infra_cli_fpd

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_fpd_infra_cli_fpd"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd fpd}", reflect.TypeOf(Fpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd:fpd", reflect.TypeOf(Fpd{}))
}

// Fpd
type Fpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // fpd config mode.
    Config Fpd_Config
}

func (fpd *Fpd) GetEntityData() *types.CommonEntityData {
    fpd.EntityData.YFilter = fpd.YFilter
    fpd.EntityData.YangName = "fpd"
    fpd.EntityData.BundleName = "cisco_ios_xr"
    fpd.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd"
    fpd.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd:fpd"
    fpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd.EntityData.Children = types.NewOrderedMap()
    fpd.EntityData.Children.Append("config", types.YChild{"Config", &fpd.Config})
    fpd.EntityData.Leafs = types.NewOrderedMap()

    fpd.EntityData.YListKeys = []string {}

    return &(fpd.EntityData)
}

// Fpd_Config
// fpd config mode
type Fpd_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is AutoUpgrade. The default value is disable.
    AutoUpgrade interface{}
}

func (config *Fpd_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "fpd"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("auto-upgrade", types.YLeaf{"AutoUpgrade", config.AutoUpgrade})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Fpd_Config_AutoUpgrade
type Fpd_Config_AutoUpgrade string

const (
    Fpd_Config_AutoUpgrade_enable Fpd_Config_AutoUpgrade = "enable"

    Fpd_Config_AutoUpgrade_disable Fpd_Config_AutoUpgrade = "disable"
)

