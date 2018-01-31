// This module contains a collection of YANG definitions
// for Cisco IOS-XR controller-otu package operational data.
// 
// This module contains definitions
// for the following management objects:
//   otu: OTU operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package controller_otu_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_otu_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-controller-otu-oper otu}", reflect.TypeOf(Otu{}))
    ydk.RegisterEntity("Cisco-IOS-XR-controller-otu-oper:otu", reflect.TypeOf(Otu{}))
}

// OtuPrbsStatus represents Otu prbs status
type OtuPrbsStatus string

const (
    // Locked
    OtuPrbsStatus_locked OtuPrbsStatus = "locked"

    // Unlocked
    OtuPrbsStatus_unlocked OtuPrbsStatus = "unlocked"

    // Not Applicable
    OtuPrbsStatus_not_applicable OtuPrbsStatus = "not-applicable"
)

// OtuPrbsPattern represents Otu prbs pattern
type OtuPrbsPattern string

const (
    // OTU PRBS pattern not applicable
    OtuPrbsPattern_not_applicable OtuPrbsPattern = "not-applicable"

    // PN31
    OtuPrbsPattern_pn31 OtuPrbsPattern = "pn31"

    // PN23
    OtuPrbsPattern_pn23 OtuPrbsPattern = "pn23"

    // PN11
    OtuPrbsPattern_pn11 OtuPrbsPattern = "pn11"

    // INVERTED PN31
    OtuPrbsPattern_inverted_pn31 OtuPrbsPattern = "inverted-pn31"

    // INVERTED PN11
    OtuPrbsPattern_inverted_pn11 OtuPrbsPattern = "inverted-pn11"

    // PN15
    OtuPrbsPattern_pn15 OtuPrbsPattern = "pn15"
)

// OtuStateEt represents Otu state et
type OtuStateEt string

const (
    // Not Ready
    OtuStateEt_not_ready OtuStateEt = "not-ready"

    // Admin Down
    OtuStateEt_admin_down OtuStateEt = "admin-down"

    // Down
    OtuStateEt_down OtuStateEt = "down"

    // Up
    OtuStateEt_up OtuStateEt = "up"

    // Shutdown
    OtuStateEt_shutdown OtuStateEt = "shutdown"

    // Error Disable
    OtuStateEt_error_disable OtuStateEt = "error-disable"

    // Down Immediate
    OtuStateEt_down_immediate OtuStateEt = "down-immediate"

    // Down Immediate Admin
    OtuStateEt_down_immediate_admin OtuStateEt = "down-immediate-admin"

    // Down Graceful
    OtuStateEt_down_graceful OtuStateEt = "down-graceful"

    // Begin Shutdown
    OtuStateEt_begin_shutdown OtuStateEt = "begin-shutdown"

    // End Shutdown
    OtuStateEt_end_shutdown OtuStateEt = "end-shutdown"

    // Begin Error Disable
    OtuStateEt_begin_error_disable OtuStateEt = "begin-error-disable"

    // End Error Disable
    OtuStateEt_end_error_disable OtuStateEt = "end-error-disable"

    // Begin Down Graceful
    OtuStateEt_begin_down_graceful OtuStateEt = "begin-down-graceful"

    // Reset
    OtuStateEt_reset OtuStateEt = "reset"

    // Operational
    OtuStateEt_operational OtuStateEt = "operational"

    // Not Operational
    OtuStateEt_not_operational OtuStateEt = "not-operational"

    // Unknown
    OtuStateEt_unknown OtuStateEt = "unknown"

    // Last
    OtuStateEt_last OtuStateEt = "last"
)

// OtuPrbsTest represents Otu prbs test
type OtuPrbsTest string

const (
    // Disable
    OtuPrbsTest_disable OtuPrbsTest = "disable"

    // Enable
    OtuPrbsTest_enable OtuPrbsTest = "enable"
)

// OtuPpFsmState represents Otu pp fsm state
type OtuPpFsmState string

const (
    // In Active
    OtuPpFsmState_otu_in_active OtuPpFsmState = "otu-in-active"

    // Disabled
    OtuPpFsmState_otu_disabled OtuPpFsmState = "otu-disabled"

    // Normal
    OtuPpFsmState_otu_normal_state OtuPpFsmState = "otu-normal-state"

    // Local Failing
    OtuPpFsmState_otu_local_failing OtuPpFsmState = "otu-local-failing"

    // Remote Failing
    OtuPpFsmState_otu_remote_failing OtuPpFsmState = "otu-remote-failing"

    // Maintance Failing
    OtuPpFsmState_otu_main_t_failing OtuPpFsmState = "otu-main-t-failing"

    // Regenerator Failing
    OtuPpFsmState_otu_regen_failing OtuPpFsmState = "otu-regen-failing"

    // Local Failed
    OtuPpFsmState_otu_local_failed OtuPpFsmState = "otu-local-failed"

    // Remote Failed
    OtuPpFsmState_otu_remote_failed OtuPpFsmState = "otu-remote-failed"

    // Maintance Failed
    OtuPpFsmState_otu_main_t_failed OtuPpFsmState = "otu-main-t-failed"

    // Regenerator Failed
    OtuPpFsmState_otu_regen_failed OtuPpFsmState = "otu-regen-failed"
)

// OtuG709FecMode represents Otu g709fec mode
type OtuG709FecMode string

const (
    // NONE
    OtuG709FecMode_otu_bag_none_fec OtuG709FecMode = "otu-bag-none-fec"

    // STANDARD
    OtuG709FecMode_otu_bag_standard_fec OtuG709FecMode = "otu-bag-standard-fec"

    // ENHANCEDI.7
    OtuG709FecMode_otu_bag_1_i_7_fec OtuG709FecMode = "otu-bag-1-i-7-fec"

    // ENHANCEDI.4
    OtuG709FecMode_otu_bag_1_i_4_fec OtuG709FecMode = "otu-bag-1-i-4-fec"

    // SWIZZLE
    OtuG709FecMode_otu_bag_swizzle_fec OtuG709FecMode = "otu-bag-swizzle-fec"

    // HIGH GAIN20
    OtuG709FecMode_otu_bag_hg20_fec OtuG709FecMode = "otu-bag-hg20-fec"

    // Enhanced High Gain 7
    OtuG709FecMode_otu_bag_enhanced_hg7_fec OtuG709FecMode = "otu-bag-enhanced-hg7-fec"

    // Soft-Decision 20
    OtuG709FecMode_otu_bag_sd20_fec OtuG709FecMode = "otu-bag-sd20-fec"

    // Soft-Decision 7
    OtuG709FecMode_otu_bag_sd7_fec OtuG709FecMode = "otu-bag-sd7-fec"

    // ALL
    OtuG709FecMode_otu_bag_all_fec OtuG709FecMode = "otu-bag-all-fec"
)

// OtuPrbsMode represents Otu prbs mode
type OtuPrbsMode string

const (
    // OTU PRBS mode not applicable
    OtuPrbsMode_not_applicable OtuPrbsMode = "not-applicable"

    // Source
    OtuPrbsMode_source OtuPrbsMode = "source"

    // Sink
    OtuPrbsMode_sink OtuPrbsMode = "sink"

    // Source Sink
    OtuPrbsMode_source_sink OtuPrbsMode = "source-sink"
)

// OtuPerMon represents Otu per mon
type OtuPerMon string

const (
    // Disable
    OtuPerMon_disable OtuPerMon = "disable"

    // Enable
    OtuPerMon_enable OtuPerMon = "enable"
)

// OtuTtiEt represents Otu tti et
type OtuTtiEt string

const (
    // ASCII
    OtuTtiEt_ascii OtuTtiEt = "ascii"

    // HEX
    OtuTtiEt_hex OtuTtiEt = "hex"

    // FULL ASCII
    OtuTtiEt_full_ascii OtuTtiEt = "full-ascii"

    // FULL HEX
    OtuTtiEt_full_hex OtuTtiEt = "full-hex"
)

// OtuPpIntfState represents Otu pp intf state
type OtuPpIntfState string

const (
    // Interface is Up
    OtuPpIntfState_otu_pp_intf_up OtuPpIntfState = "otu-pp-intf-up"

    // Interface is Going Down
    OtuPpIntfState_otu_pp_intf_failing OtuPpIntfState = "otu-pp-intf-failing"

    // Interface Down
    OtuPpIntfState_otu_pp_intf_down OtuPpIntfState = "otu-pp-intf-down"
)

// OtuSecState represents Otu sec state
type OtuSecState string

const (
    // Normal
    OtuSecState_normal OtuSecState = "normal"

    // Maintenance
    OtuSecState_maintenance OtuSecState = "maintenance"

    // Automatic In Service
    OtuSecState_ais OtuSecState = "ais"
)

// OtuLoopBackMode represents Otu loop back mode
type OtuLoopBackMode string

const (
    // None
    OtuLoopBackMode_none OtuLoopBackMode = "none"

    // Line
    OtuLoopBackMode_line OtuLoopBackMode = "line"

    // Internal
    OtuLoopBackMode_internal OtuLoopBackMode = "internal"
)

// GmplsOtuTtiMode represents Gmpls otu tti mode
type GmplsOtuTtiMode string

const (
    // Not Set
    GmplsOtuTtiMode_gmpls_otu_tti_mode_none GmplsOtuTtiMode = "gmpls-otu-tti-mode-none"

    // Section Monitoring
    GmplsOtuTtiMode_gmpls_otu_tti_mode_sm GmplsOtuTtiMode = "gmpls-otu-tti-mode-sm"

    // Path Monitoring
    GmplsOtuTtiMode_gmpls_otu_tti_mode_pm GmplsOtuTtiMode = "gmpls-otu-tti-mode-pm"

    // Tandem Connection Monitoring
    GmplsOtuTtiMode_gmpls_otu_tti_mode_tcm GmplsOtuTtiMode = "gmpls-otu-tti-mode-tcm"
)

// OtuDerState represents Otu der state
type OtuDerState string

const (
    // Out Of Service
    OtuDerState_out_of_service OtuDerState = "out-of-service"

    // In Service
    OtuDerState_in_service OtuDerState = "in-service"

    // Maintenance
    OtuDerState_maintenance OtuDerState = "maintenance"

    // Automatic In Service
    OtuDerState_ais OtuDerState = "ais"
)

// Otu
// OTU operational data
type Otu struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All OTU Port operational data.
    Controllers Otu_Controllers
}

func (otu *Otu) GetFilter() yfilter.YFilter { return otu.YFilter }

func (otu *Otu) SetFilter(yf yfilter.YFilter) { otu.YFilter = yf }

func (otu *Otu) GetGoName(yname string) string {
    if yname == "controllers" { return "Controllers" }
    return ""
}

func (otu *Otu) GetSegmentPath() string {
    return "Cisco-IOS-XR-controller-otu-oper:otu"
}

func (otu *Otu) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controllers" {
        return &otu.Controllers
    }
    return nil
}

func (otu *Otu) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["controllers"] = &otu.Controllers
    return children
}

func (otu *Otu) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (otu *Otu) GetBundleName() string { return "cisco_ios_xr" }

func (otu *Otu) GetYangName() string { return "otu" }

func (otu *Otu) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otu *Otu) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otu *Otu) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otu *Otu) SetParent(parent types.Entity) { otu.parent = parent }

func (otu *Otu) GetParent() types.Entity { return otu.parent }

func (otu *Otu) GetParentYangName() string { return "Cisco-IOS-XR-controller-otu-oper" }

// Otu_Controllers
// All OTU Port operational data
type Otu_Controllers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OTU Port operational data. The type is slice of Otu_Controllers_Controller.
    Controller []Otu_Controllers_Controller
}

func (controllers *Otu_Controllers) GetFilter() yfilter.YFilter { return controllers.YFilter }

func (controllers *Otu_Controllers) SetFilter(yf yfilter.YFilter) { controllers.YFilter = yf }

func (controllers *Otu_Controllers) GetGoName(yname string) string {
    if yname == "controller" { return "Controller" }
    return ""
}

func (controllers *Otu_Controllers) GetSegmentPath() string {
    return "controllers"
}

func (controllers *Otu_Controllers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "controller" {
        for _, c := range controllers.Controller {
            if controllers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Otu_Controllers_Controller{}
        controllers.Controller = append(controllers.Controller, child)
        return &controllers.Controller[len(controllers.Controller)-1]
    }
    return nil
}

func (controllers *Otu_Controllers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controllers.Controller {
        children[controllers.Controller[i].GetSegmentPath()] = &controllers.Controller[i]
    }
    return children
}

func (controllers *Otu_Controllers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllers *Otu_Controllers) GetBundleName() string { return "cisco_ios_xr" }

func (controllers *Otu_Controllers) GetYangName() string { return "controllers" }

func (controllers *Otu_Controllers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllers *Otu_Controllers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllers *Otu_Controllers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllers *Otu_Controllers) SetParent(parent types.Entity) { controllers.parent = parent }

func (controllers *Otu_Controllers) GetParent() types.Entity { return controllers.parent }

func (controllers *Otu_Controllers) GetParentYangName() string { return "otu" }

// Otu_Controllers_Controller
// OTU Port operational data
type Otu_Controllers_Controller struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // OTU port PRBS operational data.
    Prbs Otu_Controllers_Controller_Prbs

    // OTU port operational data.
    Info Otu_Controllers_Controller_Info
}

func (controller *Otu_Controllers_Controller) GetFilter() yfilter.YFilter { return controller.YFilter }

func (controller *Otu_Controllers_Controller) SetFilter(yf yfilter.YFilter) { controller.YFilter = yf }

func (controller *Otu_Controllers_Controller) GetGoName(yname string) string {
    if yname == "controller-name" { return "ControllerName" }
    if yname == "prbs" { return "Prbs" }
    if yname == "info" { return "Info" }
    return ""
}

func (controller *Otu_Controllers_Controller) GetSegmentPath() string {
    return "controller" + "[controller-name='" + fmt.Sprintf("%v", controller.ControllerName) + "']"
}

func (controller *Otu_Controllers_Controller) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prbs" {
        return &controller.Prbs
    }
    if childYangName == "info" {
        return &controller.Info
    }
    return nil
}

func (controller *Otu_Controllers_Controller) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prbs"] = &controller.Prbs
    children["info"] = &controller.Info
    return children
}

func (controller *Otu_Controllers_Controller) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller-name"] = controller.ControllerName
    return leafs
}

func (controller *Otu_Controllers_Controller) GetBundleName() string { return "cisco_ios_xr" }

func (controller *Otu_Controllers_Controller) GetYangName() string { return "controller" }

func (controller *Otu_Controllers_Controller) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controller *Otu_Controllers_Controller) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controller *Otu_Controllers_Controller) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controller *Otu_Controllers_Controller) SetParent(parent types.Entity) { controller.parent = parent }

func (controller *Otu_Controllers_Controller) GetParent() types.Entity { return controller.parent }

func (controller *Otu_Controllers_Controller) GetParentYangName() string { return "controllers" }

// Otu_Controllers_Controller_Prbs
// OTU port PRBS operational data
type Otu_Controllers_Controller_Prbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // otu prbs test. The type is OtuPrbsTest.
    OtuPrbsTest interface{}

    // otu prbs mode. The type is OtuPrbsMode.
    OtuPrbsMode interface{}

    // otu prbs pattern. The type is OtuPrbsPattern.
    OtuPrbsPattern interface{}

    // otu prbs status. The type is OtuPrbsStatus.
    OtuPrbsStatus interface{}
}

func (prbs *Otu_Controllers_Controller_Prbs) GetFilter() yfilter.YFilter { return prbs.YFilter }

func (prbs *Otu_Controllers_Controller_Prbs) SetFilter(yf yfilter.YFilter) { prbs.YFilter = yf }

func (prbs *Otu_Controllers_Controller_Prbs) GetGoName(yname string) string {
    if yname == "otu-prbs-test" { return "OtuPrbsTest" }
    if yname == "otu-prbs-mode" { return "OtuPrbsMode" }
    if yname == "otu-prbs-pattern" { return "OtuPrbsPattern" }
    if yname == "otu-prbs-status" { return "OtuPrbsStatus" }
    return ""
}

func (prbs *Otu_Controllers_Controller_Prbs) GetSegmentPath() string {
    return "prbs"
}

func (prbs *Otu_Controllers_Controller_Prbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prbs *Otu_Controllers_Controller_Prbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prbs *Otu_Controllers_Controller_Prbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["otu-prbs-test"] = prbs.OtuPrbsTest
    leafs["otu-prbs-mode"] = prbs.OtuPrbsMode
    leafs["otu-prbs-pattern"] = prbs.OtuPrbsPattern
    leafs["otu-prbs-status"] = prbs.OtuPrbsStatus
    return leafs
}

func (prbs *Otu_Controllers_Controller_Prbs) GetBundleName() string { return "cisco_ios_xr" }

func (prbs *Otu_Controllers_Controller_Prbs) GetYangName() string { return "prbs" }

func (prbs *Otu_Controllers_Controller_Prbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prbs *Otu_Controllers_Controller_Prbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prbs *Otu_Controllers_Controller_Prbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prbs *Otu_Controllers_Controller_Prbs) SetParent(parent types.Entity) { prbs.parent = parent }

func (prbs *Otu_Controllers_Controller_Prbs) GetParent() types.Entity { return prbs.parent }

func (prbs *Otu_Controllers_Controller_Prbs) GetParentYangName() string { return "controller" }

// Otu_Controllers_Controller_Info
// OTU port operational data
type Otu_Controllers_Controller_Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Admin State. The type is OtuStateEt.
    State interface{}

    // Interface Name. The type is string.
    Name interface{}

    // SF in the form of 1.0E - <SF>. The type is interface{} with range: 0..255.
    Sf interface{}

    // SD in the form of 1.0E - <SD>. The type is interface{} with range: 0..255.
    Sd interface{}

    // Loopback. The type is OtuLoopBackMode.
    LoopbackMode interface{}

    // FEC. The type is OtuG709FecMode.
    FecMode interface{}

    // Derived State. The type is OtuDerState.
    DerivedstateMode interface{}

    // Sec State. The type is OtuSecState.
    InheritSecState interface{}

    // Sec State. The type is OtuSecState.
    ConfigSecState interface{}

    // OTU GCC. The type is bool.
    GccMode interface{}

    // q value calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    Q interface{}

    // q margin calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    QMargin interface{}

    // Performance Monitoring. The type is OtuPerMon.
    PerformanceMonitoring interface{}

    // Average bit errors corrected. The type is interface{} with range:
    // 0..18446744073709551615.
    Ec interface{}

    // Uncorrected word count. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc interface{}

    // Pre FEC BER Value in form 0.00. The type is interface{} with range:
    // -2147483648..2147483647.
    PreFecVal interface{}

    // Pre FEC BER Mantissa in form E-<mantisaa>. The type is interface{} with
    // range: -128..127.
    PreFecMantissa interface{}

    // EC value present. The type is bool.
    EcValue interface{}

    // Uc value present. The type is bool.
    UcValue interface{}

    // Pre fec val present . The type is bool.
    PreFecBerValue interface{}

    // Pre fec val mantissa. The type is bool.
    PreFecBerMantissa interface{}

    // NV Optical support. The type is bool.
    NvOpticalSupport interface{}

    // GMPLS TTI MODE. The type is GmplsOtuTtiMode.
    GmplsTtiMode interface{}

    // GMPLS TCM ID. The type is interface{} with range: 0..255.
    GmplsTvmId interface{}

    // Auto tti flag. The type is bool.
    AutoTtiFlag interface{}

    // Controller description string. The type is string.
    Description interface{}

    // TTI.
    Local Otu_Controllers_Controller_Info_Local

    // Remote.
    Remote Otu_Controllers_Controller_Info_Remote

    // OTU TTI.
    TtiMode Otu_Controllers_Controller_Info_TtiMode

    // Network Shared Risk Link Group information.
    NetworkSrlg Otu_Controllers_Controller_Info_NetworkSrlg

    // OTU layer alarm Information.
    OtuAlarmInfo Otu_Controllers_Controller_Info_OtuAlarmInfo

    // Proactive Protection.
    Proactive Otu_Controllers_Controller_Info_Proactive

    // OTU FEC Statistics.
    OtuFecSatistics Otu_Controllers_Controller_Info_OtuFecSatistics
}

func (info *Otu_Controllers_Controller_Info) GetFilter() yfilter.YFilter { return info.YFilter }

func (info *Otu_Controllers_Controller_Info) SetFilter(yf yfilter.YFilter) { info.YFilter = yf }

func (info *Otu_Controllers_Controller_Info) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "name" { return "Name" }
    if yname == "sf" { return "Sf" }
    if yname == "sd" { return "Sd" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    if yname == "fec-mode" { return "FecMode" }
    if yname == "derivedstate-mode" { return "DerivedstateMode" }
    if yname == "inherit-sec-state" { return "InheritSecState" }
    if yname == "config-sec-state" { return "ConfigSecState" }
    if yname == "gcc-mode" { return "GccMode" }
    if yname == "q" { return "Q" }
    if yname == "q-margin" { return "QMargin" }
    if yname == "performance-monitoring" { return "PerformanceMonitoring" }
    if yname == "ec" { return "Ec" }
    if yname == "uc" { return "Uc" }
    if yname == "pre-fec-val" { return "PreFecVal" }
    if yname == "pre-fec-mantissa" { return "PreFecMantissa" }
    if yname == "ec-value" { return "EcValue" }
    if yname == "uc-value" { return "UcValue" }
    if yname == "pre-fec-ber-value" { return "PreFecBerValue" }
    if yname == "pre-fec-ber-mantissa" { return "PreFecBerMantissa" }
    if yname == "nv-optical-support" { return "NvOpticalSupport" }
    if yname == "gmpls-tti-mode" { return "GmplsTtiMode" }
    if yname == "gmpls-tvm-id" { return "GmplsTvmId" }
    if yname == "auto-tti-flag" { return "AutoTtiFlag" }
    if yname == "description" { return "Description" }
    if yname == "local" { return "Local" }
    if yname == "remote" { return "Remote" }
    if yname == "tti-mode" { return "TtiMode" }
    if yname == "network-srlg" { return "NetworkSrlg" }
    if yname == "otu-alarm-info" { return "OtuAlarmInfo" }
    if yname == "proactive" { return "Proactive" }
    if yname == "otu-fec-satistics" { return "OtuFecSatistics" }
    return ""
}

func (info *Otu_Controllers_Controller_Info) GetSegmentPath() string {
    return "info"
}

func (info *Otu_Controllers_Controller_Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local" {
        return &info.Local
    }
    if childYangName == "remote" {
        return &info.Remote
    }
    if childYangName == "tti-mode" {
        return &info.TtiMode
    }
    if childYangName == "network-srlg" {
        return &info.NetworkSrlg
    }
    if childYangName == "otu-alarm-info" {
        return &info.OtuAlarmInfo
    }
    if childYangName == "proactive" {
        return &info.Proactive
    }
    if childYangName == "otu-fec-satistics" {
        return &info.OtuFecSatistics
    }
    return nil
}

func (info *Otu_Controllers_Controller_Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local"] = &info.Local
    children["remote"] = &info.Remote
    children["tti-mode"] = &info.TtiMode
    children["network-srlg"] = &info.NetworkSrlg
    children["otu-alarm-info"] = &info.OtuAlarmInfo
    children["proactive"] = &info.Proactive
    children["otu-fec-satistics"] = &info.OtuFecSatistics
    return children
}

func (info *Otu_Controllers_Controller_Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = info.State
    leafs["name"] = info.Name
    leafs["sf"] = info.Sf
    leafs["sd"] = info.Sd
    leafs["loopback-mode"] = info.LoopbackMode
    leafs["fec-mode"] = info.FecMode
    leafs["derivedstate-mode"] = info.DerivedstateMode
    leafs["inherit-sec-state"] = info.InheritSecState
    leafs["config-sec-state"] = info.ConfigSecState
    leafs["gcc-mode"] = info.GccMode
    leafs["q"] = info.Q
    leafs["q-margin"] = info.QMargin
    leafs["performance-monitoring"] = info.PerformanceMonitoring
    leafs["ec"] = info.Ec
    leafs["uc"] = info.Uc
    leafs["pre-fec-val"] = info.PreFecVal
    leafs["pre-fec-mantissa"] = info.PreFecMantissa
    leafs["ec-value"] = info.EcValue
    leafs["uc-value"] = info.UcValue
    leafs["pre-fec-ber-value"] = info.PreFecBerValue
    leafs["pre-fec-ber-mantissa"] = info.PreFecBerMantissa
    leafs["nv-optical-support"] = info.NvOpticalSupport
    leafs["gmpls-tti-mode"] = info.GmplsTtiMode
    leafs["gmpls-tvm-id"] = info.GmplsTvmId
    leafs["auto-tti-flag"] = info.AutoTtiFlag
    leafs["description"] = info.Description
    return leafs
}

func (info *Otu_Controllers_Controller_Info) GetBundleName() string { return "cisco_ios_xr" }

func (info *Otu_Controllers_Controller_Info) GetYangName() string { return "info" }

func (info *Otu_Controllers_Controller_Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (info *Otu_Controllers_Controller_Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (info *Otu_Controllers_Controller_Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (info *Otu_Controllers_Controller_Info) SetParent(parent types.Entity) { info.parent = parent }

func (info *Otu_Controllers_Controller_Info) GetParent() types.Entity { return info.parent }

func (info *Otu_Controllers_Controller_Info) GetParentYangName() string { return "controller" }

// Otu_Controllers_Controller_Info_Local
// TTI
type Otu_Controllers_Controller_Info_Local struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router ID. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (local *Otu_Controllers_Controller_Info_Local) GetFilter() yfilter.YFilter { return local.YFilter }

func (local *Otu_Controllers_Controller_Info_Local) SetFilter(yf yfilter.YFilter) { local.YFilter = yf }

func (local *Otu_Controllers_Controller_Info_Local) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "if-index" { return "IfIndex" }
    return ""
}

func (local *Otu_Controllers_Controller_Info_Local) GetSegmentPath() string {
    return "local"
}

func (local *Otu_Controllers_Controller_Info_Local) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (local *Otu_Controllers_Controller_Info_Local) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (local *Otu_Controllers_Controller_Info_Local) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = local.RouterId
    leafs["if-index"] = local.IfIndex
    return leafs
}

func (local *Otu_Controllers_Controller_Info_Local) GetBundleName() string { return "cisco_ios_xr" }

func (local *Otu_Controllers_Controller_Info_Local) GetYangName() string { return "local" }

func (local *Otu_Controllers_Controller_Info_Local) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (local *Otu_Controllers_Controller_Info_Local) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (local *Otu_Controllers_Controller_Info_Local) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (local *Otu_Controllers_Controller_Info_Local) SetParent(parent types.Entity) { local.parent = parent }

func (local *Otu_Controllers_Controller_Info_Local) GetParent() types.Entity { return local.parent }

func (local *Otu_Controllers_Controller_Info_Local) GetParentYangName() string { return "info" }

// Otu_Controllers_Controller_Info_Remote
// Remote
type Otu_Controllers_Controller_Info_Remote struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router ID. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (remote *Otu_Controllers_Controller_Info_Remote) GetFilter() yfilter.YFilter { return remote.YFilter }

func (remote *Otu_Controllers_Controller_Info_Remote) SetFilter(yf yfilter.YFilter) { remote.YFilter = yf }

func (remote *Otu_Controllers_Controller_Info_Remote) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "if-index" { return "IfIndex" }
    return ""
}

func (remote *Otu_Controllers_Controller_Info_Remote) GetSegmentPath() string {
    return "remote"
}

func (remote *Otu_Controllers_Controller_Info_Remote) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remote *Otu_Controllers_Controller_Info_Remote) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remote *Otu_Controllers_Controller_Info_Remote) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = remote.RouterId
    leafs["if-index"] = remote.IfIndex
    return leafs
}

func (remote *Otu_Controllers_Controller_Info_Remote) GetBundleName() string { return "cisco_ios_xr" }

func (remote *Otu_Controllers_Controller_Info_Remote) GetYangName() string { return "remote" }

func (remote *Otu_Controllers_Controller_Info_Remote) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remote *Otu_Controllers_Controller_Info_Remote) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remote *Otu_Controllers_Controller_Info_Remote) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remote *Otu_Controllers_Controller_Info_Remote) SetParent(parent types.Entity) { remote.parent = parent }

func (remote *Otu_Controllers_Controller_Info_Remote) GetParent() types.Entity { return remote.parent }

func (remote *Otu_Controllers_Controller_Info_Remote) GetParentYangName() string { return "info" }

// Otu_Controllers_Controller_Info_TtiMode
// OTU TTI
type Otu_Controllers_Controller_Info_TtiMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // G709TTI sent. The type is OtuTtiEt.
    G709TtiSentMode interface{}

    // G709TTI Expected. The type is OtuTtiEt.
    G709TtiExpMode interface{}

    // G709TTI Recieved. The type is OtuTtiEt.
    G709TtiRecMode interface{}

    // Remote Interface Name. The type is string.
    RemoteInterface interface{}

    // Remote host name. The type is string.
    RemoteHostName interface{}

    // Remote host ip. The type is string.
    RemoteIpAddr interface{}

    // String Sent.
    Tx Otu_Controllers_Controller_Info_TtiMode_Tx

    // String Expected.
    Exp Otu_Controllers_Controller_Info_TtiMode_Exp

    // String Received.
    Rec Otu_Controllers_Controller_Info_TtiMode_Rec
}

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetFilter() yfilter.YFilter { return ttiMode.YFilter }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) SetFilter(yf yfilter.YFilter) { ttiMode.YFilter = yf }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetGoName(yname string) string {
    if yname == "g709tti-sent-mode" { return "G709TtiSentMode" }
    if yname == "g709tti-exp-mode" { return "G709TtiExpMode" }
    if yname == "g709tti-rec-mode" { return "G709TtiRecMode" }
    if yname == "remote-interface" { return "RemoteInterface" }
    if yname == "remote-host-name" { return "RemoteHostName" }
    if yname == "remote-ip-addr" { return "RemoteIpAddr" }
    if yname == "tx" { return "Tx" }
    if yname == "exp" { return "Exp" }
    if yname == "rec" { return "Rec" }
    return ""
}

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetSegmentPath() string {
    return "tti-mode"
}

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx" {
        return &ttiMode.Tx
    }
    if childYangName == "exp" {
        return &ttiMode.Exp
    }
    if childYangName == "rec" {
        return &ttiMode.Rec
    }
    return nil
}

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx"] = &ttiMode.Tx
    children["exp"] = &ttiMode.Exp
    children["rec"] = &ttiMode.Rec
    return children
}

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["g709tti-sent-mode"] = ttiMode.G709TtiSentMode
    leafs["g709tti-exp-mode"] = ttiMode.G709TtiExpMode
    leafs["g709tti-rec-mode"] = ttiMode.G709TtiRecMode
    leafs["remote-interface"] = ttiMode.RemoteInterface
    leafs["remote-host-name"] = ttiMode.RemoteHostName
    leafs["remote-ip-addr"] = ttiMode.RemoteIpAddr
    return leafs
}

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetBundleName() string { return "cisco_ios_xr" }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetYangName() string { return "tti-mode" }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) SetParent(parent types.Entity) { ttiMode.parent = parent }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetParent() types.Entity { return ttiMode.parent }

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetParentYangName() string { return "info" }

// Otu_Controllers_Controller_Info_TtiMode_Tx
// String Sent
type Otu_Controllers_Controller_Info_TtiMode_Tx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // full tti ascii String . The type is string.
    FullTtiAsciiString interface{}

    // tx String . The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String . The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String . The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetFilter() yfilter.YFilter { return tx.YFilter }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) SetFilter(yf yfilter.YFilter) { tx.YFilter = yf }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetGoName(yname string) string {
    if yname == "full-tti-ascii-string" { return "FullTtiAsciiString" }
    if yname == "sapi" { return "Sapi" }
    if yname == "dapi" { return "Dapi" }
    if yname == "operator-specific" { return "OperatorSpecific" }
    return ""
}

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetSegmentPath() string {
    return "tx"
}

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["full-tti-ascii-string"] = tx.FullTtiAsciiString
    leafs["sapi"] = tx.Sapi
    leafs["dapi"] = tx.Dapi
    leafs["operator-specific"] = tx.OperatorSpecific
    return leafs
}

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetBundleName() string { return "cisco_ios_xr" }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetYangName() string { return "tx" }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) SetParent(parent types.Entity) { tx.parent = parent }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetParent() types.Entity { return tx.parent }

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetParentYangName() string { return "tti-mode" }

// Otu_Controllers_Controller_Info_TtiMode_Exp
// String Expected
type Otu_Controllers_Controller_Info_TtiMode_Exp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // full tti ascii String . The type is string.
    FullTtiAsciiString interface{}

    // tx String . The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String . The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String . The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetFilter() yfilter.YFilter { return exp.YFilter }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) SetFilter(yf yfilter.YFilter) { exp.YFilter = yf }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetGoName(yname string) string {
    if yname == "full-tti-ascii-string" { return "FullTtiAsciiString" }
    if yname == "sapi" { return "Sapi" }
    if yname == "dapi" { return "Dapi" }
    if yname == "operator-specific" { return "OperatorSpecific" }
    return ""
}

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetSegmentPath() string {
    return "exp"
}

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["full-tti-ascii-string"] = exp.FullTtiAsciiString
    leafs["sapi"] = exp.Sapi
    leafs["dapi"] = exp.Dapi
    leafs["operator-specific"] = exp.OperatorSpecific
    return leafs
}

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetBundleName() string { return "cisco_ios_xr" }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetYangName() string { return "exp" }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) SetParent(parent types.Entity) { exp.parent = parent }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetParent() types.Entity { return exp.parent }

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetParentYangName() string { return "tti-mode" }

// Otu_Controllers_Controller_Info_TtiMode_Rec
// String Received
type Otu_Controllers_Controller_Info_TtiMode_Rec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // full tti ascii String . The type is string.
    FullTtiAsciiString interface{}

    // tx String . The type is slice of interface{} with range: 0..255.
    Sapi []interface{}

    // exp String . The type is slice of interface{} with range: 0..255.
    Dapi []interface{}

    // rec String . The type is slice of interface{} with range: 0..255.
    OperatorSpecific []interface{}
}

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetFilter() yfilter.YFilter { return rec.YFilter }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) SetFilter(yf yfilter.YFilter) { rec.YFilter = yf }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetGoName(yname string) string {
    if yname == "full-tti-ascii-string" { return "FullTtiAsciiString" }
    if yname == "sapi" { return "Sapi" }
    if yname == "dapi" { return "Dapi" }
    if yname == "operator-specific" { return "OperatorSpecific" }
    return ""
}

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetSegmentPath() string {
    return "rec"
}

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["full-tti-ascii-string"] = rec.FullTtiAsciiString
    leafs["sapi"] = rec.Sapi
    leafs["dapi"] = rec.Dapi
    leafs["operator-specific"] = rec.OperatorSpecific
    return leafs
}

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetBundleName() string { return "cisco_ios_xr" }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetYangName() string { return "rec" }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) SetParent(parent types.Entity) { rec.parent = parent }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetParent() types.Entity { return rec.parent }

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetParentYangName() string { return "tti-mode" }

// Otu_Controllers_Controller_Info_NetworkSrlg
// Network Shared Risk Link Group information
type Otu_Controllers_Controller_Info_NetworkSrlg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array of Network Shared Risk Link Group information. The type is slice of
    // Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo.
    SrlgInfo []Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo
}

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetFilter() yfilter.YFilter { return networkSrlg.YFilter }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) SetFilter(yf yfilter.YFilter) { networkSrlg.YFilter = yf }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetGoName(yname string) string {
    if yname == "srlg-info" { return "SrlgInfo" }
    return ""
}

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetSegmentPath() string {
    return "network-srlg"
}

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-info" {
        for _, c := range networkSrlg.SrlgInfo {
            if networkSrlg.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo{}
        networkSrlg.SrlgInfo = append(networkSrlg.SrlgInfo, child)
        return &networkSrlg.SrlgInfo[len(networkSrlg.SrlgInfo)-1]
    }
    return nil
}

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkSrlg.SrlgInfo {
        children[networkSrlg.SrlgInfo[i].GetSegmentPath()] = &networkSrlg.SrlgInfo[i]
    }
    return children
}

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetBundleName() string { return "cisco_ios_xr" }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetYangName() string { return "network-srlg" }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) SetParent(parent types.Entity) { networkSrlg.parent = parent }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetParent() types.Entity { return networkSrlg.parent }

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetParentYangName() string { return "info" }

// Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo
// Array of Network Shared Risk Link Group
// information
type Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array to maintain set id number. The type is interface{} with range:
    // 0..4294967295.
    SetId interface{}

    // Shared Risk Link Group information expressed in integer format. The type is
    // slice of interface{} with range: 0..4294967295.
    Srlg []interface{}
}

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetFilter() yfilter.YFilter { return srlgInfo.YFilter }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) SetFilter(yf yfilter.YFilter) { srlgInfo.YFilter = yf }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetGoName(yname string) string {
    if yname == "set-id" { return "SetId" }
    if yname == "srlg" { return "Srlg" }
    return ""
}

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetSegmentPath() string {
    return "srlg-info"
}

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["set-id"] = srlgInfo.SetId
    leafs["srlg"] = srlgInfo.Srlg
    return leafs
}

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetBundleName() string { return "cisco_ios_xr" }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetYangName() string { return "srlg-info" }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) SetParent(parent types.Entity) { srlgInfo.parent = parent }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetParent() types.Entity { return srlgInfo.parent }

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetParentYangName() string { return "network-srlg" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo
// OTU layer alarm Information
type Otu_Controllers_Controller_Info_OtuAlarmInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Loss of Signal.
    Los Otu_Controllers_Controller_Info_OtuAlarmInfo_Los

    // Loss of Frame.
    Lof Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof

    // Loss of MultiFrame.
    Lom Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom

    // Out of Frame.
    Oof Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof

    // Out of MultiFrame.
    Oom Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom

    // Alarm Indication Signal.
    Ais Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais

    // Incoming Alignment Error.
    Iae Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae

    // Backward Incoming Alignment Error.
    Biae Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae

    // Backward Defect Indication.
    Bdi Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi

    // Trace Identifier Mismatch.
    Tim Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim

    // GCC End of Channel.
    Eoc Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc

    // FEC mismatch alarm.
    FecMismatch Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch

    // SF BER alarm.
    SfBer Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer

    // SD BER alarm.
    SdBer Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer

    // EC alarm.
    Ec Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec

    // UC alarm.
    Uc Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc

    // FEC UnCorrected Word.
    Fecunc Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc
}

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetFilter() yfilter.YFilter { return otuAlarmInfo.YFilter }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) SetFilter(yf yfilter.YFilter) { otuAlarmInfo.YFilter = yf }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetGoName(yname string) string {
    if yname == "los" { return "Los" }
    if yname == "lof" { return "Lof" }
    if yname == "lom" { return "Lom" }
    if yname == "oof" { return "Oof" }
    if yname == "oom" { return "Oom" }
    if yname == "ais" { return "Ais" }
    if yname == "iae" { return "Iae" }
    if yname == "biae" { return "Biae" }
    if yname == "bdi" { return "Bdi" }
    if yname == "tim" { return "Tim" }
    if yname == "eoc" { return "Eoc" }
    if yname == "fec-mismatch" { return "FecMismatch" }
    if yname == "sf-ber" { return "SfBer" }
    if yname == "sd-ber" { return "SdBer" }
    if yname == "ec" { return "Ec" }
    if yname == "uc" { return "Uc" }
    if yname == "fecunc" { return "Fecunc" }
    return ""
}

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetSegmentPath() string {
    return "otu-alarm-info"
}

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "los" {
        return &otuAlarmInfo.Los
    }
    if childYangName == "lof" {
        return &otuAlarmInfo.Lof
    }
    if childYangName == "lom" {
        return &otuAlarmInfo.Lom
    }
    if childYangName == "oof" {
        return &otuAlarmInfo.Oof
    }
    if childYangName == "oom" {
        return &otuAlarmInfo.Oom
    }
    if childYangName == "ais" {
        return &otuAlarmInfo.Ais
    }
    if childYangName == "iae" {
        return &otuAlarmInfo.Iae
    }
    if childYangName == "biae" {
        return &otuAlarmInfo.Biae
    }
    if childYangName == "bdi" {
        return &otuAlarmInfo.Bdi
    }
    if childYangName == "tim" {
        return &otuAlarmInfo.Tim
    }
    if childYangName == "eoc" {
        return &otuAlarmInfo.Eoc
    }
    if childYangName == "fec-mismatch" {
        return &otuAlarmInfo.FecMismatch
    }
    if childYangName == "sf-ber" {
        return &otuAlarmInfo.SfBer
    }
    if childYangName == "sd-ber" {
        return &otuAlarmInfo.SdBer
    }
    if childYangName == "ec" {
        return &otuAlarmInfo.Ec
    }
    if childYangName == "uc" {
        return &otuAlarmInfo.Uc
    }
    if childYangName == "fecunc" {
        return &otuAlarmInfo.Fecunc
    }
    return nil
}

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["los"] = &otuAlarmInfo.Los
    children["lof"] = &otuAlarmInfo.Lof
    children["lom"] = &otuAlarmInfo.Lom
    children["oof"] = &otuAlarmInfo.Oof
    children["oom"] = &otuAlarmInfo.Oom
    children["ais"] = &otuAlarmInfo.Ais
    children["iae"] = &otuAlarmInfo.Iae
    children["biae"] = &otuAlarmInfo.Biae
    children["bdi"] = &otuAlarmInfo.Bdi
    children["tim"] = &otuAlarmInfo.Tim
    children["eoc"] = &otuAlarmInfo.Eoc
    children["fec-mismatch"] = &otuAlarmInfo.FecMismatch
    children["sf-ber"] = &otuAlarmInfo.SfBer
    children["sd-ber"] = &otuAlarmInfo.SdBer
    children["ec"] = &otuAlarmInfo.Ec
    children["uc"] = &otuAlarmInfo.Uc
    children["fecunc"] = &otuAlarmInfo.Fecunc
    return children
}

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetYangName() string { return "otu-alarm-info" }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) SetParent(parent types.Entity) { otuAlarmInfo.parent = parent }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetParent() types.Entity { return otuAlarmInfo.parent }

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetParentYangName() string { return "info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Los
// Loss of Signal
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Los struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetFilter() yfilter.YFilter { return los.YFilter }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) SetFilter(yf yfilter.YFilter) { los.YFilter = yf }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetSegmentPath() string {
    return "los"
}

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = los.ReportingEnabled
    leafs["is-detected"] = los.IsDetected
    leafs["is-asserted"] = los.IsAsserted
    leafs["counter"] = los.Counter
    return leafs
}

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetBundleName() string { return "cisco_ios_xr" }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetYangName() string { return "los" }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) SetParent(parent types.Entity) { los.parent = parent }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetParent() types.Entity { return los.parent }

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof
// Loss of Frame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetFilter() yfilter.YFilter { return lof.YFilter }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) SetFilter(yf yfilter.YFilter) { lof.YFilter = yf }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetSegmentPath() string {
    return "lof"
}

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = lof.ReportingEnabled
    leafs["is-detected"] = lof.IsDetected
    leafs["is-asserted"] = lof.IsAsserted
    leafs["counter"] = lof.Counter
    return leafs
}

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetBundleName() string { return "cisco_ios_xr" }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetYangName() string { return "lof" }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) SetParent(parent types.Entity) { lof.parent = parent }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetParent() types.Entity { return lof.parent }

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom
// Loss of MultiFrame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetFilter() yfilter.YFilter { return lom.YFilter }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) SetFilter(yf yfilter.YFilter) { lom.YFilter = yf }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetSegmentPath() string {
    return "lom"
}

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = lom.ReportingEnabled
    leafs["is-detected"] = lom.IsDetected
    leafs["is-asserted"] = lom.IsAsserted
    leafs["counter"] = lom.Counter
    return leafs
}

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetBundleName() string { return "cisco_ios_xr" }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetYangName() string { return "lom" }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) SetParent(parent types.Entity) { lom.parent = parent }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetParent() types.Entity { return lom.parent }

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof
// Out of Frame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetFilter() yfilter.YFilter { return oof.YFilter }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) SetFilter(yf yfilter.YFilter) { oof.YFilter = yf }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetSegmentPath() string {
    return "oof"
}

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = oof.ReportingEnabled
    leafs["is-detected"] = oof.IsDetected
    leafs["is-asserted"] = oof.IsAsserted
    leafs["counter"] = oof.Counter
    return leafs
}

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetBundleName() string { return "cisco_ios_xr" }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetYangName() string { return "oof" }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) SetParent(parent types.Entity) { oof.parent = parent }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetParent() types.Entity { return oof.parent }

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom
// Out of MultiFrame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetFilter() yfilter.YFilter { return oom.YFilter }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) SetFilter(yf yfilter.YFilter) { oom.YFilter = yf }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetSegmentPath() string {
    return "oom"
}

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = oom.ReportingEnabled
    leafs["is-detected"] = oom.IsDetected
    leafs["is-asserted"] = oom.IsAsserted
    leafs["counter"] = oom.Counter
    return leafs
}

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetBundleName() string { return "cisco_ios_xr" }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetYangName() string { return "oom" }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) SetParent(parent types.Entity) { oom.parent = parent }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetParent() types.Entity { return oom.parent }

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais
// Alarm Indication Signal
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetFilter() yfilter.YFilter { return ais.YFilter }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) SetFilter(yf yfilter.YFilter) { ais.YFilter = yf }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetSegmentPath() string {
    return "ais"
}

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = ais.ReportingEnabled
    leafs["is-detected"] = ais.IsDetected
    leafs["is-asserted"] = ais.IsAsserted
    leafs["counter"] = ais.Counter
    return leafs
}

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetBundleName() string { return "cisco_ios_xr" }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetYangName() string { return "ais" }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) SetParent(parent types.Entity) { ais.parent = parent }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetParent() types.Entity { return ais.parent }

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae
// Incoming Alignment Error
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetFilter() yfilter.YFilter { return iae.YFilter }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) SetFilter(yf yfilter.YFilter) { iae.YFilter = yf }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetSegmentPath() string {
    return "iae"
}

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = iae.ReportingEnabled
    leafs["is-detected"] = iae.IsDetected
    leafs["is-asserted"] = iae.IsAsserted
    leafs["counter"] = iae.Counter
    return leafs
}

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetBundleName() string { return "cisco_ios_xr" }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetYangName() string { return "iae" }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) SetParent(parent types.Entity) { iae.parent = parent }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetParent() types.Entity { return iae.parent }

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae
// Backward Incoming Alignment Error
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetFilter() yfilter.YFilter { return biae.YFilter }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) SetFilter(yf yfilter.YFilter) { biae.YFilter = yf }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetSegmentPath() string {
    return "biae"
}

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = biae.ReportingEnabled
    leafs["is-detected"] = biae.IsDetected
    leafs["is-asserted"] = biae.IsAsserted
    leafs["counter"] = biae.Counter
    return leafs
}

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetBundleName() string { return "cisco_ios_xr" }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetYangName() string { return "biae" }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) SetParent(parent types.Entity) { biae.parent = parent }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetParent() types.Entity { return biae.parent }

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi
// Backward Defect Indication
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetFilter() yfilter.YFilter { return bdi.YFilter }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) SetFilter(yf yfilter.YFilter) { bdi.YFilter = yf }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetSegmentPath() string {
    return "bdi"
}

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = bdi.ReportingEnabled
    leafs["is-detected"] = bdi.IsDetected
    leafs["is-asserted"] = bdi.IsAsserted
    leafs["counter"] = bdi.Counter
    return leafs
}

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetBundleName() string { return "cisco_ios_xr" }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetYangName() string { return "bdi" }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) SetParent(parent types.Entity) { bdi.parent = parent }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetParent() types.Entity { return bdi.parent }

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim
// Trace Identifier Mismatch
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetFilter() yfilter.YFilter { return tim.YFilter }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) SetFilter(yf yfilter.YFilter) { tim.YFilter = yf }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetSegmentPath() string {
    return "tim"
}

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = tim.ReportingEnabled
    leafs["is-detected"] = tim.IsDetected
    leafs["is-asserted"] = tim.IsAsserted
    leafs["counter"] = tim.Counter
    return leafs
}

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetBundleName() string { return "cisco_ios_xr" }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetYangName() string { return "tim" }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) SetParent(parent types.Entity) { tim.parent = parent }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetParent() types.Entity { return tim.parent }

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc
// GCC End of Channel
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetFilter() yfilter.YFilter { return eoc.YFilter }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) SetFilter(yf yfilter.YFilter) { eoc.YFilter = yf }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetSegmentPath() string {
    return "eoc"
}

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = eoc.ReportingEnabled
    leafs["is-detected"] = eoc.IsDetected
    leafs["is-asserted"] = eoc.IsAsserted
    leafs["counter"] = eoc.Counter
    return leafs
}

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetBundleName() string { return "cisco_ios_xr" }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetYangName() string { return "eoc" }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) SetParent(parent types.Entity) { eoc.parent = parent }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetParent() types.Entity { return eoc.parent }

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch
// FEC mismatch alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetFilter() yfilter.YFilter { return fecMismatch.YFilter }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) SetFilter(yf yfilter.YFilter) { fecMismatch.YFilter = yf }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetSegmentPath() string {
    return "fec-mismatch"
}

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = fecMismatch.ReportingEnabled
    leafs["is-detected"] = fecMismatch.IsDetected
    leafs["is-asserted"] = fecMismatch.IsAsserted
    leafs["counter"] = fecMismatch.Counter
    return leafs
}

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetBundleName() string { return "cisco_ios_xr" }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetYangName() string { return "fec-mismatch" }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) SetParent(parent types.Entity) { fecMismatch.parent = parent }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetParent() types.Entity { return fecMismatch.parent }

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer
// SF BER alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetFilter() yfilter.YFilter { return sfBer.YFilter }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) SetFilter(yf yfilter.YFilter) { sfBer.YFilter = yf }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetSegmentPath() string {
    return "sf-ber"
}

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = sfBer.ReportingEnabled
    leafs["is-detected"] = sfBer.IsDetected
    leafs["is-asserted"] = sfBer.IsAsserted
    leafs["counter"] = sfBer.Counter
    return leafs
}

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetBundleName() string { return "cisco_ios_xr" }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetYangName() string { return "sf-ber" }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) SetParent(parent types.Entity) { sfBer.parent = parent }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetParent() types.Entity { return sfBer.parent }

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer
// SD BER alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetFilter() yfilter.YFilter { return sdBer.YFilter }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) SetFilter(yf yfilter.YFilter) { sdBer.YFilter = yf }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetSegmentPath() string {
    return "sd-ber"
}

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = sdBer.ReportingEnabled
    leafs["is-detected"] = sdBer.IsDetected
    leafs["is-asserted"] = sdBer.IsAsserted
    leafs["counter"] = sdBer.Counter
    return leafs
}

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetBundleName() string { return "cisco_ios_xr" }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetYangName() string { return "sd-ber" }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) SetParent(parent types.Entity) { sdBer.parent = parent }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetParent() types.Entity { return sdBer.parent }

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec
// EC alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetFilter() yfilter.YFilter { return ec.YFilter }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) SetFilter(yf yfilter.YFilter) { ec.YFilter = yf }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetSegmentPath() string {
    return "ec"
}

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = ec.ReportingEnabled
    leafs["is-detected"] = ec.IsDetected
    leafs["is-asserted"] = ec.IsAsserted
    leafs["counter"] = ec.Counter
    return leafs
}

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetBundleName() string { return "cisco_ios_xr" }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetYangName() string { return "ec" }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) SetParent(parent types.Entity) { ec.parent = parent }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetParent() types.Entity { return ec.parent }

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc
// UC alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetFilter() yfilter.YFilter { return uc.YFilter }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) SetFilter(yf yfilter.YFilter) { uc.YFilter = yf }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetSegmentPath() string {
    return "uc"
}

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = uc.ReportingEnabled
    leafs["is-detected"] = uc.IsDetected
    leafs["is-asserted"] = uc.IsAsserted
    leafs["counter"] = uc.Counter
    return leafs
}

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetBundleName() string { return "cisco_ios_xr" }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetYangName() string { return "uc" }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) SetParent(parent types.Entity) { uc.parent = parent }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetParent() types.Entity { return uc.parent }

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc
// FEC UnCorrected Word
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetFilter() yfilter.YFilter { return fecunc.YFilter }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) SetFilter(yf yfilter.YFilter) { fecunc.YFilter = yf }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetSegmentPath() string {
    return "fecunc"
}

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = fecunc.ReportingEnabled
    leafs["is-detected"] = fecunc.IsDetected
    leafs["is-asserted"] = fecunc.IsAsserted
    leafs["counter"] = fecunc.Counter
    return leafs
}

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetBundleName() string { return "cisco_ios_xr" }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetYangName() string { return "fecunc" }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) SetParent(parent types.Entity) { fecunc.parent = parent }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetParent() types.Entity { return fecunc.parent }

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetParentYangName() string { return "otu-alarm-info" }

// Otu_Controllers_Controller_Info_Proactive
// Proactive Protection
type Otu_Controllers_Controller_Info_Proactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Proactive Status. The type is bool.
    ProactiveStatus interface{}

    // Secondary Admin State. The type is OtuSecState.
    InheritSecState interface{}

    // Sec State. The type is OtuSecState.
    ConfigSecState interface{}

    // Proactive FSM State. The type is OtuPpFsmState.
    ProactiveFsmState interface{}

    // Proactive FSM IF State. The type is OtuPpIntfState.
    ProactiveFsmIfState interface{}

    // Trigger threshold coefficient. The type is interface{} with range: 0..255.
    TrigThreshCoeff interface{}

    // Trigger threshold power. The type is interface{} with range: 0..255.
    TrigThreshPower interface{}

    // Revert threshold coefficient. The type is interface{} with range: 0..255.
    RvrtThreshCoeff interface{}

    // Revert threshold power. The type is interface{} with range: 0..255.
    RvrtThreshPower interface{}

    // Trigger Integration window. The type is interface{} with range:
    // 0..4294967295.
    TriggerWindow interface{}

    // Revert Integration Window. The type is interface{} with range:
    // 0..4294967295.
    RevertWindow interface{}
}

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetFilter() yfilter.YFilter { return proactive.YFilter }

func (proactive *Otu_Controllers_Controller_Info_Proactive) SetFilter(yf yfilter.YFilter) { proactive.YFilter = yf }

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetGoName(yname string) string {
    if yname == "proactive-status" { return "ProactiveStatus" }
    if yname == "inherit-sec-state" { return "InheritSecState" }
    if yname == "config-sec-state" { return "ConfigSecState" }
    if yname == "proactive-fsm-state" { return "ProactiveFsmState" }
    if yname == "proactive-fsm-if-state" { return "ProactiveFsmIfState" }
    if yname == "trig-thresh-coeff" { return "TrigThreshCoeff" }
    if yname == "trig-thresh-power" { return "TrigThreshPower" }
    if yname == "rvrt-thresh-coeff" { return "RvrtThreshCoeff" }
    if yname == "rvrt-thresh-power" { return "RvrtThreshPower" }
    if yname == "trigger-window" { return "TriggerWindow" }
    if yname == "revert-window" { return "RevertWindow" }
    return ""
}

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetSegmentPath() string {
    return "proactive"
}

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proactive-status"] = proactive.ProactiveStatus
    leafs["inherit-sec-state"] = proactive.InheritSecState
    leafs["config-sec-state"] = proactive.ConfigSecState
    leafs["proactive-fsm-state"] = proactive.ProactiveFsmState
    leafs["proactive-fsm-if-state"] = proactive.ProactiveFsmIfState
    leafs["trig-thresh-coeff"] = proactive.TrigThreshCoeff
    leafs["trig-thresh-power"] = proactive.TrigThreshPower
    leafs["rvrt-thresh-coeff"] = proactive.RvrtThreshCoeff
    leafs["rvrt-thresh-power"] = proactive.RvrtThreshPower
    leafs["trigger-window"] = proactive.TriggerWindow
    leafs["revert-window"] = proactive.RevertWindow
    return leafs
}

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetBundleName() string { return "cisco_ios_xr" }

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetYangName() string { return "proactive" }

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proactive *Otu_Controllers_Controller_Info_Proactive) SetParent(parent types.Entity) { proactive.parent = parent }

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetParent() types.Entity { return proactive.parent }

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetParentYangName() string { return "info" }

// Otu_Controllers_Controller_Info_OtuFecSatistics
// OTU FEC Statistics
type Otu_Controllers_Controller_Info_OtuFecSatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bit Error Rate After Forward Error Correction. The type is string.
    PostFecBer interface{}

    // Bit Error Rate Before Forward Error Correction. The type is string.
    PreFecBer interface{}
}

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetFilter() yfilter.YFilter { return otuFecSatistics.YFilter }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) SetFilter(yf yfilter.YFilter) { otuFecSatistics.YFilter = yf }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetGoName(yname string) string {
    if yname == "post-fec-ber" { return "PostFecBer" }
    if yname == "pre-fec-ber" { return "PreFecBer" }
    return ""
}

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetSegmentPath() string {
    return "otu-fec-satistics"
}

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["post-fec-ber"] = otuFecSatistics.PostFecBer
    leafs["pre-fec-ber"] = otuFecSatistics.PreFecBer
    return leafs
}

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetBundleName() string { return "cisco_ios_xr" }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetYangName() string { return "otu-fec-satistics" }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) SetParent(parent types.Entity) { otuFecSatistics.parent = parent }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetParent() types.Entity { return otuFecSatistics.parent }

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetParentYangName() string { return "info" }

