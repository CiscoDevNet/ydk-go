// This module contains a collection of YANG definitions
// for Cisco IOS-XR show-fpd-loc-ng package operational data.
// 
// This module contains definitions
// for the following management objects:
//   show-fpd: Show hw-module fpd
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
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

func (showFpd *ShowFpd) GetFilter() yfilter.YFilter { return showFpd.YFilter }

func (showFpd *ShowFpd) SetFilter(yf yfilter.YFilter) { showFpd.YFilter = yf }

func (showFpd *ShowFpd) GetGoName(yname string) string {
    if yname == "locations" { return "Locations" }
    if yname == "hw-module-fpd" { return "HwModuleFpd" }
    if yname == "help-locations" { return "HelpLocations" }
    if yname == "hw-module-fpd-help-fpd" { return "HwModuleFpdHelpFpd" }
    if yname == "package" { return "Package" }
    if yname == "location-help" { return "LocationHelp" }
    return ""
}

func (showFpd *ShowFpd) GetSegmentPath() string {
    return "Cisco-IOS-XR-show-fpd-loc-ng-oper:show-fpd"
}

func (showFpd *ShowFpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "locations" {
        return &showFpd.Locations
    }
    if childYangName == "hw-module-fpd" {
        return &showFpd.HwModuleFpd
    }
    if childYangName == "help-locations" {
        return &showFpd.HelpLocations
    }
    if childYangName == "hw-module-fpd-help-fpd" {
        return &showFpd.HwModuleFpdHelpFpd
    }
    if childYangName == "package" {
        return &showFpd.Package
    }
    if childYangName == "location-help" {
        return &showFpd.LocationHelp
    }
    return nil
}

func (showFpd *ShowFpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["locations"] = &showFpd.Locations
    children["hw-module-fpd"] = &showFpd.HwModuleFpd
    children["help-locations"] = &showFpd.HelpLocations
    children["hw-module-fpd-help-fpd"] = &showFpd.HwModuleFpdHelpFpd
    children["package"] = &showFpd.Package
    children["location-help"] = &showFpd.LocationHelp
    return children
}

func (showFpd *ShowFpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (showFpd *ShowFpd) GetBundleName() string { return "cisco_ios_xr" }

func (showFpd *ShowFpd) GetYangName() string { return "show-fpd" }

func (showFpd *ShowFpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (showFpd *ShowFpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (showFpd *ShowFpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (showFpd *ShowFpd) SetParent(parent types.Entity) { showFpd.parent = parent }

func (showFpd *ShowFpd) GetParent() types.Entity { return showFpd.parent }

func (showFpd *ShowFpd) GetParentYangName() string { return "Cisco-IOS-XR-show-fpd-loc-ng-oper" }

// ShowFpd_Locations
// location table
type ShowFpd_Locations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // location. The type is slice of ShowFpd_Locations_Location.
    Location []ShowFpd_Locations_Location
}

func (locations *ShowFpd_Locations) GetFilter() yfilter.YFilter { return locations.YFilter }

func (locations *ShowFpd_Locations) SetFilter(yf yfilter.YFilter) { locations.YFilter = yf }

func (locations *ShowFpd_Locations) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (locations *ShowFpd_Locations) GetSegmentPath() string {
    return "locations"
}

func (locations *ShowFpd_Locations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "location" {
        for _, c := range locations.Location {
            if locations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_Locations_Location{}
        locations.Location = append(locations.Location, child)
        return &locations.Location[len(locations.Location)-1]
    }
    return nil
}

func (locations *ShowFpd_Locations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locations.Location {
        children[locations.Location[i].GetSegmentPath()] = &locations.Location[i]
    }
    return children
}

func (locations *ShowFpd_Locations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locations *ShowFpd_Locations) GetBundleName() string { return "cisco_ios_xr" }

func (locations *ShowFpd_Locations) GetYangName() string { return "locations" }

func (locations *ShowFpd_Locations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locations *ShowFpd_Locations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locations *ShowFpd_Locations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locations *ShowFpd_Locations) SetParent(parent types.Entity) { locations.parent = parent }

func (locations *ShowFpd_Locations) GetParent() types.Entity { return locations.parent }

func (locations *ShowFpd_Locations) GetParentYangName() string { return "show-fpd" }

// ShowFpd_Locations_Location
// location
type ShowFpd_Locations_Location struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fpd location. The type is string with length:
    // 1..32.
    LocationName interface{}

    // Display fpds on given locations. The type is slice of
    // ShowFpd_Locations_Location_Fpd.
    Fpd []ShowFpd_Locations_Location_Fpd
}

func (location *ShowFpd_Locations_Location) GetFilter() yfilter.YFilter { return location.YFilter }

func (location *ShowFpd_Locations_Location) SetFilter(yf yfilter.YFilter) { location.YFilter = yf }

func (location *ShowFpd_Locations_Location) GetGoName(yname string) string {
    if yname == "location-name" { return "LocationName" }
    if yname == "fpd" { return "Fpd" }
    return ""
}

func (location *ShowFpd_Locations_Location) GetSegmentPath() string {
    return "location" + "[location-name='" + fmt.Sprintf("%v", location.LocationName) + "']"
}

func (location *ShowFpd_Locations_Location) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fpd" {
        for _, c := range location.Fpd {
            if location.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_Locations_Location_Fpd{}
        location.Fpd = append(location.Fpd, child)
        return &location.Fpd[len(location.Fpd)-1]
    }
    return nil
}

func (location *ShowFpd_Locations_Location) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range location.Fpd {
        children[location.Fpd[i].GetSegmentPath()] = &location.Fpd[i]
    }
    return children
}

func (location *ShowFpd_Locations_Location) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location-name"] = location.LocationName
    return leafs
}

func (location *ShowFpd_Locations_Location) GetBundleName() string { return "cisco_ios_xr" }

func (location *ShowFpd_Locations_Location) GetYangName() string { return "location" }

func (location *ShowFpd_Locations_Location) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (location *ShowFpd_Locations_Location) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (location *ShowFpd_Locations_Location) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (location *ShowFpd_Locations_Location) SetParent(parent types.Entity) { location.parent = parent }

func (location *ShowFpd_Locations_Location) GetParent() types.Entity { return location.parent }

func (location *ShowFpd_Locations_Location) GetParentYangName() string { return "locations" }

// ShowFpd_Locations_Location_Fpd
// Display fpds on given locations
type ShowFpd_Locations_Location_Fpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fpd Name. The type is string with length: 1..32.
    FpdName interface{}

    // fpd list with all detailes. The type is slice of
    // ShowFpd_Locations_Location_Fpd_FpdInfoDetaile.
    FpdInfoDetaile []ShowFpd_Locations_Location_Fpd_FpdInfoDetaile
}

func (fpd *ShowFpd_Locations_Location_Fpd) GetFilter() yfilter.YFilter { return fpd.YFilter }

func (fpd *ShowFpd_Locations_Location_Fpd) SetFilter(yf yfilter.YFilter) { fpd.YFilter = yf }

func (fpd *ShowFpd_Locations_Location_Fpd) GetGoName(yname string) string {
    if yname == "fpd-name" { return "FpdName" }
    if yname == "fpd-info-detaile" { return "FpdInfoDetaile" }
    return ""
}

func (fpd *ShowFpd_Locations_Location_Fpd) GetSegmentPath() string {
    return "fpd" + "[fpd-name='" + fmt.Sprintf("%v", fpd.FpdName) + "']"
}

func (fpd *ShowFpd_Locations_Location_Fpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fpd-info-detaile" {
        for _, c := range fpd.FpdInfoDetaile {
            if fpd.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_Locations_Location_Fpd_FpdInfoDetaile{}
        fpd.FpdInfoDetaile = append(fpd.FpdInfoDetaile, child)
        return &fpd.FpdInfoDetaile[len(fpd.FpdInfoDetaile)-1]
    }
    return nil
}

func (fpd *ShowFpd_Locations_Location_Fpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fpd.FpdInfoDetaile {
        children[fpd.FpdInfoDetaile[i].GetSegmentPath()] = &fpd.FpdInfoDetaile[i]
    }
    return children
}

func (fpd *ShowFpd_Locations_Location_Fpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fpd-name"] = fpd.FpdName
    return leafs
}

func (fpd *ShowFpd_Locations_Location_Fpd) GetBundleName() string { return "cisco_ios_xr" }

func (fpd *ShowFpd_Locations_Location_Fpd) GetYangName() string { return "fpd" }

func (fpd *ShowFpd_Locations_Location_Fpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fpd *ShowFpd_Locations_Location_Fpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fpd *ShowFpd_Locations_Location_Fpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fpd *ShowFpd_Locations_Location_Fpd) SetParent(parent types.Entity) { fpd.parent = parent }

func (fpd *ShowFpd_Locations_Location_Fpd) GetParent() types.Entity { return fpd.parent }

func (fpd *ShowFpd_Locations_Location_Fpd) GetParentYangName() string { return "location" }

// ShowFpd_Locations_Location_Fpd_FpdInfoDetaile
//  fpd list with all detailes
type ShowFpd_Locations_Location_Fpd_FpdInfoDetaile struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetFilter() yfilter.YFilter { return fpdInfoDetaile.YFilter }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) SetFilter(yf yfilter.YFilter) { fpdInfoDetaile.YFilter = yf }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "card-name" { return "CardName" }
    if yname == "fpd-name" { return "FpdName" }
    if yname == "hw-version" { return "HwVersion" }
    if yname == "secure-boot-attr" { return "SecureBootAttr" }
    if yname == "status" { return "Status" }
    if yname == "running-version" { return "RunningVersion" }
    if yname == "programd-version" { return "ProgramdVersion" }
    return ""
}

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetSegmentPath() string {
    return "fpd-info-detaile"
}

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = fpdInfoDetaile.Location
    leafs["card-name"] = fpdInfoDetaile.CardName
    leafs["fpd-name"] = fpdInfoDetaile.FpdName
    leafs["hw-version"] = fpdInfoDetaile.HwVersion
    leafs["secure-boot-attr"] = fpdInfoDetaile.SecureBootAttr
    leafs["status"] = fpdInfoDetaile.Status
    leafs["running-version"] = fpdInfoDetaile.RunningVersion
    leafs["programd-version"] = fpdInfoDetaile.ProgramdVersion
    return leafs
}

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetBundleName() string { return "cisco_ios_xr" }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetYangName() string { return "fpd-info-detaile" }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) SetParent(parent types.Entity) { fpdInfoDetaile.parent = parent }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetParent() types.Entity { return fpdInfoDetaile.parent }

func (fpdInfoDetaile *ShowFpd_Locations_Location_Fpd_FpdInfoDetaile) GetParentYangName() string { return "fpd" }

// ShowFpd_HwModuleFpd
// Display fpds on all locations -show hw-module
// fpd
type ShowFpd_HwModuleFpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // fpd list with all detailes. The type is slice of
    // ShowFpd_HwModuleFpd_FpdInfoDetaile.
    FpdInfoDetaile []ShowFpd_HwModuleFpd_FpdInfoDetaile
}

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetFilter() yfilter.YFilter { return hwModuleFpd.YFilter }

func (hwModuleFpd *ShowFpd_HwModuleFpd) SetFilter(yf yfilter.YFilter) { hwModuleFpd.YFilter = yf }

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetGoName(yname string) string {
    if yname == "fpd-info-detaile" { return "FpdInfoDetaile" }
    return ""
}

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetSegmentPath() string {
    return "hw-module-fpd"
}

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fpd-info-detaile" {
        for _, c := range hwModuleFpd.FpdInfoDetaile {
            if hwModuleFpd.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_HwModuleFpd_FpdInfoDetaile{}
        hwModuleFpd.FpdInfoDetaile = append(hwModuleFpd.FpdInfoDetaile, child)
        return &hwModuleFpd.FpdInfoDetaile[len(hwModuleFpd.FpdInfoDetaile)-1]
    }
    return nil
}

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwModuleFpd.FpdInfoDetaile {
        children[hwModuleFpd.FpdInfoDetaile[i].GetSegmentPath()] = &hwModuleFpd.FpdInfoDetaile[i]
    }
    return children
}

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetBundleName() string { return "cisco_ios_xr" }

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetYangName() string { return "hw-module-fpd" }

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwModuleFpd *ShowFpd_HwModuleFpd) SetParent(parent types.Entity) { hwModuleFpd.parent = parent }

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetParent() types.Entity { return hwModuleFpd.parent }

func (hwModuleFpd *ShowFpd_HwModuleFpd) GetParentYangName() string { return "show-fpd" }

// ShowFpd_HwModuleFpd_FpdInfoDetaile
//  fpd list with all detailes
type ShowFpd_HwModuleFpd_FpdInfoDetaile struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetFilter() yfilter.YFilter { return fpdInfoDetaile.YFilter }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) SetFilter(yf yfilter.YFilter) { fpdInfoDetaile.YFilter = yf }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "card-name" { return "CardName" }
    if yname == "fpd-name" { return "FpdName" }
    if yname == "hw-version" { return "HwVersion" }
    if yname == "secure-boot-attr" { return "SecureBootAttr" }
    if yname == "status" { return "Status" }
    if yname == "running-version" { return "RunningVersion" }
    if yname == "programd-version" { return "ProgramdVersion" }
    return ""
}

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetSegmentPath() string {
    return "fpd-info-detaile"
}

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = fpdInfoDetaile.Location
    leafs["card-name"] = fpdInfoDetaile.CardName
    leafs["fpd-name"] = fpdInfoDetaile.FpdName
    leafs["hw-version"] = fpdInfoDetaile.HwVersion
    leafs["secure-boot-attr"] = fpdInfoDetaile.SecureBootAttr
    leafs["status"] = fpdInfoDetaile.Status
    leafs["running-version"] = fpdInfoDetaile.RunningVersion
    leafs["programd-version"] = fpdInfoDetaile.ProgramdVersion
    return leafs
}

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetBundleName() string { return "cisco_ios_xr" }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetYangName() string { return "fpd-info-detaile" }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) SetParent(parent types.Entity) { fpdInfoDetaile.parent = parent }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetParent() types.Entity { return fpdInfoDetaile.parent }

func (fpdInfoDetaile *ShowFpd_HwModuleFpd_FpdInfoDetaile) GetParentYangName() string { return "hw-module-fpd" }

// ShowFpd_HelpLocations
// help location table
type ShowFpd_HelpLocations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // location. The type is slice of ShowFpd_HelpLocations_HelpLocation.
    HelpLocation []ShowFpd_HelpLocations_HelpLocation
}

func (helpLocations *ShowFpd_HelpLocations) GetFilter() yfilter.YFilter { return helpLocations.YFilter }

func (helpLocations *ShowFpd_HelpLocations) SetFilter(yf yfilter.YFilter) { helpLocations.YFilter = yf }

func (helpLocations *ShowFpd_HelpLocations) GetGoName(yname string) string {
    if yname == "help-location" { return "HelpLocation" }
    return ""
}

func (helpLocations *ShowFpd_HelpLocations) GetSegmentPath() string {
    return "help-locations"
}

func (helpLocations *ShowFpd_HelpLocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "help-location" {
        for _, c := range helpLocations.HelpLocation {
            if helpLocations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_HelpLocations_HelpLocation{}
        helpLocations.HelpLocation = append(helpLocations.HelpLocation, child)
        return &helpLocations.HelpLocation[len(helpLocations.HelpLocation)-1]
    }
    return nil
}

func (helpLocations *ShowFpd_HelpLocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helpLocations.HelpLocation {
        children[helpLocations.HelpLocation[i].GetSegmentPath()] = &helpLocations.HelpLocation[i]
    }
    return children
}

func (helpLocations *ShowFpd_HelpLocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helpLocations *ShowFpd_HelpLocations) GetBundleName() string { return "cisco_ios_xr" }

func (helpLocations *ShowFpd_HelpLocations) GetYangName() string { return "help-locations" }

func (helpLocations *ShowFpd_HelpLocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helpLocations *ShowFpd_HelpLocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helpLocations *ShowFpd_HelpLocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helpLocations *ShowFpd_HelpLocations) SetParent(parent types.Entity) { helpLocations.parent = parent }

func (helpLocations *ShowFpd_HelpLocations) GetParent() types.Entity { return helpLocations.parent }

func (helpLocations *ShowFpd_HelpLocations) GetParentYangName() string { return "show-fpd" }

// ShowFpd_HelpLocations_HelpLocation
// location
type ShowFpd_HelpLocations_HelpLocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fpd location. The type is string with length:
    // 1..32.
    LocationName interface{}

    // Display fpds on given locations.
    HelpFpd ShowFpd_HelpLocations_HelpLocation_HelpFpd
}

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetFilter() yfilter.YFilter { return helpLocation.YFilter }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) SetFilter(yf yfilter.YFilter) { helpLocation.YFilter = yf }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetGoName(yname string) string {
    if yname == "location-name" { return "LocationName" }
    if yname == "help-fpd" { return "HelpFpd" }
    return ""
}

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetSegmentPath() string {
    return "help-location" + "[location-name='" + fmt.Sprintf("%v", helpLocation.LocationName) + "']"
}

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "help-fpd" {
        return &helpLocation.HelpFpd
    }
    return nil
}

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["help-fpd"] = &helpLocation.HelpFpd
    return children
}

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location-name"] = helpLocation.LocationName
    return leafs
}

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetBundleName() string { return "cisco_ios_xr" }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetYangName() string { return "help-location" }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) SetParent(parent types.Entity) { helpLocation.parent = parent }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetParent() types.Entity { return helpLocation.parent }

func (helpLocation *ShowFpd_HelpLocations_HelpLocation) GetParentYangName() string { return "help-locations" }

// ShowFpd_HelpLocations_HelpLocation_HelpFpd
// Display fpds on given locations
type ShowFpd_HelpLocations_HelpLocation_HelpFpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fpd name list. The type is slice of
    // ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName.
    FpdName []ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName
}

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetFilter() yfilter.YFilter { return helpFpd.YFilter }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) SetFilter(yf yfilter.YFilter) { helpFpd.YFilter = yf }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetGoName(yname string) string {
    if yname == "fpd-name" { return "FpdName" }
    return ""
}

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetSegmentPath() string {
    return "help-fpd"
}

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fpd-name" {
        for _, c := range helpFpd.FpdName {
            if helpFpd.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName{}
        helpFpd.FpdName = append(helpFpd.FpdName, child)
        return &helpFpd.FpdName[len(helpFpd.FpdName)-1]
    }
    return nil
}

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helpFpd.FpdName {
        children[helpFpd.FpdName[i].GetSegmentPath()] = &helpFpd.FpdName[i]
    }
    return children
}

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetBundleName() string { return "cisco_ios_xr" }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetYangName() string { return "help-fpd" }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) SetParent(parent types.Entity) { helpFpd.parent = parent }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetParent() types.Entity { return helpFpd.parent }

func (helpFpd *ShowFpd_HelpLocations_HelpLocation_HelpFpd) GetParentYangName() string { return "help-location" }

// ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName
// Fpd name list
type ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // fpd location. The type is string.
    Location interface{}

    // fpd name. The type is string.
    FpdName interface{}
}

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetFilter() yfilter.YFilter { return fpdName.YFilter }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) SetFilter(yf yfilter.YFilter) { fpdName.YFilter = yf }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "fpd-name" { return "FpdName" }
    return ""
}

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetSegmentPath() string {
    return "fpd-name"
}

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = fpdName.Location
    leafs["fpd-name"] = fpdName.FpdName
    return leafs
}

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetBundleName() string { return "cisco_ios_xr" }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetYangName() string { return "fpd-name" }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) SetParent(parent types.Entity) { fpdName.parent = parent }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetParent() types.Entity { return fpdName.parent }

func (fpdName *ShowFpd_HelpLocations_HelpLocation_HelpFpd_FpdName) GetParentYangName() string { return "help-fpd" }

// ShowFpd_HwModuleFpdHelpFpd
// Display help-fpd -show hw-module fpd help-fpd
type ShowFpd_HwModuleFpdHelpFpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fpd name list. The type is slice of ShowFpd_HwModuleFpdHelpFpd_FpdName.
    FpdName []ShowFpd_HwModuleFpdHelpFpd_FpdName
}

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetFilter() yfilter.YFilter { return hwModuleFpdHelpFpd.YFilter }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) SetFilter(yf yfilter.YFilter) { hwModuleFpdHelpFpd.YFilter = yf }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetGoName(yname string) string {
    if yname == "fpd-name" { return "FpdName" }
    return ""
}

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetSegmentPath() string {
    return "hw-module-fpd-help-fpd"
}

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fpd-name" {
        for _, c := range hwModuleFpdHelpFpd.FpdName {
            if hwModuleFpdHelpFpd.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_HwModuleFpdHelpFpd_FpdName{}
        hwModuleFpdHelpFpd.FpdName = append(hwModuleFpdHelpFpd.FpdName, child)
        return &hwModuleFpdHelpFpd.FpdName[len(hwModuleFpdHelpFpd.FpdName)-1]
    }
    return nil
}

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwModuleFpdHelpFpd.FpdName {
        children[hwModuleFpdHelpFpd.FpdName[i].GetSegmentPath()] = &hwModuleFpdHelpFpd.FpdName[i]
    }
    return children
}

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetBundleName() string { return "cisco_ios_xr" }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetYangName() string { return "hw-module-fpd-help-fpd" }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) SetParent(parent types.Entity) { hwModuleFpdHelpFpd.parent = parent }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetParent() types.Entity { return hwModuleFpdHelpFpd.parent }

func (hwModuleFpdHelpFpd *ShowFpd_HwModuleFpdHelpFpd) GetParentYangName() string { return "show-fpd" }

// ShowFpd_HwModuleFpdHelpFpd_FpdName
// Fpd name list
type ShowFpd_HwModuleFpdHelpFpd_FpdName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // fpd location. The type is string.
    Location interface{}

    // fpd name. The type is string.
    FpdName interface{}
}

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetFilter() yfilter.YFilter { return fpdName.YFilter }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) SetFilter(yf yfilter.YFilter) { fpdName.YFilter = yf }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "fpd-name" { return "FpdName" }
    return ""
}

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetSegmentPath() string {
    return "fpd-name"
}

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = fpdName.Location
    leafs["fpd-name"] = fpdName.FpdName
    return leafs
}

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetBundleName() string { return "cisco_ios_xr" }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetYangName() string { return "fpd-name" }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) SetParent(parent types.Entity) { fpdName.parent = parent }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetParent() types.Entity { return fpdName.parent }

func (fpdName *ShowFpd_HwModuleFpdHelpFpd_FpdName) GetParentYangName() string { return "hw-module-fpd-help-fpd" }

// ShowFpd_Package
// gets fpd package info
type ShowFpd_Package struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // fpd pkg list . The type is slice of ShowFpd_Package_FpdPkgData.
    FpdPkgData []ShowFpd_Package_FpdPkgData
}

func (self *ShowFpd_Package) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ShowFpd_Package) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ShowFpd_Package) GetGoName(yname string) string {
    if yname == "fpd-pkg-data" { return "FpdPkgData" }
    return ""
}

func (self *ShowFpd_Package) GetSegmentPath() string {
    return "package"
}

func (self *ShowFpd_Package) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fpd-pkg-data" {
        for _, c := range self.FpdPkgData {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_Package_FpdPkgData{}
        self.FpdPkgData = append(self.FpdPkgData, child)
        return &self.FpdPkgData[len(self.FpdPkgData)-1]
    }
    return nil
}

func (self *ShowFpd_Package) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.FpdPkgData {
        children[self.FpdPkgData[i].GetSegmentPath()] = &self.FpdPkgData[i]
    }
    return children
}

func (self *ShowFpd_Package) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *ShowFpd_Package) GetBundleName() string { return "cisco_ios_xr" }

func (self *ShowFpd_Package) GetYangName() string { return "package" }

func (self *ShowFpd_Package) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ShowFpd_Package) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ShowFpd_Package) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ShowFpd_Package) SetParent(parent types.Entity) { self.parent = parent }

func (self *ShowFpd_Package) GetParent() types.Entity { return self.parent }

func (self *ShowFpd_Package) GetParentYangName() string { return "show-fpd" }

// ShowFpd_Package_FpdPkgData
//  fpd pkg list 
type ShowFpd_Package_FpdPkgData struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetFilter() yfilter.YFilter { return fpdPkgData.YFilter }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) SetFilter(yf yfilter.YFilter) { fpdPkgData.YFilter = yf }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetGoName(yname string) string {
    if yname == "card-type" { return "CardType" }
    if yname == "fpd-desc" { return "FpdDesc" }
    if yname == "upgrade-method" { return "UpgradeMethod" }
    if yname == "fpd-ver" { return "FpdVer" }
    if yname == "min-sw-ver" { return "MinSwVer" }
    if yname == "min-hw-ver" { return "MinHwVer" }
    return ""
}

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetSegmentPath() string {
    return "fpd-pkg-data"
}

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-type"] = fpdPkgData.CardType
    leafs["fpd-desc"] = fpdPkgData.FpdDesc
    leafs["upgrade-method"] = fpdPkgData.UpgradeMethod
    leafs["fpd-ver"] = fpdPkgData.FpdVer
    leafs["min-sw-ver"] = fpdPkgData.MinSwVer
    leafs["min-hw-ver"] = fpdPkgData.MinHwVer
    return leafs
}

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetBundleName() string { return "cisco_ios_xr" }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetYangName() string { return "fpd-pkg-data" }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) SetParent(parent types.Entity) { fpdPkgData.parent = parent }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetParent() types.Entity { return fpdPkgData.parent }

func (fpdPkgData *ShowFpd_Package_FpdPkgData) GetParentYangName() string { return "package" }

// ShowFpd_LocationHelp
// fpd upgradable locations
type ShowFpd_LocationHelp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // card location list. The type is slice of ShowFpd_LocationHelp_LocationName.
    LocationName []ShowFpd_LocationHelp_LocationName
}

func (locationHelp *ShowFpd_LocationHelp) GetFilter() yfilter.YFilter { return locationHelp.YFilter }

func (locationHelp *ShowFpd_LocationHelp) SetFilter(yf yfilter.YFilter) { locationHelp.YFilter = yf }

func (locationHelp *ShowFpd_LocationHelp) GetGoName(yname string) string {
    if yname == "location-name" { return "LocationName" }
    return ""
}

func (locationHelp *ShowFpd_LocationHelp) GetSegmentPath() string {
    return "location-help"
}

func (locationHelp *ShowFpd_LocationHelp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "location-name" {
        for _, c := range locationHelp.LocationName {
            if locationHelp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowFpd_LocationHelp_LocationName{}
        locationHelp.LocationName = append(locationHelp.LocationName, child)
        return &locationHelp.LocationName[len(locationHelp.LocationName)-1]
    }
    return nil
}

func (locationHelp *ShowFpd_LocationHelp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locationHelp.LocationName {
        children[locationHelp.LocationName[i].GetSegmentPath()] = &locationHelp.LocationName[i]
    }
    return children
}

func (locationHelp *ShowFpd_LocationHelp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locationHelp *ShowFpd_LocationHelp) GetBundleName() string { return "cisco_ios_xr" }

func (locationHelp *ShowFpd_LocationHelp) GetYangName() string { return "location-help" }

func (locationHelp *ShowFpd_LocationHelp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locationHelp *ShowFpd_LocationHelp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locationHelp *ShowFpd_LocationHelp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locationHelp *ShowFpd_LocationHelp) SetParent(parent types.Entity) { locationHelp.parent = parent }

func (locationHelp *ShowFpd_LocationHelp) GetParent() types.Entity { return locationHelp.parent }

func (locationHelp *ShowFpd_LocationHelp) GetParentYangName() string { return "show-fpd" }

// ShowFpd_LocationHelp_LocationName
// card location list
type ShowFpd_LocationHelp_LocationName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // card location. The type is string.
    LocationName interface{}
}

func (locationName *ShowFpd_LocationHelp_LocationName) GetFilter() yfilter.YFilter { return locationName.YFilter }

func (locationName *ShowFpd_LocationHelp_LocationName) SetFilter(yf yfilter.YFilter) { locationName.YFilter = yf }

func (locationName *ShowFpd_LocationHelp_LocationName) GetGoName(yname string) string {
    if yname == "location-name" { return "LocationName" }
    return ""
}

func (locationName *ShowFpd_LocationHelp_LocationName) GetSegmentPath() string {
    return "location-name"
}

func (locationName *ShowFpd_LocationHelp_LocationName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (locationName *ShowFpd_LocationHelp_LocationName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (locationName *ShowFpd_LocationHelp_LocationName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location-name"] = locationName.LocationName
    return leafs
}

func (locationName *ShowFpd_LocationHelp_LocationName) GetBundleName() string { return "cisco_ios_xr" }

func (locationName *ShowFpd_LocationHelp_LocationName) GetYangName() string { return "location-name" }

func (locationName *ShowFpd_LocationHelp_LocationName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locationName *ShowFpd_LocationHelp_LocationName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locationName *ShowFpd_LocationHelp_LocationName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locationName *ShowFpd_LocationHelp_LocationName) SetParent(parent types.Entity) { locationName.parent = parent }

func (locationName *ShowFpd_LocationHelp_LocationName) GetParent() types.Entity { return locationName.parent }

func (locationName *ShowFpd_LocationHelp_LocationName) GetParentYangName() string { return "location-help" }

