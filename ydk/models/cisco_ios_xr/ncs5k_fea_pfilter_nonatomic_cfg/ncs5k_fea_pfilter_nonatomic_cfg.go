// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5k-fea-pfilter-nonatomic package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware: Hardware
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs5k_fea_pfilter_nonatomic_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs5k_fea_pfilter_nonatomic_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs5k-fea-pfilter-nonatomic-cfg hardware}", reflect.TypeOf(Hardware{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs5k-fea-pfilter-nonatomic-cfg:hardware", reflect.TypeOf(Hardware{}))
}

// AtomicDisableDfltActn represents Atomic disable dflt actn
type AtomicDisableDfltActn string

const (
    // Drops traffic during modification
    AtomicDisableDfltActn_default_action_deny AtomicDisableDfltActn = "default-action-deny"

    // Forward traffic during modification
    AtomicDisableDfltActn_default_action_permit AtomicDisableDfltActn = "default-action-permit"
)

// Hardware
// Hardware
type Hardware struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access-list option.
    AccessList Hardware_AccessList
}

func (hardware *Hardware) GetEntityData() *types.CommonEntityData {
    hardware.EntityData.YFilter = hardware.YFilter
    hardware.EntityData.YangName = "hardware"
    hardware.EntityData.BundleName = "cisco_ios_xr"
    hardware.EntityData.ParentYangName = "Cisco-IOS-XR-ncs5k-fea-pfilter-nonatomic-cfg"
    hardware.EntityData.SegmentPath = "Cisco-IOS-XR-ncs5k-fea-pfilter-nonatomic-cfg:hardware"
    hardware.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardware.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardware.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardware.EntityData.Children = types.NewOrderedMap()
    hardware.EntityData.Children.Append("access-list", types.YChild{"AccessList", &hardware.AccessList})
    hardware.EntityData.Leafs = types.NewOrderedMap()

    hardware.EntityData.YListKeys = []string {}

    return &(hardware.EntityData)
}

// Hardware_AccessList
// Access-list option
type Hardware_AccessList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify Option for Atomic disable. The type is AtomicDisableDfltActn.
    AtomicDisable interface{}
}

func (accessList *Hardware_AccessList) GetEntityData() *types.CommonEntityData {
    accessList.EntityData.YFilter = accessList.YFilter
    accessList.EntityData.YangName = "access-list"
    accessList.EntityData.BundleName = "cisco_ios_xr"
    accessList.EntityData.ParentYangName = "hardware"
    accessList.EntityData.SegmentPath = "access-list"
    accessList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessList.EntityData.Children = types.NewOrderedMap()
    accessList.EntityData.Leafs = types.NewOrderedMap()
    accessList.EntityData.Leafs.Append("atomic-disable", types.YLeaf{"AtomicDisable", accessList.AtomicDisable})

    accessList.EntityData.YListKeys = []string {}

    return &(accessList.EntityData)
}

