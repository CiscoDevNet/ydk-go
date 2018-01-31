// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rsi package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vrfs: VRF configuration
//   global-af: global af
//   srlg: srlg
//   vrf-groups: vrf groups
//   selective-vrf-download: selective vrf download
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-snmp-agent-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_rsi_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rsi_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg vrfs}", reflect.TypeOf(Vrfs{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:vrfs", reflect.TypeOf(Vrfs{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg global-af}", reflect.TypeOf(GlobalAf{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:global-af", reflect.TypeOf(GlobalAf{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg srlg}", reflect.TypeOf(Srlg{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:srlg", reflect.TypeOf(Srlg{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg vrf-groups}", reflect.TypeOf(VrfGroups{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:vrf-groups", reflect.TypeOf(VrfGroups{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rsi-cfg selective-vrf-download}", reflect.TypeOf(SelectiveVrfDownload{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download", reflect.TypeOf(SelectiveVrfDownload{}))
}

// VrfAddressFamily represents Vrf address family
type VrfAddressFamily string

const (
    // IPv4
    VrfAddressFamily_ipv4 VrfAddressFamily = "ipv4"

    // IPv6
    VrfAddressFamily_ipv6 VrfAddressFamily = "ipv6"
)

// SrlgPriority represents Srlg priority
type SrlgPriority string

const (
    // Critical
    SrlgPriority_critical SrlgPriority = "critical"

    // High
    SrlgPriority_high SrlgPriority = "high"

    // Default
    SrlgPriority_default_ SrlgPriority = "default"

    // Low
    SrlgPriority_low SrlgPriority = "low"

    // Very low
    SrlgPriority_very_low SrlgPriority = "very-low"
)

// VrfSubAddressFamily represents Vrf sub address family
type VrfSubAddressFamily string

const (
    // Unicast
    VrfSubAddressFamily_unicast VrfSubAddressFamily = "unicast"

    // Multicast
    VrfSubAddressFamily_multicast VrfSubAddressFamily = "multicast"

    // Flow spec
    VrfSubAddressFamily_flow_spec VrfSubAddressFamily = "flow-spec"
)

// Vrfs
// VRF configuration
type Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF configuration. The type is slice of Vrfs_Vrf.
    Vrf []Vrfs_Vrf
}

func (vrfs *Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Vrfs) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-cfg:vrfs"
}

func (vrfs *Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Vrfs) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-cfg" }

// Vrfs_Vrf
// VRF configuration
type Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // Fallback VRF. The type is string with length: 1..32.
    FallbackVrf interface{}

    // For disabling remote route filtering for this VRF on core-facing card. The
    // type is interface{}.
    RemoteRouteFilterDisable interface{}

    // VRF global configuration. The type is interface{}.
    Create interface{}

    // Configuration enable of big VRF. The type is interface{}.
    ModeBig interface{}

    // A textual description of the VRF. The type is string with length: 1..244.
    Description interface{}

    // VPN-ID for the VRF.
    VpnId Vrfs_Vrf_VpnId

    // VRF address family configuration.
    Afs Vrfs_Vrf_Afs

    // Multicast host stack configuration.
    MulticastHost Vrfs_Vrf_MulticastHost
}

func (vrf *Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "fallback-vrf" { return "FallbackVrf" }
    if yname == "remote-route-filter-disable" { return "RemoteRouteFilterDisable" }
    if yname == "create" { return "Create" }
    if yname == "mode-big" { return "ModeBig" }
    if yname == "description" { return "Description" }
    if yname == "vpn-id" { return "VpnId" }
    if yname == "afs" { return "Afs" }
    if yname == "Cisco-IOS-XR-ip-iarm-vrf-cfg:multicast-host" { return "MulticastHost" }
    return ""
}

func (vrf *Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vpn-id" {
        return &vrf.VpnId
    }
    if childYangName == "afs" {
        return &vrf.Afs
    }
    if childYangName == "Cisco-IOS-XR-ip-iarm-vrf-cfg:multicast-host" {
        return &vrf.MulticastHost
    }
    return nil
}

func (vrf *Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vpn-id"] = &vrf.VpnId
    children["afs"] = &vrf.Afs
    children["Cisco-IOS-XR-ip-iarm-vrf-cfg:multicast-host"] = &vrf.MulticastHost
    return children
}

func (vrf *Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["fallback-vrf"] = vrf.FallbackVrf
    leafs["remote-route-filter-disable"] = vrf.RemoteRouteFilterDisable
    leafs["create"] = vrf.Create
    leafs["mode-big"] = vrf.ModeBig
    leafs["description"] = vrf.Description
    return leafs
}

func (vrf *Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Vrfs_Vrf_VpnId
// VPN-ID for the VRF
// This type is a presence type.
type Vrfs_Vrf_VpnId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OUI of VPNID OUI. The type is interface{} with range: 0..16777215. This
    // attribute is mandatory.
    VpnOui interface{}

    // Index of VPNID Index. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory.
    VpnIndex interface{}
}

func (vpnId *Vrfs_Vrf_VpnId) GetFilter() yfilter.YFilter { return vpnId.YFilter }

func (vpnId *Vrfs_Vrf_VpnId) SetFilter(yf yfilter.YFilter) { vpnId.YFilter = yf }

func (vpnId *Vrfs_Vrf_VpnId) GetGoName(yname string) string {
    if yname == "vpn-oui" { return "VpnOui" }
    if yname == "vpn-index" { return "VpnIndex" }
    return ""
}

func (vpnId *Vrfs_Vrf_VpnId) GetSegmentPath() string {
    return "vpn-id"
}

func (vpnId *Vrfs_Vrf_VpnId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpnId *Vrfs_Vrf_VpnId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpnId *Vrfs_Vrf_VpnId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpn-oui"] = vpnId.VpnOui
    leafs["vpn-index"] = vpnId.VpnIndex
    return leafs
}

func (vpnId *Vrfs_Vrf_VpnId) GetBundleName() string { return "cisco_ios_xr" }

func (vpnId *Vrfs_Vrf_VpnId) GetYangName() string { return "vpn-id" }

func (vpnId *Vrfs_Vrf_VpnId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpnId *Vrfs_Vrf_VpnId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpnId *Vrfs_Vrf_VpnId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpnId *Vrfs_Vrf_VpnId) SetParent(parent types.Entity) { vpnId.parent = parent }

func (vpnId *Vrfs_Vrf_VpnId) GetParent() types.Entity { return vpnId.parent }

func (vpnId *Vrfs_Vrf_VpnId) GetParentYangName() string { return "vrf" }

// Vrfs_Vrf_Afs
// VRF address family configuration
type Vrfs_Vrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF address family configuration. The type is slice of Vrfs_Vrf_Afs_Af.
    Af []Vrfs_Vrf_Afs_Af
}

func (afs *Vrfs_Vrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *Vrfs_Vrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *Vrfs_Vrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *Vrfs_Vrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *Vrfs_Vrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *Vrfs_Vrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *Vrfs_Vrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *Vrfs_Vrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *Vrfs_Vrf_Afs) GetYangName() string { return "afs" }

func (afs *Vrfs_Vrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *Vrfs_Vrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *Vrfs_Vrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *Vrfs_Vrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *Vrfs_Vrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *Vrfs_Vrf_Afs) GetParentYangName() string { return "vrf" }

// Vrfs_Vrf_Afs_Af
// VRF address family configuration
type Vrfs_Vrf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is VrfAddressFamily.
    AfName interface{}

    // This attribute is a key. Sub-Address family. The type is
    // VrfSubAddressFamily.
    SafName interface{}

    // This attribute is a key. Topology name. The type is string with length:
    // 1..244.
    TopologyName interface{}

    // VRF configuration for a particular address family. The type is interface{}.
    Create interface{}

    // Set maximum prefix limits.
    MaximumPrefix Vrfs_Vrf_Afs_Af_MaximumPrefix

    // BGP AF VRF config.
    Bgp Vrfs_Vrf_Afs_Af_Bgp
}

func (af *Vrfs_Vrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *Vrfs_Vrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *Vrfs_Vrf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "topology-name" { return "TopologyName" }
    if yname == "create" { return "Create" }
    if yname == "Cisco-IOS-XR-ip-rib-cfg:maximum-prefix" { return "MaximumPrefix" }
    if yname == "Cisco-IOS-XR-ipv4-bgp-cfg:bgp" { return "Bgp" }
    return ""
}

func (af *Vrfs_Vrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']" + "[saf-name='" + fmt.Sprintf("%v", af.SafName) + "']" + "[topology-name='" + fmt.Sprintf("%v", af.TopologyName) + "']"
}

func (af *Vrfs_Vrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-ip-rib-cfg:maximum-prefix" {
        return &af.MaximumPrefix
    }
    if childYangName == "Cisco-IOS-XR-ipv4-bgp-cfg:bgp" {
        return &af.Bgp
    }
    return nil
}

func (af *Vrfs_Vrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-ip-rib-cfg:maximum-prefix"] = &af.MaximumPrefix
    children["Cisco-IOS-XR-ipv4-bgp-cfg:bgp"] = &af.Bgp
    return children
}

func (af *Vrfs_Vrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["saf-name"] = af.SafName
    leafs["topology-name"] = af.TopologyName
    leafs["create"] = af.Create
    return leafs
}

func (af *Vrfs_Vrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *Vrfs_Vrf_Afs_Af) GetYangName() string { return "af" }

func (af *Vrfs_Vrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *Vrfs_Vrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *Vrfs_Vrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *Vrfs_Vrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *Vrfs_Vrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *Vrfs_Vrf_Afs_Af) GetParentYangName() string { return "afs" }

// Vrfs_Vrf_Afs_Af_MaximumPrefix
// Set maximum prefix limits
// This type is a presence type.
type Vrfs_Vrf_Afs_Af_MaximumPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set table's maximum prefix limit. The type is interface{} with range:
    // 32..5000000. This attribute is mandatory.
    PrefixLimit interface{}

    // Mid-threshold (% of maximum). The type is interface{} with range: 1..100.
    MidThreshold interface{}
}

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetFilter() yfilter.YFilter { return maximumPrefix.YFilter }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) SetFilter(yf yfilter.YFilter) { maximumPrefix.YFilter = yf }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    if yname == "mid-threshold" { return "MidThreshold" }
    return ""
}

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-rib-cfg:maximum-prefix"
}

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-limit"] = maximumPrefix.PrefixLimit
    leafs["mid-threshold"] = maximumPrefix.MidThreshold
    return leafs
}

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetYangName() string { return "maximum-prefix" }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) SetParent(parent types.Entity) { maximumPrefix.parent = parent }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetParent() types.Entity { return maximumPrefix.parent }

func (maximumPrefix *Vrfs_Vrf_Afs_Af_MaximumPrefix) GetParentYangName() string { return "af" }

// Vrfs_Vrf_Afs_Af_Bgp
// BGP AF VRF config
type Vrfs_Vrf_Afs_Af_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy for export filtering. The type is string.
    ExportRoutePolicy interface{}

    // Route policy for import filtering. The type is string.
    ImportRoutePolicy interface{}

    // TRUE Enable advertising imported paths to PEsFALSE Disable advertising
    // imported paths to PEs. The type is bool.
    ImportVrfOptions interface{}

    // Import Route targets.
    ImportRouteTargets Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets

    // Export Route targets.
    ExportRouteTargets Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets

    // Route policy for vrf to global export filtering.
    VrfToGlobalExportRoutePolicy Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy

    // Export VRF options.
    ExportVrfOptions Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions

    // Route policy for global to vrf import filtering.
    GlobalToVrfImportRoutePolicy Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy
}

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetGoName(yname string) string {
    if yname == "export-route-policy" { return "ExportRoutePolicy" }
    if yname == "import-route-policy" { return "ImportRoutePolicy" }
    if yname == "import-vrf-options" { return "ImportVrfOptions" }
    if yname == "import-route-targets" { return "ImportRouteTargets" }
    if yname == "export-route-targets" { return "ExportRouteTargets" }
    if yname == "vrf-to-global-export-route-policy" { return "VrfToGlobalExportRoutePolicy" }
    if yname == "export-vrf-options" { return "ExportVrfOptions" }
    if yname == "global-to-vrf-import-route-policy" { return "GlobalToVrfImportRoutePolicy" }
    return ""
}

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-bgp-cfg:bgp"
}

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "import-route-targets" {
        return &bgp.ImportRouteTargets
    }
    if childYangName == "export-route-targets" {
        return &bgp.ExportRouteTargets
    }
    if childYangName == "vrf-to-global-export-route-policy" {
        return &bgp.VrfToGlobalExportRoutePolicy
    }
    if childYangName == "export-vrf-options" {
        return &bgp.ExportVrfOptions
    }
    if childYangName == "global-to-vrf-import-route-policy" {
        return &bgp.GlobalToVrfImportRoutePolicy
    }
    return nil
}

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["import-route-targets"] = &bgp.ImportRouteTargets
    children["export-route-targets"] = &bgp.ExportRouteTargets
    children["vrf-to-global-export-route-policy"] = &bgp.VrfToGlobalExportRoutePolicy
    children["export-vrf-options"] = &bgp.ExportVrfOptions
    children["global-to-vrf-import-route-policy"] = &bgp.GlobalToVrfImportRoutePolicy
    return children
}

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["export-route-policy"] = bgp.ExportRoutePolicy
    leafs["import-route-policy"] = bgp.ImportRoutePolicy
    leafs["import-vrf-options"] = bgp.ImportVrfOptions
    return leafs
}

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetYangName() string { return "bgp" }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Vrfs_Vrf_Afs_Af_Bgp) GetParentYangName() string { return "af" }

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets
// Import Route targets
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route target table.
    RouteTargets Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets
}

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetFilter() yfilter.YFilter { return importRouteTargets.YFilter }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) SetFilter(yf yfilter.YFilter) { importRouteTargets.YFilter = yf }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetGoName(yname string) string {
    if yname == "route-targets" { return "RouteTargets" }
    return ""
}

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetSegmentPath() string {
    return "import-route-targets"
}

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-targets" {
        return &importRouteTargets.RouteTargets
    }
    return nil
}

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-targets"] = &importRouteTargets.RouteTargets
    return children
}

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetBundleName() string { return "cisco_ios_xr" }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetYangName() string { return "import-route-targets" }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) SetParent(parent types.Entity) { importRouteTargets.parent = parent }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetParent() types.Entity { return importRouteTargets.parent }

func (importRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets) GetParentYangName() string { return "bgp" }

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets
// Route target table
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route target. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget.
    RouteTarget []Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetFilter() yfilter.YFilter { return routeTargets.YFilter }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) SetFilter(yf yfilter.YFilter) { routeTargets.YFilter = yf }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetGoName(yname string) string {
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetSegmentPath() string {
    return "route-targets"
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        for _, c := range routeTargets.RouteTarget {
            if routeTargets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget{}
        routeTargets.RouteTarget = append(routeTargets.RouteTarget, child)
        return &routeTargets.RouteTarget[len(routeTargets.RouteTarget)-1]
    }
    return nil
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routeTargets.RouteTarget {
        children[routeTargets.RouteTarget[i].GetSegmentPath()] = &routeTargets.RouteTarget[i]
    }
    return children
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetBundleName() string { return "cisco_ios_xr" }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetYangName() string { return "route-targets" }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) SetParent(parent types.Entity) { routeTargets.parent = parent }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetParent() types.Entity { return routeTargets.parent }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets) GetParentYangName() string { return "import-route-targets" }

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget
// Route target
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type of RT. The type is BgpVrfRouteTarget.
    Type interface{}

    // as or four byte as. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs.
    AsOrFourByteAs []Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs

    // ipv4 address. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "as-or-four-byte-as" { return "AsOrFourByteAs" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    return ""
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetSegmentPath() string {
    return "route-target" + "[type='" + fmt.Sprintf("%v", routeTarget.Type) + "']"
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "as-or-four-byte-as" {
        for _, c := range routeTarget.AsOrFourByteAs {
            if routeTarget.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs{}
        routeTarget.AsOrFourByteAs = append(routeTarget.AsOrFourByteAs, child)
        return &routeTarget.AsOrFourByteAs[len(routeTarget.AsOrFourByteAs)-1]
    }
    if childYangName == "ipv4-address" {
        for _, c := range routeTarget.Ipv4Address {
            if routeTarget.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address{}
        routeTarget.Ipv4Address = append(routeTarget.Ipv4Address, child)
        return &routeTarget.Ipv4Address[len(routeTarget.Ipv4Address)-1]
    }
    return nil
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routeTarget.AsOrFourByteAs {
        children[routeTarget.AsOrFourByteAs[i].GetSegmentPath()] = &routeTarget.AsOrFourByteAs[i]
    }
    for i := range routeTarget.Ipv4Address {
        children[routeTarget.Ipv4Address[i].GetSegmentPath()] = &routeTarget.Ipv4Address[i]
    }
    return children
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = routeTarget.Type
    return leafs
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget) GetParentYangName() string { return "route-targets" }

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs
// as or four byte as
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. AS number. The type is interface{} with range:
    // 0..4294967295.
    AsXx interface{}

    // This attribute is a key. AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // This attribute is a key. AS number Index. The type is interface{} with
    // range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetFilter() yfilter.YFilter { return asOrFourByteAs.YFilter }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) SetFilter(yf yfilter.YFilter) { asOrFourByteAs.YFilter = yf }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as" { return "As" }
    if yname == "as-index" { return "AsIndex" }
    if yname == "stitching-rt" { return "StitchingRt" }
    return ""
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetSegmentPath() string {
    return "as-or-four-byte-as" + "[as-xx='" + fmt.Sprintf("%v", asOrFourByteAs.AsXx) + "']" + "[as='" + fmt.Sprintf("%v", asOrFourByteAs.As) + "']" + "[as-index='" + fmt.Sprintf("%v", asOrFourByteAs.AsIndex) + "']" + "[stitching-rt='" + fmt.Sprintf("%v", asOrFourByteAs.StitchingRt) + "']"
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = asOrFourByteAs.AsXx
    leafs["as"] = asOrFourByteAs.As
    leafs["as-index"] = asOrFourByteAs.AsIndex
    leafs["stitching-rt"] = asOrFourByteAs.StitchingRt
    return leafs
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetYangName() string { return "as-or-four-byte-as" }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) SetParent(parent types.Entity) { asOrFourByteAs.parent = parent }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetParent() types.Entity { return asOrFourByteAs.parent }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetParentYangName() string { return "route-target" }

// Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address Index. The type is interface{} with
    // range: 0..65535.
    AddressIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetFilter() yfilter.YFilter { return ipv4Address.YFilter }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) SetFilter(yf yfilter.YFilter) { ipv4Address.YFilter = yf }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "address-index" { return "AddressIndex" }
    if yname == "stitching-rt" { return "StitchingRt" }
    return ""
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetSegmentPath() string {
    return "ipv4-address" + "[address='" + fmt.Sprintf("%v", ipv4Address.Address) + "']" + "[address-index='" + fmt.Sprintf("%v", ipv4Address.AddressIndex) + "']" + "[stitching-rt='" + fmt.Sprintf("%v", ipv4Address.StitchingRt) + "']"
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv4Address.Address
    leafs["address-index"] = ipv4Address.AddressIndex
    leafs["stitching-rt"] = ipv4Address.StitchingRt
    return leafs
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetYangName() string { return "ipv4-address" }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) SetParent(parent types.Entity) { ipv4Address.parent = parent }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetParent() types.Entity { return ipv4Address.parent }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ImportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetParentYangName() string { return "route-target" }

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets
// Export Route targets
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route target table.
    RouteTargets Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets
}

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetFilter() yfilter.YFilter { return exportRouteTargets.YFilter }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) SetFilter(yf yfilter.YFilter) { exportRouteTargets.YFilter = yf }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetGoName(yname string) string {
    if yname == "route-targets" { return "RouteTargets" }
    return ""
}

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetSegmentPath() string {
    return "export-route-targets"
}

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-targets" {
        return &exportRouteTargets.RouteTargets
    }
    return nil
}

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-targets"] = &exportRouteTargets.RouteTargets
    return children
}

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetBundleName() string { return "cisco_ios_xr" }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetYangName() string { return "export-route-targets" }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) SetParent(parent types.Entity) { exportRouteTargets.parent = parent }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetParent() types.Entity { return exportRouteTargets.parent }

func (exportRouteTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets) GetParentYangName() string { return "bgp" }

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets
// Route target table
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route target. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget.
    RouteTarget []Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetFilter() yfilter.YFilter { return routeTargets.YFilter }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) SetFilter(yf yfilter.YFilter) { routeTargets.YFilter = yf }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetGoName(yname string) string {
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetSegmentPath() string {
    return "route-targets"
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        for _, c := range routeTargets.RouteTarget {
            if routeTargets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget{}
        routeTargets.RouteTarget = append(routeTargets.RouteTarget, child)
        return &routeTargets.RouteTarget[len(routeTargets.RouteTarget)-1]
    }
    return nil
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routeTargets.RouteTarget {
        children[routeTargets.RouteTarget[i].GetSegmentPath()] = &routeTargets.RouteTarget[i]
    }
    return children
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetBundleName() string { return "cisco_ios_xr" }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetYangName() string { return "route-targets" }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) SetParent(parent types.Entity) { routeTargets.parent = parent }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetParent() types.Entity { return routeTargets.parent }

func (routeTargets *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets) GetParentYangName() string { return "export-route-targets" }

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget
// Route target
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type of RT. The type is BgpVrfRouteTarget.
    Type interface{}

    // as or four byte as. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs.
    AsOrFourByteAs []Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs

    // ipv4 address. The type is slice of
    // Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "as-or-four-byte-as" { return "AsOrFourByteAs" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    return ""
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetSegmentPath() string {
    return "route-target" + "[type='" + fmt.Sprintf("%v", routeTarget.Type) + "']"
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "as-or-four-byte-as" {
        for _, c := range routeTarget.AsOrFourByteAs {
            if routeTarget.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs{}
        routeTarget.AsOrFourByteAs = append(routeTarget.AsOrFourByteAs, child)
        return &routeTarget.AsOrFourByteAs[len(routeTarget.AsOrFourByteAs)-1]
    }
    if childYangName == "ipv4-address" {
        for _, c := range routeTarget.Ipv4Address {
            if routeTarget.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address{}
        routeTarget.Ipv4Address = append(routeTarget.Ipv4Address, child)
        return &routeTarget.Ipv4Address[len(routeTarget.Ipv4Address)-1]
    }
    return nil
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routeTarget.AsOrFourByteAs {
        children[routeTarget.AsOrFourByteAs[i].GetSegmentPath()] = &routeTarget.AsOrFourByteAs[i]
    }
    for i := range routeTarget.Ipv4Address {
        children[routeTarget.Ipv4Address[i].GetSegmentPath()] = &routeTarget.Ipv4Address[i]
    }
    return children
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = routeTarget.Type
    return leafs
}

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget) GetParentYangName() string { return "route-targets" }

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs
// as or four byte as
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. AS number. The type is interface{} with range:
    // 0..4294967295.
    AsXx interface{}

    // This attribute is a key. AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // This attribute is a key. AS number Index. The type is interface{} with
    // range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetFilter() yfilter.YFilter { return asOrFourByteAs.YFilter }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) SetFilter(yf yfilter.YFilter) { asOrFourByteAs.YFilter = yf }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as" { return "As" }
    if yname == "as-index" { return "AsIndex" }
    if yname == "stitching-rt" { return "StitchingRt" }
    return ""
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetSegmentPath() string {
    return "as-or-four-byte-as" + "[as-xx='" + fmt.Sprintf("%v", asOrFourByteAs.AsXx) + "']" + "[as='" + fmt.Sprintf("%v", asOrFourByteAs.As) + "']" + "[as-index='" + fmt.Sprintf("%v", asOrFourByteAs.AsIndex) + "']" + "[stitching-rt='" + fmt.Sprintf("%v", asOrFourByteAs.StitchingRt) + "']"
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = asOrFourByteAs.AsXx
    leafs["as"] = asOrFourByteAs.As
    leafs["as-index"] = asOrFourByteAs.AsIndex
    leafs["stitching-rt"] = asOrFourByteAs.StitchingRt
    return leafs
}

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetYangName() string { return "as-or-four-byte-as" }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) SetParent(parent types.Entity) { asOrFourByteAs.parent = parent }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetParent() types.Entity { return asOrFourByteAs.parent }

func (asOrFourByteAs *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_AsOrFourByteAs) GetParentYangName() string { return "route-target" }

// Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IP address Index. The type is interface{} with
    // range: 0..65535.
    AddressIndex interface{}

    // This attribute is a key. Stitching RT. The type is interface{} with range:
    // 0..1.
    StitchingRt interface{}
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetFilter() yfilter.YFilter { return ipv4Address.YFilter }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) SetFilter(yf yfilter.YFilter) { ipv4Address.YFilter = yf }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "address-index" { return "AddressIndex" }
    if yname == "stitching-rt" { return "StitchingRt" }
    return ""
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetSegmentPath() string {
    return "ipv4-address" + "[address='" + fmt.Sprintf("%v", ipv4Address.Address) + "']" + "[address-index='" + fmt.Sprintf("%v", ipv4Address.AddressIndex) + "']" + "[stitching-rt='" + fmt.Sprintf("%v", ipv4Address.StitchingRt) + "']"
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv4Address.Address
    leafs["address-index"] = ipv4Address.AddressIndex
    leafs["stitching-rt"] = ipv4Address.StitchingRt
    return leafs
}

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetYangName() string { return "ipv4-address" }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) SetParent(parent types.Entity) { ipv4Address.parent = parent }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetParent() types.Entity { return ipv4Address.parent }

func (ipv4Address *Vrfs_Vrf_Afs_Af_Bgp_ExportRouteTargets_RouteTargets_RouteTarget_Ipv4Address) GetParentYangName() string { return "route-target" }

// Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy
// Route policy for vrf to global export filtering
// This type is a presence type.
type Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Vrf to global export route policy. The type is string. This attribute is
    // mandatory.
    RoutePolicyName interface{}

    // TRUE Enable imported VPN paths to be exported to Default VRF.FALSE Disable
    // imported VPN paths to be exported to Default VRF. The type is bool.
    AllowImportedVpn interface{}
}

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetFilter() yfilter.YFilter { return vrfToGlobalExportRoutePolicy.YFilter }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) SetFilter(yf yfilter.YFilter) { vrfToGlobalExportRoutePolicy.YFilter = yf }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "allow-imported-vpn" { return "AllowImportedVpn" }
    return ""
}

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetSegmentPath() string {
    return "vrf-to-global-export-route-policy"
}

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = vrfToGlobalExportRoutePolicy.RoutePolicyName
    leafs["allow-imported-vpn"] = vrfToGlobalExportRoutePolicy.AllowImportedVpn
    return leafs
}

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetYangName() string { return "vrf-to-global-export-route-policy" }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) SetParent(parent types.Entity) { vrfToGlobalExportRoutePolicy.parent = parent }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetParent() types.Entity { return vrfToGlobalExportRoutePolicy.parent }

func (vrfToGlobalExportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_VrfToGlobalExportRoutePolicy) GetParentYangName() string { return "bgp" }

// Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions
// Export VRF options
type Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE Enable imported VPN paths to be exported to non-default VRFFALSE
    // Disable imported VPN paths to be exported to non-default VRF. The type is
    // bool.
    AllowImportedVpn interface{}

    // TRUE Use stitchng RTs to import extranet pathsFALSE Use regular RTs to
    // import extranet paths. The type is bool.
    ImportStitchingRt interface{}
}

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetFilter() yfilter.YFilter { return exportVrfOptions.YFilter }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) SetFilter(yf yfilter.YFilter) { exportVrfOptions.YFilter = yf }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetGoName(yname string) string {
    if yname == "allow-imported-vpn" { return "AllowImportedVpn" }
    if yname == "import-stitching-rt" { return "ImportStitchingRt" }
    return ""
}

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetSegmentPath() string {
    return "export-vrf-options"
}

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-imported-vpn"] = exportVrfOptions.AllowImportedVpn
    leafs["import-stitching-rt"] = exportVrfOptions.ImportStitchingRt
    return leafs
}

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetBundleName() string { return "cisco_ios_xr" }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetYangName() string { return "export-vrf-options" }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) SetParent(parent types.Entity) { exportVrfOptions.parent = parent }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetParent() types.Entity { return exportVrfOptions.parent }

func (exportVrfOptions *Vrfs_Vrf_Afs_Af_Bgp_ExportVrfOptions) GetParentYangName() string { return "bgp" }

// Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy
// Route policy for global to vrf import filtering
// This type is a presence type.
type Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global to vrf import route policy. The type is string. This attribute is
    // mandatory.
    RoutePolicyName interface{}

    // TRUE Enable advertising imported paths to PEsFALSE Disable advertising
    // imported paths to PEs. The type is bool.
    AdvertiseAsVpn interface{}
}

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetFilter() yfilter.YFilter { return globalToVrfImportRoutePolicy.YFilter }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) SetFilter(yf yfilter.YFilter) { globalToVrfImportRoutePolicy.YFilter = yf }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetGoName(yname string) string {
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "advertise-as-vpn" { return "AdvertiseAsVpn" }
    return ""
}

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetSegmentPath() string {
    return "global-to-vrf-import-route-policy"
}

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy-name"] = globalToVrfImportRoutePolicy.RoutePolicyName
    leafs["advertise-as-vpn"] = globalToVrfImportRoutePolicy.AdvertiseAsVpn
    return leafs
}

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetYangName() string { return "global-to-vrf-import-route-policy" }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) SetParent(parent types.Entity) { globalToVrfImportRoutePolicy.parent = parent }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetParent() types.Entity { return globalToVrfImportRoutePolicy.parent }

func (globalToVrfImportRoutePolicy *Vrfs_Vrf_Afs_Af_Bgp_GlobalToVrfImportRoutePolicy) GetParentYangName() string { return "bgp" }

// Vrfs_Vrf_MulticastHost
// Multicast host stack configuration
type Vrfs_Vrf_MulticastHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 configuration.
    Ipv4 Vrfs_Vrf_MulticastHost_Ipv4

    // IPv6 configuration.
    Ipv6 Vrfs_Vrf_MulticastHost_Ipv6
}

func (multicastHost *Vrfs_Vrf_MulticastHost) GetFilter() yfilter.YFilter { return multicastHost.YFilter }

func (multicastHost *Vrfs_Vrf_MulticastHost) SetFilter(yf yfilter.YFilter) { multicastHost.YFilter = yf }

func (multicastHost *Vrfs_Vrf_MulticastHost) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (multicastHost *Vrfs_Vrf_MulticastHost) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-iarm-vrf-cfg:multicast-host"
}

func (multicastHost *Vrfs_Vrf_MulticastHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &multicastHost.Ipv4
    }
    if childYangName == "ipv6" {
        return &multicastHost.Ipv6
    }
    return nil
}

func (multicastHost *Vrfs_Vrf_MulticastHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &multicastHost.Ipv4
    children["ipv6"] = &multicastHost.Ipv6
    return children
}

func (multicastHost *Vrfs_Vrf_MulticastHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (multicastHost *Vrfs_Vrf_MulticastHost) GetBundleName() string { return "cisco_ios_xr" }

func (multicastHost *Vrfs_Vrf_MulticastHost) GetYangName() string { return "multicast-host" }

func (multicastHost *Vrfs_Vrf_MulticastHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multicastHost *Vrfs_Vrf_MulticastHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multicastHost *Vrfs_Vrf_MulticastHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multicastHost *Vrfs_Vrf_MulticastHost) SetParent(parent types.Entity) { multicastHost.parent = parent }

func (multicastHost *Vrfs_Vrf_MulticastHost) GetParent() types.Entity { return multicastHost.parent }

func (multicastHost *Vrfs_Vrf_MulticastHost) GetParentYangName() string { return "vrf" }

// Vrfs_Vrf_MulticastHost_Ipv4
// IPv4 configuration
type Vrfs_Vrf_MulticastHost_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = ipv4.Interface
    return leafs
}

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Vrfs_Vrf_MulticastHost_Ipv4) GetParentYangName() string { return "multicast-host" }

// Vrfs_Vrf_MulticastHost_Ipv6
// IPv6 configuration
type Vrfs_Vrf_MulticastHost_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default multicast host interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = ipv6.Interface
    return leafs
}

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Vrfs_Vrf_MulticastHost_Ipv6) GetParentYangName() string { return "multicast-host" }

// GlobalAf
// global af
type GlobalAf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF address family configuration.
    Afs GlobalAf_Afs
}

func (globalAf *GlobalAf) GetFilter() yfilter.YFilter { return globalAf.YFilter }

func (globalAf *GlobalAf) SetFilter(yf yfilter.YFilter) { globalAf.YFilter = yf }

func (globalAf *GlobalAf) GetGoName(yname string) string {
    if yname == "afs" { return "Afs" }
    return ""
}

func (globalAf *GlobalAf) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-cfg:global-af"
}

func (globalAf *GlobalAf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &globalAf.Afs
    }
    return nil
}

func (globalAf *GlobalAf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &globalAf.Afs
    return children
}

func (globalAf *GlobalAf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalAf *GlobalAf) GetBundleName() string { return "cisco_ios_xr" }

func (globalAf *GlobalAf) GetYangName() string { return "global-af" }

func (globalAf *GlobalAf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalAf *GlobalAf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalAf *GlobalAf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalAf *GlobalAf) SetParent(parent types.Entity) { globalAf.parent = parent }

func (globalAf *GlobalAf) GetParent() types.Entity { return globalAf.parent }

func (globalAf *GlobalAf) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-cfg" }

// GlobalAf_Afs
// VRF address family configuration
type GlobalAf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF address family configuration. The type is slice of GlobalAf_Afs_Af.
    Af []GlobalAf_Afs_Af
}

func (afs *GlobalAf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *GlobalAf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *GlobalAf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *GlobalAf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *GlobalAf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := GlobalAf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *GlobalAf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *GlobalAf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *GlobalAf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *GlobalAf_Afs) GetYangName() string { return "afs" }

func (afs *GlobalAf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *GlobalAf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *GlobalAf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *GlobalAf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *GlobalAf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *GlobalAf_Afs) GetParentYangName() string { return "global-af" }

// GlobalAf_Afs_Af
// VRF address family configuration
type GlobalAf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is VrfAddressFamily.
    AfName interface{}

    // This attribute is a key. Sub-Address family. The type is
    // VrfSubAddressFamily.
    SafName interface{}

    // This attribute is a key. Topology name. The type is string with length:
    // 1..244.
    TopologyName interface{}

    // VRF configuration for a particular address family. The type is interface{}.
    Create interface{}
}

func (af *GlobalAf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *GlobalAf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *GlobalAf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "topology-name" { return "TopologyName" }
    if yname == "create" { return "Create" }
    return ""
}

func (af *GlobalAf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']" + "[saf-name='" + fmt.Sprintf("%v", af.SafName) + "']" + "[topology-name='" + fmt.Sprintf("%v", af.TopologyName) + "']"
}

func (af *GlobalAf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (af *GlobalAf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (af *GlobalAf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["saf-name"] = af.SafName
    leafs["topology-name"] = af.TopologyName
    leafs["create"] = af.Create
    return leafs
}

func (af *GlobalAf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *GlobalAf_Afs_Af) GetYangName() string { return "af" }

func (af *GlobalAf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *GlobalAf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *GlobalAf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *GlobalAf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *GlobalAf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *GlobalAf_Afs_Af) GetParentYangName() string { return "afs" }

// Srlg
// srlg
type Srlg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable SRLG. The type is interface{}.
    Enable interface{}

    // Set of interfaces configured with SRLG.
    Interfaces Srlg_Interfaces

    // Set of SRLG name configuration.
    SrlgNames Srlg_SrlgNames

    // Set of groups configured with SRLG.
    Groups Srlg_Groups

    // Set of inherit nodes configured with SRLG.
    InheritNodes Srlg_InheritNodes
}

func (srlg *Srlg) GetFilter() yfilter.YFilter { return srlg.YFilter }

func (srlg *Srlg) SetFilter(yf yfilter.YFilter) { srlg.YFilter = yf }

func (srlg *Srlg) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "srlg-names" { return "SrlgNames" }
    if yname == "groups" { return "Groups" }
    if yname == "inherit-nodes" { return "InheritNodes" }
    return ""
}

func (srlg *Srlg) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-cfg:srlg"
}

func (srlg *Srlg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &srlg.Interfaces
    }
    if childYangName == "srlg-names" {
        return &srlg.SrlgNames
    }
    if childYangName == "groups" {
        return &srlg.Groups
    }
    if childYangName == "inherit-nodes" {
        return &srlg.InheritNodes
    }
    return nil
}

func (srlg *Srlg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &srlg.Interfaces
    children["srlg-names"] = &srlg.SrlgNames
    children["groups"] = &srlg.Groups
    children["inherit-nodes"] = &srlg.InheritNodes
    return children
}

func (srlg *Srlg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = srlg.Enable
    return leafs
}

func (srlg *Srlg) GetBundleName() string { return "cisco_ios_xr" }

func (srlg *Srlg) GetYangName() string { return "srlg" }

func (srlg *Srlg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlg *Srlg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlg *Srlg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlg *Srlg) SetParent(parent types.Entity) { srlg.parent = parent }

func (srlg *Srlg) GetParent() types.Entity { return srlg.parent }

func (srlg *Srlg) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-cfg" }

// Srlg_Interfaces
// Set of interfaces configured with SRLG
type Srlg_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface configurations. The type is slice of Srlg_Interfaces_Interface.
    Interface []Srlg_Interfaces_Interface
}

func (interfaces *Srlg_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Srlg_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Srlg_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Srlg_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Srlg_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Srlg_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Srlg_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Srlg_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Srlg_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Srlg_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Srlg_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Srlg_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Srlg_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Srlg_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Srlg_Interfaces) GetParentYangName() string { return "srlg" }

// Srlg_Interfaces_Interface
// Interface configurations
type Srlg_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enable SRLG interface. The type is interface{}.
    Enable interface{}

    // Include optical configuration for an interface.
    IncludeOptical Srlg_Interfaces_Interface_IncludeOptical

    // Group configuration for an interface.
    InterfaceGroup Srlg_Interfaces_Interface_InterfaceGroup

    // SRLG Value configuration for an interface.
    Values Srlg_Interfaces_Interface_Values

    // SRLG Name configuration for an interface.
    InterfaceSrlgNames Srlg_Interfaces_Interface_InterfaceSrlgNames
}

func (self *Srlg_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Srlg_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Srlg_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
    if yname == "include-optical" { return "IncludeOptical" }
    if yname == "interface-group" { return "InterfaceGroup" }
    if yname == "values" { return "Values" }
    if yname == "interface-srlg-names" { return "InterfaceSrlgNames" }
    return ""
}

func (self *Srlg_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Srlg_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "include-optical" {
        return &self.IncludeOptical
    }
    if childYangName == "interface-group" {
        return &self.InterfaceGroup
    }
    if childYangName == "values" {
        return &self.Values
    }
    if childYangName == "interface-srlg-names" {
        return &self.InterfaceSrlgNames
    }
    return nil
}

func (self *Srlg_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["include-optical"] = &self.IncludeOptical
    children["interface-group"] = &self.InterfaceGroup
    children["values"] = &self.Values
    children["interface-srlg-names"] = &self.InterfaceSrlgNames
    return children
}

func (self *Srlg_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["enable"] = self.Enable
    return leafs
}

func (self *Srlg_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Srlg_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Srlg_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Srlg_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Srlg_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Srlg_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Srlg_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Srlg_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Srlg_Interfaces_Interface_IncludeOptical
// Include optical configuration for an interface
type Srlg_Interfaces_Interface_IncludeOptical struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable SRLG interface include optical. The type is interface{}.
    Enable interface{}

    // Priority for optical domain values. The type is SrlgPriority. The default
    // value is default.
    Priority interface{}
}

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetFilter() yfilter.YFilter { return includeOptical.YFilter }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) SetFilter(yf yfilter.YFilter) { includeOptical.YFilter = yf }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetSegmentPath() string {
    return "include-optical"
}

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = includeOptical.Enable
    leafs["priority"] = includeOptical.Priority
    return leafs
}

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetBundleName() string { return "cisco_ios_xr" }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetYangName() string { return "include-optical" }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) SetParent(parent types.Entity) { includeOptical.parent = parent }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetParent() types.Entity { return includeOptical.parent }

func (includeOptical *Srlg_Interfaces_Interface_IncludeOptical) GetParentYangName() string { return "interface" }

// Srlg_Interfaces_Interface_InterfaceGroup
// Group configuration for an interface
type Srlg_Interfaces_Interface_InterfaceGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable SRLG interface group submode. The type is interface{}.
    Enable interface{}

    // Set of group name under an interface.
    GroupNames Srlg_Interfaces_Interface_InterfaceGroup_GroupNames
}

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetFilter() yfilter.YFilter { return interfaceGroup.YFilter }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) SetFilter(yf yfilter.YFilter) { interfaceGroup.YFilter = yf }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "group-names" { return "GroupNames" }
    return ""
}

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetSegmentPath() string {
    return "interface-group"
}

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-names" {
        return &interfaceGroup.GroupNames
    }
    return nil
}

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-names"] = &interfaceGroup.GroupNames
    return children
}

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = interfaceGroup.Enable
    return leafs
}

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetYangName() string { return "interface-group" }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) SetParent(parent types.Entity) { interfaceGroup.parent = parent }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetParent() types.Entity { return interfaceGroup.parent }

func (interfaceGroup *Srlg_Interfaces_Interface_InterfaceGroup) GetParentYangName() string { return "interface" }

// Srlg_Interfaces_Interface_InterfaceGroup_GroupNames
// Set of group name under an interface
type Srlg_Interfaces_Interface_InterfaceGroup_GroupNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group name included under interface. The type is slice of
    // Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName.
    GroupName []Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName
}

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetFilter() yfilter.YFilter { return groupNames.YFilter }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) SetFilter(yf yfilter.YFilter) { groupNames.YFilter = yf }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetGoName(yname string) string {
    if yname == "group-name" { return "GroupName" }
    return ""
}

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetSegmentPath() string {
    return "group-names"
}

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-name" {
        for _, c := range groupNames.GroupName {
            if groupNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName{}
        groupNames.GroupName = append(groupNames.GroupName, child)
        return &groupNames.GroupName[len(groupNames.GroupName)-1]
    }
    return nil
}

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupNames.GroupName {
        children[groupNames.GroupName[i].GetSegmentPath()] = &groupNames.GroupName[i]
    }
    return children
}

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetBundleName() string { return "cisco_ios_xr" }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetYangName() string { return "group-names" }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) SetParent(parent types.Entity) { groupNames.parent = parent }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetParent() types.Entity { return groupNames.parent }

func (groupNames *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames) GetParentYangName() string { return "interface-group" }

// Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName
// Group name included under interface
type Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group name index. The type is interface{} with
    // range: 0..4294967295.
    GroupNameIndex interface{}

    // Group name. The type is string. This attribute is mandatory.
    GroupName interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetFilter() yfilter.YFilter { return groupName.YFilter }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) SetFilter(yf yfilter.YFilter) { groupName.YFilter = yf }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetGoName(yname string) string {
    if yname == "group-name-index" { return "GroupNameIndex" }
    if yname == "group-name" { return "GroupName" }
    if yname == "srlg-priority" { return "SrlgPriority" }
    return ""
}

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetSegmentPath() string {
    return "group-name" + "[group-name-index='" + fmt.Sprintf("%v", groupName.GroupNameIndex) + "']"
}

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-name-index"] = groupName.GroupNameIndex
    leafs["group-name"] = groupName.GroupName
    leafs["srlg-priority"] = groupName.SrlgPriority
    return leafs
}

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetBundleName() string { return "cisco_ios_xr" }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetYangName() string { return "group-name" }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) SetParent(parent types.Entity) { groupName.parent = parent }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetParent() types.Entity { return groupName.parent }

func (groupName *Srlg_Interfaces_Interface_InterfaceGroup_GroupNames_GroupName) GetParentYangName() string { return "group-names" }

// Srlg_Interfaces_Interface_Values
// SRLG Value configuration for an interface
type Srlg_Interfaces_Interface_Values struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG value data. The type is slice of
    // Srlg_Interfaces_Interface_Values_Value.
    Value []Srlg_Interfaces_Interface_Values_Value
}

func (values *Srlg_Interfaces_Interface_Values) GetFilter() yfilter.YFilter { return values.YFilter }

func (values *Srlg_Interfaces_Interface_Values) SetFilter(yf yfilter.YFilter) { values.YFilter = yf }

func (values *Srlg_Interfaces_Interface_Values) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    return ""
}

func (values *Srlg_Interfaces_Interface_Values) GetSegmentPath() string {
    return "values"
}

func (values *Srlg_Interfaces_Interface_Values) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "value" {
        for _, c := range values.Value {
            if values.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Interfaces_Interface_Values_Value{}
        values.Value = append(values.Value, child)
        return &values.Value[len(values.Value)-1]
    }
    return nil
}

func (values *Srlg_Interfaces_Interface_Values) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range values.Value {
        children[values.Value[i].GetSegmentPath()] = &values.Value[i]
    }
    return children
}

func (values *Srlg_Interfaces_Interface_Values) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (values *Srlg_Interfaces_Interface_Values) GetBundleName() string { return "cisco_ios_xr" }

func (values *Srlg_Interfaces_Interface_Values) GetYangName() string { return "values" }

func (values *Srlg_Interfaces_Interface_Values) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (values *Srlg_Interfaces_Interface_Values) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (values *Srlg_Interfaces_Interface_Values) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (values *Srlg_Interfaces_Interface_Values) SetParent(parent types.Entity) { values.parent = parent }

func (values *Srlg_Interfaces_Interface_Values) GetParent() types.Entity { return values.parent }

func (values *Srlg_Interfaces_Interface_Values) GetParentYangName() string { return "interface" }

// Srlg_Interfaces_Interface_Values_Value
// SRLG value data
type Srlg_Interfaces_Interface_Values_Value struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG index. The type is interface{} with range:
    // 1..65535.
    SrlgIndex interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (value *Srlg_Interfaces_Interface_Values_Value) GetFilter() yfilter.YFilter { return value.YFilter }

func (value *Srlg_Interfaces_Interface_Values_Value) SetFilter(yf yfilter.YFilter) { value.YFilter = yf }

func (value *Srlg_Interfaces_Interface_Values_Value) GetGoName(yname string) string {
    if yname == "srlg-index" { return "SrlgIndex" }
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "srlg-priority" { return "SrlgPriority" }
    return ""
}

func (value *Srlg_Interfaces_Interface_Values_Value) GetSegmentPath() string {
    return "value" + "[srlg-index='" + fmt.Sprintf("%v", value.SrlgIndex) + "']"
}

func (value *Srlg_Interfaces_Interface_Values_Value) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (value *Srlg_Interfaces_Interface_Values_Value) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (value *Srlg_Interfaces_Interface_Values_Value) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-index"] = value.SrlgIndex
    leafs["srlg-value"] = value.SrlgValue
    leafs["srlg-priority"] = value.SrlgPriority
    return leafs
}

func (value *Srlg_Interfaces_Interface_Values_Value) GetBundleName() string { return "cisco_ios_xr" }

func (value *Srlg_Interfaces_Interface_Values_Value) GetYangName() string { return "value" }

func (value *Srlg_Interfaces_Interface_Values_Value) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (value *Srlg_Interfaces_Interface_Values_Value) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (value *Srlg_Interfaces_Interface_Values_Value) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (value *Srlg_Interfaces_Interface_Values_Value) SetParent(parent types.Entity) { value.parent = parent }

func (value *Srlg_Interfaces_Interface_Values_Value) GetParent() types.Entity { return value.parent }

func (value *Srlg_Interfaces_Interface_Values_Value) GetParentYangName() string { return "values" }

// Srlg_Interfaces_Interface_InterfaceSrlgNames
// SRLG Name configuration for an interface
type Srlg_Interfaces_Interface_InterfaceSrlgNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG name data. The type is slice of
    // Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName.
    InterfaceSrlgName []Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName
}

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetFilter() yfilter.YFilter { return interfaceSrlgNames.YFilter }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) SetFilter(yf yfilter.YFilter) { interfaceSrlgNames.YFilter = yf }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetGoName(yname string) string {
    if yname == "interface-srlg-name" { return "InterfaceSrlgName" }
    return ""
}

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetSegmentPath() string {
    return "interface-srlg-names"
}

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-srlg-name" {
        for _, c := range interfaceSrlgNames.InterfaceSrlgName {
            if interfaceSrlgNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName{}
        interfaceSrlgNames.InterfaceSrlgName = append(interfaceSrlgNames.InterfaceSrlgName, child)
        return &interfaceSrlgNames.InterfaceSrlgName[len(interfaceSrlgNames.InterfaceSrlgName)-1]
    }
    return nil
}

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceSrlgNames.InterfaceSrlgName {
        children[interfaceSrlgNames.InterfaceSrlgName[i].GetSegmentPath()] = &interfaceSrlgNames.InterfaceSrlgName[i]
    }
    return children
}

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetYangName() string { return "interface-srlg-names" }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) SetParent(parent types.Entity) { interfaceSrlgNames.parent = parent }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetParent() types.Entity { return interfaceSrlgNames.parent }

func (interfaceSrlgNames *Srlg_Interfaces_Interface_InterfaceSrlgNames) GetParentYangName() string { return "interface" }

// Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName
// SRLG name data
type Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}
}

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetFilter() yfilter.YFilter { return interfaceSrlgName.YFilter }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) SetFilter(yf yfilter.YFilter) { interfaceSrlgName.YFilter = yf }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetGoName(yname string) string {
    if yname == "srlg-name" { return "SrlgName" }
    return ""
}

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetSegmentPath() string {
    return "interface-srlg-name" + "[srlg-name='" + fmt.Sprintf("%v", interfaceSrlgName.SrlgName) + "']"
}

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-name"] = interfaceSrlgName.SrlgName
    return leafs
}

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetYangName() string { return "interface-srlg-name" }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) SetParent(parent types.Entity) { interfaceSrlgName.parent = parent }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetParent() types.Entity { return interfaceSrlgName.parent }

func (interfaceSrlgName *Srlg_Interfaces_Interface_InterfaceSrlgNames_InterfaceSrlgName) GetParentYangName() string { return "interface-srlg-names" }

// Srlg_SrlgNames
// Set of SRLG name configuration
type Srlg_SrlgNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRLG name configuration. The type is slice of Srlg_SrlgNames_SrlgName.
    SrlgName []Srlg_SrlgNames_SrlgName
}

func (srlgNames *Srlg_SrlgNames) GetFilter() yfilter.YFilter { return srlgNames.YFilter }

func (srlgNames *Srlg_SrlgNames) SetFilter(yf yfilter.YFilter) { srlgNames.YFilter = yf }

func (srlgNames *Srlg_SrlgNames) GetGoName(yname string) string {
    if yname == "srlg-name" { return "SrlgName" }
    return ""
}

func (srlgNames *Srlg_SrlgNames) GetSegmentPath() string {
    return "srlg-names"
}

func (srlgNames *Srlg_SrlgNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg-name" {
        for _, c := range srlgNames.SrlgName {
            if srlgNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_SrlgNames_SrlgName{}
        srlgNames.SrlgName = append(srlgNames.SrlgName, child)
        return &srlgNames.SrlgName[len(srlgNames.SrlgName)-1]
    }
    return nil
}

func (srlgNames *Srlg_SrlgNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range srlgNames.SrlgName {
        children[srlgNames.SrlgName[i].GetSegmentPath()] = &srlgNames.SrlgName[i]
    }
    return children
}

func (srlgNames *Srlg_SrlgNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srlgNames *Srlg_SrlgNames) GetBundleName() string { return "cisco_ios_xr" }

func (srlgNames *Srlg_SrlgNames) GetYangName() string { return "srlg-names" }

func (srlgNames *Srlg_SrlgNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgNames *Srlg_SrlgNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgNames *Srlg_SrlgNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgNames *Srlg_SrlgNames) SetParent(parent types.Entity) { srlgNames.parent = parent }

func (srlgNames *Srlg_SrlgNames) GetParent() types.Entity { return srlgNames.parent }

func (srlgNames *Srlg_SrlgNames) GetParentYangName() string { return "srlg" }

// Srlg_SrlgNames_SrlgName
// SRLG name configuration
type Srlg_SrlgNames_SrlgName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG name. The type is string with length: 1..64.
    SrlgName interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}
}

func (srlgName *Srlg_SrlgNames_SrlgName) GetFilter() yfilter.YFilter { return srlgName.YFilter }

func (srlgName *Srlg_SrlgNames_SrlgName) SetFilter(yf yfilter.YFilter) { srlgName.YFilter = yf }

func (srlgName *Srlg_SrlgNames_SrlgName) GetGoName(yname string) string {
    if yname == "srlg-name" { return "SrlgName" }
    if yname == "srlg-value" { return "SrlgValue" }
    return ""
}

func (srlgName *Srlg_SrlgNames_SrlgName) GetSegmentPath() string {
    return "srlg-name" + "[srlg-name='" + fmt.Sprintf("%v", srlgName.SrlgName) + "']"
}

func (srlgName *Srlg_SrlgNames_SrlgName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srlgName *Srlg_SrlgNames_SrlgName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srlgName *Srlg_SrlgNames_SrlgName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-name"] = srlgName.SrlgName
    leafs["srlg-value"] = srlgName.SrlgValue
    return leafs
}

func (srlgName *Srlg_SrlgNames_SrlgName) GetBundleName() string { return "cisco_ios_xr" }

func (srlgName *Srlg_SrlgNames_SrlgName) GetYangName() string { return "srlg-name" }

func (srlgName *Srlg_SrlgNames_SrlgName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srlgName *Srlg_SrlgNames_SrlgName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srlgName *Srlg_SrlgNames_SrlgName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srlgName *Srlg_SrlgNames_SrlgName) SetParent(parent types.Entity) { srlgName.parent = parent }

func (srlgName *Srlg_SrlgNames_SrlgName) GetParent() types.Entity { return srlgName.parent }

func (srlgName *Srlg_SrlgNames_SrlgName) GetParentYangName() string { return "srlg-names" }

// Srlg_Groups
// Set of groups configured with SRLG
type Srlg_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group configurations. The type is slice of Srlg_Groups_Group.
    Group []Srlg_Groups_Group
}

func (groups *Srlg_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Srlg_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Srlg_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Srlg_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Srlg_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Srlg_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Srlg_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Srlg_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Srlg_Groups) GetYangName() string { return "groups" }

func (groups *Srlg_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Srlg_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Srlg_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Srlg_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Srlg_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Srlg_Groups) GetParentYangName() string { return "srlg" }

// Srlg_Groups_Group
// Group configurations
type Srlg_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupName interface{}

    // Enable SRLG group. The type is interface{}.
    Enable interface{}

    // Set of SRLG values configured under a group.
    GroupValues Srlg_Groups_Group_GroupValues
}

func (group *Srlg_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Srlg_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Srlg_Groups_Group) GetGoName(yname string) string {
    if yname == "group-name" { return "GroupName" }
    if yname == "enable" { return "Enable" }
    if yname == "group-values" { return "GroupValues" }
    return ""
}

func (group *Srlg_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-name='" + fmt.Sprintf("%v", group.GroupName) + "']"
}

func (group *Srlg_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-values" {
        return &group.GroupValues
    }
    return nil
}

func (group *Srlg_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-values"] = &group.GroupValues
    return children
}

func (group *Srlg_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-name"] = group.GroupName
    leafs["enable"] = group.Enable
    return leafs
}

func (group *Srlg_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Srlg_Groups_Group) GetYangName() string { return "group" }

func (group *Srlg_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Srlg_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Srlg_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Srlg_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Srlg_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Srlg_Groups_Group) GetParentYangName() string { return "groups" }

// Srlg_Groups_Group_GroupValues
// Set of SRLG values configured under a group
type Srlg_Groups_Group_GroupValues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group SRLG values with attribute. The type is slice of
    // Srlg_Groups_Group_GroupValues_GroupValue.
    GroupValue []Srlg_Groups_Group_GroupValues_GroupValue
}

func (groupValues *Srlg_Groups_Group_GroupValues) GetFilter() yfilter.YFilter { return groupValues.YFilter }

func (groupValues *Srlg_Groups_Group_GroupValues) SetFilter(yf yfilter.YFilter) { groupValues.YFilter = yf }

func (groupValues *Srlg_Groups_Group_GroupValues) GetGoName(yname string) string {
    if yname == "group-value" { return "GroupValue" }
    return ""
}

func (groupValues *Srlg_Groups_Group_GroupValues) GetSegmentPath() string {
    return "group-values"
}

func (groupValues *Srlg_Groups_Group_GroupValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-value" {
        for _, c := range groupValues.GroupValue {
            if groupValues.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_Groups_Group_GroupValues_GroupValue{}
        groupValues.GroupValue = append(groupValues.GroupValue, child)
        return &groupValues.GroupValue[len(groupValues.GroupValue)-1]
    }
    return nil
}

func (groupValues *Srlg_Groups_Group_GroupValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupValues.GroupValue {
        children[groupValues.GroupValue[i].GetSegmentPath()] = &groupValues.GroupValue[i]
    }
    return children
}

func (groupValues *Srlg_Groups_Group_GroupValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupValues *Srlg_Groups_Group_GroupValues) GetBundleName() string { return "cisco_ios_xr" }

func (groupValues *Srlg_Groups_Group_GroupValues) GetYangName() string { return "group-values" }

func (groupValues *Srlg_Groups_Group_GroupValues) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupValues *Srlg_Groups_Group_GroupValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupValues *Srlg_Groups_Group_GroupValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupValues *Srlg_Groups_Group_GroupValues) SetParent(parent types.Entity) { groupValues.parent = parent }

func (groupValues *Srlg_Groups_Group_GroupValues) GetParent() types.Entity { return groupValues.parent }

func (groupValues *Srlg_Groups_Group_GroupValues) GetParentYangName() string { return "group" }

// Srlg_Groups_Group_GroupValues_GroupValue
// Group SRLG values with attribute
type Srlg_Groups_Group_GroupValues_GroupValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG index. The type is interface{} with range:
    // 1..65535.
    SrlgIndex interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetFilter() yfilter.YFilter { return groupValue.YFilter }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) SetFilter(yf yfilter.YFilter) { groupValue.YFilter = yf }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetGoName(yname string) string {
    if yname == "srlg-index" { return "SrlgIndex" }
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "srlg-priority" { return "SrlgPriority" }
    return ""
}

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetSegmentPath() string {
    return "group-value" + "[srlg-index='" + fmt.Sprintf("%v", groupValue.SrlgIndex) + "']"
}

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-index"] = groupValue.SrlgIndex
    leafs["srlg-value"] = groupValue.SrlgValue
    leafs["srlg-priority"] = groupValue.SrlgPriority
    return leafs
}

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetBundleName() string { return "cisco_ios_xr" }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetYangName() string { return "group-value" }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) SetParent(parent types.Entity) { groupValue.parent = parent }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetParent() types.Entity { return groupValue.parent }

func (groupValue *Srlg_Groups_Group_GroupValues_GroupValue) GetParentYangName() string { return "group-values" }

// Srlg_InheritNodes
// Set of inherit nodes configured with SRLG
type Srlg_InheritNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inherit node configurations. The type is slice of
    // Srlg_InheritNodes_InheritNode.
    InheritNode []Srlg_InheritNodes_InheritNode
}

func (inheritNodes *Srlg_InheritNodes) GetFilter() yfilter.YFilter { return inheritNodes.YFilter }

func (inheritNodes *Srlg_InheritNodes) SetFilter(yf yfilter.YFilter) { inheritNodes.YFilter = yf }

func (inheritNodes *Srlg_InheritNodes) GetGoName(yname string) string {
    if yname == "inherit-node" { return "InheritNode" }
    return ""
}

func (inheritNodes *Srlg_InheritNodes) GetSegmentPath() string {
    return "inherit-nodes"
}

func (inheritNodes *Srlg_InheritNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inherit-node" {
        for _, c := range inheritNodes.InheritNode {
            if inheritNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_InheritNodes_InheritNode{}
        inheritNodes.InheritNode = append(inheritNodes.InheritNode, child)
        return &inheritNodes.InheritNode[len(inheritNodes.InheritNode)-1]
    }
    return nil
}

func (inheritNodes *Srlg_InheritNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inheritNodes.InheritNode {
        children[inheritNodes.InheritNode[i].GetSegmentPath()] = &inheritNodes.InheritNode[i]
    }
    return children
}

func (inheritNodes *Srlg_InheritNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inheritNodes *Srlg_InheritNodes) GetBundleName() string { return "cisco_ios_xr" }

func (inheritNodes *Srlg_InheritNodes) GetYangName() string { return "inherit-nodes" }

func (inheritNodes *Srlg_InheritNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritNodes *Srlg_InheritNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritNodes *Srlg_InheritNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritNodes *Srlg_InheritNodes) SetParent(parent types.Entity) { inheritNodes.parent = parent }

func (inheritNodes *Srlg_InheritNodes) GetParent() types.Entity { return inheritNodes.parent }

func (inheritNodes *Srlg_InheritNodes) GetParentYangName() string { return "srlg" }

// Srlg_InheritNodes_InheritNode
// Inherit node configurations
type Srlg_InheritNodes_InheritNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The inherit node name. The type is string with
    // pattern: ((([a-zA-Z0-9_]*\d+)|(\*))/){2}(([a-zA-Z0-9_]*\d+)|(\*)).
    InheritNodeName interface{}

    // Enable SRLG inherit node. The type is interface{}.
    Enable interface{}

    // Set of SRLG values configured under an inherit node.
    InheritNodeValues Srlg_InheritNodes_InheritNode_InheritNodeValues
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetFilter() yfilter.YFilter { return inheritNode.YFilter }

func (inheritNode *Srlg_InheritNodes_InheritNode) SetFilter(yf yfilter.YFilter) { inheritNode.YFilter = yf }

func (inheritNode *Srlg_InheritNodes_InheritNode) GetGoName(yname string) string {
    if yname == "inherit-node-name" { return "InheritNodeName" }
    if yname == "enable" { return "Enable" }
    if yname == "inherit-node-values" { return "InheritNodeValues" }
    return ""
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetSegmentPath() string {
    return "inherit-node" + "[inherit-node-name='" + fmt.Sprintf("%v", inheritNode.InheritNodeName) + "']"
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inherit-node-values" {
        return &inheritNode.InheritNodeValues
    }
    return nil
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["inherit-node-values"] = &inheritNode.InheritNodeValues
    return children
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inherit-node-name"] = inheritNode.InheritNodeName
    leafs["enable"] = inheritNode.Enable
    return leafs
}

func (inheritNode *Srlg_InheritNodes_InheritNode) GetBundleName() string { return "cisco_ios_xr" }

func (inheritNode *Srlg_InheritNodes_InheritNode) GetYangName() string { return "inherit-node" }

func (inheritNode *Srlg_InheritNodes_InheritNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritNode *Srlg_InheritNodes_InheritNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritNode *Srlg_InheritNodes_InheritNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritNode *Srlg_InheritNodes_InheritNode) SetParent(parent types.Entity) { inheritNode.parent = parent }

func (inheritNode *Srlg_InheritNodes_InheritNode) GetParent() types.Entity { return inheritNode.parent }

func (inheritNode *Srlg_InheritNodes_InheritNode) GetParentYangName() string { return "inherit-nodes" }

// Srlg_InheritNodes_InheritNode_InheritNodeValues
// Set of SRLG values configured under an inherit
// node
type Srlg_InheritNodes_InheritNode_InheritNodeValues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inherit node SRLG value with attributes. The type is slice of
    // Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue.
    InheritNodeValue []Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue
}

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetFilter() yfilter.YFilter { return inheritNodeValues.YFilter }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) SetFilter(yf yfilter.YFilter) { inheritNodeValues.YFilter = yf }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetGoName(yname string) string {
    if yname == "inherit-node-value" { return "InheritNodeValue" }
    return ""
}

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetSegmentPath() string {
    return "inherit-node-values"
}

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inherit-node-value" {
        for _, c := range inheritNodeValues.InheritNodeValue {
            if inheritNodeValues.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue{}
        inheritNodeValues.InheritNodeValue = append(inheritNodeValues.InheritNodeValue, child)
        return &inheritNodeValues.InheritNodeValue[len(inheritNodeValues.InheritNodeValue)-1]
    }
    return nil
}

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inheritNodeValues.InheritNodeValue {
        children[inheritNodeValues.InheritNodeValue[i].GetSegmentPath()] = &inheritNodeValues.InheritNodeValue[i]
    }
    return children
}

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetBundleName() string { return "cisco_ios_xr" }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetYangName() string { return "inherit-node-values" }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) SetParent(parent types.Entity) { inheritNodeValues.parent = parent }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetParent() types.Entity { return inheritNodeValues.parent }

func (inheritNodeValues *Srlg_InheritNodes_InheritNode_InheritNodeValues) GetParentYangName() string { return "inherit-node" }

// Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue
// Inherit node SRLG value with attributes
type Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG index. The type is interface{} with range:
    // 1..65535.
    SrlgIndex interface{}

    // SRLG value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    SrlgValue interface{}

    // SRLG priority. The type is SrlgPriority. The default value is default.
    SrlgPriority interface{}
}

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetFilter() yfilter.YFilter { return inheritNodeValue.YFilter }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) SetFilter(yf yfilter.YFilter) { inheritNodeValue.YFilter = yf }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetGoName(yname string) string {
    if yname == "srlg-index" { return "SrlgIndex" }
    if yname == "srlg-value" { return "SrlgValue" }
    if yname == "srlg-priority" { return "SrlgPriority" }
    return ""
}

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetSegmentPath() string {
    return "inherit-node-value" + "[srlg-index='" + fmt.Sprintf("%v", inheritNodeValue.SrlgIndex) + "']"
}

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["srlg-index"] = inheritNodeValue.SrlgIndex
    leafs["srlg-value"] = inheritNodeValue.SrlgValue
    leafs["srlg-priority"] = inheritNodeValue.SrlgPriority
    return leafs
}

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetBundleName() string { return "cisco_ios_xr" }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetYangName() string { return "inherit-node-value" }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) SetParent(parent types.Entity) { inheritNodeValue.parent = parent }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetParent() types.Entity { return inheritNodeValue.parent }

func (inheritNodeValue *Srlg_InheritNodes_InheritNode_InheritNodeValues_InheritNodeValue) GetParentYangName() string { return "inherit-node-values" }

// VrfGroups
// vrf groups
type VrfGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF group configuration. The type is slice of VrfGroups_VrfGroup.
    VrfGroup []VrfGroups_VrfGroup
}

func (vrfGroups *VrfGroups) GetFilter() yfilter.YFilter { return vrfGroups.YFilter }

func (vrfGroups *VrfGroups) SetFilter(yf yfilter.YFilter) { vrfGroups.YFilter = yf }

func (vrfGroups *VrfGroups) GetGoName(yname string) string {
    if yname == "vrf-group" { return "VrfGroup" }
    return ""
}

func (vrfGroups *VrfGroups) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-cfg:vrf-groups"
}

func (vrfGroups *VrfGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-group" {
        for _, c := range vrfGroups.VrfGroup {
            if vrfGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VrfGroups_VrfGroup{}
        vrfGroups.VrfGroup = append(vrfGroups.VrfGroup, child)
        return &vrfGroups.VrfGroup[len(vrfGroups.VrfGroup)-1]
    }
    return nil
}

func (vrfGroups *VrfGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfGroups.VrfGroup {
        children[vrfGroups.VrfGroup[i].GetSegmentPath()] = &vrfGroups.VrfGroup[i]
    }
    return children
}

func (vrfGroups *VrfGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfGroups *VrfGroups) GetBundleName() string { return "cisco_ios_xr" }

func (vrfGroups *VrfGroups) GetYangName() string { return "vrf-groups" }

func (vrfGroups *VrfGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfGroups *VrfGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfGroups *VrfGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfGroups *VrfGroups) SetParent(parent types.Entity) { vrfGroups.parent = parent }

func (vrfGroups *VrfGroups) GetParent() types.Entity { return vrfGroups.parent }

func (vrfGroups *VrfGroups) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-cfg" }

// VrfGroups_VrfGroup
// VRF group configuration
type VrfGroups_VrfGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF group name. The type is string with length:
    // 1..32.
    VrfGroupName interface{}

    // Enable VRF group. The type is interface{}.
    Enable interface{}

    // Set of VRFs configured under a VRF group.
    Vrfs VrfGroups_VrfGroup_Vrfs
}

func (vrfGroup *VrfGroups_VrfGroup) GetFilter() yfilter.YFilter { return vrfGroup.YFilter }

func (vrfGroup *VrfGroups_VrfGroup) SetFilter(yf yfilter.YFilter) { vrfGroup.YFilter = yf }

func (vrfGroup *VrfGroups_VrfGroup) GetGoName(yname string) string {
    if yname == "vrf-group-name" { return "VrfGroupName" }
    if yname == "enable" { return "Enable" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (vrfGroup *VrfGroups_VrfGroup) GetSegmentPath() string {
    return "vrf-group" + "[vrf-group-name='" + fmt.Sprintf("%v", vrfGroup.VrfGroupName) + "']"
}

func (vrfGroup *VrfGroups_VrfGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &vrfGroup.Vrfs
    }
    return nil
}

func (vrfGroup *VrfGroups_VrfGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &vrfGroup.Vrfs
    return children
}

func (vrfGroup *VrfGroups_VrfGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-group-name"] = vrfGroup.VrfGroupName
    leafs["enable"] = vrfGroup.Enable
    return leafs
}

func (vrfGroup *VrfGroups_VrfGroup) GetBundleName() string { return "cisco_ios_xr" }

func (vrfGroup *VrfGroups_VrfGroup) GetYangName() string { return "vrf-group" }

func (vrfGroup *VrfGroups_VrfGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfGroup *VrfGroups_VrfGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfGroup *VrfGroups_VrfGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfGroup *VrfGroups_VrfGroup) SetParent(parent types.Entity) { vrfGroup.parent = parent }

func (vrfGroup *VrfGroups_VrfGroup) GetParent() types.Entity { return vrfGroup.parent }

func (vrfGroup *VrfGroups_VrfGroup) GetParentYangName() string { return "vrf-groups" }

// VrfGroups_VrfGroup_Vrfs
// Set of VRFs configured under a VRF group
type VrfGroups_VrfGroup_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF configuration. The type is slice of VrfGroups_VrfGroup_Vrfs_Vrf.
    Vrf []VrfGroups_VrfGroup_Vrfs_Vrf
}

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *VrfGroups_VrfGroup_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VrfGroups_VrfGroup_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *VrfGroups_VrfGroup_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *VrfGroups_VrfGroup_Vrfs) GetParentYangName() string { return "vrf-group" }

// VrfGroups_VrfGroup_Vrfs_Vrf
// VRF configuration
type VrfGroups_VrfGroup_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}
}

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *VrfGroups_VrfGroup_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// SelectiveVrfDownload
// selective vrf download
type SelectiveVrfDownload struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable selective VRF download. The type is interface{}.
    Disable interface{}
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetFilter() yfilter.YFilter { return selectiveVrfDownload.YFilter }

func (selectiveVrfDownload *SelectiveVrfDownload) SetFilter(yf yfilter.YFilter) { selectiveVrfDownload.YFilter = yf }

func (selectiveVrfDownload *SelectiveVrfDownload) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download"
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = selectiveVrfDownload.Disable
    return leafs
}

func (selectiveVrfDownload *SelectiveVrfDownload) GetBundleName() string { return "cisco_ios_xr" }

func (selectiveVrfDownload *SelectiveVrfDownload) GetYangName() string { return "selective-vrf-download" }

func (selectiveVrfDownload *SelectiveVrfDownload) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectiveVrfDownload *SelectiveVrfDownload) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectiveVrfDownload *SelectiveVrfDownload) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectiveVrfDownload *SelectiveVrfDownload) SetParent(parent types.Entity) { selectiveVrfDownload.parent = parent }

func (selectiveVrfDownload *SelectiveVrfDownload) GetParent() types.Entity { return selectiveVrfDownload.parent }

func (selectiveVrfDownload *SelectiveVrfDownload) GetParentYangName() string { return "Cisco-IOS-XR-infra-rsi-cfg" }

