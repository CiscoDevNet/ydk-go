// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-vpdn package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vpdn: VPDN configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
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
    L2tp Vpdn_L2tp
}

func (vpdn *Vpdn) GetEntityData() *types.CommonEntityData {
    vpdn.EntityData.YFilter = vpdn.YFilter
    vpdn.EntityData.YangName = "vpdn"
    vpdn.EntityData.BundleName = "cisco_ios_xr"
    vpdn.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-vpdn-cfg"
    vpdn.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-vpdn-cfg:vpdn"
    vpdn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdn.EntityData.Children = types.NewOrderedMap()
    vpdn.EntityData.Children.Append("history", types.YChild{"History", &vpdn.History})
    vpdn.EntityData.Children.Append("redundancy", types.YChild{"Redundancy", &vpdn.Redundancy})
    vpdn.EntityData.Children.Append("local", types.YChild{"Local", &vpdn.Local})
    vpdn.EntityData.Children.Append("templates", types.YChild{"Templates", &vpdn.Templates})
    vpdn.EntityData.Children.Append("caller-id", types.YChild{"CallerId", &vpdn.CallerId})
    vpdn.EntityData.Children.Append("vpd-ngroups", types.YChild{"VpdNgroups", &vpdn.VpdNgroups})
    vpdn.EntityData.Children.Append("loggings", types.YChild{"Loggings", &vpdn.Loggings})
    vpdn.EntityData.Children.Append("l2tp", types.YChild{"L2tp", &vpdn.L2tp})
    vpdn.EntityData.Leafs = types.NewOrderedMap()
    vpdn.EntityData.Leafs.Append("session-limit", types.YLeaf{"SessionLimit", vpdn.SessionLimit})
    vpdn.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vpdn.Enable})
    vpdn.EntityData.Leafs.Append("soft-shut", types.YLeaf{"SoftShut", vpdn.SoftShut})

    vpdn.EntityData.YListKeys = []string {}

    return &(vpdn.EntityData)
}

// Vpdn_History
// VPDN history logging
type Vpdn_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User failure. The type is interface{}.
    Failure interface{}
}

func (history *Vpdn_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "vpdn"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("failure", types.YLeaf{"Failure", history.Failure})

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Vpdn_Redundancy
// Enable VPDN redundancy
type Vpdn_Redundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Enable VPDN redundancy. Deletion of this object also causes deletion
    // of all associated objects under Redundancy. The type is interface{}.
    Enable interface{}

    // Process crash configuration.
    ProcessFailures Vpdn_Redundancy_ProcessFailures
}

func (redundancy *Vpdn_Redundancy) GetEntityData() *types.CommonEntityData {
    redundancy.EntityData.YFilter = redundancy.YFilter
    redundancy.EntityData.YangName = "redundancy"
    redundancy.EntityData.BundleName = "cisco_ios_xr"
    redundancy.EntityData.ParentYangName = "vpdn"
    redundancy.EntityData.SegmentPath = "redundancy"
    redundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy.EntityData.Children = types.NewOrderedMap()
    redundancy.EntityData.Children.Append("process-failures", types.YChild{"ProcessFailures", &redundancy.ProcessFailures})
    redundancy.EntityData.Leafs = types.NewOrderedMap()
    redundancy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", redundancy.Enable})

    redundancy.EntityData.YListKeys = []string {}

    return &(redundancy.EntityData)
}

// Vpdn_Redundancy_ProcessFailures
// Process crash configuration
type Vpdn_Redundancy_ProcessFailures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Force a switchover if the process crashes. The type is interface{}.
    Switchover interface{}
}

func (processFailures *Vpdn_Redundancy_ProcessFailures) GetEntityData() *types.CommonEntityData {
    processFailures.EntityData.YFilter = processFailures.YFilter
    processFailures.EntityData.YangName = "process-failures"
    processFailures.EntityData.BundleName = "cisco_ios_xr"
    processFailures.EntityData.ParentYangName = "redundancy"
    processFailures.EntityData.SegmentPath = "process-failures"
    processFailures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processFailures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processFailures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processFailures.EntityData.Children = types.NewOrderedMap()
    processFailures.EntityData.Leafs = types.NewOrderedMap()
    processFailures.EntityData.Leafs.Append("switchover", types.YLeaf{"Switchover", processFailures.Switchover})

    processFailures.EntityData.YListKeys = []string {}

    return &(processFailures.EntityData)
}

// Vpdn_Local
// VPDN Local radius process configuration
type Vpdn_Local struct {
    EntityData types.CommonEntityData
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

func (local *Vpdn_Local) GetEntityData() *types.CommonEntityData {
    local.EntityData.YFilter = local.YFilter
    local.EntityData.YangName = "local"
    local.EntityData.BundleName = "cisco_ios_xr"
    local.EntityData.ParentYangName = "vpdn"
    local.EntityData.SegmentPath = "local"
    local.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    local.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    local.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    local.EntityData.Children = types.NewOrderedMap()
    local.EntityData.Leafs = types.NewOrderedMap()
    local.EntityData.Leafs.Append("secret-text", types.YLeaf{"SecretText", local.SecretText})
    local.EntityData.Leafs.Append("path", types.YLeaf{"Path", local.Path})
    local.EntityData.Leafs.Append("cache-disabled", types.YLeaf{"CacheDisabled", local.CacheDisabled})
    local.EntityData.Leafs.Append("port", types.YLeaf{"Port", local.Port})

    local.EntityData.YListKeys = []string {}

    return &(local.EntityData)
}

// Vpdn_Templates
// Table of Template
type Vpdn_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN template configuration. The type is slice of Vpdn_Templates_Template.
    Template []*Vpdn_Templates_Template
}

func (templates *Vpdn_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "vpdn"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// Vpdn_Templates_Template
// VPDN template configuration
type Vpdn_Templates_Template struct {
    EntityData types.CommonEntityData
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
    L2tpClass interface{}

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

func (template *Vpdn_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Children.Append("caller-id", types.YChild{"CallerId", &template.CallerId})
    template.EntityData.Children.Append("vpn", types.YChild{"Vpn", &template.Vpn})
    template.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", &template.Tunnel})
    template.EntityData.Children.Append("ip", types.YChild{"Ip", &template.Ip})
    template.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &template.Ipv4})
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("cisco-avp100-format-e-enable", types.YLeaf{"CiscoAvp100FormatEEnable", template.CiscoAvp100FormatEEnable})
    template.EntityData.Leafs.Append("description", types.YLeaf{"Description", template.Description})
    template.EntityData.Leafs.Append("l2tp-class", types.YLeaf{"L2tpClass", template.L2tpClass})
    template.EntityData.Leafs.Append("dsl-line-forwarding", types.YLeaf{"DslLineForwarding", template.DslLineForwarding})

    template.EntityData.YListKeys = []string {"TemplateName"}

    return &(template.EntityData)
}

// Vpdn_Templates_Template_CallerId
// Options to apply on calling station id
type Vpdn_Templates_Template_CallerId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mask characters by method. The type is string with length: 1..63.
    Mask interface{}
}

func (callerId *Vpdn_Templates_Template_CallerId) GetEntityData() *types.CommonEntityData {
    callerId.EntityData.YFilter = callerId.YFilter
    callerId.EntityData.YangName = "caller-id"
    callerId.EntityData.BundleName = "cisco_ios_xr"
    callerId.EntityData.ParentYangName = "template"
    callerId.EntityData.SegmentPath = "caller-id"
    callerId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    callerId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    callerId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    callerId.EntityData.Children = types.NewOrderedMap()
    callerId.EntityData.Leafs = types.NewOrderedMap()
    callerId.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", callerId.Mask})

    callerId.EntityData.YListKeys = []string {}

    return &(callerId.EntityData)
}

// Vpdn_Templates_Template_Vpn
// VPN ID/VRF name
type Vpdn_Templates_Template_Vpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with length: 1..32.
    Vrf interface{}

    // VPN ID.
    Id Vpdn_Templates_Template_Vpn_Id
}

func (vpn *Vpdn_Templates_Template_Vpn) GetEntityData() *types.CommonEntityData {
    vpn.EntityData.YFilter = vpn.YFilter
    vpn.EntityData.YangName = "vpn"
    vpn.EntityData.BundleName = "cisco_ios_xr"
    vpn.EntityData.ParentYangName = "template"
    vpn.EntityData.SegmentPath = "vpn"
    vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpn.EntityData.Children = types.NewOrderedMap()
    vpn.EntityData.Children.Append("id", types.YChild{"Id", &vpn.Id})
    vpn.EntityData.Leafs = types.NewOrderedMap()
    vpn.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", vpn.Vrf})

    vpn.EntityData.YListKeys = []string {}

    return &(vpn.EntityData)
}

// Vpdn_Templates_Template_Vpn_Id
// VPN ID
type Vpdn_Templates_Template_Vpn_Id struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPN ID, (OUI:VPN-Index) format(hex), 3 bytes OUI Part. The type is string
    // with pattern: [0-9a-fA-F]{1,8}.
    Oui interface{}

    // VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part. The type is
    // string with pattern: [0-9a-fA-F]{1,8}.
    Index interface{}
}

func (id *Vpdn_Templates_Template_Vpn_Id) GetEntityData() *types.CommonEntityData {
    id.EntityData.YFilter = id.YFilter
    id.EntityData.YangName = "id"
    id.EntityData.BundleName = "cisco_ios_xr"
    id.EntityData.ParentYangName = "vpn"
    id.EntityData.SegmentPath = "id"
    id.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    id.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    id.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    id.EntityData.Children = types.NewOrderedMap()
    id.EntityData.Leafs = types.NewOrderedMap()
    id.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", id.Oui})
    id.EntityData.Leafs.Append("index", types.YLeaf{"Index", id.Index})

    id.EntityData.YListKeys = []string {}

    return &(id.EntityData)
}

// Vpdn_Templates_Template_Tunnel
// L2TP tunnel commands
type Vpdn_Templates_Template_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Busy time out value in seconds. The type is interface{} with range:
    // 60..65535. Units are second.
    BusyTimeout interface{}
}

func (tunnel *Vpdn_Templates_Template_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "template"
    tunnel.EntityData.SegmentPath = "tunnel"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("busy-timeout", types.YLeaf{"BusyTimeout", tunnel.BusyTimeout})

    tunnel.EntityData.YListKeys = []string {}

    return &(tunnel.EntityData)
}

// Vpdn_Templates_Template_Ip
// Set IP TOS value
type Vpdn_Templates_Template_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set constant integer. The type is interface{} with range: 0..4294967295.
    Tos interface{}
}

func (ip *Vpdn_Templates_Template_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "template"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("tos", types.YLeaf{"Tos", ip.Tos})

    ip.EntityData.YListKeys = []string {}

    return &(ip.EntityData)
}

// Vpdn_Templates_Template_Ipv4
// IPv4 settings for tunnel
type Vpdn_Templates_Template_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 don't fragment bit set/clear/reflect. The type is DfBit.
    DfBit interface{}

    // Enter an IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}
}

func (ipv4 *Vpdn_Templates_Template_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "template"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("df-bit", types.YLeaf{"DfBit", ipv4.DfBit})
    ipv4.EntityData.Leafs.Append("source", types.YLeaf{"Source", ipv4.Source})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Vpdn_CallerId
// Options to apply on calling station ID
type Vpdn_CallerId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mask characters by method. The type is string with length: 1..63.
    Mask interface{}
}

func (callerId *Vpdn_CallerId) GetEntityData() *types.CommonEntityData {
    callerId.EntityData.YFilter = callerId.YFilter
    callerId.EntityData.YangName = "caller-id"
    callerId.EntityData.BundleName = "cisco_ios_xr"
    callerId.EntityData.ParentYangName = "vpdn"
    callerId.EntityData.SegmentPath = "caller-id"
    callerId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    callerId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    callerId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    callerId.EntityData.Children = types.NewOrderedMap()
    callerId.EntityData.Leafs = types.NewOrderedMap()
    callerId.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", callerId.Mask})

    callerId.EntityData.YListKeys = []string {}

    return &(callerId.EntityData)
}

// Vpdn_VpdNgroups
// Table of VPDNgroup
type Vpdn_VpdNgroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // vpdn-group configuration. The type is slice of Vpdn_VpdNgroups_VpdNgroup.
    VpdNgroup []*Vpdn_VpdNgroups_VpdNgroup
}

func (vpdNgroups *Vpdn_VpdNgroups) GetEntityData() *types.CommonEntityData {
    vpdNgroups.EntityData.YFilter = vpdNgroups.YFilter
    vpdNgroups.EntityData.YangName = "vpd-ngroups"
    vpdNgroups.EntityData.BundleName = "cisco_ios_xr"
    vpdNgroups.EntityData.ParentYangName = "vpdn"
    vpdNgroups.EntityData.SegmentPath = "vpd-ngroups"
    vpdNgroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdNgroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdNgroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdNgroups.EntityData.Children = types.NewOrderedMap()
    vpdNgroups.EntityData.Children.Append("vpd-ngroup", types.YChild{"VpdNgroup", nil})
    for i := range vpdNgroups.VpdNgroup {
        vpdNgroups.EntityData.Children.Append(types.GetSegmentPath(vpdNgroups.VpdNgroup[i]), types.YChild{"VpdNgroup", vpdNgroups.VpdNgroup[i]})
    }
    vpdNgroups.EntityData.Leafs = types.NewOrderedMap()

    vpdNgroups.EntityData.YListKeys = []string {}

    return &(vpdNgroups.EntityData)
}

// Vpdn_VpdNgroups_VpdNgroup
// vpdn-group configuration
type Vpdn_VpdNgroups_VpdNgroup struct {
    EntityData types.CommonEntityData
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
    L2tpClass interface{}

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

func (vpdNgroup *Vpdn_VpdNgroups_VpdNgroup) GetEntityData() *types.CommonEntityData {
    vpdNgroup.EntityData.YFilter = vpdNgroup.YFilter
    vpdNgroup.EntityData.YangName = "vpd-ngroup"
    vpdNgroup.EntityData.BundleName = "cisco_ios_xr"
    vpdNgroup.EntityData.ParentYangName = "vpd-ngroups"
    vpdNgroup.EntityData.SegmentPath = "vpd-ngroup" + types.AddKeyToken(vpdNgroup.VpdNgroupname, "vpd-ngroupname")
    vpdNgroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdNgroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdNgroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdNgroup.EntityData.Children = types.NewOrderedMap()
    vpdNgroup.EntityData.Children.Append("vpn-id", types.YChild{"VpnId", &vpdNgroup.VpnId})
    vpdNgroup.EntityData.Children.Append("ip", types.YChild{"Ip", &vpdNgroup.Ip})
    vpdNgroup.EntityData.Leafs = types.NewOrderedMap()
    vpdNgroup.EntityData.Leafs.Append("vpd-ngroupname", types.YLeaf{"VpdNgroupname", vpdNgroup.VpdNgroupname})
    vpdNgroup.EntityData.Leafs.Append("dsl-line-forwarding", types.YLeaf{"DslLineForwarding", vpdNgroup.DslLineForwarding})
    vpdNgroup.EntityData.Leafs.Append("cisco-avp100-format-e-enable", types.YLeaf{"CiscoAvp100FormatEEnable", vpdNgroup.CiscoAvp100FormatEEnable})
    vpdNgroup.EntityData.Leafs.Append("desc", types.YLeaf{"Desc", vpdNgroup.Desc})
    vpdNgroup.EntityData.Leafs.Append("attribute", types.YLeaf{"Attribute", vpdNgroup.Attribute})
    vpdNgroup.EntityData.Leafs.Append("l2tp-class", types.YLeaf{"L2tpClass", vpdNgroup.L2tpClass})
    vpdNgroup.EntityData.Leafs.Append("tunnel-busy-timeout", types.YLeaf{"TunnelBusyTimeout", vpdNgroup.TunnelBusyTimeout})
    vpdNgroup.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vpdNgroup.VrfName})
    vpdNgroup.EntityData.Leafs.Append("sr-ctemplate", types.YLeaf{"SrCtemplate", vpdNgroup.SrCtemplate})

    vpdNgroup.EntityData.YListKeys = []string {"VpdNgroupname"}

    return &(vpdNgroup.EntityData)
}

// Vpdn_VpdNgroups_VpdNgroup_VpnId
// Vpn id
type Vpdn_VpdNgroups_VpdNgroup_VpnId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPN ID, (OUI:VPN-Index) format(hex), 3 bytes OUI Part. The type is string
    // with pattern: [0-9a-fA-F]{1,8}.
    VpnIdOui interface{}

    // VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part. The type is
    // string with pattern: [0-9a-fA-F]{1,8}.
    VpnIdIndex interface{}
}

func (vpnId *Vpdn_VpdNgroups_VpdNgroup_VpnId) GetEntityData() *types.CommonEntityData {
    vpnId.EntityData.YFilter = vpnId.YFilter
    vpnId.EntityData.YangName = "vpn-id"
    vpnId.EntityData.BundleName = "cisco_ios_xr"
    vpnId.EntityData.ParentYangName = "vpd-ngroup"
    vpnId.EntityData.SegmentPath = "vpn-id"
    vpnId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpnId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpnId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpnId.EntityData.Children = types.NewOrderedMap()
    vpnId.EntityData.Leafs = types.NewOrderedMap()
    vpnId.EntityData.Leafs.Append("vpn-id-oui", types.YLeaf{"VpnIdOui", vpnId.VpnIdOui})
    vpnId.EntityData.Leafs.Append("vpn-id-index", types.YLeaf{"VpnIdIndex", vpnId.VpnIdIndex})

    vpnId.EntityData.YListKeys = []string {}

    return &(vpnId.EntityData)
}

// Vpdn_VpdNgroups_VpdNgroup_Ip
// set ip tos value
type Vpdn_VpdNgroups_VpdNgroup_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ip tos value. The type is interface{} with range: 0..255.
    Tos interface{}
}

func (ip *Vpdn_VpdNgroups_VpdNgroup_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "vpd-ngroup"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("tos", types.YLeaf{"Tos", ip.Tos})

    ip.EntityData.YListKeys = []string {}

    return &(ip.EntityData)
}

// Vpdn_Loggings
// Table of Logging
type Vpdn_Loggings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure logging for VPDN. The type is slice of Vpdn_Loggings_Logging.
    Logging []*Vpdn_Loggings_Logging
}

func (loggings *Vpdn_Loggings) GetEntityData() *types.CommonEntityData {
    loggings.EntityData.YFilter = loggings.YFilter
    loggings.EntityData.YangName = "loggings"
    loggings.EntityData.BundleName = "cisco_ios_xr"
    loggings.EntityData.ParentYangName = "vpdn"
    loggings.EntityData.SegmentPath = "loggings"
    loggings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loggings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loggings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loggings.EntityData.Children = types.NewOrderedMap()
    loggings.EntityData.Children.Append("logging", types.YChild{"Logging", nil})
    for i := range loggings.Logging {
        loggings.EntityData.Children.Append(types.GetSegmentPath(loggings.Logging[i]), types.YChild{"Logging", loggings.Logging[i]})
    }
    loggings.EntityData.Leafs = types.NewOrderedMap()

    loggings.EntityData.YListKeys = []string {}

    return &(loggings.EntityData)
}

// Vpdn_Loggings_Logging
// Configure logging for VPDN
type Vpdn_Loggings_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VPDN logging options. The type is Option.
    Option interface{}
}

func (logging *Vpdn_Loggings_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "loggings"
    logging.EntityData.SegmentPath = "logging" + types.AddKeyToken(logging.Option, "option")
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("option", types.YLeaf{"Option", logging.Option})

    logging.EntityData.YListKeys = []string {"Option"}

    return &(logging.EntityData)
}

// Vpdn_L2tp
// L2TPv2 protocol commands
type Vpdn_L2tp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP MSS adjust value. The acceptable values might be further limited
    // depending on platform. The type is interface{} with range: 1280..1460.
    TcpMssAdjust interface{}

    // L2TP IP packet reassembly enable. The type is interface{}.
    Reassembly interface{}

    // Session ID commands.
    SessionId Vpdn_L2tp_SessionId
}

func (l2tp *Vpdn_L2tp) GetEntityData() *types.CommonEntityData {
    l2tp.EntityData.YFilter = l2tp.YFilter
    l2tp.EntityData.YangName = "l2tp"
    l2tp.EntityData.BundleName = "cisco_ios_xr"
    l2tp.EntityData.ParentYangName = "vpdn"
    l2tp.EntityData.SegmentPath = "l2tp"
    l2tp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tp.EntityData.Children = types.NewOrderedMap()
    l2tp.EntityData.Children.Append("session-id", types.YChild{"SessionId", &l2tp.SessionId})
    l2tp.EntityData.Leafs = types.NewOrderedMap()
    l2tp.EntityData.Leafs.Append("tcp-mss-adjust", types.YLeaf{"TcpMssAdjust", l2tp.TcpMssAdjust})
    l2tp.EntityData.Leafs.Append("reassembly", types.YLeaf{"Reassembly", l2tp.Reassembly})

    l2tp.EntityData.YListKeys = []string {}

    return &(l2tp.EntityData)
}

// Vpdn_L2tp_SessionId
// Session ID commands
type Vpdn_L2tp_SessionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session ID space commands.
    Space Vpdn_L2tp_SessionId_Space
}

func (sessionId *Vpdn_L2tp_SessionId) GetEntityData() *types.CommonEntityData {
    sessionId.EntityData.YFilter = sessionId.YFilter
    sessionId.EntityData.YangName = "session-id"
    sessionId.EntityData.BundleName = "cisco_ios_xr"
    sessionId.EntityData.ParentYangName = "l2tp"
    sessionId.EntityData.SegmentPath = "session-id"
    sessionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionId.EntityData.Children = types.NewOrderedMap()
    sessionId.EntityData.Children.Append("space", types.YChild{"Space", &sessionId.Space})
    sessionId.EntityData.Leafs = types.NewOrderedMap()

    sessionId.EntityData.YListKeys = []string {}

    return &(sessionId.EntityData)
}

// Vpdn_L2tp_SessionId_Space
// Session ID space commands
type Vpdn_L2tp_SessionId_Space struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session ID space hierarchical command. The type is interface{}.
    Hierarchy interface{}
}

func (space *Vpdn_L2tp_SessionId_Space) GetEntityData() *types.CommonEntityData {
    space.EntityData.YFilter = space.YFilter
    space.EntityData.YangName = "space"
    space.EntityData.BundleName = "cisco_ios_xr"
    space.EntityData.ParentYangName = "session-id"
    space.EntityData.SegmentPath = "space"
    space.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    space.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    space.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    space.EntityData.Children = types.NewOrderedMap()
    space.EntityData.Leafs = types.NewOrderedMap()
    space.EntityData.Leafs.Append("hierarchy", types.YLeaf{"Hierarchy", space.Hierarchy})

    space.EntityData.YListKeys = []string {}

    return &(space.EntityData)
}

