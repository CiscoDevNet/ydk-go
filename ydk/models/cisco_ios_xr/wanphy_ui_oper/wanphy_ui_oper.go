// This module contains a collection of YANG definitions
// for Cisco IOS-XR wanphy-ui package operational data.
// 
// This module contains definitions
// for the following management objects:
//   wanphy: WANPHY operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package wanphy_ui_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package wanphy_ui_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-wanphy-ui-oper wanphy}", reflect.TypeOf(Wanphy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-wanphy-ui-oper:wanphy", reflect.TypeOf(Wanphy{}))
}

// WanphyAlarmRepStatus represents WANPHY Alarm Report Status
type WanphyAlarmRepStatus string

const (
    // Alarm reporting is disable
    WanphyAlarmRepStatus_disable WanphyAlarmRepStatus = "disable"

    // Alarm reporting is enable
    WanphyAlarmRepStatus_enable WanphyAlarmRepStatus = "enable"
)

// WanphyModeInfo represents WANPHY Modes
type WanphyModeInfo string

const (
    // LAN mode
    WanphyModeInfo_lan WanphyModeInfo = "lan"

    // WAN mode
    WanphyModeInfo_wan WanphyModeInfo = "wan"
)

// Wanphy
// WANPHY operational data
type Wanphy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All WANPHY controller operational data.
    Controllers Wanphy_Controllers
}

func (wanphy *Wanphy) GetFilter() yfilter.YFilter { return wanphy.YFilter }

func (wanphy *Wanphy) SetFilter(yf yfilter.YFilter) { wanphy.YFilter = yf }

func (wanphy *Wanphy) GetGoName(yname string) string {
    if yname == "controllers" { return "Controllers" }
    return ""
}

func (wanphy *Wanphy) GetSegmentPath() string {
    return "Cisco-IOS-XR-wanphy-ui-oper:wanphy"
}

func (wanphy *Wanphy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controllers" {
        return &wanphy.Controllers
    }
    return nil
}

func (wanphy *Wanphy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["controllers"] = &wanphy.Controllers
    return children
}

func (wanphy *Wanphy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (wanphy *Wanphy) GetBundleName() string { return "cisco_ios_xr" }

func (wanphy *Wanphy) GetYangName() string { return "wanphy" }

func (wanphy *Wanphy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wanphy *Wanphy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wanphy *Wanphy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wanphy *Wanphy) SetParent(parent types.Entity) { wanphy.parent = parent }

func (wanphy *Wanphy) GetParent() types.Entity { return wanphy.parent }

func (wanphy *Wanphy) GetParentYangName() string { return "Cisco-IOS-XR-wanphy-ui-oper" }

// Wanphy_Controllers
// All WANPHY controller operational data
type Wanphy_Controllers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // WANPHY controller operational data. The type is slice of
    // Wanphy_Controllers_Controller.
    Controller []Wanphy_Controllers_Controller
}

func (controllers *Wanphy_Controllers) GetFilter() yfilter.YFilter { return controllers.YFilter }

func (controllers *Wanphy_Controllers) SetFilter(yf yfilter.YFilter) { controllers.YFilter = yf }

func (controllers *Wanphy_Controllers) GetGoName(yname string) string {
    if yname == "controller" { return "Controller" }
    return ""
}

func (controllers *Wanphy_Controllers) GetSegmentPath() string {
    return "controllers"
}

func (controllers *Wanphy_Controllers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controller" {
        for _, c := range controllers.Controller {
            if controllers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Wanphy_Controllers_Controller{}
        controllers.Controller = append(controllers.Controller, child)
        return &controllers.Controller[len(controllers.Controller)-1]
    }
    return nil
}

func (controllers *Wanphy_Controllers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controllers.Controller {
        children[controllers.Controller[i].GetSegmentPath()] = &controllers.Controller[i]
    }
    return children
}

func (controllers *Wanphy_Controllers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllers *Wanphy_Controllers) GetBundleName() string { return "cisco_ios_xr" }

func (controllers *Wanphy_Controllers) GetYangName() string { return "controllers" }

func (controllers *Wanphy_Controllers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllers *Wanphy_Controllers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllers *Wanphy_Controllers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllers *Wanphy_Controllers) SetParent(parent types.Entity) { controllers.parent = parent }

func (controllers *Wanphy_Controllers) GetParent() types.Entity { return controllers.parent }

func (controllers *Wanphy_Controllers) GetParentYangName() string { return "wanphy" }

// Wanphy_Controllers_Controller
// WANPHY controller operational data
type Wanphy_Controllers_Controller struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // WANPHY controller operational data.
    Info Wanphy_Controllers_Controller_Info
}

func (controller *Wanphy_Controllers_Controller) GetFilter() yfilter.YFilter { return controller.YFilter }

func (controller *Wanphy_Controllers_Controller) SetFilter(yf yfilter.YFilter) { controller.YFilter = yf }

func (controller *Wanphy_Controllers_Controller) GetGoName(yname string) string {
    if yname == "controller-name" { return "ControllerName" }
    if yname == "info" { return "Info" }
    return ""
}

func (controller *Wanphy_Controllers_Controller) GetSegmentPath() string {
    return "controller" + "[controller-name='" + fmt.Sprintf("%v", controller.ControllerName) + "']"
}

func (controller *Wanphy_Controllers_Controller) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "info" {
        return &controller.Info
    }
    return nil
}

func (controller *Wanphy_Controllers_Controller) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["info"] = &controller.Info
    return children
}

func (controller *Wanphy_Controllers_Controller) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller-name"] = controller.ControllerName
    return leafs
}

func (controller *Wanphy_Controllers_Controller) GetBundleName() string { return "cisco_ios_xr" }

func (controller *Wanphy_Controllers_Controller) GetYangName() string { return "controller" }

func (controller *Wanphy_Controllers_Controller) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controller *Wanphy_Controllers_Controller) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controller *Wanphy_Controllers_Controller) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controller *Wanphy_Controllers_Controller) SetParent(parent types.Entity) { controller.parent = parent }

func (controller *Wanphy_Controllers_Controller) GetParent() types.Entity { return controller.parent }

func (controller *Wanphy_Controllers_Controller) GetParentYangName() string { return "controllers" }

// Wanphy_Controllers_Controller_Info
// WANPHY controller operational data
type Wanphy_Controllers_Controller_Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration Mode. The type is WanphyModeInfo.
    AdminMode interface{}

    // Port State. The type is interface{} with range: 0..4294967295.
    PortState interface{}

    // Section LOF. The type is interface{} with range: 0..4294967295.
    SectionLof interface{}

    // Section LOS. The type is interface{} with range: 0..4294967295.
    SectionLos interface{}

    // Section BIP(B1). The type is interface{} with range:
    // 0..18446744073709551615.
    SectionBip interface{}

    // Line AIS. The type is interface{} with range: 0..4294967295.
    LineAis interface{}

    // Line RDI. The type is interface{} with range: 0..4294967295.
    LineRdi interface{}

    // Line FEBE. The type is interface{} with range: 0..18446744073709551615.
    LineFebe interface{}

    // Line BIP(B2) . The type is interface{} with range: 0..18446744073709551615.
    LineBip interface{}

    // Path AIS. The type is interface{} with range: 0..4294967295.
    PathAis interface{}

    // Path RDI. The type is interface{} with range: 0..4294967295.
    PathRdi interface{}

    // Path FEBE. The type is interface{} with range: 0..18446744073709551615.
    PathFebe interface{}

    // Path BIP(B3). The type is interface{} with range: 0..18446744073709551615.
    PathBip interface{}

    // Path LOP. The type is interface{} with range: 0..4294967295.
    PathLop interface{}

    // Path NEWPTR. The type is interface{} with range: 0..4294967295.
    PathNewptr interface{}

    // Path PSE. The type is interface{} with range: 0..4294967295.
    PathPse interface{}

    // Path NSE. The type is interface{} with range: 0..4294967295.
    PathNse interface{}

    // WIS Alarms SER. The type is interface{} with range: 0..4294967295.
    WisAlarmsSer interface{}

    // WIS Alarms FELCDP. The type is interface{} with range: 0..4294967295.
    WisAlarmsFelcdp interface{}

    // WIS Alarms FEAISP. The type is interface{} with range: 0..4294967295.
    WisAlarmsFeaisp interface{}

    // WIS Alarms WLOS. The type is interface{} with range: 0..4294967295.
    WisAlarmsWlos interface{}

    // WIS Alarms PLCD. The type is interface{} with range: 0..4294967295.
    WisAlarmsPlcd interface{}

    // WIS Alarms LFEBIP. The type is interface{} with range:
    // 0..18446744073709551615.
    WisAlarmsLfebip interface{}

    // WIS Alarms PBEC. The type is interface{} with range:
    // 0..18446744073709551615.
    WisAlarmsPbec interface{}

    // WIS Alarms PLMP. The type is interface{} with range: 0..4294967295.
    WisAlarmsPlmp interface{}

    // BER thresholds: SF. Value 'd' in 10e-%d. The type is interface{} with
    // range: 0..4294967295.
    SfBerThreshold interface{}

    // BER thresholds: SD. Value 'd' in 10e-%d. The type is interface{} with
    // range: 0..4294967295.
    SdBerThreshold interface{}

    // SF_BER Report. The type is WanphyAlarmRepStatus.
    SfBerReport interface{}

    // SD_BER Report. The type is WanphyAlarmRepStatus.
    SdBerReport interface{}

    // Operational Mode. The type is WanphyModeInfo.
    OperationalMode interface{}

    // Remote IP Address. The type is string.
    RemoteIp interface{}

    // Register P_FEBE. The type is interface{} with range: 0..4294967295.
    RegisterPFebe interface{}

    // Register L_FE_BIP. The type is interface{} with range: 0..4294967295.
    RegisterLFeBip interface{}

    // Register L_BIP. The type is interface{} with range: 0..4294967295.
    RegisterLBip interface{}

    // Register P_BEC. The type is interface{} with range: 0..4294967295.
    RegisterPBec interface{}

    // Register S_BIP. The type is interface{} with range: 0..4294967295.
    RegisterSBip interface{}

    // Register J1-Rx0. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx0 interface{}

    // Register J1-Rx1. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx1 interface{}

    // Register J1-Rx2. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx2 interface{}

    // Register J1-Rx3. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx3 interface{}

    // Register J1-Rx4. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx4 interface{}

    // Register J1-Rx5. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx5 interface{}

    // Register J1-Rx6. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx6 interface{}

    // Register J1-Rx7. The type is interface{} with range: 0..4294967295.
    RegisterJ1Rx7 interface{}

    // wanphy poll timer. The type is interface{} with range: 0..4294967295.
    WanphyPollTimer interface{}
}

func (info *Wanphy_Controllers_Controller_Info) GetFilter() yfilter.YFilter { return info.YFilter }

func (info *Wanphy_Controllers_Controller_Info) SetFilter(yf yfilter.YFilter) { info.YFilter = yf }

func (info *Wanphy_Controllers_Controller_Info) GetGoName(yname string) string {
    if yname == "admin-mode" { return "AdminMode" }
    if yname == "port-state" { return "PortState" }
    if yname == "section-lof" { return "SectionLof" }
    if yname == "section-los" { return "SectionLos" }
    if yname == "section-bip" { return "SectionBip" }
    if yname == "line-ais" { return "LineAis" }
    if yname == "line-rdi" { return "LineRdi" }
    if yname == "line-febe" { return "LineFebe" }
    if yname == "line-bip" { return "LineBip" }
    if yname == "path-ais" { return "PathAis" }
    if yname == "path-rdi" { return "PathRdi" }
    if yname == "path-febe" { return "PathFebe" }
    if yname == "path-bip" { return "PathBip" }
    if yname == "path-lop" { return "PathLop" }
    if yname == "path-newptr" { return "PathNewptr" }
    if yname == "path-pse" { return "PathPse" }
    if yname == "path-nse" { return "PathNse" }
    if yname == "wis-alarms-ser" { return "WisAlarmsSer" }
    if yname == "wis-alarms-felcdp" { return "WisAlarmsFelcdp" }
    if yname == "wis-alarms-feaisp" { return "WisAlarmsFeaisp" }
    if yname == "wis-alarms-wlos" { return "WisAlarmsWlos" }
    if yname == "wis-alarms-plcd" { return "WisAlarmsPlcd" }
    if yname == "wis-alarms-lfebip" { return "WisAlarmsLfebip" }
    if yname == "wis-alarms-pbec" { return "WisAlarmsPbec" }
    if yname == "wis-alarms-plmp" { return "WisAlarmsPlmp" }
    if yname == "sf-ber-threshold" { return "SfBerThreshold" }
    if yname == "sd-ber-threshold" { return "SdBerThreshold" }
    if yname == "sf-ber-report" { return "SfBerReport" }
    if yname == "sd-ber-report" { return "SdBerReport" }
    if yname == "operational-mode" { return "OperationalMode" }
    if yname == "remote-ip" { return "RemoteIp" }
    if yname == "register-p-febe" { return "RegisterPFebe" }
    if yname == "register-l-fe-bip" { return "RegisterLFeBip" }
    if yname == "register-l-bip" { return "RegisterLBip" }
    if yname == "register-p-bec" { return "RegisterPBec" }
    if yname == "register-s-bip" { return "RegisterSBip" }
    if yname == "register-j1-rx0" { return "RegisterJ1Rx0" }
    if yname == "register-j1-rx1" { return "RegisterJ1Rx1" }
    if yname == "register-j1-rx2" { return "RegisterJ1Rx2" }
    if yname == "register-j1-rx3" { return "RegisterJ1Rx3" }
    if yname == "register-j1-rx4" { return "RegisterJ1Rx4" }
    if yname == "register-j1-rx5" { return "RegisterJ1Rx5" }
    if yname == "register-j1-rx6" { return "RegisterJ1Rx6" }
    if yname == "register-j1-rx7" { return "RegisterJ1Rx7" }
    if yname == "wanphy-poll-timer" { return "WanphyPollTimer" }
    return ""
}

func (info *Wanphy_Controllers_Controller_Info) GetSegmentPath() string {
    return "info"
}

func (info *Wanphy_Controllers_Controller_Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (info *Wanphy_Controllers_Controller_Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (info *Wanphy_Controllers_Controller_Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-mode"] = info.AdminMode
    leafs["port-state"] = info.PortState
    leafs["section-lof"] = info.SectionLof
    leafs["section-los"] = info.SectionLos
    leafs["section-bip"] = info.SectionBip
    leafs["line-ais"] = info.LineAis
    leafs["line-rdi"] = info.LineRdi
    leafs["line-febe"] = info.LineFebe
    leafs["line-bip"] = info.LineBip
    leafs["path-ais"] = info.PathAis
    leafs["path-rdi"] = info.PathRdi
    leafs["path-febe"] = info.PathFebe
    leafs["path-bip"] = info.PathBip
    leafs["path-lop"] = info.PathLop
    leafs["path-newptr"] = info.PathNewptr
    leafs["path-pse"] = info.PathPse
    leafs["path-nse"] = info.PathNse
    leafs["wis-alarms-ser"] = info.WisAlarmsSer
    leafs["wis-alarms-felcdp"] = info.WisAlarmsFelcdp
    leafs["wis-alarms-feaisp"] = info.WisAlarmsFeaisp
    leafs["wis-alarms-wlos"] = info.WisAlarmsWlos
    leafs["wis-alarms-plcd"] = info.WisAlarmsPlcd
    leafs["wis-alarms-lfebip"] = info.WisAlarmsLfebip
    leafs["wis-alarms-pbec"] = info.WisAlarmsPbec
    leafs["wis-alarms-plmp"] = info.WisAlarmsPlmp
    leafs["sf-ber-threshold"] = info.SfBerThreshold
    leafs["sd-ber-threshold"] = info.SdBerThreshold
    leafs["sf-ber-report"] = info.SfBerReport
    leafs["sd-ber-report"] = info.SdBerReport
    leafs["operational-mode"] = info.OperationalMode
    leafs["remote-ip"] = info.RemoteIp
    leafs["register-p-febe"] = info.RegisterPFebe
    leafs["register-l-fe-bip"] = info.RegisterLFeBip
    leafs["register-l-bip"] = info.RegisterLBip
    leafs["register-p-bec"] = info.RegisterPBec
    leafs["register-s-bip"] = info.RegisterSBip
    leafs["register-j1-rx0"] = info.RegisterJ1Rx0
    leafs["register-j1-rx1"] = info.RegisterJ1Rx1
    leafs["register-j1-rx2"] = info.RegisterJ1Rx2
    leafs["register-j1-rx3"] = info.RegisterJ1Rx3
    leafs["register-j1-rx4"] = info.RegisterJ1Rx4
    leafs["register-j1-rx5"] = info.RegisterJ1Rx5
    leafs["register-j1-rx6"] = info.RegisterJ1Rx6
    leafs["register-j1-rx7"] = info.RegisterJ1Rx7
    leafs["wanphy-poll-timer"] = info.WanphyPollTimer
    return leafs
}

func (info *Wanphy_Controllers_Controller_Info) GetBundleName() string { return "cisco_ios_xr" }

func (info *Wanphy_Controllers_Controller_Info) GetYangName() string { return "info" }

func (info *Wanphy_Controllers_Controller_Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (info *Wanphy_Controllers_Controller_Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (info *Wanphy_Controllers_Controller_Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (info *Wanphy_Controllers_Controller_Info) SetParent(parent types.Entity) { info.parent = parent }

func (info *Wanphy_Controllers_Controller_Info) GetParent() types.Entity { return info.parent }

func (info *Wanphy_Controllers_Controller_Info) GetParentYangName() string { return "controller" }

