// This module contains a collection of YANG definitions
// for Cisco IOS-XR dwdm-ui package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dwdm: DWDM operational data
//   vtxp: vtxp
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dwdm_ui_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dwdm_ui_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dwdm-ui-oper dwdm}", reflect.TypeOf(Dwdm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dwdm-ui-oper:dwdm", reflect.TypeOf(Dwdm{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dwdm-ui-oper vtxp}", reflect.TypeOf(Vtxp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dwdm-ui-oper:vtxp", reflect.TypeOf(Vtxp{}))
}

// G709ppintfState represents G709ppintf state
type G709ppintfState string

const (
    // Interface is Up
    G709ppintfState_pp_intf_up G709ppintfState = "pp-intf-up"

    // Interface is Going Down
    G709ppintfState_pp_intf_failing G709ppintfState = "pp-intf-failing"

    // Interface Down
    G709ppintfState_pp_intf_down G709ppintfState = "pp-intf-down"
)

// G709ppfsmMode represents G709ppfsm mode
type G709ppfsmMode string

const (
    // OFF
    G709ppfsmMode_pp_disable G709ppfsmMode = "pp-disable"

    // ON (Default Mode)
    G709ppfsmMode_pp_default_mode G709ppfsmMode = "pp-default-mode"

    // ON (Graceful Mode)
    G709ppfsmMode_pp_graceful_mode G709ppfsmMode = "pp-graceful-mode"
)

// DwdmControllerState represents Dwdm controller state
type DwdmControllerState string

const (
    // Up
    DwdmControllerState_dwdm_ui_state_up DwdmControllerState = "dwdm-ui-state-up"

    // Down
    DwdmControllerState_dwdm_ui_state_down DwdmControllerState = "dwdm-ui-state-down"

    // Administratively Down
    DwdmControllerState_dwdm_ui_state_admin_down DwdmControllerState = "dwdm-ui-state-admin-down"
)

// DwdmtasState represents Dwdmtas state
type DwdmtasState string

const (
    // Out of Service
    DwdmtasState_tas_oos DwdmtasState = "tas-oos"

    // In Service
    DwdmtasState_tas_is DwdmtasState = "tas-is"

    // Out of Service Maintenance
    DwdmtasState_tas_oos_mt DwdmtasState = "tas-oos-mt"

    // In Service Config allowed
    DwdmtasState_tas_is_cfg DwdmtasState = "tas-is-cfg"
)

// G709prbsMode represents G709prbs mode
type G709prbsMode string

const (
    // mode source
    G709prbsMode_mode_source G709prbsMode = "mode-source"

    // mode sink
    G709prbsMode_mode_sink G709prbsMode = "mode-sink"

    // mode source sink
    G709prbsMode_mode_source_sink G709prbsMode = "mode-source-sink"

    // mode invalid
    G709prbsMode_mode_invalid G709prbsMode = "mode-invalid"
)

// G709ppfsmState represents G709ppfsm state
type G709ppfsmState string

const (
    // In Active
    G709ppfsmState_in_active G709ppfsmState = "in-active"

    // Disabled
    G709ppfsmState_disabled G709ppfsmState = "disabled"

    // Normal
    G709ppfsmState_normal_state G709ppfsmState = "normal-state"

    // Local Failing
    G709ppfsmState_local_failing G709ppfsmState = "local-failing"

    // Remote Failing
    G709ppfsmState_remote_failing G709ppfsmState = "remote-failing"

    // Maintance Failing
    G709ppfsmState_main_t_failing G709ppfsmState = "main-t-failing"

    // Regenerator Failing
    G709ppfsmState_regen_failing G709ppfsmState = "regen-failing"

    // Local Failed
    G709ppfsmState_local_failed G709ppfsmState = "local-failed"

    // Remote Failed
    G709ppfsmState_remote_failed G709ppfsmState = "remote-failed"

    // Maintance Failed
    G709ppfsmState_main_t_failed G709ppfsmState = "main-t-failed"

    // Regenerator Failed
    G709ppfsmState_regen_failed G709ppfsmState = "regen-failed"
)

// G709prbsInterval represents PRBS test interval information
type G709prbsInterval string

const (
    // Current interval
    G709prbsInterval_current_interval G709prbsInterval = "current-interval"

    // Previous interval
    G709prbsInterval_previous_interval G709prbsInterval = "previous-interval"

    // Previous interval 2
    G709prbsInterval_previous_interval2 G709prbsInterval = "previous-interval2"

    // Previous interval 3
    G709prbsInterval_previous_interval3 G709prbsInterval = "previous-interval3"

    // Previous interval 4
    G709prbsInterval_previous_interval4 G709prbsInterval = "previous-interval4"

    // Previous interval 5
    G709prbsInterval_previous_interval5 G709prbsInterval = "previous-interval5"

    // Previous interval 6
    G709prbsInterval_previous_interval6 G709prbsInterval = "previous-interval6"

    // Previous interval 7
    G709prbsInterval_previous_interval7 G709prbsInterval = "previous-interval7"

    // Previous interval 8
    G709prbsInterval_previous_interval8 G709prbsInterval = "previous-interval8"

    // Previous interval 9
    G709prbsInterval_previous_interval9 G709prbsInterval = "previous-interval9"

    // Previous interval 10
    G709prbsInterval_previous_interval10 G709prbsInterval = "previous-interval10"

    // Previous interval 11
    G709prbsInterval_previous_interval11 G709prbsInterval = "previous-interval11"

    // Previous interval 12
    G709prbsInterval_previous_interval12 G709prbsInterval = "previous-interval12"

    // Previous interval 13
    G709prbsInterval_previous_interval13 G709prbsInterval = "previous-interval13"

    // Previous interval 14
    G709prbsInterval_previous_interval14 G709prbsInterval = "previous-interval14"

    // Previous interval 15
    G709prbsInterval_previous_interval15 G709prbsInterval = "previous-interval15"

    // Previous interval 16
    G709prbsInterval_previous_interval16 G709prbsInterval = "previous-interval16"

    // Previous interval 17
    G709prbsInterval_previous_interval17 G709prbsInterval = "previous-interval17"

    // Previous interval 18
    G709prbsInterval_previous_interval18 G709prbsInterval = "previous-interval18"

    // Previous interval 19
    G709prbsInterval_previous_interval19 G709prbsInterval = "previous-interval19"

    // Previous interval 20
    G709prbsInterval_previous_interval20 G709prbsInterval = "previous-interval20"

    // Previous interval 21
    G709prbsInterval_previous_interval21 G709prbsInterval = "previous-interval21"

    // Previous interval 22
    G709prbsInterval_previous_interval22 G709prbsInterval = "previous-interval22"

    // Previous interval 23
    G709prbsInterval_previous_interval23 G709prbsInterval = "previous-interval23"

    // Previous interval 24
    G709prbsInterval_previous_interval24 G709prbsInterval = "previous-interval24"

    // Previous interval 25
    G709prbsInterval_previous_interval25 G709prbsInterval = "previous-interval25"

    // Previous interval 26
    G709prbsInterval_previous_interval26 G709prbsInterval = "previous-interval26"

    // Previous interval 27
    G709prbsInterval_previous_interval27 G709prbsInterval = "previous-interval27"

    // Previous interval 28
    G709prbsInterval_previous_interval28 G709prbsInterval = "previous-interval28"

    // Previous interval 29
    G709prbsInterval_previous_interval29 G709prbsInterval = "previous-interval29"

    // Previous interval 30
    G709prbsInterval_previous_interval30 G709prbsInterval = "previous-interval30"

    // Previous interval 31
    G709prbsInterval_previous_interval31 G709prbsInterval = "previous-interval31"

    // Previous interval 32
    G709prbsInterval_previous_interval32 G709prbsInterval = "previous-interval32"
)

// DwdmWaveChannelOwner represents Dwdm wave channel owner
type DwdmWaveChannelOwner string

const (
    // Hardware Default
    DwdmWaveChannelOwner_default_ DwdmWaveChannelOwner = "default"

    // Configuration
    DwdmWaveChannelOwner_configuration DwdmWaveChannelOwner = "configuration"

    // GMPLS Signaled
    DwdmWaveChannelOwner_gmpls DwdmWaveChannelOwner = "gmpls"
)

// G709efecMode represents G709efec mode
type G709efecMode string

const (
    // 
    G709efecMode_g975_none G709efecMode = "g975-none"

    // G975.1 I.4
    G709efecMode_g975_1_i4 G709efecMode = "g975-1-i4"

    // G975.1 I.7
    G709efecMode_g975_1_i7 G709efecMode = "g975-1-i7"
)

// G709prbsPattern represents G709prbs pattern
type G709prbsPattern string

const (
    // pattern none
    G709prbsPattern_pattern_none G709prbsPattern = "pattern-none"

    // pattern null
    G709prbsPattern_pattern_null G709prbsPattern = "pattern-null"

    // pattern pn11
    G709prbsPattern_pattern_pn11 G709prbsPattern = "pattern-pn11"

    // pattern pn23
    G709prbsPattern_pattern_pn23 G709prbsPattern = "pattern-pn23"

    // pattern pn31
    G709prbsPattern_pattern_pn31 G709prbsPattern = "pattern-pn31"
)

// G709apsByte represents G709aps byte
type G709apsByte string

const (
    // No Protection
    G709apsByte_pp_no_protect G709apsByte = "pp-no-protect"

    // No Request
    G709apsByte_pp_no_request G709apsByte = "pp-no-request"

    // Regenerator Degrade
    G709apsByte_pp_regen_degrade G709apsByte = "pp-regen-degrade"

    // Signal Degrade
    G709apsByte_pp_sig_degrade G709apsByte = "pp-sig-degrade"

    // Maintenance Request
    G709apsByte_pp_remote_main G709apsByte = "pp-remote-main"

    // Unknown
    G709apsByte_pp_aps_unknown G709apsByte = "pp-aps-unknown"
)

// Dwdm
// DWDM operational data
type Dwdm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All DWDM Port operational data.
    Ports Dwdm_Ports
}

func (dwdm *Dwdm) GetFilter() yfilter.YFilter { return dwdm.YFilter }

func (dwdm *Dwdm) SetFilter(yf yfilter.YFilter) { dwdm.YFilter = yf }

func (dwdm *Dwdm) GetGoName(yname string) string {
    if yname == "ports" { return "Ports" }
    return ""
}

func (dwdm *Dwdm) GetSegmentPath() string {
    return "Cisco-IOS-XR-dwdm-ui-oper:dwdm"
}

func (dwdm *Dwdm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ports" {
        return &dwdm.Ports
    }
    return nil
}

func (dwdm *Dwdm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ports"] = &dwdm.Ports
    return children
}

func (dwdm *Dwdm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dwdm *Dwdm) GetBundleName() string { return "cisco_ios_xr" }

func (dwdm *Dwdm) GetYangName() string { return "dwdm" }

func (dwdm *Dwdm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dwdm *Dwdm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dwdm *Dwdm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dwdm *Dwdm) SetParent(parent types.Entity) { dwdm.parent = parent }

func (dwdm *Dwdm) GetParent() types.Entity { return dwdm.parent }

func (dwdm *Dwdm) GetParentYangName() string { return "Cisco-IOS-XR-dwdm-ui-oper" }

// Dwdm_Ports
// All DWDM Port operational data
type Dwdm_Ports struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DWDM Port operational data. The type is slice of Dwdm_Ports_Port.
    Port []Dwdm_Ports_Port
}

func (ports *Dwdm_Ports) GetFilter() yfilter.YFilter { return ports.YFilter }

func (ports *Dwdm_Ports) SetFilter(yf yfilter.YFilter) { ports.YFilter = yf }

func (ports *Dwdm_Ports) GetGoName(yname string) string {
    if yname == "port" { return "Port" }
    return ""
}

func (ports *Dwdm_Ports) GetSegmentPath() string {
    return "ports"
}

func (ports *Dwdm_Ports) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port" {
        for _, c := range ports.Port {
            if ports.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dwdm_Ports_Port{}
        ports.Port = append(ports.Port, child)
        return &ports.Port[len(ports.Port)-1]
    }
    return nil
}

func (ports *Dwdm_Ports) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ports.Port {
        children[ports.Port[i].GetSegmentPath()] = &ports.Port[i]
    }
    return children
}

func (ports *Dwdm_Ports) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ports *Dwdm_Ports) GetBundleName() string { return "cisco_ios_xr" }

func (ports *Dwdm_Ports) GetYangName() string { return "ports" }

func (ports *Dwdm_Ports) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ports *Dwdm_Ports) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ports *Dwdm_Ports) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ports *Dwdm_Ports) SetParent(parent types.Entity) { ports.parent = parent }

func (ports *Dwdm_Ports) GetParent() types.Entity { return ports.parent }

func (ports *Dwdm_Ports) GetParentYangName() string { return "dwdm" }

// Dwdm_Ports_Port
// DWDM Port operational data
type Dwdm_Ports_Port struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // DWDM Port PRBS related data.
    Prbs Dwdm_Ports_Port_Prbs

    // DWDM Port optics operational data.
    Optics Dwdm_Ports_Port_Optics

    // DWDM port operational data.
    Info Dwdm_Ports_Port_Info
}

func (port *Dwdm_Ports_Port) GetFilter() yfilter.YFilter { return port.YFilter }

func (port *Dwdm_Ports_Port) SetFilter(yf yfilter.YFilter) { port.YFilter = yf }

func (port *Dwdm_Ports_Port) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "prbs" { return "Prbs" }
    if yname == "optics" { return "Optics" }
    if yname == "info" { return "Info" }
    return ""
}

func (port *Dwdm_Ports_Port) GetSegmentPath() string {
    return "port" + "[name='" + fmt.Sprintf("%v", port.Name) + "']"
}

func (port *Dwdm_Ports_Port) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prbs" {
        return &port.Prbs
    }
    if childYangName == "optics" {
        return &port.Optics
    }
    if childYangName == "info" {
        return &port.Info
    }
    return nil
}

func (port *Dwdm_Ports_Port) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prbs"] = &port.Prbs
    children["optics"] = &port.Optics
    children["info"] = &port.Info
    return children
}

func (port *Dwdm_Ports_Port) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = port.Name
    return leafs
}

func (port *Dwdm_Ports_Port) GetBundleName() string { return "cisco_ios_xr" }

func (port *Dwdm_Ports_Port) GetYangName() string { return "port" }

func (port *Dwdm_Ports_Port) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (port *Dwdm_Ports_Port) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (port *Dwdm_Ports_Port) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (port *Dwdm_Ports_Port) SetParent(parent types.Entity) { port.parent = parent }

func (port *Dwdm_Ports_Port) GetParent() types.Entity { return port.parent }

func (port *Dwdm_Ports_Port) GetParentYangName() string { return "ports" }

// Dwdm_Ports_Port_Prbs
// DWDM Port PRBS related data
type Dwdm_Ports_Port_Prbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port 24-hour PRBS statistics table.
    TwentyFourHoursBucket Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket

    // Port 15-minute PRBS statistics table.
    FifteenMinutesBucket Dwdm_Ports_Port_Prbs_FifteenMinutesBucket
}

func (prbs *Dwdm_Ports_Port_Prbs) GetFilter() yfilter.YFilter { return prbs.YFilter }

func (prbs *Dwdm_Ports_Port_Prbs) SetFilter(yf yfilter.YFilter) { prbs.YFilter = yf }

func (prbs *Dwdm_Ports_Port_Prbs) GetGoName(yname string) string {
    if yname == "twenty-four-hours-bucket" { return "TwentyFourHoursBucket" }
    if yname == "fifteen-minutes-bucket" { return "FifteenMinutesBucket" }
    return ""
}

func (prbs *Dwdm_Ports_Port_Prbs) GetSegmentPath() string {
    return "prbs"
}

func (prbs *Dwdm_Ports_Port_Prbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "twenty-four-hours-bucket" {
        return &prbs.TwentyFourHoursBucket
    }
    if childYangName == "fifteen-minutes-bucket" {
        return &prbs.FifteenMinutesBucket
    }
    return nil
}

func (prbs *Dwdm_Ports_Port_Prbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["twenty-four-hours-bucket"] = &prbs.TwentyFourHoursBucket
    children["fifteen-minutes-bucket"] = &prbs.FifteenMinutesBucket
    return children
}

func (prbs *Dwdm_Ports_Port_Prbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prbs *Dwdm_Ports_Port_Prbs) GetBundleName() string { return "cisco_ios_xr" }

func (prbs *Dwdm_Ports_Port_Prbs) GetYangName() string { return "prbs" }

func (prbs *Dwdm_Ports_Port_Prbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prbs *Dwdm_Ports_Port_Prbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prbs *Dwdm_Ports_Port_Prbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prbs *Dwdm_Ports_Port_Prbs) SetParent(parent types.Entity) { prbs.parent = parent }

func (prbs *Dwdm_Ports_Port_Prbs) GetParent() types.Entity { return prbs.parent }

func (prbs *Dwdm_Ports_Port_Prbs) GetParentYangName() string { return "port" }

// Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket
// Port 24-hour PRBS statistics table
type Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port 24-hour PRBS statistics data.
    TwentyFourHoursStatistics Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics
}

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetFilter() yfilter.YFilter { return twentyFourHoursBucket.YFilter }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) SetFilter(yf yfilter.YFilter) { twentyFourHoursBucket.YFilter = yf }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetGoName(yname string) string {
    if yname == "twenty-four-hours-statistics" { return "TwentyFourHoursStatistics" }
    return ""
}

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetSegmentPath() string {
    return "twenty-four-hours-bucket"
}

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "twenty-four-hours-statistics" {
        return &twentyFourHoursBucket.TwentyFourHoursStatistics
    }
    return nil
}

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["twenty-four-hours-statistics"] = &twentyFourHoursBucket.TwentyFourHoursStatistics
    return children
}

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetBundleName() string { return "cisco_ios_xr" }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetYangName() string { return "twenty-four-hours-bucket" }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) SetParent(parent types.Entity) { twentyFourHoursBucket.parent = parent }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetParent() types.Entity { return twentyFourHoursBucket.parent }

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetParentYangName() string { return "prbs" }

// Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics
// Port 24-hour PRBS statistics data
type Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 'True' if PRBS is enabled 'False' otherwise. The type is bool.
    IsPrbsEnabled interface{}

    // Configured mode of PRBS test. The type is G709prbsMode.
    PrbsConfigMode interface{}

    // History consists of 15-minute/24-hour intervals. The type is slice of
    // Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry.
    PrbsEntry []Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry
}

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetFilter() yfilter.YFilter { return twentyFourHoursStatistics.YFilter }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) SetFilter(yf yfilter.YFilter) { twentyFourHoursStatistics.YFilter = yf }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetGoName(yname string) string {
    if yname == "is-prbs-enabled" { return "IsPrbsEnabled" }
    if yname == "prbs-config-mode" { return "PrbsConfigMode" }
    if yname == "prbs-entry" { return "PrbsEntry" }
    return ""
}

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetSegmentPath() string {
    return "twenty-four-hours-statistics"
}

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prbs-entry" {
        for _, c := range twentyFourHoursStatistics.PrbsEntry {
            if twentyFourHoursStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry{}
        twentyFourHoursStatistics.PrbsEntry = append(twentyFourHoursStatistics.PrbsEntry, child)
        return &twentyFourHoursStatistics.PrbsEntry[len(twentyFourHoursStatistics.PrbsEntry)-1]
    }
    return nil
}

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range twentyFourHoursStatistics.PrbsEntry {
        children[twentyFourHoursStatistics.PrbsEntry[i].GetSegmentPath()] = &twentyFourHoursStatistics.PrbsEntry[i]
    }
    return children
}

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-prbs-enabled"] = twentyFourHoursStatistics.IsPrbsEnabled
    leafs["prbs-config-mode"] = twentyFourHoursStatistics.PrbsConfigMode
    return leafs
}

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetYangName() string { return "twenty-four-hours-statistics" }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) SetParent(parent types.Entity) { twentyFourHoursStatistics.parent = parent }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetParent() types.Entity { return twentyFourHoursStatistics.parent }

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetParentYangName() string { return "twenty-four-hours-bucket" }

// Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry
// History consists of 15-minute/24-hour intervals
type Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Index of bucket, current and previous. The type is G709prbsInterval.
    IntervalIndex interface{}

    // Configured pattern of PRBS test. The type is G709prbsPattern.
    ConfiguredPattern interface{}

    // Interval start timestamp. The type is string with length: 0..64.
    StartAt interface{}

    // Interval stop timestamp. The type is string with length: 0..64.
    StopAt interface{}

    // Received Pattern of PRBS Test. The type is G709prbsPattern.
    ReceivedPattern interface{}

    // Bit Error Count. The type is interface{} with range:
    // 0..18446744073709551615.
    BitErrorCount interface{}

    // Count of pattern found in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    FoundCount interface{}

    // Count of pattern lost in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    LostCount interface{}

    // Pattern first found at timestamp. The type is string with length: 0..64.
    FoundAt interface{}

    // Pattern first lost at timestamp. The type is string with length: 0..64.
    LostAt interface{}
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetFilter() yfilter.YFilter { return prbsEntry.YFilter }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) SetFilter(yf yfilter.YFilter) { prbsEntry.YFilter = yf }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetGoName(yname string) string {
    if yname == "interval-index" { return "IntervalIndex" }
    if yname == "configured-pattern" { return "ConfiguredPattern" }
    if yname == "start-at" { return "StartAt" }
    if yname == "stop-at" { return "StopAt" }
    if yname == "received-pattern" { return "ReceivedPattern" }
    if yname == "bit-error-count" { return "BitErrorCount" }
    if yname == "found-count" { return "FoundCount" }
    if yname == "lost-count" { return "LostCount" }
    if yname == "found-at" { return "FoundAt" }
    if yname == "lost-at" { return "LostAt" }
    return ""
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetSegmentPath() string {
    return "prbs-entry"
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval-index"] = prbsEntry.IntervalIndex
    leafs["configured-pattern"] = prbsEntry.ConfiguredPattern
    leafs["start-at"] = prbsEntry.StartAt
    leafs["stop-at"] = prbsEntry.StopAt
    leafs["received-pattern"] = prbsEntry.ReceivedPattern
    leafs["bit-error-count"] = prbsEntry.BitErrorCount
    leafs["found-count"] = prbsEntry.FoundCount
    leafs["lost-count"] = prbsEntry.LostCount
    leafs["found-at"] = prbsEntry.FoundAt
    leafs["lost-at"] = prbsEntry.LostAt
    return leafs
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetBundleName() string { return "cisco_ios_xr" }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetYangName() string { return "prbs-entry" }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) SetParent(parent types.Entity) { prbsEntry.parent = parent }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetParent() types.Entity { return prbsEntry.parent }

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetParentYangName() string { return "twenty-four-hours-statistics" }

// Dwdm_Ports_Port_Prbs_FifteenMinutesBucket
// Port 15-minute PRBS statistics table
type Dwdm_Ports_Port_Prbs_FifteenMinutesBucket struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port 15-minute PRBS statistics data.
    FifteenMinutesStatistics Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics
}

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetFilter() yfilter.YFilter { return fifteenMinutesBucket.YFilter }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) SetFilter(yf yfilter.YFilter) { fifteenMinutesBucket.YFilter = yf }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetGoName(yname string) string {
    if yname == "fifteen-minutes-statistics" { return "FifteenMinutesStatistics" }
    return ""
}

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetSegmentPath() string {
    return "fifteen-minutes-bucket"
}

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fifteen-minutes-statistics" {
        return &fifteenMinutesBucket.FifteenMinutesStatistics
    }
    return nil
}

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fifteen-minutes-statistics"] = &fifteenMinutesBucket.FifteenMinutesStatistics
    return children
}

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetBundleName() string { return "cisco_ios_xr" }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetYangName() string { return "fifteen-minutes-bucket" }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) SetParent(parent types.Entity) { fifteenMinutesBucket.parent = parent }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetParent() types.Entity { return fifteenMinutesBucket.parent }

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetParentYangName() string { return "prbs" }

// Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics
// Port 15-minute PRBS statistics data
type Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 'True' if PRBS is enabled 'False' otherwise. The type is bool.
    IsPrbsEnabled interface{}

    // Configured mode of PRBS test. The type is G709prbsMode.
    PrbsConfigMode interface{}

    // History consists of 15-minute/24-hour intervals. The type is slice of
    // Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry.
    PrbsEntry []Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry
}

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetFilter() yfilter.YFilter { return fifteenMinutesStatistics.YFilter }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) SetFilter(yf yfilter.YFilter) { fifteenMinutesStatistics.YFilter = yf }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetGoName(yname string) string {
    if yname == "is-prbs-enabled" { return "IsPrbsEnabled" }
    if yname == "prbs-config-mode" { return "PrbsConfigMode" }
    if yname == "prbs-entry" { return "PrbsEntry" }
    return ""
}

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetSegmentPath() string {
    return "fifteen-minutes-statistics"
}

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prbs-entry" {
        for _, c := range fifteenMinutesStatistics.PrbsEntry {
            if fifteenMinutesStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry{}
        fifteenMinutesStatistics.PrbsEntry = append(fifteenMinutesStatistics.PrbsEntry, child)
        return &fifteenMinutesStatistics.PrbsEntry[len(fifteenMinutesStatistics.PrbsEntry)-1]
    }
    return nil
}

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fifteenMinutesStatistics.PrbsEntry {
        children[fifteenMinutesStatistics.PrbsEntry[i].GetSegmentPath()] = &fifteenMinutesStatistics.PrbsEntry[i]
    }
    return children
}

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-prbs-enabled"] = fifteenMinutesStatistics.IsPrbsEnabled
    leafs["prbs-config-mode"] = fifteenMinutesStatistics.PrbsConfigMode
    return leafs
}

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetYangName() string { return "fifteen-minutes-statistics" }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) SetParent(parent types.Entity) { fifteenMinutesStatistics.parent = parent }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetParent() types.Entity { return fifteenMinutesStatistics.parent }

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetParentYangName() string { return "fifteen-minutes-bucket" }

// Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry
// History consists of 15-minute/24-hour intervals
type Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Index of bucket, current and previous. The type is G709prbsInterval.
    IntervalIndex interface{}

    // Configured pattern of PRBS test. The type is G709prbsPattern.
    ConfiguredPattern interface{}

    // Interval start timestamp. The type is string with length: 0..64.
    StartAt interface{}

    // Interval stop timestamp. The type is string with length: 0..64.
    StopAt interface{}

    // Received Pattern of PRBS Test. The type is G709prbsPattern.
    ReceivedPattern interface{}

    // Bit Error Count. The type is interface{} with range:
    // 0..18446744073709551615.
    BitErrorCount interface{}

    // Count of pattern found in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    FoundCount interface{}

    // Count of pattern lost in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    LostCount interface{}

    // Pattern first found at timestamp. The type is string with length: 0..64.
    FoundAt interface{}

    // Pattern first lost at timestamp. The type is string with length: 0..64.
    LostAt interface{}
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetFilter() yfilter.YFilter { return prbsEntry.YFilter }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) SetFilter(yf yfilter.YFilter) { prbsEntry.YFilter = yf }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetGoName(yname string) string {
    if yname == "interval-index" { return "IntervalIndex" }
    if yname == "configured-pattern" { return "ConfiguredPattern" }
    if yname == "start-at" { return "StartAt" }
    if yname == "stop-at" { return "StopAt" }
    if yname == "received-pattern" { return "ReceivedPattern" }
    if yname == "bit-error-count" { return "BitErrorCount" }
    if yname == "found-count" { return "FoundCount" }
    if yname == "lost-count" { return "LostCount" }
    if yname == "found-at" { return "FoundAt" }
    if yname == "lost-at" { return "LostAt" }
    return ""
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetSegmentPath() string {
    return "prbs-entry"
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval-index"] = prbsEntry.IntervalIndex
    leafs["configured-pattern"] = prbsEntry.ConfiguredPattern
    leafs["start-at"] = prbsEntry.StartAt
    leafs["stop-at"] = prbsEntry.StopAt
    leafs["received-pattern"] = prbsEntry.ReceivedPattern
    leafs["bit-error-count"] = prbsEntry.BitErrorCount
    leafs["found-count"] = prbsEntry.FoundCount
    leafs["lost-count"] = prbsEntry.LostCount
    leafs["found-at"] = prbsEntry.FoundAt
    leafs["lost-at"] = prbsEntry.LostAt
    return leafs
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetBundleName() string { return "cisco_ios_xr" }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetYangName() string { return "prbs-entry" }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) SetParent(parent types.Entity) { prbsEntry.parent = parent }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetParent() types.Entity { return prbsEntry.parent }

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetParentYangName() string { return "fifteen-minutes-statistics" }

// Dwdm_Ports_Port_Optics
// DWDM Port optics operational data
type Dwdm_Ports_Port_Optics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DWDM port wavelength information data.
    WaveInfo Dwdm_Ports_Port_Optics_WaveInfo
}

func (optics *Dwdm_Ports_Port_Optics) GetFilter() yfilter.YFilter { return optics.YFilter }

func (optics *Dwdm_Ports_Port_Optics) SetFilter(yf yfilter.YFilter) { optics.YFilter = yf }

func (optics *Dwdm_Ports_Port_Optics) GetGoName(yname string) string {
    if yname == "wave-info" { return "WaveInfo" }
    return ""
}

func (optics *Dwdm_Ports_Port_Optics) GetSegmentPath() string {
    return "optics"
}

func (optics *Dwdm_Ports_Port_Optics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wave-info" {
        return &optics.WaveInfo
    }
    return nil
}

func (optics *Dwdm_Ports_Port_Optics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wave-info"] = &optics.WaveInfo
    return children
}

func (optics *Dwdm_Ports_Port_Optics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (optics *Dwdm_Ports_Port_Optics) GetBundleName() string { return "cisco_ios_xr" }

func (optics *Dwdm_Ports_Port_Optics) GetYangName() string { return "optics" }

func (optics *Dwdm_Ports_Port_Optics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optics *Dwdm_Ports_Port_Optics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optics *Dwdm_Ports_Port_Optics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optics *Dwdm_Ports_Port_Optics) SetParent(parent types.Entity) { optics.parent = parent }

func (optics *Dwdm_Ports_Port_Optics) GetParent() types.Entity { return optics.parent }

func (optics *Dwdm_Ports_Port_Optics) GetParentYangName() string { return "port" }

// Dwdm_Ports_Port_Optics_WaveInfo
// DWDM port wavelength information data
type Dwdm_Ports_Port_Optics_WaveInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Wavelength band. The type is interface{} with range: 0..4294967295.
    WaveBand interface{}

    // Lowest ITU wavelength channel number supported. The type is interface{}
    // with range: 0..4294967295.
    WaveChannelMin interface{}

    // Highest ITU wavelength channel number supported. The type is interface{}
    // with range: 0..4294967295.
    WaveChannelMax interface{}
}

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetFilter() yfilter.YFilter { return waveInfo.YFilter }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) SetFilter(yf yfilter.YFilter) { waveInfo.YFilter = yf }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetGoName(yname string) string {
    if yname == "wave-band" { return "WaveBand" }
    if yname == "wave-channel-min" { return "WaveChannelMin" }
    if yname == "wave-channel-max" { return "WaveChannelMax" }
    return ""
}

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetSegmentPath() string {
    return "wave-info"
}

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["wave-band"] = waveInfo.WaveBand
    leafs["wave-channel-min"] = waveInfo.WaveChannelMin
    leafs["wave-channel-max"] = waveInfo.WaveChannelMax
    return leafs
}

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetBundleName() string { return "cisco_ios_xr" }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetYangName() string { return "wave-info" }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) SetParent(parent types.Entity) { waveInfo.parent = parent }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetParent() types.Entity { return waveInfo.parent }

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetParentYangName() string { return "optics" }

// Dwdm_Ports_Port_Info
// DWDM port operational data
type Dwdm_Ports_Port_Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DWDM controller state: Up, Down or Administratively Down. The type is
    // DwdmControllerState.
    ControllerState interface{}

    // DWDM controller TAS state: IS, OOS, OOS-MT or IS-CFG. The type is
    // DwdmtasState.
    TransportAdminState interface{}

    // DWDM port slice state Up/Down. The type is bool.
    SliceState interface{}

    // G709 operational information.
    G709Info Dwdm_Ports_Port_Info_G709Info

    // Optics operational information.
    OpticsInfo Dwdm_Ports_Port_Info_OpticsInfo

    // TDC operational information.
    TdcInfo Dwdm_Ports_Port_Info_TdcInfo

    // Network SRLG information.
    NetworkSrlgInfo Dwdm_Ports_Port_Info_NetworkSrlgInfo

    // Proactive protection information.
    Proactive Dwdm_Ports_Port_Info_Proactive

    // Signal log information.
    SignalLog Dwdm_Ports_Port_Info_SignalLog
}

func (info *Dwdm_Ports_Port_Info) GetFilter() yfilter.YFilter { return info.YFilter }

func (info *Dwdm_Ports_Port_Info) SetFilter(yf yfilter.YFilter) { info.YFilter = yf }

func (info *Dwdm_Ports_Port_Info) GetGoName(yname string) string {
    if yname == "controller-state" { return "ControllerState" }
    if yname == "transport-admin-state" { return "TransportAdminState" }
    if yname == "slice-state" { return "SliceState" }
    if yname == "g709-info" { return "G709Info" }
    if yname == "optics-info" { return "OpticsInfo" }
    if yname == "tdc-info" { return "TdcInfo" }
    if yname == "network-srlg-info" { return "NetworkSrlgInfo" }
    if yname == "proactive" { return "Proactive" }
    if yname == "signal-log" { return "SignalLog" }
    return ""
}

func (info *Dwdm_Ports_Port_Info) GetSegmentPath() string {
    return "info"
}

func (info *Dwdm_Ports_Port_Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "g709-info" {
        return &info.G709Info
    }
    if childYangName == "optics-info" {
        return &info.OpticsInfo
    }
    if childYangName == "tdc-info" {
        return &info.TdcInfo
    }
    if childYangName == "network-srlg-info" {
        return &info.NetworkSrlgInfo
    }
    if childYangName == "proactive" {
        return &info.Proactive
    }
    if childYangName == "signal-log" {
        return &info.SignalLog
    }
    return nil
}

func (info *Dwdm_Ports_Port_Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["g709-info"] = &info.G709Info
    children["optics-info"] = &info.OpticsInfo
    children["tdc-info"] = &info.TdcInfo
    children["network-srlg-info"] = &info.NetworkSrlgInfo
    children["proactive"] = &info.Proactive
    children["signal-log"] = &info.SignalLog
    return children
}

func (info *Dwdm_Ports_Port_Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller-state"] = info.ControllerState
    leafs["transport-admin-state"] = info.TransportAdminState
    leafs["slice-state"] = info.SliceState
    return leafs
}

func (info *Dwdm_Ports_Port_Info) GetBundleName() string { return "cisco_ios_xr" }

func (info *Dwdm_Ports_Port_Info) GetYangName() string { return "info" }

func (info *Dwdm_Ports_Port_Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (info *Dwdm_Ports_Port_Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (info *Dwdm_Ports_Port_Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (info *Dwdm_Ports_Port_Info) SetParent(parent types.Entity) { info.parent = parent }

func (info *Dwdm_Ports_Port_Info) GetParent() types.Entity { return info.parent }

func (info *Dwdm_Ports_Port_Info) GetParentYangName() string { return "port" }

// Dwdm_Ports_Port_Info_G709Info
// G709 operational information
type Dwdm_Ports_Port_Info_G709Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is G709 framing enabled. The type is bool.
    IsG709Enabled interface{}

    // Is Operating FEC Mode Default. The type is bool.
    IsFecModeDefault interface{}

    // FEC information. The type is interface{} with range:
    // -2147483648..2147483647.
    FecMode interface{}

    // Remote FEC information. The type is interface{} with range:
    // -2147483648..2147483647.
    RemoteFecMode interface{}

    // EFEC information. The type is G709efecMode.
    EfecMode interface{}

    // Loopback information. The type is interface{} with range:
    // -2147483648..2147483647.
    LoopbackMode interface{}

    // Corrected bit error counter . The type is interface{} with range:
    // 0..18446744073709551615.
    Ec interface{}

    // FEC Corrected bit error accumulated counter. The type is interface{} with
    // range: 0..18446744073709551615.
    EcAccum interface{}

    // FEC Uncorrected words counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc interface{}

    // pre fec ber calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    FecBer interface{}

    // pre fec ber calculated. The type is interface{} with range:
    // -2147483648..2147483647.
    FecBerMan interface{}

    // q value calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    Q interface{}

    // q margin calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    QMargin interface{}

    // FEC BER String . The type is string with length: 0..64.
    FeCstr interface{}

    // Q String . The type is string with length: 0..64.
    Qstr interface{}

    // QMargin String. The type is string with length: 0..64.
    QmarginStr interface{}

    // Network port ID. The type is string with length: 0..65.
    NetworkPortId interface{}

    // Network connection ID. The type is string with length: 0..65.
    NetworkConnId interface{}

    // 'true' if Prbs is enabled 'false' otherwise. The type is bool.
    IsPrbsEnabled interface{}

    // Configured mode of PRBS Test. The type is G709prbsMode.
    G709PrbsMode interface{}

    // Pattern of PRBS Test. The type is G709prbsPattern.
    G709PrbsPattern interface{}

    // Time stamp for prbs configuration. The type is interface{} with range:
    // 0..18446744073709551615.
    PrbsTimeStamp interface{}

    // FEC mismatch alarm.
    FecMismatch Dwdm_Ports_Port_Info_G709Info_FecMismatch

    // FEC Corrected bits TCA information.
    EcTca Dwdm_Ports_Port_Info_G709Info_EcTca

    // FEC uncorrected words TCA information.
    UcTca Dwdm_Ports_Port_Info_G709Info_UcTca

    // OTU layer information.
    OtuInfo Dwdm_Ports_Port_Info_G709Info_OtuInfo

    // ODU layer Information.
    OduInfo Dwdm_Ports_Port_Info_G709Info_OduInfo
}

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetFilter() yfilter.YFilter { return g709Info.YFilter }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) SetFilter(yf yfilter.YFilter) { g709Info.YFilter = yf }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetGoName(yname string) string {
    if yname == "is-g709-enabled" { return "IsG709Enabled" }
    if yname == "is-fec-mode-default" { return "IsFecModeDefault" }
    if yname == "fec-mode" { return "FecMode" }
    if yname == "remote-fec-mode" { return "RemoteFecMode" }
    if yname == "efec-mode" { return "EfecMode" }
    if yname == "loopback-mode" { return "LoopbackMode" }
    if yname == "ec" { return "Ec" }
    if yname == "ec-accum" { return "EcAccum" }
    if yname == "uc" { return "Uc" }
    if yname == "fec-ber" { return "FecBer" }
    if yname == "fec-ber-man" { return "FecBerMan" }
    if yname == "q" { return "Q" }
    if yname == "q-margin" { return "QMargin" }
    if yname == "fe-cstr" { return "FeCstr" }
    if yname == "qstr" { return "Qstr" }
    if yname == "qmargin-str" { return "QmarginStr" }
    if yname == "network-port-id" { return "NetworkPortId" }
    if yname == "network-conn-id" { return "NetworkConnId" }
    if yname == "is-prbs-enabled" { return "IsPrbsEnabled" }
    if yname == "g709-prbs-mode" { return "G709PrbsMode" }
    if yname == "g709-prbs-pattern" { return "G709PrbsPattern" }
    if yname == "prbs-time-stamp" { return "PrbsTimeStamp" }
    if yname == "fec-mismatch" { return "FecMismatch" }
    if yname == "ec-tca" { return "EcTca" }
    if yname == "uc-tca" { return "UcTca" }
    if yname == "otu-info" { return "OtuInfo" }
    if yname == "odu-info" { return "OduInfo" }
    return ""
}

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetSegmentPath() string {
    return "g709-info"
}

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fec-mismatch" {
        return &g709Info.FecMismatch
    }
    if childYangName == "ec-tca" {
        return &g709Info.EcTca
    }
    if childYangName == "uc-tca" {
        return &g709Info.UcTca
    }
    if childYangName == "otu-info" {
        return &g709Info.OtuInfo
    }
    if childYangName == "odu-info" {
        return &g709Info.OduInfo
    }
    return nil
}

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fec-mismatch"] = &g709Info.FecMismatch
    children["ec-tca"] = &g709Info.EcTca
    children["uc-tca"] = &g709Info.UcTca
    children["otu-info"] = &g709Info.OtuInfo
    children["odu-info"] = &g709Info.OduInfo
    return children
}

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-g709-enabled"] = g709Info.IsG709Enabled
    leafs["is-fec-mode-default"] = g709Info.IsFecModeDefault
    leafs["fec-mode"] = g709Info.FecMode
    leafs["remote-fec-mode"] = g709Info.RemoteFecMode
    leafs["efec-mode"] = g709Info.EfecMode
    leafs["loopback-mode"] = g709Info.LoopbackMode
    leafs["ec"] = g709Info.Ec
    leafs["ec-accum"] = g709Info.EcAccum
    leafs["uc"] = g709Info.Uc
    leafs["fec-ber"] = g709Info.FecBer
    leafs["fec-ber-man"] = g709Info.FecBerMan
    leafs["q"] = g709Info.Q
    leafs["q-margin"] = g709Info.QMargin
    leafs["fe-cstr"] = g709Info.FeCstr
    leafs["qstr"] = g709Info.Qstr
    leafs["qmargin-str"] = g709Info.QmarginStr
    leafs["network-port-id"] = g709Info.NetworkPortId
    leafs["network-conn-id"] = g709Info.NetworkConnId
    leafs["is-prbs-enabled"] = g709Info.IsPrbsEnabled
    leafs["g709-prbs-mode"] = g709Info.G709PrbsMode
    leafs["g709-prbs-pattern"] = g709Info.G709PrbsPattern
    leafs["prbs-time-stamp"] = g709Info.PrbsTimeStamp
    return leafs
}

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetBundleName() string { return "cisco_ios_xr" }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetYangName() string { return "g709-info" }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) SetParent(parent types.Entity) { g709Info.parent = parent }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetParent() types.Entity { return g709Info.parent }

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetParentYangName() string { return "info" }

// Dwdm_Ports_Port_Info_G709Info_FecMismatch
// FEC mismatch alarm
type Dwdm_Ports_Port_Info_G709Info_FecMismatch struct {
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

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetFilter() yfilter.YFilter { return fecMismatch.YFilter }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) SetFilter(yf yfilter.YFilter) { fecMismatch.YFilter = yf }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetSegmentPath() string {
    return "fec-mismatch"
}

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = fecMismatch.ReportingEnabled
    leafs["is-detected"] = fecMismatch.IsDetected
    leafs["is-asserted"] = fecMismatch.IsAsserted
    leafs["counter"] = fecMismatch.Counter
    return leafs
}

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetBundleName() string { return "cisco_ios_xr" }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetYangName() string { return "fec-mismatch" }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) SetParent(parent types.Entity) { fecMismatch.parent = parent }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetParent() types.Entity { return fecMismatch.parent }

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetParentYangName() string { return "g709-info" }

// Dwdm_Ports_Port_Info_G709Info_EcTca
// FEC Corrected bits TCA information
type Dwdm_Ports_Port_Info_G709Info_EcTca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetFilter() yfilter.YFilter { return ecTca.YFilter }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) SetFilter(yf yfilter.YFilter) { ecTca.YFilter = yf }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetSegmentPath() string {
    return "ec-tca"
}

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = ecTca.ReportingEnabled
    leafs["is-detected"] = ecTca.IsDetected
    leafs["is-asserted"] = ecTca.IsAsserted
    leafs["threshold"] = ecTca.Threshold
    leafs["counter"] = ecTca.Counter
    return leafs
}

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetBundleName() string { return "cisco_ios_xr" }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetYangName() string { return "ec-tca" }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) SetParent(parent types.Entity) { ecTca.parent = parent }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetParent() types.Entity { return ecTca.parent }

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetParentYangName() string { return "g709-info" }

// Dwdm_Ports_Port_Info_G709Info_UcTca
// FEC uncorrected words TCA information
type Dwdm_Ports_Port_Info_G709Info_UcTca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetFilter() yfilter.YFilter { return ucTca.YFilter }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) SetFilter(yf yfilter.YFilter) { ucTca.YFilter = yf }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetSegmentPath() string {
    return "uc-tca"
}

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = ucTca.ReportingEnabled
    leafs["is-detected"] = ucTca.IsDetected
    leafs["is-asserted"] = ucTca.IsAsserted
    leafs["threshold"] = ucTca.Threshold
    leafs["counter"] = ucTca.Counter
    return leafs
}

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetBundleName() string { return "cisco_ios_xr" }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetYangName() string { return "uc-tca" }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) SetParent(parent types.Entity) { ucTca.parent = parent }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetParent() types.Entity { return ucTca.parent }

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetParentYangName() string { return "g709-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo
// OTU layer information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Backward Error Indication counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bei interface{}

    // Bit Interleave Parity(BIP) counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bip interface{}

    // Loss of Signal information.
    Los Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los

    // Loss of Frame information.
    Lof Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof

    // Loss of MultiFrame information.
    Lom Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom

    // Out of Frame information.
    Oof Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof

    // Out of MultiFrame information.
    Oom Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom

    // Alarm Indication Signal information.
    Ais Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais

    // Incoming Alignment Error information.
    Iae Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae

    // Backward Defect Indication information.
    Bdi Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi

    // Trace Identifier Mismatch information.
    Tim Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim

    // GCC End of Channel information.
    Eoc Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc

    // Signal Fail  BER information.
    SfBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer

    // Signal Degrade BER information.
    SdBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer

    // Prefec Signal Fail BER information.
    PrefecSfBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer

    // Prefec Signal Degrade BER information.
    PrefecSdBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer

    // Backgound Block Error TCA information.
    BbeTca Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca

    // Errored Seconds TCA information.
    EsTca Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca

    // Backgound Block Error information.
    Bbe Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe

    // Errored Seconds information .
    Es Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es

    // Severly Errored Seconds information.
    Ses Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses

    // Unavailability Seconds information.
    Uas Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas

    // Failure Count information.
    Fc Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc

    // Backgound Block Error Rate information.
    Bber Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber

    // Errored Seconds Rate information.
    Esr Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr

    // Severly Errored Seconds Rate information.
    Sesr Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr

    // Trail Trace Identifier information.
    Tti Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti
}

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetFilter() yfilter.YFilter { return otuInfo.YFilter }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) SetFilter(yf yfilter.YFilter) { otuInfo.YFilter = yf }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetGoName(yname string) string {
    if yname == "bei" { return "Bei" }
    if yname == "bip" { return "Bip" }
    if yname == "los" { return "Los" }
    if yname == "lof" { return "Lof" }
    if yname == "lom" { return "Lom" }
    if yname == "oof" { return "Oof" }
    if yname == "oom" { return "Oom" }
    if yname == "ais" { return "Ais" }
    if yname == "iae" { return "Iae" }
    if yname == "bdi" { return "Bdi" }
    if yname == "tim" { return "Tim" }
    if yname == "eoc" { return "Eoc" }
    if yname == "sf-ber" { return "SfBer" }
    if yname == "sd-ber" { return "SdBer" }
    if yname == "prefec-sf-ber" { return "PrefecSfBer" }
    if yname == "prefec-sd-ber" { return "PrefecSdBer" }
    if yname == "bbe-tca" { return "BbeTca" }
    if yname == "es-tca" { return "EsTca" }
    if yname == "bbe" { return "Bbe" }
    if yname == "es" { return "Es" }
    if yname == "ses" { return "Ses" }
    if yname == "uas" { return "Uas" }
    if yname == "fc" { return "Fc" }
    if yname == "bber" { return "Bber" }
    if yname == "esr" { return "Esr" }
    if yname == "sesr" { return "Sesr" }
    if yname == "tti" { return "Tti" }
    return ""
}

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetSegmentPath() string {
    return "otu-info"
}

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "los" {
        return &otuInfo.Los
    }
    if childYangName == "lof" {
        return &otuInfo.Lof
    }
    if childYangName == "lom" {
        return &otuInfo.Lom
    }
    if childYangName == "oof" {
        return &otuInfo.Oof
    }
    if childYangName == "oom" {
        return &otuInfo.Oom
    }
    if childYangName == "ais" {
        return &otuInfo.Ais
    }
    if childYangName == "iae" {
        return &otuInfo.Iae
    }
    if childYangName == "bdi" {
        return &otuInfo.Bdi
    }
    if childYangName == "tim" {
        return &otuInfo.Tim
    }
    if childYangName == "eoc" {
        return &otuInfo.Eoc
    }
    if childYangName == "sf-ber" {
        return &otuInfo.SfBer
    }
    if childYangName == "sd-ber" {
        return &otuInfo.SdBer
    }
    if childYangName == "prefec-sf-ber" {
        return &otuInfo.PrefecSfBer
    }
    if childYangName == "prefec-sd-ber" {
        return &otuInfo.PrefecSdBer
    }
    if childYangName == "bbe-tca" {
        return &otuInfo.BbeTca
    }
    if childYangName == "es-tca" {
        return &otuInfo.EsTca
    }
    if childYangName == "bbe" {
        return &otuInfo.Bbe
    }
    if childYangName == "es" {
        return &otuInfo.Es
    }
    if childYangName == "ses" {
        return &otuInfo.Ses
    }
    if childYangName == "uas" {
        return &otuInfo.Uas
    }
    if childYangName == "fc" {
        return &otuInfo.Fc
    }
    if childYangName == "bber" {
        return &otuInfo.Bber
    }
    if childYangName == "esr" {
        return &otuInfo.Esr
    }
    if childYangName == "sesr" {
        return &otuInfo.Sesr
    }
    if childYangName == "tti" {
        return &otuInfo.Tti
    }
    return nil
}

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["los"] = &otuInfo.Los
    children["lof"] = &otuInfo.Lof
    children["lom"] = &otuInfo.Lom
    children["oof"] = &otuInfo.Oof
    children["oom"] = &otuInfo.Oom
    children["ais"] = &otuInfo.Ais
    children["iae"] = &otuInfo.Iae
    children["bdi"] = &otuInfo.Bdi
    children["tim"] = &otuInfo.Tim
    children["eoc"] = &otuInfo.Eoc
    children["sf-ber"] = &otuInfo.SfBer
    children["sd-ber"] = &otuInfo.SdBer
    children["prefec-sf-ber"] = &otuInfo.PrefecSfBer
    children["prefec-sd-ber"] = &otuInfo.PrefecSdBer
    children["bbe-tca"] = &otuInfo.BbeTca
    children["es-tca"] = &otuInfo.EsTca
    children["bbe"] = &otuInfo.Bbe
    children["es"] = &otuInfo.Es
    children["ses"] = &otuInfo.Ses
    children["uas"] = &otuInfo.Uas
    children["fc"] = &otuInfo.Fc
    children["bber"] = &otuInfo.Bber
    children["esr"] = &otuInfo.Esr
    children["sesr"] = &otuInfo.Sesr
    children["tti"] = &otuInfo.Tti
    return children
}

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bei"] = otuInfo.Bei
    leafs["bip"] = otuInfo.Bip
    return leafs
}

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetBundleName() string { return "cisco_ios_xr" }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetYangName() string { return "otu-info" }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) SetParent(parent types.Entity) { otuInfo.parent = parent }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetParent() types.Entity { return otuInfo.parent }

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetParentYangName() string { return "g709-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los
// Loss of Signal information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los struct {
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

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetFilter() yfilter.YFilter { return los.YFilter }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) SetFilter(yf yfilter.YFilter) { los.YFilter = yf }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetSegmentPath() string {
    return "los"
}

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = los.ReportingEnabled
    leafs["is-detected"] = los.IsDetected
    leafs["is-asserted"] = los.IsAsserted
    leafs["counter"] = los.Counter
    return leafs
}

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetBundleName() string { return "cisco_ios_xr" }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetYangName() string { return "los" }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) SetParent(parent types.Entity) { los.parent = parent }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetParent() types.Entity { return los.parent }

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof
// Loss of Frame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof struct {
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

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetFilter() yfilter.YFilter { return lof.YFilter }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) SetFilter(yf yfilter.YFilter) { lof.YFilter = yf }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetSegmentPath() string {
    return "lof"
}

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = lof.ReportingEnabled
    leafs["is-detected"] = lof.IsDetected
    leafs["is-asserted"] = lof.IsAsserted
    leafs["counter"] = lof.Counter
    return leafs
}

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetBundleName() string { return "cisco_ios_xr" }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetYangName() string { return "lof" }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) SetParent(parent types.Entity) { lof.parent = parent }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetParent() types.Entity { return lof.parent }

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom
// Loss of MultiFrame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom struct {
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

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetFilter() yfilter.YFilter { return lom.YFilter }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) SetFilter(yf yfilter.YFilter) { lom.YFilter = yf }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetSegmentPath() string {
    return "lom"
}

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = lom.ReportingEnabled
    leafs["is-detected"] = lom.IsDetected
    leafs["is-asserted"] = lom.IsAsserted
    leafs["counter"] = lom.Counter
    return leafs
}

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetBundleName() string { return "cisco_ios_xr" }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetYangName() string { return "lom" }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) SetParent(parent types.Entity) { lom.parent = parent }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetParent() types.Entity { return lom.parent }

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof
// Out of Frame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof struct {
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

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetFilter() yfilter.YFilter { return oof.YFilter }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) SetFilter(yf yfilter.YFilter) { oof.YFilter = yf }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetSegmentPath() string {
    return "oof"
}

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = oof.ReportingEnabled
    leafs["is-detected"] = oof.IsDetected
    leafs["is-asserted"] = oof.IsAsserted
    leafs["counter"] = oof.Counter
    return leafs
}

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetBundleName() string { return "cisco_ios_xr" }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetYangName() string { return "oof" }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) SetParent(parent types.Entity) { oof.parent = parent }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetParent() types.Entity { return oof.parent }

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom
// Out of MultiFrame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom struct {
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

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetFilter() yfilter.YFilter { return oom.YFilter }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) SetFilter(yf yfilter.YFilter) { oom.YFilter = yf }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetSegmentPath() string {
    return "oom"
}

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = oom.ReportingEnabled
    leafs["is-detected"] = oom.IsDetected
    leafs["is-asserted"] = oom.IsAsserted
    leafs["counter"] = oom.Counter
    return leafs
}

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetBundleName() string { return "cisco_ios_xr" }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetYangName() string { return "oom" }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) SetParent(parent types.Entity) { oom.parent = parent }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetParent() types.Entity { return oom.parent }

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais
// Alarm Indication Signal information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais struct {
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

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetFilter() yfilter.YFilter { return ais.YFilter }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) SetFilter(yf yfilter.YFilter) { ais.YFilter = yf }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetSegmentPath() string {
    return "ais"
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = ais.ReportingEnabled
    leafs["is-detected"] = ais.IsDetected
    leafs["is-asserted"] = ais.IsAsserted
    leafs["counter"] = ais.Counter
    return leafs
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetBundleName() string { return "cisco_ios_xr" }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetYangName() string { return "ais" }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) SetParent(parent types.Entity) { ais.parent = parent }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetParent() types.Entity { return ais.parent }

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae
// Incoming Alignment Error information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae struct {
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

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetFilter() yfilter.YFilter { return iae.YFilter }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) SetFilter(yf yfilter.YFilter) { iae.YFilter = yf }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetSegmentPath() string {
    return "iae"
}

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = iae.ReportingEnabled
    leafs["is-detected"] = iae.IsDetected
    leafs["is-asserted"] = iae.IsAsserted
    leafs["counter"] = iae.Counter
    return leafs
}

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetBundleName() string { return "cisco_ios_xr" }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetYangName() string { return "iae" }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) SetParent(parent types.Entity) { iae.parent = parent }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetParent() types.Entity { return iae.parent }

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi
// Backward Defect Indication information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi struct {
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

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetFilter() yfilter.YFilter { return bdi.YFilter }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) SetFilter(yf yfilter.YFilter) { bdi.YFilter = yf }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetSegmentPath() string {
    return "bdi"
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = bdi.ReportingEnabled
    leafs["is-detected"] = bdi.IsDetected
    leafs["is-asserted"] = bdi.IsAsserted
    leafs["counter"] = bdi.Counter
    return leafs
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetBundleName() string { return "cisco_ios_xr" }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetYangName() string { return "bdi" }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) SetParent(parent types.Entity) { bdi.parent = parent }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetParent() types.Entity { return bdi.parent }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim
// Trace Identifier Mismatch information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim struct {
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

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetFilter() yfilter.YFilter { return tim.YFilter }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) SetFilter(yf yfilter.YFilter) { tim.YFilter = yf }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetSegmentPath() string {
    return "tim"
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = tim.ReportingEnabled
    leafs["is-detected"] = tim.IsDetected
    leafs["is-asserted"] = tim.IsAsserted
    leafs["counter"] = tim.Counter
    return leafs
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetBundleName() string { return "cisco_ios_xr" }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetYangName() string { return "tim" }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) SetParent(parent types.Entity) { tim.parent = parent }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetParent() types.Entity { return tim.parent }

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc
// GCC End of Channel information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc struct {
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

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetFilter() yfilter.YFilter { return eoc.YFilter }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) SetFilter(yf yfilter.YFilter) { eoc.YFilter = yf }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetSegmentPath() string {
    return "eoc"
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = eoc.ReportingEnabled
    leafs["is-detected"] = eoc.IsDetected
    leafs["is-asserted"] = eoc.IsAsserted
    leafs["counter"] = eoc.Counter
    return leafs
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetBundleName() string { return "cisco_ios_xr" }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetYangName() string { return "eoc" }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) SetParent(parent types.Entity) { eoc.parent = parent }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetParent() types.Entity { return eoc.parent }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer
// Signal Fail  BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetFilter() yfilter.YFilter { return sfBer.YFilter }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) SetFilter(yf yfilter.YFilter) { sfBer.YFilter = yf }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetSegmentPath() string {
    return "sf-ber"
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = sfBer.ReportingEnabled
    leafs["is-detected"] = sfBer.IsDetected
    leafs["is-asserted"] = sfBer.IsAsserted
    leafs["threshold"] = sfBer.Threshold
    leafs["counter"] = sfBer.Counter
    return leafs
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetBundleName() string { return "cisco_ios_xr" }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetYangName() string { return "sf-ber" }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) SetParent(parent types.Entity) { sfBer.parent = parent }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetParent() types.Entity { return sfBer.parent }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer
// Signal Degrade BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetFilter() yfilter.YFilter { return sdBer.YFilter }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) SetFilter(yf yfilter.YFilter) { sdBer.YFilter = yf }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetSegmentPath() string {
    return "sd-ber"
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = sdBer.ReportingEnabled
    leafs["is-detected"] = sdBer.IsDetected
    leafs["is-asserted"] = sdBer.IsAsserted
    leafs["threshold"] = sdBer.Threshold
    leafs["counter"] = sdBer.Counter
    return leafs
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetBundleName() string { return "cisco_ios_xr" }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetYangName() string { return "sd-ber" }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) SetParent(parent types.Entity) { sdBer.parent = parent }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetParent() types.Entity { return sdBer.parent }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer
// Prefec Signal Fail BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetFilter() yfilter.YFilter { return prefecSfBer.YFilter }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) SetFilter(yf yfilter.YFilter) { prefecSfBer.YFilter = yf }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetSegmentPath() string {
    return "prefec-sf-ber"
}

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = prefecSfBer.ReportingEnabled
    leafs["is-detected"] = prefecSfBer.IsDetected
    leafs["is-asserted"] = prefecSfBer.IsAsserted
    leafs["threshold"] = prefecSfBer.Threshold
    leafs["counter"] = prefecSfBer.Counter
    return leafs
}

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetBundleName() string { return "cisco_ios_xr" }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetYangName() string { return "prefec-sf-ber" }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) SetParent(parent types.Entity) { prefecSfBer.parent = parent }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetParent() types.Entity { return prefecSfBer.parent }

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer
// Prefec Signal Degrade BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetFilter() yfilter.YFilter { return prefecSdBer.YFilter }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) SetFilter(yf yfilter.YFilter) { prefecSdBer.YFilter = yf }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetSegmentPath() string {
    return "prefec-sd-ber"
}

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = prefecSdBer.ReportingEnabled
    leafs["is-detected"] = prefecSdBer.IsDetected
    leafs["is-asserted"] = prefecSdBer.IsAsserted
    leafs["threshold"] = prefecSdBer.Threshold
    leafs["counter"] = prefecSdBer.Counter
    return leafs
}

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetBundleName() string { return "cisco_ios_xr" }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetYangName() string { return "prefec-sd-ber" }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) SetParent(parent types.Entity) { prefecSdBer.parent = parent }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetParent() types.Entity { return prefecSdBer.parent }

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca
//  Backgound Block Error TCA information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetFilter() yfilter.YFilter { return bbeTca.YFilter }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) SetFilter(yf yfilter.YFilter) { bbeTca.YFilter = yf }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetSegmentPath() string {
    return "bbe-tca"
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = bbeTca.ReportingEnabled
    leafs["is-detected"] = bbeTca.IsDetected
    leafs["is-asserted"] = bbeTca.IsAsserted
    leafs["threshold"] = bbeTca.Threshold
    leafs["counter"] = bbeTca.Counter
    return leafs
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetBundleName() string { return "cisco_ios_xr" }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetYangName() string { return "bbe-tca" }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) SetParent(parent types.Entity) { bbeTca.parent = parent }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetParent() types.Entity { return bbeTca.parent }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca
// Errored Seconds TCA information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetFilter() yfilter.YFilter { return esTca.YFilter }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) SetFilter(yf yfilter.YFilter) { esTca.YFilter = yf }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetSegmentPath() string {
    return "es-tca"
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = esTca.ReportingEnabled
    leafs["is-detected"] = esTca.IsDetected
    leafs["is-asserted"] = esTca.IsAsserted
    leafs["threshold"] = esTca.Threshold
    leafs["counter"] = esTca.Counter
    return leafs
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetBundleName() string { return "cisco_ios_xr" }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetYangName() string { return "es-tca" }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) SetParent(parent types.Entity) { esTca.parent = parent }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetParent() types.Entity { return esTca.parent }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe
// Backgound Block Error information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetFilter() yfilter.YFilter { return bbe.YFilter }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) SetFilter(yf yfilter.YFilter) { bbe.YFilter = yf }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetSegmentPath() string {
    return "bbe"
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = bbe.Counter
    return leafs
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetBundleName() string { return "cisco_ios_xr" }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetYangName() string { return "bbe" }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) SetParent(parent types.Entity) { bbe.parent = parent }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetParent() types.Entity { return bbe.parent }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es
// Errored Seconds information 
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetFilter() yfilter.YFilter { return es.YFilter }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) SetFilter(yf yfilter.YFilter) { es.YFilter = yf }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetSegmentPath() string {
    return "es"
}

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = es.Counter
    return leafs
}

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetBundleName() string { return "cisco_ios_xr" }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetYangName() string { return "es" }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) SetParent(parent types.Entity) { es.parent = parent }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetParent() types.Entity { return es.parent }

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses
// Severly Errored Seconds information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetFilter() yfilter.YFilter { return ses.YFilter }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) SetFilter(yf yfilter.YFilter) { ses.YFilter = yf }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetSegmentPath() string {
    return "ses"
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = ses.Counter
    return leafs
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetBundleName() string { return "cisco_ios_xr" }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetYangName() string { return "ses" }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) SetParent(parent types.Entity) { ses.parent = parent }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetParent() types.Entity { return ses.parent }

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas
// Unavailability Seconds information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetFilter() yfilter.YFilter { return uas.YFilter }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) SetFilter(yf yfilter.YFilter) { uas.YFilter = yf }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetSegmentPath() string {
    return "uas"
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = uas.Counter
    return leafs
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetBundleName() string { return "cisco_ios_xr" }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetYangName() string { return "uas" }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) SetParent(parent types.Entity) { uas.parent = parent }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetParent() types.Entity { return uas.parent }

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc
// Failure Count information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetFilter() yfilter.YFilter { return fc.YFilter }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) SetFilter(yf yfilter.YFilter) { fc.YFilter = yf }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetSegmentPath() string {
    return "fc"
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = fc.Counter
    return leafs
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetBundleName() string { return "cisco_ios_xr" }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetYangName() string { return "fc" }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) SetParent(parent types.Entity) { fc.parent = parent }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetParent() types.Entity { return fc.parent }

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber
// Backgound Block Error Rate information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetFilter() yfilter.YFilter { return bber.YFilter }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) SetFilter(yf yfilter.YFilter) { bber.YFilter = yf }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetSegmentPath() string {
    return "bber"
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = bber.Counter
    return leafs
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetBundleName() string { return "cisco_ios_xr" }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetYangName() string { return "bber" }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) SetParent(parent types.Entity) { bber.parent = parent }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetParent() types.Entity { return bber.parent }

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr
// Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetFilter() yfilter.YFilter { return esr.YFilter }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) SetFilter(yf yfilter.YFilter) { esr.YFilter = yf }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetSegmentPath() string {
    return "esr"
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = esr.Counter
    return leafs
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetBundleName() string { return "cisco_ios_xr" }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetYangName() string { return "esr" }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) SetParent(parent types.Entity) { esr.parent = parent }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetParent() types.Entity { return esr.parent }

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr
// Severly Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetFilter() yfilter.YFilter { return sesr.YFilter }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) SetFilter(yf yfilter.YFilter) { sesr.YFilter = yf }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetSegmentPath() string {
    return "sesr"
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = sesr.Counter
    return leafs
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetBundleName() string { return "cisco_ios_xr" }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetYangName() string { return "sesr" }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) SetParent(parent types.Entity) { sesr.parent = parent }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetParent() types.Entity { return sesr.parent }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti
// Trail Trace Identifier information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of String. The type is interface{} with range: 0..4294967295.
    TxStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    ExpectedStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    RxStringType interface{}

    // Tx TTI String . The type is string with length: 0..129.
    TxTti interface{}

    // Tx SAPI[0] Field. The type is string with length: 0..5.
    TxSapi0 interface{}

    // Tx SAPI[1-15] Field. The type is string with length: 0..16.
    TxSapi interface{}

    // Tx SAPI Range String. The type is string with length: 0..6.
    TxSapiRange interface{}

    // Tx DAPI[0] Field. The type is string with length: 0..5.
    TxDapi0 interface{}

    // Tx DAPI[1-15] Field. The type is string with length: 0..16.
    TxDapi interface{}

    // Tx DAPI Range String. The type is string with length: 0..6.
    TxDapiRange interface{}

    // Tx Operator Specific Field. The type is string with length: 0..33.
    TxOperSpec interface{}

    // Tx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    TxOperSpecRange interface{}

    // Rx TTI String . The type is string with length: 0..129.
    RxTti interface{}

    // Rx SAPI[0] Field. The type is string with length: 0..5.
    RxSapi0 interface{}

    // Rx SAPI[1-15] Field. The type is string with length: 0..16.
    RxSapi interface{}

    // Rx SAPI Range String. The type is string with length: 0..6.
    RxSapiRange interface{}

    // Rx DAPI[0] Field. The type is string with length: 0..5.
    RxDapi0 interface{}

    // Rx DAPI[1-15] Field. The type is string with length: 0..16.
    RxDapi interface{}

    // Rx DAPI Range String. The type is string with length: 0..6.
    RxDapiRange interface{}

    // Rx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    RxOperSpecRange interface{}

    // Rx Operator Specific Field. The type is string with length: 0..33.
    RxOperSpec interface{}

    // Expected TTI String. The type is string with length: 0..129.
    ExpectedTti interface{}

    // Expected SAPI[0] Field. The type is string with length: 0..5.
    ExpectedSapi0 interface{}

    // Expected SAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedSapi interface{}

    // Expected SAPI Range String. The type is string with length: 0..6.
    ExpSapiRange interface{}

    // Expected DAPI[0] Field. The type is string with length: 0..5.
    ExpectedDapi0 interface{}

    // Expected DAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedDapi interface{}

    // Expected DAPI Range String. The type is string with length: 0..6.
    ExpDapiRange interface{}

    // Expected Operator Specific Field. The type is string with length: 0..33.
    ExpectedOperSpec interface{}

    // Expected Operator Specific Field Range String. The type is string with
    // length: 0..6.
    ExpOperSpecRange interface{}
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetFilter() yfilter.YFilter { return tti.YFilter }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) SetFilter(yf yfilter.YFilter) { tti.YFilter = yf }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetGoName(yname string) string {
    if yname == "tx-string-type" { return "TxStringType" }
    if yname == "expected-string-type" { return "ExpectedStringType" }
    if yname == "rx-string-type" { return "RxStringType" }
    if yname == "tx-tti" { return "TxTti" }
    if yname == "tx-sapi0" { return "TxSapi0" }
    if yname == "tx-sapi" { return "TxSapi" }
    if yname == "tx-sapi-range" { return "TxSapiRange" }
    if yname == "tx-dapi0" { return "TxDapi0" }
    if yname == "tx-dapi" { return "TxDapi" }
    if yname == "tx-dapi-range" { return "TxDapiRange" }
    if yname == "tx-oper-spec" { return "TxOperSpec" }
    if yname == "tx-oper-spec-range" { return "TxOperSpecRange" }
    if yname == "rx-tti" { return "RxTti" }
    if yname == "rx-sapi0" { return "RxSapi0" }
    if yname == "rx-sapi" { return "RxSapi" }
    if yname == "rx-sapi-range" { return "RxSapiRange" }
    if yname == "rx-dapi0" { return "RxDapi0" }
    if yname == "rx-dapi" { return "RxDapi" }
    if yname == "rx-dapi-range" { return "RxDapiRange" }
    if yname == "rx-oper-spec-range" { return "RxOperSpecRange" }
    if yname == "rx-oper-spec" { return "RxOperSpec" }
    if yname == "expected-tti" { return "ExpectedTti" }
    if yname == "expected-sapi0" { return "ExpectedSapi0" }
    if yname == "expected-sapi" { return "ExpectedSapi" }
    if yname == "exp-sapi-range" { return "ExpSapiRange" }
    if yname == "expected-dapi0" { return "ExpectedDapi0" }
    if yname == "expected-dapi" { return "ExpectedDapi" }
    if yname == "exp-dapi-range" { return "ExpDapiRange" }
    if yname == "expected-oper-spec" { return "ExpectedOperSpec" }
    if yname == "exp-oper-spec-range" { return "ExpOperSpecRange" }
    return ""
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetSegmentPath() string {
    return "tti"
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-string-type"] = tti.TxStringType
    leafs["expected-string-type"] = tti.ExpectedStringType
    leafs["rx-string-type"] = tti.RxStringType
    leafs["tx-tti"] = tti.TxTti
    leafs["tx-sapi0"] = tti.TxSapi0
    leafs["tx-sapi"] = tti.TxSapi
    leafs["tx-sapi-range"] = tti.TxSapiRange
    leafs["tx-dapi0"] = tti.TxDapi0
    leafs["tx-dapi"] = tti.TxDapi
    leafs["tx-dapi-range"] = tti.TxDapiRange
    leafs["tx-oper-spec"] = tti.TxOperSpec
    leafs["tx-oper-spec-range"] = tti.TxOperSpecRange
    leafs["rx-tti"] = tti.RxTti
    leafs["rx-sapi0"] = tti.RxSapi0
    leafs["rx-sapi"] = tti.RxSapi
    leafs["rx-sapi-range"] = tti.RxSapiRange
    leafs["rx-dapi0"] = tti.RxDapi0
    leafs["rx-dapi"] = tti.RxDapi
    leafs["rx-dapi-range"] = tti.RxDapiRange
    leafs["rx-oper-spec-range"] = tti.RxOperSpecRange
    leafs["rx-oper-spec"] = tti.RxOperSpec
    leafs["expected-tti"] = tti.ExpectedTti
    leafs["expected-sapi0"] = tti.ExpectedSapi0
    leafs["expected-sapi"] = tti.ExpectedSapi
    leafs["exp-sapi-range"] = tti.ExpSapiRange
    leafs["expected-dapi0"] = tti.ExpectedDapi0
    leafs["expected-dapi"] = tti.ExpectedDapi
    leafs["exp-dapi-range"] = tti.ExpDapiRange
    leafs["expected-oper-spec"] = tti.ExpectedOperSpec
    leafs["exp-oper-spec-range"] = tti.ExpOperSpecRange
    return leafs
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetBundleName() string { return "cisco_ios_xr" }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetYangName() string { return "tti" }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) SetParent(parent types.Entity) { tti.parent = parent }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetParent() types.Entity { return tti.parent }

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetParentYangName() string { return "otu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo
// ODU layer Information
type Dwdm_Ports_Port_Info_G709Info_OduInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bit Interleave Parity(BIP) counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bip interface{}

    // Backward Error Indication counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bei interface{}

    // Open Connection Indiction information.
    Oci Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci

    // Alarm Indication Signal information.
    Ais Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais

    // Upstream Connection Locked information.
    Lck Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck

    // Backward Defect Indication information.
    Bdi Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi

    // GCC End of Channel information.
    Eoc Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc

    // Payload Type Identifier Mismatch information.
    Ptim Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim

    // Trace Identifier Mismatch information.
    Tim Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim

    // Signal Fail  BER information.
    SfBer Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer

    // Signal Degrade BER information.
    SdBer Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer

    // Background Block Error TCA information.
    BbeTca Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca

    // Errored Seconds TCA information.
    EsTca Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca

    // Background Block Error information.
    Bbe Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe

    // Errored Seconds information.
    Es Dwdm_Ports_Port_Info_G709Info_OduInfo_Es

    // Severly Errored Seconds information.
    Ses Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses

    // Unavailability Seconds information.
    Uas Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas

    // Failure count information.
    Fc Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc

    // Background Block Error Rate count information.
    Bber Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber

    // Errored Seconds Rate information.
    Esr Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr

    // Severly Errored Seconds Rate information.
    Sesr Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr

    // Trail Trace Identifier information.
    Tti Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti
}

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetFilter() yfilter.YFilter { return oduInfo.YFilter }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) SetFilter(yf yfilter.YFilter) { oduInfo.YFilter = yf }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetGoName(yname string) string {
    if yname == "bip" { return "Bip" }
    if yname == "bei" { return "Bei" }
    if yname == "oci" { return "Oci" }
    if yname == "ais" { return "Ais" }
    if yname == "lck" { return "Lck" }
    if yname == "bdi" { return "Bdi" }
    if yname == "eoc" { return "Eoc" }
    if yname == "ptim" { return "Ptim" }
    if yname == "tim" { return "Tim" }
    if yname == "sf-ber" { return "SfBer" }
    if yname == "sd-ber" { return "SdBer" }
    if yname == "bbe-tca" { return "BbeTca" }
    if yname == "es-tca" { return "EsTca" }
    if yname == "bbe" { return "Bbe" }
    if yname == "es" { return "Es" }
    if yname == "ses" { return "Ses" }
    if yname == "uas" { return "Uas" }
    if yname == "fc" { return "Fc" }
    if yname == "bber" { return "Bber" }
    if yname == "esr" { return "Esr" }
    if yname == "sesr" { return "Sesr" }
    if yname == "tti" { return "Tti" }
    return ""
}

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetSegmentPath() string {
    return "odu-info"
}

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oci" {
        return &oduInfo.Oci
    }
    if childYangName == "ais" {
        return &oduInfo.Ais
    }
    if childYangName == "lck" {
        return &oduInfo.Lck
    }
    if childYangName == "bdi" {
        return &oduInfo.Bdi
    }
    if childYangName == "eoc" {
        return &oduInfo.Eoc
    }
    if childYangName == "ptim" {
        return &oduInfo.Ptim
    }
    if childYangName == "tim" {
        return &oduInfo.Tim
    }
    if childYangName == "sf-ber" {
        return &oduInfo.SfBer
    }
    if childYangName == "sd-ber" {
        return &oduInfo.SdBer
    }
    if childYangName == "bbe-tca" {
        return &oduInfo.BbeTca
    }
    if childYangName == "es-tca" {
        return &oduInfo.EsTca
    }
    if childYangName == "bbe" {
        return &oduInfo.Bbe
    }
    if childYangName == "es" {
        return &oduInfo.Es
    }
    if childYangName == "ses" {
        return &oduInfo.Ses
    }
    if childYangName == "uas" {
        return &oduInfo.Uas
    }
    if childYangName == "fc" {
        return &oduInfo.Fc
    }
    if childYangName == "bber" {
        return &oduInfo.Bber
    }
    if childYangName == "esr" {
        return &oduInfo.Esr
    }
    if childYangName == "sesr" {
        return &oduInfo.Sesr
    }
    if childYangName == "tti" {
        return &oduInfo.Tti
    }
    return nil
}

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["oci"] = &oduInfo.Oci
    children["ais"] = &oduInfo.Ais
    children["lck"] = &oduInfo.Lck
    children["bdi"] = &oduInfo.Bdi
    children["eoc"] = &oduInfo.Eoc
    children["ptim"] = &oduInfo.Ptim
    children["tim"] = &oduInfo.Tim
    children["sf-ber"] = &oduInfo.SfBer
    children["sd-ber"] = &oduInfo.SdBer
    children["bbe-tca"] = &oduInfo.BbeTca
    children["es-tca"] = &oduInfo.EsTca
    children["bbe"] = &oduInfo.Bbe
    children["es"] = &oduInfo.Es
    children["ses"] = &oduInfo.Ses
    children["uas"] = &oduInfo.Uas
    children["fc"] = &oduInfo.Fc
    children["bber"] = &oduInfo.Bber
    children["esr"] = &oduInfo.Esr
    children["sesr"] = &oduInfo.Sesr
    children["tti"] = &oduInfo.Tti
    return children
}

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bip"] = oduInfo.Bip
    leafs["bei"] = oduInfo.Bei
    return leafs
}

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetBundleName() string { return "cisco_ios_xr" }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetYangName() string { return "odu-info" }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) SetParent(parent types.Entity) { oduInfo.parent = parent }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetParent() types.Entity { return oduInfo.parent }

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetParentYangName() string { return "g709-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci
// Open Connection Indiction information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci struct {
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

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetFilter() yfilter.YFilter { return oci.YFilter }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) SetFilter(yf yfilter.YFilter) { oci.YFilter = yf }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetSegmentPath() string {
    return "oci"
}

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = oci.ReportingEnabled
    leafs["is-detected"] = oci.IsDetected
    leafs["is-asserted"] = oci.IsAsserted
    leafs["counter"] = oci.Counter
    return leafs
}

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetBundleName() string { return "cisco_ios_xr" }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetYangName() string { return "oci" }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) SetParent(parent types.Entity) { oci.parent = parent }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetParent() types.Entity { return oci.parent }

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais
// Alarm Indication Signal information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais struct {
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

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetFilter() yfilter.YFilter { return ais.YFilter }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) SetFilter(yf yfilter.YFilter) { ais.YFilter = yf }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetSegmentPath() string {
    return "ais"
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = ais.ReportingEnabled
    leafs["is-detected"] = ais.IsDetected
    leafs["is-asserted"] = ais.IsAsserted
    leafs["counter"] = ais.Counter
    return leafs
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetBundleName() string { return "cisco_ios_xr" }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetYangName() string { return "ais" }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) SetParent(parent types.Entity) { ais.parent = parent }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetParent() types.Entity { return ais.parent }

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck
// Upstream Connection Locked information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck struct {
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

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetFilter() yfilter.YFilter { return lck.YFilter }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) SetFilter(yf yfilter.YFilter) { lck.YFilter = yf }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetSegmentPath() string {
    return "lck"
}

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = lck.ReportingEnabled
    leafs["is-detected"] = lck.IsDetected
    leafs["is-asserted"] = lck.IsAsserted
    leafs["counter"] = lck.Counter
    return leafs
}

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetBundleName() string { return "cisco_ios_xr" }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetYangName() string { return "lck" }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) SetParent(parent types.Entity) { lck.parent = parent }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetParent() types.Entity { return lck.parent }

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi
// Backward Defect Indication information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi struct {
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

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetFilter() yfilter.YFilter { return bdi.YFilter }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) SetFilter(yf yfilter.YFilter) { bdi.YFilter = yf }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetSegmentPath() string {
    return "bdi"
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = bdi.ReportingEnabled
    leafs["is-detected"] = bdi.IsDetected
    leafs["is-asserted"] = bdi.IsAsserted
    leafs["counter"] = bdi.Counter
    return leafs
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetBundleName() string { return "cisco_ios_xr" }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetYangName() string { return "bdi" }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) SetParent(parent types.Entity) { bdi.parent = parent }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetParent() types.Entity { return bdi.parent }

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc
// GCC End of Channel information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc struct {
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

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetFilter() yfilter.YFilter { return eoc.YFilter }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) SetFilter(yf yfilter.YFilter) { eoc.YFilter = yf }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetSegmentPath() string {
    return "eoc"
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = eoc.ReportingEnabled
    leafs["is-detected"] = eoc.IsDetected
    leafs["is-asserted"] = eoc.IsAsserted
    leafs["counter"] = eoc.Counter
    return leafs
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetBundleName() string { return "cisco_ios_xr" }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetYangName() string { return "eoc" }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) SetParent(parent types.Entity) { eoc.parent = parent }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetParent() types.Entity { return eoc.parent }

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim
// Payload Type Identifier Mismatch information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim struct {
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

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetFilter() yfilter.YFilter { return ptim.YFilter }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) SetFilter(yf yfilter.YFilter) { ptim.YFilter = yf }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetSegmentPath() string {
    return "ptim"
}

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = ptim.ReportingEnabled
    leafs["is-detected"] = ptim.IsDetected
    leafs["is-asserted"] = ptim.IsAsserted
    leafs["counter"] = ptim.Counter
    return leafs
}

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetBundleName() string { return "cisco_ios_xr" }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetYangName() string { return "ptim" }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) SetParent(parent types.Entity) { ptim.parent = parent }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetParent() types.Entity { return ptim.parent }

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim
// Trace Identifier Mismatch information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim struct {
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

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetFilter() yfilter.YFilter { return tim.YFilter }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) SetFilter(yf yfilter.YFilter) { tim.YFilter = yf }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetSegmentPath() string {
    return "tim"
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = tim.ReportingEnabled
    leafs["is-detected"] = tim.IsDetected
    leafs["is-asserted"] = tim.IsAsserted
    leafs["counter"] = tim.Counter
    return leafs
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetBundleName() string { return "cisco_ios_xr" }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetYangName() string { return "tim" }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) SetParent(parent types.Entity) { tim.parent = parent }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetParent() types.Entity { return tim.parent }

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer
// Signal Fail  BER information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetFilter() yfilter.YFilter { return sfBer.YFilter }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) SetFilter(yf yfilter.YFilter) { sfBer.YFilter = yf }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetSegmentPath() string {
    return "sf-ber"
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = sfBer.ReportingEnabled
    leafs["is-detected"] = sfBer.IsDetected
    leafs["is-asserted"] = sfBer.IsAsserted
    leafs["threshold"] = sfBer.Threshold
    leafs["counter"] = sfBer.Counter
    return leafs
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetBundleName() string { return "cisco_ios_xr" }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetYangName() string { return "sf-ber" }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) SetParent(parent types.Entity) { sfBer.parent = parent }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetParent() types.Entity { return sfBer.parent }

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer
// Signal Degrade BER information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetFilter() yfilter.YFilter { return sdBer.YFilter }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) SetFilter(yf yfilter.YFilter) { sdBer.YFilter = yf }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetSegmentPath() string {
    return "sd-ber"
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = sdBer.ReportingEnabled
    leafs["is-detected"] = sdBer.IsDetected
    leafs["is-asserted"] = sdBer.IsAsserted
    leafs["threshold"] = sdBer.Threshold
    leafs["counter"] = sdBer.Counter
    return leafs
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetBundleName() string { return "cisco_ios_xr" }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetYangName() string { return "sd-ber" }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) SetParent(parent types.Entity) { sdBer.parent = parent }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetParent() types.Entity { return sdBer.parent }

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca
// Background Block Error TCA information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetFilter() yfilter.YFilter { return bbeTca.YFilter }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) SetFilter(yf yfilter.YFilter) { bbeTca.YFilter = yf }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetSegmentPath() string {
    return "bbe-tca"
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = bbeTca.ReportingEnabled
    leafs["is-detected"] = bbeTca.IsDetected
    leafs["is-asserted"] = bbeTca.IsAsserted
    leafs["threshold"] = bbeTca.Threshold
    leafs["counter"] = bbeTca.Counter
    return leafs
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetBundleName() string { return "cisco_ios_xr" }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetYangName() string { return "bbe-tca" }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) SetParent(parent types.Entity) { bbeTca.parent = parent }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetParent() types.Entity { return bbeTca.parent }

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca
// Errored Seconds TCA information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetFilter() yfilter.YFilter { return esTca.YFilter }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) SetFilter(yf yfilter.YFilter) { esTca.YFilter = yf }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetGoName(yname string) string {
    if yname == "reporting-enabled" { return "ReportingEnabled" }
    if yname == "is-detected" { return "IsDetected" }
    if yname == "is-asserted" { return "IsAsserted" }
    if yname == "threshold" { return "Threshold" }
    if yname == "counter" { return "Counter" }
    return ""
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetSegmentPath() string {
    return "es-tca"
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reporting-enabled"] = esTca.ReportingEnabled
    leafs["is-detected"] = esTca.IsDetected
    leafs["is-asserted"] = esTca.IsAsserted
    leafs["threshold"] = esTca.Threshold
    leafs["counter"] = esTca.Counter
    return leafs
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetBundleName() string { return "cisco_ios_xr" }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetYangName() string { return "es-tca" }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) SetParent(parent types.Entity) { esTca.parent = parent }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetParent() types.Entity { return esTca.parent }

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe
// Background Block Error information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetFilter() yfilter.YFilter { return bbe.YFilter }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) SetFilter(yf yfilter.YFilter) { bbe.YFilter = yf }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetSegmentPath() string {
    return "bbe"
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = bbe.Counter
    return leafs
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetBundleName() string { return "cisco_ios_xr" }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetYangName() string { return "bbe" }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) SetParent(parent types.Entity) { bbe.parent = parent }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetParent() types.Entity { return bbe.parent }

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Es
// Errored Seconds information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Es struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetFilter() yfilter.YFilter { return es.YFilter }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) SetFilter(yf yfilter.YFilter) { es.YFilter = yf }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetSegmentPath() string {
    return "es"
}

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = es.Counter
    return leafs
}

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetBundleName() string { return "cisco_ios_xr" }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetYangName() string { return "es" }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) SetParent(parent types.Entity) { es.parent = parent }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetParent() types.Entity { return es.parent }

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses
// Severly Errored Seconds information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetFilter() yfilter.YFilter { return ses.YFilter }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) SetFilter(yf yfilter.YFilter) { ses.YFilter = yf }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetSegmentPath() string {
    return "ses"
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = ses.Counter
    return leafs
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetBundleName() string { return "cisco_ios_xr" }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetYangName() string { return "ses" }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) SetParent(parent types.Entity) { ses.parent = parent }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetParent() types.Entity { return ses.parent }

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas
// Unavailability Seconds information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetFilter() yfilter.YFilter { return uas.YFilter }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) SetFilter(yf yfilter.YFilter) { uas.YFilter = yf }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetSegmentPath() string {
    return "uas"
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = uas.Counter
    return leafs
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetBundleName() string { return "cisco_ios_xr" }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetYangName() string { return "uas" }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) SetParent(parent types.Entity) { uas.parent = parent }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetParent() types.Entity { return uas.parent }

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc
// Failure count information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetFilter() yfilter.YFilter { return fc.YFilter }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) SetFilter(yf yfilter.YFilter) { fc.YFilter = yf }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetSegmentPath() string {
    return "fc"
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = fc.Counter
    return leafs
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetBundleName() string { return "cisco_ios_xr" }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetYangName() string { return "fc" }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) SetParent(parent types.Entity) { fc.parent = parent }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetParent() types.Entity { return fc.parent }

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber
// Background Block Error Rate count information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetFilter() yfilter.YFilter { return bber.YFilter }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) SetFilter(yf yfilter.YFilter) { bber.YFilter = yf }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetSegmentPath() string {
    return "bber"
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = bber.Counter
    return leafs
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetBundleName() string { return "cisco_ios_xr" }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetYangName() string { return "bber" }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) SetParent(parent types.Entity) { bber.parent = parent }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetParent() types.Entity { return bber.parent }

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr
// Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetFilter() yfilter.YFilter { return esr.YFilter }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) SetFilter(yf yfilter.YFilter) { esr.YFilter = yf }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetSegmentPath() string {
    return "esr"
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = esr.Counter
    return leafs
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetBundleName() string { return "cisco_ios_xr" }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetYangName() string { return "esr" }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) SetParent(parent types.Entity) { esr.parent = parent }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetParent() types.Entity { return esr.parent }

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr
// Severly Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetFilter() yfilter.YFilter { return sesr.YFilter }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) SetFilter(yf yfilter.YFilter) { sesr.YFilter = yf }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetGoName(yname string) string {
    if yname == "counter" { return "Counter" }
    return ""
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetSegmentPath() string {
    return "sesr"
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter"] = sesr.Counter
    return leafs
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetBundleName() string { return "cisco_ios_xr" }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetYangName() string { return "sesr" }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) SetParent(parent types.Entity) { sesr.parent = parent }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetParent() types.Entity { return sesr.parent }

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti
// Trail Trace Identifier information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of String. The type is interface{} with range: 0..4294967295.
    TxStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    ExpectedStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    RxStringType interface{}

    // Tx TTI String . The type is string with length: 0..129.
    TxTti interface{}

    // Tx SAPI[0] Field. The type is string with length: 0..5.
    TxSapi0 interface{}

    // Tx SAPI[1-15] Field. The type is string with length: 0..16.
    TxSapi interface{}

    // Tx SAPI Range String. The type is string with length: 0..6.
    TxSapiRange interface{}

    // Tx DAPI[0] Field. The type is string with length: 0..5.
    TxDapi0 interface{}

    // Tx DAPI[1-15] Field. The type is string with length: 0..16.
    TxDapi interface{}

    // Tx DAPI Range String. The type is string with length: 0..6.
    TxDapiRange interface{}

    // Tx Operator Specific Field. The type is string with length: 0..33.
    TxOperSpec interface{}

    // Tx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    TxOperSpecRange interface{}

    // Rx TTI String . The type is string with length: 0..129.
    RxTti interface{}

    // Rx SAPI[0] Field. The type is string with length: 0..5.
    RxSapi0 interface{}

    // Rx SAPI[1-15] Field. The type is string with length: 0..16.
    RxSapi interface{}

    // Rx SAPI Range String. The type is string with length: 0..6.
    RxSapiRange interface{}

    // Rx DAPI[0] Field. The type is string with length: 0..5.
    RxDapi0 interface{}

    // Rx DAPI[1-15] Field. The type is string with length: 0..16.
    RxDapi interface{}

    // Rx DAPI Range String. The type is string with length: 0..6.
    RxDapiRange interface{}

    // Rx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    RxOperSpecRange interface{}

    // Rx Operator Specific Field. The type is string with length: 0..33.
    RxOperSpec interface{}

    // Expected TTI String. The type is string with length: 0..129.
    ExpectedTti interface{}

    // Expected SAPI[0] Field. The type is string with length: 0..5.
    ExpectedSapi0 interface{}

    // Expected SAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedSapi interface{}

    // Expected SAPI Range String. The type is string with length: 0..6.
    ExpSapiRange interface{}

    // Expected DAPI[0] Field. The type is string with length: 0..5.
    ExpectedDapi0 interface{}

    // Expected DAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedDapi interface{}

    // Expected DAPI Range String. The type is string with length: 0..6.
    ExpDapiRange interface{}

    // Expected Operator Specific Field. The type is string with length: 0..33.
    ExpectedOperSpec interface{}

    // Expected Operator Specific Field Range String. The type is string with
    // length: 0..6.
    ExpOperSpecRange interface{}
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetFilter() yfilter.YFilter { return tti.YFilter }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) SetFilter(yf yfilter.YFilter) { tti.YFilter = yf }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetGoName(yname string) string {
    if yname == "tx-string-type" { return "TxStringType" }
    if yname == "expected-string-type" { return "ExpectedStringType" }
    if yname == "rx-string-type" { return "RxStringType" }
    if yname == "tx-tti" { return "TxTti" }
    if yname == "tx-sapi0" { return "TxSapi0" }
    if yname == "tx-sapi" { return "TxSapi" }
    if yname == "tx-sapi-range" { return "TxSapiRange" }
    if yname == "tx-dapi0" { return "TxDapi0" }
    if yname == "tx-dapi" { return "TxDapi" }
    if yname == "tx-dapi-range" { return "TxDapiRange" }
    if yname == "tx-oper-spec" { return "TxOperSpec" }
    if yname == "tx-oper-spec-range" { return "TxOperSpecRange" }
    if yname == "rx-tti" { return "RxTti" }
    if yname == "rx-sapi0" { return "RxSapi0" }
    if yname == "rx-sapi" { return "RxSapi" }
    if yname == "rx-sapi-range" { return "RxSapiRange" }
    if yname == "rx-dapi0" { return "RxDapi0" }
    if yname == "rx-dapi" { return "RxDapi" }
    if yname == "rx-dapi-range" { return "RxDapiRange" }
    if yname == "rx-oper-spec-range" { return "RxOperSpecRange" }
    if yname == "rx-oper-spec" { return "RxOperSpec" }
    if yname == "expected-tti" { return "ExpectedTti" }
    if yname == "expected-sapi0" { return "ExpectedSapi0" }
    if yname == "expected-sapi" { return "ExpectedSapi" }
    if yname == "exp-sapi-range" { return "ExpSapiRange" }
    if yname == "expected-dapi0" { return "ExpectedDapi0" }
    if yname == "expected-dapi" { return "ExpectedDapi" }
    if yname == "exp-dapi-range" { return "ExpDapiRange" }
    if yname == "expected-oper-spec" { return "ExpectedOperSpec" }
    if yname == "exp-oper-spec-range" { return "ExpOperSpecRange" }
    return ""
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetSegmentPath() string {
    return "tti"
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-string-type"] = tti.TxStringType
    leafs["expected-string-type"] = tti.ExpectedStringType
    leafs["rx-string-type"] = tti.RxStringType
    leafs["tx-tti"] = tti.TxTti
    leafs["tx-sapi0"] = tti.TxSapi0
    leafs["tx-sapi"] = tti.TxSapi
    leafs["tx-sapi-range"] = tti.TxSapiRange
    leafs["tx-dapi0"] = tti.TxDapi0
    leafs["tx-dapi"] = tti.TxDapi
    leafs["tx-dapi-range"] = tti.TxDapiRange
    leafs["tx-oper-spec"] = tti.TxOperSpec
    leafs["tx-oper-spec-range"] = tti.TxOperSpecRange
    leafs["rx-tti"] = tti.RxTti
    leafs["rx-sapi0"] = tti.RxSapi0
    leafs["rx-sapi"] = tti.RxSapi
    leafs["rx-sapi-range"] = tti.RxSapiRange
    leafs["rx-dapi0"] = tti.RxDapi0
    leafs["rx-dapi"] = tti.RxDapi
    leafs["rx-dapi-range"] = tti.RxDapiRange
    leafs["rx-oper-spec-range"] = tti.RxOperSpecRange
    leafs["rx-oper-spec"] = tti.RxOperSpec
    leafs["expected-tti"] = tti.ExpectedTti
    leafs["expected-sapi0"] = tti.ExpectedSapi0
    leafs["expected-sapi"] = tti.ExpectedSapi
    leafs["exp-sapi-range"] = tti.ExpSapiRange
    leafs["expected-dapi0"] = tti.ExpectedDapi0
    leafs["expected-dapi"] = tti.ExpectedDapi
    leafs["exp-dapi-range"] = tti.ExpDapiRange
    leafs["expected-oper-spec"] = tti.ExpectedOperSpec
    leafs["exp-oper-spec-range"] = tti.ExpOperSpecRange
    return leafs
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetBundleName() string { return "cisco_ios_xr" }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetYangName() string { return "tti" }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) SetParent(parent types.Entity) { tti.parent = parent }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetParent() types.Entity { return tti.parent }

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetParentYangName() string { return "odu-info" }

// Dwdm_Ports_Port_Info_OpticsInfo
// Optics operational information
type Dwdm_Ports_Port_Info_OpticsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Optics type name. The type is string with length: 0..64.
    OpticsType interface{}

    // Actual transmit clock source. The type is interface{} with range: 0..255.
    ClockSource interface{}

    // Wave Frequency Information for Progressive Frequencies. The type is string
    // with length: 0..64.
    WaveFrequencyProgressiveString interface{}

    // Wavelength Information for Progressive Frequencies. The type is string with
    // length: 0..64.
    WavelengthProgressiveString interface{}

    // True if Progressive Frequency is supported by hw. The type is bool.
    IsWaveFrequencyProgressiveValid interface{}

    // Wavelength Information for Progressive Frequencies. The type is interface{}
    // with range: 0..4294967295.
    WavelengthProgressive interface{}

    // Wavelength band information. The type is interface{} with range:
    // 0..4294967295.
    WaveBand interface{}

    // Current ITU wavelength channel number. The type is interface{} with range:
    // 0..4294967295.
    WaveChannel interface{}

    // wavelenght frequency read from hw in the uint 0 .01nm. The type is
    // interface{} with range: 0..4294967295.
    WaveFrequency interface{}

    // True if hw supported wavelength frequency readback. The type is bool.
    IsWaveFrequencyValid interface{}

    // Owner of current wavelength. The type is DwdmWaveChannelOwner.
    WaveChannelOwner interface{}

    // Wavelength channel set by GMPLS. The type is interface{} with range:
    // 0..65535.
    GmplsSetWaveChannel interface{}

    // Wavelength channel set from configuration. The type is interface{} with
    // range: 0..65535.
    ConfiguredWaveChannel interface{}

    // Wavelength channel default from hardware. The type is interface{} with
    // range: 0..65535.
    DefaultWaveChannel interface{}

    // Transmit power in the unit of 0.01dbm. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPower interface{}

    // Transmit power threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPowerThreshold interface{}

    // Laser current bias value. The type is interface{} with range:
    // -2147483648..2147483647.
    LaserCurrentBias interface{}

    // Laser Current Bias threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    LaserCurrentBiasThreshold interface{}

    // Transponder receive power. The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivePower interface{}

    // TRUE if  Rx LOS thresold configurable. The type is bool.
    IsRxLosThresholdSupported interface{}

    // Rx LOS threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    RxLosThreshold interface{}

    // Transmit  power mininum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    TransmitPowerMin interface{}

    // Transmit power maximum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    TransmitPowerMax interface{}

    // Transmit optical average value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    TransmitPowerAvg interface{}

    // Recieve power mininum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePowerMin interface{}

    // Receive power maximum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePowerMax interface{}

    // Recieve power average value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePowerAvg interface{}

    // Laser bias current minimum value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserBiasCurrentMin interface{}

    // Laser bias current maxinum value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserBiasCurrentMax interface{}

    // Laser bias current average value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserBiasCurrentAvg interface{}

    // Current chromatic dispersion. The type is interface{} with range:
    // -2147483648..2147483647.
    ChromaticDispersion interface{}

    // Current differential group Delay. The type is interface{} with range:
    // -2147483648..2147483647.
    DifferentialGroupDelay interface{}

    // Current polarization mode dispersion. The type is interface{} with range:
    // -2147483648..2147483647.
    PolarizationModeDispersion interface{}

    // Current optical signal to noise ratio. The type is interface{} with range:
    // -2147483648..2147483647.
    SignalToNoiseRatio interface{}

    // Current Polarization Dependent loss. The type is interface{} with range:
    // -2147483648..2147483647.
    PolarizationDependentLoss interface{}

    // Current Polarization change rate. The type is interface{} with range:
    // 0..4294967295.
    PolarizationChangeRate interface{}

    // Current Phase Noise. The type is interface{} with range: 0..4294967295.
    PhaseNoise interface{}

    // Transmit power failure(above/belowe a threshold) count. The type is
    // interface{} with range: 0..4294967295.
    OutputPowerFail interface{}

    // Receive power failure(above/belowe a threshold) count. The type is
    // interface{} with range: 0..4294967295.
    InputPowerFail interface{}
}

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetFilter() yfilter.YFilter { return opticsInfo.YFilter }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) SetFilter(yf yfilter.YFilter) { opticsInfo.YFilter = yf }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetGoName(yname string) string {
    if yname == "optics-type" { return "OpticsType" }
    if yname == "clock-source" { return "ClockSource" }
    if yname == "wave-frequency-progressive-string" { return "WaveFrequencyProgressiveString" }
    if yname == "wavelength-progressive-string" { return "WavelengthProgressiveString" }
    if yname == "is-wave-frequency-progressive-valid" { return "IsWaveFrequencyProgressiveValid" }
    if yname == "wavelength-progressive" { return "WavelengthProgressive" }
    if yname == "wave-band" { return "WaveBand" }
    if yname == "wave-channel" { return "WaveChannel" }
    if yname == "wave-frequency" { return "WaveFrequency" }
    if yname == "is-wave-frequency-valid" { return "IsWaveFrequencyValid" }
    if yname == "wave-channel-owner" { return "WaveChannelOwner" }
    if yname == "gmpls-set-wave-channel" { return "GmplsSetWaveChannel" }
    if yname == "configured-wave-channel" { return "ConfiguredWaveChannel" }
    if yname == "default-wave-channel" { return "DefaultWaveChannel" }
    if yname == "transmit-power" { return "TransmitPower" }
    if yname == "transmit-power-threshold" { return "TransmitPowerThreshold" }
    if yname == "laser-current-bias" { return "LaserCurrentBias" }
    if yname == "laser-current-bias-threshold" { return "LaserCurrentBiasThreshold" }
    if yname == "receive-power" { return "ReceivePower" }
    if yname == "is-rx-los-threshold-supported" { return "IsRxLosThresholdSupported" }
    if yname == "rx-los-threshold" { return "RxLosThreshold" }
    if yname == "transmit-power-min" { return "TransmitPowerMin" }
    if yname == "transmit-power-max" { return "TransmitPowerMax" }
    if yname == "transmit-power-avg" { return "TransmitPowerAvg" }
    if yname == "receive-power-min" { return "ReceivePowerMin" }
    if yname == "receive-power-max" { return "ReceivePowerMax" }
    if yname == "receive-power-avg" { return "ReceivePowerAvg" }
    if yname == "laser-bias-current-min" { return "LaserBiasCurrentMin" }
    if yname == "laser-bias-current-max" { return "LaserBiasCurrentMax" }
    if yname == "laser-bias-current-avg" { return "LaserBiasCurrentAvg" }
    if yname == "chromatic-dispersion" { return "ChromaticDispersion" }
    if yname == "differential-group-delay" { return "DifferentialGroupDelay" }
    if yname == "polarization-mode-dispersion" { return "PolarizationModeDispersion" }
    if yname == "signal-to-noise-ratio" { return "SignalToNoiseRatio" }
    if yname == "polarization-dependent-loss" { return "PolarizationDependentLoss" }
    if yname == "polarization-change-rate" { return "PolarizationChangeRate" }
    if yname == "phase-noise" { return "PhaseNoise" }
    if yname == "output-power-fail" { return "OutputPowerFail" }
    if yname == "input-power-fail" { return "InputPowerFail" }
    return ""
}

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetSegmentPath() string {
    return "optics-info"
}

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["optics-type"] = opticsInfo.OpticsType
    leafs["clock-source"] = opticsInfo.ClockSource
    leafs["wave-frequency-progressive-string"] = opticsInfo.WaveFrequencyProgressiveString
    leafs["wavelength-progressive-string"] = opticsInfo.WavelengthProgressiveString
    leafs["is-wave-frequency-progressive-valid"] = opticsInfo.IsWaveFrequencyProgressiveValid
    leafs["wavelength-progressive"] = opticsInfo.WavelengthProgressive
    leafs["wave-band"] = opticsInfo.WaveBand
    leafs["wave-channel"] = opticsInfo.WaveChannel
    leafs["wave-frequency"] = opticsInfo.WaveFrequency
    leafs["is-wave-frequency-valid"] = opticsInfo.IsWaveFrequencyValid
    leafs["wave-channel-owner"] = opticsInfo.WaveChannelOwner
    leafs["gmpls-set-wave-channel"] = opticsInfo.GmplsSetWaveChannel
    leafs["configured-wave-channel"] = opticsInfo.ConfiguredWaveChannel
    leafs["default-wave-channel"] = opticsInfo.DefaultWaveChannel
    leafs["transmit-power"] = opticsInfo.TransmitPower
    leafs["transmit-power-threshold"] = opticsInfo.TransmitPowerThreshold
    leafs["laser-current-bias"] = opticsInfo.LaserCurrentBias
    leafs["laser-current-bias-threshold"] = opticsInfo.LaserCurrentBiasThreshold
    leafs["receive-power"] = opticsInfo.ReceivePower
    leafs["is-rx-los-threshold-supported"] = opticsInfo.IsRxLosThresholdSupported
    leafs["rx-los-threshold"] = opticsInfo.RxLosThreshold
    leafs["transmit-power-min"] = opticsInfo.TransmitPowerMin
    leafs["transmit-power-max"] = opticsInfo.TransmitPowerMax
    leafs["transmit-power-avg"] = opticsInfo.TransmitPowerAvg
    leafs["receive-power-min"] = opticsInfo.ReceivePowerMin
    leafs["receive-power-max"] = opticsInfo.ReceivePowerMax
    leafs["receive-power-avg"] = opticsInfo.ReceivePowerAvg
    leafs["laser-bias-current-min"] = opticsInfo.LaserBiasCurrentMin
    leafs["laser-bias-current-max"] = opticsInfo.LaserBiasCurrentMax
    leafs["laser-bias-current-avg"] = opticsInfo.LaserBiasCurrentAvg
    leafs["chromatic-dispersion"] = opticsInfo.ChromaticDispersion
    leafs["differential-group-delay"] = opticsInfo.DifferentialGroupDelay
    leafs["polarization-mode-dispersion"] = opticsInfo.PolarizationModeDispersion
    leafs["signal-to-noise-ratio"] = opticsInfo.SignalToNoiseRatio
    leafs["polarization-dependent-loss"] = opticsInfo.PolarizationDependentLoss
    leafs["polarization-change-rate"] = opticsInfo.PolarizationChangeRate
    leafs["phase-noise"] = opticsInfo.PhaseNoise
    leafs["output-power-fail"] = opticsInfo.OutputPowerFail
    leafs["input-power-fail"] = opticsInfo.InputPowerFail
    return leafs
}

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetYangName() string { return "optics-info" }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) SetParent(parent types.Entity) { opticsInfo.parent = parent }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetParent() types.Entity { return opticsInfo.parent }

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetParentYangName() string { return "info" }

// Dwdm_Ports_Port_Info_TdcInfo
// TDC operational information
type Dwdm_Ports_Port_Info_TdcInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE for Valid else Invalid. The type is bool.
    TdcValid interface{}

    // TRUE for Alarm condition else FALSE. The type is bool.
    MajorAlarm interface{}

    // TRUE for MANUAL else AUTO. The type is bool.
    OperationMode interface{}

    // TRUE if TDC Aquiring else Locked. The type is bool.
    TdcStatus interface{}

    // TDC Dispersion Offset. The type is interface{} with range:
    // -2147483648..2147483647.
    DispersionOffset interface{}

    // Reroute BER. The type is interface{} with range: -2147483648..2147483647.
    RerouteBer interface{}

    // TRUE for ENABLED else DISABLED. The type is bool.
    IsRerouteControlEnabled interface{}
}

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetFilter() yfilter.YFilter { return tdcInfo.YFilter }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) SetFilter(yf yfilter.YFilter) { tdcInfo.YFilter = yf }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetGoName(yname string) string {
    if yname == "tdc-valid" { return "TdcValid" }
    if yname == "major-alarm" { return "MajorAlarm" }
    if yname == "operation-mode" { return "OperationMode" }
    if yname == "tdc-status" { return "TdcStatus" }
    if yname == "dispersion-offset" { return "DispersionOffset" }
    if yname == "reroute-ber" { return "RerouteBer" }
    if yname == "is-reroute-control-enabled" { return "IsRerouteControlEnabled" }
    return ""
}

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetSegmentPath() string {
    return "tdc-info"
}

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tdc-valid"] = tdcInfo.TdcValid
    leafs["major-alarm"] = tdcInfo.MajorAlarm
    leafs["operation-mode"] = tdcInfo.OperationMode
    leafs["tdc-status"] = tdcInfo.TdcStatus
    leafs["dispersion-offset"] = tdcInfo.DispersionOffset
    leafs["reroute-ber"] = tdcInfo.RerouteBer
    leafs["is-reroute-control-enabled"] = tdcInfo.IsRerouteControlEnabled
    return leafs
}

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetBundleName() string { return "cisco_ios_xr" }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetYangName() string { return "tdc-info" }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) SetParent(parent types.Entity) { tdcInfo.parent = parent }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetParent() types.Entity { return tdcInfo.parent }

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetParentYangName() string { return "info" }

// Dwdm_Ports_Port_Info_NetworkSrlgInfo
// Network SRLG information
type Dwdm_Ports_Port_Info_NetworkSrlgInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network Srlg. The type is slice of interface{} with range: 0..4294967295.
    NetworkSrlg []interface{}
}

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetFilter() yfilter.YFilter { return networkSrlgInfo.YFilter }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) SetFilter(yf yfilter.YFilter) { networkSrlgInfo.YFilter = yf }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetGoName(yname string) string {
    if yname == "network-srlg" { return "NetworkSrlg" }
    return ""
}

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetSegmentPath() string {
    return "network-srlg-info"
}

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["network-srlg"] = networkSrlgInfo.NetworkSrlg
    return leafs
}

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetBundleName() string { return "cisco_ios_xr" }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetYangName() string { return "network-srlg-info" }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) SetParent(parent types.Entity) { networkSrlgInfo.parent = parent }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetParent() types.Entity { return networkSrlgInfo.parent }

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetParentYangName() string { return "info" }

// Dwdm_Ports_Port_Info_Proactive
// Proactive protection information
type Dwdm_Ports_Port_Info_Proactive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Feature Support. The type is bool.
    ProactiveFeature interface{}

    // Proactive Mode. The type is G709ppfsmMode.
    ProactiveMode interface{}

    // Proactive FSM State. The type is G709ppfsmState.
    ProactiveFsmState interface{}

    // Proactive FSM IF State. The type is G709ppintfState.
    ProactiveFsmIfState interface{}

    // TAS State. The type is DwdmtasState.
    TasState interface{}

    // Trigger threshold coefficient. The type is interface{} with range: 0..255.
    TrigThreshCoeff interface{}

    // Trigger threshold power. The type is interface{} with range: 0..255.
    TrigThreshPower interface{}

    // Revert threshold coefficient. The type is interface{} with range: 0..255.
    RvrtThreshCoeff interface{}

    // Revert threshold power. The type is interface{} with range: 0..255.
    RvrtThreshPower interface{}

    // Default Trigger threshold coefficient. The type is interface{} with range:
    // 0..255.
    DefaultTrigThreshCoeff interface{}

    // Default Trigger threshold power. The type is interface{} with range:
    // 0..255.
    DefaultTrigThreshPower interface{}

    // Default Revert threshold coefficient. The type is interface{} with range:
    // 0..255.
    DefaultRvrtThreshCoeff interface{}

    // Default Revert threshold power. The type is interface{} with range: 0..255.
    DefaultRvrtThreshPower interface{}

    // Required Trigger Samples. The type is interface{} with range: 0..255.
    TrigSamples interface{}

    // Required Revert Samples. The type is interface{} with range: 0..255.
    RvrtSamples interface{}

    // Trigger Integration window. The type is interface{} with range:
    // 0..4294967295.
    TriggerWindow interface{}

    // Revert Integration Window. The type is interface{} with range:
    // 0..4294967295.
    RevertWindow interface{}

    // Protection Trigger State. The type is bool.
    ProtectionTrigger interface{}

    // Proactive Interface Triffer. The type is bool.
    InterfaceTrigger interface{}

    // Transmitted APS Byte. The type is interface{} with range: 0..255.
    TxAps interface{}

    // Tx APS Description. The type is G709apsByte.
    TxApsDescr interface{}

    // Received APS byte. The type is interface{} with range: 0..255.
    RxAps interface{}

    // Rx APS Description. The type is G709apsByte.
    RxApsDescr interface{}

    // AlarmState. The type is bool.
    AlarmState interface{}

    // Trigger EC Cnt. The type is interface{} with range: 0..4294967295.
    TrigEcCnt interface{}

    // Revert EC Cnt. The type is interface{} with range: 0..4294967295.
    RvrtEcCnt interface{}

    // Prefec Trigger Thresh Crossed. The type is bool.
    PrefecThreshCrossed interface{}
}

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetFilter() yfilter.YFilter { return proactive.YFilter }

func (proactive *Dwdm_Ports_Port_Info_Proactive) SetFilter(yf yfilter.YFilter) { proactive.YFilter = yf }

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetGoName(yname string) string {
    if yname == "proactive-feature" { return "ProactiveFeature" }
    if yname == "proactive-mode" { return "ProactiveMode" }
    if yname == "proactive-fsm-state" { return "ProactiveFsmState" }
    if yname == "proactive-fsm-if-state" { return "ProactiveFsmIfState" }
    if yname == "tas-state" { return "TasState" }
    if yname == "trig-thresh-coeff" { return "TrigThreshCoeff" }
    if yname == "trig-thresh-power" { return "TrigThreshPower" }
    if yname == "rvrt-thresh-coeff" { return "RvrtThreshCoeff" }
    if yname == "rvrt-thresh-power" { return "RvrtThreshPower" }
    if yname == "default-trig-thresh-coeff" { return "DefaultTrigThreshCoeff" }
    if yname == "default-trig-thresh-power" { return "DefaultTrigThreshPower" }
    if yname == "default-rvrt-thresh-coeff" { return "DefaultRvrtThreshCoeff" }
    if yname == "default-rvrt-thresh-power" { return "DefaultRvrtThreshPower" }
    if yname == "trig-samples" { return "TrigSamples" }
    if yname == "rvrt-samples" { return "RvrtSamples" }
    if yname == "trigger-window" { return "TriggerWindow" }
    if yname == "revert-window" { return "RevertWindow" }
    if yname == "protection-trigger" { return "ProtectionTrigger" }
    if yname == "interface-trigger" { return "InterfaceTrigger" }
    if yname == "tx-aps" { return "TxAps" }
    if yname == "tx-aps-descr" { return "TxApsDescr" }
    if yname == "rx-aps" { return "RxAps" }
    if yname == "rx-aps-descr" { return "RxApsDescr" }
    if yname == "alarm-state" { return "AlarmState" }
    if yname == "trig-ec-cnt" { return "TrigEcCnt" }
    if yname == "rvrt-ec-cnt" { return "RvrtEcCnt" }
    if yname == "prefec-thresh-crossed" { return "PrefecThreshCrossed" }
    return ""
}

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetSegmentPath() string {
    return "proactive"
}

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proactive-feature"] = proactive.ProactiveFeature
    leafs["proactive-mode"] = proactive.ProactiveMode
    leafs["proactive-fsm-state"] = proactive.ProactiveFsmState
    leafs["proactive-fsm-if-state"] = proactive.ProactiveFsmIfState
    leafs["tas-state"] = proactive.TasState
    leafs["trig-thresh-coeff"] = proactive.TrigThreshCoeff
    leafs["trig-thresh-power"] = proactive.TrigThreshPower
    leafs["rvrt-thresh-coeff"] = proactive.RvrtThreshCoeff
    leafs["rvrt-thresh-power"] = proactive.RvrtThreshPower
    leafs["default-trig-thresh-coeff"] = proactive.DefaultTrigThreshCoeff
    leafs["default-trig-thresh-power"] = proactive.DefaultTrigThreshPower
    leafs["default-rvrt-thresh-coeff"] = proactive.DefaultRvrtThreshCoeff
    leafs["default-rvrt-thresh-power"] = proactive.DefaultRvrtThreshPower
    leafs["trig-samples"] = proactive.TrigSamples
    leafs["rvrt-samples"] = proactive.RvrtSamples
    leafs["trigger-window"] = proactive.TriggerWindow
    leafs["revert-window"] = proactive.RevertWindow
    leafs["protection-trigger"] = proactive.ProtectionTrigger
    leafs["interface-trigger"] = proactive.InterfaceTrigger
    leafs["tx-aps"] = proactive.TxAps
    leafs["tx-aps-descr"] = proactive.TxApsDescr
    leafs["rx-aps"] = proactive.RxAps
    leafs["rx-aps-descr"] = proactive.RxApsDescr
    leafs["alarm-state"] = proactive.AlarmState
    leafs["trig-ec-cnt"] = proactive.TrigEcCnt
    leafs["rvrt-ec-cnt"] = proactive.RvrtEcCnt
    leafs["prefec-thresh-crossed"] = proactive.PrefecThreshCrossed
    return leafs
}

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetBundleName() string { return "cisco_ios_xr" }

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetYangName() string { return "proactive" }

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proactive *Dwdm_Ports_Port_Info_Proactive) SetParent(parent types.Entity) { proactive.parent = parent }

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetParent() types.Entity { return proactive.parent }

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetParentYangName() string { return "info" }

// Dwdm_Ports_Port_Info_SignalLog
// Signal log information
type Dwdm_Ports_Port_Info_SignalLog struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 'true' if signal log is enabled 'false' otherwise. The type is bool.
    IsLogEnabled interface{}

    // Log file name . The type is string with length: 0..64.
    LogFilename interface{}
}

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetFilter() yfilter.YFilter { return signalLog.YFilter }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) SetFilter(yf yfilter.YFilter) { signalLog.YFilter = yf }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetGoName(yname string) string {
    if yname == "is-log-enabled" { return "IsLogEnabled" }
    if yname == "log-filename" { return "LogFilename" }
    return ""
}

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetSegmentPath() string {
    return "signal-log"
}

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-log-enabled"] = signalLog.IsLogEnabled
    leafs["log-filename"] = signalLog.LogFilename
    return leafs
}

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetBundleName() string { return "cisco_ios_xr" }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetYangName() string { return "signal-log" }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) SetParent(parent types.Entity) { signalLog.parent = parent }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetParent() types.Entity { return signalLog.parent }

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetParentYangName() string { return "info" }

// Vtxp
// vtxp
type Vtxp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DWDM operational data.
    DwdmVtxp Vtxp_DwdmVtxp
}

func (vtxp *Vtxp) GetFilter() yfilter.YFilter { return vtxp.YFilter }

func (vtxp *Vtxp) SetFilter(yf yfilter.YFilter) { vtxp.YFilter = yf }

func (vtxp *Vtxp) GetGoName(yname string) string {
    if yname == "dwdm-vtxp" { return "DwdmVtxp" }
    return ""
}

func (vtxp *Vtxp) GetSegmentPath() string {
    return "Cisco-IOS-XR-dwdm-ui-oper:vtxp"
}

func (vtxp *Vtxp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dwdm-vtxp" {
        return &vtxp.DwdmVtxp
    }
    return nil
}

func (vtxp *Vtxp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dwdm-vtxp"] = &vtxp.DwdmVtxp
    return children
}

func (vtxp *Vtxp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtxp *Vtxp) GetBundleName() string { return "cisco_ios_xr" }

func (vtxp *Vtxp) GetYangName() string { return "vtxp" }

func (vtxp *Vtxp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vtxp *Vtxp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vtxp *Vtxp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vtxp *Vtxp) SetParent(parent types.Entity) { vtxp.parent = parent }

func (vtxp *Vtxp) GetParent() types.Entity { return vtxp.parent }

func (vtxp *Vtxp) GetParentYangName() string { return "Cisco-IOS-XR-dwdm-ui-oper" }

// Vtxp_DwdmVtxp
// DWDM operational data
type Vtxp_DwdmVtxp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All DWDM Port operational data.
    PortVtxps Vtxp_DwdmVtxp_PortVtxps
}

func (dwdmVtxp *Vtxp_DwdmVtxp) GetFilter() yfilter.YFilter { return dwdmVtxp.YFilter }

func (dwdmVtxp *Vtxp_DwdmVtxp) SetFilter(yf yfilter.YFilter) { dwdmVtxp.YFilter = yf }

func (dwdmVtxp *Vtxp_DwdmVtxp) GetGoName(yname string) string {
    if yname == "port-vtxps" { return "PortVtxps" }
    return ""
}

func (dwdmVtxp *Vtxp_DwdmVtxp) GetSegmentPath() string {
    return "dwdm-vtxp"
}

func (dwdmVtxp *Vtxp_DwdmVtxp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-vtxps" {
        return &dwdmVtxp.PortVtxps
    }
    return nil
}

func (dwdmVtxp *Vtxp_DwdmVtxp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-vtxps"] = &dwdmVtxp.PortVtxps
    return children
}

func (dwdmVtxp *Vtxp_DwdmVtxp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dwdmVtxp *Vtxp_DwdmVtxp) GetBundleName() string { return "cisco_ios_xr" }

func (dwdmVtxp *Vtxp_DwdmVtxp) GetYangName() string { return "dwdm-vtxp" }

func (dwdmVtxp *Vtxp_DwdmVtxp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dwdmVtxp *Vtxp_DwdmVtxp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dwdmVtxp *Vtxp_DwdmVtxp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dwdmVtxp *Vtxp_DwdmVtxp) SetParent(parent types.Entity) { dwdmVtxp.parent = parent }

func (dwdmVtxp *Vtxp_DwdmVtxp) GetParent() types.Entity { return dwdmVtxp.parent }

func (dwdmVtxp *Vtxp_DwdmVtxp) GetParentYangName() string { return "vtxp" }

// Vtxp_DwdmVtxp_PortVtxps
// All DWDM Port operational data
type Vtxp_DwdmVtxp_PortVtxps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DWDM Port operational data. The type is slice of
    // Vtxp_DwdmVtxp_PortVtxps_PortVtxp.
    PortVtxp []Vtxp_DwdmVtxp_PortVtxps_PortVtxp
}

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetFilter() yfilter.YFilter { return portVtxps.YFilter }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) SetFilter(yf yfilter.YFilter) { portVtxps.YFilter = yf }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetGoName(yname string) string {
    if yname == "port-vtxp" { return "PortVtxp" }
    return ""
}

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetSegmentPath() string {
    return "port-vtxps"
}

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-vtxp" {
        for _, c := range portVtxps.PortVtxp {
            if portVtxps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vtxp_DwdmVtxp_PortVtxps_PortVtxp{}
        portVtxps.PortVtxp = append(portVtxps.PortVtxp, child)
        return &portVtxps.PortVtxp[len(portVtxps.PortVtxp)-1]
    }
    return nil
}

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portVtxps.PortVtxp {
        children[portVtxps.PortVtxp[i].GetSegmentPath()] = &portVtxps.PortVtxp[i]
    }
    return children
}

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetBundleName() string { return "cisco_ios_xr" }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetYangName() string { return "port-vtxps" }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) SetParent(parent types.Entity) { portVtxps.parent = parent }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetParent() types.Entity { return portVtxps.parent }

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetParentYangName() string { return "dwdm-vtxp" }

// Vtxp_DwdmVtxp_PortVtxps_PortVtxp
// DWDM Port operational data
type Vtxp_DwdmVtxp_PortVtxps_PortVtxp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // DWDM port operational data.
    Info Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info
}

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetFilter() yfilter.YFilter { return portVtxp.YFilter }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) SetFilter(yf yfilter.YFilter) { portVtxp.YFilter = yf }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "info" { return "Info" }
    return ""
}

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetSegmentPath() string {
    return "port-vtxp" + "[name='" + fmt.Sprintf("%v", portVtxp.Name) + "']"
}

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "info" {
        return &portVtxp.Info
    }
    return nil
}

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["info"] = &portVtxp.Info
    return children
}

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = portVtxp.Name
    return leafs
}

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetBundleName() string { return "cisco_ios_xr" }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetYangName() string { return "port-vtxp" }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) SetParent(parent types.Entity) { portVtxp.parent = parent }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetParent() types.Entity { return portVtxp.parent }

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetParentYangName() string { return "port-vtxps" }

// Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info
// DWDM port operational data
type Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is VTXP attribute enabled. The type is bool.
    VtxpEnable interface{}
}

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetFilter() yfilter.YFilter { return info.YFilter }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) SetFilter(yf yfilter.YFilter) { info.YFilter = yf }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetGoName(yname string) string {
    if yname == "vtxp-enable" { return "VtxpEnable" }
    return ""
}

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetSegmentPath() string {
    return "info"
}

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vtxp-enable"] = info.VtxpEnable
    return leafs
}

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetBundleName() string { return "cisco_ios_xr" }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetYangName() string { return "info" }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) SetParent(parent types.Entity) { info.parent = parent }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetParent() types.Entity { return info.parent }

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetParentYangName() string { return "port-vtxp" }

