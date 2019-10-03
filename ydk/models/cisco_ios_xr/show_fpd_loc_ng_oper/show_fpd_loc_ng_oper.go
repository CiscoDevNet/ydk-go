// This module contains a collection of YANG definitions
// for Cisco IOS-XR show-fpd-loc-ng package operational data.
// 
// This module contains definitions
// for the following management objects:
//   show-fpd: Show hw-module fpd
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package show_fpd_loc_ng_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package show_fpd_loc_ng_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-show-fpd-loc-ng-oper show-fpd}", reflect.TypeOf(ShowFpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd", reflect.TypeOf(ShowFpd{}))
}

// ShowFpd
// Show hw-module fpd
type ShowFpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // location table.
    Locations ShowFpd_Locations

    // Display fpds on all locations -show hw-module fpd.
    HwModuleFpd ShowFpd_HwModuleFpd

    // help location table.
    HelpLocations ShowFpd_HelpLocations

    // Display help-fpd -show hw-module fpd help-fpd.
    HwModuleFpdHelpFpd ShowFpd_HwModuleFpdHelpFpd

    // gets fpd package info.
    Package ShowFpd_Package

    // fpd upgradable locations.
    LocationHelp ShowFpd_LocationHelp
}

func (showFpd *ShowFpd) GetEntityData() *types.CommonEntityData {
    showFpd.EntityData.YFilter = showFpd.YFilter
    showFpd.EntityData.YangName = "show-fpd"
    showFpd.EntityData.BundleName = "cisco_ios_xr"
    showFpd.EntityData.ParentYangName = "Cisco-IOS-XR-show-fpd-loc-ng-oper"
    showFpd.EntityData.SegmentPath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd"
    showFpd.EntityData.AbsolutePath = showFpd.EntityData.SegmentPath
    showFpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    showFpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    showFpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    showFpd.EntityData.Children = types.NewOrderedMap()
    showFpd.EntityData.Children.Append("locations", types.YChild{"Locations", &showFpd.Locations})
    showFpd.EntityData.Children.Append("hw-module-fpd", types.YChild{"HwModuleFpd", &showFpd.HwModuleFpd})
    showFpd.EntityData.Children.Append("help-locations", types.YChild{"HelpLocations", &showFpd.HelpLocations})
    showFpd.EntityData.Children.Append("hw-module-fpd-help-fpd", types.YChild{"HwModuleFpdHelpFpd", &showFpd.HwModuleFpdHelpFpd})
    showFpd.EntityData.Children.Append("package", types.YChild{"Package", &showFpd.Package})
    showFpd.EntityData.Children.Append("location-help", types.YChild{"LocationHelp", &showFpd.LocationHelp})
    showFpd.EntityData.Leafs = types.NewOrderedMap()

    showFpd.EntityData.YListKeys = []string {}

    return &(showFpd.EntityData)
}

// ShowFpd_Locations
// location table
type ShowFpd_Locations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // location. The type is slice of ShowFpd_Locations_Location.
    Location []*ShowFpd_Locations_Location
}

func (locations *ShowFpd_Locations) GetEntityData() *types.CommonEntityData {
    locations.EntityData.YFilter = locations.YFilter
    locations.EntityData.YangName = "locations"
    locations.EntityData.BundleName = "cisco_ios_xr"
    locations.EntityData.ParentYangName = "show-fpd"
    locations.EntityData.SegmentPath = "locations"
    locations.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/" + locations.EntityData.SegmentPath
    locations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locations.EntityData.Children = types.NewOrderedMap()
    locations.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range locations.Location {
        locations.EntityData.Children.Append(types.GetSegmentPath(locations.Location[i]), types.YChild{"Location", locations.Location[i]})
    }
    locations.EntityData.Leafs = types.NewOrderedMap()

    locations.EntityData.YListKeys = []string {}

    return &(locations.EntityData)
}

// ShowFpd_Locations_Location
// location
type ShowFpd_Locations_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Fpd location. The type is string with length:
    // 1..32.
    LocationName interface{}

    // Display fpds on given locations.
    Fpds ShowFpd_Locations_Location_Fpds
}

func (location *ShowFpd_Locations_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "locations"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location-name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/locations/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("fpds", types.YChild{"Fpds", &location.Fpds})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// ShowFpd_Locations_Location_Fpds
// Display fpds on given locations
type ShowFpd_Locations_Location_Fpds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Display fpds detail. The type is slice of
    // ShowFpd_Locations_Location_Fpds_Fpd.
    Fpd []*ShowFpd_Locations_Location_Fpds_Fpd
}

func (fpds *ShowFpd_Locations_Location_Fpds) GetEntityData() *types.CommonEntityData {
    fpds.EntityData.YFilter = fpds.YFilter
    fpds.EntityData.YangName = "fpds"
    fpds.EntityData.BundleName = "cisco_ios_xr"
    fpds.EntityData.ParentYangName = "location"
    fpds.EntityData.SegmentPath = "fpds"
    fpds.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/locations/location/" + fpds.EntityData.SegmentPath
    fpds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpds.EntityData.Children = types.NewOrderedMap()
    fpds.EntityData.Children.Append("fpd", types.YChild{"Fpd", nil})
    for i := range fpds.Fpd {
        fpds.EntityData.Children.Append(types.GetSegmentPath(fpds.Fpd[i]), types.YChild{"Fpd", fpds.Fpd[i]})
    }
    fpds.EntityData.Leafs = types.NewOrderedMap()

    fpds.EntityData.YListKeys = []string {}

    return &(fpds.EntityData)
}

// ShowFpd_Locations_Location_Fpds_Fpd
// Display fpds detail
type ShowFpd_Locations_Location_Fpds_Fpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Fpd Name. The type is string with length: 1..32.
    FpdName interface{}

    // Either Upgrading or free used by CTC . The type is string.
    UpgradeStatus interface{}

    // fpd list with all detailes. The type is slice of
    // ShowFpd_Locations_Location_Fpds_Fpd_FpdInfoDetaile.
    FpdInfoDetaile []*ShowFpd_Locations_Location_Fpds_Fpd_FpdInfoDetaile
}

func (fpd *ShowFpd_Locations_Location_Fpds_Fpd) GetEntityData() *types.CommonEntityData {
    fpd.EntityData.YFilter = fpd.YFilter
    fpd.EntityData.YangName = "fpd"
    fpd.EntityData.BundleName = "cisco_ios_xr"
    fpd.EntityData.ParentYangName = "fpds"
    fpd.EntityData.SegmentPath = "fpd" + types.AddKeyToken(fpd.FpdName, "fpd-name")
    fpd.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/locations/location/fpds/" + fpd.EntityData.SegmentPath
    fpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd.EntityData.Children = types.NewOrderedMap()
    fpd.EntityData.Children.Append("fpd-info-detaile", types.YChild{"FpdInfoDetaile", nil})
    for i := range fpd.FpdInfoDetaile {
        types.SetYListKey(fpd.FpdInfoDetaile[i], i)
        fpd.EntityData.Children.Append(types.GetSegmentPath(fpd.FpdInfoDetaile[i]), types.YChild{"FpdInfoDetaile", fpd.FpdInfoDetaile[i]})
    }
    fpd.EntityData.Leafs = types.NewOrderedMap()
    fpd.EntityData.Leafs.Append("fpd-name", types.YLeaf{"FpdName", fpd.FpdName})
    fpd.EntityData.Leafs.Append("upgrade-status", types.YLeaf{"UpgradeStatus", fpd.UpgradeStatus})

    fpd.EntityData.YListKeys = []string {"FpdName"}

    return &(fpd.EntityData)
}

// ShowFpd_Locations_Location_Fpds_Fpd_FpdInfoDetaile
//  fpd list with all detailes
type ShowFpd_Locations_Location_Fpds_Fpd_FpdInfoDetaile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // fpd location. The type is string.
    Location interface{}

    // Name of card on which fpd is located. The type is string.
    CardName interface{}

    // fpd name. The type is string.
    FpdName interface{}

    // hadware version. The type is string.
    HwVersion interface{}

    // secure boot attribur. The type is string.
    SecureBootAttr interface{}

    // status of the fpd. The type is string.
    Status interface{}

    // image running version . The type is string.
    RunningVersion interface{}

    // image programd version. The type is string.
    ProgramdVersion interface{}
}

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpds_Fpd_FpdInfoDetaile) GetEntityData() *types.CommonEntityData {
    fpdInfoDetaile.EntityData.YFilter = fpdInfoDetaile.YFilter
    fpdInfoDetaile.EntityData.YangName = "fpd-info-detaile"
    fpdInfoDetaile.EntityData.BundleName = "cisco_ios_xr"
    fpdInfoDetaile.EntityData.ParentYangName = "fpd"
    fpdInfoDetaile.EntityData.SegmentPath = "fpd-info-detaile" + types.AddNoKeyToken(fpdInfoDetaile)
    fpdInfoDetaile.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/locations/location/fpds/fpd/" + fpdInfoDetaile.EntityData.SegmentPath
    fpdInfoDetaile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpdInfoDetaile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpdInfoDetaile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpdInfoDetaile.EntityData.Children = types.NewOrderedMap()
    fpdInfoDetaile.EntityData.Leafs = types.NewOrderedMap()
    fpdInfoDetaile.EntityData.Leafs.Append("location", types.YLeaf{"Location", fpdInfoDetaile.Location})
    fpdInfoDetaile.EntityData.Leafs.Append("card-name", types.YLeaf{"CardName", fpdInfoDetaile.CardName})
    fpdInfoDetaile.EntityData.Leafs.Append("fpd-name", types.YLeaf{"FpdName", fpdInfoDetaile.FpdName})
    fpdInfoDetaile.EntityData.Leafs.Append("hw-version", types.YLeaf{"HwVersion", fpdInfoDetaile.HwVersion})
    fpdInfoDetaile.EntityData.Leafs.Append("secure-boot-attr", types.YLeaf{"SecureBootAttr", fpdInfoDetaile.SecureBootAttr})
    fpdInfoDetaile.EntityData.Leafs.Append("status", types.YLeaf{"Status", fpdInfoDetaile.Status})
    fpdInfoDetaile.EntityData.Leafs.Append("running-version", types.YLeaf{"RunningVersion", fpdInfoDetaile.RunningVersion})
    fpdInfoDetaile.EntityData.Leafs.Append("programd-version", types.YLeaf{"ProgramdVersion", fpdInfoDetaile.ProgramdVersion})

    fpdInfoDetaile.EntityData.YListKeys = []string {}

    return &(fpdInfoDetaile.EntityData)
}

// ShowFpd_HwModuleFpd
// Display fpds on all locations -show hw-module
// fpd
type ShowFpd_HwModuleFpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Either Upgrading or free used by CTC . The type is string.
    UpgradeStatus interface{}

    // fpd list with all detailes. The type is slice of
    // ShowFpd_HwModuleFpd_FpdInfoDetaile.
    FpdInfoDetaile []*ShowFpd_HwModuleFpd_FpdInfoDetaile
}

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetEntityData() *types.CommonEntityData {
    hwModuleFpd.EntityData.YFilter = hwModuleFpd.YFilter
    hwModuleFpd.EntityData.YangName = "hw-module-fpd"
    hwModuleFpd.EntityData.BundleName = "cisco_ios_xr"
    hwModuleFpd.EntityData.ParentYangName = "show-fpd"
    hwModuleFpd.EntityData.SegmentPath = "hw-module-fpd"
    hwModuleFpd.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/" + hwModuleFpd.EntityData.SegmentPath
    hwModuleFpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModuleFpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModuleFpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModuleFpd.EntityData.Children = types.NewOrderedMap()
    hwModuleFpd.EntityData.Children.Append("fpd-info-detaile", types.YChild{"FpdInfoDetaile", nil})
    for i := range hwModuleFpd.FpdInfoDetaile {
        types.SetYListKey(hwModuleFpd.FpdInfoDetaile[i], i)
        hwModuleFpd.EntityData.Children.Append(types.GetSegmentPath(hwModuleFpd.FpdInfoDetaile[i]), types.YChild{"FpdInfoDetaile", hwModuleFpd.FpdInfoDetaile[i]})
    }
    hwModuleFpd.EntityData.Leafs = types.NewOrderedMap()
    hwModuleFpd.EntityData.Leafs.Append("upgrade-status", types.YLeaf{"UpgradeStatus", hwModuleFpd.UpgradeStatus})

    hwModuleFpd.EntityData.YListKeys = []string {}

    return &(hwModuleFpd.EntityData)
}

// ShowFpd_HwModuleFpd_FpdInfoDetaile
//  fpd list with all detailes
type ShowFpd_HwModuleFpd_FpdInfoDetaile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // fpd location. The type is string.
    Location interface{}

    // Name of card on which fpd is located. The type is string.
    CardName interface{}

    // fpd name. The type is string.
    FpdName interface{}

    // hadware version. The type is string.
    HwVersion interface{}

    // secure boot attribur. The type is string.
    SecureBootAttr interface{}

    // status of the fpd. The type is string.
    Status interface{}

    // image running version . The type is string.
    RunningVersion interface{}

    // image programd version. The type is string.
    ProgramdVersion interface{}
}

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetEntityData() *types.CommonEntityData {
    fpdInfoDetaile.EntityData.YFilter = fpdInfoDetaile.YFilter
    fpdInfoDetaile.EntityData.YangName = "fpd-info-detaile"
    fpdInfoDetaile.EntityData.BundleName = "cisco_ios_xr"
    fpdInfoDetaile.EntityData.ParentYangName = "hw-module-fpd"
    fpdInfoDetaile.EntityData.SegmentPath = "fpd-info-detaile" + types.AddNoKeyToken(fpdInfoDetaile)
    fpdInfoDetaile.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/hw-module-fpd/" + fpdInfoDetaile.EntityData.SegmentPath
    fpdInfoDetaile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpdInfoDetaile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpdInfoDetaile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpdInfoDetaile.EntityData.Children = types.NewOrderedMap()
    fpdInfoDetaile.EntityData.Leafs = types.NewOrderedMap()
    fpdInfoDetaile.EntityData.Leafs.Append("location", types.YLeaf{"Location", fpdInfoDetaile.Location})
    fpdInfoDetaile.EntityData.Leafs.Append("card-name", types.YLeaf{"CardName", fpdInfoDetaile.CardName})
    fpdInfoDetaile.EntityData.Leafs.Append("fpd-name", types.YLeaf{"FpdName", fpdInfoDetaile.FpdName})
    fpdInfoDetaile.EntityData.Leafs.Append("hw-version", types.YLeaf{"HwVersion", fpdInfoDetaile.HwVersion})
    fpdInfoDetaile.EntityData.Leafs.Append("secure-boot-attr", types.YLeaf{"SecureBootAttr", fpdInfoDetaile.SecureBootAttr})
    fpdInfoDetaile.EntityData.Leafs.Append("status", types.YLeaf{"Status", fpdInfoDetaile.Status})
    fpdInfoDetaile.EntityData.Leafs.Append("running-version", types.YLeaf{"RunningVersion", fpdInfoDetaile.RunningVersion})
    fpdInfoDetaile.EntityData.Leafs.Append("programd-version", types.YLeaf{"ProgramdVersion", fpdInfoDetaile.ProgramdVersion})

    fpdInfoDetaile.EntityData.YListKeys = []string {}

    return &(fpdInfoDetaile.EntityData)
}

// ShowFpd_HelpLocations
// help location table
type ShowFpd_HelpLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // location. The type is slice of ShowFpd_HelpLocations_HelpLocation.
    HelpLocation []*ShowFpd_HelpLocations_HelpLocation
}

func (helpLocations *ShowFpd_HelpLocations) GetEntityData() *types.CommonEntityData {
    helpLocations.EntityData.YFilter = helpLocations.YFilter
    helpLocations.EntityData.YangName = "help-locations"
    helpLocations.EntityData.BundleName = "cisco_ios_xr"
    helpLocations.EntityData.ParentYangName = "show-fpd"
    helpLocations.EntityData.SegmentPath = "help-locations"
    helpLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/" + helpLocations.EntityData.SegmentPath
    helpLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helpLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helpLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helpLocations.EntityData.Children = types.NewOrderedMap()
    helpLocations.EntityData.Children.Append("help-location", types.YChild{"HelpLocation", nil})
    for i := range helpLocations.HelpLocation {
        helpLocations.EntityData.Children.Append(types.GetSegmentPath(helpLocations.HelpLocation[i]), types.YChild{"HelpLocation", helpLocations.HelpLocation[i]})
    }
    helpLocations.EntityData.Leafs = types.NewOrderedMap()

    helpLocations.EntityData.YListKeys = []string {}

    return &(helpLocations.EntityData)
}

// ShowFpd_HelpLocations_HelpLocation
// location
type ShowFpd_HelpLocations_HelpLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Fpd location. The type is string with length:
    // 1..32.
    LocationName interface{}

    // Display fpds on given locations.
    HelpFpd ShowFpd_HelpLocations_HelpLocation_HelpFpd
}

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetEntityData() *types.CommonEntityData {
    helpLocation.EntityData.YFilter = helpLocation.YFilter
    helpLocation.EntityData.YangName = "help-location"
    helpLocation.EntityData.BundleName = "cisco_ios_xr"
    helpLocation.EntityData.ParentYangName = "help-locations"
    helpLocation.EntityData.SegmentPath = "help-location" + types.AddKeyToken(helpLocation.LocationName, "location-name")
    helpLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/help-locations/" + helpLocation.EntityData.SegmentPath
    helpLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helpLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helpLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helpLocation.EntityData.Children = types.NewOrderedMap()
    helpLocation.EntityData.Children.Append("help-fpd", types.YChild{"HelpFpd", &helpLocation.HelpFpd})
    helpLocation.EntityData.Leafs = types.NewOrderedMap()
    helpLocation.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", helpLocation.LocationName})

    helpLocation.EntityData.YListKeys = []string {"LocationName"}

    return &(helpLocation.EntityData)
}

// ShowFpd_HelpLocations_HelpLocation_HelpFpd
// Display fpds on given locations
type ShowFpd_HelpLocations_HelpLocation_HelpFpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fpd name list. The type is slice of
    // ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName.
    FpdName []*ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName
}

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetEntityData() *types.CommonEntityData {
    helpFpd.EntityData.YFilter = helpFpd.YFilter
    helpFpd.EntityData.YangName = "help-fpd"
    helpFpd.EntityData.BundleName = "cisco_ios_xr"
    helpFpd.EntityData.ParentYangName = "help-location"
    helpFpd.EntityData.SegmentPath = "help-fpd"
    helpFpd.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/help-locations/help-location/" + helpFpd.EntityData.SegmentPath
    helpFpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helpFpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helpFpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helpFpd.EntityData.Children = types.NewOrderedMap()
    helpFpd.EntityData.Children.Append("fpd-name", types.YChild{"FpdName", nil})
    for i := range helpFpd.FpdName {
        types.SetYListKey(helpFpd.FpdName[i], i)
        helpFpd.EntityData.Children.Append(types.GetSegmentPath(helpFpd.FpdName[i]), types.YChild{"FpdName", helpFpd.FpdName[i]})
    }
    helpFpd.EntityData.Leafs = types.NewOrderedMap()

    helpFpd.EntityData.YListKeys = []string {}

    return &(helpFpd.EntityData)
}

// ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName
// Fpd name list
type ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // fpd location. The type is string.
    Location interface{}

    // fpd name. The type is string.
    FpdName interface{}
}

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetEntityData() *types.CommonEntityData {
    fpdName.EntityData.YFilter = fpdName.YFilter
    fpdName.EntityData.YangName = "fpd-name"
    fpdName.EntityData.BundleName = "cisco_ios_xr"
    fpdName.EntityData.ParentYangName = "help-fpd"
    fpdName.EntityData.SegmentPath = "fpd-name" + types.AddNoKeyToken(fpdName)
    fpdName.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/help-locations/help-location/help-fpd/" + fpdName.EntityData.SegmentPath
    fpdName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpdName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpdName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpdName.EntityData.Children = types.NewOrderedMap()
    fpdName.EntityData.Leafs = types.NewOrderedMap()
    fpdName.EntityData.Leafs.Append("location", types.YLeaf{"Location", fpdName.Location})
    fpdName.EntityData.Leafs.Append("fpd-name", types.YLeaf{"FpdName", fpdName.FpdName})

    fpdName.EntityData.YListKeys = []string {}

    return &(fpdName.EntityData)
}

// ShowFpd_HwModuleFpdHelpFpd
// Display help-fpd -show hw-module fpd help-fpd
type ShowFpd_HwModuleFpdHelpFpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fpd name list. The type is slice of ShowFpd_HwModuleFpdHelpFpd_FpdName.
    FpdName []*ShowFpd_HwModuleFpdHelpFpd_FpdName
}

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetEntityData() *types.CommonEntityData {
    hwModuleFpdHelpFpd.EntityData.YFilter = hwModuleFpdHelpFpd.YFilter
    hwModuleFpdHelpFpd.EntityData.YangName = "hw-module-fpd-help-fpd"
    hwModuleFpdHelpFpd.EntityData.BundleName = "cisco_ios_xr"
    hwModuleFpdHelpFpd.EntityData.ParentYangName = "show-fpd"
    hwModuleFpdHelpFpd.EntityData.SegmentPath = "hw-module-fpd-help-fpd"
    hwModuleFpdHelpFpd.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/" + hwModuleFpdHelpFpd.EntityData.SegmentPath
    hwModuleFpdHelpFpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModuleFpdHelpFpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModuleFpdHelpFpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModuleFpdHelpFpd.EntityData.Children = types.NewOrderedMap()
    hwModuleFpdHelpFpd.EntityData.Children.Append("fpd-name", types.YChild{"FpdName", nil})
    for i := range hwModuleFpdHelpFpd.FpdName {
        types.SetYListKey(hwModuleFpdHelpFpd.FpdName[i], i)
        hwModuleFpdHelpFpd.EntityData.Children.Append(types.GetSegmentPath(hwModuleFpdHelpFpd.FpdName[i]), types.YChild{"FpdName", hwModuleFpdHelpFpd.FpdName[i]})
    }
    hwModuleFpdHelpFpd.EntityData.Leafs = types.NewOrderedMap()

    hwModuleFpdHelpFpd.EntityData.YListKeys = []string {}

    return &(hwModuleFpdHelpFpd.EntityData)
}

// ShowFpd_HwModuleFpdHelpFpd_FpdName
// Fpd name list
type ShowFpd_HwModuleFpdHelpFpd_FpdName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // fpd location. The type is string.
    Location interface{}

    // fpd name. The type is string.
    FpdName interface{}
}

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetEntityData() *types.CommonEntityData {
    fpdName.EntityData.YFilter = fpdName.YFilter
    fpdName.EntityData.YangName = "fpd-name"
    fpdName.EntityData.BundleName = "cisco_ios_xr"
    fpdName.EntityData.ParentYangName = "hw-module-fpd-help-fpd"
    fpdName.EntityData.SegmentPath = "fpd-name" + types.AddNoKeyToken(fpdName)
    fpdName.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/hw-module-fpd-help-fpd/" + fpdName.EntityData.SegmentPath
    fpdName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpdName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpdName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpdName.EntityData.Children = types.NewOrderedMap()
    fpdName.EntityData.Leafs = types.NewOrderedMap()
    fpdName.EntityData.Leafs.Append("location", types.YLeaf{"Location", fpdName.Location})
    fpdName.EntityData.Leafs.Append("fpd-name", types.YLeaf{"FpdName", fpdName.FpdName})

    fpdName.EntityData.YListKeys = []string {}

    return &(fpdName.EntityData)
}

// ShowFpd_Package
// gets fpd package info
type ShowFpd_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // fpd pkg list . The type is slice of ShowFpd_Package_FpdPkgData.
    FpdPkgData []*ShowFpd_Package_FpdPkgData
}

func (self *ShowFpd_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "show-fpd"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("fpd-pkg-data", types.YChild{"FpdPkgData", nil})
    for i := range self.FpdPkgData {
        types.SetYListKey(self.FpdPkgData[i], i)
        self.EntityData.Children.Append(types.GetSegmentPath(self.FpdPkgData[i]), types.YChild{"FpdPkgData", self.FpdPkgData[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ShowFpd_Package_FpdPkgData
//  fpd pkg list 
type ShowFpd_Package_FpdPkgData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // card type. The type is string.
    CardType interface{}

    // fpd desc. The type is string.
    FpdDesc interface{}

    // reload or not. The type is string.
    UpgradeMethod interface{}

    // fpd version. The type is string.
    FpdVer interface{}

    // minimum sw version. The type is string.
    MinSwVer interface{}

    // minimum hw version. The type is string.
    MinHwVer interface{}
}

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetEntityData() *types.CommonEntityData {
    fpdPkgData.EntityData.YFilter = fpdPkgData.YFilter
    fpdPkgData.EntityData.YangName = "fpd-pkg-data"
    fpdPkgData.EntityData.BundleName = "cisco_ios_xr"
    fpdPkgData.EntityData.ParentYangName = "package"
    fpdPkgData.EntityData.SegmentPath = "fpd-pkg-data" + types.AddNoKeyToken(fpdPkgData)
    fpdPkgData.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/package/" + fpdPkgData.EntityData.SegmentPath
    fpdPkgData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpdPkgData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpdPkgData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpdPkgData.EntityData.Children = types.NewOrderedMap()
    fpdPkgData.EntityData.Leafs = types.NewOrderedMap()
    fpdPkgData.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", fpdPkgData.CardType})
    fpdPkgData.EntityData.Leafs.Append("fpd-desc", types.YLeaf{"FpdDesc", fpdPkgData.FpdDesc})
    fpdPkgData.EntityData.Leafs.Append("upgrade-method", types.YLeaf{"UpgradeMethod", fpdPkgData.UpgradeMethod})
    fpdPkgData.EntityData.Leafs.Append("fpd-ver", types.YLeaf{"FpdVer", fpdPkgData.FpdVer})
    fpdPkgData.EntityData.Leafs.Append("min-sw-ver", types.YLeaf{"MinSwVer", fpdPkgData.MinSwVer})
    fpdPkgData.EntityData.Leafs.Append("min-hw-ver", types.YLeaf{"MinHwVer", fpdPkgData.MinHwVer})

    fpdPkgData.EntityData.YListKeys = []string {}

    return &(fpdPkgData.EntityData)
}

// ShowFpd_LocationHelp
// fpd upgradable locations
type ShowFpd_LocationHelp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card location list. The type is slice of ShowFpd_LocationHelp_LocationName.
    LocationName []*ShowFpd_LocationHelp_LocationName
}

func (locationHelp *ShowFpd_LocationHelp) GetEntityData() *types.CommonEntityData {
    locationHelp.EntityData.YFilter = locationHelp.YFilter
    locationHelp.EntityData.YangName = "location-help"
    locationHelp.EntityData.BundleName = "cisco_ios_xr"
    locationHelp.EntityData.ParentYangName = "show-fpd"
    locationHelp.EntityData.SegmentPath = "location-help"
    locationHelp.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/" + locationHelp.EntityData.SegmentPath
    locationHelp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationHelp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationHelp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationHelp.EntityData.Children = types.NewOrderedMap()
    locationHelp.EntityData.Children.Append("location-name", types.YChild{"LocationName", nil})
    for i := range locationHelp.LocationName {
        types.SetYListKey(locationHelp.LocationName[i], i)
        locationHelp.EntityData.Children.Append(types.GetSegmentPath(locationHelp.LocationName[i]), types.YChild{"LocationName", locationHelp.LocationName[i]})
    }
    locationHelp.EntityData.Leafs = types.NewOrderedMap()

    locationHelp.EntityData.YListKeys = []string {}

    return &(locationHelp.EntityData)
}

// ShowFpd_LocationHelp_LocationName
// card location list
type ShowFpd_LocationHelp_LocationName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // card location. The type is string.
    LocationName interface{}
}

func (locationName *ShowFpd_LocationHelp_LocationName) GetEntityData() *types.CommonEntityData {
    locationName.EntityData.YFilter = locationName.YFilter
    locationName.EntityData.YangName = "location-name"
    locationName.EntityData.BundleName = "cisco_ios_xr"
    locationName.EntityData.ParentYangName = "location-help"
    locationName.EntityData.SegmentPath = "location-name" + types.AddNoKeyToken(locationName)
    locationName.EntityData.AbsolutePath = "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd/location-help/" + locationName.EntityData.SegmentPath
    locationName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationName.EntityData.Children = types.NewOrderedMap()
    locationName.EntityData.Leafs = types.NewOrderedMap()
    locationName.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", locationName.LocationName})

    locationName.EntityData.YListKeys = []string {}

    return &(locationName.EntityData)
}

