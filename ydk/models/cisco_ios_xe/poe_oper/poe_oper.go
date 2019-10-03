// This module contains a collection of YANG definitions for
// monitoring power over ethernet feature in a Network Element.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package poe_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package poe_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-poe-oper poe-oper-data}", reflect.TypeOf(PoeOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-poe-oper:poe-oper-data", reflect.TypeOf(PoeOperData{}))
}

// IlpowerPdClass represents Name of the power class
type IlpowerPdClass string

const (
    // List of POE interfaces, keyed by interface name
    IlpowerPdClass_poe_null IlpowerPdClass = "poe-null"

    // Power class unknown
    IlpowerPdClass_poe_unknown IlpowerPdClass = "poe-unknown"

    // Power class cisco
    IlpowerPdClass_poe_cisco IlpowerPdClass = "poe-cisco"

    // IEEE power class 0
    IlpowerPdClass_poe_ieee0 IlpowerPdClass = "poe-ieee0"

    // IEEE power class 1
    IlpowerPdClass_poe_ieee1 IlpowerPdClass = "poe-ieee1"

    // IEEE power class 2
    IlpowerPdClass_poe_ieee2 IlpowerPdClass = "poe-ieee2"

    // IEEE power class 3
    IlpowerPdClass_poe_ieee3 IlpowerPdClass = "poe-ieee3"

    // IEEE power class 4
    IlpowerPdClass_poe_ieee4 IlpowerPdClass = "poe-ieee4"

    // IEEE power class 5
    IlpowerPdClass_poe_ieee5 IlpowerPdClass = "poe-ieee5"

    // IEEE power class unknown
    IlpowerPdClass_poe_ieee_unknown_class IlpowerPdClass = "poe-ieee-unknown-class"
)

// PoeOperData
// Informaton about POEs
type PoeOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of POE interfaces, keyed by interface name. The type is slice of
    // PoeOperData_PoePort.
    PoePort []*PoeOperData_PoePort
}

func (poeOperData *PoeOperData) GetEntityData() *types.CommonEntityData {
    poeOperData.EntityData.YFilter = poeOperData.YFilter
    poeOperData.EntityData.YangName = "poe-oper-data"
    poeOperData.EntityData.BundleName = "cisco_ios_xe"
    poeOperData.EntityData.ParentYangName = "Cisco-IOS-XE-poe-oper"
    poeOperData.EntityData.SegmentPath = "Cisco-IOS-XE-poe-oper:poe-oper-data"
    poeOperData.EntityData.AbsolutePath = poeOperData.EntityData.SegmentPath
    poeOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    poeOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    poeOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    poeOperData.EntityData.Children = types.NewOrderedMap()
    poeOperData.EntityData.Children.Append("poe-port", types.YChild{"PoePort", nil})
    for i := range poeOperData.PoePort {
        poeOperData.EntityData.Children.Append(types.GetSegmentPath(poeOperData.PoePort[i]), types.YChild{"PoePort", poeOperData.PoePort[i]})
    }
    poeOperData.EntityData.Leafs = types.NewOrderedMap()

    poeOperData.EntityData.YListKeys = []string {}

    return &(poeOperData.EntityData)
}

// PoeOperData_PoePort
// List of POE interfaces, keyed by interface name
type PoeOperData_PoePort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the POE interface. The type is string.
    IntfName interface{}

    // POE interface admin state. The type is bool.
    PoeIntfEnabled interface{}

    // Power used by PD device. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    PowerUsed interface{}

    // Class of the PD device. The type is IlpowerPdClass.
    PdClass interface{}
}

func (poePort *PoeOperData_PoePort) GetEntityData() *types.CommonEntityData {
    poePort.EntityData.YFilter = poePort.YFilter
    poePort.EntityData.YangName = "poe-port"
    poePort.EntityData.BundleName = "cisco_ios_xe"
    poePort.EntityData.ParentYangName = "poe-oper-data"
    poePort.EntityData.SegmentPath = "poe-port" + types.AddKeyToken(poePort.IntfName, "intf-name")
    poePort.EntityData.AbsolutePath = "Cisco-IOS-XE-poe-oper:poe-oper-data/" + poePort.EntityData.SegmentPath
    poePort.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    poePort.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    poePort.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    poePort.EntityData.Children = types.NewOrderedMap()
    poePort.EntityData.Leafs = types.NewOrderedMap()
    poePort.EntityData.Leafs.Append("intf-name", types.YLeaf{"IntfName", poePort.IntfName})
    poePort.EntityData.Leafs.Append("poe-intf-enabled", types.YLeaf{"PoeIntfEnabled", poePort.PoeIntfEnabled})
    poePort.EntityData.Leafs.Append("power-used", types.YLeaf{"PowerUsed", poePort.PowerUsed})
    poePort.EntityData.Leafs.Append("pd-class", types.YLeaf{"PdClass", poePort.PdClass})

    poePort.EntityData.YListKeys = []string {"IntfName"}

    return &(poePort.EntityData)
}

