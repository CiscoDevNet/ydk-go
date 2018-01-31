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
    parent types.Entity
    YFilter yfilter.YFilter

    // Access-list option.
    AccessList Hardware_AccessList
}

func (hardware *Hardware) GetFilter() yfilter.YFilter { return hardware.YFilter }

func (hardware *Hardware) SetFilter(yf yfilter.YFilter) { hardware.YFilter = yf }

func (hardware *Hardware) GetGoName(yname string) string {
    if yname == "access-list" { return "AccessList" }
    return ""
}

func (hardware *Hardware) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs5k-fea-pfilter-nonatomic-cfg:hardware"
}

func (hardware *Hardware) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list" {
        return &hardware.AccessList
    }
    return nil
}

func (hardware *Hardware) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list"] = &hardware.AccessList
    return children
}

func (hardware *Hardware) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardware *Hardware) GetBundleName() string { return "cisco_ios_xr" }

func (hardware *Hardware) GetYangName() string { return "hardware" }

func (hardware *Hardware) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardware *Hardware) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardware *Hardware) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardware *Hardware) SetParent(parent types.Entity) { hardware.parent = parent }

func (hardware *Hardware) GetParent() types.Entity { return hardware.parent }

func (hardware *Hardware) GetParentYangName() string { return "Cisco-IOS-XR-ncs5k-fea-pfilter-nonatomic-cfg" }

// Hardware_AccessList
// Access-list option
type Hardware_AccessList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify Option for Atomic disable. The type is AtomicDisableDfltActn.
    AtomicDisable interface{}
}

func (accessList *Hardware_AccessList) GetFilter() yfilter.YFilter { return accessList.YFilter }

func (accessList *Hardware_AccessList) SetFilter(yf yfilter.YFilter) { accessList.YFilter = yf }

func (accessList *Hardware_AccessList) GetGoName(yname string) string {
    if yname == "atomic-disable" { return "AtomicDisable" }
    return ""
}

func (accessList *Hardware_AccessList) GetSegmentPath() string {
    return "access-list"
}

func (accessList *Hardware_AccessList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessList *Hardware_AccessList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessList *Hardware_AccessList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["atomic-disable"] = accessList.AtomicDisable
    return leafs
}

func (accessList *Hardware_AccessList) GetBundleName() string { return "cisco_ios_xr" }

func (accessList *Hardware_AccessList) GetYangName() string { return "access-list" }

func (accessList *Hardware_AccessList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessList *Hardware_AccessList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessList *Hardware_AccessList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessList *Hardware_AccessList) SetParent(parent types.Entity) { accessList.parent = parent }

func (accessList *Hardware_AccessList) GetParent() types.Entity { return accessList.parent }

func (accessList *Hardware_AccessList) GetParentYangName() string { return "hardware" }

