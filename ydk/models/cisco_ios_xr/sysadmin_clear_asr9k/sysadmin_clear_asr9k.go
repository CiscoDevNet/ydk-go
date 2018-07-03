// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all 'clear' commands for Sysadmin.
// 
// Copyright(c) 2012-2016 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_clear_asr9k

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_clear_asr9k"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-clear-ASR9K clear}", reflect.TypeOf(Clear{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-clear-asr9k:clear", reflect.TypeOf(Clear{}))
}

// Clear
type Clear struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Controller Clear_Controller

    
    Plugin Clear_Plugin
}

func (clear *Clear) GetEntityData() *types.CommonEntityData {
    clear.EntityData.YFilter = clear.YFilter
    clear.EntityData.YangName = "clear"
    clear.EntityData.BundleName = "cisco_ios_xr"
    clear.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-clear-asr9k"
    clear.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-clear-asr9k:clear"
    clear.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clear.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clear.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clear.EntityData.Children = types.NewOrderedMap()
    clear.EntityData.Children.Append("controller", types.YChild{"Controller", &clear.Controller})
    clear.EntityData.Children.Append("plugin", types.YChild{"Plugin", &clear.Plugin})
    clear.EntityData.Leafs = types.NewOrderedMap()

    clear.EntityData.YListKeys = []string {}

    return &(clear.EntityData)
}

// Clear_Controller
type Clear_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Switch Clear_Controller_Switch
}

func (controller *Clear_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "clear"
    controller.EntityData.SegmentPath = "controller"
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("switch", types.YChild{"Switch", &controller.Switch})
    controller.EntityData.Leafs = types.NewOrderedMap()

    controller.EntityData.YListKeys = []string {}

    return &(controller.EntityData)
}

// Clear_Controller_Switch
type Clear_Controller_Switch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control Ethernet switch operational data.
    Oper Clear_Controller_Switch_Oper
}

func (self *Clear_Controller_Switch) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "switch"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "controller"
    self.EntityData.SegmentPath = "switch"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("oper", types.YChild{"Oper", &self.Oper})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Clear_Controller_Switch_Oper
// Control Ethernet switch operational data.
type Clear_Controller_Switch_Oper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Fdb Clear_Controller_Switch_Oper_Fdb

    
    Statistics Clear_Controller_Switch_Oper_Statistics
}

func (oper *Clear_Controller_Switch_Oper) GetEntityData() *types.CommonEntityData {
    oper.EntityData.YFilter = oper.YFilter
    oper.EntityData.YangName = "oper"
    oper.EntityData.BundleName = "cisco_ios_xr"
    oper.EntityData.ParentYangName = "switch"
    oper.EntityData.SegmentPath = "oper"
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = types.NewOrderedMap()
    oper.EntityData.Children.Append("fdb", types.YChild{"Fdb", &oper.Fdb})
    oper.EntityData.Children.Append("statistics", types.YChild{"Statistics", &oper.Statistics})
    oper.EntityData.Leafs = types.NewOrderedMap()

    oper.EntityData.YListKeys = []string {}

    return &(oper.EntityData)
}

// Clear_Controller_Switch_Oper_Fdb
type Clear_Controller_Switch_Oper_Fdb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Clear_Controller_Switch_Oper_Fdb_Location.
    Location []*Clear_Controller_Switch_Oper_Fdb_Location
}

func (fdb *Clear_Controller_Switch_Oper_Fdb) GetEntityData() *types.CommonEntityData {
    fdb.EntityData.YFilter = fdb.YFilter
    fdb.EntityData.YangName = "fdb"
    fdb.EntityData.BundleName = "cisco_ios_xr"
    fdb.EntityData.ParentYangName = "oper"
    fdb.EntityData.SegmentPath = "fdb"
    fdb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdb.EntityData.Children = types.NewOrderedMap()
    fdb.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range fdb.Location {
        fdb.EntityData.Children.Append(types.GetSegmentPath(fdb.Location[i]), types.YChild{"Location", fdb.Location[i]})
    }
    fdb.EntityData.Leafs = types.NewOrderedMap()

    fdb.EntityData.YListKeys = []string {}

    return &(fdb.EntityData)
}

// Clear_Controller_Switch_Oper_Fdb_Location
type Clear_Controller_Switch_Oper_Fdb_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}
}

func (location *Clear_Controller_Switch_Oper_Fdb_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "fdb"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Rack, "rack") + types.AddKeyToken(location.Card, "card") + types.AddKeyToken(location.SwitchId, "switch-id")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", location.Rack})
    location.EntityData.Leafs.Append("card", types.YLeaf{"Card", location.Card})
    location.EntityData.Leafs.Append("switch-id", types.YLeaf{"SwitchId", location.SwitchId})

    location.EntityData.YListKeys = []string {"Rack", "Card", "SwitchId"}

    return &(location.EntityData)
}

// Clear_Controller_Switch_Oper_Statistics
type Clear_Controller_Switch_Oper_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Clear_Controller_Switch_Oper_Statistics_Location.
    Location []*Clear_Controller_Switch_Oper_Statistics_Location
}

func (statistics *Clear_Controller_Switch_Oper_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "oper"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range statistics.Location {
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Location[i]), types.YChild{"Location", statistics.Location[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Clear_Controller_Switch_Oper_Statistics_Location
type Clear_Controller_Switch_Oper_Statistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}
}

func (location *Clear_Controller_Switch_Oper_Statistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "statistics"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Rack, "rack") + types.AddKeyToken(location.Card, "card") + types.AddKeyToken(location.SwitchId, "switch-id")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", location.Rack})
    location.EntityData.Leafs.Append("card", types.YLeaf{"Card", location.Card})
    location.EntityData.Leafs.Append("switch-id", types.YLeaf{"SwitchId", location.SwitchId})

    location.EntityData.YListKeys = []string {"Rack", "Card", "SwitchId"}

    return &(location.EntityData)
}

// Clear_Plugin
type Clear_Plugin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Slot Clear_Plugin_Slot
}

func (plugin *Clear_Plugin) GetEntityData() *types.CommonEntityData {
    plugin.EntityData.YFilter = plugin.YFilter
    plugin.EntityData.YangName = "plugin"
    plugin.EntityData.BundleName = "cisco_ios_xr"
    plugin.EntityData.ParentYangName = "clear"
    plugin.EntityData.SegmentPath = "plugin"
    plugin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    plugin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    plugin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    plugin.EntityData.Children = types.NewOrderedMap()
    plugin.EntityData.Children.Append("slot", types.YChild{"Slot", &plugin.Slot})
    plugin.EntityData.Leafs = types.NewOrderedMap()

    plugin.EntityData.YListKeys = []string {}

    return &(plugin.EntityData)
}

// Clear_Plugin_Slot
type Clear_Plugin_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Clear_Plugin_Slot_Location.
    Location []*Clear_Plugin_Slot_Location
}

func (slot *Clear_Plugin_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "plugin"
    slot.EntityData.SegmentPath = "slot"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range slot.Location {
        slot.EntityData.Children.Append(types.GetSegmentPath(slot.Location[i]), types.YChild{"Location", slot.Location[i]})
    }
    slot.EntityData.Leafs = types.NewOrderedMap()

    slot.EntityData.YListKeys = []string {}

    return &(slot.EntityData)
}

// Clear_Plugin_Slot_Location
type Clear_Plugin_Slot_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}
}

func (location *Clear_Plugin_Slot_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "slot"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

