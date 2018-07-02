// This module contains a collection of YANG definitions
// for Cisco IOS-XR ocni package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ocni-ni-base: An OpenConfig description of a network-instance.
//     This may be a Layer 3 forwarding construct such as a virtual
//     routing and forwarding (VRF) instance, or a Layer 2 instance
//     such as a virtual switch instance (VSI). Mixed Layer 2 and
//     Layer 3 instances are also supaported. Copyright (c) 2017 by
//     Cisco Systems, Inc.  All rights reserved.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ocni_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ocni_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ocni-oper ocni-ni-base}", reflect.TypeOf(OcniNiBase{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ocni-oper:ocni-ni-base", reflect.TypeOf(OcniNiBase{}))
}

// OcniNiBase
// An OpenConfig description of a network-instance.
// This may be a Layer 3 forwarding construct such
// as a virtual routing and forwarding (VRF)
// instance, or a Layer 2 instance such as a virtual
// switch instance (VSI). Mixed Layer 2 and Layer 3
// instances are also supaported. Copyright (c) 2017
// by Cisco Systems, Inc.  All rights reserved.
type OcniNiBase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system.
    NetworkInstances OcniNiBase_NetworkInstances
}

func (ocniNiBase *OcniNiBase) GetEntityData() *types.CommonEntityData {
    ocniNiBase.EntityData.YFilter = ocniNiBase.YFilter
    ocniNiBase.EntityData.YangName = "ocni-ni-base"
    ocniNiBase.EntityData.BundleName = "cisco_ios_xr"
    ocniNiBase.EntityData.ParentYangName = "Cisco-IOS-XR-ocni-oper"
    ocniNiBase.EntityData.SegmentPath = "Cisco-IOS-XR-ocni-oper:ocni-ni-base"
    ocniNiBase.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ocniNiBase.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ocniNiBase.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ocniNiBase.EntityData.Children = types.NewOrderedMap()
    ocniNiBase.EntityData.Children.Append("network-instances", types.YChild{"NetworkInstances", &ocniNiBase.NetworkInstances})
    ocniNiBase.EntityData.Leafs = types.NewOrderedMap()

    ocniNiBase.EntityData.YListKeys = []string {}

    return &(ocniNiBase.EntityData)
}

// OcniNiBase_NetworkInstances
// Network instances configured on the local system
type OcniNiBase_NetworkInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system. The type is slice of
    // OcniNiBase_NetworkInstances_NetworkInstance.
    NetworkInstance []*OcniNiBase_NetworkInstances_NetworkInstance
}

func (networkInstances *OcniNiBase_NetworkInstances) GetEntityData() *types.CommonEntityData {
    networkInstances.EntityData.YFilter = networkInstances.YFilter
    networkInstances.EntityData.YangName = "network-instances"
    networkInstances.EntityData.BundleName = "cisco_ios_xr"
    networkInstances.EntityData.ParentYangName = "ocni-ni-base"
    networkInstances.EntityData.SegmentPath = "network-instances"
    networkInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkInstances.EntityData.Children = types.NewOrderedMap()
    networkInstances.EntityData.Children.Append("network-instance", types.YChild{"NetworkInstance", nil})
    for i := range networkInstances.NetworkInstance {
        networkInstances.EntityData.Children.Append(types.GetSegmentPath(networkInstances.NetworkInstance[i]), types.YChild{"NetworkInstance", networkInstances.NetworkInstance[i]})
    }
    networkInstances.EntityData.Leafs = types.NewOrderedMap()

    networkInstances.EntityData.YListKeys = []string {}

    return &(networkInstances.EntityData)
}

// OcniNiBase_NetworkInstances_NetworkInstance
// Network instances configured on the local
// system
type OcniNiBase_NetworkInstances_NetworkInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique name identifying the network instance.
    // The type is string.
    Name interface{}

    // Operational state parameters relating to a network instance.
    State OcniNiBase_NetworkInstances_NetworkInstance_State
}

func (networkInstance *OcniNiBase_NetworkInstances_NetworkInstance) GetEntityData() *types.CommonEntityData {
    networkInstance.EntityData.YFilter = networkInstance.YFilter
    networkInstance.EntityData.YangName = "network-instance"
    networkInstance.EntityData.BundleName = "cisco_ios_xr"
    networkInstance.EntityData.ParentYangName = "network-instances"
    networkInstance.EntityData.SegmentPath = "network-instance" + types.AddKeyToken(networkInstance.Name, "name")
    networkInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkInstance.EntityData.Children = types.NewOrderedMap()
    networkInstance.EntityData.Children.Append("state", types.YChild{"State", &networkInstance.State})
    networkInstance.EntityData.Leafs = types.NewOrderedMap()
    networkInstance.EntityData.Leafs.Append("name", types.YLeaf{"Name", networkInstance.Name})

    networkInstance.EntityData.YListKeys = []string {"Name"}

    return &(networkInstance.EntityData)
}

// OcniNiBase_NetworkInstances_NetworkInstance_State
// Operational state parameters relating to a
// network instance
type OcniNiBase_NetworkInstances_NetworkInstance_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An operator-assigned unique name for the forwarding instance. The type is
    // string.
    Name interface{}

    // The type of network instance. The value of this leaf indicates the type of
    // forwarding entries that should be supported by this network instance. The
    // type is string.
    Type interface{}

    // Whether the network instance should be configured to be active on the
    // network element. The type is bool.
    Enabled interface{}

    // A free-form string to be used by the network operator to describe the
    // function of this network instance. The type is string.
    Description interface{}

    // The address families that are to be enabled for this network instance. The
    // type is slice of string.
    EnabledAddressFamily []interface{}
}

func (state *OcniNiBase_NetworkInstances_NetworkInstance_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "network-instance"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("enabled-address-family", types.YLeaf{"EnabledAddressFamily", state.EnabledAddressFamily})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

