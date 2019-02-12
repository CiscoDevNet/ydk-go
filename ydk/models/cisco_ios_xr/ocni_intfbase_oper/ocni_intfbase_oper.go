// This module contains a collection of YANG definitions
// for Cisco IOS-XR ocni-intfbase package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ocni-ni-intfbase: An OpenConfig description of a
//     network-instance. This may be a Layer 3 forwarding construct
//     such as a virtual routing and forwarding (VRF) instance, or
//     a Layer 2 instance such as a virtual switch instance (VSI).
//     Mixed Layer 2 and Layer 3 instances are also supported.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ocni_intfbase_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ocni_intfbase_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ocni-intfbase-oper ocni-ni-intfbase}", reflect.TypeOf(OcniNiIntfbase{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ocni-intfbase-oper:ocni-ni-intfbase", reflect.TypeOf(OcniNiIntfbase{}))
}

// OcniNiIntfbase
// An OpenConfig description of a network-instance.
// This may be a Layer 3 forwarding construct such
// as a virtual routing and forwarding (VRF)
// instance, or a Layer 2 instance such as a virtual
// switch instance (VSI). Mixed Layer 2 and Layer 3
// instances are also supported.
type OcniNiIntfbase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system.
    NetworkInstances OcniNiIntfbase_NetworkInstances
}

func (ocniNiIntfbase *OcniNiIntfbase) GetEntityData() *types.CommonEntityData {
    ocniNiIntfbase.EntityData.YFilter = ocniNiIntfbase.YFilter
    ocniNiIntfbase.EntityData.YangName = "ocni-ni-intfbase"
    ocniNiIntfbase.EntityData.BundleName = "cisco_ios_xr"
    ocniNiIntfbase.EntityData.ParentYangName = "Cisco-IOS-XR-ocni-intfbase-oper"
    ocniNiIntfbase.EntityData.SegmentPath = "Cisco-IOS-XR-ocni-intfbase-oper:ocni-ni-intfbase"
    ocniNiIntfbase.EntityData.AbsolutePath = ocniNiIntfbase.EntityData.SegmentPath
    ocniNiIntfbase.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ocniNiIntfbase.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ocniNiIntfbase.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ocniNiIntfbase.EntityData.Children = types.NewOrderedMap()
    ocniNiIntfbase.EntityData.Children.Append("network-instances", types.YChild{"NetworkInstances", &ocniNiIntfbase.NetworkInstances})
    ocniNiIntfbase.EntityData.Leafs = types.NewOrderedMap()

    ocniNiIntfbase.EntityData.YListKeys = []string {}

    return &(ocniNiIntfbase.EntityData)
}

// OcniNiIntfbase_NetworkInstances
// Network instances configured on the local system
type OcniNiIntfbase_NetworkInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network instances configured on the local system. The type is slice of
    // OcniNiIntfbase_NetworkInstances_NetworkInstance.
    NetworkInstance []*OcniNiIntfbase_NetworkInstances_NetworkInstance
}

func (networkInstances *OcniNiIntfbase_NetworkInstances) GetEntityData() *types.CommonEntityData {
    networkInstances.EntityData.YFilter = networkInstances.YFilter
    networkInstances.EntityData.YangName = "network-instances"
    networkInstances.EntityData.BundleName = "cisco_ios_xr"
    networkInstances.EntityData.ParentYangName = "ocni-ni-intfbase"
    networkInstances.EntityData.SegmentPath = "network-instances"
    networkInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-intfbase-oper:ocni-ni-intfbase/" + networkInstances.EntityData.SegmentPath
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

// OcniNiIntfbase_NetworkInstances_NetworkInstance
// Network instances configured on the local
// system
type OcniNiIntfbase_NetworkInstances_NetworkInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique name identifying the network instance.
    // The type is string.
    Name interface{}

    // An interface associated with the network instance.
    Interfaces OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces
}

func (networkInstance *OcniNiIntfbase_NetworkInstances_NetworkInstance) GetEntityData() *types.CommonEntityData {
    networkInstance.EntityData.YFilter = networkInstance.YFilter
    networkInstance.EntityData.YangName = "network-instance"
    networkInstance.EntityData.BundleName = "cisco_ios_xr"
    networkInstance.EntityData.ParentYangName = "network-instances"
    networkInstance.EntityData.SegmentPath = "network-instance" + types.AddKeyToken(networkInstance.Name, "name")
    networkInstance.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-intfbase-oper:ocni-ni-intfbase/network-instances/" + networkInstance.EntityData.SegmentPath
    networkInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkInstance.EntityData.Children = types.NewOrderedMap()
    networkInstance.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &networkInstance.Interfaces})
    networkInstance.EntityData.Leafs = types.NewOrderedMap()
    networkInstance.EntityData.Leafs.Append("name", types.YLeaf{"Name", networkInstance.Name})

    networkInstance.EntityData.YListKeys = []string {"Name"}

    return &(networkInstance.EntityData)
}

// OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces
// An interface associated with the network
// instance
type OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An interface associated with the network instance. The type is slice of
    // OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface.
    Interface []*OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface
}

func (interfaces *OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "network-instance"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-intfbase-oper:ocni-ni-intfbase/network-instances/network-instance/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface
// An interface associated with the network
// instance
type OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A reference to an identifier for this interface
    // which acts as a key for this list. The type is string.
    Id interface{}

    // Operational state parameters relating to the associated interface.
    State OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface_State
}

func (self *OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Id, "id")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-intfbase-oper:ocni-ni-intfbase/network-instances/network-instance/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("id", types.YLeaf{"Id", self.Id})

    self.EntityData.YListKeys = []string {"Id"}

    return &(self.EntityData)
}

// OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface_State
// Operational state parameters relating to the
// associated interface
type OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A unique identifier for this interface - this is expressed as a free-text
    // string. The type is string.
    Id interface{}

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string.
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // interface{} with range: 0..4294967295.
    Subinterface interface{}

    // The address families on the subinterface which are to be associated with
    // this network instance. When this leaf-list is empty and the network
    // instance requires Layer 3 information the address families for which the
    // network instance is enabled should be imported. If the value of this
    // leaf-list is specified then the association MUST only be made for those
    // address families that are included in the list. The type is slice of
    // string.
    AssociatedAddressFamily []interface{}
}

func (state *OcniNiIntfbase_NetworkInstances_NetworkInstance_Interfaces_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "Cisco-IOS-XR-ocni-intfbase-oper:ocni-ni-intfbase/network-instances/network-instance/interfaces/interface/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("id", types.YLeaf{"Id", state.Id})
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})
    state.EntityData.Leafs.Append("associated-address-family", types.YLeaf{"AssociatedAddressFamily", state.AssociatedAddressFamily})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

