// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all 'clear' commands for Sysadmin.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_clear_ncs5500

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_clear_ncs5500"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-clear-NCS5500 clear}", reflect.TypeOf(Clear{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-clear-ncs5500:clear", reflect.TypeOf(Clear{}))
}

// Clear
type Clear struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Controller Clear_Controller

    
    ClearAsicErrorsGrp Clear_ClearAsicErrorsGrp
}

func (clear *Clear) GetEntityData() *types.CommonEntityData {
    clear.EntityData.YFilter = clear.YFilter
    clear.EntityData.YangName = "clear"
    clear.EntityData.BundleName = "cisco_ios_xr"
    clear.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-clear-ncs5500"
    clear.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear"
    clear.EntityData.AbsolutePath = clear.EntityData.SegmentPath
    clear.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clear.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clear.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clear.EntityData.Children = types.NewOrderedMap()
    clear.EntityData.Children.Append("controller", types.YChild{"Controller", &clear.Controller})
    clear.EntityData.Children.Append("clear-asic-errors-grp", types.YChild{"ClearAsicErrorsGrp", &clear.ClearAsicErrorsGrp})
    clear.EntityData.Leafs = types.NewOrderedMap()

    clear.EntityData.YListKeys = []string {}

    return &(clear.EntityData)
}

// Clear_Controller
type Clear_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Switch Clear_Controller_Switch

    // Fabric resource commands.
    Fabric Clear_Controller_Fabric
}

func (controller *Clear_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "clear"
    controller.EntityData.SegmentPath = "controller"
    controller.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/" + controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("switch", types.YChild{"Switch", &controller.Switch})
    controller.EntityData.Children.Append("fabric", types.YChild{"Fabric", &controller.Fabric})
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
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/" + self.EntityData.SegmentPath
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
    oper.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/switch/" + oper.EntityData.SegmentPath
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
    fdb.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/switch/oper/" + fdb.EntityData.SegmentPath
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
    YListKey string

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
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/switch/oper/fdb/" + location.EntityData.SegmentPath
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
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/switch/oper/" + statistics.EntityData.SegmentPath
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
    YListKey string

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
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/switch/oper/statistics/" + location.EntityData.SegmentPath
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

// Clear_Controller_Fabric
// Fabric resource commands
type Clear_Controller_Fabric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Counter Clear_Controller_Fabric_Counter

    
    ClearStatistics Clear_Controller_Fabric_ClearStatistics
}

func (fabric *Clear_Controller_Fabric) GetEntityData() *types.CommonEntityData {
    fabric.EntityData.YFilter = fabric.YFilter
    fabric.EntityData.YangName = "fabric"
    fabric.EntityData.BundleName = "cisco_ios_xr"
    fabric.EntityData.ParentYangName = "controller"
    fabric.EntityData.SegmentPath = "fabric"
    fabric.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/" + fabric.EntityData.SegmentPath
    fabric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabric.EntityData.Children = types.NewOrderedMap()
    fabric.EntityData.Children.Append("counter", types.YChild{"Counter", &fabric.Counter})
    fabric.EntityData.Children.Append("clear_statistics", types.YChild{"ClearStatistics", &fabric.ClearStatistics})
    fabric.EntityData.Leafs = types.NewOrderedMap()

    fabric.EntityData.YListKeys = []string {}

    return &(fabric.EntityData)
}

// Clear_Controller_Fabric_Counter
type Clear_Controller_Fabric_Counter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Clear_Controller_Fabric_Counter_Plane.
    Plane []*Clear_Controller_Fabric_Counter_Plane
}

func (counter *Clear_Controller_Fabric_Counter) GetEntityData() *types.CommonEntityData {
    counter.EntityData.YFilter = counter.YFilter
    counter.EntityData.YangName = "counter"
    counter.EntityData.BundleName = "cisco_ios_xr"
    counter.EntityData.ParentYangName = "fabric"
    counter.EntityData.SegmentPath = "counter"
    counter.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/fabric/" + counter.EntityData.SegmentPath
    counter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counter.EntityData.Children = types.NewOrderedMap()
    counter.EntityData.Children.Append("plane", types.YChild{"Plane", nil})
    for i := range counter.Plane {
        counter.EntityData.Children.Append(types.GetSegmentPath(counter.Plane[i]), types.YChild{"Plane", counter.Plane[i]})
    }
    counter.EntityData.Leafs = types.NewOrderedMap()

    counter.EntityData.YListKeys = []string {}

    return &(counter.EntityData)
}

// Clear_Controller_Fabric_Counter_Plane
type Clear_Controller_Fabric_Counter_Plane struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern: all|[0-5].
    Planeid interface{}
}

func (plane *Clear_Controller_Fabric_Counter_Plane) GetEntityData() *types.CommonEntityData {
    plane.EntityData.YFilter = plane.YFilter
    plane.EntityData.YangName = "plane"
    plane.EntityData.BundleName = "cisco_ios_xr"
    plane.EntityData.ParentYangName = "counter"
    plane.EntityData.SegmentPath = "plane" + types.AddKeyToken(plane.Planeid, "planeid")
    plane.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/fabric/counter/" + plane.EntityData.SegmentPath
    plane.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    plane.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    plane.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    plane.EntityData.Children = types.NewOrderedMap()
    plane.EntityData.Leafs = types.NewOrderedMap()
    plane.EntityData.Leafs.Append("planeid", types.YLeaf{"Planeid", plane.Planeid})

    plane.EntityData.YListKeys = []string {"Planeid"}

    return &(plane.EntityData)
}

// Clear_Controller_Fabric_ClearStatistics
type Clear_Controller_Fabric_ClearStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Clear_Controller_Fabric_ClearStatistics_Plane.
    Plane []*Clear_Controller_Fabric_ClearStatistics_Plane
}

func (clearStatistics *Clear_Controller_Fabric_ClearStatistics) GetEntityData() *types.CommonEntityData {
    clearStatistics.EntityData.YFilter = clearStatistics.YFilter
    clearStatistics.EntityData.YangName = "clear_statistics"
    clearStatistics.EntityData.BundleName = "cisco_ios_xr"
    clearStatistics.EntityData.ParentYangName = "fabric"
    clearStatistics.EntityData.SegmentPath = "clear_statistics"
    clearStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/fabric/" + clearStatistics.EntityData.SegmentPath
    clearStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearStatistics.EntityData.Children = types.NewOrderedMap()
    clearStatistics.EntityData.Children.Append("plane", types.YChild{"Plane", nil})
    for i := range clearStatistics.Plane {
        clearStatistics.EntityData.Children.Append(types.GetSegmentPath(clearStatistics.Plane[i]), types.YChild{"Plane", clearStatistics.Plane[i]})
    }
    clearStatistics.EntityData.Leafs = types.NewOrderedMap()

    clearStatistics.EntityData.YListKeys = []string {}

    return &(clearStatistics.EntityData)
}

// Clear_Controller_Fabric_ClearStatistics_Plane
type Clear_Controller_Fabric_ClearStatistics_Plane struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern: all|[0-5].
    Planeid interface{}
}

func (plane *Clear_Controller_Fabric_ClearStatistics_Plane) GetEntityData() *types.CommonEntityData {
    plane.EntityData.YFilter = plane.YFilter
    plane.EntityData.YangName = "plane"
    plane.EntityData.BundleName = "cisco_ios_xr"
    plane.EntityData.ParentYangName = "clear_statistics"
    plane.EntityData.SegmentPath = "plane" + types.AddKeyToken(plane.Planeid, "planeid")
    plane.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/controller/fabric/clear_statistics/" + plane.EntityData.SegmentPath
    plane.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    plane.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    plane.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    plane.EntityData.Children = types.NewOrderedMap()
    plane.EntityData.Leafs = types.NewOrderedMap()
    plane.EntityData.Leafs.Append("planeid", types.YLeaf{"Planeid", plane.Planeid})

    plane.EntityData.YListKeys = []string {"Planeid"}

    return &(plane.EntityData)
}

// Clear_ClearAsicErrorsGrp
type Clear_ClearAsicErrorsGrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Clear_ClearAsicErrorsGrp_ClearDevice.
    ClearDevice []*Clear_ClearAsicErrorsGrp_ClearDevice
}

func (clearAsicErrorsGrp *Clear_ClearAsicErrorsGrp) GetEntityData() *types.CommonEntityData {
    clearAsicErrorsGrp.EntityData.YFilter = clearAsicErrorsGrp.YFilter
    clearAsicErrorsGrp.EntityData.YangName = "clear-asic-errors-grp"
    clearAsicErrorsGrp.EntityData.BundleName = "cisco_ios_xr"
    clearAsicErrorsGrp.EntityData.ParentYangName = "clear"
    clearAsicErrorsGrp.EntityData.SegmentPath = "clear-asic-errors-grp"
    clearAsicErrorsGrp.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/" + clearAsicErrorsGrp.EntityData.SegmentPath
    clearAsicErrorsGrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearAsicErrorsGrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearAsicErrorsGrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearAsicErrorsGrp.EntityData.Children = types.NewOrderedMap()
    clearAsicErrorsGrp.EntityData.Children.Append("clear-device", types.YChild{"ClearDevice", nil})
    for i := range clearAsicErrorsGrp.ClearDevice {
        clearAsicErrorsGrp.EntityData.Children.Append(types.GetSegmentPath(clearAsicErrorsGrp.ClearDevice[i]), types.YChild{"ClearDevice", clearAsicErrorsGrp.ClearDevice[i]})
    }
    clearAsicErrorsGrp.EntityData.Leafs = types.NewOrderedMap()

    clearAsicErrorsGrp.EntityData.YListKeys = []string {}

    return &(clearAsicErrorsGrp.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice
type Clear_ClearAsicErrorsGrp_ClearDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    DeviceName interface{}

    // The type is slice of Clear_ClearAsicErrorsGrp_ClearDevice_Instance.
    Instance []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance

    
    AllInstances Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances
}

func (clearDevice *Clear_ClearAsicErrorsGrp_ClearDevice) GetEntityData() *types.CommonEntityData {
    clearDevice.EntityData.YFilter = clearDevice.YFilter
    clearDevice.EntityData.YangName = "clear-device"
    clearDevice.EntityData.BundleName = "cisco_ios_xr"
    clearDevice.EntityData.ParentYangName = "clear-asic-errors-grp"
    clearDevice.EntityData.SegmentPath = "clear-device" + types.AddKeyToken(clearDevice.DeviceName, "device-name")
    clearDevice.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/" + clearDevice.EntityData.SegmentPath
    clearDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearDevice.EntityData.Children = types.NewOrderedMap()
    clearDevice.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range clearDevice.Instance {
        clearDevice.EntityData.Children.Append(types.GetSegmentPath(clearDevice.Instance[i]), types.YChild{"Instance", clearDevice.Instance[i]})
    }
    clearDevice.EntityData.Children.Append("all-instances", types.YChild{"AllInstances", &clearDevice.AllInstances})
    clearDevice.EntityData.Leafs = types.NewOrderedMap()
    clearDevice.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", clearDevice.DeviceName})

    clearDevice.EntityData.YListKeys = []string {"DeviceName"}

    return &(clearDevice.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    InstanceNum interface{}

    
    Sbe Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe

    
    Mbe Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe

    
    Parity Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity

    
    Generic Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic

    
    Crc Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc

    
    Reset Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset

    
    Barrier Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier

    
    Unexpected Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected

    
    Link Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link

    
    OorThresh Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh

    
    Bp Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp

    
    Io Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io

    
    Ucode Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode

    
    Config Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config

    
    Indirect Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect

    
    Nonerr Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr

    
    Summary Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary

    
    All Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All
}

func (instance *Clear_ClearAsicErrorsGrp_ClearDevice_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "clear-device"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.InstanceNum, "instance-num")
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("sbe", types.YChild{"Sbe", &instance.Sbe})
    instance.EntityData.Children.Append("mbe", types.YChild{"Mbe", &instance.Mbe})
    instance.EntityData.Children.Append("parity", types.YChild{"Parity", &instance.Parity})
    instance.EntityData.Children.Append("generic", types.YChild{"Generic", &instance.Generic})
    instance.EntityData.Children.Append("crc", types.YChild{"Crc", &instance.Crc})
    instance.EntityData.Children.Append("reset", types.YChild{"Reset", &instance.Reset})
    instance.EntityData.Children.Append("barrier", types.YChild{"Barrier", &instance.Barrier})
    instance.EntityData.Children.Append("unexpected", types.YChild{"Unexpected", &instance.Unexpected})
    instance.EntityData.Children.Append("link", types.YChild{"Link", &instance.Link})
    instance.EntityData.Children.Append("oor-thresh", types.YChild{"OorThresh", &instance.OorThresh})
    instance.EntityData.Children.Append("bp", types.YChild{"Bp", &instance.Bp})
    instance.EntityData.Children.Append("io", types.YChild{"Io", &instance.Io})
    instance.EntityData.Children.Append("ucode", types.YChild{"Ucode", &instance.Ucode})
    instance.EntityData.Children.Append("config", types.YChild{"Config", &instance.Config})
    instance.EntityData.Children.Append("indirect", types.YChild{"Indirect", &instance.Indirect})
    instance.EntityData.Children.Append("nonerr", types.YChild{"Nonerr", &instance.Nonerr})
    instance.EntityData.Children.Append("summary", types.YChild{"Summary", &instance.Summary})
    instance.EntityData.Children.Append("all", types.YChild{"All", &instance.All})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-num", types.YLeaf{"InstanceNum", instance.InstanceNum})

    instance.EntityData.YListKeys = []string {"InstanceNum"}

    return &(instance.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe_Location
}

func (sbe *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe) GetEntityData() *types.CommonEntityData {
    sbe.EntityData.YFilter = sbe.YFilter
    sbe.EntityData.YangName = "sbe"
    sbe.EntityData.BundleName = "cisco_ios_xr"
    sbe.EntityData.ParentYangName = "instance"
    sbe.EntityData.SegmentPath = "sbe"
    sbe.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + sbe.EntityData.SegmentPath
    sbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sbe.EntityData.Children = types.NewOrderedMap()
    sbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range sbe.Location {
        sbe.EntityData.Children.Append(types.GetSegmentPath(sbe.Location[i]), types.YChild{"Location", sbe.Location[i]})
    }
    sbe.EntityData.Leafs = types.NewOrderedMap()

    sbe.EntityData.YListKeys = []string {}

    return &(sbe.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Sbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "sbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/sbe/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe_Location
}

func (mbe *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe) GetEntityData() *types.CommonEntityData {
    mbe.EntityData.YFilter = mbe.YFilter
    mbe.EntityData.YangName = "mbe"
    mbe.EntityData.BundleName = "cisco_ios_xr"
    mbe.EntityData.ParentYangName = "instance"
    mbe.EntityData.SegmentPath = "mbe"
    mbe.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + mbe.EntityData.SegmentPath
    mbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mbe.EntityData.Children = types.NewOrderedMap()
    mbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range mbe.Location {
        mbe.EntityData.Children.Append(types.GetSegmentPath(mbe.Location[i]), types.YChild{"Location", mbe.Location[i]})
    }
    mbe.EntityData.Leafs = types.NewOrderedMap()

    mbe.EntityData.YListKeys = []string {}

    return &(mbe.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Mbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "mbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/mbe/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity_Location
}

func (parity *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity) GetEntityData() *types.CommonEntityData {
    parity.EntityData.YFilter = parity.YFilter
    parity.EntityData.YangName = "parity"
    parity.EntityData.BundleName = "cisco_ios_xr"
    parity.EntityData.ParentYangName = "instance"
    parity.EntityData.SegmentPath = "parity"
    parity.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + parity.EntityData.SegmentPath
    parity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parity.EntityData.Children = types.NewOrderedMap()
    parity.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range parity.Location {
        parity.EntityData.Children.Append(types.GetSegmentPath(parity.Location[i]), types.YChild{"Location", parity.Location[i]})
    }
    parity.EntityData.Leafs = types.NewOrderedMap()

    parity.EntityData.YListKeys = []string {}

    return &(parity.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Parity_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "parity"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/parity/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic_Location
}

func (generic *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic) GetEntityData() *types.CommonEntityData {
    generic.EntityData.YFilter = generic.YFilter
    generic.EntityData.YangName = "generic"
    generic.EntityData.BundleName = "cisco_ios_xr"
    generic.EntityData.ParentYangName = "instance"
    generic.EntityData.SegmentPath = "generic"
    generic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + generic.EntityData.SegmentPath
    generic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generic.EntityData.Children = types.NewOrderedMap()
    generic.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range generic.Location {
        generic.EntityData.Children.Append(types.GetSegmentPath(generic.Location[i]), types.YChild{"Location", generic.Location[i]})
    }
    generic.EntityData.Leafs = types.NewOrderedMap()

    generic.EntityData.YListKeys = []string {}

    return &(generic.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Generic_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "generic"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/generic/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc_Location
}

func (crc *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc) GetEntityData() *types.CommonEntityData {
    crc.EntityData.YFilter = crc.YFilter
    crc.EntityData.YangName = "crc"
    crc.EntityData.BundleName = "cisco_ios_xr"
    crc.EntityData.ParentYangName = "instance"
    crc.EntityData.SegmentPath = "crc"
    crc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + crc.EntityData.SegmentPath
    crc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crc.EntityData.Children = types.NewOrderedMap()
    crc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range crc.Location {
        crc.EntityData.Children.Append(types.GetSegmentPath(crc.Location[i]), types.YChild{"Location", crc.Location[i]})
    }
    crc.EntityData.Leafs = types.NewOrderedMap()

    crc.EntityData.YListKeys = []string {}

    return &(crc.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Crc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "crc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/crc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset_Location
}

func (reset *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset) GetEntityData() *types.CommonEntityData {
    reset.EntityData.YFilter = reset.YFilter
    reset.EntityData.YangName = "reset"
    reset.EntityData.BundleName = "cisco_ios_xr"
    reset.EntityData.ParentYangName = "instance"
    reset.EntityData.SegmentPath = "reset"
    reset.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + reset.EntityData.SegmentPath
    reset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reset.EntityData.Children = types.NewOrderedMap()
    reset.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range reset.Location {
        reset.EntityData.Children.Append(types.GetSegmentPath(reset.Location[i]), types.YChild{"Location", reset.Location[i]})
    }
    reset.EntityData.Leafs = types.NewOrderedMap()

    reset.EntityData.YListKeys = []string {}

    return &(reset.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Reset_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "reset"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/reset/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier_Location
}

func (barrier *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier) GetEntityData() *types.CommonEntityData {
    barrier.EntityData.YFilter = barrier.YFilter
    barrier.EntityData.YangName = "barrier"
    barrier.EntityData.BundleName = "cisco_ios_xr"
    barrier.EntityData.ParentYangName = "instance"
    barrier.EntityData.SegmentPath = "barrier"
    barrier.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + barrier.EntityData.SegmentPath
    barrier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    barrier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    barrier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    barrier.EntityData.Children = types.NewOrderedMap()
    barrier.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range barrier.Location {
        barrier.EntityData.Children.Append(types.GetSegmentPath(barrier.Location[i]), types.YChild{"Location", barrier.Location[i]})
    }
    barrier.EntityData.Leafs = types.NewOrderedMap()

    barrier.EntityData.YListKeys = []string {}

    return &(barrier.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Barrier_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "barrier"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/barrier/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected_Location
}

func (unexpected *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected) GetEntityData() *types.CommonEntityData {
    unexpected.EntityData.YFilter = unexpected.YFilter
    unexpected.EntityData.YangName = "unexpected"
    unexpected.EntityData.BundleName = "cisco_ios_xr"
    unexpected.EntityData.ParentYangName = "instance"
    unexpected.EntityData.SegmentPath = "unexpected"
    unexpected.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + unexpected.EntityData.SegmentPath
    unexpected.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unexpected.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unexpected.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unexpected.EntityData.Children = types.NewOrderedMap()
    unexpected.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range unexpected.Location {
        unexpected.EntityData.Children.Append(types.GetSegmentPath(unexpected.Location[i]), types.YChild{"Location", unexpected.Location[i]})
    }
    unexpected.EntityData.Leafs = types.NewOrderedMap()

    unexpected.EntityData.YListKeys = []string {}

    return &(unexpected.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Unexpected_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "unexpected"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/unexpected/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link_Location
}

func (link *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "cisco_ios_xr"
    link.EntityData.ParentYangName = "instance"
    link.EntityData.SegmentPath = "link"
    link.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + link.EntityData.SegmentPath
    link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    link.EntityData.Children = types.NewOrderedMap()
    link.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range link.Location {
        link.EntityData.Children.Append(types.GetSegmentPath(link.Location[i]), types.YChild{"Location", link.Location[i]})
    }
    link.EntityData.Leafs = types.NewOrderedMap()

    link.EntityData.YListKeys = []string {}

    return &(link.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Link_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "link"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/link/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh_Location
}

func (oorThresh *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh) GetEntityData() *types.CommonEntityData {
    oorThresh.EntityData.YFilter = oorThresh.YFilter
    oorThresh.EntityData.YangName = "oor-thresh"
    oorThresh.EntityData.BundleName = "cisco_ios_xr"
    oorThresh.EntityData.ParentYangName = "instance"
    oorThresh.EntityData.SegmentPath = "oor-thresh"
    oorThresh.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + oorThresh.EntityData.SegmentPath
    oorThresh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorThresh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorThresh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorThresh.EntityData.Children = types.NewOrderedMap()
    oorThresh.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range oorThresh.Location {
        oorThresh.EntityData.Children.Append(types.GetSegmentPath(oorThresh.Location[i]), types.YChild{"Location", oorThresh.Location[i]})
    }
    oorThresh.EntityData.Leafs = types.NewOrderedMap()

    oorThresh.EntityData.YListKeys = []string {}

    return &(oorThresh.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_OorThresh_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "oor-thresh"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/oor-thresh/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp_Location
}

func (bp *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp) GetEntityData() *types.CommonEntityData {
    bp.EntityData.YFilter = bp.YFilter
    bp.EntityData.YangName = "bp"
    bp.EntityData.BundleName = "cisco_ios_xr"
    bp.EntityData.ParentYangName = "instance"
    bp.EntityData.SegmentPath = "bp"
    bp.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + bp.EntityData.SegmentPath
    bp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bp.EntityData.Children = types.NewOrderedMap()
    bp.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range bp.Location {
        bp.EntityData.Children.Append(types.GetSegmentPath(bp.Location[i]), types.YChild{"Location", bp.Location[i]})
    }
    bp.EntityData.Leafs = types.NewOrderedMap()

    bp.EntityData.YListKeys = []string {}

    return &(bp.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Bp_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "bp"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/bp/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io_Location
}

func (io *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io) GetEntityData() *types.CommonEntityData {
    io.EntityData.YFilter = io.YFilter
    io.EntityData.YangName = "io"
    io.EntityData.BundleName = "cisco_ios_xr"
    io.EntityData.ParentYangName = "instance"
    io.EntityData.SegmentPath = "io"
    io.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + io.EntityData.SegmentPath
    io.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    io.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    io.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    io.EntityData.Children = types.NewOrderedMap()
    io.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range io.Location {
        io.EntityData.Children.Append(types.GetSegmentPath(io.Location[i]), types.YChild{"Location", io.Location[i]})
    }
    io.EntityData.Leafs = types.NewOrderedMap()

    io.EntityData.YListKeys = []string {}

    return &(io.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Io_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "io"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/io/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode_Location
}

func (ucode *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode) GetEntityData() *types.CommonEntityData {
    ucode.EntityData.YFilter = ucode.YFilter
    ucode.EntityData.YangName = "ucode"
    ucode.EntityData.BundleName = "cisco_ios_xr"
    ucode.EntityData.ParentYangName = "instance"
    ucode.EntityData.SegmentPath = "ucode"
    ucode.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + ucode.EntityData.SegmentPath
    ucode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucode.EntityData.Children = types.NewOrderedMap()
    ucode.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range ucode.Location {
        ucode.EntityData.Children.Append(types.GetSegmentPath(ucode.Location[i]), types.YChild{"Location", ucode.Location[i]})
    }
    ucode.EntityData.Leafs = types.NewOrderedMap()

    ucode.EntityData.YListKeys = []string {}

    return &(ucode.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Ucode_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "ucode"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/ucode/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config_Location
}

func (config *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "instance"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range config.Location {
        config.EntityData.Children.Append(types.GetSegmentPath(config.Location[i]), types.YChild{"Location", config.Location[i]})
    }
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Config_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "config"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/config/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect_Location
}

func (indirect *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect) GetEntityData() *types.CommonEntityData {
    indirect.EntityData.YFilter = indirect.YFilter
    indirect.EntityData.YangName = "indirect"
    indirect.EntityData.BundleName = "cisco_ios_xr"
    indirect.EntityData.ParentYangName = "instance"
    indirect.EntityData.SegmentPath = "indirect"
    indirect.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + indirect.EntityData.SegmentPath
    indirect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indirect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indirect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indirect.EntityData.Children = types.NewOrderedMap()
    indirect.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range indirect.Location {
        indirect.EntityData.Children.Append(types.GetSegmentPath(indirect.Location[i]), types.YChild{"Location", indirect.Location[i]})
    }
    indirect.EntityData.Leafs = types.NewOrderedMap()

    indirect.EntityData.YListKeys = []string {}

    return &(indirect.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Indirect_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "indirect"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/indirect/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr_Location
}

func (nonerr *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr) GetEntityData() *types.CommonEntityData {
    nonerr.EntityData.YFilter = nonerr.YFilter
    nonerr.EntityData.YangName = "nonerr"
    nonerr.EntityData.BundleName = "cisco_ios_xr"
    nonerr.EntityData.ParentYangName = "instance"
    nonerr.EntityData.SegmentPath = "nonerr"
    nonerr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + nonerr.EntityData.SegmentPath
    nonerr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonerr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonerr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonerr.EntityData.Children = types.NewOrderedMap()
    nonerr.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range nonerr.Location {
        nonerr.EntityData.Children.Append(types.GetSegmentPath(nonerr.Location[i]), types.YChild{"Location", nonerr.Location[i]})
    }
    nonerr.EntityData.Leafs = types.NewOrderedMap()

    nonerr.EntityData.YListKeys = []string {}

    return &(nonerr.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Nonerr_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "nonerr"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/nonerr/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary_Location
}

func (summary *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "instance"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range summary.Location {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.Location[i]), types.YChild{"Location", summary.Location[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_Summary_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/summary/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_Location
}

func (all *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "instance"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("history", types.YChild{"History", &all.History})
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History_Location
}

func (history *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "all"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/all/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range history.Location {
        history.EntityData.Children.Append(types.GetSegmentPath(history.Location[i]), types.YChild{"Location", history.Location[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_History_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "history"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/all/history/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_Instance_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/instance/all/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Sbe Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe

    
    Mbe Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe

    
    Parity Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity

    
    Generic Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic

    
    Crc Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc

    
    Reset Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset

    
    Barrier Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier

    
    Unexpected Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected

    
    Link Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link

    
    OorThresh Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh

    
    Bp Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp

    
    Io Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io

    
    Ucode Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode

    
    Config Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config

    
    Indirect Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect

    
    Nonerr Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr

    
    Summary Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary

    
    All Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All
}

func (allInstances *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances) GetEntityData() *types.CommonEntityData {
    allInstances.EntityData.YFilter = allInstances.YFilter
    allInstances.EntityData.YangName = "all-instances"
    allInstances.EntityData.BundleName = "cisco_ios_xr"
    allInstances.EntityData.ParentYangName = "clear-device"
    allInstances.EntityData.SegmentPath = "all-instances"
    allInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/" + allInstances.EntityData.SegmentPath
    allInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allInstances.EntityData.Children = types.NewOrderedMap()
    allInstances.EntityData.Children.Append("sbe", types.YChild{"Sbe", &allInstances.Sbe})
    allInstances.EntityData.Children.Append("mbe", types.YChild{"Mbe", &allInstances.Mbe})
    allInstances.EntityData.Children.Append("parity", types.YChild{"Parity", &allInstances.Parity})
    allInstances.EntityData.Children.Append("generic", types.YChild{"Generic", &allInstances.Generic})
    allInstances.EntityData.Children.Append("crc", types.YChild{"Crc", &allInstances.Crc})
    allInstances.EntityData.Children.Append("reset", types.YChild{"Reset", &allInstances.Reset})
    allInstances.EntityData.Children.Append("barrier", types.YChild{"Barrier", &allInstances.Barrier})
    allInstances.EntityData.Children.Append("unexpected", types.YChild{"Unexpected", &allInstances.Unexpected})
    allInstances.EntityData.Children.Append("link", types.YChild{"Link", &allInstances.Link})
    allInstances.EntityData.Children.Append("oor-thresh", types.YChild{"OorThresh", &allInstances.OorThresh})
    allInstances.EntityData.Children.Append("bp", types.YChild{"Bp", &allInstances.Bp})
    allInstances.EntityData.Children.Append("io", types.YChild{"Io", &allInstances.Io})
    allInstances.EntityData.Children.Append("ucode", types.YChild{"Ucode", &allInstances.Ucode})
    allInstances.EntityData.Children.Append("config", types.YChild{"Config", &allInstances.Config})
    allInstances.EntityData.Children.Append("indirect", types.YChild{"Indirect", &allInstances.Indirect})
    allInstances.EntityData.Children.Append("nonerr", types.YChild{"Nonerr", &allInstances.Nonerr})
    allInstances.EntityData.Children.Append("summary", types.YChild{"Summary", &allInstances.Summary})
    allInstances.EntityData.Children.Append("all", types.YChild{"All", &allInstances.All})
    allInstances.EntityData.Leafs = types.NewOrderedMap()

    allInstances.EntityData.YListKeys = []string {}

    return &(allInstances.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe_Location
}

func (sbe *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe) GetEntityData() *types.CommonEntityData {
    sbe.EntityData.YFilter = sbe.YFilter
    sbe.EntityData.YangName = "sbe"
    sbe.EntityData.BundleName = "cisco_ios_xr"
    sbe.EntityData.ParentYangName = "all-instances"
    sbe.EntityData.SegmentPath = "sbe"
    sbe.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + sbe.EntityData.SegmentPath
    sbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sbe.EntityData.Children = types.NewOrderedMap()
    sbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range sbe.Location {
        sbe.EntityData.Children.Append(types.GetSegmentPath(sbe.Location[i]), types.YChild{"Location", sbe.Location[i]})
    }
    sbe.EntityData.Leafs = types.NewOrderedMap()

    sbe.EntityData.YListKeys = []string {}

    return &(sbe.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Sbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "sbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/sbe/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe_Location
}

func (mbe *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe) GetEntityData() *types.CommonEntityData {
    mbe.EntityData.YFilter = mbe.YFilter
    mbe.EntityData.YangName = "mbe"
    mbe.EntityData.BundleName = "cisco_ios_xr"
    mbe.EntityData.ParentYangName = "all-instances"
    mbe.EntityData.SegmentPath = "mbe"
    mbe.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + mbe.EntityData.SegmentPath
    mbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mbe.EntityData.Children = types.NewOrderedMap()
    mbe.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range mbe.Location {
        mbe.EntityData.Children.Append(types.GetSegmentPath(mbe.Location[i]), types.YChild{"Location", mbe.Location[i]})
    }
    mbe.EntityData.Leafs = types.NewOrderedMap()

    mbe.EntityData.YListKeys = []string {}

    return &(mbe.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Mbe_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "mbe"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/mbe/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity_Location
}

func (parity *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity) GetEntityData() *types.CommonEntityData {
    parity.EntityData.YFilter = parity.YFilter
    parity.EntityData.YangName = "parity"
    parity.EntityData.BundleName = "cisco_ios_xr"
    parity.EntityData.ParentYangName = "all-instances"
    parity.EntityData.SegmentPath = "parity"
    parity.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + parity.EntityData.SegmentPath
    parity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parity.EntityData.Children = types.NewOrderedMap()
    parity.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range parity.Location {
        parity.EntityData.Children.Append(types.GetSegmentPath(parity.Location[i]), types.YChild{"Location", parity.Location[i]})
    }
    parity.EntityData.Leafs = types.NewOrderedMap()

    parity.EntityData.YListKeys = []string {}

    return &(parity.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Parity_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "parity"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/parity/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic_Location
}

func (generic *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic) GetEntityData() *types.CommonEntityData {
    generic.EntityData.YFilter = generic.YFilter
    generic.EntityData.YangName = "generic"
    generic.EntityData.BundleName = "cisco_ios_xr"
    generic.EntityData.ParentYangName = "all-instances"
    generic.EntityData.SegmentPath = "generic"
    generic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + generic.EntityData.SegmentPath
    generic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generic.EntityData.Children = types.NewOrderedMap()
    generic.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range generic.Location {
        generic.EntityData.Children.Append(types.GetSegmentPath(generic.Location[i]), types.YChild{"Location", generic.Location[i]})
    }
    generic.EntityData.Leafs = types.NewOrderedMap()

    generic.EntityData.YListKeys = []string {}

    return &(generic.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Generic_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "generic"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/generic/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc_Location
}

func (crc *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc) GetEntityData() *types.CommonEntityData {
    crc.EntityData.YFilter = crc.YFilter
    crc.EntityData.YangName = "crc"
    crc.EntityData.BundleName = "cisco_ios_xr"
    crc.EntityData.ParentYangName = "all-instances"
    crc.EntityData.SegmentPath = "crc"
    crc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + crc.EntityData.SegmentPath
    crc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crc.EntityData.Children = types.NewOrderedMap()
    crc.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range crc.Location {
        crc.EntityData.Children.Append(types.GetSegmentPath(crc.Location[i]), types.YChild{"Location", crc.Location[i]})
    }
    crc.EntityData.Leafs = types.NewOrderedMap()

    crc.EntityData.YListKeys = []string {}

    return &(crc.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Crc_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "crc"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/crc/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset_Location
}

func (reset *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset) GetEntityData() *types.CommonEntityData {
    reset.EntityData.YFilter = reset.YFilter
    reset.EntityData.YangName = "reset"
    reset.EntityData.BundleName = "cisco_ios_xr"
    reset.EntityData.ParentYangName = "all-instances"
    reset.EntityData.SegmentPath = "reset"
    reset.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + reset.EntityData.SegmentPath
    reset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reset.EntityData.Children = types.NewOrderedMap()
    reset.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range reset.Location {
        reset.EntityData.Children.Append(types.GetSegmentPath(reset.Location[i]), types.YChild{"Location", reset.Location[i]})
    }
    reset.EntityData.Leafs = types.NewOrderedMap()

    reset.EntityData.YListKeys = []string {}

    return &(reset.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Reset_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "reset"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/reset/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier_Location
}

func (barrier *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier) GetEntityData() *types.CommonEntityData {
    barrier.EntityData.YFilter = barrier.YFilter
    barrier.EntityData.YangName = "barrier"
    barrier.EntityData.BundleName = "cisco_ios_xr"
    barrier.EntityData.ParentYangName = "all-instances"
    barrier.EntityData.SegmentPath = "barrier"
    barrier.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + barrier.EntityData.SegmentPath
    barrier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    barrier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    barrier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    barrier.EntityData.Children = types.NewOrderedMap()
    barrier.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range barrier.Location {
        barrier.EntityData.Children.Append(types.GetSegmentPath(barrier.Location[i]), types.YChild{"Location", barrier.Location[i]})
    }
    barrier.EntityData.Leafs = types.NewOrderedMap()

    barrier.EntityData.YListKeys = []string {}

    return &(barrier.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Barrier_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "barrier"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/barrier/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected_Location
}

func (unexpected *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected) GetEntityData() *types.CommonEntityData {
    unexpected.EntityData.YFilter = unexpected.YFilter
    unexpected.EntityData.YangName = "unexpected"
    unexpected.EntityData.BundleName = "cisco_ios_xr"
    unexpected.EntityData.ParentYangName = "all-instances"
    unexpected.EntityData.SegmentPath = "unexpected"
    unexpected.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + unexpected.EntityData.SegmentPath
    unexpected.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unexpected.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unexpected.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unexpected.EntityData.Children = types.NewOrderedMap()
    unexpected.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range unexpected.Location {
        unexpected.EntityData.Children.Append(types.GetSegmentPath(unexpected.Location[i]), types.YChild{"Location", unexpected.Location[i]})
    }
    unexpected.EntityData.Leafs = types.NewOrderedMap()

    unexpected.EntityData.YListKeys = []string {}

    return &(unexpected.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Unexpected_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "unexpected"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/unexpected/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link_Location
}

func (link *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "cisco_ios_xr"
    link.EntityData.ParentYangName = "all-instances"
    link.EntityData.SegmentPath = "link"
    link.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + link.EntityData.SegmentPath
    link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    link.EntityData.Children = types.NewOrderedMap()
    link.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range link.Location {
        link.EntityData.Children.Append(types.GetSegmentPath(link.Location[i]), types.YChild{"Location", link.Location[i]})
    }
    link.EntityData.Leafs = types.NewOrderedMap()

    link.EntityData.YListKeys = []string {}

    return &(link.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Link_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "link"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/link/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh_Location
}

func (oorThresh *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh) GetEntityData() *types.CommonEntityData {
    oorThresh.EntityData.YFilter = oorThresh.YFilter
    oorThresh.EntityData.YangName = "oor-thresh"
    oorThresh.EntityData.BundleName = "cisco_ios_xr"
    oorThresh.EntityData.ParentYangName = "all-instances"
    oorThresh.EntityData.SegmentPath = "oor-thresh"
    oorThresh.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + oorThresh.EntityData.SegmentPath
    oorThresh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorThresh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorThresh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorThresh.EntityData.Children = types.NewOrderedMap()
    oorThresh.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range oorThresh.Location {
        oorThresh.EntityData.Children.Append(types.GetSegmentPath(oorThresh.Location[i]), types.YChild{"Location", oorThresh.Location[i]})
    }
    oorThresh.EntityData.Leafs = types.NewOrderedMap()

    oorThresh.EntityData.YListKeys = []string {}

    return &(oorThresh.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_OorThresh_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "oor-thresh"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/oor-thresh/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp_Location
}

func (bp *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp) GetEntityData() *types.CommonEntityData {
    bp.EntityData.YFilter = bp.YFilter
    bp.EntityData.YangName = "bp"
    bp.EntityData.BundleName = "cisco_ios_xr"
    bp.EntityData.ParentYangName = "all-instances"
    bp.EntityData.SegmentPath = "bp"
    bp.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + bp.EntityData.SegmentPath
    bp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bp.EntityData.Children = types.NewOrderedMap()
    bp.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range bp.Location {
        bp.EntityData.Children.Append(types.GetSegmentPath(bp.Location[i]), types.YChild{"Location", bp.Location[i]})
    }
    bp.EntityData.Leafs = types.NewOrderedMap()

    bp.EntityData.YListKeys = []string {}

    return &(bp.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Bp_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "bp"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/bp/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io_Location
}

func (io *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io) GetEntityData() *types.CommonEntityData {
    io.EntityData.YFilter = io.YFilter
    io.EntityData.YangName = "io"
    io.EntityData.BundleName = "cisco_ios_xr"
    io.EntityData.ParentYangName = "all-instances"
    io.EntityData.SegmentPath = "io"
    io.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + io.EntityData.SegmentPath
    io.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    io.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    io.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    io.EntityData.Children = types.NewOrderedMap()
    io.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range io.Location {
        io.EntityData.Children.Append(types.GetSegmentPath(io.Location[i]), types.YChild{"Location", io.Location[i]})
    }
    io.EntityData.Leafs = types.NewOrderedMap()

    io.EntityData.YListKeys = []string {}

    return &(io.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Io_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "io"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/io/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode_Location
}

func (ucode *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode) GetEntityData() *types.CommonEntityData {
    ucode.EntityData.YFilter = ucode.YFilter
    ucode.EntityData.YangName = "ucode"
    ucode.EntityData.BundleName = "cisco_ios_xr"
    ucode.EntityData.ParentYangName = "all-instances"
    ucode.EntityData.SegmentPath = "ucode"
    ucode.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + ucode.EntityData.SegmentPath
    ucode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucode.EntityData.Children = types.NewOrderedMap()
    ucode.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range ucode.Location {
        ucode.EntityData.Children.Append(types.GetSegmentPath(ucode.Location[i]), types.YChild{"Location", ucode.Location[i]})
    }
    ucode.EntityData.Leafs = types.NewOrderedMap()

    ucode.EntityData.YListKeys = []string {}

    return &(ucode.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Ucode_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "ucode"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/ucode/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config_Location
}

func (config *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "all-instances"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range config.Location {
        config.EntityData.Children.Append(types.GetSegmentPath(config.Location[i]), types.YChild{"Location", config.Location[i]})
    }
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Config_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "config"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/config/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect_Location
}

func (indirect *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect) GetEntityData() *types.CommonEntityData {
    indirect.EntityData.YFilter = indirect.YFilter
    indirect.EntityData.YangName = "indirect"
    indirect.EntityData.BundleName = "cisco_ios_xr"
    indirect.EntityData.ParentYangName = "all-instances"
    indirect.EntityData.SegmentPath = "indirect"
    indirect.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + indirect.EntityData.SegmentPath
    indirect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indirect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indirect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indirect.EntityData.Children = types.NewOrderedMap()
    indirect.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range indirect.Location {
        indirect.EntityData.Children.Append(types.GetSegmentPath(indirect.Location[i]), types.YChild{"Location", indirect.Location[i]})
    }
    indirect.EntityData.Leafs = types.NewOrderedMap()

    indirect.EntityData.YListKeys = []string {}

    return &(indirect.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Indirect_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "indirect"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/indirect/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr_Location
}

func (nonerr *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr) GetEntityData() *types.CommonEntityData {
    nonerr.EntityData.YFilter = nonerr.YFilter
    nonerr.EntityData.YangName = "nonerr"
    nonerr.EntityData.BundleName = "cisco_ios_xr"
    nonerr.EntityData.ParentYangName = "all-instances"
    nonerr.EntityData.SegmentPath = "nonerr"
    nonerr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + nonerr.EntityData.SegmentPath
    nonerr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonerr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonerr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonerr.EntityData.Children = types.NewOrderedMap()
    nonerr.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range nonerr.Location {
        nonerr.EntityData.Children.Append(types.GetSegmentPath(nonerr.Location[i]), types.YChild{"Location", nonerr.Location[i]})
    }
    nonerr.EntityData.Leafs = types.NewOrderedMap()

    nonerr.EntityData.YListKeys = []string {}

    return &(nonerr.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Nonerr_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "nonerr"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/nonerr/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary_Location
}

func (summary *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "all-instances"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range summary.Location {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.Location[i]), types.YChild{"Location", summary.Location[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_Summary_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/summary/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_Location
}

func (all *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "all-instances"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("history", types.YChild{"History", &all.History})
    all.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range all.Location {
        all.EntityData.Children.Append(types.GetSegmentPath(all.Location[i]), types.YChild{"Location", all.Location[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History_Location.
    Location []*Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History_Location
}

func (history *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "all"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/all/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range history.Location {
        history.EntityData.Children.Append(types.GetSegmentPath(history.Location[i]), types.YChild{"Location", history.Location[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_History_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "history"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/all/history/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_Location
type Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    LocationName interface{}
}

func (location *Clear_ClearAsicErrorsGrp_ClearDevice_AllInstances_All_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "all"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-clear-ncs5500:clear/clear-asic-errors-grp/clear-device/all-instances/all/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

