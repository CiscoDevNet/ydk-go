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

// OtuG709fecMode represents Otu g709fec mode
type OtuG709fecMode string

const (
    // NONE
    OtuG709fecMode_otu_bag_none_fec OtuG709fecMode = "otu-bag-none-fec"

    // STANDARD
    OtuG709fecMode_otu_bag_standard_fec OtuG709fecMode = "otu-bag-standard-fec"

    // ENHANCEDI.7
    OtuG709fecMode_otu_bag_1_i_7_fec OtuG709fecMode = "otu-bag-1-i-7-fec"

    // ENHANCEDI.4
    OtuG709fecMode_otu_bag_1_i_4_fec OtuG709fecMode = "otu-bag-1-i-4-fec"

    // SWIZZLE
    OtuG709fecMode_otu_bag_swizzle_fec OtuG709fecMode = "otu-bag-swizzle-fec"

    // HIGH GAIN20
    OtuG709fecMode_otu_bag_hg20_fec OtuG709fecMode = "otu-bag-hg20-fec"

    // Enhanced High Gain 7
    OtuG709fecMode_otu_bag_enhanced_hg7_fec OtuG709fecMode = "otu-bag-enhanced-hg7-fec"

    // Soft-Decision 20
    OtuG709fecMode_otu_bag_sd20_fec OtuG709fecMode = "otu-bag-sd20-fec"

    // Soft-Decision 7
    OtuG709fecMode_otu_bag_sd7_fec OtuG709fecMode = "otu-bag-sd7-fec"

    // ALL
    OtuG709fecMode_otu_bag_all_fec OtuG709fecMode = "otu-bag-all-fec"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All OTU Port operational data.
    Controllers Otu_Controllers
}

func (otu *Otu) GetEntityData() *types.CommonEntityData {
    otu.EntityData.YFilter = otu.YFilter
    otu.EntityData.YangName = "otu"
    otu.EntityData.BundleName = "cisco_ios_xr"
    otu.EntityData.ParentYangName = "Cisco-IOS-XR-controller-otu-oper"
    otu.EntityData.SegmentPath = "Cisco-IOS-XR-controller-otu-oper:otu"
    otu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otu.EntityData.Children = types.NewOrderedMap()
    otu.EntityData.Children.Append("controllers", types.YChild{"Controllers", &otu.Controllers})
    otu.EntityData.Leafs = types.NewOrderedMap()

    otu.EntityData.YListKeys = []string {}

    return &(otu.EntityData)
}

// Otu_Controllers
// All OTU Port operational data
type Otu_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OTU Port operational data. The type is slice of Otu_Controllers_Controller.
    Controller []*Otu_Controllers_Controller
}

func (controllers *Otu_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "otu"
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

// Otu_Controllers_Controller
// OTU Port operational data
type Otu_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // OTU port PRBS operational data.
    Prbs Otu_Controllers_Controller_Prbs

    // OTU port operational data.
    Info Otu_Controllers_Controller_Info
}

func (controller *Otu_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("prbs", types.YChild{"Prbs", &controller.Prbs})
    controller.EntityData.Children.Append("info", types.YChild{"Info", &controller.Info})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// Otu_Controllers_Controller_Prbs
// OTU port PRBS operational data
type Otu_Controllers_Controller_Prbs struct {
    EntityData types.CommonEntityData
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

func (prbs *Otu_Controllers_Controller_Prbs) GetEntityData() *types.CommonEntityData {
    prbs.EntityData.YFilter = prbs.YFilter
    prbs.EntityData.YangName = "prbs"
    prbs.EntityData.BundleName = "cisco_ios_xr"
    prbs.EntityData.ParentYangName = "controller"
    prbs.EntityData.SegmentPath = "prbs"
    prbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prbs.EntityData.Children = types.NewOrderedMap()
    prbs.EntityData.Leafs = types.NewOrderedMap()
    prbs.EntityData.Leafs.Append("otu-prbs-test", types.YLeaf{"OtuPrbsTest", prbs.OtuPrbsTest})
    prbs.EntityData.Leafs.Append("otu-prbs-mode", types.YLeaf{"OtuPrbsMode", prbs.OtuPrbsMode})
    prbs.EntityData.Leafs.Append("otu-prbs-pattern", types.YLeaf{"OtuPrbsPattern", prbs.OtuPrbsPattern})
    prbs.EntityData.Leafs.Append("otu-prbs-status", types.YLeaf{"OtuPrbsStatus", prbs.OtuPrbsStatus})

    prbs.EntityData.YListKeys = []string {}

    return &(prbs.EntityData)
}

// Otu_Controllers_Controller_Info
// OTU port operational data
type Otu_Controllers_Controller_Info struct {
    EntityData types.CommonEntityData
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

    // FEC. The type is OtuG709fecMode.
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

func (info *Otu_Controllers_Controller_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "controller"
    info.EntityData.SegmentPath = "info"
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("local", types.YChild{"Local", &info.Local})
    info.EntityData.Children.Append("remote", types.YChild{"Remote", &info.Remote})
    info.EntityData.Children.Append("tti-mode", types.YChild{"TtiMode", &info.TtiMode})
    info.EntityData.Children.Append("network-srlg", types.YChild{"NetworkSrlg", &info.NetworkSrlg})
    info.EntityData.Children.Append("otu-alarm-info", types.YChild{"OtuAlarmInfo", &info.OtuAlarmInfo})
    info.EntityData.Children.Append("proactive", types.YChild{"Proactive", &info.Proactive})
    info.EntityData.Children.Append("otu-fec-satistics", types.YChild{"OtuFecSatistics", &info.OtuFecSatistics})
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("state", types.YLeaf{"State", info.State})
    info.EntityData.Leafs.Append("name", types.YLeaf{"Name", info.Name})
    info.EntityData.Leafs.Append("sf", types.YLeaf{"Sf", info.Sf})
    info.EntityData.Leafs.Append("sd", types.YLeaf{"Sd", info.Sd})
    info.EntityData.Leafs.Append("loopback-mode", types.YLeaf{"LoopbackMode", info.LoopbackMode})
    info.EntityData.Leafs.Append("fec-mode", types.YLeaf{"FecMode", info.FecMode})
    info.EntityData.Leafs.Append("derivedstate-mode", types.YLeaf{"DerivedstateMode", info.DerivedstateMode})
    info.EntityData.Leafs.Append("inherit-sec-state", types.YLeaf{"InheritSecState", info.InheritSecState})
    info.EntityData.Leafs.Append("config-sec-state", types.YLeaf{"ConfigSecState", info.ConfigSecState})
    info.EntityData.Leafs.Append("gcc-mode", types.YLeaf{"GccMode", info.GccMode})
    info.EntityData.Leafs.Append("q", types.YLeaf{"Q", info.Q})
    info.EntityData.Leafs.Append("q-margin", types.YLeaf{"QMargin", info.QMargin})
    info.EntityData.Leafs.Append("performance-monitoring", types.YLeaf{"PerformanceMonitoring", info.PerformanceMonitoring})
    info.EntityData.Leafs.Append("ec", types.YLeaf{"Ec", info.Ec})
    info.EntityData.Leafs.Append("uc", types.YLeaf{"Uc", info.Uc})
    info.EntityData.Leafs.Append("pre-fec-val", types.YLeaf{"PreFecVal", info.PreFecVal})
    info.EntityData.Leafs.Append("pre-fec-mantissa", types.YLeaf{"PreFecMantissa", info.PreFecMantissa})
    info.EntityData.Leafs.Append("ec-value", types.YLeaf{"EcValue", info.EcValue})
    info.EntityData.Leafs.Append("uc-value", types.YLeaf{"UcValue", info.UcValue})
    info.EntityData.Leafs.Append("pre-fec-ber-value", types.YLeaf{"PreFecBerValue", info.PreFecBerValue})
    info.EntityData.Leafs.Append("pre-fec-ber-mantissa", types.YLeaf{"PreFecBerMantissa", info.PreFecBerMantissa})
    info.EntityData.Leafs.Append("nv-optical-support", types.YLeaf{"NvOpticalSupport", info.NvOpticalSupport})
    info.EntityData.Leafs.Append("gmpls-tti-mode", types.YLeaf{"GmplsTtiMode", info.GmplsTtiMode})
    info.EntityData.Leafs.Append("gmpls-tvm-id", types.YLeaf{"GmplsTvmId", info.GmplsTvmId})
    info.EntityData.Leafs.Append("auto-tti-flag", types.YLeaf{"AutoTtiFlag", info.AutoTtiFlag})
    info.EntityData.Leafs.Append("description", types.YLeaf{"Description", info.Description})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Otu_Controllers_Controller_Info_Local
// TTI
type Otu_Controllers_Controller_Info_Local struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (local *Otu_Controllers_Controller_Info_Local) GetEntityData() *types.CommonEntityData {
    local.EntityData.YFilter = local.YFilter
    local.EntityData.YangName = "local"
    local.EntityData.BundleName = "cisco_ios_xr"
    local.EntityData.ParentYangName = "info"
    local.EntityData.SegmentPath = "local"
    local.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    local.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    local.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    local.EntityData.Children = types.NewOrderedMap()
    local.EntityData.Leafs = types.NewOrderedMap()
    local.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", local.RouterId})
    local.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", local.IfIndex})

    local.EntityData.YListKeys = []string {}

    return &(local.EntityData)
}

// Otu_Controllers_Controller_Info_Remote
// Remote
type Otu_Controllers_Controller_Info_Remote struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (remote *Otu_Controllers_Controller_Info_Remote) GetEntityData() *types.CommonEntityData {
    remote.EntityData.YFilter = remote.YFilter
    remote.EntityData.YangName = "remote"
    remote.EntityData.BundleName = "cisco_ios_xr"
    remote.EntityData.ParentYangName = "info"
    remote.EntityData.SegmentPath = "remote"
    remote.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remote.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remote.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remote.EntityData.Children = types.NewOrderedMap()
    remote.EntityData.Leafs = types.NewOrderedMap()
    remote.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", remote.RouterId})
    remote.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", remote.IfIndex})

    remote.EntityData.YListKeys = []string {}

    return &(remote.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode
// OTU TTI
type Otu_Controllers_Controller_Info_TtiMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // G709TTI sent. The type is OtuTtiEt.
    G709ttiSentMode interface{}

    // G709TTI Expected. The type is OtuTtiEt.
    G709ttiExpMode interface{}

    // G709TTI Recieved. The type is OtuTtiEt.
    G709ttiRecMode interface{}

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

func (ttiMode *Otu_Controllers_Controller_Info_TtiMode) GetEntityData() *types.CommonEntityData {
    ttiMode.EntityData.YFilter = ttiMode.YFilter
    ttiMode.EntityData.YangName = "tti-mode"
    ttiMode.EntityData.BundleName = "cisco_ios_xr"
    ttiMode.EntityData.ParentYangName = "info"
    ttiMode.EntityData.SegmentPath = "tti-mode"
    ttiMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ttiMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ttiMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ttiMode.EntityData.Children = types.NewOrderedMap()
    ttiMode.EntityData.Children.Append("tx", types.YChild{"Tx", &ttiMode.Tx})
    ttiMode.EntityData.Children.Append("exp", types.YChild{"Exp", &ttiMode.Exp})
    ttiMode.EntityData.Children.Append("rec", types.YChild{"Rec", &ttiMode.Rec})
    ttiMode.EntityData.Leafs = types.NewOrderedMap()
    ttiMode.EntityData.Leafs.Append("g709tti-sent-mode", types.YLeaf{"G709ttiSentMode", ttiMode.G709ttiSentMode})
    ttiMode.EntityData.Leafs.Append("g709tti-exp-mode", types.YLeaf{"G709ttiExpMode", ttiMode.G709ttiExpMode})
    ttiMode.EntityData.Leafs.Append("g709tti-rec-mode", types.YLeaf{"G709ttiRecMode", ttiMode.G709ttiRecMode})
    ttiMode.EntityData.Leafs.Append("remote-interface", types.YLeaf{"RemoteInterface", ttiMode.RemoteInterface})
    ttiMode.EntityData.Leafs.Append("remote-host-name", types.YLeaf{"RemoteHostName", ttiMode.RemoteHostName})
    ttiMode.EntityData.Leafs.Append("remote-ip-addr", types.YLeaf{"RemoteIpAddr", ttiMode.RemoteIpAddr})

    ttiMode.EntityData.YListKeys = []string {}

    return &(ttiMode.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Tx
// String Sent
type Otu_Controllers_Controller_Info_TtiMode_Tx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // full tti ascii String . The type is string.
    FullTtiAsciiString interface{}

    // tx String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Tx_Sapi.
    Sapi []*Otu_Controllers_Controller_Info_TtiMode_Tx_Sapi

    // exp String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Tx_Dapi.
    Dapi []*Otu_Controllers_Controller_Info_TtiMode_Tx_Dapi

    // rec String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Tx_OperatorSpecific.
    OperatorSpecific []*Otu_Controllers_Controller_Info_TtiMode_Tx_OperatorSpecific
}

func (tx *Otu_Controllers_Controller_Info_TtiMode_Tx) GetEntityData() *types.CommonEntityData {
    tx.EntityData.YFilter = tx.YFilter
    tx.EntityData.YangName = "tx"
    tx.EntityData.BundleName = "cisco_ios_xr"
    tx.EntityData.ParentYangName = "tti-mode"
    tx.EntityData.SegmentPath = "tx"
    tx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tx.EntityData.Children = types.NewOrderedMap()
    tx.EntityData.Children.Append("sapi", types.YChild{"Sapi", nil})
    for i := range tx.Sapi {
        tx.EntityData.Children.Append(types.GetSegmentPath(tx.Sapi[i]), types.YChild{"Sapi", tx.Sapi[i]})
    }
    tx.EntityData.Children.Append("dapi", types.YChild{"Dapi", nil})
    for i := range tx.Dapi {
        tx.EntityData.Children.Append(types.GetSegmentPath(tx.Dapi[i]), types.YChild{"Dapi", tx.Dapi[i]})
    }
    tx.EntityData.Children.Append("operator-specific", types.YChild{"OperatorSpecific", nil})
    for i := range tx.OperatorSpecific {
        tx.EntityData.Children.Append(types.GetSegmentPath(tx.OperatorSpecific[i]), types.YChild{"OperatorSpecific", tx.OperatorSpecific[i]})
    }
    tx.EntityData.Leafs = types.NewOrderedMap()
    tx.EntityData.Leafs.Append("full-tti-ascii-string", types.YLeaf{"FullTtiAsciiString", tx.FullTtiAsciiString})

    tx.EntityData.YListKeys = []string {}

    return &(tx.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Tx_Sapi
// tx String 
type Otu_Controllers_Controller_Info_TtiMode_Tx_Sapi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (sapi *Otu_Controllers_Controller_Info_TtiMode_Tx_Sapi) GetEntityData() *types.CommonEntityData {
    sapi.EntityData.YFilter = sapi.YFilter
    sapi.EntityData.YangName = "sapi"
    sapi.EntityData.BundleName = "cisco_ios_xr"
    sapi.EntityData.ParentYangName = "tx"
    sapi.EntityData.SegmentPath = "sapi"
    sapi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sapi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sapi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sapi.EntityData.Children = types.NewOrderedMap()
    sapi.EntityData.Leafs = types.NewOrderedMap()
    sapi.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", sapi.Entry})

    sapi.EntityData.YListKeys = []string {}

    return &(sapi.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Tx_Dapi
// exp String 
type Otu_Controllers_Controller_Info_TtiMode_Tx_Dapi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (dapi *Otu_Controllers_Controller_Info_TtiMode_Tx_Dapi) GetEntityData() *types.CommonEntityData {
    dapi.EntityData.YFilter = dapi.YFilter
    dapi.EntityData.YangName = "dapi"
    dapi.EntityData.BundleName = "cisco_ios_xr"
    dapi.EntityData.ParentYangName = "tx"
    dapi.EntityData.SegmentPath = "dapi"
    dapi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dapi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dapi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dapi.EntityData.Children = types.NewOrderedMap()
    dapi.EntityData.Leafs = types.NewOrderedMap()
    dapi.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dapi.Entry})

    dapi.EntityData.YListKeys = []string {}

    return &(dapi.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Tx_OperatorSpecific
// rec String 
type Otu_Controllers_Controller_Info_TtiMode_Tx_OperatorSpecific struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (operatorSpecific *Otu_Controllers_Controller_Info_TtiMode_Tx_OperatorSpecific) GetEntityData() *types.CommonEntityData {
    operatorSpecific.EntityData.YFilter = operatorSpecific.YFilter
    operatorSpecific.EntityData.YangName = "operator-specific"
    operatorSpecific.EntityData.BundleName = "cisco_ios_xr"
    operatorSpecific.EntityData.ParentYangName = "tx"
    operatorSpecific.EntityData.SegmentPath = "operator-specific"
    operatorSpecific.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operatorSpecific.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operatorSpecific.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operatorSpecific.EntityData.Children = types.NewOrderedMap()
    operatorSpecific.EntityData.Leafs = types.NewOrderedMap()
    operatorSpecific.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", operatorSpecific.Entry})

    operatorSpecific.EntityData.YListKeys = []string {}

    return &(operatorSpecific.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Exp
// String Expected
type Otu_Controllers_Controller_Info_TtiMode_Exp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // full tti ascii String . The type is string.
    FullTtiAsciiString interface{}

    // tx String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Exp_Sapi.
    Sapi []*Otu_Controllers_Controller_Info_TtiMode_Exp_Sapi

    // exp String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Exp_Dapi.
    Dapi []*Otu_Controllers_Controller_Info_TtiMode_Exp_Dapi

    // rec String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Exp_OperatorSpecific.
    OperatorSpecific []*Otu_Controllers_Controller_Info_TtiMode_Exp_OperatorSpecific
}

func (exp *Otu_Controllers_Controller_Info_TtiMode_Exp) GetEntityData() *types.CommonEntityData {
    exp.EntityData.YFilter = exp.YFilter
    exp.EntityData.YangName = "exp"
    exp.EntityData.BundleName = "cisco_ios_xr"
    exp.EntityData.ParentYangName = "tti-mode"
    exp.EntityData.SegmentPath = "exp"
    exp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exp.EntityData.Children = types.NewOrderedMap()
    exp.EntityData.Children.Append("sapi", types.YChild{"Sapi", nil})
    for i := range exp.Sapi {
        exp.EntityData.Children.Append(types.GetSegmentPath(exp.Sapi[i]), types.YChild{"Sapi", exp.Sapi[i]})
    }
    exp.EntityData.Children.Append("dapi", types.YChild{"Dapi", nil})
    for i := range exp.Dapi {
        exp.EntityData.Children.Append(types.GetSegmentPath(exp.Dapi[i]), types.YChild{"Dapi", exp.Dapi[i]})
    }
    exp.EntityData.Children.Append("operator-specific", types.YChild{"OperatorSpecific", nil})
    for i := range exp.OperatorSpecific {
        exp.EntityData.Children.Append(types.GetSegmentPath(exp.OperatorSpecific[i]), types.YChild{"OperatorSpecific", exp.OperatorSpecific[i]})
    }
    exp.EntityData.Leafs = types.NewOrderedMap()
    exp.EntityData.Leafs.Append("full-tti-ascii-string", types.YLeaf{"FullTtiAsciiString", exp.FullTtiAsciiString})

    exp.EntityData.YListKeys = []string {}

    return &(exp.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Exp_Sapi
// tx String 
type Otu_Controllers_Controller_Info_TtiMode_Exp_Sapi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (sapi *Otu_Controllers_Controller_Info_TtiMode_Exp_Sapi) GetEntityData() *types.CommonEntityData {
    sapi.EntityData.YFilter = sapi.YFilter
    sapi.EntityData.YangName = "sapi"
    sapi.EntityData.BundleName = "cisco_ios_xr"
    sapi.EntityData.ParentYangName = "exp"
    sapi.EntityData.SegmentPath = "sapi"
    sapi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sapi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sapi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sapi.EntityData.Children = types.NewOrderedMap()
    sapi.EntityData.Leafs = types.NewOrderedMap()
    sapi.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", sapi.Entry})

    sapi.EntityData.YListKeys = []string {}

    return &(sapi.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Exp_Dapi
// exp String 
type Otu_Controllers_Controller_Info_TtiMode_Exp_Dapi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (dapi *Otu_Controllers_Controller_Info_TtiMode_Exp_Dapi) GetEntityData() *types.CommonEntityData {
    dapi.EntityData.YFilter = dapi.YFilter
    dapi.EntityData.YangName = "dapi"
    dapi.EntityData.BundleName = "cisco_ios_xr"
    dapi.EntityData.ParentYangName = "exp"
    dapi.EntityData.SegmentPath = "dapi"
    dapi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dapi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dapi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dapi.EntityData.Children = types.NewOrderedMap()
    dapi.EntityData.Leafs = types.NewOrderedMap()
    dapi.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dapi.Entry})

    dapi.EntityData.YListKeys = []string {}

    return &(dapi.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Exp_OperatorSpecific
// rec String 
type Otu_Controllers_Controller_Info_TtiMode_Exp_OperatorSpecific struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (operatorSpecific *Otu_Controllers_Controller_Info_TtiMode_Exp_OperatorSpecific) GetEntityData() *types.CommonEntityData {
    operatorSpecific.EntityData.YFilter = operatorSpecific.YFilter
    operatorSpecific.EntityData.YangName = "operator-specific"
    operatorSpecific.EntityData.BundleName = "cisco_ios_xr"
    operatorSpecific.EntityData.ParentYangName = "exp"
    operatorSpecific.EntityData.SegmentPath = "operator-specific"
    operatorSpecific.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operatorSpecific.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operatorSpecific.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operatorSpecific.EntityData.Children = types.NewOrderedMap()
    operatorSpecific.EntityData.Leafs = types.NewOrderedMap()
    operatorSpecific.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", operatorSpecific.Entry})

    operatorSpecific.EntityData.YListKeys = []string {}

    return &(operatorSpecific.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Rec
// String Received
type Otu_Controllers_Controller_Info_TtiMode_Rec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // full tti ascii String . The type is string.
    FullTtiAsciiString interface{}

    // tx String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Rec_Sapi.
    Sapi []*Otu_Controllers_Controller_Info_TtiMode_Rec_Sapi

    // exp String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Rec_Dapi.
    Dapi []*Otu_Controllers_Controller_Info_TtiMode_Rec_Dapi

    // rec String . The type is slice of
    // Otu_Controllers_Controller_Info_TtiMode_Rec_OperatorSpecific.
    OperatorSpecific []*Otu_Controllers_Controller_Info_TtiMode_Rec_OperatorSpecific
}

func (rec *Otu_Controllers_Controller_Info_TtiMode_Rec) GetEntityData() *types.CommonEntityData {
    rec.EntityData.YFilter = rec.YFilter
    rec.EntityData.YangName = "rec"
    rec.EntityData.BundleName = "cisco_ios_xr"
    rec.EntityData.ParentYangName = "tti-mode"
    rec.EntityData.SegmentPath = "rec"
    rec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rec.EntityData.Children = types.NewOrderedMap()
    rec.EntityData.Children.Append("sapi", types.YChild{"Sapi", nil})
    for i := range rec.Sapi {
        rec.EntityData.Children.Append(types.GetSegmentPath(rec.Sapi[i]), types.YChild{"Sapi", rec.Sapi[i]})
    }
    rec.EntityData.Children.Append("dapi", types.YChild{"Dapi", nil})
    for i := range rec.Dapi {
        rec.EntityData.Children.Append(types.GetSegmentPath(rec.Dapi[i]), types.YChild{"Dapi", rec.Dapi[i]})
    }
    rec.EntityData.Children.Append("operator-specific", types.YChild{"OperatorSpecific", nil})
    for i := range rec.OperatorSpecific {
        rec.EntityData.Children.Append(types.GetSegmentPath(rec.OperatorSpecific[i]), types.YChild{"OperatorSpecific", rec.OperatorSpecific[i]})
    }
    rec.EntityData.Leafs = types.NewOrderedMap()
    rec.EntityData.Leafs.Append("full-tti-ascii-string", types.YLeaf{"FullTtiAsciiString", rec.FullTtiAsciiString})

    rec.EntityData.YListKeys = []string {}

    return &(rec.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Rec_Sapi
// tx String 
type Otu_Controllers_Controller_Info_TtiMode_Rec_Sapi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (sapi *Otu_Controllers_Controller_Info_TtiMode_Rec_Sapi) GetEntityData() *types.CommonEntityData {
    sapi.EntityData.YFilter = sapi.YFilter
    sapi.EntityData.YangName = "sapi"
    sapi.EntityData.BundleName = "cisco_ios_xr"
    sapi.EntityData.ParentYangName = "rec"
    sapi.EntityData.SegmentPath = "sapi"
    sapi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sapi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sapi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sapi.EntityData.Children = types.NewOrderedMap()
    sapi.EntityData.Leafs = types.NewOrderedMap()
    sapi.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", sapi.Entry})

    sapi.EntityData.YListKeys = []string {}

    return &(sapi.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Rec_Dapi
// exp String 
type Otu_Controllers_Controller_Info_TtiMode_Rec_Dapi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (dapi *Otu_Controllers_Controller_Info_TtiMode_Rec_Dapi) GetEntityData() *types.CommonEntityData {
    dapi.EntityData.YFilter = dapi.YFilter
    dapi.EntityData.YangName = "dapi"
    dapi.EntityData.BundleName = "cisco_ios_xr"
    dapi.EntityData.ParentYangName = "rec"
    dapi.EntityData.SegmentPath = "dapi"
    dapi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dapi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dapi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dapi.EntityData.Children = types.NewOrderedMap()
    dapi.EntityData.Leafs = types.NewOrderedMap()
    dapi.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dapi.Entry})

    dapi.EntityData.YListKeys = []string {}

    return &(dapi.EntityData)
}

// Otu_Controllers_Controller_Info_TtiMode_Rec_OperatorSpecific
// rec String 
type Otu_Controllers_Controller_Info_TtiMode_Rec_OperatorSpecific struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (operatorSpecific *Otu_Controllers_Controller_Info_TtiMode_Rec_OperatorSpecific) GetEntityData() *types.CommonEntityData {
    operatorSpecific.EntityData.YFilter = operatorSpecific.YFilter
    operatorSpecific.EntityData.YangName = "operator-specific"
    operatorSpecific.EntityData.BundleName = "cisco_ios_xr"
    operatorSpecific.EntityData.ParentYangName = "rec"
    operatorSpecific.EntityData.SegmentPath = "operator-specific"
    operatorSpecific.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operatorSpecific.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operatorSpecific.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operatorSpecific.EntityData.Children = types.NewOrderedMap()
    operatorSpecific.EntityData.Leafs = types.NewOrderedMap()
    operatorSpecific.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", operatorSpecific.Entry})

    operatorSpecific.EntityData.YListKeys = []string {}

    return &(operatorSpecific.EntityData)
}

// Otu_Controllers_Controller_Info_NetworkSrlg
// Network Shared Risk Link Group information
type Otu_Controllers_Controller_Info_NetworkSrlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of Network Shared Risk Link Group information. The type is slice of
    // Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo.
    SrlgInfo []*Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo
}

func (networkSrlg *Otu_Controllers_Controller_Info_NetworkSrlg) GetEntityData() *types.CommonEntityData {
    networkSrlg.EntityData.YFilter = networkSrlg.YFilter
    networkSrlg.EntityData.YangName = "network-srlg"
    networkSrlg.EntityData.BundleName = "cisco_ios_xr"
    networkSrlg.EntityData.ParentYangName = "info"
    networkSrlg.EntityData.SegmentPath = "network-srlg"
    networkSrlg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkSrlg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkSrlg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkSrlg.EntityData.Children = types.NewOrderedMap()
    networkSrlg.EntityData.Children.Append("srlg-info", types.YChild{"SrlgInfo", nil})
    for i := range networkSrlg.SrlgInfo {
        networkSrlg.EntityData.Children.Append(types.GetSegmentPath(networkSrlg.SrlgInfo[i]), types.YChild{"SrlgInfo", networkSrlg.SrlgInfo[i]})
    }
    networkSrlg.EntityData.Leafs = types.NewOrderedMap()

    networkSrlg.EntityData.YListKeys = []string {}

    return &(networkSrlg.EntityData)
}

// Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo
// Array of Network Shared Risk Link Group
// information
type Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array to maintain set id number. The type is interface{} with range:
    // 0..4294967295.
    SetId interface{}

    // Shared Risk Link Group information expressed in  integer format. The type
    // is slice of Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo_Srlg.
    Srlg []*Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo_Srlg
}

func (srlgInfo *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo) GetEntityData() *types.CommonEntityData {
    srlgInfo.EntityData.YFilter = srlgInfo.YFilter
    srlgInfo.EntityData.YangName = "srlg-info"
    srlgInfo.EntityData.BundleName = "cisco_ios_xr"
    srlgInfo.EntityData.ParentYangName = "network-srlg"
    srlgInfo.EntityData.SegmentPath = "srlg-info"
    srlgInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgInfo.EntityData.Children = types.NewOrderedMap()
    srlgInfo.EntityData.Children.Append("srlg", types.YChild{"Srlg", nil})
    for i := range srlgInfo.Srlg {
        srlgInfo.EntityData.Children.Append(types.GetSegmentPath(srlgInfo.Srlg[i]), types.YChild{"Srlg", srlgInfo.Srlg[i]})
    }
    srlgInfo.EntityData.Leafs = types.NewOrderedMap()
    srlgInfo.EntityData.Leafs.Append("set-id", types.YLeaf{"SetId", srlgInfo.SetId})

    srlgInfo.EntityData.YListKeys = []string {}

    return &(srlgInfo.EntityData)
}

// Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo_Srlg
// Shared Risk Link Group information expressed in
// 
// integer format
type Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo_Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (srlg *Otu_Controllers_Controller_Info_NetworkSrlg_SrlgInfo_Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "cisco_ios_xr"
    srlg.EntityData.ParentYangName = "srlg-info"
    srlg.EntityData.SegmentPath = "srlg"
    srlg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlg.EntityData.Children = types.NewOrderedMap()
    srlg.EntityData.Leafs = types.NewOrderedMap()
    srlg.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", srlg.Entry})

    srlg.EntityData.YListKeys = []string {}

    return &(srlg.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo
// OTU layer alarm Information
type Otu_Controllers_Controller_Info_OtuAlarmInfo struct {
    EntityData types.CommonEntityData
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

func (otuAlarmInfo *Otu_Controllers_Controller_Info_OtuAlarmInfo) GetEntityData() *types.CommonEntityData {
    otuAlarmInfo.EntityData.YFilter = otuAlarmInfo.YFilter
    otuAlarmInfo.EntityData.YangName = "otu-alarm-info"
    otuAlarmInfo.EntityData.BundleName = "cisco_ios_xr"
    otuAlarmInfo.EntityData.ParentYangName = "info"
    otuAlarmInfo.EntityData.SegmentPath = "otu-alarm-info"
    otuAlarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otuAlarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otuAlarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otuAlarmInfo.EntityData.Children = types.NewOrderedMap()
    otuAlarmInfo.EntityData.Children.Append("los", types.YChild{"Los", &otuAlarmInfo.Los})
    otuAlarmInfo.EntityData.Children.Append("lof", types.YChild{"Lof", &otuAlarmInfo.Lof})
    otuAlarmInfo.EntityData.Children.Append("lom", types.YChild{"Lom", &otuAlarmInfo.Lom})
    otuAlarmInfo.EntityData.Children.Append("oof", types.YChild{"Oof", &otuAlarmInfo.Oof})
    otuAlarmInfo.EntityData.Children.Append("oom", types.YChild{"Oom", &otuAlarmInfo.Oom})
    otuAlarmInfo.EntityData.Children.Append("ais", types.YChild{"Ais", &otuAlarmInfo.Ais})
    otuAlarmInfo.EntityData.Children.Append("iae", types.YChild{"Iae", &otuAlarmInfo.Iae})
    otuAlarmInfo.EntityData.Children.Append("biae", types.YChild{"Biae", &otuAlarmInfo.Biae})
    otuAlarmInfo.EntityData.Children.Append("bdi", types.YChild{"Bdi", &otuAlarmInfo.Bdi})
    otuAlarmInfo.EntityData.Children.Append("tim", types.YChild{"Tim", &otuAlarmInfo.Tim})
    otuAlarmInfo.EntityData.Children.Append("eoc", types.YChild{"Eoc", &otuAlarmInfo.Eoc})
    otuAlarmInfo.EntityData.Children.Append("fec-mismatch", types.YChild{"FecMismatch", &otuAlarmInfo.FecMismatch})
    otuAlarmInfo.EntityData.Children.Append("sf-ber", types.YChild{"SfBer", &otuAlarmInfo.SfBer})
    otuAlarmInfo.EntityData.Children.Append("sd-ber", types.YChild{"SdBer", &otuAlarmInfo.SdBer})
    otuAlarmInfo.EntityData.Children.Append("ec", types.YChild{"Ec", &otuAlarmInfo.Ec})
    otuAlarmInfo.EntityData.Children.Append("uc", types.YChild{"Uc", &otuAlarmInfo.Uc})
    otuAlarmInfo.EntityData.Children.Append("fecunc", types.YChild{"Fecunc", &otuAlarmInfo.Fecunc})
    otuAlarmInfo.EntityData.Leafs = types.NewOrderedMap()

    otuAlarmInfo.EntityData.YListKeys = []string {}

    return &(otuAlarmInfo.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Los
// Loss of Signal
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Los struct {
    EntityData types.CommonEntityData
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

func (los *Otu_Controllers_Controller_Info_OtuAlarmInfo_Los) GetEntityData() *types.CommonEntityData {
    los.EntityData.YFilter = los.YFilter
    los.EntityData.YangName = "los"
    los.EntityData.BundleName = "cisco_ios_xr"
    los.EntityData.ParentYangName = "otu-alarm-info"
    los.EntityData.SegmentPath = "los"
    los.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    los.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    los.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    los.EntityData.Children = types.NewOrderedMap()
    los.EntityData.Leafs = types.NewOrderedMap()
    los.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", los.ReportingEnabled})
    los.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", los.IsDetected})
    los.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", los.IsAsserted})
    los.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", los.Counter})

    los.EntityData.YListKeys = []string {}

    return &(los.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof
// Loss of Frame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof struct {
    EntityData types.CommonEntityData
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

func (lof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lof) GetEntityData() *types.CommonEntityData {
    lof.EntityData.YFilter = lof.YFilter
    lof.EntityData.YangName = "lof"
    lof.EntityData.BundleName = "cisco_ios_xr"
    lof.EntityData.ParentYangName = "otu-alarm-info"
    lof.EntityData.SegmentPath = "lof"
    lof.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lof.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lof.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lof.EntityData.Children = types.NewOrderedMap()
    lof.EntityData.Leafs = types.NewOrderedMap()
    lof.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", lof.ReportingEnabled})
    lof.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lof.IsDetected})
    lof.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", lof.IsAsserted})
    lof.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lof.Counter})

    lof.EntityData.YListKeys = []string {}

    return &(lof.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom
// Loss of MultiFrame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom struct {
    EntityData types.CommonEntityData
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

func (lom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Lom) GetEntityData() *types.CommonEntityData {
    lom.EntityData.YFilter = lom.YFilter
    lom.EntityData.YangName = "lom"
    lom.EntityData.BundleName = "cisco_ios_xr"
    lom.EntityData.ParentYangName = "otu-alarm-info"
    lom.EntityData.SegmentPath = "lom"
    lom.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lom.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lom.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lom.EntityData.Children = types.NewOrderedMap()
    lom.EntityData.Leafs = types.NewOrderedMap()
    lom.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", lom.ReportingEnabled})
    lom.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", lom.IsDetected})
    lom.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", lom.IsAsserted})
    lom.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", lom.Counter})

    lom.EntityData.YListKeys = []string {}

    return &(lom.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof
// Out of Frame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof struct {
    EntityData types.CommonEntityData
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

func (oof *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oof) GetEntityData() *types.CommonEntityData {
    oof.EntityData.YFilter = oof.YFilter
    oof.EntityData.YangName = "oof"
    oof.EntityData.BundleName = "cisco_ios_xr"
    oof.EntityData.ParentYangName = "otu-alarm-info"
    oof.EntityData.SegmentPath = "oof"
    oof.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oof.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oof.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oof.EntityData.Children = types.NewOrderedMap()
    oof.EntityData.Leafs = types.NewOrderedMap()
    oof.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", oof.ReportingEnabled})
    oof.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", oof.IsDetected})
    oof.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", oof.IsAsserted})
    oof.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", oof.Counter})

    oof.EntityData.YListKeys = []string {}

    return &(oof.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom
// Out of MultiFrame
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom struct {
    EntityData types.CommonEntityData
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

func (oom *Otu_Controllers_Controller_Info_OtuAlarmInfo_Oom) GetEntityData() *types.CommonEntityData {
    oom.EntityData.YFilter = oom.YFilter
    oom.EntityData.YangName = "oom"
    oom.EntityData.BundleName = "cisco_ios_xr"
    oom.EntityData.ParentYangName = "otu-alarm-info"
    oom.EntityData.SegmentPath = "oom"
    oom.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oom.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oom.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oom.EntityData.Children = types.NewOrderedMap()
    oom.EntityData.Leafs = types.NewOrderedMap()
    oom.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", oom.ReportingEnabled})
    oom.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", oom.IsDetected})
    oom.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", oom.IsAsserted})
    oom.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", oom.Counter})

    oom.EntityData.YListKeys = []string {}

    return &(oom.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais
// Alarm Indication Signal
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais struct {
    EntityData types.CommonEntityData
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

func (ais *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ais) GetEntityData() *types.CommonEntityData {
    ais.EntityData.YFilter = ais.YFilter
    ais.EntityData.YangName = "ais"
    ais.EntityData.BundleName = "cisco_ios_xr"
    ais.EntityData.ParentYangName = "otu-alarm-info"
    ais.EntityData.SegmentPath = "ais"
    ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ais.EntityData.Children = types.NewOrderedMap()
    ais.EntityData.Leafs = types.NewOrderedMap()
    ais.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", ais.ReportingEnabled})
    ais.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ais.IsDetected})
    ais.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", ais.IsAsserted})
    ais.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ais.Counter})

    ais.EntityData.YListKeys = []string {}

    return &(ais.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae
// Incoming Alignment Error
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae struct {
    EntityData types.CommonEntityData
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

func (iae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Iae) GetEntityData() *types.CommonEntityData {
    iae.EntityData.YFilter = iae.YFilter
    iae.EntityData.YangName = "iae"
    iae.EntityData.BundleName = "cisco_ios_xr"
    iae.EntityData.ParentYangName = "otu-alarm-info"
    iae.EntityData.SegmentPath = "iae"
    iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iae.EntityData.Children = types.NewOrderedMap()
    iae.EntityData.Leafs = types.NewOrderedMap()
    iae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", iae.ReportingEnabled})
    iae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", iae.IsDetected})
    iae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", iae.IsAsserted})
    iae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", iae.Counter})

    iae.EntityData.YListKeys = []string {}

    return &(iae.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae
// Backward Incoming Alignment Error
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae struct {
    EntityData types.CommonEntityData
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

func (biae *Otu_Controllers_Controller_Info_OtuAlarmInfo_Biae) GetEntityData() *types.CommonEntityData {
    biae.EntityData.YFilter = biae.YFilter
    biae.EntityData.YangName = "biae"
    biae.EntityData.BundleName = "cisco_ios_xr"
    biae.EntityData.ParentYangName = "otu-alarm-info"
    biae.EntityData.SegmentPath = "biae"
    biae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    biae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    biae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    biae.EntityData.Children = types.NewOrderedMap()
    biae.EntityData.Leafs = types.NewOrderedMap()
    biae.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", biae.ReportingEnabled})
    biae.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", biae.IsDetected})
    biae.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", biae.IsAsserted})
    biae.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", biae.Counter})

    biae.EntityData.YListKeys = []string {}

    return &(biae.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi
// Backward Defect Indication
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi struct {
    EntityData types.CommonEntityData
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

func (bdi *Otu_Controllers_Controller_Info_OtuAlarmInfo_Bdi) GetEntityData() *types.CommonEntityData {
    bdi.EntityData.YFilter = bdi.YFilter
    bdi.EntityData.YangName = "bdi"
    bdi.EntityData.BundleName = "cisco_ios_xr"
    bdi.EntityData.ParentYangName = "otu-alarm-info"
    bdi.EntityData.SegmentPath = "bdi"
    bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdi.EntityData.Children = types.NewOrderedMap()
    bdi.EntityData.Leafs = types.NewOrderedMap()
    bdi.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", bdi.ReportingEnabled})
    bdi.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", bdi.IsDetected})
    bdi.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", bdi.IsAsserted})
    bdi.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", bdi.Counter})

    bdi.EntityData.YListKeys = []string {}

    return &(bdi.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim
// Trace Identifier Mismatch
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim struct {
    EntityData types.CommonEntityData
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

func (tim *Otu_Controllers_Controller_Info_OtuAlarmInfo_Tim) GetEntityData() *types.CommonEntityData {
    tim.EntityData.YFilter = tim.YFilter
    tim.EntityData.YangName = "tim"
    tim.EntityData.BundleName = "cisco_ios_xr"
    tim.EntityData.ParentYangName = "otu-alarm-info"
    tim.EntityData.SegmentPath = "tim"
    tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tim.EntityData.Children = types.NewOrderedMap()
    tim.EntityData.Leafs = types.NewOrderedMap()
    tim.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", tim.ReportingEnabled})
    tim.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", tim.IsDetected})
    tim.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", tim.IsAsserted})
    tim.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", tim.Counter})

    tim.EntityData.YListKeys = []string {}

    return &(tim.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc
// GCC End of Channel
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc struct {
    EntityData types.CommonEntityData
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

func (eoc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Eoc) GetEntityData() *types.CommonEntityData {
    eoc.EntityData.YFilter = eoc.YFilter
    eoc.EntityData.YangName = "eoc"
    eoc.EntityData.BundleName = "cisco_ios_xr"
    eoc.EntityData.ParentYangName = "otu-alarm-info"
    eoc.EntityData.SegmentPath = "eoc"
    eoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eoc.EntityData.Children = types.NewOrderedMap()
    eoc.EntityData.Leafs = types.NewOrderedMap()
    eoc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", eoc.ReportingEnabled})
    eoc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", eoc.IsDetected})
    eoc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", eoc.IsAsserted})
    eoc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", eoc.Counter})

    eoc.EntityData.YListKeys = []string {}

    return &(eoc.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch
// FEC mismatch alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch struct {
    EntityData types.CommonEntityData
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

func (fecMismatch *Otu_Controllers_Controller_Info_OtuAlarmInfo_FecMismatch) GetEntityData() *types.CommonEntityData {
    fecMismatch.EntityData.YFilter = fecMismatch.YFilter
    fecMismatch.EntityData.YangName = "fec-mismatch"
    fecMismatch.EntityData.BundleName = "cisco_ios_xr"
    fecMismatch.EntityData.ParentYangName = "otu-alarm-info"
    fecMismatch.EntityData.SegmentPath = "fec-mismatch"
    fecMismatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fecMismatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fecMismatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fecMismatch.EntityData.Children = types.NewOrderedMap()
    fecMismatch.EntityData.Leafs = types.NewOrderedMap()
    fecMismatch.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", fecMismatch.ReportingEnabled})
    fecMismatch.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", fecMismatch.IsDetected})
    fecMismatch.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", fecMismatch.IsAsserted})
    fecMismatch.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", fecMismatch.Counter})

    fecMismatch.EntityData.YListKeys = []string {}

    return &(fecMismatch.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer
// SF BER alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer struct {
    EntityData types.CommonEntityData
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

func (sfBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SfBer) GetEntityData() *types.CommonEntityData {
    sfBer.EntityData.YFilter = sfBer.YFilter
    sfBer.EntityData.YangName = "sf-ber"
    sfBer.EntityData.BundleName = "cisco_ios_xr"
    sfBer.EntityData.ParentYangName = "otu-alarm-info"
    sfBer.EntityData.SegmentPath = "sf-ber"
    sfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfBer.EntityData.Children = types.NewOrderedMap()
    sfBer.EntityData.Leafs = types.NewOrderedMap()
    sfBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", sfBer.ReportingEnabled})
    sfBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", sfBer.IsDetected})
    sfBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", sfBer.IsAsserted})
    sfBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", sfBer.Counter})

    sfBer.EntityData.YListKeys = []string {}

    return &(sfBer.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer
// SD BER alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer struct {
    EntityData types.CommonEntityData
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

func (sdBer *Otu_Controllers_Controller_Info_OtuAlarmInfo_SdBer) GetEntityData() *types.CommonEntityData {
    sdBer.EntityData.YFilter = sdBer.YFilter
    sdBer.EntityData.YangName = "sd-ber"
    sdBer.EntityData.BundleName = "cisco_ios_xr"
    sdBer.EntityData.ParentYangName = "otu-alarm-info"
    sdBer.EntityData.SegmentPath = "sd-ber"
    sdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdBer.EntityData.Children = types.NewOrderedMap()
    sdBer.EntityData.Leafs = types.NewOrderedMap()
    sdBer.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", sdBer.ReportingEnabled})
    sdBer.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", sdBer.IsDetected})
    sdBer.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", sdBer.IsAsserted})
    sdBer.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", sdBer.Counter})

    sdBer.EntityData.YListKeys = []string {}

    return &(sdBer.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec
// EC alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec struct {
    EntityData types.CommonEntityData
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

func (ec *Otu_Controllers_Controller_Info_OtuAlarmInfo_Ec) GetEntityData() *types.CommonEntityData {
    ec.EntityData.YFilter = ec.YFilter
    ec.EntityData.YangName = "ec"
    ec.EntityData.BundleName = "cisco_ios_xr"
    ec.EntityData.ParentYangName = "otu-alarm-info"
    ec.EntityData.SegmentPath = "ec"
    ec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ec.EntityData.Children = types.NewOrderedMap()
    ec.EntityData.Leafs = types.NewOrderedMap()
    ec.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", ec.ReportingEnabled})
    ec.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", ec.IsDetected})
    ec.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", ec.IsAsserted})
    ec.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", ec.Counter})

    ec.EntityData.YListKeys = []string {}

    return &(ec.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc
// UC alarm
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc struct {
    EntityData types.CommonEntityData
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

func (uc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Uc) GetEntityData() *types.CommonEntityData {
    uc.EntityData.YFilter = uc.YFilter
    uc.EntityData.YangName = "uc"
    uc.EntityData.BundleName = "cisco_ios_xr"
    uc.EntityData.ParentYangName = "otu-alarm-info"
    uc.EntityData.SegmentPath = "uc"
    uc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    uc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    uc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    uc.EntityData.Children = types.NewOrderedMap()
    uc.EntityData.Leafs = types.NewOrderedMap()
    uc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", uc.ReportingEnabled})
    uc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", uc.IsDetected})
    uc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", uc.IsAsserted})
    uc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", uc.Counter})

    uc.EntityData.YListKeys = []string {}

    return &(uc.EntityData)
}

// Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc
// FEC UnCorrected Word
type Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc struct {
    EntityData types.CommonEntityData
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

func (fecunc *Otu_Controllers_Controller_Info_OtuAlarmInfo_Fecunc) GetEntityData() *types.CommonEntityData {
    fecunc.EntityData.YFilter = fecunc.YFilter
    fecunc.EntityData.YangName = "fecunc"
    fecunc.EntityData.BundleName = "cisco_ios_xr"
    fecunc.EntityData.ParentYangName = "otu-alarm-info"
    fecunc.EntityData.SegmentPath = "fecunc"
    fecunc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fecunc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fecunc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fecunc.EntityData.Children = types.NewOrderedMap()
    fecunc.EntityData.Leafs = types.NewOrderedMap()
    fecunc.EntityData.Leafs.Append("reporting-enabled", types.YLeaf{"ReportingEnabled", fecunc.ReportingEnabled})
    fecunc.EntityData.Leafs.Append("is-detected", types.YLeaf{"IsDetected", fecunc.IsDetected})
    fecunc.EntityData.Leafs.Append("is-asserted", types.YLeaf{"IsAsserted", fecunc.IsAsserted})
    fecunc.EntityData.Leafs.Append("counter", types.YLeaf{"Counter", fecunc.Counter})

    fecunc.EntityData.YListKeys = []string {}

    return &(fecunc.EntityData)
}

// Otu_Controllers_Controller_Info_Proactive
// Proactive Protection
type Otu_Controllers_Controller_Info_Proactive struct {
    EntityData types.CommonEntityData
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

func (proactive *Otu_Controllers_Controller_Info_Proactive) GetEntityData() *types.CommonEntityData {
    proactive.EntityData.YFilter = proactive.YFilter
    proactive.EntityData.YangName = "proactive"
    proactive.EntityData.BundleName = "cisco_ios_xr"
    proactive.EntityData.ParentYangName = "info"
    proactive.EntityData.SegmentPath = "proactive"
    proactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proactive.EntityData.Children = types.NewOrderedMap()
    proactive.EntityData.Leafs = types.NewOrderedMap()
    proactive.EntityData.Leafs.Append("proactive-status", types.YLeaf{"ProactiveStatus", proactive.ProactiveStatus})
    proactive.EntityData.Leafs.Append("inherit-sec-state", types.YLeaf{"InheritSecState", proactive.InheritSecState})
    proactive.EntityData.Leafs.Append("config-sec-state", types.YLeaf{"ConfigSecState", proactive.ConfigSecState})
    proactive.EntityData.Leafs.Append("proactive-fsm-state", types.YLeaf{"ProactiveFsmState", proactive.ProactiveFsmState})
    proactive.EntityData.Leafs.Append("proactive-fsm-if-state", types.YLeaf{"ProactiveFsmIfState", proactive.ProactiveFsmIfState})
    proactive.EntityData.Leafs.Append("trig-thresh-coeff", types.YLeaf{"TrigThreshCoeff", proactive.TrigThreshCoeff})
    proactive.EntityData.Leafs.Append("trig-thresh-power", types.YLeaf{"TrigThreshPower", proactive.TrigThreshPower})
    proactive.EntityData.Leafs.Append("rvrt-thresh-coeff", types.YLeaf{"RvrtThreshCoeff", proactive.RvrtThreshCoeff})
    proactive.EntityData.Leafs.Append("rvrt-thresh-power", types.YLeaf{"RvrtThreshPower", proactive.RvrtThreshPower})
    proactive.EntityData.Leafs.Append("trigger-window", types.YLeaf{"TriggerWindow", proactive.TriggerWindow})
    proactive.EntityData.Leafs.Append("revert-window", types.YLeaf{"RevertWindow", proactive.RevertWindow})

    proactive.EntityData.YListKeys = []string {}

    return &(proactive.EntityData)
}

// Otu_Controllers_Controller_Info_OtuFecSatistics
// OTU FEC Statistics
type Otu_Controllers_Controller_Info_OtuFecSatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bit Error Rate After Forward Error Correction. The type is string.
    PostFecBer interface{}

    // Bit Error Rate Before Forward Error Correction. The type is string.
    PreFecBer interface{}
}

func (otuFecSatistics *Otu_Controllers_Controller_Info_OtuFecSatistics) GetEntityData() *types.CommonEntityData {
    otuFecSatistics.EntityData.YFilter = otuFecSatistics.YFilter
    otuFecSatistics.EntityData.YangName = "otu-fec-satistics"
    otuFecSatistics.EntityData.BundleName = "cisco_ios_xr"
    otuFecSatistics.EntityData.ParentYangName = "info"
    otuFecSatistics.EntityData.SegmentPath = "otu-fec-satistics"
    otuFecSatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otuFecSatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otuFecSatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otuFecSatistics.EntityData.Children = types.NewOrderedMap()
    otuFecSatistics.EntityData.Leafs = types.NewOrderedMap()
    otuFecSatistics.EntityData.Leafs.Append("post-fec-ber", types.YLeaf{"PostFecBer", otuFecSatistics.PostFecBer})
    otuFecSatistics.EntityData.Leafs.Append("pre-fec-ber", types.YLeaf{"PreFecBer", otuFecSatistics.PreFecBer})

    otuFecSatistics.EntityData.YListKeys = []string {}

    return &(otuFecSatistics.EntityData)
}

