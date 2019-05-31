// This module contains definitions
// for the Calvados model objects.
// 
// This module holds a sample validation.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package valtest

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package valtest"))
    ydk.RegisterEntity("{http://www.cisco.com/panini/calvados/valtest config}", reflect.TypeOf(Config{}))
    ydk.RegisterEntity("valtest:config", reflect.TypeOf(Config{}))
}

// Config
type Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Valtest Config_Valtest
}

func (config *Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "valtest"
    config.EntityData.SegmentPath = "valtest:config"
    config.EntityData.AbsolutePath = config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("valtest", types.YChild{"Valtest", &config.Valtest})
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Config_Valtest
type Config_Valtest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range:
    // -9223372036854775808..9223372036854775807. The default value is 42.
    ANumber interface{}

    // The type is interface{} with range:
    // -9223372036854775808..9223372036854775807. The default value is 7.
    BNumber interface{}
}

func (valtest *Config_Valtest) GetEntityData() *types.CommonEntityData {
    valtest.EntityData.YFilter = valtest.YFilter
    valtest.EntityData.YangName = "valtest"
    valtest.EntityData.BundleName = "cisco_ios_xr"
    valtest.EntityData.ParentYangName = "config"
    valtest.EntityData.SegmentPath = "valtest"
    valtest.EntityData.AbsolutePath = "valtest:config/" + valtest.EntityData.SegmentPath
    valtest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    valtest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    valtest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    valtest.EntityData.Children = types.NewOrderedMap()
    valtest.EntityData.Leafs = types.NewOrderedMap()
    valtest.EntityData.Leafs.Append("a_number", types.YLeaf{"ANumber", valtest.ANumber})
    valtest.EntityData.Leafs.Append("b_number", types.YLeaf{"BNumber", valtest.BNumber})

    valtest.EntityData.YListKeys = []string {}

    return &(valtest.EntityData)
}

