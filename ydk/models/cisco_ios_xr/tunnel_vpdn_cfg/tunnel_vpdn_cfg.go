// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-vpdn package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vpdn: VPDN configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_vpdn_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_vpdn_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-vpdn-cfg vpdn}", reflect.TypeOf(Vpdn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-vpdn-cfg:vpdn", reflect.TypeOf(Vpdn{}))
}

// DfBit represents Df bit
type DfBit string

const (
    // Clear df bit
    DfBit_clear DfBit = "clear"

    // Reflect df bit from inner ip header
    DfBit_reflect DfBit = "reflect"

    // Set df bit
    DfBit_set DfBit = "set"
)

// Option represents Option
type Option string

const (
    // Log VPDN events locally
    Option_local Option = "local"

    // Log VPDN user events
    Option_user Option = "user"

    // Log VPDN dead cache
    Option_dead_cache Option = "dead-cache"

    // Log VPDN tunnel drops
    Option_tunnel_drop Option = "tunnel-drop"
)

// Vpdn
// VPDN configuration
type Vpdn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum simultaneous VPDN sessions. The type is interface{} with range:
    // 1..131072.
    SessionLimit interface{}

    // Enable VPDN configuration. Deletion of this object also causes deletion of
    // all associated objects under VPDN. The type is interface{}.
    Enable interface{}

    // New session no longer allowed. The type is interface{}.
    SoftShut interface{}

    // VPDN history logging.
    History Vpdn_History

    // Enable VPDN redundancy.
    Redundancy Vpdn_Redundancy

    // VPDN Local radius process configuration.
    Local Vpdn_Local

    // Table of Template.
    Templates Vpdn_Templates

    // Options to apply on calling station ID.
    CallerId Vpdn_CallerId

    // Table of VPDNgroup.
    VpdNgroups Vpdn_VpdNgroups

    // Table of Logging.
    Loggings Vpdn_Loggings

    // L2TPv2 protocol commands.
    L2Tp Vpdn_L2Tp
}

func (vpdn *Vpdn) GetFilter() yfilter.YFilter { return vpdn.YFilter }

func (vpdn *Vpdn) SetFilter(yf yfilter.YFilter) { vpdn.YFilter = yf }

func (vpdn *Vpdn) GetGoName(yname string) string {
    if yname == "session-limit" { return "SessionLimit" }
    if yname == "enable" { return "Enable" }
    if yname == "soft-shut" { return "SoftShut" }
    if yname == "history" { return "History" }
    if yname == "redundancy" { return "Redundancy" }
    if yname == "local" { return "Local" }
    if yname == "templates" { return "Templates" }
    if yname == "caller-id" { return "CallerId" }
    if yname == "vpd-ngroups" { return "VpdNgroups" }
    if yname == "loggings" { return "Loggings" }
    if yname == "l2tp" { return "L2Tp" }
    return ""
}

func (vpdn *Vpdn) GetSegmentPath() string {
    return "Cisco-IOS-XR-tunnel-vpdn-cfg:vpdn"
}

func (vpdn *Vpdn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "history" {
        return &vpdn.History
    }
    if childYangName == "redundancy" {
        return &vpdn.Redundancy
    }
    if childYangName == "local" {
        return &vpdn.Local
    }
    if childYangName == "templates" {
        return &vpdn.Templates
    }
    if childYangName == "caller-id" {
        return &vpdn.CallerId
    }
    if childYangName == "vpd-ngroups" {
        return &vpdn.VpdNgroups
    }
    if childYangName == "loggings" {
        return &vpdn.Loggings
    }
    if childYangName == "l2tp" {
        return &vpdn.L2Tp
    }
    return nil
}

func (vpdn *Vpdn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["history"] = &vpdn.History
    children["redundancy"] = &vpdn.Redundancy
    children["local"] = &vpdn.Local
    children["templates"] = &vpdn.Templates
    children["caller-id"] = &vpdn.CallerId
    children["vpd-ngroups"] = &vpdn.VpdNgroups
    children["loggings"] = &vpdn.Loggings
    children["l2tp"] = &vpdn.L2Tp
    return children
}

func (vpdn *Vpdn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-limit"] = vpdn.SessionLimit
    leafs["enable"] = vpdn.Enable
    leafs["soft-shut"] = vpdn.SoftShut
    return leafs
}

func (vpdn *Vpdn) GetBundleName() string { return "cisco_ios_xr" }

func (vpdn *Vpdn) GetYangName() string { return "vpdn" }

func (vpdn *Vpdn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdn *Vpdn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdn *Vpdn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdn *Vpdn) SetParent(parent types.Entity) { vpdn.parent = parent }

func (vpdn *Vpdn) GetParent() types.Entity { return vpdn.parent }

func (vpdn *Vpdn) GetParentYangName() string { return "Cisco-IOS-XR-tunnel-vpdn-cfg" }

// Vpdn_History
// VPDN history logging
type Vpdn_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // User failure. The type is interface{}.
    Failure interface{}
}

func (history *Vpdn_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Vpdn_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Vpdn_History) GetGoName(yname string) string {
    if yname == "failure" { return "Failure" }
    return ""
}

func (history *Vpdn_History) GetSegmentPath() string {
    return "history"
}

func (history *Vpdn_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (history *Vpdn_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (history *Vpdn_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["failure"] = history.Failure
    return leafs
}

func (history *Vpdn_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Vpdn_History) GetYangName() string { return "history" }

func (history *Vpdn_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Vpdn_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Vpdn_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Vpdn_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Vpdn_History) GetParent() types.Entity { return history.parent }

func (history *Vpdn_History) GetParentYangName() string { return "vpdn" }

// Vpdn_Redundancy
// Enable VPDN redundancy
type Vpdn_Redundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Enable VPDN redundancy. Deletion of this object also causes deletion
    // of all associated objects under Redundancy. The type is interface{}.
    Enable interface{}

    // Process crash configuration.
    ProcessFailures Vpdn_Redundancy_ProcessFailures
}

func (redundancy *Vpdn_Redundancy) GetFilter() yfilter.YFilter { return redundancy.YFilter }

func (redundancy *Vpdn_Redundancy) SetFilter(yf yfilter.YFilter) { redundancy.YFilter = yf }

func (redundancy *Vpdn_Redundancy) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "process-failures" { return "ProcessFailures" }
    return ""
}

func (redundancy *Vpdn_Redundancy) GetSegmentPath() string {
    return "redundancy"
}

func (redundancy *Vpdn_Redundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-failures" {
        return &redundancy.ProcessFailures
    }
    return nil
}

func (redundancy *Vpdn_Redundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process-failures"] = &redundancy.ProcessFailures
    return children
}

func (redundancy *Vpdn_Redundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = redundancy.Enable
    return leafs
}

func (redundancy *Vpdn_Redundancy) GetBundleName() string { return "cisco_ios_xr" }

func (redundancy *Vpdn_Redundancy) GetYangName() string { return "redundancy" }

func (redundancy *Vpdn_Redundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redundancy *Vpdn_Redundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redundancy *Vpdn_Redundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redundancy *Vpdn_Redundancy) SetParent(parent types.Entity) { redundancy.parent = parent }

func (redundancy *Vpdn_Redundancy) GetParent() types.Entity { return redundancy.parent }

func (redundancy *Vpdn_Redundancy) GetParentYangName() string { return "vpdn" }

// Vpdn_Redundancy_ProcessFailures
// Process crash configuration
type Vpdn_Redundancy_ProcessFailures struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Force a switchover if the process crashes. The type is interface{}.
    Switchover interface{}
}

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetFilter() yfilter.YFilter { return processFailures.YFilter }

func (processFailures *Vpdn_Redundancy_ProcessFailures) SetFilter(yf yfilter.YFilter) { processFailures.YFilter = yf }

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetGoName(yname string) string {
    if yname == "switchover" { return "Switchover" }
    return ""
}

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetSegmentPath() string {
    return "process-failures"
}

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["switchover"] = processFailures.Switchover
    return leafs
}

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetBundleName() string { return "cisco_ios_xr" }

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetYangName() string { return "process-failures" }

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processFailures *Vpdn_Redundancy_ProcessFailures) SetParent(parent types.Entity) { processFailures.parent = parent }

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetParent() types.Entity { return processFailures.parent }

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetParentYangName() string { return "redundancy" }

// Vpdn_Local
// VPDN Local radius process configuration
type Vpdn_Local struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // secret password. The type is string with length: 1..32.
    SecretText interface{}

    // local path of the saved profile. The type is string with length: 1..64.
    Path interface{}

    // Set constant integer. The type is interface{}.
    CacheDisabled interface{}

    // port value. The type is interface{} with range: 1..65535.
    Port interface{}
}

func (local *Vpdn_Local) GetFilter() yfilter.YFilter { return local.YFilter }

func (local *Vpdn_Local) SetFilter(yf yfilter.YFilter) { local.YFilter = yf }

func (local *Vpdn_Local) GetGoName(yname string) string {
    if yname == "secret-text" { return "SecretText" }
    if yname == "path" { return "Path" }
    if yname == "cache-disabled" { return "CacheDisabled" }
    if yname == "port" { return "Port" }
    return ""
}

func (local *Vpdn_Local) GetSegmentPath() string {
    return "local"
}

func (local *Vpdn_Local) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (local *Vpdn_Local) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (local *Vpdn_Local) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["secret-text"] = local.SecretText
    leafs["path"] = local.Path
    leafs["cache-disabled"] = local.CacheDisabled
    leafs["port"] = local.Port
    return leafs
}

func (local *Vpdn_Local) GetBundleName() string { return "cisco_ios_xr" }

func (local *Vpdn_Local) GetYangName() string { return "local" }

func (local *Vpdn_Local) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (local *Vpdn_Local) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (local *Vpdn_Local) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (local *Vpdn_Local) SetParent(parent types.Entity) { local.parent = parent }

func (local *Vpdn_Local) GetParent() types.Entity { return local.parent }

func (local *Vpdn_Local) GetParentYangName() string { return "vpdn" }

// Vpdn_Templates
// Table of Template
type Vpdn_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPDN template configuration. The type is slice of Vpdn_Templates_Template.
    Template []Vpdn_Templates_Template
}

func (templates *Vpdn_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *Vpdn_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *Vpdn_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *Vpdn_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *Vpdn_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vpdn_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *Vpdn_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *Vpdn_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *Vpdn_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *Vpdn_Templates) GetYangName() string { return "templates" }

func (templates *Vpdn_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *Vpdn_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *Vpdn_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *Vpdn_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *Vpdn_Templates) GetParent() types.Entity { return templates.parent }

func (templates *Vpdn_Templates) GetParentYangName() string { return "vpdn" }

// Vpdn_Templates_Template
// VPDN template configuration
type Vpdn_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VPDN template name. The type is string with
    // length: 1..63.
    TemplateName interface{}

    // To support NAS-Port format e in Cisco AVP 100. The type is interface{}.
    CiscoAvp100FormatEEnable interface{}

    // Up to 100 characters describing this VPDN template. The type is string with
    // length: 1..100.
    Description interface{}

    // L2TP class  command. The type is string with length: 1..79.
    L2TpClass interface{}

    // Forward DSL Line Info attributes. The type is interface{}.
    DslLineForwarding interface{}

    // Options to apply on calling station id.
    CallerId Vpdn_Templates_Template_CallerId

    // VPN ID/VRF name.
    Vpn Vpdn_Templates_Template_Vpn

    // L2TP tunnel commands.
    Tunnel Vpdn_Templates_Template_Tunnel

    // Set IP TOS value.
    Ip Vpdn_Templates_Template_Ip

    // IPv4 settings for tunnel.
    Ipv4 Vpdn_Templates_Template_Ipv4
}

func (template *Vpdn_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *Vpdn_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *Vpdn_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "cisco-avp100-format-e-enable" { return "CiscoAvp100FormatEEnable" }
    if yname == "description" { return "Description" }
    if yname == "l2tp-class" { return "L2TpClass" }
    if yname == "dsl-line-forwarding" { return "DslLineForwarding" }
    if yname == "caller-id" { return "CallerId" }
    if yname == "vpn" { return "Vpn" }
    if yname == "tunnel" { return "Tunnel" }
    if yname == "ip" { return "Ip" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (template *Vpdn_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *Vpdn_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "caller-id" {
        return &template.CallerId
    }
    if childYangName == "vpn" {
        return &template.Vpn
    }
    if childYangName == "tunnel" {
        return &template.Tunnel
    }
    if childYangName == "ip" {
        return &template.Ip
    }
    if childYangName == "ipv4" {
        return &template.Ipv4
    }
    return nil
}

func (template *Vpdn_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["caller-id"] = &template.CallerId
    children["vpn"] = &template.Vpn
    children["tunnel"] = &template.Tunnel
    children["ip"] = &template.Ip
    children["ipv4"] = &template.Ipv4
    return children
}

func (template *Vpdn_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["cisco-avp100-format-e-enable"] = template.CiscoAvp100FormatEEnable
    leafs["description"] = template.Description
    leafs["l2tp-class"] = template.L2TpClass
    leafs["dsl-line-forwarding"] = template.DslLineForwarding
    return leafs
}

func (template *Vpdn_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *Vpdn_Templates_Template) GetYangName() string { return "template" }

func (template *Vpdn_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *Vpdn_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *Vpdn_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *Vpdn_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *Vpdn_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *Vpdn_Templates_Template) GetParentYangName() string { return "templates" }

// Vpdn_Templates_Template_CallerId
// Options to apply on calling station id
type Vpdn_Templates_Template_CallerId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mask characters by method. The type is string with length: 1..63.
    Mask interface{}
}

func (callerId *Vpdn_Templates_Template_CallerId) GetFilter() yfilter.YFilter { return callerId.YFilter }

func (callerId *Vpdn_Templates_Template_CallerId) SetFilter(yf yfilter.YFilter) { callerId.YFilter = yf }

func (callerId *Vpdn_Templates_Template_CallerId) GetGoName(yname string) string {
    if yname == "mask" { return "Mask" }
    return ""
}

func (callerId *Vpdn_Templates_Template_CallerId) GetSegmentPath() string {
    return "caller-id"
}

func (callerId *Vpdn_Templates_Template_CallerId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (callerId *Vpdn_Templates_Template_CallerId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (callerId *Vpdn_Templates_Template_CallerId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mask"] = callerId.Mask
    return leafs
}

func (callerId *Vpdn_Templates_Template_CallerId) GetBundleName() string { return "cisco_ios_xr" }

func (callerId *Vpdn_Templates_Template_CallerId) GetYangName() string { return "caller-id" }

func (callerId *Vpdn_Templates_Template_CallerId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (callerId *Vpdn_Templates_Template_CallerId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (callerId *Vpdn_Templates_Template_CallerId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (callerId *Vpdn_Templates_Template_CallerId) SetParent(parent types.Entity) { callerId.parent = parent }

func (callerId *Vpdn_Templates_Template_CallerId) GetParent() types.Entity { return callerId.parent }

func (callerId *Vpdn_Templates_Template_CallerId) GetParentYangName() string { return "template" }

// Vpdn_Templates_Template_Vpn
// VPN ID/VRF name
type Vpdn_Templates_Template_Vpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string with length: 1..32.
    Vrf interface{}

    // VPN ID.
    Id Vpdn_Templates_Template_Vpn_Id
}

func (vpn *Vpdn_Templates_Template_Vpn) GetFilter() yfilter.YFilter { return vpn.YFilter }

func (vpn *Vpdn_Templates_Template_Vpn) SetFilter(yf yfilter.YFilter) { vpn.YFilter = yf }

func (vpn *Vpdn_Templates_Template_Vpn) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    if yname == "id" { return "Id" }
    return ""
}

func (vpn *Vpdn_Templates_Template_Vpn) GetSegmentPath() string {
    return "vpn"
}

func (vpn *Vpdn_Templates_Template_Vpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "id" {
        return &vpn.Id
    }
    return nil
}

func (vpn *Vpdn_Templates_Template_Vpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["id"] = &vpn.Id
    return children
}

func (vpn *Vpdn_Templates_Template_Vpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf"] = vpn.Vrf
    return leafs
}

func (vpn *Vpdn_Templates_Template_Vpn) GetBundleName() string { return "cisco_ios_xr" }

func (vpn *Vpdn_Templates_Template_Vpn) GetYangName() string { return "vpn" }

func (vpn *Vpdn_Templates_Template_Vpn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpn *Vpdn_Templates_Template_Vpn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpn *Vpdn_Templates_Template_Vpn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpn *Vpdn_Templates_Template_Vpn) SetParent(parent types.Entity) { vpn.parent = parent }

func (vpn *Vpdn_Templates_Template_Vpn) GetParent() types.Entity { return vpn.parent }

func (vpn *Vpdn_Templates_Template_Vpn) GetParentYangName() string { return "template" }

// Vpdn_Templates_Template_Vpn_Id
// VPN ID
type Vpdn_Templates_Template_Vpn_Id struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPN ID, (OUI:VPN-Index) format(hex), 3 bytes OUI Part. The type is string
    // with pattern: [0-9a-fA-F]{1,8}.
    Oui interface{}

    // VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part. The type is
    // string with pattern: [0-9a-fA-F]{1,8}.
    Index interface{}
}

func (id *Vpdn_Templates_Template_Vpn_Id) GetFilter() yfilter.YFilter { return id.YFilter }

func (id *Vpdn_Templates_Template_Vpn_Id) SetFilter(yf yfilter.YFilter) { id.YFilter = yf }

func (id *Vpdn_Templates_Template_Vpn_Id) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "index" { return "Index" }
    return ""
}

func (id *Vpdn_Templates_Template_Vpn_Id) GetSegmentPath() string {
    return "id"
}

func (id *Vpdn_Templates_Template_Vpn_Id) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (id *Vpdn_Templates_Template_Vpn_Id) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (id *Vpdn_Templates_Template_Vpn_Id) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = id.Oui
    leafs["index"] = id.Index
    return leafs
}

func (id *Vpdn_Templates_Template_Vpn_Id) GetBundleName() string { return "cisco_ios_xr" }

func (id *Vpdn_Templates_Template_Vpn_Id) GetYangName() string { return "id" }

func (id *Vpdn_Templates_Template_Vpn_Id) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (id *Vpdn_Templates_Template_Vpn_Id) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (id *Vpdn_Templates_Template_Vpn_Id) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (id *Vpdn_Templates_Template_Vpn_Id) SetParent(parent types.Entity) { id.parent = parent }

func (id *Vpdn_Templates_Template_Vpn_Id) GetParent() types.Entity { return id.parent }

func (id *Vpdn_Templates_Template_Vpn_Id) GetParentYangName() string { return "vpn" }

// Vpdn_Templates_Template_Tunnel
// L2TP tunnel commands
type Vpdn_Templates_Template_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Busy time out value in seconds. The type is interface{} with range:
    // 60..65535. Units are second.
    BusyTimeout interface{}
}

func (tunnel *Vpdn_Templates_Template_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *Vpdn_Templates_Template_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *Vpdn_Templates_Template_Tunnel) GetGoName(yname string) string {
    if yname == "busy-timeout" { return "BusyTimeout" }
    return ""
}

func (tunnel *Vpdn_Templates_Template_Tunnel) GetSegmentPath() string {
    return "tunnel"
}

func (tunnel *Vpdn_Templates_Template_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnel *Vpdn_Templates_Template_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnel *Vpdn_Templates_Template_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["busy-timeout"] = tunnel.BusyTimeout
    return leafs
}

func (tunnel *Vpdn_Templates_Template_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *Vpdn_Templates_Template_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *Vpdn_Templates_Template_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *Vpdn_Templates_Template_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *Vpdn_Templates_Template_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *Vpdn_Templates_Template_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *Vpdn_Templates_Template_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *Vpdn_Templates_Template_Tunnel) GetParentYangName() string { return "template" }

// Vpdn_Templates_Template_Ip
// Set IP TOS value
type Vpdn_Templates_Template_Ip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set constant integer. The type is interface{} with range:
    // -2147483648..2147483647.
    Tos interface{}
}

func (ip *Vpdn_Templates_Template_Ip) GetFilter() yfilter.YFilter { return ip.YFilter }

func (ip *Vpdn_Templates_Template_Ip) SetFilter(yf yfilter.YFilter) { ip.YFilter = yf }

func (ip *Vpdn_Templates_Template_Ip) GetGoName(yname string) string {
    if yname == "tos" { return "Tos" }
    return ""
}

func (ip *Vpdn_Templates_Template_Ip) GetSegmentPath() string {
    return "ip"
}

func (ip *Vpdn_Templates_Template_Ip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ip *Vpdn_Templates_Template_Ip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ip *Vpdn_Templates_Template_Ip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tos"] = ip.Tos
    return leafs
}

func (ip *Vpdn_Templates_Template_Ip) GetBundleName() string { return "cisco_ios_xr" }

func (ip *Vpdn_Templates_Template_Ip) GetYangName() string { return "ip" }

func (ip *Vpdn_Templates_Template_Ip) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ip *Vpdn_Templates_Template_Ip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ip *Vpdn_Templates_Template_Ip) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ip *Vpdn_Templates_Template_Ip) SetParent(parent types.Entity) { ip.parent = parent }

func (ip *Vpdn_Templates_Template_Ip) GetParent() types.Entity { return ip.parent }

func (ip *Vpdn_Templates_Template_Ip) GetParentYangName() string { return "template" }

// Vpdn_Templates_Template_Ipv4
// IPv4 settings for tunnel
type Vpdn_Templates_Template_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 don't fragment bit set/clear/reflect. The type is DfBit.
    DfBit interface{}

    // Enter an IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}
}

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Vpdn_Templates_Template_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetGoName(yname string) string {
    if yname == "df-bit" { return "DfBit" }
    if yname == "source" { return "Source" }
    return ""
}

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["df-bit"] = ipv4.DfBit
    leafs["source"] = ipv4.Source
    return leafs
}

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Vpdn_Templates_Template_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetParentYangName() string { return "template" }

// Vpdn_CallerId
// Options to apply on calling station ID
type Vpdn_CallerId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mask characters by method. The type is string with length: 1..63.
    Mask interface{}
}

func (callerId *Vpdn_CallerId) GetFilter() yfilter.YFilter { return callerId.YFilter }

func (callerId *Vpdn_CallerId) SetFilter(yf yfilter.YFilter) { callerId.YFilter = yf }

func (callerId *Vpdn_CallerId) GetGoName(yname string) string {
    if yname == "mask" { return "Mask" }
    return ""
}

func (callerId *Vpdn_CallerId) GetSegmentPath() string {
    return "caller-id"
}

func (callerId *Vpdn_CallerId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (callerId *Vpdn_CallerId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (callerId *Vpdn_CallerId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mask"] = callerId.Mask
    return leafs
}

func (callerId *Vpdn_CallerId) GetBundleName() string { return "cisco_ios_xr" }

func (callerId *Vpdn_CallerId) GetYangName() string { return "caller-id" }

func (callerId *Vpdn_CallerId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (callerId *Vpdn_CallerId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (callerId *Vpdn_CallerId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (callerId *Vpdn_CallerId) SetParent(parent types.Entity) { callerId.parent = parent }

func (callerId *Vpdn_CallerId) GetParent() types.Entity { return callerId.parent }

func (callerId *Vpdn_CallerId) GetParentYangName() string { return "vpdn" }

// Vpdn_VpdNgroups
// Table of VPDNgroup
type Vpdn_VpdNgroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // vpdn-group configuration. The type is slice of Vpdn_VpdNgroups_VpdNgroup.
    VpdNgroup []Vpdn_VpdNgroups_VpdNgroup
}

func (vpdNgroups *Vpdn_VpdNgroups) GetFilter() yfilter.YFilter { return vpdNgroups.YFilter }

func (vpdNgroups *Vpdn_VpdNgroups) SetFilter(yf yfilter.YFilter) { vpdNgroups.YFilter = yf }

func (vpdNgroups *Vpdn_VpdNgroups) GetGoName(yname string) string {
    if yname == "vpd-ngroup" { return "VpdNgroup" }
    return ""
}

func (vpdNgroups *Vpdn_VpdNgroups) GetSegmentPath() string {
    return "vpd-ngroups"
}

func (vpdNgroups *Vpdn_VpdNgroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vpd-ngroup" {
        for _, c := range vpdNgroups.VpdNgroup {
            if vpdNgroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vpdn_VpdNgroups_VpdNgroup{}
        vpdNgroups.VpdNgroup = append(vpdNgroups.VpdNgroup, child)
        return &vpdNgroups.VpdNgroup[len(vpdNgroups.VpdNgroup)-1]
    }
    return nil
}

func (vpdNgroups *Vpdn_VpdNgroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vpdNgroups.VpdNgroup {
        children[vpdNgroups.VpdNgroup[i].GetSegmentPath()] = &vpdNgroups.VpdNgroup[i]
    }
    return children
}

func (vpdNgroups *Vpdn_VpdNgroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vpdNgroups *Vpdn_VpdNgroups) GetBundleName() string { return "cisco_ios_xr" }

func (vpdNgroups *Vpdn_VpdNgroups) GetYangName() string { return "vpd-ngroups" }

func (vpdNgroups *Vpdn_VpdNgroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdNgroups *Vpdn_VpdNgroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdNgroups *Vpdn_VpdNgroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdNgroups *Vpdn_VpdNgroups) SetParent(parent types.Entity) { vpdNgroups.parent = parent }

func (vpdNgroups *Vpdn_VpdNgroups) GetParent() types.Entity { return vpdNgroups.parent }

func (vpdNgroups *Vpdn_VpdNgroups) GetParentYangName() string { return "vpdn" }

// Vpdn_VpdNgroups_VpdNgroup
// vpdn-group configuration
type Vpdn_VpdNgroups_VpdNgroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. vpdn-group name. The type is string with length:
    // 1..63.
    VpdNgroupname interface{}

    // Forward DSL Line Info attributes. The type is interface{}.
    DslLineForwarding interface{}

    // To support NAS-Port format e in cisco AVP 100. The type is interface{}.
    CiscoAvp100FormatEEnable interface{}

    // upto 100 characters describing this VPDN group. The type is string with
    // length: 1..100.
    Desc interface{}

    // match substring. The type is string with length: 1..63.
    Attribute interface{}

    // l2tp class name. The type is string with length: 1..79.
    L2TpClass interface{}

    // Busy list timeout length. The type is interface{} with range: 1..65535.
    TunnelBusyTimeout interface{}

    // Vrf name. The type is string with length: 1..32.
    VrfName interface{}

    // Source vpdn-template. The type is string with length: 1..63.
    SrCtemplate interface{}

    // Vpn id.
    VpnId Vpdn_VpdNgroups_VpdNgroup_VpnId

    // set ip tos value.
    Ip Vpdn_VpdNgroups_VpdNgroup_Ip
}

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetFilter() yfilter.YFilter { return vpdNgroup.YFilter }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) SetFilter(yf yfilter.YFilter) { vpdNgroup.YFilter = yf }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetGoName(yname string) string {
    if yname == "vpd-ngroupname" { return "VpdNgroupname" }
    if yname == "dsl-line-forwarding" { return "DslLineForwarding" }
    if yname == "cisco-avp100-format-e-enable" { return "CiscoAvp100FormatEEnable" }
    if yname == "desc" { return "Desc" }
    if yname == "attribute" { return "Attribute" }
    if yname == "l2tp-class" { return "L2TpClass" }
    if yname == "tunnel-busy-timeout" { return "TunnelBusyTimeout" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "sr-ctemplate" { return "SrCtemplate" }
    if yname == "vpn-id" { return "VpnId" }
    if yname == "ip" { return "Ip" }
    return ""
}

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetSegmentPath() string {
    return "vpd-ngroup" + "[vpd-ngroupname='" + fmt.Sprintf("%v", vpdNgroup.VpdNgroupname) + "']"
}

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vpn-id" {
        return &vpdNgroup.VpnId
    }
    if childYangName == "ip" {
        return &vpdNgroup.Ip
    }
    return nil
}

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vpn-id"] = &vpdNgroup.VpnId
    children["ip"] = &vpdNgroup.Ip
    return children
}

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpd-ngroupname"] = vpdNgroup.VpdNgroupname
    leafs["dsl-line-forwarding"] = vpdNgroup.DslLineForwarding
    leafs["cisco-avp100-format-e-enable"] = vpdNgroup.CiscoAvp100FormatEEnable
    leafs["desc"] = vpdNgroup.Desc
    leafs["attribute"] = vpdNgroup.Attribute
    leafs["l2tp-class"] = vpdNgroup.L2TpClass
    leafs["tunnel-busy-timeout"] = vpdNgroup.TunnelBusyTimeout
    leafs["vrf-name"] = vpdNgroup.VrfName
    leafs["sr-ctemplate"] = vpdNgroup.SrCtemplate
    return leafs
}

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetBundleName() string { return "cisco_ios_xr" }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetYangName() string { return "vpd-ngroup" }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) SetParent(parent types.Entity) { vpdNgroup.parent = parent }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetParent() types.Entity { return vpdNgroup.parent }

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetParentYangName() string { return "vpd-ngroups" }

// Vpdn_VpdNgroups_VpdNgroup_VpnId
// Vpn id
type Vpdn_VpdNgroups_VpdNgroup_VpnId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPN ID, (OUI:VPN-Index) format(hex), 3 bytes OUI Part. The type is string
    // with pattern: [0-9a-fA-F]{1,8}.
    VpnIdOui interface{}

    // VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part. The type is
    // string with pattern: [0-9a-fA-F]{1,8}.
    VpnIdIndex interface{}
}

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetFilter() yfilter.YFilter { return vpnId.YFilter }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) SetFilter(yf yfilter.YFilter) { vpnId.YFilter = yf }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetGoName(yname string) string {
    if yname == "vpn-id-oui" { return "VpnIdOui" }
    if yname == "vpn-id-index" { return "VpnIdIndex" }
    return ""
}

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetSegmentPath() string {
    return "vpn-id"
}

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpn-id-oui"] = vpnId.VpnIdOui
    leafs["vpn-id-index"] = vpnId.VpnIdIndex
    return leafs
}

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetBundleName() string { return "cisco_ios_xr" }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetYangName() string { return "vpn-id" }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) SetParent(parent types.Entity) { vpnId.parent = parent }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetParent() types.Entity { return vpnId.parent }

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetParentYangName() string { return "vpd-ngroup" }

// Vpdn_VpdNgroups_VpdNgroup_Ip
// set ip tos value
type Vpdn_VpdNgroups_VpdNgroup_Ip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ip tos value. The type is interface{} with range: 0..255.
    Tos interface{}
}

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetFilter() yfilter.YFilter { return ip.YFilter }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) SetFilter(yf yfilter.YFilter) { ip.YFilter = yf }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetGoName(yname string) string {
    if yname == "tos" { return "Tos" }
    return ""
}

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetSegmentPath() string {
    return "ip"
}

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tos"] = ip.Tos
    return leafs
}

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetBundleName() string { return "cisco_ios_xr" }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetYangName() string { return "ip" }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) SetParent(parent types.Entity) { ip.parent = parent }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetParent() types.Entity { return ip.parent }

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetParentYangName() string { return "vpd-ngroup" }

// Vpdn_Loggings
// Table of Logging
type Vpdn_Loggings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure logging for VPDN. The type is slice of Vpdn_Loggings_Logging.
    Logging []Vpdn_Loggings_Logging
}

func (loggings *Vpdn_Loggings) GetFilter() yfilter.YFilter { return loggings.YFilter }

func (loggings *Vpdn_Loggings) SetFilter(yf yfilter.YFilter) { loggings.YFilter = yf }

func (loggings *Vpdn_Loggings) GetGoName(yname string) string {
    if yname == "logging" { return "Logging" }
    return ""
}

func (loggings *Vpdn_Loggings) GetSegmentPath() string {
    return "loggings"
}

func (loggings *Vpdn_Loggings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logging" {
        for _, c := range loggings.Logging {
            if loggings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vpdn_Loggings_Logging{}
        loggings.Logging = append(loggings.Logging, child)
        return &loggings.Logging[len(loggings.Logging)-1]
    }
    return nil
}

func (loggings *Vpdn_Loggings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range loggings.Logging {
        children[loggings.Logging[i].GetSegmentPath()] = &loggings.Logging[i]
    }
    return children
}

func (loggings *Vpdn_Loggings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loggings *Vpdn_Loggings) GetBundleName() string { return "cisco_ios_xr" }

func (loggings *Vpdn_Loggings) GetYangName() string { return "loggings" }

func (loggings *Vpdn_Loggings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loggings *Vpdn_Loggings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loggings *Vpdn_Loggings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loggings *Vpdn_Loggings) SetParent(parent types.Entity) { loggings.parent = parent }

func (loggings *Vpdn_Loggings) GetParent() types.Entity { return loggings.parent }

func (loggings *Vpdn_Loggings) GetParentYangName() string { return "vpdn" }

// Vpdn_Loggings_Logging
// Configure logging for VPDN
type Vpdn_Loggings_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VPDN logging options. The type is Option.
    Option interface{}
}

func (logging *Vpdn_Loggings_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Vpdn_Loggings_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Vpdn_Loggings_Logging) GetGoName(yname string) string {
    if yname == "option" { return "Option" }
    return ""
}

func (logging *Vpdn_Loggings_Logging) GetSegmentPath() string {
    return "logging" + "[option='" + fmt.Sprintf("%v", logging.Option) + "']"
}

func (logging *Vpdn_Loggings_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logging *Vpdn_Loggings_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logging *Vpdn_Loggings_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["option"] = logging.Option
    return leafs
}

func (logging *Vpdn_Loggings_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Vpdn_Loggings_Logging) GetYangName() string { return "logging" }

func (logging *Vpdn_Loggings_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Vpdn_Loggings_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Vpdn_Loggings_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Vpdn_Loggings_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Vpdn_Loggings_Logging) GetParent() types.Entity { return logging.parent }

func (logging *Vpdn_Loggings_Logging) GetParentYangName() string { return "loggings" }

// Vpdn_L2Tp
// L2TPv2 protocol commands
type Vpdn_L2Tp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP MSS adjust value. The acceptable values might be further limited
    // depending on platform. The type is interface{} with range: 1280..1460.
    TcpMssAdjust interface{}

    // L2TP IP packet reassembly enable. The type is interface{}.
    Reassembly interface{}

    // Session ID commands.
    SessionId Vpdn_L2Tp_SessionId
}

func (l2Tp *Vpdn_L2Tp) GetFilter() yfilter.YFilter { return l2Tp.YFilter }

func (l2Tp *Vpdn_L2Tp) SetFilter(yf yfilter.YFilter) { l2Tp.YFilter = yf }

func (l2Tp *Vpdn_L2Tp) GetGoName(yname string) string {
    if yname == "tcp-mss-adjust" { return "TcpMssAdjust" }
    if yname == "reassembly" { return "Reassembly" }
    if yname == "session-id" { return "SessionId" }
    return ""
}

func (l2Tp *Vpdn_L2Tp) GetSegmentPath() string {
    return "l2tp"
}

func (l2Tp *Vpdn_L2Tp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-id" {
        return &l2Tp.SessionId
    }
    return nil
}

func (l2Tp *Vpdn_L2Tp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session-id"] = &l2Tp.SessionId
    return children
}

func (l2Tp *Vpdn_L2Tp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-mss-adjust"] = l2Tp.TcpMssAdjust
    leafs["reassembly"] = l2Tp.Reassembly
    return leafs
}

func (l2Tp *Vpdn_L2Tp) GetBundleName() string { return "cisco_ios_xr" }

func (l2Tp *Vpdn_L2Tp) GetYangName() string { return "l2tp" }

func (l2Tp *Vpdn_L2Tp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Tp *Vpdn_L2Tp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Tp *Vpdn_L2Tp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Tp *Vpdn_L2Tp) SetParent(parent types.Entity) { l2Tp.parent = parent }

func (l2Tp *Vpdn_L2Tp) GetParent() types.Entity { return l2Tp.parent }

func (l2Tp *Vpdn_L2Tp) GetParentYangName() string { return "vpdn" }

// Vpdn_L2Tp_SessionId
// Session ID commands
type Vpdn_L2Tp_SessionId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID space commands.
    Space Vpdn_L2Tp_SessionId_Space
}

func (sessionId *Vpdn_L2Tp_SessionId) GetFilter() yfilter.YFilter { return sessionId.YFilter }

func (sessionId *Vpdn_L2Tp_SessionId) SetFilter(yf yfilter.YFilter) { sessionId.YFilter = yf }

func (sessionId *Vpdn_L2Tp_SessionId) GetGoName(yname string) string {
    if yname == "space" { return "Space" }
    return ""
}

func (sessionId *Vpdn_L2Tp_SessionId) GetSegmentPath() string {
    return "session-id"
}

func (sessionId *Vpdn_L2Tp_SessionId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "space" {
        return &sessionId.Space
    }
    return nil
}

func (sessionId *Vpdn_L2Tp_SessionId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["space"] = &sessionId.Space
    return children
}

func (sessionId *Vpdn_L2Tp_SessionId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessionId *Vpdn_L2Tp_SessionId) GetBundleName() string { return "cisco_ios_xr" }

func (sessionId *Vpdn_L2Tp_SessionId) GetYangName() string { return "session-id" }

func (sessionId *Vpdn_L2Tp_SessionId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionId *Vpdn_L2Tp_SessionId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionId *Vpdn_L2Tp_SessionId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionId *Vpdn_L2Tp_SessionId) SetParent(parent types.Entity) { sessionId.parent = parent }

func (sessionId *Vpdn_L2Tp_SessionId) GetParent() types.Entity { return sessionId.parent }

func (sessionId *Vpdn_L2Tp_SessionId) GetParentYangName() string { return "l2tp" }

// Vpdn_L2Tp_SessionId_Space
// Session ID space commands
type Vpdn_L2Tp_SessionId_Space struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID space hierarchical command. The type is interface{}.
    Hierarchy interface{}
}

func (space *Vpdn_L2Tp_SessionId_Space) GetFilter() yfilter.YFilter { return space.YFilter }

func (space *Vpdn_L2Tp_SessionId_Space) SetFilter(yf yfilter.YFilter) { space.YFilter = yf }

func (space *Vpdn_L2Tp_SessionId_Space) GetGoName(yname string) string {
    if yname == "hierarchy" { return "Hierarchy" }
    return ""
}

func (space *Vpdn_L2Tp_SessionId_Space) GetSegmentPath() string {
    return "space"
}

func (space *Vpdn_L2Tp_SessionId_Space) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (space *Vpdn_L2Tp_SessionId_Space) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (space *Vpdn_L2Tp_SessionId_Space) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hierarchy"] = space.Hierarchy
    return leafs
}

func (space *Vpdn_L2Tp_SessionId_Space) GetBundleName() string { return "cisco_ios_xr" }

func (space *Vpdn_L2Tp_SessionId_Space) GetYangName() string { return "space" }

func (space *Vpdn_L2Tp_SessionId_Space) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (space *Vpdn_L2Tp_SessionId_Space) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (space *Vpdn_L2Tp_SessionId_Space) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (space *Vpdn_L2Tp_SessionId_Space) SetParent(parent types.Entity) { space.parent = parent }

func (space *Vpdn_L2Tp_SessionId_Space) GetParent() types.Entity { return space.parent }

func (space *Vpdn_L2Tp_SessionId_Space) GetParentYangName() string { return "session-id" }

