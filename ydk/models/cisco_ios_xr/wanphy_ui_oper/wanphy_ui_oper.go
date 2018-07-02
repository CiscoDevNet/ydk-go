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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All WANPHY controller operational data.
    Controllers Wanphy_Controllers
}

func (wanphy *Wanphy) GetEntityData() *types.CommonEntityData {
    wanphy.EntityData.YFilter = wanphy.YFilter
    wanphy.EntityData.YangName = "wanphy"
    wanphy.EntityData.BundleName = "cisco_ios_xr"
    wanphy.EntityData.ParentYangName = "Cisco-IOS-XR-wanphy-ui-oper"
    wanphy.EntityData.SegmentPath = "Cisco-IOS-XR-wanphy-ui-oper:wanphy"
    wanphy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wanphy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wanphy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wanphy.EntityData.Children = types.NewOrderedMap()
    wanphy.EntityData.Children.Append("controllers", types.YChild{"Controllers", &wanphy.Controllers})
    wanphy.EntityData.Leafs = types.NewOrderedMap()

    wanphy.EntityData.YListKeys = []string {}

    return &(wanphy.EntityData)
}

// Wanphy_Controllers
// All WANPHY controller operational data
type Wanphy_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WANPHY controller operational data. The type is slice of
    // Wanphy_Controllers_Controller.
    Controller []*Wanphy_Controllers_Controller
}

func (controllers *Wanphy_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "wanphy"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controller", types.YChild{"Controller", nil})
    for i := range controllers.Controller {
        controllers.EntityData.Children.Append(types.GetSegmentPath(controllers.Controller[i]), types.YChild{"Controller", controllers.Controller[i]})
    }
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// Wanphy_Controllers_Controller
// WANPHY controller operational data
type Wanphy_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // WANPHY controller operational data.
    Info Wanphy_Controllers_Controller_Info
}

func (controller *Wanphy_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("info", types.YChild{"Info", &controller.Info})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// Wanphy_Controllers_Controller_Info
// WANPHY controller operational data
type Wanphy_Controllers_Controller_Info struct {
    EntityData types.CommonEntityData
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

func (info *Wanphy_Controllers_Controller_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "controller"
    info.EntityData.SegmentPath = "info"
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("admin-mode", types.YLeaf{"AdminMode", info.AdminMode})
    info.EntityData.Leafs.Append("port-state", types.YLeaf{"PortState", info.PortState})
    info.EntityData.Leafs.Append("section-lof", types.YLeaf{"SectionLof", info.SectionLof})
    info.EntityData.Leafs.Append("section-los", types.YLeaf{"SectionLos", info.SectionLos})
    info.EntityData.Leafs.Append("section-bip", types.YLeaf{"SectionBip", info.SectionBip})
    info.EntityData.Leafs.Append("line-ais", types.YLeaf{"LineAis", info.LineAis})
    info.EntityData.Leafs.Append("line-rdi", types.YLeaf{"LineRdi", info.LineRdi})
    info.EntityData.Leafs.Append("line-febe", types.YLeaf{"LineFebe", info.LineFebe})
    info.EntityData.Leafs.Append("line-bip", types.YLeaf{"LineBip", info.LineBip})
    info.EntityData.Leafs.Append("path-ais", types.YLeaf{"PathAis", info.PathAis})
    info.EntityData.Leafs.Append("path-rdi", types.YLeaf{"PathRdi", info.PathRdi})
    info.EntityData.Leafs.Append("path-febe", types.YLeaf{"PathFebe", info.PathFebe})
    info.EntityData.Leafs.Append("path-bip", types.YLeaf{"PathBip", info.PathBip})
    info.EntityData.Leafs.Append("path-lop", types.YLeaf{"PathLop", info.PathLop})
    info.EntityData.Leafs.Append("path-newptr", types.YLeaf{"PathNewptr", info.PathNewptr})
    info.EntityData.Leafs.Append("path-pse", types.YLeaf{"PathPse", info.PathPse})
    info.EntityData.Leafs.Append("path-nse", types.YLeaf{"PathNse", info.PathNse})
    info.EntityData.Leafs.Append("wis-alarms-ser", types.YLeaf{"WisAlarmsSer", info.WisAlarmsSer})
    info.EntityData.Leafs.Append("wis-alarms-felcdp", types.YLeaf{"WisAlarmsFelcdp", info.WisAlarmsFelcdp})
    info.EntityData.Leafs.Append("wis-alarms-feaisp", types.YLeaf{"WisAlarmsFeaisp", info.WisAlarmsFeaisp})
    info.EntityData.Leafs.Append("wis-alarms-wlos", types.YLeaf{"WisAlarmsWlos", info.WisAlarmsWlos})
    info.EntityData.Leafs.Append("wis-alarms-plcd", types.YLeaf{"WisAlarmsPlcd", info.WisAlarmsPlcd})
    info.EntityData.Leafs.Append("wis-alarms-lfebip", types.YLeaf{"WisAlarmsLfebip", info.WisAlarmsLfebip})
    info.EntityData.Leafs.Append("wis-alarms-pbec", types.YLeaf{"WisAlarmsPbec", info.WisAlarmsPbec})
    info.EntityData.Leafs.Append("wis-alarms-plmp", types.YLeaf{"WisAlarmsPlmp", info.WisAlarmsPlmp})
    info.EntityData.Leafs.Append("sf-ber-threshold", types.YLeaf{"SfBerThreshold", info.SfBerThreshold})
    info.EntityData.Leafs.Append("sd-ber-threshold", types.YLeaf{"SdBerThreshold", info.SdBerThreshold})
    info.EntityData.Leafs.Append("sf-ber-report", types.YLeaf{"SfBerReport", info.SfBerReport})
    info.EntityData.Leafs.Append("sd-ber-report", types.YLeaf{"SdBerReport", info.SdBerReport})
    info.EntityData.Leafs.Append("operational-mode", types.YLeaf{"OperationalMode", info.OperationalMode})
    info.EntityData.Leafs.Append("remote-ip", types.YLeaf{"RemoteIp", info.RemoteIp})
    info.EntityData.Leafs.Append("register-p-febe", types.YLeaf{"RegisterPFebe", info.RegisterPFebe})
    info.EntityData.Leafs.Append("register-l-fe-bip", types.YLeaf{"RegisterLFeBip", info.RegisterLFeBip})
    info.EntityData.Leafs.Append("register-l-bip", types.YLeaf{"RegisterLBip", info.RegisterLBip})
    info.EntityData.Leafs.Append("register-p-bec", types.YLeaf{"RegisterPBec", info.RegisterPBec})
    info.EntityData.Leafs.Append("register-s-bip", types.YLeaf{"RegisterSBip", info.RegisterSBip})
    info.EntityData.Leafs.Append("register-j1-rx0", types.YLeaf{"RegisterJ1Rx0", info.RegisterJ1Rx0})
    info.EntityData.Leafs.Append("register-j1-rx1", types.YLeaf{"RegisterJ1Rx1", info.RegisterJ1Rx1})
    info.EntityData.Leafs.Append("register-j1-rx2", types.YLeaf{"RegisterJ1Rx2", info.RegisterJ1Rx2})
    info.EntityData.Leafs.Append("register-j1-rx3", types.YLeaf{"RegisterJ1Rx3", info.RegisterJ1Rx3})
    info.EntityData.Leafs.Append("register-j1-rx4", types.YLeaf{"RegisterJ1Rx4", info.RegisterJ1Rx4})
    info.EntityData.Leafs.Append("register-j1-rx5", types.YLeaf{"RegisterJ1Rx5", info.RegisterJ1Rx5})
    info.EntityData.Leafs.Append("register-j1-rx6", types.YLeaf{"RegisterJ1Rx6", info.RegisterJ1Rx6})
    info.EntityData.Leafs.Append("register-j1-rx7", types.YLeaf{"RegisterJ1Rx7", info.RegisterJ1Rx7})
    info.EntityData.Leafs.Append("wanphy-poll-timer", types.YLeaf{"WanphyPollTimer", info.WanphyPollTimer})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

