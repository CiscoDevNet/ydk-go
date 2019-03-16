// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all hardware devices managed in Sysadmin.
// 
// Copyright(c) 2015-2016 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_controllers_ncs5501

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_controllers_ncs5501"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-controllers-NCS5501 controller}", reflect.TypeOf(Controller{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller", reflect.TypeOf(Controller{}))
}

// Controller
type Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fabric resource commands.
    Fabric Controller_Fabric

    
    CardMgr Controller_CardMgr
}

func (controller *Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-controllers-ncs5501"
    controller.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller"
    controller.EntityData.AbsolutePath = controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("fabric", types.YChild{"Fabric", &controller.Fabric})
    controller.EntityData.Children.Append("card_mgr", types.YChild{"CardMgr", &controller.CardMgr})
    controller.EntityData.Leafs = types.NewOrderedMap()

    controller.EntityData.YListKeys = []string {}

    return &(controller.EntityData)
}

// Controller_Fabric
// Fabric resource commands
type Controller_Fabric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Oper Controller_Fabric_Oper
}

func (fabric *Controller_Fabric) GetEntityData() *types.CommonEntityData {
    fabric.EntityData.YFilter = fabric.YFilter
    fabric.EntityData.YangName = "fabric"
    fabric.EntityData.BundleName = "cisco_ios_xr"
    fabric.EntityData.ParentYangName = "controller"
    fabric.EntityData.SegmentPath = "fabric"
    fabric.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/" + fabric.EntityData.SegmentPath
    fabric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabric.EntityData.Children = types.NewOrderedMap()
    fabric.EntityData.Children.Append("oper", types.YChild{"Oper", &fabric.Oper})
    fabric.EntityData.Leafs = types.NewOrderedMap()

    fabric.EntityData.YListKeys = []string {}

    return &(fabric.EntityData)
}

// Controller_Fabric_Oper
type Controller_Fabric_Oper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FGID management information.
    Fgid Controller_Fabric_Oper_Fgid
}

func (oper *Controller_Fabric_Oper) GetEntityData() *types.CommonEntityData {
    oper.EntityData.YFilter = oper.YFilter
    oper.EntityData.YangName = "oper"
    oper.EntityData.BundleName = "cisco_ios_xr"
    oper.EntityData.ParentYangName = "fabric"
    oper.EntityData.SegmentPath = "oper"
    oper.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/" + oper.EntityData.SegmentPath
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = types.NewOrderedMap()
    oper.EntityData.Children.Append("fgid", types.YChild{"Fgid", &oper.Fgid})
    oper.EntityData.Leafs = types.NewOrderedMap()

    oper.EntityData.YListKeys = []string {}

    return &(oper.EntityData)
}

// Controller_Fabric_Oper_Fgid
// FGID management information
type Controller_Fabric_Oper_Fgid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Information Controller_Fabric_Oper_Fgid_Information

    
    Resource Controller_Fabric_Oper_Fgid_Resource

    
    Statistics Controller_Fabric_Oper_Fgid_Statistics

    
    FgidMgr Controller_Fabric_Oper_Fgid_FgidMgr

    // The type is slice of Controller_Fabric_Oper_Fgid_ProgramError.
    ProgramError []*Controller_Fabric_Oper_Fgid_ProgramError
}

func (fgid *Controller_Fabric_Oper_Fgid) GetEntityData() *types.CommonEntityData {
    fgid.EntityData.YFilter = fgid.YFilter
    fgid.EntityData.YangName = "fgid"
    fgid.EntityData.BundleName = "cisco_ios_xr"
    fgid.EntityData.ParentYangName = "oper"
    fgid.EntityData.SegmentPath = "fgid"
    fgid.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/" + fgid.EntityData.SegmentPath
    fgid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fgid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fgid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fgid.EntityData.Children = types.NewOrderedMap()
    fgid.EntityData.Children.Append("information", types.YChild{"Information", &fgid.Information})
    fgid.EntityData.Children.Append("resource", types.YChild{"Resource", &fgid.Resource})
    fgid.EntityData.Children.Append("statistics", types.YChild{"Statistics", &fgid.Statistics})
    fgid.EntityData.Children.Append("fgid_mgr", types.YChild{"FgidMgr", &fgid.FgidMgr})
    fgid.EntityData.Children.Append("program_error", types.YChild{"ProgramError", nil})
    for i := range fgid.ProgramError {
        fgid.EntityData.Children.Append(types.GetSegmentPath(fgid.ProgramError[i]), types.YChild{"ProgramError", fgid.ProgramError[i]})
    }
    fgid.EntityData.Leafs = types.NewOrderedMap()

    fgid.EntityData.YListKeys = []string {}

    return &(fgid.EntityData)
}

// Controller_Fabric_Oper_Fgid_Information
type Controller_Fabric_Oper_Fgid_Information struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Fabric_Oper_Fgid_Information_Id.
    Id []*Controller_Fabric_Oper_Fgid_Information_Id
}

func (information *Controller_Fabric_Oper_Fgid_Information) GetEntityData() *types.CommonEntityData {
    information.EntityData.YFilter = information.YFilter
    information.EntityData.YangName = "information"
    information.EntityData.BundleName = "cisco_ios_xr"
    information.EntityData.ParentYangName = "fgid"
    information.EntityData.SegmentPath = "information"
    information.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/" + information.EntityData.SegmentPath
    information.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    information.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    information.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    information.EntityData.Children = types.NewOrderedMap()
    information.EntityData.Children.Append("id", types.YChild{"Id", nil})
    for i := range information.Id {
        information.EntityData.Children.Append(types.GetSegmentPath(information.Id[i]), types.YChild{"Id", information.Id[i]})
    }
    information.EntityData.Leafs = types.NewOrderedMap()

    information.EntityData.YListKeys = []string {}

    return &(information.EntityData)
}

// Controller_Fabric_Oper_Fgid_Information_Id
type Controller_Fabric_Oper_Fgid_Information_Id struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..128000.
    FgidId interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    TotalAssoFabricqIds interface{}

    // The type is string.
    AssoClientInfo interface{}

    // The type is slice of Controller_Fabric_Oper_Fgid_Information_Id_HexBitmaps.
    HexBitmaps []*Controller_Fabric_Oper_Fgid_Information_Id_HexBitmaps

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Information_Id_BinaryBitmaps.
    BinaryBitmaps []*Controller_Fabric_Oper_Fgid_Information_Id_BinaryBitmaps

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Information_Id_AssoFabricqIds.
    AssoFabricqIds []*Controller_Fabric_Oper_Fgid_Information_Id_AssoFabricqIds

    // The type is slice of Controller_Fabric_Oper_Fgid_Information_Id_Drivers.
    Drivers []*Controller_Fabric_Oper_Fgid_Information_Id_Drivers
}

func (id *Controller_Fabric_Oper_Fgid_Information_Id) GetEntityData() *types.CommonEntityData {
    id.EntityData.YFilter = id.YFilter
    id.EntityData.YangName = "id"
    id.EntityData.BundleName = "cisco_ios_xr"
    id.EntityData.ParentYangName = "information"
    id.EntityData.SegmentPath = "id" + types.AddKeyToken(id.FgidId, "fgid_id")
    id.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/information/" + id.EntityData.SegmentPath
    id.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    id.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    id.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    id.EntityData.Children = types.NewOrderedMap()
    id.EntityData.Children.Append("hex_bitmaps", types.YChild{"HexBitmaps", nil})
    for i := range id.HexBitmaps {
        id.EntityData.Children.Append(types.GetSegmentPath(id.HexBitmaps[i]), types.YChild{"HexBitmaps", id.HexBitmaps[i]})
    }
    id.EntityData.Children.Append("binary_bitmaps", types.YChild{"BinaryBitmaps", nil})
    for i := range id.BinaryBitmaps {
        id.EntityData.Children.Append(types.GetSegmentPath(id.BinaryBitmaps[i]), types.YChild{"BinaryBitmaps", id.BinaryBitmaps[i]})
    }
    id.EntityData.Children.Append("asso_fabricq_ids", types.YChild{"AssoFabricqIds", nil})
    for i := range id.AssoFabricqIds {
        id.EntityData.Children.Append(types.GetSegmentPath(id.AssoFabricqIds[i]), types.YChild{"AssoFabricqIds", id.AssoFabricqIds[i]})
    }
    id.EntityData.Children.Append("drivers", types.YChild{"Drivers", nil})
    for i := range id.Drivers {
        id.EntityData.Children.Append(types.GetSegmentPath(id.Drivers[i]), types.YChild{"Drivers", id.Drivers[i]})
    }
    id.EntityData.Leafs = types.NewOrderedMap()
    id.EntityData.Leafs.Append("fgid_id", types.YLeaf{"FgidId", id.FgidId})
    id.EntityData.Leafs.Append("total_asso_fabricq_ids", types.YLeaf{"TotalAssoFabricqIds", id.TotalAssoFabricqIds})
    id.EntityData.Leafs.Append("asso_client_info", types.YLeaf{"AssoClientInfo", id.AssoClientInfo})

    id.EntityData.YListKeys = []string {"FgidId"}

    return &(id.EntityData)
}

// Controller_Fabric_Oper_Fgid_Information_Id_HexBitmaps
type Controller_Fabric_Oper_Fgid_Information_Id_HexBitmaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    RackNumber interface{}

    // The type is string.
    Bitmap interface{}
}

func (hexBitmaps *Controller_Fabric_Oper_Fgid_Information_Id_HexBitmaps) GetEntityData() *types.CommonEntityData {
    hexBitmaps.EntityData.YFilter = hexBitmaps.YFilter
    hexBitmaps.EntityData.YangName = "hex_bitmaps"
    hexBitmaps.EntityData.BundleName = "cisco_ios_xr"
    hexBitmaps.EntityData.ParentYangName = "id"
    hexBitmaps.EntityData.SegmentPath = "hex_bitmaps" + types.AddKeyToken(hexBitmaps.RackNumber, "rack_number")
    hexBitmaps.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/information/id/" + hexBitmaps.EntityData.SegmentPath
    hexBitmaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hexBitmaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hexBitmaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hexBitmaps.EntityData.Children = types.NewOrderedMap()
    hexBitmaps.EntityData.Leafs = types.NewOrderedMap()
    hexBitmaps.EntityData.Leafs.Append("rack_number", types.YLeaf{"RackNumber", hexBitmaps.RackNumber})
    hexBitmaps.EntityData.Leafs.Append("bitmap", types.YLeaf{"Bitmap", hexBitmaps.Bitmap})

    hexBitmaps.EntityData.YListKeys = []string {"RackNumber"}

    return &(hexBitmaps.EntityData)
}

// Controller_Fabric_Oper_Fgid_Information_Id_BinaryBitmaps
type Controller_Fabric_Oper_Fgid_Information_Id_BinaryBitmaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    RackNumber interface{}

    // The type is string.
    Bitmap interface{}
}

func (binaryBitmaps *Controller_Fabric_Oper_Fgid_Information_Id_BinaryBitmaps) GetEntityData() *types.CommonEntityData {
    binaryBitmaps.EntityData.YFilter = binaryBitmaps.YFilter
    binaryBitmaps.EntityData.YangName = "binary_bitmaps"
    binaryBitmaps.EntityData.BundleName = "cisco_ios_xr"
    binaryBitmaps.EntityData.ParentYangName = "id"
    binaryBitmaps.EntityData.SegmentPath = "binary_bitmaps" + types.AddKeyToken(binaryBitmaps.RackNumber, "rack_number")
    binaryBitmaps.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/information/id/" + binaryBitmaps.EntityData.SegmentPath
    binaryBitmaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binaryBitmaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binaryBitmaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binaryBitmaps.EntityData.Children = types.NewOrderedMap()
    binaryBitmaps.EntityData.Leafs = types.NewOrderedMap()
    binaryBitmaps.EntityData.Leafs.Append("rack_number", types.YLeaf{"RackNumber", binaryBitmaps.RackNumber})
    binaryBitmaps.EntityData.Leafs.Append("bitmap", types.YLeaf{"Bitmap", binaryBitmaps.Bitmap})

    binaryBitmaps.EntityData.YListKeys = []string {"RackNumber"}

    return &(binaryBitmaps.EntityData)
}

// Controller_Fabric_Oper_Fgid_Information_Id_AssoFabricqIds
type Controller_Fabric_Oper_Fgid_Information_Id_AssoFabricqIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    FabricqId interface{}
}

func (assoFabricqIds *Controller_Fabric_Oper_Fgid_Information_Id_AssoFabricqIds) GetEntityData() *types.CommonEntityData {
    assoFabricqIds.EntityData.YFilter = assoFabricqIds.YFilter
    assoFabricqIds.EntityData.YangName = "asso_fabricq_ids"
    assoFabricqIds.EntityData.BundleName = "cisco_ios_xr"
    assoFabricqIds.EntityData.ParentYangName = "id"
    assoFabricqIds.EntityData.SegmentPath = "asso_fabricq_ids" + types.AddKeyToken(assoFabricqIds.FabricqId, "fabricq_id")
    assoFabricqIds.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/information/id/" + assoFabricqIds.EntityData.SegmentPath
    assoFabricqIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    assoFabricqIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    assoFabricqIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    assoFabricqIds.EntityData.Children = types.NewOrderedMap()
    assoFabricqIds.EntityData.Leafs = types.NewOrderedMap()
    assoFabricqIds.EntityData.Leafs.Append("fabricq_id", types.YLeaf{"FabricqId", assoFabricqIds.FabricqId})

    assoFabricqIds.EntityData.YListKeys = []string {"FabricqId"}

    return &(assoFabricqIds.EntityData)
}

// Controller_Fabric_Oper_Fgid_Information_Id_Drivers
type Controller_Fabric_Oper_Fgid_Information_Id_Drivers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    RackNumber interface{}

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Information_Id_Drivers_Clients.
    Clients []*Controller_Fabric_Oper_Fgid_Information_Id_Drivers_Clients
}

func (drivers *Controller_Fabric_Oper_Fgid_Information_Id_Drivers) GetEntityData() *types.CommonEntityData {
    drivers.EntityData.YFilter = drivers.YFilter
    drivers.EntityData.YangName = "drivers"
    drivers.EntityData.BundleName = "cisco_ios_xr"
    drivers.EntityData.ParentYangName = "id"
    drivers.EntityData.SegmentPath = "drivers" + types.AddKeyToken(drivers.RackNumber, "rack_number")
    drivers.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/information/id/" + drivers.EntityData.SegmentPath
    drivers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drivers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drivers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drivers.EntityData.Children = types.NewOrderedMap()
    drivers.EntityData.Children.Append("clients", types.YChild{"Clients", nil})
    for i := range drivers.Clients {
        drivers.EntityData.Children.Append(types.GetSegmentPath(drivers.Clients[i]), types.YChild{"Clients", drivers.Clients[i]})
    }
    drivers.EntityData.Leafs = types.NewOrderedMap()
    drivers.EntityData.Leafs.Append("rack_number", types.YLeaf{"RackNumber", drivers.RackNumber})

    drivers.EntityData.YListKeys = []string {"RackNumber"}

    return &(drivers.EntityData)
}

// Controller_Fabric_Oper_Fgid_Information_Id_Drivers_Clients
type Controller_Fabric_Oper_Fgid_Information_Id_Drivers_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    ClientIdx interface{}

    // The type is bool. The default value is false.
    ShowAsic0 interface{}

    // The type is string.
    Asic0Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic1 interface{}

    // The type is string.
    Asic1Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic2 interface{}

    // The type is string.
    Asic2Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic3 interface{}

    // The type is string.
    Asic3Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic4 interface{}

    // The type is string.
    Asic4Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic5 interface{}

    // The type is string.
    Asic5Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic6 interface{}

    // The type is string.
    Asic6Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic7 interface{}

    // The type is string.
    Asic7Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic8 interface{}

    // The type is string.
    Asic8Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic9 interface{}

    // The type is string.
    Asic9Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic10 interface{}

    // The type is string.
    Asic10Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic11 interface{}

    // The type is string.
    Asic11Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic12 interface{}

    // The type is string.
    Asic12Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic13 interface{}

    // The type is string.
    Asic13Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic14 interface{}

    // The type is string.
    Asic14Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic15 interface{}

    // The type is string.
    Asic15Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic16 interface{}

    // The type is string.
    Asic16Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic17 interface{}

    // The type is string.
    Asic17Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic18 interface{}

    // The type is string.
    Asic18Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic19 interface{}

    // The type is string.
    Asic19Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic20 interface{}

    // The type is string.
    Asic20Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic21 interface{}

    // The type is string.
    Asic21Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic22 interface{}

    // The type is string.
    Asic22Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic23 interface{}

    // The type is string.
    Asic23Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic24 interface{}

    // The type is string.
    Asic24Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic25 interface{}

    // The type is string.
    Asic25Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic26 interface{}

    // The type is string.
    Asic26Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic27 interface{}

    // The type is string.
    Asic27Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic28 interface{}

    // The type is string.
    Asic28Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic29 interface{}

    // The type is string.
    Asic29Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic30 interface{}

    // The type is string.
    Asic30Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic31 interface{}

    // The type is string.
    Asic31Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic32 interface{}

    // The type is string.
    Asic32Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic33 interface{}

    // The type is string.
    Asic33Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic34 interface{}

    // The type is string.
    Asic34Bitmap interface{}

    // The type is bool. The default value is false.
    ShowAsic35 interface{}

    // The type is string.
    Asic35Bitmap interface{}
}

func (clients *Controller_Fabric_Oper_Fgid_Information_Id_Drivers_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "drivers"
    clients.EntityData.SegmentPath = "clients" + types.AddKeyToken(clients.ClientIdx, "client_idx")
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/information/id/drivers/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Leafs = types.NewOrderedMap()
    clients.EntityData.Leafs.Append("client_idx", types.YLeaf{"ClientIdx", clients.ClientIdx})
    clients.EntityData.Leafs.Append("show_asic_0", types.YLeaf{"ShowAsic0", clients.ShowAsic0})
    clients.EntityData.Leafs.Append("asic_0_bitmap", types.YLeaf{"Asic0Bitmap", clients.Asic0Bitmap})
    clients.EntityData.Leafs.Append("show_asic_1", types.YLeaf{"ShowAsic1", clients.ShowAsic1})
    clients.EntityData.Leafs.Append("asic_1_bitmap", types.YLeaf{"Asic1Bitmap", clients.Asic1Bitmap})
    clients.EntityData.Leafs.Append("show_asic_2", types.YLeaf{"ShowAsic2", clients.ShowAsic2})
    clients.EntityData.Leafs.Append("asic_2_bitmap", types.YLeaf{"Asic2Bitmap", clients.Asic2Bitmap})
    clients.EntityData.Leafs.Append("show_asic_3", types.YLeaf{"ShowAsic3", clients.ShowAsic3})
    clients.EntityData.Leafs.Append("asic_3_bitmap", types.YLeaf{"Asic3Bitmap", clients.Asic3Bitmap})
    clients.EntityData.Leafs.Append("show_asic_4", types.YLeaf{"ShowAsic4", clients.ShowAsic4})
    clients.EntityData.Leafs.Append("asic_4_bitmap", types.YLeaf{"Asic4Bitmap", clients.Asic4Bitmap})
    clients.EntityData.Leafs.Append("show_asic_5", types.YLeaf{"ShowAsic5", clients.ShowAsic5})
    clients.EntityData.Leafs.Append("asic_5_bitmap", types.YLeaf{"Asic5Bitmap", clients.Asic5Bitmap})
    clients.EntityData.Leafs.Append("show_asic_6", types.YLeaf{"ShowAsic6", clients.ShowAsic6})
    clients.EntityData.Leafs.Append("asic_6_bitmap", types.YLeaf{"Asic6Bitmap", clients.Asic6Bitmap})
    clients.EntityData.Leafs.Append("show_asic_7", types.YLeaf{"ShowAsic7", clients.ShowAsic7})
    clients.EntityData.Leafs.Append("asic_7_bitmap", types.YLeaf{"Asic7Bitmap", clients.Asic7Bitmap})
    clients.EntityData.Leafs.Append("show_asic_8", types.YLeaf{"ShowAsic8", clients.ShowAsic8})
    clients.EntityData.Leafs.Append("asic_8_bitmap", types.YLeaf{"Asic8Bitmap", clients.Asic8Bitmap})
    clients.EntityData.Leafs.Append("show_asic_9", types.YLeaf{"ShowAsic9", clients.ShowAsic9})
    clients.EntityData.Leafs.Append("asic_9_bitmap", types.YLeaf{"Asic9Bitmap", clients.Asic9Bitmap})
    clients.EntityData.Leafs.Append("show_asic_10", types.YLeaf{"ShowAsic10", clients.ShowAsic10})
    clients.EntityData.Leafs.Append("asic_10_bitmap", types.YLeaf{"Asic10Bitmap", clients.Asic10Bitmap})
    clients.EntityData.Leafs.Append("show_asic_11", types.YLeaf{"ShowAsic11", clients.ShowAsic11})
    clients.EntityData.Leafs.Append("asic_11_bitmap", types.YLeaf{"Asic11Bitmap", clients.Asic11Bitmap})
    clients.EntityData.Leafs.Append("show_asic_12", types.YLeaf{"ShowAsic12", clients.ShowAsic12})
    clients.EntityData.Leafs.Append("asic_12_bitmap", types.YLeaf{"Asic12Bitmap", clients.Asic12Bitmap})
    clients.EntityData.Leafs.Append("show_asic_13", types.YLeaf{"ShowAsic13", clients.ShowAsic13})
    clients.EntityData.Leafs.Append("asic_13_bitmap", types.YLeaf{"Asic13Bitmap", clients.Asic13Bitmap})
    clients.EntityData.Leafs.Append("show_asic_14", types.YLeaf{"ShowAsic14", clients.ShowAsic14})
    clients.EntityData.Leafs.Append("asic_14_bitmap", types.YLeaf{"Asic14Bitmap", clients.Asic14Bitmap})
    clients.EntityData.Leafs.Append("show_asic_15", types.YLeaf{"ShowAsic15", clients.ShowAsic15})
    clients.EntityData.Leafs.Append("asic_15_bitmap", types.YLeaf{"Asic15Bitmap", clients.Asic15Bitmap})
    clients.EntityData.Leafs.Append("show_asic_16", types.YLeaf{"ShowAsic16", clients.ShowAsic16})
    clients.EntityData.Leafs.Append("asic_16_bitmap", types.YLeaf{"Asic16Bitmap", clients.Asic16Bitmap})
    clients.EntityData.Leafs.Append("show_asic_17", types.YLeaf{"ShowAsic17", clients.ShowAsic17})
    clients.EntityData.Leafs.Append("asic_17_bitmap", types.YLeaf{"Asic17Bitmap", clients.Asic17Bitmap})
    clients.EntityData.Leafs.Append("show_asic_18", types.YLeaf{"ShowAsic18", clients.ShowAsic18})
    clients.EntityData.Leafs.Append("asic_18_bitmap", types.YLeaf{"Asic18Bitmap", clients.Asic18Bitmap})
    clients.EntityData.Leafs.Append("show_asic_19", types.YLeaf{"ShowAsic19", clients.ShowAsic19})
    clients.EntityData.Leafs.Append("asic_19_bitmap", types.YLeaf{"Asic19Bitmap", clients.Asic19Bitmap})
    clients.EntityData.Leafs.Append("show_asic_20", types.YLeaf{"ShowAsic20", clients.ShowAsic20})
    clients.EntityData.Leafs.Append("asic_20_bitmap", types.YLeaf{"Asic20Bitmap", clients.Asic20Bitmap})
    clients.EntityData.Leafs.Append("show_asic_21", types.YLeaf{"ShowAsic21", clients.ShowAsic21})
    clients.EntityData.Leafs.Append("asic_21_bitmap", types.YLeaf{"Asic21Bitmap", clients.Asic21Bitmap})
    clients.EntityData.Leafs.Append("show_asic_22", types.YLeaf{"ShowAsic22", clients.ShowAsic22})
    clients.EntityData.Leafs.Append("asic_22_bitmap", types.YLeaf{"Asic22Bitmap", clients.Asic22Bitmap})
    clients.EntityData.Leafs.Append("show_asic_23", types.YLeaf{"ShowAsic23", clients.ShowAsic23})
    clients.EntityData.Leafs.Append("asic_23_bitmap", types.YLeaf{"Asic23Bitmap", clients.Asic23Bitmap})
    clients.EntityData.Leafs.Append("show_asic_24", types.YLeaf{"ShowAsic24", clients.ShowAsic24})
    clients.EntityData.Leafs.Append("asic_24_bitmap", types.YLeaf{"Asic24Bitmap", clients.Asic24Bitmap})
    clients.EntityData.Leafs.Append("show_asic_25", types.YLeaf{"ShowAsic25", clients.ShowAsic25})
    clients.EntityData.Leafs.Append("asic_25_bitmap", types.YLeaf{"Asic25Bitmap", clients.Asic25Bitmap})
    clients.EntityData.Leafs.Append("show_asic_26", types.YLeaf{"ShowAsic26", clients.ShowAsic26})
    clients.EntityData.Leafs.Append("asic_26_bitmap", types.YLeaf{"Asic26Bitmap", clients.Asic26Bitmap})
    clients.EntityData.Leafs.Append("show_asic_27", types.YLeaf{"ShowAsic27", clients.ShowAsic27})
    clients.EntityData.Leafs.Append("asic_27_bitmap", types.YLeaf{"Asic27Bitmap", clients.Asic27Bitmap})
    clients.EntityData.Leafs.Append("show_asic_28", types.YLeaf{"ShowAsic28", clients.ShowAsic28})
    clients.EntityData.Leafs.Append("asic_28_bitmap", types.YLeaf{"Asic28Bitmap", clients.Asic28Bitmap})
    clients.EntityData.Leafs.Append("show_asic_29", types.YLeaf{"ShowAsic29", clients.ShowAsic29})
    clients.EntityData.Leafs.Append("asic_29_bitmap", types.YLeaf{"Asic29Bitmap", clients.Asic29Bitmap})
    clients.EntityData.Leafs.Append("show_asic_30", types.YLeaf{"ShowAsic30", clients.ShowAsic30})
    clients.EntityData.Leafs.Append("asic_30_bitmap", types.YLeaf{"Asic30Bitmap", clients.Asic30Bitmap})
    clients.EntityData.Leafs.Append("show_asic_31", types.YLeaf{"ShowAsic31", clients.ShowAsic31})
    clients.EntityData.Leafs.Append("asic_31_bitmap", types.YLeaf{"Asic31Bitmap", clients.Asic31Bitmap})
    clients.EntityData.Leafs.Append("show_asic_32", types.YLeaf{"ShowAsic32", clients.ShowAsic32})
    clients.EntityData.Leafs.Append("asic_32_bitmap", types.YLeaf{"Asic32Bitmap", clients.Asic32Bitmap})
    clients.EntityData.Leafs.Append("show_asic_33", types.YLeaf{"ShowAsic33", clients.ShowAsic33})
    clients.EntityData.Leafs.Append("asic_33_bitmap", types.YLeaf{"Asic33Bitmap", clients.Asic33Bitmap})
    clients.EntityData.Leafs.Append("show_asic_34", types.YLeaf{"ShowAsic34", clients.ShowAsic34})
    clients.EntityData.Leafs.Append("asic_34_bitmap", types.YLeaf{"Asic34Bitmap", clients.Asic34Bitmap})
    clients.EntityData.Leafs.Append("show_asic_35", types.YLeaf{"ShowAsic35", clients.ShowAsic35})
    clients.EntityData.Leafs.Append("asic_35_bitmap", types.YLeaf{"Asic35Bitmap", clients.Asic35Bitmap})

    clients.EntityData.YListKeys = []string {"ClientIdx"}

    return &(clients.EntityData)
}

// Controller_Fabric_Oper_Fgid_Resource
type Controller_Fabric_Oper_Fgid_Resource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Fabric_Oper_Fgid_Resource_Sdr.
    Sdr []*Controller_Fabric_Oper_Fgid_Resource_Sdr
}

func (resource *Controller_Fabric_Oper_Fgid_Resource) GetEntityData() *types.CommonEntityData {
    resource.EntityData.YFilter = resource.YFilter
    resource.EntityData.YangName = "resource"
    resource.EntityData.BundleName = "cisco_ios_xr"
    resource.EntityData.ParentYangName = "fgid"
    resource.EntityData.SegmentPath = "resource"
    resource.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/" + resource.EntityData.SegmentPath
    resource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resource.EntityData.Children = types.NewOrderedMap()
    resource.EntityData.Children.Append("sdr", types.YChild{"Sdr", nil})
    for i := range resource.Sdr {
        resource.EntityData.Children.Append(types.GetSegmentPath(resource.Sdr[i]), types.YChild{"Sdr", resource.Sdr[i]})
    }
    resource.EntityData.Leafs = types.NewOrderedMap()

    resource.EntityData.YListKeys = []string {}

    return &(resource.EntityData)
}

// Controller_Fabric_Oper_Fgid_Resource_Sdr
type Controller_Fabric_Oper_Fgid_Resource_Sdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    SdrName interface{}

    // The type is string. The default value is Secure Domain Router name..
    Description interface{}

    // The type is slice of Controller_Fabric_Oper_Fgid_Resource_Sdr_Application.
    Application []*Controller_Fabric_Oper_Fgid_Resource_Sdr_Application
}

func (sdr *Controller_Fabric_Oper_Fgid_Resource_Sdr) GetEntityData() *types.CommonEntityData {
    sdr.EntityData.YFilter = sdr.YFilter
    sdr.EntityData.YangName = "sdr"
    sdr.EntityData.BundleName = "cisco_ios_xr"
    sdr.EntityData.ParentYangName = "resource"
    sdr.EntityData.SegmentPath = "sdr" + types.AddKeyToken(sdr.SdrName, "sdr_name")
    sdr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/resource/" + sdr.EntityData.SegmentPath
    sdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdr.EntityData.Children = types.NewOrderedMap()
    sdr.EntityData.Children.Append("application", types.YChild{"Application", nil})
    for i := range sdr.Application {
        sdr.EntityData.Children.Append(types.GetSegmentPath(sdr.Application[i]), types.YChild{"Application", sdr.Application[i]})
    }
    sdr.EntityData.Leafs = types.NewOrderedMap()
    sdr.EntityData.Leafs.Append("sdr_name", types.YLeaf{"SdrName", sdr.SdrName})
    sdr.EntityData.Leafs.Append("description", types.YLeaf{"Description", sdr.Description})

    sdr.EntityData.YListKeys = []string {"SdrName"}

    return &(sdr.EntityData)
}

// Controller_Fabric_Oper_Fgid_Resource_Sdr_Application
type Controller_Fabric_Oper_Fgid_Resource_Sdr_Application struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    AppName interface{}

    // The type is string. The default value is application..
    Description interface{}

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange.
    IdsRange []*Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange
}

func (application *Controller_Fabric_Oper_Fgid_Resource_Sdr_Application) GetEntityData() *types.CommonEntityData {
    application.EntityData.YFilter = application.YFilter
    application.EntityData.YangName = "application"
    application.EntityData.BundleName = "cisco_ios_xr"
    application.EntityData.ParentYangName = "sdr"
    application.EntityData.SegmentPath = "application" + types.AddKeyToken(application.AppName, "app_name")
    application.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/resource/sdr/" + application.EntityData.SegmentPath
    application.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    application.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    application.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    application.EntityData.Children = types.NewOrderedMap()
    application.EntityData.Children.Append("ids_range", types.YChild{"IdsRange", nil})
    for i := range application.IdsRange {
        application.EntityData.Children.Append(types.GetSegmentPath(application.IdsRange[i]), types.YChild{"IdsRange", application.IdsRange[i]})
    }
    application.EntityData.Leafs = types.NewOrderedMap()
    application.EntityData.Leafs.Append("app_name", types.YLeaf{"AppName", application.AppName})
    application.EntityData.Leafs.Append("description", types.YLeaf{"Description", application.Description})

    application.EntityData.YListKeys = []string {"AppName"}

    return &(application.EntityData)
}

// Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange
type Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..128000.
    Id interface{}

    // This attribute is a key. The type is interface{} with range: 0..128000.
    Elements interface{}

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange_FgidIds.
    FgidIds []*Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange_FgidIds
}

func (idsRange *Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange) GetEntityData() *types.CommonEntityData {
    idsRange.EntityData.YFilter = idsRange.YFilter
    idsRange.EntityData.YangName = "ids_range"
    idsRange.EntityData.BundleName = "cisco_ios_xr"
    idsRange.EntityData.ParentYangName = "application"
    idsRange.EntityData.SegmentPath = "ids_range" + types.AddKeyToken(idsRange.Id, "id") + types.AddKeyToken(idsRange.Elements, "elements")
    idsRange.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/resource/sdr/application/" + idsRange.EntityData.SegmentPath
    idsRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idsRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idsRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idsRange.EntityData.Children = types.NewOrderedMap()
    idsRange.EntityData.Children.Append("fgid_ids", types.YChild{"FgidIds", nil})
    for i := range idsRange.FgidIds {
        idsRange.EntityData.Children.Append(types.GetSegmentPath(idsRange.FgidIds[i]), types.YChild{"FgidIds", idsRange.FgidIds[i]})
    }
    idsRange.EntityData.Leafs = types.NewOrderedMap()
    idsRange.EntityData.Leafs.Append("id", types.YLeaf{"Id", idsRange.Id})
    idsRange.EntityData.Leafs.Append("elements", types.YLeaf{"Elements", idsRange.Elements})

    idsRange.EntityData.YListKeys = []string {"Id", "Elements"}

    return &(idsRange.EntityData)
}

// Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange_FgidIds
type Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange_FgidIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    FgidId interface{}

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    LineIdx interface{}

    // The type is string.
    SdrNameH interface{}

    // The type is string.
    AppNameH interface{}
}

func (fgidIds *Controller_Fabric_Oper_Fgid_Resource_Sdr_Application_IdsRange_FgidIds) GetEntityData() *types.CommonEntityData {
    fgidIds.EntityData.YFilter = fgidIds.YFilter
    fgidIds.EntityData.YangName = "fgid_ids"
    fgidIds.EntityData.BundleName = "cisco_ios_xr"
    fgidIds.EntityData.ParentYangName = "ids_range"
    fgidIds.EntityData.SegmentPath = "fgid_ids" + types.AddKeyToken(fgidIds.FgidId, "fgid_id") + types.AddKeyToken(fgidIds.LineIdx, "line_idx")
    fgidIds.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/resource/sdr/application/ids_range/" + fgidIds.EntityData.SegmentPath
    fgidIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fgidIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fgidIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fgidIds.EntityData.Children = types.NewOrderedMap()
    fgidIds.EntityData.Leafs = types.NewOrderedMap()
    fgidIds.EntityData.Leafs.Append("fgid_id", types.YLeaf{"FgidId", fgidIds.FgidId})
    fgidIds.EntityData.Leafs.Append("line_idx", types.YLeaf{"LineIdx", fgidIds.LineIdx})
    fgidIds.EntityData.Leafs.Append("sdr_name_h", types.YLeaf{"SdrNameH", fgidIds.SdrNameH})
    fgidIds.EntityData.Leafs.Append("app_name_h", types.YLeaf{"AppNameH", fgidIds.AppNameH})

    fgidIds.EntityData.YListKeys = []string {"FgidId", "LineIdx"}

    return &(fgidIds.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics
type Controller_Fabric_Oper_Fgid_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    All Controller_Fabric_Oper_Fgid_Statistics_All

    
    Sdr Controller_Fabric_Oper_Fgid_Statistics_Sdr

    
    Pool Controller_Fabric_Oper_Fgid_Statistics_Pool

    
    System Controller_Fabric_Oper_Fgid_Statistics_System
}

func (statistics *Controller_Fabric_Oper_Fgid_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "fgid"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("all", types.YChild{"All", &statistics.All})
    statistics.EntityData.Children.Append("sdr", types.YChild{"Sdr", &statistics.Sdr})
    statistics.EntityData.Children.Append("pool", types.YChild{"Pool", &statistics.Pool})
    statistics.EntityData.Children.Append("system", types.YChild{"System", &statistics.System})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_All
type Controller_Fabric_Oper_Fgid_Statistics_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Fabric_Oper_Fgid_Statistics_All_StatsList.
    StatsList []*Controller_Fabric_Oper_Fgid_Statistics_All_StatsList

    // The type is slice of Controller_Fabric_Oper_Fgid_Statistics_All_SdrList.
    SdrList []*Controller_Fabric_Oper_Fgid_Statistics_All_SdrList

    // The type is slice of Controller_Fabric_Oper_Fgid_Statistics_All_PoolList.
    PoolList []*Controller_Fabric_Oper_Fgid_Statistics_All_PoolList
}

func (all *Controller_Fabric_Oper_Fgid_Statistics_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "statistics"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("stats_list", types.YChild{"StatsList", nil})
    for i := range all.StatsList {
        all.EntityData.Children.Append(types.GetSegmentPath(all.StatsList[i]), types.YChild{"StatsList", all.StatsList[i]})
    }
    all.EntityData.Children.Append("sdr_list", types.YChild{"SdrList", nil})
    for i := range all.SdrList {
        all.EntityData.Children.Append(types.GetSegmentPath(all.SdrList[i]), types.YChild{"SdrList", all.SdrList[i]})
    }
    all.EntityData.Children.Append("pool_list", types.YChild{"PoolList", nil})
    for i := range all.PoolList {
        all.EntityData.Children.Append(types.GetSegmentPath(all.PoolList[i]), types.YChild{"PoolList", all.PoolList[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_All_StatsList
type Controller_Fabric_Oper_Fgid_Statistics_All_StatsList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    SystemStats interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemTotalFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemInuseFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemHwmFgids interface{}
}

func (statsList *Controller_Fabric_Oper_Fgid_Statistics_All_StatsList) GetEntityData() *types.CommonEntityData {
    statsList.EntityData.YFilter = statsList.YFilter
    statsList.EntityData.YangName = "stats_list"
    statsList.EntityData.BundleName = "cisco_ios_xr"
    statsList.EntityData.ParentYangName = "all"
    statsList.EntityData.SegmentPath = "stats_list" + types.AddKeyToken(statsList.SystemStats, "system_stats")
    statsList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/all/" + statsList.EntityData.SegmentPath
    statsList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsList.EntityData.Children = types.NewOrderedMap()
    statsList.EntityData.Leafs = types.NewOrderedMap()
    statsList.EntityData.Leafs.Append("system_stats", types.YLeaf{"SystemStats", statsList.SystemStats})
    statsList.EntityData.Leafs.Append("system_total_fgids", types.YLeaf{"SystemTotalFgids", statsList.SystemTotalFgids})
    statsList.EntityData.Leafs.Append("system_inuse_fgids", types.YLeaf{"SystemInuseFgids", statsList.SystemInuseFgids})
    statsList.EntityData.Leafs.Append("system_hwm_fgids", types.YLeaf{"SystemHwmFgids", statsList.SystemHwmFgids})

    statsList.EntityData.YListKeys = []string {"SystemStats"}

    return &(statsList.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_All_SdrList
type Controller_Fabric_Oper_Fgid_Statistics_All_SdrList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    SdrName interface{}

    // The type is string. The default value is Secure Domain Router name..
    Description interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrTotalFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrInuseFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrHwmFgids interface{}

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Statistics_All_SdrList_Application.
    Application []*Controller_Fabric_Oper_Fgid_Statistics_All_SdrList_Application
}

func (sdrList *Controller_Fabric_Oper_Fgid_Statistics_All_SdrList) GetEntityData() *types.CommonEntityData {
    sdrList.EntityData.YFilter = sdrList.YFilter
    sdrList.EntityData.YangName = "sdr_list"
    sdrList.EntityData.BundleName = "cisco_ios_xr"
    sdrList.EntityData.ParentYangName = "all"
    sdrList.EntityData.SegmentPath = "sdr_list" + types.AddKeyToken(sdrList.SdrName, "sdr_name")
    sdrList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/all/" + sdrList.EntityData.SegmentPath
    sdrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrList.EntityData.Children = types.NewOrderedMap()
    sdrList.EntityData.Children.Append("application", types.YChild{"Application", nil})
    for i := range sdrList.Application {
        sdrList.EntityData.Children.Append(types.GetSegmentPath(sdrList.Application[i]), types.YChild{"Application", sdrList.Application[i]})
    }
    sdrList.EntityData.Leafs = types.NewOrderedMap()
    sdrList.EntityData.Leafs.Append("sdr_name", types.YLeaf{"SdrName", sdrList.SdrName})
    sdrList.EntityData.Leafs.Append("description", types.YLeaf{"Description", sdrList.Description})
    sdrList.EntityData.Leafs.Append("sdr_total_fgids", types.YLeaf{"SdrTotalFgids", sdrList.SdrTotalFgids})
    sdrList.EntityData.Leafs.Append("sdr_inuse_fgids", types.YLeaf{"SdrInuseFgids", sdrList.SdrInuseFgids})
    sdrList.EntityData.Leafs.Append("sdr_hwm_fgids", types.YLeaf{"SdrHwmFgids", sdrList.SdrHwmFgids})

    sdrList.EntityData.YListKeys = []string {"SdrName"}

    return &(sdrList.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_All_SdrList_Application
type Controller_Fabric_Oper_Fgid_Statistics_All_SdrList_Application struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    AppName interface{}

    // The type is string. The default value is application..
    Description interface{}

    // The type is interface{} with range: 0..4294967295.
    AppId interface{}

    // The type is interface{} with range: 0..4294967295.
    PoolId interface{}

    // The type is interface{} with range: 0..4294967295.
    InuseFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    HwmFgids interface{}
}

func (application *Controller_Fabric_Oper_Fgid_Statistics_All_SdrList_Application) GetEntityData() *types.CommonEntityData {
    application.EntityData.YFilter = application.YFilter
    application.EntityData.YangName = "application"
    application.EntityData.BundleName = "cisco_ios_xr"
    application.EntityData.ParentYangName = "sdr_list"
    application.EntityData.SegmentPath = "application" + types.AddKeyToken(application.AppName, "app_name")
    application.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/all/sdr_list/" + application.EntityData.SegmentPath
    application.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    application.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    application.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    application.EntityData.Children = types.NewOrderedMap()
    application.EntityData.Leafs = types.NewOrderedMap()
    application.EntityData.Leafs.Append("app_name", types.YLeaf{"AppName", application.AppName})
    application.EntityData.Leafs.Append("description", types.YLeaf{"Description", application.Description})
    application.EntityData.Leafs.Append("app_id", types.YLeaf{"AppId", application.AppId})
    application.EntityData.Leafs.Append("pool_id", types.YLeaf{"PoolId", application.PoolId})
    application.EntityData.Leafs.Append("inuse_fgids", types.YLeaf{"InuseFgids", application.InuseFgids})
    application.EntityData.Leafs.Append("hwm_fgids", types.YLeaf{"HwmFgids", application.HwmFgids})

    application.EntityData.YListKeys = []string {"AppName"}

    return &(application.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_All_PoolList
type Controller_Fabric_Oper_Fgid_Statistics_All_PoolList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    PoolId interface{}

    // The type is string.
    PoolName interface{}

    // The type is string.
    PoolType interface{}

    // The type is string.
    StartFgid interface{}

    // The type is interface{} with range: 0..4294967295.
    TotalFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    CurrentFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    HwmFgids interface{}
}

func (poolList *Controller_Fabric_Oper_Fgid_Statistics_All_PoolList) GetEntityData() *types.CommonEntityData {
    poolList.EntityData.YFilter = poolList.YFilter
    poolList.EntityData.YangName = "pool_list"
    poolList.EntityData.BundleName = "cisco_ios_xr"
    poolList.EntityData.ParentYangName = "all"
    poolList.EntityData.SegmentPath = "pool_list" + types.AddKeyToken(poolList.PoolId, "pool_id")
    poolList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/all/" + poolList.EntityData.SegmentPath
    poolList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolList.EntityData.Children = types.NewOrderedMap()
    poolList.EntityData.Leafs = types.NewOrderedMap()
    poolList.EntityData.Leafs.Append("pool_id", types.YLeaf{"PoolId", poolList.PoolId})
    poolList.EntityData.Leafs.Append("pool_name", types.YLeaf{"PoolName", poolList.PoolName})
    poolList.EntityData.Leafs.Append("pool_type", types.YLeaf{"PoolType", poolList.PoolType})
    poolList.EntityData.Leafs.Append("start_fgid", types.YLeaf{"StartFgid", poolList.StartFgid})
    poolList.EntityData.Leafs.Append("total_fgids", types.YLeaf{"TotalFgids", poolList.TotalFgids})
    poolList.EntityData.Leafs.Append("current_fgids", types.YLeaf{"CurrentFgids", poolList.CurrentFgids})
    poolList.EntityData.Leafs.Append("hwm_fgids", types.YLeaf{"HwmFgids", poolList.HwmFgids})

    poolList.EntityData.YListKeys = []string {"PoolId"}

    return &(poolList.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_Sdr
type Controller_Fabric_Oper_Fgid_Statistics_Sdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList.
    SdrList []*Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList
}

func (sdr *Controller_Fabric_Oper_Fgid_Statistics_Sdr) GetEntityData() *types.CommonEntityData {
    sdr.EntityData.YFilter = sdr.YFilter
    sdr.EntityData.YangName = "sdr"
    sdr.EntityData.BundleName = "cisco_ios_xr"
    sdr.EntityData.ParentYangName = "statistics"
    sdr.EntityData.SegmentPath = "sdr"
    sdr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/" + sdr.EntityData.SegmentPath
    sdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdr.EntityData.Children = types.NewOrderedMap()
    sdr.EntityData.Children.Append("sdr_list", types.YChild{"SdrList", nil})
    for i := range sdr.SdrList {
        sdr.EntityData.Children.Append(types.GetSegmentPath(sdr.SdrList[i]), types.YChild{"SdrList", sdr.SdrList[i]})
    }
    sdr.EntityData.Leafs = types.NewOrderedMap()

    sdr.EntityData.YListKeys = []string {}

    return &(sdr.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList
type Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    SdrName interface{}

    // The type is string. The default value is Secure Domain Router name..
    Description interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrTotalFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrInuseFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrHwmFgids interface{}

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList_Application.
    Application []*Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList_Application
}

func (sdrList *Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList) GetEntityData() *types.CommonEntityData {
    sdrList.EntityData.YFilter = sdrList.YFilter
    sdrList.EntityData.YangName = "sdr_list"
    sdrList.EntityData.BundleName = "cisco_ios_xr"
    sdrList.EntityData.ParentYangName = "sdr"
    sdrList.EntityData.SegmentPath = "sdr_list" + types.AddKeyToken(sdrList.SdrName, "sdr_name")
    sdrList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/sdr/" + sdrList.EntityData.SegmentPath
    sdrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrList.EntityData.Children = types.NewOrderedMap()
    sdrList.EntityData.Children.Append("application", types.YChild{"Application", nil})
    for i := range sdrList.Application {
        sdrList.EntityData.Children.Append(types.GetSegmentPath(sdrList.Application[i]), types.YChild{"Application", sdrList.Application[i]})
    }
    sdrList.EntityData.Leafs = types.NewOrderedMap()
    sdrList.EntityData.Leafs.Append("sdr_name", types.YLeaf{"SdrName", sdrList.SdrName})
    sdrList.EntityData.Leafs.Append("description", types.YLeaf{"Description", sdrList.Description})
    sdrList.EntityData.Leafs.Append("sdr_total_fgids", types.YLeaf{"SdrTotalFgids", sdrList.SdrTotalFgids})
    sdrList.EntityData.Leafs.Append("sdr_inuse_fgids", types.YLeaf{"SdrInuseFgids", sdrList.SdrInuseFgids})
    sdrList.EntityData.Leafs.Append("sdr_hwm_fgids", types.YLeaf{"SdrHwmFgids", sdrList.SdrHwmFgids})

    sdrList.EntityData.YListKeys = []string {"SdrName"}

    return &(sdrList.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList_Application
type Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList_Application struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    AppName interface{}

    // The type is string. The default value is application..
    Description interface{}

    // The type is interface{} with range: 0..4294967295.
    AppId interface{}

    // The type is interface{} with range: 0..4294967295.
    PoolId interface{}

    // The type is interface{} with range: 0..4294967295.
    InuseFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    HwmFgids interface{}
}

func (application *Controller_Fabric_Oper_Fgid_Statistics_Sdr_SdrList_Application) GetEntityData() *types.CommonEntityData {
    application.EntityData.YFilter = application.YFilter
    application.EntityData.YangName = "application"
    application.EntityData.BundleName = "cisco_ios_xr"
    application.EntityData.ParentYangName = "sdr_list"
    application.EntityData.SegmentPath = "application" + types.AddKeyToken(application.AppName, "app_name")
    application.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/sdr/sdr_list/" + application.EntityData.SegmentPath
    application.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    application.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    application.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    application.EntityData.Children = types.NewOrderedMap()
    application.EntityData.Leafs = types.NewOrderedMap()
    application.EntityData.Leafs.Append("app_name", types.YLeaf{"AppName", application.AppName})
    application.EntityData.Leafs.Append("description", types.YLeaf{"Description", application.Description})
    application.EntityData.Leafs.Append("app_id", types.YLeaf{"AppId", application.AppId})
    application.EntityData.Leafs.Append("pool_id", types.YLeaf{"PoolId", application.PoolId})
    application.EntityData.Leafs.Append("inuse_fgids", types.YLeaf{"InuseFgids", application.InuseFgids})
    application.EntityData.Leafs.Append("hwm_fgids", types.YLeaf{"HwmFgids", application.HwmFgids})

    application.EntityData.YListKeys = []string {"AppName"}

    return &(application.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_Pool
type Controller_Fabric_Oper_Fgid_Statistics_Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Fabric_Oper_Fgid_Statistics_Pool_PoolList.
    PoolList []*Controller_Fabric_Oper_Fgid_Statistics_Pool_PoolList
}

func (pool *Controller_Fabric_Oper_Fgid_Statistics_Pool) GetEntityData() *types.CommonEntityData {
    pool.EntityData.YFilter = pool.YFilter
    pool.EntityData.YangName = "pool"
    pool.EntityData.BundleName = "cisco_ios_xr"
    pool.EntityData.ParentYangName = "statistics"
    pool.EntityData.SegmentPath = "pool"
    pool.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/" + pool.EntityData.SegmentPath
    pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pool.EntityData.Children = types.NewOrderedMap()
    pool.EntityData.Children.Append("pool_list", types.YChild{"PoolList", nil})
    for i := range pool.PoolList {
        pool.EntityData.Children.Append(types.GetSegmentPath(pool.PoolList[i]), types.YChild{"PoolList", pool.PoolList[i]})
    }
    pool.EntityData.Leafs = types.NewOrderedMap()

    pool.EntityData.YListKeys = []string {}

    return &(pool.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_Pool_PoolList
type Controller_Fabric_Oper_Fgid_Statistics_Pool_PoolList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    PoolId interface{}

    // The type is string.
    PoolName interface{}

    // The type is string.
    PoolType interface{}

    // The type is string.
    StartFgid interface{}

    // The type is interface{} with range: 0..4294967295.
    TotalFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    CurrentFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    HwmFgids interface{}
}

func (poolList *Controller_Fabric_Oper_Fgid_Statistics_Pool_PoolList) GetEntityData() *types.CommonEntityData {
    poolList.EntityData.YFilter = poolList.YFilter
    poolList.EntityData.YangName = "pool_list"
    poolList.EntityData.BundleName = "cisco_ios_xr"
    poolList.EntityData.ParentYangName = "pool"
    poolList.EntityData.SegmentPath = "pool_list" + types.AddKeyToken(poolList.PoolId, "pool_id")
    poolList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/pool/" + poolList.EntityData.SegmentPath
    poolList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolList.EntityData.Children = types.NewOrderedMap()
    poolList.EntityData.Leafs = types.NewOrderedMap()
    poolList.EntityData.Leafs.Append("pool_id", types.YLeaf{"PoolId", poolList.PoolId})
    poolList.EntityData.Leafs.Append("pool_name", types.YLeaf{"PoolName", poolList.PoolName})
    poolList.EntityData.Leafs.Append("pool_type", types.YLeaf{"PoolType", poolList.PoolType})
    poolList.EntityData.Leafs.Append("start_fgid", types.YLeaf{"StartFgid", poolList.StartFgid})
    poolList.EntityData.Leafs.Append("total_fgids", types.YLeaf{"TotalFgids", poolList.TotalFgids})
    poolList.EntityData.Leafs.Append("current_fgids", types.YLeaf{"CurrentFgids", poolList.CurrentFgids})
    poolList.EntityData.Leafs.Append("hwm_fgids", types.YLeaf{"HwmFgids", poolList.HwmFgids})

    poolList.EntityData.YListKeys = []string {"PoolId"}

    return &(poolList.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_System
type Controller_Fabric_Oper_Fgid_Statistics_System struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_Statistics_System_StatsList.
    StatsList []*Controller_Fabric_Oper_Fgid_Statistics_System_StatsList
}

func (system *Controller_Fabric_Oper_Fgid_Statistics_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xr"
    system.EntityData.ParentYangName = "statistics"
    system.EntityData.SegmentPath = "system"
    system.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/" + system.EntityData.SegmentPath
    system.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Children.Append("stats_list", types.YChild{"StatsList", nil})
    for i := range system.StatsList {
        system.EntityData.Children.Append(types.GetSegmentPath(system.StatsList[i]), types.YChild{"StatsList", system.StatsList[i]})
    }
    system.EntityData.Leafs = types.NewOrderedMap()

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// Controller_Fabric_Oper_Fgid_Statistics_System_StatsList
type Controller_Fabric_Oper_Fgid_Statistics_System_StatsList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    SystemStats interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemTotalFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemInuseFgids interface{}

    // The type is interface{} with range: 0..4294967295.
    SystemHwmFgids interface{}
}

func (statsList *Controller_Fabric_Oper_Fgid_Statistics_System_StatsList) GetEntityData() *types.CommonEntityData {
    statsList.EntityData.YFilter = statsList.YFilter
    statsList.EntityData.YangName = "stats_list"
    statsList.EntityData.BundleName = "cisco_ios_xr"
    statsList.EntityData.ParentYangName = "system"
    statsList.EntityData.SegmentPath = "stats_list" + types.AddKeyToken(statsList.SystemStats, "system_stats")
    statsList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/statistics/system/" + statsList.EntityData.SegmentPath
    statsList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsList.EntityData.Children = types.NewOrderedMap()
    statsList.EntityData.Leafs = types.NewOrderedMap()
    statsList.EntityData.Leafs.Append("system_stats", types.YLeaf{"SystemStats", statsList.SystemStats})
    statsList.EntityData.Leafs.Append("system_total_fgids", types.YLeaf{"SystemTotalFgids", statsList.SystemTotalFgids})
    statsList.EntityData.Leafs.Append("system_inuse_fgids", types.YLeaf{"SystemInuseFgids", statsList.SystemInuseFgids})
    statsList.EntityData.Leafs.Append("system_hwm_fgids", types.YLeaf{"SystemHwmFgids", statsList.SystemHwmFgids})

    statsList.EntityData.YListKeys = []string {"SystemStats"}

    return &(statsList.EntityData)
}

// Controller_Fabric_Oper_Fgid_FgidMgr
type Controller_Fabric_Oper_Fgid_FgidMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of
    // Controller_Fabric_Oper_Fgid_FgidMgr_Trace.
    Trace []*Controller_Fabric_Oper_Fgid_FgidMgr_Trace
}

func (fgidMgr *Controller_Fabric_Oper_Fgid_FgidMgr) GetEntityData() *types.CommonEntityData {
    fgidMgr.EntityData.YFilter = fgidMgr.YFilter
    fgidMgr.EntityData.YangName = "fgid_mgr"
    fgidMgr.EntityData.BundleName = "cisco_ios_xr"
    fgidMgr.EntityData.ParentYangName = "fgid"
    fgidMgr.EntityData.SegmentPath = "fgid_mgr"
    fgidMgr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/" + fgidMgr.EntityData.SegmentPath
    fgidMgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fgidMgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fgidMgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fgidMgr.EntityData.Children = types.NewOrderedMap()
    fgidMgr.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range fgidMgr.Trace {
        fgidMgr.EntityData.Children.Append(types.GetSegmentPath(fgidMgr.Trace[i]), types.YChild{"Trace", fgidMgr.Trace[i]})
    }
    fgidMgr.EntityData.Leafs = types.NewOrderedMap()

    fgidMgr.EntityData.YListKeys = []string {}

    return &(fgidMgr.EntityData)
}

// Controller_Fabric_Oper_Fgid_FgidMgr_Trace
// show traceable processes
type Controller_Fabric_Oper_Fgid_FgidMgr_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location.
    Location []*Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location
}

func (trace *Controller_Fabric_Oper_Fgid_FgidMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "fgid_mgr"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/fgid_mgr/" + trace.EntityData.SegmentPath
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range trace.Location {
        trace.EntityData.Children.Append(types.GetSegmentPath(trace.Location[i]), types.YChild{"Location", trace.Location[i]})
    }
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("buffer", types.YLeaf{"Buffer", trace.Buffer})

    trace.EntityData.YListKeys = []string {"Buffer"}

    return &(trace.EntityData)
}

// Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location
type Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions.
    AllOptions []*Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions
}

func (location *Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/fgid_mgr/trace/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("all-options", types.YChild{"AllOptions", nil})
    for i := range location.AllOptions {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AllOptions[i]), types.YChild{"AllOptions", location.AllOptions[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location_name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions
type Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/fgid_mgr/trace/location/" + allOptions.EntityData.SegmentPath
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        types.SetYListKey(allOptions.TraceBlocks[i], i)
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions_TraceBlocks
type Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Controller_Fabric_Oper_Fgid_FgidMgr_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/fgid_mgr/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// Controller_Fabric_Oper_Fgid_ProgramError
type Controller_Fabric_Oper_Fgid_ProgramError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..128000.
    Start interface{}

    // This attribute is a key. The type is interface{} with range: 0..128000.
    End interface{}

    // The type is slice of Controller_Fabric_Oper_Fgid_ProgramError_Rack.
    Rack []*Controller_Fabric_Oper_Fgid_ProgramError_Rack
}

func (programError *Controller_Fabric_Oper_Fgid_ProgramError) GetEntityData() *types.CommonEntityData {
    programError.EntityData.YFilter = programError.YFilter
    programError.EntityData.YangName = "program_error"
    programError.EntityData.BundleName = "cisco_ios_xr"
    programError.EntityData.ParentYangName = "fgid"
    programError.EntityData.SegmentPath = "program_error" + types.AddKeyToken(programError.Start, "start") + types.AddKeyToken(programError.End, "end")
    programError.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/" + programError.EntityData.SegmentPath
    programError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programError.EntityData.Children = types.NewOrderedMap()
    programError.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range programError.Rack {
        programError.EntityData.Children.Append(types.GetSegmentPath(programError.Rack[i]), types.YChild{"Rack", programError.Rack[i]})
    }
    programError.EntityData.Leafs = types.NewOrderedMap()
    programError.EntityData.Leafs.Append("start", types.YLeaf{"Start", programError.Start})
    programError.EntityData.Leafs.Append("end", types.YLeaf{"End", programError.End})

    programError.EntityData.YListKeys = []string {"Start", "End"}

    return &(programError.EntityData)
}

// Controller_Fabric_Oper_Fgid_ProgramError_Rack
type Controller_Fabric_Oper_Fgid_ProgramError_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    RackId interface{}

    // The type is string.
    RackIdStr interface{}

    // The type is string.
    FgidsInError interface{}

    // The type is bool.
    FoundFgidsInError interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    TotalErrorFgids interface{}

    // The type is bool.
    IncorrectFgidsRange interface{}

    // The type is bool.
    CmdNotSupported interface{}
}

func (rack *Controller_Fabric_Oper_Fgid_ProgramError_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "program_error"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.RackId, "rack_id")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/fabric/oper/fgid/program_error/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("rack_id", types.YLeaf{"RackId", rack.RackId})
    rack.EntityData.Leafs.Append("rack_id_str", types.YLeaf{"RackIdStr", rack.RackIdStr})
    rack.EntityData.Leafs.Append("fgids_in_error", types.YLeaf{"FgidsInError", rack.FgidsInError})
    rack.EntityData.Leafs.Append("found_fgids_in_error", types.YLeaf{"FoundFgidsInError", rack.FoundFgidsInError})
    rack.EntityData.Leafs.Append("total_error_fgids", types.YLeaf{"TotalErrorFgids", rack.TotalErrorFgids})
    rack.EntityData.Leafs.Append("incorrect_fgids_range", types.YLeaf{"IncorrectFgidsRange", rack.IncorrectFgidsRange})
    rack.EntityData.Leafs.Append("cmd_not_supported", types.YLeaf{"CmdNotSupported", rack.CmdNotSupported})

    rack.EntityData.YListKeys = []string {"RackId"}

    return &(rack.EntityData)
}

// Controller_CardMgr
type Controller_CardMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Controller_CardMgr_Trace.
    Trace []*Controller_CardMgr_Trace

    
    Inventory Controller_CardMgr_Inventory

    
    EventHistory Controller_CardMgr_EventHistory

    
    NotifHistory Controller_CardMgr_NotifHistory

    
    OirHistory Controller_CardMgr_OirHistory

    
    Iofpga Controller_CardMgr_Iofpga

    
    Bootloader Controller_CardMgr_Bootloader
}

func (cardMgr *Controller_CardMgr) GetEntityData() *types.CommonEntityData {
    cardMgr.EntityData.YFilter = cardMgr.YFilter
    cardMgr.EntityData.YangName = "card_mgr"
    cardMgr.EntityData.BundleName = "cisco_ios_xr"
    cardMgr.EntityData.ParentYangName = "controller"
    cardMgr.EntityData.SegmentPath = "card_mgr"
    cardMgr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/" + cardMgr.EntityData.SegmentPath
    cardMgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardMgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardMgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardMgr.EntityData.Children = types.NewOrderedMap()
    cardMgr.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range cardMgr.Trace {
        cardMgr.EntityData.Children.Append(types.GetSegmentPath(cardMgr.Trace[i]), types.YChild{"Trace", cardMgr.Trace[i]})
    }
    cardMgr.EntityData.Children.Append("inventory", types.YChild{"Inventory", &cardMgr.Inventory})
    cardMgr.EntityData.Children.Append("event-history", types.YChild{"EventHistory", &cardMgr.EventHistory})
    cardMgr.EntityData.Children.Append("notif-history", types.YChild{"NotifHistory", &cardMgr.NotifHistory})
    cardMgr.EntityData.Children.Append("oir-history", types.YChild{"OirHistory", &cardMgr.OirHistory})
    cardMgr.EntityData.Children.Append("iofpga", types.YChild{"Iofpga", &cardMgr.Iofpga})
    cardMgr.EntityData.Children.Append("bootloader", types.YChild{"Bootloader", &cardMgr.Bootloader})
    cardMgr.EntityData.Leafs = types.NewOrderedMap()

    cardMgr.EntityData.YListKeys = []string {}

    return &(cardMgr.EntityData)
}

// Controller_CardMgr_Trace
// show traceable processes
type Controller_CardMgr_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Controller_CardMgr_Trace_Location.
    Location []*Controller_CardMgr_Trace_Location
}

func (trace *Controller_CardMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "card_mgr"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/" + trace.EntityData.SegmentPath
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range trace.Location {
        trace.EntityData.Children.Append(types.GetSegmentPath(trace.Location[i]), types.YChild{"Location", trace.Location[i]})
    }
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("buffer", types.YLeaf{"Buffer", trace.Buffer})

    trace.EntityData.YListKeys = []string {"Buffer"}

    return &(trace.EntityData)
}

// Controller_CardMgr_Trace_Location
type Controller_CardMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Controller_CardMgr_Trace_Location_AllOptions.
    AllOptions []*Controller_CardMgr_Trace_Location_AllOptions
}

func (location *Controller_CardMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/trace/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("all-options", types.YChild{"AllOptions", nil})
    for i := range location.AllOptions {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AllOptions[i]), types.YChild{"AllOptions", location.AllOptions[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location_name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Controller_CardMgr_Trace_Location_AllOptions
type Controller_CardMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Controller_CardMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Controller_CardMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Controller_CardMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/trace/location/" + allOptions.EntityData.SegmentPath
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        types.SetYListKey(allOptions.TraceBlocks[i], i)
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Controller_CardMgr_Trace_Location_AllOptions_TraceBlocks
type Controller_CardMgr_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Controller_CardMgr_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// Controller_CardMgr_Inventory
type Controller_CardMgr_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Summary Controller_CardMgr_Inventory_Summary

    
    Detail Controller_CardMgr_Inventory_Detail
}

func (inventory *Controller_CardMgr_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "card_mgr"
    inventory.EntityData.SegmentPath = "inventory"
    inventory.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/" + inventory.EntityData.SegmentPath
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("summary", types.YChild{"Summary", &inventory.Summary})
    inventory.EntityData.Children.Append("detail", types.YChild{"Detail", &inventory.Detail})
    inventory.EntityData.Leafs = types.NewOrderedMap()

    inventory.EntityData.YListKeys = []string {}

    return &(inventory.EntityData)
}

// Controller_CardMgr_Inventory_Summary
type Controller_CardMgr_Inventory_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_CardMgr_Inventory_Summary_CardMgrInvSummary.
    CardMgrInvSummary []*Controller_CardMgr_Inventory_Summary_CardMgrInvSummary
}

func (summary *Controller_CardMgr_Inventory_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "inventory"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/inventory/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("card_mgr_inv_summary", types.YChild{"CardMgrInvSummary", nil})
    for i := range summary.CardMgrInvSummary {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.CardMgrInvSummary[i]), types.YChild{"CardMgrInvSummary", summary.CardMgrInvSummary[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Controller_CardMgr_Inventory_Summary_CardMgrInvSummary
type Controller_CardMgr_Inventory_Summary_CardMgrInvSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. This attribute is mandatory.
    Location interface{}

    // The type is string.
    CardMgrInvPIDString interface{}

    // The type is interface{} with range: 0..4294967295.
    CardMgrInvSlotNumber interface{}

    // The type is string.
    CardMgrInvSerialNumber interface{}

    // The type is string.
    CardMgrInvHwVersion interface{}

    // The type is string.
    CardMgrInvCardState interface{}
}

func (cardMgrInvSummary *Controller_CardMgr_Inventory_Summary_CardMgrInvSummary) GetEntityData() *types.CommonEntityData {
    cardMgrInvSummary.EntityData.YFilter = cardMgrInvSummary.YFilter
    cardMgrInvSummary.EntityData.YangName = "card_mgr_inv_summary"
    cardMgrInvSummary.EntityData.BundleName = "cisco_ios_xr"
    cardMgrInvSummary.EntityData.ParentYangName = "summary"
    cardMgrInvSummary.EntityData.SegmentPath = "card_mgr_inv_summary" + types.AddKeyToken(cardMgrInvSummary.Location, "location")
    cardMgrInvSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/inventory/summary/" + cardMgrInvSummary.EntityData.SegmentPath
    cardMgrInvSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardMgrInvSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardMgrInvSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardMgrInvSummary.EntityData.Children = types.NewOrderedMap()
    cardMgrInvSummary.EntityData.Leafs = types.NewOrderedMap()
    cardMgrInvSummary.EntityData.Leafs.Append("location", types.YLeaf{"Location", cardMgrInvSummary.Location})
    cardMgrInvSummary.EntityData.Leafs.Append("card_mgr_inv_PID_string", types.YLeaf{"CardMgrInvPIDString", cardMgrInvSummary.CardMgrInvPIDString})
    cardMgrInvSummary.EntityData.Leafs.Append("card_mgr_inv_slot_number", types.YLeaf{"CardMgrInvSlotNumber", cardMgrInvSummary.CardMgrInvSlotNumber})
    cardMgrInvSummary.EntityData.Leafs.Append("card_mgr_inv_serial_number", types.YLeaf{"CardMgrInvSerialNumber", cardMgrInvSummary.CardMgrInvSerialNumber})
    cardMgrInvSummary.EntityData.Leafs.Append("card_mgr_inv_hw_version", types.YLeaf{"CardMgrInvHwVersion", cardMgrInvSummary.CardMgrInvHwVersion})
    cardMgrInvSummary.EntityData.Leafs.Append("card_mgr_inv_card_state", types.YLeaf{"CardMgrInvCardState", cardMgrInvSummary.CardMgrInvCardState})

    cardMgrInvSummary.EntityData.YListKeys = []string {"Location"}

    return &(cardMgrInvSummary.EntityData)
}

// Controller_CardMgr_Inventory_Detail
type Controller_CardMgr_Inventory_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_Inventory_Detail_CardMgrInvDetail.
    CardMgrInvDetail []*Controller_CardMgr_Inventory_Detail_CardMgrInvDetail
}

func (detail *Controller_CardMgr_Inventory_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "inventory"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/inventory/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("card_mgr_inv_detail", types.YChild{"CardMgrInvDetail", nil})
    for i := range detail.CardMgrInvDetail {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.CardMgrInvDetail[i]), types.YChild{"CardMgrInvDetail", detail.CardMgrInvDetail[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Controller_CardMgr_Inventory_Detail_CardMgrInvDetail
type Controller_CardMgr_Inventory_Detail_CardMgrInvDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    
    CardMgrInvDetailList Controller_CardMgr_Inventory_Detail_CardMgrInvDetail_CardMgrInvDetailList
}

func (cardMgrInvDetail *Controller_CardMgr_Inventory_Detail_CardMgrInvDetail) GetEntityData() *types.CommonEntityData {
    cardMgrInvDetail.EntityData.YFilter = cardMgrInvDetail.YFilter
    cardMgrInvDetail.EntityData.YangName = "card_mgr_inv_detail"
    cardMgrInvDetail.EntityData.BundleName = "cisco_ios_xr"
    cardMgrInvDetail.EntityData.ParentYangName = "detail"
    cardMgrInvDetail.EntityData.SegmentPath = "card_mgr_inv_detail" + types.AddKeyToken(cardMgrInvDetail.Location, "location")
    cardMgrInvDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/inventory/detail/" + cardMgrInvDetail.EntityData.SegmentPath
    cardMgrInvDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardMgrInvDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardMgrInvDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardMgrInvDetail.EntityData.Children = types.NewOrderedMap()
    cardMgrInvDetail.EntityData.Children.Append("card_mgr_inv_detail_list", types.YChild{"CardMgrInvDetailList", &cardMgrInvDetail.CardMgrInvDetailList})
    cardMgrInvDetail.EntityData.Leafs = types.NewOrderedMap()
    cardMgrInvDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", cardMgrInvDetail.Location})

    cardMgrInvDetail.EntityData.YListKeys = []string {"Location"}

    return &(cardMgrInvDetail.EntityData)
}

// Controller_CardMgr_Inventory_Detail_CardMgrInvDetail_CardMgrInvDetailList
type Controller_CardMgr_Inventory_Detail_CardMgrInvDetail_CardMgrInvDetailList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    CardMgrInvDetailValues []interface{}
}

func (cardMgrInvDetailList *Controller_CardMgr_Inventory_Detail_CardMgrInvDetail_CardMgrInvDetailList) GetEntityData() *types.CommonEntityData {
    cardMgrInvDetailList.EntityData.YFilter = cardMgrInvDetailList.YFilter
    cardMgrInvDetailList.EntityData.YangName = "card_mgr_inv_detail_list"
    cardMgrInvDetailList.EntityData.BundleName = "cisco_ios_xr"
    cardMgrInvDetailList.EntityData.ParentYangName = "card_mgr_inv_detail"
    cardMgrInvDetailList.EntityData.SegmentPath = "card_mgr_inv_detail_list"
    cardMgrInvDetailList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/inventory/detail/card_mgr_inv_detail/" + cardMgrInvDetailList.EntityData.SegmentPath
    cardMgrInvDetailList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardMgrInvDetailList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardMgrInvDetailList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardMgrInvDetailList.EntityData.Children = types.NewOrderedMap()
    cardMgrInvDetailList.EntityData.Leafs = types.NewOrderedMap()
    cardMgrInvDetailList.EntityData.Leafs.Append("card_mgr_inv_detail_values", types.YLeaf{"CardMgrInvDetailValues", cardMgrInvDetailList.CardMgrInvDetailValues})

    cardMgrInvDetailList.EntityData.YListKeys = []string {}

    return &(cardMgrInvDetailList.EntityData)
}

// Controller_CardMgr_EventHistory
type Controller_CardMgr_EventHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Brief Controller_CardMgr_EventHistory_Brief

    
    Detail Controller_CardMgr_EventHistory_Detail
}

func (eventHistory *Controller_CardMgr_EventHistory) GetEntityData() *types.CommonEntityData {
    eventHistory.EntityData.YFilter = eventHistory.YFilter
    eventHistory.EntityData.YangName = "event-history"
    eventHistory.EntityData.BundleName = "cisco_ios_xr"
    eventHistory.EntityData.ParentYangName = "card_mgr"
    eventHistory.EntityData.SegmentPath = "event-history"
    eventHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/" + eventHistory.EntityData.SegmentPath
    eventHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventHistory.EntityData.Children = types.NewOrderedMap()
    eventHistory.EntityData.Children.Append("brief", types.YChild{"Brief", &eventHistory.Brief})
    eventHistory.EntityData.Children.Append("detail", types.YChild{"Detail", &eventHistory.Detail})
    eventHistory.EntityData.Leafs = types.NewOrderedMap()

    eventHistory.EntityData.YListKeys = []string {}

    return &(eventHistory.EntityData)
}

// Controller_CardMgr_EventHistory_Brief
type Controller_CardMgr_EventHistory_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_EventHistory_Brief_Location.
    Location []*Controller_CardMgr_EventHistory_Brief_Location
}

func (brief *Controller_CardMgr_EventHistory_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "event-history"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/event-history/" + brief.EntityData.SegmentPath
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range brief.Location {
        brief.EntityData.Children.Append(types.GetSegmentPath(brief.Location[i]), types.YChild{"Location", brief.Location[i]})
    }
    brief.EntityData.Leafs = types.NewOrderedMap()

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// Controller_CardMgr_EventHistory_Brief_Location
type Controller_CardMgr_EventHistory_Brief_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    
    CardEventHistBrief Controller_CardMgr_EventHistory_Brief_Location_CardEventHistBrief
}

func (location *Controller_CardMgr_EventHistory_Brief_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "brief"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/event-history/brief/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("card_event_hist_brief", types.YChild{"CardEventHistBrief", &location.CardEventHistBrief})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Controller_CardMgr_EventHistory_Brief_Location_CardEventHistBrief
type Controller_CardMgr_EventHistory_Brief_Location_CardEventHistBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    CardEventHistBriefValues []interface{}
}

func (cardEventHistBrief *Controller_CardMgr_EventHistory_Brief_Location_CardEventHistBrief) GetEntityData() *types.CommonEntityData {
    cardEventHistBrief.EntityData.YFilter = cardEventHistBrief.YFilter
    cardEventHistBrief.EntityData.YangName = "card_event_hist_brief"
    cardEventHistBrief.EntityData.BundleName = "cisco_ios_xr"
    cardEventHistBrief.EntityData.ParentYangName = "location"
    cardEventHistBrief.EntityData.SegmentPath = "card_event_hist_brief"
    cardEventHistBrief.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/event-history/brief/location/" + cardEventHistBrief.EntityData.SegmentPath
    cardEventHistBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardEventHistBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardEventHistBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardEventHistBrief.EntityData.Children = types.NewOrderedMap()
    cardEventHistBrief.EntityData.Leafs = types.NewOrderedMap()
    cardEventHistBrief.EntityData.Leafs.Append("card_event_hist_brief_values", types.YLeaf{"CardEventHistBriefValues", cardEventHistBrief.CardEventHistBriefValues})

    cardEventHistBrief.EntityData.YListKeys = []string {}

    return &(cardEventHistBrief.EntityData)
}

// Controller_CardMgr_EventHistory_Detail
type Controller_CardMgr_EventHistory_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_EventHistory_Detail_Location.
    Location []*Controller_CardMgr_EventHistory_Detail_Location
}

func (detail *Controller_CardMgr_EventHistory_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "event-history"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/event-history/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range detail.Location {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.Location[i]), types.YChild{"Location", detail.Location[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Controller_CardMgr_EventHistory_Detail_Location
type Controller_CardMgr_EventHistory_Detail_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    
    CardEventHistDetail Controller_CardMgr_EventHistory_Detail_Location_CardEventHistDetail
}

func (location *Controller_CardMgr_EventHistory_Detail_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "detail"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/event-history/detail/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("card_event_hist_detail", types.YChild{"CardEventHistDetail", &location.CardEventHistDetail})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Controller_CardMgr_EventHistory_Detail_Location_CardEventHistDetail
type Controller_CardMgr_EventHistory_Detail_Location_CardEventHistDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    CardEventHistDetailValues []interface{}
}

func (cardEventHistDetail *Controller_CardMgr_EventHistory_Detail_Location_CardEventHistDetail) GetEntityData() *types.CommonEntityData {
    cardEventHistDetail.EntityData.YFilter = cardEventHistDetail.YFilter
    cardEventHistDetail.EntityData.YangName = "card_event_hist_detail"
    cardEventHistDetail.EntityData.BundleName = "cisco_ios_xr"
    cardEventHistDetail.EntityData.ParentYangName = "location"
    cardEventHistDetail.EntityData.SegmentPath = "card_event_hist_detail"
    cardEventHistDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/event-history/detail/location/" + cardEventHistDetail.EntityData.SegmentPath
    cardEventHistDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardEventHistDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardEventHistDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardEventHistDetail.EntityData.Children = types.NewOrderedMap()
    cardEventHistDetail.EntityData.Leafs = types.NewOrderedMap()
    cardEventHistDetail.EntityData.Leafs.Append("card_event_hist_detail_values", types.YLeaf{"CardEventHistDetailValues", cardEventHistDetail.CardEventHistDetailValues})

    cardEventHistDetail.EntityData.YListKeys = []string {}

    return &(cardEventHistDetail.EntityData)
}

// Controller_CardMgr_NotifHistory
type Controller_CardMgr_NotifHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Brief Controller_CardMgr_NotifHistory_Brief

    
    Detail Controller_CardMgr_NotifHistory_Detail
}

func (notifHistory *Controller_CardMgr_NotifHistory) GetEntityData() *types.CommonEntityData {
    notifHistory.EntityData.YFilter = notifHistory.YFilter
    notifHistory.EntityData.YangName = "notif-history"
    notifHistory.EntityData.BundleName = "cisco_ios_xr"
    notifHistory.EntityData.ParentYangName = "card_mgr"
    notifHistory.EntityData.SegmentPath = "notif-history"
    notifHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/" + notifHistory.EntityData.SegmentPath
    notifHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notifHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notifHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notifHistory.EntityData.Children = types.NewOrderedMap()
    notifHistory.EntityData.Children.Append("brief", types.YChild{"Brief", &notifHistory.Brief})
    notifHistory.EntityData.Children.Append("detail", types.YChild{"Detail", &notifHistory.Detail})
    notifHistory.EntityData.Leafs = types.NewOrderedMap()

    notifHistory.EntityData.YListKeys = []string {}

    return &(notifHistory.EntityData)
}

// Controller_CardMgr_NotifHistory_Brief
type Controller_CardMgr_NotifHistory_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_NotifHistory_Brief_Location.
    Location []*Controller_CardMgr_NotifHistory_Brief_Location
}

func (brief *Controller_CardMgr_NotifHistory_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "notif-history"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/notif-history/" + brief.EntityData.SegmentPath
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range brief.Location {
        brief.EntityData.Children.Append(types.GetSegmentPath(brief.Location[i]), types.YChild{"Location", brief.Location[i]})
    }
    brief.EntityData.Leafs = types.NewOrderedMap()

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// Controller_CardMgr_NotifHistory_Brief_Location
type Controller_CardMgr_NotifHistory_Brief_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    
    CardNotifHistBrief Controller_CardMgr_NotifHistory_Brief_Location_CardNotifHistBrief
}

func (location *Controller_CardMgr_NotifHistory_Brief_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "brief"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/notif-history/brief/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("card_notif_hist_brief", types.YChild{"CardNotifHistBrief", &location.CardNotifHistBrief})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Controller_CardMgr_NotifHistory_Brief_Location_CardNotifHistBrief
type Controller_CardMgr_NotifHistory_Brief_Location_CardNotifHistBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    CardNotifHistBriefValues []interface{}
}

func (cardNotifHistBrief *Controller_CardMgr_NotifHistory_Brief_Location_CardNotifHistBrief) GetEntityData() *types.CommonEntityData {
    cardNotifHistBrief.EntityData.YFilter = cardNotifHistBrief.YFilter
    cardNotifHistBrief.EntityData.YangName = "card_notif_hist_brief"
    cardNotifHistBrief.EntityData.BundleName = "cisco_ios_xr"
    cardNotifHistBrief.EntityData.ParentYangName = "location"
    cardNotifHistBrief.EntityData.SegmentPath = "card_notif_hist_brief"
    cardNotifHistBrief.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/notif-history/brief/location/" + cardNotifHistBrief.EntityData.SegmentPath
    cardNotifHistBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardNotifHistBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardNotifHistBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardNotifHistBrief.EntityData.Children = types.NewOrderedMap()
    cardNotifHistBrief.EntityData.Leafs = types.NewOrderedMap()
    cardNotifHistBrief.EntityData.Leafs.Append("card_notif_hist_brief_values", types.YLeaf{"CardNotifHistBriefValues", cardNotifHistBrief.CardNotifHistBriefValues})

    cardNotifHistBrief.EntityData.YListKeys = []string {}

    return &(cardNotifHistBrief.EntityData)
}

// Controller_CardMgr_NotifHistory_Detail
type Controller_CardMgr_NotifHistory_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_NotifHistory_Detail_Location.
    Location []*Controller_CardMgr_NotifHistory_Detail_Location
}

func (detail *Controller_CardMgr_NotifHistory_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "notif-history"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/notif-history/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range detail.Location {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.Location[i]), types.YChild{"Location", detail.Location[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Controller_CardMgr_NotifHistory_Detail_Location
type Controller_CardMgr_NotifHistory_Detail_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    
    CardNotifHistDetail Controller_CardMgr_NotifHistory_Detail_Location_CardNotifHistDetail
}

func (location *Controller_CardMgr_NotifHistory_Detail_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "detail"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/notif-history/detail/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("card_notif_hist_detail", types.YChild{"CardNotifHistDetail", &location.CardNotifHistDetail})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

    return &(location.EntityData)
}

// Controller_CardMgr_NotifHistory_Detail_Location_CardNotifHistDetail
type Controller_CardMgr_NotifHistory_Detail_Location_CardNotifHistDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    CardNotifHistDetailValues []interface{}
}

func (cardNotifHistDetail *Controller_CardMgr_NotifHistory_Detail_Location_CardNotifHistDetail) GetEntityData() *types.CommonEntityData {
    cardNotifHistDetail.EntityData.YFilter = cardNotifHistDetail.YFilter
    cardNotifHistDetail.EntityData.YangName = "card_notif_hist_detail"
    cardNotifHistDetail.EntityData.BundleName = "cisco_ios_xr"
    cardNotifHistDetail.EntityData.ParentYangName = "location"
    cardNotifHistDetail.EntityData.SegmentPath = "card_notif_hist_detail"
    cardNotifHistDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/notif-history/detail/location/" + cardNotifHistDetail.EntityData.SegmentPath
    cardNotifHistDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardNotifHistDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardNotifHistDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardNotifHistDetail.EntityData.Children = types.NewOrderedMap()
    cardNotifHistDetail.EntityData.Leafs = types.NewOrderedMap()
    cardNotifHistDetail.EntityData.Leafs.Append("card_notif_hist_detail_values", types.YLeaf{"CardNotifHistDetailValues", cardNotifHistDetail.CardNotifHistDetailValues})

    cardNotifHistDetail.EntityData.YListKeys = []string {}

    return &(cardNotifHistDetail.EntityData)
}

// Controller_CardMgr_OirHistory
type Controller_CardMgr_OirHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_OirHistory_Rack.
    Rack []*Controller_CardMgr_OirHistory_Rack
}

func (oirHistory *Controller_CardMgr_OirHistory) GetEntityData() *types.CommonEntityData {
    oirHistory.EntityData.YFilter = oirHistory.YFilter
    oirHistory.EntityData.YangName = "oir-history"
    oirHistory.EntityData.BundleName = "cisco_ios_xr"
    oirHistory.EntityData.ParentYangName = "card_mgr"
    oirHistory.EntityData.SegmentPath = "oir-history"
    oirHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/" + oirHistory.EntityData.SegmentPath
    oirHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirHistory.EntityData.Children = types.NewOrderedMap()
    oirHistory.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range oirHistory.Rack {
        oirHistory.EntityData.Children.Append(types.GetSegmentPath(oirHistory.Rack[i]), types.YChild{"Rack", oirHistory.Rack[i]})
    }
    oirHistory.EntityData.Leafs = types.NewOrderedMap()

    oirHistory.EntityData.YListKeys = []string {}

    return &(oirHistory.EntityData)
}

// Controller_CardMgr_OirHistory_Rack
type Controller_CardMgr_OirHistory_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Rack interface{}

    
    CardOirHist Controller_CardMgr_OirHistory_Rack_CardOirHist
}

func (rack *Controller_CardMgr_OirHistory_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "oir-history"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Rack, "rack")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/oir-history/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("card_oir_hist", types.YChild{"CardOirHist", &rack.CardOirHist})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", rack.Rack})

    rack.EntityData.YListKeys = []string {"Rack"}

    return &(rack.EntityData)
}

// Controller_CardMgr_OirHistory_Rack_CardOirHist
type Controller_CardMgr_OirHistory_Rack_CardOirHist struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    CardOirEvents []interface{}
}

func (cardOirHist *Controller_CardMgr_OirHistory_Rack_CardOirHist) GetEntityData() *types.CommonEntityData {
    cardOirHist.EntityData.YFilter = cardOirHist.YFilter
    cardOirHist.EntityData.YangName = "card_oir_hist"
    cardOirHist.EntityData.BundleName = "cisco_ios_xr"
    cardOirHist.EntityData.ParentYangName = "rack"
    cardOirHist.EntityData.SegmentPath = "card_oir_hist"
    cardOirHist.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/oir-history/rack/" + cardOirHist.EntityData.SegmentPath
    cardOirHist.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardOirHist.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardOirHist.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardOirHist.EntityData.Children = types.NewOrderedMap()
    cardOirHist.EntityData.Leafs = types.NewOrderedMap()
    cardOirHist.EntityData.Leafs.Append("card_oir_events", types.YLeaf{"CardOirEvents", cardOirHist.CardOirEvents})

    cardOirHist.EntityData.YListKeys = []string {}

    return &(cardOirHist.EntityData)
}

// Controller_CardMgr_Iofpga
type Controller_CardMgr_Iofpga struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Register Controller_CardMgr_Iofpga_Register

    
    Flash Controller_CardMgr_Iofpga_Flash
}

func (iofpga *Controller_CardMgr_Iofpga) GetEntityData() *types.CommonEntityData {
    iofpga.EntityData.YFilter = iofpga.YFilter
    iofpga.EntityData.YangName = "iofpga"
    iofpga.EntityData.BundleName = "cisco_ios_xr"
    iofpga.EntityData.ParentYangName = "card_mgr"
    iofpga.EntityData.SegmentPath = "iofpga"
    iofpga.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/" + iofpga.EntityData.SegmentPath
    iofpga.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpga.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpga.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpga.EntityData.Children = types.NewOrderedMap()
    iofpga.EntityData.Children.Append("register", types.YChild{"Register", &iofpga.Register})
    iofpga.EntityData.Children.Append("flash", types.YChild{"Flash", &iofpga.Flash})
    iofpga.EntityData.Leafs = types.NewOrderedMap()

    iofpga.EntityData.YListKeys = []string {}

    return &(iofpga.EntityData)
}

// Controller_CardMgr_Iofpga_Register
type Controller_CardMgr_Iofpga_Register struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cpu Controller_CardMgr_Iofpga_Register_Cpu

    
    Mb Controller_CardMgr_Iofpga_Register_Mb

    
    Dc Controller_CardMgr_Iofpga_Register_Dc
}

func (register *Controller_CardMgr_Iofpga_Register) GetEntityData() *types.CommonEntityData {
    register.EntityData.YFilter = register.YFilter
    register.EntityData.YangName = "register"
    register.EntityData.BundleName = "cisco_ios_xr"
    register.EntityData.ParentYangName = "iofpga"
    register.EntityData.SegmentPath = "register"
    register.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/" + register.EntityData.SegmentPath
    register.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    register.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    register.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    register.EntityData.Children = types.NewOrderedMap()
    register.EntityData.Children.Append("cpu", types.YChild{"Cpu", &register.Cpu})
    register.EntityData.Children.Append("mb", types.YChild{"Mb", &register.Mb})
    register.EntityData.Children.Append("dc", types.YChild{"Dc", &register.Dc})
    register.EntityData.Leafs = types.NewOrderedMap()

    register.EntityData.YListKeys = []string {}

    return &(register.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu
type Controller_CardMgr_Iofpga_Register_Cpu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation.
    RegisterLocation []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation
}

func (cpu *Controller_CardMgr_Iofpga_Register_Cpu) GetEntityData() *types.CommonEntityData {
    cpu.EntityData.YFilter = cpu.YFilter
    cpu.EntityData.YangName = "cpu"
    cpu.EntityData.BundleName = "cisco_ios_xr"
    cpu.EntityData.ParentYangName = "register"
    cpu.EntityData.SegmentPath = "cpu"
    cpu.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/" + cpu.EntityData.SegmentPath
    cpu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpu.EntityData.Children = types.NewOrderedMap()
    cpu.EntityData.Children.Append("register_location", types.YChild{"RegisterLocation", nil})
    for i := range cpu.RegisterLocation {
        cpu.EntityData.Children.Append(types.GetSegmentPath(cpu.RegisterLocation[i]), types.YChild{"RegisterLocation", cpu.RegisterLocation[i]})
    }
    cpu.EntityData.Leafs = types.NewOrderedMap()

    cpu.EntityData.YListKeys = []string {}

    return &(cpu.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    RegisterLocation interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber.
    IofpgaBlockNumber []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset.
    IofpgaOffset []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress.
    IofpgaAddress []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress
}

func (registerLocation *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation) GetEntityData() *types.CommonEntityData {
    registerLocation.EntityData.YFilter = registerLocation.YFilter
    registerLocation.EntityData.YangName = "register_location"
    registerLocation.EntityData.BundleName = "cisco_ios_xr"
    registerLocation.EntityData.ParentYangName = "cpu"
    registerLocation.EntityData.SegmentPath = "register_location" + types.AddKeyToken(registerLocation.RegisterLocation, "register_location")
    registerLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/" + registerLocation.EntityData.SegmentPath
    registerLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registerLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registerLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registerLocation.EntityData.Children = types.NewOrderedMap()
    registerLocation.EntityData.Children.Append("iofpga_block_number", types.YChild{"IofpgaBlockNumber", nil})
    for i := range registerLocation.IofpgaBlockNumber {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaBlockNumber[i]), types.YChild{"IofpgaBlockNumber", registerLocation.IofpgaBlockNumber[i]})
    }
    registerLocation.EntityData.Children.Append("iofpga_offset", types.YChild{"IofpgaOffset", nil})
    for i := range registerLocation.IofpgaOffset {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaOffset[i]), types.YChild{"IofpgaOffset", registerLocation.IofpgaOffset[i]})
    }
    registerLocation.EntityData.Children.Append("iofpga_address", types.YChild{"IofpgaAddress", nil})
    for i := range registerLocation.IofpgaAddress {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaAddress[i]), types.YChild{"IofpgaAddress", registerLocation.IofpgaAddress[i]})
    }
    registerLocation.EntityData.Leafs = types.NewOrderedMap()
    registerLocation.EntityData.Leafs.Append("register_location", types.YLeaf{"RegisterLocation", registerLocation.RegisterLocation})

    registerLocation.EntityData.YListKeys = []string {"RegisterLocation"}

    return &(registerLocation.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    IofpgaBlockNum interface{}

    // The type is string.
    BlockLocation interface{}

    // The type is string.
    IofpgaBlockNm interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber.
    IofpgaRegisterNumber []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber
}

func (iofpgaBlockNumber *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber) GetEntityData() *types.CommonEntityData {
    iofpgaBlockNumber.EntityData.YFilter = iofpgaBlockNumber.YFilter
    iofpgaBlockNumber.EntityData.YangName = "iofpga_block_number"
    iofpgaBlockNumber.EntityData.BundleName = "cisco_ios_xr"
    iofpgaBlockNumber.EntityData.ParentYangName = "register_location"
    iofpgaBlockNumber.EntityData.SegmentPath = "iofpga_block_number" + types.AddKeyToken(iofpgaBlockNumber.IofpgaBlockNum, "iofpga_block_num")
    iofpgaBlockNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/" + iofpgaBlockNumber.EntityData.SegmentPath
    iofpgaBlockNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaBlockNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaBlockNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaBlockNumber.EntityData.Children = types.NewOrderedMap()
    iofpgaBlockNumber.EntityData.Children.Append("iofpga_register_number", types.YChild{"IofpgaRegisterNumber", nil})
    for i := range iofpgaBlockNumber.IofpgaRegisterNumber {
        iofpgaBlockNumber.EntityData.Children.Append(types.GetSegmentPath(iofpgaBlockNumber.IofpgaRegisterNumber[i]), types.YChild{"IofpgaRegisterNumber", iofpgaBlockNumber.IofpgaRegisterNumber[i]})
    }
    iofpgaBlockNumber.EntityData.Leafs = types.NewOrderedMap()
    iofpgaBlockNumber.EntityData.Leafs.Append("iofpga_block_num", types.YLeaf{"IofpgaBlockNum", iofpgaBlockNumber.IofpgaBlockNum})
    iofpgaBlockNumber.EntityData.Leafs.Append("block_location", types.YLeaf{"BlockLocation", iofpgaBlockNumber.BlockLocation})
    iofpgaBlockNumber.EntityData.Leafs.Append("iofpga_block_nm", types.YLeaf{"IofpgaBlockNm", iofpgaBlockNumber.IofpgaBlockNm})

    iofpgaBlockNumber.EntityData.YListKeys = []string {"IofpgaBlockNum"}

    return &(iofpgaBlockNumber.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // The type is string.
    IofpgaRegisterName interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData.
    IofpgaData []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData
}

func (iofpgaRegisterNumber *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber) GetEntityData() *types.CommonEntityData {
    iofpgaRegisterNumber.EntityData.YFilter = iofpgaRegisterNumber.YFilter
    iofpgaRegisterNumber.EntityData.YangName = "iofpga_register_number"
    iofpgaRegisterNumber.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegisterNumber.EntityData.ParentYangName = "iofpga_block_number"
    iofpgaRegisterNumber.EntityData.SegmentPath = "iofpga_register_number" + types.AddKeyToken(iofpgaRegisterNumber.Index, "index")
    iofpgaRegisterNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/iofpga_block_number/" + iofpgaRegisterNumber.EntityData.SegmentPath
    iofpgaRegisterNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegisterNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegisterNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegisterNumber.EntityData.Children = types.NewOrderedMap()
    iofpgaRegisterNumber.EntityData.Children.Append("iofpga_data", types.YChild{"IofpgaData", nil})
    for i := range iofpgaRegisterNumber.IofpgaData {
        types.SetYListKey(iofpgaRegisterNumber.IofpgaData[i], i)
        iofpgaRegisterNumber.EntityData.Children.Append(types.GetSegmentPath(iofpgaRegisterNumber.IofpgaData[i]), types.YChild{"IofpgaData", iofpgaRegisterNumber.IofpgaData[i]})
    }
    iofpgaRegisterNumber.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegisterNumber.EntityData.Leafs.Append("index", types.YLeaf{"Index", iofpgaRegisterNumber.Index})
    iofpgaRegisterNumber.EntityData.Leafs.Append("iofpga_register_name", types.YLeaf{"IofpgaRegisterName", iofpgaRegisterNumber.IofpgaRegisterName})

    iofpgaRegisterNumber.EntityData.YListKeys = []string {"Index"}

    return &(iofpgaRegisterNumber.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    Name interface{}

    // The type is interface{} with range: 0..4294967295.
    Offset interface{}

    // The type is interface{} with range: 0..4294967295.
    Value interface{}
}

func (iofpgaData *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData) GetEntityData() *types.CommonEntityData {
    iofpgaData.EntityData.YFilter = iofpgaData.YFilter
    iofpgaData.EntityData.YangName = "iofpga_data"
    iofpgaData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaData.EntityData.ParentYangName = "iofpga_register_number"
    iofpgaData.EntityData.SegmentPath = "iofpga_data" + types.AddNoKeyToken(iofpgaData)
    iofpgaData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/iofpga_block_number/iofpga_register_number/" + iofpgaData.EntityData.SegmentPath
    iofpgaData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaData.EntityData.Children = types.NewOrderedMap()
    iofpgaData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaData.EntityData.Leafs.Append("name", types.YLeaf{"Name", iofpgaData.Name})
    iofpgaData.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", iofpgaData.Offset})
    iofpgaData.EntityData.Leafs.Append("value", types.YLeaf{"Value", iofpgaData.Value})

    iofpgaData.EntityData.YListKeys = []string {}

    return &(iofpgaData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    HexOffset interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData.
    IofpgaRegOffsetData []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData
}

func (iofpgaOffset *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset) GetEntityData() *types.CommonEntityData {
    iofpgaOffset.EntityData.YFilter = iofpgaOffset.YFilter
    iofpgaOffset.EntityData.YangName = "iofpga_offset"
    iofpgaOffset.EntityData.BundleName = "cisco_ios_xr"
    iofpgaOffset.EntityData.ParentYangName = "register_location"
    iofpgaOffset.EntityData.SegmentPath = "iofpga_offset" + types.AddKeyToken(iofpgaOffset.HexOffset, "hex_offset")
    iofpgaOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/" + iofpgaOffset.EntityData.SegmentPath
    iofpgaOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaOffset.EntityData.Children = types.NewOrderedMap()
    iofpgaOffset.EntityData.Children.Append("iofpga_reg_offset_data", types.YChild{"IofpgaRegOffsetData", nil})
    for i := range iofpgaOffset.IofpgaRegOffsetData {
        types.SetYListKey(iofpgaOffset.IofpgaRegOffsetData[i], i)
        iofpgaOffset.EntityData.Children.Append(types.GetSegmentPath(iofpgaOffset.IofpgaRegOffsetData[i]), types.YChild{"IofpgaRegOffsetData", iofpgaOffset.IofpgaRegOffsetData[i]})
    }
    iofpgaOffset.EntityData.Leafs = types.NewOrderedMap()
    iofpgaOffset.EntityData.Leafs.Append("hex_offset", types.YLeaf{"HexOffset", iofpgaOffset.HexOffset})

    iofpgaOffset.EntityData.YListKeys = []string {"HexOffset"}

    return &(iofpgaOffset.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    IofpgaRegOffAddr interface{}

    // The type is string.
    RegOffValue interface{}
}

func (iofpgaRegOffsetData *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData) GetEntityData() *types.CommonEntityData {
    iofpgaRegOffsetData.EntityData.YFilter = iofpgaRegOffsetData.YFilter
    iofpgaRegOffsetData.EntityData.YangName = "iofpga_reg_offset_data"
    iofpgaRegOffsetData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegOffsetData.EntityData.ParentYangName = "iofpga_offset"
    iofpgaRegOffsetData.EntityData.SegmentPath = "iofpga_reg_offset_data" + types.AddNoKeyToken(iofpgaRegOffsetData)
    iofpgaRegOffsetData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/iofpga_offset/" + iofpgaRegOffsetData.EntityData.SegmentPath
    iofpgaRegOffsetData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegOffsetData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegOffsetData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegOffsetData.EntityData.Children = types.NewOrderedMap()
    iofpgaRegOffsetData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegOffsetData.EntityData.Leafs.Append("iofpga_reg_off_addr", types.YLeaf{"IofpgaRegOffAddr", iofpgaRegOffsetData.IofpgaRegOffAddr})
    iofpgaRegOffsetData.EntityData.Leafs.Append("reg_off_value", types.YLeaf{"RegOffValue", iofpgaRegOffsetData.RegOffValue})

    iofpgaRegOffsetData.EntityData.YListKeys = []string {}

    return &(iofpgaRegOffsetData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    StartHexAddr interface{}

    // This attribute is a key. The type is string.
    EndHexAddr interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList.
    IofpgaRegRangeAddrList []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList
}

func (iofpgaAddress *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress) GetEntityData() *types.CommonEntityData {
    iofpgaAddress.EntityData.YFilter = iofpgaAddress.YFilter
    iofpgaAddress.EntityData.YangName = "iofpga_address"
    iofpgaAddress.EntityData.BundleName = "cisco_ios_xr"
    iofpgaAddress.EntityData.ParentYangName = "register_location"
    iofpgaAddress.EntityData.SegmentPath = "iofpga_address" + types.AddKeyToken(iofpgaAddress.StartHexAddr, "start_hex_addr") + types.AddKeyToken(iofpgaAddress.EndHexAddr, "end_hex_addr")
    iofpgaAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/" + iofpgaAddress.EntityData.SegmentPath
    iofpgaAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaAddress.EntityData.Children = types.NewOrderedMap()
    iofpgaAddress.EntityData.Children.Append("iofpga_reg_range_addr_list", types.YChild{"IofpgaRegRangeAddrList", nil})
    for i := range iofpgaAddress.IofpgaRegRangeAddrList {
        iofpgaAddress.EntityData.Children.Append(types.GetSegmentPath(iofpgaAddress.IofpgaRegRangeAddrList[i]), types.YChild{"IofpgaRegRangeAddrList", iofpgaAddress.IofpgaRegRangeAddrList[i]})
    }
    iofpgaAddress.EntityData.Leafs = types.NewOrderedMap()
    iofpgaAddress.EntityData.Leafs.Append("start_hex_addr", types.YLeaf{"StartHexAddr", iofpgaAddress.StartHexAddr})
    iofpgaAddress.EntityData.Leafs.Append("end_hex_addr", types.YLeaf{"EndHexAddr", iofpgaAddress.EndHexAddr})

    iofpgaAddress.EntityData.YListKeys = []string {"StartHexAddr", "EndHexAddr"}

    return &(iofpgaAddress.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    IofpgaRegRangeAddr interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData.
    IofpgaRegData []*Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData
}

func (iofpgaRegRangeAddrList *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList) GetEntityData() *types.CommonEntityData {
    iofpgaRegRangeAddrList.EntityData.YFilter = iofpgaRegRangeAddrList.YFilter
    iofpgaRegRangeAddrList.EntityData.YangName = "iofpga_reg_range_addr_list"
    iofpgaRegRangeAddrList.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegRangeAddrList.EntityData.ParentYangName = "iofpga_address"
    iofpgaRegRangeAddrList.EntityData.SegmentPath = "iofpga_reg_range_addr_list" + types.AddKeyToken(iofpgaRegRangeAddrList.IofpgaRegRangeAddr, "iofpga_reg_range_addr")
    iofpgaRegRangeAddrList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/iofpga_address/" + iofpgaRegRangeAddrList.EntityData.SegmentPath
    iofpgaRegRangeAddrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegRangeAddrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegRangeAddrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegRangeAddrList.EntityData.Children = types.NewOrderedMap()
    iofpgaRegRangeAddrList.EntityData.Children.Append("iofpga_reg_data", types.YChild{"IofpgaRegData", nil})
    for i := range iofpgaRegRangeAddrList.IofpgaRegData {
        types.SetYListKey(iofpgaRegRangeAddrList.IofpgaRegData[i], i)
        iofpgaRegRangeAddrList.EntityData.Children.Append(types.GetSegmentPath(iofpgaRegRangeAddrList.IofpgaRegData[i]), types.YChild{"IofpgaRegData", iofpgaRegRangeAddrList.IofpgaRegData[i]})
    }
    iofpgaRegRangeAddrList.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegRangeAddrList.EntityData.Leafs.Append("iofpga_reg_range_addr", types.YLeaf{"IofpgaRegRangeAddr", iofpgaRegRangeAddrList.IofpgaRegRangeAddr})

    iofpgaRegRangeAddrList.EntityData.YListKeys = []string {"IofpgaRegRangeAddr"}

    return &(iofpgaRegRangeAddrList.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData
type Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    IofpgaRegAddr interface{}

    // The type is string.
    RegValue interface{}
}

func (iofpgaRegData *Controller_CardMgr_Iofpga_Register_Cpu_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData) GetEntityData() *types.CommonEntityData {
    iofpgaRegData.EntityData.YFilter = iofpgaRegData.YFilter
    iofpgaRegData.EntityData.YangName = "iofpga_reg_data"
    iofpgaRegData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegData.EntityData.ParentYangName = "iofpga_reg_range_addr_list"
    iofpgaRegData.EntityData.SegmentPath = "iofpga_reg_data" + types.AddNoKeyToken(iofpgaRegData)
    iofpgaRegData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/cpu/register_location/iofpga_address/iofpga_reg_range_addr_list/" + iofpgaRegData.EntityData.SegmentPath
    iofpgaRegData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegData.EntityData.Children = types.NewOrderedMap()
    iofpgaRegData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegData.EntityData.Leafs.Append("iofpga_reg_addr", types.YLeaf{"IofpgaRegAddr", iofpgaRegData.IofpgaRegAddr})
    iofpgaRegData.EntityData.Leafs.Append("reg_value", types.YLeaf{"RegValue", iofpgaRegData.RegValue})

    iofpgaRegData.EntityData.YListKeys = []string {}

    return &(iofpgaRegData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb
type Controller_CardMgr_Iofpga_Register_Mb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation.
    RegisterLocation []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation
}

func (mb *Controller_CardMgr_Iofpga_Register_Mb) GetEntityData() *types.CommonEntityData {
    mb.EntityData.YFilter = mb.YFilter
    mb.EntityData.YangName = "mb"
    mb.EntityData.BundleName = "cisco_ios_xr"
    mb.EntityData.ParentYangName = "register"
    mb.EntityData.SegmentPath = "mb"
    mb.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/" + mb.EntityData.SegmentPath
    mb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mb.EntityData.Children = types.NewOrderedMap()
    mb.EntityData.Children.Append("register_location", types.YChild{"RegisterLocation", nil})
    for i := range mb.RegisterLocation {
        mb.EntityData.Children.Append(types.GetSegmentPath(mb.RegisterLocation[i]), types.YChild{"RegisterLocation", mb.RegisterLocation[i]})
    }
    mb.EntityData.Leafs = types.NewOrderedMap()

    mb.EntityData.YListKeys = []string {}

    return &(mb.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    RegisterLocation interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber.
    IofpgaBlockNumber []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset.
    IofpgaOffset []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress.
    IofpgaAddress []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress
}

func (registerLocation *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation) GetEntityData() *types.CommonEntityData {
    registerLocation.EntityData.YFilter = registerLocation.YFilter
    registerLocation.EntityData.YangName = "register_location"
    registerLocation.EntityData.BundleName = "cisco_ios_xr"
    registerLocation.EntityData.ParentYangName = "mb"
    registerLocation.EntityData.SegmentPath = "register_location" + types.AddKeyToken(registerLocation.RegisterLocation, "register_location")
    registerLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/" + registerLocation.EntityData.SegmentPath
    registerLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registerLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registerLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registerLocation.EntityData.Children = types.NewOrderedMap()
    registerLocation.EntityData.Children.Append("iofpga_block_number", types.YChild{"IofpgaBlockNumber", nil})
    for i := range registerLocation.IofpgaBlockNumber {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaBlockNumber[i]), types.YChild{"IofpgaBlockNumber", registerLocation.IofpgaBlockNumber[i]})
    }
    registerLocation.EntityData.Children.Append("iofpga_offset", types.YChild{"IofpgaOffset", nil})
    for i := range registerLocation.IofpgaOffset {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaOffset[i]), types.YChild{"IofpgaOffset", registerLocation.IofpgaOffset[i]})
    }
    registerLocation.EntityData.Children.Append("iofpga_address", types.YChild{"IofpgaAddress", nil})
    for i := range registerLocation.IofpgaAddress {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaAddress[i]), types.YChild{"IofpgaAddress", registerLocation.IofpgaAddress[i]})
    }
    registerLocation.EntityData.Leafs = types.NewOrderedMap()
    registerLocation.EntityData.Leafs.Append("register_location", types.YLeaf{"RegisterLocation", registerLocation.RegisterLocation})

    registerLocation.EntityData.YListKeys = []string {"RegisterLocation"}

    return &(registerLocation.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    IofpgaBlockNum interface{}

    // The type is string.
    BlockLocation interface{}

    // The type is string.
    IofpgaBlockNm interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber.
    IofpgaRegisterNumber []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber
}

func (iofpgaBlockNumber *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber) GetEntityData() *types.CommonEntityData {
    iofpgaBlockNumber.EntityData.YFilter = iofpgaBlockNumber.YFilter
    iofpgaBlockNumber.EntityData.YangName = "iofpga_block_number"
    iofpgaBlockNumber.EntityData.BundleName = "cisco_ios_xr"
    iofpgaBlockNumber.EntityData.ParentYangName = "register_location"
    iofpgaBlockNumber.EntityData.SegmentPath = "iofpga_block_number" + types.AddKeyToken(iofpgaBlockNumber.IofpgaBlockNum, "iofpga_block_num")
    iofpgaBlockNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/" + iofpgaBlockNumber.EntityData.SegmentPath
    iofpgaBlockNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaBlockNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaBlockNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaBlockNumber.EntityData.Children = types.NewOrderedMap()
    iofpgaBlockNumber.EntityData.Children.Append("iofpga_register_number", types.YChild{"IofpgaRegisterNumber", nil})
    for i := range iofpgaBlockNumber.IofpgaRegisterNumber {
        iofpgaBlockNumber.EntityData.Children.Append(types.GetSegmentPath(iofpgaBlockNumber.IofpgaRegisterNumber[i]), types.YChild{"IofpgaRegisterNumber", iofpgaBlockNumber.IofpgaRegisterNumber[i]})
    }
    iofpgaBlockNumber.EntityData.Leafs = types.NewOrderedMap()
    iofpgaBlockNumber.EntityData.Leafs.Append("iofpga_block_num", types.YLeaf{"IofpgaBlockNum", iofpgaBlockNumber.IofpgaBlockNum})
    iofpgaBlockNumber.EntityData.Leafs.Append("block_location", types.YLeaf{"BlockLocation", iofpgaBlockNumber.BlockLocation})
    iofpgaBlockNumber.EntityData.Leafs.Append("iofpga_block_nm", types.YLeaf{"IofpgaBlockNm", iofpgaBlockNumber.IofpgaBlockNm})

    iofpgaBlockNumber.EntityData.YListKeys = []string {"IofpgaBlockNum"}

    return &(iofpgaBlockNumber.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // The type is string.
    IofpgaRegisterName interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData.
    IofpgaData []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData
}

func (iofpgaRegisterNumber *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber) GetEntityData() *types.CommonEntityData {
    iofpgaRegisterNumber.EntityData.YFilter = iofpgaRegisterNumber.YFilter
    iofpgaRegisterNumber.EntityData.YangName = "iofpga_register_number"
    iofpgaRegisterNumber.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegisterNumber.EntityData.ParentYangName = "iofpga_block_number"
    iofpgaRegisterNumber.EntityData.SegmentPath = "iofpga_register_number" + types.AddKeyToken(iofpgaRegisterNumber.Index, "index")
    iofpgaRegisterNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/iofpga_block_number/" + iofpgaRegisterNumber.EntityData.SegmentPath
    iofpgaRegisterNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegisterNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegisterNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegisterNumber.EntityData.Children = types.NewOrderedMap()
    iofpgaRegisterNumber.EntityData.Children.Append("iofpga_data", types.YChild{"IofpgaData", nil})
    for i := range iofpgaRegisterNumber.IofpgaData {
        types.SetYListKey(iofpgaRegisterNumber.IofpgaData[i], i)
        iofpgaRegisterNumber.EntityData.Children.Append(types.GetSegmentPath(iofpgaRegisterNumber.IofpgaData[i]), types.YChild{"IofpgaData", iofpgaRegisterNumber.IofpgaData[i]})
    }
    iofpgaRegisterNumber.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegisterNumber.EntityData.Leafs.Append("index", types.YLeaf{"Index", iofpgaRegisterNumber.Index})
    iofpgaRegisterNumber.EntityData.Leafs.Append("iofpga_register_name", types.YLeaf{"IofpgaRegisterName", iofpgaRegisterNumber.IofpgaRegisterName})

    iofpgaRegisterNumber.EntityData.YListKeys = []string {"Index"}

    return &(iofpgaRegisterNumber.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    Name interface{}

    // The type is interface{} with range: 0..4294967295.
    Offset interface{}

    // The type is interface{} with range: 0..4294967295.
    Value interface{}
}

func (iofpgaData *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData) GetEntityData() *types.CommonEntityData {
    iofpgaData.EntityData.YFilter = iofpgaData.YFilter
    iofpgaData.EntityData.YangName = "iofpga_data"
    iofpgaData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaData.EntityData.ParentYangName = "iofpga_register_number"
    iofpgaData.EntityData.SegmentPath = "iofpga_data" + types.AddNoKeyToken(iofpgaData)
    iofpgaData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/iofpga_block_number/iofpga_register_number/" + iofpgaData.EntityData.SegmentPath
    iofpgaData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaData.EntityData.Children = types.NewOrderedMap()
    iofpgaData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaData.EntityData.Leafs.Append("name", types.YLeaf{"Name", iofpgaData.Name})
    iofpgaData.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", iofpgaData.Offset})
    iofpgaData.EntityData.Leafs.Append("value", types.YLeaf{"Value", iofpgaData.Value})

    iofpgaData.EntityData.YListKeys = []string {}

    return &(iofpgaData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    HexOffset interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData.
    IofpgaRegOffsetData []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData
}

func (iofpgaOffset *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset) GetEntityData() *types.CommonEntityData {
    iofpgaOffset.EntityData.YFilter = iofpgaOffset.YFilter
    iofpgaOffset.EntityData.YangName = "iofpga_offset"
    iofpgaOffset.EntityData.BundleName = "cisco_ios_xr"
    iofpgaOffset.EntityData.ParentYangName = "register_location"
    iofpgaOffset.EntityData.SegmentPath = "iofpga_offset" + types.AddKeyToken(iofpgaOffset.HexOffset, "hex_offset")
    iofpgaOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/" + iofpgaOffset.EntityData.SegmentPath
    iofpgaOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaOffset.EntityData.Children = types.NewOrderedMap()
    iofpgaOffset.EntityData.Children.Append("iofpga_reg_offset_data", types.YChild{"IofpgaRegOffsetData", nil})
    for i := range iofpgaOffset.IofpgaRegOffsetData {
        types.SetYListKey(iofpgaOffset.IofpgaRegOffsetData[i], i)
        iofpgaOffset.EntityData.Children.Append(types.GetSegmentPath(iofpgaOffset.IofpgaRegOffsetData[i]), types.YChild{"IofpgaRegOffsetData", iofpgaOffset.IofpgaRegOffsetData[i]})
    }
    iofpgaOffset.EntityData.Leafs = types.NewOrderedMap()
    iofpgaOffset.EntityData.Leafs.Append("hex_offset", types.YLeaf{"HexOffset", iofpgaOffset.HexOffset})

    iofpgaOffset.EntityData.YListKeys = []string {"HexOffset"}

    return &(iofpgaOffset.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    IofpgaRegOffAddr interface{}

    // The type is string.
    RegOffValue interface{}
}

func (iofpgaRegOffsetData *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData) GetEntityData() *types.CommonEntityData {
    iofpgaRegOffsetData.EntityData.YFilter = iofpgaRegOffsetData.YFilter
    iofpgaRegOffsetData.EntityData.YangName = "iofpga_reg_offset_data"
    iofpgaRegOffsetData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegOffsetData.EntityData.ParentYangName = "iofpga_offset"
    iofpgaRegOffsetData.EntityData.SegmentPath = "iofpga_reg_offset_data" + types.AddNoKeyToken(iofpgaRegOffsetData)
    iofpgaRegOffsetData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/iofpga_offset/" + iofpgaRegOffsetData.EntityData.SegmentPath
    iofpgaRegOffsetData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegOffsetData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegOffsetData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegOffsetData.EntityData.Children = types.NewOrderedMap()
    iofpgaRegOffsetData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegOffsetData.EntityData.Leafs.Append("iofpga_reg_off_addr", types.YLeaf{"IofpgaRegOffAddr", iofpgaRegOffsetData.IofpgaRegOffAddr})
    iofpgaRegOffsetData.EntityData.Leafs.Append("reg_off_value", types.YLeaf{"RegOffValue", iofpgaRegOffsetData.RegOffValue})

    iofpgaRegOffsetData.EntityData.YListKeys = []string {}

    return &(iofpgaRegOffsetData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    StartHexAddr interface{}

    // This attribute is a key. The type is string.
    EndHexAddr interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList.
    IofpgaRegRangeAddrList []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList
}

func (iofpgaAddress *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress) GetEntityData() *types.CommonEntityData {
    iofpgaAddress.EntityData.YFilter = iofpgaAddress.YFilter
    iofpgaAddress.EntityData.YangName = "iofpga_address"
    iofpgaAddress.EntityData.BundleName = "cisco_ios_xr"
    iofpgaAddress.EntityData.ParentYangName = "register_location"
    iofpgaAddress.EntityData.SegmentPath = "iofpga_address" + types.AddKeyToken(iofpgaAddress.StartHexAddr, "start_hex_addr") + types.AddKeyToken(iofpgaAddress.EndHexAddr, "end_hex_addr")
    iofpgaAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/" + iofpgaAddress.EntityData.SegmentPath
    iofpgaAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaAddress.EntityData.Children = types.NewOrderedMap()
    iofpgaAddress.EntityData.Children.Append("iofpga_reg_range_addr_list", types.YChild{"IofpgaRegRangeAddrList", nil})
    for i := range iofpgaAddress.IofpgaRegRangeAddrList {
        iofpgaAddress.EntityData.Children.Append(types.GetSegmentPath(iofpgaAddress.IofpgaRegRangeAddrList[i]), types.YChild{"IofpgaRegRangeAddrList", iofpgaAddress.IofpgaRegRangeAddrList[i]})
    }
    iofpgaAddress.EntityData.Leafs = types.NewOrderedMap()
    iofpgaAddress.EntityData.Leafs.Append("start_hex_addr", types.YLeaf{"StartHexAddr", iofpgaAddress.StartHexAddr})
    iofpgaAddress.EntityData.Leafs.Append("end_hex_addr", types.YLeaf{"EndHexAddr", iofpgaAddress.EndHexAddr})

    iofpgaAddress.EntityData.YListKeys = []string {"StartHexAddr", "EndHexAddr"}

    return &(iofpgaAddress.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    IofpgaRegRangeAddr interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData.
    IofpgaRegData []*Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData
}

func (iofpgaRegRangeAddrList *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList) GetEntityData() *types.CommonEntityData {
    iofpgaRegRangeAddrList.EntityData.YFilter = iofpgaRegRangeAddrList.YFilter
    iofpgaRegRangeAddrList.EntityData.YangName = "iofpga_reg_range_addr_list"
    iofpgaRegRangeAddrList.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegRangeAddrList.EntityData.ParentYangName = "iofpga_address"
    iofpgaRegRangeAddrList.EntityData.SegmentPath = "iofpga_reg_range_addr_list" + types.AddKeyToken(iofpgaRegRangeAddrList.IofpgaRegRangeAddr, "iofpga_reg_range_addr")
    iofpgaRegRangeAddrList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/iofpga_address/" + iofpgaRegRangeAddrList.EntityData.SegmentPath
    iofpgaRegRangeAddrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegRangeAddrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegRangeAddrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegRangeAddrList.EntityData.Children = types.NewOrderedMap()
    iofpgaRegRangeAddrList.EntityData.Children.Append("iofpga_reg_data", types.YChild{"IofpgaRegData", nil})
    for i := range iofpgaRegRangeAddrList.IofpgaRegData {
        types.SetYListKey(iofpgaRegRangeAddrList.IofpgaRegData[i], i)
        iofpgaRegRangeAddrList.EntityData.Children.Append(types.GetSegmentPath(iofpgaRegRangeAddrList.IofpgaRegData[i]), types.YChild{"IofpgaRegData", iofpgaRegRangeAddrList.IofpgaRegData[i]})
    }
    iofpgaRegRangeAddrList.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegRangeAddrList.EntityData.Leafs.Append("iofpga_reg_range_addr", types.YLeaf{"IofpgaRegRangeAddr", iofpgaRegRangeAddrList.IofpgaRegRangeAddr})

    iofpgaRegRangeAddrList.EntityData.YListKeys = []string {"IofpgaRegRangeAddr"}

    return &(iofpgaRegRangeAddrList.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData
type Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    IofpgaRegAddr interface{}

    // The type is string.
    RegValue interface{}
}

func (iofpgaRegData *Controller_CardMgr_Iofpga_Register_Mb_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData) GetEntityData() *types.CommonEntityData {
    iofpgaRegData.EntityData.YFilter = iofpgaRegData.YFilter
    iofpgaRegData.EntityData.YangName = "iofpga_reg_data"
    iofpgaRegData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegData.EntityData.ParentYangName = "iofpga_reg_range_addr_list"
    iofpgaRegData.EntityData.SegmentPath = "iofpga_reg_data" + types.AddNoKeyToken(iofpgaRegData)
    iofpgaRegData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/mb/register_location/iofpga_address/iofpga_reg_range_addr_list/" + iofpgaRegData.EntityData.SegmentPath
    iofpgaRegData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegData.EntityData.Children = types.NewOrderedMap()
    iofpgaRegData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegData.EntityData.Leafs.Append("iofpga_reg_addr", types.YLeaf{"IofpgaRegAddr", iofpgaRegData.IofpgaRegAddr})
    iofpgaRegData.EntityData.Leafs.Append("reg_value", types.YLeaf{"RegValue", iofpgaRegData.RegValue})

    iofpgaRegData.EntityData.YListKeys = []string {}

    return &(iofpgaRegData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc
type Controller_CardMgr_Iofpga_Register_Dc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation.
    RegisterLocation []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation
}

func (dc *Controller_CardMgr_Iofpga_Register_Dc) GetEntityData() *types.CommonEntityData {
    dc.EntityData.YFilter = dc.YFilter
    dc.EntityData.YangName = "dc"
    dc.EntityData.BundleName = "cisco_ios_xr"
    dc.EntityData.ParentYangName = "register"
    dc.EntityData.SegmentPath = "dc"
    dc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/" + dc.EntityData.SegmentPath
    dc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dc.EntityData.Children = types.NewOrderedMap()
    dc.EntityData.Children.Append("register_location", types.YChild{"RegisterLocation", nil})
    for i := range dc.RegisterLocation {
        dc.EntityData.Children.Append(types.GetSegmentPath(dc.RegisterLocation[i]), types.YChild{"RegisterLocation", dc.RegisterLocation[i]})
    }
    dc.EntityData.Leafs = types.NewOrderedMap()

    dc.EntityData.YListKeys = []string {}

    return &(dc.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    RegisterLocation interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber.
    IofpgaBlockNumber []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset.
    IofpgaOffset []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress.
    IofpgaAddress []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress
}

func (registerLocation *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation) GetEntityData() *types.CommonEntityData {
    registerLocation.EntityData.YFilter = registerLocation.YFilter
    registerLocation.EntityData.YangName = "register_location"
    registerLocation.EntityData.BundleName = "cisco_ios_xr"
    registerLocation.EntityData.ParentYangName = "dc"
    registerLocation.EntityData.SegmentPath = "register_location" + types.AddKeyToken(registerLocation.RegisterLocation, "register_location")
    registerLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/" + registerLocation.EntityData.SegmentPath
    registerLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registerLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registerLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registerLocation.EntityData.Children = types.NewOrderedMap()
    registerLocation.EntityData.Children.Append("iofpga_block_number", types.YChild{"IofpgaBlockNumber", nil})
    for i := range registerLocation.IofpgaBlockNumber {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaBlockNumber[i]), types.YChild{"IofpgaBlockNumber", registerLocation.IofpgaBlockNumber[i]})
    }
    registerLocation.EntityData.Children.Append("iofpga_offset", types.YChild{"IofpgaOffset", nil})
    for i := range registerLocation.IofpgaOffset {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaOffset[i]), types.YChild{"IofpgaOffset", registerLocation.IofpgaOffset[i]})
    }
    registerLocation.EntityData.Children.Append("iofpga_address", types.YChild{"IofpgaAddress", nil})
    for i := range registerLocation.IofpgaAddress {
        registerLocation.EntityData.Children.Append(types.GetSegmentPath(registerLocation.IofpgaAddress[i]), types.YChild{"IofpgaAddress", registerLocation.IofpgaAddress[i]})
    }
    registerLocation.EntityData.Leafs = types.NewOrderedMap()
    registerLocation.EntityData.Leafs.Append("register_location", types.YLeaf{"RegisterLocation", registerLocation.RegisterLocation})

    registerLocation.EntityData.YListKeys = []string {"RegisterLocation"}

    return &(registerLocation.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    IofpgaBlockNum interface{}

    // The type is string.
    BlockLocation interface{}

    // The type is string.
    IofpgaBlockNm interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber.
    IofpgaRegisterNumber []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber
}

func (iofpgaBlockNumber *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber) GetEntityData() *types.CommonEntityData {
    iofpgaBlockNumber.EntityData.YFilter = iofpgaBlockNumber.YFilter
    iofpgaBlockNumber.EntityData.YangName = "iofpga_block_number"
    iofpgaBlockNumber.EntityData.BundleName = "cisco_ios_xr"
    iofpgaBlockNumber.EntityData.ParentYangName = "register_location"
    iofpgaBlockNumber.EntityData.SegmentPath = "iofpga_block_number" + types.AddKeyToken(iofpgaBlockNumber.IofpgaBlockNum, "iofpga_block_num")
    iofpgaBlockNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/" + iofpgaBlockNumber.EntityData.SegmentPath
    iofpgaBlockNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaBlockNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaBlockNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaBlockNumber.EntityData.Children = types.NewOrderedMap()
    iofpgaBlockNumber.EntityData.Children.Append("iofpga_register_number", types.YChild{"IofpgaRegisterNumber", nil})
    for i := range iofpgaBlockNumber.IofpgaRegisterNumber {
        iofpgaBlockNumber.EntityData.Children.Append(types.GetSegmentPath(iofpgaBlockNumber.IofpgaRegisterNumber[i]), types.YChild{"IofpgaRegisterNumber", iofpgaBlockNumber.IofpgaRegisterNumber[i]})
    }
    iofpgaBlockNumber.EntityData.Leafs = types.NewOrderedMap()
    iofpgaBlockNumber.EntityData.Leafs.Append("iofpga_block_num", types.YLeaf{"IofpgaBlockNum", iofpgaBlockNumber.IofpgaBlockNum})
    iofpgaBlockNumber.EntityData.Leafs.Append("block_location", types.YLeaf{"BlockLocation", iofpgaBlockNumber.BlockLocation})
    iofpgaBlockNumber.EntityData.Leafs.Append("iofpga_block_nm", types.YLeaf{"IofpgaBlockNm", iofpgaBlockNumber.IofpgaBlockNm})

    iofpgaBlockNumber.EntityData.YListKeys = []string {"IofpgaBlockNum"}

    return &(iofpgaBlockNumber.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // The type is string.
    IofpgaRegisterName interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData.
    IofpgaData []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData
}

func (iofpgaRegisterNumber *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber) GetEntityData() *types.CommonEntityData {
    iofpgaRegisterNumber.EntityData.YFilter = iofpgaRegisterNumber.YFilter
    iofpgaRegisterNumber.EntityData.YangName = "iofpga_register_number"
    iofpgaRegisterNumber.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegisterNumber.EntityData.ParentYangName = "iofpga_block_number"
    iofpgaRegisterNumber.EntityData.SegmentPath = "iofpga_register_number" + types.AddKeyToken(iofpgaRegisterNumber.Index, "index")
    iofpgaRegisterNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/iofpga_block_number/" + iofpgaRegisterNumber.EntityData.SegmentPath
    iofpgaRegisterNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegisterNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegisterNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegisterNumber.EntityData.Children = types.NewOrderedMap()
    iofpgaRegisterNumber.EntityData.Children.Append("iofpga_data", types.YChild{"IofpgaData", nil})
    for i := range iofpgaRegisterNumber.IofpgaData {
        types.SetYListKey(iofpgaRegisterNumber.IofpgaData[i], i)
        iofpgaRegisterNumber.EntityData.Children.Append(types.GetSegmentPath(iofpgaRegisterNumber.IofpgaData[i]), types.YChild{"IofpgaData", iofpgaRegisterNumber.IofpgaData[i]})
    }
    iofpgaRegisterNumber.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegisterNumber.EntityData.Leafs.Append("index", types.YLeaf{"Index", iofpgaRegisterNumber.Index})
    iofpgaRegisterNumber.EntityData.Leafs.Append("iofpga_register_name", types.YLeaf{"IofpgaRegisterName", iofpgaRegisterNumber.IofpgaRegisterName})

    iofpgaRegisterNumber.EntityData.YListKeys = []string {"Index"}

    return &(iofpgaRegisterNumber.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    Name interface{}

    // The type is interface{} with range: 0..4294967295.
    Offset interface{}

    // The type is interface{} with range: 0..4294967295.
    Value interface{}
}

func (iofpgaData *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaBlockNumber_IofpgaRegisterNumber_IofpgaData) GetEntityData() *types.CommonEntityData {
    iofpgaData.EntityData.YFilter = iofpgaData.YFilter
    iofpgaData.EntityData.YangName = "iofpga_data"
    iofpgaData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaData.EntityData.ParentYangName = "iofpga_register_number"
    iofpgaData.EntityData.SegmentPath = "iofpga_data" + types.AddNoKeyToken(iofpgaData)
    iofpgaData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/iofpga_block_number/iofpga_register_number/" + iofpgaData.EntityData.SegmentPath
    iofpgaData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaData.EntityData.Children = types.NewOrderedMap()
    iofpgaData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaData.EntityData.Leafs.Append("name", types.YLeaf{"Name", iofpgaData.Name})
    iofpgaData.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", iofpgaData.Offset})
    iofpgaData.EntityData.Leafs.Append("value", types.YLeaf{"Value", iofpgaData.Value})

    iofpgaData.EntityData.YListKeys = []string {}

    return &(iofpgaData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    HexOffset interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData.
    IofpgaRegOffsetData []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData
}

func (iofpgaOffset *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset) GetEntityData() *types.CommonEntityData {
    iofpgaOffset.EntityData.YFilter = iofpgaOffset.YFilter
    iofpgaOffset.EntityData.YangName = "iofpga_offset"
    iofpgaOffset.EntityData.BundleName = "cisco_ios_xr"
    iofpgaOffset.EntityData.ParentYangName = "register_location"
    iofpgaOffset.EntityData.SegmentPath = "iofpga_offset" + types.AddKeyToken(iofpgaOffset.HexOffset, "hex_offset")
    iofpgaOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/" + iofpgaOffset.EntityData.SegmentPath
    iofpgaOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaOffset.EntityData.Children = types.NewOrderedMap()
    iofpgaOffset.EntityData.Children.Append("iofpga_reg_offset_data", types.YChild{"IofpgaRegOffsetData", nil})
    for i := range iofpgaOffset.IofpgaRegOffsetData {
        types.SetYListKey(iofpgaOffset.IofpgaRegOffsetData[i], i)
        iofpgaOffset.EntityData.Children.Append(types.GetSegmentPath(iofpgaOffset.IofpgaRegOffsetData[i]), types.YChild{"IofpgaRegOffsetData", iofpgaOffset.IofpgaRegOffsetData[i]})
    }
    iofpgaOffset.EntityData.Leafs = types.NewOrderedMap()
    iofpgaOffset.EntityData.Leafs.Append("hex_offset", types.YLeaf{"HexOffset", iofpgaOffset.HexOffset})

    iofpgaOffset.EntityData.YListKeys = []string {"HexOffset"}

    return &(iofpgaOffset.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    IofpgaRegOffAddr interface{}

    // The type is string.
    RegOffValue interface{}
}

func (iofpgaRegOffsetData *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaOffset_IofpgaRegOffsetData) GetEntityData() *types.CommonEntityData {
    iofpgaRegOffsetData.EntityData.YFilter = iofpgaRegOffsetData.YFilter
    iofpgaRegOffsetData.EntityData.YangName = "iofpga_reg_offset_data"
    iofpgaRegOffsetData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegOffsetData.EntityData.ParentYangName = "iofpga_offset"
    iofpgaRegOffsetData.EntityData.SegmentPath = "iofpga_reg_offset_data" + types.AddNoKeyToken(iofpgaRegOffsetData)
    iofpgaRegOffsetData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/iofpga_offset/" + iofpgaRegOffsetData.EntityData.SegmentPath
    iofpgaRegOffsetData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegOffsetData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegOffsetData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegOffsetData.EntityData.Children = types.NewOrderedMap()
    iofpgaRegOffsetData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegOffsetData.EntityData.Leafs.Append("iofpga_reg_off_addr", types.YLeaf{"IofpgaRegOffAddr", iofpgaRegOffsetData.IofpgaRegOffAddr})
    iofpgaRegOffsetData.EntityData.Leafs.Append("reg_off_value", types.YLeaf{"RegOffValue", iofpgaRegOffsetData.RegOffValue})

    iofpgaRegOffsetData.EntityData.YListKeys = []string {}

    return &(iofpgaRegOffsetData.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    StartHexAddr interface{}

    // This attribute is a key. The type is string.
    EndHexAddr interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList.
    IofpgaRegRangeAddrList []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList
}

func (iofpgaAddress *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress) GetEntityData() *types.CommonEntityData {
    iofpgaAddress.EntityData.YFilter = iofpgaAddress.YFilter
    iofpgaAddress.EntityData.YangName = "iofpga_address"
    iofpgaAddress.EntityData.BundleName = "cisco_ios_xr"
    iofpgaAddress.EntityData.ParentYangName = "register_location"
    iofpgaAddress.EntityData.SegmentPath = "iofpga_address" + types.AddKeyToken(iofpgaAddress.StartHexAddr, "start_hex_addr") + types.AddKeyToken(iofpgaAddress.EndHexAddr, "end_hex_addr")
    iofpgaAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/" + iofpgaAddress.EntityData.SegmentPath
    iofpgaAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaAddress.EntityData.Children = types.NewOrderedMap()
    iofpgaAddress.EntityData.Children.Append("iofpga_reg_range_addr_list", types.YChild{"IofpgaRegRangeAddrList", nil})
    for i := range iofpgaAddress.IofpgaRegRangeAddrList {
        iofpgaAddress.EntityData.Children.Append(types.GetSegmentPath(iofpgaAddress.IofpgaRegRangeAddrList[i]), types.YChild{"IofpgaRegRangeAddrList", iofpgaAddress.IofpgaRegRangeAddrList[i]})
    }
    iofpgaAddress.EntityData.Leafs = types.NewOrderedMap()
    iofpgaAddress.EntityData.Leafs.Append("start_hex_addr", types.YLeaf{"StartHexAddr", iofpgaAddress.StartHexAddr})
    iofpgaAddress.EntityData.Leafs.Append("end_hex_addr", types.YLeaf{"EndHexAddr", iofpgaAddress.EndHexAddr})

    iofpgaAddress.EntityData.YListKeys = []string {"StartHexAddr", "EndHexAddr"}

    return &(iofpgaAddress.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    IofpgaRegRangeAddr interface{}

    // The type is slice of
    // Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData.
    IofpgaRegData []*Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData
}

func (iofpgaRegRangeAddrList *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList) GetEntityData() *types.CommonEntityData {
    iofpgaRegRangeAddrList.EntityData.YFilter = iofpgaRegRangeAddrList.YFilter
    iofpgaRegRangeAddrList.EntityData.YangName = "iofpga_reg_range_addr_list"
    iofpgaRegRangeAddrList.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegRangeAddrList.EntityData.ParentYangName = "iofpga_address"
    iofpgaRegRangeAddrList.EntityData.SegmentPath = "iofpga_reg_range_addr_list" + types.AddKeyToken(iofpgaRegRangeAddrList.IofpgaRegRangeAddr, "iofpga_reg_range_addr")
    iofpgaRegRangeAddrList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/iofpga_address/" + iofpgaRegRangeAddrList.EntityData.SegmentPath
    iofpgaRegRangeAddrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegRangeAddrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegRangeAddrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegRangeAddrList.EntityData.Children = types.NewOrderedMap()
    iofpgaRegRangeAddrList.EntityData.Children.Append("iofpga_reg_data", types.YChild{"IofpgaRegData", nil})
    for i := range iofpgaRegRangeAddrList.IofpgaRegData {
        types.SetYListKey(iofpgaRegRangeAddrList.IofpgaRegData[i], i)
        iofpgaRegRangeAddrList.EntityData.Children.Append(types.GetSegmentPath(iofpgaRegRangeAddrList.IofpgaRegData[i]), types.YChild{"IofpgaRegData", iofpgaRegRangeAddrList.IofpgaRegData[i]})
    }
    iofpgaRegRangeAddrList.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegRangeAddrList.EntityData.Leafs.Append("iofpga_reg_range_addr", types.YLeaf{"IofpgaRegRangeAddr", iofpgaRegRangeAddrList.IofpgaRegRangeAddr})

    iofpgaRegRangeAddrList.EntityData.YListKeys = []string {"IofpgaRegRangeAddr"}

    return &(iofpgaRegRangeAddrList.EntityData)
}

// Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData
type Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    IofpgaRegAddr interface{}

    // The type is string.
    RegValue interface{}
}

func (iofpgaRegData *Controller_CardMgr_Iofpga_Register_Dc_RegisterLocation_IofpgaAddress_IofpgaRegRangeAddrList_IofpgaRegData) GetEntityData() *types.CommonEntityData {
    iofpgaRegData.EntityData.YFilter = iofpgaRegData.YFilter
    iofpgaRegData.EntityData.YangName = "iofpga_reg_data"
    iofpgaRegData.EntityData.BundleName = "cisco_ios_xr"
    iofpgaRegData.EntityData.ParentYangName = "iofpga_reg_range_addr_list"
    iofpgaRegData.EntityData.SegmentPath = "iofpga_reg_data" + types.AddNoKeyToken(iofpgaRegData)
    iofpgaRegData.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/register/dc/register_location/iofpga_address/iofpga_reg_range_addr_list/" + iofpgaRegData.EntityData.SegmentPath
    iofpgaRegData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaRegData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaRegData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaRegData.EntityData.Children = types.NewOrderedMap()
    iofpgaRegData.EntityData.Leafs = types.NewOrderedMap()
    iofpgaRegData.EntityData.Leafs.Append("iofpga_reg_addr", types.YLeaf{"IofpgaRegAddr", iofpgaRegData.IofpgaRegAddr})
    iofpgaRegData.EntityData.Leafs.Append("reg_value", types.YLeaf{"RegValue", iofpgaRegData.RegValue})

    iofpgaRegData.EntityData.YListKeys = []string {}

    return &(iofpgaRegData.EntityData)
}

// Controller_CardMgr_Iofpga_Flash
type Controller_CardMgr_Iofpga_Flash struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Info Controller_CardMgr_Iofpga_Flash_Info

    
    Status Controller_CardMgr_Iofpga_Flash_Status
}

func (flash *Controller_CardMgr_Iofpga_Flash) GetEntityData() *types.CommonEntityData {
    flash.EntityData.YFilter = flash.YFilter
    flash.EntityData.YangName = "flash"
    flash.EntityData.BundleName = "cisco_ios_xr"
    flash.EntityData.ParentYangName = "iofpga"
    flash.EntityData.SegmentPath = "flash"
    flash.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/" + flash.EntityData.SegmentPath
    flash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flash.EntityData.Children = types.NewOrderedMap()
    flash.EntityData.Children.Append("info", types.YChild{"Info", &flash.Info})
    flash.EntityData.Children.Append("status", types.YChild{"Status", &flash.Status})
    flash.EntityData.Leafs = types.NewOrderedMap()

    flash.EntityData.YListKeys = []string {}

    return &(flash.EntityData)
}

// Controller_CardMgr_Iofpga_Flash_Info
type Controller_CardMgr_Iofpga_Flash_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_Iofpga_Flash_Info_FlashLocation.
    FlashLocation []*Controller_CardMgr_Iofpga_Flash_Info_FlashLocation
}

func (info *Controller_CardMgr_Iofpga_Flash_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "flash"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/flash/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("flash_location", types.YChild{"FlashLocation", nil})
    for i := range info.FlashLocation {
        info.EntityData.Children.Append(types.GetSegmentPath(info.FlashLocation[i]), types.YChild{"FlashLocation", info.FlashLocation[i]})
    }
    info.EntityData.Leafs = types.NewOrderedMap()

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Controller_CardMgr_Iofpga_Flash_Info_FlashLocation
type Controller_CardMgr_Iofpga_Flash_Info_FlashLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FlashLocation interface{}

    
    IofpgaFlashInfo Controller_CardMgr_Iofpga_Flash_Info_FlashLocation_IofpgaFlashInfo
}

func (flashLocation *Controller_CardMgr_Iofpga_Flash_Info_FlashLocation) GetEntityData() *types.CommonEntityData {
    flashLocation.EntityData.YFilter = flashLocation.YFilter
    flashLocation.EntityData.YangName = "flash_location"
    flashLocation.EntityData.BundleName = "cisco_ios_xr"
    flashLocation.EntityData.ParentYangName = "info"
    flashLocation.EntityData.SegmentPath = "flash_location" + types.AddKeyToken(flashLocation.FlashLocation, "flash_location")
    flashLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/flash/info/" + flashLocation.EntityData.SegmentPath
    flashLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flashLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flashLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flashLocation.EntityData.Children = types.NewOrderedMap()
    flashLocation.EntityData.Children.Append("iofpga_flash_info", types.YChild{"IofpgaFlashInfo", &flashLocation.IofpgaFlashInfo})
    flashLocation.EntityData.Leafs = types.NewOrderedMap()
    flashLocation.EntityData.Leafs.Append("flash_location", types.YLeaf{"FlashLocation", flashLocation.FlashLocation})

    flashLocation.EntityData.YListKeys = []string {"FlashLocation"}

    return &(flashLocation.EntityData)
}

// Controller_CardMgr_Iofpga_Flash_Info_FlashLocation_IofpgaFlashInfo
type Controller_CardMgr_Iofpga_Flash_Info_FlashLocation_IofpgaFlashInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    IofpgaFlashInfoValues []interface{}
}

func (iofpgaFlashInfo *Controller_CardMgr_Iofpga_Flash_Info_FlashLocation_IofpgaFlashInfo) GetEntityData() *types.CommonEntityData {
    iofpgaFlashInfo.EntityData.YFilter = iofpgaFlashInfo.YFilter
    iofpgaFlashInfo.EntityData.YangName = "iofpga_flash_info"
    iofpgaFlashInfo.EntityData.BundleName = "cisco_ios_xr"
    iofpgaFlashInfo.EntityData.ParentYangName = "flash_location"
    iofpgaFlashInfo.EntityData.SegmentPath = "iofpga_flash_info"
    iofpgaFlashInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/flash/info/flash_location/" + iofpgaFlashInfo.EntityData.SegmentPath
    iofpgaFlashInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaFlashInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaFlashInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaFlashInfo.EntityData.Children = types.NewOrderedMap()
    iofpgaFlashInfo.EntityData.Leafs = types.NewOrderedMap()
    iofpgaFlashInfo.EntityData.Leafs.Append("iofpga_flash_info_values", types.YLeaf{"IofpgaFlashInfoValues", iofpgaFlashInfo.IofpgaFlashInfoValues})

    iofpgaFlashInfo.EntityData.YListKeys = []string {}

    return &(iofpgaFlashInfo.EntityData)
}

// Controller_CardMgr_Iofpga_Flash_Status
type Controller_CardMgr_Iofpga_Flash_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_CardMgr_Iofpga_Flash_Status_FlashLocation.
    FlashLocation []*Controller_CardMgr_Iofpga_Flash_Status_FlashLocation
}

func (status *Controller_CardMgr_Iofpga_Flash_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xr"
    status.EntityData.ParentYangName = "flash"
    status.EntityData.SegmentPath = "status"
    status.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/flash/" + status.EntityData.SegmentPath
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Children.Append("flash_location", types.YChild{"FlashLocation", nil})
    for i := range status.FlashLocation {
        status.EntityData.Children.Append(types.GetSegmentPath(status.FlashLocation[i]), types.YChild{"FlashLocation", status.FlashLocation[i]})
    }
    status.EntityData.Leafs = types.NewOrderedMap()

    status.EntityData.YListKeys = []string {}

    return &(status.EntityData)
}

// Controller_CardMgr_Iofpga_Flash_Status_FlashLocation
type Controller_CardMgr_Iofpga_Flash_Status_FlashLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FlashLocation interface{}

    
    IofpgaFlashStatus Controller_CardMgr_Iofpga_Flash_Status_FlashLocation_IofpgaFlashStatus
}

func (flashLocation *Controller_CardMgr_Iofpga_Flash_Status_FlashLocation) GetEntityData() *types.CommonEntityData {
    flashLocation.EntityData.YFilter = flashLocation.YFilter
    flashLocation.EntityData.YangName = "flash_location"
    flashLocation.EntityData.BundleName = "cisco_ios_xr"
    flashLocation.EntityData.ParentYangName = "status"
    flashLocation.EntityData.SegmentPath = "flash_location" + types.AddKeyToken(flashLocation.FlashLocation, "flash_location")
    flashLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/flash/status/" + flashLocation.EntityData.SegmentPath
    flashLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flashLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flashLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flashLocation.EntityData.Children = types.NewOrderedMap()
    flashLocation.EntityData.Children.Append("iofpga_flash_status", types.YChild{"IofpgaFlashStatus", &flashLocation.IofpgaFlashStatus})
    flashLocation.EntityData.Leafs = types.NewOrderedMap()
    flashLocation.EntityData.Leafs.Append("flash_location", types.YLeaf{"FlashLocation", flashLocation.FlashLocation})

    flashLocation.EntityData.YListKeys = []string {"FlashLocation"}

    return &(flashLocation.EntityData)
}

// Controller_CardMgr_Iofpga_Flash_Status_FlashLocation_IofpgaFlashStatus
type Controller_CardMgr_Iofpga_Flash_Status_FlashLocation_IofpgaFlashStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    IofpgaFlashStatusValues []interface{}
}

func (iofpgaFlashStatus *Controller_CardMgr_Iofpga_Flash_Status_FlashLocation_IofpgaFlashStatus) GetEntityData() *types.CommonEntityData {
    iofpgaFlashStatus.EntityData.YFilter = iofpgaFlashStatus.YFilter
    iofpgaFlashStatus.EntityData.YangName = "iofpga_flash_status"
    iofpgaFlashStatus.EntityData.BundleName = "cisco_ios_xr"
    iofpgaFlashStatus.EntityData.ParentYangName = "flash_location"
    iofpgaFlashStatus.EntityData.SegmentPath = "iofpga_flash_status"
    iofpgaFlashStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/iofpga/flash/status/flash_location/" + iofpgaFlashStatus.EntityData.SegmentPath
    iofpgaFlashStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iofpgaFlashStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iofpgaFlashStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iofpgaFlashStatus.EntityData.Children = types.NewOrderedMap()
    iofpgaFlashStatus.EntityData.Leafs = types.NewOrderedMap()
    iofpgaFlashStatus.EntityData.Leafs.Append("iofpga_flash_status_values", types.YLeaf{"IofpgaFlashStatusValues", iofpgaFlashStatus.IofpgaFlashStatusValues})

    iofpgaFlashStatus.EntityData.YListKeys = []string {}

    return &(iofpgaFlashStatus.EntityData)
}

// Controller_CardMgr_Bootloader
type Controller_CardMgr_Bootloader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Flash Controller_CardMgr_Bootloader_Flash
}

func (bootloader *Controller_CardMgr_Bootloader) GetEntityData() *types.CommonEntityData {
    bootloader.EntityData.YFilter = bootloader.YFilter
    bootloader.EntityData.YangName = "bootloader"
    bootloader.EntityData.BundleName = "cisco_ios_xr"
    bootloader.EntityData.ParentYangName = "card_mgr"
    bootloader.EntityData.SegmentPath = "bootloader"
    bootloader.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/" + bootloader.EntityData.SegmentPath
    bootloader.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootloader.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootloader.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootloader.EntityData.Children = types.NewOrderedMap()
    bootloader.EntityData.Children.Append("flash", types.YChild{"Flash", &bootloader.Flash})
    bootloader.EntityData.Leafs = types.NewOrderedMap()

    bootloader.EntityData.YListKeys = []string {}

    return &(bootloader.EntityData)
}

// Controller_CardMgr_Bootloader_Flash
type Controller_CardMgr_Bootloader_Flash struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Info Controller_CardMgr_Bootloader_Flash_Info

    
    Status Controller_CardMgr_Bootloader_Flash_Status
}

func (flash *Controller_CardMgr_Bootloader_Flash) GetEntityData() *types.CommonEntityData {
    flash.EntityData.YFilter = flash.YFilter
    flash.EntityData.YangName = "flash"
    flash.EntityData.BundleName = "cisco_ios_xr"
    flash.EntityData.ParentYangName = "bootloader"
    flash.EntityData.SegmentPath = "flash"
    flash.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/bootloader/" + flash.EntityData.SegmentPath
    flash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flash.EntityData.Children = types.NewOrderedMap()
    flash.EntityData.Children.Append("info", types.YChild{"Info", &flash.Info})
    flash.EntityData.Children.Append("status", types.YChild{"Status", &flash.Status})
    flash.EntityData.Leafs = types.NewOrderedMap()

    flash.EntityData.YListKeys = []string {}

    return &(flash.EntityData)
}

// Controller_CardMgr_Bootloader_Flash_Info
type Controller_CardMgr_Bootloader_Flash_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_CardMgr_Bootloader_Flash_Info_FlashLocation.
    FlashLocation []*Controller_CardMgr_Bootloader_Flash_Info_FlashLocation
}

func (info *Controller_CardMgr_Bootloader_Flash_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "flash"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/bootloader/flash/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("flash_location", types.YChild{"FlashLocation", nil})
    for i := range info.FlashLocation {
        info.EntityData.Children.Append(types.GetSegmentPath(info.FlashLocation[i]), types.YChild{"FlashLocation", info.FlashLocation[i]})
    }
    info.EntityData.Leafs = types.NewOrderedMap()

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Controller_CardMgr_Bootloader_Flash_Info_FlashLocation
type Controller_CardMgr_Bootloader_Flash_Info_FlashLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FlashLocation interface{}

    
    BootldrFlashInfo Controller_CardMgr_Bootloader_Flash_Info_FlashLocation_BootldrFlashInfo
}

func (flashLocation *Controller_CardMgr_Bootloader_Flash_Info_FlashLocation) GetEntityData() *types.CommonEntityData {
    flashLocation.EntityData.YFilter = flashLocation.YFilter
    flashLocation.EntityData.YangName = "flash_location"
    flashLocation.EntityData.BundleName = "cisco_ios_xr"
    flashLocation.EntityData.ParentYangName = "info"
    flashLocation.EntityData.SegmentPath = "flash_location" + types.AddKeyToken(flashLocation.FlashLocation, "flash_location")
    flashLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/bootloader/flash/info/" + flashLocation.EntityData.SegmentPath
    flashLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flashLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flashLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flashLocation.EntityData.Children = types.NewOrderedMap()
    flashLocation.EntityData.Children.Append("bootldr_flash_info", types.YChild{"BootldrFlashInfo", &flashLocation.BootldrFlashInfo})
    flashLocation.EntityData.Leafs = types.NewOrderedMap()
    flashLocation.EntityData.Leafs.Append("flash_location", types.YLeaf{"FlashLocation", flashLocation.FlashLocation})

    flashLocation.EntityData.YListKeys = []string {"FlashLocation"}

    return &(flashLocation.EntityData)
}

// Controller_CardMgr_Bootloader_Flash_Info_FlashLocation_BootldrFlashInfo
type Controller_CardMgr_Bootloader_Flash_Info_FlashLocation_BootldrFlashInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    BootldrFlashInfoValues []interface{}
}

func (bootldrFlashInfo *Controller_CardMgr_Bootloader_Flash_Info_FlashLocation_BootldrFlashInfo) GetEntityData() *types.CommonEntityData {
    bootldrFlashInfo.EntityData.YFilter = bootldrFlashInfo.YFilter
    bootldrFlashInfo.EntityData.YangName = "bootldr_flash_info"
    bootldrFlashInfo.EntityData.BundleName = "cisco_ios_xr"
    bootldrFlashInfo.EntityData.ParentYangName = "flash_location"
    bootldrFlashInfo.EntityData.SegmentPath = "bootldr_flash_info"
    bootldrFlashInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/bootloader/flash/info/flash_location/" + bootldrFlashInfo.EntityData.SegmentPath
    bootldrFlashInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootldrFlashInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootldrFlashInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootldrFlashInfo.EntityData.Children = types.NewOrderedMap()
    bootldrFlashInfo.EntityData.Leafs = types.NewOrderedMap()
    bootldrFlashInfo.EntityData.Leafs.Append("bootldr_flash_info_values", types.YLeaf{"BootldrFlashInfoValues", bootldrFlashInfo.BootldrFlashInfoValues})

    bootldrFlashInfo.EntityData.YListKeys = []string {}

    return &(bootldrFlashInfo.EntityData)
}

// Controller_CardMgr_Bootloader_Flash_Status
type Controller_CardMgr_Bootloader_Flash_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_CardMgr_Bootloader_Flash_Status_FlashLocation.
    FlashLocation []*Controller_CardMgr_Bootloader_Flash_Status_FlashLocation
}

func (status *Controller_CardMgr_Bootloader_Flash_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xr"
    status.EntityData.ParentYangName = "flash"
    status.EntityData.SegmentPath = "status"
    status.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/bootloader/flash/" + status.EntityData.SegmentPath
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Children.Append("flash_location", types.YChild{"FlashLocation", nil})
    for i := range status.FlashLocation {
        status.EntityData.Children.Append(types.GetSegmentPath(status.FlashLocation[i]), types.YChild{"FlashLocation", status.FlashLocation[i]})
    }
    status.EntityData.Leafs = types.NewOrderedMap()

    status.EntityData.YListKeys = []string {}

    return &(status.EntityData)
}

// Controller_CardMgr_Bootloader_Flash_Status_FlashLocation
type Controller_CardMgr_Bootloader_Flash_Status_FlashLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    FlashLocation interface{}

    
    BootldrFlashStatus Controller_CardMgr_Bootloader_Flash_Status_FlashLocation_BootldrFlashStatus
}

func (flashLocation *Controller_CardMgr_Bootloader_Flash_Status_FlashLocation) GetEntityData() *types.CommonEntityData {
    flashLocation.EntityData.YFilter = flashLocation.YFilter
    flashLocation.EntityData.YangName = "flash_location"
    flashLocation.EntityData.BundleName = "cisco_ios_xr"
    flashLocation.EntityData.ParentYangName = "status"
    flashLocation.EntityData.SegmentPath = "flash_location" + types.AddKeyToken(flashLocation.FlashLocation, "flash_location")
    flashLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/bootloader/flash/status/" + flashLocation.EntityData.SegmentPath
    flashLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flashLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flashLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flashLocation.EntityData.Children = types.NewOrderedMap()
    flashLocation.EntityData.Children.Append("bootldr_flash_status", types.YChild{"BootldrFlashStatus", &flashLocation.BootldrFlashStatus})
    flashLocation.EntityData.Leafs = types.NewOrderedMap()
    flashLocation.EntityData.Leafs.Append("flash_location", types.YLeaf{"FlashLocation", flashLocation.FlashLocation})

    flashLocation.EntityData.YListKeys = []string {"FlashLocation"}

    return &(flashLocation.EntityData)
}

// Controller_CardMgr_Bootloader_Flash_Status_FlashLocation_BootldrFlashStatus
type Controller_CardMgr_Bootloader_Flash_Status_FlashLocation_BootldrFlashStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    BootldrFlashStatusValues []interface{}
}

func (bootldrFlashStatus *Controller_CardMgr_Bootloader_Flash_Status_FlashLocation_BootldrFlashStatus) GetEntityData() *types.CommonEntityData {
    bootldrFlashStatus.EntityData.YFilter = bootldrFlashStatus.YFilter
    bootldrFlashStatus.EntityData.YangName = "bootldr_flash_status"
    bootldrFlashStatus.EntityData.BundleName = "cisco_ios_xr"
    bootldrFlashStatus.EntityData.ParentYangName = "flash_location"
    bootldrFlashStatus.EntityData.SegmentPath = "bootldr_flash_status"
    bootldrFlashStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-controllers-ncs5501:controller/card_mgr/bootloader/flash/status/flash_location/" + bootldrFlashStatus.EntityData.SegmentPath
    bootldrFlashStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootldrFlashStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootldrFlashStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootldrFlashStatus.EntityData.Children = types.NewOrderedMap()
    bootldrFlashStatus.EntityData.Leafs = types.NewOrderedMap()
    bootldrFlashStatus.EntityData.Leafs.Append("bootldr_flash_status_values", types.YLeaf{"BootldrFlashStatusValues", bootldrFlashStatus.BootldrFlashStatusValues})

    bootldrFlashStatus.EntityData.YListKeys = []string {}

    return &(bootldrFlashStatus.EntityData)
}

