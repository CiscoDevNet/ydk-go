// This module contains definitions
// for the Calvados model objects.
// 
// This module contains the YANG definitions
// for the Cisco IOS-XR SysAdmin
// 'vm profile|cpu|memory' commands.
// 
// Copyright(c) 2018 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_vm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_vm"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-vm vm}", reflect.TypeOf(Vm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-vm:vm", reflect.TypeOf(Vm{}))
}

// Vm
type Vm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Config Vm_Config
}

func (vm *Vm) GetEntityData() *types.CommonEntityData {
    vm.EntityData.YFilter = vm.YFilter
    vm.EntityData.YangName = "vm"
    vm.EntityData.BundleName = "cisco_ios_xr"
    vm.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-vm"
    vm.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-vm:vm"
    vm.EntityData.AbsolutePath = vm.EntityData.SegmentPath
    vm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vm.EntityData.Children = types.NewOrderedMap()
    vm.EntityData.Children.Append("config", types.YChild{"Config", &vm.Config})
    vm.EntityData.Leafs = types.NewOrderedMap()

    vm.EntityData.YListKeys = []string {}

    return &(vm.EntityData)
}

// Vm_Config
type Vm_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    HwProfile Vm_Config_HwProfile

    
    Memory Vm_Config_Memory

    
    Cpu Vm_Config_Cpu
}

func (config *Vm_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "vm"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-vm:vm/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("hw-profile", types.YChild{"HwProfile", &config.HwProfile})
    config.EntityData.Children.Append("memory", types.YChild{"Memory", &config.Memory})
    config.EntityData.Children.Append("cpu", types.YChild{"Cpu", &config.Cpu})
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Vm_Config_HwProfile
type Vm_Config_HwProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // xrv9k profile vpe|vrr. The type is Profile.
    Profile interface{}
}

func (hwProfile *Vm_Config_HwProfile) GetEntityData() *types.CommonEntityData {
    hwProfile.EntityData.YFilter = hwProfile.YFilter
    hwProfile.EntityData.YangName = "hw-profile"
    hwProfile.EntityData.BundleName = "cisco_ios_xr"
    hwProfile.EntityData.ParentYangName = "config"
    hwProfile.EntityData.SegmentPath = "hw-profile"
    hwProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-vm:vm/config/" + hwProfile.EntityData.SegmentPath
    hwProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwProfile.EntityData.Children = types.NewOrderedMap()
    hwProfile.EntityData.Leafs = types.NewOrderedMap()
    hwProfile.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", hwProfile.Profile})

    hwProfile.EntityData.YListKeys = []string {}

    return &(hwProfile.EntityData)
}

// Vm_Config_HwProfile_Profile represents xrv9k profile vpe|vrr
type Vm_Config_HwProfile_Profile string

const (
    Vm_Config_HwProfile_Profile_vrr Vm_Config_HwProfile_Profile = "vrr"
)

// Vm_Config_Memory
type Vm_Config_Memory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // admin container memory in GB. The type is interface{} with range:
    // 0..4294967295.
    Admin interface{}

    // rp container memory in GB. The type is interface{} with range:
    // 0..4294967295.
    Rp interface{}

    // lc container memory in GB. The type is interface{} with range:
    // 0..4294967295.
    Lc interface{}
}

func (memory *Vm_Config_Memory) GetEntityData() *types.CommonEntityData {
    memory.EntityData.YFilter = memory.YFilter
    memory.EntityData.YangName = "memory"
    memory.EntityData.BundleName = "cisco_ios_xr"
    memory.EntityData.ParentYangName = "config"
    memory.EntityData.SegmentPath = "memory"
    memory.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-vm:vm/config/" + memory.EntityData.SegmentPath
    memory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memory.EntityData.Children = types.NewOrderedMap()
    memory.EntityData.Leafs = types.NewOrderedMap()
    memory.EntityData.Leafs.Append("admin", types.YLeaf{"Admin", memory.Admin})
    memory.EntityData.Leafs.Append("rp", types.YLeaf{"Rp", memory.Rp})
    memory.EntityData.Leafs.Append("lc", types.YLeaf{"Lc", memory.Lc})

    memory.EntityData.YListKeys = []string {}

    return &(memory.EntityData)
}

// Vm_Config_Cpu
type Vm_Config_Cpu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // assign cpu cores to control/data plane. The type is string with pattern:
    // b'0(-[0-9]+)?/[0-9]+(-[0-9]+)?'.
    Assign interface{}
}

func (cpu *Vm_Config_Cpu) GetEntityData() *types.CommonEntityData {
    cpu.EntityData.YFilter = cpu.YFilter
    cpu.EntityData.YangName = "cpu"
    cpu.EntityData.BundleName = "cisco_ios_xr"
    cpu.EntityData.ParentYangName = "config"
    cpu.EntityData.SegmentPath = "cpu"
    cpu.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-vm:vm/config/" + cpu.EntityData.SegmentPath
    cpu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpu.EntityData.Children = types.NewOrderedMap()
    cpu.EntityData.Leafs = types.NewOrderedMap()
    cpu.EntityData.Leafs.Append("assign", types.YLeaf{"Assign", cpu.Assign})

    cpu.EntityData.YListKeys = []string {}

    return &(cpu.EntityData)
}

